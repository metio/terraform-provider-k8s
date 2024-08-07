/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package networking_istio_io_v1

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
	_ datasource.DataSource = &NetworkingIstioIoWorkloadEntryV1Manifest{}
)

func NewNetworkingIstioIoWorkloadEntryV1Manifest() datasource.DataSource {
	return &NetworkingIstioIoWorkloadEntryV1Manifest{}
}

type NetworkingIstioIoWorkloadEntryV1Manifest struct{}

type NetworkingIstioIoWorkloadEntryV1ManifestData struct {
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
		Address        *string            `tfsdk:"address" json:"address,omitempty"`
		Labels         *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Locality       *string            `tfsdk:"locality" json:"locality,omitempty"`
		Network        *string            `tfsdk:"network" json:"network,omitempty"`
		Ports          *map[string]string `tfsdk:"ports" json:"ports,omitempty"`
		ServiceAccount *string            `tfsdk:"service_account" json:"serviceAccount,omitempty"`
		Weight         *int64             `tfsdk:"weight" json:"weight,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *NetworkingIstioIoWorkloadEntryV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_networking_istio_io_workload_entry_v1_manifest"
}

func (r *NetworkingIstioIoWorkloadEntryV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "Configuration affecting VMs onboarded into the mesh. See more details at: https://istio.io/docs/reference/config/networking/workload-entry.html",
				MarkdownDescription: "Configuration affecting VMs onboarded into the mesh. See more details at: https://istio.io/docs/reference/config/networking/workload-entry.html",
				Attributes: map[string]schema.Attribute{
					"address": schema.StringAttribute{
						Description:         "Address associated with the network endpoint without the port.",
						MarkdownDescription: "Address associated with the network endpoint without the port.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(256),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "One or more labels associated with the endpoint.",
						MarkdownDescription: "One or more labels associated with the endpoint.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"locality": schema.StringAttribute{
						Description:         "The locality associated with the endpoint.",
						MarkdownDescription: "The locality associated with the endpoint.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(2048),
						},
					},

					"network": schema.StringAttribute{
						Description:         "Network enables Istio to group endpoints resident in the same L3 domain/network.",
						MarkdownDescription: "Network enables Istio to group endpoints resident in the same L3 domain/network.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(2048),
						},
					},

					"ports": schema.MapAttribute{
						Description:         "Set of ports associated with the endpoint.",
						MarkdownDescription: "Set of ports associated with the endpoint.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service_account": schema.StringAttribute{
						Description:         "The service account associated with the workload if a sidecar is present in the workload.",
						MarkdownDescription: "The service account associated with the workload if a sidecar is present in the workload.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(253),
						},
					},

					"weight": schema.Int64Attribute{
						Description:         "The load balancing weight associated with the endpoint.",
						MarkdownDescription: "The load balancing weight associated with the endpoint.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
							int64validator.AtMost(4.294967295e+09),
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

func (r *NetworkingIstioIoWorkloadEntryV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_networking_istio_io_workload_entry_v1_manifest")

	var model NetworkingIstioIoWorkloadEntryV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("networking.istio.io/v1")
	model.Kind = pointer.String("WorkloadEntry")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
