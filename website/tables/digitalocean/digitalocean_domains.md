# Table: digitalocean_domains

This table shows data for DigitalOcean Domains.

The primary key for this table is **name**.

## Relations

The following tables depend on digitalocean_domains:
  - [digitalocean_domain_records](digitalocean_domain_records)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|name (PK)|String|
|ttl|Int|
|zone_file|String|