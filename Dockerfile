FROM golang:1.9
LABEL maintainer="web@rafalgolarz.com"
COPY *.go /go/src/github.com/rafalgolarz/passgen/
COPY passwords/*.go /go/src/github.com/rafalgolarz/passgen/passwords/
COPY passwords/*.toml /go/src/github.com/rafalgolarz/passgen/passwords/

RUN go get github.com/gin-gonic/gin
RUN go get github.com/sirupsen/logrus
RUN go get github.com/BurntSushi/toml
WORKDIR /go/src/github.com/rafalgolarz/passgen

RUN go build
RUN chmod +x passgen

ENV DEFAULT_API_PORT=:8080
ENV GIN_MODE=release

CMD ./passgen
EXPOSE 8080