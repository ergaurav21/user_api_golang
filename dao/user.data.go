package dao

import (
	"database/sql"
	"fmt"
	"log"
	"user_api/database"
)

type User struct {
	UserName  *string `json:"username"`
	Email     string  `json:"email"`
	FirstName string  `json:"firstname"`
	LastName  string  `json:"lastname"`
}

func GetUser(userName string) (*User, error) {
	row := database.DbConn.QueryRow( `SELECT 
	user_name, 
	email, 
	first_name, 
	last_name   
	FROM tbl_user 
	WHERE user_name = ?`, userName)

	user := &User{}
	err := row.Scan(
		&user.UserName,
		&user.Email,
		&user.FirstName,
		&user.LastName,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		log.Println(err)
		return nil, err
	}
	fmt.Println(user)
	return user, nil
}

func RemoveUser(userName string) error {
	_, err := database.DbConn.Exec(`DELETE FROM tbl_user where user_name = ?`, userName)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func GetUserList() ([]*User, error) {
	results, err := database.DbConn.Query(`SELECT 
	user_name, 
	email, 
	first_name, 
	last_name
	FROM tbl_user`)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	defer results.Close()

	users := make([]*User, 0)
	for results.Next() {
		var user User
		results.Scan(&user.UserName,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			)

		users = append(users, &user)
	}
	return users, nil
}

func UpdateUser(user User) error {

	_, err := database.DbConn.Exec(`UPDATE tbl_user SET 
		email=?, 
		first_name=?, 
		last_name=? 
		WHERE user_name=?`,
		user.Email,
		user.FirstName,
		user.LastName,
		user.UserName)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func InsertUser(user User) error {

	result, err := database.DbConn.Exec(`INSERT INTO tbl_user  
	(user_name, 
	email, 
	first_name, 
	last_name) VALUES (?, ?, ?, ?)`,
		*user.UserName,
		user.Email,
		user.FirstName,
		user.LastName)
	if err != nil {
		log.Println(err.Error())
		return  err
	}
	_, err = result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return  err
	}
	return  nil
}
