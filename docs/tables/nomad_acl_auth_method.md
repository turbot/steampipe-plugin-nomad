---
title: "Steampipe Table: nomad_acl_auth_method - Query Nomad ACL Auth Methods using SQL"
description: "Allows users to query Nomad ACL Auth Methods, specifically the information about the ACL Auth Methods configured in Nomad."
---

# Table: nomad_acl_auth_method - Query Nomad ACL Auth Methods using SQL

Nomad ACL Auth Method is a feature within HashiCorp Nomad that provides a way to control access to resources and operations in a Nomad cluster. It allows operators to define a set of rules that can be used to restrict the actions that a user or group of users can perform. These rules can be used to create fine-grained access control policies that are tailored to the specific needs of your organization.

## Table Usage Guide

The `nomad_acl_auth_method` table offers insights into the ACL Auth Methods configured within HashiCorp Nomad. As a system administrator or DevOps engineer, leverage this table to understand the access control policies in place, including the rules and permissions associated with each method. This table can be instrumental in auditing your Nomad cluster's security configuration and ensuring that access controls are appropriately set.

**Important Notes**
- You need to specify the `secret_id` config argument in the `nomad.spc` file to be able to query this table.

## Examples

### Basic info
Explore the creation and modification details of various access control methods in Nomad. This is useful for understanding the evolution and changes in your security settings over time.

```sql+postgres
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

```sql+sqlite
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

### List default auth methods
Explore which authentication methods are set as default in your system. This can help in understanding the primary security measures in place and when they were established.

```sql+postgres
select
  name,
  title,
  type,
  create_time
from
  nomad_acl_auth_method
where
  default_auth_method;
```

```sql+sqlite
select
  name,
  title,
  type,
  create_time
from
  nomad_acl_auth_method
where
  default_auth_method = 1;
```

### List auth methods created in the last 30 days
Discover the segments that have been authorized in the last month. This could be used to monitor recent changes in access permissions, helping to maintain system security.

```sql+postgres
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

```sql+sqlite
select
  name,
  title,
  type,
  create_time
from
  nomad_acl_auth_method
where
  create_time >= datetime('now','-30 day');
```

### List auth methods with global token locality
Explore which authentication methods have a global token locality. This is useful to understand which methods can be applied universally across your network.

```sql+postgres
select
  name,
  title,
  type,
  create_time
from
  nomad_acl_auth_method
where
  token_locality = 'global';
```

```sql+sqlite
select
  name,
  title,
  type,
  create_time
from
  nomad_acl_auth_method
where
  token_locality = 'global';
```

### Get config details of a particular auth method
Explore the configuration details of a specific authentication method to understand its settings and parameters. This is useful for auditing security settings or troubleshooting authentication issues.

```sql+postgres
select
  name,
  config ->> 'JWKSURL' as "jwks_url",
  config ->> 'JWKSCACert' as "jwks_ca_cert",
  config -> 'OIDCScopes' as "oidc_scopes",
  config -> 'BoundIssuer' as "bound_issuer",
  config -> 'SigningAlgs' as "signing_algs",
  config ->> 'OIDCClientID' as "oidc_client_id",
  config -> 'ClaimMappings' as "claim_mappings",
  config -> 'BoundAudiences' as "bound_audiences",
  config -> 'DiscoveryCaPem' as "discovery_ca_pem",
  config ->> 'ClockSkewLeeway' as "clock_skew_leeway",
  config ->> 'NotBeforeLeeway' as "not_before_leeway",
  config ->> 'ExpirationLeeway' as "expiration_leeway",
  config ->> 'OIDCClientSecret' as "oidc_client_secret",
  config ->> 'OIDCDiscoveryURL' as "oidc_discovery_url",
  config -> 'ListClaimMappings' as "list_claim_mappings",
  config -> 'AllowedRedirectURIs' as "allowed_redirect_uris",
  config -> 'JWTValidationPubKeys' as "jwt_validation_pub_keys"
from
  nomad_acl_auth_method
where
  name = 'auth-method';
```

```sql+sqlite
select
  name,
  json_extract(config, '$.JWKSURL') as "jwks_url",
  json_extract(config, '$.JWKSCACert') as "jwks_ca_cert",
  json_extract(config, '$.OIDCScopes') as "oidc_scopes",
  json_extract(config, '$.BoundIssuer') as "bound_issuer",
  json_extract(config, '$.SigningAlgs') as "signing_algs",
  json_extract(config, '$.OIDCClientID') as "oidc_client_id",
  json_extract(config, '$.ClaimMappings') as "claim_mappings",
  json_extract(config, '$.BoundAudiences') as "bound_audiences",
  json_extract(config, '$.DiscoveryCaPem') as "discovery_ca_pem",
  json_extract(config, '$.ClockSkewLeeway') as "clock_skew_leeway",
  json_extract(config, '$.NotBeforeLeeway') as "not_before_leeway",
  json_extract(config, '$.ExpirationLeeway') as "expiration_leeway",
  json_extract(config, '$.OIDCClientSecret') as "oidc_client_secret",
  json_extract(config, '$.OIDCDiscoveryURL') as "oidc_discovery_url",
  json_extract(config, '$.ListClaimMappings') as "list_claim_mappings",
  json_extract(config, '$.AllowedRedirectURIs') as "allowed_redirect_uris",
  json_extract(config, '$.JWTValidationPubKeys') as "jwt_validation_pub_keys"
from
  nomad_acl_auth_method
where
  name = 'auth-method';
```