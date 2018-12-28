setup:
	go get github.com/golang/dep
	dep init

compile:
	mkdir -p out
	go build -o out/main

build-deps:
	dep ensure

start: build-deps compile
	./out/main