# Table: nomad_volume

Volumes are named, persistent data that resides in the host in the path /var/lib/docker/volumes. When the job is run, Docker will create the volume if it does not exist, or mount it. It’s the same as interacting with docker volume create.

## Examples

### Basic info

```sql
select
  id,
  name,
  clone_id,
  capacity,
  namespace,
  provider,
  plugin_id
from
  nomad_volume;
```

### List all volumes with the access mode set to "ReadWriteOnce"

```sql
select
  id,
  name,
  clone_id,
  capacity,
  namespace,
  provider,
  plugin_id
from
  nomad_volume
where
  access_mode ->> 'mode' = 'ReadWriteOnce';
```

### List all volumes with at least one healthy controller and one healthy node

```sql
select
  id,
  name,
  clone_id,
  capacity,
  namespace,
  provider,
  plugin_id
from
  nomad_volume
where
  controllers_healthy > 0
  and nodes_healthy > 0;
```

### List all volumes with the schedulable flag set to true

```sql
select
  id,
  name,
  clone_id,
  capacity,
  namespace,
  provider,
  plugin_id
from
  nomad_volume
where
  schedulable = true;
```

### List all CSI volumes with a requested capacity between 50GB and 200GB

```sql
select
  id,
  name,
  clone_id,
  capacity,
  namespace,
  provider,
  plugin_id
from
  nomad_volume
where
  requested_capacity_min >= 50000000000
  and requested_capacity_max <= 200000000000;
```

### List all CSI volumes with a controller_required flag set to true and at least 3 healthy controllers

```sql
select
  id,
  name,
  clone_id,
  capacity,
  namespace,
  provider,
  plugin_id
from
  nomad_volume
where
  controller_required = true
  and controllers_healthy >= 3;
```

