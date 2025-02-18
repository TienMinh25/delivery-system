FROM golang:1.23 AS builder

WORKDIR /api

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go test -count=1 -v github.com/TienMinh25/delivery-system/internal/gateway
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api cmd/api/main.go

FROM scratch

COPY --from=builder /api/api /

EXPOSE 4401

# run binary api file when container run
ENTRYPOINT ["/api"]