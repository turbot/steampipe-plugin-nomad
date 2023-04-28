package nomad

import (
	"context"

	"github.com/hashicorp/nomad/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableNomadPlugin(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "nomad_plugin",
		Description: "Retrieve information about your plugins.",
		List: &plugin.ListConfig{
			Hydrate: listPlugins,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getPlugin,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the plugin.",
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "The description of the plugin.",
			},
			{
				Name:        "quota",
				Type:        proto.ColumnType_STRING,
				Description: "The quota of the plugin.",
			},
			{
				Name:        "capabilities",
				Type:        proto.ColumnType_JSON,
				Description: "The capabilities of the plugin.",
			},
			{
				Name:        "meta",
				Type:        proto.ColumnType_JSON,
				Description: "The metadata associated with the plugin.",
			},
			{
				Name:        "create_index",
				Type:        proto.ColumnType_INT,
				Description: "The index when the plugin was created.",
			},
			{
				Name:        "modify_index",
				Type:        proto.ColumnType_INT,
				Description: "The index when the plugin was last modified.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "The title of the plugin.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

func listPlugins(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("nomad_plugin.listPlugins", "connection_error", err)
		return nil, err
	}

	maxLimit := int64(1000)
	if d.QueryContext.Limit != nil {
		if *d.QueryContext.Limit < maxLimit {
			maxLimit = *d.QueryContext.Limit
		}
	}

	input := &api.QueryOptions{
		PerPage: int32(maxLimit),
	}

	// Add support for optional region qual
	if d.EqualsQuals["datacenter"] != nil {
		input.Region = d.EqualsQuals["datacenter"].GetStringValue()
	}

	for {
		plugins, metadata, err := client.CSIPlugins().List(input)
		if err != nil {
			plugin.Logger(ctx).Error("nomad_plugin.listPlugins", "query_error", err)
			return nil, err
		}

		for _, plugin := range plugins {
			d.StreamListItem(ctx, plugin)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		input.NextToken = metadata.NextToken
		if input.NextToken == "" {
			break
		}
	}

	return nil, nil
}

func getPlugin(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	var id string
	if h.Item != nil {
		id = h.Item.(*api.CSIPluginListStub).ID
	} else {
		id = d.EqualsQualString("id")
	}

	// check if id is empty
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := getClient(ctx, d)
	if err != nil {
		logger.Error("nomad_node.getPlugin", "connection_error", err)
		return nil, err
	}

	plugin, _, err := client.CSIPlugins().Info(id, &api.QueryOptions{})
	if err != nil {
		logger.Error("nomad_node.getPlugin", "api_error", err)
		return nil, err
	}

	return plugin, nil
}
