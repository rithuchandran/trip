FROM golang:1.12-alpine3.9 as builder

WORKDIR /code
RUN apk --update add build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . /code

RUN make compile


# Step2: Build a small image for running the program
FROM gcr.io/distroless/base:latest

COPY --from=builder /code/out/trip /go/bin/trip
COPY --from=builder /code/data /data

EXPOSE 8080
CMD ["/go/bin/trip"]