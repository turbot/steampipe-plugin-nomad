package nomad

import (
	"context"
	"errors"
	"os"

	"github.com/hashicorp/nomad/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type nomadConfig struct {
	Address  *string `cty:"address"`
	SecretID *string `cty:"secret_id"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"address": {
		Type: schema.TypeString,
	},
	"secret_id": {
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

	namespace := "*"
	address := os.Getenv("NOMAD_ADDR")
	namespace = os.Getenv("NOMAD_NAMESPACE")
	secretId := os.Getenv("NOMAD_TOKEN")

	if nomadConfig.Address != nil {
		address = *nomadConfig.Address
	}
	if nomadConfig.SecretID != nil {
		secretId = *nomadConfig.SecretID
	}

	if address != "" {
		con := api.DefaultConfig()
		con.Address = address
		con.SecretID = secretId
		con.Namespace = namespace
		client, _ := api.NewClient(con)
		return client, nil
	}

	return nil, errors.New("'token' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe.")
}
