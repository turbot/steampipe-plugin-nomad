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
		Description: "Retrieve information about your nodes.",
		List: &plugin.ListConfig{
			Hydrate: listNodes,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getNode,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The id of the node.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the node.",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "The status of the node.",
			},
			{
				Name:        "address",
				Type:        proto.ColumnType_STRING,
				Description: "The address of the node.",
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
			{
				Name:        "http_address",
				Type:        proto.ColumnType_STRING,
				Description: "HTTP address for the node",
				Hydrate:     getNode,
			},

			{
				Name:        "tls_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether TLS is enabled for the node's HTTP address",
				Hydrate:     getNode,
			},

			{
				Name:        "resources",
				Type:        proto.ColumnType_JSON,
				Description: "Resources allocated to the node",
				Hydrate:     getNode,
			},

			{
				Name:        "links",
				Type:        proto.ColumnType_JSON,
				Description: "Links to other nodes or entities",
				Hydrate:     getNode,
			},

			{
				Name:        "meta",
				Type:        proto.ColumnType_JSON,
				Description: "Metadata associated with the node",
				Hydrate:     getNode,
			},

			{
				Name:        "cgroup_parent",
				Type:        proto.ColumnType_STRING,
				Description: "The parent cgroup for the node",
				Hydrate:     getNode,
			},

			{
				Name:        "drain_strategy",
				Type:        proto.ColumnType_JSON,
				Description: "The strategy used to drain the node",
				Hydrate:     getNode,
			},

			{
				Name:        "status_updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The timestamp of the last status update for the node",
				Hydrate:     getNode,
				Transform:   transform.FromField("StatusUpdatedAt").Transform(transform.UnixToTimestamp),
			},

			{
				Name:        "events",
				Type:        proto.ColumnType_JSON,
				Description: "Events associated with the node",
				Hydrate:     getNode,
			},

			{
				Name:        "host_volumes",
				Type:        proto.ColumnType_JSON,
				Description: "Volumes attached to the node",
				Hydrate:     getNode,
			},

			{
				Name:        "host_networks",
				Type:        proto.ColumnType_JSON,
				Description: "Networks attached to the node",
				Hydrate:     getNode,
			},

			{
				Name:        "csi_controller_plugins",
				Type:        proto.ColumnType_JSON,
				Description: "CSI controller plugins attached to the node",
				Hydrate:     getNode,
			},

			{
				Name:        "csi_node_plugins",
				Type:        proto.ColumnType_JSON,
				Description: "CSI node plugins attached to the node",
				Hydrate:     getNode,
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "The title of the node.",
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
		if *d.QueryContext.Limit < maxLimit {
			maxLimit = *d.QueryContext.Limit
		}
	}

	nodeClient := client.Nodes()
	input := &api.QueryOptions{
		PerPage: int32(maxLimit),
	}

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
	var id string
	if h.Item != nil {
		id = h.Item.(*api.NodeListStub).ID
	} else {
		id = d.EqualsQualString("id")
	}

	// Create client
	client, err := getClient(ctx, d)
	nodeClient := client.Nodes()
	if err != nil {
		logger.Error("nomad_node.getNode", "connection_error", err)
		return nil, err
	}

	node, _, err := nodeClient.Info(id, &api.QueryOptions{})
	if err != nil {
		logger.Error("nomad_node.getNode", "api_error", err)
		return nil, err
	}

	return node, nil
}
