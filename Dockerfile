FROM golang:1.18 AS builder
WORKDIR /app-go
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app-go ./cmd/api

FROM scratch
COPY --from=builder ./app-go/app-go /usr/bin/app-go
EXPOSE 4001
ENTRYPOINT ["/usr/bin/app-go"]