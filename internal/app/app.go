package app

import (
	"fmt"
	"github.com/DarabiMohammad/backend-app/internal/database"
)

func StartApp() {

	var err = database.Connect()

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database")
}
