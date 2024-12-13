package main

import (
	"fmt"
	"os"
)

func main() {
	// Путь к бинарному файлу
	binFilePath := "C:/Projects_Go/wled/WLED/flash_contents.bin"

	// Путь к текстовому файлу
	txtFilePath := "flash_contents.txt"

	// Открываем бинарный файл для чтения
	binFile, err := os.Open(binFilePath)
	if err != nil {
		fmt.Printf("Ошибка при открытии бинарного файла: %v\n", err)
		return
	}
	defer binFile.Close()

	// Создаем текстовый файл для записи
	txtFile, err := os.Create(txtFilePath)
	if err != nil {
		fmt.Printf("Ошибка при создании текстового файла: %v\n", err)
		return
	}
	defer txtFile.Close()

	// Буфер для чтения данных
	buffer := make([]byte, 16)

	// Читаем бинарный файл и записываем в текстовый файл
	for {
		// Читаем данные в буфер
		n, err := binFile.Read(buffer)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Printf("Ошибка при чтении бинарного файла: %v\n", err)
			return
		}

		// Преобразуем данные в шестнадцатеричный формат
		for i := 0; i < n; i++ {
			fmt.Fprintf(txtFile, "%02X ", buffer[i])
		}
		fmt.Fprintln(txtFile)
	}

	fmt.Println("Содержимое бинарного файла успешно сохранено в текстовый файл.")
}
