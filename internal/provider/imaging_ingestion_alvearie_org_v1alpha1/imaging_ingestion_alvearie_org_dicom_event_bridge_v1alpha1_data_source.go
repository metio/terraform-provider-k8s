/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package imaging_ingestion_alvearie_org_v1alpha1

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
	_ datasource.DataSource              = &ImagingIngestionAlvearieOrgDicomEventBridgeV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &ImagingIngestionAlvearieOrgDicomEventBridgeV1Alpha1DataSource{}
)

func NewImagingIngestionAlvearieOrgDicomEventBridgeV1Alpha1DataSource() datasource.DataSource {
	return &ImagingIngestionAlvearieOrgDicomEventBridgeV1Alpha1DataSource{}
}

type ImagingIngestionAlvearieOrgDicomEventBridgeV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type ImagingIngestionAlvearieOrgDicomEventBridgeV1Alpha1DataSourceData struct {
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
		DicomEventDrivenIngestionName *string `tfsdk:"dicom_event_driven_ingestion_name" json:"dicomEventDrivenIngestionName,omitempty"`
		EdgeMailbox                   *string `tfsdk:"edge_mailbox" json:"edgeMailbox,omitempty"`
		EventBridge                   *struct {
			Image *string `tfsdk:"image" json:"image,omitempty"`
		} `tfsdk:"event_bridge" json:"eventBridge,omitempty"`
		ImagePullPolicy  *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
		ImagePullSecrets *[]struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
		NatsSecure      *bool   `tfsdk:"nats_secure" json:"natsSecure,omitempty"`
		NatsSubjectRoot *string `tfsdk:"nats_subject_root" json:"natsSubjectRoot,omitempty"`
		NatsTokenSecret *string `tfsdk:"nats_token_secret" json:"natsTokenSecret,omitempty"`
		NatsUrl         *string `tfsdk:"nats_url" json:"natsUrl,omitempty"`
		Role            *string `tfsdk:"role" json:"role,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ImagingIngestionAlvearieOrgDicomEventBridgeV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_imaging_ingestion_alvearie_org_dicom_event_bridge_v1alpha1"
}

func (r *ImagingIngestionAlvearieOrgDicomEventBridgeV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DicomEventBridge is the Schema for the dicomeventbridges API",
		MarkdownDescription: "DicomEventBridge is the Schema for the dicomeventbridges API",
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
				Description:         "DicomEventBridgeSpec defines the desired state of DicomEventBridge",
				MarkdownDescription: "DicomEventBridgeSpec defines the desired state of DicomEventBridge",
				Attributes: map[string]schema.Attribute{
					"dicom_event_driven_ingestion_name": schema.StringAttribute{
						Description:         "DICOM Event Driven Ingestion Name",
						MarkdownDescription: "DICOM Event Driven Ingestion Name",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"edge_mailbox": schema.StringAttribute{
						Description:         "Event Bridge Edge Mailbox. Required when Role is edge.",
						MarkdownDescription: "Event Bridge Edge Mailbox. Required when Role is edge.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"event_bridge": schema.SingleNestedAttribute{
						Description:         "Event Bridge Deployment Spec",
						MarkdownDescription: "Event Bridge Deployment Spec",
						Attributes: map[string]schema.Attribute{
							"image": schema.StringAttribute{
								Description:         "Image",
								MarkdownDescription: "Image",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"image_pull_policy": schema.StringAttribute{
						Description:         "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
						MarkdownDescription: "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"image_pull_secrets": schema.ListNestedAttribute{
						Description:         "Image Pull Secrets",
						MarkdownDescription: "Image Pull Secrets",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
									MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

					"nats_secure": schema.BoolAttribute{
						Description:         "Make NATS Connection Secure",
						MarkdownDescription: "Make NATS Connection Secure",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"nats_subject_root": schema.StringAttribute{
						Description:         "NATS Subject Root",
						MarkdownDescription: "NATS Subject Root",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"nats_token_secret": schema.StringAttribute{
						Description:         "NATS Token Secret Name",
						MarkdownDescription: "NATS Token Secret Name",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"nats_url": schema.StringAttribute{
						Description:         "NATS URL",
						MarkdownDescription: "NATS URL",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"role": schema.StringAttribute{
						Description:         "Event Bridge Role",
						MarkdownDescription: "Event Bridge Role",
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

func (r *ImagingIngestionAlvearieOrgDicomEventBridgeV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *ImagingIngestionAlvearieOrgDicomEventBridgeV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_imaging_ingestion_alvearie_org_dicom_event_bridge_v1alpha1")

	var data ImagingIngestionAlvearieOrgDicomEventBridgeV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "imaging-ingestion.alvearie.org", Version: "v1alpha1", Resource: "DicomEventBridge"}).
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

	var readResponse ImagingIngestionAlvearieOrgDicomEventBridgeV1Alpha1DataSourceData
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
	data.ApiVersion = pointer.String("imaging-ingestion.alvearie.org/v1alpha1")
	data.Kind = pointer.String("DicomEventBridge")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}