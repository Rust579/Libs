package websoket_client

import (
	"Libs/configs"
	"Libs/tcp/clients"
	"encoding/json"
	"fmt"
	"github.com/fasthttp/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type ErrorItem struct {
	Key     int    `json:"key"`
	Message string `json:"message"`
}

type Response struct {
	Status      bool          `json:"status"`
	NeedFields  bool          `json:"need_fields"`
	Errors      []ErrorItem   `json:"errors"`
	Values      []interface{} `json:"values"`
	TmRequest   string        `json:"tm_req"`
	TmRequestSt time.Time     `json:"-"`
}

func WSClient() {
	// Настроим заголовки для аутентификации
	headers := http.Header{
		"key":     []string{configs.Cfg.ServiceKeys.ServiceAssessment},
		"service": []string{clients.ServiceAssessment},
	}

	// Подключаемся к серверу
	dialer := websocket.DefaultDialer
	conn, resp, err := dialer.Dial("ws://localhost:8989/api/v1/getconn", headers)
	if err != nil {
		log.Printf("Error connecting to WebSocket server: %v", err)
		if resp != nil {
			log.Printf("HTTP Response Status: %s", resp.Status)

			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println("Error reading response body:", err)
				return
			}

			// Декодирование JSON-данных в структуру Response
			var response Response
			err = json.Unmarshal(bodyBytes, &response)
			if err != nil {
				log.Println("Error unmarshalling JSON:", err)
				return
			}

			log.Printf("HTTP Response body: %v", response)
		}
		return
	}
	defer conn.Close()

	// Считываем первое сообщение от сервера
	_, message, err := conn.ReadMessage()
	if err != nil {
		log.Fatalf("Error reading message: %v", err)
	}
	log.Printf("Received: %s", message)

	// Отправляем команду getallusers
	err = conn.WriteMessage(websocket.TextMessage, []byte("getalluserss"))
	if err != nil {
		log.Fatalf("Error writing message: %v", err)
	}

	// Настраиваем пинг-понг
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	go func() {
		for {
			<-ticker.C
			err := conn.WriteMessage(websocket.TextMessage, []byte("ping"))
			if err != nil {
				log.Printf("Error sending ping: %v", err)
				return
			}
			fmt.Println("ping")
		}
	}()

	// Чтение ответов от сервера
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			return
		}
		log.Printf("Received: %s", message)

		// Other messages
		/*if string(message) == "ping" {
			err = conn.WriteMessage(websocket.TextMessage, []byte("pong"))
			if err != nil {
				log.Fatalf("Error writing pong message: %v", err)
			}
			log.Println("pong sent to server")
		}*/

		time.Sleep(3 * time.Second)
	}
}
