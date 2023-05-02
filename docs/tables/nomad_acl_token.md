# Table: nomad_acl_token

Nomad uses tokens to authenticate requests to the cluster. These tokens are created using the `nomad acl token create` command. When a token is being created, the operator can specify one or more policies to apply to the token. These policies determine if any action specified by the token bearer is authorized.

## Examples

### Basic info

```sql
select
  name,
  accessor_id,
  secret_id,
  type,
  global,
  create_time,
  expiration_ttl
from
  nomad_acl_token;
```

### List management tokens

```sql
select
  name,
  accessor_id,
  secret_id,
  type,
  global,
  create_time,
  expiration_ttl
from
  nomad_acl_token
where
  type = 'management';
```

### List global tokens

```sql
select
  name,
  accessor_id,
  secret_id,
  type,
  global,
  create_time,
  expiration_ttl
from
  nomad_acl_token
where
  global;
```

### List tokens which will not expire

```sql
select
  name,
  accessor_id,
  secret_id,
  type,
  global,
  create_time,
  expiration_ttl
from
  nomad_acl_token
where
  expiration_time is null;
```

### List tokens which are not associated with any role

```sql
select
  name,
  accessor_id,
  secret_id,
  type,
  global,
  create_time,
  expiration_ttl
from
  nomad_acl_token
where
  roles is null;
```
