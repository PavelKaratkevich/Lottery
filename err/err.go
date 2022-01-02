package errorApp

import (
	"encoding/json"
	"net/http"
)

// helper functions
func SendSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err) // Dan - this might be a bit of an overreaction - panicking would shut down your app and
		// this is not something you would like to handle often =D
	}
	// Dan - technically this return is not entirely necessary, because the functions returns automatically at the end of its body
	// anyway and you do not have any values to return.
	return
}
func SendError(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		// Dan - same as above.
		panic(err)
	}
	// Dan - same as above.
	return
}

// Dan - it appears that the two functions are quite similar, and we could make this package DRYer.
// You could have one unexported function to send any status and use two wrapper-functions named as above to send success/error.
// Don't hesitate to ask me for more clarifications if necessary.

// TODO - avoid panic
// TODO - remove returns
// TODO - refactor to make the code DRY.
