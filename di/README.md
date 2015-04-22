## discover - di

`discover` is a cli app to get a valid address for a service using consul


This implementation uses dependency injection to make testing a little easier.


### example usage
_assuming you are running consul with DNSMasq to facilitate resolving `*.consul` names with your default dns server_

	discover -service-name test
	http://10.1.3.73:8080


### build
	go build -o discover

### test
	./test.sh