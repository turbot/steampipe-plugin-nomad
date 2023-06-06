# Table: nomad_acl_auth_method

Authentication methods dictate how Nomad should talk to SSO providers when a user requests to authenticate using one. Currently, Nomad only supports the OpenID Connect (OIDC) SSO workflow which allows users to log in to Nomad via applications such as Auth0, Okta, and Vault.

You need to specify the `secret_id` parameter in the `nomad.spc` file to be able to query this table.

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

### List default auth methods

```sql
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

### List auth methods created in the last 30 days

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

### List auth methods with global token locality

```sql
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

```sql
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
