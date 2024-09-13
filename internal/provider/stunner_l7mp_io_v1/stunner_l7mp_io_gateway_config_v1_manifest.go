/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package stunner_l7mp_io_v1

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
	_ datasource.DataSource = &StunnerL7MpIoGatewayConfigV1Manifest{}
)

func NewStunnerL7MpIoGatewayConfigV1Manifest() datasource.DataSource {
	return &StunnerL7MpIoGatewayConfigV1Manifest{}
}

type StunnerL7MpIoGatewayConfigV1Manifest struct{}

type StunnerL7MpIoGatewayConfigV1ManifestData struct {
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
		AuthLifetime *int64 `tfsdk:"auth_lifetime" json:"authLifetime,omitempty"`
		AuthRef      *struct {
			Group     *string `tfsdk:"group" json:"group,omitempty"`
			Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"auth_ref" json:"authRef,omitempty"`
		AuthType                       *string            `tfsdk:"auth_type" json:"authType,omitempty"`
		Dataplane                      *string            `tfsdk:"dataplane" json:"dataplane,omitempty"`
		LoadBalancerServiceAnnotations *map[string]string `tfsdk:"load_balancer_service_annotations" json:"loadBalancerServiceAnnotations,omitempty"`
		LogLevel                       *string            `tfsdk:"log_level" json:"logLevel,omitempty"`
		Password                       *string            `tfsdk:"password" json:"password,omitempty"`
		Realm                          *string            `tfsdk:"realm" json:"realm,omitempty"`
		SharedSecret                   *string            `tfsdk:"shared_secret" json:"sharedSecret,omitempty"`
		UserName                       *string            `tfsdk:"user_name" json:"userName,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *StunnerL7MpIoGatewayConfigV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_stunner_l7mp_io_gateway_config_v1_manifest"
}

func (r *StunnerL7MpIoGatewayConfigV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "GatewayConfig is the Schema for the gatewayconfigs API",
		MarkdownDescription: "GatewayConfig is the Schema for the gatewayconfigs API",
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
				Description:         "GatewayConfigSpec defines the desired state of GatewayConfig",
				MarkdownDescription: "GatewayConfigSpec defines the desired state of GatewayConfig",
				Attributes: map[string]schema.Attribute{
					"auth_lifetime": schema.Int64Attribute{
						Description:         "AuthLifetime defines the lifetime of 'longterm' authentication credentials in seconds.",
						MarkdownDescription: "AuthLifetime defines the lifetime of 'longterm' authentication credentials in seconds.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"auth_ref": schema.SingleNestedAttribute{
						Description:         "Note that externally set credentials override any inline auth credentials (AuthType, AuthUsername, etc.): if AuthRef is nonempty then it is expected that the referenced Secret exists and *all* authentication credentials are correctly set in the referenced Secret (username/password or shared secret). Mixing of credential sources (inline/external) is not supported.",
						MarkdownDescription: "Note that externally set credentials override any inline auth credentials (AuthType, AuthUsername, etc.): if AuthRef is nonempty then it is expected that the referenced Secret exists and *all* authentication credentials are correctly set in the referenced Secret (username/password or shared secret). Mixing of credential sources (inline/external) is not supported.",
						Attributes: map[string]schema.Attribute{
							"group": schema.StringAttribute{
								Description:         "Group is the group of the referent. For example, 'gateway.networking.k8s.io'. When unspecified or empty string, core API group is inferred.",
								MarkdownDescription: "Group is the group of the referent. For example, 'gateway.networking.k8s.io'. When unspecified or empty string, core API group is inferred.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtMost(253),
									stringvalidator.RegexMatches(regexp.MustCompile(`^$|^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
								},
							},

							"kind": schema.StringAttribute{
								Description:         "Kind is kind of the referent. For example 'Secret'.",
								MarkdownDescription: "Kind is kind of the referent. For example 'Secret'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(63),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z]([-a-zA-Z0-9]*[a-zA-Z0-9])?$`), ""),
								},
							},

							"name": schema.StringAttribute{
								Description:         "Name is the name of the referent.",
								MarkdownDescription: "Name is the name of the referent.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(253),
								},
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace is the namespace of the referenced object. When unspecified, the local namespace is inferred. Note that when a namespace different than the local namespace is specified, a ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details. Support: Core",
								MarkdownDescription: "Namespace is the namespace of the referenced object. When unspecified, the local namespace is inferred. Note that when a namespace different than the local namespace is specified, a ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details. Support: Core",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"auth_type": schema.StringAttribute{
						Description:         "AuthType is the type of the STUN/TURN authentication mechanism.",
						MarkdownDescription: "AuthType is the type of the STUN/TURN authentication mechanism.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^plaintext|static|longterm|ephemeral|timewindowed$`), ""),
						},
					},

					"dataplane": schema.StringAttribute{
						Description:         "Dataplane defines the dataplane (stunnerd image, version, etc) for STUNner gateways using this GatewayConfig.",
						MarkdownDescription: "Dataplane defines the dataplane (stunnerd image, version, etc) for STUNner gateways using this GatewayConfig.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"load_balancer_service_annotations": schema.MapAttribute{
						Description:         "LoadBalancerServiceAnnotations is a list of annotations that will go into the LoadBalancer services created automatically by the operator to wrap Gateways. NOTE: removing annotations from a GatewayConfig will not result in the removal of the corresponding annotations from the LoadBalancer service, in order to prevent the accidental removal of an annotation installed there by Kubernetes or the cloud provider. If you really want to remove an annotation, do this manually or simply remove all Gateways (which will remove the corresponding LoadBalancer services), update the GatewayConfig and then recreate the Gateways, so that the newly created LoadBalancer services will contain the required annotations.",
						MarkdownDescription: "LoadBalancerServiceAnnotations is a list of annotations that will go into the LoadBalancer services created automatically by the operator to wrap Gateways. NOTE: removing annotations from a GatewayConfig will not result in the removal of the corresponding annotations from the LoadBalancer service, in order to prevent the accidental removal of an annotation installed there by Kubernetes or the cloud provider. If you really want to remove an annotation, do this manually or simply remove all Gateways (which will remove the corresponding LoadBalancer services), update the GatewayConfig and then recreate the Gateways, so that the newly created LoadBalancer services will contain the required annotations.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"log_level": schema.StringAttribute{
						Description:         "LogLevel specifies the default loglevel for the STUNner daemon.",
						MarkdownDescription: "LogLevel specifies the default loglevel for the STUNner daemon.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"password": schema.StringAttribute{
						Description:         "Password defines the 'password' credential for 'plaintext' authentication.",
						MarkdownDescription: "Password defines the 'password' credential for 'plaintext' authentication.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
						},
					},

					"realm": schema.StringAttribute{
						Description:         "Realm defines the STUN/TURN authentication realm to be used for clients toauthenticate with STUNner. The realm must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character. No other punctuation is allowed.",
						MarkdownDescription: "Realm defines the STUN/TURN authentication realm to be used for clients toauthenticate with STUNner. The realm must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character. No other punctuation is allowed.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
						},
					},

					"shared_secret": schema.StringAttribute{
						Description:         "SharedSecret defines the shared secret to be used for 'longterm' authentication.",
						MarkdownDescription: "SharedSecret defines the shared secret to be used for 'longterm' authentication.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"user_name": schema.StringAttribute{
						Description:         "Username defines the 'username' credential for 'plaintext' authentication.",
						MarkdownDescription: "Username defines the 'username' credential for 'plaintext' authentication.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
						},
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *StunnerL7MpIoGatewayConfigV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_stunner_l7mp_io_gateway_config_v1_manifest")

	var model StunnerL7MpIoGatewayConfigV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("stunner.l7mp.io/v1")
	model.Kind = pointer.String("GatewayConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
