package nomad

import (
	"context"

	"github.com/hashicorp/nomad/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableNomadNamespace(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "nomad_namespace",
		Description: "Retrieve information about your namespaces.",
		List: &plugin.ListConfig{
			Hydrate: listNamespaces,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    getNamespace,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "A string representing the name of the namespace.",
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "A string providing a description or summary of the namespace.",
			},
			{
				Name:        "quota",
				Type:        proto.ColumnType_STRING,
				Description: "A string specifying the maximum usage limit for the namespace.",
			},
			{
				Name:        "create_index",
				Type:        proto.ColumnType_INT,
				Description: "An unsigned 64-bit integer representing the index at which the namespace was created.",
			},
			{
				Name:        "modify_index",
				Type:        proto.ColumnType_INT,
				Description: "An unsigned 64-bit integer representing the index at which the namespace was last modified.",
			},
			{
				Name:        "capabilities",
				Type:        proto.ColumnType_JSON,
				Description: "A pointer to a NamespaceCapabilities struct that defines the capabilities of the namespace.",
			},
			{
				Name:        "meta",
				Type:        proto.ColumnType_JSON,
				Description: "A map containing additional metadata associated with the namespace.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "The title of the namespace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

func listNamespaces(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("nomad_namespace.listNamespaces", "connection_error", err)
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
		namespaces, metadata, err := client.Namespaces().List(input)
		if err != nil {
			plugin.Logger(ctx).Error("nomad_namespace.listNamespaces", "api_error", err)
			return nil, err
		}

		for _, namespace := range namespaces {
			d.StreamListItem(ctx, namespace)

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

func getNamespace(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	name := d.EqualsQualString("name")

	// check if name is empty
	if name == "" {
		return nil, nil
	}

	// Create client
	client, err := getClient(ctx, d)
	if err != nil {
		logger.Error("nomad_node.getNamespace", "connection_error", err)
		return nil, err
	}

	namespace, _, err := client.Namespaces().Info(name, &api.QueryOptions{})
	if err != nil {
		logger.Error("nomad_node.getNamespace", "api_error", err)
		return nil, err
	}

	return namespace, nil
}
