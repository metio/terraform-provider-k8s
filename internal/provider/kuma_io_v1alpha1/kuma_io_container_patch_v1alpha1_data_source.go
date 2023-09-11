/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kuma_io_v1alpha1

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
	_ datasource.DataSource              = &KumaIoContainerPatchV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &KumaIoContainerPatchV1Alpha1DataSource{}
)

func NewKumaIoContainerPatchV1Alpha1DataSource() datasource.DataSource {
	return &KumaIoContainerPatchV1Alpha1DataSource{}
}

type KumaIoContainerPatchV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type KumaIoContainerPatchV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Mesh *string `tfsdk:"mesh" json:"mesh,omitempty"`
	Spec *struct {
		InitPatch *[]struct {
			From  *string `tfsdk:"from" json:"from,omitempty"`
			Op    *string `tfsdk:"op" json:"op,omitempty"`
			Path  *string `tfsdk:"path" json:"path,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"init_patch" json:"initPatch,omitempty"`
		SidecarPatch *[]struct {
			From  *string `tfsdk:"from" json:"from,omitempty"`
			Op    *string `tfsdk:"op" json:"op,omitempty"`
			Path  *string `tfsdk:"path" json:"path,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"sidecar_patch" json:"sidecarPatch,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KumaIoContainerPatchV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kuma_io_container_patch_v1alpha1"
}

func (r *KumaIoContainerPatchV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ContainerPatch stores a list of patches to apply to init and sidecar containers.",
		MarkdownDescription: "ContainerPatch stores a list of patches to apply to init and sidecar containers.",
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

			"mesh": schema.StringAttribute{
				Description:         "",
				MarkdownDescription: "",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "ContainerPatchSpec specifies the options available for a ContainerPatch",
				MarkdownDescription: "ContainerPatchSpec specifies the options available for a ContainerPatch",
				Attributes: map[string]schema.Attribute{
					"init_patch": schema.ListNestedAttribute{
						Description:         "InitPatch specifies jsonpatch to apply to an init container.",
						MarkdownDescription: "InitPatch specifies jsonpatch to apply to an init container.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"from": schema.StringAttribute{
									Description:         "From is a jsonpatch from string, used by move and copy operations.",
									MarkdownDescription: "From is a jsonpatch from string, used by move and copy operations.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"op": schema.StringAttribute{
									Description:         "Op is a jsonpatch operation string.",
									MarkdownDescription: "Op is a jsonpatch operation string.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"path": schema.StringAttribute{
									Description:         "Path is a jsonpatch path string.",
									MarkdownDescription: "Path is a jsonpatch path string.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"value": schema.StringAttribute{
									Description:         "Value must be a string representing a valid json object used by replace and add operations. String has to be escaped with ' to be valid a json object.",
									MarkdownDescription: "Value must be a string representing a valid json object used by replace and add operations. String has to be escaped with ' to be valid a json object.",
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

					"sidecar_patch": schema.ListNestedAttribute{
						Description:         "SidecarPatch specifies jsonpatch to apply to a sidecar container.",
						MarkdownDescription: "SidecarPatch specifies jsonpatch to apply to a sidecar container.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"from": schema.StringAttribute{
									Description:         "From is a jsonpatch from string, used by move and copy operations.",
									MarkdownDescription: "From is a jsonpatch from string, used by move and copy operations.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"op": schema.StringAttribute{
									Description:         "Op is a jsonpatch operation string.",
									MarkdownDescription: "Op is a jsonpatch operation string.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"path": schema.StringAttribute{
									Description:         "Path is a jsonpatch path string.",
									MarkdownDescription: "Path is a jsonpatch path string.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"value": schema.StringAttribute{
									Description:         "Value must be a string representing a valid json object used by replace and add operations. String has to be escaped with ' to be valid a json object.",
									MarkdownDescription: "Value must be a string representing a valid json object used by replace and add operations. String has to be escaped with ' to be valid a json object.",
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

func (r *KumaIoContainerPatchV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *KumaIoContainerPatchV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_kuma_io_container_patch_v1alpha1")

	var data KumaIoContainerPatchV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "kuma.io", Version: "v1alpha1", Resource: "containerpatches"}).
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

	var readResponse KumaIoContainerPatchV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("kuma.io/v1alpha1")
	data.Kind = pointer.String("ContainerPatch")
	data.Metadata = readResponse.Metadata
	data.Mesh = readResponse.Mesh
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
