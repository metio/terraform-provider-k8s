/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package site_superedge_io_v1alpha1

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
	_ datasource.DataSource              = &SiteSuperedgeIoNodeUnitV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &SiteSuperedgeIoNodeUnitV1Alpha1DataSource{}
)

func NewSiteSuperedgeIoNodeUnitV1Alpha1DataSource() datasource.DataSource {
	return &SiteSuperedgeIoNodeUnitV1Alpha1DataSource{}
}

type SiteSuperedgeIoNodeUnitV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type SiteSuperedgeIoNodeUnitV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Nodes    *[]string `tfsdk:"nodes" json:"nodes,omitempty"`
		Selector *struct {
			Annotations      *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"selector" json:"selector,omitempty"`
		Setnode *struct {
			Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Taints      *[]struct {
				Effect    *string `tfsdk:"effect" json:"effect,omitempty"`
				Key       *string `tfsdk:"key" json:"key,omitempty"`
				TimeAdded *string `tfsdk:"time_added" json:"timeAdded,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"taints" json:"taints,omitempty"`
		} `tfsdk:"setnode" json:"setnode,omitempty"`
		Taints *[]struct {
			Effect    *string `tfsdk:"effect" json:"effect,omitempty"`
			Key       *string `tfsdk:"key" json:"key,omitempty"`
			TimeAdded *string `tfsdk:"time_added" json:"timeAdded,omitempty"`
			Value     *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"taints" json:"taints,omitempty"`
		Type          *string `tfsdk:"type" json:"type,omitempty"`
		Unschedulable *bool   `tfsdk:"unschedulable" json:"unschedulable,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SiteSuperedgeIoNodeUnitV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_site_superedge_io_node_unit_v1alpha1"
}

func (r *SiteSuperedgeIoNodeUnitV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "NodeUnit is the Schema for the nodeunits API",
		MarkdownDescription: "NodeUnit is the Schema for the nodeunits API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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
				Description:         "NodeUnitSpec defines the desired state of NodeUnit",
				MarkdownDescription: "NodeUnitSpec defines the desired state of NodeUnit",
				Attributes: map[string]schema.Attribute{
					"nodes": schema.ListAttribute{
						Description:         "If specified, If node exists, join nodeunit directly",
						MarkdownDescription: "If specified, If node exists, join nodeunit directly",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"selector": schema.SingleNestedAttribute{
						Description:         "If specified, Label selector for nodes.",
						MarkdownDescription: "If specified, Label selector for nodes.",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "If specified, select node to join nodeUnit according to Annotations",
								MarkdownDescription: "If specified, select node to join nodeUnit according to Annotations",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"match_expressions": schema.ListNestedAttribute{
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "key is the label key that the selector applies to.",
											MarkdownDescription: "key is the label key that the selector applies to.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"operator": schema.StringAttribute{
											Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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

							"match_labels": schema.MapAttribute{
								Description:         "matchLabels is a map of {key,value} pairs.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs.",
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

					"setnode": schema.SingleNestedAttribute{
						Description:         "If specified, set the relevant properties to the node of nodeunit.",
						MarkdownDescription: "If specified, set the relevant properties to the node of nodeunit.",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "If specified, set annotations to all nodes of nodeunit",
								MarkdownDescription: "If specified, set annotations to all nodes of nodeunit",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"labels": schema.MapAttribute{
								Description:         "If specified, set labels to all nodes of nodeunit",
								MarkdownDescription: "If specified, set labels to all nodes of nodeunit",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"taints": schema.ListNestedAttribute{
								Description:         "If specified, set taints to all nodes of nodeunit",
								MarkdownDescription: "If specified, set taints to all nodes of nodeunit",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"effect": schema.StringAttribute{
											Description:         "Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
											MarkdownDescription: "Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"key": schema.StringAttribute{
											Description:         "Required. The taint key to be applied to a node.",
											MarkdownDescription: "Required. The taint key to be applied to a node.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"time_added": schema.StringAttribute{
											Description:         "TimeAdded represents the time at which the taint was added. It is only written for NoExecute taints.",
											MarkdownDescription: "TimeAdded represents the time at which the taint was added. It is only written for NoExecute taints.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"value": schema.StringAttribute{
											Description:         "The taint value corresponding to the taint key.",
											MarkdownDescription: "The taint value corresponding to the taint key.",
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

					"taints": schema.ListNestedAttribute{
						Description:         "If specified, allow to set taints to nodeunit for the scheduler to choose",
						MarkdownDescription: "If specified, allow to set taints to nodeunit for the scheduler to choose",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"effect": schema.StringAttribute{
									Description:         "Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
									MarkdownDescription: "Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"key": schema.StringAttribute{
									Description:         "Required. The taint key to be applied to a node.",
									MarkdownDescription: "Required. The taint key to be applied to a node.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"time_added": schema.StringAttribute{
									Description:         "TimeAdded represents the time at which the taint was added. It is only written for NoExecute taints.",
									MarkdownDescription: "TimeAdded represents the time at which the taint was added. It is only written for NoExecute taints.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"value": schema.StringAttribute{
									Description:         "The taint value corresponding to the taint key.",
									MarkdownDescription: "The taint value corresponding to the taint key.",
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

					"type": schema.StringAttribute{
						Description:         "Type of nodeunit， vaule: Cloud、Edge",
						MarkdownDescription: "Type of nodeunit， vaule: Cloud、Edge",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"unschedulable": schema.BoolAttribute{
						Description:         "Unschedulable controls nodeUnit schedulability of new workwolads. By default, nodeUnit is schedulable.",
						MarkdownDescription: "Unschedulable controls nodeUnit schedulability of new workwolads. By default, nodeUnit is schedulable.",
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

func (r *SiteSuperedgeIoNodeUnitV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *SiteSuperedgeIoNodeUnitV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_site_superedge_io_node_unit_v1alpha1")

	var data SiteSuperedgeIoNodeUnitV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "site.superedge.io", Version: "v1alpha1", Resource: "NodeUnit"}).
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

	var readResponse SiteSuperedgeIoNodeUnitV1Alpha1DataSourceData
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

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("site.superedge.io/v1alpha1")
	data.Kind = pointer.String("NodeUnit")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}