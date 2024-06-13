package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CodeReq struct {
	Code string `json:"code"`
}

func main() {
	code := `package QR

import (
	"context"
	"fmt"
	"github.com/skip2/go-qrcode"
)

func QrCode(ctx *context.Context) {
	url := "https://habr.com/ru/companies/slurm/articles/704208/"
	fmt.Println(*ctx)
	q, _ := qrcode.New(url, qrcode.Medium)

	str := q.ToSmallString(false)

	fmt.Println(str)
}
`

	codeReq := CodeReq{Code: code}
	reqBody, err := json.Marshal(codeReq)
	if err != nil {
		fmt.Printf("Error marshalling request: %v\n", err)
		return
	}

	resp, err := http.Post("http://localhost:8080/check", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}

	fmt.Printf("Response: %s\n", body)
}
