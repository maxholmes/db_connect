package member

//function creates a new member
//input: First_Name, Last_Name, email
//output: nil - updates member table with a unique id

// connecting to a PostgreSQL database with Go's database/sql package
import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Member struct {
}

type SqlMember struct {
	Member_id  int
	first_name string
	last_name  string
	Email      string
}

func NewMember() *Member {
	return &Member{}

}

func (m *Member) CreateMember(first_name string, last_name string, email string) error {

	connStr := "user=postgres dbname=first_db password=Password1! host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println("Unable to connect to Database" + err.Error())
		return err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Printf("Failed to ping database!" + err.Error())
		return err
	}
	fmt.Printf("Successfully connected to database!\n")

	sqlStatement := `INSERT INTO member (first_name, last_name, email)
	VALUES ($1, $2, $3)`
	_, err = db.Exec(sqlStatement, first_name, last_name, email)
	if err != nil {
		fmt.Printf("Failed to Insert into Member!" + err.Error())
		return err
	}
	fmt.Println("\nRow inserted successfully!")
	return nil

}

func (m *Member) FetchMemberInfo(first_name string, last_name string) ([]SqlMember, error) {

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
	defer rows.Close()

	var sql_members []SqlMember

	for rows.Next() {
		var member_id int
		var email string

		if err := rows.Scan(&member_id, &email); err != nil {
			fmt.Printf("Failed to Scan Rows!" + err.Error())
			return nil, err
		}
		fmt.Printf("%d is %s\n", member_id, email)
		sql_member := SqlMember{
			Member_id: member_id,
			email:     email,
		}
		sql_members = append(sql_members, sql_member)
	}
	if err := rows.Err(); err != nil {
		fmt.Printf("Failed to Iterate Rows!" + err.Error())
		return nil, err
	}
	return sql_members, nil
}
