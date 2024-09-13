/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package generator

import (
	"fmt"
	"github.com/pb33f/libopenapi/datamodel/high/base"
	v2high "github.com/pb33f/libopenapi/datamodel/high/v2"
	"k8s.io/utils/strings/slices"
	"sort"
	"strings"
)

func ConvertOpenAPIv2(schemas []v2high.Swagger) []*TemplateData {
	data := make([]*TemplateData, 0)
	for _, schema := range schemas {
		for name, definition := range schema.Definitions.Definitions.FromNewest() {
			if supportedOpenAPIv2Object(name, definition) {
				namespaced := isNamespacedObject(name)
				if templateData := openAPIv2AsTemplateData(definition, namespaced); templateData != nil {
					data = append(data, templateData)
				}
			}
		}
	}
	return data
}

func openAPIv2AsTemplateData(definition *base.SchemaProxy, namespaced bool) *TemplateData {
	var group string
	var version string
	var kind string
	if node, present := definition.Schema().Extensions.Get("x-kubernetes-group-version-kind"); present {
		var gvk []map[string]string
		err := node.Decode(&gvk)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		group = gvk[0]["group"]
		version = gvk[0]["version"]
		kind = gvk[0]["kind"]
	} else {
		return nil
	}
	schema := definition.Schema()
	// remove manually managed or otherwise ignored properties
	schema.Properties.Delete("metadata")
	schema.Properties.Delete("status")
	schema.Properties.Delete("apiVersion")
	schema.Properties.Delete("kind")

	imports := AdditionalImports{}
	typeName := resourceTypeName(group, kind, version)

	return &TemplateData{
		BT:          "`",
		Package:     goPackageName(group, version),
		Group:       group,
		Version:     version,
		Kind:        kind,
		PluralKind:  pluralForm(kind),
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
		Properties:        openAPIv2Properties(schema, &imports, "", typeName),
	}
}

func openAPIv2Properties(schema *base.Schema, imports *AdditionalImports, path string, terraformResourceName string) []*Property {
	props := make([]*Property, 0)

	if schema != nil {
		for name, prop := range schema.Properties.FromNewest() {
			if prop.Schema() != nil {
				propPath := propertyPath(path, name)
				if ignored, ok := ignoredAttributes[terraformResourceName]; ok {
					if slices.Contains(ignored, propPath) {
						continue
					}
				}

				schemaType := prop.Schema().Type[0]

				var nestedProperties []*Property
				if schemaType == "array" && prop.Schema().Items != nil && prop.Schema().Items.IsA() && prop.Schema().Items.A.Schema().Type[0] == "object" {
					nestedProperties = openAPIv2Properties(prop.Schema().Items.A.Schema(), imports, propPath, terraformResourceName)
				} else if schemaType == "object" && prop.Schema().AdditionalProperties != nil && prop.Schema().AdditionalProperties.IsA() && prop.Schema().AdditionalProperties.A.Schema().Type[0] == "object" {
					nestedProperties = openAPIv2Properties(prop.Schema().AdditionalProperties.A.Schema(), imports, propPath, terraformResourceName)
				} else {
					nestedProperties = openAPIv2Properties(prop.Schema(), imports, propPath, terraformResourceName)
				}

				attributeType, valueType, elementType, goType, customType := translateTypeWith(&openapiv2TypeTranslator{property: prop.Schema()}, terraformResourceName, propPath)

				validators := validatorsFor(&openapiv2ValidatorExtractor{
					property: prop.Schema(),
					imports:  imports,
				}, terraformResourceName, propPath, imports)

				if goType == "big.Float" {
					imports.MathBig = true
				}

				props = append(props, &Property{
					BT:                     "`",
					Name:                   name,
					GoName:                 goName(name),
					GoType:                 goType,
					TerraformAttributeName: terraformAttributeName(name, path == ""),
					TerraformAttributeType: attributeType,
					TerraformElementType:   elementType,
					TerraformCustomType:    customType,
					TerraformValueType:     valueType,
					Description:            description(prop.Schema().Description),
					Required:               slices.Contains(schema.Required, name),
					Optional:               !slices.Contains(schema.Required, name),
					Computed:               false,
					Properties:             nestedProperties,
					ValidatorsType:         mapAttributeTypeToValidatorsType(attributeType),
					ValidatorsPackage:      mapAttributeTypeToValidatorsPackage(attributeType),
					Validators:             validators,
				})
			}
		}
	}

	sort.SliceStable(props, func(i, j int) bool {
		return props[i].Name < props[j].Name
	})

	if schema.MinProperties != nil && schema.MaxProperties != nil {
		minProperties := *schema.MinProperties
		maxProperties := schema.MaxProperties

		if minProperties == 1 && *maxProperties == 1 {
			for _, outer := range props {
				var pathExpressions []string
				for _, inner := range props {
					if outer.Name != inner.Name {
						pathExpressions = append(pathExpressions, fmt.Sprintf(`path.MatchRelative().AtParent().AtName("%s")`, inner.TerraformAttributeName))
					}
				}
				validator := fmt.Sprintf(`%s.ExactlyOneOf(%v)`, outer.ValidatorsPackage, strings.Join(pathExpressions, ", "))
				outer.Validators = append(outer.Validators, validator)
				addValidatorImports(outer, imports)
				imports.Path = true
			}
		}
	} else if schema.MinProperties != nil && schema.MaxProperties == nil {
		minProperties := *schema.MinProperties

		if minProperties == 1 {
			for _, outer := range props {
				var pathExpressions []string
				for _, inner := range props {
					if outer.Name != inner.Name {
						pathExpressions = append(pathExpressions, fmt.Sprintf(`path.MatchRelative().AtParent().AtName("%s")`, inner.TerraformAttributeName))
					}
				}
				validator := fmt.Sprintf(`%s.AtLeastOneOf(%v)`, outer.ValidatorsPackage, strings.Join(pathExpressions, ", "))
				outer.Validators = append(outer.Validators, validator)
				addValidatorImports(outer, imports)
				imports.Path = true
			}
		} else if minProperties > 1 && minProperties == int64(len(props)) {
			for _, prop := range props {
				prop.Required = true
				prop.Optional = false
			}
		}
	}

	return props
}
