FROM golang

ADD . /go/src/github.com/sausheong/ws-d
WORKDIR /go/src/github.com/sausheong/ws-d

RUN go get github.com/lib/pq
RUN go install github.com/sausheong/ws-d

ENTRYPOINT /go/bin/ws-d

EXPOSE 8080