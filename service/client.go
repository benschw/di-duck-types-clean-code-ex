package main

import (
	"fmt"
	"net/http"

	"github.com/benschw/opin-go/rest"

	"github.com/benschw/dns-clb-go/clb"
	"github.com/benschw/dns-clb-go/dns"
)

const ServiceAddress = "greeting.service.consul"

// Interface for Load Balancer
type AddressGetter interface {
	GetAddress(string) (dns.Address, error)
}

// Client Factory
func NewGreetingClient() *GreetingClient {
	return &GreetingClient{
		Lb:      clb.NewClb("localhost", "53", clb.Random),
		Address: ServiceAddress,
	}
}

// Client
type GreetingClient struct {
	Lb      AddressGetter
	Address string
}

func (c *GreetingClient) GetGreeting() ([]byte, error) {
	host, _ := c.Lb.GetAddress(c.Address)
	r, _ := rest.MakeRequest("GET", fmt.Sprintf("http://%s/greeting", host), nil)
	return rest.ProcessResponseBytes(r, http.StatusOK)
}
