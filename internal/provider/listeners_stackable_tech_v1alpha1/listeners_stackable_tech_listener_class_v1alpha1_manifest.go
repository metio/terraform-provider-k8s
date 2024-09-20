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
	_ datasource.DataSource = &ListenersStackableTechListenerClassV1Alpha1Manifest{}
)

func NewListenersStackableTechListenerClassV1Alpha1Manifest() datasource.DataSource {
	return &ListenersStackableTechListenerClassV1Alpha1Manifest{}
}

type ListenersStackableTechListenerClassV1Alpha1Manifest struct{}

type ListenersStackableTechListenerClassV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ServiceAnnotations           *map[string]string `tfsdk:"service_annotations" json:"serviceAnnotations,omitempty"`
		ServiceExternalTrafficPolicy *string            `tfsdk:"service_external_traffic_policy" json:"serviceExternalTrafficPolicy,omitempty"`
		ServiceType                  *string            `tfsdk:"service_type" json:"serviceType,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ListenersStackableTechListenerClassV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_listeners_stackable_tech_listener_class_v1alpha1_manifest"
}

func (r *ListenersStackableTechListenerClassV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Auto-generated derived type for ListenerClassSpec via 'CustomResource'",
		MarkdownDescription: "Auto-generated derived type for ListenerClassSpec via 'CustomResource'",
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
				Description:         "Defines a policy for how [Listeners](https://docs.stackable.tech/home/nightly/listener-operator/listener) should be exposed. Read the [ListenerClass documentation](https://docs.stackable.tech/home/nightly/listener-operator/listenerclass) for more information.",
				MarkdownDescription: "Defines a policy for how [Listeners](https://docs.stackable.tech/home/nightly/listener-operator/listener) should be exposed. Read the [ListenerClass documentation](https://docs.stackable.tech/home/nightly/listener-operator/listenerclass) for more information.",
				Attributes: map[string]schema.Attribute{
					"service_annotations": schema.MapAttribute{
						Description:         "Annotations that should be added to the Service object.",
						MarkdownDescription: "Annotations that should be added to the Service object.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service_external_traffic_policy": schema.StringAttribute{
						Description:         "'externalTrafficPolicy' that should be set on the created ['Service'] objects. The default is 'Local' (in contrast to 'Cluster'), as we aim to direct traffic to a node running the workload and we should keep testing that as the primary configuration. Cluster is a fallback option for providers that break Local mode (IONOS so far).",
						MarkdownDescription: "'externalTrafficPolicy' that should be set on the created ['Service'] objects. The default is 'Local' (in contrast to 'Cluster'), as we aim to direct traffic to a node running the workload and we should keep testing that as the primary configuration. Cluster is a fallback option for providers that break Local mode (IONOS so far).",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Cluster", "Local"),
						},
					},

					"service_type": schema.StringAttribute{
						Description:         "The method used to access the services.",
						MarkdownDescription: "The method used to access the services.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("NodePort", "LoadBalancer", "ClusterIP"),
						},
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *ListenersStackableTechListenerClassV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_listeners_stackable_tech_listener_class_v1alpha1_manifest")

	var model ListenersStackableTechListenerClassV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("listeners.stackable.tech/v1alpha1")
	model.Kind = pointer.String("ListenerClass")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
