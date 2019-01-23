FROM golang:1.11.4

# locale & timezone (Asia/Tokyo)
# https://github.com/moby/moby/issues/12084
ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

# system update
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
RUN go install api/vendor/github.com/goadesign/goa/goagen && \
    go install api/vendor/golang.org/x/lint/golint/ && \
    go install api/vendor/gopkg.in/godo.v2/cmd/godo

# build
RUN go build -o server

EXPOSE 8080
CMD ["./server"]
