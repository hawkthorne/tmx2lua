.PHONY = clean fmt tar upload

all: fmt osx/tmx2lua.osx.zip linux/tmx2lua.linux.tar.gz linux64/tmx2lua.linux64.tar.gz \
	windows32/tmx2lua.win32.zip windows64/tmx2lua.win64.zip

GOPATH=$(HOME)/gopath
export GOPATH

install:
	go get github.com/kyleconroy/go-tmx/tmx


linux64/tmx2lua.linux64.tar.gz: tmx2lua.go
	mkdir -p linux64
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o linux64/tmx2lua
	cd linux64 && tar -cvzf tmx2lua.linux64.tar.gz tmx2lua

linux/tmx2lua.linux.tar.gz: tmx2lua.go
	mkdir -p linux
	GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o linux/tmx2lua
	cd linux && tar -cvzf tmx2lua.linux.tar.gz tmx2lua

osx/tmx2lua.osx.zip: tmx2lua.go
	mkdir -p osx
	GOOS=darwin GOARCH=amd64 go build -o osx/tmx2lua
	cd osx && zip tmx2lua.osx.zip tmx2lua

windows64/tmx2lua.win64.zip: tmx2lua.go
	mkdir -p windows64
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o windows64/tmx2lua.exe
	cd windows64 && zip tmx2lua.win64.zip tmx2lua.exe

windows32/tmx2lua.win32.zip: tmx2lua.go
	mkdir -p windows32
	GOOS=windows GOARCH=386 CGO_ENABLED=0 go build -o windows32/tmx2lua.exe
	cd windows32 && zip tmx2lua.win32.zip tmx2lua.exe

fmt: 
	go fmt tmx2lua.go

clean: 
	go clean
	rm -rf windows32
	rm -rf windows64
	rm -rf osx
	rm -rf linux
	rm -rf linux64
