package main

import (
	"fmt"

	"github.com/tyhopp/lenv"
)

func main() {
	paths := lenv.ReadLenvFile()

	for _, paths := range paths {
		fmt.Println(paths)
	}
}
