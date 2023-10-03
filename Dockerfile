FROM golang:1.20-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o go-mailer main.go


FROM alpine:latest

WORKDIR /app

COPY --from=builder ./app/go-mailer .
COPY environment.yaml .

CMD ["/app/go-mailer"]