package handlers

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "ошибка генерации страницы", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write(file)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) // 10 MB

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "ошибка при получении файла", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка при чтении файла", http.StatusInternalServerError)
		return
	}
	converted := service.Convert(string(fileBytes))

	if err := os.Mkdir("uploads", 0755); err != nil && !errors.Is(err, os.ErrExist) {
		log.Fatal(err)
	}

	fileName := filepath.Join("uploads", time.Now().UTC().String()+filepath.Ext(handler.Filename))
	os.WriteFile(fileName, []byte(converted), 0755)

	io.WriteString(w, converted)
}
