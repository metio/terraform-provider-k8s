/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package services_k8s_aws_v1alpha1

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
	_ datasource.DataSource              = &ServicesK8SAwsFieldExportV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &ServicesK8SAwsFieldExportV1Alpha1DataSource{}
)

func NewServicesK8SAwsFieldExportV1Alpha1DataSource() datasource.DataSource {
	return &ServicesK8SAwsFieldExportV1Alpha1DataSource{}
}

type ServicesK8SAwsFieldExportV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type ServicesK8SAwsFieldExportV1Alpha1DataSourceData struct {
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
		From *struct {
			Path     *string `tfsdk:"path" json:"path,omitempty"`
			Resource *struct {
				Group *string `tfsdk:"group" json:"group,omitempty"`
				Kind  *string `tfsdk:"kind" json:"kind,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"resource" json:"resource,omitempty"`
		} `tfsdk:"from" json:"from,omitempty"`
		To *struct {
			Key       *string `tfsdk:"key" json:"key,omitempty"`
			Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"to" json:"to,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ServicesK8SAwsFieldExportV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_services_k8s_aws_field_export_v1alpha1"
}

func (r *ServicesK8SAwsFieldExportV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "FieldExport is the schema for the FieldExport API.",
		MarkdownDescription: "FieldExport is the schema for the FieldExport API.",
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
				Description:         "FieldExportSpec defines the desired state of the FieldExport.",
				MarkdownDescription: "FieldExportSpec defines the desired state of the FieldExport.",
				Attributes: map[string]schema.Attribute{
					"from": schema.SingleNestedAttribute{
						Description:         "ResourceFieldSelector provides the values necessary to identify an individual field on an individual K8s resource.",
						MarkdownDescription: "ResourceFieldSelector provides the values necessary to identify an individual field on an individual K8s resource.",
						Attributes: map[string]schema.Attribute{
							"path": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"resource": schema.SingleNestedAttribute{
								Description:         "NamespacedResource provides all the values necessary to identify an ACK resource of a given type (within the same namespace as the custom resource containing this type).",
								MarkdownDescription: "NamespacedResource provides all the values necessary to identify an ACK resource of a given type (within the same namespace as the custom resource containing this type).",
								Attributes: map[string]schema.Attribute{
									"group": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"kind": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
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
						Required: false,
						Optional: false,
						Computed: true,
					},

					"to": schema.SingleNestedAttribute{
						Description:         "FieldExportTarget provides the values necessary to identify the output path for a field export.",
						MarkdownDescription: "FieldExportTarget provides the values necessary to identify the output path for a field export.",
						Attributes: map[string]schema.Attribute{
							"key": schema.StringAttribute{
								Description:         "Key overrides the default value ('<namespace>.<FieldExport-resource-name>') for the FieldExport target",
								MarkdownDescription: "Key overrides the default value ('<namespace>.<FieldExport-resource-name>') for the FieldExport target",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"kind": schema.StringAttribute{
								Description:         "FieldExportOutputType represents all types that can be produced by a field export operation",
								MarkdownDescription: "FieldExportOutputType represents all types that can be produced by a field export operation",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace is marked as optional, so we cannot compose 'NamespacedName'",
								MarkdownDescription: "Namespace is marked as optional, so we cannot compose 'NamespacedName'",
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
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *ServicesK8SAwsFieldExportV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *ServicesK8SAwsFieldExportV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_services_k8s_aws_field_export_v1alpha1")

	var data ServicesK8SAwsFieldExportV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "services.k8s.aws", Version: "v1alpha1", Resource: "fieldexports"}).
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

	var readResponse ServicesK8SAwsFieldExportV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("services.k8s.aws/v1alpha1")
	data.Kind = pointer.String("FieldExport")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
