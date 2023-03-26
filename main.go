package main

import (
	"flag"
	"log"
	"net"
	"strings"
)

type DnsData struct {
	DNS   string   `json:"dns"`
	CName string   `json:"cname"`
	IpsV4 []string `json:"ips_v4"`
	IpsV6 []string `json:"ips_v6"`
}

func main() {
	dnsPtr := flag.String("dns", "", "DNS name to look up")
	flag.Parse()

	if *dnsPtr == "" {
		log.Fatal("Missing DNS name argument")
	}

	incomingDns := strings.TrimSpace(*dnsPtr)
	ips, err := net.LookupIP(incomingDns)
	if err != nil {
		log.Fatalf("Error looking up IP address for %s: %v", incomingDns, err)
	}

	data := &DnsData{
		DNS:   incomingDns,
		CName: "",
		IpsV4: make([]string, 0),
		IpsV6: make([]string, 0),
	}

	data.CName, _ = net.LookupCNAME(data.DNS)

	for _, ip := range ips {
		if ip.To4() != nil {
			data.IpsV4 = append(data.IpsV4, ip.String())
		} else if ip.To16() != nil {
			data.IpsV6 = append(data.IpsV6, ip.String())
		}
	}

	log.Printf("DNS:   %s", data.DNS)
	log.Printf("CName: %s", data.CName)
	log.Printf("IpsV4: %v", data.IpsV4)
	log.Printf("IpsV6: %v", data.IpsV6)
}
