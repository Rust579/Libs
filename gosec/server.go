package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

type CodeRequest struct {
	Code string `json:"code"`
}

type GosecResult struct {
	GolanErrors  interface{}   `json:"Golang errors"`
	Issues       []interface{} `json:"issues"`
	Stats        interface{}   `json:"stats"`
	GosecVersion interface{}   `json:"GosecVersion"`
}

func main() {
	http.HandleFunc("/check", checkHandler)
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var codeReq CodeRequest
	err := json.NewDecoder(r.Body).Decode(&codeReq)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Создание временной директории для кода
	tempDir, err := ioutil.TempDir("", "code-")
	if err != nil {
		http.Error(w, "Could not create temporary directory", http.StatusInternalServerError)
		return
	}
	defer os.RemoveAll(tempDir) // Удалить директорию после завершения обработки

	// Создание временного файла в этой директории
	tempFile := filepath.Join(tempDir, "temp.go")
	err = ioutil.WriteFile(tempFile, []byte(codeReq.Code), 0644)
	if err != nil {
		http.Error(w, "Could not write to temporary file", http.StatusInternalServerError)
		return
	}

	// Запуск gosec
	cmd := exec.Command("gosec", "-fmt", "json", tempDir)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		fmt.Printf("gosec error: %s\n", stderr.String()) // Логирование stderr
		http.Error(w, fmt.Sprintf("Error running gosec: %s", stderr.String()), http.StatusInternalServerError)
		return
	}

	var result GosecResult
	err = json.Unmarshal(out.Bytes(), &result)
	if err != nil {
		http.Error(w, "Error parsing gosec output", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
