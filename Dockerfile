FROM golang:1.16-alpine AS builder

WORKDIR $GOPATH/src/go-api/

COPY . .

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/go-api

FROM scratch

WORKDIR $GOPATH/src/go-api/

COPY --from=builder /go/bin/go-api /go/bin/go-api

ENTRYPOINT ["/go/bin/go-api"]