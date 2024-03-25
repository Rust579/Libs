package clients

import (
	"Libs/configs"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net"
)

func SetUserInfo() {

	usIdObj, _ := primitive.ObjectIDFromHex("658bc54e38e210cc706603b7")
	//usIdObj, _ := primitive.ObjectIDFromHex("657aed51e0cd039559a65200")

	// Создаем структуру для вашего запроса.
	request := ReqSetUserData{
		Key:     configs.Cfg.ServiceKeys.ServiceRobocode,
		Method:  MethodSetUserInfo,
		Service: ServiceRobocode,
		Data: UserDataByRobocode{
			Id: usIdObj,
			//Name:       "- Рус тем - фы ва ",
			//Surname:    " -Ахм еров -фы вп",
			//Patronymic: " Рави левич фы впа",
			//Snils:      "595 832 266 45111",
			Sex: "male",
			//Email:      "axmerovrustem@list.ru",
			//Phone:      "89173494694111111111111",
			Avatar: "https://learn.unionepro.ru/uploads/pictures/6fa543ac9868498889d83fab577ffc04.jpg",
			//Avatar: "https://avatars.mds.yandex.net/get-yapic/20749/TlcP7cQFsAmecgpwjmsBJNSFiE-1/islands-200",
			Birthday: &UserInfoBirthday{
				Year:  1990,
				Month: 10,
				Day:   1,
			},
		},
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
