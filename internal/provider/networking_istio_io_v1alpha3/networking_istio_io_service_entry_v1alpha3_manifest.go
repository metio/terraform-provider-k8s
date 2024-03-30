/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package networking_istio_io_v1alpha3

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
	_ datasource.DataSource = &NetworkingIstioIoServiceEntryV1Alpha3Manifest{}
)

func NewNetworkingIstioIoServiceEntryV1Alpha3Manifest() datasource.DataSource {
	return &NetworkingIstioIoServiceEntryV1Alpha3Manifest{}
}

type NetworkingIstioIoServiceEntryV1Alpha3Manifest struct{}

type NetworkingIstioIoServiceEntryV1Alpha3ManifestData struct {
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
		Addresses *[]string `tfsdk:"addresses" json:"addresses,omitempty"`
		Endpoints *[]struct {
			Address        *string            `tfsdk:"address" json:"address,omitempty"`
			Labels         *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Locality       *string            `tfsdk:"locality" json:"locality,omitempty"`
			Network        *string            `tfsdk:"network" json:"network,omitempty"`
			Ports          *map[string]string `tfsdk:"ports" json:"ports,omitempty"`
			ServiceAccount *string            `tfsdk:"service_account" json:"serviceAccount,omitempty"`
			Weight         *int64             `tfsdk:"weight" json:"weight,omitempty"`
		} `tfsdk:"endpoints" json:"endpoints,omitempty"`
		ExportTo *[]string `tfsdk:"export_to" json:"exportTo,omitempty"`
		Hosts    *[]string `tfsdk:"hosts" json:"hosts,omitempty"`
		Location *string   `tfsdk:"location" json:"location,omitempty"`
		Ports    *[]struct {
			Name       *string `tfsdk:"name" json:"name,omitempty"`
			Number     *int64  `tfsdk:"number" json:"number,omitempty"`
			Protocol   *string `tfsdk:"protocol" json:"protocol,omitempty"`
			TargetPort *int64  `tfsdk:"target_port" json:"targetPort,omitempty"`
		} `tfsdk:"ports" json:"ports,omitempty"`
		Resolution       *string   `tfsdk:"resolution" json:"resolution,omitempty"`
		SubjectAltNames  *[]string `tfsdk:"subject_alt_names" json:"subjectAltNames,omitempty"`
		WorkloadSelector *struct {
			Labels *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		} `tfsdk:"workload_selector" json:"workloadSelector,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *NetworkingIstioIoServiceEntryV1Alpha3Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_networking_istio_io_service_entry_v1alpha3_manifest"
}

func (r *NetworkingIstioIoServiceEntryV1Alpha3Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "Configuration affecting service registry. See more details at: https://istio.io/docs/reference/config/networking/service-entry.html",
				MarkdownDescription: "Configuration affecting service registry. See more details at: https://istio.io/docs/reference/config/networking/service-entry.html",
				Attributes: map[string]schema.Attribute{
					"addresses": schema.ListAttribute{
						Description:         "The virtual IP addresses associated with the service.",
						MarkdownDescription: "The virtual IP addresses associated with the service.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"endpoints": schema.ListNestedAttribute{
						Description:         "One or more endpoints associated with the service.",
						MarkdownDescription: "One or more endpoints associated with the service.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"address": schema.StringAttribute{
									Description:         "Address associated with the network endpoint without the port.",
									MarkdownDescription: "Address associated with the network endpoint without the port.",
									Required:            false,
									Optional:            true,
									Computed:            false,
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
								},

								"network": schema.StringAttribute{
									Description:         "Network enables Istio to group endpoints resident in the same L3 domain/network.",
									MarkdownDescription: "Network enables Istio to group endpoints resident in the same L3 domain/network.",
									Required:            false,
									Optional:            true,
									Computed:            false,
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
								},

								"weight": schema.Int64Attribute{
									Description:         "The load balancing weight associated with the endpoint.",
									MarkdownDescription: "The load balancing weight associated with the endpoint.",
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

					"export_to": schema.ListAttribute{
						Description:         "A list of namespaces to which this service is exported.",
						MarkdownDescription: "A list of namespaces to which this service is exported.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"hosts": schema.ListAttribute{
						Description:         "The hosts associated with the ServiceEntry.",
						MarkdownDescription: "The hosts associated with the ServiceEntry.",
						ElementType:         types.StringType,
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"location": schema.StringAttribute{
						Description:         "Specify whether the service should be considered external to the mesh or part of the mesh.Valid Options: MESH_EXTERNAL, MESH_INTERNAL",
						MarkdownDescription: "Specify whether the service should be considered external to the mesh or part of the mesh.Valid Options: MESH_EXTERNAL, MESH_INTERNAL",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("MESH_EXTERNAL", "MESH_INTERNAL"),
						},
					},

					"ports": schema.ListNestedAttribute{
						Description:         "The ports associated with the external service.",
						MarkdownDescription: "The ports associated with the external service.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Label assigned to the port.",
									MarkdownDescription: "Label assigned to the port.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"number": schema.Int64Attribute{
									Description:         "A valid non-negative integer port number.",
									MarkdownDescription: "A valid non-negative integer port number.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"protocol": schema.StringAttribute{
									Description:         "The protocol exposed on the port.",
									MarkdownDescription: "The protocol exposed on the port.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"target_port": schema.Int64Attribute{
									Description:         "The port number on the endpoint where the traffic will be received.",
									MarkdownDescription: "The port number on the endpoint where the traffic will be received.",
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

					"resolution": schema.StringAttribute{
						Description:         "Service resolution mode for the hosts.Valid Options: NONE, STATIC, DNS, DNS_ROUND_ROBIN",
						MarkdownDescription: "Service resolution mode for the hosts.Valid Options: NONE, STATIC, DNS, DNS_ROUND_ROBIN",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("NONE", "STATIC", "DNS", "DNS_ROUND_ROBIN"),
						},
					},

					"subject_alt_names": schema.ListAttribute{
						Description:         "If specified, the proxy will verify that the server certificate's subject alternate name matches one of the specified values.",
						MarkdownDescription: "If specified, the proxy will verify that the server certificate's subject alternate name matches one of the specified values.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"workload_selector": schema.SingleNestedAttribute{
						Description:         "Applicable only for MESH_INTERNAL services.",
						MarkdownDescription: "Applicable only for MESH_INTERNAL services.",
						Attributes: map[string]schema.Attribute{
							"labels": schema.MapAttribute{
								Description:         "One or more labels that indicate a specific set of pods/VMs on which the configuration should be applied.",
								MarkdownDescription: "One or more labels that indicate a specific set of pods/VMs on which the configuration should be applied.",
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
		},
	}
}

func (r *NetworkingIstioIoServiceEntryV1Alpha3Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_networking_istio_io_service_entry_v1alpha3_manifest")

	var model NetworkingIstioIoServiceEntryV1Alpha3ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("networking.istio.io/v1alpha3")
	model.Kind = pointer.String("ServiceEntry")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
