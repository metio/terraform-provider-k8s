/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package storage_kubeblocks_io_v1alpha1

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
	_ datasource.DataSource = &StorageKubeblocksIoStorageProviderV1Alpha1Manifest{}
)

func NewStorageKubeblocksIoStorageProviderV1Alpha1Manifest() datasource.DataSource {
	return &StorageKubeblocksIoStorageProviderV1Alpha1Manifest{}
}

type StorageKubeblocksIoStorageProviderV1Alpha1Manifest struct{}

type StorageKubeblocksIoStorageProviderV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		CsiDriverName           *string `tfsdk:"csi_driver_name" json:"csiDriverName,omitempty"`
		CsiDriverSecretTemplate *string `tfsdk:"csi_driver_secret_template" json:"csiDriverSecretTemplate,omitempty"`
		DatasafedConfigTemplate *string `tfsdk:"datasafed_config_template" json:"datasafedConfigTemplate,omitempty"`
		ParametersSchema        *struct {
			CredentialFields *[]string          `tfsdk:"credential_fields" json:"credentialFields,omitempty"`
			OpenAPIV3Schema  *map[string]string `tfsdk:"open_apiv3_schema" json:"openAPIV3Schema,omitempty"`
		} `tfsdk:"parameters_schema" json:"parametersSchema,omitempty"`
		PersistentVolumeClaimTemplate *string `tfsdk:"persistent_volume_claim_template" json:"persistentVolumeClaimTemplate,omitempty"`
		StorageClassTemplate          *string `tfsdk:"storage_class_template" json:"storageClassTemplate,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *StorageKubeblocksIoStorageProviderV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_storage_kubeblocks_io_storage_provider_v1alpha1_manifest"
}

func (r *StorageKubeblocksIoStorageProviderV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "StorageProvider comprises specifications that provide guidance on accessing remote storage. Currently the supported access methods are via a dedicated CSI driver or the 'datasafed' tool. In case of CSI driver, the specification expounds on provisioning PVCs for that driver. As for the 'datasafed' tool, the specification provides insights on generating the necessary configuration file.",
		MarkdownDescription: "StorageProvider comprises specifications that provide guidance on accessing remote storage. Currently the supported access methods are via a dedicated CSI driver or the 'datasafed' tool. In case of CSI driver, the specification expounds on provisioning PVCs for that driver. As for the 'datasafed' tool, the specification provides insights on generating the necessary configuration file.",
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
				Description:         "StorageProviderSpec defines the desired state of 'StorageProvider'.",
				MarkdownDescription: "StorageProviderSpec defines the desired state of 'StorageProvider'.",
				Attributes: map[string]schema.Attribute{
					"csi_driver_name": schema.StringAttribute{
						Description:         "Specifies the name of the CSI driver used to access remote storage. This field can be empty, it indicates that the storage is not accessible via CSI.",
						MarkdownDescription: "Specifies the name of the CSI driver used to access remote storage. This field can be empty, it indicates that the storage is not accessible via CSI.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"csi_driver_secret_template": schema.StringAttribute{
						Description:         "A Go template that used to render and generate 'k8s.io/api/core/v1.Secret' resources for a specific CSI driver. For example, 'accessKey' and 'secretKey' needed by CSI-S3 are stored in this 'Secret' resource.",
						MarkdownDescription: "A Go template that used to render and generate 'k8s.io/api/core/v1.Secret' resources for a specific CSI driver. For example, 'accessKey' and 'secretKey' needed by CSI-S3 are stored in this 'Secret' resource.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"datasafed_config_template": schema.StringAttribute{
						Description:         "A Go template used to render and generate 'k8s.io/api/core/v1.Secret'. This 'Secret' involves the configuration details required by the 'datasafed' tool to access remote storage. For example, the 'Secret' should contain 'endpoint', 'bucket', 'region', 'accessKey', 'secretKey', or something else for S3 storage. This field can be empty, it means this kind of storage is not accessible via the 'datasafed' tool.",
						MarkdownDescription: "A Go template used to render and generate 'k8s.io/api/core/v1.Secret'. This 'Secret' involves the configuration details required by the 'datasafed' tool to access remote storage. For example, the 'Secret' should contain 'endpoint', 'bucket', 'region', 'accessKey', 'secretKey', or something else for S3 storage. This field can be empty, it means this kind of storage is not accessible via the 'datasafed' tool.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"parameters_schema": schema.SingleNestedAttribute{
						Description:         "Describes the parameters required for storage. The parameters defined here can be referenced in the above templates, and 'kbcli' uses this definition for dynamic command-line parameter parsing.",
						MarkdownDescription: "Describes the parameters required for storage. The parameters defined here can be referenced in the above templates, and 'kbcli' uses this definition for dynamic command-line parameter parsing.",
						Attributes: map[string]schema.Attribute{
							"credential_fields": schema.ListAttribute{
								Description:         "Defines which parameters are credential fields, which need to be handled specifically. For instance, these should be stored in a 'Secret' instead of a 'ConfigMap'.",
								MarkdownDescription: "Defines which parameters are credential fields, which need to be handled specifically. For instance, these should be stored in a 'Secret' instead of a 'ConfigMap'.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"open_apiv3_schema": schema.MapAttribute{
								Description:         "Defines the parameters in OpenAPI V3.",
								MarkdownDescription: "Defines the parameters in OpenAPI V3.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"persistent_volume_claim_template": schema.StringAttribute{
						Description:         "A Go template that renders and generates 'k8s.io/api/core/v1.PersistentVolumeClaim' resources. This PVC can reference the 'StorageClass' created from 'storageClassTemplate', allowing Pods to access remote storage by mounting the PVC.",
						MarkdownDescription: "A Go template that renders and generates 'k8s.io/api/core/v1.PersistentVolumeClaim' resources. This PVC can reference the 'StorageClass' created from 'storageClassTemplate', allowing Pods to access remote storage by mounting the PVC.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"storage_class_template": schema.StringAttribute{
						Description:         "A Go template utilized to render and generate 'kubernetes.storage.k8s.io.v1.StorageClass' resources. The 'StorageClass' created by this template is aimed at using the CSI driver.",
						MarkdownDescription: "A Go template utilized to render and generate 'kubernetes.storage.k8s.io.v1.StorageClass' resources. The 'StorageClass' created by this template is aimed at using the CSI driver.",
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

func (r *StorageKubeblocksIoStorageProviderV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_storage_kubeblocks_io_storage_provider_v1alpha1_manifest")

	var model StorageKubeblocksIoStorageProviderV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("storage.kubeblocks.io/v1alpha1")
	model.Kind = pointer.String("StorageProvider")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
