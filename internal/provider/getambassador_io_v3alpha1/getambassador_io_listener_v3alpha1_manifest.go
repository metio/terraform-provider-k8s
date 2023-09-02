/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package getambassador_io_v3alpha1

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &GetambassadorIoListenerV3Alpha1Manifest{}
)

func NewGetambassadorIoListenerV3Alpha1Manifest() datasource.DataSource {
	return &GetambassadorIoListenerV3Alpha1Manifest{}
}

type GetambassadorIoListenerV3Alpha1Manifest struct{}

type GetambassadorIoListenerV3Alpha1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
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
		Ambassador_id *[]string `tfsdk:"ambassador_id" json:"ambassador_id,omitempty"`
		HostBinding   *struct {
			Namespace *struct {
				From *string `tfsdk:"from" json:"from,omitempty"`
			} `tfsdk:"namespace" json:"namespace,omitempty"`
			Selector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"selector" json:"selector,omitempty"`
		} `tfsdk:"host_binding" json:"hostBinding,omitempty"`
		L7Depth       *int64    `tfsdk:"l7_depth" json:"l7Depth,omitempty"`
		Port          *int64    `tfsdk:"port" json:"port,omitempty"`
		Protocol      *string   `tfsdk:"protocol" json:"protocol,omitempty"`
		ProtocolStack *[]string `tfsdk:"protocol_stack" json:"protocolStack,omitempty"`
		SecurityModel *string   `tfsdk:"security_model" json:"securityModel,omitempty"`
		StatsPrefix   *string   `tfsdk:"stats_prefix" json:"statsPrefix,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GetambassadorIoListenerV3Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_getambassador_io_listener_v3alpha1_manifest"
}

func (r *GetambassadorIoListenerV3Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Listener is the Schema for the hosts API",
		MarkdownDescription: "Listener is the Schema for the hosts API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

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
				Description:         "ListenerSpec defines the desired state of this Port",
				MarkdownDescription: "ListenerSpec defines the desired state of this Port",
				Attributes: map[string]schema.Attribute{
					"ambassador_id": schema.ListAttribute{
						Description:         "AmbassadorID declares which Ambassador instances should pay attention to this resource. If no value is provided, the default is:  ambassador_id: - 'default'",
						MarkdownDescription: "AmbassadorID declares which Ambassador instances should pay attention to this resource. If no value is provided, the default is:  ambassador_id: - 'default'",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"host_binding": schema.SingleNestedAttribute{
						Description:         "HostBinding allows restricting which Hosts will be used for this Listener.",
						MarkdownDescription: "HostBinding allows restricting which Hosts will be used for this Listener.",
						Attributes: map[string]schema.Attribute{
							"namespace": schema.SingleNestedAttribute{
								Description:         "NamespaceBindingType defines we we specify which namespaces to look for Hosts in.",
								MarkdownDescription: "NamespaceBindingType defines we we specify which namespaces to look for Hosts in.",
								Attributes: map[string]schema.Attribute{
									"from": schema.StringAttribute{
										Description:         "NamespaceFromType defines how we evaluate a NamespaceBindingType.",
										MarkdownDescription: "NamespaceFromType defines how we evaluate a NamespaceBindingType.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("SELF", "ALL", "SELECTOR"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
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
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"l7_depth": schema.Int64Attribute{
						Description:         "L7Depth specifies how many layer 7 load balancers are between us and the edge of the network.",
						MarkdownDescription: "L7Depth specifies how many layer 7 load balancers are between us and the edge of the network.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"port": schema.Int64Attribute{
						Description:         "Port is the network port. Only one Listener can use a given port.",
						MarkdownDescription: "Port is the network port. Only one Listener can use a given port.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
							int64validator.AtMost(65535),
						},
					},

					"protocol": schema.StringAttribute{
						Description:         "Protocol is a shorthand for certain predefined stacks. Exactly one of Protocol or ProtocolStack must be supplied.",
						MarkdownDescription: "Protocol is a shorthand for certain predefined stacks. Exactly one of Protocol or ProtocolStack must be supplied.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("HTTP", "HTTPS", "HTTPPROXY", "HTTPSPROXY", "TCP", "TLS", "UDP"),
						},
					},

					"protocol_stack": schema.ListAttribute{
						Description:         "ProtocolStack explicitly specifies the protocol stack to set up. Exactly one of Protocol or ProtocolStack must be supplied.",
						MarkdownDescription: "ProtocolStack explicitly specifies the protocol stack to set up. Exactly one of Protocol or ProtocolStack must be supplied.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"security_model": schema.StringAttribute{
						Description:         "SecurityModel specifies how to determine whether connections to this port are secure or insecure.",
						MarkdownDescription: "SecurityModel specifies how to determine whether connections to this port are secure or insecure.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("XFP", "SECURE", "INSECURE"),
						},
					},

					"stats_prefix": schema.StringAttribute{
						Description:         "StatsPrefix specifies the prefix for statistics sent by Envoy about this Listener. The default depends on the protocol: 'ingress-http', 'ingress-https', 'ingress-tls-$port', or 'ingress-$port'.",
						MarkdownDescription: "StatsPrefix specifies the prefix for statistics sent by Envoy about this Listener. The default depends on the protocol: 'ingress-http', 'ingress-https', 'ingress-tls-$port', or 'ingress-$port'.",
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

func (r *GetambassadorIoListenerV3Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_getambassador_io_listener_v3alpha1_manifest")

	var model GetambassadorIoListenerV3Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("getambassador.io/v3alpha1")
	model.Kind = pointer.String("Listener")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"YAML Error: "+err.Error(),
		)
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
