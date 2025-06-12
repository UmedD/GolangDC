package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var db *sqlx.DB

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func main() {
	dsn := "host=localhost port=5432 user=postgres password=2021 dbname=postgres sslmode=disable"
	var err error
	db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalln("Connection error:", err)
	}
	fmt.Println("Connected to PostgreSQL!")

	sqlQuery := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100),
			email VARCHAR(100),
			age INT
		);`
	db.MustExec(sqlQuery)

	us := User{
		Name:  "Umed",
		Email: "um@gmail.com",
		Age:   29,
	}

	err = insertUser(us)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	users, err := getAllUsers()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%#+v", users)

	err = updatePrice("Umed", 30)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	us, err = getProductByID(1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%#+v", us)

	err = deleteUser("Umed")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}

func insertUser(u User) error {
	_, err := db.Exec(`
			INSERT INTO users (name, email, age)
			VALUES ($1, $2, $3) RETURNING id`, u.Name, u.Email, u.Age)
	return err
}

func getAllUsers() ([]User, error) {
	var users []User
	err := db.Select(&users, "SELECT name, email, age FROM users")
	return users, err
}

func getProductByID(id int) (User, error) {
	p := User{}
	err := db.Get(&p, "SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		return User{}, err
	}

	return p, nil
}

func updatePrice(name string, newAge int) error {
	_, err := db.Exec("UPDATE users SET age = $1 WHERE name = $2", newAge, name)
	return err
}

func deleteUser(name string) error {
	_, err := db.Exec("DELETE FROM users WHERE name = $1", name)
	return err
}

func GetConn() *sqlx.DB {
	return db
}
