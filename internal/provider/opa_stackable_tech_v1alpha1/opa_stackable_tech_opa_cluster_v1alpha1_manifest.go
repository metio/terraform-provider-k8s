/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package opa_stackable_tech_v1alpha1

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
	_ datasource.DataSource = &OpaStackableTechOpaClusterV1Alpha1Manifest{}
)

func NewOpaStackableTechOpaClusterV1Alpha1Manifest() datasource.DataSource {
	return &OpaStackableTechOpaClusterV1Alpha1Manifest{}
}

type OpaStackableTechOpaClusterV1Alpha1Manifest struct{}

type OpaStackableTechOpaClusterV1Alpha1ManifestData struct {
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
		ClusterConfig *struct {
			ListenerClass *string `tfsdk:"listener_class" json:"listenerClass,omitempty"`
			UserInfo      *struct {
				Backend *struct {
					ExperimentalXfscAas *struct {
						Hostname *string `tfsdk:"hostname" json:"hostname,omitempty"`
						Port     *int64  `tfsdk:"port" json:"port,omitempty"`
					} `tfsdk:"experimental_xfsc_aas" json:"experimentalXfscAas,omitempty"`
					Keycloak *struct {
						AdminRealm              *string `tfsdk:"admin_realm" json:"adminRealm,omitempty"`
						ClientCredentialsSecret *string `tfsdk:"client_credentials_secret" json:"clientCredentialsSecret,omitempty"`
						Hostname                *string `tfsdk:"hostname" json:"hostname,omitempty"`
						Port                    *int64  `tfsdk:"port" json:"port,omitempty"`
						RootPath                *string `tfsdk:"root_path" json:"rootPath,omitempty"`
						Tls                     *struct {
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
						UserRealm *string `tfsdk:"user_realm" json:"userRealm,omitempty"`
					} `tfsdk:"keycloak" json:"keycloak,omitempty"`
					None *map[string]string `tfsdk:"none" json:"none,omitempty"`
				} `tfsdk:"backend" json:"backend,omitempty"`
				Cache *struct {
					EntryTimeToLive *string `tfsdk:"entry_time_to_live" json:"entryTimeToLive,omitempty"`
				} `tfsdk:"cache" json:"cache,omitempty"`
			} `tfsdk:"user_info" json:"userInfo,omitempty"`
			VectorAggregatorConfigMapName *string `tfsdk:"vector_aggregator_config_map_name" json:"vectorAggregatorConfigMapName,omitempty"`
		} `tfsdk:"cluster_config" json:"clusterConfig,omitempty"`
		ClusterOperation *struct {
			ReconciliationPaused *bool `tfsdk:"reconciliation_paused" json:"reconciliationPaused,omitempty"`
			Stopped              *bool `tfsdk:"stopped" json:"stopped,omitempty"`
		} `tfsdk:"cluster_operation" json:"clusterOperation,omitempty"`
		Image *struct {
			Custom         *string `tfsdk:"custom" json:"custom,omitempty"`
			ProductVersion *string `tfsdk:"product_version" json:"productVersion,omitempty"`
			PullPolicy     *string `tfsdk:"pull_policy" json:"pullPolicy,omitempty"`
			PullSecrets    *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"pull_secrets" json:"pullSecrets,omitempty"`
			Repo             *string `tfsdk:"repo" json:"repo,omitempty"`
			StackableVersion *string `tfsdk:"stackable_version" json:"stackableVersion,omitempty"`
		} `tfsdk:"image" json:"image,omitempty"`
		Servers *struct {
			CliOverrides *map[string]string `tfsdk:"cli_overrides" json:"cliOverrides,omitempty"`
			Config       *struct {
				Affinity *struct {
					NodeAffinity    *map[string]string `tfsdk:"node_affinity" json:"nodeAffinity,omitempty"`
					NodeSelector    *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
					PodAffinity     *map[string]string `tfsdk:"pod_affinity" json:"podAffinity,omitempty"`
					PodAntiAffinity *map[string]string `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
				} `tfsdk:"affinity" json:"affinity,omitempty"`
				GracefulShutdownTimeout *string `tfsdk:"graceful_shutdown_timeout" json:"gracefulShutdownTimeout,omitempty"`
				Logging                 *struct {
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
			} `tfsdk:"config" json:"config,omitempty"`
			ConfigOverrides *map[string]map[string]string `tfsdk:"config_overrides" json:"configOverrides,omitempty"`
			EnvOverrides    *map[string]string            `tfsdk:"env_overrides" json:"envOverrides,omitempty"`
			PodOverrides    *map[string]string            `tfsdk:"pod_overrides" json:"podOverrides,omitempty"`
			RoleConfig      *map[string]string            `tfsdk:"role_config" json:"roleConfig,omitempty"`
			RoleGroups      *struct {
				CliOverrides *map[string]string `tfsdk:"cli_overrides" json:"cliOverrides,omitempty"`
				Config       *struct {
					Affinity *struct {
						NodeAffinity    *map[string]string `tfsdk:"node_affinity" json:"nodeAffinity,omitempty"`
						NodeSelector    *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
						PodAffinity     *map[string]string `tfsdk:"pod_affinity" json:"podAffinity,omitempty"`
						PodAntiAffinity *map[string]string `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
					} `tfsdk:"affinity" json:"affinity,omitempty"`
					GracefulShutdownTimeout *string `tfsdk:"graceful_shutdown_timeout" json:"gracefulShutdownTimeout,omitempty"`
					Logging                 *struct {
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
				} `tfsdk:"config" json:"config,omitempty"`
				ConfigOverrides *map[string]map[string]string `tfsdk:"config_overrides" json:"configOverrides,omitempty"`
				EnvOverrides    *map[string]string            `tfsdk:"env_overrides" json:"envOverrides,omitempty"`
				PodOverrides    *map[string]string            `tfsdk:"pod_overrides" json:"podOverrides,omitempty"`
				Replicas        *int64                        `tfsdk:"replicas" json:"replicas,omitempty"`
			} `tfsdk:"role_groups" json:"roleGroups,omitempty"`
		} `tfsdk:"servers" json:"servers,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OpaStackableTechOpaClusterV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_opa_stackable_tech_opa_cluster_v1alpha1_manifest"
}

func (r *OpaStackableTechOpaClusterV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Auto-generated derived type for OpaSpec via 'CustomResource'",
		MarkdownDescription: "Auto-generated derived type for OpaSpec via 'CustomResource'",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"cluster_config": schema.SingleNestedAttribute{
						Description:         "Global OPA cluster configuration that applies to all roles and role groups.",
						MarkdownDescription: "Global OPA cluster configuration that applies to all roles and role groups.",
						Attributes: map[string]schema.Attribute{
							"listener_class": schema.StringAttribute{
								Description:         "This field controls which type of Service the operator creates for this OpaCluster: * cluster-internal: Use a ClusterIP service * external-unstable: Use a NodePort service * external-stable: Use a LoadBalancer service This is a temporary solution with the goal to keep yaml manifests forward compatible. In the future, this setting will control which ListenerClass <https://docs.stackable.tech/home/stable/listener-operator/listenerclass.html> will be used to expose the service, and ListenerClass names will stay the same, allowing for a non-breaking change.",
								MarkdownDescription: "This field controls which type of Service the operator creates for this OpaCluster: * cluster-internal: Use a ClusterIP service * external-unstable: Use a NodePort service * external-stable: Use a LoadBalancer service This is a temporary solution with the goal to keep yaml manifests forward compatible. In the future, this setting will control which ListenerClass <https://docs.stackable.tech/home/stable/listener-operator/listenerclass.html> will be used to expose the service, and ListenerClass names will stay the same, allowing for a non-breaking change.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("cluster-internal", "external-unstable", "external-stable"),
								},
							},

							"user_info": schema.SingleNestedAttribute{
								Description:         "Configures how to fetch additional metadata about users (such as group memberships) from an external directory service.",
								MarkdownDescription: "Configures how to fetch additional metadata about users (such as group memberships) from an external directory service.",
								Attributes: map[string]schema.Attribute{
									"backend": schema.SingleNestedAttribute{
										Description:         "The backend directory service to use.",
										MarkdownDescription: "The backend directory service to use.",
										Attributes: map[string]schema.Attribute{
											"experimental_xfsc_aas": schema.SingleNestedAttribute{
												Description:         "Backend that fetches user information from the Gaia-X Cross Federation Services Components (XFSC) Authentication & Authorization Service.",
												MarkdownDescription: "Backend that fetches user information from the Gaia-X Cross Federation Services Components (XFSC) Authentication & Authorization Service.",
												Attributes: map[string]schema.Attribute{
													"hostname": schema.StringAttribute{
														Description:         "Hostname of the identity provider, e.g. 'my.aas.corp'.",
														MarkdownDescription: "Hostname of the identity provider, e.g. 'my.aas.corp'.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"port": schema.Int64Attribute{
														Description:         "Port of the identity provider. Defaults to port 5000.",
														MarkdownDescription: "Port of the identity provider. Defaults to port 5000.",
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

											"keycloak": schema.SingleNestedAttribute{
												Description:         "Backend that fetches user information from Keycloak.",
												MarkdownDescription: "Backend that fetches user information from Keycloak.",
												Attributes: map[string]schema.Attribute{
													"admin_realm": schema.StringAttribute{
														Description:         "The Keycloak realm that OPA's Keycloak account (as specified by 'credentialsSecretName' exists in). Typically 'master'.",
														MarkdownDescription: "The Keycloak realm that OPA's Keycloak account (as specified by 'credentialsSecretName' exists in). Typically 'master'.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"client_credentials_secret": schema.StringAttribute{
														Description:         "Name of a Secret that contains client credentials of a Keycloak account with permission to read user metadata. Must contain the fields 'clientId' and 'clientSecret'.",
														MarkdownDescription: "Name of a Secret that contains client credentials of a Keycloak account with permission to read user metadata. Must contain the fields 'clientId' and 'clientSecret'.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"hostname": schema.StringAttribute{
														Description:         "Hostname of the identity provider, e.g. 'my.keycloak.corp'.",
														MarkdownDescription: "Hostname of the identity provider, e.g. 'my.keycloak.corp'.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"port": schema.Int64Attribute{
														Description:         "Port of the identity provider. If TLS is used defaults to '443', otherwise to '80'.",
														MarkdownDescription: "Port of the identity provider. If TLS is used defaults to '443', otherwise to '80'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(0),
														},
													},

													"root_path": schema.StringAttribute{
														Description:         "Root HTTP path of the identity provider. Defaults to '/'.",
														MarkdownDescription: "Root HTTP path of the identity provider. Defaults to '/'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
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

													"user_realm": schema.StringAttribute{
														Description:         "The Keycloak realm that user metadata should be resolved from.",
														MarkdownDescription: "The Keycloak realm that user metadata should be resolved from.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"none": schema.MapAttribute{
												Description:         "Dummy backend that adds no extra user information.",
												MarkdownDescription: "Dummy backend that adds no extra user information.",
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

									"cache": schema.SingleNestedAttribute{
										Description:         "Caching configuration.",
										MarkdownDescription: "Caching configuration.",
										Attributes: map[string]schema.Attribute{
											"entry_time_to_live": schema.StringAttribute{
												Description:         "How long metadata about each user should be cached for.",
												MarkdownDescription: "How long metadata about each user should be cached for.",
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

							"vector_aggregator_config_map_name": schema.StringAttribute{
								Description:         "Name of the Vector aggregator discovery ConfigMap. It must contain the key 'ADDRESS' with the address of the Vector aggregator.",
								MarkdownDescription: "Name of the Vector aggregator discovery ConfigMap. It must contain the key 'ADDRESS' with the address of the Vector aggregator.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"cluster_operation": schema.SingleNestedAttribute{
						Description:         "Cluster operations like pause reconciliation or cluster stop.",
						MarkdownDescription: "Cluster operations like pause reconciliation or cluster stop.",
						Attributes: map[string]schema.Attribute{
							"reconciliation_paused": schema.BoolAttribute{
								Description:         "Flag to stop cluster reconciliation by the operator. This means that all changes in the custom resource spec are ignored until this flag is set to false or removed. The operator will however still watch the deployed resources at the time and update the custom resource status field. If applied at the same time with 'stopped', 'reconciliationPaused' will take precedence over 'stopped' and stop the reconciliation immediately.",
								MarkdownDescription: "Flag to stop cluster reconciliation by the operator. This means that all changes in the custom resource spec are ignored until this flag is set to false or removed. The operator will however still watch the deployed resources at the time and update the custom resource status field. If applied at the same time with 'stopped', 'reconciliationPaused' will take precedence over 'stopped' and stop the reconciliation immediately.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"stopped": schema.BoolAttribute{
								Description:         "Flag to stop the cluster. This means all deployed resources (e.g. Services, StatefulSets, ConfigMaps) are kept but all deployed Pods (e.g. replicas from a StatefulSet) are scaled to 0 and therefore stopped and removed. If applied at the same time with 'reconciliationPaused', the latter will pause reconciliation and 'stopped' will take no effect until 'reconciliationPaused' is set to false or removed.",
								MarkdownDescription: "Flag to stop the cluster. This means all deployed resources (e.g. Services, StatefulSets, ConfigMaps) are kept but all deployed Pods (e.g. replicas from a StatefulSet) are scaled to 0 and therefore stopped and removed. If applied at the same time with 'reconciliationPaused', the latter will pause reconciliation and 'stopped' will take no effect until 'reconciliationPaused' is set to false or removed.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"image": schema.SingleNestedAttribute{
						Description:         "The OPA image to use",
						MarkdownDescription: "The OPA image to use",
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

					"servers": schema.SingleNestedAttribute{
						Description:         "OPA server configuration.",
						MarkdownDescription: "OPA server configuration.",
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

									"graceful_shutdown_timeout": schema.StringAttribute{
										Description:         "Time period Pods have to gracefully shut down, e.g. '30m', '1h' or '2d'. Consult the operator documentation for details.",
										MarkdownDescription: "Time period Pods have to gracefully shut down, e.g. '30m', '1h' or '2d'. Consult the operator documentation for details.",
										Required:            false,
										Optional:            true,
										Computed:            false,
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

							"role_config": schema.MapAttribute{
								Description:         "This is a product-agnostic RoleConfig, with nothing in it. It is used e.g. by products that have nothing configurable at role level.",
								MarkdownDescription: "This is a product-agnostic RoleConfig, with nothing in it. It is used e.g. by products that have nothing configurable at role level.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"role_groups": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
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

											"graceful_shutdown_timeout": schema.StringAttribute{
												Description:         "Time period Pods have to gracefully shut down, e.g. '30m', '1h' or '2d'. Consult the operator documentation for details.",
												MarkdownDescription: "Time period Pods have to gracefully shut down, e.g. '30m', '1h' or '2d'. Consult the operator documentation for details.",
												Required:            false,
												Optional:            true,
												Computed:            false,
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
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *OpaStackableTechOpaClusterV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_opa_stackable_tech_opa_cluster_v1alpha1_manifest")

	var model OpaStackableTechOpaClusterV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("opa.stackable.tech/v1alpha1")
	model.Kind = pointer.String("OpaCluster")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
