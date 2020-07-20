package setddns

import (
	"crypto/sha256"
	"encoding/hex"
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

// SetHostnameIPv4Addr set's the IPv4 Address for a particular hostname
func SetHostnameIPv4Addr(hostname string, apiKey string, sharedSecret string, WANIPv4 string, apiURL string) *Response {
	preSha256SumConcatenation := fmt.Sprintf("%s%s.%s", WANIPv4, hostname, sharedSecret)
	sha256SumOfConcatenation := sha256.Sum256([]byte(preSha256SumConcatenation))
	sha256SumAsString := hex.EncodeToString(sha256SumOfConcatenation[:])

	if !strings.HasSuffix(apiURL, "/") {
		apiURL = apiURL + "/"
	}

	constructedAPIEndpoint := fmt.Sprintf("%s?mode=set&hostname=%s&hash=%s", apiURL, hostname, sha256SumAsString)

	APIclient := &http.Client{}
	req, err := http.NewRequest("GET", constructedAPIEndpoint, nil)
	if err != nil {
		log.Fatalln(err, "Error creating set-ipv4 request before the request could be sent out")
	}
	req.Header.Add("x-api-key", apiKey)

	resp, err := APIclient.Do(req)
	if err != nil {
		log.Fatalln(err, "Error executing set-ipv4 request to DDNS API")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err, "Error reading set-ipv4 response body from DDNS API")
	}

	response := &Response{}

	jsonerr := json.Unmarshal([]byte(body), response)
	if jsonerr != nil {
		log.Fatalln(jsonerr)
	}

	return response
}
