package main

import (
	"github.com/steampipe-plugin-nomad/nomad"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: nomad.Plugin})
}
