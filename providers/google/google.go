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
		"google_apikeys_key":                                             42,
		"google_app_engine_application":                                  43,
		"google_app_engine_application_url_dispatch_rules":               44,
		"google_app_engine_domain_mapping":                               45,
		"google_app_engine_firewall_rule":                                46,
		"google_app_engine_flexible_app_version":                         47,
		"google_app_engine_service_network_settings":                     48,
		"google_app_engine_service_split_traffic":                        49,
		"google_app_engine_standard_app_version":                         50,
		"google_artifact_registry_repository":                            51,
		"google_artifact_registry_repository_iam_binding":                52,
		"google_artifact_registry_repository_iam_member":                 53,
		"google_artifact_registry_repository_iam_policy":                 54,
		"google_assured_workloads_workload":                              55,
		"google_beyondcorp_app_connection":                               56,
		"google_beyondcorp_app_connector":                                57,
		"google_beyondcorp_app_gateway":                                  58,
		"google_biglake_catalog":                                         59,
		"google_biglake_database":                                        60,
		"google_bigquery_analytics_hub_data_exchange":                    61,
		"google_bigquery_analytics_hub_data_exchange_iam_binding":        62,
		"google_bigquery_analytics_hub_data_exchange_iam_member":         63,
		"google_bigquery_analytics_hub_data_exchange_iam_policy":         64,
		"google_bigquery_analytics_hub_listing":                          65,
		"google_bigquery_analytics_hub_listing_iam_binding":              66,
		"google_bigquery_analytics_hub_listing_iam_member":               67,
		"google_bigquery_analytics_hub_listing_iam_policy":               68,
		"google_bigquery_bi_reservation":                                 69,
		"google_bigquery_capacity_commitment":                            70,
		"google_bigquery_connection":                                     71,
		"google_bigquery_connection_iam_binding":                         72,
		"google_bigquery_connection_iam_member":                          73,
		"google_bigquery_connection_iam_policy":                          74,
		"google_bigquery_data_transfer_config":                           75,
		"google_bigquery_datapolicy_data_policy":                         76,
		"google_bigquery_datapolicy_data_policy_iam_binding":             77,
		"google_bigquery_datapolicy_data_policy_iam_member":              78,
		"google_bigquery_datapolicy_data_policy_iam_policy":              79,
		"google_bigquery_dataset":                                        80,
		"google_bigquery_dataset_access":                                 81,
		"google_bigquery_dataset_iam_binding":                            82,
		"google_bigquery_dataset_iam_member":                             83,
		"google_bigquery_dataset_iam_policy":                             84,
		"google_bigquery_job":                                            85,
		"google_bigquery_reservation":                                    86,
		"google_bigquery_reservation_assignment":                         87,
		"google_bigquery_routine":                                        88,
		"google_bigquery_table":                                          89,
		"google_bigquery_table_iam_binding":                              90,
		"google_bigquery_table_iam_member":                               91,
		"google_bigquery_table_iam_policy":                               92,
		"google_bigtable_app_profile":                                    93,
		"google_bigtable_gc_policy":                                      94,
		"google_bigtable_instance":                                       95,
		"google_bigtable_instance_iam_binding":                           96,
		"google_bigtable_instance_iam_member":                            97,
		"google_bigtable_instance_iam_policy":                            98,
		"google_bigtable_table":                                          99,
		"google_bigtable_table_iam_binding":                              100,
		"google_bigtable_table_iam_member":                               101,
		"google_bigtable_table_iam_policy":                               102,
		"google_billing_account_iam_binding":                             103,
		"google_billing_account_iam_member":                              104,
		"google_billing_account_iam_policy":                              105,
		"google_billing_budget":                                          106,
		"google_billing_project_info":                                    107,
		"google_billing_subaccount":                                      108,
		"google_binary_authorization_attestor":                           109,
		"google_binary_authorization_attestor_iam_binding":               110,
		"google_binary_authorization_attestor_iam_member":                111,
		"google_binary_authorization_attestor_iam_policy":                112,
		"google_binary_authorization_policy":                             113,
		"google_certificate_manager_certificate":                         114,
		"google_certificate_manager_certificate_issuance_config":         115,
		"google_certificate_manager_certificate_map":                     116,
		"google_certificate_manager_certificate_map_entry":               117,
		"google_certificate_manager_dns_authorization":                   118,
		"google_certificate_manager_trust_config":                        119,
		"google_cloud_asset_folder_feed":                                 120,
		"google_cloud_asset_organization_feed":                           121,
		"google_cloud_asset_project_feed":                                122,
		"google_cloud_identity_group":                                    123,
		"google_cloud_identity_group_membership":                         124,
		"google_cloud_ids_endpoint":                                      125,
		"google_cloud_run_domain_mapping":                                126,
		"google_cloud_run_service":                                       127,
		"google_cloud_run_service_iam_binding":                           128,
		"google_cloud_run_service_iam_member":                            129,
		"google_cloud_run_service_iam_policy":                            130,
		"google_cloud_run_v2_job":                                        131,
		"google_cloud_run_v2_job_iam_binding":                            132,
		"google_cloud_run_v2_job_iam_member":                             133,
		"google_cloud_run_v2_job_iam_policy":                             134,
		"google_cloud_run_v2_service":                                    135,
		"google_cloud_run_v2_service_iam_binding":                        136,
		"google_cloud_run_v2_service_iam_member":                         137,
		"google_cloud_run_v2_service_iam_policy":                         138,
		"google_cloud_scheduler_job":                                     139,
		"google_cloud_tasks_queue":                                       140,
		"google_cloud_tasks_queue_iam_binding":                           141,
		"google_cloud_tasks_queue_iam_member":                            142,
		"google_cloud_tasks_queue_iam_policy":                            143,
		"google_cloudbuild_bitbucket_server_config":                      144,
		"google_cloudbuild_trigger":                                      145,
		"google_cloudbuild_worker_pool":                                  146,
		"google_cloudbuildv2_connection":                                 147,
		"google_cloudbuildv2_connection_iam_binding":                     148,
		"google_cloudbuildv2_connection_iam_member":                      149,
		"google_cloudbuildv2_connection_iam_policy":                      150,
		"google_cloudbuildv2_repository":                                 151,
		"google_clouddeploy_delivery_pipeline":                           152,
		"google_clouddeploy_target":                                      153,
		"google_cloudfunctions2_function":                                154,
		"google_cloudfunctions2_function_iam_binding":                    155,
		"google_cloudfunctions2_function_iam_member":                     156,
		"google_cloudfunctions2_function_iam_policy":                     157,
		"google_cloudfunctions_function":                                 158,
		"google_cloudfunctions_function_iam_binding":                     159,
		"google_cloudfunctions_function_iam_member":                      160,
		"google_cloudfunctions_function_iam_policy":                      161,
		"google_cloudiot_device":                                         162,
		"google_cloudiot_registry":                                       163,
		"google_cloudiot_registry_iam_binding":                           164,
		"google_cloudiot_registry_iam_member":                            165,
		"google_cloudiot_registry_iam_policy":                            166,
		"google_composer_environment":                                    167,
		"google_compute_address":                                         168,
		"google_compute_attached_disk":                                   169,
		"google_compute_autoscaler":                                      170,
		"google_compute_backend_bucket":                                  171,
		"google_compute_backend_bucket_signed_url_key":                   172,
		"google_compute_backend_service":                                 173,
		"google_compute_backend_service_signed_url_key":                  174,
		"google_compute_disk":                                            175,
		"google_compute_disk_async_replication":                          176,
		"google_compute_disk_iam_binding":                                177,
		"google_compute_disk_iam_member":                                 178,
		"google_compute_disk_iam_policy":                                 179,
		"google_compute_disk_resource_policy_attachment":                 180,
		"google_compute_external_vpn_gateway":                            181,
		"google_compute_firewall":                                        182,
		"google_compute_firewall_policy":                                 183,
		"google_compute_firewall_policy_association":                     184,
		"google_compute_firewall_policy_rule":                            185,
		"google_compute_forwarding_rule":                                 186,
		"google_compute_global_address":                                  187,
		"google_compute_global_forwarding_rule":                          188,
		"google_compute_global_network_endpoint":                         189,
		"google_compute_global_network_endpoint_group":                   190,
		"google_compute_ha_vpn_gateway":                                  191,
		"google_compute_health_check":                                    192,
		"google_compute_http_health_check":                               193,
		"google_compute_https_health_check":                              194,
		"google_compute_image":                                           195,
		"google_compute_image_iam_binding":                               196,
		"google_compute_image_iam_member":                                197,
		"google_compute_image_iam_policy":                                198,
		"google_compute_instance":                                        199,
		"google_compute_instance_from_template":                          200,
		"google_compute_instance_group":                                  201,
		"google_compute_instance_group_manager":                          202,
		"google_compute_instance_group_named_port":                       203,
		"google_compute_instance_iam_binding":                            204,
		"google_compute_instance_iam_member":                             205,
		"google_compute_instance_iam_policy":                             206,
		"google_compute_instance_template":                               207,
		"google_compute_interconnect_attachment":                         208,
		"google_compute_managed_ssl_certificate":                         209,
		"google_compute_network":                                         210,
		"google_compute_network_endpoint":                                211,
		"google_compute_network_endpoint_group":                          212,
		"google_compute_network_endpoints":                               213,
		"google_compute_network_firewall_policy":                         214,
		"google_compute_network_firewall_policy_association":             215,
		"google_compute_network_firewall_policy_rule":                    216,
		"google_compute_network_peering":                                 217,
		"google_compute_network_peering_routes_config":                   218,
		"google_compute_node_group":                                      219,
		"google_compute_node_template":                                   220,
		"google_compute_packet_mirroring":                                221,
		"google_compute_per_instance_config":                             222,
		"google_compute_project_default_network_tier":                    223,
		"google_compute_project_metadata":                                224,
		"google_compute_project_metadata_item":                           225,
		"google_compute_public_advertised_prefix":                        226,
		"google_compute_public_delegated_prefix":                         227,
		"google_compute_region_autoscaler":                               228,
		"google_compute_region_backend_service":                          229,
		"google_compute_region_commitment":                               230,
		"google_compute_region_disk":                                     231,
		"google_compute_region_disk_iam_binding":                         232,
		"google_compute_region_disk_iam_member":                          233,
		"google_compute_region_disk_iam_policy":                          234,
		"google_compute_region_disk_resource_policy_attachment":          235,
		"google_compute_region_health_check":                             236,
		"google_compute_region_instance_group_manager":                   237,
		"google_compute_region_instance_template":                        238,
		"google_compute_region_network_endpoint_group":                   239,
		"google_compute_region_network_firewall_policy":                  240,
		"google_compute_region_network_firewall_policy_association":      241,
		"google_compute_region_network_firewall_policy_rule":             242,
		"google_compute_region_per_instance_config":                      243,
		"google_compute_region_ssl_certificate":                          244,
		"google_compute_region_ssl_policy":                               245,
		"google_compute_region_target_http_proxy":                        246,
		"google_compute_region_target_https_proxy":                       247,
		"google_compute_region_target_tcp_proxy":                         248,
		"google_compute_region_url_map":                                  249,
		"google_compute_reservation":                                     250,
		"google_compute_resource_policy":                                 251,
		"google_compute_route":                                           252,
		"google_compute_router":                                          253,
		"google_compute_router_interface":                                254,
		"google_compute_router_nat":                                      255,
		"google_compute_router_peer":                                     256,
		"google_compute_security_policy":                                 257,
		"google_compute_service_attachment":                              258,
		"google_compute_shared_vpc_host_project":                         259,
		"google_compute_shared_vpc_service_project":                      260,
		"google_compute_snapshot":                                        261,
		"google_compute_snapshot_iam_binding":                            262,
		"google_compute_snapshot_iam_member":                             263,
		"google_compute_snapshot_iam_policy":                             264,
		"google_compute_ssl_certificate":                                 265,
		"google_compute_ssl_policy":                                      266,
		"google_compute_subnetwork":                                      267,
		"google_compute_subnetwork_iam_binding":                          268,
		"google_compute_subnetwork_iam_member":                           269,
		"google_compute_subnetwork_iam_policy":                           270,
		"google_compute_target_grpc_proxy":                               271,
		"google_compute_target_http_proxy":                               272,
		"google_compute_target_https_proxy":                              273,
		"google_compute_target_instance":                                 274,
		"google_compute_target_pool":                                     275,
		"google_compute_target_ssl_proxy":                                276,
		"google_compute_target_tcp_proxy":                                277,
		"google_compute_url_map":                                         278,
		"google_compute_vpn_gateway":                                     279,
		"google_compute_vpn_tunnel":                                      280,
		"google_container_analysis_note":                                 281,
		"google_container_analysis_note_iam_binding":                     282,
		"google_container_analysis_note_iam_member":                      283,
		"google_container_analysis_note_iam_policy":                      284,
		"google_container_analysis_occurrence":                           285,
		"google_container_attached_cluster":                              286,
		"google_container_aws_cluster":                                   287,
		"google_container_aws_node_pool":                                 288,
		"google_container_azure_client":                                  289,
		"google_container_azure_cluster":                                 290,
		"google_container_azure_node_pool":                               291,
		"google_container_cluster":                                       292,
		"google_container_node_pool":                                     293,
		"google_container_registry":                                      294,
		"google_data_catalog_entry":                                      295,
		"google_data_catalog_entry_group":                                296,
		"google_data_catalog_entry_group_iam_binding":                    297,
		"google_data_catalog_entry_group_iam_member":                     298,
		"google_data_catalog_entry_group_iam_policy":                     299,
		"google_data_catalog_policy_tag":                                 300,
		"google_data_catalog_policy_tag_iam_binding":                     301,
		"google_data_catalog_policy_tag_iam_member":                      302,
		"google_data_catalog_policy_tag_iam_policy":                      303,
		"google_data_catalog_tag":                                        304,
		"google_data_catalog_tag_template":                               305,
		"google_data_catalog_tag_template_iam_binding":                   306,
		"google_data_catalog_tag_template_iam_member":                    307,
		"google_data_catalog_tag_template_iam_policy":                    308,
		"google_data_catalog_taxonomy":                                   309,
		"google_data_catalog_taxonomy_iam_binding":                       310,
		"google_data_catalog_taxonomy_iam_member":                        311,
		"google_data_catalog_taxonomy_iam_policy":                        312,
		"google_data_fusion_instance":                                    313,
		"google_data_fusion_instance_iam_binding":                        314,
		"google_data_fusion_instance_iam_member":                         315,
		"google_data_fusion_instance_iam_policy":                         316,
		"google_data_loss_prevention_deidentify_template":                317,
		"google_data_loss_prevention_inspect_template":                   318,
		"google_data_loss_prevention_job_trigger":                        319,
		"google_data_loss_prevention_stored_info_type":                   320,
		"google_database_migration_service_connection_profile":           321,
		"google_dataflow_job":                                            322,
		"google_dataplex_asset":                                          323,
		"google_dataplex_asset_iam_binding":                              324,
		"google_dataplex_asset_iam_member":                               325,
		"google_dataplex_asset_iam_policy":                               326,
		"google_dataplex_datascan":                                       327,
		"google_dataplex_datascan_iam_binding":                           328,
		"google_dataplex_datascan_iam_member":                            329,
		"google_dataplex_datascan_iam_policy":                            330,
		"google_dataplex_lake":                                           331,
		"google_dataplex_lake_iam_binding":                               332,
		"google_dataplex_lake_iam_member":                                333,
		"google_dataplex_lake_iam_policy":                                334,
		"google_dataplex_task":                                           335,
		"google_dataplex_task_iam_binding":                               336,
		"google_dataplex_task_iam_member":                                337,
		"google_dataplex_task_iam_policy":                                338,
		"google_dataplex_zone":                                           339,
		"google_dataplex_zone_iam_binding":                               340,
		"google_dataplex_zone_iam_member":                                341,
		"google_dataplex_zone_iam_policy":                                342,
		"google_dataproc_autoscaling_policy":                             343,
		"google_dataproc_autoscaling_policy_iam_binding":                 344,
		"google_dataproc_autoscaling_policy_iam_member":                  345,
		"google_dataproc_autoscaling_policy_iam_policy":                  346,
		"google_dataproc_cluster":                                        347,
		"google_dataproc_cluster_iam_binding":                            348,
		"google_dataproc_cluster_iam_member":                             349,
		"google_dataproc_cluster_iam_policy":                             350,
		"google_dataproc_job":                                            351,
		"google_dataproc_job_iam_binding":                                352,
		"google_dataproc_job_iam_member":                                 353,
		"google_dataproc_job_iam_policy":                                 354,
		"google_dataproc_metastore_service":                              355,
		"google_dataproc_metastore_service_iam_binding":                  356,
		"google_dataproc_metastore_service_iam_member":                   357,
		"google_dataproc_metastore_service_iam_policy":                   358,
		"google_dataproc_workflow_template":                              359,
		"google_datastore_index":                                         360,
		"google_datastream_connection_profile":                           361,
		"google_datastream_private_connection":                           362,
		"google_datastream_stream":                                       363,
		"google_deployment_manager_deployment":                           364,
		"google_dialogflow_agent":                                        365,
		"google_dialogflow_cx_agent":                                     366,
		"google_dialogflow_cx_entity_type":                               367,
		"google_dialogflow_cx_environment":                               368,
		"google_dialogflow_cx_flow":                                      369,
		"google_dialogflow_cx_intent":                                    370,
		"google_dialogflow_cx_page":                                      371,
		"google_dialogflow_cx_version":                                   372,
		"google_dialogflow_cx_webhook":                                   373,
		"google_dialogflow_entity_type":                                  374,
		"google_dialogflow_fulfillment":                                  375,
		"google_dialogflow_intent":                                       376,
		"google_dns_managed_zone":                                        377,
		"google_dns_managed_zone_iam_binding":                            378,
		"google_dns_managed_zone_iam_member":                             379,
		"google_dns_managed_zone_iam_policy":                             380,
		"google_dns_policy":                                              381,
		"google_dns_record_set":                                          382,
		"google_dns_response_policy":                                     383,
		"google_dns_response_policy_rule":                                384,
		"google_document_ai_processor":                                   385,
		"google_document_ai_processor_default_version":                   386,
		"google_document_ai_warehouse_document_schema":                   387,
		"google_document_ai_warehouse_location":                          388,
		"google_endpoints_service":                                       389,
		"google_endpoints_service_consumers_iam_binding":                 390,
		"google_endpoints_service_consumers_iam_member":                  391,
		"google_endpoints_service_consumers_iam_policy":                  392,
		"google_endpoints_service_iam_binding":                           393,
		"google_endpoints_service_iam_member":                            394,
		"google_endpoints_service_iam_policy":                            395,
		"google_essential_contacts_contact":                              396,
		"google_eventarc_channel":                                        397,
		"google_eventarc_google_channel_config":                          398,
		"google_eventarc_trigger":                                        399,
		"google_filestore_backup":                                        400,
		"google_filestore_instance":                                      401,
		"google_filestore_snapshot":                                      402,
		"google_firebaserules_release":                                   403,
		"google_firebaserules_ruleset":                                   404,
		"google_firestore_database":                                      405,
		"google_firestore_document":                                      406,
		"google_firestore_field":                                         407,
		"google_firestore_index":                                         408,
		"google_folder":                                                  409,
		"google_folder_access_approval_settings":                         410,
		"google_folder_iam_audit_config":                                 411,
		"google_folder_iam_binding":                                      412,
		"google_folder_iam_member":                                       413,
		"google_folder_iam_policy":                                       414,
		"google_folder_organization_policy":                              415,
		"google_game_services_game_server_cluster":                       416,
		"google_game_services_game_server_config":                        417,
		"google_game_services_game_server_deployment":                    418,
		"google_game_services_game_server_deployment_rollout":            419,
		"google_game_services_realm":                                     420,
		"google_gke_backup_backup_plan":                                  421,
		"google_gke_backup_backup_plan_iam_binding":                      422,
		"google_gke_backup_backup_plan_iam_member":                       423,
		"google_gke_backup_backup_plan_iam_policy":                       424,
		"google_gke_hub_feature":                                         425,
		"google_gke_hub_feature_iam_binding":                             426,
		"google_gke_hub_feature_iam_member":                              427,
		"google_gke_hub_feature_iam_policy":                              428,
		"google_gke_hub_feature_membership":                              429,
		"google_gke_hub_membership":                                      430,
		"google_gke_hub_membership_binding":                              431,
		"google_gke_hub_membership_iam_binding":                          432,
		"google_gke_hub_membership_iam_member":                           433,
		"google_gke_hub_membership_iam_policy":                           434,
		"google_gke_hub_namespace":                                       435,
		"google_gke_hub_scope":                                           436,
		"google_gke_hub_scope_iam_binding":                               437,
		"google_gke_hub_scope_iam_member":                                438,
		"google_gke_hub_scope_iam_policy":                                439,
		"google_gke_hub_scope_rbac_role_binding":                         440,
		"google_healthcare_consent_store":                                441,
		"google_healthcare_consent_store_iam_binding":                    442,
		"google_healthcare_consent_store_iam_member":                     443,
		"google_healthcare_consent_store_iam_policy":                     444,
		"google_healthcare_dataset":                                      445,
		"google_healthcare_dataset_iam_binding":                          446,
		"google_healthcare_dataset_iam_member":                           447,
		"google_healthcare_dataset_iam_policy":                           448,
		"google_healthcare_dicom_store":                                  449,
		"google_healthcare_dicom_store_iam_binding":                      450,
		"google_healthcare_dicom_store_iam_member":                       451,
		"google_healthcare_dicom_store_iam_policy":                       452,
		"google_healthcare_fhir_store":                                   453,
		"google_healthcare_fhir_store_iam_binding":                       454,
		"google_healthcare_fhir_store_iam_member":                        455,
		"google_healthcare_fhir_store_iam_policy":                        456,
		"google_healthcare_hl7_v2_store":                                 457,
		"google_healthcare_hl7_v2_store_iam_binding":                     458,
		"google_healthcare_hl7_v2_store_iam_member":                      459,
		"google_healthcare_hl7_v2_store_iam_policy":                      460,
		"google_iam_access_boundary_policy":                              461,
		"google_iam_deny_policy":                                         462,
		"google_iam_workforce_pool":                                      463,
		"google_iam_workforce_pool_provider":                             464,
		"google_iam_workload_identity_pool":                              465,
		"google_iam_workload_identity_pool_provider":                     466,
		"google_iap_app_engine_service_iam_binding":                      467,
		"google_iap_app_engine_service_iam_member":                       468,
		"google_iap_app_engine_service_iam_policy":                       469,
		"google_iap_app_engine_version_iam_binding":                      470,
		"google_iap_app_engine_version_iam_member":                       471,
		"google_iap_app_engine_version_iam_policy":                       472,
		"google_iap_brand":                                               473,
		"google_iap_client":                                              474,
		"google_iap_tunnel_iam_binding":                                  475,
		"google_iap_tunnel_iam_member":                                   476,
		"google_iap_tunnel_iam_policy":                                   477,
		"google_iap_tunnel_instance_iam_binding":                         478,
		"google_iap_tunnel_instance_iam_member":                          479,
		"google_iap_tunnel_instance_iam_policy":                          480,
		"google_iap_web_backend_service_iam_binding":                     481,
		"google_iap_web_backend_service_iam_member":                      482,
		"google_iap_web_backend_service_iam_policy":                      483,
		"google_iap_web_iam_binding":                                     484,
		"google_iap_web_iam_member":                                      485,
		"google_iap_web_iam_policy":                                      486,
		"google_iap_web_region_backend_service_iam_binding":              487,
		"google_iap_web_region_backend_service_iam_member":               488,
		"google_iap_web_region_backend_service_iam_policy":               489,
		"google_iap_web_type_app_engine_iam_binding":                     490,
		"google_iap_web_type_app_engine_iam_member":                      491,
		"google_iap_web_type_app_engine_iam_policy":                      492,
		"google_iap_web_type_compute_iam_binding":                        493,
		"google_iap_web_type_compute_iam_member":                         494,
		"google_iap_web_type_compute_iam_policy":                         495,
		"google_identity_platform_config":                                496,
		"google_identity_platform_default_supported_idp_config":          497,
		"google_identity_platform_inbound_saml_config":                   498,
		"google_identity_platform_oauth_idp_config":                      499,
		"google_identity_platform_project_default_config":                500,
		"google_identity_platform_tenant":                                501,
		"google_identity_platform_tenant_default_supported_idp_config":   502,
		"google_identity_platform_tenant_inbound_saml_config":            503,
		"google_identity_platform_tenant_oauth_idp_config":               504,
		"google_kms_crypto_key":                                          505,
		"google_kms_crypto_key_iam_binding":                              506,
		"google_kms_crypto_key_iam_member":                               507,
		"google_kms_crypto_key_iam_policy":                               508,
		"google_kms_crypto_key_version":                                  509,
		"google_kms_key_ring":                                            510,
		"google_kms_key_ring_iam_binding":                                511,
		"google_kms_key_ring_iam_member":                                 512,
		"google_kms_key_ring_iam_policy":                                 513,
		"google_kms_key_ring_import_job":                                 514,
		"google_kms_secret_ciphertext":                                   515,
		"google_logging_billing_account_bucket_config":                   516,
		"google_logging_billing_account_exclusion":                       517,
		"google_logging_billing_account_sink":                            518,
		"google_logging_folder_bucket_config":                            519,
		"google_logging_folder_exclusion":                                520,
		"google_logging_folder_sink":                                     521,
		"google_logging_linked_dataset":                                  522,
		"google_logging_log_view":                                        523,
		"google_logging_metric":                                          524,
		"google_logging_organization_bucket_config":                      525,
		"google_logging_organization_exclusion":                          526,
		"google_logging_organization_sink":                               527,
		"google_logging_project_bucket_config":                           528,
		"google_logging_project_exclusion":                               529,
		"google_logging_project_sink":                                    530,
		"google_looker_instance":                                         531,
		"google_memcache_instance":                                       532,
		"google_ml_engine_model":                                         533,
		"google_monitoring_alert_policy":                                 534,
		"google_monitoring_custom_service":                               535,
		"google_monitoring_dashboard":                                    536,
		"google_monitoring_group":                                        537,
		"google_monitoring_metric_descriptor":                            538,
		"google_monitoring_monitored_project":                            539,
		"google_monitoring_notification_channel":                         540,
		"google_monitoring_service":                                      541,
		"google_monitoring_slo":                                          542,
		"google_monitoring_uptime_check_config":                          543,
		"google_network_connectivity_hub":                                544,
		"google_network_connectivity_service_connection_policy":          545,
		"google_network_connectivity_spoke":                              546,
		"google_network_management_connectivity_test":                    547,
		"google_network_security_address_group":                          548,
		"google_network_security_gateway_security_policy":                549,
		"google_network_security_gateway_security_policy_rule":           550,
		"google_network_security_url_lists":                              551,
		"google_network_services_edge_cache_keyset":                      552,
		"google_network_services_edge_cache_origin":                      553,
		"google_network_services_edge_cache_service":                     554,
		"google_network_services_gateway":                                555,
		"google_notebooks_environment":                                   556,
		"google_notebooks_instance":                                      557,
		"google_notebooks_instance_iam_binding":                          558,
		"google_notebooks_instance_iam_member":                           559,
		"google_notebooks_instance_iam_policy":                           560,
		"google_notebooks_location":                                      561,
		"google_notebooks_runtime":                                       562,
		"google_notebooks_runtime_iam_binding":                           563,
		"google_notebooks_runtime_iam_member":                            564,
		"google_notebooks_runtime_iam_policy":                            565,
		"google_org_policy_policy":                                       566,
		"google_organization_access_approval_settings":                   567,
		"google_organization_iam_audit_config":                           568,
		"google_organization_iam_binding":                                569,
		"google_organization_iam_custom_role":                            570,
		"google_organization_iam_member":                                 571,
		"google_organization_iam_policy":                                 572,
		"google_organization_policy":                                     573,
		"google_os_config_os_policy_assignment":                          574,
		"google_os_config_patch_deployment":                              575,
		"google_os_login_ssh_public_key":                                 576,
		"google_privateca_ca_pool":                                       577,
		"google_privateca_ca_pool_iam_binding":                           578,
		"google_privateca_ca_pool_iam_member":                            579,
		"google_privateca_ca_pool_iam_policy":                            580,
		"google_privateca_certificate":                                   581,
		"google_privateca_certificate_authority":                         582,
		"google_privateca_certificate_template":                          583,
		"google_privateca_certificate_template_iam_binding":              584,
		"google_privateca_certificate_template_iam_member":               585,
		"google_privateca_certificate_template_iam_policy":               586,
		"google_project":                                                 587,
		"google_project_access_approval_settings":                        588,
		"google_project_default_service_accounts":                        589,
		"google_project_iam_audit_config":                                590,
		"google_project_iam_binding":                                     591,
		"google_project_iam_custom_role":                                 592,
		"google_project_iam_member":                                      593,
		"google_project_iam_policy":                                      594,
		"google_project_organization_policy":                             595,
		"google_project_service":                                         596,
		"google_project_usage_export_bucket":                             597,
		"google_public_ca_external_account_key":                          598,
		"google_pubsub_lite_reservation":                                 599,
		"google_pubsub_lite_subscription":                                600,
		"google_pubsub_lite_topic":                                       601,
		"google_pubsub_schema":                                           602,
		"google_pubsub_subscription":                                     603,
		"google_pubsub_subscription_iam_binding":                         604,
		"google_pubsub_subscription_iam_member":                          605,
		"google_pubsub_subscription_iam_policy":                          606,
		"google_pubsub_topic":                                            607,
		"google_pubsub_topic_iam_binding":                                608,
		"google_pubsub_topic_iam_member":                                 609,
		"google_pubsub_topic_iam_policy":                                 610,
		"google_recaptcha_enterprise_key":                                611,
		"google_redis_instance":                                          612,
		"google_resource_manager_lien":                                   613,
		"google_scc_mute_config":                                         614,
		"google_scc_notification_config":                                 615,
		"google_scc_source":                                              616,
		"google_scc_source_iam_binding":                                  617,
		"google_scc_source_iam_member":                                   618,
		"google_scc_source_iam_policy":                                   619,
		"google_secret_manager_secret":                                   620,
		"google_secret_manager_secret_iam_binding":                       621,
		"google_secret_manager_secret_iam_member":                        622,
		"google_secret_manager_secret_iam_policy":                        623,
		"google_secret_manager_secret_version":                           624,
		"google_service_account":                                         625,
		"google_service_account_iam_binding":                             626,
		"google_service_account_iam_member":                              627,
		"google_service_account_iam_policy":                              628,
		"google_service_account_key":                                     629,
		"google_service_networking_connection":                           630,
		"google_service_networking_peered_dns_domain":                    631,
		"google_sourcerepo_repository":                                   632,
		"google_sourcerepo_repository_iam_binding":                       633,
		"google_sourcerepo_repository_iam_member":                        634,
		"google_sourcerepo_repository_iam_policy":                        635,
		"google_spanner_database":                                        636,
		"google_spanner_database_iam_binding":                            637,
		"google_spanner_database_iam_member":                             638,
		"google_spanner_database_iam_policy":                             639,
		"google_spanner_instance":                                        640,
		"google_spanner_instance_iam_binding":                            641,
		"google_spanner_instance_iam_member":                             642,
		"google_spanner_instance_iam_policy":                             643,
		"google_sql_database":                                            644,
		"google_sql_database_instance":                                   645,
		"google_sql_source_representation_instance":                      646,
		"google_sql_ssl_cert":                                            647,
		"google_sql_user":                                                648,
		"google_storage_bucket":                                          649,
		"google_storage_bucket_access_control":                           650,
		"google_storage_bucket_acl":                                      651,
		"google_storage_bucket_iam_binding":                              652,
		"google_storage_bucket_iam_member":                               653,
		"google_storage_bucket_iam_policy":                               654,
		"google_storage_bucket_object":                                   655,
		"google_storage_default_object_access_control":                   656,
		"google_storage_default_object_acl":                              657,
		"google_storage_hmac_key":                                        658,
		"google_storage_notification":                                    659,
		"google_storage_object_access_control":                           660,
		"google_storage_object_acl":                                      661,
		"google_storage_transfer_agent_pool":                             662,
		"google_storage_transfer_job":                                    663,
		"google_tags_location_tag_binding":                               664,
		"google_tags_tag_binding":                                        665,
		"google_tags_tag_key":                                            666,
		"google_tags_tag_key_iam_binding":                                667,
		"google_tags_tag_key_iam_member":                                 668,
		"google_tags_tag_key_iam_policy":                                 669,
		"google_tags_tag_value":                                          670,
		"google_tags_tag_value_iam_binding":                              671,
		"google_tags_tag_value_iam_member":                               672,
		"google_tags_tag_value_iam_policy":                               673,
		"google_tpu_node":                                                674,
		"google_vertex_ai_dataset":                                       675,
		"google_vertex_ai_endpoint":                                      676,
		"google_vertex_ai_featurestore":                                  677,
		"google_vertex_ai_featurestore_entitytype":                       678,
		"google_vertex_ai_featurestore_entitytype_feature":               679,
		"google_vertex_ai_index":                                         680,
		"google_vertex_ai_index_endpoint":                                681,
		"google_vertex_ai_tensorboard":                                   682,
		"google_vpc_access_connector":                                    683,
		"google_workflows_workflow":                                      684,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
