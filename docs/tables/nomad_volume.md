---
title: "Steampipe Table: nomad_volume - Query Nomad Volumes using SQL"
description: "Allows users to query Nomad Volumes, specifically retrieving details about volume configuration and usage. This provides insights into the management of storage resources within a Nomad cluster."
---

# Table: nomad_volume - Query Nomad Volumes using SQL

Nomad is a workload orchestrator that enables the deployment and management of applications across a fleet of servers, providing features like service discovery, global deployments, and horizontal scalability. One of its resources is the Nomad Volume, which is a logical unit of storage provisioned from a storage provider in the cluster. The Nomad Volume holds data that can be accessed by applications running within the Nomad cluster.

## Table Usage Guide

The `nomad_volume` table provides insights into the volumes within the Nomad orchestration system. As a systems administrator or DevOps engineer, you can explore volume-specific details through this table, including configuration parameters, usage statistics, and associated metadata. Utilize it to uncover information about volume allocation, such as which tasks are using a volume, the current state of the volume, and the verification of volume configurations.

## Examples

### Basic info
Explore the fundamental details of your storage volumes, such as their identification, capacity, and associated provider, to better understand and manage your resources. This information can be particularly useful for identifying capacity issues or for tracking volumes across different providers.

```sql
select
  id,
  name,
  clone_id,
  capacity,
  namespace,
  provider,
  plugin_id
from
  nomad_volume;
```

### List volumes with access mode set to `ReadWriteOnce`
Explore which volumes have their access mode set to 'ReadWriteOnce'. This can be useful for managing data accessibility and ensuring specific volumes are not simultaneously written by multiple users.

```sql
select
  id,
  name,
  clone_id,
  capacity,
  namespace,
  provider,
  plugin_id
from
  nomad_volume
where
  access_mode ->> 'mode' = 'ReadWriteOnce';
```

### List volumes with at least one healthy node
Explore which volumes in your system are functioning optimally by identifying those with at least one healthy node. This can be beneficial for system maintenance and troubleshooting.

```sql
select
  id,
  name,
  clone_id,
  capacity,
  namespace,
  provider,
  plugin_id
from
  nomad_volume
where
  nodes_healthy > 0;
```

### List schedulable volumes
Discover the volumes that can be scheduled for tasks, helping in managing and optimizing resource allocation. This can be particularly useful in understanding and maximizing the utilization of your storage resources.

```sql
select
  id,
  name,
  clone_id,
  capacity,
  namespace,
  provider,
  plugin_id
from
  nomad_volume
where
  schedulable;
```

### List volumes present in default namespace
Explore the storage volumes present within the default operational space. This is useful for understanding the current storage usage and capacity management within your default namespace.

```sql
select
  id,
  name,
  clone_id,
  capacity,
  namespace,
  provider,
  plugin_id
from
  nomad_volume
where
  namespace = 'default';
```