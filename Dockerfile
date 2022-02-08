FROM golang:1.17

WORKDIR /go/src/github.com/rdbell/fileserver-httpbasicauth
COPY . .

RUN go install

CMD ["fileserver-httpbasicauth"]
