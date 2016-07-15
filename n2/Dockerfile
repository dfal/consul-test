FROM golang

ENV GOBIN /go/bin

ADD ./src /go/src/github.com/twinfield/service2

RUN cd /go/src/github.com/twinfield/service2 && \
	go get && \
	go install service2.go

ENTRYPOINT /go/bin/service2
