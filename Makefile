compile:
	mkdir -p out
	go build -o out/big-life-backend

test:
	go test ./...

start: compile
	./out/big-life-backend