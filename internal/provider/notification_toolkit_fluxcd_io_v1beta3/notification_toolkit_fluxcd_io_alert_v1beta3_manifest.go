/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package notification_toolkit_fluxcd_io_v1beta3

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
	_ datasource.DataSource = &NotificationToolkitFluxcdIoAlertV1Beta3Manifest{}
)

func NewNotificationToolkitFluxcdIoAlertV1Beta3Manifest() datasource.DataSource {
	return &NotificationToolkitFluxcdIoAlertV1Beta3Manifest{}
}

type NotificationToolkitFluxcdIoAlertV1Beta3Manifest struct{}

type NotificationToolkitFluxcdIoAlertV1Beta3ManifestData struct {
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
		EventMetadata *map[string]string `tfsdk:"event_metadata" json:"eventMetadata,omitempty"`
		EventSeverity *string            `tfsdk:"event_severity" json:"eventSeverity,omitempty"`
		EventSources  *[]struct {
			ApiVersion  *string            `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind        *string            `tfsdk:"kind" json:"kind,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			Name        *string            `tfsdk:"name" json:"name,omitempty"`
			Namespace   *string            `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"event_sources" json:"eventSources,omitempty"`
		ExclusionList *[]string `tfsdk:"exclusion_list" json:"exclusionList,omitempty"`
		InclusionList *[]string `tfsdk:"inclusion_list" json:"inclusionList,omitempty"`
		ProviderRef   *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"provider_ref" json:"providerRef,omitempty"`
		Summary *string `tfsdk:"summary" json:"summary,omitempty"`
		Suspend *bool   `tfsdk:"suspend" json:"suspend,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *NotificationToolkitFluxcdIoAlertV1Beta3Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_notification_toolkit_fluxcd_io_alert_v1beta3_manifest"
}

func (r *NotificationToolkitFluxcdIoAlertV1Beta3Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Alert is the Schema for the alerts API",
		MarkdownDescription: "Alert is the Schema for the alerts API",
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
				Description:         "AlertSpec defines an alerting rule for events involving a list of objects.",
				MarkdownDescription: "AlertSpec defines an alerting rule for events involving a list of objects.",
				Attributes: map[string]schema.Attribute{
					"event_metadata": schema.MapAttribute{
						Description:         "EventMetadata is an optional field for adding metadata to events dispatched by thecontroller. This can be used for enhancing the context of the event. If a fieldwould override one already present on the original event as generated by the emitter,then the override doesn't happen, i.e. the original value is preserved, and an infolog is printed.",
						MarkdownDescription: "EventMetadata is an optional field for adding metadata to events dispatched by thecontroller. This can be used for enhancing the context of the event. If a fieldwould override one already present on the original event as generated by the emitter,then the override doesn't happen, i.e. the original value is preserved, and an infolog is printed.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"event_severity": schema.StringAttribute{
						Description:         "EventSeverity specifies how to filter events based on severity.If set to 'info' no events will be filtered.",
						MarkdownDescription: "EventSeverity specifies how to filter events based on severity.If set to 'info' no events will be filtered.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("info", "error"),
						},
					},

					"event_sources": schema.ListNestedAttribute{
						Description:         "EventSources specifies how to filter events basedon the involved object kind, name and namespace.",
						MarkdownDescription: "EventSources specifies how to filter events basedon the involved object kind, name and namespace.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"api_version": schema.StringAttribute{
									Description:         "API version of the referent",
									MarkdownDescription: "API version of the referent",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"kind": schema.StringAttribute{
									Description:         "Kind of the referent",
									MarkdownDescription: "Kind of the referent",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("Bucket", "GitRepository", "Kustomization", "HelmRelease", "HelmChart", "HelmRepository", "ImageRepository", "ImagePolicy", "ImageUpdateAutomation", "OCIRepository"),
									},
								},

								"match_labels": schema.MapAttribute{
									Description:         "MatchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.MatchLabels requires the name to be set to '*'.",
									MarkdownDescription: "MatchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.MatchLabels requires the name to be set to '*'.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name of the referentIf multiple resources are targeted '*' may be set.",
									MarkdownDescription: "Name of the referentIf multiple resources are targeted '*' may be set.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(53),
									},
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace of the referent",
									MarkdownDescription: "Namespace of the referent",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(53),
									},
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"exclusion_list": schema.ListAttribute{
						Description:         "ExclusionList specifies a list of Golang regular expressionsto be used for excluding messages.",
						MarkdownDescription: "ExclusionList specifies a list of Golang regular expressionsto be used for excluding messages.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"inclusion_list": schema.ListAttribute{
						Description:         "InclusionList specifies a list of Golang regular expressionsto be used for including messages.",
						MarkdownDescription: "InclusionList specifies a list of Golang regular expressionsto be used for including messages.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"provider_ref": schema.SingleNestedAttribute{
						Description:         "ProviderRef specifies which Provider this Alert should use.",
						MarkdownDescription: "ProviderRef specifies which Provider this Alert should use.",
						Attributes: map[string]schema.Attribute{
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

					"summary": schema.StringAttribute{
						Description:         "Summary holds a short description of the impact and affected cluster.",
						MarkdownDescription: "Summary holds a short description of the impact and affected cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(255),
						},
					},

					"suspend": schema.BoolAttribute{
						Description:         "Suspend tells the controller to suspend subsequentevents handling for this Alert.",
						MarkdownDescription: "Suspend tells the controller to suspend subsequentevents handling for this Alert.",
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

func (r *NotificationToolkitFluxcdIoAlertV1Beta3Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_notification_toolkit_fluxcd_io_alert_v1beta3_manifest")

	var model NotificationToolkitFluxcdIoAlertV1Beta3ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("notification.toolkit.fluxcd.io/v1beta3")
	model.Kind = pointer.String("Alert")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
