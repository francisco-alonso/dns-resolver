FROM golang:1.23

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o tld-server ./tld-server/main.go

EXPOSE 53/udp

CMD ["./tld-server"]
