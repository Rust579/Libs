package main

import (
	"fmt"
	"regexp"
)

// Структура для хранения ответа в формате JSON
type JSONResponse struct {
	Answers []Answer `json:"answers"`
}

type Answer struct {
	LineNumber string `json:"line number"`
	Issue      string `json:"issue"`
}

// Функция для парсинга JSON ответа
func parseJSONResponse(jsonText string) (*JSONResponse, error) {

	// Используем регулярное выражение для извлечения line number и issue
	re := regexp.MustCompile(`{"line number": "(\d+)", "issue": "([^"]+)"}`)
	matches := re.FindAllStringSubmatch(jsonText, -1)

	var answers []Answer
	for _, match := range matches {
		if len(match) == 3 {
			answers = append(answers, Answer{
				LineNumber: match[1],
				Issue:      match[2],
			})
		}
	}

	return &JSONResponse{Answers: answers}, nil
}

func main() {

	str1 := "```json\n{\n    \"answers\": [\n"
	str2 := "{\"line number\": \"21\", \"issue\": \"Comment does not match the code functionality, because it describes bubble sort, but the function 'FileProcess' does not implement sorting.\"},\n"
	str3 := "{\"line number\": \"101\", \"issue\": \"Comment does not match the code functionality, because it describes converting hunk to patch and applying it to the target file, but the function 'FileProcessV2' also includes logic for handling file status 'statusAdd' and 'statusDelete'.\"},\n "
	str4 := "{\"line number\": \"200\", \"issue\": \"Comment does not match the code functionality, because it describes checking the existence of a directory, but the function 'CheckExistsDir' checks both the existence of a directory in the cache and on the disk.\"}.\n"

	str := str1 + str2 + str3 + str4

	response, err := parseJSONResponse(str)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("response:", response)

}
