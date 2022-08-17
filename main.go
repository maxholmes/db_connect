package main

import (
	"db_connect/member"
	"fmt"
)

func main() {

	fmt.Println("welcome")

	var member = member.NewMember()
	member.CreateMember("TreeJ", "TV", "TreeJTVgmail.com")

	// fetched_members, err := member.FetchMemberInfo("Tony", "TheTiger")
	// if err == nil {
	// 	for _, fetched_member := range fetched_members {
	// 		fmt.Printf("Getched Member!" + fetched_member.Email)
	// 	}
	// }

	/*
	   variables required for connection string: connStr

	   user= (using default user for postgres database)
	   dbname= (using default database that comes with postgres)
	   password = (password used during initial setup)
	   host = (hostname or IP Address of server)
	   sslmode = (must be set to disabled unless using SSL)
	*/

	// connect to db
	// ************create an insert statement for each table************

	// connStr := "user=postgres dbname=first_db password=Password1! host=localhost sslmode=disable"
	// db, err := sql.Open("postgres", connStr)
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("\nSuccessfully connected to database!\n")

	// // insert a row
	// sqlStatement := `INSERT INTO member ( first_name, last_name, email)
	// VALUES ($1, $2, $3)`
	// _, err = db.Exec(sqlStatement, "Cole ", "Brittner", "cole.brittner@route.com")
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("\nRow inserted successfully!")
	// }

	// // update a row in table
	// sqlStatementUpdt1 := `
	//     UPDATE member
	//     SET first_name = $2 , last_name = $3
	// 	where member_id = $1`
	// res, err := db.Exec(sqlStatementUpdt1, 1, "Max", "Holmes")
	// if err != nil {
	// 	panic(err)
	// }
	// count, err := res.RowsAffected()

	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("rows updated: %v\n", count)

}
