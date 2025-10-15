package models

import (
	"minha-primeira-api/internal/models/database"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:password`
}

func GetAllUsers() ([]User, error) {
	rows, err := database.DB.Query("SELECT id, name, age FROM USERS")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func InsertUsers(u *User) error {
	return database.DB.QueryRow(
		"INSERT INTO users (name, age) VALUES ($1,$2) RETURNING id",
		u.Name, u.Age,
	).Scan(&u.ID)
}

func DeleteUsersById(id int) error {
	_, err := database.DB.Exec("DELETE FROM USERS WHERE id=$1", id)
	return err
}

func UpdateUser(u *User) error {
	_, err := database.DB.Exec("UPDATE users SET name=$1, age=$2, id=$3", u.Name, u.Age, u.ID)
	return err
}
