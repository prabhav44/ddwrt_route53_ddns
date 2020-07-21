package main

import (
	"log"

	flags "../internal/flags"
	getddns "../pkg/getddns"
	setddns "../pkg/setddns"
)

func main() {
	var operation flags.Operation = flags.GetFlags()
	switch operation.Name {
	case "set":
		getIPv4response, err := getddns.GetPublicIPv4Addr(operation.Parameters.APIkey, operation.Parameters.APIURL)
		if err != nil {
			log.Fatalln(err)
		}
		var ip string = getIPv4response.ReturnMessage
		setIPv4response, err := setddns.SetHostnameIPv4Addr(operation.Parameters.Hostname, operation.Parameters.APIkey, operation.Parameters.SharedSecret, ip, operation.Parameters.APIURL)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(setIPv4response.ReturnStatus)
		log.Println(setIPv4response.ReturnMessage)
	case "get":
		getIPv4response, err := getddns.GetPublicIPv4Addr(operation.Parameters.APIkey, operation.Parameters.APIURL)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(getIPv4response.ReturnStatus)
		log.Println(getIPv4response.ReturnMessage)
	}
}
