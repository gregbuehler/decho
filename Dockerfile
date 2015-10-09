FROM golang

ADD . /go/src/github.com/gregbuehler/decho

RUN go install github.com/gregbuehler/decho

ENTRYPOINT /go/bin/decho

EXPOSE 8765
