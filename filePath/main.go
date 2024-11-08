package main

import (
	"fmt"
	"path/filepath"
)

const (
	goMod        = "/go.mod"
	requirements = "/requirements.txt"
)

func main() {

	pId := "aaa"

	goModFilePath := filepath.Join(pId, goMod)
	fmt.Println("goModFilePath:", goModFilePath)

}
