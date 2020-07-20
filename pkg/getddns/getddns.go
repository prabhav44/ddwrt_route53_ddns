package getddns

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Response struct {
	ReturnMessage string `json:"return_message"`
	ReturnStatus  string `json:"return_status"`
}

func GetPublicIPv4Addr(apiKey string, apiURL string) *Response {
	if !strings.HasSuffix(apiURL, "/") {
		apiURL = apiURL + "/"
	}

	constructedAPIEndpoint := fmt.Sprintf("%s?mode=get", apiURL)

	APIclient := &http.Client{}
	req, err := http.NewRequest("GET", constructedAPIEndpoint, nil)
	if err != nil {
		log.Fatalln(err, "Error creating get-ipv4-address request before the request could be sent out")
	}
	req.Header.Add("x-api-key", apiKey)

	resp, err := APIclient.Do(req)
	if err != nil {
		log.Fatalln(err, "Error executing get-ipv4 request to DDNS API")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err, "Error reading get-ipv4 response body from DDNS API")
	}

	response := &Response{}

	jsonerr := json.Unmarshal([]byte(body), response)
	if jsonerr != nil {
		log.Fatalln(jsonerr)
	}

	return response
}
