FROM golang:1.18 AS builder
WORKDIR /app-go
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app-go ./

FROM scratch
COPY --from=builder ./app-go/app-go /usr/bin/app-go
EXPOSE 8000
ENTRYPOINT ["/usr/bin/app-go"]