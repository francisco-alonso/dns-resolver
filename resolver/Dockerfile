FROM golang:1.23

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o dns-resolver ./resolver/main.go

EXPOSE 53/udp

CMD ["./dns-resolver"]
