package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func readFile(f string) ([]byte, error) {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func TestDockerfileContent(t *testing.T) {
	dockerfileStruct{"ola"}.dockerfile()
	dockerfileContent()

	actual, err := readFile("Dockerfile")
	if err != nil {
		t.Errorf("Cannot read actual file: '%v'", err)
	}

	expected, err := readFile("tests/expectedDockerfile")
	if err != nil {
		t.Errorf("Cannot read expected file: '%v'", err)
	}

	if !bytes.Equal(expected, actual) {
		t.Errorf("Expected: '%s'. Actual: '%s'", expected, actual)
	}
}
