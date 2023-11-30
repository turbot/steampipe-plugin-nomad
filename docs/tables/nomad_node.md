---
title: "Steampipe Table: nomad_node - Query Nomad Nodes using SQL"
description: "Allows users to query Nomad Nodes, specifically the information about each node registered with a Nomad cluster, providing insights into node status, resources, and metadata."
---

# Table: nomad_node - Query Nomad Nodes using SQL

A Nomad Node is a physical or virtual machine that has the Nomad agent running on it. It is responsible for running tasks and reporting on their status. Nodes are registered with a Nomad cluster and are managed by the server agents.

## Table Usage Guide

The `nomad_node` table provides insights into nodes within HashiCorp Nomad. As a DevOps engineer, explore node-specific details through this table, including status, resources, and associated metadata. Utilize it to uncover information about nodes, such as their availability, resource utilization, and the tasks they are running.

## Examples

### Basic info
Explore the status and configuration of nodes within your network to understand their performance and identify any potential issues. This can help in maintaining optimal network performance and preemptively addressing any problems.

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

### List nodes with drain enabled
Uncover the details of nodes that have the drain feature enabled. This is particularly useful for identifying nodes that are temporarily not accepting any new allocations, which can aid in resource management and troubleshooting.

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
Determine the areas in which nodes are not yet ready for operation. This is beneficial for identifying potential issues in your network and addressing them proactively to avoid disruptions.

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
Discover the segments that have Transport Layer Security (TLS) disabled in your network. This is essential to identify potential security vulnerabilities and ensure all nodes in your network are secure.

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
Discover the segments that are eligible for scheduling in a data center, allowing you to better manage resource allocation and workload distribution. This can help you optimize performance and ensure the smooth operation of your systems.

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