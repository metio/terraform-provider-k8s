/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package security_profiles_operator_x_k8s_io_v1beta1

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	_ datasource.DataSource = &SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1Manifest{}
)

func NewSecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1Manifest() datasource.DataSource {
	return &SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1Manifest{}
}

type SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1Manifest struct{}

type SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1ManifestData struct {
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
		Architectures    *[]string `tfsdk:"architectures" json:"architectures,omitempty"`
		BaseProfileName  *string   `tfsdk:"base_profile_name" json:"baseProfileName,omitempty"`
		DefaultAction    *string   `tfsdk:"default_action" json:"defaultAction,omitempty"`
		Disabled         *bool     `tfsdk:"disabled" json:"disabled,omitempty"`
		Flags            *[]string `tfsdk:"flags" json:"flags,omitempty"`
		ListenerMetadata *string   `tfsdk:"listener_metadata" json:"listenerMetadata,omitempty"`
		ListenerPath     *string   `tfsdk:"listener_path" json:"listenerPath,omitempty"`
		Syscalls         *[]struct {
			Action *string `tfsdk:"action" json:"action,omitempty"`
			Args   *[]struct {
				Index    *int64  `tfsdk:"index" json:"index,omitempty"`
				Op       *string `tfsdk:"op" json:"op,omitempty"`
				Value    *int64  `tfsdk:"value" json:"value,omitempty"`
				ValueTwo *int64  `tfsdk:"value_two" json:"valueTwo,omitempty"`
			} `tfsdk:"args" json:"args,omitempty"`
			ErrnoRet *int64    `tfsdk:"errno_ret" json:"errnoRet,omitempty"`
			Names    *[]string `tfsdk:"names" json:"names,omitempty"`
		} `tfsdk:"syscalls" json:"syscalls,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_security_profiles_operator_x_k8s_io_seccomp_profile_v1beta1_manifest"
}

func (r *SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "SeccompProfile is a cluster level specification for a seccomp profile. See https://github.com/opencontainers/runtime-spec/blob/master/config-linux.md#seccomp",
		MarkdownDescription: "SeccompProfile is a cluster level specification for a seccomp profile. See https://github.com/opencontainers/runtime-spec/blob/master/config-linux.md#seccomp",
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
				Description:         "SeccompProfileSpec defines the desired state of SeccompProfile.",
				MarkdownDescription: "SeccompProfileSpec defines the desired state of SeccompProfile.",
				Attributes: map[string]schema.Attribute{
					"architectures": schema.ListAttribute{
						Description:         "the architecture used for system calls",
						MarkdownDescription: "the architecture used for system calls",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"base_profile_name": schema.StringAttribute{
						Description:         "BaseProfileName is the name of base profile (in the same namespace) that will be unioned into this profile. Base profiles can be references as remote OCI artifacts as well when prefixed with 'oci://'.",
						MarkdownDescription: "BaseProfileName is the name of base profile (in the same namespace) that will be unioned into this profile. Base profiles can be references as remote OCI artifacts as well when prefixed with 'oci://'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"default_action": schema.StringAttribute{
						Description:         "the default action for seccomp",
						MarkdownDescription: "the default action for seccomp",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("SCMP_ACT_KILL", "SCMP_ACT_KILL_PROCESS", "SCMP_ACT_KILL_THREAD", "SCMP_ACT_TRAP", "SCMP_ACT_ERRNO", "SCMP_ACT_TRACE", "SCMP_ACT_ALLOW", "SCMP_ACT_LOG", "SCMP_ACT_NOTIFY"),
						},
					},

					"disabled": schema.BoolAttribute{
						Description:         "Whether the profile is disabled and should be skipped during reconciliation.",
						MarkdownDescription: "Whether the profile is disabled and should be skipped during reconciliation.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"flags": schema.ListAttribute{
						Description:         "list of flags to use with seccomp(2)",
						MarkdownDescription: "list of flags to use with seccomp(2)",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"listener_metadata": schema.StringAttribute{
						Description:         "opaque data to pass to the seccomp agent",
						MarkdownDescription: "opaque data to pass to the seccomp agent",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"listener_path": schema.StringAttribute{
						Description:         "path of UNIX domain socket to contact a seccomp agent for SCMP_ACT_NOTIFY",
						MarkdownDescription: "path of UNIX domain socket to contact a seccomp agent for SCMP_ACT_NOTIFY",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"syscalls": schema.ListNestedAttribute{
						Description:         "match a syscall in seccomp. While this property is OPTIONAL, some values of defaultAction are not useful without syscalls entries. For example, if defaultAction is SCMP_ACT_KILL and syscalls is empty or unset, the kernel will kill the container process on its first syscall",
						MarkdownDescription: "match a syscall in seccomp. While this property is OPTIONAL, some values of defaultAction are not useful without syscalls entries. For example, if defaultAction is SCMP_ACT_KILL and syscalls is empty or unset, the kernel will kill the container process on its first syscall",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action": schema.StringAttribute{
									Description:         "the action for seccomp rules",
									MarkdownDescription: "the action for seccomp rules",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("SCMP_ACT_KILL", "SCMP_ACT_KILL_PROCESS", "SCMP_ACT_KILL_THREAD", "SCMP_ACT_TRAP", "SCMP_ACT_ERRNO", "SCMP_ACT_TRACE", "SCMP_ACT_ALLOW", "SCMP_ACT_LOG", "SCMP_ACT_NOTIFY"),
									},
								},

								"args": schema.ListNestedAttribute{
									Description:         "the specific syscall in seccomp",
									MarkdownDescription: "the specific syscall in seccomp",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"index": schema.Int64Attribute{
												Description:         "the index for syscall arguments in seccomp",
												MarkdownDescription: "the index for syscall arguments in seccomp",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"op": schema.StringAttribute{
												Description:         "the operator for syscall arguments in seccomp",
												MarkdownDescription: "the operator for syscall arguments in seccomp",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("SCMP_CMP_NE", "SCMP_CMP_LT", "SCMP_CMP_LE", "SCMP_CMP_EQ", "SCMP_CMP_GE", "SCMP_CMP_GT", "SCMP_CMP_MASKED_EQ"),
												},
											},

											"value": schema.Int64Attribute{
												Description:         "the value for syscall arguments in seccomp",
												MarkdownDescription: "the value for syscall arguments in seccomp",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"value_two": schema.Int64Attribute{
												Description:         "the value for syscall arguments in seccomp",
												MarkdownDescription: "the value for syscall arguments in seccomp",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"errno_ret": schema.Int64Attribute{
									Description:         "the errno return code to use. Some actions like SCMP_ACT_ERRNO and SCMP_ACT_TRACE allow to specify the errno code to return",
									MarkdownDescription: "the errno return code to use. Some actions like SCMP_ACT_ERRNO and SCMP_ACT_TRACE allow to specify the errno code to return",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"names": schema.ListAttribute{
									Description:         "the names of the syscalls",
									MarkdownDescription: "the names of the syscalls",
									ElementType:         types.StringType,
									Required:            true,
									Optional:            false,
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
		},
	}
}

func (r *SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_security_profiles_operator_x_k8s_io_seccomp_profile_v1beta1_manifest")

	var model SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("security-profiles-operator.x-k8s.io/v1beta1")
	model.Kind = pointer.String("SeccompProfile")

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
