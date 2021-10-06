package models

import (
	"database/sql"
	"echo-rest/db"
	"echo-rest/helpers"
	"fmt"
)

type Account struct {
	idAccount int    `json:"idaccount"`
	Email     string `json:"email"`
}

func CheckLogin(email, password string) (bool, error) {
	var obj Account
	var pwd string

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM account WHERE email = ?"

	err := con.QueryRow(sqlStatement, email).Scan(
		&obj.idAccount, &obj.Email, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Println("Email not found")
		return false, err

	}

	if err != nil {
		fmt.Println("Query error")
		return false, err
	}

	match, err := helpers.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Password doesn't match")
		return false, err
	}

	return true, nil

}
