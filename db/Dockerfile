FROM golang:1.18-alpine as base
WORKDIR /app
COPY . .

FROM base as build
RUN go build -o migrate migrate.go

CMD ["/app/migrate"]