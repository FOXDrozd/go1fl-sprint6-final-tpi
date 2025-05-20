package handlers

import (
	"io"
	"net/http"
	"os"
	"time"

	"go1fl-sprint6-final/internal/service"
)


func GetHtmlFormat(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		http.Error(w, "No correction method", http.StatusBadRequest)
	}

    http.ServeFile(w, r, "./index.html")
}

func UploadFile(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w, "No correction method", http.StatusBadRequest)
	}

	fileUploaded, _, err := r.FormFile("myFile")
  	defer fileUploaded.Close()


	if err != nil {
		http.Error(w, "Error get file", http.StatusBadRequest)
		return
	}

	data, err := io.ReadAll(fileUploaded)

	if err != nil {
		http.Error(w, "Not read file", http.StatusInternalServerError)
		return
	}

	str, err := service.FormatData(string(data)) 

	if err != nil {
		http.Error(w,"Check data: %s",  http.StatusBadRequest)
		return 
	}

	file, err := os.OpenFile(time.Now().UTC().Format("02-01_03-04-05"), os.O_CREATE | os.O_APPEND | os.O_RDWR, 0755)

	if err != nil {
		http.Error(w, "Error read file", http.StatusBadRequest)
		return
	}


	defer file.Close()

	_, err = file.WriteString(str) 

	if err != nil {
		http.Error(w,"Error write file",  http.StatusBadRequest)
		return
	}
}