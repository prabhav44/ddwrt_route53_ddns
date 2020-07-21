package common

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	ReturnMessage string `json:"return_message"`
	ReturnStatus  string `json:"return_status"`
}

func SendHTTPRequest(constructedAPIEndpoint string, apiKey string) (Response, error) {
	response := Response{}

	APIclient := &http.Client{}
	req, err := http.NewRequest("GET", constructedAPIEndpoint, nil)
	if err != nil {
		log.Println(err)
		return response, errors.New("Error creating get-ipv4-address request before the request could be sent out")
	}
	req.Header.Add("x-api-key", apiKey)

	resp, err := APIclient.Do(req)
	if err != nil {
		log.Println(err)
		return response, errors.New("Error executing get-ipv4 request to DDNS API")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return response, errors.New("Error reading get-ipv4 response body from DDNS API")
	}

	jsonerr := json.Unmarshal([]byte(body), &response)
	if jsonerr != nil {
		log.Println(jsonerr)
		return response, errors.New("Error unmarshalling JSON response body from DDNS API")
	}

	return response, nil
}
