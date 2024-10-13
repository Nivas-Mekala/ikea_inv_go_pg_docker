FROM golang:1.22-alpine as builder
WORKDIR /app
COPY go.* .
RUN go mod download

COPY . .

WORKDIR cmd/app
RUN CGO_ENABLES=0 GOOSE=linux go build -mod=readonly -v -o main

FROM alpine:latest
RUN apk add --no-cache ca-certificates

ARG SERVICE
COPY --from=builder /app/cmd/app/main /main
COPY --from=builder /app/.env .

CMD [ "/main" ]