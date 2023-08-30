/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package generator

import (
	"github.com/getkin/kin-openapi/openapi3"
	"k8s.io/utils/strings/slices"
	"strings"
)

var supportedKubernetesApiObjects = []string{
	"io.k8s.api.admissionregistration.v1.MutatingWebhookConfiguration",
	"io.k8s.api.admissionregistration.v1.ValidatingWebhookConfiguration",
	"io.k8s.api.apps.v1.DaemonSet",
	"io.k8s.api.apps.v1.Deployment",
	"io.k8s.api.apps.v1.ReplicaSet",
	"io.k8s.api.apps.v1.StatefulSet",
	"io.k8s.api.autoscaling.v1.HorizontalPodAutoscaler",
	"io.k8s.api.autoscaling.v2.HorizontalPodAutoscaler",
	"io.k8s.api.batch.v1.CronJob",
	"io.k8s.api.batch.v1.Job",
	"io.k8s.api.certificates.v1.CertificateSigningRequest",
	"io.k8s.api.core.v1.ConfigMap",
	"io.k8s.api.core.v1.Endpoints",
	"io.k8s.api.core.v1.LimitRange",
	"io.k8s.api.core.v1.Namespace",
	"io.k8s.api.core.v1.PersistentVolume",
	"io.k8s.api.core.v1.PersistentVolumeClaim",
	"io.k8s.api.core.v1.Pod",
	"io.k8s.api.core.v1.ReplicationController",
	"io.k8s.api.core.v1.Secret",
	"io.k8s.api.core.v1.Service",
	"io.k8s.api.core.v1.ServiceAccount",
	"io.k8s.api.discovery.v1.EndpointSlice",
	"io.k8s.api.events.v1.Event",
	"io.k8s.api.flowcontrol.v1beta2.FlowSchema",
	"io.k8s.api.flowcontrol.v1beta2.PriorityLevelConfiguration",
	"io.k8s.api.flowcontrol.v1beta3.FlowSchema",
	"io.k8s.api.flowcontrol.v1beta3.PriorityLevelConfiguration",
	"io.k8s.api.networking.v1.Ingress",
	"io.k8s.api.networking.v1.IngressClass",
	"io.k8s.api.networking.v1.NetworkPolicy",
	"io.k8s.api.policy.v1.PodDisruptionBudget",
	"io.k8s.api.rbac.v1.ClusterRole",
	"io.k8s.api.rbac.v1.ClusterRoleBinding",
	"io.k8s.api.rbac.v1.Role",
	"io.k8s.api.rbac.v1.RoleBinding",
	"io.k8s.api.scheduling.v1.PriorityClass",
	"io.k8s.api.storage.v1.CSIDriver",
	"io.k8s.api.storage.v1.CSINode",
	"io.k8s.api.storage.v1.StorageClass",
	"io.k8s.api.storage.v1.VolumeAttachment",
	"io.k8s.kube-aggregator.pkg.apis.apiregistration.v1.APIService",
}

func supportedOpenAPIv3Object(name string, definition *openapi3.SchemaRef) bool {
	if !strings.HasPrefix(name, "io.k8s") || slices.Contains(supportedKubernetesApiObjects, name) {
		if _, ok := definition.Value.Extensions["x-kubernetes-group-version-kind"]; ok {
			if len(definition.Value.Properties) > 0 {
				return true
			}
		}
	}

	return false
}
