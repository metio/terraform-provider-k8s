/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package gateway_networking_k8s_io_v1alpha2

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &GatewayNetworkingK8SIoBackendLbpolicyV1Alpha2Manifest{}
)

func NewGatewayNetworkingK8SIoBackendLbpolicyV1Alpha2Manifest() datasource.DataSource {
	return &GatewayNetworkingK8SIoBackendLbpolicyV1Alpha2Manifest{}
}

type GatewayNetworkingK8SIoBackendLbpolicyV1Alpha2Manifest struct{}

type GatewayNetworkingK8SIoBackendLbpolicyV1Alpha2ManifestData struct {
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
		SessionPersistence *struct {
			AbsoluteTimeout *string `tfsdk:"absolute_timeout" json:"absoluteTimeout,omitempty"`
			CookieConfig    *struct {
				LifetimeType *string `tfsdk:"lifetime_type" json:"lifetimeType,omitempty"`
			} `tfsdk:"cookie_config" json:"cookieConfig,omitempty"`
			IdleTimeout *string `tfsdk:"idle_timeout" json:"idleTimeout,omitempty"`
			SessionName *string `tfsdk:"session_name" json:"sessionName,omitempty"`
			Type        *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"session_persistence" json:"sessionPersistence,omitempty"`
		TargetRefs *[]struct {
			Group *string `tfsdk:"group" json:"group,omitempty"`
			Kind  *string `tfsdk:"kind" json:"kind,omitempty"`
			Name  *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"target_refs" json:"targetRefs,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GatewayNetworkingK8SIoBackendLbpolicyV1Alpha2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_gateway_networking_k8s_io_backend_lb_policy_v1alpha2_manifest"
}

func (r *GatewayNetworkingK8SIoBackendLbpolicyV1Alpha2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "BackendLBPolicy provides a way to define load balancing rules for a backend.",
		MarkdownDescription: "BackendLBPolicy provides a way to define load balancing rules for a backend.",
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
				Description:         "Spec defines the desired state of BackendLBPolicy.",
				MarkdownDescription: "Spec defines the desired state of BackendLBPolicy.",
				Attributes: map[string]schema.Attribute{
					"session_persistence": schema.SingleNestedAttribute{
						Description:         "SessionPersistence defines and configures session persistence for the backend. Support: Extended",
						MarkdownDescription: "SessionPersistence defines and configures session persistence for the backend. Support: Extended",
						Attributes: map[string]schema.Attribute{
							"absolute_timeout": schema.StringAttribute{
								Description:         "AbsoluteTimeout defines the absolute timeout of the persistent session. Once the AbsoluteTimeout duration has elapsed, the session becomes invalid. Support: Extended",
								MarkdownDescription: "AbsoluteTimeout defines the absolute timeout of the persistent session. Once the AbsoluteTimeout duration has elapsed, the session becomes invalid. Support: Extended",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]{1,5}(h|m|s|ms)){1,4}$`), ""),
								},
							},

							"cookie_config": schema.SingleNestedAttribute{
								Description:         "CookieConfig provides configuration settings that are specific to cookie-based session persistence. Support: Core",
								MarkdownDescription: "CookieConfig provides configuration settings that are specific to cookie-based session persistence. Support: Core",
								Attributes: map[string]schema.Attribute{
									"lifetime_type": schema.StringAttribute{
										Description:         "LifetimeType specifies whether the cookie has a permanent or session-based lifetime. A permanent cookie persists until its specified expiry time, defined by the Expires or Max-Age cookie attributes, while a session cookie is deleted when the current session ends. When set to 'Permanent', AbsoluteTimeout indicates the cookie's lifetime via the Expires or Max-Age cookie attributes and is required. When set to 'Session', AbsoluteTimeout indicates the absolute lifetime of the cookie tracked by the gateway and is optional. Support: Core for 'Session' type Support: Extended for 'Permanent' type",
										MarkdownDescription: "LifetimeType specifies whether the cookie has a permanent or session-based lifetime. A permanent cookie persists until its specified expiry time, defined by the Expires or Max-Age cookie attributes, while a session cookie is deleted when the current session ends. When set to 'Permanent', AbsoluteTimeout indicates the cookie's lifetime via the Expires or Max-Age cookie attributes and is required. When set to 'Session', AbsoluteTimeout indicates the absolute lifetime of the cookie tracked by the gateway and is optional. Support: Core for 'Session' type Support: Extended for 'Permanent' type",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Permanent", "Session"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"idle_timeout": schema.StringAttribute{
								Description:         "IdleTimeout defines the idle timeout of the persistent session. Once the session has been idle for more than the specified IdleTimeout duration, the session becomes invalid. Support: Extended",
								MarkdownDescription: "IdleTimeout defines the idle timeout of the persistent session. Once the session has been idle for more than the specified IdleTimeout duration, the session becomes invalid. Support: Extended",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]{1,5}(h|m|s|ms)){1,4}$`), ""),
								},
							},

							"session_name": schema.StringAttribute{
								Description:         "SessionName defines the name of the persistent session token which may be reflected in the cookie or the header. Users should avoid reusing session names to prevent unintended consequences, such as rejection or unpredictable behavior. Support: Implementation-specific",
								MarkdownDescription: "SessionName defines the name of the persistent session token which may be reflected in the cookie or the header. Users should avoid reusing session names to prevent unintended consequences, such as rejection or unpredictable behavior. Support: Implementation-specific",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtMost(128),
								},
							},

							"type": schema.StringAttribute{
								Description:         "Type defines the type of session persistence such as through the use a header or cookie. Defaults to cookie based session persistence. Support: Core for 'Cookie' type Support: Extended for 'Header' type",
								MarkdownDescription: "Type defines the type of session persistence such as through the use a header or cookie. Defaults to cookie based session persistence. Support: Core for 'Cookie' type Support: Extended for 'Header' type",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Cookie", "Header"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"target_refs": schema.ListNestedAttribute{
						Description:         "TargetRef identifies an API object to apply policy to. Currently, Backends (i.e. Service, ServiceImport, or any implementation-specific backendRef) are the only valid API target references.",
						MarkdownDescription: "TargetRef identifies an API object to apply policy to. Currently, Backends (i.e. Service, ServiceImport, or any implementation-specific backendRef) are the only valid API target references.",
						NestedObject: schema.NestedAttributeObject{
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

func (r *GatewayNetworkingK8SIoBackendLbpolicyV1Alpha2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_gateway_networking_k8s_io_backend_lb_policy_v1alpha2_manifest")

	var model GatewayNetworkingK8SIoBackendLbpolicyV1Alpha2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("gateway.networking.k8s.io/v1alpha2")
	model.Kind = pointer.String("BackendLBPolicy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
