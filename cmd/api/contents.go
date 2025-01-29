package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	file, header, _ := r.FormFile("file")
	defer file.Close()

	// create a destination file
	dst, _ := os.Create(filepath.Join("./", header.Filename))
	defer dst.Close()

	// upload the file to destination path
	nb_bytes, _ := io.Copy(dst, file)

	fmt.Println("File uploaded successfully")
}
