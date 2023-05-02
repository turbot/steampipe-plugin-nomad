# Table: nomad_job

The Nomad job specification (or "jobspec" for short) defines the schema for Nomad jobs. Nomad jobs are specified in HCL, which aims to strike a balance between human readable and editable, and machine-friendly.

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
  submit_time,
  title
from
  nomad_job;
```

### List all the stable jobs

```sql
select
  id,
  name,
  namespace,
  status,
  type,
  region,
  modify_index,
  title
from
  nomad_job
where
  stable;
```

### List the multiregion jobs

```sql
select
  id,
  name,
  namespace,
  status,
  type,
  region,
  modify_index,
  title
from
  nomad_job
where
  multiregion is not null;
```

### List the jobs with autorevert enabled

```sql
select
  id,
  name,
  namespace,
  status,
  type,
  region,
  modify_index,
  title,
  update ->> 'AutoRevert' as auto_revert
from
  nomad_job
where
  update ->> 'AutoRevert' = 'true';
```

### Describe the CSI plugfin configuration for the job

```sql
select
  id as job_id,
  name as job_name,
  t -> 'CSIPluginConfig' -> 'ID' as csi_plugin_id,
  t -> 'CSIPluginConfig' -> 'Type' as csi_plugin_type,
  t -> 'CSIPluginConfig' -> 'HealthTimeout' as csi_plugin_timeout
from
  nomad_job,
  jsonb_array_elements(task_groups) as tg,
  jsonb_array_elements(tg -> 'Tasks') as t
```