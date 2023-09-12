/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package config_karmada_io_v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	k8sTypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"time"
)

var (
	_ resource.Resource                = &ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1Resource{}
	_ resource.ResourceWithConfigure   = &ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1Resource{}
	_ resource.ResourceWithImportState = &ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1Resource{}
)

func NewConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1Resource() resource.Resource {
	return &ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1Resource{}
}

type ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1ResourceData struct {
	ID                  types.String `tfsdk:"id" json:"-"`
	ForceConflicts      types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager        types.String `tfsdk:"field_manager" json:"-"`
	DeletionPropagation types.String `tfsdk:"deletion_propagation" json:"-"`
	WaitForUpsert       types.List   `tfsdk:"wait_for_upsert" json:"-"`
	WaitForDelete       types.Object `tfsdk:"wait_for_delete" json:"-"`

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

func (r *ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_config_karmada_io_resource_interpreter_customization_v1alpha1"
}

func (r *ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
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

			"force_conflicts": schema.BoolAttribute{
				Description:         "If 'true', server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "If `true`, server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"field_manager": schema.StringAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},

			"deletion_propagation": schema.StringAttribute{
				Description:         "Decides if a deletion will propagate to the dependents of the object, and how the garbage collector will handle the propagation.",
				MarkdownDescription: "Decides if a deletion will propagate to the dependents of the object, and how the garbage collector will handle the propagation.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("Orphan", "Background", "Foreground"),
				},
			},

			"wait_for_upsert": schema.ListNestedAttribute{
				Description:         "Wait for specific conditions after create/update of resources.",
				MarkdownDescription: "Wait for specific conditions after create/update of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"jsonpath": schema.StringAttribute{
							Description:         "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							MarkdownDescription: "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},
						"value": schema.StringAttribute{
							Description:         "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							MarkdownDescription: "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							Required:            false,
							Optional:            true,
							Computed:            true,
						},
						"timeout": schema.Int64Attribute{
							Description:         "The number of seconds to wait before giving up. Zero means check once and don't wait.",
							MarkdownDescription: "The number of seconds to wait before giving up. Zero means check once and don't wait.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(30),
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
						"poll_interval": schema.Int64Attribute{
							Description:         "The number of seconds to wait before checking again.",
							MarkdownDescription: "The number of seconds to wait before checking again.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(5),
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
					},
				},
			},

			"wait_for_delete": schema.SingleNestedAttribute{
				Description:         "Wait for deletion of resources.",
				MarkdownDescription: "Wait for deletion of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"timeout": schema.Int64Attribute{
						Description:         "The number of seconds to wait before giving up. Zero means check once and don't wait.",
						MarkdownDescription: "The number of seconds to wait before giving up. Zero means check once and don't wait.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             int64default.StaticInt64(30),
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
					"poll_interval": schema.Int64Attribute{
						Description:         "The number of seconds to wait before checking again.",
						MarkdownDescription: "The number of seconds to wait before checking again.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             int64default.StaticInt64(5),
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
				},
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            true,
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
						Computed:            true,
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
										Description:         "LuaScript holds the Lua script that is used to interpret the dependencies of a specific resource. The script should implement a function as follows: luaScript: > function GetDependencies(desiredObj) dependencies = {} if desiredObj.spec.serviceAccountName ~= nil and desiredObj.spec.serviceAccountName ~= 'default' then dependency = {} dependency.apiVersion = 'v1' dependency.kind = 'ServiceAccount' dependency.name = desiredObj.spec.serviceAccountName dependency.namespace = desiredObj.namespace dependencies[1] = {} dependencies[1] = dependency end return dependencies end  The content of the LuaScript needs to be a whole function including both declaration and implementation.  The parameters will be supplied by the system: - desiredObj: the object represents the configuration to be applied to the member cluster.  The returned value should be expressed by a slice of DependentObjectReference.",
										MarkdownDescription: "LuaScript holds the Lua script that is used to interpret the dependencies of a specific resource. The script should implement a function as follows: luaScript: > function GetDependencies(desiredObj) dependencies = {} if desiredObj.spec.serviceAccountName ~= nil and desiredObj.spec.serviceAccountName ~= 'default' then dependency = {} dependency.apiVersion = 'v1' dependency.kind = 'ServiceAccount' dependency.name = desiredObj.spec.serviceAccountName dependency.namespace = desiredObj.namespace dependencies[1] = {} dependencies[1] = dependency end return dependencies end  The content of the LuaScript needs to be a whole function including both declaration and implementation.  The parameters will be supplied by the system: - desiredObj: the object represents the configuration to be applied to the member cluster.  The returned value should be expressed by a slice of DependentObjectReference.",
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
										Description:         "LuaScript holds the Lua script that is used to assess the health state of a specific resource. The script should implement a function as follows: luaScript: > function InterpretHealth(observedObj) if observedObj.status.readyReplicas == observedObj.spec.replicas then return true end end  The content of the LuaScript needs to be a whole function including both declaration and implementation.  The parameters will be supplied by the system: - observedObj: the object represents the configuration that is observed from a specific member cluster.  The returned boolean value indicates the health status.",
										MarkdownDescription: "LuaScript holds the Lua script that is used to assess the health state of a specific resource. The script should implement a function as follows: luaScript: > function InterpretHealth(observedObj) if observedObj.status.readyReplicas == observedObj.spec.replicas then return true end end  The content of the LuaScript needs to be a whole function including both declaration and implementation.  The parameters will be supplied by the system: - observedObj: the object represents the configuration that is observed from a specific member cluster.  The returned boolean value indicates the health status.",
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
										Description:         "LuaScript holds the Lua script that is used to discover the resource's replica as well as resource requirements  The script should implement a function as follows: luaScript: > function GetReplicas(desiredObj) replica = desiredObj.spec.replicas requirement = {} requirement.nodeClaim = {} requirement.nodeClaim.nodeSelector = desiredObj.spec.template.spec.nodeSelector requirement.nodeClaim.tolerations = desiredObj.spec.template.spec.tolerations requirement.resourceRequest = desiredObj.spec.template.spec.containers[1].resources.limits return replica, requirement end  The content of the LuaScript needs to be a whole function including both declaration and implementation.  The parameters will be supplied by the system: - desiredObj: the object represents the configuration to be applied to the member cluster.  The function expects two return values: - replica: the declared replica number - requirement: the resource required by each replica expressed with a ResourceBindingSpec.ReplicaRequirements. The returned values will be set into a ResourceBinding or ClusterResourceBinding.",
										MarkdownDescription: "LuaScript holds the Lua script that is used to discover the resource's replica as well as resource requirements  The script should implement a function as follows: luaScript: > function GetReplicas(desiredObj) replica = desiredObj.spec.replicas requirement = {} requirement.nodeClaim = {} requirement.nodeClaim.nodeSelector = desiredObj.spec.template.spec.nodeSelector requirement.nodeClaim.tolerations = desiredObj.spec.template.spec.tolerations requirement.resourceRequest = desiredObj.spec.template.spec.containers[1].resources.limits return replica, requirement end  The content of the LuaScript needs to be a whole function including both declaration and implementation.  The parameters will be supplied by the system: - desiredObj: the object represents the configuration to be applied to the member cluster.  The function expects two return values: - replica: the declared replica number - requirement: the resource required by each replica expressed with a ResourceBindingSpec.ReplicaRequirements. The returned values will be set into a ResourceBinding or ClusterResourceBinding.",
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
										Description:         "LuaScript holds the Lua script that is used to revise replicas in the desired specification. The script should implement a function as follows: luaScript: > function ReviseReplica(desiredObj, desiredReplica) desiredObj.spec.replicas = desiredReplica return desiredObj end  The content of the LuaScript needs to be a whole function including both declaration and implementation.  The parameters will be supplied by the system: - desiredObj: the object represents the configuration to be applied to the member cluster. - desiredReplica: the replica number should be applied with.  The returned object should be a revised configuration which will be applied to member cluster eventually.",
										MarkdownDescription: "LuaScript holds the Lua script that is used to revise replicas in the desired specification. The script should implement a function as follows: luaScript: > function ReviseReplica(desiredObj, desiredReplica) desiredObj.spec.replicas = desiredReplica return desiredObj end  The content of the LuaScript needs to be a whole function including both declaration and implementation.  The parameters will be supplied by the system: - desiredObj: the object represents the configuration to be applied to the member cluster. - desiredReplica: the replica number should be applied with.  The returned object should be a revised configuration which will be applied to member cluster eventually.",
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
										Description:         "LuaScript holds the Lua script that is used to retain runtime values to the desired specification.  The script should implement a function as follows: luaScript: > function Retain(desiredObj, observedObj) desiredObj.spec.fieldFoo = observedObj.spec.fieldFoo return desiredObj end  The content of the LuaScript needs to be a whole function including both declaration and implementation.  The parameters will be supplied by the system: - desiredObj: the object represents the configuration to be applied to the member cluster. - observedObj: the object represents the configuration that is observed from a specific member cluster.  The returned object should be a retained configuration which will be applied to member cluster eventually.",
										MarkdownDescription: "LuaScript holds the Lua script that is used to retain runtime values to the desired specification.  The script should implement a function as follows: luaScript: > function Retain(desiredObj, observedObj) desiredObj.spec.fieldFoo = observedObj.spec.fieldFoo return desiredObj end  The content of the LuaScript needs to be a whole function including both declaration and implementation.  The parameters will be supplied by the system: - desiredObj: the object represents the configuration to be applied to the member cluster. - observedObj: the object represents the configuration that is observed from a specific member cluster.  The returned object should be a retained configuration which will be applied to member cluster eventually.",
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
										Description:         "LuaScript holds the Lua script that is used to aggregate decentralized statuses to the desired specification. The script should implement a function as follows: luaScript: > function AggregateStatus(desiredObj, statusItems) for i = 1, #statusItems do desiredObj.status.readyReplicas = desiredObj.status.readyReplicas + items[i].readyReplicas end return desiredObj end  The content of the LuaScript needs to be a whole function including both declaration and implementation.  The parameters will be supplied by the system: - desiredObj: the object represents a resource template. - statusItems: the slice of status expressed with AggregatedStatusItem.  The returned object should be a whole object with status aggregated.",
										MarkdownDescription: "LuaScript holds the Lua script that is used to aggregate decentralized statuses to the desired specification. The script should implement a function as follows: luaScript: > function AggregateStatus(desiredObj, statusItems) for i = 1, #statusItems do desiredObj.status.readyReplicas = desiredObj.status.readyReplicas + items[i].readyReplicas end return desiredObj end  The content of the LuaScript needs to be a whole function including both declaration and implementation.  The parameters will be supplied by the system: - desiredObj: the object represents a resource template. - statusItems: the slice of status expressed with AggregatedStatusItem.  The returned object should be a whole object with status aggregated.",
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
										Description:         "LuaScript holds the Lua script that is used to get the status from the observed specification. The script should implement a function as follows: luaScript: > function ReflectStatus(observedObj) status = {} status.readyReplicas = observedObj.status.observedObj return status end  The content of the LuaScript needs to be a whole function including both declaration and implementation.  The parameters will be supplied by the system: - observedObj: the object represents the configuration that is observed from a specific member cluster.  The returned status could be the whole status or part of it and will be set into both Work and ResourceBinding(ClusterResourceBinding).",
										MarkdownDescription: "LuaScript holds the Lua script that is used to get the status from the observed specification. The script should implement a function as follows: luaScript: > function ReflectStatus(observedObj) status = {} status.readyReplicas = observedObj.status.observedObj return status end  The content of the LuaScript needs to be a whole function including both declaration and implementation.  The parameters will be supplied by the system: - observedObj: the object represents the configuration that is observed from a specific member cluster.  The returned status could be the whole status or part of it and will be set into both Work and ResourceBinding(ClusterResourceBinding).",
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

func (r *ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedResourceDataError(request.ProviderData))
	}
}

func (r *ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_config_karmada_io_resource_interpreter_customization_v1alpha1")

	var model ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("config.karmada.io/v1alpha1")
	model.Kind = pointer.String("ResourceInterpreterCustomization")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "config.karmada.io", Version: "v1alpha1", Resource: "resourceinterpretercustomizations"}).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec
	if model.ForceConflicts.IsUnknown() {
		model.ForceConflicts = types.BoolNull()
	}
	if model.FieldManager.IsUnknown() {
		model.FieldManager = types.StringNull()
	}
	if model.DeletionPropagation.IsUnknown() {
		model.DeletionPropagation = types.StringNull()
	}
	if model.WaitForUpsert.IsUnknown() {
		model.WaitForUpsert = types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"jsonpath":      types.StringType,
				"value":         types.StringType,
				"timeout":       types.Int64Type,
				"poll_interval": types.Int64Type,
			},
		})
	}
	if model.WaitForDelete.IsUnknown() {
		model.WaitForDelete = types.ObjectNull(map[string]attr.Type{
			"timeout":       types.Int64Type,
			"poll_interval": types.Int64Type,
		})
	}

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_config_karmada_io_resource_interpreter_customization_v1alpha1")

	var data ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "config.karmada.io", Version: "v1alpha1", Resource: "resourceinterpretercustomizations"}).
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

	var readResponse ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec
	if data.ForceConflicts.IsUnknown() {
		data.ForceConflicts = types.BoolNull()
	}
	if data.FieldManager.IsUnknown() {
		data.FieldManager = types.StringNull()
	}
	if data.DeletionPropagation.IsUnknown() {
		data.DeletionPropagation = types.StringNull()
	}
	if data.WaitForUpsert.IsUnknown() {
		data.WaitForUpsert = types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"jsonpath":      types.StringType,
				"value":         types.StringType,
				"timeout":       types.Int64Type,
				"poll_interval": types.Int64Type,
			},
		})
	}
	if data.WaitForDelete.IsUnknown() {
		data.WaitForDelete = types.ObjectNull(map[string]attr.Type{
			"timeout":       types.Int64Type,
			"poll_interval": types.Int64Type,
		})
	}

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_config_karmada_io_resource_interpreter_customization_v1alpha1")

	var model ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("config.karmada.io/v1alpha1")
	model.Kind = pointer.String("ResourceInterpreterCustomization")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "config.karmada.io", Version: "v1alpha1", Resource: "resourceinterpretercustomizations"}).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_config_karmada_io_resource_interpreter_customization_v1alpha1")

	var data ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	deleteOptions := meta.DeleteOptions{}
	if !data.DeletionPropagation.IsNull() && !data.DeletionPropagation.IsUnknown() {
		deleteOptions.PropagationPolicy = utilities.MapDeletionPropagation(data.DeletionPropagation.ValueString())
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "config.karmada.io", Version: "v1alpha1", Resource: "resourceinterpretercustomizations"}).
		Delete(ctx, data.Metadata.Name, deleteOptions)
	if utilities.IsDeletionError(err) {
		response.Diagnostics.Append(utilities.DeleteError(err))
		return
	}

	if !data.WaitForDelete.IsNull() && !data.WaitForDelete.IsUnknown() {
		timeout := utilities.DetermineTimeout(data.WaitForDelete.Attributes())
		pollInterval := utilities.DeterminePollInterval(data.WaitForDelete.Attributes())

		startTime := time.Now()
		for {
			_, err := r.kubernetesClient.
				Resource(k8sSchema.GroupVersionResource{Group: "config.karmada.io", Version: "v1alpha1", Resource: "resourceinterpretercustomizations"}).
				Get(ctx, data.Metadata.Name, meta.GetOptions{})
			if utilities.IsNotFound(err) || timeout.Milliseconds() == 0 {
				break
			}
			if time.Now().After(startTime.Add(timeout)) {
				response.Diagnostics.Append(utilities.WaitTimeoutExceeded())
				return
			}
			time.Sleep(pollInterval)
		}
	}
}

func (r *ConfigKarmadaIoResourceInterpreterCustomizationV1Alpha1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	if request.ID == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'name' Got: '%q'", request.ID),
		)
		return
	}
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	resource.ImportStatePassthroughID(ctx, path.Root("metadata").AtName("name"), request, response)
}
