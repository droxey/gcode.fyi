FROM golang:1.4 AS build

# -> Set environment variables:
ENV GOOS=linux
ENV GOARCH=amd64
ENV BINARY_NAME="backend"
ENV BACKEND_SRC="${GOPATH}/src/gcode.fyi/${BINARY_NAME}"
ENV BACKEND_BIN="${GOPATH}/bin/gcode.fyi"


# -> First add modules list to better utilize caching:
RUN mkdir -p ${BACKEND_SRC}
COPY go.sum go.mod ${BACKEND_SRC}
WORKDIR ${BACKEND_SRC}

# -> Download module dependencies:
RUN go mod download

# -> Copy code:
COPY . ${BACKEND_SRC}

# -> Build components:
RUN go install -installsuffix cgo -ldflags="-w -s"
RUN mkdir ${BACKEND_BIN}
RUN cp -r ${BACKEND_BIN}/backend ${BACKEND_SRC}

# -> Runtime container:
FROM alpine
LABEL maintainer="Dani Roxberry <dani@bitoriented.com>"

# -> Set environment variables:
ENV TZ="America/Los_Angeles"
ENV PATH="${BACKEND_BIN}:${PATH}"

# -> Install dependencies:
RUN apk add --update --no-cache sqlite tzdata ca-certificates bash

# -> Update time zone:
RUN cp --remove-destination /usr/share/zoneinfo/${TZ} /etc/localtime

# -> Fix "binary not found in path" bug.
# => See http://stackoverflow.com/questions/34729748/installed-go-binary-not-found-in-path-on-alpine-linux-docker
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

# -> Run backend executable:
WORKDIR ${BACKEND_SRC}
COPY --from=build . ${BACKEND_SRC}
EXPOSE 8585
CMD ["./backend"]
