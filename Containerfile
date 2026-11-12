FROM golang:1.27-alpine AS builder

WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -buildvcs=false -trimpath -o /out/fcupcntr ./cmd/fcupcntr

FROM alpine:3.22
RUN mkdir /data && chown 999:999 /data
WORKDIR /app
COPY --from=builder /out/fcupcntr /app/fcupcntr
USER 999:999
EXPOSE 8080
ENV PORT=8080 DSN=/data/counters.db MAX_COUNTERS=100
CMD ["/app/fcupcntr"]
