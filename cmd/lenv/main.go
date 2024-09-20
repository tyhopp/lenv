package main

import (
	"flag"
	"log"
	"os"

	"github.com/tyhopp/lenv"
)

// getPaths retrieves the source path and destination paths based on the given environment.
func getPaths(logger *log.Logger, env string) (string, []string) {
	source, err := lenv.GetEnvFilePath(env)
	if err != nil {
		logger.Fatal(err)
	}

	destinations, err := lenv.ReadLenvFile()
	if err != nil {
		logger.Fatal(err)
	}

	return source, destinations
}

func main() {
	logger := log.New(os.Stderr, "", 0)

	envPtr := flag.String("env", ".env", "name of the environment file")
	flag.Parse()

	if len(flag.Args()) < 1 {
		logger.Fatal("lenv: missing subcommand, available: check, link, unlink")
	}

	subcommand := flag.Args()[0]

	switch subcommand {
	case "check":
		source, destinations := getPaths(logger, *envPtr)
		err := lenv.Check(source, destinations)
		if err != nil {
			logger.Fatal(err)
		}
	case "link":
		source, destinations := getPaths(logger, *envPtr)
		err := lenv.Link(source, destinations)
		if err != nil {
			logger.Fatal(err)
		}
	case "unlink":
		_, destinations := getPaths(logger, *envPtr)
		err := lenv.Unlink(destinations)
		if err != nil {
			logger.Fatal(err)
		}
	default:
		logger.Fatal("lenv: unknown subcommand, available: check, link, unlink")
	}
}
