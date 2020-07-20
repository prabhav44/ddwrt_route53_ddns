package main

import (
	"fmt"

	flags "../internal/flags"
	publicip "../pkg/publicip"
	setddns "../pkg/setddns"
)

func main() {
	var operation flags.Operation = flags.GetFlags()
	switch operation.Name {
	case "set":
		ip := publicip.GetPublicIP()
		setddns.SetHostnameIPv4Addr(operation.Parameters.Hostname, operation.Parameters.APIkey, operation.Parameters.SharedSecret, ip, operation.Parameters.APIURL)
	case "get":
		fmt.Println("start get workflow")
	}
}
