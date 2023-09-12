/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package cilium_io_v2

import (
	"context"
	"encoding/json"
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
	_ datasource.DataSource              = &CiliumIoCiliumIdentityV2DataSource{}
	_ datasource.DataSourceWithConfigure = &CiliumIoCiliumIdentityV2DataSource{}
)

func NewCiliumIoCiliumIdentityV2DataSource() datasource.DataSource {
	return &CiliumIoCiliumIdentityV2DataSource{}
}

type CiliumIoCiliumIdentityV2DataSource struct {
	kubernetesClient dynamic.Interface
}

type CiliumIoCiliumIdentityV2DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Security_labels *map[string]string `tfsdk:"security_labels" json:"security-labels,omitempty"`
}

func (r *CiliumIoCiliumIdentityV2DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cilium_io_cilium_identity_v2"
}

func (r *CiliumIoCiliumIdentityV2DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CiliumIdentity is a CRD that represents an identity managed by Cilium. It is intended as a backing store for identity allocation, acting as the global coordination backend, and can be used in place of a KVStore (such as etcd). The name of the CRD is the numeric identity and the labels on the CRD object are the kubernetes sourced labels seen by cilium. This is currently the only label source possible when running under kubernetes. Non-kubernetes labels are filtered but all labels, from all sources, are places in the SecurityLabels field. These also include the source and are used to define the identity. The labels under metav1.ObjectMeta can be used when searching for CiliumIdentity instances that include particular labels. This can be done with invocations such as:  kubectl get ciliumid -l 'foo=bar'",
		MarkdownDescription: "CiliumIdentity is a CRD that represents an identity managed by Cilium. It is intended as a backing store for identity allocation, acting as the global coordination backend, and can be used in place of a KVStore (such as etcd). The name of the CRD is the numeric identity and the labels on the CRD object are the kubernetes sourced labels seen by cilium. This is currently the only label source possible when running under kubernetes. Non-kubernetes labels are filtered but all labels, from all sources, are places in the SecurityLabels field. These also include the source and are used to define the identity. The labels under metav1.ObjectMeta can be used when searching for CiliumIdentity instances that include particular labels. This can be done with invocations such as:  kubectl get ciliumid -l 'foo=bar'",
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

			"security_labels": schema.MapAttribute{
				Description:         "SecurityLabels is the source-of-truth set of labels for this identity.",
				MarkdownDescription: "SecurityLabels is the source-of-truth set of labels for this identity.",
				ElementType:         types.StringType,
				Required:            false,
				Optional:            false,
				Computed:            true,
			},
		},
	}
}

func (r *CiliumIoCiliumIdentityV2DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *CiliumIoCiliumIdentityV2DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_cilium_io_cilium_identity_v2")

	var data CiliumIoCiliumIdentityV2DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "cilium.io", Version: "v2", Resource: "ciliumidentities"}).
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

	var readResponse CiliumIoCiliumIdentityV2DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("cilium.io/v2")
	data.Kind = pointer.String("CiliumIdentity")
	data.Metadata = readResponse.Metadata
	data.Security_labels = readResponse.Security_labels

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
