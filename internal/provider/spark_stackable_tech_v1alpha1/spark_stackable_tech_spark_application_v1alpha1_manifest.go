/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package spark_stackable_tech_v1alpha1

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
	_ datasource.DataSource = &SparkStackableTechSparkApplicationV1Alpha1Manifest{}
)

func NewSparkStackableTechSparkApplicationV1Alpha1Manifest() datasource.DataSource {
	return &SparkStackableTechSparkApplicationV1Alpha1Manifest{}
}

type SparkStackableTechSparkApplicationV1Alpha1Manifest struct{}

type SparkStackableTechSparkApplicationV1Alpha1ManifestData struct {
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
		Args *[]string `tfsdk:"args" json:"args,omitempty"`
		Deps *struct {
			ExcludePackages *[]string `tfsdk:"exclude_packages" json:"excludePackages,omitempty"`
			Packages        *[]string `tfsdk:"packages" json:"packages,omitempty"`
			Repositories    *[]string `tfsdk:"repositories" json:"repositories,omitempty"`
			Requirements    *[]string `tfsdk:"requirements" json:"requirements,omitempty"`
		} `tfsdk:"deps" json:"deps,omitempty"`
		Driver *struct {
			CliOverrides *map[string]string `tfsdk:"cli_overrides" json:"cliOverrides,omitempty"`
			Config       *struct {
				Affinity *struct {
					NodeAffinity    *map[string]string `tfsdk:"node_affinity" json:"nodeAffinity,omitempty"`
					NodeSelector    *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
					PodAffinity     *map[string]string `tfsdk:"pod_affinity" json:"podAffinity,omitempty"`
					PodAntiAffinity *map[string]string `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
				} `tfsdk:"affinity" json:"affinity,omitempty"`
				Logging *struct {
					Containers *struct {
						Console *struct {
							Level *string `tfsdk:"level" json:"level,omitempty"`
						} `tfsdk:"console" json:"console,omitempty"`
						Custom *struct {
							ConfigMap *string `tfsdk:"config_map" json:"configMap,omitempty"`
						} `tfsdk:"custom" json:"custom,omitempty"`
						File *struct {
							Level *string `tfsdk:"level" json:"level,omitempty"`
						} `tfsdk:"file" json:"file,omitempty"`
						Loggers *struct {
							Level *string `tfsdk:"level" json:"level,omitempty"`
						} `tfsdk:"loggers" json:"loggers,omitempty"`
					} `tfsdk:"containers" json:"containers,omitempty"`
					EnableVectorAgent *bool `tfsdk:"enable_vector_agent" json:"enableVectorAgent,omitempty"`
				} `tfsdk:"logging" json:"logging,omitempty"`
				Resources *struct {
					Cpu *struct {
						Max *string `tfsdk:"max" json:"max,omitempty"`
						Min *string `tfsdk:"min" json:"min,omitempty"`
					} `tfsdk:"cpu" json:"cpu,omitempty"`
					Memory *struct {
						Limit         *string            `tfsdk:"limit" json:"limit,omitempty"`
						RuntimeLimits *map[string]string `tfsdk:"runtime_limits" json:"runtimeLimits,omitempty"`
					} `tfsdk:"memory" json:"memory,omitempty"`
					Storage *map[string]string `tfsdk:"storage" json:"storage,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				VolumeMounts *[]map[string]string `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
			} `tfsdk:"config" json:"config,omitempty"`
			ConfigOverrides *map[string]map[string]string `tfsdk:"config_overrides" json:"configOverrides,omitempty"`
			EnvOverrides    *map[string]string            `tfsdk:"env_overrides" json:"envOverrides,omitempty"`
			PodOverrides    *map[string]string            `tfsdk:"pod_overrides" json:"podOverrides,omitempty"`
		} `tfsdk:"driver" json:"driver,omitempty"`
		Env *[]struct {
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
		} `tfsdk:"env" json:"env,omitempty"`
		Executor *struct {
			CliOverrides *map[string]string `tfsdk:"cli_overrides" json:"cliOverrides,omitempty"`
			Config       *struct {
				Affinity *struct {
					NodeAffinity    *map[string]string `tfsdk:"node_affinity" json:"nodeAffinity,omitempty"`
					NodeSelector    *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
					PodAffinity     *map[string]string `tfsdk:"pod_affinity" json:"podAffinity,omitempty"`
					PodAntiAffinity *map[string]string `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
				} `tfsdk:"affinity" json:"affinity,omitempty"`
				Logging *struct {
					Containers *struct {
						Console *struct {
							Level *string `tfsdk:"level" json:"level,omitempty"`
						} `tfsdk:"console" json:"console,omitempty"`
						Custom *struct {
							ConfigMap *string `tfsdk:"config_map" json:"configMap,omitempty"`
						} `tfsdk:"custom" json:"custom,omitempty"`
						File *struct {
							Level *string `tfsdk:"level" json:"level,omitempty"`
						} `tfsdk:"file" json:"file,omitempty"`
						Loggers *struct {
							Level *string `tfsdk:"level" json:"level,omitempty"`
						} `tfsdk:"loggers" json:"loggers,omitempty"`
					} `tfsdk:"containers" json:"containers,omitempty"`
					EnableVectorAgent *bool `tfsdk:"enable_vector_agent" json:"enableVectorAgent,omitempty"`
				} `tfsdk:"logging" json:"logging,omitempty"`
				Resources *struct {
					Cpu *struct {
						Max *string `tfsdk:"max" json:"max,omitempty"`
						Min *string `tfsdk:"min" json:"min,omitempty"`
					} `tfsdk:"cpu" json:"cpu,omitempty"`
					Memory *struct {
						Limit         *string            `tfsdk:"limit" json:"limit,omitempty"`
						RuntimeLimits *map[string]string `tfsdk:"runtime_limits" json:"runtimeLimits,omitempty"`
					} `tfsdk:"memory" json:"memory,omitempty"`
					Storage *map[string]string `tfsdk:"storage" json:"storage,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				VolumeMounts *[]map[string]string `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
			} `tfsdk:"config" json:"config,omitempty"`
			ConfigOverrides *map[string]map[string]string `tfsdk:"config_overrides" json:"configOverrides,omitempty"`
			EnvOverrides    *map[string]string            `tfsdk:"env_overrides" json:"envOverrides,omitempty"`
			PodOverrides    *map[string]string            `tfsdk:"pod_overrides" json:"podOverrides,omitempty"`
			Replicas        *int64                        `tfsdk:"replicas" json:"replicas,omitempty"`
		} `tfsdk:"executor" json:"executor,omitempty"`
		Image *string `tfsdk:"image" json:"image,omitempty"`
		Job   *struct {
			CliOverrides *map[string]string `tfsdk:"cli_overrides" json:"cliOverrides,omitempty"`
			Config       *struct {
				Resources *struct {
					Cpu *struct {
						Max *string `tfsdk:"max" json:"max,omitempty"`
						Min *string `tfsdk:"min" json:"min,omitempty"`
					} `tfsdk:"cpu" json:"cpu,omitempty"`
					Memory *struct {
						Limit         *string            `tfsdk:"limit" json:"limit,omitempty"`
						RuntimeLimits *map[string]string `tfsdk:"runtime_limits" json:"runtimeLimits,omitempty"`
					} `tfsdk:"memory" json:"memory,omitempty"`
					Storage *map[string]string `tfsdk:"storage" json:"storage,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				VolumeMounts *[]map[string]string `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
			} `tfsdk:"config" json:"config,omitempty"`
			ConfigOverrides *map[string]map[string]string `tfsdk:"config_overrides" json:"configOverrides,omitempty"`
			EnvOverrides    *map[string]string            `tfsdk:"env_overrides" json:"envOverrides,omitempty"`
			PodOverrides    *map[string]string            `tfsdk:"pod_overrides" json:"podOverrides,omitempty"`
		} `tfsdk:"job" json:"job,omitempty"`
		LogFileDirectory *struct {
			S3 *struct {
				Bucket *struct {
					Inline *struct {
						BucketName *string `tfsdk:"bucket_name" json:"bucketName,omitempty"`
						Connection *struct {
							Inline *struct {
								AccessStyle *string `tfsdk:"access_style" json:"accessStyle,omitempty"`
								Credentials *struct {
									Scope *struct {
										ListenerVolumes *[]string `tfsdk:"listener_volumes" json:"listenerVolumes,omitempty"`
										Node            *bool     `tfsdk:"node" json:"node,omitempty"`
										Pod             *bool     `tfsdk:"pod" json:"pod,omitempty"`
										Services        *[]string `tfsdk:"services" json:"services,omitempty"`
									} `tfsdk:"scope" json:"scope,omitempty"`
									SecretClass *string `tfsdk:"secret_class" json:"secretClass,omitempty"`
								} `tfsdk:"credentials" json:"credentials,omitempty"`
								Host *string `tfsdk:"host" json:"host,omitempty"`
								Port *int64  `tfsdk:"port" json:"port,omitempty"`
								Tls  *struct {
									Verification *struct {
										None   *map[string]string `tfsdk:"none" json:"none,omitempty"`
										Server *struct {
											CaCert *struct {
												SecretClass *string            `tfsdk:"secret_class" json:"secretClass,omitempty"`
												WebPki      *map[string]string `tfsdk:"web_pki" json:"webPki,omitempty"`
											} `tfsdk:"ca_cert" json:"caCert,omitempty"`
										} `tfsdk:"server" json:"server,omitempty"`
									} `tfsdk:"verification" json:"verification,omitempty"`
								} `tfsdk:"tls" json:"tls,omitempty"`
							} `tfsdk:"inline" json:"inline,omitempty"`
							Reference *string `tfsdk:"reference" json:"reference,omitempty"`
						} `tfsdk:"connection" json:"connection,omitempty"`
					} `tfsdk:"inline" json:"inline,omitempty"`
					Reference *string `tfsdk:"reference" json:"reference,omitempty"`
				} `tfsdk:"bucket" json:"bucket,omitempty"`
				Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
			} `tfsdk:"s3" json:"s3,omitempty"`
		} `tfsdk:"log_file_directory" json:"logFileDirectory,omitempty"`
		MainApplicationFile *string `tfsdk:"main_application_file" json:"mainApplicationFile,omitempty"`
		MainClass           *string `tfsdk:"main_class" json:"mainClass,omitempty"`
		Mode                *string `tfsdk:"mode" json:"mode,omitempty"`
		S3connection        *struct {
			Inline *struct {
				AccessStyle *string `tfsdk:"access_style" json:"accessStyle,omitempty"`
				Credentials *struct {
					Scope *struct {
						ListenerVolumes *[]string `tfsdk:"listener_volumes" json:"listenerVolumes,omitempty"`
						Node            *bool     `tfsdk:"node" json:"node,omitempty"`
						Pod             *bool     `tfsdk:"pod" json:"pod,omitempty"`
						Services        *[]string `tfsdk:"services" json:"services,omitempty"`
					} `tfsdk:"scope" json:"scope,omitempty"`
					SecretClass *string `tfsdk:"secret_class" json:"secretClass,omitempty"`
				} `tfsdk:"credentials" json:"credentials,omitempty"`
				Host *string `tfsdk:"host" json:"host,omitempty"`
				Port *int64  `tfsdk:"port" json:"port,omitempty"`
				Tls  *struct {
					Verification *struct {
						None   *map[string]string `tfsdk:"none" json:"none,omitempty"`
						Server *struct {
							CaCert *struct {
								SecretClass *string            `tfsdk:"secret_class" json:"secretClass,omitempty"`
								WebPki      *map[string]string `tfsdk:"web_pki" json:"webPki,omitempty"`
							} `tfsdk:"ca_cert" json:"caCert,omitempty"`
						} `tfsdk:"server" json:"server,omitempty"`
					} `tfsdk:"verification" json:"verification,omitempty"`
				} `tfsdk:"tls" json:"tls,omitempty"`
			} `tfsdk:"inline" json:"inline,omitempty"`
			Reference *string `tfsdk:"reference" json:"reference,omitempty"`
		} `tfsdk:"s3connection" json:"s3connection,omitempty"`
		SparkConf  *map[string]string `tfsdk:"spark_conf" json:"sparkConf,omitempty"`
		SparkImage *struct {
			Custom         *string `tfsdk:"custom" json:"custom,omitempty"`
			ProductVersion *string `tfsdk:"product_version" json:"productVersion,omitempty"`
			PullPolicy     *string `tfsdk:"pull_policy" json:"pullPolicy,omitempty"`
			PullSecrets    *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"pull_secrets" json:"pullSecrets,omitempty"`
			Repo             *string `tfsdk:"repo" json:"repo,omitempty"`
			StackableVersion *string `tfsdk:"stackable_version" json:"stackableVersion,omitempty"`
		} `tfsdk:"spark_image" json:"sparkImage,omitempty"`
		VectorAggregatorConfigMapName *string              `tfsdk:"vector_aggregator_config_map_name" json:"vectorAggregatorConfigMapName,omitempty"`
		Volumes                       *[]map[string]string `tfsdk:"volumes" json:"volumes,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SparkStackableTechSparkApplicationV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_spark_stackable_tech_spark_application_v1alpha1_manifest"
}

func (r *SparkStackableTechSparkApplicationV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Auto-generated derived type for SparkApplicationSpec via 'CustomResource'",
		MarkdownDescription: "Auto-generated derived type for SparkApplicationSpec via 'CustomResource'",
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
				Description:         "A Spark cluster stacklet. This resource is managed by the Stackable operator for Apache Spark. Find more information on how to use it and the resources that the operator generates in the [operator documentation](https://docs.stackable.tech/home/nightly/spark-k8s/). The SparkApplication CRD looks a little different than the CRDs of the other products on the Stackable Data Platform.",
				MarkdownDescription: "A Spark cluster stacklet. This resource is managed by the Stackable operator for Apache Spark. Find more information on how to use it and the resources that the operator generates in the [operator documentation](https://docs.stackable.tech/home/nightly/spark-k8s/). The SparkApplication CRD looks a little different than the CRDs of the other products on the Stackable Data Platform.",
				Attributes: map[string]schema.Attribute{
					"args": schema.ListAttribute{
						Description:         "Arguments passed directly to the job artifact.",
						MarkdownDescription: "Arguments passed directly to the job artifact.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"deps": schema.SingleNestedAttribute{
						Description:         "Job dependencies: a list of python packages that will be installed via pip, a list of packages or repositories that is passed directly to spark-submit, or a list of excluded packages (also passed directly to spark-submit).",
						MarkdownDescription: "Job dependencies: a list of python packages that will be installed via pip, a list of packages or repositories that is passed directly to spark-submit, or a list of excluded packages (also passed directly to spark-submit).",
						Attributes: map[string]schema.Attribute{
							"exclude_packages": schema.ListAttribute{
								Description:         "A list of excluded packages that is passed directly to 'spark-submit'.",
								MarkdownDescription: "A list of excluded packages that is passed directly to 'spark-submit'.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"packages": schema.ListAttribute{
								Description:         "A list of packages that is passed directly to 'spark-submit'.",
								MarkdownDescription: "A list of packages that is passed directly to 'spark-submit'.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"repositories": schema.ListAttribute{
								Description:         "A list of repositories that is passed directly to 'spark-submit'.",
								MarkdownDescription: "A list of repositories that is passed directly to 'spark-submit'.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"requirements": schema.ListAttribute{
								Description:         "Under the 'requirements' you can specify Python dependencies that will be installed with 'pip'. Example: 'tabulate==0.8.9'",
								MarkdownDescription: "Under the 'requirements' you can specify Python dependencies that will be installed with 'pip'. Example: 'tabulate==0.8.9'",
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

					"driver": schema.SingleNestedAttribute{
						Description:         "The driver role specifies the configuration that, together with the driver pod template, is used by Spark to create driver pods.",
						MarkdownDescription: "The driver role specifies the configuration that, together with the driver pod template, is used by Spark to create driver pods.",
						Attributes: map[string]schema.Attribute{
							"cli_overrides": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"config": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"affinity": schema.SingleNestedAttribute{
										Description:         "These configuration settings control [Pod placement](https://docs.stackable.tech/home/nightly/concepts/operations/pod_placement).",
										MarkdownDescription: "These configuration settings control [Pod placement](https://docs.stackable.tech/home/nightly/concepts/operations/pod_placement).",
										Attributes: map[string]schema.Attribute{
											"node_affinity": schema.MapAttribute{
												Description:         "Same as the 'spec.affinity.nodeAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
												MarkdownDescription: "Same as the 'spec.affinity.nodeAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"node_selector": schema.MapAttribute{
												Description:         "Simple key-value pairs forming a nodeSelector, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
												MarkdownDescription: "Simple key-value pairs forming a nodeSelector, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"pod_affinity": schema.MapAttribute{
												Description:         "Same as the 'spec.affinity.podAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
												MarkdownDescription: "Same as the 'spec.affinity.podAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"pod_anti_affinity": schema.MapAttribute{
												Description:         "Same as the 'spec.affinity.podAntiAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
												MarkdownDescription: "Same as the 'spec.affinity.podAntiAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
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

									"logging": schema.SingleNestedAttribute{
										Description:         "Logging configuration, learn more in the [logging concept documentation](https://docs.stackable.tech/home/nightly/concepts/logging).",
										MarkdownDescription: "Logging configuration, learn more in the [logging concept documentation](https://docs.stackable.tech/home/nightly/concepts/logging).",
										Attributes: map[string]schema.Attribute{
											"containers": schema.SingleNestedAttribute{
												Description:         "Log configuration per container.",
												MarkdownDescription: "Log configuration per container.",
												Attributes: map[string]schema.Attribute{
													"console": schema.SingleNestedAttribute{
														Description:         "Configuration for the console appender",
														MarkdownDescription: "Configuration for the console appender",
														Attributes: map[string]schema.Attribute{
															"level": schema.StringAttribute{
																Description:         "The log level threshold. Log events with a lower log level are discarded.",
																MarkdownDescription: "The log level threshold. Log events with a lower log level are discarded.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", "NONE"),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"custom": schema.SingleNestedAttribute{
														Description:         "Custom log configuration provided in a ConfigMap",
														MarkdownDescription: "Custom log configuration provided in a ConfigMap",
														Attributes: map[string]schema.Attribute{
															"config_map": schema.StringAttribute{
																Description:         "ConfigMap containing the log configuration files",
																MarkdownDescription: "ConfigMap containing the log configuration files",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"file": schema.SingleNestedAttribute{
														Description:         "Configuration for the file appender",
														MarkdownDescription: "Configuration for the file appender",
														Attributes: map[string]schema.Attribute{
															"level": schema.StringAttribute{
																Description:         "The log level threshold. Log events with a lower log level are discarded.",
																MarkdownDescription: "The log level threshold. Log events with a lower log level are discarded.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", "NONE"),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"loggers": schema.SingleNestedAttribute{
														Description:         "Configuration per logger",
														MarkdownDescription: "Configuration per logger",
														Attributes: map[string]schema.Attribute{
															"level": schema.StringAttribute{
																Description:         "The log level threshold. Log events with a lower log level are discarded.",
																MarkdownDescription: "The log level threshold. Log events with a lower log level are discarded.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", "NONE"),
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

											"enable_vector_agent": schema.BoolAttribute{
												Description:         "Wether or not to deploy a container with the Vector log agent.",
												MarkdownDescription: "Wether or not to deploy a container with the Vector log agent.",
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
										Description:         "Resource usage is configured here, this includes CPU usage, memory usage and disk storage usage, if this role needs any.",
										MarkdownDescription: "Resource usage is configured here, this includes CPU usage, memory usage and disk storage usage, if this role needs any.",
										Attributes: map[string]schema.Attribute{
											"cpu": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"max": schema.StringAttribute{
														Description:         "The maximum amount of CPU cores that can be requested by Pods. Equivalent to the 'limit' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.",
														MarkdownDescription: "The maximum amount of CPU cores that can be requested by Pods. Equivalent to the 'limit' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"min": schema.StringAttribute{
														Description:         "The minimal amount of CPU cores that Pods need to run. Equivalent to the 'request' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.",
														MarkdownDescription: "The minimal amount of CPU cores that Pods need to run. Equivalent to the 'request' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"memory": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"limit": schema.StringAttribute{
														Description:         "The maximum amount of memory that should be available to the Pod. Specified as a byte [Quantity](https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/), which means these suffixes are supported: E, P, T, G, M, k. You can also use the power-of-two equivalents: Ei, Pi, Ti, Gi, Mi, Ki. For example, the following represent roughly the same value: '128974848, 129e6, 129M, 128974848000m, 123Mi'",
														MarkdownDescription: "The maximum amount of memory that should be available to the Pod. Specified as a byte [Quantity](https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/), which means these suffixes are supported: E, P, T, G, M, k. You can also use the power-of-two equivalents: Ei, Pi, Ti, Gi, Mi, Ki. For example, the following represent roughly the same value: '128974848, 129e6, 129M, 128974848000m, 123Mi'",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"runtime_limits": schema.MapAttribute{
														Description:         "Additional options that can be specified.",
														MarkdownDescription: "Additional options that can be specified.",
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

											"storage": schema.MapAttribute{
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

									"volume_mounts": schema.ListAttribute{
										Description:         "Volume mounts for the spark-submit, driver and executor pods.",
										MarkdownDescription: "Volume mounts for the spark-submit, driver and executor pods.",
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

							"config_overrides": schema.MapAttribute{
								Description:         "The 'configOverrides' can be used to configure properties in product config files that are not exposed in the CRD. Read the [config overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#config-overrides) and consult the operator specific usage guide documentation for details on the available config files and settings for the specific product.",
								MarkdownDescription: "The 'configOverrides' can be used to configure properties in product config files that are not exposed in the CRD. Read the [config overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#config-overrides) and consult the operator specific usage guide documentation for details on the available config files and settings for the specific product.",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"env_overrides": schema.MapAttribute{
								Description:         "'envOverrides' configure environment variables to be set in the Pods. It is a map from strings to strings - environment variables and the value to set. Read the [environment variable overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#env-overrides) for more information and consult the operator specific usage guide to find out about the product specific environment variables that are available.",
								MarkdownDescription: "'envOverrides' configure environment variables to be set in the Pods. It is a map from strings to strings - environment variables and the value to set. Read the [environment variable overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#env-overrides) for more information and consult the operator specific usage guide to find out about the product specific environment variables that are available.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pod_overrides": schema.MapAttribute{
								Description:         "In the 'podOverrides' property you can define a [PodTemplateSpec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#podtemplatespec-v1-core) to override any property that can be set on a Kubernetes Pod. Read the [Pod overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#pod-overrides) for more information.",
								MarkdownDescription: "In the 'podOverrides' property you can define a [PodTemplateSpec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#podtemplatespec-v1-core) to override any property that can be set on a Kubernetes Pod. Read the [Pod overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#pod-overrides) for more information.",
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

					"env": schema.ListNestedAttribute{
						Description:         "A list of environment variables that will be set in the job pod and the driver and executor pod templates.",
						MarkdownDescription: "A list of environment variables that will be set in the job pod and the driver and executor pod templates.",
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
									Description:         "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
									MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
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
													Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
													MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
													Required:            true,
													Optional:            false,
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
											Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
											MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
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
											Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
											MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
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
													Description:         "The key of the secret to select from. Must be a valid secret key.",
													MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
													MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
													Required:            true,
													Optional:            false,
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

					"executor": schema.SingleNestedAttribute{
						Description:         "The executor role specifies the configuration that, together with the driver pod template, is used by Spark to create the executor pods. This is RoleGroup instead of plain CommonConfiguration because it needs to allows for the number of replicas. to be specified.",
						MarkdownDescription: "The executor role specifies the configuration that, together with the driver pod template, is used by Spark to create the executor pods. This is RoleGroup instead of plain CommonConfiguration because it needs to allows for the number of replicas. to be specified.",
						Attributes: map[string]schema.Attribute{
							"cli_overrides": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"config": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"affinity": schema.SingleNestedAttribute{
										Description:         "These configuration settings control [Pod placement](https://docs.stackable.tech/home/nightly/concepts/operations/pod_placement).",
										MarkdownDescription: "These configuration settings control [Pod placement](https://docs.stackable.tech/home/nightly/concepts/operations/pod_placement).",
										Attributes: map[string]schema.Attribute{
											"node_affinity": schema.MapAttribute{
												Description:         "Same as the 'spec.affinity.nodeAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
												MarkdownDescription: "Same as the 'spec.affinity.nodeAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"node_selector": schema.MapAttribute{
												Description:         "Simple key-value pairs forming a nodeSelector, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
												MarkdownDescription: "Simple key-value pairs forming a nodeSelector, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"pod_affinity": schema.MapAttribute{
												Description:         "Same as the 'spec.affinity.podAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
												MarkdownDescription: "Same as the 'spec.affinity.podAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"pod_anti_affinity": schema.MapAttribute{
												Description:         "Same as the 'spec.affinity.podAntiAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
												MarkdownDescription: "Same as the 'spec.affinity.podAntiAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
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

									"logging": schema.SingleNestedAttribute{
										Description:         "Logging configuration, learn more in the [logging concept documentation](https://docs.stackable.tech/home/nightly/concepts/logging).",
										MarkdownDescription: "Logging configuration, learn more in the [logging concept documentation](https://docs.stackable.tech/home/nightly/concepts/logging).",
										Attributes: map[string]schema.Attribute{
											"containers": schema.SingleNestedAttribute{
												Description:         "Log configuration per container.",
												MarkdownDescription: "Log configuration per container.",
												Attributes: map[string]schema.Attribute{
													"console": schema.SingleNestedAttribute{
														Description:         "Configuration for the console appender",
														MarkdownDescription: "Configuration for the console appender",
														Attributes: map[string]schema.Attribute{
															"level": schema.StringAttribute{
																Description:         "The log level threshold. Log events with a lower log level are discarded.",
																MarkdownDescription: "The log level threshold. Log events with a lower log level are discarded.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", "NONE"),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"custom": schema.SingleNestedAttribute{
														Description:         "Custom log configuration provided in a ConfigMap",
														MarkdownDescription: "Custom log configuration provided in a ConfigMap",
														Attributes: map[string]schema.Attribute{
															"config_map": schema.StringAttribute{
																Description:         "ConfigMap containing the log configuration files",
																MarkdownDescription: "ConfigMap containing the log configuration files",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"file": schema.SingleNestedAttribute{
														Description:         "Configuration for the file appender",
														MarkdownDescription: "Configuration for the file appender",
														Attributes: map[string]schema.Attribute{
															"level": schema.StringAttribute{
																Description:         "The log level threshold. Log events with a lower log level are discarded.",
																MarkdownDescription: "The log level threshold. Log events with a lower log level are discarded.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", "NONE"),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"loggers": schema.SingleNestedAttribute{
														Description:         "Configuration per logger",
														MarkdownDescription: "Configuration per logger",
														Attributes: map[string]schema.Attribute{
															"level": schema.StringAttribute{
																Description:         "The log level threshold. Log events with a lower log level are discarded.",
																MarkdownDescription: "The log level threshold. Log events with a lower log level are discarded.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", "NONE"),
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

											"enable_vector_agent": schema.BoolAttribute{
												Description:         "Wether or not to deploy a container with the Vector log agent.",
												MarkdownDescription: "Wether or not to deploy a container with the Vector log agent.",
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
										Description:         "Resource usage is configured here, this includes CPU usage, memory usage and disk storage usage, if this role needs any.",
										MarkdownDescription: "Resource usage is configured here, this includes CPU usage, memory usage and disk storage usage, if this role needs any.",
										Attributes: map[string]schema.Attribute{
											"cpu": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"max": schema.StringAttribute{
														Description:         "The maximum amount of CPU cores that can be requested by Pods. Equivalent to the 'limit' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.",
														MarkdownDescription: "The maximum amount of CPU cores that can be requested by Pods. Equivalent to the 'limit' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"min": schema.StringAttribute{
														Description:         "The minimal amount of CPU cores that Pods need to run. Equivalent to the 'request' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.",
														MarkdownDescription: "The minimal amount of CPU cores that Pods need to run. Equivalent to the 'request' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"memory": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"limit": schema.StringAttribute{
														Description:         "The maximum amount of memory that should be available to the Pod. Specified as a byte [Quantity](https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/), which means these suffixes are supported: E, P, T, G, M, k. You can also use the power-of-two equivalents: Ei, Pi, Ti, Gi, Mi, Ki. For example, the following represent roughly the same value: '128974848, 129e6, 129M, 128974848000m, 123Mi'",
														MarkdownDescription: "The maximum amount of memory that should be available to the Pod. Specified as a byte [Quantity](https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/), which means these suffixes are supported: E, P, T, G, M, k. You can also use the power-of-two equivalents: Ei, Pi, Ti, Gi, Mi, Ki. For example, the following represent roughly the same value: '128974848, 129e6, 129M, 128974848000m, 123Mi'",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"runtime_limits": schema.MapAttribute{
														Description:         "Additional options that can be specified.",
														MarkdownDescription: "Additional options that can be specified.",
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

											"storage": schema.MapAttribute{
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

									"volume_mounts": schema.ListAttribute{
										Description:         "Volume mounts for the spark-submit, driver and executor pods.",
										MarkdownDescription: "Volume mounts for the spark-submit, driver and executor pods.",
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

							"config_overrides": schema.MapAttribute{
								Description:         "The 'configOverrides' can be used to configure properties in product config files that are not exposed in the CRD. Read the [config overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#config-overrides) and consult the operator specific usage guide documentation for details on the available config files and settings for the specific product.",
								MarkdownDescription: "The 'configOverrides' can be used to configure properties in product config files that are not exposed in the CRD. Read the [config overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#config-overrides) and consult the operator specific usage guide documentation for details on the available config files and settings for the specific product.",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"env_overrides": schema.MapAttribute{
								Description:         "'envOverrides' configure environment variables to be set in the Pods. It is a map from strings to strings - environment variables and the value to set. Read the [environment variable overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#env-overrides) for more information and consult the operator specific usage guide to find out about the product specific environment variables that are available.",
								MarkdownDescription: "'envOverrides' configure environment variables to be set in the Pods. It is a map from strings to strings - environment variables and the value to set. Read the [environment variable overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#env-overrides) for more information and consult the operator specific usage guide to find out about the product specific environment variables that are available.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pod_overrides": schema.MapAttribute{
								Description:         "In the 'podOverrides' property you can define a [PodTemplateSpec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#podtemplatespec-v1-core) to override any property that can be set on a Kubernetes Pod. Read the [Pod overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#pod-overrides) for more information.",
								MarkdownDescription: "In the 'podOverrides' property you can define a [PodTemplateSpec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#podtemplatespec-v1-core) to override any property that can be set on a Kubernetes Pod. Read the [Pod overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#pod-overrides) for more information.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"replicas": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"image": schema.StringAttribute{
						Description:         "User-supplied image containing spark-job dependencies that will be copied to the specified volume mount. See the [examples](https://docs.stackable.tech/home/nightly/spark-k8s/usage-guide/examples).",
						MarkdownDescription: "User-supplied image containing spark-job dependencies that will be copied to the specified volume mount. See the [examples](https://docs.stackable.tech/home/nightly/spark-k8s/usage-guide/examples).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"job": schema.SingleNestedAttribute{
						Description:         "The job builds a spark-submit command, complete with arguments and referenced dependencies such as templates, and passes it on to Spark. The reason this property uses its own type (SubmitConfigFragment) is because logging is not supported for spark-submit processes.",
						MarkdownDescription: "The job builds a spark-submit command, complete with arguments and referenced dependencies such as templates, and passes it on to Spark. The reason this property uses its own type (SubmitConfigFragment) is because logging is not supported for spark-submit processes.",
						Attributes: map[string]schema.Attribute{
							"cli_overrides": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"config": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"resources": schema.SingleNestedAttribute{
										Description:         "Resource usage is configured here, this includes CPU usage, memory usage and disk storage usage, if this role needs any.",
										MarkdownDescription: "Resource usage is configured here, this includes CPU usage, memory usage and disk storage usage, if this role needs any.",
										Attributes: map[string]schema.Attribute{
											"cpu": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"max": schema.StringAttribute{
														Description:         "The maximum amount of CPU cores that can be requested by Pods. Equivalent to the 'limit' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.",
														MarkdownDescription: "The maximum amount of CPU cores that can be requested by Pods. Equivalent to the 'limit' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"min": schema.StringAttribute{
														Description:         "The minimal amount of CPU cores that Pods need to run. Equivalent to the 'request' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.",
														MarkdownDescription: "The minimal amount of CPU cores that Pods need to run. Equivalent to the 'request' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"memory": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"limit": schema.StringAttribute{
														Description:         "The maximum amount of memory that should be available to the Pod. Specified as a byte [Quantity](https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/), which means these suffixes are supported: E, P, T, G, M, k. You can also use the power-of-two equivalents: Ei, Pi, Ti, Gi, Mi, Ki. For example, the following represent roughly the same value: '128974848, 129e6, 129M, 128974848000m, 123Mi'",
														MarkdownDescription: "The maximum amount of memory that should be available to the Pod. Specified as a byte [Quantity](https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/), which means these suffixes are supported: E, P, T, G, M, k. You can also use the power-of-two equivalents: Ei, Pi, Ti, Gi, Mi, Ki. For example, the following represent roughly the same value: '128974848, 129e6, 129M, 128974848000m, 123Mi'",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"runtime_limits": schema.MapAttribute{
														Description:         "Additional options that can be specified.",
														MarkdownDescription: "Additional options that can be specified.",
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

											"storage": schema.MapAttribute{
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

									"volume_mounts": schema.ListAttribute{
										Description:         "Volume mounts for the spark-submit, driver and executor pods.",
										MarkdownDescription: "Volume mounts for the spark-submit, driver and executor pods.",
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

							"config_overrides": schema.MapAttribute{
								Description:         "The 'configOverrides' can be used to configure properties in product config files that are not exposed in the CRD. Read the [config overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#config-overrides) and consult the operator specific usage guide documentation for details on the available config files and settings for the specific product.",
								MarkdownDescription: "The 'configOverrides' can be used to configure properties in product config files that are not exposed in the CRD. Read the [config overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#config-overrides) and consult the operator specific usage guide documentation for details on the available config files and settings for the specific product.",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"env_overrides": schema.MapAttribute{
								Description:         "'envOverrides' configure environment variables to be set in the Pods. It is a map from strings to strings - environment variables and the value to set. Read the [environment variable overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#env-overrides) for more information and consult the operator specific usage guide to find out about the product specific environment variables that are available.",
								MarkdownDescription: "'envOverrides' configure environment variables to be set in the Pods. It is a map from strings to strings - environment variables and the value to set. Read the [environment variable overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#env-overrides) for more information and consult the operator specific usage guide to find out about the product specific environment variables that are available.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pod_overrides": schema.MapAttribute{
								Description:         "In the 'podOverrides' property you can define a [PodTemplateSpec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#podtemplatespec-v1-core) to override any property that can be set on a Kubernetes Pod. Read the [Pod overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#pod-overrides) for more information.",
								MarkdownDescription: "In the 'podOverrides' property you can define a [PodTemplateSpec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#podtemplatespec-v1-core) to override any property that can be set on a Kubernetes Pod. Read the [Pod overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#pod-overrides) for more information.",
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

					"log_file_directory": schema.SingleNestedAttribute{
						Description:         "The log file directory definition used by the Spark history server. Currently only S3 buckets are supported.",
						MarkdownDescription: "The log file directory definition used by the Spark history server. Currently only S3 buckets are supported.",
						Attributes: map[string]schema.Attribute{
							"s3": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"bucket": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"inline": schema.SingleNestedAttribute{
												Description:         "S3 bucket specification containing the bucket name and an inlined or referenced connection specification. Learn more on the [S3 concept documentation](https://docs.stackable.tech/home/nightly/concepts/s3).",
												MarkdownDescription: "S3 bucket specification containing the bucket name and an inlined or referenced connection specification. Learn more on the [S3 concept documentation](https://docs.stackable.tech/home/nightly/concepts/s3).",
												Attributes: map[string]schema.Attribute{
													"bucket_name": schema.StringAttribute{
														Description:         "The name of the S3 bucket.",
														MarkdownDescription: "The name of the S3 bucket.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"connection": schema.SingleNestedAttribute{
														Description:         "The definition of an S3 connection, either inline or as a reference.",
														MarkdownDescription: "The definition of an S3 connection, either inline or as a reference.",
														Attributes: map[string]schema.Attribute{
															"inline": schema.SingleNestedAttribute{
																Description:         "S3 connection definition as a resource. Learn more on the [S3 concept documentation](https://docs.stackable.tech/home/nightly/concepts/s3).",
																MarkdownDescription: "S3 connection definition as a resource. Learn more on the [S3 concept documentation](https://docs.stackable.tech/home/nightly/concepts/s3).",
																Attributes: map[string]schema.Attribute{
																	"access_style": schema.StringAttribute{
																		Description:         "Which access style to use. Defaults to virtual hosted-style as most of the data products out there. Have a look at the [AWS documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/VirtualHosting.html).",
																		MarkdownDescription: "Which access style to use. Defaults to virtual hosted-style as most of the data products out there. Have a look at the [AWS documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/VirtualHosting.html).",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("Path", "VirtualHosted"),
																		},
																	},

																	"credentials": schema.SingleNestedAttribute{
																		Description:         "If the S3 uses authentication you have to specify you S3 credentials. In the most cases a [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) providing 'accessKey' and 'secretKey' is sufficient.",
																		MarkdownDescription: "If the S3 uses authentication you have to specify you S3 credentials. In the most cases a [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) providing 'accessKey' and 'secretKey' is sufficient.",
																		Attributes: map[string]schema.Attribute{
																			"scope": schema.SingleNestedAttribute{
																				Description:         "[Scope](https://docs.stackable.tech/home/nightly/secret-operator/scope) of the [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass).",
																				MarkdownDescription: "[Scope](https://docs.stackable.tech/home/nightly/secret-operator/scope) of the [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass).",
																				Attributes: map[string]schema.Attribute{
																					"listener_volumes": schema.ListAttribute{
																						Description:         "The listener volume scope allows Node and Service scopes to be inferred from the applicable listeners. This must correspond to Volume names in the Pod that mount Listeners.",
																						MarkdownDescription: "The listener volume scope allows Node and Service scopes to be inferred from the applicable listeners. This must correspond to Volume names in the Pod that mount Listeners.",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"node": schema.BoolAttribute{
																						Description:         "The node scope is resolved to the name of the Kubernetes Node object that the Pod is running on. This will typically be the DNS name of the node.",
																						MarkdownDescription: "The node scope is resolved to the name of the Kubernetes Node object that the Pod is running on. This will typically be the DNS name of the node.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"pod": schema.BoolAttribute{
																						Description:         "The pod scope is resolved to the name of the Kubernetes Pod. This allows the secret to differentiate between StatefulSet replicas.",
																						MarkdownDescription: "The pod scope is resolved to the name of the Kubernetes Pod. This allows the secret to differentiate between StatefulSet replicas.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"services": schema.ListAttribute{
																						Description:         "The service scope allows Pod objects to specify custom scopes. This should typically correspond to Service objects that the Pod participates in.",
																						MarkdownDescription: "The service scope allows Pod objects to specify custom scopes. This should typically correspond to Service objects that the Pod participates in.",
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

																			"secret_class": schema.StringAttribute{
																				Description:         "[SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) containing the LDAP bind credentials.",
																				MarkdownDescription: "[SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) containing the LDAP bind credentials.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"host": schema.StringAttribute{
																		Description:         "Host of the S3 server without any protocol or port. For example: 'west1.my-cloud.com'.",
																		MarkdownDescription: "Host of the S3 server without any protocol or port. For example: 'west1.my-cloud.com'.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"port": schema.Int64Attribute{
																		Description:         "Port the S3 server listens on. If not specified the product will determine the port to use.",
																		MarkdownDescription: "Port the S3 server listens on. If not specified the product will determine the port to use.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.Int64{
																			int64validator.AtLeast(0),
																		},
																	},

																	"tls": schema.SingleNestedAttribute{
																		Description:         "Use a TLS connection. If not specified no TLS will be used.",
																		MarkdownDescription: "Use a TLS connection. If not specified no TLS will be used.",
																		Attributes: map[string]schema.Attribute{
																			"verification": schema.SingleNestedAttribute{
																				Description:         "The verification method used to verify the certificates of the server and/or the client.",
																				MarkdownDescription: "The verification method used to verify the certificates of the server and/or the client.",
																				Attributes: map[string]schema.Attribute{
																					"none": schema.MapAttribute{
																						Description:         "Use TLS but don't verify certificates.",
																						MarkdownDescription: "Use TLS but don't verify certificates.",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"server": schema.SingleNestedAttribute{
																						Description:         "Use TLS and a CA certificate to verify the server.",
																						MarkdownDescription: "Use TLS and a CA certificate to verify the server.",
																						Attributes: map[string]schema.Attribute{
																							"ca_cert": schema.SingleNestedAttribute{
																								Description:         "CA cert to verify the server.",
																								MarkdownDescription: "CA cert to verify the server.",
																								Attributes: map[string]schema.Attribute{
																									"secret_class": schema.StringAttribute{
																										Description:         "Name of the [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) which will provide the CA certificate. Note that a SecretClass does not need to have a key but can also work with just a CA certificate, so if you got provided with a CA cert but don't have access to the key you can still use this method.",
																										MarkdownDescription: "Name of the [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) which will provide the CA certificate. Note that a SecretClass does not need to have a key but can also work with just a CA certificate, so if you got provided with a CA cert but don't have access to the key you can still use this method.",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},

																									"web_pki": schema.MapAttribute{
																										Description:         "Use TLS and the CA certificates trusted by the common web browsers to verify the server. This can be useful when you e.g. use public AWS S3 or other public available services.",
																										MarkdownDescription: "Use TLS and the CA certificates trusted by the common web browsers to verify the server. This can be useful when you e.g. use public AWS S3 or other public available services.",
																										ElementType:         types.StringType,
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
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
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"reference": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
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

											"reference": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"prefix": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
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

					"main_application_file": schema.StringAttribute{
						Description:         "The actual application file that will be called by 'spark-submit'.",
						MarkdownDescription: "The actual application file that will be called by 'spark-submit'.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"main_class": schema.StringAttribute{
						Description:         "The main class - i.e. entry point - for JVM artifacts.",
						MarkdownDescription: "The main class - i.e. entry point - for JVM artifacts.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"mode": schema.StringAttribute{
						Description:         "Mode: cluster or client. Currently only cluster is supported.",
						MarkdownDescription: "Mode: cluster or client. Currently only cluster is supported.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("cluster", "client"),
						},
					},

					"s3connection": schema.SingleNestedAttribute{
						Description:         "Configure an S3 connection that the SparkApplication has access to. Read more in the [Spark S3 usage guide](https://docs.stackable.tech/home/nightly/spark-k8s/usage-guide/s3).",
						MarkdownDescription: "Configure an S3 connection that the SparkApplication has access to. Read more in the [Spark S3 usage guide](https://docs.stackable.tech/home/nightly/spark-k8s/usage-guide/s3).",
						Attributes: map[string]schema.Attribute{
							"inline": schema.SingleNestedAttribute{
								Description:         "S3 connection definition as a resource. Learn more on the [S3 concept documentation](https://docs.stackable.tech/home/nightly/concepts/s3).",
								MarkdownDescription: "S3 connection definition as a resource. Learn more on the [S3 concept documentation](https://docs.stackable.tech/home/nightly/concepts/s3).",
								Attributes: map[string]schema.Attribute{
									"access_style": schema.StringAttribute{
										Description:         "Which access style to use. Defaults to virtual hosted-style as most of the data products out there. Have a look at the [AWS documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/VirtualHosting.html).",
										MarkdownDescription: "Which access style to use. Defaults to virtual hosted-style as most of the data products out there. Have a look at the [AWS documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/VirtualHosting.html).",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Path", "VirtualHosted"),
										},
									},

									"credentials": schema.SingleNestedAttribute{
										Description:         "If the S3 uses authentication you have to specify you S3 credentials. In the most cases a [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) providing 'accessKey' and 'secretKey' is sufficient.",
										MarkdownDescription: "If the S3 uses authentication you have to specify you S3 credentials. In the most cases a [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) providing 'accessKey' and 'secretKey' is sufficient.",
										Attributes: map[string]schema.Attribute{
											"scope": schema.SingleNestedAttribute{
												Description:         "[Scope](https://docs.stackable.tech/home/nightly/secret-operator/scope) of the [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass).",
												MarkdownDescription: "[Scope](https://docs.stackable.tech/home/nightly/secret-operator/scope) of the [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass).",
												Attributes: map[string]schema.Attribute{
													"listener_volumes": schema.ListAttribute{
														Description:         "The listener volume scope allows Node and Service scopes to be inferred from the applicable listeners. This must correspond to Volume names in the Pod that mount Listeners.",
														MarkdownDescription: "The listener volume scope allows Node and Service scopes to be inferred from the applicable listeners. This must correspond to Volume names in the Pod that mount Listeners.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"node": schema.BoolAttribute{
														Description:         "The node scope is resolved to the name of the Kubernetes Node object that the Pod is running on. This will typically be the DNS name of the node.",
														MarkdownDescription: "The node scope is resolved to the name of the Kubernetes Node object that the Pod is running on. This will typically be the DNS name of the node.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"pod": schema.BoolAttribute{
														Description:         "The pod scope is resolved to the name of the Kubernetes Pod. This allows the secret to differentiate between StatefulSet replicas.",
														MarkdownDescription: "The pod scope is resolved to the name of the Kubernetes Pod. This allows the secret to differentiate between StatefulSet replicas.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"services": schema.ListAttribute{
														Description:         "The service scope allows Pod objects to specify custom scopes. This should typically correspond to Service objects that the Pod participates in.",
														MarkdownDescription: "The service scope allows Pod objects to specify custom scopes. This should typically correspond to Service objects that the Pod participates in.",
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

											"secret_class": schema.StringAttribute{
												Description:         "[SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) containing the LDAP bind credentials.",
												MarkdownDescription: "[SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) containing the LDAP bind credentials.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"host": schema.StringAttribute{
										Description:         "Host of the S3 server without any protocol or port. For example: 'west1.my-cloud.com'.",
										MarkdownDescription: "Host of the S3 server without any protocol or port. For example: 'west1.my-cloud.com'.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"port": schema.Int64Attribute{
										Description:         "Port the S3 server listens on. If not specified the product will determine the port to use.",
										MarkdownDescription: "Port the S3 server listens on. If not specified the product will determine the port to use.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"tls": schema.SingleNestedAttribute{
										Description:         "Use a TLS connection. If not specified no TLS will be used.",
										MarkdownDescription: "Use a TLS connection. If not specified no TLS will be used.",
										Attributes: map[string]schema.Attribute{
											"verification": schema.SingleNestedAttribute{
												Description:         "The verification method used to verify the certificates of the server and/or the client.",
												MarkdownDescription: "The verification method used to verify the certificates of the server and/or the client.",
												Attributes: map[string]schema.Attribute{
													"none": schema.MapAttribute{
														Description:         "Use TLS but don't verify certificates.",
														MarkdownDescription: "Use TLS but don't verify certificates.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"server": schema.SingleNestedAttribute{
														Description:         "Use TLS and a CA certificate to verify the server.",
														MarkdownDescription: "Use TLS and a CA certificate to verify the server.",
														Attributes: map[string]schema.Attribute{
															"ca_cert": schema.SingleNestedAttribute{
																Description:         "CA cert to verify the server.",
																MarkdownDescription: "CA cert to verify the server.",
																Attributes: map[string]schema.Attribute{
																	"secret_class": schema.StringAttribute{
																		Description:         "Name of the [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) which will provide the CA certificate. Note that a SecretClass does not need to have a key but can also work with just a CA certificate, so if you got provided with a CA cert but don't have access to the key you can still use this method.",
																		MarkdownDescription: "Name of the [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) which will provide the CA certificate. Note that a SecretClass does not need to have a key but can also work with just a CA certificate, so if you got provided with a CA cert but don't have access to the key you can still use this method.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"web_pki": schema.MapAttribute{
																		Description:         "Use TLS and the CA certificates trusted by the common web browsers to verify the server. This can be useful when you e.g. use public AWS S3 or other public available services.",
																		MarkdownDescription: "Use TLS and the CA certificates trusted by the common web browsers to verify the server. This can be useful when you e.g. use public AWS S3 or other public available services.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"reference": schema.StringAttribute{
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

					"spark_conf": schema.MapAttribute{
						Description:         "A map of key/value strings that will be passed directly to spark-submit.",
						MarkdownDescription: "A map of key/value strings that will be passed directly to spark-submit.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"spark_image": schema.SingleNestedAttribute{
						Description:         "Specify which image to use, the easiest way is to only configure the 'productVersion'. You can also configure a custom image registry to pull from, as well as completely custom images. Consult the [Product image selection documentation](https://docs.stackable.tech/home/nightly/concepts/product_image_selection) for details.",
						MarkdownDescription: "Specify which image to use, the easiest way is to only configure the 'productVersion'. You can also configure a custom image registry to pull from, as well as completely custom images. Consult the [Product image selection documentation](https://docs.stackable.tech/home/nightly/concepts/product_image_selection) for details.",
						Attributes: map[string]schema.Attribute{
							"custom": schema.StringAttribute{
								Description:         "Overwrite the docker image. Specify the full docker image name, e.g. 'docker.stackable.tech/stackable/superset:1.4.1-stackable2.1.0'",
								MarkdownDescription: "Overwrite the docker image. Specify the full docker image name, e.g. 'docker.stackable.tech/stackable/superset:1.4.1-stackable2.1.0'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"product_version": schema.StringAttribute{
								Description:         "Version of the product, e.g. '1.4.1'.",
								MarkdownDescription: "Version of the product, e.g. '1.4.1'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pull_policy": schema.StringAttribute{
								Description:         "[Pull policy](https://kubernetes.io/docs/concepts/containers/images/#image-pull-policy) used when pulling the image.",
								MarkdownDescription: "[Pull policy](https://kubernetes.io/docs/concepts/containers/images/#image-pull-policy) used when pulling the image.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("IfNotPresent", "Always", "Never"),
								},
							},

							"pull_secrets": schema.ListNestedAttribute{
								Description:         "[Image pull secrets](https://kubernetes.io/docs/concepts/containers/images/#specifying-imagepullsecrets-on-a-pod) to pull images from a private registry.",
								MarkdownDescription: "[Image pull secrets](https://kubernetes.io/docs/concepts/containers/images/#specifying-imagepullsecrets-on-a-pod) to pull images from a private registry.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
											MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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

							"repo": schema.StringAttribute{
								Description:         "Name of the docker repo, e.g. 'docker.stackable.tech/stackable'",
								MarkdownDescription: "Name of the docker repo, e.g. 'docker.stackable.tech/stackable'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"stackable_version": schema.StringAttribute{
								Description:         "Stackable version of the product, e.g. '23.4', '23.4.1' or '0.0.0-dev'. If not specified, the operator will use its own version, e.g. '23.4.1'. When using a nightly operator or a pr version, it will use the nightly '0.0.0-dev' image.",
								MarkdownDescription: "Stackable version of the product, e.g. '23.4', '23.4.1' or '0.0.0-dev'. If not specified, the operator will use its own version, e.g. '23.4.1'. When using a nightly operator or a pr version, it will use the nightly '0.0.0-dev' image.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"vector_aggregator_config_map_name": schema.StringAttribute{
						Description:         "Name of the Vector aggregator [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery). It must contain the key 'ADDRESS' with the address of the Vector aggregator. Follow the [logging tutorial](https://docs.stackable.tech/home/nightly/tutorials/logging-vector-aggregator) to learn how to configure log aggregation with Vector.",
						MarkdownDescription: "Name of the Vector aggregator [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery). It must contain the key 'ADDRESS' with the address of the Vector aggregator. Follow the [logging tutorial](https://docs.stackable.tech/home/nightly/tutorials/logging-vector-aggregator) to learn how to configure log aggregation with Vector.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"volumes": schema.ListAttribute{
						Description:         "A list of volumes that can be made available to the job, driver or executors via their volume mounts.",
						MarkdownDescription: "A list of volumes that can be made available to the job, driver or executors via their volume mounts.",
						ElementType:         types.MapType{ElemType: types.StringType},
						Required:            false,
						Optional:            true,
						Computed:            false,
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *SparkStackableTechSparkApplicationV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_spark_stackable_tech_spark_application_v1alpha1_manifest")

	var model SparkStackableTechSparkApplicationV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("spark.stackable.tech/v1alpha1")
	model.Kind = pointer.String("SparkApplication")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
