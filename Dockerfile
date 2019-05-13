FROM golang:1.12.4-alpine as builder
COPY . ./a2d/
WORKDIR a2d
RUN adduser -D -g '' a2d && \
    apk add git && \
    CGO_ENABLED=0 go build && \
    cp a2d /a2d && \
    chmod 100 /a2d

FROM scratch
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder --chown=a2d:a2d /a2d /usr/local/a2d
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
USER a2d
ENTRYPOINT ["/usr/local/a2d"]
