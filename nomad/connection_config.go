package nomad

import(
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type nomadConfig struct {
	address *string `cty:address`
}

var ConfigSchema = map[string]*schema.Attribute{
	"address": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &nomadConfig{}
}

func GetConfig(connection *plugin.Connection) nomadConfig {
	if connection == nil || connection.Config == nil {
		return nomadConfig{}
	}
	config, _ := connection.Config.(nomadConfig)
	return config
}
