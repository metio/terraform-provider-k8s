/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package security_profiles_operator_x_k8s_io_v1beta1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	k8sTypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"strings"
)

var (
	_ resource.Resource                = &SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1Resource{}
	_ resource.ResourceWithConfigure   = &SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1Resource{}
	_ resource.ResourceWithImportState = &SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1Resource{}
)

func NewSecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1Resource() resource.Resource {
	return &SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1Resource{}
}

type SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitFor        types.List   `tfsdk:"wait_for" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

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

func (r *SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_security_profiles_operator_x_k8s_io_seccomp_profile_v1beta1"
}

func (r *SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
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

			"force_conflicts": schema.BoolAttribute{
				Description:         "If 'true', server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "If `true`, server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"field_manager": schema.BoolAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"wait_for": schema.ListNestedAttribute{
				Description:         "Wait for specific conditions after create/update of resources.",
				MarkdownDescription: "Wait for specific conditions after create/update of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"jsonpath": schema.StringAttribute{
							Description:         "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							MarkdownDescription: "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},
						"value": schema.StringAttribute{
							Description:         "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							MarkdownDescription: "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							Required:            false,
							Optional:            true,
							Computed:            true,
						},
						"timeout": schema.StringAttribute{
							Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("30s"),
						},
					},
				},
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            true,
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
						Computed:            true,
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

func (r *SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *dynamic.DynamicClient, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_security_profiles_operator_x_k8s_io_seccomp_profile_v1beta1")

	var model SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("security-profiles-operator.x-k8s.io/v1beta1")
	model.Kind = pointer.String("SeccompProfile")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "security-profiles-operator.x-k8s.io", Version: "v1beta1", Resource: "SeccompProfile"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while creating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_security_profiles_operator_x_k8s_io_seccomp_profile_v1beta1")

	var data SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "security-profiles-operator.x-k8s.io", Version: "v1beta1", Resource: "SeccompProfile"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_security_profiles_operator_x_k8s_io_seccomp_profile_v1beta1")

	var model SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("security-profiles-operator.x-k8s.io/v1beta1")
	model.Kind = pointer.String("SeccompProfile")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "security-profiles-operator.x-k8s.io", Version: "v1beta1", Resource: "SeccompProfile"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while updating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_security_profiles_operator_x_k8s_io_seccomp_profile_v1beta1")

	var data SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "security-profiles-operator.x-k8s.io", Version: "v1beta1", Resource: "SeccompProfile"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to DELETE resource",
			"An unexpected error occurred while deleting the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"DELETE Error: "+err.Error(),
		)
		return
	}
}

func (r *SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	idParts := strings.Split(request.ID, "/")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'namespace/name' Got: '%q'", request.ID),
		)
		return
	}

	namespace := idParts[0]
	name := idParts[1]
	tflog.Trace(ctx, "parsed import ID", map[string]interface{}{
		"namespace": namespace,
		"name":      name,
	})
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("namespace"), namespace)...)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("name"), name)...)
}
