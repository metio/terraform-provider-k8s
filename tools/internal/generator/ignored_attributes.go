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
	"security_istio_io_request_authentication_v1beta1": {
		"spec.jwtRules.jwks_uri",
	},
	"security_istio_io_request_authentication_v1": {
		"spec.jwtRules.jwks_uri",
	},
	//"apps_stateful_set_v1": {
	//	"spec.volumeClaimTemplates.apiVersion",
	//	"spec.volumeClaimTemplates.kind",
	//	"spec.volumeClaimTemplates.metadata",
	//	"spec.volumeClaimTemplates.status",
	//},
}
