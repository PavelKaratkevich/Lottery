package domain

type Request struct {
	First_name string
	Last_name string
	Id_number string
}

type Response struct {
	Ticket_id int
}

type Error struct {
	Code int
	Message string
}

type ExistingCustomer struct {
	First_name string
	Last_name string
	Id_number string
}