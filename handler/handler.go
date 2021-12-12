package handler

import (
	"Lottery/domain"
	errorApp "Lottery/err"
	"Lottery/server"
	"encoding/json"
	"net/http"
)

type LotteryHandlers struct {
	repository server.LotteryRepositoryDb
}

func (ch *LotteryHandlers) BuyTicket(w http.ResponseWriter, r *http.Request) {
	var request domain.Request
	// Decoding the request into Request struct
	err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			errorApp.SendError(w, http.StatusBadRequest, err.Error())
		}
	// Checking if all the fields were filled in
	if request.First_name == "" || request.Last_name == "" || request.Id_number == "" {
		errorApp.SendError(w, http.StatusBadRequest, "Please fill out all the required fields")
	} else {
	// Invoking the Ticket Purchase function
	resp, error := ch.repository.BuyLottery(request)
		if error != nil {
			errorApp.SendError(w, error.Code, error.Message)
		} else {
			errorApp.SendSuccess(w, resp)
		}
	}
}

// helper function
func NewLotteryService(repository server.LotteryRepositoryDb) LotteryHandlers {
	return LotteryHandlers{
		repository,
	}
}
