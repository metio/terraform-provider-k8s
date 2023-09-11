/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package metacontroller_k8s_io_v1alpha1

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
	_ datasource.DataSource              = &MetacontrollerK8SIoCompositeControllerV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &MetacontrollerK8SIoCompositeControllerV1Alpha1DataSource{}
)

func NewMetacontrollerK8SIoCompositeControllerV1Alpha1DataSource() datasource.DataSource {
	return &MetacontrollerK8SIoCompositeControllerV1Alpha1DataSource{}
}

type MetacontrollerK8SIoCompositeControllerV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type MetacontrollerK8SIoCompositeControllerV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ChildResources *[]struct {
			ApiVersion     *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Resource       *string `tfsdk:"resource" json:"resource,omitempty"`
			UpdateStrategy *struct {
				Method       *string `tfsdk:"method" json:"method,omitempty"`
				StatusChecks *struct {
					Conditions *[]struct {
						Reason *string `tfsdk:"reason" json:"reason,omitempty"`
						Status *string `tfsdk:"status" json:"status,omitempty"`
						Type   *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"conditions" json:"conditions,omitempty"`
				} `tfsdk:"status_checks" json:"statusChecks,omitempty"`
			} `tfsdk:"update_strategy" json:"updateStrategy,omitempty"`
		} `tfsdk:"child_resources" json:"childResources,omitempty"`
		GenerateSelector *bool `tfsdk:"generate_selector" json:"generateSelector,omitempty"`
		Hooks            *struct {
			Customize *struct {
				Version *string `tfsdk:"version" json:"version,omitempty"`
				Webhook *struct {
					Etag *struct {
						CacheCleanupSeconds *int64 `tfsdk:"cache_cleanup_seconds" json:"cacheCleanupSeconds,omitempty"`
						CacheTimeoutSeconds *int64 `tfsdk:"cache_timeout_seconds" json:"cacheTimeoutSeconds,omitempty"`
						Enabled             *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
					} `tfsdk:"etag" json:"etag,omitempty"`
					Path                   *string `tfsdk:"path" json:"path,omitempty"`
					ResponseUnMarshallMode *string `tfsdk:"response_un_marshall_mode" json:"responseUnMarshallMode,omitempty"`
					Service                *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						Port      *int64  `tfsdk:"port" json:"port,omitempty"`
						Protocol  *string `tfsdk:"protocol" json:"protocol,omitempty"`
					} `tfsdk:"service" json:"service,omitempty"`
					Timeout *string `tfsdk:"timeout" json:"timeout,omitempty"`
					Url     *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"webhook" json:"webhook,omitempty"`
			} `tfsdk:"customize" json:"customize,omitempty"`
			Finalize *struct {
				Version *string `tfsdk:"version" json:"version,omitempty"`
				Webhook *struct {
					Etag *struct {
						CacheCleanupSeconds *int64 `tfsdk:"cache_cleanup_seconds" json:"cacheCleanupSeconds,omitempty"`
						CacheTimeoutSeconds *int64 `tfsdk:"cache_timeout_seconds" json:"cacheTimeoutSeconds,omitempty"`
						Enabled             *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
					} `tfsdk:"etag" json:"etag,omitempty"`
					Path                   *string `tfsdk:"path" json:"path,omitempty"`
					ResponseUnMarshallMode *string `tfsdk:"response_un_marshall_mode" json:"responseUnMarshallMode,omitempty"`
					Service                *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						Port      *int64  `tfsdk:"port" json:"port,omitempty"`
						Protocol  *string `tfsdk:"protocol" json:"protocol,omitempty"`
					} `tfsdk:"service" json:"service,omitempty"`
					Timeout *string `tfsdk:"timeout" json:"timeout,omitempty"`
					Url     *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"webhook" json:"webhook,omitempty"`
			} `tfsdk:"finalize" json:"finalize,omitempty"`
			PostUpdateChild *struct {
				Version *string `tfsdk:"version" json:"version,omitempty"`
				Webhook *struct {
					Etag *struct {
						CacheCleanupSeconds *int64 `tfsdk:"cache_cleanup_seconds" json:"cacheCleanupSeconds,omitempty"`
						CacheTimeoutSeconds *int64 `tfsdk:"cache_timeout_seconds" json:"cacheTimeoutSeconds,omitempty"`
						Enabled             *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
					} `tfsdk:"etag" json:"etag,omitempty"`
					Path                   *string `tfsdk:"path" json:"path,omitempty"`
					ResponseUnMarshallMode *string `tfsdk:"response_un_marshall_mode" json:"responseUnMarshallMode,omitempty"`
					Service                *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						Port      *int64  `tfsdk:"port" json:"port,omitempty"`
						Protocol  *string `tfsdk:"protocol" json:"protocol,omitempty"`
					} `tfsdk:"service" json:"service,omitempty"`
					Timeout *string `tfsdk:"timeout" json:"timeout,omitempty"`
					Url     *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"webhook" json:"webhook,omitempty"`
			} `tfsdk:"post_update_child" json:"postUpdateChild,omitempty"`
			PreUpdateChild *struct {
				Version *string `tfsdk:"version" json:"version,omitempty"`
				Webhook *struct {
					Etag *struct {
						CacheCleanupSeconds *int64 `tfsdk:"cache_cleanup_seconds" json:"cacheCleanupSeconds,omitempty"`
						CacheTimeoutSeconds *int64 `tfsdk:"cache_timeout_seconds" json:"cacheTimeoutSeconds,omitempty"`
						Enabled             *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
					} `tfsdk:"etag" json:"etag,omitempty"`
					Path                   *string `tfsdk:"path" json:"path,omitempty"`
					ResponseUnMarshallMode *string `tfsdk:"response_un_marshall_mode" json:"responseUnMarshallMode,omitempty"`
					Service                *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						Port      *int64  `tfsdk:"port" json:"port,omitempty"`
						Protocol  *string `tfsdk:"protocol" json:"protocol,omitempty"`
					} `tfsdk:"service" json:"service,omitempty"`
					Timeout *string `tfsdk:"timeout" json:"timeout,omitempty"`
					Url     *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"webhook" json:"webhook,omitempty"`
			} `tfsdk:"pre_update_child" json:"preUpdateChild,omitempty"`
			Sync *struct {
				Version *string `tfsdk:"version" json:"version,omitempty"`
				Webhook *struct {
					Etag *struct {
						CacheCleanupSeconds *int64 `tfsdk:"cache_cleanup_seconds" json:"cacheCleanupSeconds,omitempty"`
						CacheTimeoutSeconds *int64 `tfsdk:"cache_timeout_seconds" json:"cacheTimeoutSeconds,omitempty"`
						Enabled             *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
					} `tfsdk:"etag" json:"etag,omitempty"`
					Path                   *string `tfsdk:"path" json:"path,omitempty"`
					ResponseUnMarshallMode *string `tfsdk:"response_un_marshall_mode" json:"responseUnMarshallMode,omitempty"`
					Service                *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						Port      *int64  `tfsdk:"port" json:"port,omitempty"`
						Protocol  *string `tfsdk:"protocol" json:"protocol,omitempty"`
					} `tfsdk:"service" json:"service,omitempty"`
					Timeout *string `tfsdk:"timeout" json:"timeout,omitempty"`
					Url     *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"webhook" json:"webhook,omitempty"`
			} `tfsdk:"sync" json:"sync,omitempty"`
		} `tfsdk:"hooks" json:"hooks,omitempty"`
		ParentResource *struct {
			ApiVersion    *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			LabelSelector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
			Resource        *string `tfsdk:"resource" json:"resource,omitempty"`
			RevisionHistory *struct {
				FieldPaths *[]string `tfsdk:"field_paths" json:"fieldPaths,omitempty"`
			} `tfsdk:"revision_history" json:"revisionHistory,omitempty"`
		} `tfsdk:"parent_resource" json:"parentResource,omitempty"`
		ResyncPeriodSeconds *int64 `tfsdk:"resync_period_seconds" json:"resyncPeriodSeconds,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *MetacontrollerK8SIoCompositeControllerV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_metacontroller_k8s_io_composite_controller_v1alpha1"
}

func (r *MetacontrollerK8SIoCompositeControllerV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CompositeController",
		MarkdownDescription: "CompositeController",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"child_resources": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"api_version": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"resource": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"update_strategy": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"method": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"status_checks": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"conditions": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"reason": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"status": schema.StringAttribute{
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"generate_selector": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"hooks": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"customize": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"version": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"webhook": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"etag": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"cache_cleanup_seconds": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"cache_timeout_seconds": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"enabled": schema.BoolAttribute{
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

											"path": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"response_un_marshall_mode": schema.StringAttribute{
												Description:         "Sets the json unmarshall mode. One of the 'loose' or 'strict'. In 'strict' mode additional checks are performed to detect unknown and duplicated fields.",
												MarkdownDescription: "Sets the json unmarshall mode. One of the 'loose' or 'strict'. In 'strict' mode additional checks are performed to detect unknown and duplicated fields.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"service": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"namespace": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"port": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"protocol": schema.StringAttribute{
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

											"timeout": schema.StringAttribute{
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
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"finalize": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"version": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"webhook": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"etag": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"cache_cleanup_seconds": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"cache_timeout_seconds": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"enabled": schema.BoolAttribute{
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

											"path": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"response_un_marshall_mode": schema.StringAttribute{
												Description:         "Sets the json unmarshall mode. One of the 'loose' or 'strict'. In 'strict' mode additional checks are performed to detect unknown and duplicated fields.",
												MarkdownDescription: "Sets the json unmarshall mode. One of the 'loose' or 'strict'. In 'strict' mode additional checks are performed to detect unknown and duplicated fields.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"service": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"namespace": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"port": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"protocol": schema.StringAttribute{
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

											"timeout": schema.StringAttribute{
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
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"post_update_child": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"version": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"webhook": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"etag": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"cache_cleanup_seconds": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"cache_timeout_seconds": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"enabled": schema.BoolAttribute{
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

											"path": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"response_un_marshall_mode": schema.StringAttribute{
												Description:         "Sets the json unmarshall mode. One of the 'loose' or 'strict'. In 'strict' mode additional checks are performed to detect unknown and duplicated fields.",
												MarkdownDescription: "Sets the json unmarshall mode. One of the 'loose' or 'strict'. In 'strict' mode additional checks are performed to detect unknown and duplicated fields.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"service": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"namespace": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"port": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"protocol": schema.StringAttribute{
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

											"timeout": schema.StringAttribute{
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
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"pre_update_child": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"version": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"webhook": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"etag": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"cache_cleanup_seconds": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"cache_timeout_seconds": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"enabled": schema.BoolAttribute{
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

											"path": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"response_un_marshall_mode": schema.StringAttribute{
												Description:         "Sets the json unmarshall mode. One of the 'loose' or 'strict'. In 'strict' mode additional checks are performed to detect unknown and duplicated fields.",
												MarkdownDescription: "Sets the json unmarshall mode. One of the 'loose' or 'strict'. In 'strict' mode additional checks are performed to detect unknown and duplicated fields.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"service": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"namespace": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"port": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"protocol": schema.StringAttribute{
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

											"timeout": schema.StringAttribute{
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
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"sync": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"version": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"webhook": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"etag": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"cache_cleanup_seconds": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"cache_timeout_seconds": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"enabled": schema.BoolAttribute{
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

											"path": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"response_un_marshall_mode": schema.StringAttribute{
												Description:         "Sets the json unmarshall mode. One of the 'loose' or 'strict'. In 'strict' mode additional checks are performed to detect unknown and duplicated fields.",
												MarkdownDescription: "Sets the json unmarshall mode. One of the 'loose' or 'strict'. In 'strict' mode additional checks are performed to detect unknown and duplicated fields.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"service": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"namespace": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"port": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"protocol": schema.StringAttribute{
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

											"timeout": schema.StringAttribute{
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

					"parent_resource": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"label_selector": schema.SingleNestedAttribute{
								Description:         "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
								MarkdownDescription: "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
								Attributes: map[string]schema.Attribute{
									"match_expressions": schema.ListNestedAttribute{
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "key is the label key that the selector applies to.",
													MarkdownDescription: "key is the label key that the selector applies to.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"operator": schema.StringAttribute{
													Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
													MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"values": schema.ListAttribute{
													Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
													MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
													ElementType:         types.StringType,
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

									"match_labels": schema.MapAttribute{
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

							"resource": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"revision_history": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"field_paths": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
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

					"resync_period_seconds": schema.Int64Attribute{
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
	}
}

func (r *MetacontrollerK8SIoCompositeControllerV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *MetacontrollerK8SIoCompositeControllerV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_metacontroller_k8s_io_composite_controller_v1alpha1")

	var data MetacontrollerK8SIoCompositeControllerV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "metacontroller.k8s.io", Version: "v1alpha1", Resource: "compositecontrollers"}).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		var statusError *k8sErrors.StatusError
		if errors.As(err, &statusError) {
			if statusError.Status().Code == http.StatusNotFound {
				response.Diagnostics.AddError(
					"Unable to find resource",
					fmt.Sprintf("The requested resource cannot be found. "+
						"Make sure that it does exist in your cluster and you have set the correct name configured.\n\n"+
						"Name: %s", data.Metadata.Name),
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

	var readResponse MetacontrollerK8SIoCompositeControllerV1Alpha1DataSourceData
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
	data.ApiVersion = pointer.String("metacontroller.k8s.io/v1alpha1")
	data.Kind = pointer.String("CompositeController")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
