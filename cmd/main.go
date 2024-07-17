package main

import (
	"Libs/configs"
	"Libs/grpc"
	"github.com/sirupsen/logrus"
)

var gfsd = make(chan bool)

func main() {

	err := configs.Init()
	if err != nil {
		logrus.Fatalf("error open config file: %s", err.Error())
	}

	grpc.GRPCZip()

	/*clients.GetUserData()

	hagrid.InitRmq()

	websoket_client.WSClient()

	QR.QrCode("https://habr.com/ru/companies/slurm/articles/704208/")*/

	//<-gfsd

}
