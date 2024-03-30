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
	_ datasource.DataSource = &ImagingIngestionAlvearieOrgDicomwebIngestionServiceV1Alpha1Manifest{}
)

func NewImagingIngestionAlvearieOrgDicomwebIngestionServiceV1Alpha1Manifest() datasource.DataSource {
	return &ImagingIngestionAlvearieOrgDicomwebIngestionServiceV1Alpha1Manifest{}
}

type ImagingIngestionAlvearieOrgDicomwebIngestionServiceV1Alpha1Manifest struct{}

type ImagingIngestionAlvearieOrgDicomwebIngestionServiceV1Alpha1ManifestData struct {
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
		BucketConfigName              *string `tfsdk:"bucket_config_name" json:"bucketConfigName,omitempty"`
		BucketSecretName              *string `tfsdk:"bucket_secret_name" json:"bucketSecretName,omitempty"`
		DicomEventDrivenIngestionName *string `tfsdk:"dicom_event_driven_ingestion_name" json:"dicomEventDrivenIngestionName,omitempty"`
		ImagePullPolicy               *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
		ImagePullSecrets              *[]struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
		ProviderName *string `tfsdk:"provider_name" json:"providerName,omitempty"`
		StowService  *struct {
			Concurrency *int64  `tfsdk:"concurrency" json:"concurrency,omitempty"`
			Image       *string `tfsdk:"image" json:"image,omitempty"`
			MaxReplicas *int64  `tfsdk:"max_replicas" json:"maxReplicas,omitempty"`
			MinReplicas *int64  `tfsdk:"min_replicas" json:"minReplicas,omitempty"`
		} `tfsdk:"stow_service" json:"stowService,omitempty"`
		WadoService *struct {
			Concurrency *int64  `tfsdk:"concurrency" json:"concurrency,omitempty"`
			Image       *string `tfsdk:"image" json:"image,omitempty"`
			MaxReplicas *int64  `tfsdk:"max_replicas" json:"maxReplicas,omitempty"`
			MinReplicas *int64  `tfsdk:"min_replicas" json:"minReplicas,omitempty"`
		} `tfsdk:"wado_service" json:"wadoService,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ImagingIngestionAlvearieOrgDicomwebIngestionServiceV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_imaging_ingestion_alvearie_org_dicomweb_ingestion_service_v1alpha1_manifest"
}

func (r *ImagingIngestionAlvearieOrgDicomwebIngestionServiceV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Provides DICOMweb WADO-RS and STOW-RS access to a S3 bucket",
		MarkdownDescription: "Provides DICOMweb WADO-RS and STOW-RS access to a S3 bucket",
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
				Description:         "DicomwebIngestionServiceSpec defines the desired state of DicomwebIngestionService",
				MarkdownDescription: "DicomwebIngestionServiceSpec defines the desired state of DicomwebIngestionService",
				Attributes: map[string]schema.Attribute{
					"bucket_config_name": schema.StringAttribute{
						Description:         "Bucket Config Name",
						MarkdownDescription: "Bucket Config Name",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bucket_secret_name": schema.StringAttribute{
						Description:         "Bucket Secret Name",
						MarkdownDescription: "Bucket Secret Name",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"dicom_event_driven_ingestion_name": schema.StringAttribute{
						Description:         "DICOM Event Driven Ingestion Name",
						MarkdownDescription: "DICOM Event Driven Ingestion Name",
						Required:            true,
						Optional:            false,
						Computed:            false,
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

					"provider_name": schema.StringAttribute{
						Description:         "Provider Name",
						MarkdownDescription: "Provider Name",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"stow_service": schema.SingleNestedAttribute{
						Description:         "STOW Service Spec",
						MarkdownDescription: "STOW Service Spec",
						Attributes: map[string]schema.Attribute{
							"concurrency": schema.Int64Attribute{
								Description:         "Container Concurrency",
								MarkdownDescription: "Container Concurrency",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image": schema.StringAttribute{
								Description:         "Image",
								MarkdownDescription: "Image",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_replicas": schema.Int64Attribute{
								Description:         "Max Replicas",
								MarkdownDescription: "Max Replicas",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"min_replicas": schema.Int64Attribute{
								Description:         "Min Replicas",
								MarkdownDescription: "Min Replicas",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"wado_service": schema.SingleNestedAttribute{
						Description:         "WADO Service Spec",
						MarkdownDescription: "WADO Service Spec",
						Attributes: map[string]schema.Attribute{
							"concurrency": schema.Int64Attribute{
								Description:         "Container Concurrency",
								MarkdownDescription: "Container Concurrency",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image": schema.StringAttribute{
								Description:         "Image",
								MarkdownDescription: "Image",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_replicas": schema.Int64Attribute{
								Description:         "Max Replicas",
								MarkdownDescription: "Max Replicas",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"min_replicas": schema.Int64Attribute{
								Description:         "Min Replicas",
								MarkdownDescription: "Min Replicas",
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

func (r *ImagingIngestionAlvearieOrgDicomwebIngestionServiceV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_imaging_ingestion_alvearie_org_dicomweb_ingestion_service_v1alpha1_manifest")

	var model ImagingIngestionAlvearieOrgDicomwebIngestionServiceV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("imaging-ingestion.alvearie.org/v1alpha1")
	model.Kind = pointer.String("DicomwebIngestionService")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
