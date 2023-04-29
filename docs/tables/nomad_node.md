# Table: nomad_node

Nomad Nodes are designed to keep both the end user and the node provider safe. All in one App.

## Examples

### Basic info

```sql
select
  name,
  id,
  node_class,
  status,
  datacenter,
  cgroup_parent
from
  nomad_node;
```

### List all nodes with drain set to true

```sql
select
  name,
  id,
  node_class,
  status,
  datacenter,
  cgroup_parent
from
  nomad_node
where
  drain = true;
```

### List all nodes with TLS enabled

```sql
select
  name,
  id,
  node_class,
  status,
  datacenter,
  cgroup_parent
from
  nomad_node
where
  tls_enabled = true;
```

### List all nodes with the scheduling eligibility set to "eligible"

```sql
select
  name,
  id,
  node_class,
  status,
  datacenter,
  cgroup_parent
from
  nomad_node
where
  scheduling_eligibility = 'eligible';
```

### List all nodes with the drain strategy set to "graceful"

```sql
select
  name,
  id,
  node_class,
  status,
  datacenter,
  cgroup_parent
from
  nomad_node
where
  drain_strategy ->> 'type' = 'graceful';
```

### List the nodes that have been drained in the last week, along with their drain metadata

```sql
select
  name,
  id,
  node_class,
  status,
  datacenter,
  cgroup_parent
from
  nomad_node
where
  (last_drain ->> 'end_time') > (NOW() - INTERVAL '7 days')::text
  and drain = true;
```