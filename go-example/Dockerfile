FROM golang:alpine
RUN mkdir /app
WORKDIR /app
CMD ["ls"]
COPY go.mod .
COPY go.sum .

RUN go mod download
ADD . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app

FROM scratch
COPY --from=build-env /go/bin/app /go/bin/app
ENTRYPOINT ["/go/bin/app"]