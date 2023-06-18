#!/bin/bash
NAME="ssltool"
VERSION=`cat VERSION`
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/${NAME}_linux_amd64_${VERSION}
CGO_ENABLED=0 GOOS=windows go build -o build/${NAME}_win_amd64_${VERSION}.exe
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/${NAME}_mac_amd64_${VERSION}
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o build/${NAME}_mac_arm64_${VERSION}
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o build/${NAME}_linux_arm_${VERSION}
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o build/${NAME}_linux_arm64_${VERSION}

#find build -name "*" -exec strip {} \;
