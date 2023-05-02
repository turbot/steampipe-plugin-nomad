package nomad

import (
	"context"

	"github.com/hashicorp/nomad/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableNomadACLPolicy(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "nomad_acl_policy",
		Description: "Retrieve information about your ACL policies.",
		List: &plugin.ListConfig{
			Hydrate: listACLPolicies,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    getACLPolicy,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the acl policy.",
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "The description of the acl policy.",
			},
			{
				Name:        "rules",
				Type:        proto.ColumnType_STRING,
				Description: "The set of rules of the acl policy.",
				Hydrate:     getACLPolicy,
			},
			{
				Name:        "create_index",
				Type:        proto.ColumnType_INT,
				Description: "The index when the acl policy was created.",
			},
			{
				Name:        "modify_index",
				Type:        proto.ColumnType_INT,
				Description: "The index when the acl policy was last modified.",
			},
			{
				Name:        "job_acl",
				Type:        proto.ColumnType_JSON,
				Description: "The capabilities of the acl policy.",
				Transform:   transform.FromField("JobACL"),
				Hydrate:     getACLPolicy,
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "The title of the acl policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

func listACLPolicies(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("nomad_acl_policy.listACLPolicies", "connection_error", err)
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
		policies, metadata, err := client.ACLPolicies().List(input)
		if err != nil {
			plugin.Logger(ctx).Error("nomad_acl_policy.listACLPolicies", "api_error", err)
			return nil, err
		}

		for _, policy := range policies {
			d.StreamListItem(ctx, policy)

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

func getACLPolicy(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	var name string
	if h.Item != nil {
		name = h.Item.(*api.ACLPolicyListStub).Name
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
		logger.Error("nomad_acl_policy.getACLPolicy", "connection_error", err)
		return nil, err
	}

	policy, _, err := client.ACLPolicies().Info(name, &api.QueryOptions{})
	if err != nil {
		logger.Error("nomad_acl_policy.getACLPolicy", "api_error", err)
		return nil, err
	}

	return policy, nil
}
