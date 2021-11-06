package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type UserTitle struct {
	UserID    uint32
	TitleID   uint8
	YearID    uint8
	IsDefault bool
}

func main() {
	// create connection to the mysql database and execute query. dont forget imports
	db, err := sql.Open("mysql", "root:1002@tcp(localhost:3306)/app")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// query the database
	results, err := db.Query("SELECT user_id, title_id, year_id, is_default FROM user_title WHERE year_id = 10")
	if err != nil {
		panic(err.Error())
	}

	// print results
	for results.Next() {
		// scan the result and assign to UserTitle
		var userTitle UserTitle
		err = results.Scan(&userTitle.UserID, &userTitle.TitleID, &userTitle.YearID, &userTitle.IsDefault)
		if err != nil {
			panic(err.Error())
		}
		// print the result
		// fmt.Print(userTitle)
	}

	// close the mysql database connection
	defer db.Close()
}
