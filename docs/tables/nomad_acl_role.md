---
title: "Steampipe Table: nomad_acl_role - Query Nomad ACL Roles using SQL"
description: "Allows users to query ACL Roles in Nomad, specifically the role's name, type, policies, and related metadata, providing insights into access control configurations and potential security risks."
---

# Table: nomad_acl_role - Query Nomad ACL Roles using SQL

Nomad ACL Roles are a resource in HashiCorp's Nomad that allows you to define permissions for clients and servers. They provide a flexible way to manage access control, allowing you to specify which actions a client or server can perform and on which resources. ACL Roles are an integral part of Nomad's security model, which aims to provide secure, multi-tenant environments.

## Table Usage Guide

The `nomad_acl_role` table provides insights into ACL roles within HashiCorp's Nomad. As a DevOps engineer, explore role-specific details through this table, including role names, types, and associated policies. Utilize it to uncover information about roles, such as their permissions, the resources they have access to, and potential security risks in your Nomad environment.

**Important Notes**
- You need to specify the `secret_id` config argument in the `nomad.spc` file to be able to query this table.

## Examples

### Basic info
Explore which access control roles have been created and modified in your Nomad environment, allowing you to maintain security and manage user permissions effectively.

```sql+postgres
select
  id,
  name,
  description,
  create_index,
  modify_index
from
  nomad_acl_role;
```

```sql+sqlite
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
Identify the access control policies linked to a specific role to understand its permissions and restrictions. This could be useful in auditing or updating security measures.

```sql+postgres
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

```sql+sqlite
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
      json_extract(p.value, '$.Name')
    from
      nomad_acl_role,
      json_each(policies) as p
    where
      name = 'aclRole'
  );
```

### List roles which are attached to ACL tokens
Determine the roles associated with ACL tokens in your system to understand their permissions and access levels. This can be useful in managing security and ensuring proper access control within your network.

```sql+postgres
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

```sql+sqlite
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
      json_each.value
    from
      nomad_acl_token,
      json_each(roles)
  );
```