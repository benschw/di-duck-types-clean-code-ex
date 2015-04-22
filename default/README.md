## discover - default

`discover` is a cli app to get a valid address for a service using consul


This implementation is not testable. a consul server must be run and cli output from the artifact must be analyzed.

### example usage
_assuming you are running consul with DNSMasq to facilitate resolving `*.consul` names with your default dns server_

	discover -service-name test
	http://10.1.3.73:8080


### build
	go build -o discover

### test
	./test.sh