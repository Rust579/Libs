package diffapplytofile

import (
	"bufio"
	"fmt"
	"os"
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
				OldStart: 10,
				OldLines: 0,
				NewStart: 10,
				NewLines: 1,
				Lines: []string{
					"+",
				},
			},
			{
				OldStart: 52,
				OldLines: 4,
				NewStart: 52,
				NewLines: 0,
				Lines: []string{
					"-if len(lr) == 0 {",
					"-return nil, nil",
					"-}",
					"-",
				},
			},
		},
	}

	filePath := "C:\\Projects Go\\copy-main.go"

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

	for _, hunk := range patch.Hunks {

		var newLines []string
		for _, line := range hunk.Lines {
			if strings.HasPrefix(line, "+") {
				newLines = append(newLines, line[1:])
			}
		}

		lines = append(lines[:hunk.OldStart-1], append(newLines, lines[hunk.NewStart+hunk.OldLines-1:]...)...)
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
