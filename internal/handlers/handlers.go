package handlers

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method now allowed", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "./index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, "ошибка выделения памяти", http.StatusInternalServerError)
		return
	}

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
		http.Error(w, "внутренняя ошибка", http.StatusInternalServerError)
		return
	}

	fileName := filepath.Join("uploads", time.Now().UTC().String()+filepath.Ext(handler.Filename))
	os.WriteFile(fileName, []byte(converted), 0755)

	if _, err = io.WriteString(w, converted); err != nil {
		http.Error(w, "внутренняя ошибка", http.StatusInternalServerError)
	}
}
