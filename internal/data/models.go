package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Vendors     VendorModel
	Permissions PermissionModel
	Tokens      TokenModel
	Users       UserModel
	UserProfile UserProfileModel
	Contents    ContentModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Vendors:     VendorModel{DB: db},
		Permissions: PermissionModel{DB: db},
		Tokens:      TokenModel{DB: db},
		Users:       UserModel{DB: db},
		UserProfile: UserProfileModel{DB: db},
		Contents:    ContentModel{DB: db},
	}
}
