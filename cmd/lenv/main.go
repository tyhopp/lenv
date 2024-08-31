package main

import (
	"github.com/tyhopp/lenv"
)

func main() {
	logger := lenv.Logger()
	source := lenv.GetEnvFilePath(logger)
	destinations := lenv.ReadLenvFile(logger)
	lenv.Check(logger, source, destinations)
}
