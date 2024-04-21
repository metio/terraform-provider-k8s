/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package generator

var ignoredAttributes = map[string][]string{
	"acid_zalan_do_postgresql_v1": {
		"spec.init_containers",
		"spec.pod_priority_class_name",
	},
	"crd_projectcalico_org_ip_pool_v1": {
		"spec.nat-outgoing",
	},
	"getambassador_io_host_v2": {
		"spec.ambassadorId",
	},
	"logging_banzaicloud_io_fluentbit_agent_v1beta1": {
		"spec.bufferStorageVolume.host_path",
		"spec.position_db",
		"spec.positiondb.host_path",
	},
	"logging_banzaicloud_io_logging_v1beta1": {
		"spec.fluentbit.bufferStorageVolume.host_path",
		"spec.fluentbit.position_db",
		"spec.fluentbit.positiondb.host_path",
		"spec.fluentd.bufferStorageVolume.host_path",
		"spec.fluentd.extraVolumes.volume.host_path",
		"spec.fluentd.fluentdPvcSpec.host_path",
		"spec.nodeAgents.nodeAgentFluentbit.bufferStorageVolume.host_path",
		"spec.nodeAgents.nodeAgentFluentbit.positiondb.host_path",
	},
	"logging_banzaicloud_io_node_agent_v1beta1": {
		"spec.nodeAgentFluentbit.bufferStorageVolume.host_path",
		"spec.nodeAgentFluentbit.positiondb.host_path",
	},
	"logging_extensions_banzaicloud_io_event_tailer_v1alpha1": {
		"spec.positionVolume.host_path",
	},
	"networking_istio_io_virtual_service_v1alpha3": {
		"spec.http.mirror_percent",
	},
	"networking_istio_io_virtual_service_v1beta1": {
		"spec.http.mirror_percent",
	},
	"networking_istio_io_virtual_service_v1": {
		"spec.http.mirror_percent",
	},
	"security_istio_io_request_authentication_v1beta1": {
		"spec.jwtRules.jwks_uri",
	},
	"security_istio_io_request_authentication_v1": {
		"spec.jwtRules.jwks_uri",
	},
	"operator_victoriametrics_com_vm_agent_v1beta1": {
		"spec.inlineRelabelConfig.source_labels",
		"spec.inlineRelabelConfig.target_label",
		"spec.nodeScrapeRelabelTemplate.source_labels",
		"spec.nodeScrapeRelabelTemplate.target_label",
		"spec.podScrapeRelabelTemplate.source_labels",
		"spec.podScrapeRelabelTemplate.target_label",
		"spec.probeScrapeRelabelTemplate.source_labels",
		"spec.probeScrapeRelabelTemplate.target_label",
		"spec.remoteWrite.inlineUrlRelabelConfig.source_labels",
		"spec.remoteWrite.inlineUrlRelabelConfig.target_label",
		"spec.remoteWrite.streamAggrConfig.rules.input_relabel_configs.source_labels",
		"spec.remoteWrite.streamAggrConfig.rules.input_relabel_configs.target_label",
		"spec.remoteWrite.streamAggrConfig.rules.output_relabel_configs.source_labels",
		"spec.remoteWrite.streamAggrConfig.rules.output_relabel_configs.target_label",
		"spec.scrapeConfigRelabelTemplate.source_labels",
		"spec.scrapeConfigRelabelTemplate.target_label",
		"spec.serviceScrapeRelabelTemplate.source_labels",
		"spec.serviceScrapeRelabelTemplate.target_label",
		"spec.staticScrapeRelabelTemplate.source_labels",
		"spec.staticScrapeRelabelTemplate.target_label",
	},
	"operator_victoriametrics_com_vm_node_scrape_v1beta1": {
		"spec.metricRelabelConfigs.source_labels",
		"spec.metricRelabelConfigs.target_label",
		"spec.relabelConfigs.source_labels",
		"spec.relabelConfigs.target_label",
	},
	"operator_victoriametrics_com_vm_pod_scrape_v1beta1": {
		"spec.podMetricsEndpoints.metricRelabelConfigs.source_labels",
		"spec.podMetricsEndpoints.metricRelabelConfigs.target_label",
		"spec.podMetricsEndpoints.relabelConfigs.source_labels",
		"spec.podMetricsEndpoints.relabelConfigs.target_label",
	},
	"operator_victoriametrics_com_vm_probe_v1beta1": {
		"spec.targets.ingress.relabelingConfigs.source_labels",
		"spec.targets.ingress.relabelingConfigs.target_label",
		"spec.targets.staticConfig.relabelingConfigs.source_labels",
		"spec.targets.staticConfig.relabelingConfigs.target_label",
	},
	"operator_victoriametrics_com_vm_service_scrape_v1beta1": {
		"spec.endpoints.metricRelabelConfigs.source_labels",
		"spec.endpoints.metricRelabelConfigs.target_label",
		"spec.endpoints.relabelConfigs.source_labels",
		"spec.endpoints.relabelConfigs.target_label",
	},
	"operator_victoriametrics_com_vm_single_v1beta1": {
		"spec.streamAggrConfig.rules.input_relabel_configs.source_labels",
		"spec.streamAggrConfig.rules.input_relabel_configs.target_label",
		"spec.streamAggrConfig.rules.output_relabel_configs.source_labels",
		"spec.streamAggrConfig.rules.output_relabel_configs.target_label",
	},
	"operator_victoriametrics_com_vm_static_scrape_v1beta1": {
		"spec.targetEndpoints.metricRelabelConfigs.source_labels",
		"spec.targetEndpoints.metricRelabelConfigs.target_label",
		"spec.targetEndpoints.relabelConfigs.source_labels",
		"spec.targetEndpoints.relabelConfigs.target_label",
	},
	//"apps_stateful_set_v1": {
	//	"spec.volumeClaimTemplates.apiVersion",
	//	"spec.volumeClaimTemplates.kind",
	//	"spec.volumeClaimTemplates.metadata",
	//	"spec.volumeClaimTemplates.status",
	//},
}
