---
title: "Steampipe Table: nomad_plugin - Query Nomad Plugins using SQL"
description: "Allows users to query Plugins in Nomad, specifically providing details on plugin configuration, status, and driver details."
---

# Table: nomad_plugin - Query Nomad Plugins using SQL

A Nomad Plugin is a crucial component in the Nomad ecosystem, enabling extensibility and customization. Plugins are leveraged to integrate with various external systems, providing a wide range of functionality, from task drivers to device plugins and log storage. It is an essential tool for managing and extending the capabilities of your Nomad environment.

## Table Usage Guide

The `nomad_plugin` table provides insights into Plugins within Nomad. As a DevOps engineer, explore plugin-specific details through this table, including its configuration, status, and driver details. Utilize it to uncover information about plugins, such as their current state, associated tasks, and the overall health of the plugin.

## Examples

### Basic info
Explore which plugins are required for your controller by determining their version and provider. This aids in assessing the elements within your system for better management and updates.

```sql
select
  id,
  controller_required,
  version,
  create_index,
  modify_index,
  provider
from
  nomad_plugin;
```

### List CSI plugins that require a controller
Analyze the settings to understand which CSI plugins necessitate a controller. This allows you to pinpoint the specific locations where controllers are required, ensuring your system configuration is optimized.

```sql
select
  id,
  title,
  version,
  create_index,
  modify_index,
  provider
from
  nomad_plugin
where
  controller_required;
```

### Show the number of nodes and controllers associated with each plugin
Discover the health status of your system by identifying the number of healthy nodes and controllers for each plugin. This information can be instrumental in assessing system performance and identifying potential areas for improvement.

```sql
select
  id,
  version,
  nodes_healthy,
  controllers_healthy
from
  nomad_plugin;
```

### Show the number of expected nodes for each plugin
Assess the elements within each plugin to understand the expected number of nodes. This information can be useful in planning and allocating resources effectively.

```sql
select
  id,
  version,
  nodes_healthy,
  nodes_expected
from
  nomad_plugin;
```