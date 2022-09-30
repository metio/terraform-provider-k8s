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

type InfinispanOrgInfinispanV1Resource struct{}

var (
	_ resource.Resource = (*InfinispanOrgInfinispanV1Resource)(nil)
)

type InfinispanOrgInfinispanV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type InfinispanOrgInfinispanV1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		ConfigListener *struct {
			Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
		} `tfsdk:"config_listener" yaml:"configListener,omitempty"`

		ConfigMapName *string `tfsdk:"config_map_name" yaml:"configMapName,omitempty"`

		Container *struct {
			CliExtraJvmOpts *string `tfsdk:"cli_extra_jvm_opts" yaml:"cliExtraJvmOpts,omitempty"`

			Cpu *string `tfsdk:"cpu" yaml:"cpu,omitempty"`

			ExtraJvmOpts *string `tfsdk:"extra_jvm_opts" yaml:"extraJvmOpts,omitempty"`

			Memory *string `tfsdk:"memory" yaml:"memory,omitempty"`

			RouterExtraJvmOpts *string `tfsdk:"router_extra_jvm_opts" yaml:"routerExtraJvmOpts,omitempty"`
		} `tfsdk:"container" yaml:"container,omitempty"`

		Image *string `tfsdk:"image" yaml:"image,omitempty"`

		Logging *struct {
			Categories *map[string]string `tfsdk:"categories" yaml:"categories,omitempty"`
		} `tfsdk:"logging" yaml:"logging,omitempty"`

		Service *struct {
			Container *struct {
				StorageClassName *string `tfsdk:"storage_class_name" yaml:"storageClassName,omitempty"`

				EphemeralStorage *bool `tfsdk:"ephemeral_storage" yaml:"ephemeralStorage,omitempty"`

				Storage *string `tfsdk:"storage" yaml:"storage,omitempty"`
			} `tfsdk:"container" yaml:"container,omitempty"`

			ReplicationFactor *int64 `tfsdk:"replication_factor" yaml:"replicationFactor,omitempty"`

			Sites *struct {
				Local *struct {
					Discovery *struct {
						LaunchGossipRouter *bool `tfsdk:"launch_gossip_router" yaml:"launchGossipRouter,omitempty"`

						Type *string `tfsdk:"type" yaml:"type,omitempty"`
					} `tfsdk:"discovery" yaml:"discovery,omitempty"`

					Encryption *struct {
						Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`

						RouterKeyStore *struct {
							Alias *string `tfsdk:"alias" yaml:"alias,omitempty"`

							Filename *string `tfsdk:"filename" yaml:"filename,omitempty"`

							SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
						} `tfsdk:"router_key_store" yaml:"routerKeyStore,omitempty"`

						TransportKeyStore *struct {
							Filename *string `tfsdk:"filename" yaml:"filename,omitempty"`

							SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`

							Alias *string `tfsdk:"alias" yaml:"alias,omitempty"`
						} `tfsdk:"transport_key_store" yaml:"transportKeyStore,omitempty"`

						TrustStore *struct {
							Filename *string `tfsdk:"filename" yaml:"filename,omitempty"`

							SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
						} `tfsdk:"trust_store" yaml:"trustStore,omitempty"`
					} `tfsdk:"encryption" yaml:"encryption,omitempty"`

					Expose *struct {
						Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

						NodePort *int64 `tfsdk:"node_port" yaml:"nodePort,omitempty"`

						Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

						RouteHostName *string `tfsdk:"route_host_name" yaml:"routeHostName,omitempty"`

						Type *string `tfsdk:"type" yaml:"type,omitempty"`
					} `tfsdk:"expose" yaml:"expose,omitempty"`

					MaxRelayNodes *int64 `tfsdk:"max_relay_nodes" yaml:"maxRelayNodes,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"local" yaml:"local,omitempty"`

				Locations *[]struct {
					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

					Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

					SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`

					Url *string `tfsdk:"url" yaml:"url,omitempty"`

					ClusterName *string `tfsdk:"cluster_name" yaml:"clusterName,omitempty"`

					Host *string `tfsdk:"host" yaml:"host,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"locations" yaml:"locations,omitempty"`
			} `tfsdk:"sites" yaml:"sites,omitempty"`

			Type *string `tfsdk:"type" yaml:"type,omitempty"`
		} `tfsdk:"service" yaml:"service,omitempty"`

		Upgrades *struct {
			Type *string `tfsdk:"type" yaml:"type,omitempty"`
		} `tfsdk:"upgrades" yaml:"upgrades,omitempty"`

		Dependencies *struct {
			Artifacts *[]struct {
				Url *string `tfsdk:"url" yaml:"url,omitempty"`

				Hash *string `tfsdk:"hash" yaml:"hash,omitempty"`

				Maven *string `tfsdk:"maven" yaml:"maven,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`
			} `tfsdk:"artifacts" yaml:"artifacts,omitempty"`

			VolumeClaimName *string `tfsdk:"volume_claim_name" yaml:"volumeClaimName,omitempty"`
		} `tfsdk:"dependencies" yaml:"dependencies,omitempty"`

		Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`

		Autoscale *struct {
			MaxReplicas *int64 `tfsdk:"max_replicas" yaml:"maxReplicas,omitempty"`

			MinMemUsagePercent *int64 `tfsdk:"min_mem_usage_percent" yaml:"minMemUsagePercent,omitempty"`

			MinReplicas *int64 `tfsdk:"min_replicas" yaml:"minReplicas,omitempty"`

			Disabled *bool `tfsdk:"disabled" yaml:"disabled,omitempty"`

			MaxMemUsagePercent *int64 `tfsdk:"max_mem_usage_percent" yaml:"maxMemUsagePercent,omitempty"`
		} `tfsdk:"autoscale" yaml:"autoscale,omitempty"`

		Version *string `tfsdk:"version" yaml:"version,omitempty"`

		Affinity *struct {
			NodeAffinity *struct {
				PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
					Preference *struct {
						MatchFields *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"match_fields" yaml:"matchFields,omitempty"`

						MatchExpressions *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`
					} `tfsdk:"preference" yaml:"preference,omitempty"`

					Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`
				} `tfsdk:"preferred_during_scheduling_ignored_during_execution" yaml:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`

				RequiredDuringSchedulingIgnoredDuringExecution *struct {
					NodeSelectorTerms *[]struct {
						MatchExpressions *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

						MatchFields *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"match_fields" yaml:"matchFields,omitempty"`
					} `tfsdk:"node_selector_terms" yaml:"nodeSelectorTerms,omitempty"`
				} `tfsdk:"required_during_scheduling_ignored_during_execution" yaml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
			} `tfsdk:"node_affinity" yaml:"nodeAffinity,omitempty"`

			PodAffinity *struct {
				PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
					PodAffinityTerm *struct {
						Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

						TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`

						LabelSelector *struct {
							MatchExpressions *[]struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

								Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
							} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

							MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
						} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`
					} `tfsdk:"pod_affinity_term" yaml:"podAffinityTerm,omitempty"`

					Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`
				} `tfsdk:"preferred_during_scheduling_ignored_during_execution" yaml:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`

				RequiredDuringSchedulingIgnoredDuringExecution *[]struct {
					LabelSelector *struct {
						MatchExpressions *[]struct {
							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`

							Key *string `tfsdk:"key" yaml:"key,omitempty"`
						} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

						MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
					} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

					Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

					TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`
				} `tfsdk:"required_during_scheduling_ignored_during_execution" yaml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
			} `tfsdk:"pod_affinity" yaml:"podAffinity,omitempty"`

			PodAntiAffinity *struct {
				PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
					PodAffinityTerm *struct {
						LabelSelector *struct {
							MatchExpressions *[]struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

								Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
							} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

							MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
						} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

						Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

						TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`
					} `tfsdk:"pod_affinity_term" yaml:"podAffinityTerm,omitempty"`

					Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`
				} `tfsdk:"preferred_during_scheduling_ignored_during_execution" yaml:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`

				RequiredDuringSchedulingIgnoredDuringExecution *[]struct {
					LabelSelector *struct {
						MatchExpressions *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

						MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
					} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

					Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

					TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`
				} `tfsdk:"required_during_scheduling_ignored_during_execution" yaml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
			} `tfsdk:"pod_anti_affinity" yaml:"podAntiAffinity,omitempty"`
		} `tfsdk:"affinity" yaml:"affinity,omitempty"`

		CloudEvents *struct {
			CacheEntriesTopic *string `tfsdk:"cache_entries_topic" yaml:"cacheEntriesTopic,omitempty"`

			Acks *string `tfsdk:"acks" yaml:"acks,omitempty"`

			BootstrapServers *string `tfsdk:"bootstrap_servers" yaml:"bootstrapServers,omitempty"`
		} `tfsdk:"cloud_events" yaml:"cloudEvents,omitempty"`

		Expose *struct {
			NodePort *int64 `tfsdk:"node_port" yaml:"nodePort,omitempty"`

			Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

			Type *string `tfsdk:"type" yaml:"type,omitempty"`

			Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

			Host *string `tfsdk:"host" yaml:"host,omitempty"`
		} `tfsdk:"expose" yaml:"expose,omitempty"`

		Security *struct {
			EndpointEncryption *struct {
				ClientCertSecretName *string `tfsdk:"client_cert_secret_name" yaml:"clientCertSecretName,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`

				CertSecretName *string `tfsdk:"cert_secret_name" yaml:"certSecretName,omitempty"`

				CertServiceName *string `tfsdk:"cert_service_name" yaml:"certServiceName,omitempty"`

				ClientCert *string `tfsdk:"client_cert" yaml:"clientCert,omitempty"`
			} `tfsdk:"endpoint_encryption" yaml:"endpointEncryption,omitempty"`

			EndpointSecretName *string `tfsdk:"endpoint_secret_name" yaml:"endpointSecretName,omitempty"`

			Authorization *struct {
				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Roles *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Permissions *[]string `tfsdk:"permissions" yaml:"permissions,omitempty"`
				} `tfsdk:"roles" yaml:"roles,omitempty"`
			} `tfsdk:"authorization" yaml:"authorization,omitempty"`

			EndpointAuthentication *bool `tfsdk:"endpoint_authentication" yaml:"endpointAuthentication,omitempty"`
		} `tfsdk:"security" yaml:"security,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewInfinispanOrgInfinispanV1Resource() resource.Resource {
	return &InfinispanOrgInfinispanV1Resource{}
}

func (r *InfinispanOrgInfinispanV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_infinispan_org_infinispan_v1"
}

func (r *InfinispanOrgInfinispanV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Infinispan is the Schema for the infinispans API",
		MarkdownDescription: "Infinispan is the Schema for the infinispans API",
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
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.RequiresReplace(),
						},
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
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
				Description:         "InfinispanSpec defines the desired state of Infinispan",
				MarkdownDescription: "InfinispanSpec defines the desired state of Infinispan",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"config_listener": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"enabled": {
								Description:         "If true, a dedicated pod is used to ensure that all config resources created on the Infinispan server have a matching CR resource",
								MarkdownDescription: "If true, a dedicated pod is used to ensure that all config resources created on the Infinispan server have a matching CR resource",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"config_map_name": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"container": {
						Description:         "InfinispanContainerSpec specify resource requirements per container",
						MarkdownDescription: "InfinispanContainerSpec specify resource requirements per container",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"cli_extra_jvm_opts": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cpu": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"extra_jvm_opts": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"memory": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"router_extra_jvm_opts": {
								Description:         "",
								MarkdownDescription: "",

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

					"image": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"logging": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"categories": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"service": {
						Description:         "InfinispanServiceSpec specify configuration for specific service",
						MarkdownDescription: "InfinispanServiceSpec specify configuration for specific service",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"container": {
								Description:         "InfinispanServiceContainerSpec resource requirements specific for service",
								MarkdownDescription: "InfinispanServiceContainerSpec resource requirements specific for service",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"storage_class_name": {
										Description:         "The storage class object for persistent volume claims",
										MarkdownDescription: "The storage class object for persistent volume claims",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"ephemeral_storage": {
										Description:         "Enable/disable container ephemeral storage",
										MarkdownDescription: "Enable/disable container ephemeral storage",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"storage": {
										Description:         "The amount of storage for the persistent volume claim.",
										MarkdownDescription: "The amount of storage for the persistent volume claim.",

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

							"replication_factor": {
								Description:         "Cache replication factor, or number of copies for each entry.",
								MarkdownDescription: "Cache replication factor, or number of copies for each entry.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"sites": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"local": {
										Description:         "InfinispanSitesLocalSpec enables cross-site replication",
										MarkdownDescription: "InfinispanSitesLocalSpec enables cross-site replication",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"discovery": {
												Description:         "DiscoverySiteSpec configures the corss-site replication discovery",
												MarkdownDescription: "DiscoverySiteSpec configures the corss-site replication discovery",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"launch_gossip_router": {
														Description:         "Enables (default) or disables the Gossip Router pod and cross-site services",
														MarkdownDescription: "Enables (default) or disables the Gossip Router pod and cross-site services",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"type": {
														Description:         "Configures the discovery mode for cross-site replication",
														MarkdownDescription: "Configures the discovery mode for cross-site replication",

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

											"encryption": {
												Description:         "EncryptionSiteSpec enables TLS for cross-site replication",
												MarkdownDescription: "EncryptionSiteSpec enables TLS for cross-site replication",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"protocol": {
														Description:         "TLSProtocol specifies the TLS protocol",
														MarkdownDescription: "TLSProtocol specifies the TLS protocol",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"router_key_store": {
														Description:         "CrossSiteKeyStore keystore configuration for cross-site replication with TLS",
														MarkdownDescription: "CrossSiteKeyStore keystore configuration for cross-site replication with TLS",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"alias": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"filename": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},
														}),

														Required: true,
														Optional: false,
														Computed: false,
													},

													"transport_key_store": {
														Description:         "CrossSiteKeyStore keystore configuration for cross-site replication with TLS",
														MarkdownDescription: "CrossSiteKeyStore keystore configuration for cross-site replication with TLS",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"filename": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"alias": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: true,
														Optional: false,
														Computed: false,
													},

													"trust_store": {
														Description:         "CrossSiteTrustStore truststore configuration for cross-site replication with TLS",
														MarkdownDescription: "CrossSiteTrustStore truststore configuration for cross-site replication with TLS",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"filename": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"expose": {
												Description:         "CrossSiteExposeSpec describe how Infinispan Cross-Site service will be exposed externally",
												MarkdownDescription: "CrossSiteExposeSpec describe how Infinispan Cross-Site service will be exposed externally",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"node_port": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"port": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"route_host_name": {
														Description:         "RouteHostName optionally, specifies a custom hostname to be used by Openshift Route.",
														MarkdownDescription: "RouteHostName optionally, specifies a custom hostname to be used by Openshift Route.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"type": {
														Description:         "Type specifies different exposition methods for data grid",
														MarkdownDescription: "Type specifies different exposition methods for data grid",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"max_relay_nodes": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},

									"locations": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"namespace": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"port": {
												Description:         "Deprecated and to be removed on subsequent release. Use .URL with infinispan+xsite schema instead.",
												MarkdownDescription: "Deprecated and to be removed on subsequent release. Use .URL with infinispan+xsite schema instead.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_name": {
												Description:         "The access secret that allows backups to a remote site",
												MarkdownDescription: "The access secret that allows backups to a remote site",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"url": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"cluster_name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"host": {
												Description:         "Deprecated and to be removed on subsequent release. Use .URL with infinispan+xsite schema instead.",
												MarkdownDescription: "Deprecated and to be removed on subsequent release. Use .URL with infinispan+xsite schema instead.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": {
								Description:         "The service type",
								MarkdownDescription: "The service type",

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

					"upgrades": {
						Description:         "Strategy to use when doing upgrades",
						MarkdownDescription: "Strategy to use when doing upgrades",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"type": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"dependencies": {
						Description:         "External dependencies needed by the Infinispan cluster",
						MarkdownDescription: "External dependencies needed by the Infinispan cluster",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"artifacts": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"url": {
										Description:         "URL of the file you want to download.",
										MarkdownDescription: "URL of the file you want to download.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"hash": {
										Description:         "Checksum that you can use to verify downloaded files.",
										MarkdownDescription: "Checksum that you can use to verify downloaded files.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"maven": {
										Description:         "Coordinates of a maven artifact in the 'groupId:artifactId:version' format, for example 'org.postgresql:postgresql:42.3.1'.",
										MarkdownDescription: "Coordinates of a maven artifact in the 'groupId:artifactId:version' format, for example 'org.postgresql:postgresql:42.3.1'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": {
										Description:         "Deprecated, no longer has any effect. Specifies the type of file you want to download. If not specified, the file type is automatically determined from the extension.",
										MarkdownDescription: "Deprecated, no longer has any effect. Specifies the type of file you want to download. If not specified, the file type is automatically determined from the extension.",

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

							"volume_claim_name": {
								Description:         "The Persistent Volume Claim that holds custom libraries",
								MarkdownDescription: "The Persistent Volume Claim that holds custom libraries",

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

					"replicas": {
						Description:         "The number of nodes in the Infinispan cluster.",
						MarkdownDescription: "The number of nodes in the Infinispan cluster.",

						Type: types.Int64Type,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"autoscale": {
						Description:         "Autoscale describe autoscaling configuration for the cluster",
						MarkdownDescription: "Autoscale describe autoscaling configuration for the cluster",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"max_replicas": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"min_mem_usage_percent": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"min_replicas": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"disabled": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"max_mem_usage_percent": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"version": {
						Description:         "The semantic version of the Infinispan cluster.",
						MarkdownDescription: "The semantic version of the Infinispan cluster.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"affinity": {
						Description:         "Affinity is a group of affinity scheduling rules.",
						MarkdownDescription: "Affinity is a group of affinity scheduling rules.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"node_affinity": {
								Description:         "Describes node affinity scheduling rules for the pod.",
								MarkdownDescription: "Describes node affinity scheduling rules for the pod.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"preferred_during_scheduling_ignored_during_execution": {
										Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",
										MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"preference": {
												Description:         "A node selector term, associated with the corresponding weight.",
												MarkdownDescription: "A node selector term, associated with the corresponding weight.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"match_fields": {
														Description:         "A list of node selector requirements by node's fields.",
														MarkdownDescription: "A list of node selector requirements by node's fields.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The label key that the selector applies to.",
																MarkdownDescription: "The label key that the selector applies to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"operator": {
																Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"values": {
																Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"match_expressions": {
														Description:         "A list of node selector requirements by node's labels.",
														MarkdownDescription: "A list of node selector requirements by node's labels.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The label key that the selector applies to.",
																MarkdownDescription: "The label key that the selector applies to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"operator": {
																Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"values": {
																Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",

																Type: types.ListType{ElemType: types.StringType},

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

											"weight": {
												Description:         "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
												MarkdownDescription: "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"required_during_scheduling_ignored_during_execution": {
										Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",
										MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"node_selector_terms": {
												Description:         "Required. A list of node selector terms. The terms are ORed.",
												MarkdownDescription: "Required. A list of node selector terms. The terms are ORed.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"match_expressions": {
														Description:         "A list of node selector requirements by node's labels.",
														MarkdownDescription: "A list of node selector requirements by node's labels.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The label key that the selector applies to.",
																MarkdownDescription: "The label key that the selector applies to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"operator": {
																Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"values": {
																Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"match_fields": {
														Description:         "A list of node selector requirements by node's fields.",
														MarkdownDescription: "A list of node selector requirements by node's fields.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The label key that the selector applies to.",
																MarkdownDescription: "The label key that the selector applies to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"operator": {
																Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"values": {
																Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",

																Type: types.ListType{ElemType: types.StringType},

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pod_affinity": {
								Description:         "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
								MarkdownDescription: "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"preferred_during_scheduling_ignored_during_execution": {
										Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
										MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"pod_affinity_term": {
												Description:         "Required. A pod affinity term, associated with the corresponding weight.",
												MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"namespaces": {
														Description:         "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'",
														MarkdownDescription: "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"topology_key": {
														Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
														MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"label_selector": {
														Description:         "A label query over a set of resources, in this case pods.",
														MarkdownDescription: "A label query over a set of resources, in this case pods.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"match_expressions": {
																Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"operator": {
																		Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"values": {
																		Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"match_labels": {
																Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

																Type: types.MapType{ElemType: types.StringType},

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

											"weight": {
												Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
												MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"required_during_scheduling_ignored_during_execution": {
										Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
										MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"label_selector": {
												Description:         "A label query over a set of resources, in this case pods.",
												MarkdownDescription: "A label query over a set of resources, in this case pods.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"match_expressions": {
														Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
														MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"operator": {
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"values": {
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"key": {
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"match_labels": {
														Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
														MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespaces": {
												Description:         "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'",
												MarkdownDescription: "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"topology_key": {
												Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
												MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pod_anti_affinity": {
								Description:         "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
								MarkdownDescription: "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"preferred_during_scheduling_ignored_during_execution": {
										Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
										MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"pod_affinity_term": {
												Description:         "Required. A pod affinity term, associated with the corresponding weight.",
												MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"label_selector": {
														Description:         "A label query over a set of resources, in this case pods.",
														MarkdownDescription: "A label query over a set of resources, in this case pods.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"match_expressions": {
																Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"operator": {
																		Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"values": {
																		Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"match_labels": {
																Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespaces": {
														Description:         "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'",
														MarkdownDescription: "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"topology_key": {
														Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
														MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"weight": {
												Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
												MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"required_during_scheduling_ignored_during_execution": {
										Description:         "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
										MarkdownDescription: "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"label_selector": {
												Description:         "A label query over a set of resources, in this case pods.",
												MarkdownDescription: "A label query over a set of resources, in this case pods.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"match_expressions": {
														Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
														MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"operator": {
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"values": {
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"match_labels": {
														Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
														MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespaces": {
												Description:         "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'",
												MarkdownDescription: "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"topology_key": {
												Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
												MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

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

						Required: false,
						Optional: true,
						Computed: false,
					},

					"cloud_events": {
						Description:         "InfinispanCloudEvents describes how Infinispan is connected with Cloud Event, see Kafka docs for more info",
						MarkdownDescription: "InfinispanCloudEvents describes how Infinispan is connected with Cloud Event, see Kafka docs for more info",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"cache_entries_topic": {
								Description:         "CacheEntriesTopic is the name of the topic on which events will be published",
								MarkdownDescription: "CacheEntriesTopic is the name of the topic on which events will be published",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"acks": {
								Description:         "Acks configuration for the producer ack-value",
								MarkdownDescription: "Acks configuration for the producer ack-value",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"bootstrap_servers": {
								Description:         "BootstrapServers is comma separated list of boostrap server:port addresses",
								MarkdownDescription: "BootstrapServers is comma separated list of boostrap server:port addresses",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"expose": {
						Description:         "ExposeSpec describe how Infinispan will be exposed externally",
						MarkdownDescription: "ExposeSpec describe how Infinispan will be exposed externally",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"node_port": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": {
								Description:         "Type specifies different exposition methods for data grid",
								MarkdownDescription: "Type specifies different exposition methods for data grid",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"annotations": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"host": {
								Description:         "The network hostname for your Infinispan cluster",
								MarkdownDescription: "The network hostname for your Infinispan cluster",

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

					"security": {
						Description:         "InfinispanSecurity info for the user application connection",
						MarkdownDescription: "InfinispanSecurity info for the user application connection",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"endpoint_encryption": {
								Description:         "EndpointEncryption configuration",
								MarkdownDescription: "EndpointEncryption configuration",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"client_cert_secret_name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": {
										Description:         "Disable or modify endpoint encryption.",
										MarkdownDescription: "Disable or modify endpoint encryption.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"cert_secret_name": {
										Description:         "The secret that contains TLS certificates",
										MarkdownDescription: "The secret that contains TLS certificates",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"cert_service_name": {
										Description:         "A service that provides TLS certificates",
										MarkdownDescription: "A service that provides TLS certificates",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"client_cert": {
										Description:         "ClientCertType specifies a client certificate validation mechanism.",
										MarkdownDescription: "ClientCertType specifies a client certificate validation mechanism.",

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

							"endpoint_secret_name": {
								Description:         "The secret that contains user credentials.",
								MarkdownDescription: "The secret that contains user credentials.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"authorization": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"enabled": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"roles": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"permissions": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"endpoint_authentication": {
								Description:         "Enable or disable user authentication",
								MarkdownDescription: "Enable or disable user authentication",

								Type: types.BoolType,

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

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *InfinispanOrgInfinispanV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_infinispan_org_infinispan_v1")

	var state InfinispanOrgInfinispanV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel InfinispanOrgInfinispanV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("infinispan.org/v1")
	goModel.Kind = utilities.Ptr("Infinispan")

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

func (r *InfinispanOrgInfinispanV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_infinispan_org_infinispan_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *InfinispanOrgInfinispanV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_infinispan_org_infinispan_v1")

	var state InfinispanOrgInfinispanV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel InfinispanOrgInfinispanV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("infinispan.org/v1")
	goModel.Kind = utilities.Ptr("Infinispan")

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

func (r *InfinispanOrgInfinispanV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_infinispan_org_infinispan_v1")
	// NO-OP: Terraform removes the state automatically for us
}
