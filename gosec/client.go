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
	"log"
)

func QrCode(ctx *context.Context) {
	url := "https://habr.com/ru/companies/slurm/articles/704208/"
	a := *ctx
	fmt.Println(a)

	token := "9952f1088915435fac3f8f52064bfad86517530116166f8ce5854dd7bef3332e"
	fmt.Println(token)
	q, _ := qrcode.New(url, qrcode.Medium)
	str := q.ToSmallString(false)
	fmt.Println(str)

	log.Fatal(fmt.Sprintf("error occured while running swagger server: %s", "err"))
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

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}

	//fmt.Printf("Response: %s\n", body)
}
