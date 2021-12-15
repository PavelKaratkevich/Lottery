package main

import (
	"Lottery/lotterypb"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type LotteryRepositoryDb struct {
	client *sqlx.DB
}

func main() {
	db := ConnectDB()
	startServer(db)
}

func (l LotteryRepositoryDb) BuyLotteryTicket(ctx context.Context, req *lotterypb.Request) (*lotterypb.Response, error) {
	var output lotterypb.Response
	var ExistingCustomer lotterypb.ExistingCustomer
	if req.FirstName == "" || req.LastName == "" || req.IdNumber == "" {
		return nil, status.Errorf(codes.Internal, "All fields must be filled in")
	}
	error := l.client.Get(&ExistingCustomer, "Select first_name, last_name, id_number from Lottery where id_number = ?", req.IdNumber)
	log.Printf("Error: %v", error)
	switch {
	case error == sql.ErrNoRows:
		res, err := l.client.Exec("INSERT INTO Lottery (first_name, last_name, id_number) VALUES (?, ?, ?)", req.FirstName, req.LastName, req.IdNumber)
		if err != nil {
			if strings.Contains(err.Error(), "Error 1644") {
				return nil, status.Errorf(codes.ResourceExhausted, "Tickets were sold out")
			} else {
				return nil, status.Errorf(codes.Internal, "Unknown server error")
			}
		}
		insertID, err1 := res.LastInsertId()
		if err1 != nil {
			return nil, status.Errorf(codes.Internal, "Error while retrieving ID of a purchased ticket")
		}
		output = lotterypb.Response{TicketId: int32(insertID)}
	case error == nil:
		return nil, status.Errorf(codes.AlreadyExists, fmt.Sprintf("The user with ID Number %v already purchased a ticket", ExistingCustomer.IdNumber))
	case strings.Contains(error.Error(), "Error 1146"):
		return nil, status.Errorf(codes.Internal, "The table not found")
	}
	return &output, nil
}

func NewLotteryRepositoryDb(dbClient *sqlx.DB) LotteryRepositoryDb {
	return LotteryRepositoryDb{
		dbClient,
	}
}

func ConnectDB() *sqlx.DB {
	// load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// get environment variables
	db_name := os.Getenv("DB_NAME")
	db_port := os.Getenv("DB_PORT")
	db_address := os.Getenv("DB_ADDRESS")
	db_pswd := os.Getenv("DB_PSWD")
	dataSource := fmt.Sprintf("root:%s@tcp(%s:%s)/%s", db_pswd, db_address, db_port, db_name)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil || client == nil {
		log.Fatal("Error while opening DB: ", err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}

func startServer(db *sqlx.DB) {
	// Set flags if our code crashes - it will show the line where code crahsed
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// create SSL server
	s := grpc.NewServer()
	// register service
	lss := NewLotteryRepositoryDb(db)
	lotterypb.RegisterLotteryServiceServer(s, lss)
	reflection.Register(s)
	// establish connection
	ls, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatalf("Error while listening to the network: %v", err)
	}
	// binding
	go func() {
		fmt.Println("Server is running...")
		if err := s.Serve(ls); err != nil {
			log.Fatalf("Error while serving the connection: %v", err)
		}
	}()
	// Waiting for CTRL C
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	// Block until a signal is received
	<-ch
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("Closing the listener")
	ls.Close()
	fmt.Println("Stopping DB connection")
	db.Close()
	fmt.Println("Program has been completed")
}
