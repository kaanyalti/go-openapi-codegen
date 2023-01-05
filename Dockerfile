FROM golang:1.19-alpine as base
WORKDIR /app
COPY . .

FROM base as build
RUN go build -o build main.go

CMD ["/app/build"]