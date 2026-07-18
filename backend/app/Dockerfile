FROM golang:1.22.1-alpine3.18 as builder

WORKDIR /app
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

COPY backend/app/go.mod backend/app/go.sum ./
RUN go mod download

COPY backend/app/ .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates curl
RUN apk add --no-cache tzdata
RUN adduser -D -g '' golang

WORKDIR /app


COPY --from=builder /app/main .
COPY --from=builder /app/templates ./templates
COPY  backend/app/migrations ./migrations
COPY  backend/app/docs/swagger.json ./docs/swagger.json

USER golang

ENV APP_ENV stg
ENV APP_HOST 0.0.0.0

EXPOSE 8080

CMD ["/app/main"]
