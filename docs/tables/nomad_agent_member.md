# Table: nomad_agent_member

The Nomad agent is a long running process which runs on every machine that is part of the Nomad cluster. The behavior of the agent depends on if it is running in client or server mode. Clients are responsible for running tasks, while servers are responsible for managing the cluster.

## Examples

### Basic info

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
