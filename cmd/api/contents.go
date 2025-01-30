package main

import (
	"fmt"
	"github.com/pistolricks/go-template-api/internal/data"
	"github.com/pistolricks/go-template-api/internal/validator"
	"net/http"
	"time"
)

func (app *application) uploadImageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	var input struct {
		ID        string    `json:"id"`
		CreatedAt time.Time `json:"-"`
		Name      string    `json:"name"`
		Src       string    `json:"src"`
		Type      string    `json:"type"`
		Size      float32   `json:"size"`
		Width     float32   `json:"width"`
		Height    float32   `json:"height"`
		SortOrder int16     `json:"sort_order"`
		UserID    string    `json:"user_id"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	fmt.Println("Past readJSON")

	content := &data.Content{
		ID:        input.ID,
		CreatedAt: input.CreatedAt,
		Name:      input.Name,
		Src:       input.Src,
		Type:      input.Type,
		Size:      input.Size,
		Width:     input.Width,
		Height:    input.Height,
		SortOrder: input.SortOrder,
		UserID:    input.UserID,
	}

	v := validator.New()

	if data.ValidateContent(v, content); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Contents.EncodeWebP(content)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/contents/%s", content.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"content": content}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
