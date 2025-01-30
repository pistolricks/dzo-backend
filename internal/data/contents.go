package data

import (
	"database/sql"
	"fmt"
	"github.com/kolesa-team/go-webp/decoder"
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
	"github.com/pistolricks/go-template-api/internal/validator"
	"image/jpeg"
	"log"
	"os"
	"time"
)

type Content struct {
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

func ValidateContent(v *validator.Validator, content *Content) {
	v.Check(content.Name != "", "name", "is required")
	v.Check(content.Type != "", "type", "is required")
	v.Check(content.Size > 0, "size", "This content doesn't have any data to it")
	v.Check(content.SortOrder > 0, "order", "order must be greater than zero")
}

type ContentModel struct {
	DB *sql.DB
}

func (m ContentModel) EncodeWebP(content *Content) error {
	fmt.Println("Made it to EncodeWebP")
	file, err := os.Open(content.Src)
	if err != nil {
		log.Fatalln(err)
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatalln(err)
	}

	output, err := os.Create("storage/output_encode.webp")

	if err != nil {
		log.Fatal(err)
	}
	//noinspection GoUnhandledErrorResult
	defer output.Close()

	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
	if err != nil {
		log.Fatalln(err)
	}

	if err := webp.Encode(output, img, options); err != nil {
		log.Fatalln(err)
	}

	return nil
}

func (m ContentModel) DecodeWebP(content *Content) error {

	file, err := os.Open(content.Src)
	if err != nil {
		log.Fatalln(err)
	}

	fileName := fmt.Sprintf("./storage/%d.jpg", content.ID)
	output, err := os.Create(fileName)

	if err != nil {
		log.Fatal(err)
	}

	img, err := webp.Decode(file, &decoder.Options{})
	if err != nil {
		log.Fatalln(err)
	}

	if err = jpeg.Encode(output, img, &jpeg.Options{Quality: 75}); err != nil {
		log.Fatalln(err)
	}

	return nil
}
