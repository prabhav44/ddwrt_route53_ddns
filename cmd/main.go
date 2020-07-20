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
		var ip string = getddns.GetPublicIPv4Addr(operation.Parameters.APIkey, operation.Parameters.APIURL).ReturnMessage
		response := setddns.SetHostnameIPv4Addr(operation.Parameters.Hostname, operation.Parameters.APIkey, operation.Parameters.SharedSecret, ip, operation.Parameters.APIURL)
		log.Println(response.ReturnStatus)
		log.Println(response.ReturnMessage)
	case "get":
		response := getddns.GetPublicIPv4Addr(operation.Parameters.APIkey, operation.Parameters.APIURL)
		log.Println(response.ReturnStatus)
		log.Println(response.ReturnMessage)
	}
}
