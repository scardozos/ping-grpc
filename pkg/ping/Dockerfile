FROM golang:1.13-alpine

ENV PORT=8080

WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build ping.go 

EXPOSE 8080

CMD [ "./ping" ]

