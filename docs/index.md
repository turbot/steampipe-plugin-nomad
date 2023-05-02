---
organization: Turbot
category: ["software development"]
icon_url: "/images/plugins/turbot/nomad.svg"
brand_color: "#00CA8E"
display_name: "Nomad"
short_name: "nomad"
description: "Steampipe plugin to query nodes, jobs, deployments and more from Nomad."
og_description: "Query Nomad with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/nomad-social-graphic.png"
---

# Nomad + Steampipe

[Nomad](https://www.nomadproject.io/) is a simple and flexible scheduler and orchestrator for managing containers and non-containerized applications across on-prem and clouds at scale.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

List your Nomad jobs:

```sql
select
  id,
  name,
  status,
  dispatched,
  namespace,
  priority,
  region
from
  nomad_job;
```

```
+------+------+---------+------------+-----------+----------+--------+
| id   | name | status  | dispatched | namespace | priority | region |
+------+------+---------+------------+-----------+----------+--------+
| docs | docs | pending | false      | default   | 50       | global |
+------+------+---------+------------+-----------+----------+--------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/nomad/tables)**

## Quick start

### Install

Download and install the latest Nomad plugin:

```sh
steampipe plugin install nomad
```

### Credentials

| Item        | Description                                                                                                                                                                              |
| ----------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Credentials | Nomad requires an Address and Namespace or Address, Namespace and [Secret ID](https://developer.hashicorp.com/nomad/tutorials/access-control/access-control-tokens) for all requests.    |
| Permissions | The permission scope of Secret IDs is set by the Admin at the creation time of the ACL tokens.                                                                                           |
| Radius      | Each connection represents a single Nomad Installation.                                                                                                                                  |
| Resolution  | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/nomad.spc`)<br />2. Credentials specified in environment variables, e.g., `NOMAD_ADDR` and `NOMAD_TOKEN`. |

### Configuration

Installing the latest nomad plugin will create a config file (`~/.steampipe/config/nomad.spc`) with a single connection named `nomad`:

Configure your account details in `~/.steampipe/config/nomad.spc`:

```hcl
connection "nomad" {
  plugin = "nomad"

  # `address` - The address of the Nomad server.
  # Can also be set with the NOMAD_ADDR environment variable.
  # address = "http://18.118.164.168:4646"

  # `namespace` - The Nomad cluster namespace.
  # For more information on the Namespace, please see https://developer.hashicorp.com/nomad/tutorials/manage-clusters/namespaces.
  # Can also be set with the NOMAD_NAMESPACE environment variable.
  # "*" indicates all the namespaces available.
  namespace = "*"

  # `secret_id` - The SecretID of an ACL token.
  # The SecretID is required to make requests for ACL-enabled clusters.
  # For more information on the ACL Token, please see https://developer.hashicorp.com/nomad/tutorials/access-control/access-control-tokens.
  # Can also be set with the NOMAD_TOKEN environment variable.
  # secret_id = "c178b810-8b18-6f38-016f-725ddec5d58"
}
```

## Configuring Nomad Credentials

You may specify the Address and Namespace to authenticate:

- `address`: The address of the nomad server.
- `namespace`: The Nomad Cluster namespace.

```hcl
connection "nomad" {
  plugin    = "nomad"
  address   = "http://18.118.144.168:4646"
  namespace = "*"
}
```

or you may specify the Address, Namespace and SecretID to authenticate:

- `address`: The address of the nomad server.
- `namespace`: The Nomad Cluster namespace.
- `secret_id`: The SecretID of an ACL token.

```hcl
connection "nomad" {
  plugin    = "nomad"
  address   = "http://18.118.144.168:4646"
  namespace = "*"
  secret_id = "c178b810-8b18-6f38-016f-725ddec5d58"
}
```

or through environment variables

The Nomad plugin will use the Nomad environment variable to obtain credentials **only if the `address`,`namespace` and `secret_id` is not specified** in the connection:

```sh
export NOMAD_ADDR="http://18.118.144.168:4646"
export NOMAD_NAMESPACE="*"
export NOMAD_TOKEN="c178b810-8b18-6f38-016f-725ddec5d58"
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-nomad
- Community: [Slack Channel](https://steampipe.io/community/join)
