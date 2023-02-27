build-win32:
	env GOOS=windows GOARCH=386 go build -o bin/server_win32.exe

build-win64:
	env GOOS=windows GOARCH=amd64 go build -o bin/server_win64.exe

build-linux:
	env GOOS=linux go build -o bin/server_linux

build-macos:
	env GOOS=darwin go build -o bin/server_macos

build-all: build-win32 build-win64 build-linux build-macos
