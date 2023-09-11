/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package security_profiles_operator_x_k8s_io_v1beta1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"net/http"
)

var (
	_ datasource.DataSource              = &SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1DataSource{}
	_ datasource.DataSourceWithConfigure = &SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1DataSource{}
)

func NewSecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1DataSource() datasource.DataSource {
	return &SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1DataSource{}
}

type SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1DataSource struct {
	kubernetesClient dynamic.Interface
}

type SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

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

func (r *SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_security_profiles_operator_x_k8s_io_seccomp_profile_v1beta1"
}

func (r *SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
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
						Optional:            false,
						Computed:            true,
					},

					"base_profile_name": schema.StringAttribute{
						Description:         "BaseProfileName is the name of base profile (in the same namespace) that will be unioned into this profile. Base profiles can be references as remote OCI artifacts as well when prefixed with 'oci://'.",
						MarkdownDescription: "BaseProfileName is the name of base profile (in the same namespace) that will be unioned into this profile. Base profiles can be references as remote OCI artifacts as well when prefixed with 'oci://'.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"default_action": schema.StringAttribute{
						Description:         "the default action for seccomp",
						MarkdownDescription: "the default action for seccomp",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"disabled": schema.BoolAttribute{
						Description:         "Whether the profile is disabled and should be skipped during reconciliation.",
						MarkdownDescription: "Whether the profile is disabled and should be skipped during reconciliation.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"flags": schema.ListAttribute{
						Description:         "list of flags to use with seccomp(2)",
						MarkdownDescription: "list of flags to use with seccomp(2)",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"listener_metadata": schema.StringAttribute{
						Description:         "opaque data to pass to the seccomp agent",
						MarkdownDescription: "opaque data to pass to the seccomp agent",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"listener_path": schema.StringAttribute{
						Description:         "path of UNIX domain socket to contact a seccomp agent for SCMP_ACT_NOTIFY",
						MarkdownDescription: "path of UNIX domain socket to contact a seccomp agent for SCMP_ACT_NOTIFY",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"syscalls": schema.ListNestedAttribute{
						Description:         "match a syscall in seccomp. While this property is OPTIONAL, some values of defaultAction are not useful without syscalls entries. For example, if defaultAction is SCMP_ACT_KILL and syscalls is empty or unset, the kernel will kill the container process on its first syscall",
						MarkdownDescription: "match a syscall in seccomp. While this property is OPTIONAL, some values of defaultAction are not useful without syscalls entries. For example, if defaultAction is SCMP_ACT_KILL and syscalls is empty or unset, the kernel will kill the container process on its first syscall",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action": schema.StringAttribute{
									Description:         "the action for seccomp rules",
									MarkdownDescription: "the action for seccomp rules",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"args": schema.ListNestedAttribute{
									Description:         "the specific syscall in seccomp",
									MarkdownDescription: "the specific syscall in seccomp",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"index": schema.Int64Attribute{
												Description:         "the index for syscall arguments in seccomp",
												MarkdownDescription: "the index for syscall arguments in seccomp",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"op": schema.StringAttribute{
												Description:         "the operator for syscall arguments in seccomp",
												MarkdownDescription: "the operator for syscall arguments in seccomp",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"value": schema.Int64Attribute{
												Description:         "the value for syscall arguments in seccomp",
												MarkdownDescription: "the value for syscall arguments in seccomp",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"value_two": schema.Int64Attribute{
												Description:         "the value for syscall arguments in seccomp",
												MarkdownDescription: "the value for syscall arguments in seccomp",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"errno_ret": schema.Int64Attribute{
									Description:         "the errno return code to use. Some actions like SCMP_ACT_ERRNO and SCMP_ACT_TRACE allow to specify the errno code to return",
									MarkdownDescription: "the errno return code to use. Some actions like SCMP_ACT_ERRNO and SCMP_ACT_TRACE allow to specify the errno code to return",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"names": schema.ListAttribute{
									Description:         "the names of the syscalls",
									MarkdownDescription: "the names of the syscalls",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_security_profiles_operator_x_k8s_io_seccomp_profile_v1beta1")

	var data SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "security-profiles-operator.x-k8s.io", Version: "v1beta1", Resource: "seccompprofiles"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		var statusError *k8sErrors.StatusError
		if errors.As(err, &statusError) {
			if statusError.Status().Code == http.StatusNotFound {
				response.Diagnostics.AddError(
					"Unable to find resource",
					fmt.Sprintf("The requested resource cannot be found. "+
						"Make sure that it does exist in your cluster and you have set the correct name and namespace configured.\n\n"+
						"Namespace: %s\n"+
						"Name: %s", data.Metadata.Namespace, data.Metadata.Name),
				)
				return
			}
		} else {
			response.Diagnostics.AddError(
				"Unable to GET resource",
				fmt.Sprintf("An unexpected error occurred while reading the resource. "+
					"Please report this issue to the provider developers.\n\n"+
					"GET Error (%T): %s", err, err.Error()),
			)
		}
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

	var readResponse SecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1DataSourceData
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

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("security-profiles-operator.x-k8s.io/v1beta1")
	data.Kind = pointer.String("SeccompProfile")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
