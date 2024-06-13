package clients

import (
	"Libs/configs"
	"encoding/json"
	"fmt"
	"net"
)

func GetUserData() {

	request := ReqSetUserData{
		Key:     configs.Cfg.ServiceKeys.ServiceAssessment,
		Method:  MethodGetUserData,
		Token:   "6dff98811cce3371bced1c368af79ba6462a58d5dc400618d683475bace5da20",
		Service: ServiceAssessment,
		//IsService: true,
	}

	var resp *Response

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
