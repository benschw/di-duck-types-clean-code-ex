package main

import "flag"

func main() {
	bind := flag.String("bind", "0.0.0.0:8080", "address to bind http server to")
	flag.Parse()

	RunServer(*bind)
}
