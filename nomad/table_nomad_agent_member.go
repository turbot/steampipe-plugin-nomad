package nomad

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableNomadAgentMember(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "nomad_agent_member",
		Description: "Retrieve information about your agent members.",
		List: &plugin.ListConfig{
			Hydrate: listAgentMembers,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the agent member.",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "The current operational status of the agent member.",
			},
			{
				Name:        "address",
				Type:        proto.ColumnType_STRING,
				Description: "The IP address or hostname of the agent member.",
			},
			{
				Name:        "port",
				Type:        proto.ColumnType_INT,
				Description: "The network port number used by the agent member.",
			},
			{
				Name:        "protocol_min",
				Type:        proto.ColumnType_INT,
				Description: "The minimum version of the protocol supported by the agent member.",
			},
			{
				Name:        "protocol_max",
				Type:        proto.ColumnType_INT,
				Description: "The maximum version of the protocol supported by the agent member.",
			},
			{
				Name:        "protocol_cur",
				Type:        proto.ColumnType_INT,
				Description: "The current version of the protocol used by the agent member.",
			},
			{
				Name:        "delegate_min",
				Type:        proto.ColumnType_INT,
				Description: " The minimum number of delegations allowed by the agent member.",
			},
			{
				Name:        "delegate_max",
				Type:        proto.ColumnType_INT,
				Description: "The maximum number of delegations allowed by the agent member.",
			},
			{
				Name:        "delegate_cur",
				Type:        proto.ColumnType_INT,
				Description: "The current number of delegations held by the agent member.",
			},
			{
				Name:        "tags",
				Type:        proto.ColumnType_JSON,
				Description: "A set of key-value pairs that describe the agent member.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "The title of the agent member.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

func listAgentMembers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("nomad_agent_member.listAgentMembers", "connection_error", err)
		return nil, err
	}

	servers, err := client.Agent().Members()
	if err != nil {
		plugin.Logger(ctx).Error("nomad_agent_member.listAgentMembers", "api_error", err)
		return nil, err
	}

	for _, member := range servers.Members {
		d.StreamListItem(ctx, member)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
