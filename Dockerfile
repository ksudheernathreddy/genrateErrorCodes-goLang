FROM golang AS builder
WORKDIR /go/src/app
COPY ./fault.go ./
ENV GOPATH /go/src/app
ENV GOBIN $GOPATH/bin
ENV GOOS linux
ENV GOARCH amd64
ENV CGO_ENABLED 0
RUN go get -d -v ./...
RUN go install -v ./...

FROM alpine
COPY --from=builder /go/src/app/bin .
RUN chmod +x ./app
CMD ["./app"]
