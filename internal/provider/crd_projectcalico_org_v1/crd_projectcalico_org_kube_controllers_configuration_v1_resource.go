/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package crd_projectcalico_org_v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
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
	_ resource.Resource                = &CrdProjectcalicoOrgKubeControllersConfigurationV1Resource{}
	_ resource.ResourceWithConfigure   = &CrdProjectcalicoOrgKubeControllersConfigurationV1Resource{}
	_ resource.ResourceWithImportState = &CrdProjectcalicoOrgKubeControllersConfigurationV1Resource{}
)

func NewCrdProjectcalicoOrgKubeControllersConfigurationV1Resource() resource.Resource {
	return &CrdProjectcalicoOrgKubeControllersConfigurationV1Resource{}
}

type CrdProjectcalicoOrgKubeControllersConfigurationV1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type CrdProjectcalicoOrgKubeControllersConfigurationV1ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitForUpsert  types.List   `tfsdk:"wait_for_upsert" json:"-"`
	WaitForDelete  types.Object `tfsdk:"wait_for_delete" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Controllers *struct {
			Namespace *struct {
				ReconcilerPeriod *string `tfsdk:"reconciler_period" json:"reconcilerPeriod,omitempty"`
			} `tfsdk:"namespace" json:"namespace,omitempty"`
			Node *struct {
				HostEndpoint *struct {
					AutoCreate *string `tfsdk:"auto_create" json:"autoCreate,omitempty"`
				} `tfsdk:"host_endpoint" json:"hostEndpoint,omitempty"`
				LeakGracePeriod  *string `tfsdk:"leak_grace_period" json:"leakGracePeriod,omitempty"`
				ReconcilerPeriod *string `tfsdk:"reconciler_period" json:"reconcilerPeriod,omitempty"`
				SyncLabels       *string `tfsdk:"sync_labels" json:"syncLabels,omitempty"`
			} `tfsdk:"node" json:"node,omitempty"`
			Policy *struct {
				ReconcilerPeriod *string `tfsdk:"reconciler_period" json:"reconcilerPeriod,omitempty"`
			} `tfsdk:"policy" json:"policy,omitempty"`
			ServiceAccount *struct {
				ReconcilerPeriod *string `tfsdk:"reconciler_period" json:"reconcilerPeriod,omitempty"`
			} `tfsdk:"service_account" json:"serviceAccount,omitempty"`
			WorkloadEndpoint *struct {
				ReconcilerPeriod *string `tfsdk:"reconciler_period" json:"reconcilerPeriod,omitempty"`
			} `tfsdk:"workload_endpoint" json:"workloadEndpoint,omitempty"`
		} `tfsdk:"controllers" json:"controllers,omitempty"`
		DebugProfilePort       *int64  `tfsdk:"debug_profile_port" json:"debugProfilePort,omitempty"`
		EtcdV3CompactionPeriod *string `tfsdk:"etcd_v3_compaction_period" json:"etcdV3CompactionPeriod,omitempty"`
		HealthChecks           *string `tfsdk:"health_checks" json:"healthChecks,omitempty"`
		LogSeverityScreen      *string `tfsdk:"log_severity_screen" json:"logSeverityScreen,omitempty"`
		PrometheusMetricsPort  *int64  `tfsdk:"prometheus_metrics_port" json:"prometheusMetricsPort,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CrdProjectcalicoOrgKubeControllersConfigurationV1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_crd_projectcalico_org_kube_controllers_configuration_v1"
}

func (r *CrdProjectcalicoOrgKubeControllersConfigurationV1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
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

			"field_manager": schema.BoolAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
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
						"timeout": schema.StringAttribute{
							Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("30s"),
						},
						"poll_interval": schema.StringAttribute{
							Description:         "The length of time to wait before checking again.",
							MarkdownDescription: "The length of time to wait before checking again.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("5s"),
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
					"timeout": schema.StringAttribute{
						Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
						MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString("30s"),
					},
					"poll_interval": schema.StringAttribute{
						Description:         "The length of time to wait before checking again.",
						MarkdownDescription: "The length of time to wait before checking again.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString("5s"),
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
				Description:         "KubeControllersConfigurationSpec contains the values of the Kubernetes controllers configuration.",
				MarkdownDescription: "KubeControllersConfigurationSpec contains the values of the Kubernetes controllers configuration.",
				Attributes: map[string]schema.Attribute{
					"controllers": schema.SingleNestedAttribute{
						Description:         "Controllers enables and configures individual Kubernetes controllers",
						MarkdownDescription: "Controllers enables and configures individual Kubernetes controllers",
						Attributes: map[string]schema.Attribute{
							"namespace": schema.SingleNestedAttribute{
								Description:         "Namespace enables and configures the namespace controller. Enabled by default, set to nil to disable.",
								MarkdownDescription: "Namespace enables and configures the namespace controller. Enabled by default, set to nil to disable.",
								Attributes: map[string]schema.Attribute{
									"reconciler_period": schema.StringAttribute{
										Description:         "ReconcilerPeriod is the period to perform reconciliation with the Calico datastore. [Default: 5m]",
										MarkdownDescription: "ReconcilerPeriod is the period to perform reconciliation with the Calico datastore. [Default: 5m]",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"node": schema.SingleNestedAttribute{
								Description:         "Node enables and configures the node controller. Enabled by default, set to nil to disable.",
								MarkdownDescription: "Node enables and configures the node controller. Enabled by default, set to nil to disable.",
								Attributes: map[string]schema.Attribute{
									"host_endpoint": schema.SingleNestedAttribute{
										Description:         "HostEndpoint controls syncing nodes to host endpoints. Disabled by default, set to nil to disable.",
										MarkdownDescription: "HostEndpoint controls syncing nodes to host endpoints. Disabled by default, set to nil to disable.",
										Attributes: map[string]schema.Attribute{
											"auto_create": schema.StringAttribute{
												Description:         "AutoCreate enables automatic creation of host endpoints for every node. [Default: Disabled]",
												MarkdownDescription: "AutoCreate enables automatic creation of host endpoints for every node. [Default: Disabled]",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"leak_grace_period": schema.StringAttribute{
										Description:         "LeakGracePeriod is the period used by the controller to determine if an IP address has been leaked. Set to 0 to disable IP garbage collection. [Default: 15m]",
										MarkdownDescription: "LeakGracePeriod is the period used by the controller to determine if an IP address has been leaked. Set to 0 to disable IP garbage collection. [Default: 15m]",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"reconciler_period": schema.StringAttribute{
										Description:         "ReconcilerPeriod is the period to perform reconciliation with the Calico datastore. [Default: 5m]",
										MarkdownDescription: "ReconcilerPeriod is the period to perform reconciliation with the Calico datastore. [Default: 5m]",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"sync_labels": schema.StringAttribute{
										Description:         "SyncLabels controls whether to copy Kubernetes node labels to Calico nodes. [Default: Enabled]",
										MarkdownDescription: "SyncLabels controls whether to copy Kubernetes node labels to Calico nodes. [Default: Enabled]",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"policy": schema.SingleNestedAttribute{
								Description:         "Policy enables and configures the policy controller. Enabled by default, set to nil to disable.",
								MarkdownDescription: "Policy enables and configures the policy controller. Enabled by default, set to nil to disable.",
								Attributes: map[string]schema.Attribute{
									"reconciler_period": schema.StringAttribute{
										Description:         "ReconcilerPeriod is the period to perform reconciliation with the Calico datastore. [Default: 5m]",
										MarkdownDescription: "ReconcilerPeriod is the period to perform reconciliation with the Calico datastore. [Default: 5m]",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"service_account": schema.SingleNestedAttribute{
								Description:         "ServiceAccount enables and configures the service account controller. Enabled by default, set to nil to disable.",
								MarkdownDescription: "ServiceAccount enables and configures the service account controller. Enabled by default, set to nil to disable.",
								Attributes: map[string]schema.Attribute{
									"reconciler_period": schema.StringAttribute{
										Description:         "ReconcilerPeriod is the period to perform reconciliation with the Calico datastore. [Default: 5m]",
										MarkdownDescription: "ReconcilerPeriod is the period to perform reconciliation with the Calico datastore. [Default: 5m]",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"workload_endpoint": schema.SingleNestedAttribute{
								Description:         "WorkloadEndpoint enables and configures the workload endpoint controller. Enabled by default, set to nil to disable.",
								MarkdownDescription: "WorkloadEndpoint enables and configures the workload endpoint controller. Enabled by default, set to nil to disable.",
								Attributes: map[string]schema.Attribute{
									"reconciler_period": schema.StringAttribute{
										Description:         "ReconcilerPeriod is the period to perform reconciliation with the Calico datastore. [Default: 5m]",
										MarkdownDescription: "ReconcilerPeriod is the period to perform reconciliation with the Calico datastore. [Default: 5m]",
										Required:            false,
										Optional:            true,
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

					"debug_profile_port": schema.Int64Attribute{
						Description:         "DebugProfilePort configures the port to serve memory and cpu profiles on. If not specified, profiling is disabled.",
						MarkdownDescription: "DebugProfilePort configures the port to serve memory and cpu profiles on. If not specified, profiling is disabled.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"etcd_v3_compaction_period": schema.StringAttribute{
						Description:         "EtcdV3CompactionPeriod is the period between etcdv3 compaction requests. Set to 0 to disable. [Default: 10m]",
						MarkdownDescription: "EtcdV3CompactionPeriod is the period between etcdv3 compaction requests. Set to 0 to disable. [Default: 10m]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"health_checks": schema.StringAttribute{
						Description:         "HealthChecks enables or disables support for health checks [Default: Enabled]",
						MarkdownDescription: "HealthChecks enables or disables support for health checks [Default: Enabled]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"log_severity_screen": schema.StringAttribute{
						Description:         "LogSeverityScreen is the log severity above which logs are sent to the stdout. [Default: Info]",
						MarkdownDescription: "LogSeverityScreen is the log severity above which logs are sent to the stdout. [Default: Info]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"prometheus_metrics_port": schema.Int64Attribute{
						Description:         "PrometheusMetricsPort is the TCP port that the Prometheus metrics server should bind to. Set to 0 to disable. [Default: 9094]",
						MarkdownDescription: "PrometheusMetricsPort is the TCP port that the Prometheus metrics server should bind to. Set to 0 to disable. [Default: 9094]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *CrdProjectcalicoOrgKubeControllersConfigurationV1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
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

func (r *CrdProjectcalicoOrgKubeControllersConfigurationV1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_crd_projectcalico_org_kube_controllers_configuration_v1")

	var model CrdProjectcalicoOrgKubeControllersConfigurationV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("crd.projectcalico.org/v1")
	model.Kind = pointer.String("KubeControllersConfiguration")

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
		Resource(k8sSchema.GroupVersionResource{Group: "crd.projectcalico.org", Version: "v1", Resource: "kubecontrollersconfigurations"}).
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

	var readResponse CrdProjectcalicoOrgKubeControllersConfigurationV1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *CrdProjectcalicoOrgKubeControllersConfigurationV1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_crd_projectcalico_org_kube_controllers_configuration_v1")

	var data CrdProjectcalicoOrgKubeControllersConfigurationV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "crd.projectcalico.org", Version: "v1", Resource: "kubecontrollersconfigurations"}).
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

	var readResponse CrdProjectcalicoOrgKubeControllersConfigurationV1ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *CrdProjectcalicoOrgKubeControllersConfigurationV1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_crd_projectcalico_org_kube_controllers_configuration_v1")

	var model CrdProjectcalicoOrgKubeControllersConfigurationV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("crd.projectcalico.org/v1")
	model.Kind = pointer.String("KubeControllersConfiguration")

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
		Resource(k8sSchema.GroupVersionResource{Group: "crd.projectcalico.org", Version: "v1", Resource: "kubecontrollersconfigurations"}).
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

	var readResponse CrdProjectcalicoOrgKubeControllersConfigurationV1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *CrdProjectcalicoOrgKubeControllersConfigurationV1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_crd_projectcalico_org_kube_controllers_configuration_v1")

	var data CrdProjectcalicoOrgKubeControllersConfigurationV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "crd.projectcalico.org", Version: "v1", Resource: "kubecontrollersconfigurations"}).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if utilities.IsDeletionError(err) {
		response.Diagnostics.Append(utilities.DeleteError(err))
		return
	}

	if !data.WaitForDelete.IsNull() {
		timeout := utilities.DetermineTimeout(data.WaitForDelete.Attributes())
		pollInterval := utilities.DeterminePollInterval(data.WaitForDelete.Attributes())

		startTime := time.Now()
		for {
			_, err := r.kubernetesClient.
				Resource(k8sSchema.GroupVersionResource{Group: "crd.projectcalico.org", Version: "v1", Resource: "kubecontrollersconfigurations"}).
				Get(ctx, data.Metadata.Name, meta.GetOptions{})
			if utilities.IsNotFound(err) || timeout == time.Second*0 {
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

func (r *CrdProjectcalicoOrgKubeControllersConfigurationV1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
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
