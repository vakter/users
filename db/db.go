package db

import (
 "time"
 "github.com/go-pg/pg/v10"
 "github.com/vakter/libs/dbhandler"

)

// User represents the users table in the database
type User struct {
 ID           int       pg:"id,pk"                // Primary key
 ChatID       int64     pg:"chat_id,unique,notnull" // Unique chat ID
 TelegramID   int64     pg:"telegram_id,unique,notnull" // Unique Telegram ID
 FirstName    string    pg:"first_name"           // First name of the user
 LastName     string    pg:"last_name"            // Last name of the user
 LanguageCode string    pg:"language_code"        // Language code of the user
 Username     string    pg:"username"             // Username of the user
 State        string    pg:"state"                // State of the user
 CreationTime time.Time pg:"creation_time,default:now()" // Timestamp of creation
}

// UsersDBHandler extends DBHandler with user-specific methods
type UsersDBHandler struct {
 dbhandler.DBHandler
}

// CreateUser inserts a new user into the database
func (udb *UsersDBHandler) CreateUser(chatID, telegramID int64, firstName, lastName, languageCode, username string) error {
 user := &User{
  ChatID:       chatID,
  TelegramID:   telegramID,
  FirstName:    firstName,
  LastName:     lastName,
  LanguageCode: languageCode,
  Username:     username,
  State:        "active", // Default state
 }
 _, err := udb.Conn.Model(user).Insert()
 return err
}

// GetUser retrieves a user by chatID
func (udb *UsersDBHandler) GetUser(chatID int64) (*User, error) {
 user := &User{}
 err := udb.Conn.Model(user).Where("chat_id = ?", chatID).Select()
 if err != nil {
  return nil, err
 }
 return user, nil
}

// UpdateUser updates an existing user's information
func (udb *UsersDBHandler) UpdateUser(chatID int64, firstName, lastName, languageCode, username string) error {
 user := &User{
  ChatID:       chatID,
  FirstName:    firstName,
  LastName:     lastName,
  LanguageCode: languageCode,
  Username:     username,
 }
 _, err := udb.Conn.Model(user).Where("chat_id = ?", chatID).Update()
 return err
}

// ListUsers retrieves all users from the database
func (udb *UsersDBHandler) ListUsers() ([]User, error) {
 var users []User
 err := udb.Conn.Model(&users).Select()
 if err != nil {
  return nil, err
 }
 return users, nil
}

// DeleteUser removes a user from the database by chatID
func (udb *UsersDBHandler) DeleteUser(chatID int64) error {
 user := &User{ChatID: chatID}
 _, err := udb.Conn.Model(user).Where("chat_id = ?", chatID).Delete()
 return err
}
