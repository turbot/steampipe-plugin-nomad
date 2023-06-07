# Table: nomad_acl_policy

Policies consist of a set of rules defining the capabilities or actions to be granted. For example, a readonly policy might only grant the ability to list and inspect running jobs, but not to submit new ones. No permissions are granted by default, making Nomad a default-deny system.

You need to specify the `secret_id` config argument in the `nomad.spc` file to be able to query this table.

## Examples

### Basic info

```sql
select
  name,
  rules,
  description,
  create_index,
  modify_index
from
  nomad_acl_policy;
```

### List policies that are attached to any job

```sql
select
  name,
  rules,
  description,
  create_index,
  modify_index
from
  nomad_acl_policy
where
  job_acl is not null;
```

### List policies which are attached to ACL tokens

```sql
select
  name,
  rules,
  description,
  create_index,
  modify_index
from
  nomad_acl_policy
where
  name in
  (
    select
      jsonb_array_elements_text(policies)
    from
      nomad_acl_token
  );
```