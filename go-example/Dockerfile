FROM golang as builder
ENV GO111MODULE=on
WORKDIR /app
ADD go.mod go.sum /app/
RUN go mod download
ADD . .
RUN go build -o /app main.go

FROM gcr.io/distroless/base
EXPOSE 8080
WORKDIR /
COPY --from=builder /app /usr/bin/app
ENTRYPOINT ["/usr/bin/app"]