# Table: nomad_acl_policy

Policies consist of a set of rules defining the capabilities or actions to be granted. For example, a readonly policy might only grant the ability to list and inspect running jobs, but not to submit new ones. No permissions are granted by default, making Nomad a default-deny system.

## Examples

### Basic info

```sql
select
  name,
  rules,
  title,
  description,
  create_index,
  modify_index
from
  nomad_acl_policy;
```

### List the policies that are attached to a job

```sql
select
  name,
  rules,
  title,
  description,
  jsonb_pretty(job_acl)
from
  nomad_acl_policy
where
  job_acl is not null;
```