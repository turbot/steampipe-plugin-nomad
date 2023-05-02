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
				Description: "A string representing the unique identifier of the node.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "A string representing the name of the node.",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "A string representing the current status of the node.",
			},
			{
				Name:        "create_index",
				Type:        proto.ColumnType_INT,
				Description: "An unsigned 64-bit integer representing the index at which the node was created.",
			},
			{
				Name:        "datacenter",
				Type:        proto.ColumnType_STRING,
				Description: "A string specifying the datacenter in which the node is located.",
			},
			{
				Name:        "drain",
				Type:        proto.ColumnType_BOOL,
				Description: "A boolean value indicating whether the node is currently being drained.",
			},
			{
				Name:        "modify_index",
				Type:        proto.ColumnType_INT,
				Description: "An unsigned 64-bit integer representing the index at which the node was last modified.",
			},
			{
				Name:        "node_class",
				Type:        proto.ColumnType_STRING,
				Description: "A string specifying the class of the node.",
			},
			{
				Name:        "scheduling_eligibility",
				Type:        proto.ColumnType_STRING,
				Description: "A string specifying the node's scheduling eligibility.",
			},
			{
				Name:        "status_description",
				Type:        proto.ColumnType_STRING,
				Description: "A string providing a description of the node's current status.",
			},
			{
				Name:        "http_address",
				Type:        proto.ColumnType_STRING,
				Description: "A string specifying the HTTP address of the node.",
				Transform:   transform.FromField("HTTPAddr"),
				Hydrate:     getNode,
			},
			{
				Name:        "tls_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "A boolean value indicating whether TLS is enabled for the node.",
				Transform:   transform.FromField("TLSEnabled"),
				Hydrate:     getNode,
			},
			{
				Name:        "cgroup_parent",
				Type:        proto.ColumnType_STRING,
				Description: "A string specifying the parent cgroup for the node.",
				Hydrate:     getNode,
			},
			{
				Name:        "status_updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "An integer representing the timestamp at which the node's status was last updated.",
				Hydrate:     getNode,
				Transform:   transform.FromField("StatusUpdatedAt").Transform(transform.UnixToTimestamp),
			},
			{
				Name:        "events",
				Type:        proto.ColumnType_JSON,
				Description: "Represents events associated with the node.",
				Hydrate:     getNode,
			},
			{
				Name:        "host_volumes",
				Type:        proto.ColumnType_JSON,
				Description: "A map containing information about the volumes attached to the node.",
				Hydrate:     getNode,
			},
			{
				Name:        "host_networks",
				Type:        proto.ColumnType_JSON,
				Description: "A map containing information about the networks attached to the node.",
				Hydrate:     getNode,
			},
			{
				Name:        "csi_controller_plugins",
				Type:        proto.ColumnType_JSON,
				Description: "A map containing information about the CSI controller plugins installed on the node.",
				Hydrate:     getNode,
			},
			{
				Name:        "csi_node_plugins",
				Type:        proto.ColumnType_JSON,
				Description: "A map containing information about the CSI node plugins installed on the node.",
				Hydrate:     getNode,
			},
			{
				Name:        "attributes",
				Type:        proto.ColumnType_JSON,
				Description: "A map containing user-defined attributes associated with the node.",
			},
			{
				Name:        "drivers",
				Type:        proto.ColumnType_JSON,
				Description: "A map containing information about the drivers installed on the node.",
			},
			{
				Name:        "last_drain",
				Type:        proto.ColumnType_JSON,
				Description: "Represents the metadata for the last drain operation performed on the node.",
			},
			{
				Name:        "node_resources",
				Type:        proto.ColumnType_JSON,
				Description: "Represents the resources allocated to the node.",
			},
			{
				Name:        "reserved_resources",
				Type:        proto.ColumnType_JSON,
				Description: "Represents the reserved resources on the node.",
			},
			{
				Name:        "resources",
				Type:        proto.ColumnType_JSON,
				Description: "Represents the total resources available on the node.",
				Hydrate:     getNode,
			},
			{
				Name:        "reserved",
				Type:        proto.ColumnType_JSON,
				Description: "Represents the reserved resources on the node.",
				Hydrate:     getNode,
			},
			{
				Name:        "links",
				Type:        proto.ColumnType_JSON,
				Description: "A map containing links to related resources associated with the node.",
				Hydrate:     getNode,
			},
			{
				Name:        "meta",
				Type:        proto.ColumnType_JSON,
				Description: "A map containing additional metadata associated with the node.",
				Hydrate:     getNode,
			},
			{
				Name:        "drain_strategy",
				Type:        proto.ColumnType_JSON,
				Description: "Represents the strategy used for draining the node.",
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

	input := &api.QueryOptions{
		PerPage: int32(maxLimit),
	}

	for {
		nodes, metadata, err := client.Nodes().List(input)
		if err != nil {
			plugin.Logger(ctx).Error("nomad_node.listNodes", "api_error", err)
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
	if err != nil {
		logger.Error("nomad_node.getNode", "connection_error", err)
		return nil, err
	}

	node, _, err := client.Nodes().Info(id, &api.QueryOptions{})
	if err != nil {
		logger.Error("nomad_node.getNode", "api_error", err)
		return nil, err
	}

	return node, nil
}
