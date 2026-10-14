FROM golang:1.27-alpine AS builder

WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -buildvcs=false -trimpath -o /out/fcupcntr ./cmd/fcupcntr

FROM alpine:3.22
RUN addgroup -S app && adduser -S -G app app && mkdir /data && chown app:app /data
WORKDIR /app
COPY --from=builder /out/fcupcntr /app/fcupcntr
USER app
EXPOSE 8080
ENV PORT=8080 DB_PATH=/data/counters.db MAX_COUNTERS=100
CMD ["/app/fcupcntr"]
