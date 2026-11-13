package main

import (
	"github.com/turbot/steampipe-plugin-nomad/nomad"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: nomad.Plugin})
}
