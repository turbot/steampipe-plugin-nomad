package nomad

import (
	"context"
	"errors"

	"github.com/hashicorp/nomad/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type nomadConfig struct {
	Address *string `cty:"address"`
	Region  *string `cty:"region"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"address": {
		Type: schema.TypeString,
	},
	"region": {
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

func getClient(ctx context.Context, d *plugin.QueryData) (*api.Client, error) {
	nomadConfig := GetConfig(d.Connection)

	// token := os.Getenv("OP_CONNECT_TOKEN")
	// url := os.Getenv("OP_CONNECT_HOST")
	address := ""
	region := ""
	if nomadConfig.Address != nil {
		address = *nomadConfig.Address
	}
	// if nomadConfig.Region != nil {
	// 	region = *nomadConfig.Region
	// }

	if address != "" {
		con := api.DefaultConfig()
		con.Address = address
		con.Region = region
		con.SecretID = "c178b810-8b18-6f38-016f-725ddec5d587"
		client, _ := api.NewClient(con)
		return client, nil
	}

	return nil, errors.New("'token' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe.")
}
