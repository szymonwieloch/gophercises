package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Database phone number normalizer")
	os.Remove("./numbers.db")

	db, err := sql.Open("sqlite3", "./numbers.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = createDB(db)
	if err != nil {
		log.Fatal(err)
	}

	phoneNumbers, err := getAllPhoneNumbers(db)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(phoneNumbers)

	for idx := range phoneNumbers {
		phoneNumbers[idx].number = normalize(phoneNumbers[idx].number)
	}

	unique, duplicates := findDuplicates(phoneNumbers)

	err = updateNumbers(db, unique)
	if err != nil {
		log.Fatal(err)
	}
	err = deleteDuplicates(db, duplicates)
	if err != nil {
		log.Fatal(err)
	}
	// Print the outcome of the cleanup
	phoneNumbers, err = getAllPhoneNumbers(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Unique and normalized phone numbers:")
	for _, ph := range phoneNumbers {
		fmt.Println("id=", ph.id, ",number=", ph.number)
	}
}

// initial numbers to be cleaned by the rest of the application
var initialNumbers []string = []string{
	"1234567890",
	"123 456 7891",
	"(123) 456 7892",
	"(123) 456-7893",
	"123-456-7894",
	"123-456-7890",
	"1234567892",
	"(123)456-7892",
}

// Crates fake DB with hardcoded list of numbers that need normalization
func createDB(db *sql.DB) error {
	createTable := `
	create table numbers (
		id integer not null primary key autoincrement,
		number text not null
	);
	`
	_, err := db.Exec(createTable)
	if err != nil {
		return err
	}

	insertNumbers := `
		insert into numbers (number) values($1)
	`

	for _, num := range initialNumbers {
		_, err = db.Exec(insertNumbers, num)
		if err != nil {
			return err
		}
	}
	return nil
}

type PhoneNumber struct {
	id     int32
	number string
}

// Gets all phone numbers from the DB
func getAllPhoneNumbers(db *sql.DB) ([]PhoneNumber, error) {
	query := `select id, number from numbers`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := []PhoneNumber{}
	for rows.Next() {
		var id int32
		var number string
		err = rows.Scan(&id, &number)
		if err != nil {
			return nil, err
		}
		result = append(result, PhoneNumber{id: id, number: number})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return result, err
}

// Finds duplicates in the provided list.
// Returns unique phone numbers and the list of duplicated IDs
func findDuplicates(numbers []PhoneNumber) ([]PhoneNumber, []int32) {
	unique := []PhoneNumber{}
	duplicates := []int32{}
	known := map[string]bool{}
	for _, pn := range numbers {
		if known[pn.number] {
			duplicates = append(duplicates, pn.id)
		} else {
			known[pn.number] = true
			unique = append(unique, pn)
		}
	}
	return unique, duplicates
}

func deleteDuplicates(db *sql.DB, duplicates []int32) error {
	if len(duplicates) == 0 {
		return nil
	}
	stringIDs := []string{}
	for _, dup := range duplicates {
		stringIDs = append(stringIDs, fmt.Sprint(dup))
	}
	query := fmt.Sprint("delete from numbers where id in(", strings.Join(stringIDs, ", "), ")")
	_, err := db.Exec(query)
	return err
}

func updateNumbers(db *sql.DB, numbers []PhoneNumber) error {
	query := `
		update numbers set number = $1 where id = $2
	`
	for _, num := range numbers {
		_, err := db.Exec(query, num.number, num.id)
		if err != nil {
			return err
		}
	}

	return nil //TODO
}
