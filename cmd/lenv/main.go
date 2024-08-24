package main

import (
	"github.com/tyhopp/lenv"
)

func main() {
	// todo: parse command line arguments and conditionally call these functions
	lenv.Check()
	lenv.Link()
	lenv.Unlink()
}
