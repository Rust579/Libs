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

type CodeGCILRequest struct {
	Code string `json:"code"`
}

type LintResult struct {
	Issues []interface{} `json:"Issues"`
}

func main() {
	http.HandleFunc("/check", checkHandlerGCIL)
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}

func checkHandlerGCIL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var codeReq CodeGCILRequest
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

	// Запуск golangci-lint
	cmd := exec.Command("golangci-lint", "run", "--out-format", "json", tempFile)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		fmt.Printf("golangci-lint error: %s\n", stderr.String()) // Логирование stderr
		http.Error(w, fmt.Sprintf("Error running golangci-lint: %s", stderr.String()), http.StatusInternalServerError)
		fmt.Println("err", err)
		//return
	}

	var result interface{}
	err = json.Unmarshal(out.Bytes(), &result)
	if err != nil {
		http.Error(w, "Error parsing golangci-lint output", http.StatusInternalServerError)
		return
	}

	localDir := "./gosec" // Ваша локальная директория, например, "./results"
	filePath := filepath.Join(localDir, "golangci-lint-result.json")
	jsonResult, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		http.Error(w, "Error encoding result to JSON", http.StatusInternalServerError)
		return
	}

	err = ioutil.WriteFile(filePath, jsonResult, 0644)
	if err != nil {
		http.Error(w, "Error saving JSON result", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
