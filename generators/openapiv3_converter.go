//go:build generators

/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package main

import (
	"encoding/json"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"k8s.io/utils/strings/slices"
	"sort"
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
}

func convertOpenAPIv3(schemas []map[string]*openapi3.SchemaRef, pkg string) []*TemplateData {
	data := make([]*TemplateData, 0)
	for _, schema := range schemas {
		for name, definition := range schema {
			if !strings.HasPrefix(name, "io.k8s") || slices.Contains(supportedKubernetesApiObjects, name) {
				if _, ok := definition.Value.ExtensionProps.Extensions["x-kubernetes-group-version-kind"]; ok {
					if templateData := openAPIv3AsTemplateData(definition, pkg); templateData != nil {
						data = append(data, templateData)
					}
				}
			}
		}
	}
	return data
}

type GVK struct {
	Group   string `json:"group"`
	Version string `json:"version"`
	Kind    string `json:"kind"`
}

func openAPIv3AsTemplateData(definition *openapi3.SchemaRef, pkg string) *TemplateData {
	var group string
	var version string
	var kind string
	if gvkExt, ok := definition.Value.ExtensionProps.Extensions["x-kubernetes-group-version-kind"]; ok {
		raw := gvkExt.(json.RawMessage)
		var gvks []GVK
		if err := json.Unmarshal(raw, &gvks); err != nil {
			return nil
		}
		if len(gvks) != 1 {
			return nil
		}
		gvk := gvks[0]
		group = gvk.Group
		version = gvk.Version
		kind = gvk.Kind
	}
	schema := definition.Value
	// remove manually managed or otherwise ignored properties
	delete(schema.Properties, "metadata")
	delete(schema.Properties, "status")
	delete(schema.Properties, "apiVersion")
	delete(schema.Properties, "kind")

	if len(schema.Properties) == 0 {
		return nil
	}

	imports := AdditionalImports{}
	trn := terraformResourceName(group, kind, version)

	return &TemplateData{
		BT:                    "`",
		Package:               pkg,
		File:                  terraformResourceFile(group, kind, version),
		Group:                 group,
		Version:               version,
		Kind:                  kind,
		Namespaced:            true,
		Description:           description(schema.Description),
		TerraformResourceType: terraformResourceType(group, kind, version),
		TerraformModelType:    terraformModelType(group, kind, version),
		GoModelType:           goModelType(group, kind, version),
		Properties:            openAPIv3Properties(schema, &imports, "", trn),
		TerraformResourceName: trn,
		AdditionalImports:     imports,
	}
}

func openAPIv3Properties(schema *openapi3.Schema, imports *AdditionalImports, path string, terraformResourceName string) []*Property {
	props := make([]*Property, 0)

	if schema != nil {
		for name, prop := range schema.Properties {
			if prop.Value != nil {
				propPath := propertyPath(path, name)
				if ignored, ok := ignoredAttributes[terraformResourceName]; ok {
					if slices.Contains(ignored, propPath) {
						continue
					}
				}

				var nestedProperties []*Property
				if prop.Value.Type == "array" && prop.Value.Items != nil && prop.Value.Items.Value != nil && prop.Value.Items.Value.Type == "object" {
					nestedProperties = openAPIv3Properties(prop.Value.Items.Value, imports, propPath, terraformResourceName)
				} else if prop.Value.Type == "object" && prop.Value.AdditionalProperties != nil && prop.Value.AdditionalProperties.Value.Type == "object" {
					nestedProperties = openAPIv3Properties(prop.Value.AdditionalProperties.Value, imports, propPath, terraformResourceName)
				} else {
					nestedProperties = openAPIv3Properties(prop.Value, imports, propPath, terraformResourceName)
				}

				attributeType, valueType, goType := translateTypeWith(&openapiv3TypeTranslator{property: prop.Value})

				validators := validatorsFor(&openapiv3ValidatorExtractor{
					property: prop.Value,
					imports:  imports,
				}, terraformResourceName, propPath, imports)

				props = append(props, &Property{
					BT:                     "`",
					Name:                   name,
					GoName:                 goName(name),
					GoType:                 goType,
					TerraformAttributeName: terraformAttributeName(name),
					TerraformAttributeType: attributeType,
					TerraformValueType:     valueType,
					Description:            description(prop.Value.Description),
					Required:               slices.Contains(schema.Required, name),
					Optional:               !slices.Contains(schema.Required, name),
					Computed:               false,
					Properties:             nestedProperties,
					Validators:             validators,
				})
			}
		}
	}

	sort.SliceStable(props, func(i, j int) bool {
		return props[i].Name < props[j].Name
	})

	if schema.MinProps > 0 && schema.MaxProps != nil {
		min := schema.MinProps
		max := schema.MaxProps

		if min == 1 && *max == 1 {
			imports.SchemaValidator = true

			for _, outer := range props {
				for _, inner := range props {
					if outer.Name != inner.Name {
						validator := fmt.Sprintf(`schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("%s"))`, inner.TerraformAttributeName)
						outer.Validators = append(outer.Validators, validator)
					}
				}
			}
		}
	} else if schema.MinProps > 0 && schema.MaxProps == nil {
		min := schema.MinProps

		if min == 1 {
			imports.SchemaValidator = true

			for _, outer := range props {
				for _, inner := range props {
					if outer.Name != inner.Name {
						validator := fmt.Sprintf(`schemavalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("%s"))`, inner.TerraformAttributeName)
						outer.Validators = append(outer.Validators, validator)
					}
				}
			}
		} else if min > 1 && min == uint64(len(props)) {
			for _, prop := range props {
				prop.Required = true
				prop.Optional = false
			}
		}
	}

	return props
}
