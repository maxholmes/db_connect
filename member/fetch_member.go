package member

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type GetMemberInfo struct {
}

func FetchMember() *GetMemberInfo {
	return &GetMemberInfo{}

}

const (
	connStr = "user=postgres dbname=first_db password=**** host=localhost sslmode=disable"
)

// input first and last name
//return email and ID
func (g *GetMemberInfo) FetchMemberInfo(first_name string, last_name string) (*sql.Rows, error) {

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println("Unable to connect to Database" + err.Error())
		return nil, err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Printf("Failed to ping database!" + err.Error())
		return nil, err
	}
	fmt.Printf("Successfully connected to database!\n")

	sqlStatement := `select member_id, email from member where first_name = $1 and last_name = $2`
	rows, err := db.Query(sqlStatement, first_name, last_name)

	if err != nil {
		fmt.Printf("Failed to Query Row!" + err.Error())
		return nil, err

	}

	fmt.Println("Fetched Row!")

	return rows, nil
}

// age := 27
// rows, err := db.Query("SELECT name FROM users WHERE age=?", age)
// if err != nil {
//         log.Fatal(err)
// }
// defer rows.Close()
// for rows.Next() {
//         var name string
//         if err := rows.Scan(&name); err != nil {
//                 log.Fatal(err)
//         }
//         fmt.Printf("%s is %d\n", name, age)
// }
// if err := rows.Err(); err != nil {
//         log.Fatal(err)
// }
