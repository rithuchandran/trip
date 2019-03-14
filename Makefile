compile:
	mkdir -p out
	git config --global core.eol lf
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o out/trip --tags netgo --ldflags '-extldflags "-lm -lstdc++ -static"'

run: compile
	go run .

test:
	go test ./...

start: compile
	./out/trip
