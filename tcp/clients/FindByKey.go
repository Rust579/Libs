package clients

import (
	"Libs/configs"
	"encoding/json"
	"fmt"
	"net"
)

func FindByKey() {

	//usIdObj, _ := primitive.ObjectIDFromHex("626538013814c85c474dd491")
	// Создаем структуру для вашего запроса.
	request := RequestFindByKey{
		Key:    configs.Cfg.ServiceKeys.ServiceSupport,
		Method: MethodFindByKey,
		//Token:     "479590fd42464b304a3b8469399c0cc01c4236da49509ef803e5e609a2726029",
		Service:   ServiceSupport,
		Data:      "657aed51e0cd039559a65200",
		IsService: true,
	}

	//var resp *Response
	conn, err := net.Dial("tcp", configs.Cfg.OneId.LocalTcpUrl)
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
	buffer := make([]byte, 40960)
	_, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	fmt.Println(string(buffer))
}
