/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package listeners_stackable_tech_v1alpha1

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
	_ datasource.DataSource = &ListenersStackableTechListenerV1Alpha1Manifest{}
)

func NewListenersStackableTechListenerV1Alpha1Manifest() datasource.DataSource {
	return &ListenersStackableTechListenerV1Alpha1Manifest{}
}

type ListenersStackableTechListenerV1Alpha1Manifest struct{}

type ListenersStackableTechListenerV1Alpha1ManifestData struct {
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
		ClassName              *string            `tfsdk:"class_name" json:"className,omitempty"`
		ExtraPodSelectorLabels *map[string]string `tfsdk:"extra_pod_selector_labels" json:"extraPodSelectorLabels,omitempty"`
		Ports                  *[]struct {
			Name     *string `tfsdk:"name" json:"name,omitempty"`
			Port     *int64  `tfsdk:"port" json:"port,omitempty"`
			Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
		} `tfsdk:"ports" json:"ports,omitempty"`
		PublishNotReadyAddresses *bool `tfsdk:"publish_not_ready_addresses" json:"publishNotReadyAddresses,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ListenersStackableTechListenerV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_listeners_stackable_tech_listener_v1alpha1_manifest"
}

func (r *ListenersStackableTechListenerV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Auto-generated derived type for ListenerSpec via 'CustomResource'",
		MarkdownDescription: "Auto-generated derived type for ListenerSpec via 'CustomResource'",
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
				Description:         "Exposes a set of pods to the outside world. Essentially a Stackable extension of a Kubernetes Service. Compared to a Service, a Listener changes three things: 1. It uses a cluster-level policy object (ListenerClass) to define how exactly the exposure works 2. It has a consistent API for reading back the exposed address(es) of the service 3. The Pod must mount a Volume referring to the Listener, which also allows ['sticky' scheduling](https://docs.stackable.tech/home/nightly/listener-operator/listener#_sticky_scheduling). Learn more in the [Listener documentation](https://docs.stackable.tech/home/nightly/listener-operator/listener).",
				MarkdownDescription: "Exposes a set of pods to the outside world. Essentially a Stackable extension of a Kubernetes Service. Compared to a Service, a Listener changes three things: 1. It uses a cluster-level policy object (ListenerClass) to define how exactly the exposure works 2. It has a consistent API for reading back the exposed address(es) of the service 3. The Pod must mount a Volume referring to the Listener, which also allows ['sticky' scheduling](https://docs.stackable.tech/home/nightly/listener-operator/listener#_sticky_scheduling). Learn more in the [Listener documentation](https://docs.stackable.tech/home/nightly/listener-operator/listener).",
				Attributes: map[string]schema.Attribute{
					"class_name": schema.StringAttribute{
						Description:         "The name of the [ListenerClass](https://docs.stackable.tech/home/nightly/listener-operator/listenerclass).",
						MarkdownDescription: "The name of the [ListenerClass](https://docs.stackable.tech/home/nightly/listener-operator/listenerclass).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"extra_pod_selector_labels": schema.MapAttribute{
						Description:         "Extra labels that the Pods must match in order to be exposed. They must _also_ still have a Volume referring to the Listener.",
						MarkdownDescription: "Extra labels that the Pods must match in order to be exposed. They must _also_ still have a Volume referring to the Listener.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ports": schema.ListNestedAttribute{
						Description:         "Ports that should be exposed.",
						MarkdownDescription: "Ports that should be exposed.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "The name of the port. The name of each port *must* be unique within a single Listener.",
									MarkdownDescription: "The name of the port. The name of each port *must* be unique within a single Listener.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"port": schema.Int64Attribute{
									Description:         "The port number.",
									MarkdownDescription: "The port number.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"protocol": schema.StringAttribute{
									Description:         "The layer-4 protocol ('TCP' or 'UDP').",
									MarkdownDescription: "The layer-4 protocol ('TCP' or 'UDP').",
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

					"publish_not_ready_addresses": schema.BoolAttribute{
						Description:         "Whether incoming traffic should also be directed to Pods that are not 'Ready'.",
						MarkdownDescription: "Whether incoming traffic should also be directed to Pods that are not 'Ready'.",
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

func (r *ListenersStackableTechListenerV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_listeners_stackable_tech_listener_v1alpha1_manifest")

	var model ListenersStackableTechListenerV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("listeners.stackable.tech/v1alpha1")
	model.Kind = pointer.String("Listener")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
