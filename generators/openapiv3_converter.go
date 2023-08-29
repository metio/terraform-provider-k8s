//go:build generators

/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package main

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"k8s.io/utils/strings/slices"
	"sort"
	"strings"
)

func convertOpenAPIv3(schemas []map[string]*openapi3.SchemaRef) []*TemplateData {
	data := make([]*TemplateData, 0)
	for _, schema := range schemas {
		for name, definition := range schema {
			if supportedOpenAPIv3Object(name, definition) {
				namespaced := isNamespacedObject(name)
				templateData := openAPIv3AsTemplateData(definition, namespaced)
				data = append(data, templateData)
			}
		}
	}
	return data
}

func openAPIv3AsTemplateData(definition *openapi3.SchemaRef, namespaced bool) *TemplateData {
	var group string
	var version string
	var kind string
	if gvkExt, ok := definition.Value.Extensions["x-kubernetes-group-version-kind"]; ok {
		raw := gvkExt.([]interface{})
		gvk := raw[0].(map[string]interface{})
		group = gvk["group"].(string)
		version = gvk["version"].(string)
		kind = gvk["kind"].(string)
	}
	schema := definition.Value
	// remove manually managed or otherwise ignored properties
	delete(schema.Properties, "metadata")
	delete(schema.Properties, "status")
	delete(schema.Properties, "apiVersion")
	delete(schema.Properties, "kind")

	imports := AdditionalImports{}
	typeName := resourceTypeName(group, kind, version)

	return &TemplateData{
		BT:          "`",
		Package:     goPackageName(group, version),
		Group:       group,
		Version:     version,
		Kind:        kind,
		Namespaced:  namespaced,
		Description: description(schema.Description),

		ResourceFile:         resourceFile(group, kind, version),
		ResourceTestFile:     resourceTestFile(group, kind, version),
		ResourceWorkflowFile: resourceWorkflowFile(group, kind, version),
		ResourceTypeName:     typeName,
		FullResourceTypeName: fmt.Sprintf("k8s_%s", typeName),
		ResourceTypeStruct:   resourceTypeStruct(group, kind, version),
		ResourceDataStruct:   resourceDataStruct(group, kind, version),
		ResourceTypeTest:     resourceTypeTest(group, kind, version),

		DataSourceFile:         dataSourceFile(group, kind, version),
		DataSourceTestFile:     dataSourceTestFile(group, kind, version),
		DataSourceWorkflowFile: dataSourceWorkflowFile(group, kind, version),
		DataSourceTypeName:     typeName,
		FullDataSourceTypeName: fmt.Sprintf("k8s_%s", typeName),
		DataSourceTypeStruct:   dataSourceTypeStruct(group, kind, version),
		DataSourceDataStruct:   dataSourceDataStruct(group, kind, version),
		DataSourceTypeTest:     dataSourceTypeTest(group, kind, version),

		ManifestFile:         manifestFile(group, kind, version),
		ManifestTestFile:     manifestTestFile(group, kind, version),
		ManifestWorkflowFile: manifestWorkflowFile(group, kind, version),
		ManifestTypeName:     fmt.Sprintf("%s_manifest", typeName),
		FullManifestTypeName: fmt.Sprintf("k8s_%s_manifest", typeName),
		ManifestTypeStruct:   manifestTypeStruct(group, kind, version),
		ManifestDataStruct:   manifestDataStruct(group, kind, version),
		ManifestTypeTest:     manifestTypeTest(group, kind, version),

		TerraformModelType: terraformModelType(group, kind, version),

		AdditionalImports: imports,
		Properties:        openAPIv3Properties(schema, &imports, "", typeName),
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
				} else if prop.Value.Type == "object" && prop.Value.AdditionalProperties.Schema != nil && prop.Value.AdditionalProperties.Schema.Value.Type == "object" {
					nestedProperties = openAPIv3Properties(prop.Value.AdditionalProperties.Schema.Value, imports, propPath, terraformResourceName)
				} else {
					nestedProperties = openAPIv3Properties(prop.Value, imports, propPath, terraformResourceName)
				}

				attributeType, valueType, elementType, goType := translateTypeWith(&openapiv3TypeTranslator{property: prop.Value})

				validators := validatorsFor(&openapiv3ValidatorExtractor{
					property: prop.Value,
					imports:  imports,
				}, terraformResourceName, propPath, imports)

				props = append(props, &Property{
					BT:                     "`",
					Name:                   name,
					GoName:                 goName(name),
					GoType:                 goType,
					TerraformAttributeName: terraformAttributeName(name, path == ""),
					TerraformAttributeType: attributeType,
					TerraformElementType:   elementType,
					TerraformValueType:     valueType,
					Description:            description(prop.Value.Description),
					Required:               slices.Contains(schema.Required, name),
					Optional:               !slices.Contains(schema.Required, name),
					Computed:               false,
					Properties:             nestedProperties,
					ValidatorsType:         mapAttributeTypeToValidatorsType(attributeType),
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
				var pathExpressions []string
				for _, inner := range props {
					if outer.Name != inner.Name {
						pathExpressions = append(pathExpressions, fmt.Sprintf(`path.MatchRelative().AtParent().AtName("%s")`, inner.TerraformAttributeName))
					}
				}
				validator := fmt.Sprintf(`schemavalidator.ExactlyOneOf(%v)`, strings.Join(pathExpressions, ", "))
				outer.Validators = append(outer.Validators, validator)
			}
		}
	} else if schema.MinProps > 0 && schema.MaxProps == nil {
		min := schema.MinProps

		if min == 1 {
			imports.SchemaValidator = true

			for _, outer := range props {
				var pathExpressions []string
				for _, inner := range props {
					if outer.Name != inner.Name {
						pathExpressions = append(pathExpressions, fmt.Sprintf(`path.MatchRelative().AtParent().AtName("%s")`, inner.TerraformAttributeName))
					}
				}
				validator := fmt.Sprintf(`schemavalidator.AtLeastOneOf(%v)`, strings.Join(pathExpressions, ", "))
				outer.Validators = append(outer.Validators, validator)
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
