# Table: nomad_plugin

CSI is the container storage interface. It is a plugin for Kubernetes and other container orchestrators that allows storage suppliers to expose their products to containerised applications as persistent storage.

## Examples

### Basic info

```sql
select
  id,
  title,
  version,
  create_index,
  modify_index,
  provider
from
  nomad_plugin;
```

### List the CSI plugins that require a controller

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
  controller_rquired;
```

### Find the number of nodes and controllers associated with each CSI plugin

```sql
select
  id,
  title,
  nodes_healthy,
  controllers_healthy
from
  csi_plugins;
```

### Find the total number of expected and healthy nodes for each CSI plugin

```sql
select
  id,
  title,
  nodes_healthy,
  nodes_expected
from
  csi_plugins;
```

### Find the plugin with the highest number of expected nodes

```sql
select
  id,
  title,
  nodes_healthy,
  nodes_expected
from
  csi_plugins
order by
  nodes_expected desc limit 1;
```
