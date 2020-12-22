thrift-gen:
	./develop/gen-script.sh
run-client:
	go run ./client/main.go
run-server:
	go run ./server/main.go