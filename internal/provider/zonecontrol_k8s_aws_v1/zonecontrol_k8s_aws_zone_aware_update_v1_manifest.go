/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package zonecontrol_k8s_aws_v1

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
	_ datasource.DataSource = &ZonecontrolK8SAwsZoneAwareUpdateV1Manifest{}
)

func NewZonecontrolK8SAwsZoneAwareUpdateV1Manifest() datasource.DataSource {
	return &ZonecontrolK8SAwsZoneAwareUpdateV1Manifest{}
}

type ZonecontrolK8SAwsZoneAwareUpdateV1Manifest struct{}

type ZonecontrolK8SAwsZoneAwareUpdateV1ManifestData struct {
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
		DryRun            *bool   `tfsdk:"dry_run" json:"dryRun,omitempty"`
		ExponentialFactor *string `tfsdk:"exponential_factor" json:"exponentialFactor,omitempty"`
		IgnoreAlarm       *bool   `tfsdk:"ignore_alarm" json:"ignoreAlarm,omitempty"`
		MaxUnavailable    *string `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
		PauseRolloutAlarm *string `tfsdk:"pause_rollout_alarm" json:"pauseRolloutAlarm,omitempty"`
		Statefulset       *string `tfsdk:"statefulset" json:"statefulset,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ZonecontrolK8SAwsZoneAwareUpdateV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_zonecontrol_k8s_aws_zone_aware_update_v1_manifest"
}

func (r *ZonecontrolK8SAwsZoneAwareUpdateV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ZoneAwareUpdate is the Schema for the zoneawareupdates API",
		MarkdownDescription: "ZoneAwareUpdate is the Schema for the zoneawareupdates API",
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
				Description:         "ZoneAwareUpdateSpec defines the desired state of ZoneAwareUpdate",
				MarkdownDescription: "ZoneAwareUpdateSpec defines the desired state of ZoneAwareUpdate",
				Attributes: map[string]schema.Attribute{
					"dry_run": schema.BoolAttribute{
						Description:         "Dryn-run mode that can be used to test the new controller before enable it",
						MarkdownDescription: "Dryn-run mode that can be used to test the new controller before enable it",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"exponential_factor": schema.StringAttribute{
						Description:         "The exponential growth rate in float string. Default value is 2.0. It's possible to disable exponential updates by setting the ExponentialFactor to 0. In this case, the number of pods updated at each step is defined only by the MaxUnavailable param.",
						MarkdownDescription: "The exponential growth rate in float string. Default value is 2.0. It's possible to disable exponential updates by setting the ExponentialFactor to 0. In this case, the number of pods updated at each step is defined only by the MaxUnavailable param.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ignore_alarm": schema.BoolAttribute{
						Description:         "Flag to ignore the PauseRolloutAlarm (default false)",
						MarkdownDescription: "Flag to ignore the PauseRolloutAlarm (default false)",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"max_unavailable": schema.StringAttribute{
						Description:         "Max number (or %) of pods that can be updated at the same time.",
						MarkdownDescription: "Max number (or %) of pods that can be updated at the same time.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"pause_rollout_alarm": schema.StringAttribute{
						Description:         "CW alarm name used to pause/skip updates. Alarm should be on the same account and region.",
						MarkdownDescription: "CW alarm name used to pause/skip updates. Alarm should be on the same account and region.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"statefulset": schema.StringAttribute{
						Description:         "The name of the StatefulSet for which the ZoneAwareUpdate applies to.",
						MarkdownDescription: "The name of the StatefulSet for which the ZoneAwareUpdate applies to.",
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

func (r *ZonecontrolK8SAwsZoneAwareUpdateV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_zonecontrol_k8s_aws_zone_aware_update_v1_manifest")

	var model ZonecontrolK8SAwsZoneAwareUpdateV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("zonecontrol.k8s.aws/v1")
	model.Kind = pointer.String("ZoneAwareUpdate")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
