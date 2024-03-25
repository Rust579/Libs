package clients

import (
	"Libs/configs"
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

func GetUserData() {

	request := ReqSetUserData{
		Key:     configs.Cfg.ServiceKeys.ServiceAssessment,
		Method:  MethodGetUserData,
		Token:   "7ca656b7ed8e1b3c8c5dc7843c1eac553ab87a3fc1aa8bdc507b0bb66ba7b57f",
		Service: ServiceAssessment,
		//IsService: true,
	}

	var resp *Response

	reqBytes, err := json.Marshal(request)
	if err != nil {
		return
	}

	conn, err := net.Dial("tcp", configs.Cfg.OneId.LocalTcpUrl)
	if err != nil {
		return
	}
	defer conn.Close()
	_, err = conn.Write(reqBytes)
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

	message, _ := bufio.NewReader(conn).ReadString('\n')
	if err = json.Unmarshal([]byte(message), &resp); err != nil {
		return
	}

	fmt.Println("resp:", resp)

}
