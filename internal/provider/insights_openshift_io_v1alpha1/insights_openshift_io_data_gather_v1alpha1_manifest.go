/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package insights_openshift_io_v1alpha1

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
	_ datasource.DataSource = &InsightsOpenshiftIoDataGatherV1Alpha1Manifest{}
)

func NewInsightsOpenshiftIoDataGatherV1Alpha1Manifest() datasource.DataSource {
	return &InsightsOpenshiftIoDataGatherV1Alpha1Manifest{}
}

type InsightsOpenshiftIoDataGatherV1Alpha1Manifest struct{}

type InsightsOpenshiftIoDataGatherV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		DataPolicy *string `tfsdk:"data_policy" json:"dataPolicy,omitempty"`
		Gatherers  *[]struct {
			Name  *string `tfsdk:"name" json:"name,omitempty"`
			State *string `tfsdk:"state" json:"state,omitempty"`
		} `tfsdk:"gatherers" json:"gatherers,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *InsightsOpenshiftIoDataGatherV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_insights_openshift_io_data_gather_v1alpha1_manifest"
}

func (r *InsightsOpenshiftIoDataGatherV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DataGather provides data gather configuration options and status for the particular Insights data gathering.  Compatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
		MarkdownDescription: "DataGather provides data gather configuration options and status for the particular Insights data gathering.  Compatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
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
				Description:         "spec holds user settable values for configuration",
				MarkdownDescription: "spec holds user settable values for configuration",
				Attributes: map[string]schema.Attribute{
					"data_policy": schema.StringAttribute{
						Description:         "dataPolicy allows user to enable additional global obfuscation of the IP addresses and base domain in the Insights archive data. Valid values are 'ClearText' and 'ObfuscateNetworking'. When set to ClearText the data is not obfuscated. When set to ObfuscateNetworking the IP addresses and the cluster domain name are obfuscated. When omitted, this means no opinion and the platform is left to choose a reasonable default, which is subject to change over time. The current default is ClearText.",
						MarkdownDescription: "dataPolicy allows user to enable additional global obfuscation of the IP addresses and base domain in the Insights archive data. Valid values are 'ClearText' and 'ObfuscateNetworking'. When set to ClearText the data is not obfuscated. When set to ObfuscateNetworking the IP addresses and the cluster domain name are obfuscated. When omitted, this means no opinion and the platform is left to choose a reasonable default, which is subject to change over time. The current default is ClearText.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("", "ClearText", "ObfuscateNetworking"),
						},
					},

					"gatherers": schema.ListNestedAttribute{
						Description:         "gatherers is a list of gatherers configurations. The particular gatherers IDs can be found at https://github.com/openshift/insights-operator/blob/master/docs/gathered-data.md. Run the following command to get the names of last active gatherers: 'oc get insightsoperators.operator.openshift.io cluster -o json | jq '.status.gatherStatus.gatherers[].name''",
						MarkdownDescription: "gatherers is a list of gatherers configurations. The particular gatherers IDs can be found at https://github.com/openshift/insights-operator/blob/master/docs/gathered-data.md. Run the following command to get the names of last active gatherers: 'oc get insightsoperators.operator.openshift.io cluster -o json | jq '.status.gatherStatus.gatherers[].name''",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "name is the name of specific gatherer",
									MarkdownDescription: "name is the name of specific gatherer",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"state": schema.StringAttribute{
									Description:         "state allows you to configure specific gatherer. Valid values are 'Enabled', 'Disabled' and omitted. When omitted, this means no opinion and the platform is left to choose a reasonable default. The current default is Enabled.",
									MarkdownDescription: "state allows you to configure specific gatherer. Valid values are 'Enabled', 'Disabled' and omitted. When omitted, this means no opinion and the platform is left to choose a reasonable default. The current default is Enabled.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("", "Enabled", "Disabled"),
									},
								},
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
		},
	}
}

func (r *InsightsOpenshiftIoDataGatherV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_insights_openshift_io_data_gather_v1alpha1_manifest")

	var model InsightsOpenshiftIoDataGatherV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("insights.openshift.io/v1alpha1")
	model.Kind = pointer.String("DataGather")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
