/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kuma_io_v1alpha1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &KumaIoMeshAccessLogV1Alpha1Manifest{}
)

func NewKumaIoMeshAccessLogV1Alpha1Manifest() datasource.DataSource {
	return &KumaIoMeshAccessLogV1Alpha1Manifest{}
}

type KumaIoMeshAccessLogV1Alpha1Manifest struct{}

type KumaIoMeshAccessLogV1Alpha1ManifestData struct {
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
		From *[]struct {
			Default *struct {
				Backends *[]struct {
					File *struct {
						Format *struct {
							Json *[]struct {
								Key   *string `tfsdk:"key" json:"key,omitempty"`
								Value *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"json" json:"json,omitempty"`
							OmitEmptyValues *bool   `tfsdk:"omit_empty_values" json:"omitEmptyValues,omitempty"`
							Plain           *string `tfsdk:"plain" json:"plain,omitempty"`
							Type            *string `tfsdk:"type" json:"type,omitempty"`
						} `tfsdk:"format" json:"format,omitempty"`
						Path *string `tfsdk:"path" json:"path,omitempty"`
					} `tfsdk:"file" json:"file,omitempty"`
					OpenTelemetry *struct {
						Attributes *[]struct {
							Key   *string `tfsdk:"key" json:"key,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"attributes" json:"attributes,omitempty"`
						Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
						Endpoint *string            `tfsdk:"endpoint" json:"endpoint,omitempty"`
					} `tfsdk:"open_telemetry" json:"openTelemetry,omitempty"`
					Tcp *struct {
						Address *string `tfsdk:"address" json:"address,omitempty"`
						Format  *struct {
							Json *[]struct {
								Key   *string `tfsdk:"key" json:"key,omitempty"`
								Value *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"json" json:"json,omitempty"`
							OmitEmptyValues *bool   `tfsdk:"omit_empty_values" json:"omitEmptyValues,omitempty"`
							Plain           *string `tfsdk:"plain" json:"plain,omitempty"`
							Type            *string `tfsdk:"type" json:"type,omitempty"`
						} `tfsdk:"format" json:"format,omitempty"`
					} `tfsdk:"tcp" json:"tcp,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"backends" json:"backends,omitempty"`
			} `tfsdk:"default" json:"default,omitempty"`
			TargetRef *struct {
				Kind       *string            `tfsdk:"kind" json:"kind,omitempty"`
				Mesh       *string            `tfsdk:"mesh" json:"mesh,omitempty"`
				Name       *string            `tfsdk:"name" json:"name,omitempty"`
				ProxyTypes *[]string          `tfsdk:"proxy_types" json:"proxyTypes,omitempty"`
				Tags       *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
			} `tfsdk:"target_ref" json:"targetRef,omitempty"`
		} `tfsdk:"from" json:"from,omitempty"`
		TargetRef *struct {
			Kind       *string            `tfsdk:"kind" json:"kind,omitempty"`
			Mesh       *string            `tfsdk:"mesh" json:"mesh,omitempty"`
			Name       *string            `tfsdk:"name" json:"name,omitempty"`
			ProxyTypes *[]string          `tfsdk:"proxy_types" json:"proxyTypes,omitempty"`
			Tags       *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
		} `tfsdk:"target_ref" json:"targetRef,omitempty"`
		To *[]struct {
			Default *struct {
				Backends *[]struct {
					File *struct {
						Format *struct {
							Json *[]struct {
								Key   *string `tfsdk:"key" json:"key,omitempty"`
								Value *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"json" json:"json,omitempty"`
							OmitEmptyValues *bool   `tfsdk:"omit_empty_values" json:"omitEmptyValues,omitempty"`
							Plain           *string `tfsdk:"plain" json:"plain,omitempty"`
							Type            *string `tfsdk:"type" json:"type,omitempty"`
						} `tfsdk:"format" json:"format,omitempty"`
						Path *string `tfsdk:"path" json:"path,omitempty"`
					} `tfsdk:"file" json:"file,omitempty"`
					OpenTelemetry *struct {
						Attributes *[]struct {
							Key   *string `tfsdk:"key" json:"key,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"attributes" json:"attributes,omitempty"`
						Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
						Endpoint *string            `tfsdk:"endpoint" json:"endpoint,omitempty"`
					} `tfsdk:"open_telemetry" json:"openTelemetry,omitempty"`
					Tcp *struct {
						Address *string `tfsdk:"address" json:"address,omitempty"`
						Format  *struct {
							Json *[]struct {
								Key   *string `tfsdk:"key" json:"key,omitempty"`
								Value *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"json" json:"json,omitempty"`
							OmitEmptyValues *bool   `tfsdk:"omit_empty_values" json:"omitEmptyValues,omitempty"`
							Plain           *string `tfsdk:"plain" json:"plain,omitempty"`
							Type            *string `tfsdk:"type" json:"type,omitempty"`
						} `tfsdk:"format" json:"format,omitempty"`
					} `tfsdk:"tcp" json:"tcp,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"backends" json:"backends,omitempty"`
			} `tfsdk:"default" json:"default,omitempty"`
			TargetRef *struct {
				Kind       *string            `tfsdk:"kind" json:"kind,omitempty"`
				Mesh       *string            `tfsdk:"mesh" json:"mesh,omitempty"`
				Name       *string            `tfsdk:"name" json:"name,omitempty"`
				ProxyTypes *[]string          `tfsdk:"proxy_types" json:"proxyTypes,omitempty"`
				Tags       *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
			} `tfsdk:"target_ref" json:"targetRef,omitempty"`
		} `tfsdk:"to" json:"to,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KumaIoMeshAccessLogV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kuma_io_mesh_access_log_v1alpha1_manifest"
}

func (r *KumaIoMeshAccessLogV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
		Attributes: map[string]schema.Attribute{
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
				Description:         "Spec is the specification of the Kuma MeshAccessLog resource.",
				MarkdownDescription: "Spec is the specification of the Kuma MeshAccessLog resource.",
				Attributes: map[string]schema.Attribute{
					"from": schema.ListNestedAttribute{
						Description:         "From list makes a match between clients and corresponding configurations",
						MarkdownDescription: "From list makes a match between clients and corresponding configurations",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"default": schema.SingleNestedAttribute{
									Description:         "Default is a configuration specific to the group of clients referenced in'targetRef'",
									MarkdownDescription: "Default is a configuration specific to the group of clients referenced in'targetRef'",
									Attributes: map[string]schema.Attribute{
										"backends": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"file": schema.SingleNestedAttribute{
														Description:         "FileBackend defines configuration for file based access logs",
														MarkdownDescription: "FileBackend defines configuration for file based access logs",
														Attributes: map[string]schema.Attribute{
															"format": schema.SingleNestedAttribute{
																Description:         "Format of access logs. Placeholders available onhttps://www.envoyproxy.io/docs/envoy/latest/configuration/observability/access_log/usage#command-operators",
																MarkdownDescription: "Format of access logs. Placeholders available onhttps://www.envoyproxy.io/docs/envoy/latest/configuration/observability/access_log/usage#command-operators",
																Attributes: map[string]schema.Attribute{
																	"json": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"key": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"value": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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

																	"omit_empty_values": schema.BoolAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"plain": schema.StringAttribute{
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
																			stringvalidator.OneOf("Plain", "Json"),
																		},
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"path": schema.StringAttribute{
																Description:         "Path to a file that logs will be written to",
																MarkdownDescription: "Path to a file that logs will be written to",
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

													"open_telemetry": schema.SingleNestedAttribute{
														Description:         "Defines an OpenTelemetry logging backend.",
														MarkdownDescription: "Defines an OpenTelemetry logging backend.",
														Attributes: map[string]schema.Attribute{
															"attributes": schema.ListNestedAttribute{
																Description:         "Attributes can contain placeholders available onhttps://www.envoyproxy.io/docs/envoy/latest/configuration/observability/access_log/usage#command-operators",
																MarkdownDescription: "Attributes can contain placeholders available onhttps://www.envoyproxy.io/docs/envoy/latest/configuration/observability/access_log/usage#command-operators",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"value": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

															"body": schema.MapAttribute{
																Description:         "Body is a raw string or an OTLP any value as described athttps://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/logs/data-model.md#field-bodyIt can contain placeholders available onhttps://www.envoyproxy.io/docs/envoy/latest/configuration/observability/access_log/usage#command-operators",
																MarkdownDescription: "Body is a raw string or an OTLP any value as described athttps://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/logs/data-model.md#field-bodyIt can contain placeholders available onhttps://www.envoyproxy.io/docs/envoy/latest/configuration/observability/access_log/usage#command-operators",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "Endpoint of OpenTelemetry collector. An empty port defaults to 4317.",
																MarkdownDescription: "Endpoint of OpenTelemetry collector. An empty port defaults to 4317.",
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

													"tcp": schema.SingleNestedAttribute{
														Description:         "TCPBackend defines a TCP logging backend.",
														MarkdownDescription: "TCPBackend defines a TCP logging backend.",
														Attributes: map[string]schema.Attribute{
															"address": schema.StringAttribute{
																Description:         "Address of the TCP logging backend",
																MarkdownDescription: "Address of the TCP logging backend",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},

															"format": schema.SingleNestedAttribute{
																Description:         "Format of access logs. Placeholders available onhttps://www.envoyproxy.io/docs/envoy/latest/configuration/observability/access_log/usage#command-operators",
																MarkdownDescription: "Format of access logs. Placeholders available onhttps://www.envoyproxy.io/docs/envoy/latest/configuration/observability/access_log/usage#command-operators",
																Attributes: map[string]schema.Attribute{
																	"json": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"key": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"value": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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

																	"omit_empty_values": schema.BoolAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"plain": schema.StringAttribute{
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
																			stringvalidator.OneOf("Plain", "Json"),
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
															stringvalidator.OneOf("Tcp", "File", "OpenTelemetry"),
														},
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

								"target_ref": schema.SingleNestedAttribute{
									Description:         "TargetRef is a reference to the resource that represents a group ofclients.",
									MarkdownDescription: "TargetRef is a reference to the resource that represents a group ofclients.",
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
											Description:         "Name of the referenced resource. Can only be used with kinds: 'MeshService','MeshServiceSubset' and 'MeshGatewayRoute'",
											MarkdownDescription: "Name of the referenced resource. Can only be used with kinds: 'MeshService','MeshServiceSubset' and 'MeshGatewayRoute'",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"proxy_types": schema.ListAttribute{
											Description:         "ProxyTypes specifies the data plane types that are subject to the policy. When not specified,all data plane types are targeted by the policy.",
											MarkdownDescription: "ProxyTypes specifies the data plane types that are subject to the policy. When not specified,all data plane types are targeted by the policy.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tags": schema.MapAttribute{
											Description:         "Tags used to select a subset of proxies by tags. Can only be used with kinds'MeshSubset' and 'MeshServiceSubset'",
											MarkdownDescription: "Tags used to select a subset of proxies by tags. Can only be used with kinds'MeshSubset' and 'MeshServiceSubset'",
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
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"target_ref": schema.SingleNestedAttribute{
						Description:         "TargetRef is a reference to the resource the policy takes an effect on.The resource could be either a real store object or virtual resourcedefined in-place.",
						MarkdownDescription: "TargetRef is a reference to the resource the policy takes an effect on.The resource could be either a real store object or virtual resourcedefined in-place.",
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
								Description:         "Name of the referenced resource. Can only be used with kinds: 'MeshService','MeshServiceSubset' and 'MeshGatewayRoute'",
								MarkdownDescription: "Name of the referenced resource. Can only be used with kinds: 'MeshService','MeshServiceSubset' and 'MeshGatewayRoute'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"proxy_types": schema.ListAttribute{
								Description:         "ProxyTypes specifies the data plane types that are subject to the policy. When not specified,all data plane types are targeted by the policy.",
								MarkdownDescription: "ProxyTypes specifies the data plane types that are subject to the policy. When not specified,all data plane types are targeted by the policy.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tags": schema.MapAttribute{
								Description:         "Tags used to select a subset of proxies by tags. Can only be used with kinds'MeshSubset' and 'MeshServiceSubset'",
								MarkdownDescription: "Tags used to select a subset of proxies by tags. Can only be used with kinds'MeshSubset' and 'MeshServiceSubset'",
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

					"to": schema.ListNestedAttribute{
						Description:         "To list makes a match between the consumed services and corresponding configurations",
						MarkdownDescription: "To list makes a match between the consumed services and corresponding configurations",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"default": schema.SingleNestedAttribute{
									Description:         "Default is a configuration specific to the group of destinations referenced in'targetRef'",
									MarkdownDescription: "Default is a configuration specific to the group of destinations referenced in'targetRef'",
									Attributes: map[string]schema.Attribute{
										"backends": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"file": schema.SingleNestedAttribute{
														Description:         "FileBackend defines configuration for file based access logs",
														MarkdownDescription: "FileBackend defines configuration for file based access logs",
														Attributes: map[string]schema.Attribute{
															"format": schema.SingleNestedAttribute{
																Description:         "Format of access logs. Placeholders available onhttps://www.envoyproxy.io/docs/envoy/latest/configuration/observability/access_log/usage#command-operators",
																MarkdownDescription: "Format of access logs. Placeholders available onhttps://www.envoyproxy.io/docs/envoy/latest/configuration/observability/access_log/usage#command-operators",
																Attributes: map[string]schema.Attribute{
																	"json": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"key": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"value": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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

																	"omit_empty_values": schema.BoolAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"plain": schema.StringAttribute{
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
																			stringvalidator.OneOf("Plain", "Json"),
																		},
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"path": schema.StringAttribute{
																Description:         "Path to a file that logs will be written to",
																MarkdownDescription: "Path to a file that logs will be written to",
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

													"open_telemetry": schema.SingleNestedAttribute{
														Description:         "Defines an OpenTelemetry logging backend.",
														MarkdownDescription: "Defines an OpenTelemetry logging backend.",
														Attributes: map[string]schema.Attribute{
															"attributes": schema.ListNestedAttribute{
																Description:         "Attributes can contain placeholders available onhttps://www.envoyproxy.io/docs/envoy/latest/configuration/observability/access_log/usage#command-operators",
																MarkdownDescription: "Attributes can contain placeholders available onhttps://www.envoyproxy.io/docs/envoy/latest/configuration/observability/access_log/usage#command-operators",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"value": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

															"body": schema.MapAttribute{
																Description:         "Body is a raw string or an OTLP any value as described athttps://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/logs/data-model.md#field-bodyIt can contain placeholders available onhttps://www.envoyproxy.io/docs/envoy/latest/configuration/observability/access_log/usage#command-operators",
																MarkdownDescription: "Body is a raw string or an OTLP any value as described athttps://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/logs/data-model.md#field-bodyIt can contain placeholders available onhttps://www.envoyproxy.io/docs/envoy/latest/configuration/observability/access_log/usage#command-operators",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "Endpoint of OpenTelemetry collector. An empty port defaults to 4317.",
																MarkdownDescription: "Endpoint of OpenTelemetry collector. An empty port defaults to 4317.",
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

													"tcp": schema.SingleNestedAttribute{
														Description:         "TCPBackend defines a TCP logging backend.",
														MarkdownDescription: "TCPBackend defines a TCP logging backend.",
														Attributes: map[string]schema.Attribute{
															"address": schema.StringAttribute{
																Description:         "Address of the TCP logging backend",
																MarkdownDescription: "Address of the TCP logging backend",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},

															"format": schema.SingleNestedAttribute{
																Description:         "Format of access logs. Placeholders available onhttps://www.envoyproxy.io/docs/envoy/latest/configuration/observability/access_log/usage#command-operators",
																MarkdownDescription: "Format of access logs. Placeholders available onhttps://www.envoyproxy.io/docs/envoy/latest/configuration/observability/access_log/usage#command-operators",
																Attributes: map[string]schema.Attribute{
																	"json": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"key": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"value": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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

																	"omit_empty_values": schema.BoolAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"plain": schema.StringAttribute{
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
																			stringvalidator.OneOf("Plain", "Json"),
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
															stringvalidator.OneOf("Tcp", "File", "OpenTelemetry"),
														},
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

								"target_ref": schema.SingleNestedAttribute{
									Description:         "TargetRef is a reference to the resource that represents a group ofdestinations.",
									MarkdownDescription: "TargetRef is a reference to the resource that represents a group ofdestinations.",
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
											Description:         "Name of the referenced resource. Can only be used with kinds: 'MeshService','MeshServiceSubset' and 'MeshGatewayRoute'",
											MarkdownDescription: "Name of the referenced resource. Can only be used with kinds: 'MeshService','MeshServiceSubset' and 'MeshGatewayRoute'",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"proxy_types": schema.ListAttribute{
											Description:         "ProxyTypes specifies the data plane types that are subject to the policy. When not specified,all data plane types are targeted by the policy.",
											MarkdownDescription: "ProxyTypes specifies the data plane types that are subject to the policy. When not specified,all data plane types are targeted by the policy.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tags": schema.MapAttribute{
											Description:         "Tags used to select a subset of proxies by tags. Can only be used with kinds'MeshSubset' and 'MeshServiceSubset'",
											MarkdownDescription: "Tags used to select a subset of proxies by tags. Can only be used with kinds'MeshSubset' and 'MeshServiceSubset'",
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

func (r *KumaIoMeshAccessLogV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kuma_io_mesh_access_log_v1alpha1_manifest")

	var model KumaIoMeshAccessLogV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("kuma.io/v1alpha1")
	model.Kind = pointer.String("MeshAccessLog")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
