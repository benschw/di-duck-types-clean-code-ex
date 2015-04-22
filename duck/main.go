package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/benschw/dns-clb-go/clb"
	"github.com/benschw/dns-clb-go/dns"
)

type AddressGetter interface {
	GetAddress(string) (dns.Address, error)
}

func Discover(lb AddressGetter, addr string) string {
	address, _ := lb.GetAddress(addr + ".service.consul")

	return fmt.Sprintf("http://%s:%d", address.Address, address.Port)
}

func main() {
	dnsServer := flag.String("dns", "127.0.1.1:53", "dns server address")
	svcName := flag.String("service-name", "", "consul service name")
	flag.Parse()

	dnsParts := strings.Split(*dnsServer, ":")
	lb := clb.NewClb(dnsParts[0], dnsParts[1], clb.Random)

	fmt.Print(Discover(lb, *svcName))
}
