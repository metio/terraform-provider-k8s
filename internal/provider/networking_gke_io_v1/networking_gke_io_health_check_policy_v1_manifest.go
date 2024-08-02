/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package networking_gke_io_v1

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
	_ datasource.DataSource = &NetworkingGkeIoHealthCheckPolicyV1Manifest{}
)

func NewNetworkingGkeIoHealthCheckPolicyV1Manifest() datasource.DataSource {
	return &NetworkingGkeIoHealthCheckPolicyV1Manifest{}
}

type NetworkingGkeIoHealthCheckPolicyV1Manifest struct{}

type NetworkingGkeIoHealthCheckPolicyV1ManifestData struct {
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
		Default *struct {
			CheckIntervalSec *int64 `tfsdk:"check_interval_sec" json:"checkIntervalSec,omitempty"`
			Config           *struct {
				GrpcHealthCheck *struct {
					GrpcServiceName   *string `tfsdk:"grpc_service_name" json:"grpcServiceName,omitempty"`
					Port              *int64  `tfsdk:"port" json:"port,omitempty"`
					PortName          *string `tfsdk:"port_name" json:"portName,omitempty"`
					PortSpecification *string `tfsdk:"port_specification" json:"portSpecification,omitempty"`
				} `tfsdk:"grpc_health_check" json:"grpcHealthCheck,omitempty"`
				Http2HealthCheck *struct {
					Host              *string `tfsdk:"host" json:"host,omitempty"`
					Port              *int64  `tfsdk:"port" json:"port,omitempty"`
					PortName          *string `tfsdk:"port_name" json:"portName,omitempty"`
					PortSpecification *string `tfsdk:"port_specification" json:"portSpecification,omitempty"`
					ProxyHeader       *string `tfsdk:"proxy_header" json:"proxyHeader,omitempty"`
					RequestPath       *string `tfsdk:"request_path" json:"requestPath,omitempty"`
					Response          *string `tfsdk:"response" json:"response,omitempty"`
				} `tfsdk:"http2_health_check" json:"http2HealthCheck,omitempty"`
				HttpHealthCheck *struct {
					Host              *string `tfsdk:"host" json:"host,omitempty"`
					Port              *int64  `tfsdk:"port" json:"port,omitempty"`
					PortName          *string `tfsdk:"port_name" json:"portName,omitempty"`
					PortSpecification *string `tfsdk:"port_specification" json:"portSpecification,omitempty"`
					ProxyHeader       *string `tfsdk:"proxy_header" json:"proxyHeader,omitempty"`
					RequestPath       *string `tfsdk:"request_path" json:"requestPath,omitempty"`
					Response          *string `tfsdk:"response" json:"response,omitempty"`
				} `tfsdk:"http_health_check" json:"httpHealthCheck,omitempty"`
				HttpsHealthCheck *struct {
					Host              *string `tfsdk:"host" json:"host,omitempty"`
					Port              *int64  `tfsdk:"port" json:"port,omitempty"`
					PortName          *string `tfsdk:"port_name" json:"portName,omitempty"`
					PortSpecification *string `tfsdk:"port_specification" json:"portSpecification,omitempty"`
					ProxyHeader       *string `tfsdk:"proxy_header" json:"proxyHeader,omitempty"`
					RequestPath       *string `tfsdk:"request_path" json:"requestPath,omitempty"`
					Response          *string `tfsdk:"response" json:"response,omitempty"`
				} `tfsdk:"https_health_check" json:"httpsHealthCheck,omitempty"`
				TcpHealthCheck *struct {
					Port              *int64  `tfsdk:"port" json:"port,omitempty"`
					PortName          *string `tfsdk:"port_name" json:"portName,omitempty"`
					PortSpecification *string `tfsdk:"port_specification" json:"portSpecification,omitempty"`
					ProxyHeader       *string `tfsdk:"proxy_header" json:"proxyHeader,omitempty"`
					Request           *string `tfsdk:"request" json:"request,omitempty"`
					Response          *string `tfsdk:"response" json:"response,omitempty"`
				} `tfsdk:"tcp_health_check" json:"tcpHealthCheck,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"config" json:"config,omitempty"`
			HealthyThreshold *int64 `tfsdk:"healthy_threshold" json:"healthyThreshold,omitempty"`
			LogConfig        *struct {
				Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
			} `tfsdk:"log_config" json:"logConfig,omitempty"`
			TimeoutSec         *int64 `tfsdk:"timeout_sec" json:"timeoutSec,omitempty"`
			UnhealthyThreshold *int64 `tfsdk:"unhealthy_threshold" json:"unhealthyThreshold,omitempty"`
		} `tfsdk:"default" json:"default,omitempty"`
		TargetRef *struct {
			Group     *string `tfsdk:"group" json:"group,omitempty"`
			Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"target_ref" json:"targetRef,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *NetworkingGkeIoHealthCheckPolicyV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_networking_gke_io_health_check_policy_v1_manifest"
}

func (r *NetworkingGkeIoHealthCheckPolicyV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "HealthCheckPolicy provides a way to create and attach a HealthCheck to a BackendService with the GKE implementation of the Gateway API. This policy can only be attached to a BackendService.",
		MarkdownDescription: "HealthCheckPolicy provides a way to create and attach a HealthCheck to a BackendService with the GKE implementation of the Gateway API. This policy can only be attached to a BackendService.",
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
				Description:         "Spec defines the desired state of HealthCheckPolicy.",
				MarkdownDescription: "Spec defines the desired state of HealthCheckPolicy.",
				Attributes: map[string]schema.Attribute{
					"default": schema.SingleNestedAttribute{
						Description:         "Default defines default policy configuration for the targeted resource.",
						MarkdownDescription: "Default defines default policy configuration for the targeted resource.",
						Attributes: map[string]schema.Attribute{
							"check_interval_sec": schema.Int64Attribute{
								Description:         "How often (in seconds) to send a health check. If not specified, a default value of 5 seconds will be used.",
								MarkdownDescription: "How often (in seconds) to send a health check. If not specified, a default value of 5 seconds will be used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(300),
								},
							},

							"config": schema.SingleNestedAttribute{
								Description:         "Specifies the type of the healthCheck, either TCP, HTTP, HTTPS, HTTP2 or GRPC. Exactly one of the protocol-specific health check field must be specified, which must match type field. Config contains per protocol (i.e. HTTP, HTTPS, HTTP2, TCP, GRPC) configuration. If not specified, health check type defaults to HTTP.",
								MarkdownDescription: "Specifies the type of the healthCheck, either TCP, HTTP, HTTPS, HTTP2 or GRPC. Exactly one of the protocol-specific health check field must be specified, which must match type field. Config contains per protocol (i.e. HTTP, HTTPS, HTTP2, TCP, GRPC) configuration. If not specified, health check type defaults to HTTP.",
								Attributes: map[string]schema.Attribute{
									"grpc_health_check": schema.SingleNestedAttribute{
										Description:         "GRPC is the health check configuration of type GRPC.",
										MarkdownDescription: "GRPC is the health check configuration of type GRPC.",
										Attributes: map[string]schema.Attribute{
											"grpc_service_name": schema.StringAttribute{
												Description:         "The gRPC service name for the health check. This field is optional. The value of grpcServiceName has the following meanings by convention: - Empty serviceName means the overall status of all services at the backend. - Non-empty serviceName means the health of that gRPC service, as defined by   the owner of the service. The grpcServiceName can only be ASCII.",
												MarkdownDescription: "The gRPC service name for the health check. This field is optional. The value of grpcServiceName has the following meanings by convention: - Empty serviceName means the overall status of all services at the backend. - Non-empty serviceName means the health of that gRPC service, as defined by   the owner of the service. The grpcServiceName can only be ASCII.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(1024),
													stringvalidator.RegexMatches(regexp.MustCompile(`[\x00-\xFF]+`), ""),
												},
											},

											"port": schema.Int64Attribute{
												Description:         "The TCP port number for the health check request. Valid values are 1 through 65535.",
												MarkdownDescription: "The TCP port number for the health check request. Valid values are 1 through 65535.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
													int64validator.AtMost(65535),
												},
											},

											"port_name": schema.StringAttribute{
												Description:         "Port name as defined in InstanceGroup#NamedPort#name. If both port and portName are defined, port takes precedence.",
												MarkdownDescription: "Port name as defined in InstanceGroup#NamedPort#name. If both port and portName are defined, port takes precedence.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(63),
													stringvalidator.RegexMatches(regexp.MustCompile(`[a-z]([-a-z0-9]*[a-z0-9])?`), ""),
												},
											},

											"port_specification": schema.StringAttribute{
												Description:         "Specifies how port is selected for health checking, can be one of following values:  USE_FIXED_PORT: The port number in port is used for health checking. USE_NAMED_PORT: The portName is used for health checking. USE_SERVING_PORT: For NetworkEndpointGroup, the port specified for each network endpoint is used for health checking. For other backends, the port or named port specified in the Backend Service is used for health checking.  If not specified, Protocol health check follows behavior specified in port and portName fields. If neither Port nor PortName is specified, this defaults to USE_SERVING_PORT.",
												MarkdownDescription: "Specifies how port is selected for health checking, can be one of following values:  USE_FIXED_PORT: The port number in port is used for health checking. USE_NAMED_PORT: The portName is used for health checking. USE_SERVING_PORT: For NetworkEndpointGroup, the port specified for each network endpoint is used for health checking. For other backends, the port or named port specified in the Backend Service is used for health checking.  If not specified, Protocol health check follows behavior specified in port and portName fields. If neither Port nor PortName is specified, this defaults to USE_SERVING_PORT.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("USE_FIXED_PORT", "USE_NAMED_PORT", "USE_SERVING_PORT"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"http2_health_check": schema.SingleNestedAttribute{
										Description:         "HTTP2 is the health check configuration of type HTTP2.",
										MarkdownDescription: "HTTP2 is the health check configuration of type HTTP2.",
										Attributes: map[string]schema.Attribute{
											"host": schema.StringAttribute{
												Description:         "Host is the value of the host header in the HTTP health check request. This matches the RFC 1123 definition of a hostname with 1 notable exception that numeric IP addresses are not allowed. If not specified or left empty, the IP on behalf of which this health check is performed will be used.",
												MarkdownDescription: "Host is the value of the host header in the HTTP health check request. This matches the RFC 1123 definition of a hostname with 1 notable exception that numeric IP addresses are not allowed. If not specified or left empty, the IP on behalf of which this health check is performed will be used.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(2048),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
												},
											},

											"port": schema.Int64Attribute{
												Description:         "The TCP port number for the health check request. Valid values are 1 through 65535.",
												MarkdownDescription: "The TCP port number for the health check request. Valid values are 1 through 65535.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
													int64validator.AtMost(65535),
												},
											},

											"port_name": schema.StringAttribute{
												Description:         "Port name as defined in InstanceGroup#NamedPort#name. If both port and portName are defined, port takes precedence.",
												MarkdownDescription: "Port name as defined in InstanceGroup#NamedPort#name. If both port and portName are defined, port takes precedence.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(63),
													stringvalidator.RegexMatches(regexp.MustCompile(`[a-z]([-a-z0-9]*[a-z0-9])?`), ""),
												},
											},

											"port_specification": schema.StringAttribute{
												Description:         "Specifies how port is selected for health checking, can be one of following values:  USE_FIXED_PORT: The port number in port is used for health checking. USE_NAMED_PORT: The portName is used for health checking. USE_SERVING_PORT: For NetworkEndpointGroup, the port specified for each network endpoint is used for health checking. For other backends, the port or named port specified in the Backend Service is used for health checking.  If not specified, Protocol health check follows behavior specified in port and portName fields. If neither Port nor PortName is specified, this defaults to USE_SERVING_PORT.",
												MarkdownDescription: "Specifies how port is selected for health checking, can be one of following values:  USE_FIXED_PORT: The port number in port is used for health checking. USE_NAMED_PORT: The portName is used for health checking. USE_SERVING_PORT: For NetworkEndpointGroup, the port specified for each network endpoint is used for health checking. For other backends, the port or named port specified in the Backend Service is used for health checking.  If not specified, Protocol health check follows behavior specified in port and portName fields. If neither Port nor PortName is specified, this defaults to USE_SERVING_PORT.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("USE_FIXED_PORT", "USE_NAMED_PORT", "USE_SERVING_PORT"),
												},
											},

											"proxy_header": schema.StringAttribute{
												Description:         "Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. If not specified, this defaults to NONE.",
												MarkdownDescription: "Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. If not specified, this defaults to NONE.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("NONE", "PROXY_V1"),
												},
											},

											"request_path": schema.StringAttribute{
												Description:         "The request path of the HTTP health check request. If not specified or left empty, a default value of '/' is used.",
												MarkdownDescription: "The request path of the HTTP health check request. If not specified or left empty, a default value of '/' is used.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(2048),
													stringvalidator.RegexMatches(regexp.MustCompile(`\/[A-Za-z0-9\/\-._~%!$&'()*+,;=:]*$`), ""),
												},
											},

											"response": schema.StringAttribute{
												Description:         "The string to match anywhere in the first 1024 bytes of the response body. If not specified or left empty, the status code determines health. The response data can only be ASCII.",
												MarkdownDescription: "The string to match anywhere in the first 1024 bytes of the response body. If not specified or left empty, the status code determines health. The response data can only be ASCII.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(1024),
													stringvalidator.RegexMatches(regexp.MustCompile(`[\x00-\xFF]+`), ""),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"http_health_check": schema.SingleNestedAttribute{
										Description:         "HTTP is the health check configuration of type HTTP.",
										MarkdownDescription: "HTTP is the health check configuration of type HTTP.",
										Attributes: map[string]schema.Attribute{
											"host": schema.StringAttribute{
												Description:         "Host is the value of the host header in the HTTP health check request. This matches the RFC 1123 definition of a hostname with 1 notable exception that numeric IP addresses are not allowed. If not specified or left empty, the IP on behalf of which this health check is performed will be used.",
												MarkdownDescription: "Host is the value of the host header in the HTTP health check request. This matches the RFC 1123 definition of a hostname with 1 notable exception that numeric IP addresses are not allowed. If not specified or left empty, the IP on behalf of which this health check is performed will be used.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(2048),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
												},
											},

											"port": schema.Int64Attribute{
												Description:         "The TCP port number for the health check request. Valid values are 1 through 65535.",
												MarkdownDescription: "The TCP port number for the health check request. Valid values are 1 through 65535.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
													int64validator.AtMost(65535),
												},
											},

											"port_name": schema.StringAttribute{
												Description:         "Port name as defined in InstanceGroup#NamedPort#name. If both port and portName are defined, port takes precedence.",
												MarkdownDescription: "Port name as defined in InstanceGroup#NamedPort#name. If both port and portName are defined, port takes precedence.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(63),
													stringvalidator.RegexMatches(regexp.MustCompile(`[a-z]([-a-z0-9]*[a-z0-9])?`), ""),
												},
											},

											"port_specification": schema.StringAttribute{
												Description:         "Specifies how port is selected for health checking, can be one of following values:  USE_FIXED_PORT: The port number in port is used for health checking. USE_NAMED_PORT: The portName is used for health checking. USE_SERVING_PORT: For NetworkEndpointGroup, the port specified for each network endpoint is used for health checking. For other backends, the port or named port specified in the Backend Service is used for health checking.  If not specified, Protocol health check follows behavior specified in port and portName fields. If neither Port nor PortName is specified, this defaults to USE_SERVING_PORT.",
												MarkdownDescription: "Specifies how port is selected for health checking, can be one of following values:  USE_FIXED_PORT: The port number in port is used for health checking. USE_NAMED_PORT: The portName is used for health checking. USE_SERVING_PORT: For NetworkEndpointGroup, the port specified for each network endpoint is used for health checking. For other backends, the port or named port specified in the Backend Service is used for health checking.  If not specified, Protocol health check follows behavior specified in port and portName fields. If neither Port nor PortName is specified, this defaults to USE_SERVING_PORT.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("USE_FIXED_PORT", "USE_NAMED_PORT", "USE_SERVING_PORT"),
												},
											},

											"proxy_header": schema.StringAttribute{
												Description:         "Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. If not specified, this defaults to NONE.",
												MarkdownDescription: "Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. If not specified, this defaults to NONE.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("NONE", "PROXY_V1"),
												},
											},

											"request_path": schema.StringAttribute{
												Description:         "The request path of the HTTP health check request. If not specified or left empty, a default value of '/' is used.",
												MarkdownDescription: "The request path of the HTTP health check request. If not specified or left empty, a default value of '/' is used.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(2048),
													stringvalidator.RegexMatches(regexp.MustCompile(`\/[A-Za-z0-9\/\-._~%!$&'()*+,;=:]*$`), ""),
												},
											},

											"response": schema.StringAttribute{
												Description:         "The string to match anywhere in the first 1024 bytes of the response body. If not specified or left empty, the status code determines health. The response data can only be ASCII.",
												MarkdownDescription: "The string to match anywhere in the first 1024 bytes of the response body. If not specified or left empty, the status code determines health. The response data can only be ASCII.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(1024),
													stringvalidator.RegexMatches(regexp.MustCompile(`[\x00-\xFF]+`), ""),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"https_health_check": schema.SingleNestedAttribute{
										Description:         "HTTPS is the health check configuration of type HTTPS.",
										MarkdownDescription: "HTTPS is the health check configuration of type HTTPS.",
										Attributes: map[string]schema.Attribute{
											"host": schema.StringAttribute{
												Description:         "Host is the value of the host header in the HTTP health check request. This matches the RFC 1123 definition of a hostname with 1 notable exception that numeric IP addresses are not allowed. If not specified or left empty, the IP on behalf of which this health check is performed will be used.",
												MarkdownDescription: "Host is the value of the host header in the HTTP health check request. This matches the RFC 1123 definition of a hostname with 1 notable exception that numeric IP addresses are not allowed. If not specified or left empty, the IP on behalf of which this health check is performed will be used.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(2048),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
												},
											},

											"port": schema.Int64Attribute{
												Description:         "The TCP port number for the health check request. Valid values are 1 through 65535.",
												MarkdownDescription: "The TCP port number for the health check request. Valid values are 1 through 65535.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
													int64validator.AtMost(65535),
												},
											},

											"port_name": schema.StringAttribute{
												Description:         "Port name as defined in InstanceGroup#NamedPort#name. If both port and portName are defined, port takes precedence.",
												MarkdownDescription: "Port name as defined in InstanceGroup#NamedPort#name. If both port and portName are defined, port takes precedence.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(63),
													stringvalidator.RegexMatches(regexp.MustCompile(`[a-z]([-a-z0-9]*[a-z0-9])?`), ""),
												},
											},

											"port_specification": schema.StringAttribute{
												Description:         "Specifies how port is selected for health checking, can be one of following values:  USE_FIXED_PORT: The port number in port is used for health checking. USE_NAMED_PORT: The portName is used for health checking. USE_SERVING_PORT: For NetworkEndpointGroup, the port specified for each network endpoint is used for health checking. For other backends, the port or named port specified in the Backend Service is used for health checking.  If not specified, Protocol health check follows behavior specified in port and portName fields. If neither Port nor PortName is specified, this defaults to USE_SERVING_PORT.",
												MarkdownDescription: "Specifies how port is selected for health checking, can be one of following values:  USE_FIXED_PORT: The port number in port is used for health checking. USE_NAMED_PORT: The portName is used for health checking. USE_SERVING_PORT: For NetworkEndpointGroup, the port specified for each network endpoint is used for health checking. For other backends, the port or named port specified in the Backend Service is used for health checking.  If not specified, Protocol health check follows behavior specified in port and portName fields. If neither Port nor PortName is specified, this defaults to USE_SERVING_PORT.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("USE_FIXED_PORT", "USE_NAMED_PORT", "USE_SERVING_PORT"),
												},
											},

											"proxy_header": schema.StringAttribute{
												Description:         "Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. If not specified, this defaults to NONE.",
												MarkdownDescription: "Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. If not specified, this defaults to NONE.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("NONE", "PROXY_V1"),
												},
											},

											"request_path": schema.StringAttribute{
												Description:         "The request path of the HTTP health check request. If not specified or left empty, a default value of '/' is used.",
												MarkdownDescription: "The request path of the HTTP health check request. If not specified or left empty, a default value of '/' is used.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(2048),
													stringvalidator.RegexMatches(regexp.MustCompile(`\/[A-Za-z0-9\/\-._~%!$&'()*+,;=:]*$`), ""),
												},
											},

											"response": schema.StringAttribute{
												Description:         "The string to match anywhere in the first 1024 bytes of the response body. If not specified or left empty, the status code determines health. The response data can only be ASCII.",
												MarkdownDescription: "The string to match anywhere in the first 1024 bytes of the response body. If not specified or left empty, the status code determines health. The response data can only be ASCII.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(1024),
													stringvalidator.RegexMatches(regexp.MustCompile(`[\x00-\xFF]+`), ""),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"tcp_health_check": schema.SingleNestedAttribute{
										Description:         "TCP is the health check configuration of type TCP.",
										MarkdownDescription: "TCP is the health check configuration of type TCP.",
										Attributes: map[string]schema.Attribute{
											"port": schema.Int64Attribute{
												Description:         "The TCP port number for the health check request. Valid values are 1 through 65535.",
												MarkdownDescription: "The TCP port number for the health check request. Valid values are 1 through 65535.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
													int64validator.AtMost(65535),
												},
											},

											"port_name": schema.StringAttribute{
												Description:         "Port name as defined in InstanceGroup#NamedPort#name. If both port and portName are defined, port takes precedence.",
												MarkdownDescription: "Port name as defined in InstanceGroup#NamedPort#name. If both port and portName are defined, port takes precedence.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(63),
													stringvalidator.RegexMatches(regexp.MustCompile(`[a-z]([-a-z0-9]*[a-z0-9])?`), ""),
												},
											},

											"port_specification": schema.StringAttribute{
												Description:         "Specifies how port is selected for health checking, can be one of following values:  USE_FIXED_PORT: The port number in port is used for health checking. USE_NAMED_PORT: The portName is used for health checking. USE_SERVING_PORT: For NetworkEndpointGroup, the port specified for each network endpoint is used for health checking. For other backends, the port or named port specified in the Backend Service is used for health checking.  If not specified, Protocol health check follows behavior specified in port and portName fields. If neither Port nor PortName is specified, this defaults to USE_SERVING_PORT.",
												MarkdownDescription: "Specifies how port is selected for health checking, can be one of following values:  USE_FIXED_PORT: The port number in port is used for health checking. USE_NAMED_PORT: The portName is used for health checking. USE_SERVING_PORT: For NetworkEndpointGroup, the port specified for each network endpoint is used for health checking. For other backends, the port or named port specified in the Backend Service is used for health checking.  If not specified, Protocol health check follows behavior specified in port and portName fields. If neither Port nor PortName is specified, this defaults to USE_SERVING_PORT.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("USE_FIXED_PORT", "USE_NAMED_PORT", "USE_SERVING_PORT"),
												},
											},

											"proxy_header": schema.StringAttribute{
												Description:         "Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. If not specified, this defaults to NONE.",
												MarkdownDescription: "Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. If not specified, this defaults to NONE.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("NONE", "PROXY_V1"),
												},
											},

											"request": schema.StringAttribute{
												Description:         "The application data to send once the TCP connection has been established. If not specified, this defaults to empty. If both request and response are empty, the connection establishment alone will indicate health. The request data can only be ASCII.",
												MarkdownDescription: "The application data to send once the TCP connection has been established. If not specified, this defaults to empty. If both request and response are empty, the connection establishment alone will indicate health. The request data can only be ASCII.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(1024),
													stringvalidator.RegexMatches(regexp.MustCompile(`[\x00-\xFF]+`), ""),
												},
											},

											"response": schema.StringAttribute{
												Description:         "The bytes to match against the beginning of the response data. If not specified or left empty, any response will indicate health. The response data can only be ASCII.",
												MarkdownDescription: "The bytes to match against the beginning of the response data. If not specified or left empty, any response will indicate health. The response data can only be ASCII.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(1024),
													stringvalidator.RegexMatches(regexp.MustCompile(`[\x00-\xFF]+`), ""),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": schema.StringAttribute{
										Description:         "Specifies the type of the healthCheck, either TCP, HTTP, HTTPS, HTTP2 or GRPC. Exactly one of the protocol-specific health check field must be specified, which must match type field.",
										MarkdownDescription: "Specifies the type of the healthCheck, either TCP, HTTP, HTTPS, HTTP2 or GRPC. Exactly one of the protocol-specific health check field must be specified, which must match type field.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("TCP", "HTTP", "HTTPS", "HTTP2", "GRPC"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"healthy_threshold": schema.Int64Attribute{
								Description:         "A so-far unhealthy instance will be marked healthy after this many consecutive successes. If not specified, a default value of 2 will be used.",
								MarkdownDescription: "A so-far unhealthy instance will be marked healthy after this many consecutive successes. If not specified, a default value of 2 will be used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(10),
								},
							},

							"log_config": schema.SingleNestedAttribute{
								Description:         "LogConfig configures logging on this health check.",
								MarkdownDescription: "LogConfig configures logging on this health check.",
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Description:         "Enabled indicates whether or not to export health check logs. If not specified, this defaults to false, which means health check logging will be disabled.",
										MarkdownDescription: "Enabled indicates whether or not to export health check logs. If not specified, this defaults to false, which means health check logging will be disabled.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"timeout_sec": schema.Int64Attribute{
								Description:         "How long (in seconds) to wait before claiming failure. If not specified, a default value of 5 seconds will be used. It is invalid for timeoutSec to have greater value than checkIntervalSec.",
								MarkdownDescription: "How long (in seconds) to wait before claiming failure. If not specified, a default value of 5 seconds will be used. It is invalid for timeoutSec to have greater value than checkIntervalSec.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(300),
								},
							},

							"unhealthy_threshold": schema.Int64Attribute{
								Description:         "A so-far healthy instance will be marked unhealthy after this many consecutive failures. If not specified, a default value of 2 will be used.",
								MarkdownDescription: "A so-far healthy instance will be marked unhealthy after this many consecutive failures. If not specified, a default value of 2 will be used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(10),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"target_ref": schema.SingleNestedAttribute{
						Description:         "TargetRef identifies an API object to apply policy to.",
						MarkdownDescription: "TargetRef identifies an API object to apply policy to.",
						Attributes: map[string]schema.Attribute{
							"group": schema.StringAttribute{
								Description:         "Group is the group of the target resource.",
								MarkdownDescription: "Group is the group of the target resource.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtMost(253),
									stringvalidator.RegexMatches(regexp.MustCompile(`^$|^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
								},
							},

							"kind": schema.StringAttribute{
								Description:         "Kind is kind of the target resource.",
								MarkdownDescription: "Kind is kind of the target resource.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(63),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z]([-a-zA-Z0-9]*[a-zA-Z0-9])?$`), ""),
								},
							},

							"name": schema.StringAttribute{
								Description:         "Name is the name of the target resource.",
								MarkdownDescription: "Name is the name of the target resource.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(253),
								},
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace is the namespace of the referent. When unspecified, the local namespace is inferred. Even when policy targets a resource in a different namespace, it MUST only apply to traffic originating from the same namespace as the policy.",
								MarkdownDescription: "Namespace is the namespace of the referent. When unspecified, the local namespace is inferred. Even when policy targets a resource in a different namespace, it MUST only apply to traffic originating from the same namespace as the policy.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(63),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
								},
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
	}
}

func (r *NetworkingGkeIoHealthCheckPolicyV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_networking_gke_io_health_check_policy_v1_manifest")

	var model NetworkingGkeIoHealthCheckPolicyV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("networking.gke.io/v1")
	model.Kind = pointer.String("HealthCheckPolicy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
