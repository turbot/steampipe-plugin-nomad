package nomad

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableNomadAgent(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "nomad_agent",
		Description: "Retrieve information about your agents.",
		List: &plugin.ListConfig{
			Hydrate: listAgents,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The id of the agent.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the agent.",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "The status of the agent.",
			},
			{
				Name:        "address",
				Type:        proto.ColumnType_STRING,
				Description: "The address of the agent.",
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
				Name:        "agent_class",
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
				Name:        "agent_resources",
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
				Description: "HTTP address for the agent",
			},

			{
				Name:        "tls_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether TLS is enabled for the agent's HTTP address",
			},

			{
				Name:        "resources",
				Type:        proto.ColumnType_JSON,
				Description: "Resources allocated to the agent",
			},

			{
				Name:        "links",
				Type:        proto.ColumnType_JSON,
				Description: "Links to other agents or entities",
			},

			{
				Name:        "meta",
				Type:        proto.ColumnType_JSON,
				Description: "Metadata associated with the agent",
			},

			{
				Name:        "cgroup_parent",
				Type:        proto.ColumnType_STRING,
				Description: "The parent cgroup for the agent",
			},

			{
				Name:        "drain_strategy",
				Type:        proto.ColumnType_JSON,
				Description: "The strategy used to drain the agent",
			},

			{
				Name:        "status_updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The timestamp of the last status update for the agent",

				Transform: transform.FromField("StatusUpdatedAt").Transform(transform.UnixToTimestamp),
			},

			{
				Name:        "events",
				Type:        proto.ColumnType_JSON,
				Description: "Events associated with the agent",
			},

			{
				Name:        "host_volumes",
				Type:        proto.ColumnType_JSON,
				Description: "Volumes attached to the agent",
			},

			{
				Name:        "host_networks",
				Type:        proto.ColumnType_JSON,
				Description: "Networks attached to the agent",
			},

			{
				Name:        "csi_controller_plugins",
				Type:        proto.ColumnType_JSON,
				Description: "CSI controller plugins attached to the agent",
			},

			{
				Name:        "csi_agent_plugins",
				Type:        proto.ColumnType_JSON,
				Description: "CSI agent plugins attached to the agent",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "The title of the agent.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

func listAgents(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("nomad_agent.listAgents", "connection_error", err)
		return nil, err
	}

	maxLimit := int64(1000)
	if d.QueryContext.Limit != nil {
		if *d.QueryContext.Limit < maxLimit {
			maxLimit = *d.QueryContext.Limit
		}
	}

	servers, err := client.Agent().Members()
	if err != nil {
		plugin.Logger(ctx).Error("nomad_agent.listAgents", "query_error", err)
		return nil, err
	}

	for _, server := range servers.Members {
		d.StreamListItem(ctx, server)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
