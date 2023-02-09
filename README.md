# HCS Hours

## Build

GOOS=windows GOARCH=amd64 go build -o bin/hcs-hours.exe main.go
GOOS=darwin GOARCH=amd64 go build -o bin/hcs-hours-darwin main.go
GOOS=linux GOARCH=amd64 go build -o bin/hcs-hours-linux main.go