/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/metio/terraform-provider-k8s/internal/provider/admissionregistration_k8s_io_v1"
	"github.com/metio/terraform-provider-k8s/internal/provider/apiregistration_k8s_io_v1"
	"github.com/metio/terraform-provider-k8s/internal/provider/apps_v1"
	"github.com/metio/terraform-provider-k8s/internal/provider/autoscaling_v1"
	"github.com/metio/terraform-provider-k8s/internal/provider/autoscaling_v2"
	"github.com/metio/terraform-provider-k8s/internal/provider/batch_v1"
	"github.com/metio/terraform-provider-k8s/internal/provider/certificates_k8s_io_v1"
	"github.com/metio/terraform-provider-k8s/internal/provider/core_v1"
	"github.com/metio/terraform-provider-k8s/internal/provider/discovery_k8s_io_v1"
	"github.com/metio/terraform-provider-k8s/internal/provider/events_k8s_io_v1"
	"github.com/metio/terraform-provider-k8s/internal/provider/flowcontrol_apiserver_k8s_io_v1beta2"
	"github.com/metio/terraform-provider-k8s/internal/provider/flowcontrol_apiserver_k8s_io_v1beta3"
	"github.com/metio/terraform-provider-k8s/internal/provider/networking_k8s_io_v1"
	"github.com/metio/terraform-provider-k8s/internal/provider/policy_v1"
	"github.com/metio/terraform-provider-k8s/internal/provider/rbac_authorization_k8s_io_v1"
	"github.com/metio/terraform-provider-k8s/internal/provider/scheduling_k8s_io_v1"
	"github.com/metio/terraform-provider-k8s/internal/provider/storage_k8s_io_v1"
)

func allResources() []func() resource.Resource {
	return []func() resource.Resource{
		admissionregistration_k8s_io_v1.NewAdmissionregistrationK8SIoMutatingWebhookConfigurationV1Resource,
		admissionregistration_k8s_io_v1.NewAdmissionregistrationK8SIoValidatingWebhookConfigurationV1Resource,
		apiregistration_k8s_io_v1.NewApiregistrationK8SIoAPIServiceV1Resource,
		apps_v1.NewAppsDaemonSetV1Resource,
		apps_v1.NewAppsDeploymentV1Resource,
		apps_v1.NewAppsReplicaSetV1Resource,
		apps_v1.NewAppsStatefulSetV1Resource,
		autoscaling_v1.NewAutoscalingHorizontalPodAutoscalerV1Resource,
		autoscaling_v2.NewAutoscalingHorizontalPodAutoscalerV2Resource,
		batch_v1.NewBatchCronJobV1Resource,
		batch_v1.NewBatchJobV1Resource,
		certificates_k8s_io_v1.NewCertificatesK8SIoCertificateSigningRequestV1Resource,
		core_v1.NewConfigMapV1Resource,
		core_v1.NewEndpointsV1Resource,
		core_v1.NewLimitRangeV1Resource,
		core_v1.NewNamespaceV1Resource,
		core_v1.NewPersistentVolumeClaimV1Resource,
		core_v1.NewPersistentVolumeV1Resource,
		core_v1.NewPodV1Resource,
		core_v1.NewReplicationControllerV1Resource,
		core_v1.NewSecretV1Resource,
		core_v1.NewServiceAccountV1Resource,
		core_v1.NewServiceV1Resource,
		discovery_k8s_io_v1.NewDiscoveryK8SIoEndpointSliceV1Resource,
		events_k8s_io_v1.NewEventsK8SIoEventV1Resource,
		flowcontrol_apiserver_k8s_io_v1beta2.NewFlowcontrolApiserverK8SIoFlowSchemaV1Beta2Resource,
		flowcontrol_apiserver_k8s_io_v1beta2.NewFlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta2Resource,
		flowcontrol_apiserver_k8s_io_v1beta3.NewFlowcontrolApiserverK8SIoFlowSchemaV1Beta3Resource,
		flowcontrol_apiserver_k8s_io_v1beta3.NewFlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta3Resource,
		networking_k8s_io_v1.NewNetworkingK8SIoIngressClassV1Resource,
		networking_k8s_io_v1.NewNetworkingK8SIoIngressV1Resource,
		networking_k8s_io_v1.NewNetworkingK8SIoNetworkPolicyV1Resource,
		policy_v1.NewPolicyPodDisruptionBudgetV1Resource,
		rbac_authorization_k8s_io_v1.NewRbacAuthorizationK8SIoClusterRoleBindingV1Resource,
		rbac_authorization_k8s_io_v1.NewRbacAuthorizationK8SIoClusterRoleV1Resource,
		rbac_authorization_k8s_io_v1.NewRbacAuthorizationK8SIoRoleBindingV1Resource,
		rbac_authorization_k8s_io_v1.NewRbacAuthorizationK8SIoRoleV1Resource,
		scheduling_k8s_io_v1.NewSchedulingK8SIoPriorityClassV1Resource,
		storage_k8s_io_v1.NewStorageK8SIoCSIDriverV1Resource,
		storage_k8s_io_v1.NewStorageK8SIoCSINodeV1Resource,
		storage_k8s_io_v1.NewStorageK8SIoStorageClassV1Resource,
		storage_k8s_io_v1.NewStorageK8SIoVolumeAttachmentV1Resource,
	}
}
