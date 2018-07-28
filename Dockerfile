FROM golang:1.10-alpine as builder

WORKDIR /go/src/app
COPY . .

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app .

FROM scratch

COPY --from=builder /go/bin/app /

ENTRYPOINT ["/app"]

EXPOSE 8080
