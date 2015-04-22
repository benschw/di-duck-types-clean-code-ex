#!/bin/bash

mkdir -p build

# download consul
if [ ! -f ./build/consul ]; then
	wget -O ./build/consul.zip https://dl.bintray.com/mitchellh/consul/0.5.0_linux_amd64.zip
	unzip ./build/consul.zip -d ./build/ 
	chmod 755 ./build/consul
fi 


mkdir -p ./build/consul.d
echo '{"service": {"name": "test", "port": 8080}}' \
    > ./build/consul.d/test.json

# Start Consul
./build/consul agent -server -bind 0.0.0.0 -bootstrap-expect 1 \
	-data-dir ./build/consul-data -config-dir ./build/consul.d > /dev/null &

CONSUL_PID=$!
sleep 5

go test

kill $CONSUL_PID