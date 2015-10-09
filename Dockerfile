FROM golang

ADD . /go/src/gitlab.com/gregbuehler/decho

RUN go install gitlab.com/gregbuehler/decho

ENTRYPOINT /go/bin/decho

EXPOSE 8765
