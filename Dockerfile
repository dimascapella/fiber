FROM golang:1.18-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download && go mod verify
RUN go build -o main .

FROM alpine:latest

WORKDIR /root
COPY --from=builder /app/main .
COPY .env.example .
RUN mv .env.example .env
COPY wait-for.sh .
COPY start.sh .
RUN chmod +x ./start.sh

CMD [ "./main" ]