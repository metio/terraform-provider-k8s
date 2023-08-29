//go:build generators

/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package main

import (
	"fmt"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/utils/strings/slices"
	"sort"
	"strings"
)

func convertCRDv1(crds []*apiextensionsv1.CustomResourceDefinition) []*TemplateData {
	data := make([]*TemplateData, 0)
	for _, crd := range crds {
		if templateData := crdV1AsTemplateData(crd); templateData != nil {
			data = append(data, templateData)
		}
	}
	return data
}

func crdV1AsTemplateData(crd *apiextensionsv1.CustomResourceDefinition) *TemplateData {
	group := crd.Spec.Group
	version := crd.Spec.Versions[0]
	kind := crd.Spec.Names.Kind

	schema := version.Schema.OpenAPIV3Schema
	// remove manually managed or otherwise ignored properties
	delete(schema.Properties, "metadata")
	delete(schema.Properties, "status")
	delete(schema.Properties, "apiVersion")
	delete(schema.Properties, "kind")

	if len(schema.Properties) == 0 {
		return nil
	}

	imports := AdditionalImports{}
	typeName := resourceTypeName(group, kind, version.Name)

	return &TemplateData{
		BT:          "`",
		Package:     goPackageName(group, version.Name),
		Group:       group,
		Version:     version.Name,
		Kind:        kind,
		Namespaced:  crd.Spec.Scope == apiextensionsv1.NamespaceScoped,
		Description: description(schema.Description),

		ResourceFile:         resourceFile(group, kind, version.Name),
		ResourceTestFile:     resourceTestFile(group, kind, version.Name),
		ResourceWorkflowFile: resourceWorkflowFile(group, kind, version.Name),
		ResourceTypeName:     typeName,
		FullResourceTypeName: fmt.Sprintf("k8s_%s", typeName),
		ResourceTypeStruct:   resourceTypeStruct(group, kind, version.Name),
		ResourceDataStruct:   resourceDataStruct(group, kind, version.Name),
		ResourceTypeTest:     resourceTypeTest(group, kind, version.Name),

		DataSourceFile:         dataSourceFile(group, kind, version.Name),
		DataSourceTestFile:     dataSourceTestFile(group, kind, version.Name),
		DataSourceWorkflowFile: dataSourceWorkflowFile(group, kind, version.Name),
		DataSourceTypeName:     typeName,
		FullDataSourceTypeName: fmt.Sprintf("k8s_%s", typeName),
		DataSourceTypeStruct:   dataSourceTypeStruct(group, kind, version.Name),
		DataSourceDataStruct:   dataSourceDataStruct(group, kind, version.Name),
		DataSourceTypeTest:     dataSourceTypeTest(group, kind, version.Name),

		ManifestFile:         manifestFile(group, kind, version.Name),
		ManifestTestFile:     manifestTestFile(group, kind, version.Name),
		ManifestWorkflowFile: manifestWorkflowFile(group, kind, version.Name),
		ManifestTypeName:     fmt.Sprintf("%s_manifest", typeName),
		FullManifestTypeName: fmt.Sprintf("k8s_%s_manifest", typeName),
		ManifestTypeStruct:   manifestTypeStruct(group, kind, version.Name),
		ManifestDataStruct:   manifestDataStruct(group, kind, version.Name),
		ManifestTypeTest:     manifestTypeTest(group, kind, version.Name),

		TerraformModelType: terraformModelType(group, kind, version.Name),

		AdditionalImports: imports,
		Properties:        crdV1Properties(schema, &imports, "", typeName),
	}
}

func crdV1Properties(schema *apiextensionsv1.JSONSchemaProps, imports *AdditionalImports, path string, terraformResourceName string) []*Property {
	props := make([]*Property, 0)

	for name, prop := range schema.Properties {
		propPath := propertyPath(path, name)
		if ignored, ok := ignoredAttributes[terraformResourceName]; ok {
			if slices.Contains(ignored, propPath) {
				continue
			}
		}

		var nestedProperties []*Property
		if prop.Type == "array" && prop.Items.Schema.Type == "object" {
			nestedProperties = crdV1Properties(prop.Items.Schema, imports, propPath, terraformResourceName)
		} else if prop.Type == "object" && prop.AdditionalProperties != nil && prop.AdditionalProperties.Schema.Type == "object" {
			nestedProperties = crdV1Properties(prop.AdditionalProperties.Schema, imports, propPath, terraformResourceName)
		} else {
			nestedProperties = crdV1Properties(&prop, imports, propPath, terraformResourceName)
		}

		attributeType, valueType, elementType, goType := translateTypeWith(&crd1TypeTranslator{property: &prop})

		validators := validatorsFor(&crdv1ValidatorExtractor{
			property: &prop,
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
			Description:            description(prop.Description),
			Required:               slices.Contains(schema.Required, name),
			Optional:               !slices.Contains(schema.Required, name),
			Computed:               false,
			Properties:             nestedProperties,
			ValidatorsType:         mapAttributeTypeToValidatorsType(attributeType),
			Validators:             validators,
		})
	}

	sort.SliceStable(props, func(i, j int) bool {
		return props[i].Name < props[j].Name
	})

	if schema.MinProperties != nil && schema.MaxProperties != nil {
		min := *schema.MinProperties
		max := *schema.MaxProperties

		if min == 1 && max == 1 {
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
	} else if schema.MinProperties != nil && schema.MaxProperties == nil {
		min := *schema.MinProperties

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
		} else if min > 1 && min == int64(len(props)) {
			for _, prop := range props {
				prop.Required = true
				prop.Optional = false
			}
		}
	}

	return props
}
