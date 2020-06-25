package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

func main() {
	interfaceArgPtr := flag.String("interface", "vlan2", "Enter the name of the interface you would like to grab your WAN ip address from")
	flag.Parse()

	WANipInterface, err := net.InterfaceByName(*interfaceArgPtr)
	if err != nil {
		log.Fatalln(err)
	}

	WANip, err := WANipInterface.Addrs()
	if err != nil {
		log.Fatalln(err)
	} else if len(WANip) > 1 {
		log.Fatalln("Interface provided has more than one IP address associated with it")
	}

	// The WAN interface (vlan2 on my dd wrt instance) should only have a single address
	// here we take the single IP address and convert it to a string
	var primaryWANip string = WANip[0].String()

	// Parse the CIDR string into the IP address itself and the network separate
	// Note that we don't need the subnet value for this
	WANipv4Addr, _, err := net.ParseCIDR(primaryWANip)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(WANipv4Addr)
}
