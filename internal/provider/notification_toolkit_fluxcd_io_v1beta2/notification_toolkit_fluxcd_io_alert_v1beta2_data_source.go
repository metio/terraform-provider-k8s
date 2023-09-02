/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package notification_toolkit_fluxcd_io_v1beta2

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &NotificationToolkitFluxcdIoAlertV1Beta2DataSource{}
	_ datasource.DataSourceWithConfigure = &NotificationToolkitFluxcdIoAlertV1Beta2DataSource{}
)

func NewNotificationToolkitFluxcdIoAlertV1Beta2DataSource() datasource.DataSource {
	return &NotificationToolkitFluxcdIoAlertV1Beta2DataSource{}
}

type NotificationToolkitFluxcdIoAlertV1Beta2DataSource struct {
	kubernetesClient dynamic.Interface
}

type NotificationToolkitFluxcdIoAlertV1Beta2DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

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

func (r *NotificationToolkitFluxcdIoAlertV1Beta2DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_notification_toolkit_fluxcd_io_alert_v1beta2"
}

func (r *NotificationToolkitFluxcdIoAlertV1Beta2DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Alert is the Schema for the alerts API",
		MarkdownDescription: "Alert is the Schema for the alerts API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "AlertSpec defines an alerting rule for events involving a list of objects.",
				MarkdownDescription: "AlertSpec defines an alerting rule for events involving a list of objects.",
				Attributes: map[string]schema.Attribute{
					"event_metadata": schema.MapAttribute{
						Description:         "EventMetadata is an optional field for adding metadata to events dispatched by the controller. This can be used for enhancing the context of the event. If a field would override one already present on the original event as generated by the emitter, then the override doesn't happen, i.e. the original value is preserved, and an info log is printed.",
						MarkdownDescription: "EventMetadata is an optional field for adding metadata to events dispatched by the controller. This can be used for enhancing the context of the event. If a field would override one already present on the original event as generated by the emitter, then the override doesn't happen, i.e. the original value is preserved, and an info log is printed.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"event_severity": schema.StringAttribute{
						Description:         "EventSeverity specifies how to filter events based on severity. If set to 'info' no events will be filtered.",
						MarkdownDescription: "EventSeverity specifies how to filter events based on severity. If set to 'info' no events will be filtered.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"event_sources": schema.ListNestedAttribute{
						Description:         "EventSources specifies how to filter events based on the involved object kind, name and namespace.",
						MarkdownDescription: "EventSources specifies how to filter events based on the involved object kind, name and namespace.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"api_version": schema.StringAttribute{
									Description:         "API version of the referent",
									MarkdownDescription: "API version of the referent",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"kind": schema.StringAttribute{
									Description:         "Kind of the referent",
									MarkdownDescription: "Kind of the referent",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"match_labels": schema.MapAttribute{
									Description:         "MatchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed. MatchLabels requires the name to be set to '*'.",
									MarkdownDescription: "MatchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed. MatchLabels requires the name to be set to '*'.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"name": schema.StringAttribute{
									Description:         "Name of the referent If multiple resources are targeted '*' may be set.",
									MarkdownDescription: "Name of the referent If multiple resources are targeted '*' may be set.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace of the referent",
									MarkdownDescription: "Namespace of the referent",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"exclusion_list": schema.ListAttribute{
						Description:         "ExclusionList specifies a list of Golang regular expressions to be used for excluding messages.",
						MarkdownDescription: "ExclusionList specifies a list of Golang regular expressions to be used for excluding messages.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"inclusion_list": schema.ListAttribute{
						Description:         "InclusionList specifies a list of Golang regular expressions to be used for including messages.",
						MarkdownDescription: "InclusionList specifies a list of Golang regular expressions to be used for including messages.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"provider_ref": schema.SingleNestedAttribute{
						Description:         "ProviderRef specifies which Provider this Alert should use.",
						MarkdownDescription: "ProviderRef specifies which Provider this Alert should use.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the referent.",
								MarkdownDescription: "Name of the referent.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"summary": schema.StringAttribute{
						Description:         "Summary holds a short description of the impact and affected cluster.",
						MarkdownDescription: "Summary holds a short description of the impact and affected cluster.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"suspend": schema.BoolAttribute{
						Description:         "Suspend tells the controller to suspend subsequent events handling for this Alert.",
						MarkdownDescription: "Suspend tells the controller to suspend subsequent events handling for this Alert.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *NotificationToolkitFluxcdIoAlertV1Beta2DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *NotificationToolkitFluxcdIoAlertV1Beta2DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_notification_toolkit_fluxcd_io_alert_v1beta2")

	var data NotificationToolkitFluxcdIoAlertV1Beta2DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "notification.toolkit.fluxcd.io", Version: "v1beta2", Resource: "Alert"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse NotificationToolkitFluxcdIoAlertV1Beta2DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("notification.toolkit.fluxcd.io/v1beta2")
	data.Kind = pointer.String("Alert")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
