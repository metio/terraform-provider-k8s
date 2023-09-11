/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package chaos_mesh_org_v1alpha1

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
	_ datasource.DataSource              = &ChaosMeshOrgJvmchaosV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &ChaosMeshOrgJvmchaosV1Alpha1DataSource{}
)

func NewChaosMeshOrgJvmchaosV1Alpha1DataSource() datasource.DataSource {
	return &ChaosMeshOrgJvmchaosV1Alpha1DataSource{}
}

type ChaosMeshOrgJvmchaosV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type ChaosMeshOrgJvmchaosV1Alpha1DataSourceData struct {
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
		Action                *string   `tfsdk:"action" json:"action,omitempty"`
		Class                 *string   `tfsdk:"class" json:"class,omitempty"`
		ContainerNames        *[]string `tfsdk:"container_names" json:"containerNames,omitempty"`
		CpuCount              *int64    `tfsdk:"cpu_count" json:"cpuCount,omitempty"`
		Database              *string   `tfsdk:"database" json:"database,omitempty"`
		Duration              *string   `tfsdk:"duration" json:"duration,omitempty"`
		Exception             *string   `tfsdk:"exception" json:"exception,omitempty"`
		Latency               *int64    `tfsdk:"latency" json:"latency,omitempty"`
		MemType               *string   `tfsdk:"mem_type" json:"memType,omitempty"`
		Method                *string   `tfsdk:"method" json:"method,omitempty"`
		Mode                  *string   `tfsdk:"mode" json:"mode,omitempty"`
		MysqlConnectorVersion *string   `tfsdk:"mysql_connector_version" json:"mysqlConnectorVersion,omitempty"`
		Name                  *string   `tfsdk:"name" json:"name,omitempty"`
		Pid                   *int64    `tfsdk:"pid" json:"pid,omitempty"`
		Port                  *int64    `tfsdk:"port" json:"port,omitempty"`
		RemoteCluster         *string   `tfsdk:"remote_cluster" json:"remoteCluster,omitempty"`
		RuleData              *string   `tfsdk:"rule_data" json:"ruleData,omitempty"`
		Selector              *struct {
			AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" json:"annotationSelectors,omitempty"`
			ExpressionSelectors *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"expression_selectors" json:"expressionSelectors,omitempty"`
			FieldSelectors    *map[string]string   `tfsdk:"field_selectors" json:"fieldSelectors,omitempty"`
			LabelSelectors    *map[string]string   `tfsdk:"label_selectors" json:"labelSelectors,omitempty"`
			Namespaces        *[]string            `tfsdk:"namespaces" json:"namespaces,omitempty"`
			NodeSelectors     *map[string]string   `tfsdk:"node_selectors" json:"nodeSelectors,omitempty"`
			Nodes             *[]string            `tfsdk:"nodes" json:"nodes,omitempty"`
			PodPhaseSelectors *[]string            `tfsdk:"pod_phase_selectors" json:"podPhaseSelectors,omitempty"`
			Pods              *map[string][]string `tfsdk:"pods" json:"pods,omitempty"`
		} `tfsdk:"selector" json:"selector,omitempty"`
		SqlType *string `tfsdk:"sql_type" json:"sqlType,omitempty"`
		Table   *string `tfsdk:"table" json:"table,omitempty"`
		Value   *string `tfsdk:"value" json:"value,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ChaosMeshOrgJvmchaosV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_chaos_mesh_org_jvm_chaos_v1alpha1"
}

func (r *ChaosMeshOrgJvmchaosV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "JVMChaos is the Schema for the jvmchaos API",
		MarkdownDescription: "JVMChaos is the Schema for the jvmchaos API",
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
				Description:         "JVMChaosSpec defines the desired state of JVMChaos",
				MarkdownDescription: "JVMChaosSpec defines the desired state of JVMChaos",
				Attributes: map[string]schema.Attribute{
					"action": schema.StringAttribute{
						Description:         "Action defines the specific jvm chaos action. Supported action: latency;return;exception;stress;gc;ruleData",
						MarkdownDescription: "Action defines the specific jvm chaos action. Supported action: latency;return;exception;stress;gc;ruleData",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"class": schema.StringAttribute{
						Description:         "Java class",
						MarkdownDescription: "Java class",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"container_names": schema.ListAttribute{
						Description:         "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
						MarkdownDescription: "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"cpu_count": schema.Int64Attribute{
						Description:         "the CPU core number needs to use, only set it when action is stress",
						MarkdownDescription: "the CPU core number needs to use, only set it when action is stress",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"database": schema.StringAttribute{
						Description:         "the match database default value is '', means match all database",
						MarkdownDescription: "the match database default value is '', means match all database",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"duration": schema.StringAttribute{
						Description:         "Duration represents the duration of the chaos action",
						MarkdownDescription: "Duration represents the duration of the chaos action",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"exception": schema.StringAttribute{
						Description:         "the exception which needs to throw for action 'exception' or the exception message needs to throw in action 'mysql'",
						MarkdownDescription: "the exception which needs to throw for action 'exception' or the exception message needs to throw in action 'mysql'",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"latency": schema.Int64Attribute{
						Description:         "the latency duration for action 'latency', unit ms or the latency duration in action 'mysql'",
						MarkdownDescription: "the latency duration for action 'latency', unit ms or the latency duration in action 'mysql'",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"mem_type": schema.StringAttribute{
						Description:         "the memory type needs to locate, only set it when action is stress, the value can be 'stack' or 'heap'",
						MarkdownDescription: "the memory type needs to locate, only set it when action is stress, the value can be 'stack' or 'heap'",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"method": schema.StringAttribute{
						Description:         "the method in Java class",
						MarkdownDescription: "the method in Java class",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"mode": schema.StringAttribute{
						Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
						MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"mysql_connector_version": schema.StringAttribute{
						Description:         "the version of mysql-connector-java, only support 5.X.X(set to '5') and 8.X.X(set to '8') now",
						MarkdownDescription: "the version of mysql-connector-java, only support 5.X.X(set to '5') and 8.X.X(set to '8') now",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"name": schema.StringAttribute{
						Description:         "byteman rule name, should be unique, and will generate one if not set",
						MarkdownDescription: "byteman rule name, should be unique, and will generate one if not set",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"pid": schema.Int64Attribute{
						Description:         "the pid of Java process which needs to attach",
						MarkdownDescription: "the pid of Java process which needs to attach",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"port": schema.Int64Attribute{
						Description:         "the port of agent server, default 9277",
						MarkdownDescription: "the port of agent server, default 9277",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"remote_cluster": schema.StringAttribute{
						Description:         "RemoteCluster represents the remote cluster where the chaos will be deployed",
						MarkdownDescription: "RemoteCluster represents the remote cluster where the chaos will be deployed",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"rule_data": schema.StringAttribute{
						Description:         "the byteman rule's data for action 'ruleData'",
						MarkdownDescription: "the byteman rule's data for action 'ruleData'",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"selector": schema.SingleNestedAttribute{
						Description:         "Selector is used to select pods that are used to inject chaos action.",
						MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",
						Attributes: map[string]schema.Attribute{
							"annotation_selectors": schema.MapAttribute{
								Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
								MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"expression_selectors": schema.ListNestedAttribute{
								Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
								MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
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

							"field_selectors": schema.MapAttribute{
								Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
								MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"label_selectors": schema.MapAttribute{
								Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
								MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"namespaces": schema.ListAttribute{
								Description:         "Namespaces is a set of namespace to which objects belong.",
								MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"node_selectors": schema.MapAttribute{
								Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
								MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"nodes": schema.ListAttribute{
								Description:         "Nodes is a set of node name and objects must belong to these nodes.",
								MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"pod_phase_selectors": schema.ListAttribute{
								Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
								MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"pods": schema.MapAttribute{
								Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
								MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
								ElementType:         types.ListType{ElemType: types.StringType},
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"sql_type": schema.StringAttribute{
						Description:         "the match sql type default value is '', means match all SQL type. The value can be 'select', 'insert', 'update', 'delete', 'replace'.",
						MarkdownDescription: "the match sql type default value is '', means match all SQL type. The value can be 'select', 'insert', 'update', 'delete', 'replace'.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"table": schema.StringAttribute{
						Description:         "the match table default value is '', means match all table",
						MarkdownDescription: "the match table default value is '', means match all table",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"value": schema.StringAttribute{
						Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
						MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
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

func (r *ChaosMeshOrgJvmchaosV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *ChaosMeshOrgJvmchaosV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_chaos_mesh_org_jvm_chaos_v1alpha1")

	var data ChaosMeshOrgJvmchaosV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "chaos-mesh.org", Version: "v1alpha1", Resource: "jvmchaos"}).
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

	var readResponse ChaosMeshOrgJvmchaosV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("chaos-mesh.org/v1alpha1")
	data.Kind = pointer.String("JVMChaos")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
