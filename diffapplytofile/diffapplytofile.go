package diffapplytofile

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

type UpdateRequest struct {
	LineNumber int    `json:"line_number"`
	Text       string `json:"text"`
	Action     string `json:"action"` // "add", "delete", "update"
}

func ApplyDiffToFile() {
	// Пример запроса
	requests := []UpdateRequest{
		{
			LineNumber: 21,
			Text:       "	// add 30",
			Action:     "delete",
		},
		{
			LineNumber: 28,
			Text:       "	// add 30",
			Action:     "delete",
		},
		{
			LineNumber: 30,
			Text:       "	// add 30",
			Action:     "delete",
		},
		{
			LineNumber: 34,
			Text:       "	// add 34",
			Action:     "add",
		},
		{
			LineNumber: 22,
			Text:       "	// upd 22",
			Action:     "update",
		},
		{
			LineNumber: 31,
			Text:       "	defer file.Close() // upd 31",
			Action:     "update",
		},
	}

	filePath := "C:\\Projects Go\\service.go"

	t1 := time.Now()

	err := processUpdateRequests(filePath, requests)
	if err != nil {
		log.Fatalf("Failed to process update request: %v", err)
	}

	fmt.Println("processUpdateRequest time taken:", time.Since(t1))
	fmt.Println("File updated successfully")
}

func processUpdateRequests(filePath string, requests []UpdateRequest) error {

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	lines := strings.Split(string(fileContent), "\n")

	// Сортировка запросов по типу действий
	var deleteRequests, updateRequests, addRequests []UpdateRequest
	for _, request := range requests {
		switch request.Action {
		case "delete":
			deleteRequests = append(deleteRequests, request)
		case "update":
			updateRequests = append(updateRequests, request)
		case "add":
			addRequests = append(addRequests, request)
		}
	}

	// Сортировка запросов на удаление по номеру строки в обратном порядке
	sort.Slice(deleteRequests, func(i, j int) bool {
		return deleteRequests[i].LineNumber > deleteRequests[j].LineNumber
	})

	// Сортировка запросов на добавление по номеру строки в обратном порядке
	sort.Slice(addRequests, func(i, j int) bool {
		return addRequests[i].LineNumber > addRequests[j].LineNumber
	})

	// Сначала обновляем
	for _, request := range updateRequests {
		lineIndex := request.LineNumber - 1
		if lineIndex < 0 || lineIndex >= len(lines) {
			return fmt.Errorf("invalid line number for update action")
		}
		lines[lineIndex] = request.Text
	}

	// Затем удаляем
	for _, request := range deleteRequests {
		lineIndex := request.LineNumber - 1
		if lineIndex < 0 || lineIndex >= len(lines) {
			return fmt.Errorf("invalid line number for delete action")
		}
		lines = append(lines[:lineIndex], lines[lineIndex+1:]...)
	}

	// И добавляем с учетом смещения после удалений
	for _, request := range addRequests {
		offset := 0

		for _, v := range deleteRequests {
			if request.LineNumber > v.LineNumber {
				offset--
			}
		}

		insertPosition := request.LineNumber + offset - 1
		if insertPosition < 0 || insertPosition > len(lines) {
			return fmt.Errorf("invalid line number for add action")
		}
		lines = append(lines[:insertPosition], append([]string{request.Text}, lines[insertPosition:]...)...)
	}

	updatedContent := strings.Join(lines, "\n")
	err = os.WriteFile(filePath, []byte(updatedContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write updated content to file: %w", err)
	}

	return nil
}
