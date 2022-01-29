package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

type DnsData struct {
	DNS   string   `json:"dns"`
	CName string   `json:"cname"`
	IpsV4 []string `json:"ips_v4"`
	IpsV6 []string `json:"ips_v6"`
}

func main() {

	input := os.Args[1]
	incomingDns := strings.TrimSuffix(input, "\n")
	ips, err := net.LookupIP(incomingDns)
	if err != nil {
		panic(err)
	}

	data := new(DnsData)
	for _, ip := range ips {
		if ip.To4() != nil {
			data.IpsV4 = append(data.IpsV4, ip.String())
		} else if ip.To16() != nil {
			data.IpsV6 = append(data.IpsV6, ip.String())
		}
	}

	data.DNS = incomingDns
	data.CName, _ = net.LookupCNAME(data.DNS)

	fmt.Printf("DNS:   %s \n", data.DNS)
	fmt.Printf("CName: %s \n", data.CName)
	fmt.Printf("IpsV4: %s \n", data.IpsV4)
	fmt.Printf("IpsV6: %s", data.IpsV6)
}
