ARG GO_VERSION=1.16.6
FROM golang:${GO_VERSION}-alpine AS builder

RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

RUN apk update && \
    apk add --no-cache git ca-certificates && \
    update-ca-certificates

# For faster builds
ENV CGO_ENABLED=0

WORKDIR /src
RUN unset GOPATH

# install dependencies
COPY ./go.mod .
COPY ./go.sum .

RUN go mod download
RUN go mod verify

COPY . .

RUN go test ./...
RUN go build \
    -installsuffix 'static' \
    -o /app

RUN go get github.com/go-swagger/go-swagger/cmd/swagger
RUN mkdir /doc
RUN swagger generate spec -o /doc/swagger.json

RUN mkdir /migrations
COPY ./migrations /migrations

FROM scratch AS final
COPY --from=builder /user/group /user/passwd /etc/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app /app
COPY --from=builder /doc /doc
COPY --from=builder /migrations /migrations

EXPOSE 8080

USER nobody:nobody

ENTRYPOINT ["/app"]
