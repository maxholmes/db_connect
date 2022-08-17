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

	sqlStatement := `INSERT INTO member ( first_name, last_name, email)
	VALUES ($1, $2, $3)`
	_, err = db.Exec(sqlStatement, first_name, last_name, email)
	if err != nil {
		fmt.Printf("Failed to Insert into Member!" + err.Error())
		return err
	}
	fmt.Println("\nRow inserted successfully!")
	return nil

}
