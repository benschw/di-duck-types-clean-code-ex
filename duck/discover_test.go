package main

import (
	"fmt"
	"testing"

	"github.com/benschw/dns-clb-go/dns"
)

type StaticAddressGetter struct {
	Val dns.Address
}

func (lb *StaticAddressGetter) GetAddress(address string) (dns.Address, error) {
	if address == "test.service.consul" {
		return lb.Val, nil
	}
	return dns.Address{}, fmt.Errorf("invalid service name")
}

func TestDiscover(t *testing.T) {
	//given
	expected := "http://foo:8080"
	lb := &StaticAddressGetter{Val: dns.Address{Address: "foo", Port: 8080}}

	// when
	found := Discover(lb, "test")

	// then
	if found != expected {
		t.Errorf("\"%s\" not equal to \"%s\"", found, expected)
	}
}
