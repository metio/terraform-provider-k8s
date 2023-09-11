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
	_ datasource.DataSource              = &ImagingIngestionAlvearieOrgDicomInstanceBindingV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &ImagingIngestionAlvearieOrgDicomInstanceBindingV1Alpha1DataSource{}
)

func NewImagingIngestionAlvearieOrgDicomInstanceBindingV1Alpha1DataSource() datasource.DataSource {
	return &ImagingIngestionAlvearieOrgDicomInstanceBindingV1Alpha1DataSource{}
}

type ImagingIngestionAlvearieOrgDicomInstanceBindingV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type ImagingIngestionAlvearieOrgDicomInstanceBindingV1Alpha1DataSourceData struct {
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
		BindingConfigName             *string `tfsdk:"binding_config_name" json:"bindingConfigName,omitempty"`
		BindingSecretName             *string `tfsdk:"binding_secret_name" json:"bindingSecretName,omitempty"`
		DicomEventDrivenIngestionName *string `tfsdk:"dicom_event_driven_ingestion_name" json:"dicomEventDrivenIngestionName,omitempty"`
		ImagePullPolicy               *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
		ImagePullSecrets              *[]struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
		InstanceBinding *struct {
			Concurrency *int64  `tfsdk:"concurrency" json:"concurrency,omitempty"`
			Image       *string `tfsdk:"image" json:"image,omitempty"`
			MaxReplicas *int64  `tfsdk:"max_replicas" json:"maxReplicas,omitempty"`
			MinReplicas *int64  `tfsdk:"min_replicas" json:"minReplicas,omitempty"`
		} `tfsdk:"instance_binding" json:"instanceBinding,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ImagingIngestionAlvearieOrgDicomInstanceBindingV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_imaging_ingestion_alvearie_org_dicom_instance_binding_v1alpha1"
}

func (r *ImagingIngestionAlvearieOrgDicomInstanceBindingV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Fan-out notification of DICOM instance data (STOW-RS)",
		MarkdownDescription: "Fan-out notification of DICOM instance data (STOW-RS)",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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
				Description:         "DicomInstanceBindingSpec defines the desired state of DicomInstanceBinding",
				MarkdownDescription: "DicomInstanceBindingSpec defines the desired state of DicomInstanceBinding",
				Attributes: map[string]schema.Attribute{
					"binding_config_name": schema.StringAttribute{
						Description:         "Binding Config Name",
						MarkdownDescription: "Binding Config Name",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"binding_secret_name": schema.StringAttribute{
						Description:         "Binding Secret Name",
						MarkdownDescription: "Binding Secret Name",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"dicom_event_driven_ingestion_name": schema.StringAttribute{
						Description:         "DICOM Event Driven Ingestion Name",
						MarkdownDescription: "DICOM Event Driven Ingestion Name",
						Required:            false,
						Optional:            false,
						Computed:            true,
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

					"instance_binding": schema.SingleNestedAttribute{
						Description:         "Instance Binding Spec",
						MarkdownDescription: "Instance Binding Spec",
						Attributes: map[string]schema.Attribute{
							"concurrency": schema.Int64Attribute{
								Description:         "Container Concurrency",
								MarkdownDescription: "Container Concurrency",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"image": schema.StringAttribute{
								Description:         "Image",
								MarkdownDescription: "Image",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"max_replicas": schema.Int64Attribute{
								Description:         "Max Replicas",
								MarkdownDescription: "Max Replicas",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"min_replicas": schema.Int64Attribute{
								Description:         "Min Replicas",
								MarkdownDescription: "Min Replicas",
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
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *ImagingIngestionAlvearieOrgDicomInstanceBindingV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *ImagingIngestionAlvearieOrgDicomInstanceBindingV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_imaging_ingestion_alvearie_org_dicom_instance_binding_v1alpha1")

	var data ImagingIngestionAlvearieOrgDicomInstanceBindingV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "imaging-ingestion.alvearie.org", Version: "v1alpha1", Resource: "dicominstancebindings"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse ImagingIngestionAlvearieOrgDicomInstanceBindingV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("imaging-ingestion.alvearie.org/v1alpha1")
	data.Kind = pointer.String("DicomInstanceBinding")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
