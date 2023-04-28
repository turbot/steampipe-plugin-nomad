package nomad

import (
	"context"

	"github.com/hashicorp/nomad/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableNomadACLBindingRule(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "nomad_acl_binding_rule",
		Description: "Retrieve information about your ACL binding rules.",
		List: &plugin.ListConfig{
			Hydrate: listACLBindingRules,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    getACLBindingRule,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the acl_binding_rule.",
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "The description of the acl_binding_rule.",
			},
			{
				Name:        "quota",
				Type:        proto.ColumnType_STRING,
				Description: "The quota of the acl_binding_rule.",
			},
			{
				Name:        "capabilities",
				Type:        proto.ColumnType_JSON,
				Description: "The capabilities of the acl_binding_rule.",
			},
			{
				Name:        "meta",
				Type:        proto.ColumnType_JSON,
				Description: "The metadata associated with the acl_binding_rule.",
			},
			{
				Name:        "create_index",
				Type:        proto.ColumnType_INT,
				Description: "The index when the acl_binding_rule was created.",
			},
			{
				Name:        "modify_index",
				Type:        proto.ColumnType_INT,
				Description: "The index when the acl_binding_rule was last modified.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "The title of the acl_binding_rule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

func listACLBindingRules(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("nomad_acl_binding_rule.listACLBindingRules", "connection_error", err)
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
		bindingRules, metadata, err := client.ACLBindingRules().List(input)
		if err != nil {
			plugin.Logger(ctx).Error("nomad_acl_binding_rule.listACLBindingRules", "query_error", err)
			return nil, err
		}

		for _, bindingRule := range bindingRules {
			d.StreamListItem(ctx, bindingRule)

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

func getACLBindingRule(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	var id string
	if h.Item != nil {
		id = h.Item.(*api.ACLBindingRuleListStub).ID
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
		logger.Error("nomad_acl_binding_rule.getACLBindingRule", "connection_error", err)
		return nil, err
	}

	bindingRule, _, err := client.ACLBindingRules().Get(id, &api.QueryOptions{})
	if err != nil {
		logger.Error("nomad_acl_binding_rule.getACLBindingRule", "api_error", err)
		return nil, err
	}

	return bindingRule, nil
}
