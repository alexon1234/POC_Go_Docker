FROM golang as builder
ENV GO111MODULE=on
WORKDIR /app
ADD . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch
COPY --from=builder /app /usr/bin/app
ENTRYPOINT ["/usr/bin/app"]