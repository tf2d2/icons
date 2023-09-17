package aws

import (
	"fmt"

	"github.com/tf2d2/icons/internal/resource"
)

var (
	Resources = []*resource.Resource{
		{
			Name: "azurerm_aadb2c_directory",
		},
		{
			Name: "azurerm_active_directory_domain_service",
		},
		{
			Name: "azurerm_active_directory_domain_service_replica_set",
		},
		{
			Name: "azurerm_active_directory_domain_service_trust",
		},
		{
			Name: "azurerm_advanced_threat_protection",
		},
		{
			Name: "azurerm_analysis_services_server",
		},
		{
			Name: "azurerm_api_connection",
		},
		{
			Name: "azurerm_api_management",
		},
		{
			Name: "azurerm_api_management_api",
		},
		{
			Name: "azurerm_api_management_api_diagnostic",
		},
		{
			Name: "azurerm_api_management_api_operation",
		},
		{
			Name: "azurerm_api_management_api_operation_policy",
		},
		{
			Name: "azurerm_api_management_api_operation_tag",
		},
		{
			Name: "azurerm_api_management_api_policy",
		},
		{
			Name: "azurerm_api_management_api_release",
		},
		{
			Name: "azurerm_api_management_api_schema",
		},
		{
			Name: "azurerm_api_management_api_tag",
		},
		{
			Name: "azurerm_api_management_api_tag_description",
		},
		{
			Name: "azurerm_api_management_api_version_set",
		},
		{
			Name: "azurerm_api_management_authorization_server",
		},
		{
			Name: "azurerm_api_management_backend",
		},
		{
			Name: "azurerm_api_management_certificate",
		},
		{
			Name: "azurerm_api_management_custom_domain",
		},
		{
			Name: "azurerm_api_management_diagnostic",
		},
		{
			Name: "azurerm_api_management_email_template",
		},
		{
			Name: "azurerm_api_management_gateway",
		},
		{
			Name: "azurerm_api_management_gateway_api",
		},
		{
			Name: "azurerm_api_management_gateway_certificate_authority",
		},
		{
			Name: "azurerm_api_management_gateway_host_name_configuration",
		},
		{
			Name: "azurerm_api_management_global_schema",
		},
		{
			Name: "azurerm_api_management_group",
		},
		{
			Name: "azurerm_api_management_group_user",
		},
		{
			Name: "azurerm_api_management_identity_provider_aad",
		},
		{
			Name: "azurerm_api_management_identity_provider_aadb2c",
		},
		{
			Name: "azurerm_api_management_identity_provider_facebook",
		},
		{
			Name: "azurerm_api_management_identity_provider_google",
		},
		{
			Name: "azurerm_api_management_identity_provider_microsoft",
		},
		{
			Name: "azurerm_api_management_identity_provider_twitter",
		},
		{
			Name: "azurerm_api_management_logger",
		},
		{
			Name: "azurerm_api_management_named_value",
		},
		{
			Name: "azurerm_api_management_notification_recipient_email",
		},
		{
			Name: "azurerm_api_management_notification_recipient_user",
		},
		{
			Name: "azurerm_api_management_openid_connect_provider",
		},
		{
			Name: "azurerm_api_management_policy",
		},
		{
			Name: "azurerm_api_management_product",
		},
		{
			Name: "azurerm_api_management_product_api",
		},
		{
			Name: "azurerm_api_management_product_group",
		},
		{
			Name: "azurerm_api_management_product_policy",
		},
		{
			Name: "azurerm_api_management_product_tag",
		},
		{
			Name: "azurerm_api_management_redis_cache",
		},
		{
			Name: "azurerm_api_management_subscription",
		},
		{
			Name: "azurerm_api_management_tag",
		},
		{
			Name: "azurerm_api_management_user",
		},
		{
			Name: "azurerm_app_configuration",
		},
		{
			Name: "azurerm_app_configuration_feature",
		},
		{
			Name: "azurerm_app_configuration_key",
		},
		{
			Name: "azurerm_app_service",
		},
		{
			Name: "azurerm_app_service_active_slot",
		},
		{
			Name: "azurerm_app_service_certificate",
		},
		{
			Name: "azurerm_app_service_certificate_binding",
		},
		{
			Name: "azurerm_app_service_certificate_order",
		},
		{
			Name: "azurerm_app_service_connection",
		},
		{
			Name: "azurerm_app_service_custom_hostname_binding",
		},
		{
			Name: "azurerm_app_service_environment",
		},
		{
			Name: "azurerm_app_service_environment_v3",
		},
		{
			Name: "azurerm_app_service_hybrid_connection",
		},
		{
			Name: "azurerm_app_service_managed_certificate",
		},
		{
			Name: "azurerm_app_service_plan",
		},
		{
			Name: "azurerm_app_service_public_certificate",
		},
		{
			Name: "azurerm_app_service_slot",
		},
		{
			Name: "azurerm_app_service_slot_custom_hostname_binding",
		},
		{
			Name: "azurerm_app_service_slot_virtual_network_swift_connection",
		},
		{
			Name: "azurerm_app_service_source_control",
		},
		{
			Name: "azurerm_app_service_source_control_slot",
		},
		{
			Name: "azurerm_app_service_source_control_token",
		},
		{
			Name: "azurerm_app_service_virtual_network_swift_connection",
		},
		{
			Name: "azurerm_application_gateway",
		},
		{
			Name: "azurerm_application_insights",
		},
		{
			Name: "azurerm_application_insights_analytics_item",
		},
		{
			Name: "azurerm_application_insights_api_key",
		},
		{
			Name: "azurerm_application_insights_smart_detection_rule",
		},
		{
			Name: "azurerm_application_insights_standard_web_test",
		},
		{
			Name: "azurerm_application_insights_web_test",
		},
		{
			Name: "azurerm_application_insights_workbook",
		},
		{
			Name: "azurerm_application_insights_workbook_template",
		},
		{
			Name: "azurerm_application_security_group",
		},
		{
			Name: "azurerm_arc_kubernetes_cluster",
		},
		{
			Name: "azurerm_arc_kubernetes_cluster_extension",
		},
		{
			Name: "azurerm_arc_kubernetes_flux_configuration",
		},
		{
			Name: "azurerm_arc_machine_extension",
		},
		{
			Name: "azurerm_arc_private_link_scope",
		},
		{
			Name: "azurerm_attestation_provider",
		},
		{
			Name: "azurerm_automanage_configuration",
		},
		{
			Name: "azurerm_automation_account",
		},
		{
			Name: "azurerm_automation_certificate",
		},
		{
			Name: "azurerm_automation_connection",
		},
		{
			Name: "azurerm_automation_connection_certificate",
		},
		{
			Name: "azurerm_automation_connection_classic_certificate",
		},
		{
			Name: "azurerm_automation_connection_service_principal",
		},
		{
			Name: "azurerm_automation_connection_type",
		},
		{
			Name: "azurerm_automation_credential",
		},
		{
			Name: "azurerm_automation_dsc_configuration",
		},
		{
			Name: "azurerm_automation_dsc_nodeconfiguration",
		},
		{
			Name: "azurerm_automation_hybrid_runbook_worker",
		},
		{
			Name: "azurerm_automation_hybrid_runbook_worker_group",
		},
		{
			Name: "azurerm_automation_job_schedule",
		},
		{
			Name: "azurerm_automation_module",
		},
		{
			Name: "azurerm_automation_python3_package",
		},
		{
			Name: "azurerm_automation_runbook",
		},
		{
			Name: "azurerm_automation_schedule",
		},
		{
			Name: "azurerm_automation_software_update_configuration",
		},
		{
			Name: "azurerm_automation_source_control",
		},
		{
			Name: "azurerm_automation_variable_bool",
		},
		{
			Name: "azurerm_automation_variable_datetime",
		},
		{
			Name: "azurerm_automation_variable_int",
		},
		{
			Name: "azurerm_automation_variable_object",
		},
		{
			Name: "azurerm_automation_variable_string",
		},
		{
			Name: "azurerm_automation_watcher",
		},
		{
			Name: "azurerm_automation_webhook",
		},
		{
			Name: "azurerm_availability_set",
		},
		{
			Name: "azurerm_backup_container_storage_account",
		},
		{
			Name: "azurerm_backup_policy_file_share",
		},
		{
			Name: "azurerm_backup_policy_vm",
		},
		{
			Name: "azurerm_backup_policy_vm_workload",
		},
		{
			Name: "azurerm_backup_protected_file_share",
		},
		{
			Name: "azurerm_backup_protected_vm",
		},
		{
			Name: "azurerm_bastion_host",
		},
		{
			Name: "azurerm_batch_account",
		},
		{
			Name: "azurerm_batch_application",
		},
		{
			Name: "azurerm_batch_certificate",
		},
		{
			Name: "azurerm_batch_job",
		},
		{
			Name: "azurerm_batch_pool",
		},
		{
			Name: "azurerm_billing_account_cost_management_export",
		},
		{
			Name: "azurerm_blueprint_assignment",
		},
		{
			Name: "azurerm_bot_channel_alexa",
		},
		{
			Name: "azurerm_bot_channel_direct_line_speech",
		},
		{
			Name: "azurerm_bot_channel_directline",
		},
		{
			Name: "azurerm_bot_channel_email",
		},
		{
			Name: "azurerm_bot_channel_facebook",
		},
		{
			Name: "azurerm_bot_channel_line",
		},
		{
			Name: "azurerm_bot_channel_ms_teams",
		},
		{
			Name: "azurerm_bot_channel_slack",
		},
		{
			Name: "azurerm_bot_channel_sms",
		},
		{
			Name: "azurerm_bot_channel_web_chat",
		},
		{
			Name: "azurerm_bot_channels_registration",
		},
		{
			Name: "azurerm_bot_connection",
		},
		{
			Name: "azurerm_bot_service_azure_bot",
		},
		{
			Name: "azurerm_bot_web_app",
		},
		{
			Name: "azurerm_capacity_reservation",
		},
		{
			Name: "azurerm_capacity_reservation_group",
		},
		{
			Name: "azurerm_cdn_endpoint",
		},
		{
			Name: "azurerm_cdn_endpoint_custom_domain",
		},
		{
			Name: "azurerm_cdn_frontdoor_custom_domain",
		},
		{
			Name: "azurerm_cdn_frontdoor_custom_domain_association",
		},
		{
			Name: "azurerm_cdn_frontdoor_endpoint",
		},
		{
			Name: "azurerm_cdn_frontdoor_firewall_policy",
		},
		{
			Name: "azurerm_cdn_frontdoor_origin",
		},
		{
			Name: "azurerm_cdn_frontdoor_origin_group",
		},
		{
			Name: "azurerm_cdn_frontdoor_profile",
		},
		{
			Name: "azurerm_cdn_frontdoor_route",
		},
		{
			Name: "azurerm_cdn_frontdoor_route_disable_link_to_default_domain",
		},
		{
			Name: "azurerm_cdn_frontdoor_rule",
		},
		{
			Name: "azurerm_cdn_frontdoor_rule_set",
		},
		{
			Name: "azurerm_cdn_frontdoor_secret",
		},
		{
			Name: "azurerm_cdn_frontdoor_security_policy",
		},
		{
			Name: "azurerm_cdn_profile",
		},
		{
			Name: "azurerm_cognitive_account",
		},
		{
			Name: "azurerm_cognitive_account_customer_managed_key",
		},
		{
			Name: "azurerm_cognitive_deployment",
		},
		{
			Name: "azurerm_communication_service",
		},
		{
			Name: "azurerm_confidential_ledger",
		},
		{
			Name: "azurerm_consumption_budget_management_group",
		},
		{
			Name: "azurerm_consumption_budget_resource_group",
		},
		{
			Name: "azurerm_consumption_budget_subscription",
		},
		{
			Name: "azurerm_container_app",
		},
		{
			Name: "azurerm_container_app_environment",
		},
		{
			Name: "azurerm_container_app_environment_certificate",
		},
		{
			Name: "azurerm_container_app_environment_dapr_component",
		},
		{
			Name: "azurerm_container_app_environment_storage",
		},
		{
			Name: "azurerm_container_connected_registry",
		},
		{
			Name: "azurerm_container_group",
		},
		{
			Name: "azurerm_container_registry",
		},
		{
			Name: "azurerm_container_registry_agent_pool",
		},
		{
			Name: "azurerm_container_registry_scope_map",
		},
		{
			Name: "azurerm_container_registry_task",
		},
		{
			Name: "azurerm_container_registry_task_schedule_run_now",
		},
		{
			Name: "azurerm_container_registry_token",
		},
		{
			Name: "azurerm_container_registry_token_password",
		},
		{
			Name: "azurerm_container_registry_webhook",
		},
		{
			Name: "azurerm_cosmosdb_account",
		},
		{
			Name: "azurerm_cosmosdb_cassandra_cluster",
		},
		{
			Name: "azurerm_cosmosdb_cassandra_datacenter",
		},
		{
			Name: "azurerm_cosmosdb_cassandra_keyspace",
		},
		{
			Name: "azurerm_cosmosdb_cassandra_table",
		},
		{
			Name: "azurerm_cosmosdb_gremlin_database",
		},
		{
			Name: "azurerm_cosmosdb_gremlin_graph",
		},
		{
			Name: "azurerm_cosmosdb_mongo_collection",
		},
		{
			Name: "azurerm_cosmosdb_mongo_database",
		},
		{
			Name: "azurerm_cosmosdb_mongo_role_definition",
		},
		{
			Name: "azurerm_cosmosdb_mongo_user_definition",
		},
		{
			Name: "azurerm_cosmosdb_notebook_workspace",
		},
		{
			Name: "azurerm_cosmosdb_postgresql_cluster",
		},
		{
			Name: "azurerm_cosmosdb_postgresql_coordinator_configuration",
		},
		{
			Name: "azurerm_cosmosdb_postgresql_firewall_rule",
		},
		{
			Name: "azurerm_cosmosdb_postgresql_node_configuration",
		},
		{
			Name: "azurerm_cosmosdb_postgresql_role",
		},
		{
			Name: "azurerm_cosmosdb_sql_container",
		},
		{
			Name: "azurerm_cosmosdb_sql_database",
		},
		{
			Name: "azurerm_cosmosdb_sql_dedicated_gateway",
		},
		{
			Name: "azurerm_cosmosdb_sql_function",
		},
		{
			Name: "azurerm_cosmosdb_sql_role_assignment",
		},
		{
			Name: "azurerm_cosmosdb_sql_role_definition",
		},
		{
			Name: "azurerm_cosmosdb_sql_stored_procedure",
		},
		{
			Name: "azurerm_cosmosdb_sql_trigger",
		},
		{
			Name: "azurerm_cosmosdb_table",
		},
		{
			Name: "azurerm_cost_anomaly_alert",
		},
		{
			Name: "azurerm_cost_management_scheduled_action",
		},
		{
			Name: "azurerm_custom_ip_prefix",
		},
		{
			Name: "azurerm_custom_provider",
		},
		{
			Name: "azurerm_dashboard",
		},
		{
			Name: "azurerm_dashboard_grafana",
		},
		{
			Name: "azurerm_data_factory",
		},
		{
			Name: "azurerm_data_factory_custom_dataset",
		},
		{
			Name: "azurerm_data_factory_data_flow",
		},
		{
			Name: "azurerm_data_factory_dataset_azure_blob",
		},
		{
			Name: "azurerm_data_factory_dataset_binary",
		},
		{
			Name: "azurerm_data_factory_dataset_cosmosdb_sqlapi",
		},
		{
			Name: "azurerm_data_factory_dataset_delimited_text",
		},
		{
			Name: "azurerm_data_factory_dataset_http",
		},
		{
			Name: "azurerm_data_factory_dataset_json",
		},
		{
			Name: "azurerm_data_factory_dataset_mysql",
		},
		{
			Name: "azurerm_data_factory_dataset_parquet",
		},
		{
			Name: "azurerm_data_factory_dataset_postgresql",
		},
		{
			Name: "azurerm_data_factory_dataset_snowflake",
		},
		{
			Name: "azurerm_data_factory_dataset_sql_server_table",
		},
		{
			Name: "azurerm_data_factory_flowlet_data_flow",
		},
		{
			Name: "azurerm_data_factory_integration_runtime_azure",
		},
		{
			Name: "azurerm_data_factory_integration_runtime_azure_ssis",
		},
		{
			Name: "azurerm_data_factory_integration_runtime_managed",
		},
		{
			Name: "azurerm_data_factory_integration_runtime_self_hosted",
		},
		{
			Name: "azurerm_data_factory_linked_custom_service",
		},
		{
			Name: "azurerm_data_factory_linked_service_azure_blob_storage",
		},
		{
			Name: "azurerm_data_factory_linked_service_azure_databricks",
		},
		{
			Name: "azurerm_data_factory_linked_service_azure_file_storage",
		},
		{
			Name: "azurerm_data_factory_linked_service_azure_function",
		},
		{
			Name: "azurerm_data_factory_linked_service_azure_search",
		},
		{
			Name: "azurerm_data_factory_linked_service_azure_sql_database",
		},
		{
			Name: "azurerm_data_factory_linked_service_azure_table_storage",
		},
		{
			Name: "azurerm_data_factory_linked_service_cosmosdb",
		},
		{
			Name: "azurerm_data_factory_linked_service_cosmosdb_mongoapi",
		},
		{
			Name: "azurerm_data_factory_linked_service_data_lake_storage_gen2",
		},
		{
			Name: "azurerm_data_factory_linked_service_key_vault",
		},
		{
			Name: "azurerm_data_factory_linked_service_kusto",
		},
		{
			Name: "azurerm_data_factory_linked_service_mysql",
		},
		{
			Name: "azurerm_data_factory_linked_service_odata",
		},
		{
			Name: "azurerm_data_factory_linked_service_odbc",
		},
		{
			Name: "azurerm_data_factory_linked_service_postgresql",
		},
		{
			Name: "azurerm_data_factory_linked_service_sftp",
		},
		{
			Name: "azurerm_data_factory_linked_service_snowflake",
		},
		{
			Name: "azurerm_data_factory_linked_service_sql_server",
		},
		{
			Name: "azurerm_data_factory_linked_service_synapse",
		},
		{
			Name: "azurerm_data_factory_linked_service_web",
		},
		{
			Name: "azurerm_data_factory_managed_private_endpoint",
		},
		{
			Name: "azurerm_data_factory_pipeline",
		},
		{
			Name: "azurerm_data_factory_trigger_blob_event",
		},
		{
			Name: "azurerm_data_factory_trigger_custom_event",
		},
		{
			Name: "azurerm_data_factory_trigger_schedule",
		},
		{
			Name: "azurerm_data_factory_trigger_tumbling_window",
		},
		{
			Name: "azurerm_data_protection_backup_instance_blob_storage",
		},
		{
			Name: "azurerm_data_protection_backup_instance_disk",
		},
		{
			Name: "azurerm_data_protection_backup_instance_postgresql",
		},
		{
			Name: "azurerm_data_protection_backup_policy_blob_storage",
		},
		{
			Name: "azurerm_data_protection_backup_policy_disk",
		},
		{
			Name: "azurerm_data_protection_backup_policy_postgresql",
		},
		{
			Name: "azurerm_data_protection_backup_vault",
		},
		{
			Name: "azurerm_data_protection_resource_guard",
		},
		{
			Name: "azurerm_data_share",
		},
		{
			Name: "azurerm_data_share_account",
		},
		{
			Name: "azurerm_data_share_dataset_blob_storage",
		},
		{
			Name: "azurerm_data_share_dataset_data_lake_gen2",
		},
		{
			Name: "azurerm_data_share_dataset_kusto_cluster",
		},
		{
			Name: "azurerm_data_share_dataset_kusto_database",
		},
		{
			Name: "azurerm_database_migration_project",
		},
		{
			Name: "azurerm_database_migration_service",
		},
		{
			Name: "azurerm_databox_edge_device",
		},
		{
			Name: "azurerm_databox_edge_order",
		},
		{
			Name: "azurerm_databricks_access_connector",
		},
		{
			Name: "azurerm_databricks_virtual_network_peering",
		},
		{
			Name: "azurerm_databricks_workspace",
		},
		{
			Name: "azurerm_databricks_workspace_customer_managed_key",
		},
		{
			Name: "azurerm_databricks_workspace_root_dbfs_customer_managed_key",
		},
		{
			Name: "azurerm_datadog_monitor",
		},
		{
			Name: "azurerm_datadog_monitor_sso_configuration",
		},
		{
			Name: "azurerm_datadog_monitor_tag_rule",
		},
		{
			Name: "azurerm_dedicated_hardware_security_module",
		},
		{
			Name: "azurerm_dedicated_host",
		},
		{
			Name: "azurerm_dedicated_host_group",
		},
		{
			Name: "azurerm_dev_test_global_vm_shutdown_schedule",
		},
		{
			Name: "azurerm_dev_test_lab",
		},
		{
			Name: "azurerm_dev_test_linux_virtual_machine",
		},
		{
			Name: "azurerm_dev_test_policy",
		},
		{
			Name: "azurerm_dev_test_schedule",
		},
		{
			Name: "azurerm_dev_test_virtual_network",
		},
		{
			Name: "azurerm_dev_test_windows_virtual_machine",
		},
		{
			Name: "azurerm_digital_twins_endpoint_eventgrid",
		},
		{
			Name: "azurerm_digital_twins_endpoint_eventhub",
		},
		{
			Name: "azurerm_digital_twins_endpoint_servicebus",
		},
		{
			Name: "azurerm_digital_twins_instance",
		},
		{
			Name: "azurerm_digital_twins_time_series_database_connection",
		},
		{
			Name: "azurerm_disk_access",
		},
		{
			Name: "azurerm_disk_encryption_set",
		},
		{
			Name: "azurerm_disk_pool",
		},
		{
			Name: "azurerm_disk_pool_iscsi_target",
		},
		{
			Name: "azurerm_disk_pool_iscsi_target_lun",
		},
		{
			Name: "azurerm_disk_pool_managed_disk_attachment",
		},
		{
			Name: "azurerm_dns_a_record",
		},
		{
			Name: "azurerm_dns_aaaa_record",
		},
		{
			Name: "azurerm_dns_caa_record",
		},
		{
			Name: "azurerm_dns_cname_record",
		},
		{
			Name: "azurerm_dns_mx_record",
		},
		{
			Name: "azurerm_dns_ns_record",
		},
		{
			Name: "azurerm_dns_ptr_record",
		},
		{
			Name: "azurerm_dns_srv_record",
		},
		{
			Name: "azurerm_dns_txt_record",
		},
		{
			Name: "azurerm_dns_zone",
		},
		{
			Name: "azurerm_elastic_cloud_elasticsearch",
		},
		{
			Name: "azurerm_email_communication_service",
		},
		{
			Name: "azurerm_eventgrid_domain",
		},
		{
			Name: "azurerm_eventgrid_domain_topic",
		},
		{
			Name: "azurerm_eventgrid_event_subscription",
		},
		{
			Name: "azurerm_eventgrid_system_topic",
		},
		{
			Name: "azurerm_eventgrid_system_topic_event_subscription",
		},
		{
			Name: "azurerm_eventgrid_topic",
		},
		{
			Name: "azurerm_eventhub",
		},
		{
			Name: "azurerm_eventhub_authorization_rule",
		},
		{
			Name: "azurerm_eventhub_cluster",
		},
		{
			Name: "azurerm_eventhub_consumer_group",
		},
		{
			Name: "azurerm_eventhub_namespace",
		},
		{
			Name: "azurerm_eventhub_namespace_authorization_rule",
		},
		{
			Name: "azurerm_eventhub_namespace_customer_managed_key",
		},
		{
			Name: "azurerm_eventhub_namespace_disaster_recovery_config",
		},
		{
			Name: "azurerm_eventhub_namespace_schema_group",
		},
		{
			Name: "azurerm_express_route_circuit",
		},
		{
			Name: "azurerm_express_route_circuit_authorization",
		},
		{
			Name: "azurerm_express_route_circuit_connection",
		},
		{
			Name: "azurerm_express_route_circuit_peering",
		},
		{
			Name: "azurerm_express_route_connection",
		},
		{
			Name: "azurerm_express_route_gateway",
		},
		{
			Name: "azurerm_express_route_port",
		},
		{
			Name: "azurerm_express_route_port_authorization",
		},
		{
			Name: "azurerm_federated_identity_credential",
		},
		{
			Name: "azurerm_firewall",
		},
		{
			Name: "azurerm_firewall_application_rule_collection",
		},
		{
			Name: "azurerm_firewall_nat_rule_collection",
		},
		{
			Name: "azurerm_firewall_network_rule_collection",
		},
		{
			Name: "azurerm_firewall_policy",
		},
		{
			Name: "azurerm_firewall_policy_rule_collection_group",
		},
		{
			Name: "azurerm_fluid_relay_server",
		},
		{
			Name: "azurerm_frontdoor",
		},
		{
			Name: "azurerm_frontdoor_custom_https_configuration",
		},
		{
			Name: "azurerm_frontdoor_firewall_policy",
		},
		{
			Name: "azurerm_frontdoor_rules_engine",
		},
		{
			Name: "azurerm_function_app",
		},
		{
			Name: "azurerm_function_app_active_slot",
		},
		{
			Name: "azurerm_function_app_function",
		},
		{
			Name: "azurerm_function_app_hybrid_connection",
		},
		{
			Name: "azurerm_function_app_slot",
		},
		{
			Name: "azurerm_gallery_application",
		},
		{
			Name: "azurerm_gallery_application_version",
		},
		{
			Name: "azurerm_graph_account",
		},
		{
			Name: "azurerm_graph_services_account",
		},
		{
			Name: "azurerm_hdinsight_hadoop_cluster",
		},
		{
			Name: "azurerm_hdinsight_hbase_cluster",
		},
		{
			Name: "azurerm_hdinsight_interactive_query_cluster",
		},
		{
			Name: "azurerm_hdinsight_kafka_cluster",
		},
		{
			Name: "azurerm_hdinsight_spark_cluster",
		},
		{
			Name: "azurerm_healthbot",
		},
		{
			Name: "azurerm_healthcare_dicom_service",
		},
		{
			Name: "azurerm_healthcare_fhir_service",
		},
		{
			Name: "azurerm_healthcare_medtech_service",
		},
		{
			Name: "azurerm_healthcare_medtech_service_fhir_destination",
		},
		{
			Name: "azurerm_healthcare_service",
		},
		{
			Name: "azurerm_healthcare_workspace",
		},
		{
			Name: "azurerm_hpc_cache",
		},
		{
			Name: "azurerm_hpc_cache_access_policy",
		},
		{
			Name: "azurerm_hpc_cache_blob_nfs_target",
		},
		{
			Name: "azurerm_hpc_cache_blob_target",
		},
		{
			Name: "azurerm_hpc_cache_nfs_target",
		},
		{
			Name: "azurerm_image",
		},
		{
			Name: "azurerm_integration_service_environment",
		},
		{
			Name: "azurerm_iot_security_device_group",
		},
		{
			Name: "azurerm_iot_security_solution",
		},
		{
			Name: "azurerm_iot_time_series_insights_access_policy",
		},
		{
			Name: "azurerm_iot_time_series_insights_event_source_eventhub",
		},
		{
			Name: "azurerm_iot_time_series_insights_event_source_iothub",
		},
		{
			Name: "azurerm_iot_time_series_insights_gen2_environment",
		},
		{
			Name: "azurerm_iot_time_series_insights_reference_data_set",
		},
		{
			Name: "azurerm_iot_time_series_insights_standard_environment",
		},
		{
			Name: "azurerm_iotcentral_application",
		},
		{
			Name: "azurerm_iotcentral_application_network_rule_set",
		},
		{
			Name: "azurerm_iothub",
		},
		{
			Name: "azurerm_iothub_certificate",
		},
		{
			Name: "azurerm_iothub_consumer_group",
		},
		{
			Name: "azurerm_iothub_device_update_account",
		},
		{
			Name: "azurerm_iothub_device_update_instance",
		},
		{
			Name: "azurerm_iothub_dps",
		},
		{
			Name: "azurerm_iothub_dps_certificate",
		},
		{
			Name: "azurerm_iothub_dps_shared_access_policy",
		},
		{
			Name: "azurerm_iothub_endpoint_cosmosdb_account",
		},
		{
			Name: "azurerm_iothub_endpoint_eventhub",
		},
		{
			Name: "azurerm_iothub_endpoint_servicebus_queue",
		},
		{
			Name: "azurerm_iothub_endpoint_servicebus_topic",
		},
		{
			Name: "azurerm_iothub_endpoint_storage_container",
		},
		{
			Name: "azurerm_iothub_enrichment",
		},
		{
			Name: "azurerm_iothub_fallback_route",
		},
		{
			Name: "azurerm_iothub_file_upload",
		},
		{
			Name: "azurerm_iothub_route",
		},
		{
			Name: "azurerm_iothub_shared_access_policy",
		},
		{
			Name: "azurerm_ip_group",
		},
		{
			Name: "azurerm_ip_group_cidr",
		},
		{
			Name: "azurerm_key_vault",
		},
		{
			Name: "azurerm_key_vault_access_policy",
		},
		{
			Name: "azurerm_key_vault_certificate",
		},
		{
			Name: "azurerm_key_vault_certificate_contacts",
		},
		{
			Name: "azurerm_key_vault_certificate_issuer",
		},
		{
			Name: "azurerm_key_vault_key",
		},
		{
			Name: "azurerm_key_vault_managed_hardware_security_module",
		},
		{
			Name: "azurerm_key_vault_managed_storage_account",
		},
		{
			Name: "azurerm_key_vault_managed_storage_account_sas_token_definition",
		},
		{
			Name: "azurerm_key_vault_secret",
		},
		{
			Name: "azurerm_kubernetes_cluster",
		},
		{
			Name: "azurerm_kubernetes_cluster_extension",
		},
		{
			Name: "azurerm_kubernetes_cluster_node_pool",
		},
		{
			Name: "azurerm_kubernetes_cluster_trusted_access_role_binding",
		},
		{
			Name: "azurerm_kubernetes_fleet_manager",
		},
		{
			Name: "azurerm_kubernetes_flux_configuration",
		},
		{
			Name: "azurerm_kusto_attached_database_configuration",
		},
		{
			Name: "azurerm_kusto_cluster",
		},
		{
			Name: "azurerm_kusto_cluster_customer_managed_key",
		},
		{
			Name: "azurerm_kusto_cluster_managed_private_endpoint",
		},
		{
			Name: "azurerm_kusto_cluster_principal_assignment",
		},
		{
			Name: "azurerm_kusto_cosmosdb_data_connection",
		},
		{
			Name: "azurerm_kusto_database",
		},
		{
			Name: "azurerm_kusto_database_principal_assignment",
		},
		{
			Name: "azurerm_kusto_eventgrid_data_connection",
		},
		{
			Name: "azurerm_kusto_eventhub_data_connection",
		},
		{
			Name: "azurerm_kusto_iothub_data_connection",
		},
		{
			Name: "azurerm_kusto_script",
		},
		{
			Name: "azurerm_lab_service_lab",
		},
		{
			Name: "azurerm_lab_service_plan",
		},
		{
			Name: "azurerm_lab_service_schedule",
		},
		{
			Name: "azurerm_lab_service_user",
		},
		{
			Name: "azurerm_lb",
		},
		{
			Name: "azurerm_lb_backend_address_pool",
		},
		{
			Name: "azurerm_lb_backend_address_pool_address",
		},
		{
			Name: "azurerm_lb_nat_pool",
		},
		{
			Name: "azurerm_lb_nat_rule",
		},
		{
			Name: "azurerm_lb_outbound_rule",
		},
		{
			Name: "azurerm_lb_probe",
		},
		{
			Name: "azurerm_lb_rule",
		},
		{
			Name: "azurerm_lighthouse_assignment",
		},
		{
			Name: "azurerm_lighthouse_definition",
		},
		{
			Name: "azurerm_linux_function_app",
		},
		{
			Name: "azurerm_linux_function_app_slot",
		},
		{
			Name: "azurerm_linux_virtual_machine",
		},
		{
			Name: "azurerm_linux_virtual_machine_scale_set",
		},
		{
			Name: "azurerm_linux_web_app",
		},
		{
			Name: "azurerm_linux_web_app_slot",
		},
		{
			Name: "azurerm_load_test",
		},
		{
			Name: "azurerm_local_network_gateway",
		},
		{
			Name: "azurerm_log_analytics_cluster",
		},
		{
			Name: "azurerm_log_analytics_cluster_customer_managed_key",
		},
		{
			Name: "azurerm_log_analytics_data_export_rule",
		},
		{
			Name: "azurerm_log_analytics_datasource_windows_event",
		},
		{
			Name: "azurerm_log_analytics_datasource_windows_performance_counter",
		},
		{
			Name: "azurerm_log_analytics_linked_service",
		},
		{
			Name: "azurerm_log_analytics_linked_storage_account",
		},
		{
			Name: "azurerm_log_analytics_query_pack",
		},
		{
			Name: "azurerm_log_analytics_query_pack_query",
		},
		{
			Name: "azurerm_log_analytics_saved_search",
		},
		{
			Name: "azurerm_log_analytics_solution",
		},
		{
			Name: "azurerm_log_analytics_storage_insights",
		},
		{
			Name: "azurerm_log_analytics_workspace",
		},
		{
			Name: "azurerm_logic_app_action_custom",
		},
		{
			Name: "azurerm_logic_app_action_http",
		},
		{
			Name: "azurerm_logic_app_integration_account",
		},
		{
			Name: "azurerm_logic_app_integration_account_agreement",
		},
		{
			Name: "azurerm_logic_app_integration_account_assembly",
		},
		{
			Name: "azurerm_logic_app_integration_account_batch_configuration",
		},
		{
			Name: "azurerm_logic_app_integration_account_certificate",
		},
		{
			Name: "azurerm_logic_app_integration_account_map",
		},
		{
			Name: "azurerm_logic_app_integration_account_partner",
		},
		{
			Name: "azurerm_logic_app_integration_account_schema",
		},
		{
			Name: "azurerm_logic_app_integration_account_session",
		},
		{
			Name: "azurerm_logic_app_standard",
		},
		{
			Name: "azurerm_logic_app_trigger_custom",
		},
		{
			Name: "azurerm_logic_app_trigger_http_request",
		},
		{
			Name: "azurerm_logic_app_trigger_recurrence",
		},
		{
			Name: "azurerm_logic_app_workflow",
		},
		{
			Name: "azurerm_logz_monitor",
		},
		{
			Name: "azurerm_logz_sub_account",
		},
		{
			Name: "azurerm_logz_sub_account_tag_rule",
		},
		{
			Name: "azurerm_logz_tag_rule",
		},
		{
			Name: "azurerm_machine_learning_compute_cluster",
		},
		{
			Name: "azurerm_machine_learning_compute_instance",
		},
		{
			Name: "azurerm_machine_learning_datastore_blobstorage",
		},
		{
			Name: "azurerm_machine_learning_datastore_datalake_gen2",
		},
		{
			Name: "azurerm_machine_learning_datastore_fileshare",
		},
		{
			Name: "azurerm_machine_learning_inference_cluster",
		},
		{
			Name: "azurerm_machine_learning_synapse_spark",
		},
		{
			Name: "azurerm_machine_learning_workspace",
		},
		{
			Name: "azurerm_maintenance_assignment_dedicated_host",
		},
		{
			Name: "azurerm_maintenance_assignment_virtual_machine",
		},
		{
			Name: "azurerm_maintenance_assignment_virtual_machine_scale_set",
		},
		{
			Name: "azurerm_maintenance_configuration",
		},
		{
			Name: "azurerm_managed_application",
		},
		{
			Name: "azurerm_managed_application_definition",
		},
		{
			Name: "azurerm_managed_disk",
		},
		{
			Name: "azurerm_managed_disk_sas_token",
		},
		{
			Name: "azurerm_managed_lustre_file_system",
		},
		{
			Name: "azurerm_management_group",
		},
		{
			Name: "azurerm_management_group_policy_assignment",
		},
		{
			Name: "azurerm_management_group_policy_exemption",
		},
		{
			Name: "azurerm_management_group_policy_remediation",
		},
		{
			Name: "azurerm_management_group_subscription_association",
		},
		{
			Name: "azurerm_management_group_template_deployment",
		},
		{
			Name: "azurerm_management_lock",
		},
		{
			Name: "azurerm_maps_account",
		},
		{
			Name: "azurerm_maps_creator",
		},
		{
			Name: "azurerm_mariadb_configuration",
		},
		{
			Name: "azurerm_mariadb_database",
		},
		{
			Name: "azurerm_mariadb_firewall_rule",
		},
		{
			Name: "azurerm_mariadb_server",
		},
		{
			Name: "azurerm_mariadb_virtual_network_rule",
		},
		{
			Name: "azurerm_marketplace_agreement",
		},
		{
			Name: "azurerm_marketplace_role_assignment",
		},
		{
			Name: "azurerm_media_asset",
		},
		{
			Name: "azurerm_media_asset_filter",
		},
		{
			Name: "azurerm_media_content_key_policy",
		},
		{
			Name: "azurerm_media_job",
		},
		{
			Name: "azurerm_media_live_event",
		},
		{
			Name: "azurerm_media_live_event_output",
		},
		{
			Name: "azurerm_media_services_account",
		},
		{
			Name: "azurerm_media_services_account_filter",
		},
		{
			Name: "azurerm_media_streaming_endpoint",
		},
		{
			Name: "azurerm_media_streaming_locator",
		},
		{
			Name: "azurerm_media_streaming_policy",
		},
		{
			Name: "azurerm_media_transform",
		},
		{
			Name: "azurerm_mobile_network",
		},
		{
			Name: "azurerm_mobile_network_attached_data_network",
		},
		{
			Name: "azurerm_mobile_network_data_network",
		},
		{
			Name: "azurerm_mobile_network_packet_core_control_plane",
		},
		{
			Name: "azurerm_mobile_network_packet_core_data_plane",
		},
		{
			Name: "azurerm_mobile_network_service",
		},
		{
			Name: "azurerm_mobile_network_sim",
		},
		{
			Name: "azurerm_mobile_network_sim_group",
		},
		{
			Name: "azurerm_mobile_network_sim_policy",
		},
		{
			Name: "azurerm_mobile_network_site",
		},
		{
			Name: "azurerm_mobile_network_slice",
		},
		{
			Name: "azurerm_monitor_aad_diagnostic_setting",
		},
		{
			Name: "azurerm_monitor_action_group",
		},
		{
			Name: "azurerm_monitor_action_rule_action_group",
		},
		{
			Name: "azurerm_monitor_action_rule_suppression",
		},
		{
			Name: "azurerm_monitor_activity_log_alert",
		},
		{
			Name: "azurerm_monitor_alert_processing_rule_action_group",
		},
		{
			Name: "azurerm_monitor_alert_processing_rule_suppression",
		},
		{
			Name: "azurerm_monitor_alert_prometheus_rule_group",
		},
		{
			Name: "azurerm_monitor_autoscale_setting",
		},
		{
			Name: "azurerm_monitor_data_collection_endpoint",
		},
		{
			Name: "azurerm_monitor_data_collection_rule",
		},
		{
			Name: "azurerm_monitor_data_collection_rule_association",
		},
		{
			Name: "azurerm_monitor_diagnostic_setting",
		},
		{
			Name: "azurerm_monitor_log_profile",
		},
		{
			Name: "azurerm_monitor_metric_alert",
		},
		{
			Name: "azurerm_monitor_private_link_scope",
		},
		{
			Name: "azurerm_monitor_private_link_scoped_service",
		},
		{
			Name: "azurerm_monitor_scheduled_query_rules_alert",
		},
		{
			Name: "azurerm_monitor_scheduled_query_rules_alert_v2",
		},
		{
			Name: "azurerm_monitor_scheduled_query_rules_log",
		},
		{
			Name: "azurerm_monitor_smart_detector_alert_rule",
		},
		{
			Name: "azurerm_monitor_workspace",
		},
		{
			Name: "azurerm_mssql_database",
		},
		{
			Name: "azurerm_mssql_database_extended_auditing_policy",
		},
		{
			Name: "azurerm_mssql_database_vulnerability_assessment_rule_baseline",
		},
		{
			Name: "azurerm_mssql_elasticpool",
		},
		{
			Name: "azurerm_mssql_failover_group",
		},
		{
			Name: "azurerm_mssql_firewall_rule",
		},
		{
			Name: "azurerm_mssql_job_agent",
		},
		{
			Name: "azurerm_mssql_job_credential",
		},
		{
			Name: "azurerm_mssql_managed_database",
		},
		{
			Name: "azurerm_mssql_managed_instance",
		},
		{
			Name: "azurerm_mssql_managed_instance_active_directory_administrator",
		},
		{
			Name: "azurerm_mssql_managed_instance_failover_group",
		},
		{
			Name: "azurerm_mssql_managed_instance_security_alert_policy",
		},
		{
			Name: "azurerm_mssql_managed_instance_transparent_data_encryption",
		},
		{
			Name: "azurerm_mssql_managed_instance_vulnerability_assessment",
		},
		{
			Name: "azurerm_mssql_outbound_firewall_rule",
		},
		{
			Name: "azurerm_mssql_server",
		},
		{
			Name: "azurerm_mssql_server_dns_alias",
		},
		{
			Name: "azurerm_mssql_server_extended_auditing_policy",
		},
		{
			Name: "azurerm_mssql_server_microsoft_support_auditing_policy",
		},
		{
			Name: "azurerm_mssql_server_security_alert_policy",
		},
		{
			Name: "azurerm_mssql_server_transparent_data_encryption",
		},
		{
			Name: "azurerm_mssql_server_vulnerability_assessment",
		},
		{
			Name: "azurerm_mssql_virtual_machine",
		},
		{
			Name: "azurerm_mssql_virtual_machine_availability_group_listener",
		},
		{
			Name: "azurerm_mssql_virtual_machine_group",
		},
		{
			Name: "azurerm_mssql_virtual_network_rule",
		},
		{
			Name: "azurerm_mysql_active_directory_administrator",
		},
		{
			Name: "azurerm_mysql_configuration",
		},
		{
			Name: "azurerm_mysql_database",
		},
		{
			Name: "azurerm_mysql_firewall_rule",
		},
		{
			Name: "azurerm_mysql_flexible_database",
		},
		{
			Name: "azurerm_mysql_flexible_server",
		},
		{
			Name: "azurerm_mysql_flexible_server_active_directory_administrator",
		},
		{
			Name: "azurerm_mysql_flexible_server_configuration",
		},
		{
			Name: "azurerm_mysql_flexible_server_firewall_rule",
		},
		{
			Name: "azurerm_mysql_server",
		},
		{
			Name: "azurerm_mysql_server_key",
		},
		{
			Name: "azurerm_mysql_virtual_network_rule",
		},
		{
			Name: "azurerm_nat_gateway",
		},
		{
			Name: "azurerm_nat_gateway_public_ip_association",
		},
		{
			Name: "azurerm_nat_gateway_public_ip_prefix_association",
		},
		{
			Name: "azurerm_netapp_account",
		},
		{
			Name: "azurerm_netapp_pool",
		},
		{
			Name: "azurerm_netapp_snapshot",
		},
		{
			Name: "azurerm_netapp_snapshot_policy",
		},
		{
			Name: "azurerm_netapp_volume",
		},
		{
			Name: "azurerm_netapp_volume_group_sap_hana",
		},
		{
			Name: "azurerm_netapp_volume_quota_rule",
		},
		{
			Name: "azurerm_network_connection_monitor",
		},
		{
			Name: "azurerm_network_ddos_protection_plan",
		},
		{
			Name: "azurerm_network_function_azure_traffic_collector",
		},
		{
			Name: "azurerm_network_interface",
		},
		{
			Name: "azurerm_network_interface_application_gateway_backend_address_pool_association",
		},
		{
			Name: "azurerm_network_interface_application_security_group_association",
		},
		{
			Name: "azurerm_network_interface_backend_address_pool_association",
		},
		{
			Name: "azurerm_network_interface_nat_rule_association",
		},
		{
			Name: "azurerm_network_interface_security_group_association",
		},
		{
			Name: "azurerm_network_manager",
		},
		{
			Name: "azurerm_network_manager_admin_rule",
		},
		{
			Name: "azurerm_network_manager_admin_rule_collection",
		},
		{
			Name: "azurerm_network_manager_connectivity_configuration",
		},
		{
			Name: "azurerm_network_manager_deployment",
		},
		{
			Name: "azurerm_network_manager_management_group_connection",
		},
		{
			Name: "azurerm_network_manager_network_group",
		},
		{
			Name: "azurerm_network_manager_scope_connection",
		},
		{
			Name: "azurerm_network_manager_security_admin_configuration",
		},
		{
			Name: "azurerm_network_manager_static_member",
		},
		{
			Name: "azurerm_network_manager_subscription_connection",
		},
		{
			Name: "azurerm_network_packet_capture",
		},
		{
			Name: "azurerm_network_profile",
		},
		{
			Name: "azurerm_network_security_group",
		},
		{
			Name: "azurerm_network_security_rule",
		},
		{
			Name: "azurerm_network_watcher",
		},
		{
			Name: "azurerm_network_watcher_flow_log",
		},
		{
			Name: "azurerm_new_relic_monitor",
		},
		{
			Name: "azurerm_nginx_certificate",
		},
		{
			Name: "azurerm_nginx_configuration",
		},
		{
			Name: "azurerm_nginx_deployment",
		},
		{
			Name: "azurerm_notification_hub",
		},
		{
			Name: "azurerm_notification_hub_authorization_rule",
		},
		{
			Name: "azurerm_notification_hub_namespace",
		},
		{
			Name: "azurerm_orbital_contact",
		},
		{
			Name: "azurerm_orbital_contact_profile",
		},
		{
			Name: "azurerm_orbital_spacecraft",
		},
		{
			Name: "azurerm_orchestrated_virtual_machine_scale_set",
		},
		{
			Name: "azurerm_palo_alto_local_rulestack",
		},
		{
			Name: "azurerm_palo_alto_local_rulestack_certificate",
		},
		{
			Name: "azurerm_palo_alto_local_rulestack_fqdn_list",
		},
		{
			Name: "azurerm_palo_alto_local_rulestack_outbound_trust_certificate_association",
		},
		{
			Name: "azurerm_palo_alto_local_rulestack_outbound_untrust_certificate_association",
		},
		{
			Name: "azurerm_palo_alto_local_rulestack_prefix_list",
		},
		{
			Name: "azurerm_palo_alto_local_rulestack_rule",
		},
		{
			Name: "azurerm_palo_alto_next_generation_firewall_virtual_hub_local_rulestack",
		},
		{
			Name: "azurerm_palo_alto_next_generation_firewall_virtual_hub_panorama",
		},
		{
			Name: "azurerm_palo_alto_next_generation_firewall_virtual_network_local_rulestack",
		},
		{
			Name: "azurerm_palo_alto_next_generation_firewall_virtual_network_panorama",
		},
		{
			Name: "azurerm_palo_alto_virtual_network_appliance",
		},
		{
			Name: "azurerm_pim_active_role_assignment",
		},
		{
			Name: "azurerm_pim_eligible_role_assignment",
		},
		{
			Name: "azurerm_point_to_site_vpn_gateway",
		},
		{
			Name: "azurerm_policy_definition",
		},
		{
			Name: "azurerm_policy_set_definition",
		},
		{
			Name: "azurerm_policy_virtual_machine_configuration_assignment",
		},
		{
			Name: "azurerm_portal_dashboard",
		},
		{
			Name: "azurerm_portal_tenant_configuration",
		},
		{
			Name: "azurerm_postgresql_active_directory_administrator",
		},
		{
			Name: "azurerm_postgresql_configuration",
		},
		{
			Name: "azurerm_postgresql_database",
		},
		{
			Name: "azurerm_postgresql_firewall_rule",
		},
		{
			Name: "azurerm_postgresql_flexible_server",
		},
		{
			Name: "azurerm_postgresql_flexible_server_active_directory_administrator",
		},
		{
			Name: "azurerm_postgresql_flexible_server_configuration",
		},
		{
			Name: "azurerm_postgresql_flexible_server_database",
		},
		{
			Name: "azurerm_postgresql_flexible_server_firewall_rule",
		},
		{
			Name: "azurerm_postgresql_server",
		},
		{
			Name: "azurerm_postgresql_server_key",
		},
		{
			Name: "azurerm_postgresql_virtual_network_rule",
		},
		{
			Name: "azurerm_powerbi_embedded",
		},
		{
			Name: "azurerm_private_dns_a_record",
		},
		{
			Name: "azurerm_private_dns_aaaa_record",
		},
		{
			Name: "azurerm_private_dns_cname_record",
		},
		{
			Name: "azurerm_private_dns_mx_record",
		},
		{
			Name: "azurerm_private_dns_ptr_record",
		},
		{
			Name: "azurerm_private_dns_resolver",
		},
		{
			Name: "azurerm_private_dns_resolver_dns_forwarding_ruleset",
		},
		{
			Name: "azurerm_private_dns_resolver_forwarding_rule",
		},
		{
			Name: "azurerm_private_dns_resolver_inbound_endpoint",
		},
		{
			Name: "azurerm_private_dns_resolver_outbound_endpoint",
		},
		{
			Name: "azurerm_private_dns_resolver_virtual_network_link",
		},
		{
			Name: "azurerm_private_dns_srv_record",
		},
		{
			Name: "azurerm_private_dns_txt_record",
		},
		{
			Name: "azurerm_private_dns_zone",
		},
		{
			Name: "azurerm_private_dns_zone_virtual_network_link",
		},
		{
			Name: "azurerm_private_endpoint",
		},
		{
			Name: "azurerm_private_endpoint_application_security_group_association",
		},
		{
			Name: "azurerm_private_link_service",
		},
		{
			Name: "azurerm_proximity_placement_group",
		},
		{
			Name: "azurerm_public_ip",
		},
		{
			Name: "azurerm_public_ip_prefix",
		},
		{
			Name: "azurerm_purview_account",
		},
		{
			Name: "azurerm_recovery_services_vault",
		},
		{
			Name: "azurerm_recovery_services_vault_resource_guard_association",
		},
		{
			Name: "azurerm_redis_cache",
		},
		{
			Name: "azurerm_redis_enterprise_cluster",
		},
		{
			Name: "azurerm_redis_enterprise_database",
		},
		{
			Name: "azurerm_redis_firewall_rule",
		},
		{
			Name: "azurerm_redis_linked_server",
		},
		{
			Name: "azurerm_relay_hybrid_connection",
		},
		{
			Name: "azurerm_relay_hybrid_connection_authorization_rule",
		},
		{
			Name: "azurerm_relay_namespace",
		},
		{
			Name: "azurerm_relay_namespace_authorization_rule",
		},
		{
			Name: "azurerm_resource_deployment_script_azure_cli",
		},
		{
			Name: "azurerm_resource_deployment_script_azure_power_shell",
		},
		{
			Name: "azurerm_resource_group",
		},
		{
			Name: "azurerm_resource_group_cost_management_export",
		},
		{
			Name: "azurerm_resource_group_cost_management_view",
		},
		{
			Name: "azurerm_resource_group_policy_assignment",
		},
		{
			Name: "azurerm_resource_group_policy_exemption",
		},
		{
			Name: "azurerm_resource_group_policy_remediation",
		},
		{
			Name: "azurerm_resource_group_template_deployment",
		},
		{
			Name: "azurerm_resource_policy_assignment",
		},
		{
			Name: "azurerm_resource_policy_exemption",
		},
		{
			Name: "azurerm_resource_policy_remediation",
		},
		{
			Name: "azurerm_resource_provider_registration",
		},
		{
			Name: "azurerm_role_assignment",
		},
		{
			Name: "azurerm_role_definition",
		},
		{
			Name: "azurerm_route",
		},
		{
			Name: "azurerm_route_filter",
		},
		{
			Name: "azurerm_route_map",
		},
		{
			Name: "azurerm_route_server",
		},
		{
			Name: "azurerm_route_server_bgp_connection",
		},
		{
			Name: "azurerm_route_table",
		},
		{
			Name: "azurerm_search_service",
		},
		{
			Name: "azurerm_search_shared_private_link_service",
		},
		{
			Name: "azurerm_security_center_assessment",
		},
		{
			Name: "azurerm_security_center_assessment_policy",
		},
		{
			Name: "azurerm_security_center_auto_provisioning",
		},
		{
			Name: "azurerm_security_center_automation",
		},
		{
			Name: "azurerm_security_center_contact",
		},
		{
			Name: "azurerm_security_center_server_vulnerability_assessment",
		},
		{
			Name: "azurerm_security_center_server_vulnerability_assessment_virtual_machine",
		},
		{
			Name: "azurerm_security_center_setting",
		},
		{
			Name: "azurerm_security_center_subscription_pricing",
		},
		{
			Name: "azurerm_security_center_workspace",
		},
		{
			Name: "azurerm_sentinel_alert_rule_anomaly_built_in",
		},
		{
			Name: "azurerm_sentinel_alert_rule_anomaly_duplicate",
		},
		{
			Name: "azurerm_sentinel_alert_rule_fusion",
		},
		{
			Name: "azurerm_sentinel_alert_rule_machine_learning_behavior_analytics",
		},
		{
			Name: "azurerm_sentinel_alert_rule_ms_security_incident",
		},
		{
			Name: "azurerm_sentinel_alert_rule_nrt",
		},
		{
			Name: "azurerm_sentinel_alert_rule_scheduled",
		},
		{
			Name: "azurerm_sentinel_alert_rule_threat_intelligence",
		},
		{
			Name: "azurerm_sentinel_automation_rule",
		},
		{
			Name: "azurerm_sentinel_data_connector_aws_cloud_trail",
		},
		{
			Name: "azurerm_sentinel_data_connector_aws_s3",
		},
		{
			Name: "azurerm_sentinel_data_connector_azure_active_directory",
		},
		{
			Name: "azurerm_sentinel_data_connector_azure_advanced_threat_protection",
		},
		{
			Name: "azurerm_sentinel_data_connector_azure_security_center",
		},
		{
			Name: "azurerm_sentinel_data_connector_dynamics_365",
		},
		{
			Name: "azurerm_sentinel_data_connector_iot",
		},
		{
			Name: "azurerm_sentinel_data_connector_microsoft_cloud_app_security",
		},
		{
			Name: "azurerm_sentinel_data_connector_microsoft_defender_advanced_threat_protection",
		},
		{
			Name: "azurerm_sentinel_data_connector_microsoft_threat_intelligence",
		},
		{
			Name: "azurerm_sentinel_data_connector_microsoft_threat_protection",
		},
		{
			Name: "azurerm_sentinel_data_connector_office_365",
		},
		{
			Name: "azurerm_sentinel_data_connector_office_365_project",
		},
		{
			Name: "azurerm_sentinel_data_connector_office_atp",
		},
		{
			Name: "azurerm_sentinel_data_connector_office_irm",
		},
		{
			Name: "azurerm_sentinel_data_connector_office_power_bi",
		},
		{
			Name: "azurerm_sentinel_data_connector_threat_intelligence",
		},
		{
			Name: "azurerm_sentinel_data_connector_threat_intelligence_taxii",
		},
		{
			Name: "azurerm_sentinel_log_analytics_workspace_onboarding",
		},
		{
			Name: "azurerm_sentinel_metadata",
		},
		{
			Name: "azurerm_sentinel_threat_intelligence_indicator",
		},
		{
			Name: "azurerm_sentinel_watchlist",
		},
		{
			Name: "azurerm_sentinel_watchlist_item",
		},
		{
			Name: "azurerm_service_fabric_cluster",
		},
		{
			Name: "azurerm_service_fabric_managed_cluster",
		},
		{
			Name: "azurerm_service_plan",
		},
		{
			Name: "azurerm_servicebus_namespace",
		},
		{
			Name: "azurerm_servicebus_namespace_authorization_rule",
		},
		{
			Name: "azurerm_servicebus_namespace_disaster_recovery_config",
		},
		{
			Name: "azurerm_servicebus_namespace_network_rule_set",
		},
		{
			Name: "azurerm_servicebus_queue",
		},
		{
			Name: "azurerm_servicebus_queue_authorization_rule",
		},
		{
			Name: "azurerm_servicebus_subscription",
		},
		{
			Name: "azurerm_servicebus_subscription_rule",
		},
		{
			Name: "azurerm_servicebus_topic",
		},
		{
			Name: "azurerm_servicebus_topic_authorization_rule",
		},
		{
			Name: "azurerm_shared_image",
		},
		{
			Name: "azurerm_shared_image_gallery",
		},
		{
			Name: "azurerm_shared_image_version",
		},
		{
			Name: "azurerm_signalr_service",
		},
		{
			Name: "azurerm_signalr_service_custom_certificate",
		},
		{
			Name: "azurerm_signalr_service_custom_domain",
		},
		{
			Name: "azurerm_signalr_service_network_acl",
		},
		{
			Name: "azurerm_signalr_shared_private_link_resource",
		},
		{
			Name: "azurerm_site_recovery_fabric",
		},
		{
			Name: "azurerm_site_recovery_hyperv_network_mapping",
		},
		{
			Name: "azurerm_site_recovery_hyperv_replication_policy",
		},
		{
			Name: "azurerm_site_recovery_hyperv_replication_policy_association",
		},
		{
			Name: "azurerm_site_recovery_network_mapping",
		},
		{
			Name: "azurerm_site_recovery_protection_container",
		},
		{
			Name: "azurerm_site_recovery_protection_container_mapping",
		},
		{
			Name: "azurerm_site_recovery_replicated_vm",
		},
		{
			Name: "azurerm_site_recovery_replication_policy",
		},
		{
			Name: "azurerm_site_recovery_replication_recovery_plan",
		},
		{
			Name: "azurerm_site_recovery_services_vault_hyperv_site",
		},
		{
			Name: "azurerm_site_recovery_vmware_replication_policy",
		},
		{
			Name: "azurerm_site_recovery_vmware_replication_policy_association",
		},
		{
			Name: "azurerm_snapshot",
		},
		{
			Name: "azurerm_source_control_token",
		},
		{
			Name: "azurerm_spatial_anchors_account",
		},
		{
			Name: "azurerm_spring_cloud_accelerator",
		},
		{
			Name: "azurerm_spring_cloud_active_deployment",
		},
		{
			Name: "azurerm_spring_cloud_api_portal",
		},
		{
			Name: "azurerm_spring_cloud_api_portal_custom_domain",
		},
		{
			Name: "azurerm_spring_cloud_app",
		},
		{
			Name: "azurerm_spring_cloud_app_cosmosdb_association",
		},
		{
			Name: "azurerm_spring_cloud_app_mysql_association",
		},
		{
			Name: "azurerm_spring_cloud_app_redis_association",
		},
		{
			Name: "azurerm_spring_cloud_application_live_view",
		},
		{
			Name: "azurerm_spring_cloud_build_deployment",
		},
		{
			Name: "azurerm_spring_cloud_build_pack_binding",
		},
		{
			Name: "azurerm_spring_cloud_builder",
		},
		{
			Name: "azurerm_spring_cloud_certificate",
		},
		{
			Name: "azurerm_spring_cloud_configuration_service",
		},
		{
			Name: "azurerm_spring_cloud_connection",
		},
		{
			Name: "azurerm_spring_cloud_container_deployment",
		},
		{
			Name: "azurerm_spring_cloud_custom_domain",
		},
		{
			Name: "azurerm_spring_cloud_customized_accelerator",
		},
		{
			Name: "azurerm_spring_cloud_dev_tool_portal",
		},
		{
			Name: "azurerm_spring_cloud_gateway",
		},
		{
			Name: "azurerm_spring_cloud_gateway_custom_domain",
		},
		{
			Name: "azurerm_spring_cloud_gateway_route_config",
		},
		{
			Name: "azurerm_spring_cloud_java_deployment",
		},
		{
			Name: "azurerm_spring_cloud_service",
		},
		{
			Name: "azurerm_spring_cloud_storage",
		},
		{
			Name: "azurerm_sql_active_directory_administrator",
		},
		{
			Name: "azurerm_sql_database",
		},
		{
			Name: "azurerm_sql_elasticpool",
		},
		{
			Name: "azurerm_sql_failover_group",
		},
		{
			Name: "azurerm_sql_firewall_rule",
		},
		{
			Name: "azurerm_sql_managed_database",
		},
		{
			Name: "azurerm_sql_managed_instance",
		},
		{
			Name: "azurerm_sql_managed_instance_active_directory_administrator",
		},
		{
			Name: "azurerm_sql_managed_instance_failover_group",
		},
		{
			Name: "azurerm_sql_server",
		},
		{
			Name: "azurerm_sql_virtual_network_rule",
		},
		{
			Name: "azurerm_ssh_public_key",
		},
		{
			Name: "azurerm_stack_hci_cluster",
		},
		{
			Name: "azurerm_static_site",
		},
		{
			Name: "azurerm_static_site_custom_domain",
		},
		{
			Name: "azurerm_storage_account",
		},
		{
			Name: "azurerm_storage_account_customer_managed_key",
		},
		{
			Name: "azurerm_storage_account_local_user",
		},
		{
			Name: "azurerm_storage_account_network_rules",
		},
		{
			Name: "azurerm_storage_blob",
		},
		{
			Name: "azurerm_storage_blob_inventory_policy",
		},
		{
			Name: "azurerm_storage_container",
		},
		{
			Name: "azurerm_storage_data_lake_gen2_filesystem",
		},
		{
			Name: "azurerm_storage_data_lake_gen2_path",
		},
		{
			Name: "azurerm_storage_encryption_scope",
		},
		{
			Name: "azurerm_storage_management_policy",
		},
		{
			Name: "azurerm_storage_mover",
		},
		{
			Name: "azurerm_storage_mover_agent",
		},
		{
			Name: "azurerm_storage_mover_job_definition",
		},
		{
			Name: "azurerm_storage_mover_project",
		},
		{
			Name: "azurerm_storage_mover_source_endpoint",
		},
		{
			Name: "azurerm_storage_mover_target_endpoint",
		},
		{
			Name: "azurerm_storage_object_replication",
		},
		{
			Name: "azurerm_storage_queue",
		},
		{
			Name: "azurerm_storage_share",
		},
		{
			Name: "azurerm_storage_share_directory",
		},
		{
			Name: "azurerm_storage_share_file",
		},
		{
			Name: "azurerm_storage_sync",
		},
		{
			Name: "azurerm_storage_sync_cloud_endpoint",
		},
		{
			Name: "azurerm_storage_sync_group",
		},
		{
			Name: "azurerm_storage_table",
		},
		{
			Name: "azurerm_storage_table_entity",
		},
		{
			Name: "azurerm_stream_analytics_cluster",
		},
		{
			Name: "azurerm_stream_analytics_function_javascript_uda",
		},
		{
			Name: "azurerm_stream_analytics_function_javascript_udf",
		},
		{
			Name: "azurerm_stream_analytics_job",
		},
		{
			Name: "azurerm_stream_analytics_job_schedule",
		},
		{
			Name: "azurerm_stream_analytics_managed_private_endpoint",
		},
		{
			Name: "azurerm_stream_analytics_output_blob",
		},
		{
			Name: "azurerm_stream_analytics_output_cosmosdb",
		},
		{
			Name: "azurerm_stream_analytics_output_eventhub",
		},
		{
			Name: "azurerm_stream_analytics_output_function",
		},
		{
			Name: "azurerm_stream_analytics_output_mssql",
		},
		{
			Name: "azurerm_stream_analytics_output_powerbi",
		},
		{
			Name: "azurerm_stream_analytics_output_servicebus_queue",
		},
		{
			Name: "azurerm_stream_analytics_output_servicebus_topic",
		},
		{
			Name: "azurerm_stream_analytics_output_synapse",
		},
		{
			Name: "azurerm_stream_analytics_output_table",
		},
		{
			Name: "azurerm_stream_analytics_reference_input_blob",
		},
		{
			Name: "azurerm_stream_analytics_reference_input_mssql",
		},
		{
			Name: "azurerm_stream_analytics_stream_input_blob",
		},
		{
			Name: "azurerm_stream_analytics_stream_input_eventhub",
		},
		{
			Name: "azurerm_stream_analytics_stream_input_eventhub_v2",
		},
		{
			Name: "azurerm_stream_analytics_stream_input_iothub",
		},
		{
			Name: "azurerm_subnet",
		},
		{
			Name: "azurerm_subnet_nat_gateway_association",
		},
		{
			Name: "azurerm_subnet_network_security_group_association",
		},
		{
			Name: "azurerm_subnet_route_table_association",
		},
		{
			Name: "azurerm_subnet_service_endpoint_storage_policy",
		},
		{
			Name: "azurerm_subscription",
		},
		{
			Name: "azurerm_subscription_cost_management_export",
		},
		{
			Name: "azurerm_subscription_cost_management_view",
		},
		{
			Name: "azurerm_subscription_policy_assignment",
		},
		{
			Name: "azurerm_subscription_policy_exemption",
		},
		{
			Name: "azurerm_subscription_policy_remediation",
		},
		{
			Name: "azurerm_subscription_template_deployment",
		},
		{
			Name: "azurerm_synapse_firewall_rule",
		},
		{
			Name: "azurerm_synapse_integration_runtime_azure",
		},
		{
			Name: "azurerm_synapse_integration_runtime_self_hosted",
		},
		{
			Name: "azurerm_synapse_linked_service",
		},
		{
			Name: "azurerm_synapse_managed_private_endpoint",
		},
		{
			Name: "azurerm_synapse_private_link_hub",
		},
		{
			Name: "azurerm_synapse_role_assignment",
		},
		{
			Name: "azurerm_synapse_spark_pool",
		},
		{
			Name: "azurerm_synapse_sql_pool",
		},
		{
			Name: "azurerm_synapse_sql_pool_extended_auditing_policy",
		},
		{
			Name: "azurerm_synapse_sql_pool_security_alert_policy",
		},
		{
			Name: "azurerm_synapse_sql_pool_vulnerability_assessment",
		},
		{
			Name: "azurerm_synapse_sql_pool_vulnerability_assessment_baseline",
		},
		{
			Name: "azurerm_synapse_sql_pool_workload_classifier",
		},
		{
			Name: "azurerm_synapse_sql_pool_workload_group",
		},
		{
			Name: "azurerm_synapse_workspace",
		},
		{
			Name: "azurerm_synapse_workspace_aad_admin",
		},
		{
			Name: "azurerm_synapse_workspace_extended_auditing_policy",
		},
		{
			Name: "azurerm_synapse_workspace_key",
		},
		{
			Name: "azurerm_synapse_workspace_security_alert_policy",
		},
		{
			Name: "azurerm_synapse_workspace_sql_aad_admin",
		},
		{
			Name: "azurerm_synapse_workspace_vulnerability_assessment",
		},
		{
			Name: "azurerm_template_deployment",
		},
		{
			Name: "azurerm_tenant_template_deployment",
		},
		{
			Name: "azurerm_traffic_manager_azure_endpoint",
		},
		{
			Name: "azurerm_traffic_manager_external_endpoint",
		},
		{
			Name: "azurerm_traffic_manager_nested_endpoint",
		},
		{
			Name: "azurerm_traffic_manager_profile",
		},
		{
			Name: "azurerm_user_assigned_identity",
		},
		{
			Name: "azurerm_video_analyzer",
		},
		{
			Name: "azurerm_video_analyzer_edge_module",
		},
		{
			Name: "azurerm_virtual_desktop_application",
		},
		{
			Name: "azurerm_virtual_desktop_application_group",
		},
		{
			Name: "azurerm_virtual_desktop_host_pool",
		},
		{
			Name: "azurerm_virtual_desktop_host_pool_registration_info",
		},
		{
			Name: "azurerm_virtual_desktop_scaling_plan",
		},
		{
			Name: "azurerm_virtual_desktop_workspace",
		},
		{
			Name: "azurerm_virtual_desktop_workspace_application_group_association",
		},
		{
			Name: "azurerm_virtual_hub",
		},
		{
			Name: "azurerm_virtual_hub_bgp_connection",
		},
		{
			Name: "azurerm_virtual_hub_connection",
		},
		{
			Name: "azurerm_virtual_hub_ip",
		},
		{
			Name: "azurerm_virtual_hub_route_table",
		},
		{
			Name: "azurerm_virtual_hub_route_table_route",
		},
		{
			Name: "azurerm_virtual_hub_routing_intent",
		},
		{
			Name: "azurerm_virtual_hub_security_partner_provider",
		},
		{
			Name: "azurerm_virtual_machine",
		},
		{
			Name: "azurerm_virtual_machine_data_disk_attachment",
		},
		{
			Name: "azurerm_virtual_machine_extension",
		},
		{
			Name: "azurerm_virtual_machine_packet_capture",
		},
		{
			Name: "azurerm_virtual_machine_scale_set",
		},
		{
			Name: "azurerm_virtual_machine_scale_set_extension",
		},
		{
			Name: "azurerm_virtual_machine_scale_set_packet_capture",
		},
		{
			Name: "azurerm_virtual_network",
		},
		{
			Name: "azurerm_virtual_network_dns_servers",
		},
		{
			Name: "azurerm_virtual_network_gateway",
		},
		{
			Name: "azurerm_virtual_network_gateway_connection",
		},
		{
			Name: "azurerm_virtual_network_gateway_nat_rule",
		},
		{
			Name: "azurerm_virtual_network_peering",
		},
		{
			Name: "azurerm_virtual_wan",
		},
		{
			Name: "azurerm_vmware_cluster",
		},
		{
			Name: "azurerm_vmware_express_route_authorization",
		},
		{
			Name: "azurerm_vmware_netapp_volume_attachment",
		},
		{
			Name: "azurerm_vmware_private_cloud",
		},
		{
			Name: "azurerm_voice_services_communications_gateway",
		},
		{
			Name: "azurerm_voice_services_communications_gateway_test_line",
		},
		{
			Name: "azurerm_vpn_gateway",
		},
		{
			Name: "azurerm_vpn_gateway_connection",
		},
		{
			Name: "azurerm_vpn_gateway_nat_rule",
		},
		{
			Name: "azurerm_vpn_server_configuration",
		},
		{
			Name: "azurerm_vpn_server_configuration_policy_group",
		},
		{
			Name: "azurerm_vpn_site",
		},
		{
			Name: "azurerm_web_app_active_slot",
		},
		{
			Name: "azurerm_web_app_hybrid_connection",
		},
		{
			Name: "azurerm_web_application_firewall_policy",
		},
		{
			Name: "azurerm_web_pubsub",
		},
		{
			Name: "azurerm_web_pubsub_custom_certificate",
		},
		{
			Name: "azurerm_web_pubsub_custom_domain",
		},
		{
			Name: "azurerm_web_pubsub_hub",
		},
		{
			Name: "azurerm_web_pubsub_network_acl",
		},
		{
			Name: "azurerm_web_pubsub_shared_private_link_resource",
		},
		{
			Name: "azurerm_windows_function_app",
		},
		{
			Name: "azurerm_windows_function_app_slot",
		},
		{
			Name: "azurerm_windows_virtual_machine",
		},
		{
			Name: "azurerm_windows_virtual_machine_scale_set",
		},
		{
			Name: "azurerm_windows_web_app",
		},
		{
			Name: "azurerm_windows_web_app_slot",
		},
	}
	resourcesMap = map[string]int{
		"azurerm_aadb2c_directory":                                                       0,
		"azurerm_active_directory_domain_service":                                        1,
		"azurerm_active_directory_domain_service_replica_set":                            2,
		"azurerm_active_directory_domain_service_trust":                                  3,
		"azurerm_advanced_threat_protection":                                             4,
		"azurerm_analysis_services_server":                                               5,
		"azurerm_api_connection":                                                         6,
		"azurerm_api_management":                                                         7,
		"azurerm_api_management_api":                                                     8,
		"azurerm_api_management_api_diagnostic":                                          9,
		"azurerm_api_management_api_operation":                                           10,
		"azurerm_api_management_api_operation_policy":                                    11,
		"azurerm_api_management_api_operation_tag":                                       12,
		"azurerm_api_management_api_policy":                                              13,
		"azurerm_api_management_api_release":                                             14,
		"azurerm_api_management_api_schema":                                              15,
		"azurerm_api_management_api_tag":                                                 16,
		"azurerm_api_management_api_tag_description":                                     17,
		"azurerm_api_management_api_version_set":                                         18,
		"azurerm_api_management_authorization_server":                                    19,
		"azurerm_api_management_backend":                                                 20,
		"azurerm_api_management_certificate":                                             21,
		"azurerm_api_management_custom_domain":                                           22,
		"azurerm_api_management_diagnostic":                                              23,
		"azurerm_api_management_email_template":                                          24,
		"azurerm_api_management_gateway":                                                 25,
		"azurerm_api_management_gateway_api":                                             26,
		"azurerm_api_management_gateway_certificate_authority":                           27,
		"azurerm_api_management_gateway_host_name_configuration":                         28,
		"azurerm_api_management_global_schema":                                           29,
		"azurerm_api_management_group":                                                   30,
		"azurerm_api_management_group_user":                                              31,
		"azurerm_api_management_identity_provider_aad":                                   32,
		"azurerm_api_management_identity_provider_aadb2c":                                33,
		"azurerm_api_management_identity_provider_facebook":                              34,
		"azurerm_api_management_identity_provider_google":                                35,
		"azurerm_api_management_identity_provider_microsoft":                             36,
		"azurerm_api_management_identity_provider_twitter":                               37,
		"azurerm_api_management_logger":                                                  38,
		"azurerm_api_management_named_value":                                             39,
		"azurerm_api_management_notification_recipient_email":                            40,
		"azurerm_api_management_notification_recipient_user":                             41,
		"azurerm_api_management_openid_connect_provider":                                 42,
		"azurerm_api_management_policy":                                                  43,
		"azurerm_api_management_product":                                                 44,
		"azurerm_api_management_product_api":                                             45,
		"azurerm_api_management_product_group":                                           46,
		"azurerm_api_management_product_policy":                                          47,
		"azurerm_api_management_product_tag":                                             48,
		"azurerm_api_management_redis_cache":                                             49,
		"azurerm_api_management_subscription":                                            50,
		"azurerm_api_management_tag":                                                     51,
		"azurerm_api_management_user":                                                    52,
		"azurerm_app_configuration":                                                      53,
		"azurerm_app_configuration_feature":                                              54,
		"azurerm_app_configuration_key":                                                  55,
		"azurerm_app_service":                                                            56,
		"azurerm_app_service_active_slot":                                                57,
		"azurerm_app_service_certificate":                                                58,
		"azurerm_app_service_certificate_binding":                                        59,
		"azurerm_app_service_certificate_order":                                          60,
		"azurerm_app_service_connection":                                                 61,
		"azurerm_app_service_custom_hostname_binding":                                    62,
		"azurerm_app_service_environment":                                                63,
		"azurerm_app_service_environment_v3":                                             64,
		"azurerm_app_service_hybrid_connection":                                          65,
		"azurerm_app_service_managed_certificate":                                        66,
		"azurerm_app_service_plan":                                                       67,
		"azurerm_app_service_public_certificate":                                         68,
		"azurerm_app_service_slot":                                                       69,
		"azurerm_app_service_slot_custom_hostname_binding":                               70,
		"azurerm_app_service_slot_virtual_network_swift_connection":                      71,
		"azurerm_app_service_source_control":                                             72,
		"azurerm_app_service_source_control_slot":                                        73,
		"azurerm_app_service_source_control_token":                                       74,
		"azurerm_app_service_virtual_network_swift_connection":                           75,
		"azurerm_application_gateway":                                                    76,
		"azurerm_application_insights":                                                   77,
		"azurerm_application_insights_analytics_item":                                    78,
		"azurerm_application_insights_api_key":                                           79,
		"azurerm_application_insights_smart_detection_rule":                              80,
		"azurerm_application_insights_standard_web_test":                                 81,
		"azurerm_application_insights_web_test":                                          82,
		"azurerm_application_insights_workbook":                                          83,
		"azurerm_application_insights_workbook_template":                                 84,
		"azurerm_application_security_group":                                             85,
		"azurerm_arc_kubernetes_cluster":                                                 86,
		"azurerm_arc_kubernetes_cluster_extension":                                       87,
		"azurerm_arc_kubernetes_flux_configuration":                                      88,
		"azurerm_arc_machine_extension":                                                  89,
		"azurerm_arc_private_link_scope":                                                 90,
		"azurerm_attestation_provider":                                                   91,
		"azurerm_automanage_configuration":                                               92,
		"azurerm_automation_account":                                                     93,
		"azurerm_automation_certificate":                                                 94,
		"azurerm_automation_connection":                                                  95,
		"azurerm_automation_connection_certificate":                                      96,
		"azurerm_automation_connection_classic_certificate":                              97,
		"azurerm_automation_connection_service_principal":                                98,
		"azurerm_automation_connection_type":                                             99,
		"azurerm_automation_credential":                                                  100,
		"azurerm_automation_dsc_configuration":                                           101,
		"azurerm_automation_dsc_nodeconfiguration":                                       102,
		"azurerm_automation_hybrid_runbook_worker":                                       103,
		"azurerm_automation_hybrid_runbook_worker_group":                                 104,
		"azurerm_automation_job_schedule":                                                105,
		"azurerm_automation_module":                                                      106,
		"azurerm_automation_python3_package":                                             107,
		"azurerm_automation_runbook":                                                     108,
		"azurerm_automation_schedule":                                                    109,
		"azurerm_automation_software_update_configuration":                               110,
		"azurerm_automation_source_control":                                              111,
		"azurerm_automation_variable_bool":                                               112,
		"azurerm_automation_variable_datetime":                                           113,
		"azurerm_automation_variable_int":                                                114,
		"azurerm_automation_variable_object":                                             115,
		"azurerm_automation_variable_string":                                             116,
		"azurerm_automation_watcher":                                                     117,
		"azurerm_automation_webhook":                                                     118,
		"azurerm_availability_set":                                                       119,
		"azurerm_backup_container_storage_account":                                       120,
		"azurerm_backup_policy_file_share":                                               121,
		"azurerm_backup_policy_vm":                                                       122,
		"azurerm_backup_policy_vm_workload":                                              123,
		"azurerm_backup_protected_file_share":                                            124,
		"azurerm_backup_protected_vm":                                                    125,
		"azurerm_bastion_host":                                                           126,
		"azurerm_batch_account":                                                          127,
		"azurerm_batch_application":                                                      128,
		"azurerm_batch_certificate":                                                      129,
		"azurerm_batch_job":                                                              130,
		"azurerm_batch_pool":                                                             131,
		"azurerm_billing_account_cost_management_export":                                 132,
		"azurerm_blueprint_assignment":                                                   133,
		"azurerm_bot_channel_alexa":                                                      134,
		"azurerm_bot_channel_direct_line_speech":                                         135,
		"azurerm_bot_channel_directline":                                                 136,
		"azurerm_bot_channel_email":                                                      137,
		"azurerm_bot_channel_facebook":                                                   138,
		"azurerm_bot_channel_line":                                                       139,
		"azurerm_bot_channel_ms_teams":                                                   140,
		"azurerm_bot_channel_slack":                                                      141,
		"azurerm_bot_channel_sms":                                                        142,
		"azurerm_bot_channel_web_chat":                                                   143,
		"azurerm_bot_channels_registration":                                              144,
		"azurerm_bot_connection":                                                         145,
		"azurerm_bot_service_azure_bot":                                                  146,
		"azurerm_bot_web_app":                                                            147,
		"azurerm_capacity_reservation":                                                   148,
		"azurerm_capacity_reservation_group":                                             149,
		"azurerm_cdn_endpoint":                                                           150,
		"azurerm_cdn_endpoint_custom_domain":                                             151,
		"azurerm_cdn_frontdoor_custom_domain":                                            152,
		"azurerm_cdn_frontdoor_custom_domain_association":                                153,
		"azurerm_cdn_frontdoor_endpoint":                                                 154,
		"azurerm_cdn_frontdoor_firewall_policy":                                          155,
		"azurerm_cdn_frontdoor_origin":                                                   156,
		"azurerm_cdn_frontdoor_origin_group":                                             157,
		"azurerm_cdn_frontdoor_profile":                                                  158,
		"azurerm_cdn_frontdoor_route":                                                    159,
		"azurerm_cdn_frontdoor_route_disable_link_to_default_domain":                     160,
		"azurerm_cdn_frontdoor_rule":                                                     161,
		"azurerm_cdn_frontdoor_rule_set":                                                 162,
		"azurerm_cdn_frontdoor_secret":                                                   163,
		"azurerm_cdn_frontdoor_security_policy":                                          164,
		"azurerm_cdn_profile":                                                            165,
		"azurerm_cognitive_account":                                                      166,
		"azurerm_cognitive_account_customer_managed_key":                                 167,
		"azurerm_cognitive_deployment":                                                   168,
		"azurerm_communication_service":                                                  169,
		"azurerm_confidential_ledger":                                                    170,
		"azurerm_consumption_budget_management_group":                                    171,
		"azurerm_consumption_budget_resource_group":                                      172,
		"azurerm_consumption_budget_subscription":                                        173,
		"azurerm_container_app":                                                          174,
		"azurerm_container_app_environment":                                              175,
		"azurerm_container_app_environment_certificate":                                  176,
		"azurerm_container_app_environment_dapr_component":                               177,
		"azurerm_container_app_environment_storage":                                      178,
		"azurerm_container_connected_registry":                                           179,
		"azurerm_container_group":                                                        180,
		"azurerm_container_registry":                                                     181,
		"azurerm_container_registry_agent_pool":                                          182,
		"azurerm_container_registry_scope_map":                                           183,
		"azurerm_container_registry_task":                                                184,
		"azurerm_container_registry_task_schedule_run_now":                               185,
		"azurerm_container_registry_token":                                               186,
		"azurerm_container_registry_token_password":                                      187,
		"azurerm_container_registry_webhook":                                             188,
		"azurerm_cosmosdb_account":                                                       189,
		"azurerm_cosmosdb_cassandra_cluster":                                             190,
		"azurerm_cosmosdb_cassandra_datacenter":                                          191,
		"azurerm_cosmosdb_cassandra_keyspace":                                            192,
		"azurerm_cosmosdb_cassandra_table":                                               193,
		"azurerm_cosmosdb_gremlin_database":                                              194,
		"azurerm_cosmosdb_gremlin_graph":                                                 195,
		"azurerm_cosmosdb_mongo_collection":                                              196,
		"azurerm_cosmosdb_mongo_database":                                                197,
		"azurerm_cosmosdb_mongo_role_definition":                                         198,
		"azurerm_cosmosdb_mongo_user_definition":                                         199,
		"azurerm_cosmosdb_notebook_workspace":                                            200,
		"azurerm_cosmosdb_postgresql_cluster":                                            201,
		"azurerm_cosmosdb_postgresql_coordinator_configuration":                          202,
		"azurerm_cosmosdb_postgresql_firewall_rule":                                      203,
		"azurerm_cosmosdb_postgresql_node_configuration":                                 204,
		"azurerm_cosmosdb_postgresql_role":                                               205,
		"azurerm_cosmosdb_sql_container":                                                 206,
		"azurerm_cosmosdb_sql_database":                                                  207,
		"azurerm_cosmosdb_sql_dedicated_gateway":                                         208,
		"azurerm_cosmosdb_sql_function":                                                  209,
		"azurerm_cosmosdb_sql_role_assignment":                                           210,
		"azurerm_cosmosdb_sql_role_definition":                                           211,
		"azurerm_cosmosdb_sql_stored_procedure":                                          212,
		"azurerm_cosmosdb_sql_trigger":                                                   213,
		"azurerm_cosmosdb_table":                                                         214,
		"azurerm_cost_anomaly_alert":                                                     215,
		"azurerm_cost_management_scheduled_action":                                       216,
		"azurerm_custom_ip_prefix":                                                       217,
		"azurerm_custom_provider":                                                        218,
		"azurerm_dashboard":                                                              219,
		"azurerm_dashboard_grafana":                                                      220,
		"azurerm_data_factory":                                                           221,
		"azurerm_data_factory_custom_dataset":                                            222,
		"azurerm_data_factory_data_flow":                                                 223,
		"azurerm_data_factory_dataset_azure_blob":                                        224,
		"azurerm_data_factory_dataset_binary":                                            225,
		"azurerm_data_factory_dataset_cosmosdb_sqlapi":                                   226,
		"azurerm_data_factory_dataset_delimited_text":                                    227,
		"azurerm_data_factory_dataset_http":                                              228,
		"azurerm_data_factory_dataset_json":                                              229,
		"azurerm_data_factory_dataset_mysql":                                             230,
		"azurerm_data_factory_dataset_parquet":                                           231,
		"azurerm_data_factory_dataset_postgresql":                                        232,
		"azurerm_data_factory_dataset_snowflake":                                         233,
		"azurerm_data_factory_dataset_sql_server_table":                                  234,
		"azurerm_data_factory_flowlet_data_flow":                                         235,
		"azurerm_data_factory_integration_runtime_azure":                                 236,
		"azurerm_data_factory_integration_runtime_azure_ssis":                            237,
		"azurerm_data_factory_integration_runtime_managed":                               238,
		"azurerm_data_factory_integration_runtime_self_hosted":                           239,
		"azurerm_data_factory_linked_custom_service":                                     240,
		"azurerm_data_factory_linked_service_azure_blob_storage":                         241,
		"azurerm_data_factory_linked_service_azure_databricks":                           242,
		"azurerm_data_factory_linked_service_azure_file_storage":                         243,
		"azurerm_data_factory_linked_service_azure_function":                             244,
		"azurerm_data_factory_linked_service_azure_search":                               245,
		"azurerm_data_factory_linked_service_azure_sql_database":                         246,
		"azurerm_data_factory_linked_service_azure_table_storage":                        247,
		"azurerm_data_factory_linked_service_cosmosdb":                                   248,
		"azurerm_data_factory_linked_service_cosmosdb_mongoapi":                          249,
		"azurerm_data_factory_linked_service_data_lake_storage_gen2":                     250,
		"azurerm_data_factory_linked_service_key_vault":                                  251,
		"azurerm_data_factory_linked_service_kusto":                                      252,
		"azurerm_data_factory_linked_service_mysql":                                      253,
		"azurerm_data_factory_linked_service_odata":                                      254,
		"azurerm_data_factory_linked_service_odbc":                                       255,
		"azurerm_data_factory_linked_service_postgresql":                                 256,
		"azurerm_data_factory_linked_service_sftp":                                       257,
		"azurerm_data_factory_linked_service_snowflake":                                  258,
		"azurerm_data_factory_linked_service_sql_server":                                 259,
		"azurerm_data_factory_linked_service_synapse":                                    260,
		"azurerm_data_factory_linked_service_web":                                        261,
		"azurerm_data_factory_managed_private_endpoint":                                  262,
		"azurerm_data_factory_pipeline":                                                  263,
		"azurerm_data_factory_trigger_blob_event":                                        264,
		"azurerm_data_factory_trigger_custom_event":                                      265,
		"azurerm_data_factory_trigger_schedule":                                          266,
		"azurerm_data_factory_trigger_tumbling_window":                                   267,
		"azurerm_data_protection_backup_instance_blob_storage":                           268,
		"azurerm_data_protection_backup_instance_disk":                                   269,
		"azurerm_data_protection_backup_instance_postgresql":                             270,
		"azurerm_data_protection_backup_policy_blob_storage":                             271,
		"azurerm_data_protection_backup_policy_disk":                                     272,
		"azurerm_data_protection_backup_policy_postgresql":                               273,
		"azurerm_data_protection_backup_vault":                                           274,
		"azurerm_data_protection_resource_guard":                                         275,
		"azurerm_data_share":                                                             276,
		"azurerm_data_share_account":                                                     277,
		"azurerm_data_share_dataset_blob_storage":                                        278,
		"azurerm_data_share_dataset_data_lake_gen2":                                      279,
		"azurerm_data_share_dataset_kusto_cluster":                                       280,
		"azurerm_data_share_dataset_kusto_database":                                      281,
		"azurerm_database_migration_project":                                             282,
		"azurerm_database_migration_service":                                             283,
		"azurerm_databox_edge_device":                                                    284,
		"azurerm_databox_edge_order":                                                     285,
		"azurerm_databricks_access_connector":                                            286,
		"azurerm_databricks_virtual_network_peering":                                     287,
		"azurerm_databricks_workspace":                                                   288,
		"azurerm_databricks_workspace_customer_managed_key":                              289,
		"azurerm_databricks_workspace_root_dbfs_customer_managed_key":                    290,
		"azurerm_datadog_monitor":                                                        291,
		"azurerm_datadog_monitor_sso_configuration":                                      292,
		"azurerm_datadog_monitor_tag_rule":                                               293,
		"azurerm_dedicated_hardware_security_module":                                     294,
		"azurerm_dedicated_host":                                                         295,
		"azurerm_dedicated_host_group":                                                   296,
		"azurerm_dev_test_global_vm_shutdown_schedule":                                   297,
		"azurerm_dev_test_lab":                                                           298,
		"azurerm_dev_test_linux_virtual_machine":                                         299,
		"azurerm_dev_test_policy":                                                        300,
		"azurerm_dev_test_schedule":                                                      301,
		"azurerm_dev_test_virtual_network":                                               302,
		"azurerm_dev_test_windows_virtual_machine":                                       303,
		"azurerm_digital_twins_endpoint_eventgrid":                                       304,
		"azurerm_digital_twins_endpoint_eventhub":                                        305,
		"azurerm_digital_twins_endpoint_servicebus":                                      306,
		"azurerm_digital_twins_instance":                                                 307,
		"azurerm_digital_twins_time_series_database_connection":                          308,
		"azurerm_disk_access":                                                            309,
		"azurerm_disk_encryption_set":                                                    310,
		"azurerm_disk_pool":                                                              311,
		"azurerm_disk_pool_iscsi_target":                                                 312,
		"azurerm_disk_pool_iscsi_target_lun":                                             313,
		"azurerm_disk_pool_managed_disk_attachment":                                      314,
		"azurerm_dns_a_record":                                                           315,
		"azurerm_dns_aaaa_record":                                                        316,
		"azurerm_dns_caa_record":                                                         317,
		"azurerm_dns_cname_record":                                                       318,
		"azurerm_dns_mx_record":                                                          319,
		"azurerm_dns_ns_record":                                                          320,
		"azurerm_dns_ptr_record":                                                         321,
		"azurerm_dns_srv_record":                                                         322,
		"azurerm_dns_txt_record":                                                         323,
		"azurerm_dns_zone":                                                               324,
		"azurerm_elastic_cloud_elasticsearch":                                            325,
		"azurerm_email_communication_service":                                            326,
		"azurerm_eventgrid_domain":                                                       327,
		"azurerm_eventgrid_domain_topic":                                                 328,
		"azurerm_eventgrid_event_subscription":                                           329,
		"azurerm_eventgrid_system_topic":                                                 330,
		"azurerm_eventgrid_system_topic_event_subscription":                              331,
		"azurerm_eventgrid_topic":                                                        332,
		"azurerm_eventhub":                                                               333,
		"azurerm_eventhub_authorization_rule":                                            334,
		"azurerm_eventhub_cluster":                                                       335,
		"azurerm_eventhub_consumer_group":                                                336,
		"azurerm_eventhub_namespace":                                                     337,
		"azurerm_eventhub_namespace_authorization_rule":                                  338,
		"azurerm_eventhub_namespace_customer_managed_key":                                339,
		"azurerm_eventhub_namespace_disaster_recovery_config":                            340,
		"azurerm_eventhub_namespace_schema_group":                                        341,
		"azurerm_express_route_circuit":                                                  342,
		"azurerm_express_route_circuit_authorization":                                    343,
		"azurerm_express_route_circuit_connection":                                       344,
		"azurerm_express_route_circuit_peering":                                          345,
		"azurerm_express_route_connection":                                               346,
		"azurerm_express_route_gateway":                                                  347,
		"azurerm_express_route_port":                                                     348,
		"azurerm_express_route_port_authorization":                                       349,
		"azurerm_federated_identity_credential":                                          350,
		"azurerm_firewall":                                                               351,
		"azurerm_firewall_application_rule_collection":                                   352,
		"azurerm_firewall_nat_rule_collection":                                           353,
		"azurerm_firewall_network_rule_collection":                                       354,
		"azurerm_firewall_policy":                                                        355,
		"azurerm_firewall_policy_rule_collection_group":                                  356,
		"azurerm_fluid_relay_server":                                                     357,
		"azurerm_frontdoor":                                                              358,
		"azurerm_frontdoor_custom_https_configuration":                                   359,
		"azurerm_frontdoor_firewall_policy":                                              360,
		"azurerm_frontdoor_rules_engine":                                                 361,
		"azurerm_function_app":                                                           362,
		"azurerm_function_app_active_slot":                                               363,
		"azurerm_function_app_function":                                                  364,
		"azurerm_function_app_hybrid_connection":                                         365,
		"azurerm_function_app_slot":                                                      366,
		"azurerm_gallery_application":                                                    367,
		"azurerm_gallery_application_version":                                            368,
		"azurerm_graph_account":                                                          369,
		"azurerm_graph_services_account":                                                 370,
		"azurerm_hdinsight_hadoop_cluster":                                               371,
		"azurerm_hdinsight_hbase_cluster":                                                372,
		"azurerm_hdinsight_interactive_query_cluster":                                    373,
		"azurerm_hdinsight_kafka_cluster":                                                374,
		"azurerm_hdinsight_spark_cluster":                                                375,
		"azurerm_healthbot":                                                              376,
		"azurerm_healthcare_dicom_service":                                               377,
		"azurerm_healthcare_fhir_service":                                                378,
		"azurerm_healthcare_medtech_service":                                             379,
		"azurerm_healthcare_medtech_service_fhir_destination":                            380,
		"azurerm_healthcare_service":                                                     381,
		"azurerm_healthcare_workspace":                                                   382,
		"azurerm_hpc_cache":                                                              383,
		"azurerm_hpc_cache_access_policy":                                                384,
		"azurerm_hpc_cache_blob_nfs_target":                                              385,
		"azurerm_hpc_cache_blob_target":                                                  386,
		"azurerm_hpc_cache_nfs_target":                                                   387,
		"azurerm_image":                                                                  388,
		"azurerm_integration_service_environment":                                        389,
		"azurerm_iot_security_device_group":                                              390,
		"azurerm_iot_security_solution":                                                  391,
		"azurerm_iot_time_series_insights_access_policy":                                 392,
		"azurerm_iot_time_series_insights_event_source_eventhub":                         393,
		"azurerm_iot_time_series_insights_event_source_iothub":                           394,
		"azurerm_iot_time_series_insights_gen2_environment":                              395,
		"azurerm_iot_time_series_insights_reference_data_set":                            396,
		"azurerm_iot_time_series_insights_standard_environment":                          397,
		"azurerm_iotcentral_application":                                                 398,
		"azurerm_iotcentral_application_network_rule_set":                                399,
		"azurerm_iothub":                                                                 400,
		"azurerm_iothub_certificate":                                                     401,
		"azurerm_iothub_consumer_group":                                                  402,
		"azurerm_iothub_device_update_account":                                           403,
		"azurerm_iothub_device_update_instance":                                          404,
		"azurerm_iothub_dps":                                                             405,
		"azurerm_iothub_dps_certificate":                                                 406,
		"azurerm_iothub_dps_shared_access_policy":                                        407,
		"azurerm_iothub_endpoint_cosmosdb_account":                                       408,
		"azurerm_iothub_endpoint_eventhub":                                               409,
		"azurerm_iothub_endpoint_servicebus_queue":                                       410,
		"azurerm_iothub_endpoint_servicebus_topic":                                       411,
		"azurerm_iothub_endpoint_storage_container":                                      412,
		"azurerm_iothub_enrichment":                                                      413,
		"azurerm_iothub_fallback_route":                                                  414,
		"azurerm_iothub_file_upload":                                                     415,
		"azurerm_iothub_route":                                                           416,
		"azurerm_iothub_shared_access_policy":                                            417,
		"azurerm_ip_group":                                                               418,
		"azurerm_ip_group_cidr":                                                          419,
		"azurerm_key_vault":                                                              420,
		"azurerm_key_vault_access_policy":                                                421,
		"azurerm_key_vault_certificate":                                                  422,
		"azurerm_key_vault_certificate_contacts":                                         423,
		"azurerm_key_vault_certificate_issuer":                                           424,
		"azurerm_key_vault_key":                                                          425,
		"azurerm_key_vault_managed_hardware_security_module":                             426,
		"azurerm_key_vault_managed_storage_account":                                      427,
		"azurerm_key_vault_managed_storage_account_sas_token_definition":                 428,
		"azurerm_key_vault_secret":                                                       429,
		"azurerm_kubernetes_cluster":                                                     430,
		"azurerm_kubernetes_cluster_extension":                                           431,
		"azurerm_kubernetes_cluster_node_pool":                                           432,
		"azurerm_kubernetes_cluster_trusted_access_role_binding":                         433,
		"azurerm_kubernetes_fleet_manager":                                               434,
		"azurerm_kubernetes_flux_configuration":                                          435,
		"azurerm_kusto_attached_database_configuration":                                  436,
		"azurerm_kusto_cluster":                                                          437,
		"azurerm_kusto_cluster_customer_managed_key":                                     438,
		"azurerm_kusto_cluster_managed_private_endpoint":                                 439,
		"azurerm_kusto_cluster_principal_assignment":                                     440,
		"azurerm_kusto_cosmosdb_data_connection":                                         441,
		"azurerm_kusto_database":                                                         442,
		"azurerm_kusto_database_principal_assignment":                                    443,
		"azurerm_kusto_eventgrid_data_connection":                                        444,
		"azurerm_kusto_eventhub_data_connection":                                         445,
		"azurerm_kusto_iothub_data_connection":                                           446,
		"azurerm_kusto_script":                                                           447,
		"azurerm_lab_service_lab":                                                        448,
		"azurerm_lab_service_plan":                                                       449,
		"azurerm_lab_service_schedule":                                                   450,
		"azurerm_lab_service_user":                                                       451,
		"azurerm_lb":                                                                     452,
		"azurerm_lb_backend_address_pool":                                                453,
		"azurerm_lb_backend_address_pool_address":                                        454,
		"azurerm_lb_nat_pool":                                                            455,
		"azurerm_lb_nat_rule":                                                            456,
		"azurerm_lb_outbound_rule":                                                       457,
		"azurerm_lb_probe":                                                               458,
		"azurerm_lb_rule":                                                                459,
		"azurerm_lighthouse_assignment":                                                  460,
		"azurerm_lighthouse_definition":                                                  461,
		"azurerm_linux_function_app":                                                     462,
		"azurerm_linux_function_app_slot":                                                463,
		"azurerm_linux_virtual_machine":                                                  464,
		"azurerm_linux_virtual_machine_scale_set":                                        465,
		"azurerm_linux_web_app":                                                          466,
		"azurerm_linux_web_app_slot":                                                     467,
		"azurerm_load_test":                                                              468,
		"azurerm_local_network_gateway":                                                  469,
		"azurerm_log_analytics_cluster":                                                  470,
		"azurerm_log_analytics_cluster_customer_managed_key":                             471,
		"azurerm_log_analytics_data_export_rule":                                         472,
		"azurerm_log_analytics_datasource_windows_event":                                 473,
		"azurerm_log_analytics_datasource_windows_performance_counter":                   474,
		"azurerm_log_analytics_linked_service":                                           475,
		"azurerm_log_analytics_linked_storage_account":                                   476,
		"azurerm_log_analytics_query_pack":                                               477,
		"azurerm_log_analytics_query_pack_query":                                         478,
		"azurerm_log_analytics_saved_search":                                             479,
		"azurerm_log_analytics_solution":                                                 480,
		"azurerm_log_analytics_storage_insights":                                         481,
		"azurerm_log_analytics_workspace":                                                482,
		"azurerm_logic_app_action_custom":                                                483,
		"azurerm_logic_app_action_http":                                                  484,
		"azurerm_logic_app_integration_account":                                          485,
		"azurerm_logic_app_integration_account_agreement":                                486,
		"azurerm_logic_app_integration_account_assembly":                                 487,
		"azurerm_logic_app_integration_account_batch_configuration":                      488,
		"azurerm_logic_app_integration_account_certificate":                              489,
		"azurerm_logic_app_integration_account_map":                                      490,
		"azurerm_logic_app_integration_account_partner":                                  491,
		"azurerm_logic_app_integration_account_schema":                                   492,
		"azurerm_logic_app_integration_account_session":                                  493,
		"azurerm_logic_app_standard":                                                     494,
		"azurerm_logic_app_trigger_custom":                                               495,
		"azurerm_logic_app_trigger_http_request":                                         496,
		"azurerm_logic_app_trigger_recurrence":                                           497,
		"azurerm_logic_app_workflow":                                                     498,
		"azurerm_logz_monitor":                                                           499,
		"azurerm_logz_sub_account":                                                       500,
		"azurerm_logz_sub_account_tag_rule":                                              501,
		"azurerm_logz_tag_rule":                                                          502,
		"azurerm_machine_learning_compute_cluster":                                       503,
		"azurerm_machine_learning_compute_instance":                                      504,
		"azurerm_machine_learning_datastore_blobstorage":                                 505,
		"azurerm_machine_learning_datastore_datalake_gen2":                               506,
		"azurerm_machine_learning_datastore_fileshare":                                   507,
		"azurerm_machine_learning_inference_cluster":                                     508,
		"azurerm_machine_learning_synapse_spark":                                         509,
		"azurerm_machine_learning_workspace":                                             510,
		"azurerm_maintenance_assignment_dedicated_host":                                  511,
		"azurerm_maintenance_assignment_virtual_machine":                                 512,
		"azurerm_maintenance_assignment_virtual_machine_scale_set":                       513,
		"azurerm_maintenance_configuration":                                              514,
		"azurerm_managed_application":                                                    515,
		"azurerm_managed_application_definition":                                         516,
		"azurerm_managed_disk":                                                           517,
		"azurerm_managed_disk_sas_token":                                                 518,
		"azurerm_managed_lustre_file_system":                                             519,
		"azurerm_management_group":                                                       520,
		"azurerm_management_group_policy_assignment":                                     521,
		"azurerm_management_group_policy_exemption":                                      522,
		"azurerm_management_group_policy_remediation":                                    523,
		"azurerm_management_group_subscription_association":                              524,
		"azurerm_management_group_template_deployment":                                   525,
		"azurerm_management_lock":                                                        526,
		"azurerm_maps_account":                                                           527,
		"azurerm_maps_creator":                                                           528,
		"azurerm_mariadb_configuration":                                                  529,
		"azurerm_mariadb_database":                                                       530,
		"azurerm_mariadb_firewall_rule":                                                  531,
		"azurerm_mariadb_server":                                                         532,
		"azurerm_mariadb_virtual_network_rule":                                           533,
		"azurerm_marketplace_agreement":                                                  534,
		"azurerm_marketplace_role_assignment":                                            535,
		"azurerm_media_asset":                                                            536,
		"azurerm_media_asset_filter":                                                     537,
		"azurerm_media_content_key_policy":                                               538,
		"azurerm_media_job":                                                              539,
		"azurerm_media_live_event":                                                       540,
		"azurerm_media_live_event_output":                                                541,
		"azurerm_media_services_account":                                                 542,
		"azurerm_media_services_account_filter":                                          543,
		"azurerm_media_streaming_endpoint":                                               544,
		"azurerm_media_streaming_locator":                                                545,
		"azurerm_media_streaming_policy":                                                 546,
		"azurerm_media_transform":                                                        547,
		"azurerm_mobile_network":                                                         548,
		"azurerm_mobile_network_attached_data_network":                                   549,
		"azurerm_mobile_network_data_network":                                            550,
		"azurerm_mobile_network_packet_core_control_plane":                               551,
		"azurerm_mobile_network_packet_core_data_plane":                                  552,
		"azurerm_mobile_network_service":                                                 553,
		"azurerm_mobile_network_sim":                                                     554,
		"azurerm_mobile_network_sim_group":                                               555,
		"azurerm_mobile_network_sim_policy":                                              556,
		"azurerm_mobile_network_site":                                                    557,
		"azurerm_mobile_network_slice":                                                   558,
		"azurerm_monitor_aad_diagnostic_setting":                                         559,
		"azurerm_monitor_action_group":                                                   560,
		"azurerm_monitor_action_rule_action_group":                                       561,
		"azurerm_monitor_action_rule_suppression":                                        562,
		"azurerm_monitor_activity_log_alert":                                             563,
		"azurerm_monitor_alert_processing_rule_action_group":                             564,
		"azurerm_monitor_alert_processing_rule_suppression":                              565,
		"azurerm_monitor_alert_prometheus_rule_group":                                    566,
		"azurerm_monitor_autoscale_setting":                                              567,
		"azurerm_monitor_data_collection_endpoint":                                       568,
		"azurerm_monitor_data_collection_rule":                                           569,
		"azurerm_monitor_data_collection_rule_association":                               570,
		"azurerm_monitor_diagnostic_setting":                                             571,
		"azurerm_monitor_log_profile":                                                    572,
		"azurerm_monitor_metric_alert":                                                   573,
		"azurerm_monitor_private_link_scope":                                             574,
		"azurerm_monitor_private_link_scoped_service":                                    575,
		"azurerm_monitor_scheduled_query_rules_alert":                                    576,
		"azurerm_monitor_scheduled_query_rules_alert_v2":                                 577,
		"azurerm_monitor_scheduled_query_rules_log":                                      578,
		"azurerm_monitor_smart_detector_alert_rule":                                      579,
		"azurerm_monitor_workspace":                                                      580,
		"azurerm_mssql_database":                                                         581,
		"azurerm_mssql_database_extended_auditing_policy":                                582,
		"azurerm_mssql_database_vulnerability_assessment_rule_baseline":                  583,
		"azurerm_mssql_elasticpool":                                                      584,
		"azurerm_mssql_failover_group":                                                   585,
		"azurerm_mssql_firewall_rule":                                                    586,
		"azurerm_mssql_job_agent":                                                        587,
		"azurerm_mssql_job_credential":                                                   588,
		"azurerm_mssql_managed_database":                                                 589,
		"azurerm_mssql_managed_instance":                                                 590,
		"azurerm_mssql_managed_instance_active_directory_administrator":                  591,
		"azurerm_mssql_managed_instance_failover_group":                                  592,
		"azurerm_mssql_managed_instance_security_alert_policy":                           593,
		"azurerm_mssql_managed_instance_transparent_data_encryption":                     594,
		"azurerm_mssql_managed_instance_vulnerability_assessment":                        595,
		"azurerm_mssql_outbound_firewall_rule":                                           596,
		"azurerm_mssql_server":                                                           597,
		"azurerm_mssql_server_dns_alias":                                                 598,
		"azurerm_mssql_server_extended_auditing_policy":                                  599,
		"azurerm_mssql_server_microsoft_support_auditing_policy":                         600,
		"azurerm_mssql_server_security_alert_policy":                                     601,
		"azurerm_mssql_server_transparent_data_encryption":                               602,
		"azurerm_mssql_server_vulnerability_assessment":                                  603,
		"azurerm_mssql_virtual_machine":                                                  604,
		"azurerm_mssql_virtual_machine_availability_group_listener":                      605,
		"azurerm_mssql_virtual_machine_group":                                            606,
		"azurerm_mssql_virtual_network_rule":                                             607,
		"azurerm_mysql_active_directory_administrator":                                   608,
		"azurerm_mysql_configuration":                                                    609,
		"azurerm_mysql_database":                                                         610,
		"azurerm_mysql_firewall_rule":                                                    611,
		"azurerm_mysql_flexible_database":                                                612,
		"azurerm_mysql_flexible_server":                                                  613,
		"azurerm_mysql_flexible_server_active_directory_administrator":                   614,
		"azurerm_mysql_flexible_server_configuration":                                    615,
		"azurerm_mysql_flexible_server_firewall_rule":                                    616,
		"azurerm_mysql_server":                                                           617,
		"azurerm_mysql_server_key":                                                       618,
		"azurerm_mysql_virtual_network_rule":                                             619,
		"azurerm_nat_gateway":                                                            620,
		"azurerm_nat_gateway_public_ip_association":                                      621,
		"azurerm_nat_gateway_public_ip_prefix_association":                               622,
		"azurerm_netapp_account":                                                         623,
		"azurerm_netapp_pool":                                                            624,
		"azurerm_netapp_snapshot":                                                        625,
		"azurerm_netapp_snapshot_policy":                                                 626,
		"azurerm_netapp_volume":                                                          627,
		"azurerm_netapp_volume_group_sap_hana":                                           628,
		"azurerm_netapp_volume_quota_rule":                                               629,
		"azurerm_network_connection_monitor":                                             630,
		"azurerm_network_ddos_protection_plan":                                           631,
		"azurerm_network_function_azure_traffic_collector":                               632,
		"azurerm_network_interface":                                                      633,
		"azurerm_network_interface_application_gateway_backend_address_pool_association": 634,
		"azurerm_network_interface_application_security_group_association":               635,
		"azurerm_network_interface_backend_address_pool_association":                     636,
		"azurerm_network_interface_nat_rule_association":                                 637,
		"azurerm_network_interface_security_group_association":                           638,
		"azurerm_network_manager":                                                        639,
		"azurerm_network_manager_admin_rule":                                             640,
		"azurerm_network_manager_admin_rule_collection":                                  641,
		"azurerm_network_manager_connectivity_configuration":                             642,
		"azurerm_network_manager_deployment":                                             643,
		"azurerm_network_manager_management_group_connection":                            644,
		"azurerm_network_manager_network_group":                                          645,
		"azurerm_network_manager_scope_connection":                                       646,
		"azurerm_network_manager_security_admin_configuration":                           647,
		"azurerm_network_manager_static_member":                                          648,
		"azurerm_network_manager_subscription_connection":                                649,
		"azurerm_network_packet_capture":                                                 650,
		"azurerm_network_profile":                                                        651,
		"azurerm_network_security_group":                                                 652,
		"azurerm_network_security_rule":                                                  653,
		"azurerm_network_watcher":                                                        654,
		"azurerm_network_watcher_flow_log":                                               655,
		"azurerm_new_relic_monitor":                                                      656,
		"azurerm_nginx_certificate":                                                      657,
		"azurerm_nginx_configuration":                                                    658,
		"azurerm_nginx_deployment":                                                       659,
		"azurerm_notification_hub":                                                       660,
		"azurerm_notification_hub_authorization_rule":                                    661,
		"azurerm_notification_hub_namespace":                                             662,
		"azurerm_orbital_contact":                                                        663,
		"azurerm_orbital_contact_profile":                                                664,
		"azurerm_orbital_spacecraft":                                                     665,
		"azurerm_orchestrated_virtual_machine_scale_set":                                 666,
		"azurerm_palo_alto_local_rulestack":                                              667,
		"azurerm_palo_alto_local_rulestack_certificate":                                  668,
		"azurerm_palo_alto_local_rulestack_fqdn_list":                                    669,
		"azurerm_palo_alto_local_rulestack_outbound_trust_certificate_association":       670,
		"azurerm_palo_alto_local_rulestack_outbound_untrust_certificate_association":     671,
		"azurerm_palo_alto_local_rulestack_prefix_list":                                  672,
		"azurerm_palo_alto_local_rulestack_rule":                                         673,
		"azurerm_palo_alto_next_generation_firewall_virtual_hub_local_rulestack":         674,
		"azurerm_palo_alto_next_generation_firewall_virtual_hub_panorama":                675,
		"azurerm_palo_alto_next_generation_firewall_virtual_network_local_rulestack":     676,
		"azurerm_palo_alto_next_generation_firewall_virtual_network_panorama":            677,
		"azurerm_palo_alto_virtual_network_appliance":                                    678,
		"azurerm_pim_active_role_assignment":                                             679,
		"azurerm_pim_eligible_role_assignment":                                           680,
		"azurerm_point_to_site_vpn_gateway":                                              681,
		"azurerm_policy_definition":                                                      682,
		"azurerm_policy_set_definition":                                                  683,
		"azurerm_policy_virtual_machine_configuration_assignment":                        684,
		"azurerm_portal_dashboard":                                                       685,
		"azurerm_portal_tenant_configuration":                                            686,
		"azurerm_postgresql_active_directory_administrator":                              687,
		"azurerm_postgresql_configuration":                                               688,
		"azurerm_postgresql_database":                                                    689,
		"azurerm_postgresql_firewall_rule":                                               690,
		"azurerm_postgresql_flexible_server":                                             691,
		"azurerm_postgresql_flexible_server_active_directory_administrator":              692,
		"azurerm_postgresql_flexible_server_configuration":                               693,
		"azurerm_postgresql_flexible_server_database":                                    694,
		"azurerm_postgresql_flexible_server_firewall_rule":                               695,
		"azurerm_postgresql_server":                                                      696,
		"azurerm_postgresql_server_key":                                                  697,
		"azurerm_postgresql_virtual_network_rule":                                        698,
		"azurerm_powerbi_embedded":                                                       699,
		"azurerm_private_dns_a_record":                                                   700,
		"azurerm_private_dns_aaaa_record":                                                701,
		"azurerm_private_dns_cname_record":                                               702,
		"azurerm_private_dns_mx_record":                                                  703,
		"azurerm_private_dns_ptr_record":                                                 704,
		"azurerm_private_dns_resolver":                                                   705,
		"azurerm_private_dns_resolver_dns_forwarding_ruleset":                            706,
		"azurerm_private_dns_resolver_forwarding_rule":                                   707,
		"azurerm_private_dns_resolver_inbound_endpoint":                                  708,
		"azurerm_private_dns_resolver_outbound_endpoint":                                 709,
		"azurerm_private_dns_resolver_virtual_network_link":                              710,
		"azurerm_private_dns_srv_record":                                                 711,
		"azurerm_private_dns_txt_record":                                                 712,
		"azurerm_private_dns_zone":                                                       713,
		"azurerm_private_dns_zone_virtual_network_link":                                  714,
		"azurerm_private_endpoint":                                                       715,
		"azurerm_private_endpoint_application_security_group_association":                716,
		"azurerm_private_link_service":                                                   717,
		"azurerm_proximity_placement_group":                                              718,
		"azurerm_public_ip":                                                              719,
		"azurerm_public_ip_prefix":                                                       720,
		"azurerm_purview_account":                                                        721,
		"azurerm_recovery_services_vault":                                                722,
		"azurerm_recovery_services_vault_resource_guard_association":                     723,
		"azurerm_redis_cache":                                                            724,
		"azurerm_redis_enterprise_cluster":                                               725,
		"azurerm_redis_enterprise_database":                                              726,
		"azurerm_redis_firewall_rule":                                                    727,
		"azurerm_redis_linked_server":                                                    728,
		"azurerm_relay_hybrid_connection":                                                729,
		"azurerm_relay_hybrid_connection_authorization_rule":                             730,
		"azurerm_relay_namespace":                                                        731,
		"azurerm_relay_namespace_authorization_rule":                                     732,
		"azurerm_resource_deployment_script_azure_cli":                                   733,
		"azurerm_resource_deployment_script_azure_power_shell":                           734,
		"azurerm_resource_group":                                                         735,
		"azurerm_resource_group_cost_management_export":                                  736,
		"azurerm_resource_group_cost_management_view":                                    737,
		"azurerm_resource_group_policy_assignment":                                       738,
		"azurerm_resource_group_policy_exemption":                                        739,
		"azurerm_resource_group_policy_remediation":                                      740,
		"azurerm_resource_group_template_deployment":                                     741,
		"azurerm_resource_policy_assignment":                                             742,
		"azurerm_resource_policy_exemption":                                              743,
		"azurerm_resource_policy_remediation":                                            744,
		"azurerm_resource_provider_registration":                                         745,
		"azurerm_role_assignment":                                                        746,
		"azurerm_role_definition":                                                        747,
		"azurerm_route":                                                                  748,
		"azurerm_route_filter":                                                           749,
		"azurerm_route_map":                                                              750,
		"azurerm_route_server":                                                           751,
		"azurerm_route_server_bgp_connection":                                            752,
		"azurerm_route_table":                                                            753,
		"azurerm_search_service":                                                         754,
		"azurerm_search_shared_private_link_service":                                     755,
		"azurerm_security_center_assessment":                                             756,
		"azurerm_security_center_assessment_policy":                                      757,
		"azurerm_security_center_auto_provisioning":                                      758,
		"azurerm_security_center_automation":                                             759,
		"azurerm_security_center_contact":                                                760,
		"azurerm_security_center_server_vulnerability_assessment":                        761,
		"azurerm_security_center_server_vulnerability_assessment_virtual_machine":        762,
		"azurerm_security_center_setting":                                                763,
		"azurerm_security_center_subscription_pricing":                                   764,
		"azurerm_security_center_workspace":                                              765,
		"azurerm_sentinel_alert_rule_anomaly_built_in":                                   766,
		"azurerm_sentinel_alert_rule_anomaly_duplicate":                                  767,
		"azurerm_sentinel_alert_rule_fusion":                                             768,
		"azurerm_sentinel_alert_rule_machine_learning_behavior_analytics":                769,
		"azurerm_sentinel_alert_rule_ms_security_incident":                               770,
		"azurerm_sentinel_alert_rule_nrt":                                                771,
		"azurerm_sentinel_alert_rule_scheduled":                                          772,
		"azurerm_sentinel_alert_rule_threat_intelligence":                                773,
		"azurerm_sentinel_automation_rule":                                               774,
		"azurerm_sentinel_data_connector_aws_cloud_trail":                                775,
		"azurerm_sentinel_data_connector_aws_s3":                                         776,
		"azurerm_sentinel_data_connector_azure_active_directory":                         777,
		"azurerm_sentinel_data_connector_azure_advanced_threat_protection":               778,
		"azurerm_sentinel_data_connector_azure_security_center":                          779,
		"azurerm_sentinel_data_connector_dynamics_365":                                   780,
		"azurerm_sentinel_data_connector_iot":                                            781,
		"azurerm_sentinel_data_connector_microsoft_cloud_app_security":                   782,
		"azurerm_sentinel_data_connector_microsoft_defender_advanced_threat_protection":  783,
		"azurerm_sentinel_data_connector_microsoft_threat_intelligence":                  784,
		"azurerm_sentinel_data_connector_microsoft_threat_protection":                    785,
		"azurerm_sentinel_data_connector_office_365":                                     786,
		"azurerm_sentinel_data_connector_office_365_project":                             787,
		"azurerm_sentinel_data_connector_office_atp":                                     788,
		"azurerm_sentinel_data_connector_office_irm":                                     789,
		"azurerm_sentinel_data_connector_office_power_bi":                                790,
		"azurerm_sentinel_data_connector_threat_intelligence":                            791,
		"azurerm_sentinel_data_connector_threat_intelligence_taxii":                      792,
		"azurerm_sentinel_log_analytics_workspace_onboarding":                            793,
		"azurerm_sentinel_metadata":                                                      794,
		"azurerm_sentinel_threat_intelligence_indicator":                                 795,
		"azurerm_sentinel_watchlist":                                                     796,
		"azurerm_sentinel_watchlist_item":                                                797,
		"azurerm_service_fabric_cluster":                                                 798,
		"azurerm_service_fabric_managed_cluster":                                         799,
		"azurerm_service_plan":                                                           800,
		"azurerm_servicebus_namespace":                                                   801,
		"azurerm_servicebus_namespace_authorization_rule":                                802,
		"azurerm_servicebus_namespace_disaster_recovery_config":                          803,
		"azurerm_servicebus_namespace_network_rule_set":                                  804,
		"azurerm_servicebus_queue":                                                       805,
		"azurerm_servicebus_queue_authorization_rule":                                    806,
		"azurerm_servicebus_subscription":                                                807,
		"azurerm_servicebus_subscription_rule":                                           808,
		"azurerm_servicebus_topic":                                                       809,
		"azurerm_servicebus_topic_authorization_rule":                                    810,
		"azurerm_shared_image":                                                           811,
		"azurerm_shared_image_gallery":                                                   812,
		"azurerm_shared_image_version":                                                   813,
		"azurerm_signalr_service":                                                        814,
		"azurerm_signalr_service_custom_certificate":                                     815,
		"azurerm_signalr_service_custom_domain":                                          816,
		"azurerm_signalr_service_network_acl":                                            817,
		"azurerm_signalr_shared_private_link_resource":                                   818,
		"azurerm_site_recovery_fabric":                                                   819,
		"azurerm_site_recovery_hyperv_network_mapping":                                   820,
		"azurerm_site_recovery_hyperv_replication_policy":                                821,
		"azurerm_site_recovery_hyperv_replication_policy_association":                    822,
		"azurerm_site_recovery_network_mapping":                                          823,
		"azurerm_site_recovery_protection_container":                                     824,
		"azurerm_site_recovery_protection_container_mapping":                             825,
		"azurerm_site_recovery_replicated_vm":                                            826,
		"azurerm_site_recovery_replication_policy":                                       827,
		"azurerm_site_recovery_replication_recovery_plan":                                828,
		"azurerm_site_recovery_services_vault_hyperv_site":                               829,
		"azurerm_site_recovery_vmware_replication_policy":                                830,
		"azurerm_site_recovery_vmware_replication_policy_association":                    831,
		"azurerm_snapshot":                                                               832,
		"azurerm_source_control_token":                                                   833,
		"azurerm_spatial_anchors_account":                                                834,
		"azurerm_spring_cloud_accelerator":                                               835,
		"azurerm_spring_cloud_active_deployment":                                         836,
		"azurerm_spring_cloud_api_portal":                                                837,
		"azurerm_spring_cloud_api_portal_custom_domain":                                  838,
		"azurerm_spring_cloud_app":                                                       839,
		"azurerm_spring_cloud_app_cosmosdb_association":                                  840,
		"azurerm_spring_cloud_app_mysql_association":                                     841,
		"azurerm_spring_cloud_app_redis_association":                                     842,
		"azurerm_spring_cloud_application_live_view":                                     843,
		"azurerm_spring_cloud_build_deployment":                                          844,
		"azurerm_spring_cloud_build_pack_binding":                                        845,
		"azurerm_spring_cloud_builder":                                                   846,
		"azurerm_spring_cloud_certificate":                                               847,
		"azurerm_spring_cloud_configuration_service":                                     848,
		"azurerm_spring_cloud_connection":                                                849,
		"azurerm_spring_cloud_container_deployment":                                      850,
		"azurerm_spring_cloud_custom_domain":                                             851,
		"azurerm_spring_cloud_customized_accelerator":                                    852,
		"azurerm_spring_cloud_dev_tool_portal":                                           853,
		"azurerm_spring_cloud_gateway":                                                   854,
		"azurerm_spring_cloud_gateway_custom_domain":                                     855,
		"azurerm_spring_cloud_gateway_route_config":                                      856,
		"azurerm_spring_cloud_java_deployment":                                           857,
		"azurerm_spring_cloud_service":                                                   858,
		"azurerm_spring_cloud_storage":                                                   859,
		"azurerm_sql_active_directory_administrator":                                     860,
		"azurerm_sql_database":                                                           861,
		"azurerm_sql_elasticpool":                                                        862,
		"azurerm_sql_failover_group":                                                     863,
		"azurerm_sql_firewall_rule":                                                      864,
		"azurerm_sql_managed_database":                                                   865,
		"azurerm_sql_managed_instance":                                                   866,
		"azurerm_sql_managed_instance_active_directory_administrator":                    867,
		"azurerm_sql_managed_instance_failover_group":                                    868,
		"azurerm_sql_server":                                                             869,
		"azurerm_sql_virtual_network_rule":                                               870,
		"azurerm_ssh_public_key":                                                         871,
		"azurerm_stack_hci_cluster":                                                      872,
		"azurerm_static_site":                                                            873,
		"azurerm_static_site_custom_domain":                                              874,
		"azurerm_storage_account":                                                        875,
		"azurerm_storage_account_customer_managed_key":                                   876,
		"azurerm_storage_account_local_user":                                             877,
		"azurerm_storage_account_network_rules":                                          878,
		"azurerm_storage_blob":                                                           879,
		"azurerm_storage_blob_inventory_policy":                                          880,
		"azurerm_storage_container":                                                      881,
		"azurerm_storage_data_lake_gen2_filesystem":                                      882,
		"azurerm_storage_data_lake_gen2_path":                                            883,
		"azurerm_storage_encryption_scope":                                               884,
		"azurerm_storage_management_policy":                                              885,
		"azurerm_storage_mover":                                                          886,
		"azurerm_storage_mover_agent":                                                    887,
		"azurerm_storage_mover_job_definition":                                           888,
		"azurerm_storage_mover_project":                                                  889,
		"azurerm_storage_mover_source_endpoint":                                          890,
		"azurerm_storage_mover_target_endpoint":                                          891,
		"azurerm_storage_object_replication":                                             892,
		"azurerm_storage_queue":                                                          893,
		"azurerm_storage_share":                                                          894,
		"azurerm_storage_share_directory":                                                895,
		"azurerm_storage_share_file":                                                     896,
		"azurerm_storage_sync":                                                           897,
		"azurerm_storage_sync_cloud_endpoint":                                            898,
		"azurerm_storage_sync_group":                                                     899,
		"azurerm_storage_table":                                                          900,
		"azurerm_storage_table_entity":                                                   901,
		"azurerm_stream_analytics_cluster":                                               902,
		"azurerm_stream_analytics_function_javascript_uda":                               903,
		"azurerm_stream_analytics_function_javascript_udf":                               904,
		"azurerm_stream_analytics_job":                                                   905,
		"azurerm_stream_analytics_job_schedule":                                          906,
		"azurerm_stream_analytics_managed_private_endpoint":                              907,
		"azurerm_stream_analytics_output_blob":                                           908,
		"azurerm_stream_analytics_output_cosmosdb":                                       909,
		"azurerm_stream_analytics_output_eventhub":                                       910,
		"azurerm_stream_analytics_output_function":                                       911,
		"azurerm_stream_analytics_output_mssql":                                          912,
		"azurerm_stream_analytics_output_powerbi":                                        913,
		"azurerm_stream_analytics_output_servicebus_queue":                               914,
		"azurerm_stream_analytics_output_servicebus_topic":                               915,
		"azurerm_stream_analytics_output_synapse":                                        916,
		"azurerm_stream_analytics_output_table":                                          917,
		"azurerm_stream_analytics_reference_input_blob":                                  918,
		"azurerm_stream_analytics_reference_input_mssql":                                 919,
		"azurerm_stream_analytics_stream_input_blob":                                     920,
		"azurerm_stream_analytics_stream_input_eventhub":                                 921,
		"azurerm_stream_analytics_stream_input_eventhub_v2":                              922,
		"azurerm_stream_analytics_stream_input_iothub":                                   923,
		"azurerm_subnet":                                                                 924,
		"azurerm_subnet_nat_gateway_association":                                         925,
		"azurerm_subnet_network_security_group_association":                              926,
		"azurerm_subnet_route_table_association":                                         927,
		"azurerm_subnet_service_endpoint_storage_policy":                                 928,
		"azurerm_subscription":                                                           929,
		"azurerm_subscription_cost_management_export":                                    930,
		"azurerm_subscription_cost_management_view":                                      931,
		"azurerm_subscription_policy_assignment":                                         932,
		"azurerm_subscription_policy_exemption":                                          933,
		"azurerm_subscription_policy_remediation":                                        934,
		"azurerm_subscription_template_deployment":                                       935,
		"azurerm_synapse_firewall_rule":                                                  936,
		"azurerm_synapse_integration_runtime_azure":                                      937,
		"azurerm_synapse_integration_runtime_self_hosted":                                938,
		"azurerm_synapse_linked_service":                                                 939,
		"azurerm_synapse_managed_private_endpoint":                                       940,
		"azurerm_synapse_private_link_hub":                                               941,
		"azurerm_synapse_role_assignment":                                                942,
		"azurerm_synapse_spark_pool":                                                     943,
		"azurerm_synapse_sql_pool":                                                       944,
		"azurerm_synapse_sql_pool_extended_auditing_policy":                              945,
		"azurerm_synapse_sql_pool_security_alert_policy":                                 946,
		"azurerm_synapse_sql_pool_vulnerability_assessment":                              947,
		"azurerm_synapse_sql_pool_vulnerability_assessment_baseline":                     948,
		"azurerm_synapse_sql_pool_workload_classifier":                                   949,
		"azurerm_synapse_sql_pool_workload_group":                                        950,
		"azurerm_synapse_workspace":                                                      951,
		"azurerm_synapse_workspace_aad_admin":                                            952,
		"azurerm_synapse_workspace_extended_auditing_policy":                             953,
		"azurerm_synapse_workspace_key":                                                  954,
		"azurerm_synapse_workspace_security_alert_policy":                                955,
		"azurerm_synapse_workspace_sql_aad_admin":                                        956,
		"azurerm_synapse_workspace_vulnerability_assessment":                             957,
		"azurerm_template_deployment":                                                    958,
		"azurerm_tenant_template_deployment":                                             959,
		"azurerm_traffic_manager_azure_endpoint":                                         960,
		"azurerm_traffic_manager_external_endpoint":                                      961,
		"azurerm_traffic_manager_nested_endpoint":                                        962,
		"azurerm_traffic_manager_profile":                                                963,
		"azurerm_user_assigned_identity":                                                 964,
		"azurerm_video_analyzer":                                                         965,
		"azurerm_video_analyzer_edge_module":                                             966,
		"azurerm_virtual_desktop_application":                                            967,
		"azurerm_virtual_desktop_application_group":                                      968,
		"azurerm_virtual_desktop_host_pool":                                              969,
		"azurerm_virtual_desktop_host_pool_registration_info":                            970,
		"azurerm_virtual_desktop_scaling_plan":                                           971,
		"azurerm_virtual_desktop_workspace":                                              972,
		"azurerm_virtual_desktop_workspace_application_group_association":                973,
		"azurerm_virtual_hub":                                                            974,
		"azurerm_virtual_hub_bgp_connection":                                             975,
		"azurerm_virtual_hub_connection":                                                 976,
		"azurerm_virtual_hub_ip":                                                         977,
		"azurerm_virtual_hub_route_table":                                                978,
		"azurerm_virtual_hub_route_table_route":                                          979,
		"azurerm_virtual_hub_routing_intent":                                             980,
		"azurerm_virtual_hub_security_partner_provider":                                  981,
		"azurerm_virtual_machine":                                                        982,
		"azurerm_virtual_machine_data_disk_attachment":                                   983,
		"azurerm_virtual_machine_extension":                                              984,
		"azurerm_virtual_machine_packet_capture":                                         985,
		"azurerm_virtual_machine_scale_set":                                              986,
		"azurerm_virtual_machine_scale_set_extension":                                    987,
		"azurerm_virtual_machine_scale_set_packet_capture":                               988,
		"azurerm_virtual_network":                                                        989,
		"azurerm_virtual_network_dns_servers":                                            990,
		"azurerm_virtual_network_gateway":                                                991,
		"azurerm_virtual_network_gateway_connection":                                     992,
		"azurerm_virtual_network_gateway_nat_rule":                                       993,
		"azurerm_virtual_network_peering":                                                994,
		"azurerm_virtual_wan":                                                            995,
		"azurerm_vmware_cluster":                                                         996,
		"azurerm_vmware_express_route_authorization":                                     997,
		"azurerm_vmware_netapp_volume_attachment":                                        998,
		"azurerm_vmware_private_cloud":                                                   999,
		"azurerm_voice_services_communications_gateway":                                  1000,
		"azurerm_voice_services_communications_gateway_test_line":                        1001,
		"azurerm_vpn_gateway":                                                            1002,
		"azurerm_vpn_gateway_connection":                                                 1003,
		"azurerm_vpn_gateway_nat_rule":                                                   1004,
		"azurerm_vpn_server_configuration":                                               1005,
		"azurerm_vpn_server_configuration_policy_group":                                  1006,
		"azurerm_vpn_site":                                                               1007,
		"azurerm_web_app_active_slot":                                                    1008,
		"azurerm_web_app_hybrid_connection":                                              1009,
		"azurerm_web_application_firewall_policy":                                        1010,
		"azurerm_web_pubsub":                                                             1011,
		"azurerm_web_pubsub_custom_certificate":                                          1012,
		"azurerm_web_pubsub_custom_domain":                                               1013,
		"azurerm_web_pubsub_hub":                                                         1014,
		"azurerm_web_pubsub_network_acl":                                                 1015,
		"azurerm_web_pubsub_shared_private_link_resource":                                1016,
		"azurerm_windows_function_app":                                                   1017,
		"azurerm_windows_function_app_slot":                                              1018,
		"azurerm_windows_virtual_machine":                                                1019,
		"azurerm_windows_virtual_machine_scale_set":                                      1020,
		"azurerm_windows_web_app":                                                        1021,
		"azurerm_windows_web_app_slot":                                                   1022,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
