package nomad

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-nomad",
		DefaultTransform: transform.FromCamel(),
		DefaultIgnoreConfig: &plugin.IgnoreConfig{
			ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"404"}),
		},
		DefaultRetryConfig: &plugin.RetryConfig{
			ShouldRetryErrorFunc: shouldRetryError([]string{"429"})},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"nomad_node":             tableNomadNode(ctx),
			"nomad_job":              tableNomadJob(ctx),
			"nomad_namespace":        tableNomadNamespace(ctx),
			"nomad_deployment":       tableNomadDeployment(ctx),
			"nomad_acl_auth_method":  tableNomadACLAuthMethod(ctx),
			"nomad_acl_binding_rule": tableNomadACLBindingRule(ctx),
			"nomad_acl_policy":       tableNomadACLPolicy(ctx),
			"nomad_volume":           tableNomadVolume(ctx),
			"nomad_plugin":           tableNomadPlugin(ctx),
			"nomad_agent":            tableNomadAgent(ctx),
		},
	}
	return p
}
