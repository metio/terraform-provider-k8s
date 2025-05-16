/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package clickhouse_altinity_com_v1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	_ datasource.DataSource = &ClickhouseAltinityComClickHouseOperatorConfigurationV1Manifest{}
)

func NewClickhouseAltinityComClickHouseOperatorConfigurationV1Manifest() datasource.DataSource {
	return &ClickhouseAltinityComClickHouseOperatorConfigurationV1Manifest{}
}

type ClickhouseAltinityComClickHouseOperatorConfigurationV1Manifest struct{}

type ClickhouseAltinityComClickHouseOperatorConfigurationV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Annotation *struct {
			Exclude *[]string `tfsdk:"exclude" json:"exclude,omitempty"`
			Include *[]string `tfsdk:"include" json:"include,omitempty"`
		} `tfsdk:"annotation" json:"annotation,omitempty"`
		Clickhouse *struct {
			Access *struct {
				Password *string `tfsdk:"password" json:"password,omitempty"`
				Port     *int64  `tfsdk:"port" json:"port,omitempty"`
				RootCA   *string `tfsdk:"root_ca" json:"rootCA,omitempty"`
				Scheme   *string `tfsdk:"scheme" json:"scheme,omitempty"`
				Secret   *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"secret" json:"secret,omitempty"`
				Timeouts *struct {
					Connect *int64 `tfsdk:"connect" json:"connect,omitempty"`
					Query   *int64 `tfsdk:"query" json:"query,omitempty"`
				} `tfsdk:"timeouts" json:"timeouts,omitempty"`
				Username *string `tfsdk:"username" json:"username,omitempty"`
			} `tfsdk:"access" json:"access,omitempty"`
			Configuration *struct {
				File *struct {
					Path *struct {
						Common *string `tfsdk:"common" json:"common,omitempty"`
						Host   *string `tfsdk:"host" json:"host,omitempty"`
						User   *string `tfsdk:"user" json:"user,omitempty"`
					} `tfsdk:"path" json:"path,omitempty"`
				} `tfsdk:"file" json:"file,omitempty"`
				Network *struct {
					HostRegexpTemplate *string `tfsdk:"host_regexp_template" json:"hostRegexpTemplate,omitempty"`
				} `tfsdk:"network" json:"network,omitempty"`
				User *struct {
					Default *struct {
						NetworksIP *[]string `tfsdk:"networks_ip" json:"networksIP,omitempty"`
						Password   *string   `tfsdk:"password" json:"password,omitempty"`
						Profile    *string   `tfsdk:"profile" json:"profile,omitempty"`
						Quota      *string   `tfsdk:"quota" json:"quota,omitempty"`
					} `tfsdk:"default" json:"default,omitempty"`
				} `tfsdk:"user" json:"user,omitempty"`
			} `tfsdk:"configuration" json:"configuration,omitempty"`
			ConfigurationRestartPolicy *struct {
				Rules *[]struct {
					Rules   *[]map[string]string `tfsdk:"rules" json:"rules,omitempty"`
					Version *string              `tfsdk:"version" json:"version,omitempty"`
				} `tfsdk:"rules" json:"rules,omitempty"`
			} `tfsdk:"configuration_restart_policy" json:"configurationRestartPolicy,omitempty"`
			Metrics *struct {
				Timeouts *struct {
					Collect *int64 `tfsdk:"collect" json:"collect,omitempty"`
				} `tfsdk:"timeouts" json:"timeouts,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
		} `tfsdk:"clickhouse" json:"clickhouse,omitempty"`
		Label *struct {
			AppendScope *string   `tfsdk:"append_scope" json:"appendScope,omitempty"`
			Exclude     *[]string `tfsdk:"exclude" json:"exclude,omitempty"`
			Include     *[]string `tfsdk:"include" json:"include,omitempty"`
		} `tfsdk:"label" json:"label,omitempty"`
		Logger *struct {
			Alsologtostderr  *string `tfsdk:"alsologtostderr" json:"alsologtostderr,omitempty"`
			Log_backtrace_at *string `tfsdk:"log_backtrace_at" json:"log_backtrace_at,omitempty"`
			Logtostderr      *string `tfsdk:"logtostderr" json:"logtostderr,omitempty"`
			Stderrthreshold  *string `tfsdk:"stderrthreshold" json:"stderrthreshold,omitempty"`
			V                *string `tfsdk:"v" json:"v,omitempty"`
			Vmodule          *string `tfsdk:"vmodule" json:"vmodule,omitempty"`
		} `tfsdk:"logger" json:"logger,omitempty"`
		Metrics *struct {
			Labels *struct {
				Exclude *[]string `tfsdk:"exclude" json:"exclude,omitempty"`
			} `tfsdk:"labels" json:"labels,omitempty"`
		} `tfsdk:"metrics" json:"metrics,omitempty"`
		Pod *struct {
			TerminationGracePeriod *int64 `tfsdk:"termination_grace_period" json:"terminationGracePeriod,omitempty"`
		} `tfsdk:"pod" json:"pod,omitempty"`
		Reconcile *struct {
			Host *struct {
				Wait *struct {
					Exclude *string `tfsdk:"exclude" json:"exclude,omitempty"`
					Include *string `tfsdk:"include" json:"include,omitempty"`
					Queries *string `tfsdk:"queries" json:"queries,omitempty"`
				} `tfsdk:"wait" json:"wait,omitempty"`
			} `tfsdk:"host" json:"host,omitempty"`
			Runtime *struct {
				ReconcileCHIsThreadsNumber           *int64 `tfsdk:"reconcile_ch_is_threads_number" json:"reconcileCHIsThreadsNumber,omitempty"`
				ReconcileShardsMaxConcurrencyPercent *int64 `tfsdk:"reconcile_shards_max_concurrency_percent" json:"reconcileShardsMaxConcurrencyPercent,omitempty"`
				ReconcileShardsThreadsNumber         *int64 `tfsdk:"reconcile_shards_threads_number" json:"reconcileShardsThreadsNumber,omitempty"`
			} `tfsdk:"runtime" json:"runtime,omitempty"`
			StatefulSet *struct {
				Create *struct {
					OnFailure *string `tfsdk:"on_failure" json:"onFailure,omitempty"`
				} `tfsdk:"create" json:"create,omitempty"`
				Update *struct {
					OnFailure    *string `tfsdk:"on_failure" json:"onFailure,omitempty"`
					PollInterval *int64  `tfsdk:"poll_interval" json:"pollInterval,omitempty"`
					Timeout      *int64  `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"update" json:"update,omitempty"`
			} `tfsdk:"stateful_set" json:"statefulSet,omitempty"`
		} `tfsdk:"reconcile" json:"reconcile,omitempty"`
		StatefulSet *struct {
			RevisionHistoryLimit *int64 `tfsdk:"revision_history_limit" json:"revisionHistoryLimit,omitempty"`
		} `tfsdk:"stateful_set" json:"statefulSet,omitempty"`
		Status *struct {
			Fields *struct {
				Action  *string `tfsdk:"action" json:"action,omitempty"`
				Actions *string `tfsdk:"actions" json:"actions,omitempty"`
				Error   *string `tfsdk:"error" json:"error,omitempty"`
				Errors  *string `tfsdk:"errors" json:"errors,omitempty"`
			} `tfsdk:"fields" json:"fields,omitempty"`
		} `tfsdk:"status" json:"status,omitempty"`
		Template *struct {
			Chi *struct {
				Path   *string `tfsdk:"path" json:"path,omitempty"`
				Policy *string `tfsdk:"policy" json:"policy,omitempty"`
			} `tfsdk:"chi" json:"chi,omitempty"`
		} `tfsdk:"template" json:"template,omitempty"`
		Watch *struct {
			Namespaces *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
		} `tfsdk:"watch" json:"watch,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ClickhouseAltinityComClickHouseOperatorConfigurationV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_clickhouse_altinity_com_click_house_operator_configuration_v1_manifest"
}

func (r *ClickhouseAltinityComClickHouseOperatorConfigurationV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "allows customize 'clickhouse-operator' settings, need restart clickhouse-operator pod after adding, more details https://github.com/Altinity/clickhouse-operator/blob/master/docs/operator_configuration.md",
		MarkdownDescription: "allows customize 'clickhouse-operator' settings, need restart clickhouse-operator pod after adding, more details https://github.com/Altinity/clickhouse-operator/blob/master/docs/operator_configuration.md",
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
				Description:         "Allows to define settings of the clickhouse-operator. More info: https://github.com/Altinity/clickhouse-operator/blob/master/config/config.yaml Check into etc-clickhouse-operator* ConfigMaps if you need more control ",
				MarkdownDescription: "Allows to define settings of the clickhouse-operator. More info: https://github.com/Altinity/clickhouse-operator/blob/master/config/config.yaml Check into etc-clickhouse-operator* ConfigMaps if you need more control ",
				Attributes: map[string]schema.Attribute{
					"annotation": schema.SingleNestedAttribute{
						Description:         "defines which metadata.annotations items will include or exclude during render StatefulSet, Pod, PVC resources",
						MarkdownDescription: "defines which metadata.annotations items will include or exclude during render StatefulSet, Pod, PVC resources",
						Attributes: map[string]schema.Attribute{
							"exclude": schema.ListAttribute{
								Description:         "When propagating labels from the chi's 'metadata.annotations' section to child objects' 'metadata.annotations', exclude annotations with names from the following list ",
								MarkdownDescription: "When propagating labels from the chi's 'metadata.annotations' section to child objects' 'metadata.annotations', exclude annotations with names from the following list ",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"include": schema.ListAttribute{
								Description:         "When propagating labels from the chi's 'metadata.annotations' section to child objects' 'metadata.annotations', include annotations with names from the following list ",
								MarkdownDescription: "When propagating labels from the chi's 'metadata.annotations' section to child objects' 'metadata.annotations', include annotations with names from the following list ",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"clickhouse": schema.SingleNestedAttribute{
						Description:         "Clickhouse related parameters used by clickhouse-operator",
						MarkdownDescription: "Clickhouse related parameters used by clickhouse-operator",
						Attributes: map[string]schema.Attribute{
							"access": schema.SingleNestedAttribute{
								Description:         "parameters which use for connect to clickhouse from clickhouse-operator deployment",
								MarkdownDescription: "parameters which use for connect to clickhouse from clickhouse-operator deployment",
								Attributes: map[string]schema.Attribute{
									"password": schema.StringAttribute{
										Description:         "ClickHouse password to be used by operator to connect to ClickHouse instances, deprecated, use chCredentialsSecretName",
										MarkdownDescription: "ClickHouse password to be used by operator to connect to ClickHouse instances, deprecated, use chCredentialsSecretName",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"port": schema.Int64Attribute{
										Description:         "Port to be used by operator to connect to ClickHouse instances",
										MarkdownDescription: "Port to be used by operator to connect to ClickHouse instances",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
											int64validator.AtMost(65535),
										},
									},

									"root_ca": schema.StringAttribute{
										Description:         "Root certificate authority that clients use when verifying server certificates. Used for https connection to ClickHouse",
										MarkdownDescription: "Root certificate authority that clients use when verifying server certificates. Used for https connection to ClickHouse",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"scheme": schema.StringAttribute{
										Description:         "The scheme to user for connecting to ClickHouse. Possible values: http, https, auto",
										MarkdownDescription: "The scheme to user for connecting to ClickHouse. Possible values: http, https, auto",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"secret": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of k8s Secret with username and password to be used by operator to connect to ClickHouse instances",
												MarkdownDescription: "Name of k8s Secret with username and password to be used by operator to connect to ClickHouse instances",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "Location of k8s Secret with username and password to be used by operator to connect to ClickHouse instances",
												MarkdownDescription: "Location of k8s Secret with username and password to be used by operator to connect to ClickHouse instances",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"timeouts": schema.SingleNestedAttribute{
										Description:         "Timeouts used to limit connection and queries from the operator to ClickHouse instances, In seconds",
										MarkdownDescription: "Timeouts used to limit connection and queries from the operator to ClickHouse instances, In seconds",
										Attributes: map[string]schema.Attribute{
											"connect": schema.Int64Attribute{
												Description:         "Timout to setup connection from the operator to ClickHouse instances. In seconds.",
												MarkdownDescription: "Timout to setup connection from the operator to ClickHouse instances. In seconds.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
													int64validator.AtMost(10),
												},
											},

											"query": schema.Int64Attribute{
												Description:         "Timout to perform SQL query from the operator to ClickHouse instances. In seconds.",
												MarkdownDescription: "Timout to perform SQL query from the operator to ClickHouse instances. In seconds.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
													int64validator.AtMost(600),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"username": schema.StringAttribute{
										Description:         "ClickHouse username to be used by operator to connect to ClickHouse instances, deprecated, use chCredentialsSecretName",
										MarkdownDescription: "ClickHouse username to be used by operator to connect to ClickHouse instances, deprecated, use chCredentialsSecretName",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"configuration": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"file": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"path": schema.SingleNestedAttribute{
												Description:         "Each 'path' can be either absolute or relative. In case path is absolute - it is used as is. In case path is relative - it is relative to the folder where configuration file you are reading right now is located. ",
												MarkdownDescription: "Each 'path' can be either absolute or relative. In case path is absolute - it is used as is. In case path is relative - it is relative to the folder where configuration file you are reading right now is located. ",
												Attributes: map[string]schema.Attribute{
													"common": schema.StringAttribute{
														Description:         "Path to the folder where ClickHouse configuration files common for all instances within a CHI are located. Default value - config.d ",
														MarkdownDescription: "Path to the folder where ClickHouse configuration files common for all instances within a CHI are located. Default value - config.d ",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"host": schema.StringAttribute{
														Description:         "Path to the folder where ClickHouse configuration files unique for each instance (host) within a CHI are located. Default value - conf.d ",
														MarkdownDescription: "Path to the folder where ClickHouse configuration files unique for each instance (host) within a CHI are located. Default value - conf.d ",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"user": schema.StringAttribute{
														Description:         "Path to the folder where ClickHouse configuration files with users settings are located. Files are common for all instances within a CHI. Default value - users.d ",
														MarkdownDescription: "Path to the folder where ClickHouse configuration files with users settings are located. Files are common for all instances within a CHI. Default value - users.d ",
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
										Required: false,
										Optional: true,
										Computed: false,
									},

									"network": schema.SingleNestedAttribute{
										Description:         "Default network parameters for any user which will create",
										MarkdownDescription: "Default network parameters for any user which will create",
										Attributes: map[string]schema.Attribute{
											"host_regexp_template": schema.StringAttribute{
												Description:         "ClickHouse server configuration '<host_regexp>...</host_regexp>' for any <user>",
												MarkdownDescription: "ClickHouse server configuration '<host_regexp>...</host_regexp>' for any <user>",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"user": schema.SingleNestedAttribute{
										Description:         "Default parameters for any user which will create",
										MarkdownDescription: "Default parameters for any user which will create",
										Attributes: map[string]schema.Attribute{
											"default": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"networks_ip": schema.ListAttribute{
														Description:         "ClickHouse server configuration '<networks><ip>...</ip></networks>' for any <user>",
														MarkdownDescription: "ClickHouse server configuration '<networks><ip>...</ip></networks>' for any <user>",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"password": schema.StringAttribute{
														Description:         "ClickHouse server configuration '<password>...</password>' for any <user>",
														MarkdownDescription: "ClickHouse server configuration '<password>...</password>' for any <user>",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"profile": schema.StringAttribute{
														Description:         "ClickHouse server configuration '<profile>...</profile>' for any <user>",
														MarkdownDescription: "ClickHouse server configuration '<profile>...</profile>' for any <user>",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"quota": schema.StringAttribute{
														Description:         "ClickHouse server configuration '<quota>...</quota>' for any <user>",
														MarkdownDescription: "ClickHouse server configuration '<quota>...</quota>' for any <user>",
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
										Required: false,
										Optional: true,
										Computed: false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"configuration_restart_policy": schema.SingleNestedAttribute{
								Description:         "Configuration restart policy describes what configuration changes require ClickHouse restart",
								MarkdownDescription: "Configuration restart policy describes what configuration changes require ClickHouse restart",
								Attributes: map[string]schema.Attribute{
									"rules": schema.ListNestedAttribute{
										Description:         "Array of set of rules per specified ClickHouse versions",
										MarkdownDescription: "Array of set of rules per specified ClickHouse versions",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"rules": schema.ListAttribute{
													Description:         "Set of configuration rules for specified ClickHouse version",
													MarkdownDescription: "Set of configuration rules for specified ClickHouse version",
													ElementType:         types.MapType{ElemType: types.StringType},
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"version": schema.StringAttribute{
													Description:         "ClickHouse version expression",
													MarkdownDescription: "ClickHouse version expression",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"metrics": schema.SingleNestedAttribute{
								Description:         "parameters which use for connect to fetch metrics from clickhouse by clickhouse-operator",
								MarkdownDescription: "parameters which use for connect to fetch metrics from clickhouse by clickhouse-operator",
								Attributes: map[string]schema.Attribute{
									"timeouts": schema.SingleNestedAttribute{
										Description:         "Timeouts used to limit connection and queries from the metrics exporter to ClickHouse instances Specified in seconds. ",
										MarkdownDescription: "Timeouts used to limit connection and queries from the metrics exporter to ClickHouse instances Specified in seconds. ",
										Attributes: map[string]schema.Attribute{
											"collect": schema.Int64Attribute{
												Description:         "Timeout used to limit metrics collection request. In seconds. Upon reaching this timeout metrics collection is aborted and no more metrics are collected in this cycle. All collected metrics are returned. ",
												MarkdownDescription: "Timeout used to limit metrics collection request. In seconds. Upon reaching this timeout metrics collection is aborted and no more metrics are collected in this cycle. All collected metrics are returned. ",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
													int64validator.AtMost(600),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"label": schema.SingleNestedAttribute{
						Description:         "defines which metadata.labels will include or exclude during render StatefulSet, Pod, PVC resources",
						MarkdownDescription: "defines which metadata.labels will include or exclude during render StatefulSet, Pod, PVC resources",
						Attributes: map[string]schema.Attribute{
							"append_scope": schema.StringAttribute{
								Description:         "Whether to append *Scope* labels to StatefulSet and Pod - 'LabelShardScopeIndex' - 'LabelReplicaScopeIndex' - 'LabelCHIScopeIndex' - 'LabelCHIScopeCycleSize' - 'LabelCHIScopeCycleIndex' - 'LabelCHIScopeCycleOffset' - 'LabelClusterScopeIndex' - 'LabelClusterScopeCycleSize' - 'LabelClusterScopeCycleIndex' - 'LabelClusterScopeCycleOffset' ",
								MarkdownDescription: "Whether to append *Scope* labels to StatefulSet and Pod - 'LabelShardScopeIndex' - 'LabelReplicaScopeIndex' - 'LabelCHIScopeIndex' - 'LabelCHIScopeCycleSize' - 'LabelCHIScopeCycleIndex' - 'LabelCHIScopeCycleOffset' - 'LabelClusterScopeIndex' - 'LabelClusterScopeCycleSize' - 'LabelClusterScopeCycleIndex' - 'LabelClusterScopeCycleOffset' ",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("", "0", "1", "False", "false", "True", "true", "No", "no", "Yes", "yes", "Off", "off", "On", "on", "Disable", "disable", "Enable", "enable", "Disabled", "disabled", "Enabled", "enabled"),
								},
							},

							"exclude": schema.ListAttribute{
								Description:         "When propagating labels from the chi's 'metadata.labels' section to child objects' 'metadata.labels', exclude labels from the following list ",
								MarkdownDescription: "When propagating labels from the chi's 'metadata.labels' section to child objects' 'metadata.labels', exclude labels from the following list ",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"include": schema.ListAttribute{
								Description:         "When propagating labels from the chi's 'metadata.labels' section to child objects' 'metadata.labels', include labels from the following list ",
								MarkdownDescription: "When propagating labels from the chi's 'metadata.labels' section to child objects' 'metadata.labels', include labels from the following list ",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"logger": schema.SingleNestedAttribute{
						Description:         "allow setup clickhouse-operator logger behavior",
						MarkdownDescription: "allow setup clickhouse-operator logger behavior",
						Attributes: map[string]schema.Attribute{
							"alsologtostderr": schema.StringAttribute{
								Description:         "boolean allows logs to stderr and files both",
								MarkdownDescription: "boolean allows logs to stderr and files both",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"log_backtrace_at": schema.StringAttribute{
								Description:         "It can be set to a file and line number with a logging line. Ex.: file.go:123 Each time when this line is being executed, a stack trace will be written to the Info log. ",
								MarkdownDescription: "It can be set to a file and line number with a logging line. Ex.: file.go:123 Each time when this line is being executed, a stack trace will be written to the Info log. ",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"logtostderr": schema.StringAttribute{
								Description:         "boolean, allows logs to stderr",
								MarkdownDescription: "boolean, allows logs to stderr",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"stderrthreshold": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"v": schema.StringAttribute{
								Description:         "verbosity level of clickhouse-operator log, default - 1 max - 9",
								MarkdownDescription: "verbosity level of clickhouse-operator log, default - 1 max - 9",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"vmodule": schema.StringAttribute{
								Description:         "Comma-separated list of filename=N, where filename (can be a pattern) must have no .go ext, and N is a V level. Ex.: file*=2 sets the 'V' to 2 in all files with names like file*. ",
								MarkdownDescription: "Comma-separated list of filename=N, where filename (can be a pattern) must have no .go ext, and N is a V level. Ex.: file*=2 sets the 'V' to 2 in all files with names like file*. ",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"metrics": schema.SingleNestedAttribute{
						Description:         "defines metrics exporter options",
						MarkdownDescription: "defines metrics exporter options",
						Attributes: map[string]schema.Attribute{
							"labels": schema.SingleNestedAttribute{
								Description:         "defines metric labels options",
								MarkdownDescription: "defines metric labels options",
								Attributes: map[string]schema.Attribute{
									"exclude": schema.ListAttribute{
										Description:         "When adding labels to a metric exclude labels with names from the following list ",
										MarkdownDescription: "When adding labels to a metric exclude labels with names from the following list ",
										ElementType:         types.StringType,
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"pod": schema.SingleNestedAttribute{
						Description:         "define pod specific parameters",
						MarkdownDescription: "define pod specific parameters",
						Attributes: map[string]schema.Attribute{
							"termination_grace_period": schema.Int64Attribute{
								Description:         "Optional duration in seconds the pod needs to terminate gracefully. Look details in 'pod.spec.terminationGracePeriodSeconds' ",
								MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully. Look details in 'pod.spec.terminationGracePeriodSeconds' ",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"reconcile": schema.SingleNestedAttribute{
						Description:         "allow tuning reconciling process",
						MarkdownDescription: "allow tuning reconciling process",
						Attributes: map[string]schema.Attribute{
							"host": schema.SingleNestedAttribute{
								Description:         "Whether the operator during reconcile procedure should wait for a ClickHouse host: - to be excluded from a ClickHouse cluster - to complete all running queries - to be included into a ClickHouse cluster respectfully before moving forward ",
								MarkdownDescription: "Whether the operator during reconcile procedure should wait for a ClickHouse host: - to be excluded from a ClickHouse cluster - to complete all running queries - to be included into a ClickHouse cluster respectfully before moving forward ",
								Attributes: map[string]schema.Attribute{
									"wait": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"exclude": schema.StringAttribute{
												Description:         "Whether the operator during reconcile procedure should wait for a ClickHouse host to be excluded from a ClickHouse cluster",
												MarkdownDescription: "Whether the operator during reconcile procedure should wait for a ClickHouse host to be excluded from a ClickHouse cluster",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("", "0", "1", "False", "false", "True", "true", "No", "no", "Yes", "yes", "Off", "off", "On", "on", "Disable", "disable", "Enable", "enable", "Disabled", "disabled", "Enabled", "enabled"),
												},
											},

											"include": schema.StringAttribute{
												Description:         "Whether the operator during reconcile procedure should wait for a ClickHouse host to be included into a ClickHouse cluster",
												MarkdownDescription: "Whether the operator during reconcile procedure should wait for a ClickHouse host to be included into a ClickHouse cluster",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("", "0", "1", "False", "false", "True", "true", "No", "no", "Yes", "yes", "Off", "off", "On", "on", "Disable", "disable", "Enable", "enable", "Disabled", "disabled", "Enabled", "enabled"),
												},
											},

											"queries": schema.StringAttribute{
												Description:         "Whether the operator during reconcile procedure should wait for a ClickHouse host to complete all running queries",
												MarkdownDescription: "Whether the operator during reconcile procedure should wait for a ClickHouse host to complete all running queries",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("", "0", "1", "False", "false", "True", "true", "No", "no", "Yes", "yes", "Off", "off", "On", "on", "Disable", "disable", "Enable", "enable", "Disabled", "disabled", "Enabled", "enabled"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"runtime": schema.SingleNestedAttribute{
								Description:         "runtime parameters for clickhouse-operator process which are used during reconcile cycle",
								MarkdownDescription: "runtime parameters for clickhouse-operator process which are used during reconcile cycle",
								Attributes: map[string]schema.Attribute{
									"reconcile_ch_is_threads_number": schema.Int64Attribute{
										Description:         "How many goroutines will be used to reconcile CHIs in parallel, 10 by default",
										MarkdownDescription: "How many goroutines will be used to reconcile CHIs in parallel, 10 by default",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
											int64validator.AtMost(65535),
										},
									},

									"reconcile_shards_max_concurrency_percent": schema.Int64Attribute{
										Description:         "The maximum percentage of cluster shards that may be reconciled in parallel, 50 percent by default.",
										MarkdownDescription: "The maximum percentage of cluster shards that may be reconciled in parallel, 50 percent by default.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
											int64validator.AtMost(100),
										},
									},

									"reconcile_shards_threads_number": schema.Int64Attribute{
										Description:         "How many goroutines will be used to reconcile shards of a cluster in parallel, 1 by default",
										MarkdownDescription: "How many goroutines will be used to reconcile shards of a cluster in parallel, 1 by default",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
											int64validator.AtMost(65535),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"stateful_set": schema.SingleNestedAttribute{
								Description:         "Allow change default behavior for reconciling StatefulSet which generated by clickhouse-operator",
								MarkdownDescription: "Allow change default behavior for reconciling StatefulSet which generated by clickhouse-operator",
								Attributes: map[string]schema.Attribute{
									"create": schema.SingleNestedAttribute{
										Description:         "Behavior during create StatefulSet",
										MarkdownDescription: "Behavior during create StatefulSet",
										Attributes: map[string]schema.Attribute{
											"on_failure": schema.StringAttribute{
												Description:         "What to do in case created StatefulSet is not in Ready after 'statefulSetUpdateTimeout' seconds Possible options: 1. abort - do nothing, just break the process and wait for admin. 2. delete - delete newly created problematic StatefulSet. 3. ignore (default) - ignore error, pretend nothing happened and move on to the next StatefulSet. ",
												MarkdownDescription: "What to do in case created StatefulSet is not in Ready after 'statefulSetUpdateTimeout' seconds Possible options: 1. abort - do nothing, just break the process and wait for admin. 2. delete - delete newly created problematic StatefulSet. 3. ignore (default) - ignore error, pretend nothing happened and move on to the next StatefulSet. ",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"update": schema.SingleNestedAttribute{
										Description:         "Behavior during update StatefulSet",
										MarkdownDescription: "Behavior during update StatefulSet",
										Attributes: map[string]schema.Attribute{
											"on_failure": schema.StringAttribute{
												Description:         "What to do in case updated StatefulSet is not in Ready after 'statefulSetUpdateTimeout' seconds Possible options: 1. abort - do nothing, just break the process and wait for admin. 2. rollback (default) - delete Pod and rollback StatefulSet to previous Generation. Pod would be recreated by StatefulSet based on rollback-ed configuration. 3. ignore - ignore error, pretend nothing happened and move on to the next StatefulSet. ",
												MarkdownDescription: "What to do in case updated StatefulSet is not in Ready after 'statefulSetUpdateTimeout' seconds Possible options: 1. abort - do nothing, just break the process and wait for admin. 2. rollback (default) - delete Pod and rollback StatefulSet to previous Generation. Pod would be recreated by StatefulSet based on rollback-ed configuration. 3. ignore - ignore error, pretend nothing happened and move on to the next StatefulSet. ",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"poll_interval": schema.Int64Attribute{
												Description:         "How many seconds to wait between checks for created/updated StatefulSet status",
												MarkdownDescription: "How many seconds to wait between checks for created/updated StatefulSet status",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"timeout": schema.Int64Attribute{
												Description:         "How many seconds to wait for created/updated StatefulSet to be Ready",
												MarkdownDescription: "How many seconds to wait for created/updated StatefulSet to be Ready",
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
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"stateful_set": schema.SingleNestedAttribute{
						Description:         "define StatefulSet-specific parameters",
						MarkdownDescription: "define StatefulSet-specific parameters",
						Attributes: map[string]schema.Attribute{
							"revision_history_limit": schema.Int64Attribute{
								Description:         "revisionHistoryLimit is the maximum number of revisions that will be maintained in the StatefulSet's revision history. Look details in 'statefulset.spec.revisionHistoryLimit' ",
								MarkdownDescription: "revisionHistoryLimit is the maximum number of revisions that will be maintained in the StatefulSet's revision history. Look details in 'statefulset.spec.revisionHistoryLimit' ",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"status": schema.SingleNestedAttribute{
						Description:         "defines status options",
						MarkdownDescription: "defines status options",
						Attributes: map[string]schema.Attribute{
							"fields": schema.SingleNestedAttribute{
								Description:         "defines status fields options",
								MarkdownDescription: "defines status fields options",
								Attributes: map[string]schema.Attribute{
									"action": schema.StringAttribute{
										Description:         "Whether the operator should fill status field 'action'",
										MarkdownDescription: "Whether the operator should fill status field 'action'",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("", "0", "1", "False", "false", "True", "true", "No", "no", "Yes", "yes", "Off", "off", "On", "on", "Disable", "disable", "Enable", "enable", "Disabled", "disabled", "Enabled", "enabled"),
										},
									},

									"actions": schema.StringAttribute{
										Description:         "Whether the operator should fill status field 'actions'",
										MarkdownDescription: "Whether the operator should fill status field 'actions'",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("", "0", "1", "False", "false", "True", "true", "No", "no", "Yes", "yes", "Off", "off", "On", "on", "Disable", "disable", "Enable", "enable", "Disabled", "disabled", "Enabled", "enabled"),
										},
									},

									"error": schema.StringAttribute{
										Description:         "Whether the operator should fill status field 'error'",
										MarkdownDescription: "Whether the operator should fill status field 'error'",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("", "0", "1", "False", "false", "True", "true", "No", "no", "Yes", "yes", "Off", "off", "On", "on", "Disable", "disable", "Enable", "enable", "Disabled", "disabled", "Enabled", "enabled"),
										},
									},

									"errors": schema.StringAttribute{
										Description:         "Whether the operator should fill status field 'errors'",
										MarkdownDescription: "Whether the operator should fill status field 'errors'",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("", "0", "1", "False", "false", "True", "true", "No", "no", "Yes", "yes", "Off", "off", "On", "on", "Disable", "disable", "Enable", "enable", "Disabled", "disabled", "Enabled", "enabled"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"template": schema.SingleNestedAttribute{
						Description:         "Parameters which are used if you want to generate ClickHouseInstallationTemplate custom resources from files which are stored inside clickhouse-operator deployment",
						MarkdownDescription: "Parameters which are used if you want to generate ClickHouseInstallationTemplate custom resources from files which are stored inside clickhouse-operator deployment",
						Attributes: map[string]schema.Attribute{
							"chi": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"path": schema.StringAttribute{
										Description:         "Path to folder where ClickHouseInstallationTemplate .yaml manifests are located.",
										MarkdownDescription: "Path to folder where ClickHouseInstallationTemplate .yaml manifests are located.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"policy": schema.StringAttribute{
										Description:         "CHI template updates handling policy Possible policy values: - ReadOnStart. Accept CHIT updates on the operators start only. - ApplyOnNextReconcile. Accept CHIT updates at all time. Apply news CHITs on next regular reconcile of the CHI ",
										MarkdownDescription: "CHI template updates handling policy Possible policy values: - ReadOnStart. Accept CHIT updates on the operators start only. - ApplyOnNextReconcile. Accept CHIT updates at all time. Apply news CHITs on next regular reconcile of the CHI ",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("", "ReadOnStart", "ApplyOnNextReconcile"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"watch": schema.SingleNestedAttribute{
						Description:         "Parameters for watch kubernetes resources which used by clickhouse-operator deployment",
						MarkdownDescription: "Parameters for watch kubernetes resources which used by clickhouse-operator deployment",
						Attributes: map[string]schema.Attribute{
							"namespaces": schema.ListAttribute{
								Description:         "List of namespaces where clickhouse-operator watches for events.",
								MarkdownDescription: "List of namespaces where clickhouse-operator watches for events.",
								ElementType:         types.StringType,
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *ClickhouseAltinityComClickHouseOperatorConfigurationV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_clickhouse_altinity_com_click_house_operator_configuration_v1_manifest")

	var model ClickhouseAltinityComClickHouseOperatorConfigurationV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("clickhouse.altinity.com/v1")
	model.Kind = pointer.String("ClickHouseOperatorConfiguration")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
