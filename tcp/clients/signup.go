package clients

import (
	"Libs/configs"
	"encoding/json"
	"fmt"
	"net"
)

func SignUp() {

	request := ReqSignUp{
		Key:    configs.Cfg.ServiceKeys.ServiceRobocode,
		Method: MethodSignUp,
		//Token:   "9952f1088915435fac3f8f52064bfad86517530116166f8ce5854dd7bef3332e",
		Service: ServiceRobocode,
		Data: ReqSignUpInfo{
			Phone: "8612345678912",
			Email: "palakat96@mail.ru",
			UserInfo: UsInfo{
				Passport: FIO{
					Name:       "ххх",
					SurName:    "ууу",
					Patronymic: "ааа",
				},
			},
		},
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
