/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kuma_io_v1alpha1

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
	_ datasource.DataSource              = &KumaIoMeshHttprouteV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &KumaIoMeshHttprouteV1Alpha1DataSource{}
)

func NewKumaIoMeshHttprouteV1Alpha1DataSource() datasource.DataSource {
	return &KumaIoMeshHttprouteV1Alpha1DataSource{}
}

type KumaIoMeshHttprouteV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type KumaIoMeshHttprouteV1Alpha1DataSourceData struct {
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

func (r *KumaIoMeshHttprouteV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kuma_io_mesh_http_route_v1alpha1"
}

func (r *KumaIoMeshHttprouteV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
								Optional:            false,
								Computed:            true,
							},

							"mesh": schema.StringAttribute{
								Description:         "Mesh is reserved for future use to identify cross mesh resources.",
								MarkdownDescription: "Mesh is reserved for future use to identify cross mesh resources.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
								MarkdownDescription: "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"tags": schema.MapAttribute{
								Description:         "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
								MarkdownDescription: "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
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
																	Optional:            false,
																	Computed:            true,
																},

																"mesh": schema.StringAttribute{
																	Description:         "Mesh is reserved for future use to identify cross mesh resources.",
																	MarkdownDescription: "Mesh is reserved for future use to identify cross mesh resources.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
																	MarkdownDescription: "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"tags": schema.MapAttribute{
																	Description:         "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
																	MarkdownDescription: "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"weight": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"value": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

																		"remove": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"set": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"value": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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
																					Optional:            false,
																					Computed:            true,
																				},

																				"mesh": schema.StringAttribute{
																					Description:         "Mesh is reserved for future use to identify cross mesh resources.",
																					MarkdownDescription: "Mesh is reserved for future use to identify cross mesh resources.",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"name": schema.StringAttribute{
																					Description:         "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
																					MarkdownDescription: "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"tags": schema.MapAttribute{
																					Description:         "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
																					MarkdownDescription: "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
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

																		"percentage": schema.StringAttribute{
																			Description:         "Percentage of requests to mirror. If not specified, all requests to the target cluster will be mirrored.",
																			MarkdownDescription: "Percentage of requests to mirror. If not specified, all requests to the target cluster will be mirrored.",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"request_redirect": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"hostname": schema.StringAttribute{
																			Description:         "PreciseHostname is the fully qualified domain name of a network host. This matches the RFC 1123 definition of a hostname with 1 notable exception that numeric IP addresses are not allowed.  Note that as per RFC1035 and RFC1123, a *label* must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character. No other punctuation is allowed.",
																			MarkdownDescription: "PreciseHostname is the fully qualified domain name of a network host. This matches the RFC 1123 definition of a hostname with 1 notable exception that numeric IP addresses are not allowed.  Note that as per RFC1035 and RFC1123, a *label* must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character. No other punctuation is allowed.",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"path": schema.SingleNestedAttribute{
																			Description:         "Path defines parameters used to modify the path of the incoming request. The modified path is then used to construct the location header. When empty, the request path is used as-is.",
																			MarkdownDescription: "Path defines parameters used to modify the path of the incoming request. The modified path is then used to construct the location header. When empty, the request path is used as-is.",
																			Attributes: map[string]schema.Attribute{
																				"replace_full_path": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"replace_prefix_match": schema.StringAttribute{
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

																		"port": schema.Int64Attribute{
																			Description:         "Port is the port to be used in the value of the 'Location' header in the response. When empty, port (if specified) of the request is used.",
																			MarkdownDescription: "Port is the port to be used in the value of the 'Location' header in the response. When empty, port (if specified) of the request is used.",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"scheme": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"status_code": schema.Int64Attribute{
																			Description:         "StatusCode is the HTTP status code to be used in response.",
																			MarkdownDescription: "StatusCode is the HTTP status code to be used in response.",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"value": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

																		"remove": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"set": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"value": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

																"type": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"url_rewrite": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"hostname": schema.StringAttribute{
																			Description:         "Hostname is the value to be used to replace the host header value during forwarding.",
																			MarkdownDescription: "Hostname is the value to be used to replace the host header value during forwarding.",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"path": schema.SingleNestedAttribute{
																			Description:         "Path defines a path rewrite.",
																			MarkdownDescription: "Path defines a path rewrite.",
																			Attributes: map[string]schema.Attribute{
																				"replace_full_path": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"replace_prefix_match": schema.StringAttribute{
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"type": schema.StringAttribute{
																		Description:         "Type specifies how to match against the value of the header.",
																		MarkdownDescription: "Type specifies how to match against the value of the header.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"value": schema.StringAttribute{
																		Description:         "Value is the value of HTTP Header to be matched.",
																		MarkdownDescription: "Value is the value of HTTP Header to be matched.",
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

														"method": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"path": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"type": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"value": schema.StringAttribute{
																	Description:         "Exact or prefix matches must be an absolute path. A prefix matches only if separated by a slash or the entire path.",
																	MarkdownDescription: "Exact or prefix matches must be an absolute path. A prefix matches only if separated by a slash or the entire path.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"query_params": schema.ListNestedAttribute{
															Description:         "QueryParams matches based on HTTP URL query parameters. Multiple matches are ANDed together such that all listed matches must succeed.",
															MarkdownDescription: "QueryParams matches based on HTTP URL query parameters. Multiple matches are ANDed together such that all listed matches must succeed.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
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

																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"target_ref": schema.SingleNestedAttribute{
									Description:         "TargetRef is a reference to the resource that represents a group of request destinations.",
									MarkdownDescription: "TargetRef is a reference to the resource that represents a group of request destinations.",
									Attributes: map[string]schema.Attribute{
										"kind": schema.StringAttribute{
											Description:         "Kind of the referenced resource",
											MarkdownDescription: "Kind of the referenced resource",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"mesh": schema.StringAttribute{
											Description:         "Mesh is reserved for future use to identify cross mesh resources.",
											MarkdownDescription: "Mesh is reserved for future use to identify cross mesh resources.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
											MarkdownDescription: "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"tags": schema.MapAttribute{
											Description:         "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
											MarkdownDescription: "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
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

func (r *KumaIoMeshHttprouteV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *KumaIoMeshHttprouteV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_kuma_io_mesh_http_route_v1alpha1")

	var data KumaIoMeshHttprouteV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "kuma.io", Version: "v1alpha1", Resource: "meshhttproutes"}).
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

	var readResponse KumaIoMeshHttprouteV1Alpha1DataSourceData
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
	data.ApiVersion = pointer.String("kuma.io/v1alpha1")
	data.Kind = pointer.String("MeshHTTPRoute")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
