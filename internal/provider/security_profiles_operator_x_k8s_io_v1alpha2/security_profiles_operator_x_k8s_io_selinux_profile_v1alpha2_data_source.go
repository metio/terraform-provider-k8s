/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package security_profiles_operator_x_k8s_io_v1alpha2

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &SecurityProfilesOperatorXK8SIoSelinuxProfileV1Alpha2DataSource{}
	_ datasource.DataSourceWithConfigure = &SecurityProfilesOperatorXK8SIoSelinuxProfileV1Alpha2DataSource{}
)

func NewSecurityProfilesOperatorXK8SIoSelinuxProfileV1Alpha2DataSource() datasource.DataSource {
	return &SecurityProfilesOperatorXK8SIoSelinuxProfileV1Alpha2DataSource{}
}

type SecurityProfilesOperatorXK8SIoSelinuxProfileV1Alpha2DataSource struct {
	kubernetesClient dynamic.Interface
}

type SecurityProfilesOperatorXK8SIoSelinuxProfileV1Alpha2DataSourceData struct {
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
		Allow    *map[string]map[string][]string `tfsdk:"allow" json:"allow,omitempty"`
		Disabled *bool                           `tfsdk:"disabled" json:"disabled,omitempty"`
		Inherit  *[]struct {
			Kind *string `tfsdk:"kind" json:"kind,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"inherit" json:"inherit,omitempty"`
		Permissive *bool `tfsdk:"permissive" json:"permissive,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SecurityProfilesOperatorXK8SIoSelinuxProfileV1Alpha2DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_security_profiles_operator_x_k8s_io_selinux_profile_v1alpha2"
}

func (r *SecurityProfilesOperatorXK8SIoSelinuxProfileV1Alpha2DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "SelinuxProfile is the Schema for the selinuxprofiles API.",
		MarkdownDescription: "SelinuxProfile is the Schema for the selinuxprofiles API.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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
				Description:         "SelinuxProfileSpec defines the desired state of SelinuxProfile.",
				MarkdownDescription: "SelinuxProfileSpec defines the desired state of SelinuxProfile.",
				Attributes: map[string]schema.Attribute{
					"allow": schema.MapAttribute{
						Description:         "Defines the allow policy for the profile",
						MarkdownDescription: "Defines the allow policy for the profile",
						ElementType:         types.MapType{ElemType: types.ListType{ElemType: types.StringType}},
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

					"inherit": schema.ListNestedAttribute{
						Description:         "A SELinuxProfile or set of profiles that this inherits from. Note that they need to be in the same namespace.",
						MarkdownDescription: "A SELinuxProfile or set of profiles that this inherits from. Note that they need to be in the same namespace.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"kind": schema.StringAttribute{
									Description:         "The Kind of the policy that this inherits from. Can be a SelinuxProfile object Or 'System' if an already installed policy will be used. The allowed 'System' policies are available in the SecurityProfilesOperatorDaemon instance.",
									MarkdownDescription: "The Kind of the policy that this inherits from. Can be a SelinuxProfile object Or 'System' if an already installed policy will be used. The allowed 'System' policies are available in the SecurityProfilesOperatorDaemon instance.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"name": schema.StringAttribute{
									Description:         "The name of the policy that this inherits from.",
									MarkdownDescription: "The name of the policy that this inherits from.",
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

					"permissive": schema.BoolAttribute{
						Description:         "Permissive, when true will cause the SELinux profile to only log violations instead of enforcing them.",
						MarkdownDescription: "Permissive, when true will cause the SELinux profile to only log violations instead of enforcing them.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *SecurityProfilesOperatorXK8SIoSelinuxProfileV1Alpha2DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *SecurityProfilesOperatorXK8SIoSelinuxProfileV1Alpha2DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_security_profiles_operator_x_k8s_io_selinux_profile_v1alpha2")

	var data SecurityProfilesOperatorXK8SIoSelinuxProfileV1Alpha2DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "security-profiles-operator.x-k8s.io", Version: "v1alpha2", Resource: "selinuxprofiles"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse SecurityProfilesOperatorXK8SIoSelinuxProfileV1Alpha2DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("security-profiles-operator.x-k8s.io/v1alpha2")
	data.Kind = pointer.String("SelinuxProfile")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
