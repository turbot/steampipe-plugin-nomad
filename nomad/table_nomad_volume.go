package nomad

import (
	"context"

	"github.com/hashicorp/nomad/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableNomadVolume(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "nomad_volume",
		Description: "Retrieve information about your volumes.",
		List: &plugin.ListConfig{
			Hydrate: listVolumes,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getVolume,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the volume.",
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "The description of the volume.",
			},
			{
				Name:        "quota",
				Type:        proto.ColumnType_STRING,
				Description: "The quota of the volume.",
			},
			{
				Name:        "capabilities",
				Type:        proto.ColumnType_JSON,
				Description: "The capabilities of the volume.",
			},
			{
				Name:        "meta",
				Type:        proto.ColumnType_JSON,
				Description: "The metadata associated with the volume.",
			},
			{
				Name:        "create_index",
				Type:        proto.ColumnType_INT,
				Description: "The index when the volume was created.",
			},
			{
				Name:        "modify_index",
				Type:        proto.ColumnType_INT,
				Description: "The index when the volume was last modified.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "The title of the volume.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

func listVolumes(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("nomad_volume.listVolumes", "connection_error", err)
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
		volumes, metadata, err := client.CSIVolumes().List(input)
		if err != nil {
			plugin.Logger(ctx).Error("nomad_volume.listVolumes", "query_error", err)
			return nil, err
		}

		for _, volume := range volumes {
			d.StreamListItem(ctx, volume)

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

func getVolume(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	var id string
	if h.Item != nil {
		id = h.Item.(*api.CSIVolumeListStub).ID
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
		logger.Error("nomad_node.getVolume", "connection_error", err)
		return nil, err
	}

	volume, _, err := client.CSIVolumes().Info(id, &api.QueryOptions{})
	if err != nil {
		logger.Error("nomad_node.getVolume", "api_error", err)
		return nil, err
	}

	return volume, nil
}
