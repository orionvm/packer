package main

import (
	"github.com/mitchellh/packer/builder/xen"
	"github.com/mitchellh/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(xen.Builder))
	server.Serve()
}
