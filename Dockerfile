FROM golang:1.22.5-alpine AS  builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /url-shortener

FROM alpine:latest

COPY --from=builder /url-shortener /url-shortener
COPY --from=builder /app/.env .

EXPOSE 8080

CMD [ "/url-shortener" ]
