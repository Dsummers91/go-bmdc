FROM golang:latest 

WORKDIR /go/src/github.com/dsummers91
RUN mkdir ./go-bmdc
ADD . ./go-bmdc/ 
WORKDIR ./go-bmdc 
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN dep ensure -vendor-only
RUN go build -o bmdc-server . 
CMD ["./bmdc-server"]
