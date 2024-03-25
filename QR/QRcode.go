package QR

import (
	"fmt"
	"github.com/skip2/go-qrcode"
)

func QrCode() {
	url := "https://habr.com/ru/companies/slurm/articles/704208/"

	q, _ := qrcode.New(url, qrcode.Medium)

	str := q.ToSmallString(false)

	fmt.Println(str)
}
