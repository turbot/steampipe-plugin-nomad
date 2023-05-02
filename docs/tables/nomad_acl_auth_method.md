# Table: nomad_acl_auth_method

Authentication methods dictate how Nomad should talk to SSO providers when a user requests to authenticate using one. Currently, Nomad only supports the OpenID Connect (OIDC) SSO workflow which allows users to log in to Nomad via applications such as Auth0, Okta, and Vault.

## Examples

### Basic info

```sql
select
  name,
  title,
  type,
  create_time,
  create_index,
  modify_index,
  modify_time
from
  nomad_acl_auth_method;
```

### List the default auth methods

```sql
select
  name,
  title,
  type,
  create_time
from
  nomad_acl_auth_method
where
  "default";
```

### List the auth methods created in the last 30 days

```sql
select
  name,
  title,
  type,
  create_time
from
  nomad_acl_auth_method
where
  create_time >= now() - interval '30' day;
```

### List of access tokens which haven't been modified in the last 30 days

```sql
select
  name,
  title,
  type,
  create_time
from
  nomad_acl_auth_method
where
  modify_time <= now() - interval '30' day;
```

### List all the auth methods that create a global token with SSO login

```sql
select
  name,
  title,
  type,
  create_time
from
  nomad_acl_auth_method
where
  token_locality like '%global%';
  ```