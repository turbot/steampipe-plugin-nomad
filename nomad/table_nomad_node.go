package nomad

import (
	"context"

	"github.com/hashicorp/nomad/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableNomadNode(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "nomad_node",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate: listNodes,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate: getNode,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The UUID of the vault.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the vault.",
			},
			{
				Name:        "address",
				Type:        proto.ColumnType_STRING,
				Description: "The version of the vault metadata.",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "The version of the vault metadata.",
			},
			{
				Name:        "create_index",
				Type:        proto.ColumnType_INT,
				Description: "The version of the vault metadata.",
			},
			{
				Name:        "datacenter",
				Type:        proto.ColumnType_STRING,
				Description: "The version of the vault metadata.",
			},
			{
				Name:        "drain",
				Type:        proto.ColumnType_BOOL,
				Description: "The version of the vault metadata.",
			},
			{
				Name:        "modify_index",
				Type:        proto.ColumnType_INT,
				Description: "The version of the vault metadata.",
			},
			{
				Name:        "node_class",
				Type:        proto.ColumnType_STRING,
				Description: "Date and time when the vault was created.",
			},
			{
				Name:        "scheduling_eligibility",
				Type:        proto.ColumnType_STRING,
				Description: "The version of the vault contents.",
			},
			{
				Name:        "status_description",
				Type:        proto.ColumnType_STRING,
				Description: "The description for the vault.",
			},
			{
				Name:        "version",
				Type:        proto.ColumnType_STRING,
				Description: "Number of active items in the vault.",
			},
			{
				Name:        "attributes",
				Type:        proto.ColumnType_JSON,
				Description: "The type of vault. Possible values are EVERYONE, PERSONAL and USER_CREATED.",
			},
			{
				Name:        "drivers",
				Type:        proto.ColumnType_JSON,
				Description: "Date and time when the vault or its contents were last changed.",
			},
			{
				Name:        "last_drain",
				Type:        proto.ColumnType_JSON,
				Description: "Date and time when the vault or its contents were last changed.",
			},
			{
				Name:        "node_resources",
				Type:        proto.ColumnType_JSON,
				Description: "Date and time when the vault or its contents were last changed.",
			},
			{
				Name:        "reserved_resources",
				Type:        proto.ColumnType_JSON,
				Description: "Date and time when the vault or its contents were last changed.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "The title of the vault.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

func listNodes(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("nomad_node.listNodes", "connection_error", err)
		return nil, err
	}

		maxLimit := int64(1000)
		if d.QueryContext.Limit != nil {
			if (*d.QueryContext.Limit < maxLimit){
				maxLimit = *d.QueryContext.Limit
			}
		}

	nodeClient := client.Nodes()
	input := &api.QueryOptions{
		PerPage: int32(maxLimit),
	}

	// Add support for optional region qual
	if d.EqualsQuals["datacenter"] != nil {
		input.Region = d.EqualsQuals["datacenter"].GetStringValue()
	}

	for {
		nodes, metadata, err := nodeClient.List(input)
		if err != nil {
			plugin.Logger(ctx).Error("nomad_node.listNodes", "query_error", err)
			return nil, err
		}

		for _, node := range nodes {
			d.StreamListItem(ctx, node)

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

func getNode(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("id")
	// Create client
	input := &api.QueryOptions{}
	client, err := getClient(ctx, d)
	nodeClient := client.Nodes()
	if err != nil {
		logger.Error("nomad_node.getNode", "connection_error", err)
		return nil, err
	}

	node, _, err := nodeClient.Info(id, input)
	if err != nil {
		logger.Error("nomad_node.getNode", "api_error", err)
		return nil, err
	}

	return node, nil
}
