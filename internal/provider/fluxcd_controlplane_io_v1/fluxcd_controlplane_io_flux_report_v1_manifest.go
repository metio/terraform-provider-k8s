/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package fluxcd_controlplane_io_v1

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
	_ datasource.DataSource = &FluxcdControlplaneIoFluxReportV1Manifest{}
)

func NewFluxcdControlplaneIoFluxReportV1Manifest() datasource.DataSource {
	return &FluxcdControlplaneIoFluxReportV1Manifest{}
}

type FluxcdControlplaneIoFluxReportV1Manifest struct{}

type FluxcdControlplaneIoFluxReportV1ManifestData struct {
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
		Components *[]struct {
			Image  *string `tfsdk:"image" json:"image,omitempty"`
			Name   *string `tfsdk:"name" json:"name,omitempty"`
			Ready  *bool   `tfsdk:"ready" json:"ready,omitempty"`
			Status *string `tfsdk:"status" json:"status,omitempty"`
		} `tfsdk:"components" json:"components,omitempty"`
		Distribution *struct {
			Entitlement *string `tfsdk:"entitlement" json:"entitlement,omitempty"`
			ManagedBy   *string `tfsdk:"managed_by" json:"managedBy,omitempty"`
			Status      *string `tfsdk:"status" json:"status,omitempty"`
			Version     *string `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"distribution" json:"distribution,omitempty"`
		Reconcilers *[]struct {
			ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
			Stats      *struct {
				Failing   *int64  `tfsdk:"failing" json:"failing,omitempty"`
				Running   *int64  `tfsdk:"running" json:"running,omitempty"`
				Suspended *int64  `tfsdk:"suspended" json:"suspended,omitempty"`
				TotalSize *string `tfsdk:"total_size" json:"totalSize,omitempty"`
			} `tfsdk:"stats" json:"stats,omitempty"`
		} `tfsdk:"reconcilers" json:"reconcilers,omitempty"`
		Sync *struct {
			Id     *string `tfsdk:"id" json:"id,omitempty"`
			Path   *string `tfsdk:"path" json:"path,omitempty"`
			Ready  *bool   `tfsdk:"ready" json:"ready,omitempty"`
			Source *string `tfsdk:"source" json:"source,omitempty"`
			Status *string `tfsdk:"status" json:"status,omitempty"`
		} `tfsdk:"sync" json:"sync,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *FluxcdControlplaneIoFluxReportV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_fluxcd_controlplane_io_flux_report_v1_manifest"
}

func (r *FluxcdControlplaneIoFluxReportV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "FluxReport is the Schema for the fluxreports API.",
		MarkdownDescription: "FluxReport is the Schema for the fluxreports API.",
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
				Description:         "FluxReportSpec defines the observed state of a Flux installation.",
				MarkdownDescription: "FluxReportSpec defines the observed state of a Flux installation.",
				Attributes: map[string]schema.Attribute{
					"components": schema.ListNestedAttribute{
						Description:         "ComponentsStatus is the status of the Flux controller deployments.",
						MarkdownDescription: "ComponentsStatus is the status of the Flux controller deployments.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"image": schema.StringAttribute{
									Description:         "Image is the container image of the Flux component.",
									MarkdownDescription: "Image is the container image of the Flux component.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name is the name of the Flux component.",
									MarkdownDescription: "Name is the name of the Flux component.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"ready": schema.BoolAttribute{
									Description:         "Ready is the readiness status of the Flux component.",
									MarkdownDescription: "Ready is the readiness status of the Flux component.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"status": schema.StringAttribute{
									Description:         "Status is a human-readable message indicating details about the Flux component observed state.",
									MarkdownDescription: "Status is a human-readable message indicating details about the Flux component observed state.",
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

					"distribution": schema.SingleNestedAttribute{
						Description:         "Distribution is the version information of the Flux installation.",
						MarkdownDescription: "Distribution is the version information of the Flux installation.",
						Attributes: map[string]schema.Attribute{
							"entitlement": schema.StringAttribute{
								Description:         "Entitlement is the entitlement verification status.",
								MarkdownDescription: "Entitlement is the entitlement verification status.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"managed_by": schema.StringAttribute{
								Description:         "ManagedBy is the name of the operator managing the Flux instance.",
								MarkdownDescription: "ManagedBy is the name of the operator managing the Flux instance.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"status": schema.StringAttribute{
								Description:         "Status is a human-readable message indicating details about the distribution observed state.",
								MarkdownDescription: "Status is a human-readable message indicating details about the distribution observed state.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"version": schema.StringAttribute{
								Description:         "Version is the version of the Flux instance.",
								MarkdownDescription: "Version is the version of the Flux instance.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"reconcilers": schema.ListNestedAttribute{
						Description:         "ReconcilersStatus is the list of Flux reconcilers and their statistics grouped by API kind.",
						MarkdownDescription: "ReconcilersStatus is the list of Flux reconcilers and their statistics grouped by API kind.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"api_version": schema.StringAttribute{
									Description:         "APIVersion is the API version of the Flux resource.",
									MarkdownDescription: "APIVersion is the API version of the Flux resource.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"kind": schema.StringAttribute{
									Description:         "Kind is the kind of the Flux resource.",
									MarkdownDescription: "Kind is the kind of the Flux resource.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"stats": schema.SingleNestedAttribute{
									Description:         "Stats is the reconcile statics of the Flux resource kind.",
									MarkdownDescription: "Stats is the reconcile statics of the Flux resource kind.",
									Attributes: map[string]schema.Attribute{
										"failing": schema.Int64Attribute{
											Description:         "Failing is the number of reconciled resources in the Failing state.",
											MarkdownDescription: "Failing is the number of reconciled resources in the Failing state.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"running": schema.Int64Attribute{
											Description:         "Running is the number of reconciled resources in the Running state.",
											MarkdownDescription: "Running is the number of reconciled resources in the Running state.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"suspended": schema.Int64Attribute{
											Description:         "Suspended is the number of reconciled resources in the Suspended state.",
											MarkdownDescription: "Suspended is the number of reconciled resources in the Suspended state.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"total_size": schema.StringAttribute{
											Description:         "TotalSize is the total size of the artifacts in storage.",
											MarkdownDescription: "TotalSize is the total size of the artifacts in storage.",
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

					"sync": schema.SingleNestedAttribute{
						Description:         "SyncStatus is the status of the cluster sync Source and Kustomization resources.",
						MarkdownDescription: "SyncStatus is the status of the cluster sync Source and Kustomization resources.",
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Description:         "ID is the identifier of the sync.",
								MarkdownDescription: "ID is the identifier of the sync.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"path": schema.StringAttribute{
								Description:         "Path is the kustomize path of the sync.",
								MarkdownDescription: "Path is the kustomize path of the sync.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ready": schema.BoolAttribute{
								Description:         "Ready is the readiness status of the sync.",
								MarkdownDescription: "Ready is the readiness status of the sync.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"source": schema.StringAttribute{
								Description:         "Source is the URL of the source repository.",
								MarkdownDescription: "Source is the URL of the source repository.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"status": schema.StringAttribute{
								Description:         "Status is a human-readable message indicating details about the sync observed state.",
								MarkdownDescription: "Status is a human-readable message indicating details about the sync observed state.",
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *FluxcdControlplaneIoFluxReportV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_fluxcd_controlplane_io_flux_report_v1_manifest")

	var model FluxcdControlplaneIoFluxReportV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("fluxcd.controlplane.io/v1")
	model.Kind = pointer.String("FluxReport")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
