FROM golang:1.10
LABEL maintainer="web@rafalgolarz.com"

COPY cmd/passgend/  /go/src/github.com/rafalgolarz/passgen/cmd/passgend/
COPY pkg/  /go/src/github.com/rafalgolarz/passgen/pkg/
COPY vendor/ /go/src/
WORKDIR /go/src/github.com/rafalgolarz/passgen/cmd/passgend

#
# -ldflags "-X main.build=" sets a build number to the build timestamp (UTC)
# build number is available from /health request

RUN go build -ldflags "-X main.build=`date -u +%Y%m%d.%H%M%S`.UTC"
RUN ls -al
RUN chmod +x passgend

ENV GIN_MODE=release
ENV DEFAULT_API_PORT=:8080

CMD ./passgend
EXPOSE 8080
