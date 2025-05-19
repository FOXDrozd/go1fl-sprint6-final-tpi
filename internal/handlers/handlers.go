package handlers

import (
	"fmt"
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
	}

	str, err := service.FormatData(string(data)) 

	if err != nil {
		fmt.Printf("Check data: %s", err.Error())
		return 
	}

	file, err := os.OpenFile(time.Now().UTC().Format("02-01-2006 03-04-05"), os.O_CREATE | os.O_APPEND | os.O_RDWR, 0755)

	if err != nil {
		fmt.Printf("Error read file: %s", err.Error())
		return
	}


	defer file.Close()

	_, err = file.WriteString(str) 

	if err != nil {
		fmt.Printf("Error write file: %s", err.Error())
	}
}