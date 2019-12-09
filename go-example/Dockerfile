FROM golang as builder
ENV GO111MODULE=on
WORKDIR /app
ADD go.mod go.sum /app/
RUN go mod download
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch
EXPOSE 8080
WORKDIR /
COPY --from=builder /app /usr/bin/app
ENTRYPOINT ["/usr/bin/app"]