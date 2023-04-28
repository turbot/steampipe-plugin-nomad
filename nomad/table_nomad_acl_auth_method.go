package nomad

import (
	"context"

	"github.com/hashicorp/nomad/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableNomadACLAuthMethod(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "nomad_acl_auth_method",
		Description: "Retrieve information about your ACL auth methods.",
		List: &plugin.ListConfig{
			Hydrate: listACLAuthMethods,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    getACLAuthMethod,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the acl_auth_method.",
			},
			{
				Name:        "type",
				Type:        proto.ColumnType_STRING,
				Description: "Type is the SSO identifier this auth-method is.",
			},
			{
				Name:        "token_locality",
				Type:        proto.ColumnType_STRING,
				Description: "Defines whether the auth-method creates a local or global token when performing SSO login.",
				Hydrate:     getACLAuthMethod,
			},
			{
				Name:        "max_token_ttl",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The maximum life of a token created by this method.",
				Transform:   transform.FromField("MaxTokenTTL").Transform(transform.UnixToTimestamp),
				Hydrate:     getACLAuthMethod,
			},
			{
				Name:        "default",
				Type:        proto.ColumnType_BOOL,
				Description: "Default identifies whether this is the default auth-method to use when attempting to login without specifying an auth-method name to use.",
			},
			{
				Name:        "config",
				Type:        proto.ColumnType_JSON,
				Description: "Config contains the detailed configuration which is specific to the auth-method.",
				Hydrate:     getACLAuthMethod,
			},
			{
				Name:        "create_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The creation time of the auth method.",
				Hydrate:     getACLAuthMethod,
			},
			{
				Name:        "modify_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The last modification time of the auth method.",
				Hydrate:     getACLAuthMethod,
			},
			{
				Name:        "create_index",
				Type:        proto.ColumnType_INT,
				Description: "Create index of the auth method.",
			},
			{
				Name:        "modify_index",
				Type:        proto.ColumnType_INT,
				Description: "Modify index of the auth method.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "The title of the acl auth method.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

func listACLAuthMethods(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("nomad_acl_auth_method.listACLAuthMethods", "connection_error", err)
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
		authMethods, metadata, err := client.ACLAuthMethods().List(input)
		if err != nil {
			plugin.Logger(ctx).Error("nomad_acl_auth_method.listACLAuthMethods", "query_error", err)
			return nil, err
		}

		for _, authMethod := range authMethods {
			d.StreamListItem(ctx, authMethod)

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

func getACLAuthMethod(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	var name string
	if h.Item != nil {
		name = h.Item.(*api.ACLAuthMethodListStub).Name
	} else {
		name = d.EqualsQualString("name")
	}

	// check if name is empty
	if name == "" {
		return nil, nil
	}

	// Create client
	client, err := getClient(ctx, d)
	if err != nil {
		logger.Error("nomad_acl_auth_method.getACLAuthMethod", "connection_error", err)
		return nil, err
	}

	authMethod, _, err := client.ACLAuthMethods().Get(name, &api.QueryOptions{})
	if err != nil {
		logger.Error("nomad_acl_auth_method.getACLAuthMethod", "api_error", err)
		return nil, err
	}

	return authMethod, nil
}
