---
title: "Steampipe Table: nomad_acl_binding_rule - Query Nomad ACL Binding Rules using SQL"
description: "Allows users to query Nomad ACL Binding Rules, specifically the details of each rule, facilitating the management of access control lists in Nomad."
---

# Table: nomad_acl_binding_rule - Query Nomad ACL Binding Rules using SQL

Nomad ACL Binding Rule is a resource in HashiCorp's Nomad that allows users to map identities from an external trusted source to Nomad ACL Policies. It provides a way to define how trusted identities from an external source are granted Nomad ACL policies. Nomad ACL Binding Rule helps you manage access control lists by defining the mapping between the trusted identity attributes and Nomad ACL policies.

## Table Usage Guide

The `nomad_acl_binding_rule` table provides insights into ACL Binding Rules within Nomad. As a DevOps engineer or a security analyst, you can explore rule-specific details through this table, including the associated policies and identity sources. Utilize it to manage and monitor access control lists in your Nomad environment, ensuring secure and efficient operation.

**Important Notes**
- You need to specify the `secret_id` config argument in the `nomad.spc` file to be able to query this table.

## Examples

### Basic info
Explore the different authorization methods and types used within your system, along with their creation details. This information can help you understand your access control list (ACL) binding rules, making it easier to manage user permissions and security.

```sql+postgres
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

```sql+sqlite
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
Discover the specific binding rules that are associated with 'role' types. This can be particularly useful in managing access control lists (ACLs) within the Nomad system.

```sql+postgres
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

```sql+sqlite
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
Discover the recently created binding rules in your Nomad ACL to stay updated on changes and modifications. This helps you maintain an oversight of your authorization methods, ensuring they are up-to-date and secure.

```sql+postgres
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

```sql+sqlite
select
  id,
  auth_method,
  bind_name,
  bind_type,
  create_time
from
  nomad_acl_binding_rule
where
  create_time >= datetime('now', '-30 days');
```

### Show auth methods related to the binding rule
Explore the relationship between authentication methods and binding rules to better understand security protocols. This can help identify potential vulnerabilities or areas for improvement in your system's security.

```sql+postgres
select
  a.name as auth_method_name,
  a.type as auth_method_type,
  a.create_time as auth_method_creation_time,
  b.id as binding_rule_id,
  b.bind_type as bind_type
from
  nomad_acl_binding_rule as b
  left join nomad_acl_auth_method as a on b.auth_method = a.name;
```

```sql+sqlite
select
  a.name as auth_method_name,
  a.type as auth_method_type,
  a.create_time as auth_method_creation_time,
  b.id as binding_rule_id,
  b.bind_type as bind_type
from
  nomad_acl_binding_rule as b
  left join nomad_acl_auth_method as a on b.auth_method = a.name;
```