/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package getambassador_io_v1

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/schemavalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &GetambassadorIoMappingV1Manifest{}
)

func NewGetambassadorIoMappingV1Manifest() datasource.DataSource {
	return &GetambassadorIoMappingV1Manifest{}
}

type GetambassadorIoMappingV1Manifest struct{}

type GetambassadorIoMappingV1ManifestData struct {
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
		Add_linkerd_headers             *bool              `tfsdk:"add_linkerd_headers" json:"add_linkerd_headers,omitempty"`
		Add_request_headers             *map[string]string `tfsdk:"add_request_headers" json:"add_request_headers,omitempty"`
		Add_response_headers            *map[string]string `tfsdk:"add_response_headers" json:"add_response_headers,omitempty"`
		Allow_upgrade                   *[]string          `tfsdk:"allow_upgrade" json:"allow_upgrade,omitempty"`
		Ambassador_id                   *[]string          `tfsdk:"ambassador_id" json:"ambassador_id,omitempty"`
		Auth_context_extensions         *map[string]string `tfsdk:"auth_context_extensions" json:"auth_context_extensions,omitempty"`
		Auto_host_rewrite               *bool              `tfsdk:"auto_host_rewrite" json:"auto_host_rewrite,omitempty"`
		Bypass_auth                     *bool              `tfsdk:"bypass_auth" json:"bypass_auth,omitempty"`
		Bypass_error_response_overrides *bool              `tfsdk:"bypass_error_response_overrides" json:"bypass_error_response_overrides,omitempty"`
		Case_sensitive                  *bool              `tfsdk:"case_sensitive" json:"case_sensitive,omitempty"`
		Circuit_breakers                *[]struct {
			Max_connections      *int64  `tfsdk:"max_connections" json:"max_connections,omitempty"`
			Max_pending_requests *int64  `tfsdk:"max_pending_requests" json:"max_pending_requests,omitempty"`
			Max_requests         *int64  `tfsdk:"max_requests" json:"max_requests,omitempty"`
			Max_retries          *int64  `tfsdk:"max_retries" json:"max_retries,omitempty"`
			Priority             *string `tfsdk:"priority" json:"priority,omitempty"`
		} `tfsdk:"circuit_breakers" json:"circuit_breakers,omitempty"`
		Cluster_idle_timeout_ms            *int64  `tfsdk:"cluster_idle_timeout_ms" json:"cluster_idle_timeout_ms,omitempty"`
		Cluster_max_connection_lifetime_ms *int64  `tfsdk:"cluster_max_connection_lifetime_ms" json:"cluster_max_connection_lifetime_ms,omitempty"`
		Cluster_tag                        *string `tfsdk:"cluster_tag" json:"cluster_tag,omitempty"`
		Connect_timeout_ms                 *int64  `tfsdk:"connect_timeout_ms" json:"connect_timeout_ms,omitempty"`
		Cors                               *struct {
			Credentials     *bool     `tfsdk:"credentials" json:"credentials,omitempty"`
			Exposed_headers *[]string `tfsdk:"exposed_headers" json:"exposed_headers,omitempty"`
			Headers         *[]string `tfsdk:"headers" json:"headers,omitempty"`
			Max_age         *string   `tfsdk:"max_age" json:"max_age,omitempty"`
			Methods         *[]string `tfsdk:"methods" json:"methods,omitempty"`
			Origins         *[]string `tfsdk:"origins" json:"origins,omitempty"`
		} `tfsdk:"cors" json:"cors,omitempty"`
		Dns_type *string `tfsdk:"dns_type" json:"dns_type,omitempty"`
		Docs     *struct {
			Display_name *string `tfsdk:"display_name" json:"display_name,omitempty"`
			Ignored      *bool   `tfsdk:"ignored" json:"ignored,omitempty"`
			Path         *string `tfsdk:"path" json:"path,omitempty"`
			Timeout_ms   *int64  `tfsdk:"timeout_ms" json:"timeout_ms,omitempty"`
			Url          *string `tfsdk:"url" json:"url,omitempty"`
		} `tfsdk:"docs" json:"docs,omitempty"`
		Enable_ipv4              *bool              `tfsdk:"enable_ipv4" json:"enable_ipv4,omitempty"`
		Enable_ipv6              *bool              `tfsdk:"enable_ipv6" json:"enable_ipv6,omitempty"`
		Envoy_override           *map[string]string `tfsdk:"envoy_override" json:"envoy_override,omitempty"`
		Error_response_overrides *[]struct {
			Body *struct {
				Content_type       *string            `tfsdk:"content_type" json:"content_type,omitempty"`
				Json_format        *map[string]string `tfsdk:"json_format" json:"json_format,omitempty"`
				Text_format        *string            `tfsdk:"text_format" json:"text_format,omitempty"`
				Text_format_source *struct {
					Filename *string `tfsdk:"filename" json:"filename,omitempty"`
				} `tfsdk:"text_format_source" json:"text_format_source,omitempty"`
			} `tfsdk:"body" json:"body,omitempty"`
			On_status_code *int64 `tfsdk:"on_status_code" json:"on_status_code,omitempty"`
		} `tfsdk:"error_response_overrides" json:"error_response_overrides,omitempty"`
		Grpc            *bool              `tfsdk:"grpc" json:"grpc,omitempty"`
		Headers         *map[string]string `tfsdk:"headers" json:"headers,omitempty"`
		Host            *string            `tfsdk:"host" json:"host,omitempty"`
		Host_redirect   *bool              `tfsdk:"host_redirect" json:"host_redirect,omitempty"`
		Host_regex      *bool              `tfsdk:"host_regex" json:"host_regex,omitempty"`
		Host_rewrite    *string            `tfsdk:"host_rewrite" json:"host_rewrite,omitempty"`
		Idle_timeout_ms *int64             `tfsdk:"idle_timeout_ms" json:"idle_timeout_ms,omitempty"`
		Keepalive       *struct {
			Idle_time *int64 `tfsdk:"idle_time" json:"idle_time,omitempty"`
			Interval  *int64 `tfsdk:"interval" json:"interval,omitempty"`
			Probes    *int64 `tfsdk:"probes" json:"probes,omitempty"`
		} `tfsdk:"keepalive" json:"keepalive,omitempty"`
		Labels        *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Load_balancer *struct {
			Cookie *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
				Path *string `tfsdk:"path" json:"path,omitempty"`
				Ttl  *string `tfsdk:"ttl" json:"ttl,omitempty"`
			} `tfsdk:"cookie" json:"cookie,omitempty"`
			Header    *string `tfsdk:"header" json:"header,omitempty"`
			Policy    *string `tfsdk:"policy" json:"policy,omitempty"`
			Source_ip *bool   `tfsdk:"source_ip" json:"source_ip,omitempty"`
		} `tfsdk:"load_balancer" json:"load_balancer,omitempty"`
		Method                 *string              `tfsdk:"method" json:"method,omitempty"`
		Method_regex           *bool                `tfsdk:"method_regex" json:"method_regex,omitempty"`
		Modules                *[]map[string]string `tfsdk:"modules" json:"modules,omitempty"`
		Outlier_detection      *string              `tfsdk:"outlier_detection" json:"outlier_detection,omitempty"`
		Path_redirect          *string              `tfsdk:"path_redirect" json:"path_redirect,omitempty"`
		Precedence             *int64               `tfsdk:"precedence" json:"precedence,omitempty"`
		Prefix                 *string              `tfsdk:"prefix" json:"prefix,omitempty"`
		Prefix_exact           *bool                `tfsdk:"prefix_exact" json:"prefix_exact,omitempty"`
		Prefix_redirect        *string              `tfsdk:"prefix_redirect" json:"prefix_redirect,omitempty"`
		Prefix_regex           *bool                `tfsdk:"prefix_regex" json:"prefix_regex,omitempty"`
		Priority               *string              `tfsdk:"priority" json:"priority,omitempty"`
		Query_parameters       *map[string]string   `tfsdk:"query_parameters" json:"query_parameters,omitempty"`
		Redirect_response_code *int64               `tfsdk:"redirect_response_code" json:"redirect_response_code,omitempty"`
		Regex_headers          *map[string]string   `tfsdk:"regex_headers" json:"regex_headers,omitempty"`
		Regex_query_parameters *map[string]string   `tfsdk:"regex_query_parameters" json:"regex_query_parameters,omitempty"`
		Regex_redirect         *struct {
			Pattern      *string `tfsdk:"pattern" json:"pattern,omitempty"`
			Substitution *string `tfsdk:"substitution" json:"substitution,omitempty"`
		} `tfsdk:"regex_redirect" json:"regex_redirect,omitempty"`
		Regex_rewrite *struct {
			Pattern      *string `tfsdk:"pattern" json:"pattern,omitempty"`
			Substitution *string `tfsdk:"substitution" json:"substitution,omitempty"`
		} `tfsdk:"regex_rewrite" json:"regex_rewrite,omitempty"`
		Remove_request_headers  *[]string `tfsdk:"remove_request_headers" json:"remove_request_headers,omitempty"`
		Remove_response_headers *[]string `tfsdk:"remove_response_headers" json:"remove_response_headers,omitempty"`
		Resolver                *string   `tfsdk:"resolver" json:"resolver,omitempty"`
		Respect_dns_ttl         *bool     `tfsdk:"respect_dns_ttl" json:"respect_dns_ttl,omitempty"`
		Retry_policy            *struct {
			Num_retries     *int64  `tfsdk:"num_retries" json:"num_retries,omitempty"`
			Per_try_timeout *string `tfsdk:"per_try_timeout" json:"per_try_timeout,omitempty"`
			Retry_on        *string `tfsdk:"retry_on" json:"retry_on,omitempty"`
		} `tfsdk:"retry_policy" json:"retry_policy,omitempty"`
		Rewrite         *string `tfsdk:"rewrite" json:"rewrite,omitempty"`
		Service         *string `tfsdk:"service" json:"service,omitempty"`
		Shadow          *bool   `tfsdk:"shadow" json:"shadow,omitempty"`
		Timeout_ms      *int64  `tfsdk:"timeout_ms" json:"timeout_ms,omitempty"`
		Tls             *bool   `tfsdk:"tls" json:"tls,omitempty"`
		Use_websocket   *bool   `tfsdk:"use_websocket" json:"use_websocket,omitempty"`
		V3StatsName     *string `tfsdk:"v3_stats_name" json:"v3StatsName,omitempty"`
		V3health_checks *[]struct {
			Health_check *struct {
				Grpc *struct {
					Authority     *string `tfsdk:"authority" json:"authority,omitempty"`
					Upstream_name *string `tfsdk:"upstream_name" json:"upstream_name,omitempty"`
				} `tfsdk:"grpc" json:"grpc,omitempty"`
				Http *struct {
					Add_request_headers *struct {
						Append           *bool   `tfsdk:"append" json:"append,omitempty"`
						V2Representation *string `tfsdk:"v2_representation" json:"v2Representation,omitempty"`
						Value            *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"add_request_headers" json:"add_request_headers,omitempty"`
					Expected_statuses *[]struct {
						Max *int64 `tfsdk:"max" json:"max,omitempty"`
						Min *int64 `tfsdk:"min" json:"min,omitempty"`
					} `tfsdk:"expected_statuses" json:"expected_statuses,omitempty"`
					Hostname               *string   `tfsdk:"hostname" json:"hostname,omitempty"`
					Path                   *string   `tfsdk:"path" json:"path,omitempty"`
					Remove_request_headers *[]string `tfsdk:"remove_request_headers" json:"remove_request_headers,omitempty"`
				} `tfsdk:"http" json:"http,omitempty"`
			} `tfsdk:"health_check" json:"health_check,omitempty"`
			Healthy_threshold   *int64  `tfsdk:"healthy_threshold" json:"healthy_threshold,omitempty"`
			Interval            *string `tfsdk:"interval" json:"interval,omitempty"`
			Timeout             *string `tfsdk:"timeout" json:"timeout,omitempty"`
			Unhealthy_threshold *int64  `tfsdk:"unhealthy_threshold" json:"unhealthy_threshold,omitempty"`
		} `tfsdk:"v3health_checks" json:"v3health_checks,omitempty"`
		Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GetambassadorIoMappingV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_getambassador_io_mapping_v1_manifest"
}

func (r *GetambassadorIoMappingV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Mapping is the Schema for the mappings API",
		MarkdownDescription: "Mapping is the Schema for the mappings API",
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
				Description:         "MappingSpec defines the desired state of Mapping",
				MarkdownDescription: "MappingSpec defines the desired state of Mapping",
				Attributes: map[string]schema.Attribute{
					"add_linkerd_headers": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"add_request_headers": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"add_response_headers": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"allow_upgrade": schema.ListAttribute{
						Description:         "A case-insensitive list of the non-HTTP protocols to allow 'upgrading' to from HTTP via the 'Connection: upgrade' mechanism[1].  After the upgrade, Ambassador does not interpret the traffic, and behaves similarly to how it does for TCPMappings.  [1]: https://tools.ietf.org/html/rfc7230#section-6.7  For example, if your upstream service supports WebSockets, you would write  allow_upgrade: - websocket  Or if your upstream service supports upgrading from HTTP to SPDY (as the Kubernetes apiserver does for 'kubectl exec' functionality), you would write  allow_upgrade: - spdy/3.1",
						MarkdownDescription: "A case-insensitive list of the non-HTTP protocols to allow 'upgrading' to from HTTP via the 'Connection: upgrade' mechanism[1].  After the upgrade, Ambassador does not interpret the traffic, and behaves similarly to how it does for TCPMappings.  [1]: https://tools.ietf.org/html/rfc7230#section-6.7  For example, if your upstream service supports WebSockets, you would write  allow_upgrade: - websocket  Or if your upstream service supports upgrading from HTTP to SPDY (as the Kubernetes apiserver does for 'kubectl exec' functionality), you would write  allow_upgrade: - spdy/3.1",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ambassador_id": schema.ListAttribute{
						Description:         "AmbassadorID declares which Ambassador instances should pay attention to this resource.  May either be a string or a list of strings.  If no value is provided, the default is:  ambassador_id: - 'default'",
						MarkdownDescription: "AmbassadorID declares which Ambassador instances should pay attention to this resource.  May either be a string or a list of strings.  If no value is provided, the default is:  ambassador_id: - 'default'",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"auth_context_extensions": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"auto_host_rewrite": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bypass_auth": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bypass_error_response_overrides": schema.BoolAttribute{
						Description:         "If true, bypasses any 'error_response_overrides' set on the Ambassador module.",
						MarkdownDescription: "If true, bypasses any 'error_response_overrides' set on the Ambassador module.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"case_sensitive": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"circuit_breakers": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"max_connections": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"max_pending_requests": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"max_requests": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"max_retries": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"priority": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("default", "high"),
									},
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"cluster_idle_timeout_ms": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cluster_max_connection_lifetime_ms": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cluster_tag": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"connect_timeout_ms": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cors": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"credentials": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"exposed_headers": schema.ListAttribute{
								Description:         "StringOrStringList is just what it says on the tin, but note that it will always marshal as a list of strings right now.",
								MarkdownDescription: "StringOrStringList is just what it says on the tin, but note that it will always marshal as a list of strings right now.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"headers": schema.ListAttribute{
								Description:         "StringOrStringList is just what it says on the tin, but note that it will always marshal as a list of strings right now.",
								MarkdownDescription: "StringOrStringList is just what it says on the tin, but note that it will always marshal as a list of strings right now.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_age": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"methods": schema.ListAttribute{
								Description:         "StringOrStringList is just what it says on the tin, but note that it will always marshal as a list of strings right now.",
								MarkdownDescription: "StringOrStringList is just what it says on the tin, but note that it will always marshal as a list of strings right now.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"origins": schema.ListAttribute{
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

					"dns_type": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"docs": schema.SingleNestedAttribute{
						Description:         "DocsInfo provides some extra information about the docs for the Mapping (used by the Dev Portal)",
						MarkdownDescription: "DocsInfo provides some extra information about the docs for the Mapping (used by the Dev Portal)",
						Attributes: map[string]schema.Attribute{
							"display_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ignored": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"path": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"timeout_ms": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"url": schema.StringAttribute{
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

					"enable_ipv4": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_ipv6": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"envoy_override": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"error_response_overrides": schema.ListNestedAttribute{
						Description:         "Error response overrides for this Mapping. Replaces all of the 'error_response_overrides' set on the Ambassador module, if any.",
						MarkdownDescription: "Error response overrides for this Mapping. Replaces all of the 'error_response_overrides' set on the Ambassador module, if any.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"body": schema.SingleNestedAttribute{
									Description:         "The new response body",
									MarkdownDescription: "The new response body",
									Attributes: map[string]schema.Attribute{
										"content_type": schema.StringAttribute{
											Description:         "The content type to set on the error response body when using text_format or text_format_source. Defaults to 'text/plain'.",
											MarkdownDescription: "The content type to set on the error response body when using text_format or text_format_source. Defaults to 'text/plain'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_format": schema.MapAttribute{
											Description:         "A JSON response with content-type: application/json. The values can contain format text like in text_format.",
											MarkdownDescription: "A JSON response with content-type: application/json. The values can contain format text like in text_format.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"text_format": schema.StringAttribute{
											Description:         "A format string representing a text response body. Content-Type can be set using the 'content_type' field below.",
											MarkdownDescription: "A format string representing a text response body. Content-Type can be set using the 'content_type' field below.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"text_format_source": schema.SingleNestedAttribute{
											Description:         "A format string sourced from a file on the Ambassador container. Useful for larger response bodies that should not be placed inline in configuration.",
											MarkdownDescription: "A format string sourced from a file on the Ambassador container. Useful for larger response bodies that should not be placed inline in configuration.",
											Attributes: map[string]schema.Attribute{
												"filename": schema.StringAttribute{
													Description:         "The name of a file on the Ambassador pod that contains a format text string.",
													MarkdownDescription: "The name of a file on the Ambassador pod that contains a format text string.",
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

								"on_status_code": schema.Int64Attribute{
									Description:         "The status code to match on -- not a pointer because it's required.",
									MarkdownDescription: "The status code to match on -- not a pointer because it's required.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(400),
										int64validator.AtMost(599),
									},
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"grpc": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"headers": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"host": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"host_redirect": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"host_regex": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"host_rewrite": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"idle_timeout_ms": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"keepalive": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"idle_time": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"interval": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"probes": schema.Int64Attribute{
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

					"labels": schema.MapAttribute{
						Description:         "A DomainMap is the overall Mapping.spec.Labels type. It maps domains (kind of like namespaces for Mapping labels) to arrays of label groups.",
						MarkdownDescription: "A DomainMap is the overall Mapping.spec.Labels type. It maps domains (kind of like namespaces for Mapping labels) to arrays of label groups.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"load_balancer": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"cookie": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"path": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ttl": schema.StringAttribute{
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

							"header": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"policy": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("round_robin", "ring_hash", "maglev", "least_request"),
								},
							},

							"source_ip": schema.BoolAttribute{
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

					"method": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"method_regex": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"modules": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.MapType{ElemType: types.StringType},
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"outlier_detection": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"path_redirect": schema.StringAttribute{
						Description:         "Path replacement to use when generating an HTTP redirect. Used with 'host_redirect'.",
						MarkdownDescription: "Path replacement to use when generating an HTTP redirect. Used with 'host_redirect'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"precedence": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"prefix": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"prefix_exact": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"prefix_redirect": schema.StringAttribute{
						Description:         "Prefix rewrite to use when generating an HTTP redirect. Used with 'host_redirect'.",
						MarkdownDescription: "Prefix rewrite to use when generating an HTTP redirect. Used with 'host_redirect'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"prefix_regex": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"priority": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"query_parameters": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"redirect_response_code": schema.Int64Attribute{
						Description:         "The response code to use when generating an HTTP redirect. Defaults to 301. Used with 'host_redirect'.",
						MarkdownDescription: "The response code to use when generating an HTTP redirect. Defaults to 301. Used with 'host_redirect'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.OneOf(301, 302, 303, 307, 308),
						},
					},

					"regex_headers": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"regex_query_parameters": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"regex_redirect": schema.SingleNestedAttribute{
						Description:         "Prefix regex rewrite to use when generating an HTTP redirect. Used with 'host_redirect'.",
						MarkdownDescription: "Prefix regex rewrite to use when generating an HTTP redirect. Used with 'host_redirect'.",
						Attributes: map[string]schema.Attribute{
							"pattern": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"substitution": schema.StringAttribute{
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

					"regex_rewrite": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"pattern": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"substitution": schema.StringAttribute{
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

					"remove_request_headers": schema.ListAttribute{
						Description:         "StringOrStringList is just what it says on the tin, but note that it will always marshal as a list of strings right now.",
						MarkdownDescription: "StringOrStringList is just what it says on the tin, but note that it will always marshal as a list of strings right now.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"remove_response_headers": schema.ListAttribute{
						Description:         "StringOrStringList is just what it says on the tin, but note that it will always marshal as a list of strings right now.",
						MarkdownDescription: "StringOrStringList is just what it says on the tin, but note that it will always marshal as a list of strings right now.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"resolver": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"respect_dns_ttl": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"retry_policy": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"num_retries": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"per_try_timeout": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"retry_on": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("5xx", "gateway-error", "connect-failure", "retriable-4xx", "refused-stream", "retriable-status-codes"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"rewrite": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"shadow": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"timeout_ms": schema.Int64Attribute{
						Description:         "The timeout for requests that use this Mapping. Overrides 'cluster_request_timeout_ms' set on the Ambassador Module, if it exists.",
						MarkdownDescription: "The timeout for requests that use this Mapping. Overrides 'cluster_request_timeout_ms' set on the Ambassador Module, if it exists.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tls": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"use_websocket": schema.BoolAttribute{
						Description:         "use_websocket is deprecated, and is equivlaent to setting 'allow_upgrade: ['websocket']'",
						MarkdownDescription: "use_websocket is deprecated, and is equivlaent to setting 'allow_upgrade: ['websocket']'",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"v3_stats_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"v3health_checks": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"health_check": schema.SingleNestedAttribute{
									Description:         "Configuration for where the healthcheck request should be made to",
									MarkdownDescription: "Configuration for where the healthcheck request should be made to",
									Attributes: map[string]schema.Attribute{
										"grpc": schema.SingleNestedAttribute{
											Description:         "HealthCheck for gRPC upstreams. Only one of grpc_health_check or http_health_check may be specified",
											MarkdownDescription: "HealthCheck for gRPC upstreams. Only one of grpc_health_check or http_health_check may be specified",
											Attributes: map[string]schema.Attribute{
												"authority": schema.StringAttribute{
													Description:         "The value of the :authority header in the gRPC health check request. If left empty the upstream name will be used.",
													MarkdownDescription: "The value of the :authority header in the gRPC health check request. If left empty the upstream name will be used.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"upstream_name": schema.StringAttribute{
													Description:         "The upstream name parameter which will be sent to gRPC service in the health check message",
													MarkdownDescription: "The upstream name parameter which will be sent to gRPC service in the health check message",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
											Validators: []UNKNOWN{
												schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("http")),
											},
										},

										"http": schema.SingleNestedAttribute{
											Description:         "HealthCheck for HTTP upstreams. Only one of http_health_check or grpc_health_check may be specified",
											MarkdownDescription: "HealthCheck for HTTP upstreams. Only one of http_health_check or grpc_health_check may be specified",
											Attributes: map[string]schema.Attribute{
												"add_request_headers": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"append": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"v2_representation": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("", "string", "null"),
															},
														},

														"value": schema.StringAttribute{
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

												"expected_statuses": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"max": schema.Int64Attribute{
																Description:         "End of the statuses to include. Must be between 100 and 599 (inclusive)",
																MarkdownDescription: "End of the statuses to include. Must be between 100 and 599 (inclusive)",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(100),
																	int64validator.AtMost(599),
																},
															},

															"min": schema.Int64Attribute{
																Description:         "Start of the statuses to include. Must be between 100 and 599 (inclusive)",
																MarkdownDescription: "Start of the statuses to include. Must be between 100 and 599 (inclusive)",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(100),
																	int64validator.AtMost(599),
																},
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"hostname": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"path": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"remove_request_headers": schema.ListAttribute{
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
											Validators: []UNKNOWN{
												schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("grpc")),
											},
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"healthy_threshold": schema.Int64Attribute{
									Description:         "Number of expected responses for the upstream to be considered healthy. Defaults to 1.",
									MarkdownDescription: "Number of expected responses for the upstream to be considered healthy. Defaults to 1.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"interval": schema.StringAttribute{
									Description:         "Interval between health checks. Defaults to every 5 seconds.",
									MarkdownDescription: "Interval between health checks. Defaults to every 5 seconds.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"timeout": schema.StringAttribute{
									Description:         "Timeout for connecting to the health checking endpoint. Defaults to 3 seconds.",
									MarkdownDescription: "Timeout for connecting to the health checking endpoint. Defaults to 3 seconds.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"unhealthy_threshold": schema.Int64Attribute{
									Description:         "Number of non-expected responses for the upstream to be considered unhealthy. A single 503 will mark the upstream as unhealthy regardless of the threshold. Defaults to 2.",
									MarkdownDescription: "Number of non-expected responses for the upstream to be considered unhealthy. A single 503 will mark the upstream as unhealthy regardless of the threshold. Defaults to 2.",
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

					"weight": schema.Int64Attribute{
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
	}
}

func (r *GetambassadorIoMappingV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_getambassador_io_mapping_v1_manifest")

	var model GetambassadorIoMappingV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("getambassador.io/v1")
	model.Kind = pointer.String("Mapping")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"YAML Error: "+err.Error(),
		)
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
