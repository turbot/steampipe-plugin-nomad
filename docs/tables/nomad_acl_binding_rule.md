# Table: nomad_acl_binding_rule

Binding rules provide a mapping between a Nomad user's SSO authorization claims and internal Nomad objects such as ACL Roles and ACL Policies. A binding rule is directly related to a single auth method, and therefore only evaluated by login attempts using that method. All binding rules mapped to an auth method are evaluated during each login attempt.

## Examples

### Basic info

```sql
select
  id,
  auth_method,
  bind_name,
  bind_type,
  create_time,
  create_index
from
  nomad_acl_binding_rule;
```

### List role type binding rules

```sql
select
  id,
  auth_method,
  bind_name,
  bind_type,
  create_time,
  create_index
from
  nomad_acl_binding_rule
where
  bind_type = 'role';
```

### List binding rules created in the last 30 days

```sql
select
  id,
  auth_method,
  bind_name,
  bind_type,
  create_time
from
  nomad_acl_binding_rule
where
  create_time >= now() - interval '30' day;
```

### Show auth methods related to the binding rule

```sql
select
  a.name as auth_method_name,
  a.type as auth_method_type,
  a.create_time as auth_method_creation_time,
  b.id as binding_rule_id,
  b.bind_type as bind_type
from
  nomad_acl_binding_rule as b
  left join
    nomad_acl_auth_method as a
    on b.auth_method = a.name;
```
