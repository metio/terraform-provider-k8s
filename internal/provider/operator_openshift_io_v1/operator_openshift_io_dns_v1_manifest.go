/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package operator_openshift_io_v1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &OperatorOpenshiftIoDnsV1Manifest{}
)

func NewOperatorOpenshiftIoDnsV1Manifest() datasource.DataSource {
	return &OperatorOpenshiftIoDnsV1Manifest{}
}

type OperatorOpenshiftIoDnsV1Manifest struct{}

type OperatorOpenshiftIoDnsV1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Cache *struct {
			NegativeTTL *string `tfsdk:"negative_ttl" json:"negativeTTL,omitempty"`
			PositiveTTL *string `tfsdk:"positive_ttl" json:"positiveTTL,omitempty"`
		} `tfsdk:"cache" json:"cache,omitempty"`
		LogLevel        *string `tfsdk:"log_level" json:"logLevel,omitempty"`
		ManagementState *string `tfsdk:"management_state" json:"managementState,omitempty"`
		NodePlacement   *struct {
			NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
			Tolerations  *[]struct {
				Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
				Key               *string `tfsdk:"key" json:"key,omitempty"`
				Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
				TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
				Value             *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"tolerations" json:"tolerations,omitempty"`
		} `tfsdk:"node_placement" json:"nodePlacement,omitempty"`
		OperatorLogLevel *string `tfsdk:"operator_log_level" json:"operatorLogLevel,omitempty"`
		Servers          *[]struct {
			ForwardPlugin *struct {
				Policy           *string `tfsdk:"policy" json:"policy,omitempty"`
				ProtocolStrategy *string `tfsdk:"protocol_strategy" json:"protocolStrategy,omitempty"`
				TransportConfig  *struct {
					Tls *struct {
						CaBundle *struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"ca_bundle" json:"caBundle,omitempty"`
						ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
					} `tfsdk:"tls" json:"tls,omitempty"`
					Transport *string `tfsdk:"transport" json:"transport,omitempty"`
				} `tfsdk:"transport_config" json:"transportConfig,omitempty"`
				Upstreams *[]string `tfsdk:"upstreams" json:"upstreams,omitempty"`
			} `tfsdk:"forward_plugin" json:"forwardPlugin,omitempty"`
			Name  *string   `tfsdk:"name" json:"name,omitempty"`
			Zones *[]string `tfsdk:"zones" json:"zones,omitempty"`
		} `tfsdk:"servers" json:"servers,omitempty"`
		UpstreamResolvers *struct {
			Policy           *string `tfsdk:"policy" json:"policy,omitempty"`
			ProtocolStrategy *string `tfsdk:"protocol_strategy" json:"protocolStrategy,omitempty"`
			TransportConfig  *struct {
				Tls *struct {
					CaBundle *struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"ca_bundle" json:"caBundle,omitempty"`
					ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
				} `tfsdk:"tls" json:"tls,omitempty"`
				Transport *string `tfsdk:"transport" json:"transport,omitempty"`
			} `tfsdk:"transport_config" json:"transportConfig,omitempty"`
			Upstreams *[]struct {
				Address *string `tfsdk:"address" json:"address,omitempty"`
				Port    *int64  `tfsdk:"port" json:"port,omitempty"`
				Type    *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"upstreams" json:"upstreams,omitempty"`
		} `tfsdk:"upstream_resolvers" json:"upstreamResolvers,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OperatorOpenshiftIoDnsV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operator_openshift_io_dns_v1_manifest"
}

func (r *OperatorOpenshiftIoDnsV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DNS manages the CoreDNS component to provide a name resolution service for pods and services in the cluster.  This supports the DNS-based service discovery specification: https://github.com/kubernetes/dns/blob/master/docs/specification.md  More details: https://kubernetes.io/docs/tasks/administer-cluster/coredns  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		MarkdownDescription: "DNS manages the CoreDNS component to provide a name resolution service for pods and services in the cluster.  This supports the DNS-based service discovery specification: https://github.com/kubernetes/dns/blob/master/docs/specification.md  More details: https://kubernetes.io/docs/tasks/administer-cluster/coredns  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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
				Description:         "spec is the specification of the desired behavior of the DNS.",
				MarkdownDescription: "spec is the specification of the desired behavior of the DNS.",
				Attributes: map[string]schema.Attribute{
					"cache": schema.SingleNestedAttribute{
						Description:         "cache describes the caching configuration that applies to all server blocks listed in the Corefile. This field allows a cluster admin to optionally configure: * positiveTTL which is a duration for which positive responses should be cached. * negativeTTL which is a duration for which negative responses should be cached. If this is not configured, OpenShift will configure positive and negative caching with a default value that is subject to change. At the time of writing, the default positiveTTL is 900 seconds and the default negativeTTL is 30 seconds or as noted in the respective Corefile for your version of OpenShift.",
						MarkdownDescription: "cache describes the caching configuration that applies to all server blocks listed in the Corefile. This field allows a cluster admin to optionally configure: * positiveTTL which is a duration for which positive responses should be cached. * negativeTTL which is a duration for which negative responses should be cached. If this is not configured, OpenShift will configure positive and negative caching with a default value that is subject to change. At the time of writing, the default positiveTTL is 900 seconds and the default negativeTTL is 30 seconds or as noted in the respective Corefile for your version of OpenShift.",
						Attributes: map[string]schema.Attribute{
							"negative_ttl": schema.StringAttribute{
								Description:         "negativeTTL is optional and specifies the amount of time that a negative response should be cached.  If configured, it must be a value of 1s (1 second) or greater up to a theoretical maximum of several years. This field expects an unsigned duration string of decimal numbers, each with optional fraction and a unit suffix, e.g. '100s', '1m30s', '12h30m10s'. Values that are fractions of a second are rounded down to the nearest second. If the configured value is less than 1s, the default value will be used. If not configured, the value will be 0s and OpenShift will use a default value of 30 seconds unless noted otherwise in the respective Corefile for your version of OpenShift. The default value of 30 seconds is subject to change.",
								MarkdownDescription: "negativeTTL is optional and specifies the amount of time that a negative response should be cached.  If configured, it must be a value of 1s (1 second) or greater up to a theoretical maximum of several years. This field expects an unsigned duration string of decimal numbers, each with optional fraction and a unit suffix, e.g. '100s', '1m30s', '12h30m10s'. Values that are fractions of a second are rounded down to the nearest second. If the configured value is less than 1s, the default value will be used. If not configured, the value will be 0s and OpenShift will use a default value of 30 seconds unless noted otherwise in the respective Corefile for your version of OpenShift. The default value of 30 seconds is subject to change.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(0|([0-9]+(\.[0-9]+)?(ns|us|µs|μs|ms|s|m|h))+)$`), ""),
								},
							},

							"positive_ttl": schema.StringAttribute{
								Description:         "positiveTTL is optional and specifies the amount of time that a positive response should be cached.  If configured, it must be a value of 1s (1 second) or greater up to a theoretical maximum of several years. This field expects an unsigned duration string of decimal numbers, each with optional fraction and a unit suffix, e.g. '100s', '1m30s', '12h30m10s'. Values that are fractions of a second are rounded down to the nearest second. If the configured value is less than 1s, the default value will be used. If not configured, the value will be 0s and OpenShift will use a default value of 900 seconds unless noted otherwise in the respective Corefile for your version of OpenShift. The default value of 900 seconds is subject to change.",
								MarkdownDescription: "positiveTTL is optional and specifies the amount of time that a positive response should be cached.  If configured, it must be a value of 1s (1 second) or greater up to a theoretical maximum of several years. This field expects an unsigned duration string of decimal numbers, each with optional fraction and a unit suffix, e.g. '100s', '1m30s', '12h30m10s'. Values that are fractions of a second are rounded down to the nearest second. If the configured value is less than 1s, the default value will be used. If not configured, the value will be 0s and OpenShift will use a default value of 900 seconds unless noted otherwise in the respective Corefile for your version of OpenShift. The default value of 900 seconds is subject to change.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(0|([0-9]+(\.[0-9]+)?(ns|us|µs|μs|ms|s|m|h))+)$`), ""),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"log_level": schema.StringAttribute{
						Description:         "logLevel describes the desired logging verbosity for CoreDNS. Any one of the following values may be specified: * Normal logs errors from upstream resolvers. * Debug logs errors, NXDOMAIN responses, and NODATA responses. * Trace logs errors and all responses. Setting logLevel: Trace will produce extremely verbose logs. Valid values are: 'Normal', 'Debug', 'Trace'. Defaults to 'Normal'.",
						MarkdownDescription: "logLevel describes the desired logging verbosity for CoreDNS. Any one of the following values may be specified: * Normal logs errors from upstream resolvers. * Debug logs errors, NXDOMAIN responses, and NODATA responses. * Trace logs errors and all responses. Setting logLevel: Trace will produce extremely verbose logs. Valid values are: 'Normal', 'Debug', 'Trace'. Defaults to 'Normal'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Normal", "Debug", "Trace"),
						},
					},

					"management_state": schema.StringAttribute{
						Description:         "managementState indicates whether the DNS operator should manage cluster DNS",
						MarkdownDescription: "managementState indicates whether the DNS operator should manage cluster DNS",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(Managed|Unmanaged|Force|Removed)$`), ""),
						},
					},

					"node_placement": schema.SingleNestedAttribute{
						Description:         "nodePlacement provides explicit control over the scheduling of DNS pods.  Generally, it is useful to run a DNS pod on every node so that DNS queries are always handled by a local DNS pod instead of going over the network to a DNS pod on another node.  However, security policies may require restricting the placement of DNS pods to specific nodes. For example, if a security policy prohibits pods on arbitrary nodes from communicating with the API, a node selector can be specified to restrict DNS pods to nodes that are permitted to communicate with the API.  Conversely, if running DNS pods on nodes with a particular taint is desired, a toleration can be specified for that taint.  If unset, defaults are used. See nodePlacement for more details.",
						MarkdownDescription: "nodePlacement provides explicit control over the scheduling of DNS pods.  Generally, it is useful to run a DNS pod on every node so that DNS queries are always handled by a local DNS pod instead of going over the network to a DNS pod on another node.  However, security policies may require restricting the placement of DNS pods to specific nodes. For example, if a security policy prohibits pods on arbitrary nodes from communicating with the API, a node selector can be specified to restrict DNS pods to nodes that are permitted to communicate with the API.  Conversely, if running DNS pods on nodes with a particular taint is desired, a toleration can be specified for that taint.  If unset, defaults are used. See nodePlacement for more details.",
						Attributes: map[string]schema.Attribute{
							"node_selector": schema.MapAttribute{
								Description:         "nodeSelector is the node selector applied to DNS pods.  If empty, the default is used, which is currently the following:  kubernetes.io/os: linux  This default is subject to change.  If set, the specified selector is used and replaces the default.",
								MarkdownDescription: "nodeSelector is the node selector applied to DNS pods.  If empty, the default is used, which is currently the following:  kubernetes.io/os: linux  This default is subject to change.  If set, the specified selector is used and replaces the default.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tolerations": schema.ListNestedAttribute{
								Description:         "tolerations is a list of tolerations applied to DNS pods.  If empty, the DNS operator sets a toleration for the 'node-role.kubernetes.io/master' taint.  This default is subject to change.  Specifying tolerations without including a toleration for the 'node-role.kubernetes.io/master' taint may be risky as it could lead to an outage if all worker nodes become unavailable.  Note that the daemon controller adds some tolerations as well.  See https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/",
								MarkdownDescription: "tolerations is a list of tolerations applied to DNS pods.  If empty, the DNS operator sets a toleration for the 'node-role.kubernetes.io/master' taint.  This default is subject to change.  Specifying tolerations without including a toleration for the 'node-role.kubernetes.io/master' taint may be risky as it could lead to an outage if all worker nodes become unavailable.  Note that the daemon controller adds some tolerations as well.  See https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"effect": schema.StringAttribute{
											Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
											MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"key": schema.StringAttribute{
											Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
											MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
											MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"toleration_seconds": schema.Int64Attribute{
											Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
											MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
											MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"operator_log_level": schema.StringAttribute{
						Description:         "operatorLogLevel controls the logging level of the DNS Operator. Valid values are: 'Normal', 'Debug', 'Trace'. Defaults to 'Normal'. setting operatorLogLevel: Trace will produce extremely verbose logs.",
						MarkdownDescription: "operatorLogLevel controls the logging level of the DNS Operator. Valid values are: 'Normal', 'Debug', 'Trace'. Defaults to 'Normal'. setting operatorLogLevel: Trace will produce extremely verbose logs.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Normal", "Debug", "Trace"),
						},
					},

					"servers": schema.ListNestedAttribute{
						Description:         "servers is a list of DNS resolvers that provide name query delegation for one or more subdomains outside the scope of the cluster domain. If servers consists of more than one Server, longest suffix match will be used to determine the Server.  For example, if there are two Servers, one for 'foo.com' and another for 'a.foo.com', and the name query is for 'www.a.foo.com', it will be routed to the Server with Zone 'a.foo.com'.  If this field is nil, no servers are created.",
						MarkdownDescription: "servers is a list of DNS resolvers that provide name query delegation for one or more subdomains outside the scope of the cluster domain. If servers consists of more than one Server, longest suffix match will be used to determine the Server.  For example, if there are two Servers, one for 'foo.com' and another for 'a.foo.com', and the name query is for 'www.a.foo.com', it will be routed to the Server with Zone 'a.foo.com'.  If this field is nil, no servers are created.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"forward_plugin": schema.SingleNestedAttribute{
									Description:         "forwardPlugin defines a schema for configuring CoreDNS to proxy DNS messages to upstream resolvers.",
									MarkdownDescription: "forwardPlugin defines a schema for configuring CoreDNS to proxy DNS messages to upstream resolvers.",
									Attributes: map[string]schema.Attribute{
										"policy": schema.StringAttribute{
											Description:         "policy is used to determine the order in which upstream servers are selected for querying. Any one of the following values may be specified:  * 'Random' picks a random upstream server for each query. * 'RoundRobin' picks upstream servers in a round-robin order, moving to the next server for each new query. * 'Sequential' tries querying upstream servers in a sequential order until one responds, starting with the first server for each new query.  The default value is 'Random'",
											MarkdownDescription: "policy is used to determine the order in which upstream servers are selected for querying. Any one of the following values may be specified:  * 'Random' picks a random upstream server for each query. * 'RoundRobin' picks upstream servers in a round-robin order, moving to the next server for each new query. * 'Sequential' tries querying upstream servers in a sequential order until one responds, starting with the first server for each new query.  The default value is 'Random'",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Random", "RoundRobin", "Sequential"),
											},
										},

										"protocol_strategy": schema.StringAttribute{
											Description:         "protocolStrategy specifies the protocol to use for upstream DNS requests. Valid values for protocolStrategy are 'TCP' and omitted. When omitted, this means no opinion and the platform is left to choose a reasonable default, which is subject to change over time. The current default is to use the protocol of the original client request. 'TCP' specifies that the platform should use TCP for all upstream DNS requests, even if the client request uses UDP. 'TCP' is useful for UDP-specific issues such as those created by non-compliant upstream resolvers, but may consume more bandwidth or increase DNS response time. Note that protocolStrategy only affects the protocol of DNS requests that CoreDNS makes to upstream resolvers. It does not affect the protocol of DNS requests between clients and CoreDNS.",
											MarkdownDescription: "protocolStrategy specifies the protocol to use for upstream DNS requests. Valid values for protocolStrategy are 'TCP' and omitted. When omitted, this means no opinion and the platform is left to choose a reasonable default, which is subject to change over time. The current default is to use the protocol of the original client request. 'TCP' specifies that the platform should use TCP for all upstream DNS requests, even if the client request uses UDP. 'TCP' is useful for UDP-specific issues such as those created by non-compliant upstream resolvers, but may consume more bandwidth or increase DNS response time. Note that protocolStrategy only affects the protocol of DNS requests that CoreDNS makes to upstream resolvers. It does not affect the protocol of DNS requests between clients and CoreDNS.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("TCP", ""),
											},
										},

										"transport_config": schema.SingleNestedAttribute{
											Description:         "transportConfig is used to configure the transport type, server name, and optional custom CA or CA bundle to use when forwarding DNS requests to an upstream resolver.  The default value is '' (empty) which results in a standard cleartext connection being used when forwarding DNS requests to an upstream resolver.",
											MarkdownDescription: "transportConfig is used to configure the transport type, server name, and optional custom CA or CA bundle to use when forwarding DNS requests to an upstream resolver.  The default value is '' (empty) which results in a standard cleartext connection being used when forwarding DNS requests to an upstream resolver.",
											Attributes: map[string]schema.Attribute{
												"tls": schema.SingleNestedAttribute{
													Description:         "tls contains the additional configuration options to use when Transport is set to 'TLS'.",
													MarkdownDescription: "tls contains the additional configuration options to use when Transport is set to 'TLS'.",
													Attributes: map[string]schema.Attribute{
														"ca_bundle": schema.SingleNestedAttribute{
															Description:         "caBundle references a ConfigMap that must contain either a single CA Certificate or a CA Bundle. This allows cluster administrators to provide their own CA or CA bundle for validating the certificate of upstream resolvers.  1. The configmap must contain a 'ca-bundle.crt' key. 2. The value must be a PEM encoded CA certificate or CA bundle. 3. The administrator must create this configmap in the openshift-config namespace. 4. The upstream server certificate must contain a Subject Alternative Name (SAN) that matches ServerName.",
															MarkdownDescription: "caBundle references a ConfigMap that must contain either a single CA Certificate or a CA Bundle. This allows cluster administrators to provide their own CA or CA bundle for validating the certificate of upstream resolvers.  1. The configmap must contain a 'ca-bundle.crt' key. 2. The value must be a PEM encoded CA certificate or CA bundle. 3. The administrator must create this configmap in the openshift-config namespace. 4. The upstream server certificate must contain a Subject Alternative Name (SAN) that matches ServerName.",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "name is the metadata.name of the referenced config map",
																	MarkdownDescription: "name is the metadata.name of the referenced config map",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"server_name": schema.StringAttribute{
															Description:         "serverName is the upstream server to connect to when forwarding DNS queries. This is required when Transport is set to 'TLS'. ServerName will be validated against the DNS naming conventions in RFC 1123 and should match the TLS certificate installed in the upstream resolver(s).",
															MarkdownDescription: "serverName is the upstream server to connect to when forwarding DNS queries. This is required when Transport is set to 'TLS'. ServerName will be validated against the DNS naming conventions in RFC 1123 and should match the TLS certificate installed in the upstream resolver(s).",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtMost(253),
																stringvalidator.RegexMatches(regexp.MustCompile(`^([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])(\.([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9]))*$`), ""),
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"transport": schema.StringAttribute{
													Description:         "transport allows cluster administrators to opt-in to using a DNS-over-TLS connection between cluster DNS and an upstream resolver(s). Configuring TLS as the transport at this level without configuring a CABundle will result in the system certificates being used to verify the serving certificate of the upstream resolver(s).  Possible values: '' (empty) - This means no explicit choice has been made and the platform chooses the default which is subject to change over time. The current default is 'Cleartext'. 'Cleartext' - Cluster admin specified cleartext option. This results in the same functionality as an empty value but may be useful when a cluster admin wants to be more explicit about the transport, or wants to switch from 'TLS' to 'Cleartext' explicitly. 'TLS' - This indicates that DNS queries should be sent over a TLS connection. If Transport is set to TLS, you MUST also set ServerName. If a port is not included with the upstream IP, port 853 will be tried by default per RFC 7858 section 3.1; https://datatracker.ietf.org/doc/html/rfc7858#section-3.1.",
													MarkdownDescription: "transport allows cluster administrators to opt-in to using a DNS-over-TLS connection between cluster DNS and an upstream resolver(s). Configuring TLS as the transport at this level without configuring a CABundle will result in the system certificates being used to verify the serving certificate of the upstream resolver(s).  Possible values: '' (empty) - This means no explicit choice has been made and the platform chooses the default which is subject to change over time. The current default is 'Cleartext'. 'Cleartext' - Cluster admin specified cleartext option. This results in the same functionality as an empty value but may be useful when a cluster admin wants to be more explicit about the transport, or wants to switch from 'TLS' to 'Cleartext' explicitly. 'TLS' - This indicates that DNS queries should be sent over a TLS connection. If Transport is set to TLS, you MUST also set ServerName. If a port is not included with the upstream IP, port 853 will be tried by default per RFC 7858 section 3.1; https://datatracker.ietf.org/doc/html/rfc7858#section-3.1.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("TLS", "Cleartext", ""),
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"upstreams": schema.ListAttribute{
											Description:         "upstreams is a list of resolvers to forward name queries for subdomains of Zones. Each instance of CoreDNS performs health checking of Upstreams. When a healthy upstream returns an error during the exchange, another resolver is tried from Upstreams. The Upstreams are selected in the order specified in Policy. Each upstream is represented by an IP address or IP:port if the upstream listens on a port other than 53.  A maximum of 15 upstreams is allowed per ForwardPlugin.",
											MarkdownDescription: "upstreams is a list of resolvers to forward name queries for subdomains of Zones. Each instance of CoreDNS performs health checking of Upstreams. When a healthy upstream returns an error during the exchange, another resolver is tried from Upstreams. The Upstreams are selected in the order specified in Policy. Each upstream is represented by an IP address or IP:port if the upstream listens on a port other than 53.  A maximum of 15 upstreams is allowed per ForwardPlugin.",
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

								"name": schema.StringAttribute{
									Description:         "name is required and specifies a unique name for the server. Name must comply with the Service Name Syntax of rfc6335.",
									MarkdownDescription: "name is required and specifies a unique name for the server. Name must comply with the Service Name Syntax of rfc6335.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"zones": schema.ListAttribute{
									Description:         "zones is required and specifies the subdomains that Server is authoritative for. Zones must conform to the rfc1123 definition of a subdomain. Specifying the cluster domain (i.e., 'cluster.local') is invalid.",
									MarkdownDescription: "zones is required and specifies the subdomains that Server is authoritative for. Zones must conform to the rfc1123 definition of a subdomain. Specifying the cluster domain (i.e., 'cluster.local') is invalid.",
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

					"upstream_resolvers": schema.SingleNestedAttribute{
						Description:         "upstreamResolvers defines a schema for configuring CoreDNS to proxy DNS messages to upstream resolvers for the case of the default ('.') server  If this field is not specified, the upstream used will default to /etc/resolv.conf, with policy 'sequential'",
						MarkdownDescription: "upstreamResolvers defines a schema for configuring CoreDNS to proxy DNS messages to upstream resolvers for the case of the default ('.') server  If this field is not specified, the upstream used will default to /etc/resolv.conf, with policy 'sequential'",
						Attributes: map[string]schema.Attribute{
							"policy": schema.StringAttribute{
								Description:         "Policy is used to determine the order in which upstream servers are selected for querying. Any one of the following values may be specified:  * 'Random' picks a random upstream server for each query. * 'RoundRobin' picks upstream servers in a round-robin order, moving to the next server for each new query. * 'Sequential' tries querying upstream servers in a sequential order until one responds, starting with the first server for each new query.  The default value is 'Sequential'",
								MarkdownDescription: "Policy is used to determine the order in which upstream servers are selected for querying. Any one of the following values may be specified:  * 'Random' picks a random upstream server for each query. * 'RoundRobin' picks upstream servers in a round-robin order, moving to the next server for each new query. * 'Sequential' tries querying upstream servers in a sequential order until one responds, starting with the first server for each new query.  The default value is 'Sequential'",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Random", "RoundRobin", "Sequential"),
								},
							},

							"protocol_strategy": schema.StringAttribute{
								Description:         "protocolStrategy specifies the protocol to use for upstream DNS requests. Valid values for protocolStrategy are 'TCP' and omitted. When omitted, this means no opinion and the platform is left to choose a reasonable default, which is subject to change over time. The current default is to use the protocol of the original client request. 'TCP' specifies that the platform should use TCP for all upstream DNS requests, even if the client request uses UDP. 'TCP' is useful for UDP-specific issues such as those created by non-compliant upstream resolvers, but may consume more bandwidth or increase DNS response time. Note that protocolStrategy only affects the protocol of DNS requests that CoreDNS makes to upstream resolvers. It does not affect the protocol of DNS requests between clients and CoreDNS.",
								MarkdownDescription: "protocolStrategy specifies the protocol to use for upstream DNS requests. Valid values for protocolStrategy are 'TCP' and omitted. When omitted, this means no opinion and the platform is left to choose a reasonable default, which is subject to change over time. The current default is to use the protocol of the original client request. 'TCP' specifies that the platform should use TCP for all upstream DNS requests, even if the client request uses UDP. 'TCP' is useful for UDP-specific issues such as those created by non-compliant upstream resolvers, but may consume more bandwidth or increase DNS response time. Note that protocolStrategy only affects the protocol of DNS requests that CoreDNS makes to upstream resolvers. It does not affect the protocol of DNS requests between clients and CoreDNS.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("TCP", ""),
								},
							},

							"transport_config": schema.SingleNestedAttribute{
								Description:         "transportConfig is used to configure the transport type, server name, and optional custom CA or CA bundle to use when forwarding DNS requests to an upstream resolver.  The default value is '' (empty) which results in a standard cleartext connection being used when forwarding DNS requests to an upstream resolver.",
								MarkdownDescription: "transportConfig is used to configure the transport type, server name, and optional custom CA or CA bundle to use when forwarding DNS requests to an upstream resolver.  The default value is '' (empty) which results in a standard cleartext connection being used when forwarding DNS requests to an upstream resolver.",
								Attributes: map[string]schema.Attribute{
									"tls": schema.SingleNestedAttribute{
										Description:         "tls contains the additional configuration options to use when Transport is set to 'TLS'.",
										MarkdownDescription: "tls contains the additional configuration options to use when Transport is set to 'TLS'.",
										Attributes: map[string]schema.Attribute{
											"ca_bundle": schema.SingleNestedAttribute{
												Description:         "caBundle references a ConfigMap that must contain either a single CA Certificate or a CA Bundle. This allows cluster administrators to provide their own CA or CA bundle for validating the certificate of upstream resolvers.  1. The configmap must contain a 'ca-bundle.crt' key. 2. The value must be a PEM encoded CA certificate or CA bundle. 3. The administrator must create this configmap in the openshift-config namespace. 4. The upstream server certificate must contain a Subject Alternative Name (SAN) that matches ServerName.",
												MarkdownDescription: "caBundle references a ConfigMap that must contain either a single CA Certificate or a CA Bundle. This allows cluster administrators to provide their own CA or CA bundle for validating the certificate of upstream resolvers.  1. The configmap must contain a 'ca-bundle.crt' key. 2. The value must be a PEM encoded CA certificate or CA bundle. 3. The administrator must create this configmap in the openshift-config namespace. 4. The upstream server certificate must contain a Subject Alternative Name (SAN) that matches ServerName.",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "name is the metadata.name of the referenced config map",
														MarkdownDescription: "name is the metadata.name of the referenced config map",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"server_name": schema.StringAttribute{
												Description:         "serverName is the upstream server to connect to when forwarding DNS queries. This is required when Transport is set to 'TLS'. ServerName will be validated against the DNS naming conventions in RFC 1123 and should match the TLS certificate installed in the upstream resolver(s).",
												MarkdownDescription: "serverName is the upstream server to connect to when forwarding DNS queries. This is required when Transport is set to 'TLS'. ServerName will be validated against the DNS naming conventions in RFC 1123 and should match the TLS certificate installed in the upstream resolver(s).",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(253),
													stringvalidator.RegexMatches(regexp.MustCompile(`^([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])(\.([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9]))*$`), ""),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"transport": schema.StringAttribute{
										Description:         "transport allows cluster administrators to opt-in to using a DNS-over-TLS connection between cluster DNS and an upstream resolver(s). Configuring TLS as the transport at this level without configuring a CABundle will result in the system certificates being used to verify the serving certificate of the upstream resolver(s).  Possible values: '' (empty) - This means no explicit choice has been made and the platform chooses the default which is subject to change over time. The current default is 'Cleartext'. 'Cleartext' - Cluster admin specified cleartext option. This results in the same functionality as an empty value but may be useful when a cluster admin wants to be more explicit about the transport, or wants to switch from 'TLS' to 'Cleartext' explicitly. 'TLS' - This indicates that DNS queries should be sent over a TLS connection. If Transport is set to TLS, you MUST also set ServerName. If a port is not included with the upstream IP, port 853 will be tried by default per RFC 7858 section 3.1; https://datatracker.ietf.org/doc/html/rfc7858#section-3.1.",
										MarkdownDescription: "transport allows cluster administrators to opt-in to using a DNS-over-TLS connection between cluster DNS and an upstream resolver(s). Configuring TLS as the transport at this level without configuring a CABundle will result in the system certificates being used to verify the serving certificate of the upstream resolver(s).  Possible values: '' (empty) - This means no explicit choice has been made and the platform chooses the default which is subject to change over time. The current default is 'Cleartext'. 'Cleartext' - Cluster admin specified cleartext option. This results in the same functionality as an empty value but may be useful when a cluster admin wants to be more explicit about the transport, or wants to switch from 'TLS' to 'Cleartext' explicitly. 'TLS' - This indicates that DNS queries should be sent over a TLS connection. If Transport is set to TLS, you MUST also set ServerName. If a port is not included with the upstream IP, port 853 will be tried by default per RFC 7858 section 3.1; https://datatracker.ietf.org/doc/html/rfc7858#section-3.1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("TLS", "Cleartext", ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"upstreams": schema.ListNestedAttribute{
								Description:         "Upstreams is a list of resolvers to forward name queries for the '.' domain. Each instance of CoreDNS performs health checking of Upstreams. When a healthy upstream returns an error during the exchange, another resolver is tried from Upstreams. The Upstreams are selected in the order specified in Policy.  A maximum of 15 upstreams is allowed per ForwardPlugin. If no Upstreams are specified, /etc/resolv.conf is used by default",
								MarkdownDescription: "Upstreams is a list of resolvers to forward name queries for the '.' domain. Each instance of CoreDNS performs health checking of Upstreams. When a healthy upstream returns an error during the exchange, another resolver is tried from Upstreams. The Upstreams are selected in the order specified in Policy.  A maximum of 15 upstreams is allowed per ForwardPlugin. If no Upstreams are specified, /etc/resolv.conf is used by default",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"address": schema.StringAttribute{
											Description:         "Address must be defined when Type is set to Network. It will be ignored otherwise. It must be a valid ipv4 or ipv6 address.",
											MarkdownDescription: "Address must be defined when Type is set to Network. It will be ignored otherwise. It must be a valid ipv4 or ipv6 address.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"port": schema.Int64Attribute{
											Description:         "Port may be defined when Type is set to Network. It will be ignored otherwise. Port must be between 65535",
											MarkdownDescription: "Port may be defined when Type is set to Network. It will be ignored otherwise. Port must be between 65535",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(1),
												int64validator.AtMost(65535),
											},
										},

										"type": schema.StringAttribute{
											Description:         "Type defines whether this upstream contains an IP/IP:port resolver or the local /etc/resolv.conf. Type accepts 2 possible values: SystemResolvConf or Network.  * When SystemResolvConf is used, the Upstream structure does not require any further fields to be defined: /etc/resolv.conf will be used * When Network is used, the Upstream structure must contain at least an Address",
											MarkdownDescription: "Type defines whether this upstream contains an IP/IP:port resolver or the local /etc/resolv.conf. Type accepts 2 possible values: SystemResolvConf or Network.  * When SystemResolvConf is used, the Upstream structure does not require any further fields to be defined: /etc/resolv.conf will be used * When Network is used, the Upstream structure must contain at least an Address",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("SystemResolvConf", "Network", ""),
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *OperatorOpenshiftIoDnsV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_operator_openshift_io_dns_v1_manifest")

	var model OperatorOpenshiftIoDnsV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("operator.openshift.io/v1")
	model.Kind = pointer.String("DNS")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
