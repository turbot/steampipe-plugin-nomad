# Table: nomad_plugin

Nomad CSI (Container Storage Interface) plugin is a plugin architecture for integrating container orchestration systems like Nomad with different storage providers. CSI allows for the abstraction of storage vendors, creating a standard interface for container orchestration systems to manage storage resources across various infrastructure providers.

## Examples

### Basic info

```sql
select
  id,
  controller_required,
  version,
  create_index,
  modify_index,
  provider
from
  nomad_plugin;
```

### List CSI plugins that require a controller

```sql
select
  id,
  title,
  version,
  create_index,
  modify_index,
  provider
from
  nomad_plugin
where
  controller_required;
```

### Show the number of nodes and controllers associated with each plugin

```sql
select
  id,
  version,
  nodes_healthy,
  controllers_healthy
from
  nomad_plugin;
```

### Show the number of expected nodes for each plugin

```sql
select
  id,
  version,
  nodes_healthy,
  nodes_expected
from
  nomad_plugin;
```
