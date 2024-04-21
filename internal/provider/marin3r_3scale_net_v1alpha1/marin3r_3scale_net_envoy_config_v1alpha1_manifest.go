/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package marin3r_3scale_net_v1alpha1

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
	_ datasource.DataSource = &Marin3R3ScaleNetEnvoyConfigV1Alpha1Manifest{}
)

func NewMarin3R3ScaleNetEnvoyConfigV1Alpha1Manifest() datasource.DataSource {
	return &Marin3R3ScaleNetEnvoyConfigV1Alpha1Manifest{}
}

type Marin3R3ScaleNetEnvoyConfigV1Alpha1Manifest struct{}

type Marin3R3ScaleNetEnvoyConfigV1Alpha1ManifestData struct {
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
		EnvoyAPI       *string `tfsdk:"envoy_api" json:"envoyAPI,omitempty"`
		EnvoyResources *struct {
			Clusters *[]struct {
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"clusters" json:"clusters,omitempty"`
			Endpoints *[]struct {
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"endpoints" json:"endpoints,omitempty"`
			ExtensionConfigs *[]struct {
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"extension_configs" json:"extensionConfigs,omitempty"`
			Listeners *[]struct {
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"listeners" json:"listeners,omitempty"`
			Routes *[]struct {
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"routes" json:"routes,omitempty"`
			Runtimes *[]struct {
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"runtimes" json:"runtimes,omitempty"`
			ScopedRoutes *[]struct {
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"scoped_routes" json:"scopedRoutes,omitempty"`
			Secrets *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
				Ref  *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"ref" json:"ref,omitempty"`
			} `tfsdk:"secrets" json:"secrets,omitempty"`
		} `tfsdk:"envoy_resources" json:"envoyResources,omitempty"`
		NodeID    *string `tfsdk:"node_id" json:"nodeID,omitempty"`
		Resources *[]struct {
			Blueprint                  *string `tfsdk:"blueprint" json:"blueprint,omitempty"`
			GenerateFromEndpointSlices *struct {
				ClusterName *string `tfsdk:"cluster_name" json:"clusterName,omitempty"`
				Selector    *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
				} `tfsdk:"selector" json:"selector,omitempty"`
				TargetPort *string `tfsdk:"target_port" json:"targetPort,omitempty"`
			} `tfsdk:"generate_from_endpoint_slices" json:"generateFromEndpointSlices,omitempty"`
			GenerateFromOpaqueSecret *struct {
				Alias *string `tfsdk:"alias" json:"alias,omitempty"`
				Key   *string `tfsdk:"key" json:"key,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"generate_from_opaque_secret" json:"generateFromOpaqueSecret,omitempty"`
			GenerateFromTlsSecret *string            `tfsdk:"generate_from_tls_secret" json:"generateFromTlsSecret,omitempty"`
			Type                  *string            `tfsdk:"type" json:"type,omitempty"`
			Value                 *map[string]string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"resources" json:"resources,omitempty"`
		Serialization *string `tfsdk:"serialization" json:"serialization,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *Marin3R3ScaleNetEnvoyConfigV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_marin3r_3scale_net_envoy_config_v1alpha1_manifest"
}

func (r *Marin3R3ScaleNetEnvoyConfigV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "EnvoyConfig holds the configuration for a given envoy nodeID. The spec of an EnvoyConfig object holds the Envoy resources that conform the desired configuration for the given nodeID and that the discovery service will send to any envoy client that identifies itself with that nodeID.",
		MarkdownDescription: "EnvoyConfig holds the configuration for a given envoy nodeID. The spec of an EnvoyConfig object holds the Envoy resources that conform the desired configuration for the given nodeID and that the discovery service will send to any envoy client that identifies itself with that nodeID.",
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
				Description:         "EnvoyConfigSpec defines the desired state of EnvoyConfig",
				MarkdownDescription: "EnvoyConfigSpec defines the desired state of EnvoyConfig",
				Attributes: map[string]schema.Attribute{
					"envoy_api": schema.StringAttribute{
						Description:         "EnvoyAPI is the version of envoy's API to use. Defaults to v3.",
						MarkdownDescription: "EnvoyAPI is the version of envoy's API to use. Defaults to v3.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("v3"),
						},
					},

					"envoy_resources": schema.SingleNestedAttribute{
						Description:         "EnvoyResources holds the different types of resources suported by the envoy discovery service DEPRECATED. Use the 'resources' field instead.",
						MarkdownDescription: "EnvoyResources holds the different types of resources suported by the envoy discovery service DEPRECATED. Use the 'resources' field instead.",
						Attributes: map[string]schema.Attribute{
							"clusters": schema.ListNestedAttribute{
								Description:         "Clusters is a list of the envoy Cluster resource type. API V3 reference: https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/cluster/v3/cluster.proto",
								MarkdownDescription: "Clusters is a list of the envoy Cluster resource type. API V3 reference: https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/cluster/v3/cluster.proto",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of the envoy resource. DEPRECATED: this field has no effect and will be removed in an upcoming release. The name of the resources for discovery purposes is included in the resource itself. Refer to the envoy API reference to check how the name is specified for each resource type.",
											MarkdownDescription: "Name of the envoy resource. DEPRECATED: this field has no effect and will be removed in an upcoming release. The name of the resources for discovery purposes is included in the resource itself. Refer to the envoy API reference to check how the name is specified for each resource type.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Value is the serialized representation of the envoy resource",
											MarkdownDescription: "Value is the serialized representation of the envoy resource",
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

							"endpoints": schema.ListNestedAttribute{
								Description:         "Endpoints is a list of the envoy ClusterLoadAssignment resource type. API V3 reference: https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/endpoint/v3/endpoint.proto",
								MarkdownDescription: "Endpoints is a list of the envoy ClusterLoadAssignment resource type. API V3 reference: https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/endpoint/v3/endpoint.proto",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of the envoy resource. DEPRECATED: this field has no effect and will be removed in an upcoming release. The name of the resources for discovery purposes is included in the resource itself. Refer to the envoy API reference to check how the name is specified for each resource type.",
											MarkdownDescription: "Name of the envoy resource. DEPRECATED: this field has no effect and will be removed in an upcoming release. The name of the resources for discovery purposes is included in the resource itself. Refer to the envoy API reference to check how the name is specified for each resource type.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Value is the serialized representation of the envoy resource",
											MarkdownDescription: "Value is the serialized representation of the envoy resource",
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

							"extension_configs": schema.ListNestedAttribute{
								Description:         "ExtensionConfigs is a list of the envoy ExtensionConfig resource type API V3 reference: https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/core/v3/extension.proto",
								MarkdownDescription: "ExtensionConfigs is a list of the envoy ExtensionConfig resource type API V3 reference: https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/core/v3/extension.proto",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of the envoy resource. DEPRECATED: this field has no effect and will be removed in an upcoming release. The name of the resources for discovery purposes is included in the resource itself. Refer to the envoy API reference to check how the name is specified for each resource type.",
											MarkdownDescription: "Name of the envoy resource. DEPRECATED: this field has no effect and will be removed in an upcoming release. The name of the resources for discovery purposes is included in the resource itself. Refer to the envoy API reference to check how the name is specified for each resource type.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Value is the serialized representation of the envoy resource",
											MarkdownDescription: "Value is the serialized representation of the envoy resource",
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

							"listeners": schema.ListNestedAttribute{
								Description:         "Listeners is a list of the envoy Listener resource type. API V3 reference: https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/listener/v3/listener.proto",
								MarkdownDescription: "Listeners is a list of the envoy Listener resource type. API V3 reference: https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/listener/v3/listener.proto",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of the envoy resource. DEPRECATED: this field has no effect and will be removed in an upcoming release. The name of the resources for discovery purposes is included in the resource itself. Refer to the envoy API reference to check how the name is specified for each resource type.",
											MarkdownDescription: "Name of the envoy resource. DEPRECATED: this field has no effect and will be removed in an upcoming release. The name of the resources for discovery purposes is included in the resource itself. Refer to the envoy API reference to check how the name is specified for each resource type.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Value is the serialized representation of the envoy resource",
											MarkdownDescription: "Value is the serialized representation of the envoy resource",
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

							"routes": schema.ListNestedAttribute{
								Description:         "Routes is a list of the envoy Route resource type. API V3 reference: https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/route/v3/route.proto",
								MarkdownDescription: "Routes is a list of the envoy Route resource type. API V3 reference: https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/route/v3/route.proto",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of the envoy resource. DEPRECATED: this field has no effect and will be removed in an upcoming release. The name of the resources for discovery purposes is included in the resource itself. Refer to the envoy API reference to check how the name is specified for each resource type.",
											MarkdownDescription: "Name of the envoy resource. DEPRECATED: this field has no effect and will be removed in an upcoming release. The name of the resources for discovery purposes is included in the resource itself. Refer to the envoy API reference to check how the name is specified for each resource type.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Value is the serialized representation of the envoy resource",
											MarkdownDescription: "Value is the serialized representation of the envoy resource",
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

							"runtimes": schema.ListNestedAttribute{
								Description:         "Runtimes is a list of the envoy Runtime resource type. API V3 reference: https://www.envoyproxy.io/docs/envoy/latest/api-v3/service/runtime/v3/rtds.proto",
								MarkdownDescription: "Runtimes is a list of the envoy Runtime resource type. API V3 reference: https://www.envoyproxy.io/docs/envoy/latest/api-v3/service/runtime/v3/rtds.proto",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of the envoy resource. DEPRECATED: this field has no effect and will be removed in an upcoming release. The name of the resources for discovery purposes is included in the resource itself. Refer to the envoy API reference to check how the name is specified for each resource type.",
											MarkdownDescription: "Name of the envoy resource. DEPRECATED: this field has no effect and will be removed in an upcoming release. The name of the resources for discovery purposes is included in the resource itself. Refer to the envoy API reference to check how the name is specified for each resource type.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Value is the serialized representation of the envoy resource",
											MarkdownDescription: "Value is the serialized representation of the envoy resource",
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

							"scoped_routes": schema.ListNestedAttribute{
								Description:         "ScopedRoutes is a list of the envoy ScopeRoute resource type. API V3 reference: https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/route/v3/scoped_route.proto",
								MarkdownDescription: "ScopedRoutes is a list of the envoy ScopeRoute resource type. API V3 reference: https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/route/v3/scoped_route.proto",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of the envoy resource. DEPRECATED: this field has no effect and will be removed in an upcoming release. The name of the resources for discovery purposes is included in the resource itself. Refer to the envoy API reference to check how the name is specified for each resource type.",
											MarkdownDescription: "Name of the envoy resource. DEPRECATED: this field has no effect and will be removed in an upcoming release. The name of the resources for discovery purposes is included in the resource itself. Refer to the envoy API reference to check how the name is specified for each resource type.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Value is the serialized representation of the envoy resource",
											MarkdownDescription: "Value is the serialized representation of the envoy resource",
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

							"secrets": schema.ListNestedAttribute{
								Description:         "Secrets is a list of references to Kubernetes Secret objects.",
								MarkdownDescription: "Secrets is a list of references to Kubernetes Secret objects.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of the envoy tslCerticate secret resource. The certificate will be fetched from a Kubernetes Secrets of type 'kubernetes.io/tls' with this same name.",
											MarkdownDescription: "Name of the envoy tslCerticate secret resource. The certificate will be fetched from a Kubernetes Secrets of type 'kubernetes.io/tls' with this same name.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"ref": schema.SingleNestedAttribute{
											Description:         "DEPRECATED: this field is deprecated and it's value will be ignored. The 'name' of the Kubernetes Secret must match the 'name' field.",
											MarkdownDescription: "DEPRECATED: this field is deprecated and it's value will be ignored. The 'name' of the Kubernetes Secret must match the 'name' field.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is unique within a namespace to reference a secret resource.",
													MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace": schema.StringAttribute{
													Description:         "namespace defines the space within which the secret name must be unique.",
													MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
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

					"node_id": schema.StringAttribute{
						Description:         "NodeID holds the envoy identifier for the discovery service to know which set of resources to send to each of the envoy clients that connect to it.",
						MarkdownDescription: "NodeID holds the envoy identifier for the discovery service to know which set of resources to send to each of the envoy clients that connect to it.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"resources": schema.ListNestedAttribute{
						Description:         "Resources holds the different types of resources suported by the envoy discovery service",
						MarkdownDescription: "Resources holds the different types of resources suported by the envoy discovery service",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"blueprint": schema.StringAttribute{
									Description:         "Blueprint specifies a template to generate a configuration proto. It is currently only supported to generate secret configuration resources from k8s Secrets",
									MarkdownDescription: "Blueprint specifies a template to generate a configuration proto. It is currently only supported to generate secret configuration resources from k8s Secrets",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("tlsCertificate", "validationContext"),
									},
								},

								"generate_from_endpoint_slices": schema.SingleNestedAttribute{
									Description:         "Specifies a label selector to watch for EndpointSlices that will be used to generate the endpoint resource",
									MarkdownDescription: "Specifies a label selector to watch for EndpointSlices that will be used to generate the endpoint resource",
									Attributes: map[string]schema.Attribute{
										"cluster_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"selector": schema.SingleNestedAttribute{
											Description:         "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
											MarkdownDescription: "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
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
											Required: true,
											Optional: false,
											Computed: false,
										},

										"target_port": schema.StringAttribute{
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

								"generate_from_opaque_secret": schema.SingleNestedAttribute{
									Description:         "The name of a Kubernetes Secret of type 'Opaque'. It will generate an envoy 'generic secret' proto.",
									MarkdownDescription: "The name of a Kubernetes Secret of type 'Opaque'. It will generate an envoy 'generic secret' proto.",
									Attributes: map[string]schema.Attribute{
										"alias": schema.StringAttribute{
											Description:         "A unique name to refer to the name:key combination",
											MarkdownDescription: "A unique name to refer to the name:key combination",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"key": schema.StringAttribute{
											Description:         "The key of the secret to select from.  Must be a valid secret key.",
											MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "The name of the secret in the pod's namespace to select from.",
											MarkdownDescription: "The name of the secret in the pod's namespace to select from.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"generate_from_tls_secret": schema.StringAttribute{
									Description:         "The name of a Kubernetes Secret of type 'kubernetes.io/tls'",
									MarkdownDescription: "The name of a Kubernetes Secret of type 'kubernetes.io/tls'",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"type": schema.StringAttribute{
									Description:         "Type is the type url for the protobuf message",
									MarkdownDescription: "Type is the type url for the protobuf message",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("listener", "route", "scopedRoute", "cluster", "endpoint", "secret", "runtime", "extensionConfig"),
									},
								},

								"value": schema.MapAttribute{
									Description:         "Value is the protobufer message that configures the resource. The proto must match the envoy configuration API v3 specification for the given resource type (https://www.envoyproxy.io/docs/envoy/latest/api-docs/xds_protocol#resource-types)",
									MarkdownDescription: "Value is the protobufer message that configures the resource. The proto must match the envoy configuration API v3 specification for the given resource type (https://www.envoyproxy.io/docs/envoy/latest/api-docs/xds_protocol#resource-types)",
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

					"serialization": schema.StringAttribute{
						Description:         "Serialization specicifies the serialization format used to describe the resources. 'json' and 'yaml' are supported. 'json' is used if unset.",
						MarkdownDescription: "Serialization specicifies the serialization format used to describe the resources. 'json' and 'yaml' are supported. 'json' is used if unset.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("json", "yaml"),
						},
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *Marin3R3ScaleNetEnvoyConfigV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_marin3r_3scale_net_envoy_config_v1alpha1_manifest")

	var model Marin3R3ScaleNetEnvoyConfigV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("marin3r.3scale.net/v1alpha1")
	model.Kind = pointer.String("EnvoyConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
