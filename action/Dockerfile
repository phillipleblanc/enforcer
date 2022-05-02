FROM golang as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /app
COPY . .
RUN go build -o /app/main

FROM alpine
RUN apk add --no-cache ca-certificates
RUN update-ca-certificates
COPY --from=builder /app/main /bin/main
ENTRYPOINT main