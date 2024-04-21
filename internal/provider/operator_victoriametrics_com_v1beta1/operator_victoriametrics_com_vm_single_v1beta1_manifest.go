/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package operator_victoriametrics_com_v1beta1

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
	_ datasource.DataSource = &OperatorVictoriametricsComVmsingleV1Beta1Manifest{}
)

func NewOperatorVictoriametricsComVmsingleV1Beta1Manifest() datasource.DataSource {
	return &OperatorVictoriametricsComVmsingleV1Beta1Manifest{}
}

type OperatorVictoriametricsComVmsingleV1Beta1Manifest struct{}

type OperatorVictoriametricsComVmsingleV1Beta1ManifestData struct {
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
		Affinity   *map[string]string   `tfsdk:"affinity" json:"affinity,omitempty"`
		ConfigMaps *[]string            `tfsdk:"config_maps" json:"configMaps,omitempty"`
		Containers *[]map[string]string `tfsdk:"containers" json:"containers,omitempty"`
		DnsConfig  *struct {
			Nameservers *[]string `tfsdk:"nameservers" json:"nameservers,omitempty"`
			Options     *[]struct {
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"options" json:"options,omitempty"`
			Searches *[]string `tfsdk:"searches" json:"searches,omitempty"`
		} `tfsdk:"dns_config" json:"dnsConfig,omitempty"`
		DnsPolicy   *string              `tfsdk:"dns_policy" json:"dnsPolicy,omitempty"`
		ExtraArgs   *map[string]string   `tfsdk:"extra_args" json:"extraArgs,omitempty"`
		ExtraEnvs   *[]map[string]string `tfsdk:"extra_envs" json:"extraEnvs,omitempty"`
		HostAliases *[]struct {
			Hostnames *[]string `tfsdk:"hostnames" json:"hostnames,omitempty"`
			Ip        *string   `tfsdk:"ip" json:"ip,omitempty"`
		} `tfsdk:"host_aliases" json:"hostAliases,omitempty"`
		HostNetwork *bool `tfsdk:"host_network" json:"hostNetwork,omitempty"`
		Image       *struct {
			PullPolicy *string `tfsdk:"pull_policy" json:"pullPolicy,omitempty"`
			Repository *string `tfsdk:"repository" json:"repository,omitempty"`
			Tag        *string `tfsdk:"tag" json:"tag,omitempty"`
		} `tfsdk:"image" json:"image,omitempty"`
		ImagePullSecrets *[]struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
		InitContainers *[]map[string]string `tfsdk:"init_containers" json:"initContainers,omitempty"`
		InsertPorts    *struct {
			GraphitePort     *string `tfsdk:"graphite_port" json:"graphitePort,omitempty"`
			InfluxPort       *string `tfsdk:"influx_port" json:"influxPort,omitempty"`
			OpenTSDBHTTPPort *string `tfsdk:"open_tsdbhttp_port" json:"openTSDBHTTPPort,omitempty"`
			OpenTSDBPort     *string `tfsdk:"open_tsdb_port" json:"openTSDBPort,omitempty"`
		} `tfsdk:"insert_ports" json:"insertPorts,omitempty"`
		License *struct {
			Key    *string `tfsdk:"key" json:"key,omitempty"`
			KeyRef *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"key_ref" json:"keyRef,omitempty"`
		} `tfsdk:"license" json:"license,omitempty"`
		LivenessProbe *map[string]string `tfsdk:"liveness_probe" json:"livenessProbe,omitempty"`
		LogFormat     *string            `tfsdk:"log_format" json:"logFormat,omitempty"`
		LogLevel      *string            `tfsdk:"log_level" json:"logLevel,omitempty"`
		NodeSelector  *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
		PodMetadata   *struct {
			Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Name        *string            `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"pod_metadata" json:"podMetadata,omitempty"`
		Port              *string `tfsdk:"port" json:"port,omitempty"`
		PriorityClassName *string `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
		ReadinessGates    *[]struct {
			ConditionType *string `tfsdk:"condition_type" json:"conditionType,omitempty"`
		} `tfsdk:"readiness_gates" json:"readinessGates,omitempty"`
		ReadinessProbe       *map[string]string `tfsdk:"readiness_probe" json:"readinessProbe,omitempty"`
		RemovePvcAfterDelete *bool              `tfsdk:"remove_pvc_after_delete" json:"removePvcAfterDelete,omitempty"`
		ReplicaCount         *int64             `tfsdk:"replica_count" json:"replicaCount,omitempty"`
		Resources            *struct {
			Claims *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"claims" json:"claims,omitempty"`
			Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
			Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"resources" json:"resources,omitempty"`
		RetentionPeriod           *string            `tfsdk:"retention_period" json:"retentionPeriod,omitempty"`
		RevisionHistoryLimitCount *int64             `tfsdk:"revision_history_limit_count" json:"revisionHistoryLimitCount,omitempty"`
		RuntimeClassName          *string            `tfsdk:"runtime_class_name" json:"runtimeClassName,omitempty"`
		SchedulerName             *string            `tfsdk:"scheduler_name" json:"schedulerName,omitempty"`
		Secrets                   *[]string          `tfsdk:"secrets" json:"secrets,omitempty"`
		SecurityContext           *map[string]string `tfsdk:"security_context" json:"securityContext,omitempty"`
		ServiceAccountName        *string            `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
		ServiceScrapeSpec         *map[string]string `tfsdk:"service_scrape_spec" json:"serviceScrapeSpec,omitempty"`
		ServiceSpec               *struct {
			Metadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				Name        *string            `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
			Spec         *map[string]string `tfsdk:"spec" json:"spec,omitempty"`
			UseAsDefault *bool              `tfsdk:"use_as_default" json:"useAsDefault,omitempty"`
		} `tfsdk:"service_spec" json:"serviceSpec,omitempty"`
		StartupProbe *map[string]string `tfsdk:"startup_probe" json:"startupProbe,omitempty"`
		Storage      *struct {
			AccessModes *[]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
			DataSource  *struct {
				ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
				Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"data_source" json:"dataSource,omitempty"`
			DataSourceRef *struct {
				ApiGroup  *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
				Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"data_source_ref" json:"dataSourceRef,omitempty"`
			Resources *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			Selector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"selector" json:"selector,omitempty"`
			StorageClassName *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
			VolumeMode       *string `tfsdk:"volume_mode" json:"volumeMode,omitempty"`
			VolumeName       *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
		} `tfsdk:"storage" json:"storage,omitempty"`
		StorageDataPath *string `tfsdk:"storage_data_path" json:"storageDataPath,omitempty"`
		StorageMetadata *struct {
			Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Name        *string            `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"storage_metadata" json:"storageMetadata,omitempty"`
		StreamAggrConfig *struct {
			DedupInterval *string `tfsdk:"dedup_interval" json:"dedupInterval,omitempty"`
			DropInput     *bool   `tfsdk:"drop_input" json:"dropInput,omitempty"`
			KeepInput     *bool   `tfsdk:"keep_input" json:"keepInput,omitempty"`
			Rules         *[]struct {
				By                    *[]string `tfsdk:"by" json:"by,omitempty"`
				Flush_on_shutdown     *bool     `tfsdk:"flush_on_shutdown" json:"flush_on_shutdown,omitempty"`
				Input_relabel_configs *[]struct {
					Action       *string            `tfsdk:"action" json:"action,omitempty"`
					If           *map[string]string `tfsdk:"if" json:"if,omitempty"`
					Labels       *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Match        *string            `tfsdk:"match" json:"match,omitempty"`
					Modulus      *int64             `tfsdk:"modulus" json:"modulus,omitempty"`
					Regex        *map[string]string `tfsdk:"regex" json:"regex,omitempty"`
					Replacement  *string            `tfsdk:"replacement" json:"replacement,omitempty"`
					Separator    *string            `tfsdk:"separator" json:"separator,omitempty"`
					SourceLabels *[]string          `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
					TargetLabel  *string            `tfsdk:"target_label" json:"targetLabel,omitempty"`
				} `tfsdk:"input_relabel_configs" json:"input_relabel_configs,omitempty"`
				Interval               *string            `tfsdk:"interval" json:"interval,omitempty"`
				Match                  *map[string]string `tfsdk:"match" json:"match,omitempty"`
				Output_relabel_configs *[]struct {
					Action       *string            `tfsdk:"action" json:"action,omitempty"`
					If           *map[string]string `tfsdk:"if" json:"if,omitempty"`
					Labels       *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Match        *string            `tfsdk:"match" json:"match,omitempty"`
					Modulus      *int64             `tfsdk:"modulus" json:"modulus,omitempty"`
					Regex        *map[string]string `tfsdk:"regex" json:"regex,omitempty"`
					Replacement  *string            `tfsdk:"replacement" json:"replacement,omitempty"`
					Separator    *string            `tfsdk:"separator" json:"separator,omitempty"`
					SourceLabels *[]string          `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
					TargetLabel  *string            `tfsdk:"target_label" json:"targetLabel,omitempty"`
				} `tfsdk:"output_relabel_configs" json:"output_relabel_configs,omitempty"`
				Outputs            *[]string `tfsdk:"outputs" json:"outputs,omitempty"`
				Staleness_interval *string   `tfsdk:"staleness_interval" json:"staleness_interval,omitempty"`
				Without            *[]string `tfsdk:"without" json:"without,omitempty"`
			} `tfsdk:"rules" json:"rules,omitempty"`
		} `tfsdk:"stream_aggr_config" json:"streamAggrConfig,omitempty"`
		TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" json:"terminationGracePeriodSeconds,omitempty"`
		Tolerations                   *[]struct {
			Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
			Key               *string `tfsdk:"key" json:"key,omitempty"`
			Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
			TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
			Value             *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tolerations" json:"tolerations,omitempty"`
		TopologySpreadConstraints *[]map[string]string `tfsdk:"topology_spread_constraints" json:"topologySpreadConstraints,omitempty"`
		UseStrictSecurity         *bool                `tfsdk:"use_strict_security" json:"useStrictSecurity,omitempty"`
		VmBackup                  *struct {
			AcceptEULA        *bool  `tfsdk:"accept_eula" json:"acceptEULA,omitempty"`
			Concurrency       *int64 `tfsdk:"concurrency" json:"concurrency,omitempty"`
			CredentialsSecret *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"credentials_secret" json:"credentialsSecret,omitempty"`
			CustomS3Endpoint            *string            `tfsdk:"custom_s3_endpoint" json:"customS3Endpoint,omitempty"`
			Destination                 *string            `tfsdk:"destination" json:"destination,omitempty"`
			DestinationDisableSuffixAdd *bool              `tfsdk:"destination_disable_suffix_add" json:"destinationDisableSuffixAdd,omitempty"`
			DisableDaily                *bool              `tfsdk:"disable_daily" json:"disableDaily,omitempty"`
			DisableHourly               *bool              `tfsdk:"disable_hourly" json:"disableHourly,omitempty"`
			DisableMonthly              *bool              `tfsdk:"disable_monthly" json:"disableMonthly,omitempty"`
			DisableWeekly               *bool              `tfsdk:"disable_weekly" json:"disableWeekly,omitempty"`
			ExtraArgs                   *map[string]string `tfsdk:"extra_args" json:"extraArgs,omitempty"`
			ExtraEnvs                   *[]struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					FieldRef *struct {
						ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
						FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
					} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
					ResourceFieldRef *struct {
						ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
						Divisor       *string `tfsdk:"divisor" json:"divisor,omitempty"`
						Resource      *string `tfsdk:"resource" json:"resource,omitempty"`
					} `tfsdk:"resource_field_ref" json:"resourceFieldRef,omitempty"`
					SecretKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"extra_envs" json:"extraEnvs,omitempty"`
			Image *struct {
				PullPolicy *string `tfsdk:"pull_policy" json:"pullPolicy,omitempty"`
				Repository *string `tfsdk:"repository" json:"repository,omitempty"`
				Tag        *string `tfsdk:"tag" json:"tag,omitempty"`
			} `tfsdk:"image" json:"image,omitempty"`
			LogFormat *string `tfsdk:"log_format" json:"logFormat,omitempty"`
			LogLevel  *string `tfsdk:"log_level" json:"logLevel,omitempty"`
			Port      *string `tfsdk:"port" json:"port,omitempty"`
			Resources *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			Restore *struct {
				OnStart *struct {
					Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
				} `tfsdk:"on_start" json:"onStart,omitempty"`
			} `tfsdk:"restore" json:"restore,omitempty"`
			SnapshotCreateURL *string `tfsdk:"snapshot_create_url" json:"snapshotCreateURL,omitempty"`
			SnapshotDeleteURL *string `tfsdk:"snapshot_delete_url" json:"snapshotDeleteURL,omitempty"`
			VolumeMounts      *[]struct {
				MountPath        *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
				MountPropagation *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
				Name             *string `tfsdk:"name" json:"name,omitempty"`
				ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
				SubPath          *string `tfsdk:"sub_path" json:"subPath,omitempty"`
				SubPathExpr      *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
			} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
		} `tfsdk:"vm_backup" json:"vmBackup,omitempty"`
		VolumeMounts *[]struct {
			MountPath        *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
			MountPropagation *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
			Name             *string `tfsdk:"name" json:"name,omitempty"`
			ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
			SubPath          *string `tfsdk:"sub_path" json:"subPath,omitempty"`
			SubPathExpr      *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
		} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
		Volumes *[]map[string]string `tfsdk:"volumes" json:"volumes,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OperatorVictoriametricsComVmsingleV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operator_victoriametrics_com_vm_single_v1beta1_manifest"
}

func (r *OperatorVictoriametricsComVmsingleV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VMSingle  is fast, cost-effective and scalable time-series database.",
		MarkdownDescription: "VMSingle  is fast, cost-effective and scalable time-series database.",
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
				Description:         "VMSingleSpec defines the desired state of VMSingle",
				MarkdownDescription: "VMSingleSpec defines the desired state of VMSingle",
				Attributes: map[string]schema.Attribute{
					"affinity": schema.MapAttribute{
						Description:         "Affinity If specified, the pod's scheduling constraints.",
						MarkdownDescription: "Affinity If specified, the pod's scheduling constraints.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"config_maps": schema.ListAttribute{
						Description:         "ConfigMaps is a list of ConfigMaps in the same namespace as the VMSingleobject, which shall be mounted into the VMSingle Pods.",
						MarkdownDescription: "ConfigMaps is a list of ConfigMaps in the same namespace as the VMSingleobject, which shall be mounted into the VMSingle Pods.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"containers": schema.ListAttribute{
						Description:         "Containers property allows to inject additions sidecars or to patch existing containers.It can be useful for proxies, backup, etc.",
						MarkdownDescription: "Containers property allows to inject additions sidecars or to patch existing containers.It can be useful for proxies, backup, etc.",
						ElementType:         types.MapType{ElemType: types.StringType},
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"dns_config": schema.SingleNestedAttribute{
						Description:         "Specifies the DNS parameters of a pod.Parameters specified here will be merged to the generated DNSconfiguration based on DNSPolicy.",
						MarkdownDescription: "Specifies the DNS parameters of a pod.Parameters specified here will be merged to the generated DNSconfiguration based on DNSPolicy.",
						Attributes: map[string]schema.Attribute{
							"nameservers": schema.ListAttribute{
								Description:         "A list of DNS name server IP addresses.This will be appended to the base nameservers generated from DNSPolicy.Duplicated nameservers will be removed.",
								MarkdownDescription: "A list of DNS name server IP addresses.This will be appended to the base nameservers generated from DNSPolicy.Duplicated nameservers will be removed.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"options": schema.ListNestedAttribute{
								Description:         "A list of DNS resolver options.This will be merged with the base options generated from DNSPolicy.Duplicated entries will be removed. Resolution options given in Optionswill override those that appear in the base DNSPolicy.",
								MarkdownDescription: "A list of DNS resolver options.This will be merged with the base options generated from DNSPolicy.Duplicated entries will be removed. Resolution options given in Optionswill override those that appear in the base DNSPolicy.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Required.",
											MarkdownDescription: "Required.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
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

							"searches": schema.ListAttribute{
								Description:         "A list of DNS search domains for host-name lookup.This will be appended to the base search paths generated from DNSPolicy.Duplicated search paths will be removed.",
								MarkdownDescription: "A list of DNS search domains for host-name lookup.This will be appended to the base search paths generated from DNSPolicy.Duplicated search paths will be removed.",
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

					"dns_policy": schema.StringAttribute{
						Description:         "DNSPolicy sets DNS policy for the pod",
						MarkdownDescription: "DNSPolicy sets DNS policy for the pod",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"extra_args": schema.MapAttribute{
						Description:         "ExtraArgs that will be passed to  VMSingle podfor example remoteWrite.tmpDataPath: /tmp",
						MarkdownDescription: "ExtraArgs that will be passed to  VMSingle podfor example remoteWrite.tmpDataPath: /tmp",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"extra_envs": schema.ListAttribute{
						Description:         "ExtraEnvs that will be added to VMSingle pod",
						MarkdownDescription: "ExtraEnvs that will be added to VMSingle pod",
						ElementType:         types.MapType{ElemType: types.StringType},
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"host_aliases": schema.ListNestedAttribute{
						Description:         "HostAliases provides mapping for ip and hostname,that would be propagated to pod,cannot be used with HostNetwork.",
						MarkdownDescription: "HostAliases provides mapping for ip and hostname,that would be propagated to pod,cannot be used with HostNetwork.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"hostnames": schema.ListAttribute{
									Description:         "Hostnames for the above IP address.",
									MarkdownDescription: "Hostnames for the above IP address.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ip": schema.StringAttribute{
									Description:         "IP address of the host file entry.",
									MarkdownDescription: "IP address of the host file entry.",
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

					"host_network": schema.BoolAttribute{
						Description:         "HostNetwork controls whether the pod may use the node network namespace",
						MarkdownDescription: "HostNetwork controls whether the pod may use the node network namespace",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image": schema.SingleNestedAttribute{
						Description:         "Image - docker image settings for VMSingleif no specified operator uses default config version",
						MarkdownDescription: "Image - docker image settings for VMSingleif no specified operator uses default config version",
						Attributes: map[string]schema.Attribute{
							"pull_policy": schema.StringAttribute{
								Description:         "PullPolicy describes how to pull docker image",
								MarkdownDescription: "PullPolicy describes how to pull docker image",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"repository": schema.StringAttribute{
								Description:         "Repository contains name of docker image + it's repository if needed",
								MarkdownDescription: "Repository contains name of docker image + it's repository if needed",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tag": schema.StringAttribute{
								Description:         "Tag contains desired docker image version",
								MarkdownDescription: "Tag contains desired docker image version",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"image_pull_secrets": schema.ListNestedAttribute{
						Description:         "ImagePullSecrets An optional list of references to secrets in the same namespaceto use for pulling images from registriessee https://kubernetes.io/docs/concepts/containers/images/#referring-to-an-imagepullsecrets-on-a-pod",
						MarkdownDescription: "ImagePullSecrets An optional list of references to secrets in the same namespaceto use for pulling images from registriessee https://kubernetes.io/docs/concepts/containers/images/#referring-to-an-imagepullsecrets-on-a-pod",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
									MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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

					"init_containers": schema.ListAttribute{
						Description:         "InitContainers allows adding initContainers to the pod definition. Those can be used to e.g.fetch secrets for injection into the vmSingle configuration from external sources. Anyerrors during the execution of an initContainer will lead to a restart of the Pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/Using initContainers for any use case other then secret fetching is entirely outside the scopeof what the maintainers will support and by doing so, you accept that this behaviour may breakat any time without notice.",
						MarkdownDescription: "InitContainers allows adding initContainers to the pod definition. Those can be used to e.g.fetch secrets for injection into the vmSingle configuration from external sources. Anyerrors during the execution of an initContainer will lead to a restart of the Pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/Using initContainers for any use case other then secret fetching is entirely outside the scopeof what the maintainers will support and by doing so, you accept that this behaviour may breakat any time without notice.",
						ElementType:         types.MapType{ElemType: types.StringType},
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"insert_ports": schema.SingleNestedAttribute{
						Description:         "InsertPorts - additional listen ports for data ingestion.",
						MarkdownDescription: "InsertPorts - additional listen ports for data ingestion.",
						Attributes: map[string]schema.Attribute{
							"graphite_port": schema.StringAttribute{
								Description:         "GraphitePort listen port",
								MarkdownDescription: "GraphitePort listen port",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"influx_port": schema.StringAttribute{
								Description:         "InfluxPort listen port",
								MarkdownDescription: "InfluxPort listen port",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"open_tsdbhttp_port": schema.StringAttribute{
								Description:         "OpenTSDBHTTPPort for http connections.",
								MarkdownDescription: "OpenTSDBHTTPPort for http connections.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"open_tsdb_port": schema.StringAttribute{
								Description:         "OpenTSDBPort for tcp and udp listen",
								MarkdownDescription: "OpenTSDBPort for tcp and udp listen",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"license": schema.SingleNestedAttribute{
						Description:         "License allows to configure license key to be used for enterprise features.Using license key is supported starting from VictoriaMetrics v1.94.0.See: https://docs.victoriametrics.com/enterprise.html",
						MarkdownDescription: "License allows to configure license key to be used for enterprise features.Using license key is supported starting from VictoriaMetrics v1.94.0.See: https://docs.victoriametrics.com/enterprise.html",
						Attributes: map[string]schema.Attribute{
							"key": schema.StringAttribute{
								Description:         "Enterprise license key. This flag is available only in VictoriaMetrics enterprise.Documentation - https://docs.victoriametrics.com/enterprise.htmlfor more information, visit https://victoriametrics.com/products/enterprise/ .To request a trial license, go to https://victoriametrics.com/products/enterprise/trial/",
								MarkdownDescription: "Enterprise license key. This flag is available only in VictoriaMetrics enterprise.Documentation - https://docs.victoriametrics.com/enterprise.htmlfor more information, visit https://victoriametrics.com/products/enterprise/ .To request a trial license, go to https://victoriametrics.com/products/enterprise/trial/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"key_ref": schema.SingleNestedAttribute{
								Description:         "KeyRef is reference to secret with license key for enterprise features.",
								MarkdownDescription: "KeyRef is reference to secret with license key for enterprise features.",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "The key of the secret to select from.  Must be a valid secret key.",
										MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"optional": schema.BoolAttribute{
										Description:         "Specify whether the Secret or its key must be defined",
										MarkdownDescription: "Specify whether the Secret or its key must be defined",
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

					"liveness_probe": schema.MapAttribute{
						Description:         "LivenessProbe that will be added CRD pod",
						MarkdownDescription: "LivenessProbe that will be added CRD pod",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"log_format": schema.StringAttribute{
						Description:         "LogFormat for VMSingle to be configured with.",
						MarkdownDescription: "LogFormat for VMSingle to be configured with.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("default", "json"),
						},
					},

					"log_level": schema.StringAttribute{
						Description:         "LogLevel for victoria metrics single to be configured with.",
						MarkdownDescription: "LogLevel for victoria metrics single to be configured with.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("INFO", "WARN", "ERROR", "FATAL", "PANIC"),
						},
					},

					"node_selector": schema.MapAttribute{
						Description:         "NodeSelector Define which Nodes the Pods are scheduled on.",
						MarkdownDescription: "NodeSelector Define which Nodes the Pods are scheduled on.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"pod_metadata": schema.SingleNestedAttribute{
						Description:         "PodMetadata configures Labels and Annotations which are propagated to the VMSingle pods.",
						MarkdownDescription: "PodMetadata configures Labels and Annotations which are propagated to the VMSingle pods.",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
								MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"labels": schema.MapAttribute{
								Description:         "Labels Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
								MarkdownDescription: "Labels Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name must be unique within a namespace. Is required when creating resources, althoughsome resources may allow a client to request the generation of an appropriate nameautomatically. Name is primarily intended for creation idempotence and configurationdefinition.Cannot be updated.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
								MarkdownDescription: "Name must be unique within a namespace. Is required when creating resources, althoughsome resources may allow a client to request the generation of an appropriate nameautomatically. Name is primarily intended for creation idempotence and configurationdefinition.Cannot be updated.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"port": schema.StringAttribute{
						Description:         "Port listen port",
						MarkdownDescription: "Port listen port",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"priority_class_name": schema.StringAttribute{
						Description:         "PriorityClassName assigned to the Pods",
						MarkdownDescription: "PriorityClassName assigned to the Pods",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"readiness_gates": schema.ListNestedAttribute{
						Description:         "ReadinessGates defines pod readiness gates",
						MarkdownDescription: "ReadinessGates defines pod readiness gates",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"condition_type": schema.StringAttribute{
									Description:         "ConditionType refers to a condition in the pod's condition list with matching type.",
									MarkdownDescription: "ConditionType refers to a condition in the pod's condition list with matching type.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"readiness_probe": schema.MapAttribute{
						Description:         "ReadinessProbe that will be added CRD pod",
						MarkdownDescription: "ReadinessProbe that will be added CRD pod",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"remove_pvc_after_delete": schema.BoolAttribute{
						Description:         "RemovePvcAfterDelete - if true, controller adds ownership to pvcand after VMSingle objest deletion - pvc will be garbage collectedby controller manager",
						MarkdownDescription: "RemovePvcAfterDelete - if true, controller adds ownership to pvcand after VMSingle objest deletion - pvc will be garbage collectedby controller manager",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"replica_count": schema.Int64Attribute{
						Description:         "ReplicaCount is the expected size of the VMSingleit can be 0 or 1if you need more - use vm cluster",
						MarkdownDescription: "ReplicaCount is the expected size of the VMSingleit can be 0 or 1if you need more - use vm cluster",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"resources": schema.SingleNestedAttribute{
						Description:         "Resources container resource request and limits, https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/if not defined default resources from operator config will be used",
						MarkdownDescription: "Resources container resource request and limits, https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/if not defined default resources from operator config will be used",
						Attributes: map[string]schema.Attribute{
							"claims": schema.ListNestedAttribute{
								Description:         "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
								MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
											MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"limits": schema.MapAttribute{
								Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"requests": schema.MapAttribute{
								Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

					"retention_period": schema.StringAttribute{
						Description:         "RetentionPeriod for the stored metricsNote VictoriaMetrics has data/ and indexdb/ foldersmetrics from data/ removed eventually as soon as partition leaves retention periodreverse index data at indexdb rotates once at the half of configured retention periodhttps://docs.victoriametrics.com/Single-server-VictoriaMetrics.html#retention",
						MarkdownDescription: "RetentionPeriod for the stored metricsNote VictoriaMetrics has data/ and indexdb/ foldersmetrics from data/ removed eventually as soon as partition leaves retention periodreverse index data at indexdb rotates once at the half of configured retention periodhttps://docs.victoriametrics.com/Single-server-VictoriaMetrics.html#retention",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"revision_history_limit_count": schema.Int64Attribute{
						Description:         "The number of old ReplicaSets to retain to allow rollback in deployment ormaximum number of revisions that will be maintained in the StatefulSet's revision history.Defaults to 10.",
						MarkdownDescription: "The number of old ReplicaSets to retain to allow rollback in deployment ormaximum number of revisions that will be maintained in the StatefulSet's revision history.Defaults to 10.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"runtime_class_name": schema.StringAttribute{
						Description:         "RuntimeClassName - defines runtime class for kubernetes pod.https://kubernetes.io/docs/concepts/containers/runtime-class/",
						MarkdownDescription: "RuntimeClassName - defines runtime class for kubernetes pod.https://kubernetes.io/docs/concepts/containers/runtime-class/",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"scheduler_name": schema.StringAttribute{
						Description:         "SchedulerName - defines kubernetes scheduler name",
						MarkdownDescription: "SchedulerName - defines kubernetes scheduler name",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"secrets": schema.ListAttribute{
						Description:         "Secrets is a list of Secrets in the same namespace as the VMSingleobject, which shall be mounted into the VMSingle Pods.",
						MarkdownDescription: "Secrets is a list of Secrets in the same namespace as the VMSingleobject, which shall be mounted into the VMSingle Pods.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"security_context": schema.MapAttribute{
						Description:         "SecurityContext holds pod-level security attributes and common container settings.This defaults to the default PodSecurityContext.",
						MarkdownDescription: "SecurityContext holds pod-level security attributes and common container settings.This defaults to the default PodSecurityContext.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service_account_name": schema.StringAttribute{
						Description:         "ServiceAccountName is the name of the ServiceAccount to use to run theVMSingle Pods.",
						MarkdownDescription: "ServiceAccountName is the name of the ServiceAccount to use to run theVMSingle Pods.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service_scrape_spec": schema.MapAttribute{
						Description:         "ServiceScrapeSpec that will be added to vmsingle VMServiceScrape spec",
						MarkdownDescription: "ServiceScrapeSpec that will be added to vmsingle VMServiceScrape spec",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service_spec": schema.SingleNestedAttribute{
						Description:         "ServiceSpec that will be added to vmsingle service spec",
						MarkdownDescription: "ServiceSpec that will be added to vmsingle service spec",
						Attributes: map[string]schema.Attribute{
							"metadata": schema.SingleNestedAttribute{
								Description:         "EmbeddedObjectMetadata defines objectMeta for additional service.",
								MarkdownDescription: "EmbeddedObjectMetadata defines objectMeta for additional service.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
										MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
										Description:         "Labels Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
										MarkdownDescription: "Labels Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name must be unique within a namespace. Is required when creating resources, althoughsome resources may allow a client to request the generation of an appropriate nameautomatically. Name is primarily intended for creation idempotence and configurationdefinition.Cannot be updated.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
										MarkdownDescription: "Name must be unique within a namespace. Is required when creating resources, althoughsome resources may allow a client to request the generation of an appropriate nameautomatically. Name is primarily intended for creation idempotence and configurationdefinition.Cannot be updated.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"spec": schema.MapAttribute{
								Description:         "ServiceSpec describes the attributes that a user creates on a service.More info: https://kubernetes.io/docs/concepts/services-networking/service/",
								MarkdownDescription: "ServiceSpec describes the attributes that a user creates on a service.More info: https://kubernetes.io/docs/concepts/services-networking/service/",
								ElementType:         types.StringType,
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"use_as_default": schema.BoolAttribute{
								Description:         "UseAsDefault applies changes from given service definition to the main object ServiceChaning from headless service to clusterIP or loadbalancer may break cross-component communication",
								MarkdownDescription: "UseAsDefault applies changes from given service definition to the main object ServiceChaning from headless service to clusterIP or loadbalancer may break cross-component communication",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"startup_probe": schema.MapAttribute{
						Description:         "StartupProbe that will be added to CRD pod",
						MarkdownDescription: "StartupProbe that will be added to CRD pod",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"storage": schema.SingleNestedAttribute{
						Description:         "Storage is the definition of how storage will be used by the VMSingleby default it's empty dir",
						MarkdownDescription: "Storage is the definition of how storage will be used by the VMSingleby default it's empty dir",
						Attributes: map[string]schema.Attribute{
							"access_modes": schema.ListAttribute{
								Description:         "accessModes contains the desired access modes the volume should have.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
								MarkdownDescription: "accessModes contains the desired access modes the volume should have.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"data_source": schema.SingleNestedAttribute{
								Description:         "dataSource field can be used to specify either:* An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot)* An existing PVC (PersistentVolumeClaim)If the provisioner or an external controller can support the specified data source,it will create a new volume based on the contents of the specified data source.When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef,and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified.If the namespace is specified, then dataSourceRef will not be copied to dataSource.",
								MarkdownDescription: "dataSource field can be used to specify either:* An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot)* An existing PVC (PersistentVolumeClaim)If the provisioner or an external controller can support the specified data source,it will create a new volume based on the contents of the specified data source.When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef,and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified.If the namespace is specified, then dataSourceRef will not be copied to dataSource.",
								Attributes: map[string]schema.Attribute{
									"api_group": schema.StringAttribute{
										Description:         "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
										MarkdownDescription: "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"kind": schema.StringAttribute{
										Description:         "Kind is the type of resource being referenced",
										MarkdownDescription: "Kind is the type of resource being referenced",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name is the name of resource being referenced",
										MarkdownDescription: "Name is the name of resource being referenced",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"data_source_ref": schema.SingleNestedAttribute{
								Description:         "dataSourceRef specifies the object from which to populate the volume with data, if a non-emptyvolume is desired. This may be any object from a non-empty API group (noncore object) or a PersistentVolumeClaim object.When this field is specified, volume binding will only succeed if the type ofthe specified object matches some installed volume populator or dynamicprovisioner.This field will replace the functionality of the dataSource field and as suchif both fields are non-empty, they must have the same value. For backwardscompatibility, when namespace isn't specified in dataSourceRef,both fields (dataSource and dataSourceRef) will be set to the samevalue automatically if one of them is empty and the other is non-empty.When namespace is specified in dataSourceRef,dataSource isn't set to the same value and must be empty.There are three important differences between dataSource and dataSourceRef:* While dataSource only allows two specific types of objects, dataSourceRef  allows any non-core object, as well as PersistentVolumeClaim objects.* While dataSource ignores disallowed values (dropping them), dataSourceRef  preserves all values, and generates an error if a disallowed value is  specified.* While dataSource only allows local objects, dataSourceRef allows objects  in any namespaces.(Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.(Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
								MarkdownDescription: "dataSourceRef specifies the object from which to populate the volume with data, if a non-emptyvolume is desired. This may be any object from a non-empty API group (noncore object) or a PersistentVolumeClaim object.When this field is specified, volume binding will only succeed if the type ofthe specified object matches some installed volume populator or dynamicprovisioner.This field will replace the functionality of the dataSource field and as suchif both fields are non-empty, they must have the same value. For backwardscompatibility, when namespace isn't specified in dataSourceRef,both fields (dataSource and dataSourceRef) will be set to the samevalue automatically if one of them is empty and the other is non-empty.When namespace is specified in dataSourceRef,dataSource isn't set to the same value and must be empty.There are three important differences between dataSource and dataSourceRef:* While dataSource only allows two specific types of objects, dataSourceRef  allows any non-core object, as well as PersistentVolumeClaim objects.* While dataSource ignores disallowed values (dropping them), dataSourceRef  preserves all values, and generates an error if a disallowed value is  specified.* While dataSource only allows local objects, dataSourceRef allows objects  in any namespaces.(Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.(Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
								Attributes: map[string]schema.Attribute{
									"api_group": schema.StringAttribute{
										Description:         "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
										MarkdownDescription: "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"kind": schema.StringAttribute{
										Description:         "Kind is the type of resource being referenced",
										MarkdownDescription: "Kind is the type of resource being referenced",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name is the name of resource being referenced",
										MarkdownDescription: "Name is the name of resource being referenced",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "Namespace is the namespace of resource being referencedNote that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.(Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
										MarkdownDescription: "Namespace is the namespace of resource being referencedNote that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.(Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"resources": schema.SingleNestedAttribute{
								Description:         "resources represents the minimum resources the volume should have.If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirementsthat are lower than previous value but must still be higher than capacity recorded in thestatus field of the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
								MarkdownDescription: "resources represents the minimum resources the volume should have.If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirementsthat are lower than previous value but must still be higher than capacity recorded in thestatus field of the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
								Attributes: map[string]schema.Attribute{
									"claims": schema.ListNestedAttribute{
										Description:         "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
										MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
													MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

							"selector": schema.SingleNestedAttribute{
								Description:         "selector is a label query over volumes to consider for binding.",
								MarkdownDescription: "selector is a label query over volumes to consider for binding.",
								Attributes: map[string]schema.Attribute{
									"match_expressions": schema.ListNestedAttribute{
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "key is the label key that the selector applies to.",
													MarkdownDescription: "key is the label key that the selector applies to.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"operator": schema.StringAttribute{
													Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
													MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"values": schema.ListAttribute{
													Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
													MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
													ElementType:         types.StringType,
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

									"match_labels": schema.MapAttribute{
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

							"storage_class_name": schema.StringAttribute{
								Description:         "storageClassName is the name of the StorageClass required by the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
								MarkdownDescription: "storageClassName is the name of the StorageClass required by the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"volume_mode": schema.StringAttribute{
								Description:         "volumeMode defines what type of volume is required by the claim.Value of Filesystem is implied when not included in claim spec.",
								MarkdownDescription: "volumeMode defines what type of volume is required by the claim.Value of Filesystem is implied when not included in claim spec.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"volume_name": schema.StringAttribute{
								Description:         "volumeName is the binding reference to the PersistentVolume backing this claim.",
								MarkdownDescription: "volumeName is the binding reference to the PersistentVolume backing this claim.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"storage_data_path": schema.StringAttribute{
						Description:         "StorageDataPath disables spec.storage option and overrides arg for victoria-metrics binary --storageDataPath,its users responsibility to mount proper device into given path.",
						MarkdownDescription: "StorageDataPath disables spec.storage option and overrides arg for victoria-metrics binary --storageDataPath,its users responsibility to mount proper device into given path.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"storage_metadata": schema.SingleNestedAttribute{
						Description:         "StorageMeta defines annotations and labels attached to PVC for given vmsingle CR",
						MarkdownDescription: "StorageMeta defines annotations and labels attached to PVC for given vmsingle CR",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
								MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"labels": schema.MapAttribute{
								Description:         "Labels Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
								MarkdownDescription: "Labels Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name must be unique within a namespace. Is required when creating resources, althoughsome resources may allow a client to request the generation of an appropriate nameautomatically. Name is primarily intended for creation idempotence and configurationdefinition.Cannot be updated.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
								MarkdownDescription: "Name must be unique within a namespace. Is required when creating resources, althoughsome resources may allow a client to request the generation of an appropriate nameautomatically. Name is primarily intended for creation idempotence and configurationdefinition.Cannot be updated.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"stream_aggr_config": schema.SingleNestedAttribute{
						Description:         "StreamAggrConfig defines stream aggregation configuration for VMSingle",
						MarkdownDescription: "StreamAggrConfig defines stream aggregation configuration for VMSingle",
						Attributes: map[string]schema.Attribute{
							"dedup_interval": schema.StringAttribute{
								Description:         "Allows setting different de-duplication intervals per each configured remote storage",
								MarkdownDescription: "Allows setting different de-duplication intervals per each configured remote storage",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"drop_input": schema.BoolAttribute{
								Description:         "Allow drop all the input samples after the aggregation",
								MarkdownDescription: "Allow drop all the input samples after the aggregation",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"keep_input": schema.BoolAttribute{
								Description:         "Allows writing both raw and aggregate data",
								MarkdownDescription: "Allows writing both raw and aggregate data",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"rules": schema.ListNestedAttribute{
								Description:         "Stream aggregation rules",
								MarkdownDescription: "Stream aggregation rules",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"by": schema.ListAttribute{
											Description:         "By is an optional list of labels for grouping input series.See also Without.If neither By nor Without are set, then the Outputs are calculatedindividually per each input time series.",
											MarkdownDescription: "By is an optional list of labels for grouping input series.See also Without.If neither By nor Without are set, then the Outputs are calculatedindividually per each input time series.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"flush_on_shutdown": schema.BoolAttribute{
											Description:         "FlushOnShutdown defines whether to flush the aggregation state on process terminationor config reload. Is 'false' by default.It is not recommended changing this setting, unless unfinished aggregations statesare preferred to missing data points.",
											MarkdownDescription: "FlushOnShutdown defines whether to flush the aggregation state on process terminationor config reload. Is 'false' by default.It is not recommended changing this setting, unless unfinished aggregations statesare preferred to missing data points.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"input_relabel_configs": schema.ListNestedAttribute{
											Description:         "InputRelabelConfigs is an optional relabeling rules, which are applied on the inputbefore aggregation.",
											MarkdownDescription: "InputRelabelConfigs is an optional relabeling rules, which are applied on the inputbefore aggregation.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"action": schema.StringAttribute{
														Description:         "Action to perform based on regex matching. Default is 'replace'",
														MarkdownDescription: "Action to perform based on regex matching. Default is 'replace'",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"if": schema.MapAttribute{
														Description:         "If represents metricsQL match expression (or list of expressions): '{__name__=~'foo_.*'}'",
														MarkdownDescription: "If represents metricsQL match expression (or list of expressions): '{__name__=~'foo_.*'}'",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"labels": schema.MapAttribute{
														Description:         "Labels is used together with Match for 'action: graphite'",
														MarkdownDescription: "Labels is used together with Match for 'action: graphite'",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"match": schema.StringAttribute{
														Description:         "Match is used together with Labels for 'action: graphite'",
														MarkdownDescription: "Match is used together with Labels for 'action: graphite'",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"modulus": schema.Int64Attribute{
														Description:         "Modulus to take of the hash of the source label values.",
														MarkdownDescription: "Modulus to take of the hash of the source label values.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"regex": schema.MapAttribute{
														Description:         "Regular expression against which the extracted value is matched. Default is '(.*)'victoriaMetrics supports multiline regex joined with |https://docs.victoriametrics.com/vmagent/#relabeling-enhancements",
														MarkdownDescription: "Regular expression against which the extracted value is matched. Default is '(.*)'victoriaMetrics supports multiline regex joined with |https://docs.victoriametrics.com/vmagent/#relabeling-enhancements",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"replacement": schema.StringAttribute{
														Description:         "Replacement value against which a regex replace is performed if theregular expression matches. Regex capture groups are available. Default is '$1'",
														MarkdownDescription: "Replacement value against which a regex replace is performed if theregular expression matches. Regex capture groups are available. Default is '$1'",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"separator": schema.StringAttribute{
														Description:         "Separator placed between concatenated source label values. default is ';'.",
														MarkdownDescription: "Separator placed between concatenated source label values. default is ';'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"source_labels": schema.ListAttribute{
														Description:         "The source labels select values from existing labels. Their content is concatenatedusing the configured separator and matched against the configured regular expressionfor the replace, keep, and drop actions.",
														MarkdownDescription: "The source labels select values from existing labels. Their content is concatenatedusing the configured separator and matched against the configured regular expressionfor the replace, keep, and drop actions.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"target_label": schema.StringAttribute{
														Description:         "Label to which the resulting value is written in a replace action.It is mandatory for replace actions. Regex capture groups are available.",
														MarkdownDescription: "Label to which the resulting value is written in a replace action.It is mandatory for replace actions. Regex capture groups are available.",
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

										"interval": schema.StringAttribute{
											Description:         "Interval is the interval between aggregations.",
											MarkdownDescription: "Interval is the interval between aggregations.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"match": schema.MapAttribute{
											Description:         "Match is a label selector (or list of label selectors) for filtering time series for the given selector.If the match isn't set, then all the input time series are processed.",
											MarkdownDescription: "Match is a label selector (or list of label selectors) for filtering time series for the given selector.If the match isn't set, then all the input time series are processed.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"output_relabel_configs": schema.ListNestedAttribute{
											Description:         "OutputRelabelConfigs is an optional relabeling rules, which are appliedon the aggregated output before being sent to remote storage.",
											MarkdownDescription: "OutputRelabelConfigs is an optional relabeling rules, which are appliedon the aggregated output before being sent to remote storage.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"action": schema.StringAttribute{
														Description:         "Action to perform based on regex matching. Default is 'replace'",
														MarkdownDescription: "Action to perform based on regex matching. Default is 'replace'",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"if": schema.MapAttribute{
														Description:         "If represents metricsQL match expression (or list of expressions): '{__name__=~'foo_.*'}'",
														MarkdownDescription: "If represents metricsQL match expression (or list of expressions): '{__name__=~'foo_.*'}'",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"labels": schema.MapAttribute{
														Description:         "Labels is used together with Match for 'action: graphite'",
														MarkdownDescription: "Labels is used together with Match for 'action: graphite'",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"match": schema.StringAttribute{
														Description:         "Match is used together with Labels for 'action: graphite'",
														MarkdownDescription: "Match is used together with Labels for 'action: graphite'",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"modulus": schema.Int64Attribute{
														Description:         "Modulus to take of the hash of the source label values.",
														MarkdownDescription: "Modulus to take of the hash of the source label values.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"regex": schema.MapAttribute{
														Description:         "Regular expression against which the extracted value is matched. Default is '(.*)'victoriaMetrics supports multiline regex joined with |https://docs.victoriametrics.com/vmagent/#relabeling-enhancements",
														MarkdownDescription: "Regular expression against which the extracted value is matched. Default is '(.*)'victoriaMetrics supports multiline regex joined with |https://docs.victoriametrics.com/vmagent/#relabeling-enhancements",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"replacement": schema.StringAttribute{
														Description:         "Replacement value against which a regex replace is performed if theregular expression matches. Regex capture groups are available. Default is '$1'",
														MarkdownDescription: "Replacement value against which a regex replace is performed if theregular expression matches. Regex capture groups are available. Default is '$1'",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"separator": schema.StringAttribute{
														Description:         "Separator placed between concatenated source label values. default is ';'.",
														MarkdownDescription: "Separator placed between concatenated source label values. default is ';'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"source_labels": schema.ListAttribute{
														Description:         "The source labels select values from existing labels. Their content is concatenatedusing the configured separator and matched against the configured regular expressionfor the replace, keep, and drop actions.",
														MarkdownDescription: "The source labels select values from existing labels. Their content is concatenatedusing the configured separator and matched against the configured regular expressionfor the replace, keep, and drop actions.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"target_label": schema.StringAttribute{
														Description:         "Label to which the resulting value is written in a replace action.It is mandatory for replace actions. Regex capture groups are available.",
														MarkdownDescription: "Label to which the resulting value is written in a replace action.It is mandatory for replace actions. Regex capture groups are available.",
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

										"outputs": schema.ListAttribute{
											Description:         "Outputs is a list of output aggregate functions to produce.The following names are allowed:- total - aggregates input counters- increase - counts the increase over input counters- count_series - counts the input series- count_samples - counts the input samples- sum_samples - sums the input samples- last - the last biggest sample value- min - the minimum sample value- max - the maximum sample value- avg - the average value across all the samples- stddev - standard deviation across all the samples- stdvar - standard variance across all the samples- histogram_bucket - creates VictoriaMetrics histogram for input samples- quantiles(phi1, ..., phiN) - quantiles' estimation for phi in the range [0..1]The output time series will have the following names:  input_name:aggr_<interval>_<output>",
											MarkdownDescription: "Outputs is a list of output aggregate functions to produce.The following names are allowed:- total - aggregates input counters- increase - counts the increase over input counters- count_series - counts the input series- count_samples - counts the input samples- sum_samples - sums the input samples- last - the last biggest sample value- min - the minimum sample value- max - the maximum sample value- avg - the average value across all the samples- stddev - standard deviation across all the samples- stdvar - standard variance across all the samples- histogram_bucket - creates VictoriaMetrics histogram for input samples- quantiles(phi1, ..., phiN) - quantiles' estimation for phi in the range [0..1]The output time series will have the following names:  input_name:aggr_<interval>_<output>",
											ElementType:         types.StringType,
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"staleness_interval": schema.StringAttribute{
											Description:         "StalenessInterval defines an interval after which the series state will be reset if no samples have been sent during it.",
											MarkdownDescription: "StalenessInterval defines an interval after which the series state will be reset if no samples have been sent during it.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"without": schema.ListAttribute{
											Description:         "Without is an optional list of labels, which must be excluded when grouping input series.See also By.If neither By nor Without are set, then the Outputs are calculatedindividually per each input time series.",
											MarkdownDescription: "Without is an optional list of labels, which must be excluded when grouping input series.See also By.If neither By nor Without are set, then the Outputs are calculatedindividually per each input time series.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"termination_grace_period_seconds": schema.Int64Attribute{
						Description:         "TerminationGracePeriodSeconds period for container graceful termination",
						MarkdownDescription: "TerminationGracePeriodSeconds period for container graceful termination",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tolerations": schema.ListNestedAttribute{
						Description:         "Tolerations If specified, the pod's tolerations.",
						MarkdownDescription: "Tolerations If specified, the pod's tolerations.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"effect": schema.StringAttribute{
									Description:         "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
									MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"key": schema.StringAttribute{
									Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
									MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"operator": schema.StringAttribute{
									Description:         "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
									MarkdownDescription: "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"toleration_seconds": schema.Int64Attribute{
									Description:         "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
									MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
									MarkdownDescription: "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
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

					"topology_spread_constraints": schema.ListAttribute{
						Description:         "TopologySpreadConstraints embedded kubernetes pod configuration option,controls how pods are spread across your cluster among failure-domainssuch as regions, zones, nodes, and other user-defined topology domainshttps://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/",
						MarkdownDescription: "TopologySpreadConstraints embedded kubernetes pod configuration option,controls how pods are spread across your cluster among failure-domainssuch as regions, zones, nodes, and other user-defined topology domainshttps://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/",
						ElementType:         types.MapType{ElemType: types.StringType},
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"use_strict_security": schema.BoolAttribute{
						Description:         "UseStrictSecurity enables strict security mode for componentit restricts disk writes accessuses non-root user out of the boxdrops not needed security permissions",
						MarkdownDescription: "UseStrictSecurity enables strict security mode for componentit restricts disk writes accessuses non-root user out of the boxdrops not needed security permissions",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"vm_backup": schema.SingleNestedAttribute{
						Description:         "VMBackup configuration for backup",
						MarkdownDescription: "VMBackup configuration for backup",
						Attributes: map[string]schema.Attribute{
							"accept_eula": schema.BoolAttribute{
								Description:         "AcceptEULA accepts enterprise feature usage, must be set to true.otherwise backupmanager cannot be added to single/cluster version.https://victoriametrics.com/legal/esa/",
								MarkdownDescription: "AcceptEULA accepts enterprise feature usage, must be set to true.otherwise backupmanager cannot be added to single/cluster version.https://victoriametrics.com/legal/esa/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"concurrency": schema.Int64Attribute{
								Description:         "Defines number of concurrent workers. Higher concurrency may reduce backup duration (default 10)",
								MarkdownDescription: "Defines number of concurrent workers. Higher concurrency may reduce backup duration (default 10)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"credentials_secret": schema.SingleNestedAttribute{
								Description:         "CredentialsSecret is secret in the same namespace for access to remote storageThe secret is mounted into /etc/vm/creds.",
								MarkdownDescription: "CredentialsSecret is secret in the same namespace for access to remote storageThe secret is mounted into /etc/vm/creds.",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "The key of the secret to select from.  Must be a valid secret key.",
										MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"optional": schema.BoolAttribute{
										Description:         "Specify whether the Secret or its key must be defined",
										MarkdownDescription: "Specify whether the Secret or its key must be defined",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"custom_s3_endpoint": schema.StringAttribute{
								Description:         "Custom S3 endpoint for use with S3-compatible storages (e.g. MinIO). S3 is used if not set",
								MarkdownDescription: "Custom S3 endpoint for use with S3-compatible storages (e.g. MinIO). S3 is used if not set",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"destination": schema.StringAttribute{
								Description:         "Defines destination for backup",
								MarkdownDescription: "Defines destination for backup",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"destination_disable_suffix_add": schema.BoolAttribute{
								Description:         "DestinationDisableSuffixAdd - disables suffix adding for cluster version backupseach vmstorage backup must have unique backup folderso operator adds POD_NAME as suffix for backup destination folder.",
								MarkdownDescription: "DestinationDisableSuffixAdd - disables suffix adding for cluster version backupseach vmstorage backup must have unique backup folderso operator adds POD_NAME as suffix for backup destination folder.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"disable_daily": schema.BoolAttribute{
								Description:         "Defines if daily backups disabled (default false)",
								MarkdownDescription: "Defines if daily backups disabled (default false)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"disable_hourly": schema.BoolAttribute{
								Description:         "Defines if hourly backups disabled (default false)",
								MarkdownDescription: "Defines if hourly backups disabled (default false)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"disable_monthly": schema.BoolAttribute{
								Description:         "Defines if monthly backups disabled (default false)",
								MarkdownDescription: "Defines if monthly backups disabled (default false)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"disable_weekly": schema.BoolAttribute{
								Description:         "Defines if weekly backups disabled (default false)",
								MarkdownDescription: "Defines if weekly backups disabled (default false)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"extra_args": schema.MapAttribute{
								Description:         "extra args like maxBytesPerSecond default 0",
								MarkdownDescription: "extra args like maxBytesPerSecond default 0",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"extra_envs": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
											MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Variable references $(VAR_NAME) are expandedusing the previously defined environment variables in the container andany service environment variables. If a variable cannot be resolved,the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.'$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'.Escaped references will never be expanded, regardless of whether the variableexists or not.Defaults to ''.",
											MarkdownDescription: "Variable references $(VAR_NAME) are expandedusing the previously defined environment variables in the container andany service environment variables. If a variable cannot be resolved,the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.'$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'.Escaped references will never be expanded, regardless of whether the variableexists or not.Defaults to ''.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value_from": schema.SingleNestedAttribute{
											Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
											MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
													Description:         "Selects a key of a ConfigMap.",
													MarkdownDescription: "Selects a key of a ConfigMap.",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key to select.",
															MarkdownDescription: "The key to select.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
															MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"optional": schema.BoolAttribute{
															Description:         "Specify whether the ConfigMap or its key must be defined",
															MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"field_ref": schema.SingleNestedAttribute{
													Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']',spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
													MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']',spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
													Attributes: map[string]schema.Attribute{
														"api_version": schema.StringAttribute{
															Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
															MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"field_path": schema.StringAttribute{
															Description:         "Path of the field to select in the specified API version.",
															MarkdownDescription: "Path of the field to select in the specified API version.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"resource_field_ref": schema.SingleNestedAttribute{
													Description:         "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
													MarkdownDescription: "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
													Attributes: map[string]schema.Attribute{
														"container_name": schema.StringAttribute{
															Description:         "Container name: required for volumes, optional for env vars",
															MarkdownDescription: "Container name: required for volumes, optional for env vars",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"divisor": schema.StringAttribute{
															Description:         "Specifies the output format of the exposed resources, defaults to '1'",
															MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"resource": schema.StringAttribute{
															Description:         "Required: resource to select",
															MarkdownDescription: "Required: resource to select",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"secret_key_ref": schema.SingleNestedAttribute{
													Description:         "Selects a key of a secret in the pod's namespace",
													MarkdownDescription: "Selects a key of a secret in the pod's namespace",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key of the secret to select from.  Must be a valid secret key.",
															MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
															MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"optional": schema.BoolAttribute{
															Description:         "Specify whether the Secret or its key must be defined",
															MarkdownDescription: "Specify whether the Secret or its key must be defined",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"image": schema.SingleNestedAttribute{
								Description:         "Image - docker image settings for VMBackuper",
								MarkdownDescription: "Image - docker image settings for VMBackuper",
								Attributes: map[string]schema.Attribute{
									"pull_policy": schema.StringAttribute{
										Description:         "PullPolicy describes how to pull docker image",
										MarkdownDescription: "PullPolicy describes how to pull docker image",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"repository": schema.StringAttribute{
										Description:         "Repository contains name of docker image + it's repository if needed",
										MarkdownDescription: "Repository contains name of docker image + it's repository if needed",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tag": schema.StringAttribute{
										Description:         "Tag contains desired docker image version",
										MarkdownDescription: "Tag contains desired docker image version",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"log_format": schema.StringAttribute{
								Description:         "LogFormat for VMBackup to be configured with.default or json",
								MarkdownDescription: "LogFormat for VMBackup to be configured with.default or json",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("default", "json"),
								},
							},

							"log_level": schema.StringAttribute{
								Description:         "LogLevel for VMBackup to be configured with.",
								MarkdownDescription: "LogLevel for VMBackup to be configured with.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("INFO", "WARN", "ERROR", "FATAL", "PANIC"),
								},
							},

							"port": schema.StringAttribute{
								Description:         "Port for health check connections",
								MarkdownDescription: "Port for health check connections",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resources": schema.SingleNestedAttribute{
								Description:         "Resources container resource request and limits, https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/if not defined default resources from operator config will be used",
								MarkdownDescription: "Resources container resource request and limits, https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/if not defined default resources from operator config will be used",
								Attributes: map[string]schema.Attribute{
									"claims": schema.ListNestedAttribute{
										Description:         "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
										MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
													MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

							"restore": schema.SingleNestedAttribute{
								Description:         "Restore Allows to enable restore options for podRead more: https://docs.victoriametrics.com/vmbackupmanager.html#restore-commands",
								MarkdownDescription: "Restore Allows to enable restore options for podRead more: https://docs.victoriametrics.com/vmbackupmanager.html#restore-commands",
								Attributes: map[string]schema.Attribute{
									"on_start": schema.SingleNestedAttribute{
										Description:         "OnStart defines configuration for restore on pod start",
										MarkdownDescription: "OnStart defines configuration for restore on pod start",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Enabled defines if restore on start enabled",
												MarkdownDescription: "Enabled defines if restore on start enabled",
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

							"snapshot_create_url": schema.StringAttribute{
								Description:         "SnapshotCreateURL overwrites url for snapshot create",
								MarkdownDescription: "SnapshotCreateURL overwrites url for snapshot create",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"snapshot_delete_url": schema.StringAttribute{
								Description:         "SnapShotDeleteURL overwrites url for snapshot delete",
								MarkdownDescription: "SnapShotDeleteURL overwrites url for snapshot delete",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"volume_mounts": schema.ListNestedAttribute{
								Description:         "VolumeMounts allows configuration of additional VolumeMounts on the output Deployment definition.VolumeMounts specified will be appended to other VolumeMounts in the vmbackupmanager container,that are generated as a result of StorageSpec objects.",
								MarkdownDescription: "VolumeMounts allows configuration of additional VolumeMounts on the output Deployment definition.VolumeMounts specified will be appended to other VolumeMounts in the vmbackupmanager container,that are generated as a result of StorageSpec objects.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"mount_path": schema.StringAttribute{
											Description:         "Path within the container at which the volume should be mounted.  Mustnot contain ':'.",
											MarkdownDescription: "Path within the container at which the volume should be mounted.  Mustnot contain ':'.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"mount_propagation": schema.StringAttribute{
											Description:         "mountPropagation determines how mounts are propagated from the hostto container and the other way around.When not set, MountPropagationNone is used.This field is beta in 1.10.",
											MarkdownDescription: "mountPropagation determines how mounts are propagated from the hostto container and the other way around.When not set, MountPropagationNone is used.This field is beta in 1.10.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "This must match the Name of a Volume.",
											MarkdownDescription: "This must match the Name of a Volume.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"read_only": schema.BoolAttribute{
											Description:         "Mounted read-only if true, read-write otherwise (false or unspecified).Defaults to false.",
											MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified).Defaults to false.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"sub_path": schema.StringAttribute{
											Description:         "Path within the volume from which the container's volume should be mounted.Defaults to '' (volume's root).",
											MarkdownDescription: "Path within the volume from which the container's volume should be mounted.Defaults to '' (volume's root).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"sub_path_expr": schema.StringAttribute{
											Description:         "Expanded path within the volume from which the container's volume should be mounted.Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.Defaults to '' (volume's root).SubPathExpr and SubPath are mutually exclusive.",
											MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted.Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.Defaults to '' (volume's root).SubPathExpr and SubPath are mutually exclusive.",
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

					"volume_mounts": schema.ListNestedAttribute{
						Description:         "VolumeMounts allows configuration of additional VolumeMounts on the output Deployment definition.VolumeMounts specified will be appended to other VolumeMounts in the VMSingle container,that are generated as a result of StorageSpec objects.",
						MarkdownDescription: "VolumeMounts allows configuration of additional VolumeMounts on the output Deployment definition.VolumeMounts specified will be appended to other VolumeMounts in the VMSingle container,that are generated as a result of StorageSpec objects.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"mount_path": schema.StringAttribute{
									Description:         "Path within the container at which the volume should be mounted.  Mustnot contain ':'.",
									MarkdownDescription: "Path within the container at which the volume should be mounted.  Mustnot contain ':'.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"mount_propagation": schema.StringAttribute{
									Description:         "mountPropagation determines how mounts are propagated from the hostto container and the other way around.When not set, MountPropagationNone is used.This field is beta in 1.10.",
									MarkdownDescription: "mountPropagation determines how mounts are propagated from the hostto container and the other way around.When not set, MountPropagationNone is used.This field is beta in 1.10.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "This must match the Name of a Volume.",
									MarkdownDescription: "This must match the Name of a Volume.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"read_only": schema.BoolAttribute{
									Description:         "Mounted read-only if true, read-write otherwise (false or unspecified).Defaults to false.",
									MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified).Defaults to false.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"sub_path": schema.StringAttribute{
									Description:         "Path within the volume from which the container's volume should be mounted.Defaults to '' (volume's root).",
									MarkdownDescription: "Path within the volume from which the container's volume should be mounted.Defaults to '' (volume's root).",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"sub_path_expr": schema.StringAttribute{
									Description:         "Expanded path within the volume from which the container's volume should be mounted.Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.Defaults to '' (volume's root).SubPathExpr and SubPath are mutually exclusive.",
									MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted.Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.Defaults to '' (volume's root).SubPathExpr and SubPath are mutually exclusive.",
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

					"volumes": schema.ListAttribute{
						Description:         "Volumes allows configuration of additional volumes on the output deploy definition.Volumes specified will be appended to other volumes that are generated as a result ofStorageSpec objects.",
						MarkdownDescription: "Volumes allows configuration of additional volumes on the output deploy definition.Volumes specified will be appended to other volumes that are generated as a result ofStorageSpec objects.",
						ElementType:         types.MapType{ElemType: types.StringType},
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

func (r *OperatorVictoriametricsComVmsingleV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_operator_victoriametrics_com_vm_single_v1beta1_manifest")

	var model OperatorVictoriametricsComVmsingleV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("operator.victoriametrics.com/v1beta1")
	model.Kind = pointer.String("VMSingle")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
