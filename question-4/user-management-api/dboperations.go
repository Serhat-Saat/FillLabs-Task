package main

import (
	"database/sql"
	"fmt"
	"user-management-api/errors"
	"user-management-api/queries"
)

// Add a new user
func CreateUser(user *User) error {
	query := queries.CreateUser
	result, err := db.Exec(query, user.UserName, user.UserEmail, user.UserPhone)
	if err != nil {
		return fmt.Errorf(errors.ErrCreateUser)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf(errors.ErrLastInsertId)
	}
	user.ID = int(id)

	return nil
}

// fetch a user with a specific ID
func GetUserByID(id int) (*User, error) {
	query := queries.GetUserByID
	row := db.QueryRow(query, id)
	var user User
	err := row.Scan(&user.ID, &user.UserName, &user.UserEmail, &user.UserPhone)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf(errors.ErrNotFound)
		}
		return nil, fmt.Errorf(errors.ErrReadUsers)
	}

	return &user, nil
}

// Fetch all users
func GetAllUsers() ([]User, error) {
	query := queries.GetAllUsers
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf(errors.ErrGetUsers)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.UserName, &user.UserEmail, &user.UserPhone); err != nil {
			return nil, fmt.Errorf(errors.ErrReadUsers)
		}
		users = append(users, user)
	}
	return users, nil
}

// Update user information
func UpdateUser(user *User) error {
	query := queries.UpdateUser
	result, err := db.Exec(query, user.UserName, user.UserEmail, user.UserPhone, user.ID)
	if err != nil {
		return fmt.Errorf(errors.ErrUpdateUsers)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf(errors.ErrRowsAffected)
	}
	if rowsAffected == 0 {
		return fmt.Errorf(errors.ErrNotFound)
	}
	return nil
}

// Delete a user
func DeleteUser(id int) error {
	query := queries.DeleteUser
	result, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf(errors.ErrDeleteUser)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf(errors.ErrRowsAffected)
	}
	if rowsAffected == 0 {
		return fmt.Errorf(errors.ErrNotFound)
	}
	return nil
}
