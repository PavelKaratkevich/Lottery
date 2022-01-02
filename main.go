package main

import (
	"Lottery/app"
)

func main() {
	// Dan - this is OK to have just that in main. Alternative you could have put all the contents of func StartLottery()
	// here and that would be a bit easier to understand how the app works straight from the main.
	app.StartLottery()
}
