/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type SourceToolkitFluxcdIoHelmChartV1Beta1Resource struct{}

var (
	_ resource.Resource = (*SourceToolkitFluxcdIoHelmChartV1Beta1Resource)(nil)
)

type SourceToolkitFluxcdIoHelmChartV1Beta1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type SourceToolkitFluxcdIoHelmChartV1Beta1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		AccessFrom *struct {
			NamespaceSelectors *[]struct {
				MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
			} `tfsdk:"namespace_selectors" yaml:"namespaceSelectors,omitempty"`
		} `tfsdk:"access_from" yaml:"accessFrom,omitempty"`

		Interval *string `tfsdk:"interval" yaml:"interval,omitempty"`

		ReconcileStrategy *string `tfsdk:"reconcile_strategy" yaml:"reconcileStrategy,omitempty"`

		SourceRef *struct {
			ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"source_ref" yaml:"sourceRef,omitempty"`

		ValuesFile *string `tfsdk:"values_file" yaml:"valuesFile,omitempty"`

		ValuesFiles *[]string `tfsdk:"values_files" yaml:"valuesFiles,omitempty"`

		Chart *string `tfsdk:"chart" yaml:"chart,omitempty"`

		Suspend *bool `tfsdk:"suspend" yaml:"suspend,omitempty"`

		Version *string `tfsdk:"version" yaml:"version,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewSourceToolkitFluxcdIoHelmChartV1Beta1Resource() resource.Resource {
	return &SourceToolkitFluxcdIoHelmChartV1Beta1Resource{}
}

func (r *SourceToolkitFluxcdIoHelmChartV1Beta1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_toolkit_fluxcd_io_helm_chart_v1beta1"
}

func (r *SourceToolkitFluxcdIoHelmChartV1Beta1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "HelmChart is the Schema for the helmcharts API",
		MarkdownDescription: "HelmChart is the Schema for the helmcharts API",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.RequiresReplace(),
						},
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "HelmChartSpec defines the desired state of a Helm chart.",
				MarkdownDescription: "HelmChartSpec defines the desired state of a Helm chart.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"access_from": {
						Description:         "AccessFrom defines an Access Control List for allowing cross-namespace references to this object.",
						MarkdownDescription: "AccessFrom defines an Access Control List for allowing cross-namespace references to this object.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"namespace_selectors": {
								Description:         "NamespaceSelectors is the list of namespace selectors to which this ACL applies. Items in this list are evaluated using a logical OR operation.",
								MarkdownDescription: "NamespaceSelectors is the list of namespace selectors to which this ACL applies. Items in this list are evaluated using a logical OR operation.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"match_labels": {
										Description:         "MatchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "MatchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"interval": {
						Description:         "The interval at which to check the Source for updates.",
						MarkdownDescription: "The interval at which to check the Source for updates.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"reconcile_strategy": {
						Description:         "Determines what enables the creation of a new artifact. Valid values are ('ChartVersion', 'Revision'). See the documentation of the values for an explanation on their behavior. Defaults to ChartVersion when omitted.",
						MarkdownDescription: "Determines what enables the creation of a new artifact. Valid values are ('ChartVersion', 'Revision'). See the documentation of the values for an explanation on their behavior. Defaults to ChartVersion when omitted.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"source_ref": {
						Description:         "The reference to the Source the chart is available at.",
						MarkdownDescription: "The reference to the Source the chart is available at.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"api_version": {
								Description:         "APIVersion of the referent.",
								MarkdownDescription: "APIVersion of the referent.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"kind": {
								Description:         "Kind of the referent, valid values are ('HelmRepository', 'GitRepository', 'Bucket').",
								MarkdownDescription: "Kind of the referent, valid values are ('HelmRepository', 'GitRepository', 'Bucket').",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"name": {
								Description:         "Name of the referent.",
								MarkdownDescription: "Name of the referent.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"values_file": {
						Description:         "Alternative values file to use as the default chart values, expected to be a relative path in the SourceRef. Deprecated in favor of ValuesFiles, for backwards compatibility the file defined here is merged before the ValuesFiles items. Ignored when omitted.",
						MarkdownDescription: "Alternative values file to use as the default chart values, expected to be a relative path in the SourceRef. Deprecated in favor of ValuesFiles, for backwards compatibility the file defined here is merged before the ValuesFiles items. Ignored when omitted.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"values_files": {
						Description:         "Alternative list of values files to use as the chart values (values.yaml is not included by default), expected to be a relative path in the SourceRef. Values files are merged in the order of this list with the last file overriding the first. Ignored when omitted.",
						MarkdownDescription: "Alternative list of values files to use as the chart values (values.yaml is not included by default), expected to be a relative path in the SourceRef. Values files are merged in the order of this list with the last file overriding the first. Ignored when omitted.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"chart": {
						Description:         "The name or path the Helm chart is available at in the SourceRef.",
						MarkdownDescription: "The name or path the Helm chart is available at in the SourceRef.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"suspend": {
						Description:         "This flag tells the controller to suspend the reconciliation of this source.",
						MarkdownDescription: "This flag tells the controller to suspend the reconciliation of this source.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"version": {
						Description:         "The chart version semver expression, ignored for charts from GitRepository and Bucket sources. Defaults to latest when omitted.",
						MarkdownDescription: "The chart version semver expression, ignored for charts from GitRepository and Bucket sources. Defaults to latest when omitted.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *SourceToolkitFluxcdIoHelmChartV1Beta1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_source_toolkit_fluxcd_io_helm_chart_v1beta1")

	var state SourceToolkitFluxcdIoHelmChartV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SourceToolkitFluxcdIoHelmChartV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("source.toolkit.fluxcd.io/v1beta1")
	goModel.Kind = utilities.Ptr("HelmChart")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *SourceToolkitFluxcdIoHelmChartV1Beta1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_source_toolkit_fluxcd_io_helm_chart_v1beta1")
	// NO-OP: All data is already in Terraform state
}

func (r *SourceToolkitFluxcdIoHelmChartV1Beta1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_source_toolkit_fluxcd_io_helm_chart_v1beta1")

	var state SourceToolkitFluxcdIoHelmChartV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SourceToolkitFluxcdIoHelmChartV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("source.toolkit.fluxcd.io/v1beta1")
	goModel.Kind = utilities.Ptr("HelmChart")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *SourceToolkitFluxcdIoHelmChartV1Beta1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_source_toolkit_fluxcd_io_helm_chart_v1beta1")
	// NO-OP: Terraform removes the state automatically for us
}
