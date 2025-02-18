FROM golang:1.23 AS builder

WORKDIR /api

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go test -count=1 -v github.com/TienMinh25/delivery-system/internal/notifications/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o notifications cmd/notifications/main.go

FROM scratch

COPY --from=builder /api/notifications /

EXPOSE 4402

# run binary api file when container run
ENTRYPOINT ["/notifications"]