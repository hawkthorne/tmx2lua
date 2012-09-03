.PHONY = clean fmt tar

all: fmt build/darwin/tmx2lua build/windows/tmx2lua build/linux/tmx2lua

tar: all build/tmx2lua.windows.tar.gz build/tmx2lua.darwin.tar.gz build/tmx2lua.linux.tar.gz

build/%/tmx2lua: tmx2lua.go
	mkdir -p build/$*
	GOOS=$* GOARCH=amd64 go build -o $@

build/tmx2lua.%.tar.gz: build/%/tmx2lua
	cd build/$* && tar -cf tmx2lua.$*.tar.gz tmx2lua
	mv build/$*/tmx2lua.$*.tar.gz build

fmt: 
	go fmt tmx2lua.go

clean: 
	rm -rf build
