---
title: "Steampipe Table: nomad_agent_member - Query Nomad Agent Members using SQL"
description: "Allows users to query Nomad Agent Members, specifically providing details about the members of the Nomad agent, such as their role, status, and protocol versions."
---

# Table: nomad_agent_member - Query Nomad Agent Members using SQL

Nomad is a simple and flexible workload orchestrator to deploy and manage containers and non-containerized applications across on-prem and clouds at scale. Nomad is widely adopted and used in production by PagerDuty, Trivago, Deluxe Entertainment, and more. An agent member in Nomad represents a single instance of the Nomad agent.

## Table Usage Guide

The `nomad_agent_member` table provides insights into the members of a Nomad agent. As a DevOps engineer or system administrator, explore member-specific details through this table, including their role, status, and protocol versions. Utilize it to uncover information about members, such as their status, roles, and the protocol versions they are using.

## Examples

### Basic info
Explore the status and details of various agents within the Nomad system. This can be useful in monitoring network health and identifying potential issues with specific agents.

```sql
select
  name,
  address,
  port,
  status,
  protocol_cur
from
  nomad_agent_member;
```

### List global agents
Explore which global agents are currently active in your network. This can help you manage your resources more efficiently and identify potential issues with network connectivity or performance.

```sql
select
  name,
  address,
  port,
  status,
  protocol_cur
from
  nomad_agent_member
where
  tags ->> 'region' = 'global';
```

### List agents which are not `alive`
Discover the details of agents that are not currently active. This query can be used to identify potential issues or disruptions in your network by pinpointing inactive agents.

```sql
select
  name,
  address,
  port,
  status,
  protocol_cur
from
  nomad_agent_member
where
  status <> 'alive';
```