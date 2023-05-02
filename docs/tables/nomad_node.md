# Table: nomad_node

A node refers to a physical or virtual machine that is capable of running jobs. A node is typically a server or a worker machine that is part of a cluster, and it can run one or more Nomad agents.

## Examples

### Basic info

```sql
select
  name,
  id,
  node_class,
  drain,
  status,
  datacenter,
  cgroup_parent
from
  nomad_node;
```

### List nodes with drain set to true

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
  drain;
```

### List nodes which are not ready

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
  status <> 'ready';
```

### List nodes with TLS disabled

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
  not tls_enabled;
```

### List nodes which are eligible for scheduling

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
