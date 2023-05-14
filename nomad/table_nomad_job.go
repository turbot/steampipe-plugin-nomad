package nomad

import (
	"context"
	"fmt"

	"github.com/hashicorp/nomad/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableNomadJob(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "nomad_job",
		Description: "Retrieve information about your jobs.",
		List: &plugin.ListConfig{
			Hydrate: listJobs,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "namespace",
					Require: plugin.Optional,
				},
				{
					Name:    "create_index",
					Require: plugin.Optional,
				},
				{
					Name:    "name",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getJob,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "Generated UUID for the job.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the job.",
			},
			{
				Name:        "status",
				Description: "The status of the job.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "all_at_once",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether all tasks should be run in parallel or not.",
			},
			{
				Name:        "consul_namespace",
				Description: "The Consul namespace used by the job.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getJob,
			},
			{
				Name:        "consul_token",
				Description: "Consul token used by the job.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getJob,
			},
			{
				Name:        "create_index",
				Description: "Create index of the job.",
				Type:        proto.ColumnType_INT,
				Hydrate:     getJob,
			},
			{
				Name:        "dispatch_idempotency_token",
				Description: "The dispatch idempotency token used by the job.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getJob,
			},
			{
				Name:        "dispatched",
				Description: "Indicates whether the job has been dispatched.",
				Type:        proto.ColumnType_BOOL,
				Hydrate:     getJob,
			},
			{
				Name:        "job_modify_index",
				Description: "Job modify index of the job.",
				Type:        proto.ColumnType_INT,
				Hydrate:     getJob,
			},
			{
				Name:        "modify_index",
				Description: "Modify index of the job.",
				Type:        proto.ColumnType_INT,
				Hydrate:     getJob,
			},
			{
				Name:        "namespace",
				Type:        proto.ColumnType_STRING,
				Description: "The namespace associated with the job.",
			},
			{
				Name:        "parent_id",
				Description: "The parent ID of the job.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ParentID"),
			},
			{
				Name:        "priority",
				Type:        proto.ColumnType_INT,
				Description: "The priority of the job.",
			},
			{
				Name:        "region",
				Type:        proto.ColumnType_STRING,
				Description: "The region where the Nomad job is running.",
				Hydrate:     getJob,
			},
			{
				Name:        "stable",
				Description: "Indicates whether the job is stable.",
				Type:        proto.ColumnType_BOOL,
				Hydrate:     getJob,
			},
			{
				Name:        "status_description",
				Description: "The description of the status of the job.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "stop",
				Description: "Indicates whether the job should be stopped.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "submit_time",
				Description: "The time when the job was submitted.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("SubmitTime").Transform(convertNanoSecToTimestamp),
			},
			{
				Name:        "type",
				Type:        proto.ColumnType_STRING,
				Description: "The type of job.",
			},
			{
				Name:        "vault_namespace",
				Description: "The vault namespace used by the job.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getJob,
			},
			{
				Name:        "vault_token",
				Description: "Vault token used by the job.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getJob,
			},
			{
				Name:        "version",
				Description: "The version of the job.",
				Type:        proto.ColumnType_INT,
				Hydrate:     getJob,
			},
			{
				Name:        "affinities",
				Type:        proto.ColumnType_JSON,
				Description: "The list of affinities for the job.",
				Hydrate:     getJob,
			},
			{
				Name:        "constraints",
				Type:        proto.ColumnType_JSON,
				Description: "The list of constraints for the job.",
				Hydrate:     getJob,
			},
			{
				Name:        "datacenters",
				Type:        proto.ColumnType_JSON,
				Description: "The list of datacenters where the job can be run.",
			},
			{
				Name:        "migrate",
				Type:        proto.ColumnType_JSON,
				Description: "The migration strategy for the job.",
				Hydrate:     getJob,
			},
			{
				Name:        "meta",
				Description: "Metadata associated with the job.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "multiregion",
				Type:        proto.ColumnType_JSON,
				Description: "The multi-region settings for the job.",
				Hydrate:     getJob,
			},
			{
				Name:        "parameterized_job",
				Type:        proto.ColumnType_JSON,
				Description: "The parameterized job configuration for the job.",
			},
			{
				Name:        "payload",
				Description: "The payload of the job.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getJob,
			},
			{
				Name:        "periodic",
				Type:        proto.ColumnType_JSON,
				Description: "The periodic configuration for the job.",
			},
			{
				Name:        "reschedule",
				Type:        proto.ColumnType_JSON,
				Description: "The rescheduling policy for the job.",
				Hydrate:     getJob,
			},
			{
				Name:        "spreads",
				Type:        proto.ColumnType_JSON,
				Description: "The list of spread configurations for the job.",
				Hydrate:     getJob,
			},
			{
				Name:        "task_groups",
				Type:        proto.ColumnType_JSON,
				Description: "The list of task groups for the job.",
				Hydrate:     getJob,
			},
			{
				Name:        "update",
				Type:        proto.ColumnType_JSON,
				Description: "The update strategy for the job.",
				Hydrate:     getJob,
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "The title of the job.",
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
		if *d.QueryContext.Limit < maxLimit {
			maxLimit = *d.QueryContext.Limit
		}
	}

	input := &api.QueryOptions{
		PerPage: int32(maxLimit),
	}

	if d.EqualsQualString("namespace") != "" {
		input.Namespace = d.EqualsQualString("namespace")
	}
	if d.EqualsQuals["create_index"] != nil {
		input.Prefix = d.EqualsQuals["create_index"].GetStringValue()
	}
	if d.EqualsQualString("name") != "" {
		filter := fmt.Sprintf("Name== %q\n", d.EqualsQualString("name"))
		input.Filter = filter
	}

	for {
		jobs, metadata, err := client.Jobs().List(input)
		if err != nil {
			plugin.Logger(ctx).Error("nomad_job.listJobs", "api_error", err)
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

func getJob(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	var id string
	if h.Item != nil {
		id = h.Item.(*api.JobListStub).ID
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
		logger.Error("nomad_node.getJob", "connection_error", err)
		return nil, err
	}

	job, _, err := client.Jobs().Info(id, &api.QueryOptions{})
	if err != nil {
		logger.Error("nomad_node.getJob", "api_error", err)
		return nil, err
	}

	return job, nil
}
