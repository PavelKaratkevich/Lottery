package domain

type Request struct {
	First_name string
	Last_name  string
	// Dan - can we trust a user to pass us his/her id? What if a user wants to exploit our system?
	Id_number string
}

type Response struct {
	Ticket_id int
}

type Error struct {
	Code    int
	Message string
}

type ExistingCustomer struct {
	First_name string
	Last_name  string
	Id_number  string
}

// TODO - rename fields to match the style of Go
// TODO - think about the use of ExistingCustomer struct - it has the same fields ar request. Maybe make a User struct and reuse it?
// TODO - think about the security issue with the user ID. Fix, if possible.
