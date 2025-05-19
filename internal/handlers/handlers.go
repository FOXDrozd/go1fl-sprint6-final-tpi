package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"go1fl-sprint6-final/internal/service"
)


func GetHtmlFormat(w http.ResponseWriter, r *http.Request){
	currencyDir, err := os.Getwd()

	if err != nil {
		fmt.Printf("Error Getwd in getHtmlFormat %s\n",err.Error())
	}

	data, err := os.ReadFile(filepath.Join( filepath.Dir(currencyDir), "index.html"))

	if err != nil {
		fmt.Printf("Error read file index.html %s\n", err.Error())
	}

	defer 

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func UploadFile(w http.ResponseWriter, r *http.Request){
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
	}

	currencyDir, err := os.Getwd()

	file, err := os.OpenFile(filepath.Join(filepath.Dir(currencyDir), time.Now().UTC().Format("02-01-2006 03-04-05")), os.O_CREATE | os.O_APPEND | os.O_RDWR, 0755)

	if err != nil {
		fmt.Printf("Error read file: %s", err.Error())
	}


	defer file.Close()

	_, err = file.WriteString(str) 

	if err != nil {
		fmt.Printf("Error write file: %s", err.Error())
	}
}