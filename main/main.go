// main.go
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/invazions/MSSQL_INTERACTION/models"
	"github.com/invazions/MSSQL_INTERACTION/repositories"
	"github.com/invazions/MSSQL_INTERACTION/services"
)

func main() {
	connString := "sqlserver://username:password@localhost:1433?database=yourdb"
	db, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Error opening database: ", err.Error())
	}
	defer db.Close()

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	// Пример добавления пользователя
	newUser := models.User{Name: "John Doe", Email: "john@example.com"}
	err = userService.CreateUser(newUser)
	if err != nil {
		log.Fatal("Error creating user: ", err.Error())
	}

	// Пример чтения пользователя
	user, err := userService.GetUserByID(1)
	if err != nil {
		log.Fatal("Error getting user: ", err.Error())
	}
	fmt.Printf("Retrieved User: %+v\n", user)

	// Пример удаления пользователя
	err = userService.DeleteUser(1)
	if err != nil {
		log.Fatal("Error deleting user: ", err.Error())
	}
}
