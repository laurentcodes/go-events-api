package models

import (
	"errors"

	"example.com/events-api/db"
	"example.com/events-api/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

func (u User) Save() error {
	query := `
    INSERT INTO users (email, password) 
    VALUES (?, ?)
  `

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId

	return err
}

func (u *User) ValidateCredentials() error {
	query := `
		SELECT id, email, password 
		FROM users 
		WHERE email = ?
	`

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	var id int64

	err := row.Scan(&id, &u.Email, &retrievedPassword)

	if err != nil {
		return errors.New("invalid credentials")
	}

	passwordValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordValid {
		return errors.New("invalid credentials")
	}

	u.ID = id

	return nil
}
