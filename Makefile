run:
	./bin/terrace

build:
	go build -x -o ./bin/terrace ./src
	sudo setcap CAP_NET_BIND_SERVICE=+eip ./bin/terrace 
