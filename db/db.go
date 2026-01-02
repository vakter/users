package db

import (
 "time"

 "github.com/go-pg/pg/v10"
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
