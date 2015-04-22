## discover - duck

`discover` is a cli app to get a valid address for a service using consul


This implementation uses dependency injection and duck typing to facilitate injecting a mock dns resolver in tests. This makes the application much easier to test and allows us to only test our app, not the dns library we are using or Consul itself.


### example usage
_assuming you are running consul with DNSMasq to facilitate resolving `*.consul` names with your default dns server_

	discover -service-name test
	http://10.1.3.73:8080


### build
	go build -o discover

### test
	go test