package main

import (
	"net/http"
)

func (app *application) createUserProfileHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Username           string   `json:"username"`
		Title              string   `json:"title"`
		FullName           []string `json:"full_name"`
		Images             []string `json:"images"`
		PhoneNumber        string   `json:"phone_number"`
		Email              string   `json:"email"`
		DisplayContactInfo []bool   `json:"display_contact_info"`
		Answers            []string `json:"answers"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

}
