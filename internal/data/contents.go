package data

import (
	"database/sql"
	"fmt"
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
	"image/jpeg"
	"log"
	"os"
	"time"
)

type Content struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Name      string    `json:"name"`
	Src       string    `json:"src"`
	Type      string    `json:"type"`
	Size      float32   `json:"size"`
	Order     int16     `json:"order"`
	UserID    int64     `json:"-"`
}

type ContentModel struct {
	DB *sql.DB
}

func (m ContentModel) EncodeWebP(content *Content) error {
	file, err := os.Open(content.Src)
	if err != nil {
		log.Fatalln(err)
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatalln(err)
	}

	fileName := fmt.Sprintf("./storage/image-%d-%d.webp", content.ID, content.Order)

	output, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
	if err != nil {
		log.Fatalln(err)
	}

	if err := webp.Encode(output, img, options); err != nil {
		log.Fatalln(err)
	}
	if err != nil {
		return err
	}
	return nil
}
