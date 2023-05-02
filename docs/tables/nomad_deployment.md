# Table: nomad_deployment

Deployments are the mechanism by which Nomad rolls out changes to cluster state in a step-by-step fashion. Deployments are only available for Jobs with the type service. When an Evaluation is processed, the scheduler creates only the number of Allocations permitted by the update block and the current state of the cluster. The Deployment is used to monitor the health of those Allocations and emit a new Evaluation for the next step of the update.

## Examples

### Basic info

```sql
select
  id,
  namespace,
  status,
  is_multiregion,
  job_id,
  job_version
from
  nomad_deployment;
```

### List deployments which are not running

```sql
select
  id,
  namespace,
  status,
  is_multiregion,
  job_id,
  job_version
from
  nomad_deployment
where
  status <> 'running';
```

### List multi region deployments

```sql
select
  id,
  namespace,
  status,
  is_multiregion,
  job_id,
  job_version
from
  nomad_deployment
where
  is_multiregion;
```

### Show task group details of the deployments

```sql
select
  id,
  namespace,
  status,
  is_multiregion,
  jsonb_pretty(task_groups) as task_groups
from
  nomad_deployment;
```
