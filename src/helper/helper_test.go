package helper

import (
	"bytes"
	"errors"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	LoadEnv("./../../.env")
}

func TestIfEnvFileNotFound(t *testing.T) {
	LoadEnv("./.env.prod")
}

func TestHandleException(t *testing.T) {
	err := errors.New("Some Exception")
	w := new(bytes.Buffer)
	HandleExeption(w, err)

	if w.String() != "Something went wrong: Some Exception" {
		t.Errorf("Expected %q, got: %s", "Something went wrong: Some Exception", w.String())
	}
}
