#!/bin/bash
NAME="ssltool"
CGO_ENABLED=0 go build -o build/${NAME}_linux_amd64
CGO_ENABLED=0 GOOS=windows go build -o build/${NAME}_win_amd64.exe
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/${NAME}_mac_amd64
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o build/${NAME}_mac_arm64
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o build/${NAME}_linux_arm
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o build/${NAME}_linux_arm64

find build -name "*" -exec strip {} \;