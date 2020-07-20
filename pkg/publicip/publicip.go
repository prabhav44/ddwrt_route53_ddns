package publicip

import (
	"io/ioutil"
	"log"
	"net/http"
)

//GetPublicIP contacts the ipify API and obtains the public IP of the WAN the device running this program is a member of
func GetPublicIP() string {
	res, err := http.Get("https://api.ipify.org")
	if err != nil {
		log.Fatalln("Unable to obtain public IP from ipify API")
	}
	defer res.Body.Close()
	ip, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("Unable to parse public IP obtained from ipify API")
	}
	return string(ip)
}
