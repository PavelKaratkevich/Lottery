package app

import (
	"Lottery/handler"
	"Lottery/server"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func StartLottery() {
	//set router
	router := mux.NewRouter()
	// establish DB connection
	db := ConnectDB()
	// wiring
	LotteryRepositoryDb := server.NewLotteryRepositoryDb(db)
	ls := handler.NewLotteryService(LotteryRepositoryDb)
	// declare routes
	router.HandleFunc("/", ls.BuyTicket).Methods(http.MethodPost)
	// set environment variables
	address := os.Getenv("ADDRESS_NAME")
	port := os.Getenv("PORT_NAME")
	// listen and serve
	log.Printf("Server is running at %s:%s", address, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

// Dan - It looks like this function is not used outside this package - that's a perfect oportunity to make it unexported =)
// TODO - make the function unexported
func ConnectDB() *sqlx.DB {
	// load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Cannot load .env file:", err)
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

// TODO - add more newlines to increase readability
// TODO - rename variables to match the style of Go (Hint - there are 5)
// TODO - add .env.example file
