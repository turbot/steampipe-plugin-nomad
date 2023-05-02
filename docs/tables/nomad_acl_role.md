# Table: nomad_acl_role

Roles group one or more ACL policies into a container which can then be used to generate ACL tokens for authorization. This abstraction allows easier control and updating of ACL permissions, particularly in larger, more diverse clusters.

## Examples

### Basic info

```sql
select
  id,
  name,
  description,
  create_index,
  modify_index
from
  nomad_acl_role;
```

### Show ACL policies attached to a particular ACL role

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
      p ->> 'Name'
    from
      nomad_acl_role,
      jsonb_array_elements(policies) as p
    where
      name = 'aclRole'
  );
```

### List roles which are attached to ACL tokens

```sql
select
  id,
  name,
  description,
  create_index,
  modify_index
from
  nomad_acl_role
where
  name in
  (
    select
      jsonb_array_elements_text(roles)
    from
      nomad_acl_token
  );
```
