# builder
FROM golang:1.14.4-alpine3.12 as builder

RUN mkdir -p /go/src/github.com/magicalbanana/bentobox/

WORKDIR /go/src/github.com/magicalbanana/bentobox/

COPY . .

RUN apk add --update --no-cache alpine-sdk git

RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -mod vendor -v -a -installsuffix cgo -o bentobox-cli \
    main.go

# actual container
FROM alpine:3.12

RUN apk add --update --no-cache bash ca-certificates

RUN mkdir -p /etc/bentobox/bin

COPY --from=builder /go/src/github.com/magicalbanana/bentobox/bentobox-cli /etc/bentobox/bin/

# add the binary to path
ENV PATH="/etc/bentobox/bin:${PATH}"

# this is the apps directory
RUN mkdir -p /app

ENTRYPOINT ["bentobox-cli"]
