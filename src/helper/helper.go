package helper

import (
	"io"
	"log"
	"os"
	"strings"
)

func LoadEnv(file string) {

	f, err := os.Open(file)

	if err != nil {
		log.Fatalf("Could not load the %s file: %v", file, err)
	}
	defer f.Close()

	content, err := io.ReadAll(f)

	if err != nil {
		log.Fatalf("Something went wrong while reading the %s file %v", file, err)
	}

	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		env := strings.Split(line, "=")
		os.Setenv(env[0], env[1])
	}
}
