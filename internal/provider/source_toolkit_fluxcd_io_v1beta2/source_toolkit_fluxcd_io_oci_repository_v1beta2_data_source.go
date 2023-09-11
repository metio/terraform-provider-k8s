/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package source_toolkit_fluxcd_io_v1beta2

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"net/http"
)

var (
	_ datasource.DataSource              = &SourceToolkitFluxcdIoOcirepositoryV1Beta2DataSource{}
	_ datasource.DataSourceWithConfigure = &SourceToolkitFluxcdIoOcirepositoryV1Beta2DataSource{}
)

func NewSourceToolkitFluxcdIoOcirepositoryV1Beta2DataSource() datasource.DataSource {
	return &SourceToolkitFluxcdIoOcirepositoryV1Beta2DataSource{}
}

type SourceToolkitFluxcdIoOcirepositoryV1Beta2DataSource struct {
	kubernetesClient dynamic.Interface
}

type SourceToolkitFluxcdIoOcirepositoryV1Beta2DataSourceData struct {
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
		CertSecretRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"cert_secret_ref" json:"certSecretRef,omitempty"`
		Ignore        *string `tfsdk:"ignore" json:"ignore,omitempty"`
		Insecure      *bool   `tfsdk:"insecure" json:"insecure,omitempty"`
		Interval      *string `tfsdk:"interval" json:"interval,omitempty"`
		LayerSelector *struct {
			MediaType *string `tfsdk:"media_type" json:"mediaType,omitempty"`
			Operation *string `tfsdk:"operation" json:"operation,omitempty"`
		} `tfsdk:"layer_selector" json:"layerSelector,omitempty"`
		Provider *string `tfsdk:"provider" json:"provider,omitempty"`
		Ref      *struct {
			Digest *string `tfsdk:"digest" json:"digest,omitempty"`
			Semver *string `tfsdk:"semver" json:"semver,omitempty"`
			Tag    *string `tfsdk:"tag" json:"tag,omitempty"`
		} `tfsdk:"ref" json:"ref,omitempty"`
		SecretRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
		ServiceAccountName *string `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
		Suspend            *bool   `tfsdk:"suspend" json:"suspend,omitempty"`
		Timeout            *string `tfsdk:"timeout" json:"timeout,omitempty"`
		Url                *string `tfsdk:"url" json:"url,omitempty"`
		Verify             *struct {
			Provider  *string `tfsdk:"provider" json:"provider,omitempty"`
			SecretRef *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
		} `tfsdk:"verify" json:"verify,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SourceToolkitFluxcdIoOcirepositoryV1Beta2DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_source_toolkit_fluxcd_io_oci_repository_v1beta2"
}

func (r *SourceToolkitFluxcdIoOcirepositoryV1Beta2DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "OCIRepository is the Schema for the ocirepositories API",
		MarkdownDescription: "OCIRepository is the Schema for the ocirepositories API",
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
				Description:         "OCIRepositorySpec defines the desired state of OCIRepository",
				MarkdownDescription: "OCIRepositorySpec defines the desired state of OCIRepository",
				Attributes: map[string]schema.Attribute{
					"cert_secret_ref": schema.SingleNestedAttribute{
						Description:         "CertSecretRef can be given the name of a Secret containing either or both of  - a PEM-encoded client certificate ('tls.crt') and private key ('tls.key'); - a PEM-encoded CA certificate ('ca.crt')  and whichever are supplied, will be used for connecting to the registry. The client cert and key are useful if you are authenticating with a certificate; the CA cert is useful if you are using a self-signed server certificate. The Secret must be of type 'Opaque' or 'kubernetes.io/tls'.  Note: Support for the 'caFile', 'certFile' and 'keyFile' keys have been deprecated.",
						MarkdownDescription: "CertSecretRef can be given the name of a Secret containing either or both of  - a PEM-encoded client certificate ('tls.crt') and private key ('tls.key'); - a PEM-encoded CA certificate ('ca.crt')  and whichever are supplied, will be used for connecting to the registry. The client cert and key are useful if you are authenticating with a certificate; the CA cert is useful if you are using a self-signed server certificate. The Secret must be of type 'Opaque' or 'kubernetes.io/tls'.  Note: Support for the 'caFile', 'certFile' and 'keyFile' keys have been deprecated.",
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

					"ignore": schema.StringAttribute{
						Description:         "Ignore overrides the set of excluded patterns in the .sourceignore format (which is the same as .gitignore). If not provided, a default will be used, consult the documentation for your version to find out what those are.",
						MarkdownDescription: "Ignore overrides the set of excluded patterns in the .sourceignore format (which is the same as .gitignore). If not provided, a default will be used, consult the documentation for your version to find out what those are.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"insecure": schema.BoolAttribute{
						Description:         "Insecure allows connecting to a non-TLS HTTP container registry.",
						MarkdownDescription: "Insecure allows connecting to a non-TLS HTTP container registry.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"interval": schema.StringAttribute{
						Description:         "Interval at which the OCIRepository URL is checked for updates. This interval is approximate and may be subject to jitter to ensure efficient use of resources.",
						MarkdownDescription: "Interval at which the OCIRepository URL is checked for updates. This interval is approximate and may be subject to jitter to ensure efficient use of resources.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"layer_selector": schema.SingleNestedAttribute{
						Description:         "LayerSelector specifies which layer should be extracted from the OCI artifact. When not specified, the first layer found in the artifact is selected.",
						MarkdownDescription: "LayerSelector specifies which layer should be extracted from the OCI artifact. When not specified, the first layer found in the artifact is selected.",
						Attributes: map[string]schema.Attribute{
							"media_type": schema.StringAttribute{
								Description:         "MediaType specifies the OCI media type of the layer which should be extracted from the OCI Artifact. The first layer matching this type is selected.",
								MarkdownDescription: "MediaType specifies the OCI media type of the layer which should be extracted from the OCI Artifact. The first layer matching this type is selected.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"operation": schema.StringAttribute{
								Description:         "Operation specifies how the selected layer should be processed. By default, the layer compressed content is extracted to storage. When the operation is set to 'copy', the layer compressed content is persisted to storage as it is.",
								MarkdownDescription: "Operation specifies how the selected layer should be processed. By default, the layer compressed content is extracted to storage. When the operation is set to 'copy', the layer compressed content is persisted to storage as it is.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"provider": schema.StringAttribute{
						Description:         "The provider used for authentication, can be 'aws', 'azure', 'gcp' or 'generic'. When not specified, defaults to 'generic'.",
						MarkdownDescription: "The provider used for authentication, can be 'aws', 'azure', 'gcp' or 'generic'. When not specified, defaults to 'generic'.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"ref": schema.SingleNestedAttribute{
						Description:         "The OCI reference to pull and monitor for changes, defaults to the latest tag.",
						MarkdownDescription: "The OCI reference to pull and monitor for changes, defaults to the latest tag.",
						Attributes: map[string]schema.Attribute{
							"digest": schema.StringAttribute{
								Description:         "Digest is the image digest to pull, takes precedence over SemVer. The value should be in the format 'sha256:<HASH>'.",
								MarkdownDescription: "Digest is the image digest to pull, takes precedence over SemVer. The value should be in the format 'sha256:<HASH>'.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"semver": schema.StringAttribute{
								Description:         "SemVer is the range of tags to pull selecting the latest within the range, takes precedence over Tag.",
								MarkdownDescription: "SemVer is the range of tags to pull selecting the latest within the range, takes precedence over Tag.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"tag": schema.StringAttribute{
								Description:         "Tag is the image tag to pull, defaults to latest.",
								MarkdownDescription: "Tag is the image tag to pull, defaults to latest.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"secret_ref": schema.SingleNestedAttribute{
						Description:         "SecretRef contains the secret name containing the registry login credentials to resolve image metadata. The secret must be of type kubernetes.io/dockerconfigjson.",
						MarkdownDescription: "SecretRef contains the secret name containing the registry login credentials to resolve image metadata. The secret must be of type kubernetes.io/dockerconfigjson.",
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

					"service_account_name": schema.StringAttribute{
						Description:         "ServiceAccountName is the name of the Kubernetes ServiceAccount used to authenticate the image pull if the service account has attached pull secrets. For more information: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#add-imagepullsecrets-to-a-service-account",
						MarkdownDescription: "ServiceAccountName is the name of the Kubernetes ServiceAccount used to authenticate the image pull if the service account has attached pull secrets. For more information: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#add-imagepullsecrets-to-a-service-account",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"suspend": schema.BoolAttribute{
						Description:         "This flag tells the controller to suspend the reconciliation of this source.",
						MarkdownDescription: "This flag tells the controller to suspend the reconciliation of this source.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"timeout": schema.StringAttribute{
						Description:         "The timeout for remote OCI Repository operations like pulling, defaults to 60s.",
						MarkdownDescription: "The timeout for remote OCI Repository operations like pulling, defaults to 60s.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"url": schema.StringAttribute{
						Description:         "URL is a reference to an OCI artifact repository hosted on a remote container registry.",
						MarkdownDescription: "URL is a reference to an OCI artifact repository hosted on a remote container registry.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"verify": schema.SingleNestedAttribute{
						Description:         "Verify contains the secret name containing the trusted public keys used to verify the signature and specifies which provider to use to check whether OCI image is authentic.",
						MarkdownDescription: "Verify contains the secret name containing the trusted public keys used to verify the signature and specifies which provider to use to check whether OCI image is authentic.",
						Attributes: map[string]schema.Attribute{
							"provider": schema.StringAttribute{
								Description:         "Provider specifies the technology used to sign the OCI Artifact.",
								MarkdownDescription: "Provider specifies the technology used to sign the OCI Artifact.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"secret_ref": schema.SingleNestedAttribute{
								Description:         "SecretRef specifies the Kubernetes Secret containing the trusted public keys.",
								MarkdownDescription: "SecretRef specifies the Kubernetes Secret containing the trusted public keys.",
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

func (r *SourceToolkitFluxcdIoOcirepositoryV1Beta2DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *SourceToolkitFluxcdIoOcirepositoryV1Beta2DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_source_toolkit_fluxcd_io_oci_repository_v1beta2")

	var data SourceToolkitFluxcdIoOcirepositoryV1Beta2DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "source.toolkit.fluxcd.io", Version: "v1beta2", Resource: "ocirepositories"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		var statusError *k8sErrors.StatusError
		if errors.As(err, &statusError) {
			if statusError.Status().Code == http.StatusNotFound {
				response.Diagnostics.AddError(
					"Unable to find resource",
					fmt.Sprintf("The requested resource cannot be found. "+
						"Make sure that it does exist in your cluster and you have set the correct name and namespace configured.\n\n"+
						"Namespace: %s\n"+
						"Name: %s", data.Metadata.Namespace, data.Metadata.Name),
				)
				return
			}
		} else {
			response.Diagnostics.AddError(
				"Unable to GET resource",
				fmt.Sprintf("An unexpected error occurred while reading the resource. "+
					"Please report this issue to the provider developers.\n\n"+
					"GET Error (%T): %s", err, err.Error()),
			)
		}
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

	var readResponse SourceToolkitFluxcdIoOcirepositoryV1Beta2DataSourceData
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
	data.ApiVersion = pointer.String("source.toolkit.fluxcd.io/v1beta2")
	data.Kind = pointer.String("OCIRepository")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
