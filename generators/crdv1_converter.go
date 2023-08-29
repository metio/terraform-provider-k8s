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

func convertCRDv1(crds []*apiextensionsv1.CustomResourceDefinition, pkg string) []*TemplateData {
	data := make([]*TemplateData, 0)
	for _, crd := range crds {
		if templateData := crdv1AsTemplateData(crd, pkg); templateData != nil {
			data = append(data, templateData)
		}
	}
	return data
}

func crdv1AsTemplateData(crd *apiextensionsv1.CustomResourceDefinition, pkg string) *TemplateData {
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
	trn := terraformResourceName(group, kind, version.Name)

	return &TemplateData{
		BT:                    "`",
		Package:               pkg,
		File:                  terraformResourceFile(group, kind, version.Name),
		Group:                 group,
		Version:               version.Name,
		Kind:                  kind,
		Namespaced:            crd.Spec.Scope == apiextensionsv1.NamespaceScoped,
		Description:           description(schema.Description),
		TerraformResourceType: terraformResourceType(group, kind, version.Name),
		TerraformModelType:    terraformModelType(group, kind, version.Name),
		GoModelType:           goModelType(group, kind, version.Name),
		Properties:            crdv1Properties(schema, &imports, "", trn),
		TerraformResourceName: trn,
		AdditionalImports:     imports,
	}
}

func crdv1Properties(schema *apiextensionsv1.JSONSchemaProps, imports *AdditionalImports, path string, terraformResourceName string) []*Property {
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
			nestedProperties = crdv1Properties(prop.Items.Schema, imports, propPath, terraformResourceName)
		} else if prop.Type == "object" && prop.AdditionalProperties != nil && prop.AdditionalProperties.Schema.Type == "object" {
			nestedProperties = crdv1Properties(prop.AdditionalProperties.Schema, imports, propPath, terraformResourceName)
		} else {
			nestedProperties = crdv1Properties(&prop, imports, propPath, terraformResourceName)
		}

		attributeType, valueType, goType := translateTypeWith(&crd1TypeTranslator{property: &prop})

		validators := validatorsFor(&crdv1ValidatorExtractor{
			property: &prop,
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
			Description:            description(prop.Description),
			Required:               slices.Contains(schema.Required, name),
			Optional:               !slices.Contains(schema.Required, name),
			Computed:               false,
			Properties:             nestedProperties,
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
