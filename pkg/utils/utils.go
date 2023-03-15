package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error while Reading the request Body %v", err)
		return
	}
	err = json.Unmarshal([]byte(body), x)
	if err != nil {
		fmt.Printf("Error while Unmarshalling %v", err)
		return
	}
}
