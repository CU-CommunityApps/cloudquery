# Table: aws_ecs_cluster_task_sets

This table shows data for Amazon Elastic Container Service (ECS) Cluster Task Sets.

https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_TaskSet.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_ecs_cluster_services](aws_ecs_cluster_services).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|capacity_provider_strategy|JSON|
|cluster_arn|String|
|computed_desired_count|Int|
|created_at|Timestamp|
|external_id|String|
|id|String|
|launch_type|String|
|load_balancers|JSON|
|network_configuration|JSON|
|pending_count|Int|
|platform_family|String|
|platform_version|String|
|running_count|Int|
|scale|JSON|
|service_arn|String|
|service_registries|JSON|
|stability_status|String|
|stability_status_at|Timestamp|
|started_by|String|
|status|String|
|task_definition|String|
|task_set_arn|String|
|updated_at|Timestamp|