FROM golang:1.23 AS builder

WORKDIR /api

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go test -count=1 -v github.com/TienMinh25/delivery-system/internal/orders/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o orders cmd/orders/main.go

FROM scratch

COPY --from=builder /api/orders /

EXPOSE 4403

# run binary api file when container run
ENTRYPOINT ["/orders"]