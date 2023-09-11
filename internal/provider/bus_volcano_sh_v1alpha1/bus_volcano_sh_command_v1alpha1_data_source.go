/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package bus_volcano_sh_v1alpha1

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
	_ datasource.DataSource              = &BusVolcanoShCommandV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &BusVolcanoShCommandV1Alpha1DataSource{}
)

func NewBusVolcanoShCommandV1Alpha1DataSource() datasource.DataSource {
	return &BusVolcanoShCommandV1Alpha1DataSource{}
}

type BusVolcanoShCommandV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type BusVolcanoShCommandV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Action  *string `tfsdk:"action" json:"action,omitempty"`
	Message *string `tfsdk:"message" json:"message,omitempty"`
	Reason  *string `tfsdk:"reason" json:"reason,omitempty"`
	Target  *struct {
		ApiVersion         *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
		BlockOwnerDeletion *bool   `tfsdk:"block_owner_deletion" json:"blockOwnerDeletion,omitempty"`
		Controller         *bool   `tfsdk:"controller" json:"controller,omitempty"`
		Kind               *string `tfsdk:"kind" json:"kind,omitempty"`
		Name               *string `tfsdk:"name" json:"name,omitempty"`
		Uid                *string `tfsdk:"uid" json:"uid,omitempty"`
	} `tfsdk:"target" json:"target,omitempty"`
}

func (r *BusVolcanoShCommandV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_bus_volcano_sh_command_v1alpha1"
}

func (r *BusVolcanoShCommandV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Command defines command structure.",
		MarkdownDescription: "Command defines command structure.",
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

			"action": schema.StringAttribute{
				Description:         "Action defines the action that will be took to the target object.",
				MarkdownDescription: "Action defines the action that will be took to the target object.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"message": schema.StringAttribute{
				Description:         "Human-readable message indicating details of this command.",
				MarkdownDescription: "Human-readable message indicating details of this command.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"reason": schema.StringAttribute{
				Description:         "Unique, one-word, CamelCase reason for this command.",
				MarkdownDescription: "Unique, one-word, CamelCase reason for this command.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"target": schema.SingleNestedAttribute{
				Description:         "TargetObject defines the target object of this command.",
				MarkdownDescription: "TargetObject defines the target object of this command.",
				Attributes: map[string]schema.Attribute{
					"api_version": schema.StringAttribute{
						Description:         "API version of the referent.",
						MarkdownDescription: "API version of the referent.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"block_owner_deletion": schema.BoolAttribute{
						Description:         "If true, AND if the owner has the 'foregroundDeletion' finalizer, then the owner cannot be deleted from the key-value store until this reference is removed. See https://kubernetes.io/docs/concepts/architecture/garbage-collection/#foreground-deletion for how the garbage collector interacts with this field and enforces the foreground deletion. Defaults to false. To set this field, a user needs 'delete' permission of the owner, otherwise 422 (Unprocessable Entity) will be returned.",
						MarkdownDescription: "If true, AND if the owner has the 'foregroundDeletion' finalizer, then the owner cannot be deleted from the key-value store until this reference is removed. See https://kubernetes.io/docs/concepts/architecture/garbage-collection/#foreground-deletion for how the garbage collector interacts with this field and enforces the foreground deletion. Defaults to false. To set this field, a user needs 'delete' permission of the owner, otherwise 422 (Unprocessable Entity) will be returned.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"controller": schema.BoolAttribute{
						Description:         "If true, this reference points to the managing controller.",
						MarkdownDescription: "If true, this reference points to the managing controller.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"kind": schema.StringAttribute{
						Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
						MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"name": schema.StringAttribute{
						Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
						MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"uid": schema.StringAttribute{
						Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#uids",
						MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#uids",
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

func (r *BusVolcanoShCommandV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *BusVolcanoShCommandV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_bus_volcano_sh_command_v1alpha1")

	var data BusVolcanoShCommandV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "bus.volcano.sh", Version: "v1alpha1", Resource: "commands"}).
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

	var readResponse BusVolcanoShCommandV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("bus.volcano.sh/v1alpha1")
	data.Kind = pointer.String("Command")
	data.Metadata = readResponse.Metadata
	data.Action = readResponse.Action
	data.Message = readResponse.Message
	data.Reason = readResponse.Reason
	data.Target = readResponse.Target

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
