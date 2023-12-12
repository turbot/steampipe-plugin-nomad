---
title: "Steampipe Table: nomad_namespace - Query Nomad Namespaces using SQL"
description: "Allows users to query Nomad Namespaces, specifically the details and configurations of each namespace, providing insights into the structure and organization of tasks within the Nomad orchestration system."
---

# Table: nomad_namespace - Query Nomad Namespaces using SQL

Nomad Namespaces is a feature of HashiCorp's Nomad, an orchestration system that enables the organization and segmentation of tasks. It allows for the isolation of tasks and resources, providing a way to separate teams or environments within a single Nomad cluster. This feature enhances the efficiency and security of task management in the system.

## Table Usage Guide

The `nomad_namespace` table provides insights into namespaces within HashiCorp's Nomad orchestration system. As a DevOps engineer, explore namespace-specific details through this table, including their configurations and associated metadata. Utilize it to understand the structure and organization of tasks within your Nomad system, and to enhance the efficiency and security of your task management.

**Important Notes**
- You need to specify the `namespace` config argument in the `nomad.spc` file to be able to query this table.

## Examples

### Basic info
Explore the basic information of your namespaces in Nomad to gain insights into their creation and modification indices. This can be useful for tracking changes and understanding the lifecycle of your namespaces.

```sql+postgres
select
  name,
  description,
  create_index,
  modify_index
from
  nomad_namespace;
```

```sql+sqlite
select
  name,
  description,
  create_index,
  modify_index
from
  nomad_namespace;
```

### List the disabled task drivers of namespaces
Uncover the details of task drivers that are inactive within various namespaces. This allows for an assessment of system capabilities and potential areas of improvement.

```sql+postgres
select
  name,
  description,
  capabilities -> 'DisabledTaskDrivers' as disabled_task_drivers
from
  nomad_namespace;
```

```sql+sqlite
select
  name,
  description,
  json_extract(capabilities, '$.DisabledTaskDrivers') as disabled_task_drivers
from
  nomad_namespace;
```