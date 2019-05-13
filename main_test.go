package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
)

const (
	testDockerFile = "someDockerfile"
)

// See https://stackoverflow.com/a/34102842/2777965
func TestMain(m *testing.M) {
	setup()
	m.Run()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

var ds = dockerfileStruct{"ola", testDockerFile}

func setup() {
	ds.dockerfile()
}

func readFile(f string) ([]byte, error) {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func TestDockerfileContent(t *testing.T) {
	actual, err := ds.dockerfileContent()
	if err != nil {
		t.Error(err)
	}

	dsExpected := dockerfileStruct{"", "tests/expectedDockerfile"}
	expected, err := dsExpected.dockerfileContent()
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(expected, actual) {
		t.Errorf("Expected: '%s'. Actual: '%s'", expected, actual)
	}
}

func shutdown() {
	err := os.Remove(testDockerFile)
	if err != nil {
		log.Fatal(err)
	}
}
