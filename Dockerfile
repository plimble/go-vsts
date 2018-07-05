FROM golang:alpine

RUN apk --update upgrade && \
   apk add git ca-certificates && \
   update-ca-certificates && \
   rm -rf /var/cache/apk/*

WORKDIR /go/src/app
COPY . /go/src/app
RUN go build -o app .

CMD ["./app"]
