/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package security_profiles_operator_x_k8s_io_v1alpha1

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
	_ datasource.DataSource = &SecurityProfilesOperatorXK8SIoAppArmorProfileV1Alpha1Manifest{}
)

func NewSecurityProfilesOperatorXK8SIoAppArmorProfileV1Alpha1Manifest() datasource.DataSource {
	return &SecurityProfilesOperatorXK8SIoAppArmorProfileV1Alpha1Manifest{}
}

type SecurityProfilesOperatorXK8SIoAppArmorProfileV1Alpha1Manifest struct{}

type SecurityProfilesOperatorXK8SIoAppArmorProfileV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Abstract *struct {
			Capability *struct {
				AllowedCapabilities *[]string `tfsdk:"allowed_capabilities" json:"allowedCapabilities,omitempty"`
			} `tfsdk:"capability" json:"capability,omitempty"`
			Executable *struct {
				AllowedExecutables *[]string `tfsdk:"allowed_executables" json:"allowedExecutables,omitempty"`
				AllowedLibraries   *[]string `tfsdk:"allowed_libraries" json:"allowedLibraries,omitempty"`
			} `tfsdk:"executable" json:"executable,omitempty"`
			Filesystem *struct {
				ReadOnlyPaths  *[]string `tfsdk:"read_only_paths" json:"readOnlyPaths,omitempty"`
				ReadWritePaths *[]string `tfsdk:"read_write_paths" json:"readWritePaths,omitempty"`
				WriteOnlyPaths *[]string `tfsdk:"write_only_paths" json:"writeOnlyPaths,omitempty"`
			} `tfsdk:"filesystem" json:"filesystem,omitempty"`
			Network *struct {
				AllowRaw         *bool `tfsdk:"allow_raw" json:"allowRaw,omitempty"`
				AllowedProtocols *struct {
					AllowTcp *bool `tfsdk:"allow_tcp" json:"allowTcp,omitempty"`
					AllowUdp *bool `tfsdk:"allow_udp" json:"allowUdp,omitempty"`
				} `tfsdk:"allowed_protocols" json:"allowedProtocols,omitempty"`
			} `tfsdk:"network" json:"network,omitempty"`
		} `tfsdk:"abstract" json:"abstract,omitempty"`
		ComplainMode *bool `tfsdk:"complain_mode" json:"complainMode,omitempty"`
		Disabled     *bool `tfsdk:"disabled" json:"disabled,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SecurityProfilesOperatorXK8SIoAppArmorProfileV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_security_profiles_operator_x_k8s_io_app_armor_profile_v1alpha1_manifest"
}

func (r *SecurityProfilesOperatorXK8SIoAppArmorProfileV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "AppArmorProfile is a cluster level specification for an AppArmor profile.",
		MarkdownDescription: "AppArmorProfile is a cluster level specification for an AppArmor profile.",
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
				Description:         "AppArmorProfileSpec defines the desired state of AppArmorProfile.",
				MarkdownDescription: "AppArmorProfileSpec defines the desired state of AppArmorProfile.",
				Attributes: map[string]schema.Attribute{
					"abstract": schema.SingleNestedAttribute{
						Description:         "Abstract stores the apparmor profile allow lists for executable, file, network and capabilities access.",
						MarkdownDescription: "Abstract stores the apparmor profile allow lists for executable, file, network and capabilities access.",
						Attributes: map[string]schema.Attribute{
							"capability": schema.SingleNestedAttribute{
								Description:         "Capability rules for Linux capabilities.",
								MarkdownDescription: "Capability rules for Linux capabilities.",
								Attributes: map[string]schema.Attribute{
									"allowed_capabilities": schema.ListAttribute{
										Description:         "AllowedCapabilities lost of allowed capabilities.",
										MarkdownDescription: "AllowedCapabilities lost of allowed capabilities.",
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

							"executable": schema.SingleNestedAttribute{
								Description:         "Executable rules for allowed executables.",
								MarkdownDescription: "Executable rules for allowed executables.",
								Attributes: map[string]schema.Attribute{
									"allowed_executables": schema.ListAttribute{
										Description:         "AllowedExecutables list of allowed executables.",
										MarkdownDescription: "AllowedExecutables list of allowed executables.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"allowed_libraries": schema.ListAttribute{
										Description:         "AllowedLibraries list of allowed libraries.",
										MarkdownDescription: "AllowedLibraries list of allowed libraries.",
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

							"filesystem": schema.SingleNestedAttribute{
								Description:         "Filesystem rules for filesystem access.",
								MarkdownDescription: "Filesystem rules for filesystem access.",
								Attributes: map[string]schema.Attribute{
									"read_only_paths": schema.ListAttribute{
										Description:         "ReadOnlyPaths list of allowed read only file paths.",
										MarkdownDescription: "ReadOnlyPaths list of allowed read only file paths.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"read_write_paths": schema.ListAttribute{
										Description:         "ReadWritePaths list of allowed read write file paths.",
										MarkdownDescription: "ReadWritePaths list of allowed read write file paths.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"write_only_paths": schema.ListAttribute{
										Description:         "WriteOnlyPaths list of allowed write only file paths.",
										MarkdownDescription: "WriteOnlyPaths list of allowed write only file paths.",
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

							"network": schema.SingleNestedAttribute{
								Description:         "Network rules for network access.",
								MarkdownDescription: "Network rules for network access.",
								Attributes: map[string]schema.Attribute{
									"allow_raw": schema.BoolAttribute{
										Description:         "AllowRaw allows raw sockets.",
										MarkdownDescription: "AllowRaw allows raw sockets.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"allowed_protocols": schema.SingleNestedAttribute{
										Description:         "Protocols keeps the allowed networking protocols.",
										MarkdownDescription: "Protocols keeps the allowed networking protocols.",
										Attributes: map[string]schema.Attribute{
											"allow_tcp": schema.BoolAttribute{
												Description:         "AllowTCP allows TCP socket connections.",
												MarkdownDescription: "AllowTCP allows TCP socket connections.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"allow_udp": schema.BoolAttribute{
												Description:         "AllowUDP allows UDP sockets connections.",
												MarkdownDescription: "AllowUDP allows UDP sockets connections.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"complain_mode": schema.BoolAttribute{
						Description:         "ComplainMode places the apparmor profile into 'complain' mode, by default is placed in 'enforce' mode. In complain mode, if a given action is not allowed, it will be allowed, but this violation will be logged with a tag of access being 'ALLOWED unconfined'.",
						MarkdownDescription: "ComplainMode places the apparmor profile into 'complain' mode, by default is placed in 'enforce' mode. In complain mode, if a given action is not allowed, it will be allowed, but this violation will be logged with a tag of access being 'ALLOWED unconfined'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"disabled": schema.BoolAttribute{
						Description:         "Whether the profile is disabled and should be skipped during reconciliation.",
						MarkdownDescription: "Whether the profile is disabled and should be skipped during reconciliation.",
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

func (r *SecurityProfilesOperatorXK8SIoAppArmorProfileV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_security_profiles_operator_x_k8s_io_app_armor_profile_v1alpha1_manifest")

	var model SecurityProfilesOperatorXK8SIoAppArmorProfileV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("security-profiles-operator.x-k8s.io/v1alpha1")
	model.Kind = pointer.String("AppArmorProfile")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
