run:
	sudo ./bin/terrace

build:
	go build -x -o ./bin/terrace ./src
	sudo setcap CAP_NET_BIND_SERVICE=+eip ./bin/terrace 

update:
	/usr/local/go/bin/go build -x -o ./bin/terrace ./src
	sudo setcap CAP_NET_BIND_SERVICE=+eip ./bin/terrace

deploy:
	git push git@54.180.97.185:~/terrace_hook master
