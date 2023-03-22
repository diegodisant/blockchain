FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY . ./

RUN go mod download

FROM builder AS development

COPY --from=builder /app /app

CMD ["go", "run", "main.go"]

FROM builder AS production

COPY --from=builder /app /app

RUN go build -o /blockchain

CMD ["/blockchain"]
