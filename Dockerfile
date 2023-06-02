ARG GO_VERSION=1.20

FROM golang:${GO_VERSION}-alpine AS builder
ARG GOOS=linux
ARG GOARCH=amd64
ENV GO111MODULE on
RUN mkdir -p /app \
    && apk add --no-cache ca-certificates git gcc libc-dev curl make
WORKDIR /app
COPY . /app/
RUN go mod download \
    && make build-${GOOS}-${GOARCH}

FROM golang:${GO_VERSION}-alpine
LABEL maintainer="pooyan.info"
LABEL description="OSS License Audit Tool"
ARG BIN_NAME=olaudit
ENV GO111MODULE on

RUN apk add --no-cache ca-certificates git gcc libc-dev openssh bash
COPY --from=builder /app/build/${BIN_NAME} /bin/${BIN_NAME}
COPY scripts/entrypoint.sh /bin/entrypoint.sh
ENTRYPOINT ["/bin/entrypoint.sh"]