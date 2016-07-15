FROM golang

ENV GOBIN /go/bin

ADD ./src/service1.go /go/src/github.com/twinfield/service1/

RUN cd /go/src/github.com/twinfield/service1 && \
  go get && \
  go install service1.go

ENTRYPOINT /go/bin/service1

EXPOSE 9090
