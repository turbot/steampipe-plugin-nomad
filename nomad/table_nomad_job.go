package nomad

import (
	"context"

	"github.com/hashicorp/nomad/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableNomadJob(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "nomad_job",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate: listJobs,
		},
		// Get: &plugin.GetConfig{
		// 	KeyColumns: plugin.SingleColumn("id"),
		// 	Hydrate: getNode,
		// },
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The UUID of the vault.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "parent_id",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the vault.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The version of the vault metadata.",
			},
			{
				Name:        "namespace",
				Type:        proto.ColumnType_STRING,
				Description: "The version of the vault metadata.",
			},
			{
				Name:        "datacentre",
				Type:        proto.ColumnType_JSON,
				Description: "The version of the vault metadata.",
			},
			{
				Name:        "type",
				Type:        proto.ColumnType_STRING,
				Description: "The version of the vault metadata.",
			},
			{
				Name:        "priority",
				Type:        proto.ColumnType_INT,
				Description: "The version of the vault metadata.",
			},
			{
				Name:        "periodic",
				Type:        proto.ColumnType_BOOL,
				Description: "The version of the vault metadata.",
			},
			{
				Name:        "parameterized_job",
				Type:        proto.ColumnType_BOOL,
				Description: "Date and time when the vault was created.",
			},
			{
				Name:        "stop",
				Type:        proto.ColumnType_BOOL,
				Description: "The version of the vault contents.",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "The description for the vault.",
			},
			{
				Name:        "status_description",
				Type:        proto.ColumnType_STRING,
				Description: "Number of active items in the vault.",
			},
			{
				Name:        "job_summary",
				Type:        proto.ColumnType_JSON,
				Description: "The type of vault. Possible values are EVERYONE, PERSONAL and USER_CREATED.",
			},
			{
				Name:        "create_index",
				Type:        proto.ColumnType_INT,
				Description: "Date and time when the vault or its contents were last changed.",
			},
			{
				Name:        "modify_index",
				Type:        proto.ColumnType_INT,
				Description: "Date and time when the vault or its contents were last changed.",
			},
			{
				Name:        "submit_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Date and time when the vault or its contents were last changed.",
				Transform:   transform.FromField("SubmitTime").Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "meta",
				Type:        proto.ColumnType_JSON,
				Description: "Date and time when the vault or its contents were last changed.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "The title of the vault.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

func listJobs(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("nomad_job.listJobs", "connection_error", err)
		return nil, err
	}

		maxLimit := int64(1000)
		if d.QueryContext.Limit != nil {
			if (*d.QueryContext.Limit < maxLimit){
				maxLimit = *d.QueryContext.Limit
			}
		}

	jobClient := client.Jobs()
	input := &api.QueryOptions{
		PerPage: int32(maxLimit),
	}

	// Add support for optional region qual
	if d.EqualsQuals["datacenter"] != nil {
		input.Region = d.EqualsQuals["datacenter"].GetStringValue()
	}

	for {
		jobs, metadata, err := jobClient.List(input)
		if err != nil {
			plugin.Logger(ctx).Error("nomad_job.listJobs", "query_error", err)
			return nil, err
		}

		for _, job := range jobs {
			d.StreamListItem(ctx, job)

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

// func getNode(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
// 	logger := plugin.Logger(ctx)
// 	id := d.EqualsQualString("id")
// 	// Create client
// 	input := &api.QueryOptions{}
// 	client, err := getClient(ctx, d)
// 	nodeClient := client.Nodes()
// 	if err != nil {
// 		logger.Error("nomad_node.getNode", "connection_error", err)
// 		return nil, err
// 	}

// 	node, _, err := nodeClient.Info(id, input)
// 	if err != nil {
// 		logger.Error("nomad_node.getNode", "api_error", err)
// 		return nil, err
// 	}

// 	return node, nil
// }
