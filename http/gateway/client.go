package gateway

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/unione-pro/core/pkg/libs/httpclient"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ProctoringResultV2 struct {
	Id        primitive.ObjectID           ` json:"_id" bson:"_id,omitempty"`
	ProctoID  string                       `json:"procto_id" bson:"procto_id,omitempty"`
	Type      string                       `json:"type" bson:"type"`
	UserID    primitive.ObjectID           `json:"user_id" bson:"user_id"`
	CourseID  primitive.ObjectID           `json:"course_id" bson:"course_id"`
	StageID   primitive.ObjectID           `json:"stage_id" bson:"stage_id"`
	Score     float64                      `json:"score" bson:"score"`
	Created   time.Time                    `json:"created_at" bson:"created_at"`
	Updated   time.Time                    `json:"updated_at" bson:"updated_at"`
	UpdAuthor *ProctoringResultUpdAuthorV2 `json:"author_updated,omitempty" bson:"author_updated,omitempty"`
}

type ProctoringResultUpdAuthorV2 struct {
	Id            primitive.ObjectID `json:"id_auth" bson:"id_auth"`
	EventDateTime time.Time          `json:"date" bson:"date"`
}

type GwayWriteResponse struct {
	Status bool          `json:"status"`
	Errors []Error       `json:"errors"`
	Values []interface{} `json:"values"`
}

type Error struct {
	Key     int    `json:"key"`
	Message string `json:"message"`
}

const (
	WriteProctoringResults = "v2/write-proctoring-results"
)

var client *httpclient.FastHttpClient

func Init() {
	client = httpclient.NewFastHttpClient("http://0.0.0.0:8090/")
}

func WriteProctoringResultsV2(proctoringResult *ProctoringResultV2) error {
	return sendWriteRequest(WriteProctoringResults, []ProctoringResultV2{*proctoringResult})
}

func sendWriteRequest(url string, req interface{}) error {

	respBody, err := client.SendGetRequest(url, req)
	if err != nil || respBody == nil {
		fmt.Println(err)
		fmt.Println(respBody)
		return errors.New("error: GATEWAY endpoint " + url + ": " + err.Error())
	}

	out := GwayWriteResponse{}

	if err = json.Unmarshal(*respBody, &out); err != nil {
		fmt.Println(err)
		return errors.New("error to send item to gway: " + err.Error())
	}

	if out.Status == false {
		if len(out.Errors) != 0 {
			fmt.Println(err)
			return errors.New(out.Errors[0].Message)
		}
		fmt.Println(err)
		return errors.New("error: GATEWAY status 'false' but no error")
	}
	fmt.Println("--------------")
	fmt.Println(out)
	fmt.Println("--------------")
	return nil
}
