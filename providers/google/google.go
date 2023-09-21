package aws

import (
	"fmt"

	"github.com/tf2d2/icons/internal/resource"
)

var (
	Resources = []*resource.Resource{
		{
			Name: "google_access_context_manager_access_level",
		},
		{
			Name: "google_access_context_manager_access_level_condition",
		},
		{
			Name: "google_access_context_manager_access_levels",
		},
		{
			Name: "google_access_context_manager_access_policy",
		},
		{
			Name: "google_access_context_manager_access_policy_iam_binding",
		},
		{
			Name: "google_access_context_manager_access_policy_iam_member",
		},
		{
			Name: "google_access_context_manager_access_policy_iam_policy",
		},
		{
			Name: "google_access_context_manager_authorized_orgs_desc",
		},
		{
			Name: "google_access_context_manager_egress_policy",
		},
		{
			Name: "google_access_context_manager_gcp_user_access_binding",
		},
		{
			Name: "google_access_context_manager_ingress_policy",
		},
		{
			Name: "google_access_context_manager_service_perimeter",
		},
		{
			Name: "google_access_context_manager_service_perimeter_egress_policy",
		},
		{
			Name: "google_access_context_manager_service_perimeter_ingress_policy",
		},
		{
			Name: "google_access_context_manager_service_perimeter_resource",
		},
		{
			Name: "google_access_context_manager_service_perimeters",
		},
		{
			Name: "google_active_directory_domain",
		},
		{
			Name: "google_active_directory_domain_trust",
		},
		{
			Name: "google_alloydb_backup",
		},
		{
			Name: "google_alloydb_cluster",
		},
		{
			Name: "google_alloydb_instance",
		},
		{
			Name: "google_apigee_addons_config",
		},
		{
			Name: "google_apigee_endpoint_attachment",
		},
		{
			Name: "google_apigee_env_keystore",
		},
		{
			Name: "google_apigee_env_references",
		},
		{
			Name: "google_apigee_envgroup",
		},
		{
			Name: "google_apigee_envgroup_attachment",
		},
		{
			Name: "google_apigee_environment",
		},
		{
			Name: "google_apigee_environment_iam_binding",
		},
		{
			Name: "google_apigee_environment_iam_member",
		},
		{
			Name: "google_apigee_environment_iam_policy",
		},
		{
			Name: "google_apigee_flowhook",
		},
		{
			Name: "google_apigee_instance",
		},
		{
			Name: "google_apigee_instance_attachment",
		},
		{
			Name: "google_apigee_keystores_aliases_key_cert_file",
		},
		{
			Name: "google_apigee_keystores_aliases_pkcs12",
		},
		{
			Name: "google_apigee_keystores_aliases_self_signed_cert",
		},
		{
			Name: "google_apigee_nat_address",
		},
		{
			Name: "google_apigee_organization",
		},
		{
			Name: "google_apigee_sharedflow",
		},
		{
			Name: "google_apigee_sharedflow_deployment",
		},
		{
			Name: "google_apigee_sync_authorization",
		},
		{
			Name: "google_apigee_target_server",
		},
		{
			Name: "google_apikeys_key",
		},
		{
			Name: "google_app_engine_application",
		},
		{
			Name: "google_app_engine_application_url_dispatch_rules",
		},
		{
			Name: "google_app_engine_domain_mapping",
		},
		{
			Name: "google_app_engine_firewall_rule",
		},
		{
			Name: "google_app_engine_flexible_app_version",
		},
		{
			Name: "google_app_engine_service_network_settings",
		},
		{
			Name: "google_app_engine_service_split_traffic",
		},
		{
			Name: "google_app_engine_standard_app_version",
		},
		{
			Name: "google_artifact_registry_repository",
		},
		{
			Name: "google_artifact_registry_repository_iam_binding",
		},
		{
			Name: "google_artifact_registry_repository_iam_member",
		},
		{
			Name: "google_artifact_registry_repository_iam_policy",
		},
		{
			Name: "google_assured_workloads_workload",
		},
		{
			Name: "google_beyondcorp_app_connection",
		},
		{
			Name: "google_beyondcorp_app_connector",
		},
		{
			Name: "google_beyondcorp_app_gateway",
		},
		{
			Name: "google_biglake_catalog",
		},
		{
			Name: "google_biglake_database",
		},
		{
			Name: "google_biglake_table",
		},
		{
			Name: "google_bigquery_analytics_hub_data_exchange",
		},
		{
			Name: "google_bigquery_analytics_hub_data_exchange_iam_binding",
		},
		{
			Name: "google_bigquery_analytics_hub_data_exchange_iam_member",
		},
		{
			Name: "google_bigquery_analytics_hub_data_exchange_iam_policy",
		},
		{
			Name: "google_bigquery_analytics_hub_listing",
		},
		{
			Name: "google_bigquery_analytics_hub_listing_iam_binding",
		},
		{
			Name: "google_bigquery_analytics_hub_listing_iam_member",
		},
		{
			Name: "google_bigquery_analytics_hub_listing_iam_policy",
		},
		{
			Name: "google_bigquery_bi_reservation",
		},
		{
			Name: "google_bigquery_capacity_commitment",
		},
		{
			Name: "google_bigquery_connection",
		},
		{
			Name: "google_bigquery_connection_iam_binding",
		},
		{
			Name: "google_bigquery_connection_iam_member",
		},
		{
			Name: "google_bigquery_connection_iam_policy",
		},
		{
			Name: "google_bigquery_data_transfer_config",
		},
		{
			Name: "google_bigquery_datapolicy_data_policy",
		},
		{
			Name: "google_bigquery_datapolicy_data_policy_iam_binding",
		},
		{
			Name: "google_bigquery_datapolicy_data_policy_iam_member",
		},
		{
			Name: "google_bigquery_datapolicy_data_policy_iam_policy",
		},
		{
			Name: "google_bigquery_dataset",
		},
		{
			Name: "google_bigquery_dataset_access",
		},
		{
			Name: "google_bigquery_dataset_iam_binding",
		},
		{
			Name: "google_bigquery_dataset_iam_member",
		},
		{
			Name: "google_bigquery_dataset_iam_policy",
		},
		{
			Name: "google_bigquery_job",
		},
		{
			Name: "google_bigquery_reservation",
		},
		{
			Name: "google_bigquery_reservation_assignment",
		},
		{
			Name: "google_bigquery_routine",
		},
		{
			Name: "google_bigquery_table",
		},
		{
			Name: "google_bigquery_table_iam_binding",
		},
		{
			Name: "google_bigquery_table_iam_member",
		},
		{
			Name: "google_bigquery_table_iam_policy",
		},
		{
			Name: "google_bigtable_app_profile",
		},
		{
			Name: "google_bigtable_gc_policy",
		},
		{
			Name: "google_bigtable_instance",
		},
		{
			Name: "google_bigtable_instance_iam_binding",
		},
		{
			Name: "google_bigtable_instance_iam_member",
		},
		{
			Name: "google_bigtable_instance_iam_policy",
		},
		{
			Name: "google_bigtable_table",
		},
		{
			Name: "google_bigtable_table_iam_binding",
		},
		{
			Name: "google_bigtable_table_iam_member",
		},
		{
			Name: "google_bigtable_table_iam_policy",
		},
		{
			Name: "google_billing_account_iam_binding",
		},
		{
			Name: "google_billing_account_iam_member",
		},
		{
			Name: "google_billing_account_iam_policy",
		},
		{
			Name: "google_billing_budget",
		},
		{
			Name: "google_billing_project_info",
		},
		{
			Name: "google_billing_subaccount",
		},
		{
			Name: "google_binary_authorization_attestor",
		},
		{
			Name: "google_binary_authorization_attestor_iam_binding",
		},
		{
			Name: "google_binary_authorization_attestor_iam_member",
		},
		{
			Name: "google_binary_authorization_attestor_iam_policy",
		},
		{
			Name: "google_binary_authorization_policy",
		},
		{
			Name: "google_certificate_manager_certificate",
		},
		{
			Name: "google_certificate_manager_certificate_issuance_config",
		},
		{
			Name: "google_certificate_manager_certificate_map",
		},
		{
			Name: "google_certificate_manager_certificate_map_entry",
		},
		{
			Name: "google_certificate_manager_dns_authorization",
		},
		{
			Name: "google_certificate_manager_trust_config",
		},
		{
			Name: "google_cloud_asset_folder_feed",
		},
		{
			Name: "google_cloud_asset_organization_feed",
		},
		{
			Name: "google_cloud_asset_project_feed",
		},
		{
			Name: "google_cloud_identity_group",
		},
		{
			Name: "google_cloud_identity_group_membership",
		},
		{
			Name: "google_cloud_ids_endpoint",
		},
		{
			Name: "google_cloud_run_domain_mapping",
		},
		{
			Name: "google_cloud_run_service",
		},
		{
			Name: "google_cloud_run_service_iam_binding",
		},
		{
			Name: "google_cloud_run_service_iam_member",
		},
		{
			Name: "google_cloud_run_service_iam_policy",
		},
		{
			Name: "google_cloud_run_v2_job",
		},
		{
			Name: "google_cloud_run_v2_job_iam_binding",
		},
		{
			Name: "google_cloud_run_v2_job_iam_member",
		},
		{
			Name: "google_cloud_run_v2_job_iam_policy",
		},
		{
			Name: "google_cloud_run_v2_service",
		},
		{
			Name: "google_cloud_run_v2_service_iam_binding",
		},
		{
			Name: "google_cloud_run_v2_service_iam_member",
		},
		{
			Name: "google_cloud_run_v2_service_iam_policy",
		},
		{
			Name: "google_cloud_scheduler_job",
		},
		{
			Name: "google_cloud_tasks_queue",
		},
		{
			Name: "google_cloud_tasks_queue_iam_binding",
		},
		{
			Name: "google_cloud_tasks_queue_iam_member",
		},
		{
			Name: "google_cloud_tasks_queue_iam_policy",
		},
		{
			Name: "google_cloudbuild_bitbucket_server_config",
		},
		{
			Name: "google_cloudbuild_trigger",
		},
		{
			Name: "google_cloudbuild_worker_pool",
		},
		{
			Name: "google_cloudbuildv2_connection",
		},
		{
			Name: "google_cloudbuildv2_connection_iam_binding",
		},
		{
			Name: "google_cloudbuildv2_connection_iam_member",
		},
		{
			Name: "google_cloudbuildv2_connection_iam_policy",
		},
		{
			Name: "google_cloudbuildv2_repository",
		},
		{
			Name: "google_clouddeploy_delivery_pipeline",
		},
		{
			Name: "google_clouddeploy_target",
		},
		{
			Name: "google_cloudfunctions2_function",
		},
		{
			Name: "google_cloudfunctions2_function_iam_binding",
		},
		{
			Name: "google_cloudfunctions2_function_iam_member",
		},
		{
			Name: "google_cloudfunctions2_function_iam_policy",
		},
		{
			Name: "google_cloudfunctions_function",
		},
		{
			Name: "google_cloudfunctions_function_iam_binding",
		},
		{
			Name: "google_cloudfunctions_function_iam_member",
		},
		{
			Name: "google_cloudfunctions_function_iam_policy",
		},
		{
			Name: "google_cloudiot_device",
		},
		{
			Name: "google_cloudiot_registry",
		},
		{
			Name: "google_cloudiot_registry_iam_binding",
		},
		{
			Name: "google_cloudiot_registry_iam_member",
		},
		{
			Name: "google_cloudiot_registry_iam_policy",
		},
		{
			Name: "google_composer_environment",
		},
		{
			Name: "google_compute_address",
		},
		{
			Name: "google_compute_attached_disk",
		},
		{
			Name: "google_compute_autoscaler",
		},
		{
			Name: "google_compute_backend_bucket",
		},
		{
			Name: "google_compute_backend_bucket_signed_url_key",
		},
		{
			Name: "google_compute_backend_service",
		},
		{
			Name: "google_compute_backend_service_signed_url_key",
		},
		{
			Name: "google_compute_disk",
		},
		{
			Name: "google_compute_disk_async_replication",
		},
		{
			Name: "google_compute_disk_iam_binding",
		},
		{
			Name: "google_compute_disk_iam_member",
		},
		{
			Name: "google_compute_disk_iam_policy",
		},
		{
			Name: "google_compute_disk_resource_policy_attachment",
		},
		{
			Name: "google_compute_external_vpn_gateway",
		},
		{
			Name: "google_compute_firewall",
		},
		{
			Name: "google_compute_firewall_policy",
		},
		{
			Name: "google_compute_firewall_policy_association",
		},
		{
			Name: "google_compute_firewall_policy_rule",
		},
		{
			Name: "google_compute_forwarding_rule",
		},
		{
			Name: "google_compute_global_address",
		},
		{
			Name: "google_compute_global_forwarding_rule",
		},
		{
			Name: "google_compute_global_network_endpoint",
		},
		{
			Name: "google_compute_global_network_endpoint_group",
		},
		{
			Name: "google_compute_ha_vpn_gateway",
		},
		{
			Name: "google_compute_health_check",
		},
		{
			Name: "google_compute_http_health_check",
		},
		{
			Name: "google_compute_https_health_check",
		},
		{
			Name: "google_compute_image",
		},
		{
			Name: "google_compute_image_iam_binding",
		},
		{
			Name: "google_compute_image_iam_member",
		},
		{
			Name: "google_compute_image_iam_policy",
		},
		{
			Name: "google_compute_instance",
		},
		{
			Name: "google_compute_instance_from_template",
		},
		{
			Name: "google_compute_instance_group",
		},
		{
			Name: "google_compute_instance_group_manager",
		},
		{
			Name: "google_compute_instance_group_named_port",
		},
		{
			Name: "google_compute_instance_iam_binding",
		},
		{
			Name: "google_compute_instance_iam_member",
		},
		{
			Name: "google_compute_instance_iam_policy",
		},
		{
			Name: "google_compute_instance_template",
		},
		{
			Name: "google_compute_interconnect_attachment",
		},
		{
			Name: "google_compute_managed_ssl_certificate",
		},
		{
			Name: "google_compute_network",
		},
		{
			Name: "google_compute_network_endpoint",
		},
		{
			Name: "google_compute_network_endpoint_group",
		},
		{
			Name: "google_compute_network_endpoints",
		},
		{
			Name: "google_compute_network_firewall_policy",
		},
		{
			Name: "google_compute_network_firewall_policy_association",
		},
		{
			Name: "google_compute_network_firewall_policy_rule",
		},
		{
			Name: "google_compute_network_peering",
		},
		{
			Name: "google_compute_network_peering_routes_config",
		},
		{
			Name: "google_compute_node_group",
		},
		{
			Name: "google_compute_node_template",
		},
		{
			Name: "google_compute_packet_mirroring",
		},
		{
			Name: "google_compute_per_instance_config",
		},
		{
			Name: "google_compute_project_default_network_tier",
		},
		{
			Name: "google_compute_project_metadata",
		},
		{
			Name: "google_compute_project_metadata_item",
		},
		{
			Name: "google_compute_public_advertised_prefix",
		},
		{
			Name: "google_compute_public_delegated_prefix",
		},
		{
			Name: "google_compute_region_autoscaler",
		},
		{
			Name: "google_compute_region_backend_service",
		},
		{
			Name: "google_compute_region_commitment",
		},
		{
			Name: "google_compute_region_disk",
		},
		{
			Name: "google_compute_region_disk_iam_binding",
		},
		{
			Name: "google_compute_region_disk_iam_member",
		},
		{
			Name: "google_compute_region_disk_iam_policy",
		},
		{
			Name: "google_compute_region_disk_resource_policy_attachment",
		},
		{
			Name: "google_compute_region_health_check",
		},
		{
			Name: "google_compute_region_instance_group_manager",
		},
		{
			Name: "google_compute_region_instance_template",
		},
		{
			Name: "google_compute_region_network_endpoint_group",
		},
		{
			Name: "google_compute_region_network_firewall_policy",
		},
		{
			Name: "google_compute_region_network_firewall_policy_association",
		},
		{
			Name: "google_compute_region_network_firewall_policy_rule",
		},
		{
			Name: "google_compute_region_per_instance_config",
		},
		{
			Name: "google_compute_region_ssl_certificate",
		},
		{
			Name: "google_compute_region_ssl_policy",
		},
		{
			Name: "google_compute_region_target_http_proxy",
		},
		{
			Name: "google_compute_region_target_https_proxy",
		},
		{
			Name: "google_compute_region_target_tcp_proxy",
		},
		{
			Name: "google_compute_region_url_map",
		},
		{
			Name: "google_compute_reservation",
		},
		{
			Name: "google_compute_resource_policy",
		},
		{
			Name: "google_compute_route",
		},
		{
			Name: "google_compute_router",
		},
		{
			Name: "google_compute_router_interface",
		},
		{
			Name: "google_compute_router_nat",
		},
		{
			Name: "google_compute_router_peer",
		},
		{
			Name: "google_compute_security_policy",
		},
		{
			Name: "google_compute_service_attachment",
		},
		{
			Name: "google_compute_shared_vpc_host_project",
		},
		{
			Name: "google_compute_shared_vpc_service_project",
		},
		{
			Name: "google_compute_snapshot",
		},
		{
			Name: "google_compute_snapshot_iam_binding",
		},
		{
			Name: "google_compute_snapshot_iam_member",
		},
		{
			Name: "google_compute_snapshot_iam_policy",
		},
		{
			Name: "google_compute_ssl_certificate",
		},
		{
			Name: "google_compute_ssl_policy",
		},
		{
			Name: "google_compute_subnetwork",
		},
		{
			Name: "google_compute_subnetwork_iam_binding",
		},
		{
			Name: "google_compute_subnetwork_iam_member",
		},
		{
			Name: "google_compute_subnetwork_iam_policy",
		},
		{
			Name: "google_compute_target_grpc_proxy",
		},
		{
			Name: "google_compute_target_http_proxy",
		},
		{
			Name: "google_compute_target_https_proxy",
		},
		{
			Name: "google_compute_target_instance",
		},
		{
			Name: "google_compute_target_pool",
		},
		{
			Name: "google_compute_target_ssl_proxy",
		},
		{
			Name: "google_compute_target_tcp_proxy",
		},
		{
			Name: "google_compute_url_map",
		},
		{
			Name: "google_compute_vpn_gateway",
		},
		{
			Name: "google_compute_vpn_tunnel",
		},
		{
			Name: "google_container_analysis_note",
		},
		{
			Name: "google_container_analysis_note_iam_binding",
		},
		{
			Name: "google_container_analysis_note_iam_member",
		},
		{
			Name: "google_container_analysis_note_iam_policy",
		},
		{
			Name: "google_container_analysis_occurrence",
		},
		{
			Name: "google_container_attached_cluster",
		},
		{
			Name: "google_container_aws_cluster",
		},
		{
			Name: "google_container_aws_node_pool",
		},
		{
			Name: "google_container_azure_client",
		},
		{
			Name: "google_container_azure_cluster",
		},
		{
			Name: "google_container_azure_node_pool",
		},
		{
			Name: "google_container_cluster",
		},
		{
			Name: "google_container_node_pool",
		},
		{
			Name: "google_container_registry",
		},
		{
			Name: "google_data_catalog_entry",
		},
		{
			Name: "google_data_catalog_entry_group",
		},
		{
			Name: "google_data_catalog_entry_group_iam_binding",
		},
		{
			Name: "google_data_catalog_entry_group_iam_member",
		},
		{
			Name: "google_data_catalog_entry_group_iam_policy",
		},
		{
			Name: "google_data_catalog_policy_tag",
		},
		{
			Name: "google_data_catalog_policy_tag_iam_binding",
		},
		{
			Name: "google_data_catalog_policy_tag_iam_member",
		},
		{
			Name: "google_data_catalog_policy_tag_iam_policy",
		},
		{
			Name: "google_data_catalog_tag",
		},
		{
			Name: "google_data_catalog_tag_template",
		},
		{
			Name: "google_data_catalog_tag_template_iam_binding",
		},
		{
			Name: "google_data_catalog_tag_template_iam_member",
		},
		{
			Name: "google_data_catalog_tag_template_iam_policy",
		},
		{
			Name: "google_data_catalog_taxonomy",
		},
		{
			Name: "google_data_catalog_taxonomy_iam_binding",
		},
		{
			Name: "google_data_catalog_taxonomy_iam_member",
		},
		{
			Name: "google_data_catalog_taxonomy_iam_policy",
		},
		{
			Name: "google_data_fusion_instance",
		},
		{
			Name: "google_data_fusion_instance_iam_binding",
		},
		{
			Name: "google_data_fusion_instance_iam_member",
		},
		{
			Name: "google_data_fusion_instance_iam_policy",
		},
		{
			Name: "google_data_loss_prevention_deidentify_template",
		},
		{
			Name: "google_data_loss_prevention_inspect_template",
		},
		{
			Name: "google_data_loss_prevention_job_trigger",
		},
		{
			Name: "google_data_loss_prevention_stored_info_type",
		},
		{
			Name: "google_data_pipeline_pipeline",
		},
		{
			Name: "google_database_migration_service_connection_profile",
		},
		{
			Name: "google_dataflow_job",
		},
		{
			Name: "google_dataplex_asset",
		},
		{
			Name: "google_dataplex_asset_iam_binding",
		},
		{
			Name: "google_dataplex_asset_iam_member",
		},
		{
			Name: "google_dataplex_asset_iam_policy",
		},
		{
			Name: "google_dataplex_datascan",
		},
		{
			Name: "google_dataplex_datascan_iam_binding",
		},
		{
			Name: "google_dataplex_datascan_iam_member",
		},
		{
			Name: "google_dataplex_datascan_iam_policy",
		},
		{
			Name: "google_dataplex_lake",
		},
		{
			Name: "google_dataplex_lake_iam_binding",
		},
		{
			Name: "google_dataplex_lake_iam_member",
		},
		{
			Name: "google_dataplex_lake_iam_policy",
		},
		{
			Name: "google_dataplex_task",
		},
		{
			Name: "google_dataplex_task_iam_binding",
		},
		{
			Name: "google_dataplex_task_iam_member",
		},
		{
			Name: "google_dataplex_task_iam_policy",
		},
		{
			Name: "google_dataplex_zone",
		},
		{
			Name: "google_dataplex_zone_iam_binding",
		},
		{
			Name: "google_dataplex_zone_iam_member",
		},
		{
			Name: "google_dataplex_zone_iam_policy",
		},
		{
			Name: "google_dataproc_autoscaling_policy",
		},
		{
			Name: "google_dataproc_autoscaling_policy_iam_binding",
		},
		{
			Name: "google_dataproc_autoscaling_policy_iam_member",
		},
		{
			Name: "google_dataproc_autoscaling_policy_iam_policy",
		},
		{
			Name: "google_dataproc_cluster",
		},
		{
			Name: "google_dataproc_cluster_iam_binding",
		},
		{
			Name: "google_dataproc_cluster_iam_member",
		},
		{
			Name: "google_dataproc_cluster_iam_policy",
		},
		{
			Name: "google_dataproc_job",
		},
		{
			Name: "google_dataproc_job_iam_binding",
		},
		{
			Name: "google_dataproc_job_iam_member",
		},
		{
			Name: "google_dataproc_job_iam_policy",
		},
		{
			Name: "google_dataproc_metastore_service",
		},
		{
			Name: "google_dataproc_metastore_service_iam_binding",
		},
		{
			Name: "google_dataproc_metastore_service_iam_member",
		},
		{
			Name: "google_dataproc_metastore_service_iam_policy",
		},
		{
			Name: "google_dataproc_workflow_template",
		},
		{
			Name: "google_datastore_index",
		},
		{
			Name: "google_datastream_connection_profile",
		},
		{
			Name: "google_datastream_private_connection",
		},
		{
			Name: "google_datastream_stream",
		},
		{
			Name: "google_deployment_manager_deployment",
		},
		{
			Name: "google_dialogflow_agent",
		},
		{
			Name: "google_dialogflow_cx_agent",
		},
		{
			Name: "google_dialogflow_cx_entity_type",
		},
		{
			Name: "google_dialogflow_cx_environment",
		},
		{
			Name: "google_dialogflow_cx_flow",
		},
		{
			Name: "google_dialogflow_cx_intent",
		},
		{
			Name: "google_dialogflow_cx_page",
		},
		{
			Name: "google_dialogflow_cx_test_case",
		},
		{
			Name: "google_dialogflow_cx_version",
		},
		{
			Name: "google_dialogflow_cx_webhook",
		},
		{
			Name: "google_dialogflow_entity_type",
		},
		{
			Name: "google_dialogflow_fulfillment",
		},
		{
			Name: "google_dialogflow_intent",
		},
		{
			Name: "google_dns_managed_zone",
		},
		{
			Name: "google_dns_managed_zone_iam_binding",
		},
		{
			Name: "google_dns_managed_zone_iam_member",
		},
		{
			Name: "google_dns_managed_zone_iam_policy",
		},
		{
			Name: "google_dns_policy",
		},
		{
			Name: "google_dns_record_set",
		},
		{
			Name: "google_dns_response_policy",
		},
		{
			Name: "google_dns_response_policy_rule",
		},
		{
			Name: "google_document_ai_processor",
		},
		{
			Name: "google_document_ai_processor_default_version",
		},
		{
			Name: "google_document_ai_warehouse_document_schema",
		},
		{
			Name: "google_document_ai_warehouse_location",
		},
		{
			Name: "google_endpoints_service",
		},
		{
			Name: "google_endpoints_service_consumers_iam_binding",
		},
		{
			Name: "google_endpoints_service_consumers_iam_member",
		},
		{
			Name: "google_endpoints_service_consumers_iam_policy",
		},
		{
			Name: "google_endpoints_service_iam_binding",
		},
		{
			Name: "google_endpoints_service_iam_member",
		},
		{
			Name: "google_endpoints_service_iam_policy",
		},
		{
			Name: "google_essential_contacts_contact",
		},
		{
			Name: "google_eventarc_channel",
		},
		{
			Name: "google_eventarc_google_channel_config",
		},
		{
			Name: "google_eventarc_trigger",
		},
		{
			Name: "google_filestore_backup",
		},
		{
			Name: "google_filestore_instance",
		},
		{
			Name: "google_filestore_snapshot",
		},
		{
			Name: "google_firebaserules_release",
		},
		{
			Name: "google_firebaserules_ruleset",
		},
		{
			Name: "google_firestore_database",
		},
		{
			Name: "google_firestore_document",
		},
		{
			Name: "google_firestore_field",
		},
		{
			Name: "google_firestore_index",
		},
		{
			Name: "google_folder",
		},
		{
			Name: "google_folder_access_approval_settings",
		},
		{
			Name: "google_folder_iam_audit_config",
		},
		{
			Name: "google_folder_iam_binding",
		},
		{
			Name: "google_folder_iam_member",
		},
		{
			Name: "google_folder_iam_policy",
		},
		{
			Name: "google_folder_organization_policy",
		},
		{
			Name: "google_game_services_game_server_cluster",
		},
		{
			Name: "google_game_services_game_server_config",
		},
		{
			Name: "google_game_services_game_server_deployment",
		},
		{
			Name: "google_game_services_game_server_deployment_rollout",
		},
		{
			Name: "google_game_services_realm",
		},
		{
			Name: "google_gke_backup_backup_plan",
		},
		{
			Name: "google_gke_backup_backup_plan_iam_binding",
		},
		{
			Name: "google_gke_backup_backup_plan_iam_member",
		},
		{
			Name: "google_gke_backup_backup_plan_iam_policy",
		},
		{
			Name: "google_gke_hub_feature",
		},
		{
			Name: "google_gke_hub_feature_iam_binding",
		},
		{
			Name: "google_gke_hub_feature_iam_member",
		},
		{
			Name: "google_gke_hub_feature_iam_policy",
		},
		{
			Name: "google_gke_hub_feature_membership",
		},
		{
			Name: "google_gke_hub_membership",
		},
		{
			Name: "google_gke_hub_membership_binding",
		},
		{
			Name: "google_gke_hub_membership_iam_binding",
		},
		{
			Name: "google_gke_hub_membership_iam_member",
		},
		{
			Name: "google_gke_hub_membership_iam_policy",
		},
		{
			Name: "google_gke_hub_namespace",
		},
		{
			Name: "google_gke_hub_scope",
		},
		{
			Name: "google_gke_hub_scope_iam_binding",
		},
		{
			Name: "google_gke_hub_scope_iam_member",
		},
		{
			Name: "google_gke_hub_scope_iam_policy",
		},
		{
			Name: "google_gke_hub_scope_rbac_role_binding",
		},
		{
			Name: "google_healthcare_consent_store",
		},
		{
			Name: "google_healthcare_consent_store_iam_binding",
		},
		{
			Name: "google_healthcare_consent_store_iam_member",
		},
		{
			Name: "google_healthcare_consent_store_iam_policy",
		},
		{
			Name: "google_healthcare_dataset",
		},
		{
			Name: "google_healthcare_dataset_iam_binding",
		},
		{
			Name: "google_healthcare_dataset_iam_member",
		},
		{
			Name: "google_healthcare_dataset_iam_policy",
		},
		{
			Name: "google_healthcare_dicom_store",
		},
		{
			Name: "google_healthcare_dicom_store_iam_binding",
		},
		{
			Name: "google_healthcare_dicom_store_iam_member",
		},
		{
			Name: "google_healthcare_dicom_store_iam_policy",
		},
		{
			Name: "google_healthcare_fhir_store",
		},
		{
			Name: "google_healthcare_fhir_store_iam_binding",
		},
		{
			Name: "google_healthcare_fhir_store_iam_member",
		},
		{
			Name: "google_healthcare_fhir_store_iam_policy",
		},
		{
			Name: "google_healthcare_hl7_v2_store",
		},
		{
			Name: "google_healthcare_hl7_v2_store_iam_binding",
		},
		{
			Name: "google_healthcare_hl7_v2_store_iam_member",
		},
		{
			Name: "google_healthcare_hl7_v2_store_iam_policy",
		},
		{
			Name: "google_iam_access_boundary_policy",
		},
		{
			Name: "google_iam_deny_policy",
		},
		{
			Name: "google_iam_workforce_pool",
		},
		{
			Name: "google_iam_workforce_pool_provider",
		},
		{
			Name: "google_iam_workload_identity_pool",
		},
		{
			Name: "google_iam_workload_identity_pool_provider",
		},
		{
			Name: "google_iap_app_engine_service_iam_binding",
		},
		{
			Name: "google_iap_app_engine_service_iam_member",
		},
		{
			Name: "google_iap_app_engine_service_iam_policy",
		},
		{
			Name: "google_iap_app_engine_version_iam_binding",
		},
		{
			Name: "google_iap_app_engine_version_iam_member",
		},
		{
			Name: "google_iap_app_engine_version_iam_policy",
		},
		{
			Name: "google_iap_brand",
		},
		{
			Name: "google_iap_client",
		},
		{
			Name: "google_iap_tunnel_iam_binding",
		},
		{
			Name: "google_iap_tunnel_iam_member",
		},
		{
			Name: "google_iap_tunnel_iam_policy",
		},
		{
			Name: "google_iap_tunnel_instance_iam_binding",
		},
		{
			Name: "google_iap_tunnel_instance_iam_member",
		},
		{
			Name: "google_iap_tunnel_instance_iam_policy",
		},
		{
			Name: "google_iap_web_backend_service_iam_binding",
		},
		{
			Name: "google_iap_web_backend_service_iam_member",
		},
		{
			Name: "google_iap_web_backend_service_iam_policy",
		},
		{
			Name: "google_iap_web_iam_binding",
		},
		{
			Name: "google_iap_web_iam_member",
		},
		{
			Name: "google_iap_web_iam_policy",
		},
		{
			Name: "google_iap_web_region_backend_service_iam_binding",
		},
		{
			Name: "google_iap_web_region_backend_service_iam_member",
		},
		{
			Name: "google_iap_web_region_backend_service_iam_policy",
		},
		{
			Name: "google_iap_web_type_app_engine_iam_binding",
		},
		{
			Name: "google_iap_web_type_app_engine_iam_member",
		},
		{
			Name: "google_iap_web_type_app_engine_iam_policy",
		},
		{
			Name: "google_iap_web_type_compute_iam_binding",
		},
		{
			Name: "google_iap_web_type_compute_iam_member",
		},
		{
			Name: "google_iap_web_type_compute_iam_policy",
		},
		{
			Name: "google_identity_platform_config",
		},
		{
			Name: "google_identity_platform_default_supported_idp_config",
		},
		{
			Name: "google_identity_platform_inbound_saml_config",
		},
		{
			Name: "google_identity_platform_oauth_idp_config",
		},
		{
			Name: "google_identity_platform_project_default_config",
		},
		{
			Name: "google_identity_platform_tenant",
		},
		{
			Name: "google_identity_platform_tenant_default_supported_idp_config",
		},
		{
			Name: "google_identity_platform_tenant_inbound_saml_config",
		},
		{
			Name: "google_identity_platform_tenant_oauth_idp_config",
		},
		{
			Name: "google_kms_crypto_key",
		},
		{
			Name: "google_kms_crypto_key_iam_binding",
		},
		{
			Name: "google_kms_crypto_key_iam_member",
		},
		{
			Name: "google_kms_crypto_key_iam_policy",
		},
		{
			Name: "google_kms_crypto_key_version",
		},
		{
			Name: "google_kms_key_ring",
		},
		{
			Name: "google_kms_key_ring_iam_binding",
		},
		{
			Name: "google_kms_key_ring_iam_member",
		},
		{
			Name: "google_kms_key_ring_iam_policy",
		},
		{
			Name: "google_kms_key_ring_import_job",
		},
		{
			Name: "google_kms_secret_ciphertext",
		},
		{
			Name: "google_logging_billing_account_bucket_config",
		},
		{
			Name: "google_logging_billing_account_exclusion",
		},
		{
			Name: "google_logging_billing_account_sink",
		},
		{
			Name: "google_logging_folder_bucket_config",
		},
		{
			Name: "google_logging_folder_exclusion",
		},
		{
			Name: "google_logging_folder_sink",
		},
		{
			Name: "google_logging_linked_dataset",
		},
		{
			Name: "google_logging_log_view",
		},
		{
			Name: "google_logging_metric",
		},
		{
			Name: "google_logging_organization_bucket_config",
		},
		{
			Name: "google_logging_organization_exclusion",
		},
		{
			Name: "google_logging_organization_sink",
		},
		{
			Name: "google_logging_project_bucket_config",
		},
		{
			Name: "google_logging_project_exclusion",
		},
		{
			Name: "google_logging_project_sink",
		},
		{
			Name: "google_looker_instance",
		},
		{
			Name: "google_memcache_instance",
		},
		{
			Name: "google_ml_engine_model",
		},
		{
			Name: "google_monitoring_alert_policy",
		},
		{
			Name: "google_monitoring_custom_service",
		},
		{
			Name: "google_monitoring_dashboard",
		},
		{
			Name: "google_monitoring_group",
		},
		{
			Name: "google_monitoring_metric_descriptor",
		},
		{
			Name: "google_monitoring_monitored_project",
		},
		{
			Name: "google_monitoring_notification_channel",
		},
		{
			Name: "google_monitoring_service",
		},
		{
			Name: "google_monitoring_slo",
		},
		{
			Name: "google_monitoring_uptime_check_config",
		},
		{
			Name: "google_network_connectivity_hub",
		},
		{
			Name: "google_network_connectivity_service_connection_policy",
		},
		{
			Name: "google_network_connectivity_spoke",
		},
		{
			Name: "google_network_management_connectivity_test",
		},
		{
			Name: "google_network_security_address_group",
		},
		{
			Name: "google_network_security_gateway_security_policy",
		},
		{
			Name: "google_network_security_gateway_security_policy_rule",
		},
		{
			Name: "google_network_security_url_lists",
		},
		{
			Name: "google_network_services_edge_cache_keyset",
		},
		{
			Name: "google_network_services_edge_cache_origin",
		},
		{
			Name: "google_network_services_edge_cache_service",
		},
		{
			Name: "google_network_services_gateway",
		},
		{
			Name: "google_notebooks_environment",
		},
		{
			Name: "google_notebooks_instance",
		},
		{
			Name: "google_notebooks_instance_iam_binding",
		},
		{
			Name: "google_notebooks_instance_iam_member",
		},
		{
			Name: "google_notebooks_instance_iam_policy",
		},
		{
			Name: "google_notebooks_location",
		},
		{
			Name: "google_notebooks_runtime",
		},
		{
			Name: "google_notebooks_runtime_iam_binding",
		},
		{
			Name: "google_notebooks_runtime_iam_member",
		},
		{
			Name: "google_notebooks_runtime_iam_policy",
		},
		{
			Name: "google_org_policy_policy",
		},
		{
			Name: "google_organization_access_approval_settings",
		},
		{
			Name: "google_organization_iam_audit_config",
		},
		{
			Name: "google_organization_iam_binding",
		},
		{
			Name: "google_organization_iam_custom_role",
		},
		{
			Name: "google_organization_iam_member",
		},
		{
			Name: "google_organization_iam_policy",
		},
		{
			Name: "google_organization_policy",
		},
		{
			Name: "google_os_config_os_policy_assignment",
		},
		{
			Name: "google_os_config_patch_deployment",
		},
		{
			Name: "google_os_login_ssh_public_key",
		},
		{
			Name: "google_privateca_ca_pool",
		},
		{
			Name: "google_privateca_ca_pool_iam_binding",
		},
		{
			Name: "google_privateca_ca_pool_iam_member",
		},
		{
			Name: "google_privateca_ca_pool_iam_policy",
		},
		{
			Name: "google_privateca_certificate",
		},
		{
			Name: "google_privateca_certificate_authority",
		},
		{
			Name: "google_privateca_certificate_template",
		},
		{
			Name: "google_privateca_certificate_template_iam_binding",
		},
		{
			Name: "google_privateca_certificate_template_iam_member",
		},
		{
			Name: "google_privateca_certificate_template_iam_policy",
		},
		{
			Name: "google_project",
		},
		{
			Name: "google_project_access_approval_settings",
		},
		{
			Name: "google_project_default_service_accounts",
		},
		{
			Name: "google_project_iam_audit_config",
		},
		{
			Name: "google_project_iam_binding",
		},
		{
			Name: "google_project_iam_custom_role",
		},
		{
			Name: "google_project_iam_member",
		},
		{
			Name: "google_project_iam_policy",
		},
		{
			Name: "google_project_organization_policy",
		},
		{
			Name: "google_project_service",
		},
		{
			Name: "google_project_usage_export_bucket",
		},
		{
			Name: "google_public_ca_external_account_key",
		},
		{
			Name: "google_pubsub_lite_reservation",
		},
		{
			Name: "google_pubsub_lite_subscription",
		},
		{
			Name: "google_pubsub_lite_topic",
		},
		{
			Name: "google_pubsub_schema",
		},
		{
			Name: "google_pubsub_subscription",
		},
		{
			Name: "google_pubsub_subscription_iam_binding",
		},
		{
			Name: "google_pubsub_subscription_iam_member",
		},
		{
			Name: "google_pubsub_subscription_iam_policy",
		},
		{
			Name: "google_pubsub_topic",
		},
		{
			Name: "google_pubsub_topic_iam_binding",
		},
		{
			Name: "google_pubsub_topic_iam_member",
		},
		{
			Name: "google_pubsub_topic_iam_policy",
		},
		{
			Name: "google_recaptcha_enterprise_key",
		},
		{
			Name: "google_redis_instance",
		},
		{
			Name: "google_resource_manager_lien",
		},
		{
			Name: "google_scc_mute_config",
		},
		{
			Name: "google_scc_notification_config",
		},
		{
			Name: "google_scc_source",
		},
		{
			Name: "google_scc_source_iam_binding",
		},
		{
			Name: "google_scc_source_iam_member",
		},
		{
			Name: "google_scc_source_iam_policy",
		},
		{
			Name: "google_secret_manager_secret",
		},
		{
			Name: "google_secret_manager_secret_iam_binding",
		},
		{
			Name: "google_secret_manager_secret_iam_member",
		},
		{
			Name: "google_secret_manager_secret_iam_policy",
		},
		{
			Name: "google_secret_manager_secret_version",
		},
		{
			Name: "google_service_account",
		},
		{
			Name: "google_service_account_iam_binding",
		},
		{
			Name: "google_service_account_iam_member",
		},
		{
			Name: "google_service_account_iam_policy",
		},
		{
			Name: "google_service_account_key",
		},
		{
			Name: "google_service_networking_connection",
		},
		{
			Name: "google_service_networking_peered_dns_domain",
		},
		{
			Name: "google_sourcerepo_repository",
		},
		{
			Name: "google_sourcerepo_repository_iam_binding",
		},
		{
			Name: "google_sourcerepo_repository_iam_member",
		},
		{
			Name: "google_sourcerepo_repository_iam_policy",
		},
		{
			Name: "google_spanner_database",
		},
		{
			Name: "google_spanner_database_iam_binding",
		},
		{
			Name: "google_spanner_database_iam_member",
		},
		{
			Name: "google_spanner_database_iam_policy",
		},
		{
			Name: "google_spanner_instance",
		},
		{
			Name: "google_spanner_instance_iam_binding",
		},
		{
			Name: "google_spanner_instance_iam_member",
		},
		{
			Name: "google_spanner_instance_iam_policy",
		},
		{
			Name: "google_sql_database",
		},
		{
			Name: "google_sql_database_instance",
		},
		{
			Name: "google_sql_source_representation_instance",
		},
		{
			Name: "google_sql_ssl_cert",
		},
		{
			Name: "google_sql_user",
		},
		{
			Name: "google_storage_bucket",
		},
		{
			Name: "google_storage_bucket_access_control",
		},
		{
			Name: "google_storage_bucket_acl",
		},
		{
			Name: "google_storage_bucket_iam_binding",
		},
		{
			Name: "google_storage_bucket_iam_member",
		},
		{
			Name: "google_storage_bucket_iam_policy",
		},
		{
			Name: "google_storage_bucket_object",
		},
		{
			Name: "google_storage_default_object_access_control",
		},
		{
			Name: "google_storage_default_object_acl",
		},
		{
			Name: "google_storage_hmac_key",
		},
		{
			Name: "google_storage_insights_report_config",
		},
		{
			Name: "google_storage_notification",
		},
		{
			Name: "google_storage_object_access_control",
		},
		{
			Name: "google_storage_object_acl",
		},
		{
			Name: "google_storage_transfer_agent_pool",
		},
		{
			Name: "google_storage_transfer_job",
		},
		{
			Name: "google_tags_location_tag_binding",
		},
		{
			Name: "google_tags_tag_binding",
		},
		{
			Name: "google_tags_tag_key",
		},
		{
			Name: "google_tags_tag_key_iam_binding",
		},
		{
			Name: "google_tags_tag_key_iam_member",
		},
		{
			Name: "google_tags_tag_key_iam_policy",
		},
		{
			Name: "google_tags_tag_value",
		},
		{
			Name: "google_tags_tag_value_iam_binding",
		},
		{
			Name: "google_tags_tag_value_iam_member",
		},
		{
			Name: "google_tags_tag_value_iam_policy",
		},
		{
			Name: "google_tpu_node",
		},
		{
			Name: "google_vertex_ai_dataset",
		},
		{
			Name: "google_vertex_ai_endpoint",
		},
		{
			Name: "google_vertex_ai_featurestore",
		},
		{
			Name: "google_vertex_ai_featurestore_entitytype",
		},
		{
			Name: "google_vertex_ai_featurestore_entitytype_feature",
		},
		{
			Name: "google_vertex_ai_index",
		},
		{
			Name: "google_vertex_ai_index_endpoint",
		},
		{
			Name: "google_vertex_ai_tensorboard",
		},
		{
			Name: "google_vpc_access_connector",
		},
		{
			Name: "google_workflows_workflow",
		},
	}
	resourcesMap = map[string]int{
		"google_access_context_manager_access_level":                     0,
		"google_access_context_manager_access_level_condition":           1,
		"google_access_context_manager_access_levels":                    2,
		"google_access_context_manager_access_policy":                    3,
		"google_access_context_manager_access_policy_iam_binding":        4,
		"google_access_context_manager_access_policy_iam_member":         5,
		"google_access_context_manager_access_policy_iam_policy":         6,
		"google_access_context_manager_authorized_orgs_desc":             7,
		"google_access_context_manager_egress_policy":                    8,
		"google_access_context_manager_gcp_user_access_binding":          9,
		"google_access_context_manager_ingress_policy":                   10,
		"google_access_context_manager_service_perimeter":                11,
		"google_access_context_manager_service_perimeter_egress_policy":  12,
		"google_access_context_manager_service_perimeter_ingress_policy": 13,
		"google_access_context_manager_service_perimeter_resource":       14,
		"google_access_context_manager_service_perimeters":               15,
		"google_active_directory_domain":                                 16,
		"google_active_directory_domain_trust":                           17,
		"google_alloydb_backup":                                          18,
		"google_alloydb_cluster":                                         19,
		"google_alloydb_instance":                                        20,
		"google_apigee_addons_config":                                    21,
		"google_apigee_endpoint_attachment":                              22,
		"google_apigee_env_keystore":                                     23,
		"google_apigee_env_references":                                   24,
		"google_apigee_envgroup":                                         25,
		"google_apigee_envgroup_attachment":                              26,
		"google_apigee_environment":                                      27,
		"google_apigee_environment_iam_binding":                          28,
		"google_apigee_environment_iam_member":                           29,
		"google_apigee_environment_iam_policy":                           30,
		"google_apigee_flowhook":                                         31,
		"google_apigee_instance":                                         32,
		"google_apigee_instance_attachment":                              33,
		"google_apigee_keystores_aliases_key_cert_file":                  34,
		"google_apigee_keystores_aliases_pkcs12":                         35,
		"google_apigee_keystores_aliases_self_signed_cert":               36,
		"google_apigee_nat_address":                                      37,
		"google_apigee_organization":                                     38,
		"google_apigee_sharedflow":                                       39,
		"google_apigee_sharedflow_deployment":                            40,
		"google_apigee_sync_authorization":                               41,
		"google_apigee_target_server":                                    42,
		"google_apikeys_key":                                             43,
		"google_app_engine_application":                                  44,
		"google_app_engine_application_url_dispatch_rules":               45,
		"google_app_engine_domain_mapping":                               46,
		"google_app_engine_firewall_rule":                                47,
		"google_app_engine_flexible_app_version":                         48,
		"google_app_engine_service_network_settings":                     49,
		"google_app_engine_service_split_traffic":                        50,
		"google_app_engine_standard_app_version":                         51,
		"google_artifact_registry_repository":                            52,
		"google_artifact_registry_repository_iam_binding":                53,
		"google_artifact_registry_repository_iam_member":                 54,
		"google_artifact_registry_repository_iam_policy":                 55,
		"google_assured_workloads_workload":                              56,
		"google_beyondcorp_app_connection":                               57,
		"google_beyondcorp_app_connector":                                58,
		"google_beyondcorp_app_gateway":                                  59,
		"google_biglake_catalog":                                         60,
		"google_biglake_database":                                        61,
		"google_biglake_table":                                           62,
		"google_bigquery_analytics_hub_data_exchange":                    63,
		"google_bigquery_analytics_hub_data_exchange_iam_binding":        64,
		"google_bigquery_analytics_hub_data_exchange_iam_member":         65,
		"google_bigquery_analytics_hub_data_exchange_iam_policy":         66,
		"google_bigquery_analytics_hub_listing":                          67,
		"google_bigquery_analytics_hub_listing_iam_binding":              68,
		"google_bigquery_analytics_hub_listing_iam_member":               69,
		"google_bigquery_analytics_hub_listing_iam_policy":               70,
		"google_bigquery_bi_reservation":                                 71,
		"google_bigquery_capacity_commitment":                            72,
		"google_bigquery_connection":                                     73,
		"google_bigquery_connection_iam_binding":                         74,
		"google_bigquery_connection_iam_member":                          75,
		"google_bigquery_connection_iam_policy":                          76,
		"google_bigquery_data_transfer_config":                           77,
		"google_bigquery_datapolicy_data_policy":                         78,
		"google_bigquery_datapolicy_data_policy_iam_binding":             79,
		"google_bigquery_datapolicy_data_policy_iam_member":              80,
		"google_bigquery_datapolicy_data_policy_iam_policy":              81,
		"google_bigquery_dataset":                                        82,
		"google_bigquery_dataset_access":                                 83,
		"google_bigquery_dataset_iam_binding":                            84,
		"google_bigquery_dataset_iam_member":                             85,
		"google_bigquery_dataset_iam_policy":                             86,
		"google_bigquery_job":                                            87,
		"google_bigquery_reservation":                                    88,
		"google_bigquery_reservation_assignment":                         89,
		"google_bigquery_routine":                                        90,
		"google_bigquery_table":                                          91,
		"google_bigquery_table_iam_binding":                              92,
		"google_bigquery_table_iam_member":                               93,
		"google_bigquery_table_iam_policy":                               94,
		"google_bigtable_app_profile":                                    95,
		"google_bigtable_gc_policy":                                      96,
		"google_bigtable_instance":                                       97,
		"google_bigtable_instance_iam_binding":                           98,
		"google_bigtable_instance_iam_member":                            99,
		"google_bigtable_instance_iam_policy":                            100,
		"google_bigtable_table":                                          101,
		"google_bigtable_table_iam_binding":                              102,
		"google_bigtable_table_iam_member":                               103,
		"google_bigtable_table_iam_policy":                               104,
		"google_billing_account_iam_binding":                             105,
		"google_billing_account_iam_member":                              106,
		"google_billing_account_iam_policy":                              107,
		"google_billing_budget":                                          108,
		"google_billing_project_info":                                    109,
		"google_billing_subaccount":                                      110,
		"google_binary_authorization_attestor":                           111,
		"google_binary_authorization_attestor_iam_binding":               112,
		"google_binary_authorization_attestor_iam_member":                113,
		"google_binary_authorization_attestor_iam_policy":                114,
		"google_binary_authorization_policy":                             115,
		"google_certificate_manager_certificate":                         116,
		"google_certificate_manager_certificate_issuance_config":         117,
		"google_certificate_manager_certificate_map":                     118,
		"google_certificate_manager_certificate_map_entry":               119,
		"google_certificate_manager_dns_authorization":                   120,
		"google_certificate_manager_trust_config":                        121,
		"google_cloud_asset_folder_feed":                                 122,
		"google_cloud_asset_organization_feed":                           123,
		"google_cloud_asset_project_feed":                                124,
		"google_cloud_identity_group":                                    125,
		"google_cloud_identity_group_membership":                         126,
		"google_cloud_ids_endpoint":                                      127,
		"google_cloud_run_domain_mapping":                                128,
		"google_cloud_run_service":                                       129,
		"google_cloud_run_service_iam_binding":                           130,
		"google_cloud_run_service_iam_member":                            131,
		"google_cloud_run_service_iam_policy":                            132,
		"google_cloud_run_v2_job":                                        133,
		"google_cloud_run_v2_job_iam_binding":                            134,
		"google_cloud_run_v2_job_iam_member":                             135,
		"google_cloud_run_v2_job_iam_policy":                             136,
		"google_cloud_run_v2_service":                                    137,
		"google_cloud_run_v2_service_iam_binding":                        138,
		"google_cloud_run_v2_service_iam_member":                         139,
		"google_cloud_run_v2_service_iam_policy":                         140,
		"google_cloud_scheduler_job":                                     141,
		"google_cloud_tasks_queue":                                       142,
		"google_cloud_tasks_queue_iam_binding":                           143,
		"google_cloud_tasks_queue_iam_member":                            144,
		"google_cloud_tasks_queue_iam_policy":                            145,
		"google_cloudbuild_bitbucket_server_config":                      146,
		"google_cloudbuild_trigger":                                      147,
		"google_cloudbuild_worker_pool":                                  148,
		"google_cloudbuildv2_connection":                                 149,
		"google_cloudbuildv2_connection_iam_binding":                     150,
		"google_cloudbuildv2_connection_iam_member":                      151,
		"google_cloudbuildv2_connection_iam_policy":                      152,
		"google_cloudbuildv2_repository":                                 153,
		"google_clouddeploy_delivery_pipeline":                           154,
		"google_clouddeploy_target":                                      155,
		"google_cloudfunctions2_function":                                156,
		"google_cloudfunctions2_function_iam_binding":                    157,
		"google_cloudfunctions2_function_iam_member":                     158,
		"google_cloudfunctions2_function_iam_policy":                     159,
		"google_cloudfunctions_function":                                 160,
		"google_cloudfunctions_function_iam_binding":                     161,
		"google_cloudfunctions_function_iam_member":                      162,
		"google_cloudfunctions_function_iam_policy":                      163,
		"google_cloudiot_device":                                         164,
		"google_cloudiot_registry":                                       165,
		"google_cloudiot_registry_iam_binding":                           166,
		"google_cloudiot_registry_iam_member":                            167,
		"google_cloudiot_registry_iam_policy":                            168,
		"google_composer_environment":                                    169,
		"google_compute_address":                                         170,
		"google_compute_attached_disk":                                   171,
		"google_compute_autoscaler":                                      172,
		"google_compute_backend_bucket":                                  173,
		"google_compute_backend_bucket_signed_url_key":                   174,
		"google_compute_backend_service":                                 175,
		"google_compute_backend_service_signed_url_key":                  176,
		"google_compute_disk":                                            177,
		"google_compute_disk_async_replication":                          178,
		"google_compute_disk_iam_binding":                                179,
		"google_compute_disk_iam_member":                                 180,
		"google_compute_disk_iam_policy":                                 181,
		"google_compute_disk_resource_policy_attachment":                 182,
		"google_compute_external_vpn_gateway":                            183,
		"google_compute_firewall":                                        184,
		"google_compute_firewall_policy":                                 185,
		"google_compute_firewall_policy_association":                     186,
		"google_compute_firewall_policy_rule":                            187,
		"google_compute_forwarding_rule":                                 188,
		"google_compute_global_address":                                  189,
		"google_compute_global_forwarding_rule":                          190,
		"google_compute_global_network_endpoint":                         191,
		"google_compute_global_network_endpoint_group":                   192,
		"google_compute_ha_vpn_gateway":                                  193,
		"google_compute_health_check":                                    194,
		"google_compute_http_health_check":                               195,
		"google_compute_https_health_check":                              196,
		"google_compute_image":                                           197,
		"google_compute_image_iam_binding":                               198,
		"google_compute_image_iam_member":                                199,
		"google_compute_image_iam_policy":                                200,
		"google_compute_instance":                                        201,
		"google_compute_instance_from_template":                          202,
		"google_compute_instance_group":                                  203,
		"google_compute_instance_group_manager":                          204,
		"google_compute_instance_group_named_port":                       205,
		"google_compute_instance_iam_binding":                            206,
		"google_compute_instance_iam_member":                             207,
		"google_compute_instance_iam_policy":                             208,
		"google_compute_instance_template":                               209,
		"google_compute_interconnect_attachment":                         210,
		"google_compute_managed_ssl_certificate":                         211,
		"google_compute_network":                                         212,
		"google_compute_network_endpoint":                                213,
		"google_compute_network_endpoint_group":                          214,
		"google_compute_network_endpoints":                               215,
		"google_compute_network_firewall_policy":                         216,
		"google_compute_network_firewall_policy_association":             217,
		"google_compute_network_firewall_policy_rule":                    218,
		"google_compute_network_peering":                                 219,
		"google_compute_network_peering_routes_config":                   220,
		"google_compute_node_group":                                      221,
		"google_compute_node_template":                                   222,
		"google_compute_packet_mirroring":                                223,
		"google_compute_per_instance_config":                             224,
		"google_compute_project_default_network_tier":                    225,
		"google_compute_project_metadata":                                226,
		"google_compute_project_metadata_item":                           227,
		"google_compute_public_advertised_prefix":                        228,
		"google_compute_public_delegated_prefix":                         229,
		"google_compute_region_autoscaler":                               230,
		"google_compute_region_backend_service":                          231,
		"google_compute_region_commitment":                               232,
		"google_compute_region_disk":                                     233,
		"google_compute_region_disk_iam_binding":                         234,
		"google_compute_region_disk_iam_member":                          235,
		"google_compute_region_disk_iam_policy":                          236,
		"google_compute_region_disk_resource_policy_attachment":          237,
		"google_compute_region_health_check":                             238,
		"google_compute_region_instance_group_manager":                   239,
		"google_compute_region_instance_template":                        240,
		"google_compute_region_network_endpoint_group":                   241,
		"google_compute_region_network_firewall_policy":                  242,
		"google_compute_region_network_firewall_policy_association":      243,
		"google_compute_region_network_firewall_policy_rule":             244,
		"google_compute_region_per_instance_config":                      245,
		"google_compute_region_ssl_certificate":                          246,
		"google_compute_region_ssl_policy":                               247,
		"google_compute_region_target_http_proxy":                        248,
		"google_compute_region_target_https_proxy":                       249,
		"google_compute_region_target_tcp_proxy":                         250,
		"google_compute_region_url_map":                                  251,
		"google_compute_reservation":                                     252,
		"google_compute_resource_policy":                                 253,
		"google_compute_route":                                           254,
		"google_compute_router":                                          255,
		"google_compute_router_interface":                                256,
		"google_compute_router_nat":                                      257,
		"google_compute_router_peer":                                     258,
		"google_compute_security_policy":                                 259,
		"google_compute_service_attachment":                              260,
		"google_compute_shared_vpc_host_project":                         261,
		"google_compute_shared_vpc_service_project":                      262,
		"google_compute_snapshot":                                        263,
		"google_compute_snapshot_iam_binding":                            264,
		"google_compute_snapshot_iam_member":                             265,
		"google_compute_snapshot_iam_policy":                             266,
		"google_compute_ssl_certificate":                                 267,
		"google_compute_ssl_policy":                                      268,
		"google_compute_subnetwork":                                      269,
		"google_compute_subnetwork_iam_binding":                          270,
		"google_compute_subnetwork_iam_member":                           271,
		"google_compute_subnetwork_iam_policy":                           272,
		"google_compute_target_grpc_proxy":                               273,
		"google_compute_target_http_proxy":                               274,
		"google_compute_target_https_proxy":                              275,
		"google_compute_target_instance":                                 276,
		"google_compute_target_pool":                                     277,
		"google_compute_target_ssl_proxy":                                278,
		"google_compute_target_tcp_proxy":                                279,
		"google_compute_url_map":                                         280,
		"google_compute_vpn_gateway":                                     281,
		"google_compute_vpn_tunnel":                                      282,
		"google_container_analysis_note":                                 283,
		"google_container_analysis_note_iam_binding":                     284,
		"google_container_analysis_note_iam_member":                      285,
		"google_container_analysis_note_iam_policy":                      286,
		"google_container_analysis_occurrence":                           287,
		"google_container_attached_cluster":                              288,
		"google_container_aws_cluster":                                   289,
		"google_container_aws_node_pool":                                 290,
		"google_container_azure_client":                                  291,
		"google_container_azure_cluster":                                 292,
		"google_container_azure_node_pool":                               293,
		"google_container_cluster":                                       294,
		"google_container_node_pool":                                     295,
		"google_container_registry":                                      296,
		"google_data_catalog_entry":                                      297,
		"google_data_catalog_entry_group":                                298,
		"google_data_catalog_entry_group_iam_binding":                    299,
		"google_data_catalog_entry_group_iam_member":                     300,
		"google_data_catalog_entry_group_iam_policy":                     301,
		"google_data_catalog_policy_tag":                                 302,
		"google_data_catalog_policy_tag_iam_binding":                     303,
		"google_data_catalog_policy_tag_iam_member":                      304,
		"google_data_catalog_policy_tag_iam_policy":                      305,
		"google_data_catalog_tag":                                        306,
		"google_data_catalog_tag_template":                               307,
		"google_data_catalog_tag_template_iam_binding":                   308,
		"google_data_catalog_tag_template_iam_member":                    309,
		"google_data_catalog_tag_template_iam_policy":                    310,
		"google_data_catalog_taxonomy":                                   311,
		"google_data_catalog_taxonomy_iam_binding":                       312,
		"google_data_catalog_taxonomy_iam_member":                        313,
		"google_data_catalog_taxonomy_iam_policy":                        314,
		"google_data_fusion_instance":                                    315,
		"google_data_fusion_instance_iam_binding":                        316,
		"google_data_fusion_instance_iam_member":                         317,
		"google_data_fusion_instance_iam_policy":                         318,
		"google_data_loss_prevention_deidentify_template":                319,
		"google_data_loss_prevention_inspect_template":                   320,
		"google_data_loss_prevention_job_trigger":                        321,
		"google_data_loss_prevention_stored_info_type":                   322,
		"google_data_pipeline_pipeline":                                  323,
		"google_database_migration_service_connection_profile":           324,
		"google_dataflow_job":                                            325,
		"google_dataplex_asset":                                          326,
		"google_dataplex_asset_iam_binding":                              327,
		"google_dataplex_asset_iam_member":                               328,
		"google_dataplex_asset_iam_policy":                               329,
		"google_dataplex_datascan":                                       330,
		"google_dataplex_datascan_iam_binding":                           331,
		"google_dataplex_datascan_iam_member":                            332,
		"google_dataplex_datascan_iam_policy":                            333,
		"google_dataplex_lake":                                           334,
		"google_dataplex_lake_iam_binding":                               335,
		"google_dataplex_lake_iam_member":                                336,
		"google_dataplex_lake_iam_policy":                                337,
		"google_dataplex_task":                                           338,
		"google_dataplex_task_iam_binding":                               339,
		"google_dataplex_task_iam_member":                                340,
		"google_dataplex_task_iam_policy":                                341,
		"google_dataplex_zone":                                           342,
		"google_dataplex_zone_iam_binding":                               343,
		"google_dataplex_zone_iam_member":                                344,
		"google_dataplex_zone_iam_policy":                                345,
		"google_dataproc_autoscaling_policy":                             346,
		"google_dataproc_autoscaling_policy_iam_binding":                 347,
		"google_dataproc_autoscaling_policy_iam_member":                  348,
		"google_dataproc_autoscaling_policy_iam_policy":                  349,
		"google_dataproc_cluster":                                        350,
		"google_dataproc_cluster_iam_binding":                            351,
		"google_dataproc_cluster_iam_member":                             352,
		"google_dataproc_cluster_iam_policy":                             353,
		"google_dataproc_job":                                            354,
		"google_dataproc_job_iam_binding":                                355,
		"google_dataproc_job_iam_member":                                 356,
		"google_dataproc_job_iam_policy":                                 357,
		"google_dataproc_metastore_service":                              358,
		"google_dataproc_metastore_service_iam_binding":                  359,
		"google_dataproc_metastore_service_iam_member":                   360,
		"google_dataproc_metastore_service_iam_policy":                   361,
		"google_dataproc_workflow_template":                              362,
		"google_datastore_index":                                         363,
		"google_datastream_connection_profile":                           364,
		"google_datastream_private_connection":                           365,
		"google_datastream_stream":                                       366,
		"google_deployment_manager_deployment":                           367,
		"google_dialogflow_agent":                                        368,
		"google_dialogflow_cx_agent":                                     369,
		"google_dialogflow_cx_entity_type":                               370,
		"google_dialogflow_cx_environment":                               371,
		"google_dialogflow_cx_flow":                                      372,
		"google_dialogflow_cx_intent":                                    373,
		"google_dialogflow_cx_page":                                      374,
		"google_dialogflow_cx_test_case":                                 375,
		"google_dialogflow_cx_version":                                   376,
		"google_dialogflow_cx_webhook":                                   377,
		"google_dialogflow_entity_type":                                  378,
		"google_dialogflow_fulfillment":                                  379,
		"google_dialogflow_intent":                                       380,
		"google_dns_managed_zone":                                        381,
		"google_dns_managed_zone_iam_binding":                            382,
		"google_dns_managed_zone_iam_member":                             383,
		"google_dns_managed_zone_iam_policy":                             384,
		"google_dns_policy":                                              385,
		"google_dns_record_set":                                          386,
		"google_dns_response_policy":                                     387,
		"google_dns_response_policy_rule":                                388,
		"google_document_ai_processor":                                   389,
		"google_document_ai_processor_default_version":                   390,
		"google_document_ai_warehouse_document_schema":                   391,
		"google_document_ai_warehouse_location":                          392,
		"google_endpoints_service":                                       393,
		"google_endpoints_service_consumers_iam_binding":                 394,
		"google_endpoints_service_consumers_iam_member":                  395,
		"google_endpoints_service_consumers_iam_policy":                  396,
		"google_endpoints_service_iam_binding":                           397,
		"google_endpoints_service_iam_member":                            398,
		"google_endpoints_service_iam_policy":                            399,
		"google_essential_contacts_contact":                              400,
		"google_eventarc_channel":                                        401,
		"google_eventarc_google_channel_config":                          402,
		"google_eventarc_trigger":                                        403,
		"google_filestore_backup":                                        404,
		"google_filestore_instance":                                      405,
		"google_filestore_snapshot":                                      406,
		"google_firebaserules_release":                                   407,
		"google_firebaserules_ruleset":                                   408,
		"google_firestore_database":                                      409,
		"google_firestore_document":                                      410,
		"google_firestore_field":                                         411,
		"google_firestore_index":                                         412,
		"google_folder":                                                  413,
		"google_folder_access_approval_settings":                         414,
		"google_folder_iam_audit_config":                                 415,
		"google_folder_iam_binding":                                      416,
		"google_folder_iam_member":                                       417,
		"google_folder_iam_policy":                                       418,
		"google_folder_organization_policy":                              419,
		"google_game_services_game_server_cluster":                       420,
		"google_game_services_game_server_config":                        421,
		"google_game_services_game_server_deployment":                    422,
		"google_game_services_game_server_deployment_rollout":            423,
		"google_game_services_realm":                                     424,
		"google_gke_backup_backup_plan":                                  425,
		"google_gke_backup_backup_plan_iam_binding":                      426,
		"google_gke_backup_backup_plan_iam_member":                       427,
		"google_gke_backup_backup_plan_iam_policy":                       428,
		"google_gke_hub_feature":                                         429,
		"google_gke_hub_feature_iam_binding":                             430,
		"google_gke_hub_feature_iam_member":                              431,
		"google_gke_hub_feature_iam_policy":                              432,
		"google_gke_hub_feature_membership":                              433,
		"google_gke_hub_membership":                                      434,
		"google_gke_hub_membership_binding":                              435,
		"google_gke_hub_membership_iam_binding":                          436,
		"google_gke_hub_membership_iam_member":                           437,
		"google_gke_hub_membership_iam_policy":                           438,
		"google_gke_hub_namespace":                                       439,
		"google_gke_hub_scope":                                           440,
		"google_gke_hub_scope_iam_binding":                               441,
		"google_gke_hub_scope_iam_member":                                442,
		"google_gke_hub_scope_iam_policy":                                443,
		"google_gke_hub_scope_rbac_role_binding":                         444,
		"google_healthcare_consent_store":                                445,
		"google_healthcare_consent_store_iam_binding":                    446,
		"google_healthcare_consent_store_iam_member":                     447,
		"google_healthcare_consent_store_iam_policy":                     448,
		"google_healthcare_dataset":                                      449,
		"google_healthcare_dataset_iam_binding":                          450,
		"google_healthcare_dataset_iam_member":                           451,
		"google_healthcare_dataset_iam_policy":                           452,
		"google_healthcare_dicom_store":                                  453,
		"google_healthcare_dicom_store_iam_binding":                      454,
		"google_healthcare_dicom_store_iam_member":                       455,
		"google_healthcare_dicom_store_iam_policy":                       456,
		"google_healthcare_fhir_store":                                   457,
		"google_healthcare_fhir_store_iam_binding":                       458,
		"google_healthcare_fhir_store_iam_member":                        459,
		"google_healthcare_fhir_store_iam_policy":                        460,
		"google_healthcare_hl7_v2_store":                                 461,
		"google_healthcare_hl7_v2_store_iam_binding":                     462,
		"google_healthcare_hl7_v2_store_iam_member":                      463,
		"google_healthcare_hl7_v2_store_iam_policy":                      464,
		"google_iam_access_boundary_policy":                              465,
		"google_iam_deny_policy":                                         466,
		"google_iam_workforce_pool":                                      467,
		"google_iam_workforce_pool_provider":                             468,
		"google_iam_workload_identity_pool":                              469,
		"google_iam_workload_identity_pool_provider":                     470,
		"google_iap_app_engine_service_iam_binding":                      471,
		"google_iap_app_engine_service_iam_member":                       472,
		"google_iap_app_engine_service_iam_policy":                       473,
		"google_iap_app_engine_version_iam_binding":                      474,
		"google_iap_app_engine_version_iam_member":                       475,
		"google_iap_app_engine_version_iam_policy":                       476,
		"google_iap_brand":                                               477,
		"google_iap_client":                                              478,
		"google_iap_tunnel_iam_binding":                                  479,
		"google_iap_tunnel_iam_member":                                   480,
		"google_iap_tunnel_iam_policy":                                   481,
		"google_iap_tunnel_instance_iam_binding":                         482,
		"google_iap_tunnel_instance_iam_member":                          483,
		"google_iap_tunnel_instance_iam_policy":                          484,
		"google_iap_web_backend_service_iam_binding":                     485,
		"google_iap_web_backend_service_iam_member":                      486,
		"google_iap_web_backend_service_iam_policy":                      487,
		"google_iap_web_iam_binding":                                     488,
		"google_iap_web_iam_member":                                      489,
		"google_iap_web_iam_policy":                                      490,
		"google_iap_web_region_backend_service_iam_binding":              491,
		"google_iap_web_region_backend_service_iam_member":               492,
		"google_iap_web_region_backend_service_iam_policy":               493,
		"google_iap_web_type_app_engine_iam_binding":                     494,
		"google_iap_web_type_app_engine_iam_member":                      495,
		"google_iap_web_type_app_engine_iam_policy":                      496,
		"google_iap_web_type_compute_iam_binding":                        497,
		"google_iap_web_type_compute_iam_member":                         498,
		"google_iap_web_type_compute_iam_policy":                         499,
		"google_identity_platform_config":                                500,
		"google_identity_platform_default_supported_idp_config":          501,
		"google_identity_platform_inbound_saml_config":                   502,
		"google_identity_platform_oauth_idp_config":                      503,
		"google_identity_platform_project_default_config":                504,
		"google_identity_platform_tenant":                                505,
		"google_identity_platform_tenant_default_supported_idp_config":   506,
		"google_identity_platform_tenant_inbound_saml_config":            507,
		"google_identity_platform_tenant_oauth_idp_config":               508,
		"google_kms_crypto_key":                                          509,
		"google_kms_crypto_key_iam_binding":                              510,
		"google_kms_crypto_key_iam_member":                               511,
		"google_kms_crypto_key_iam_policy":                               512,
		"google_kms_crypto_key_version":                                  513,
		"google_kms_key_ring":                                            514,
		"google_kms_key_ring_iam_binding":                                515,
		"google_kms_key_ring_iam_member":                                 516,
		"google_kms_key_ring_iam_policy":                                 517,
		"google_kms_key_ring_import_job":                                 518,
		"google_kms_secret_ciphertext":                                   519,
		"google_logging_billing_account_bucket_config":                   520,
		"google_logging_billing_account_exclusion":                       521,
		"google_logging_billing_account_sink":                            522,
		"google_logging_folder_bucket_config":                            523,
		"google_logging_folder_exclusion":                                524,
		"google_logging_folder_sink":                                     525,
		"google_logging_linked_dataset":                                  526,
		"google_logging_log_view":                                        527,
		"google_logging_metric":                                          528,
		"google_logging_organization_bucket_config":                      529,
		"google_logging_organization_exclusion":                          530,
		"google_logging_organization_sink":                               531,
		"google_logging_project_bucket_config":                           532,
		"google_logging_project_exclusion":                               533,
		"google_logging_project_sink":                                    534,
		"google_looker_instance":                                         535,
		"google_memcache_instance":                                       536,
		"google_ml_engine_model":                                         537,
		"google_monitoring_alert_policy":                                 538,
		"google_monitoring_custom_service":                               539,
		"google_monitoring_dashboard":                                    540,
		"google_monitoring_group":                                        541,
		"google_monitoring_metric_descriptor":                            542,
		"google_monitoring_monitored_project":                            543,
		"google_monitoring_notification_channel":                         544,
		"google_monitoring_service":                                      545,
		"google_monitoring_slo":                                          546,
		"google_monitoring_uptime_check_config":                          547,
		"google_network_connectivity_hub":                                548,
		"google_network_connectivity_service_connection_policy":          549,
		"google_network_connectivity_spoke":                              550,
		"google_network_management_connectivity_test":                    551,
		"google_network_security_address_group":                          552,
		"google_network_security_gateway_security_policy":                553,
		"google_network_security_gateway_security_policy_rule":           554,
		"google_network_security_url_lists":                              555,
		"google_network_services_edge_cache_keyset":                      556,
		"google_network_services_edge_cache_origin":                      557,
		"google_network_services_edge_cache_service":                     558,
		"google_network_services_gateway":                                559,
		"google_notebooks_environment":                                   560,
		"google_notebooks_instance":                                      561,
		"google_notebooks_instance_iam_binding":                          562,
		"google_notebooks_instance_iam_member":                           563,
		"google_notebooks_instance_iam_policy":                           564,
		"google_notebooks_location":                                      565,
		"google_notebooks_runtime":                                       566,
		"google_notebooks_runtime_iam_binding":                           567,
		"google_notebooks_runtime_iam_member":                            568,
		"google_notebooks_runtime_iam_policy":                            569,
		"google_org_policy_policy":                                       570,
		"google_organization_access_approval_settings":                   571,
		"google_organization_iam_audit_config":                           572,
		"google_organization_iam_binding":                                573,
		"google_organization_iam_custom_role":                            574,
		"google_organization_iam_member":                                 575,
		"google_organization_iam_policy":                                 576,
		"google_organization_policy":                                     577,
		"google_os_config_os_policy_assignment":                          578,
		"google_os_config_patch_deployment":                              579,
		"google_os_login_ssh_public_key":                                 580,
		"google_privateca_ca_pool":                                       581,
		"google_privateca_ca_pool_iam_binding":                           582,
		"google_privateca_ca_pool_iam_member":                            583,
		"google_privateca_ca_pool_iam_policy":                            584,
		"google_privateca_certificate":                                   585,
		"google_privateca_certificate_authority":                         586,
		"google_privateca_certificate_template":                          587,
		"google_privateca_certificate_template_iam_binding":              588,
		"google_privateca_certificate_template_iam_member":               589,
		"google_privateca_certificate_template_iam_policy":               590,
		"google_project":                                                 591,
		"google_project_access_approval_settings":                        592,
		"google_project_default_service_accounts":                        593,
		"google_project_iam_audit_config":                                594,
		"google_project_iam_binding":                                     595,
		"google_project_iam_custom_role":                                 596,
		"google_project_iam_member":                                      597,
		"google_project_iam_policy":                                      598,
		"google_project_organization_policy":                             599,
		"google_project_service":                                         600,
		"google_project_usage_export_bucket":                             601,
		"google_public_ca_external_account_key":                          602,
		"google_pubsub_lite_reservation":                                 603,
		"google_pubsub_lite_subscription":                                604,
		"google_pubsub_lite_topic":                                       605,
		"google_pubsub_schema":                                           606,
		"google_pubsub_subscription":                                     607,
		"google_pubsub_subscription_iam_binding":                         608,
		"google_pubsub_subscription_iam_member":                          609,
		"google_pubsub_subscription_iam_policy":                          610,
		"google_pubsub_topic":                                            611,
		"google_pubsub_topic_iam_binding":                                612,
		"google_pubsub_topic_iam_member":                                 613,
		"google_pubsub_topic_iam_policy":                                 614,
		"google_recaptcha_enterprise_key":                                615,
		"google_redis_instance":                                          616,
		"google_resource_manager_lien":                                   617,
		"google_scc_mute_config":                                         618,
		"google_scc_notification_config":                                 619,
		"google_scc_source":                                              620,
		"google_scc_source_iam_binding":                                  621,
		"google_scc_source_iam_member":                                   622,
		"google_scc_source_iam_policy":                                   623,
		"google_secret_manager_secret":                                   624,
		"google_secret_manager_secret_iam_binding":                       625,
		"google_secret_manager_secret_iam_member":                        626,
		"google_secret_manager_secret_iam_policy":                        627,
		"google_secret_manager_secret_version":                           628,
		"google_service_account":                                         629,
		"google_service_account_iam_binding":                             630,
		"google_service_account_iam_member":                              631,
		"google_service_account_iam_policy":                              632,
		"google_service_account_key":                                     633,
		"google_service_networking_connection":                           634,
		"google_service_networking_peered_dns_domain":                    635,
		"google_sourcerepo_repository":                                   636,
		"google_sourcerepo_repository_iam_binding":                       637,
		"google_sourcerepo_repository_iam_member":                        638,
		"google_sourcerepo_repository_iam_policy":                        639,
		"google_spanner_database":                                        640,
		"google_spanner_database_iam_binding":                            641,
		"google_spanner_database_iam_member":                             642,
		"google_spanner_database_iam_policy":                             643,
		"google_spanner_instance":                                        644,
		"google_spanner_instance_iam_binding":                            645,
		"google_spanner_instance_iam_member":                             646,
		"google_spanner_instance_iam_policy":                             647,
		"google_sql_database":                                            648,
		"google_sql_database_instance":                                   649,
		"google_sql_source_representation_instance":                      650,
		"google_sql_ssl_cert":                                            651,
		"google_sql_user":                                                652,
		"google_storage_bucket":                                          653,
		"google_storage_bucket_access_control":                           654,
		"google_storage_bucket_acl":                                      655,
		"google_storage_bucket_iam_binding":                              656,
		"google_storage_bucket_iam_member":                               657,
		"google_storage_bucket_iam_policy":                               658,
		"google_storage_bucket_object":                                   659,
		"google_storage_default_object_access_control":                   660,
		"google_storage_default_object_acl":                              661,
		"google_storage_hmac_key":                                        662,
		"google_storage_insights_report_config":                          663,
		"google_storage_notification":                                    664,
		"google_storage_object_access_control":                           665,
		"google_storage_object_acl":                                      666,
		"google_storage_transfer_agent_pool":                             667,
		"google_storage_transfer_job":                                    668,
		"google_tags_location_tag_binding":                               669,
		"google_tags_tag_binding":                                        670,
		"google_tags_tag_key":                                            671,
		"google_tags_tag_key_iam_binding":                                672,
		"google_tags_tag_key_iam_member":                                 673,
		"google_tags_tag_key_iam_policy":                                 674,
		"google_tags_tag_value":                                          675,
		"google_tags_tag_value_iam_binding":                              676,
		"google_tags_tag_value_iam_member":                               677,
		"google_tags_tag_value_iam_policy":                               678,
		"google_tpu_node":                                                679,
		"google_vertex_ai_dataset":                                       680,
		"google_vertex_ai_endpoint":                                      681,
		"google_vertex_ai_featurestore":                                  682,
		"google_vertex_ai_featurestore_entitytype":                       683,
		"google_vertex_ai_featurestore_entitytype_feature":               684,
		"google_vertex_ai_index":                                         685,
		"google_vertex_ai_index_endpoint":                                686,
		"google_vertex_ai_tensorboard":                                   687,
		"google_vpc_access_connector":                                    688,
		"google_workflows_workflow":                                      689,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
