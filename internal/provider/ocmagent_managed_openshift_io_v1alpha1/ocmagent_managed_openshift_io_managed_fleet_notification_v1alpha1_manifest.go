/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package ocmagent_managed_openshift_io_v1alpha1

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
	_ datasource.DataSource = &OcmagentManagedOpenshiftIoManagedFleetNotificationV1Alpha1Manifest{}
)

func NewOcmagentManagedOpenshiftIoManagedFleetNotificationV1Alpha1Manifest() datasource.DataSource {
	return &OcmagentManagedOpenshiftIoManagedFleetNotificationV1Alpha1Manifest{}
}

type OcmagentManagedOpenshiftIoManagedFleetNotificationV1Alpha1Manifest struct{}

type OcmagentManagedOpenshiftIoManagedFleetNotificationV1Alpha1ManifestData struct {
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
		FleetNotification *struct {
			LimitedSupport      *bool     `tfsdk:"limited_support" json:"limitedSupport,omitempty"`
			LogType             *string   `tfsdk:"log_type" json:"logType,omitempty"`
			Name                *string   `tfsdk:"name" json:"name,omitempty"`
			NotificationMessage *string   `tfsdk:"notification_message" json:"notificationMessage,omitempty"`
			References          *[]string `tfsdk:"references" json:"references,omitempty"`
			ResendWait          *int64    `tfsdk:"resend_wait" json:"resendWait,omitempty"`
			Severity            *string   `tfsdk:"severity" json:"severity,omitempty"`
			Summary             *string   `tfsdk:"summary" json:"summary,omitempty"`
		} `tfsdk:"fleet_notification" json:"fleetNotification,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OcmagentManagedOpenshiftIoManagedFleetNotificationV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_ocmagent_managed_openshift_io_managed_fleet_notification_v1alpha1_manifest"
}

func (r *OcmagentManagedOpenshiftIoManagedFleetNotificationV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ManagedFleetNotification is the Schema for the managedfleetnotifications API",
		MarkdownDescription: "ManagedFleetNotification is the Schema for the managedfleetnotifications API",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"fleet_notification": schema.SingleNestedAttribute{
						Description:         "FleetNotification defines the desired spec of ManagedFleetNotification",
						MarkdownDescription: "FleetNotification defines the desired spec of ManagedFleetNotification",
						Attributes: map[string]schema.Attribute{
							"limited_support": schema.BoolAttribute{
								Description:         "Whether or not limited support should be sent for this notification",
								MarkdownDescription: "Whether or not limited support should be sent for this notification",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"log_type": schema.StringAttribute{
								Description:         "LogType is a categorization property that can be used to group service logs for aggregation and managing notification preferences.",
								MarkdownDescription: "LogType is a categorization property that can be used to group service logs for aggregation and managing notification preferences.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "The name of the notification used to associate with an alert",
								MarkdownDescription: "The name of the notification used to associate with an alert",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"notification_message": schema.StringAttribute{
								Description:         "The body text of the notification when the alert is active",
								MarkdownDescription: "The body text of the notification when the alert is active",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"references": schema.ListAttribute{
								Description:         "References useful for context or remediation - this could be links to documentation, KB articles, etc",
								MarkdownDescription: "References useful for context or remediation - this could be links to documentation, KB articles, etc",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resend_wait": schema.Int64Attribute{
								Description:         "Measured in hours. The minimum time interval that must elapse between active notifications",
								MarkdownDescription: "Measured in hours. The minimum time interval that must elapse between active notifications",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"severity": schema.StringAttribute{
								Description:         "Re-use the severity definitation in managednotification_types",
								MarkdownDescription: "Re-use the severity definitation in managednotification_types",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Debug", "Info", "Warning", "Error", "Fatal"),
								},
							},

							"summary": schema.StringAttribute{
								Description:         "The summary line of the notification",
								MarkdownDescription: "The summary line of the notification",
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *OcmagentManagedOpenshiftIoManagedFleetNotificationV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_ocmagent_managed_openshift_io_managed_fleet_notification_v1alpha1_manifest")

	var model OcmagentManagedOpenshiftIoManagedFleetNotificationV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("ocmagent.managed.openshift.io/v1alpha1")
	model.Kind = pointer.String("ManagedFleetNotification")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
