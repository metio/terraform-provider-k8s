/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package apiextensions_crossplane_io_v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &ApiextensionsCrossplaneIoCompositeResourceDefinitionV1DataSource{}
	_ datasource.DataSourceWithConfigure = &ApiextensionsCrossplaneIoCompositeResourceDefinitionV1DataSource{}
)

func NewApiextensionsCrossplaneIoCompositeResourceDefinitionV1DataSource() datasource.DataSource {
	return &ApiextensionsCrossplaneIoCompositeResourceDefinitionV1DataSource{}
}

type ApiextensionsCrossplaneIoCompositeResourceDefinitionV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type ApiextensionsCrossplaneIoCompositeResourceDefinitionV1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ClaimNames *struct {
			Categories *[]string `tfsdk:"categories" json:"categories,omitempty"`
			Kind       *string   `tfsdk:"kind" json:"kind,omitempty"`
			ListKind   *string   `tfsdk:"list_kind" json:"listKind,omitempty"`
			Plural     *string   `tfsdk:"plural" json:"plural,omitempty"`
			ShortNames *[]string `tfsdk:"short_names" json:"shortNames,omitempty"`
			Singular   *string   `tfsdk:"singular" json:"singular,omitempty"`
		} `tfsdk:"claim_names" json:"claimNames,omitempty"`
		ConnectionSecretKeys *[]string `tfsdk:"connection_secret_keys" json:"connectionSecretKeys,omitempty"`
		Conversion           *struct {
			Strategy *string `tfsdk:"strategy" json:"strategy,omitempty"`
			Webhook  *struct {
				ClientConfig *struct {
					CaBundle *string `tfsdk:"ca_bundle" json:"caBundle,omitempty"`
					Service  *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						Path      *string `tfsdk:"path" json:"path,omitempty"`
						Port      *int64  `tfsdk:"port" json:"port,omitempty"`
					} `tfsdk:"service" json:"service,omitempty"`
					Url *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"client_config" json:"clientConfig,omitempty"`
				ConversionReviewVersions *[]string `tfsdk:"conversion_review_versions" json:"conversionReviewVersions,omitempty"`
			} `tfsdk:"webhook" json:"webhook,omitempty"`
		} `tfsdk:"conversion" json:"conversion,omitempty"`
		DefaultCompositeDeletePolicy *string `tfsdk:"default_composite_delete_policy" json:"defaultCompositeDeletePolicy,omitempty"`
		DefaultCompositionRef        *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"default_composition_ref" json:"defaultCompositionRef,omitempty"`
		DefaultCompositionUpdatePolicy *string `tfsdk:"default_composition_update_policy" json:"defaultCompositionUpdatePolicy,omitempty"`
		EnforcedCompositionRef         *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"enforced_composition_ref" json:"enforcedCompositionRef,omitempty"`
		Group    *string `tfsdk:"group" json:"group,omitempty"`
		Metadata *struct {
			Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		} `tfsdk:"metadata" json:"metadata,omitempty"`
		Names *struct {
			Categories *[]string `tfsdk:"categories" json:"categories,omitempty"`
			Kind       *string   `tfsdk:"kind" json:"kind,omitempty"`
			ListKind   *string   `tfsdk:"list_kind" json:"listKind,omitempty"`
			Plural     *string   `tfsdk:"plural" json:"plural,omitempty"`
			ShortNames *[]string `tfsdk:"short_names" json:"shortNames,omitempty"`
			Singular   *string   `tfsdk:"singular" json:"singular,omitempty"`
		} `tfsdk:"names" json:"names,omitempty"`
		Versions *[]struct {
			AdditionalPrinterColumns *[]struct {
				Description *string `tfsdk:"description" json:"description,omitempty"`
				Format      *string `tfsdk:"format" json:"format,omitempty"`
				JsonPath    *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Name        *string `tfsdk:"name" json:"name,omitempty"`
				Priority    *int64  `tfsdk:"priority" json:"priority,omitempty"`
				Type        *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"additional_printer_columns" json:"additionalPrinterColumns,omitempty"`
			Deprecated         *bool   `tfsdk:"deprecated" json:"deprecated,omitempty"`
			DeprecationWarning *string `tfsdk:"deprecation_warning" json:"deprecationWarning,omitempty"`
			Name               *string `tfsdk:"name" json:"name,omitempty"`
			Referenceable      *bool   `tfsdk:"referenceable" json:"referenceable,omitempty"`
			Schema             *struct {
				OpenAPIV3Schema *map[string]string `tfsdk:"open_apiv3_schema" json:"openAPIV3Schema,omitempty"`
			} `tfsdk:"schema" json:"schema,omitempty"`
			Served *bool `tfsdk:"served" json:"served,omitempty"`
		} `tfsdk:"versions" json:"versions,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ApiextensionsCrossplaneIoCompositeResourceDefinitionV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apiextensions_crossplane_io_composite_resource_definition_v1"
}

func (r *ApiextensionsCrossplaneIoCompositeResourceDefinitionV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "A CompositeResourceDefinition defines a new kind of composite infrastructure resource. The new resource is composed of other composite or managed infrastructure resources.",
		MarkdownDescription: "A CompositeResourceDefinition defines a new kind of composite infrastructure resource. The new resource is composed of other composite or managed infrastructure resources.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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
				Description:         "CompositeResourceDefinitionSpec specifies the desired state of the definition.",
				MarkdownDescription: "CompositeResourceDefinitionSpec specifies the desired state of the definition.",
				Attributes: map[string]schema.Attribute{
					"claim_names": schema.SingleNestedAttribute{
						Description:         "ClaimNames specifies the names of an optional composite resource claim. When claim names are specified Crossplane will create a namespaced 'composite resource claim' CRD that corresponds to the defined composite resource. This composite resource claim acts as a namespaced proxy for the composite resource; creating, updating, or deleting the claim will create, update, or delete a corresponding composite resource. You may add claim names to an existing CompositeResourceDefinition, but they cannot be changed or removed once they have been set.",
						MarkdownDescription: "ClaimNames specifies the names of an optional composite resource claim. When claim names are specified Crossplane will create a namespaced 'composite resource claim' CRD that corresponds to the defined composite resource. This composite resource claim acts as a namespaced proxy for the composite resource; creating, updating, or deleting the claim will create, update, or delete a corresponding composite resource. You may add claim names to an existing CompositeResourceDefinition, but they cannot be changed or removed once they have been set.",
						Attributes: map[string]schema.Attribute{
							"categories": schema.ListAttribute{
								Description:         "categories is a list of grouped resources this custom resource belongs to (e.g. 'all'). This is published in API discovery documents, and used by clients to support invocations like 'kubectl get all'.",
								MarkdownDescription: "categories is a list of grouped resources this custom resource belongs to (e.g. 'all'). This is published in API discovery documents, and used by clients to support invocations like 'kubectl get all'.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"kind": schema.StringAttribute{
								Description:         "kind is the serialized kind of the resource. It is normally CamelCase and singular. Custom resource instances will use this value as the 'kind' attribute in API calls.",
								MarkdownDescription: "kind is the serialized kind of the resource. It is normally CamelCase and singular. Custom resource instances will use this value as the 'kind' attribute in API calls.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"list_kind": schema.StringAttribute{
								Description:         "listKind is the serialized kind of the list for this resource. Defaults to ''kind'List'.",
								MarkdownDescription: "listKind is the serialized kind of the list for this resource. Defaults to ''kind'List'.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"plural": schema.StringAttribute{
								Description:         "plural is the plural name of the resource to serve. The custom resources are served under '/apis/<group>/<version>/.../<plural>'. Must match the name of the CustomResourceDefinition (in the form '<names.plural>.<group>'). Must be all lowercase.",
								MarkdownDescription: "plural is the plural name of the resource to serve. The custom resources are served under '/apis/<group>/<version>/.../<plural>'. Must match the name of the CustomResourceDefinition (in the form '<names.plural>.<group>'). Must be all lowercase.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"short_names": schema.ListAttribute{
								Description:         "shortNames are short names for the resource, exposed in API discovery documents, and used by clients to support invocations like 'kubectl get <shortname>'. It must be all lowercase.",
								MarkdownDescription: "shortNames are short names for the resource, exposed in API discovery documents, and used by clients to support invocations like 'kubectl get <shortname>'. It must be all lowercase.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"singular": schema.StringAttribute{
								Description:         "singular is the singular name of the resource. It must be all lowercase. Defaults to lowercased 'kind'.",
								MarkdownDescription: "singular is the singular name of the resource. It must be all lowercase. Defaults to lowercased 'kind'.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"connection_secret_keys": schema.ListAttribute{
						Description:         "ConnectionSecretKeys is the list of keys that will be exposed to the end user of the defined kind. If the list is empty, all keys will be published.",
						MarkdownDescription: "ConnectionSecretKeys is the list of keys that will be exposed to the end user of the defined kind. If the list is empty, all keys will be published.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"conversion": schema.SingleNestedAttribute{
						Description:         "Conversion defines all conversion settings for the defined Composite resource.",
						MarkdownDescription: "Conversion defines all conversion settings for the defined Composite resource.",
						Attributes: map[string]schema.Attribute{
							"strategy": schema.StringAttribute{
								Description:         "strategy specifies how custom resources are converted between versions. Allowed values are: - ''None'': The converter only change the apiVersion and would not touch any other field in the custom resource. - ''Webhook'': API Server will call to an external webhook to do the conversion. Additional information is needed for this option. This requires spec.preserveUnknownFields to be false, and spec.conversion.webhook to be set.",
								MarkdownDescription: "strategy specifies how custom resources are converted between versions. Allowed values are: - ''None'': The converter only change the apiVersion and would not touch any other field in the custom resource. - ''Webhook'': API Server will call to an external webhook to do the conversion. Additional information is needed for this option. This requires spec.preserveUnknownFields to be false, and spec.conversion.webhook to be set.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"webhook": schema.SingleNestedAttribute{
								Description:         "webhook describes how to call the conversion webhook. Required when 'strategy' is set to ''Webhook''.",
								MarkdownDescription: "webhook describes how to call the conversion webhook. Required when 'strategy' is set to ''Webhook''.",
								Attributes: map[string]schema.Attribute{
									"client_config": schema.SingleNestedAttribute{
										Description:         "clientConfig is the instructions for how to call the webhook if strategy is 'Webhook'.",
										MarkdownDescription: "clientConfig is the instructions for how to call the webhook if strategy is 'Webhook'.",
										Attributes: map[string]schema.Attribute{
											"ca_bundle": schema.StringAttribute{
												Description:         "caBundle is a PEM encoded CA bundle which will be used to validate the webhook's server certificate. If unspecified, system trust roots on the apiserver are used.",
												MarkdownDescription: "caBundle is a PEM encoded CA bundle which will be used to validate the webhook's server certificate. If unspecified, system trust roots on the apiserver are used.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"service": schema.SingleNestedAttribute{
												Description:         "service is a reference to the service for this webhook. Either service or url must be specified.  If the webhook is running within the cluster, then you should use 'service'.",
												MarkdownDescription: "service is a reference to the service for this webhook. Either service or url must be specified.  If the webhook is running within the cluster, then you should use 'service'.",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "name is the name of the service. Required",
														MarkdownDescription: "name is the name of the service. Required",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"namespace": schema.StringAttribute{
														Description:         "namespace is the namespace of the service. Required",
														MarkdownDescription: "namespace is the namespace of the service. Required",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"path": schema.StringAttribute{
														Description:         "path is an optional URL path at which the webhook will be contacted.",
														MarkdownDescription: "path is an optional URL path at which the webhook will be contacted.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"port": schema.Int64Attribute{
														Description:         "port is an optional service port at which the webhook will be contacted. 'port' should be a valid port number (1-65535, inclusive). Defaults to 443 for backward compatibility.",
														MarkdownDescription: "port is an optional service port at which the webhook will be contacted. 'port' should be a valid port number (1-65535, inclusive). Defaults to 443 for backward compatibility.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"url": schema.StringAttribute{
												Description:         "url gives the location of the webhook, in standard URL form ('scheme://host:port/path'). Exactly one of 'url' or 'service' must be specified.  The 'host' should not refer to a service running in the cluster; use the 'service' field instead. The host might be resolved via external DNS in some apiservers (e.g., 'kube-apiserver' cannot resolve in-cluster DNS as that would be a layering violation). 'host' may also be an IP address.  Please note that using 'localhost' or '127.0.0.1' as a 'host' is risky unless you take great care to run this webhook on all hosts which run an apiserver which might need to make calls to this webhook. Such installs are likely to be non-portable, i.e., not easy to turn up in a new cluster.  The scheme must be 'https'; the URL must begin with 'https://'.  A path is optional, and if present may be any string permissible in a URL. You may use the path to pass an arbitrary string to the webhook, for example, a cluster identifier.  Attempting to use a user or basic auth e.g. 'user:password@' is not allowed. Fragments ('#...') and query parameters ('?...') are not allowed, either.",
												MarkdownDescription: "url gives the location of the webhook, in standard URL form ('scheme://host:port/path'). Exactly one of 'url' or 'service' must be specified.  The 'host' should not refer to a service running in the cluster; use the 'service' field instead. The host might be resolved via external DNS in some apiservers (e.g., 'kube-apiserver' cannot resolve in-cluster DNS as that would be a layering violation). 'host' may also be an IP address.  Please note that using 'localhost' or '127.0.0.1' as a 'host' is risky unless you take great care to run this webhook on all hosts which run an apiserver which might need to make calls to this webhook. Such installs are likely to be non-portable, i.e., not easy to turn up in a new cluster.  The scheme must be 'https'; the URL must begin with 'https://'.  A path is optional, and if present may be any string permissible in a URL. You may use the path to pass an arbitrary string to the webhook, for example, a cluster identifier.  Attempting to use a user or basic auth e.g. 'user:password@' is not allowed. Fragments ('#...') and query parameters ('?...') are not allowed, either.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"conversion_review_versions": schema.ListAttribute{
										Description:         "conversionReviewVersions is an ordered list of preferred 'ConversionReview' versions the Webhook expects. The API server will use the first version in the list which it supports. If none of the versions specified in this list are supported by API server, conversion will fail for the custom resource. If a persisted Webhook configuration specifies allowed versions and does not include any versions known to the API Server, calls to the webhook will fail.",
										MarkdownDescription: "conversionReviewVersions is an ordered list of preferred 'ConversionReview' versions the Webhook expects. The API server will use the first version in the list which it supports. If none of the versions specified in this list are supported by API server, conversion will fail for the custom resource. If a persisted Webhook configuration specifies allowed versions and does not include any versions known to the API Server, calls to the webhook will fail.",
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"default_composite_delete_policy": schema.StringAttribute{
						Description:         "DefaultCompositeDeletePolicy is the policy used when deleting the Composite that is associated with the Claim if no policy has been specified.",
						MarkdownDescription: "DefaultCompositeDeletePolicy is the policy used when deleting the Composite that is associated with the Claim if no policy has been specified.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"default_composition_ref": schema.SingleNestedAttribute{
						Description:         "DefaultCompositionRef refers to the Composition resource that will be used in case no composition selector is given.",
						MarkdownDescription: "DefaultCompositionRef refers to the Composition resource that will be used in case no composition selector is given.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the Composition.",
								MarkdownDescription: "Name of the Composition.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"default_composition_update_policy": schema.StringAttribute{
						Description:         "DefaultCompositionUpdatePolicy is the policy used when updating composites after a new Composition Revision has been created if no policy has been specified on the composite.",
						MarkdownDescription: "DefaultCompositionUpdatePolicy is the policy used when updating composites after a new Composition Revision has been created if no policy has been specified on the composite.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"enforced_composition_ref": schema.SingleNestedAttribute{
						Description:         "EnforcedCompositionRef refers to the Composition resource that will be used by all composite instances whose schema is defined by this definition.",
						MarkdownDescription: "EnforcedCompositionRef refers to the Composition resource that will be used by all composite instances whose schema is defined by this definition.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the Composition.",
								MarkdownDescription: "Name of the Composition.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"group": schema.StringAttribute{
						Description:         "Group specifies the API group of the defined composite resource. Composite resources are served under '/apis/<group>/...'. Must match the name of the XRD (in the form '<names.plural>.<group>').",
						MarkdownDescription: "Group specifies the API group of the defined composite resource. Composite resources are served under '/apis/<group>/...'. Must match the name of the XRD (in the form '<names.plural>.<group>').",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"metadata": schema.SingleNestedAttribute{
						Description:         "Metadata specifies the desired metadata for the defined composite resource and claim CRD's.",
						MarkdownDescription: "Metadata specifies the desired metadata for the defined composite resource and claim CRD's.",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
								MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"labels": schema.MapAttribute{
								Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels and services. These labels are added to the composite resource and claim CRD's in addition to any labels defined by 'CompositionResourceDefinition' 'metadata.labels'.",
								MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels and services. These labels are added to the composite resource and claim CRD's in addition to any labels defined by 'CompositionResourceDefinition' 'metadata.labels'.",
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

					"names": schema.SingleNestedAttribute{
						Description:         "Names specifies the resource and kind names of the defined composite resource.",
						MarkdownDescription: "Names specifies the resource and kind names of the defined composite resource.",
						Attributes: map[string]schema.Attribute{
							"categories": schema.ListAttribute{
								Description:         "categories is a list of grouped resources this custom resource belongs to (e.g. 'all'). This is published in API discovery documents, and used by clients to support invocations like 'kubectl get all'.",
								MarkdownDescription: "categories is a list of grouped resources this custom resource belongs to (e.g. 'all'). This is published in API discovery documents, and used by clients to support invocations like 'kubectl get all'.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"kind": schema.StringAttribute{
								Description:         "kind is the serialized kind of the resource. It is normally CamelCase and singular. Custom resource instances will use this value as the 'kind' attribute in API calls.",
								MarkdownDescription: "kind is the serialized kind of the resource. It is normally CamelCase and singular. Custom resource instances will use this value as the 'kind' attribute in API calls.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"list_kind": schema.StringAttribute{
								Description:         "listKind is the serialized kind of the list for this resource. Defaults to ''kind'List'.",
								MarkdownDescription: "listKind is the serialized kind of the list for this resource. Defaults to ''kind'List'.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"plural": schema.StringAttribute{
								Description:         "plural is the plural name of the resource to serve. The custom resources are served under '/apis/<group>/<version>/.../<plural>'. Must match the name of the CustomResourceDefinition (in the form '<names.plural>.<group>'). Must be all lowercase.",
								MarkdownDescription: "plural is the plural name of the resource to serve. The custom resources are served under '/apis/<group>/<version>/.../<plural>'. Must match the name of the CustomResourceDefinition (in the form '<names.plural>.<group>'). Must be all lowercase.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"short_names": schema.ListAttribute{
								Description:         "shortNames are short names for the resource, exposed in API discovery documents, and used by clients to support invocations like 'kubectl get <shortname>'. It must be all lowercase.",
								MarkdownDescription: "shortNames are short names for the resource, exposed in API discovery documents, and used by clients to support invocations like 'kubectl get <shortname>'. It must be all lowercase.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"singular": schema.StringAttribute{
								Description:         "singular is the singular name of the resource. It must be all lowercase. Defaults to lowercased 'kind'.",
								MarkdownDescription: "singular is the singular name of the resource. It must be all lowercase. Defaults to lowercased 'kind'.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"versions": schema.ListNestedAttribute{
						Description:         "Versions is the list of all API versions of the defined composite resource. Version names are used to compute the order in which served versions are listed in API discovery. If the version string is 'kube-like', it will sort above non 'kube-like' version strings, which are ordered lexicographically. 'Kube-like' versions start with a 'v', then are followed by a number (the major version), then optionally the string 'alpha' or 'beta' and another number (the minor version). These are sorted first by GA > beta > alpha (where GA is a version with no suffix such as beta or alpha), and then by comparing major version, then minor version. An example sorted list of versions: v10, v2, v1, v11beta2, v10beta3, v3beta1, v12alpha1, v11alpha2, foo1, foo10.",
						MarkdownDescription: "Versions is the list of all API versions of the defined composite resource. Version names are used to compute the order in which served versions are listed in API discovery. If the version string is 'kube-like', it will sort above non 'kube-like' version strings, which are ordered lexicographically. 'Kube-like' versions start with a 'v', then are followed by a number (the major version), then optionally the string 'alpha' or 'beta' and another number (the minor version). These are sorted first by GA > beta > alpha (where GA is a version with no suffix such as beta or alpha), and then by comparing major version, then minor version. An example sorted list of versions: v10, v2, v1, v11beta2, v10beta3, v3beta1, v12alpha1, v11alpha2, foo1, foo10.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"additional_printer_columns": schema.ListNestedAttribute{
									Description:         "AdditionalPrinterColumns specifies additional columns returned in Table output. If no columns are specified, a single column displaying the age of the custom resource is used. See the following link for details: https://kubernetes.io/docs/reference/using-api/api-concepts/#receiving-resources-as-tables",
									MarkdownDescription: "AdditionalPrinterColumns specifies additional columns returned in Table output. If no columns are specified, a single column displaying the age of the custom resource is used. See the following link for details: https://kubernetes.io/docs/reference/using-api/api-concepts/#receiving-resources-as-tables",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"description": schema.StringAttribute{
												Description:         "description is a human readable description of this column.",
												MarkdownDescription: "description is a human readable description of this column.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"format": schema.StringAttribute{
												Description:         "format is an optional OpenAPI type definition for this column. The 'name' format is applied to the primary identifier column to assist in clients identifying column is the resource name. See https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#data-types for details.",
												MarkdownDescription: "format is an optional OpenAPI type definition for this column. The 'name' format is applied to the primary identifier column to assist in clients identifying column is the resource name. See https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#data-types for details.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"json_path": schema.StringAttribute{
												Description:         "jsonPath is a simple JSON path (i.e. with array notation) which is evaluated against each custom resource to produce the value for this column.",
												MarkdownDescription: "jsonPath is a simple JSON path (i.e. with array notation) which is evaluated against each custom resource to produce the value for this column.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "name is a human readable name for the column.",
												MarkdownDescription: "name is a human readable name for the column.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"priority": schema.Int64Attribute{
												Description:         "priority is an integer defining the relative importance of this column compared to others. Lower numbers are considered higher priority. Columns that may be omitted in limited space scenarios should be given a priority greater than 0.",
												MarkdownDescription: "priority is an integer defining the relative importance of this column compared to others. Lower numbers are considered higher priority. Columns that may be omitted in limited space scenarios should be given a priority greater than 0.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"type": schema.StringAttribute{
												Description:         "type is an OpenAPI type definition for this column. See https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#data-types for details.",
												MarkdownDescription: "type is an OpenAPI type definition for this column. See https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#data-types for details.",
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

								"deprecated": schema.BoolAttribute{
									Description:         "The deprecated field specifies that this version is deprecated and should not be used.",
									MarkdownDescription: "The deprecated field specifies that this version is deprecated and should not be used.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"deprecation_warning": schema.StringAttribute{
									Description:         "DeprecationWarning specifies the message that should be shown to the user when using this version.",
									MarkdownDescription: "DeprecationWarning specifies the message that should be shown to the user when using this version.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"name": schema.StringAttribute{
									Description:         "Name of this version, e.g. “v1”, “v2beta1”, etc. Composite resources are served under this version at '/apis/<group>/<version>/...' if 'served' is true.",
									MarkdownDescription: "Name of this version, e.g. “v1”, “v2beta1”, etc. Composite resources are served under this version at '/apis/<group>/<version>/...' if 'served' is true.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"referenceable": schema.BoolAttribute{
									Description:         "Referenceable specifies that this version may be referenced by a Composition in order to configure which resources an XR may be composed of. Exactly one version must be marked as referenceable; all Compositions must target only the referenceable version. The referenceable version must be served. It's mapped to the CRD's 'spec.versions[*].storage' field.",
									MarkdownDescription: "Referenceable specifies that this version may be referenced by a Composition in order to configure which resources an XR may be composed of. Exactly one version must be marked as referenceable; all Compositions must target only the referenceable version. The referenceable version must be served. It's mapped to the CRD's 'spec.versions[*].storage' field.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"schema": schema.SingleNestedAttribute{
									Description:         "Schema describes the schema used for validation, pruning, and defaulting of this version of the defined composite resource. Fields required by all composite resources will be injected into this schema automatically, and will override equivalently named fields in this schema. Omitting this schema results in a schema that contains only the fields required by all composite resources.",
									MarkdownDescription: "Schema describes the schema used for validation, pruning, and defaulting of this version of the defined composite resource. Fields required by all composite resources will be injected into this schema automatically, and will override equivalently named fields in this schema. Omitting this schema results in a schema that contains only the fields required by all composite resources.",
									Attributes: map[string]schema.Attribute{
										"open_apiv3_schema": schema.MapAttribute{
											Description:         "OpenAPIV3Schema is the OpenAPI v3 schema to use for validation and pruning.",
											MarkdownDescription: "OpenAPIV3Schema is the OpenAPI v3 schema to use for validation and pruning.",
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

								"served": schema.BoolAttribute{
									Description:         "Served specifies that this version should be served via REST APIs.",
									MarkdownDescription: "Served specifies that this version should be served via REST APIs.",
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
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *ApiextensionsCrossplaneIoCompositeResourceDefinitionV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *ApiextensionsCrossplaneIoCompositeResourceDefinitionV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_apiextensions_crossplane_io_composite_resource_definition_v1")

	var data ApiextensionsCrossplaneIoCompositeResourceDefinitionV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "apiextensions.crossplane.io", Version: "v1", Resource: "CompositeResourceDefinition"}).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
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

	var readResponse ApiextensionsCrossplaneIoCompositeResourceDefinitionV1DataSourceData
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

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("apiextensions.crossplane.io/v1")
	data.Kind = pointer.String("CompositeResourceDefinition")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
