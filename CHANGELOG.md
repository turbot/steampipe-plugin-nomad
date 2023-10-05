## v0.1.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#6](https://github.com/turbot/steampipe-plugin-nomad/pull/6))

## v0.1.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#4](https://github.com/turbot/steampipe-plugin-nomad/pull/4))
- Recompiled plugin with Go version `1.21`. ([#4](https://github.com/turbot/steampipe-plugin-nomad/pull/4))

## v0.0.1 [2023-06-07]

_What's new?_

- New tables added
  - [nomad_acl_auth_method](https://hub.steampipe.io/plugins/turbot/nomad/tables/nomad_acl_auth_method)
  - [nomad_acl_binding_rule](https://hub.steampipe.io/plugins/turbot/nomad/tables/nomad_acl_binding_rule)
  - [nomad_acl_policy](https://hub.steampipe.io/plugins/turbot/nomad/tables/nomad_acl_policy)
  - [nomad_acl_role](https://hub.steampipe.io/plugins/turbot/nomad/tables/nomad_acl_role)
  - [nomad_acl_token](https://hub.steampipe.io/plugins/turbot/nomad/tables/nomad_acl_token)
  - [nomad_agent_member](https://hub.steampipe.io/plugins/turbot/nomad/tables/nomad_agent_member)
  - [nomad_deployment](https://hub.steampipe.io/plugins/turbot/nomad/tables/nomad_deployment)
  - [nomad_job](https://hub.steampipe.io/plugins/turbot/nomad/tables/nomad_job)
  - [nomad_namespace](https://hub.steampipe.io/plugins/turbot/nomad/tables/nomad_namespace)
  - [nomad_node](https://hub.steampipe.io/plugins/turbot/nomad/tables/nomad_node)
  - [nomad_plugin](https://hub.steampipe.io/plugins/turbot/nomad/tables/nomad_plugin)
  - [nomad_volume](https://hub.steampipe.io/plugins/turbot/nomad/tables/nomad_volume)
