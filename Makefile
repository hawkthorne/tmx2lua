.PHONY = clean fmt tar upload

all: fmt build/osx/tmx2lua build/windows32/tmx2lua.exe build/windows64/tmx2lua.exe build/linux/tmx2lua build/linux64/tmx2lua

tar: all build/tmx2lua.osx.tar build/tmx2lua.linux.tar build/tmx2lua.linux64.tar build/tmx2lua.windows32.zip build/tmx2lua.windows64.zip

build/osx/tmx2lua: tmx2lua.go
	mkdir -p build/osx
	GOOS=darwin GOARCH=amd64 go build -o $@

build/linux64/tmx2lua: tmx2lua.go
	mkdir -p build/linux64
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $@

build/linux/tmx2lua: tmx2lua.go
	mkdir -p build/linux
	GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o $@

build/windows64/tmx2lua.exe: tmx2lua.go
	mkdir -p build/windows64
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o $@

build/windows32/tmx2lua.exe: tmx2lua.go
	mkdir -p build/windows32
	GOOS=windows GOARCH=386 CGO_ENABLED=0 go build -o $@

build/tmx2lua.%.zip: build/%/tmx2lua.exe
	cd build/$* && zip tmx2lua.$*.zip tmx2lua.exe
	mv build/$*/tmx2lua.$*.zip build

build/tmx2lua.%.tar: build/%/tmx2lua
	cd build/$* && tar -cf tmx2lua.$*.tar tmx2lua
	mv build/$*/tmx2lua.$*.tar build

fmt: 
	go fmt tmx2lua.go

upload: tar
	python upload.py build/tmx2lua.linux.tar
	python upload.py build/tmx2lua.linux64.tar
	python upload.py build/tmx2lua.osx.tar
	python upload.py build/tmx2lua.windows32.zip
	python upload.py build/tmx2lua.windows64.zip

clean: 
	go clean
	rm -rf build

install: all
	cp build/darwin/tmx2lua ~/bin
