all: build strip
build: ssltool_linux_amd64 ssltool_linux_arm64 ssltool_linux_arm ssltool_mac_amd64 ssltool_mac_arm64 ssltool_mac_amd64
ssltool_linux_amd64:
	CGO_ENABLED=0 go build -o build/ssltool_linux_amd64
ssltool_windows_amd64.exe:
	CGO_ENABLED=0 GOOS=windows go build -o build/ssltool_win_amd64.exe
ssltool_mac_amd64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/ssltool_mac_amd64
ssltool_mac_arm64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o build/ssltool_mac_arm64
ssltool_linux_arm:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o build/ssltool_linux_arm
ssltool_linux_arm64:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o build/ssltool_linux_arm64

strip:
	find build -name "*" -exec strip {} \;

zip:
	cd build
	zip -r build/ssltool.zip build/*

prod: build strip zip
