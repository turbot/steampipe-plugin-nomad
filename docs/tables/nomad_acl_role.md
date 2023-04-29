# Table: nomad_acl_role

Roles group one or more ACL policies into a container which can then be used to generate ACL tokens for authorisation. This abstraction allows easier control and updating of ACL permissions, particularly in larger, more diverse clusters.

## Examples

### Basic info

```sql
select
  id,
  name,
  title,
  description,
  create_index,
  modify_index
from
  nomad_acl_role;
```

### Name the roles attached to an ACL policy

```sql
select
  id,
  name,
  title,
  description,
  create_index,
  modify_index
from
  nomad_acl_role
where
  policies is not null;
```