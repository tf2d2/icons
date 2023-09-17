package aws

import (
	"fmt"

	"github.com/tf2d2/icons/internal/resource"
)

var (
	Resources = []*resource.Resource{
		{
			Name:    "aws_accessanalyzer_analyzer",
			IconURL: "https://raw.githubusercontent.com/tf2d2/icons/main/aws/resource/Security-Identity-Compliance/AWS-Identity-Access-Management_IAM-Access-Analyzer.svg",
		},
		{
			Name: "aws_accessanalyzer_archive_rule",
		},
		{
			Name: "aws_account_alternate_contact",
		},
		{
			Name: "aws_account_primary_contact",
		},
		{
			Name: "aws_acm_certificate",
		},
		{
			Name: "aws_acm_certificate_validation",
		},
		{
			Name: "aws_acmpca_certificate",
		},
		{
			Name: "aws_acmpca_certificate_authority",
		},
		{
			Name: "aws_acmpca_certificate_authority_certificate",
		},
		{
			Name: "aws_acmpca_permission",
		},
		{
			Name: "aws_acmpca_policy",
		},
		{
			Name: "aws_alb",
		},
		{
			Name: "aws_alb_listener",
		},
		{
			Name: "aws_alb_listener_certificate",
		},
		{
			Name: "aws_alb_listener_rule",
		},
		{
			Name: "aws_alb_target_group",
		},
		{
			Name: "aws_alb_target_group_attachment",
		},
		{
			Name: "aws_ami",
		},
		{
			Name: "aws_ami_copy",
		},
		{
			Name: "aws_ami_from_instance",
		},
		{
			Name: "aws_ami_launch_permission",
		},
		{
			Name: "aws_amplify_app",
		},
		{
			Name: "aws_amplify_backend_environment",
		},
		{
			Name: "aws_amplify_branch",
		},
		{
			Name: "aws_amplify_domain_association",
		},
		{
			Name: "aws_amplify_webhook",
		},
		{
			Name: "aws_api_gateway_account",
		},
		{
			Name: "aws_api_gateway_api_key",
		},
		{
			Name: "aws_api_gateway_authorizer",
		},
		{
			Name: "aws_api_gateway_base_path_mapping",
		},
		{
			Name: "aws_api_gateway_client_certificate",
		},
		{
			Name: "aws_api_gateway_deployment",
		},
		{
			Name: "aws_api_gateway_documentation_part",
		},
		{
			Name: "aws_api_gateway_documentation_version",
		},
		{
			Name: "aws_api_gateway_domain_name",
		},
		{
			Name: "aws_api_gateway_gateway_response",
		},
		{
			Name: "aws_api_gateway_integration",
		},
		{
			Name: "aws_api_gateway_integration_response",
		},
		{
			Name: "aws_api_gateway_method",
		},
		{
			Name: "aws_api_gateway_method_response",
		},
		{
			Name: "aws_api_gateway_method_settings",
		},
		{
			Name: "aws_api_gateway_model",
		},
		{
			Name: "aws_api_gateway_request_validator",
		},
		{
			Name: "aws_api_gateway_resource",
		},
		{
			Name: "aws_api_gateway_rest_api",
		},
		{
			Name: "aws_api_gateway_rest_api_policy",
		},
		{
			Name: "aws_api_gateway_stage",
		},
		{
			Name: "aws_api_gateway_usage_plan",
		},
		{
			Name: "aws_api_gateway_usage_plan_key",
		},
		{
			Name: "aws_api_gateway_vpc_link",
		},
		{
			Name: "aws_apigatewayv2_api",
		},
		{
			Name: "aws_apigatewayv2_api_mapping",
		},
		{
			Name: "aws_apigatewayv2_authorizer",
		},
		{
			Name: "aws_apigatewayv2_deployment",
		},
		{
			Name: "aws_apigatewayv2_domain_name",
		},
		{
			Name: "aws_apigatewayv2_integration",
		},
		{
			Name: "aws_apigatewayv2_integration_response",
		},
		{
			Name: "aws_apigatewayv2_model",
		},
		{
			Name: "aws_apigatewayv2_route",
		},
		{
			Name: "aws_apigatewayv2_route_response",
		},
		{
			Name: "aws_apigatewayv2_stage",
		},
		{
			Name: "aws_apigatewayv2_vpc_link",
		},
		{
			Name: "aws_app_cookie_stickiness_policy",
		},
		{
			Name: "aws_appautoscaling_policy",
		},
		{
			Name: "aws_appautoscaling_scheduled_action",
		},
		{
			Name: "aws_appautoscaling_target",
		},
		{
			Name: "aws_appconfig_application",
		},
		{
			Name: "aws_appconfig_configuration_profile",
		},
		{
			Name: "aws_appconfig_deployment",
		},
		{
			Name: "aws_appconfig_deployment_strategy",
		},
		{
			Name: "aws_appconfig_environment",
		},
		{
			Name: "aws_appconfig_extension",
		},
		{
			Name: "aws_appconfig_extension_association",
		},
		{
			Name: "aws_appconfig_hosted_configuration_version",
		},
		{
			Name: "aws_appflow_connector_profile",
		},
		{
			Name: "aws_appflow_flow",
		},
		{
			Name: "aws_appintegrations_data_integration",
		},
		{
			Name: "aws_appintegrations_event_integration",
		},
		{
			Name: "aws_applicationinsights_application",
		},
		{
			Name: "aws_appmesh_gateway_route",
		},
		{
			Name: "aws_appmesh_mesh",
		},
		{
			Name: "aws_appmesh_route",
		},
		{
			Name: "aws_appmesh_virtual_gateway",
		},
		{
			Name: "aws_appmesh_virtual_node",
		},
		{
			Name: "aws_appmesh_virtual_router",
		},
		{
			Name: "aws_appmesh_virtual_service",
		},
		{
			Name: "aws_apprunner_auto_scaling_configuration_version",
		},
		{
			Name: "aws_apprunner_connection",
		},
		{
			Name: "aws_apprunner_custom_domain_association",
		},
		{
			Name: "aws_apprunner_observability_configuration",
		},
		{
			Name: "aws_apprunner_service",
		},
		{
			Name: "aws_apprunner_vpc_connector",
		},
		{
			Name: "aws_apprunner_vpc_ingress_connection",
		},
		{
			Name: "aws_appstream_directory_config",
		},
		{
			Name: "aws_appstream_fleet",
		},
		{
			Name: "aws_appstream_fleet_stack_association",
		},
		{
			Name: "aws_appstream_image_builder",
		},
		{
			Name: "aws_appstream_stack",
		},
		{
			Name: "aws_appstream_user",
		},
		{
			Name: "aws_appstream_user_stack_association",
		},
		{
			Name: "aws_appsync_api_cache",
		},
		{
			Name: "aws_appsync_api_key",
		},
		{
			Name: "aws_appsync_datasource",
		},
		{
			Name: "aws_appsync_domain_name",
		},
		{
			Name: "aws_appsync_domain_name_api_association",
		},
		{
			Name: "aws_appsync_function",
		},
		{
			Name: "aws_appsync_graphql_api",
		},
		{
			Name: "aws_appsync_resolver",
		},
		{
			Name: "aws_appsync_type",
		},
		{
			Name: "aws_athena_data_catalog",
		},
		{
			Name: "aws_athena_database",
		},
		{
			Name: "aws_athena_named_query",
		},
		{
			Name: "aws_athena_workgroup",
		},
		{
			Name: "aws_auditmanager_account_registration",
		},
		{
			Name: "aws_auditmanager_assessment",
		},
		{
			Name: "aws_auditmanager_assessment_delegation",
		},
		{
			Name: "aws_auditmanager_assessment_report",
		},
		{
			Name: "aws_auditmanager_control",
		},
		{
			Name: "aws_auditmanager_framework",
		},
		{
			Name: "aws_auditmanager_framework_share",
		},
		{
			Name: "aws_auditmanager_organization_admin_account_registration",
		},
		{
			Name: "aws_autoscaling_attachment",
		},
		{
			Name: "aws_autoscaling_group",
		},
		{
			Name: "aws_autoscaling_group_tag",
		},
		{
			Name: "aws_autoscaling_lifecycle_hook",
		},
		{
			Name: "aws_autoscaling_notification",
		},
		{
			Name: "aws_autoscaling_policy",
		},
		{
			Name: "aws_autoscaling_schedule",
		},
		{
			Name: "aws_autoscaling_traffic_source_attachment",
		},
		{
			Name: "aws_autoscalingplans_scaling_plan",
		},
		{
			Name: "aws_backup_framework",
		},
		{
			Name: "aws_backup_global_settings",
		},
		{
			Name: "aws_backup_plan",
		},
		{
			Name: "aws_backup_region_settings",
		},
		{
			Name: "aws_backup_report_plan",
		},
		{
			Name: "aws_backup_selection",
		},
		{
			Name: "aws_backup_vault",
		},
		{
			Name: "aws_backup_vault_lock_configuration",
		},
		{
			Name: "aws_backup_vault_notifications",
		},
		{
			Name: "aws_backup_vault_policy",
		},
		{
			Name: "aws_batch_compute_environment",
		},
		{
			Name: "aws_batch_job_definition",
		},
		{
			Name: "aws_batch_job_queue",
		},
		{
			Name: "aws_batch_scheduling_policy",
		},
		{
			Name: "aws_budgets_budget",
		},
		{
			Name: "aws_budgets_budget_action",
		},
		{
			Name: "aws_ce_anomaly_monitor",
		},
		{
			Name: "aws_ce_anomaly_subscription",
		},
		{
			Name: "aws_ce_cost_allocation_tag",
		},
		{
			Name: "aws_ce_cost_category",
		},
		{
			Name: "aws_chime_voice_connector",
		},
		{
			Name: "aws_chime_voice_connector_group",
		},
		{
			Name: "aws_chime_voice_connector_logging",
		},
		{
			Name: "aws_chime_voice_connector_origination",
		},
		{
			Name: "aws_chime_voice_connector_streaming",
		},
		{
			Name: "aws_chime_voice_connector_termination",
		},
		{
			Name: "aws_chime_voice_connector_termination_credentials",
		},
		{
			Name: "aws_chimesdkmediapipelines_media_insights_pipeline_configuration",
		},
		{
			Name: "aws_chimesdkvoice_global_settings",
		},
		{
			Name: "aws_chimesdkvoice_sip_media_application",
		},
		{
			Name: "aws_chimesdkvoice_sip_rule",
		},
		{
			Name: "aws_chimesdkvoice_voice_profile_domain",
		},
		{
			Name: "aws_cleanrooms_collaboration",
		},
		{
			Name: "aws_cloud9_environment_ec2",
		},
		{
			Name: "aws_cloud9_environment_membership",
		},
		{
			Name: "aws_cloudcontrolapi_resource",
		},
		{
			Name: "aws_cloudformation_stack",
		},
		{
			Name: "aws_cloudformation_stack_set",
		},
		{
			Name: "aws_cloudformation_stack_set_instance",
		},
		{
			Name: "aws_cloudformation_type",
		},
		{
			Name: "aws_cloudfront_cache_policy",
		},
		{
			Name: "aws_cloudfront_continuous_deployment_policy",
		},
		{
			Name: "aws_cloudfront_distribution",
		},
		{
			Name: "aws_cloudfront_field_level_encryption_config",
		},
		{
			Name: "aws_cloudfront_field_level_encryption_profile",
		},
		{
			Name: "aws_cloudfront_function",
		},
		{
			Name: "aws_cloudfront_key_group",
		},
		{
			Name: "aws_cloudfront_monitoring_subscription",
		},
		{
			Name: "aws_cloudfront_origin_access_control",
		},
		{
			Name: "aws_cloudfront_origin_access_identity",
		},
		{
			Name: "aws_cloudfront_origin_request_policy",
		},
		{
			Name: "aws_cloudfront_public_key",
		},
		{
			Name: "aws_cloudfront_realtime_log_config",
		},
		{
			Name: "aws_cloudfront_response_headers_policy",
		},
		{
			Name: "aws_cloudhsm_v2_cluster",
		},
		{
			Name: "aws_cloudhsm_v2_hsm",
		},
		{
			Name: "aws_cloudsearch_domain",
		},
		{
			Name: "aws_cloudsearch_domain_service_access_policy",
		},
		{
			Name: "aws_cloudtrail",
		},
		{
			Name: "aws_cloudtrail_event_data_store",
		},
		{
			Name: "aws_cloudwatch_composite_alarm",
		},
		{
			Name: "aws_cloudwatch_dashboard",
		},
		{
			Name: "aws_cloudwatch_event_api_destination",
		},
		{
			Name: "aws_cloudwatch_event_archive",
		},
		{
			Name: "aws_cloudwatch_event_bus",
		},
		{
			Name: "aws_cloudwatch_event_bus_policy",
		},
		{
			Name: "aws_cloudwatch_event_connection",
		},
		{
			Name: "aws_cloudwatch_event_endpoint",
		},
		{
			Name: "aws_cloudwatch_event_permission",
		},
		{
			Name: "aws_cloudwatch_event_rule",
		},
		{
			Name: "aws_cloudwatch_event_target",
		},
		{
			Name: "aws_cloudwatch_log_data_protection_policy",
		},
		{
			Name: "aws_cloudwatch_log_destination",
		},
		{
			Name: "aws_cloudwatch_log_destination_policy",
		},
		{
			Name: "aws_cloudwatch_log_group",
		},
		{
			Name: "aws_cloudwatch_log_metric_filter",
		},
		{
			Name: "aws_cloudwatch_log_resource_policy",
		},
		{
			Name: "aws_cloudwatch_log_stream",
		},
		{
			Name: "aws_cloudwatch_log_subscription_filter",
		},
		{
			Name: "aws_cloudwatch_metric_alarm",
		},
		{
			Name: "aws_cloudwatch_metric_stream",
		},
		{
			Name: "aws_cloudwatch_query_definition",
		},
		{
			Name: "aws_codeartifact_domain",
		},
		{
			Name: "aws_codeartifact_domain_permissions_policy",
		},
		{
			Name: "aws_codeartifact_repository",
		},
		{
			Name: "aws_codeartifact_repository_permissions_policy",
		},
		{
			Name: "aws_codebuild_project",
		},
		{
			Name: "aws_codebuild_report_group",
		},
		{
			Name: "aws_codebuild_resource_policy",
		},
		{
			Name: "aws_codebuild_source_credential",
		},
		{
			Name: "aws_codebuild_webhook",
		},
		{
			Name: "aws_codecatalyst_dev_environment",
		},
		{
			Name: "aws_codecatalyst_project",
		},
		{
			Name: "aws_codecatalyst_source_repository",
		},
		{
			Name: "aws_codecommit_approval_rule_template",
		},
		{
			Name: "aws_codecommit_approval_rule_template_association",
		},
		{
			Name: "aws_codecommit_repository",
		},
		{
			Name: "aws_codecommit_trigger",
		},
		{
			Name: "aws_codedeploy_app",
		},
		{
			Name: "aws_codedeploy_deployment_config",
		},
		{
			Name: "aws_codedeploy_deployment_group",
		},
		{
			Name: "aws_codegurureviewer_repository_association",
		},
		{
			Name: "aws_codepipeline",
		},
		{
			Name: "aws_codepipeline_custom_action_type",
		},
		{
			Name: "aws_codepipeline_webhook",
		},
		{
			Name: "aws_codestarconnections_connection",
		},
		{
			Name: "aws_codestarconnections_host",
		},
		{
			Name: "aws_codestarnotifications_notification_rule",
		},
		{
			Name: "aws_cognito_identity_pool",
		},
		{
			Name: "aws_cognito_identity_pool_provider_principal_tag",
		},
		{
			Name: "aws_cognito_identity_pool_roles_attachment",
		},
		{
			Name: "aws_cognito_identity_provider",
		},
		{
			Name: "aws_cognito_managed_user_pool_client",
		},
		{
			Name: "aws_cognito_resource_server",
		},
		{
			Name: "aws_cognito_risk_configuration",
		},
		{
			Name: "aws_cognito_user",
		},
		{
			Name: "aws_cognito_user_group",
		},
		{
			Name: "aws_cognito_user_in_group",
		},
		{
			Name: "aws_cognito_user_pool",
		},
		{
			Name: "aws_cognito_user_pool_client",
		},
		{
			Name: "aws_cognito_user_pool_domain",
		},
		{
			Name: "aws_cognito_user_pool_ui_customization",
		},
		{
			Name: "aws_comprehend_document_classifier",
		},
		{
			Name: "aws_comprehend_entity_recognizer",
		},
		{
			Name: "aws_config_aggregate_authorization",
		},
		{
			Name: "aws_config_config_rule",
		},
		{
			Name: "aws_config_configuration_aggregator",
		},
		{
			Name: "aws_config_configuration_recorder",
		},
		{
			Name: "aws_config_configuration_recorder_status",
		},
		{
			Name: "aws_config_conformance_pack",
		},
		{
			Name: "aws_config_delivery_channel",
		},
		{
			Name: "aws_config_organization_conformance_pack",
		},
		{
			Name: "aws_config_organization_custom_policy_rule",
		},
		{
			Name: "aws_config_organization_custom_rule",
		},
		{
			Name: "aws_config_organization_managed_rule",
		},
		{
			Name: "aws_config_remediation_configuration",
		},
		{
			Name: "aws_connect_bot_association",
		},
		{
			Name: "aws_connect_contact_flow",
		},
		{
			Name: "aws_connect_contact_flow_module",
		},
		{
			Name: "aws_connect_hours_of_operation",
		},
		{
			Name: "aws_connect_instance",
		},
		{
			Name: "aws_connect_instance_storage_config",
		},
		{
			Name: "aws_connect_lambda_function_association",
		},
		{
			Name: "aws_connect_phone_number",
		},
		{
			Name: "aws_connect_queue",
		},
		{
			Name: "aws_connect_quick_connect",
		},
		{
			Name: "aws_connect_routing_profile",
		},
		{
			Name: "aws_connect_security_profile",
		},
		{
			Name: "aws_connect_user",
		},
		{
			Name: "aws_connect_user_hierarchy_group",
		},
		{
			Name: "aws_connect_user_hierarchy_structure",
		},
		{
			Name: "aws_connect_vocabulary",
		},
		{
			Name: "aws_controltower_control",
		},
		{
			Name: "aws_cur_report_definition",
		},
		{
			Name: "aws_customer_gateway",
		},
		{
			Name: "aws_dataexchange_data_set",
		},
		{
			Name: "aws_dataexchange_revision",
		},
		{
			Name: "aws_datapipeline_pipeline",
		},
		{
			Name: "aws_datapipeline_pipeline_definition",
		},
		{
			Name: "aws_datasync_agent",
		},
		{
			Name: "aws_datasync_location_azure_blob",
		},
		{
			Name: "aws_datasync_location_efs",
		},
		{
			Name: "aws_datasync_location_fsx_lustre_file_system",
		},
		{
			Name: "aws_datasync_location_fsx_ontap_file_system",
		},
		{
			Name: "aws_datasync_location_fsx_openzfs_file_system",
		},
		{
			Name: "aws_datasync_location_fsx_windows_file_system",
		},
		{
			Name: "aws_datasync_location_hdfs",
		},
		{
			Name: "aws_datasync_location_nfs",
		},
		{
			Name: "aws_datasync_location_object_storage",
		},
		{
			Name: "aws_datasync_location_s3",
		},
		{
			Name: "aws_datasync_location_smb",
		},
		{
			Name: "aws_datasync_task",
		},
		{
			Name: "aws_dax_cluster",
		},
		{
			Name: "aws_dax_parameter_group",
		},
		{
			Name: "aws_dax_subnet_group",
		},
		{
			Name: "aws_db_cluster_snapshot",
		},
		{
			Name: "aws_db_event_subscription",
		},
		{
			Name: "aws_db_instance",
		},
		{
			Name: "aws_db_instance_automated_backups_replication",
		},
		{
			Name: "aws_db_instance_role_association",
		},
		{
			Name: "aws_db_option_group",
		},
		{
			Name: "aws_db_parameter_group",
		},
		{
			Name: "aws_db_proxy",
		},
		{
			Name: "aws_db_proxy_default_target_group",
		},
		{
			Name: "aws_db_proxy_endpoint",
		},
		{
			Name: "aws_db_proxy_target",
		},
		{
			Name: "aws_db_snapshot",
		},
		{
			Name: "aws_db_snapshot_copy",
		},
		{
			Name: "aws_db_subnet_group",
		},
		{
			Name: "aws_default_network_acl",
		},
		{
			Name: "aws_default_route_table",
		},
		{
			Name: "aws_default_security_group",
		},
		{
			Name: "aws_default_subnet",
		},
		{
			Name: "aws_default_vpc",
		},
		{
			Name: "aws_default_vpc_dhcp_options",
		},
		{
			Name: "aws_detective_graph",
		},
		{
			Name: "aws_detective_invitation_accepter",
		},
		{
			Name: "aws_detective_member",
		},
		{
			Name: "aws_devicefarm_device_pool",
		},
		{
			Name: "aws_devicefarm_instance_profile",
		},
		{
			Name: "aws_devicefarm_network_profile",
		},
		{
			Name: "aws_devicefarm_project",
		},
		{
			Name: "aws_devicefarm_test_grid_project",
		},
		{
			Name: "aws_devicefarm_upload",
		},
		{
			Name: "aws_directory_service_conditional_forwarder",
		},
		{
			Name: "aws_directory_service_directory",
		},
		{
			Name: "aws_directory_service_log_subscription",
		},
		{
			Name: "aws_directory_service_radius_settings",
		},
		{
			Name: "aws_directory_service_region",
		},
		{
			Name: "aws_directory_service_shared_directory",
		},
		{
			Name: "aws_directory_service_shared_directory_accepter",
		},
		{
			Name: "aws_directory_service_trust",
		},
		{
			Name: "aws_dlm_lifecycle_policy",
		},
		{
			Name: "aws_dms_certificate",
		},
		{
			Name: "aws_dms_endpoint",
		},
		{
			Name: "aws_dms_event_subscription",
		},
		{
			Name: "aws_dms_replication_instance",
		},
		{
			Name: "aws_dms_replication_subnet_group",
		},
		{
			Name: "aws_dms_replication_task",
		},
		{
			Name: "aws_dms_s3_endpoint",
		},
		{
			Name: "aws_docdb_cluster",
		},
		{
			Name: "aws_docdb_cluster_instance",
		},
		{
			Name: "aws_docdb_cluster_parameter_group",
		},
		{
			Name: "aws_docdb_cluster_snapshot",
		},
		{
			Name: "aws_docdb_event_subscription",
		},
		{
			Name: "aws_docdb_global_cluster",
		},
		{
			Name: "aws_docdb_subnet_group",
		},
		{
			Name: "aws_dx_bgp_peer",
		},
		{
			Name: "aws_dx_connection",
		},
		{
			Name: "aws_dx_connection_association",
		},
		{
			Name: "aws_dx_connection_confirmation",
		},
		{
			Name: "aws_dx_gateway",
		},
		{
			Name: "aws_dx_gateway_association",
		},
		{
			Name: "aws_dx_gateway_association_proposal",
		},
		{
			Name: "aws_dx_hosted_connection",
		},
		{
			Name: "aws_dx_hosted_private_virtual_interface",
		},
		{
			Name: "aws_dx_hosted_private_virtual_interface_accepter",
		},
		{
			Name: "aws_dx_hosted_public_virtual_interface",
		},
		{
			Name: "aws_dx_hosted_public_virtual_interface_accepter",
		},
		{
			Name: "aws_dx_hosted_transit_virtual_interface",
		},
		{
			Name: "aws_dx_hosted_transit_virtual_interface_accepter",
		},
		{
			Name: "aws_dx_lag",
		},
		{
			Name: "aws_dx_macsec_key_association",
		},
		{
			Name: "aws_dx_private_virtual_interface",
		},
		{
			Name: "aws_dx_public_virtual_interface",
		},
		{
			Name: "aws_dx_transit_virtual_interface",
		},
		{
			Name: "aws_dynamodb_contributor_insights",
		},
		{
			Name: "aws_dynamodb_global_table",
		},
		{
			Name: "aws_dynamodb_kinesis_streaming_destination",
		},
		{
			Name: "aws_dynamodb_table",
		},
		{
			Name: "aws_dynamodb_table_item",
		},
		{
			Name: "aws_dynamodb_table_replica",
		},
		{
			Name: "aws_dynamodb_tag",
		},
		{
			Name: "aws_ebs_default_kms_key",
		},
		{
			Name: "aws_ebs_encryption_by_default",
		},
		{
			Name: "aws_ebs_snapshot",
		},
		{
			Name: "aws_ebs_snapshot_copy",
		},
		{
			Name: "aws_ebs_snapshot_import",
		},
		{
			Name: "aws_ebs_volume",
		},
		{
			Name: "aws_ec2_availability_zone_group",
		},
		{
			Name: "aws_ec2_capacity_reservation",
		},
		{
			Name: "aws_ec2_carrier_gateway",
		},
		{
			Name: "aws_ec2_client_vpn_authorization_rule",
		},
		{
			Name: "aws_ec2_client_vpn_endpoint",
		},
		{
			Name: "aws_ec2_client_vpn_network_association",
		},
		{
			Name: "aws_ec2_client_vpn_route",
		},
		{
			Name: "aws_ec2_fleet",
		},
		{
			Name: "aws_ec2_host",
		},
		{
			Name: "aws_ec2_instance_connect_endpoint",
		},
		{
			Name: "aws_ec2_instance_state",
		},
		{
			Name: "aws_ec2_local_gateway_route",
		},
		{
			Name: "aws_ec2_local_gateway_route_table_vpc_association",
		},
		{
			Name: "aws_ec2_managed_prefix_list",
		},
		{
			Name: "aws_ec2_managed_prefix_list_entry",
		},
		{
			Name: "aws_ec2_network_insights_analysis",
		},
		{
			Name: "aws_ec2_network_insights_path",
		},
		{
			Name: "aws_ec2_serial_console_access",
		},
		{
			Name: "aws_ec2_subnet_cidr_reservation",
		},
		{
			Name: "aws_ec2_tag",
		},
		{
			Name: "aws_ec2_traffic_mirror_filter",
		},
		{
			Name: "aws_ec2_traffic_mirror_filter_rule",
		},
		{
			Name: "aws_ec2_traffic_mirror_session",
		},
		{
			Name: "aws_ec2_traffic_mirror_target",
		},
		{
			Name: "aws_ec2_transit_gateway",
		},
		{
			Name: "aws_ec2_transit_gateway_connect",
		},
		{
			Name: "aws_ec2_transit_gateway_connect_peer",
		},
		{
			Name: "aws_ec2_transit_gateway_multicast_domain",
		},
		{
			Name: "aws_ec2_transit_gateway_multicast_domain_association",
		},
		{
			Name: "aws_ec2_transit_gateway_multicast_group_member",
		},
		{
			Name: "aws_ec2_transit_gateway_multicast_group_source",
		},
		{
			Name: "aws_ec2_transit_gateway_peering_attachment",
		},
		{
			Name: "aws_ec2_transit_gateway_peering_attachment_accepter",
		},
		{
			Name: "aws_ec2_transit_gateway_policy_table",
		},
		{
			Name: "aws_ec2_transit_gateway_policy_table_association",
		},
		{
			Name: "aws_ec2_transit_gateway_prefix_list_reference",
		},
		{
			Name: "aws_ec2_transit_gateway_route",
		},
		{
			Name: "aws_ec2_transit_gateway_route_table",
		},
		{
			Name: "aws_ec2_transit_gateway_route_table_association",
		},
		{
			Name: "aws_ec2_transit_gateway_route_table_propagation",
		},
		{
			Name: "aws_ec2_transit_gateway_vpc_attachment",
		},
		{
			Name: "aws_ec2_transit_gateway_vpc_attachment_accepter",
		},
		{
			Name: "aws_ecr_lifecycle_policy",
		},
		{
			Name: "aws_ecr_pull_through_cache_rule",
		},
		{
			Name: "aws_ecr_registry_policy",
		},
		{
			Name: "aws_ecr_registry_scanning_configuration",
		},
		{
			Name: "aws_ecr_replication_configuration",
		},
		{
			Name: "aws_ecr_repository",
		},
		{
			Name: "aws_ecr_repository_policy",
		},
		{
			Name: "aws_ecrpublic_repository",
		},
		{
			Name: "aws_ecrpublic_repository_policy",
		},
		{
			Name: "aws_ecs_account_setting_default",
		},
		{
			Name: "aws_ecs_capacity_provider",
		},
		{
			Name: "aws_ecs_cluster",
		},
		{
			Name: "aws_ecs_cluster_capacity_providers",
		},
		{
			Name: "aws_ecs_service",
		},
		{
			Name: "aws_ecs_tag",
		},
		{
			Name: "aws_ecs_task_definition",
		},
		{
			Name: "aws_ecs_task_set",
		},
		{
			Name: "aws_efs_access_point",
		},
		{
			Name: "aws_efs_backup_policy",
		},
		{
			Name: "aws_efs_file_system",
		},
		{
			Name: "aws_efs_file_system_policy",
		},
		{
			Name: "aws_efs_mount_target",
		},
		{
			Name: "aws_efs_replication_configuration",
		},
		{
			Name: "aws_egress_only_internet_gateway",
		},
		{
			Name: "aws_eip",
		},
		{
			Name: "aws_eip_association",
		},
		{
			Name: "aws_eks_addon",
		},
		{
			Name: "aws_eks_cluster",
		},
		{
			Name: "aws_eks_fargate_profile",
		},
		{
			Name: "aws_eks_identity_provider_config",
		},
		{
			Name: "aws_eks_node_group",
		},
		{
			Name: "aws_elastic_beanstalk_application",
		},
		{
			Name: "aws_elastic_beanstalk_application_version",
		},
		{
			Name: "aws_elastic_beanstalk_configuration_template",
		},
		{
			Name: "aws_elastic_beanstalk_environment",
		},
		{
			Name: "aws_elasticache_cluster",
		},
		{
			Name: "aws_elasticache_global_replication_group",
		},
		{
			Name: "aws_elasticache_parameter_group",
		},
		{
			Name: "aws_elasticache_replication_group",
		},
		{
			Name: "aws_elasticache_subnet_group",
		},
		{
			Name: "aws_elasticache_user",
		},
		{
			Name: "aws_elasticache_user_group",
		},
		{
			Name: "aws_elasticache_user_group_association",
		},
		{
			Name: "aws_elasticsearch_domain",
		},
		{
			Name: "aws_elasticsearch_domain_policy",
		},
		{
			Name: "aws_elasticsearch_domain_saml_options",
		},
		{
			Name: "aws_elastictranscoder_pipeline",
		},
		{
			Name: "aws_elastictranscoder_preset",
		},
		{
			Name: "aws_elb",
		},
		{
			Name: "aws_elb_attachment",
		},
		{
			Name: "aws_emr_block_public_access_configuration",
		},
		{
			Name: "aws_emr_cluster",
		},
		{
			Name: "aws_emr_instance_fleet",
		},
		{
			Name: "aws_emr_instance_group",
		},
		{
			Name: "aws_emr_managed_scaling_policy",
		},
		{
			Name: "aws_emr_security_configuration",
		},
		{
			Name: "aws_emr_studio",
		},
		{
			Name: "aws_emr_studio_session_mapping",
		},
		{
			Name: "aws_emrcontainers_job_template",
		},
		{
			Name: "aws_emrcontainers_virtual_cluster",
		},
		{
			Name: "aws_emrserverless_application",
		},
		{
			Name: "aws_evidently_feature",
		},
		{
			Name: "aws_evidently_launch",
		},
		{
			Name: "aws_evidently_project",
		},
		{
			Name: "aws_evidently_segment",
		},
		{
			Name: "aws_finspace_kx_cluster",
		},
		{
			Name: "aws_finspace_kx_database",
		},
		{
			Name: "aws_finspace_kx_environment",
		},
		{
			Name: "aws_finspace_kx_user",
		},
		{
			Name: "aws_fis_experiment_template",
		},
		{
			Name: "aws_flow_log",
		},
		{
			Name: "aws_fms_admin_account",
		},
		{
			Name: "aws_fms_policy",
		},
		{
			Name: "aws_fsx_backup",
		},
		{
			Name: "aws_fsx_data_repository_association",
		},
		{
			Name: "aws_fsx_file_cache",
		},
		{
			Name: "aws_fsx_lustre_file_system",
		},
		{
			Name: "aws_fsx_ontap_file_system",
		},
		{
			Name: "aws_fsx_ontap_storage_virtual_machine",
		},
		{
			Name: "aws_fsx_ontap_volume",
		},
		{
			Name: "aws_fsx_openzfs_file_system",
		},
		{
			Name: "aws_fsx_openzfs_snapshot",
		},
		{
			Name: "aws_fsx_openzfs_volume",
		},
		{
			Name: "aws_fsx_windows_file_system",
		},
		{
			Name: "aws_gamelift_alias",
		},
		{
			Name: "aws_gamelift_build",
		},
		{
			Name: "aws_gamelift_fleet",
		},
		{
			Name: "aws_gamelift_game_server_group",
		},
		{
			Name: "aws_gamelift_game_session_queue",
		},
		{
			Name: "aws_gamelift_script",
		},
		{
			Name: "aws_glacier_vault",
		},
		{
			Name: "aws_glacier_vault_lock",
		},
		{
			Name: "aws_globalaccelerator_accelerator",
		},
		{
			Name: "aws_globalaccelerator_custom_routing_accelerator",
		},
		{
			Name: "aws_globalaccelerator_custom_routing_endpoint_group",
		},
		{
			Name: "aws_globalaccelerator_custom_routing_listener",
		},
		{
			Name: "aws_globalaccelerator_endpoint_group",
		},
		{
			Name: "aws_globalaccelerator_listener",
		},
		{
			Name: "aws_glue_catalog_database",
		},
		{
			Name: "aws_glue_catalog_table",
		},
		{
			Name: "aws_glue_classifier",
		},
		{
			Name: "aws_glue_connection",
		},
		{
			Name: "aws_glue_crawler",
		},
		{
			Name: "aws_glue_data_catalog_encryption_settings",
		},
		{
			Name: "aws_glue_data_quality_ruleset",
		},
		{
			Name: "aws_glue_dev_endpoint",
		},
		{
			Name: "aws_glue_job",
		},
		{
			Name: "aws_glue_ml_transform",
		},
		{
			Name: "aws_glue_partition",
		},
		{
			Name: "aws_glue_partition_index",
		},
		{
			Name: "aws_glue_registry",
		},
		{
			Name: "aws_glue_resource_policy",
		},
		{
			Name: "aws_glue_schema",
		},
		{
			Name: "aws_glue_security_configuration",
		},
		{
			Name: "aws_glue_trigger",
		},
		{
			Name: "aws_glue_user_defined_function",
		},
		{
			Name: "aws_glue_workflow",
		},
		{
			Name: "aws_grafana_license_association",
		},
		{
			Name: "aws_grafana_role_association",
		},
		{
			Name: "aws_grafana_workspace",
		},
		{
			Name: "aws_grafana_workspace_api_key",
		},
		{
			Name: "aws_grafana_workspace_saml_configuration",
		},
		{
			Name: "aws_guardduty_detector",
		},
		{
			Name: "aws_guardduty_filter",
		},
		{
			Name: "aws_guardduty_invite_accepter",
		},
		{
			Name: "aws_guardduty_ipset",
		},
		{
			Name: "aws_guardduty_member",
		},
		{
			Name: "aws_guardduty_organization_admin_account",
		},
		{
			Name: "aws_guardduty_organization_configuration",
		},
		{
			Name: "aws_guardduty_publishing_destination",
		},
		{
			Name: "aws_guardduty_threatintelset",
		},
		{
			Name: "aws_iam_access_key",
		},
		{
			Name: "aws_iam_account_alias",
		},
		{
			Name: "aws_iam_account_password_policy",
		},
		{
			Name: "aws_iam_group",
		},
		{
			Name: "aws_iam_group_membership",
		},
		{
			Name: "aws_iam_group_policy",
		},
		{
			Name: "aws_iam_group_policy_attachment",
		},
		{
			Name: "aws_iam_instance_profile",
		},
		{
			Name: "aws_iam_openid_connect_provider",
		},
		{
			Name: "aws_iam_policy",
		},
		{
			Name: "aws_iam_policy_attachment",
		},
		{
			Name:    "aws_iam_role",
			IconURL: "https://raw.githubusercontent.com/tf2d2/icons/main/aws/resource/Security-Identity-Compliance/AWS-Identity-Access-Management_Role.svg",
		},
		{
			Name: "aws_iam_role_policy",
		},
		{
			Name: "aws_iam_role_policy_attachment",
		},
		{
			Name: "aws_iam_saml_provider",
		},
		{
			Name: "aws_iam_security_token_service_preferences",
		},
		{
			Name: "aws_iam_server_certificate",
		},
		{
			Name: "aws_iam_service_linked_role",
		},
		{
			Name: "aws_iam_service_specific_credential",
		},
		{
			Name: "aws_iam_signing_certificate",
		},
		{
			Name: "aws_iam_user",
		},
		{
			Name: "aws_iam_user_group_membership",
		},
		{
			Name: "aws_iam_user_login_profile",
		},
		{
			Name: "aws_iam_user_policy",
		},
		{
			Name: "aws_iam_user_policy_attachment",
		},
		{
			Name: "aws_iam_user_ssh_key",
		},
		{
			Name: "aws_iam_virtual_mfa_device",
		},
		{
			Name: "aws_identitystore_group",
		},
		{
			Name: "aws_identitystore_group_membership",
		},
		{
			Name: "aws_identitystore_user",
		},
		{
			Name: "aws_imagebuilder_component",
		},
		{
			Name: "aws_imagebuilder_container_recipe",
		},
		{
			Name: "aws_imagebuilder_distribution_configuration",
		},
		{
			Name: "aws_imagebuilder_image",
		},
		{
			Name: "aws_imagebuilder_image_pipeline",
		},
		{
			Name: "aws_imagebuilder_image_recipe",
		},
		{
			Name: "aws_imagebuilder_infrastructure_configuration",
		},
		{
			Name: "aws_inspector2_delegated_admin_account",
		},
		{
			Name: "aws_inspector2_enabler",
		},
		{
			Name: "aws_inspector2_member_association",
		},
		{
			Name: "aws_inspector2_organization_configuration",
		},
		{
			Name: "aws_inspector_assessment_target",
		},
		{
			Name: "aws_inspector_assessment_template",
		},
		{
			Name: "aws_inspector_resource_group",
		},
		{
			Name: "aws_instance",
		},
		{
			Name: "aws_internet_gateway",
		},
		{
			Name: "aws_internet_gateway_attachment",
		},
		{
			Name: "aws_internetmonitor_monitor",
		},
		{
			Name: "aws_iot_authorizer",
		},
		{
			Name: "aws_iot_certificate",
		},
		{
			Name: "aws_iot_indexing_configuration",
		},
		{
			Name: "aws_iot_logging_options",
		},
		{
			Name: "aws_iot_policy",
		},
		{
			Name: "aws_iot_policy_attachment",
		},
		{
			Name: "aws_iot_provisioning_template",
		},
		{
			Name: "aws_iot_role_alias",
		},
		{
			Name: "aws_iot_thing",
		},
		{
			Name: "aws_iot_thing_group",
		},
		{
			Name: "aws_iot_thing_group_membership",
		},
		{
			Name: "aws_iot_thing_principal_attachment",
		},
		{
			Name: "aws_iot_thing_type",
		},
		{
			Name: "aws_iot_topic_rule",
		},
		{
			Name: "aws_iot_topic_rule_destination",
		},
		{
			Name: "aws_ivs_channel",
		},
		{
			Name: "aws_ivs_playback_key_pair",
		},
		{
			Name: "aws_ivs_recording_configuration",
		},
		{
			Name: "aws_ivschat_logging_configuration",
		},
		{
			Name: "aws_ivschat_room",
		},
		{
			Name: "aws_kendra_data_source",
		},
		{
			Name: "aws_kendra_experience",
		},
		{
			Name: "aws_kendra_faq",
		},
		{
			Name: "aws_kendra_index",
		},
		{
			Name: "aws_kendra_query_suggestions_block_list",
		},
		{
			Name: "aws_kendra_thesaurus",
		},
		{
			Name: "aws_key_pair",
		},
		{
			Name: "aws_keyspaces_keyspace",
		},
		{
			Name: "aws_keyspaces_table",
		},
		{
			Name: "aws_kinesis_analytics_application",
		},
		{
			Name: "aws_kinesis_firehose_delivery_stream",
		},
		{
			Name: "aws_kinesis_stream",
		},
		{
			Name: "aws_kinesis_stream_consumer",
		},
		{
			Name: "aws_kinesis_video_stream",
		},
		{
			Name: "aws_kinesisanalyticsv2_application",
		},
		{
			Name: "aws_kinesisanalyticsv2_application_snapshot",
		},
		{
			Name: "aws_kms_alias",
		},
		{
			Name: "aws_kms_ciphertext",
		},
		{
			Name: "aws_kms_custom_key_store",
		},
		{
			Name: "aws_kms_external_key",
		},
		{
			Name: "aws_kms_grant",
		},
		{
			Name: "aws_kms_key",
		},
		{
			Name: "aws_kms_key_policy",
		},
		{
			Name: "aws_kms_replica_external_key",
		},
		{
			Name: "aws_kms_replica_key",
		},
		{
			Name: "aws_lakeformation_data_lake_settings",
		},
		{
			Name: "aws_lakeformation_lf_tag",
		},
		{
			Name: "aws_lakeformation_permissions",
		},
		{
			Name: "aws_lakeformation_resource",
		},
		{
			Name: "aws_lakeformation_resource_lf_tags",
		},
		{
			Name: "aws_lambda_alias",
		},
		{
			Name: "aws_lambda_code_signing_config",
		},
		{
			Name: "aws_lambda_event_source_mapping",
		},
		{
			Name: "aws_lambda_function",
		},
		{
			Name: "aws_lambda_function_event_invoke_config",
		},
		{
			Name: "aws_lambda_function_url",
		},
		{
			Name: "aws_lambda_invocation",
		},
		{
			Name: "aws_lambda_layer_version",
		},
		{
			Name: "aws_lambda_layer_version_permission",
		},
		{
			Name: "aws_lambda_permission",
		},
		{
			Name: "aws_lambda_provisioned_concurrency_config",
		},
		{
			Name: "aws_launch_configuration",
		},
		{
			Name: "aws_launch_template",
		},
		{
			Name: "aws_lb",
		},
		{
			Name: "aws_lb_cookie_stickiness_policy",
		},
		{
			Name: "aws_lb_listener",
		},
		{
			Name: "aws_lb_listener_certificate",
		},
		{
			Name: "aws_lb_listener_rule",
		},
		{
			Name: "aws_lb_ssl_negotiation_policy",
		},
		{
			Name: "aws_lb_target_group",
		},
		{
			Name: "aws_lb_target_group_attachment",
		},
		{
			Name: "aws_lex_bot",
		},
		{
			Name: "aws_lex_bot_alias",
		},
		{
			Name: "aws_lex_intent",
		},
		{
			Name: "aws_lex_slot_type",
		},
		{
			Name: "aws_licensemanager_association",
		},
		{
			Name: "aws_licensemanager_grant",
		},
		{
			Name: "aws_licensemanager_grant_accepter",
		},
		{
			Name: "aws_licensemanager_license_configuration",
		},
		{
			Name: "aws_lightsail_bucket",
		},
		{
			Name: "aws_lightsail_bucket_access_key",
		},
		{
			Name: "aws_lightsail_bucket_resource_access",
		},
		{
			Name: "aws_lightsail_certificate",
		},
		{
			Name: "aws_lightsail_container_service",
		},
		{
			Name: "aws_lightsail_container_service_deployment_version",
		},
		{
			Name: "aws_lightsail_database",
		},
		{
			Name: "aws_lightsail_disk",
		},
		{
			Name: "aws_lightsail_disk_attachment",
		},
		{
			Name: "aws_lightsail_distribution",
		},
		{
			Name: "aws_lightsail_domain",
		},
		{
			Name: "aws_lightsail_domain_entry",
		},
		{
			Name: "aws_lightsail_instance",
		},
		{
			Name: "aws_lightsail_instance_public_ports",
		},
		{
			Name: "aws_lightsail_key_pair",
		},
		{
			Name: "aws_lightsail_lb",
		},
		{
			Name: "aws_lightsail_lb_attachment",
		},
		{
			Name: "aws_lightsail_lb_certificate",
		},
		{
			Name: "aws_lightsail_lb_certificate_attachment",
		},
		{
			Name: "aws_lightsail_lb_https_redirection_policy",
		},
		{
			Name: "aws_lightsail_lb_stickiness_policy",
		},
		{
			Name: "aws_lightsail_static_ip",
		},
		{
			Name: "aws_lightsail_static_ip_attachment",
		},
		{
			Name: "aws_load_balancer_backend_server_policy",
		},
		{
			Name: "aws_load_balancer_listener_policy",
		},
		{
			Name: "aws_load_balancer_policy",
		},
		{
			Name: "aws_location_geofence_collection",
		},
		{
			Name: "aws_location_map",
		},
		{
			Name: "aws_location_place_index",
		},
		{
			Name: "aws_location_route_calculator",
		},
		{
			Name: "aws_location_tracker",
		},
		{
			Name: "aws_location_tracker_association",
		},
		{
			Name: "aws_macie2_account",
		},
		{
			Name: "aws_macie2_classification_export_configuration",
		},
		{
			Name: "aws_macie2_classification_job",
		},
		{
			Name: "aws_macie2_custom_data_identifier",
		},
		{
			Name: "aws_macie2_findings_filter",
		},
		{
			Name: "aws_macie2_invitation_accepter",
		},
		{
			Name: "aws_macie2_member",
		},
		{
			Name: "aws_macie2_organization_admin_account",
		},
		{
			Name: "aws_main_route_table_association",
		},
		{
			Name: "aws_media_convert_queue",
		},
		{
			Name: "aws_media_package_channel",
		},
		{
			Name: "aws_media_store_container",
		},
		{
			Name: "aws_media_store_container_policy",
		},
		{
			Name: "aws_medialive_channel",
		},
		{
			Name: "aws_medialive_input",
		},
		{
			Name: "aws_medialive_input_security_group",
		},
		{
			Name: "aws_medialive_multiplex",
		},
		{
			Name: "aws_medialive_multiplex_program",
		},
		{
			Name: "aws_memorydb_acl",
		},
		{
			Name: "aws_memorydb_cluster",
		},
		{
			Name: "aws_memorydb_parameter_group",
		},
		{
			Name: "aws_memorydb_snapshot",
		},
		{
			Name: "aws_memorydb_subnet_group",
		},
		{
			Name: "aws_memorydb_user",
		},
		{
			Name: "aws_mq_broker",
		},
		{
			Name: "aws_mq_configuration",
		},
		{
			Name: "aws_msk_cluster",
		},
		{
			Name: "aws_msk_cluster_policy",
		},
		{
			Name: "aws_msk_configuration",
		},
		{
			Name: "aws_msk_scram_secret_association",
		},
		{
			Name: "aws_msk_serverless_cluster",
		},
		{
			Name: "aws_msk_vpc_connection",
		},
		{
			Name: "aws_mskconnect_connector",
		},
		{
			Name: "aws_mskconnect_custom_plugin",
		},
		{
			Name: "aws_mskconnect_worker_configuration",
		},
		{
			Name: "aws_mwaa_environment",
		},
		{
			Name: "aws_nat_gateway",
		},
		{
			Name: "aws_neptune_cluster",
		},
		{
			Name: "aws_neptune_cluster_endpoint",
		},
		{
			Name: "aws_neptune_cluster_instance",
		},
		{
			Name: "aws_neptune_cluster_parameter_group",
		},
		{
			Name: "aws_neptune_cluster_snapshot",
		},
		{
			Name: "aws_neptune_event_subscription",
		},
		{
			Name: "aws_neptune_global_cluster",
		},
		{
			Name: "aws_neptune_parameter_group",
		},
		{
			Name: "aws_neptune_subnet_group",
		},
		{
			Name: "aws_network_acl",
		},
		{
			Name: "aws_network_acl_association",
		},
		{
			Name: "aws_network_acl_rule",
		},
		{
			Name: "aws_network_interface",
		},
		{
			Name: "aws_network_interface_attachment",
		},
		{
			Name: "aws_network_interface_sg_attachment",
		},
		{
			Name: "aws_networkfirewall_firewall",
		},
		{
			Name: "aws_networkfirewall_firewall_policy",
		},
		{
			Name: "aws_networkfirewall_logging_configuration",
		},
		{
			Name: "aws_networkfirewall_resource_policy",
		},
		{
			Name: "aws_networkfirewall_rule_group",
		},
		{
			Name: "aws_networkmanager_attachment_accepter",
		},
		{
			Name: "aws_networkmanager_connect_attachment",
		},
		{
			Name: "aws_networkmanager_connect_peer",
		},
		{
			Name: "aws_networkmanager_connection",
		},
		{
			Name: "aws_networkmanager_core_network",
		},
		{
			Name: "aws_networkmanager_core_network_policy_attachment",
		},
		{
			Name: "aws_networkmanager_customer_gateway_association",
		},
		{
			Name: "aws_networkmanager_device",
		},
		{
			Name: "aws_networkmanager_global_network",
		},
		{
			Name: "aws_networkmanager_link",
		},
		{
			Name: "aws_networkmanager_link_association",
		},
		{
			Name: "aws_networkmanager_site",
		},
		{
			Name: "aws_networkmanager_site_to_site_vpn_attachment",
		},
		{
			Name: "aws_networkmanager_transit_gateway_connect_peer_association",
		},
		{
			Name: "aws_networkmanager_transit_gateway_peering",
		},
		{
			Name: "aws_networkmanager_transit_gateway_registration",
		},
		{
			Name: "aws_networkmanager_transit_gateway_route_table_attachment",
		},
		{
			Name: "aws_networkmanager_vpc_attachment",
		},
		{
			Name: "aws_oam_link",
		},
		{
			Name: "aws_oam_sink",
		},
		{
			Name: "aws_oam_sink_policy",
		},
		{
			Name: "aws_opensearch_domain",
		},
		{
			Name: "aws_opensearch_domain_policy",
		},
		{
			Name: "aws_opensearch_domain_saml_options",
		},
		{
			Name: "aws_opensearch_inbound_connection_accepter",
		},
		{
			Name: "aws_opensearch_outbound_connection",
		},
		{
			Name: "aws_opensearch_vpc_endpoint",
		},
		{
			Name: "aws_opensearchserverless_access_policy",
		},
		{
			Name: "aws_opensearchserverless_collection",
		},
		{
			Name: "aws_opensearchserverless_security_config",
		},
		{
			Name: "aws_opensearchserverless_security_policy",
		},
		{
			Name: "aws_opensearchserverless_vpc_endpoint",
		},
		{
			Name: "aws_opsworks_application",
		},
		{
			Name: "aws_opsworks_custom_layer",
		},
		{
			Name: "aws_opsworks_ecs_cluster_layer",
		},
		{
			Name: "aws_opsworks_ganglia_layer",
		},
		{
			Name: "aws_opsworks_haproxy_layer",
		},
		{
			Name: "aws_opsworks_instance",
		},
		{
			Name: "aws_opsworks_java_app_layer",
		},
		{
			Name: "aws_opsworks_memcached_layer",
		},
		{
			Name: "aws_opsworks_mysql_layer",
		},
		{
			Name: "aws_opsworks_nodejs_app_layer",
		},
		{
			Name: "aws_opsworks_permission",
		},
		{
			Name: "aws_opsworks_php_app_layer",
		},
		{
			Name: "aws_opsworks_rails_app_layer",
		},
		{
			Name: "aws_opsworks_rds_db_instance",
		},
		{
			Name: "aws_opsworks_stack",
		},
		{
			Name: "aws_opsworks_static_web_layer",
		},
		{
			Name: "aws_opsworks_user_profile",
		},
		{
			Name: "aws_organizations_account",
		},
		{
			Name: "aws_organizations_delegated_administrator",
		},
		{
			Name: "aws_organizations_organization",
		},
		{
			Name: "aws_organizations_organizational_unit",
		},
		{
			Name: "aws_organizations_policy",
		},
		{
			Name: "aws_organizations_policy_attachment",
		},
		{
			Name: "aws_organizations_resource_policy",
		},
		{
			Name: "aws_pinpoint_adm_channel",
		},
		{
			Name: "aws_pinpoint_apns_channel",
		},
		{
			Name: "aws_pinpoint_apns_sandbox_channel",
		},
		{
			Name: "aws_pinpoint_apns_voip_channel",
		},
		{
			Name: "aws_pinpoint_apns_voip_sandbox_channel",
		},
		{
			Name: "aws_pinpoint_app",
		},
		{
			Name: "aws_pinpoint_baidu_channel",
		},
		{
			Name: "aws_pinpoint_email_channel",
		},
		{
			Name: "aws_pinpoint_event_stream",
		},
		{
			Name: "aws_pinpoint_gcm_channel",
		},
		{
			Name: "aws_pinpoint_sms_channel",
		},
		{
			Name: "aws_pipes_pipe",
		},
		{
			Name: "aws_placement_group",
		},
		{
			Name: "aws_prometheus_alert_manager_definition",
		},
		{
			Name: "aws_prometheus_rule_group_namespace",
		},
		{
			Name: "aws_prometheus_workspace",
		},
		{
			Name: "aws_proxy_protocol_policy",
		},
		{
			Name: "aws_qldb_ledger",
		},
		{
			Name: "aws_qldb_stream",
		},
		{
			Name: "aws_quicksight_account_subscription",
		},
		{
			Name: "aws_quicksight_analysis",
		},
		{
			Name: "aws_quicksight_dashboard",
		},
		{
			Name: "aws_quicksight_data_set",
		},
		{
			Name: "aws_quicksight_data_source",
		},
		{
			Name: "aws_quicksight_folder",
		},
		{
			Name: "aws_quicksight_folder_membership",
		},
		{
			Name: "aws_quicksight_group",
		},
		{
			Name: "aws_quicksight_group_membership",
		},
		{
			Name: "aws_quicksight_iam_policy_assignment",
		},
		{
			Name: "aws_quicksight_ingestion",
		},
		{
			Name: "aws_quicksight_namespace",
		},
		{
			Name: "aws_quicksight_refresh_schedule",
		},
		{
			Name: "aws_quicksight_template",
		},
		{
			Name: "aws_quicksight_template_alias",
		},
		{
			Name: "aws_quicksight_theme",
		},
		{
			Name: "aws_quicksight_user",
		},
		{
			Name: "aws_quicksight_vpc_connection",
		},
		{
			Name: "aws_ram_principal_association",
		},
		{
			Name: "aws_ram_resource_association",
		},
		{
			Name: "aws_ram_resource_share",
		},
		{
			Name: "aws_ram_resource_share_accepter",
		},
		{
			Name: "aws_ram_sharing_with_organization",
		},
		{
			Name: "aws_rbin_rule",
		},
		{
			Name: "aws_rds_cluster",
		},
		{
			Name: "aws_rds_cluster_activity_stream",
		},
		{
			Name: "aws_rds_cluster_endpoint",
		},
		{
			Name: "aws_rds_cluster_instance",
		},
		{
			Name: "aws_rds_cluster_parameter_group",
		},
		{
			Name: "aws_rds_cluster_role_association",
		},
		{
			Name: "aws_rds_export_task",
		},
		{
			Name: "aws_rds_global_cluster",
		},
		{
			Name: "aws_rds_reserved_instance",
		},
		{
			Name: "aws_redshift_authentication_profile",
		},
		{
			Name: "aws_redshift_cluster",
		},
		{
			Name: "aws_redshift_cluster_iam_roles",
		},
		{
			Name: "aws_redshift_cluster_snapshot",
		},
		{
			Name: "aws_redshift_endpoint_access",
		},
		{
			Name: "aws_redshift_endpoint_authorization",
		},
		{
			Name: "aws_redshift_event_subscription",
		},
		{
			Name: "aws_redshift_hsm_client_certificate",
		},
		{
			Name: "aws_redshift_hsm_configuration",
		},
		{
			Name: "aws_redshift_parameter_group",
		},
		{
			Name: "aws_redshift_partner",
		},
		{
			Name: "aws_redshift_scheduled_action",
		},
		{
			Name: "aws_redshift_snapshot_copy_grant",
		},
		{
			Name: "aws_redshift_snapshot_schedule",
		},
		{
			Name: "aws_redshift_snapshot_schedule_association",
		},
		{
			Name: "aws_redshift_subnet_group",
		},
		{
			Name: "aws_redshift_usage_limit",
		},
		{
			Name: "aws_redshiftdata_statement",
		},
		{
			Name: "aws_redshiftserverless_endpoint_access",
		},
		{
			Name: "aws_redshiftserverless_namespace",
		},
		{
			Name: "aws_redshiftserverless_resource_policy",
		},
		{
			Name: "aws_redshiftserverless_snapshot",
		},
		{
			Name: "aws_redshiftserverless_usage_limit",
		},
		{
			Name: "aws_redshiftserverless_workgroup",
		},
		{
			Name: "aws_resourceexplorer2_index",
		},
		{
			Name: "aws_resourceexplorer2_view",
		},
		{
			Name: "aws_resourcegroups_group",
		},
		{
			Name: "aws_resourcegroups_resource",
		},
		{
			Name: "aws_rolesanywhere_profile",
		},
		{
			Name: "aws_rolesanywhere_trust_anchor",
		},
		{
			Name: "aws_route",
		},
		{
			Name: "aws_route53_cidr_collection",
		},
		{
			Name: "aws_route53_cidr_location",
		},
		{
			Name: "aws_route53_delegation_set",
		},
		{
			Name: "aws_route53_health_check",
		},
		{
			Name: "aws_route53_hosted_zone_dnssec",
		},
		{
			Name: "aws_route53_key_signing_key",
		},
		{
			Name: "aws_route53_query_log",
		},
		{
			Name: "aws_route53_record",
		},
		{
			Name: "aws_route53_resolver_config",
		},
		{
			Name: "aws_route53_resolver_dnssec_config",
		},
		{
			Name: "aws_route53_resolver_endpoint",
		},
		{
			Name: "aws_route53_resolver_firewall_config",
		},
		{
			Name: "aws_route53_resolver_firewall_domain_list",
		},
		{
			Name: "aws_route53_resolver_firewall_rule",
		},
		{
			Name: "aws_route53_resolver_firewall_rule_group",
		},
		{
			Name: "aws_route53_resolver_firewall_rule_group_association",
		},
		{
			Name: "aws_route53_resolver_query_log_config",
		},
		{
			Name: "aws_route53_resolver_query_log_config_association",
		},
		{
			Name: "aws_route53_resolver_rule",
		},
		{
			Name: "aws_route53_resolver_rule_association",
		},
		{
			Name: "aws_route53_traffic_policy",
		},
		{
			Name: "aws_route53_traffic_policy_instance",
		},
		{
			Name: "aws_route53_vpc_association_authorization",
		},
		{
			Name: "aws_route53_zone",
		},
		{
			Name: "aws_route53_zone_association",
		},
		{
			Name: "aws_route53domains_registered_domain",
		},
		{
			Name: "aws_route53recoverycontrolconfig_cluster",
		},
		{
			Name: "aws_route53recoverycontrolconfig_control_panel",
		},
		{
			Name: "aws_route53recoverycontrolconfig_routing_control",
		},
		{
			Name: "aws_route53recoverycontrolconfig_safety_rule",
		},
		{
			Name: "aws_route53recoveryreadiness_cell",
		},
		{
			Name: "aws_route53recoveryreadiness_readiness_check",
		},
		{
			Name: "aws_route53recoveryreadiness_recovery_group",
		},
		{
			Name: "aws_route53recoveryreadiness_resource_set",
		},
		{
			Name: "aws_route_table",
		},
		{
			Name: "aws_route_table_association",
		},
		{
			Name: "aws_rum_app_monitor",
		},
		{
			Name: "aws_rum_metrics_destination",
		},
		{
			Name: "aws_s3_access_point",
		},
		{
			Name: "aws_s3_account_public_access_block",
		},
		{
			Name: "aws_s3_bucket",
		},
		{
			Name: "aws_s3_bucket_accelerate_configuration",
		},
		{
			Name: "aws_s3_bucket_acl",
		},
		{
			Name: "aws_s3_bucket_analytics_configuration",
		},
		{
			Name: "aws_s3_bucket_cors_configuration",
		},
		{
			Name: "aws_s3_bucket_intelligent_tiering_configuration",
		},
		{
			Name: "aws_s3_bucket_inventory",
		},
		{
			Name: "aws_s3_bucket_lifecycle_configuration",
		},
		{
			Name: "aws_s3_bucket_logging",
		},
		{
			Name: "aws_s3_bucket_metric",
		},
		{
			Name: "aws_s3_bucket_notification",
		},
		{
			Name: "aws_s3_bucket_object",
		},
		{
			Name: "aws_s3_bucket_object_lock_configuration",
		},
		{
			Name: "aws_s3_bucket_ownership_controls",
		},
		{
			Name: "aws_s3_bucket_policy",
		},
		{
			Name: "aws_s3_bucket_public_access_block",
		},
		{
			Name: "aws_s3_bucket_replication_configuration",
		},
		{
			Name: "aws_s3_bucket_request_payment_configuration",
		},
		{
			Name: "aws_s3_bucket_server_side_encryption_configuration",
		},
		{
			Name: "aws_s3_bucket_versioning",
		},
		{
			Name: "aws_s3_bucket_website_configuration",
		},
		{
			Name: "aws_s3_object",
		},
		{
			Name: "aws_s3_object_copy",
		},
		{
			Name: "aws_s3control_access_point_policy",
		},
		{
			Name: "aws_s3control_bucket",
		},
		{
			Name: "aws_s3control_bucket_lifecycle_configuration",
		},
		{
			Name: "aws_s3control_bucket_policy",
		},
		{
			Name: "aws_s3control_multi_region_access_point",
		},
		{
			Name: "aws_s3control_multi_region_access_point_policy",
		},
		{
			Name: "aws_s3control_object_lambda_access_point",
		},
		{
			Name: "aws_s3control_object_lambda_access_point_policy",
		},
		{
			Name: "aws_s3control_storage_lens_configuration",
		},
		{
			Name: "aws_s3outposts_endpoint",
		},
		{
			Name: "aws_sagemaker_app",
		},
		{
			Name: "aws_sagemaker_app_image_config",
		},
		{
			Name: "aws_sagemaker_code_repository",
		},
		{
			Name: "aws_sagemaker_data_quality_job_definition",
		},
		{
			Name: "aws_sagemaker_device",
		},
		{
			Name: "aws_sagemaker_device_fleet",
		},
		{
			Name: "aws_sagemaker_domain",
		},
		{
			Name: "aws_sagemaker_endpoint",
		},
		{
			Name: "aws_sagemaker_endpoint_configuration",
		},
		{
			Name: "aws_sagemaker_feature_group",
		},
		{
			Name: "aws_sagemaker_flow_definition",
		},
		{
			Name: "aws_sagemaker_human_task_ui",
		},
		{
			Name: "aws_sagemaker_image",
		},
		{
			Name: "aws_sagemaker_image_version",
		},
		{
			Name: "aws_sagemaker_model",
		},
		{
			Name: "aws_sagemaker_model_package_group",
		},
		{
			Name: "aws_sagemaker_model_package_group_policy",
		},
		{
			Name: "aws_sagemaker_monitoring_schedule",
		},
		{
			Name: "aws_sagemaker_notebook_instance",
		},
		{
			Name: "aws_sagemaker_notebook_instance_lifecycle_configuration",
		},
		{
			Name: "aws_sagemaker_pipeline",
		},
		{
			Name: "aws_sagemaker_project",
		},
		{
			Name: "aws_sagemaker_servicecatalog_portfolio_status",
		},
		{
			Name: "aws_sagemaker_space",
		},
		{
			Name: "aws_sagemaker_studio_lifecycle_config",
		},
		{
			Name: "aws_sagemaker_user_profile",
		},
		{
			Name: "aws_sagemaker_workforce",
		},
		{
			Name: "aws_sagemaker_workteam",
		},
		{
			Name: "aws_scheduler_schedule",
		},
		{
			Name: "aws_scheduler_schedule_group",
		},
		{
			Name: "aws_schemas_discoverer",
		},
		{
			Name: "aws_schemas_registry",
		},
		{
			Name: "aws_schemas_registry_policy",
		},
		{
			Name: "aws_schemas_schema",
		},
		{
			Name: "aws_secretsmanager_secret",
		},
		{
			Name: "aws_secretsmanager_secret_policy",
		},
		{
			Name: "aws_secretsmanager_secret_rotation",
		},
		{
			Name: "aws_secretsmanager_secret_version",
		},
		{
			Name: "aws_security_group",
		},
		{
			Name: "aws_security_group_rule",
		},
		{
			Name: "aws_securityhub_account",
		},
		{
			Name: "aws_securityhub_action_target",
		},
		{
			Name: "aws_securityhub_finding_aggregator",
		},
		{
			Name: "aws_securityhub_insight",
		},
		{
			Name: "aws_securityhub_invite_accepter",
		},
		{
			Name: "aws_securityhub_member",
		},
		{
			Name: "aws_securityhub_organization_admin_account",
		},
		{
			Name: "aws_securityhub_organization_configuration",
		},
		{
			Name: "aws_securityhub_product_subscription",
		},
		{
			Name: "aws_securityhub_standards_control",
		},
		{
			Name: "aws_securityhub_standards_subscription",
		},
		{
			Name: "aws_serverlessapplicationrepository_cloudformation_stack",
		},
		{
			Name: "aws_service_discovery_http_namespace",
		},
		{
			Name: "aws_service_discovery_instance",
		},
		{
			Name: "aws_service_discovery_private_dns_namespace",
		},
		{
			Name: "aws_service_discovery_public_dns_namespace",
		},
		{
			Name: "aws_service_discovery_service",
		},
		{
			Name: "aws_servicecatalog_budget_resource_association",
		},
		{
			Name: "aws_servicecatalog_constraint",
		},
		{
			Name: "aws_servicecatalog_organizations_access",
		},
		{
			Name: "aws_servicecatalog_portfolio",
		},
		{
			Name: "aws_servicecatalog_portfolio_share",
		},
		{
			Name: "aws_servicecatalog_principal_portfolio_association",
		},
		{
			Name: "aws_servicecatalog_product",
		},
		{
			Name: "aws_servicecatalog_product_portfolio_association",
		},
		{
			Name: "aws_servicecatalog_provisioned_product",
		},
		{
			Name: "aws_servicecatalog_provisioning_artifact",
		},
		{
			Name: "aws_servicecatalog_service_action",
		},
		{
			Name: "aws_servicecatalog_tag_option",
		},
		{
			Name: "aws_servicecatalog_tag_option_resource_association",
		},
		{
			Name: "aws_servicequotas_service_quota",
		},
		{
			Name: "aws_ses_active_receipt_rule_set",
		},
		{
			Name: "aws_ses_configuration_set",
		},
		{
			Name: "aws_ses_domain_dkim",
		},
		{
			Name: "aws_ses_domain_identity",
		},
		{
			Name: "aws_ses_domain_identity_verification",
		},
		{
			Name: "aws_ses_domain_mail_from",
		},
		{
			Name: "aws_ses_email_identity",
		},
		{
			Name: "aws_ses_event_destination",
		},
		{
			Name: "aws_ses_identity_notification_topic",
		},
		{
			Name: "aws_ses_identity_policy",
		},
		{
			Name: "aws_ses_receipt_filter",
		},
		{
			Name: "aws_ses_receipt_rule",
		},
		{
			Name: "aws_ses_receipt_rule_set",
		},
		{
			Name: "aws_ses_template",
		},
		{
			Name: "aws_sesv2_configuration_set",
		},
		{
			Name: "aws_sesv2_configuration_set_event_destination",
		},
		{
			Name: "aws_sesv2_contact_list",
		},
		{
			Name: "aws_sesv2_dedicated_ip_assignment",
		},
		{
			Name: "aws_sesv2_dedicated_ip_pool",
		},
		{
			Name: "aws_sesv2_email_identity",
		},
		{
			Name: "aws_sesv2_email_identity_feedback_attributes",
		},
		{
			Name: "aws_sesv2_email_identity_mail_from_attributes",
		},
		{
			Name: "aws_sfn_activity",
		},
		{
			Name: "aws_sfn_alias",
		},
		{
			Name: "aws_sfn_state_machine",
		},
		{
			Name: "aws_shield_application_layer_automatic_response",
		},
		{
			Name: "aws_shield_drt_access_log_bucket_association",
		},
		{
			Name: "aws_shield_drt_access_role_arn_association",
		},
		{
			Name: "aws_shield_protection",
		},
		{
			Name: "aws_shield_protection_group",
		},
		{
			Name: "aws_shield_protection_health_check_association",
		},
		{
			Name: "aws_signer_signing_job",
		},
		{
			Name: "aws_signer_signing_profile",
		},
		{
			Name: "aws_signer_signing_profile_permission",
		},
		{
			Name: "aws_simpledb_domain",
		},
		{
			Name: "aws_snapshot_create_volume_permission",
		},
		{
			Name: "aws_sns_platform_application",
		},
		{
			Name: "aws_sns_sms_preferences",
		},
		{
			Name: "aws_sns_topic",
		},
		{
			Name: "aws_sns_topic_data_protection_policy",
		},
		{
			Name: "aws_sns_topic_policy",
		},
		{
			Name: "aws_sns_topic_subscription",
		},
		{
			Name: "aws_spot_datafeed_subscription",
		},
		{
			Name: "aws_spot_fleet_request",
		},
		{
			Name: "aws_spot_instance_request",
		},
		{
			Name: "aws_sqs_queue",
		},
		{
			Name: "aws_sqs_queue_policy",
		},
		{
			Name: "aws_sqs_queue_redrive_allow_policy",
		},
		{
			Name: "aws_sqs_queue_redrive_policy",
		},
		{
			Name: "aws_ssm_activation",
		},
		{
			Name: "aws_ssm_association",
		},
		{
			Name: "aws_ssm_default_patch_baseline",
		},
		{
			Name: "aws_ssm_document",
		},
		{
			Name: "aws_ssm_maintenance_window",
		},
		{
			Name: "aws_ssm_maintenance_window_target",
		},
		{
			Name: "aws_ssm_maintenance_window_task",
		},
		{
			Name: "aws_ssm_parameter",
		},
		{
			Name: "aws_ssm_patch_baseline",
		},
		{
			Name: "aws_ssm_patch_group",
		},
		{
			Name: "aws_ssm_resource_data_sync",
		},
		{
			Name: "aws_ssm_service_setting",
		},
		{
			Name: "aws_ssmcontacts_contact",
		},
		{
			Name: "aws_ssmcontacts_contact_channel",
		},
		{
			Name: "aws_ssmcontacts_plan",
		},
		{
			Name: "aws_ssmincidents_replication_set",
		},
		{
			Name: "aws_ssmincidents_response_plan",
		},
		{
			Name: "aws_ssoadmin_account_assignment",
		},
		{
			Name: "aws_ssoadmin_customer_managed_policy_attachment",
		},
		{
			Name: "aws_ssoadmin_instance_access_control_attributes",
		},
		{
			Name: "aws_ssoadmin_managed_policy_attachment",
		},
		{
			Name: "aws_ssoadmin_permission_set",
		},
		{
			Name: "aws_ssoadmin_permission_set_inline_policy",
		},
		{
			Name: "aws_ssoadmin_permissions_boundary_attachment",
		},
		{
			Name: "aws_storagegateway_cache",
		},
		{
			Name: "aws_storagegateway_cached_iscsi_volume",
		},
		{
			Name: "aws_storagegateway_file_system_association",
		},
		{
			Name: "aws_storagegateway_gateway",
		},
		{
			Name: "aws_storagegateway_nfs_file_share",
		},
		{
			Name: "aws_storagegateway_smb_file_share",
		},
		{
			Name: "aws_storagegateway_stored_iscsi_volume",
		},
		{
			Name: "aws_storagegateway_tape_pool",
		},
		{
			Name: "aws_storagegateway_upload_buffer",
		},
		{
			Name: "aws_storagegateway_working_storage",
		},
		{
			Name: "aws_subnet",
		},
		{
			Name: "aws_swf_domain",
		},
		{
			Name: "aws_synthetics_canary",
		},
		{
			Name: "aws_synthetics_group",
		},
		{
			Name: "aws_synthetics_group_association",
		},
		{
			Name: "aws_timestreamwrite_database",
		},
		{
			Name: "aws_timestreamwrite_table",
		},
		{
			Name: "aws_transcribe_language_model",
		},
		{
			Name: "aws_transcribe_medical_vocabulary",
		},
		{
			Name: "aws_transcribe_vocabulary",
		},
		{
			Name: "aws_transcribe_vocabulary_filter",
		},
		{
			Name: "aws_transfer_access",
		},
		{
			Name: "aws_transfer_agreement",
		},
		{
			Name: "aws_transfer_certificate",
		},
		{
			Name: "aws_transfer_connector",
		},
		{
			Name: "aws_transfer_profile",
		},
		{
			Name: "aws_transfer_server",
		},
		{
			Name: "aws_transfer_ssh_key",
		},
		{
			Name: "aws_transfer_tag",
		},
		{
			Name: "aws_transfer_user",
		},
		{
			Name: "aws_transfer_workflow",
		},
		{
			Name: "aws_verifiedaccess_instance",
		},
		{
			Name: "aws_verifiedaccess_trust_provider",
		},
		{
			Name: "aws_volume_attachment",
		},
		{
			Name: "aws_vpc",
		},
		{
			Name: "aws_vpc_dhcp_options",
		},
		{
			Name: "aws_vpc_dhcp_options_association",
		},
		{
			Name: "aws_vpc_endpoint",
		},
		{
			Name: "aws_vpc_endpoint_connection_accepter",
		},
		{
			Name: "aws_vpc_endpoint_connection_notification",
		},
		{
			Name: "aws_vpc_endpoint_policy",
		},
		{
			Name: "aws_vpc_endpoint_route_table_association",
		},
		{
			Name: "aws_vpc_endpoint_security_group_association",
		},
		{
			Name: "aws_vpc_endpoint_service",
		},
		{
			Name: "aws_vpc_endpoint_service_allowed_principal",
		},
		{
			Name: "aws_vpc_endpoint_subnet_association",
		},
		{
			Name: "aws_vpc_ipam",
		},
		{
			Name: "aws_vpc_ipam_organization_admin_account",
		},
		{
			Name: "aws_vpc_ipam_pool",
		},
		{
			Name: "aws_vpc_ipam_pool_cidr",
		},
		{
			Name: "aws_vpc_ipam_pool_cidr_allocation",
		},
		{
			Name: "aws_vpc_ipam_preview_next_cidr",
		},
		{
			Name: "aws_vpc_ipam_resource_discovery",
		},
		{
			Name: "aws_vpc_ipam_resource_discovery_association",
		},
		{
			Name: "aws_vpc_ipam_scope",
		},
		{
			Name: "aws_vpc_ipv4_cidr_block_association",
		},
		{
			Name: "aws_vpc_ipv6_cidr_block_association",
		},
		{
			Name: "aws_vpc_network_performance_metric_subscription",
		},
		{
			Name: "aws_vpc_peering_connection",
		},
		{
			Name: "aws_vpc_peering_connection_accepter",
		},
		{
			Name: "aws_vpc_peering_connection_options",
		},
		{
			Name: "aws_vpc_security_group_egress_rule",
		},
		{
			Name: "aws_vpc_security_group_ingress_rule",
		},
		{
			Name: "aws_vpclattice_access_log_subscription",
		},
		{
			Name: "aws_vpclattice_auth_policy",
		},
		{
			Name: "aws_vpclattice_listener",
		},
		{
			Name: "aws_vpclattice_listener_rule",
		},
		{
			Name: "aws_vpclattice_resource_policy",
		},
		{
			Name: "aws_vpclattice_service",
		},
		{
			Name: "aws_vpclattice_service_network",
		},
		{
			Name: "aws_vpclattice_service_network_service_association",
		},
		{
			Name: "aws_vpclattice_service_network_vpc_association",
		},
		{
			Name: "aws_vpclattice_target_group",
		},
		{
			Name: "aws_vpclattice_target_group_attachment",
		},
		{
			Name: "aws_vpn_connection",
		},
		{
			Name: "aws_vpn_connection_route",
		},
		{
			Name: "aws_vpn_gateway",
		},
		{
			Name: "aws_vpn_gateway_attachment",
		},
		{
			Name: "aws_vpn_gateway_route_propagation",
		},
		{
			Name: "aws_waf_byte_match_set",
		},
		{
			Name: "aws_waf_geo_match_set",
		},
		{
			Name: "aws_waf_ipset",
		},
		{
			Name: "aws_waf_rate_based_rule",
		},
		{
			Name: "aws_waf_regex_match_set",
		},
		{
			Name: "aws_waf_regex_pattern_set",
		},
		{
			Name: "aws_waf_rule",
		},
		{
			Name: "aws_waf_rule_group",
		},
		{
			Name: "aws_waf_size_constraint_set",
		},
		{
			Name: "aws_waf_sql_injection_match_set",
		},
		{
			Name: "aws_waf_web_acl",
		},
		{
			Name: "aws_waf_xss_match_set",
		},
		{
			Name: "aws_wafregional_byte_match_set",
		},
		{
			Name: "aws_wafregional_geo_match_set",
		},
		{
			Name: "aws_wafregional_ipset",
		},
		{
			Name: "aws_wafregional_rate_based_rule",
		},
		{
			Name: "aws_wafregional_regex_match_set",
		},
		{
			Name: "aws_wafregional_regex_pattern_set",
		},
		{
			Name: "aws_wafregional_rule",
		},
		{
			Name: "aws_wafregional_rule_group",
		},
		{
			Name: "aws_wafregional_size_constraint_set",
		},
		{
			Name: "aws_wafregional_sql_injection_match_set",
		},
		{
			Name: "aws_wafregional_web_acl",
		},
		{
			Name: "aws_wafregional_web_acl_association",
		},
		{
			Name: "aws_wafregional_xss_match_set",
		},
		{
			Name: "aws_wafv2_ip_set",
		},
		{
			Name: "aws_wafv2_regex_pattern_set",
		},
		{
			Name: "aws_wafv2_rule_group",
		},
		{
			Name: "aws_wafv2_web_acl",
		},
		{
			Name: "aws_wafv2_web_acl_association",
		},
		{
			Name: "aws_wafv2_web_acl_logging_configuration",
		},
		{
			Name: "aws_worklink_fleet",
		},
		{
			Name: "aws_worklink_website_certificate_authority_association",
		},
		{
			Name: "aws_workspaces_connection_alias",
		},
		{
			Name: "aws_workspaces_directory",
		},
		{
			Name: "aws_workspaces_ip_group",
		},
		{
			Name: "aws_workspaces_workspace",
		},
		{
			Name: "aws_xray_encryption_config",
		},
		{
			Name: "aws_xray_group",
		},
		{
			Name: "aws_xray_sampling_rule",
		},
	}
	resourcesMap = map[string]int{
		"aws_accessanalyzer_analyzer":                                      0,
		"aws_accessanalyzer_archive_rule":                                  1,
		"aws_account_alternate_contact":                                    2,
		"aws_account_primary_contact":                                      3,
		"aws_acm_certificate":                                              4,
		"aws_acm_certificate_validation":                                   5,
		"aws_acmpca_certificate":                                           6,
		"aws_acmpca_certificate_authority":                                 7,
		"aws_acmpca_certificate_authority_certificate":                     8,
		"aws_acmpca_permission":                                            9,
		"aws_acmpca_policy":                                                10,
		"aws_alb":                                                          11,
		"aws_alb_listener":                                                 12,
		"aws_alb_listener_certificate":                                     13,
		"aws_alb_listener_rule":                                            14,
		"aws_alb_target_group":                                             15,
		"aws_alb_target_group_attachment":                                  16,
		"aws_ami":                                                          17,
		"aws_ami_copy":                                                     18,
		"aws_ami_from_instance":                                            19,
		"aws_ami_launch_permission":                                        20,
		"aws_amplify_app":                                                  21,
		"aws_amplify_backend_environment":                                  22,
		"aws_amplify_branch":                                               23,
		"aws_amplify_domain_association":                                   24,
		"aws_amplify_webhook":                                              25,
		"aws_api_gateway_account":                                          26,
		"aws_api_gateway_api_key":                                          27,
		"aws_api_gateway_authorizer":                                       28,
		"aws_api_gateway_base_path_mapping":                                29,
		"aws_api_gateway_client_certificate":                               30,
		"aws_api_gateway_deployment":                                       31,
		"aws_api_gateway_documentation_part":                               32,
		"aws_api_gateway_documentation_version":                            33,
		"aws_api_gateway_domain_name":                                      34,
		"aws_api_gateway_gateway_response":                                 35,
		"aws_api_gateway_integration":                                      36,
		"aws_api_gateway_integration_response":                             37,
		"aws_api_gateway_method":                                           38,
		"aws_api_gateway_method_response":                                  39,
		"aws_api_gateway_method_settings":                                  40,
		"aws_api_gateway_model":                                            41,
		"aws_api_gateway_request_validator":                                42,
		"aws_api_gateway_resource":                                         43,
		"aws_api_gateway_rest_api":                                         44,
		"aws_api_gateway_rest_api_policy":                                  45,
		"aws_api_gateway_stage":                                            46,
		"aws_api_gateway_usage_plan":                                       47,
		"aws_api_gateway_usage_plan_key":                                   48,
		"aws_api_gateway_vpc_link":                                         49,
		"aws_apigatewayv2_api":                                             50,
		"aws_apigatewayv2_api_mapping":                                     51,
		"aws_apigatewayv2_authorizer":                                      52,
		"aws_apigatewayv2_deployment":                                      53,
		"aws_apigatewayv2_domain_name":                                     54,
		"aws_apigatewayv2_integration":                                     55,
		"aws_apigatewayv2_integration_response":                            56,
		"aws_apigatewayv2_model":                                           57,
		"aws_apigatewayv2_route":                                           58,
		"aws_apigatewayv2_route_response":                                  59,
		"aws_apigatewayv2_stage":                                           60,
		"aws_apigatewayv2_vpc_link":                                        61,
		"aws_app_cookie_stickiness_policy":                                 62,
		"aws_appautoscaling_policy":                                        63,
		"aws_appautoscaling_scheduled_action":                              64,
		"aws_appautoscaling_target":                                        65,
		"aws_appconfig_application":                                        66,
		"aws_appconfig_configuration_profile":                              67,
		"aws_appconfig_deployment":                                         68,
		"aws_appconfig_deployment_strategy":                                69,
		"aws_appconfig_environment":                                        70,
		"aws_appconfig_extension":                                          71,
		"aws_appconfig_extension_association":                              72,
		"aws_appconfig_hosted_configuration_version":                       73,
		"aws_appflow_connector_profile":                                    74,
		"aws_appflow_flow":                                                 75,
		"aws_appintegrations_data_integration":                             76,
		"aws_appintegrations_event_integration":                            77,
		"aws_applicationinsights_application":                              78,
		"aws_appmesh_gateway_route":                                        79,
		"aws_appmesh_mesh":                                                 80,
		"aws_appmesh_route":                                                81,
		"aws_appmesh_virtual_gateway":                                      82,
		"aws_appmesh_virtual_node":                                         83,
		"aws_appmesh_virtual_router":                                       84,
		"aws_appmesh_virtual_service":                                      85,
		"aws_apprunner_auto_scaling_configuration_version":                 86,
		"aws_apprunner_connection":                                         87,
		"aws_apprunner_custom_domain_association":                          88,
		"aws_apprunner_observability_configuration":                        89,
		"aws_apprunner_service":                                            90,
		"aws_apprunner_vpc_connector":                                      91,
		"aws_apprunner_vpc_ingress_connection":                             92,
		"aws_appstream_directory_config":                                   93,
		"aws_appstream_fleet":                                              94,
		"aws_appstream_fleet_stack_association":                            95,
		"aws_appstream_image_builder":                                      96,
		"aws_appstream_stack":                                              97,
		"aws_appstream_user":                                               98,
		"aws_appstream_user_stack_association":                             99,
		"aws_appsync_api_cache":                                            100,
		"aws_appsync_api_key":                                              101,
		"aws_appsync_datasource":                                           102,
		"aws_appsync_domain_name":                                          103,
		"aws_appsync_domain_name_api_association":                          104,
		"aws_appsync_function":                                             105,
		"aws_appsync_graphql_api":                                          106,
		"aws_appsync_resolver":                                             107,
		"aws_appsync_type":                                                 108,
		"aws_athena_data_catalog":                                          109,
		"aws_athena_database":                                              110,
		"aws_athena_named_query":                                           111,
		"aws_athena_workgroup":                                             112,
		"aws_auditmanager_account_registration":                            113,
		"aws_auditmanager_assessment":                                      114,
		"aws_auditmanager_assessment_delegation":                           115,
		"aws_auditmanager_assessment_report":                               116,
		"aws_auditmanager_control":                                         117,
		"aws_auditmanager_framework":                                       118,
		"aws_auditmanager_framework_share":                                 119,
		"aws_auditmanager_organization_admin_account_registration":         120,
		"aws_autoscaling_attachment":                                       121,
		"aws_autoscaling_group":                                            122,
		"aws_autoscaling_group_tag":                                        123,
		"aws_autoscaling_lifecycle_hook":                                   124,
		"aws_autoscaling_notification":                                     125,
		"aws_autoscaling_policy":                                           126,
		"aws_autoscaling_schedule":                                         127,
		"aws_autoscaling_traffic_source_attachment":                        128,
		"aws_autoscalingplans_scaling_plan":                                129,
		"aws_backup_framework":                                             130,
		"aws_backup_global_settings":                                       131,
		"aws_backup_plan":                                                  132,
		"aws_backup_region_settings":                                       133,
		"aws_backup_report_plan":                                           134,
		"aws_backup_selection":                                             135,
		"aws_backup_vault":                                                 136,
		"aws_backup_vault_lock_configuration":                              137,
		"aws_backup_vault_notifications":                                   138,
		"aws_backup_vault_policy":                                          139,
		"aws_batch_compute_environment":                                    140,
		"aws_batch_job_definition":                                         141,
		"aws_batch_job_queue":                                              142,
		"aws_batch_scheduling_policy":                                      143,
		"aws_budgets_budget":                                               144,
		"aws_budgets_budget_action":                                        145,
		"aws_ce_anomaly_monitor":                                           146,
		"aws_ce_anomaly_subscription":                                      147,
		"aws_ce_cost_allocation_tag":                                       148,
		"aws_ce_cost_category":                                             149,
		"aws_chime_voice_connector":                                        150,
		"aws_chime_voice_connector_group":                                  151,
		"aws_chime_voice_connector_logging":                                152,
		"aws_chime_voice_connector_origination":                            153,
		"aws_chime_voice_connector_streaming":                              154,
		"aws_chime_voice_connector_termination":                            155,
		"aws_chime_voice_connector_termination_credentials":                156,
		"aws_chimesdkmediapipelines_media_insights_pipeline_configuration": 157,
		"aws_chimesdkvoice_global_settings":                                158,
		"aws_chimesdkvoice_sip_media_application":                          159,
		"aws_chimesdkvoice_sip_rule":                                       160,
		"aws_chimesdkvoice_voice_profile_domain":                           161,
		"aws_cleanrooms_collaboration":                                     162,
		"aws_cloud9_environment_ec2":                                       163,
		"aws_cloud9_environment_membership":                                164,
		"aws_cloudcontrolapi_resource":                                     165,
		"aws_cloudformation_stack":                                         166,
		"aws_cloudformation_stack_set":                                     167,
		"aws_cloudformation_stack_set_instance":                            168,
		"aws_cloudformation_type":                                          169,
		"aws_cloudfront_cache_policy":                                      170,
		"aws_cloudfront_continuous_deployment_policy":                      171,
		"aws_cloudfront_distribution":                                      172,
		"aws_cloudfront_field_level_encryption_config":                     173,
		"aws_cloudfront_field_level_encryption_profile":                    174,
		"aws_cloudfront_function":                                          175,
		"aws_cloudfront_key_group":                                         176,
		"aws_cloudfront_monitoring_subscription":                           177,
		"aws_cloudfront_origin_access_control":                             178,
		"aws_cloudfront_origin_access_identity":                            179,
		"aws_cloudfront_origin_request_policy":                             180,
		"aws_cloudfront_public_key":                                        181,
		"aws_cloudfront_realtime_log_config":                               182,
		"aws_cloudfront_response_headers_policy":                           183,
		"aws_cloudhsm_v2_cluster":                                          184,
		"aws_cloudhsm_v2_hsm":                                              185,
		"aws_cloudsearch_domain":                                           186,
		"aws_cloudsearch_domain_service_access_policy":                     187,
		"aws_cloudtrail":                                                   188,
		"aws_cloudtrail_event_data_store":                                  189,
		"aws_cloudwatch_composite_alarm":                                   190,
		"aws_cloudwatch_dashboard":                                         191,
		"aws_cloudwatch_event_api_destination":                             192,
		"aws_cloudwatch_event_archive":                                     193,
		"aws_cloudwatch_event_bus":                                         194,
		"aws_cloudwatch_event_bus_policy":                                  195,
		"aws_cloudwatch_event_connection":                                  196,
		"aws_cloudwatch_event_endpoint":                                    197,
		"aws_cloudwatch_event_permission":                                  198,
		"aws_cloudwatch_event_rule":                                        199,
		"aws_cloudwatch_event_target":                                      200,
		"aws_cloudwatch_log_data_protection_policy":                        201,
		"aws_cloudwatch_log_destination":                                   202,
		"aws_cloudwatch_log_destination_policy":                            203,
		"aws_cloudwatch_log_group":                                         204,
		"aws_cloudwatch_log_metric_filter":                                 205,
		"aws_cloudwatch_log_resource_policy":                               206,
		"aws_cloudwatch_log_stream":                                        207,
		"aws_cloudwatch_log_subscription_filter":                           208,
		"aws_cloudwatch_metric_alarm":                                      209,
		"aws_cloudwatch_metric_stream":                                     210,
		"aws_cloudwatch_query_definition":                                  211,
		"aws_codeartifact_domain":                                          212,
		"aws_codeartifact_domain_permissions_policy":                       213,
		"aws_codeartifact_repository":                                      214,
		"aws_codeartifact_repository_permissions_policy":                   215,
		"aws_codebuild_project":                                            216,
		"aws_codebuild_report_group":                                       217,
		"aws_codebuild_resource_policy":                                    218,
		"aws_codebuild_source_credential":                                  219,
		"aws_codebuild_webhook":                                            220,
		"aws_codecatalyst_dev_environment":                                 221,
		"aws_codecatalyst_project":                                         222,
		"aws_codecatalyst_source_repository":                               223,
		"aws_codecommit_approval_rule_template":                            224,
		"aws_codecommit_approval_rule_template_association":                225,
		"aws_codecommit_repository":                                        226,
		"aws_codecommit_trigger":                                           227,
		"aws_codedeploy_app":                                               228,
		"aws_codedeploy_deployment_config":                                 229,
		"aws_codedeploy_deployment_group":                                  230,
		"aws_codegurureviewer_repository_association":                      231,
		"aws_codepipeline":                                                 232,
		"aws_codepipeline_custom_action_type":                              233,
		"aws_codepipeline_webhook":                                         234,
		"aws_codestarconnections_connection":                               235,
		"aws_codestarconnections_host":                                     236,
		"aws_codestarnotifications_notification_rule":                      237,
		"aws_cognito_identity_pool":                                        238,
		"aws_cognito_identity_pool_provider_principal_tag":                 239,
		"aws_cognito_identity_pool_roles_attachment":                       240,
		"aws_cognito_identity_provider":                                    241,
		"aws_cognito_managed_user_pool_client":                             242,
		"aws_cognito_resource_server":                                      243,
		"aws_cognito_risk_configuration":                                   244,
		"aws_cognito_user":                                                 245,
		"aws_cognito_user_group":                                           246,
		"aws_cognito_user_in_group":                                        247,
		"aws_cognito_user_pool":                                            248,
		"aws_cognito_user_pool_client":                                     249,
		"aws_cognito_user_pool_domain":                                     250,
		"aws_cognito_user_pool_ui_customization":                           251,
		"aws_comprehend_document_classifier":                               252,
		"aws_comprehend_entity_recognizer":                                 253,
		"aws_config_aggregate_authorization":                               254,
		"aws_config_config_rule":                                           255,
		"aws_config_configuration_aggregator":                              256,
		"aws_config_configuration_recorder":                                257,
		"aws_config_configuration_recorder_status":                         258,
		"aws_config_conformance_pack":                                      259,
		"aws_config_delivery_channel":                                      260,
		"aws_config_organization_conformance_pack":                         261,
		"aws_config_organization_custom_policy_rule":                       262,
		"aws_config_organization_custom_rule":                              263,
		"aws_config_organization_managed_rule":                             264,
		"aws_config_remediation_configuration":                             265,
		"aws_connect_bot_association":                                      266,
		"aws_connect_contact_flow":                                         267,
		"aws_connect_contact_flow_module":                                  268,
		"aws_connect_hours_of_operation":                                   269,
		"aws_connect_instance":                                             270,
		"aws_connect_instance_storage_config":                              271,
		"aws_connect_lambda_function_association":                          272,
		"aws_connect_phone_number":                                         273,
		"aws_connect_queue":                                                274,
		"aws_connect_quick_connect":                                        275,
		"aws_connect_routing_profile":                                      276,
		"aws_connect_security_profile":                                     277,
		"aws_connect_user":                                                 278,
		"aws_connect_user_hierarchy_group":                                 279,
		"aws_connect_user_hierarchy_structure":                             280,
		"aws_connect_vocabulary":                                           281,
		"aws_controltower_control":                                         282,
		"aws_cur_report_definition":                                        283,
		"aws_customer_gateway":                                             284,
		"aws_dataexchange_data_set":                                        285,
		"aws_dataexchange_revision":                                        286,
		"aws_datapipeline_pipeline":                                        287,
		"aws_datapipeline_pipeline_definition":                             288,
		"aws_datasync_agent":                                               289,
		"aws_datasync_location_azure_blob":                                 290,
		"aws_datasync_location_efs":                                        291,
		"aws_datasync_location_fsx_lustre_file_system":                     292,
		"aws_datasync_location_fsx_ontap_file_system":                      293,
		"aws_datasync_location_fsx_openzfs_file_system":                    294,
		"aws_datasync_location_fsx_windows_file_system":                    295,
		"aws_datasync_location_hdfs":                                       296,
		"aws_datasync_location_nfs":                                        297,
		"aws_datasync_location_object_storage":                             298,
		"aws_datasync_location_s3":                                         299,
		"aws_datasync_location_smb":                                        300,
		"aws_datasync_task":                                                301,
		"aws_dax_cluster":                                                  302,
		"aws_dax_parameter_group":                                          303,
		"aws_dax_subnet_group":                                             304,
		"aws_db_cluster_snapshot":                                          305,
		"aws_db_event_subscription":                                        306,
		"aws_db_instance":                                                  307,
		"aws_db_instance_automated_backups_replication":                    308,
		"aws_db_instance_role_association":                                 309,
		"aws_db_option_group":                                              310,
		"aws_db_parameter_group":                                           311,
		"aws_db_proxy":                                                     312,
		"aws_db_proxy_default_target_group":                                313,
		"aws_db_proxy_endpoint":                                            314,
		"aws_db_proxy_target":                                              315,
		"aws_db_snapshot":                                                  316,
		"aws_db_snapshot_copy":                                             317,
		"aws_db_subnet_group":                                              318,
		"aws_default_network_acl":                                          319,
		"aws_default_route_table":                                          320,
		"aws_default_security_group":                                       321,
		"aws_default_subnet":                                               322,
		"aws_default_vpc":                                                  323,
		"aws_default_vpc_dhcp_options":                                     324,
		"aws_detective_graph":                                              325,
		"aws_detective_invitation_accepter":                                326,
		"aws_detective_member":                                             327,
		"aws_devicefarm_device_pool":                                       328,
		"aws_devicefarm_instance_profile":                                  329,
		"aws_devicefarm_network_profile":                                   330,
		"aws_devicefarm_project":                                           331,
		"aws_devicefarm_test_grid_project":                                 332,
		"aws_devicefarm_upload":                                            333,
		"aws_directory_service_conditional_forwarder":                      334,
		"aws_directory_service_directory":                                  335,
		"aws_directory_service_log_subscription":                           336,
		"aws_directory_service_radius_settings":                            337,
		"aws_directory_service_region":                                     338,
		"aws_directory_service_shared_directory":                           339,
		"aws_directory_service_shared_directory_accepter":                  340,
		"aws_directory_service_trust":                                      341,
		"aws_dlm_lifecycle_policy":                                         342,
		"aws_dms_certificate":                                              343,
		"aws_dms_endpoint":                                                 344,
		"aws_dms_event_subscription":                                       345,
		"aws_dms_replication_instance":                                     346,
		"aws_dms_replication_subnet_group":                                 347,
		"aws_dms_replication_task":                                         348,
		"aws_dms_s3_endpoint":                                              349,
		"aws_docdb_cluster":                                                350,
		"aws_docdb_cluster_instance":                                       351,
		"aws_docdb_cluster_parameter_group":                                352,
		"aws_docdb_cluster_snapshot":                                       353,
		"aws_docdb_event_subscription":                                     354,
		"aws_docdb_global_cluster":                                         355,
		"aws_docdb_subnet_group":                                           356,
		"aws_dx_bgp_peer":                                                  357,
		"aws_dx_connection":                                                358,
		"aws_dx_connection_association":                                    359,
		"aws_dx_connection_confirmation":                                   360,
		"aws_dx_gateway":                                                   361,
		"aws_dx_gateway_association":                                       362,
		"aws_dx_gateway_association_proposal":                              363,
		"aws_dx_hosted_connection":                                         364,
		"aws_dx_hosted_private_virtual_interface":                          365,
		"aws_dx_hosted_private_virtual_interface_accepter":                 366,
		"aws_dx_hosted_public_virtual_interface":                           367,
		"aws_dx_hosted_public_virtual_interface_accepter":                  368,
		"aws_dx_hosted_transit_virtual_interface":                          369,
		"aws_dx_hosted_transit_virtual_interface_accepter":                 370,
		"aws_dx_lag":                                                       371,
		"aws_dx_macsec_key_association":                                    372,
		"aws_dx_private_virtual_interface":                                 373,
		"aws_dx_public_virtual_interface":                                  374,
		"aws_dx_transit_virtual_interface":                                 375,
		"aws_dynamodb_contributor_insights":                                376,
		"aws_dynamodb_global_table":                                        377,
		"aws_dynamodb_kinesis_streaming_destination":                       378,
		"aws_dynamodb_table":                                               379,
		"aws_dynamodb_table_item":                                          380,
		"aws_dynamodb_table_replica":                                       381,
		"aws_dynamodb_tag":                                                 382,
		"aws_ebs_default_kms_key":                                          383,
		"aws_ebs_encryption_by_default":                                    384,
		"aws_ebs_snapshot":                                                 385,
		"aws_ebs_snapshot_copy":                                            386,
		"aws_ebs_snapshot_import":                                          387,
		"aws_ebs_volume":                                                   388,
		"aws_ec2_availability_zone_group":                                  389,
		"aws_ec2_capacity_reservation":                                     390,
		"aws_ec2_carrier_gateway":                                          391,
		"aws_ec2_client_vpn_authorization_rule":                            392,
		"aws_ec2_client_vpn_endpoint":                                      393,
		"aws_ec2_client_vpn_network_association":                           394,
		"aws_ec2_client_vpn_route":                                         395,
		"aws_ec2_fleet":                                                    396,
		"aws_ec2_host":                                                     397,
		"aws_ec2_instance_connect_endpoint":                                398,
		"aws_ec2_instance_state":                                           399,
		"aws_ec2_local_gateway_route":                                      400,
		"aws_ec2_local_gateway_route_table_vpc_association":                401,
		"aws_ec2_managed_prefix_list":                                      402,
		"aws_ec2_managed_prefix_list_entry":                                403,
		"aws_ec2_network_insights_analysis":                                404,
		"aws_ec2_network_insights_path":                                    405,
		"aws_ec2_serial_console_access":                                    406,
		"aws_ec2_subnet_cidr_reservation":                                  407,
		"aws_ec2_tag":                                                      408,
		"aws_ec2_traffic_mirror_filter":                                    409,
		"aws_ec2_traffic_mirror_filter_rule":                               410,
		"aws_ec2_traffic_mirror_session":                                   411,
		"aws_ec2_traffic_mirror_target":                                    412,
		"aws_ec2_transit_gateway":                                          413,
		"aws_ec2_transit_gateway_connect":                                  414,
		"aws_ec2_transit_gateway_connect_peer":                             415,
		"aws_ec2_transit_gateway_multicast_domain":                         416,
		"aws_ec2_transit_gateway_multicast_domain_association":             417,
		"aws_ec2_transit_gateway_multicast_group_member":                   418,
		"aws_ec2_transit_gateway_multicast_group_source":                   419,
		"aws_ec2_transit_gateway_peering_attachment":                       420,
		"aws_ec2_transit_gateway_peering_attachment_accepter":              421,
		"aws_ec2_transit_gateway_policy_table":                             422,
		"aws_ec2_transit_gateway_policy_table_association":                 423,
		"aws_ec2_transit_gateway_prefix_list_reference":                    424,
		"aws_ec2_transit_gateway_route":                                    425,
		"aws_ec2_transit_gateway_route_table":                              426,
		"aws_ec2_transit_gateway_route_table_association":                  427,
		"aws_ec2_transit_gateway_route_table_propagation":                  428,
		"aws_ec2_transit_gateway_vpc_attachment":                           429,
		"aws_ec2_transit_gateway_vpc_attachment_accepter":                  430,
		"aws_ecr_lifecycle_policy":                                         431,
		"aws_ecr_pull_through_cache_rule":                                  432,
		"aws_ecr_registry_policy":                                          433,
		"aws_ecr_registry_scanning_configuration":                          434,
		"aws_ecr_replication_configuration":                                435,
		"aws_ecr_repository":                                               436,
		"aws_ecr_repository_policy":                                        437,
		"aws_ecrpublic_repository":                                         438,
		"aws_ecrpublic_repository_policy":                                  439,
		"aws_ecs_account_setting_default":                                  440,
		"aws_ecs_capacity_provider":                                        441,
		"aws_ecs_cluster":                                                  442,
		"aws_ecs_cluster_capacity_providers":                               443,
		"aws_ecs_service":                                                  444,
		"aws_ecs_tag":                                                      445,
		"aws_ecs_task_definition":                                          446,
		"aws_ecs_task_set":                                                 447,
		"aws_efs_access_point":                                             448,
		"aws_efs_backup_policy":                                            449,
		"aws_efs_file_system":                                              450,
		"aws_efs_file_system_policy":                                       451,
		"aws_efs_mount_target":                                             452,
		"aws_efs_replication_configuration":                                453,
		"aws_egress_only_internet_gateway":                                 454,
		"aws_eip":                                                          455,
		"aws_eip_association":                                              456,
		"aws_eks_addon":                                                    457,
		"aws_eks_cluster":                                                  458,
		"aws_eks_fargate_profile":                                          459,
		"aws_eks_identity_provider_config":                                 460,
		"aws_eks_node_group":                                               461,
		"aws_elastic_beanstalk_application":                                462,
		"aws_elastic_beanstalk_application_version":                        463,
		"aws_elastic_beanstalk_configuration_template":                     464,
		"aws_elastic_beanstalk_environment":                                465,
		"aws_elasticache_cluster":                                          466,
		"aws_elasticache_global_replication_group":                         467,
		"aws_elasticache_parameter_group":                                  468,
		"aws_elasticache_replication_group":                                469,
		"aws_elasticache_subnet_group":                                     470,
		"aws_elasticache_user":                                             471,
		"aws_elasticache_user_group":                                       472,
		"aws_elasticache_user_group_association":                           473,
		"aws_elasticsearch_domain":                                         474,
		"aws_elasticsearch_domain_policy":                                  475,
		"aws_elasticsearch_domain_saml_options":                            476,
		"aws_elastictranscoder_pipeline":                                   477,
		"aws_elastictranscoder_preset":                                     478,
		"aws_elb":                                                          479,
		"aws_elb_attachment":                                               480,
		"aws_emr_block_public_access_configuration":                        481,
		"aws_emr_cluster":                                                  482,
		"aws_emr_instance_fleet":                                           483,
		"aws_emr_instance_group":                                           484,
		"aws_emr_managed_scaling_policy":                                   485,
		"aws_emr_security_configuration":                                   486,
		"aws_emr_studio":                                                   487,
		"aws_emr_studio_session_mapping":                                   488,
		"aws_emrcontainers_job_template":                                   489,
		"aws_emrcontainers_virtual_cluster":                                490,
		"aws_emrserverless_application":                                    491,
		"aws_evidently_feature":                                            492,
		"aws_evidently_launch":                                             493,
		"aws_evidently_project":                                            494,
		"aws_evidently_segment":                                            495,
		"aws_finspace_kx_cluster":                                          496,
		"aws_finspace_kx_database":                                         497,
		"aws_finspace_kx_environment":                                      498,
		"aws_finspace_kx_user":                                             499,
		"aws_fis_experiment_template":                                      500,
		"aws_flow_log":                                                     501,
		"aws_fms_admin_account":                                            502,
		"aws_fms_policy":                                                   503,
		"aws_fsx_backup":                                                   504,
		"aws_fsx_data_repository_association":                              505,
		"aws_fsx_file_cache":                                               506,
		"aws_fsx_lustre_file_system":                                       507,
		"aws_fsx_ontap_file_system":                                        508,
		"aws_fsx_ontap_storage_virtual_machine":                            509,
		"aws_fsx_ontap_volume":                                             510,
		"aws_fsx_openzfs_file_system":                                      511,
		"aws_fsx_openzfs_snapshot":                                         512,
		"aws_fsx_openzfs_volume":                                           513,
		"aws_fsx_windows_file_system":                                      514,
		"aws_gamelift_alias":                                               515,
		"aws_gamelift_build":                                               516,
		"aws_gamelift_fleet":                                               517,
		"aws_gamelift_game_server_group":                                   518,
		"aws_gamelift_game_session_queue":                                  519,
		"aws_gamelift_script":                                              520,
		"aws_glacier_vault":                                                521,
		"aws_glacier_vault_lock":                                           522,
		"aws_globalaccelerator_accelerator":                                523,
		"aws_globalaccelerator_custom_routing_accelerator":                 524,
		"aws_globalaccelerator_custom_routing_endpoint_group":              525,
		"aws_globalaccelerator_custom_routing_listener":                    526,
		"aws_globalaccelerator_endpoint_group":                             527,
		"aws_globalaccelerator_listener":                                   528,
		"aws_glue_catalog_database":                                        529,
		"aws_glue_catalog_table":                                           530,
		"aws_glue_classifier":                                              531,
		"aws_glue_connection":                                              532,
		"aws_glue_crawler":                                                 533,
		"aws_glue_data_catalog_encryption_settings":                        534,
		"aws_glue_data_quality_ruleset":                                    535,
		"aws_glue_dev_endpoint":                                            536,
		"aws_glue_job":                                                     537,
		"aws_glue_ml_transform":                                            538,
		"aws_glue_partition":                                               539,
		"aws_glue_partition_index":                                         540,
		"aws_glue_registry":                                                541,
		"aws_glue_resource_policy":                                         542,
		"aws_glue_schema":                                                  543,
		"aws_glue_security_configuration":                                  544,
		"aws_glue_trigger":                                                 545,
		"aws_glue_user_defined_function":                                   546,
		"aws_glue_workflow":                                                547,
		"aws_grafana_license_association":                                  548,
		"aws_grafana_role_association":                                     549,
		"aws_grafana_workspace":                                            550,
		"aws_grafana_workspace_api_key":                                    551,
		"aws_grafana_workspace_saml_configuration":                         552,
		"aws_guardduty_detector":                                           553,
		"aws_guardduty_filter":                                             554,
		"aws_guardduty_invite_accepter":                                    555,
		"aws_guardduty_ipset":                                              556,
		"aws_guardduty_member":                                             557,
		"aws_guardduty_organization_admin_account":                         558,
		"aws_guardduty_organization_configuration":                         559,
		"aws_guardduty_publishing_destination":                             560,
		"aws_guardduty_threatintelset":                                     561,
		"aws_iam_access_key":                                               562,
		"aws_iam_account_alias":                                            563,
		"aws_iam_account_password_policy":                                  564,
		"aws_iam_group":                                                    565,
		"aws_iam_group_membership":                                         566,
		"aws_iam_group_policy":                                             567,
		"aws_iam_group_policy_attachment":                                  568,
		"aws_iam_instance_profile":                                         569,
		"aws_iam_openid_connect_provider":                                  570,
		"aws_iam_policy":                                                   571,
		"aws_iam_policy_attachment":                                        572,
		"aws_iam_role":                                                     573,
		"aws_iam_role_policy":                                              574,
		"aws_iam_role_policy_attachment":                                   575,
		"aws_iam_saml_provider":                                            576,
		"aws_iam_security_token_service_preferences":                       577,
		"aws_iam_server_certificate":                                       578,
		"aws_iam_service_linked_role":                                      579,
		"aws_iam_service_specific_credential":                              580,
		"aws_iam_signing_certificate":                                      581,
		"aws_iam_user":                                                     582,
		"aws_iam_user_group_membership":                                    583,
		"aws_iam_user_login_profile":                                       584,
		"aws_iam_user_policy":                                              585,
		"aws_iam_user_policy_attachment":                                   586,
		"aws_iam_user_ssh_key":                                             587,
		"aws_iam_virtual_mfa_device":                                       588,
		"aws_identitystore_group":                                          589,
		"aws_identitystore_group_membership":                               590,
		"aws_identitystore_user":                                           591,
		"aws_imagebuilder_component":                                       592,
		"aws_imagebuilder_container_recipe":                                593,
		"aws_imagebuilder_distribution_configuration":                      594,
		"aws_imagebuilder_image":                                           595,
		"aws_imagebuilder_image_pipeline":                                  596,
		"aws_imagebuilder_image_recipe":                                    597,
		"aws_imagebuilder_infrastructure_configuration":                    598,
		"aws_inspector2_delegated_admin_account":                           599,
		"aws_inspector2_enabler":                                           600,
		"aws_inspector2_member_association":                                601,
		"aws_inspector2_organization_configuration":                        602,
		"aws_inspector_assessment_target":                                  603,
		"aws_inspector_assessment_template":                                604,
		"aws_inspector_resource_group":                                     605,
		"aws_instance":                                                     606,
		"aws_internet_gateway":                                             607,
		"aws_internet_gateway_attachment":                                  608,
		"aws_internetmonitor_monitor":                                      609,
		"aws_iot_authorizer":                                               610,
		"aws_iot_certificate":                                              611,
		"aws_iot_indexing_configuration":                                   612,
		"aws_iot_logging_options":                                          613,
		"aws_iot_policy":                                                   614,
		"aws_iot_policy_attachment":                                        615,
		"aws_iot_provisioning_template":                                    616,
		"aws_iot_role_alias":                                               617,
		"aws_iot_thing":                                                    618,
		"aws_iot_thing_group":                                              619,
		"aws_iot_thing_group_membership":                                   620,
		"aws_iot_thing_principal_attachment":                               621,
		"aws_iot_thing_type":                                               622,
		"aws_iot_topic_rule":                                               623,
		"aws_iot_topic_rule_destination":                                   624,
		"aws_ivs_channel":                                                  625,
		"aws_ivs_playback_key_pair":                                        626,
		"aws_ivs_recording_configuration":                                  627,
		"aws_ivschat_logging_configuration":                                628,
		"aws_ivschat_room":                                                 629,
		"aws_kendra_data_source":                                           630,
		"aws_kendra_experience":                                            631,
		"aws_kendra_faq":                                                   632,
		"aws_kendra_index":                                                 633,
		"aws_kendra_query_suggestions_block_list":                          634,
		"aws_kendra_thesaurus":                                             635,
		"aws_key_pair":                                                     636,
		"aws_keyspaces_keyspace":                                           637,
		"aws_keyspaces_table":                                              638,
		"aws_kinesis_analytics_application":                                639,
		"aws_kinesis_firehose_delivery_stream":                             640,
		"aws_kinesis_stream":                                               641,
		"aws_kinesis_stream_consumer":                                      642,
		"aws_kinesis_video_stream":                                         643,
		"aws_kinesisanalyticsv2_application":                               644,
		"aws_kinesisanalyticsv2_application_snapshot":                      645,
		"aws_kms_alias":                                                    646,
		"aws_kms_ciphertext":                                               647,
		"aws_kms_custom_key_store":                                         648,
		"aws_kms_external_key":                                             649,
		"aws_kms_grant":                                                    650,
		"aws_kms_key":                                                      651,
		"aws_kms_key_policy":                                               652,
		"aws_kms_replica_external_key":                                     653,
		"aws_kms_replica_key":                                              654,
		"aws_lakeformation_data_lake_settings":                             655,
		"aws_lakeformation_lf_tag":                                         656,
		"aws_lakeformation_permissions":                                    657,
		"aws_lakeformation_resource":                                       658,
		"aws_lakeformation_resource_lf_tags":                               659,
		"aws_lambda_alias":                                                 660,
		"aws_lambda_code_signing_config":                                   661,
		"aws_lambda_event_source_mapping":                                  662,
		"aws_lambda_function":                                              663,
		"aws_lambda_function_event_invoke_config":                          664,
		"aws_lambda_function_url":                                          665,
		"aws_lambda_invocation":                                            666,
		"aws_lambda_layer_version":                                         667,
		"aws_lambda_layer_version_permission":                              668,
		"aws_lambda_permission":                                            669,
		"aws_lambda_provisioned_concurrency_config":                        670,
		"aws_launch_configuration":                                         671,
		"aws_launch_template":                                              672,
		"aws_lb":                                                           673,
		"aws_lb_cookie_stickiness_policy":                                  674,
		"aws_lb_listener":                                                  675,
		"aws_lb_listener_certificate":                                      676,
		"aws_lb_listener_rule":                                             677,
		"aws_lb_ssl_negotiation_policy":                                    678,
		"aws_lb_target_group":                                              679,
		"aws_lb_target_group_attachment":                                   680,
		"aws_lex_bot":                                                      681,
		"aws_lex_bot_alias":                                                682,
		"aws_lex_intent":                                                   683,
		"aws_lex_slot_type":                                                684,
		"aws_licensemanager_association":                                   685,
		"aws_licensemanager_grant":                                         686,
		"aws_licensemanager_grant_accepter":                                687,
		"aws_licensemanager_license_configuration":                         688,
		"aws_lightsail_bucket":                                             689,
		"aws_lightsail_bucket_access_key":                                  690,
		"aws_lightsail_bucket_resource_access":                             691,
		"aws_lightsail_certificate":                                        692,
		"aws_lightsail_container_service":                                  693,
		"aws_lightsail_container_service_deployment_version":               694,
		"aws_lightsail_database":                                           695,
		"aws_lightsail_disk":                                               696,
		"aws_lightsail_disk_attachment":                                    697,
		"aws_lightsail_distribution":                                       698,
		"aws_lightsail_domain":                                             699,
		"aws_lightsail_domain_entry":                                       700,
		"aws_lightsail_instance":                                           701,
		"aws_lightsail_instance_public_ports":                              702,
		"aws_lightsail_key_pair":                                           703,
		"aws_lightsail_lb":                                                 704,
		"aws_lightsail_lb_attachment":                                      705,
		"aws_lightsail_lb_certificate":                                     706,
		"aws_lightsail_lb_certificate_attachment":                          707,
		"aws_lightsail_lb_https_redirection_policy":                        708,
		"aws_lightsail_lb_stickiness_policy":                               709,
		"aws_lightsail_static_ip":                                          710,
		"aws_lightsail_static_ip_attachment":                               711,
		"aws_load_balancer_backend_server_policy":                          712,
		"aws_load_balancer_listener_policy":                                713,
		"aws_load_balancer_policy":                                         714,
		"aws_location_geofence_collection":                                 715,
		"aws_location_map":                                                 716,
		"aws_location_place_index":                                         717,
		"aws_location_route_calculator":                                    718,
		"aws_location_tracker":                                             719,
		"aws_location_tracker_association":                                 720,
		"aws_macie2_account":                                               721,
		"aws_macie2_classification_export_configuration":                   722,
		"aws_macie2_classification_job":                                    723,
		"aws_macie2_custom_data_identifier":                                724,
		"aws_macie2_findings_filter":                                       725,
		"aws_macie2_invitation_accepter":                                   726,
		"aws_macie2_member":                                                727,
		"aws_macie2_organization_admin_account":                            728,
		"aws_main_route_table_association":                                 729,
		"aws_media_convert_queue":                                          730,
		"aws_media_package_channel":                                        731,
		"aws_media_store_container":                                        732,
		"aws_media_store_container_policy":                                 733,
		"aws_medialive_channel":                                            734,
		"aws_medialive_input":                                              735,
		"aws_medialive_input_security_group":                               736,
		"aws_medialive_multiplex":                                          737,
		"aws_medialive_multiplex_program":                                  738,
		"aws_memorydb_acl":                                                 739,
		"aws_memorydb_cluster":                                             740,
		"aws_memorydb_parameter_group":                                     741,
		"aws_memorydb_snapshot":                                            742,
		"aws_memorydb_subnet_group":                                        743,
		"aws_memorydb_user":                                                744,
		"aws_mq_broker":                                                    745,
		"aws_mq_configuration":                                             746,
		"aws_msk_cluster":                                                  747,
		"aws_msk_cluster_policy":                                           748,
		"aws_msk_configuration":                                            749,
		"aws_msk_scram_secret_association":                                 750,
		"aws_msk_serverless_cluster":                                       751,
		"aws_msk_vpc_connection":                                           752,
		"aws_mskconnect_connector":                                         753,
		"aws_mskconnect_custom_plugin":                                     754,
		"aws_mskconnect_worker_configuration":                              755,
		"aws_mwaa_environment":                                             756,
		"aws_nat_gateway":                                                  757,
		"aws_neptune_cluster":                                              758,
		"aws_neptune_cluster_endpoint":                                     759,
		"aws_neptune_cluster_instance":                                     760,
		"aws_neptune_cluster_parameter_group":                              761,
		"aws_neptune_cluster_snapshot":                                     762,
		"aws_neptune_event_subscription":                                   763,
		"aws_neptune_global_cluster":                                       764,
		"aws_neptune_parameter_group":                                      765,
		"aws_neptune_subnet_group":                                         766,
		"aws_network_acl":                                                  767,
		"aws_network_acl_association":                                      768,
		"aws_network_acl_rule":                                             769,
		"aws_network_interface":                                            770,
		"aws_network_interface_attachment":                                 771,
		"aws_network_interface_sg_attachment":                              772,
		"aws_networkfirewall_firewall":                                     773,
		"aws_networkfirewall_firewall_policy":                              774,
		"aws_networkfirewall_logging_configuration":                        775,
		"aws_networkfirewall_resource_policy":                              776,
		"aws_networkfirewall_rule_group":                                   777,
		"aws_networkmanager_attachment_accepter":                           778,
		"aws_networkmanager_connect_attachment":                            779,
		"aws_networkmanager_connect_peer":                                  780,
		"aws_networkmanager_connection":                                    781,
		"aws_networkmanager_core_network":                                  782,
		"aws_networkmanager_core_network_policy_attachment":                783,
		"aws_networkmanager_customer_gateway_association":                  784,
		"aws_networkmanager_device":                                        785,
		"aws_networkmanager_global_network":                                786,
		"aws_networkmanager_link":                                          787,
		"aws_networkmanager_link_association":                              788,
		"aws_networkmanager_site":                                          789,
		"aws_networkmanager_site_to_site_vpn_attachment":                   790,
		"aws_networkmanager_transit_gateway_connect_peer_association":      791,
		"aws_networkmanager_transit_gateway_peering":                       792,
		"aws_networkmanager_transit_gateway_registration":                  793,
		"aws_networkmanager_transit_gateway_route_table_attachment":        794,
		"aws_networkmanager_vpc_attachment":                                795,
		"aws_oam_link":                                                     796,
		"aws_oam_sink":                                                     797,
		"aws_oam_sink_policy":                                              798,
		"aws_opensearch_domain":                                            799,
		"aws_opensearch_domain_policy":                                     800,
		"aws_opensearch_domain_saml_options":                               801,
		"aws_opensearch_inbound_connection_accepter":                       802,
		"aws_opensearch_outbound_connection":                               803,
		"aws_opensearch_vpc_endpoint":                                      804,
		"aws_opensearchserverless_access_policy":                           805,
		"aws_opensearchserverless_collection":                              806,
		"aws_opensearchserverless_security_config":                         807,
		"aws_opensearchserverless_security_policy":                         808,
		"aws_opensearchserverless_vpc_endpoint":                            809,
		"aws_opsworks_application":                                         810,
		"aws_opsworks_custom_layer":                                        811,
		"aws_opsworks_ecs_cluster_layer":                                   812,
		"aws_opsworks_ganglia_layer":                                       813,
		"aws_opsworks_haproxy_layer":                                       814,
		"aws_opsworks_instance":                                            815,
		"aws_opsworks_java_app_layer":                                      816,
		"aws_opsworks_memcached_layer":                                     817,
		"aws_opsworks_mysql_layer":                                         818,
		"aws_opsworks_nodejs_app_layer":                                    819,
		"aws_opsworks_permission":                                          820,
		"aws_opsworks_php_app_layer":                                       821,
		"aws_opsworks_rails_app_layer":                                     822,
		"aws_opsworks_rds_db_instance":                                     823,
		"aws_opsworks_stack":                                               824,
		"aws_opsworks_static_web_layer":                                    825,
		"aws_opsworks_user_profile":                                        826,
		"aws_organizations_account":                                        827,
		"aws_organizations_delegated_administrator":                        828,
		"aws_organizations_organization":                                   829,
		"aws_organizations_organizational_unit":                            830,
		"aws_organizations_policy":                                         831,
		"aws_organizations_policy_attachment":                              832,
		"aws_organizations_resource_policy":                                833,
		"aws_pinpoint_adm_channel":                                         834,
		"aws_pinpoint_apns_channel":                                        835,
		"aws_pinpoint_apns_sandbox_channel":                                836,
		"aws_pinpoint_apns_voip_channel":                                   837,
		"aws_pinpoint_apns_voip_sandbox_channel":                           838,
		"aws_pinpoint_app":                                                 839,
		"aws_pinpoint_baidu_channel":                                       840,
		"aws_pinpoint_email_channel":                                       841,
		"aws_pinpoint_event_stream":                                        842,
		"aws_pinpoint_gcm_channel":                                         843,
		"aws_pinpoint_sms_channel":                                         844,
		"aws_pipes_pipe":                                                   845,
		"aws_placement_group":                                              846,
		"aws_prometheus_alert_manager_definition":                          847,
		"aws_prometheus_rule_group_namespace":                              848,
		"aws_prometheus_workspace":                                         849,
		"aws_proxy_protocol_policy":                                        850,
		"aws_qldb_ledger":                                                  851,
		"aws_qldb_stream":                                                  852,
		"aws_quicksight_account_subscription":                              853,
		"aws_quicksight_analysis":                                          854,
		"aws_quicksight_dashboard":                                         855,
		"aws_quicksight_data_set":                                          856,
		"aws_quicksight_data_source":                                       857,
		"aws_quicksight_folder":                                            858,
		"aws_quicksight_folder_membership":                                 859,
		"aws_quicksight_group":                                             860,
		"aws_quicksight_group_membership":                                  861,
		"aws_quicksight_iam_policy_assignment":                             862,
		"aws_quicksight_ingestion":                                         863,
		"aws_quicksight_namespace":                                         864,
		"aws_quicksight_refresh_schedule":                                  865,
		"aws_quicksight_template":                                          866,
		"aws_quicksight_template_alias":                                    867,
		"aws_quicksight_theme":                                             868,
		"aws_quicksight_user":                                              869,
		"aws_quicksight_vpc_connection":                                    870,
		"aws_ram_principal_association":                                    871,
		"aws_ram_resource_association":                                     872,
		"aws_ram_resource_share":                                           873,
		"aws_ram_resource_share_accepter":                                  874,
		"aws_ram_sharing_with_organization":                                875,
		"aws_rbin_rule":                                                    876,
		"aws_rds_cluster":                                                  877,
		"aws_rds_cluster_activity_stream":                                  878,
		"aws_rds_cluster_endpoint":                                         879,
		"aws_rds_cluster_instance":                                         880,
		"aws_rds_cluster_parameter_group":                                  881,
		"aws_rds_cluster_role_association":                                 882,
		"aws_rds_export_task":                                              883,
		"aws_rds_global_cluster":                                           884,
		"aws_rds_reserved_instance":                                        885,
		"aws_redshift_authentication_profile":                              886,
		"aws_redshift_cluster":                                             887,
		"aws_redshift_cluster_iam_roles":                                   888,
		"aws_redshift_cluster_snapshot":                                    889,
		"aws_redshift_endpoint_access":                                     890,
		"aws_redshift_endpoint_authorization":                              891,
		"aws_redshift_event_subscription":                                  892,
		"aws_redshift_hsm_client_certificate":                              893,
		"aws_redshift_hsm_configuration":                                   894,
		"aws_redshift_parameter_group":                                     895,
		"aws_redshift_partner":                                             896,
		"aws_redshift_scheduled_action":                                    897,
		"aws_redshift_snapshot_copy_grant":                                 898,
		"aws_redshift_snapshot_schedule":                                   899,
		"aws_redshift_snapshot_schedule_association":                       900,
		"aws_redshift_subnet_group":                                        901,
		"aws_redshift_usage_limit":                                         902,
		"aws_redshiftdata_statement":                                       903,
		"aws_redshiftserverless_endpoint_access":                           904,
		"aws_redshiftserverless_namespace":                                 905,
		"aws_redshiftserverless_resource_policy":                           906,
		"aws_redshiftserverless_snapshot":                                  907,
		"aws_redshiftserverless_usage_limit":                               908,
		"aws_redshiftserverless_workgroup":                                 909,
		"aws_resourceexplorer2_index":                                      910,
		"aws_resourceexplorer2_view":                                       911,
		"aws_resourcegroups_group":                                         912,
		"aws_resourcegroups_resource":                                      913,
		"aws_rolesanywhere_profile":                                        914,
		"aws_rolesanywhere_trust_anchor":                                   915,
		"aws_route":                                                        916,
		"aws_route53_cidr_collection":                                      917,
		"aws_route53_cidr_location":                                        918,
		"aws_route53_delegation_set":                                       919,
		"aws_route53_health_check":                                         920,
		"aws_route53_hosted_zone_dnssec":                                   921,
		"aws_route53_key_signing_key":                                      922,
		"aws_route53_query_log":                                            923,
		"aws_route53_record":                                               924,
		"aws_route53_resolver_config":                                      925,
		"aws_route53_resolver_dnssec_config":                               926,
		"aws_route53_resolver_endpoint":                                    927,
		"aws_route53_resolver_firewall_config":                             928,
		"aws_route53_resolver_firewall_domain_list":                        929,
		"aws_route53_resolver_firewall_rule":                               930,
		"aws_route53_resolver_firewall_rule_group":                         931,
		"aws_route53_resolver_firewall_rule_group_association":             932,
		"aws_route53_resolver_query_log_config":                            933,
		"aws_route53_resolver_query_log_config_association":                934,
		"aws_route53_resolver_rule":                                        935,
		"aws_route53_resolver_rule_association":                            936,
		"aws_route53_traffic_policy":                                       937,
		"aws_route53_traffic_policy_instance":                              938,
		"aws_route53_vpc_association_authorization":                        939,
		"aws_route53_zone":                                                 940,
		"aws_route53_zone_association":                                     941,
		"aws_route53domains_registered_domain":                             942,
		"aws_route53recoverycontrolconfig_cluster":                         943,
		"aws_route53recoverycontrolconfig_control_panel":                   944,
		"aws_route53recoverycontrolconfig_routing_control":                 945,
		"aws_route53recoverycontrolconfig_safety_rule":                     946,
		"aws_route53recoveryreadiness_cell":                                947,
		"aws_route53recoveryreadiness_readiness_check":                     948,
		"aws_route53recoveryreadiness_recovery_group":                      949,
		"aws_route53recoveryreadiness_resource_set":                        950,
		"aws_route_table":                                                  951,
		"aws_route_table_association":                                      952,
		"aws_rum_app_monitor":                                              953,
		"aws_rum_metrics_destination":                                      954,
		"aws_s3_access_point":                                              955,
		"aws_s3_account_public_access_block":                               956,
		"aws_s3_bucket":                                                    957,
		"aws_s3_bucket_accelerate_configuration":                           958,
		"aws_s3_bucket_acl":                                                959,
		"aws_s3_bucket_analytics_configuration":                            960,
		"aws_s3_bucket_cors_configuration":                                 961,
		"aws_s3_bucket_intelligent_tiering_configuration":                  962,
		"aws_s3_bucket_inventory":                                          963,
		"aws_s3_bucket_lifecycle_configuration":                            964,
		"aws_s3_bucket_logging":                                            965,
		"aws_s3_bucket_metric":                                             966,
		"aws_s3_bucket_notification":                                       967,
		"aws_s3_bucket_object":                                             968,
		"aws_s3_bucket_object_lock_configuration":                          969,
		"aws_s3_bucket_ownership_controls":                                 970,
		"aws_s3_bucket_policy":                                             971,
		"aws_s3_bucket_public_access_block":                                972,
		"aws_s3_bucket_replication_configuration":                          973,
		"aws_s3_bucket_request_payment_configuration":                      974,
		"aws_s3_bucket_server_side_encryption_configuration":               975,
		"aws_s3_bucket_versioning":                                         976,
		"aws_s3_bucket_website_configuration":                              977,
		"aws_s3_object":                                                    978,
		"aws_s3_object_copy":                                               979,
		"aws_s3control_access_point_policy":                                980,
		"aws_s3control_bucket":                                             981,
		"aws_s3control_bucket_lifecycle_configuration":                     982,
		"aws_s3control_bucket_policy":                                      983,
		"aws_s3control_multi_region_access_point":                          984,
		"aws_s3control_multi_region_access_point_policy":                   985,
		"aws_s3control_object_lambda_access_point":                         986,
		"aws_s3control_object_lambda_access_point_policy":                  987,
		"aws_s3control_storage_lens_configuration":                         988,
		"aws_s3outposts_endpoint":                                          989,
		"aws_sagemaker_app":                                                990,
		"aws_sagemaker_app_image_config":                                   991,
		"aws_sagemaker_code_repository":                                    992,
		"aws_sagemaker_data_quality_job_definition":                        993,
		"aws_sagemaker_device":                                             994,
		"aws_sagemaker_device_fleet":                                       995,
		"aws_sagemaker_domain":                                             996,
		"aws_sagemaker_endpoint":                                           997,
		"aws_sagemaker_endpoint_configuration":                             998,
		"aws_sagemaker_feature_group":                                      999,
		"aws_sagemaker_flow_definition":                                    1000,
		"aws_sagemaker_human_task_ui":                                      1001,
		"aws_sagemaker_image":                                              1002,
		"aws_sagemaker_image_version":                                      1003,
		"aws_sagemaker_model":                                              1004,
		"aws_sagemaker_model_package_group":                                1005,
		"aws_sagemaker_model_package_group_policy":                         1006,
		"aws_sagemaker_monitoring_schedule":                                1007,
		"aws_sagemaker_notebook_instance":                                  1008,
		"aws_sagemaker_notebook_instance_lifecycle_configuration":          1009,
		"aws_sagemaker_pipeline":                                           1010,
		"aws_sagemaker_project":                                            1011,
		"aws_sagemaker_servicecatalog_portfolio_status":                    1012,
		"aws_sagemaker_space":                                              1013,
		"aws_sagemaker_studio_lifecycle_config":                            1014,
		"aws_sagemaker_user_profile":                                       1015,
		"aws_sagemaker_workforce":                                          1016,
		"aws_sagemaker_workteam":                                           1017,
		"aws_scheduler_schedule":                                           1018,
		"aws_scheduler_schedule_group":                                     1019,
		"aws_schemas_discoverer":                                           1020,
		"aws_schemas_registry":                                             1021,
		"aws_schemas_registry_policy":                                      1022,
		"aws_schemas_schema":                                               1023,
		"aws_secretsmanager_secret":                                        1024,
		"aws_secretsmanager_secret_policy":                                 1025,
		"aws_secretsmanager_secret_rotation":                               1026,
		"aws_secretsmanager_secret_version":                                1027,
		"aws_security_group":                                               1028,
		"aws_security_group_rule":                                          1029,
		"aws_securityhub_account":                                          1030,
		"aws_securityhub_action_target":                                    1031,
		"aws_securityhub_finding_aggregator":                               1032,
		"aws_securityhub_insight":                                          1033,
		"aws_securityhub_invite_accepter":                                  1034,
		"aws_securityhub_member":                                           1035,
		"aws_securityhub_organization_admin_account":                       1036,
		"aws_securityhub_organization_configuration":                       1037,
		"aws_securityhub_product_subscription":                             1038,
		"aws_securityhub_standards_control":                                1039,
		"aws_securityhub_standards_subscription":                           1040,
		"aws_serverlessapplicationrepository_cloudformation_stack":         1041,
		"aws_service_discovery_http_namespace":                             1042,
		"aws_service_discovery_instance":                                   1043,
		"aws_service_discovery_private_dns_namespace":                      1044,
		"aws_service_discovery_public_dns_namespace":                       1045,
		"aws_service_discovery_service":                                    1046,
		"aws_servicecatalog_budget_resource_association":                   1047,
		"aws_servicecatalog_constraint":                                    1048,
		"aws_servicecatalog_organizations_access":                          1049,
		"aws_servicecatalog_portfolio":                                     1050,
		"aws_servicecatalog_portfolio_share":                               1051,
		"aws_servicecatalog_principal_portfolio_association":               1052,
		"aws_servicecatalog_product":                                       1053,
		"aws_servicecatalog_product_portfolio_association":                 1054,
		"aws_servicecatalog_provisioned_product":                           1055,
		"aws_servicecatalog_provisioning_artifact":                         1056,
		"aws_servicecatalog_service_action":                                1057,
		"aws_servicecatalog_tag_option":                                    1058,
		"aws_servicecatalog_tag_option_resource_association":               1059,
		"aws_servicequotas_service_quota":                                  1060,
		"aws_ses_active_receipt_rule_set":                                  1061,
		"aws_ses_configuration_set":                                        1062,
		"aws_ses_domain_dkim":                                              1063,
		"aws_ses_domain_identity":                                          1064,
		"aws_ses_domain_identity_verification":                             1065,
		"aws_ses_domain_mail_from":                                         1066,
		"aws_ses_email_identity":                                           1067,
		"aws_ses_event_destination":                                        1068,
		"aws_ses_identity_notification_topic":                              1069,
		"aws_ses_identity_policy":                                          1070,
		"aws_ses_receipt_filter":                                           1071,
		"aws_ses_receipt_rule":                                             1072,
		"aws_ses_receipt_rule_set":                                         1073,
		"aws_ses_template":                                                 1074,
		"aws_sesv2_configuration_set":                                      1075,
		"aws_sesv2_configuration_set_event_destination":                    1076,
		"aws_sesv2_contact_list":                                           1077,
		"aws_sesv2_dedicated_ip_assignment":                                1078,
		"aws_sesv2_dedicated_ip_pool":                                      1079,
		"aws_sesv2_email_identity":                                         1080,
		"aws_sesv2_email_identity_feedback_attributes":                     1081,
		"aws_sesv2_email_identity_mail_from_attributes":                    1082,
		"aws_sfn_activity":                                                 1083,
		"aws_sfn_alias":                                                    1084,
		"aws_sfn_state_machine":                                            1085,
		"aws_shield_application_layer_automatic_response":                  1086,
		"aws_shield_drt_access_log_bucket_association":                     1087,
		"aws_shield_drt_access_role_arn_association":                       1088,
		"aws_shield_protection":                                            1089,
		"aws_shield_protection_group":                                      1090,
		"aws_shield_protection_health_check_association":                   1091,
		"aws_signer_signing_job":                                           1092,
		"aws_signer_signing_profile":                                       1093,
		"aws_signer_signing_profile_permission":                            1094,
		"aws_simpledb_domain":                                              1095,
		"aws_snapshot_create_volume_permission":                            1096,
		"aws_sns_platform_application":                                     1097,
		"aws_sns_sms_preferences":                                          1098,
		"aws_sns_topic":                                                    1099,
		"aws_sns_topic_data_protection_policy":                             1100,
		"aws_sns_topic_policy":                                             1101,
		"aws_sns_topic_subscription":                                       1102,
		"aws_spot_datafeed_subscription":                                   1103,
		"aws_spot_fleet_request":                                           1104,
		"aws_spot_instance_request":                                        1105,
		"aws_sqs_queue":                                                    1106,
		"aws_sqs_queue_policy":                                             1107,
		"aws_sqs_queue_redrive_allow_policy":                               1108,
		"aws_sqs_queue_redrive_policy":                                     1109,
		"aws_ssm_activation":                                               1110,
		"aws_ssm_association":                                              1111,
		"aws_ssm_default_patch_baseline":                                   1112,
		"aws_ssm_document":                                                 1113,
		"aws_ssm_maintenance_window":                                       1114,
		"aws_ssm_maintenance_window_target":                                1115,
		"aws_ssm_maintenance_window_task":                                  1116,
		"aws_ssm_parameter":                                                1117,
		"aws_ssm_patch_baseline":                                           1118,
		"aws_ssm_patch_group":                                              1119,
		"aws_ssm_resource_data_sync":                                       1120,
		"aws_ssm_service_setting":                                          1121,
		"aws_ssmcontacts_contact":                                          1122,
		"aws_ssmcontacts_contact_channel":                                  1123,
		"aws_ssmcontacts_plan":                                             1124,
		"aws_ssmincidents_replication_set":                                 1125,
		"aws_ssmincidents_response_plan":                                   1126,
		"aws_ssoadmin_account_assignment":                                  1127,
		"aws_ssoadmin_customer_managed_policy_attachment":                  1128,
		"aws_ssoadmin_instance_access_control_attributes":                  1129,
		"aws_ssoadmin_managed_policy_attachment":                           1130,
		"aws_ssoadmin_permission_set":                                      1131,
		"aws_ssoadmin_permission_set_inline_policy":                        1132,
		"aws_ssoadmin_permissions_boundary_attachment":                     1133,
		"aws_storagegateway_cache":                                         1134,
		"aws_storagegateway_cached_iscsi_volume":                           1135,
		"aws_storagegateway_file_system_association":                       1136,
		"aws_storagegateway_gateway":                                       1137,
		"aws_storagegateway_nfs_file_share":                                1138,
		"aws_storagegateway_smb_file_share":                                1139,
		"aws_storagegateway_stored_iscsi_volume":                           1140,
		"aws_storagegateway_tape_pool":                                     1141,
		"aws_storagegateway_upload_buffer":                                 1142,
		"aws_storagegateway_working_storage":                               1143,
		"aws_subnet":                                                       1144,
		"aws_swf_domain":                                                   1145,
		"aws_synthetics_canary":                                            1146,
		"aws_synthetics_group":                                             1147,
		"aws_synthetics_group_association":                                 1148,
		"aws_timestreamwrite_database":                                     1149,
		"aws_timestreamwrite_table":                                        1150,
		"aws_transcribe_language_model":                                    1151,
		"aws_transcribe_medical_vocabulary":                                1152,
		"aws_transcribe_vocabulary":                                        1153,
		"aws_transcribe_vocabulary_filter":                                 1154,
		"aws_transfer_access":                                              1155,
		"aws_transfer_agreement":                                           1156,
		"aws_transfer_certificate":                                         1157,
		"aws_transfer_connector":                                           1158,
		"aws_transfer_profile":                                             1159,
		"aws_transfer_server":                                              1160,
		"aws_transfer_ssh_key":                                             1161,
		"aws_transfer_tag":                                                 1162,
		"aws_transfer_user":                                                1163,
		"aws_transfer_workflow":                                            1164,
		"aws_verifiedaccess_instance":                                      1165,
		"aws_verifiedaccess_trust_provider":                                1166,
		"aws_volume_attachment":                                            1167,
		"aws_vpc":                                                          1168,
		"aws_vpc_dhcp_options":                                             1169,
		"aws_vpc_dhcp_options_association":                                 1170,
		"aws_vpc_endpoint":                                                 1171,
		"aws_vpc_endpoint_connection_accepter":                             1172,
		"aws_vpc_endpoint_connection_notification":                         1173,
		"aws_vpc_endpoint_policy":                                          1174,
		"aws_vpc_endpoint_route_table_association":                         1175,
		"aws_vpc_endpoint_security_group_association":                      1176,
		"aws_vpc_endpoint_service":                                         1177,
		"aws_vpc_endpoint_service_allowed_principal":                       1178,
		"aws_vpc_endpoint_subnet_association":                              1179,
		"aws_vpc_ipam":                                                     1180,
		"aws_vpc_ipam_organization_admin_account":                          1181,
		"aws_vpc_ipam_pool":                                                1182,
		"aws_vpc_ipam_pool_cidr":                                           1183,
		"aws_vpc_ipam_pool_cidr_allocation":                                1184,
		"aws_vpc_ipam_preview_next_cidr":                                   1185,
		"aws_vpc_ipam_resource_discovery":                                  1186,
		"aws_vpc_ipam_resource_discovery_association":                      1187,
		"aws_vpc_ipam_scope":                                               1188,
		"aws_vpc_ipv4_cidr_block_association":                              1189,
		"aws_vpc_ipv6_cidr_block_association":                              1190,
		"aws_vpc_network_performance_metric_subscription":                  1191,
		"aws_vpc_peering_connection":                                       1192,
		"aws_vpc_peering_connection_accepter":                              1193,
		"aws_vpc_peering_connection_options":                               1194,
		"aws_vpc_security_group_egress_rule":                               1195,
		"aws_vpc_security_group_ingress_rule":                              1196,
		"aws_vpclattice_access_log_subscription":                           1197,
		"aws_vpclattice_auth_policy":                                       1198,
		"aws_vpclattice_listener":                                          1199,
		"aws_vpclattice_listener_rule":                                     1200,
		"aws_vpclattice_resource_policy":                                   1201,
		"aws_vpclattice_service":                                           1202,
		"aws_vpclattice_service_network":                                   1203,
		"aws_vpclattice_service_network_service_association":               1204,
		"aws_vpclattice_service_network_vpc_association":                   1205,
		"aws_vpclattice_target_group":                                      1206,
		"aws_vpclattice_target_group_attachment":                           1207,
		"aws_vpn_connection":                                               1208,
		"aws_vpn_connection_route":                                         1209,
		"aws_vpn_gateway":                                                  1210,
		"aws_vpn_gateway_attachment":                                       1211,
		"aws_vpn_gateway_route_propagation":                                1212,
		"aws_waf_byte_match_set":                                           1213,
		"aws_waf_geo_match_set":                                            1214,
		"aws_waf_ipset":                                                    1215,
		"aws_waf_rate_based_rule":                                          1216,
		"aws_waf_regex_match_set":                                          1217,
		"aws_waf_regex_pattern_set":                                        1218,
		"aws_waf_rule":                                                     1219,
		"aws_waf_rule_group":                                               1220,
		"aws_waf_size_constraint_set":                                      1221,
		"aws_waf_sql_injection_match_set":                                  1222,
		"aws_waf_web_acl":                                                  1223,
		"aws_waf_xss_match_set":                                            1224,
		"aws_wafregional_byte_match_set":                                   1225,
		"aws_wafregional_geo_match_set":                                    1226,
		"aws_wafregional_ipset":                                            1227,
		"aws_wafregional_rate_based_rule":                                  1228,
		"aws_wafregional_regex_match_set":                                  1229,
		"aws_wafregional_regex_pattern_set":                                1230,
		"aws_wafregional_rule":                                             1231,
		"aws_wafregional_rule_group":                                       1232,
		"aws_wafregional_size_constraint_set":                              1233,
		"aws_wafregional_sql_injection_match_set":                          1234,
		"aws_wafregional_web_acl":                                          1235,
		"aws_wafregional_web_acl_association":                              1236,
		"aws_wafregional_xss_match_set":                                    1237,
		"aws_wafv2_ip_set":                                                 1238,
		"aws_wafv2_regex_pattern_set":                                      1239,
		"aws_wafv2_rule_group":                                             1240,
		"aws_wafv2_web_acl":                                                1241,
		"aws_wafv2_web_acl_association":                                    1242,
		"aws_wafv2_web_acl_logging_configuration":                          1243,
		"aws_worklink_fleet":                                               1244,
		"aws_worklink_website_certificate_authority_association":           1245,
		"aws_workspaces_connection_alias":                                  1246,
		"aws_workspaces_directory":                                         1247,
		"aws_workspaces_ip_group":                                          1248,
		"aws_workspaces_workspace":                                         1249,
		"aws_xray_encryption_config":                                       1250,
		"aws_xray_group":                                                   1251,
		"aws_xray_sampling_rule":                                           1252,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
