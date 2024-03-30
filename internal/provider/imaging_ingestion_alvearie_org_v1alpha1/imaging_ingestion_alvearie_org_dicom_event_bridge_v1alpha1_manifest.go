/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package imaging_ingestion_alvearie_org_v1alpha1

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
	_ datasource.DataSource = &ImagingIngestionAlvearieOrgDicomEventBridgeV1Alpha1Manifest{}
)

func NewImagingIngestionAlvearieOrgDicomEventBridgeV1Alpha1Manifest() datasource.DataSource {
	return &ImagingIngestionAlvearieOrgDicomEventBridgeV1Alpha1Manifest{}
}

type ImagingIngestionAlvearieOrgDicomEventBridgeV1Alpha1Manifest struct{}

type ImagingIngestionAlvearieOrgDicomEventBridgeV1Alpha1ManifestData struct {
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

func (r *ImagingIngestionAlvearieOrgDicomEventBridgeV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_imaging_ingestion_alvearie_org_dicom_event_bridge_v1alpha1_manifest"
}

func (r *ImagingIngestionAlvearieOrgDicomEventBridgeV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DicomEventBridge is the Schema for the dicomeventbridges API",
		MarkdownDescription: "DicomEventBridge is the Schema for the dicomeventbridges API",
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
				Description:         "DicomEventBridgeSpec defines the desired state of DicomEventBridge",
				MarkdownDescription: "DicomEventBridgeSpec defines the desired state of DicomEventBridge",
				Attributes: map[string]schema.Attribute{
					"dicom_event_driven_ingestion_name": schema.StringAttribute{
						Description:         "DICOM Event Driven Ingestion Name",
						MarkdownDescription: "DICOM Event Driven Ingestion Name",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"edge_mailbox": schema.StringAttribute{
						Description:         "Event Bridge Edge Mailbox. Required when Role is edge.",
						MarkdownDescription: "Event Bridge Edge Mailbox. Required when Role is edge.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"event_bridge": schema.SingleNestedAttribute{
						Description:         "Event Bridge Deployment Spec",
						MarkdownDescription: "Event Bridge Deployment Spec",
						Attributes: map[string]schema.Attribute{
							"image": schema.StringAttribute{
								Description:         "Image",
								MarkdownDescription: "Image",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"image_pull_policy": schema.StringAttribute{
						Description:         "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
						MarkdownDescription: "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
						Required:            false,
						Optional:            true,
						Computed:            false,
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
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"nats_secure": schema.BoolAttribute{
						Description:         "Make NATS Connection Secure",
						MarkdownDescription: "Make NATS Connection Secure",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"nats_subject_root": schema.StringAttribute{
						Description:         "NATS Subject Root",
						MarkdownDescription: "NATS Subject Root",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"nats_token_secret": schema.StringAttribute{
						Description:         "NATS Token Secret Name",
						MarkdownDescription: "NATS Token Secret Name",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"nats_url": schema.StringAttribute{
						Description:         "NATS URL",
						MarkdownDescription: "NATS URL",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"role": schema.StringAttribute{
						Description:         "Event Bridge Role",
						MarkdownDescription: "Event Bridge Role",
						Required:            true,
						Optional:            false,
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

func (r *ImagingIngestionAlvearieOrgDicomEventBridgeV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_imaging_ingestion_alvearie_org_dicom_event_bridge_v1alpha1_manifest")

	var model ImagingIngestionAlvearieOrgDicomEventBridgeV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("imaging-ingestion.alvearie.org/v1alpha1")
	model.Kind = pointer.String("DicomEventBridge")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
