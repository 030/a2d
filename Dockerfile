FROM golang:1.12.4-alpine as builder
COPY . ./ola/
WORKDIR projectName
RUN adduser -D -g '' ola && \
    apk add git && \
    CGO_ENABLED=0 go build && \
    cp ola /ola && \
    chmod 100 /ola

FROM scratch
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder --chown=ola:ola /ola /usr/local/ola
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
USER ola
ENTRYPOINT ["/usr/local/ola"]
