/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type SourceToolkitFluxcdIoOCIRepositoryV1Beta2Resource struct{}

var (
	_ resource.Resource = (*SourceToolkitFluxcdIoOCIRepositoryV1Beta2Resource)(nil)
)

type SourceToolkitFluxcdIoOCIRepositoryV1Beta2TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type SourceToolkitFluxcdIoOCIRepositoryV1Beta2GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		CertSecretRef *struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"cert_secret_ref" yaml:"certSecretRef,omitempty"`

		Ignore *string `tfsdk:"ignore" yaml:"ignore,omitempty"`

		Insecure *bool `tfsdk:"insecure" yaml:"insecure,omitempty"`

		Interval *string `tfsdk:"interval" yaml:"interval,omitempty"`

		LayerSelector *struct {
			MediaType *string `tfsdk:"media_type" yaml:"mediaType,omitempty"`

			Operation *string `tfsdk:"operation" yaml:"operation,omitempty"`
		} `tfsdk:"layer_selector" yaml:"layerSelector,omitempty"`

		Provider *string `tfsdk:"provider" yaml:"provider,omitempty"`

		Ref *struct {
			Digest *string `tfsdk:"digest" yaml:"digest,omitempty"`

			Semver *string `tfsdk:"semver" yaml:"semver,omitempty"`

			Tag *string `tfsdk:"tag" yaml:"tag,omitempty"`
		} `tfsdk:"ref" yaml:"ref,omitempty"`

		SecretRef *struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

		ServiceAccountName *string `tfsdk:"service_account_name" yaml:"serviceAccountName,omitempty"`

		Suspend *bool `tfsdk:"suspend" yaml:"suspend,omitempty"`

		Timeout *string `tfsdk:"timeout" yaml:"timeout,omitempty"`

		Url *string `tfsdk:"url" yaml:"url,omitempty"`

		Verify *struct {
			Provider *string `tfsdk:"provider" yaml:"provider,omitempty"`

			SecretRef *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`
		} `tfsdk:"verify" yaml:"verify,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewSourceToolkitFluxcdIoOCIRepositoryV1Beta2Resource() resource.Resource {
	return &SourceToolkitFluxcdIoOCIRepositoryV1Beta2Resource{}
}

func (r *SourceToolkitFluxcdIoOCIRepositoryV1Beta2Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_toolkit_fluxcd_io_oci_repository_v1beta2"
}

func (r *SourceToolkitFluxcdIoOCIRepositoryV1Beta2Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "OCIRepository is the Schema for the ocirepositories API",
		MarkdownDescription: "OCIRepository is the Schema for the ocirepositories API",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "OCIRepositorySpec defines the desired state of OCIRepository",
				MarkdownDescription: "OCIRepositorySpec defines the desired state of OCIRepository",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"cert_secret_ref": {
						Description:         "CertSecretRef can be given the name of a secret containing either or both of  - a PEM-encoded client certificate ('certFile') and private key ('keyFile'); - a PEM-encoded CA certificate ('caFile')  and whichever are supplied, will be used for connecting to the registry. The client cert and key are useful if you are authenticating with a certificate; the CA cert is useful if you are using a self-signed server certificate.",
						MarkdownDescription: "CertSecretRef can be given the name of a secret containing either or both of  - a PEM-encoded client certificate ('certFile') and private key ('keyFile'); - a PEM-encoded CA certificate ('caFile')  and whichever are supplied, will be used for connecting to the registry. The client cert and key are useful if you are authenticating with a certificate; the CA cert is useful if you are using a self-signed server certificate.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name of the referent.",
								MarkdownDescription: "Name of the referent.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ignore": {
						Description:         "Ignore overrides the set of excluded patterns in the .sourceignore format (which is the same as .gitignore). If not provided, a default will be used, consult the documentation for your version to find out what those are.",
						MarkdownDescription: "Ignore overrides the set of excluded patterns in the .sourceignore format (which is the same as .gitignore). If not provided, a default will be used, consult the documentation for your version to find out what those are.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"insecure": {
						Description:         "Insecure allows connecting to a non-TLS HTTP container registry.",
						MarkdownDescription: "Insecure allows connecting to a non-TLS HTTP container registry.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"interval": {
						Description:         "The interval at which to check for image updates.",
						MarkdownDescription: "The interval at which to check for image updates.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ms|s|m|h))+$`), ""),
						},
					},

					"layer_selector": {
						Description:         "LayerSelector specifies which layer should be extracted from the OCI artifact. When not specified, the first layer found in the artifact is selected.",
						MarkdownDescription: "LayerSelector specifies which layer should be extracted from the OCI artifact. When not specified, the first layer found in the artifact is selected.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"media_type": {
								Description:         "MediaType specifies the OCI media type of the layer which should be extracted from the OCI Artifact. The first layer matching this type is selected.",
								MarkdownDescription: "MediaType specifies the OCI media type of the layer which should be extracted from the OCI Artifact. The first layer matching this type is selected.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"operation": {
								Description:         "Operation specifies how the selected layer should be processed. By default, the layer compressed content is extracted to storage. When the operation is set to 'copy', the layer compressed content is persisted to storage as it is.",
								MarkdownDescription: "Operation specifies how the selected layer should be processed. By default, the layer compressed content is extracted to storage. When the operation is set to 'copy', the layer compressed content is persisted to storage as it is.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("extract", "copy"),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"provider": {
						Description:         "The provider used for authentication, can be 'aws', 'azure', 'gcp' or 'generic'. When not specified, defaults to 'generic'.",
						MarkdownDescription: "The provider used for authentication, can be 'aws', 'azure', 'gcp' or 'generic'. When not specified, defaults to 'generic'.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("generic", "aws", "azure", "gcp"),
						},
					},

					"ref": {
						Description:         "The OCI reference to pull and monitor for changes, defaults to the latest tag.",
						MarkdownDescription: "The OCI reference to pull and monitor for changes, defaults to the latest tag.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"digest": {
								Description:         "Digest is the image digest to pull, takes precedence over SemVer. The value should be in the format 'sha256:<HASH>'.",
								MarkdownDescription: "Digest is the image digest to pull, takes precedence over SemVer. The value should be in the format 'sha256:<HASH>'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"semver": {
								Description:         "SemVer is the range of tags to pull selecting the latest within the range, takes precedence over Tag.",
								MarkdownDescription: "SemVer is the range of tags to pull selecting the latest within the range, takes precedence over Tag.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tag": {
								Description:         "Tag is the image tag to pull, defaults to latest.",
								MarkdownDescription: "Tag is the image tag to pull, defaults to latest.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"secret_ref": {
						Description:         "SecretRef contains the secret name containing the registry login credentials to resolve image metadata. The secret must be of type kubernetes.io/dockerconfigjson.",
						MarkdownDescription: "SecretRef contains the secret name containing the registry login credentials to resolve image metadata. The secret must be of type kubernetes.io/dockerconfigjson.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name of the referent.",
								MarkdownDescription: "Name of the referent.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"service_account_name": {
						Description:         "ServiceAccountName is the name of the Kubernetes ServiceAccount used to authenticate the image pull if the service account has attached pull secrets. For more information: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#add-imagepullsecrets-to-a-service-account",
						MarkdownDescription: "ServiceAccountName is the name of the Kubernetes ServiceAccount used to authenticate the image pull if the service account has attached pull secrets. For more information: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#add-imagepullsecrets-to-a-service-account",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"suspend": {
						Description:         "This flag tells the controller to suspend the reconciliation of this source.",
						MarkdownDescription: "This flag tells the controller to suspend the reconciliation of this source.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"timeout": {
						Description:         "The timeout for remote OCI Repository operations like pulling, defaults to 60s.",
						MarkdownDescription: "The timeout for remote OCI Repository operations like pulling, defaults to 60s.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ms|s|m))+$`), ""),
						},
					},

					"url": {
						Description:         "URL is a reference to an OCI artifact repository hosted on a remote container registry.",
						MarkdownDescription: "URL is a reference to an OCI artifact repository hosted on a remote container registry.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.RegexMatches(regexp.MustCompile(`^oci://.*$`), ""),
						},
					},

					"verify": {
						Description:         "Verify contains the secret name containing the trusted public keys used to verify the signature and specifies which provider to use to check whether OCI image is authentic.",
						MarkdownDescription: "Verify contains the secret name containing the trusted public keys used to verify the signature and specifies which provider to use to check whether OCI image is authentic.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"provider": {
								Description:         "Provider specifies the technology used to sign the OCI Artifact.",
								MarkdownDescription: "Provider specifies the technology used to sign the OCI Artifact.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("cosign"),
								},
							},

							"secret_ref": {
								Description:         "SecretRef specifies the Kubernetes Secret containing the trusted public keys.",
								MarkdownDescription: "SecretRef specifies the Kubernetes Secret containing the trusted public keys.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name of the referent.",
										MarkdownDescription: "Name of the referent.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *SourceToolkitFluxcdIoOCIRepositoryV1Beta2Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_source_toolkit_fluxcd_io_oci_repository_v1beta2")

	var state SourceToolkitFluxcdIoOCIRepositoryV1Beta2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SourceToolkitFluxcdIoOCIRepositoryV1Beta2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("source.toolkit.fluxcd.io/v1beta2")
	goModel.Kind = utilities.Ptr("OCIRepository")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *SourceToolkitFluxcdIoOCIRepositoryV1Beta2Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_source_toolkit_fluxcd_io_oci_repository_v1beta2")
	// NO-OP: All data is already in Terraform state
}

func (r *SourceToolkitFluxcdIoOCIRepositoryV1Beta2Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_source_toolkit_fluxcd_io_oci_repository_v1beta2")

	var state SourceToolkitFluxcdIoOCIRepositoryV1Beta2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SourceToolkitFluxcdIoOCIRepositoryV1Beta2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("source.toolkit.fluxcd.io/v1beta2")
	goModel.Kind = utilities.Ptr("OCIRepository")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *SourceToolkitFluxcdIoOCIRepositoryV1Beta2Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_source_toolkit_fluxcd_io_oci_repository_v1beta2")
	// NO-OP: Terraform removes the state automatically for us
}
