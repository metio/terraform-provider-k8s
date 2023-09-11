/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package capsule_clastix_io_v1alpha1

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
	_ datasource.DataSource              = &CapsuleClastixIoCapsuleConfigurationV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &CapsuleClastixIoCapsuleConfigurationV1Alpha1DataSource{}
)

func NewCapsuleClastixIoCapsuleConfigurationV1Alpha1DataSource() datasource.DataSource {
	return &CapsuleClastixIoCapsuleConfigurationV1Alpha1DataSource{}
}

type CapsuleClastixIoCapsuleConfigurationV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type CapsuleClastixIoCapsuleConfigurationV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ForceTenantPrefix       *bool     `tfsdk:"force_tenant_prefix" json:"forceTenantPrefix,omitempty"`
		ProtectedNamespaceRegex *string   `tfsdk:"protected_namespace_regex" json:"protectedNamespaceRegex,omitempty"`
		UserGroups              *[]string `tfsdk:"user_groups" json:"userGroups,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CapsuleClastixIoCapsuleConfigurationV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_capsule_clastix_io_capsule_configuration_v1alpha1"
}

func (r *CapsuleClastixIoCapsuleConfigurationV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CapsuleConfiguration is the Schema for the Capsule configuration API.",
		MarkdownDescription: "CapsuleConfiguration is the Schema for the Capsule configuration API.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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
				Description:         "CapsuleConfigurationSpec defines the Capsule configuration.",
				MarkdownDescription: "CapsuleConfigurationSpec defines the Capsule configuration.",
				Attributes: map[string]schema.Attribute{
					"force_tenant_prefix": schema.BoolAttribute{
						Description:         "Enforces the Tenant owner, during Namespace creation, to name it using the selected Tenant name as prefix, separated by a dash. This is useful to avoid Namespace name collision in a public CaaS environment.",
						MarkdownDescription: "Enforces the Tenant owner, during Namespace creation, to name it using the selected Tenant name as prefix, separated by a dash. This is useful to avoid Namespace name collision in a public CaaS environment.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"protected_namespace_regex": schema.StringAttribute{
						Description:         "Disallow creation of namespaces, whose name matches this regexp",
						MarkdownDescription: "Disallow creation of namespaces, whose name matches this regexp",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"user_groups": schema.ListAttribute{
						Description:         "Names of the groups for Capsule users.",
						MarkdownDescription: "Names of the groups for Capsule users.",
						ElementType:         types.StringType,
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

func (r *CapsuleClastixIoCapsuleConfigurationV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *CapsuleClastixIoCapsuleConfigurationV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_capsule_clastix_io_capsule_configuration_v1alpha1")

	var data CapsuleClastixIoCapsuleConfigurationV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "capsule.clastix.io", Version: "v1alpha1", Resource: "capsuleconfigurations"}).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetResourceError(err, data.Metadata.Name))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse CapsuleClastixIoCapsuleConfigurationV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("capsule.clastix.io/v1alpha1")
	data.Kind = pointer.String("CapsuleConfiguration")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
