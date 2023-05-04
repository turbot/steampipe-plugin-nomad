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
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "namespace",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getVolume,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The unique identifier of the CSI volume.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the CSI volume.",
			},
			{
				Name:        "external_id",
				Type:        proto.ColumnType_STRING,
				Description: "The external ID of the CSI volume.",
				Transform:   transform.FromField("ExternalID"),
			},
			{
				Name:        "namespace",
				Type:        proto.ColumnType_STRING,
				Description: "The namespace of the CSI volume.",
			},
			{
				Name:        "capacity",
				Type:        proto.ColumnType_INT,
				Description: "The capacity of the CSI volume.",
				Hydrate:     getVolume,
			},
			{
				Name:        "requested_capacity_min",
				Type:        proto.ColumnType_INT,
				Description: "The minimum requested capacity of the CSI volume.",
				Hydrate:     getVolume,
			},
			{
				Name:        "requested_capacity_max",
				Type:        proto.ColumnType_INT,
				Description: "The maximum requested capacity of the CSI volume.",
				Hydrate:     getVolume,
			},
			{
				Name:        "clone_id",
				Type:        proto.ColumnType_STRING,
				Description: "The clone ID of the CSI volume.",
				Transform:   transform.FromField("CloneID"),
				Hydrate:     getVolume,
			},
			{
				Name:        "snapshot_id",
				Type:        proto.ColumnType_STRING,
				Description: "The snapshot ID of the CSI volume.",
				Transform:   transform.FromField("SnapshotID"),
				Hydrate:     getVolume,
			},
			{
				Name:        "schedulable",
				Type:        proto.ColumnType_BOOL,
				Description: "A flag that indicates if all the denormalized plugin health fields are true.",
			},
			{
				Name:        "plugin_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the CSI plugin.",
				Transform:   transform.FromField("PluginID"),
			},
			{
				Name:        "provider",
				Type:        proto.ColumnType_STRING,
				Description: "The provider of the CSI volume.",
			},
			{
				Name:        "provider_version",
				Type:        proto.ColumnType_STRING,
				Description: "The version of the provider of the CSI volume.",
			},
			{
				Name:        "controller_required",
				Type:        proto.ColumnType_BOOL,
				Description: "A flag that indicates if a controller is required for the CSI volume.",
			},
			{
				Name:        "controllers_healthy",
				Type:        proto.ColumnType_INT,
				Description: "The number of healthy controllers for the CSI volume.",
			},
			{
				Name:        "controllers_expected",
				Type:        proto.ColumnType_INT,
				Description: "The expected number of controllers for the CSI volume.",
			},
			{
				Name:        "nodes_healthy",
				Type:        proto.ColumnType_INT,
				Description: "The number of healthy nodes for the CSI volume.",
			},
			{
				Name:        "nodes_expected",
				Type:        proto.ColumnType_INT,
				Description: "The expected number of nodes for the CSI volume.",
			},
			{
				Name:        "resource_exhausted",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time when the CSI volume's resource is exhausted.",
			},
			{
				Name:        "create_index",
				Type:        proto.ColumnType_INT,
				Description: "The index of when the CSI volume was created.",
			},
			{
				Name:        "modify_index",
				Type:        proto.ColumnType_INT,
				Description: "The index of when the CSI volume was last modified.",
			},
			{
				Name:        "requested_topologies",
				Type:        proto.ColumnType_JSON,
				Description: "The topologies that were submitted as options to the storage provider at the time the volume was created.",
			},
			{
				Name:        "topologies",
				Type:        proto.ColumnType_JSON,
				Description: "The topologies returned by the storage provider based on the RequestedTopologies and what the storage provider could support.",
			},
			{
				Name:        "access_mode",
				Type:        proto.ColumnType_JSON,
				Description: "The access mode of the CSI volume.",
			},
			{
				Name:        "attachment_mode",
				Type:        proto.ColumnType_JSON,
				Description: "The attachment mode of the CSI volume.",
			},
			{
				Name:        "mount_options",
				Type:        proto.ColumnType_JSON,
				Description: "The mount options of the CSI volume.",
				Hydrate:     getVolume,
			},
			{
				Name:        "secrets",
				Type:        proto.ColumnType_JSON,
				Description: "The secrets of the CSI volume.",
				Hydrate:     getVolume,
			},
			{
				Name:        "parameters",
				Type:        proto.ColumnType_JSON,
				Description: "The parameters of the CSI volume.",
				Hydrate:     getVolume,
			},
			{
				Name:        "context",
				Type:        proto.ColumnType_JSON,
				Description: "The context of the CSI volume.",
				Hydrate:     getVolume,
			},
			{
				Name:        "requested_capabilities",
				Type:        proto.ColumnType_JSON,
				Description: "The requested capabilities of the CSI volume.",
			},
			{
				Name:        "read_allocs",
				Type:        proto.ColumnType_JSON,
				Description: "The map of allocation IDs for tracking reader claim status.",
				Hydrate:     getVolume,
			},
			{
				Name:        "write_allocs",
				Type:        proto.ColumnType_JSON,
				Description: "The map of allocation IDs for tracking writer claim status.",
				Hydrate:     getVolume,
			},
			{
				Name:        "allocations",
				Type:        proto.ColumnType_JSON,
				Description: "The list of combined readers and writers allocations.",
				Hydrate:     getVolume,
			},
			{
				Name:        "extra_keys_hcl",
				Type:        proto.ColumnType_JSON,
				Description: "The list of extra keys used by the hcl parser to report unexpected keys.",
				Transform:   transform.FromField("ExtraKeysHCL"),
				Hydrate:     getVolume,
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

	if d.EqualsQuals["namespace"] != nil {
		input.Namespace = d.EqualsQualString("namespace")
	}

	for {
		volumes, metadata, err := client.CSIVolumes().List(input)
		if err != nil {
			plugin.Logger(ctx).Error("nomad_volume.listVolumes", "api_error", err)
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
