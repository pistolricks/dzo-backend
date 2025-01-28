package data

import (
	"context"
	"database/sql"
	"errors"
	"github.com/pistolricks/go-template-api/internal/validator"
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
	UserID             int64     `json:"-"`
}

func ValidateUserProfileEmail(v *validator.Validator, email string) {
	v.Check(email != "", "email", "must be provided")
	v.Check(validator.Matches(email, validator.EmailRX), "email", "must be a valid email address")
}

func ValidateUserProfile(v *validator.Validator, u *UserProfile) {
	v.Check(u.Title != "", "title", "is required")
	v.Check(u.Username != "", "username", "is required")
	v.Check(len(u.FullName) != 0, "full_name", "must not be empty")
	v.Check(len(u.Images) != 0, "images", "must not be empty")
	v.Check(len(u.Answers) != 0, "answers", "must not be empty")
	v.Check(len(u.DisplayContactInfo) != 0, "display_contact_info", "must not be empty")

	v.Check(u.PhoneNumber != "", "phone_number", "must be provided")
	v.Check(u.Email != "", "email", "must be provided")
	v.Check(validator.Matches(u.Email, validator.EmailRX), "email", "must be a valid email address")

}

type UserProfileModel struct {
	DB *sql.DB
}

func (m UserProfileModel) Insert(user *UserProfile) error {
	query := `
	INSERT INTO user_profile (username, title, full_name, images, phone_number, email, display_contact_info, answers, user_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING id, created_at, username`

	args := []any{user.Username, user.Title, user.FullName, user.Images, user.PhoneNumber, user.Email, user.DisplayContactInfo, user.Answers, user.UserID}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.ID, &user.CreatedAt, &user.Username)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		default:
			return err
		}
	}

	return nil
}

func (m UserProfileModel) Update(user *UserProfile) error {
	query := `
UPDATE user_profile
SET username = $1, title = $2, full_name = $3, images = $4, phone_number = $5, email = $6, display_contact_info = $7, answers = $8
WHERE id = $9
RETURNING username`

	args := []any{
		user.Username,
		user.Title,
		user.FullName,
		user.Images,
		user.PhoneNumber,
		user.Email,
		user.DisplayContactInfo,
		user.Answers,
		user.UserID,
		user.ID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.Username)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}
	return nil
}
