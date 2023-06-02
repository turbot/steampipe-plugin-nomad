![image](https://hub.steampipe.io/images/plugins/turbot/nomad-social-graphic.png)

# Nomad Plugin for Steampipe

Use SQL to query nodes, jobs, deployments and more from Nomad.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/nomad)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/nomad/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-nomad/issues)

## Quick start

### Install

Download and install the latest Nomad plugin:

```bash
steampipe plugin install nomad
```

Configure your [credentials](https://hub.steampipe.io/plugins/turbot/nomad#credentials) and [config file](https://hub.steampipe.io/plugins/turbot/nomad#configuration).

Add your configuration details in `~/.steampipe/config/nomad.spc`:

```hcl
connection "nomad" {
  plugin    = "nomad"
  # Authentication information
  address   = "http://52.14.112.248:8500"
  token     = "c178b810-8b18-6f38-016f-725ddec5d58"
}
```

Or through environment variables:

```sh
export NOMAD_ADDR=http://18.118.144.168:4646
export NOMAD_TOKEN=c178b810-8b18-6f38-016f-725ddec5d58
```

Run steampipe:

```shell
steampipe query
```

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

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-nomad.git
cd steampipe-plugin-nomad
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/nomad.spc
```

Try it!

```
steampipe query
> .inspect nomad
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-nomad/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Nomad Plugin](https://github.com/turbot/steampipe-plugin-nomad/labels/help%20wanted)
