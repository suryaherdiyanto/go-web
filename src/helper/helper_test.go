package helper

import (
	"testing"
)

func TestLoadEnv(t *testing.T) {
	LoadEnv("./../../.env")
}

func TestIfEnvFileNotFound(t *testing.T) {
	LoadEnv("./.env.prod")
}
