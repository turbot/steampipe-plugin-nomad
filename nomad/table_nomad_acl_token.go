package nomad

import (
	"context"

	"github.com/hashicorp/nomad/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableNomadACLToken(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "nomad_acl_token",
		Description: "Retrieve information about your ACL tokens.",
		List: &plugin.ListConfig{
			Hydrate: listACLTokens,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("accessor_id"),
			Hydrate:    getACLToken,
		},
		Columns: []*plugin.Column{
			{
				Name:        "accessor_id",
				Type:        proto.ColumnType_STRING,
				Description: "The accessor ID of the acl token.",
				Transform:   transform.FromField("AccessorID"),
			},
			{
				Name:        "secret_id",
				Type:        proto.ColumnType_STRING,
				Description: "The secret ID of the acl token.",
				Transform:   transform.FromField("SecretID"),
				Hydrate:     getACLToken,
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the acl token.",
			},
			{
				Name:        "type",
				Type:        proto.ColumnType_STRING,
				Description: "The type of the acl token.",
			},
			{
				Name:        "global",
				Type:        proto.ColumnType_BOOL,
				Description: "Check whether the token is global or not.",
			},
			{
				Name:        "create_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The creation time of the acl token.",
			},
			{
				Name:        "expiration_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The expiration time of the acl token.",
			},
			{
				Name:        "expiration_ttl",
				Type:        proto.ColumnType_STRING,
				Description: "The maximum life of the acl token.",
				Transform:   transform.FromField("ExpirationTTL"),
				Hydrate:     getACLToken,
			},
			{
				Name:        "create_index",
				Type:        proto.ColumnType_INT,
				Description: "Create index of the acl token.",
			},
			{
				Name:        "modify_index",
				Type:        proto.ColumnType_INT,
				Description: "Modify index of the acl token.",
			},
			{
				Name:        "policies",
				Type:        proto.ColumnType_JSON,
				Description: "Policies attached to the acl token.",
			},
			{
				Name:        "roles",
				Type:        proto.ColumnType_JSON,
				Description: "Roles attached to the acl token.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "The title of the acl token.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

func listACLTokens(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("nomad_acl_token.listACLTokens", "connection_error", err)
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
		tokens, metadata, err := client.ACLTokens().List(input)
		if err != nil {
			plugin.Logger(ctx).Error("nomad_acl_token.listACLTokens", "api_error", err)
			return nil, err
		}

		for _, token := range tokens {
			d.StreamListItem(ctx, token)

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

func getACLToken(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	var accessorID string
	if h.Item != nil {
		accessorID = h.Item.(*api.ACLTokenListStub).AccessorID
	} else {
		accessorID = d.EqualsQualString("accessor_id")
	}

	// check if accessorID is empty
	if accessorID == "" {
		return nil, nil
	}

	// Create client
	client, err := getClient(ctx, d)
	if err != nil {
		logger.Error("nomad_acl_token.getACLToken", "connection_error", err)
		return nil, err
	}

	token, _, err := client.ACLTokens().Info(accessorID, &api.QueryOptions{})
	if err != nil {
		logger.Error("nomad_acl_token.getACLToken", "api_error", err)
		return nil, err
	}

	return token, nil
}
