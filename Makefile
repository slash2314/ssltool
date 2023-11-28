VERSION := $$(cat VERSION)
all: build strip
build: ssltool_linux_amd64 ssltool_linux_arm64 ssltool_linux_arm ssltool_linux_arm64 ssltool_mac_amd64 ssltool_mac_arm64 ssltool_mac_amd64 ssltool_windows_amd64.exe ssltool_windows_arm64.exe
ssltool_linux_amd64:
	CGO_ENABLED=0 go build -o build/ssltool_linux_amd64_$(VERSION)
ssltool_windows_amd64.exe:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build/ssltool_win_amd64_$(VERSION).exe
ssltool_windows_arm64.exe:
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -o build/ssltool_win_arm64_$(VERSION).exe
ssltool_mac_amd64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/ssltool_mac_amd64_$(VERSION)
ssltool_mac_arm64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o build/ssltool_mac_arm64_$(VERSION)
ssltool_linux_arm:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o build/ssltool_linux_arm_$(VERSION)
ssltool_linux_arm64:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o build/ssltool_linux_arm64_$(VERSION)

strip:
	find build -name "build/*" -exec strip {} \;

zip:
	cd build
	zip -r build/ssltool.zip build/*

prod: build strip zip

clean:
	rm build/*
