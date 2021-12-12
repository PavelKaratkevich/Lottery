package server

import (
	"Lottery/domain"
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"github.com/jmoiron/sqlx"
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
	error := l.client.Get(&existingCustomer, "Select first_name, last_name, id_number from Lottery where id_number = ?", request.Id_number)
// If no ticket holder with such ID were identified, then we invoke an INSERT function		
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
				Code:    http.StatusForbidden,
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
