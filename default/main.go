package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/benschw/dns-clb-go/clb"
)

func main() {
	dnsServer := flag.String("dns", "127.0.1.1:53", "dns server address")
	svcName := flag.String("service-name", "", "consul service name")
	flag.Parse()

	dnsParts := strings.Split(*dnsServer, ":")
	lb := clb.NewClb(dnsParts[0], dnsParts[1], clb.Random)

	address, _ := lb.GetAddress(*svcName + ".service.consul")

	fmt.Printf("http://%s:%d", address.Address, address.Port)
}
