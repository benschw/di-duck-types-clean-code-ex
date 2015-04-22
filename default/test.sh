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

go build -o ./build/discover || exit 1


IP=`ifconfig | grep "inet addr" | awk -F: '{print $2}' | awk '{print $1}' | head -n 1`

EXPECTED=http://$IP:8080
FOUND=`./build/discover -dns localhost:8600 -service-name test`

kill $CONSUL_PID

if [ "$EXPECTED" == "$FOUND" ]; then
	echo pass
else
	echo fail: $EXPECTED not equal to $FOUND
	exit 1
fi


