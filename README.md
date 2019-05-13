# a2d

Application to Docker (A2D) is a tool that is able to dockerize a single application.

## usage

```
docker run utrecht/a2d:1.0.0 projectName
```

returns

```
FROM golang:1.12.4-alpine as builder
COPY . ./projectName/
WORKDIR projectName
RUN adduser -D -g '' projectName && \
    apk add git && \
    CGO_ENABLED=0 go build && \
    cp projectName /projectName && \
    chmod 100 /projectName

FROM scratch
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder --chown=projectName:projectName /projectName /usr/local/projectName
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
USER projectName
ENTRYPOINT ["/usr/local/projectName"]
```

or if one does not have docker installed:

```
./a2d projectName
```
