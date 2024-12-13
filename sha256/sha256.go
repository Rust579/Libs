package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("C:\\Projects_Go\\TestProjectsSAST\\PyTest\\sast_reqs.py")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	hash := sha256.New()

	if _, err := io.Copy(hash, file); err != nil {
		fmt.Println(err)
	}
	hashSum := hex.EncodeToString(hash.Sum(nil))

	fmt.Println("hash sum:", hashSum)

}
