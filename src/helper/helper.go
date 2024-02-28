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
		log.Printf("Could not load the %s \n Exeption: %v. Skipping... \n", file, err)
		f.Close()
		return
	}
	defer f.Close()

	content, err := io.ReadAll(f)

	if err != nil {
		log.Fatalf("Something went wrong while reading the %s \n Exception: %v", file, err)
	}

	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		env := strings.Split(line, "=")
		os.Setenv(env[0], env[1])
	}
}
