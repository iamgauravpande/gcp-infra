# BINARY_NAME=gcp-infra-golang
 .DEFAULT_GOAL=run

# build:
    
# 	GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux-amd64 cmd/main/main.go

run:
	go run cmd/main/main.go
 
# clean:
# 	go clean
# 	rm -rf bin/${BINARY_NAME}-linux-amd64