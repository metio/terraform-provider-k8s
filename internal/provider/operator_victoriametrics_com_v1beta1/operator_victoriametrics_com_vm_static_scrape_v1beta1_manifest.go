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
	_ datasource.DataSource = &OperatorVictoriametricsComVmstaticScrapeV1Beta1Manifest{}
)

func NewOperatorVictoriametricsComVmstaticScrapeV1Beta1Manifest() datasource.DataSource {
	return &OperatorVictoriametricsComVmstaticScrapeV1Beta1Manifest{}
}

type OperatorVictoriametricsComVmstaticScrapeV1Beta1Manifest struct{}

type OperatorVictoriametricsComVmstaticScrapeV1Beta1ManifestData struct {
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
		JobName         *string `tfsdk:"job_name" json:"jobName,omitempty"`
		SampleLimit     *int64  `tfsdk:"sample_limit" json:"sampleLimit,omitempty"`
		TargetEndpoints *[]struct {
			Authorization *struct {
				Credentials *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"credentials" json:"credentials,omitempty"`
				CredentialsFile *string `tfsdk:"credentials_file" json:"credentialsFile,omitempty"`
				Type            *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"authorization" json:"authorization,omitempty"`
			BasicAuth *struct {
				Password *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"password" json:"password,omitempty"`
				Password_file *string `tfsdk:"password_file" json:"password_file,omitempty"`
				Username      *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"username" json:"username,omitempty"`
			} `tfsdk:"basic_auth" json:"basicAuth,omitempty"`
			BearerTokenFile   *string `tfsdk:"bearer_token_file" json:"bearerTokenFile,omitempty"`
			BearerTokenSecret *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"bearer_token_secret" json:"bearerTokenSecret,omitempty"`
			Follow_redirects     *bool              `tfsdk:"follow_redirects" json:"follow_redirects,omitempty"`
			HonorLabels          *bool              `tfsdk:"honor_labels" json:"honorLabels,omitempty"`
			HonorTimestamps      *bool              `tfsdk:"honor_timestamps" json:"honorTimestamps,omitempty"`
			Interval             *string            `tfsdk:"interval" json:"interval,omitempty"`
			Labels               *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			MetricRelabelConfigs *[]struct {
				Action       *string            `tfsdk:"action" json:"action,omitempty"`
				If           *map[string]string `tfsdk:"if" json:"if,omitempty"`
				Labels       *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				Match        *string            `tfsdk:"match" json:"match,omitempty"`
				Modulus      *int64             `tfsdk:"modulus" json:"modulus,omitempty"`
				Regex        *string            `tfsdk:"regex" json:"regex,omitempty"`
				Replacement  *string            `tfsdk:"replacement" json:"replacement,omitempty"`
				Separator    *string            `tfsdk:"separator" json:"separator,omitempty"`
				SourceLabels *[]string          `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
				TargetLabel  *string            `tfsdk:"target_label" json:"targetLabel,omitempty"`
			} `tfsdk:"metric_relabel_configs" json:"metricRelabelConfigs,omitempty"`
			Oauth2 *struct {
				Client_id *struct {
					ConfigMap *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"config_map" json:"configMap,omitempty"`
					Secret *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"secret" json:"secret,omitempty"`
				} `tfsdk:"client_id" json:"client_id,omitempty"`
				Client_secret *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"client_secret" json:"client_secret,omitempty"`
				Client_secret_file *string            `tfsdk:"client_secret_file" json:"client_secret_file,omitempty"`
				Endpoint_params    *map[string]string `tfsdk:"endpoint_params" json:"endpoint_params,omitempty"`
				Scopes             *[]string          `tfsdk:"scopes" json:"scopes,omitempty"`
				Token_url          *string            `tfsdk:"token_url" json:"token_url,omitempty"`
			} `tfsdk:"oauth2" json:"oauth2,omitempty"`
			Params         *map[string][]string `tfsdk:"params" json:"params,omitempty"`
			Path           *string              `tfsdk:"path" json:"path,omitempty"`
			Port           *string              `tfsdk:"port" json:"port,omitempty"`
			ProxyURL       *string              `tfsdk:"proxy_url" json:"proxyURL,omitempty"`
			RelabelConfigs *[]struct {
				Action       *string            `tfsdk:"action" json:"action,omitempty"`
				If           *map[string]string `tfsdk:"if" json:"if,omitempty"`
				Labels       *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				Match        *string            `tfsdk:"match" json:"match,omitempty"`
				Modulus      *int64             `tfsdk:"modulus" json:"modulus,omitempty"`
				Regex        *string            `tfsdk:"regex" json:"regex,omitempty"`
				Replacement  *string            `tfsdk:"replacement" json:"replacement,omitempty"`
				Separator    *string            `tfsdk:"separator" json:"separator,omitempty"`
				SourceLabels *[]string          `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
				TargetLabel  *string            `tfsdk:"target_label" json:"targetLabel,omitempty"`
			} `tfsdk:"relabel_configs" json:"relabelConfigs,omitempty"`
			SampleLimit     *int64    `tfsdk:"sample_limit" json:"sampleLimit,omitempty"`
			Scheme          *string   `tfsdk:"scheme" json:"scheme,omitempty"`
			ScrapeTimeout   *string   `tfsdk:"scrape_timeout" json:"scrapeTimeout,omitempty"`
			Scrape_interval *string   `tfsdk:"scrape_interval" json:"scrape_interval,omitempty"`
			Targets         *[]string `tfsdk:"targets" json:"targets,omitempty"`
			TlsConfig       *struct {
				Ca *struct {
					ConfigMap *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"config_map" json:"configMap,omitempty"`
					Secret *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"secret" json:"secret,omitempty"`
				} `tfsdk:"ca" json:"ca,omitempty"`
				CaFile *string `tfsdk:"ca_file" json:"caFile,omitempty"`
				Cert   *struct {
					ConfigMap *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"config_map" json:"configMap,omitempty"`
					Secret *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"secret" json:"secret,omitempty"`
				} `tfsdk:"cert" json:"cert,omitempty"`
				CertFile           *string `tfsdk:"cert_file" json:"certFile,omitempty"`
				InsecureSkipVerify *bool   `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
				KeyFile            *string `tfsdk:"key_file" json:"keyFile,omitempty"`
				KeySecret          *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"key_secret" json:"keySecret,omitempty"`
				ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
			} `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
			Vm_scrape_params *struct {
				Disable_compression  *bool     `tfsdk:"disable_compression" json:"disable_compression,omitempty"`
				Disable_keep_alive   *bool     `tfsdk:"disable_keep_alive" json:"disable_keep_alive,omitempty"`
				Headers              *[]string `tfsdk:"headers" json:"headers,omitempty"`
				Metric_relabel_debug *bool     `tfsdk:"metric_relabel_debug" json:"metric_relabel_debug,omitempty"`
				No_stale_markers     *bool     `tfsdk:"no_stale_markers" json:"no_stale_markers,omitempty"`
				Proxy_client_config  *struct {
					Basic_auth *struct {
						Password *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"password" json:"password,omitempty"`
						Password_file *string `tfsdk:"password_file" json:"password_file,omitempty"`
						Username      *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"username" json:"username,omitempty"`
					} `tfsdk:"basic_auth" json:"basic_auth,omitempty"`
					Bearer_token *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"bearer_token" json:"bearer_token,omitempty"`
					Bearer_token_file *string `tfsdk:"bearer_token_file" json:"bearer_token_file,omitempty"`
					Tls_config        *struct {
						Ca *struct {
							ConfigMap *struct {
								Key      *string `tfsdk:"key" json:"key,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
							} `tfsdk:"config_map" json:"configMap,omitempty"`
							Secret *struct {
								Key      *string `tfsdk:"key" json:"key,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
							} `tfsdk:"secret" json:"secret,omitempty"`
						} `tfsdk:"ca" json:"ca,omitempty"`
						CaFile *string `tfsdk:"ca_file" json:"caFile,omitempty"`
						Cert   *struct {
							ConfigMap *struct {
								Key      *string `tfsdk:"key" json:"key,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
							} `tfsdk:"config_map" json:"configMap,omitempty"`
							Secret *struct {
								Key      *string `tfsdk:"key" json:"key,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
							} `tfsdk:"secret" json:"secret,omitempty"`
						} `tfsdk:"cert" json:"cert,omitempty"`
						CertFile           *string `tfsdk:"cert_file" json:"certFile,omitempty"`
						InsecureSkipVerify *bool   `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
						KeyFile            *string `tfsdk:"key_file" json:"keyFile,omitempty"`
						KeySecret          *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"key_secret" json:"keySecret,omitempty"`
						ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
					} `tfsdk:"tls_config" json:"tls_config,omitempty"`
				} `tfsdk:"proxy_client_config" json:"proxy_client_config,omitempty"`
				Relabel_debug         *bool   `tfsdk:"relabel_debug" json:"relabel_debug,omitempty"`
				Scrape_align_interval *string `tfsdk:"scrape_align_interval" json:"scrape_align_interval,omitempty"`
				Scrape_offset         *string `tfsdk:"scrape_offset" json:"scrape_offset,omitempty"`
				Stream_parse          *bool   `tfsdk:"stream_parse" json:"stream_parse,omitempty"`
			} `tfsdk:"vm_scrape_params" json:"vm_scrape_params,omitempty"`
		} `tfsdk:"target_endpoints" json:"targetEndpoints,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OperatorVictoriametricsComVmstaticScrapeV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operator_victoriametrics_com_vm_static_scrape_v1beta1_manifest"
}

func (r *OperatorVictoriametricsComVmstaticScrapeV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VMStaticScrape  defines static targets configuration for scraping.",
		MarkdownDescription: "VMStaticScrape  defines static targets configuration for scraping.",
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
				Description:         "VMStaticScrapeSpec defines the desired state of VMStaticScrape.",
				MarkdownDescription: "VMStaticScrapeSpec defines the desired state of VMStaticScrape.",
				Attributes: map[string]schema.Attribute{
					"job_name": schema.StringAttribute{
						Description:         "JobName name of job.",
						MarkdownDescription: "JobName name of job.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"sample_limit": schema.Int64Attribute{
						Description:         "SampleLimit defines per-scrape limit on number of scraped samples that will be accepted.",
						MarkdownDescription: "SampleLimit defines per-scrape limit on number of scraped samples that will be accepted.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"target_endpoints": schema.ListNestedAttribute{
						Description:         "A list of target endpoints to scrape metrics from.",
						MarkdownDescription: "A list of target endpoints to scrape metrics from.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"authorization": schema.SingleNestedAttribute{
									Description:         "Authorization with http header Authorization",
									MarkdownDescription: "Authorization with http header Authorization",
									Attributes: map[string]schema.Attribute{
										"credentials": schema.SingleNestedAttribute{
											Description:         "Reference to the secret with value for authorization",
											MarkdownDescription: "Reference to the secret with value for authorization",
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

										"credentials_file": schema.StringAttribute{
											Description:         "File with value for authorization",
											MarkdownDescription: "File with value for authorization",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"type": schema.StringAttribute{
											Description:         "Type of authorization, default to bearer",
											MarkdownDescription: "Type of authorization, default to bearer",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"basic_auth": schema.SingleNestedAttribute{
									Description:         "BasicAuth allow an endpoint to authenticate over basic authenticationMore info: https://prometheus.io/docs/operating/configuration/#endpoints",
									MarkdownDescription: "BasicAuth allow an endpoint to authenticate over basic authenticationMore info: https://prometheus.io/docs/operating/configuration/#endpoints",
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
											Required: false,
											Optional: true,
											Computed: false,
										},

										"password_file": schema.StringAttribute{
											Description:         "PasswordFile defines path to password file at disk",
											MarkdownDescription: "PasswordFile defines path to password file at disk",
											Required:            false,
											Optional:            true,
											Computed:            false,
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
											Required: false,
											Optional: true,
											Computed: false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"bearer_token_file": schema.StringAttribute{
									Description:         "File to read bearer token for scraping targets.",
									MarkdownDescription: "File to read bearer token for scraping targets.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"bearer_token_secret": schema.SingleNestedAttribute{
									Description:         "Secret to mount to read bearer token for scraping targets. The secretneeds to be in the same namespace as the service scrape and accessible bythe victoria-metrics operator.",
									MarkdownDescription: "Secret to mount to read bearer token for scraping targets. The secretneeds to be in the same namespace as the service scrape and accessible bythe victoria-metrics operator.",
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

								"follow_redirects": schema.BoolAttribute{
									Description:         "FollowRedirects controls redirects for scraping.",
									MarkdownDescription: "FollowRedirects controls redirects for scraping.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"honor_labels": schema.BoolAttribute{
									Description:         "HonorLabels chooses the metric's labels on collisions with target labels.",
									MarkdownDescription: "HonorLabels chooses the metric's labels on collisions with target labels.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"honor_timestamps": schema.BoolAttribute{
									Description:         "HonorTimestamps controls whether vmagent respects the timestamps present in scraped data.",
									MarkdownDescription: "HonorTimestamps controls whether vmagent respects the timestamps present in scraped data.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"interval": schema.StringAttribute{
									Description:         "Interval at which metrics should be scraped",
									MarkdownDescription: "Interval at which metrics should be scraped",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels static labels for targets.",
									MarkdownDescription: "Labels static labels for targets.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metric_relabel_configs": schema.ListNestedAttribute{
									Description:         "MetricRelabelConfigs to apply to samples before ingestion.",
									MarkdownDescription: "MetricRelabelConfigs to apply to samples before ingestion.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"action": schema.StringAttribute{
												Description:         "Action to perform based on regex matching. Default is 'replace'",
												MarkdownDescription: "Action to perform based on regex matching. Default is 'replace'",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"if": schema.MapAttribute{
												Description:         "If represents metricsQL match expression (or list of expressions): '{__name__=~'foo_.*'}'",
												MarkdownDescription: "If represents metricsQL match expression (or list of expressions): '{__name__=~'foo_.*'}'",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"labels": schema.MapAttribute{
												Description:         "Labels is used together with Match for 'action: graphite'",
												MarkdownDescription: "Labels is used together with Match for 'action: graphite'",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"match": schema.StringAttribute{
												Description:         "Match is used together with Labels for 'action: graphite'",
												MarkdownDescription: "Match is used together with Labels for 'action: graphite'",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"modulus": schema.Int64Attribute{
												Description:         "Modulus to take of the hash of the source label values.",
												MarkdownDescription: "Modulus to take of the hash of the source label values.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"regex": schema.StringAttribute{
												Description:         "Regular expression against which the extracted value is matched. Default is '(.*)'",
												MarkdownDescription: "Regular expression against which the extracted value is matched. Default is '(.*)'",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"replacement": schema.StringAttribute{
												Description:         "Replacement value against which a regex replace is performed if theregular expression matches. Regex capture groups are available. Default is '$1'",
												MarkdownDescription: "Replacement value against which a regex replace is performed if theregular expression matches. Regex capture groups are available. Default is '$1'",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"separator": schema.StringAttribute{
												Description:         "Separator placed between concatenated source label values. default is ';'.",
												MarkdownDescription: "Separator placed between concatenated source label values. default is ';'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"source_labels": schema.ListAttribute{
												Description:         "The source labels select values from existing labels. Their content is concatenatedusing the configured separator and matched against the configured regular expressionfor the replace, keep, and drop actions.",
												MarkdownDescription: "The source labels select values from existing labels. Their content is concatenatedusing the configured separator and matched against the configured regular expressionfor the replace, keep, and drop actions.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"target_label": schema.StringAttribute{
												Description:         "Label to which the resulting value is written in a replace action.It is mandatory for replace actions. Regex capture groups are available.",
												MarkdownDescription: "Label to which the resulting value is written in a replace action.It is mandatory for replace actions. Regex capture groups are available.",
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

								"oauth2": schema.SingleNestedAttribute{
									Description:         "OAuth2 defines auth configuration",
									MarkdownDescription: "OAuth2 defines auth configuration",
									Attributes: map[string]schema.Attribute{
										"client_id": schema.SingleNestedAttribute{
											Description:         "The secret or configmap containing the OAuth2 client id",
											MarkdownDescription: "The secret or configmap containing the OAuth2 client id",
											Attributes: map[string]schema.Attribute{
												"config_map": schema.SingleNestedAttribute{
													Description:         "ConfigMap containing data to use for the targets.",
													MarkdownDescription: "ConfigMap containing data to use for the targets.",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key to select.",
															MarkdownDescription: "The key to select.",
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
															Description:         "Specify whether the ConfigMap or its key must be defined",
															MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"secret": schema.SingleNestedAttribute{
													Description:         "Secret containing data to use for the targets.",
													MarkdownDescription: "Secret containing data to use for the targets.",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"client_secret": schema.SingleNestedAttribute{
											Description:         "The secret containing the OAuth2 client secret",
											MarkdownDescription: "The secret containing the OAuth2 client secret",
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

										"client_secret_file": schema.StringAttribute{
											Description:         "ClientSecretFile defines path for client secret file.",
											MarkdownDescription: "ClientSecretFile defines path for client secret file.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"endpoint_params": schema.MapAttribute{
											Description:         "Parameters to append to the token URL",
											MarkdownDescription: "Parameters to append to the token URL",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"scopes": schema.ListAttribute{
											Description:         "OAuth2 scopes used for the token request",
											MarkdownDescription: "OAuth2 scopes used for the token request",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"token_url": schema.StringAttribute{
											Description:         "The URL to fetch the token from",
											MarkdownDescription: "The URL to fetch the token from",
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

								"params": schema.MapAttribute{
									Description:         "Optional HTTP URL parameters",
									MarkdownDescription: "Optional HTTP URL parameters",
									ElementType:         types.ListType{ElemType: types.StringType},
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"path": schema.StringAttribute{
									Description:         "HTTP path to scrape for metrics.",
									MarkdownDescription: "HTTP path to scrape for metrics.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"port": schema.StringAttribute{
									Description:         "Default port for target.",
									MarkdownDescription: "Default port for target.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"proxy_url": schema.StringAttribute{
									Description:         "ProxyURL eg http://proxyserver:2195 Directs scrapes to proxy through this endpoint.",
									MarkdownDescription: "ProxyURL eg http://proxyserver:2195 Directs scrapes to proxy through this endpoint.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"relabel_configs": schema.ListNestedAttribute{
									Description:         "RelabelConfigs to apply to samples before scraping.More info: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config",
									MarkdownDescription: "RelabelConfigs to apply to samples before scraping.More info: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"action": schema.StringAttribute{
												Description:         "Action to perform based on regex matching. Default is 'replace'",
												MarkdownDescription: "Action to perform based on regex matching. Default is 'replace'",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"if": schema.MapAttribute{
												Description:         "If represents metricsQL match expression (or list of expressions): '{__name__=~'foo_.*'}'",
												MarkdownDescription: "If represents metricsQL match expression (or list of expressions): '{__name__=~'foo_.*'}'",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"labels": schema.MapAttribute{
												Description:         "Labels is used together with Match for 'action: graphite'",
												MarkdownDescription: "Labels is used together with Match for 'action: graphite'",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"match": schema.StringAttribute{
												Description:         "Match is used together with Labels for 'action: graphite'",
												MarkdownDescription: "Match is used together with Labels for 'action: graphite'",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"modulus": schema.Int64Attribute{
												Description:         "Modulus to take of the hash of the source label values.",
												MarkdownDescription: "Modulus to take of the hash of the source label values.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"regex": schema.StringAttribute{
												Description:         "Regular expression against which the extracted value is matched. Default is '(.*)'",
												MarkdownDescription: "Regular expression against which the extracted value is matched. Default is '(.*)'",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"replacement": schema.StringAttribute{
												Description:         "Replacement value against which a regex replace is performed if theregular expression matches. Regex capture groups are available. Default is '$1'",
												MarkdownDescription: "Replacement value against which a regex replace is performed if theregular expression matches. Regex capture groups are available. Default is '$1'",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"separator": schema.StringAttribute{
												Description:         "Separator placed between concatenated source label values. default is ';'.",
												MarkdownDescription: "Separator placed between concatenated source label values. default is ';'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"source_labels": schema.ListAttribute{
												Description:         "The source labels select values from existing labels. Their content is concatenatedusing the configured separator and matched against the configured regular expressionfor the replace, keep, and drop actions.",
												MarkdownDescription: "The source labels select values from existing labels. Their content is concatenatedusing the configured separator and matched against the configured regular expressionfor the replace, keep, and drop actions.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"target_label": schema.StringAttribute{
												Description:         "Label to which the resulting value is written in a replace action.It is mandatory for replace actions. Regex capture groups are available.",
												MarkdownDescription: "Label to which the resulting value is written in a replace action.It is mandatory for replace actions. Regex capture groups are available.",
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

								"sample_limit": schema.Int64Attribute{
									Description:         "SampleLimit defines per-scrape limit on number of scraped samples that will be accepted.",
									MarkdownDescription: "SampleLimit defines per-scrape limit on number of scraped samples that will be accepted.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"scheme": schema.StringAttribute{
									Description:         "HTTP scheme to use for scraping.",
									MarkdownDescription: "HTTP scheme to use for scraping.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("http", "https"),
									},
								},

								"scrape_timeout": schema.StringAttribute{
									Description:         "Timeout after which the scrape is ended",
									MarkdownDescription: "Timeout after which the scrape is ended",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"scrape_interval": schema.StringAttribute{
									Description:         "ScrapeInterval is the same as Interval and has priority over it.one of scrape_interval or interval can be used",
									MarkdownDescription: "ScrapeInterval is the same as Interval and has priority over it.one of scrape_interval or interval can be used",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"targets": schema.ListAttribute{
									Description:         "Targets static targets addresses in form of ['192.122.55.55:9100','some-name:9100'].",
									MarkdownDescription: "Targets static targets addresses in form of ['192.122.55.55:9100','some-name:9100'].",
									ElementType:         types.StringType,
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"tls_config": schema.SingleNestedAttribute{
									Description:         "TLSConfig configuration to use when scraping the endpoint",
									MarkdownDescription: "TLSConfig configuration to use when scraping the endpoint",
									Attributes: map[string]schema.Attribute{
										"ca": schema.SingleNestedAttribute{
											Description:         "Stuct containing the CA cert to use for the targets.",
											MarkdownDescription: "Stuct containing the CA cert to use for the targets.",
											Attributes: map[string]schema.Attribute{
												"config_map": schema.SingleNestedAttribute{
													Description:         "ConfigMap containing data to use for the targets.",
													MarkdownDescription: "ConfigMap containing data to use for the targets.",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key to select.",
															MarkdownDescription: "The key to select.",
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
															Description:         "Specify whether the ConfigMap or its key must be defined",
															MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"secret": schema.SingleNestedAttribute{
													Description:         "Secret containing data to use for the targets.",
													MarkdownDescription: "Secret containing data to use for the targets.",
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"ca_file": schema.StringAttribute{
											Description:         "Path to the CA cert in the container to use for the targets.",
											MarkdownDescription: "Path to the CA cert in the container to use for the targets.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"cert": schema.SingleNestedAttribute{
											Description:         "Struct containing the client cert file for the targets.",
											MarkdownDescription: "Struct containing the client cert file for the targets.",
											Attributes: map[string]schema.Attribute{
												"config_map": schema.SingleNestedAttribute{
													Description:         "ConfigMap containing data to use for the targets.",
													MarkdownDescription: "ConfigMap containing data to use for the targets.",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key to select.",
															MarkdownDescription: "The key to select.",
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
															Description:         "Specify whether the ConfigMap or its key must be defined",
															MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"secret": schema.SingleNestedAttribute{
													Description:         "Secret containing data to use for the targets.",
													MarkdownDescription: "Secret containing data to use for the targets.",
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"cert_file": schema.StringAttribute{
											Description:         "Path to the client cert file in the container for the targets.",
											MarkdownDescription: "Path to the client cert file in the container for the targets.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"insecure_skip_verify": schema.BoolAttribute{
											Description:         "Disable target certificate validation.",
											MarkdownDescription: "Disable target certificate validation.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"key_file": schema.StringAttribute{
											Description:         "Path to the client key file in the container for the targets.",
											MarkdownDescription: "Path to the client key file in the container for the targets.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"key_secret": schema.SingleNestedAttribute{
											Description:         "Secret containing the client key file for the targets.",
											MarkdownDescription: "Secret containing the client key file for the targets.",
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

										"server_name": schema.StringAttribute{
											Description:         "Used to verify the hostname for the targets.",
											MarkdownDescription: "Used to verify the hostname for the targets.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"vm_scrape_params": schema.SingleNestedAttribute{
									Description:         "VMScrapeParams defines VictoriaMetrics specific scrape parametrs",
									MarkdownDescription: "VMScrapeParams defines VictoriaMetrics specific scrape parametrs",
									Attributes: map[string]schema.Attribute{
										"disable_compression": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"disable_keep_alive": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"headers": schema.ListAttribute{
											Description:         "Headers allows sending custom headers to scrape targetsmust be in of semicolon separated header with it's valueeg:headerName: headerValuevmagent supports since 1.79.0 version",
											MarkdownDescription: "Headers allows sending custom headers to scrape targetsmust be in of semicolon separated header with it's valueeg:headerName: headerValuevmagent supports since 1.79.0 version",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"metric_relabel_debug": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"no_stale_markers": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"proxy_client_config": schema.SingleNestedAttribute{
											Description:         "ProxyClientConfig configures proxy auth settings for scrapingSee feature description https://docs.victoriametrics.com/vmagent.html#scraping-targets-via-a-proxy",
											MarkdownDescription: "ProxyClientConfig configures proxy auth settings for scrapingSee feature description https://docs.victoriametrics.com/vmagent.html#scraping-targets-via-a-proxy",
											Attributes: map[string]schema.Attribute{
												"basic_auth": schema.SingleNestedAttribute{
													Description:         "BasicAuth allow an endpoint to authenticate over basic authentication",
													MarkdownDescription: "BasicAuth allow an endpoint to authenticate over basic authentication",
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
															Required: false,
															Optional: true,
															Computed: false,
														},

														"password_file": schema.StringAttribute{
															Description:         "PasswordFile defines path to password file at disk",
															MarkdownDescription: "PasswordFile defines path to password file at disk",
															Required:            false,
															Optional:            true,
															Computed:            false,
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
															Required: false,
															Optional: true,
															Computed: false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"bearer_token": schema.SingleNestedAttribute{
													Description:         "SecretKeySelector selects a key of a Secret.",
													MarkdownDescription: "SecretKeySelector selects a key of a Secret.",
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

												"bearer_token_file": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tls_config": schema.SingleNestedAttribute{
													Description:         "TLSConfig specifies TLSConfig configuration parameters.",
													MarkdownDescription: "TLSConfig specifies TLSConfig configuration parameters.",
													Attributes: map[string]schema.Attribute{
														"ca": schema.SingleNestedAttribute{
															Description:         "Stuct containing the CA cert to use for the targets.",
															MarkdownDescription: "Stuct containing the CA cert to use for the targets.",
															Attributes: map[string]schema.Attribute{
																"config_map": schema.SingleNestedAttribute{
																	Description:         "ConfigMap containing data to use for the targets.",
																	MarkdownDescription: "ConfigMap containing data to use for the targets.",
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "The key to select.",
																			MarkdownDescription: "The key to select.",
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
																			Description:         "Specify whether the ConfigMap or its key must be defined",
																			MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"secret": schema.SingleNestedAttribute{
																	Description:         "Secret containing data to use for the targets.",
																	MarkdownDescription: "Secret containing data to use for the targets.",
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
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"ca_file": schema.StringAttribute{
															Description:         "Path to the CA cert in the container to use for the targets.",
															MarkdownDescription: "Path to the CA cert in the container to use for the targets.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"cert": schema.SingleNestedAttribute{
															Description:         "Struct containing the client cert file for the targets.",
															MarkdownDescription: "Struct containing the client cert file for the targets.",
															Attributes: map[string]schema.Attribute{
																"config_map": schema.SingleNestedAttribute{
																	Description:         "ConfigMap containing data to use for the targets.",
																	MarkdownDescription: "ConfigMap containing data to use for the targets.",
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "The key to select.",
																			MarkdownDescription: "The key to select.",
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
																			Description:         "Specify whether the ConfigMap or its key must be defined",
																			MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"secret": schema.SingleNestedAttribute{
																	Description:         "Secret containing data to use for the targets.",
																	MarkdownDescription: "Secret containing data to use for the targets.",
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
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"cert_file": schema.StringAttribute{
															Description:         "Path to the client cert file in the container for the targets.",
															MarkdownDescription: "Path to the client cert file in the container for the targets.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"insecure_skip_verify": schema.BoolAttribute{
															Description:         "Disable target certificate validation.",
															MarkdownDescription: "Disable target certificate validation.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"key_file": schema.StringAttribute{
															Description:         "Path to the client key file in the container for the targets.",
															MarkdownDescription: "Path to the client key file in the container for the targets.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"key_secret": schema.SingleNestedAttribute{
															Description:         "Secret containing the client key file for the targets.",
															MarkdownDescription: "Secret containing the client key file for the targets.",
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

														"server_name": schema.StringAttribute{
															Description:         "Used to verify the hostname for the targets.",
															MarkdownDescription: "Used to verify the hostname for the targets.",
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

										"relabel_debug": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"scrape_align_interval": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"scrape_offset": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"stream_parse": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
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
	}
}

func (r *OperatorVictoriametricsComVmstaticScrapeV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_operator_victoriametrics_com_vm_static_scrape_v1beta1_manifest")

	var model OperatorVictoriametricsComVmstaticScrapeV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("operator.victoriametrics.com/v1beta1")
	model.Kind = pointer.String("VMStaticScrape")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
