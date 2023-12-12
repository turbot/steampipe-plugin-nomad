---
title: "Steampipe Table: nomad_acl_token - Query Nomad ACL Tokens using SQL"
description: "Allows users to query Nomad ACL Tokens, providing detailed information about each token's access control list (ACL) policies, type, and other related metadata."
---

# Table: nomad_acl_token - Query Nomad ACL Tokens using SQL

Nomad ACL Tokens are used in HashiCorp Nomad to provide a flexible, capability-based access control system. They allow operators to restrict access to certain data and APIs, making Nomad more secure. Each ACL token is associated with a set of policies, and these policies dictate the token's specific capabilities.

## Table Usage Guide

The `nomad_acl_token` table provides insights into ACL Tokens within HashiCorp Nomad. As a security analyst, explore token-specific details through this table, including associated policies, token type, and related metadata. Utilize it to uncover information about tokens, such as those with broad permissions, and to verify the security of your Nomad deployment.

**Important Notes**
- You need to specify the `secret_id` config argument in the `nomad.spc` file to be able to query this table.

## Examples

### Basic info
Assess the elements within your Nomad ACL tokens to understand the different types, their global status, creation time, and expiration time-to-live (TTL). This can help manage and track the lifecycle and accessibility of each token.

```sql+postgres
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

```sql+sqlite
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
Explore which management tokens are currently active in your system to understand their creation and expiration timelines. This can be beneficial for assessing your system's security by identifying potential vulnerabilities due to outdated or globally accessible tokens.

```sql+postgres
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

```sql+sqlite
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
Explore which access control list (ACL) tokens in Nomad are set as global. This is useful in identifying potential security risks associated with globally accessible tokens.

```sql+postgres
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

```sql+sqlite
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
  global = 1;
```

### List tokens which will never expire
Uncover the details of access control list (ACL) tokens that have been set without an expiration time, thus identifying potential security risks due to tokens that will never expire. This is useful for maintaining secure practices by ensuring all tokens have a designated expiry.

```sql+postgres
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

```sql+sqlite
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
Discover the segments that consist of tokens which are not linked to any role. This is useful for identifying potential security risks, as these tokens may have been created without proper role assignments.

```sql+postgres
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

```sql+sqlite
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