---
title: "Steampipe Table: nomad_acl_policy - Query Nomad ACL Policies using SQL"
description: "Allows users to query Nomad ACL Policies, specifically retrieving details about the policy name, description, rules, and creation time."
---

# Table: nomad_acl_policy - Query Nomad ACL Policies using SQL

Nomad Access Control List (ACL) Policies are a crucial aspect of Nomad's security model, enabling fine-grained authorization of Nomad's API. These policies provide a flexible way to grant permissions to Nomad's API, allowing for a variety of permissions based on the specific needs of your applications. ACL Policies are defined in HashiCorp Configuration Language (HCL) or JSON, and can be managed via the API or CLI.

## Table Usage Guide

The `nomad_acl_policy` table provides insights into ACL policies within HashiCorp Nomad. As a security analyst, explore policy-specific details through this table, including policy names, descriptions, rules, and creation times. Utilize it to uncover information about policies, such as their specific permissions and the resources they apply to, providing a comprehensive view of your Nomad environment's security posture.

**Important Notes**
- You need to specify the `secret_id` config argument in the `nomad.spc` file to be able to query this table.

## Examples

### Basic info
Explore the policies in your Nomad cluster to understand their rules and descriptions, as well as when they were created or last modified. This could be useful for auditing purposes or to ensure compliance with security protocols.

```sql+postgres
select
  name,
  rules,
  description,
  create_index,
  modify_index
from
  nomad_acl_policy;
```

```sql+sqlite
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
Explore which policies are linked to a job to gain insights into their rules and descriptions, useful for understanding the permissions and restrictions associated with different jobs. This can help in effectively managing and modifying job-related policies.

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
  job_acl is not null;
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
  job_acl is not null;
```

### List policies which are attached to ACL tokens
Explore which policies are related to ACL tokens, allowing you to understand the rules and descriptions associated with each policy. This can help in managing and modifying your ACL tokens more effectively.

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
      jsonb_array_elements_text(policies)
    from
      nomad_acl_token
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
      json_each.value
    from
      nomad_acl_token,
      json_each(policies)
  );
```