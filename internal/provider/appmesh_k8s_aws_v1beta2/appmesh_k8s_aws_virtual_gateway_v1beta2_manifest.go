/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package appmesh_k8s_aws_v1beta2

import (
	"context"
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
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &AppmeshK8SAwsVirtualGatewayV1Beta2Manifest{}
)

func NewAppmeshK8SAwsVirtualGatewayV1Beta2Manifest() datasource.DataSource {
	return &AppmeshK8SAwsVirtualGatewayV1Beta2Manifest{}
}

type AppmeshK8SAwsVirtualGatewayV1Beta2Manifest struct{}

type AppmeshK8SAwsVirtualGatewayV1Beta2ManifestData struct {
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
		AwsName         *string `tfsdk:"aws_name" json:"awsName,omitempty"`
		BackendDefaults *struct {
			ClientPolicy *struct {
				Tls *struct {
					Certificate *struct {
						File *struct {
							CertificateChain *string `tfsdk:"certificate_chain" json:"certificateChain,omitempty"`
							PrivateKey       *string `tfsdk:"private_key" json:"privateKey,omitempty"`
						} `tfsdk:"file" json:"file,omitempty"`
						Sds *struct {
							SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
						} `tfsdk:"sds" json:"sds,omitempty"`
					} `tfsdk:"certificate" json:"certificate,omitempty"`
					Enforce    *bool     `tfsdk:"enforce" json:"enforce,omitempty"`
					Ports      *[]string `tfsdk:"ports" json:"ports,omitempty"`
					Validation *struct {
						SubjectAlternativeNames *struct {
							Match *struct {
								Exact *[]string `tfsdk:"exact" json:"exact,omitempty"`
							} `tfsdk:"match" json:"match,omitempty"`
						} `tfsdk:"subject_alternative_names" json:"subjectAlternativeNames,omitempty"`
						Trust *struct {
							Acm *struct {
								CertificateAuthorityARNs *[]string `tfsdk:"certificate_authority_ar_ns" json:"certificateAuthorityARNs,omitempty"`
							} `tfsdk:"acm" json:"acm,omitempty"`
							File *struct {
								CertificateChain *string `tfsdk:"certificate_chain" json:"certificateChain,omitempty"`
							} `tfsdk:"file" json:"file,omitempty"`
							Sds *struct {
								SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
							} `tfsdk:"sds" json:"sds,omitempty"`
						} `tfsdk:"trust" json:"trust,omitempty"`
					} `tfsdk:"validation" json:"validation,omitempty"`
				} `tfsdk:"tls" json:"tls,omitempty"`
			} `tfsdk:"client_policy" json:"clientPolicy,omitempty"`
		} `tfsdk:"backend_defaults" json:"backendDefaults,omitempty"`
		GatewayRouteSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"gateway_route_selector" json:"gatewayRouteSelector,omitempty"`
		Listeners *[]struct {
			ConnectionPool *struct {
				Grpc *struct {
					MaxRequests *int64 `tfsdk:"max_requests" json:"maxRequests,omitempty"`
				} `tfsdk:"grpc" json:"grpc,omitempty"`
				Http *struct {
					MaxConnections     *int64 `tfsdk:"max_connections" json:"maxConnections,omitempty"`
					MaxPendingRequests *int64 `tfsdk:"max_pending_requests" json:"maxPendingRequests,omitempty"`
				} `tfsdk:"http" json:"http,omitempty"`
				Http2 *struct {
					MaxRequests *int64 `tfsdk:"max_requests" json:"maxRequests,omitempty"`
				} `tfsdk:"http2" json:"http2,omitempty"`
			} `tfsdk:"connection_pool" json:"connectionPool,omitempty"`
			HealthCheck *struct {
				HealthyThreshold   *int64  `tfsdk:"healthy_threshold" json:"healthyThreshold,omitempty"`
				IntervalMillis     *int64  `tfsdk:"interval_millis" json:"intervalMillis,omitempty"`
				Path               *string `tfsdk:"path" json:"path,omitempty"`
				Port               *int64  `tfsdk:"port" json:"port,omitempty"`
				Protocol           *string `tfsdk:"protocol" json:"protocol,omitempty"`
				TimeoutMillis      *int64  `tfsdk:"timeout_millis" json:"timeoutMillis,omitempty"`
				UnhealthyThreshold *int64  `tfsdk:"unhealthy_threshold" json:"unhealthyThreshold,omitempty"`
			} `tfsdk:"health_check" json:"healthCheck,omitempty"`
			PortMapping *struct {
				Port     *int64  `tfsdk:"port" json:"port,omitempty"`
				Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
			} `tfsdk:"port_mapping" json:"portMapping,omitempty"`
			Tls *struct {
				Certificate *struct {
					Acm *struct {
						CertificateARN *string `tfsdk:"certificate_arn" json:"certificateARN,omitempty"`
					} `tfsdk:"acm" json:"acm,omitempty"`
					File *struct {
						CertificateChain *string `tfsdk:"certificate_chain" json:"certificateChain,omitempty"`
						PrivateKey       *string `tfsdk:"private_key" json:"privateKey,omitempty"`
					} `tfsdk:"file" json:"file,omitempty"`
					Sds *struct {
						SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
					} `tfsdk:"sds" json:"sds,omitempty"`
				} `tfsdk:"certificate" json:"certificate,omitempty"`
				Mode       *string `tfsdk:"mode" json:"mode,omitempty"`
				Validation *struct {
					SubjectAlternativeNames *struct {
						Match *struct {
							Exact *[]string `tfsdk:"exact" json:"exact,omitempty"`
						} `tfsdk:"match" json:"match,omitempty"`
					} `tfsdk:"subject_alternative_names" json:"subjectAlternativeNames,omitempty"`
					Trust *struct {
						Acm *struct {
							CertificateAuthorityARNs *[]string `tfsdk:"certificate_authority_ar_ns" json:"certificateAuthorityARNs,omitempty"`
						} `tfsdk:"acm" json:"acm,omitempty"`
						File *struct {
							CertificateChain *string `tfsdk:"certificate_chain" json:"certificateChain,omitempty"`
						} `tfsdk:"file" json:"file,omitempty"`
						Sds *struct {
							SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
						} `tfsdk:"sds" json:"sds,omitempty"`
					} `tfsdk:"trust" json:"trust,omitempty"`
				} `tfsdk:"validation" json:"validation,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
		} `tfsdk:"listeners" json:"listeners,omitempty"`
		Logging *struct {
			AccessLog *struct {
				File *struct {
					Format *struct {
						Json *[]struct {
							Key   *string `tfsdk:"key" json:"key,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"json" json:"json,omitempty"`
						Text *string `tfsdk:"text" json:"text,omitempty"`
					} `tfsdk:"format" json:"format,omitempty"`
					Path *string `tfsdk:"path" json:"path,omitempty"`
				} `tfsdk:"file" json:"file,omitempty"`
			} `tfsdk:"access_log" json:"accessLog,omitempty"`
		} `tfsdk:"logging" json:"logging,omitempty"`
		MeshRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
			Uid  *string `tfsdk:"uid" json:"uid,omitempty"`
		} `tfsdk:"mesh_ref" json:"meshRef,omitempty"`
		NamespaceSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
		PodSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"pod_selector" json:"podSelector,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppmeshK8SAwsVirtualGatewayV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_appmesh_k8s_aws_virtual_gateway_v1beta2_manifest"
}

func (r *AppmeshK8SAwsVirtualGatewayV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VirtualGateway is the Schema for the virtualgateways API",
		MarkdownDescription: "VirtualGateway is the Schema for the virtualgateways API",
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
				Description:         "VirtualGatewaySpec defines the desired state of VirtualGateway refers to https://docs.aws.amazon.com/app-mesh/latest/userguide/virtual_gateways.html",
				MarkdownDescription: "VirtualGatewaySpec defines the desired state of VirtualGateway refers to https://docs.aws.amazon.com/app-mesh/latest/userguide/virtual_gateways.html",
				Attributes: map[string]schema.Attribute{
					"aws_name": schema.StringAttribute{
						Description:         "AWSName is the AppMesh VirtualGateway object's name. If unspecified or empty, it defaults to be '${name}_${namespace}' of k8s VirtualGateway",
						MarkdownDescription: "AWSName is the AppMesh VirtualGateway object's name. If unspecified or empty, it defaults to be '${name}_${namespace}' of k8s VirtualGateway",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"backend_defaults": schema.SingleNestedAttribute{
						Description:         "A reference to an object that represents the defaults for backend GatewayRoutes.",
						MarkdownDescription: "A reference to an object that represents the defaults for backend GatewayRoutes.",
						Attributes: map[string]schema.Attribute{
							"client_policy": schema.SingleNestedAttribute{
								Description:         "A reference to an object that represents a client policy.",
								MarkdownDescription: "A reference to an object that represents a client policy.",
								Attributes: map[string]schema.Attribute{
									"tls": schema.SingleNestedAttribute{
										Description:         "A reference to an object that represents a Transport Layer Security (TLS) client policy.",
										MarkdownDescription: "A reference to an object that represents a Transport Layer Security (TLS) client policy.",
										Attributes: map[string]schema.Attribute{
											"certificate": schema.SingleNestedAttribute{
												Description:         "A reference to an object that represents TLS certificate.",
												MarkdownDescription: "A reference to an object that represents TLS certificate.",
												Attributes: map[string]schema.Attribute{
													"file": schema.SingleNestedAttribute{
														Description:         "An object that represents a TLS cert via a local file",
														MarkdownDescription: "An object that represents a TLS cert via a local file",
														Attributes: map[string]schema.Attribute{
															"certificate_chain": schema.StringAttribute{
																Description:         "The certificate chain for the certificate.",
																MarkdownDescription: "The certificate chain for the certificate.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																	stringvalidator.LengthAtMost(255),
																},
															},

															"private_key": schema.StringAttribute{
																Description:         "The private key for a certificate stored on the file system of the virtual Gateway.",
																MarkdownDescription: "The private key for a certificate stored on the file system of the virtual Gateway.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																	stringvalidator.LengthAtMost(255),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"sds": schema.SingleNestedAttribute{
														Description:         "An object that represents a TLS cert via SDS entry",
														MarkdownDescription: "An object that represents a TLS cert via SDS entry",
														Attributes: map[string]schema.Attribute{
															"secret_name": schema.StringAttribute{
																Description:         "The certificate trust chain for a certificate issued via SDS cluster",
																MarkdownDescription: "The certificate trust chain for a certificate issued via SDS cluster",
																Required:            true,
																Optional:            false,
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

											"enforce": schema.BoolAttribute{
												Description:         "Whether the policy is enforced. If unspecified, default settings from AWS API will be applied. Refer to AWS Docs for default settings.",
												MarkdownDescription: "Whether the policy is enforced. If unspecified, default settings from AWS API will be applied. Refer to AWS Docs for default settings.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ports": schema.ListAttribute{
												Description:         "The range of ports that the policy is enforced for.",
												MarkdownDescription: "The range of ports that the policy is enforced for.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"validation": schema.SingleNestedAttribute{
												Description:         "A reference to an object that represents a TLS validation context.",
												MarkdownDescription: "A reference to an object that represents a TLS validation context.",
												Attributes: map[string]schema.Attribute{
													"subject_alternative_names": schema.SingleNestedAttribute{
														Description:         "Possible alternative names to consider",
														MarkdownDescription: "Possible alternative names to consider",
														Attributes: map[string]schema.Attribute{
															"match": schema.SingleNestedAttribute{
																Description:         "Match is a required field",
																MarkdownDescription: "Match is a required field",
																Attributes: map[string]schema.Attribute{
																	"exact": schema.ListAttribute{
																		Description:         "Exact is a required field",
																		MarkdownDescription: "Exact is a required field",
																		ElementType:         types.StringType,
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},
																},
																Required: true,
																Optional: false,
																Computed: false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"trust": schema.SingleNestedAttribute{
														Description:         "A reference to an object that represents a TLS validation context trust",
														MarkdownDescription: "A reference to an object that represents a TLS validation context trust",
														Attributes: map[string]schema.Attribute{
															"acm": schema.SingleNestedAttribute{
																Description:         "A reference to an object that represents a TLS validation context trust for an AWS Certicate Manager (ACM) certificate.",
																MarkdownDescription: "A reference to an object that represents a TLS validation context trust for an AWS Certicate Manager (ACM) certificate.",
																Attributes: map[string]schema.Attribute{
																	"certificate_authority_ar_ns": schema.ListAttribute{
																		Description:         "One or more ACM Amazon Resource Name (ARN)s.",
																		MarkdownDescription: "One or more ACM Amazon Resource Name (ARN)s.",
																		ElementType:         types.StringType,
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"file": schema.SingleNestedAttribute{
																Description:         "An object that represents a TLS validation context trust for a local file.",
																MarkdownDescription: "An object that represents a TLS validation context trust for a local file.",
																Attributes: map[string]schema.Attribute{
																	"certificate_chain": schema.StringAttribute{
																		Description:         "The certificate trust chain for a certificate stored on the file system of the virtual Gateway.",
																		MarkdownDescription: "The certificate trust chain for a certificate stored on the file system of the virtual Gateway.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.LengthAtLeast(1),
																			stringvalidator.LengthAtMost(255),
																		},
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"sds": schema.SingleNestedAttribute{
																Description:         "An object that represents a TLS validation context trust for a SDS certificate",
																MarkdownDescription: "An object that represents a TLS validation context trust for a SDS certificate",
																Attributes: map[string]schema.Attribute{
																	"secret_name": schema.StringAttribute{
																		Description:         "The certificate trust chain for a certificate issued via SDS.",
																		MarkdownDescription: "The certificate trust chain for a certificate issued via SDS.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
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
												},
												Required: true,
												Optional: false,
												Computed: false,
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"gateway_route_selector": schema.SingleNestedAttribute{
						Description:         "GatewayRouteSelector selects GatewayRoutes using labels to designate GatewayRoute membership. If not specified it selects all GatewayRoutes in that namespace.",
						MarkdownDescription: "GatewayRouteSelector selects GatewayRoutes using labels to designate GatewayRoute membership. If not specified it selects all GatewayRoutes in that namespace.",
						Attributes: map[string]schema.Attribute{
							"match_expressions": schema.ListNestedAttribute{
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "key is the label key that the selector applies to.",
											MarkdownDescription: "key is the label key that the selector applies to.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											ElementType:         types.StringType,
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

							"match_labels": schema.MapAttribute{
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

					"listeners": schema.ListNestedAttribute{
						Description:         "The listener that the virtual gateway is expected to receive inbound traffic from",
						MarkdownDescription: "The listener that the virtual gateway is expected to receive inbound traffic from",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"connection_pool": schema.SingleNestedAttribute{
									Description:         "The connection pool settings for the listener",
									MarkdownDescription: "The connection pool settings for the listener",
									Attributes: map[string]schema.Attribute{
										"grpc": schema.SingleNestedAttribute{
											Description:         "Specifies grpc connection pool settings for the virtual gateway listener",
											MarkdownDescription: "Specifies grpc connection pool settings for the virtual gateway listener",
											Attributes: map[string]schema.Attribute{
												"max_requests": schema.Int64Attribute{
													Description:         "Represents the maximum number of inflight requests that an envoy can concurrently support across all the hosts in the upstream cluster",
													MarkdownDescription: "Represents the maximum number of inflight requests that an envoy can concurrently support across all the hosts in the upstream cluster",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(1),
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"http": schema.SingleNestedAttribute{
											Description:         "Specifies http connection pool settings for the virtual gateway listener",
											MarkdownDescription: "Specifies http connection pool settings for the virtual gateway listener",
											Attributes: map[string]schema.Attribute{
												"max_connections": schema.Int64Attribute{
													Description:         "Represents the maximum number of outbound TCP connections the envoy can establish concurrently with all the hosts in the upstream cluster.",
													MarkdownDescription: "Represents the maximum number of outbound TCP connections the envoy can establish concurrently with all the hosts in the upstream cluster.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(1),
													},
												},

												"max_pending_requests": schema.Int64Attribute{
													Description:         "Represents the number of overflowing requests after max_connections that an envoy will queue to an upstream cluster.",
													MarkdownDescription: "Represents the number of overflowing requests after max_connections that an envoy will queue to an upstream cluster.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(1),
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"http2": schema.SingleNestedAttribute{
											Description:         "Specifies http2 connection pool settings for the virtual gateway listener",
											MarkdownDescription: "Specifies http2 connection pool settings for the virtual gateway listener",
											Attributes: map[string]schema.Attribute{
												"max_requests": schema.Int64Attribute{
													Description:         "Represents the maximum number of inflight requests that an envoy can concurrently support across all the hosts in the upstream cluster",
													MarkdownDescription: "Represents the maximum number of inflight requests that an envoy can concurrently support across all the hosts in the upstream cluster",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(1),
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

								"health_check": schema.SingleNestedAttribute{
									Description:         "The health check information for the listener.",
									MarkdownDescription: "The health check information for the listener.",
									Attributes: map[string]schema.Attribute{
										"healthy_threshold": schema.Int64Attribute{
											Description:         "The number of consecutive successful health checks that must occur before declaring listener healthy.",
											MarkdownDescription: "The number of consecutive successful health checks that must occur before declaring listener healthy.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(2),
												int64validator.AtMost(10),
											},
										},

										"interval_millis": schema.Int64Attribute{
											Description:         "The time period in milliseconds between each health check execution.",
											MarkdownDescription: "The time period in milliseconds between each health check execution.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(5000),
												int64validator.AtMost(300000),
											},
										},

										"path": schema.StringAttribute{
											Description:         "The destination path for the health check request. This value is only used if the specified protocol is http or http2. For any other protocol, this value is ignored.",
											MarkdownDescription: "The destination path for the health check request. This value is only used if the specified protocol is http or http2. For any other protocol, this value is ignored.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"port": schema.Int64Attribute{
											Description:         "The destination port for the health check request.",
											MarkdownDescription: "The destination port for the health check request.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(1),
												int64validator.AtMost(65535),
											},
										},

										"protocol": schema.StringAttribute{
											Description:         "The protocol for the health check request",
											MarkdownDescription: "The protocol for the health check request",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("grpc", "http", "http2"),
											},
										},

										"timeout_millis": schema.Int64Attribute{
											Description:         "The amount of time to wait when receiving a response from the health check, in milliseconds.",
											MarkdownDescription: "The amount of time to wait when receiving a response from the health check, in milliseconds.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(2000),
												int64validator.AtMost(60000),
											},
										},

										"unhealthy_threshold": schema.Int64Attribute{
											Description:         "The number of consecutive failed health checks that must occur before declaring a virtual Gateway unhealthy.",
											MarkdownDescription: "The number of consecutive failed health checks that must occur before declaring a virtual Gateway unhealthy.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(2),
												int64validator.AtMost(10),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"port_mapping": schema.SingleNestedAttribute{
									Description:         "The port mapping information for the listener.",
									MarkdownDescription: "The port mapping information for the listener.",
									Attributes: map[string]schema.Attribute{
										"port": schema.Int64Attribute{
											Description:         "The port used for the port mapping.",
											MarkdownDescription: "The port used for the port mapping.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(1),
												int64validator.AtMost(65535),
											},
										},

										"protocol": schema.StringAttribute{
											Description:         "The protocol used for the port mapping.",
											MarkdownDescription: "The protocol used for the port mapping.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("grpc", "http", "http2"),
											},
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"tls": schema.SingleNestedAttribute{
									Description:         "A reference to an object that represents the Transport Layer Security (TLS) properties for a listener.",
									MarkdownDescription: "A reference to an object that represents the Transport Layer Security (TLS) properties for a listener.",
									Attributes: map[string]schema.Attribute{
										"certificate": schema.SingleNestedAttribute{
											Description:         "A reference to an object that represents a listener's TLS certificate.",
											MarkdownDescription: "A reference to an object that represents a listener's TLS certificate.",
											Attributes: map[string]schema.Attribute{
												"acm": schema.SingleNestedAttribute{
													Description:         "A reference to an object that represents an AWS Certificate Manager (ACM) certificate.",
													MarkdownDescription: "A reference to an object that represents an AWS Certificate Manager (ACM) certificate.",
													Attributes: map[string]schema.Attribute{
														"certificate_arn": schema.StringAttribute{
															Description:         "The Amazon Resource Name (ARN) for the certificate.",
															MarkdownDescription: "The Amazon Resource Name (ARN) for the certificate.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"file": schema.SingleNestedAttribute{
													Description:         "A reference to an object that represents a local file certificate.",
													MarkdownDescription: "A reference to an object that represents a local file certificate.",
													Attributes: map[string]schema.Attribute{
														"certificate_chain": schema.StringAttribute{
															Description:         "The certificate chain for the certificate.",
															MarkdownDescription: "The certificate chain for the certificate.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtLeast(1),
																stringvalidator.LengthAtMost(255),
															},
														},

														"private_key": schema.StringAttribute{
															Description:         "The private key for a certificate stored on the file system of the virtual Gateway.",
															MarkdownDescription: "The private key for a certificate stored on the file system of the virtual Gateway.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtLeast(1),
																stringvalidator.LengthAtMost(255),
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"sds": schema.SingleNestedAttribute{
													Description:         "A reference to an object that represents an SDS issued certificate",
													MarkdownDescription: "A reference to an object that represents an SDS issued certificate",
													Attributes: map[string]schema.Attribute{
														"secret_name": schema.StringAttribute{
															Description:         "The certificate trust chain for a certificate issued via SDS cluster",
															MarkdownDescription: "The certificate trust chain for a certificate issued via SDS cluster",
															Required:            true,
															Optional:            false,
															Computed:            false,
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

										"mode": schema.StringAttribute{
											Description:         "ListenerTLS mode",
											MarkdownDescription: "ListenerTLS mode",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("DISABLED", "PERMISSIVE", "STRICT"),
											},
										},

										"validation": schema.SingleNestedAttribute{
											Description:         "A reference to an object that represents Validation context",
											MarkdownDescription: "A reference to an object that represents Validation context",
											Attributes: map[string]schema.Attribute{
												"subject_alternative_names": schema.SingleNestedAttribute{
													Description:         "Possible alternate names to consider",
													MarkdownDescription: "Possible alternate names to consider",
													Attributes: map[string]schema.Attribute{
														"match": schema.SingleNestedAttribute{
															Description:         "Match is a required field",
															MarkdownDescription: "Match is a required field",
															Attributes: map[string]schema.Attribute{
																"exact": schema.ListAttribute{
																	Description:         "Exact is a required field",
																	MarkdownDescription: "Exact is a required field",
																	ElementType:         types.StringType,
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"trust": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"acm": schema.SingleNestedAttribute{
															Description:         "A reference to an object that represents a TLS validation context trust for an AWS Certicate Manager (ACM) certificate.",
															MarkdownDescription: "A reference to an object that represents a TLS validation context trust for an AWS Certicate Manager (ACM) certificate.",
															Attributes: map[string]schema.Attribute{
																"certificate_authority_ar_ns": schema.ListAttribute{
																	Description:         "One or more ACM Amazon Resource Name (ARN)s.",
																	MarkdownDescription: "One or more ACM Amazon Resource Name (ARN)s.",
																	ElementType:         types.StringType,
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"file": schema.SingleNestedAttribute{
															Description:         "An object that represents a TLS validation context trust for a local file.",
															MarkdownDescription: "An object that represents a TLS validation context trust for a local file.",
															Attributes: map[string]schema.Attribute{
																"certificate_chain": schema.StringAttribute{
																	Description:         "The certificate trust chain for a certificate stored on the file system of the virtual Gateway.",
																	MarkdownDescription: "The certificate trust chain for a certificate stored on the file system of the virtual Gateway.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																		stringvalidator.LengthAtMost(255),
																	},
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"sds": schema.SingleNestedAttribute{
															Description:         "An object that represents a TLS validation context trust for an SDS system",
															MarkdownDescription: "An object that represents a TLS validation context trust for an SDS system",
															Attributes: map[string]schema.Attribute{
																"secret_name": schema.StringAttribute{
																	Description:         "The certificate trust chain for a certificate issued via SDS.",
																	MarkdownDescription: "The certificate trust chain for a certificate issued via SDS.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
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

					"logging": schema.SingleNestedAttribute{
						Description:         "The inbound and outbound access logging information for the virtual gateway.",
						MarkdownDescription: "The inbound and outbound access logging information for the virtual gateway.",
						Attributes: map[string]schema.Attribute{
							"access_log": schema.SingleNestedAttribute{
								Description:         "The access log configuration for a virtual Gateway.",
								MarkdownDescription: "The access log configuration for a virtual Gateway.",
								Attributes: map[string]schema.Attribute{
									"file": schema.SingleNestedAttribute{
										Description:         "The file object to send virtual gateway access logs to.",
										MarkdownDescription: "The file object to send virtual gateway access logs to.",
										Attributes: map[string]schema.Attribute{
											"format": schema.SingleNestedAttribute{
												Description:         "Structured access log output format",
												MarkdownDescription: "Structured access log output format",
												Attributes: map[string]schema.Attribute{
													"json": schema.ListNestedAttribute{
														Description:         "Output specified fields as a JSON object",
														MarkdownDescription: "Output specified fields as a JSON object",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "The name of the field in the JSON object",
																	MarkdownDescription: "The name of the field in the JSON object",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																	},
																},

																"value": schema.StringAttribute{
																	Description:         "The format string",
																	MarkdownDescription: "The format string",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																	},
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"text": schema.StringAttribute{
														Description:         "Custom format string",
														MarkdownDescription: "Custom format string",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"path": schema.StringAttribute{
												Description:         "The file path to write access logs to.",
												MarkdownDescription: "The file path to write access logs to.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(255),
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"mesh_ref": schema.SingleNestedAttribute{
						Description:         "A reference to k8s Mesh CR that this VirtualGateway belongs to. The admission controller populates it using Meshes's selector, and prevents users from setting this field.  Populated by the system. Read-only.",
						MarkdownDescription: "A reference to k8s Mesh CR that this VirtualGateway belongs to. The admission controller populates it using Meshes's selector, and prevents users from setting this field.  Populated by the system. Read-only.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name is the name of Mesh CR",
								MarkdownDescription: "Name is the name of Mesh CR",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"uid": schema.StringAttribute{
								Description:         "UID is the UID of Mesh CR",
								MarkdownDescription: "UID is the UID of Mesh CR",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"namespace_selector": schema.SingleNestedAttribute{
						Description:         "NamespaceSelector selects Namespaces using labels to designate GatewayRoute membership. This field follows standard label selector semantics; if present but empty, it selects all namespaces.",
						MarkdownDescription: "NamespaceSelector selects Namespaces using labels to designate GatewayRoute membership. This field follows standard label selector semantics; if present but empty, it selects all namespaces.",
						Attributes: map[string]schema.Attribute{
							"match_expressions": schema.ListNestedAttribute{
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "key is the label key that the selector applies to.",
											MarkdownDescription: "key is the label key that the selector applies to.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											ElementType:         types.StringType,
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

							"match_labels": schema.MapAttribute{
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

					"pod_selector": schema.SingleNestedAttribute{
						Description:         "PodSelector selects Pods using labels to designate VirtualGateway membership. This field follows standard label selector semantics: 	if present but empty, it selects all pods within namespace. 	if absent, it selects no pod.",
						MarkdownDescription: "PodSelector selects Pods using labels to designate VirtualGateway membership. This field follows standard label selector semantics: 	if present but empty, it selects all pods within namespace. 	if absent, it selects no pod.",
						Attributes: map[string]schema.Attribute{
							"match_expressions": schema.ListNestedAttribute{
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "key is the label key that the selector applies to.",
											MarkdownDescription: "key is the label key that the selector applies to.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											ElementType:         types.StringType,
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

							"match_labels": schema.MapAttribute{
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *AppmeshK8SAwsVirtualGatewayV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_appmesh_k8s_aws_virtual_gateway_v1beta2_manifest")

	var model AppmeshK8SAwsVirtualGatewayV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("appmesh.k8s.aws/v1beta2")
	model.Kind = pointer.String("VirtualGateway")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
