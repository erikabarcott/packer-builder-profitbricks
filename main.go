package main

import (
	"github.com/StackPointCloud/packer-builder-profitbricks/builder"
	"github.com/mitchellh/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()

	if err != nil {
		panic(err)
	}
	builder := profitbricks.Builder{}
	server.RegisterBuilder(&builder)
	server.Serve()
}
