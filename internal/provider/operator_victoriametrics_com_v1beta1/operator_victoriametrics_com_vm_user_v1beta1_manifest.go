/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package operator_victoriametrics_com_v1beta1

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
	_ datasource.DataSource = &OperatorVictoriametricsComVmuserV1Beta1Manifest{}
)

func NewOperatorVictoriametricsComVmuserV1Beta1Manifest() datasource.DataSource {
	return &OperatorVictoriametricsComVmuserV1Beta1Manifest{}
}

type OperatorVictoriametricsComVmuserV1Beta1Manifest struct{}

type OperatorVictoriametricsComVmuserV1Beta1ManifestData struct {
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
		BearerToken                *string   `tfsdk:"bearer_token" json:"bearerToken,omitempty"`
		Default_url                *[]string `tfsdk:"default_url" json:"default_url,omitempty"`
		Disable_secret_creation    *bool     `tfsdk:"disable_secret_creation" json:"disable_secret_creation,omitempty"`
		Drop_src_path_prefix_parts *int64    `tfsdk:"drop_src_path_prefix_parts" json:"drop_src_path_prefix_parts,omitempty"`
		GeneratePassword           *bool     `tfsdk:"generate_password" json:"generatePassword,omitempty"`
		Headers                    *[]string `tfsdk:"headers" json:"headers,omitempty"`
		Ip_filters                 *struct {
			Allow_list *[]string `tfsdk:"allow_list" json:"allow_list,omitempty"`
			Deny_list  *[]string `tfsdk:"deny_list" json:"deny_list,omitempty"`
		} `tfsdk:"ip_filters" json:"ip_filters,omitempty"`
		Load_balancing_policy   *string            `tfsdk:"load_balancing_policy" json:"load_balancing_policy,omitempty"`
		Max_concurrent_requests *int64             `tfsdk:"max_concurrent_requests" json:"max_concurrent_requests,omitempty"`
		Metric_labels           *map[string]string `tfsdk:"metric_labels" json:"metric_labels,omitempty"`
		Name                    *string            `tfsdk:"name" json:"name,omitempty"`
		Password                *string            `tfsdk:"password" json:"password,omitempty"`
		PasswordRef             *struct {
			Key      *string `tfsdk:"key" json:"key,omitempty"`
			Name     *string `tfsdk:"name" json:"name,omitempty"`
			Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
		} `tfsdk:"password_ref" json:"passwordRef,omitempty"`
		Response_headers   *[]string `tfsdk:"response_headers" json:"response_headers,omitempty"`
		Retry_status_codes *[]string `tfsdk:"retry_status_codes" json:"retry_status_codes,omitempty"`
		TargetRefs         *[]struct {
			Crd *struct {
				Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"crd" json:"crd,omitempty"`
			Drop_src_path_prefix_parts *int64    `tfsdk:"drop_src_path_prefix_parts" json:"drop_src_path_prefix_parts,omitempty"`
			Headers                    *[]string `tfsdk:"headers" json:"headers,omitempty"`
			Hosts                      *[]string `tfsdk:"hosts" json:"hosts,omitempty"`
			Load_balancing_policy      *string   `tfsdk:"load_balancing_policy" json:"load_balancing_policy,omitempty"`
			Paths                      *[]string `tfsdk:"paths" json:"paths,omitempty"`
			Response_headers           *[]string `tfsdk:"response_headers" json:"response_headers,omitempty"`
			Retry_status_codes         *[]string `tfsdk:"retry_status_codes" json:"retry_status_codes,omitempty"`
			Static                     *struct {
				Url  *string   `tfsdk:"url" json:"url,omitempty"`
				Urls *[]string `tfsdk:"urls" json:"urls,omitempty"`
			} `tfsdk:"static" json:"static,omitempty"`
			TargetRefBasicAuth *struct {
				Password *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"password" json:"password,omitempty"`
				Username *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"username" json:"username,omitempty"`
			} `tfsdk:"target_ref_basic_auth" json:"targetRefBasicAuth,omitempty"`
			Target_path_suffix *string `tfsdk:"target_path_suffix" json:"target_path_suffix,omitempty"`
		} `tfsdk:"target_refs" json:"targetRefs,omitempty"`
		Tls_insecure_skip_verify *bool `tfsdk:"tls_insecure_skip_verify" json:"tls_insecure_skip_verify,omitempty"`
		TokenRef                 *struct {
			Key      *string `tfsdk:"key" json:"key,omitempty"`
			Name     *string `tfsdk:"name" json:"name,omitempty"`
			Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
		} `tfsdk:"token_ref" json:"tokenRef,omitempty"`
		Username *string `tfsdk:"username" json:"username,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OperatorVictoriametricsComVmuserV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operator_victoriametrics_com_vm_user_v1beta1_manifest"
}

func (r *OperatorVictoriametricsComVmuserV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VMUser is the Schema for the vmusers API",
		MarkdownDescription: "VMUser is the Schema for the vmusers API",
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
				Description:         "VMUserSpec defines the desired state of VMUser",
				MarkdownDescription: "VMUserSpec defines the desired state of VMUser",
				Attributes: map[string]schema.Attribute{
					"bearer_token": schema.StringAttribute{
						Description:         "BearerToken Authorization header value for accessing protected endpoint.",
						MarkdownDescription: "BearerToken Authorization header value for accessing protected endpoint.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"default_url": schema.ListAttribute{
						Description:         "DefaultURLs backend url for non-matching paths filterusually used for default backend with error message",
						MarkdownDescription: "DefaultURLs backend url for non-matching paths filterusually used for default backend with error message",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"disable_secret_creation": schema.BoolAttribute{
						Description:         "DisableSecretCreation skips related secret creation for vmuser",
						MarkdownDescription: "DisableSecretCreation skips related secret creation for vmuser",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"drop_src_path_prefix_parts": schema.Int64Attribute{
						Description:         "DropSrcPathPrefixParts is the number of '/'-delimited request path prefix parts to drop before proxying the request to backend.See https://docs.victoriametrics.com/vmauth.html#dropping-request-path-prefix for more details.",
						MarkdownDescription: "DropSrcPathPrefixParts is the number of '/'-delimited request path prefix parts to drop before proxying the request to backend.See https://docs.victoriametrics.com/vmauth.html#dropping-request-path-prefix for more details.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"generate_password": schema.BoolAttribute{
						Description:         "GeneratePassword instructs operator to generate password for userif spec.password if empty.",
						MarkdownDescription: "GeneratePassword instructs operator to generate password for userif spec.password if empty.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"headers": schema.ListAttribute{
						Description:         "Headers represent additional http headers, that vmauth usesin form of ['header_key: header_value']multiple values for header key:['header_key: value1,value2']it's available since 1.68.0 version of vmauth",
						MarkdownDescription: "Headers represent additional http headers, that vmauth usesin form of ['header_key: header_value']multiple values for header key:['header_key: value1,value2']it's available since 1.68.0 version of vmauth",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ip_filters": schema.SingleNestedAttribute{
						Description:         "IPFilters defines per target src ip filterssupported only with enterprise version of vmauthhttps://docs.victoriametrics.com/vmauth.html#ip-filters",
						MarkdownDescription: "IPFilters defines per target src ip filterssupported only with enterprise version of vmauthhttps://docs.victoriametrics.com/vmauth.html#ip-filters",
						Attributes: map[string]schema.Attribute{
							"allow_list": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"deny_list": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
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

					"load_balancing_policy": schema.StringAttribute{
						Description:         "LoadBalancingPolicy defines load balancing policy to use for backend urls.Supported policies: least_loaded, first_available.See https://docs.victoriametrics.com/vmauth.html#load-balancing for more details (default 'least_loaded')",
						MarkdownDescription: "LoadBalancingPolicy defines load balancing policy to use for backend urls.Supported policies: least_loaded, first_available.See https://docs.victoriametrics.com/vmauth.html#load-balancing for more details (default 'least_loaded')",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("least_loaded", "first_available"),
						},
					},

					"max_concurrent_requests": schema.Int64Attribute{
						Description:         "MaxConcurrentRequests defines max concurrent requests per user300 is default value for vmauth",
						MarkdownDescription: "MaxConcurrentRequests defines max concurrent requests per user300 is default value for vmauth",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"metric_labels": schema.MapAttribute{
						Description:         "MetricLabels - additional labels for metrics exported by vmauth for given user.",
						MarkdownDescription: "MetricLabels - additional labels for metrics exported by vmauth for given user.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "Name of the VMUser object.",
						MarkdownDescription: "Name of the VMUser object.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"password": schema.StringAttribute{
						Description:         "Password basic auth password for accessing protected endpoint.",
						MarkdownDescription: "Password basic auth password for accessing protected endpoint.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"password_ref": schema.SingleNestedAttribute{
						Description:         "PasswordRef allows fetching password from user-create secret by its name and key.",
						MarkdownDescription: "PasswordRef allows fetching password from user-create secret by its name and key.",
						Attributes: map[string]schema.Attribute{
							"key": schema.StringAttribute{
								Description:         "The key of the secret to select from.  Must be a valid secret key.",
								MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
								MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"optional": schema.BoolAttribute{
								Description:         "Specify whether the Secret or its key must be defined",
								MarkdownDescription: "Specify whether the Secret or its key must be defined",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"response_headers": schema.ListAttribute{
						Description:         "ResponseHeaders represent additional http headers, that vmauth adds for request responsein form of ['header_key: header_value']multiple values for header key:['header_key: value1,value2']it's available since 1.93.0 version of vmauth",
						MarkdownDescription: "ResponseHeaders represent additional http headers, that vmauth adds for request responsein form of ['header_key: header_value']multiple values for header key:['header_key: value1,value2']it's available since 1.93.0 version of vmauth",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"retry_status_codes": schema.ListAttribute{
						Description:         "RetryStatusCodes defines http status codes in numeric format for request retriese.g. [429,503]",
						MarkdownDescription: "RetryStatusCodes defines http status codes in numeric format for request retriese.g. [429,503]",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"target_refs": schema.ListNestedAttribute{
						Description:         "TargetRefs - reference to endpoints, which user may access.",
						MarkdownDescription: "TargetRefs - reference to endpoints, which user may access.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"crd": schema.SingleNestedAttribute{
									Description:         "CRD describes exist operator's CRD object,operator generates access url based on CRD params.",
									MarkdownDescription: "CRD describes exist operator's CRD object,operator generates access url based on CRD params.",
									Attributes: map[string]schema.Attribute{
										"kind": schema.StringAttribute{
											Description:         "Kind one of:VMAgent VMAlert VMCluster VMSingle or VMAlertManager",
											MarkdownDescription: "Kind one of:VMAgent VMAlert VMCluster VMSingle or VMAlertManager",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name target CRD object name",
											MarkdownDescription: "Name target CRD object name",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"namespace": schema.StringAttribute{
											Description:         "Namespace target CRD object namespace.",
											MarkdownDescription: "Namespace target CRD object namespace.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"drop_src_path_prefix_parts": schema.Int64Attribute{
									Description:         "DropSrcPathPrefixParts is the number of '/'-delimited request path prefix parts to drop before proxying the request to backend.See https://docs.victoriametrics.com/vmauth.html#dropping-request-path-prefix for more details.",
									MarkdownDescription: "DropSrcPathPrefixParts is the number of '/'-delimited request path prefix parts to drop before proxying the request to backend.See https://docs.victoriametrics.com/vmauth.html#dropping-request-path-prefix for more details.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"headers": schema.ListAttribute{
									Description:         "Headers represent additional http headers, that vmauth usesin form of ['header_key: header_value']multiple values for header key:['header_key: value1,value2']it's available since 1.68.0 version of vmauth",
									MarkdownDescription: "Headers represent additional http headers, that vmauth usesin form of ['header_key: header_value']multiple values for header key:['header_key: value1,value2']it's available since 1.68.0 version of vmauth",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"hosts": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"load_balancing_policy": schema.StringAttribute{
									Description:         "LoadBalancingPolicy defines load balancing policy to use for backend urls.Supported policies: least_loaded, first_available.See https://docs.victoriametrics.com/vmauth.html#load-balancing for more details (default 'least_loaded')",
									MarkdownDescription: "LoadBalancingPolicy defines load balancing policy to use for backend urls.Supported policies: least_loaded, first_available.See https://docs.victoriametrics.com/vmauth.html#load-balancing for more details (default 'least_loaded')",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("least_loaded", "first_available"),
									},
								},

								"paths": schema.ListAttribute{
									Description:         "Paths - matched path to route.",
									MarkdownDescription: "Paths - matched path to route.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"response_headers": schema.ListAttribute{
									Description:         "ResponseHeaders represent additional http headers, that vmauth adds for request responsein form of ['header_key: header_value']multiple values for header key:['header_key: value1,value2']it's available since 1.93.0 version of vmauth",
									MarkdownDescription: "ResponseHeaders represent additional http headers, that vmauth adds for request responsein form of ['header_key: header_value']multiple values for header key:['header_key: value1,value2']it's available since 1.93.0 version of vmauth",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"retry_status_codes": schema.ListAttribute{
									Description:         "RetryStatusCodes defines http status codes in numeric format for request retriesCan be defined per target or at VMUser.spec levele.g. [429,503]",
									MarkdownDescription: "RetryStatusCodes defines http status codes in numeric format for request retriesCan be defined per target or at VMUser.spec levele.g. [429,503]",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"static": schema.SingleNestedAttribute{
									Description:         "Static - user defined url for traffic forward,for instance http://vmsingle:8429",
									MarkdownDescription: "Static - user defined url for traffic forward,for instance http://vmsingle:8429",
									Attributes: map[string]schema.Attribute{
										"url": schema.StringAttribute{
											Description:         "URL http url for given staticRef.",
											MarkdownDescription: "URL http url for given staticRef.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"urls": schema.ListAttribute{
											Description:         "URLs allows setting multiple urls for load-balancing at vmauth-side.",
											MarkdownDescription: "URLs allows setting multiple urls for load-balancing at vmauth-side.",
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

								"target_ref_basic_auth": schema.SingleNestedAttribute{
									Description:         "TargetRefBasicAuth allow an target endpoint to authenticate over basic authentication",
									MarkdownDescription: "TargetRefBasicAuth allow an target endpoint to authenticate over basic authentication",
									Attributes: map[string]schema.Attribute{
										"password": schema.SingleNestedAttribute{
											Description:         "The secret in the service scrape namespace that contains the passwordfor authentication.It must be at them same namespace as CRD",
											MarkdownDescription: "The secret in the service scrape namespace that contains the passwordfor authentication.It must be at them same namespace as CRD",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "The key of the secret to select from.  Must be a valid secret key.",
													MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"optional": schema.BoolAttribute{
													Description:         "Specify whether the Secret or its key must be defined",
													MarkdownDescription: "Specify whether the Secret or its key must be defined",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"username": schema.SingleNestedAttribute{
											Description:         "The secret in the service scrape namespace that contains the usernamefor authentication.It must be at them same namespace as CRD",
											MarkdownDescription: "The secret in the service scrape namespace that contains the usernamefor authentication.It must be at them same namespace as CRD",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "The key of the secret to select from.  Must be a valid secret key.",
													MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"optional": schema.BoolAttribute{
													Description:         "Specify whether the Secret or its key must be defined",
													MarkdownDescription: "Specify whether the Secret or its key must be defined",
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
									Required: false,
									Optional: true,
									Computed: false,
								},

								"target_path_suffix": schema.StringAttribute{
									Description:         "QueryParams []string 'json:'queryParams,omitempty''TargetPathSuffix allows to add some suffix to the target pathIt allows to hide tenant configuration from user with crd as ref.it also may contain any url encoded params.",
									MarkdownDescription: "QueryParams []string 'json:'queryParams,omitempty''TargetPathSuffix allows to add some suffix to the target pathIt allows to hide tenant configuration from user with crd as ref.it also may contain any url encoded params.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"tls_insecure_skip_verify": schema.BoolAttribute{
						Description:         "TLSInsecureSkipVerify - whether to skip TLS verification when connecting to backend over HTTPS.See https://docs.victoriametrics.com/vmauth.html#backend-tls-setup",
						MarkdownDescription: "TLSInsecureSkipVerify - whether to skip TLS verification when connecting to backend over HTTPS.See https://docs.victoriametrics.com/vmauth.html#backend-tls-setup",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"token_ref": schema.SingleNestedAttribute{
						Description:         "TokenRef allows fetching token from user-created secrets by its name and key.",
						MarkdownDescription: "TokenRef allows fetching token from user-created secrets by its name and key.",
						Attributes: map[string]schema.Attribute{
							"key": schema.StringAttribute{
								Description:         "The key of the secret to select from.  Must be a valid secret key.",
								MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
								MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"optional": schema.BoolAttribute{
								Description:         "Specify whether the Secret or its key must be defined",
								MarkdownDescription: "Specify whether the Secret or its key must be defined",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"username": schema.StringAttribute{
						Description:         "UserName basic auth user name for accessing protected endpoint,will be replaced with metadata.name of VMUser if omitted.",
						MarkdownDescription: "UserName basic auth user name for accessing protected endpoint,will be replaced with metadata.name of VMUser if omitted.",
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
	}
}

func (r *OperatorVictoriametricsComVmuserV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_operator_victoriametrics_com_vm_user_v1beta1_manifest")

	var model OperatorVictoriametricsComVmuserV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("operator.victoriametrics.com/v1beta1")
	model.Kind = pointer.String("VMUser")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
