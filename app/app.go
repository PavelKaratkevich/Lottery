package app

import (
	"Lottery/server"
	_"github.com/go-sql-driver/mysql"
	"Lottery/handler"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
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