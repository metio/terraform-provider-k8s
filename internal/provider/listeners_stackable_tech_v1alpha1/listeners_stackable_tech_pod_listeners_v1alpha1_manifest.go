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
	_ datasource.DataSource = &ListenersStackableTechPodListenersV1Alpha1Manifest{}
)

func NewListenersStackableTechPodListenersV1Alpha1Manifest() datasource.DataSource {
	return &ListenersStackableTechPodListenersV1Alpha1Manifest{}
}

type ListenersStackableTechPodListenersV1Alpha1Manifest struct{}

type ListenersStackableTechPodListenersV1Alpha1ManifestData struct {
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
		Listeners *struct {
			IngressAddresses *[]struct {
				Address     *string            `tfsdk:"address" json:"address,omitempty"`
				AddressType *string            `tfsdk:"address_type" json:"addressType,omitempty"`
				Ports       *map[string]string `tfsdk:"ports" json:"ports,omitempty"`
			} `tfsdk:"ingress_addresses" json:"ingressAddresses,omitempty"`
			Scope *string `tfsdk:"scope" json:"scope,omitempty"`
		} `tfsdk:"listeners" json:"listeners,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ListenersStackableTechPodListenersV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_listeners_stackable_tech_pod_listeners_v1alpha1_manifest"
}

func (r *ListenersStackableTechPodListenersV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Auto-generated derived type for PodListenersSpec via 'CustomResource'",
		MarkdownDescription: "Auto-generated derived type for PodListenersSpec via 'CustomResource'",
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
				Description:         "Informs users about Listeners that are bound by a given Pod. This is not expected to be created or modified by users. It will be created by the Stackable Listener Operator when mounting the listener volume, and is always named 'pod-{pod.metadata.uid}'.",
				MarkdownDescription: "Informs users about Listeners that are bound by a given Pod. This is not expected to be created or modified by users. It will be created by the Stackable Listener Operator when mounting the listener volume, and is always named 'pod-{pod.metadata.uid}'.",
				Attributes: map[string]schema.Attribute{
					"listeners": schema.SingleNestedAttribute{
						Description:         "All Listeners currently bound by the Pod. Indexed by Volume name (not PersistentVolume or PersistentVolumeClaim).",
						MarkdownDescription: "All Listeners currently bound by the Pod. Indexed by Volume name (not PersistentVolume or PersistentVolumeClaim).",
						Attributes: map[string]schema.Attribute{
							"ingress_addresses": schema.ListNestedAttribute{
								Description:         "Addresses allowing access to this Pod. Compared to 'ingress_addresses' on the Listener status, this list is restricted to addresses that can access this Pod. This field is intended to be equivalent to the files mounted into the Listener volume.",
								MarkdownDescription: "Addresses allowing access to this Pod. Compared to 'ingress_addresses' on the Listener status, this list is restricted to addresses that can access this Pod. This field is intended to be equivalent to the files mounted into the Listener volume.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"address": schema.StringAttribute{
											Description:         "The hostname or IP address to the Listener.",
											MarkdownDescription: "The hostname or IP address to the Listener.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"address_type": schema.StringAttribute{
											Description:         "The type of address ('Hostname' or 'IP').",
											MarkdownDescription: "The type of address ('Hostname' or 'IP').",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Hostname", "IP"),
											},
										},

										"ports": schema.MapAttribute{
											Description:         "Port mapping table.",
											MarkdownDescription: "Port mapping table.",
											ElementType:         types.StringType,
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

							"scope": schema.StringAttribute{
								Description:         "'Node' if this address only allows access to Pods hosted on a specific Kubernetes Node, otherwise 'Cluster'.",
								MarkdownDescription: "'Node' if this address only allows access to Pods hosted on a specific Kubernetes Node, otherwise 'Cluster'.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Node", "Cluster"),
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
	}
}

func (r *ListenersStackableTechPodListenersV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_listeners_stackable_tech_pod_listeners_v1alpha1_manifest")

	var model ListenersStackableTechPodListenersV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("listeners.stackable.tech/v1alpha1")
	model.Kind = pointer.String("PodListeners")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
