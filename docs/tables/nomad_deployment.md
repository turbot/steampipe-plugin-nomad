---
title: "Steampipe Table: nomad_deployment - Query Nomad Deployments using SQL"
description: "Allows users to query Nomad Deployments, specifically to obtain information about the status, configuration, and progress of deployments."
---

# Table: nomad_deployment - Query Nomad Deployments using SQL

Nomad is a flexible workload orchestrator designed to deploy, manage, and scale containerized and non-containerized applications. Deployments in Nomad represent a series of actions to be taken on a particular job version to transition it from a previous job version. It provides details about the status, configuration, and progress of a deployment.

## Table Usage Guide

The `nomad_deployment` table provides insights into deployments within Nomad. As a DevOps engineer, explore deployment-specific details through this table, including status, configuration, and progress. Utilize it to uncover information about deployments, such as their current state, the job version they are associated with, and the tasks involved in the deployment.

## Examples

### Basic info
Explore the status and version details of various jobs within a deployment to assess their performance and identify any potential issues. This can be useful for maintaining efficiency and troubleshooting in multi-region deployments.

```sql+postgres
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

```sql+sqlite
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
Explore which deployments are not currently active. This can be useful in identifying potential issues or inefficiencies within your system.

```sql+postgres
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

```sql+sqlite
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
  status != 'running';
```

### List multi region deployments
Discover the segments that have multiple region deployments to better understand and manage your distributed resources. This can help in identifying potential areas for optimization and risk mitigation.

```sql+postgres
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

```sql+sqlite
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
Explore the status and configuration of deployments, including whether they span multiple regions. This can be useful for understanding the complexity and reach of your deployments.

```sql+postgres
select
  id,
  namespace,
  status,
  is_multiregion,
  jsonb_pretty(task_groups) as task_groups
from
  nomad_deployment;
```

```sql+sqlite
select
  id,
  namespace,
  status,
  is_multiregion,
  task_groups
from
  nomad_deployment;
```