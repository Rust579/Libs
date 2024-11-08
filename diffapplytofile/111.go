package diffapplytofile

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Patch struct {
	OldFileName string `json:"oldFileName"`
	NewFileName string `json:"newFileName"`
	OldHeader   string `json:"oldHeader"`
	NewHeader   string `json:"newHeader"`
	Hunks       []Hunk `json:"hunks"`
}

type Hunk struct {
	OldStart int      `json:"oldStart"`
	OldLines int      `json:"oldLines"`
	NewStart int      `json:"newStart"`
	NewLines int      `json:"newLines"`
	Lines    []string `json:"lines"`
}

func ApplyPatch() {

	patch := Patch{
		Hunks: []Hunk{
			{
				OldStart: 1,
				OldLines: 1,
				NewStart: 1,
				NewLines: 1,
				Lines: []string{
					"-package linter",
					"+package liter",
				},
			},
			{
				OldStart: 3,
				OldLines: 4,
				NewStart: 3,
				NewLines: 1,
				Lines: []string{
					"-import (",
					"-\t\"os\"",
					"-\t\"tea.gitpark.ru/sast/dockerwrapper/internal/configs\"",
					"-)",
					"+import (\"os\", \"tea.gitpark.ru/sast/dockerwrapper/internal/configs\")",
				},
			},
			{
				OldStart: 22,
				OldLines: 0,
				NewStart: 19,
				NewLines: 3,
				Lines: []string{
					"+\t// implementation",
					"+\t// 1",
					"+\t// 2",
				},
			},
		},
	}

	filePath := "C:\\Projects Go\\copy-trivy.go"

	err := applyPatch(filePath, patch)
	if err != nil {
		fmt.Printf("Error applying patch: %v\n", err)
	}

}

// applyPatch применяет patch к файлу
func applyPatch(filePath string, patch Patch) error {

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("could not read file: %v", err)
	}

	// Сортировка хунков по OldStart в порядке убывания
	sort.Slice(patch.Hunks, func(i, j int) bool {
		return patch.Hunks[i].OldStart > patch.Hunks[j].OldStart
	})

	for _, hunk := range patch.Hunks {
		oldStart := hunk.OldStart - 1
		newLines := hunk.Lines

		// Удаление старых строк
		lines = append(lines[:oldStart], lines[oldStart+hunk.OldLines:]...)

		// Вставка новых строк
		for i := len(newLines) - 1; i >= 0; i-- {
			line := newLines[i]
			if strings.HasPrefix(line, "-") {
				continue
			} else if strings.HasPrefix(line, "+") {
				line = strings.TrimPrefix(line, "+")
			}
			lines = append(lines[:oldStart], append([]string{line}, lines[oldStart:]...)...)
		}
	}

	outputFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("could not create output file: %v", err)
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("could not write to output file: %v", err)
		}
	}
	if err := writer.Flush(); err != nil {
		return fmt.Errorf("could not flush writer: %v", err)
	}

	return nil
}
