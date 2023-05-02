package nomad

import (
	"context"

	"github.com/hashicorp/nomad/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableNomadACLRole(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "nomad_acl_role",
		Description: "Retrieve information about your ACL roles.",
		List: &plugin.ListConfig{
			Hydrate: listACLRoles,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getACLRole,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the ACL role.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the ACL role.",
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "A human-readable, operator set description that can provide additional context about the ACL role.",
			},
			{
				Name:        "create_index",
				Type:        proto.ColumnType_INT,
				Description: "The index when the ACL role was created.",
			},
			{
				Name:        "modify_index",
				Type:        proto.ColumnType_INT,
				Description: "The index when the ACL role was last modified.",
			},
			{
				Name:        "policies",
				Type:        proto.ColumnType_JSON,
				Description: "An array of ACL policy links.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "The title of the acl role.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

func listACLRoles(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("nomad_acl_role.listACLRoles", "connection_error", err)
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
		roles, metadata, err := client.ACLRoles().List(input)
		if err != nil {
			plugin.Logger(ctx).Error("nomad_acl_role.listACLRoles", "api_error", err)
			return nil, err
		}

		for _, role := range roles {
			d.StreamListItem(ctx, role)

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

func getACLRole(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	var id string
	if h.Item != nil {
		id = h.Item.(*api.ACLRoleListStub).ID
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
		logger.Error("nomad_acl_role.getACLRole", "connection_error", err)
		return nil, err
	}

	role, _, err := client.ACLRoles().Get(id, &api.QueryOptions{})
	if err != nil {
		logger.Error("nomad_acl_role.getACLRole", "api_error", err)
		return nil, err
	}

	return role, nil
}
