package handler

// import (
// 	errorApp "Lottery/err"
// 	"Lottery/lotterypb"
// 	"Lottery/server"
// 	"context"
// 	"encoding/json"
// 	"log"
// 	"net/http"

// 	"google.golang.org/grpc/status"
// )

// type LotteryHandlers struct {
// 	repository server.LotteryRepositoryDb
// }

// func (ch *LotteryHandlers) BuyTicket(r *lotterypb.Request) {
// 	var request lotterypb.Request
// 	// Decoding the request into Request struct
// 	err := json.NewDecoder(r.Body).Decode(&request)
// 		if err != nil {
// 			errorApp.SendError(w, http.StatusBadRequest, err.Error())
// 		}
// 	// Checking if all the fields were filled in
// 	if request.FirstName == "" || request.LastName == "" || request.IdNumber == "" {
// 		errorApp.SendError(w, http.StatusBadRequest, "Please fill out all the required fields")
// 	} else {
// 	// Invoking the Ticket Purchase function
// 	resp, error := ch.repository.BuyLotteryTicket(context.Background(), &request)
		
	
// 	if error != nil {
// 		status, ok := status.FromError(err)
// 		if ok == true {
// 			log.Printf("Error message from server: %v\n", status.Message())
// 			log.Println(status.Code())
// 			return
// 		} else {
// 			log.Fatalf("Some big trouble: %v", err)
// 			return
// 		}
// 	}
// 	}
// }

// // helper function
// func NewLotteryService(repository server.LotteryRepositoryDb) LotteryHandlers {
// 	return LotteryHandlers{
// 		repository,
// 	}
// }