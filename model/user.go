package model

import (
	"context"
	"fmt"
	"os"

	"github.com/s-yuk/go-bank-api/database"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u User) GetUser() ([]User, error) {
	conn := database.ConnectDatabase()

	rows, err := conn.Query(context.Background(), "SELECT * FROM users")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	var users []User

	for rows.Next() {

		if err := rows.Scan(&u.ID, u.Username, u.Email, u.Password); err != nil {
			return []User{}, err
		}

		user := User{
			ID:       u.ID,
			Username: u.Username,
			Email:    u.Email,
			Password: u.Password,
		}

		users = append(users, user)
	}

	return users, nil
}
