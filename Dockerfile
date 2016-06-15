FROM daocloud.io/library/golang:1.6.2

# Build app
ADD . $GOPATH/src/weupay

WORKDIR $GOPATH/src/weupay
RUN go get -t weupay
RUN go build

EXPOSE 8080
CMD $GOPATH/src/weupay/weupay
