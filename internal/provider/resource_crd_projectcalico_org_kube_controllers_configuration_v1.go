/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type CrdProjectcalicoOrgKubeControllersConfigurationV1Resource struct{}

var (
	_ resource.Resource = (*CrdProjectcalicoOrgKubeControllersConfigurationV1Resource)(nil)
)

type CrdProjectcalicoOrgKubeControllersConfigurationV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type CrdProjectcalicoOrgKubeControllersConfigurationV1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Controllers *struct {
			Namespace *struct {
				ReconcilerPeriod *string `tfsdk:"reconciler_period" yaml:"reconcilerPeriod,omitempty"`
			} `tfsdk:"namespace" yaml:"namespace,omitempty"`

			Node *struct {
				HostEndpoint *struct {
					AutoCreate *string `tfsdk:"auto_create" yaml:"autoCreate,omitempty"`
				} `tfsdk:"host_endpoint" yaml:"hostEndpoint,omitempty"`

				LeakGracePeriod *string `tfsdk:"leak_grace_period" yaml:"leakGracePeriod,omitempty"`

				ReconcilerPeriod *string `tfsdk:"reconciler_period" yaml:"reconcilerPeriod,omitempty"`

				SyncLabels *string `tfsdk:"sync_labels" yaml:"syncLabels,omitempty"`
			} `tfsdk:"node" yaml:"node,omitempty"`

			Policy *struct {
				ReconcilerPeriod *string `tfsdk:"reconciler_period" yaml:"reconcilerPeriod,omitempty"`
			} `tfsdk:"policy" yaml:"policy,omitempty"`

			ServiceAccount *struct {
				ReconcilerPeriod *string `tfsdk:"reconciler_period" yaml:"reconcilerPeriod,omitempty"`
			} `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

			WorkloadEndpoint *struct {
				ReconcilerPeriod *string `tfsdk:"reconciler_period" yaml:"reconcilerPeriod,omitempty"`
			} `tfsdk:"workload_endpoint" yaml:"workloadEndpoint,omitempty"`
		} `tfsdk:"controllers" yaml:"controllers,omitempty"`

		DebugProfilePort *int64 `tfsdk:"debug_profile_port" yaml:"debugProfilePort,omitempty"`

		EtcdV3CompactionPeriod *string `tfsdk:"etcd_v3_compaction_period" yaml:"etcdV3CompactionPeriod,omitempty"`

		HealthChecks *string `tfsdk:"health_checks" yaml:"healthChecks,omitempty"`

		LogSeverityScreen *string `tfsdk:"log_severity_screen" yaml:"logSeverityScreen,omitempty"`

		PrometheusMetricsPort *int64 `tfsdk:"prometheus_metrics_port" yaml:"prometheusMetricsPort,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewCrdProjectcalicoOrgKubeControllersConfigurationV1Resource() resource.Resource {
	return &CrdProjectcalicoOrgKubeControllersConfigurationV1Resource{}
}

func (r *CrdProjectcalicoOrgKubeControllersConfigurationV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_crd_projectcalico_org_kube_controllers_configuration_v1"
}

func (r *CrdProjectcalicoOrgKubeControllersConfigurationV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "",
		MarkdownDescription: "",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "KubeControllersConfigurationSpec contains the values of the Kubernetes controllers configuration.",
				MarkdownDescription: "KubeControllersConfigurationSpec contains the values of the Kubernetes controllers configuration.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"controllers": {
						Description:         "Controllers enables and configures individual Kubernetes controllers",
						MarkdownDescription: "Controllers enables and configures individual Kubernetes controllers",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"namespace": {
								Description:         "Namespace enables and configures the namespace controller. Enabled by default, set to nil to disable.",
								MarkdownDescription: "Namespace enables and configures the namespace controller. Enabled by default, set to nil to disable.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"reconciler_period": {
										Description:         "ReconcilerPeriod is the period to perform reconciliation with the Calico datastore. [Default: 5m]",
										MarkdownDescription: "ReconcilerPeriod is the period to perform reconciliation with the Calico datastore. [Default: 5m]",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"node": {
								Description:         "Node enables and configures the node controller. Enabled by default, set to nil to disable.",
								MarkdownDescription: "Node enables and configures the node controller. Enabled by default, set to nil to disable.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"host_endpoint": {
										Description:         "HostEndpoint controls syncing nodes to host endpoints. Disabled by default, set to nil to disable.",
										MarkdownDescription: "HostEndpoint controls syncing nodes to host endpoints. Disabled by default, set to nil to disable.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"auto_create": {
												Description:         "AutoCreate enables automatic creation of host endpoints for every node. [Default: Disabled]",
												MarkdownDescription: "AutoCreate enables automatic creation of host endpoints for every node. [Default: Disabled]",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"leak_grace_period": {
										Description:         "LeakGracePeriod is the period used by the controller to determine if an IP address has been leaked. Set to 0 to disable IP garbage collection. [Default: 15m]",
										MarkdownDescription: "LeakGracePeriod is the period used by the controller to determine if an IP address has been leaked. Set to 0 to disable IP garbage collection. [Default: 15m]",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"reconciler_period": {
										Description:         "ReconcilerPeriod is the period to perform reconciliation with the Calico datastore. [Default: 5m]",
										MarkdownDescription: "ReconcilerPeriod is the period to perform reconciliation with the Calico datastore. [Default: 5m]",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"sync_labels": {
										Description:         "SyncLabels controls whether to copy Kubernetes node labels to Calico nodes. [Default: Enabled]",
										MarkdownDescription: "SyncLabels controls whether to copy Kubernetes node labels to Calico nodes. [Default: Enabled]",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"policy": {
								Description:         "Policy enables and configures the policy controller. Enabled by default, set to nil to disable.",
								MarkdownDescription: "Policy enables and configures the policy controller. Enabled by default, set to nil to disable.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"reconciler_period": {
										Description:         "ReconcilerPeriod is the period to perform reconciliation with the Calico datastore. [Default: 5m]",
										MarkdownDescription: "ReconcilerPeriod is the period to perform reconciliation with the Calico datastore. [Default: 5m]",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"service_account": {
								Description:         "ServiceAccount enables and configures the service account controller. Enabled by default, set to nil to disable.",
								MarkdownDescription: "ServiceAccount enables and configures the service account controller. Enabled by default, set to nil to disable.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"reconciler_period": {
										Description:         "ReconcilerPeriod is the period to perform reconciliation with the Calico datastore. [Default: 5m]",
										MarkdownDescription: "ReconcilerPeriod is the period to perform reconciliation with the Calico datastore. [Default: 5m]",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"workload_endpoint": {
								Description:         "WorkloadEndpoint enables and configures the workload endpoint controller. Enabled by default, set to nil to disable.",
								MarkdownDescription: "WorkloadEndpoint enables and configures the workload endpoint controller. Enabled by default, set to nil to disable.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"reconciler_period": {
										Description:         "ReconcilerPeriod is the period to perform reconciliation with the Calico datastore. [Default: 5m]",
										MarkdownDescription: "ReconcilerPeriod is the period to perform reconciliation with the Calico datastore. [Default: 5m]",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"debug_profile_port": {
						Description:         "DebugProfilePort configures the port to serve memory and cpu profiles on. If not specified, profiling is disabled.",
						MarkdownDescription: "DebugProfilePort configures the port to serve memory and cpu profiles on. If not specified, profiling is disabled.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"etcd_v3_compaction_period": {
						Description:         "EtcdV3CompactionPeriod is the period between etcdv3 compaction requests. Set to 0 to disable. [Default: 10m]",
						MarkdownDescription: "EtcdV3CompactionPeriod is the period between etcdv3 compaction requests. Set to 0 to disable. [Default: 10m]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"health_checks": {
						Description:         "HealthChecks enables or disables support for health checks [Default: Enabled]",
						MarkdownDescription: "HealthChecks enables or disables support for health checks [Default: Enabled]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"log_severity_screen": {
						Description:         "LogSeverityScreen is the log severity above which logs are sent to the stdout. [Default: Info]",
						MarkdownDescription: "LogSeverityScreen is the log severity above which logs are sent to the stdout. [Default: Info]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"prometheus_metrics_port": {
						Description:         "PrometheusMetricsPort is the TCP port that the Prometheus metrics server should bind to. Set to 0 to disable. [Default: 9094]",
						MarkdownDescription: "PrometheusMetricsPort is the TCP port that the Prometheus metrics server should bind to. Set to 0 to disable. [Default: 9094]",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *CrdProjectcalicoOrgKubeControllersConfigurationV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_crd_projectcalico_org_kube_controllers_configuration_v1")

	var state CrdProjectcalicoOrgKubeControllersConfigurationV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CrdProjectcalicoOrgKubeControllersConfigurationV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("crd.projectcalico.org/v1")
	goModel.Kind = utilities.Ptr("KubeControllersConfiguration")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CrdProjectcalicoOrgKubeControllersConfigurationV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_crd_projectcalico_org_kube_controllers_configuration_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *CrdProjectcalicoOrgKubeControllersConfigurationV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_crd_projectcalico_org_kube_controllers_configuration_v1")

	var state CrdProjectcalicoOrgKubeControllersConfigurationV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CrdProjectcalicoOrgKubeControllersConfigurationV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("crd.projectcalico.org/v1")
	goModel.Kind = utilities.Ptr("KubeControllersConfiguration")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CrdProjectcalicoOrgKubeControllersConfigurationV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_crd_projectcalico_org_kube_controllers_configuration_v1")
	// NO-OP: Terraform removes the state automatically for us
}
