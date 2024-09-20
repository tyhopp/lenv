package main

import (
	"log"
	"os"

	"github.com/tyhopp/lenv"
)

func main() {
	logger := log.New(os.Stdout, "", 0)

	if len(os.Args) < 2 {
		logger.Fatal("lenv: missing subcommand, available: check, link, unlink")
	}

	getPaths := func() (string, []string) {
		source, err := lenv.GetEnvFilePath()
		if err != nil {
			logger.Fatal(err)
		}

		destinations, err := lenv.ReadLenvFile()
		if err != nil {
			logger.Fatal(err)
		}

		return source, destinations
	}

	switch os.Args[1] {
	case "check":
		source, destinations := getPaths()
		err := lenv.Check(source, destinations)
		if err != nil {
			logger.Fatal(err)
		}
	case "link":
		source, destinations := getPaths()
		err := lenv.Link(source, destinations)
		if err != nil {
			logger.Fatal(err)
		}
	case "unlink":
		_, destinations := getPaths()
		err := lenv.Unlink(destinations)
		if err != nil {
			logger.Fatal(err)
		}
	default:
		logger.Fatal("lenv: unknown subcommand, available: check, link, unlink")
	}
}
