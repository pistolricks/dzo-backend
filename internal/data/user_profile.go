package data

import (
	"database/sql"
	"time"
)

type UserProfile struct {
	ID                 int64     `json:"id"`
	CreatedAt          time.Time `json:"created_at"`
	Username           string    `json:"username"`
	Title              string    `json:"title"`
	FullName           []string  `json:"full_name"`
	Images             []string  `json:"images"`
	PhoneNumber        string    `json:"phone_number"`
	Email              string    `json:"email"`
	DisplayContactInfo []bool    `json:"display_contact_info"`
	Answers            []string  `json:"answers"`
}

type UserProfileModel struct {
	DB *sql.DB
}
