package codeBERT

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func analyzeCodeAndCommentWithCodeBERT(code, comment string) (string, error) {
	combined := fmt.Sprintf("%s\n%s", code, comment)
	cmd := exec.Command("python3", "C:\\Projects Go\\Libs\\codeBERT\\codebert_script.py")
	cmd.Stdin = bytes.NewBufferString(combined)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("command error: %v, stderr: %s", err, stderr.String())
	}

	return out.String(), nil
}

func filterWarnings(output string) string {
	lines := strings.Split(output, "\n")
	var filteredLines []string

	for _, line := range lines {
		// Убираем строки, содержащие предупреждения или необработанные сообщения
		if strings.Contains(line, "FutureWarning") || strings.Contains(line, "not initialized") {
			continue
		}
		filteredLines = append(filteredLines, line)
	}

	return strings.Join(filteredLines, "\n")
}

func AnalyzeCodeWithCodeBERT() {

	code := `
func QrCode() {
	url := "https://habr.com/ru/companies/slurm/articles/704208/"
	q, _ := qrcode.New(url, qrcode.Medium)
	str := q.ToSmallString(false)
	fmt.Println(str)
}
    `

	result, err := analyzeCodeAndCommentWithCodeBERT(code, "// QrCode")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("CodeBERT analysis result:")
	fmt.Println(result)
}
