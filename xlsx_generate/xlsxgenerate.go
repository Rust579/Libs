package xlsx_generate

import (
	"github.com/tealeg/xlsx"
	"unicode/utf8"
)

func GenerateExcelFile(filename string) error {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		return err
	}

	// Добавляем заголовок
	headerRow := sheet.AddRow()
	addCellWithStyle(headerRow, "Name", "Arial", 12)
	addCellWithStyle(headerRow, "оценка руководителя", "Arial", 12)
	addCellWithStyle(headerRow, "средняя оценка подчиненных", "Arial", 12)
	addCellWithStyle(headerRow, "средняя оценка коллег", "Arial", 12)
	addCellWithStyle(headerRow, "итого по экспертам", "Arial", 12)
	addCellWithStyle(headerRow, "самооценка", "Arial", 12)
	addCellWithStyle(headerRow, "расхождения", "Arial", 12)

	headers := []string{"Name", "оценка руководителя", "средняя оценка подчиненных", "средняя оценка коллег", "итого по экспертам", "самооценка", "расхождения"}

	// Добавляем данные
	data := [][]string{
		{"Работа в команде", "", "", "", "", "", ""},
		{"Ориентация на цели", "4", "1", "3", "5", "2", "3"},
	}

	for _, row := range data {
		newRow := sheet.AddRow()
		for _, cellValue := range row {
			cell := newRow.AddCell()
			cell.Value = cellValue
		}
	}

	// Устанавливаем ширину столбца на основе содержимого ячеек
	for i := 0; i < len(headerRow.Cells); i++ {
		col := sheet.Col(i)
		col.Width = autoWidth(data, headers, i)
	}

	// Сохраняем файл
	err = file.Save(filename)
	if err != nil {
		return err
	}

	return nil
}

func addCellWithStyle(row *xlsx.Row, value, fontName string, fontSize float64) {
	cell := row.AddCell()
	cell.Value = value
	style := xlsx.NewStyle()
	font := *xlsx.NewFont(int(fontSize), fontName)
	style.Font = font
	cell.SetStyle(style)
}

// Функция для определения ширины столбца на основе содержимого ячеек
func autoWidth(data [][]string, headers []string, columnIndex int) float64 {
	var maxWidth float64
	for _, row := range data {
		widthData := utf8.RuneCountInString(row[columnIndex]) // коэффициент для учета ширины символов
		widthHeader := utf8.RuneCountInString(headers[columnIndex])
		if widthData >= widthHeader {
			maxWidth = float64(widthData)
		} else {
			maxWidth = float64(widthHeader)
		}
	}
	return maxWidth
}
