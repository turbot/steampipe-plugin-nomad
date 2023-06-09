# Table: nomad_namespace

Nomad has support for namespaces, which allow jobs and their associated objects to be segmented from each other and other users of the cluster. Nomad places all jobs and their derived objects into namespaces. These include jobs, allocations, deployments, and evaluations.

You need to specify the `namespace` config argument in the `nomad.spc` file to be able to query this table.

## Examples

### Basic info

```sql
select
  name,
  description,
  create_index,
  modify_index
from
  nomad_namespace;
```

### List the disabled task drivers of namespaces

```sql
select
  name,
  description,
  capabilities -> 'DisabledTaskDrivers' as disabled_task_drivers
from
  nomad_namespace;
```
