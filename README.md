# a2d

[![Build Status](https://travis-ci.org/030/a2d.svg?branch=master)](https://travis-ci.org/030/a2d)
[![Go Report Card](https://goreportcard.com/badge/github.com/030/a2d)](https://goreportcard.com/report/github.com/030/a2d)
![DevOps SE Questions](https://img.shields.io/stackexchange/devops/t/a2d.svg)
![Docker Pulls](https://img.shields.io/docker/pulls/utrecht/a2d.svg)
![Issues](https://img.shields.io/github/issues-raw/030/a2d.svg)
![Pull requests](https://img.shields.io/github/issues-pr-raw/030/a2d.svg)
![Total downloads](https://img.shields.io/github/downloads/030/a2d/total.svg)
![License](https://img.shields.io/github/license/030/a2d.svg)
![Repository Size](https://img.shields.io/github/repo-size/030/a2d.svg)
![Contributors](https://img.shields.io/github/contributors/030/a2d.svg)
![Commit activity](https://img.shields.io/github/commit-activity/m/030/a2d.svg)
![Last commit](https://img.shields.io/github/last-commit/030/a2d.svg)
![Release date](https://img.shields.io/github/release-date/030/a2d.svg)
![Latest Production Release Version](https://img.shields.io/github/release/030/a2d.svg)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=030_a2d&metric=bugs)](https://sonarcloud.io/dashboard?id=030_a2d)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=030_a2d&metric=code_smells)](https://sonarcloud.io/dashboard?id=030_a2d)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=030_a2d&metric=coverage)](https://sonarcloud.io/dashboard?id=030_a2d)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=030_a2d&metric=duplicated_lines_density)](https://sonarcloud.io/dashboard?id=030_a2d)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=030_a2d&metric=ncloc)](https://sonarcloud.io/dashboard?id=030_a2d)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=030_a2d&metric=sqale_rating)](https://sonarcloud.io/dashboard?id=030_a2d)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=030_a2d&metric=alert_status)](https://sonarcloud.io/dashboard?id=030_a2d)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=030_a2d&metric=reliability_rating)](https://sonarcloud.io/dashboard?id=030_a2d)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=030_a2d&metric=security_rating)](https://sonarcloud.io/dashboard?id=030_a2d)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=030_a2d&metric=sqale_index)](https://sonarcloud.io/dashboard?id=030_a2d)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=030_a2d&metric=vulnerabilities)](https://sonarcloud.io/dashboard?id=030_a2d)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/2812/badge)](https://bestpractices.coreinfrastructure.org/projects/2812)

Application to Docker (A2D) is a tool that is able to dockerize a single application.

## usage

```
docker run -v $PWD:/projectName utrecht/a2d:1.0.0 \
       -projectName someProjectName -dockerfile /projectName/Dockerfile
```

creates a Dockerfile. If one does not have docker installed, one could issue:

```
./a2d -h
```

and the following output would be returned:

```
FROM golang:1.12.4-alpine as builder
COPY . ./someProjectName/
WORKDIR someProjectName
RUN adduser -D -g '' someProjectName && \
    apk add git && \
    CGO_ENABLED=0 go build && \
    cp someProjectName /someProjectName && \
    chmod 100 /someProjectName

FROM scratch
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder --chown=someProjectName:someProjectName /someProjectName /usr/local/someProjectName
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
USER someProjectName
ENTRYPOINT ["/usr/local/someProjectName"]
```
