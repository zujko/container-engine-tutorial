FROM golang:1.8.3-alpine3.6

RUN apk add --no-cache git

WORKDIR /go/src/mini-url
COPY . .

RUN go-wrapper download
RUN go-wrapper install
CMD ["go-wrapper", "run"]