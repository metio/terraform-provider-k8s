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
	_ datasource.DataSource = &NetworkingGkeIoGcpbackendPolicyV1Manifest{}
)

func NewNetworkingGkeIoGcpbackendPolicyV1Manifest() datasource.DataSource {
	return &NetworkingGkeIoGcpbackendPolicyV1Manifest{}
}

type NetworkingGkeIoGcpbackendPolicyV1Manifest struct{}

type NetworkingGkeIoGcpbackendPolicyV1ManifestData struct {
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
			ConnectionDraining *struct {
				DrainingTimeoutSec *int64 `tfsdk:"draining_timeout_sec" json:"drainingTimeoutSec,omitempty"`
			} `tfsdk:"connection_draining" json:"connectionDraining,omitempty"`
			Iap *struct {
				ClientID           *string `tfsdk:"client_id" json:"clientID,omitempty"`
				Enabled            *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
				Oauth2ClientSecret *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"oauth2_client_secret" json:"oauth2ClientSecret,omitempty"`
			} `tfsdk:"iap" json:"iap,omitempty"`
			Logging *struct {
				Enabled    *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
				SampleRate *int64 `tfsdk:"sample_rate" json:"sampleRate,omitempty"`
			} `tfsdk:"logging" json:"logging,omitempty"`
			SecurityPolicy  *string `tfsdk:"security_policy" json:"securityPolicy,omitempty"`
			SessionAffinity *struct {
				CookieTtlSec *int64  `tfsdk:"cookie_ttl_sec" json:"cookieTtlSec,omitempty"`
				Type         *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"session_affinity" json:"sessionAffinity,omitempty"`
			TimeoutSec *int64 `tfsdk:"timeout_sec" json:"timeoutSec,omitempty"`
		} `tfsdk:"default" json:"default,omitempty"`
		TargetRef *struct {
			Group     *string `tfsdk:"group" json:"group,omitempty"`
			Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"target_ref" json:"targetRef,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *NetworkingGkeIoGcpbackendPolicyV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_networking_gke_io_gcp_backend_policy_v1_manifest"
}

func (r *NetworkingGkeIoGcpbackendPolicyV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "GCPBackendPolicy provides a way to apply LoadBalancer policy configuration with the GKE implementation of the Gateway API.",
		MarkdownDescription: "GCPBackendPolicy provides a way to apply LoadBalancer policy configuration with the GKE implementation of the Gateway API.",
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
				Description:         "Spec defines the desired state of GCPBackendPolicy.",
				MarkdownDescription: "Spec defines the desired state of GCPBackendPolicy.",
				Attributes: map[string]schema.Attribute{
					"default": schema.SingleNestedAttribute{
						Description:         "Default defines default policy configuration for the targeted resource.",
						MarkdownDescription: "Default defines default policy configuration for the targeted resource.",
						Attributes: map[string]schema.Attribute{
							"connection_draining": schema.SingleNestedAttribute{
								Description:         "ConnectionDraining contains configuration for connection draining",
								MarkdownDescription: "ConnectionDraining contains configuration for connection draining",
								Attributes: map[string]schema.Attribute{
									"draining_timeout_sec": schema.Int64Attribute{
										Description:         "DrainingTimeoutSec is a BackendService parameter. It is used during removal of VMs from instance groups. This guarantees that for the specified time all existing connections to a VM will remain untouched, but no new connections will be accepted. Set timeout to zero to disable connection draining. Enable the feature by specifying a timeout of up to one hour. If the field is omitted, a default value (0s) will be used. See https://cloud.google.com/compute/docs/reference/rest/v1/backendServices",
										MarkdownDescription: "DrainingTimeoutSec is a BackendService parameter. It is used during removal of VMs from instance groups. This guarantees that for the specified time all existing connections to a VM will remain untouched, but no new connections will be accepted. Set timeout to zero to disable connection draining. Enable the feature by specifying a timeout of up to one hour. If the field is omitted, a default value (0s) will be used. See https://cloud.google.com/compute/docs/reference/rest/v1/backendServices",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
											int64validator.AtMost(3600),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"iap": schema.SingleNestedAttribute{
								Description:         "IAP contains the configurations for Identity-Aware Proxy. Identity-Aware Proxy manages access control policies for backend services associated with a HTTPRoute, so they can be accessed only by authenticated users or applications with correct Identity and Access Management (IAM) role. See https://cloud.google.com/compute/docs/reference/rest/v1/backendServices",
								MarkdownDescription: "IAP contains the configurations for Identity-Aware Proxy. Identity-Aware Proxy manages access control policies for backend services associated with a HTTPRoute, so they can be accessed only by authenticated users or applications with correct Identity and Access Management (IAM) role. See https://cloud.google.com/compute/docs/reference/rest/v1/backendServices",
								Attributes: map[string]schema.Attribute{
									"client_id": schema.StringAttribute{
										Description:         "ClientID is the OAuth2 client ID to use for the authentication flow. See iap.oauth2ClientId in https://cloud.google.com/compute/docs/reference/rest/v1/backendServices ClientID must be set if Enabled is set to true.",
										MarkdownDescription: "ClientID is the OAuth2 client ID to use for the authentication flow. See iap.oauth2ClientId in https://cloud.google.com/compute/docs/reference/rest/v1/backendServices ClientID must be set if Enabled is set to true.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Enabled denotes whether the serving infrastructure will authenticate and authorize all incoming requests. If true, the ClientID and Oauth2ClientSecret fields must be non-empty. If not specified, this defaults to false, which means Identity-Aware Proxy is disabled by default.",
										MarkdownDescription: "Enabled denotes whether the serving infrastructure will authenticate and authorize all incoming requests. If true, the ClientID and Oauth2ClientSecret fields must be non-empty. If not specified, this defaults to false, which means Identity-Aware Proxy is disabled by default.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"oauth2_client_secret": schema.SingleNestedAttribute{
										Description:         "Oauth2ClientSecret contains the OAuth2 client secret to use for the authentication flow. See https://cloud.google.com/compute/docs/reference/rest/v1/backendServices Oauth2ClientSecret must be set if Enabled is set to true.",
										MarkdownDescription: "Oauth2ClientSecret contains the OAuth2 client secret to use for the authentication flow. See https://cloud.google.com/compute/docs/reference/rest/v1/backendServices Oauth2ClientSecret must be set if Enabled is set to true.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name is the reference to the secret resource.",
												MarkdownDescription: "Name is the reference to the secret resource.",
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

							"logging": schema.SingleNestedAttribute{
								Description:         "LoggingConfig contains configuration for logging.",
								MarkdownDescription: "LoggingConfig contains configuration for logging.",
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Description:         "Enabled denotes whether to enable logging for the load balancer traffic served by this backend service. If not specified, this defaults to false, which means logging is disabled by default.",
										MarkdownDescription: "Enabled denotes whether to enable logging for the load balancer traffic served by this backend service. If not specified, this defaults to false, which means logging is disabled by default.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"sample_rate": schema.Int64Attribute{
										Description:         "This field can only be specified if logging is enabled for this backend service. The value of the field must be in range [0, 1e6]. This is converted to a floating point value in the range [0, 1] by dividing by 1e6 for use with the GCE api and interpreted as the proportion of requests that will be logged. By default all requests will be logged.",
										MarkdownDescription: "This field can only be specified if logging is enabled for this backend service. The value of the field must be in range [0, 1e6]. This is converted to a floating point value in the range [0, 1] by dividing by 1e6 for use with the GCE api and interpreted as the proportion of requests that will be logged. By default all requests will be logged.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
											int64validator.AtMost(1e+06),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"security_policy": schema.StringAttribute{
								Description:         "SecurityPolicy is a reference to a GCP Cloud Armor SecurityPolicy resource.",
								MarkdownDescription: "SecurityPolicy is a reference to a GCP Cloud Armor SecurityPolicy resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"session_affinity": schema.SingleNestedAttribute{
								Description:         "SessionAffinityConfig contains configuration for stickiness parameters.",
								MarkdownDescription: "SessionAffinityConfig contains configuration for stickiness parameters.",
								Attributes: map[string]schema.Attribute{
									"cookie_ttl_sec": schema.Int64Attribute{
										Description:         "CookieTTLSec specifies the lifetime of cookies in seconds. This setting requires GENERATED_COOKIE or HTTP_COOKIE session affinity. If set to 0, the cookie is non-persistent and lasts only until the end of the browser session (or equivalent). The maximum allowed value is two weeks (1,209,600).",
										MarkdownDescription: "CookieTTLSec specifies the lifetime of cookies in seconds. This setting requires GENERATED_COOKIE or HTTP_COOKIE session affinity. If set to 0, the cookie is non-persistent and lasts only until the end of the browser session (or equivalent). The maximum allowed value is two weeks (1,209,600).",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
											int64validator.AtMost(1.2096e+06),
										},
									},

									"type": schema.StringAttribute{
										Description:         "Type specifies the type of session affinity to use. If not specified, this defaults to NONE.",
										MarkdownDescription: "Type specifies the type of session affinity to use. If not specified, this defaults to NONE.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("CLIENT_IP", "CLIENT_IP_PORT_PROTO", "CLIENT_IP_PROTO", "GENERATED_COOKIE", "HEADER_FIELD", "HTTP_COOKIE", "NONE"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"timeout_sec": schema.Int64Attribute{
								Description:         "TimeoutSec is a BackendService parameter. See https://cloud.google.com/compute/docs/reference/rest/v1/backendServices. If the field is omitted, a default value (30s) will be used.",
								MarkdownDescription: "TimeoutSec is a BackendService parameter. See https://cloud.google.com/compute/docs/reference/rest/v1/backendServices. If the field is omitted, a default value (30s) will be used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(2.147483647e+09),
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

func (r *NetworkingGkeIoGcpbackendPolicyV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_networking_gke_io_gcp_backend_policy_v1_manifest")

	var model NetworkingGkeIoGcpbackendPolicyV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("networking.gke.io/v1")
	model.Kind = pointer.String("GCPBackendPolicy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
