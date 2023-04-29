# Table: nomad_deployment

Nomad deployment refers to the practice of deploying software applications or services in a way that allows them to be easily moved between different computing environments, such as on-premises data centers and public or private clouds, without requiring significant changes to the application or its infrastructure.

## Examples

### Basic info

```sql
select
  id,
  namespace,
  status,
  is_multiregion,
  title
from
  nomad_deployment;
```

### Describe the job the deployment is created for

```sql
select
  id as deployment_id,
  namespace,
  job_id,
  job_create_index,
  job_modify_index,
  job_spec_modify_index,
  job_version
from
  nomad_deployment;
```

### List the deployments in running state

```sql
select
  id,
  namespace,
  status,
  is_multiregion,
  title
from
  nomad_deployment
where
  status = 'running';
```

### List the multiregion deployments

```sql
select
  id,
  namespace,
  status,
  is_multiregion,
  title
from
  nomad_deployment
where
  is_multiregion;
```

### List the deployments with auto revert of tasks enabled

```sql
select
  id,
  namespace,
  status,
  is_multiregion,
  title,
  task_groups -> 'my-task' ->> 'AutoReert' as auto_revert
from
  nomad_deployment
where
  task_groups -> 'my-task' ->> 'AutoRevert' = 'true';
```