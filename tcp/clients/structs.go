package clients

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net"
	"time"
)

type Response struct {
	Status      bool          `json:"status"`
	NeedFields  bool          `json:"need_fields"`
	Errors      []ErrorItem   `json:"errors"`
	Values      []interface{} `json:"values"`
	TmRequest   string        `json:"tm_req"`
	TmRequestSt time.Time     `json:"-"`
}

type ErrorItem struct {
	Key     int    `json:"key"`
	Message string `json:"message"`
}

type RequestСreateToken struct {
	Conn      net.Conn           `json:"-"`
	Key       string             `json:"key"`
	Method    string             `json:"method"`
	Token     string             `json:"token"`
	Service   string             `json:"service"`
	Data      primitive.ObjectID `json:"data"`
	IsService bool               `json:"is_service"`
}

type RequestFindByIds struct {
	Conn      net.Conn             `json:"-"`
	Key       string               `json:"key"`
	Method    string               `json:"method"`
	Token     string               `json:"token"`
	Service   string               `json:"service"`
	Data      []primitive.ObjectID `json:"data"`
	IsService bool                 `json:"is_service"`
}

type ReqSetUserData struct {
	Conn      net.Conn           `json:"-"`
	Key       string             `json:"key"`
	Method    string             `json:"method"`
	Token     string             `json:"token"`
	Service   string             `json:"service"`
	Data      UserDataByRobocode `json:"data"`
	IsService bool               `json:"is_service"`
}

type UserInfoBirthday struct {
	Year  int `json:"year" bson:"year"`
	Month int `json:"month" bson:"month"`
	Day   int `json:"day" bson:"day"`
}

type UserDataByRobocode struct {
	Id         primitive.ObjectID `json:"id" bson:"_id"`
	Name       string             `json:"name" bson:"name"`
	Surname    string             `json:"surname" bson:"surname"`
	Patronymic string             `json:"patronymic" bson:"patronymic"`
	Snils      string             `json:"snils" bson:"snils"`
	Birthday   *UserInfoBirthday  `json:"birthday" bson:"birthday"`
	Phone      string             `json:"phone" bson:"phone"`
	Email      string             `json:"email" bson:"email"`
	Sex        string             `json:"sex" bson:"sex"`
	Avatar     string             `json:"avatar" bson:"avatar"`
}

type ReqSignUp struct {
	Conn    net.Conn      `json:"-"`
	Key     string        `json:"key"`
	Method  string        `json:"method"`
	Token   string        `json:"token"`
	Service string        `json:"service"`
	Data    ReqSignUpInfo `json:"data"`
}

type ReqSignUpInfo struct {
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	UserInfo UsInfo `json:"user_info"`
}

type UsInfo struct {
	Passport FIO `json:"passport"`
}

type FIO struct {
	Name       string `json:"name"`
	SurName    string `json:"sur_name"`
	Patronymic string `json:"patronymic"`
}

type ReqGetAllUsers struct {
	Conn    net.Conn   `json:"-"`
	Key     string     `json:"key"`
	Method  string     `json:"method"`
	Token   string     `json:"token"`
	Service string     `json:"service"`
	Data    Pagination `json:"data"`
}

type Pagination struct {
	Offset int `json:"offset"`
	Page   int `json:"page"`
	Limit  int `json:"limit,omitempty"`
}

type RobocodeUser struct {
	Name        string `json:"name" bson:"name"`
	Surname     string `json:"surname" bson:"surname"`
	Patronymic  string `json:"patronymic" bson:"patronymic"`
	Phone       string `json:"phone" bson:"phone"`
	Email       string `json:"email" bson:"email"`
	Snils       string `json:"snils" bson:"snils"`
	Sex         string `json:"sex" bson:"sex"`
	BirthDate   string `json:"birth_date" bson:"birth_date"`
	AccountId   string `json:"account_id" bson:"account_id"`
	PrincipalId string `json:"principal_id" bson:"principal_id"`
	LeaderId    string `json:"leader_id" bson:"leader_id"`
	UntiId      string `json:"unti_id" bson:"unti_id"`
}

type ReqRobocodeUser struct {
	Conn      net.Conn       `json:"-"`
	Key       string         `json:"key"`
	Method    string         `json:"method"`
	Token     string         `json:"token"`
	Service   string         `json:"service"`
	Data      []RobocodeUser `json:"data"`
	IsService bool           `json:"is_service"`
}

type RequestFindByKeyCount struct {
	Conn      net.Conn           `json:"-"`
	Key       string             `json:"key"`
	Method    string             `json:"method"`
	Token     string             `json:"token"`
	Service   string             `json:"service"`
	Data      FindByKeyCountData `json:"data"`
	IsService bool               `json:"is_service"`
}

type RequestFindByKey struct {
	Conn      net.Conn `json:"-"`
	Key       string   `json:"key"`
	Method    string   `json:"method"`
	Token     string   `json:"token"`
	Service   string   `json:"service"`
	Data      string   `json:"data"`
	IsService bool     `json:"is_service"`
}

type FindByKeyCountData struct {
	Query      string     `json:"query"`
	Pages      Pagination `json:"pages"`
	NotDeleted bool       `json:"not_deleted"`
}
