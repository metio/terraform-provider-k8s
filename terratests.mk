# SPDX-FileCopyrightText: The terraform-provider-k8s Authors
# SPDX-License-Identifier: 0BSD

out/terratest-sentinel-about_k8s_io_cluster_property_v1alpha1_manifest_test.go: out/install-sentinel terratest/about_k8s_io_v1alpha1/about_k8s_io_cluster_property_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_about_k8s_io_cluster_property_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/about_k8s_io_v1alpha1/about_k8s_io_cluster_property_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-acid_zalan_do_operator_configuration_v1_manifest_test.go: out/install-sentinel terratest/acid_zalan_do_v1/acid_zalan_do_operator_configuration_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_acid_zalan_do_operator_configuration_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/acid_zalan_do_v1/acid_zalan_do_operator_configuration_v1_manifest_test.go
	touch $@
out/terratest-sentinel-acid_zalan_do_postgres_team_v1_manifest_test.go: out/install-sentinel terratest/acid_zalan_do_v1/acid_zalan_do_postgres_team_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_acid_zalan_do_postgres_team_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/acid_zalan_do_v1/acid_zalan_do_postgres_team_v1_manifest_test.go
	touch $@
out/terratest-sentinel-acid_zalan_do_postgresql_v1_manifest_test.go: out/install-sentinel terratest/acid_zalan_do_v1/acid_zalan_do_postgresql_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_acid_zalan_do_postgresql_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/acid_zalan_do_v1/acid_zalan_do_postgresql_v1_manifest_test.go
	touch $@
out/terratest-sentinel-acme_cert_manager_io_challenge_v1_manifest_test.go: out/install-sentinel terratest/acme_cert_manager_io_v1/acme_cert_manager_io_challenge_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_acme_cert_manager_io_challenge_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/acme_cert_manager_io_v1/acme_cert_manager_io_challenge_v1_manifest_test.go
	touch $@
out/terratest-sentinel-acme_cert_manager_io_order_v1_manifest_test.go: out/install-sentinel terratest/acme_cert_manager_io_v1/acme_cert_manager_io_order_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_acme_cert_manager_io_order_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/acme_cert_manager_io_v1/acme_cert_manager_io_order_v1_manifest_test.go
	touch $@
out/terratest-sentinel-acmpca_services_k8s_aws_certificate_authority_activation_v1alpha1_manifest_test.go: out/install-sentinel terratest/acmpca_services_k8s_aws_v1alpha1/acmpca_services_k8s_aws_certificate_authority_activation_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_acmpca_services_k8s_aws_certificate_authority_activation_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/acmpca_services_k8s_aws_v1alpha1/acmpca_services_k8s_aws_certificate_authority_activation_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-acmpca_services_k8s_aws_certificate_authority_v1alpha1_manifest_test.go: out/install-sentinel terratest/acmpca_services_k8s_aws_v1alpha1/acmpca_services_k8s_aws_certificate_authority_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_acmpca_services_k8s_aws_certificate_authority_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/acmpca_services_k8s_aws_v1alpha1/acmpca_services_k8s_aws_certificate_authority_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-acmpca_services_k8s_aws_certificate_v1alpha1_manifest_test.go: out/install-sentinel terratest/acmpca_services_k8s_aws_v1alpha1/acmpca_services_k8s_aws_certificate_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_acmpca_services_k8s_aws_certificate_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/acmpca_services_k8s_aws_v1alpha1/acmpca_services_k8s_aws_certificate_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-actions_github_com_autoscaling_listener_v1alpha1_manifest_test.go: out/install-sentinel terratest/actions_github_com_v1alpha1/actions_github_com_autoscaling_listener_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_actions_github_com_autoscaling_listener_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/actions_github_com_v1alpha1/actions_github_com_autoscaling_listener_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-actions_github_com_autoscaling_runner_set_v1alpha1_manifest_test.go: out/install-sentinel terratest/actions_github_com_v1alpha1/actions_github_com_autoscaling_runner_set_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_actions_github_com_autoscaling_runner_set_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/actions_github_com_v1alpha1/actions_github_com_autoscaling_runner_set_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-actions_github_com_ephemeral_runner_set_v1alpha1_manifest_test.go: out/install-sentinel terratest/actions_github_com_v1alpha1/actions_github_com_ephemeral_runner_set_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_actions_github_com_ephemeral_runner_set_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/actions_github_com_v1alpha1/actions_github_com_ephemeral_runner_set_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-actions_github_com_ephemeral_runner_v1alpha1_manifest_test.go: out/install-sentinel terratest/actions_github_com_v1alpha1/actions_github_com_ephemeral_runner_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_actions_github_com_ephemeral_runner_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/actions_github_com_v1alpha1/actions_github_com_ephemeral_runner_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-actions_summerwind_dev_horizontal_runner_autoscaler_v1alpha1_manifest_test.go: out/install-sentinel terratest/actions_summerwind_dev_v1alpha1/actions_summerwind_dev_horizontal_runner_autoscaler_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_actions_summerwind_dev_horizontal_runner_autoscaler_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/actions_summerwind_dev_v1alpha1/actions_summerwind_dev_horizontal_runner_autoscaler_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-actions_summerwind_dev_runner_deployment_v1alpha1_manifest_test.go: out/install-sentinel terratest/actions_summerwind_dev_v1alpha1/actions_summerwind_dev_runner_deployment_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_actions_summerwind_dev_runner_deployment_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/actions_summerwind_dev_v1alpha1/actions_summerwind_dev_runner_deployment_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-actions_summerwind_dev_runner_replica_set_v1alpha1_manifest_test.go: out/install-sentinel terratest/actions_summerwind_dev_v1alpha1/actions_summerwind_dev_runner_replica_set_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_actions_summerwind_dev_runner_replica_set_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/actions_summerwind_dev_v1alpha1/actions_summerwind_dev_runner_replica_set_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-actions_summerwind_dev_runner_set_v1alpha1_manifest_test.go: out/install-sentinel terratest/actions_summerwind_dev_v1alpha1/actions_summerwind_dev_runner_set_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_actions_summerwind_dev_runner_set_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/actions_summerwind_dev_v1alpha1/actions_summerwind_dev_runner_set_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-actions_summerwind_dev_runner_v1alpha1_manifest_test.go: out/install-sentinel terratest/actions_summerwind_dev_v1alpha1/actions_summerwind_dev_runner_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_actions_summerwind_dev_runner_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/actions_summerwind_dev_v1alpha1/actions_summerwind_dev_runner_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-addons_cluster_x_k8s_io_cluster_resource_set_binding_v1alpha3_manifest_test.go: out/install-sentinel terratest/addons_cluster_x_k8s_io_v1alpha3/addons_cluster_x_k8s_io_cluster_resource_set_binding_v1alpha3_manifest_test.go $(shell find ./examples/data-sources/k8s_addons_cluster_x_k8s_io_cluster_resource_set_binding_v1alpha3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/addons_cluster_x_k8s_io_v1alpha3/addons_cluster_x_k8s_io_cluster_resource_set_binding_v1alpha3_manifest_test.go
	touch $@
out/terratest-sentinel-addons_cluster_x_k8s_io_cluster_resource_set_v1alpha3_manifest_test.go: out/install-sentinel terratest/addons_cluster_x_k8s_io_v1alpha3/addons_cluster_x_k8s_io_cluster_resource_set_v1alpha3_manifest_test.go $(shell find ./examples/data-sources/k8s_addons_cluster_x_k8s_io_cluster_resource_set_v1alpha3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/addons_cluster_x_k8s_io_v1alpha3/addons_cluster_x_k8s_io_cluster_resource_set_v1alpha3_manifest_test.go
	touch $@
out/terratest-sentinel-addons_cluster_x_k8s_io_cluster_resource_set_binding_v1alpha4_manifest_test.go: out/install-sentinel terratest/addons_cluster_x_k8s_io_v1alpha4/addons_cluster_x_k8s_io_cluster_resource_set_binding_v1alpha4_manifest_test.go $(shell find ./examples/data-sources/k8s_addons_cluster_x_k8s_io_cluster_resource_set_binding_v1alpha4_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/addons_cluster_x_k8s_io_v1alpha4/addons_cluster_x_k8s_io_cluster_resource_set_binding_v1alpha4_manifest_test.go
	touch $@
out/terratest-sentinel-addons_cluster_x_k8s_io_cluster_resource_set_v1alpha4_manifest_test.go: out/install-sentinel terratest/addons_cluster_x_k8s_io_v1alpha4/addons_cluster_x_k8s_io_cluster_resource_set_v1alpha4_manifest_test.go $(shell find ./examples/data-sources/k8s_addons_cluster_x_k8s_io_cluster_resource_set_v1alpha4_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/addons_cluster_x_k8s_io_v1alpha4/addons_cluster_x_k8s_io_cluster_resource_set_v1alpha4_manifest_test.go
	touch $@
out/terratest-sentinel-addons_cluster_x_k8s_io_cluster_resource_set_binding_v1beta1_manifest_test.go: out/install-sentinel terratest/addons_cluster_x_k8s_io_v1beta1/addons_cluster_x_k8s_io_cluster_resource_set_binding_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_addons_cluster_x_k8s_io_cluster_resource_set_binding_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/addons_cluster_x_k8s_io_v1beta1/addons_cluster_x_k8s_io_cluster_resource_set_binding_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-addons_cluster_x_k8s_io_cluster_resource_set_v1beta1_manifest_test.go: out/install-sentinel terratest/addons_cluster_x_k8s_io_v1beta1/addons_cluster_x_k8s_io_cluster_resource_set_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_addons_cluster_x_k8s_io_cluster_resource_set_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/addons_cluster_x_k8s_io_v1beta1/addons_cluster_x_k8s_io_cluster_resource_set_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-admissionregistration_k8s_io_mutating_webhook_configuration_v1_manifest_test.go: out/install-sentinel terratest/admissionregistration_k8s_io_v1/admissionregistration_k8s_io_mutating_webhook_configuration_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_admissionregistration_k8s_io_mutating_webhook_configuration_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/admissionregistration_k8s_io_v1/admissionregistration_k8s_io_mutating_webhook_configuration_v1_manifest_test.go
	touch $@
out/terratest-sentinel-admissionregistration_k8s_io_validating_webhook_configuration_v1_manifest_test.go: out/install-sentinel terratest/admissionregistration_k8s_io_v1/admissionregistration_k8s_io_validating_webhook_configuration_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_admissionregistration_k8s_io_validating_webhook_configuration_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/admissionregistration_k8s_io_v1/admissionregistration_k8s_io_validating_webhook_configuration_v1_manifest_test.go
	touch $@
out/terratest-sentinel-agent_k8s_elastic_co_agent_v1alpha1_manifest_test.go: out/install-sentinel terratest/agent_k8s_elastic_co_v1alpha1/agent_k8s_elastic_co_agent_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_agent_k8s_elastic_co_agent_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/agent_k8s_elastic_co_v1alpha1/agent_k8s_elastic_co_agent_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-anywhere_eks_amazonaws_com_aws_datacenter_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_aws_datacenter_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_anywhere_eks_amazonaws_com_aws_datacenter_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_aws_datacenter_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-anywhere_eks_amazonaws_com_aws_iam_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_aws_iam_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_anywhere_eks_amazonaws_com_aws_iam_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_aws_iam_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-anywhere_eks_amazonaws_com_bundles_v1alpha1_manifest_test.go: out/install-sentinel terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_bundles_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_anywhere_eks_amazonaws_com_bundles_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_bundles_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-anywhere_eks_amazonaws_com_cloud_stack_datacenter_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_cloud_stack_datacenter_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_anywhere_eks_amazonaws_com_cloud_stack_datacenter_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_cloud_stack_datacenter_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-anywhere_eks_amazonaws_com_cloud_stack_machine_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_cloud_stack_machine_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_anywhere_eks_amazonaws_com_cloud_stack_machine_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_cloud_stack_machine_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-anywhere_eks_amazonaws_com_cluster_v1alpha1_manifest_test.go: out/install-sentinel terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_cluster_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_anywhere_eks_amazonaws_com_cluster_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_cluster_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-anywhere_eks_amazonaws_com_control_plane_upgrade_v1alpha1_manifest_test.go: out/install-sentinel terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_control_plane_upgrade_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_anywhere_eks_amazonaws_com_control_plane_upgrade_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_control_plane_upgrade_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-anywhere_eks_amazonaws_com_docker_datacenter_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_docker_datacenter_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_anywhere_eks_amazonaws_com_docker_datacenter_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_docker_datacenter_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-anywhere_eks_amazonaws_com_eksa_release_v1alpha1_manifest_test.go: out/install-sentinel terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_eksa_release_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_anywhere_eks_amazonaws_com_eksa_release_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_eksa_release_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-anywhere_eks_amazonaws_com_flux_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_flux_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_anywhere_eks_amazonaws_com_flux_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_flux_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-anywhere_eks_amazonaws_com_git_ops_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_git_ops_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_anywhere_eks_amazonaws_com_git_ops_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_git_ops_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-anywhere_eks_amazonaws_com_machine_deployment_upgrade_v1alpha1_manifest_test.go: out/install-sentinel terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_machine_deployment_upgrade_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_anywhere_eks_amazonaws_com_machine_deployment_upgrade_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_machine_deployment_upgrade_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-anywhere_eks_amazonaws_com_node_upgrade_v1alpha1_manifest_test.go: out/install-sentinel terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_node_upgrade_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_anywhere_eks_amazonaws_com_node_upgrade_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_node_upgrade_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-anywhere_eks_amazonaws_com_nutanix_datacenter_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_nutanix_datacenter_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_anywhere_eks_amazonaws_com_nutanix_datacenter_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_nutanix_datacenter_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-anywhere_eks_amazonaws_com_nutanix_machine_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_nutanix_machine_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_anywhere_eks_amazonaws_com_nutanix_machine_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_nutanix_machine_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-anywhere_eks_amazonaws_com_oidc_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_oidc_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_anywhere_eks_amazonaws_com_oidc_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_oidc_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-anywhere_eks_amazonaws_com_snow_datacenter_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_snow_datacenter_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_anywhere_eks_amazonaws_com_snow_datacenter_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_snow_datacenter_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-anywhere_eks_amazonaws_com_snow_ip_pool_v1alpha1_manifest_test.go: out/install-sentinel terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_snow_ip_pool_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_anywhere_eks_amazonaws_com_snow_ip_pool_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_snow_ip_pool_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-anywhere_eks_amazonaws_com_snow_machine_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_snow_machine_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_anywhere_eks_amazonaws_com_snow_machine_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_snow_machine_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-anywhere_eks_amazonaws_com_tinkerbell_datacenter_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_tinkerbell_datacenter_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_anywhere_eks_amazonaws_com_tinkerbell_datacenter_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_tinkerbell_datacenter_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-anywhere_eks_amazonaws_com_tinkerbell_machine_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_tinkerbell_machine_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_anywhere_eks_amazonaws_com_tinkerbell_machine_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_tinkerbell_machine_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-anywhere_eks_amazonaws_com_tinkerbell_template_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_tinkerbell_template_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_anywhere_eks_amazonaws_com_tinkerbell_template_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_tinkerbell_template_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-anywhere_eks_amazonaws_com_v_sphere_datacenter_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_v_sphere_datacenter_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_anywhere_eks_amazonaws_com_v_sphere_datacenter_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_v_sphere_datacenter_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-anywhere_eks_amazonaws_com_v_sphere_machine_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_v_sphere_machine_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_anywhere_eks_amazonaws_com_v_sphere_machine_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/anywhere_eks_amazonaws_com_v1alpha1/anywhere_eks_amazonaws_com_v_sphere_machine_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apacheweb_arsenal_dev_apacheweb_v1alpha1_manifest_test.go: out/install-sentinel terratest/apacheweb_arsenal_dev_v1alpha1/apacheweb_arsenal_dev_apacheweb_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apacheweb_arsenal_dev_apacheweb_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apacheweb_arsenal_dev_v1alpha1/apacheweb_arsenal_dev_apacheweb_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-api_clever_cloud_com_config_provider_v1_manifest_test.go: out/install-sentinel terratest/api_clever_cloud_com_v1/api_clever_cloud_com_config_provider_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_api_clever_cloud_com_config_provider_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/api_clever_cloud_com_v1/api_clever_cloud_com_config_provider_v1_manifest_test.go
	touch $@
out/terratest-sentinel-api_clever_cloud_com_elastic_search_v1_manifest_test.go: out/install-sentinel terratest/api_clever_cloud_com_v1/api_clever_cloud_com_elastic_search_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_api_clever_cloud_com_elastic_search_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/api_clever_cloud_com_v1/api_clever_cloud_com_elastic_search_v1_manifest_test.go
	touch $@
out/terratest-sentinel-api_clever_cloud_com_mongo_db_v1_manifest_test.go: out/install-sentinel terratest/api_clever_cloud_com_v1/api_clever_cloud_com_mongo_db_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_api_clever_cloud_com_mongo_db_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/api_clever_cloud_com_v1/api_clever_cloud_com_mongo_db_v1_manifest_test.go
	touch $@
out/terratest-sentinel-api_clever_cloud_com_my_sql_v1_manifest_test.go: out/install-sentinel terratest/api_clever_cloud_com_v1/api_clever_cloud_com_my_sql_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_api_clever_cloud_com_my_sql_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/api_clever_cloud_com_v1/api_clever_cloud_com_my_sql_v1_manifest_test.go
	touch $@
out/terratest-sentinel-api_clever_cloud_com_postgre_sql_v1_manifest_test.go: out/install-sentinel terratest/api_clever_cloud_com_v1/api_clever_cloud_com_postgre_sql_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_api_clever_cloud_com_postgre_sql_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/api_clever_cloud_com_v1/api_clever_cloud_com_postgre_sql_v1_manifest_test.go
	touch $@
out/terratest-sentinel-api_clever_cloud_com_redis_v1_manifest_test.go: out/install-sentinel terratest/api_clever_cloud_com_v1/api_clever_cloud_com_redis_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_api_clever_cloud_com_redis_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/api_clever_cloud_com_v1/api_clever_cloud_com_redis_v1_manifest_test.go
	touch $@
out/terratest-sentinel-api_clever_cloud_com_pulsar_v1beta1_manifest_test.go: out/install-sentinel terratest/api_clever_cloud_com_v1beta1/api_clever_cloud_com_pulsar_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_api_clever_cloud_com_pulsar_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/api_clever_cloud_com_v1beta1/api_clever_cloud_com_pulsar_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-api_kubemod_io_mod_rule_v1beta1_manifest_test.go: out/install-sentinel terratest/api_kubemod_io_v1beta1/api_kubemod_io_mod_rule_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_api_kubemod_io_mod_rule_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/api_kubemod_io_v1beta1/api_kubemod_io_mod_rule_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-apicodegen_apimatic_io_api_matic_v1beta1_manifest_test.go: out/install-sentinel terratest/apicodegen_apimatic_io_v1beta1/apicodegen_apimatic_io_api_matic_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_apicodegen_apimatic_io_api_matic_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apicodegen_apimatic_io_v1beta1/apicodegen_apimatic_io_api_matic_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-apiextensions_crossplane_io_composite_resource_definition_v1_manifest_test.go: out/install-sentinel terratest/apiextensions_crossplane_io_v1/apiextensions_crossplane_io_composite_resource_definition_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_apiextensions_crossplane_io_composite_resource_definition_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apiextensions_crossplane_io_v1/apiextensions_crossplane_io_composite_resource_definition_v1_manifest_test.go
	touch $@
out/terratest-sentinel-apiextensions_crossplane_io_composition_revision_v1_manifest_test.go: out/install-sentinel terratest/apiextensions_crossplane_io_v1/apiextensions_crossplane_io_composition_revision_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_apiextensions_crossplane_io_composition_revision_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apiextensions_crossplane_io_v1/apiextensions_crossplane_io_composition_revision_v1_manifest_test.go
	touch $@
out/terratest-sentinel-apiextensions_crossplane_io_composition_v1_manifest_test.go: out/install-sentinel terratest/apiextensions_crossplane_io_v1/apiextensions_crossplane_io_composition_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_apiextensions_crossplane_io_composition_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apiextensions_crossplane_io_v1/apiextensions_crossplane_io_composition_v1_manifest_test.go
	touch $@
out/terratest-sentinel-apiextensions_crossplane_io_composition_revision_v1beta1_manifest_test.go: out/install-sentinel terratest/apiextensions_crossplane_io_v1beta1/apiextensions_crossplane_io_composition_revision_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_apiextensions_crossplane_io_composition_revision_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apiextensions_crossplane_io_v1beta1/apiextensions_crossplane_io_composition_revision_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-apigatewayv2_services_k8s_aws_api_v1alpha1_manifest_test.go: out/install-sentinel terratest/apigatewayv2_services_k8s_aws_v1alpha1/apigatewayv2_services_k8s_aws_api_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apigatewayv2_services_k8s_aws_api_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apigatewayv2_services_k8s_aws_v1alpha1/apigatewayv2_services_k8s_aws_api_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apigatewayv2_services_k8s_aws_authorizer_v1alpha1_manifest_test.go: out/install-sentinel terratest/apigatewayv2_services_k8s_aws_v1alpha1/apigatewayv2_services_k8s_aws_authorizer_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apigatewayv2_services_k8s_aws_authorizer_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apigatewayv2_services_k8s_aws_v1alpha1/apigatewayv2_services_k8s_aws_authorizer_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apigatewayv2_services_k8s_aws_deployment_v1alpha1_manifest_test.go: out/install-sentinel terratest/apigatewayv2_services_k8s_aws_v1alpha1/apigatewayv2_services_k8s_aws_deployment_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apigatewayv2_services_k8s_aws_deployment_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apigatewayv2_services_k8s_aws_v1alpha1/apigatewayv2_services_k8s_aws_deployment_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apigatewayv2_services_k8s_aws_integration_v1alpha1_manifest_test.go: out/install-sentinel terratest/apigatewayv2_services_k8s_aws_v1alpha1/apigatewayv2_services_k8s_aws_integration_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apigatewayv2_services_k8s_aws_integration_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apigatewayv2_services_k8s_aws_v1alpha1/apigatewayv2_services_k8s_aws_integration_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apigatewayv2_services_k8s_aws_route_v1alpha1_manifest_test.go: out/install-sentinel terratest/apigatewayv2_services_k8s_aws_v1alpha1/apigatewayv2_services_k8s_aws_route_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apigatewayv2_services_k8s_aws_route_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apigatewayv2_services_k8s_aws_v1alpha1/apigatewayv2_services_k8s_aws_route_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apigatewayv2_services_k8s_aws_stage_v1alpha1_manifest_test.go: out/install-sentinel terratest/apigatewayv2_services_k8s_aws_v1alpha1/apigatewayv2_services_k8s_aws_stage_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apigatewayv2_services_k8s_aws_stage_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apigatewayv2_services_k8s_aws_v1alpha1/apigatewayv2_services_k8s_aws_stage_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apigatewayv2_services_k8s_aws_vpc_link_v1alpha1_manifest_test.go: out/install-sentinel terratest/apigatewayv2_services_k8s_aws_v1alpha1/apigatewayv2_services_k8s_aws_vpc_link_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apigatewayv2_services_k8s_aws_vpc_link_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apigatewayv2_services_k8s_aws_v1alpha1/apigatewayv2_services_k8s_aws_vpc_link_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apiregistration_k8s_io_api_service_v1_manifest_test.go: out/install-sentinel terratest/apiregistration_k8s_io_v1/apiregistration_k8s_io_api_service_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_apiregistration_k8s_io_api_service_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apiregistration_k8s_io_v1/apiregistration_k8s_io_api_service_v1_manifest_test.go
	touch $@
out/terratest-sentinel-apisix_apache_org_apisix_cluster_config_v2_manifest_test.go: out/install-sentinel terratest/apisix_apache_org_v2/apisix_apache_org_apisix_cluster_config_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_apisix_apache_org_apisix_cluster_config_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apisix_apache_org_v2/apisix_apache_org_apisix_cluster_config_v2_manifest_test.go
	touch $@
out/terratest-sentinel-apisix_apache_org_apisix_consumer_v2_manifest_test.go: out/install-sentinel terratest/apisix_apache_org_v2/apisix_apache_org_apisix_consumer_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_apisix_apache_org_apisix_consumer_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apisix_apache_org_v2/apisix_apache_org_apisix_consumer_v2_manifest_test.go
	touch $@
out/terratest-sentinel-apisix_apache_org_apisix_global_rule_v2_manifest_test.go: out/install-sentinel terratest/apisix_apache_org_v2/apisix_apache_org_apisix_global_rule_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_apisix_apache_org_apisix_global_rule_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apisix_apache_org_v2/apisix_apache_org_apisix_global_rule_v2_manifest_test.go
	touch $@
out/terratest-sentinel-apisix_apache_org_apisix_plugin_config_v2_manifest_test.go: out/install-sentinel terratest/apisix_apache_org_v2/apisix_apache_org_apisix_plugin_config_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_apisix_apache_org_apisix_plugin_config_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apisix_apache_org_v2/apisix_apache_org_apisix_plugin_config_v2_manifest_test.go
	touch $@
out/terratest-sentinel-apisix_apache_org_apisix_route_v2_manifest_test.go: out/install-sentinel terratest/apisix_apache_org_v2/apisix_apache_org_apisix_route_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_apisix_apache_org_apisix_route_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apisix_apache_org_v2/apisix_apache_org_apisix_route_v2_manifest_test.go
	touch $@
out/terratest-sentinel-apisix_apache_org_apisix_tls_v2_manifest_test.go: out/install-sentinel terratest/apisix_apache_org_v2/apisix_apache_org_apisix_tls_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_apisix_apache_org_apisix_tls_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apisix_apache_org_v2/apisix_apache_org_apisix_tls_v2_manifest_test.go
	touch $@
out/terratest-sentinel-apisix_apache_org_apisix_upstream_v2_manifest_test.go: out/install-sentinel terratest/apisix_apache_org_v2/apisix_apache_org_apisix_upstream_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_apisix_apache_org_apisix_upstream_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apisix_apache_org_v2/apisix_apache_org_apisix_upstream_v2_manifest_test.go
	touch $@
out/terratest-sentinel-apm_k8s_elastic_co_apm_server_v1_manifest_test.go: out/install-sentinel terratest/apm_k8s_elastic_co_v1/apm_k8s_elastic_co_apm_server_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_apm_k8s_elastic_co_apm_server_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apm_k8s_elastic_co_v1/apm_k8s_elastic_co_apm_server_v1_manifest_test.go
	touch $@
out/terratest-sentinel-apm_k8s_elastic_co_apm_server_v1beta1_manifest_test.go: out/install-sentinel terratest/apm_k8s_elastic_co_v1beta1/apm_k8s_elastic_co_apm_server_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_apm_k8s_elastic_co_apm_server_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apm_k8s_elastic_co_v1beta1/apm_k8s_elastic_co_apm_server_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-app_kiegroup_org_kogito_build_v1beta1_manifest_test.go: out/install-sentinel terratest/app_kiegroup_org_v1beta1/app_kiegroup_org_kogito_build_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_app_kiegroup_org_kogito_build_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/app_kiegroup_org_v1beta1/app_kiegroup_org_kogito_build_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-app_kiegroup_org_kogito_infra_v1beta1_manifest_test.go: out/install-sentinel terratest/app_kiegroup_org_v1beta1/app_kiegroup_org_kogito_infra_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_app_kiegroup_org_kogito_infra_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/app_kiegroup_org_v1beta1/app_kiegroup_org_kogito_infra_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-app_kiegroup_org_kogito_runtime_v1beta1_manifest_test.go: out/install-sentinel terratest/app_kiegroup_org_v1beta1/app_kiegroup_org_kogito_runtime_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_app_kiegroup_org_kogito_runtime_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/app_kiegroup_org_v1beta1/app_kiegroup_org_kogito_runtime_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-app_kiegroup_org_kogito_supporting_service_v1beta1_manifest_test.go: out/install-sentinel terratest/app_kiegroup_org_v1beta1/app_kiegroup_org_kogito_supporting_service_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_app_kiegroup_org_kogito_supporting_service_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/app_kiegroup_org_v1beta1/app_kiegroup_org_kogito_supporting_service_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-app_lightbend_com_akka_cluster_v1alpha1_manifest_test.go: out/install-sentinel terratest/app_lightbend_com_v1alpha1/app_lightbend_com_akka_cluster_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_app_lightbend_com_akka_cluster_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/app_lightbend_com_v1alpha1/app_lightbend_com_akka_cluster_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-app_redislabs_com_redis_enterprise_cluster_v1_manifest_test.go: out/install-sentinel terratest/app_redislabs_com_v1/app_redislabs_com_redis_enterprise_cluster_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_app_redislabs_com_redis_enterprise_cluster_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/app_redislabs_com_v1/app_redislabs_com_redis_enterprise_cluster_v1_manifest_test.go
	touch $@
out/terratest-sentinel-app_redislabs_com_redis_enterprise_active_active_database_v1alpha1_manifest_test.go: out/install-sentinel terratest/app_redislabs_com_v1alpha1/app_redislabs_com_redis_enterprise_active_active_database_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_app_redislabs_com_redis_enterprise_active_active_database_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/app_redislabs_com_v1alpha1/app_redislabs_com_redis_enterprise_active_active_database_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-app_redislabs_com_redis_enterprise_cluster_v1alpha1_manifest_test.go: out/install-sentinel terratest/app_redislabs_com_v1alpha1/app_redislabs_com_redis_enterprise_cluster_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_app_redislabs_com_redis_enterprise_cluster_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/app_redislabs_com_v1alpha1/app_redislabs_com_redis_enterprise_cluster_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-app_redislabs_com_redis_enterprise_database_v1alpha1_manifest_test.go: out/install-sentinel terratest/app_redislabs_com_v1alpha1/app_redislabs_com_redis_enterprise_database_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_app_redislabs_com_redis_enterprise_database_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/app_redislabs_com_v1alpha1/app_redislabs_com_redis_enterprise_database_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-app_redislabs_com_redis_enterprise_remote_cluster_v1alpha1_manifest_test.go: out/install-sentinel terratest/app_redislabs_com_v1alpha1/app_redislabs_com_redis_enterprise_remote_cluster_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_app_redislabs_com_redis_enterprise_remote_cluster_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/app_redislabs_com_v1alpha1/app_redislabs_com_redis_enterprise_remote_cluster_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-app_terraform_io_agent_pool_v1alpha2_manifest_test.go: out/install-sentinel terratest/app_terraform_io_v1alpha2/app_terraform_io_agent_pool_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_app_terraform_io_agent_pool_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/app_terraform_io_v1alpha2/app_terraform_io_agent_pool_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-app_terraform_io_module_v1alpha2_manifest_test.go: out/install-sentinel terratest/app_terraform_io_v1alpha2/app_terraform_io_module_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_app_terraform_io_module_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/app_terraform_io_v1alpha2/app_terraform_io_module_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-app_terraform_io_workspace_v1alpha2_manifest_test.go: out/install-sentinel terratest/app_terraform_io_v1alpha2/app_terraform_io_workspace_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_app_terraform_io_workspace_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/app_terraform_io_v1alpha2/app_terraform_io_workspace_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-application_networking_k8s_aws_access_log_policy_v1alpha1_manifest_test.go: out/install-sentinel terratest/application_networking_k8s_aws_v1alpha1/application_networking_k8s_aws_access_log_policy_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_application_networking_k8s_aws_access_log_policy_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/application_networking_k8s_aws_v1alpha1/application_networking_k8s_aws_access_log_policy_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-application_networking_k8s_aws_iam_auth_policy_v1alpha1_manifest_test.go: out/install-sentinel terratest/application_networking_k8s_aws_v1alpha1/application_networking_k8s_aws_iam_auth_policy_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_application_networking_k8s_aws_iam_auth_policy_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/application_networking_k8s_aws_v1alpha1/application_networking_k8s_aws_iam_auth_policy_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-application_networking_k8s_aws_service_import_v1alpha1_manifest_test.go: out/install-sentinel terratest/application_networking_k8s_aws_v1alpha1/application_networking_k8s_aws_service_import_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_application_networking_k8s_aws_service_import_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/application_networking_k8s_aws_v1alpha1/application_networking_k8s_aws_service_import_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-application_networking_k8s_aws_target_group_policy_v1alpha1_manifest_test.go: out/install-sentinel terratest/application_networking_k8s_aws_v1alpha1/application_networking_k8s_aws_target_group_policy_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_application_networking_k8s_aws_target_group_policy_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/application_networking_k8s_aws_v1alpha1/application_networking_k8s_aws_target_group_policy_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-application_networking_k8s_aws_vpc_association_policy_v1alpha1_manifest_test.go: out/install-sentinel terratest/application_networking_k8s_aws_v1alpha1/application_networking_k8s_aws_vpc_association_policy_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_application_networking_k8s_aws_vpc_association_policy_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/application_networking_k8s_aws_v1alpha1/application_networking_k8s_aws_vpc_association_policy_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-applicationautoscaling_services_k8s_aws_scalable_target_v1alpha1_manifest_test.go: out/install-sentinel terratest/applicationautoscaling_services_k8s_aws_v1alpha1/applicationautoscaling_services_k8s_aws_scalable_target_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_applicationautoscaling_services_k8s_aws_scalable_target_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/applicationautoscaling_services_k8s_aws_v1alpha1/applicationautoscaling_services_k8s_aws_scalable_target_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-applicationautoscaling_services_k8s_aws_scaling_policy_v1alpha1_manifest_test.go: out/install-sentinel terratest/applicationautoscaling_services_k8s_aws_v1alpha1/applicationautoscaling_services_k8s_aws_scaling_policy_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_applicationautoscaling_services_k8s_aws_scaling_policy_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/applicationautoscaling_services_k8s_aws_v1alpha1/applicationautoscaling_services_k8s_aws_scaling_policy_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-appmesh_k8s_aws_backend_group_v1beta2_manifest_test.go: out/install-sentinel terratest/appmesh_k8s_aws_v1beta2/appmesh_k8s_aws_backend_group_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_appmesh_k8s_aws_backend_group_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/appmesh_k8s_aws_v1beta2/appmesh_k8s_aws_backend_group_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-appmesh_k8s_aws_gateway_route_v1beta2_manifest_test.go: out/install-sentinel terratest/appmesh_k8s_aws_v1beta2/appmesh_k8s_aws_gateway_route_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_appmesh_k8s_aws_gateway_route_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/appmesh_k8s_aws_v1beta2/appmesh_k8s_aws_gateway_route_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-appmesh_k8s_aws_mesh_v1beta2_manifest_test.go: out/install-sentinel terratest/appmesh_k8s_aws_v1beta2/appmesh_k8s_aws_mesh_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_appmesh_k8s_aws_mesh_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/appmesh_k8s_aws_v1beta2/appmesh_k8s_aws_mesh_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-appmesh_k8s_aws_virtual_gateway_v1beta2_manifest_test.go: out/install-sentinel terratest/appmesh_k8s_aws_v1beta2/appmesh_k8s_aws_virtual_gateway_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_appmesh_k8s_aws_virtual_gateway_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/appmesh_k8s_aws_v1beta2/appmesh_k8s_aws_virtual_gateway_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-appmesh_k8s_aws_virtual_node_v1beta2_manifest_test.go: out/install-sentinel terratest/appmesh_k8s_aws_v1beta2/appmesh_k8s_aws_virtual_node_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_appmesh_k8s_aws_virtual_node_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/appmesh_k8s_aws_v1beta2/appmesh_k8s_aws_virtual_node_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-appmesh_k8s_aws_virtual_router_v1beta2_manifest_test.go: out/install-sentinel terratest/appmesh_k8s_aws_v1beta2/appmesh_k8s_aws_virtual_router_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_appmesh_k8s_aws_virtual_router_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/appmesh_k8s_aws_v1beta2/appmesh_k8s_aws_virtual_router_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-appmesh_k8s_aws_virtual_service_v1beta2_manifest_test.go: out/install-sentinel terratest/appmesh_k8s_aws_v1beta2/appmesh_k8s_aws_virtual_service_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_appmesh_k8s_aws_virtual_service_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/appmesh_k8s_aws_v1beta2/appmesh_k8s_aws_virtual_service_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-appprotect_f5_com_ap_log_conf_v1beta1_manifest_test.go: out/install-sentinel terratest/appprotect_f5_com_v1beta1/appprotect_f5_com_ap_log_conf_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_appprotect_f5_com_ap_log_conf_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/appprotect_f5_com_v1beta1/appprotect_f5_com_ap_log_conf_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-appprotect_f5_com_ap_policy_v1beta1_manifest_test.go: out/install-sentinel terratest/appprotect_f5_com_v1beta1/appprotect_f5_com_ap_policy_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_appprotect_f5_com_ap_policy_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/appprotect_f5_com_v1beta1/appprotect_f5_com_ap_policy_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-appprotect_f5_com_ap_user_sig_v1beta1_manifest_test.go: out/install-sentinel terratest/appprotect_f5_com_v1beta1/appprotect_f5_com_ap_user_sig_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_appprotect_f5_com_ap_user_sig_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/appprotect_f5_com_v1beta1/appprotect_f5_com_ap_user_sig_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-appprotectdos_f5_com_ap_dos_log_conf_v1beta1_manifest_test.go: out/install-sentinel terratest/appprotectdos_f5_com_v1beta1/appprotectdos_f5_com_ap_dos_log_conf_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_appprotectdos_f5_com_ap_dos_log_conf_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/appprotectdos_f5_com_v1beta1/appprotectdos_f5_com_ap_dos_log_conf_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-appprotectdos_f5_com_ap_dos_policy_v1beta1_manifest_test.go: out/install-sentinel terratest/appprotectdos_f5_com_v1beta1/appprotectdos_f5_com_ap_dos_policy_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_appprotectdos_f5_com_ap_dos_policy_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/appprotectdos_f5_com_v1beta1/appprotectdos_f5_com_ap_dos_policy_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-appprotectdos_f5_com_dos_protected_resource_v1beta1_manifest_test.go: out/install-sentinel terratest/appprotectdos_f5_com_v1beta1/appprotectdos_f5_com_dos_protected_resource_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_appprotectdos_f5_com_dos_protected_resource_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/appprotectdos_f5_com_v1beta1/appprotectdos_f5_com_dos_protected_resource_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_3scale_net_ap_icast_v1alpha1_manifest_test.go: out/install-sentinel terratest/apps_3scale_net_v1alpha1/apps_3scale_net_ap_icast_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_3scale_net_ap_icast_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_3scale_net_v1alpha1/apps_3scale_net_ap_icast_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_3scale_net_api_manager_backup_v1alpha1_manifest_test.go: out/install-sentinel terratest/apps_3scale_net_v1alpha1/apps_3scale_net_api_manager_backup_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_3scale_net_api_manager_backup_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_3scale_net_v1alpha1/apps_3scale_net_api_manager_backup_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_3scale_net_api_manager_restore_v1alpha1_manifest_test.go: out/install-sentinel terratest/apps_3scale_net_v1alpha1/apps_3scale_net_api_manager_restore_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_3scale_net_api_manager_restore_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_3scale_net_v1alpha1/apps_3scale_net_api_manager_restore_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_3scale_net_api_manager_v1alpha1_manifest_test.go: out/install-sentinel terratest/apps_3scale_net_v1alpha1/apps_3scale_net_api_manager_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_3scale_net_api_manager_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_3scale_net_v1alpha1/apps_3scale_net_api_manager_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_clusternet_io_base_v1alpha1_manifest_test.go: out/install-sentinel terratest/apps_clusternet_io_v1alpha1/apps_clusternet_io_base_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_clusternet_io_base_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_clusternet_io_v1alpha1/apps_clusternet_io_base_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_clusternet_io_description_v1alpha1_manifest_test.go: out/install-sentinel terratest/apps_clusternet_io_v1alpha1/apps_clusternet_io_description_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_clusternet_io_description_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_clusternet_io_v1alpha1/apps_clusternet_io_description_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_clusternet_io_feed_inventory_v1alpha1_manifest_test.go: out/install-sentinel terratest/apps_clusternet_io_v1alpha1/apps_clusternet_io_feed_inventory_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_clusternet_io_feed_inventory_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_clusternet_io_v1alpha1/apps_clusternet_io_feed_inventory_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_clusternet_io_globalization_v1alpha1_manifest_test.go: out/install-sentinel terratest/apps_clusternet_io_v1alpha1/apps_clusternet_io_globalization_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_clusternet_io_globalization_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_clusternet_io_v1alpha1/apps_clusternet_io_globalization_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_clusternet_io_helm_chart_v1alpha1_manifest_test.go: out/install-sentinel terratest/apps_clusternet_io_v1alpha1/apps_clusternet_io_helm_chart_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_clusternet_io_helm_chart_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_clusternet_io_v1alpha1/apps_clusternet_io_helm_chart_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_clusternet_io_helm_release_v1alpha1_manifest_test.go: out/install-sentinel terratest/apps_clusternet_io_v1alpha1/apps_clusternet_io_helm_release_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_clusternet_io_helm_release_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_clusternet_io_v1alpha1/apps_clusternet_io_helm_release_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_clusternet_io_localization_v1alpha1_manifest_test.go: out/install-sentinel terratest/apps_clusternet_io_v1alpha1/apps_clusternet_io_localization_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_clusternet_io_localization_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_clusternet_io_v1alpha1/apps_clusternet_io_localization_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_clusternet_io_manifest_v1alpha1_manifest_test.go: out/install-sentinel terratest/apps_clusternet_io_v1alpha1/apps_clusternet_io_manifest_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_clusternet_io_manifest_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_clusternet_io_v1alpha1/apps_clusternet_io_manifest_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_clusternet_io_subscription_v1alpha1_manifest_test.go: out/install-sentinel terratest/apps_clusternet_io_v1alpha1/apps_clusternet_io_subscription_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_clusternet_io_subscription_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_clusternet_io_v1alpha1/apps_clusternet_io_subscription_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_emqx_io_emqx_broker_v1beta3_manifest_test.go: out/install-sentinel terratest/apps_emqx_io_v1beta3/apps_emqx_io_emqx_broker_v1beta3_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_emqx_io_emqx_broker_v1beta3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_emqx_io_v1beta3/apps_emqx_io_emqx_broker_v1beta3_manifest_test.go
	touch $@
out/terratest-sentinel-apps_emqx_io_emqx_enterprise_v1beta3_manifest_test.go: out/install-sentinel terratest/apps_emqx_io_v1beta3/apps_emqx_io_emqx_enterprise_v1beta3_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_emqx_io_emqx_enterprise_v1beta3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_emqx_io_v1beta3/apps_emqx_io_emqx_enterprise_v1beta3_manifest_test.go
	touch $@
out/terratest-sentinel-apps_emqx_io_emqx_plugin_v1beta3_manifest_test.go: out/install-sentinel terratest/apps_emqx_io_v1beta3/apps_emqx_io_emqx_plugin_v1beta3_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_emqx_io_emqx_plugin_v1beta3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_emqx_io_v1beta3/apps_emqx_io_emqx_plugin_v1beta3_manifest_test.go
	touch $@
out/terratest-sentinel-apps_emqx_io_emqx_broker_v1beta4_manifest_test.go: out/install-sentinel terratest/apps_emqx_io_v1beta4/apps_emqx_io_emqx_broker_v1beta4_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_emqx_io_emqx_broker_v1beta4_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_emqx_io_v1beta4/apps_emqx_io_emqx_broker_v1beta4_manifest_test.go
	touch $@
out/terratest-sentinel-apps_emqx_io_emqx_enterprise_v1beta4_manifest_test.go: out/install-sentinel terratest/apps_emqx_io_v1beta4/apps_emqx_io_emqx_enterprise_v1beta4_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_emqx_io_emqx_enterprise_v1beta4_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_emqx_io_v1beta4/apps_emqx_io_emqx_enterprise_v1beta4_manifest_test.go
	touch $@
out/terratest-sentinel-apps_emqx_io_emqx_plugin_v1beta4_manifest_test.go: out/install-sentinel terratest/apps_emqx_io_v1beta4/apps_emqx_io_emqx_plugin_v1beta4_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_emqx_io_emqx_plugin_v1beta4_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_emqx_io_v1beta4/apps_emqx_io_emqx_plugin_v1beta4_manifest_test.go
	touch $@
out/terratest-sentinel-apps_emqx_io_rebalance_v1beta4_manifest_test.go: out/install-sentinel terratest/apps_emqx_io_v1beta4/apps_emqx_io_rebalance_v1beta4_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_emqx_io_rebalance_v1beta4_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_emqx_io_v1beta4/apps_emqx_io_rebalance_v1beta4_manifest_test.go
	touch $@
out/terratest-sentinel-apps_emqx_io_emqx_v2alpha1_manifest_test.go: out/install-sentinel terratest/apps_emqx_io_v2alpha1/apps_emqx_io_emqx_v2alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_emqx_io_emqx_v2alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_emqx_io_v2alpha1/apps_emqx_io_emqx_v2alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_emqx_io_emqx_v2beta1_manifest_test.go: out/install-sentinel terratest/apps_emqx_io_v2beta1/apps_emqx_io_emqx_v2beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_emqx_io_emqx_v2beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_emqx_io_v2beta1/apps_emqx_io_emqx_v2beta1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_emqx_io_rebalance_v2beta1_manifest_test.go: out/install-sentinel terratest/apps_emqx_io_v2beta1/apps_emqx_io_rebalance_v2beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_emqx_io_rebalance_v2beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_emqx_io_v2beta1/apps_emqx_io_rebalance_v2beta1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_gitlab_com_git_lab_v1beta1_manifest_test.go: out/install-sentinel terratest/apps_gitlab_com_v1beta1/apps_gitlab_com_git_lab_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_gitlab_com_git_lab_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_gitlab_com_v1beta1/apps_gitlab_com_git_lab_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_gitlab_com_runner_v1beta2_manifest_test.go: out/install-sentinel terratest/apps_gitlab_com_v1beta2/apps_gitlab_com_runner_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_gitlab_com_runner_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_gitlab_com_v1beta2/apps_gitlab_com_runner_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-apps_kubeblocks_io_backup_policy_template_v1alpha1_manifest_test.go: out/install-sentinel terratest/apps_kubeblocks_io_v1alpha1/apps_kubeblocks_io_backup_policy_template_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_kubeblocks_io_backup_policy_template_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_kubeblocks_io_v1alpha1/apps_kubeblocks_io_backup_policy_template_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_kubeblocks_io_cluster_definition_v1alpha1_manifest_test.go: out/install-sentinel terratest/apps_kubeblocks_io_v1alpha1/apps_kubeblocks_io_cluster_definition_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_kubeblocks_io_cluster_definition_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_kubeblocks_io_v1alpha1/apps_kubeblocks_io_cluster_definition_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_kubeblocks_io_cluster_v1alpha1_manifest_test.go: out/install-sentinel terratest/apps_kubeblocks_io_v1alpha1/apps_kubeblocks_io_cluster_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_kubeblocks_io_cluster_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_kubeblocks_io_v1alpha1/apps_kubeblocks_io_cluster_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_kubeblocks_io_cluster_version_v1alpha1_manifest_test.go: out/install-sentinel terratest/apps_kubeblocks_io_v1alpha1/apps_kubeblocks_io_cluster_version_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_kubeblocks_io_cluster_version_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_kubeblocks_io_v1alpha1/apps_kubeblocks_io_cluster_version_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_kubeblocks_io_component_class_definition_v1alpha1_manifest_test.go: out/install-sentinel terratest/apps_kubeblocks_io_v1alpha1/apps_kubeblocks_io_component_class_definition_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_kubeblocks_io_component_class_definition_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_kubeblocks_io_v1alpha1/apps_kubeblocks_io_component_class_definition_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_kubeblocks_io_component_definition_v1alpha1_manifest_test.go: out/install-sentinel terratest/apps_kubeblocks_io_v1alpha1/apps_kubeblocks_io_component_definition_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_kubeblocks_io_component_definition_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_kubeblocks_io_v1alpha1/apps_kubeblocks_io_component_definition_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_kubeblocks_io_component_resource_constraint_v1alpha1_manifest_test.go: out/install-sentinel terratest/apps_kubeblocks_io_v1alpha1/apps_kubeblocks_io_component_resource_constraint_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_kubeblocks_io_component_resource_constraint_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_kubeblocks_io_v1alpha1/apps_kubeblocks_io_component_resource_constraint_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_kubeblocks_io_component_v1alpha1_manifest_test.go: out/install-sentinel terratest/apps_kubeblocks_io_v1alpha1/apps_kubeblocks_io_component_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_kubeblocks_io_component_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_kubeblocks_io_v1alpha1/apps_kubeblocks_io_component_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_kubeblocks_io_component_version_v1alpha1_manifest_test.go: out/install-sentinel terratest/apps_kubeblocks_io_v1alpha1/apps_kubeblocks_io_component_version_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_kubeblocks_io_component_version_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_kubeblocks_io_v1alpha1/apps_kubeblocks_io_component_version_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_kubeblocks_io_config_constraint_v1alpha1_manifest_test.go: out/install-sentinel terratest/apps_kubeblocks_io_v1alpha1/apps_kubeblocks_io_config_constraint_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_kubeblocks_io_config_constraint_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_kubeblocks_io_v1alpha1/apps_kubeblocks_io_config_constraint_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_kubeblocks_io_configuration_v1alpha1_manifest_test.go: out/install-sentinel terratest/apps_kubeblocks_io_v1alpha1/apps_kubeblocks_io_configuration_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_kubeblocks_io_configuration_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_kubeblocks_io_v1alpha1/apps_kubeblocks_io_configuration_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_kubeblocks_io_ops_definition_v1alpha1_manifest_test.go: out/install-sentinel terratest/apps_kubeblocks_io_v1alpha1/apps_kubeblocks_io_ops_definition_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_kubeblocks_io_ops_definition_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_kubeblocks_io_v1alpha1/apps_kubeblocks_io_ops_definition_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_kubeblocks_io_ops_request_v1alpha1_manifest_test.go: out/install-sentinel terratest/apps_kubeblocks_io_v1alpha1/apps_kubeblocks_io_ops_request_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_kubeblocks_io_ops_request_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_kubeblocks_io_v1alpha1/apps_kubeblocks_io_ops_request_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_kubeblocks_io_service_descriptor_v1alpha1_manifest_test.go: out/install-sentinel terratest/apps_kubeblocks_io_v1alpha1/apps_kubeblocks_io_service_descriptor_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_kubeblocks_io_service_descriptor_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_kubeblocks_io_v1alpha1/apps_kubeblocks_io_service_descriptor_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_kubeblocks_io_config_constraint_v1beta1_manifest_test.go: out/install-sentinel terratest/apps_kubeblocks_io_v1beta1/apps_kubeblocks_io_config_constraint_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_kubeblocks_io_config_constraint_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_kubeblocks_io_v1beta1/apps_kubeblocks_io_config_constraint_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_kubedl_io_cron_v1alpha1_manifest_test.go: out/install-sentinel terratest/apps_kubedl_io_v1alpha1/apps_kubedl_io_cron_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_kubedl_io_cron_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_kubedl_io_v1alpha1/apps_kubedl_io_cron_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_kubeedge_io_edge_application_v1alpha1_manifest_test.go: out/install-sentinel terratest/apps_kubeedge_io_v1alpha1/apps_kubeedge_io_edge_application_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_kubeedge_io_edge_application_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_kubeedge_io_v1alpha1/apps_kubeedge_io_edge_application_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_kubeedge_io_node_group_v1alpha1_manifest_test.go: out/install-sentinel terratest/apps_kubeedge_io_v1alpha1/apps_kubeedge_io_node_group_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_kubeedge_io_node_group_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_kubeedge_io_v1alpha1/apps_kubeedge_io_node_group_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_m88i_io_nexus_v1alpha1_manifest_test.go: out/install-sentinel terratest/apps_m88i_io_v1alpha1/apps_m88i_io_nexus_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_m88i_io_nexus_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_m88i_io_v1alpha1/apps_m88i_io_nexus_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_redhat_com_cluster_impairment_v1alpha1_manifest_test.go: out/install-sentinel terratest/apps_redhat_com_v1alpha1/apps_redhat_com_cluster_impairment_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_redhat_com_cluster_impairment_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_redhat_com_v1alpha1/apps_redhat_com_cluster_impairment_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_daemon_set_v1_manifest_test.go: out/install-sentinel terratest/apps_v1/apps_daemon_set_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_daemon_set_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_v1/apps_daemon_set_v1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_deployment_v1_manifest_test.go: out/install-sentinel terratest/apps_v1/apps_deployment_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_deployment_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_v1/apps_deployment_v1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_replica_set_v1_manifest_test.go: out/install-sentinel terratest/apps_v1/apps_replica_set_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_replica_set_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_v1/apps_replica_set_v1_manifest_test.go
	touch $@
out/terratest-sentinel-apps_stateful_set_v1_manifest_test.go: out/install-sentinel terratest/apps_v1/apps_stateful_set_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_apps_stateful_set_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/apps_v1/apps_stateful_set_v1_manifest_test.go
	touch $@
out/terratest-sentinel-aquasecurity_github_io_aqua_starboard_v1alpha1_manifest_test.go: out/install-sentinel terratest/aquasecurity_github_io_v1alpha1/aquasecurity_github_io_aqua_starboard_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_aquasecurity_github_io_aqua_starboard_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/aquasecurity_github_io_v1alpha1/aquasecurity_github_io_aqua_starboard_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-argoproj_io_app_project_v1alpha1_manifest_test.go: out/install-sentinel terratest/argoproj_io_v1alpha1/argoproj_io_app_project_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_argoproj_io_app_project_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/argoproj_io_v1alpha1/argoproj_io_app_project_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-argoproj_io_application_set_v1alpha1_manifest_test.go: out/install-sentinel terratest/argoproj_io_v1alpha1/argoproj_io_application_set_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_argoproj_io_application_set_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/argoproj_io_v1alpha1/argoproj_io_application_set_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-argoproj_io_application_v1alpha1_manifest_test.go: out/install-sentinel terratest/argoproj_io_v1alpha1/argoproj_io_application_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_argoproj_io_application_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/argoproj_io_v1alpha1/argoproj_io_application_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-argoproj_io_argo_cd_v1alpha1_manifest_test.go: out/install-sentinel terratest/argoproj_io_v1alpha1/argoproj_io_argo_cd_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_argoproj_io_argo_cd_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/argoproj_io_v1alpha1/argoproj_io_argo_cd_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-argoproj_io_argo_cd_export_v1alpha1_manifest_test.go: out/install-sentinel terratest/argoproj_io_v1alpha1/argoproj_io_argo_cd_export_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_argoproj_io_argo_cd_export_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/argoproj_io_v1alpha1/argoproj_io_argo_cd_export_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-argoproj_io_argo_cd_v1beta1_manifest_test.go: out/install-sentinel terratest/argoproj_io_v1beta1/argoproj_io_argo_cd_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_argoproj_io_argo_cd_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/argoproj_io_v1beta1/argoproj_io_argo_cd_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-asdb_aerospike_com_aerospike_cluster_v1_manifest_test.go: out/install-sentinel terratest/asdb_aerospike_com_v1/asdb_aerospike_com_aerospike_cluster_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_asdb_aerospike_com_aerospike_cluster_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/asdb_aerospike_com_v1/asdb_aerospike_com_aerospike_cluster_v1_manifest_test.go
	touch $@
out/terratest-sentinel-asdb_aerospike_com_aerospike_cluster_v1beta1_manifest_test.go: out/install-sentinel terratest/asdb_aerospike_com_v1beta1/asdb_aerospike_com_aerospike_cluster_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_asdb_aerospike_com_aerospike_cluster_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/asdb_aerospike_com_v1beta1/asdb_aerospike_com_aerospike_cluster_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-atlasmap_io_atlas_map_v1alpha1_manifest_test.go: out/install-sentinel terratest/atlasmap_io_v1alpha1/atlasmap_io_atlas_map_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_atlasmap_io_atlas_map_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/atlasmap_io_v1alpha1/atlasmap_io_atlas_map_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-auth_ops42_org_aws_auth_sync_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/auth_ops42_org_v1alpha1/auth_ops42_org_aws_auth_sync_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_auth_ops42_org_aws_auth_sync_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/auth_ops42_org_v1alpha1/auth_ops42_org_aws_auth_sync_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-authzed_com_spice_db_cluster_v1alpha1_manifest_test.go: out/install-sentinel terratest/authzed_com_v1alpha1/authzed_com_spice_db_cluster_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_authzed_com_spice_db_cluster_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/authzed_com_v1alpha1/authzed_com_spice_db_cluster_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-automation_kubensync_com_managed_resource_v1alpha1_manifest_test.go: out/install-sentinel terratest/automation_kubensync_com_v1alpha1/automation_kubensync_com_managed_resource_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_automation_kubensync_com_managed_resource_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/automation_kubensync_com_v1alpha1/automation_kubensync_com_managed_resource_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-autoscaling_k8s_elastic_co_elasticsearch_autoscaler_v1alpha1_manifest_test.go: out/install-sentinel terratest/autoscaling_k8s_elastic_co_v1alpha1/autoscaling_k8s_elastic_co_elasticsearch_autoscaler_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_autoscaling_k8s_elastic_co_elasticsearch_autoscaler_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/autoscaling_k8s_elastic_co_v1alpha1/autoscaling_k8s_elastic_co_elasticsearch_autoscaler_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-autoscaling_k8s_io_vertical_pod_autoscaler_checkpoint_v1_manifest_test.go: out/install-sentinel terratest/autoscaling_k8s_io_v1/autoscaling_k8s_io_vertical_pod_autoscaler_checkpoint_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_autoscaling_k8s_io_vertical_pod_autoscaler_checkpoint_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/autoscaling_k8s_io_v1/autoscaling_k8s_io_vertical_pod_autoscaler_checkpoint_v1_manifest_test.go
	touch $@
out/terratest-sentinel-autoscaling_k8s_io_vertical_pod_autoscaler_v1_manifest_test.go: out/install-sentinel terratest/autoscaling_k8s_io_v1/autoscaling_k8s_io_vertical_pod_autoscaler_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_autoscaling_k8s_io_vertical_pod_autoscaler_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/autoscaling_k8s_io_v1/autoscaling_k8s_io_vertical_pod_autoscaler_v1_manifest_test.go
	touch $@
out/terratest-sentinel-autoscaling_k8s_io_vertical_pod_autoscaler_checkpoint_v1beta2_manifest_test.go: out/install-sentinel terratest/autoscaling_k8s_io_v1beta2/autoscaling_k8s_io_vertical_pod_autoscaler_checkpoint_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_autoscaling_k8s_io_vertical_pod_autoscaler_checkpoint_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/autoscaling_k8s_io_v1beta2/autoscaling_k8s_io_vertical_pod_autoscaler_checkpoint_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-autoscaling_k8s_io_vertical_pod_autoscaler_v1beta2_manifest_test.go: out/install-sentinel terratest/autoscaling_k8s_io_v1beta2/autoscaling_k8s_io_vertical_pod_autoscaler_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_autoscaling_k8s_io_vertical_pod_autoscaler_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/autoscaling_k8s_io_v1beta2/autoscaling_k8s_io_vertical_pod_autoscaler_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-autoscaling_karmada_io_cron_federated_hpa_v1alpha1_manifest_test.go: out/install-sentinel terratest/autoscaling_karmada_io_v1alpha1/autoscaling_karmada_io_cron_federated_hpa_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_autoscaling_karmada_io_cron_federated_hpa_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/autoscaling_karmada_io_v1alpha1/autoscaling_karmada_io_cron_federated_hpa_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-autoscaling_karmada_io_federated_hpa_v1alpha1_manifest_test.go: out/install-sentinel terratest/autoscaling_karmada_io_v1alpha1/autoscaling_karmada_io_federated_hpa_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_autoscaling_karmada_io_federated_hpa_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/autoscaling_karmada_io_v1alpha1/autoscaling_karmada_io_federated_hpa_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-autoscaling_horizontal_pod_autoscaler_v1_manifest_test.go: out/install-sentinel terratest/autoscaling_v1/autoscaling_horizontal_pod_autoscaler_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_autoscaling_horizontal_pod_autoscaler_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/autoscaling_v1/autoscaling_horizontal_pod_autoscaler_v1_manifest_test.go
	touch $@
out/terratest-sentinel-autoscaling_horizontal_pod_autoscaler_v2_manifest_test.go: out/install-sentinel terratest/autoscaling_v2/autoscaling_horizontal_pod_autoscaler_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_autoscaling_horizontal_pod_autoscaler_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/autoscaling_v2/autoscaling_horizontal_pod_autoscaler_v2_manifest_test.go
	touch $@
out/terratest-sentinel-awx_ansible_com_awx_v1beta1_manifest_test.go: out/install-sentinel terratest/awx_ansible_com_v1beta1/awx_ansible_com_awx_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_awx_ansible_com_awx_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/awx_ansible_com_v1beta1/awx_ansible_com_awx_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-awx_ansible_com_awx_backup_v1beta1_manifest_test.go: out/install-sentinel terratest/awx_ansible_com_v1beta1/awx_ansible_com_awx_backup_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_awx_ansible_com_awx_backup_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/awx_ansible_com_v1beta1/awx_ansible_com_awx_backup_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-awx_ansible_com_awx_restore_v1beta1_manifest_test.go: out/install-sentinel terratest/awx_ansible_com_v1beta1/awx_ansible_com_awx_restore_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_awx_ansible_com_awx_restore_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/awx_ansible_com_v1beta1/awx_ansible_com_awx_restore_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_apim_service_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_apim_service_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_apim_service_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_apim_service_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_api_mgmt_api_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_api_mgmt_api_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_api_mgmt_api_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_api_mgmt_api_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_app_insights_api_key_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_app_insights_api_key_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_app_insights_api_key_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_app_insights_api_key_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_app_insights_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_app_insights_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_app_insights_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_app_insights_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_azure_load_balancer_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_azure_load_balancer_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_azure_load_balancer_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_azure_load_balancer_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_azure_network_interface_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_azure_network_interface_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_azure_network_interface_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_azure_network_interface_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_azure_public_ip_address_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_azure_public_ip_address_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_azure_public_ip_address_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_azure_public_ip_address_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_azure_sql_action_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_azure_sql_action_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_azure_sql_action_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_azure_sql_action_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_azure_sql_database_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_azure_sql_database_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_azure_sql_database_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_azure_sql_database_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_azure_sql_failover_group_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_azure_sql_failover_group_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_azure_sql_failover_group_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_azure_sql_failover_group_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_azure_sql_firewall_rule_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_azure_sql_firewall_rule_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_azure_sql_firewall_rule_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_azure_sql_firewall_rule_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_azure_sql_server_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_azure_sql_server_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_azure_sql_server_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_azure_sql_server_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_azure_sql_managed_user_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_azure_sql_managed_user_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_azure_sql_managed_user_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_azure_sql_managed_user_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_azure_sql_user_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_azure_sql_user_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_azure_sql_user_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_azure_sql_user_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_azure_sqlv_net_rule_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_azure_sqlv_net_rule_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_azure_sqlv_net_rule_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_azure_sqlv_net_rule_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_azure_virtual_machine_extension_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_azure_virtual_machine_extension_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_azure_virtual_machine_extension_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_azure_virtual_machine_extension_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_azure_virtual_machine_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_azure_virtual_machine_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_azure_virtual_machine_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_azure_virtual_machine_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_azure_vm_scale_set_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_azure_vm_scale_set_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_azure_vm_scale_set_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_azure_vm_scale_set_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_blob_container_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_blob_container_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_blob_container_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_blob_container_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_consumer_group_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_consumer_group_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_consumer_group_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_consumer_group_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_cosmos_db_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_cosmos_db_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_cosmos_db_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_cosmos_db_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_eventhub_namespace_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_eventhub_namespace_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_eventhub_namespace_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_eventhub_namespace_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_eventhub_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_eventhub_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_eventhub_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_eventhub_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_key_vault_key_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_key_vault_key_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_key_vault_key_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_key_vault_key_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_key_vault_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_key_vault_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_key_vault_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_key_vault_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_my_sqlaad_user_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_my_sqlaad_user_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_my_sqlaad_user_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_my_sqlaad_user_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_my_sql_database_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_my_sql_database_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_my_sql_database_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_my_sql_database_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_my_sql_firewall_rule_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_my_sql_firewall_rule_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_my_sql_firewall_rule_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_my_sql_firewall_rule_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_my_sql_server_administrator_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_my_sql_server_administrator_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_my_sql_server_administrator_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_my_sql_server_administrator_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_my_sql_server_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_my_sql_server_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_my_sql_server_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_my_sql_server_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_my_sql_user_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_my_sql_user_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_my_sql_user_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_my_sql_user_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_my_sqlv_net_rule_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_my_sqlv_net_rule_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_my_sqlv_net_rule_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_my_sqlv_net_rule_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_postgre_sql_database_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_postgre_sql_database_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_postgre_sql_database_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_postgre_sql_database_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_postgre_sql_firewall_rule_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_postgre_sql_firewall_rule_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_postgre_sql_firewall_rule_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_postgre_sql_firewall_rule_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_postgre_sql_server_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_postgre_sql_server_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_postgre_sql_server_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_postgre_sql_server_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_postgre_sql_user_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_postgre_sql_user_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_postgre_sql_user_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_postgre_sql_user_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_postgre_sqlv_net_rule_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_postgre_sqlv_net_rule_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_postgre_sqlv_net_rule_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_postgre_sqlv_net_rule_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_redis_cache_action_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_redis_cache_action_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_redis_cache_action_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_redis_cache_action_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_redis_cache_firewall_rule_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_redis_cache_firewall_rule_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_redis_cache_firewall_rule_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_redis_cache_firewall_rule_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_resource_group_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_resource_group_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_resource_group_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_resource_group_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_storage_account_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_storage_account_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_storage_account_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_storage_account_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_virtual_network_v1alpha1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_virtual_network_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_virtual_network_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha1/azure_microsoft_com_virtual_network_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_blob_container_v1alpha2_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha2/azure_microsoft_com_blob_container_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_blob_container_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha2/azure_microsoft_com_blob_container_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_my_sqlaad_user_v1alpha2_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha2/azure_microsoft_com_my_sqlaad_user_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_my_sqlaad_user_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha2/azure_microsoft_com_my_sqlaad_user_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_my_sql_server_v1alpha2_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha2/azure_microsoft_com_my_sql_server_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_my_sql_server_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha2/azure_microsoft_com_my_sql_server_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_my_sql_user_v1alpha2_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha2/azure_microsoft_com_my_sql_user_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_my_sql_user_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha2/azure_microsoft_com_my_sql_user_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_postgre_sql_server_v1alpha2_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1alpha2/azure_microsoft_com_postgre_sql_server_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_postgre_sql_server_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1alpha2/azure_microsoft_com_postgre_sql_server_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_azure_sql_database_v1beta1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1beta1/azure_microsoft_com_azure_sql_database_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_azure_sql_database_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1beta1/azure_microsoft_com_azure_sql_database_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_azure_sql_failover_group_v1beta1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1beta1/azure_microsoft_com_azure_sql_failover_group_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_azure_sql_failover_group_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1beta1/azure_microsoft_com_azure_sql_failover_group_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_azure_sql_firewall_rule_v1beta1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1beta1/azure_microsoft_com_azure_sql_firewall_rule_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_azure_sql_firewall_rule_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1beta1/azure_microsoft_com_azure_sql_firewall_rule_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-azure_microsoft_com_azure_sql_server_v1beta1_manifest_test.go: out/install-sentinel terratest/azure_microsoft_com_v1beta1/azure_microsoft_com_azure_sql_server_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_azure_microsoft_com_azure_sql_server_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/azure_microsoft_com_v1beta1/azure_microsoft_com_azure_sql_server_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-b3scale_infra_run_bbb_frontend_v1_manifest_test.go: out/install-sentinel terratest/b3scale_infra_run_v1/b3scale_infra_run_bbb_frontend_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_b3scale_infra_run_bbb_frontend_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/b3scale_infra_run_v1/b3scale_infra_run_bbb_frontend_v1_manifest_test.go
	touch $@
out/terratest-sentinel-b3scale_io_bbb_frontend_v1_manifest_test.go: out/install-sentinel terratest/b3scale_io_v1/b3scale_io_bbb_frontend_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_b3scale_io_bbb_frontend_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/b3scale_io_v1/b3scale_io_bbb_frontend_v1_manifest_test.go
	touch $@
out/terratest-sentinel-batch_cron_job_v1_manifest_test.go: out/install-sentinel terratest/batch_v1/batch_cron_job_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_batch_cron_job_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/batch_v1/batch_cron_job_v1_manifest_test.go
	touch $@
out/terratest-sentinel-batch_job_v1_manifest_test.go: out/install-sentinel terratest/batch_v1/batch_job_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_batch_job_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/batch_v1/batch_job_v1_manifest_test.go
	touch $@
out/terratest-sentinel-batch_volcano_sh_job_v1alpha1_manifest_test.go: out/install-sentinel terratest/batch_volcano_sh_v1alpha1/batch_volcano_sh_job_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_batch_volcano_sh_job_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/batch_volcano_sh_v1alpha1/batch_volcano_sh_job_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-beat_k8s_elastic_co_beat_v1beta1_manifest_test.go: out/install-sentinel terratest/beat_k8s_elastic_co_v1beta1/beat_k8s_elastic_co_beat_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_beat_k8s_elastic_co_beat_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/beat_k8s_elastic_co_v1beta1/beat_k8s_elastic_co_beat_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-beegfs_csi_netapp_com_beegfs_driver_v1_manifest_test.go: out/install-sentinel terratest/beegfs_csi_netapp_com_v1/beegfs_csi_netapp_com_beegfs_driver_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_beegfs_csi_netapp_com_beegfs_driver_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/beegfs_csi_netapp_com_v1/beegfs_csi_netapp_com_beegfs_driver_v1_manifest_test.go
	touch $@
out/terratest-sentinel-binding_operators_coreos_com_service_binding_v1alpha1_manifest_test.go: out/install-sentinel terratest/binding_operators_coreos_com_v1alpha1/binding_operators_coreos_com_service_binding_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_binding_operators_coreos_com_service_binding_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/binding_operators_coreos_com_v1alpha1/binding_operators_coreos_com_service_binding_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-bitnami_com_sealed_secret_v1alpha1_manifest_test.go: out/install-sentinel terratest/bitnami_com_v1alpha1/bitnami_com_sealed_secret_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_bitnami_com_sealed_secret_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/bitnami_com_v1alpha1/bitnami_com_sealed_secret_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-bmc_tinkerbell_org_job_v1alpha1_manifest_test.go: out/install-sentinel terratest/bmc_tinkerbell_org_v1alpha1/bmc_tinkerbell_org_job_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_bmc_tinkerbell_org_job_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/bmc_tinkerbell_org_v1alpha1/bmc_tinkerbell_org_job_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-bmc_tinkerbell_org_machine_v1alpha1_manifest_test.go: out/install-sentinel terratest/bmc_tinkerbell_org_v1alpha1/bmc_tinkerbell_org_machine_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_bmc_tinkerbell_org_machine_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/bmc_tinkerbell_org_v1alpha1/bmc_tinkerbell_org_machine_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-bmc_tinkerbell_org_task_v1alpha1_manifest_test.go: out/install-sentinel terratest/bmc_tinkerbell_org_v1alpha1/bmc_tinkerbell_org_task_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_bmc_tinkerbell_org_task_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/bmc_tinkerbell_org_v1alpha1/bmc_tinkerbell_org_task_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-boskos_k8s_io_drlc_object_v1_manifest_test.go: out/install-sentinel terratest/boskos_k8s_io_v1/boskos_k8s_io_drlc_object_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_boskos_k8s_io_drlc_object_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/boskos_k8s_io_v1/boskos_k8s_io_drlc_object_v1_manifest_test.go
	touch $@
out/terratest-sentinel-boskos_k8s_io_resource_object_v1_manifest_test.go: out/install-sentinel terratest/boskos_k8s_io_v1/boskos_k8s_io_resource_object_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_boskos_k8s_io_resource_object_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/boskos_k8s_io_v1/boskos_k8s_io_resource_object_v1_manifest_test.go
	touch $@
out/terratest-sentinel-bpfman_io_bpf_program_v1alpha1_manifest_test.go: out/install-sentinel terratest/bpfman_io_v1alpha1/bpfman_io_bpf_program_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_bpfman_io_bpf_program_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/bpfman_io_v1alpha1/bpfman_io_bpf_program_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-bpfman_io_fentry_program_v1alpha1_manifest_test.go: out/install-sentinel terratest/bpfman_io_v1alpha1/bpfman_io_fentry_program_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_bpfman_io_fentry_program_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/bpfman_io_v1alpha1/bpfman_io_fentry_program_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-bpfman_io_fexit_program_v1alpha1_manifest_test.go: out/install-sentinel terratest/bpfman_io_v1alpha1/bpfman_io_fexit_program_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_bpfman_io_fexit_program_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/bpfman_io_v1alpha1/bpfman_io_fexit_program_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-bpfman_io_kprobe_program_v1alpha1_manifest_test.go: out/install-sentinel terratest/bpfman_io_v1alpha1/bpfman_io_kprobe_program_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_bpfman_io_kprobe_program_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/bpfman_io_v1alpha1/bpfman_io_kprobe_program_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-bpfman_io_tc_program_v1alpha1_manifest_test.go: out/install-sentinel terratest/bpfman_io_v1alpha1/bpfman_io_tc_program_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_bpfman_io_tc_program_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/bpfman_io_v1alpha1/bpfman_io_tc_program_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-bpfman_io_tracepoint_program_v1alpha1_manifest_test.go: out/install-sentinel terratest/bpfman_io_v1alpha1/bpfman_io_tracepoint_program_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_bpfman_io_tracepoint_program_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/bpfman_io_v1alpha1/bpfman_io_tracepoint_program_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-bpfman_io_uprobe_program_v1alpha1_manifest_test.go: out/install-sentinel terratest/bpfman_io_v1alpha1/bpfman_io_uprobe_program_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_bpfman_io_uprobe_program_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/bpfman_io_v1alpha1/bpfman_io_uprobe_program_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-bpfman_io_xdp_program_v1alpha1_manifest_test.go: out/install-sentinel terratest/bpfman_io_v1alpha1/bpfman_io_xdp_program_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_bpfman_io_xdp_program_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/bpfman_io_v1alpha1/bpfman_io_xdp_program_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-bus_volcano_sh_command_v1alpha1_manifest_test.go: out/install-sentinel terratest/bus_volcano_sh_v1alpha1/bus_volcano_sh_command_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_bus_volcano_sh_command_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/bus_volcano_sh_v1alpha1/bus_volcano_sh_command_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-cache_kubedl_io_cache_backend_v1alpha1_manifest_test.go: out/install-sentinel terratest/cache_kubedl_io_v1alpha1/cache_kubedl_io_cache_backend_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_cache_kubedl_io_cache_backend_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cache_kubedl_io_v1alpha1/cache_kubedl_io_cache_backend_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-caching_ibm_com_varnish_cluster_v1alpha1_manifest_test.go: out/install-sentinel terratest/caching_ibm_com_v1alpha1/caching_ibm_com_varnish_cluster_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_caching_ibm_com_varnish_cluster_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/caching_ibm_com_v1alpha1/caching_ibm_com_varnish_cluster_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-camel_apache_org_build_v1_manifest_test.go: out/install-sentinel terratest/camel_apache_org_v1/camel_apache_org_build_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_camel_apache_org_build_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/camel_apache_org_v1/camel_apache_org_build_v1_manifest_test.go
	touch $@
out/terratest-sentinel-camel_apache_org_camel_catalog_v1_manifest_test.go: out/install-sentinel terratest/camel_apache_org_v1/camel_apache_org_camel_catalog_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_camel_apache_org_camel_catalog_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/camel_apache_org_v1/camel_apache_org_camel_catalog_v1_manifest_test.go
	touch $@
out/terratest-sentinel-camel_apache_org_integration_kit_v1_manifest_test.go: out/install-sentinel terratest/camel_apache_org_v1/camel_apache_org_integration_kit_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_camel_apache_org_integration_kit_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/camel_apache_org_v1/camel_apache_org_integration_kit_v1_manifest_test.go
	touch $@
out/terratest-sentinel-camel_apache_org_integration_platform_v1_manifest_test.go: out/install-sentinel terratest/camel_apache_org_v1/camel_apache_org_integration_platform_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_camel_apache_org_integration_platform_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/camel_apache_org_v1/camel_apache_org_integration_platform_v1_manifest_test.go
	touch $@
out/terratest-sentinel-camel_apache_org_integration_profile_v1_manifest_test.go: out/install-sentinel terratest/camel_apache_org_v1/camel_apache_org_integration_profile_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_camel_apache_org_integration_profile_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/camel_apache_org_v1/camel_apache_org_integration_profile_v1_manifest_test.go
	touch $@
out/terratest-sentinel-camel_apache_org_integration_v1_manifest_test.go: out/install-sentinel terratest/camel_apache_org_v1/camel_apache_org_integration_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_camel_apache_org_integration_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/camel_apache_org_v1/camel_apache_org_integration_v1_manifest_test.go
	touch $@
out/terratest-sentinel-camel_apache_org_kamelet_v1_manifest_test.go: out/install-sentinel terratest/camel_apache_org_v1/camel_apache_org_kamelet_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_camel_apache_org_kamelet_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/camel_apache_org_v1/camel_apache_org_kamelet_v1_manifest_test.go
	touch $@
out/terratest-sentinel-camel_apache_org_pipe_v1_manifest_test.go: out/install-sentinel terratest/camel_apache_org_v1/camel_apache_org_pipe_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_camel_apache_org_pipe_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/camel_apache_org_v1/camel_apache_org_pipe_v1_manifest_test.go
	touch $@
out/terratest-sentinel-camel_apache_org_kamelet_binding_v1alpha1_manifest_test.go: out/install-sentinel terratest/camel_apache_org_v1alpha1/camel_apache_org_kamelet_binding_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_camel_apache_org_kamelet_binding_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/camel_apache_org_v1alpha1/camel_apache_org_kamelet_binding_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-camel_apache_org_kamelet_v1alpha1_manifest_test.go: out/install-sentinel terratest/camel_apache_org_v1alpha1/camel_apache_org_kamelet_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_camel_apache_org_kamelet_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/camel_apache_org_v1alpha1/camel_apache_org_kamelet_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-canaries_flanksource_com_canary_v1_manifest_test.go: out/install-sentinel terratest/canaries_flanksource_com_v1/canaries_flanksource_com_canary_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_canaries_flanksource_com_canary_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/canaries_flanksource_com_v1/canaries_flanksource_com_canary_v1_manifest_test.go
	touch $@
out/terratest-sentinel-canaries_flanksource_com_component_v1_manifest_test.go: out/install-sentinel terratest/canaries_flanksource_com_v1/canaries_flanksource_com_component_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_canaries_flanksource_com_component_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/canaries_flanksource_com_v1/canaries_flanksource_com_component_v1_manifest_test.go
	touch $@
out/terratest-sentinel-canaries_flanksource_com_topology_v1_manifest_test.go: out/install-sentinel terratest/canaries_flanksource_com_v1/canaries_flanksource_com_topology_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_canaries_flanksource_com_topology_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/canaries_flanksource_com_v1/canaries_flanksource_com_topology_v1_manifest_test.go
	touch $@
out/terratest-sentinel-capabilities_3scale_net_tenant_v1alpha1_manifest_test.go: out/install-sentinel terratest/capabilities_3scale_net_v1alpha1/capabilities_3scale_net_tenant_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_capabilities_3scale_net_tenant_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/capabilities_3scale_net_v1alpha1/capabilities_3scale_net_tenant_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-capabilities_3scale_net_active_doc_v1beta1_manifest_test.go: out/install-sentinel terratest/capabilities_3scale_net_v1beta1/capabilities_3scale_net_active_doc_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_capabilities_3scale_net_active_doc_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/capabilities_3scale_net_v1beta1/capabilities_3scale_net_active_doc_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-capabilities_3scale_net_application_v1beta1_manifest_test.go: out/install-sentinel terratest/capabilities_3scale_net_v1beta1/capabilities_3scale_net_application_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_capabilities_3scale_net_application_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/capabilities_3scale_net_v1beta1/capabilities_3scale_net_application_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-capabilities_3scale_net_backend_v1beta1_manifest_test.go: out/install-sentinel terratest/capabilities_3scale_net_v1beta1/capabilities_3scale_net_backend_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_capabilities_3scale_net_backend_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/capabilities_3scale_net_v1beta1/capabilities_3scale_net_backend_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-capabilities_3scale_net_custom_policy_definition_v1beta1_manifest_test.go: out/install-sentinel terratest/capabilities_3scale_net_v1beta1/capabilities_3scale_net_custom_policy_definition_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_capabilities_3scale_net_custom_policy_definition_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/capabilities_3scale_net_v1beta1/capabilities_3scale_net_custom_policy_definition_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-capabilities_3scale_net_developer_account_v1beta1_manifest_test.go: out/install-sentinel terratest/capabilities_3scale_net_v1beta1/capabilities_3scale_net_developer_account_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_capabilities_3scale_net_developer_account_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/capabilities_3scale_net_v1beta1/capabilities_3scale_net_developer_account_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-capabilities_3scale_net_developer_user_v1beta1_manifest_test.go: out/install-sentinel terratest/capabilities_3scale_net_v1beta1/capabilities_3scale_net_developer_user_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_capabilities_3scale_net_developer_user_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/capabilities_3scale_net_v1beta1/capabilities_3scale_net_developer_user_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-capabilities_3scale_net_open_api_v1beta1_manifest_test.go: out/install-sentinel terratest/capabilities_3scale_net_v1beta1/capabilities_3scale_net_open_api_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_capabilities_3scale_net_open_api_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/capabilities_3scale_net_v1beta1/capabilities_3scale_net_open_api_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-capabilities_3scale_net_product_v1beta1_manifest_test.go: out/install-sentinel terratest/capabilities_3scale_net_v1beta1/capabilities_3scale_net_product_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_capabilities_3scale_net_product_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/capabilities_3scale_net_v1beta1/capabilities_3scale_net_product_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-capabilities_3scale_net_proxy_config_promote_v1beta1_manifest_test.go: out/install-sentinel terratest/capabilities_3scale_net_v1beta1/capabilities_3scale_net_proxy_config_promote_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_capabilities_3scale_net_proxy_config_promote_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/capabilities_3scale_net_v1beta1/capabilities_3scale_net_proxy_config_promote_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-capsule_clastix_io_capsule_configuration_v1alpha1_manifest_test.go: out/install-sentinel terratest/capsule_clastix_io_v1alpha1/capsule_clastix_io_capsule_configuration_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_capsule_clastix_io_capsule_configuration_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/capsule_clastix_io_v1alpha1/capsule_clastix_io_capsule_configuration_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-capsule_clastix_io_tenant_v1alpha1_manifest_test.go: out/install-sentinel terratest/capsule_clastix_io_v1alpha1/capsule_clastix_io_tenant_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_capsule_clastix_io_tenant_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/capsule_clastix_io_v1alpha1/capsule_clastix_io_tenant_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-capsule_clastix_io_tenant_v1beta1_manifest_test.go: out/install-sentinel terratest/capsule_clastix_io_v1beta1/capsule_clastix_io_tenant_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_capsule_clastix_io_tenant_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/capsule_clastix_io_v1beta1/capsule_clastix_io_tenant_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-capsule_clastix_io_capsule_configuration_v1beta2_manifest_test.go: out/install-sentinel terratest/capsule_clastix_io_v1beta2/capsule_clastix_io_capsule_configuration_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_capsule_clastix_io_capsule_configuration_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/capsule_clastix_io_v1beta2/capsule_clastix_io_capsule_configuration_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-capsule_clastix_io_tenant_v1beta2_manifest_test.go: out/install-sentinel terratest/capsule_clastix_io_v1beta2/capsule_clastix_io_tenant_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_capsule_clastix_io_tenant_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/capsule_clastix_io_v1beta2/capsule_clastix_io_tenant_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-cassandra_datastax_com_cassandra_datacenter_v1beta1_manifest_test.go: out/install-sentinel terratest/cassandra_datastax_com_v1beta1/cassandra_datastax_com_cassandra_datacenter_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_cassandra_datastax_com_cassandra_datacenter_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cassandra_datastax_com_v1beta1/cassandra_datastax_com_cassandra_datacenter_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-ceph_rook_io_ceph_block_pool_rados_namespace_v1_manifest_test.go: out/install-sentinel terratest/ceph_rook_io_v1/ceph_rook_io_ceph_block_pool_rados_namespace_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_ceph_rook_io_ceph_block_pool_rados_namespace_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ceph_rook_io_v1/ceph_rook_io_ceph_block_pool_rados_namespace_v1_manifest_test.go
	touch $@
out/terratest-sentinel-ceph_rook_io_ceph_block_pool_v1_manifest_test.go: out/install-sentinel terratest/ceph_rook_io_v1/ceph_rook_io_ceph_block_pool_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_ceph_rook_io_ceph_block_pool_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ceph_rook_io_v1/ceph_rook_io_ceph_block_pool_v1_manifest_test.go
	touch $@
out/terratest-sentinel-ceph_rook_io_ceph_bucket_notification_v1_manifest_test.go: out/install-sentinel terratest/ceph_rook_io_v1/ceph_rook_io_ceph_bucket_notification_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_ceph_rook_io_ceph_bucket_notification_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ceph_rook_io_v1/ceph_rook_io_ceph_bucket_notification_v1_manifest_test.go
	touch $@
out/terratest-sentinel-ceph_rook_io_ceph_bucket_topic_v1_manifest_test.go: out/install-sentinel terratest/ceph_rook_io_v1/ceph_rook_io_ceph_bucket_topic_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_ceph_rook_io_ceph_bucket_topic_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ceph_rook_io_v1/ceph_rook_io_ceph_bucket_topic_v1_manifest_test.go
	touch $@
out/terratest-sentinel-ceph_rook_io_ceph_client_v1_manifest_test.go: out/install-sentinel terratest/ceph_rook_io_v1/ceph_rook_io_ceph_client_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_ceph_rook_io_ceph_client_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ceph_rook_io_v1/ceph_rook_io_ceph_client_v1_manifest_test.go
	touch $@
out/terratest-sentinel-ceph_rook_io_ceph_cluster_v1_manifest_test.go: out/install-sentinel terratest/ceph_rook_io_v1/ceph_rook_io_ceph_cluster_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_ceph_rook_io_ceph_cluster_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ceph_rook_io_v1/ceph_rook_io_ceph_cluster_v1_manifest_test.go
	touch $@
out/terratest-sentinel-ceph_rook_io_ceph_cosi_driver_v1_manifest_test.go: out/install-sentinel terratest/ceph_rook_io_v1/ceph_rook_io_ceph_cosi_driver_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_ceph_rook_io_ceph_cosi_driver_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ceph_rook_io_v1/ceph_rook_io_ceph_cosi_driver_v1_manifest_test.go
	touch $@
out/terratest-sentinel-ceph_rook_io_ceph_filesystem_mirror_v1_manifest_test.go: out/install-sentinel terratest/ceph_rook_io_v1/ceph_rook_io_ceph_filesystem_mirror_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_ceph_rook_io_ceph_filesystem_mirror_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ceph_rook_io_v1/ceph_rook_io_ceph_filesystem_mirror_v1_manifest_test.go
	touch $@
out/terratest-sentinel-ceph_rook_io_ceph_filesystem_sub_volume_group_v1_manifest_test.go: out/install-sentinel terratest/ceph_rook_io_v1/ceph_rook_io_ceph_filesystem_sub_volume_group_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_ceph_rook_io_ceph_filesystem_sub_volume_group_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ceph_rook_io_v1/ceph_rook_io_ceph_filesystem_sub_volume_group_v1_manifest_test.go
	touch $@
out/terratest-sentinel-ceph_rook_io_ceph_filesystem_v1_manifest_test.go: out/install-sentinel terratest/ceph_rook_io_v1/ceph_rook_io_ceph_filesystem_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_ceph_rook_io_ceph_filesystem_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ceph_rook_io_v1/ceph_rook_io_ceph_filesystem_v1_manifest_test.go
	touch $@
out/terratest-sentinel-ceph_rook_io_ceph_nfs_v1_manifest_test.go: out/install-sentinel terratest/ceph_rook_io_v1/ceph_rook_io_ceph_nfs_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_ceph_rook_io_ceph_nfs_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ceph_rook_io_v1/ceph_rook_io_ceph_nfs_v1_manifest_test.go
	touch $@
out/terratest-sentinel-ceph_rook_io_ceph_object_realm_v1_manifest_test.go: out/install-sentinel terratest/ceph_rook_io_v1/ceph_rook_io_ceph_object_realm_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_ceph_rook_io_ceph_object_realm_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ceph_rook_io_v1/ceph_rook_io_ceph_object_realm_v1_manifest_test.go
	touch $@
out/terratest-sentinel-ceph_rook_io_ceph_object_store_user_v1_manifest_test.go: out/install-sentinel terratest/ceph_rook_io_v1/ceph_rook_io_ceph_object_store_user_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_ceph_rook_io_ceph_object_store_user_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ceph_rook_io_v1/ceph_rook_io_ceph_object_store_user_v1_manifest_test.go
	touch $@
out/terratest-sentinel-ceph_rook_io_ceph_object_store_v1_manifest_test.go: out/install-sentinel terratest/ceph_rook_io_v1/ceph_rook_io_ceph_object_store_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_ceph_rook_io_ceph_object_store_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ceph_rook_io_v1/ceph_rook_io_ceph_object_store_v1_manifest_test.go
	touch $@
out/terratest-sentinel-ceph_rook_io_ceph_object_zone_group_v1_manifest_test.go: out/install-sentinel terratest/ceph_rook_io_v1/ceph_rook_io_ceph_object_zone_group_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_ceph_rook_io_ceph_object_zone_group_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ceph_rook_io_v1/ceph_rook_io_ceph_object_zone_group_v1_manifest_test.go
	touch $@
out/terratest-sentinel-ceph_rook_io_ceph_object_zone_v1_manifest_test.go: out/install-sentinel terratest/ceph_rook_io_v1/ceph_rook_io_ceph_object_zone_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_ceph_rook_io_ceph_object_zone_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ceph_rook_io_v1/ceph_rook_io_ceph_object_zone_v1_manifest_test.go
	touch $@
out/terratest-sentinel-ceph_rook_io_ceph_rbd_mirror_v1_manifest_test.go: out/install-sentinel terratest/ceph_rook_io_v1/ceph_rook_io_ceph_rbd_mirror_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_ceph_rook_io_ceph_rbd_mirror_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ceph_rook_io_v1/ceph_rook_io_ceph_rbd_mirror_v1_manifest_test.go
	touch $@
out/terratest-sentinel-cert_manager_io_certificate_request_v1_manifest_test.go: out/install-sentinel terratest/cert_manager_io_v1/cert_manager_io_certificate_request_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_cert_manager_io_certificate_request_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cert_manager_io_v1/cert_manager_io_certificate_request_v1_manifest_test.go
	touch $@
out/terratest-sentinel-cert_manager_io_certificate_v1_manifest_test.go: out/install-sentinel terratest/cert_manager_io_v1/cert_manager_io_certificate_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_cert_manager_io_certificate_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cert_manager_io_v1/cert_manager_io_certificate_v1_manifest_test.go
	touch $@
out/terratest-sentinel-cert_manager_io_cluster_issuer_v1_manifest_test.go: out/install-sentinel terratest/cert_manager_io_v1/cert_manager_io_cluster_issuer_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_cert_manager_io_cluster_issuer_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cert_manager_io_v1/cert_manager_io_cluster_issuer_v1_manifest_test.go
	touch $@
out/terratest-sentinel-cert_manager_io_issuer_v1_manifest_test.go: out/install-sentinel terratest/cert_manager_io_v1/cert_manager_io_issuer_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_cert_manager_io_issuer_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cert_manager_io_v1/cert_manager_io_issuer_v1_manifest_test.go
	touch $@
out/terratest-sentinel-certificates_k8s_io_certificate_signing_request_v1_manifest_test.go: out/install-sentinel terratest/certificates_k8s_io_v1/certificates_k8s_io_certificate_signing_request_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_certificates_k8s_io_certificate_signing_request_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/certificates_k8s_io_v1/certificates_k8s_io_certificate_signing_request_v1_manifest_test.go
	touch $@
out/terratest-sentinel-chainsaw_kyverno_io_configuration_v1alpha1_manifest_test.go: out/install-sentinel terratest/chainsaw_kyverno_io_v1alpha1/chainsaw_kyverno_io_configuration_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_chainsaw_kyverno_io_configuration_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/chainsaw_kyverno_io_v1alpha1/chainsaw_kyverno_io_configuration_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-chainsaw_kyverno_io_test_v1alpha1_manifest_test.go: out/install-sentinel terratest/chainsaw_kyverno_io_v1alpha1/chainsaw_kyverno_io_test_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_chainsaw_kyverno_io_test_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/chainsaw_kyverno_io_v1alpha1/chainsaw_kyverno_io_test_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-chainsaw_kyverno_io_configuration_v1alpha2_manifest_test.go: out/install-sentinel terratest/chainsaw_kyverno_io_v1alpha2/chainsaw_kyverno_io_configuration_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_chainsaw_kyverno_io_configuration_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/chainsaw_kyverno_io_v1alpha2/chainsaw_kyverno_io_configuration_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-chainsaw_kyverno_io_test_v1alpha2_manifest_test.go: out/install-sentinel terratest/chainsaw_kyverno_io_v1alpha2/chainsaw_kyverno_io_test_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_chainsaw_kyverno_io_test_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/chainsaw_kyverno_io_v1alpha2/chainsaw_kyverno_io_test_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-chaos_mesh_org_aws_chaos_v1alpha1_manifest_test.go: out/install-sentinel terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_aws_chaos_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_chaos_mesh_org_aws_chaos_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_aws_chaos_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-chaos_mesh_org_azure_chaos_v1alpha1_manifest_test.go: out/install-sentinel terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_azure_chaos_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_chaos_mesh_org_azure_chaos_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_azure_chaos_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-chaos_mesh_org_block_chaos_v1alpha1_manifest_test.go: out/install-sentinel terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_block_chaos_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_chaos_mesh_org_block_chaos_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_block_chaos_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-chaos_mesh_org_dns_chaos_v1alpha1_manifest_test.go: out/install-sentinel terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_dns_chaos_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_chaos_mesh_org_dns_chaos_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_dns_chaos_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-chaos_mesh_org_gcp_chaos_v1alpha1_manifest_test.go: out/install-sentinel terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_gcp_chaos_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_chaos_mesh_org_gcp_chaos_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_gcp_chaos_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-chaos_mesh_org_http_chaos_v1alpha1_manifest_test.go: out/install-sentinel terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_http_chaos_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_chaos_mesh_org_http_chaos_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_http_chaos_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-chaos_mesh_org_io_chaos_v1alpha1_manifest_test.go: out/install-sentinel terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_io_chaos_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_chaos_mesh_org_io_chaos_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_io_chaos_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-chaos_mesh_org_jvm_chaos_v1alpha1_manifest_test.go: out/install-sentinel terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_jvm_chaos_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_chaos_mesh_org_jvm_chaos_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_jvm_chaos_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-chaos_mesh_org_kernel_chaos_v1alpha1_manifest_test.go: out/install-sentinel terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_kernel_chaos_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_chaos_mesh_org_kernel_chaos_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_kernel_chaos_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-chaos_mesh_org_network_chaos_v1alpha1_manifest_test.go: out/install-sentinel terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_network_chaos_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_chaos_mesh_org_network_chaos_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_network_chaos_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-chaos_mesh_org_physical_machine_chaos_v1alpha1_manifest_test.go: out/install-sentinel terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_physical_machine_chaos_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_chaos_mesh_org_physical_machine_chaos_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_physical_machine_chaos_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-chaos_mesh_org_physical_machine_v1alpha1_manifest_test.go: out/install-sentinel terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_physical_machine_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_chaos_mesh_org_physical_machine_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_physical_machine_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-chaos_mesh_org_pod_chaos_v1alpha1_manifest_test.go: out/install-sentinel terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_pod_chaos_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_chaos_mesh_org_pod_chaos_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_pod_chaos_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-chaos_mesh_org_pod_http_chaos_v1alpha1_manifest_test.go: out/install-sentinel terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_pod_http_chaos_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_chaos_mesh_org_pod_http_chaos_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_pod_http_chaos_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-chaos_mesh_org_pod_io_chaos_v1alpha1_manifest_test.go: out/install-sentinel terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_pod_io_chaos_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_chaos_mesh_org_pod_io_chaos_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_pod_io_chaos_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-chaos_mesh_org_pod_network_chaos_v1alpha1_manifest_test.go: out/install-sentinel terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_pod_network_chaos_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_chaos_mesh_org_pod_network_chaos_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_pod_network_chaos_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-chaos_mesh_org_remote_cluster_v1alpha1_manifest_test.go: out/install-sentinel terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_remote_cluster_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_chaos_mesh_org_remote_cluster_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_remote_cluster_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-chaos_mesh_org_schedule_v1alpha1_manifest_test.go: out/install-sentinel terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_schedule_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_chaos_mesh_org_schedule_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_schedule_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-chaos_mesh_org_status_check_v1alpha1_manifest_test.go: out/install-sentinel terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_status_check_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_chaos_mesh_org_status_check_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_status_check_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-chaos_mesh_org_stress_chaos_v1alpha1_manifest_test.go: out/install-sentinel terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_stress_chaos_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_chaos_mesh_org_stress_chaos_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_stress_chaos_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-chaos_mesh_org_time_chaos_v1alpha1_manifest_test.go: out/install-sentinel terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_time_chaos_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_chaos_mesh_org_time_chaos_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_time_chaos_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-chaos_mesh_org_workflow_node_v1alpha1_manifest_test.go: out/install-sentinel terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_workflow_node_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_chaos_mesh_org_workflow_node_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_workflow_node_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-chaos_mesh_org_workflow_v1alpha1_manifest_test.go: out/install-sentinel terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_workflow_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_chaos_mesh_org_workflow_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/chaos_mesh_org_v1alpha1/chaos_mesh_org_workflow_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-chaosblade_io_chaos_blade_v1alpha1_manifest_test.go: out/install-sentinel terratest/chaosblade_io_v1alpha1/chaosblade_io_chaos_blade_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_chaosblade_io_chaos_blade_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/chaosblade_io_v1alpha1/chaosblade_io_chaos_blade_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-charts_amd_com_amdgpu_v1alpha1_manifest_test.go: out/install-sentinel terratest/charts_amd_com_v1alpha1/charts_amd_com_amdgpu_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_charts_amd_com_amdgpu_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/charts_amd_com_v1alpha1/charts_amd_com_amdgpu_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-charts_flagsmith_com_flagsmith_v1alpha1_manifest_test.go: out/install-sentinel terratest/charts_flagsmith_com_v1alpha1/charts_flagsmith_com_flagsmith_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_charts_flagsmith_com_flagsmith_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/charts_flagsmith_com_v1alpha1/charts_flagsmith_com_flagsmith_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-charts_helm_k8s_io_snyk_monitor_v1alpha1_manifest_test.go: out/install-sentinel terratest/charts_helm_k8s_io_v1alpha1/charts_helm_k8s_io_snyk_monitor_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_charts_helm_k8s_io_snyk_monitor_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/charts_helm_k8s_io_v1alpha1/charts_helm_k8s_io_snyk_monitor_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-charts_opdev_io_synapse_v1alpha1_manifest_test.go: out/install-sentinel terratest/charts_opdev_io_v1alpha1/charts_opdev_io_synapse_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_charts_opdev_io_synapse_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/charts_opdev_io_v1alpha1/charts_opdev_io_synapse_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-charts_operatorhub_io_cockroachdb_v1alpha1_manifest_test.go: out/install-sentinel terratest/charts_operatorhub_io_v1alpha1/charts_operatorhub_io_cockroachdb_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_charts_operatorhub_io_cockroachdb_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/charts_operatorhub_io_v1alpha1/charts_operatorhub_io_cockroachdb_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-che_eclipse_org_kubernetes_image_puller_v1alpha1_manifest_test.go: out/install-sentinel terratest/che_eclipse_org_v1alpha1/che_eclipse_org_kubernetes_image_puller_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_che_eclipse_org_kubernetes_image_puller_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/che_eclipse_org_v1alpha1/che_eclipse_org_kubernetes_image_puller_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-chisel_operator_io_exit_node_provisioner_v1_manifest_test.go: out/install-sentinel terratest/chisel_operator_io_v1/chisel_operator_io_exit_node_provisioner_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_chisel_operator_io_exit_node_provisioner_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/chisel_operator_io_v1/chisel_operator_io_exit_node_provisioner_v1_manifest_test.go
	touch $@
out/terratest-sentinel-chisel_operator_io_exit_node_v1_manifest_test.go: out/install-sentinel terratest/chisel_operator_io_v1/chisel_operator_io_exit_node_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_chisel_operator_io_exit_node_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/chisel_operator_io_v1/chisel_operator_io_exit_node_v1_manifest_test.go
	touch $@
out/terratest-sentinel-chisel_operator_io_exit_node_v2_manifest_test.go: out/install-sentinel terratest/chisel_operator_io_v2/chisel_operator_io_exit_node_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_chisel_operator_io_exit_node_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/chisel_operator_io_v2/chisel_operator_io_exit_node_v2_manifest_test.go
	touch $@
out/terratest-sentinel-cilium_io_cilium_clusterwide_envoy_config_v2_manifest_test.go: out/install-sentinel terratest/cilium_io_v2/cilium_io_cilium_clusterwide_envoy_config_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_cilium_io_cilium_clusterwide_envoy_config_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cilium_io_v2/cilium_io_cilium_clusterwide_envoy_config_v2_manifest_test.go
	touch $@
out/terratest-sentinel-cilium_io_cilium_clusterwide_network_policy_v2_manifest_test.go: out/install-sentinel terratest/cilium_io_v2/cilium_io_cilium_clusterwide_network_policy_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_cilium_io_cilium_clusterwide_network_policy_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cilium_io_v2/cilium_io_cilium_clusterwide_network_policy_v2_manifest_test.go
	touch $@
out/terratest-sentinel-cilium_io_cilium_egress_gateway_policy_v2_manifest_test.go: out/install-sentinel terratest/cilium_io_v2/cilium_io_cilium_egress_gateway_policy_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_cilium_io_cilium_egress_gateway_policy_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cilium_io_v2/cilium_io_cilium_egress_gateway_policy_v2_manifest_test.go
	touch $@
out/terratest-sentinel-cilium_io_cilium_envoy_config_v2_manifest_test.go: out/install-sentinel terratest/cilium_io_v2/cilium_io_cilium_envoy_config_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_cilium_io_cilium_envoy_config_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cilium_io_v2/cilium_io_cilium_envoy_config_v2_manifest_test.go
	touch $@
out/terratest-sentinel-cilium_io_cilium_external_workload_v2_manifest_test.go: out/install-sentinel terratest/cilium_io_v2/cilium_io_cilium_external_workload_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_cilium_io_cilium_external_workload_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cilium_io_v2/cilium_io_cilium_external_workload_v2_manifest_test.go
	touch $@
out/terratest-sentinel-cilium_io_cilium_identity_v2_manifest_test.go: out/install-sentinel terratest/cilium_io_v2/cilium_io_cilium_identity_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_cilium_io_cilium_identity_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cilium_io_v2/cilium_io_cilium_identity_v2_manifest_test.go
	touch $@
out/terratest-sentinel-cilium_io_cilium_local_redirect_policy_v2_manifest_test.go: out/install-sentinel terratest/cilium_io_v2/cilium_io_cilium_local_redirect_policy_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_cilium_io_cilium_local_redirect_policy_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cilium_io_v2/cilium_io_cilium_local_redirect_policy_v2_manifest_test.go
	touch $@
out/terratest-sentinel-cilium_io_cilium_network_policy_v2_manifest_test.go: out/install-sentinel terratest/cilium_io_v2/cilium_io_cilium_network_policy_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_cilium_io_cilium_network_policy_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cilium_io_v2/cilium_io_cilium_network_policy_v2_manifest_test.go
	touch $@
out/terratest-sentinel-cilium_io_cilium_node_v2_manifest_test.go: out/install-sentinel terratest/cilium_io_v2/cilium_io_cilium_node_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_cilium_io_cilium_node_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cilium_io_v2/cilium_io_cilium_node_v2_manifest_test.go
	touch $@
out/terratest-sentinel-cilium_io_cilium_bgp_peering_policy_v2alpha1_manifest_test.go: out/install-sentinel terratest/cilium_io_v2alpha1/cilium_io_cilium_bgp_peering_policy_v2alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_cilium_io_cilium_bgp_peering_policy_v2alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cilium_io_v2alpha1/cilium_io_cilium_bgp_peering_policy_v2alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-cilium_io_cilium_cidr_group_v2alpha1_manifest_test.go: out/install-sentinel terratest/cilium_io_v2alpha1/cilium_io_cilium_cidr_group_v2alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_cilium_io_cilium_cidr_group_v2alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cilium_io_v2alpha1/cilium_io_cilium_cidr_group_v2alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-cilium_io_cilium_endpoint_slice_v2alpha1_manifest_test.go: out/install-sentinel terratest/cilium_io_v2alpha1/cilium_io_cilium_endpoint_slice_v2alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_cilium_io_cilium_endpoint_slice_v2alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cilium_io_v2alpha1/cilium_io_cilium_endpoint_slice_v2alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-cilium_io_cilium_l2_announcement_policy_v2alpha1_manifest_test.go: out/install-sentinel terratest/cilium_io_v2alpha1/cilium_io_cilium_l2_announcement_policy_v2alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_cilium_io_cilium_l2_announcement_policy_v2alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cilium_io_v2alpha1/cilium_io_cilium_l2_announcement_policy_v2alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-cilium_io_cilium_load_balancer_ip_pool_v2alpha1_manifest_test.go: out/install-sentinel terratest/cilium_io_v2alpha1/cilium_io_cilium_load_balancer_ip_pool_v2alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_cilium_io_cilium_load_balancer_ip_pool_v2alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cilium_io_v2alpha1/cilium_io_cilium_load_balancer_ip_pool_v2alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-cilium_io_cilium_node_config_v2alpha1_manifest_test.go: out/install-sentinel terratest/cilium_io_v2alpha1/cilium_io_cilium_node_config_v2alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_cilium_io_cilium_node_config_v2alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cilium_io_v2alpha1/cilium_io_cilium_node_config_v2alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-cilium_io_cilium_pod_ip_pool_v2alpha1_manifest_test.go: out/install-sentinel terratest/cilium_io_v2alpha1/cilium_io_cilium_pod_ip_pool_v2alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_cilium_io_cilium_pod_ip_pool_v2alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cilium_io_v2alpha1/cilium_io_cilium_pod_ip_pool_v2alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-claudie_io_input_manifest_v1beta1_manifest_test.go: out/install-sentinel terratest/claudie_io_v1beta1/claudie_io_input_manifest_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_claudie_io_input_manifest_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/claudie_io_v1beta1/claudie_io_input_manifest_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-cloudformation_linki_space_stack_v1alpha1_manifest_test.go: out/install-sentinel terratest/cloudformation_linki_space_v1alpha1/cloudformation_linki_space_stack_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_cloudformation_linki_space_stack_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cloudformation_linki_space_v1alpha1/cloudformation_linki_space_stack_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-cloudfront_services_k8s_aws_cache_policy_v1alpha1_manifest_test.go: out/install-sentinel terratest/cloudfront_services_k8s_aws_v1alpha1/cloudfront_services_k8s_aws_cache_policy_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_cloudfront_services_k8s_aws_cache_policy_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cloudfront_services_k8s_aws_v1alpha1/cloudfront_services_k8s_aws_cache_policy_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-cloudfront_services_k8s_aws_distribution_v1alpha1_manifest_test.go: out/install-sentinel terratest/cloudfront_services_k8s_aws_v1alpha1/cloudfront_services_k8s_aws_distribution_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_cloudfront_services_k8s_aws_distribution_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cloudfront_services_k8s_aws_v1alpha1/cloudfront_services_k8s_aws_distribution_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-cloudfront_services_k8s_aws_function_v1alpha1_manifest_test.go: out/install-sentinel terratest/cloudfront_services_k8s_aws_v1alpha1/cloudfront_services_k8s_aws_function_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_cloudfront_services_k8s_aws_function_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cloudfront_services_k8s_aws_v1alpha1/cloudfront_services_k8s_aws_function_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-cloudfront_services_k8s_aws_origin_request_policy_v1alpha1_manifest_test.go: out/install-sentinel terratest/cloudfront_services_k8s_aws_v1alpha1/cloudfront_services_k8s_aws_origin_request_policy_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_cloudfront_services_k8s_aws_origin_request_policy_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cloudfront_services_k8s_aws_v1alpha1/cloudfront_services_k8s_aws_origin_request_policy_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-cloudfront_services_k8s_aws_response_headers_policy_v1alpha1_manifest_test.go: out/install-sentinel terratest/cloudfront_services_k8s_aws_v1alpha1/cloudfront_services_k8s_aws_response_headers_policy_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_cloudfront_services_k8s_aws_response_headers_policy_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cloudfront_services_k8s_aws_v1alpha1/cloudfront_services_k8s_aws_response_headers_policy_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-cloudtrail_services_k8s_aws_event_data_store_v1alpha1_manifest_test.go: out/install-sentinel terratest/cloudtrail_services_k8s_aws_v1alpha1/cloudtrail_services_k8s_aws_event_data_store_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_cloudtrail_services_k8s_aws_event_data_store_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cloudtrail_services_k8s_aws_v1alpha1/cloudtrail_services_k8s_aws_event_data_store_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-cloudtrail_services_k8s_aws_trail_v1alpha1_manifest_test.go: out/install-sentinel terratest/cloudtrail_services_k8s_aws_v1alpha1/cloudtrail_services_k8s_aws_trail_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_cloudtrail_services_k8s_aws_trail_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cloudtrail_services_k8s_aws_v1alpha1/cloudtrail_services_k8s_aws_trail_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-cloudwatch_aws_amazon_com_amazon_cloud_watch_agent_v1alpha1_manifest_test.go: out/install-sentinel terratest/cloudwatch_aws_amazon_com_v1alpha1/cloudwatch_aws_amazon_com_amazon_cloud_watch_agent_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_cloudwatch_aws_amazon_com_amazon_cloud_watch_agent_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cloudwatch_aws_amazon_com_v1alpha1/cloudwatch_aws_amazon_com_amazon_cloud_watch_agent_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-cloudwatch_aws_amazon_com_instrumentation_v1alpha1_manifest_test.go: out/install-sentinel terratest/cloudwatch_aws_amazon_com_v1alpha1/cloudwatch_aws_amazon_com_instrumentation_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_cloudwatch_aws_amazon_com_instrumentation_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cloudwatch_aws_amazon_com_v1alpha1/cloudwatch_aws_amazon_com_instrumentation_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-cloudwatch_services_k8s_aws_metric_alarm_v1alpha1_manifest_test.go: out/install-sentinel terratest/cloudwatch_services_k8s_aws_v1alpha1/cloudwatch_services_k8s_aws_metric_alarm_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_cloudwatch_services_k8s_aws_metric_alarm_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cloudwatch_services_k8s_aws_v1alpha1/cloudwatch_services_k8s_aws_metric_alarm_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-cloudwatchlogs_services_k8s_aws_log_group_v1alpha1_manifest_test.go: out/install-sentinel terratest/cloudwatchlogs_services_k8s_aws_v1alpha1/cloudwatchlogs_services_k8s_aws_log_group_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_cloudwatchlogs_services_k8s_aws_log_group_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cloudwatchlogs_services_k8s_aws_v1alpha1/cloudwatchlogs_services_k8s_aws_log_group_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-cluster_clusterpedia_io_cluster_sync_resources_v1alpha2_manifest_test.go: out/install-sentinel terratest/cluster_clusterpedia_io_v1alpha2/cluster_clusterpedia_io_cluster_sync_resources_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_cluster_clusterpedia_io_cluster_sync_resources_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cluster_clusterpedia_io_v1alpha2/cluster_clusterpedia_io_cluster_sync_resources_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-cluster_clusterpedia_io_pedia_cluster_v1alpha2_manifest_test.go: out/install-sentinel terratest/cluster_clusterpedia_io_v1alpha2/cluster_clusterpedia_io_pedia_cluster_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_cluster_clusterpedia_io_pedia_cluster_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cluster_clusterpedia_io_v1alpha2/cluster_clusterpedia_io_pedia_cluster_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-cluster_ipfs_io_circuit_relay_v1alpha1_manifest_test.go: out/install-sentinel terratest/cluster_ipfs_io_v1alpha1/cluster_ipfs_io_circuit_relay_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_cluster_ipfs_io_circuit_relay_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cluster_ipfs_io_v1alpha1/cluster_ipfs_io_circuit_relay_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-cluster_ipfs_io_ipfs_cluster_v1alpha1_manifest_test.go: out/install-sentinel terratest/cluster_ipfs_io_v1alpha1/cluster_ipfs_io_ipfs_cluster_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_cluster_ipfs_io_ipfs_cluster_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cluster_ipfs_io_v1alpha1/cluster_ipfs_io_ipfs_cluster_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-cluster_x_k8s_io_cluster_v1alpha3_manifest_test.go: out/install-sentinel terratest/cluster_x_k8s_io_v1alpha3/cluster_x_k8s_io_cluster_v1alpha3_manifest_test.go $(shell find ./examples/data-sources/k8s_cluster_x_k8s_io_cluster_v1alpha3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cluster_x_k8s_io_v1alpha3/cluster_x_k8s_io_cluster_v1alpha3_manifest_test.go
	touch $@
out/terratest-sentinel-cluster_x_k8s_io_machine_deployment_v1alpha3_manifest_test.go: out/install-sentinel terratest/cluster_x_k8s_io_v1alpha3/cluster_x_k8s_io_machine_deployment_v1alpha3_manifest_test.go $(shell find ./examples/data-sources/k8s_cluster_x_k8s_io_machine_deployment_v1alpha3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cluster_x_k8s_io_v1alpha3/cluster_x_k8s_io_machine_deployment_v1alpha3_manifest_test.go
	touch $@
out/terratest-sentinel-cluster_x_k8s_io_machine_health_check_v1alpha3_manifest_test.go: out/install-sentinel terratest/cluster_x_k8s_io_v1alpha3/cluster_x_k8s_io_machine_health_check_v1alpha3_manifest_test.go $(shell find ./examples/data-sources/k8s_cluster_x_k8s_io_machine_health_check_v1alpha3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cluster_x_k8s_io_v1alpha3/cluster_x_k8s_io_machine_health_check_v1alpha3_manifest_test.go
	touch $@
out/terratest-sentinel-cluster_x_k8s_io_machine_pool_v1alpha3_manifest_test.go: out/install-sentinel terratest/cluster_x_k8s_io_v1alpha3/cluster_x_k8s_io_machine_pool_v1alpha3_manifest_test.go $(shell find ./examples/data-sources/k8s_cluster_x_k8s_io_machine_pool_v1alpha3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cluster_x_k8s_io_v1alpha3/cluster_x_k8s_io_machine_pool_v1alpha3_manifest_test.go
	touch $@
out/terratest-sentinel-cluster_x_k8s_io_machine_set_v1alpha3_manifest_test.go: out/install-sentinel terratest/cluster_x_k8s_io_v1alpha3/cluster_x_k8s_io_machine_set_v1alpha3_manifest_test.go $(shell find ./examples/data-sources/k8s_cluster_x_k8s_io_machine_set_v1alpha3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cluster_x_k8s_io_v1alpha3/cluster_x_k8s_io_machine_set_v1alpha3_manifest_test.go
	touch $@
out/terratest-sentinel-cluster_x_k8s_io_machine_v1alpha3_manifest_test.go: out/install-sentinel terratest/cluster_x_k8s_io_v1alpha3/cluster_x_k8s_io_machine_v1alpha3_manifest_test.go $(shell find ./examples/data-sources/k8s_cluster_x_k8s_io_machine_v1alpha3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cluster_x_k8s_io_v1alpha3/cluster_x_k8s_io_machine_v1alpha3_manifest_test.go
	touch $@
out/terratest-sentinel-cluster_x_k8s_io_cluster_class_v1alpha4_manifest_test.go: out/install-sentinel terratest/cluster_x_k8s_io_v1alpha4/cluster_x_k8s_io_cluster_class_v1alpha4_manifest_test.go $(shell find ./examples/data-sources/k8s_cluster_x_k8s_io_cluster_class_v1alpha4_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cluster_x_k8s_io_v1alpha4/cluster_x_k8s_io_cluster_class_v1alpha4_manifest_test.go
	touch $@
out/terratest-sentinel-cluster_x_k8s_io_cluster_v1alpha4_manifest_test.go: out/install-sentinel terratest/cluster_x_k8s_io_v1alpha4/cluster_x_k8s_io_cluster_v1alpha4_manifest_test.go $(shell find ./examples/data-sources/k8s_cluster_x_k8s_io_cluster_v1alpha4_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cluster_x_k8s_io_v1alpha4/cluster_x_k8s_io_cluster_v1alpha4_manifest_test.go
	touch $@
out/terratest-sentinel-cluster_x_k8s_io_machine_deployment_v1alpha4_manifest_test.go: out/install-sentinel terratest/cluster_x_k8s_io_v1alpha4/cluster_x_k8s_io_machine_deployment_v1alpha4_manifest_test.go $(shell find ./examples/data-sources/k8s_cluster_x_k8s_io_machine_deployment_v1alpha4_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cluster_x_k8s_io_v1alpha4/cluster_x_k8s_io_machine_deployment_v1alpha4_manifest_test.go
	touch $@
out/terratest-sentinel-cluster_x_k8s_io_machine_health_check_v1alpha4_manifest_test.go: out/install-sentinel terratest/cluster_x_k8s_io_v1alpha4/cluster_x_k8s_io_machine_health_check_v1alpha4_manifest_test.go $(shell find ./examples/data-sources/k8s_cluster_x_k8s_io_machine_health_check_v1alpha4_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cluster_x_k8s_io_v1alpha4/cluster_x_k8s_io_machine_health_check_v1alpha4_manifest_test.go
	touch $@
out/terratest-sentinel-cluster_x_k8s_io_machine_pool_v1alpha4_manifest_test.go: out/install-sentinel terratest/cluster_x_k8s_io_v1alpha4/cluster_x_k8s_io_machine_pool_v1alpha4_manifest_test.go $(shell find ./examples/data-sources/k8s_cluster_x_k8s_io_machine_pool_v1alpha4_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cluster_x_k8s_io_v1alpha4/cluster_x_k8s_io_machine_pool_v1alpha4_manifest_test.go
	touch $@
out/terratest-sentinel-cluster_x_k8s_io_machine_set_v1alpha4_manifest_test.go: out/install-sentinel terratest/cluster_x_k8s_io_v1alpha4/cluster_x_k8s_io_machine_set_v1alpha4_manifest_test.go $(shell find ./examples/data-sources/k8s_cluster_x_k8s_io_machine_set_v1alpha4_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cluster_x_k8s_io_v1alpha4/cluster_x_k8s_io_machine_set_v1alpha4_manifest_test.go
	touch $@
out/terratest-sentinel-cluster_x_k8s_io_machine_v1alpha4_manifest_test.go: out/install-sentinel terratest/cluster_x_k8s_io_v1alpha4/cluster_x_k8s_io_machine_v1alpha4_manifest_test.go $(shell find ./examples/data-sources/k8s_cluster_x_k8s_io_machine_v1alpha4_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cluster_x_k8s_io_v1alpha4/cluster_x_k8s_io_machine_v1alpha4_manifest_test.go
	touch $@
out/terratest-sentinel-cluster_x_k8s_io_cluster_class_v1beta1_manifest_test.go: out/install-sentinel terratest/cluster_x_k8s_io_v1beta1/cluster_x_k8s_io_cluster_class_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_cluster_x_k8s_io_cluster_class_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cluster_x_k8s_io_v1beta1/cluster_x_k8s_io_cluster_class_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-cluster_x_k8s_io_cluster_v1beta1_manifest_test.go: out/install-sentinel terratest/cluster_x_k8s_io_v1beta1/cluster_x_k8s_io_cluster_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_cluster_x_k8s_io_cluster_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cluster_x_k8s_io_v1beta1/cluster_x_k8s_io_cluster_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-cluster_x_k8s_io_machine_deployment_v1beta1_manifest_test.go: out/install-sentinel terratest/cluster_x_k8s_io_v1beta1/cluster_x_k8s_io_machine_deployment_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_cluster_x_k8s_io_machine_deployment_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cluster_x_k8s_io_v1beta1/cluster_x_k8s_io_machine_deployment_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-cluster_x_k8s_io_machine_health_check_v1beta1_manifest_test.go: out/install-sentinel terratest/cluster_x_k8s_io_v1beta1/cluster_x_k8s_io_machine_health_check_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_cluster_x_k8s_io_machine_health_check_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cluster_x_k8s_io_v1beta1/cluster_x_k8s_io_machine_health_check_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-cluster_x_k8s_io_machine_pool_v1beta1_manifest_test.go: out/install-sentinel terratest/cluster_x_k8s_io_v1beta1/cluster_x_k8s_io_machine_pool_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_cluster_x_k8s_io_machine_pool_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cluster_x_k8s_io_v1beta1/cluster_x_k8s_io_machine_pool_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-cluster_x_k8s_io_machine_set_v1beta1_manifest_test.go: out/install-sentinel terratest/cluster_x_k8s_io_v1beta1/cluster_x_k8s_io_machine_set_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_cluster_x_k8s_io_machine_set_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cluster_x_k8s_io_v1beta1/cluster_x_k8s_io_machine_set_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-cluster_x_k8s_io_machine_v1beta1_manifest_test.go: out/install-sentinel terratest/cluster_x_k8s_io_v1beta1/cluster_x_k8s_io_machine_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_cluster_x_k8s_io_machine_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/cluster_x_k8s_io_v1beta1/cluster_x_k8s_io_machine_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-clusters_clusternet_io_cluster_registration_request_v1beta1_manifest_test.go: out/install-sentinel terratest/clusters_clusternet_io_v1beta1/clusters_clusternet_io_cluster_registration_request_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_clusters_clusternet_io_cluster_registration_request_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/clusters_clusternet_io_v1beta1/clusters_clusternet_io_cluster_registration_request_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-clusters_clusternet_io_managed_cluster_v1beta1_manifest_test.go: out/install-sentinel terratest/clusters_clusternet_io_v1beta1/clusters_clusternet_io_managed_cluster_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_clusters_clusternet_io_managed_cluster_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/clusters_clusternet_io_v1beta1/clusters_clusternet_io_managed_cluster_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-clustertemplate_openshift_io_cluster_template_instance_v1alpha1_manifest_test.go: out/install-sentinel terratest/clustertemplate_openshift_io_v1alpha1/clustertemplate_openshift_io_cluster_template_instance_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_clustertemplate_openshift_io_cluster_template_instance_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/clustertemplate_openshift_io_v1alpha1/clustertemplate_openshift_io_cluster_template_instance_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-clustertemplate_openshift_io_cluster_template_quota_v1alpha1_manifest_test.go: out/install-sentinel terratest/clustertemplate_openshift_io_v1alpha1/clustertemplate_openshift_io_cluster_template_quota_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_clustertemplate_openshift_io_cluster_template_quota_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/clustertemplate_openshift_io_v1alpha1/clustertemplate_openshift_io_cluster_template_quota_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-clustertemplate_openshift_io_cluster_template_setup_v1alpha1_manifest_test.go: out/install-sentinel terratest/clustertemplate_openshift_io_v1alpha1/clustertemplate_openshift_io_cluster_template_setup_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_clustertemplate_openshift_io_cluster_template_setup_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/clustertemplate_openshift_io_v1alpha1/clustertemplate_openshift_io_cluster_template_setup_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-clustertemplate_openshift_io_cluster_template_v1alpha1_manifest_test.go: out/install-sentinel terratest/clustertemplate_openshift_io_v1alpha1/clustertemplate_openshift_io_cluster_template_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_clustertemplate_openshift_io_cluster_template_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/clustertemplate_openshift_io_v1alpha1/clustertemplate_openshift_io_cluster_template_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-clustertemplate_openshift_io_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/clustertemplate_openshift_io_v1alpha1/clustertemplate_openshift_io_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_clustertemplate_openshift_io_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/clustertemplate_openshift_io_v1alpha1/clustertemplate_openshift_io_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-confidentialcontainers_org_cc_runtime_v1beta1_manifest_test.go: out/install-sentinel terratest/confidentialcontainers_org_v1beta1/confidentialcontainers_org_cc_runtime_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_confidentialcontainers_org_cc_runtime_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/confidentialcontainers_org_v1beta1/confidentialcontainers_org_cc_runtime_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-config_gatekeeper_sh_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/config_gatekeeper_sh_v1alpha1/config_gatekeeper_sh_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_config_gatekeeper_sh_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/config_gatekeeper_sh_v1alpha1/config_gatekeeper_sh_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-config_grafana_com_project_config_v1_manifest_test.go: out/install-sentinel terratest/config_grafana_com_v1/config_grafana_com_project_config_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_config_grafana_com_project_config_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/config_grafana_com_v1/config_grafana_com_project_config_v1_manifest_test.go
	touch $@
out/terratest-sentinel-config_karmada_io_resource_interpreter_customization_v1alpha1_manifest_test.go: out/install-sentinel terratest/config_karmada_io_v1alpha1/config_karmada_io_resource_interpreter_customization_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_config_karmada_io_resource_interpreter_customization_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/config_karmada_io_v1alpha1/config_karmada_io_resource_interpreter_customization_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-config_karmada_io_resource_interpreter_webhook_configuration_v1alpha1_manifest_test.go: out/install-sentinel terratest/config_karmada_io_v1alpha1/config_karmada_io_resource_interpreter_webhook_configuration_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_config_karmada_io_resource_interpreter_webhook_configuration_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/config_karmada_io_v1alpha1/config_karmada_io_resource_interpreter_webhook_configuration_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-config_koordinator_sh_cluster_colocation_profile_v1alpha1_manifest_test.go: out/install-sentinel terratest/config_koordinator_sh_v1alpha1/config_koordinator_sh_cluster_colocation_profile_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_config_koordinator_sh_cluster_colocation_profile_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/config_koordinator_sh_v1alpha1/config_koordinator_sh_cluster_colocation_profile_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-config_storageos_com_operator_config_v1_manifest_test.go: out/install-sentinel terratest/config_storageos_com_v1/config_storageos_com_operator_config_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_config_storageos_com_operator_config_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/config_storageos_com_v1/config_storageos_com_operator_config_v1_manifest_test.go
	touch $@
out/terratest-sentinel-control_k8ssandra_io_cassandra_task_v1alpha1_manifest_test.go: out/install-sentinel terratest/control_k8ssandra_io_v1alpha1/control_k8ssandra_io_cassandra_task_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_control_k8ssandra_io_cassandra_task_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/control_k8ssandra_io_v1alpha1/control_k8ssandra_io_cassandra_task_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-core_kubeadmiral_io_cluster_collected_status_v1alpha1_manifest_test.go: out/install-sentinel terratest/core_kubeadmiral_io_v1alpha1/core_kubeadmiral_io_cluster_collected_status_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_core_kubeadmiral_io_cluster_collected_status_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/core_kubeadmiral_io_v1alpha1/core_kubeadmiral_io_cluster_collected_status_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-core_kubeadmiral_io_cluster_federated_object_v1alpha1_manifest_test.go: out/install-sentinel terratest/core_kubeadmiral_io_v1alpha1/core_kubeadmiral_io_cluster_federated_object_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_core_kubeadmiral_io_cluster_federated_object_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/core_kubeadmiral_io_v1alpha1/core_kubeadmiral_io_cluster_federated_object_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-core_kubeadmiral_io_cluster_override_policy_v1alpha1_manifest_test.go: out/install-sentinel terratest/core_kubeadmiral_io_v1alpha1/core_kubeadmiral_io_cluster_override_policy_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_core_kubeadmiral_io_cluster_override_policy_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/core_kubeadmiral_io_v1alpha1/core_kubeadmiral_io_cluster_override_policy_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-core_kubeadmiral_io_cluster_propagation_policy_v1alpha1_manifest_test.go: out/install-sentinel terratest/core_kubeadmiral_io_v1alpha1/core_kubeadmiral_io_cluster_propagation_policy_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_core_kubeadmiral_io_cluster_propagation_policy_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/core_kubeadmiral_io_v1alpha1/core_kubeadmiral_io_cluster_propagation_policy_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-core_kubeadmiral_io_collected_status_v1alpha1_manifest_test.go: out/install-sentinel terratest/core_kubeadmiral_io_v1alpha1/core_kubeadmiral_io_collected_status_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_core_kubeadmiral_io_collected_status_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/core_kubeadmiral_io_v1alpha1/core_kubeadmiral_io_collected_status_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-core_kubeadmiral_io_federated_cluster_v1alpha1_manifest_test.go: out/install-sentinel terratest/core_kubeadmiral_io_v1alpha1/core_kubeadmiral_io_federated_cluster_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_core_kubeadmiral_io_federated_cluster_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/core_kubeadmiral_io_v1alpha1/core_kubeadmiral_io_federated_cluster_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-core_kubeadmiral_io_federated_object_v1alpha1_manifest_test.go: out/install-sentinel terratest/core_kubeadmiral_io_v1alpha1/core_kubeadmiral_io_federated_object_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_core_kubeadmiral_io_federated_object_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/core_kubeadmiral_io_v1alpha1/core_kubeadmiral_io_federated_object_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-core_kubeadmiral_io_federated_type_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/core_kubeadmiral_io_v1alpha1/core_kubeadmiral_io_federated_type_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_core_kubeadmiral_io_federated_type_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/core_kubeadmiral_io_v1alpha1/core_kubeadmiral_io_federated_type_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-core_kubeadmiral_io_override_policy_v1alpha1_manifest_test.go: out/install-sentinel terratest/core_kubeadmiral_io_v1alpha1/core_kubeadmiral_io_override_policy_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_core_kubeadmiral_io_override_policy_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/core_kubeadmiral_io_v1alpha1/core_kubeadmiral_io_override_policy_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-core_kubeadmiral_io_propagation_policy_v1alpha1_manifest_test.go: out/install-sentinel terratest/core_kubeadmiral_io_v1alpha1/core_kubeadmiral_io_propagation_policy_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_core_kubeadmiral_io_propagation_policy_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/core_kubeadmiral_io_v1alpha1/core_kubeadmiral_io_propagation_policy_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-core_kubeadmiral_io_scheduler_plugin_webhook_configuration_v1alpha1_manifest_test.go: out/install-sentinel terratest/core_kubeadmiral_io_v1alpha1/core_kubeadmiral_io_scheduler_plugin_webhook_configuration_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_core_kubeadmiral_io_scheduler_plugin_webhook_configuration_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/core_kubeadmiral_io_v1alpha1/core_kubeadmiral_io_scheduler_plugin_webhook_configuration_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-core_kubeadmiral_io_scheduling_profile_v1alpha1_manifest_test.go: out/install-sentinel terratest/core_kubeadmiral_io_v1alpha1/core_kubeadmiral_io_scheduling_profile_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_core_kubeadmiral_io_scheduling_profile_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/core_kubeadmiral_io_v1alpha1/core_kubeadmiral_io_scheduling_profile_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-core_linuxsuren_github_com_a_test_v1alpha1_manifest_test.go: out/install-sentinel terratest/core_linuxsuren_github_com_v1alpha1/core_linuxsuren_github_com_a_test_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_core_linuxsuren_github_com_a_test_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/core_linuxsuren_github_com_v1alpha1/core_linuxsuren_github_com_a_test_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-core_openfeature_dev_feature_flag_configuration_v1alpha1_manifest_test.go: out/install-sentinel terratest/core_openfeature_dev_v1alpha1/core_openfeature_dev_feature_flag_configuration_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_core_openfeature_dev_feature_flag_configuration_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/core_openfeature_dev_v1alpha1/core_openfeature_dev_feature_flag_configuration_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-core_openfeature_dev_feature_flag_configuration_v1alpha2_manifest_test.go: out/install-sentinel terratest/core_openfeature_dev_v1alpha2/core_openfeature_dev_feature_flag_configuration_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_core_openfeature_dev_feature_flag_configuration_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/core_openfeature_dev_v1alpha2/core_openfeature_dev_feature_flag_configuration_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-core_strimzi_io_strimzi_pod_set_v1beta2_manifest_test.go: out/install-sentinel terratest/core_strimzi_io_v1beta2/core_strimzi_io_strimzi_pod_set_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_core_strimzi_io_strimzi_pod_set_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/core_strimzi_io_v1beta2/core_strimzi_io_strimzi_pod_set_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-config_map_v1_manifest_test.go: out/install-sentinel terratest/core_v1/config_map_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_config_map_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/core_v1/config_map_v1_manifest_test.go
	touch $@
out/terratest-sentinel-endpoints_v1_manifest_test.go: out/install-sentinel terratest/core_v1/endpoints_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_endpoints_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/core_v1/endpoints_v1_manifest_test.go
	touch $@
out/terratest-sentinel-limit_range_v1_manifest_test.go: out/install-sentinel terratest/core_v1/limit_range_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_limit_range_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/core_v1/limit_range_v1_manifest_test.go
	touch $@
out/terratest-sentinel-namespace_v1_manifest_test.go: out/install-sentinel terratest/core_v1/namespace_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_namespace_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/core_v1/namespace_v1_manifest_test.go
	touch $@
out/terratest-sentinel-persistent_volume_claim_v1_manifest_test.go: out/install-sentinel terratest/core_v1/persistent_volume_claim_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_persistent_volume_claim_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/core_v1/persistent_volume_claim_v1_manifest_test.go
	touch $@
out/terratest-sentinel-persistent_volume_v1_manifest_test.go: out/install-sentinel terratest/core_v1/persistent_volume_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_persistent_volume_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/core_v1/persistent_volume_v1_manifest_test.go
	touch $@
out/terratest-sentinel-pod_v1_manifest_test.go: out/install-sentinel terratest/core_v1/pod_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_pod_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/core_v1/pod_v1_manifest_test.go
	touch $@
out/terratest-sentinel-replication_controller_v1_manifest_test.go: out/install-sentinel terratest/core_v1/replication_controller_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_replication_controller_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/core_v1/replication_controller_v1_manifest_test.go
	touch $@
out/terratest-sentinel-secret_v1_manifest_test.go: out/install-sentinel terratest/core_v1/secret_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_secret_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/core_v1/secret_v1_manifest_test.go
	touch $@
out/terratest-sentinel-service_account_v1_manifest_test.go: out/install-sentinel terratest/core_v1/service_account_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_service_account_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/core_v1/service_account_v1_manifest_test.go
	touch $@
out/terratest-sentinel-service_v1_manifest_test.go: out/install-sentinel terratest/core_v1/service_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_service_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/core_v1/service_v1_manifest_test.go
	touch $@
out/terratest-sentinel-couchbase_com_couchbase_autoscaler_v2_manifest_test.go: out/install-sentinel terratest/couchbase_com_v2/couchbase_com_couchbase_autoscaler_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_couchbase_com_couchbase_autoscaler_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/couchbase_com_v2/couchbase_com_couchbase_autoscaler_v2_manifest_test.go
	touch $@
out/terratest-sentinel-couchbase_com_couchbase_backup_restore_v2_manifest_test.go: out/install-sentinel terratest/couchbase_com_v2/couchbase_com_couchbase_backup_restore_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_couchbase_com_couchbase_backup_restore_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/couchbase_com_v2/couchbase_com_couchbase_backup_restore_v2_manifest_test.go
	touch $@
out/terratest-sentinel-couchbase_com_couchbase_backup_v2_manifest_test.go: out/install-sentinel terratest/couchbase_com_v2/couchbase_com_couchbase_backup_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_couchbase_com_couchbase_backup_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/couchbase_com_v2/couchbase_com_couchbase_backup_v2_manifest_test.go
	touch $@
out/terratest-sentinel-couchbase_com_couchbase_bucket_v2_manifest_test.go: out/install-sentinel terratest/couchbase_com_v2/couchbase_com_couchbase_bucket_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_couchbase_com_couchbase_bucket_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/couchbase_com_v2/couchbase_com_couchbase_bucket_v2_manifest_test.go
	touch $@
out/terratest-sentinel-couchbase_com_couchbase_cluster_v2_manifest_test.go: out/install-sentinel terratest/couchbase_com_v2/couchbase_com_couchbase_cluster_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_couchbase_com_couchbase_cluster_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/couchbase_com_v2/couchbase_com_couchbase_cluster_v2_manifest_test.go
	touch $@
out/terratest-sentinel-couchbase_com_couchbase_collection_group_v2_manifest_test.go: out/install-sentinel terratest/couchbase_com_v2/couchbase_com_couchbase_collection_group_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_couchbase_com_couchbase_collection_group_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/couchbase_com_v2/couchbase_com_couchbase_collection_group_v2_manifest_test.go
	touch $@
out/terratest-sentinel-couchbase_com_couchbase_collection_v2_manifest_test.go: out/install-sentinel terratest/couchbase_com_v2/couchbase_com_couchbase_collection_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_couchbase_com_couchbase_collection_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/couchbase_com_v2/couchbase_com_couchbase_collection_v2_manifest_test.go
	touch $@
out/terratest-sentinel-couchbase_com_couchbase_ephemeral_bucket_v2_manifest_test.go: out/install-sentinel terratest/couchbase_com_v2/couchbase_com_couchbase_ephemeral_bucket_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_couchbase_com_couchbase_ephemeral_bucket_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/couchbase_com_v2/couchbase_com_couchbase_ephemeral_bucket_v2_manifest_test.go
	touch $@
out/terratest-sentinel-couchbase_com_couchbase_group_v2_manifest_test.go: out/install-sentinel terratest/couchbase_com_v2/couchbase_com_couchbase_group_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_couchbase_com_couchbase_group_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/couchbase_com_v2/couchbase_com_couchbase_group_v2_manifest_test.go
	touch $@
out/terratest-sentinel-couchbase_com_couchbase_memcached_bucket_v2_manifest_test.go: out/install-sentinel terratest/couchbase_com_v2/couchbase_com_couchbase_memcached_bucket_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_couchbase_com_couchbase_memcached_bucket_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/couchbase_com_v2/couchbase_com_couchbase_memcached_bucket_v2_manifest_test.go
	touch $@
out/terratest-sentinel-couchbase_com_couchbase_migration_replication_v2_manifest_test.go: out/install-sentinel terratest/couchbase_com_v2/couchbase_com_couchbase_migration_replication_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_couchbase_com_couchbase_migration_replication_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/couchbase_com_v2/couchbase_com_couchbase_migration_replication_v2_manifest_test.go
	touch $@
out/terratest-sentinel-couchbase_com_couchbase_replication_v2_manifest_test.go: out/install-sentinel terratest/couchbase_com_v2/couchbase_com_couchbase_replication_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_couchbase_com_couchbase_replication_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/couchbase_com_v2/couchbase_com_couchbase_replication_v2_manifest_test.go
	touch $@
out/terratest-sentinel-couchbase_com_couchbase_role_binding_v2_manifest_test.go: out/install-sentinel terratest/couchbase_com_v2/couchbase_com_couchbase_role_binding_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_couchbase_com_couchbase_role_binding_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/couchbase_com_v2/couchbase_com_couchbase_role_binding_v2_manifest_test.go
	touch $@
out/terratest-sentinel-couchbase_com_couchbase_scope_group_v2_manifest_test.go: out/install-sentinel terratest/couchbase_com_v2/couchbase_com_couchbase_scope_group_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_couchbase_com_couchbase_scope_group_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/couchbase_com_v2/couchbase_com_couchbase_scope_group_v2_manifest_test.go
	touch $@
out/terratest-sentinel-couchbase_com_couchbase_scope_v2_manifest_test.go: out/install-sentinel terratest/couchbase_com_v2/couchbase_com_couchbase_scope_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_couchbase_com_couchbase_scope_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/couchbase_com_v2/couchbase_com_couchbase_scope_v2_manifest_test.go
	touch $@
out/terratest-sentinel-couchbase_com_couchbase_user_v2_manifest_test.go: out/install-sentinel terratest/couchbase_com_v2/couchbase_com_couchbase_user_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_couchbase_com_couchbase_user_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/couchbase_com_v2/couchbase_com_couchbase_user_v2_manifest_test.go
	touch $@
out/terratest-sentinel-craftypath_github_io_sops_secret_v1alpha1_manifest_test.go: out/install-sentinel terratest/craftypath_github_io_v1alpha1/craftypath_github_io_sops_secret_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_craftypath_github_io_sops_secret_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/craftypath_github_io_v1alpha1/craftypath_github_io_sops_secret_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-crane_konveyor_io_operator_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/crane_konveyor_io_v1alpha1/crane_konveyor_io_operator_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_crane_konveyor_io_operator_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/crane_konveyor_io_v1alpha1/crane_konveyor_io_operator_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-crd_projectcalico_org_bgp_configuration_v1_manifest_test.go: out/install-sentinel terratest/crd_projectcalico_org_v1/crd_projectcalico_org_bgp_configuration_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_crd_projectcalico_org_bgp_configuration_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/crd_projectcalico_org_v1/crd_projectcalico_org_bgp_configuration_v1_manifest_test.go
	touch $@
out/terratest-sentinel-crd_projectcalico_org_bgp_filter_v1_manifest_test.go: out/install-sentinel terratest/crd_projectcalico_org_v1/crd_projectcalico_org_bgp_filter_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_crd_projectcalico_org_bgp_filter_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/crd_projectcalico_org_v1/crd_projectcalico_org_bgp_filter_v1_manifest_test.go
	touch $@
out/terratest-sentinel-crd_projectcalico_org_bgp_peer_v1_manifest_test.go: out/install-sentinel terratest/crd_projectcalico_org_v1/crd_projectcalico_org_bgp_peer_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_crd_projectcalico_org_bgp_peer_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/crd_projectcalico_org_v1/crd_projectcalico_org_bgp_peer_v1_manifest_test.go
	touch $@
out/terratest-sentinel-crd_projectcalico_org_block_affinity_v1_manifest_test.go: out/install-sentinel terratest/crd_projectcalico_org_v1/crd_projectcalico_org_block_affinity_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_crd_projectcalico_org_block_affinity_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/crd_projectcalico_org_v1/crd_projectcalico_org_block_affinity_v1_manifest_test.go
	touch $@
out/terratest-sentinel-crd_projectcalico_org_calico_node_status_v1_manifest_test.go: out/install-sentinel terratest/crd_projectcalico_org_v1/crd_projectcalico_org_calico_node_status_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_crd_projectcalico_org_calico_node_status_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/crd_projectcalico_org_v1/crd_projectcalico_org_calico_node_status_v1_manifest_test.go
	touch $@
out/terratest-sentinel-crd_projectcalico_org_cluster_information_v1_manifest_test.go: out/install-sentinel terratest/crd_projectcalico_org_v1/crd_projectcalico_org_cluster_information_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_crd_projectcalico_org_cluster_information_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/crd_projectcalico_org_v1/crd_projectcalico_org_cluster_information_v1_manifest_test.go
	touch $@
out/terratest-sentinel-crd_projectcalico_org_felix_configuration_v1_manifest_test.go: out/install-sentinel terratest/crd_projectcalico_org_v1/crd_projectcalico_org_felix_configuration_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_crd_projectcalico_org_felix_configuration_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/crd_projectcalico_org_v1/crd_projectcalico_org_felix_configuration_v1_manifest_test.go
	touch $@
out/terratest-sentinel-crd_projectcalico_org_global_network_policy_v1_manifest_test.go: out/install-sentinel terratest/crd_projectcalico_org_v1/crd_projectcalico_org_global_network_policy_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_crd_projectcalico_org_global_network_policy_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/crd_projectcalico_org_v1/crd_projectcalico_org_global_network_policy_v1_manifest_test.go
	touch $@
out/terratest-sentinel-crd_projectcalico_org_global_network_set_v1_manifest_test.go: out/install-sentinel terratest/crd_projectcalico_org_v1/crd_projectcalico_org_global_network_set_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_crd_projectcalico_org_global_network_set_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/crd_projectcalico_org_v1/crd_projectcalico_org_global_network_set_v1_manifest_test.go
	touch $@
out/terratest-sentinel-crd_projectcalico_org_host_endpoint_v1_manifest_test.go: out/install-sentinel terratest/crd_projectcalico_org_v1/crd_projectcalico_org_host_endpoint_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_crd_projectcalico_org_host_endpoint_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/crd_projectcalico_org_v1/crd_projectcalico_org_host_endpoint_v1_manifest_test.go
	touch $@
out/terratest-sentinel-crd_projectcalico_org_ipam_block_v1_manifest_test.go: out/install-sentinel terratest/crd_projectcalico_org_v1/crd_projectcalico_org_ipam_block_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_crd_projectcalico_org_ipam_block_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/crd_projectcalico_org_v1/crd_projectcalico_org_ipam_block_v1_manifest_test.go
	touch $@
out/terratest-sentinel-crd_projectcalico_org_ipam_config_v1_manifest_test.go: out/install-sentinel terratest/crd_projectcalico_org_v1/crd_projectcalico_org_ipam_config_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_crd_projectcalico_org_ipam_config_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/crd_projectcalico_org_v1/crd_projectcalico_org_ipam_config_v1_manifest_test.go
	touch $@
out/terratest-sentinel-crd_projectcalico_org_ipam_handle_v1_manifest_test.go: out/install-sentinel terratest/crd_projectcalico_org_v1/crd_projectcalico_org_ipam_handle_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_crd_projectcalico_org_ipam_handle_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/crd_projectcalico_org_v1/crd_projectcalico_org_ipam_handle_v1_manifest_test.go
	touch $@
out/terratest-sentinel-crd_projectcalico_org_ip_pool_v1_manifest_test.go: out/install-sentinel terratest/crd_projectcalico_org_v1/crd_projectcalico_org_ip_pool_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_crd_projectcalico_org_ip_pool_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/crd_projectcalico_org_v1/crd_projectcalico_org_ip_pool_v1_manifest_test.go
	touch $@
out/terratest-sentinel-crd_projectcalico_org_ip_reservation_v1_manifest_test.go: out/install-sentinel terratest/crd_projectcalico_org_v1/crd_projectcalico_org_ip_reservation_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_crd_projectcalico_org_ip_reservation_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/crd_projectcalico_org_v1/crd_projectcalico_org_ip_reservation_v1_manifest_test.go
	touch $@
out/terratest-sentinel-crd_projectcalico_org_kube_controllers_configuration_v1_manifest_test.go: out/install-sentinel terratest/crd_projectcalico_org_v1/crd_projectcalico_org_kube_controllers_configuration_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_crd_projectcalico_org_kube_controllers_configuration_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/crd_projectcalico_org_v1/crd_projectcalico_org_kube_controllers_configuration_v1_manifest_test.go
	touch $@
out/terratest-sentinel-crd_projectcalico_org_network_policy_v1_manifest_test.go: out/install-sentinel terratest/crd_projectcalico_org_v1/crd_projectcalico_org_network_policy_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_crd_projectcalico_org_network_policy_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/crd_projectcalico_org_v1/crd_projectcalico_org_network_policy_v1_manifest_test.go
	touch $@
out/terratest-sentinel-crd_projectcalico_org_network_set_v1_manifest_test.go: out/install-sentinel terratest/crd_projectcalico_org_v1/crd_projectcalico_org_network_set_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_crd_projectcalico_org_network_set_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/crd_projectcalico_org_v1/crd_projectcalico_org_network_set_v1_manifest_test.go
	touch $@
out/terratest-sentinel-data_fluid_io_alluxio_runtime_v1alpha1_manifest_test.go: out/install-sentinel terratest/data_fluid_io_v1alpha1/data_fluid_io_alluxio_runtime_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_data_fluid_io_alluxio_runtime_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/data_fluid_io_v1alpha1/data_fluid_io_alluxio_runtime_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-data_fluid_io_data_backup_v1alpha1_manifest_test.go: out/install-sentinel terratest/data_fluid_io_v1alpha1/data_fluid_io_data_backup_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_data_fluid_io_data_backup_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/data_fluid_io_v1alpha1/data_fluid_io_data_backup_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-data_fluid_io_data_load_v1alpha1_manifest_test.go: out/install-sentinel terratest/data_fluid_io_v1alpha1/data_fluid_io_data_load_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_data_fluid_io_data_load_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/data_fluid_io_v1alpha1/data_fluid_io_data_load_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-data_fluid_io_dataset_v1alpha1_manifest_test.go: out/install-sentinel terratest/data_fluid_io_v1alpha1/data_fluid_io_dataset_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_data_fluid_io_dataset_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/data_fluid_io_v1alpha1/data_fluid_io_dataset_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-data_fluid_io_goose_fs_runtime_v1alpha1_manifest_test.go: out/install-sentinel terratest/data_fluid_io_v1alpha1/data_fluid_io_goose_fs_runtime_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_data_fluid_io_goose_fs_runtime_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/data_fluid_io_v1alpha1/data_fluid_io_goose_fs_runtime_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-data_fluid_io_jindo_runtime_v1alpha1_manifest_test.go: out/install-sentinel terratest/data_fluid_io_v1alpha1/data_fluid_io_jindo_runtime_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_data_fluid_io_jindo_runtime_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/data_fluid_io_v1alpha1/data_fluid_io_jindo_runtime_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-data_fluid_io_juice_fs_runtime_v1alpha1_manifest_test.go: out/install-sentinel terratest/data_fluid_io_v1alpha1/data_fluid_io_juice_fs_runtime_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_data_fluid_io_juice_fs_runtime_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/data_fluid_io_v1alpha1/data_fluid_io_juice_fs_runtime_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-data_fluid_io_thin_runtime_profile_v1alpha1_manifest_test.go: out/install-sentinel terratest/data_fluid_io_v1alpha1/data_fluid_io_thin_runtime_profile_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_data_fluid_io_thin_runtime_profile_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/data_fluid_io_v1alpha1/data_fluid_io_thin_runtime_profile_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-data_fluid_io_thin_runtime_v1alpha1_manifest_test.go: out/install-sentinel terratest/data_fluid_io_v1alpha1/data_fluid_io_thin_runtime_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_data_fluid_io_thin_runtime_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/data_fluid_io_v1alpha1/data_fluid_io_thin_runtime_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-databases_schemahero_io_database_v1alpha4_manifest_test.go: out/install-sentinel terratest/databases_schemahero_io_v1alpha4/databases_schemahero_io_database_v1alpha4_manifest_test.go $(shell find ./examples/data-sources/k8s_databases_schemahero_io_database_v1alpha4_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/databases_schemahero_io_v1alpha4/databases_schemahero_io_database_v1alpha4_manifest_test.go
	touch $@
out/terratest-sentinel-databases_spotahome_com_redis_failover_v1_manifest_test.go: out/install-sentinel terratest/databases_spotahome_com_v1/databases_spotahome_com_redis_failover_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_databases_spotahome_com_redis_failover_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/databases_spotahome_com_v1/databases_spotahome_com_redis_failover_v1_manifest_test.go
	touch $@
out/terratest-sentinel-datadoghq_com_datadog_agent_v1alpha1_manifest_test.go: out/install-sentinel terratest/datadoghq_com_v1alpha1/datadoghq_com_datadog_agent_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_datadoghq_com_datadog_agent_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/datadoghq_com_v1alpha1/datadoghq_com_datadog_agent_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-datadoghq_com_datadog_metric_v1alpha1_manifest_test.go: out/install-sentinel terratest/datadoghq_com_v1alpha1/datadoghq_com_datadog_metric_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_datadoghq_com_datadog_metric_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/datadoghq_com_v1alpha1/datadoghq_com_datadog_metric_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-datadoghq_com_datadog_monitor_v1alpha1_manifest_test.go: out/install-sentinel terratest/datadoghq_com_v1alpha1/datadoghq_com_datadog_monitor_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_datadoghq_com_datadog_monitor_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/datadoghq_com_v1alpha1/datadoghq_com_datadog_monitor_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-datadoghq_com_datadog_slo_v1alpha1_manifest_test.go: out/install-sentinel terratest/datadoghq_com_v1alpha1/datadoghq_com_datadog_slo_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_datadoghq_com_datadog_slo_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/datadoghq_com_v1alpha1/datadoghq_com_datadog_slo_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-datadoghq_com_datadog_agent_v2alpha1_manifest_test.go: out/install-sentinel terratest/datadoghq_com_v2alpha1/datadoghq_com_datadog_agent_v2alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_datadoghq_com_datadog_agent_v2alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/datadoghq_com_v2alpha1/datadoghq_com_datadog_agent_v2alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-dataprotection_kubeblocks_io_action_set_v1alpha1_manifest_test.go: out/install-sentinel terratest/dataprotection_kubeblocks_io_v1alpha1/dataprotection_kubeblocks_io_action_set_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_dataprotection_kubeblocks_io_action_set_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/dataprotection_kubeblocks_io_v1alpha1/dataprotection_kubeblocks_io_action_set_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-dataprotection_kubeblocks_io_backup_policy_v1alpha1_manifest_test.go: out/install-sentinel terratest/dataprotection_kubeblocks_io_v1alpha1/dataprotection_kubeblocks_io_backup_policy_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_dataprotection_kubeblocks_io_backup_policy_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/dataprotection_kubeblocks_io_v1alpha1/dataprotection_kubeblocks_io_backup_policy_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-dataprotection_kubeblocks_io_backup_repo_v1alpha1_manifest_test.go: out/install-sentinel terratest/dataprotection_kubeblocks_io_v1alpha1/dataprotection_kubeblocks_io_backup_repo_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_dataprotection_kubeblocks_io_backup_repo_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/dataprotection_kubeblocks_io_v1alpha1/dataprotection_kubeblocks_io_backup_repo_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-dataprotection_kubeblocks_io_backup_schedule_v1alpha1_manifest_test.go: out/install-sentinel terratest/dataprotection_kubeblocks_io_v1alpha1/dataprotection_kubeblocks_io_backup_schedule_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_dataprotection_kubeblocks_io_backup_schedule_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/dataprotection_kubeblocks_io_v1alpha1/dataprotection_kubeblocks_io_backup_schedule_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-dataprotection_kubeblocks_io_backup_v1alpha1_manifest_test.go: out/install-sentinel terratest/dataprotection_kubeblocks_io_v1alpha1/dataprotection_kubeblocks_io_backup_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_dataprotection_kubeblocks_io_backup_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/dataprotection_kubeblocks_io_v1alpha1/dataprotection_kubeblocks_io_backup_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-dataprotection_kubeblocks_io_restore_v1alpha1_manifest_test.go: out/install-sentinel terratest/dataprotection_kubeblocks_io_v1alpha1/dataprotection_kubeblocks_io_restore_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_dataprotection_kubeblocks_io_restore_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/dataprotection_kubeblocks_io_v1alpha1/dataprotection_kubeblocks_io_restore_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-designer_kaoto_io_kaoto_v1alpha1_manifest_test.go: out/install-sentinel terratest/designer_kaoto_io_v1alpha1/designer_kaoto_io_kaoto_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_designer_kaoto_io_kaoto_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/designer_kaoto_io_v1alpha1/designer_kaoto_io_kaoto_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-devices_kubeedge_io_device_model_v1alpha2_manifest_test.go: out/install-sentinel terratest/devices_kubeedge_io_v1alpha2/devices_kubeedge_io_device_model_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_devices_kubeedge_io_device_model_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/devices_kubeedge_io_v1alpha2/devices_kubeedge_io_device_model_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-devices_kubeedge_io_device_v1alpha2_manifest_test.go: out/install-sentinel terratest/devices_kubeedge_io_v1alpha2/devices_kubeedge_io_device_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_devices_kubeedge_io_device_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/devices_kubeedge_io_v1alpha2/devices_kubeedge_io_device_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-devices_kubeedge_io_device_model_v1beta1_manifest_test.go: out/install-sentinel terratest/devices_kubeedge_io_v1beta1/devices_kubeedge_io_device_model_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_devices_kubeedge_io_device_model_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/devices_kubeedge_io_v1beta1/devices_kubeedge_io_device_model_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-devices_kubeedge_io_device_v1beta1_manifest_test.go: out/install-sentinel terratest/devices_kubeedge_io_v1beta1/devices_kubeedge_io_device_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_devices_kubeedge_io_device_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/devices_kubeedge_io_v1beta1/devices_kubeedge_io_device_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-devops_kubesphere_io_releaser_controller_v1alpha1_manifest_test.go: out/install-sentinel terratest/devops_kubesphere_io_v1alpha1/devops_kubesphere_io_releaser_controller_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_devops_kubesphere_io_releaser_controller_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/devops_kubesphere_io_v1alpha1/devops_kubesphere_io_releaser_controller_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-devops_kubesphere_io_releaser_v1alpha1_manifest_test.go: out/install-sentinel terratest/devops_kubesphere_io_v1alpha1/devops_kubesphere_io_releaser_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_devops_kubesphere_io_releaser_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/devops_kubesphere_io_v1alpha1/devops_kubesphere_io_releaser_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-dex_gpu_ninja_com_dex_identity_provider_v1alpha1_manifest_test.go: out/install-sentinel terratest/dex_gpu_ninja_com_v1alpha1/dex_gpu_ninja_com_dex_identity_provider_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_dex_gpu_ninja_com_dex_identity_provider_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/dex_gpu_ninja_com_v1alpha1/dex_gpu_ninja_com_dex_identity_provider_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-dex_gpu_ninja_com_dex_o_auth2_client_v1alpha1_manifest_test.go: out/install-sentinel terratest/dex_gpu_ninja_com_v1alpha1/dex_gpu_ninja_com_dex_o_auth2_client_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_dex_gpu_ninja_com_dex_o_auth2_client_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/dex_gpu_ninja_com_v1alpha1/dex_gpu_ninja_com_dex_o_auth2_client_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-dex_gpu_ninja_com_dex_user_v1alpha1_manifest_test.go: out/install-sentinel terratest/dex_gpu_ninja_com_v1alpha1/dex_gpu_ninja_com_dex_user_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_dex_gpu_ninja_com_dex_user_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/dex_gpu_ninja_com_v1alpha1/dex_gpu_ninja_com_dex_user_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-digitalis_io_vals_secret_v1_manifest_test.go: out/install-sentinel terratest/digitalis_io_v1/digitalis_io_vals_secret_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_digitalis_io_vals_secret_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/digitalis_io_v1/digitalis_io_vals_secret_v1_manifest_test.go
	touch $@
out/terratest-sentinel-digitalis_io_db_secret_v1beta1_manifest_test.go: out/install-sentinel terratest/digitalis_io_v1beta1/digitalis_io_db_secret_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_digitalis_io_db_secret_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/digitalis_io_v1beta1/digitalis_io_db_secret_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-discovery_k8s_io_endpoint_slice_v1_manifest_test.go: out/install-sentinel terratest/discovery_k8s_io_v1/discovery_k8s_io_endpoint_slice_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_discovery_k8s_io_endpoint_slice_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/discovery_k8s_io_v1/discovery_k8s_io_endpoint_slice_v1_manifest_test.go
	touch $@
out/terratest-sentinel-documentdb_services_k8s_aws_db_cluster_v1alpha1_manifest_test.go: out/install-sentinel terratest/documentdb_services_k8s_aws_v1alpha1/documentdb_services_k8s_aws_db_cluster_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_documentdb_services_k8s_aws_db_cluster_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/documentdb_services_k8s_aws_v1alpha1/documentdb_services_k8s_aws_db_cluster_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-documentdb_services_k8s_aws_db_instance_v1alpha1_manifest_test.go: out/install-sentinel terratest/documentdb_services_k8s_aws_v1alpha1/documentdb_services_k8s_aws_db_instance_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_documentdb_services_k8s_aws_db_instance_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/documentdb_services_k8s_aws_v1alpha1/documentdb_services_k8s_aws_db_instance_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-documentdb_services_k8s_aws_db_subnet_group_v1alpha1_manifest_test.go: out/install-sentinel terratest/documentdb_services_k8s_aws_v1alpha1/documentdb_services_k8s_aws_db_subnet_group_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_documentdb_services_k8s_aws_db_subnet_group_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/documentdb_services_k8s_aws_v1alpha1/documentdb_services_k8s_aws_db_subnet_group_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-druid_apache_org_druid_v1alpha1_manifest_test.go: out/install-sentinel terratest/druid_apache_org_v1alpha1/druid_apache_org_druid_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_druid_apache_org_druid_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/druid_apache_org_v1alpha1/druid_apache_org_druid_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-dynamodb_services_k8s_aws_backup_v1alpha1_manifest_test.go: out/install-sentinel terratest/dynamodb_services_k8s_aws_v1alpha1/dynamodb_services_k8s_aws_backup_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_dynamodb_services_k8s_aws_backup_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/dynamodb_services_k8s_aws_v1alpha1/dynamodb_services_k8s_aws_backup_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-dynamodb_services_k8s_aws_global_table_v1alpha1_manifest_test.go: out/install-sentinel terratest/dynamodb_services_k8s_aws_v1alpha1/dynamodb_services_k8s_aws_global_table_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_dynamodb_services_k8s_aws_global_table_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/dynamodb_services_k8s_aws_v1alpha1/dynamodb_services_k8s_aws_global_table_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-dynamodb_services_k8s_aws_table_v1alpha1_manifest_test.go: out/install-sentinel terratest/dynamodb_services_k8s_aws_v1alpha1/dynamodb_services_k8s_aws_table_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_dynamodb_services_k8s_aws_table_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/dynamodb_services_k8s_aws_v1alpha1/dynamodb_services_k8s_aws_table_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-ec2_services_k8s_aws_dhcp_options_v1alpha1_manifest_test.go: out/install-sentinel terratest/ec2_services_k8s_aws_v1alpha1/ec2_services_k8s_aws_dhcp_options_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_ec2_services_k8s_aws_dhcp_options_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ec2_services_k8s_aws_v1alpha1/ec2_services_k8s_aws_dhcp_options_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-ec2_services_k8s_aws_elastic_ip_address_v1alpha1_manifest_test.go: out/install-sentinel terratest/ec2_services_k8s_aws_v1alpha1/ec2_services_k8s_aws_elastic_ip_address_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_ec2_services_k8s_aws_elastic_ip_address_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ec2_services_k8s_aws_v1alpha1/ec2_services_k8s_aws_elastic_ip_address_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-ec2_services_k8s_aws_instance_v1alpha1_manifest_test.go: out/install-sentinel terratest/ec2_services_k8s_aws_v1alpha1/ec2_services_k8s_aws_instance_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_ec2_services_k8s_aws_instance_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ec2_services_k8s_aws_v1alpha1/ec2_services_k8s_aws_instance_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-ec2_services_k8s_aws_internet_gateway_v1alpha1_manifest_test.go: out/install-sentinel terratest/ec2_services_k8s_aws_v1alpha1/ec2_services_k8s_aws_internet_gateway_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_ec2_services_k8s_aws_internet_gateway_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ec2_services_k8s_aws_v1alpha1/ec2_services_k8s_aws_internet_gateway_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-ec2_services_k8s_aws_nat_gateway_v1alpha1_manifest_test.go: out/install-sentinel terratest/ec2_services_k8s_aws_v1alpha1/ec2_services_k8s_aws_nat_gateway_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_ec2_services_k8s_aws_nat_gateway_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ec2_services_k8s_aws_v1alpha1/ec2_services_k8s_aws_nat_gateway_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-ec2_services_k8s_aws_route_table_v1alpha1_manifest_test.go: out/install-sentinel terratest/ec2_services_k8s_aws_v1alpha1/ec2_services_k8s_aws_route_table_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_ec2_services_k8s_aws_route_table_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ec2_services_k8s_aws_v1alpha1/ec2_services_k8s_aws_route_table_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-ec2_services_k8s_aws_security_group_v1alpha1_manifest_test.go: out/install-sentinel terratest/ec2_services_k8s_aws_v1alpha1/ec2_services_k8s_aws_security_group_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_ec2_services_k8s_aws_security_group_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ec2_services_k8s_aws_v1alpha1/ec2_services_k8s_aws_security_group_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-ec2_services_k8s_aws_subnet_v1alpha1_manifest_test.go: out/install-sentinel terratest/ec2_services_k8s_aws_v1alpha1/ec2_services_k8s_aws_subnet_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_ec2_services_k8s_aws_subnet_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ec2_services_k8s_aws_v1alpha1/ec2_services_k8s_aws_subnet_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-ec2_services_k8s_aws_transit_gateway_v1alpha1_manifest_test.go: out/install-sentinel terratest/ec2_services_k8s_aws_v1alpha1/ec2_services_k8s_aws_transit_gateway_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_ec2_services_k8s_aws_transit_gateway_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ec2_services_k8s_aws_v1alpha1/ec2_services_k8s_aws_transit_gateway_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-ec2_services_k8s_aws_vpc_v1alpha1_manifest_test.go: out/install-sentinel terratest/ec2_services_k8s_aws_v1alpha1/ec2_services_k8s_aws_vpc_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_ec2_services_k8s_aws_vpc_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ec2_services_k8s_aws_v1alpha1/ec2_services_k8s_aws_vpc_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-ec2_services_k8s_aws_vpc_endpoint_v1alpha1_manifest_test.go: out/install-sentinel terratest/ec2_services_k8s_aws_v1alpha1/ec2_services_k8s_aws_vpc_endpoint_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_ec2_services_k8s_aws_vpc_endpoint_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ec2_services_k8s_aws_v1alpha1/ec2_services_k8s_aws_vpc_endpoint_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-ecr_services_k8s_aws_pull_through_cache_rule_v1alpha1_manifest_test.go: out/install-sentinel terratest/ecr_services_k8s_aws_v1alpha1/ecr_services_k8s_aws_pull_through_cache_rule_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_ecr_services_k8s_aws_pull_through_cache_rule_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ecr_services_k8s_aws_v1alpha1/ecr_services_k8s_aws_pull_through_cache_rule_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-ecr_services_k8s_aws_repository_v1alpha1_manifest_test.go: out/install-sentinel terratest/ecr_services_k8s_aws_v1alpha1/ecr_services_k8s_aws_repository_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_ecr_services_k8s_aws_repository_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ecr_services_k8s_aws_v1alpha1/ecr_services_k8s_aws_repository_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-efs_services_k8s_aws_access_point_v1alpha1_manifest_test.go: out/install-sentinel terratest/efs_services_k8s_aws_v1alpha1/efs_services_k8s_aws_access_point_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_efs_services_k8s_aws_access_point_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/efs_services_k8s_aws_v1alpha1/efs_services_k8s_aws_access_point_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-efs_services_k8s_aws_file_system_v1alpha1_manifest_test.go: out/install-sentinel terratest/efs_services_k8s_aws_v1alpha1/efs_services_k8s_aws_file_system_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_efs_services_k8s_aws_file_system_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/efs_services_k8s_aws_v1alpha1/efs_services_k8s_aws_file_system_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-efs_services_k8s_aws_mount_target_v1alpha1_manifest_test.go: out/install-sentinel terratest/efs_services_k8s_aws_v1alpha1/efs_services_k8s_aws_mount_target_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_efs_services_k8s_aws_mount_target_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/efs_services_k8s_aws_v1alpha1/efs_services_k8s_aws_mount_target_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-eks_services_k8s_aws_addon_v1alpha1_manifest_test.go: out/install-sentinel terratest/eks_services_k8s_aws_v1alpha1/eks_services_k8s_aws_addon_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_eks_services_k8s_aws_addon_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/eks_services_k8s_aws_v1alpha1/eks_services_k8s_aws_addon_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-eks_services_k8s_aws_cluster_v1alpha1_manifest_test.go: out/install-sentinel terratest/eks_services_k8s_aws_v1alpha1/eks_services_k8s_aws_cluster_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_eks_services_k8s_aws_cluster_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/eks_services_k8s_aws_v1alpha1/eks_services_k8s_aws_cluster_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-eks_services_k8s_aws_fargate_profile_v1alpha1_manifest_test.go: out/install-sentinel terratest/eks_services_k8s_aws_v1alpha1/eks_services_k8s_aws_fargate_profile_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_eks_services_k8s_aws_fargate_profile_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/eks_services_k8s_aws_v1alpha1/eks_services_k8s_aws_fargate_profile_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-eks_services_k8s_aws_nodegroup_v1alpha1_manifest_test.go: out/install-sentinel terratest/eks_services_k8s_aws_v1alpha1/eks_services_k8s_aws_nodegroup_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_eks_services_k8s_aws_nodegroup_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/eks_services_k8s_aws_v1alpha1/eks_services_k8s_aws_nodegroup_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-elasticache_services_k8s_aws_cache_parameter_group_v1alpha1_manifest_test.go: out/install-sentinel terratest/elasticache_services_k8s_aws_v1alpha1/elasticache_services_k8s_aws_cache_parameter_group_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_elasticache_services_k8s_aws_cache_parameter_group_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/elasticache_services_k8s_aws_v1alpha1/elasticache_services_k8s_aws_cache_parameter_group_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-elasticache_services_k8s_aws_cache_subnet_group_v1alpha1_manifest_test.go: out/install-sentinel terratest/elasticache_services_k8s_aws_v1alpha1/elasticache_services_k8s_aws_cache_subnet_group_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_elasticache_services_k8s_aws_cache_subnet_group_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/elasticache_services_k8s_aws_v1alpha1/elasticache_services_k8s_aws_cache_subnet_group_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-elasticache_services_k8s_aws_replication_group_v1alpha1_manifest_test.go: out/install-sentinel terratest/elasticache_services_k8s_aws_v1alpha1/elasticache_services_k8s_aws_replication_group_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_elasticache_services_k8s_aws_replication_group_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/elasticache_services_k8s_aws_v1alpha1/elasticache_services_k8s_aws_replication_group_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-elasticache_services_k8s_aws_snapshot_v1alpha1_manifest_test.go: out/install-sentinel terratest/elasticache_services_k8s_aws_v1alpha1/elasticache_services_k8s_aws_snapshot_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_elasticache_services_k8s_aws_snapshot_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/elasticache_services_k8s_aws_v1alpha1/elasticache_services_k8s_aws_snapshot_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-elasticache_services_k8s_aws_user_group_v1alpha1_manifest_test.go: out/install-sentinel terratest/elasticache_services_k8s_aws_v1alpha1/elasticache_services_k8s_aws_user_group_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_elasticache_services_k8s_aws_user_group_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/elasticache_services_k8s_aws_v1alpha1/elasticache_services_k8s_aws_user_group_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-elasticache_services_k8s_aws_user_v1alpha1_manifest_test.go: out/install-sentinel terratest/elasticache_services_k8s_aws_v1alpha1/elasticache_services_k8s_aws_user_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_elasticache_services_k8s_aws_user_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/elasticache_services_k8s_aws_v1alpha1/elasticache_services_k8s_aws_user_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-elasticsearch_k8s_elastic_co_elasticsearch_v1_manifest_test.go: out/install-sentinel terratest/elasticsearch_k8s_elastic_co_v1/elasticsearch_k8s_elastic_co_elasticsearch_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_elasticsearch_k8s_elastic_co_elasticsearch_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/elasticsearch_k8s_elastic_co_v1/elasticsearch_k8s_elastic_co_elasticsearch_v1_manifest_test.go
	touch $@
out/terratest-sentinel-elasticsearch_k8s_elastic_co_elasticsearch_v1beta1_manifest_test.go: out/install-sentinel terratest/elasticsearch_k8s_elastic_co_v1beta1/elasticsearch_k8s_elastic_co_elasticsearch_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_elasticsearch_k8s_elastic_co_elasticsearch_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/elasticsearch_k8s_elastic_co_v1beta1/elasticsearch_k8s_elastic_co_elasticsearch_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-elbv2_k8s_aws_target_group_binding_v1alpha1_manifest_test.go: out/install-sentinel terratest/elbv2_k8s_aws_v1alpha1/elbv2_k8s_aws_target_group_binding_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_elbv2_k8s_aws_target_group_binding_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/elbv2_k8s_aws_v1alpha1/elbv2_k8s_aws_target_group_binding_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-elbv2_k8s_aws_ingress_class_params_v1beta1_manifest_test.go: out/install-sentinel terratest/elbv2_k8s_aws_v1beta1/elbv2_k8s_aws_ingress_class_params_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_elbv2_k8s_aws_ingress_class_params_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/elbv2_k8s_aws_v1beta1/elbv2_k8s_aws_ingress_class_params_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-elbv2_k8s_aws_target_group_binding_v1beta1_manifest_test.go: out/install-sentinel terratest/elbv2_k8s_aws_v1beta1/elbv2_k8s_aws_target_group_binding_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_elbv2_k8s_aws_target_group_binding_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/elbv2_k8s_aws_v1beta1/elbv2_k8s_aws_target_group_binding_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-emrcontainers_services_k8s_aws_job_run_v1alpha1_manifest_test.go: out/install-sentinel terratest/emrcontainers_services_k8s_aws_v1alpha1/emrcontainers_services_k8s_aws_job_run_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_emrcontainers_services_k8s_aws_job_run_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/emrcontainers_services_k8s_aws_v1alpha1/emrcontainers_services_k8s_aws_job_run_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-emrcontainers_services_k8s_aws_virtual_cluster_v1alpha1_manifest_test.go: out/install-sentinel terratest/emrcontainers_services_k8s_aws_v1alpha1/emrcontainers_services_k8s_aws_virtual_cluster_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_emrcontainers_services_k8s_aws_virtual_cluster_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/emrcontainers_services_k8s_aws_v1alpha1/emrcontainers_services_k8s_aws_virtual_cluster_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-ensembleoss_io_cluster_v1_manifest_test.go: out/install-sentinel terratest/ensembleoss_io_v1/ensembleoss_io_cluster_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_ensembleoss_io_cluster_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ensembleoss_io_v1/ensembleoss_io_cluster_v1_manifest_test.go
	touch $@
out/terratest-sentinel-ensembleoss_io_resource_v1_manifest_test.go: out/install-sentinel terratest/ensembleoss_io_v1/ensembleoss_io_resource_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_ensembleoss_io_resource_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ensembleoss_io_v1/ensembleoss_io_resource_v1_manifest_test.go
	touch $@
out/terratest-sentinel-enterprise_gloo_solo_io_auth_config_v1_manifest_test.go: out/install-sentinel terratest/enterprise_gloo_solo_io_v1/enterprise_gloo_solo_io_auth_config_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_enterprise_gloo_solo_io_auth_config_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/enterprise_gloo_solo_io_v1/enterprise_gloo_solo_io_auth_config_v1_manifest_test.go
	touch $@
out/terratest-sentinel-enterprisesearch_k8s_elastic_co_enterprise_search_v1_manifest_test.go: out/install-sentinel terratest/enterprisesearch_k8s_elastic_co_v1/enterprisesearch_k8s_elastic_co_enterprise_search_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_enterprisesearch_k8s_elastic_co_enterprise_search_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/enterprisesearch_k8s_elastic_co_v1/enterprisesearch_k8s_elastic_co_enterprise_search_v1_manifest_test.go
	touch $@
out/terratest-sentinel-enterprisesearch_k8s_elastic_co_enterprise_search_v1beta1_manifest_test.go: out/install-sentinel terratest/enterprisesearch_k8s_elastic_co_v1beta1/enterprisesearch_k8s_elastic_co_enterprise_search_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_enterprisesearch_k8s_elastic_co_enterprise_search_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/enterprisesearch_k8s_elastic_co_v1beta1/enterprisesearch_k8s_elastic_co_enterprise_search_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-events_k8s_io_event_v1_manifest_test.go: out/install-sentinel terratest/events_k8s_io_v1/events_k8s_io_event_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_events_k8s_io_event_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/events_k8s_io_v1/events_k8s_io_event_v1_manifest_test.go
	touch $@
out/terratest-sentinel-everest_percona_com_backup_storage_v1alpha1_manifest_test.go: out/install-sentinel terratest/everest_percona_com_v1alpha1/everest_percona_com_backup_storage_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_everest_percona_com_backup_storage_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/everest_percona_com_v1alpha1/everest_percona_com_backup_storage_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-everest_percona_com_database_cluster_backup_v1alpha1_manifest_test.go: out/install-sentinel terratest/everest_percona_com_v1alpha1/everest_percona_com_database_cluster_backup_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_everest_percona_com_database_cluster_backup_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/everest_percona_com_v1alpha1/everest_percona_com_database_cluster_backup_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-everest_percona_com_database_cluster_restore_v1alpha1_manifest_test.go: out/install-sentinel terratest/everest_percona_com_v1alpha1/everest_percona_com_database_cluster_restore_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_everest_percona_com_database_cluster_restore_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/everest_percona_com_v1alpha1/everest_percona_com_database_cluster_restore_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-everest_percona_com_database_cluster_v1alpha1_manifest_test.go: out/install-sentinel terratest/everest_percona_com_v1alpha1/everest_percona_com_database_cluster_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_everest_percona_com_database_cluster_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/everest_percona_com_v1alpha1/everest_percona_com_database_cluster_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-everest_percona_com_database_engine_v1alpha1_manifest_test.go: out/install-sentinel terratest/everest_percona_com_v1alpha1/everest_percona_com_database_engine_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_everest_percona_com_database_engine_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/everest_percona_com_v1alpha1/everest_percona_com_database_engine_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-everest_percona_com_monitoring_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/everest_percona_com_v1alpha1/everest_percona_com_monitoring_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_everest_percona_com_monitoring_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/everest_percona_com_v1alpha1/everest_percona_com_monitoring_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-execution_furiko_io_job_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/execution_furiko_io_v1alpha1/execution_furiko_io_job_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_execution_furiko_io_job_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/execution_furiko_io_v1alpha1/execution_furiko_io_job_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-execution_furiko_io_job_v1alpha1_manifest_test.go: out/install-sentinel terratest/execution_furiko_io_v1alpha1/execution_furiko_io_job_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_execution_furiko_io_job_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/execution_furiko_io_v1alpha1/execution_furiko_io_job_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-executor_testkube_io_executor_v1_manifest_test.go: out/install-sentinel terratest/executor_testkube_io_v1/executor_testkube_io_executor_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_executor_testkube_io_executor_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/executor_testkube_io_v1/executor_testkube_io_executor_v1_manifest_test.go
	touch $@
out/terratest-sentinel-executor_testkube_io_webhook_v1_manifest_test.go: out/install-sentinel terratest/executor_testkube_io_v1/executor_testkube_io_webhook_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_executor_testkube_io_webhook_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/executor_testkube_io_v1/executor_testkube_io_webhook_v1_manifest_test.go
	touch $@
out/terratest-sentinel-expansion_gatekeeper_sh_expansion_template_v1alpha1_manifest_test.go: out/install-sentinel terratest/expansion_gatekeeper_sh_v1alpha1/expansion_gatekeeper_sh_expansion_template_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_expansion_gatekeeper_sh_expansion_template_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/expansion_gatekeeper_sh_v1alpha1/expansion_gatekeeper_sh_expansion_template_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-expansion_gatekeeper_sh_expansion_template_v1beta1_manifest_test.go: out/install-sentinel terratest/expansion_gatekeeper_sh_v1beta1/expansion_gatekeeper_sh_expansion_template_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_expansion_gatekeeper_sh_expansion_template_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/expansion_gatekeeper_sh_v1beta1/expansion_gatekeeper_sh_expansion_template_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-extensions_istio_io_wasm_plugin_v1alpha1_manifest_test.go: out/install-sentinel terratest/extensions_istio_io_v1alpha1/extensions_istio_io_wasm_plugin_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_extensions_istio_io_wasm_plugin_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/extensions_istio_io_v1alpha1/extensions_istio_io_wasm_plugin_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-extensions_kubeblocks_io_addon_v1alpha1_manifest_test.go: out/install-sentinel terratest/extensions_kubeblocks_io_v1alpha1/extensions_kubeblocks_io_addon_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_extensions_kubeblocks_io_addon_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/extensions_kubeblocks_io_v1alpha1/extensions_kubeblocks_io_addon_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-external_secrets_io_cluster_secret_store_v1alpha1_manifest_test.go: out/install-sentinel terratest/external_secrets_io_v1alpha1/external_secrets_io_cluster_secret_store_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_external_secrets_io_cluster_secret_store_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/external_secrets_io_v1alpha1/external_secrets_io_cluster_secret_store_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-external_secrets_io_external_secret_v1alpha1_manifest_test.go: out/install-sentinel terratest/external_secrets_io_v1alpha1/external_secrets_io_external_secret_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_external_secrets_io_external_secret_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/external_secrets_io_v1alpha1/external_secrets_io_external_secret_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-external_secrets_io_secret_store_v1alpha1_manifest_test.go: out/install-sentinel terratest/external_secrets_io_v1alpha1/external_secrets_io_secret_store_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_external_secrets_io_secret_store_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/external_secrets_io_v1alpha1/external_secrets_io_secret_store_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-external_secrets_io_cluster_external_secret_v1beta1_manifest_test.go: out/install-sentinel terratest/external_secrets_io_v1beta1/external_secrets_io_cluster_external_secret_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_external_secrets_io_cluster_external_secret_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/external_secrets_io_v1beta1/external_secrets_io_cluster_external_secret_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-external_secrets_io_cluster_secret_store_v1beta1_manifest_test.go: out/install-sentinel terratest/external_secrets_io_v1beta1/external_secrets_io_cluster_secret_store_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_external_secrets_io_cluster_secret_store_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/external_secrets_io_v1beta1/external_secrets_io_cluster_secret_store_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-external_secrets_io_external_secret_v1beta1_manifest_test.go: out/install-sentinel terratest/external_secrets_io_v1beta1/external_secrets_io_external_secret_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_external_secrets_io_external_secret_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/external_secrets_io_v1beta1/external_secrets_io_external_secret_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-external_secrets_io_secret_store_v1beta1_manifest_test.go: out/install-sentinel terratest/external_secrets_io_v1beta1/external_secrets_io_secret_store_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_external_secrets_io_secret_store_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/external_secrets_io_v1beta1/external_secrets_io_secret_store_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-externaldata_gatekeeper_sh_provider_v1alpha1_manifest_test.go: out/install-sentinel terratest/externaldata_gatekeeper_sh_v1alpha1/externaldata_gatekeeper_sh_provider_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_externaldata_gatekeeper_sh_provider_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/externaldata_gatekeeper_sh_v1alpha1/externaldata_gatekeeper_sh_provider_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-externaldata_gatekeeper_sh_provider_v1beta1_manifest_test.go: out/install-sentinel terratest/externaldata_gatekeeper_sh_v1beta1/externaldata_gatekeeper_sh_provider_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_externaldata_gatekeeper_sh_provider_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/externaldata_gatekeeper_sh_v1beta1/externaldata_gatekeeper_sh_provider_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-externaldns_k8s_io_dns_endpoint_v1alpha1_manifest_test.go: out/install-sentinel terratest/externaldns_k8s_io_v1alpha1/externaldns_k8s_io_dns_endpoint_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_externaldns_k8s_io_dns_endpoint_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/externaldns_k8s_io_v1alpha1/externaldns_k8s_io_dns_endpoint_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-externaldns_nginx_org_dns_endpoint_v1_manifest_test.go: out/install-sentinel terratest/externaldns_nginx_org_v1/externaldns_nginx_org_dns_endpoint_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_externaldns_nginx_org_dns_endpoint_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/externaldns_nginx_org_v1/externaldns_nginx_org_dns_endpoint_v1_manifest_test.go
	touch $@
out/terratest-sentinel-fence_agents_remediation_medik8s_io_fence_agents_remediation_template_v1alpha1_manifest_test.go: out/install-sentinel terratest/fence_agents_remediation_medik8s_io_v1alpha1/fence_agents_remediation_medik8s_io_fence_agents_remediation_template_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_fence_agents_remediation_medik8s_io_fence_agents_remediation_template_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/fence_agents_remediation_medik8s_io_v1alpha1/fence_agents_remediation_medik8s_io_fence_agents_remediation_template_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-fence_agents_remediation_medik8s_io_fence_agents_remediation_v1alpha1_manifest_test.go: out/install-sentinel terratest/fence_agents_remediation_medik8s_io_v1alpha1/fence_agents_remediation_medik8s_io_fence_agents_remediation_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_fence_agents_remediation_medik8s_io_fence_agents_remediation_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/fence_agents_remediation_medik8s_io_v1alpha1/fence_agents_remediation_medik8s_io_fence_agents_remediation_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-flagger_app_alert_provider_v1beta1_manifest_test.go: out/install-sentinel terratest/flagger_app_v1beta1/flagger_app_alert_provider_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_flagger_app_alert_provider_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/flagger_app_v1beta1/flagger_app_alert_provider_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-flagger_app_canary_v1beta1_manifest_test.go: out/install-sentinel terratest/flagger_app_v1beta1/flagger_app_canary_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_flagger_app_canary_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/flagger_app_v1beta1/flagger_app_canary_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-flagger_app_metric_template_v1beta1_manifest_test.go: out/install-sentinel terratest/flagger_app_v1beta1/flagger_app_metric_template_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_flagger_app_metric_template_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/flagger_app_v1beta1/flagger_app_metric_template_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-flink_apache_org_flink_deployment_v1beta1_manifest_test.go: out/install-sentinel terratest/flink_apache_org_v1beta1/flink_apache_org_flink_deployment_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_flink_apache_org_flink_deployment_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/flink_apache_org_v1beta1/flink_apache_org_flink_deployment_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-flink_apache_org_flink_session_job_v1beta1_manifest_test.go: out/install-sentinel terratest/flink_apache_org_v1beta1/flink_apache_org_flink_session_job_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_flink_apache_org_flink_session_job_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/flink_apache_org_v1beta1/flink_apache_org_flink_session_job_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-flow_volcano_sh_job_flow_v1alpha1_manifest_test.go: out/install-sentinel terratest/flow_volcano_sh_v1alpha1/flow_volcano_sh_job_flow_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_flow_volcano_sh_job_flow_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/flow_volcano_sh_v1alpha1/flow_volcano_sh_job_flow_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-flow_volcano_sh_job_template_v1alpha1_manifest_test.go: out/install-sentinel terratest/flow_volcano_sh_v1alpha1/flow_volcano_sh_job_template_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_flow_volcano_sh_job_template_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/flow_volcano_sh_v1alpha1/flow_volcano_sh_job_template_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-flowcontrol_apiserver_k8s_io_flow_schema_v1beta3_manifest_test.go: out/install-sentinel terratest/flowcontrol_apiserver_k8s_io_v1beta3/flowcontrol_apiserver_k8s_io_flow_schema_v1beta3_manifest_test.go $(shell find ./examples/data-sources/k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1beta3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/flowcontrol_apiserver_k8s_io_v1beta3/flowcontrol_apiserver_k8s_io_flow_schema_v1beta3_manifest_test.go
	touch $@
out/terratest-sentinel-flowcontrol_apiserver_k8s_io_priority_level_configuration_v1beta3_manifest_test.go: out/install-sentinel terratest/flowcontrol_apiserver_k8s_io_v1beta3/flowcontrol_apiserver_k8s_io_priority_level_configuration_v1beta3_manifest_test.go $(shell find ./examples/data-sources/k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1beta3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/flowcontrol_apiserver_k8s_io_v1beta3/flowcontrol_apiserver_k8s_io_priority_level_configuration_v1beta3_manifest_test.go
	touch $@
out/terratest-sentinel-flows_netobserv_io_flow_collector_v1alpha1_manifest_test.go: out/install-sentinel terratest/flows_netobserv_io_v1alpha1/flows_netobserv_io_flow_collector_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_flows_netobserv_io_flow_collector_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/flows_netobserv_io_v1alpha1/flows_netobserv_io_flow_collector_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-flows_netobserv_io_flow_collector_v1beta1_manifest_test.go: out/install-sentinel terratest/flows_netobserv_io_v1beta1/flows_netobserv_io_flow_collector_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_flows_netobserv_io_flow_collector_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/flows_netobserv_io_v1beta1/flows_netobserv_io_flow_collector_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-flows_netobserv_io_flow_collector_v1beta2_manifest_test.go: out/install-sentinel terratest/flows_netobserv_io_v1beta2/flows_netobserv_io_flow_collector_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_flows_netobserv_io_flow_collector_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/flows_netobserv_io_v1beta2/flows_netobserv_io_flow_collector_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-fluentbit_fluent_io_cluster_filter_v1alpha2_manifest_test.go: out/install-sentinel terratest/fluentbit_fluent_io_v1alpha2/fluentbit_fluent_io_cluster_filter_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_fluentbit_fluent_io_cluster_filter_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/fluentbit_fluent_io_v1alpha2/fluentbit_fluent_io_cluster_filter_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-fluentbit_fluent_io_cluster_fluent_bit_config_v1alpha2_manifest_test.go: out/install-sentinel terratest/fluentbit_fluent_io_v1alpha2/fluentbit_fluent_io_cluster_fluent_bit_config_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_fluentbit_fluent_io_cluster_fluent_bit_config_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/fluentbit_fluent_io_v1alpha2/fluentbit_fluent_io_cluster_fluent_bit_config_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-fluentbit_fluent_io_cluster_input_v1alpha2_manifest_test.go: out/install-sentinel terratest/fluentbit_fluent_io_v1alpha2/fluentbit_fluent_io_cluster_input_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_fluentbit_fluent_io_cluster_input_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/fluentbit_fluent_io_v1alpha2/fluentbit_fluent_io_cluster_input_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-fluentbit_fluent_io_cluster_output_v1alpha2_manifest_test.go: out/install-sentinel terratest/fluentbit_fluent_io_v1alpha2/fluentbit_fluent_io_cluster_output_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_fluentbit_fluent_io_cluster_output_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/fluentbit_fluent_io_v1alpha2/fluentbit_fluent_io_cluster_output_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-fluentbit_fluent_io_cluster_parser_v1alpha2_manifest_test.go: out/install-sentinel terratest/fluentbit_fluent_io_v1alpha2/fluentbit_fluent_io_cluster_parser_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_fluentbit_fluent_io_cluster_parser_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/fluentbit_fluent_io_v1alpha2/fluentbit_fluent_io_cluster_parser_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-fluentbit_fluent_io_collector_v1alpha2_manifest_test.go: out/install-sentinel terratest/fluentbit_fluent_io_v1alpha2/fluentbit_fluent_io_collector_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_fluentbit_fluent_io_collector_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/fluentbit_fluent_io_v1alpha2/fluentbit_fluent_io_collector_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-fluentbit_fluent_io_filter_v1alpha2_manifest_test.go: out/install-sentinel terratest/fluentbit_fluent_io_v1alpha2/fluentbit_fluent_io_filter_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_fluentbit_fluent_io_filter_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/fluentbit_fluent_io_v1alpha2/fluentbit_fluent_io_filter_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-fluentbit_fluent_io_fluent_bit_config_v1alpha2_manifest_test.go: out/install-sentinel terratest/fluentbit_fluent_io_v1alpha2/fluentbit_fluent_io_fluent_bit_config_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_fluentbit_fluent_io_fluent_bit_config_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/fluentbit_fluent_io_v1alpha2/fluentbit_fluent_io_fluent_bit_config_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-fluentbit_fluent_io_fluent_bit_v1alpha2_manifest_test.go: out/install-sentinel terratest/fluentbit_fluent_io_v1alpha2/fluentbit_fluent_io_fluent_bit_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_fluentbit_fluent_io_fluent_bit_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/fluentbit_fluent_io_v1alpha2/fluentbit_fluent_io_fluent_bit_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-fluentbit_fluent_io_output_v1alpha2_manifest_test.go: out/install-sentinel terratest/fluentbit_fluent_io_v1alpha2/fluentbit_fluent_io_output_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_fluentbit_fluent_io_output_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/fluentbit_fluent_io_v1alpha2/fluentbit_fluent_io_output_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-fluentbit_fluent_io_parser_v1alpha2_manifest_test.go: out/install-sentinel terratest/fluentbit_fluent_io_v1alpha2/fluentbit_fluent_io_parser_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_fluentbit_fluent_io_parser_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/fluentbit_fluent_io_v1alpha2/fluentbit_fluent_io_parser_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-fluentd_fluent_io_cluster_filter_v1alpha1_manifest_test.go: out/install-sentinel terratest/fluentd_fluent_io_v1alpha1/fluentd_fluent_io_cluster_filter_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_fluentd_fluent_io_cluster_filter_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/fluentd_fluent_io_v1alpha1/fluentd_fluent_io_cluster_filter_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-fluentd_fluent_io_cluster_fluentd_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/fluentd_fluent_io_v1alpha1/fluentd_fluent_io_cluster_fluentd_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_fluentd_fluent_io_cluster_fluentd_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/fluentd_fluent_io_v1alpha1/fluentd_fluent_io_cluster_fluentd_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-fluentd_fluent_io_cluster_input_v1alpha1_manifest_test.go: out/install-sentinel terratest/fluentd_fluent_io_v1alpha1/fluentd_fluent_io_cluster_input_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_fluentd_fluent_io_cluster_input_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/fluentd_fluent_io_v1alpha1/fluentd_fluent_io_cluster_input_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-fluentd_fluent_io_cluster_output_v1alpha1_manifest_test.go: out/install-sentinel terratest/fluentd_fluent_io_v1alpha1/fluentd_fluent_io_cluster_output_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_fluentd_fluent_io_cluster_output_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/fluentd_fluent_io_v1alpha1/fluentd_fluent_io_cluster_output_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-fluentd_fluent_io_filter_v1alpha1_manifest_test.go: out/install-sentinel terratest/fluentd_fluent_io_v1alpha1/fluentd_fluent_io_filter_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_fluentd_fluent_io_filter_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/fluentd_fluent_io_v1alpha1/fluentd_fluent_io_filter_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-fluentd_fluent_io_fluentd_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/fluentd_fluent_io_v1alpha1/fluentd_fluent_io_fluentd_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_fluentd_fluent_io_fluentd_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/fluentd_fluent_io_v1alpha1/fluentd_fluent_io_fluentd_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-fluentd_fluent_io_fluentd_v1alpha1_manifest_test.go: out/install-sentinel terratest/fluentd_fluent_io_v1alpha1/fluentd_fluent_io_fluentd_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_fluentd_fluent_io_fluentd_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/fluentd_fluent_io_v1alpha1/fluentd_fluent_io_fluentd_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-fluentd_fluent_io_input_v1alpha1_manifest_test.go: out/install-sentinel terratest/fluentd_fluent_io_v1alpha1/fluentd_fluent_io_input_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_fluentd_fluent_io_input_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/fluentd_fluent_io_v1alpha1/fluentd_fluent_io_input_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-fluentd_fluent_io_output_v1alpha1_manifest_test.go: out/install-sentinel terratest/fluentd_fluent_io_v1alpha1/fluentd_fluent_io_output_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_fluentd_fluent_io_output_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/fluentd_fluent_io_v1alpha1/fluentd_fluent_io_output_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-flux_framework_org_mini_cluster_v1alpha1_manifest_test.go: out/install-sentinel terratest/flux_framework_org_v1alpha1/flux_framework_org_mini_cluster_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_flux_framework_org_mini_cluster_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/flux_framework_org_v1alpha1/flux_framework_org_mini_cluster_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-flux_framework_org_mini_cluster_v1alpha2_manifest_test.go: out/install-sentinel terratest/flux_framework_org_v1alpha2/flux_framework_org_mini_cluster_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_flux_framework_org_mini_cluster_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/flux_framework_org_v1alpha2/flux_framework_org_mini_cluster_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-forklift_konveyor_io_forklift_controller_v1beta1_manifest_test.go: out/install-sentinel terratest/forklift_konveyor_io_v1beta1/forklift_konveyor_io_forklift_controller_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_forklift_konveyor_io_forklift_controller_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/forklift_konveyor_io_v1beta1/forklift_konveyor_io_forklift_controller_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-forklift_konveyor_io_hook_v1beta1_manifest_test.go: out/install-sentinel terratest/forklift_konveyor_io_v1beta1/forklift_konveyor_io_hook_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_forklift_konveyor_io_hook_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/forklift_konveyor_io_v1beta1/forklift_konveyor_io_hook_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-forklift_konveyor_io_host_v1beta1_manifest_test.go: out/install-sentinel terratest/forklift_konveyor_io_v1beta1/forklift_konveyor_io_host_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_forklift_konveyor_io_host_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/forklift_konveyor_io_v1beta1/forklift_konveyor_io_host_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-forklift_konveyor_io_migration_v1beta1_manifest_test.go: out/install-sentinel terratest/forklift_konveyor_io_v1beta1/forklift_konveyor_io_migration_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_forklift_konveyor_io_migration_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/forklift_konveyor_io_v1beta1/forklift_konveyor_io_migration_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-forklift_konveyor_io_network_map_v1beta1_manifest_test.go: out/install-sentinel terratest/forklift_konveyor_io_v1beta1/forklift_konveyor_io_network_map_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_forklift_konveyor_io_network_map_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/forklift_konveyor_io_v1beta1/forklift_konveyor_io_network_map_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-forklift_konveyor_io_openstack_volume_populator_v1beta1_manifest_test.go: out/install-sentinel terratest/forklift_konveyor_io_v1beta1/forklift_konveyor_io_openstack_volume_populator_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_forklift_konveyor_io_openstack_volume_populator_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/forklift_konveyor_io_v1beta1/forklift_konveyor_io_openstack_volume_populator_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-forklift_konveyor_io_ovirt_volume_populator_v1beta1_manifest_test.go: out/install-sentinel terratest/forklift_konveyor_io_v1beta1/forklift_konveyor_io_ovirt_volume_populator_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_forklift_konveyor_io_ovirt_volume_populator_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/forklift_konveyor_io_v1beta1/forklift_konveyor_io_ovirt_volume_populator_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-forklift_konveyor_io_plan_v1beta1_manifest_test.go: out/install-sentinel terratest/forklift_konveyor_io_v1beta1/forklift_konveyor_io_plan_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_forklift_konveyor_io_plan_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/forklift_konveyor_io_v1beta1/forklift_konveyor_io_plan_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-forklift_konveyor_io_provider_v1beta1_manifest_test.go: out/install-sentinel terratest/forklift_konveyor_io_v1beta1/forklift_konveyor_io_provider_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_forklift_konveyor_io_provider_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/forklift_konveyor_io_v1beta1/forklift_konveyor_io_provider_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-forklift_konveyor_io_storage_map_v1beta1_manifest_test.go: out/install-sentinel terratest/forklift_konveyor_io_v1beta1/forklift_konveyor_io_storage_map_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_forklift_konveyor_io_storage_map_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/forklift_konveyor_io_v1beta1/forklift_konveyor_io_storage_map_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-fossul_io_backup_config_v1_manifest_test.go: out/install-sentinel terratest/fossul_io_v1/fossul_io_backup_config_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_fossul_io_backup_config_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/fossul_io_v1/fossul_io_backup_config_v1_manifest_test.go
	touch $@
out/terratest-sentinel-fossul_io_backup_schedule_v1_manifest_test.go: out/install-sentinel terratest/fossul_io_v1/fossul_io_backup_schedule_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_fossul_io_backup_schedule_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/fossul_io_v1/fossul_io_backup_schedule_v1_manifest_test.go
	touch $@
out/terratest-sentinel-fossul_io_backup_v1_manifest_test.go: out/install-sentinel terratest/fossul_io_v1/fossul_io_backup_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_fossul_io_backup_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/fossul_io_v1/fossul_io_backup_v1_manifest_test.go
	touch $@
out/terratest-sentinel-fossul_io_fossul_v1_manifest_test.go: out/install-sentinel terratest/fossul_io_v1/fossul_io_fossul_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_fossul_io_fossul_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/fossul_io_v1/fossul_io_fossul_v1_manifest_test.go
	touch $@
out/terratest-sentinel-fossul_io_restore_v1_manifest_test.go: out/install-sentinel terratest/fossul_io_v1/fossul_io_restore_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_fossul_io_restore_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/fossul_io_v1/fossul_io_restore_v1_manifest_test.go
	touch $@
out/terratest-sentinel-gateway_networking_k8s_io_gateway_class_v1_manifest_test.go: out/install-sentinel terratest/gateway_networking_k8s_io_v1/gateway_networking_k8s_io_gateway_class_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_gateway_networking_k8s_io_gateway_class_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/gateway_networking_k8s_io_v1/gateway_networking_k8s_io_gateway_class_v1_manifest_test.go
	touch $@
out/terratest-sentinel-gateway_networking_k8s_io_gateway_v1_manifest_test.go: out/install-sentinel terratest/gateway_networking_k8s_io_v1/gateway_networking_k8s_io_gateway_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_gateway_networking_k8s_io_gateway_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/gateway_networking_k8s_io_v1/gateway_networking_k8s_io_gateway_v1_manifest_test.go
	touch $@
out/terratest-sentinel-gateway_networking_k8s_io_grpc_route_v1_manifest_test.go: out/install-sentinel terratest/gateway_networking_k8s_io_v1/gateway_networking_k8s_io_grpc_route_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_gateway_networking_k8s_io_grpc_route_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/gateway_networking_k8s_io_v1/gateway_networking_k8s_io_grpc_route_v1_manifest_test.go
	touch $@
out/terratest-sentinel-gateway_networking_k8s_io_http_route_v1_manifest_test.go: out/install-sentinel terratest/gateway_networking_k8s_io_v1/gateway_networking_k8s_io_http_route_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_gateway_networking_k8s_io_http_route_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/gateway_networking_k8s_io_v1/gateway_networking_k8s_io_http_route_v1_manifest_test.go
	touch $@
out/terratest-sentinel-gateway_networking_k8s_io_grpc_route_v1alpha2_manifest_test.go: out/install-sentinel terratest/gateway_networking_k8s_io_v1alpha2/gateway_networking_k8s_io_grpc_route_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_gateway_networking_k8s_io_grpc_route_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/gateway_networking_k8s_io_v1alpha2/gateway_networking_k8s_io_grpc_route_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-gateway_networking_k8s_io_reference_grant_v1alpha2_manifest_test.go: out/install-sentinel terratest/gateway_networking_k8s_io_v1alpha2/gateway_networking_k8s_io_reference_grant_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_gateway_networking_k8s_io_reference_grant_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/gateway_networking_k8s_io_v1alpha2/gateway_networking_k8s_io_reference_grant_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-gateway_networking_k8s_io_tcp_route_v1alpha2_manifest_test.go: out/install-sentinel terratest/gateway_networking_k8s_io_v1alpha2/gateway_networking_k8s_io_tcp_route_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_gateway_networking_k8s_io_tcp_route_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/gateway_networking_k8s_io_v1alpha2/gateway_networking_k8s_io_tcp_route_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-gateway_networking_k8s_io_tls_route_v1alpha2_manifest_test.go: out/install-sentinel terratest/gateway_networking_k8s_io_v1alpha2/gateway_networking_k8s_io_tls_route_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_gateway_networking_k8s_io_tls_route_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/gateway_networking_k8s_io_v1alpha2/gateway_networking_k8s_io_tls_route_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-gateway_networking_k8s_io_udp_route_v1alpha2_manifest_test.go: out/install-sentinel terratest/gateway_networking_k8s_io_v1alpha2/gateway_networking_k8s_io_udp_route_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_gateway_networking_k8s_io_udp_route_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/gateway_networking_k8s_io_v1alpha2/gateway_networking_k8s_io_udp_route_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-gateway_networking_k8s_io_gateway_class_v1beta1_manifest_test.go: out/install-sentinel terratest/gateway_networking_k8s_io_v1beta1/gateway_networking_k8s_io_gateway_class_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_gateway_networking_k8s_io_gateway_class_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/gateway_networking_k8s_io_v1beta1/gateway_networking_k8s_io_gateway_class_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-gateway_networking_k8s_io_gateway_v1beta1_manifest_test.go: out/install-sentinel terratest/gateway_networking_k8s_io_v1beta1/gateway_networking_k8s_io_gateway_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_gateway_networking_k8s_io_gateway_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/gateway_networking_k8s_io_v1beta1/gateway_networking_k8s_io_gateway_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-gateway_networking_k8s_io_http_route_v1beta1_manifest_test.go: out/install-sentinel terratest/gateway_networking_k8s_io_v1beta1/gateway_networking_k8s_io_http_route_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_gateway_networking_k8s_io_http_route_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/gateway_networking_k8s_io_v1beta1/gateway_networking_k8s_io_http_route_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-gateway_networking_k8s_io_reference_grant_v1beta1_manifest_test.go: out/install-sentinel terratest/gateway_networking_k8s_io_v1beta1/gateway_networking_k8s_io_reference_grant_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_gateway_networking_k8s_io_reference_grant_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/gateway_networking_k8s_io_v1beta1/gateway_networking_k8s_io_reference_grant_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-gateway_nginx_org_client_settings_policy_v1alpha1_manifest_test.go: out/install-sentinel terratest/gateway_nginx_org_v1alpha1/gateway_nginx_org_client_settings_policy_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_gateway_nginx_org_client_settings_policy_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/gateway_nginx_org_v1alpha1/gateway_nginx_org_client_settings_policy_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-gateway_nginx_org_nginx_gateway_v1alpha1_manifest_test.go: out/install-sentinel terratest/gateway_nginx_org_v1alpha1/gateway_nginx_org_nginx_gateway_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_gateway_nginx_org_nginx_gateway_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/gateway_nginx_org_v1alpha1/gateway_nginx_org_nginx_gateway_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-gateway_nginx_org_nginx_proxy_v1alpha1_manifest_test.go: out/install-sentinel terratest/gateway_nginx_org_v1alpha1/gateway_nginx_org_nginx_proxy_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_gateway_nginx_org_nginx_proxy_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/gateway_nginx_org_v1alpha1/gateway_nginx_org_nginx_proxy_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-gateway_nginx_org_observability_policy_v1alpha1_manifest_test.go: out/install-sentinel terratest/gateway_nginx_org_v1alpha1/gateway_nginx_org_observability_policy_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_gateway_nginx_org_observability_policy_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/gateway_nginx_org_v1alpha1/gateway_nginx_org_observability_policy_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-gateway_solo_io_gateway_v1_manifest_test.go: out/install-sentinel terratest/gateway_solo_io_v1/gateway_solo_io_gateway_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_gateway_solo_io_gateway_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/gateway_solo_io_v1/gateway_solo_io_gateway_v1_manifest_test.go
	touch $@
out/terratest-sentinel-gateway_solo_io_matchable_http_gateway_v1_manifest_test.go: out/install-sentinel terratest/gateway_solo_io_v1/gateway_solo_io_matchable_http_gateway_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_gateway_solo_io_matchable_http_gateway_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/gateway_solo_io_v1/gateway_solo_io_matchable_http_gateway_v1_manifest_test.go
	touch $@
out/terratest-sentinel-gateway_solo_io_route_option_v1_manifest_test.go: out/install-sentinel terratest/gateway_solo_io_v1/gateway_solo_io_route_option_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_gateway_solo_io_route_option_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/gateway_solo_io_v1/gateway_solo_io_route_option_v1_manifest_test.go
	touch $@
out/terratest-sentinel-gateway_solo_io_route_table_v1_manifest_test.go: out/install-sentinel terratest/gateway_solo_io_v1/gateway_solo_io_route_table_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_gateway_solo_io_route_table_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/gateway_solo_io_v1/gateway_solo_io_route_table_v1_manifest_test.go
	touch $@
out/terratest-sentinel-gateway_solo_io_virtual_host_option_v1_manifest_test.go: out/install-sentinel terratest/gateway_solo_io_v1/gateway_solo_io_virtual_host_option_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_gateway_solo_io_virtual_host_option_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/gateway_solo_io_v1/gateway_solo_io_virtual_host_option_v1_manifest_test.go
	touch $@
out/terratest-sentinel-gateway_solo_io_virtual_service_v1_manifest_test.go: out/install-sentinel terratest/gateway_solo_io_v1/gateway_solo_io_virtual_service_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_gateway_solo_io_virtual_service_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/gateway_solo_io_v1/gateway_solo_io_virtual_service_v1_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_auth_service_v1_manifest_test.go: out/install-sentinel terratest/getambassador_io_v1/getambassador_io_auth_service_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_auth_service_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v1/getambassador_io_auth_service_v1_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_consul_resolver_v1_manifest_test.go: out/install-sentinel terratest/getambassador_io_v1/getambassador_io_consul_resolver_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_consul_resolver_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v1/getambassador_io_consul_resolver_v1_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_dev_portal_v1_manifest_test.go: out/install-sentinel terratest/getambassador_io_v1/getambassador_io_dev_portal_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_dev_portal_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v1/getambassador_io_dev_portal_v1_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_kubernetes_endpoint_resolver_v1_manifest_test.go: out/install-sentinel terratest/getambassador_io_v1/getambassador_io_kubernetes_endpoint_resolver_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_kubernetes_endpoint_resolver_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v1/getambassador_io_kubernetes_endpoint_resolver_v1_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_kubernetes_service_resolver_v1_manifest_test.go: out/install-sentinel terratest/getambassador_io_v1/getambassador_io_kubernetes_service_resolver_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_kubernetes_service_resolver_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v1/getambassador_io_kubernetes_service_resolver_v1_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_log_service_v1_manifest_test.go: out/install-sentinel terratest/getambassador_io_v1/getambassador_io_log_service_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_log_service_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v1/getambassador_io_log_service_v1_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_mapping_v1_manifest_test.go: out/install-sentinel terratest/getambassador_io_v1/getambassador_io_mapping_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_mapping_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v1/getambassador_io_mapping_v1_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_module_v1_manifest_test.go: out/install-sentinel terratest/getambassador_io_v1/getambassador_io_module_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_module_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v1/getambassador_io_module_v1_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_rate_limit_service_v1_manifest_test.go: out/install-sentinel terratest/getambassador_io_v1/getambassador_io_rate_limit_service_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_rate_limit_service_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v1/getambassador_io_rate_limit_service_v1_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_tcp_mapping_v1_manifest_test.go: out/install-sentinel terratest/getambassador_io_v1/getambassador_io_tcp_mapping_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_tcp_mapping_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v1/getambassador_io_tcp_mapping_v1_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_tls_context_v1_manifest_test.go: out/install-sentinel terratest/getambassador_io_v1/getambassador_io_tls_context_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_tls_context_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v1/getambassador_io_tls_context_v1_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_tracing_service_v1_manifest_test.go: out/install-sentinel terratest/getambassador_io_v1/getambassador_io_tracing_service_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_tracing_service_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v1/getambassador_io_tracing_service_v1_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_auth_service_v2_manifest_test.go: out/install-sentinel terratest/getambassador_io_v2/getambassador_io_auth_service_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_auth_service_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v2/getambassador_io_auth_service_v2_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_consul_resolver_v2_manifest_test.go: out/install-sentinel terratest/getambassador_io_v2/getambassador_io_consul_resolver_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_consul_resolver_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v2/getambassador_io_consul_resolver_v2_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_dev_portal_v2_manifest_test.go: out/install-sentinel terratest/getambassador_io_v2/getambassador_io_dev_portal_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_dev_portal_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v2/getambassador_io_dev_portal_v2_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_host_v2_manifest_test.go: out/install-sentinel terratest/getambassador_io_v2/getambassador_io_host_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_host_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v2/getambassador_io_host_v2_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_kubernetes_endpoint_resolver_v2_manifest_test.go: out/install-sentinel terratest/getambassador_io_v2/getambassador_io_kubernetes_endpoint_resolver_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_kubernetes_endpoint_resolver_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v2/getambassador_io_kubernetes_endpoint_resolver_v2_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_kubernetes_service_resolver_v2_manifest_test.go: out/install-sentinel terratest/getambassador_io_v2/getambassador_io_kubernetes_service_resolver_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_kubernetes_service_resolver_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v2/getambassador_io_kubernetes_service_resolver_v2_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_log_service_v2_manifest_test.go: out/install-sentinel terratest/getambassador_io_v2/getambassador_io_log_service_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_log_service_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v2/getambassador_io_log_service_v2_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_mapping_v2_manifest_test.go: out/install-sentinel terratest/getambassador_io_v2/getambassador_io_mapping_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_mapping_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v2/getambassador_io_mapping_v2_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_module_v2_manifest_test.go: out/install-sentinel terratest/getambassador_io_v2/getambassador_io_module_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_module_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v2/getambassador_io_module_v2_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_rate_limit_service_v2_manifest_test.go: out/install-sentinel terratest/getambassador_io_v2/getambassador_io_rate_limit_service_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_rate_limit_service_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v2/getambassador_io_rate_limit_service_v2_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_tcp_mapping_v2_manifest_test.go: out/install-sentinel terratest/getambassador_io_v2/getambassador_io_tcp_mapping_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_tcp_mapping_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v2/getambassador_io_tcp_mapping_v2_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_tls_context_v2_manifest_test.go: out/install-sentinel terratest/getambassador_io_v2/getambassador_io_tls_context_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_tls_context_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v2/getambassador_io_tls_context_v2_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_tracing_service_v2_manifest_test.go: out/install-sentinel terratest/getambassador_io_v2/getambassador_io_tracing_service_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_tracing_service_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v2/getambassador_io_tracing_service_v2_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_auth_service_v3alpha1_manifest_test.go: out/install-sentinel terratest/getambassador_io_v3alpha1/getambassador_io_auth_service_v3alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_auth_service_v3alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v3alpha1/getambassador_io_auth_service_v3alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_consul_resolver_v3alpha1_manifest_test.go: out/install-sentinel terratest/getambassador_io_v3alpha1/getambassador_io_consul_resolver_v3alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_consul_resolver_v3alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v3alpha1/getambassador_io_consul_resolver_v3alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_dev_portal_v3alpha1_manifest_test.go: out/install-sentinel terratest/getambassador_io_v3alpha1/getambassador_io_dev_portal_v3alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_dev_portal_v3alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v3alpha1/getambassador_io_dev_portal_v3alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_host_v3alpha1_manifest_test.go: out/install-sentinel terratest/getambassador_io_v3alpha1/getambassador_io_host_v3alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_host_v3alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v3alpha1/getambassador_io_host_v3alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_kubernetes_endpoint_resolver_v3alpha1_manifest_test.go: out/install-sentinel terratest/getambassador_io_v3alpha1/getambassador_io_kubernetes_endpoint_resolver_v3alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_kubernetes_endpoint_resolver_v3alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v3alpha1/getambassador_io_kubernetes_endpoint_resolver_v3alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_kubernetes_service_resolver_v3alpha1_manifest_test.go: out/install-sentinel terratest/getambassador_io_v3alpha1/getambassador_io_kubernetes_service_resolver_v3alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_kubernetes_service_resolver_v3alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v3alpha1/getambassador_io_kubernetes_service_resolver_v3alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_listener_v3alpha1_manifest_test.go: out/install-sentinel terratest/getambassador_io_v3alpha1/getambassador_io_listener_v3alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_listener_v3alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v3alpha1/getambassador_io_listener_v3alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_log_service_v3alpha1_manifest_test.go: out/install-sentinel terratest/getambassador_io_v3alpha1/getambassador_io_log_service_v3alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_log_service_v3alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v3alpha1/getambassador_io_log_service_v3alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_mapping_v3alpha1_manifest_test.go: out/install-sentinel terratest/getambassador_io_v3alpha1/getambassador_io_mapping_v3alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_mapping_v3alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v3alpha1/getambassador_io_mapping_v3alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_module_v3alpha1_manifest_test.go: out/install-sentinel terratest/getambassador_io_v3alpha1/getambassador_io_module_v3alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_module_v3alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v3alpha1/getambassador_io_module_v3alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_rate_limit_service_v3alpha1_manifest_test.go: out/install-sentinel terratest/getambassador_io_v3alpha1/getambassador_io_rate_limit_service_v3alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_rate_limit_service_v3alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v3alpha1/getambassador_io_rate_limit_service_v3alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_tcp_mapping_v3alpha1_manifest_test.go: out/install-sentinel terratest/getambassador_io_v3alpha1/getambassador_io_tcp_mapping_v3alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_tcp_mapping_v3alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v3alpha1/getambassador_io_tcp_mapping_v3alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_tls_context_v3alpha1_manifest_test.go: out/install-sentinel terratest/getambassador_io_v3alpha1/getambassador_io_tls_context_v3alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_tls_context_v3alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v3alpha1/getambassador_io_tls_context_v3alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-getambassador_io_tracing_service_v3alpha1_manifest_test.go: out/install-sentinel terratest/getambassador_io_v3alpha1/getambassador_io_tracing_service_v3alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_getambassador_io_tracing_service_v3alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/getambassador_io_v3alpha1/getambassador_io_tracing_service_v3alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-gitops_hybrid_cloud_patterns_io_pattern_v1alpha1_manifest_test.go: out/install-sentinel terratest/gitops_hybrid_cloud_patterns_io_v1alpha1/gitops_hybrid_cloud_patterns_io_pattern_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_gitops_hybrid_cloud_patterns_io_pattern_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/gitops_hybrid_cloud_patterns_io_v1alpha1/gitops_hybrid_cloud_patterns_io_pattern_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-gloo_solo_io_proxy_v1_manifest_test.go: out/install-sentinel terratest/gloo_solo_io_v1/gloo_solo_io_proxy_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_gloo_solo_io_proxy_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/gloo_solo_io_v1/gloo_solo_io_proxy_v1_manifest_test.go
	touch $@
out/terratest-sentinel-gloo_solo_io_settings_v1_manifest_test.go: out/install-sentinel terratest/gloo_solo_io_v1/gloo_solo_io_settings_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_gloo_solo_io_settings_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/gloo_solo_io_v1/gloo_solo_io_settings_v1_manifest_test.go
	touch $@
out/terratest-sentinel-gloo_solo_io_upstream_group_v1_manifest_test.go: out/install-sentinel terratest/gloo_solo_io_v1/gloo_solo_io_upstream_group_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_gloo_solo_io_upstream_group_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/gloo_solo_io_v1/gloo_solo_io_upstream_group_v1_manifest_test.go
	touch $@
out/terratest-sentinel-gloo_solo_io_upstream_v1_manifest_test.go: out/install-sentinel terratest/gloo_solo_io_v1/gloo_solo_io_upstream_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_gloo_solo_io_upstream_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/gloo_solo_io_v1/gloo_solo_io_upstream_v1_manifest_test.go
	touch $@
out/terratest-sentinel-grafana_integreatly_org_grafana_dashboard_v1beta1_manifest_test.go: out/install-sentinel terratest/grafana_integreatly_org_v1beta1/grafana_integreatly_org_grafana_dashboard_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_grafana_integreatly_org_grafana_dashboard_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/grafana_integreatly_org_v1beta1/grafana_integreatly_org_grafana_dashboard_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-grafana_integreatly_org_grafana_datasource_v1beta1_manifest_test.go: out/install-sentinel terratest/grafana_integreatly_org_v1beta1/grafana_integreatly_org_grafana_datasource_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_grafana_integreatly_org_grafana_datasource_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/grafana_integreatly_org_v1beta1/grafana_integreatly_org_grafana_datasource_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-grafana_integreatly_org_grafana_folder_v1beta1_manifest_test.go: out/install-sentinel terratest/grafana_integreatly_org_v1beta1/grafana_integreatly_org_grafana_folder_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_grafana_integreatly_org_grafana_folder_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/grafana_integreatly_org_v1beta1/grafana_integreatly_org_grafana_folder_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-grafana_integreatly_org_grafana_v1beta1_manifest_test.go: out/install-sentinel terratest/grafana_integreatly_org_v1beta1/grafana_integreatly_org_grafana_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_grafana_integreatly_org_grafana_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/grafana_integreatly_org_v1beta1/grafana_integreatly_org_grafana_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-graphql_gloo_solo_io_graph_ql_api_v1beta1_manifest_test.go: out/install-sentinel terratest/graphql_gloo_solo_io_v1beta1/graphql_gloo_solo_io_graph_ql_api_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_graphql_gloo_solo_io_graph_ql_api_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/graphql_gloo_solo_io_v1beta1/graphql_gloo_solo_io_graph_ql_api_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-groupsnapshot_storage_k8s_io_volume_group_snapshot_class_v1alpha1_manifest_test.go: out/install-sentinel terratest/groupsnapshot_storage_k8s_io_v1alpha1/groupsnapshot_storage_k8s_io_volume_group_snapshot_class_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_groupsnapshot_storage_k8s_io_volume_group_snapshot_class_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/groupsnapshot_storage_k8s_io_v1alpha1/groupsnapshot_storage_k8s_io_volume_group_snapshot_class_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-groupsnapshot_storage_k8s_io_volume_group_snapshot_content_v1alpha1_manifest_test.go: out/install-sentinel terratest/groupsnapshot_storage_k8s_io_v1alpha1/groupsnapshot_storage_k8s_io_volume_group_snapshot_content_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_groupsnapshot_storage_k8s_io_volume_group_snapshot_content_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/groupsnapshot_storage_k8s_io_v1alpha1/groupsnapshot_storage_k8s_io_volume_group_snapshot_content_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-groupsnapshot_storage_k8s_io_volume_group_snapshot_v1alpha1_manifest_test.go: out/install-sentinel terratest/groupsnapshot_storage_k8s_io_v1alpha1/groupsnapshot_storage_k8s_io_volume_group_snapshot_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_groupsnapshot_storage_k8s_io_volume_group_snapshot_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/groupsnapshot_storage_k8s_io_v1alpha1/groupsnapshot_storage_k8s_io_volume_group_snapshot_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-hazelcast_com_cron_hot_backup_v1alpha1_manifest_test.go: out/install-sentinel terratest/hazelcast_com_v1alpha1/hazelcast_com_cron_hot_backup_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_hazelcast_com_cron_hot_backup_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/hazelcast_com_v1alpha1/hazelcast_com_cron_hot_backup_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-hazelcast_com_hazelcast_v1alpha1_manifest_test.go: out/install-sentinel terratest/hazelcast_com_v1alpha1/hazelcast_com_hazelcast_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_hazelcast_com_hazelcast_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/hazelcast_com_v1alpha1/hazelcast_com_hazelcast_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-hazelcast_com_hot_backup_v1alpha1_manifest_test.go: out/install-sentinel terratest/hazelcast_com_v1alpha1/hazelcast_com_hot_backup_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_hazelcast_com_hot_backup_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/hazelcast_com_v1alpha1/hazelcast_com_hot_backup_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-hazelcast_com_management_center_v1alpha1_manifest_test.go: out/install-sentinel terratest/hazelcast_com_v1alpha1/hazelcast_com_management_center_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_hazelcast_com_management_center_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/hazelcast_com_v1alpha1/hazelcast_com_management_center_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-hazelcast_com_map_v1alpha1_manifest_test.go: out/install-sentinel terratest/hazelcast_com_v1alpha1/hazelcast_com_map_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_hazelcast_com_map_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/hazelcast_com_v1alpha1/hazelcast_com_map_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-hazelcast_com_wan_replication_v1alpha1_manifest_test.go: out/install-sentinel terratest/hazelcast_com_v1alpha1/hazelcast_com_wan_replication_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_hazelcast_com_wan_replication_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/hazelcast_com_v1alpha1/hazelcast_com_wan_replication_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-helm_sigstore_dev_rekor_v1alpha1_manifest_test.go: out/install-sentinel terratest/helm_sigstore_dev_v1alpha1/helm_sigstore_dev_rekor_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_helm_sigstore_dev_rekor_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/helm_sigstore_dev_v1alpha1/helm_sigstore_dev_rekor_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-helm_toolkit_fluxcd_io_helm_release_v2_manifest_test.go: out/install-sentinel terratest/helm_toolkit_fluxcd_io_v2/helm_toolkit_fluxcd_io_helm_release_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_helm_toolkit_fluxcd_io_helm_release_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/helm_toolkit_fluxcd_io_v2/helm_toolkit_fluxcd_io_helm_release_v2_manifest_test.go
	touch $@
out/terratest-sentinel-helm_toolkit_fluxcd_io_helm_release_v2beta1_manifest_test.go: out/install-sentinel terratest/helm_toolkit_fluxcd_io_v2beta1/helm_toolkit_fluxcd_io_helm_release_v2beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_helm_toolkit_fluxcd_io_helm_release_v2beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/helm_toolkit_fluxcd_io_v2beta1/helm_toolkit_fluxcd_io_helm_release_v2beta1_manifest_test.go
	touch $@
out/terratest-sentinel-helm_toolkit_fluxcd_io_helm_release_v2beta2_manifest_test.go: out/install-sentinel terratest/helm_toolkit_fluxcd_io_v2beta2/helm_toolkit_fluxcd_io_helm_release_v2beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_helm_toolkit_fluxcd_io_helm_release_v2beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/helm_toolkit_fluxcd_io_v2beta2/helm_toolkit_fluxcd_io_helm_release_v2beta2_manifest_test.go
	touch $@
out/terratest-sentinel-hive_openshift_io_checkpoint_v1_manifest_test.go: out/install-sentinel terratest/hive_openshift_io_v1/hive_openshift_io_checkpoint_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_hive_openshift_io_checkpoint_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/hive_openshift_io_v1/hive_openshift_io_checkpoint_v1_manifest_test.go
	touch $@
out/terratest-sentinel-hive_openshift_io_cluster_claim_v1_manifest_test.go: out/install-sentinel terratest/hive_openshift_io_v1/hive_openshift_io_cluster_claim_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_hive_openshift_io_cluster_claim_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/hive_openshift_io_v1/hive_openshift_io_cluster_claim_v1_manifest_test.go
	touch $@
out/terratest-sentinel-hive_openshift_io_cluster_deployment_customization_v1_manifest_test.go: out/install-sentinel terratest/hive_openshift_io_v1/hive_openshift_io_cluster_deployment_customization_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_hive_openshift_io_cluster_deployment_customization_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/hive_openshift_io_v1/hive_openshift_io_cluster_deployment_customization_v1_manifest_test.go
	touch $@
out/terratest-sentinel-hive_openshift_io_cluster_deployment_v1_manifest_test.go: out/install-sentinel terratest/hive_openshift_io_v1/hive_openshift_io_cluster_deployment_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_hive_openshift_io_cluster_deployment_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/hive_openshift_io_v1/hive_openshift_io_cluster_deployment_v1_manifest_test.go
	touch $@
out/terratest-sentinel-hive_openshift_io_cluster_deprovision_v1_manifest_test.go: out/install-sentinel terratest/hive_openshift_io_v1/hive_openshift_io_cluster_deprovision_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_hive_openshift_io_cluster_deprovision_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/hive_openshift_io_v1/hive_openshift_io_cluster_deprovision_v1_manifest_test.go
	touch $@
out/terratest-sentinel-hive_openshift_io_cluster_image_set_v1_manifest_test.go: out/install-sentinel terratest/hive_openshift_io_v1/hive_openshift_io_cluster_image_set_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_hive_openshift_io_cluster_image_set_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/hive_openshift_io_v1/hive_openshift_io_cluster_image_set_v1_manifest_test.go
	touch $@
out/terratest-sentinel-hive_openshift_io_cluster_pool_v1_manifest_test.go: out/install-sentinel terratest/hive_openshift_io_v1/hive_openshift_io_cluster_pool_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_hive_openshift_io_cluster_pool_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/hive_openshift_io_v1/hive_openshift_io_cluster_pool_v1_manifest_test.go
	touch $@
out/terratest-sentinel-hive_openshift_io_cluster_provision_v1_manifest_test.go: out/install-sentinel terratest/hive_openshift_io_v1/hive_openshift_io_cluster_provision_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_hive_openshift_io_cluster_provision_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/hive_openshift_io_v1/hive_openshift_io_cluster_provision_v1_manifest_test.go
	touch $@
out/terratest-sentinel-hive_openshift_io_cluster_relocate_v1_manifest_test.go: out/install-sentinel terratest/hive_openshift_io_v1/hive_openshift_io_cluster_relocate_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_hive_openshift_io_cluster_relocate_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/hive_openshift_io_v1/hive_openshift_io_cluster_relocate_v1_manifest_test.go
	touch $@
out/terratest-sentinel-hive_openshift_io_cluster_state_v1_manifest_test.go: out/install-sentinel terratest/hive_openshift_io_v1/hive_openshift_io_cluster_state_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_hive_openshift_io_cluster_state_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/hive_openshift_io_v1/hive_openshift_io_cluster_state_v1_manifest_test.go
	touch $@
out/terratest-sentinel-hive_openshift_io_dns_zone_v1_manifest_test.go: out/install-sentinel terratest/hive_openshift_io_v1/hive_openshift_io_dns_zone_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_hive_openshift_io_dns_zone_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/hive_openshift_io_v1/hive_openshift_io_dns_zone_v1_manifest_test.go
	touch $@
out/terratest-sentinel-hive_openshift_io_hive_config_v1_manifest_test.go: out/install-sentinel terratest/hive_openshift_io_v1/hive_openshift_io_hive_config_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_hive_openshift_io_hive_config_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/hive_openshift_io_v1/hive_openshift_io_hive_config_v1_manifest_test.go
	touch $@
out/terratest-sentinel-hive_openshift_io_machine_pool_name_lease_v1_manifest_test.go: out/install-sentinel terratest/hive_openshift_io_v1/hive_openshift_io_machine_pool_name_lease_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_hive_openshift_io_machine_pool_name_lease_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/hive_openshift_io_v1/hive_openshift_io_machine_pool_name_lease_v1_manifest_test.go
	touch $@
out/terratest-sentinel-hive_openshift_io_machine_pool_v1_manifest_test.go: out/install-sentinel terratest/hive_openshift_io_v1/hive_openshift_io_machine_pool_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_hive_openshift_io_machine_pool_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/hive_openshift_io_v1/hive_openshift_io_machine_pool_v1_manifest_test.go
	touch $@
out/terratest-sentinel-hive_openshift_io_selector_sync_identity_provider_v1_manifest_test.go: out/install-sentinel terratest/hive_openshift_io_v1/hive_openshift_io_selector_sync_identity_provider_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_hive_openshift_io_selector_sync_identity_provider_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/hive_openshift_io_v1/hive_openshift_io_selector_sync_identity_provider_v1_manifest_test.go
	touch $@
out/terratest-sentinel-hive_openshift_io_selector_sync_set_v1_manifest_test.go: out/install-sentinel terratest/hive_openshift_io_v1/hive_openshift_io_selector_sync_set_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_hive_openshift_io_selector_sync_set_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/hive_openshift_io_v1/hive_openshift_io_selector_sync_set_v1_manifest_test.go
	touch $@
out/terratest-sentinel-hive_openshift_io_sync_identity_provider_v1_manifest_test.go: out/install-sentinel terratest/hive_openshift_io_v1/hive_openshift_io_sync_identity_provider_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_hive_openshift_io_sync_identity_provider_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/hive_openshift_io_v1/hive_openshift_io_sync_identity_provider_v1_manifest_test.go
	touch $@
out/terratest-sentinel-hive_openshift_io_sync_set_v1_manifest_test.go: out/install-sentinel terratest/hive_openshift_io_v1/hive_openshift_io_sync_set_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_hive_openshift_io_sync_set_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/hive_openshift_io_v1/hive_openshift_io_sync_set_v1_manifest_test.go
	touch $@
out/terratest-sentinel-hiveinternal_openshift_io_cluster_sync_lease_v1alpha1_manifest_test.go: out/install-sentinel terratest/hiveinternal_openshift_io_v1alpha1/hiveinternal_openshift_io_cluster_sync_lease_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_hiveinternal_openshift_io_cluster_sync_lease_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/hiveinternal_openshift_io_v1alpha1/hiveinternal_openshift_io_cluster_sync_lease_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-hiveinternal_openshift_io_cluster_sync_v1alpha1_manifest_test.go: out/install-sentinel terratest/hiveinternal_openshift_io_v1alpha1/hiveinternal_openshift_io_cluster_sync_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_hiveinternal_openshift_io_cluster_sync_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/hiveinternal_openshift_io_v1alpha1/hiveinternal_openshift_io_cluster_sync_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-hiveinternal_openshift_io_fake_cluster_install_v1alpha1_manifest_test.go: out/install-sentinel terratest/hiveinternal_openshift_io_v1alpha1/hiveinternal_openshift_io_fake_cluster_install_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_hiveinternal_openshift_io_fake_cluster_install_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/hiveinternal_openshift_io_v1alpha1/hiveinternal_openshift_io_fake_cluster_install_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-hnc_x_k8s_io_hierarchical_resource_quota_v1alpha2_manifest_test.go: out/install-sentinel terratest/hnc_x_k8s_io_v1alpha2/hnc_x_k8s_io_hierarchical_resource_quota_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_hnc_x_k8s_io_hierarchical_resource_quota_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/hnc_x_k8s_io_v1alpha2/hnc_x_k8s_io_hierarchical_resource_quota_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-hnc_x_k8s_io_hierarchy_configuration_v1alpha2_manifest_test.go: out/install-sentinel terratest/hnc_x_k8s_io_v1alpha2/hnc_x_k8s_io_hierarchy_configuration_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_hnc_x_k8s_io_hierarchy_configuration_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/hnc_x_k8s_io_v1alpha2/hnc_x_k8s_io_hierarchy_configuration_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-hnc_x_k8s_io_hnc_configuration_v1alpha2_manifest_test.go: out/install-sentinel terratest/hnc_x_k8s_io_v1alpha2/hnc_x_k8s_io_hnc_configuration_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_hnc_x_k8s_io_hnc_configuration_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/hnc_x_k8s_io_v1alpha2/hnc_x_k8s_io_hnc_configuration_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-hnc_x_k8s_io_subnamespace_anchor_v1alpha2_manifest_test.go: out/install-sentinel terratest/hnc_x_k8s_io_v1alpha2/hnc_x_k8s_io_subnamespace_anchor_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_hnc_x_k8s_io_subnamespace_anchor_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/hnc_x_k8s_io_v1alpha2/hnc_x_k8s_io_subnamespace_anchor_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-hyperfoil_io_horreum_v1alpha1_manifest_test.go: out/install-sentinel terratest/hyperfoil_io_v1alpha1/hyperfoil_io_horreum_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_hyperfoil_io_horreum_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/hyperfoil_io_v1alpha1/hyperfoil_io_horreum_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-hyperfoil_io_hyperfoil_v1alpha2_manifest_test.go: out/install-sentinel terratest/hyperfoil_io_v1alpha2/hyperfoil_io_hyperfoil_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_hyperfoil_io_hyperfoil_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/hyperfoil_io_v1alpha2/hyperfoil_io_hyperfoil_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-iam_services_k8s_aws_group_v1alpha1_manifest_test.go: out/install-sentinel terratest/iam_services_k8s_aws_v1alpha1/iam_services_k8s_aws_group_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_iam_services_k8s_aws_group_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/iam_services_k8s_aws_v1alpha1/iam_services_k8s_aws_group_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-iam_services_k8s_aws_instance_profile_v1alpha1_manifest_test.go: out/install-sentinel terratest/iam_services_k8s_aws_v1alpha1/iam_services_k8s_aws_instance_profile_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_iam_services_k8s_aws_instance_profile_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/iam_services_k8s_aws_v1alpha1/iam_services_k8s_aws_instance_profile_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-iam_services_k8s_aws_open_id_connect_provider_v1alpha1_manifest_test.go: out/install-sentinel terratest/iam_services_k8s_aws_v1alpha1/iam_services_k8s_aws_open_id_connect_provider_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_iam_services_k8s_aws_open_id_connect_provider_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/iam_services_k8s_aws_v1alpha1/iam_services_k8s_aws_open_id_connect_provider_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-iam_services_k8s_aws_policy_v1alpha1_manifest_test.go: out/install-sentinel terratest/iam_services_k8s_aws_v1alpha1/iam_services_k8s_aws_policy_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_iam_services_k8s_aws_policy_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/iam_services_k8s_aws_v1alpha1/iam_services_k8s_aws_policy_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-iam_services_k8s_aws_role_v1alpha1_manifest_test.go: out/install-sentinel terratest/iam_services_k8s_aws_v1alpha1/iam_services_k8s_aws_role_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_iam_services_k8s_aws_role_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/iam_services_k8s_aws_v1alpha1/iam_services_k8s_aws_role_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-iam_services_k8s_aws_user_v1alpha1_manifest_test.go: out/install-sentinel terratest/iam_services_k8s_aws_v1alpha1/iam_services_k8s_aws_user_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_iam_services_k8s_aws_user_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/iam_services_k8s_aws_v1alpha1/iam_services_k8s_aws_user_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-ibmcloud_ibm_com_composable_v1alpha1_manifest_test.go: out/install-sentinel terratest/ibmcloud_ibm_com_v1alpha1/ibmcloud_ibm_com_composable_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_ibmcloud_ibm_com_composable_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ibmcloud_ibm_com_v1alpha1/ibmcloud_ibm_com_composable_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-image_toolkit_fluxcd_io_image_policy_v1beta1_manifest_test.go: out/install-sentinel terratest/image_toolkit_fluxcd_io_v1beta1/image_toolkit_fluxcd_io_image_policy_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_image_toolkit_fluxcd_io_image_policy_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/image_toolkit_fluxcd_io_v1beta1/image_toolkit_fluxcd_io_image_policy_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-image_toolkit_fluxcd_io_image_repository_v1beta1_manifest_test.go: out/install-sentinel terratest/image_toolkit_fluxcd_io_v1beta1/image_toolkit_fluxcd_io_image_repository_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_image_toolkit_fluxcd_io_image_repository_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/image_toolkit_fluxcd_io_v1beta1/image_toolkit_fluxcd_io_image_repository_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-image_toolkit_fluxcd_io_image_update_automation_v1beta1_manifest_test.go: out/install-sentinel terratest/image_toolkit_fluxcd_io_v1beta1/image_toolkit_fluxcd_io_image_update_automation_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_image_toolkit_fluxcd_io_image_update_automation_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/image_toolkit_fluxcd_io_v1beta1/image_toolkit_fluxcd_io_image_update_automation_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-image_toolkit_fluxcd_io_image_policy_v1beta2_manifest_test.go: out/install-sentinel terratest/image_toolkit_fluxcd_io_v1beta2/image_toolkit_fluxcd_io_image_policy_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_image_toolkit_fluxcd_io_image_policy_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/image_toolkit_fluxcd_io_v1beta2/image_toolkit_fluxcd_io_image_policy_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-image_toolkit_fluxcd_io_image_repository_v1beta2_manifest_test.go: out/install-sentinel terratest/image_toolkit_fluxcd_io_v1beta2/image_toolkit_fluxcd_io_image_repository_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_image_toolkit_fluxcd_io_image_repository_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/image_toolkit_fluxcd_io_v1beta2/image_toolkit_fluxcd_io_image_repository_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-image_toolkit_fluxcd_io_image_update_automation_v1beta2_manifest_test.go: out/install-sentinel terratest/image_toolkit_fluxcd_io_v1beta2/image_toolkit_fluxcd_io_image_update_automation_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_image_toolkit_fluxcd_io_image_update_automation_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/image_toolkit_fluxcd_io_v1beta2/image_toolkit_fluxcd_io_image_update_automation_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-imaging_ingestion_alvearie_org_dicom_event_bridge_v1alpha1_manifest_test.go: out/install-sentinel terratest/imaging_ingestion_alvearie_org_v1alpha1/imaging_ingestion_alvearie_org_dicom_event_bridge_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_imaging_ingestion_alvearie_org_dicom_event_bridge_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/imaging_ingestion_alvearie_org_v1alpha1/imaging_ingestion_alvearie_org_dicom_event_bridge_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-imaging_ingestion_alvearie_org_dicom_event_driven_ingestion_v1alpha1_manifest_test.go: out/install-sentinel terratest/imaging_ingestion_alvearie_org_v1alpha1/imaging_ingestion_alvearie_org_dicom_event_driven_ingestion_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_imaging_ingestion_alvearie_org_dicom_event_driven_ingestion_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/imaging_ingestion_alvearie_org_v1alpha1/imaging_ingestion_alvearie_org_dicom_event_driven_ingestion_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-imaging_ingestion_alvearie_org_dicom_instance_binding_v1alpha1_manifest_test.go: out/install-sentinel terratest/imaging_ingestion_alvearie_org_v1alpha1/imaging_ingestion_alvearie_org_dicom_instance_binding_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_imaging_ingestion_alvearie_org_dicom_instance_binding_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/imaging_ingestion_alvearie_org_v1alpha1/imaging_ingestion_alvearie_org_dicom_instance_binding_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-imaging_ingestion_alvearie_org_dicom_study_binding_v1alpha1_manifest_test.go: out/install-sentinel terratest/imaging_ingestion_alvearie_org_v1alpha1/imaging_ingestion_alvearie_org_dicom_study_binding_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_imaging_ingestion_alvearie_org_dicom_study_binding_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/imaging_ingestion_alvearie_org_v1alpha1/imaging_ingestion_alvearie_org_dicom_study_binding_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-imaging_ingestion_alvearie_org_dicomweb_ingestion_service_v1alpha1_manifest_test.go: out/install-sentinel terratest/imaging_ingestion_alvearie_org_v1alpha1/imaging_ingestion_alvearie_org_dicomweb_ingestion_service_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_imaging_ingestion_alvearie_org_dicomweb_ingestion_service_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/imaging_ingestion_alvearie_org_v1alpha1/imaging_ingestion_alvearie_org_dicomweb_ingestion_service_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-imaging_ingestion_alvearie_org_dimse_ingestion_service_v1alpha1_manifest_test.go: out/install-sentinel terratest/imaging_ingestion_alvearie_org_v1alpha1/imaging_ingestion_alvearie_org_dimse_ingestion_service_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_imaging_ingestion_alvearie_org_dimse_ingestion_service_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/imaging_ingestion_alvearie_org_v1alpha1/imaging_ingestion_alvearie_org_dimse_ingestion_service_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-imaging_ingestion_alvearie_org_dimse_proxy_v1alpha1_manifest_test.go: out/install-sentinel terratest/imaging_ingestion_alvearie_org_v1alpha1/imaging_ingestion_alvearie_org_dimse_proxy_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_imaging_ingestion_alvearie_org_dimse_proxy_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/imaging_ingestion_alvearie_org_v1alpha1/imaging_ingestion_alvearie_org_dimse_proxy_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-inference_kubedl_io_elastic_batch_job_v1alpha1_manifest_test.go: out/install-sentinel terratest/inference_kubedl_io_v1alpha1/inference_kubedl_io_elastic_batch_job_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_inference_kubedl_io_elastic_batch_job_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/inference_kubedl_io_v1alpha1/inference_kubedl_io_elastic_batch_job_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-infinispan_org_infinispan_v1_manifest_test.go: out/install-sentinel terratest/infinispan_org_v1/infinispan_org_infinispan_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_infinispan_org_infinispan_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infinispan_org_v1/infinispan_org_infinispan_v1_manifest_test.go
	touch $@
out/terratest-sentinel-infinispan_org_backup_v2alpha1_manifest_test.go: out/install-sentinel terratest/infinispan_org_v2alpha1/infinispan_org_backup_v2alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_infinispan_org_backup_v2alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infinispan_org_v2alpha1/infinispan_org_backup_v2alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-infinispan_org_batch_v2alpha1_manifest_test.go: out/install-sentinel terratest/infinispan_org_v2alpha1/infinispan_org_batch_v2alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_infinispan_org_batch_v2alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infinispan_org_v2alpha1/infinispan_org_batch_v2alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-infinispan_org_cache_v2alpha1_manifest_test.go: out/install-sentinel terratest/infinispan_org_v2alpha1/infinispan_org_cache_v2alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_infinispan_org_cache_v2alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infinispan_org_v2alpha1/infinispan_org_cache_v2alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-infinispan_org_restore_v2alpha1_manifest_test.go: out/install-sentinel terratest/infinispan_org_v2alpha1/infinispan_org_restore_v2alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_infinispan_org_restore_v2alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infinispan_org_v2alpha1/infinispan_org_restore_v2alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-infra_contrib_fluxcd_io_terraform_v1alpha1_manifest_test.go: out/install-sentinel terratest/infra_contrib_fluxcd_io_v1alpha1/infra_contrib_fluxcd_io_terraform_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_infra_contrib_fluxcd_io_terraform_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infra_contrib_fluxcd_io_v1alpha1/infra_contrib_fluxcd_io_terraform_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-infra_contrib_fluxcd_io_terraform_v1alpha2_manifest_test.go: out/install-sentinel terratest/infra_contrib_fluxcd_io_v1alpha2/infra_contrib_fluxcd_io_terraform_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_infra_contrib_fluxcd_io_terraform_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infra_contrib_fluxcd_io_v1alpha2/infra_contrib_fluxcd_io_terraform_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_kubevirt_cluster_template_v1alpha1_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1alpha1/infrastructure_cluster_x_k8s_io_kubevirt_cluster_template_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_kubevirt_cluster_template_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1alpha1/infrastructure_cluster_x_k8s_io_kubevirt_cluster_template_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_kubevirt_cluster_v1alpha1_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1alpha1/infrastructure_cluster_x_k8s_io_kubevirt_cluster_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_kubevirt_cluster_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1alpha1/infrastructure_cluster_x_k8s_io_kubevirt_cluster_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_kubevirt_machine_template_v1alpha1_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1alpha1/infrastructure_cluster_x_k8s_io_kubevirt_machine_template_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_kubevirt_machine_template_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1alpha1/infrastructure_cluster_x_k8s_io_kubevirt_machine_template_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_kubevirt_machine_v1alpha1_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1alpha1/infrastructure_cluster_x_k8s_io_kubevirt_machine_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_kubevirt_machine_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1alpha1/infrastructure_cluster_x_k8s_io_kubevirt_machine_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_cluster_identity_v1alpha3_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1alpha3/infrastructure_cluster_x_k8s_io_v_sphere_cluster_identity_v1alpha3_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_v_sphere_cluster_identity_v1alpha3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1alpha3/infrastructure_cluster_x_k8s_io_v_sphere_cluster_identity_v1alpha3_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_cluster_v1alpha3_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1alpha3/infrastructure_cluster_x_k8s_io_v_sphere_cluster_v1alpha3_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_v_sphere_cluster_v1alpha3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1alpha3/infrastructure_cluster_x_k8s_io_v_sphere_cluster_v1alpha3_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_deployment_zone_v1alpha3_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1alpha3/infrastructure_cluster_x_k8s_io_v_sphere_deployment_zone_v1alpha3_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_v_sphere_deployment_zone_v1alpha3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1alpha3/infrastructure_cluster_x_k8s_io_v_sphere_deployment_zone_v1alpha3_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_failure_domain_v1alpha3_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1alpha3/infrastructure_cluster_x_k8s_io_v_sphere_failure_domain_v1alpha3_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_v_sphere_failure_domain_v1alpha3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1alpha3/infrastructure_cluster_x_k8s_io_v_sphere_failure_domain_v1alpha3_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_machine_template_v1alpha3_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1alpha3/infrastructure_cluster_x_k8s_io_v_sphere_machine_template_v1alpha3_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_v_sphere_machine_template_v1alpha3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1alpha3/infrastructure_cluster_x_k8s_io_v_sphere_machine_template_v1alpha3_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_machine_v1alpha3_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1alpha3/infrastructure_cluster_x_k8s_io_v_sphere_machine_v1alpha3_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_v_sphere_machine_v1alpha3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1alpha3/infrastructure_cluster_x_k8s_io_v_sphere_machine_v1alpha3_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_vm_v1alpha3_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1alpha3/infrastructure_cluster_x_k8s_io_v_sphere_vm_v1alpha3_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_v_sphere_vm_v1alpha3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1alpha3/infrastructure_cluster_x_k8s_io_v_sphere_vm_v1alpha3_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_cluster_identity_v1alpha4_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1alpha4/infrastructure_cluster_x_k8s_io_v_sphere_cluster_identity_v1alpha4_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_v_sphere_cluster_identity_v1alpha4_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1alpha4/infrastructure_cluster_x_k8s_io_v_sphere_cluster_identity_v1alpha4_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_cluster_template_v1alpha4_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1alpha4/infrastructure_cluster_x_k8s_io_v_sphere_cluster_template_v1alpha4_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_v_sphere_cluster_template_v1alpha4_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1alpha4/infrastructure_cluster_x_k8s_io_v_sphere_cluster_template_v1alpha4_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_cluster_v1alpha4_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1alpha4/infrastructure_cluster_x_k8s_io_v_sphere_cluster_v1alpha4_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_v_sphere_cluster_v1alpha4_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1alpha4/infrastructure_cluster_x_k8s_io_v_sphere_cluster_v1alpha4_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_deployment_zone_v1alpha4_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1alpha4/infrastructure_cluster_x_k8s_io_v_sphere_deployment_zone_v1alpha4_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_v_sphere_deployment_zone_v1alpha4_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1alpha4/infrastructure_cluster_x_k8s_io_v_sphere_deployment_zone_v1alpha4_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_failure_domain_v1alpha4_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1alpha4/infrastructure_cluster_x_k8s_io_v_sphere_failure_domain_v1alpha4_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_v_sphere_failure_domain_v1alpha4_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1alpha4/infrastructure_cluster_x_k8s_io_v_sphere_failure_domain_v1alpha4_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_machine_template_v1alpha4_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1alpha4/infrastructure_cluster_x_k8s_io_v_sphere_machine_template_v1alpha4_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_v_sphere_machine_template_v1alpha4_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1alpha4/infrastructure_cluster_x_k8s_io_v_sphere_machine_template_v1alpha4_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_machine_v1alpha4_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1alpha4/infrastructure_cluster_x_k8s_io_v_sphere_machine_v1alpha4_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_v_sphere_machine_v1alpha4_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1alpha4/infrastructure_cluster_x_k8s_io_v_sphere_machine_v1alpha4_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_vm_v1alpha4_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1alpha4/infrastructure_cluster_x_k8s_io_v_sphere_vm_v1alpha4_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_v_sphere_vm_v1alpha4_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1alpha4/infrastructure_cluster_x_k8s_io_v_sphere_vm_v1alpha4_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_ibm_power_vs_cluster_template_v1beta1_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_ibm_power_vs_cluster_template_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_ibm_power_vs_cluster_template_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_ibm_power_vs_cluster_template_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_ibm_power_vs_cluster_v1beta1_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_ibm_power_vs_cluster_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_ibm_power_vs_cluster_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_ibm_power_vs_cluster_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_ibm_power_vs_image_v1beta1_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_ibm_power_vs_image_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_ibm_power_vs_image_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_ibm_power_vs_image_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_ibm_power_vs_machine_template_v1beta1_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_ibm_power_vs_machine_template_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_ibm_power_vs_machine_template_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_ibm_power_vs_machine_template_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_ibm_power_vs_machine_v1beta1_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_ibm_power_vs_machine_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_ibm_power_vs_machine_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_ibm_power_vs_machine_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_ibmvpc_cluster_v1beta1_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_ibmvpc_cluster_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_ibmvpc_cluster_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_ibmvpc_cluster_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_ibmvpc_machine_template_v1beta1_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_ibmvpc_machine_template_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_ibmvpc_machine_template_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_ibmvpc_machine_template_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_ibmvpc_machine_v1beta1_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_ibmvpc_machine_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_ibmvpc_machine_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_ibmvpc_machine_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_tinkerbell_cluster_v1beta1_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_tinkerbell_cluster_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_tinkerbell_cluster_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_tinkerbell_cluster_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_tinkerbell_machine_template_v1beta1_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_tinkerbell_machine_template_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_tinkerbell_machine_template_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_tinkerbell_machine_template_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_tinkerbell_machine_v1beta1_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_tinkerbell_machine_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_tinkerbell_machine_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_tinkerbell_machine_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_cluster_identity_v1beta1_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_v_sphere_cluster_identity_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_v_sphere_cluster_identity_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_v_sphere_cluster_identity_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_cluster_template_v1beta1_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_v_sphere_cluster_template_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_v_sphere_cluster_template_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_v_sphere_cluster_template_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_cluster_v1beta1_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_v_sphere_cluster_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_v_sphere_cluster_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_v_sphere_cluster_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_deployment_zone_v1beta1_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_v_sphere_deployment_zone_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_v_sphere_deployment_zone_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_v_sphere_deployment_zone_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_failure_domain_v1beta1_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_v_sphere_failure_domain_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_v_sphere_failure_domain_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_v_sphere_failure_domain_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_machine_template_v1beta1_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_v_sphere_machine_template_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_v_sphere_machine_template_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_v_sphere_machine_template_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_machine_v1beta1_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_v_sphere_machine_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_v_sphere_machine_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_v_sphere_machine_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_vm_v1beta1_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_v_sphere_vm_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_v_sphere_vm_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1beta1/infrastructure_cluster_x_k8s_io_v_sphere_vm_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_ibm_power_vs_cluster_template_v1beta2_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1beta2/infrastructure_cluster_x_k8s_io_ibm_power_vs_cluster_template_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_ibm_power_vs_cluster_template_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1beta2/infrastructure_cluster_x_k8s_io_ibm_power_vs_cluster_template_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_ibm_power_vs_cluster_v1beta2_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1beta2/infrastructure_cluster_x_k8s_io_ibm_power_vs_cluster_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_ibm_power_vs_cluster_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1beta2/infrastructure_cluster_x_k8s_io_ibm_power_vs_cluster_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_ibm_power_vs_image_v1beta2_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1beta2/infrastructure_cluster_x_k8s_io_ibm_power_vs_image_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_ibm_power_vs_image_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1beta2/infrastructure_cluster_x_k8s_io_ibm_power_vs_image_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_ibm_power_vs_machine_template_v1beta2_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1beta2/infrastructure_cluster_x_k8s_io_ibm_power_vs_machine_template_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_ibm_power_vs_machine_template_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1beta2/infrastructure_cluster_x_k8s_io_ibm_power_vs_machine_template_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_ibm_power_vs_machine_v1beta2_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1beta2/infrastructure_cluster_x_k8s_io_ibm_power_vs_machine_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_ibm_power_vs_machine_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1beta2/infrastructure_cluster_x_k8s_io_ibm_power_vs_machine_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_ibmvpc_cluster_v1beta2_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1beta2/infrastructure_cluster_x_k8s_io_ibmvpc_cluster_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_ibmvpc_cluster_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1beta2/infrastructure_cluster_x_k8s_io_ibmvpc_cluster_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_ibmvpc_machine_template_v1beta2_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1beta2/infrastructure_cluster_x_k8s_io_ibmvpc_machine_template_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_ibmvpc_machine_template_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1beta2/infrastructure_cluster_x_k8s_io_ibmvpc_machine_template_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-infrastructure_cluster_x_k8s_io_ibmvpc_machine_v1beta2_manifest_test.go: out/install-sentinel terratest/infrastructure_cluster_x_k8s_io_v1beta2/infrastructure_cluster_x_k8s_io_ibmvpc_machine_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_ibmvpc_machine_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/infrastructure_cluster_x_k8s_io_v1beta2/infrastructure_cluster_x_k8s_io_ibmvpc_machine_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-installation_mattermost_com_mattermost_v1beta1_manifest_test.go: out/install-sentinel terratest/installation_mattermost_com_v1beta1/installation_mattermost_com_mattermost_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_installation_mattermost_com_mattermost_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/installation_mattermost_com_v1beta1/installation_mattermost_com_mattermost_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-instana_io_instana_agent_v1_manifest_test.go: out/install-sentinel terratest/instana_io_v1/instana_io_instana_agent_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_instana_io_instana_agent_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/instana_io_v1/instana_io_instana_agent_v1_manifest_test.go
	touch $@
out/terratest-sentinel-integration_rock8s_com_deferred_resource_v1beta1_manifest_test.go: out/install-sentinel terratest/integration_rock8s_com_v1beta1/integration_rock8s_com_deferred_resource_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_integration_rock8s_com_deferred_resource_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/integration_rock8s_com_v1beta1/integration_rock8s_com_deferred_resource_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-integration_rock8s_com_plug_v1beta1_manifest_test.go: out/install-sentinel terratest/integration_rock8s_com_v1beta1/integration_rock8s_com_plug_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_integration_rock8s_com_plug_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/integration_rock8s_com_v1beta1/integration_rock8s_com_plug_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-integration_rock8s_com_socket_v1beta1_manifest_test.go: out/install-sentinel terratest/integration_rock8s_com_v1beta1/integration_rock8s_com_socket_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_integration_rock8s_com_socket_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/integration_rock8s_com_v1beta1/integration_rock8s_com_socket_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-iot_eclipse_org_ditto_v1alpha1_manifest_test.go: out/install-sentinel terratest/iot_eclipse_org_v1alpha1/iot_eclipse_org_ditto_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_iot_eclipse_org_ditto_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/iot_eclipse_org_v1alpha1/iot_eclipse_org_ditto_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-iot_eclipse_org_hawkbit_v1alpha1_manifest_test.go: out/install-sentinel terratest/iot_eclipse_org_v1alpha1/iot_eclipse_org_hawkbit_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_iot_eclipse_org_hawkbit_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/iot_eclipse_org_v1alpha1/iot_eclipse_org_hawkbit_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-ipam_cluster_x_k8s_io_ip_address_claim_v1alpha1_manifest_test.go: out/install-sentinel terratest/ipam_cluster_x_k8s_io_v1alpha1/ipam_cluster_x_k8s_io_ip_address_claim_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_ipam_cluster_x_k8s_io_ip_address_claim_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ipam_cluster_x_k8s_io_v1alpha1/ipam_cluster_x_k8s_io_ip_address_claim_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-ipam_cluster_x_k8s_io_ip_address_v1alpha1_manifest_test.go: out/install-sentinel terratest/ipam_cluster_x_k8s_io_v1alpha1/ipam_cluster_x_k8s_io_ip_address_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_ipam_cluster_x_k8s_io_ip_address_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ipam_cluster_x_k8s_io_v1alpha1/ipam_cluster_x_k8s_io_ip_address_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-ipam_cluster_x_k8s_io_ip_address_claim_v1beta1_manifest_test.go: out/install-sentinel terratest/ipam_cluster_x_k8s_io_v1beta1/ipam_cluster_x_k8s_io_ip_address_claim_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_ipam_cluster_x_k8s_io_ip_address_claim_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ipam_cluster_x_k8s_io_v1beta1/ipam_cluster_x_k8s_io_ip_address_claim_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-ipam_cluster_x_k8s_io_ip_address_v1beta1_manifest_test.go: out/install-sentinel terratest/ipam_cluster_x_k8s_io_v1beta1/ipam_cluster_x_k8s_io_ip_address_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_ipam_cluster_x_k8s_io_ip_address_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ipam_cluster_x_k8s_io_v1beta1/ipam_cluster_x_k8s_io_ip_address_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-isindir_github_com_sops_secret_v1alpha1_manifest_test.go: out/install-sentinel terratest/isindir_github_com_v1alpha1/isindir_github_com_sops_secret_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_isindir_github_com_sops_secret_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/isindir_github_com_v1alpha1/isindir_github_com_sops_secret_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-isindir_github_com_sops_secret_v1alpha2_manifest_test.go: out/install-sentinel terratest/isindir_github_com_v1alpha2/isindir_github_com_sops_secret_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_isindir_github_com_sops_secret_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/isindir_github_com_v1alpha2/isindir_github_com_sops_secret_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-isindir_github_com_sops_secret_v1alpha3_manifest_test.go: out/install-sentinel terratest/isindir_github_com_v1alpha3/isindir_github_com_sops_secret_v1alpha3_manifest_test.go $(shell find ./examples/data-sources/k8s_isindir_github_com_sops_secret_v1alpha3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/isindir_github_com_v1alpha3/isindir_github_com_sops_secret_v1alpha3_manifest_test.go
	touch $@
out/terratest-sentinel-jaegertracing_io_jaeger_v1_manifest_test.go: out/install-sentinel terratest/jaegertracing_io_v1/jaegertracing_io_jaeger_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_jaegertracing_io_jaeger_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/jaegertracing_io_v1/jaegertracing_io_jaeger_v1_manifest_test.go
	touch $@
out/terratest-sentinel-jobset_x_k8s_io_job_set_v1alpha2_manifest_test.go: out/install-sentinel terratest/jobset_x_k8s_io_v1alpha2/jobset_x_k8s_io_job_set_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_jobset_x_k8s_io_job_set_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/jobset_x_k8s_io_v1alpha2/jobset_x_k8s_io_job_set_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-jobsmanager_raczylo_com_managed_job_v1beta1_manifest_test.go: out/install-sentinel terratest/jobsmanager_raczylo_com_v1beta1/jobsmanager_raczylo_com_managed_job_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_jobsmanager_raczylo_com_managed_job_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/jobsmanager_raczylo_com_v1beta1/jobsmanager_raczylo_com_managed_job_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-k6_io_k6_v1alpha1_manifest_test.go: out/install-sentinel terratest/k6_io_v1alpha1/k6_io_k6_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_k6_io_k6_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k6_io_v1alpha1/k6_io_k6_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-k6_io_private_load_zone_v1alpha1_manifest_test.go: out/install-sentinel terratest/k6_io_v1alpha1/k6_io_private_load_zone_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_k6_io_private_load_zone_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k6_io_v1alpha1/k6_io_private_load_zone_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-k6_io_test_run_v1alpha1_manifest_test.go: out/install-sentinel terratest/k6_io_v1alpha1/k6_io_test_run_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_k6_io_test_run_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k6_io_v1alpha1/k6_io_test_run_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-k8gb_absa_oss_gslb_v1beta1_manifest_test.go: out/install-sentinel terratest/k8gb_absa_oss_v1beta1/k8gb_absa_oss_gslb_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_k8gb_absa_oss_gslb_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k8gb_absa_oss_v1beta1/k8gb_absa_oss_gslb_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-k8s_keycloak_org_keycloak_realm_import_v2alpha1_manifest_test.go: out/install-sentinel terratest/k8s_keycloak_org_v2alpha1/k8s_keycloak_org_keycloak_realm_import_v2alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_k8s_keycloak_org_keycloak_realm_import_v2alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k8s_keycloak_org_v2alpha1/k8s_keycloak_org_keycloak_realm_import_v2alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-k8s_keycloak_org_keycloak_v2alpha1_manifest_test.go: out/install-sentinel terratest/k8s_keycloak_org_v2alpha1/k8s_keycloak_org_keycloak_v2alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_k8s_keycloak_org_keycloak_v2alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k8s_keycloak_org_v2alpha1/k8s_keycloak_org_keycloak_v2alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-k8s_mariadb_com_backup_v1alpha1_manifest_test.go: out/install-sentinel terratest/k8s_mariadb_com_v1alpha1/k8s_mariadb_com_backup_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_k8s_mariadb_com_backup_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k8s_mariadb_com_v1alpha1/k8s_mariadb_com_backup_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-k8s_mariadb_com_connection_v1alpha1_manifest_test.go: out/install-sentinel terratest/k8s_mariadb_com_v1alpha1/k8s_mariadb_com_connection_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_k8s_mariadb_com_connection_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k8s_mariadb_com_v1alpha1/k8s_mariadb_com_connection_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-k8s_mariadb_com_database_v1alpha1_manifest_test.go: out/install-sentinel terratest/k8s_mariadb_com_v1alpha1/k8s_mariadb_com_database_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_k8s_mariadb_com_database_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k8s_mariadb_com_v1alpha1/k8s_mariadb_com_database_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-k8s_mariadb_com_grant_v1alpha1_manifest_test.go: out/install-sentinel terratest/k8s_mariadb_com_v1alpha1/k8s_mariadb_com_grant_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_k8s_mariadb_com_grant_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k8s_mariadb_com_v1alpha1/k8s_mariadb_com_grant_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-k8s_mariadb_com_maria_db_v1alpha1_manifest_test.go: out/install-sentinel terratest/k8s_mariadb_com_v1alpha1/k8s_mariadb_com_maria_db_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_k8s_mariadb_com_maria_db_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k8s_mariadb_com_v1alpha1/k8s_mariadb_com_maria_db_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-k8s_mariadb_com_max_scale_v1alpha1_manifest_test.go: out/install-sentinel terratest/k8s_mariadb_com_v1alpha1/k8s_mariadb_com_max_scale_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_k8s_mariadb_com_max_scale_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k8s_mariadb_com_v1alpha1/k8s_mariadb_com_max_scale_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-k8s_mariadb_com_restore_v1alpha1_manifest_test.go: out/install-sentinel terratest/k8s_mariadb_com_v1alpha1/k8s_mariadb_com_restore_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_k8s_mariadb_com_restore_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k8s_mariadb_com_v1alpha1/k8s_mariadb_com_restore_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-k8s_mariadb_com_sql_job_v1alpha1_manifest_test.go: out/install-sentinel terratest/k8s_mariadb_com_v1alpha1/k8s_mariadb_com_sql_job_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_k8s_mariadb_com_sql_job_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k8s_mariadb_com_v1alpha1/k8s_mariadb_com_sql_job_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-k8s_mariadb_com_user_v1alpha1_manifest_test.go: out/install-sentinel terratest/k8s_mariadb_com_v1alpha1/k8s_mariadb_com_user_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_k8s_mariadb_com_user_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k8s_mariadb_com_v1alpha1/k8s_mariadb_com_user_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-k8s_nginx_org_global_configuration_v1_manifest_test.go: out/install-sentinel terratest/k8s_nginx_org_v1/k8s_nginx_org_global_configuration_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_k8s_nginx_org_global_configuration_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k8s_nginx_org_v1/k8s_nginx_org_global_configuration_v1_manifest_test.go
	touch $@
out/terratest-sentinel-k8s_nginx_org_policy_v1_manifest_test.go: out/install-sentinel terratest/k8s_nginx_org_v1/k8s_nginx_org_policy_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_k8s_nginx_org_policy_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k8s_nginx_org_v1/k8s_nginx_org_policy_v1_manifest_test.go
	touch $@
out/terratest-sentinel-k8s_nginx_org_transport_server_v1_manifest_test.go: out/install-sentinel terratest/k8s_nginx_org_v1/k8s_nginx_org_transport_server_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_k8s_nginx_org_transport_server_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k8s_nginx_org_v1/k8s_nginx_org_transport_server_v1_manifest_test.go
	touch $@
out/terratest-sentinel-k8s_nginx_org_virtual_server_route_v1_manifest_test.go: out/install-sentinel terratest/k8s_nginx_org_v1/k8s_nginx_org_virtual_server_route_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_k8s_nginx_org_virtual_server_route_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k8s_nginx_org_v1/k8s_nginx_org_virtual_server_route_v1_manifest_test.go
	touch $@
out/terratest-sentinel-k8s_nginx_org_virtual_server_v1_manifest_test.go: out/install-sentinel terratest/k8s_nginx_org_v1/k8s_nginx_org_virtual_server_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_k8s_nginx_org_virtual_server_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k8s_nginx_org_v1/k8s_nginx_org_virtual_server_v1_manifest_test.go
	touch $@
out/terratest-sentinel-k8s_nginx_org_global_configuration_v1alpha1_manifest_test.go: out/install-sentinel terratest/k8s_nginx_org_v1alpha1/k8s_nginx_org_global_configuration_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_k8s_nginx_org_global_configuration_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k8s_nginx_org_v1alpha1/k8s_nginx_org_global_configuration_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-k8s_nginx_org_policy_v1alpha1_manifest_test.go: out/install-sentinel terratest/k8s_nginx_org_v1alpha1/k8s_nginx_org_policy_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_k8s_nginx_org_policy_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k8s_nginx_org_v1alpha1/k8s_nginx_org_policy_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-k8s_nginx_org_transport_server_v1alpha1_manifest_test.go: out/install-sentinel terratest/k8s_nginx_org_v1alpha1/k8s_nginx_org_transport_server_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_k8s_nginx_org_transport_server_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k8s_nginx_org_v1alpha1/k8s_nginx_org_transport_server_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-k8s_otterize_com_client_intents_v1alpha2_manifest_test.go: out/install-sentinel terratest/k8s_otterize_com_v1alpha2/k8s_otterize_com_client_intents_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_k8s_otterize_com_client_intents_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k8s_otterize_com_v1alpha2/k8s_otterize_com_client_intents_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-k8s_otterize_com_kafka_server_config_v1alpha2_manifest_test.go: out/install-sentinel terratest/k8s_otterize_com_v1alpha2/k8s_otterize_com_kafka_server_config_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_k8s_otterize_com_kafka_server_config_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k8s_otterize_com_v1alpha2/k8s_otterize_com_kafka_server_config_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-k8s_otterize_com_protected_service_v1alpha2_manifest_test.go: out/install-sentinel terratest/k8s_otterize_com_v1alpha2/k8s_otterize_com_protected_service_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_k8s_otterize_com_protected_service_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k8s_otterize_com_v1alpha2/k8s_otterize_com_protected_service_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-k8s_otterize_com_client_intents_v1alpha3_manifest_test.go: out/install-sentinel terratest/k8s_otterize_com_v1alpha3/k8s_otterize_com_client_intents_v1alpha3_manifest_test.go $(shell find ./examples/data-sources/k8s_k8s_otterize_com_client_intents_v1alpha3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k8s_otterize_com_v1alpha3/k8s_otterize_com_client_intents_v1alpha3_manifest_test.go
	touch $@
out/terratest-sentinel-k8s_otterize_com_kafka_server_config_v1alpha3_manifest_test.go: out/install-sentinel terratest/k8s_otterize_com_v1alpha3/k8s_otterize_com_kafka_server_config_v1alpha3_manifest_test.go $(shell find ./examples/data-sources/k8s_k8s_otterize_com_kafka_server_config_v1alpha3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k8s_otterize_com_v1alpha3/k8s_otterize_com_kafka_server_config_v1alpha3_manifest_test.go
	touch $@
out/terratest-sentinel-k8s_otterize_com_protected_service_v1alpha3_manifest_test.go: out/install-sentinel terratest/k8s_otterize_com_v1alpha3/k8s_otterize_com_protected_service_v1alpha3_manifest_test.go $(shell find ./examples/data-sources/k8s_k8s_otterize_com_protected_service_v1alpha3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k8s_otterize_com_v1alpha3/k8s_otterize_com_protected_service_v1alpha3_manifest_test.go
	touch $@
out/terratest-sentinel-k8up_io_archive_v1_manifest_test.go: out/install-sentinel terratest/k8up_io_v1/k8up_io_archive_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_k8up_io_archive_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k8up_io_v1/k8up_io_archive_v1_manifest_test.go
	touch $@
out/terratest-sentinel-k8up_io_backup_v1_manifest_test.go: out/install-sentinel terratest/k8up_io_v1/k8up_io_backup_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_k8up_io_backup_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k8up_io_v1/k8up_io_backup_v1_manifest_test.go
	touch $@
out/terratest-sentinel-k8up_io_check_v1_manifest_test.go: out/install-sentinel terratest/k8up_io_v1/k8up_io_check_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_k8up_io_check_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k8up_io_v1/k8up_io_check_v1_manifest_test.go
	touch $@
out/terratest-sentinel-k8up_io_pre_backup_pod_v1_manifest_test.go: out/install-sentinel terratest/k8up_io_v1/k8up_io_pre_backup_pod_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_k8up_io_pre_backup_pod_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k8up_io_v1/k8up_io_pre_backup_pod_v1_manifest_test.go
	touch $@
out/terratest-sentinel-k8up_io_prune_v1_manifest_test.go: out/install-sentinel terratest/k8up_io_v1/k8up_io_prune_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_k8up_io_prune_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k8up_io_v1/k8up_io_prune_v1_manifest_test.go
	touch $@
out/terratest-sentinel-k8up_io_restore_v1_manifest_test.go: out/install-sentinel terratest/k8up_io_v1/k8up_io_restore_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_k8up_io_restore_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k8up_io_v1/k8up_io_restore_v1_manifest_test.go
	touch $@
out/terratest-sentinel-k8up_io_schedule_v1_manifest_test.go: out/install-sentinel terratest/k8up_io_v1/k8up_io_schedule_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_k8up_io_schedule_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k8up_io_v1/k8up_io_schedule_v1_manifest_test.go
	touch $@
out/terratest-sentinel-k8up_io_snapshot_v1_manifest_test.go: out/install-sentinel terratest/k8up_io_v1/k8up_io_snapshot_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_k8up_io_snapshot_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/k8up_io_v1/k8up_io_snapshot_v1_manifest_test.go
	touch $@
out/terratest-sentinel-kafka_banzaicloud_io_cruise_control_operation_v1alpha1_manifest_test.go: out/install-sentinel terratest/kafka_banzaicloud_io_v1alpha1/kafka_banzaicloud_io_cruise_control_operation_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kafka_banzaicloud_io_cruise_control_operation_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kafka_banzaicloud_io_v1alpha1/kafka_banzaicloud_io_cruise_control_operation_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kafka_banzaicloud_io_kafka_topic_v1alpha1_manifest_test.go: out/install-sentinel terratest/kafka_banzaicloud_io_v1alpha1/kafka_banzaicloud_io_kafka_topic_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kafka_banzaicloud_io_kafka_topic_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kafka_banzaicloud_io_v1alpha1/kafka_banzaicloud_io_kafka_topic_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kafka_banzaicloud_io_kafka_user_v1alpha1_manifest_test.go: out/install-sentinel terratest/kafka_banzaicloud_io_v1alpha1/kafka_banzaicloud_io_kafka_user_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kafka_banzaicloud_io_kafka_user_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kafka_banzaicloud_io_v1alpha1/kafka_banzaicloud_io_kafka_user_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kafka_banzaicloud_io_kafka_cluster_v1beta1_manifest_test.go: out/install-sentinel terratest/kafka_banzaicloud_io_v1beta1/kafka_banzaicloud_io_kafka_cluster_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_kafka_banzaicloud_io_kafka_cluster_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kafka_banzaicloud_io_v1beta1/kafka_banzaicloud_io_kafka_cluster_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-kafka_services_k8s_aws_cluster_v1alpha1_manifest_test.go: out/install-sentinel terratest/kafka_services_k8s_aws_v1alpha1/kafka_services_k8s_aws_cluster_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kafka_services_k8s_aws_cluster_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kafka_services_k8s_aws_v1alpha1/kafka_services_k8s_aws_cluster_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kafka_strimzi_io_kafka_topic_v1alpha1_manifest_test.go: out/install-sentinel terratest/kafka_strimzi_io_v1alpha1/kafka_strimzi_io_kafka_topic_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kafka_strimzi_io_kafka_topic_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kafka_strimzi_io_v1alpha1/kafka_strimzi_io_kafka_topic_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kafka_strimzi_io_kafka_user_v1alpha1_manifest_test.go: out/install-sentinel terratest/kafka_strimzi_io_v1alpha1/kafka_strimzi_io_kafka_user_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kafka_strimzi_io_kafka_user_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kafka_strimzi_io_v1alpha1/kafka_strimzi_io_kafka_user_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kafka_strimzi_io_kafka_topic_v1beta1_manifest_test.go: out/install-sentinel terratest/kafka_strimzi_io_v1beta1/kafka_strimzi_io_kafka_topic_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_kafka_strimzi_io_kafka_topic_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kafka_strimzi_io_v1beta1/kafka_strimzi_io_kafka_topic_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-kafka_strimzi_io_kafka_user_v1beta1_manifest_test.go: out/install-sentinel terratest/kafka_strimzi_io_v1beta1/kafka_strimzi_io_kafka_user_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_kafka_strimzi_io_kafka_user_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kafka_strimzi_io_v1beta1/kafka_strimzi_io_kafka_user_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-kafka_strimzi_io_kafka_bridge_v1beta2_manifest_test.go: out/install-sentinel terratest/kafka_strimzi_io_v1beta2/kafka_strimzi_io_kafka_bridge_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_kafka_strimzi_io_kafka_bridge_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kafka_strimzi_io_v1beta2/kafka_strimzi_io_kafka_bridge_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-kafka_strimzi_io_kafka_connect_v1beta2_manifest_test.go: out/install-sentinel terratest/kafka_strimzi_io_v1beta2/kafka_strimzi_io_kafka_connect_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_kafka_strimzi_io_kafka_connect_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kafka_strimzi_io_v1beta2/kafka_strimzi_io_kafka_connect_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-kafka_strimzi_io_kafka_connector_v1beta2_manifest_test.go: out/install-sentinel terratest/kafka_strimzi_io_v1beta2/kafka_strimzi_io_kafka_connector_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_kafka_strimzi_io_kafka_connector_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kafka_strimzi_io_v1beta2/kafka_strimzi_io_kafka_connector_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-kafka_strimzi_io_kafka_mirror_maker2_v1beta2_manifest_test.go: out/install-sentinel terratest/kafka_strimzi_io_v1beta2/kafka_strimzi_io_kafka_mirror_maker2_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_kafka_strimzi_io_kafka_mirror_maker2_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kafka_strimzi_io_v1beta2/kafka_strimzi_io_kafka_mirror_maker2_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-kafka_strimzi_io_kafka_mirror_maker_v1beta2_manifest_test.go: out/install-sentinel terratest/kafka_strimzi_io_v1beta2/kafka_strimzi_io_kafka_mirror_maker_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_kafka_strimzi_io_kafka_mirror_maker_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kafka_strimzi_io_v1beta2/kafka_strimzi_io_kafka_mirror_maker_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-kafka_strimzi_io_kafka_rebalance_v1beta2_manifest_test.go: out/install-sentinel terratest/kafka_strimzi_io_v1beta2/kafka_strimzi_io_kafka_rebalance_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_kafka_strimzi_io_kafka_rebalance_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kafka_strimzi_io_v1beta2/kafka_strimzi_io_kafka_rebalance_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-kafka_strimzi_io_kafka_topic_v1beta2_manifest_test.go: out/install-sentinel terratest/kafka_strimzi_io_v1beta2/kafka_strimzi_io_kafka_topic_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_kafka_strimzi_io_kafka_topic_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kafka_strimzi_io_v1beta2/kafka_strimzi_io_kafka_topic_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-kafka_strimzi_io_kafka_user_v1beta2_manifest_test.go: out/install-sentinel terratest/kafka_strimzi_io_v1beta2/kafka_strimzi_io_kafka_user_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_kafka_strimzi_io_kafka_user_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kafka_strimzi_io_v1beta2/kafka_strimzi_io_kafka_user_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-kafka_strimzi_io_kafka_v1beta2_manifest_test.go: out/install-sentinel terratest/kafka_strimzi_io_v1beta2/kafka_strimzi_io_kafka_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_kafka_strimzi_io_kafka_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kafka_strimzi_io_v1beta2/kafka_strimzi_io_kafka_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-kamaji_clastix_io_data_store_v1alpha1_manifest_test.go: out/install-sentinel terratest/kamaji_clastix_io_v1alpha1/kamaji_clastix_io_data_store_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kamaji_clastix_io_data_store_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kamaji_clastix_io_v1alpha1/kamaji_clastix_io_data_store_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kamaji_clastix_io_tenant_control_plane_v1alpha1_manifest_test.go: out/install-sentinel terratest/kamaji_clastix_io_v1alpha1/kamaji_clastix_io_tenant_control_plane_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kamaji_clastix_io_tenant_control_plane_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kamaji_clastix_io_v1alpha1/kamaji_clastix_io_tenant_control_plane_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-karpenter_k8s_aws_ec2_node_class_v1_manifest_test.go: out/install-sentinel terratest/karpenter_k8s_aws_v1/karpenter_k8s_aws_ec2_node_class_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_karpenter_k8s_aws_ec2_node_class_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/karpenter_k8s_aws_v1/karpenter_k8s_aws_ec2_node_class_v1_manifest_test.go
	touch $@
out/terratest-sentinel-karpenter_k8s_aws_ec2_node_class_v1beta1_manifest_test.go: out/install-sentinel terratest/karpenter_k8s_aws_v1beta1/karpenter_k8s_aws_ec2_node_class_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_karpenter_k8s_aws_ec2_node_class_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/karpenter_k8s_aws_v1beta1/karpenter_k8s_aws_ec2_node_class_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-karpenter_sh_node_claim_v1_manifest_test.go: out/install-sentinel terratest/karpenter_sh_v1/karpenter_sh_node_claim_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_karpenter_sh_node_claim_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/karpenter_sh_v1/karpenter_sh_node_claim_v1_manifest_test.go
	touch $@
out/terratest-sentinel-karpenter_sh_node_pool_v1_manifest_test.go: out/install-sentinel terratest/karpenter_sh_v1/karpenter_sh_node_pool_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_karpenter_sh_node_pool_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/karpenter_sh_v1/karpenter_sh_node_pool_v1_manifest_test.go
	touch $@
out/terratest-sentinel-karpenter_sh_node_claim_v1beta1_manifest_test.go: out/install-sentinel terratest/karpenter_sh_v1beta1/karpenter_sh_node_claim_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_karpenter_sh_node_claim_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/karpenter_sh_v1beta1/karpenter_sh_node_claim_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-karpenter_sh_node_pool_v1beta1_manifest_test.go: out/install-sentinel terratest/karpenter_sh_v1beta1/karpenter_sh_node_pool_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_karpenter_sh_node_pool_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/karpenter_sh_v1beta1/karpenter_sh_node_pool_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-keda_sh_cluster_trigger_authentication_v1alpha1_manifest_test.go: out/install-sentinel terratest/keda_sh_v1alpha1/keda_sh_cluster_trigger_authentication_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_keda_sh_cluster_trigger_authentication_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/keda_sh_v1alpha1/keda_sh_cluster_trigger_authentication_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-keda_sh_scaled_job_v1alpha1_manifest_test.go: out/install-sentinel terratest/keda_sh_v1alpha1/keda_sh_scaled_job_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_keda_sh_scaled_job_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/keda_sh_v1alpha1/keda_sh_scaled_job_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-keda_sh_scaled_object_v1alpha1_manifest_test.go: out/install-sentinel terratest/keda_sh_v1alpha1/keda_sh_scaled_object_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_keda_sh_scaled_object_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/keda_sh_v1alpha1/keda_sh_scaled_object_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-keda_sh_trigger_authentication_v1alpha1_manifest_test.go: out/install-sentinel terratest/keda_sh_v1alpha1/keda_sh_trigger_authentication_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_keda_sh_trigger_authentication_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/keda_sh_v1alpha1/keda_sh_trigger_authentication_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-keycloak_k8s_reddec_net_keycloak_client_v1alpha1_manifest_test.go: out/install-sentinel terratest/keycloak_k8s_reddec_net_v1alpha1/keycloak_k8s_reddec_net_keycloak_client_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_keycloak_k8s_reddec_net_keycloak_client_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/keycloak_k8s_reddec_net_v1alpha1/keycloak_k8s_reddec_net_keycloak_client_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-keycloak_org_keycloak_backup_v1alpha1_manifest_test.go: out/install-sentinel terratest/keycloak_org_v1alpha1/keycloak_org_keycloak_backup_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_keycloak_org_keycloak_backup_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/keycloak_org_v1alpha1/keycloak_org_keycloak_backup_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-keycloak_org_keycloak_client_v1alpha1_manifest_test.go: out/install-sentinel terratest/keycloak_org_v1alpha1/keycloak_org_keycloak_client_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_keycloak_org_keycloak_client_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/keycloak_org_v1alpha1/keycloak_org_keycloak_client_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-keycloak_org_keycloak_realm_v1alpha1_manifest_test.go: out/install-sentinel terratest/keycloak_org_v1alpha1/keycloak_org_keycloak_realm_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_keycloak_org_keycloak_realm_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/keycloak_org_v1alpha1/keycloak_org_keycloak_realm_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-keycloak_org_keycloak_user_v1alpha1_manifest_test.go: out/install-sentinel terratest/keycloak_org_v1alpha1/keycloak_org_keycloak_user_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_keycloak_org_keycloak_user_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/keycloak_org_v1alpha1/keycloak_org_keycloak_user_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-keycloak_org_keycloak_v1alpha1_manifest_test.go: out/install-sentinel terratest/keycloak_org_v1alpha1/keycloak_org_keycloak_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_keycloak_org_keycloak_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/keycloak_org_v1alpha1/keycloak_org_keycloak_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-keyspaces_services_k8s_aws_keyspace_v1alpha1_manifest_test.go: out/install-sentinel terratest/keyspaces_services_k8s_aws_v1alpha1/keyspaces_services_k8s_aws_keyspace_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_keyspaces_services_k8s_aws_keyspace_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/keyspaces_services_k8s_aws_v1alpha1/keyspaces_services_k8s_aws_keyspace_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-keyspaces_services_k8s_aws_table_v1alpha1_manifest_test.go: out/install-sentinel terratest/keyspaces_services_k8s_aws_v1alpha1/keyspaces_services_k8s_aws_table_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_keyspaces_services_k8s_aws_table_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/keyspaces_services_k8s_aws_v1alpha1/keyspaces_services_k8s_aws_table_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kiali_io_kiali_v1alpha1_manifest_test.go: out/install-sentinel terratest/kiali_io_v1alpha1/kiali_io_kiali_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kiali_io_kiali_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kiali_io_v1alpha1/kiali_io_kiali_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kibana_k8s_elastic_co_kibana_v1_manifest_test.go: out/install-sentinel terratest/kibana_k8s_elastic_co_v1/kibana_k8s_elastic_co_kibana_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_kibana_k8s_elastic_co_kibana_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kibana_k8s_elastic_co_v1/kibana_k8s_elastic_co_kibana_v1_manifest_test.go
	touch $@
out/terratest-sentinel-kibana_k8s_elastic_co_kibana_v1beta1_manifest_test.go: out/install-sentinel terratest/kibana_k8s_elastic_co_v1beta1/kibana_k8s_elastic_co_kibana_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_kibana_k8s_elastic_co_kibana_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kibana_k8s_elastic_co_v1beta1/kibana_k8s_elastic_co_kibana_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-kinesis_services_k8s_aws_stream_v1alpha1_manifest_test.go: out/install-sentinel terratest/kinesis_services_k8s_aws_v1alpha1/kinesis_services_k8s_aws_stream_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kinesis_services_k8s_aws_stream_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kinesis_services_k8s_aws_v1alpha1/kinesis_services_k8s_aws_stream_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kmm_sigs_x_k8s_io_module_v1beta1_manifest_test.go: out/install-sentinel terratest/kmm_sigs_x_k8s_io_v1beta1/kmm_sigs_x_k8s_io_module_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_kmm_sigs_x_k8s_io_module_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kmm_sigs_x_k8s_io_v1beta1/kmm_sigs_x_k8s_io_module_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-kmm_sigs_x_k8s_io_node_modules_config_v1beta1_manifest_test.go: out/install-sentinel terratest/kmm_sigs_x_k8s_io_v1beta1/kmm_sigs_x_k8s_io_node_modules_config_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_kmm_sigs_x_k8s_io_node_modules_config_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kmm_sigs_x_k8s_io_v1beta1/kmm_sigs_x_k8s_io_node_modules_config_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-kmm_sigs_x_k8s_io_preflight_validation_v1beta1_manifest_test.go: out/install-sentinel terratest/kmm_sigs_x_k8s_io_v1beta1/kmm_sigs_x_k8s_io_preflight_validation_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_kmm_sigs_x_k8s_io_preflight_validation_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kmm_sigs_x_k8s_io_v1beta1/kmm_sigs_x_k8s_io_preflight_validation_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-kmm_sigs_x_k8s_io_preflight_validation_v1beta2_manifest_test.go: out/install-sentinel terratest/kmm_sigs_x_k8s_io_v1beta2/kmm_sigs_x_k8s_io_preflight_validation_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_kmm_sigs_x_k8s_io_preflight_validation_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kmm_sigs_x_k8s_io_v1beta2/kmm_sigs_x_k8s_io_preflight_validation_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-kms_services_k8s_aws_alias_v1alpha1_manifest_test.go: out/install-sentinel terratest/kms_services_k8s_aws_v1alpha1/kms_services_k8s_aws_alias_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kms_services_k8s_aws_alias_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kms_services_k8s_aws_v1alpha1/kms_services_k8s_aws_alias_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kms_services_k8s_aws_grant_v1alpha1_manifest_test.go: out/install-sentinel terratest/kms_services_k8s_aws_v1alpha1/kms_services_k8s_aws_grant_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kms_services_k8s_aws_grant_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kms_services_k8s_aws_v1alpha1/kms_services_k8s_aws_grant_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kms_services_k8s_aws_key_v1alpha1_manifest_test.go: out/install-sentinel terratest/kms_services_k8s_aws_v1alpha1/kms_services_k8s_aws_key_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kms_services_k8s_aws_key_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kms_services_k8s_aws_v1alpha1/kms_services_k8s_aws_key_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuadrant_io_dns_record_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuadrant_io_v1alpha1/kuadrant_io_dns_record_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuadrant_io_dns_record_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuadrant_io_v1alpha1/kuadrant_io_dns_record_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuadrant_io_managed_zone_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuadrant_io_v1alpha1/kuadrant_io_managed_zone_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuadrant_io_managed_zone_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuadrant_io_v1alpha1/kuadrant_io_managed_zone_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuadrant_io_kuadrant_v1beta1_manifest_test.go: out/install-sentinel terratest/kuadrant_io_v1beta1/kuadrant_io_kuadrant_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuadrant_io_kuadrant_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuadrant_io_v1beta1/kuadrant_io_kuadrant_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-kuadrant_io_auth_policy_v1beta2_manifest_test.go: out/install-sentinel terratest/kuadrant_io_v1beta2/kuadrant_io_auth_policy_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_kuadrant_io_auth_policy_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuadrant_io_v1beta2/kuadrant_io_auth_policy_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-kuadrant_io_rate_limit_policy_v1beta2_manifest_test.go: out/install-sentinel terratest/kuadrant_io_v1beta2/kuadrant_io_rate_limit_policy_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_kuadrant_io_rate_limit_policy_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuadrant_io_v1beta2/kuadrant_io_rate_limit_policy_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-kube_green_com_sleep_info_v1alpha1_manifest_test.go: out/install-sentinel terratest/kube_green_com_v1alpha1/kube_green_com_sleep_info_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kube_green_com_sleep_info_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kube_green_com_v1alpha1/kube_green_com_sleep_info_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kubean_io_cluster_operation_v1alpha1_manifest_test.go: out/install-sentinel terratest/kubean_io_v1alpha1/kubean_io_cluster_operation_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kubean_io_cluster_operation_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kubean_io_v1alpha1/kubean_io_cluster_operation_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kubean_io_cluster_v1alpha1_manifest_test.go: out/install-sentinel terratest/kubean_io_v1alpha1/kubean_io_cluster_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kubean_io_cluster_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kubean_io_v1alpha1/kubean_io_cluster_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kubean_io_local_artifact_set_v1alpha1_manifest_test.go: out/install-sentinel terratest/kubean_io_v1alpha1/kubean_io_local_artifact_set_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kubean_io_local_artifact_set_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kubean_io_v1alpha1/kubean_io_local_artifact_set_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kubean_io_manifest_v1alpha1_manifest_test.go: out/install-sentinel terratest/kubean_io_v1alpha1/kubean_io_manifest_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kubean_io_manifest_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kubean_io_v1alpha1/kubean_io_manifest_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kubecost_com_turndown_schedule_v1alpha1_manifest_test.go: out/install-sentinel terratest/kubecost_com_v1alpha1/kubecost_com_turndown_schedule_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kubecost_com_turndown_schedule_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kubecost_com_v1alpha1/kubecost_com_turndown_schedule_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kubevious_io_workload_profile_v1alpha1_manifest_test.go: out/install-sentinel terratest/kubevious_io_v1alpha1/kubevious_io_workload_profile_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kubevious_io_workload_profile_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kubevious_io_v1alpha1/kubevious_io_workload_profile_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kubevious_io_workload_v1alpha1_manifest_test.go: out/install-sentinel terratest/kubevious_io_v1alpha1/kubevious_io_workload_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kubevious_io_workload_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kubevious_io_v1alpha1/kubevious_io_workload_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kueue_x_k8s_io_admission_check_v1beta1_manifest_test.go: out/install-sentinel terratest/kueue_x_k8s_io_v1beta1/kueue_x_k8s_io_admission_check_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_kueue_x_k8s_io_admission_check_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kueue_x_k8s_io_v1beta1/kueue_x_k8s_io_admission_check_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-kueue_x_k8s_io_cluster_queue_v1beta1_manifest_test.go: out/install-sentinel terratest/kueue_x_k8s_io_v1beta1/kueue_x_k8s_io_cluster_queue_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_kueue_x_k8s_io_cluster_queue_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kueue_x_k8s_io_v1beta1/kueue_x_k8s_io_cluster_queue_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-kueue_x_k8s_io_local_queue_v1beta1_manifest_test.go: out/install-sentinel terratest/kueue_x_k8s_io_v1beta1/kueue_x_k8s_io_local_queue_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_kueue_x_k8s_io_local_queue_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kueue_x_k8s_io_v1beta1/kueue_x_k8s_io_local_queue_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-kueue_x_k8s_io_resource_flavor_v1beta1_manifest_test.go: out/install-sentinel terratest/kueue_x_k8s_io_v1beta1/kueue_x_k8s_io_resource_flavor_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_kueue_x_k8s_io_resource_flavor_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kueue_x_k8s_io_v1beta1/kueue_x_k8s_io_resource_flavor_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-kueue_x_k8s_io_workload_v1beta1_manifest_test.go: out/install-sentinel terratest/kueue_x_k8s_io_v1beta1/kueue_x_k8s_io_workload_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_kueue_x_k8s_io_workload_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kueue_x_k8s_io_v1beta1/kueue_x_k8s_io_workload_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_circuit_breaker_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_circuit_breaker_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_circuit_breaker_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_circuit_breaker_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_container_patch_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_container_patch_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_container_patch_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_container_patch_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_dataplane_insight_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_dataplane_insight_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_dataplane_insight_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_dataplane_insight_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_dataplane_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_dataplane_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_dataplane_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_dataplane_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_external_service_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_external_service_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_external_service_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_external_service_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_fault_injection_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_fault_injection_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_fault_injection_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_fault_injection_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_health_check_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_health_check_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_health_check_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_health_check_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_mesh_access_log_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_mesh_access_log_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_mesh_access_log_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_mesh_access_log_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_mesh_circuit_breaker_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_mesh_circuit_breaker_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_mesh_circuit_breaker_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_mesh_circuit_breaker_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_mesh_fault_injection_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_mesh_fault_injection_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_mesh_fault_injection_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_mesh_fault_injection_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_mesh_gateway_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_mesh_gateway_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_mesh_gateway_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_mesh_gateway_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_mesh_gateway_instance_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_mesh_gateway_instance_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_mesh_gateway_instance_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_mesh_gateway_instance_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_mesh_gateway_route_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_mesh_gateway_route_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_mesh_gateway_route_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_mesh_gateway_route_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_mesh_gateway_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_mesh_gateway_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_mesh_gateway_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_mesh_gateway_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_mesh_health_check_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_mesh_health_check_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_mesh_health_check_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_mesh_health_check_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_mesh_http_route_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_mesh_http_route_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_mesh_http_route_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_mesh_http_route_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_mesh_insight_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_mesh_insight_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_mesh_insight_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_mesh_insight_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_mesh_load_balancing_strategy_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_mesh_load_balancing_strategy_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_mesh_load_balancing_strategy_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_mesh_load_balancing_strategy_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_mesh_proxy_patch_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_mesh_proxy_patch_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_mesh_proxy_patch_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_mesh_proxy_patch_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_mesh_rate_limit_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_mesh_rate_limit_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_mesh_rate_limit_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_mesh_rate_limit_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_mesh_retry_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_mesh_retry_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_mesh_retry_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_mesh_retry_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_mesh_tcp_route_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_mesh_tcp_route_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_mesh_tcp_route_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_mesh_tcp_route_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_mesh_timeout_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_mesh_timeout_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_mesh_timeout_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_mesh_timeout_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_mesh_trace_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_mesh_trace_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_mesh_trace_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_mesh_trace_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_mesh_traffic_permission_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_mesh_traffic_permission_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_mesh_traffic_permission_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_mesh_traffic_permission_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_mesh_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_mesh_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_mesh_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_mesh_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_proxy_template_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_proxy_template_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_proxy_template_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_proxy_template_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_rate_limit_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_rate_limit_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_rate_limit_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_rate_limit_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_retry_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_retry_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_retry_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_retry_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_service_insight_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_service_insight_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_service_insight_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_service_insight_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_timeout_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_timeout_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_timeout_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_timeout_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_traffic_log_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_traffic_log_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_traffic_log_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_traffic_log_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_traffic_permission_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_traffic_permission_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_traffic_permission_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_traffic_permission_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_traffic_route_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_traffic_route_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_traffic_route_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_traffic_route_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_traffic_trace_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_traffic_trace_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_traffic_trace_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_traffic_trace_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_virtual_outbound_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_virtual_outbound_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_virtual_outbound_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_virtual_outbound_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_zone_egress_insight_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_zone_egress_insight_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_zone_egress_insight_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_zone_egress_insight_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_zone_egress_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_zone_egress_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_zone_egress_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_zone_egress_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_zone_ingress_insight_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_zone_ingress_insight_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_zone_ingress_insight_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_zone_ingress_insight_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_zone_ingress_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_zone_ingress_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_zone_ingress_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_zone_ingress_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_zone_insight_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_zone_insight_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_zone_insight_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_zone_insight_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kuma_io_zone_v1alpha1_manifest_test.go: out/install-sentinel terratest/kuma_io_v1alpha1/kuma_io_zone_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kuma_io_zone_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kuma_io_v1alpha1/kuma_io_zone_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kustomize_toolkit_fluxcd_io_kustomization_v1_manifest_test.go: out/install-sentinel terratest/kustomize_toolkit_fluxcd_io_v1/kustomize_toolkit_fluxcd_io_kustomization_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_kustomize_toolkit_fluxcd_io_kustomization_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kustomize_toolkit_fluxcd_io_v1/kustomize_toolkit_fluxcd_io_kustomization_v1_manifest_test.go
	touch $@
out/terratest-sentinel-kustomize_toolkit_fluxcd_io_kustomization_v1beta1_manifest_test.go: out/install-sentinel terratest/kustomize_toolkit_fluxcd_io_v1beta1/kustomize_toolkit_fluxcd_io_kustomization_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_kustomize_toolkit_fluxcd_io_kustomization_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kustomize_toolkit_fluxcd_io_v1beta1/kustomize_toolkit_fluxcd_io_kustomization_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-kustomize_toolkit_fluxcd_io_kustomization_v1beta2_manifest_test.go: out/install-sentinel terratest/kustomize_toolkit_fluxcd_io_v1beta2/kustomize_toolkit_fluxcd_io_kustomization_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_kustomize_toolkit_fluxcd_io_kustomization_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kustomize_toolkit_fluxcd_io_v1beta2/kustomize_toolkit_fluxcd_io_kustomization_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-kyverno_io_cluster_policy_v1_manifest_test.go: out/install-sentinel terratest/kyverno_io_v1/kyverno_io_cluster_policy_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_kyverno_io_cluster_policy_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kyverno_io_v1/kyverno_io_cluster_policy_v1_manifest_test.go
	touch $@
out/terratest-sentinel-kyverno_io_policy_v1_manifest_test.go: out/install-sentinel terratest/kyverno_io_v1/kyverno_io_policy_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_kyverno_io_policy_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kyverno_io_v1/kyverno_io_policy_v1_manifest_test.go
	touch $@
out/terratest-sentinel-kyverno_io_admission_report_v1alpha2_manifest_test.go: out/install-sentinel terratest/kyverno_io_v1alpha2/kyverno_io_admission_report_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_kyverno_io_admission_report_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kyverno_io_v1alpha2/kyverno_io_admission_report_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-kyverno_io_background_scan_report_v1alpha2_manifest_test.go: out/install-sentinel terratest/kyverno_io_v1alpha2/kyverno_io_background_scan_report_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_kyverno_io_background_scan_report_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kyverno_io_v1alpha2/kyverno_io_background_scan_report_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-kyverno_io_cluster_admission_report_v1alpha2_manifest_test.go: out/install-sentinel terratest/kyverno_io_v1alpha2/kyverno_io_cluster_admission_report_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_kyverno_io_cluster_admission_report_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kyverno_io_v1alpha2/kyverno_io_cluster_admission_report_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-kyverno_io_cluster_background_scan_report_v1alpha2_manifest_test.go: out/install-sentinel terratest/kyverno_io_v1alpha2/kyverno_io_cluster_background_scan_report_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_kyverno_io_cluster_background_scan_report_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kyverno_io_v1alpha2/kyverno_io_cluster_background_scan_report_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-kyverno_io_update_request_v1beta1_manifest_test.go: out/install-sentinel terratest/kyverno_io_v1beta1/kyverno_io_update_request_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_kyverno_io_update_request_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kyverno_io_v1beta1/kyverno_io_update_request_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-kyverno_io_admission_report_v2_manifest_test.go: out/install-sentinel terratest/kyverno_io_v2/kyverno_io_admission_report_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_kyverno_io_admission_report_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kyverno_io_v2/kyverno_io_admission_report_v2_manifest_test.go
	touch $@
out/terratest-sentinel-kyverno_io_background_scan_report_v2_manifest_test.go: out/install-sentinel terratest/kyverno_io_v2/kyverno_io_background_scan_report_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_kyverno_io_background_scan_report_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kyverno_io_v2/kyverno_io_background_scan_report_v2_manifest_test.go
	touch $@
out/terratest-sentinel-kyverno_io_cleanup_policy_v2_manifest_test.go: out/install-sentinel terratest/kyverno_io_v2/kyverno_io_cleanup_policy_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_kyverno_io_cleanup_policy_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kyverno_io_v2/kyverno_io_cleanup_policy_v2_manifest_test.go
	touch $@
out/terratest-sentinel-kyverno_io_cluster_admission_report_v2_manifest_test.go: out/install-sentinel terratest/kyverno_io_v2/kyverno_io_cluster_admission_report_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_kyverno_io_cluster_admission_report_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kyverno_io_v2/kyverno_io_cluster_admission_report_v2_manifest_test.go
	touch $@
out/terratest-sentinel-kyverno_io_cluster_background_scan_report_v2_manifest_test.go: out/install-sentinel terratest/kyverno_io_v2/kyverno_io_cluster_background_scan_report_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_kyverno_io_cluster_background_scan_report_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kyverno_io_v2/kyverno_io_cluster_background_scan_report_v2_manifest_test.go
	touch $@
out/terratest-sentinel-kyverno_io_cluster_cleanup_policy_v2_manifest_test.go: out/install-sentinel terratest/kyverno_io_v2/kyverno_io_cluster_cleanup_policy_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_kyverno_io_cluster_cleanup_policy_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kyverno_io_v2/kyverno_io_cluster_cleanup_policy_v2_manifest_test.go
	touch $@
out/terratest-sentinel-kyverno_io_policy_exception_v2_manifest_test.go: out/install-sentinel terratest/kyverno_io_v2/kyverno_io_policy_exception_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_kyverno_io_policy_exception_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kyverno_io_v2/kyverno_io_policy_exception_v2_manifest_test.go
	touch $@
out/terratest-sentinel-kyverno_io_update_request_v2_manifest_test.go: out/install-sentinel terratest/kyverno_io_v2/kyverno_io_update_request_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_kyverno_io_update_request_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kyverno_io_v2/kyverno_io_update_request_v2_manifest_test.go
	touch $@
out/terratest-sentinel-kyverno_io_cleanup_policy_v2alpha1_manifest_test.go: out/install-sentinel terratest/kyverno_io_v2alpha1/kyverno_io_cleanup_policy_v2alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kyverno_io_cleanup_policy_v2alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kyverno_io_v2alpha1/kyverno_io_cleanup_policy_v2alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kyverno_io_cluster_cleanup_policy_v2alpha1_manifest_test.go: out/install-sentinel terratest/kyverno_io_v2alpha1/kyverno_io_cluster_cleanup_policy_v2alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kyverno_io_cluster_cleanup_policy_v2alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kyverno_io_v2alpha1/kyverno_io_cluster_cleanup_policy_v2alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kyverno_io_global_context_entry_v2alpha1_manifest_test.go: out/install-sentinel terratest/kyverno_io_v2alpha1/kyverno_io_global_context_entry_v2alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kyverno_io_global_context_entry_v2alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kyverno_io_v2alpha1/kyverno_io_global_context_entry_v2alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kyverno_io_policy_exception_v2alpha1_manifest_test.go: out/install-sentinel terratest/kyverno_io_v2alpha1/kyverno_io_policy_exception_v2alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_kyverno_io_policy_exception_v2alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kyverno_io_v2alpha1/kyverno_io_policy_exception_v2alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-kyverno_io_cleanup_policy_v2beta1_manifest_test.go: out/install-sentinel terratest/kyverno_io_v2beta1/kyverno_io_cleanup_policy_v2beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_kyverno_io_cleanup_policy_v2beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kyverno_io_v2beta1/kyverno_io_cleanup_policy_v2beta1_manifest_test.go
	touch $@
out/terratest-sentinel-kyverno_io_cluster_cleanup_policy_v2beta1_manifest_test.go: out/install-sentinel terratest/kyverno_io_v2beta1/kyverno_io_cluster_cleanup_policy_v2beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_kyverno_io_cluster_cleanup_policy_v2beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kyverno_io_v2beta1/kyverno_io_cluster_cleanup_policy_v2beta1_manifest_test.go
	touch $@
out/terratest-sentinel-kyverno_io_cluster_policy_v2beta1_manifest_test.go: out/install-sentinel terratest/kyverno_io_v2beta1/kyverno_io_cluster_policy_v2beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_kyverno_io_cluster_policy_v2beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kyverno_io_v2beta1/kyverno_io_cluster_policy_v2beta1_manifest_test.go
	touch $@
out/terratest-sentinel-kyverno_io_policy_exception_v2beta1_manifest_test.go: out/install-sentinel terratest/kyverno_io_v2beta1/kyverno_io_policy_exception_v2beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_kyverno_io_policy_exception_v2beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kyverno_io_v2beta1/kyverno_io_policy_exception_v2beta1_manifest_test.go
	touch $@
out/terratest-sentinel-kyverno_io_policy_v2beta1_manifest_test.go: out/install-sentinel terratest/kyverno_io_v2beta1/kyverno_io_policy_v2beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_kyverno_io_policy_v2beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/kyverno_io_v2beta1/kyverno_io_policy_v2beta1_manifest_test.go
	touch $@
out/terratest-sentinel-lambda_services_k8s_aws_alias_v1alpha1_manifest_test.go: out/install-sentinel terratest/lambda_services_k8s_aws_v1alpha1/lambda_services_k8s_aws_alias_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_lambda_services_k8s_aws_alias_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/lambda_services_k8s_aws_v1alpha1/lambda_services_k8s_aws_alias_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-lambda_services_k8s_aws_code_signing_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/lambda_services_k8s_aws_v1alpha1/lambda_services_k8s_aws_code_signing_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_lambda_services_k8s_aws_code_signing_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/lambda_services_k8s_aws_v1alpha1/lambda_services_k8s_aws_code_signing_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-lambda_services_k8s_aws_event_source_mapping_v1alpha1_manifest_test.go: out/install-sentinel terratest/lambda_services_k8s_aws_v1alpha1/lambda_services_k8s_aws_event_source_mapping_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_lambda_services_k8s_aws_event_source_mapping_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/lambda_services_k8s_aws_v1alpha1/lambda_services_k8s_aws_event_source_mapping_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-lambda_services_k8s_aws_function_url_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/lambda_services_k8s_aws_v1alpha1/lambda_services_k8s_aws_function_url_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_lambda_services_k8s_aws_function_url_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/lambda_services_k8s_aws_v1alpha1/lambda_services_k8s_aws_function_url_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-lambda_services_k8s_aws_function_v1alpha1_manifest_test.go: out/install-sentinel terratest/lambda_services_k8s_aws_v1alpha1/lambda_services_k8s_aws_function_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_lambda_services_k8s_aws_function_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/lambda_services_k8s_aws_v1alpha1/lambda_services_k8s_aws_function_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-lambda_services_k8s_aws_layer_version_v1alpha1_manifest_test.go: out/install-sentinel terratest/lambda_services_k8s_aws_v1alpha1/lambda_services_k8s_aws_layer_version_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_lambda_services_k8s_aws_layer_version_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/lambda_services_k8s_aws_v1alpha1/lambda_services_k8s_aws_layer_version_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-lambda_services_k8s_aws_version_v1alpha1_manifest_test.go: out/install-sentinel terratest/lambda_services_k8s_aws_v1alpha1/lambda_services_k8s_aws_version_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_lambda_services_k8s_aws_version_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/lambda_services_k8s_aws_v1alpha1/lambda_services_k8s_aws_version_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-lb_lbconfig_carlosedp_com_external_load_balancer_v1_manifest_test.go: out/install-sentinel terratest/lb_lbconfig_carlosedp_com_v1/lb_lbconfig_carlosedp_com_external_load_balancer_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_lb_lbconfig_carlosedp_com_external_load_balancer_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/lb_lbconfig_carlosedp_com_v1/lb_lbconfig_carlosedp_com_external_load_balancer_v1_manifest_test.go
	touch $@
out/terratest-sentinel-leaksignal_com_cluster_leaksignal_istio_v1_manifest_test.go: out/install-sentinel terratest/leaksignal_com_v1/leaksignal_com_cluster_leaksignal_istio_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_leaksignal_com_cluster_leaksignal_istio_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/leaksignal_com_v1/leaksignal_com_cluster_leaksignal_istio_v1_manifest_test.go
	touch $@
out/terratest-sentinel-leaksignal_com_leaksignal_istio_v1_manifest_test.go: out/install-sentinel terratest/leaksignal_com_v1/leaksignal_com_leaksignal_istio_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_leaksignal_com_leaksignal_istio_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/leaksignal_com_v1/leaksignal_com_leaksignal_istio_v1_manifest_test.go
	touch $@
out/terratest-sentinel-lerentis_uploadfilter24_eu_bitwarden_secret_v1beta4_manifest_test.go: out/install-sentinel terratest/lerentis_uploadfilter24_eu_v1beta4/lerentis_uploadfilter24_eu_bitwarden_secret_v1beta4_manifest_test.go $(shell find ./examples/data-sources/k8s_lerentis_uploadfilter24_eu_bitwarden_secret_v1beta4_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/lerentis_uploadfilter24_eu_v1beta4/lerentis_uploadfilter24_eu_bitwarden_secret_v1beta4_manifest_test.go
	touch $@
out/terratest-sentinel-lerentis_uploadfilter24_eu_bitwarden_template_v1beta4_manifest_test.go: out/install-sentinel terratest/lerentis_uploadfilter24_eu_v1beta4/lerentis_uploadfilter24_eu_bitwarden_template_v1beta4_manifest_test.go $(shell find ./examples/data-sources/k8s_lerentis_uploadfilter24_eu_bitwarden_template_v1beta4_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/lerentis_uploadfilter24_eu_v1beta4/lerentis_uploadfilter24_eu_bitwarden_template_v1beta4_manifest_test.go
	touch $@
out/terratest-sentinel-lerentis_uploadfilter24_eu_registry_credential_v1beta4_manifest_test.go: out/install-sentinel terratest/lerentis_uploadfilter24_eu_v1beta4/lerentis_uploadfilter24_eu_registry_credential_v1beta4_manifest_test.go $(shell find ./examples/data-sources/k8s_lerentis_uploadfilter24_eu_registry_credential_v1beta4_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/lerentis_uploadfilter24_eu_v1beta4/lerentis_uploadfilter24_eu_registry_credential_v1beta4_manifest_test.go
	touch $@
out/terratest-sentinel-lerentis_uploadfilter24_eu_bitwarden_secret_v1beta5_manifest_test.go: out/install-sentinel terratest/lerentis_uploadfilter24_eu_v1beta5/lerentis_uploadfilter24_eu_bitwarden_secret_v1beta5_manifest_test.go $(shell find ./examples/data-sources/k8s_lerentis_uploadfilter24_eu_bitwarden_secret_v1beta5_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/lerentis_uploadfilter24_eu_v1beta5/lerentis_uploadfilter24_eu_bitwarden_secret_v1beta5_manifest_test.go
	touch $@
out/terratest-sentinel-lerentis_uploadfilter24_eu_bitwarden_template_v1beta5_manifest_test.go: out/install-sentinel terratest/lerentis_uploadfilter24_eu_v1beta5/lerentis_uploadfilter24_eu_bitwarden_template_v1beta5_manifest_test.go $(shell find ./examples/data-sources/k8s_lerentis_uploadfilter24_eu_bitwarden_template_v1beta5_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/lerentis_uploadfilter24_eu_v1beta5/lerentis_uploadfilter24_eu_bitwarden_template_v1beta5_manifest_test.go
	touch $@
out/terratest-sentinel-lerentis_uploadfilter24_eu_registry_credential_v1beta5_manifest_test.go: out/install-sentinel terratest/lerentis_uploadfilter24_eu_v1beta5/lerentis_uploadfilter24_eu_registry_credential_v1beta5_manifest_test.go $(shell find ./examples/data-sources/k8s_lerentis_uploadfilter24_eu_registry_credential_v1beta5_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/lerentis_uploadfilter24_eu_v1beta5/lerentis_uploadfilter24_eu_registry_credential_v1beta5_manifest_test.go
	touch $@
out/terratest-sentinel-lerentis_uploadfilter24_eu_bitwarden_secret_v1beta6_manifest_test.go: out/install-sentinel terratest/lerentis_uploadfilter24_eu_v1beta6/lerentis_uploadfilter24_eu_bitwarden_secret_v1beta6_manifest_test.go $(shell find ./examples/data-sources/k8s_lerentis_uploadfilter24_eu_bitwarden_secret_v1beta6_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/lerentis_uploadfilter24_eu_v1beta6/lerentis_uploadfilter24_eu_bitwarden_secret_v1beta6_manifest_test.go
	touch $@
out/terratest-sentinel-lerentis_uploadfilter24_eu_bitwarden_template_v1beta6_manifest_test.go: out/install-sentinel terratest/lerentis_uploadfilter24_eu_v1beta6/lerentis_uploadfilter24_eu_bitwarden_template_v1beta6_manifest_test.go $(shell find ./examples/data-sources/k8s_lerentis_uploadfilter24_eu_bitwarden_template_v1beta6_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/lerentis_uploadfilter24_eu_v1beta6/lerentis_uploadfilter24_eu_bitwarden_template_v1beta6_manifest_test.go
	touch $@
out/terratest-sentinel-lerentis_uploadfilter24_eu_registry_credential_v1beta6_manifest_test.go: out/install-sentinel terratest/lerentis_uploadfilter24_eu_v1beta6/lerentis_uploadfilter24_eu_registry_credential_v1beta6_manifest_test.go $(shell find ./examples/data-sources/k8s_lerentis_uploadfilter24_eu_registry_credential_v1beta6_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/lerentis_uploadfilter24_eu_v1beta6/lerentis_uploadfilter24_eu_registry_credential_v1beta6_manifest_test.go
	touch $@
out/terratest-sentinel-lerentis_uploadfilter24_eu_bitwarden_secret_v1beta7_manifest_test.go: out/install-sentinel terratest/lerentis_uploadfilter24_eu_v1beta7/lerentis_uploadfilter24_eu_bitwarden_secret_v1beta7_manifest_test.go $(shell find ./examples/data-sources/k8s_lerentis_uploadfilter24_eu_bitwarden_secret_v1beta7_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/lerentis_uploadfilter24_eu_v1beta7/lerentis_uploadfilter24_eu_bitwarden_secret_v1beta7_manifest_test.go
	touch $@
out/terratest-sentinel-lerentis_uploadfilter24_eu_bitwarden_template_v1beta7_manifest_test.go: out/install-sentinel terratest/lerentis_uploadfilter24_eu_v1beta7/lerentis_uploadfilter24_eu_bitwarden_template_v1beta7_manifest_test.go $(shell find ./examples/data-sources/k8s_lerentis_uploadfilter24_eu_bitwarden_template_v1beta7_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/lerentis_uploadfilter24_eu_v1beta7/lerentis_uploadfilter24_eu_bitwarden_template_v1beta7_manifest_test.go
	touch $@
out/terratest-sentinel-lerentis_uploadfilter24_eu_registry_credential_v1beta7_manifest_test.go: out/install-sentinel terratest/lerentis_uploadfilter24_eu_v1beta7/lerentis_uploadfilter24_eu_registry_credential_v1beta7_manifest_test.go $(shell find ./examples/data-sources/k8s_lerentis_uploadfilter24_eu_registry_credential_v1beta7_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/lerentis_uploadfilter24_eu_v1beta7/lerentis_uploadfilter24_eu_registry_credential_v1beta7_manifest_test.go
	touch $@
out/terratest-sentinel-limitador_kuadrant_io_limitador_v1alpha1_manifest_test.go: out/install-sentinel terratest/limitador_kuadrant_io_v1alpha1/limitador_kuadrant_io_limitador_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_limitador_kuadrant_io_limitador_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/limitador_kuadrant_io_v1alpha1/limitador_kuadrant_io_limitador_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-litmuschaos_io_chaos_engine_v1alpha1_manifest_test.go: out/install-sentinel terratest/litmuschaos_io_v1alpha1/litmuschaos_io_chaos_engine_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_litmuschaos_io_chaos_engine_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/litmuschaos_io_v1alpha1/litmuschaos_io_chaos_engine_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-litmuschaos_io_chaos_experiment_v1alpha1_manifest_test.go: out/install-sentinel terratest/litmuschaos_io_v1alpha1/litmuschaos_io_chaos_experiment_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_litmuschaos_io_chaos_experiment_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/litmuschaos_io_v1alpha1/litmuschaos_io_chaos_experiment_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-litmuschaos_io_chaos_result_v1alpha1_manifest_test.go: out/install-sentinel terratest/litmuschaos_io_v1alpha1/litmuschaos_io_chaos_result_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_litmuschaos_io_chaos_result_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/litmuschaos_io_v1alpha1/litmuschaos_io_chaos_result_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-logging_banzaicloud_io_cluster_flow_v1alpha1_manifest_test.go: out/install-sentinel terratest/logging_banzaicloud_io_v1alpha1/logging_banzaicloud_io_cluster_flow_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_logging_banzaicloud_io_cluster_flow_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/logging_banzaicloud_io_v1alpha1/logging_banzaicloud_io_cluster_flow_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-logging_banzaicloud_io_cluster_output_v1alpha1_manifest_test.go: out/install-sentinel terratest/logging_banzaicloud_io_v1alpha1/logging_banzaicloud_io_cluster_output_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_logging_banzaicloud_io_cluster_output_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/logging_banzaicloud_io_v1alpha1/logging_banzaicloud_io_cluster_output_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-logging_banzaicloud_io_flow_v1alpha1_manifest_test.go: out/install-sentinel terratest/logging_banzaicloud_io_v1alpha1/logging_banzaicloud_io_flow_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_logging_banzaicloud_io_flow_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/logging_banzaicloud_io_v1alpha1/logging_banzaicloud_io_flow_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-logging_banzaicloud_io_logging_v1alpha1_manifest_test.go: out/install-sentinel terratest/logging_banzaicloud_io_v1alpha1/logging_banzaicloud_io_logging_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_logging_banzaicloud_io_logging_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/logging_banzaicloud_io_v1alpha1/logging_banzaicloud_io_logging_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-logging_banzaicloud_io_output_v1alpha1_manifest_test.go: out/install-sentinel terratest/logging_banzaicloud_io_v1alpha1/logging_banzaicloud_io_output_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_logging_banzaicloud_io_output_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/logging_banzaicloud_io_v1alpha1/logging_banzaicloud_io_output_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-logging_banzaicloud_io_cluster_flow_v1beta1_manifest_test.go: out/install-sentinel terratest/logging_banzaicloud_io_v1beta1/logging_banzaicloud_io_cluster_flow_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_logging_banzaicloud_io_cluster_flow_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/logging_banzaicloud_io_v1beta1/logging_banzaicloud_io_cluster_flow_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-logging_banzaicloud_io_cluster_output_v1beta1_manifest_test.go: out/install-sentinel terratest/logging_banzaicloud_io_v1beta1/logging_banzaicloud_io_cluster_output_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_logging_banzaicloud_io_cluster_output_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/logging_banzaicloud_io_v1beta1/logging_banzaicloud_io_cluster_output_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-logging_banzaicloud_io_flow_v1beta1_manifest_test.go: out/install-sentinel terratest/logging_banzaicloud_io_v1beta1/logging_banzaicloud_io_flow_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_logging_banzaicloud_io_flow_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/logging_banzaicloud_io_v1beta1/logging_banzaicloud_io_flow_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-logging_banzaicloud_io_fluentbit_agent_v1beta1_manifest_test.go: out/install-sentinel terratest/logging_banzaicloud_io_v1beta1/logging_banzaicloud_io_fluentbit_agent_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_logging_banzaicloud_io_fluentbit_agent_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/logging_banzaicloud_io_v1beta1/logging_banzaicloud_io_fluentbit_agent_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-logging_banzaicloud_io_logging_v1beta1_manifest_test.go: out/install-sentinel terratest/logging_banzaicloud_io_v1beta1/logging_banzaicloud_io_logging_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_logging_banzaicloud_io_logging_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/logging_banzaicloud_io_v1beta1/logging_banzaicloud_io_logging_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-logging_banzaicloud_io_node_agent_v1beta1_manifest_test.go: out/install-sentinel terratest/logging_banzaicloud_io_v1beta1/logging_banzaicloud_io_node_agent_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_logging_banzaicloud_io_node_agent_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/logging_banzaicloud_io_v1beta1/logging_banzaicloud_io_node_agent_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-logging_banzaicloud_io_output_v1beta1_manifest_test.go: out/install-sentinel terratest/logging_banzaicloud_io_v1beta1/logging_banzaicloud_io_output_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_logging_banzaicloud_io_output_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/logging_banzaicloud_io_v1beta1/logging_banzaicloud_io_output_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-logging_banzaicloud_io_syslog_ng_cluster_flow_v1beta1_manifest_test.go: out/install-sentinel terratest/logging_banzaicloud_io_v1beta1/logging_banzaicloud_io_syslog_ng_cluster_flow_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_logging_banzaicloud_io_syslog_ng_cluster_flow_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/logging_banzaicloud_io_v1beta1/logging_banzaicloud_io_syslog_ng_cluster_flow_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-logging_banzaicloud_io_syslog_ng_cluster_output_v1beta1_manifest_test.go: out/install-sentinel terratest/logging_banzaicloud_io_v1beta1/logging_banzaicloud_io_syslog_ng_cluster_output_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_logging_banzaicloud_io_syslog_ng_cluster_output_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/logging_banzaicloud_io_v1beta1/logging_banzaicloud_io_syslog_ng_cluster_output_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-logging_banzaicloud_io_syslog_ng_flow_v1beta1_manifest_test.go: out/install-sentinel terratest/logging_banzaicloud_io_v1beta1/logging_banzaicloud_io_syslog_ng_flow_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_logging_banzaicloud_io_syslog_ng_flow_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/logging_banzaicloud_io_v1beta1/logging_banzaicloud_io_syslog_ng_flow_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-logging_banzaicloud_io_syslog_ng_output_v1beta1_manifest_test.go: out/install-sentinel terratest/logging_banzaicloud_io_v1beta1/logging_banzaicloud_io_syslog_ng_output_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_logging_banzaicloud_io_syslog_ng_output_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/logging_banzaicloud_io_v1beta1/logging_banzaicloud_io_syslog_ng_output_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-logging_extensions_banzaicloud_io_event_tailer_v1alpha1_manifest_test.go: out/install-sentinel terratest/logging_extensions_banzaicloud_io_v1alpha1/logging_extensions_banzaicloud_io_event_tailer_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_logging_extensions_banzaicloud_io_event_tailer_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/logging_extensions_banzaicloud_io_v1alpha1/logging_extensions_banzaicloud_io_event_tailer_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-logging_extensions_banzaicloud_io_host_tailer_v1alpha1_manifest_test.go: out/install-sentinel terratest/logging_extensions_banzaicloud_io_v1alpha1/logging_extensions_banzaicloud_io_host_tailer_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_logging_extensions_banzaicloud_io_host_tailer_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/logging_extensions_banzaicloud_io_v1alpha1/logging_extensions_banzaicloud_io_host_tailer_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-loki_grafana_com_alerting_rule_v1_manifest_test.go: out/install-sentinel terratest/loki_grafana_com_v1/loki_grafana_com_alerting_rule_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_loki_grafana_com_alerting_rule_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/loki_grafana_com_v1/loki_grafana_com_alerting_rule_v1_manifest_test.go
	touch $@
out/terratest-sentinel-loki_grafana_com_loki_stack_v1_manifest_test.go: out/install-sentinel terratest/loki_grafana_com_v1/loki_grafana_com_loki_stack_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_loki_grafana_com_loki_stack_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/loki_grafana_com_v1/loki_grafana_com_loki_stack_v1_manifest_test.go
	touch $@
out/terratest-sentinel-loki_grafana_com_recording_rule_v1_manifest_test.go: out/install-sentinel terratest/loki_grafana_com_v1/loki_grafana_com_recording_rule_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_loki_grafana_com_recording_rule_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/loki_grafana_com_v1/loki_grafana_com_recording_rule_v1_manifest_test.go
	touch $@
out/terratest-sentinel-loki_grafana_com_ruler_config_v1_manifest_test.go: out/install-sentinel terratest/loki_grafana_com_v1/loki_grafana_com_ruler_config_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_loki_grafana_com_ruler_config_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/loki_grafana_com_v1/loki_grafana_com_ruler_config_v1_manifest_test.go
	touch $@
out/terratest-sentinel-loki_grafana_com_alerting_rule_v1beta1_manifest_test.go: out/install-sentinel terratest/loki_grafana_com_v1beta1/loki_grafana_com_alerting_rule_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_loki_grafana_com_alerting_rule_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/loki_grafana_com_v1beta1/loki_grafana_com_alerting_rule_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-loki_grafana_com_loki_stack_v1beta1_manifest_test.go: out/install-sentinel terratest/loki_grafana_com_v1beta1/loki_grafana_com_loki_stack_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_loki_grafana_com_loki_stack_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/loki_grafana_com_v1beta1/loki_grafana_com_loki_stack_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-loki_grafana_com_recording_rule_v1beta1_manifest_test.go: out/install-sentinel terratest/loki_grafana_com_v1beta1/loki_grafana_com_recording_rule_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_loki_grafana_com_recording_rule_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/loki_grafana_com_v1beta1/loki_grafana_com_recording_rule_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-loki_grafana_com_ruler_config_v1beta1_manifest_test.go: out/install-sentinel terratest/loki_grafana_com_v1beta1/loki_grafana_com_ruler_config_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_loki_grafana_com_ruler_config_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/loki_grafana_com_v1beta1/loki_grafana_com_ruler_config_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_backing_image_data_source_v1beta1_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta1/longhorn_io_backing_image_data_source_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_backing_image_data_source_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta1/longhorn_io_backing_image_data_source_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_backing_image_manager_v1beta1_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta1/longhorn_io_backing_image_manager_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_backing_image_manager_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta1/longhorn_io_backing_image_manager_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_backing_image_v1beta1_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta1/longhorn_io_backing_image_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_backing_image_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta1/longhorn_io_backing_image_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_backup_target_v1beta1_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta1/longhorn_io_backup_target_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_backup_target_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta1/longhorn_io_backup_target_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_backup_v1beta1_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta1/longhorn_io_backup_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_backup_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta1/longhorn_io_backup_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_backup_volume_v1beta1_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta1/longhorn_io_backup_volume_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_backup_volume_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta1/longhorn_io_backup_volume_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_engine_image_v1beta1_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta1/longhorn_io_engine_image_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_engine_image_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta1/longhorn_io_engine_image_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_engine_v1beta1_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta1/longhorn_io_engine_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_engine_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta1/longhorn_io_engine_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_instance_manager_v1beta1_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta1/longhorn_io_instance_manager_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_instance_manager_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta1/longhorn_io_instance_manager_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_node_v1beta1_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta1/longhorn_io_node_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_node_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta1/longhorn_io_node_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_recurring_job_v1beta1_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta1/longhorn_io_recurring_job_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_recurring_job_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta1/longhorn_io_recurring_job_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_replica_v1beta1_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta1/longhorn_io_replica_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_replica_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta1/longhorn_io_replica_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_setting_v1beta1_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta1/longhorn_io_setting_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_setting_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta1/longhorn_io_setting_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_share_manager_v1beta1_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta1/longhorn_io_share_manager_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_share_manager_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta1/longhorn_io_share_manager_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_volume_v1beta1_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta1/longhorn_io_volume_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_volume_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta1/longhorn_io_volume_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_backing_image_data_source_v1beta2_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta2/longhorn_io_backing_image_data_source_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_backing_image_data_source_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta2/longhorn_io_backing_image_data_source_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_backing_image_manager_v1beta2_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta2/longhorn_io_backing_image_manager_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_backing_image_manager_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta2/longhorn_io_backing_image_manager_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_backing_image_v1beta2_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta2/longhorn_io_backing_image_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_backing_image_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta2/longhorn_io_backing_image_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_backup_backing_image_v1beta2_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta2/longhorn_io_backup_backing_image_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_backup_backing_image_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta2/longhorn_io_backup_backing_image_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_backup_target_v1beta2_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta2/longhorn_io_backup_target_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_backup_target_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta2/longhorn_io_backup_target_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_backup_v1beta2_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta2/longhorn_io_backup_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_backup_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta2/longhorn_io_backup_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_backup_volume_v1beta2_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta2/longhorn_io_backup_volume_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_backup_volume_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta2/longhorn_io_backup_volume_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_engine_image_v1beta2_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta2/longhorn_io_engine_image_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_engine_image_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta2/longhorn_io_engine_image_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_engine_v1beta2_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta2/longhorn_io_engine_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_engine_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta2/longhorn_io_engine_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_instance_manager_v1beta2_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta2/longhorn_io_instance_manager_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_instance_manager_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta2/longhorn_io_instance_manager_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_node_v1beta2_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta2/longhorn_io_node_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_node_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta2/longhorn_io_node_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_orphan_v1beta2_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta2/longhorn_io_orphan_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_orphan_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta2/longhorn_io_orphan_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_recurring_job_v1beta2_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta2/longhorn_io_recurring_job_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_recurring_job_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta2/longhorn_io_recurring_job_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_replica_v1beta2_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta2/longhorn_io_replica_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_replica_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta2/longhorn_io_replica_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_setting_v1beta2_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta2/longhorn_io_setting_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_setting_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta2/longhorn_io_setting_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_share_manager_v1beta2_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta2/longhorn_io_share_manager_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_share_manager_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta2/longhorn_io_share_manager_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_snapshot_v1beta2_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta2/longhorn_io_snapshot_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_snapshot_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta2/longhorn_io_snapshot_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_support_bundle_v1beta2_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta2/longhorn_io_support_bundle_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_support_bundle_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta2/longhorn_io_support_bundle_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_system_backup_v1beta2_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta2/longhorn_io_system_backup_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_system_backup_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta2/longhorn_io_system_backup_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_system_restore_v1beta2_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta2/longhorn_io_system_restore_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_system_restore_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta2/longhorn_io_system_restore_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_volume_attachment_v1beta2_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta2/longhorn_io_volume_attachment_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_volume_attachment_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta2/longhorn_io_volume_attachment_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-longhorn_io_volume_v1beta2_manifest_test.go: out/install-sentinel terratest/longhorn_io_v1beta2/longhorn_io_volume_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_longhorn_io_volume_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/longhorn_io_v1beta2/longhorn_io_volume_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-m4e_krestomat_io_moodle_v1alpha1_manifest_test.go: out/install-sentinel terratest/m4e_krestomat_io_v1alpha1/m4e_krestomat_io_moodle_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_m4e_krestomat_io_moodle_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/m4e_krestomat_io_v1alpha1/m4e_krestomat_io_moodle_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-m4e_krestomat_io_nginx_v1alpha1_manifest_test.go: out/install-sentinel terratest/m4e_krestomat_io_v1alpha1/m4e_krestomat_io_nginx_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_m4e_krestomat_io_nginx_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/m4e_krestomat_io_v1alpha1/m4e_krestomat_io_nginx_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-m4e_krestomat_io_phpfpm_v1alpha1_manifest_test.go: out/install-sentinel terratest/m4e_krestomat_io_v1alpha1/m4e_krestomat_io_phpfpm_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_m4e_krestomat_io_phpfpm_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/m4e_krestomat_io_v1alpha1/m4e_krestomat_io_phpfpm_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-m4e_krestomat_io_routine_v1alpha1_manifest_test.go: out/install-sentinel terratest/m4e_krestomat_io_v1alpha1/m4e_krestomat_io_routine_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_m4e_krestomat_io_routine_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/m4e_krestomat_io_v1alpha1/m4e_krestomat_io_routine_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-machine_deletion_remediation_medik8s_io_machine_deletion_remediation_template_v1alpha1_manifest_test.go: out/install-sentinel terratest/machine_deletion_remediation_medik8s_io_v1alpha1/machine_deletion_remediation_medik8s_io_machine_deletion_remediation_template_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_machine_deletion_remediation_medik8s_io_machine_deletion_remediation_template_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/machine_deletion_remediation_medik8s_io_v1alpha1/machine_deletion_remediation_medik8s_io_machine_deletion_remediation_template_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-machine_deletion_remediation_medik8s_io_machine_deletion_remediation_v1alpha1_manifest_test.go: out/install-sentinel terratest/machine_deletion_remediation_medik8s_io_v1alpha1/machine_deletion_remediation_medik8s_io_machine_deletion_remediation_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_machine_deletion_remediation_medik8s_io_machine_deletion_remediation_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/machine_deletion_remediation_medik8s_io_v1alpha1/machine_deletion_remediation_medik8s_io_machine_deletion_remediation_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-maps_k8s_elastic_co_elastic_maps_server_v1alpha1_manifest_test.go: out/install-sentinel terratest/maps_k8s_elastic_co_v1alpha1/maps_k8s_elastic_co_elastic_maps_server_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_maps_k8s_elastic_co_elastic_maps_server_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/maps_k8s_elastic_co_v1alpha1/maps_k8s_elastic_co_elastic_maps_server_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-mariadb_mmontes_io_backup_v1alpha1_manifest_test.go: out/install-sentinel terratest/mariadb_mmontes_io_v1alpha1/mariadb_mmontes_io_backup_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_mariadb_mmontes_io_backup_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/mariadb_mmontes_io_v1alpha1/mariadb_mmontes_io_backup_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-mariadb_mmontes_io_connection_v1alpha1_manifest_test.go: out/install-sentinel terratest/mariadb_mmontes_io_v1alpha1/mariadb_mmontes_io_connection_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_mariadb_mmontes_io_connection_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/mariadb_mmontes_io_v1alpha1/mariadb_mmontes_io_connection_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-mariadb_mmontes_io_database_v1alpha1_manifest_test.go: out/install-sentinel terratest/mariadb_mmontes_io_v1alpha1/mariadb_mmontes_io_database_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_mariadb_mmontes_io_database_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/mariadb_mmontes_io_v1alpha1/mariadb_mmontes_io_database_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-mariadb_mmontes_io_grant_v1alpha1_manifest_test.go: out/install-sentinel terratest/mariadb_mmontes_io_v1alpha1/mariadb_mmontes_io_grant_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_mariadb_mmontes_io_grant_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/mariadb_mmontes_io_v1alpha1/mariadb_mmontes_io_grant_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-mariadb_mmontes_io_maria_db_v1alpha1_manifest_test.go: out/install-sentinel terratest/mariadb_mmontes_io_v1alpha1/mariadb_mmontes_io_maria_db_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_mariadb_mmontes_io_maria_db_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/mariadb_mmontes_io_v1alpha1/mariadb_mmontes_io_maria_db_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-mariadb_mmontes_io_restore_v1alpha1_manifest_test.go: out/install-sentinel terratest/mariadb_mmontes_io_v1alpha1/mariadb_mmontes_io_restore_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_mariadb_mmontes_io_restore_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/mariadb_mmontes_io_v1alpha1/mariadb_mmontes_io_restore_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-mariadb_mmontes_io_sql_job_v1alpha1_manifest_test.go: out/install-sentinel terratest/mariadb_mmontes_io_v1alpha1/mariadb_mmontes_io_sql_job_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_mariadb_mmontes_io_sql_job_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/mariadb_mmontes_io_v1alpha1/mariadb_mmontes_io_sql_job_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-mariadb_mmontes_io_user_v1alpha1_manifest_test.go: out/install-sentinel terratest/mariadb_mmontes_io_v1alpha1/mariadb_mmontes_io_user_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_mariadb_mmontes_io_user_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/mariadb_mmontes_io_v1alpha1/mariadb_mmontes_io_user_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-marin3r_3scale_net_envoy_config_revision_v1alpha1_manifest_test.go: out/install-sentinel terratest/marin3r_3scale_net_v1alpha1/marin3r_3scale_net_envoy_config_revision_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_marin3r_3scale_net_envoy_config_revision_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/marin3r_3scale_net_v1alpha1/marin3r_3scale_net_envoy_config_revision_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-marin3r_3scale_net_envoy_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/marin3r_3scale_net_v1alpha1/marin3r_3scale_net_envoy_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_marin3r_3scale_net_envoy_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/marin3r_3scale_net_v1alpha1/marin3r_3scale_net_envoy_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-mattermost_com_cluster_installation_v1alpha1_manifest_test.go: out/install-sentinel terratest/mattermost_com_v1alpha1/mattermost_com_cluster_installation_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_mattermost_com_cluster_installation_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/mattermost_com_v1alpha1/mattermost_com_cluster_installation_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-mattermost_com_mattermost_restore_db_v1alpha1_manifest_test.go: out/install-sentinel terratest/mattermost_com_v1alpha1/mattermost_com_mattermost_restore_db_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_mattermost_com_mattermost_restore_db_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/mattermost_com_v1alpha1/mattermost_com_mattermost_restore_db_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-memorydb_services_k8s_aws_acl_v1alpha1_manifest_test.go: out/install-sentinel terratest/memorydb_services_k8s_aws_v1alpha1/memorydb_services_k8s_aws_acl_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_memorydb_services_k8s_aws_acl_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/memorydb_services_k8s_aws_v1alpha1/memorydb_services_k8s_aws_acl_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-memorydb_services_k8s_aws_cluster_v1alpha1_manifest_test.go: out/install-sentinel terratest/memorydb_services_k8s_aws_v1alpha1/memorydb_services_k8s_aws_cluster_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_memorydb_services_k8s_aws_cluster_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/memorydb_services_k8s_aws_v1alpha1/memorydb_services_k8s_aws_cluster_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-memorydb_services_k8s_aws_parameter_group_v1alpha1_manifest_test.go: out/install-sentinel terratest/memorydb_services_k8s_aws_v1alpha1/memorydb_services_k8s_aws_parameter_group_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_memorydb_services_k8s_aws_parameter_group_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/memorydb_services_k8s_aws_v1alpha1/memorydb_services_k8s_aws_parameter_group_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-memorydb_services_k8s_aws_snapshot_v1alpha1_manifest_test.go: out/install-sentinel terratest/memorydb_services_k8s_aws_v1alpha1/memorydb_services_k8s_aws_snapshot_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_memorydb_services_k8s_aws_snapshot_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/memorydb_services_k8s_aws_v1alpha1/memorydb_services_k8s_aws_snapshot_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-memorydb_services_k8s_aws_subnet_group_v1alpha1_manifest_test.go: out/install-sentinel terratest/memorydb_services_k8s_aws_v1alpha1/memorydb_services_k8s_aws_subnet_group_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_memorydb_services_k8s_aws_subnet_group_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/memorydb_services_k8s_aws_v1alpha1/memorydb_services_k8s_aws_subnet_group_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-memorydb_services_k8s_aws_user_v1alpha1_manifest_test.go: out/install-sentinel terratest/memorydb_services_k8s_aws_v1alpha1/memorydb_services_k8s_aws_user_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_memorydb_services_k8s_aws_user_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/memorydb_services_k8s_aws_v1alpha1/memorydb_services_k8s_aws_user_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-metacontroller_k8s_io_composite_controller_v1alpha1_manifest_test.go: out/install-sentinel terratest/metacontroller_k8s_io_v1alpha1/metacontroller_k8s_io_composite_controller_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_metacontroller_k8s_io_composite_controller_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/metacontroller_k8s_io_v1alpha1/metacontroller_k8s_io_composite_controller_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-metacontroller_k8s_io_controller_revision_v1alpha1_manifest_test.go: out/install-sentinel terratest/metacontroller_k8s_io_v1alpha1/metacontroller_k8s_io_controller_revision_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_metacontroller_k8s_io_controller_revision_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/metacontroller_k8s_io_v1alpha1/metacontroller_k8s_io_controller_revision_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-metacontroller_k8s_io_decorator_controller_v1alpha1_manifest_test.go: out/install-sentinel terratest/metacontroller_k8s_io_v1alpha1/metacontroller_k8s_io_decorator_controller_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_metacontroller_k8s_io_decorator_controller_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/metacontroller_k8s_io_v1alpha1/metacontroller_k8s_io_decorator_controller_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-metal3_io_bare_metal_host_v1alpha1_manifest_test.go: out/install-sentinel terratest/metal3_io_v1alpha1/metal3_io_bare_metal_host_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_metal3_io_bare_metal_host_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/metal3_io_v1alpha1/metal3_io_bare_metal_host_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-metal3_io_bmc_event_subscription_v1alpha1_manifest_test.go: out/install-sentinel terratest/metal3_io_v1alpha1/metal3_io_bmc_event_subscription_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_metal3_io_bmc_event_subscription_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/metal3_io_v1alpha1/metal3_io_bmc_event_subscription_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-metal3_io_data_image_v1alpha1_manifest_test.go: out/install-sentinel terratest/metal3_io_v1alpha1/metal3_io_data_image_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_metal3_io_data_image_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/metal3_io_v1alpha1/metal3_io_data_image_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-metal3_io_firmware_schema_v1alpha1_manifest_test.go: out/install-sentinel terratest/metal3_io_v1alpha1/metal3_io_firmware_schema_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_metal3_io_firmware_schema_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/metal3_io_v1alpha1/metal3_io_firmware_schema_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-metal3_io_hardware_data_v1alpha1_manifest_test.go: out/install-sentinel terratest/metal3_io_v1alpha1/metal3_io_hardware_data_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_metal3_io_hardware_data_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/metal3_io_v1alpha1/metal3_io_hardware_data_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-metal3_io_host_firmware_components_v1alpha1_manifest_test.go: out/install-sentinel terratest/metal3_io_v1alpha1/metal3_io_host_firmware_components_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_metal3_io_host_firmware_components_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/metal3_io_v1alpha1/metal3_io_host_firmware_components_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-metal3_io_host_firmware_settings_v1alpha1_manifest_test.go: out/install-sentinel terratest/metal3_io_v1alpha1/metal3_io_host_firmware_settings_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_metal3_io_host_firmware_settings_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/metal3_io_v1alpha1/metal3_io_host_firmware_settings_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-metal3_io_preprovisioning_image_v1alpha1_manifest_test.go: out/install-sentinel terratest/metal3_io_v1alpha1/metal3_io_preprovisioning_image_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_metal3_io_preprovisioning_image_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/metal3_io_v1alpha1/metal3_io_preprovisioning_image_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-minio_min_io_tenant_v2_manifest_test.go: out/install-sentinel terratest/minio_min_io_v2/minio_min_io_tenant_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_minio_min_io_tenant_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/minio_min_io_v2/minio_min_io_tenant_v2_manifest_test.go
	touch $@
out/terratest-sentinel-mirrors_kts_studio_secret_mirror_v1alpha1_manifest_test.go: out/install-sentinel terratest/mirrors_kts_studio_v1alpha1/mirrors_kts_studio_secret_mirror_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_mirrors_kts_studio_secret_mirror_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/mirrors_kts_studio_v1alpha1/mirrors_kts_studio_secret_mirror_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-mirrors_kts_studio_secret_mirror_v1alpha2_manifest_test.go: out/install-sentinel terratest/mirrors_kts_studio_v1alpha2/mirrors_kts_studio_secret_mirror_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_mirrors_kts_studio_secret_mirror_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/mirrors_kts_studio_v1alpha2/mirrors_kts_studio_secret_mirror_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-model_kubedl_io_model_v1alpha1_manifest_test.go: out/install-sentinel terratest/model_kubedl_io_v1alpha1/model_kubedl_io_model_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_model_kubedl_io_model_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/model_kubedl_io_v1alpha1/model_kubedl_io_model_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-model_kubedl_io_model_version_v1alpha1_manifest_test.go: out/install-sentinel terratest/model_kubedl_io_v1alpha1/model_kubedl_io_model_version_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_model_kubedl_io_model_version_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/model_kubedl_io_v1alpha1/model_kubedl_io_model_version_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-monitoring_coreos_com_alertmanager_v1_manifest_test.go: out/install-sentinel terratest/monitoring_coreos_com_v1/monitoring_coreos_com_alertmanager_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_monitoring_coreos_com_alertmanager_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/monitoring_coreos_com_v1/monitoring_coreos_com_alertmanager_v1_manifest_test.go
	touch $@
out/terratest-sentinel-monitoring_coreos_com_pod_monitor_v1_manifest_test.go: out/install-sentinel terratest/monitoring_coreos_com_v1/monitoring_coreos_com_pod_monitor_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_monitoring_coreos_com_pod_monitor_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/monitoring_coreos_com_v1/monitoring_coreos_com_pod_monitor_v1_manifest_test.go
	touch $@
out/terratest-sentinel-monitoring_coreos_com_probe_v1_manifest_test.go: out/install-sentinel terratest/monitoring_coreos_com_v1/monitoring_coreos_com_probe_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_monitoring_coreos_com_probe_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/monitoring_coreos_com_v1/monitoring_coreos_com_probe_v1_manifest_test.go
	touch $@
out/terratest-sentinel-monitoring_coreos_com_prometheus_rule_v1_manifest_test.go: out/install-sentinel terratest/monitoring_coreos_com_v1/monitoring_coreos_com_prometheus_rule_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_monitoring_coreos_com_prometheus_rule_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/monitoring_coreos_com_v1/monitoring_coreos_com_prometheus_rule_v1_manifest_test.go
	touch $@
out/terratest-sentinel-monitoring_coreos_com_prometheus_v1_manifest_test.go: out/install-sentinel terratest/monitoring_coreos_com_v1/monitoring_coreos_com_prometheus_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_monitoring_coreos_com_prometheus_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/monitoring_coreos_com_v1/monitoring_coreos_com_prometheus_v1_manifest_test.go
	touch $@
out/terratest-sentinel-monitoring_coreos_com_service_monitor_v1_manifest_test.go: out/install-sentinel terratest/monitoring_coreos_com_v1/monitoring_coreos_com_service_monitor_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_monitoring_coreos_com_service_monitor_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/monitoring_coreos_com_v1/monitoring_coreos_com_service_monitor_v1_manifest_test.go
	touch $@
out/terratest-sentinel-monitoring_coreos_com_thanos_ruler_v1_manifest_test.go: out/install-sentinel terratest/monitoring_coreos_com_v1/monitoring_coreos_com_thanos_ruler_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_monitoring_coreos_com_thanos_ruler_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/monitoring_coreos_com_v1/monitoring_coreos_com_thanos_ruler_v1_manifest_test.go
	touch $@
out/terratest-sentinel-monitoring_coreos_com_alertmanager_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/monitoring_coreos_com_v1alpha1/monitoring_coreos_com_alertmanager_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_monitoring_coreos_com_alertmanager_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/monitoring_coreos_com_v1alpha1/monitoring_coreos_com_alertmanager_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-monitoring_coreos_com_prometheus_agent_v1alpha1_manifest_test.go: out/install-sentinel terratest/monitoring_coreos_com_v1alpha1/monitoring_coreos_com_prometheus_agent_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_monitoring_coreos_com_prometheus_agent_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/monitoring_coreos_com_v1alpha1/monitoring_coreos_com_prometheus_agent_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-monitoring_coreos_com_scrape_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/monitoring_coreos_com_v1alpha1/monitoring_coreos_com_scrape_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_monitoring_coreos_com_scrape_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/monitoring_coreos_com_v1alpha1/monitoring_coreos_com_scrape_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-monitoring_coreos_com_alertmanager_config_v1beta1_manifest_test.go: out/install-sentinel terratest/monitoring_coreos_com_v1beta1/monitoring_coreos_com_alertmanager_config_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_monitoring_coreos_com_alertmanager_config_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/monitoring_coreos_com_v1beta1/monitoring_coreos_com_alertmanager_config_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-monocle_monocle_change_metrics_io_monocle_v1alpha1_manifest_test.go: out/install-sentinel terratest/monocle_monocle_change_metrics_io_v1alpha1/monocle_monocle_change_metrics_io_monocle_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_monocle_monocle_change_metrics_io_monocle_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/monocle_monocle_change_metrics_io_v1alpha1/monocle_monocle_change_metrics_io_monocle_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-mq_services_k8s_aws_broker_v1alpha1_manifest_test.go: out/install-sentinel terratest/mq_services_k8s_aws_v1alpha1/mq_services_k8s_aws_broker_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_mq_services_k8s_aws_broker_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/mq_services_k8s_aws_v1alpha1/mq_services_k8s_aws_broker_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-multicluster_crd_antrea_io_cluster_info_import_v1alpha1_manifest_test.go: out/install-sentinel terratest/multicluster_crd_antrea_io_v1alpha1/multicluster_crd_antrea_io_cluster_info_import_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_multicluster_crd_antrea_io_cluster_info_import_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/multicluster_crd_antrea_io_v1alpha1/multicluster_crd_antrea_io_cluster_info_import_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-multicluster_crd_antrea_io_cluster_set_v1alpha1_manifest_test.go: out/install-sentinel terratest/multicluster_crd_antrea_io_v1alpha1/multicluster_crd_antrea_io_cluster_set_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_multicluster_crd_antrea_io_cluster_set_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/multicluster_crd_antrea_io_v1alpha1/multicluster_crd_antrea_io_cluster_set_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-multicluster_crd_antrea_io_gateway_v1alpha1_manifest_test.go: out/install-sentinel terratest/multicluster_crd_antrea_io_v1alpha1/multicluster_crd_antrea_io_gateway_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_multicluster_crd_antrea_io_gateway_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/multicluster_crd_antrea_io_v1alpha1/multicluster_crd_antrea_io_gateway_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-multicluster_crd_antrea_io_label_identity_v1alpha1_manifest_test.go: out/install-sentinel terratest/multicluster_crd_antrea_io_v1alpha1/multicluster_crd_antrea_io_label_identity_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_multicluster_crd_antrea_io_label_identity_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/multicluster_crd_antrea_io_v1alpha1/multicluster_crd_antrea_io_label_identity_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-multicluster_crd_antrea_io_member_cluster_announce_v1alpha1_manifest_test.go: out/install-sentinel terratest/multicluster_crd_antrea_io_v1alpha1/multicluster_crd_antrea_io_member_cluster_announce_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_multicluster_crd_antrea_io_member_cluster_announce_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/multicluster_crd_antrea_io_v1alpha1/multicluster_crd_antrea_io_member_cluster_announce_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-multicluster_crd_antrea_io_multi_cluster_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/multicluster_crd_antrea_io_v1alpha1/multicluster_crd_antrea_io_multi_cluster_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_multicluster_crd_antrea_io_multi_cluster_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/multicluster_crd_antrea_io_v1alpha1/multicluster_crd_antrea_io_multi_cluster_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-multicluster_crd_antrea_io_resource_export_v1alpha1_manifest_test.go: out/install-sentinel terratest/multicluster_crd_antrea_io_v1alpha1/multicluster_crd_antrea_io_resource_export_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_multicluster_crd_antrea_io_resource_export_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/multicluster_crd_antrea_io_v1alpha1/multicluster_crd_antrea_io_resource_export_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-multicluster_crd_antrea_io_resource_import_v1alpha1_manifest_test.go: out/install-sentinel terratest/multicluster_crd_antrea_io_v1alpha1/multicluster_crd_antrea_io_resource_import_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_multicluster_crd_antrea_io_resource_import_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/multicluster_crd_antrea_io_v1alpha1/multicluster_crd_antrea_io_resource_import_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-multicluster_crd_antrea_io_cluster_claim_v1alpha2_manifest_test.go: out/install-sentinel terratest/multicluster_crd_antrea_io_v1alpha2/multicluster_crd_antrea_io_cluster_claim_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_multicluster_crd_antrea_io_cluster_claim_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/multicluster_crd_antrea_io_v1alpha2/multicluster_crd_antrea_io_cluster_claim_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-multicluster_crd_antrea_io_cluster_set_v1alpha2_manifest_test.go: out/install-sentinel terratest/multicluster_crd_antrea_io_v1alpha2/multicluster_crd_antrea_io_cluster_set_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_multicluster_crd_antrea_io_cluster_set_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/multicluster_crd_antrea_io_v1alpha2/multicluster_crd_antrea_io_cluster_set_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-multicluster_x_k8s_io_applied_work_v1alpha1_manifest_test.go: out/install-sentinel terratest/multicluster_x_k8s_io_v1alpha1/multicluster_x_k8s_io_applied_work_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_multicluster_x_k8s_io_applied_work_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/multicluster_x_k8s_io_v1alpha1/multicluster_x_k8s_io_applied_work_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-multicluster_x_k8s_io_service_import_v1alpha1_manifest_test.go: out/install-sentinel terratest/multicluster_x_k8s_io_v1alpha1/multicluster_x_k8s_io_service_import_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_multicluster_x_k8s_io_service_import_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/multicluster_x_k8s_io_v1alpha1/multicluster_x_k8s_io_service_import_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-multicluster_x_k8s_io_work_v1alpha1_manifest_test.go: out/install-sentinel terratest/multicluster_x_k8s_io_v1alpha1/multicluster_x_k8s_io_work_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_multicluster_x_k8s_io_work_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/multicluster_x_k8s_io_v1alpha1/multicluster_x_k8s_io_work_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-mutations_gatekeeper_sh_assign_metadata_v1_manifest_test.go: out/install-sentinel terratest/mutations_gatekeeper_sh_v1/mutations_gatekeeper_sh_assign_metadata_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_mutations_gatekeeper_sh_assign_metadata_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/mutations_gatekeeper_sh_v1/mutations_gatekeeper_sh_assign_metadata_v1_manifest_test.go
	touch $@
out/terratest-sentinel-mutations_gatekeeper_sh_assign_v1_manifest_test.go: out/install-sentinel terratest/mutations_gatekeeper_sh_v1/mutations_gatekeeper_sh_assign_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_mutations_gatekeeper_sh_assign_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/mutations_gatekeeper_sh_v1/mutations_gatekeeper_sh_assign_v1_manifest_test.go
	touch $@
out/terratest-sentinel-mutations_gatekeeper_sh_modify_set_v1_manifest_test.go: out/install-sentinel terratest/mutations_gatekeeper_sh_v1/mutations_gatekeeper_sh_modify_set_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_mutations_gatekeeper_sh_modify_set_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/mutations_gatekeeper_sh_v1/mutations_gatekeeper_sh_modify_set_v1_manifest_test.go
	touch $@
out/terratest-sentinel-mutations_gatekeeper_sh_assign_image_v1alpha1_manifest_test.go: out/install-sentinel terratest/mutations_gatekeeper_sh_v1alpha1/mutations_gatekeeper_sh_assign_image_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_mutations_gatekeeper_sh_assign_image_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/mutations_gatekeeper_sh_v1alpha1/mutations_gatekeeper_sh_assign_image_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-mutations_gatekeeper_sh_assign_metadata_v1alpha1_manifest_test.go: out/install-sentinel terratest/mutations_gatekeeper_sh_v1alpha1/mutations_gatekeeper_sh_assign_metadata_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_mutations_gatekeeper_sh_assign_metadata_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/mutations_gatekeeper_sh_v1alpha1/mutations_gatekeeper_sh_assign_metadata_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-mutations_gatekeeper_sh_assign_v1alpha1_manifest_test.go: out/install-sentinel terratest/mutations_gatekeeper_sh_v1alpha1/mutations_gatekeeper_sh_assign_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_mutations_gatekeeper_sh_assign_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/mutations_gatekeeper_sh_v1alpha1/mutations_gatekeeper_sh_assign_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-mutations_gatekeeper_sh_modify_set_v1alpha1_manifest_test.go: out/install-sentinel terratest/mutations_gatekeeper_sh_v1alpha1/mutations_gatekeeper_sh_modify_set_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_mutations_gatekeeper_sh_modify_set_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/mutations_gatekeeper_sh_v1alpha1/mutations_gatekeeper_sh_modify_set_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-mutations_gatekeeper_sh_assign_metadata_v1beta1_manifest_test.go: out/install-sentinel terratest/mutations_gatekeeper_sh_v1beta1/mutations_gatekeeper_sh_assign_metadata_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_mutations_gatekeeper_sh_assign_metadata_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/mutations_gatekeeper_sh_v1beta1/mutations_gatekeeper_sh_assign_metadata_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-mutations_gatekeeper_sh_assign_v1beta1_manifest_test.go: out/install-sentinel terratest/mutations_gatekeeper_sh_v1beta1/mutations_gatekeeper_sh_assign_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_mutations_gatekeeper_sh_assign_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/mutations_gatekeeper_sh_v1beta1/mutations_gatekeeper_sh_assign_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-mutations_gatekeeper_sh_modify_set_v1beta1_manifest_test.go: out/install-sentinel terratest/mutations_gatekeeper_sh_v1beta1/mutations_gatekeeper_sh_modify_set_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_mutations_gatekeeper_sh_modify_set_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/mutations_gatekeeper_sh_v1beta1/mutations_gatekeeper_sh_modify_set_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-nativestor_alauda_io_raw_device_v1_manifest_test.go: out/install-sentinel terratest/nativestor_alauda_io_v1/nativestor_alauda_io_raw_device_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_nativestor_alauda_io_raw_device_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/nativestor_alauda_io_v1/nativestor_alauda_io_raw_device_v1_manifest_test.go
	touch $@
out/terratest-sentinel-netchecks_io_network_assertion_v1_manifest_test.go: out/install-sentinel terratest/netchecks_io_v1/netchecks_io_network_assertion_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_netchecks_io_network_assertion_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/netchecks_io_v1/netchecks_io_network_assertion_v1_manifest_test.go
	touch $@
out/terratest-sentinel-networkfirewall_services_k8s_aws_firewall_policy_v1alpha1_manifest_test.go: out/install-sentinel terratest/networkfirewall_services_k8s_aws_v1alpha1/networkfirewall_services_k8s_aws_firewall_policy_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_networkfirewall_services_k8s_aws_firewall_policy_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networkfirewall_services_k8s_aws_v1alpha1/networkfirewall_services_k8s_aws_firewall_policy_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-networkfirewall_services_k8s_aws_firewall_v1alpha1_manifest_test.go: out/install-sentinel terratest/networkfirewall_services_k8s_aws_v1alpha1/networkfirewall_services_k8s_aws_firewall_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_networkfirewall_services_k8s_aws_firewall_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networkfirewall_services_k8s_aws_v1alpha1/networkfirewall_services_k8s_aws_firewall_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-networkfirewall_services_k8s_aws_rule_group_v1alpha1_manifest_test.go: out/install-sentinel terratest/networkfirewall_services_k8s_aws_v1alpha1/networkfirewall_services_k8s_aws_rule_group_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_networkfirewall_services_k8s_aws_rule_group_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networkfirewall_services_k8s_aws_v1alpha1/networkfirewall_services_k8s_aws_rule_group_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-networking_gke_io_gcp_backend_policy_v1_manifest_test.go: out/install-sentinel terratest/networking_gke_io_v1/networking_gke_io_gcp_backend_policy_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_networking_gke_io_gcp_backend_policy_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networking_gke_io_v1/networking_gke_io_gcp_backend_policy_v1_manifest_test.go
	touch $@
out/terratest-sentinel-networking_gke_io_gcp_gateway_policy_v1_manifest_test.go: out/install-sentinel terratest/networking_gke_io_v1/networking_gke_io_gcp_gateway_policy_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_networking_gke_io_gcp_gateway_policy_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networking_gke_io_v1/networking_gke_io_gcp_gateway_policy_v1_manifest_test.go
	touch $@
out/terratest-sentinel-networking_gke_io_health_check_policy_v1_manifest_test.go: out/install-sentinel terratest/networking_gke_io_v1/networking_gke_io_health_check_policy_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_networking_gke_io_health_check_policy_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networking_gke_io_v1/networking_gke_io_health_check_policy_v1_manifest_test.go
	touch $@
out/terratest-sentinel-networking_gke_io_lb_policy_v1_manifest_test.go: out/install-sentinel terratest/networking_gke_io_v1/networking_gke_io_lb_policy_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_networking_gke_io_lb_policy_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networking_gke_io_v1/networking_gke_io_lb_policy_v1_manifest_test.go
	touch $@
out/terratest-sentinel-networking_gke_io_managed_certificate_v1_manifest_test.go: out/install-sentinel terratest/networking_gke_io_v1/networking_gke_io_managed_certificate_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_networking_gke_io_managed_certificate_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networking_gke_io_v1/networking_gke_io_managed_certificate_v1_manifest_test.go
	touch $@
out/terratest-sentinel-networking_istio_io_destination_rule_v1_manifest_test.go: out/install-sentinel terratest/networking_istio_io_v1/networking_istio_io_destination_rule_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_networking_istio_io_destination_rule_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networking_istio_io_v1/networking_istio_io_destination_rule_v1_manifest_test.go
	touch $@
out/terratest-sentinel-networking_istio_io_gateway_v1_manifest_test.go: out/install-sentinel terratest/networking_istio_io_v1/networking_istio_io_gateway_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_networking_istio_io_gateway_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networking_istio_io_v1/networking_istio_io_gateway_v1_manifest_test.go
	touch $@
out/terratest-sentinel-networking_istio_io_service_entry_v1_manifest_test.go: out/install-sentinel terratest/networking_istio_io_v1/networking_istio_io_service_entry_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_networking_istio_io_service_entry_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networking_istio_io_v1/networking_istio_io_service_entry_v1_manifest_test.go
	touch $@
out/terratest-sentinel-networking_istio_io_sidecar_v1_manifest_test.go: out/install-sentinel terratest/networking_istio_io_v1/networking_istio_io_sidecar_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_networking_istio_io_sidecar_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networking_istio_io_v1/networking_istio_io_sidecar_v1_manifest_test.go
	touch $@
out/terratest-sentinel-networking_istio_io_virtual_service_v1_manifest_test.go: out/install-sentinel terratest/networking_istio_io_v1/networking_istio_io_virtual_service_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_networking_istio_io_virtual_service_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networking_istio_io_v1/networking_istio_io_virtual_service_v1_manifest_test.go
	touch $@
out/terratest-sentinel-networking_istio_io_workload_entry_v1_manifest_test.go: out/install-sentinel terratest/networking_istio_io_v1/networking_istio_io_workload_entry_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_networking_istio_io_workload_entry_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networking_istio_io_v1/networking_istio_io_workload_entry_v1_manifest_test.go
	touch $@
out/terratest-sentinel-networking_istio_io_workload_group_v1_manifest_test.go: out/install-sentinel terratest/networking_istio_io_v1/networking_istio_io_workload_group_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_networking_istio_io_workload_group_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networking_istio_io_v1/networking_istio_io_workload_group_v1_manifest_test.go
	touch $@
out/terratest-sentinel-networking_istio_io_destination_rule_v1alpha3_manifest_test.go: out/install-sentinel terratest/networking_istio_io_v1alpha3/networking_istio_io_destination_rule_v1alpha3_manifest_test.go $(shell find ./examples/data-sources/k8s_networking_istio_io_destination_rule_v1alpha3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networking_istio_io_v1alpha3/networking_istio_io_destination_rule_v1alpha3_manifest_test.go
	touch $@
out/terratest-sentinel-networking_istio_io_envoy_filter_v1alpha3_manifest_test.go: out/install-sentinel terratest/networking_istio_io_v1alpha3/networking_istio_io_envoy_filter_v1alpha3_manifest_test.go $(shell find ./examples/data-sources/k8s_networking_istio_io_envoy_filter_v1alpha3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networking_istio_io_v1alpha3/networking_istio_io_envoy_filter_v1alpha3_manifest_test.go
	touch $@
out/terratest-sentinel-networking_istio_io_gateway_v1alpha3_manifest_test.go: out/install-sentinel terratest/networking_istio_io_v1alpha3/networking_istio_io_gateway_v1alpha3_manifest_test.go $(shell find ./examples/data-sources/k8s_networking_istio_io_gateway_v1alpha3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networking_istio_io_v1alpha3/networking_istio_io_gateway_v1alpha3_manifest_test.go
	touch $@
out/terratest-sentinel-networking_istio_io_service_entry_v1alpha3_manifest_test.go: out/install-sentinel terratest/networking_istio_io_v1alpha3/networking_istio_io_service_entry_v1alpha3_manifest_test.go $(shell find ./examples/data-sources/k8s_networking_istio_io_service_entry_v1alpha3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networking_istio_io_v1alpha3/networking_istio_io_service_entry_v1alpha3_manifest_test.go
	touch $@
out/terratest-sentinel-networking_istio_io_sidecar_v1alpha3_manifest_test.go: out/install-sentinel terratest/networking_istio_io_v1alpha3/networking_istio_io_sidecar_v1alpha3_manifest_test.go $(shell find ./examples/data-sources/k8s_networking_istio_io_sidecar_v1alpha3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networking_istio_io_v1alpha3/networking_istio_io_sidecar_v1alpha3_manifest_test.go
	touch $@
out/terratest-sentinel-networking_istio_io_virtual_service_v1alpha3_manifest_test.go: out/install-sentinel terratest/networking_istio_io_v1alpha3/networking_istio_io_virtual_service_v1alpha3_manifest_test.go $(shell find ./examples/data-sources/k8s_networking_istio_io_virtual_service_v1alpha3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networking_istio_io_v1alpha3/networking_istio_io_virtual_service_v1alpha3_manifest_test.go
	touch $@
out/terratest-sentinel-networking_istio_io_workload_entry_v1alpha3_manifest_test.go: out/install-sentinel terratest/networking_istio_io_v1alpha3/networking_istio_io_workload_entry_v1alpha3_manifest_test.go $(shell find ./examples/data-sources/k8s_networking_istio_io_workload_entry_v1alpha3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networking_istio_io_v1alpha3/networking_istio_io_workload_entry_v1alpha3_manifest_test.go
	touch $@
out/terratest-sentinel-networking_istio_io_workload_group_v1alpha3_manifest_test.go: out/install-sentinel terratest/networking_istio_io_v1alpha3/networking_istio_io_workload_group_v1alpha3_manifest_test.go $(shell find ./examples/data-sources/k8s_networking_istio_io_workload_group_v1alpha3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networking_istio_io_v1alpha3/networking_istio_io_workload_group_v1alpha3_manifest_test.go
	touch $@
out/terratest-sentinel-networking_istio_io_destination_rule_v1beta1_manifest_test.go: out/install-sentinel terratest/networking_istio_io_v1beta1/networking_istio_io_destination_rule_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_networking_istio_io_destination_rule_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networking_istio_io_v1beta1/networking_istio_io_destination_rule_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-networking_istio_io_gateway_v1beta1_manifest_test.go: out/install-sentinel terratest/networking_istio_io_v1beta1/networking_istio_io_gateway_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_networking_istio_io_gateway_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networking_istio_io_v1beta1/networking_istio_io_gateway_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-networking_istio_io_proxy_config_v1beta1_manifest_test.go: out/install-sentinel terratest/networking_istio_io_v1beta1/networking_istio_io_proxy_config_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_networking_istio_io_proxy_config_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networking_istio_io_v1beta1/networking_istio_io_proxy_config_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-networking_istio_io_service_entry_v1beta1_manifest_test.go: out/install-sentinel terratest/networking_istio_io_v1beta1/networking_istio_io_service_entry_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_networking_istio_io_service_entry_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networking_istio_io_v1beta1/networking_istio_io_service_entry_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-networking_istio_io_sidecar_v1beta1_manifest_test.go: out/install-sentinel terratest/networking_istio_io_v1beta1/networking_istio_io_sidecar_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_networking_istio_io_sidecar_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networking_istio_io_v1beta1/networking_istio_io_sidecar_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-networking_istio_io_virtual_service_v1beta1_manifest_test.go: out/install-sentinel terratest/networking_istio_io_v1beta1/networking_istio_io_virtual_service_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_networking_istio_io_virtual_service_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networking_istio_io_v1beta1/networking_istio_io_virtual_service_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-networking_istio_io_workload_entry_v1beta1_manifest_test.go: out/install-sentinel terratest/networking_istio_io_v1beta1/networking_istio_io_workload_entry_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_networking_istio_io_workload_entry_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networking_istio_io_v1beta1/networking_istio_io_workload_entry_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-networking_istio_io_workload_group_v1beta1_manifest_test.go: out/install-sentinel terratest/networking_istio_io_v1beta1/networking_istio_io_workload_group_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_networking_istio_io_workload_group_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networking_istio_io_v1beta1/networking_istio_io_workload_group_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-networking_k8s_aws_policy_endpoint_v1alpha1_manifest_test.go: out/install-sentinel terratest/networking_k8s_aws_v1alpha1/networking_k8s_aws_policy_endpoint_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_networking_k8s_aws_policy_endpoint_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networking_k8s_aws_v1alpha1/networking_k8s_aws_policy_endpoint_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-networking_k8s_io_ingress_class_v1_manifest_test.go: out/install-sentinel terratest/networking_k8s_io_v1/networking_k8s_io_ingress_class_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_networking_k8s_io_ingress_class_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networking_k8s_io_v1/networking_k8s_io_ingress_class_v1_manifest_test.go
	touch $@
out/terratest-sentinel-networking_k8s_io_ingress_v1_manifest_test.go: out/install-sentinel terratest/networking_k8s_io_v1/networking_k8s_io_ingress_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_networking_k8s_io_ingress_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networking_k8s_io_v1/networking_k8s_io_ingress_v1_manifest_test.go
	touch $@
out/terratest-sentinel-networking_k8s_io_network_policy_v1_manifest_test.go: out/install-sentinel terratest/networking_k8s_io_v1/networking_k8s_io_network_policy_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_networking_k8s_io_network_policy_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networking_k8s_io_v1/networking_k8s_io_network_policy_v1_manifest_test.go
	touch $@
out/terratest-sentinel-networking_karmada_io_multi_cluster_ingress_v1alpha1_manifest_test.go: out/install-sentinel terratest/networking_karmada_io_v1alpha1/networking_karmada_io_multi_cluster_ingress_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_networking_karmada_io_multi_cluster_ingress_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networking_karmada_io_v1alpha1/networking_karmada_io_multi_cluster_ingress_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-networking_karmada_io_multi_cluster_service_v1alpha1_manifest_test.go: out/install-sentinel terratest/networking_karmada_io_v1alpha1/networking_karmada_io_multi_cluster_service_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_networking_karmada_io_multi_cluster_service_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/networking_karmada_io_v1alpha1/networking_karmada_io_multi_cluster_service_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-nfd_k8s_sigs_io_node_feature_rule_v1alpha1_manifest_test.go: out/install-sentinel terratest/nfd_k8s_sigs_io_v1alpha1/nfd_k8s_sigs_io_node_feature_rule_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_nfd_k8s_sigs_io_node_feature_rule_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/nfd_k8s_sigs_io_v1alpha1/nfd_k8s_sigs_io_node_feature_rule_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-nfd_kubernetes_io_node_feature_discovery_v1_manifest_test.go: out/install-sentinel terratest/nfd_kubernetes_io_v1/nfd_kubernetes_io_node_feature_discovery_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_nfd_kubernetes_io_node_feature_discovery_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/nfd_kubernetes_io_v1/nfd_kubernetes_io_node_feature_discovery_v1_manifest_test.go
	touch $@
out/terratest-sentinel-nfd_kubernetes_io_node_feature_rule_v1alpha1_manifest_test.go: out/install-sentinel terratest/nfd_kubernetes_io_v1alpha1/nfd_kubernetes_io_node_feature_rule_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_nfd_kubernetes_io_node_feature_rule_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/nfd_kubernetes_io_v1alpha1/nfd_kubernetes_io_node_feature_rule_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-nodeinfo_volcano_sh_numatopology_v1alpha1_manifest_test.go: out/install-sentinel terratest/nodeinfo_volcano_sh_v1alpha1/nodeinfo_volcano_sh_numatopology_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_nodeinfo_volcano_sh_numatopology_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/nodeinfo_volcano_sh_v1alpha1/nodeinfo_volcano_sh_numatopology_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-notebook_kubedl_io_notebook_v1alpha1_manifest_test.go: out/install-sentinel terratest/notebook_kubedl_io_v1alpha1/notebook_kubedl_io_notebook_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_notebook_kubedl_io_notebook_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/notebook_kubedl_io_v1alpha1/notebook_kubedl_io_notebook_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-notification_toolkit_fluxcd_io_receiver_v1_manifest_test.go: out/install-sentinel terratest/notification_toolkit_fluxcd_io_v1/notification_toolkit_fluxcd_io_receiver_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_notification_toolkit_fluxcd_io_receiver_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/notification_toolkit_fluxcd_io_v1/notification_toolkit_fluxcd_io_receiver_v1_manifest_test.go
	touch $@
out/terratest-sentinel-notification_toolkit_fluxcd_io_alert_v1beta1_manifest_test.go: out/install-sentinel terratest/notification_toolkit_fluxcd_io_v1beta1/notification_toolkit_fluxcd_io_alert_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_notification_toolkit_fluxcd_io_alert_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/notification_toolkit_fluxcd_io_v1beta1/notification_toolkit_fluxcd_io_alert_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-notification_toolkit_fluxcd_io_provider_v1beta1_manifest_test.go: out/install-sentinel terratest/notification_toolkit_fluxcd_io_v1beta1/notification_toolkit_fluxcd_io_provider_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_notification_toolkit_fluxcd_io_provider_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/notification_toolkit_fluxcd_io_v1beta1/notification_toolkit_fluxcd_io_provider_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-notification_toolkit_fluxcd_io_receiver_v1beta1_manifest_test.go: out/install-sentinel terratest/notification_toolkit_fluxcd_io_v1beta1/notification_toolkit_fluxcd_io_receiver_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_notification_toolkit_fluxcd_io_receiver_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/notification_toolkit_fluxcd_io_v1beta1/notification_toolkit_fluxcd_io_receiver_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-notification_toolkit_fluxcd_io_alert_v1beta2_manifest_test.go: out/install-sentinel terratest/notification_toolkit_fluxcd_io_v1beta2/notification_toolkit_fluxcd_io_alert_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_notification_toolkit_fluxcd_io_alert_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/notification_toolkit_fluxcd_io_v1beta2/notification_toolkit_fluxcd_io_alert_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-notification_toolkit_fluxcd_io_provider_v1beta2_manifest_test.go: out/install-sentinel terratest/notification_toolkit_fluxcd_io_v1beta2/notification_toolkit_fluxcd_io_provider_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_notification_toolkit_fluxcd_io_provider_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/notification_toolkit_fluxcd_io_v1beta2/notification_toolkit_fluxcd_io_provider_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-notification_toolkit_fluxcd_io_receiver_v1beta2_manifest_test.go: out/install-sentinel terratest/notification_toolkit_fluxcd_io_v1beta2/notification_toolkit_fluxcd_io_receiver_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_notification_toolkit_fluxcd_io_receiver_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/notification_toolkit_fluxcd_io_v1beta2/notification_toolkit_fluxcd_io_receiver_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-notification_toolkit_fluxcd_io_alert_v1beta3_manifest_test.go: out/install-sentinel terratest/notification_toolkit_fluxcd_io_v1beta3/notification_toolkit_fluxcd_io_alert_v1beta3_manifest_test.go $(shell find ./examples/data-sources/k8s_notification_toolkit_fluxcd_io_alert_v1beta3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/notification_toolkit_fluxcd_io_v1beta3/notification_toolkit_fluxcd_io_alert_v1beta3_manifest_test.go
	touch $@
out/terratest-sentinel-notification_toolkit_fluxcd_io_provider_v1beta3_manifest_test.go: out/install-sentinel terratest/notification_toolkit_fluxcd_io_v1beta3/notification_toolkit_fluxcd_io_provider_v1beta3_manifest_test.go $(shell find ./examples/data-sources/k8s_notification_toolkit_fluxcd_io_provider_v1beta3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/notification_toolkit_fluxcd_io_v1beta3/notification_toolkit_fluxcd_io_provider_v1beta3_manifest_test.go
	touch $@
out/terratest-sentinel-objectbucket_io_object_bucket_claim_v1alpha1_manifest_test.go: out/install-sentinel terratest/objectbucket_io_v1alpha1/objectbucket_io_object_bucket_claim_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_objectbucket_io_object_bucket_claim_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/objectbucket_io_v1alpha1/objectbucket_io_object_bucket_claim_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-objectbucket_io_object_bucket_v1alpha1_manifest_test.go: out/install-sentinel terratest/objectbucket_io_v1alpha1/objectbucket_io_object_bucket_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_objectbucket_io_object_bucket_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/objectbucket_io_v1alpha1/objectbucket_io_object_bucket_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-onepassword_com_one_password_item_v1_manifest_test.go: out/install-sentinel terratest/onepassword_com_v1/onepassword_com_one_password_item_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_onepassword_com_one_password_item_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/onepassword_com_v1/onepassword_com_one_password_item_v1_manifest_test.go
	touch $@
out/terratest-sentinel-opensearchservice_services_k8s_aws_domain_v1alpha1_manifest_test.go: out/install-sentinel terratest/opensearchservice_services_k8s_aws_v1alpha1/opensearchservice_services_k8s_aws_domain_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_opensearchservice_services_k8s_aws_domain_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/opensearchservice_services_k8s_aws_v1alpha1/opensearchservice_services_k8s_aws_domain_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-opentelemetry_io_instrumentation_v1alpha1_manifest_test.go: out/install-sentinel terratest/opentelemetry_io_v1alpha1/opentelemetry_io_instrumentation_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_opentelemetry_io_instrumentation_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/opentelemetry_io_v1alpha1/opentelemetry_io_instrumentation_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-opentelemetry_io_op_amp_bridge_v1alpha1_manifest_test.go: out/install-sentinel terratest/opentelemetry_io_v1alpha1/opentelemetry_io_op_amp_bridge_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_opentelemetry_io_op_amp_bridge_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/opentelemetry_io_v1alpha1/opentelemetry_io_op_amp_bridge_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-opentelemetry_io_open_telemetry_collector_v1alpha1_manifest_test.go: out/install-sentinel terratest/opentelemetry_io_v1alpha1/opentelemetry_io_open_telemetry_collector_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_opentelemetry_io_open_telemetry_collector_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/opentelemetry_io_v1alpha1/opentelemetry_io_open_telemetry_collector_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-opentelemetry_io_open_telemetry_collector_v1beta1_manifest_test.go: out/install-sentinel terratest/opentelemetry_io_v1beta1/opentelemetry_io_open_telemetry_collector_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_opentelemetry_io_open_telemetry_collector_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/opentelemetry_io_v1beta1/opentelemetry_io_open_telemetry_collector_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-operations_kubeedge_io_node_upgrade_job_v1alpha1_manifest_test.go: out/install-sentinel terratest/operations_kubeedge_io_v1alpha1/operations_kubeedge_io_node_upgrade_job_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_operations_kubeedge_io_node_upgrade_job_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operations_kubeedge_io_v1alpha1/operations_kubeedge_io_node_upgrade_job_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_aquasec_com_aqua_csp_v1alpha1_manifest_test.go: out/install-sentinel terratest/operator_aquasec_com_v1alpha1/operator_aquasec_com_aqua_csp_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_aquasec_com_aqua_csp_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_aquasec_com_v1alpha1/operator_aquasec_com_aqua_csp_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_aquasec_com_aqua_database_v1alpha1_manifest_test.go: out/install-sentinel terratest/operator_aquasec_com_v1alpha1/operator_aquasec_com_aqua_database_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_aquasec_com_aqua_database_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_aquasec_com_v1alpha1/operator_aquasec_com_aqua_database_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_aquasec_com_aqua_enforcer_v1alpha1_manifest_test.go: out/install-sentinel terratest/operator_aquasec_com_v1alpha1/operator_aquasec_com_aqua_enforcer_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_aquasec_com_aqua_enforcer_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_aquasec_com_v1alpha1/operator_aquasec_com_aqua_enforcer_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_aquasec_com_aqua_gateway_v1alpha1_manifest_test.go: out/install-sentinel terratest/operator_aquasec_com_v1alpha1/operator_aquasec_com_aqua_gateway_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_aquasec_com_aqua_gateway_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_aquasec_com_v1alpha1/operator_aquasec_com_aqua_gateway_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_aquasec_com_aqua_kube_enforcer_v1alpha1_manifest_test.go: out/install-sentinel terratest/operator_aquasec_com_v1alpha1/operator_aquasec_com_aqua_kube_enforcer_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_aquasec_com_aqua_kube_enforcer_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_aquasec_com_v1alpha1/operator_aquasec_com_aqua_kube_enforcer_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_aquasec_com_aqua_scanner_v1alpha1_manifest_test.go: out/install-sentinel terratest/operator_aquasec_com_v1alpha1/operator_aquasec_com_aqua_scanner_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_aquasec_com_aqua_scanner_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_aquasec_com_v1alpha1/operator_aquasec_com_aqua_scanner_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_aquasec_com_aqua_server_v1alpha1_manifest_test.go: out/install-sentinel terratest/operator_aquasec_com_v1alpha1/operator_aquasec_com_aqua_server_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_aquasec_com_aqua_server_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_aquasec_com_v1alpha1/operator_aquasec_com_aqua_server_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_authorino_kuadrant_io_authorino_v1beta1_manifest_test.go: out/install-sentinel terratest/operator_authorino_kuadrant_io_v1beta1/operator_authorino_kuadrant_io_authorino_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_authorino_kuadrant_io_authorino_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_authorino_kuadrant_io_v1beta1/operator_authorino_kuadrant_io_authorino_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_cluster_x_k8s_io_bootstrap_provider_v1alpha1_manifest_test.go: out/install-sentinel terratest/operator_cluster_x_k8s_io_v1alpha1/operator_cluster_x_k8s_io_bootstrap_provider_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_cluster_x_k8s_io_bootstrap_provider_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_cluster_x_k8s_io_v1alpha1/operator_cluster_x_k8s_io_bootstrap_provider_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_cluster_x_k8s_io_control_plane_provider_v1alpha1_manifest_test.go: out/install-sentinel terratest/operator_cluster_x_k8s_io_v1alpha1/operator_cluster_x_k8s_io_control_plane_provider_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_cluster_x_k8s_io_control_plane_provider_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_cluster_x_k8s_io_v1alpha1/operator_cluster_x_k8s_io_control_plane_provider_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_cluster_x_k8s_io_core_provider_v1alpha1_manifest_test.go: out/install-sentinel terratest/operator_cluster_x_k8s_io_v1alpha1/operator_cluster_x_k8s_io_core_provider_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_cluster_x_k8s_io_core_provider_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_cluster_x_k8s_io_v1alpha1/operator_cluster_x_k8s_io_core_provider_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_cluster_x_k8s_io_infrastructure_provider_v1alpha1_manifest_test.go: out/install-sentinel terratest/operator_cluster_x_k8s_io_v1alpha1/operator_cluster_x_k8s_io_infrastructure_provider_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_cluster_x_k8s_io_infrastructure_provider_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_cluster_x_k8s_io_v1alpha1/operator_cluster_x_k8s_io_infrastructure_provider_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_cluster_x_k8s_io_addon_provider_v1alpha2_manifest_test.go: out/install-sentinel terratest/operator_cluster_x_k8s_io_v1alpha2/operator_cluster_x_k8s_io_addon_provider_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_cluster_x_k8s_io_addon_provider_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_cluster_x_k8s_io_v1alpha2/operator_cluster_x_k8s_io_addon_provider_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-operator_cluster_x_k8s_io_bootstrap_provider_v1alpha2_manifest_test.go: out/install-sentinel terratest/operator_cluster_x_k8s_io_v1alpha2/operator_cluster_x_k8s_io_bootstrap_provider_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_cluster_x_k8s_io_bootstrap_provider_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_cluster_x_k8s_io_v1alpha2/operator_cluster_x_k8s_io_bootstrap_provider_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-operator_cluster_x_k8s_io_control_plane_provider_v1alpha2_manifest_test.go: out/install-sentinel terratest/operator_cluster_x_k8s_io_v1alpha2/operator_cluster_x_k8s_io_control_plane_provider_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_cluster_x_k8s_io_control_plane_provider_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_cluster_x_k8s_io_v1alpha2/operator_cluster_x_k8s_io_control_plane_provider_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-operator_cluster_x_k8s_io_core_provider_v1alpha2_manifest_test.go: out/install-sentinel terratest/operator_cluster_x_k8s_io_v1alpha2/operator_cluster_x_k8s_io_core_provider_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_cluster_x_k8s_io_core_provider_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_cluster_x_k8s_io_v1alpha2/operator_cluster_x_k8s_io_core_provider_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-operator_cluster_x_k8s_io_infrastructure_provider_v1alpha2_manifest_test.go: out/install-sentinel terratest/operator_cluster_x_k8s_io_v1alpha2/operator_cluster_x_k8s_io_infrastructure_provider_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_cluster_x_k8s_io_infrastructure_provider_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_cluster_x_k8s_io_v1alpha2/operator_cluster_x_k8s_io_infrastructure_provider_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-operator_cryostat_io_cryostat_v1beta1_manifest_test.go: out/install-sentinel terratest/operator_cryostat_io_v1beta1/operator_cryostat_io_cryostat_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_cryostat_io_cryostat_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_cryostat_io_v1beta1/operator_cryostat_io_cryostat_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_cryostat_io_cryostat_v1beta2_manifest_test.go: out/install-sentinel terratest/operator_cryostat_io_v1beta2/operator_cryostat_io_cryostat_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_cryostat_io_cryostat_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_cryostat_io_v1beta2/operator_cryostat_io_cryostat_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-operator_knative_dev_knative_eventing_v1beta1_manifest_test.go: out/install-sentinel terratest/operator_knative_dev_v1beta1/operator_knative_dev_knative_eventing_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_knative_dev_knative_eventing_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_knative_dev_v1beta1/operator_knative_dev_knative_eventing_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_knative_dev_knative_serving_v1beta1_manifest_test.go: out/install-sentinel terratest/operator_knative_dev_v1beta1/operator_knative_dev_knative_serving_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_knative_dev_knative_serving_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_knative_dev_v1beta1/operator_knative_dev_knative_serving_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_marin3r_3scale_net_discovery_service_certificate_v1alpha1_manifest_test.go: out/install-sentinel terratest/operator_marin3r_3scale_net_v1alpha1/operator_marin3r_3scale_net_discovery_service_certificate_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_marin3r_3scale_net_discovery_service_certificate_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_marin3r_3scale_net_v1alpha1/operator_marin3r_3scale_net_discovery_service_certificate_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_marin3r_3scale_net_discovery_service_v1alpha1_manifest_test.go: out/install-sentinel terratest/operator_marin3r_3scale_net_v1alpha1/operator_marin3r_3scale_net_discovery_service_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_marin3r_3scale_net_discovery_service_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_marin3r_3scale_net_v1alpha1/operator_marin3r_3scale_net_discovery_service_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_marin3r_3scale_net_envoy_deployment_v1alpha1_manifest_test.go: out/install-sentinel terratest/operator_marin3r_3scale_net_v1alpha1/operator_marin3r_3scale_net_envoy_deployment_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_marin3r_3scale_net_envoy_deployment_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_marin3r_3scale_net_v1alpha1/operator_marin3r_3scale_net_envoy_deployment_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_open_cluster_management_io_cluster_manager_v1_manifest_test.go: out/install-sentinel terratest/operator_open_cluster_management_io_v1/operator_open_cluster_management_io_cluster_manager_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_open_cluster_management_io_cluster_manager_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_open_cluster_management_io_v1/operator_open_cluster_management_io_cluster_manager_v1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_open_cluster_management_io_klusterlet_v1_manifest_test.go: out/install-sentinel terratest/operator_open_cluster_management_io_v1/operator_open_cluster_management_io_klusterlet_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_open_cluster_management_io_klusterlet_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_open_cluster_management_io_v1/operator_open_cluster_management_io_klusterlet_v1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_shipwright_io_shipwright_build_v1alpha1_manifest_test.go: out/install-sentinel terratest/operator_shipwright_io_v1alpha1/operator_shipwright_io_shipwright_build_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_shipwright_io_shipwright_build_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_shipwright_io_v1alpha1/operator_shipwright_io_shipwright_build_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_tigera_io_amazon_cloud_integration_v1_manifest_test.go: out/install-sentinel terratest/operator_tigera_io_v1/operator_tigera_io_amazon_cloud_integration_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_tigera_io_amazon_cloud_integration_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_tigera_io_v1/operator_tigera_io_amazon_cloud_integration_v1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_tigera_io_api_server_v1_manifest_test.go: out/install-sentinel terratest/operator_tigera_io_v1/operator_tigera_io_api_server_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_tigera_io_api_server_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_tigera_io_v1/operator_tigera_io_api_server_v1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_tigera_io_application_layer_v1_manifest_test.go: out/install-sentinel terratest/operator_tigera_io_v1/operator_tigera_io_application_layer_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_tigera_io_application_layer_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_tigera_io_v1/operator_tigera_io_application_layer_v1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_tigera_io_authentication_v1_manifest_test.go: out/install-sentinel terratest/operator_tigera_io_v1/operator_tigera_io_authentication_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_tigera_io_authentication_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_tigera_io_v1/operator_tigera_io_authentication_v1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_tigera_io_compliance_v1_manifest_test.go: out/install-sentinel terratest/operator_tigera_io_v1/operator_tigera_io_compliance_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_tigera_io_compliance_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_tigera_io_v1/operator_tigera_io_compliance_v1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_tigera_io_egress_gateway_v1_manifest_test.go: out/install-sentinel terratest/operator_tigera_io_v1/operator_tigera_io_egress_gateway_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_tigera_io_egress_gateway_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_tigera_io_v1/operator_tigera_io_egress_gateway_v1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_tigera_io_image_set_v1_manifest_test.go: out/install-sentinel terratest/operator_tigera_io_v1/operator_tigera_io_image_set_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_tigera_io_image_set_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_tigera_io_v1/operator_tigera_io_image_set_v1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_tigera_io_installation_v1_manifest_test.go: out/install-sentinel terratest/operator_tigera_io_v1/operator_tigera_io_installation_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_tigera_io_installation_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_tigera_io_v1/operator_tigera_io_installation_v1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_tigera_io_intrusion_detection_v1_manifest_test.go: out/install-sentinel terratest/operator_tigera_io_v1/operator_tigera_io_intrusion_detection_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_tigera_io_intrusion_detection_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_tigera_io_v1/operator_tigera_io_intrusion_detection_v1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_tigera_io_log_collector_v1_manifest_test.go: out/install-sentinel terratest/operator_tigera_io_v1/operator_tigera_io_log_collector_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_tigera_io_log_collector_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_tigera_io_v1/operator_tigera_io_log_collector_v1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_tigera_io_log_storage_v1_manifest_test.go: out/install-sentinel terratest/operator_tigera_io_v1/operator_tigera_io_log_storage_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_tigera_io_log_storage_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_tigera_io_v1/operator_tigera_io_log_storage_v1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_tigera_io_management_cluster_connection_v1_manifest_test.go: out/install-sentinel terratest/operator_tigera_io_v1/operator_tigera_io_management_cluster_connection_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_tigera_io_management_cluster_connection_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_tigera_io_v1/operator_tigera_io_management_cluster_connection_v1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_tigera_io_management_cluster_v1_manifest_test.go: out/install-sentinel terratest/operator_tigera_io_v1/operator_tigera_io_management_cluster_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_tigera_io_management_cluster_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_tigera_io_v1/operator_tigera_io_management_cluster_v1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_tigera_io_manager_v1_manifest_test.go: out/install-sentinel terratest/operator_tigera_io_v1/operator_tigera_io_manager_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_tigera_io_manager_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_tigera_io_v1/operator_tigera_io_manager_v1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_tigera_io_monitor_v1_manifest_test.go: out/install-sentinel terratest/operator_tigera_io_v1/operator_tigera_io_monitor_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_tigera_io_monitor_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_tigera_io_v1/operator_tigera_io_monitor_v1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_tigera_io_packet_capture_v1_manifest_test.go: out/install-sentinel terratest/operator_tigera_io_v1/operator_tigera_io_packet_capture_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_tigera_io_packet_capture_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_tigera_io_v1/operator_tigera_io_packet_capture_v1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_tigera_io_policy_recommendation_v1_manifest_test.go: out/install-sentinel terratest/operator_tigera_io_v1/operator_tigera_io_policy_recommendation_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_tigera_io_policy_recommendation_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_tigera_io_v1/operator_tigera_io_policy_recommendation_v1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_tigera_io_tenant_v1_manifest_test.go: out/install-sentinel terratest/operator_tigera_io_v1/operator_tigera_io_tenant_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_tigera_io_tenant_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_tigera_io_v1/operator_tigera_io_tenant_v1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_tigera_io_tigera_status_v1_manifest_test.go: out/install-sentinel terratest/operator_tigera_io_v1/operator_tigera_io_tigera_status_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_tigera_io_tigera_status_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_tigera_io_v1/operator_tigera_io_tigera_status_v1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_tigera_io_tls_pass_through_route_v1_manifest_test.go: out/install-sentinel terratest/operator_tigera_io_v1/operator_tigera_io_tls_pass_through_route_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_tigera_io_tls_pass_through_route_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_tigera_io_v1/operator_tigera_io_tls_pass_through_route_v1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_tigera_io_tls_terminated_route_v1_manifest_test.go: out/install-sentinel terratest/operator_tigera_io_v1/operator_tigera_io_tls_terminated_route_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_tigera_io_tls_terminated_route_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_tigera_io_v1/operator_tigera_io_tls_terminated_route_v1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_tigera_io_amazon_cloud_integration_v1beta1_manifest_test.go: out/install-sentinel terratest/operator_tigera_io_v1beta1/operator_tigera_io_amazon_cloud_integration_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_tigera_io_amazon_cloud_integration_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_tigera_io_v1beta1/operator_tigera_io_amazon_cloud_integration_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_victoriametrics_com_vm_agent_v1beta1_manifest_test.go: out/install-sentinel terratest/operator_victoriametrics_com_v1beta1/operator_victoriametrics_com_vm_agent_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_victoriametrics_com_vm_agent_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_victoriametrics_com_v1beta1/operator_victoriametrics_com_vm_agent_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_victoriametrics_com_vm_alert_v1beta1_manifest_test.go: out/install-sentinel terratest/operator_victoriametrics_com_v1beta1/operator_victoriametrics_com_vm_alert_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_victoriametrics_com_vm_alert_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_victoriametrics_com_v1beta1/operator_victoriametrics_com_vm_alert_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_victoriametrics_com_vm_alertmanager_config_v1beta1_manifest_test.go: out/install-sentinel terratest/operator_victoriametrics_com_v1beta1/operator_victoriametrics_com_vm_alertmanager_config_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_victoriametrics_com_vm_alertmanager_config_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_victoriametrics_com_v1beta1/operator_victoriametrics_com_vm_alertmanager_config_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_victoriametrics_com_vm_alertmanager_v1beta1_manifest_test.go: out/install-sentinel terratest/operator_victoriametrics_com_v1beta1/operator_victoriametrics_com_vm_alertmanager_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_victoriametrics_com_vm_alertmanager_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_victoriametrics_com_v1beta1/operator_victoriametrics_com_vm_alertmanager_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_victoriametrics_com_vm_auth_v1beta1_manifest_test.go: out/install-sentinel terratest/operator_victoriametrics_com_v1beta1/operator_victoriametrics_com_vm_auth_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_victoriametrics_com_vm_auth_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_victoriametrics_com_v1beta1/operator_victoriametrics_com_vm_auth_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_victoriametrics_com_vm_cluster_v1beta1_manifest_test.go: out/install-sentinel terratest/operator_victoriametrics_com_v1beta1/operator_victoriametrics_com_vm_cluster_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_victoriametrics_com_vm_cluster_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_victoriametrics_com_v1beta1/operator_victoriametrics_com_vm_cluster_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_victoriametrics_com_vm_node_scrape_v1beta1_manifest_test.go: out/install-sentinel terratest/operator_victoriametrics_com_v1beta1/operator_victoriametrics_com_vm_node_scrape_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_victoriametrics_com_vm_node_scrape_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_victoriametrics_com_v1beta1/operator_victoriametrics_com_vm_node_scrape_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_victoriametrics_com_vm_pod_scrape_v1beta1_manifest_test.go: out/install-sentinel terratest/operator_victoriametrics_com_v1beta1/operator_victoriametrics_com_vm_pod_scrape_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_victoriametrics_com_vm_pod_scrape_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_victoriametrics_com_v1beta1/operator_victoriametrics_com_vm_pod_scrape_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_victoriametrics_com_vm_probe_v1beta1_manifest_test.go: out/install-sentinel terratest/operator_victoriametrics_com_v1beta1/operator_victoriametrics_com_vm_probe_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_victoriametrics_com_vm_probe_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_victoriametrics_com_v1beta1/operator_victoriametrics_com_vm_probe_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_victoriametrics_com_vm_rule_v1beta1_manifest_test.go: out/install-sentinel terratest/operator_victoriametrics_com_v1beta1/operator_victoriametrics_com_vm_rule_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_victoriametrics_com_vm_rule_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_victoriametrics_com_v1beta1/operator_victoriametrics_com_vm_rule_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_victoriametrics_com_vm_service_scrape_v1beta1_manifest_test.go: out/install-sentinel terratest/operator_victoriametrics_com_v1beta1/operator_victoriametrics_com_vm_service_scrape_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_victoriametrics_com_vm_service_scrape_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_victoriametrics_com_v1beta1/operator_victoriametrics_com_vm_service_scrape_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_victoriametrics_com_vm_single_v1beta1_manifest_test.go: out/install-sentinel terratest/operator_victoriametrics_com_v1beta1/operator_victoriametrics_com_vm_single_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_victoriametrics_com_vm_single_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_victoriametrics_com_v1beta1/operator_victoriametrics_com_vm_single_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_victoriametrics_com_vm_static_scrape_v1beta1_manifest_test.go: out/install-sentinel terratest/operator_victoriametrics_com_v1beta1/operator_victoriametrics_com_vm_static_scrape_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_victoriametrics_com_vm_static_scrape_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_victoriametrics_com_v1beta1/operator_victoriametrics_com_vm_static_scrape_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-operator_victoriametrics_com_vm_user_v1beta1_manifest_test.go: out/install-sentinel terratest/operator_victoriametrics_com_v1beta1/operator_victoriametrics_com_vm_user_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_operator_victoriametrics_com_vm_user_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/operator_victoriametrics_com_v1beta1/operator_victoriametrics_com_vm_user_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-oracle_db_anthosapis_com_backup_schedule_v1alpha1_manifest_test.go: out/install-sentinel terratest/oracle_db_anthosapis_com_v1alpha1/oracle_db_anthosapis_com_backup_schedule_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_oracle_db_anthosapis_com_backup_schedule_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/oracle_db_anthosapis_com_v1alpha1/oracle_db_anthosapis_com_backup_schedule_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-oracle_db_anthosapis_com_backup_v1alpha1_manifest_test.go: out/install-sentinel terratest/oracle_db_anthosapis_com_v1alpha1/oracle_db_anthosapis_com_backup_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_oracle_db_anthosapis_com_backup_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/oracle_db_anthosapis_com_v1alpha1/oracle_db_anthosapis_com_backup_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-oracle_db_anthosapis_com_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/oracle_db_anthosapis_com_v1alpha1/oracle_db_anthosapis_com_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_oracle_db_anthosapis_com_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/oracle_db_anthosapis_com_v1alpha1/oracle_db_anthosapis_com_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-oracle_db_anthosapis_com_cron_anything_v1alpha1_manifest_test.go: out/install-sentinel terratest/oracle_db_anthosapis_com_v1alpha1/oracle_db_anthosapis_com_cron_anything_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_oracle_db_anthosapis_com_cron_anything_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/oracle_db_anthosapis_com_v1alpha1/oracle_db_anthosapis_com_cron_anything_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-oracle_db_anthosapis_com_database_v1alpha1_manifest_test.go: out/install-sentinel terratest/oracle_db_anthosapis_com_v1alpha1/oracle_db_anthosapis_com_database_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_oracle_db_anthosapis_com_database_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/oracle_db_anthosapis_com_v1alpha1/oracle_db_anthosapis_com_database_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-oracle_db_anthosapis_com_export_v1alpha1_manifest_test.go: out/install-sentinel terratest/oracle_db_anthosapis_com_v1alpha1/oracle_db_anthosapis_com_export_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_oracle_db_anthosapis_com_export_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/oracle_db_anthosapis_com_v1alpha1/oracle_db_anthosapis_com_export_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-oracle_db_anthosapis_com_import_v1alpha1_manifest_test.go: out/install-sentinel terratest/oracle_db_anthosapis_com_v1alpha1/oracle_db_anthosapis_com_import_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_oracle_db_anthosapis_com_import_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/oracle_db_anthosapis_com_v1alpha1/oracle_db_anthosapis_com_import_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-oracle_db_anthosapis_com_instance_v1alpha1_manifest_test.go: out/install-sentinel terratest/oracle_db_anthosapis_com_v1alpha1/oracle_db_anthosapis_com_instance_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_oracle_db_anthosapis_com_instance_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/oracle_db_anthosapis_com_v1alpha1/oracle_db_anthosapis_com_instance_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-oracle_db_anthosapis_com_pitr_v1alpha1_manifest_test.go: out/install-sentinel terratest/oracle_db_anthosapis_com_v1alpha1/oracle_db_anthosapis_com_pitr_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_oracle_db_anthosapis_com_pitr_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/oracle_db_anthosapis_com_v1alpha1/oracle_db_anthosapis_com_pitr_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-oracle_db_anthosapis_com_release_v1alpha1_manifest_test.go: out/install-sentinel terratest/oracle_db_anthosapis_com_v1alpha1/oracle_db_anthosapis_com_release_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_oracle_db_anthosapis_com_release_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/oracle_db_anthosapis_com_v1alpha1/oracle_db_anthosapis_com_release_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-org_eclipse_che_che_cluster_v1_manifest_test.go: out/install-sentinel terratest/org_eclipse_che_v1/org_eclipse_che_che_cluster_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_org_eclipse_che_che_cluster_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/org_eclipse_che_v1/org_eclipse_che_che_cluster_v1_manifest_test.go
	touch $@
out/terratest-sentinel-org_eclipse_che_che_cluster_v2_manifest_test.go: out/install-sentinel terratest/org_eclipse_che_v2/org_eclipse_che_che_cluster_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_org_eclipse_che_che_cluster_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/org_eclipse_che_v2/org_eclipse_che_che_cluster_v2_manifest_test.go
	touch $@
out/terratest-sentinel-organizations_services_k8s_aws_organizational_unit_v1alpha1_manifest_test.go: out/install-sentinel terratest/organizations_services_k8s_aws_v1alpha1/organizations_services_k8s_aws_organizational_unit_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_organizations_services_k8s_aws_organizational_unit_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/organizations_services_k8s_aws_v1alpha1/organizations_services_k8s_aws_organizational_unit_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-pgv2_percona_com_percona_pg_backup_v2_manifest_test.go: out/install-sentinel terratest/pgv2_percona_com_v2/pgv2_percona_com_percona_pg_backup_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_pgv2_percona_com_percona_pg_backup_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/pgv2_percona_com_v2/pgv2_percona_com_percona_pg_backup_v2_manifest_test.go
	touch $@
out/terratest-sentinel-pgv2_percona_com_percona_pg_cluster_v2_manifest_test.go: out/install-sentinel terratest/pgv2_percona_com_v2/pgv2_percona_com_percona_pg_cluster_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_pgv2_percona_com_percona_pg_cluster_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/pgv2_percona_com_v2/pgv2_percona_com_percona_pg_cluster_v2_manifest_test.go
	touch $@
out/terratest-sentinel-pgv2_percona_com_percona_pg_restore_v2_manifest_test.go: out/install-sentinel terratest/pgv2_percona_com_v2/pgv2_percona_com_percona_pg_restore_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_pgv2_percona_com_percona_pg_restore_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/pgv2_percona_com_v2/pgv2_percona_com_percona_pg_restore_v2_manifest_test.go
	touch $@
out/terratest-sentinel-pgv2_percona_com_percona_pg_upgrade_v2_manifest_test.go: out/install-sentinel terratest/pgv2_percona_com_v2/pgv2_percona_com_percona_pg_upgrade_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_pgv2_percona_com_percona_pg_upgrade_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/pgv2_percona_com_v2/pgv2_percona_com_percona_pg_upgrade_v2_manifest_test.go
	touch $@
out/terratest-sentinel-pipes_services_k8s_aws_pipe_v1alpha1_manifest_test.go: out/install-sentinel terratest/pipes_services_k8s_aws_v1alpha1/pipes_services_k8s_aws_pipe_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_pipes_services_k8s_aws_pipe_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/pipes_services_k8s_aws_v1alpha1/pipes_services_k8s_aws_pipe_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-pkg_crossplane_io_configuration_revision_v1_manifest_test.go: out/install-sentinel terratest/pkg_crossplane_io_v1/pkg_crossplane_io_configuration_revision_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_pkg_crossplane_io_configuration_revision_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/pkg_crossplane_io_v1/pkg_crossplane_io_configuration_revision_v1_manifest_test.go
	touch $@
out/terratest-sentinel-pkg_crossplane_io_configuration_v1_manifest_test.go: out/install-sentinel terratest/pkg_crossplane_io_v1/pkg_crossplane_io_configuration_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_pkg_crossplane_io_configuration_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/pkg_crossplane_io_v1/pkg_crossplane_io_configuration_v1_manifest_test.go
	touch $@
out/terratest-sentinel-pkg_crossplane_io_provider_revision_v1_manifest_test.go: out/install-sentinel terratest/pkg_crossplane_io_v1/pkg_crossplane_io_provider_revision_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_pkg_crossplane_io_provider_revision_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/pkg_crossplane_io_v1/pkg_crossplane_io_provider_revision_v1_manifest_test.go
	touch $@
out/terratest-sentinel-pkg_crossplane_io_provider_v1_manifest_test.go: out/install-sentinel terratest/pkg_crossplane_io_v1/pkg_crossplane_io_provider_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_pkg_crossplane_io_provider_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/pkg_crossplane_io_v1/pkg_crossplane_io_provider_v1_manifest_test.go
	touch $@
out/terratest-sentinel-pkg_crossplane_io_controller_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/pkg_crossplane_io_v1alpha1/pkg_crossplane_io_controller_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_pkg_crossplane_io_controller_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/pkg_crossplane_io_v1alpha1/pkg_crossplane_io_controller_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-pkg_crossplane_io_lock_v1beta1_manifest_test.go: out/install-sentinel terratest/pkg_crossplane_io_v1beta1/pkg_crossplane_io_lock_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_pkg_crossplane_io_lock_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/pkg_crossplane_io_v1beta1/pkg_crossplane_io_lock_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-policy_clusterpedia_io_cluster_import_policy_v1alpha1_manifest_test.go: out/install-sentinel terratest/policy_clusterpedia_io_v1alpha1/policy_clusterpedia_io_cluster_import_policy_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_policy_clusterpedia_io_cluster_import_policy_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/policy_clusterpedia_io_v1alpha1/policy_clusterpedia_io_cluster_import_policy_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-policy_clusterpedia_io_pedia_cluster_lifecycle_v1alpha1_manifest_test.go: out/install-sentinel terratest/policy_clusterpedia_io_v1alpha1/policy_clusterpedia_io_pedia_cluster_lifecycle_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_policy_clusterpedia_io_pedia_cluster_lifecycle_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/policy_clusterpedia_io_v1alpha1/policy_clusterpedia_io_pedia_cluster_lifecycle_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-policy_karmada_io_cluster_override_policy_v1alpha1_manifest_test.go: out/install-sentinel terratest/policy_karmada_io_v1alpha1/policy_karmada_io_cluster_override_policy_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_policy_karmada_io_cluster_override_policy_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/policy_karmada_io_v1alpha1/policy_karmada_io_cluster_override_policy_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-policy_karmada_io_cluster_propagation_policy_v1alpha1_manifest_test.go: out/install-sentinel terratest/policy_karmada_io_v1alpha1/policy_karmada_io_cluster_propagation_policy_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_policy_karmada_io_cluster_propagation_policy_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/policy_karmada_io_v1alpha1/policy_karmada_io_cluster_propagation_policy_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-policy_karmada_io_federated_resource_quota_v1alpha1_manifest_test.go: out/install-sentinel terratest/policy_karmada_io_v1alpha1/policy_karmada_io_federated_resource_quota_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_policy_karmada_io_federated_resource_quota_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/policy_karmada_io_v1alpha1/policy_karmada_io_federated_resource_quota_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-policy_karmada_io_override_policy_v1alpha1_manifest_test.go: out/install-sentinel terratest/policy_karmada_io_v1alpha1/policy_karmada_io_override_policy_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_policy_karmada_io_override_policy_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/policy_karmada_io_v1alpha1/policy_karmada_io_override_policy_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-policy_karmada_io_propagation_policy_v1alpha1_manifest_test.go: out/install-sentinel terratest/policy_karmada_io_v1alpha1/policy_karmada_io_propagation_policy_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_policy_karmada_io_propagation_policy_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/policy_karmada_io_v1alpha1/policy_karmada_io_propagation_policy_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-policy_kubeedge_io_service_account_access_v1alpha1_manifest_test.go: out/install-sentinel terratest/policy_kubeedge_io_v1alpha1/policy_kubeedge_io_service_account_access_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_policy_kubeedge_io_service_account_access_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/policy_kubeedge_io_v1alpha1/policy_kubeedge_io_service_account_access_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-policy_networking_k8s_io_admin_network_policy_v1alpha1_manifest_test.go: out/install-sentinel terratest/policy_networking_k8s_io_v1alpha1/policy_networking_k8s_io_admin_network_policy_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_policy_networking_k8s_io_admin_network_policy_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/policy_networking_k8s_io_v1alpha1/policy_networking_k8s_io_admin_network_policy_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-policy_networking_k8s_io_baseline_admin_network_policy_v1alpha1_manifest_test.go: out/install-sentinel terratest/policy_networking_k8s_io_v1alpha1/policy_networking_k8s_io_baseline_admin_network_policy_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_policy_networking_k8s_io_baseline_admin_network_policy_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/policy_networking_k8s_io_v1alpha1/policy_networking_k8s_io_baseline_admin_network_policy_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-policy_pod_disruption_budget_v1_manifest_test.go: out/install-sentinel terratest/policy_v1/policy_pod_disruption_budget_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_policy_pod_disruption_budget_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/policy_v1/policy_pod_disruption_budget_v1_manifest_test.go
	touch $@
out/terratest-sentinel-postgres_operator_crunchydata_com_pg_admin_v1beta1_manifest_test.go: out/install-sentinel terratest/postgres_operator_crunchydata_com_v1beta1/postgres_operator_crunchydata_com_pg_admin_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_postgres_operator_crunchydata_com_pg_admin_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/postgres_operator_crunchydata_com_v1beta1/postgres_operator_crunchydata_com_pg_admin_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-postgres_operator_crunchydata_com_pg_upgrade_v1beta1_manifest_test.go: out/install-sentinel terratest/postgres_operator_crunchydata_com_v1beta1/postgres_operator_crunchydata_com_pg_upgrade_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_postgres_operator_crunchydata_com_pg_upgrade_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/postgres_operator_crunchydata_com_v1beta1/postgres_operator_crunchydata_com_pg_upgrade_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-postgres_operator_crunchydata_com_postgres_cluster_v1beta1_manifest_test.go: out/install-sentinel terratest/postgres_operator_crunchydata_com_v1beta1/postgres_operator_crunchydata_com_postgres_cluster_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_postgres_operator_crunchydata_com_postgres_cluster_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/postgres_operator_crunchydata_com_v1beta1/postgres_operator_crunchydata_com_postgres_cluster_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-postgresql_cnpg_io_backup_v1_manifest_test.go: out/install-sentinel terratest/postgresql_cnpg_io_v1/postgresql_cnpg_io_backup_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_postgresql_cnpg_io_backup_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/postgresql_cnpg_io_v1/postgresql_cnpg_io_backup_v1_manifest_test.go
	touch $@
out/terratest-sentinel-postgresql_cnpg_io_cluster_v1_manifest_test.go: out/install-sentinel terratest/postgresql_cnpg_io_v1/postgresql_cnpg_io_cluster_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_postgresql_cnpg_io_cluster_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/postgresql_cnpg_io_v1/postgresql_cnpg_io_cluster_v1_manifest_test.go
	touch $@
out/terratest-sentinel-postgresql_cnpg_io_pooler_v1_manifest_test.go: out/install-sentinel terratest/postgresql_cnpg_io_v1/postgresql_cnpg_io_pooler_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_postgresql_cnpg_io_pooler_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/postgresql_cnpg_io_v1/postgresql_cnpg_io_pooler_v1_manifest_test.go
	touch $@
out/terratest-sentinel-postgresql_cnpg_io_scheduled_backup_v1_manifest_test.go: out/install-sentinel terratest/postgresql_cnpg_io_v1/postgresql_cnpg_io_scheduled_backup_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_postgresql_cnpg_io_scheduled_backup_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/postgresql_cnpg_io_v1/postgresql_cnpg_io_scheduled_backup_v1_manifest_test.go
	touch $@
out/terratest-sentinel-projectcontour_io_http_proxy_v1_manifest_test.go: out/install-sentinel terratest/projectcontour_io_v1/projectcontour_io_http_proxy_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_projectcontour_io_http_proxy_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/projectcontour_io_v1/projectcontour_io_http_proxy_v1_manifest_test.go
	touch $@
out/terratest-sentinel-projectcontour_io_tls_certificate_delegation_v1_manifest_test.go: out/install-sentinel terratest/projectcontour_io_v1/projectcontour_io_tls_certificate_delegation_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_projectcontour_io_tls_certificate_delegation_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/projectcontour_io_v1/projectcontour_io_tls_certificate_delegation_v1_manifest_test.go
	touch $@
out/terratest-sentinel-projectcontour_io_contour_configuration_v1alpha1_manifest_test.go: out/install-sentinel terratest/projectcontour_io_v1alpha1/projectcontour_io_contour_configuration_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_projectcontour_io_contour_configuration_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/projectcontour_io_v1alpha1/projectcontour_io_contour_configuration_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-projectcontour_io_contour_deployment_v1alpha1_manifest_test.go: out/install-sentinel terratest/projectcontour_io_v1alpha1/projectcontour_io_contour_deployment_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_projectcontour_io_contour_deployment_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/projectcontour_io_v1alpha1/projectcontour_io_contour_deployment_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-projectcontour_io_extension_service_v1alpha1_manifest_test.go: out/install-sentinel terratest/projectcontour_io_v1alpha1/projectcontour_io_extension_service_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_projectcontour_io_extension_service_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/projectcontour_io_v1alpha1/projectcontour_io_extension_service_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-prometheusservice_services_k8s_aws_alert_manager_definition_v1alpha1_manifest_test.go: out/install-sentinel terratest/prometheusservice_services_k8s_aws_v1alpha1/prometheusservice_services_k8s_aws_alert_manager_definition_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_prometheusservice_services_k8s_aws_alert_manager_definition_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/prometheusservice_services_k8s_aws_v1alpha1/prometheusservice_services_k8s_aws_alert_manager_definition_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-prometheusservice_services_k8s_aws_logging_configuration_v1alpha1_manifest_test.go: out/install-sentinel terratest/prometheusservice_services_k8s_aws_v1alpha1/prometheusservice_services_k8s_aws_logging_configuration_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_prometheusservice_services_k8s_aws_logging_configuration_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/prometheusservice_services_k8s_aws_v1alpha1/prometheusservice_services_k8s_aws_logging_configuration_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-prometheusservice_services_k8s_aws_rule_groups_namespace_v1alpha1_manifest_test.go: out/install-sentinel terratest/prometheusservice_services_k8s_aws_v1alpha1/prometheusservice_services_k8s_aws_rule_groups_namespace_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_prometheusservice_services_k8s_aws_rule_groups_namespace_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/prometheusservice_services_k8s_aws_v1alpha1/prometheusservice_services_k8s_aws_rule_groups_namespace_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-prometheusservice_services_k8s_aws_workspace_v1alpha1_manifest_test.go: out/install-sentinel terratest/prometheusservice_services_k8s_aws_v1alpha1/prometheusservice_services_k8s_aws_workspace_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_prometheusservice_services_k8s_aws_workspace_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/prometheusservice_services_k8s_aws_v1alpha1/prometheusservice_services_k8s_aws_workspace_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-ps_percona_com_percona_server_my_sql_v1alpha1_manifest_test.go: out/install-sentinel terratest/ps_percona_com_v1alpha1/ps_percona_com_percona_server_my_sql_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_ps_percona_com_percona_server_my_sql_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ps_percona_com_v1alpha1/ps_percona_com_percona_server_my_sql_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-ps_percona_com_percona_server_my_sql_backup_v1alpha1_manifest_test.go: out/install-sentinel terratest/ps_percona_com_v1alpha1/ps_percona_com_percona_server_my_sql_backup_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_ps_percona_com_percona_server_my_sql_backup_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ps_percona_com_v1alpha1/ps_percona_com_percona_server_my_sql_backup_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-ps_percona_com_percona_server_my_sql_restore_v1alpha1_manifest_test.go: out/install-sentinel terratest/ps_percona_com_v1alpha1/ps_percona_com_percona_server_my_sql_restore_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_ps_percona_com_percona_server_my_sql_restore_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ps_percona_com_v1alpha1/ps_percona_com_percona_server_my_sql_restore_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-psmdb_percona_com_percona_server_mongo_db_v1_manifest_test.go: out/install-sentinel terratest/psmdb_percona_com_v1/psmdb_percona_com_percona_server_mongo_db_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_psmdb_percona_com_percona_server_mongo_db_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/psmdb_percona_com_v1/psmdb_percona_com_percona_server_mongo_db_v1_manifest_test.go
	touch $@
out/terratest-sentinel-psmdb_percona_com_percona_server_mongo_db_backup_v1_manifest_test.go: out/install-sentinel terratest/psmdb_percona_com_v1/psmdb_percona_com_percona_server_mongo_db_backup_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_psmdb_percona_com_percona_server_mongo_db_backup_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/psmdb_percona_com_v1/psmdb_percona_com_percona_server_mongo_db_backup_v1_manifest_test.go
	touch $@
out/terratest-sentinel-psmdb_percona_com_percona_server_mongo_db_restore_v1_manifest_test.go: out/install-sentinel terratest/psmdb_percona_com_v1/psmdb_percona_com_percona_server_mongo_db_restore_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_psmdb_percona_com_percona_server_mongo_db_restore_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/psmdb_percona_com_v1/psmdb_percona_com_percona_server_mongo_db_restore_v1_manifest_test.go
	touch $@
out/terratest-sentinel-ptp_openshift_io_node_ptp_device_v1_manifest_test.go: out/install-sentinel terratest/ptp_openshift_io_v1/ptp_openshift_io_node_ptp_device_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_ptp_openshift_io_node_ptp_device_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ptp_openshift_io_v1/ptp_openshift_io_node_ptp_device_v1_manifest_test.go
	touch $@
out/terratest-sentinel-ptp_openshift_io_ptp_config_v1_manifest_test.go: out/install-sentinel terratest/ptp_openshift_io_v1/ptp_openshift_io_ptp_config_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_ptp_openshift_io_ptp_config_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ptp_openshift_io_v1/ptp_openshift_io_ptp_config_v1_manifest_test.go
	touch $@
out/terratest-sentinel-ptp_openshift_io_ptp_operator_config_v1_manifest_test.go: out/install-sentinel terratest/ptp_openshift_io_v1/ptp_openshift_io_ptp_operator_config_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_ptp_openshift_io_ptp_operator_config_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ptp_openshift_io_v1/ptp_openshift_io_ptp_operator_config_v1_manifest_test.go
	touch $@
out/terratest-sentinel-pubsubplus_solace_com_pub_sub_plus_event_broker_v1beta1_manifest_test.go: out/install-sentinel terratest/pubsubplus_solace_com_v1beta1/pubsubplus_solace_com_pub_sub_plus_event_broker_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_pubsubplus_solace_com_pub_sub_plus_event_broker_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/pubsubplus_solace_com_v1beta1/pubsubplus_solace_com_pub_sub_plus_event_broker_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-pxc_percona_com_percona_xtra_db_cluster_backup_v1_manifest_test.go: out/install-sentinel terratest/pxc_percona_com_v1/pxc_percona_com_percona_xtra_db_cluster_backup_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_pxc_percona_com_percona_xtra_db_cluster_backup_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/pxc_percona_com_v1/pxc_percona_com_percona_xtra_db_cluster_backup_v1_manifest_test.go
	touch $@
out/terratest-sentinel-pxc_percona_com_percona_xtra_db_cluster_restore_v1_manifest_test.go: out/install-sentinel terratest/pxc_percona_com_v1/pxc_percona_com_percona_xtra_db_cluster_restore_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_pxc_percona_com_percona_xtra_db_cluster_restore_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/pxc_percona_com_v1/pxc_percona_com_percona_xtra_db_cluster_restore_v1_manifest_test.go
	touch $@
out/terratest-sentinel-pxc_percona_com_percona_xtra_db_cluster_v1_manifest_test.go: out/install-sentinel terratest/pxc_percona_com_v1/pxc_percona_com_percona_xtra_db_cluster_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_pxc_percona_com_percona_xtra_db_cluster_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/pxc_percona_com_v1/pxc_percona_com_percona_xtra_db_cluster_v1_manifest_test.go
	touch $@
out/terratest-sentinel-quay_redhat_com_quay_registry_v1_manifest_test.go: out/install-sentinel terratest/quay_redhat_com_v1/quay_redhat_com_quay_registry_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_quay_redhat_com_quay_registry_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/quay_redhat_com_v1/quay_redhat_com_quay_registry_v1_manifest_test.go
	touch $@
out/terratest-sentinel-quota_codeflare_dev_quota_subtree_v1alpha1_manifest_test.go: out/install-sentinel terratest/quota_codeflare_dev_v1alpha1/quota_codeflare_dev_quota_subtree_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_quota_codeflare_dev_quota_subtree_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/quota_codeflare_dev_v1alpha1/quota_codeflare_dev_quota_subtree_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-ray_io_ray_cluster_v1_manifest_test.go: out/install-sentinel terratest/ray_io_v1/ray_io_ray_cluster_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_ray_io_ray_cluster_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ray_io_v1/ray_io_ray_cluster_v1_manifest_test.go
	touch $@
out/terratest-sentinel-ray_io_ray_job_v1_manifest_test.go: out/install-sentinel terratest/ray_io_v1/ray_io_ray_job_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_ray_io_ray_job_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ray_io_v1/ray_io_ray_job_v1_manifest_test.go
	touch $@
out/terratest-sentinel-ray_io_ray_service_v1_manifest_test.go: out/install-sentinel terratest/ray_io_v1/ray_io_ray_service_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_ray_io_ray_service_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ray_io_v1/ray_io_ray_service_v1_manifest_test.go
	touch $@
out/terratest-sentinel-ray_io_ray_cluster_v1alpha1_manifest_test.go: out/install-sentinel terratest/ray_io_v1alpha1/ray_io_ray_cluster_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_ray_io_ray_cluster_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ray_io_v1alpha1/ray_io_ray_cluster_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-ray_io_ray_job_v1alpha1_manifest_test.go: out/install-sentinel terratest/ray_io_v1alpha1/ray_io_ray_job_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_ray_io_ray_job_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ray_io_v1alpha1/ray_io_ray_job_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-ray_io_ray_service_v1alpha1_manifest_test.go: out/install-sentinel terratest/ray_io_v1alpha1/ray_io_ray_service_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_ray_io_ray_service_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ray_io_v1alpha1/ray_io_ray_service_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-rbac_authorization_k8s_io_cluster_role_binding_v1_manifest_test.go: out/install-sentinel terratest/rbac_authorization_k8s_io_v1/rbac_authorization_k8s_io_cluster_role_binding_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_rbac_authorization_k8s_io_cluster_role_binding_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/rbac_authorization_k8s_io_v1/rbac_authorization_k8s_io_cluster_role_binding_v1_manifest_test.go
	touch $@
out/terratest-sentinel-rbac_authorization_k8s_io_cluster_role_v1_manifest_test.go: out/install-sentinel terratest/rbac_authorization_k8s_io_v1/rbac_authorization_k8s_io_cluster_role_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_rbac_authorization_k8s_io_cluster_role_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/rbac_authorization_k8s_io_v1/rbac_authorization_k8s_io_cluster_role_v1_manifest_test.go
	touch $@
out/terratest-sentinel-rbac_authorization_k8s_io_role_binding_v1_manifest_test.go: out/install-sentinel terratest/rbac_authorization_k8s_io_v1/rbac_authorization_k8s_io_role_binding_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_rbac_authorization_k8s_io_role_binding_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/rbac_authorization_k8s_io_v1/rbac_authorization_k8s_io_role_binding_v1_manifest_test.go
	touch $@
out/terratest-sentinel-rbac_authorization_k8s_io_role_v1_manifest_test.go: out/install-sentinel terratest/rbac_authorization_k8s_io_v1/rbac_authorization_k8s_io_role_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_rbac_authorization_k8s_io_role_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/rbac_authorization_k8s_io_v1/rbac_authorization_k8s_io_role_v1_manifest_test.go
	touch $@
out/terratest-sentinel-rbacmanager_reactiveops_io_rbac_definition_v1beta1_manifest_test.go: out/install-sentinel terratest/rbacmanager_reactiveops_io_v1beta1/rbacmanager_reactiveops_io_rbac_definition_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_rbacmanager_reactiveops_io_rbac_definition_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/rbacmanager_reactiveops_io_v1beta1/rbacmanager_reactiveops_io_rbac_definition_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-rc_app_stacks_runtime_component_v1_manifest_test.go: out/install-sentinel terratest/rc_app_stacks_v1/rc_app_stacks_runtime_component_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_rc_app_stacks_runtime_component_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/rc_app_stacks_v1/rc_app_stacks_runtime_component_v1_manifest_test.go
	touch $@
out/terratest-sentinel-rc_app_stacks_runtime_operation_v1_manifest_test.go: out/install-sentinel terratest/rc_app_stacks_v1/rc_app_stacks_runtime_operation_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_rc_app_stacks_runtime_operation_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/rc_app_stacks_v1/rc_app_stacks_runtime_operation_v1_manifest_test.go
	touch $@
out/terratest-sentinel-rc_app_stacks_runtime_component_v1beta2_manifest_test.go: out/install-sentinel terratest/rc_app_stacks_v1beta2/rc_app_stacks_runtime_component_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_rc_app_stacks_runtime_component_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/rc_app_stacks_v1beta2/rc_app_stacks_runtime_component_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-rc_app_stacks_runtime_operation_v1beta2_manifest_test.go: out/install-sentinel terratest/rc_app_stacks_v1beta2/rc_app_stacks_runtime_operation_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_rc_app_stacks_runtime_operation_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/rc_app_stacks_v1beta2/rc_app_stacks_runtime_operation_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-rds_services_k8s_aws_db_cluster_parameter_group_v1alpha1_manifest_test.go: out/install-sentinel terratest/rds_services_k8s_aws_v1alpha1/rds_services_k8s_aws_db_cluster_parameter_group_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_rds_services_k8s_aws_db_cluster_parameter_group_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/rds_services_k8s_aws_v1alpha1/rds_services_k8s_aws_db_cluster_parameter_group_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-rds_services_k8s_aws_db_cluster_v1alpha1_manifest_test.go: out/install-sentinel terratest/rds_services_k8s_aws_v1alpha1/rds_services_k8s_aws_db_cluster_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_rds_services_k8s_aws_db_cluster_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/rds_services_k8s_aws_v1alpha1/rds_services_k8s_aws_db_cluster_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-rds_services_k8s_aws_db_instance_v1alpha1_manifest_test.go: out/install-sentinel terratest/rds_services_k8s_aws_v1alpha1/rds_services_k8s_aws_db_instance_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_rds_services_k8s_aws_db_instance_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/rds_services_k8s_aws_v1alpha1/rds_services_k8s_aws_db_instance_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-rds_services_k8s_aws_db_parameter_group_v1alpha1_manifest_test.go: out/install-sentinel terratest/rds_services_k8s_aws_v1alpha1/rds_services_k8s_aws_db_parameter_group_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_rds_services_k8s_aws_db_parameter_group_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/rds_services_k8s_aws_v1alpha1/rds_services_k8s_aws_db_parameter_group_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-rds_services_k8s_aws_db_proxy_v1alpha1_manifest_test.go: out/install-sentinel terratest/rds_services_k8s_aws_v1alpha1/rds_services_k8s_aws_db_proxy_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_rds_services_k8s_aws_db_proxy_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/rds_services_k8s_aws_v1alpha1/rds_services_k8s_aws_db_proxy_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-rds_services_k8s_aws_db_subnet_group_v1alpha1_manifest_test.go: out/install-sentinel terratest/rds_services_k8s_aws_v1alpha1/rds_services_k8s_aws_db_subnet_group_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_rds_services_k8s_aws_db_subnet_group_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/rds_services_k8s_aws_v1alpha1/rds_services_k8s_aws_db_subnet_group_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-rds_services_k8s_aws_global_cluster_v1alpha1_manifest_test.go: out/install-sentinel terratest/rds_services_k8s_aws_v1alpha1/rds_services_k8s_aws_global_cluster_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_rds_services_k8s_aws_global_cluster_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/rds_services_k8s_aws_v1alpha1/rds_services_k8s_aws_global_cluster_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-redhatcop_redhat_io_group_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/redhatcop_redhat_io_v1alpha1/redhatcop_redhat_io_group_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_redhatcop_redhat_io_group_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/redhatcop_redhat_io_v1alpha1/redhatcop_redhat_io_group_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-redhatcop_redhat_io_keepalived_group_v1alpha1_manifest_test.go: out/install-sentinel terratest/redhatcop_redhat_io_v1alpha1/redhatcop_redhat_io_keepalived_group_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_redhatcop_redhat_io_keepalived_group_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/redhatcop_redhat_io_v1alpha1/redhatcop_redhat_io_keepalived_group_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-redhatcop_redhat_io_namespace_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/redhatcop_redhat_io_v1alpha1/redhatcop_redhat_io_namespace_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_redhatcop_redhat_io_namespace_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/redhatcop_redhat_io_v1alpha1/redhatcop_redhat_io_namespace_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-redhatcop_redhat_io_patch_v1alpha1_manifest_test.go: out/install-sentinel terratest/redhatcop_redhat_io_v1alpha1/redhatcop_redhat_io_patch_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_redhatcop_redhat_io_patch_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/redhatcop_redhat_io_v1alpha1/redhatcop_redhat_io_patch_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-redhatcop_redhat_io_user_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/redhatcop_redhat_io_v1alpha1/redhatcop_redhat_io_user_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_redhatcop_redhat_io_user_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/redhatcop_redhat_io_v1alpha1/redhatcop_redhat_io_user_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-registry_apicur_io_apicurio_registry_v1_manifest_test.go: out/install-sentinel terratest/registry_apicur_io_v1/registry_apicur_io_apicurio_registry_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_registry_apicur_io_apicurio_registry_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/registry_apicur_io_v1/registry_apicur_io_apicurio_registry_v1_manifest_test.go
	touch $@
out/terratest-sentinel-registry_devfile_io_cluster_devfile_registries_list_v1alpha1_manifest_test.go: out/install-sentinel terratest/registry_devfile_io_v1alpha1/registry_devfile_io_cluster_devfile_registries_list_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_registry_devfile_io_cluster_devfile_registries_list_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/registry_devfile_io_v1alpha1/registry_devfile_io_cluster_devfile_registries_list_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-registry_devfile_io_devfile_registries_list_v1alpha1_manifest_test.go: out/install-sentinel terratest/registry_devfile_io_v1alpha1/registry_devfile_io_devfile_registries_list_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_registry_devfile_io_devfile_registries_list_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/registry_devfile_io_v1alpha1/registry_devfile_io_devfile_registries_list_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-registry_devfile_io_devfile_registry_v1alpha1_manifest_test.go: out/install-sentinel terratest/registry_devfile_io_v1alpha1/registry_devfile_io_devfile_registry_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_registry_devfile_io_devfile_registry_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/registry_devfile_io_v1alpha1/registry_devfile_io_devfile_registry_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-reliablesyncs_kubeedge_io_cluster_object_sync_v1alpha1_manifest_test.go: out/install-sentinel terratest/reliablesyncs_kubeedge_io_v1alpha1/reliablesyncs_kubeedge_io_cluster_object_sync_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_reliablesyncs_kubeedge_io_cluster_object_sync_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/reliablesyncs_kubeedge_io_v1alpha1/reliablesyncs_kubeedge_io_cluster_object_sync_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-reliablesyncs_kubeedge_io_object_sync_v1alpha1_manifest_test.go: out/install-sentinel terratest/reliablesyncs_kubeedge_io_v1alpha1/reliablesyncs_kubeedge_io_object_sync_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_reliablesyncs_kubeedge_io_object_sync_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/reliablesyncs_kubeedge_io_v1alpha1/reliablesyncs_kubeedge_io_object_sync_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-remediation_medik8s_io_node_health_check_v1alpha1_manifest_test.go: out/install-sentinel terratest/remediation_medik8s_io_v1alpha1/remediation_medik8s_io_node_health_check_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_remediation_medik8s_io_node_health_check_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/remediation_medik8s_io_v1alpha1/remediation_medik8s_io_node_health_check_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-repo_manager_pulpproject_org_pulp_backup_v1beta2_manifest_test.go: out/install-sentinel terratest/repo_manager_pulpproject_org_v1beta2/repo_manager_pulpproject_org_pulp_backup_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_repo_manager_pulpproject_org_pulp_backup_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/repo_manager_pulpproject_org_v1beta2/repo_manager_pulpproject_org_pulp_backup_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-repo_manager_pulpproject_org_pulp_restore_v1beta2_manifest_test.go: out/install-sentinel terratest/repo_manager_pulpproject_org_v1beta2/repo_manager_pulpproject_org_pulp_restore_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_repo_manager_pulpproject_org_pulp_restore_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/repo_manager_pulpproject_org_v1beta2/repo_manager_pulpproject_org_pulp_restore_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-repo_manager_pulpproject_org_pulp_v1beta2_manifest_test.go: out/install-sentinel terratest/repo_manager_pulpproject_org_v1beta2/repo_manager_pulpproject_org_pulp_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_repo_manager_pulpproject_org_pulp_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/repo_manager_pulpproject_org_v1beta2/repo_manager_pulpproject_org_pulp_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-reports_kyverno_io_cluster_ephemeral_report_v1_manifest_test.go: out/install-sentinel terratest/reports_kyverno_io_v1/reports_kyverno_io_cluster_ephemeral_report_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_reports_kyverno_io_cluster_ephemeral_report_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/reports_kyverno_io_v1/reports_kyverno_io_cluster_ephemeral_report_v1_manifest_test.go
	touch $@
out/terratest-sentinel-reports_kyverno_io_ephemeral_report_v1_manifest_test.go: out/install-sentinel terratest/reports_kyverno_io_v1/reports_kyverno_io_ephemeral_report_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_reports_kyverno_io_ephemeral_report_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/reports_kyverno_io_v1/reports_kyverno_io_ephemeral_report_v1_manifest_test.go
	touch $@
out/terratest-sentinel-resources_teleport_dev_teleport_login_rule_v1_manifest_test.go: out/install-sentinel terratest/resources_teleport_dev_v1/resources_teleport_dev_teleport_login_rule_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_resources_teleport_dev_teleport_login_rule_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/resources_teleport_dev_v1/resources_teleport_dev_teleport_login_rule_v1_manifest_test.go
	touch $@
out/terratest-sentinel-resources_teleport_dev_teleport_okta_import_rule_v1_manifest_test.go: out/install-sentinel terratest/resources_teleport_dev_v1/resources_teleport_dev_teleport_okta_import_rule_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_resources_teleport_dev_teleport_okta_import_rule_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/resources_teleport_dev_v1/resources_teleport_dev_teleport_okta_import_rule_v1_manifest_test.go
	touch $@
out/terratest-sentinel-resources_teleport_dev_teleport_provision_token_v2_manifest_test.go: out/install-sentinel terratest/resources_teleport_dev_v2/resources_teleport_dev_teleport_provision_token_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_resources_teleport_dev_teleport_provision_token_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/resources_teleport_dev_v2/resources_teleport_dev_teleport_provision_token_v2_manifest_test.go
	touch $@
out/terratest-sentinel-resources_teleport_dev_teleport_saml_connector_v2_manifest_test.go: out/install-sentinel terratest/resources_teleport_dev_v2/resources_teleport_dev_teleport_saml_connector_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_resources_teleport_dev_teleport_saml_connector_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/resources_teleport_dev_v2/resources_teleport_dev_teleport_saml_connector_v2_manifest_test.go
	touch $@
out/terratest-sentinel-resources_teleport_dev_teleport_user_v2_manifest_test.go: out/install-sentinel terratest/resources_teleport_dev_v2/resources_teleport_dev_teleport_user_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_resources_teleport_dev_teleport_user_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/resources_teleport_dev_v2/resources_teleport_dev_teleport_user_v2_manifest_test.go
	touch $@
out/terratest-sentinel-resources_teleport_dev_teleport_github_connector_v3_manifest_test.go: out/install-sentinel terratest/resources_teleport_dev_v3/resources_teleport_dev_teleport_github_connector_v3_manifest_test.go $(shell find ./examples/data-sources/k8s_resources_teleport_dev_teleport_github_connector_v3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/resources_teleport_dev_v3/resources_teleport_dev_teleport_github_connector_v3_manifest_test.go
	touch $@
out/terratest-sentinel-resources_teleport_dev_teleport_oidc_connector_v3_manifest_test.go: out/install-sentinel terratest/resources_teleport_dev_v3/resources_teleport_dev_teleport_oidc_connector_v3_manifest_test.go $(shell find ./examples/data-sources/k8s_resources_teleport_dev_teleport_oidc_connector_v3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/resources_teleport_dev_v3/resources_teleport_dev_teleport_oidc_connector_v3_manifest_test.go
	touch $@
out/terratest-sentinel-resources_teleport_dev_teleport_role_v5_manifest_test.go: out/install-sentinel terratest/resources_teleport_dev_v5/resources_teleport_dev_teleport_role_v5_manifest_test.go $(shell find ./examples/data-sources/k8s_resources_teleport_dev_teleport_role_v5_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/resources_teleport_dev_v5/resources_teleport_dev_teleport_role_v5_manifest_test.go
	touch $@
out/terratest-sentinel-resources_teleport_dev_teleport_role_v6_manifest_test.go: out/install-sentinel terratest/resources_teleport_dev_v6/resources_teleport_dev_teleport_role_v6_manifest_test.go $(shell find ./examples/data-sources/k8s_resources_teleport_dev_teleport_role_v6_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/resources_teleport_dev_v6/resources_teleport_dev_teleport_role_v6_manifest_test.go
	touch $@
out/terratest-sentinel-ripsaw_cloudbulldozer_io_benchmark_v1alpha1_manifest_test.go: out/install-sentinel terratest/ripsaw_cloudbulldozer_io_v1alpha1/ripsaw_cloudbulldozer_io_benchmark_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_ripsaw_cloudbulldozer_io_benchmark_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/ripsaw_cloudbulldozer_io_v1alpha1/ripsaw_cloudbulldozer_io_benchmark_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-rocketmq_apache_org_broker_v1alpha1_manifest_test.go: out/install-sentinel terratest/rocketmq_apache_org_v1alpha1/rocketmq_apache_org_broker_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_rocketmq_apache_org_broker_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/rocketmq_apache_org_v1alpha1/rocketmq_apache_org_broker_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-rocketmq_apache_org_console_v1alpha1_manifest_test.go: out/install-sentinel terratest/rocketmq_apache_org_v1alpha1/rocketmq_apache_org_console_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_rocketmq_apache_org_console_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/rocketmq_apache_org_v1alpha1/rocketmq_apache_org_console_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-rocketmq_apache_org_name_service_v1alpha1_manifest_test.go: out/install-sentinel terratest/rocketmq_apache_org_v1alpha1/rocketmq_apache_org_name_service_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_rocketmq_apache_org_name_service_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/rocketmq_apache_org_v1alpha1/rocketmq_apache_org_name_service_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-rocketmq_apache_org_topic_transfer_v1alpha1_manifest_test.go: out/install-sentinel terratest/rocketmq_apache_org_v1alpha1/rocketmq_apache_org_topic_transfer_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_rocketmq_apache_org_topic_transfer_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/rocketmq_apache_org_v1alpha1/rocketmq_apache_org_topic_transfer_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-route53_services_k8s_aws_hosted_zone_v1alpha1_manifest_test.go: out/install-sentinel terratest/route53_services_k8s_aws_v1alpha1/route53_services_k8s_aws_hosted_zone_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_route53_services_k8s_aws_hosted_zone_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/route53_services_k8s_aws_v1alpha1/route53_services_k8s_aws_hosted_zone_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-route53_services_k8s_aws_record_set_v1alpha1_manifest_test.go: out/install-sentinel terratest/route53_services_k8s_aws_v1alpha1/route53_services_k8s_aws_record_set_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_route53_services_k8s_aws_record_set_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/route53_services_k8s_aws_v1alpha1/route53_services_k8s_aws_record_set_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-route53resolver_services_k8s_aws_resolver_endpoint_v1alpha1_manifest_test.go: out/install-sentinel terratest/route53resolver_services_k8s_aws_v1alpha1/route53resolver_services_k8s_aws_resolver_endpoint_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_route53resolver_services_k8s_aws_resolver_endpoint_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/route53resolver_services_k8s_aws_v1alpha1/route53resolver_services_k8s_aws_resolver_endpoint_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-route53resolver_services_k8s_aws_resolver_rule_v1alpha1_manifest_test.go: out/install-sentinel terratest/route53resolver_services_k8s_aws_v1alpha1/route53resolver_services_k8s_aws_resolver_rule_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_route53resolver_services_k8s_aws_resolver_rule_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/route53resolver_services_k8s_aws_v1alpha1/route53resolver_services_k8s_aws_resolver_rule_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-rules_kubeedge_io_rule_endpoint_v1_manifest_test.go: out/install-sentinel terratest/rules_kubeedge_io_v1/rules_kubeedge_io_rule_endpoint_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_rules_kubeedge_io_rule_endpoint_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/rules_kubeedge_io_v1/rules_kubeedge_io_rule_endpoint_v1_manifest_test.go
	touch $@
out/terratest-sentinel-rules_kubeedge_io_rule_v1_manifest_test.go: out/install-sentinel terratest/rules_kubeedge_io_v1/rules_kubeedge_io_rule_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_rules_kubeedge_io_rule_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/rules_kubeedge_io_v1/rules_kubeedge_io_rule_v1_manifest_test.go
	touch $@
out/terratest-sentinel-runtime_cluster_x_k8s_io_extension_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/runtime_cluster_x_k8s_io_v1alpha1/runtime_cluster_x_k8s_io_extension_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_runtime_cluster_x_k8s_io_extension_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/runtime_cluster_x_k8s_io_v1alpha1/runtime_cluster_x_k8s_io_extension_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-s3_services_k8s_aws_bucket_v1alpha1_manifest_test.go: out/install-sentinel terratest/s3_services_k8s_aws_v1alpha1/s3_services_k8s_aws_bucket_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_s3_services_k8s_aws_bucket_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/s3_services_k8s_aws_v1alpha1/s3_services_k8s_aws_bucket_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-s3_snappcloud_io_s3_bucket_v1alpha1_manifest_test.go: out/install-sentinel terratest/s3_snappcloud_io_v1alpha1/s3_snappcloud_io_s3_bucket_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_s3_snappcloud_io_s3_bucket_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/s3_snappcloud_io_v1alpha1/s3_snappcloud_io_s3_bucket_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-s3_snappcloud_io_s3_user_claim_v1alpha1_manifest_test.go: out/install-sentinel terratest/s3_snappcloud_io_v1alpha1/s3_snappcloud_io_s3_user_claim_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_s3_snappcloud_io_s3_user_claim_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/s3_snappcloud_io_v1alpha1/s3_snappcloud_io_s3_user_claim_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-s3_snappcloud_io_s3_user_v1alpha1_manifest_test.go: out/install-sentinel terratest/s3_snappcloud_io_v1alpha1/s3_snappcloud_io_s3_user_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_s3_snappcloud_io_s3_user_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/s3_snappcloud_io_v1alpha1/s3_snappcloud_io_s3_user_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-sagemaker_services_k8s_aws_app_v1alpha1_manifest_test.go: out/install-sentinel terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_app_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_sagemaker_services_k8s_aws_app_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_app_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-sagemaker_services_k8s_aws_data_quality_job_definition_v1alpha1_manifest_test.go: out/install-sentinel terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_data_quality_job_definition_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_sagemaker_services_k8s_aws_data_quality_job_definition_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_data_quality_job_definition_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-sagemaker_services_k8s_aws_domain_v1alpha1_manifest_test.go: out/install-sentinel terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_domain_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_sagemaker_services_k8s_aws_domain_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_domain_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-sagemaker_services_k8s_aws_endpoint_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_endpoint_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_sagemaker_services_k8s_aws_endpoint_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_endpoint_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-sagemaker_services_k8s_aws_endpoint_v1alpha1_manifest_test.go: out/install-sentinel terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_endpoint_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_sagemaker_services_k8s_aws_endpoint_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_endpoint_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-sagemaker_services_k8s_aws_feature_group_v1alpha1_manifest_test.go: out/install-sentinel terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_feature_group_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_sagemaker_services_k8s_aws_feature_group_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_feature_group_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-sagemaker_services_k8s_aws_hyper_parameter_tuning_job_v1alpha1_manifest_test.go: out/install-sentinel terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_hyper_parameter_tuning_job_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_sagemaker_services_k8s_aws_hyper_parameter_tuning_job_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_hyper_parameter_tuning_job_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-sagemaker_services_k8s_aws_model_bias_job_definition_v1alpha1_manifest_test.go: out/install-sentinel terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_model_bias_job_definition_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_sagemaker_services_k8s_aws_model_bias_job_definition_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_model_bias_job_definition_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-sagemaker_services_k8s_aws_model_explainability_job_definition_v1alpha1_manifest_test.go: out/install-sentinel terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_model_explainability_job_definition_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_sagemaker_services_k8s_aws_model_explainability_job_definition_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_model_explainability_job_definition_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-sagemaker_services_k8s_aws_model_package_group_v1alpha1_manifest_test.go: out/install-sentinel terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_model_package_group_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_sagemaker_services_k8s_aws_model_package_group_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_model_package_group_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-sagemaker_services_k8s_aws_model_package_v1alpha1_manifest_test.go: out/install-sentinel terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_model_package_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_sagemaker_services_k8s_aws_model_package_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_model_package_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-sagemaker_services_k8s_aws_model_quality_job_definition_v1alpha1_manifest_test.go: out/install-sentinel terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_model_quality_job_definition_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_sagemaker_services_k8s_aws_model_quality_job_definition_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_model_quality_job_definition_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-sagemaker_services_k8s_aws_model_v1alpha1_manifest_test.go: out/install-sentinel terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_model_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_sagemaker_services_k8s_aws_model_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_model_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-sagemaker_services_k8s_aws_monitoring_schedule_v1alpha1_manifest_test.go: out/install-sentinel terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_monitoring_schedule_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_sagemaker_services_k8s_aws_monitoring_schedule_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_monitoring_schedule_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-sagemaker_services_k8s_aws_notebook_instance_lifecycle_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_notebook_instance_lifecycle_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_sagemaker_services_k8s_aws_notebook_instance_lifecycle_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_notebook_instance_lifecycle_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-sagemaker_services_k8s_aws_notebook_instance_v1alpha1_manifest_test.go: out/install-sentinel terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_notebook_instance_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_sagemaker_services_k8s_aws_notebook_instance_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_notebook_instance_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-sagemaker_services_k8s_aws_processing_job_v1alpha1_manifest_test.go: out/install-sentinel terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_processing_job_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_sagemaker_services_k8s_aws_processing_job_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_processing_job_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-sagemaker_services_k8s_aws_training_job_v1alpha1_manifest_test.go: out/install-sentinel terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_training_job_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_sagemaker_services_k8s_aws_training_job_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_training_job_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-sagemaker_services_k8s_aws_transform_job_v1alpha1_manifest_test.go: out/install-sentinel terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_transform_job_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_sagemaker_services_k8s_aws_transform_job_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_transform_job_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-sagemaker_services_k8s_aws_user_profile_v1alpha1_manifest_test.go: out/install-sentinel terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_user_profile_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_sagemaker_services_k8s_aws_user_profile_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/sagemaker_services_k8s_aws_v1alpha1/sagemaker_services_k8s_aws_user_profile_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-scheduling_k8s_io_priority_class_v1_manifest_test.go: out/install-sentinel terratest/scheduling_k8s_io_v1/scheduling_k8s_io_priority_class_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_scheduling_k8s_io_priority_class_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/scheduling_k8s_io_v1/scheduling_k8s_io_priority_class_v1_manifest_test.go
	touch $@
out/terratest-sentinel-scheduling_koordinator_sh_device_v1alpha1_manifest_test.go: out/install-sentinel terratest/scheduling_koordinator_sh_v1alpha1/scheduling_koordinator_sh_device_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_scheduling_koordinator_sh_device_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/scheduling_koordinator_sh_v1alpha1/scheduling_koordinator_sh_device_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-scheduling_koordinator_sh_pod_migration_job_v1alpha1_manifest_test.go: out/install-sentinel terratest/scheduling_koordinator_sh_v1alpha1/scheduling_koordinator_sh_pod_migration_job_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_scheduling_koordinator_sh_pod_migration_job_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/scheduling_koordinator_sh_v1alpha1/scheduling_koordinator_sh_pod_migration_job_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-scheduling_koordinator_sh_reservation_v1alpha1_manifest_test.go: out/install-sentinel terratest/scheduling_koordinator_sh_v1alpha1/scheduling_koordinator_sh_reservation_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_scheduling_koordinator_sh_reservation_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/scheduling_koordinator_sh_v1alpha1/scheduling_koordinator_sh_reservation_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-scheduling_sigs_k8s_io_elastic_quota_v1alpha1_manifest_test.go: out/install-sentinel terratest/scheduling_sigs_k8s_io_v1alpha1/scheduling_sigs_k8s_io_elastic_quota_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_scheduling_sigs_k8s_io_elastic_quota_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/scheduling_sigs_k8s_io_v1alpha1/scheduling_sigs_k8s_io_elastic_quota_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-scheduling_sigs_k8s_io_pod_group_v1alpha1_manifest_test.go: out/install-sentinel terratest/scheduling_sigs_k8s_io_v1alpha1/scheduling_sigs_k8s_io_pod_group_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_scheduling_sigs_k8s_io_pod_group_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/scheduling_sigs_k8s_io_v1alpha1/scheduling_sigs_k8s_io_pod_group_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-scheduling_volcano_sh_pod_group_v1beta1_manifest_test.go: out/install-sentinel terratest/scheduling_volcano_sh_v1beta1/scheduling_volcano_sh_pod_group_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_scheduling_volcano_sh_pod_group_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/scheduling_volcano_sh_v1beta1/scheduling_volcano_sh_pod_group_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-scheduling_volcano_sh_queue_v1beta1_manifest_test.go: out/install-sentinel terratest/scheduling_volcano_sh_v1beta1/scheduling_volcano_sh_queue_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_scheduling_volcano_sh_queue_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/scheduling_volcano_sh_v1beta1/scheduling_volcano_sh_queue_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-schemas_schemahero_io_data_type_v1alpha4_manifest_test.go: out/install-sentinel terratest/schemas_schemahero_io_v1alpha4/schemas_schemahero_io_data_type_v1alpha4_manifest_test.go $(shell find ./examples/data-sources/k8s_schemas_schemahero_io_data_type_v1alpha4_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/schemas_schemahero_io_v1alpha4/schemas_schemahero_io_data_type_v1alpha4_manifest_test.go
	touch $@
out/terratest-sentinel-schemas_schemahero_io_migration_v1alpha4_manifest_test.go: out/install-sentinel terratest/schemas_schemahero_io_v1alpha4/schemas_schemahero_io_migration_v1alpha4_manifest_test.go $(shell find ./examples/data-sources/k8s_schemas_schemahero_io_migration_v1alpha4_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/schemas_schemahero_io_v1alpha4/schemas_schemahero_io_migration_v1alpha4_manifest_test.go
	touch $@
out/terratest-sentinel-schemas_schemahero_io_table_v1alpha4_manifest_test.go: out/install-sentinel terratest/schemas_schemahero_io_v1alpha4/schemas_schemahero_io_table_v1alpha4_manifest_test.go $(shell find ./examples/data-sources/k8s_schemas_schemahero_io_table_v1alpha4_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/schemas_schemahero_io_v1alpha4/schemas_schemahero_io_table_v1alpha4_manifest_test.go
	touch $@
out/terratest-sentinel-scylla_scylladb_com_scylla_cluster_v1_manifest_test.go: out/install-sentinel terratest/scylla_scylladb_com_v1/scylla_scylladb_com_scylla_cluster_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_scylla_scylladb_com_scylla_cluster_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/scylla_scylladb_com_v1/scylla_scylladb_com_scylla_cluster_v1_manifest_test.go
	touch $@
out/terratest-sentinel-scylla_scylladb_com_node_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/scylla_scylladb_com_v1alpha1/scylla_scylladb_com_node_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_scylla_scylladb_com_node_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/scylla_scylladb_com_v1alpha1/scylla_scylladb_com_node_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-scylla_scylladb_com_scylla_operator_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/scylla_scylladb_com_v1alpha1/scylla_scylladb_com_scylla_operator_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_scylla_scylladb_com_scylla_operator_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/scylla_scylladb_com_v1alpha1/scylla_scylladb_com_scylla_operator_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-secretgenerator_mittwald_de_basic_auth_v1alpha1_manifest_test.go: out/install-sentinel terratest/secretgenerator_mittwald_de_v1alpha1/secretgenerator_mittwald_de_basic_auth_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_secretgenerator_mittwald_de_basic_auth_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/secretgenerator_mittwald_de_v1alpha1/secretgenerator_mittwald_de_basic_auth_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-secretgenerator_mittwald_de_ssh_key_pair_v1alpha1_manifest_test.go: out/install-sentinel terratest/secretgenerator_mittwald_de_v1alpha1/secretgenerator_mittwald_de_ssh_key_pair_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_secretgenerator_mittwald_de_ssh_key_pair_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/secretgenerator_mittwald_de_v1alpha1/secretgenerator_mittwald_de_ssh_key_pair_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-secretgenerator_mittwald_de_string_secret_v1alpha1_manifest_test.go: out/install-sentinel terratest/secretgenerator_mittwald_de_v1alpha1/secretgenerator_mittwald_de_string_secret_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_secretgenerator_mittwald_de_string_secret_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/secretgenerator_mittwald_de_v1alpha1/secretgenerator_mittwald_de_string_secret_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-secrets_crossplane_io_store_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/secrets_crossplane_io_v1alpha1/secrets_crossplane_io_store_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_secrets_crossplane_io_store_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/secrets_crossplane_io_v1alpha1/secrets_crossplane_io_store_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-secrets_doppler_com_doppler_secret_v1alpha1_manifest_test.go: out/install-sentinel terratest/secrets_doppler_com_v1alpha1/secrets_doppler_com_doppler_secret_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_secrets_doppler_com_doppler_secret_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/secrets_doppler_com_v1alpha1/secrets_doppler_com_doppler_secret_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-secrets_hashicorp_com_hcp_auth_v1beta1_manifest_test.go: out/install-sentinel terratest/secrets_hashicorp_com_v1beta1/secrets_hashicorp_com_hcp_auth_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_secrets_hashicorp_com_hcp_auth_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/secrets_hashicorp_com_v1beta1/secrets_hashicorp_com_hcp_auth_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-secrets_hashicorp_com_hcp_vault_secrets_app_v1beta1_manifest_test.go: out/install-sentinel terratest/secrets_hashicorp_com_v1beta1/secrets_hashicorp_com_hcp_vault_secrets_app_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_secrets_hashicorp_com_hcp_vault_secrets_app_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/secrets_hashicorp_com_v1beta1/secrets_hashicorp_com_hcp_vault_secrets_app_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-secrets_hashicorp_com_vault_auth_v1beta1_manifest_test.go: out/install-sentinel terratest/secrets_hashicorp_com_v1beta1/secrets_hashicorp_com_vault_auth_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_secrets_hashicorp_com_vault_auth_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/secrets_hashicorp_com_v1beta1/secrets_hashicorp_com_vault_auth_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-secrets_hashicorp_com_vault_connection_v1beta1_manifest_test.go: out/install-sentinel terratest/secrets_hashicorp_com_v1beta1/secrets_hashicorp_com_vault_connection_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_secrets_hashicorp_com_vault_connection_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/secrets_hashicorp_com_v1beta1/secrets_hashicorp_com_vault_connection_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-secrets_hashicorp_com_vault_dynamic_secret_v1beta1_manifest_test.go: out/install-sentinel terratest/secrets_hashicorp_com_v1beta1/secrets_hashicorp_com_vault_dynamic_secret_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_secrets_hashicorp_com_vault_dynamic_secret_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/secrets_hashicorp_com_v1beta1/secrets_hashicorp_com_vault_dynamic_secret_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-secrets_hashicorp_com_vault_pki_secret_v1beta1_manifest_test.go: out/install-sentinel terratest/secrets_hashicorp_com_v1beta1/secrets_hashicorp_com_vault_pki_secret_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_secrets_hashicorp_com_vault_pki_secret_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/secrets_hashicorp_com_v1beta1/secrets_hashicorp_com_vault_pki_secret_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-secrets_hashicorp_com_vault_static_secret_v1beta1_manifest_test.go: out/install-sentinel terratest/secrets_hashicorp_com_v1beta1/secrets_hashicorp_com_vault_static_secret_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_secrets_hashicorp_com_vault_static_secret_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/secrets_hashicorp_com_v1beta1/secrets_hashicorp_com_vault_static_secret_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-secrets_store_csi_x_k8s_io_secret_provider_class_v1_manifest_test.go: out/install-sentinel terratest/secrets_store_csi_x_k8s_io_v1/secrets_store_csi_x_k8s_io_secret_provider_class_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_secrets_store_csi_x_k8s_io_secret_provider_class_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/secrets_store_csi_x_k8s_io_v1/secrets_store_csi_x_k8s_io_secret_provider_class_v1_manifest_test.go
	touch $@
out/terratest-sentinel-secrets_store_csi_x_k8s_io_secret_provider_class_v1alpha1_manifest_test.go: out/install-sentinel terratest/secrets_store_csi_x_k8s_io_v1alpha1/secrets_store_csi_x_k8s_io_secret_provider_class_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_secrets_store_csi_x_k8s_io_secret_provider_class_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/secrets_store_csi_x_k8s_io_v1alpha1/secrets_store_csi_x_k8s_io_secret_provider_class_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-secretsmanager_services_k8s_aws_secret_v1alpha1_manifest_test.go: out/install-sentinel terratest/secretsmanager_services_k8s_aws_v1alpha1/secretsmanager_services_k8s_aws_secret_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_secretsmanager_services_k8s_aws_secret_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/secretsmanager_services_k8s_aws_v1alpha1/secretsmanager_services_k8s_aws_secret_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-secscan_quay_redhat_com_image_manifest_vuln_v1alpha1_manifest_test.go: out/install-sentinel terratest/secscan_quay_redhat_com_v1alpha1/secscan_quay_redhat_com_image_manifest_vuln_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_secscan_quay_redhat_com_image_manifest_vuln_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/secscan_quay_redhat_com_v1alpha1/secscan_quay_redhat_com_image_manifest_vuln_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-security_istio_io_authorization_policy_v1_manifest_test.go: out/install-sentinel terratest/security_istio_io_v1/security_istio_io_authorization_policy_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_security_istio_io_authorization_policy_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/security_istio_io_v1/security_istio_io_authorization_policy_v1_manifest_test.go
	touch $@
out/terratest-sentinel-security_istio_io_peer_authentication_v1_manifest_test.go: out/install-sentinel terratest/security_istio_io_v1/security_istio_io_peer_authentication_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_security_istio_io_peer_authentication_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/security_istio_io_v1/security_istio_io_peer_authentication_v1_manifest_test.go
	touch $@
out/terratest-sentinel-security_istio_io_request_authentication_v1_manifest_test.go: out/install-sentinel terratest/security_istio_io_v1/security_istio_io_request_authentication_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_security_istio_io_request_authentication_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/security_istio_io_v1/security_istio_io_request_authentication_v1_manifest_test.go
	touch $@
out/terratest-sentinel-security_istio_io_authorization_policy_v1beta1_manifest_test.go: out/install-sentinel terratest/security_istio_io_v1beta1/security_istio_io_authorization_policy_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_security_istio_io_authorization_policy_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/security_istio_io_v1beta1/security_istio_io_authorization_policy_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-security_istio_io_peer_authentication_v1beta1_manifest_test.go: out/install-sentinel terratest/security_istio_io_v1beta1/security_istio_io_peer_authentication_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_security_istio_io_peer_authentication_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/security_istio_io_v1beta1/security_istio_io_peer_authentication_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-security_istio_io_request_authentication_v1beta1_manifest_test.go: out/install-sentinel terratest/security_istio_io_v1beta1/security_istio_io_request_authentication_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_security_istio_io_request_authentication_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/security_istio_io_v1beta1/security_istio_io_request_authentication_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-security_profiles_operator_x_k8s_io_app_armor_profile_v1alpha1_manifest_test.go: out/install-sentinel terratest/security_profiles_operator_x_k8s_io_v1alpha1/security_profiles_operator_x_k8s_io_app_armor_profile_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_security_profiles_operator_x_k8s_io_app_armor_profile_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/security_profiles_operator_x_k8s_io_v1alpha1/security_profiles_operator_x_k8s_io_app_armor_profile_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-security_profiles_operator_x_k8s_io_profile_binding_v1alpha1_manifest_test.go: out/install-sentinel terratest/security_profiles_operator_x_k8s_io_v1alpha1/security_profiles_operator_x_k8s_io_profile_binding_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_security_profiles_operator_x_k8s_io_profile_binding_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/security_profiles_operator_x_k8s_io_v1alpha1/security_profiles_operator_x_k8s_io_profile_binding_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-security_profiles_operator_x_k8s_io_profile_recording_v1alpha1_manifest_test.go: out/install-sentinel terratest/security_profiles_operator_x_k8s_io_v1alpha1/security_profiles_operator_x_k8s_io_profile_recording_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_security_profiles_operator_x_k8s_io_profile_recording_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/security_profiles_operator_x_k8s_io_v1alpha1/security_profiles_operator_x_k8s_io_profile_recording_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-security_profiles_operator_x_k8s_io_security_profile_node_status_v1alpha1_manifest_test.go: out/install-sentinel terratest/security_profiles_operator_x_k8s_io_v1alpha1/security_profiles_operator_x_k8s_io_security_profile_node_status_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_security_profiles_operator_x_k8s_io_security_profile_node_status_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/security_profiles_operator_x_k8s_io_v1alpha1/security_profiles_operator_x_k8s_io_security_profile_node_status_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-security_profiles_operator_x_k8s_io_security_profiles_operator_daemon_v1alpha1_manifest_test.go: out/install-sentinel terratest/security_profiles_operator_x_k8s_io_v1alpha1/security_profiles_operator_x_k8s_io_security_profiles_operator_daemon_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_security_profiles_operator_x_k8s_io_security_profiles_operator_daemon_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/security_profiles_operator_x_k8s_io_v1alpha1/security_profiles_operator_x_k8s_io_security_profiles_operator_daemon_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-security_profiles_operator_x_k8s_io_raw_selinux_profile_v1alpha2_manifest_test.go: out/install-sentinel terratest/security_profiles_operator_x_k8s_io_v1alpha2/security_profiles_operator_x_k8s_io_raw_selinux_profile_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_security_profiles_operator_x_k8s_io_raw_selinux_profile_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/security_profiles_operator_x_k8s_io_v1alpha2/security_profiles_operator_x_k8s_io_raw_selinux_profile_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-security_profiles_operator_x_k8s_io_selinux_profile_v1alpha2_manifest_test.go: out/install-sentinel terratest/security_profiles_operator_x_k8s_io_v1alpha2/security_profiles_operator_x_k8s_io_selinux_profile_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_security_profiles_operator_x_k8s_io_selinux_profile_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/security_profiles_operator_x_k8s_io_v1alpha2/security_profiles_operator_x_k8s_io_selinux_profile_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-security_profiles_operator_x_k8s_io_seccomp_profile_v1beta1_manifest_test.go: out/install-sentinel terratest/security_profiles_operator_x_k8s_io_v1beta1/security_profiles_operator_x_k8s_io_seccomp_profile_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_security_profiles_operator_x_k8s_io_seccomp_profile_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/security_profiles_operator_x_k8s_io_v1beta1/security_profiles_operator_x_k8s_io_seccomp_profile_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-self_node_remediation_medik8s_io_self_node_remediation_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/self_node_remediation_medik8s_io_v1alpha1/self_node_remediation_medik8s_io_self_node_remediation_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_self_node_remediation_medik8s_io_self_node_remediation_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/self_node_remediation_medik8s_io_v1alpha1/self_node_remediation_medik8s_io_self_node_remediation_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-self_node_remediation_medik8s_io_self_node_remediation_template_v1alpha1_manifest_test.go: out/install-sentinel terratest/self_node_remediation_medik8s_io_v1alpha1/self_node_remediation_medik8s_io_self_node_remediation_template_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_self_node_remediation_medik8s_io_self_node_remediation_template_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/self_node_remediation_medik8s_io_v1alpha1/self_node_remediation_medik8s_io_self_node_remediation_template_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-self_node_remediation_medik8s_io_self_node_remediation_v1alpha1_manifest_test.go: out/install-sentinel terratest/self_node_remediation_medik8s_io_v1alpha1/self_node_remediation_medik8s_io_self_node_remediation_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_self_node_remediation_medik8s_io_self_node_remediation_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/self_node_remediation_medik8s_io_v1alpha1/self_node_remediation_medik8s_io_self_node_remediation_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-sematext_com_sematext_agent_v1_manifest_test.go: out/install-sentinel terratest/sematext_com_v1/sematext_com_sematext_agent_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_sematext_com_sematext_agent_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/sematext_com_v1/sematext_com_sematext_agent_v1_manifest_test.go
	touch $@
out/terratest-sentinel-servicebinding_io_cluster_workload_resource_mapping_v1alpha3_manifest_test.go: out/install-sentinel terratest/servicebinding_io_v1alpha3/servicebinding_io_cluster_workload_resource_mapping_v1alpha3_manifest_test.go $(shell find ./examples/data-sources/k8s_servicebinding_io_cluster_workload_resource_mapping_v1alpha3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/servicebinding_io_v1alpha3/servicebinding_io_cluster_workload_resource_mapping_v1alpha3_manifest_test.go
	touch $@
out/terratest-sentinel-servicebinding_io_service_binding_v1alpha3_manifest_test.go: out/install-sentinel terratest/servicebinding_io_v1alpha3/servicebinding_io_service_binding_v1alpha3_manifest_test.go $(shell find ./examples/data-sources/k8s_servicebinding_io_service_binding_v1alpha3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/servicebinding_io_v1alpha3/servicebinding_io_service_binding_v1alpha3_manifest_test.go
	touch $@
out/terratest-sentinel-servicebinding_io_cluster_workload_resource_mapping_v1beta1_manifest_test.go: out/install-sentinel terratest/servicebinding_io_v1beta1/servicebinding_io_cluster_workload_resource_mapping_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_servicebinding_io_cluster_workload_resource_mapping_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/servicebinding_io_v1beta1/servicebinding_io_cluster_workload_resource_mapping_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-servicebinding_io_service_binding_v1beta1_manifest_test.go: out/install-sentinel terratest/servicebinding_io_v1beta1/servicebinding_io_service_binding_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_servicebinding_io_service_binding_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/servicebinding_io_v1beta1/servicebinding_io_service_binding_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-servicemesh_cisco_com_istio_control_plane_v1alpha1_manifest_test.go: out/install-sentinel terratest/servicemesh_cisco_com_v1alpha1/servicemesh_cisco_com_istio_control_plane_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_servicemesh_cisco_com_istio_control_plane_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/servicemesh_cisco_com_v1alpha1/servicemesh_cisco_com_istio_control_plane_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-servicemesh_cisco_com_istio_mesh_gateway_v1alpha1_manifest_test.go: out/install-sentinel terratest/servicemesh_cisco_com_v1alpha1/servicemesh_cisco_com_istio_mesh_gateway_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_servicemesh_cisco_com_istio_mesh_gateway_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/servicemesh_cisco_com_v1alpha1/servicemesh_cisco_com_istio_mesh_gateway_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-servicemesh_cisco_com_istio_mesh_v1alpha1_manifest_test.go: out/install-sentinel terratest/servicemesh_cisco_com_v1alpha1/servicemesh_cisco_com_istio_mesh_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_servicemesh_cisco_com_istio_mesh_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/servicemesh_cisco_com_v1alpha1/servicemesh_cisco_com_istio_mesh_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-servicemesh_cisco_com_peer_istio_control_plane_v1alpha1_manifest_test.go: out/install-sentinel terratest/servicemesh_cisco_com_v1alpha1/servicemesh_cisco_com_peer_istio_control_plane_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_servicemesh_cisco_com_peer_istio_control_plane_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/servicemesh_cisco_com_v1alpha1/servicemesh_cisco_com_peer_istio_control_plane_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-services_k8s_aws_adopted_resource_v1alpha1_manifest_test.go: out/install-sentinel terratest/services_k8s_aws_v1alpha1/services_k8s_aws_adopted_resource_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_services_k8s_aws_adopted_resource_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/services_k8s_aws_v1alpha1/services_k8s_aws_adopted_resource_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-services_k8s_aws_field_export_v1alpha1_manifest_test.go: out/install-sentinel terratest/services_k8s_aws_v1alpha1/services_k8s_aws_field_export_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_services_k8s_aws_field_export_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/services_k8s_aws_v1alpha1/services_k8s_aws_field_export_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-serving_kubedl_io_inference_v1alpha1_manifest_test.go: out/install-sentinel terratest/serving_kubedl_io_v1alpha1/serving_kubedl_io_inference_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_serving_kubedl_io_inference_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/serving_kubedl_io_v1alpha1/serving_kubedl_io_inference_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-sfn_services_k8s_aws_activity_v1alpha1_manifest_test.go: out/install-sentinel terratest/sfn_services_k8s_aws_v1alpha1/sfn_services_k8s_aws_activity_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_sfn_services_k8s_aws_activity_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/sfn_services_k8s_aws_v1alpha1/sfn_services_k8s_aws_activity_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-sfn_services_k8s_aws_state_machine_v1alpha1_manifest_test.go: out/install-sentinel terratest/sfn_services_k8s_aws_v1alpha1/sfn_services_k8s_aws_state_machine_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_sfn_services_k8s_aws_state_machine_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/sfn_services_k8s_aws_v1alpha1/sfn_services_k8s_aws_state_machine_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-site_superedge_io_node_group_v1alpha1_manifest_test.go: out/install-sentinel terratest/site_superedge_io_v1alpha1/site_superedge_io_node_group_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_site_superedge_io_node_group_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/site_superedge_io_v1alpha1/site_superedge_io_node_group_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-site_superedge_io_node_unit_v1alpha1_manifest_test.go: out/install-sentinel terratest/site_superedge_io_v1alpha1/site_superedge_io_node_unit_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_site_superedge_io_node_unit_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/site_superedge_io_v1alpha1/site_superedge_io_node_unit_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-slo_koordinator_sh_node_metric_v1alpha1_manifest_test.go: out/install-sentinel terratest/slo_koordinator_sh_v1alpha1/slo_koordinator_sh_node_metric_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_slo_koordinator_sh_node_metric_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/slo_koordinator_sh_v1alpha1/slo_koordinator_sh_node_metric_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-slo_koordinator_sh_node_slo_v1alpha1_manifest_test.go: out/install-sentinel terratest/slo_koordinator_sh_v1alpha1/slo_koordinator_sh_node_slo_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_slo_koordinator_sh_node_slo_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/slo_koordinator_sh_v1alpha1/slo_koordinator_sh_node_slo_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-sloth_slok_dev_prometheus_service_level_v1_manifest_test.go: out/install-sentinel terratest/sloth_slok_dev_v1/sloth_slok_dev_prometheus_service_level_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_sloth_slok_dev_prometheus_service_level_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/sloth_slok_dev_v1/sloth_slok_dev_prometheus_service_level_v1_manifest_test.go
	touch $@
out/terratest-sentinel-snapscheduler_backube_snapshot_schedule_v1_manifest_test.go: out/install-sentinel terratest/snapscheduler_backube_v1/snapscheduler_backube_snapshot_schedule_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_snapscheduler_backube_snapshot_schedule_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/snapscheduler_backube_v1/snapscheduler_backube_snapshot_schedule_v1_manifest_test.go
	touch $@
out/terratest-sentinel-snapshot_storage_k8s_io_volume_snapshot_class_v1_manifest_test.go: out/install-sentinel terratest/snapshot_storage_k8s_io_v1/snapshot_storage_k8s_io_volume_snapshot_class_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_snapshot_storage_k8s_io_volume_snapshot_class_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/snapshot_storage_k8s_io_v1/snapshot_storage_k8s_io_volume_snapshot_class_v1_manifest_test.go
	touch $@
out/terratest-sentinel-snapshot_storage_k8s_io_volume_snapshot_content_v1_manifest_test.go: out/install-sentinel terratest/snapshot_storage_k8s_io_v1/snapshot_storage_k8s_io_volume_snapshot_content_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_snapshot_storage_k8s_io_volume_snapshot_content_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/snapshot_storage_k8s_io_v1/snapshot_storage_k8s_io_volume_snapshot_content_v1_manifest_test.go
	touch $@
out/terratest-sentinel-snapshot_storage_k8s_io_volume_snapshot_v1_manifest_test.go: out/install-sentinel terratest/snapshot_storage_k8s_io_v1/snapshot_storage_k8s_io_volume_snapshot_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_snapshot_storage_k8s_io_volume_snapshot_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/snapshot_storage_k8s_io_v1/snapshot_storage_k8s_io_volume_snapshot_v1_manifest_test.go
	touch $@
out/terratest-sentinel-snapshot_storage_k8s_io_volume_snapshot_class_v1beta1_manifest_test.go: out/install-sentinel terratest/snapshot_storage_k8s_io_v1beta1/snapshot_storage_k8s_io_volume_snapshot_class_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_snapshot_storage_k8s_io_volume_snapshot_class_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/snapshot_storage_k8s_io_v1beta1/snapshot_storage_k8s_io_volume_snapshot_class_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-snapshot_storage_k8s_io_volume_snapshot_content_v1beta1_manifest_test.go: out/install-sentinel terratest/snapshot_storage_k8s_io_v1beta1/snapshot_storage_k8s_io_volume_snapshot_content_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_snapshot_storage_k8s_io_volume_snapshot_content_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/snapshot_storage_k8s_io_v1beta1/snapshot_storage_k8s_io_volume_snapshot_content_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-snapshot_storage_k8s_io_volume_snapshot_v1beta1_manifest_test.go: out/install-sentinel terratest/snapshot_storage_k8s_io_v1beta1/snapshot_storage_k8s_io_volume_snapshot_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_snapshot_storage_k8s_io_volume_snapshot_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/snapshot_storage_k8s_io_v1beta1/snapshot_storage_k8s_io_volume_snapshot_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-sns_services_k8s_aws_platform_application_v1alpha1_manifest_test.go: out/install-sentinel terratest/sns_services_k8s_aws_v1alpha1/sns_services_k8s_aws_platform_application_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_sns_services_k8s_aws_platform_application_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/sns_services_k8s_aws_v1alpha1/sns_services_k8s_aws_platform_application_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-sns_services_k8s_aws_platform_endpoint_v1alpha1_manifest_test.go: out/install-sentinel terratest/sns_services_k8s_aws_v1alpha1/sns_services_k8s_aws_platform_endpoint_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_sns_services_k8s_aws_platform_endpoint_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/sns_services_k8s_aws_v1alpha1/sns_services_k8s_aws_platform_endpoint_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-sns_services_k8s_aws_subscription_v1alpha1_manifest_test.go: out/install-sentinel terratest/sns_services_k8s_aws_v1alpha1/sns_services_k8s_aws_subscription_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_sns_services_k8s_aws_subscription_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/sns_services_k8s_aws_v1alpha1/sns_services_k8s_aws_subscription_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-sns_services_k8s_aws_topic_v1alpha1_manifest_test.go: out/install-sentinel terratest/sns_services_k8s_aws_v1alpha1/sns_services_k8s_aws_topic_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_sns_services_k8s_aws_topic_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/sns_services_k8s_aws_v1alpha1/sns_services_k8s_aws_topic_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-sonataflow_org_sonata_flow_build_v1alpha08_manifest_test.go: out/install-sentinel terratest/sonataflow_org_v1alpha08/sonataflow_org_sonata_flow_build_v1alpha08_manifest_test.go $(shell find ./examples/data-sources/k8s_sonataflow_org_sonata_flow_build_v1alpha08_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/sonataflow_org_v1alpha08/sonataflow_org_sonata_flow_build_v1alpha08_manifest_test.go
	touch $@
out/terratest-sentinel-sonataflow_org_sonata_flow_platform_v1alpha08_manifest_test.go: out/install-sentinel terratest/sonataflow_org_v1alpha08/sonataflow_org_sonata_flow_platform_v1alpha08_manifest_test.go $(shell find ./examples/data-sources/k8s_sonataflow_org_sonata_flow_platform_v1alpha08_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/sonataflow_org_v1alpha08/sonataflow_org_sonata_flow_platform_v1alpha08_manifest_test.go
	touch $@
out/terratest-sentinel-sonataflow_org_sonata_flow_v1alpha08_manifest_test.go: out/install-sentinel terratest/sonataflow_org_v1alpha08/sonataflow_org_sonata_flow_v1alpha08_manifest_test.go $(shell find ./examples/data-sources/k8s_sonataflow_org_sonata_flow_v1alpha08_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/sonataflow_org_v1alpha08/sonataflow_org_sonata_flow_v1alpha08_manifest_test.go
	touch $@
out/terratest-sentinel-source_toolkit_fluxcd_io_git_repository_v1_manifest_test.go: out/install-sentinel terratest/source_toolkit_fluxcd_io_v1/source_toolkit_fluxcd_io_git_repository_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_source_toolkit_fluxcd_io_git_repository_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/source_toolkit_fluxcd_io_v1/source_toolkit_fluxcd_io_git_repository_v1_manifest_test.go
	touch $@
out/terratest-sentinel-source_toolkit_fluxcd_io_helm_chart_v1_manifest_test.go: out/install-sentinel terratest/source_toolkit_fluxcd_io_v1/source_toolkit_fluxcd_io_helm_chart_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_source_toolkit_fluxcd_io_helm_chart_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/source_toolkit_fluxcd_io_v1/source_toolkit_fluxcd_io_helm_chart_v1_manifest_test.go
	touch $@
out/terratest-sentinel-source_toolkit_fluxcd_io_helm_repository_v1_manifest_test.go: out/install-sentinel terratest/source_toolkit_fluxcd_io_v1/source_toolkit_fluxcd_io_helm_repository_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_source_toolkit_fluxcd_io_helm_repository_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/source_toolkit_fluxcd_io_v1/source_toolkit_fluxcd_io_helm_repository_v1_manifest_test.go
	touch $@
out/terratest-sentinel-source_toolkit_fluxcd_io_bucket_v1beta1_manifest_test.go: out/install-sentinel terratest/source_toolkit_fluxcd_io_v1beta1/source_toolkit_fluxcd_io_bucket_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_source_toolkit_fluxcd_io_bucket_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/source_toolkit_fluxcd_io_v1beta1/source_toolkit_fluxcd_io_bucket_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-source_toolkit_fluxcd_io_git_repository_v1beta1_manifest_test.go: out/install-sentinel terratest/source_toolkit_fluxcd_io_v1beta1/source_toolkit_fluxcd_io_git_repository_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_source_toolkit_fluxcd_io_git_repository_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/source_toolkit_fluxcd_io_v1beta1/source_toolkit_fluxcd_io_git_repository_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-source_toolkit_fluxcd_io_helm_chart_v1beta1_manifest_test.go: out/install-sentinel terratest/source_toolkit_fluxcd_io_v1beta1/source_toolkit_fluxcd_io_helm_chart_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_source_toolkit_fluxcd_io_helm_chart_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/source_toolkit_fluxcd_io_v1beta1/source_toolkit_fluxcd_io_helm_chart_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-source_toolkit_fluxcd_io_helm_repository_v1beta1_manifest_test.go: out/install-sentinel terratest/source_toolkit_fluxcd_io_v1beta1/source_toolkit_fluxcd_io_helm_repository_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_source_toolkit_fluxcd_io_helm_repository_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/source_toolkit_fluxcd_io_v1beta1/source_toolkit_fluxcd_io_helm_repository_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-source_toolkit_fluxcd_io_bucket_v1beta2_manifest_test.go: out/install-sentinel terratest/source_toolkit_fluxcd_io_v1beta2/source_toolkit_fluxcd_io_bucket_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_source_toolkit_fluxcd_io_bucket_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/source_toolkit_fluxcd_io_v1beta2/source_toolkit_fluxcd_io_bucket_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-source_toolkit_fluxcd_io_git_repository_v1beta2_manifest_test.go: out/install-sentinel terratest/source_toolkit_fluxcd_io_v1beta2/source_toolkit_fluxcd_io_git_repository_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_source_toolkit_fluxcd_io_git_repository_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/source_toolkit_fluxcd_io_v1beta2/source_toolkit_fluxcd_io_git_repository_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-source_toolkit_fluxcd_io_helm_chart_v1beta2_manifest_test.go: out/install-sentinel terratest/source_toolkit_fluxcd_io_v1beta2/source_toolkit_fluxcd_io_helm_chart_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_source_toolkit_fluxcd_io_helm_chart_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/source_toolkit_fluxcd_io_v1beta2/source_toolkit_fluxcd_io_helm_chart_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-source_toolkit_fluxcd_io_helm_repository_v1beta2_manifest_test.go: out/install-sentinel terratest/source_toolkit_fluxcd_io_v1beta2/source_toolkit_fluxcd_io_helm_repository_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_source_toolkit_fluxcd_io_helm_repository_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/source_toolkit_fluxcd_io_v1beta2/source_toolkit_fluxcd_io_helm_repository_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-source_toolkit_fluxcd_io_oci_repository_v1beta2_manifest_test.go: out/install-sentinel terratest/source_toolkit_fluxcd_io_v1beta2/source_toolkit_fluxcd_io_oci_repository_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_source_toolkit_fluxcd_io_oci_repository_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/source_toolkit_fluxcd_io_v1beta2/source_toolkit_fluxcd_io_oci_repository_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-sparkoperator_k8s_io_scheduled_spark_application_v1beta2_manifest_test.go: out/install-sentinel terratest/sparkoperator_k8s_io_v1beta2/sparkoperator_k8s_io_scheduled_spark_application_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_sparkoperator_k8s_io_scheduled_spark_application_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/sparkoperator_k8s_io_v1beta2/sparkoperator_k8s_io_scheduled_spark_application_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-sparkoperator_k8s_io_spark_application_v1beta2_manifest_test.go: out/install-sentinel terratest/sparkoperator_k8s_io_v1beta2/sparkoperator_k8s_io_spark_application_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_sparkoperator_k8s_io_spark_application_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/sparkoperator_k8s_io_v1beta2/sparkoperator_k8s_io_spark_application_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-spv_no_azure_key_vault_secret_v1_manifest_test.go: out/install-sentinel terratest/spv_no_v1/spv_no_azure_key_vault_secret_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_spv_no_azure_key_vault_secret_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/spv_no_v1/spv_no_azure_key_vault_secret_v1_manifest_test.go
	touch $@
out/terratest-sentinel-spv_no_azure_key_vault_identity_v1alpha1_manifest_test.go: out/install-sentinel terratest/spv_no_v1alpha1/spv_no_azure_key_vault_identity_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_spv_no_azure_key_vault_identity_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/spv_no_v1alpha1/spv_no_azure_key_vault_identity_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-spv_no_azure_key_vault_secret_v1alpha1_manifest_test.go: out/install-sentinel terratest/spv_no_v1alpha1/spv_no_azure_key_vault_secret_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_spv_no_azure_key_vault_secret_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/spv_no_v1alpha1/spv_no_azure_key_vault_secret_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-spv_no_azure_managed_identity_v1alpha1_manifest_test.go: out/install-sentinel terratest/spv_no_v1alpha1/spv_no_azure_managed_identity_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_spv_no_azure_managed_identity_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/spv_no_v1alpha1/spv_no_azure_managed_identity_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-spv_no_azure_key_vault_secret_v2alpha1_manifest_test.go: out/install-sentinel terratest/spv_no_v2alpha1/spv_no_azure_key_vault_secret_v2alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_spv_no_azure_key_vault_secret_v2alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/spv_no_v2alpha1/spv_no_azure_key_vault_secret_v2alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-spv_no_azure_key_vault_secret_v2beta1_manifest_test.go: out/install-sentinel terratest/spv_no_v2beta1/spv_no_azure_key_vault_secret_v2beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_spv_no_azure_key_vault_secret_v2beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/spv_no_v2beta1/spv_no_azure_key_vault_secret_v2beta1_manifest_test.go
	touch $@
out/terratest-sentinel-sqs_services_k8s_aws_queue_v1alpha1_manifest_test.go: out/install-sentinel terratest/sqs_services_k8s_aws_v1alpha1/sqs_services_k8s_aws_queue_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_sqs_services_k8s_aws_queue_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/sqs_services_k8s_aws_v1alpha1/sqs_services_k8s_aws_queue_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-storage_k8s_io_csi_driver_v1_manifest_test.go: out/install-sentinel terratest/storage_k8s_io_v1/storage_k8s_io_csi_driver_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_storage_k8s_io_csi_driver_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/storage_k8s_io_v1/storage_k8s_io_csi_driver_v1_manifest_test.go
	touch $@
out/terratest-sentinel-storage_k8s_io_csi_node_v1_manifest_test.go: out/install-sentinel terratest/storage_k8s_io_v1/storage_k8s_io_csi_node_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_storage_k8s_io_csi_node_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/storage_k8s_io_v1/storage_k8s_io_csi_node_v1_manifest_test.go
	touch $@
out/terratest-sentinel-storage_k8s_io_storage_class_v1_manifest_test.go: out/install-sentinel terratest/storage_k8s_io_v1/storage_k8s_io_storage_class_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_storage_k8s_io_storage_class_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/storage_k8s_io_v1/storage_k8s_io_storage_class_v1_manifest_test.go
	touch $@
out/terratest-sentinel-storage_k8s_io_volume_attachment_v1_manifest_test.go: out/install-sentinel terratest/storage_k8s_io_v1/storage_k8s_io_volume_attachment_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_storage_k8s_io_volume_attachment_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/storage_k8s_io_v1/storage_k8s_io_volume_attachment_v1_manifest_test.go
	touch $@
out/terratest-sentinel-storage_kubeblocks_io_storage_provider_v1alpha1_manifest_test.go: out/install-sentinel terratest/storage_kubeblocks_io_v1alpha1/storage_kubeblocks_io_storage_provider_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_storage_kubeblocks_io_storage_provider_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/storage_kubeblocks_io_v1alpha1/storage_kubeblocks_io_storage_provider_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-storageos_com_storage_os_cluster_v1_manifest_test.go: out/install-sentinel terratest/storageos_com_v1/storageos_com_storage_os_cluster_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_storageos_com_storage_os_cluster_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/storageos_com_v1/storageos_com_storage_os_cluster_v1_manifest_test.go
	touch $@
out/terratest-sentinel-sts_min_io_policy_binding_v1alpha1_manifest_test.go: out/install-sentinel terratest/sts_min_io_v1alpha1/sts_min_io_policy_binding_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_sts_min_io_policy_binding_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/sts_min_io_v1alpha1/sts_min_io_policy_binding_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-sts_min_io_policy_binding_v1beta1_manifest_test.go: out/install-sentinel terratest/sts_min_io_v1beta1/sts_min_io_policy_binding_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_sts_min_io_policy_binding_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/sts_min_io_v1beta1/sts_min_io_policy_binding_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-stunner_l7mp_io_dataplane_v1_manifest_test.go: out/install-sentinel terratest/stunner_l7mp_io_v1/stunner_l7mp_io_dataplane_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_stunner_l7mp_io_dataplane_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/stunner_l7mp_io_v1/stunner_l7mp_io_dataplane_v1_manifest_test.go
	touch $@
out/terratest-sentinel-stunner_l7mp_io_gateway_config_v1_manifest_test.go: out/install-sentinel terratest/stunner_l7mp_io_v1/stunner_l7mp_io_gateway_config_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_stunner_l7mp_io_gateway_config_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/stunner_l7mp_io_v1/stunner_l7mp_io_gateway_config_v1_manifest_test.go
	touch $@
out/terratest-sentinel-stunner_l7mp_io_static_service_v1_manifest_test.go: out/install-sentinel terratest/stunner_l7mp_io_v1/stunner_l7mp_io_static_service_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_stunner_l7mp_io_static_service_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/stunner_l7mp_io_v1/stunner_l7mp_io_static_service_v1_manifest_test.go
	touch $@
out/terratest-sentinel-stunner_l7mp_io_udp_route_v1_manifest_test.go: out/install-sentinel terratest/stunner_l7mp_io_v1/stunner_l7mp_io_udp_route_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_stunner_l7mp_io_udp_route_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/stunner_l7mp_io_v1/stunner_l7mp_io_udp_route_v1_manifest_test.go
	touch $@
out/terratest-sentinel-stunner_l7mp_io_dataplane_v1alpha1_manifest_test.go: out/install-sentinel terratest/stunner_l7mp_io_v1alpha1/stunner_l7mp_io_dataplane_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_stunner_l7mp_io_dataplane_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/stunner_l7mp_io_v1alpha1/stunner_l7mp_io_dataplane_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-stunner_l7mp_io_gateway_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/stunner_l7mp_io_v1alpha1/stunner_l7mp_io_gateway_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_stunner_l7mp_io_gateway_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/stunner_l7mp_io_v1alpha1/stunner_l7mp_io_gateway_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-stunner_l7mp_io_static_service_v1alpha1_manifest_test.go: out/install-sentinel terratest/stunner_l7mp_io_v1alpha1/stunner_l7mp_io_static_service_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_stunner_l7mp_io_static_service_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/stunner_l7mp_io_v1alpha1/stunner_l7mp_io_static_service_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-submariner_io_broker_v1alpha1_manifest_test.go: out/install-sentinel terratest/submariner_io_v1alpha1/submariner_io_broker_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_submariner_io_broker_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/submariner_io_v1alpha1/submariner_io_broker_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-submariner_io_service_discovery_v1alpha1_manifest_test.go: out/install-sentinel terratest/submariner_io_v1alpha1/submariner_io_service_discovery_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_submariner_io_service_discovery_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/submariner_io_v1alpha1/submariner_io_service_discovery_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-submariner_io_submariner_v1alpha1_manifest_test.go: out/install-sentinel terratest/submariner_io_v1alpha1/submariner_io_submariner_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_submariner_io_submariner_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/submariner_io_v1alpha1/submariner_io_submariner_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-telemetry_istio_io_telemetry_v1_manifest_test.go: out/install-sentinel terratest/telemetry_istio_io_v1/telemetry_istio_io_telemetry_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_telemetry_istio_io_telemetry_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/telemetry_istio_io_v1/telemetry_istio_io_telemetry_v1_manifest_test.go
	touch $@
out/terratest-sentinel-telemetry_istio_io_telemetry_v1alpha1_manifest_test.go: out/install-sentinel terratest/telemetry_istio_io_v1alpha1/telemetry_istio_io_telemetry_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_telemetry_istio_io_telemetry_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/telemetry_istio_io_v1alpha1/telemetry_istio_io_telemetry_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-templates_gatekeeper_sh_constraint_template_v1_manifest_test.go: out/install-sentinel terratest/templates_gatekeeper_sh_v1/templates_gatekeeper_sh_constraint_template_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_templates_gatekeeper_sh_constraint_template_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/templates_gatekeeper_sh_v1/templates_gatekeeper_sh_constraint_template_v1_manifest_test.go
	touch $@
out/terratest-sentinel-templates_gatekeeper_sh_constraint_template_v1alpha1_manifest_test.go: out/install-sentinel terratest/templates_gatekeeper_sh_v1alpha1/templates_gatekeeper_sh_constraint_template_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_templates_gatekeeper_sh_constraint_template_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/templates_gatekeeper_sh_v1alpha1/templates_gatekeeper_sh_constraint_template_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-templates_gatekeeper_sh_constraint_template_v1beta1_manifest_test.go: out/install-sentinel terratest/templates_gatekeeper_sh_v1beta1/templates_gatekeeper_sh_constraint_template_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_templates_gatekeeper_sh_constraint_template_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/templates_gatekeeper_sh_v1beta1/templates_gatekeeper_sh_constraint_template_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-tempo_grafana_com_tempo_monolithic_v1alpha1_manifest_test.go: out/install-sentinel terratest/tempo_grafana_com_v1alpha1/tempo_grafana_com_tempo_monolithic_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_tempo_grafana_com_tempo_monolithic_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tempo_grafana_com_v1alpha1/tempo_grafana_com_tempo_monolithic_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-tempo_grafana_com_tempo_stack_v1alpha1_manifest_test.go: out/install-sentinel terratest/tempo_grafana_com_v1alpha1/tempo_grafana_com_tempo_stack_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_tempo_grafana_com_tempo_stack_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tempo_grafana_com_v1alpha1/tempo_grafana_com_tempo_stack_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-temporal_io_temporal_cluster_client_v1beta1_manifest_test.go: out/install-sentinel terratest/temporal_io_v1beta1/temporal_io_temporal_cluster_client_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_temporal_io_temporal_cluster_client_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/temporal_io_v1beta1/temporal_io_temporal_cluster_client_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-temporal_io_temporal_cluster_v1beta1_manifest_test.go: out/install-sentinel terratest/temporal_io_v1beta1/temporal_io_temporal_cluster_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_temporal_io_temporal_cluster_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/temporal_io_v1beta1/temporal_io_temporal_cluster_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-temporal_io_temporal_namespace_v1beta1_manifest_test.go: out/install-sentinel terratest/temporal_io_v1beta1/temporal_io_temporal_namespace_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_temporal_io_temporal_namespace_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/temporal_io_v1beta1/temporal_io_temporal_namespace_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-temporal_io_temporal_worker_process_v1beta1_manifest_test.go: out/install-sentinel terratest/temporal_io_v1beta1/temporal_io_temporal_worker_process_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_temporal_io_temporal_worker_process_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/temporal_io_v1beta1/temporal_io_temporal_worker_process_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-tests_testkube_io_script_v1_manifest_test.go: out/install-sentinel terratest/tests_testkube_io_v1/tests_testkube_io_script_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_tests_testkube_io_script_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tests_testkube_io_v1/tests_testkube_io_script_v1_manifest_test.go
	touch $@
out/terratest-sentinel-tests_testkube_io_test_execution_v1_manifest_test.go: out/install-sentinel terratest/tests_testkube_io_v1/tests_testkube_io_test_execution_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_tests_testkube_io_test_execution_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tests_testkube_io_v1/tests_testkube_io_test_execution_v1_manifest_test.go
	touch $@
out/terratest-sentinel-tests_testkube_io_test_source_v1_manifest_test.go: out/install-sentinel terratest/tests_testkube_io_v1/tests_testkube_io_test_source_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_tests_testkube_io_test_source_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tests_testkube_io_v1/tests_testkube_io_test_source_v1_manifest_test.go
	touch $@
out/terratest-sentinel-tests_testkube_io_test_suite_execution_v1_manifest_test.go: out/install-sentinel terratest/tests_testkube_io_v1/tests_testkube_io_test_suite_execution_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_tests_testkube_io_test_suite_execution_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tests_testkube_io_v1/tests_testkube_io_test_suite_execution_v1_manifest_test.go
	touch $@
out/terratest-sentinel-tests_testkube_io_test_suite_v1_manifest_test.go: out/install-sentinel terratest/tests_testkube_io_v1/tests_testkube_io_test_suite_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_tests_testkube_io_test_suite_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tests_testkube_io_v1/tests_testkube_io_test_suite_v1_manifest_test.go
	touch $@
out/terratest-sentinel-tests_testkube_io_test_trigger_v1_manifest_test.go: out/install-sentinel terratest/tests_testkube_io_v1/tests_testkube_io_test_trigger_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_tests_testkube_io_test_trigger_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tests_testkube_io_v1/tests_testkube_io_test_trigger_v1_manifest_test.go
	touch $@
out/terratest-sentinel-tests_testkube_io_test_v1_manifest_test.go: out/install-sentinel terratest/tests_testkube_io_v1/tests_testkube_io_test_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_tests_testkube_io_test_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tests_testkube_io_v1/tests_testkube_io_test_v1_manifest_test.go
	touch $@
out/terratest-sentinel-tests_testkube_io_script_v2_manifest_test.go: out/install-sentinel terratest/tests_testkube_io_v2/tests_testkube_io_script_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_tests_testkube_io_script_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tests_testkube_io_v2/tests_testkube_io_script_v2_manifest_test.go
	touch $@
out/terratest-sentinel-tests_testkube_io_test_suite_v2_manifest_test.go: out/install-sentinel terratest/tests_testkube_io_v2/tests_testkube_io_test_suite_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_tests_testkube_io_test_suite_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tests_testkube_io_v2/tests_testkube_io_test_suite_v2_manifest_test.go
	touch $@
out/terratest-sentinel-tests_testkube_io_test_v2_manifest_test.go: out/install-sentinel terratest/tests_testkube_io_v2/tests_testkube_io_test_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_tests_testkube_io_test_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tests_testkube_io_v2/tests_testkube_io_test_v2_manifest_test.go
	touch $@
out/terratest-sentinel-tests_testkube_io_test_suite_v3_manifest_test.go: out/install-sentinel terratest/tests_testkube_io_v3/tests_testkube_io_test_suite_v3_manifest_test.go $(shell find ./examples/data-sources/k8s_tests_testkube_io_test_suite_v3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tests_testkube_io_v3/tests_testkube_io_test_suite_v3_manifest_test.go
	touch $@
out/terratest-sentinel-tests_testkube_io_test_v3_manifest_test.go: out/install-sentinel terratest/tests_testkube_io_v3/tests_testkube_io_test_v3_manifest_test.go $(shell find ./examples/data-sources/k8s_tests_testkube_io_test_v3_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tests_testkube_io_v3/tests_testkube_io_test_v3_manifest_test.go
	touch $@
out/terratest-sentinel-tf_tungsten_io_analytics_alarm_v1alpha1_manifest_test.go: out/install-sentinel terratest/tf_tungsten_io_v1alpha1/tf_tungsten_io_analytics_alarm_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_tf_tungsten_io_analytics_alarm_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tf_tungsten_io_v1alpha1/tf_tungsten_io_analytics_alarm_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-tf_tungsten_io_analytics_snmp_v1alpha1_manifest_test.go: out/install-sentinel terratest/tf_tungsten_io_v1alpha1/tf_tungsten_io_analytics_snmp_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_tf_tungsten_io_analytics_snmp_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tf_tungsten_io_v1alpha1/tf_tungsten_io_analytics_snmp_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-tf_tungsten_io_analytics_v1alpha1_manifest_test.go: out/install-sentinel terratest/tf_tungsten_io_v1alpha1/tf_tungsten_io_analytics_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_tf_tungsten_io_analytics_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tf_tungsten_io_v1alpha1/tf_tungsten_io_analytics_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-tf_tungsten_io_cassandra_v1alpha1_manifest_test.go: out/install-sentinel terratest/tf_tungsten_io_v1alpha1/tf_tungsten_io_cassandra_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_tf_tungsten_io_cassandra_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tf_tungsten_io_v1alpha1/tf_tungsten_io_cassandra_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-tf_tungsten_io_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/tf_tungsten_io_v1alpha1/tf_tungsten_io_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_tf_tungsten_io_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tf_tungsten_io_v1alpha1/tf_tungsten_io_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-tf_tungsten_io_control_v1alpha1_manifest_test.go: out/install-sentinel terratest/tf_tungsten_io_v1alpha1/tf_tungsten_io_control_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_tf_tungsten_io_control_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tf_tungsten_io_v1alpha1/tf_tungsten_io_control_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-tf_tungsten_io_kubemanager_v1alpha1_manifest_test.go: out/install-sentinel terratest/tf_tungsten_io_v1alpha1/tf_tungsten_io_kubemanager_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_tf_tungsten_io_kubemanager_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tf_tungsten_io_v1alpha1/tf_tungsten_io_kubemanager_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-tf_tungsten_io_manager_v1alpha1_manifest_test.go: out/install-sentinel terratest/tf_tungsten_io_v1alpha1/tf_tungsten_io_manager_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_tf_tungsten_io_manager_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tf_tungsten_io_v1alpha1/tf_tungsten_io_manager_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-tf_tungsten_io_query_engine_v1alpha1_manifest_test.go: out/install-sentinel terratest/tf_tungsten_io_v1alpha1/tf_tungsten_io_query_engine_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_tf_tungsten_io_query_engine_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tf_tungsten_io_v1alpha1/tf_tungsten_io_query_engine_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-tf_tungsten_io_rabbitmq_v1alpha1_manifest_test.go: out/install-sentinel terratest/tf_tungsten_io_v1alpha1/tf_tungsten_io_rabbitmq_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_tf_tungsten_io_rabbitmq_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tf_tungsten_io_v1alpha1/tf_tungsten_io_rabbitmq_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-tf_tungsten_io_redis_v1alpha1_manifest_test.go: out/install-sentinel terratest/tf_tungsten_io_v1alpha1/tf_tungsten_io_redis_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_tf_tungsten_io_redis_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tf_tungsten_io_v1alpha1/tf_tungsten_io_redis_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-tf_tungsten_io_vrouter_v1alpha1_manifest_test.go: out/install-sentinel terratest/tf_tungsten_io_v1alpha1/tf_tungsten_io_vrouter_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_tf_tungsten_io_vrouter_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tf_tungsten_io_v1alpha1/tf_tungsten_io_vrouter_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-tf_tungsten_io_webui_v1alpha1_manifest_test.go: out/install-sentinel terratest/tf_tungsten_io_v1alpha1/tf_tungsten_io_webui_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_tf_tungsten_io_webui_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tf_tungsten_io_v1alpha1/tf_tungsten_io_webui_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-tf_tungsten_io_zookeeper_v1alpha1_manifest_test.go: out/install-sentinel terratest/tf_tungsten_io_v1alpha1/tf_tungsten_io_zookeeper_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_tf_tungsten_io_zookeeper_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tf_tungsten_io_v1alpha1/tf_tungsten_io_zookeeper_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-theketch_io_app_v1beta1_manifest_test.go: out/install-sentinel terratest/theketch_io_v1beta1/theketch_io_app_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_theketch_io_app_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/theketch_io_v1beta1/theketch_io_app_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-theketch_io_job_v1beta1_manifest_test.go: out/install-sentinel terratest/theketch_io_v1beta1/theketch_io_job_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_theketch_io_job_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/theketch_io_v1beta1/theketch_io_job_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-tinkerbell_org_hardware_v1alpha1_manifest_test.go: out/install-sentinel terratest/tinkerbell_org_v1alpha1/tinkerbell_org_hardware_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_tinkerbell_org_hardware_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tinkerbell_org_v1alpha1/tinkerbell_org_hardware_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-tinkerbell_org_osie_v1alpha1_manifest_test.go: out/install-sentinel terratest/tinkerbell_org_v1alpha1/tinkerbell_org_osie_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_tinkerbell_org_osie_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tinkerbell_org_v1alpha1/tinkerbell_org_osie_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-tinkerbell_org_stack_v1alpha1_manifest_test.go: out/install-sentinel terratest/tinkerbell_org_v1alpha1/tinkerbell_org_stack_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_tinkerbell_org_stack_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tinkerbell_org_v1alpha1/tinkerbell_org_stack_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-tinkerbell_org_template_v1alpha1_manifest_test.go: out/install-sentinel terratest/tinkerbell_org_v1alpha1/tinkerbell_org_template_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_tinkerbell_org_template_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tinkerbell_org_v1alpha1/tinkerbell_org_template_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-tinkerbell_org_workflow_v1alpha1_manifest_test.go: out/install-sentinel terratest/tinkerbell_org_v1alpha1/tinkerbell_org_workflow_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_tinkerbell_org_workflow_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tinkerbell_org_v1alpha1/tinkerbell_org_workflow_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-tinkerbell_org_hardware_v1alpha2_manifest_test.go: out/install-sentinel terratest/tinkerbell_org_v1alpha2/tinkerbell_org_hardware_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_tinkerbell_org_hardware_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tinkerbell_org_v1alpha2/tinkerbell_org_hardware_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-tinkerbell_org_osie_v1alpha2_manifest_test.go: out/install-sentinel terratest/tinkerbell_org_v1alpha2/tinkerbell_org_osie_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_tinkerbell_org_osie_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tinkerbell_org_v1alpha2/tinkerbell_org_osie_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-tinkerbell_org_template_v1alpha2_manifest_test.go: out/install-sentinel terratest/tinkerbell_org_v1alpha2/tinkerbell_org_template_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_tinkerbell_org_template_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tinkerbell_org_v1alpha2/tinkerbell_org_template_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-tinkerbell_org_workflow_v1alpha2_manifest_test.go: out/install-sentinel terratest/tinkerbell_org_v1alpha2/tinkerbell_org_workflow_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_tinkerbell_org_workflow_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/tinkerbell_org_v1alpha2/tinkerbell_org_workflow_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-topology_node_k8s_io_node_resource_topology_v1alpha1_manifest_test.go: out/install-sentinel terratest/topology_node_k8s_io_v1alpha1/topology_node_k8s_io_node_resource_topology_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_topology_node_k8s_io_node_resource_topology_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/topology_node_k8s_io_v1alpha1/topology_node_k8s_io_node_resource_topology_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-topolvm_cybozu_com_logical_volume_v1_manifest_test.go: out/install-sentinel terratest/topolvm_cybozu_com_v1/topolvm_cybozu_com_logical_volume_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_topolvm_cybozu_com_logical_volume_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/topolvm_cybozu_com_v1/topolvm_cybozu_com_logical_volume_v1_manifest_test.go
	touch $@
out/terratest-sentinel-topolvm_cybozu_com_topolvm_cluster_v2_manifest_test.go: out/install-sentinel terratest/topolvm_cybozu_com_v2/topolvm_cybozu_com_topolvm_cluster_v2_manifest_test.go $(shell find ./examples/data-sources/k8s_topolvm_cybozu_com_topolvm_cluster_v2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/topolvm_cybozu_com_v2/topolvm_cybozu_com_topolvm_cluster_v2_manifest_test.go
	touch $@
out/terratest-sentinel-traefik_io_ingress_route_tcp_v1alpha1_manifest_test.go: out/install-sentinel terratest/traefik_io_v1alpha1/traefik_io_ingress_route_tcp_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_traefik_io_ingress_route_tcp_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/traefik_io_v1alpha1/traefik_io_ingress_route_tcp_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-traefik_io_ingress_route_udp_v1alpha1_manifest_test.go: out/install-sentinel terratest/traefik_io_v1alpha1/traefik_io_ingress_route_udp_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_traefik_io_ingress_route_udp_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/traefik_io_v1alpha1/traefik_io_ingress_route_udp_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-traefik_io_ingress_route_v1alpha1_manifest_test.go: out/install-sentinel terratest/traefik_io_v1alpha1/traefik_io_ingress_route_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_traefik_io_ingress_route_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/traefik_io_v1alpha1/traefik_io_ingress_route_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-traefik_io_middleware_tcp_v1alpha1_manifest_test.go: out/install-sentinel terratest/traefik_io_v1alpha1/traefik_io_middleware_tcp_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_traefik_io_middleware_tcp_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/traefik_io_v1alpha1/traefik_io_middleware_tcp_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-traefik_io_middleware_v1alpha1_manifest_test.go: out/install-sentinel terratest/traefik_io_v1alpha1/traefik_io_middleware_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_traefik_io_middleware_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/traefik_io_v1alpha1/traefik_io_middleware_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-traefik_io_servers_transport_tcp_v1alpha1_manifest_test.go: out/install-sentinel terratest/traefik_io_v1alpha1/traefik_io_servers_transport_tcp_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_traefik_io_servers_transport_tcp_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/traefik_io_v1alpha1/traefik_io_servers_transport_tcp_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-traefik_io_servers_transport_v1alpha1_manifest_test.go: out/install-sentinel terratest/traefik_io_v1alpha1/traefik_io_servers_transport_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_traefik_io_servers_transport_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/traefik_io_v1alpha1/traefik_io_servers_transport_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-traefik_io_tls_option_v1alpha1_manifest_test.go: out/install-sentinel terratest/traefik_io_v1alpha1/traefik_io_tls_option_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_traefik_io_tls_option_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/traefik_io_v1alpha1/traefik_io_tls_option_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-traefik_io_tls_store_v1alpha1_manifest_test.go: out/install-sentinel terratest/traefik_io_v1alpha1/traefik_io_tls_store_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_traefik_io_tls_store_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/traefik_io_v1alpha1/traefik_io_tls_store_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-traefik_io_traefik_service_v1alpha1_manifest_test.go: out/install-sentinel terratest/traefik_io_v1alpha1/traefik_io_traefik_service_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_traefik_io_traefik_service_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/traefik_io_v1alpha1/traefik_io_traefik_service_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-training_kubedl_io_elastic_dl_job_v1alpha1_manifest_test.go: out/install-sentinel terratest/training_kubedl_io_v1alpha1/training_kubedl_io_elastic_dl_job_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_training_kubedl_io_elastic_dl_job_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/training_kubedl_io_v1alpha1/training_kubedl_io_elastic_dl_job_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-training_kubedl_io_mars_job_v1alpha1_manifest_test.go: out/install-sentinel terratest/training_kubedl_io_v1alpha1/training_kubedl_io_mars_job_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_training_kubedl_io_mars_job_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/training_kubedl_io_v1alpha1/training_kubedl_io_mars_job_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-training_kubedl_io_mpi_job_v1alpha1_manifest_test.go: out/install-sentinel terratest/training_kubedl_io_v1alpha1/training_kubedl_io_mpi_job_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_training_kubedl_io_mpi_job_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/training_kubedl_io_v1alpha1/training_kubedl_io_mpi_job_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-training_kubedl_io_py_torch_job_v1alpha1_manifest_test.go: out/install-sentinel terratest/training_kubedl_io_v1alpha1/training_kubedl_io_py_torch_job_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_training_kubedl_io_py_torch_job_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/training_kubedl_io_v1alpha1/training_kubedl_io_py_torch_job_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-training_kubedl_io_tf_job_v1alpha1_manifest_test.go: out/install-sentinel terratest/training_kubedl_io_v1alpha1/training_kubedl_io_tf_job_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_training_kubedl_io_tf_job_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/training_kubedl_io_v1alpha1/training_kubedl_io_tf_job_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-training_kubedl_io_xdl_job_v1alpha1_manifest_test.go: out/install-sentinel terratest/training_kubedl_io_v1alpha1/training_kubedl_io_xdl_job_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_training_kubedl_io_xdl_job_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/training_kubedl_io_v1alpha1/training_kubedl_io_xdl_job_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-training_kubedl_io_xg_boost_job_v1alpha1_manifest_test.go: out/install-sentinel terratest/training_kubedl_io_v1alpha1/training_kubedl_io_xg_boost_job_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_training_kubedl_io_xg_boost_job_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/training_kubedl_io_v1alpha1/training_kubedl_io_xg_boost_job_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-trust_cert_manager_io_bundle_v1alpha1_manifest_test.go: out/install-sentinel terratest/trust_cert_manager_io_v1alpha1/trust_cert_manager_io_bundle_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_trust_cert_manager_io_bundle_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/trust_cert_manager_io_v1alpha1/trust_cert_manager_io_bundle_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-upgrade_cattle_io_plan_v1_manifest_test.go: out/install-sentinel terratest/upgrade_cattle_io_v1/upgrade_cattle_io_plan_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_upgrade_cattle_io_plan_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/upgrade_cattle_io_v1/upgrade_cattle_io_plan_v1_manifest_test.go
	touch $@
out/terratest-sentinel-upgrade_managed_openshift_io_upgrade_config_v1alpha1_manifest_test.go: out/install-sentinel terratest/upgrade_managed_openshift_io_v1alpha1/upgrade_managed_openshift_io_upgrade_config_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_upgrade_managed_openshift_io_upgrade_config_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/upgrade_managed_openshift_io_v1alpha1/upgrade_managed_openshift_io_upgrade_config_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-velero_io_backup_repository_v1_manifest_test.go: out/install-sentinel terratest/velero_io_v1/velero_io_backup_repository_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_velero_io_backup_repository_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/velero_io_v1/velero_io_backup_repository_v1_manifest_test.go
	touch $@
out/terratest-sentinel-velero_io_backup_storage_location_v1_manifest_test.go: out/install-sentinel terratest/velero_io_v1/velero_io_backup_storage_location_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_velero_io_backup_storage_location_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/velero_io_v1/velero_io_backup_storage_location_v1_manifest_test.go
	touch $@
out/terratest-sentinel-velero_io_backup_v1_manifest_test.go: out/install-sentinel terratest/velero_io_v1/velero_io_backup_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_velero_io_backup_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/velero_io_v1/velero_io_backup_v1_manifest_test.go
	touch $@
out/terratest-sentinel-velero_io_delete_backup_request_v1_manifest_test.go: out/install-sentinel terratest/velero_io_v1/velero_io_delete_backup_request_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_velero_io_delete_backup_request_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/velero_io_v1/velero_io_delete_backup_request_v1_manifest_test.go
	touch $@
out/terratest-sentinel-velero_io_download_request_v1_manifest_test.go: out/install-sentinel terratest/velero_io_v1/velero_io_download_request_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_velero_io_download_request_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/velero_io_v1/velero_io_download_request_v1_manifest_test.go
	touch $@
out/terratest-sentinel-velero_io_pod_volume_backup_v1_manifest_test.go: out/install-sentinel terratest/velero_io_v1/velero_io_pod_volume_backup_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_velero_io_pod_volume_backup_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/velero_io_v1/velero_io_pod_volume_backup_v1_manifest_test.go
	touch $@
out/terratest-sentinel-velero_io_pod_volume_restore_v1_manifest_test.go: out/install-sentinel terratest/velero_io_v1/velero_io_pod_volume_restore_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_velero_io_pod_volume_restore_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/velero_io_v1/velero_io_pod_volume_restore_v1_manifest_test.go
	touch $@
out/terratest-sentinel-velero_io_restore_v1_manifest_test.go: out/install-sentinel terratest/velero_io_v1/velero_io_restore_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_velero_io_restore_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/velero_io_v1/velero_io_restore_v1_manifest_test.go
	touch $@
out/terratest-sentinel-velero_io_schedule_v1_manifest_test.go: out/install-sentinel terratest/velero_io_v1/velero_io_schedule_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_velero_io_schedule_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/velero_io_v1/velero_io_schedule_v1_manifest_test.go
	touch $@
out/terratest-sentinel-velero_io_server_status_request_v1_manifest_test.go: out/install-sentinel terratest/velero_io_v1/velero_io_server_status_request_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_velero_io_server_status_request_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/velero_io_v1/velero_io_server_status_request_v1_manifest_test.go
	touch $@
out/terratest-sentinel-velero_io_volume_snapshot_location_v1_manifest_test.go: out/install-sentinel terratest/velero_io_v1/velero_io_volume_snapshot_location_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_velero_io_volume_snapshot_location_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/velero_io_v1/velero_io_volume_snapshot_location_v1_manifest_test.go
	touch $@
out/terratest-sentinel-velero_io_data_download_v2alpha1_manifest_test.go: out/install-sentinel terratest/velero_io_v2alpha1/velero_io_data_download_v2alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_velero_io_data_download_v2alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/velero_io_v2alpha1/velero_io_data_download_v2alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-velero_io_data_upload_v2alpha1_manifest_test.go: out/install-sentinel terratest/velero_io_v2alpha1/velero_io_data_upload_v2alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_velero_io_data_upload_v2alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/velero_io_v2alpha1/velero_io_data_upload_v2alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-virt_virtink_smartx_com_virtual_machine_migration_v1alpha1_manifest_test.go: out/install-sentinel terratest/virt_virtink_smartx_com_v1alpha1/virt_virtink_smartx_com_virtual_machine_migration_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_virt_virtink_smartx_com_virtual_machine_migration_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/virt_virtink_smartx_com_v1alpha1/virt_virtink_smartx_com_virtual_machine_migration_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-virt_virtink_smartx_com_virtual_machine_v1alpha1_manifest_test.go: out/install-sentinel terratest/virt_virtink_smartx_com_v1alpha1/virt_virtink_smartx_com_virtual_machine_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_virt_virtink_smartx_com_virtual_machine_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/virt_virtink_smartx_com_v1alpha1/virt_virtink_smartx_com_virtual_machine_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-volsync_backube_replication_destination_v1alpha1_manifest_test.go: out/install-sentinel terratest/volsync_backube_v1alpha1/volsync_backube_replication_destination_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_volsync_backube_replication_destination_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/volsync_backube_v1alpha1/volsync_backube_replication_destination_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-volsync_backube_replication_source_v1alpha1_manifest_test.go: out/install-sentinel terratest/volsync_backube_v1alpha1/volsync_backube_replication_source_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_volsync_backube_replication_source_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/volsync_backube_v1alpha1/volsync_backube_replication_source_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-vpcresources_k8s_aws_cni_node_v1alpha1_manifest_test.go: out/install-sentinel terratest/vpcresources_k8s_aws_v1alpha1/vpcresources_k8s_aws_cni_node_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_vpcresources_k8s_aws_cni_node_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/vpcresources_k8s_aws_v1alpha1/vpcresources_k8s_aws_cni_node_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-vpcresources_k8s_aws_security_group_policy_v1beta1_manifest_test.go: out/install-sentinel terratest/vpcresources_k8s_aws_v1beta1/vpcresources_k8s_aws_security_group_policy_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_vpcresources_k8s_aws_security_group_policy_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/vpcresources_k8s_aws_v1beta1/vpcresources_k8s_aws_security_group_policy_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-wgpolicyk8s_io_cluster_policy_report_v1alpha1_manifest_test.go: out/install-sentinel terratest/wgpolicyk8s_io_v1alpha1/wgpolicyk8s_io_cluster_policy_report_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_wgpolicyk8s_io_cluster_policy_report_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/wgpolicyk8s_io_v1alpha1/wgpolicyk8s_io_cluster_policy_report_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-wgpolicyk8s_io_policy_report_v1alpha1_manifest_test.go: out/install-sentinel terratest/wgpolicyk8s_io_v1alpha1/wgpolicyk8s_io_policy_report_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_wgpolicyk8s_io_policy_report_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/wgpolicyk8s_io_v1alpha1/wgpolicyk8s_io_policy_report_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-wgpolicyk8s_io_cluster_policy_report_v1alpha2_manifest_test.go: out/install-sentinel terratest/wgpolicyk8s_io_v1alpha2/wgpolicyk8s_io_cluster_policy_report_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_wgpolicyk8s_io_cluster_policy_report_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/wgpolicyk8s_io_v1alpha2/wgpolicyk8s_io_cluster_policy_report_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-wgpolicyk8s_io_policy_report_v1alpha2_manifest_test.go: out/install-sentinel terratest/wgpolicyk8s_io_v1alpha2/wgpolicyk8s_io_policy_report_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_wgpolicyk8s_io_policy_report_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/wgpolicyk8s_io_v1alpha2/wgpolicyk8s_io_policy_report_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-wgpolicyk8s_io_cluster_policy_report_v1beta1_manifest_test.go: out/install-sentinel terratest/wgpolicyk8s_io_v1beta1/wgpolicyk8s_io_cluster_policy_report_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_wgpolicyk8s_io_cluster_policy_report_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/wgpolicyk8s_io_v1beta1/wgpolicyk8s_io_cluster_policy_report_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-wgpolicyk8s_io_policy_report_v1beta1_manifest_test.go: out/install-sentinel terratest/wgpolicyk8s_io_v1beta1/wgpolicyk8s_io_policy_report_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_wgpolicyk8s_io_policy_report_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/wgpolicyk8s_io_v1beta1/wgpolicyk8s_io_policy_report_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-wildfly_org_wild_fly_server_v1alpha1_manifest_test.go: out/install-sentinel terratest/wildfly_org_v1alpha1/wildfly_org_wild_fly_server_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_wildfly_org_wild_fly_server_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/wildfly_org_v1alpha1/wildfly_org_wild_fly_server_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-work_karmada_io_cluster_resource_binding_v1alpha1_manifest_test.go: out/install-sentinel terratest/work_karmada_io_v1alpha1/work_karmada_io_cluster_resource_binding_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_work_karmada_io_cluster_resource_binding_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/work_karmada_io_v1alpha1/work_karmada_io_cluster_resource_binding_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-work_karmada_io_resource_binding_v1alpha1_manifest_test.go: out/install-sentinel terratest/work_karmada_io_v1alpha1/work_karmada_io_resource_binding_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_work_karmada_io_resource_binding_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/work_karmada_io_v1alpha1/work_karmada_io_resource_binding_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-work_karmada_io_work_v1alpha1_manifest_test.go: out/install-sentinel terratest/work_karmada_io_v1alpha1/work_karmada_io_work_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_work_karmada_io_work_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/work_karmada_io_v1alpha1/work_karmada_io_work_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-work_karmada_io_cluster_resource_binding_v1alpha2_manifest_test.go: out/install-sentinel terratest/work_karmada_io_v1alpha2/work_karmada_io_cluster_resource_binding_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_work_karmada_io_cluster_resource_binding_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/work_karmada_io_v1alpha2/work_karmada_io_cluster_resource_binding_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-work_karmada_io_resource_binding_v1alpha2_manifest_test.go: out/install-sentinel terratest/work_karmada_io_v1alpha2/work_karmada_io_resource_binding_v1alpha2_manifest_test.go $(shell find ./examples/data-sources/k8s_work_karmada_io_resource_binding_v1alpha2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/work_karmada_io_v1alpha2/work_karmada_io_resource_binding_v1alpha2_manifest_test.go
	touch $@
out/terratest-sentinel-workload_codeflare_dev_app_wrapper_v1beta1_manifest_test.go: out/install-sentinel terratest/workload_codeflare_dev_v1beta1/workload_codeflare_dev_app_wrapper_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_workload_codeflare_dev_app_wrapper_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/workload_codeflare_dev_v1beta1/workload_codeflare_dev_app_wrapper_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-workload_codeflare_dev_scheduling_spec_v1beta1_manifest_test.go: out/install-sentinel terratest/workload_codeflare_dev_v1beta1/workload_codeflare_dev_scheduling_spec_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_workload_codeflare_dev_scheduling_spec_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/workload_codeflare_dev_v1beta1/workload_codeflare_dev_scheduling_spec_v1beta1_manifest_test.go
	touch $@
out/terratest-sentinel-workload_codeflare_dev_app_wrapper_v1beta2_manifest_test.go: out/install-sentinel terratest/workload_codeflare_dev_v1beta2/workload_codeflare_dev_app_wrapper_v1beta2_manifest_test.go $(shell find ./examples/data-sources/k8s_workload_codeflare_dev_app_wrapper_v1beta2_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/workload_codeflare_dev_v1beta2/workload_codeflare_dev_app_wrapper_v1beta2_manifest_test.go
	touch $@
out/terratest-sentinel-workloads_kubeblocks_io_instance_set_v1alpha1_manifest_test.go: out/install-sentinel terratest/workloads_kubeblocks_io_v1alpha1/workloads_kubeblocks_io_instance_set_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_workloads_kubeblocks_io_instance_set_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/workloads_kubeblocks_io_v1alpha1/workloads_kubeblocks_io_instance_set_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-workloads_kubeblocks_io_replicated_state_machine_v1alpha1_manifest_test.go: out/install-sentinel terratest/workloads_kubeblocks_io_v1alpha1/workloads_kubeblocks_io_replicated_state_machine_v1alpha1_manifest_test.go $(shell find ./examples/data-sources/k8s_workloads_kubeblocks_io_replicated_state_machine_v1alpha1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/workloads_kubeblocks_io_v1alpha1/workloads_kubeblocks_io_replicated_state_machine_v1alpha1_manifest_test.go
	touch $@
out/terratest-sentinel-zonecontrol_k8s_aws_zone_aware_update_v1_manifest_test.go: out/install-sentinel terratest/zonecontrol_k8s_aws_v1/zonecontrol_k8s_aws_zone_aware_update_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_zonecontrol_k8s_aws_zone_aware_update_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/zonecontrol_k8s_aws_v1/zonecontrol_k8s_aws_zone_aware_update_v1_manifest_test.go
	touch $@
out/terratest-sentinel-zonecontrol_k8s_aws_zone_disruption_budget_v1_manifest_test.go: out/install-sentinel terratest/zonecontrol_k8s_aws_v1/zonecontrol_k8s_aws_zone_disruption_budget_v1_manifest_test.go $(shell find ./examples/data-sources/k8s_zonecontrol_k8s_aws_zone_disruption_budget_v1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/zonecontrol_k8s_aws_v1/zonecontrol_k8s_aws_zone_disruption_budget_v1_manifest_test.go
	touch $@
out/terratest-sentinel-zookeeper_pravega_io_zookeeper_cluster_v1beta1_manifest_test.go: out/install-sentinel terratest/zookeeper_pravega_io_v1beta1/zookeeper_pravega_io_zookeeper_cluster_v1beta1_manifest_test.go $(shell find ./examples/data-sources/k8s_zookeeper_pravega_io_zookeeper_cluster_v1beta1_manifest -type f -name '*.tf')
	mkdir --parents $(@D)
	go test -timeout=120s ./terratest/zookeeper_pravega_io_v1beta1/zookeeper_pravega_io_zookeeper_cluster_v1beta1_manifest_test.go
	touch $@


.PHONY: terratests
terratests: out/terratest-sentinel-about_k8s_io_cluster_property_v1alpha1_manifest_test.go out/terratest-sentinel-acid_zalan_do_operator_configuration_v1_manifest_test.go out/terratest-sentinel-acid_zalan_do_postgres_team_v1_manifest_test.go out/terratest-sentinel-acid_zalan_do_postgresql_v1_manifest_test.go out/terratest-sentinel-acme_cert_manager_io_challenge_v1_manifest_test.go out/terratest-sentinel-acme_cert_manager_io_order_v1_manifest_test.go out/terratest-sentinel-acmpca_services_k8s_aws_certificate_authority_activation_v1alpha1_manifest_test.go out/terratest-sentinel-acmpca_services_k8s_aws_certificate_authority_v1alpha1_manifest_test.go out/terratest-sentinel-acmpca_services_k8s_aws_certificate_v1alpha1_manifest_test.go out/terratest-sentinel-actions_github_com_autoscaling_listener_v1alpha1_manifest_test.go out/terratest-sentinel-actions_github_com_autoscaling_runner_set_v1alpha1_manifest_test.go out/terratest-sentinel-actions_github_com_ephemeral_runner_set_v1alpha1_manifest_test.go out/terratest-sentinel-actions_github_com_ephemeral_runner_v1alpha1_manifest_test.go out/terratest-sentinel-actions_summerwind_dev_horizontal_runner_autoscaler_v1alpha1_manifest_test.go out/terratest-sentinel-actions_summerwind_dev_runner_deployment_v1alpha1_manifest_test.go out/terratest-sentinel-actions_summerwind_dev_runner_replica_set_v1alpha1_manifest_test.go out/terratest-sentinel-actions_summerwind_dev_runner_set_v1alpha1_manifest_test.go out/terratest-sentinel-actions_summerwind_dev_runner_v1alpha1_manifest_test.go out/terratest-sentinel-addons_cluster_x_k8s_io_cluster_resource_set_binding_v1alpha3_manifest_test.go out/terratest-sentinel-addons_cluster_x_k8s_io_cluster_resource_set_v1alpha3_manifest_test.go out/terratest-sentinel-addons_cluster_x_k8s_io_cluster_resource_set_binding_v1alpha4_manifest_test.go out/terratest-sentinel-addons_cluster_x_k8s_io_cluster_resource_set_v1alpha4_manifest_test.go out/terratest-sentinel-addons_cluster_x_k8s_io_cluster_resource_set_binding_v1beta1_manifest_test.go out/terratest-sentinel-addons_cluster_x_k8s_io_cluster_resource_set_v1beta1_manifest_test.go out/terratest-sentinel-admissionregistration_k8s_io_mutating_webhook_configuration_v1_manifest_test.go out/terratest-sentinel-admissionregistration_k8s_io_validating_webhook_configuration_v1_manifest_test.go out/terratest-sentinel-agent_k8s_elastic_co_agent_v1alpha1_manifest_test.go out/terratest-sentinel-anywhere_eks_amazonaws_com_aws_datacenter_config_v1alpha1_manifest_test.go out/terratest-sentinel-anywhere_eks_amazonaws_com_aws_iam_config_v1alpha1_manifest_test.go out/terratest-sentinel-anywhere_eks_amazonaws_com_bundles_v1alpha1_manifest_test.go out/terratest-sentinel-anywhere_eks_amazonaws_com_cloud_stack_datacenter_config_v1alpha1_manifest_test.go out/terratest-sentinel-anywhere_eks_amazonaws_com_cloud_stack_machine_config_v1alpha1_manifest_test.go out/terratest-sentinel-anywhere_eks_amazonaws_com_cluster_v1alpha1_manifest_test.go out/terratest-sentinel-anywhere_eks_amazonaws_com_control_plane_upgrade_v1alpha1_manifest_test.go out/terratest-sentinel-anywhere_eks_amazonaws_com_docker_datacenter_config_v1alpha1_manifest_test.go out/terratest-sentinel-anywhere_eks_amazonaws_com_eksa_release_v1alpha1_manifest_test.go out/terratest-sentinel-anywhere_eks_amazonaws_com_flux_config_v1alpha1_manifest_test.go out/terratest-sentinel-anywhere_eks_amazonaws_com_git_ops_config_v1alpha1_manifest_test.go out/terratest-sentinel-anywhere_eks_amazonaws_com_machine_deployment_upgrade_v1alpha1_manifest_test.go out/terratest-sentinel-anywhere_eks_amazonaws_com_node_upgrade_v1alpha1_manifest_test.go out/terratest-sentinel-anywhere_eks_amazonaws_com_nutanix_datacenter_config_v1alpha1_manifest_test.go out/terratest-sentinel-anywhere_eks_amazonaws_com_nutanix_machine_config_v1alpha1_manifest_test.go out/terratest-sentinel-anywhere_eks_amazonaws_com_oidc_config_v1alpha1_manifest_test.go out/terratest-sentinel-anywhere_eks_amazonaws_com_snow_datacenter_config_v1alpha1_manifest_test.go out/terratest-sentinel-anywhere_eks_amazonaws_com_snow_ip_pool_v1alpha1_manifest_test.go out/terratest-sentinel-anywhere_eks_amazonaws_com_snow_machine_config_v1alpha1_manifest_test.go out/terratest-sentinel-anywhere_eks_amazonaws_com_tinkerbell_datacenter_config_v1alpha1_manifest_test.go out/terratest-sentinel-anywhere_eks_amazonaws_com_tinkerbell_machine_config_v1alpha1_manifest_test.go out/terratest-sentinel-anywhere_eks_amazonaws_com_tinkerbell_template_config_v1alpha1_manifest_test.go out/terratest-sentinel-anywhere_eks_amazonaws_com_v_sphere_datacenter_config_v1alpha1_manifest_test.go out/terratest-sentinel-anywhere_eks_amazonaws_com_v_sphere_machine_config_v1alpha1_manifest_test.go out/terratest-sentinel-apacheweb_arsenal_dev_apacheweb_v1alpha1_manifest_test.go out/terratest-sentinel-api_clever_cloud_com_config_provider_v1_manifest_test.go out/terratest-sentinel-api_clever_cloud_com_elastic_search_v1_manifest_test.go out/terratest-sentinel-api_clever_cloud_com_mongo_db_v1_manifest_test.go out/terratest-sentinel-api_clever_cloud_com_my_sql_v1_manifest_test.go out/terratest-sentinel-api_clever_cloud_com_postgre_sql_v1_manifest_test.go out/terratest-sentinel-api_clever_cloud_com_redis_v1_manifest_test.go out/terratest-sentinel-api_clever_cloud_com_pulsar_v1beta1_manifest_test.go out/terratest-sentinel-api_kubemod_io_mod_rule_v1beta1_manifest_test.go out/terratest-sentinel-apicodegen_apimatic_io_api_matic_v1beta1_manifest_test.go out/terratest-sentinel-apiextensions_crossplane_io_composite_resource_definition_v1_manifest_test.go out/terratest-sentinel-apiextensions_crossplane_io_composition_revision_v1_manifest_test.go out/terratest-sentinel-apiextensions_crossplane_io_composition_v1_manifest_test.go out/terratest-sentinel-apiextensions_crossplane_io_composition_revision_v1beta1_manifest_test.go out/terratest-sentinel-apigatewayv2_services_k8s_aws_api_v1alpha1_manifest_test.go out/terratest-sentinel-apigatewayv2_services_k8s_aws_authorizer_v1alpha1_manifest_test.go out/terratest-sentinel-apigatewayv2_services_k8s_aws_deployment_v1alpha1_manifest_test.go out/terratest-sentinel-apigatewayv2_services_k8s_aws_integration_v1alpha1_manifest_test.go out/terratest-sentinel-apigatewayv2_services_k8s_aws_route_v1alpha1_manifest_test.go out/terratest-sentinel-apigatewayv2_services_k8s_aws_stage_v1alpha1_manifest_test.go out/terratest-sentinel-apigatewayv2_services_k8s_aws_vpc_link_v1alpha1_manifest_test.go out/terratest-sentinel-apiregistration_k8s_io_api_service_v1_manifest_test.go out/terratest-sentinel-apisix_apache_org_apisix_cluster_config_v2_manifest_test.go out/terratest-sentinel-apisix_apache_org_apisix_consumer_v2_manifest_test.go out/terratest-sentinel-apisix_apache_org_apisix_global_rule_v2_manifest_test.go out/terratest-sentinel-apisix_apache_org_apisix_plugin_config_v2_manifest_test.go out/terratest-sentinel-apisix_apache_org_apisix_route_v2_manifest_test.go out/terratest-sentinel-apisix_apache_org_apisix_tls_v2_manifest_test.go out/terratest-sentinel-apisix_apache_org_apisix_upstream_v2_manifest_test.go out/terratest-sentinel-apm_k8s_elastic_co_apm_server_v1_manifest_test.go out/terratest-sentinel-apm_k8s_elastic_co_apm_server_v1beta1_manifest_test.go out/terratest-sentinel-app_kiegroup_org_kogito_build_v1beta1_manifest_test.go out/terratest-sentinel-app_kiegroup_org_kogito_infra_v1beta1_manifest_test.go out/terratest-sentinel-app_kiegroup_org_kogito_runtime_v1beta1_manifest_test.go out/terratest-sentinel-app_kiegroup_org_kogito_supporting_service_v1beta1_manifest_test.go out/terratest-sentinel-app_lightbend_com_akka_cluster_v1alpha1_manifest_test.go out/terratest-sentinel-app_redislabs_com_redis_enterprise_cluster_v1_manifest_test.go out/terratest-sentinel-app_redislabs_com_redis_enterprise_active_active_database_v1alpha1_manifest_test.go out/terratest-sentinel-app_redislabs_com_redis_enterprise_cluster_v1alpha1_manifest_test.go out/terratest-sentinel-app_redislabs_com_redis_enterprise_database_v1alpha1_manifest_test.go out/terratest-sentinel-app_redislabs_com_redis_enterprise_remote_cluster_v1alpha1_manifest_test.go out/terratest-sentinel-app_terraform_io_agent_pool_v1alpha2_manifest_test.go out/terratest-sentinel-app_terraform_io_module_v1alpha2_manifest_test.go out/terratest-sentinel-app_terraform_io_workspace_v1alpha2_manifest_test.go out/terratest-sentinel-application_networking_k8s_aws_access_log_policy_v1alpha1_manifest_test.go out/terratest-sentinel-application_networking_k8s_aws_iam_auth_policy_v1alpha1_manifest_test.go out/terratest-sentinel-application_networking_k8s_aws_service_import_v1alpha1_manifest_test.go out/terratest-sentinel-application_networking_k8s_aws_target_group_policy_v1alpha1_manifest_test.go out/terratest-sentinel-application_networking_k8s_aws_vpc_association_policy_v1alpha1_manifest_test.go out/terratest-sentinel-applicationautoscaling_services_k8s_aws_scalable_target_v1alpha1_manifest_test.go out/terratest-sentinel-applicationautoscaling_services_k8s_aws_scaling_policy_v1alpha1_manifest_test.go out/terratest-sentinel-appmesh_k8s_aws_backend_group_v1beta2_manifest_test.go out/terratest-sentinel-appmesh_k8s_aws_gateway_route_v1beta2_manifest_test.go out/terratest-sentinel-appmesh_k8s_aws_mesh_v1beta2_manifest_test.go out/terratest-sentinel-appmesh_k8s_aws_virtual_gateway_v1beta2_manifest_test.go out/terratest-sentinel-appmesh_k8s_aws_virtual_node_v1beta2_manifest_test.go out/terratest-sentinel-appmesh_k8s_aws_virtual_router_v1beta2_manifest_test.go out/terratest-sentinel-appmesh_k8s_aws_virtual_service_v1beta2_manifest_test.go out/terratest-sentinel-appprotect_f5_com_ap_log_conf_v1beta1_manifest_test.go out/terratest-sentinel-appprotect_f5_com_ap_policy_v1beta1_manifest_test.go out/terratest-sentinel-appprotect_f5_com_ap_user_sig_v1beta1_manifest_test.go out/terratest-sentinel-appprotectdos_f5_com_ap_dos_log_conf_v1beta1_manifest_test.go out/terratest-sentinel-appprotectdos_f5_com_ap_dos_policy_v1beta1_manifest_test.go out/terratest-sentinel-appprotectdos_f5_com_dos_protected_resource_v1beta1_manifest_test.go out/terratest-sentinel-apps_3scale_net_ap_icast_v1alpha1_manifest_test.go out/terratest-sentinel-apps_3scale_net_api_manager_backup_v1alpha1_manifest_test.go out/terratest-sentinel-apps_3scale_net_api_manager_restore_v1alpha1_manifest_test.go out/terratest-sentinel-apps_3scale_net_api_manager_v1alpha1_manifest_test.go out/terratest-sentinel-apps_clusternet_io_base_v1alpha1_manifest_test.go out/terratest-sentinel-apps_clusternet_io_description_v1alpha1_manifest_test.go out/terratest-sentinel-apps_clusternet_io_feed_inventory_v1alpha1_manifest_test.go out/terratest-sentinel-apps_clusternet_io_globalization_v1alpha1_manifest_test.go out/terratest-sentinel-apps_clusternet_io_helm_chart_v1alpha1_manifest_test.go out/terratest-sentinel-apps_clusternet_io_helm_release_v1alpha1_manifest_test.go out/terratest-sentinel-apps_clusternet_io_localization_v1alpha1_manifest_test.go out/terratest-sentinel-apps_clusternet_io_manifest_v1alpha1_manifest_test.go out/terratest-sentinel-apps_clusternet_io_subscription_v1alpha1_manifest_test.go out/terratest-sentinel-apps_emqx_io_emqx_broker_v1beta3_manifest_test.go out/terratest-sentinel-apps_emqx_io_emqx_enterprise_v1beta3_manifest_test.go out/terratest-sentinel-apps_emqx_io_emqx_plugin_v1beta3_manifest_test.go out/terratest-sentinel-apps_emqx_io_emqx_broker_v1beta4_manifest_test.go out/terratest-sentinel-apps_emqx_io_emqx_enterprise_v1beta4_manifest_test.go out/terratest-sentinel-apps_emqx_io_emqx_plugin_v1beta4_manifest_test.go out/terratest-sentinel-apps_emqx_io_rebalance_v1beta4_manifest_test.go out/terratest-sentinel-apps_emqx_io_emqx_v2alpha1_manifest_test.go out/terratest-sentinel-apps_emqx_io_emqx_v2beta1_manifest_test.go out/terratest-sentinel-apps_emqx_io_rebalance_v2beta1_manifest_test.go out/terratest-sentinel-apps_gitlab_com_git_lab_v1beta1_manifest_test.go out/terratest-sentinel-apps_gitlab_com_runner_v1beta2_manifest_test.go out/terratest-sentinel-apps_kubeblocks_io_backup_policy_template_v1alpha1_manifest_test.go out/terratest-sentinel-apps_kubeblocks_io_cluster_definition_v1alpha1_manifest_test.go out/terratest-sentinel-apps_kubeblocks_io_cluster_v1alpha1_manifest_test.go out/terratest-sentinel-apps_kubeblocks_io_cluster_version_v1alpha1_manifest_test.go out/terratest-sentinel-apps_kubeblocks_io_component_class_definition_v1alpha1_manifest_test.go out/terratest-sentinel-apps_kubeblocks_io_component_definition_v1alpha1_manifest_test.go out/terratest-sentinel-apps_kubeblocks_io_component_resource_constraint_v1alpha1_manifest_test.go out/terratest-sentinel-apps_kubeblocks_io_component_v1alpha1_manifest_test.go out/terratest-sentinel-apps_kubeblocks_io_component_version_v1alpha1_manifest_test.go out/terratest-sentinel-apps_kubeblocks_io_config_constraint_v1alpha1_manifest_test.go out/terratest-sentinel-apps_kubeblocks_io_configuration_v1alpha1_manifest_test.go out/terratest-sentinel-apps_kubeblocks_io_ops_definition_v1alpha1_manifest_test.go out/terratest-sentinel-apps_kubeblocks_io_ops_request_v1alpha1_manifest_test.go out/terratest-sentinel-apps_kubeblocks_io_service_descriptor_v1alpha1_manifest_test.go out/terratest-sentinel-apps_kubeblocks_io_config_constraint_v1beta1_manifest_test.go out/terratest-sentinel-apps_kubedl_io_cron_v1alpha1_manifest_test.go out/terratest-sentinel-apps_kubeedge_io_edge_application_v1alpha1_manifest_test.go out/terratest-sentinel-apps_kubeedge_io_node_group_v1alpha1_manifest_test.go out/terratest-sentinel-apps_m88i_io_nexus_v1alpha1_manifest_test.go out/terratest-sentinel-apps_redhat_com_cluster_impairment_v1alpha1_manifest_test.go out/terratest-sentinel-apps_daemon_set_v1_manifest_test.go out/terratest-sentinel-apps_deployment_v1_manifest_test.go out/terratest-sentinel-apps_replica_set_v1_manifest_test.go out/terratest-sentinel-apps_stateful_set_v1_manifest_test.go out/terratest-sentinel-aquasecurity_github_io_aqua_starboard_v1alpha1_manifest_test.go out/terratest-sentinel-argoproj_io_app_project_v1alpha1_manifest_test.go out/terratest-sentinel-argoproj_io_application_set_v1alpha1_manifest_test.go out/terratest-sentinel-argoproj_io_application_v1alpha1_manifest_test.go out/terratest-sentinel-argoproj_io_argo_cd_v1alpha1_manifest_test.go out/terratest-sentinel-argoproj_io_argo_cd_export_v1alpha1_manifest_test.go out/terratest-sentinel-argoproj_io_argo_cd_v1beta1_manifest_test.go out/terratest-sentinel-asdb_aerospike_com_aerospike_cluster_v1_manifest_test.go out/terratest-sentinel-asdb_aerospike_com_aerospike_cluster_v1beta1_manifest_test.go out/terratest-sentinel-atlasmap_io_atlas_map_v1alpha1_manifest_test.go out/terratest-sentinel-auth_ops42_org_aws_auth_sync_config_v1alpha1_manifest_test.go out/terratest-sentinel-authzed_com_spice_db_cluster_v1alpha1_manifest_test.go out/terratest-sentinel-automation_kubensync_com_managed_resource_v1alpha1_manifest_test.go out/terratest-sentinel-autoscaling_k8s_elastic_co_elasticsearch_autoscaler_v1alpha1_manifest_test.go out/terratest-sentinel-autoscaling_k8s_io_vertical_pod_autoscaler_checkpoint_v1_manifest_test.go out/terratest-sentinel-autoscaling_k8s_io_vertical_pod_autoscaler_v1_manifest_test.go out/terratest-sentinel-autoscaling_k8s_io_vertical_pod_autoscaler_checkpoint_v1beta2_manifest_test.go out/terratest-sentinel-autoscaling_k8s_io_vertical_pod_autoscaler_v1beta2_manifest_test.go out/terratest-sentinel-autoscaling_karmada_io_cron_federated_hpa_v1alpha1_manifest_test.go out/terratest-sentinel-autoscaling_karmada_io_federated_hpa_v1alpha1_manifest_test.go out/terratest-sentinel-autoscaling_horizontal_pod_autoscaler_v1_manifest_test.go out/terratest-sentinel-autoscaling_horizontal_pod_autoscaler_v2_manifest_test.go out/terratest-sentinel-awx_ansible_com_awx_v1beta1_manifest_test.go out/terratest-sentinel-awx_ansible_com_awx_backup_v1beta1_manifest_test.go out/terratest-sentinel-awx_ansible_com_awx_restore_v1beta1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_apim_service_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_api_mgmt_api_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_app_insights_api_key_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_app_insights_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_azure_load_balancer_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_azure_network_interface_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_azure_public_ip_address_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_azure_sql_action_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_azure_sql_database_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_azure_sql_failover_group_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_azure_sql_firewall_rule_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_azure_sql_server_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_azure_sql_managed_user_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_azure_sql_user_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_azure_sqlv_net_rule_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_azure_virtual_machine_extension_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_azure_virtual_machine_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_azure_vm_scale_set_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_blob_container_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_consumer_group_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_cosmos_db_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_eventhub_namespace_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_eventhub_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_key_vault_key_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_key_vault_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_my_sqlaad_user_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_my_sql_database_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_my_sql_firewall_rule_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_my_sql_server_administrator_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_my_sql_server_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_my_sql_user_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_my_sqlv_net_rule_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_postgre_sql_database_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_postgre_sql_firewall_rule_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_postgre_sql_server_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_postgre_sql_user_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_postgre_sqlv_net_rule_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_redis_cache_action_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_redis_cache_firewall_rule_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_resource_group_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_storage_account_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_virtual_network_v1alpha1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_blob_container_v1alpha2_manifest_test.go out/terratest-sentinel-azure_microsoft_com_my_sqlaad_user_v1alpha2_manifest_test.go out/terratest-sentinel-azure_microsoft_com_my_sql_server_v1alpha2_manifest_test.go out/terratest-sentinel-azure_microsoft_com_my_sql_user_v1alpha2_manifest_test.go out/terratest-sentinel-azure_microsoft_com_postgre_sql_server_v1alpha2_manifest_test.go out/terratest-sentinel-azure_microsoft_com_azure_sql_database_v1beta1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_azure_sql_failover_group_v1beta1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_azure_sql_firewall_rule_v1beta1_manifest_test.go out/terratest-sentinel-azure_microsoft_com_azure_sql_server_v1beta1_manifest_test.go out/terratest-sentinel-b3scale_infra_run_bbb_frontend_v1_manifest_test.go out/terratest-sentinel-b3scale_io_bbb_frontend_v1_manifest_test.go out/terratest-sentinel-batch_cron_job_v1_manifest_test.go out/terratest-sentinel-batch_job_v1_manifest_test.go out/terratest-sentinel-batch_volcano_sh_job_v1alpha1_manifest_test.go out/terratest-sentinel-beat_k8s_elastic_co_beat_v1beta1_manifest_test.go out/terratest-sentinel-beegfs_csi_netapp_com_beegfs_driver_v1_manifest_test.go out/terratest-sentinel-binding_operators_coreos_com_service_binding_v1alpha1_manifest_test.go out/terratest-sentinel-bitnami_com_sealed_secret_v1alpha1_manifest_test.go out/terratest-sentinel-bmc_tinkerbell_org_job_v1alpha1_manifest_test.go out/terratest-sentinel-bmc_tinkerbell_org_machine_v1alpha1_manifest_test.go out/terratest-sentinel-bmc_tinkerbell_org_task_v1alpha1_manifest_test.go out/terratest-sentinel-boskos_k8s_io_drlc_object_v1_manifest_test.go out/terratest-sentinel-boskos_k8s_io_resource_object_v1_manifest_test.go out/terratest-sentinel-bpfman_io_bpf_program_v1alpha1_manifest_test.go out/terratest-sentinel-bpfman_io_fentry_program_v1alpha1_manifest_test.go out/terratest-sentinel-bpfman_io_fexit_program_v1alpha1_manifest_test.go out/terratest-sentinel-bpfman_io_kprobe_program_v1alpha1_manifest_test.go out/terratest-sentinel-bpfman_io_tc_program_v1alpha1_manifest_test.go out/terratest-sentinel-bpfman_io_tracepoint_program_v1alpha1_manifest_test.go out/terratest-sentinel-bpfman_io_uprobe_program_v1alpha1_manifest_test.go out/terratest-sentinel-bpfman_io_xdp_program_v1alpha1_manifest_test.go out/terratest-sentinel-bus_volcano_sh_command_v1alpha1_manifest_test.go out/terratest-sentinel-cache_kubedl_io_cache_backend_v1alpha1_manifest_test.go out/terratest-sentinel-caching_ibm_com_varnish_cluster_v1alpha1_manifest_test.go out/terratest-sentinel-camel_apache_org_build_v1_manifest_test.go out/terratest-sentinel-camel_apache_org_camel_catalog_v1_manifest_test.go out/terratest-sentinel-camel_apache_org_integration_kit_v1_manifest_test.go out/terratest-sentinel-camel_apache_org_integration_platform_v1_manifest_test.go out/terratest-sentinel-camel_apache_org_integration_profile_v1_manifest_test.go out/terratest-sentinel-camel_apache_org_integration_v1_manifest_test.go out/terratest-sentinel-camel_apache_org_kamelet_v1_manifest_test.go out/terratest-sentinel-camel_apache_org_pipe_v1_manifest_test.go out/terratest-sentinel-camel_apache_org_kamelet_binding_v1alpha1_manifest_test.go out/terratest-sentinel-camel_apache_org_kamelet_v1alpha1_manifest_test.go out/terratest-sentinel-canaries_flanksource_com_canary_v1_manifest_test.go out/terratest-sentinel-canaries_flanksource_com_component_v1_manifest_test.go out/terratest-sentinel-canaries_flanksource_com_topology_v1_manifest_test.go out/terratest-sentinel-capabilities_3scale_net_tenant_v1alpha1_manifest_test.go out/terratest-sentinel-capabilities_3scale_net_active_doc_v1beta1_manifest_test.go out/terratest-sentinel-capabilities_3scale_net_application_v1beta1_manifest_test.go out/terratest-sentinel-capabilities_3scale_net_backend_v1beta1_manifest_test.go out/terratest-sentinel-capabilities_3scale_net_custom_policy_definition_v1beta1_manifest_test.go out/terratest-sentinel-capabilities_3scale_net_developer_account_v1beta1_manifest_test.go out/terratest-sentinel-capabilities_3scale_net_developer_user_v1beta1_manifest_test.go out/terratest-sentinel-capabilities_3scale_net_open_api_v1beta1_manifest_test.go out/terratest-sentinel-capabilities_3scale_net_product_v1beta1_manifest_test.go out/terratest-sentinel-capabilities_3scale_net_proxy_config_promote_v1beta1_manifest_test.go out/terratest-sentinel-capsule_clastix_io_capsule_configuration_v1alpha1_manifest_test.go out/terratest-sentinel-capsule_clastix_io_tenant_v1alpha1_manifest_test.go out/terratest-sentinel-capsule_clastix_io_tenant_v1beta1_manifest_test.go out/terratest-sentinel-capsule_clastix_io_capsule_configuration_v1beta2_manifest_test.go out/terratest-sentinel-capsule_clastix_io_tenant_v1beta2_manifest_test.go out/terratest-sentinel-cassandra_datastax_com_cassandra_datacenter_v1beta1_manifest_test.go out/terratest-sentinel-ceph_rook_io_ceph_block_pool_rados_namespace_v1_manifest_test.go out/terratest-sentinel-ceph_rook_io_ceph_block_pool_v1_manifest_test.go out/terratest-sentinel-ceph_rook_io_ceph_bucket_notification_v1_manifest_test.go out/terratest-sentinel-ceph_rook_io_ceph_bucket_topic_v1_manifest_test.go out/terratest-sentinel-ceph_rook_io_ceph_client_v1_manifest_test.go out/terratest-sentinel-ceph_rook_io_ceph_cluster_v1_manifest_test.go out/terratest-sentinel-ceph_rook_io_ceph_cosi_driver_v1_manifest_test.go out/terratest-sentinel-ceph_rook_io_ceph_filesystem_mirror_v1_manifest_test.go out/terratest-sentinel-ceph_rook_io_ceph_filesystem_sub_volume_group_v1_manifest_test.go out/terratest-sentinel-ceph_rook_io_ceph_filesystem_v1_manifest_test.go out/terratest-sentinel-ceph_rook_io_ceph_nfs_v1_manifest_test.go out/terratest-sentinel-ceph_rook_io_ceph_object_realm_v1_manifest_test.go out/terratest-sentinel-ceph_rook_io_ceph_object_store_user_v1_manifest_test.go out/terratest-sentinel-ceph_rook_io_ceph_object_store_v1_manifest_test.go out/terratest-sentinel-ceph_rook_io_ceph_object_zone_group_v1_manifest_test.go out/terratest-sentinel-ceph_rook_io_ceph_object_zone_v1_manifest_test.go out/terratest-sentinel-ceph_rook_io_ceph_rbd_mirror_v1_manifest_test.go out/terratest-sentinel-cert_manager_io_certificate_request_v1_manifest_test.go out/terratest-sentinel-cert_manager_io_certificate_v1_manifest_test.go out/terratest-sentinel-cert_manager_io_cluster_issuer_v1_manifest_test.go out/terratest-sentinel-cert_manager_io_issuer_v1_manifest_test.go out/terratest-sentinel-certificates_k8s_io_certificate_signing_request_v1_manifest_test.go out/terratest-sentinel-chainsaw_kyverno_io_configuration_v1alpha1_manifest_test.go out/terratest-sentinel-chainsaw_kyverno_io_test_v1alpha1_manifest_test.go out/terratest-sentinel-chainsaw_kyverno_io_configuration_v1alpha2_manifest_test.go out/terratest-sentinel-chainsaw_kyverno_io_test_v1alpha2_manifest_test.go out/terratest-sentinel-chaos_mesh_org_aws_chaos_v1alpha1_manifest_test.go out/terratest-sentinel-chaos_mesh_org_azure_chaos_v1alpha1_manifest_test.go out/terratest-sentinel-chaos_mesh_org_block_chaos_v1alpha1_manifest_test.go out/terratest-sentinel-chaos_mesh_org_dns_chaos_v1alpha1_manifest_test.go out/terratest-sentinel-chaos_mesh_org_gcp_chaos_v1alpha1_manifest_test.go out/terratest-sentinel-chaos_mesh_org_http_chaos_v1alpha1_manifest_test.go out/terratest-sentinel-chaos_mesh_org_io_chaos_v1alpha1_manifest_test.go out/terratest-sentinel-chaos_mesh_org_jvm_chaos_v1alpha1_manifest_test.go out/terratest-sentinel-chaos_mesh_org_kernel_chaos_v1alpha1_manifest_test.go out/terratest-sentinel-chaos_mesh_org_network_chaos_v1alpha1_manifest_test.go out/terratest-sentinel-chaos_mesh_org_physical_machine_chaos_v1alpha1_manifest_test.go out/terratest-sentinel-chaos_mesh_org_physical_machine_v1alpha1_manifest_test.go out/terratest-sentinel-chaos_mesh_org_pod_chaos_v1alpha1_manifest_test.go out/terratest-sentinel-chaos_mesh_org_pod_http_chaos_v1alpha1_manifest_test.go out/terratest-sentinel-chaos_mesh_org_pod_io_chaos_v1alpha1_manifest_test.go out/terratest-sentinel-chaos_mesh_org_pod_network_chaos_v1alpha1_manifest_test.go out/terratest-sentinel-chaos_mesh_org_remote_cluster_v1alpha1_manifest_test.go out/terratest-sentinel-chaos_mesh_org_schedule_v1alpha1_manifest_test.go out/terratest-sentinel-chaos_mesh_org_status_check_v1alpha1_manifest_test.go out/terratest-sentinel-chaos_mesh_org_stress_chaos_v1alpha1_manifest_test.go out/terratest-sentinel-chaos_mesh_org_time_chaos_v1alpha1_manifest_test.go out/terratest-sentinel-chaos_mesh_org_workflow_node_v1alpha1_manifest_test.go out/terratest-sentinel-chaos_mesh_org_workflow_v1alpha1_manifest_test.go out/terratest-sentinel-chaosblade_io_chaos_blade_v1alpha1_manifest_test.go out/terratest-sentinel-charts_amd_com_amdgpu_v1alpha1_manifest_test.go out/terratest-sentinel-charts_flagsmith_com_flagsmith_v1alpha1_manifest_test.go out/terratest-sentinel-charts_helm_k8s_io_snyk_monitor_v1alpha1_manifest_test.go out/terratest-sentinel-charts_opdev_io_synapse_v1alpha1_manifest_test.go out/terratest-sentinel-charts_operatorhub_io_cockroachdb_v1alpha1_manifest_test.go out/terratest-sentinel-che_eclipse_org_kubernetes_image_puller_v1alpha1_manifest_test.go out/terratest-sentinel-chisel_operator_io_exit_node_provisioner_v1_manifest_test.go out/terratest-sentinel-chisel_operator_io_exit_node_v1_manifest_test.go out/terratest-sentinel-chisel_operator_io_exit_node_v2_manifest_test.go out/terratest-sentinel-cilium_io_cilium_clusterwide_envoy_config_v2_manifest_test.go out/terratest-sentinel-cilium_io_cilium_clusterwide_network_policy_v2_manifest_test.go out/terratest-sentinel-cilium_io_cilium_egress_gateway_policy_v2_manifest_test.go out/terratest-sentinel-cilium_io_cilium_envoy_config_v2_manifest_test.go out/terratest-sentinel-cilium_io_cilium_external_workload_v2_manifest_test.go out/terratest-sentinel-cilium_io_cilium_identity_v2_manifest_test.go out/terratest-sentinel-cilium_io_cilium_local_redirect_policy_v2_manifest_test.go out/terratest-sentinel-cilium_io_cilium_network_policy_v2_manifest_test.go out/terratest-sentinel-cilium_io_cilium_node_v2_manifest_test.go out/terratest-sentinel-cilium_io_cilium_bgp_peering_policy_v2alpha1_manifest_test.go out/terratest-sentinel-cilium_io_cilium_cidr_group_v2alpha1_manifest_test.go out/terratest-sentinel-cilium_io_cilium_endpoint_slice_v2alpha1_manifest_test.go out/terratest-sentinel-cilium_io_cilium_l2_announcement_policy_v2alpha1_manifest_test.go out/terratest-sentinel-cilium_io_cilium_load_balancer_ip_pool_v2alpha1_manifest_test.go out/terratest-sentinel-cilium_io_cilium_node_config_v2alpha1_manifest_test.go out/terratest-sentinel-cilium_io_cilium_pod_ip_pool_v2alpha1_manifest_test.go out/terratest-sentinel-claudie_io_input_manifest_v1beta1_manifest_test.go out/terratest-sentinel-cloudformation_linki_space_stack_v1alpha1_manifest_test.go out/terratest-sentinel-cloudfront_services_k8s_aws_cache_policy_v1alpha1_manifest_test.go out/terratest-sentinel-cloudfront_services_k8s_aws_distribution_v1alpha1_manifest_test.go out/terratest-sentinel-cloudfront_services_k8s_aws_function_v1alpha1_manifest_test.go out/terratest-sentinel-cloudfront_services_k8s_aws_origin_request_policy_v1alpha1_manifest_test.go out/terratest-sentinel-cloudfront_services_k8s_aws_response_headers_policy_v1alpha1_manifest_test.go out/terratest-sentinel-cloudtrail_services_k8s_aws_event_data_store_v1alpha1_manifest_test.go out/terratest-sentinel-cloudtrail_services_k8s_aws_trail_v1alpha1_manifest_test.go out/terratest-sentinel-cloudwatch_aws_amazon_com_amazon_cloud_watch_agent_v1alpha1_manifest_test.go out/terratest-sentinel-cloudwatch_aws_amazon_com_instrumentation_v1alpha1_manifest_test.go out/terratest-sentinel-cloudwatch_services_k8s_aws_metric_alarm_v1alpha1_manifest_test.go out/terratest-sentinel-cloudwatchlogs_services_k8s_aws_log_group_v1alpha1_manifest_test.go out/terratest-sentinel-cluster_clusterpedia_io_cluster_sync_resources_v1alpha2_manifest_test.go out/terratest-sentinel-cluster_clusterpedia_io_pedia_cluster_v1alpha2_manifest_test.go out/terratest-sentinel-cluster_ipfs_io_circuit_relay_v1alpha1_manifest_test.go out/terratest-sentinel-cluster_ipfs_io_ipfs_cluster_v1alpha1_manifest_test.go out/terratest-sentinel-cluster_x_k8s_io_cluster_v1alpha3_manifest_test.go out/terratest-sentinel-cluster_x_k8s_io_machine_deployment_v1alpha3_manifest_test.go out/terratest-sentinel-cluster_x_k8s_io_machine_health_check_v1alpha3_manifest_test.go out/terratest-sentinel-cluster_x_k8s_io_machine_pool_v1alpha3_manifest_test.go out/terratest-sentinel-cluster_x_k8s_io_machine_set_v1alpha3_manifest_test.go out/terratest-sentinel-cluster_x_k8s_io_machine_v1alpha3_manifest_test.go out/terratest-sentinel-cluster_x_k8s_io_cluster_class_v1alpha4_manifest_test.go out/terratest-sentinel-cluster_x_k8s_io_cluster_v1alpha4_manifest_test.go out/terratest-sentinel-cluster_x_k8s_io_machine_deployment_v1alpha4_manifest_test.go out/terratest-sentinel-cluster_x_k8s_io_machine_health_check_v1alpha4_manifest_test.go out/terratest-sentinel-cluster_x_k8s_io_machine_pool_v1alpha4_manifest_test.go out/terratest-sentinel-cluster_x_k8s_io_machine_set_v1alpha4_manifest_test.go out/terratest-sentinel-cluster_x_k8s_io_machine_v1alpha4_manifest_test.go out/terratest-sentinel-cluster_x_k8s_io_cluster_class_v1beta1_manifest_test.go out/terratest-sentinel-cluster_x_k8s_io_cluster_v1beta1_manifest_test.go out/terratest-sentinel-cluster_x_k8s_io_machine_deployment_v1beta1_manifest_test.go out/terratest-sentinel-cluster_x_k8s_io_machine_health_check_v1beta1_manifest_test.go out/terratest-sentinel-cluster_x_k8s_io_machine_pool_v1beta1_manifest_test.go out/terratest-sentinel-cluster_x_k8s_io_machine_set_v1beta1_manifest_test.go out/terratest-sentinel-cluster_x_k8s_io_machine_v1beta1_manifest_test.go out/terratest-sentinel-clusters_clusternet_io_cluster_registration_request_v1beta1_manifest_test.go out/terratest-sentinel-clusters_clusternet_io_managed_cluster_v1beta1_manifest_test.go out/terratest-sentinel-clustertemplate_openshift_io_cluster_template_instance_v1alpha1_manifest_test.go out/terratest-sentinel-clustertemplate_openshift_io_cluster_template_quota_v1alpha1_manifest_test.go out/terratest-sentinel-clustertemplate_openshift_io_cluster_template_setup_v1alpha1_manifest_test.go out/terratest-sentinel-clustertemplate_openshift_io_cluster_template_v1alpha1_manifest_test.go out/terratest-sentinel-clustertemplate_openshift_io_config_v1alpha1_manifest_test.go out/terratest-sentinel-confidentialcontainers_org_cc_runtime_v1beta1_manifest_test.go out/terratest-sentinel-config_gatekeeper_sh_config_v1alpha1_manifest_test.go out/terratest-sentinel-config_grafana_com_project_config_v1_manifest_test.go out/terratest-sentinel-config_karmada_io_resource_interpreter_customization_v1alpha1_manifest_test.go out/terratest-sentinel-config_karmada_io_resource_interpreter_webhook_configuration_v1alpha1_manifest_test.go out/terratest-sentinel-config_koordinator_sh_cluster_colocation_profile_v1alpha1_manifest_test.go out/terratest-sentinel-config_storageos_com_operator_config_v1_manifest_test.go out/terratest-sentinel-control_k8ssandra_io_cassandra_task_v1alpha1_manifest_test.go out/terratest-sentinel-core_kubeadmiral_io_cluster_collected_status_v1alpha1_manifest_test.go out/terratest-sentinel-core_kubeadmiral_io_cluster_federated_object_v1alpha1_manifest_test.go out/terratest-sentinel-core_kubeadmiral_io_cluster_override_policy_v1alpha1_manifest_test.go out/terratest-sentinel-core_kubeadmiral_io_cluster_propagation_policy_v1alpha1_manifest_test.go out/terratest-sentinel-core_kubeadmiral_io_collected_status_v1alpha1_manifest_test.go out/terratest-sentinel-core_kubeadmiral_io_federated_cluster_v1alpha1_manifest_test.go out/terratest-sentinel-core_kubeadmiral_io_federated_object_v1alpha1_manifest_test.go out/terratest-sentinel-core_kubeadmiral_io_federated_type_config_v1alpha1_manifest_test.go out/terratest-sentinel-core_kubeadmiral_io_override_policy_v1alpha1_manifest_test.go out/terratest-sentinel-core_kubeadmiral_io_propagation_policy_v1alpha1_manifest_test.go out/terratest-sentinel-core_kubeadmiral_io_scheduler_plugin_webhook_configuration_v1alpha1_manifest_test.go out/terratest-sentinel-core_kubeadmiral_io_scheduling_profile_v1alpha1_manifest_test.go out/terratest-sentinel-core_linuxsuren_github_com_a_test_v1alpha1_manifest_test.go out/terratest-sentinel-core_openfeature_dev_feature_flag_configuration_v1alpha1_manifest_test.go out/terratest-sentinel-core_openfeature_dev_feature_flag_configuration_v1alpha2_manifest_test.go out/terratest-sentinel-core_strimzi_io_strimzi_pod_set_v1beta2_manifest_test.go out/terratest-sentinel-config_map_v1_manifest_test.go out/terratest-sentinel-endpoints_v1_manifest_test.go out/terratest-sentinel-limit_range_v1_manifest_test.go out/terratest-sentinel-namespace_v1_manifest_test.go out/terratest-sentinel-persistent_volume_claim_v1_manifest_test.go out/terratest-sentinel-persistent_volume_v1_manifest_test.go out/terratest-sentinel-pod_v1_manifest_test.go out/terratest-sentinel-replication_controller_v1_manifest_test.go out/terratest-sentinel-secret_v1_manifest_test.go out/terratest-sentinel-service_account_v1_manifest_test.go out/terratest-sentinel-service_v1_manifest_test.go out/terratest-sentinel-couchbase_com_couchbase_autoscaler_v2_manifest_test.go out/terratest-sentinel-couchbase_com_couchbase_backup_restore_v2_manifest_test.go out/terratest-sentinel-couchbase_com_couchbase_backup_v2_manifest_test.go out/terratest-sentinel-couchbase_com_couchbase_bucket_v2_manifest_test.go out/terratest-sentinel-couchbase_com_couchbase_cluster_v2_manifest_test.go out/terratest-sentinel-couchbase_com_couchbase_collection_group_v2_manifest_test.go out/terratest-sentinel-couchbase_com_couchbase_collection_v2_manifest_test.go out/terratest-sentinel-couchbase_com_couchbase_ephemeral_bucket_v2_manifest_test.go out/terratest-sentinel-couchbase_com_couchbase_group_v2_manifest_test.go out/terratest-sentinel-couchbase_com_couchbase_memcached_bucket_v2_manifest_test.go out/terratest-sentinel-couchbase_com_couchbase_migration_replication_v2_manifest_test.go out/terratest-sentinel-couchbase_com_couchbase_replication_v2_manifest_test.go out/terratest-sentinel-couchbase_com_couchbase_role_binding_v2_manifest_test.go out/terratest-sentinel-couchbase_com_couchbase_scope_group_v2_manifest_test.go out/terratest-sentinel-couchbase_com_couchbase_scope_v2_manifest_test.go out/terratest-sentinel-couchbase_com_couchbase_user_v2_manifest_test.go out/terratest-sentinel-craftypath_github_io_sops_secret_v1alpha1_manifest_test.go out/terratest-sentinel-crane_konveyor_io_operator_config_v1alpha1_manifest_test.go out/terratest-sentinel-crd_projectcalico_org_bgp_configuration_v1_manifest_test.go out/terratest-sentinel-crd_projectcalico_org_bgp_filter_v1_manifest_test.go out/terratest-sentinel-crd_projectcalico_org_bgp_peer_v1_manifest_test.go out/terratest-sentinel-crd_projectcalico_org_block_affinity_v1_manifest_test.go out/terratest-sentinel-crd_projectcalico_org_calico_node_status_v1_manifest_test.go out/terratest-sentinel-crd_projectcalico_org_cluster_information_v1_manifest_test.go out/terratest-sentinel-crd_projectcalico_org_felix_configuration_v1_manifest_test.go out/terratest-sentinel-crd_projectcalico_org_global_network_policy_v1_manifest_test.go out/terratest-sentinel-crd_projectcalico_org_global_network_set_v1_manifest_test.go out/terratest-sentinel-crd_projectcalico_org_host_endpoint_v1_manifest_test.go out/terratest-sentinel-crd_projectcalico_org_ipam_block_v1_manifest_test.go out/terratest-sentinel-crd_projectcalico_org_ipam_config_v1_manifest_test.go out/terratest-sentinel-crd_projectcalico_org_ipam_handle_v1_manifest_test.go out/terratest-sentinel-crd_projectcalico_org_ip_pool_v1_manifest_test.go out/terratest-sentinel-crd_projectcalico_org_ip_reservation_v1_manifest_test.go out/terratest-sentinel-crd_projectcalico_org_kube_controllers_configuration_v1_manifest_test.go out/terratest-sentinel-crd_projectcalico_org_network_policy_v1_manifest_test.go out/terratest-sentinel-crd_projectcalico_org_network_set_v1_manifest_test.go out/terratest-sentinel-data_fluid_io_alluxio_runtime_v1alpha1_manifest_test.go out/terratest-sentinel-data_fluid_io_data_backup_v1alpha1_manifest_test.go out/terratest-sentinel-data_fluid_io_data_load_v1alpha1_manifest_test.go out/terratest-sentinel-data_fluid_io_dataset_v1alpha1_manifest_test.go out/terratest-sentinel-data_fluid_io_goose_fs_runtime_v1alpha1_manifest_test.go out/terratest-sentinel-data_fluid_io_jindo_runtime_v1alpha1_manifest_test.go out/terratest-sentinel-data_fluid_io_juice_fs_runtime_v1alpha1_manifest_test.go out/terratest-sentinel-data_fluid_io_thin_runtime_profile_v1alpha1_manifest_test.go out/terratest-sentinel-data_fluid_io_thin_runtime_v1alpha1_manifest_test.go out/terratest-sentinel-databases_schemahero_io_database_v1alpha4_manifest_test.go out/terratest-sentinel-databases_spotahome_com_redis_failover_v1_manifest_test.go out/terratest-sentinel-datadoghq_com_datadog_agent_v1alpha1_manifest_test.go out/terratest-sentinel-datadoghq_com_datadog_metric_v1alpha1_manifest_test.go out/terratest-sentinel-datadoghq_com_datadog_monitor_v1alpha1_manifest_test.go out/terratest-sentinel-datadoghq_com_datadog_slo_v1alpha1_manifest_test.go out/terratest-sentinel-datadoghq_com_datadog_agent_v2alpha1_manifest_test.go out/terratest-sentinel-dataprotection_kubeblocks_io_action_set_v1alpha1_manifest_test.go out/terratest-sentinel-dataprotection_kubeblocks_io_backup_policy_v1alpha1_manifest_test.go out/terratest-sentinel-dataprotection_kubeblocks_io_backup_repo_v1alpha1_manifest_test.go out/terratest-sentinel-dataprotection_kubeblocks_io_backup_schedule_v1alpha1_manifest_test.go out/terratest-sentinel-dataprotection_kubeblocks_io_backup_v1alpha1_manifest_test.go out/terratest-sentinel-dataprotection_kubeblocks_io_restore_v1alpha1_manifest_test.go out/terratest-sentinel-designer_kaoto_io_kaoto_v1alpha1_manifest_test.go out/terratest-sentinel-devices_kubeedge_io_device_model_v1alpha2_manifest_test.go out/terratest-sentinel-devices_kubeedge_io_device_v1alpha2_manifest_test.go out/terratest-sentinel-devices_kubeedge_io_device_model_v1beta1_manifest_test.go out/terratest-sentinel-devices_kubeedge_io_device_v1beta1_manifest_test.go out/terratest-sentinel-devops_kubesphere_io_releaser_controller_v1alpha1_manifest_test.go out/terratest-sentinel-devops_kubesphere_io_releaser_v1alpha1_manifest_test.go out/terratest-sentinel-dex_gpu_ninja_com_dex_identity_provider_v1alpha1_manifest_test.go out/terratest-sentinel-dex_gpu_ninja_com_dex_o_auth2_client_v1alpha1_manifest_test.go out/terratest-sentinel-dex_gpu_ninja_com_dex_user_v1alpha1_manifest_test.go out/terratest-sentinel-digitalis_io_vals_secret_v1_manifest_test.go out/terratest-sentinel-digitalis_io_db_secret_v1beta1_manifest_test.go out/terratest-sentinel-discovery_k8s_io_endpoint_slice_v1_manifest_test.go out/terratest-sentinel-documentdb_services_k8s_aws_db_cluster_v1alpha1_manifest_test.go out/terratest-sentinel-documentdb_services_k8s_aws_db_instance_v1alpha1_manifest_test.go out/terratest-sentinel-documentdb_services_k8s_aws_db_subnet_group_v1alpha1_manifest_test.go out/terratest-sentinel-druid_apache_org_druid_v1alpha1_manifest_test.go out/terratest-sentinel-dynamodb_services_k8s_aws_backup_v1alpha1_manifest_test.go out/terratest-sentinel-dynamodb_services_k8s_aws_global_table_v1alpha1_manifest_test.go out/terratest-sentinel-dynamodb_services_k8s_aws_table_v1alpha1_manifest_test.go out/terratest-sentinel-ec2_services_k8s_aws_dhcp_options_v1alpha1_manifest_test.go out/terratest-sentinel-ec2_services_k8s_aws_elastic_ip_address_v1alpha1_manifest_test.go out/terratest-sentinel-ec2_services_k8s_aws_instance_v1alpha1_manifest_test.go out/terratest-sentinel-ec2_services_k8s_aws_internet_gateway_v1alpha1_manifest_test.go out/terratest-sentinel-ec2_services_k8s_aws_nat_gateway_v1alpha1_manifest_test.go out/terratest-sentinel-ec2_services_k8s_aws_route_table_v1alpha1_manifest_test.go out/terratest-sentinel-ec2_services_k8s_aws_security_group_v1alpha1_manifest_test.go out/terratest-sentinel-ec2_services_k8s_aws_subnet_v1alpha1_manifest_test.go out/terratest-sentinel-ec2_services_k8s_aws_transit_gateway_v1alpha1_manifest_test.go out/terratest-sentinel-ec2_services_k8s_aws_vpc_v1alpha1_manifest_test.go out/terratest-sentinel-ec2_services_k8s_aws_vpc_endpoint_v1alpha1_manifest_test.go out/terratest-sentinel-ecr_services_k8s_aws_pull_through_cache_rule_v1alpha1_manifest_test.go out/terratest-sentinel-ecr_services_k8s_aws_repository_v1alpha1_manifest_test.go out/terratest-sentinel-efs_services_k8s_aws_access_point_v1alpha1_manifest_test.go out/terratest-sentinel-efs_services_k8s_aws_file_system_v1alpha1_manifest_test.go out/terratest-sentinel-efs_services_k8s_aws_mount_target_v1alpha1_manifest_test.go out/terratest-sentinel-eks_services_k8s_aws_addon_v1alpha1_manifest_test.go out/terratest-sentinel-eks_services_k8s_aws_cluster_v1alpha1_manifest_test.go out/terratest-sentinel-eks_services_k8s_aws_fargate_profile_v1alpha1_manifest_test.go out/terratest-sentinel-eks_services_k8s_aws_nodegroup_v1alpha1_manifest_test.go out/terratest-sentinel-elasticache_services_k8s_aws_cache_parameter_group_v1alpha1_manifest_test.go out/terratest-sentinel-elasticache_services_k8s_aws_cache_subnet_group_v1alpha1_manifest_test.go out/terratest-sentinel-elasticache_services_k8s_aws_replication_group_v1alpha1_manifest_test.go out/terratest-sentinel-elasticache_services_k8s_aws_snapshot_v1alpha1_manifest_test.go out/terratest-sentinel-elasticache_services_k8s_aws_user_group_v1alpha1_manifest_test.go out/terratest-sentinel-elasticache_services_k8s_aws_user_v1alpha1_manifest_test.go out/terratest-sentinel-elasticsearch_k8s_elastic_co_elasticsearch_v1_manifest_test.go out/terratest-sentinel-elasticsearch_k8s_elastic_co_elasticsearch_v1beta1_manifest_test.go out/terratest-sentinel-elbv2_k8s_aws_target_group_binding_v1alpha1_manifest_test.go out/terratest-sentinel-elbv2_k8s_aws_ingress_class_params_v1beta1_manifest_test.go out/terratest-sentinel-elbv2_k8s_aws_target_group_binding_v1beta1_manifest_test.go out/terratest-sentinel-emrcontainers_services_k8s_aws_job_run_v1alpha1_manifest_test.go out/terratest-sentinel-emrcontainers_services_k8s_aws_virtual_cluster_v1alpha1_manifest_test.go out/terratest-sentinel-ensembleoss_io_cluster_v1_manifest_test.go out/terratest-sentinel-ensembleoss_io_resource_v1_manifest_test.go out/terratest-sentinel-enterprise_gloo_solo_io_auth_config_v1_manifest_test.go out/terratest-sentinel-enterprisesearch_k8s_elastic_co_enterprise_search_v1_manifest_test.go out/terratest-sentinel-enterprisesearch_k8s_elastic_co_enterprise_search_v1beta1_manifest_test.go out/terratest-sentinel-events_k8s_io_event_v1_manifest_test.go out/terratest-sentinel-everest_percona_com_backup_storage_v1alpha1_manifest_test.go out/terratest-sentinel-everest_percona_com_database_cluster_backup_v1alpha1_manifest_test.go out/terratest-sentinel-everest_percona_com_database_cluster_restore_v1alpha1_manifest_test.go out/terratest-sentinel-everest_percona_com_database_cluster_v1alpha1_manifest_test.go out/terratest-sentinel-everest_percona_com_database_engine_v1alpha1_manifest_test.go out/terratest-sentinel-everest_percona_com_monitoring_config_v1alpha1_manifest_test.go out/terratest-sentinel-execution_furiko_io_job_config_v1alpha1_manifest_test.go out/terratest-sentinel-execution_furiko_io_job_v1alpha1_manifest_test.go out/terratest-sentinel-executor_testkube_io_executor_v1_manifest_test.go out/terratest-sentinel-executor_testkube_io_webhook_v1_manifest_test.go out/terratest-sentinel-expansion_gatekeeper_sh_expansion_template_v1alpha1_manifest_test.go out/terratest-sentinel-expansion_gatekeeper_sh_expansion_template_v1beta1_manifest_test.go out/terratest-sentinel-extensions_istio_io_wasm_plugin_v1alpha1_manifest_test.go out/terratest-sentinel-extensions_kubeblocks_io_addon_v1alpha1_manifest_test.go out/terratest-sentinel-external_secrets_io_cluster_secret_store_v1alpha1_manifest_test.go out/terratest-sentinel-external_secrets_io_external_secret_v1alpha1_manifest_test.go out/terratest-sentinel-external_secrets_io_secret_store_v1alpha1_manifest_test.go out/terratest-sentinel-external_secrets_io_cluster_external_secret_v1beta1_manifest_test.go out/terratest-sentinel-external_secrets_io_cluster_secret_store_v1beta1_manifest_test.go out/terratest-sentinel-external_secrets_io_external_secret_v1beta1_manifest_test.go out/terratest-sentinel-external_secrets_io_secret_store_v1beta1_manifest_test.go out/terratest-sentinel-externaldata_gatekeeper_sh_provider_v1alpha1_manifest_test.go out/terratest-sentinel-externaldata_gatekeeper_sh_provider_v1beta1_manifest_test.go out/terratest-sentinel-externaldns_k8s_io_dns_endpoint_v1alpha1_manifest_test.go out/terratest-sentinel-externaldns_nginx_org_dns_endpoint_v1_manifest_test.go out/terratest-sentinel-fence_agents_remediation_medik8s_io_fence_agents_remediation_template_v1alpha1_manifest_test.go out/terratest-sentinel-fence_agents_remediation_medik8s_io_fence_agents_remediation_v1alpha1_manifest_test.go out/terratest-sentinel-flagger_app_alert_provider_v1beta1_manifest_test.go out/terratest-sentinel-flagger_app_canary_v1beta1_manifest_test.go out/terratest-sentinel-flagger_app_metric_template_v1beta1_manifest_test.go out/terratest-sentinel-flink_apache_org_flink_deployment_v1beta1_manifest_test.go out/terratest-sentinel-flink_apache_org_flink_session_job_v1beta1_manifest_test.go out/terratest-sentinel-flow_volcano_sh_job_flow_v1alpha1_manifest_test.go out/terratest-sentinel-flow_volcano_sh_job_template_v1alpha1_manifest_test.go out/terratest-sentinel-flowcontrol_apiserver_k8s_io_flow_schema_v1beta3_manifest_test.go out/terratest-sentinel-flowcontrol_apiserver_k8s_io_priority_level_configuration_v1beta3_manifest_test.go out/terratest-sentinel-flows_netobserv_io_flow_collector_v1alpha1_manifest_test.go out/terratest-sentinel-flows_netobserv_io_flow_collector_v1beta1_manifest_test.go out/terratest-sentinel-flows_netobserv_io_flow_collector_v1beta2_manifest_test.go out/terratest-sentinel-fluentbit_fluent_io_cluster_filter_v1alpha2_manifest_test.go out/terratest-sentinel-fluentbit_fluent_io_cluster_fluent_bit_config_v1alpha2_manifest_test.go out/terratest-sentinel-fluentbit_fluent_io_cluster_input_v1alpha2_manifest_test.go out/terratest-sentinel-fluentbit_fluent_io_cluster_output_v1alpha2_manifest_test.go out/terratest-sentinel-fluentbit_fluent_io_cluster_parser_v1alpha2_manifest_test.go out/terratest-sentinel-fluentbit_fluent_io_collector_v1alpha2_manifest_test.go out/terratest-sentinel-fluentbit_fluent_io_filter_v1alpha2_manifest_test.go out/terratest-sentinel-fluentbit_fluent_io_fluent_bit_config_v1alpha2_manifest_test.go out/terratest-sentinel-fluentbit_fluent_io_fluent_bit_v1alpha2_manifest_test.go out/terratest-sentinel-fluentbit_fluent_io_output_v1alpha2_manifest_test.go out/terratest-sentinel-fluentbit_fluent_io_parser_v1alpha2_manifest_test.go out/terratest-sentinel-fluentd_fluent_io_cluster_filter_v1alpha1_manifest_test.go out/terratest-sentinel-fluentd_fluent_io_cluster_fluentd_config_v1alpha1_manifest_test.go out/terratest-sentinel-fluentd_fluent_io_cluster_input_v1alpha1_manifest_test.go out/terratest-sentinel-fluentd_fluent_io_cluster_output_v1alpha1_manifest_test.go out/terratest-sentinel-fluentd_fluent_io_filter_v1alpha1_manifest_test.go out/terratest-sentinel-fluentd_fluent_io_fluentd_config_v1alpha1_manifest_test.go out/terratest-sentinel-fluentd_fluent_io_fluentd_v1alpha1_manifest_test.go out/terratest-sentinel-fluentd_fluent_io_input_v1alpha1_manifest_test.go out/terratest-sentinel-fluentd_fluent_io_output_v1alpha1_manifest_test.go out/terratest-sentinel-flux_framework_org_mini_cluster_v1alpha1_manifest_test.go out/terratest-sentinel-flux_framework_org_mini_cluster_v1alpha2_manifest_test.go out/terratest-sentinel-forklift_konveyor_io_forklift_controller_v1beta1_manifest_test.go out/terratest-sentinel-forklift_konveyor_io_hook_v1beta1_manifest_test.go out/terratest-sentinel-forklift_konveyor_io_host_v1beta1_manifest_test.go out/terratest-sentinel-forklift_konveyor_io_migration_v1beta1_manifest_test.go out/terratest-sentinel-forklift_konveyor_io_network_map_v1beta1_manifest_test.go out/terratest-sentinel-forklift_konveyor_io_openstack_volume_populator_v1beta1_manifest_test.go out/terratest-sentinel-forklift_konveyor_io_ovirt_volume_populator_v1beta1_manifest_test.go out/terratest-sentinel-forklift_konveyor_io_plan_v1beta1_manifest_test.go out/terratest-sentinel-forklift_konveyor_io_provider_v1beta1_manifest_test.go out/terratest-sentinel-forklift_konveyor_io_storage_map_v1beta1_manifest_test.go out/terratest-sentinel-fossul_io_backup_config_v1_manifest_test.go out/terratest-sentinel-fossul_io_backup_schedule_v1_manifest_test.go out/terratest-sentinel-fossul_io_backup_v1_manifest_test.go out/terratest-sentinel-fossul_io_fossul_v1_manifest_test.go out/terratest-sentinel-fossul_io_restore_v1_manifest_test.go out/terratest-sentinel-gateway_networking_k8s_io_gateway_class_v1_manifest_test.go out/terratest-sentinel-gateway_networking_k8s_io_gateway_v1_manifest_test.go out/terratest-sentinel-gateway_networking_k8s_io_grpc_route_v1_manifest_test.go out/terratest-sentinel-gateway_networking_k8s_io_http_route_v1_manifest_test.go out/terratest-sentinel-gateway_networking_k8s_io_grpc_route_v1alpha2_manifest_test.go out/terratest-sentinel-gateway_networking_k8s_io_reference_grant_v1alpha2_manifest_test.go out/terratest-sentinel-gateway_networking_k8s_io_tcp_route_v1alpha2_manifest_test.go out/terratest-sentinel-gateway_networking_k8s_io_tls_route_v1alpha2_manifest_test.go out/terratest-sentinel-gateway_networking_k8s_io_udp_route_v1alpha2_manifest_test.go out/terratest-sentinel-gateway_networking_k8s_io_gateway_class_v1beta1_manifest_test.go out/terratest-sentinel-gateway_networking_k8s_io_gateway_v1beta1_manifest_test.go out/terratest-sentinel-gateway_networking_k8s_io_http_route_v1beta1_manifest_test.go out/terratest-sentinel-gateway_networking_k8s_io_reference_grant_v1beta1_manifest_test.go out/terratest-sentinel-gateway_nginx_org_client_settings_policy_v1alpha1_manifest_test.go out/terratest-sentinel-gateway_nginx_org_nginx_gateway_v1alpha1_manifest_test.go out/terratest-sentinel-gateway_nginx_org_nginx_proxy_v1alpha1_manifest_test.go out/terratest-sentinel-gateway_nginx_org_observability_policy_v1alpha1_manifest_test.go out/terratest-sentinel-gateway_solo_io_gateway_v1_manifest_test.go out/terratest-sentinel-gateway_solo_io_matchable_http_gateway_v1_manifest_test.go out/terratest-sentinel-gateway_solo_io_route_option_v1_manifest_test.go out/terratest-sentinel-gateway_solo_io_route_table_v1_manifest_test.go out/terratest-sentinel-gateway_solo_io_virtual_host_option_v1_manifest_test.go out/terratest-sentinel-gateway_solo_io_virtual_service_v1_manifest_test.go out/terratest-sentinel-getambassador_io_auth_service_v1_manifest_test.go out/terratest-sentinel-getambassador_io_consul_resolver_v1_manifest_test.go out/terratest-sentinel-getambassador_io_dev_portal_v1_manifest_test.go out/terratest-sentinel-getambassador_io_kubernetes_endpoint_resolver_v1_manifest_test.go out/terratest-sentinel-getambassador_io_kubernetes_service_resolver_v1_manifest_test.go out/terratest-sentinel-getambassador_io_log_service_v1_manifest_test.go out/terratest-sentinel-getambassador_io_mapping_v1_manifest_test.go out/terratest-sentinel-getambassador_io_module_v1_manifest_test.go out/terratest-sentinel-getambassador_io_rate_limit_service_v1_manifest_test.go out/terratest-sentinel-getambassador_io_tcp_mapping_v1_manifest_test.go out/terratest-sentinel-getambassador_io_tls_context_v1_manifest_test.go out/terratest-sentinel-getambassador_io_tracing_service_v1_manifest_test.go out/terratest-sentinel-getambassador_io_auth_service_v2_manifest_test.go out/terratest-sentinel-getambassador_io_consul_resolver_v2_manifest_test.go out/terratest-sentinel-getambassador_io_dev_portal_v2_manifest_test.go out/terratest-sentinel-getambassador_io_host_v2_manifest_test.go out/terratest-sentinel-getambassador_io_kubernetes_endpoint_resolver_v2_manifest_test.go out/terratest-sentinel-getambassador_io_kubernetes_service_resolver_v2_manifest_test.go out/terratest-sentinel-getambassador_io_log_service_v2_manifest_test.go out/terratest-sentinel-getambassador_io_mapping_v2_manifest_test.go out/terratest-sentinel-getambassador_io_module_v2_manifest_test.go out/terratest-sentinel-getambassador_io_rate_limit_service_v2_manifest_test.go out/terratest-sentinel-getambassador_io_tcp_mapping_v2_manifest_test.go out/terratest-sentinel-getambassador_io_tls_context_v2_manifest_test.go out/terratest-sentinel-getambassador_io_tracing_service_v2_manifest_test.go out/terratest-sentinel-getambassador_io_auth_service_v3alpha1_manifest_test.go out/terratest-sentinel-getambassador_io_consul_resolver_v3alpha1_manifest_test.go out/terratest-sentinel-getambassador_io_dev_portal_v3alpha1_manifest_test.go out/terratest-sentinel-getambassador_io_host_v3alpha1_manifest_test.go out/terratest-sentinel-getambassador_io_kubernetes_endpoint_resolver_v3alpha1_manifest_test.go out/terratest-sentinel-getambassador_io_kubernetes_service_resolver_v3alpha1_manifest_test.go out/terratest-sentinel-getambassador_io_listener_v3alpha1_manifest_test.go out/terratest-sentinel-getambassador_io_log_service_v3alpha1_manifest_test.go out/terratest-sentinel-getambassador_io_mapping_v3alpha1_manifest_test.go out/terratest-sentinel-getambassador_io_module_v3alpha1_manifest_test.go out/terratest-sentinel-getambassador_io_rate_limit_service_v3alpha1_manifest_test.go out/terratest-sentinel-getambassador_io_tcp_mapping_v3alpha1_manifest_test.go out/terratest-sentinel-getambassador_io_tls_context_v3alpha1_manifest_test.go out/terratest-sentinel-getambassador_io_tracing_service_v3alpha1_manifest_test.go out/terratest-sentinel-gitops_hybrid_cloud_patterns_io_pattern_v1alpha1_manifest_test.go out/terratest-sentinel-gloo_solo_io_proxy_v1_manifest_test.go out/terratest-sentinel-gloo_solo_io_settings_v1_manifest_test.go out/terratest-sentinel-gloo_solo_io_upstream_group_v1_manifest_test.go out/terratest-sentinel-gloo_solo_io_upstream_v1_manifest_test.go out/terratest-sentinel-grafana_integreatly_org_grafana_dashboard_v1beta1_manifest_test.go out/terratest-sentinel-grafana_integreatly_org_grafana_datasource_v1beta1_manifest_test.go out/terratest-sentinel-grafana_integreatly_org_grafana_folder_v1beta1_manifest_test.go out/terratest-sentinel-grafana_integreatly_org_grafana_v1beta1_manifest_test.go out/terratest-sentinel-graphql_gloo_solo_io_graph_ql_api_v1beta1_manifest_test.go out/terratest-sentinel-groupsnapshot_storage_k8s_io_volume_group_snapshot_class_v1alpha1_manifest_test.go out/terratest-sentinel-groupsnapshot_storage_k8s_io_volume_group_snapshot_content_v1alpha1_manifest_test.go out/terratest-sentinel-groupsnapshot_storage_k8s_io_volume_group_snapshot_v1alpha1_manifest_test.go out/terratest-sentinel-hazelcast_com_cron_hot_backup_v1alpha1_manifest_test.go out/terratest-sentinel-hazelcast_com_hazelcast_v1alpha1_manifest_test.go out/terratest-sentinel-hazelcast_com_hot_backup_v1alpha1_manifest_test.go out/terratest-sentinel-hazelcast_com_management_center_v1alpha1_manifest_test.go out/terratest-sentinel-hazelcast_com_map_v1alpha1_manifest_test.go out/terratest-sentinel-hazelcast_com_wan_replication_v1alpha1_manifest_test.go out/terratest-sentinel-helm_sigstore_dev_rekor_v1alpha1_manifest_test.go out/terratest-sentinel-helm_toolkit_fluxcd_io_helm_release_v2_manifest_test.go out/terratest-sentinel-helm_toolkit_fluxcd_io_helm_release_v2beta1_manifest_test.go out/terratest-sentinel-helm_toolkit_fluxcd_io_helm_release_v2beta2_manifest_test.go out/terratest-sentinel-hive_openshift_io_checkpoint_v1_manifest_test.go out/terratest-sentinel-hive_openshift_io_cluster_claim_v1_manifest_test.go out/terratest-sentinel-hive_openshift_io_cluster_deployment_customization_v1_manifest_test.go out/terratest-sentinel-hive_openshift_io_cluster_deployment_v1_manifest_test.go out/terratest-sentinel-hive_openshift_io_cluster_deprovision_v1_manifest_test.go out/terratest-sentinel-hive_openshift_io_cluster_image_set_v1_manifest_test.go out/terratest-sentinel-hive_openshift_io_cluster_pool_v1_manifest_test.go out/terratest-sentinel-hive_openshift_io_cluster_provision_v1_manifest_test.go out/terratest-sentinel-hive_openshift_io_cluster_relocate_v1_manifest_test.go out/terratest-sentinel-hive_openshift_io_cluster_state_v1_manifest_test.go out/terratest-sentinel-hive_openshift_io_dns_zone_v1_manifest_test.go out/terratest-sentinel-hive_openshift_io_hive_config_v1_manifest_test.go out/terratest-sentinel-hive_openshift_io_machine_pool_name_lease_v1_manifest_test.go out/terratest-sentinel-hive_openshift_io_machine_pool_v1_manifest_test.go out/terratest-sentinel-hive_openshift_io_selector_sync_identity_provider_v1_manifest_test.go out/terratest-sentinel-hive_openshift_io_selector_sync_set_v1_manifest_test.go out/terratest-sentinel-hive_openshift_io_sync_identity_provider_v1_manifest_test.go out/terratest-sentinel-hive_openshift_io_sync_set_v1_manifest_test.go out/terratest-sentinel-hiveinternal_openshift_io_cluster_sync_lease_v1alpha1_manifest_test.go out/terratest-sentinel-hiveinternal_openshift_io_cluster_sync_v1alpha1_manifest_test.go out/terratest-sentinel-hiveinternal_openshift_io_fake_cluster_install_v1alpha1_manifest_test.go out/terratest-sentinel-hnc_x_k8s_io_hierarchical_resource_quota_v1alpha2_manifest_test.go out/terratest-sentinel-hnc_x_k8s_io_hierarchy_configuration_v1alpha2_manifest_test.go out/terratest-sentinel-hnc_x_k8s_io_hnc_configuration_v1alpha2_manifest_test.go out/terratest-sentinel-hnc_x_k8s_io_subnamespace_anchor_v1alpha2_manifest_test.go out/terratest-sentinel-hyperfoil_io_horreum_v1alpha1_manifest_test.go out/terratest-sentinel-hyperfoil_io_hyperfoil_v1alpha2_manifest_test.go out/terratest-sentinel-iam_services_k8s_aws_group_v1alpha1_manifest_test.go out/terratest-sentinel-iam_services_k8s_aws_instance_profile_v1alpha1_manifest_test.go out/terratest-sentinel-iam_services_k8s_aws_open_id_connect_provider_v1alpha1_manifest_test.go out/terratest-sentinel-iam_services_k8s_aws_policy_v1alpha1_manifest_test.go out/terratest-sentinel-iam_services_k8s_aws_role_v1alpha1_manifest_test.go out/terratest-sentinel-iam_services_k8s_aws_user_v1alpha1_manifest_test.go out/terratest-sentinel-ibmcloud_ibm_com_composable_v1alpha1_manifest_test.go out/terratest-sentinel-image_toolkit_fluxcd_io_image_policy_v1beta1_manifest_test.go out/terratest-sentinel-image_toolkit_fluxcd_io_image_repository_v1beta1_manifest_test.go out/terratest-sentinel-image_toolkit_fluxcd_io_image_update_automation_v1beta1_manifest_test.go out/terratest-sentinel-image_toolkit_fluxcd_io_image_policy_v1beta2_manifest_test.go out/terratest-sentinel-image_toolkit_fluxcd_io_image_repository_v1beta2_manifest_test.go out/terratest-sentinel-image_toolkit_fluxcd_io_image_update_automation_v1beta2_manifest_test.go out/terratest-sentinel-imaging_ingestion_alvearie_org_dicom_event_bridge_v1alpha1_manifest_test.go out/terratest-sentinel-imaging_ingestion_alvearie_org_dicom_event_driven_ingestion_v1alpha1_manifest_test.go out/terratest-sentinel-imaging_ingestion_alvearie_org_dicom_instance_binding_v1alpha1_manifest_test.go out/terratest-sentinel-imaging_ingestion_alvearie_org_dicom_study_binding_v1alpha1_manifest_test.go out/terratest-sentinel-imaging_ingestion_alvearie_org_dicomweb_ingestion_service_v1alpha1_manifest_test.go out/terratest-sentinel-imaging_ingestion_alvearie_org_dimse_ingestion_service_v1alpha1_manifest_test.go out/terratest-sentinel-imaging_ingestion_alvearie_org_dimse_proxy_v1alpha1_manifest_test.go out/terratest-sentinel-inference_kubedl_io_elastic_batch_job_v1alpha1_manifest_test.go out/terratest-sentinel-infinispan_org_infinispan_v1_manifest_test.go out/terratest-sentinel-infinispan_org_backup_v2alpha1_manifest_test.go out/terratest-sentinel-infinispan_org_batch_v2alpha1_manifest_test.go out/terratest-sentinel-infinispan_org_cache_v2alpha1_manifest_test.go out/terratest-sentinel-infinispan_org_restore_v2alpha1_manifest_test.go out/terratest-sentinel-infra_contrib_fluxcd_io_terraform_v1alpha1_manifest_test.go out/terratest-sentinel-infra_contrib_fluxcd_io_terraform_v1alpha2_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_kubevirt_cluster_template_v1alpha1_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_kubevirt_cluster_v1alpha1_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_kubevirt_machine_template_v1alpha1_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_kubevirt_machine_v1alpha1_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_cluster_identity_v1alpha3_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_cluster_v1alpha3_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_deployment_zone_v1alpha3_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_failure_domain_v1alpha3_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_machine_template_v1alpha3_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_machine_v1alpha3_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_vm_v1alpha3_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_cluster_identity_v1alpha4_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_cluster_template_v1alpha4_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_cluster_v1alpha4_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_deployment_zone_v1alpha4_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_failure_domain_v1alpha4_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_machine_template_v1alpha4_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_machine_v1alpha4_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_vm_v1alpha4_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_ibm_power_vs_cluster_template_v1beta1_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_ibm_power_vs_cluster_v1beta1_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_ibm_power_vs_image_v1beta1_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_ibm_power_vs_machine_template_v1beta1_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_ibm_power_vs_machine_v1beta1_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_ibmvpc_cluster_v1beta1_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_ibmvpc_machine_template_v1beta1_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_ibmvpc_machine_v1beta1_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_tinkerbell_cluster_v1beta1_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_tinkerbell_machine_template_v1beta1_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_tinkerbell_machine_v1beta1_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_cluster_identity_v1beta1_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_cluster_template_v1beta1_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_cluster_v1beta1_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_deployment_zone_v1beta1_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_failure_domain_v1beta1_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_machine_template_v1beta1_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_machine_v1beta1_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_v_sphere_vm_v1beta1_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_ibm_power_vs_cluster_template_v1beta2_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_ibm_power_vs_cluster_v1beta2_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_ibm_power_vs_image_v1beta2_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_ibm_power_vs_machine_template_v1beta2_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_ibm_power_vs_machine_v1beta2_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_ibmvpc_cluster_v1beta2_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_ibmvpc_machine_template_v1beta2_manifest_test.go out/terratest-sentinel-infrastructure_cluster_x_k8s_io_ibmvpc_machine_v1beta2_manifest_test.go out/terratest-sentinel-installation_mattermost_com_mattermost_v1beta1_manifest_test.go out/terratest-sentinel-instana_io_instana_agent_v1_manifest_test.go out/terratest-sentinel-integration_rock8s_com_deferred_resource_v1beta1_manifest_test.go out/terratest-sentinel-integration_rock8s_com_plug_v1beta1_manifest_test.go out/terratest-sentinel-integration_rock8s_com_socket_v1beta1_manifest_test.go out/terratest-sentinel-iot_eclipse_org_ditto_v1alpha1_manifest_test.go out/terratest-sentinel-iot_eclipse_org_hawkbit_v1alpha1_manifest_test.go out/terratest-sentinel-ipam_cluster_x_k8s_io_ip_address_claim_v1alpha1_manifest_test.go out/terratest-sentinel-ipam_cluster_x_k8s_io_ip_address_v1alpha1_manifest_test.go out/terratest-sentinel-ipam_cluster_x_k8s_io_ip_address_claim_v1beta1_manifest_test.go out/terratest-sentinel-ipam_cluster_x_k8s_io_ip_address_v1beta1_manifest_test.go out/terratest-sentinel-isindir_github_com_sops_secret_v1alpha1_manifest_test.go out/terratest-sentinel-isindir_github_com_sops_secret_v1alpha2_manifest_test.go out/terratest-sentinel-isindir_github_com_sops_secret_v1alpha3_manifest_test.go out/terratest-sentinel-jaegertracing_io_jaeger_v1_manifest_test.go out/terratest-sentinel-jobset_x_k8s_io_job_set_v1alpha2_manifest_test.go out/terratest-sentinel-jobsmanager_raczylo_com_managed_job_v1beta1_manifest_test.go out/terratest-sentinel-k6_io_k6_v1alpha1_manifest_test.go out/terratest-sentinel-k6_io_private_load_zone_v1alpha1_manifest_test.go out/terratest-sentinel-k6_io_test_run_v1alpha1_manifest_test.go out/terratest-sentinel-k8gb_absa_oss_gslb_v1beta1_manifest_test.go out/terratest-sentinel-k8s_keycloak_org_keycloak_realm_import_v2alpha1_manifest_test.go out/terratest-sentinel-k8s_keycloak_org_keycloak_v2alpha1_manifest_test.go out/terratest-sentinel-k8s_mariadb_com_backup_v1alpha1_manifest_test.go out/terratest-sentinel-k8s_mariadb_com_connection_v1alpha1_manifest_test.go out/terratest-sentinel-k8s_mariadb_com_database_v1alpha1_manifest_test.go out/terratest-sentinel-k8s_mariadb_com_grant_v1alpha1_manifest_test.go out/terratest-sentinel-k8s_mariadb_com_maria_db_v1alpha1_manifest_test.go out/terratest-sentinel-k8s_mariadb_com_max_scale_v1alpha1_manifest_test.go out/terratest-sentinel-k8s_mariadb_com_restore_v1alpha1_manifest_test.go out/terratest-sentinel-k8s_mariadb_com_sql_job_v1alpha1_manifest_test.go out/terratest-sentinel-k8s_mariadb_com_user_v1alpha1_manifest_test.go out/terratest-sentinel-k8s_nginx_org_global_configuration_v1_manifest_test.go out/terratest-sentinel-k8s_nginx_org_policy_v1_manifest_test.go out/terratest-sentinel-k8s_nginx_org_transport_server_v1_manifest_test.go out/terratest-sentinel-k8s_nginx_org_virtual_server_route_v1_manifest_test.go out/terratest-sentinel-k8s_nginx_org_virtual_server_v1_manifest_test.go out/terratest-sentinel-k8s_nginx_org_global_configuration_v1alpha1_manifest_test.go out/terratest-sentinel-k8s_nginx_org_policy_v1alpha1_manifest_test.go out/terratest-sentinel-k8s_nginx_org_transport_server_v1alpha1_manifest_test.go out/terratest-sentinel-k8s_otterize_com_client_intents_v1alpha2_manifest_test.go out/terratest-sentinel-k8s_otterize_com_kafka_server_config_v1alpha2_manifest_test.go out/terratest-sentinel-k8s_otterize_com_protected_service_v1alpha2_manifest_test.go out/terratest-sentinel-k8s_otterize_com_client_intents_v1alpha3_manifest_test.go out/terratest-sentinel-k8s_otterize_com_kafka_server_config_v1alpha3_manifest_test.go out/terratest-sentinel-k8s_otterize_com_protected_service_v1alpha3_manifest_test.go out/terratest-sentinel-k8up_io_archive_v1_manifest_test.go out/terratest-sentinel-k8up_io_backup_v1_manifest_test.go out/terratest-sentinel-k8up_io_check_v1_manifest_test.go out/terratest-sentinel-k8up_io_pre_backup_pod_v1_manifest_test.go out/terratest-sentinel-k8up_io_prune_v1_manifest_test.go out/terratest-sentinel-k8up_io_restore_v1_manifest_test.go out/terratest-sentinel-k8up_io_schedule_v1_manifest_test.go out/terratest-sentinel-k8up_io_snapshot_v1_manifest_test.go out/terratest-sentinel-kafka_banzaicloud_io_cruise_control_operation_v1alpha1_manifest_test.go out/terratest-sentinel-kafka_banzaicloud_io_kafka_topic_v1alpha1_manifest_test.go out/terratest-sentinel-kafka_banzaicloud_io_kafka_user_v1alpha1_manifest_test.go out/terratest-sentinel-kafka_banzaicloud_io_kafka_cluster_v1beta1_manifest_test.go out/terratest-sentinel-kafka_services_k8s_aws_cluster_v1alpha1_manifest_test.go out/terratest-sentinel-kafka_strimzi_io_kafka_topic_v1alpha1_manifest_test.go out/terratest-sentinel-kafka_strimzi_io_kafka_user_v1alpha1_manifest_test.go out/terratest-sentinel-kafka_strimzi_io_kafka_topic_v1beta1_manifest_test.go out/terratest-sentinel-kafka_strimzi_io_kafka_user_v1beta1_manifest_test.go out/terratest-sentinel-kafka_strimzi_io_kafka_bridge_v1beta2_manifest_test.go out/terratest-sentinel-kafka_strimzi_io_kafka_connect_v1beta2_manifest_test.go out/terratest-sentinel-kafka_strimzi_io_kafka_connector_v1beta2_manifest_test.go out/terratest-sentinel-kafka_strimzi_io_kafka_mirror_maker2_v1beta2_manifest_test.go out/terratest-sentinel-kafka_strimzi_io_kafka_mirror_maker_v1beta2_manifest_test.go out/terratest-sentinel-kafka_strimzi_io_kafka_rebalance_v1beta2_manifest_test.go out/terratest-sentinel-kafka_strimzi_io_kafka_topic_v1beta2_manifest_test.go out/terratest-sentinel-kafka_strimzi_io_kafka_user_v1beta2_manifest_test.go out/terratest-sentinel-kafka_strimzi_io_kafka_v1beta2_manifest_test.go out/terratest-sentinel-kamaji_clastix_io_data_store_v1alpha1_manifest_test.go out/terratest-sentinel-kamaji_clastix_io_tenant_control_plane_v1alpha1_manifest_test.go out/terratest-sentinel-karpenter_k8s_aws_ec2_node_class_v1_manifest_test.go out/terratest-sentinel-karpenter_k8s_aws_ec2_node_class_v1beta1_manifest_test.go out/terratest-sentinel-karpenter_sh_node_claim_v1_manifest_test.go out/terratest-sentinel-karpenter_sh_node_pool_v1_manifest_test.go out/terratest-sentinel-karpenter_sh_node_claim_v1beta1_manifest_test.go out/terratest-sentinel-karpenter_sh_node_pool_v1beta1_manifest_test.go out/terratest-sentinel-keda_sh_cluster_trigger_authentication_v1alpha1_manifest_test.go out/terratest-sentinel-keda_sh_scaled_job_v1alpha1_manifest_test.go out/terratest-sentinel-keda_sh_scaled_object_v1alpha1_manifest_test.go out/terratest-sentinel-keda_sh_trigger_authentication_v1alpha1_manifest_test.go out/terratest-sentinel-keycloak_k8s_reddec_net_keycloak_client_v1alpha1_manifest_test.go out/terratest-sentinel-keycloak_org_keycloak_backup_v1alpha1_manifest_test.go out/terratest-sentinel-keycloak_org_keycloak_client_v1alpha1_manifest_test.go out/terratest-sentinel-keycloak_org_keycloak_realm_v1alpha1_manifest_test.go out/terratest-sentinel-keycloak_org_keycloak_user_v1alpha1_manifest_test.go out/terratest-sentinel-keycloak_org_keycloak_v1alpha1_manifest_test.go out/terratest-sentinel-keyspaces_services_k8s_aws_keyspace_v1alpha1_manifest_test.go out/terratest-sentinel-keyspaces_services_k8s_aws_table_v1alpha1_manifest_test.go out/terratest-sentinel-kiali_io_kiali_v1alpha1_manifest_test.go out/terratest-sentinel-kibana_k8s_elastic_co_kibana_v1_manifest_test.go out/terratest-sentinel-kibana_k8s_elastic_co_kibana_v1beta1_manifest_test.go out/terratest-sentinel-kinesis_services_k8s_aws_stream_v1alpha1_manifest_test.go out/terratest-sentinel-kmm_sigs_x_k8s_io_module_v1beta1_manifest_test.go out/terratest-sentinel-kmm_sigs_x_k8s_io_node_modules_config_v1beta1_manifest_test.go out/terratest-sentinel-kmm_sigs_x_k8s_io_preflight_validation_v1beta1_manifest_test.go out/terratest-sentinel-kmm_sigs_x_k8s_io_preflight_validation_v1beta2_manifest_test.go out/terratest-sentinel-kms_services_k8s_aws_alias_v1alpha1_manifest_test.go out/terratest-sentinel-kms_services_k8s_aws_grant_v1alpha1_manifest_test.go out/terratest-sentinel-kms_services_k8s_aws_key_v1alpha1_manifest_test.go out/terratest-sentinel-kuadrant_io_dns_record_v1alpha1_manifest_test.go out/terratest-sentinel-kuadrant_io_managed_zone_v1alpha1_manifest_test.go out/terratest-sentinel-kuadrant_io_kuadrant_v1beta1_manifest_test.go out/terratest-sentinel-kuadrant_io_auth_policy_v1beta2_manifest_test.go out/terratest-sentinel-kuadrant_io_rate_limit_policy_v1beta2_manifest_test.go out/terratest-sentinel-kube_green_com_sleep_info_v1alpha1_manifest_test.go out/terratest-sentinel-kubean_io_cluster_operation_v1alpha1_manifest_test.go out/terratest-sentinel-kubean_io_cluster_v1alpha1_manifest_test.go out/terratest-sentinel-kubean_io_local_artifact_set_v1alpha1_manifest_test.go out/terratest-sentinel-kubean_io_manifest_v1alpha1_manifest_test.go out/terratest-sentinel-kubecost_com_turndown_schedule_v1alpha1_manifest_test.go out/terratest-sentinel-kubevious_io_workload_profile_v1alpha1_manifest_test.go out/terratest-sentinel-kubevious_io_workload_v1alpha1_manifest_test.go out/terratest-sentinel-kueue_x_k8s_io_admission_check_v1beta1_manifest_test.go out/terratest-sentinel-kueue_x_k8s_io_cluster_queue_v1beta1_manifest_test.go out/terratest-sentinel-kueue_x_k8s_io_local_queue_v1beta1_manifest_test.go out/terratest-sentinel-kueue_x_k8s_io_resource_flavor_v1beta1_manifest_test.go out/terratest-sentinel-kueue_x_k8s_io_workload_v1beta1_manifest_test.go out/terratest-sentinel-kuma_io_circuit_breaker_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_container_patch_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_dataplane_insight_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_dataplane_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_external_service_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_fault_injection_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_health_check_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_mesh_access_log_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_mesh_circuit_breaker_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_mesh_fault_injection_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_mesh_gateway_config_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_mesh_gateway_instance_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_mesh_gateway_route_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_mesh_gateway_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_mesh_health_check_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_mesh_http_route_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_mesh_insight_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_mesh_load_balancing_strategy_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_mesh_proxy_patch_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_mesh_rate_limit_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_mesh_retry_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_mesh_tcp_route_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_mesh_timeout_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_mesh_trace_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_mesh_traffic_permission_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_mesh_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_proxy_template_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_rate_limit_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_retry_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_service_insight_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_timeout_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_traffic_log_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_traffic_permission_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_traffic_route_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_traffic_trace_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_virtual_outbound_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_zone_egress_insight_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_zone_egress_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_zone_ingress_insight_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_zone_ingress_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_zone_insight_v1alpha1_manifest_test.go out/terratest-sentinel-kuma_io_zone_v1alpha1_manifest_test.go out/terratest-sentinel-kustomize_toolkit_fluxcd_io_kustomization_v1_manifest_test.go out/terratest-sentinel-kustomize_toolkit_fluxcd_io_kustomization_v1beta1_manifest_test.go out/terratest-sentinel-kustomize_toolkit_fluxcd_io_kustomization_v1beta2_manifest_test.go out/terratest-sentinel-kyverno_io_cluster_policy_v1_manifest_test.go out/terratest-sentinel-kyverno_io_policy_v1_manifest_test.go out/terratest-sentinel-kyverno_io_admission_report_v1alpha2_manifest_test.go out/terratest-sentinel-kyverno_io_background_scan_report_v1alpha2_manifest_test.go out/terratest-sentinel-kyverno_io_cluster_admission_report_v1alpha2_manifest_test.go out/terratest-sentinel-kyverno_io_cluster_background_scan_report_v1alpha2_manifest_test.go out/terratest-sentinel-kyverno_io_update_request_v1beta1_manifest_test.go out/terratest-sentinel-kyverno_io_admission_report_v2_manifest_test.go out/terratest-sentinel-kyverno_io_background_scan_report_v2_manifest_test.go out/terratest-sentinel-kyverno_io_cleanup_policy_v2_manifest_test.go out/terratest-sentinel-kyverno_io_cluster_admission_report_v2_manifest_test.go out/terratest-sentinel-kyverno_io_cluster_background_scan_report_v2_manifest_test.go out/terratest-sentinel-kyverno_io_cluster_cleanup_policy_v2_manifest_test.go out/terratest-sentinel-kyverno_io_policy_exception_v2_manifest_test.go out/terratest-sentinel-kyverno_io_update_request_v2_manifest_test.go out/terratest-sentinel-kyverno_io_cleanup_policy_v2alpha1_manifest_test.go out/terratest-sentinel-kyverno_io_cluster_cleanup_policy_v2alpha1_manifest_test.go out/terratest-sentinel-kyverno_io_global_context_entry_v2alpha1_manifest_test.go out/terratest-sentinel-kyverno_io_policy_exception_v2alpha1_manifest_test.go out/terratest-sentinel-kyverno_io_cleanup_policy_v2beta1_manifest_test.go out/terratest-sentinel-kyverno_io_cluster_cleanup_policy_v2beta1_manifest_test.go out/terratest-sentinel-kyverno_io_cluster_policy_v2beta1_manifest_test.go out/terratest-sentinel-kyverno_io_policy_exception_v2beta1_manifest_test.go out/terratest-sentinel-kyverno_io_policy_v2beta1_manifest_test.go out/terratest-sentinel-lambda_services_k8s_aws_alias_v1alpha1_manifest_test.go out/terratest-sentinel-lambda_services_k8s_aws_code_signing_config_v1alpha1_manifest_test.go out/terratest-sentinel-lambda_services_k8s_aws_event_source_mapping_v1alpha1_manifest_test.go out/terratest-sentinel-lambda_services_k8s_aws_function_url_config_v1alpha1_manifest_test.go out/terratest-sentinel-lambda_services_k8s_aws_function_v1alpha1_manifest_test.go out/terratest-sentinel-lambda_services_k8s_aws_layer_version_v1alpha1_manifest_test.go out/terratest-sentinel-lambda_services_k8s_aws_version_v1alpha1_manifest_test.go out/terratest-sentinel-lb_lbconfig_carlosedp_com_external_load_balancer_v1_manifest_test.go out/terratest-sentinel-leaksignal_com_cluster_leaksignal_istio_v1_manifest_test.go out/terratest-sentinel-leaksignal_com_leaksignal_istio_v1_manifest_test.go out/terratest-sentinel-lerentis_uploadfilter24_eu_bitwarden_secret_v1beta4_manifest_test.go out/terratest-sentinel-lerentis_uploadfilter24_eu_bitwarden_template_v1beta4_manifest_test.go out/terratest-sentinel-lerentis_uploadfilter24_eu_registry_credential_v1beta4_manifest_test.go out/terratest-sentinel-lerentis_uploadfilter24_eu_bitwarden_secret_v1beta5_manifest_test.go out/terratest-sentinel-lerentis_uploadfilter24_eu_bitwarden_template_v1beta5_manifest_test.go out/terratest-sentinel-lerentis_uploadfilter24_eu_registry_credential_v1beta5_manifest_test.go out/terratest-sentinel-lerentis_uploadfilter24_eu_bitwarden_secret_v1beta6_manifest_test.go out/terratest-sentinel-lerentis_uploadfilter24_eu_bitwarden_template_v1beta6_manifest_test.go out/terratest-sentinel-lerentis_uploadfilter24_eu_registry_credential_v1beta6_manifest_test.go out/terratest-sentinel-lerentis_uploadfilter24_eu_bitwarden_secret_v1beta7_manifest_test.go out/terratest-sentinel-lerentis_uploadfilter24_eu_bitwarden_template_v1beta7_manifest_test.go out/terratest-sentinel-lerentis_uploadfilter24_eu_registry_credential_v1beta7_manifest_test.go out/terratest-sentinel-limitador_kuadrant_io_limitador_v1alpha1_manifest_test.go out/terratest-sentinel-litmuschaos_io_chaos_engine_v1alpha1_manifest_test.go out/terratest-sentinel-litmuschaos_io_chaos_experiment_v1alpha1_manifest_test.go out/terratest-sentinel-litmuschaos_io_chaos_result_v1alpha1_manifest_test.go out/terratest-sentinel-logging_banzaicloud_io_cluster_flow_v1alpha1_manifest_test.go out/terratest-sentinel-logging_banzaicloud_io_cluster_output_v1alpha1_manifest_test.go out/terratest-sentinel-logging_banzaicloud_io_flow_v1alpha1_manifest_test.go out/terratest-sentinel-logging_banzaicloud_io_logging_v1alpha1_manifest_test.go out/terratest-sentinel-logging_banzaicloud_io_output_v1alpha1_manifest_test.go out/terratest-sentinel-logging_banzaicloud_io_cluster_flow_v1beta1_manifest_test.go out/terratest-sentinel-logging_banzaicloud_io_cluster_output_v1beta1_manifest_test.go out/terratest-sentinel-logging_banzaicloud_io_flow_v1beta1_manifest_test.go out/terratest-sentinel-logging_banzaicloud_io_fluentbit_agent_v1beta1_manifest_test.go out/terratest-sentinel-logging_banzaicloud_io_logging_v1beta1_manifest_test.go out/terratest-sentinel-logging_banzaicloud_io_node_agent_v1beta1_manifest_test.go out/terratest-sentinel-logging_banzaicloud_io_output_v1beta1_manifest_test.go out/terratest-sentinel-logging_banzaicloud_io_syslog_ng_cluster_flow_v1beta1_manifest_test.go out/terratest-sentinel-logging_banzaicloud_io_syslog_ng_cluster_output_v1beta1_manifest_test.go out/terratest-sentinel-logging_banzaicloud_io_syslog_ng_flow_v1beta1_manifest_test.go out/terratest-sentinel-logging_banzaicloud_io_syslog_ng_output_v1beta1_manifest_test.go out/terratest-sentinel-logging_extensions_banzaicloud_io_event_tailer_v1alpha1_manifest_test.go out/terratest-sentinel-logging_extensions_banzaicloud_io_host_tailer_v1alpha1_manifest_test.go out/terratest-sentinel-loki_grafana_com_alerting_rule_v1_manifest_test.go out/terratest-sentinel-loki_grafana_com_loki_stack_v1_manifest_test.go out/terratest-sentinel-loki_grafana_com_recording_rule_v1_manifest_test.go out/terratest-sentinel-loki_grafana_com_ruler_config_v1_manifest_test.go out/terratest-sentinel-loki_grafana_com_alerting_rule_v1beta1_manifest_test.go out/terratest-sentinel-loki_grafana_com_loki_stack_v1beta1_manifest_test.go out/terratest-sentinel-loki_grafana_com_recording_rule_v1beta1_manifest_test.go out/terratest-sentinel-loki_grafana_com_ruler_config_v1beta1_manifest_test.go out/terratest-sentinel-longhorn_io_backing_image_data_source_v1beta1_manifest_test.go out/terratest-sentinel-longhorn_io_backing_image_manager_v1beta1_manifest_test.go out/terratest-sentinel-longhorn_io_backing_image_v1beta1_manifest_test.go out/terratest-sentinel-longhorn_io_backup_target_v1beta1_manifest_test.go out/terratest-sentinel-longhorn_io_backup_v1beta1_manifest_test.go out/terratest-sentinel-longhorn_io_backup_volume_v1beta1_manifest_test.go out/terratest-sentinel-longhorn_io_engine_image_v1beta1_manifest_test.go out/terratest-sentinel-longhorn_io_engine_v1beta1_manifest_test.go out/terratest-sentinel-longhorn_io_instance_manager_v1beta1_manifest_test.go out/terratest-sentinel-longhorn_io_node_v1beta1_manifest_test.go out/terratest-sentinel-longhorn_io_recurring_job_v1beta1_manifest_test.go out/terratest-sentinel-longhorn_io_replica_v1beta1_manifest_test.go out/terratest-sentinel-longhorn_io_setting_v1beta1_manifest_test.go out/terratest-sentinel-longhorn_io_share_manager_v1beta1_manifest_test.go out/terratest-sentinel-longhorn_io_volume_v1beta1_manifest_test.go out/terratest-sentinel-longhorn_io_backing_image_data_source_v1beta2_manifest_test.go out/terratest-sentinel-longhorn_io_backing_image_manager_v1beta2_manifest_test.go out/terratest-sentinel-longhorn_io_backing_image_v1beta2_manifest_test.go out/terratest-sentinel-longhorn_io_backup_backing_image_v1beta2_manifest_test.go out/terratest-sentinel-longhorn_io_backup_target_v1beta2_manifest_test.go out/terratest-sentinel-longhorn_io_backup_v1beta2_manifest_test.go out/terratest-sentinel-longhorn_io_backup_volume_v1beta2_manifest_test.go out/terratest-sentinel-longhorn_io_engine_image_v1beta2_manifest_test.go out/terratest-sentinel-longhorn_io_engine_v1beta2_manifest_test.go out/terratest-sentinel-longhorn_io_instance_manager_v1beta2_manifest_test.go out/terratest-sentinel-longhorn_io_node_v1beta2_manifest_test.go out/terratest-sentinel-longhorn_io_orphan_v1beta2_manifest_test.go out/terratest-sentinel-longhorn_io_recurring_job_v1beta2_manifest_test.go out/terratest-sentinel-longhorn_io_replica_v1beta2_manifest_test.go out/terratest-sentinel-longhorn_io_setting_v1beta2_manifest_test.go out/terratest-sentinel-longhorn_io_share_manager_v1beta2_manifest_test.go out/terratest-sentinel-longhorn_io_snapshot_v1beta2_manifest_test.go out/terratest-sentinel-longhorn_io_support_bundle_v1beta2_manifest_test.go out/terratest-sentinel-longhorn_io_system_backup_v1beta2_manifest_test.go out/terratest-sentinel-longhorn_io_system_restore_v1beta2_manifest_test.go out/terratest-sentinel-longhorn_io_volume_attachment_v1beta2_manifest_test.go out/terratest-sentinel-longhorn_io_volume_v1beta2_manifest_test.go out/terratest-sentinel-m4e_krestomat_io_moodle_v1alpha1_manifest_test.go out/terratest-sentinel-m4e_krestomat_io_nginx_v1alpha1_manifest_test.go out/terratest-sentinel-m4e_krestomat_io_phpfpm_v1alpha1_manifest_test.go out/terratest-sentinel-m4e_krestomat_io_routine_v1alpha1_manifest_test.go out/terratest-sentinel-machine_deletion_remediation_medik8s_io_machine_deletion_remediation_template_v1alpha1_manifest_test.go out/terratest-sentinel-machine_deletion_remediation_medik8s_io_machine_deletion_remediation_v1alpha1_manifest_test.go out/terratest-sentinel-maps_k8s_elastic_co_elastic_maps_server_v1alpha1_manifest_test.go out/terratest-sentinel-mariadb_mmontes_io_backup_v1alpha1_manifest_test.go out/terratest-sentinel-mariadb_mmontes_io_connection_v1alpha1_manifest_test.go out/terratest-sentinel-mariadb_mmontes_io_database_v1alpha1_manifest_test.go out/terratest-sentinel-mariadb_mmontes_io_grant_v1alpha1_manifest_test.go out/terratest-sentinel-mariadb_mmontes_io_maria_db_v1alpha1_manifest_test.go out/terratest-sentinel-mariadb_mmontes_io_restore_v1alpha1_manifest_test.go out/terratest-sentinel-mariadb_mmontes_io_sql_job_v1alpha1_manifest_test.go out/terratest-sentinel-mariadb_mmontes_io_user_v1alpha1_manifest_test.go out/terratest-sentinel-marin3r_3scale_net_envoy_config_revision_v1alpha1_manifest_test.go out/terratest-sentinel-marin3r_3scale_net_envoy_config_v1alpha1_manifest_test.go out/terratest-sentinel-mattermost_com_cluster_installation_v1alpha1_manifest_test.go out/terratest-sentinel-mattermost_com_mattermost_restore_db_v1alpha1_manifest_test.go out/terratest-sentinel-memorydb_services_k8s_aws_acl_v1alpha1_manifest_test.go out/terratest-sentinel-memorydb_services_k8s_aws_cluster_v1alpha1_manifest_test.go out/terratest-sentinel-memorydb_services_k8s_aws_parameter_group_v1alpha1_manifest_test.go out/terratest-sentinel-memorydb_services_k8s_aws_snapshot_v1alpha1_manifest_test.go out/terratest-sentinel-memorydb_services_k8s_aws_subnet_group_v1alpha1_manifest_test.go out/terratest-sentinel-memorydb_services_k8s_aws_user_v1alpha1_manifest_test.go out/terratest-sentinel-metacontroller_k8s_io_composite_controller_v1alpha1_manifest_test.go out/terratest-sentinel-metacontroller_k8s_io_controller_revision_v1alpha1_manifest_test.go out/terratest-sentinel-metacontroller_k8s_io_decorator_controller_v1alpha1_manifest_test.go out/terratest-sentinel-metal3_io_bare_metal_host_v1alpha1_manifest_test.go out/terratest-sentinel-metal3_io_bmc_event_subscription_v1alpha1_manifest_test.go out/terratest-sentinel-metal3_io_data_image_v1alpha1_manifest_test.go out/terratest-sentinel-metal3_io_firmware_schema_v1alpha1_manifest_test.go out/terratest-sentinel-metal3_io_hardware_data_v1alpha1_manifest_test.go out/terratest-sentinel-metal3_io_host_firmware_components_v1alpha1_manifest_test.go out/terratest-sentinel-metal3_io_host_firmware_settings_v1alpha1_manifest_test.go out/terratest-sentinel-metal3_io_preprovisioning_image_v1alpha1_manifest_test.go out/terratest-sentinel-minio_min_io_tenant_v2_manifest_test.go out/terratest-sentinel-mirrors_kts_studio_secret_mirror_v1alpha1_manifest_test.go out/terratest-sentinel-mirrors_kts_studio_secret_mirror_v1alpha2_manifest_test.go out/terratest-sentinel-model_kubedl_io_model_v1alpha1_manifest_test.go out/terratest-sentinel-model_kubedl_io_model_version_v1alpha1_manifest_test.go out/terratest-sentinel-monitoring_coreos_com_alertmanager_v1_manifest_test.go out/terratest-sentinel-monitoring_coreos_com_pod_monitor_v1_manifest_test.go out/terratest-sentinel-monitoring_coreos_com_probe_v1_manifest_test.go out/terratest-sentinel-monitoring_coreos_com_prometheus_rule_v1_manifest_test.go out/terratest-sentinel-monitoring_coreos_com_prometheus_v1_manifest_test.go out/terratest-sentinel-monitoring_coreos_com_service_monitor_v1_manifest_test.go out/terratest-sentinel-monitoring_coreos_com_thanos_ruler_v1_manifest_test.go out/terratest-sentinel-monitoring_coreos_com_alertmanager_config_v1alpha1_manifest_test.go out/terratest-sentinel-monitoring_coreos_com_prometheus_agent_v1alpha1_manifest_test.go out/terratest-sentinel-monitoring_coreos_com_scrape_config_v1alpha1_manifest_test.go out/terratest-sentinel-monitoring_coreos_com_alertmanager_config_v1beta1_manifest_test.go out/terratest-sentinel-monocle_monocle_change_metrics_io_monocle_v1alpha1_manifest_test.go out/terratest-sentinel-mq_services_k8s_aws_broker_v1alpha1_manifest_test.go out/terratest-sentinel-multicluster_crd_antrea_io_cluster_info_import_v1alpha1_manifest_test.go out/terratest-sentinel-multicluster_crd_antrea_io_cluster_set_v1alpha1_manifest_test.go out/terratest-sentinel-multicluster_crd_antrea_io_gateway_v1alpha1_manifest_test.go out/terratest-sentinel-multicluster_crd_antrea_io_label_identity_v1alpha1_manifest_test.go out/terratest-sentinel-multicluster_crd_antrea_io_member_cluster_announce_v1alpha1_manifest_test.go out/terratest-sentinel-multicluster_crd_antrea_io_multi_cluster_config_v1alpha1_manifest_test.go out/terratest-sentinel-multicluster_crd_antrea_io_resource_export_v1alpha1_manifest_test.go out/terratest-sentinel-multicluster_crd_antrea_io_resource_import_v1alpha1_manifest_test.go out/terratest-sentinel-multicluster_crd_antrea_io_cluster_claim_v1alpha2_manifest_test.go out/terratest-sentinel-multicluster_crd_antrea_io_cluster_set_v1alpha2_manifest_test.go out/terratest-sentinel-multicluster_x_k8s_io_applied_work_v1alpha1_manifest_test.go out/terratest-sentinel-multicluster_x_k8s_io_service_import_v1alpha1_manifest_test.go out/terratest-sentinel-multicluster_x_k8s_io_work_v1alpha1_manifest_test.go out/terratest-sentinel-mutations_gatekeeper_sh_assign_metadata_v1_manifest_test.go out/terratest-sentinel-mutations_gatekeeper_sh_assign_v1_manifest_test.go out/terratest-sentinel-mutations_gatekeeper_sh_modify_set_v1_manifest_test.go out/terratest-sentinel-mutations_gatekeeper_sh_assign_image_v1alpha1_manifest_test.go out/terratest-sentinel-mutations_gatekeeper_sh_assign_metadata_v1alpha1_manifest_test.go out/terratest-sentinel-mutations_gatekeeper_sh_assign_v1alpha1_manifest_test.go out/terratest-sentinel-mutations_gatekeeper_sh_modify_set_v1alpha1_manifest_test.go out/terratest-sentinel-mutations_gatekeeper_sh_assign_metadata_v1beta1_manifest_test.go out/terratest-sentinel-mutations_gatekeeper_sh_assign_v1beta1_manifest_test.go out/terratest-sentinel-mutations_gatekeeper_sh_modify_set_v1beta1_manifest_test.go out/terratest-sentinel-nativestor_alauda_io_raw_device_v1_manifest_test.go out/terratest-sentinel-netchecks_io_network_assertion_v1_manifest_test.go out/terratest-sentinel-networkfirewall_services_k8s_aws_firewall_policy_v1alpha1_manifest_test.go out/terratest-sentinel-networkfirewall_services_k8s_aws_firewall_v1alpha1_manifest_test.go out/terratest-sentinel-networkfirewall_services_k8s_aws_rule_group_v1alpha1_manifest_test.go out/terratest-sentinel-networking_gke_io_gcp_backend_policy_v1_manifest_test.go out/terratest-sentinel-networking_gke_io_gcp_gateway_policy_v1_manifest_test.go out/terratest-sentinel-networking_gke_io_health_check_policy_v1_manifest_test.go out/terratest-sentinel-networking_gke_io_lb_policy_v1_manifest_test.go out/terratest-sentinel-networking_gke_io_managed_certificate_v1_manifest_test.go out/terratest-sentinel-networking_istio_io_destination_rule_v1_manifest_test.go out/terratest-sentinel-networking_istio_io_gateway_v1_manifest_test.go out/terratest-sentinel-networking_istio_io_service_entry_v1_manifest_test.go out/terratest-sentinel-networking_istio_io_sidecar_v1_manifest_test.go out/terratest-sentinel-networking_istio_io_virtual_service_v1_manifest_test.go out/terratest-sentinel-networking_istio_io_workload_entry_v1_manifest_test.go out/terratest-sentinel-networking_istio_io_workload_group_v1_manifest_test.go out/terratest-sentinel-networking_istio_io_destination_rule_v1alpha3_manifest_test.go out/terratest-sentinel-networking_istio_io_envoy_filter_v1alpha3_manifest_test.go out/terratest-sentinel-networking_istio_io_gateway_v1alpha3_manifest_test.go out/terratest-sentinel-networking_istio_io_service_entry_v1alpha3_manifest_test.go out/terratest-sentinel-networking_istio_io_sidecar_v1alpha3_manifest_test.go out/terratest-sentinel-networking_istio_io_virtual_service_v1alpha3_manifest_test.go out/terratest-sentinel-networking_istio_io_workload_entry_v1alpha3_manifest_test.go out/terratest-sentinel-networking_istio_io_workload_group_v1alpha3_manifest_test.go out/terratest-sentinel-networking_istio_io_destination_rule_v1beta1_manifest_test.go out/terratest-sentinel-networking_istio_io_gateway_v1beta1_manifest_test.go out/terratest-sentinel-networking_istio_io_proxy_config_v1beta1_manifest_test.go out/terratest-sentinel-networking_istio_io_service_entry_v1beta1_manifest_test.go out/terratest-sentinel-networking_istio_io_sidecar_v1beta1_manifest_test.go out/terratest-sentinel-networking_istio_io_virtual_service_v1beta1_manifest_test.go out/terratest-sentinel-networking_istio_io_workload_entry_v1beta1_manifest_test.go out/terratest-sentinel-networking_istio_io_workload_group_v1beta1_manifest_test.go out/terratest-sentinel-networking_k8s_aws_policy_endpoint_v1alpha1_manifest_test.go out/terratest-sentinel-networking_k8s_io_ingress_class_v1_manifest_test.go out/terratest-sentinel-networking_k8s_io_ingress_v1_manifest_test.go out/terratest-sentinel-networking_k8s_io_network_policy_v1_manifest_test.go out/terratest-sentinel-networking_karmada_io_multi_cluster_ingress_v1alpha1_manifest_test.go out/terratest-sentinel-networking_karmada_io_multi_cluster_service_v1alpha1_manifest_test.go out/terratest-sentinel-nfd_k8s_sigs_io_node_feature_rule_v1alpha1_manifest_test.go out/terratest-sentinel-nfd_kubernetes_io_node_feature_discovery_v1_manifest_test.go out/terratest-sentinel-nfd_kubernetes_io_node_feature_rule_v1alpha1_manifest_test.go out/terratest-sentinel-nodeinfo_volcano_sh_numatopology_v1alpha1_manifest_test.go out/terratest-sentinel-notebook_kubedl_io_notebook_v1alpha1_manifest_test.go out/terratest-sentinel-notification_toolkit_fluxcd_io_receiver_v1_manifest_test.go out/terratest-sentinel-notification_toolkit_fluxcd_io_alert_v1beta1_manifest_test.go out/terratest-sentinel-notification_toolkit_fluxcd_io_provider_v1beta1_manifest_test.go out/terratest-sentinel-notification_toolkit_fluxcd_io_receiver_v1beta1_manifest_test.go out/terratest-sentinel-notification_toolkit_fluxcd_io_alert_v1beta2_manifest_test.go out/terratest-sentinel-notification_toolkit_fluxcd_io_provider_v1beta2_manifest_test.go out/terratest-sentinel-notification_toolkit_fluxcd_io_receiver_v1beta2_manifest_test.go out/terratest-sentinel-notification_toolkit_fluxcd_io_alert_v1beta3_manifest_test.go out/terratest-sentinel-notification_toolkit_fluxcd_io_provider_v1beta3_manifest_test.go out/terratest-sentinel-objectbucket_io_object_bucket_claim_v1alpha1_manifest_test.go out/terratest-sentinel-objectbucket_io_object_bucket_v1alpha1_manifest_test.go out/terratest-sentinel-onepassword_com_one_password_item_v1_manifest_test.go out/terratest-sentinel-opensearchservice_services_k8s_aws_domain_v1alpha1_manifest_test.go out/terratest-sentinel-opentelemetry_io_instrumentation_v1alpha1_manifest_test.go out/terratest-sentinel-opentelemetry_io_op_amp_bridge_v1alpha1_manifest_test.go out/terratest-sentinel-opentelemetry_io_open_telemetry_collector_v1alpha1_manifest_test.go out/terratest-sentinel-opentelemetry_io_open_telemetry_collector_v1beta1_manifest_test.go out/terratest-sentinel-operations_kubeedge_io_node_upgrade_job_v1alpha1_manifest_test.go out/terratest-sentinel-operator_aquasec_com_aqua_csp_v1alpha1_manifest_test.go out/terratest-sentinel-operator_aquasec_com_aqua_database_v1alpha1_manifest_test.go out/terratest-sentinel-operator_aquasec_com_aqua_enforcer_v1alpha1_manifest_test.go out/terratest-sentinel-operator_aquasec_com_aqua_gateway_v1alpha1_manifest_test.go out/terratest-sentinel-operator_aquasec_com_aqua_kube_enforcer_v1alpha1_manifest_test.go out/terratest-sentinel-operator_aquasec_com_aqua_scanner_v1alpha1_manifest_test.go out/terratest-sentinel-operator_aquasec_com_aqua_server_v1alpha1_manifest_test.go out/terratest-sentinel-operator_authorino_kuadrant_io_authorino_v1beta1_manifest_test.go out/terratest-sentinel-operator_cluster_x_k8s_io_bootstrap_provider_v1alpha1_manifest_test.go out/terratest-sentinel-operator_cluster_x_k8s_io_control_plane_provider_v1alpha1_manifest_test.go out/terratest-sentinel-operator_cluster_x_k8s_io_core_provider_v1alpha1_manifest_test.go out/terratest-sentinel-operator_cluster_x_k8s_io_infrastructure_provider_v1alpha1_manifest_test.go out/terratest-sentinel-operator_cluster_x_k8s_io_addon_provider_v1alpha2_manifest_test.go out/terratest-sentinel-operator_cluster_x_k8s_io_bootstrap_provider_v1alpha2_manifest_test.go out/terratest-sentinel-operator_cluster_x_k8s_io_control_plane_provider_v1alpha2_manifest_test.go out/terratest-sentinel-operator_cluster_x_k8s_io_core_provider_v1alpha2_manifest_test.go out/terratest-sentinel-operator_cluster_x_k8s_io_infrastructure_provider_v1alpha2_manifest_test.go out/terratest-sentinel-operator_cryostat_io_cryostat_v1beta1_manifest_test.go out/terratest-sentinel-operator_cryostat_io_cryostat_v1beta2_manifest_test.go out/terratest-sentinel-operator_knative_dev_knative_eventing_v1beta1_manifest_test.go out/terratest-sentinel-operator_knative_dev_knative_serving_v1beta1_manifest_test.go out/terratest-sentinel-operator_marin3r_3scale_net_discovery_service_certificate_v1alpha1_manifest_test.go out/terratest-sentinel-operator_marin3r_3scale_net_discovery_service_v1alpha1_manifest_test.go out/terratest-sentinel-operator_marin3r_3scale_net_envoy_deployment_v1alpha1_manifest_test.go out/terratest-sentinel-operator_open_cluster_management_io_cluster_manager_v1_manifest_test.go out/terratest-sentinel-operator_open_cluster_management_io_klusterlet_v1_manifest_test.go out/terratest-sentinel-operator_shipwright_io_shipwright_build_v1alpha1_manifest_test.go out/terratest-sentinel-operator_tigera_io_amazon_cloud_integration_v1_manifest_test.go out/terratest-sentinel-operator_tigera_io_api_server_v1_manifest_test.go out/terratest-sentinel-operator_tigera_io_application_layer_v1_manifest_test.go out/terratest-sentinel-operator_tigera_io_authentication_v1_manifest_test.go out/terratest-sentinel-operator_tigera_io_compliance_v1_manifest_test.go out/terratest-sentinel-operator_tigera_io_egress_gateway_v1_manifest_test.go out/terratest-sentinel-operator_tigera_io_image_set_v1_manifest_test.go out/terratest-sentinel-operator_tigera_io_installation_v1_manifest_test.go out/terratest-sentinel-operator_tigera_io_intrusion_detection_v1_manifest_test.go out/terratest-sentinel-operator_tigera_io_log_collector_v1_manifest_test.go out/terratest-sentinel-operator_tigera_io_log_storage_v1_manifest_test.go out/terratest-sentinel-operator_tigera_io_management_cluster_connection_v1_manifest_test.go out/terratest-sentinel-operator_tigera_io_management_cluster_v1_manifest_test.go out/terratest-sentinel-operator_tigera_io_manager_v1_manifest_test.go out/terratest-sentinel-operator_tigera_io_monitor_v1_manifest_test.go out/terratest-sentinel-operator_tigera_io_packet_capture_v1_manifest_test.go out/terratest-sentinel-operator_tigera_io_policy_recommendation_v1_manifest_test.go out/terratest-sentinel-operator_tigera_io_tenant_v1_manifest_test.go out/terratest-sentinel-operator_tigera_io_tigera_status_v1_manifest_test.go out/terratest-sentinel-operator_tigera_io_tls_pass_through_route_v1_manifest_test.go out/terratest-sentinel-operator_tigera_io_tls_terminated_route_v1_manifest_test.go out/terratest-sentinel-operator_tigera_io_amazon_cloud_integration_v1beta1_manifest_test.go out/terratest-sentinel-operator_victoriametrics_com_vm_agent_v1beta1_manifest_test.go out/terratest-sentinel-operator_victoriametrics_com_vm_alert_v1beta1_manifest_test.go out/terratest-sentinel-operator_victoriametrics_com_vm_alertmanager_config_v1beta1_manifest_test.go out/terratest-sentinel-operator_victoriametrics_com_vm_alertmanager_v1beta1_manifest_test.go out/terratest-sentinel-operator_victoriametrics_com_vm_auth_v1beta1_manifest_test.go out/terratest-sentinel-operator_victoriametrics_com_vm_cluster_v1beta1_manifest_test.go out/terratest-sentinel-operator_victoriametrics_com_vm_node_scrape_v1beta1_manifest_test.go out/terratest-sentinel-operator_victoriametrics_com_vm_pod_scrape_v1beta1_manifest_test.go out/terratest-sentinel-operator_victoriametrics_com_vm_probe_v1beta1_manifest_test.go out/terratest-sentinel-operator_victoriametrics_com_vm_rule_v1beta1_manifest_test.go out/terratest-sentinel-operator_victoriametrics_com_vm_service_scrape_v1beta1_manifest_test.go out/terratest-sentinel-operator_victoriametrics_com_vm_single_v1beta1_manifest_test.go out/terratest-sentinel-operator_victoriametrics_com_vm_static_scrape_v1beta1_manifest_test.go out/terratest-sentinel-operator_victoriametrics_com_vm_user_v1beta1_manifest_test.go out/terratest-sentinel-oracle_db_anthosapis_com_backup_schedule_v1alpha1_manifest_test.go out/terratest-sentinel-oracle_db_anthosapis_com_backup_v1alpha1_manifest_test.go out/terratest-sentinel-oracle_db_anthosapis_com_config_v1alpha1_manifest_test.go out/terratest-sentinel-oracle_db_anthosapis_com_cron_anything_v1alpha1_manifest_test.go out/terratest-sentinel-oracle_db_anthosapis_com_database_v1alpha1_manifest_test.go out/terratest-sentinel-oracle_db_anthosapis_com_export_v1alpha1_manifest_test.go out/terratest-sentinel-oracle_db_anthosapis_com_import_v1alpha1_manifest_test.go out/terratest-sentinel-oracle_db_anthosapis_com_instance_v1alpha1_manifest_test.go out/terratest-sentinel-oracle_db_anthosapis_com_pitr_v1alpha1_manifest_test.go out/terratest-sentinel-oracle_db_anthosapis_com_release_v1alpha1_manifest_test.go out/terratest-sentinel-org_eclipse_che_che_cluster_v1_manifest_test.go out/terratest-sentinel-org_eclipse_che_che_cluster_v2_manifest_test.go out/terratest-sentinel-organizations_services_k8s_aws_organizational_unit_v1alpha1_manifest_test.go out/terratest-sentinel-pgv2_percona_com_percona_pg_backup_v2_manifest_test.go out/terratest-sentinel-pgv2_percona_com_percona_pg_cluster_v2_manifest_test.go out/terratest-sentinel-pgv2_percona_com_percona_pg_restore_v2_manifest_test.go out/terratest-sentinel-pgv2_percona_com_percona_pg_upgrade_v2_manifest_test.go out/terratest-sentinel-pipes_services_k8s_aws_pipe_v1alpha1_manifest_test.go out/terratest-sentinel-pkg_crossplane_io_configuration_revision_v1_manifest_test.go out/terratest-sentinel-pkg_crossplane_io_configuration_v1_manifest_test.go out/terratest-sentinel-pkg_crossplane_io_provider_revision_v1_manifest_test.go out/terratest-sentinel-pkg_crossplane_io_provider_v1_manifest_test.go out/terratest-sentinel-pkg_crossplane_io_controller_config_v1alpha1_manifest_test.go out/terratest-sentinel-pkg_crossplane_io_lock_v1beta1_manifest_test.go out/terratest-sentinel-policy_clusterpedia_io_cluster_import_policy_v1alpha1_manifest_test.go out/terratest-sentinel-policy_clusterpedia_io_pedia_cluster_lifecycle_v1alpha1_manifest_test.go out/terratest-sentinel-policy_karmada_io_cluster_override_policy_v1alpha1_manifest_test.go out/terratest-sentinel-policy_karmada_io_cluster_propagation_policy_v1alpha1_manifest_test.go out/terratest-sentinel-policy_karmada_io_federated_resource_quota_v1alpha1_manifest_test.go out/terratest-sentinel-policy_karmada_io_override_policy_v1alpha1_manifest_test.go out/terratest-sentinel-policy_karmada_io_propagation_policy_v1alpha1_manifest_test.go out/terratest-sentinel-policy_kubeedge_io_service_account_access_v1alpha1_manifest_test.go out/terratest-sentinel-policy_networking_k8s_io_admin_network_policy_v1alpha1_manifest_test.go out/terratest-sentinel-policy_networking_k8s_io_baseline_admin_network_policy_v1alpha1_manifest_test.go out/terratest-sentinel-policy_pod_disruption_budget_v1_manifest_test.go out/terratest-sentinel-postgres_operator_crunchydata_com_pg_admin_v1beta1_manifest_test.go out/terratest-sentinel-postgres_operator_crunchydata_com_pg_upgrade_v1beta1_manifest_test.go out/terratest-sentinel-postgres_operator_crunchydata_com_postgres_cluster_v1beta1_manifest_test.go out/terratest-sentinel-postgresql_cnpg_io_backup_v1_manifest_test.go out/terratest-sentinel-postgresql_cnpg_io_cluster_v1_manifest_test.go out/terratest-sentinel-postgresql_cnpg_io_pooler_v1_manifest_test.go out/terratest-sentinel-postgresql_cnpg_io_scheduled_backup_v1_manifest_test.go out/terratest-sentinel-projectcontour_io_http_proxy_v1_manifest_test.go out/terratest-sentinel-projectcontour_io_tls_certificate_delegation_v1_manifest_test.go out/terratest-sentinel-projectcontour_io_contour_configuration_v1alpha1_manifest_test.go out/terratest-sentinel-projectcontour_io_contour_deployment_v1alpha1_manifest_test.go out/terratest-sentinel-projectcontour_io_extension_service_v1alpha1_manifest_test.go out/terratest-sentinel-prometheusservice_services_k8s_aws_alert_manager_definition_v1alpha1_manifest_test.go out/terratest-sentinel-prometheusservice_services_k8s_aws_logging_configuration_v1alpha1_manifest_test.go out/terratest-sentinel-prometheusservice_services_k8s_aws_rule_groups_namespace_v1alpha1_manifest_test.go out/terratest-sentinel-prometheusservice_services_k8s_aws_workspace_v1alpha1_manifest_test.go out/terratest-sentinel-ps_percona_com_percona_server_my_sql_v1alpha1_manifest_test.go out/terratest-sentinel-ps_percona_com_percona_server_my_sql_backup_v1alpha1_manifest_test.go out/terratest-sentinel-ps_percona_com_percona_server_my_sql_restore_v1alpha1_manifest_test.go out/terratest-sentinel-psmdb_percona_com_percona_server_mongo_db_v1_manifest_test.go out/terratest-sentinel-psmdb_percona_com_percona_server_mongo_db_backup_v1_manifest_test.go out/terratest-sentinel-psmdb_percona_com_percona_server_mongo_db_restore_v1_manifest_test.go out/terratest-sentinel-ptp_openshift_io_node_ptp_device_v1_manifest_test.go out/terratest-sentinel-ptp_openshift_io_ptp_config_v1_manifest_test.go out/terratest-sentinel-ptp_openshift_io_ptp_operator_config_v1_manifest_test.go out/terratest-sentinel-pubsubplus_solace_com_pub_sub_plus_event_broker_v1beta1_manifest_test.go out/terratest-sentinel-pxc_percona_com_percona_xtra_db_cluster_backup_v1_manifest_test.go out/terratest-sentinel-pxc_percona_com_percona_xtra_db_cluster_restore_v1_manifest_test.go out/terratest-sentinel-pxc_percona_com_percona_xtra_db_cluster_v1_manifest_test.go out/terratest-sentinel-quay_redhat_com_quay_registry_v1_manifest_test.go out/terratest-sentinel-quota_codeflare_dev_quota_subtree_v1alpha1_manifest_test.go out/terratest-sentinel-ray_io_ray_cluster_v1_manifest_test.go out/terratest-sentinel-ray_io_ray_job_v1_manifest_test.go out/terratest-sentinel-ray_io_ray_service_v1_manifest_test.go out/terratest-sentinel-ray_io_ray_cluster_v1alpha1_manifest_test.go out/terratest-sentinel-ray_io_ray_job_v1alpha1_manifest_test.go out/terratest-sentinel-ray_io_ray_service_v1alpha1_manifest_test.go out/terratest-sentinel-rbac_authorization_k8s_io_cluster_role_binding_v1_manifest_test.go out/terratest-sentinel-rbac_authorization_k8s_io_cluster_role_v1_manifest_test.go out/terratest-sentinel-rbac_authorization_k8s_io_role_binding_v1_manifest_test.go out/terratest-sentinel-rbac_authorization_k8s_io_role_v1_manifest_test.go out/terratest-sentinel-rbacmanager_reactiveops_io_rbac_definition_v1beta1_manifest_test.go out/terratest-sentinel-rc_app_stacks_runtime_component_v1_manifest_test.go out/terratest-sentinel-rc_app_stacks_runtime_operation_v1_manifest_test.go out/terratest-sentinel-rc_app_stacks_runtime_component_v1beta2_manifest_test.go out/terratest-sentinel-rc_app_stacks_runtime_operation_v1beta2_manifest_test.go out/terratest-sentinel-rds_services_k8s_aws_db_cluster_parameter_group_v1alpha1_manifest_test.go out/terratest-sentinel-rds_services_k8s_aws_db_cluster_v1alpha1_manifest_test.go out/terratest-sentinel-rds_services_k8s_aws_db_instance_v1alpha1_manifest_test.go out/terratest-sentinel-rds_services_k8s_aws_db_parameter_group_v1alpha1_manifest_test.go out/terratest-sentinel-rds_services_k8s_aws_db_proxy_v1alpha1_manifest_test.go out/terratest-sentinel-rds_services_k8s_aws_db_subnet_group_v1alpha1_manifest_test.go out/terratest-sentinel-rds_services_k8s_aws_global_cluster_v1alpha1_manifest_test.go out/terratest-sentinel-redhatcop_redhat_io_group_config_v1alpha1_manifest_test.go out/terratest-sentinel-redhatcop_redhat_io_keepalived_group_v1alpha1_manifest_test.go out/terratest-sentinel-redhatcop_redhat_io_namespace_config_v1alpha1_manifest_test.go out/terratest-sentinel-redhatcop_redhat_io_patch_v1alpha1_manifest_test.go out/terratest-sentinel-redhatcop_redhat_io_user_config_v1alpha1_manifest_test.go out/terratest-sentinel-registry_apicur_io_apicurio_registry_v1_manifest_test.go out/terratest-sentinel-registry_devfile_io_cluster_devfile_registries_list_v1alpha1_manifest_test.go out/terratest-sentinel-registry_devfile_io_devfile_registries_list_v1alpha1_manifest_test.go out/terratest-sentinel-registry_devfile_io_devfile_registry_v1alpha1_manifest_test.go out/terratest-sentinel-reliablesyncs_kubeedge_io_cluster_object_sync_v1alpha1_manifest_test.go out/terratest-sentinel-reliablesyncs_kubeedge_io_object_sync_v1alpha1_manifest_test.go out/terratest-sentinel-remediation_medik8s_io_node_health_check_v1alpha1_manifest_test.go out/terratest-sentinel-repo_manager_pulpproject_org_pulp_backup_v1beta2_manifest_test.go out/terratest-sentinel-repo_manager_pulpproject_org_pulp_restore_v1beta2_manifest_test.go out/terratest-sentinel-repo_manager_pulpproject_org_pulp_v1beta2_manifest_test.go out/terratest-sentinel-reports_kyverno_io_cluster_ephemeral_report_v1_manifest_test.go out/terratest-sentinel-reports_kyverno_io_ephemeral_report_v1_manifest_test.go out/terratest-sentinel-resources_teleport_dev_teleport_login_rule_v1_manifest_test.go out/terratest-sentinel-resources_teleport_dev_teleport_okta_import_rule_v1_manifest_test.go out/terratest-sentinel-resources_teleport_dev_teleport_provision_token_v2_manifest_test.go out/terratest-sentinel-resources_teleport_dev_teleport_saml_connector_v2_manifest_test.go out/terratest-sentinel-resources_teleport_dev_teleport_user_v2_manifest_test.go out/terratest-sentinel-resources_teleport_dev_teleport_github_connector_v3_manifest_test.go out/terratest-sentinel-resources_teleport_dev_teleport_oidc_connector_v3_manifest_test.go out/terratest-sentinel-resources_teleport_dev_teleport_role_v5_manifest_test.go out/terratest-sentinel-resources_teleport_dev_teleport_role_v6_manifest_test.go out/terratest-sentinel-ripsaw_cloudbulldozer_io_benchmark_v1alpha1_manifest_test.go out/terratest-sentinel-rocketmq_apache_org_broker_v1alpha1_manifest_test.go out/terratest-sentinel-rocketmq_apache_org_console_v1alpha1_manifest_test.go out/terratest-sentinel-rocketmq_apache_org_name_service_v1alpha1_manifest_test.go out/terratest-sentinel-rocketmq_apache_org_topic_transfer_v1alpha1_manifest_test.go out/terratest-sentinel-route53_services_k8s_aws_hosted_zone_v1alpha1_manifest_test.go out/terratest-sentinel-route53_services_k8s_aws_record_set_v1alpha1_manifest_test.go out/terratest-sentinel-route53resolver_services_k8s_aws_resolver_endpoint_v1alpha1_manifest_test.go out/terratest-sentinel-route53resolver_services_k8s_aws_resolver_rule_v1alpha1_manifest_test.go out/terratest-sentinel-rules_kubeedge_io_rule_endpoint_v1_manifest_test.go out/terratest-sentinel-rules_kubeedge_io_rule_v1_manifest_test.go out/terratest-sentinel-runtime_cluster_x_k8s_io_extension_config_v1alpha1_manifest_test.go out/terratest-sentinel-s3_services_k8s_aws_bucket_v1alpha1_manifest_test.go out/terratest-sentinel-s3_snappcloud_io_s3_bucket_v1alpha1_manifest_test.go out/terratest-sentinel-s3_snappcloud_io_s3_user_claim_v1alpha1_manifest_test.go out/terratest-sentinel-s3_snappcloud_io_s3_user_v1alpha1_manifest_test.go out/terratest-sentinel-sagemaker_services_k8s_aws_app_v1alpha1_manifest_test.go out/terratest-sentinel-sagemaker_services_k8s_aws_data_quality_job_definition_v1alpha1_manifest_test.go out/terratest-sentinel-sagemaker_services_k8s_aws_domain_v1alpha1_manifest_test.go out/terratest-sentinel-sagemaker_services_k8s_aws_endpoint_config_v1alpha1_manifest_test.go out/terratest-sentinel-sagemaker_services_k8s_aws_endpoint_v1alpha1_manifest_test.go out/terratest-sentinel-sagemaker_services_k8s_aws_feature_group_v1alpha1_manifest_test.go out/terratest-sentinel-sagemaker_services_k8s_aws_hyper_parameter_tuning_job_v1alpha1_manifest_test.go out/terratest-sentinel-sagemaker_services_k8s_aws_model_bias_job_definition_v1alpha1_manifest_test.go out/terratest-sentinel-sagemaker_services_k8s_aws_model_explainability_job_definition_v1alpha1_manifest_test.go out/terratest-sentinel-sagemaker_services_k8s_aws_model_package_group_v1alpha1_manifest_test.go out/terratest-sentinel-sagemaker_services_k8s_aws_model_package_v1alpha1_manifest_test.go out/terratest-sentinel-sagemaker_services_k8s_aws_model_quality_job_definition_v1alpha1_manifest_test.go out/terratest-sentinel-sagemaker_services_k8s_aws_model_v1alpha1_manifest_test.go out/terratest-sentinel-sagemaker_services_k8s_aws_monitoring_schedule_v1alpha1_manifest_test.go out/terratest-sentinel-sagemaker_services_k8s_aws_notebook_instance_lifecycle_config_v1alpha1_manifest_test.go out/terratest-sentinel-sagemaker_services_k8s_aws_notebook_instance_v1alpha1_manifest_test.go out/terratest-sentinel-sagemaker_services_k8s_aws_processing_job_v1alpha1_manifest_test.go out/terratest-sentinel-sagemaker_services_k8s_aws_training_job_v1alpha1_manifest_test.go out/terratest-sentinel-sagemaker_services_k8s_aws_transform_job_v1alpha1_manifest_test.go out/terratest-sentinel-sagemaker_services_k8s_aws_user_profile_v1alpha1_manifest_test.go out/terratest-sentinel-scheduling_k8s_io_priority_class_v1_manifest_test.go out/terratest-sentinel-scheduling_koordinator_sh_device_v1alpha1_manifest_test.go out/terratest-sentinel-scheduling_koordinator_sh_pod_migration_job_v1alpha1_manifest_test.go out/terratest-sentinel-scheduling_koordinator_sh_reservation_v1alpha1_manifest_test.go out/terratest-sentinel-scheduling_sigs_k8s_io_elastic_quota_v1alpha1_manifest_test.go out/terratest-sentinel-scheduling_sigs_k8s_io_pod_group_v1alpha1_manifest_test.go out/terratest-sentinel-scheduling_volcano_sh_pod_group_v1beta1_manifest_test.go out/terratest-sentinel-scheduling_volcano_sh_queue_v1beta1_manifest_test.go out/terratest-sentinel-schemas_schemahero_io_data_type_v1alpha4_manifest_test.go out/terratest-sentinel-schemas_schemahero_io_migration_v1alpha4_manifest_test.go out/terratest-sentinel-schemas_schemahero_io_table_v1alpha4_manifest_test.go out/terratest-sentinel-scylla_scylladb_com_scylla_cluster_v1_manifest_test.go out/terratest-sentinel-scylla_scylladb_com_node_config_v1alpha1_manifest_test.go out/terratest-sentinel-scylla_scylladb_com_scylla_operator_config_v1alpha1_manifest_test.go out/terratest-sentinel-secretgenerator_mittwald_de_basic_auth_v1alpha1_manifest_test.go out/terratest-sentinel-secretgenerator_mittwald_de_ssh_key_pair_v1alpha1_manifest_test.go out/terratest-sentinel-secretgenerator_mittwald_de_string_secret_v1alpha1_manifest_test.go out/terratest-sentinel-secrets_crossplane_io_store_config_v1alpha1_manifest_test.go out/terratest-sentinel-secrets_doppler_com_doppler_secret_v1alpha1_manifest_test.go out/terratest-sentinel-secrets_hashicorp_com_hcp_auth_v1beta1_manifest_test.go out/terratest-sentinel-secrets_hashicorp_com_hcp_vault_secrets_app_v1beta1_manifest_test.go out/terratest-sentinel-secrets_hashicorp_com_vault_auth_v1beta1_manifest_test.go out/terratest-sentinel-secrets_hashicorp_com_vault_connection_v1beta1_manifest_test.go out/terratest-sentinel-secrets_hashicorp_com_vault_dynamic_secret_v1beta1_manifest_test.go out/terratest-sentinel-secrets_hashicorp_com_vault_pki_secret_v1beta1_manifest_test.go out/terratest-sentinel-secrets_hashicorp_com_vault_static_secret_v1beta1_manifest_test.go out/terratest-sentinel-secrets_store_csi_x_k8s_io_secret_provider_class_v1_manifest_test.go out/terratest-sentinel-secrets_store_csi_x_k8s_io_secret_provider_class_v1alpha1_manifest_test.go out/terratest-sentinel-secretsmanager_services_k8s_aws_secret_v1alpha1_manifest_test.go out/terratest-sentinel-secscan_quay_redhat_com_image_manifest_vuln_v1alpha1_manifest_test.go out/terratest-sentinel-security_istio_io_authorization_policy_v1_manifest_test.go out/terratest-sentinel-security_istio_io_peer_authentication_v1_manifest_test.go out/terratest-sentinel-security_istio_io_request_authentication_v1_manifest_test.go out/terratest-sentinel-security_istio_io_authorization_policy_v1beta1_manifest_test.go out/terratest-sentinel-security_istio_io_peer_authentication_v1beta1_manifest_test.go out/terratest-sentinel-security_istio_io_request_authentication_v1beta1_manifest_test.go out/terratest-sentinel-security_profiles_operator_x_k8s_io_app_armor_profile_v1alpha1_manifest_test.go out/terratest-sentinel-security_profiles_operator_x_k8s_io_profile_binding_v1alpha1_manifest_test.go out/terratest-sentinel-security_profiles_operator_x_k8s_io_profile_recording_v1alpha1_manifest_test.go out/terratest-sentinel-security_profiles_operator_x_k8s_io_security_profile_node_status_v1alpha1_manifest_test.go out/terratest-sentinel-security_profiles_operator_x_k8s_io_security_profiles_operator_daemon_v1alpha1_manifest_test.go out/terratest-sentinel-security_profiles_operator_x_k8s_io_raw_selinux_profile_v1alpha2_manifest_test.go out/terratest-sentinel-security_profiles_operator_x_k8s_io_selinux_profile_v1alpha2_manifest_test.go out/terratest-sentinel-security_profiles_operator_x_k8s_io_seccomp_profile_v1beta1_manifest_test.go out/terratest-sentinel-self_node_remediation_medik8s_io_self_node_remediation_config_v1alpha1_manifest_test.go out/terratest-sentinel-self_node_remediation_medik8s_io_self_node_remediation_template_v1alpha1_manifest_test.go out/terratest-sentinel-self_node_remediation_medik8s_io_self_node_remediation_v1alpha1_manifest_test.go out/terratest-sentinel-sematext_com_sematext_agent_v1_manifest_test.go out/terratest-sentinel-servicebinding_io_cluster_workload_resource_mapping_v1alpha3_manifest_test.go out/terratest-sentinel-servicebinding_io_service_binding_v1alpha3_manifest_test.go out/terratest-sentinel-servicebinding_io_cluster_workload_resource_mapping_v1beta1_manifest_test.go out/terratest-sentinel-servicebinding_io_service_binding_v1beta1_manifest_test.go out/terratest-sentinel-servicemesh_cisco_com_istio_control_plane_v1alpha1_manifest_test.go out/terratest-sentinel-servicemesh_cisco_com_istio_mesh_gateway_v1alpha1_manifest_test.go out/terratest-sentinel-servicemesh_cisco_com_istio_mesh_v1alpha1_manifest_test.go out/terratest-sentinel-servicemesh_cisco_com_peer_istio_control_plane_v1alpha1_manifest_test.go out/terratest-sentinel-services_k8s_aws_adopted_resource_v1alpha1_manifest_test.go out/terratest-sentinel-services_k8s_aws_field_export_v1alpha1_manifest_test.go out/terratest-sentinel-serving_kubedl_io_inference_v1alpha1_manifest_test.go out/terratest-sentinel-sfn_services_k8s_aws_activity_v1alpha1_manifest_test.go out/terratest-sentinel-sfn_services_k8s_aws_state_machine_v1alpha1_manifest_test.go out/terratest-sentinel-site_superedge_io_node_group_v1alpha1_manifest_test.go out/terratest-sentinel-site_superedge_io_node_unit_v1alpha1_manifest_test.go out/terratest-sentinel-slo_koordinator_sh_node_metric_v1alpha1_manifest_test.go out/terratest-sentinel-slo_koordinator_sh_node_slo_v1alpha1_manifest_test.go out/terratest-sentinel-sloth_slok_dev_prometheus_service_level_v1_manifest_test.go out/terratest-sentinel-snapscheduler_backube_snapshot_schedule_v1_manifest_test.go out/terratest-sentinel-snapshot_storage_k8s_io_volume_snapshot_class_v1_manifest_test.go out/terratest-sentinel-snapshot_storage_k8s_io_volume_snapshot_content_v1_manifest_test.go out/terratest-sentinel-snapshot_storage_k8s_io_volume_snapshot_v1_manifest_test.go out/terratest-sentinel-snapshot_storage_k8s_io_volume_snapshot_class_v1beta1_manifest_test.go out/terratest-sentinel-snapshot_storage_k8s_io_volume_snapshot_content_v1beta1_manifest_test.go out/terratest-sentinel-snapshot_storage_k8s_io_volume_snapshot_v1beta1_manifest_test.go out/terratest-sentinel-sns_services_k8s_aws_platform_application_v1alpha1_manifest_test.go out/terratest-sentinel-sns_services_k8s_aws_platform_endpoint_v1alpha1_manifest_test.go out/terratest-sentinel-sns_services_k8s_aws_subscription_v1alpha1_manifest_test.go out/terratest-sentinel-sns_services_k8s_aws_topic_v1alpha1_manifest_test.go out/terratest-sentinel-sonataflow_org_sonata_flow_build_v1alpha08_manifest_test.go out/terratest-sentinel-sonataflow_org_sonata_flow_platform_v1alpha08_manifest_test.go out/terratest-sentinel-sonataflow_org_sonata_flow_v1alpha08_manifest_test.go out/terratest-sentinel-source_toolkit_fluxcd_io_git_repository_v1_manifest_test.go out/terratest-sentinel-source_toolkit_fluxcd_io_helm_chart_v1_manifest_test.go out/terratest-sentinel-source_toolkit_fluxcd_io_helm_repository_v1_manifest_test.go out/terratest-sentinel-source_toolkit_fluxcd_io_bucket_v1beta1_manifest_test.go out/terratest-sentinel-source_toolkit_fluxcd_io_git_repository_v1beta1_manifest_test.go out/terratest-sentinel-source_toolkit_fluxcd_io_helm_chart_v1beta1_manifest_test.go out/terratest-sentinel-source_toolkit_fluxcd_io_helm_repository_v1beta1_manifest_test.go out/terratest-sentinel-source_toolkit_fluxcd_io_bucket_v1beta2_manifest_test.go out/terratest-sentinel-source_toolkit_fluxcd_io_git_repository_v1beta2_manifest_test.go out/terratest-sentinel-source_toolkit_fluxcd_io_helm_chart_v1beta2_manifest_test.go out/terratest-sentinel-source_toolkit_fluxcd_io_helm_repository_v1beta2_manifest_test.go out/terratest-sentinel-source_toolkit_fluxcd_io_oci_repository_v1beta2_manifest_test.go out/terratest-sentinel-sparkoperator_k8s_io_scheduled_spark_application_v1beta2_manifest_test.go out/terratest-sentinel-sparkoperator_k8s_io_spark_application_v1beta2_manifest_test.go out/terratest-sentinel-spv_no_azure_key_vault_secret_v1_manifest_test.go out/terratest-sentinel-spv_no_azure_key_vault_identity_v1alpha1_manifest_test.go out/terratest-sentinel-spv_no_azure_key_vault_secret_v1alpha1_manifest_test.go out/terratest-sentinel-spv_no_azure_managed_identity_v1alpha1_manifest_test.go out/terratest-sentinel-spv_no_azure_key_vault_secret_v2alpha1_manifest_test.go out/terratest-sentinel-spv_no_azure_key_vault_secret_v2beta1_manifest_test.go out/terratest-sentinel-sqs_services_k8s_aws_queue_v1alpha1_manifest_test.go out/terratest-sentinel-storage_k8s_io_csi_driver_v1_manifest_test.go out/terratest-sentinel-storage_k8s_io_csi_node_v1_manifest_test.go out/terratest-sentinel-storage_k8s_io_storage_class_v1_manifest_test.go out/terratest-sentinel-storage_k8s_io_volume_attachment_v1_manifest_test.go out/terratest-sentinel-storage_kubeblocks_io_storage_provider_v1alpha1_manifest_test.go out/terratest-sentinel-storageos_com_storage_os_cluster_v1_manifest_test.go out/terratest-sentinel-sts_min_io_policy_binding_v1alpha1_manifest_test.go out/terratest-sentinel-sts_min_io_policy_binding_v1beta1_manifest_test.go out/terratest-sentinel-stunner_l7mp_io_dataplane_v1_manifest_test.go out/terratest-sentinel-stunner_l7mp_io_gateway_config_v1_manifest_test.go out/terratest-sentinel-stunner_l7mp_io_static_service_v1_manifest_test.go out/terratest-sentinel-stunner_l7mp_io_udp_route_v1_manifest_test.go out/terratest-sentinel-stunner_l7mp_io_dataplane_v1alpha1_manifest_test.go out/terratest-sentinel-stunner_l7mp_io_gateway_config_v1alpha1_manifest_test.go out/terratest-sentinel-stunner_l7mp_io_static_service_v1alpha1_manifest_test.go out/terratest-sentinel-submariner_io_broker_v1alpha1_manifest_test.go out/terratest-sentinel-submariner_io_service_discovery_v1alpha1_manifest_test.go out/terratest-sentinel-submariner_io_submariner_v1alpha1_manifest_test.go out/terratest-sentinel-telemetry_istio_io_telemetry_v1_manifest_test.go out/terratest-sentinel-telemetry_istio_io_telemetry_v1alpha1_manifest_test.go out/terratest-sentinel-templates_gatekeeper_sh_constraint_template_v1_manifest_test.go out/terratest-sentinel-templates_gatekeeper_sh_constraint_template_v1alpha1_manifest_test.go out/terratest-sentinel-templates_gatekeeper_sh_constraint_template_v1beta1_manifest_test.go out/terratest-sentinel-tempo_grafana_com_tempo_monolithic_v1alpha1_manifest_test.go out/terratest-sentinel-tempo_grafana_com_tempo_stack_v1alpha1_manifest_test.go out/terratest-sentinel-temporal_io_temporal_cluster_client_v1beta1_manifest_test.go out/terratest-sentinel-temporal_io_temporal_cluster_v1beta1_manifest_test.go out/terratest-sentinel-temporal_io_temporal_namespace_v1beta1_manifest_test.go out/terratest-sentinel-temporal_io_temporal_worker_process_v1beta1_manifest_test.go out/terratest-sentinel-tests_testkube_io_script_v1_manifest_test.go out/terratest-sentinel-tests_testkube_io_test_execution_v1_manifest_test.go out/terratest-sentinel-tests_testkube_io_test_source_v1_manifest_test.go out/terratest-sentinel-tests_testkube_io_test_suite_execution_v1_manifest_test.go out/terratest-sentinel-tests_testkube_io_test_suite_v1_manifest_test.go out/terratest-sentinel-tests_testkube_io_test_trigger_v1_manifest_test.go out/terratest-sentinel-tests_testkube_io_test_v1_manifest_test.go out/terratest-sentinel-tests_testkube_io_script_v2_manifest_test.go out/terratest-sentinel-tests_testkube_io_test_suite_v2_manifest_test.go out/terratest-sentinel-tests_testkube_io_test_v2_manifest_test.go out/terratest-sentinel-tests_testkube_io_test_suite_v3_manifest_test.go out/terratest-sentinel-tests_testkube_io_test_v3_manifest_test.go out/terratest-sentinel-tf_tungsten_io_analytics_alarm_v1alpha1_manifest_test.go out/terratest-sentinel-tf_tungsten_io_analytics_snmp_v1alpha1_manifest_test.go out/terratest-sentinel-tf_tungsten_io_analytics_v1alpha1_manifest_test.go out/terratest-sentinel-tf_tungsten_io_cassandra_v1alpha1_manifest_test.go out/terratest-sentinel-tf_tungsten_io_config_v1alpha1_manifest_test.go out/terratest-sentinel-tf_tungsten_io_control_v1alpha1_manifest_test.go out/terratest-sentinel-tf_tungsten_io_kubemanager_v1alpha1_manifest_test.go out/terratest-sentinel-tf_tungsten_io_manager_v1alpha1_manifest_test.go out/terratest-sentinel-tf_tungsten_io_query_engine_v1alpha1_manifest_test.go out/terratest-sentinel-tf_tungsten_io_rabbitmq_v1alpha1_manifest_test.go out/terratest-sentinel-tf_tungsten_io_redis_v1alpha1_manifest_test.go out/terratest-sentinel-tf_tungsten_io_vrouter_v1alpha1_manifest_test.go out/terratest-sentinel-tf_tungsten_io_webui_v1alpha1_manifest_test.go out/terratest-sentinel-tf_tungsten_io_zookeeper_v1alpha1_manifest_test.go out/terratest-sentinel-theketch_io_app_v1beta1_manifest_test.go out/terratest-sentinel-theketch_io_job_v1beta1_manifest_test.go out/terratest-sentinel-tinkerbell_org_hardware_v1alpha1_manifest_test.go out/terratest-sentinel-tinkerbell_org_osie_v1alpha1_manifest_test.go out/terratest-sentinel-tinkerbell_org_stack_v1alpha1_manifest_test.go out/terratest-sentinel-tinkerbell_org_template_v1alpha1_manifest_test.go out/terratest-sentinel-tinkerbell_org_workflow_v1alpha1_manifest_test.go out/terratest-sentinel-tinkerbell_org_hardware_v1alpha2_manifest_test.go out/terratest-sentinel-tinkerbell_org_osie_v1alpha2_manifest_test.go out/terratest-sentinel-tinkerbell_org_template_v1alpha2_manifest_test.go out/terratest-sentinel-tinkerbell_org_workflow_v1alpha2_manifest_test.go out/terratest-sentinel-topology_node_k8s_io_node_resource_topology_v1alpha1_manifest_test.go out/terratest-sentinel-topolvm_cybozu_com_logical_volume_v1_manifest_test.go out/terratest-sentinel-topolvm_cybozu_com_topolvm_cluster_v2_manifest_test.go out/terratest-sentinel-traefik_io_ingress_route_tcp_v1alpha1_manifest_test.go out/terratest-sentinel-traefik_io_ingress_route_udp_v1alpha1_manifest_test.go out/terratest-sentinel-traefik_io_ingress_route_v1alpha1_manifest_test.go out/terratest-sentinel-traefik_io_middleware_tcp_v1alpha1_manifest_test.go out/terratest-sentinel-traefik_io_middleware_v1alpha1_manifest_test.go out/terratest-sentinel-traefik_io_servers_transport_tcp_v1alpha1_manifest_test.go out/terratest-sentinel-traefik_io_servers_transport_v1alpha1_manifest_test.go out/terratest-sentinel-traefik_io_tls_option_v1alpha1_manifest_test.go out/terratest-sentinel-traefik_io_tls_store_v1alpha1_manifest_test.go out/terratest-sentinel-traefik_io_traefik_service_v1alpha1_manifest_test.go out/terratest-sentinel-training_kubedl_io_elastic_dl_job_v1alpha1_manifest_test.go out/terratest-sentinel-training_kubedl_io_mars_job_v1alpha1_manifest_test.go out/terratest-sentinel-training_kubedl_io_mpi_job_v1alpha1_manifest_test.go out/terratest-sentinel-training_kubedl_io_py_torch_job_v1alpha1_manifest_test.go out/terratest-sentinel-training_kubedl_io_tf_job_v1alpha1_manifest_test.go out/terratest-sentinel-training_kubedl_io_xdl_job_v1alpha1_manifest_test.go out/terratest-sentinel-training_kubedl_io_xg_boost_job_v1alpha1_manifest_test.go out/terratest-sentinel-trust_cert_manager_io_bundle_v1alpha1_manifest_test.go out/terratest-sentinel-upgrade_cattle_io_plan_v1_manifest_test.go out/terratest-sentinel-upgrade_managed_openshift_io_upgrade_config_v1alpha1_manifest_test.go out/terratest-sentinel-velero_io_backup_repository_v1_manifest_test.go out/terratest-sentinel-velero_io_backup_storage_location_v1_manifest_test.go out/terratest-sentinel-velero_io_backup_v1_manifest_test.go out/terratest-sentinel-velero_io_delete_backup_request_v1_manifest_test.go out/terratest-sentinel-velero_io_download_request_v1_manifest_test.go out/terratest-sentinel-velero_io_pod_volume_backup_v1_manifest_test.go out/terratest-sentinel-velero_io_pod_volume_restore_v1_manifest_test.go out/terratest-sentinel-velero_io_restore_v1_manifest_test.go out/terratest-sentinel-velero_io_schedule_v1_manifest_test.go out/terratest-sentinel-velero_io_server_status_request_v1_manifest_test.go out/terratest-sentinel-velero_io_volume_snapshot_location_v1_manifest_test.go out/terratest-sentinel-velero_io_data_download_v2alpha1_manifest_test.go out/terratest-sentinel-velero_io_data_upload_v2alpha1_manifest_test.go out/terratest-sentinel-virt_virtink_smartx_com_virtual_machine_migration_v1alpha1_manifest_test.go out/terratest-sentinel-virt_virtink_smartx_com_virtual_machine_v1alpha1_manifest_test.go out/terratest-sentinel-volsync_backube_replication_destination_v1alpha1_manifest_test.go out/terratest-sentinel-volsync_backube_replication_source_v1alpha1_manifest_test.go out/terratest-sentinel-vpcresources_k8s_aws_cni_node_v1alpha1_manifest_test.go out/terratest-sentinel-vpcresources_k8s_aws_security_group_policy_v1beta1_manifest_test.go out/terratest-sentinel-wgpolicyk8s_io_cluster_policy_report_v1alpha1_manifest_test.go out/terratest-sentinel-wgpolicyk8s_io_policy_report_v1alpha1_manifest_test.go out/terratest-sentinel-wgpolicyk8s_io_cluster_policy_report_v1alpha2_manifest_test.go out/terratest-sentinel-wgpolicyk8s_io_policy_report_v1alpha2_manifest_test.go out/terratest-sentinel-wgpolicyk8s_io_cluster_policy_report_v1beta1_manifest_test.go out/terratest-sentinel-wgpolicyk8s_io_policy_report_v1beta1_manifest_test.go out/terratest-sentinel-wildfly_org_wild_fly_server_v1alpha1_manifest_test.go out/terratest-sentinel-work_karmada_io_cluster_resource_binding_v1alpha1_manifest_test.go out/terratest-sentinel-work_karmada_io_resource_binding_v1alpha1_manifest_test.go out/terratest-sentinel-work_karmada_io_work_v1alpha1_manifest_test.go out/terratest-sentinel-work_karmada_io_cluster_resource_binding_v1alpha2_manifest_test.go out/terratest-sentinel-work_karmada_io_resource_binding_v1alpha2_manifest_test.go out/terratest-sentinel-workload_codeflare_dev_app_wrapper_v1beta1_manifest_test.go out/terratest-sentinel-workload_codeflare_dev_scheduling_spec_v1beta1_manifest_test.go out/terratest-sentinel-workload_codeflare_dev_app_wrapper_v1beta2_manifest_test.go out/terratest-sentinel-workloads_kubeblocks_io_instance_set_v1alpha1_manifest_test.go out/terratest-sentinel-workloads_kubeblocks_io_replicated_state_machine_v1alpha1_manifest_test.go out/terratest-sentinel-zonecontrol_k8s_aws_zone_aware_update_v1_manifest_test.go out/terratest-sentinel-zonecontrol_k8s_aws_zone_disruption_budget_v1_manifest_test.go out/terratest-sentinel-zookeeper_pravega_io_zookeeper_cluster_v1beta1_manifest_test.go  ## run all terratest tests
