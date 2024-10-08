/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package config_karmada_io_v1alpha1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1Manifest{}
)

func NewConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1Manifest() datasource.DataSource {
	return &ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1Manifest{}
}

type ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1Manifest struct{}

type ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

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

func (r *ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_config_karmada_io_resource_interpreter_customization_v1alpha1_manifest"
}

func (r *ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ResourceInterpreterCustomization describes the configuration of a specific resource for Karmada to get the structure. It has higher precedence than the default interpreter and the interpreter webhook.",
		MarkdownDescription: "ResourceInterpreterCustomization describes the configuration of a specific resource for Karmada to get the structure. It has higher precedence than the default interpreter and the interpreter webhook.",
		Attributes: map[string]schema.Attribute{
			"yaml": schema.StringAttribute{
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
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
						Optional:            true,
						Computed:            false,
						Validators: []validator.Map{
							validators.LabelValidator(),
						},
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Map{
							validators.AnnotationValidator(),
						},
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
										Description:         "LuaScript holds the Lua script that is used to interpret the dependencies of a specific resource. The script should implement a function as follows: ''' luaScript: > function GetDependencies(desiredObj) dependencies = {} serviceAccountName = desiredObj.spec.template.spec.serviceAccountName if serviceAccountName ~= nil and serviceAccountName ~= 'default' then dependency = {} dependency.apiVersion = 'v1' dependency.kind = 'ServiceAccount' dependency.name = serviceAccountName dependency.namespace = desiredObj.metadata.namespace dependencies[1] = dependency end return dependencies end ''' The content of the LuaScript needs to be a whole function including both declaration and implementation. The parameters will be supplied by the system: - desiredObj: the object represents the configuration to be applied to the member cluster. The returned value should be expressed by a slice of DependentObjectReference.",
										MarkdownDescription: "LuaScript holds the Lua script that is used to interpret the dependencies of a specific resource. The script should implement a function as follows: ''' luaScript: > function GetDependencies(desiredObj) dependencies = {} serviceAccountName = desiredObj.spec.template.spec.serviceAccountName if serviceAccountName ~= nil and serviceAccountName ~= 'default' then dependency = {} dependency.apiVersion = 'v1' dependency.kind = 'ServiceAccount' dependency.name = serviceAccountName dependency.namespace = desiredObj.metadata.namespace dependencies[1] = dependency end return dependencies end ''' The content of the LuaScript needs to be a whole function including both declaration and implementation. The parameters will be supplied by the system: - desiredObj: the object represents the configuration to be applied to the member cluster. The returned value should be expressed by a slice of DependentObjectReference.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"health_interpretation": schema.SingleNestedAttribute{
								Description:         "HealthInterpretation describes the health assessment rules by which Karmada can assess the health state of the resource type.",
								MarkdownDescription: "HealthInterpretation describes the health assessment rules by which Karmada can assess the health state of the resource type.",
								Attributes: map[string]schema.Attribute{
									"lua_script": schema.StringAttribute{
										Description:         "LuaScript holds the Lua script that is used to assess the health state of a specific resource. The script should implement a function as follows: ''' luaScript: > function InterpretHealth(observedObj) if observedObj.status.readyReplicas == observedObj.spec.replicas then return true end end ''' The content of the LuaScript needs to be a whole function including both declaration and implementation. The parameters will be supplied by the system: - observedObj: the object represents the configuration that is observed from a specific member cluster. The returned boolean value indicates the health status.",
										MarkdownDescription: "LuaScript holds the Lua script that is used to assess the health state of a specific resource. The script should implement a function as follows: ''' luaScript: > function InterpretHealth(observedObj) if observedObj.status.readyReplicas == observedObj.spec.replicas then return true end end ''' The content of the LuaScript needs to be a whole function including both declaration and implementation. The parameters will be supplied by the system: - observedObj: the object represents the configuration that is observed from a specific member cluster. The returned boolean value indicates the health status.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"replica_resource": schema.SingleNestedAttribute{
								Description:         "ReplicaResource describes the rules for Karmada to discover the resource's replica as well as resource requirements. It would be useful for those CRD resources that declare workload types like Deployment. It is usually not needed for Kubernetes native resources(Deployment, Job) as Karmada knows how to discover info from them. But if it is set, the built-in discovery rules will be ignored.",
								MarkdownDescription: "ReplicaResource describes the rules for Karmada to discover the resource's replica as well as resource requirements. It would be useful for those CRD resources that declare workload types like Deployment. It is usually not needed for Kubernetes native resources(Deployment, Job) as Karmada knows how to discover info from them. But if it is set, the built-in discovery rules will be ignored.",
								Attributes: map[string]schema.Attribute{
									"lua_script": schema.StringAttribute{
										Description:         "LuaScript holds the Lua script that is used to discover the resource's replica as well as resource requirements The script should implement a function as follows: ''' luaScript: > function GetReplicas(desiredObj) replica = desiredObj.spec.replicas requirement = {} requirement.nodeClaim = {} requirement.nodeClaim.nodeSelector = desiredObj.spec.template.spec.nodeSelector requirement.nodeClaim.tolerations = desiredObj.spec.template.spec.tolerations requirement.resourceRequest = desiredObj.spec.template.spec.containers[1].resources.limits return replica, requirement end ''' The content of the LuaScript needs to be a whole function including both declaration and implementation. The parameters will be supplied by the system: - desiredObj: the object represents the configuration to be applied to the member cluster. The function expects two return values: - replica: the declared replica number - requirement: the resource required by each replica expressed with a ResourceBindingSpec.ReplicaRequirements. The returned values will be set into a ResourceBinding or ClusterResourceBinding.",
										MarkdownDescription: "LuaScript holds the Lua script that is used to discover the resource's replica as well as resource requirements The script should implement a function as follows: ''' luaScript: > function GetReplicas(desiredObj) replica = desiredObj.spec.replicas requirement = {} requirement.nodeClaim = {} requirement.nodeClaim.nodeSelector = desiredObj.spec.template.spec.nodeSelector requirement.nodeClaim.tolerations = desiredObj.spec.template.spec.tolerations requirement.resourceRequest = desiredObj.spec.template.spec.containers[1].resources.limits return replica, requirement end ''' The content of the LuaScript needs to be a whole function including both declaration and implementation. The parameters will be supplied by the system: - desiredObj: the object represents the configuration to be applied to the member cluster. The function expects two return values: - replica: the declared replica number - requirement: the resource required by each replica expressed with a ResourceBindingSpec.ReplicaRequirements. The returned values will be set into a ResourceBinding or ClusterResourceBinding.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"replica_revision": schema.SingleNestedAttribute{
								Description:         "ReplicaRevision describes the rules for Karmada to revise the resource's replica. It would be useful for those CRD resources that declare workload types like Deployment. It is usually not needed for Kubernetes native resources(Deployment, Job) as Karmada knows how to revise replicas for them. But if it is set, the built-in revision rules will be ignored.",
								MarkdownDescription: "ReplicaRevision describes the rules for Karmada to revise the resource's replica. It would be useful for those CRD resources that declare workload types like Deployment. It is usually not needed for Kubernetes native resources(Deployment, Job) as Karmada knows how to revise replicas for them. But if it is set, the built-in revision rules will be ignored.",
								Attributes: map[string]schema.Attribute{
									"lua_script": schema.StringAttribute{
										Description:         "LuaScript holds the Lua script that is used to revise replicas in the desired specification. The script should implement a function as follows: ''' luaScript: > function ReviseReplica(desiredObj, desiredReplica) desiredObj.spec.replicas = desiredReplica return desiredObj end ''' The content of the LuaScript needs to be a whole function including both declaration and implementation. The parameters will be supplied by the system: - desiredObj: the object represents the configuration to be applied to the member cluster. - desiredReplica: the replica number should be applied with. The returned object should be a revised configuration which will be applied to member cluster eventually.",
										MarkdownDescription: "LuaScript holds the Lua script that is used to revise replicas in the desired specification. The script should implement a function as follows: ''' luaScript: > function ReviseReplica(desiredObj, desiredReplica) desiredObj.spec.replicas = desiredReplica return desiredObj end ''' The content of the LuaScript needs to be a whole function including both declaration and implementation. The parameters will be supplied by the system: - desiredObj: the object represents the configuration to be applied to the member cluster. - desiredReplica: the replica number should be applied with. The returned object should be a revised configuration which will be applied to member cluster eventually.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"retention": schema.SingleNestedAttribute{
								Description:         "Retention describes the desired behavior that Karmada should react on the changes made by member cluster components. This avoids system running into a meaningless loop that Karmada resource controller and the member cluster component continually applying opposite values of a field. For example, the 'replicas' of Deployment might be changed by the HPA controller on member cluster. In this case, Karmada should retain the 'replicas' and not try to change it.",
								MarkdownDescription: "Retention describes the desired behavior that Karmada should react on the changes made by member cluster components. This avoids system running into a meaningless loop that Karmada resource controller and the member cluster component continually applying opposite values of a field. For example, the 'replicas' of Deployment might be changed by the HPA controller on member cluster. In this case, Karmada should retain the 'replicas' and not try to change it.",
								Attributes: map[string]schema.Attribute{
									"lua_script": schema.StringAttribute{
										Description:         "LuaScript holds the Lua script that is used to retain runtime values to the desired specification. The script should implement a function as follows: ''' luaScript: > function Retain(desiredObj, observedObj) desiredObj.spec.fieldFoo = observedObj.spec.fieldFoo return desiredObj end ''' The content of the LuaScript needs to be a whole function including both declaration and implementation. The parameters will be supplied by the system: - desiredObj: the object represents the configuration to be applied to the member cluster. - observedObj: the object represents the configuration that is observed from a specific member cluster. The returned object should be a retained configuration which will be applied to member cluster eventually.",
										MarkdownDescription: "LuaScript holds the Lua script that is used to retain runtime values to the desired specification. The script should implement a function as follows: ''' luaScript: > function Retain(desiredObj, observedObj) desiredObj.spec.fieldFoo = observedObj.spec.fieldFoo return desiredObj end ''' The content of the LuaScript needs to be a whole function including both declaration and implementation. The parameters will be supplied by the system: - desiredObj: the object represents the configuration to be applied to the member cluster. - observedObj: the object represents the configuration that is observed from a specific member cluster. The returned object should be a retained configuration which will be applied to member cluster eventually.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"status_aggregation": schema.SingleNestedAttribute{
								Description:         "StatusAggregation describes the rules for Karmada to aggregate status collected from member clusters to resource template. Karmada provides built-in rules for several standard Kubernetes types, see: https://karmada.io/docs/userguide/globalview/customizing-resource-interpreter/#aggregatestatus If StatusAggregation is set, the built-in rules will be ignored.",
								MarkdownDescription: "StatusAggregation describes the rules for Karmada to aggregate status collected from member clusters to resource template. Karmada provides built-in rules for several standard Kubernetes types, see: https://karmada.io/docs/userguide/globalview/customizing-resource-interpreter/#aggregatestatus If StatusAggregation is set, the built-in rules will be ignored.",
								Attributes: map[string]schema.Attribute{
									"lua_script": schema.StringAttribute{
										Description:         "LuaScript holds the Lua script that is used to aggregate decentralized statuses to the desired specification. The script should implement a function as follows: ''' luaScript: > function AggregateStatus(desiredObj, statusItems) for i = 1, #statusItems do desiredObj.status.readyReplicas = desiredObj.status.readyReplicas + items[i].readyReplicas end return desiredObj end ''' The content of the LuaScript needs to be a whole function including both declaration and implementation. The parameters will be supplied by the system: - desiredObj: the object represents a resource template. - statusItems: the slice of status expressed with AggregatedStatusItem. The returned object should be a whole object with status aggregated.",
										MarkdownDescription: "LuaScript holds the Lua script that is used to aggregate decentralized statuses to the desired specification. The script should implement a function as follows: ''' luaScript: > function AggregateStatus(desiredObj, statusItems) for i = 1, #statusItems do desiredObj.status.readyReplicas = desiredObj.status.readyReplicas + items[i].readyReplicas end return desiredObj end ''' The content of the LuaScript needs to be a whole function including both declaration and implementation. The parameters will be supplied by the system: - desiredObj: the object represents a resource template. - statusItems: the slice of status expressed with AggregatedStatusItem. The returned object should be a whole object with status aggregated.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"status_reflection": schema.SingleNestedAttribute{
								Description:         "StatusReflection describes the rules for Karmada to pick the resource's status. Karmada provides built-in rules for several standard Kubernetes types, see: https://karmada.io/docs/userguide/globalview/customizing-resource-interpreter/#interpretstatus If StatusReflection is set, the built-in rules will be ignored.",
								MarkdownDescription: "StatusReflection describes the rules for Karmada to pick the resource's status. Karmada provides built-in rules for several standard Kubernetes types, see: https://karmada.io/docs/userguide/globalview/customizing-resource-interpreter/#interpretstatus If StatusReflection is set, the built-in rules will be ignored.",
								Attributes: map[string]schema.Attribute{
									"lua_script": schema.StringAttribute{
										Description:         "LuaScript holds the Lua script that is used to get the status from the observed specification. The script should implement a function as follows: ''' luaScript: > function ReflectStatus(observedObj) status = {} status.readyReplicas = observedObj.status.observedObj return status end ''' The content of the LuaScript needs to be a whole function including both declaration and implementation. The parameters will be supplied by the system: - observedObj: the object represents the configuration that is observed from a specific member cluster. The returned status could be the whole status or part of it and will be set into both Work and ResourceBinding(ClusterResourceBinding).",
										MarkdownDescription: "LuaScript holds the Lua script that is used to get the status from the observed specification. The script should implement a function as follows: ''' luaScript: > function ReflectStatus(observedObj) status = {} status.readyReplicas = observedObj.status.observedObj return status end ''' The content of the LuaScript needs to be a whole function including both declaration and implementation. The parameters will be supplied by the system: - observedObj: the object represents the configuration that is observed from a specific member cluster. The returned status could be the whole status or part of it and will be set into both Work and ResourceBinding(ClusterResourceBinding).",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"target": schema.SingleNestedAttribute{
						Description:         "CustomizationTarget represents the resource type that the customization applies to.",
						MarkdownDescription: "CustomizationTarget represents the resource type that the customization applies to.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "APIVersion represents the API version of the target resource.",
								MarkdownDescription: "APIVersion represents the API version of the target resource.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind represents the Kind of target resources.",
								MarkdownDescription: "Kind represents the Kind of target resources.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_config_karmada_io_resource_interpreter_customization_v1alpha1_manifest")

	var model ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("config.karmada.io/v1alpha1")
	model.Kind = pointer.String("ResourceInterpreterCustomization")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
