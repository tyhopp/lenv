package main

import (
	"log"
	"os"

	"github.com/tyhopp/lenv"
)

func main() {
	logger := log.New(os.Stdout, "", 0)

	source, err := lenv.GetEnvFilePath()
	if err != nil {
		logger.Fatal(err)
	}

	destinations, err := lenv.ReadLenvFile()
	if err != nil {
		logger.Fatal(err)
	}

	err = lenv.Check(source, destinations)
	if err != nil {
		logger.Fatal(err)
	}
}
