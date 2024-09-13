/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package ptp_openshift_io_v1

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
	_ datasource.DataSource = &PtpOpenshiftIoPtpOperatorConfigV1Manifest{}
)

func NewPtpOpenshiftIoPtpOperatorConfigV1Manifest() datasource.DataSource {
	return &PtpOpenshiftIoPtpOperatorConfigV1Manifest{}
}

type PtpOpenshiftIoPtpOperatorConfigV1Manifest struct{}

type PtpOpenshiftIoPtpOperatorConfigV1ManifestData struct {
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
		DaemonNodeSelector *map[string]string `tfsdk:"daemon_node_selector" json:"daemonNodeSelector,omitempty"`
		Plugins            *map[string]string `tfsdk:"plugins" json:"plugins,omitempty"`
		PtpEventConfig     *struct {
			ApiVersion           *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			EnableEventPublisher *bool   `tfsdk:"enable_event_publisher" json:"enableEventPublisher,omitempty"`
			StorageType          *string `tfsdk:"storage_type" json:"storageType,omitempty"`
			TransportHost        *string `tfsdk:"transport_host" json:"transportHost,omitempty"`
		} `tfsdk:"ptp_event_config" json:"ptpEventConfig,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *PtpOpenshiftIoPtpOperatorConfigV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_ptp_openshift_io_ptp_operator_config_v1_manifest"
}

func (r *PtpOpenshiftIoPtpOperatorConfigV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "PtpOperatorConfig is the Schema for the ptpoperatorconfigs API",
		MarkdownDescription: "PtpOperatorConfig is the Schema for the ptpoperatorconfigs API",
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
				Description:         "PtpOperatorConfigSpec defines the desired state of PtpOperatorConfig.",
				MarkdownDescription: "PtpOperatorConfigSpec defines the desired state of PtpOperatorConfig.",
				Attributes: map[string]schema.Attribute{
					"daemon_node_selector": schema.MapAttribute{
						Description:         "DaemonNodeSelector specifies the node selector for the linuxptp daemon. This is a map of key-value pairs used to select the nodes where the linuxptp daemon will run. If empty {}, the linuxptp daemon will be deployed on each node of the cluster.",
						MarkdownDescription: "DaemonNodeSelector specifies the node selector for the linuxptp daemon. This is a map of key-value pairs used to select the nodes where the linuxptp daemon will run. If empty {}, the linuxptp daemon will be deployed on each node of the cluster.",
						ElementType:         types.StringType,
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"plugins": schema.MapAttribute{
						Description:         "EnabledPlugins is a map of plugin names to their configuration settings. Each entry in the map specifies the configuration for a specific plugin. This field is optional and can be omitted if no plugins are enabled.",
						MarkdownDescription: "EnabledPlugins is a map of plugin names to their configuration settings. Each entry in the map specifies the configuration for a specific plugin. This field is optional and can be omitted if no plugins are enabled.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ptp_event_config": schema.SingleNestedAttribute{
						Description:         "EventConfig contains the configuration settings for the PTP event sidecar. This field is optional and can be omitted if event sidecar configuration is not required.",
						MarkdownDescription: "EventConfig contains the configuration settings for the PTP event sidecar. This field is optional and can be omitted if event sidecar configuration is not required.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "ApiVersion is used to determine which API is used for the event service 1.0: default version. event service is mapped to internal REST-API. 2.x: event service is mapped to O-RAN v3.0 Compliant O-Cloud Notification REST-API.",
								MarkdownDescription: "ApiVersion is used to determine which API is used for the event service 1.0: default version. event service is mapped to internal REST-API. 2.x: event service is mapped to O-RAN v3.0 Compliant O-Cloud Notification REST-API.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_event_publisher": schema.BoolAttribute{
								Description:         "EnableEventPublisher will deploy event proxy as a sidecar",
								MarkdownDescription: "EnableEventPublisher will deploy event proxy as a sidecar",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"storage_type": schema.StringAttribute{
								Description:         "StorageType is the type of storage to store subscription data",
								MarkdownDescription: "StorageType is the type of storage to store subscription data",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"transport_host": schema.StringAttribute{
								Description:         "TransportHost format is <protocol>://<transport-service>.<namespace>.svc.cluster.local:<transport-port> Example HTTP transport: 'http://ptp-event-publisher-service-NODE_NAME.openshift-ptp.svc.cluster.local:9043'",
								MarkdownDescription: "TransportHost format is <protocol>://<transport-service>.<namespace>.svc.cluster.local:<transport-port> Example HTTP transport: 'http://ptp-event-publisher-service-NODE_NAME.openshift-ptp.svc.cluster.local:9043'",
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *PtpOpenshiftIoPtpOperatorConfigV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_ptp_openshift_io_ptp_operator_config_v1_manifest")

	var model PtpOpenshiftIoPtpOperatorConfigV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("ptp.openshift.io/v1")
	model.Kind = pointer.String("PtpOperatorConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
