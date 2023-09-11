/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package core_v1

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
	_ datasource.DataSource              = &LimitRangeV1DataSource{}
	_ datasource.DataSourceWithConfigure = &LimitRangeV1DataSource{}
)

func NewLimitRangeV1DataSource() datasource.DataSource {
	return &LimitRangeV1DataSource{}
}

type LimitRangeV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type LimitRangeV1DataSourceData struct {
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
		Limits *[]struct {
			Default              *map[string]string `tfsdk:"default" json:"default,omitempty"`
			DefaultRequest       *map[string]string `tfsdk:"default_request" json:"defaultRequest,omitempty"`
			Max                  *map[string]string `tfsdk:"max" json:"max,omitempty"`
			MaxLimitRequestRatio *map[string]string `tfsdk:"max_limit_request_ratio" json:"maxLimitRequestRatio,omitempty"`
			Min                  *map[string]string `tfsdk:"min" json:"min,omitempty"`
			Type                 *string            `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"limits" json:"limits,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *LimitRangeV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_limit_range_v1"
}

func (r *LimitRangeV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "LimitRange sets resource usage limits for each kind of resource in a Namespace.",
		MarkdownDescription: "LimitRange sets resource usage limits for each kind of resource in a Namespace.",
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
				Description:         "LimitRangeSpec defines a min/max usage limit for resources that match on kind.",
				MarkdownDescription: "LimitRangeSpec defines a min/max usage limit for resources that match on kind.",
				Attributes: map[string]schema.Attribute{
					"limits": schema.ListNestedAttribute{
						Description:         "Limits is the list of LimitRangeItem objects that are enforced.",
						MarkdownDescription: "Limits is the list of LimitRangeItem objects that are enforced.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"default": schema.MapAttribute{
									Description:         "Default resource requirement limit value by resource name if resource limit is omitted.",
									MarkdownDescription: "Default resource requirement limit value by resource name if resource limit is omitted.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"default_request": schema.MapAttribute{
									Description:         "DefaultRequest is the default resource requirement request value by resource name if resource request is omitted.",
									MarkdownDescription: "DefaultRequest is the default resource requirement request value by resource name if resource request is omitted.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"max": schema.MapAttribute{
									Description:         "Max usage constraints on this kind by resource name.",
									MarkdownDescription: "Max usage constraints on this kind by resource name.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"max_limit_request_ratio": schema.MapAttribute{
									Description:         "MaxLimitRequestRatio if specified, the named resource must have a request and limit that are both non-zero where limit divided by request is less than or equal to the enumerated value; this represents the max burst for the named resource.",
									MarkdownDescription: "MaxLimitRequestRatio if specified, the named resource must have a request and limit that are both non-zero where limit divided by request is less than or equal to the enumerated value; this represents the max burst for the named resource.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"min": schema.MapAttribute{
									Description:         "Min usage constraints on this kind by resource name.",
									MarkdownDescription: "Min usage constraints on this kind by resource name.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"type": schema.StringAttribute{
									Description:         "Type of resource that this limit applies to.",
									MarkdownDescription: "Type of resource that this limit applies to.",
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

func (r *LimitRangeV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *LimitRangeV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_limit_range_v1")

	var data LimitRangeV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "", Version: "v1", Resource: "limitranges"}).
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

	var readResponse LimitRangeV1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("v1")
	data.Kind = pointer.String("LimitRange")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
