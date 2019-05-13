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
	fileName    string
}

func (ds dockerfileStruct) dockerfile() error {
	f, err := os.Create(ds.fileName)
	if err != nil {
		return err
	}

	t, err := template.New(ds.fileName).Parse(`FROM golang:1.12.4-alpine as builder
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

func (ds dockerfileStruct) dockerfileContent() ([]byte, error) {
	data, err := ioutil.ReadFile(ds.fileName)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func validation() (string, string, error) {
	pn := flag.String("projectName", "", "The name of the app that has to be dockerized")
	f := flag.String("dockerfile", "Dockerfile", "The name of the Dockerfile that has to be created")
	d := flag.Bool("debug", false, "Whether debug level should be enabled")

	flag.Parse()

	if *pn == "" {
		return "", "", errors.New("projectName should not be empty")
	}

	if *d {
		log.SetLevel(log.DebugLevel)
	}

	return *pn, *f, nil
}

func main() {
	pn, f, err := validation()
	if err != nil {
		log.Fatal(err)
	}

	ds := dockerfileStruct{pn, f}
	ds.dockerfile()

	b, err := ds.dockerfileContent()
	if err != nil {
		log.Fatal(err)
	}
	log.Debug(string(b))
}
