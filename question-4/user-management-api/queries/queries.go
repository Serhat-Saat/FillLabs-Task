package queries

// providing a centralized management for SQL queries
const (
	CreateUser = "INSERT INTO users (userName, userEmail, userPhone) VALUES (?, ?, ?)"

	GetAllUsers = "SELECT id, userName, userEmail, userPhone FROM users"

	GetUserByID = "SELECT id, userName, userEmail, userPhone FROM users WHERE id = ?"

	UpdateUser = "UPDATE users SET userName = ?, userEmail = ?, userPhone = ? WHERE id = ?"

	DeleteUser = "DELETE FROM users WHERE id = ?"
)
