package main

import (
	"Libs/configs"
	"github.com/sirupsen/logrus"
)

var gfsd = make(chan bool)

func main() {

	err := configs.Init()
	if err != nil {
		logrus.Fatalf("error open config file: %s", err.Error())
	}

	//clients.GetUserData()

	//hagrid.InitRmq()

	//websoket_client.WSClient()

	//<-gfsd

}
