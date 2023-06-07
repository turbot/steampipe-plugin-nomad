# Table: nomad_job

A job is a specification provided by users that declares a workload for Nomad. A job is a form of desired state; the user is expressing that the job should be running, but not where it should be run. The responsibility of Nomad is to make sure the actual state matches the user desired state. A job is composed of one or more task groups.

## Examples

### Basic info

```sql
select
  id,
  name,
  namespace,
  status,
  type,
  region,
  modify_index,
  submit_time
from
  nomad_job;
```

### List unstable jobs

```sql
select
  id,
  name,
  namespace,
  status,
  type,
  region
from
  nomad_job
where
  not stable;
```

### List multi-region jobs

```sql
select
  id,
  name,
  namespace,
  status,
  region,
  multiregion
from
  nomad_job
where
  multiregion is not null;
```

### List pending jobs

```sql
select
  id,
  name,
  namespace,
  status,
  region,
  multiregion
from
  nomad_job
where
  status = 'pending';
```

### List jobs with `autorevert` enabled

```sql
select
  id,
  name,
  namespace,
  status,
  update ->> 'AutoRevert' as auto_revert
from
  nomad_job
where
  update ->> 'AutoRevert' = 'true';
```

### Show the CSI plugin configuration of the jobs

```sql
select
  id as job_id,
  name as job_name,
  t -> 'CSIPluginConfig' ->> 'ID' as csi_plugin_id,
  t -> 'CSIPluginConfig' ->> 'Type' as csi_plugin_type,
  t -> 'CSIPluginConfig' ->> 'HealthTimeout' as csi_plugin_timeout
from
  nomad_job,
  jsonb_array_elements(task_groups) as tg,
  jsonb_array_elements(tg -> 'Tasks') as t;
```
