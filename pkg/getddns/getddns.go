package getddns

import (
	"fmt"
	"strings"

	common "../common"
)

// GetPublicIPv4Addr gets the public IPv4 address of the network the device calling the function is on
func GetPublicIPv4Addr(apiKey string, apiURL string) (common.Response, error) {
	if !strings.HasSuffix(apiURL, "/") {
		apiURL = apiURL + "/"
	}

	constructedAPIEndpoint := fmt.Sprintf("%s?mode=get", apiURL)

	return common.SendHTTPRequest(constructedAPIEndpoint, apiKey)
}
