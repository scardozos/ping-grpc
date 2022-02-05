FROM golang:1.13-alpine

WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build ./pong.go

EXPOSE 8080

CMD [ "/pong" ]
