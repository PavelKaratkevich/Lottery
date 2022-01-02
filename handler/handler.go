package handler

import (
	"Lottery/domain"
	errorApp "Lottery/err"
	"Lottery/server"
	"encoding/json"
	"net/http"
)

// Dan - this is a minor thing, but, as a famous quote states, naming is hard =) https://martinfowler.com/bliki/TwoHardThings.html
// I could recommend naming it LotteryService and making it an interface for further extensibility and loose coupling considerations.
type LotteryHandlers struct {
	repository server.LotteryRepositoryDb
}

// Dan - I couldn't resist asking - why exactly "ch"?
func (ch *LotteryHandlers) BuyTicket(w http.ResponseWriter, r *http.Request) {
	var request domain.Request
	// Decoding the request into Request struct
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		errorApp.SendError(w, http.StatusBadRequest, err.Error()) // Dan - it could be risky to send back to an end user
		// the raw error output - in some cases it may result in implementation reveal and be exploited by malicious users.
	}
	// Checking if all the fields were filled in

	// Dan - if possible it would be better to avoid nested if statements as they are hard to read and are bug prone.
	// Usually 1 level is enough, a maximum of 2 - kindly consider refactoring this.
	if request.First_name == "" || request.Last_name == "" || request.Id_number == "" {
		errorApp.SendError(w, http.StatusBadRequest, "Please fill out all the required fields")
	} else {
		// Invoking the Ticket Purchase function

		// Dan - this is what IDE tells and I have to agree - "Variable 'error' collides with the builtin interface".
		// Just call it "err", it won't complain =)
		resp, error := ch.repository.BuyLottery(request)
		if error != nil {
			errorApp.SendError(w, error.Code, error.Message)
		} else {
			errorApp.SendSuccess(w, resp)
		}
	}
}

// helper function
// Dan - this one is actually called a constructor and it belongs to the top, next to the LotteryHandlers.
func NewLotteryService(repository server.LotteryRepositoryDb) LotteryHandlers {
	return LotteryHandlers{
		repository,
	}
}

// TODO - make use of the interface
// TODO - fix and move the constructor
// TODO - rename error variable
// TODO - fix the security issue with the raw error data in response
// TODO - add request validation function as you mentioned in the video
// TODO - refactor if
