ARG GO_VERSION=1.23-alpine

# Builder image
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

# Final production image
FROM golang:${GO_VERSION}-alpine
LABEL maintainer="pooyan.info"
LABEL description="OSS License Audit Tool"
ARG BIN_NAME=olaudit
ENV GO111MODULE on

# Create a non-root user with UID and GID 1000 and home directory /app
RUN addgroup -g 1000 appgroup \
    && adduser -D -u 1000 -G appgroup -h /app -s /sbin/nologin appuser

RUN apk add --no-cache ca-certificates git gcc libc-dev bash openssl>=3.1.1-r0
COPY --from=builder /app/build/${BIN_NAME} /bin/${BIN_NAME}
COPY scripts/entrypoint.sh /bin/entrypoint.sh

# Grant read and execute permissions to appuser for /bin/${BIN_NAME} and /bin/entrypoint.sh
RUN chown appuser:appgroup /bin/${BIN_NAME} /bin/entrypoint.sh \
    && chmod 550 /bin/${BIN_NAME} /bin/entrypoint.sh

USER appuser

ENTRYPOINT ["/bin/entrypoint.sh"]
