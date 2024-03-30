/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package networking_k8s_aws_v1alpha1

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
	_ datasource.DataSource = &NetworkingK8SAwsPolicyEndpointV1Alpha1Manifest{}
)

func NewNetworkingK8SAwsPolicyEndpointV1Alpha1Manifest() datasource.DataSource {
	return &NetworkingK8SAwsPolicyEndpointV1Alpha1Manifest{}
}

type NetworkingK8SAwsPolicyEndpointV1Alpha1Manifest struct{}

type NetworkingK8SAwsPolicyEndpointV1Alpha1ManifestData struct {
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
		Egress *[]struct {
			Cidr   *string   `tfsdk:"cidr" json:"cidr,omitempty"`
			Except *[]string `tfsdk:"except" json:"except,omitempty"`
			Ports  *[]struct {
				EndPort  *int64  `tfsdk:"end_port" json:"endPort,omitempty"`
				Port     *int64  `tfsdk:"port" json:"port,omitempty"`
				Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
			} `tfsdk:"ports" json:"ports,omitempty"`
		} `tfsdk:"egress" json:"egress,omitempty"`
		Ingress *[]struct {
			Cidr   *string   `tfsdk:"cidr" json:"cidr,omitempty"`
			Except *[]string `tfsdk:"except" json:"except,omitempty"`
			Ports  *[]struct {
				EndPort  *int64  `tfsdk:"end_port" json:"endPort,omitempty"`
				Port     *int64  `tfsdk:"port" json:"port,omitempty"`
				Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
			} `tfsdk:"ports" json:"ports,omitempty"`
		} `tfsdk:"ingress" json:"ingress,omitempty"`
		PodIsolation *[]string `tfsdk:"pod_isolation" json:"podIsolation,omitempty"`
		PodSelector  *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"pod_selector" json:"podSelector,omitempty"`
		PodSelectorEndpoints *[]struct {
			HostIP    *string `tfsdk:"host_ip" json:"hostIP,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			PodIP     *string `tfsdk:"pod_ip" json:"podIP,omitempty"`
		} `tfsdk:"pod_selector_endpoints" json:"podSelectorEndpoints,omitempty"`
		PolicyRef *struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"policy_ref" json:"policyRef,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *NetworkingK8SAwsPolicyEndpointV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_networking_k8s_aws_policy_endpoint_v1alpha1_manifest"
}

func (r *NetworkingK8SAwsPolicyEndpointV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "PolicyEndpoint is the Schema for the policyendpoints API",
		MarkdownDescription: "PolicyEndpoint is the Schema for the policyendpoints API",
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
				Description:         "PolicyEndpointSpec defines the desired state of PolicyEndpoint",
				MarkdownDescription: "PolicyEndpointSpec defines the desired state of PolicyEndpoint",
				Attributes: map[string]schema.Attribute{
					"egress": schema.ListNestedAttribute{
						Description:         "Egress is the list of egress rules containing resolved network addresses",
						MarkdownDescription: "Egress is the list of egress rules containing resolved network addresses",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"cidr": schema.StringAttribute{
									Description:         "CIDR is the network address(s) of the endpoint",
									MarkdownDescription: "CIDR is the network address(s) of the endpoint",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"except": schema.ListAttribute{
									Description:         "Except is the exceptions to the CIDR ranges mentioned above.",
									MarkdownDescription: "Except is the exceptions to the CIDR ranges mentioned above.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ports": schema.ListNestedAttribute{
									Description:         "Ports is the list of ports",
									MarkdownDescription: "Ports is the list of ports",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"end_port": schema.Int64Attribute{
												Description:         "Endport specifies the port range port to endPort port must be defined and an integer, endPort > port",
												MarkdownDescription: "Endport specifies the port range port to endPort port must be defined and an integer, endPort > port",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"port": schema.Int64Attribute{
												Description:         "Port specifies the numerical port for the protocol. If empty applies to all ports",
												MarkdownDescription: "Port specifies the numerical port for the protocol. If empty applies to all ports",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"protocol": schema.StringAttribute{
												Description:         "Protocol specifies the transport protocol, default TCP",
												MarkdownDescription: "Protocol specifies the transport protocol, default TCP",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"ingress": schema.ListNestedAttribute{
						Description:         "Ingress is the list of ingress rules containing resolved network addresses",
						MarkdownDescription: "Ingress is the list of ingress rules containing resolved network addresses",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"cidr": schema.StringAttribute{
									Description:         "CIDR is the network address(s) of the endpoint",
									MarkdownDescription: "CIDR is the network address(s) of the endpoint",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"except": schema.ListAttribute{
									Description:         "Except is the exceptions to the CIDR ranges mentioned above.",
									MarkdownDescription: "Except is the exceptions to the CIDR ranges mentioned above.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ports": schema.ListNestedAttribute{
									Description:         "Ports is the list of ports",
									MarkdownDescription: "Ports is the list of ports",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"end_port": schema.Int64Attribute{
												Description:         "Endport specifies the port range port to endPort port must be defined and an integer, endPort > port",
												MarkdownDescription: "Endport specifies the port range port to endPort port must be defined and an integer, endPort > port",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"port": schema.Int64Attribute{
												Description:         "Port specifies the numerical port for the protocol. If empty applies to all ports",
												MarkdownDescription: "Port specifies the numerical port for the protocol. If empty applies to all ports",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"protocol": schema.StringAttribute{
												Description:         "Protocol specifies the transport protocol, default TCP",
												MarkdownDescription: "Protocol specifies the transport protocol, default TCP",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"pod_isolation": schema.ListAttribute{
						Description:         "PodIsolation specifies whether the pod needs to be isolated for a particular traffic direction Ingress or Egress, or both. If default isolation is not specified, and there are no ingress/egress rules, then the pod is not isolated from the point of view of this policy. This follows the NetworkPolicy spec.PolicyTypes.",
						MarkdownDescription: "PodIsolation specifies whether the pod needs to be isolated for a particular traffic direction Ingress or Egress, or both. If default isolation is not specified, and there are no ingress/egress rules, then the pod is not isolated from the point of view of this policy. This follows the NetworkPolicy spec.PolicyTypes.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"pod_selector": schema.SingleNestedAttribute{
						Description:         "PodSelector is the podSelector from the policy resource",
						MarkdownDescription: "PodSelector is the podSelector from the policy resource",
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

					"pod_selector_endpoints": schema.ListNestedAttribute{
						Description:         "PodSelectorEndpoints contains information about the pods matching the podSelector",
						MarkdownDescription: "PodSelectorEndpoints contains information about the pods matching the podSelector",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"host_ip": schema.StringAttribute{
									Description:         "HostIP is the IP address of the host the pod is currently running on",
									MarkdownDescription: "HostIP is the IP address of the host the pod is currently running on",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name is the pod name",
									MarkdownDescription: "Name is the pod name",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace is the pod namespace",
									MarkdownDescription: "Namespace is the pod namespace",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"pod_ip": schema.StringAttribute{
									Description:         "PodIP is the IP address of the pod",
									MarkdownDescription: "PodIP is the IP address of the pod",
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

					"policy_ref": schema.SingleNestedAttribute{
						Description:         "PolicyRef is a reference to the Kubernetes NetworkPolicy resource.",
						MarkdownDescription: "PolicyRef is a reference to the Kubernetes NetworkPolicy resource.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name is the name of the Policy",
								MarkdownDescription: "Name is the name of the Policy",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace is the namespace of the Policy",
								MarkdownDescription: "Namespace is the namespace of the Policy",
								Required:            true,
								Optional:            false,
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
	}
}

func (r *NetworkingK8SAwsPolicyEndpointV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_networking_k8s_aws_policy_endpoint_v1alpha1_manifest")

	var model NetworkingK8SAwsPolicyEndpointV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("networking.k8s.aws/v1alpha1")
	model.Kind = pointer.String("PolicyEndpoint")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
