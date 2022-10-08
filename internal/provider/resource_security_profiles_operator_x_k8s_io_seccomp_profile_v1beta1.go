/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"

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

type SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1Resource struct{}

var (
	_ resource.Resource = (*SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1Resource)(nil)
)

type SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1GoModel struct {
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
		Architectures *[]string `tfsdk:"architectures" yaml:"architectures,omitempty"`

		BaseProfileName *string `tfsdk:"base_profile_name" yaml:"baseProfileName,omitempty"`

		DefaultAction *string `tfsdk:"default_action" yaml:"defaultAction,omitempty"`

		Flags *[]string `tfsdk:"flags" yaml:"flags,omitempty"`

		ListenerMetadata *string `tfsdk:"listener_metadata" yaml:"listenerMetadata,omitempty"`

		ListenerPath *string `tfsdk:"listener_path" yaml:"listenerPath,omitempty"`

		Syscalls *[]struct {
			Action *string `tfsdk:"action" yaml:"action,omitempty"`

			Args *[]struct {
				Index *int64 `tfsdk:"index" yaml:"index,omitempty"`

				Op *string `tfsdk:"op" yaml:"op,omitempty"`

				Value *int64 `tfsdk:"value" yaml:"value,omitempty"`

				ValueTwo *int64 `tfsdk:"value_two" yaml:"valueTwo,omitempty"`
			} `tfsdk:"args" yaml:"args,omitempty"`

			ErrnoRet *string `tfsdk:"errno_ret" yaml:"errnoRet,omitempty"`

			Names *[]string `tfsdk:"names" yaml:"names,omitempty"`
		} `tfsdk:"syscalls" yaml:"syscalls,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewSecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1Resource() resource.Resource {
	return &SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1Resource{}
}

func (r *SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_profiles_operator_x_k8s_io_seccomp_profile_v1beta1"
}

func (r *SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "SeccompProfile is a cluster level specification for a seccomp profile. See https://github.com/opencontainers/runtime-spec/blob/master/config-linux.md#seccomp",
		MarkdownDescription: "SeccompProfile is a cluster level specification for a seccomp profile. See https://github.com/opencontainers/runtime-spec/blob/master/config-linux.md#seccomp",
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
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.RequiresReplace(),
						},
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
				Description:         "SeccompProfileSpec defines the desired state of SeccompProfile.",
				MarkdownDescription: "SeccompProfileSpec defines the desired state of SeccompProfile.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"architectures": {
						Description:         "the architecture used for system calls",
						MarkdownDescription: "the architecture used for system calls",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"base_profile_name": {
						Description:         "name of base profile (in the same namespace) what will be unioned into this profile",
						MarkdownDescription: "name of base profile (in the same namespace) what will be unioned into this profile",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"default_action": {
						Description:         "the default action for seccomp",
						MarkdownDescription: "the default action for seccomp",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"flags": {
						Description:         "list of flags to use with seccomp(2)",
						MarkdownDescription: "list of flags to use with seccomp(2)",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"listener_metadata": {
						Description:         "opaque data to pass to the seccomp agent",
						MarkdownDescription: "opaque data to pass to the seccomp agent",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"listener_path": {
						Description:         "path of UNIX domain socket to contact a seccomp agent for SCMP_ACT_NOTIFY",
						MarkdownDescription: "path of UNIX domain socket to contact a seccomp agent for SCMP_ACT_NOTIFY",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"syscalls": {
						Description:         "match a syscall in seccomp. While this property is OPTIONAL, some values of defaultAction are not useful without syscalls entries. For example, if defaultAction is SCMP_ACT_KILL and syscalls is empty or unset, the kernel will kill the container process on its first syscall",
						MarkdownDescription: "match a syscall in seccomp. While this property is OPTIONAL, some values of defaultAction are not useful without syscalls entries. For example, if defaultAction is SCMP_ACT_KILL and syscalls is empty or unset, the kernel will kill the container process on its first syscall",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"action": {
								Description:         "the action for seccomp rules",
								MarkdownDescription: "the action for seccomp rules",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"args": {
								Description:         "the specific syscall in seccomp",
								MarkdownDescription: "the specific syscall in seccomp",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"index": {
										Description:         "the index for syscall arguments in seccomp",
										MarkdownDescription: "the index for syscall arguments in seccomp",

										Type: types.Int64Type,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(0),
										},
									},

									"op": {
										Description:         "the operator for syscall arguments in seccomp",
										MarkdownDescription: "the operator for syscall arguments in seccomp",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"value": {
										Description:         "the value for syscall arguments in seccomp",
										MarkdownDescription: "the value for syscall arguments in seccomp",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(0),
										},
									},

									"value_two": {
										Description:         "the value for syscall arguments in seccomp",
										MarkdownDescription: "the value for syscall arguments in seccomp",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(0),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"errno_ret": {
								Description:         "the errno return code to use. Some actions like SCMP_ACT_ERRNO and SCMP_ACT_TRACE allow to specify the errno code to return",
								MarkdownDescription: "the errno return code to use. Some actions like SCMP_ACT_ERRNO and SCMP_ACT_TRACE allow to specify the errno code to return",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"names": {
								Description:         "the names of the syscalls",
								MarkdownDescription: "the names of the syscalls",

								Type: types.ListType{ElemType: types.StringType},

								Required: true,
								Optional: false,
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
		},
	}, nil
}

func (r *SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_security_profiles_operator_x_k8s_io_seccomp_profile_v1beta1")

	var state SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("security-profiles-operator.x-k8s.io/v1beta1")
	goModel.Kind = utilities.Ptr("SeccompProfile")

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

func (r *SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_security_profiles_operator_x_k8s_io_seccomp_profile_v1beta1")
	// NO-OP: All data is already in Terraform state
}

func (r *SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_security_profiles_operator_x_k8s_io_seccomp_profile_v1beta1")

	var state SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("security-profiles-operator.x-k8s.io/v1beta1")
	goModel.Kind = utilities.Ptr("SeccompProfile")

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

func (r *SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_security_profiles_operator_x_k8s_io_seccomp_profile_v1beta1")
	// NO-OP: Terraform removes the state automatically for us
}
