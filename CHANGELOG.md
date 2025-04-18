## v1.1.1 [2025-04-18]

_Bug fixes_

- Fixed Linux AMD64 plugin build failures for `Postgres 14 FDW`, `Postgres 15 FDW`, and `SQLite Extension` by upgrading GitHub Actions runners from `ubuntu-20.04` to `ubuntu-22.04`.

## v1.1.0 [2025-04-17]

_Dependencies_

- Recompiled plugin with Go version `1.23.1`. ([#23](https://github.com/turbot/steampipe-plugin-nomad/pull/23))
- Recompiled plugin with [steampipe-plugin-sdk v5.11.5](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.11.5/CHANGELOG.md#v5115-2025-03-31) that addresses critical and high vulnerabilities in dependent packages. ([#23](https://github.com/turbot/steampipe-plugin-nomad/pull/23))

## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#21](https://github.com/turbot/steampipe-plugin-nomad/pull/21))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#21](https://github.com/turbot/steampipe-plugin-nomad/pull/21))

## v0.2.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#17](https://github.com/turbot/steampipe-plugin-nomad/pull/17))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#17](https://github.com/turbot/steampipe-plugin-nomad/pull/17))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-nomad/blob/main/docs/LICENSE). ([#17](https://github.com/turbot/steampipe-plugin-nomad/pull/17))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#16](https://github.com/turbot/steampipe-plugin-nomad/pull/16))

## v0.1.2 [2023-12-06]

_Bug fixes_

- Fixed the invalid Go module path of the plugin. ([#13](https://github.com/turbot/steampipe-plugin-nomad/pull/13))

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
