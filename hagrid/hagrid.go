package hagrid

import (
	"github.com/sirupsen/logrus"
	broker "github.com/unione-pro/core/pkg/broker/amqp"
	"log"
	"time"
)

type RabbitCfg struct {
	Host     string
	VHost    string
	UserName string
	Password string
}

var queues []string

func InitRmq() error {
	rabbitCfg := RabbitCfg{
		Host:     "10.62.10.74",
		VHost:    "",
		UserName: "guest",
		Password: "guest",
	}

	var addr string
	if rabbitCfg.VHost != "" {
		addr = "amqp://" + rabbitCfg.UserName + ":" + rabbitCfg.Password + "@" + rabbitCfg.Host + "/" + rabbitCfg.VHost
	} else {
		addr = "amqp://" + rabbitCfg.UserName + ":" + rabbitCfg.Password + "@" + rabbitCfg.Host
	}

	queues = append(queues, "hagrid_request_test")
	if err := broker.Init(addr, queues); err != nil {
		logrus.Println("Could not create a connection with ", addr)
		time.Sleep(1 * time.Second)
		return err
	}

	if err := broker.InitFanout(addr, []string{}, "hagrid_exc_test"); err != nil {
		return err
	}

	type DataComp struct {
		Title string `json:"title"`
		Level string `json:"level"`
	}

	type Data struct {
		Name       string     `json:"name"`
		CourseId   string     `json:"course_id"`
		CourseName string     `json:"course_name"`
		OrgName    string     `json:"org_name"`
		EndDate    string     `json:"end_date"`
		QrCode     bool       `json:"qr_code"`
		CertNumber string     `json:"cert_number"`
		DataComp   []DataComp `json:"data_comp"`
		Background string     `json:"background"`
		TestOption string     `json:"test_option"`
	}

	type CertGen struct {
		Pattern   string `json:"pattern"`
		UseriD    string `json:"user_id"`
		TimeStamp string `json:"time_stamp"`
		Data      Data   `json:"data"`
	}

	/*usId, _ := primitive.ObjectIDFromHex("65141ac749cb1b180e5d97f8")
	cId, _ := primitive.ObjectIDFromHex("651bbbd73fcdd1c0dafd01e8")*/

	/*msg := CertGen{
		Pattern:   "cert_generated_v2",
		UseriD:    "65141ac749cb1b180e5d97f8",
		TimeStamp: "2023-10-06T15:04:05Z",
		Data: Data{
			Name:       "ТЕСТ Руст Ахмеров V2222",
			CourseId:   "651bbbd73fcdd1c0dafd01e8",
			CourseName: "тест курса 06.10",
			OrgName:    "ТЕСТ Руст Ахмеров",
			EndDate:    "06.10.2023",
			QrCode:     true,
			CertNumber: "000001",
			DataComp: []DataComp{{
				Title: "компетенция/ руст V2222",
				Level: "Экспертный",
			}},
		},
	}*/

	msgV2 := CertGen{
		Pattern:   "cert_generated_v2",
		UseriD:    "65141ac749cb1b180e5d97f8",
		TimeStamp: "2023-10-06T15:04:05Z",
		Data: Data{
			Name:       "ТЕСТ Руст Ахмеров V2222",
			CourseId:   "651bbbd73fcdd1c0dafd01e8",
			CourseName: "тест курса 06.10",
			OrgName:    "ТЕСТ Руст Ахмеров",
			EndDate:    "06.10.2023",
			QrCode:     true,
			CertNumber: "000001",
			Background: "https://disk.yandex.ru/i/wc-BIWpdYQfjdg",
			TestOption: "STUDY",
			DataComp: []DataComp{{
				Title: "компетенция/ руст V2222",
				Level: "Экспертный",
			}},
		},
	}

	//publish(msg)
	publish(msgV2)

	return nil
}

/*func PublishEventExchange(data interface{}) error {
	if err := broker.SendV2(data, "", ExchangeAssEngine); err != nil {
		return err
	}
	log.Println("msg published to "+ExchangeAssEngine, data)
	return nil
}*/

// publish sends msg to RMQ server to prescribed queue
func publish(msg interface{}) {
	err := broker.SendV2(msg, "hagrid_request_test", "")
	if err != nil {
		log.Println("Failed to publish to queue ERROR: "+err.Error(), msg)
		logrus.Println("Failed to publish to queue")
	}
}
