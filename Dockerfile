FROM daocloud.io/library/golang:1.6.2

# Build app
ADD . $GOPATH/src/weupay

RUN go get -t weupay
RUN go install weupay

EXPOSE 80
CMD ["weupay"]
