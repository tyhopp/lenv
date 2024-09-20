package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tyhopp/lenv"
)

// printUsage prints the usage instructions for the lenv command.
func printUsage() {
	fmt.Println("Usage: lenv [options] <subcommand>")
	fmt.Println("Options:")
	flag.PrintDefaults()
	fmt.Println("Subcommands:")
	fmt.Println("  check   - Check status of symlinks between source env file and destinations")
	fmt.Println("  link    - Symlink source env file to destinations")
	fmt.Println("  unlink  - Remove symlinks from destinations")
}

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
	helpPtr := flag.Bool("help", false, "display help information")
	flag.Parse()

	if *helpPtr {
		printUsage()
		return
	}

	if len(flag.Args()) < 1 {
		logger.Fatal("lenv: missing subcommand, use -help flag for usage instructions")
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
		logger.Fatal("lenv: unknown subcommand, use -help flag for usage instructions")
	}
}
