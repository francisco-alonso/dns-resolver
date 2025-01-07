FROM golang:1.23

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o root-server ./root-server/main.go

EXPOSE 53/udp

CMD ["./root-server"]