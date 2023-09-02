/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package config_karmada_io_v1alpha1

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
	_ datasource.DataSource              = &ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1DataSource{}
)

func NewConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1DataSource() datasource.DataSource {
	return &ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1DataSource{}
}

type ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Customizations *struct {
			DependencyInterpretation *struct {
				LuaScript *string `tfsdk:"lua_script" json:"luaScript,omitempty"`
			} `tfsdk:"dependency_interpretation" json:"dependencyInterpretation,omitempty"`
			HealthInterpretation *struct {
				LuaScript *string `tfsdk:"lua_script" json:"luaScript,omitempty"`
			} `tfsdk:"health_interpretation" json:"healthInterpretation,omitempty"`
			ReplicaResource *struct {
				LuaScript *string `tfsdk:"lua_script" json:"luaScript,omitempty"`
			} `tfsdk:"replica_resource" json:"replicaResource,omitempty"`
			ReplicaRevision *struct {
				LuaScript *string `tfsdk:"lua_script" json:"luaScript,omitempty"`
			} `tfsdk:"replica_revision" json:"replicaRevision,omitempty"`
			Retention *struct {
				LuaScript *string `tfsdk:"lua_script" json:"luaScript,omitempty"`
			} `tfsdk:"retention" json:"retention,omitempty"`
			StatusAggregation *struct {
				LuaScript *string `tfsdk:"lua_script" json:"luaScript,omitempty"`
			} `tfsdk:"status_aggregation" json:"statusAggregation,omitempty"`
			StatusReflection *struct {
				LuaScript *string `tfsdk:"lua_script" json:"luaScript,omitempty"`
			} `tfsdk:"status_reflection" json:"statusReflection,omitempty"`
		} `tfsdk:"customizations" json:"customizations,omitempty"`
		Target *struct {
			ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
		} `tfsdk:"target" json:"target,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_config_karmada_io_resource_interpreter_customization_v1alpha1"
}

func (r *ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ResourceInterpreterCustomization describes the configuration of a specific resource for Karmada to get the structure. It has higher precedence than the default interpreter and the interpreter webhook.",
		MarkdownDescription: "ResourceInterpreterCustomization describes the configuration of a specific resource for Karmada to get the structure. It has higher precedence than the default interpreter and the interpreter webhook.",
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
				Description:         "Spec describes the configuration in detail.",
				MarkdownDescription: "Spec describes the configuration in detail.",
				Attributes: map[string]schema.Attribute{
					"customizations": schema.SingleNestedAttribute{
						Description:         "Customizations describe the interpretation rules.",
						MarkdownDescription: "Customizations describe the interpretation rules.",
						Attributes: map[string]schema.Attribute{
							"dependency_interpretation": schema.SingleNestedAttribute{
								Description:         "DependencyInterpretation describes the rules for Karmada to analyze the dependent resources. Karmada provides built-in rules for several standard Kubernetes types, see: https://karmada.io/docs/userguide/globalview/customizing-resource-interpreter/#interpretdependency If DependencyInterpretation is set, the built-in rules will be ignored.",
								MarkdownDescription: "DependencyInterpretation describes the rules for Karmada to analyze the dependent resources. Karmada provides built-in rules for several standard Kubernetes types, see: https://karmada.io/docs/userguide/globalview/customizing-resource-interpreter/#interpretdependency If DependencyInterpretation is set, the built-in rules will be ignored.",
								Attributes: map[string]schema.Attribute{
									"lua_script": schema.StringAttribute{
										Description:         "LuaScript holds the Lua script that is used to interpret the dependencies of a specific resource. The script should implement a function as follows: luaScript: > function GetDependencies(desiredObj) dependencies = {} if desiredObj.spec.serviceAccountName ~= '' and desiredObj.spec.serviceAccountName ~= 'default' then dependency = {} dependency.apiVersion = 'v1' dependency.kind = 'ServiceAccount' dependency.name = desiredObj.spec.serviceAccountName dependency.namespace = desiredObj.namespace dependencies[1] = {} dependencies[1] = dependency end return dependencies end  The content of the LuaScript needs to be a whole function including both declaration and implementation.  The parameters will be supplied by the system: - desiredObj: the object represents the configuration to be applied to the member cluster.  The returned value should be expressed by a slice of DependentObjectReference.",
										MarkdownDescription: "LuaScript holds the Lua script that is used to interpret the dependencies of a specific resource. The script should implement a function as follows: luaScript: > function GetDependencies(desiredObj) dependencies = {} if desiredObj.spec.serviceAccountName ~= '' and desiredObj.spec.serviceAccountName ~= 'default' then dependency = {} dependency.apiVersion = 'v1' dependency.kind = 'ServiceAccount' dependency.name = desiredObj.spec.serviceAccountName dependency.namespace = desiredObj.namespace dependencies[1] = {} dependencies[1] = dependency end return dependencies end  The content of the LuaScript needs to be a whole function including both declaration and implementation.  The parameters will be supplied by the system: - desiredObj: the object represents the configuration to be applied to the member cluster.  The returned value should be expressed by a slice of DependentObjectReference.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"health_interpretation": schema.SingleNestedAttribute{
								Description:         "HealthInterpretation describes the health assessment rules by which Karmada can assess the health state of the resource type.",
								MarkdownDescription: "HealthInterpretation describes the health assessment rules by which Karmada can assess the health state of the resource type.",
								Attributes: map[string]schema.Attribute{
									"lua_script": schema.StringAttribute{
										Description:         "LuaScript holds the Lua script that is used to assess the health state of a specific resource. The script should implement a function as follows: luaScript: > function InterpretHealth(observedObj) if observedObj.status.readyReplicas == observedObj.spec.replicas then return true end end  The content of the LuaScript needs to be a whole function including both declaration and implementation.  The parameters will be supplied by the system: - observedObj: the object represents the configuration that is observed from a specific member cluster.  The returned boolean value indicates the health status.",
										MarkdownDescription: "LuaScript holds the Lua script that is used to assess the health state of a specific resource. The script should implement a function as follows: luaScript: > function InterpretHealth(observedObj) if observedObj.status.readyReplicas == observedObj.spec.replicas then return true end end  The content of the LuaScript needs to be a whole function including both declaration and implementation.  The parameters will be supplied by the system: - observedObj: the object represents the configuration that is observed from a specific member cluster.  The returned boolean value indicates the health status.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"replica_resource": schema.SingleNestedAttribute{
								Description:         "ReplicaResource describes the rules for Karmada to discover the resource's replica as well as resource requirements. It would be useful for those CRD resources that declare workload types like Deployment. It is usually not needed for Kubernetes native resources(Deployment, Job) as Karmada knows how to discover info from them. But if it is set, the built-in discovery rules will be ignored.",
								MarkdownDescription: "ReplicaResource describes the rules for Karmada to discover the resource's replica as well as resource requirements. It would be useful for those CRD resources that declare workload types like Deployment. It is usually not needed for Kubernetes native resources(Deployment, Job) as Karmada knows how to discover info from them. But if it is set, the built-in discovery rules will be ignored.",
								Attributes: map[string]schema.Attribute{
									"lua_script": schema.StringAttribute{
										Description:         "LuaScript holds the Lua script that is used to discover the resource's replica as well as resource requirements  The script should implement a function as follows: luaScript: > function GetReplicas(desiredObj) replica = desiredObj.spec.replicas requirement = {} requirement.nodeClaim = {} requirement.nodeClaim.nodeSelector = desiredObj.spec.template.spec.nodeSelector requirement.nodeClaim.tolerations = desiredObj.spec.template.spec.tolerations requirement.resourceRequest = desiredObj.spec.template.spec.containers[1].resources.limits return replica, requirement end  The content of the LuaScript needs to be a whole function including both declaration and implementation.  The parameters will be supplied by the system: - desiredObj: the object represents the configuration to be applied to the member cluster.  The function expects two return values: - replica: the declared replica number - requirement: the resource required by each replica expressed with a ResourceBindingSpec.ReplicaRequirements. The returned values will be set into a ResourceBinding or ClusterResourceBinding.",
										MarkdownDescription: "LuaScript holds the Lua script that is used to discover the resource's replica as well as resource requirements  The script should implement a function as follows: luaScript: > function GetReplicas(desiredObj) replica = desiredObj.spec.replicas requirement = {} requirement.nodeClaim = {} requirement.nodeClaim.nodeSelector = desiredObj.spec.template.spec.nodeSelector requirement.nodeClaim.tolerations = desiredObj.spec.template.spec.tolerations requirement.resourceRequest = desiredObj.spec.template.spec.containers[1].resources.limits return replica, requirement end  The content of the LuaScript needs to be a whole function including both declaration and implementation.  The parameters will be supplied by the system: - desiredObj: the object represents the configuration to be applied to the member cluster.  The function expects two return values: - replica: the declared replica number - requirement: the resource required by each replica expressed with a ResourceBindingSpec.ReplicaRequirements. The returned values will be set into a ResourceBinding or ClusterResourceBinding.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"replica_revision": schema.SingleNestedAttribute{
								Description:         "ReplicaRevision describes the rules for Karmada to revise the resource's replica. It would be useful for those CRD resources that declare workload types like Deployment. It is usually not needed for Kubernetes native resources(Deployment, Job) as Karmada knows how to revise replicas for them. But if it is set, the built-in revision rules will be ignored.",
								MarkdownDescription: "ReplicaRevision describes the rules for Karmada to revise the resource's replica. It would be useful for those CRD resources that declare workload types like Deployment. It is usually not needed for Kubernetes native resources(Deployment, Job) as Karmada knows how to revise replicas for them. But if it is set, the built-in revision rules will be ignored.",
								Attributes: map[string]schema.Attribute{
									"lua_script": schema.StringAttribute{
										Description:         "LuaScript holds the Lua script that is used to revise replicas in the desired specification. The script should implement a function as follows: luaScript: > function ReviseReplica(desiredObj, desiredReplica) desiredObj.spec.replicas = desiredReplica return desiredObj end  The content of the LuaScript needs to be a whole function including both declaration and implementation.  The parameters will be supplied by the system: - desiredObj: the object represents the configuration to be applied to the member cluster. - desiredReplica: the replica number should be applied with.  The returned object should be a revised configuration which will be applied to member cluster eventually.",
										MarkdownDescription: "LuaScript holds the Lua script that is used to revise replicas in the desired specification. The script should implement a function as follows: luaScript: > function ReviseReplica(desiredObj, desiredReplica) desiredObj.spec.replicas = desiredReplica return desiredObj end  The content of the LuaScript needs to be a whole function including both declaration and implementation.  The parameters will be supplied by the system: - desiredObj: the object represents the configuration to be applied to the member cluster. - desiredReplica: the replica number should be applied with.  The returned object should be a revised configuration which will be applied to member cluster eventually.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"retention": schema.SingleNestedAttribute{
								Description:         "Retention describes the desired behavior that Karmada should react on the changes made by member cluster components. This avoids system running into a meaningless loop that Karmada resource controller and the member cluster component continually applying opposite values of a field. For example, the 'replicas' of Deployment might be changed by the HPA controller on member cluster. In this case, Karmada should retain the 'replicas' and not try to change it.",
								MarkdownDescription: "Retention describes the desired behavior that Karmada should react on the changes made by member cluster components. This avoids system running into a meaningless loop that Karmada resource controller and the member cluster component continually applying opposite values of a field. For example, the 'replicas' of Deployment might be changed by the HPA controller on member cluster. In this case, Karmada should retain the 'replicas' and not try to change it.",
								Attributes: map[string]schema.Attribute{
									"lua_script": schema.StringAttribute{
										Description:         "LuaScript holds the Lua script that is used to retain runtime values to the desired specification.  The script should implement a function as follows: luaScript: > function Retain(desiredObj, observedObj) desiredObj.spec.fieldFoo = observedObj.spec.fieldFoo return desiredObj end  The content of the LuaScript needs to be a whole function including both declaration and implementation.  The parameters will be supplied by the system: - desiredObj: the object represents the configuration to be applied to the member cluster. - observedObj: the object represents the configuration that is observed from a specific member cluster.  The returned object should be a retained configuration which will be applied to member cluster eventually.",
										MarkdownDescription: "LuaScript holds the Lua script that is used to retain runtime values to the desired specification.  The script should implement a function as follows: luaScript: > function Retain(desiredObj, observedObj) desiredObj.spec.fieldFoo = observedObj.spec.fieldFoo return desiredObj end  The content of the LuaScript needs to be a whole function including both declaration and implementation.  The parameters will be supplied by the system: - desiredObj: the object represents the configuration to be applied to the member cluster. - observedObj: the object represents the configuration that is observed from a specific member cluster.  The returned object should be a retained configuration which will be applied to member cluster eventually.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"status_aggregation": schema.SingleNestedAttribute{
								Description:         "StatusAggregation describes the rules for Karmada to aggregate status collected from member clusters to resource template. Karmada provides built-in rules for several standard Kubernetes types, see: https://karmada.io/docs/userguide/globalview/customizing-resource-interpreter/#aggregatestatus If StatusAggregation is set, the built-in rules will be ignored.",
								MarkdownDescription: "StatusAggregation describes the rules for Karmada to aggregate status collected from member clusters to resource template. Karmada provides built-in rules for several standard Kubernetes types, see: https://karmada.io/docs/userguide/globalview/customizing-resource-interpreter/#aggregatestatus If StatusAggregation is set, the built-in rules will be ignored.",
								Attributes: map[string]schema.Attribute{
									"lua_script": schema.StringAttribute{
										Description:         "LuaScript holds the Lua script that is used to aggregate decentralized statuses to the desired specification. The script should implement a function as follows: luaScript: > function AggregateStatus(desiredObj, statusItems) for i = 1, #statusItems do desiredObj.status.readyReplicas = desiredObj.status.readyReplicas + items[i].readyReplicas end return desiredObj end  The content of the LuaScript needs to be a whole function including both declaration and implementation.  The parameters will be supplied by the system: - desiredObj: the object represents a resource template. - statusItems: the slice of status expressed with AggregatedStatusItem.  The returned object should be a whole object with status aggregated.",
										MarkdownDescription: "LuaScript holds the Lua script that is used to aggregate decentralized statuses to the desired specification. The script should implement a function as follows: luaScript: > function AggregateStatus(desiredObj, statusItems) for i = 1, #statusItems do desiredObj.status.readyReplicas = desiredObj.status.readyReplicas + items[i].readyReplicas end return desiredObj end  The content of the LuaScript needs to be a whole function including both declaration and implementation.  The parameters will be supplied by the system: - desiredObj: the object represents a resource template. - statusItems: the slice of status expressed with AggregatedStatusItem.  The returned object should be a whole object with status aggregated.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"status_reflection": schema.SingleNestedAttribute{
								Description:         "StatusReflection describes the rules for Karmada to pick the resource's status. Karmada provides built-in rules for several standard Kubernetes types, see: https://karmada.io/docs/userguide/globalview/customizing-resource-interpreter/#interpretstatus If StatusReflection is set, the built-in rules will be ignored.",
								MarkdownDescription: "StatusReflection describes the rules for Karmada to pick the resource's status. Karmada provides built-in rules for several standard Kubernetes types, see: https://karmada.io/docs/userguide/globalview/customizing-resource-interpreter/#interpretstatus If StatusReflection is set, the built-in rules will be ignored.",
								Attributes: map[string]schema.Attribute{
									"lua_script": schema.StringAttribute{
										Description:         "LuaScript holds the Lua script that is used to get the status from the observed specification. The script should implement a function as follows: luaScript: > function ReflectStatus(observedObj) status = {} status.readyReplicas = observedObj.status.observedObj return status end  The content of the LuaScript needs to be a whole function including both declaration and implementation.  The parameters will be supplied by the system: - observedObj: the object represents the configuration that is observed from a specific member cluster.  The returned status could be the whole status or part of it and will be set into both Work and ResourceBinding(ClusterResourceBinding).",
										MarkdownDescription: "LuaScript holds the Lua script that is used to get the status from the observed specification. The script should implement a function as follows: luaScript: > function ReflectStatus(observedObj) status = {} status.readyReplicas = observedObj.status.observedObj return status end  The content of the LuaScript needs to be a whole function including both declaration and implementation.  The parameters will be supplied by the system: - observedObj: the object represents the configuration that is observed from a specific member cluster.  The returned status could be the whole status or part of it and will be set into both Work and ResourceBinding(ClusterResourceBinding).",
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

					"target": schema.SingleNestedAttribute{
						Description:         "CustomizationTarget represents the resource type that the customization applies to.",
						MarkdownDescription: "CustomizationTarget represents the resource type that the customization applies to.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "APIVersion represents the API version of the target resource.",
								MarkdownDescription: "APIVersion represents the API version of the target resource.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind represents the Kind of target resources.",
								MarkdownDescription: "Kind represents the Kind of target resources.",
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

func (r *ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_config_karmada_io_resource_interpreter_customization_v1alpha1")

	var data ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "config.karmada.io", Version: "v1alpha1", Resource: "ResourceInterpreterCustomization"}).
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

	var readResponse ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1DataSourceData
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
	data.ApiVersion = pointer.String("config.karmada.io/v1alpha1")
	data.Kind = pointer.String("ResourceInterpreterCustomization")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
