package setddns

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

// SetHostnameIPv4Addr set's the IPv4 Address for a particular hostname
func SetHostnameIPv4Addr(hostname string, apiKey string, sharedSecret string, WANIPv4 string, apiURL string) {
	preSha256SumConcatenation := fmt.Sprintf("%s%s.%s", WANIPv4, hostname, sharedSecret)
	sha256SumOfConcatenation := sha256.Sum256([]byte(preSha256SumConcatenation))
	sha256SumAsString := hex.EncodeToString(sha256SumOfConcatenation[:])

	if !strings.HasSuffix(apiURL, "/") {
		apiURL = apiURL + "/"
	}

	constructedAPIEndpoint := fmt.Sprintf("%s?mode=set&hostname=%s&hash=%s", apiURL, hostname, sha256SumAsString)
	fmt.Println(constructedAPIEndpoint)
}
