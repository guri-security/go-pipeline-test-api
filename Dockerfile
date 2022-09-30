FROM golang:1.19.1-alpine AS go-pipeline-test-api

WORKDIR /go/src/github.com/guri-security/go-pipeline-test-api

COPY . .

RUN cd /go/src/github.com/guri-security/go-pipeline-test-api/src \
    ; go build -o go-pipeline-test-api .

FROM golang:1.19.1-alpine

COPY --from=go-pipeline-test-api /go/src/github.com/guri-security/go-pipeline-test-api/src/go-pipeline-test-api /usr/bin/go-pipeline-test-api

CMD ["/usr/bin/go-pipeline-test-api"]