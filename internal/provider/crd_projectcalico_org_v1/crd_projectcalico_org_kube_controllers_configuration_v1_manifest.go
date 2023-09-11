/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package crd_projectcalico_org_v1

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
	_ datasource.DataSource = &CrdProjectcalicoOrgKubeControllersConfigurationV1Manifest{}
)

func NewCrdProjectcalicoOrgKubeControllersConfigurationV1Manifest() datasource.DataSource {
	return &CrdProjectcalicoOrgKubeControllersConfigurationV1Manifest{}
}

type CrdProjectcalicoOrgKubeControllersConfigurationV1Manifest struct{}

type CrdProjectcalicoOrgKubeControllersConfigurationV1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
	YAML types.String `tfsdk:"yaml" json:"-"`

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

func (r *CrdProjectcalicoOrgKubeControllersConfigurationV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_crd_projectcalico_org_kube_controllers_configuration_v1_manifest"
}

func (r *CrdProjectcalicoOrgKubeControllersConfigurationV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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

func (r *CrdProjectcalicoOrgKubeControllersConfigurationV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_crd_projectcalico_org_kube_controllers_configuration_v1_manifest")

	var model CrdProjectcalicoOrgKubeControllersConfigurationV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("crd.projectcalico.org/v1")
	model.Kind = pointer.String("KubeControllersConfiguration")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
