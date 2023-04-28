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
		Description: "Retrieve information about your jobs.",
		List: &plugin.ListConfig{
			Hydrate: listJobs,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getJob,
		},
		Columns: []*plugin.Column{
			{
				Name:        "region",
				Type:        proto.ColumnType_STRING,
				Description: "The region where the Nomad client is running.",
				Hydrate:     getJob,
			},
			{
				Name:        "namespace",
				Type:        proto.ColumnType_STRING,
				Description: "The namespace associated with the job.",
			},
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "Generated UUID for the deployment.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the job.",
			},
			{
				Name:        "type",
				Type:        proto.ColumnType_STRING,
				Description: "The type of job.",
			},
			{
				Name:        "priority",
				Type:        proto.ColumnType_INT,
				Description: "The priority of the job.",
			},
			{
				Name:        "all_at_once",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether all tasks should be run in parallel or not.",
			},
			{
				Name:        "datacenters",
				Type:        proto.ColumnType_JSON,
				Description: "The list of datacenters where the job can be run.",
			},
			{
				Name:        "constraints",
				Type:        proto.ColumnType_JSON,
				Description: "The list of constraints for the job.",
				Hydrate:     getJob,
			},
			{
				Name:        "affinities",
				Type:        proto.ColumnType_JSON,
				Description: "The list of affinities for the job.",
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
			{
				Name:        "multiregion",
				Type:        proto.ColumnType_JSON,
				Description: "The multi-region settings for the job.",
				Hydrate:     getJob,
			},
			{
				Name:        "spreads",
				Type:        proto.ColumnType_JSON,
				Description: "The list of spread configurations for the job.",
				Hydrate:     getJob,
			},
			{
				Name:        "periodic",
				Type:        proto.ColumnType_JSON,
				Description: "The periodic configuration for the job.",
			},
			{
				Name:        "parameterized_job",
				Type:        proto.ColumnType_JSON,
				Description: "The parameterized job configuration for the job.",
			},
			{
				Name:        "reschedule",
				Type:        proto.ColumnType_JSON,
				Description: "The rescheduling policy for the job.",
				Hydrate:     getJob,
			},
			{
				Name:        "migrate",
				Type:        proto.ColumnType_JSON,
				Description: "The migration strategy for the job.",
				Hydrate:     getJob,
			},
			{
				Name:        "meta",
				Description: "Metadata associated with the test",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "consul_token",
				Description: "Consul token used by the test",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getJob,
			},
			{
				Name:        "vault_token",
				Description: "Vault token used by the test",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getJob,
			},
			{
				Name:        "stop",
				Description: "Indicates whether the test should be stopped",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "parent_id",
				Description: "The parent ID of the test",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ParentID"),
				Hydrate:     getJob,
			},
			{
				Name:        "dispatched",
				Description: "Indicates whether the test has been dispatched",
				Type:        proto.ColumnType_BOOL,
				Hydrate:     getJob,
			},
			{
				Name:        "dispatch_idempotency_token",
				Description: "The dispatch idempotency token used by the test",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getJob,
			},
			{
				Name:        "payload",
				Description: "The payload of the test",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getJob,
			},
			{
				Name:        "consul_namespace",
				Description: "The Consul namespace used by the test",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getJob,
			},
			{
				Name:        "vault_namespace",
				Description: "The Vault namespace used by the test",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getJob,
			},
			{
				Name:        "nomad_token_id",
				Description: "The Nomad token ID used by the test",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("NomadTokenID"),
				Hydrate:     getJob,
			},
			{
				Name:        "status",
				Description: "The status of the test",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "status_description",
				Description: "The description of the status of the test",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "stable",
				Description: "Indicates whether the test is stable",
				Type:        proto.ColumnType_BOOL,
				Hydrate:     getJob,
			},
			{
				Name:        "version",
				Description: "The version of the test",
				Type:        proto.ColumnType_INT,
				Hydrate:     getJob,
			},
			{
				Name:        "submit_time",
				Description: "The time when the test was submitted",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("SubmitTime").Transform(convertJobSubmitTimestamp),
			},
			{
				Name:        "create_index",
				Description: "The create index of the test",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "modify_index",
				Description: "The modify index of the test",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "job_modify_index",
				Description: "The modify index of the job associated with the test",
				Type:        proto.ColumnType_INT,
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

	// Add support for optional region qual
	if d.EqualsQuals["datacenter"] != nil {
		input.Region = d.EqualsQuals["datacenter"].GetStringValue()
	}

	for {
		jobs, metadata, err := client.Jobs().List(input)
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
