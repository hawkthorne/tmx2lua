.PHONY = clean fmt

all: fmt build/darwin/tmx2lua build/linux/tmx2lua build/windows/tmx2lua

build/darwin/tmx2lua: tmx2lua.go
	mkdir -p build/darwin
	GOOS=darwin go build -o $@

build/linux/tmx2lua: tmx2lua.go
	mkdir -p build/linux
	GOOS=linux GOARCH=amd64 go build -o $@

build/windows/tmx2lua: tmx2lua.go
	mkdir -p build/windows
	GOOS=windows GOARCH=amd64 go build -o $@


fmt: 
	go fmt tmx2lua.go

clean: 
	rm -rf build
