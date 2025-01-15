FROM golang:1.23

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o auth-server ./auth-server/main.go

EXPOSE 53/udp

CMD ["./auth-server"]
