FROM golang:1.19-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download && go mod verify
RUN go build -o .bin/build main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/.bin/build .

CMD [ "/app/.bin/build" ]