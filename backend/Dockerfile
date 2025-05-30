FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Tell Go where the module root is (WORKDIR=/app and module=operary)
ENV GO111MODULE=on

RUN go build -o operary ./main.go

CMD ["./operary"]
