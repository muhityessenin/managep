FROM golang:1.22-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext musl-dev


#dependencies
COPY ["go.sum", "go.mod", "./"]
RUN go mod download

#build
COPY ./ ./
RUN go build -o app ./cmd

FROM alpine AS runner

COPY --from=builder /usr/local/src/app /
COPY config/config.yml config/config.yml
COPY .env ./.env

CMD ["./app"]