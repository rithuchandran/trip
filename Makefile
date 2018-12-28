setup:
	pwd
	go get -u github.com/golang/dep/cmd/dep
	dep init

compile:
	mkdir -p out
	go build -o out/main

build-deps:
	dep ensure

start: build-deps compile
	pwd
	./out/main