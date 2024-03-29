/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package main

import "k8s.io/utils/strings/slices"

var namespacedObjects = []string{
	"io.k8s.api.apps.v1.DaemonSet",
	"io.k8s.api.apps.v1.Deployment",
	"io.k8s.api.apps.v1.ReplicaSet",
	"io.k8s.api.apps.v1.StatefulSet",
	"io.k8s.api.autoscaling.v1.HorizontalPodAutoscaler",
	"io.k8s.api.autoscaling.v2.HorizontalPodAutoscaler",
	"io.k8s.api.batch.v1.CronJob",
	"io.k8s.api.batch.v1.Job",
	"io.k8s.api.core.v1.ConfigMap",
	"io.k8s.api.core.v1.Endpoints",
	"io.k8s.api.core.v1.LimitRange",
	"io.k8s.api.core.v1.PersistentVolume",
	"io.k8s.api.core.v1.PersistentVolumeClaim",
	"io.k8s.api.core.v1.Pod",
	"io.k8s.api.core.v1.ReplicationController",
	"io.k8s.api.core.v1.Secret",
	"io.k8s.api.core.v1.Service",
	"io.k8s.api.core.v1.ServiceAccount",
	"io.k8s.api.discovery.v1.EndpointSlice",
	"io.k8s.api.events.v1.Event",
	"io.k8s.api.networking.v1.Ingress",
	"io.k8s.api.networking.v1.NetworkPolicy",
	"io.k8s.api.policy.v1.PodDisruptionBudget",
	"io.k8s.api.rbac.v1.Role",
	"io.k8s.api.rbac.v1.RoleBinding",
}

func isNamespacedObject(name string) bool {
	return slices.Contains(namespacedObjects, name)
}
