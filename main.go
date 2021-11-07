package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// create connection to the mysql database and execute query. dont forget imports
	db, err := sql.Open("mysql", "root:1002@tcp(localhost:3306)/app")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	YearID = 9
	FillStore(db)
	FillUsers()
	ComputeUserNetworksIDs()
	fmt.Println(len(Users[12991].NetworkUserIDS))
	fmt.Println(len(Users[66471].NetworkUserIDS))
	fmt.Println(len(Users[16494].NetworkUserIDS))
}
