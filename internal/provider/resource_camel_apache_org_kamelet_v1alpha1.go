/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type CamelApacheOrgKameletV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*CamelApacheOrgKameletV1Alpha1Resource)(nil)
)

type CamelApacheOrgKameletV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type CamelApacheOrgKameletV1Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Authorization *map[string]string `tfsdk:"authorization" yaml:"authorization,omitempty"`

		Definition *struct {
			Dollarschema *string `tfsdk:"dollarschema" yaml:"$schema,omitempty"`

			Description *string `tfsdk:"description" yaml:"description,omitempty"`

			Example utilities.Dynamic `tfsdk:"example" yaml:"example,omitempty"`

			ExternalDocs *struct {
				Description *string `tfsdk:"description" yaml:"description,omitempty"`

				Url *string `tfsdk:"url" yaml:"url,omitempty"`
			} `tfsdk:"external_docs" yaml:"externalDocs,omitempty"`

			Id *string `tfsdk:"id" yaml:"id,omitempty"`

			Properties *struct {
				Default utilities.Dynamic `tfsdk:"default" yaml:"default,omitempty"`

				Description *string `tfsdk:"description" yaml:"description,omitempty"`

				Enum *[]string `tfsdk:"enum" yaml:"enum,omitempty"`

				Example utilities.Dynamic `tfsdk:"example" yaml:"example,omitempty"`

				ExclusiveMaximum *bool `tfsdk:"exclusive_maximum" yaml:"exclusiveMaximum,omitempty"`

				ExclusiveMinimum *bool `tfsdk:"exclusive_minimum" yaml:"exclusiveMinimum,omitempty"`

				Format *string `tfsdk:"format" yaml:"format,omitempty"`

				Id *string `tfsdk:"id" yaml:"id,omitempty"`

				MaxItems *int64 `tfsdk:"max_items" yaml:"maxItems,omitempty"`

				MaxLength *int64 `tfsdk:"max_length" yaml:"maxLength,omitempty"`

				MaxProperties *int64 `tfsdk:"max_properties" yaml:"maxProperties,omitempty"`

				Maximum *string `tfsdk:"maximum" yaml:"maximum,omitempty"`

				MinItems *int64 `tfsdk:"min_items" yaml:"minItems,omitempty"`

				MinLength *int64 `tfsdk:"min_length" yaml:"minLength,omitempty"`

				MinProperties *int64 `tfsdk:"min_properties" yaml:"minProperties,omitempty"`

				Minimum *string `tfsdk:"minimum" yaml:"minimum,omitempty"`

				MultipleOf *string `tfsdk:"multiple_of" yaml:"multipleOf,omitempty"`

				Nullable *bool `tfsdk:"nullable" yaml:"nullable,omitempty"`

				Pattern *string `tfsdk:"pattern" yaml:"pattern,omitempty"`

				Title *string `tfsdk:"title" yaml:"title,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`

				UniqueItems *bool `tfsdk:"unique_items" yaml:"uniqueItems,omitempty"`

				X_descriptors *[]string `tfsdk:"x_descriptors" yaml:"x-descriptors,omitempty"`
			} `tfsdk:"properties" yaml:"properties,omitempty"`

			Required *[]string `tfsdk:"required" yaml:"required,omitempty"`

			Title *string `tfsdk:"title" yaml:"title,omitempty"`

			Type *string `tfsdk:"type" yaml:"type,omitempty"`
		} `tfsdk:"definition" yaml:"definition,omitempty"`

		Dependencies *[]string `tfsdk:"dependencies" yaml:"dependencies,omitempty"`

		Flow utilities.Dynamic `tfsdk:"flow" yaml:"flow,omitempty"`

		Sources *[]struct {
			Compression *bool `tfsdk:"compression" yaml:"compression,omitempty"`

			Content *string `tfsdk:"content" yaml:"content,omitempty"`

			ContentKey *string `tfsdk:"content_key" yaml:"contentKey,omitempty"`

			ContentRef *string `tfsdk:"content_ref" yaml:"contentRef,omitempty"`

			ContentType *string `tfsdk:"content_type" yaml:"contentType,omitempty"`

			Interceptors *[]string `tfsdk:"interceptors" yaml:"interceptors,omitempty"`

			Language *string `tfsdk:"language" yaml:"language,omitempty"`

			Loader *string `tfsdk:"loader" yaml:"loader,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Path *string `tfsdk:"path" yaml:"path,omitempty"`

			Property_names *[]string `tfsdk:"property_names" yaml:"property-names,omitempty"`

			RawContent *string `tfsdk:"raw_content" yaml:"rawContent,omitempty"`

			Type *string `tfsdk:"type" yaml:"type,omitempty"`
		} `tfsdk:"sources" yaml:"sources,omitempty"`

		Template utilities.Dynamic `tfsdk:"template" yaml:"template,omitempty"`

		Types *struct {
			MediaType *string `tfsdk:"media_type" yaml:"mediaType,omitempty"`

			Schema *struct {
				Dollarschema *string `tfsdk:"dollarschema" yaml:"$schema,omitempty"`

				Description *string `tfsdk:"description" yaml:"description,omitempty"`

				Example utilities.Dynamic `tfsdk:"example" yaml:"example,omitempty"`

				ExternalDocs *struct {
					Description *string `tfsdk:"description" yaml:"description,omitempty"`

					Url *string `tfsdk:"url" yaml:"url,omitempty"`
				} `tfsdk:"external_docs" yaml:"externalDocs,omitempty"`

				Id *string `tfsdk:"id" yaml:"id,omitempty"`

				Properties *struct {
					Default utilities.Dynamic `tfsdk:"default" yaml:"default,omitempty"`

					Description *string `tfsdk:"description" yaml:"description,omitempty"`

					Enum *[]string `tfsdk:"enum" yaml:"enum,omitempty"`

					Example utilities.Dynamic `tfsdk:"example" yaml:"example,omitempty"`

					ExclusiveMaximum *bool `tfsdk:"exclusive_maximum" yaml:"exclusiveMaximum,omitempty"`

					ExclusiveMinimum *bool `tfsdk:"exclusive_minimum" yaml:"exclusiveMinimum,omitempty"`

					Format *string `tfsdk:"format" yaml:"format,omitempty"`

					Id *string `tfsdk:"id" yaml:"id,omitempty"`

					MaxItems *int64 `tfsdk:"max_items" yaml:"maxItems,omitempty"`

					MaxLength *int64 `tfsdk:"max_length" yaml:"maxLength,omitempty"`

					MaxProperties *int64 `tfsdk:"max_properties" yaml:"maxProperties,omitempty"`

					Maximum *string `tfsdk:"maximum" yaml:"maximum,omitempty"`

					MinItems *int64 `tfsdk:"min_items" yaml:"minItems,omitempty"`

					MinLength *int64 `tfsdk:"min_length" yaml:"minLength,omitempty"`

					MinProperties *int64 `tfsdk:"min_properties" yaml:"minProperties,omitempty"`

					Minimum *string `tfsdk:"minimum" yaml:"minimum,omitempty"`

					MultipleOf *string `tfsdk:"multiple_of" yaml:"multipleOf,omitempty"`

					Nullable *bool `tfsdk:"nullable" yaml:"nullable,omitempty"`

					Pattern *string `tfsdk:"pattern" yaml:"pattern,omitempty"`

					Title *string `tfsdk:"title" yaml:"title,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`

					UniqueItems *bool `tfsdk:"unique_items" yaml:"uniqueItems,omitempty"`

					X_descriptors *[]string `tfsdk:"x_descriptors" yaml:"x-descriptors,omitempty"`
				} `tfsdk:"properties" yaml:"properties,omitempty"`

				Required *[]string `tfsdk:"required" yaml:"required,omitempty"`

				Title *string `tfsdk:"title" yaml:"title,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`
			} `tfsdk:"schema" yaml:"schema,omitempty"`
		} `tfsdk:"types" yaml:"types,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewCamelApacheOrgKameletV1Alpha1Resource() resource.Resource {
	return &CamelApacheOrgKameletV1Alpha1Resource{}
}

func (r *CamelApacheOrgKameletV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_camel_apache_org_kamelet_v1alpha1"
}

func (r *CamelApacheOrgKameletV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Kamelet is the Schema for the kamelets API",
		MarkdownDescription: "Kamelet is the Schema for the kamelets API",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "the desired specification",
				MarkdownDescription: "the desired specification",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"authorization": {
						Description:         "Deprecated: unused",
						MarkdownDescription: "Deprecated: unused",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"definition": {
						Description:         "defines the formal configuration of the Kamelet",
						MarkdownDescription: "defines the formal configuration of the Kamelet",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"dollarschema": {
								Description:         "JSONSchemaURL represents a schema url.",
								MarkdownDescription: "JSONSchemaURL represents a schema url.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"description": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"example": {
								Description:         "JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.",
								MarkdownDescription: "JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.",

								Type: utilities.DynamicType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"external_docs": {
								Description:         "ExternalDocumentation allows referencing an external resource for extended documentation.",
								MarkdownDescription: "ExternalDocumentation allows referencing an external resource for extended documentation.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"description": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"url": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"id": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"properties": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"default": {
										Description:         "default is a default value for undefined object fields.",
										MarkdownDescription: "default is a default value for undefined object fields.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"description": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enum": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"example": {
										Description:         "JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.",
										MarkdownDescription: "JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"exclusive_maximum": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"exclusive_minimum": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"format": {
										Description:         "format is an OpenAPI v3 format string. Unknown formats are ignored. The following formats are validated:  - bsonobjectid: a bson object ID, i.e. a 24 characters hex string - uri: an URI as parsed by Golang net/url.ParseRequestURI - email: an email address as parsed by Golang net/mail.ParseAddress - hostname: a valid representation for an Internet host name, as defined by RFC 1034, section 3.1 [RFC1034]. - ipv4: an IPv4 IP as parsed by Golang net.ParseIP - ipv6: an IPv6 IP as parsed by Golang net.ParseIP - cidr: a CIDR as parsed by Golang net.ParseCIDR - mac: a MAC address as parsed by Golang net.ParseMAC - uuid: an UUID that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid3: an UUID3 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?3[0-9a-f]{3}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid4: an UUID4 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?4[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - uuid5: an UUID5 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?5[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - isbn: an ISBN10 or ISBN13 number string like '0321751043' or '978-0321751041' - isbn10: an ISBN10 number string like '0321751043' - isbn13: an ISBN13 number string like '978-0321751041' - creditcard: a credit card number defined by the regex ^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35d{3})d{11})$ with any non digit characters mixed in - ssn: a U.S. social security number following the regex ^d{3}[- ]?d{2}[- ]?d{4}$ - hexcolor: an hexadecimal color code like '#FFFFFF' following the regex ^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$ - rgbcolor: an RGB color code like rgb like 'rgb(255,255,255)' - byte: base64 encoded binary data - password: any kind of string - date: a date string like '2006-01-02' as defined by full-date in RFC3339 - duration: a duration string like '22 ns' as parsed by Golang time.ParseDuration or compatible with Scala duration format - datetime: a date time string like '2014-12-15T19:30:20.000Z' as defined by date-time in RFC3339.",
										MarkdownDescription: "format is an OpenAPI v3 format string. Unknown formats are ignored. The following formats are validated:  - bsonobjectid: a bson object ID, i.e. a 24 characters hex string - uri: an URI as parsed by Golang net/url.ParseRequestURI - email: an email address as parsed by Golang net/mail.ParseAddress - hostname: a valid representation for an Internet host name, as defined by RFC 1034, section 3.1 [RFC1034]. - ipv4: an IPv4 IP as parsed by Golang net.ParseIP - ipv6: an IPv6 IP as parsed by Golang net.ParseIP - cidr: a CIDR as parsed by Golang net.ParseCIDR - mac: a MAC address as parsed by Golang net.ParseMAC - uuid: an UUID that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid3: an UUID3 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?3[0-9a-f]{3}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid4: an UUID4 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?4[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - uuid5: an UUID5 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?5[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - isbn: an ISBN10 or ISBN13 number string like '0321751043' or '978-0321751041' - isbn10: an ISBN10 number string like '0321751043' - isbn13: an ISBN13 number string like '978-0321751041' - creditcard: a credit card number defined by the regex ^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35d{3})d{11})$ with any non digit characters mixed in - ssn: a U.S. social security number following the regex ^d{3}[- ]?d{2}[- ]?d{4}$ - hexcolor: an hexadecimal color code like '#FFFFFF' following the regex ^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$ - rgbcolor: an RGB color code like rgb like 'rgb(255,255,255)' - byte: base64 encoded binary data - password: any kind of string - date: a date string like '2006-01-02' as defined by full-date in RFC3339 - duration: a duration string like '22 ns' as parsed by Golang time.ParseDuration or compatible with Scala duration format - datetime: a date time string like '2014-12-15T19:30:20.000Z' as defined by date-time in RFC3339.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"id": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"max_items": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"max_length": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"max_properties": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"maximum": {
										Description:         "A Number represents a JSON number literal.",
										MarkdownDescription: "A Number represents a JSON number literal.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"min_items": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"min_length": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"min_properties": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"minimum": {
										Description:         "A Number represents a JSON number literal.",
										MarkdownDescription: "A Number represents a JSON number literal.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"multiple_of": {
										Description:         "A Number represents a JSON number literal.",
										MarkdownDescription: "A Number represents a JSON number literal.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"nullable": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"pattern": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"title": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"unique_items": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"x_descriptors": {
										Description:         "XDescriptors is a list of extended properties that trigger a custom behavior in external systems",
										MarkdownDescription: "XDescriptors is a list of extended properties that trigger a custom behavior in external systems",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"required": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"title": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"dependencies": {
						Description:         "Camel dependencies needed by the Kamelet",
						MarkdownDescription: "Camel dependencies needed by the Kamelet",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"flow": {
						Description:         "Deprecated: use Template instead the main source in YAML DSL",
						MarkdownDescription: "Deprecated: use Template instead the main source in YAML DSL",

						Type: utilities.DynamicType{},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"sources": {
						Description:         "sources in any Camel DSL supported",
						MarkdownDescription: "sources in any Camel DSL supported",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"compression": {
								Description:         "if the content is compressed (base64 encrypted)",
								MarkdownDescription: "if the content is compressed (base64 encrypted)",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"content": {
								Description:         "the source code (plain text)",
								MarkdownDescription: "the source code (plain text)",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"content_key": {
								Description:         "the confimap key holding the source content",
								MarkdownDescription: "the confimap key holding the source content",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"content_ref": {
								Description:         "the confimap reference holding the source content",
								MarkdownDescription: "the confimap reference holding the source content",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"content_type": {
								Description:         "the content type (tipically text or binary)",
								MarkdownDescription: "the content type (tipically text or binary)",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"interceptors": {
								Description:         "Interceptors are optional identifiers the org.apache.camel.k.RoutesLoader uses to pre/post process sources",
								MarkdownDescription: "Interceptors are optional identifiers the org.apache.camel.k.RoutesLoader uses to pre/post process sources",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"language": {
								Description:         "specify which is the language (Camel DSL) used to interpret this source code",
								MarkdownDescription: "specify which is the language (Camel DSL) used to interpret this source code",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"loader": {
								Description:         "Loader is an optional id of the org.apache.camel.k.RoutesLoader that will interpret this source at runtime",
								MarkdownDescription: "Loader is an optional id of the org.apache.camel.k.RoutesLoader that will interpret this source at runtime",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "the name of the specification",
								MarkdownDescription: "the name of the specification",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"path": {
								Description:         "the path where the file is stored",
								MarkdownDescription: "the path where the file is stored",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"property_names": {
								Description:         "List of property names defined in the source (e.g. if type is 'template')",
								MarkdownDescription: "List of property names defined in the source (e.g. if type is 'template')",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"raw_content": {
								Description:         "the source code (binary)",
								MarkdownDescription: "the source code (binary)",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									validators.Base64Validator(),
								},
							},

							"type": {
								Description:         "Type defines the kind of source described by this object",
								MarkdownDescription: "Type defines the kind of source described by this object",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"template": {
						Description:         "the main source in YAML DSL",
						MarkdownDescription: "the main source in YAML DSL",

						Type: utilities.DynamicType{},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"types": {
						Description:         "data specification types for the events consumed/produced by the Kamelet",
						MarkdownDescription: "data specification types for the events consumed/produced by the Kamelet",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"media_type": {
								Description:         "media type as expected for HTTP media types (ie, application/json)",
								MarkdownDescription: "media type as expected for HTTP media types (ie, application/json)",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"schema": {
								Description:         "the expected schema for the event",
								MarkdownDescription: "the expected schema for the event",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"dollarschema": {
										Description:         "JSONSchemaURL represents a schema url.",
										MarkdownDescription: "JSONSchemaURL represents a schema url.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"description": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"example": {
										Description:         "JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.",
										MarkdownDescription: "JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"external_docs": {
										Description:         "ExternalDocumentation allows referencing an external resource for extended documentation.",
										MarkdownDescription: "ExternalDocumentation allows referencing an external resource for extended documentation.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"description": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"url": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"id": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"properties": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"default": {
												Description:         "default is a default value for undefined object fields.",
												MarkdownDescription: "default is a default value for undefined object fields.",

												Type: utilities.DynamicType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"description": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"enum": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"example": {
												Description:         "JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.",
												MarkdownDescription: "JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.",

												Type: utilities.DynamicType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"exclusive_maximum": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"exclusive_minimum": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"format": {
												Description:         "format is an OpenAPI v3 format string. Unknown formats are ignored. The following formats are validated:  - bsonobjectid: a bson object ID, i.e. a 24 characters hex string - uri: an URI as parsed by Golang net/url.ParseRequestURI - email: an email address as parsed by Golang net/mail.ParseAddress - hostname: a valid representation for an Internet host name, as defined by RFC 1034, section 3.1 [RFC1034]. - ipv4: an IPv4 IP as parsed by Golang net.ParseIP - ipv6: an IPv6 IP as parsed by Golang net.ParseIP - cidr: a CIDR as parsed by Golang net.ParseCIDR - mac: a MAC address as parsed by Golang net.ParseMAC - uuid: an UUID that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid3: an UUID3 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?3[0-9a-f]{3}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid4: an UUID4 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?4[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - uuid5: an UUID5 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?5[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - isbn: an ISBN10 or ISBN13 number string like '0321751043' or '978-0321751041' - isbn10: an ISBN10 number string like '0321751043' - isbn13: an ISBN13 number string like '978-0321751041' - creditcard: a credit card number defined by the regex ^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35d{3})d{11})$ with any non digit characters mixed in - ssn: a U.S. social security number following the regex ^d{3}[- ]?d{2}[- ]?d{4}$ - hexcolor: an hexadecimal color code like '#FFFFFF' following the regex ^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$ - rgbcolor: an RGB color code like rgb like 'rgb(255,255,255)' - byte: base64 encoded binary data - password: any kind of string - date: a date string like '2006-01-02' as defined by full-date in RFC3339 - duration: a duration string like '22 ns' as parsed by Golang time.ParseDuration or compatible with Scala duration format - datetime: a date time string like '2014-12-15T19:30:20.000Z' as defined by date-time in RFC3339.",
												MarkdownDescription: "format is an OpenAPI v3 format string. Unknown formats are ignored. The following formats are validated:  - bsonobjectid: a bson object ID, i.e. a 24 characters hex string - uri: an URI as parsed by Golang net/url.ParseRequestURI - email: an email address as parsed by Golang net/mail.ParseAddress - hostname: a valid representation for an Internet host name, as defined by RFC 1034, section 3.1 [RFC1034]. - ipv4: an IPv4 IP as parsed by Golang net.ParseIP - ipv6: an IPv6 IP as parsed by Golang net.ParseIP - cidr: a CIDR as parsed by Golang net.ParseCIDR - mac: a MAC address as parsed by Golang net.ParseMAC - uuid: an UUID that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid3: an UUID3 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?3[0-9a-f]{3}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid4: an UUID4 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?4[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - uuid5: an UUID5 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?5[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - isbn: an ISBN10 or ISBN13 number string like '0321751043' or '978-0321751041' - isbn10: an ISBN10 number string like '0321751043' - isbn13: an ISBN13 number string like '978-0321751041' - creditcard: a credit card number defined by the regex ^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35d{3})d{11})$ with any non digit characters mixed in - ssn: a U.S. social security number following the regex ^d{3}[- ]?d{2}[- ]?d{4}$ - hexcolor: an hexadecimal color code like '#FFFFFF' following the regex ^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$ - rgbcolor: an RGB color code like rgb like 'rgb(255,255,255)' - byte: base64 encoded binary data - password: any kind of string - date: a date string like '2006-01-02' as defined by full-date in RFC3339 - duration: a duration string like '22 ns' as parsed by Golang time.ParseDuration or compatible with Scala duration format - datetime: a date time string like '2014-12-15T19:30:20.000Z' as defined by date-time in RFC3339.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"id": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_items": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_length": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_properties": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"maximum": {
												Description:         "A Number represents a JSON number literal.",
												MarkdownDescription: "A Number represents a JSON number literal.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"min_items": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"min_length": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"min_properties": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"minimum": {
												Description:         "A Number represents a JSON number literal.",
												MarkdownDescription: "A Number represents a JSON number literal.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"multiple_of": {
												Description:         "A Number represents a JSON number literal.",
												MarkdownDescription: "A Number represents a JSON number literal.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"nullable": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pattern": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"title": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"type": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"unique_items": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"x_descriptors": {
												Description:         "XDescriptors is a list of extended properties that trigger a custom behavior in external systems",
												MarkdownDescription: "XDescriptors is a list of extended properties that trigger a custom behavior in external systems",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"required": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"title": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *CamelApacheOrgKameletV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_camel_apache_org_kamelet_v1alpha1")

	var state CamelApacheOrgKameletV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CamelApacheOrgKameletV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("camel.apache.org/v1alpha1")
	goModel.Kind = utilities.Ptr("Kamelet")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CamelApacheOrgKameletV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_camel_apache_org_kamelet_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *CamelApacheOrgKameletV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_camel_apache_org_kamelet_v1alpha1")

	var state CamelApacheOrgKameletV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CamelApacheOrgKameletV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("camel.apache.org/v1alpha1")
	goModel.Kind = utilities.Ptr("Kamelet")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CamelApacheOrgKameletV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_camel_apache_org_kamelet_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
