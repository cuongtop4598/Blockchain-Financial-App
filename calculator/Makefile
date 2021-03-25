gen-cal:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ~/Desktop/Golang/src/github.com/cuongtop4598/calculator/calculator.proto
run-server:
	go run server/server.go
run-client:
	go run client/client.go
