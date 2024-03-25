package maskEmail

import (
	"fmt"
	"strings"
)

// MaskEmail Для строк вида aaa://xxxx:yyyy@bbb, где хххх и уууу - логин и пароль
// На выходе будет aaa://x**x:y**y@bbb
// Символы :, //, @ обязательны, если одного из нет или их порядок отличается - вернется оригинальная строка
func MaskEmail(config string) string {

	parts := strings.SplitN(config, "@", 2)
	if len(parts) < 2 {
		return config
	}

	result := fmt.Sprintf("%s@%s", MaskLoginPassword(parts[0]), parts[1])

	return result
}

// MaskLoginPassword Для обычной строки, например xxxxxx
// На выходе:
// 6 символов и больше - xx**xx
// от 3 до 5 символов - x**x
// 2 символа и меньше - вернется оригинальная строка
func MaskLoginPassword(psw string) string {
	if len(psw) <= 2 {
		return psw
	}
	if len(psw) <= 5 {
		return psw[:1] + strings.Repeat("*", len(psw)-2) + psw[len(psw)-1:]

	} else {
		return psw[:2] + strings.Repeat("*", len(psw)-4) + psw[len(psw)-2:]
	}
}
