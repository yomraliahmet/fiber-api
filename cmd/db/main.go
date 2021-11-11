package main

import (
	"fmt"
	"os"

	"github.com/yomraliahmet/fiber-api/database"
	"github.com/yomraliahmet/fiber-api/models"
)

func main() {

	args := os.Args[1:]

	var errMessage string = "You did not enter a valid command!"

	if len(args) > 0 {

		argName := args[0]

		if argName == "reset" {

			database.ConnectDb()
			database.Database.Db.Migrator().DropTable(&models.User{})
			database.Database.Db.Migrator().DropTable(&models.Product{})
			database.Database.Db.Migrator().DropTable(&models.Order{})

			fmt.Println("Database reset successfully.")
		} else {
			fmt.Println(errMessage)
		}

	} else {
		fmt.Println(errMessage)
	}
}
