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
	"networking_istio_io_virtual_service_v1alpha3": {
		"spec.http.mirror_percent",
	},
	"networking_istio_io_virtual_service_v1beta1": {
		"spec.http.mirror_percent",
	},
	"security_istio_io_request_authentication_v1beta1": {
		"spec.jwtRules.jwks_uri",
	},
	//"apps_stateful_set_v1": {
	//	"spec.volumeClaimTemplates.apiVersion",
	//	"spec.volumeClaimTemplates.kind",
	//	"spec.volumeClaimTemplates.metadata",
	//	"spec.volumeClaimTemplates.status",
	//},
}
