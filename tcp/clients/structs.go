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

type UsersShort struct {
	Id           primitive.ObjectID    `json:"id,omitempty" bson:"_id,omitempty"`
	Phone        string                `json:"phone" bson:"phone"`
	Email        string                `json:"email" bson:"email"`
	UserSpec     bool                  `json:"user_spec" bson:"user_spec"`
	Confirm      UserPublicDataConfirm `json:"confirm" bson:"confirm"`
	Info         UserInfo              `json:"user_info" bson:"user_info"`
	Login        string                `json:"login" bson:"login"`
	Sex          string                `json:"sex" bson:"sex"`
	Status       string                `json:"status" bson:"status"`
	Avatar       string                `json:"avatar" bson:"avatar"`
	Roles        []UserRole            `json:"roles" bson:"roles"`
	IsSupport    bool                  `json:"is_support" bson:"-"`
	Moderation   ModerationInfo        `json:"moderation" bson:"moderation"`
	RobocodeUser bool                  `json:"robocode_user" bson:"robocode_user"`
	RobocodeInfo RobocodeInfo          `json:"robocode_info" bson:"robocode_info"`
	CreatedAt    time.Time             `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time             `json:"updated_at" bson:"updated_at"`
	DeletedAt    time.Time             `json:"deleted_at" bson:"deleted_at"`
}

type UserPublicDataConfirm struct {
	Phone          bool      `json:"phone" bson:"phone"`
	PhoneNew       string    `json:"-" bson:"phone_new"`
	PhoneCodeTm    time.Time `json:"phone_code_tm" bson:"phone_code_tm"`
	Email          bool      `json:"email" bson:"email"`
	EmailNew       string    `json:"-" bson:"email_new"`
	EmailCodeTm    time.Time `json:"email_code_tm" bson:"email_code_tm"`
	Telegram       bool      `json:"telegram" bson:"telegram"`
	TelegramCodeTm time.Time `json:"telegram_code_tm" bson:"telegram_code_tm"`
}

type UserInfo struct {
	Geo      UserInfoGeo      `json:"geo" bson:"geo"`
	Work     UserInfoWork     `json:"work" bson:"work"`
	Birthday UserInfoBirthday `json:"birthday" bson:"birthday"`
	Passport UserInfoPassport `json:"passport" bson:"passport"`
	Docs     UserDocs         `json:"docs" bson:"docs"`
}

type UserInfoPassport struct {
	Name       string `json:"name" bson:"name"`
	SurName    string `json:"sur_name" bson:"sur_name"`
	Patronymic string `json:"patronymic" bson:"patronymic"`
}

type UserInfoGeo struct {
	Country  string `json:"country" bson:"country"`
	City     string `json:"city" bson:"city"`
	FiasId   string `json:"fias_id" bson:"fias_id"`
	FiasAddr string `json:"fias_addr" bson:"fias_addr"`
}

type UserDocs struct {
	Snils         string    `json:"snils" bson:"snils"`
	SnilsDop      string    `json:"snils_dop" bson:"snils_dop"`
	IdDocName     string    `json:"id_doc_name" bson:"id_doc_name"`
	IdDoc         string    `json:"id_doc" bson:"id_doc"`
	Images        DocImages `json:"images" bson:"images"`
	MsgProctoring string    `json:"msg_proct" bson:"msg_proct"`
}

type DocImage struct {
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Type        string `json:"type" bson:"type"`
	Image       string `json:"image" bson:"image"`
}

type DocImages struct {
	Passport      []DocImage `json:"passport" bson:"passport"`
	Diploma       []DocImage `json:"diploma" bson:"diploma"`
	Qualification []DocImage `json:"qualification" bson:"qualification"`
	ProctStatus   bool       `json:"status" bson:"status"`
	Proctoring    []DocImage `json:"proctoring" bson:"proctoring"`
}
type UserInfoWork struct {
	Place    string `json:"place" bson:"place"`
	Industry string `json:"industry" bson:"industry"`
	Position string `json:"position" bson:"position"`
	Ogrn     string `json:"ogrn" bson:"ogrn"`
	Inn      string `json:"inn" bson:"inn"`
}

type ModerationInfo struct {
	IsModerated    bool      `json:"is_moderated" bson:"is_moderated"`
	ModerationTime time.Time `json:"moderation_time" bson:"moderation_time"`
	Comment        string    `json:"comment" bson:"comment"`
}

type UserRole struct {
	Service string `json:"service" bson:"service"`
	Role    string `json:"role" bson:"role"`
}

type RobocodeInfo struct {
	AccountId   string `json:"account_id" bson:"account_id"`
	PrincipalId string `json:"principal_id" bson:"principal_id"`
	LeaderId    string `json:"leader_id" bson:"leader_id"`
	UntiId      string `json:"unti_id" bson:"unti_id"`
}
