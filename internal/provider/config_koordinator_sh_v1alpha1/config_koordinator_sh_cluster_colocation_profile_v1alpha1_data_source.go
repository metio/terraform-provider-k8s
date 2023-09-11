/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package config_koordinator_sh_v1alpha1

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
	_ datasource.DataSource              = &ConfigKoordinatorShClusterColocationProfileV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &ConfigKoordinatorShClusterColocationProfileV1Alpha1DataSource{}
)

func NewConfigKoordinatorShClusterColocationProfileV1Alpha1DataSource() datasource.DataSource {
	return &ConfigKoordinatorShClusterColocationProfileV1Alpha1DataSource{}
}

type ConfigKoordinatorShClusterColocationProfileV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type ConfigKoordinatorShClusterColocationProfileV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Annotations         *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
		KoordinatorPriority *int64             `tfsdk:"koordinator_priority" json:"koordinatorPriority,omitempty"`
		Labels              *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		NamespaceSelector   *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
		Patch             *map[string]string `tfsdk:"patch" json:"patch,omitempty"`
		PriorityClassName *string            `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
		Probability       *string            `tfsdk:"probability" json:"probability,omitempty"`
		QosClass          *string            `tfsdk:"qos_class" json:"qosClass,omitempty"`
		SchedulerName     *string            `tfsdk:"scheduler_name" json:"schedulerName,omitempty"`
		Selector          *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"selector" json:"selector,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ConfigKoordinatorShClusterColocationProfileV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_config_koordinator_sh_cluster_colocation_profile_v1alpha1"
}

func (r *ConfigKoordinatorShClusterColocationProfileV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClusterColocationProfile is the Schema for the ClusterColocationProfile API",
		MarkdownDescription: "ClusterColocationProfile is the Schema for the ClusterColocationProfile API",
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
				Description:         "ClusterColocationProfileSpec is a description of a ClusterColocationProfile.",
				MarkdownDescription: "ClusterColocationProfileSpec is a description of a ClusterColocationProfile.",
				Attributes: map[string]schema.Attribute{
					"annotations": schema.MapAttribute{
						Description:         "Annotations describes the k/v pair that needs to inject into Pod.Annotations",
						MarkdownDescription: "Annotations describes the k/v pair that needs to inject into Pod.Annotations",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"koordinator_priority": schema.Int64Attribute{
						Description:         "KoordinatorPriority defines the Pod sub-priority in Koordinator. The priority value will be injected into Pod as label koordinator.sh/priority. Various Koordinator components determine the priority of the Pod in the Koordinator through KoordinatorPriority and the priority value in PriorityClassName. The higher the value, the higher the priority.",
						MarkdownDescription: "KoordinatorPriority defines the Pod sub-priority in Koordinator. The priority value will be injected into Pod as label koordinator.sh/priority. Various Koordinator components determine the priority of the Pod in the Koordinator through KoordinatorPriority and the priority value in PriorityClassName. The higher the value, the higher the priority.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"labels": schema.MapAttribute{
						Description:         "Labels describes the k/v pair that needs to inject into Pod.Labels",
						MarkdownDescription: "Labels describes the k/v pair that needs to inject into Pod.Labels",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"namespace_selector": schema.SingleNestedAttribute{
						Description:         "NamespaceSelector decides whether to mutate/validate Pods if the namespace matches the selector. Default to the empty LabelSelector, which matches everything.",
						MarkdownDescription: "NamespaceSelector decides whether to mutate/validate Pods if the namespace matches the selector. Default to the empty LabelSelector, which matches everything.",
						Attributes: map[string]schema.Attribute{
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
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

					"patch": schema.MapAttribute{
						Description:         "Patch indicates patching podTemplate that will be injected to the Pod.",
						MarkdownDescription: "Patch indicates patching podTemplate that will be injected to the Pod.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"priority_class_name": schema.StringAttribute{
						Description:         "If specified, the priorityClassName and the priority value defined in PriorityClass will be injected into the Pod. The PriorityClassName, priority value in PriorityClassName and KoordinatorPriority will affect the scheduling, preemption and other behaviors of Koordinator system.",
						MarkdownDescription: "If specified, the priorityClassName and the priority value defined in PriorityClass will be injected into the Pod. The PriorityClassName, priority value in PriorityClassName and KoordinatorPriority will affect the scheduling, preemption and other behaviors of Koordinator system.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"probability": schema.StringAttribute{
						Description:         "Probability indicates profile will make effect with a probability.",
						MarkdownDescription: "Probability indicates profile will make effect with a probability.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"qos_class": schema.StringAttribute{
						Description:         "QoSClass describes the type of Koordinator QoS that the Pod is running. The value will be injected into Pod as label koordinator.sh/qosClass. Options are LSE/LSR/LS/BE/SYSTEM.",
						MarkdownDescription: "QoSClass describes the type of Koordinator QoS that the Pod is running. The value will be injected into Pod as label koordinator.sh/qosClass. Options are LSE/LSR/LS/BE/SYSTEM.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"scheduler_name": schema.StringAttribute{
						Description:         "If specified, the pod will be dispatched by specified scheduler.",
						MarkdownDescription: "If specified, the pod will be dispatched by specified scheduler.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"selector": schema.SingleNestedAttribute{
						Description:         "Selector decides whether to mutate/validate Pods if the Pod matches the selector. Default to the empty LabelSelector, which matches everything.",
						MarkdownDescription: "Selector decides whether to mutate/validate Pods if the Pod matches the selector. Default to the empty LabelSelector, which matches everything.",
						Attributes: map[string]schema.Attribute{
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
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *ConfigKoordinatorShClusterColocationProfileV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *ConfigKoordinatorShClusterColocationProfileV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_config_koordinator_sh_cluster_colocation_profile_v1alpha1")

	var data ConfigKoordinatorShClusterColocationProfileV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "config.koordinator.sh", Version: "v1alpha1", Resource: "clustercolocationprofiles"}).
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

	var readResponse ConfigKoordinatorShClusterColocationProfileV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("config.koordinator.sh/v1alpha1")
	data.Kind = pointer.String("ClusterColocationProfile")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
