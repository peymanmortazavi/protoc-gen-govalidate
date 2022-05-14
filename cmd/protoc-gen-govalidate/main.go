package main

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"
)

var (
	singleFunc       = flag.Bool("singleFunc", true, "whether or not to include all models in a single function.")
	exportFunc       = flag.Bool("exportFunc", false, "whether or not to export a function to register tags.")
	useInit          = flag.Bool("useInit", true, "create init function to register all tags.")
	registerFunction = flag.String("registerFunc", "", "register function in form of <import-path>:<function-name>")
)

func main() {
	opts := protogen.Options{
		ParamFunc: flag.Set,
	}

	opts.Run(process)
}

func process(p *protogen.Plugin) error {
	return nil
}
