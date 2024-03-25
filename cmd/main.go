package main

import (
	"Libs/configs"
	"Libs/tcp/clients"
	"github.com/sirupsen/logrus"
)

func main() {

	err := configs.Init()
	if err != nil {
		logrus.Fatalf("error open config file: %s", err.Error())
	}

	clients.GetUserData()

}
