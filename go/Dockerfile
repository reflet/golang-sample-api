FROM golang:1.11.4

RUN apt-get update && \
    apt-get install -y vim less

# copy application code from host.
ADD src /go/src
WORKDIR /go/src/api

# install package management tool (dep).
ARG DEP_RELEASE_TAG=v0.5.0
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# install packages. (vendoring: /go/bin/api/vendor/)
RUN dep ensure -v -vendor-only=true

# install development tools.
RUN go install api/vendor/github.com/goadesign/goa/goagen

# build
RUN go build -o server

EXPOSE 8080
CMD ["./server"]
