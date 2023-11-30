---
title: "Steampipe Table: nomad_job - Query Nomad Jobs using SQL"
description: "Allows users to query Nomad Jobs, specifically their configurations, statuses, and other metadata, providing insights into job orchestration and management."
---

# Table: nomad_job - Query Nomad Jobs using SQL

Nomad is a simple and flexible workload orchestrator to deploy and manage containers and non-containerized applications across on-prem and clouds at scale. Nomad Jobs are the primary configuration that users interact with when using Nomad. These jobs are specified in files which are then passed to Nomad to schedule onto the cluster.

## Table Usage Guide

The `nomad_job` table provides insights into the jobs scheduled within HashiCorp Nomad. As a DevOps engineer, explore job-specific details through this table, including job configurations, statuses, and other metadata. Utilize it to manage and monitor your workloads, understand job dependencies, and optimize resource allocation.

## Examples

### Basic info
Explore which jobs are currently active by determining their status and type. This can help in understanding the workload distribution across different regions, and track any changes over time.

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
Discover the segments that contain unstable jobs across various regions and namespaces, helping you identify areas that may require troubleshooting or further investigation.

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
Discover the segments that have jobs spanning multiple regions, enabling you to manage and monitor tasks across different geographical areas more effectively.

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
Discover the segments that consist of jobs awaiting execution, enabling you to manage and prioritize your workflow more effectively. This can be particularly useful in scenarios where resource allocation and task scheduling are critical.

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
Explore which jobs have the 'autorevert' feature enabled. This can be particularly useful for understanding and managing job configurations in a Nomad cluster.

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
Analyze the configuration of job-specific CSI plugins to understand their types and health timeout settings. This can help in monitoring and managing the performance and health of these plugins.

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