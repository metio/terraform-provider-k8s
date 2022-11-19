/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"regexp"

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

type GetambassadorIoMappingV3Alpha1Resource struct{}

var (
	_ resource.Resource = (*GetambassadorIoMappingV3Alpha1Resource)(nil)
)

type GetambassadorIoMappingV3Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type GetambassadorIoMappingV3Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Add_linkerd_headers *bool `tfsdk:"add_linkerd_headers" yaml:"add_linkerd_headers,omitempty"`

		Add_request_headers *struct {
			Append *bool `tfsdk:"append" yaml:"append,omitempty"`

			V2Representation *string `tfsdk:"v2_representation" yaml:"v2Representation,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"add_request_headers" yaml:"add_request_headers,omitempty"`

		Add_response_headers *struct {
			Append *bool `tfsdk:"append" yaml:"append,omitempty"`

			V2Representation *string `tfsdk:"v2_representation" yaml:"v2Representation,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"add_response_headers" yaml:"add_response_headers,omitempty"`

		Allow_upgrade *[]string `tfsdk:"allow_upgrade" yaml:"allow_upgrade,omitempty"`

		Ambassador_id *[]string `tfsdk:"ambassador_id" yaml:"ambassador_id,omitempty"`

		Auth_context_extensions *map[string]string `tfsdk:"auth_context_extensions" yaml:"auth_context_extensions,omitempty"`

		Auto_host_rewrite *bool `tfsdk:"auto_host_rewrite" yaml:"auto_host_rewrite,omitempty"`

		Bypass_auth *bool `tfsdk:"bypass_auth" yaml:"bypass_auth,omitempty"`

		Bypass_error_response_overrides *bool `tfsdk:"bypass_error_response_overrides" yaml:"bypass_error_response_overrides,omitempty"`

		Case_sensitive *bool `tfsdk:"case_sensitive" yaml:"case_sensitive,omitempty"`

		Circuit_breakers *[]struct {
			Max_connections *int64 `tfsdk:"max_connections" yaml:"max_connections,omitempty"`

			Max_pending_requests *int64 `tfsdk:"max_pending_requests" yaml:"max_pending_requests,omitempty"`

			Max_requests *int64 `tfsdk:"max_requests" yaml:"max_requests,omitempty"`

			Max_retries *int64 `tfsdk:"max_retries" yaml:"max_retries,omitempty"`

			Priority *string `tfsdk:"priority" yaml:"priority,omitempty"`
		} `tfsdk:"circuit_breakers" yaml:"circuit_breakers,omitempty"`

		Cluster_idle_timeout_ms *int64 `tfsdk:"cluster_idle_timeout_ms" yaml:"cluster_idle_timeout_ms,omitempty"`

		Cluster_max_connection_lifetime_ms *int64 `tfsdk:"cluster_max_connection_lifetime_ms" yaml:"cluster_max_connection_lifetime_ms,omitempty"`

		Cluster_tag *string `tfsdk:"cluster_tag" yaml:"cluster_tag,omitempty"`

		Connect_timeout_ms *int64 `tfsdk:"connect_timeout_ms" yaml:"connect_timeout_ms,omitempty"`

		Cors *struct {
			Credentials *bool `tfsdk:"credentials" yaml:"credentials,omitempty"`

			Exposed_headers *[]string `tfsdk:"exposed_headers" yaml:"exposed_headers,omitempty"`

			Headers *[]string `tfsdk:"headers" yaml:"headers,omitempty"`

			Max_age *string `tfsdk:"max_age" yaml:"max_age,omitempty"`

			Methods *[]string `tfsdk:"methods" yaml:"methods,omitempty"`

			Origins *[]string `tfsdk:"origins" yaml:"origins,omitempty"`

			V2CommaSeparatedOrigins *bool `tfsdk:"v2_comma_separated_origins" yaml:"v2CommaSeparatedOrigins,omitempty"`
		} `tfsdk:"cors" yaml:"cors,omitempty"`

		Dns_type *string `tfsdk:"dns_type" yaml:"dns_type,omitempty"`

		Docs *struct {
			Display_name *string `tfsdk:"display_name" yaml:"display_name,omitempty"`

			Ignored *bool `tfsdk:"ignored" yaml:"ignored,omitempty"`

			Path *string `tfsdk:"path" yaml:"path,omitempty"`

			Timeout_ms *int64 `tfsdk:"timeout_ms" yaml:"timeout_ms,omitempty"`

			Url *string `tfsdk:"url" yaml:"url,omitempty"`
		} `tfsdk:"docs" yaml:"docs,omitempty"`

		Enable_ipv4 *bool `tfsdk:"enable_ipv4" yaml:"enable_ipv4,omitempty"`

		Enable_ipv6 *bool `tfsdk:"enable_ipv6" yaml:"enable_ipv6,omitempty"`

		Envoy_override utilities.Dynamic `tfsdk:"envoy_override" yaml:"envoy_override,omitempty"`

		Error_response_overrides *[]struct {
			Body *struct {
				Content_type *string `tfsdk:"content_type" yaml:"content_type,omitempty"`

				Json_format *map[string]string `tfsdk:"json_format" yaml:"json_format,omitempty"`

				Text_format *string `tfsdk:"text_format" yaml:"text_format,omitempty"`

				Text_format_source *struct {
					Filename *string `tfsdk:"filename" yaml:"filename,omitempty"`
				} `tfsdk:"text_format_source" yaml:"text_format_source,omitempty"`
			} `tfsdk:"body" yaml:"body,omitempty"`

			On_status_code *int64 `tfsdk:"on_status_code" yaml:"on_status_code,omitempty"`
		} `tfsdk:"error_response_overrides" yaml:"error_response_overrides,omitempty"`

		Grpc *bool `tfsdk:"grpc" yaml:"grpc,omitempty"`

		Headers *map[string]string `tfsdk:"headers" yaml:"headers,omitempty"`

		Host *string `tfsdk:"host" yaml:"host,omitempty"`

		Host_redirect *bool `tfsdk:"host_redirect" yaml:"host_redirect,omitempty"`

		Host_regex *bool `tfsdk:"host_regex" yaml:"host_regex,omitempty"`

		Host_rewrite *string `tfsdk:"host_rewrite" yaml:"host_rewrite,omitempty"`

		Hostname *string `tfsdk:"hostname" yaml:"hostname,omitempty"`

		Idle_timeout_ms *int64 `tfsdk:"idle_timeout_ms" yaml:"idle_timeout_ms,omitempty"`

		Keepalive *struct {
			Idle_time *int64 `tfsdk:"idle_time" yaml:"idle_time,omitempty"`

			Interval *int64 `tfsdk:"interval" yaml:"interval,omitempty"`

			Probes *int64 `tfsdk:"probes" yaml:"probes,omitempty"`
		} `tfsdk:"keepalive" yaml:"keepalive,omitempty"`

		Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

		Load_balancer *struct {
			Cookie *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Path *string `tfsdk:"path" yaml:"path,omitempty"`

				Ttl *string `tfsdk:"ttl" yaml:"ttl,omitempty"`
			} `tfsdk:"cookie" yaml:"cookie,omitempty"`

			Header *string `tfsdk:"header" yaml:"header,omitempty"`

			Policy *string `tfsdk:"policy" yaml:"policy,omitempty"`

			Source_ip *bool `tfsdk:"source_ip" yaml:"source_ip,omitempty"`
		} `tfsdk:"load_balancer" yaml:"load_balancer,omitempty"`

		Method *string `tfsdk:"method" yaml:"method,omitempty"`

		Method_regex *bool `tfsdk:"method_regex" yaml:"method_regex,omitempty"`

		Modules *[]map[string]string `tfsdk:"modules" yaml:"modules,omitempty"`

		Outlier_detection *string `tfsdk:"outlier_detection" yaml:"outlier_detection,omitempty"`

		Path_redirect *string `tfsdk:"path_redirect" yaml:"path_redirect,omitempty"`

		Precedence *int64 `tfsdk:"precedence" yaml:"precedence,omitempty"`

		Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

		Prefix_exact *bool `tfsdk:"prefix_exact" yaml:"prefix_exact,omitempty"`

		Prefix_redirect *string `tfsdk:"prefix_redirect" yaml:"prefix_redirect,omitempty"`

		Prefix_regex *bool `tfsdk:"prefix_regex" yaml:"prefix_regex,omitempty"`

		Priority *string `tfsdk:"priority" yaml:"priority,omitempty"`

		Query_parameters *map[string]string `tfsdk:"query_parameters" yaml:"query_parameters,omitempty"`

		Redirect_response_code *int64 `tfsdk:"redirect_response_code" yaml:"redirect_response_code,omitempty"`

		Regex_headers *map[string]string `tfsdk:"regex_headers" yaml:"regex_headers,omitempty"`

		Regex_query_parameters *map[string]string `tfsdk:"regex_query_parameters" yaml:"regex_query_parameters,omitempty"`

		Regex_redirect *struct {
			Pattern *string `tfsdk:"pattern" yaml:"pattern,omitempty"`

			Substitution *string `tfsdk:"substitution" yaml:"substitution,omitempty"`
		} `tfsdk:"regex_redirect" yaml:"regex_redirect,omitempty"`

		Regex_rewrite *struct {
			Pattern *string `tfsdk:"pattern" yaml:"pattern,omitempty"`

			Substitution *string `tfsdk:"substitution" yaml:"substitution,omitempty"`
		} `tfsdk:"regex_rewrite" yaml:"regex_rewrite,omitempty"`

		Remove_request_headers *[]string `tfsdk:"remove_request_headers" yaml:"remove_request_headers,omitempty"`

		Remove_response_headers *[]string `tfsdk:"remove_response_headers" yaml:"remove_response_headers,omitempty"`

		Resolver *string `tfsdk:"resolver" yaml:"resolver,omitempty"`

		Respect_dns_ttl *bool `tfsdk:"respect_dns_ttl" yaml:"respect_dns_ttl,omitempty"`

		Retry_policy *struct {
			Num_retries *int64 `tfsdk:"num_retries" yaml:"num_retries,omitempty"`

			Per_try_timeout *string `tfsdk:"per_try_timeout" yaml:"per_try_timeout,omitempty"`

			Retry_on *string `tfsdk:"retry_on" yaml:"retry_on,omitempty"`
		} `tfsdk:"retry_policy" yaml:"retry_policy,omitempty"`

		Rewrite *string `tfsdk:"rewrite" yaml:"rewrite,omitempty"`

		Service *string `tfsdk:"service" yaml:"service,omitempty"`

		Shadow *bool `tfsdk:"shadow" yaml:"shadow,omitempty"`

		Stats_name *string `tfsdk:"stats_name" yaml:"stats_name,omitempty"`

		Timeout_ms *int64 `tfsdk:"timeout_ms" yaml:"timeout_ms,omitempty"`

		Tls *string `tfsdk:"tls" yaml:"tls,omitempty"`

		Use_websocket *bool `tfsdk:"use_websocket" yaml:"use_websocket,omitempty"`

		V2BoolHeaders *[]string `tfsdk:"v2_bool_headers" yaml:"v2BoolHeaders,omitempty"`

		V2BoolQueryParameters *[]string `tfsdk:"v2_bool_query_parameters" yaml:"v2BoolQueryParameters,omitempty"`

		V2ExplicitTLS *struct {
			ServiceScheme *string `tfsdk:"service_scheme" yaml:"serviceScheme,omitempty"`

			Tls *string `tfsdk:"tls" yaml:"tls,omitempty"`
		} `tfsdk:"v2_explicit_tls" yaml:"v2ExplicitTLS,omitempty"`

		Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewGetambassadorIoMappingV3Alpha1Resource() resource.Resource {
	return &GetambassadorIoMappingV3Alpha1Resource{}
}

func (r *GetambassadorIoMappingV3Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_getambassador_io_mapping_v3alpha1"
}

func (r *GetambassadorIoMappingV3Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Mapping is the Schema for the mappings API",
		MarkdownDescription: "Mapping is the Schema for the mappings API",
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
				Description:         "MappingSpec defines the desired state of Mapping",
				MarkdownDescription: "MappingSpec defines the desired state of Mapping",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"add_linkerd_headers": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"add_request_headers": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"append": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"v2_representation": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("", "string", "null"),
								},
							},

							"value": {
								Description:         "",
								MarkdownDescription: "",

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

					"add_response_headers": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"append": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"v2_representation": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("", "string", "null"),
								},
							},

							"value": {
								Description:         "",
								MarkdownDescription: "",

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

					"allow_upgrade": {
						Description:         "A case-insensitive list of the non-HTTP protocols to allow 'upgrading' to from HTTP via the 'Connection: upgrade' mechanism[1].  After the upgrade, Ambassador does not interpret the traffic, and behaves similarly to how it does for TCPMappings.  [1]: https://tools.ietf.org/html/rfc7230#section-6.7  For example, if your upstream service supports WebSockets, you would write     allow_upgrade:    - websocket  Or if your upstream service supports upgrading from HTTP to SPDY (as the Kubernetes apiserver does for 'kubectl exec' functionality), you would write     allow_upgrade:    - spdy/3.1",
						MarkdownDescription: "A case-insensitive list of the non-HTTP protocols to allow 'upgrading' to from HTTP via the 'Connection: upgrade' mechanism[1].  After the upgrade, Ambassador does not interpret the traffic, and behaves similarly to how it does for TCPMappings.  [1]: https://tools.ietf.org/html/rfc7230#section-6.7  For example, if your upstream service supports WebSockets, you would write     allow_upgrade:    - websocket  Or if your upstream service supports upgrading from HTTP to SPDY (as the Kubernetes apiserver does for 'kubectl exec' functionality), you would write     allow_upgrade:    - spdy/3.1",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ambassador_id": {
						Description:         "AmbassadorID declares which Ambassador instances should pay attention to this resource. If no value is provided, the default is:  	ambassador_id: 	- 'default'  TODO(lukeshu): In v3alpha2, consider renaming all of the 'ambassador_id' (singular) fields to 'ambassador_ids' (plural).",
						MarkdownDescription: "AmbassadorID declares which Ambassador instances should pay attention to this resource. If no value is provided, the default is:  	ambassador_id: 	- 'default'  TODO(lukeshu): In v3alpha2, consider renaming all of the 'ambassador_id' (singular) fields to 'ambassador_ids' (plural).",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"auth_context_extensions": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"auto_host_rewrite": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"bypass_auth": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"bypass_error_response_overrides": {
						Description:         "If true, bypasses any 'error_response_overrides' set on the Ambassador module.",
						MarkdownDescription: "If true, bypasses any 'error_response_overrides' set on the Ambassador module.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"case_sensitive": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"circuit_breakers": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"max_connections": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"max_pending_requests": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"max_requests": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"max_retries": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"priority": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("default", "high"),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"cluster_idle_timeout_ms": {
						Description:         "TODO(lukeshu): In v3alpha2, change all of the '{foo}_ms'/'MillisecondDuration' fields to '{foo}'/'metav1.Duration'.",
						MarkdownDescription: "TODO(lukeshu): In v3alpha2, change all of the '{foo}_ms'/'MillisecondDuration' fields to '{foo}'/'metav1.Duration'.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"cluster_max_connection_lifetime_ms": {
						Description:         "TODO(lukeshu): In v3alpha2, change all of the '{foo}_ms'/'MillisecondDuration' fields to '{foo}'/'metav1.Duration'.",
						MarkdownDescription: "TODO(lukeshu): In v3alpha2, change all of the '{foo}_ms'/'MillisecondDuration' fields to '{foo}'/'metav1.Duration'.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"cluster_tag": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"connect_timeout_ms": {
						Description:         "TODO(lukeshu): In v3alpha2, change all of the '{foo}_ms'/'MillisecondDuration' fields to '{foo}'/'metav1.Duration'.",
						MarkdownDescription: "TODO(lukeshu): In v3alpha2, change all of the '{foo}_ms'/'MillisecondDuration' fields to '{foo}'/'metav1.Duration'.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"cors": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"credentials": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"exposed_headers": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"headers": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"max_age": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"methods": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"origins": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"v2_comma_separated_origins": {
								Description:         "",
								MarkdownDescription: "",

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

					"dns_type": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"docs": {
						Description:         "DocsInfo provides some extra information about the docs for the Mapping. Docs is used by both the agent and the DevPortal.",
						MarkdownDescription: "DocsInfo provides some extra information about the docs for the Mapping. Docs is used by both the agent and the DevPortal.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"display_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ignored": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"path": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"timeout_ms": {
								Description:         "TODO(lukeshu): In v3alpha2, change all of the '{foo}_ms'/'MillisecondDuration' fields to '{foo}'/'metav1.Duration'.",
								MarkdownDescription: "TODO(lukeshu): In v3alpha2, change all of the '{foo}_ms'/'MillisecondDuration' fields to '{foo}'/'metav1.Duration'.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"url": {
								Description:         "",
								MarkdownDescription: "",

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

					"enable_ipv4": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_ipv6": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"envoy_override": {
						Description:         "UntypedDict is relatively opaque as a Go type, but it preserves its contents in a roundtrippable way.",
						MarkdownDescription: "UntypedDict is relatively opaque as a Go type, but it preserves its contents in a roundtrippable way.",

						Type: utilities.DynamicType{},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"error_response_overrides": {
						Description:         "Error response overrides for this Mapping. Replaces all of the 'error_response_overrides' set on the Ambassador module, if any.",
						MarkdownDescription: "Error response overrides for this Mapping. Replaces all of the 'error_response_overrides' set on the Ambassador module, if any.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"body": {
								Description:         "The new response body",
								MarkdownDescription: "The new response body",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"content_type": {
										Description:         "The content type to set on the error response body when using text_format or text_format_source. Defaults to 'text/plain'.",
										MarkdownDescription: "The content type to set on the error response body when using text_format or text_format_source. Defaults to 'text/plain'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"json_format": {
										Description:         "A JSON response with content-type: application/json. The values can contain format text like in text_format.",
										MarkdownDescription: "A JSON response with content-type: application/json. The values can contain format text like in text_format.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"text_format": {
										Description:         "A format string representing a text response body. Content-Type can be set using the 'content_type' field below.",
										MarkdownDescription: "A format string representing a text response body. Content-Type can be set using the 'content_type' field below.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"text_format_source": {
										Description:         "A format string sourced from a file on the Ambassador container. Useful for larger response bodies that should not be placed inline in configuration.",
										MarkdownDescription: "A format string sourced from a file on the Ambassador container. Useful for larger response bodies that should not be placed inline in configuration.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"filename": {
												Description:         "The name of a file on the Ambassador pod that contains a format text string.",
												MarkdownDescription: "The name of a file on the Ambassador pod that contains a format text string.",

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

							"on_status_code": {
								Description:         "The status code to match on -- not a pointer because it's required.",
								MarkdownDescription: "The status code to match on -- not a pointer because it's required.",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(400),

									int64validator.AtMost(599),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"grpc": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"headers": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"host": {
						Description:         "Exact match for the hostname of a request if HostRegex is false; regex match for the hostname if HostRegex is true.  Host specifies both a match for the ':authority' header of a request, as well as a match criterion for Host CRDs: a Mapping that specifies Host will not associate with a Host that doesn't have a matching Hostname.  If both Host and Hostname are set, an error is logged, Host is ignored, and Hostname is used.  DEPRECATED: Host is either an exact match or a regex, depending on HostRegex. Use HostName instead.  TODO(lukeshu): In v3alpha2, get rid of MappingSpec.host and MappingSpec.host_regex in favor of a MappingSpec.deprecated_hostname_regex.",
						MarkdownDescription: "Exact match for the hostname of a request if HostRegex is false; regex match for the hostname if HostRegex is true.  Host specifies both a match for the ':authority' header of a request, as well as a match criterion for Host CRDs: a Mapping that specifies Host will not associate with a Host that doesn't have a matching Hostname.  If both Host and Hostname are set, an error is logged, Host is ignored, and Hostname is used.  DEPRECATED: Host is either an exact match or a regex, depending on HostRegex. Use HostName instead.  TODO(lukeshu): In v3alpha2, get rid of MappingSpec.host and MappingSpec.host_regex in favor of a MappingSpec.deprecated_hostname_regex.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"host_redirect": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"host_regex": {
						Description:         "DEPRECATED: Host is either an exact match or a regex, depending on HostRegex. Use HostName instead.  TODO(lukeshu): In v3alpha2, get rid of MappingSpec.host and MappingSpec.host_regex in favor of a MappingSpec.deprecated_hostname_regex.",
						MarkdownDescription: "DEPRECATED: Host is either an exact match or a regex, depending on HostRegex. Use HostName instead.  TODO(lukeshu): In v3alpha2, get rid of MappingSpec.host and MappingSpec.host_regex in favor of a MappingSpec.deprecated_hostname_regex.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"host_rewrite": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"hostname": {
						Description:         "Hostname is a DNS glob specifying the hosts to which this Mapping applies.  Hostname specifies both a match for the ':authority' header of a request, as well as a match criterion for Host CRDs: a Mapping that specifies Hostname will not associate with a Host that doesn't have a matching Hostname.  If both Host and Hostname are set, an error is logged, Host is ignored, and Hostname is used.",
						MarkdownDescription: "Hostname is a DNS glob specifying the hosts to which this Mapping applies.  Hostname specifies both a match for the ':authority' header of a request, as well as a match criterion for Host CRDs: a Mapping that specifies Hostname will not associate with a Host that doesn't have a matching Hostname.  If both Host and Hostname are set, an error is logged, Host is ignored, and Hostname is used.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"idle_timeout_ms": {
						Description:         "TODO(lukeshu): In v3alpha2, change all of the '{foo}_ms'/'MillisecondDuration' fields to '{foo}'/'metav1.Duration'.",
						MarkdownDescription: "TODO(lukeshu): In v3alpha2, change all of the '{foo}_ms'/'MillisecondDuration' fields to '{foo}'/'metav1.Duration'.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"keepalive": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"idle_time": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"interval": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"probes": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"labels": {
						Description:         "A DomainMap is the overall Mapping.spec.Labels type. It maps domains (kind of like namespaces for Mapping labels) to arrays of label groups.",
						MarkdownDescription: "A DomainMap is the overall Mapping.spec.Labels type. It maps domains (kind of like namespaces for Mapping labels) to arrays of label groups.",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"load_balancer": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"cookie": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"path": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"ttl": {
										Description:         "",
										MarkdownDescription: "",

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

							"header": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"policy": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("round_robin", "ring_hash", "maglev", "least_request"),
								},
							},

							"source_ip": {
								Description:         "",
								MarkdownDescription: "",

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

					"method": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"method_regex": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"modules": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.ListType{ElemType: types.MapType{ElemType: types.StringType}},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"outlier_detection": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"path_redirect": {
						Description:         "Path replacement to use when generating an HTTP redirect. Used with 'host_redirect'.",
						MarkdownDescription: "Path replacement to use when generating an HTTP redirect. Used with 'host_redirect'.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"precedence": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"prefix": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"prefix_exact": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"prefix_redirect": {
						Description:         "Prefix rewrite to use when generating an HTTP redirect. Used with 'host_redirect'.",
						MarkdownDescription: "Prefix rewrite to use when generating an HTTP redirect. Used with 'host_redirect'.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"prefix_regex": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"priority": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"query_parameters": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"redirect_response_code": {
						Description:         "The response code to use when generating an HTTP redirect. Defaults to 301. Used with 'host_redirect'.",
						MarkdownDescription: "The response code to use when generating an HTTP redirect. Defaults to 301. Used with 'host_redirect'.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							int64validator.OneOf(301, 302, 303, 307, 308),
						},
					},

					"regex_headers": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"regex_query_parameters": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"regex_redirect": {
						Description:         "Prefix regex rewrite to use when generating an HTTP redirect. Used with 'host_redirect'.",
						MarkdownDescription: "Prefix regex rewrite to use when generating an HTTP redirect. Used with 'host_redirect'.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"pattern": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"substitution": {
								Description:         "",
								MarkdownDescription: "",

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

					"regex_rewrite": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"pattern": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"substitution": {
								Description:         "",
								MarkdownDescription: "",

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

					"remove_request_headers": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"remove_response_headers": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"resolver": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"respect_dns_ttl": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"retry_policy": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"num_retries": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"per_try_timeout": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"retry_on": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("5xx", "gateway-error", "connect-failure", "retriable-4xx", "refused-stream", "retriable-status-codes"),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"rewrite": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"service": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"shadow": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"stats_name": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"timeout_ms": {
						Description:         "The timeout for requests that use this Mapping. Overrides 'cluster_request_timeout_ms' set on the Ambassador Module, if it exists.",
						MarkdownDescription: "The timeout for requests that use this Mapping. Overrides 'cluster_request_timeout_ms' set on the Ambassador Module, if it exists.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"tls": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"use_websocket": {
						Description:         "use_websocket is deprecated, and is equivlaent to setting 'allow_upgrade: ['websocket']'  TODO(lukeshu): In v3alpha2, get rid of MappingSpec.DeprecatedUseWebsocket.",
						MarkdownDescription: "use_websocket is deprecated, and is equivlaent to setting 'allow_upgrade: ['websocket']'  TODO(lukeshu): In v3alpha2, get rid of MappingSpec.DeprecatedUseWebsocket.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"v2_bool_headers": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"v2_bool_query_parameters": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"v2_explicit_tls": {
						Description:         "V2ExplicitTLS controls some vanity/stylistic elements when converting from v3alpha1 to v2.  The values in an V2ExplicitTLS should not in any way affect the runtime operation of Emissary; except that it may affect internal names in the Envoy config, which may in turn affect stats names.  But it should not affect any end-user observable behavior.",
						MarkdownDescription: "V2ExplicitTLS controls some vanity/stylistic elements when converting from v3alpha1 to v2.  The values in an V2ExplicitTLS should not in any way affect the runtime operation of Emissary; except that it may affect internal names in the Envoy config, which may in turn affect stats names.  But it should not affect any end-user observable behavior.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"service_scheme": {
								Description:         "ServiceScheme specifies how to spell and capitalize the scheme-part of the service URL.  Acceptable values are 'http://' (case-insensitive), 'https://' (case-insensitive), or ''.  The value is used if it agrees with whether or not this resource enables TLS origination, or if something else in the resource overrides the scheme.",
								MarkdownDescription: "ServiceScheme specifies how to spell and capitalize the scheme-part of the service URL.  Acceptable values are 'http://' (case-insensitive), 'https://' (case-insensitive), or ''.  The value is used if it agrees with whether or not this resource enables TLS origination, or if something else in the resource overrides the scheme.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.RegexMatches(regexp.MustCompile(`^([hH][tT][tT][pP][sS]?://)?$`), ""),
								},
							},

							"tls": {
								Description:         "TLS controls whether and how to represent the 'tls' field when its value could be implied by the 'service' field.  In v2, there were a lot of different ways to spell an 'empty' value, and this field specifies which way to spell it (and will therefore only be used if the value will indeed be empty).   | Value        | Representation                        | Meaning of representation          |  |--------------+---------------------------------------+------------------------------------|  | ''           | omit the field                        | defer to service (no TLSContext)   |  | 'null'       | store an explicit 'null' in the field | defer to service (no TLSContext)   |  | 'string'     | store an empty string in the field    | defer to service (no TLSContext)   |  | 'bool:false' | store a Boolean 'false' in the field  | defer to service (no TLSContext)   |  | 'bool:true'  | store a Boolean 'true' in the field   | originate TLS (no TLSContext)      |  If the meaning of the representation contradicts anything else (if a TLSContext is to be used, or in the case of 'bool:true' if TLS is not to be originated), then this field is ignored.",
								MarkdownDescription: "TLS controls whether and how to represent the 'tls' field when its value could be implied by the 'service' field.  In v2, there were a lot of different ways to spell an 'empty' value, and this field specifies which way to spell it (and will therefore only be used if the value will indeed be empty).   | Value        | Representation                        | Meaning of representation          |  |--------------+---------------------------------------+------------------------------------|  | ''           | omit the field                        | defer to service (no TLSContext)   |  | 'null'       | store an explicit 'null' in the field | defer to service (no TLSContext)   |  | 'string'     | store an empty string in the field    | defer to service (no TLSContext)   |  | 'bool:false' | store a Boolean 'false' in the field  | defer to service (no TLSContext)   |  | 'bool:true'  | store a Boolean 'true' in the field   | originate TLS (no TLSContext)      |  If the meaning of the representation contradicts anything else (if a TLSContext is to be used, or in the case of 'bool:true' if TLS is not to be originated), then this field is ignored.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("", "null", "bool:true", "bool:false", "string"),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"weight": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *GetambassadorIoMappingV3Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_getambassador_io_mapping_v3alpha1")

	var state GetambassadorIoMappingV3Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel GetambassadorIoMappingV3Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("getambassador.io/v3alpha1")
	goModel.Kind = utilities.Ptr("Mapping")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *GetambassadorIoMappingV3Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_getambassador_io_mapping_v3alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *GetambassadorIoMappingV3Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_getambassador_io_mapping_v3alpha1")

	var state GetambassadorIoMappingV3Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel GetambassadorIoMappingV3Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("getambassador.io/v3alpha1")
	goModel.Kind = utilities.Ptr("Mapping")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *GetambassadorIoMappingV3Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_getambassador_io_mapping_v3alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
