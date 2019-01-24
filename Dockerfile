FROM golang

ENV ROOT_PACKAGE=gitlab.com/proemergotech/yafuds-go-server

ADD . $GOPATH/src/$ROOT_PACKAGE

WORKDIR $GOPATH/src/$ROOT_PACKAGE

RUN go build -o loadbalance-test .

RUN chmod +x loadbalance-test

CMD ./loadbalance-test $(hostname -I)
