# Hello Fiber
Go application in Go Fiber Framework

## Requirements
- Go 1.25
- Fiber 

## Installation
- Create module
```
go mod init HelloFiber
```
- Install Fiber
```
go get github.com/gofiber/fiber/v2
```
- Install air (Hot Reload) and configure

Download air executable and copy in go PATH
```
go install github.com/air-verse/air@latest
```
or in ./bin folder
```
curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
```
- Initialize
```
air init
```
- Configure
Add ./srv in final of this line in .air.toml file
```
cmd = "go build -o ./tmp/main ./src"
```
- Loads env vars from a .env file
```
go get github.com/joho/godotenv
```
## Formatter
```
gofmt -w ./src 
```
## Run
```
go run src/main.go 
```
## Hot Reload
```
air
```
## Build
```
go build -o HelloFiber ./src/main.go
```