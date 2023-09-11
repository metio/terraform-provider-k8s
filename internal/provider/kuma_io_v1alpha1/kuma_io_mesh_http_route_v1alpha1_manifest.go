/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kuma_io_v1alpha1

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &KumaIoMeshHttprouteV1Alpha1Manifest{}
)

func NewKumaIoMeshHttprouteV1Alpha1Manifest() datasource.DataSource {
	return &KumaIoMeshHttprouteV1Alpha1Manifest{}
}

type KumaIoMeshHttprouteV1Alpha1Manifest struct{}

type KumaIoMeshHttprouteV1Alpha1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		TargetRef *struct {
			Kind *string            `tfsdk:"kind" json:"kind,omitempty"`
			Mesh *string            `tfsdk:"mesh" json:"mesh,omitempty"`
			Name *string            `tfsdk:"name" json:"name,omitempty"`
			Tags *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
		} `tfsdk:"target_ref" json:"targetRef,omitempty"`
		To *[]struct {
			Rules *[]struct {
				Default *struct {
					BackendRefs *[]struct {
						Kind   *string            `tfsdk:"kind" json:"kind,omitempty"`
						Mesh   *string            `tfsdk:"mesh" json:"mesh,omitempty"`
						Name   *string            `tfsdk:"name" json:"name,omitempty"`
						Tags   *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
						Weight *int64             `tfsdk:"weight" json:"weight,omitempty"`
					} `tfsdk:"backend_refs" json:"backendRefs,omitempty"`
					Filters *[]struct {
						RequestHeaderModifier *struct {
							Add *[]struct {
								Name  *string `tfsdk:"name" json:"name,omitempty"`
								Value *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"add" json:"add,omitempty"`
							Remove *[]string `tfsdk:"remove" json:"remove,omitempty"`
							Set    *[]struct {
								Name  *string `tfsdk:"name" json:"name,omitempty"`
								Value *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"set" json:"set,omitempty"`
						} `tfsdk:"request_header_modifier" json:"requestHeaderModifier,omitempty"`
						RequestMirror *struct {
							BackendRef *struct {
								Kind *string            `tfsdk:"kind" json:"kind,omitempty"`
								Mesh *string            `tfsdk:"mesh" json:"mesh,omitempty"`
								Name *string            `tfsdk:"name" json:"name,omitempty"`
								Tags *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
							} `tfsdk:"backend_ref" json:"backendRef,omitempty"`
							Percentage *string `tfsdk:"percentage" json:"percentage,omitempty"`
						} `tfsdk:"request_mirror" json:"requestMirror,omitempty"`
						RequestRedirect *struct {
							Hostname *string `tfsdk:"hostname" json:"hostname,omitempty"`
							Path     *struct {
								ReplaceFullPath    *string `tfsdk:"replace_full_path" json:"replaceFullPath,omitempty"`
								ReplacePrefixMatch *string `tfsdk:"replace_prefix_match" json:"replacePrefixMatch,omitempty"`
								Type               *string `tfsdk:"type" json:"type,omitempty"`
							} `tfsdk:"path" json:"path,omitempty"`
							Port       *int64  `tfsdk:"port" json:"port,omitempty"`
							Scheme     *string `tfsdk:"scheme" json:"scheme,omitempty"`
							StatusCode *int64  `tfsdk:"status_code" json:"statusCode,omitempty"`
						} `tfsdk:"request_redirect" json:"requestRedirect,omitempty"`
						ResponseHeaderModifier *struct {
							Add *[]struct {
								Name  *string `tfsdk:"name" json:"name,omitempty"`
								Value *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"add" json:"add,omitempty"`
							Remove *[]string `tfsdk:"remove" json:"remove,omitempty"`
							Set    *[]struct {
								Name  *string `tfsdk:"name" json:"name,omitempty"`
								Value *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"set" json:"set,omitempty"`
						} `tfsdk:"response_header_modifier" json:"responseHeaderModifier,omitempty"`
						Type       *string `tfsdk:"type" json:"type,omitempty"`
						UrlRewrite *struct {
							Hostname *string `tfsdk:"hostname" json:"hostname,omitempty"`
							Path     *struct {
								ReplaceFullPath    *string `tfsdk:"replace_full_path" json:"replaceFullPath,omitempty"`
								ReplacePrefixMatch *string `tfsdk:"replace_prefix_match" json:"replacePrefixMatch,omitempty"`
								Type               *string `tfsdk:"type" json:"type,omitempty"`
							} `tfsdk:"path" json:"path,omitempty"`
						} `tfsdk:"url_rewrite" json:"urlRewrite,omitempty"`
					} `tfsdk:"filters" json:"filters,omitempty"`
				} `tfsdk:"default" json:"default,omitempty"`
				Matches *[]struct {
					Headers *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Type  *string `tfsdk:"type" json:"type,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"headers" json:"headers,omitempty"`
					Method *string `tfsdk:"method" json:"method,omitempty"`
					Path   *struct {
						Type  *string `tfsdk:"type" json:"type,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"path" json:"path,omitempty"`
					QueryParams *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Type  *string `tfsdk:"type" json:"type,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"query_params" json:"queryParams,omitempty"`
				} `tfsdk:"matches" json:"matches,omitempty"`
			} `tfsdk:"rules" json:"rules,omitempty"`
			TargetRef *struct {
				Kind *string            `tfsdk:"kind" json:"kind,omitempty"`
				Mesh *string            `tfsdk:"mesh" json:"mesh,omitempty"`
				Name *string            `tfsdk:"name" json:"name,omitempty"`
				Tags *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
			} `tfsdk:"target_ref" json:"targetRef,omitempty"`
		} `tfsdk:"to" json:"to,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KumaIoMeshHttprouteV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kuma_io_mesh_http_route_v1alpha1_manifest"
}

func (r *KumaIoMeshHttprouteV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"yaml": schema.StringAttribute{
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
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
						Optional:            true,
						Computed:            false,
						Validators: []validator.Map{
							validators.LabelValidator(),
						},
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Map{
							validators.AnnotationValidator(),
						},
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "Spec is the specification of the Kuma MeshHTTPRoute resource.",
				MarkdownDescription: "Spec is the specification of the Kuma MeshHTTPRoute resource.",
				Attributes: map[string]schema.Attribute{
					"target_ref": schema.SingleNestedAttribute{
						Description:         "TargetRef is a reference to the resource the policy takes an effect on. The resource could be either a real store object or virtual resource defined inplace.",
						MarkdownDescription: "TargetRef is a reference to the resource the policy takes an effect on. The resource could be either a real store object or virtual resource defined inplace.",
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Description:         "Kind of the referenced resource",
								MarkdownDescription: "Kind of the referenced resource",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Mesh", "MeshSubset", "MeshGateway", "MeshService", "MeshServiceSubset", "MeshHTTPRoute"),
								},
							},

							"mesh": schema.StringAttribute{
								Description:         "Mesh is reserved for future use to identify cross mesh resources.",
								MarkdownDescription: "Mesh is reserved for future use to identify cross mesh resources.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
								MarkdownDescription: "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tags": schema.MapAttribute{
								Description:         "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
								MarkdownDescription: "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"to": schema.ListNestedAttribute{
						Description:         "To matches destination services of requests and holds configuration.",
						MarkdownDescription: "To matches destination services of requests and holds configuration.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"rules": schema.ListNestedAttribute{
									Description:         "Rules contains the routing rules applies to a combination of top-level targetRef and the targetRef in this entry.",
									MarkdownDescription: "Rules contains the routing rules applies to a combination of top-level targetRef and the targetRef in this entry.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"default": schema.SingleNestedAttribute{
												Description:         "Default holds routing rules that can be merged with rules from other policies.",
												MarkdownDescription: "Default holds routing rules that can be merged with rules from other policies.",
												Attributes: map[string]schema.Attribute{
													"backend_refs": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"kind": schema.StringAttribute{
																	Description:         "Kind of the referenced resource",
																	MarkdownDescription: "Kind of the referenced resource",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("Mesh", "MeshSubset", "MeshGateway", "MeshService", "MeshServiceSubset", "MeshHTTPRoute"),
																	},
																},

																"mesh": schema.StringAttribute{
																	Description:         "Mesh is reserved for future use to identify cross mesh resources.",
																	MarkdownDescription: "Mesh is reserved for future use to identify cross mesh resources.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
																	MarkdownDescription: "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"tags": schema.MapAttribute{
																	Description:         "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
																	MarkdownDescription: "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"weight": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(0),
																	},
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"filters": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"request_header_modifier": schema.SingleNestedAttribute{
																	Description:         "Only one action is supported per header name. Configuration to set or add multiple values for a header must use RFC 7230 header value formatting, separating each value with a comma.",
																	MarkdownDescription: "Only one action is supported per header name. Configuration to set or add multiple values for a header must use RFC 7230 header value formatting, separating each value with a comma.",
																	Attributes: map[string]schema.Attribute{
																		"add": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																						Validators: []validator.String{
																							stringvalidator.LengthAtLeast(1),
																							stringvalidator.LengthAtMost(256),
																							stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
																						},
																					},

																					"value": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"remove": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"set": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																						Validators: []validator.String{
																							stringvalidator.LengthAtLeast(1),
																							stringvalidator.LengthAtMost(256),
																							stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
																						},
																					},

																					"value": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"request_mirror": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"backend_ref": schema.SingleNestedAttribute{
																			Description:         "TargetRef defines structure that allows attaching policy to various objects",
																			MarkdownDescription: "TargetRef defines structure that allows attaching policy to various objects",
																			Attributes: map[string]schema.Attribute{
																				"kind": schema.StringAttribute{
																					Description:         "Kind of the referenced resource",
																					MarkdownDescription: "Kind of the referenced resource",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																					Validators: []validator.String{
																						stringvalidator.OneOf("Mesh", "MeshSubset", "MeshGateway", "MeshService", "MeshServiceSubset", "MeshHTTPRoute"),
																					},
																				},

																				"mesh": schema.StringAttribute{
																					Description:         "Mesh is reserved for future use to identify cross mesh resources.",
																					MarkdownDescription: "Mesh is reserved for future use to identify cross mesh resources.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"name": schema.StringAttribute{
																					Description:         "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
																					MarkdownDescription: "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"tags": schema.MapAttribute{
																					Description:         "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
																					MarkdownDescription: "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: true,
																			Optional: false,
																			Computed: false,
																		},

																		"percentage": schema.StringAttribute{
																			Description:         "Percentage of requests to mirror. If not specified, all requests to the target cluster will be mirrored.",
																			MarkdownDescription: "Percentage of requests to mirror. If not specified, all requests to the target cluster will be mirrored.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"request_redirect": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"hostname": schema.StringAttribute{
																			Description:         "PreciseHostname is the fully qualified domain name of a network host. This matches the RFC 1123 definition of a hostname with 1 notable exception that numeric IP addresses are not allowed.  Note that as per RFC1035 and RFC1123, a *label* must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character. No other punctuation is allowed.",
																			MarkdownDescription: "PreciseHostname is the fully qualified domain name of a network host. This matches the RFC 1123 definition of a hostname with 1 notable exception that numeric IP addresses are not allowed.  Note that as per RFC1035 and RFC1123, a *label* must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character. No other punctuation is allowed.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.LengthAtLeast(1),
																				stringvalidator.LengthAtMost(253),
																				stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
																			},
																		},

																		"path": schema.SingleNestedAttribute{
																			Description:         "Path defines parameters used to modify the path of the incoming request. The modified path is then used to construct the location header. When empty, the request path is used as-is.",
																			MarkdownDescription: "Path defines parameters used to modify the path of the incoming request. The modified path is then used to construct the location header. When empty, the request path is used as-is.",
																			Attributes: map[string]schema.Attribute{
																				"replace_full_path": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"replace_prefix_match": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"type": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																					Validators: []validator.String{
																						stringvalidator.OneOf("ReplaceFullPath", "ReplacePrefixMatch"),
																					},
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"port": schema.Int64Attribute{
																			Description:         "Port is the port to be used in the value of the 'Location' header in the response. When empty, port (if specified) of the request is used.",
																			MarkdownDescription: "Port is the port to be used in the value of the 'Location' header in the response. When empty, port (if specified) of the request is used.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.Int64{
																				int64validator.AtLeast(1),
																				int64validator.AtMost(65535),
																			},
																		},

																		"scheme": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.OneOf("http", "https"),
																			},
																		},

																		"status_code": schema.Int64Attribute{
																			Description:         "StatusCode is the HTTP status code to be used in response.",
																			MarkdownDescription: "StatusCode is the HTTP status code to be used in response.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.Int64{
																				int64validator.OneOf(301, 302, 303, 307, 308),
																			},
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"response_header_modifier": schema.SingleNestedAttribute{
																	Description:         "Only one action is supported per header name. Configuration to set or add multiple values for a header must use RFC 7230 header value formatting, separating each value with a comma.",
																	MarkdownDescription: "Only one action is supported per header name. Configuration to set or add multiple values for a header must use RFC 7230 header value formatting, separating each value with a comma.",
																	Attributes: map[string]schema.Attribute{
																		"add": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																						Validators: []validator.String{
																							stringvalidator.LengthAtLeast(1),
																							stringvalidator.LengthAtMost(256),
																							stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
																						},
																					},

																					"value": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"remove": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"set": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																						Validators: []validator.String{
																							stringvalidator.LengthAtLeast(1),
																							stringvalidator.LengthAtMost(256),
																							stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
																						},
																					},

																					"value": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"type": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("RequestHeaderModifier", "ResponseHeaderModifier", "RequestRedirect", "URLRewrite", "RequestMirror"),
																	},
																},

																"url_rewrite": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"hostname": schema.StringAttribute{
																			Description:         "Hostname is the value to be used to replace the host header value during forwarding.",
																			MarkdownDescription: "Hostname is the value to be used to replace the host header value during forwarding.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.LengthAtLeast(1),
																				stringvalidator.LengthAtMost(253),
																				stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
																			},
																		},

																		"path": schema.SingleNestedAttribute{
																			Description:         "Path defines a path rewrite.",
																			MarkdownDescription: "Path defines a path rewrite.",
																			Attributes: map[string]schema.Attribute{
																				"replace_full_path": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"replace_prefix_match": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"type": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																					Validators: []validator.String{
																						stringvalidator.OneOf("ReplaceFullPath", "ReplacePrefixMatch"),
																					},
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},
												},
												Required: true,
												Optional: false,
												Computed: false,
											},

											"matches": schema.ListNestedAttribute{
												Description:         "Matches describes how to match HTTP requests this rule should be applied to.",
												MarkdownDescription: "Matches describes how to match HTTP requests this rule should be applied to.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"headers": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "Name is the name of the HTTP Header to be matched. Name MUST be lower case as they will be handled with case insensitivity (See https://tools.ietf.org/html/rfc7230#section-3.2).",
																		MarkdownDescription: "Name is the name of the HTTP Header to be matched. Name MUST be lower case as they will be handled with case insensitivity (See https://tools.ietf.org/html/rfc7230#section-3.2).",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.LengthAtLeast(1),
																			stringvalidator.LengthAtMost(256),
																			stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
																		},
																	},

																	"type": schema.StringAttribute{
																		Description:         "Type specifies how to match against the value of the header.",
																		MarkdownDescription: "Type specifies how to match against the value of the header.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("Exact", "Present", "RegularExpression", "Absent", "Prefix"),
																		},
																	},

																	"value": schema.StringAttribute{
																		Description:         "Value is the value of HTTP Header to be matched.",
																		MarkdownDescription: "Value is the value of HTTP Header to be matched.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"method": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("CONNECT", "DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT", "TRACE"),
															},
														},

														"path": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"type": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("Exact", "PathPrefix", "RegularExpression"),
																	},
																},

																"value": schema.StringAttribute{
																	Description:         "Exact or prefix matches must be an absolute path. A prefix matches only if separated by a slash or the entire path.",
																	MarkdownDescription: "Exact or prefix matches must be an absolute path. A prefix matches only if separated by a slash or the entire path.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																	},
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"query_params": schema.ListNestedAttribute{
															Description:         "QueryParams matches based on HTTP URL query parameters. Multiple matches are ANDed together such that all listed matches must succeed.",
															MarkdownDescription: "QueryParams matches based on HTTP URL query parameters. Multiple matches are ANDed together such that all listed matches must succeed.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.LengthAtLeast(1),
																		},
																	},

																	"type": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("Exact", "RegularExpression"),
																		},
																	},

																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},
													},
												},
												Required: true,
												Optional: false,
												Computed: false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"target_ref": schema.SingleNestedAttribute{
									Description:         "TargetRef is a reference to the resource that represents a group of request destinations.",
									MarkdownDescription: "TargetRef is a reference to the resource that represents a group of request destinations.",
									Attributes: map[string]schema.Attribute{
										"kind": schema.StringAttribute{
											Description:         "Kind of the referenced resource",
											MarkdownDescription: "Kind of the referenced resource",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Mesh", "MeshSubset", "MeshGateway", "MeshService", "MeshServiceSubset", "MeshHTTPRoute"),
											},
										},

										"mesh": schema.StringAttribute{
											Description:         "Mesh is reserved for future use to identify cross mesh resources.",
											MarkdownDescription: "Mesh is reserved for future use to identify cross mesh resources.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
											MarkdownDescription: "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tags": schema.MapAttribute{
											Description:         "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
											MarkdownDescription: "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *KumaIoMeshHttprouteV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kuma_io_mesh_http_route_v1alpha1_manifest")

	var model KumaIoMeshHttprouteV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("kuma.io/v1alpha1")
	model.Kind = pointer.String("MeshHTTPRoute")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
