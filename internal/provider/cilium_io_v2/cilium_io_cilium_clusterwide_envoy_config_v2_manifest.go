/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package cilium_io_v2

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
	_ datasource.DataSource = &CiliumIoCiliumClusterwideEnvoyConfigV2Manifest{}
)

func NewCiliumIoCiliumClusterwideEnvoyConfigV2Manifest() datasource.DataSource {
	return &CiliumIoCiliumClusterwideEnvoyConfigV2Manifest{}
}

type CiliumIoCiliumClusterwideEnvoyConfigV2Manifest struct{}

type CiliumIoCiliumClusterwideEnvoyConfigV2ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		BackendServices *[]struct {
			Name      *string   `tfsdk:"name" json:"name,omitempty"`
			Namespace *string   `tfsdk:"namespace" json:"namespace,omitempty"`
			Number    *[]string `tfsdk:"number" json:"number,omitempty"`
		} `tfsdk:"backend_services" json:"backendServices,omitempty"`
		NodeSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
		Resources *[]map[string]string `tfsdk:"resources" json:"resources,omitempty"`
		Services  *[]struct {
			Listener  *string   `tfsdk:"listener" json:"listener,omitempty"`
			Name      *string   `tfsdk:"name" json:"name,omitempty"`
			Namespace *string   `tfsdk:"namespace" json:"namespace,omitempty"`
			Ports     *[]string `tfsdk:"ports" json:"ports,omitempty"`
		} `tfsdk:"services" json:"services,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CiliumIoCiliumClusterwideEnvoyConfigV2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cilium_io_cilium_clusterwide_envoy_config_v2_manifest"
}

func (r *CiliumIoCiliumClusterwideEnvoyConfigV2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
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
					"backend_services": schema.ListNestedAttribute{
						Description:         "BackendServices specifies Kubernetes services whose backends are automatically synced to Envoy using EDS. Traffic for these services is not forwarded to an Envoy listener. This allows an Envoy listener load balance traffic to these backends while normal Cilium service load balancing takes care of balancing traffic for these services at the same time.",
						MarkdownDescription: "BackendServices specifies Kubernetes services whose backends are automatically synced to Envoy using EDS. Traffic for these services is not forwarded to an Envoy listener. This allows an Envoy listener load balance traffic to these backends while normal Cilium service load balancing takes care of balancing traffic for these services at the same time.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name is the name of a destination Kubernetes service that identifies traffic to be redirected.",
									MarkdownDescription: "Name is the name of a destination Kubernetes service that identifies traffic to be redirected.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace is the Kubernetes service namespace. In CiliumEnvoyConfig namespace defaults to the namespace of the CEC, In CiliumClusterwideEnvoyConfig namespace defaults to 'default'.",
									MarkdownDescription: "Namespace is the Kubernetes service namespace. In CiliumEnvoyConfig namespace defaults to the namespace of the CEC, In CiliumClusterwideEnvoyConfig namespace defaults to 'default'.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"number": schema.ListAttribute{
									Description:         "Ports is a set of port numbers, which can be used for filtering in case of underlying is exposing multiple port numbers.",
									MarkdownDescription: "Ports is a set of port numbers, which can be used for filtering in case of underlying is exposing multiple port numbers.",
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

					"node_selector": schema.SingleNestedAttribute{
						Description:         "NodeSelector is a label selector that determines to which nodes this configuration applies. If nil, then this config applies to all nodes.",
						MarkdownDescription: "NodeSelector is a label selector that determines to which nodes this configuration applies. If nil, then this config applies to all nodes.",
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
											Validators: []validator.String{
												stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
											},
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

					"resources": schema.ListAttribute{
						Description:         "Envoy xDS resources, a list of the following Envoy resource types: type.googleapis.com/envoy.config.listener.v3.Listener, type.googleapis.com/envoy.config.route.v3.RouteConfiguration, type.googleapis.com/envoy.config.cluster.v3.Cluster, type.googleapis.com/envoy.config.endpoint.v3.ClusterLoadAssignment, and type.googleapis.com/envoy.extensions.transport_sockets.tls.v3.Secret.",
						MarkdownDescription: "Envoy xDS resources, a list of the following Envoy resource types: type.googleapis.com/envoy.config.listener.v3.Listener, type.googleapis.com/envoy.config.route.v3.RouteConfiguration, type.googleapis.com/envoy.config.cluster.v3.Cluster, type.googleapis.com/envoy.config.endpoint.v3.ClusterLoadAssignment, and type.googleapis.com/envoy.extensions.transport_sockets.tls.v3.Secret.",
						ElementType:         types.MapType{ElemType: types.StringType},
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"services": schema.ListNestedAttribute{
						Description:         "Services specifies Kubernetes services for which traffic is forwarded to an Envoy listener for L7 load balancing. Backends of these services are automatically synced to Envoy usign EDS.",
						MarkdownDescription: "Services specifies Kubernetes services for which traffic is forwarded to an Envoy listener for L7 load balancing. Backends of these services are automatically synced to Envoy usign EDS.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"listener": schema.StringAttribute{
									Description:         "Listener specifies the name of the Envoy listener the service traffic is redirected to. The listener must be specified in the Envoy 'resources' of the same CiliumEnvoyConfig. If omitted, the first listener specified in 'resources' is used.",
									MarkdownDescription: "Listener specifies the name of the Envoy listener the service traffic is redirected to. The listener must be specified in the Envoy 'resources' of the same CiliumEnvoyConfig. If omitted, the first listener specified in 'resources' is used.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name is the name of a destination Kubernetes service that identifies traffic to be redirected.",
									MarkdownDescription: "Name is the name of a destination Kubernetes service that identifies traffic to be redirected.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace is the Kubernetes service namespace. In CiliumEnvoyConfig namespace this is overridden to the namespace of the CEC, In CiliumClusterwideEnvoyConfig namespace defaults to 'default'.",
									MarkdownDescription: "Namespace is the Kubernetes service namespace. In CiliumEnvoyConfig namespace this is overridden to the namespace of the CEC, In CiliumClusterwideEnvoyConfig namespace defaults to 'default'.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ports": schema.ListAttribute{
									Description:         "Ports is a set of service's frontend ports that should be redirected to the Envoy listener. By default all frontend ports of the service are redirected.",
									MarkdownDescription: "Ports is a set of service's frontend ports that should be redirected to the Envoy listener. By default all frontend ports of the service are redirected.",
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *CiliumIoCiliumClusterwideEnvoyConfigV2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cilium_io_cilium_clusterwide_envoy_config_v2_manifest")

	var model CiliumIoCiliumClusterwideEnvoyConfigV2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("cilium.io/v2")
	model.Kind = pointer.String("CiliumClusterwideEnvoyConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
