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
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getACLBindingRule,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "An internally generated UUID for this rule and is controlled by Nomad.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "The description of the acl_binding_rule.",
			},
			{
				Name:        "auth_method",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the auth method for which this rule applies to.",
			},
			{
				Name:        "selector",
				Type:        proto.ColumnType_STRING,
				Description: "An expression that matches against verified identity attributes returned from the auth method during login.",
				Hydrate:     getACLBindingRule,
			},
			{
				Name:        "bind_type",
				Type:        proto.ColumnType_STRING,
				Description: "Adjusts how this binding rule is applied at login time.",
				Hydrate:     getACLBindingRule,
			},
			{
				Name:        "bind_name",
				Type:        proto.ColumnType_STRING,
				Description: "BindName is the target of the binding.",
				Hydrate:     getACLBindingRule,
			},
			{
				Name:        "create_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Create time of the ACL binding rule.",
				Hydrate:     getACLBindingRule,
			},
			{
				Name:        "modify_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Last modify time of the ACL binding rule.",
				Hydrate:     getACLBindingRule,
			},
			{
				Name:        "create_index",
				Type:        proto.ColumnType_INT,
				Description: "Create index of the ACL binding rule.",
			},
			{
				Name:        "modify_index",
				Type:        proto.ColumnType_INT,
				Description: "Modify index of the ACL binding rule.",
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
