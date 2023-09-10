/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package autoscaling_karmada_io_v1alpha1

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &AutoscalingKarmadaIoCronFederatedHpaV1Alpha1Manifest{}
)

func NewAutoscalingKarmadaIoCronFederatedHpaV1Alpha1Manifest() datasource.DataSource {
	return &AutoscalingKarmadaIoCronFederatedHpaV1Alpha1Manifest{}
}

type AutoscalingKarmadaIoCronFederatedHpaV1Alpha1Manifest struct{}

type AutoscalingKarmadaIoCronFederatedHpaV1Alpha1ManifestData struct {
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
		Rules *[]struct {
			FailedHistoryLimit     *int64  `tfsdk:"failed_history_limit" json:"failedHistoryLimit,omitempty"`
			Name                   *string `tfsdk:"name" json:"name,omitempty"`
			Schedule               *string `tfsdk:"schedule" json:"schedule,omitempty"`
			SuccessfulHistoryLimit *int64  `tfsdk:"successful_history_limit" json:"successfulHistoryLimit,omitempty"`
			Suspend                *bool   `tfsdk:"suspend" json:"suspend,omitempty"`
			TargetMaxReplicas      *int64  `tfsdk:"target_max_replicas" json:"targetMaxReplicas,omitempty"`
			TargetMinReplicas      *int64  `tfsdk:"target_min_replicas" json:"targetMinReplicas,omitempty"`
			TargetReplicas         *int64  `tfsdk:"target_replicas" json:"targetReplicas,omitempty"`
			TimeZone               *string `tfsdk:"time_zone" json:"timeZone,omitempty"`
		} `tfsdk:"rules" json:"rules,omitempty"`
		ScaleTargetRef *struct {
			ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"scale_target_ref" json:"scaleTargetRef,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AutoscalingKarmadaIoCronFederatedHpaV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_autoscaling_karmada_io_cron_federated_hpa_v1alpha1_manifest"
}

func (r *AutoscalingKarmadaIoCronFederatedHpaV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CronFederatedHPA represents a collection of repeating schedule to scale replica number of a specific workload. It can scale any resource implementing the scale subresource as well as FederatedHPA.",
		MarkdownDescription: "CronFederatedHPA represents a collection of repeating schedule to scale replica number of a specific workload. It can scale any resource implementing the scale subresource as well as FederatedHPA.",
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
				Description:         "Spec is the specification of the CronFederatedHPA.",
				MarkdownDescription: "Spec is the specification of the CronFederatedHPA.",
				Attributes: map[string]schema.Attribute{
					"rules": schema.ListNestedAttribute{
						Description:         "Rules contains a collection of schedules that declares when and how the referencing target resource should be scaled.",
						MarkdownDescription: "Rules contains a collection of schedules that declares when and how the referencing target resource should be scaled.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"failed_history_limit": schema.Int64Attribute{
									Description:         "FailedHistoryLimit represents the count of failed execution items for each rule. The value must be a positive integer. It defaults to 3.",
									MarkdownDescription: "FailedHistoryLimit represents the count of failed execution items for each rule. The value must be a positive integer. It defaults to 3.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(0),
										int64validator.AtMost(32),
									},
								},

								"name": schema.StringAttribute{
									Description:         "Name of the rule. Each rule in a CronFederatedHPA must have a unique name.  Note: the name will be used as an identifier to record its execution history. Changing the name will be considered as deleting the old rule and adding a new rule, that means the original execution history will be discarded.",
									MarkdownDescription: "Name of the rule. Each rule in a CronFederatedHPA must have a unique name.  Note: the name will be used as an identifier to record its execution history. Changing the name will be considered as deleting the old rule and adding a new rule, that means the original execution history will be discarded.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(32),
									},
								},

								"schedule": schema.StringAttribute{
									Description:         "Schedule is the cron expression that represents a periodical time. The syntax follows https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/#schedule-syntax.",
									MarkdownDescription: "Schedule is the cron expression that represents a periodical time. The syntax follows https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/#schedule-syntax.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"successful_history_limit": schema.Int64Attribute{
									Description:         "SuccessfulHistoryLimit represents the count of successful execution items for each rule. The value must be a positive integer. It defaults to 3.",
									MarkdownDescription: "SuccessfulHistoryLimit represents the count of successful execution items for each rule. The value must be a positive integer. It defaults to 3.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(1),
										int64validator.AtMost(32),
									},
								},

								"suspend": schema.BoolAttribute{
									Description:         "Suspend tells the controller to suspend subsequent executions. Defaults to false.",
									MarkdownDescription: "Suspend tells the controller to suspend subsequent executions. Defaults to false.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"target_max_replicas": schema.Int64Attribute{
									Description:         "TargetMaxReplicas is the target MaxReplicas to be set for FederatedHPA. Only needed when referencing resource is FederatedHPA. TargetMinReplicas and TargetMaxReplicas can be specified together or either one can be specified alone. nil means the MaxReplicas(.spec.maxReplicas) of the referencing FederatedHPA will not be updated.",
									MarkdownDescription: "TargetMaxReplicas is the target MaxReplicas to be set for FederatedHPA. Only needed when referencing resource is FederatedHPA. TargetMinReplicas and TargetMaxReplicas can be specified together or either one can be specified alone. nil means the MaxReplicas(.spec.maxReplicas) of the referencing FederatedHPA will not be updated.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"target_min_replicas": schema.Int64Attribute{
									Description:         "TargetMinReplicas is the target MinReplicas to be set for FederatedHPA. Only needed when referencing resource is FederatedHPA. TargetMinReplicas and TargetMaxReplicas can be specified together or either one can be specified alone. nil means the MinReplicas(.spec.minReplicas) of the referencing FederatedHPA will not be updated.",
									MarkdownDescription: "TargetMinReplicas is the target MinReplicas to be set for FederatedHPA. Only needed when referencing resource is FederatedHPA. TargetMinReplicas and TargetMaxReplicas can be specified together or either one can be specified alone. nil means the MinReplicas(.spec.minReplicas) of the referencing FederatedHPA will not be updated.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"target_replicas": schema.Int64Attribute{
									Description:         "TargetReplicas is the target replicas to be scaled for resources referencing by ScaleTargetRef of this CronFederatedHPA. Only needed when referencing resource is not FederatedHPA.",
									MarkdownDescription: "TargetReplicas is the target replicas to be scaled for resources referencing by ScaleTargetRef of this CronFederatedHPA. Only needed when referencing resource is not FederatedHPA.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"time_zone": schema.StringAttribute{
									Description:         "TimeZone for the giving schedule. If not specified, this will default to the time zone of the karmada-controller-manager process. Invalid TimeZone will be rejected when applying by karmada-webhook. see https://en.wikipedia.org/wiki/List_of_tz_database_time_zones for the all timezones.",
									MarkdownDescription: "TimeZone for the giving schedule. If not specified, this will default to the time zone of the karmada-controller-manager process. Invalid TimeZone will be rejected when applying by karmada-webhook. see https://en.wikipedia.org/wiki/List_of_tz_database_time_zones for the all timezones.",
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

					"scale_target_ref": schema.SingleNestedAttribute{
						Description:         "ScaleTargetRef points to the target resource to scale. Target resource could be any resource that implementing the scale subresource like Deployment, or FederatedHPA.",
						MarkdownDescription: "ScaleTargetRef points to the target resource to scale. Target resource could be any resource that implementing the scale subresource like Deployment, or FederatedHPA.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "apiVersion is the API version of the referent",
								MarkdownDescription: "apiVersion is the API version of the referent",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "kind is the kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								MarkdownDescription: "kind is the kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "name is the name of the referent; More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								MarkdownDescription: "name is the name of the referent; More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								Required:            true,
								Optional:            false,
								Computed:            false,
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

func (r *AutoscalingKarmadaIoCronFederatedHpaV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_autoscaling_karmada_io_cron_federated_hpa_v1alpha1_manifest")

	var model AutoscalingKarmadaIoCronFederatedHpaV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("autoscaling.karmada.io/v1alpha1")
	model.Kind = pointer.String("CronFederatedHPA")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"YAML Error: "+err.Error(),
		)
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
