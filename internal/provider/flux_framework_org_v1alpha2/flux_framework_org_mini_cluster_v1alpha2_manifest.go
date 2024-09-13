/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package flux_framework_org_v1alpha2

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
	_ datasource.DataSource = &FluxFrameworkOrgMiniClusterV1Alpha2Manifest{}
)

func NewFluxFrameworkOrgMiniClusterV1Alpha2Manifest() datasource.DataSource {
	return &FluxFrameworkOrgMiniClusterV1Alpha2Manifest{}
}

type FluxFrameworkOrgMiniClusterV1Alpha2Manifest struct{}

type FluxFrameworkOrgMiniClusterV1Alpha2ManifestData struct {
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
		Archive *struct {
			Path *string `tfsdk:"path" json:"path,omitempty"`
		} `tfsdk:"archive" json:"archive,omitempty"`
		Cleanup    *bool `tfsdk:"cleanup" json:"cleanup,omitempty"`
		Containers *[]struct {
			Batch    *bool   `tfsdk:"batch" json:"batch,omitempty"`
			BatchRaw *bool   `tfsdk:"batch_raw" json:"batchRaw,omitempty"`
			Command  *string `tfsdk:"command" json:"command,omitempty"`
			Commands *struct {
				BrokerPre  *string `tfsdk:"broker_pre" json:"brokerPre,omitempty"`
				Init       *string `tfsdk:"init" json:"init,omitempty"`
				Post       *string `tfsdk:"post" json:"post,omitempty"`
				Pre        *string `tfsdk:"pre" json:"pre,omitempty"`
				Prefix     *string `tfsdk:"prefix" json:"prefix,omitempty"`
				Script     *string `tfsdk:"script" json:"script,omitempty"`
				ServicePre *string `tfsdk:"service_pre" json:"servicePre,omitempty"`
				WorkerPre  *string `tfsdk:"worker_pre" json:"workerPre,omitempty"`
			} `tfsdk:"commands" json:"commands,omitempty"`
			Environment     *map[string]string `tfsdk:"environment" json:"environment,omitempty"`
			Image           *string            `tfsdk:"image" json:"image,omitempty"`
			ImagePullSecret *string            `tfsdk:"image_pull_secret" json:"imagePullSecret,omitempty"`
			Launcher        *bool              `tfsdk:"launcher" json:"launcher,omitempty"`
			LifeCycle       *struct {
				PostStartExec *string `tfsdk:"post_start_exec" json:"postStartExec,omitempty"`
				PreStopExec   *string `tfsdk:"pre_stop_exec" json:"preStopExec,omitempty"`
			} `tfsdk:"life_cycle" json:"lifeCycle,omitempty"`
			Logs             *string   `tfsdk:"logs" json:"logs,omitempty"`
			Name             *string   `tfsdk:"name" json:"name,omitempty"`
			NoWrapEntrypoint *bool     `tfsdk:"no_wrap_entrypoint" json:"noWrapEntrypoint,omitempty"`
			Ports            *[]string `tfsdk:"ports" json:"ports,omitempty"`
			PullAlways       *bool     `tfsdk:"pull_always" json:"pullAlways,omitempty"`
			Resources        *struct {
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			RunFlux *bool `tfsdk:"run_flux" json:"runFlux,omitempty"`
			Secrets *struct {
				Key  *string `tfsdk:"key" json:"key,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"secrets" json:"secrets,omitempty"`
			SecurityContext *struct {
				AddCapabilities *[]string `tfsdk:"add_capabilities" json:"addCapabilities,omitempty"`
				Privileged      *bool     `tfsdk:"privileged" json:"privileged,omitempty"`
			} `tfsdk:"security_context" json:"securityContext,omitempty"`
			Volumes *struct {
				ClaimName         *string            `tfsdk:"claim_name" json:"claimName,omitempty"`
				ConfigMapName     *string            `tfsdk:"config_map_name" json:"configMapName,omitempty"`
				EmptyDir          *bool              `tfsdk:"empty_dir" json:"emptyDir,omitempty"`
				EmptyDirMedium    *string            `tfsdk:"empty_dir_medium" json:"emptyDirMedium,omitempty"`
				EmptyDirSizeLimit *string            `tfsdk:"empty_dir_size_limit" json:"emptyDirSizeLimit,omitempty"`
				HostPath          *string            `tfsdk:"host_path" json:"hostPath,omitempty"`
				Items             *map[string]string `tfsdk:"items" json:"items,omitempty"`
				Path              *string            `tfsdk:"path" json:"path,omitempty"`
				ReadOnly          *bool              `tfsdk:"read_only" json:"readOnly,omitempty"`
				SecretName        *string            `tfsdk:"secret_name" json:"secretName,omitempty"`
			} `tfsdk:"volumes" json:"volumes,omitempty"`
			WorkingDir *string `tfsdk:"working_dir" json:"workingDir,omitempty"`
		} `tfsdk:"containers" json:"containers,omitempty"`
		DeadlineSeconds *int64 `tfsdk:"deadline_seconds" json:"deadlineSeconds,omitempty"`
		Flux            *struct {
			Arch         *string `tfsdk:"arch" json:"arch,omitempty"`
			BrokerConfig *string `tfsdk:"broker_config" json:"brokerConfig,omitempty"`
			Bursting     *struct {
				Clusters *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
					Size *int64  `tfsdk:"size" json:"size,omitempty"`
				} `tfsdk:"clusters" json:"clusters,omitempty"`
				Hostlist   *string `tfsdk:"hostlist" json:"hostlist,omitempty"`
				LeadBroker *struct {
					Address *string `tfsdk:"address" json:"address,omitempty"`
					Name    *string `tfsdk:"name" json:"name,omitempty"`
					Port    *int64  `tfsdk:"port" json:"port,omitempty"`
					Size    *int64  `tfsdk:"size" json:"size,omitempty"`
				} `tfsdk:"lead_broker" json:"leadBroker,omitempty"`
			} `tfsdk:"bursting" json:"bursting,omitempty"`
			CompleteWorkers *bool   `tfsdk:"complete_workers" json:"completeWorkers,omitempty"`
			ConnectTimeout  *string `tfsdk:"connect_timeout" json:"connectTimeout,omitempty"`
			Container       *struct {
				Disable         *bool   `tfsdk:"disable" json:"disable,omitempty"`
				Image           *string `tfsdk:"image" json:"image,omitempty"`
				ImagePullSecret *string `tfsdk:"image_pull_secret" json:"imagePullSecret,omitempty"`
				MountPath       *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
				Name            *string `tfsdk:"name" json:"name,omitempty"`
				PullAlways      *bool   `tfsdk:"pull_always" json:"pullAlways,omitempty"`
				PythonPath      *string `tfsdk:"python_path" json:"pythonPath,omitempty"`
				Resources       *struct {
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				WorkingDir *string `tfsdk:"working_dir" json:"workingDir,omitempty"`
			} `tfsdk:"container" json:"container,omitempty"`
			CurveCert      *string `tfsdk:"curve_cert" json:"curveCert,omitempty"`
			LogLevel       *int64  `tfsdk:"log_level" json:"logLevel,omitempty"`
			MinimalService *bool   `tfsdk:"minimal_service" json:"minimalService,omitempty"`
			MungeSecret    *string `tfsdk:"munge_secret" json:"mungeSecret,omitempty"`
			NoWaitSocket   *bool   `tfsdk:"no_wait_socket" json:"noWaitSocket,omitempty"`
			OptionFlags    *string `tfsdk:"option_flags" json:"optionFlags,omitempty"`
			Scheduler      *struct {
				QueuePolicy *string `tfsdk:"queue_policy" json:"queuePolicy,omitempty"`
			} `tfsdk:"scheduler" json:"scheduler,omitempty"`
			SubmitCommand *string `tfsdk:"submit_command" json:"submitCommand,omitempty"`
			Wrap          *string `tfsdk:"wrap" json:"wrap,omitempty"`
		} `tfsdk:"flux" json:"flux,omitempty"`
		Interactive *bool              `tfsdk:"interactive" json:"interactive,omitempty"`
		JobLabels   *map[string]string `tfsdk:"job_labels" json:"jobLabels,omitempty"`
		Logging     *struct {
			Debug  *bool `tfsdk:"debug" json:"debug,omitempty"`
			Quiet  *bool `tfsdk:"quiet" json:"quiet,omitempty"`
			Strict *bool `tfsdk:"strict" json:"strict,omitempty"`
			Timed  *bool `tfsdk:"timed" json:"timed,omitempty"`
			Zeromq *bool `tfsdk:"zeromq" json:"zeromq,omitempty"`
		} `tfsdk:"logging" json:"logging,omitempty"`
		MaxSize *int64 `tfsdk:"max_size" json:"maxSize,omitempty"`
		MinSize *int64 `tfsdk:"min_size" json:"minSize,omitempty"`
		Network *struct {
			DisableAffinity *bool   `tfsdk:"disable_affinity" json:"disableAffinity,omitempty"`
			HeadlessName    *string `tfsdk:"headless_name" json:"headlessName,omitempty"`
		} `tfsdk:"network" json:"network,omitempty"`
		Pod *struct {
			Annotations                  *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			AutomountServiceAccountToken *bool              `tfsdk:"automount_service_account_token" json:"automountServiceAccountToken,omitempty"`
			Labels                       *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			NodeSelector                 *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
			Resources                    *map[string]string `tfsdk:"resources" json:"resources,omitempty"`
			RestartPolicy                *string            `tfsdk:"restart_policy" json:"restartPolicy,omitempty"`
			RuntimeClassName             *string            `tfsdk:"runtime_class_name" json:"runtimeClassName,omitempty"`
			SchedulerName                *string            `tfsdk:"scheduler_name" json:"schedulerName,omitempty"`
			ServiceAccountName           *string            `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
		} `tfsdk:"pod" json:"pod,omitempty"`
		Services *[]struct {
			Batch    *bool   `tfsdk:"batch" json:"batch,omitempty"`
			BatchRaw *bool   `tfsdk:"batch_raw" json:"batchRaw,omitempty"`
			Command  *string `tfsdk:"command" json:"command,omitempty"`
			Commands *struct {
				BrokerPre  *string `tfsdk:"broker_pre" json:"brokerPre,omitempty"`
				Init       *string `tfsdk:"init" json:"init,omitempty"`
				Post       *string `tfsdk:"post" json:"post,omitempty"`
				Pre        *string `tfsdk:"pre" json:"pre,omitempty"`
				Prefix     *string `tfsdk:"prefix" json:"prefix,omitempty"`
				Script     *string `tfsdk:"script" json:"script,omitempty"`
				ServicePre *string `tfsdk:"service_pre" json:"servicePre,omitempty"`
				WorkerPre  *string `tfsdk:"worker_pre" json:"workerPre,omitempty"`
			} `tfsdk:"commands" json:"commands,omitempty"`
			Environment     *map[string]string `tfsdk:"environment" json:"environment,omitempty"`
			Image           *string            `tfsdk:"image" json:"image,omitempty"`
			ImagePullSecret *string            `tfsdk:"image_pull_secret" json:"imagePullSecret,omitempty"`
			Launcher        *bool              `tfsdk:"launcher" json:"launcher,omitempty"`
			LifeCycle       *struct {
				PostStartExec *string `tfsdk:"post_start_exec" json:"postStartExec,omitempty"`
				PreStopExec   *string `tfsdk:"pre_stop_exec" json:"preStopExec,omitempty"`
			} `tfsdk:"life_cycle" json:"lifeCycle,omitempty"`
			Logs             *string   `tfsdk:"logs" json:"logs,omitempty"`
			Name             *string   `tfsdk:"name" json:"name,omitempty"`
			NoWrapEntrypoint *bool     `tfsdk:"no_wrap_entrypoint" json:"noWrapEntrypoint,omitempty"`
			Ports            *[]string `tfsdk:"ports" json:"ports,omitempty"`
			PullAlways       *bool     `tfsdk:"pull_always" json:"pullAlways,omitempty"`
			Resources        *struct {
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			RunFlux *bool `tfsdk:"run_flux" json:"runFlux,omitempty"`
			Secrets *struct {
				Key  *string `tfsdk:"key" json:"key,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"secrets" json:"secrets,omitempty"`
			SecurityContext *struct {
				AddCapabilities *[]string `tfsdk:"add_capabilities" json:"addCapabilities,omitempty"`
				Privileged      *bool     `tfsdk:"privileged" json:"privileged,omitempty"`
			} `tfsdk:"security_context" json:"securityContext,omitempty"`
			Volumes *struct {
				ClaimName         *string            `tfsdk:"claim_name" json:"claimName,omitempty"`
				ConfigMapName     *string            `tfsdk:"config_map_name" json:"configMapName,omitempty"`
				EmptyDir          *bool              `tfsdk:"empty_dir" json:"emptyDir,omitempty"`
				EmptyDirMedium    *string            `tfsdk:"empty_dir_medium" json:"emptyDirMedium,omitempty"`
				EmptyDirSizeLimit *string            `tfsdk:"empty_dir_size_limit" json:"emptyDirSizeLimit,omitempty"`
				HostPath          *string            `tfsdk:"host_path" json:"hostPath,omitempty"`
				Items             *map[string]string `tfsdk:"items" json:"items,omitempty"`
				Path              *string            `tfsdk:"path" json:"path,omitempty"`
				ReadOnly          *bool              `tfsdk:"read_only" json:"readOnly,omitempty"`
				SecretName        *string            `tfsdk:"secret_name" json:"secretName,omitempty"`
			} `tfsdk:"volumes" json:"volumes,omitempty"`
			WorkingDir *string `tfsdk:"working_dir" json:"workingDir,omitempty"`
		} `tfsdk:"services" json:"services,omitempty"`
		ShareProcessNamespace *bool  `tfsdk:"share_process_namespace" json:"shareProcessNamespace,omitempty"`
		Size                  *int64 `tfsdk:"size" json:"size,omitempty"`
		Tasks                 *int64 `tfsdk:"tasks" json:"tasks,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *FluxFrameworkOrgMiniClusterV1Alpha2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_flux_framework_org_mini_cluster_v1alpha2_manifest"
}

func (r *FluxFrameworkOrgMiniClusterV1Alpha2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "MiniCluster is the Schema for a Flux job launcher on K8s",
		MarkdownDescription: "MiniCluster is the Schema for a Flux job launcher on K8s",
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
				Description:         "MiniCluster is an HPC cluster in Kubernetes you can control Either to submit a single job (and go away) or for a persistent single- or multi- user cluster",
				MarkdownDescription: "MiniCluster is an HPC cluster in Kubernetes you can control Either to submit a single job (and go away) or for a persistent single- or multi- user cluster",
				Attributes: map[string]schema.Attribute{
					"archive": schema.SingleNestedAttribute{
						Description:         "Archive to load or save",
						MarkdownDescription: "Archive to load or save",
						Attributes: map[string]schema.Attribute{
							"path": schema.StringAttribute{
								Description:         "Save or load from this directory path",
								MarkdownDescription: "Save or load from this directory path",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"cleanup": schema.BoolAttribute{
						Description:         "Cleanup the pods and storage when the index broker pod is complete",
						MarkdownDescription: "Cleanup the pods and storage when the index broker pod is complete",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"containers": schema.ListNestedAttribute{
						Description:         "Containers is one or more containers to be created in a pod. There should only be one container to run flux with runFlux",
						MarkdownDescription: "Containers is one or more containers to be created in a pod. There should only be one container to run flux with runFlux",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"batch": schema.BoolAttribute{
									Description:         "Indicate that the command is a batch job that will be written to a file to submit",
									MarkdownDescription: "Indicate that the command is a batch job that will be written to a file to submit",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"batch_raw": schema.BoolAttribute{
									Description:         "Don't wrap batch commands in flux submit (provide custom logic myself)",
									MarkdownDescription: "Don't wrap batch commands in flux submit (provide custom logic myself)",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"command": schema.StringAttribute{
									Description:         "Single user executable to provide to flux start",
									MarkdownDescription: "Single user executable to provide to flux start",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"commands": schema.SingleNestedAttribute{
									Description:         "More specific or detailed commands for just workers/broker",
									MarkdownDescription: "More specific or detailed commands for just workers/broker",
									Attributes: map[string]schema.Attribute{
										"broker_pre": schema.StringAttribute{
											Description:         "A single command for only the broker to run",
											MarkdownDescription: "A single command for only the broker to run",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"init": schema.StringAttribute{
											Description:         "init command is run before anything",
											MarkdownDescription: "init command is run before anything",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"post": schema.StringAttribute{
											Description:         "post command is run in the entrypoint when the broker exits / finishes",
											MarkdownDescription: "post command is run in the entrypoint when the broker exits / finishes",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"pre": schema.StringAttribute{
											Description:         "pre command is run after global PreCommand, after asFlux is set (can override)",
											MarkdownDescription: "pre command is run after global PreCommand, after asFlux is set (can override)",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"prefix": schema.StringAttribute{
											Description:         "Prefix to flux start / submit / broker Typically used for a wrapper command to mount, etc.",
											MarkdownDescription: "Prefix to flux start / submit / broker Typically used for a wrapper command to mount, etc.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"script": schema.StringAttribute{
											Description:         "Custom script for submit (e.g., multiple lines)",
											MarkdownDescription: "Custom script for submit (e.g., multiple lines)",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"service_pre": schema.StringAttribute{
											Description:         "A command only for service start.sh tor run",
											MarkdownDescription: "A command only for service start.sh tor run",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"worker_pre": schema.StringAttribute{
											Description:         "A command only for workers to run",
											MarkdownDescription: "A command only for workers to run",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"environment": schema.MapAttribute{
									Description:         "Key/value pairs for the environment",
									MarkdownDescription: "Key/value pairs for the environment",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"image": schema.StringAttribute{
									Description:         "Container image must contain flux and flux-sched install",
									MarkdownDescription: "Container image must contain flux and flux-sched install",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"image_pull_secret": schema.StringAttribute{
									Description:         "Allow the user to pull authenticated images By default no secret is selected. Setting this with the name of an already existing imagePullSecret will specify that secret in the pod spec.",
									MarkdownDescription: "Allow the user to pull authenticated images By default no secret is selected. Setting this with the name of an already existing imagePullSecret will specify that secret in the pod spec.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"launcher": schema.BoolAttribute{
									Description:         "Indicate that the command is a launcher that will ask for its own jobs (and provided directly to flux start)",
									MarkdownDescription: "Indicate that the command is a launcher that will ask for its own jobs (and provided directly to flux start)",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"life_cycle": schema.SingleNestedAttribute{
									Description:         "Lifecycle can handle post start commands, etc.",
									MarkdownDescription: "Lifecycle can handle post start commands, etc.",
									Attributes: map[string]schema.Attribute{
										"post_start_exec": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"pre_stop_exec": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"logs": schema.StringAttribute{
									Description:         "Log output directory",
									MarkdownDescription: "Log output directory",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Container name is only required for non flux runners",
									MarkdownDescription: "Container name is only required for non flux runners",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"no_wrap_entrypoint": schema.BoolAttribute{
									Description:         "Do not wrap the entrypoint to wait for flux, add to path, etc?",
									MarkdownDescription: "Do not wrap the entrypoint to wait for flux, add to path, etc?",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ports": schema.ListAttribute{
									Description:         "Ports to be exposed to other containers in the cluster We take a single list of integers and map to the same",
									MarkdownDescription: "Ports to be exposed to other containers in the cluster We take a single list of integers and map to the same",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"pull_always": schema.BoolAttribute{
									Description:         "Allow the user to dictate pulling By default we pull if not present. Setting this to true will indicate to pull always",
									MarkdownDescription: "Allow the user to dictate pulling By default we pull if not present. Setting this to true will indicate to pull always",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"resources": schema.SingleNestedAttribute{
									Description:         "Resources include limits and requests",
									MarkdownDescription: "Resources include limits and requests",
									Attributes: map[string]schema.Attribute{
										"limits": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"requests": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
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

								"run_flux": schema.BoolAttribute{
									Description:         "Application container intended to run flux (broker)",
									MarkdownDescription: "Application container intended to run flux (broker)",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"secrets": schema.SingleNestedAttribute{
									Description:         "Secrets that will be added to the environment The user is expected to create their own secrets for the operator to find",
									MarkdownDescription: "Secrets that will be added to the environment The user is expected to create their own secrets for the operator to find",
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "Key under secretKeyRef->Key",
											MarkdownDescription: "Key under secretKeyRef->Key",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name under secretKeyRef->Name",
											MarkdownDescription: "Name under secretKeyRef->Name",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"security_context": schema.SingleNestedAttribute{
									Description:         "Security Context https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",
									MarkdownDescription: "Security Context https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",
									Attributes: map[string]schema.Attribute{
										"add_capabilities": schema.ListAttribute{
											Description:         "Capabilities to add",
											MarkdownDescription: "Capabilities to add",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"privileged": schema.BoolAttribute{
											Description:         "Privileged container",
											MarkdownDescription: "Privileged container",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"volumes": schema.SingleNestedAttribute{
									Description:         "Existing volumes that can be mounted",
									MarkdownDescription: "Existing volumes that can be mounted",
									Attributes: map[string]schema.Attribute{
										"claim_name": schema.StringAttribute{
											Description:         "Claim name if the existing volume is a PVC",
											MarkdownDescription: "Claim name if the existing volume is a PVC",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"config_map_name": schema.StringAttribute{
											Description:         "Config map name if the existing volume is a config map You should also define items if you are using this",
											MarkdownDescription: "Config map name if the existing volume is a config map You should also define items if you are using this",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"empty_dir": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"empty_dir_medium": schema.StringAttribute{
											Description:         "Add an empty directory custom type",
											MarkdownDescription: "Add an empty directory custom type",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"empty_dir_size_limit": schema.StringAttribute{
											Description:         "Add an empty directory sizeLimit",
											MarkdownDescription: "Add an empty directory sizeLimit",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"host_path": schema.StringAttribute{
											Description:         "An existing hostPath to bind to path",
											MarkdownDescription: "An existing hostPath to bind to path",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"items": schema.MapAttribute{
											Description:         "Items (key and paths) for the config map",
											MarkdownDescription: "Items (key and paths) for the config map",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"path": schema.StringAttribute{
											Description:         "Path and claim name are always required if a secret isn't defined",
											MarkdownDescription: "Path and claim name are always required if a secret isn't defined",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"read_only": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret_name": schema.StringAttribute{
											Description:         "An existing secret",
											MarkdownDescription: "An existing secret",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"working_dir": schema.StringAttribute{
									Description:         "Working directory to run command from",
									MarkdownDescription: "Working directory to run command from",
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

					"deadline_seconds": schema.Int64Attribute{
						Description:         "Should the job be limited to a particular number of seconds? Approximately one year. This cannot be zero or job won't start",
						MarkdownDescription: "Should the job be limited to a particular number of seconds? Approximately one year. This cannot be zero or job won't start",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"flux": schema.SingleNestedAttribute{
						Description:         "Flux options for the broker, shared across cluster",
						MarkdownDescription: "Flux options for the broker, shared across cluster",
						Attributes: map[string]schema.Attribute{
							"arch": schema.StringAttribute{
								Description:         "Change the arch string - determines the binaries that are downloaded to run the entrypoint",
								MarkdownDescription: "Change the arch string - determines the binaries that are downloaded to run the entrypoint",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"broker_config": schema.StringAttribute{
								Description:         "Optionally provide a manually created broker config this is intended for bursting to remote clusters",
								MarkdownDescription: "Optionally provide a manually created broker config this is intended for bursting to remote clusters",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"bursting": schema.SingleNestedAttribute{
								Description:         "Bursting - one or more external clusters to burst to We assume a single, central MiniCluster with an ipaddress that all connect to.",
								MarkdownDescription: "Bursting - one or more external clusters to burst to We assume a single, central MiniCluster with an ipaddress that all connect to.",
								Attributes: map[string]schema.Attribute{
									"clusters": schema.ListNestedAttribute{
										Description:         "External clusters to burst to. Each external cluster must share the same listing to align ranks",
										MarkdownDescription: "External clusters to burst to. Each external cluster must share the same listing to align ranks",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "The hostnames for the bursted clusters If set, the user is responsible for ensuring uniqueness. The operator will set to burst-N",
													MarkdownDescription: "The hostnames for the bursted clusters If set, the user is responsible for ensuring uniqueness. The operator will set to burst-N",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"size": schema.Int64Attribute{
													Description:         "Size of bursted cluster. Defaults to same size as local minicluster if not set",
													MarkdownDescription: "Size of bursted cluster. Defaults to same size as local minicluster if not set",
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

									"hostlist": schema.StringAttribute{
										Description:         "Hostlist is a custom hostlist for the broker.toml that includes the local plus bursted cluster. This is typically used for bursting to another resource type, where we can predict the hostnames but they don't follow the same convention as the Flux Operator",
										MarkdownDescription: "Hostlist is a custom hostlist for the broker.toml that includes the local plus bursted cluster. This is typically used for bursting to another resource type, where we can predict the hostnames but they don't follow the same convention as the Flux Operator",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"lead_broker": schema.SingleNestedAttribute{
										Description:         "The lead broker ip address to join to. E.g., if we burst to cluster 2, this is the address to connect to cluster 1 For the first cluster, this should not be defined",
										MarkdownDescription: "The lead broker ip address to join to. E.g., if we burst to cluster 2, this is the address to connect to cluster 1 For the first cluster, this should not be defined",
										Attributes: map[string]schema.Attribute{
											"address": schema.StringAttribute{
												Description:         "Lead broker address (ip or hostname)",
												MarkdownDescription: "Lead broker address (ip or hostname)",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "We need the name of the lead job to assemble the hostnames",
												MarkdownDescription: "We need the name of the lead job to assemble the hostnames",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"port": schema.Int64Attribute{
												Description:         "Lead broker port - should only be used for external cluster",
												MarkdownDescription: "Lead broker port - should only be used for external cluster",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"size": schema.Int64Attribute{
												Description:         "Lead broker size",
												MarkdownDescription: "Lead broker size",
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
								Required: false,
								Optional: true,
								Computed: false,
							},

							"complete_workers": schema.BoolAttribute{
								Description:         "Complete workers when they fail This is ideal if you don't want them to restart",
								MarkdownDescription: "Complete workers when they fail This is ideal if you don't want them to restart",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"connect_timeout": schema.StringAttribute{
								Description:         "Single user executable to provide to flux start",
								MarkdownDescription: "Single user executable to provide to flux start",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"container": schema.SingleNestedAttribute{
								Description:         "Container base for flux",
								MarkdownDescription: "Container base for flux",
								Attributes: map[string]schema.Attribute{
									"disable": schema.BoolAttribute{
										Description:         "Disable the sidecar container, assuming that the main application container has flux",
										MarkdownDescription: "Disable the sidecar container, assuming that the main application container has flux",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"image": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"image_pull_secret": schema.StringAttribute{
										Description:         "Allow the user to pull authenticated images By default no secret is selected. Setting this with the name of an already existing imagePullSecret will specify that secret in the pod spec.",
										MarkdownDescription: "Allow the user to pull authenticated images By default no secret is selected. Setting this with the name of an already existing imagePullSecret will specify that secret in the pod spec.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"mount_path": schema.StringAttribute{
										Description:         "Mount path for flux to be at (will be added to path)",
										MarkdownDescription: "Mount path for flux to be at (will be added to path)",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Container name is only required for non flux runners",
										MarkdownDescription: "Container name is only required for non flux runners",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"pull_always": schema.BoolAttribute{
										Description:         "Allow the user to dictate pulling By default we pull if not present. Setting this to true will indicate to pull always",
										MarkdownDescription: "Allow the user to dictate pulling By default we pull if not present. Setting this to true will indicate to pull always",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"python_path": schema.StringAttribute{
										Description:         "Customize python path for flux",
										MarkdownDescription: "Customize python path for flux",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "Resources include limits and requests These must be defined for cpu and memory for the QoS to be Guaranteed",
										MarkdownDescription: "Resources include limits and requests These must be defined for cpu and memory for the QoS to be Guaranteed",
										Attributes: map[string]schema.Attribute{
											"limits": schema.MapAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"requests": schema.MapAttribute{
												Description:         "",
												MarkdownDescription: "",
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

									"working_dir": schema.StringAttribute{
										Description:         "Working directory to run command from",
										MarkdownDescription: "Working directory to run command from",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"curve_cert": schema.StringAttribute{
								Description:         "Optionally provide an already existing curve certificate This is not recommended in favor of providing the secret name as curveCertSecret, below",
								MarkdownDescription: "Optionally provide an already existing curve certificate This is not recommended in favor of providing the secret name as curveCertSecret, below",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"log_level": schema.Int64Attribute{
								Description:         "Log level to use for flux logging (only in non TestMode)",
								MarkdownDescription: "Log level to use for flux logging (only in non TestMode)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"minimal_service": schema.BoolAttribute{
								Description:         "Only expose the broker service (to reduce load on DNS)",
								MarkdownDescription: "Only expose the broker service (to reduce load on DNS)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"munge_secret": schema.StringAttribute{
								Description:         "Expect a secret (named according to this string) for a munge key. This is intended for bursting. Assumed to be at /etc/munge/munge.key This is binary data.",
								MarkdownDescription: "Expect a secret (named according to this string) for a munge key. This is intended for bursting. Assumed to be at /etc/munge/munge.key This is binary data.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"no_wait_socket": schema.BoolAttribute{
								Description:         "Do not wait for the socket",
								MarkdownDescription: "Do not wait for the socket",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"option_flags": schema.StringAttribute{
								Description:         "Flux option flags, usually provided with -o optional - if needed, default option flags for the server These can also be set in the user interface to override here. This is only valid for a FluxRunner 'runFlux' true",
								MarkdownDescription: "Flux option flags, usually provided with -o optional - if needed, default option flags for the server These can also be set in the user interface to override here. This is only valid for a FluxRunner 'runFlux' true",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"scheduler": schema.SingleNestedAttribute{
								Description:         "Custom attributes for the fluxion scheduler",
								MarkdownDescription: "Custom attributes for the fluxion scheduler",
								Attributes: map[string]schema.Attribute{
									"queue_policy": schema.StringAttribute{
										Description:         "Scheduler queue policy, defaults to 'fcfs' can also be 'easy'",
										MarkdownDescription: "Scheduler queue policy, defaults to 'fcfs' can also be 'easy'",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"submit_command": schema.StringAttribute{
								Description:         "Modify flux submit to be something else",
								MarkdownDescription: "Modify flux submit to be something else",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"wrap": schema.StringAttribute{
								Description:         "Commands for flux start --wrap",
								MarkdownDescription: "Commands for flux start --wrap",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"interactive": schema.BoolAttribute{
						Description:         "Run a single-user, interactive minicluster",
						MarkdownDescription: "Run a single-user, interactive minicluster",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"job_labels": schema.MapAttribute{
						Description:         "Labels for the job",
						MarkdownDescription: "Labels for the job",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"logging": schema.SingleNestedAttribute{
						Description:         "Logging modes determine the output you see in the job log",
						MarkdownDescription: "Logging modes determine the output you see in the job log",
						Attributes: map[string]schema.Attribute{
							"debug": schema.BoolAttribute{
								Description:         "Debug mode adds extra verbosity to Flux",
								MarkdownDescription: "Debug mode adds extra verbosity to Flux",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"quiet": schema.BoolAttribute{
								Description:         "Quiet mode silences all output so the job only shows the test running",
								MarkdownDescription: "Quiet mode silences all output so the job only shows the test running",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"strict": schema.BoolAttribute{
								Description:         "Strict mode ensures any failure will not continue in the job entrypoint",
								MarkdownDescription: "Strict mode ensures any failure will not continue in the job entrypoint",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"timed": schema.BoolAttribute{
								Description:         "Timed mode adds timing to Flux commands",
								MarkdownDescription: "Timed mode adds timing to Flux commands",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"zeromq": schema.BoolAttribute{
								Description:         "Enable Zeromq logging",
								MarkdownDescription: "Enable Zeromq logging",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"max_size": schema.Int64Attribute{
						Description:         "MaxSize (maximum number of pods to allow scaling to)",
						MarkdownDescription: "MaxSize (maximum number of pods to allow scaling to)",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"min_size": schema.Int64Attribute{
						Description:         "MinSize (minimum number of pods that must be up for Flux) Note that this option does not edit the number of tasks, so a job could run with fewer (and then not start)",
						MarkdownDescription: "MinSize (minimum number of pods that must be up for Flux) Note that this option does not edit the number of tasks, so a job could run with fewer (and then not start)",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"network": schema.SingleNestedAttribute{
						Description:         "A spec for exposing or defining the cluster headless service",
						MarkdownDescription: "A spec for exposing or defining the cluster headless service",
						Attributes: map[string]schema.Attribute{
							"disable_affinity": schema.BoolAttribute{
								Description:         "Disable affinity rules that guarantee one network address / node",
								MarkdownDescription: "Disable affinity rules that guarantee one network address / node",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"headless_name": schema.StringAttribute{
								Description:         "Name for cluster headless service",
								MarkdownDescription: "Name for cluster headless service",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"pod": schema.SingleNestedAttribute{
						Description:         "Pod spec details",
						MarkdownDescription: "Pod spec details",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "Annotations for each pod",
								MarkdownDescription: "Annotations for each pod",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"automount_service_account_token": schema.BoolAttribute{
								Description:         "Automatically mount the service account name",
								MarkdownDescription: "Automatically mount the service account name",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"labels": schema.MapAttribute{
								Description:         "Labels for each pod",
								MarkdownDescription: "Labels for each pod",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_selector": schema.MapAttribute{
								Description:         "NodeSelectors for a pod",
								MarkdownDescription: "NodeSelectors for a pod",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resources": schema.MapAttribute{
								Description:         "Resources include limits and requests",
								MarkdownDescription: "Resources include limits and requests",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"restart_policy": schema.StringAttribute{
								Description:         "Restart Policy",
								MarkdownDescription: "Restart Policy",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"runtime_class_name": schema.StringAttribute{
								Description:         "RuntimeClassName for the pod",
								MarkdownDescription: "RuntimeClassName for the pod",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"scheduler_name": schema.StringAttribute{
								Description:         "Scheduler name for the pod",
								MarkdownDescription: "Scheduler name for the pod",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_account_name": schema.StringAttribute{
								Description:         "Service account name for the pod",
								MarkdownDescription: "Service account name for the pod",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"services": schema.ListNestedAttribute{
						Description:         "Services are one or more service containers to bring up alongside the MiniCluster.",
						MarkdownDescription: "Services are one or more service containers to bring up alongside the MiniCluster.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"batch": schema.BoolAttribute{
									Description:         "Indicate that the command is a batch job that will be written to a file to submit",
									MarkdownDescription: "Indicate that the command is a batch job that will be written to a file to submit",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"batch_raw": schema.BoolAttribute{
									Description:         "Don't wrap batch commands in flux submit (provide custom logic myself)",
									MarkdownDescription: "Don't wrap batch commands in flux submit (provide custom logic myself)",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"command": schema.StringAttribute{
									Description:         "Single user executable to provide to flux start",
									MarkdownDescription: "Single user executable to provide to flux start",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"commands": schema.SingleNestedAttribute{
									Description:         "More specific or detailed commands for just workers/broker",
									MarkdownDescription: "More specific or detailed commands for just workers/broker",
									Attributes: map[string]schema.Attribute{
										"broker_pre": schema.StringAttribute{
											Description:         "A single command for only the broker to run",
											MarkdownDescription: "A single command for only the broker to run",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"init": schema.StringAttribute{
											Description:         "init command is run before anything",
											MarkdownDescription: "init command is run before anything",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"post": schema.StringAttribute{
											Description:         "post command is run in the entrypoint when the broker exits / finishes",
											MarkdownDescription: "post command is run in the entrypoint when the broker exits / finishes",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"pre": schema.StringAttribute{
											Description:         "pre command is run after global PreCommand, after asFlux is set (can override)",
											MarkdownDescription: "pre command is run after global PreCommand, after asFlux is set (can override)",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"prefix": schema.StringAttribute{
											Description:         "Prefix to flux start / submit / broker Typically used for a wrapper command to mount, etc.",
											MarkdownDescription: "Prefix to flux start / submit / broker Typically used for a wrapper command to mount, etc.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"script": schema.StringAttribute{
											Description:         "Custom script for submit (e.g., multiple lines)",
											MarkdownDescription: "Custom script for submit (e.g., multiple lines)",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"service_pre": schema.StringAttribute{
											Description:         "A command only for service start.sh tor run",
											MarkdownDescription: "A command only for service start.sh tor run",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"worker_pre": schema.StringAttribute{
											Description:         "A command only for workers to run",
											MarkdownDescription: "A command only for workers to run",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"environment": schema.MapAttribute{
									Description:         "Key/value pairs for the environment",
									MarkdownDescription: "Key/value pairs for the environment",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"image": schema.StringAttribute{
									Description:         "Container image must contain flux and flux-sched install",
									MarkdownDescription: "Container image must contain flux and flux-sched install",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"image_pull_secret": schema.StringAttribute{
									Description:         "Allow the user to pull authenticated images By default no secret is selected. Setting this with the name of an already existing imagePullSecret will specify that secret in the pod spec.",
									MarkdownDescription: "Allow the user to pull authenticated images By default no secret is selected. Setting this with the name of an already existing imagePullSecret will specify that secret in the pod spec.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"launcher": schema.BoolAttribute{
									Description:         "Indicate that the command is a launcher that will ask for its own jobs (and provided directly to flux start)",
									MarkdownDescription: "Indicate that the command is a launcher that will ask for its own jobs (and provided directly to flux start)",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"life_cycle": schema.SingleNestedAttribute{
									Description:         "Lifecycle can handle post start commands, etc.",
									MarkdownDescription: "Lifecycle can handle post start commands, etc.",
									Attributes: map[string]schema.Attribute{
										"post_start_exec": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"pre_stop_exec": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"logs": schema.StringAttribute{
									Description:         "Log output directory",
									MarkdownDescription: "Log output directory",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Container name is only required for non flux runners",
									MarkdownDescription: "Container name is only required for non flux runners",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"no_wrap_entrypoint": schema.BoolAttribute{
									Description:         "Do not wrap the entrypoint to wait for flux, add to path, etc?",
									MarkdownDescription: "Do not wrap the entrypoint to wait for flux, add to path, etc?",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ports": schema.ListAttribute{
									Description:         "Ports to be exposed to other containers in the cluster We take a single list of integers and map to the same",
									MarkdownDescription: "Ports to be exposed to other containers in the cluster We take a single list of integers and map to the same",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"pull_always": schema.BoolAttribute{
									Description:         "Allow the user to dictate pulling By default we pull if not present. Setting this to true will indicate to pull always",
									MarkdownDescription: "Allow the user to dictate pulling By default we pull if not present. Setting this to true will indicate to pull always",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"resources": schema.SingleNestedAttribute{
									Description:         "Resources include limits and requests",
									MarkdownDescription: "Resources include limits and requests",
									Attributes: map[string]schema.Attribute{
										"limits": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"requests": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
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

								"run_flux": schema.BoolAttribute{
									Description:         "Application container intended to run flux (broker)",
									MarkdownDescription: "Application container intended to run flux (broker)",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"secrets": schema.SingleNestedAttribute{
									Description:         "Secrets that will be added to the environment The user is expected to create their own secrets for the operator to find",
									MarkdownDescription: "Secrets that will be added to the environment The user is expected to create their own secrets for the operator to find",
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "Key under secretKeyRef->Key",
											MarkdownDescription: "Key under secretKeyRef->Key",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name under secretKeyRef->Name",
											MarkdownDescription: "Name under secretKeyRef->Name",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"security_context": schema.SingleNestedAttribute{
									Description:         "Security Context https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",
									MarkdownDescription: "Security Context https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",
									Attributes: map[string]schema.Attribute{
										"add_capabilities": schema.ListAttribute{
											Description:         "Capabilities to add",
											MarkdownDescription: "Capabilities to add",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"privileged": schema.BoolAttribute{
											Description:         "Privileged container",
											MarkdownDescription: "Privileged container",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"volumes": schema.SingleNestedAttribute{
									Description:         "Existing volumes that can be mounted",
									MarkdownDescription: "Existing volumes that can be mounted",
									Attributes: map[string]schema.Attribute{
										"claim_name": schema.StringAttribute{
											Description:         "Claim name if the existing volume is a PVC",
											MarkdownDescription: "Claim name if the existing volume is a PVC",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"config_map_name": schema.StringAttribute{
											Description:         "Config map name if the existing volume is a config map You should also define items if you are using this",
											MarkdownDescription: "Config map name if the existing volume is a config map You should also define items if you are using this",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"empty_dir": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"empty_dir_medium": schema.StringAttribute{
											Description:         "Add an empty directory custom type",
											MarkdownDescription: "Add an empty directory custom type",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"empty_dir_size_limit": schema.StringAttribute{
											Description:         "Add an empty directory sizeLimit",
											MarkdownDescription: "Add an empty directory sizeLimit",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"host_path": schema.StringAttribute{
											Description:         "An existing hostPath to bind to path",
											MarkdownDescription: "An existing hostPath to bind to path",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"items": schema.MapAttribute{
											Description:         "Items (key and paths) for the config map",
											MarkdownDescription: "Items (key and paths) for the config map",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"path": schema.StringAttribute{
											Description:         "Path and claim name are always required if a secret isn't defined",
											MarkdownDescription: "Path and claim name are always required if a secret isn't defined",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"read_only": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret_name": schema.StringAttribute{
											Description:         "An existing secret",
											MarkdownDescription: "An existing secret",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"working_dir": schema.StringAttribute{
									Description:         "Working directory to run command from",
									MarkdownDescription: "Working directory to run command from",
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

					"share_process_namespace": schema.BoolAttribute{
						Description:         "Share process namespace?",
						MarkdownDescription: "Share process namespace?",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"size": schema.Int64Attribute{
						Description:         "Size (number of job pods to run, size of minicluster in pods) This is also the minimum number required to start Flux",
						MarkdownDescription: "Size (number of job pods to run, size of minicluster in pods) This is also the minimum number required to start Flux",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tasks": schema.Int64Attribute{
						Description:         "Total number of CPUs being run across entire cluster",
						MarkdownDescription: "Total number of CPUs being run across entire cluster",
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

func (r *FluxFrameworkOrgMiniClusterV1Alpha2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_flux_framework_org_mini_cluster_v1alpha2_manifest")

	var model FluxFrameworkOrgMiniClusterV1Alpha2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("flux-framework.org/v1alpha2")
	model.Kind = pointer.String("MiniCluster")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
