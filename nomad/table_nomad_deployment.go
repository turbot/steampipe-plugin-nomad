package nomad

import (
	"context"

	"github.com/hashicorp/nomad/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableNomadDeployment(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "nomad_deployment",
		Description: "Retrieve information about your deployments.",
		List: &plugin.ListConfig{
			Hydrate: listDeployments,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "namespace",
					Require: plugin.Optional,
				},
				{
					Name:    "status",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getDeployment,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "Generated UUID for the deployment.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "job_id",
				Type:        proto.ColumnType_STRING,
				Description: "Job the deployment is created for.",
				Transform:   transform.FromField("JobID"),
			},
			{
				Name:        "is_multiregion",
				Type:        proto.ColumnType_BOOL,
				Description: "Specifies if this deployment is part of a multi-region deployment.",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "Status of the deployment.",
			},
			{
				Name:        "create_index",
				Type:        proto.ColumnType_INT,
				Description: "Create index of the deployment.",
			},
			{
				Name:        "job_create_index",
				Type:        proto.ColumnType_INT,
				Description: "Create index of the job which the deployment is tracking.",
			},
			{
				Name:        "job_modify_index",
				Type:        proto.ColumnType_INT,
				Description: "ModifyIndex of the job which the deployment is tracking.",
			},
			{
				Name:        "job_spec_modify_index",
				Type:        proto.ColumnType_INT,
				Description: "JobModifyIndex of the job which the deployment is tracking.",
			},
			{
				Name:        "job_version",
				Type:        proto.ColumnType_INT,
				Description: "Version of the job at which the deployment is tracking.",
			},
			{
				Name:        "modify_index",
				Type:        proto.ColumnType_INT,
				Description: "ModifyIndex of the deployment.",
			},
			{
				Name:        "namespace",
				Type:        proto.ColumnType_STRING,
				Description: "Namespace the deployment is created in.",
			},
			{
				Name:        "status_description",
				Type:        proto.ColumnType_STRING,
				Description: "Human-readable description of the deployment status.",
			},
			{
				Name:        "task_groups",
				Type:        proto.ColumnType_JSON,
				Description: "Set of task groups effected by the deployment and their current deployment status.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "The title of the deployment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
		},
	}
}

func listDeployments(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("nomad_deployment.listDeployments", "connection_error", err)
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

	if d.EqualsQuals["namespace"] != nil {
		input.Namespace = d.EqualsQualString("namespace")
	}

	for {
		deployments, metadata, err := client.Deployments().List(input)
		if err != nil {
			plugin.Logger(ctx).Error("nomad_deployment.listDeployments", "api_error", err)
			return nil, err
		}

		for _, deployment := range deployments {
			d.StreamListItem(ctx, deployment)

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

func getDeployment(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("id")

	// check if id is empty
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := getClient(ctx, d)
	if err != nil {
		logger.Error("nomad_node.getDeployment", "connection_error", err)
		return nil, err
	}

	deployment, _, err := client.Deployments().Info(id, &api.QueryOptions{})
	if err != nil {
		logger.Error("nomad_node.getDeployment", "api_error", err)
		return nil, err
	}

	return deployment, nil
}
