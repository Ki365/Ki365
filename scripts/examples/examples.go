package main

import (
	"github.com/ki365/ki365/abatement"
)

// Generates initial example projects from examples/source to examples/build
func main() {
	abatement.GenerateExamples("./examples/source/", "./examples/build/", true)
}
