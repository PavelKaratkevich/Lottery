package server

import (
	"Lottery/domain"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strings"
)

type LotteryRepositoryDb struct {
	client *sqlx.DB
}

func (l LotteryRepositoryDb) BuyLottery(request domain.Request) (*domain.Response, *domain.Error) {
	// Creating struct for a Response
	var output domain.Response
	// Creating struct for an existing ticket holder
	var existingCustomer domain.ExistingCustomer
	// Checking if the ticket was already purchased by the ID holder
	/* we can also do this check by returning 1062 error code from the MySQL database, however, it will make the DB skip
	some of the auto-incremented fields 'id of the ticket'*/

	// Dan - it's great that you are writing your own SQL queries - this is a sign of a hardcore developer \ -_-/
	// However there are a couple of things I have to mention:
	// 1. It's best practise to use a separate package for repository logics only. You definitely should use prepared
	// statements there as well - this would safeguard you from a common vulnerability called SQL Injection - read more hear - https://www.w3schools.com/sql/sql_injection.asp
	error := l.client.Get(&existingCustomer, "Select first_name, last_name, id_number from Lottery where id_number = ?", request.Id_number)
	// If no ticket holder with such ID were identified, then we invoke an INSERT function

	// Dan - if possible it would be better to avoid nested if statements as they are hard to read and are bug prone.
	// Usually 1 level is enough, a maximum of 2 - kindly consider refactoring this.
	if error == sql.ErrNoRows {
		res, err := l.client.Exec("INSERT INTO Lottery (first_name, last_name, id_number) VALUES (?, ?, ?)", request.First_name, request.Last_name, request.Id_number)
		if err != nil {
			if strings.Contains(err.Error(), "Error 1644") {
				return nil, &domain.Error{
					Code:    http.StatusGone,
					Message: "Tickets were sold out",
				}
			} else {
				return nil, &domain.Error{
					Code:    http.StatusInternalServerError,
					Message: "Unknown server error",
				}
			}
		}
		// Getting the last inserted ticket ID
		// Dan - you can name it "err" - it will be rewritten anyway, no worries
		insertID, err1 := res.LastInsertId()
		if err1 != nil {
			return nil, &domain.Error{
				Code:    http.StatusInternalServerError,
				Message: "Error while retrieving ID of a purchased ticket",
			}
		}
		output = domain.Response{Ticket_id: int(insertID)}
	} else {
		// If holder of ticket was identified, we send to him/her the notification
		return nil, &domain.Error{
			Code: http.StatusForbidden,
			// Dan - looks like another avoidable implementation reveal - it's recommended to always assume that your use
			// is a nasty hacker that wants to steal your data and do God knows what with your web-site/app. Please consider
			// changing the message to avoid the reveal.
			Message: fmt.Sprintf("The user with ID Number %v already purchased a ticket", existingCustomer.Id_number),
		}
	}
	return &output, nil
}

func NewLotteryRepositoryDb(dbClient *sqlx.DB) LotteryRepositoryDb {
	return LotteryRepositoryDb{
		dbClient,
	}
}

// TODO - fix and move the constructor
// TODO - rename error variable
// TODO - fix SQL injection - create a separate DB logics package and use prepared statements http://jmoiron.github.io/sqlx/
// TODO - fix reveling implementation
// TODO - refactor if
