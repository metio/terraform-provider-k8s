/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kuma_io_v1alpha1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &KumaIoMeshRetryV1Alpha1Manifest{}
)

func NewKumaIoMeshRetryV1Alpha1Manifest() datasource.DataSource {
	return &KumaIoMeshRetryV1Alpha1Manifest{}
}

type KumaIoMeshRetryV1Alpha1Manifest struct{}

type KumaIoMeshRetryV1Alpha1ManifestData struct {
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
			Default *struct {
				Grpc *struct {
					BackOff *struct {
						BaseInterval *string `tfsdk:"base_interval" json:"baseInterval,omitempty"`
						MaxInterval  *string `tfsdk:"max_interval" json:"maxInterval,omitempty"`
					} `tfsdk:"back_off" json:"backOff,omitempty"`
					NumRetries         *int64  `tfsdk:"num_retries" json:"numRetries,omitempty"`
					PerTryTimeout      *string `tfsdk:"per_try_timeout" json:"perTryTimeout,omitempty"`
					RateLimitedBackOff *struct {
						MaxInterval  *string `tfsdk:"max_interval" json:"maxInterval,omitempty"`
						ResetHeaders *[]struct {
							Format *string `tfsdk:"format" json:"format,omitempty"`
							Name   *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"reset_headers" json:"resetHeaders,omitempty"`
					} `tfsdk:"rate_limited_back_off" json:"rateLimitedBackOff,omitempty"`
					RetryOn *[]string `tfsdk:"retry_on" json:"retryOn,omitempty"`
				} `tfsdk:"grpc" json:"grpc,omitempty"`
				Http *struct {
					BackOff *struct {
						BaseInterval *string `tfsdk:"base_interval" json:"baseInterval,omitempty"`
						MaxInterval  *string `tfsdk:"max_interval" json:"maxInterval,omitempty"`
					} `tfsdk:"back_off" json:"backOff,omitempty"`
					HostSelection *[]struct {
						Predicate       *string            `tfsdk:"predicate" json:"predicate,omitempty"`
						Tags            *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
						UpdateFrequency *int64             `tfsdk:"update_frequency" json:"updateFrequency,omitempty"`
					} `tfsdk:"host_selection" json:"hostSelection,omitempty"`
					HostSelectionMaxAttempts *int64  `tfsdk:"host_selection_max_attempts" json:"hostSelectionMaxAttempts,omitempty"`
					NumRetries               *int64  `tfsdk:"num_retries" json:"numRetries,omitempty"`
					PerTryTimeout            *string `tfsdk:"per_try_timeout" json:"perTryTimeout,omitempty"`
					RateLimitedBackOff       *struct {
						MaxInterval  *string `tfsdk:"max_interval" json:"maxInterval,omitempty"`
						ResetHeaders *[]struct {
							Format *string `tfsdk:"format" json:"format,omitempty"`
							Name   *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"reset_headers" json:"resetHeaders,omitempty"`
					} `tfsdk:"rate_limited_back_off" json:"rateLimitedBackOff,omitempty"`
					RetriableRequestHeaders *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Type  *string `tfsdk:"type" json:"type,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"retriable_request_headers" json:"retriableRequestHeaders,omitempty"`
					RetriableResponseHeaders *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Type  *string `tfsdk:"type" json:"type,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"retriable_response_headers" json:"retriableResponseHeaders,omitempty"`
					RetryOn *[]string `tfsdk:"retry_on" json:"retryOn,omitempty"`
				} `tfsdk:"http" json:"http,omitempty"`
				Tcp *struct {
					MaxConnectAttempt *int64 `tfsdk:"max_connect_attempt" json:"maxConnectAttempt,omitempty"`
				} `tfsdk:"tcp" json:"tcp,omitempty"`
			} `tfsdk:"default" json:"default,omitempty"`
			TargetRef *struct {
				Kind *string            `tfsdk:"kind" json:"kind,omitempty"`
				Mesh *string            `tfsdk:"mesh" json:"mesh,omitempty"`
				Name *string            `tfsdk:"name" json:"name,omitempty"`
				Tags *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
			} `tfsdk:"target_ref" json:"targetRef,omitempty"`
		} `tfsdk:"to" json:"to,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KumaIoMeshRetryV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kuma_io_mesh_retry_v1alpha1_manifest"
}

func (r *KumaIoMeshRetryV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "Spec is the specification of the Kuma MeshRetry resource.",
				MarkdownDescription: "Spec is the specification of the Kuma MeshRetry resource.",
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
									Description:         "Default is a configuration specific to the group of destinations referenced in 'targetRef'",
									MarkdownDescription: "Default is a configuration specific to the group of destinations referenced in 'targetRef'",
									Attributes: map[string]schema.Attribute{
										"grpc": schema.SingleNestedAttribute{
											Description:         "GRPC defines a configuration of retries for GRPC traffic",
											MarkdownDescription: "GRPC defines a configuration of retries for GRPC traffic",
											Attributes: map[string]schema.Attribute{
												"back_off": schema.SingleNestedAttribute{
													Description:         "BackOff is a configuration of durations which will be used in exponential backoff strategy between retries.",
													MarkdownDescription: "BackOff is a configuration of durations which will be used in exponential backoff strategy between retries.",
													Attributes: map[string]schema.Attribute{
														"base_interval": schema.StringAttribute{
															Description:         "BaseInterval is an amount of time which should be taken between retries. Must be greater than zero. Values less than 1 ms are rounded up to 1 ms. Default is 25ms.",
															MarkdownDescription: "BaseInterval is an amount of time which should be taken between retries. Must be greater than zero. Values less than 1 ms are rounded up to 1 ms. Default is 25ms.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"max_interval": schema.StringAttribute{
															Description:         "MaxInterval is a maximal amount of time which will be taken between retries. Default is 10 times the 'BaseInterval'.",
															MarkdownDescription: "MaxInterval is a maximal amount of time which will be taken between retries. Default is 10 times the 'BaseInterval'.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"num_retries": schema.Int64Attribute{
													Description:         "NumRetries is the number of attempts that will be made on failed (and retriable) requests.",
													MarkdownDescription: "NumRetries is the number of attempts that will be made on failed (and retriable) requests.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"per_try_timeout": schema.StringAttribute{
													Description:         "PerTryTimeout is the amount of time after which retry attempt should timeout. Setting this timeout to 0 will disable it. Default is 15s.",
													MarkdownDescription: "PerTryTimeout is the amount of time after which retry attempt should timeout. Setting this timeout to 0 will disable it. Default is 15s.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"rate_limited_back_off": schema.SingleNestedAttribute{
													Description:         "RateLimitedBackOff is a configuration of backoff which will be used when the upstream returns one of the headers configured.",
													MarkdownDescription: "RateLimitedBackOff is a configuration of backoff which will be used when the upstream returns one of the headers configured.",
													Attributes: map[string]schema.Attribute{
														"max_interval": schema.StringAttribute{
															Description:         "MaxInterval is a maximal amount of time which will be taken between retries. Default is 300 seconds.",
															MarkdownDescription: "MaxInterval is a maximal amount of time which will be taken between retries. Default is 300 seconds.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"reset_headers": schema.ListNestedAttribute{
															Description:         "ResetHeaders specifies the list of headers (like Retry-After or X-RateLimit-Reset) to match against the response. Headers are tried in order, and matched case-insensitive. The first header to be parsed successfully is used. If no headers match the default exponential BackOff is used instead.",
															MarkdownDescription: "ResetHeaders specifies the list of headers (like Retry-After or X-RateLimit-Reset) to match against the response. Headers are tried in order, and matched case-insensitive. The first header to be parsed successfully is used. If no headers match the default exponential BackOff is used instead.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"format": schema.StringAttribute{
																		Description:         "The format of the reset header, either Seconds or UnixTimestamp.",
																		MarkdownDescription: "The format of the reset header, either Seconds or UnixTimestamp.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("Seconds", "UnixTimestamp"),
																		},
																	},

																	"name": schema.StringAttribute{
																		Description:         "The Name of the reset header.",
																		MarkdownDescription: "The Name of the reset header.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.LengthAtLeast(1),
																			stringvalidator.LengthAtMost(256),
																			stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
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

												"retry_on": schema.ListAttribute{
													Description:         "RetryOn is a list of conditions which will cause a retry. Available values are: [Canceled, DeadlineExceeded, Internal, ResourceExhausted, Unavailable].",
													MarkdownDescription: "RetryOn is a list of conditions which will cause a retry. Available values are: [Canceled, DeadlineExceeded, Internal, ResourceExhausted, Unavailable].",
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

										"http": schema.SingleNestedAttribute{
											Description:         "HTTP defines a configuration of retries for HTTP traffic",
											MarkdownDescription: "HTTP defines a configuration of retries for HTTP traffic",
											Attributes: map[string]schema.Attribute{
												"back_off": schema.SingleNestedAttribute{
													Description:         "BackOff is a configuration of durations which will be used in exponential backoff strategy between retries",
													MarkdownDescription: "BackOff is a configuration of durations which will be used in exponential backoff strategy between retries",
													Attributes: map[string]schema.Attribute{
														"base_interval": schema.StringAttribute{
															Description:         "BaseInterval is an amount of time which should be taken between retries. Must be greater than zero. Values less than 1 ms are rounded up to 1 ms. Default is 25ms.",
															MarkdownDescription: "BaseInterval is an amount of time which should be taken between retries. Must be greater than zero. Values less than 1 ms are rounded up to 1 ms. Default is 25ms.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"max_interval": schema.StringAttribute{
															Description:         "MaxInterval is a maximal amount of time which will be taken between retries. Default is 10 times the 'BaseInterval'.",
															MarkdownDescription: "MaxInterval is a maximal amount of time which will be taken between retries. Default is 10 times the 'BaseInterval'.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"host_selection": schema.ListNestedAttribute{
													Description:         "HostSelection is a list of predicates that dictate how hosts should be selected when requests are retried.",
													MarkdownDescription: "HostSelection is a list of predicates that dictate how hosts should be selected when requests are retried.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"predicate": schema.StringAttribute{
																Description:         "Type is requested predicate mode. Available values are OmitPreviousHosts, OmitHostsWithTags, and OmitPreviousPriorities.",
																MarkdownDescription: "Type is requested predicate mode. Available values are OmitPreviousHosts, OmitHostsWithTags, and OmitPreviousPriorities.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"tags": schema.MapAttribute{
																Description:         "Tags is a map of metadata to match against for selecting the omitted hosts. Required if Type is OmitHostsWithTags",
																MarkdownDescription: "Tags is a map of metadata to match against for selecting the omitted hosts. Required if Type is OmitHostsWithTags",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"update_frequency": schema.Int64Attribute{
																Description:         "UpdateFrequency is how often the priority load should be updated based on previously attempted priorities. Used for OmitPreviousPriorities. Default is 2 if not set.",
																MarkdownDescription: "UpdateFrequency is how often the priority load should be updated based on previously attempted priorities. Used for OmitPreviousPriorities. Default is 2 if not set.",
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

												"host_selection_max_attempts": schema.Int64Attribute{
													Description:         "HostSelectionMaxAttempts is the maximum number of times host selection will be reattempted before giving up, at which point the host that was last selected will be routed to. If unspecified, this will default to retrying once.",
													MarkdownDescription: "HostSelectionMaxAttempts is the maximum number of times host selection will be reattempted before giving up, at which point the host that was last selected will be routed to. If unspecified, this will default to retrying once.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"num_retries": schema.Int64Attribute{
													Description:         "NumRetries is the number of attempts that will be made on failed (and retriable) requests",
													MarkdownDescription: "NumRetries is the number of attempts that will be made on failed (and retriable) requests",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"per_try_timeout": schema.StringAttribute{
													Description:         "PerTryTimeout is the amount of time after which retry attempt should timeout. Setting this timeout to 0 will disable it. Default is 15s.",
													MarkdownDescription: "PerTryTimeout is the amount of time after which retry attempt should timeout. Setting this timeout to 0 will disable it. Default is 15s.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"rate_limited_back_off": schema.SingleNestedAttribute{
													Description:         "RateLimitedBackOff is a configuration of backoff which will be used when the upstream returns one of the headers configured.",
													MarkdownDescription: "RateLimitedBackOff is a configuration of backoff which will be used when the upstream returns one of the headers configured.",
													Attributes: map[string]schema.Attribute{
														"max_interval": schema.StringAttribute{
															Description:         "MaxInterval is a maximal amount of time which will be taken between retries. Default is 300 seconds.",
															MarkdownDescription: "MaxInterval is a maximal amount of time which will be taken between retries. Default is 300 seconds.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"reset_headers": schema.ListNestedAttribute{
															Description:         "ResetHeaders specifies the list of headers (like Retry-After or X-RateLimit-Reset) to match against the response. Headers are tried in order, and matched case-insensitive. The first header to be parsed successfully is used. If no headers match the default exponential BackOff is used instead.",
															MarkdownDescription: "ResetHeaders specifies the list of headers (like Retry-After or X-RateLimit-Reset) to match against the response. Headers are tried in order, and matched case-insensitive. The first header to be parsed successfully is used. If no headers match the default exponential BackOff is used instead.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"format": schema.StringAttribute{
																		Description:         "The format of the reset header, either Seconds or UnixTimestamp.",
																		MarkdownDescription: "The format of the reset header, either Seconds or UnixTimestamp.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("Seconds", "UnixTimestamp"),
																		},
																	},

																	"name": schema.StringAttribute{
																		Description:         "The Name of the reset header.",
																		MarkdownDescription: "The Name of the reset header.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.LengthAtLeast(1),
																			stringvalidator.LengthAtMost(256),
																			stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
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

												"retriable_request_headers": schema.ListNestedAttribute{
													Description:         "RetriableRequestHeaders is an HTTP headers which must be present in the request for retries to be attempted.",
													MarkdownDescription: "RetriableRequestHeaders is an HTTP headers which must be present in the request for retries to be attempted.",
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

												"retriable_response_headers": schema.ListNestedAttribute{
													Description:         "RetriableResponseHeaders is an HTTP response headers that trigger a retry if present in the response. A retry will be triggered if any of the header matches match the upstream response headers.",
													MarkdownDescription: "RetriableResponseHeaders is an HTTP response headers that trigger a retry if present in the response. A retry will be triggered if any of the header matches match the upstream response headers.",
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

												"retry_on": schema.ListAttribute{
													Description:         "RetryOn is a list of conditions which will cause a retry. Available values are: [5XX, GatewayError, Reset, Retriable4xx, ConnectFailure, EnvoyRatelimited, RefusedStream, Http3PostConnectFailure, HttpMethodConnect, HttpMethodDelete, HttpMethodGet, HttpMethodHead, HttpMethodOptions, HttpMethodPatch, HttpMethodPost, HttpMethodPut, HttpMethodTrace]. Also, any HTTP status code (500, 503, etc).",
													MarkdownDescription: "RetryOn is a list of conditions which will cause a retry. Available values are: [5XX, GatewayError, Reset, Retriable4xx, ConnectFailure, EnvoyRatelimited, RefusedStream, Http3PostConnectFailure, HttpMethodConnect, HttpMethodDelete, HttpMethodGet, HttpMethodHead, HttpMethodOptions, HttpMethodPatch, HttpMethodPost, HttpMethodPut, HttpMethodTrace]. Also, any HTTP status code (500, 503, etc).",
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

										"tcp": schema.SingleNestedAttribute{
											Description:         "TCP defines a configuration of retries for TCP traffic",
											MarkdownDescription: "TCP defines a configuration of retries for TCP traffic",
											Attributes: map[string]schema.Attribute{
												"max_connect_attempt": schema.Int64Attribute{
													Description:         "MaxConnectAttempt is a maximal amount of TCP connection attempts which will be made before giving up",
													MarkdownDescription: "MaxConnectAttempt is a maximal amount of TCP connection attempts which will be made before giving up",
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
									Required: false,
									Optional: true,
									Computed: false,
								},

								"target_ref": schema.SingleNestedAttribute{
									Description:         "TargetRef is a reference to the resource that represents a group of destinations.",
									MarkdownDescription: "TargetRef is a reference to the resource that represents a group of destinations.",
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

func (r *KumaIoMeshRetryV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kuma_io_mesh_retry_v1alpha1_manifest")

	var model KumaIoMeshRetryV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("kuma.io/v1alpha1")
	model.Kind = pointer.String("MeshRetry")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
