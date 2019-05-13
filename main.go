package main

import (
	"errors"
	"flag"
	"html/template"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
)

type dockerfileStruct struct {
	ProjectName string
}

func (ds dockerfileStruct) dockerfile() error {
	f, err := os.Create("Dockerfile")
	if err != nil {
		return err
	}

	t, err := template.New("Dockerfile").Parse(`FROM golang:1.12.4-alpine as builder
COPY . ./{{.ProjectName}}/
WORKDIR projectName
RUN adduser -D -g '' {{.ProjectName}} && \
    apk add git && \
    CGO_ENABLED=0 go build && \
    cp {{.ProjectName}} /{{.ProjectName}} && \
    chmod 100 /{{.ProjectName}}

FROM scratch
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder --chown={{.ProjectName}}:{{.ProjectName}} /{{.ProjectName}} /usr/local/{{.ProjectName}}
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
USER {{.ProjectName}}
ENTRYPOINT ["/usr/local/{{.ProjectName}}"]
`)
	if err != nil {
		return err
	}

	err = t.Execute(f, ds)
	if err != nil {
		return err
	}
	return nil
}

func dockerfileContent() error {
	data, err := ioutil.ReadFile("Dockerfile")
	if err != nil {
		return err
	}
	log.Debug(string(data))
	return nil
}

func validation() (string, error) {
	pn := flag.String("projectName", "", "The name of the app that has to be dockerized")
	d := flag.Bool("debug", false, "Whether debug level should be enabled")

	flag.Parse()

	if *pn == "" {
		return "", errors.New("projectName should not be empty")
	}

	if *d {
		log.SetLevel(log.DebugLevel)
	}

	return *pn, nil
}

func main() {
	pn, err := validation()
	if err != nil {
		log.Fatal(err)
	}

	dockerfileStruct{pn}.dockerfile()

	err = dockerfileContent()
	if err != nil {
		log.Fatal(err)
	}
}
