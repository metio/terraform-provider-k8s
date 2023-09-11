/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package camel_apache_org_v1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"net/http"
)

var (
	_ datasource.DataSource              = &CamelApacheOrgKameletV1DataSource{}
	_ datasource.DataSourceWithConfigure = &CamelApacheOrgKameletV1DataSource{}
)

func NewCamelApacheOrgKameletV1DataSource() datasource.DataSource {
	return &CamelApacheOrgKameletV1DataSource{}
}

type CamelApacheOrgKameletV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type CamelApacheOrgKameletV1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		DataTypes *struct {
			Default *string `tfsdk:"default" json:"default,omitempty"`
			Headers *struct {
				Default     *string `tfsdk:"default" json:"default,omitempty"`
				Description *string `tfsdk:"description" json:"description,omitempty"`
				Required    *bool   `tfsdk:"required" json:"required,omitempty"`
				Title       *string `tfsdk:"title" json:"title,omitempty"`
				Type        *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"headers" json:"headers,omitempty"`
			Types *struct {
				Dependencies *[]string `tfsdk:"dependencies" json:"dependencies,omitempty"`
				Description  *string   `tfsdk:"description" json:"description,omitempty"`
				Format       *string   `tfsdk:"format" json:"format,omitempty"`
				Headers      *struct {
					Default     *string `tfsdk:"default" json:"default,omitempty"`
					Description *string `tfsdk:"description" json:"description,omitempty"`
					Required    *bool   `tfsdk:"required" json:"required,omitempty"`
					Title       *string `tfsdk:"title" json:"title,omitempty"`
					Type        *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"headers" json:"headers,omitempty"`
				MediaType *string `tfsdk:"media_type" json:"mediaType,omitempty"`
				Schema    *struct {
					Dollarschema *string            `tfsdk:"dollarschema" json:"$schema,omitempty"`
					Description  *string            `tfsdk:"description" json:"description,omitempty"`
					Example      *map[string]string `tfsdk:"example" json:"example,omitempty"`
					ExternalDocs *struct {
						Description *string `tfsdk:"description" json:"description,omitempty"`
						Url         *string `tfsdk:"url" json:"url,omitempty"`
					} `tfsdk:"external_docs" json:"externalDocs,omitempty"`
					Id         *string `tfsdk:"id" json:"id,omitempty"`
					Properties *struct {
						Default          *map[string]string `tfsdk:"default" json:"default,omitempty"`
						Deprecated       *bool              `tfsdk:"deprecated" json:"deprecated,omitempty"`
						Description      *string            `tfsdk:"description" json:"description,omitempty"`
						Enum             *[]string          `tfsdk:"enum" json:"enum,omitempty"`
						Example          *map[string]string `tfsdk:"example" json:"example,omitempty"`
						ExclusiveMaximum *bool              `tfsdk:"exclusive_maximum" json:"exclusiveMaximum,omitempty"`
						ExclusiveMinimum *bool              `tfsdk:"exclusive_minimum" json:"exclusiveMinimum,omitempty"`
						Format           *string            `tfsdk:"format" json:"format,omitempty"`
						Id               *string            `tfsdk:"id" json:"id,omitempty"`
						MaxItems         *int64             `tfsdk:"max_items" json:"maxItems,omitempty"`
						MaxLength        *int64             `tfsdk:"max_length" json:"maxLength,omitempty"`
						MaxProperties    *int64             `tfsdk:"max_properties" json:"maxProperties,omitempty"`
						Maximum          *string            `tfsdk:"maximum" json:"maximum,omitempty"`
						MinItems         *int64             `tfsdk:"min_items" json:"minItems,omitempty"`
						MinLength        *int64             `tfsdk:"min_length" json:"minLength,omitempty"`
						MinProperties    *int64             `tfsdk:"min_properties" json:"minProperties,omitempty"`
						Minimum          *string            `tfsdk:"minimum" json:"minimum,omitempty"`
						MultipleOf       *string            `tfsdk:"multiple_of" json:"multipleOf,omitempty"`
						Nullable         *bool              `tfsdk:"nullable" json:"nullable,omitempty"`
						Pattern          *string            `tfsdk:"pattern" json:"pattern,omitempty"`
						Title            *string            `tfsdk:"title" json:"title,omitempty"`
						Type             *string            `tfsdk:"type" json:"type,omitempty"`
						UniqueItems      *bool              `tfsdk:"unique_items" json:"uniqueItems,omitempty"`
						X_descriptors    *[]string          `tfsdk:"x_descriptors" json:"x-descriptors,omitempty"`
					} `tfsdk:"properties" json:"properties,omitempty"`
					Required *[]string `tfsdk:"required" json:"required,omitempty"`
					Title    *string   `tfsdk:"title" json:"title,omitempty"`
					Type     *string   `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"schema" json:"schema,omitempty"`
				Scheme *string `tfsdk:"scheme" json:"scheme,omitempty"`
			} `tfsdk:"types" json:"types,omitempty"`
		} `tfsdk:"data_types" json:"dataTypes,omitempty"`
		Definition *struct {
			Dollarschema *string            `tfsdk:"dollarschema" json:"$schema,omitempty"`
			Description  *string            `tfsdk:"description" json:"description,omitempty"`
			Example      *map[string]string `tfsdk:"example" json:"example,omitempty"`
			ExternalDocs *struct {
				Description *string `tfsdk:"description" json:"description,omitempty"`
				Url         *string `tfsdk:"url" json:"url,omitempty"`
			} `tfsdk:"external_docs" json:"externalDocs,omitempty"`
			Id         *string `tfsdk:"id" json:"id,omitempty"`
			Properties *struct {
				Default          *map[string]string `tfsdk:"default" json:"default,omitempty"`
				Deprecated       *bool              `tfsdk:"deprecated" json:"deprecated,omitempty"`
				Description      *string            `tfsdk:"description" json:"description,omitempty"`
				Enum             *[]string          `tfsdk:"enum" json:"enum,omitempty"`
				Example          *map[string]string `tfsdk:"example" json:"example,omitempty"`
				ExclusiveMaximum *bool              `tfsdk:"exclusive_maximum" json:"exclusiveMaximum,omitempty"`
				ExclusiveMinimum *bool              `tfsdk:"exclusive_minimum" json:"exclusiveMinimum,omitempty"`
				Format           *string            `tfsdk:"format" json:"format,omitempty"`
				Id               *string            `tfsdk:"id" json:"id,omitempty"`
				MaxItems         *int64             `tfsdk:"max_items" json:"maxItems,omitempty"`
				MaxLength        *int64             `tfsdk:"max_length" json:"maxLength,omitempty"`
				MaxProperties    *int64             `tfsdk:"max_properties" json:"maxProperties,omitempty"`
				Maximum          *string            `tfsdk:"maximum" json:"maximum,omitempty"`
				MinItems         *int64             `tfsdk:"min_items" json:"minItems,omitempty"`
				MinLength        *int64             `tfsdk:"min_length" json:"minLength,omitempty"`
				MinProperties    *int64             `tfsdk:"min_properties" json:"minProperties,omitempty"`
				Minimum          *string            `tfsdk:"minimum" json:"minimum,omitempty"`
				MultipleOf       *string            `tfsdk:"multiple_of" json:"multipleOf,omitempty"`
				Nullable         *bool              `tfsdk:"nullable" json:"nullable,omitempty"`
				Pattern          *string            `tfsdk:"pattern" json:"pattern,omitempty"`
				Title            *string            `tfsdk:"title" json:"title,omitempty"`
				Type             *string            `tfsdk:"type" json:"type,omitempty"`
				UniqueItems      *bool              `tfsdk:"unique_items" json:"uniqueItems,omitempty"`
				X_descriptors    *[]string          `tfsdk:"x_descriptors" json:"x-descriptors,omitempty"`
			} `tfsdk:"properties" json:"properties,omitempty"`
			Required *[]string `tfsdk:"required" json:"required,omitempty"`
			Title    *string   `tfsdk:"title" json:"title,omitempty"`
			Type     *string   `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"definition" json:"definition,omitempty"`
		Dependencies *[]string `tfsdk:"dependencies" json:"dependencies,omitempty"`
		Sources      *[]struct {
			Compression    *bool     `tfsdk:"compression" json:"compression,omitempty"`
			Content        *string   `tfsdk:"content" json:"content,omitempty"`
			ContentKey     *string   `tfsdk:"content_key" json:"contentKey,omitempty"`
			ContentRef     *string   `tfsdk:"content_ref" json:"contentRef,omitempty"`
			ContentType    *string   `tfsdk:"content_type" json:"contentType,omitempty"`
			Interceptors   *[]string `tfsdk:"interceptors" json:"interceptors,omitempty"`
			Language       *string   `tfsdk:"language" json:"language,omitempty"`
			Loader         *string   `tfsdk:"loader" json:"loader,omitempty"`
			Name           *string   `tfsdk:"name" json:"name,omitempty"`
			Path           *string   `tfsdk:"path" json:"path,omitempty"`
			Property_names *[]string `tfsdk:"property_names" json:"property-names,omitempty"`
			RawContent     *string   `tfsdk:"raw_content" json:"rawContent,omitempty"`
			Type           *string   `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"sources" json:"sources,omitempty"`
		Template *map[string]string `tfsdk:"template" json:"template,omitempty"`
		Types    *struct {
			MediaType *string `tfsdk:"media_type" json:"mediaType,omitempty"`
			Schema    *struct {
				Dollarschema *string            `tfsdk:"dollarschema" json:"$schema,omitempty"`
				Description  *string            `tfsdk:"description" json:"description,omitempty"`
				Example      *map[string]string `tfsdk:"example" json:"example,omitempty"`
				ExternalDocs *struct {
					Description *string `tfsdk:"description" json:"description,omitempty"`
					Url         *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"external_docs" json:"externalDocs,omitempty"`
				Id         *string `tfsdk:"id" json:"id,omitempty"`
				Properties *struct {
					Default          *map[string]string `tfsdk:"default" json:"default,omitempty"`
					Deprecated       *bool              `tfsdk:"deprecated" json:"deprecated,omitempty"`
					Description      *string            `tfsdk:"description" json:"description,omitempty"`
					Enum             *[]string          `tfsdk:"enum" json:"enum,omitempty"`
					Example          *map[string]string `tfsdk:"example" json:"example,omitempty"`
					ExclusiveMaximum *bool              `tfsdk:"exclusive_maximum" json:"exclusiveMaximum,omitempty"`
					ExclusiveMinimum *bool              `tfsdk:"exclusive_minimum" json:"exclusiveMinimum,omitempty"`
					Format           *string            `tfsdk:"format" json:"format,omitempty"`
					Id               *string            `tfsdk:"id" json:"id,omitempty"`
					MaxItems         *int64             `tfsdk:"max_items" json:"maxItems,omitempty"`
					MaxLength        *int64             `tfsdk:"max_length" json:"maxLength,omitempty"`
					MaxProperties    *int64             `tfsdk:"max_properties" json:"maxProperties,omitempty"`
					Maximum          *string            `tfsdk:"maximum" json:"maximum,omitempty"`
					MinItems         *int64             `tfsdk:"min_items" json:"minItems,omitempty"`
					MinLength        *int64             `tfsdk:"min_length" json:"minLength,omitempty"`
					MinProperties    *int64             `tfsdk:"min_properties" json:"minProperties,omitempty"`
					Minimum          *string            `tfsdk:"minimum" json:"minimum,omitempty"`
					MultipleOf       *string            `tfsdk:"multiple_of" json:"multipleOf,omitempty"`
					Nullable         *bool              `tfsdk:"nullable" json:"nullable,omitempty"`
					Pattern          *string            `tfsdk:"pattern" json:"pattern,omitempty"`
					Title            *string            `tfsdk:"title" json:"title,omitempty"`
					Type             *string            `tfsdk:"type" json:"type,omitempty"`
					UniqueItems      *bool              `tfsdk:"unique_items" json:"uniqueItems,omitempty"`
					X_descriptors    *[]string          `tfsdk:"x_descriptors" json:"x-descriptors,omitempty"`
				} `tfsdk:"properties" json:"properties,omitempty"`
				Required *[]string `tfsdk:"required" json:"required,omitempty"`
				Title    *string   `tfsdk:"title" json:"title,omitempty"`
				Type     *string   `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"schema" json:"schema,omitempty"`
		} `tfsdk:"types" json:"types,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CamelApacheOrgKameletV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_camel_apache_org_kamelet_v1"
}

func (r *CamelApacheOrgKameletV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Kamelet is the Schema for the kamelets API.",
		MarkdownDescription: "Kamelet is the Schema for the kamelets API.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"metadata": schema.SingleNestedAttribute{
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Optional:            false,
				Computed:            false,
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							validators.NameValidator(),
							stringvalidator.LengthAtLeast(1),
						},
					},

					"namespace": schema.StringAttribute{
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							validators.NameValidator(),
							stringvalidator.LengthAtLeast(1),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "the desired specification",
				MarkdownDescription: "the desired specification",
				Attributes: map[string]schema.Attribute{
					"data_types": schema.SingleNestedAttribute{
						Description:         "data specification types for the events consumed/produced by the Kamelet",
						MarkdownDescription: "data specification types for the events consumed/produced by the Kamelet",
						Attributes: map[string]schema.Attribute{
							"default": schema.StringAttribute{
								Description:         "the default data type for this Kamelet",
								MarkdownDescription: "the default data type for this Kamelet",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"headers": schema.SingleNestedAttribute{
								Description:         "one to many header specifications",
								MarkdownDescription: "one to many header specifications",
								Attributes: map[string]schema.Attribute{
									"default": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"description": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"required": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"title": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"type": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"types": schema.SingleNestedAttribute{
								Description:         "one to many data type specifications",
								MarkdownDescription: "one to many data type specifications",
								Attributes: map[string]schema.Attribute{
									"dependencies": schema.ListAttribute{
										Description:         "the list of Camel or Maven dependencies required by the data type",
										MarkdownDescription: "the list of Camel or Maven dependencies required by the data type",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"description": schema.StringAttribute{
										Description:         "optional description",
										MarkdownDescription: "optional description",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"format": schema.StringAttribute{
										Description:         "the data type format name",
										MarkdownDescription: "the data type format name",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"headers": schema.SingleNestedAttribute{
										Description:         "one to many header specifications",
										MarkdownDescription: "one to many header specifications",
										Attributes: map[string]schema.Attribute{
											"default": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"description": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"required": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"title": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"media_type": schema.StringAttribute{
										Description:         "media type as expected for HTTP media types (ie, application/json)",
										MarkdownDescription: "media type as expected for HTTP media types (ie, application/json)",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"schema": schema.SingleNestedAttribute{
										Description:         "the expected schema for the data type",
										MarkdownDescription: "the expected schema for the data type",
										Attributes: map[string]schema.Attribute{
											"dollarschema": schema.StringAttribute{
												Description:         "JSONSchemaURL represents a schema url.",
												MarkdownDescription: "JSONSchemaURL represents a schema url.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"description": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"example": schema.MapAttribute{
												Description:         "JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.",
												MarkdownDescription: "JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"external_docs": schema.SingleNestedAttribute{
												Description:         "ExternalDocumentation allows referencing an external resource for extended documentation.",
												MarkdownDescription: "ExternalDocumentation allows referencing an external resource for extended documentation.",
												Attributes: map[string]schema.Attribute{
													"description": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"url": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"id": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"properties": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"default": schema.MapAttribute{
														Description:         "default is a default value for undefined object fields.",
														MarkdownDescription: "default is a default value for undefined object fields.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"deprecated": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"description": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"enum": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"example": schema.MapAttribute{
														Description:         "JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.",
														MarkdownDescription: "JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"exclusive_maximum": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"exclusive_minimum": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"format": schema.StringAttribute{
														Description:         "format is an OpenAPI v3 format string. Unknown formats are ignored. The following formats are validated:  - bsonobjectid: a bson object ID, i.e. a 24 characters hex string - uri: an URI as parsed by Golang net/url.ParseRequestURI - email: an email address as parsed by Golang net/mail.ParseAddress - hostname: a valid representation for an Internet host name, as defined by RFC 1034, section 3.1 [RFC1034]. - ipv4: an IPv4 IP as parsed by Golang net.ParseIP - ipv6: an IPv6 IP as parsed by Golang net.ParseIP - cidr: a CIDR as parsed by Golang net.ParseCIDR - mac: a MAC address as parsed by Golang net.ParseMAC - uuid: an UUID that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid3: an UUID3 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?3[0-9a-f]{3}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid4: an UUID4 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?4[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - uuid5: an UUID5 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?5[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - isbn: an ISBN10 or ISBN13 number string like '0321751043' or '978-0321751041' - isbn10: an ISBN10 number string like '0321751043' - isbn13: an ISBN13 number string like '978-0321751041' - creditcard: a credit card number defined by the regex ^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35d{3})d{11})$ with any non digit characters mixed in - ssn: a U.S. social security number following the regex ^d{3}[- ]?d{2}[- ]?d{4}$ - hexcolor: an hexadecimal color code like '#FFFFFF' following the regex ^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$ - rgbcolor: an RGB color code like rgb like 'rgb(255,255,255)' - byte: base64 encoded binary data - password: any kind of string - date: a date string like '2006-01-02' as defined by full-date in RFC3339 - duration: a duration string like '22 ns' as parsed by Golang time.ParseDuration or compatible with Scala duration format - datetime: a date time string like '2014-12-15T19:30:20.000Z' as defined by date-time in RFC3339.",
														MarkdownDescription: "format is an OpenAPI v3 format string. Unknown formats are ignored. The following formats are validated:  - bsonobjectid: a bson object ID, i.e. a 24 characters hex string - uri: an URI as parsed by Golang net/url.ParseRequestURI - email: an email address as parsed by Golang net/mail.ParseAddress - hostname: a valid representation for an Internet host name, as defined by RFC 1034, section 3.1 [RFC1034]. - ipv4: an IPv4 IP as parsed by Golang net.ParseIP - ipv6: an IPv6 IP as parsed by Golang net.ParseIP - cidr: a CIDR as parsed by Golang net.ParseCIDR - mac: a MAC address as parsed by Golang net.ParseMAC - uuid: an UUID that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid3: an UUID3 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?3[0-9a-f]{3}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid4: an UUID4 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?4[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - uuid5: an UUID5 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?5[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - isbn: an ISBN10 or ISBN13 number string like '0321751043' or '978-0321751041' - isbn10: an ISBN10 number string like '0321751043' - isbn13: an ISBN13 number string like '978-0321751041' - creditcard: a credit card number defined by the regex ^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35d{3})d{11})$ with any non digit characters mixed in - ssn: a U.S. social security number following the regex ^d{3}[- ]?d{2}[- ]?d{4}$ - hexcolor: an hexadecimal color code like '#FFFFFF' following the regex ^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$ - rgbcolor: an RGB color code like rgb like 'rgb(255,255,255)' - byte: base64 encoded binary data - password: any kind of string - date: a date string like '2006-01-02' as defined by full-date in RFC3339 - duration: a duration string like '22 ns' as parsed by Golang time.ParseDuration or compatible with Scala duration format - datetime: a date time string like '2014-12-15T19:30:20.000Z' as defined by date-time in RFC3339.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"id": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"max_items": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"max_length": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"max_properties": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"maximum": schema.StringAttribute{
														Description:         "A Number represents a JSON number literal.",
														MarkdownDescription: "A Number represents a JSON number literal.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"min_items": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"min_length": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"min_properties": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"minimum": schema.StringAttribute{
														Description:         "A Number represents a JSON number literal.",
														MarkdownDescription: "A Number represents a JSON number literal.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"multiple_of": schema.StringAttribute{
														Description:         "A Number represents a JSON number literal.",
														MarkdownDescription: "A Number represents a JSON number literal.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"nullable": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"pattern": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"title": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"type": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"unique_items": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"x_descriptors": schema.ListAttribute{
														Description:         "XDescriptors is a list of extended properties that trigger a custom behavior in external systems",
														MarkdownDescription: "XDescriptors is a list of extended properties that trigger a custom behavior in external systems",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"required": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"title": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"scheme": schema.StringAttribute{
										Description:         "the data type component scheme",
										MarkdownDescription: "the data type component scheme",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"definition": schema.SingleNestedAttribute{
						Description:         "defines the formal configuration of the Kamelet",
						MarkdownDescription: "defines the formal configuration of the Kamelet",
						Attributes: map[string]schema.Attribute{
							"dollarschema": schema.StringAttribute{
								Description:         "JSONSchemaURL represents a schema url.",
								MarkdownDescription: "JSONSchemaURL represents a schema url.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"description": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"example": schema.MapAttribute{
								Description:         "JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.",
								MarkdownDescription: "JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"external_docs": schema.SingleNestedAttribute{
								Description:         "ExternalDocumentation allows referencing an external resource for extended documentation.",
								MarkdownDescription: "ExternalDocumentation allows referencing an external resource for extended documentation.",
								Attributes: map[string]schema.Attribute{
									"description": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"url": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"id": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"properties": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"default": schema.MapAttribute{
										Description:         "default is a default value for undefined object fields.",
										MarkdownDescription: "default is a default value for undefined object fields.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"deprecated": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"description": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enum": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"example": schema.MapAttribute{
										Description:         "JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.",
										MarkdownDescription: "JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"exclusive_maximum": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"exclusive_minimum": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"format": schema.StringAttribute{
										Description:         "format is an OpenAPI v3 format string. Unknown formats are ignored. The following formats are validated:  - bsonobjectid: a bson object ID, i.e. a 24 characters hex string - uri: an URI as parsed by Golang net/url.ParseRequestURI - email: an email address as parsed by Golang net/mail.ParseAddress - hostname: a valid representation for an Internet host name, as defined by RFC 1034, section 3.1 [RFC1034]. - ipv4: an IPv4 IP as parsed by Golang net.ParseIP - ipv6: an IPv6 IP as parsed by Golang net.ParseIP - cidr: a CIDR as parsed by Golang net.ParseCIDR - mac: a MAC address as parsed by Golang net.ParseMAC - uuid: an UUID that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid3: an UUID3 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?3[0-9a-f]{3}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid4: an UUID4 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?4[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - uuid5: an UUID5 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?5[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - isbn: an ISBN10 or ISBN13 number string like '0321751043' or '978-0321751041' - isbn10: an ISBN10 number string like '0321751043' - isbn13: an ISBN13 number string like '978-0321751041' - creditcard: a credit card number defined by the regex ^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35d{3})d{11})$ with any non digit characters mixed in - ssn: a U.S. social security number following the regex ^d{3}[- ]?d{2}[- ]?d{4}$ - hexcolor: an hexadecimal color code like '#FFFFFF' following the regex ^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$ - rgbcolor: an RGB color code like rgb like 'rgb(255,255,255)' - byte: base64 encoded binary data - password: any kind of string - date: a date string like '2006-01-02' as defined by full-date in RFC3339 - duration: a duration string like '22 ns' as parsed by Golang time.ParseDuration or compatible with Scala duration format - datetime: a date time string like '2014-12-15T19:30:20.000Z' as defined by date-time in RFC3339.",
										MarkdownDescription: "format is an OpenAPI v3 format string. Unknown formats are ignored. The following formats are validated:  - bsonobjectid: a bson object ID, i.e. a 24 characters hex string - uri: an URI as parsed by Golang net/url.ParseRequestURI - email: an email address as parsed by Golang net/mail.ParseAddress - hostname: a valid representation for an Internet host name, as defined by RFC 1034, section 3.1 [RFC1034]. - ipv4: an IPv4 IP as parsed by Golang net.ParseIP - ipv6: an IPv6 IP as parsed by Golang net.ParseIP - cidr: a CIDR as parsed by Golang net.ParseCIDR - mac: a MAC address as parsed by Golang net.ParseMAC - uuid: an UUID that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid3: an UUID3 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?3[0-9a-f]{3}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid4: an UUID4 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?4[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - uuid5: an UUID5 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?5[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - isbn: an ISBN10 or ISBN13 number string like '0321751043' or '978-0321751041' - isbn10: an ISBN10 number string like '0321751043' - isbn13: an ISBN13 number string like '978-0321751041' - creditcard: a credit card number defined by the regex ^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35d{3})d{11})$ with any non digit characters mixed in - ssn: a U.S. social security number following the regex ^d{3}[- ]?d{2}[- ]?d{4}$ - hexcolor: an hexadecimal color code like '#FFFFFF' following the regex ^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$ - rgbcolor: an RGB color code like rgb like 'rgb(255,255,255)' - byte: base64 encoded binary data - password: any kind of string - date: a date string like '2006-01-02' as defined by full-date in RFC3339 - duration: a duration string like '22 ns' as parsed by Golang time.ParseDuration or compatible with Scala duration format - datetime: a date time string like '2014-12-15T19:30:20.000Z' as defined by date-time in RFC3339.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"id": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"max_items": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"max_length": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"max_properties": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"maximum": schema.StringAttribute{
										Description:         "A Number represents a JSON number literal.",
										MarkdownDescription: "A Number represents a JSON number literal.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"min_items": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"min_length": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"min_properties": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"minimum": schema.StringAttribute{
										Description:         "A Number represents a JSON number literal.",
										MarkdownDescription: "A Number represents a JSON number literal.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"multiple_of": schema.StringAttribute{
										Description:         "A Number represents a JSON number literal.",
										MarkdownDescription: "A Number represents a JSON number literal.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"nullable": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"pattern": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"title": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"type": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"unique_items": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"x_descriptors": schema.ListAttribute{
										Description:         "XDescriptors is a list of extended properties that trigger a custom behavior in external systems",
										MarkdownDescription: "XDescriptors is a list of extended properties that trigger a custom behavior in external systems",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"required": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"title": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"type": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"dependencies": schema.ListAttribute{
						Description:         "Camel dependencies needed by the Kamelet",
						MarkdownDescription: "Camel dependencies needed by the Kamelet",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"sources": schema.ListNestedAttribute{
						Description:         "sources in any Camel DSL supported",
						MarkdownDescription: "sources in any Camel DSL supported",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"compression": schema.BoolAttribute{
									Description:         "if the content is compressed (base64 encrypted)",
									MarkdownDescription: "if the content is compressed (base64 encrypted)",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"content": schema.StringAttribute{
									Description:         "the source code (plain text)",
									MarkdownDescription: "the source code (plain text)",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"content_key": schema.StringAttribute{
									Description:         "the confimap key holding the source content",
									MarkdownDescription: "the confimap key holding the source content",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"content_ref": schema.StringAttribute{
									Description:         "the confimap reference holding the source content",
									MarkdownDescription: "the confimap reference holding the source content",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"content_type": schema.StringAttribute{
									Description:         "the content type (tipically text or binary)",
									MarkdownDescription: "the content type (tipically text or binary)",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"interceptors": schema.ListAttribute{
									Description:         "Interceptors are optional identifiers the org.apache.camel.k.RoutesLoader uses to pre/post process sources",
									MarkdownDescription: "Interceptors are optional identifiers the org.apache.camel.k.RoutesLoader uses to pre/post process sources",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"language": schema.StringAttribute{
									Description:         "specify which is the language (Camel DSL) used to interpret this source code",
									MarkdownDescription: "specify which is the language (Camel DSL) used to interpret this source code",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"loader": schema.StringAttribute{
									Description:         "Loader is an optional id of the org.apache.camel.k.RoutesLoader that will interpret this source at runtime",
									MarkdownDescription: "Loader is an optional id of the org.apache.camel.k.RoutesLoader that will interpret this source at runtime",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"name": schema.StringAttribute{
									Description:         "the name of the specification",
									MarkdownDescription: "the name of the specification",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"path": schema.StringAttribute{
									Description:         "the path where the file is stored",
									MarkdownDescription: "the path where the file is stored",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"property_names": schema.ListAttribute{
									Description:         "List of property names defined in the source (e.g. if type is 'template')",
									MarkdownDescription: "List of property names defined in the source (e.g. if type is 'template')",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"raw_content": schema.StringAttribute{
									Description:         "the source code (binary)",
									MarkdownDescription: "the source code (binary)",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"type": schema.StringAttribute{
									Description:         "Type defines the kind of source described by this object",
									MarkdownDescription: "Type defines the kind of source described by this object",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"template": schema.MapAttribute{
						Description:         "the main source in YAML DSL",
						MarkdownDescription: "the main source in YAML DSL",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"types": schema.SingleNestedAttribute{
						Description:         "data specification types for the events consumed/produced by the Kamelet Deprecated: In favor of using DataTypes",
						MarkdownDescription: "data specification types for the events consumed/produced by the Kamelet Deprecated: In favor of using DataTypes",
						Attributes: map[string]schema.Attribute{
							"media_type": schema.StringAttribute{
								Description:         "media type as expected for HTTP media types (ie, application/json)",
								MarkdownDescription: "media type as expected for HTTP media types (ie, application/json)",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"schema": schema.SingleNestedAttribute{
								Description:         "the expected schema for the event",
								MarkdownDescription: "the expected schema for the event",
								Attributes: map[string]schema.Attribute{
									"dollarschema": schema.StringAttribute{
										Description:         "JSONSchemaURL represents a schema url.",
										MarkdownDescription: "JSONSchemaURL represents a schema url.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"description": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"example": schema.MapAttribute{
										Description:         "JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.",
										MarkdownDescription: "JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"external_docs": schema.SingleNestedAttribute{
										Description:         "ExternalDocumentation allows referencing an external resource for extended documentation.",
										MarkdownDescription: "ExternalDocumentation allows referencing an external resource for extended documentation.",
										Attributes: map[string]schema.Attribute{
											"description": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"url": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"id": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"properties": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"default": schema.MapAttribute{
												Description:         "default is a default value for undefined object fields.",
												MarkdownDescription: "default is a default value for undefined object fields.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"deprecated": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"description": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"enum": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"example": schema.MapAttribute{
												Description:         "JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.",
												MarkdownDescription: "JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"exclusive_maximum": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"exclusive_minimum": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"format": schema.StringAttribute{
												Description:         "format is an OpenAPI v3 format string. Unknown formats are ignored. The following formats are validated:  - bsonobjectid: a bson object ID, i.e. a 24 characters hex string - uri: an URI as parsed by Golang net/url.ParseRequestURI - email: an email address as parsed by Golang net/mail.ParseAddress - hostname: a valid representation for an Internet host name, as defined by RFC 1034, section 3.1 [RFC1034]. - ipv4: an IPv4 IP as parsed by Golang net.ParseIP - ipv6: an IPv6 IP as parsed by Golang net.ParseIP - cidr: a CIDR as parsed by Golang net.ParseCIDR - mac: a MAC address as parsed by Golang net.ParseMAC - uuid: an UUID that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid3: an UUID3 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?3[0-9a-f]{3}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid4: an UUID4 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?4[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - uuid5: an UUID5 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?5[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - isbn: an ISBN10 or ISBN13 number string like '0321751043' or '978-0321751041' - isbn10: an ISBN10 number string like '0321751043' - isbn13: an ISBN13 number string like '978-0321751041' - creditcard: a credit card number defined by the regex ^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35d{3})d{11})$ with any non digit characters mixed in - ssn: a U.S. social security number following the regex ^d{3}[- ]?d{2}[- ]?d{4}$ - hexcolor: an hexadecimal color code like '#FFFFFF' following the regex ^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$ - rgbcolor: an RGB color code like rgb like 'rgb(255,255,255)' - byte: base64 encoded binary data - password: any kind of string - date: a date string like '2006-01-02' as defined by full-date in RFC3339 - duration: a duration string like '22 ns' as parsed by Golang time.ParseDuration or compatible with Scala duration format - datetime: a date time string like '2014-12-15T19:30:20.000Z' as defined by date-time in RFC3339.",
												MarkdownDescription: "format is an OpenAPI v3 format string. Unknown formats are ignored. The following formats are validated:  - bsonobjectid: a bson object ID, i.e. a 24 characters hex string - uri: an URI as parsed by Golang net/url.ParseRequestURI - email: an email address as parsed by Golang net/mail.ParseAddress - hostname: a valid representation for an Internet host name, as defined by RFC 1034, section 3.1 [RFC1034]. - ipv4: an IPv4 IP as parsed by Golang net.ParseIP - ipv6: an IPv6 IP as parsed by Golang net.ParseIP - cidr: a CIDR as parsed by Golang net.ParseCIDR - mac: a MAC address as parsed by Golang net.ParseMAC - uuid: an UUID that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid3: an UUID3 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?3[0-9a-f]{3}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid4: an UUID4 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?4[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - uuid5: an UUID5 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?5[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - isbn: an ISBN10 or ISBN13 number string like '0321751043' or '978-0321751041' - isbn10: an ISBN10 number string like '0321751043' - isbn13: an ISBN13 number string like '978-0321751041' - creditcard: a credit card number defined by the regex ^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35d{3})d{11})$ with any non digit characters mixed in - ssn: a U.S. social security number following the regex ^d{3}[- ]?d{2}[- ]?d{4}$ - hexcolor: an hexadecimal color code like '#FFFFFF' following the regex ^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$ - rgbcolor: an RGB color code like rgb like 'rgb(255,255,255)' - byte: base64 encoded binary data - password: any kind of string - date: a date string like '2006-01-02' as defined by full-date in RFC3339 - duration: a duration string like '22 ns' as parsed by Golang time.ParseDuration or compatible with Scala duration format - datetime: a date time string like '2014-12-15T19:30:20.000Z' as defined by date-time in RFC3339.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"id": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"max_items": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"max_length": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"max_properties": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"maximum": schema.StringAttribute{
												Description:         "A Number represents a JSON number literal.",
												MarkdownDescription: "A Number represents a JSON number literal.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"min_items": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"min_length": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"min_properties": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"minimum": schema.StringAttribute{
												Description:         "A Number represents a JSON number literal.",
												MarkdownDescription: "A Number represents a JSON number literal.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"multiple_of": schema.StringAttribute{
												Description:         "A Number represents a JSON number literal.",
												MarkdownDescription: "A Number represents a JSON number literal.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"nullable": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"pattern": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"title": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"unique_items": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"x_descriptors": schema.ListAttribute{
												Description:         "XDescriptors is a list of extended properties that trigger a custom behavior in external systems",
												MarkdownDescription: "XDescriptors is a list of extended properties that trigger a custom behavior in external systems",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"required": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"title": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"type": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *CamelApacheOrgKameletV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *CamelApacheOrgKameletV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_camel_apache_org_kamelet_v1")

	var data CamelApacheOrgKameletV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "camel.apache.org", Version: "v1", Resource: "kamelets"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		var statusError *k8sErrors.StatusError
		if errors.As(err, &statusError) {
			if statusError.Status().Code == http.StatusNotFound {
				response.Diagnostics.AddError(
					"Unable to find resource",
					fmt.Sprintf("The requested resource cannot be found. "+
						"Make sure that it does exist in your cluster and you have set the correct name and namespace configured.\n\n"+
						"Namespace: %s\n"+
						"Name: %s", data.Metadata.Namespace, data.Metadata.Name),
				)
				return
			}
		} else {
			response.Diagnostics.AddError(
				"Unable to GET resource",
				fmt.Sprintf("An unexpected error occurred while reading the resource. "+
					"Please report this issue to the provider developers.\n\n"+
					"GET Error (%T): %s", err, err.Error()),
			)
		}
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse CamelApacheOrgKameletV1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("camel.apache.org/v1")
	data.Kind = pointer.String("Kamelet")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
