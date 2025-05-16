/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package hyperspike_io_v1

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
	_ datasource.DataSource = &HyperspikeIoValkeyV1Manifest{}
)

func NewHyperspikeIoValkeyV1Manifest() datasource.DataSource {
	return &HyperspikeIoValkeyV1Manifest{}
}

type HyperspikeIoValkeyV1Manifest struct{}

type HyperspikeIoValkeyV1ManifestData struct {
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
		AnonymousAuth                *bool   `tfsdk:"anonymous_auth" json:"anonymousAuth,omitempty"`
		CertIssuer                   *string `tfsdk:"cert_issuer" json:"certIssuer,omitempty"`
		CertIssuerType               *string `tfsdk:"cert_issuer_type" json:"certIssuerType,omitempty"`
		ClusterDomain                *string `tfsdk:"cluster_domain" json:"clusterDomain,omitempty"`
		ClusterPreferredEndpointType *string `tfsdk:"cluster_preferred_endpoint_type" json:"clusterPreferredEndpointType,omitempty"`
		ExporterImage                *string `tfsdk:"exporter_image" json:"exporterImage,omitempty"`
		ExternalAccess               *struct {
			CertIssuer     *string `tfsdk:"cert_issuer" json:"certIssuer,omitempty"`
			CertIssuerType *string `tfsdk:"cert_issuer_type" json:"certIssuerType,omitempty"`
			Enabled        *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			ExternalDNS    *bool   `tfsdk:"external_dns" json:"externalDNS,omitempty"`
			LoadBalancer   *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			} `tfsdk:"load_balancer" json:"loadBalancer,omitempty"`
			Proxy *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				ExtraConfig *string            `tfsdk:"extra_config" json:"extraConfig,omitempty"`
				Hostname    *string            `tfsdk:"hostname" json:"hostname,omitempty"`
				Image       *string            `tfsdk:"image" json:"image,omitempty"`
				Replicas    *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
				Resources   *struct {
					Claims *[]struct {
						Name    *string `tfsdk:"name" json:"name,omitempty"`
						Request *string `tfsdk:"request" json:"request,omitempty"`
					} `tfsdk:"claims" json:"claims,omitempty"`
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
			} `tfsdk:"proxy" json:"proxy,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"external_access" json:"externalAccess,omitempty"`
		Image            *string            `tfsdk:"image" json:"image,omitempty"`
		NodeSelector     *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
		Nodes            *int64             `tfsdk:"nodes" json:"nodes,omitempty"`
		Prometheus       *bool              `tfsdk:"prometheus" json:"prometheus,omitempty"`
		PrometheusLabels *map[string]string `tfsdk:"prometheus_labels" json:"prometheusLabels,omitempty"`
		Replicas         *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
		Resources        *struct {
			Claims *[]struct {
				Name    *string `tfsdk:"name" json:"name,omitempty"`
				Request *string `tfsdk:"request" json:"request,omitempty"`
			} `tfsdk:"claims" json:"claims,omitempty"`
			Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
			Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"resources" json:"resources,omitempty"`
		ServicePassword *struct {
			Key      *string `tfsdk:"key" json:"key,omitempty"`
			Name     *string `tfsdk:"name" json:"name,omitempty"`
			Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
		} `tfsdk:"service_password" json:"servicePassword,omitempty"`
		Storage *struct {
			ApiVersion *string            `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind       *string            `tfsdk:"kind" json:"kind,omitempty"`
			Metadata   *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
			Spec       *struct {
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
				StorageClassName          *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
				VolumeAttributesClassName *string `tfsdk:"volume_attributes_class_name" json:"volumeAttributesClassName,omitempty"`
				VolumeMode                *string `tfsdk:"volume_mode" json:"volumeMode,omitempty"`
				VolumeName                *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
			Status *struct {
				AccessModes               *[]string          `tfsdk:"access_modes" json:"accessModes,omitempty"`
				AllocatedResourceStatuses *map[string]string `tfsdk:"allocated_resource_statuses" json:"allocatedResourceStatuses,omitempty"`
				AllocatedResources        *map[string]string `tfsdk:"allocated_resources" json:"allocatedResources,omitempty"`
				Capacity                  *map[string]string `tfsdk:"capacity" json:"capacity,omitempty"`
				Conditions                *[]struct {
					LastProbeTime      *string `tfsdk:"last_probe_time" json:"lastProbeTime,omitempty"`
					LastTransitionTime *string `tfsdk:"last_transition_time" json:"lastTransitionTime,omitempty"`
					Message            *string `tfsdk:"message" json:"message,omitempty"`
					Reason             *string `tfsdk:"reason" json:"reason,omitempty"`
					Status             *string `tfsdk:"status" json:"status,omitempty"`
					Type               *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"conditions" json:"conditions,omitempty"`
				CurrentVolumeAttributesClassName *string `tfsdk:"current_volume_attributes_class_name" json:"currentVolumeAttributesClassName,omitempty"`
				ModifyVolumeStatus               *struct {
					Status                          *string `tfsdk:"status" json:"status,omitempty"`
					TargetVolumeAttributesClassName *string `tfsdk:"target_volume_attributes_class_name" json:"targetVolumeAttributesClassName,omitempty"`
				} `tfsdk:"modify_volume_status" json:"modifyVolumeStatus,omitempty"`
				Phase *string `tfsdk:"phase" json:"phase,omitempty"`
			} `tfsdk:"status" json:"status,omitempty"`
		} `tfsdk:"storage" json:"storage,omitempty"`
		Tls         *bool `tfsdk:"tls" json:"tls,omitempty"`
		Tolerations *[]struct {
			Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
			Key               *string `tfsdk:"key" json:"key,omitempty"`
			Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
			TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
			Value             *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tolerations" json:"tolerations,omitempty"`
		VolumePermissions *bool `tfsdk:"volume_permissions" json:"volumePermissions,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *HyperspikeIoValkeyV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_hyperspike_io_valkey_v1_manifest"
}

func (r *HyperspikeIoValkeyV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Valkey is the Schema for the valkeys API",
		MarkdownDescription: "Valkey is the Schema for the valkeys API",
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
				Description:         "ValkeySpec defines the desired state of Valkey",
				MarkdownDescription: "ValkeySpec defines the desired state of Valkey",
				Attributes: map[string]schema.Attribute{
					"anonymous_auth": schema.BoolAttribute{
						Description:         "Anonymous Auth",
						MarkdownDescription: "Anonymous Auth",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"cert_issuer": schema.StringAttribute{
						Description:         "Certificate Issuer",
						MarkdownDescription: "Certificate Issuer",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cert_issuer_type": schema.StringAttribute{
						Description:         "Certificate Issuer Type",
						MarkdownDescription: "Certificate Issuer Type",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("ClusterIssuer", "Issuer"),
						},
					},

					"cluster_domain": schema.StringAttribute{
						Description:         "Cluster Domain - used for DNS",
						MarkdownDescription: "Cluster Domain - used for DNS",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"cluster_preferred_endpoint_type": schema.StringAttribute{
						Description:         "Which endpoint is shown as the preferred endpoint valid values are 'ip', 'hostname', or 'unknown-endpoint'.",
						MarkdownDescription: "Which endpoint is shown as the preferred endpoint valid values are 'ip', 'hostname', or 'unknown-endpoint'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("ip", "hostname", "unknown-endpoint"),
						},
					},

					"exporter_image": schema.StringAttribute{
						Description:         "Exporter Image to use",
						MarkdownDescription: "Exporter Image to use",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"external_access": schema.SingleNestedAttribute{
						Description:         "External access configuration",
						MarkdownDescription: "External access configuration",
						Attributes: map[string]schema.Attribute{
							"cert_issuer": schema.StringAttribute{
								Description:         "Cert Issuer for external access TLS certificate",
								MarkdownDescription: "Cert Issuer for external access TLS certificate",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cert_issuer_type": schema.StringAttribute{
								Description:         "Cert Issuer Type for external access TLS certificate",
								MarkdownDescription: "Cert Issuer Type for external access TLS certificate",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("ClusterIssuer", "Issuer"),
								},
							},

							"enabled": schema.BoolAttribute{
								Description:         "Enable external access",
								MarkdownDescription: "Enable external access",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"external_dns": schema.BoolAttribute{
								Description:         "Support External DNS",
								MarkdownDescription: "Support External DNS",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"load_balancer": schema.SingleNestedAttribute{
								Description:         "LoadBalancer Settings",
								MarkdownDescription: "LoadBalancer Settings",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations for the load balancer service",
										MarkdownDescription: "Annotations for the load balancer service",
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

							"proxy": schema.SingleNestedAttribute{
								Description:         "Proxy Settings",
								MarkdownDescription: "Proxy Settings",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations for the proxy service",
										MarkdownDescription: "Annotations for the proxy service",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"extra_config": schema.StringAttribute{
										Description:         "Extra Envoy configuration",
										MarkdownDescription: "Extra Envoy configuration",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"hostname": schema.StringAttribute{
										Description:         "External Hostname for the proxy",
										MarkdownDescription: "External Hostname for the proxy",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"image": schema.StringAttribute{
										Description:         "Image to use for the proxy",
										MarkdownDescription: "Image to use for the proxy",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"replicas": schema.Int64Attribute{
										Description:         "Replicas for the proxy",
										MarkdownDescription: "Replicas for the proxy",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "Resources requirements and limits for the proxy container",
										MarkdownDescription: "Resources requirements and limits for the proxy container",
										Attributes: map[string]schema.Attribute{
											"claims": schema.ListNestedAttribute{
												Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
												MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
															MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"request": schema.StringAttribute{
															Description:         "Request is the name chosen for a request in the referenced claim. If empty, everything from the claim is made available, otherwise only the result of this request.",
															MarkdownDescription: "Request is the name chosen for a request in the referenced claim. If empty, everything from the claim is made available, otherwise only the result of this request.",
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

											"limits": schema.MapAttribute{
												Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"requests": schema.MapAttribute{
												Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

							"type": schema.StringAttribute{
								Description:         "External access type LoadBalancer or Proxy, the LoadBalancer type will create a LoadBalancer service for each Valkey Shard (master node) The Proxy type will create a single LoadBalancer service and use an envoy proxy to route traffic to the Valkey Shards",
								MarkdownDescription: "External access type LoadBalancer or Proxy, the LoadBalancer type will create a LoadBalancer service for each Valkey Shard (master node) The Proxy type will create a single LoadBalancer service and use an envoy proxy to route traffic to the Valkey Shards",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("LoadBalancer", "Proxy"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"image": schema.StringAttribute{
						Description:         "Image to use",
						MarkdownDescription: "Image to use",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"node_selector": schema.MapAttribute{
						Description:         "Node Selector",
						MarkdownDescription: "Node Selector",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"nodes": schema.Int64Attribute{
						Description:         "Number of shards",
						MarkdownDescription: "Number of shards",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"prometheus": schema.BoolAttribute{
						Description:         "Enable prometheus",
						MarkdownDescription: "Enable prometheus",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"prometheus_labels": schema.MapAttribute{
						Description:         "Extra prometheus labels for operator targeting",
						MarkdownDescription: "Extra prometheus labels for operator targeting",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"replicas": schema.Int64Attribute{
						Description:         "Number of replicas",
						MarkdownDescription: "Number of replicas",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"resources": schema.SingleNestedAttribute{
						Description:         "Resources requirements and limits for the Valkey Server container",
						MarkdownDescription: "Resources requirements and limits for the Valkey Server container",
						Attributes: map[string]schema.Attribute{
							"claims": schema.ListNestedAttribute{
								Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
								MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
											MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"request": schema.StringAttribute{
											Description:         "Request is the name chosen for a request in the referenced claim. If empty, everything from the claim is made available, otherwise only the result of this request.",
											MarkdownDescription: "Request is the name chosen for a request in the referenced claim. If empty, everything from the claim is made available, otherwise only the result of this request.",
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

							"limits": schema.MapAttribute{
								Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"requests": schema.MapAttribute{
								Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

					"service_password": schema.SingleNestedAttribute{
						Description:         "Service Password",
						MarkdownDescription: "Service Password",
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

					"storage": schema.SingleNestedAttribute{
						Description:         "Persistent volume claim",
						MarkdownDescription: "Persistent volume claim",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
								MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"metadata": schema.MapAttribute{
								Description:         "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
								MarkdownDescription: "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"spec": schema.SingleNestedAttribute{
								Description:         "spec defines the desired characteristics of a volume requested by a pod author. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
								MarkdownDescription: "spec defines the desired characteristics of a volume requested by a pod author. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
								Attributes: map[string]schema.Attribute{
									"access_modes": schema.ListAttribute{
										Description:         "accessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
										MarkdownDescription: "accessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"data_source": schema.SingleNestedAttribute{
										Description:         "dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef, and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified. If the namespace is specified, then dataSourceRef will not be copied to dataSource.",
										MarkdownDescription: "dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef, and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified. If the namespace is specified, then dataSourceRef will not be copied to dataSource.",
										Attributes: map[string]schema.Attribute{
											"api_group": schema.StringAttribute{
												Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
												MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
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
										Description:         "dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the dataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, when namespace isn't specified in dataSourceRef, both fields (dataSource and dataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. When namespace is specified in dataSourceRef, dataSource isn't set to the same value and must be empty. There are three important differences between dataSource and dataSourceRef: * While dataSource only allows two specific types of objects, dataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects. * While dataSource ignores disallowed values (dropping them), dataSourceRef preserves all values, and generates an error if a disallowed value is specified. * While dataSource only allows local objects, dataSourceRef allows objects in any namespaces. (Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled. (Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
										MarkdownDescription: "dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the dataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, when namespace isn't specified in dataSourceRef, both fields (dataSource and dataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. When namespace is specified in dataSourceRef, dataSource isn't set to the same value and must be empty. There are three important differences between dataSource and dataSourceRef: * While dataSource only allows two specific types of objects, dataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects. * While dataSource ignores disallowed values (dropping them), dataSourceRef preserves all values, and generates an error if a disallowed value is specified. * While dataSource only allows local objects, dataSourceRef allows objects in any namespaces. (Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled. (Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
										Attributes: map[string]schema.Attribute{
											"api_group": schema.StringAttribute{
												Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
												MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
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
												Description:         "Namespace is the namespace of resource being referenced Note that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details. (Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
												MarkdownDescription: "Namespace is the namespace of resource being referenced Note that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details. (Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
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
										Description:         "resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
										MarkdownDescription: "resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
										Attributes: map[string]schema.Attribute{
											"limits": schema.MapAttribute{
												Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"requests": schema.MapAttribute{
												Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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
															Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
															MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"values": schema.ListAttribute{
															Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
															MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
												Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
										Description:         "storageClassName is the name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
										MarkdownDescription: "storageClassName is the name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"volume_attributes_class_name": schema.StringAttribute{
										Description:         "volumeAttributesClassName may be used to set the VolumeAttributesClass used by this claim. If specified, the CSI driver will create or update the volume with the attributes defined in the corresponding VolumeAttributesClass. This has a different purpose than storageClassName, it can be changed after the claim is created. An empty string value means that no VolumeAttributesClass will be applied to the claim but it's not allowed to reset this field to empty string once it is set. If unspecified and the PersistentVolumeClaim is unbound, the default VolumeAttributesClass will be set by the persistentvolume controller if it exists. If the resource referred to by volumeAttributesClass does not exist, this PersistentVolumeClaim will be set to a Pending state, as reflected by the modifyVolumeStatus field, until such as a resource exists. More info: https://kubernetes.io/docs/concepts/storage/volume-attributes-classes/ (Beta) Using this field requires the VolumeAttributesClass feature gate to be enabled (off by default).",
										MarkdownDescription: "volumeAttributesClassName may be used to set the VolumeAttributesClass used by this claim. If specified, the CSI driver will create or update the volume with the attributes defined in the corresponding VolumeAttributesClass. This has a different purpose than storageClassName, it can be changed after the claim is created. An empty string value means that no VolumeAttributesClass will be applied to the claim but it's not allowed to reset this field to empty string once it is set. If unspecified and the PersistentVolumeClaim is unbound, the default VolumeAttributesClass will be set by the persistentvolume controller if it exists. If the resource referred to by volumeAttributesClass does not exist, this PersistentVolumeClaim will be set to a Pending state, as reflected by the modifyVolumeStatus field, until such as a resource exists. More info: https://kubernetes.io/docs/concepts/storage/volume-attributes-classes/ (Beta) Using this field requires the VolumeAttributesClass feature gate to be enabled (off by default).",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"volume_mode": schema.StringAttribute{
										Description:         "volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.",
										MarkdownDescription: "volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.",
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

							"status": schema.SingleNestedAttribute{
								Description:         "status represents the current information/status of a persistent volume claim. Read-only. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
								MarkdownDescription: "status represents the current information/status of a persistent volume claim. Read-only. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
								Attributes: map[string]schema.Attribute{
									"access_modes": schema.ListAttribute{
										Description:         "accessModes contains the actual access modes the volume backing the PVC has. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
										MarkdownDescription: "accessModes contains the actual access modes the volume backing the PVC has. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"allocated_resource_statuses": schema.MapAttribute{
										Description:         "allocatedResourceStatuses stores status of resource being resized for the given PVC. Key names follow standard Kubernetes label syntax. Valid values are either: * Un-prefixed keys: - storage - the capacity of the volume. * Custom resources must use implementation-defined prefixed names such as 'example.com/my-custom-resource' Apart from above values - keys that are unprefixed or have kubernetes.io prefix are considered reserved and hence may not be used. ClaimResourceStatus can be in any of following states: - ControllerResizeInProgress: State set when resize controller starts resizing the volume in control-plane. - ControllerResizeFailed: State set when resize has failed in resize controller with a terminal error. - NodeResizePending: State set when resize controller has finished resizing the volume but further resizing of volume is needed on the node. - NodeResizeInProgress: State set when kubelet starts resizing the volume. - NodeResizeFailed: State set when resizing has failed in kubelet with a terminal error. Transient errors don't set NodeResizeFailed. For example: if expanding a PVC for more capacity - this field can be one of the following states: - pvc.status.allocatedResourceStatus['storage'] = 'ControllerResizeInProgress' - pvc.status.allocatedResourceStatus['storage'] = 'ControllerResizeFailed' - pvc.status.allocatedResourceStatus['storage'] = 'NodeResizePending' - pvc.status.allocatedResourceStatus['storage'] = 'NodeResizeInProgress' - pvc.status.allocatedResourceStatus['storage'] = 'NodeResizeFailed' When this field is not set, it means that no resize operation is in progress for the given PVC. A controller that receives PVC update with previously unknown resourceName or ClaimResourceStatus should ignore the update for the purpose it was designed. For example - a controller that only is responsible for resizing capacity of the volume, should ignore PVC updates that change other valid resources associated with PVC. This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.",
										MarkdownDescription: "allocatedResourceStatuses stores status of resource being resized for the given PVC. Key names follow standard Kubernetes label syntax. Valid values are either: * Un-prefixed keys: - storage - the capacity of the volume. * Custom resources must use implementation-defined prefixed names such as 'example.com/my-custom-resource' Apart from above values - keys that are unprefixed or have kubernetes.io prefix are considered reserved and hence may not be used. ClaimResourceStatus can be in any of following states: - ControllerResizeInProgress: State set when resize controller starts resizing the volume in control-plane. - ControllerResizeFailed: State set when resize has failed in resize controller with a terminal error. - NodeResizePending: State set when resize controller has finished resizing the volume but further resizing of volume is needed on the node. - NodeResizeInProgress: State set when kubelet starts resizing the volume. - NodeResizeFailed: State set when resizing has failed in kubelet with a terminal error. Transient errors don't set NodeResizeFailed. For example: if expanding a PVC for more capacity - this field can be one of the following states: - pvc.status.allocatedResourceStatus['storage'] = 'ControllerResizeInProgress' - pvc.status.allocatedResourceStatus['storage'] = 'ControllerResizeFailed' - pvc.status.allocatedResourceStatus['storage'] = 'NodeResizePending' - pvc.status.allocatedResourceStatus['storage'] = 'NodeResizeInProgress' - pvc.status.allocatedResourceStatus['storage'] = 'NodeResizeFailed' When this field is not set, it means that no resize operation is in progress for the given PVC. A controller that receives PVC update with previously unknown resourceName or ClaimResourceStatus should ignore the update for the purpose it was designed. For example - a controller that only is responsible for resizing capacity of the volume, should ignore PVC updates that change other valid resources associated with PVC. This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"allocated_resources": schema.MapAttribute{
										Description:         "allocatedResources tracks the resources allocated to a PVC including its capacity. Key names follow standard Kubernetes label syntax. Valid values are either: * Un-prefixed keys: - storage - the capacity of the volume. * Custom resources must use implementation-defined prefixed names such as 'example.com/my-custom-resource' Apart from above values - keys that are unprefixed or have kubernetes.io prefix are considered reserved and hence may not be used. Capacity reported here may be larger than the actual capacity when a volume expansion operation is requested. For storage quota, the larger value from allocatedResources and PVC.spec.resources is used. If allocatedResources is not set, PVC.spec.resources alone is used for quota calculation. If a volume expansion capacity request is lowered, allocatedResources is only lowered if there are no expansion operations in progress and if the actual volume capacity is equal or lower than the requested capacity. A controller that receives PVC update with previously unknown resourceName should ignore the update for the purpose it was designed. For example - a controller that only is responsible for resizing capacity of the volume, should ignore PVC updates that change other valid resources associated with PVC. This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.",
										MarkdownDescription: "allocatedResources tracks the resources allocated to a PVC including its capacity. Key names follow standard Kubernetes label syntax. Valid values are either: * Un-prefixed keys: - storage - the capacity of the volume. * Custom resources must use implementation-defined prefixed names such as 'example.com/my-custom-resource' Apart from above values - keys that are unprefixed or have kubernetes.io prefix are considered reserved and hence may not be used. Capacity reported here may be larger than the actual capacity when a volume expansion operation is requested. For storage quota, the larger value from allocatedResources and PVC.spec.resources is used. If allocatedResources is not set, PVC.spec.resources alone is used for quota calculation. If a volume expansion capacity request is lowered, allocatedResources is only lowered if there are no expansion operations in progress and if the actual volume capacity is equal or lower than the requested capacity. A controller that receives PVC update with previously unknown resourceName should ignore the update for the purpose it was designed. For example - a controller that only is responsible for resizing capacity of the volume, should ignore PVC updates that change other valid resources associated with PVC. This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"capacity": schema.MapAttribute{
										Description:         "capacity represents the actual resources of the underlying volume.",
										MarkdownDescription: "capacity represents the actual resources of the underlying volume.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"conditions": schema.ListNestedAttribute{
										Description:         "conditions is the current Condition of persistent volume claim. If underlying persistent volume is being resized then the Condition will be set to 'Resizing'.",
										MarkdownDescription: "conditions is the current Condition of persistent volume claim. If underlying persistent volume is being resized then the Condition will be set to 'Resizing'.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"last_probe_time": schema.StringAttribute{
													Description:         "lastProbeTime is the time we probed the condition.",
													MarkdownDescription: "lastProbeTime is the time we probed the condition.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														validators.DateTime64Validator(),
													},
												},

												"last_transition_time": schema.StringAttribute{
													Description:         "lastTransitionTime is the time the condition transitioned from one status to another.",
													MarkdownDescription: "lastTransitionTime is the time the condition transitioned from one status to another.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														validators.DateTime64Validator(),
													},
												},

												"message": schema.StringAttribute{
													Description:         "message is the human-readable message indicating details about last transition.",
													MarkdownDescription: "message is the human-readable message indicating details about last transition.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"reason": schema.StringAttribute{
													Description:         "reason is a unique, this should be a short, machine understandable string that gives the reason for condition's last transition. If it reports 'Resizing' that means the underlying persistent volume is being resized.",
													MarkdownDescription: "reason is a unique, this should be a short, machine understandable string that gives the reason for condition's last transition. If it reports 'Resizing' that means the underlying persistent volume is being resized.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"status": schema.StringAttribute{
													Description:         "Status is the status of the condition. Can be True, False, Unknown. More info: https://kubernetes.io/docs/reference/kubernetes-api/config-and-storage-resources/persistent-volume-claim-v1/#:~:text=state%20of%20pvc-,conditions.status,-(string)%2C%20required",
													MarkdownDescription: "Status is the status of the condition. Can be True, False, Unknown. More info: https://kubernetes.io/docs/reference/kubernetes-api/config-and-storage-resources/persistent-volume-claim-v1/#:~:text=state%20of%20pvc-,conditions.status,-(string)%2C%20required",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"type": schema.StringAttribute{
													Description:         "Type is the type of the condition. More info: https://kubernetes.io/docs/reference/kubernetes-api/config-and-storage-resources/persistent-volume-claim-v1/#:~:text=set%20to%20%27ResizeStarted%27.-,PersistentVolumeClaimCondition,-contains%20details%20about",
													MarkdownDescription: "Type is the type of the condition. More info: https://kubernetes.io/docs/reference/kubernetes-api/config-and-storage-resources/persistent-volume-claim-v1/#:~:text=set%20to%20%27ResizeStarted%27.-,PersistentVolumeClaimCondition,-contains%20details%20about",
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

									"current_volume_attributes_class_name": schema.StringAttribute{
										Description:         "currentVolumeAttributesClassName is the current name of the VolumeAttributesClass the PVC is using. When unset, there is no VolumeAttributeClass applied to this PersistentVolumeClaim This is a beta field and requires enabling VolumeAttributesClass feature (off by default).",
										MarkdownDescription: "currentVolumeAttributesClassName is the current name of the VolumeAttributesClass the PVC is using. When unset, there is no VolumeAttributeClass applied to this PersistentVolumeClaim This is a beta field and requires enabling VolumeAttributesClass feature (off by default).",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"modify_volume_status": schema.SingleNestedAttribute{
										Description:         "ModifyVolumeStatus represents the status object of ControllerModifyVolume operation. When this is unset, there is no ModifyVolume operation being attempted. This is a beta field and requires enabling VolumeAttributesClass feature (off by default).",
										MarkdownDescription: "ModifyVolumeStatus represents the status object of ControllerModifyVolume operation. When this is unset, there is no ModifyVolume operation being attempted. This is a beta field and requires enabling VolumeAttributesClass feature (off by default).",
										Attributes: map[string]schema.Attribute{
											"status": schema.StringAttribute{
												Description:         "status is the status of the ControllerModifyVolume operation. It can be in any of following states: - Pending Pending indicates that the PersistentVolumeClaim cannot be modified due to unmet requirements, such as the specified VolumeAttributesClass not existing. - InProgress InProgress indicates that the volume is being modified. - Infeasible Infeasible indicates that the request has been rejected as invalid by the CSI driver. To resolve the error, a valid VolumeAttributesClass needs to be specified. Note: New statuses can be added in the future. Consumers should check for unknown statuses and fail appropriately.",
												MarkdownDescription: "status is the status of the ControllerModifyVolume operation. It can be in any of following states: - Pending Pending indicates that the PersistentVolumeClaim cannot be modified due to unmet requirements, such as the specified VolumeAttributesClass not existing. - InProgress InProgress indicates that the volume is being modified. - Infeasible Infeasible indicates that the request has been rejected as invalid by the CSI driver. To resolve the error, a valid VolumeAttributesClass needs to be specified. Note: New statuses can be added in the future. Consumers should check for unknown statuses and fail appropriately.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"target_volume_attributes_class_name": schema.StringAttribute{
												Description:         "targetVolumeAttributesClassName is the name of the VolumeAttributesClass the PVC currently being reconciled",
												MarkdownDescription: "targetVolumeAttributesClassName is the name of the VolumeAttributesClass the PVC currently being reconciled",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"phase": schema.StringAttribute{
										Description:         "phase represents the current phase of PersistentVolumeClaim.",
										MarkdownDescription: "phase represents the current phase of PersistentVolumeClaim.",
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

					"tls": schema.BoolAttribute{
						Description:         "TLS Support",
						MarkdownDescription: "TLS Support",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tolerations": schema.ListNestedAttribute{
						Description:         "Tolerations",
						MarkdownDescription: "Tolerations",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"effect": schema.StringAttribute{
									Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
									MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"key": schema.StringAttribute{
									Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
									MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"operator": schema.StringAttribute{
									Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
									MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"toleration_seconds": schema.Int64Attribute{
									Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
									MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
									MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
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

					"volume_permissions": schema.BoolAttribute{
						Description:         "Turn on an init container to set permissions on the persistent volume",
						MarkdownDescription: "Turn on an init container to set permissions on the persistent volume",
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
	}
}

func (r *HyperspikeIoValkeyV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_hyperspike_io_valkey_v1_manifest")

	var model HyperspikeIoValkeyV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("hyperspike.io/v1")
	model.Kind = pointer.String("Valkey")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
