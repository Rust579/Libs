package gateway360

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type ErrorItem struct {
	Field   string `json:"field"`
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

func GetUser() {
	data := `{
	  "filter": {
	    "or": [
	      {
	        "condition": "size",
	        "field": "hard_skills",
	        "value": 0
	      }
	    ]
	  },
	   "pages": {
	        "page": 1,
	        "limit": 1
	    }
	}`

	req, err := http.NewRequest("POST", "http://localhost:8090/v1/get-user", strings.NewReader(data))
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Status: %v\n", response.Status)
	fmt.Printf("NeedFields: %v\n", response.NeedFields)
	fmt.Printf("Errors: %v\n", response.Errors)
	fmt.Printf("Values: %v\n", response.Values)
	fmt.Printf("TmRequest: %v\n", response.TmRequest)
	fmt.Printf("TmRequestSt: %v\n", response.TmRequestSt)
}
