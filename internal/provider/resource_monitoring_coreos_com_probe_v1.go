/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type MonitoringCoreosComProbeV1Resource struct{}

var (
	_ resource.Resource = (*MonitoringCoreosComProbeV1Resource)(nil)
)

type MonitoringCoreosComProbeV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type MonitoringCoreosComProbeV1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Authorization *struct {
			Credentials *struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
			} `tfsdk:"credentials" yaml:"credentials,omitempty"`

			Type *string `tfsdk:"type" yaml:"type,omitempty"`
		} `tfsdk:"authorization" yaml:"authorization,omitempty"`

		BasicAuth *struct {
			Password *struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
			} `tfsdk:"password" yaml:"password,omitempty"`

			Username *struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
			} `tfsdk:"username" yaml:"username,omitempty"`
		} `tfsdk:"basic_auth" yaml:"basicAuth,omitempty"`

		BearerTokenSecret *struct {
			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
		} `tfsdk:"bearer_token_secret" yaml:"bearerTokenSecret,omitempty"`

		Interval *string `tfsdk:"interval" yaml:"interval,omitempty"`

		JobName *string `tfsdk:"job_name" yaml:"jobName,omitempty"`

		LabelLimit *int64 `tfsdk:"label_limit" yaml:"labelLimit,omitempty"`

		LabelNameLengthLimit *int64 `tfsdk:"label_name_length_limit" yaml:"labelNameLengthLimit,omitempty"`

		LabelValueLengthLimit *int64 `tfsdk:"label_value_length_limit" yaml:"labelValueLengthLimit,omitempty"`

		MetricRelabelings *[]struct {
			Action *string `tfsdk:"action" yaml:"action,omitempty"`

			Modulus *int64 `tfsdk:"modulus" yaml:"modulus,omitempty"`

			Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

			Replacement *string `tfsdk:"replacement" yaml:"replacement,omitempty"`

			Separator *string `tfsdk:"separator" yaml:"separator,omitempty"`

			SourceLabels *[]string `tfsdk:"source_labels" yaml:"sourceLabels,omitempty"`

			TargetLabel *string `tfsdk:"target_label" yaml:"targetLabel,omitempty"`
		} `tfsdk:"metric_relabelings" yaml:"metricRelabelings,omitempty"`

		Module *string `tfsdk:"module" yaml:"module,omitempty"`

		Oauth2 *struct {
			ClientId *struct {
				ConfigMap *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"config_map" yaml:"configMap,omitempty"`

				Secret *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"secret" yaml:"secret,omitempty"`
			} `tfsdk:"client_id" yaml:"clientId,omitempty"`

			ClientSecret *struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
			} `tfsdk:"client_secret" yaml:"clientSecret,omitempty"`

			EndpointParams *map[string]string `tfsdk:"endpoint_params" yaml:"endpointParams,omitempty"`

			Scopes *[]string `tfsdk:"scopes" yaml:"scopes,omitempty"`

			TokenUrl *string `tfsdk:"token_url" yaml:"tokenUrl,omitempty"`
		} `tfsdk:"oauth2" yaml:"oauth2,omitempty"`

		Prober *struct {
			Path *string `tfsdk:"path" yaml:"path,omitempty"`

			ProxyUrl *string `tfsdk:"proxy_url" yaml:"proxyUrl,omitempty"`

			Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`

			Url *string `tfsdk:"url" yaml:"url,omitempty"`
		} `tfsdk:"prober" yaml:"prober,omitempty"`

		SampleLimit *int64 `tfsdk:"sample_limit" yaml:"sampleLimit,omitempty"`

		ScrapeTimeout *string `tfsdk:"scrape_timeout" yaml:"scrapeTimeout,omitempty"`

		TargetLimit *int64 `tfsdk:"target_limit" yaml:"targetLimit,omitempty"`

		Targets *struct {
			Ingress *struct {
				NamespaceSelector *struct {
					Any *bool `tfsdk:"any" yaml:"any,omitempty"`

					MatchNames *[]string `tfsdk:"match_names" yaml:"matchNames,omitempty"`
				} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

				RelabelingConfigs *[]struct {
					Action *string `tfsdk:"action" yaml:"action,omitempty"`

					Modulus *int64 `tfsdk:"modulus" yaml:"modulus,omitempty"`

					Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

					Replacement *string `tfsdk:"replacement" yaml:"replacement,omitempty"`

					Separator *string `tfsdk:"separator" yaml:"separator,omitempty"`

					SourceLabels *[]string `tfsdk:"source_labels" yaml:"sourceLabels,omitempty"`

					TargetLabel *string `tfsdk:"target_label" yaml:"targetLabel,omitempty"`
				} `tfsdk:"relabeling_configs" yaml:"relabelingConfigs,omitempty"`

				Selector *struct {
					MatchExpressions *[]struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

						Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
					} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

					MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
				} `tfsdk:"selector" yaml:"selector,omitempty"`
			} `tfsdk:"ingress" yaml:"ingress,omitempty"`

			StaticConfig *struct {
				Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

				RelabelingConfigs *[]struct {
					Action *string `tfsdk:"action" yaml:"action,omitempty"`

					Modulus *int64 `tfsdk:"modulus" yaml:"modulus,omitempty"`

					Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

					Replacement *string `tfsdk:"replacement" yaml:"replacement,omitempty"`

					Separator *string `tfsdk:"separator" yaml:"separator,omitempty"`

					SourceLabels *[]string `tfsdk:"source_labels" yaml:"sourceLabels,omitempty"`

					TargetLabel *string `tfsdk:"target_label" yaml:"targetLabel,omitempty"`
				} `tfsdk:"relabeling_configs" yaml:"relabelingConfigs,omitempty"`

				Static *[]string `tfsdk:"static" yaml:"static,omitempty"`
			} `tfsdk:"static_config" yaml:"staticConfig,omitempty"`
		} `tfsdk:"targets" yaml:"targets,omitempty"`

		TlsConfig *struct {
			Ca *struct {
				ConfigMap *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"config_map" yaml:"configMap,omitempty"`

				Secret *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"secret" yaml:"secret,omitempty"`
			} `tfsdk:"ca" yaml:"ca,omitempty"`

			Cert *struct {
				ConfigMap *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"config_map" yaml:"configMap,omitempty"`

				Secret *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"secret" yaml:"secret,omitempty"`
			} `tfsdk:"cert" yaml:"cert,omitempty"`

			InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" yaml:"insecureSkipVerify,omitempty"`

			KeySecret *struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
			} `tfsdk:"key_secret" yaml:"keySecret,omitempty"`

			ServerName *string `tfsdk:"server_name" yaml:"serverName,omitempty"`
		} `tfsdk:"tls_config" yaml:"tlsConfig,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewMonitoringCoreosComProbeV1Resource() resource.Resource {
	return &MonitoringCoreosComProbeV1Resource{}
}

func (r *MonitoringCoreosComProbeV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_monitoring_coreos_com_probe_v1"
}

func (r *MonitoringCoreosComProbeV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Probe defines monitoring for a set of static targets or ingresses.",
		MarkdownDescription: "Probe defines monitoring for a set of static targets or ingresses.",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "Specification of desired Ingress selection for target discovery by Prometheus.",
				MarkdownDescription: "Specification of desired Ingress selection for target discovery by Prometheus.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"authorization": {
						Description:         "Authorization section for this endpoint",
						MarkdownDescription: "Authorization section for this endpoint",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"credentials": {
								Description:         "The secret's key that contains the credentials of the request",
								MarkdownDescription: "The secret's key that contains the credentials of the request",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"key": {
										Description:         "The key of the secret to select from.  Must be a valid secret key.",
										MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"name": {
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"optional": {
										Description:         "Specify whether the Secret or its key must be defined",
										MarkdownDescription: "Specify whether the Secret or its key must be defined",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": {
								Description:         "Set the authentication type. Defaults to Bearer, Basic will cause an error",
								MarkdownDescription: "Set the authentication type. Defaults to Bearer, Basic will cause an error",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"basic_auth": {
						Description:         "BasicAuth allow an endpoint to authenticate over basic authentication. More info: https://prometheus.io/docs/operating/configuration/#endpoint",
						MarkdownDescription: "BasicAuth allow an endpoint to authenticate over basic authentication. More info: https://prometheus.io/docs/operating/configuration/#endpoint",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"password": {
								Description:         "The secret in the service monitor namespace that contains the password for authentication.",
								MarkdownDescription: "The secret in the service monitor namespace that contains the password for authentication.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"key": {
										Description:         "The key of the secret to select from.  Must be a valid secret key.",
										MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"name": {
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"optional": {
										Description:         "Specify whether the Secret or its key must be defined",
										MarkdownDescription: "Specify whether the Secret or its key must be defined",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"username": {
								Description:         "The secret in the service monitor namespace that contains the username for authentication.",
								MarkdownDescription: "The secret in the service monitor namespace that contains the username for authentication.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"key": {
										Description:         "The key of the secret to select from.  Must be a valid secret key.",
										MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"name": {
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"optional": {
										Description:         "Specify whether the Secret or its key must be defined",
										MarkdownDescription: "Specify whether the Secret or its key must be defined",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"bearer_token_secret": {
						Description:         "Secret to mount to read bearer token for scraping targets. The secret needs to be in the same namespace as the probe and accessible by the Prometheus Operator.",
						MarkdownDescription: "Secret to mount to read bearer token for scraping targets. The secret needs to be in the same namespace as the probe and accessible by the Prometheus Operator.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"key": {
								Description:         "The key of the secret to select from.  Must be a valid secret key.",
								MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"name": {
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"optional": {
								Description:         "Specify whether the Secret or its key must be defined",
								MarkdownDescription: "Specify whether the Secret or its key must be defined",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"interval": {
						Description:         "Interval at which targets are probed using the configured prober. If not specified Prometheus' global scrape interval is used.",
						MarkdownDescription: "Interval at which targets are probed using the configured prober. If not specified Prometheus' global scrape interval is used.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"job_name": {
						Description:         "The job name assigned to scraped metrics by default.",
						MarkdownDescription: "The job name assigned to scraped metrics by default.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"label_limit": {
						Description:         "Per-scrape limit on number of labels that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.",
						MarkdownDescription: "Per-scrape limit on number of labels that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"label_name_length_limit": {
						Description:         "Per-scrape limit on length of labels name that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.",
						MarkdownDescription: "Per-scrape limit on length of labels name that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"label_value_length_limit": {
						Description:         "Per-scrape limit on length of labels value that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.",
						MarkdownDescription: "Per-scrape limit on length of labels value that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"metric_relabelings": {
						Description:         "MetricRelabelConfigs to apply to samples before ingestion.",
						MarkdownDescription: "MetricRelabelConfigs to apply to samples before ingestion.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"action": {
								Description:         "Action to perform based on regex matching. Default is 'replace'. uppercase and lowercase actions require Prometheus >= 2.36.",
								MarkdownDescription: "Action to perform based on regex matching. Default is 'replace'. uppercase and lowercase actions require Prometheus >= 2.36.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"modulus": {
								Description:         "Modulus to take of the hash of the source label values.",
								MarkdownDescription: "Modulus to take of the hash of the source label values.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"regex": {
								Description:         "Regular expression against which the extracted value is matched. Default is '(.*)'",
								MarkdownDescription: "Regular expression against which the extracted value is matched. Default is '(.*)'",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"replacement": {
								Description:         "Replacement value against which a regex replace is performed if the regular expression matches. Regex capture groups are available. Default is '$1'",
								MarkdownDescription: "Replacement value against which a regex replace is performed if the regular expression matches. Regex capture groups are available. Default is '$1'",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"separator": {
								Description:         "Separator placed between concatenated source label values. default is ';'.",
								MarkdownDescription: "Separator placed between concatenated source label values. default is ';'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"source_labels": {
								Description:         "The source labels select values from existing labels. Their content is concatenated using the configured separator and matched against the configured regular expression for the replace, keep, and drop actions.",
								MarkdownDescription: "The source labels select values from existing labels. Their content is concatenated using the configured separator and matched against the configured regular expression for the replace, keep, and drop actions.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"target_label": {
								Description:         "Label to which the resulting value is written in a replace action. It is mandatory for replace actions. Regex capture groups are available.",
								MarkdownDescription: "Label to which the resulting value is written in a replace action. It is mandatory for replace actions. Regex capture groups are available.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"module": {
						Description:         "The module to use for probing specifying how to probe the target. Example module configuring in the blackbox exporter: https://github.com/prometheus/blackbox_exporter/blob/master/example.yml",
						MarkdownDescription: "The module to use for probing specifying how to probe the target. Example module configuring in the blackbox exporter: https://github.com/prometheus/blackbox_exporter/blob/master/example.yml",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"oauth2": {
						Description:         "OAuth2 for the URL. Only valid in Prometheus versions 2.27.0 and newer.",
						MarkdownDescription: "OAuth2 for the URL. Only valid in Prometheus versions 2.27.0 and newer.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"client_id": {
								Description:         "The secret or configmap containing the OAuth2 client id",
								MarkdownDescription: "The secret or configmap containing the OAuth2 client id",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"config_map": {
										Description:         "ConfigMap containing data to use for the targets.",
										MarkdownDescription: "ConfigMap containing data to use for the targets.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "The key to select.",
												MarkdownDescription: "The key to select.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"optional": {
												Description:         "Specify whether the ConfigMap or its key must be defined",
												MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret": {
										Description:         "Secret containing data to use for the targets.",
										MarkdownDescription: "Secret containing data to use for the targets.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "The key of the secret to select from.  Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"optional": {
												Description:         "Specify whether the Secret or its key must be defined",
												MarkdownDescription: "Specify whether the Secret or its key must be defined",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: true,
								Optional: false,
								Computed: false,
							},

							"client_secret": {
								Description:         "The secret containing the OAuth2 client secret",
								MarkdownDescription: "The secret containing the OAuth2 client secret",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"key": {
										Description:         "The key of the secret to select from.  Must be a valid secret key.",
										MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"name": {
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"optional": {
										Description:         "Specify whether the Secret or its key must be defined",
										MarkdownDescription: "Specify whether the Secret or its key must be defined",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: true,
								Optional: false,
								Computed: false,
							},

							"endpoint_params": {
								Description:         "Parameters to append to the token URL",
								MarkdownDescription: "Parameters to append to the token URL",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"scopes": {
								Description:         "OAuth2 scopes used for the token request",
								MarkdownDescription: "OAuth2 scopes used for the token request",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"token_url": {
								Description:         "The URL to fetch the token from",
								MarkdownDescription: "The URL to fetch the token from",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.LengthAtLeast(1),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"prober": {
						Description:         "Specification for the prober to use for probing targets. The prober.URL parameter is required. Targets cannot be probed if left empty.",
						MarkdownDescription: "Specification for the prober to use for probing targets. The prober.URL parameter is required. Targets cannot be probed if left empty.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"path": {
								Description:         "Path to collect metrics from. Defaults to '/probe'.",
								MarkdownDescription: "Path to collect metrics from. Defaults to '/probe'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"proxy_url": {
								Description:         "Optional ProxyURL.",
								MarkdownDescription: "Optional ProxyURL.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"scheme": {
								Description:         "HTTP scheme to use for scraping. Defaults to 'http'.",
								MarkdownDescription: "HTTP scheme to use for scraping. Defaults to 'http'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"url": {
								Description:         "Mandatory URL of the prober.",
								MarkdownDescription: "Mandatory URL of the prober.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"sample_limit": {
						Description:         "SampleLimit defines per-scrape limit on number of scraped samples that will be accepted.",
						MarkdownDescription: "SampleLimit defines per-scrape limit on number of scraped samples that will be accepted.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"scrape_timeout": {
						Description:         "Timeout for scraping metrics from the Prometheus exporter. If not specified, the Prometheus global scrape interval is used.",
						MarkdownDescription: "Timeout for scraping metrics from the Prometheus exporter. If not specified, the Prometheus global scrape interval is used.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"target_limit": {
						Description:         "TargetLimit defines a limit on the number of scraped targets that will be accepted.",
						MarkdownDescription: "TargetLimit defines a limit on the number of scraped targets that will be accepted.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"targets": {
						Description:         "Targets defines a set of static or dynamically discovered targets to probe.",
						MarkdownDescription: "Targets defines a set of static or dynamically discovered targets to probe.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"ingress": {
								Description:         "ingress defines the Ingress objects to probe and the relabeling configuration. If 'staticConfig' is also defined, 'staticConfig' takes precedence.",
								MarkdownDescription: "ingress defines the Ingress objects to probe and the relabeling configuration. If 'staticConfig' is also defined, 'staticConfig' takes precedence.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"namespace_selector": {
										Description:         "From which namespaces to select Ingress objects.",
										MarkdownDescription: "From which namespaces to select Ingress objects.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"any": {
												Description:         "Boolean describing whether all namespaces are selected in contrast to a list restricting them.",
												MarkdownDescription: "Boolean describing whether all namespaces are selected in contrast to a list restricting them.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"match_names": {
												Description:         "List of namespace names to select from.",
												MarkdownDescription: "List of namespace names to select from.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"relabeling_configs": {
										Description:         "RelabelConfigs to apply to the label set of the target before it gets scraped. The original ingress address is available via the '__tmp_prometheus_ingress_address' label. It can be used to customize the probed URL. The original scrape job's name is available via the '__tmp_prometheus_job_name' label. More info: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config",
										MarkdownDescription: "RelabelConfigs to apply to the label set of the target before it gets scraped. The original ingress address is available via the '__tmp_prometheus_ingress_address' label. It can be used to customize the probed URL. The original scrape job's name is available via the '__tmp_prometheus_job_name' label. More info: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"action": {
												Description:         "Action to perform based on regex matching. Default is 'replace'. uppercase and lowercase actions require Prometheus >= 2.36.",
												MarkdownDescription: "Action to perform based on regex matching. Default is 'replace'. uppercase and lowercase actions require Prometheus >= 2.36.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"modulus": {
												Description:         "Modulus to take of the hash of the source label values.",
												MarkdownDescription: "Modulus to take of the hash of the source label values.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"regex": {
												Description:         "Regular expression against which the extracted value is matched. Default is '(.*)'",
												MarkdownDescription: "Regular expression against which the extracted value is matched. Default is '(.*)'",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"replacement": {
												Description:         "Replacement value against which a regex replace is performed if the regular expression matches. Regex capture groups are available. Default is '$1'",
												MarkdownDescription: "Replacement value against which a regex replace is performed if the regular expression matches. Regex capture groups are available. Default is '$1'",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"separator": {
												Description:         "Separator placed between concatenated source label values. default is ';'.",
												MarkdownDescription: "Separator placed between concatenated source label values. default is ';'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"source_labels": {
												Description:         "The source labels select values from existing labels. Their content is concatenated using the configured separator and matched against the configured regular expression for the replace, keep, and drop actions.",
												MarkdownDescription: "The source labels select values from existing labels. Their content is concatenated using the configured separator and matched against the configured regular expression for the replace, keep, and drop actions.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"target_label": {
												Description:         "Label to which the resulting value is written in a replace action. It is mandatory for replace actions. Regex capture groups are available.",
												MarkdownDescription: "Label to which the resulting value is written in a replace action. It is mandatory for replace actions. Regex capture groups are available.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"selector": {
										Description:         "Selector to select the Ingress objects.",
										MarkdownDescription: "Selector to select the Ingress objects.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"match_expressions": {
												Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "key is the label key that the selector applies to.",
														MarkdownDescription: "key is the label key that the selector applies to.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"operator": {
														Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"values": {
														Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"match_labels": {
												Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"static_config": {
								Description:         "staticConfig defines the static list of targets to probe and the relabeling configuration. If 'ingress' is also defined, 'staticConfig' takes precedence. More info: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#static_config.",
								MarkdownDescription: "staticConfig defines the static list of targets to probe and the relabeling configuration. If 'ingress' is also defined, 'staticConfig' takes precedence. More info: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#static_config.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"labels": {
										Description:         "Labels assigned to all metrics scraped from the targets.",
										MarkdownDescription: "Labels assigned to all metrics scraped from the targets.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"relabeling_configs": {
										Description:         "RelabelConfigs to apply to the label set of the targets before it gets scraped. More info: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config",
										MarkdownDescription: "RelabelConfigs to apply to the label set of the targets before it gets scraped. More info: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"action": {
												Description:         "Action to perform based on regex matching. Default is 'replace'. uppercase and lowercase actions require Prometheus >= 2.36.",
												MarkdownDescription: "Action to perform based on regex matching. Default is 'replace'. uppercase and lowercase actions require Prometheus >= 2.36.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"modulus": {
												Description:         "Modulus to take of the hash of the source label values.",
												MarkdownDescription: "Modulus to take of the hash of the source label values.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"regex": {
												Description:         "Regular expression against which the extracted value is matched. Default is '(.*)'",
												MarkdownDescription: "Regular expression against which the extracted value is matched. Default is '(.*)'",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"replacement": {
												Description:         "Replacement value against which a regex replace is performed if the regular expression matches. Regex capture groups are available. Default is '$1'",
												MarkdownDescription: "Replacement value against which a regex replace is performed if the regular expression matches. Regex capture groups are available. Default is '$1'",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"separator": {
												Description:         "Separator placed between concatenated source label values. default is ';'.",
												MarkdownDescription: "Separator placed between concatenated source label values. default is ';'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"source_labels": {
												Description:         "The source labels select values from existing labels. Their content is concatenated using the configured separator and matched against the configured regular expression for the replace, keep, and drop actions.",
												MarkdownDescription: "The source labels select values from existing labels. Their content is concatenated using the configured separator and matched against the configured regular expression for the replace, keep, and drop actions.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"target_label": {
												Description:         "Label to which the resulting value is written in a replace action. It is mandatory for replace actions. Regex capture groups are available.",
												MarkdownDescription: "Label to which the resulting value is written in a replace action. It is mandatory for replace actions. Regex capture groups are available.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"static": {
										Description:         "The list of hosts to probe.",
										MarkdownDescription: "The list of hosts to probe.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"tls_config": {
						Description:         "TLS configuration to use when scraping the endpoint.",
						MarkdownDescription: "TLS configuration to use when scraping the endpoint.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"ca": {
								Description:         "Struct containing the CA cert to use for the targets.",
								MarkdownDescription: "Struct containing the CA cert to use for the targets.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"config_map": {
										Description:         "ConfigMap containing data to use for the targets.",
										MarkdownDescription: "ConfigMap containing data to use for the targets.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "The key to select.",
												MarkdownDescription: "The key to select.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"optional": {
												Description:         "Specify whether the ConfigMap or its key must be defined",
												MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret": {
										Description:         "Secret containing data to use for the targets.",
										MarkdownDescription: "Secret containing data to use for the targets.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "The key of the secret to select from.  Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"optional": {
												Description:         "Specify whether the Secret or its key must be defined",
												MarkdownDescription: "Specify whether the Secret or its key must be defined",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cert": {
								Description:         "Struct containing the client cert file for the targets.",
								MarkdownDescription: "Struct containing the client cert file for the targets.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"config_map": {
										Description:         "ConfigMap containing data to use for the targets.",
										MarkdownDescription: "ConfigMap containing data to use for the targets.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "The key to select.",
												MarkdownDescription: "The key to select.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"optional": {
												Description:         "Specify whether the ConfigMap or its key must be defined",
												MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret": {
										Description:         "Secret containing data to use for the targets.",
										MarkdownDescription: "Secret containing data to use for the targets.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "The key of the secret to select from.  Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"optional": {
												Description:         "Specify whether the Secret or its key must be defined",
												MarkdownDescription: "Specify whether the Secret or its key must be defined",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"insecure_skip_verify": {
								Description:         "Disable target certificate validation.",
								MarkdownDescription: "Disable target certificate validation.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"key_secret": {
								Description:         "Secret containing the client key file for the targets.",
								MarkdownDescription: "Secret containing the client key file for the targets.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"key": {
										Description:         "The key of the secret to select from.  Must be a valid secret key.",
										MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"name": {
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"optional": {
										Description:         "Specify whether the Secret or its key must be defined",
										MarkdownDescription: "Specify whether the Secret or its key must be defined",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"server_name": {
								Description:         "Used to verify the hostname for the targets.",
								MarkdownDescription: "Used to verify the hostname for the targets.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}, nil
}

func (r *MonitoringCoreosComProbeV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_monitoring_coreos_com_probe_v1")

	var state MonitoringCoreosComProbeV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel MonitoringCoreosComProbeV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("monitoring.coreos.com/v1")
	goModel.Kind = utilities.Ptr("Probe")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *MonitoringCoreosComProbeV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_monitoring_coreos_com_probe_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *MonitoringCoreosComProbeV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_monitoring_coreos_com_probe_v1")

	var state MonitoringCoreosComProbeV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel MonitoringCoreosComProbeV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("monitoring.coreos.com/v1")
	goModel.Kind = utilities.Ptr("Probe")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *MonitoringCoreosComProbeV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_monitoring_coreos_com_probe_v1")
	// NO-OP: Terraform removes the state automatically for us
}
