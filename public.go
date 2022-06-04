package getmyip

import (
	"encoding/json"
	"net/http"
)

func GetPublicIP() string {
	response, err := http.Get("http://ip-api.com/json")
	if err != nil {
		return ""
	}

	var resp = new(requestPublicIP)
	err = json.NewDecoder(response.Body).Decode(resp)
	if err != nil {
		return ""
	}
	return resp.Query
}

type requestPublicIP struct {
	Query  string `json:"query"`
	Status string `json:"success"`
}
