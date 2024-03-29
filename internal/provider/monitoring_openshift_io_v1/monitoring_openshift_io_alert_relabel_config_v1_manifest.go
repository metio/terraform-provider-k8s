/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package monitoring_openshift_io_v1

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
	_ datasource.DataSource = &MonitoringOpenshiftIoAlertRelabelConfigV1Manifest{}
)

func NewMonitoringOpenshiftIoAlertRelabelConfigV1Manifest() datasource.DataSource {
	return &MonitoringOpenshiftIoAlertRelabelConfigV1Manifest{}
}

type MonitoringOpenshiftIoAlertRelabelConfigV1Manifest struct{}

type MonitoringOpenshiftIoAlertRelabelConfigV1ManifestData struct {
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
		Configs *[]struct {
			Action       *string   `tfsdk:"action" json:"action,omitempty"`
			Modulus      *int64    `tfsdk:"modulus" json:"modulus,omitempty"`
			Regex        *string   `tfsdk:"regex" json:"regex,omitempty"`
			Replacement  *string   `tfsdk:"replacement" json:"replacement,omitempty"`
			Separator    *string   `tfsdk:"separator" json:"separator,omitempty"`
			SourceLabels *[]string `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
			TargetLabel  *string   `tfsdk:"target_label" json:"targetLabel,omitempty"`
		} `tfsdk:"configs" json:"configs,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *MonitoringOpenshiftIoAlertRelabelConfigV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_monitoring_openshift_io_alert_relabel_config_v1_manifest"
}

func (r *MonitoringOpenshiftIoAlertRelabelConfigV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "AlertRelabelConfig defines a set of relabel configs for alerts.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		MarkdownDescription: "AlertRelabelConfig defines a set of relabel configs for alerts.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
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
				Description:         "spec describes the desired state of this AlertRelabelConfig object.",
				MarkdownDescription: "spec describes the desired state of this AlertRelabelConfig object.",
				Attributes: map[string]schema.Attribute{
					"configs": schema.ListNestedAttribute{
						Description:         "configs is a list of sequentially evaluated alert relabel configs.",
						MarkdownDescription: "configs is a list of sequentially evaluated alert relabel configs.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action": schema.StringAttribute{
									Description:         "action to perform based on regex matching. Must be one of: 'Replace', 'Keep', 'Drop', 'HashMod', 'LabelMap', 'LabelDrop', or 'LabelKeep'. Default is: 'Replace'",
									MarkdownDescription: "action to perform based on regex matching. Must be one of: 'Replace', 'Keep', 'Drop', 'HashMod', 'LabelMap', 'LabelDrop', or 'LabelKeep'. Default is: 'Replace'",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("Replace", "Keep", "Drop", "HashMod", "LabelMap", "LabelDrop", "LabelKeep"),
									},
								},

								"modulus": schema.Int64Attribute{
									Description:         "modulus to take of the hash of the source label values.  This can be combined with the 'HashMod' action to set 'target_label' to the 'modulus' of a hash of the concatenated 'source_labels'. This is only valid if sourceLabels is not empty and action is not 'LabelKeep' or 'LabelDrop'.",
									MarkdownDescription: "modulus to take of the hash of the source label values.  This can be combined with the 'HashMod' action to set 'target_label' to the 'modulus' of a hash of the concatenated 'source_labels'. This is only valid if sourceLabels is not empty and action is not 'LabelKeep' or 'LabelDrop'.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"regex": schema.StringAttribute{
									Description:         "regex against which the extracted value is matched. Default is: '(.*)' regex is required for all actions except 'HashMod'",
									MarkdownDescription: "regex against which the extracted value is matched. Default is: '(.*)' regex is required for all actions except 'HashMod'",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(2048),
									},
								},

								"replacement": schema.StringAttribute{
									Description:         "replacement value against which a regex replace is performed if the regular expression matches. This is required if the action is 'Replace' or 'LabelMap' and forbidden for actions 'LabelKeep' and 'LabelDrop'. Regex capture groups are available. Default is: '$1'",
									MarkdownDescription: "replacement value against which a regex replace is performed if the regular expression matches. This is required if the action is 'Replace' or 'LabelMap' and forbidden for actions 'LabelKeep' and 'LabelDrop'. Regex capture groups are available. Default is: '$1'",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(2048),
									},
								},

								"separator": schema.StringAttribute{
									Description:         "separator placed between concatenated source label values. When omitted, Prometheus will use its default value of ';'.",
									MarkdownDescription: "separator placed between concatenated source label values. When omitted, Prometheus will use its default value of ';'.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(2048),
									},
								},

								"source_labels": schema.ListAttribute{
									Description:         "sourceLabels select values from existing labels. Their content is concatenated using the configured separator and matched against the configured regular expression for the 'Replace', 'Keep', and 'Drop' actions. Not allowed for actions 'LabelKeep' and 'LabelDrop'.",
									MarkdownDescription: "sourceLabels select values from existing labels. Their content is concatenated using the configured separator and matched against the configured regular expression for the 'Replace', 'Keep', and 'Drop' actions. Not allowed for actions 'LabelKeep' and 'LabelDrop'.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"target_label": schema.StringAttribute{
									Description:         "targetLabel to which the resulting value is written in a 'Replace' action. It is required for 'Replace' and 'HashMod' actions and forbidden for actions 'LabelKeep' and 'LabelDrop'. Regex capture groups are available.",
									MarkdownDescription: "targetLabel to which the resulting value is written in a 'Replace' action. It is required for 'Replace' and 'HashMod' actions and forbidden for actions 'LabelKeep' and 'LabelDrop'. Regex capture groups are available.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(2048),
									},
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

func (r *MonitoringOpenshiftIoAlertRelabelConfigV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_monitoring_openshift_io_alert_relabel_config_v1_manifest")

	var model MonitoringOpenshiftIoAlertRelabelConfigV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("monitoring.openshift.io/v1")
	model.Kind = pointer.String("AlertRelabelConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
