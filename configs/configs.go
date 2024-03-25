package configs

import (
	"encoding/json"
	"fmt"
	"os"
)

var Cfg Configs

func Init() error {

	envFile, errFile := os.Open("env.json")
	if errFile != nil {
		return errFile
	}

	decoder := json.NewDecoder(envFile)

	if errCfg := decoder.Decode(&Cfg); errCfg != nil {
		return errCfg
	}

	fmt.Println(Cfg)
	return nil
}
