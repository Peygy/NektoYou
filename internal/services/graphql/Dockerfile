FROM golang:alpine AS builder

WORKDIR /app

ADD go.mod .

COPY . .

RUN go build -o main ./cmd


FROM alpine

WORKDIR /build

COPY --from=builder /app/main /build/main
COPY --from=builder /app/config/config.dev.yml /build/config/config.dev.yml

EXPOSE 4000

CMD ["/build/main"]
# docker build -t graphql-server .
# docker run -d -p 4000:4000 --name graphql-server graphql-server