package clients

import (
	"Libs/configs"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net"
)

func FindByIds() {

	usIdObj, _ := primitive.ObjectIDFromHex("65eec57c5b804ea2af8e4ee3")
	// Создаем структуру для вашего запроса.
	request := RequestFindByIds{
		Key:    configs.Cfg.ServiceKeys.ServiceSupport,
		Method: MethodFindByIds,
		//Token:     "479590fd42464b304a3b8469399c0cc01c4236da49509ef803e5e609a2726029",
		Service:   ServiceSupport,
		Data:      []primitive.ObjectID{usIdObj},
		IsService: true,
	}

	var resp *Response

	conn, err := net.Dial("tcp", configs.Cfg.OneId.ProdTcpUrl)
	if err != nil {
		fmt.Println("Ошибка при подключении к серверу:", err)
		return
	}
	fmt.Println("Подключение к серверу")
	defer conn.Close()

	requestJSON, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Ошибка при преобразовании в JSON:", err)
		return
	}

	// Отправляем запрос как массив байтов
	_, err = conn.Write(requestJSON)
	if err != nil {
		fmt.Println("Ошибка при отправке данных:", err)
		return
	}

	// Отправляем символ новой строки для завершения запроса на сервере
	_, err = conn.Write([]byte("\n"))
	if err != nil {
		fmt.Println("Ошибка при отправке символа новой строки:", err)
		return
	}

	// Создаем буфер для чтения ответа от сервера
	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	// Декодируем ответ из JSON
	if err = json.Unmarshal(buffer[:n], &resp); err != nil {
		fmt.Println("Ошибка при декодировании JSON:", err)
		return
	}

	fmt.Println("Ответ от сервера:", resp)
}
