# DOCKERFILE 

FROM golang:alpine AS builder

WORKDIR /go/src/app
COPY . .

RUN go build -o tester

FROM alpine
WORKDIR /app
COPY --from=builder /go/src/app /app/
CMD ["./tester"]


