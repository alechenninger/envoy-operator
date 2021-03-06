package main

import (
	"flag"
	"log"

	"github.com/solo-io/envoy-operator/pkg/downward"
)

func main() {
	inputfile := flag.String("input", "", "input file")
	outfile := flag.String("output", "", "output file")
	flag.Parse()
	transformer := downward.NewTransformer()
	err := transformer.TransformFiles(*inputfile, *outfile)
	if err != nil {
		log.Fatalf("initializer failed: %v", err)
	}
}
