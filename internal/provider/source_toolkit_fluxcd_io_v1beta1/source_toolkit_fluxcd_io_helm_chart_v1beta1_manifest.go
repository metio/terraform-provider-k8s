/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package source_toolkit_fluxcd_io_v1beta1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &SourceToolkitFluxcdIoHelmChartV1Beta1Manifest{}
)

func NewSourceToolkitFluxcdIoHelmChartV1Beta1Manifest() datasource.DataSource {
	return &SourceToolkitFluxcdIoHelmChartV1Beta1Manifest{}
}

type SourceToolkitFluxcdIoHelmChartV1Beta1Manifest struct{}

type SourceToolkitFluxcdIoHelmChartV1Beta1ManifestData struct {
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
		AccessFrom *struct {
			NamespaceSelectors *[]struct {
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"namespace_selectors" json:"namespaceSelectors,omitempty"`
		} `tfsdk:"access_from" json:"accessFrom,omitempty"`
		Chart             *string `tfsdk:"chart" json:"chart,omitempty"`
		Interval          *string `tfsdk:"interval" json:"interval,omitempty"`
		ReconcileStrategy *string `tfsdk:"reconcile_strategy" json:"reconcileStrategy,omitempty"`
		SourceRef         *struct {
			ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"source_ref" json:"sourceRef,omitempty"`
		Suspend     *bool     `tfsdk:"suspend" json:"suspend,omitempty"`
		ValuesFile  *string   `tfsdk:"values_file" json:"valuesFile,omitempty"`
		ValuesFiles *[]string `tfsdk:"values_files" json:"valuesFiles,omitempty"`
		Version     *string   `tfsdk:"version" json:"version,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SourceToolkitFluxcdIoHelmChartV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_source_toolkit_fluxcd_io_helm_chart_v1beta1_manifest"
}

func (r *SourceToolkitFluxcdIoHelmChartV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "HelmChart is the Schema for the helmcharts API",
		MarkdownDescription: "HelmChart is the Schema for the helmcharts API",
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
				Description:         "HelmChartSpec defines the desired state of a Helm chart.",
				MarkdownDescription: "HelmChartSpec defines the desired state of a Helm chart.",
				Attributes: map[string]schema.Attribute{
					"access_from": schema.SingleNestedAttribute{
						Description:         "AccessFrom defines an Access Control List for allowing cross-namespace references to this object.",
						MarkdownDescription: "AccessFrom defines an Access Control List for allowing cross-namespace references to this object.",
						Attributes: map[string]schema.Attribute{
							"namespace_selectors": schema.ListNestedAttribute{
								Description:         "NamespaceSelectors is the list of namespace selectors to which this ACL applies. Items in this list are evaluated using a logical OR operation.",
								MarkdownDescription: "NamespaceSelectors is the list of namespace selectors to which this ACL applies. Items in this list are evaluated using a logical OR operation.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"match_labels": schema.MapAttribute{
											Description:         "MatchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
											MarkdownDescription: "MatchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
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

					"chart": schema.StringAttribute{
						Description:         "The name or path the Helm chart is available at in the SourceRef.",
						MarkdownDescription: "The name or path the Helm chart is available at in the SourceRef.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"interval": schema.StringAttribute{
						Description:         "The interval at which to check the Source for updates.",
						MarkdownDescription: "The interval at which to check the Source for updates.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"reconcile_strategy": schema.StringAttribute{
						Description:         "Determines what enables the creation of a new artifact. Valid values are ('ChartVersion', 'Revision'). See the documentation of the values for an explanation on their behavior. Defaults to ChartVersion when omitted.",
						MarkdownDescription: "Determines what enables the creation of a new artifact. Valid values are ('ChartVersion', 'Revision'). See the documentation of the values for an explanation on their behavior. Defaults to ChartVersion when omitted.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("ChartVersion", "Revision"),
						},
					},

					"source_ref": schema.SingleNestedAttribute{
						Description:         "The reference to the Source the chart is available at.",
						MarkdownDescription: "The reference to the Source the chart is available at.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "APIVersion of the referent.",
								MarkdownDescription: "APIVersion of the referent.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the referent, valid values are ('HelmRepository', 'GitRepository', 'Bucket').",
								MarkdownDescription: "Kind of the referent, valid values are ('HelmRepository', 'GitRepository', 'Bucket').",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("HelmRepository", "GitRepository", "Bucket"),
								},
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referent.",
								MarkdownDescription: "Name of the referent.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"suspend": schema.BoolAttribute{
						Description:         "This flag tells the controller to suspend the reconciliation of this source.",
						MarkdownDescription: "This flag tells the controller to suspend the reconciliation of this source.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"values_file": schema.StringAttribute{
						Description:         "Alternative values file to use as the default chart values, expected to be a relative path in the SourceRef. Deprecated in favor of ValuesFiles, for backwards compatibility the file defined here is merged before the ValuesFiles items. Ignored when omitted.",
						MarkdownDescription: "Alternative values file to use as the default chart values, expected to be a relative path in the SourceRef. Deprecated in favor of ValuesFiles, for backwards compatibility the file defined here is merged before the ValuesFiles items. Ignored when omitted.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"values_files": schema.ListAttribute{
						Description:         "Alternative list of values files to use as the chart values (values.yaml is not included by default), expected to be a relative path in the SourceRef. Values files are merged in the order of this list with the last file overriding the first. Ignored when omitted.",
						MarkdownDescription: "Alternative list of values files to use as the chart values (values.yaml is not included by default), expected to be a relative path in the SourceRef. Values files are merged in the order of this list with the last file overriding the first. Ignored when omitted.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"version": schema.StringAttribute{
						Description:         "The chart version semver expression, ignored for charts from GitRepository and Bucket sources. Defaults to latest when omitted.",
						MarkdownDescription: "The chart version semver expression, ignored for charts from GitRepository and Bucket sources. Defaults to latest when omitted.",
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

func (r *SourceToolkitFluxcdIoHelmChartV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_source_toolkit_fluxcd_io_helm_chart_v1beta1_manifest")

	var model SourceToolkitFluxcdIoHelmChartV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("source.toolkit.fluxcd.io/v1beta1")
	model.Kind = pointer.String("HelmChart")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
