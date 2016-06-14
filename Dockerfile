FROM golang:1.5.1

# Build app
ADD . $GOPATH/src/weupay

RUN go get -t weupay
RUN go install weupay

EXPOSE 80
CMD ["weupay"]
