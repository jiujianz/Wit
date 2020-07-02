FROM golang:1.14

RUN mkdir /go/src/app

WORKDIR /go/src/app

RUN go get github.com/gin-gonic/gin
RUN go get github.com/codegangsta/gin

CMD ["gin", "--port", "8080"]

ADD . /go/src/app