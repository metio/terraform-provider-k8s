/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package generator

var pluralForms = map[string]string{
	"Binding":                        "bindings",
	"ComponentStatus":                "componentstatuses",
	"ConfigMap":                      "configmaps",
	"Endpoints":                      "endpoints",
	"Event":                          "events",
	"LimitRange":                     "limitranges",
	"Namespace":                      "namespaces",
	"Node":                           "nodes",
	"PersistentVolumeClaim":          "persistentvolumeclaims",
	"PersistentVolume":               "persistentvolumes",
	"Pod":                            "pods",
	"PodTemplate":                    "podtemplates",
	"ReplicationController":          "replicationcontrollers",
	"ResourceQuota":                  "resourcequotas",
	"Secret":                         "secrets",
	"ServiceAccount":                 "serviceaccounts",
	"Service":                        "services",
	"MutatingWebhookConfiguration":   "mutatingwebhookconfigurations",
	"ValidatingWebhookConfiguration": "validatingwebhookconfigurations",
	"CustomResourceDefinition":       "customresourcedefinitions",
	"APIService":                     "apiservices",
	"ControllerRevision":             "controllerrevisions",
	"DaemonSet":                      "daemonsets",
	"Deployment":                     "deployments",
	"ReplicaSet":                     "replicasets",
	"StatefulSet":                    "statefulsets",
	"TokenReview":                    "tokenreviews",
	"LocalSubjectAccessReview":       "localsubjectaccessreviews",
	"SelfSubjectAccessReview":        "selfsubjectaccessreviews",
	"SelfSubjectRulesReview":         "selfsubjectrulesreviews",
	"SubjectAccessReview":            "subjectaccessreviews",
	"HorizontalPodAutoscaler":        "horizontalpodautoscalers",
	"CronJob":                        "cronjobs",
	"Job":                            "jobs",
	"CertificateSigningRequest":      "certificatesigningrequests",
	"Lease":                          "leases",
	"EndpointSlice":                  "endpointslices",
	"FlowSchema":                     "flowschemas",
	"PriorityLevelConfiguration":     "prioritylevelconfigurations",
	"NodeMetrics":                    "nodes",
	"PodMetrics":                     "pods",
	"IngressClass":                   "ingressclasses",
	"Ingress":                        "ingresses",
	"NetworkPolicy":                  "networkpolicies",
	"RuntimeClass":                   "runtimeclasses",
	"PodDisruptionBudget":            "poddisruptionbudgets",
	"ClusterRoleBinding":             "clusterrolebindings",
	"ClusterRole":                    "clusterroles",
	"RoleBinding":                    "rolebindings",
	"Role":                           "roles",
	"PriorityClass":                  "priorityclasses",
	"VolumeSnapshotClass":            "volumesnapshotclasses",
	"VolumeSnapshotContent":          "volumesnapshotcontents",
	"VolumeSnapshot":                 "volumesnapshots",
	"CSIDriver":                      "csidrivers",
	"CSINode":                        "csinodes",
	"CSIStorageCapacity":             "csistoragecapacities",
	"StorageClass":                   "storageclasses",
	"VolumeAttachment":               "volumeattachments",
}

func pluralForm(original string) string {
	if plural, ok := pluralForms[original]; ok {
		return plural
	}
	return original
}
