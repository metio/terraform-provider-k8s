/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package source_toolkit_fluxcd_io_v1beta2

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &SourceToolkitFluxcdIoOcirepositoryV1Beta2Manifest{}
)

func NewSourceToolkitFluxcdIoOcirepositoryV1Beta2Manifest() datasource.DataSource {
	return &SourceToolkitFluxcdIoOcirepositoryV1Beta2Manifest{}
}

type SourceToolkitFluxcdIoOcirepositoryV1Beta2Manifest struct{}

type SourceToolkitFluxcdIoOcirepositoryV1Beta2ManifestData struct {
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
			MatchOIDCIdentity *[]struct {
				Issuer  *string `tfsdk:"issuer" json:"issuer,omitempty"`
				Subject *string `tfsdk:"subject" json:"subject,omitempty"`
			} `tfsdk:"match_oidc_identity" json:"matchOIDCIdentity,omitempty"`
			Provider  *string `tfsdk:"provider" json:"provider,omitempty"`
			SecretRef *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
		} `tfsdk:"verify" json:"verify,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SourceToolkitFluxcdIoOcirepositoryV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_source_toolkit_fluxcd_io_oci_repository_v1beta2_manifest"
}

func (r *SourceToolkitFluxcdIoOcirepositoryV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "OCIRepository is the Schema for the ocirepositories API",
		MarkdownDescription: "OCIRepository is the Schema for the ocirepositories API",
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
				Description:         "OCIRepositorySpec defines the desired state of OCIRepository",
				MarkdownDescription: "OCIRepositorySpec defines the desired state of OCIRepository",
				Attributes: map[string]schema.Attribute{
					"cert_secret_ref": schema.SingleNestedAttribute{
						Description:         "CertSecretRef can be given the name of a Secret containingeither or both of- a PEM-encoded client certificate ('tls.crt') and privatekey ('tls.key');- a PEM-encoded CA certificate ('ca.crt')and whichever are supplied, will be used for connecting to theregistry. The client cert and key are useful if you areauthenticating with a certificate; the CA cert is useful ifyou are using a self-signed server certificate. The Secret mustbe of type 'Opaque' or 'kubernetes.io/tls'.Note: Support for the 'caFile', 'certFile' and 'keyFile' keys havebeen deprecated.",
						MarkdownDescription: "CertSecretRef can be given the name of a Secret containingeither or both of- a PEM-encoded client certificate ('tls.crt') and privatekey ('tls.key');- a PEM-encoded CA certificate ('ca.crt')and whichever are supplied, will be used for connecting to theregistry. The client cert and key are useful if you areauthenticating with a certificate; the CA cert is useful ifyou are using a self-signed server certificate. The Secret mustbe of type 'Opaque' or 'kubernetes.io/tls'.Note: Support for the 'caFile', 'certFile' and 'keyFile' keys havebeen deprecated.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the referent.",
								MarkdownDescription: "Name of the referent.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"ignore": schema.StringAttribute{
						Description:         "Ignore overrides the set of excluded patterns in the .sourceignore format(which is the same as .gitignore). If not provided, a default will be used,consult the documentation for your version to find out what those are.",
						MarkdownDescription: "Ignore overrides the set of excluded patterns in the .sourceignore format(which is the same as .gitignore). If not provided, a default will be used,consult the documentation for your version to find out what those are.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"insecure": schema.BoolAttribute{
						Description:         "Insecure allows connecting to a non-TLS HTTP container registry.",
						MarkdownDescription: "Insecure allows connecting to a non-TLS HTTP container registry.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"interval": schema.StringAttribute{
						Description:         "Interval at which the OCIRepository URL is checked for updates.This interval is approximate and may be subject to jitter to ensureefficient use of resources.",
						MarkdownDescription: "Interval at which the OCIRepository URL is checked for updates.This interval is approximate and may be subject to jitter to ensureefficient use of resources.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ms|s|m|h))+$`), ""),
						},
					},

					"layer_selector": schema.SingleNestedAttribute{
						Description:         "LayerSelector specifies which layer should be extracted from the OCI artifact.When not specified, the first layer found in the artifact is selected.",
						MarkdownDescription: "LayerSelector specifies which layer should be extracted from the OCI artifact.When not specified, the first layer found in the artifact is selected.",
						Attributes: map[string]schema.Attribute{
							"media_type": schema.StringAttribute{
								Description:         "MediaType specifies the OCI media type of the layerwhich should be extracted from the OCI Artifact. Thefirst layer matching this type is selected.",
								MarkdownDescription: "MediaType specifies the OCI media type of the layerwhich should be extracted from the OCI Artifact. Thefirst layer matching this type is selected.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"operation": schema.StringAttribute{
								Description:         "Operation specifies how the selected layer should be processed.By default, the layer compressed content is extracted to storage.When the operation is set to 'copy', the layer compressed contentis persisted to storage as it is.",
								MarkdownDescription: "Operation specifies how the selected layer should be processed.By default, the layer compressed content is extracted to storage.When the operation is set to 'copy', the layer compressed contentis persisted to storage as it is.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("extract", "copy"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"provider": schema.StringAttribute{
						Description:         "The provider used for authentication, can be 'aws', 'azure', 'gcp' or 'generic'.When not specified, defaults to 'generic'.",
						MarkdownDescription: "The provider used for authentication, can be 'aws', 'azure', 'gcp' or 'generic'.When not specified, defaults to 'generic'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("generic", "aws", "azure", "gcp"),
						},
					},

					"ref": schema.SingleNestedAttribute{
						Description:         "The OCI reference to pull and monitor for changes,defaults to the latest tag.",
						MarkdownDescription: "The OCI reference to pull and monitor for changes,defaults to the latest tag.",
						Attributes: map[string]schema.Attribute{
							"digest": schema.StringAttribute{
								Description:         "Digest is the image digest to pull, takes precedence over SemVer.The value should be in the format 'sha256:<HASH>'.",
								MarkdownDescription: "Digest is the image digest to pull, takes precedence over SemVer.The value should be in the format 'sha256:<HASH>'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"semver": schema.StringAttribute{
								Description:         "SemVer is the range of tags to pull selecting the latest withinthe range, takes precedence over Tag.",
								MarkdownDescription: "SemVer is the range of tags to pull selecting the latest withinthe range, takes precedence over Tag.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tag": schema.StringAttribute{
								Description:         "Tag is the image tag to pull, defaults to latest.",
								MarkdownDescription: "Tag is the image tag to pull, defaults to latest.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"secret_ref": schema.SingleNestedAttribute{
						Description:         "SecretRef contains the secret name containing the registry logincredentials to resolve image metadata.The secret must be of type kubernetes.io/dockerconfigjson.",
						MarkdownDescription: "SecretRef contains the secret name containing the registry logincredentials to resolve image metadata.The secret must be of type kubernetes.io/dockerconfigjson.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the referent.",
								MarkdownDescription: "Name of the referent.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"service_account_name": schema.StringAttribute{
						Description:         "ServiceAccountName is the name of the Kubernetes ServiceAccount used to authenticatethe image pull if the service account has attached pull secrets. For more information:https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#add-imagepullsecrets-to-a-service-account",
						MarkdownDescription: "ServiceAccountName is the name of the Kubernetes ServiceAccount used to authenticatethe image pull if the service account has attached pull secrets. For more information:https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#add-imagepullsecrets-to-a-service-account",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"suspend": schema.BoolAttribute{
						Description:         "This flag tells the controller to suspend the reconciliation of this source.",
						MarkdownDescription: "This flag tells the controller to suspend the reconciliation of this source.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"timeout": schema.StringAttribute{
						Description:         "The timeout for remote OCI Repository operations like pulling, defaults to 60s.",
						MarkdownDescription: "The timeout for remote OCI Repository operations like pulling, defaults to 60s.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ms|s|m))+$`), ""),
						},
					},

					"url": schema.StringAttribute{
						Description:         "URL is a reference to an OCI artifact repository hostedon a remote container registry.",
						MarkdownDescription: "URL is a reference to an OCI artifact repository hostedon a remote container registry.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^oci://.*$`), ""),
						},
					},

					"verify": schema.SingleNestedAttribute{
						Description:         "Verify contains the secret name containing the trusted public keysused to verify the signature and specifies which provider to use to checkwhether OCI image is authentic.",
						MarkdownDescription: "Verify contains the secret name containing the trusted public keysused to verify the signature and specifies which provider to use to checkwhether OCI image is authentic.",
						Attributes: map[string]schema.Attribute{
							"match_oidc_identity": schema.ListNestedAttribute{
								Description:         "MatchOIDCIdentity specifies the identity matching criteria to usewhile verifying an OCI artifact which was signed using Cosign keylesssigning. The artifact's identity is deemed to be verified if any of thespecified matchers match against the identity.",
								MarkdownDescription: "MatchOIDCIdentity specifies the identity matching criteria to usewhile verifying an OCI artifact which was signed using Cosign keylesssigning. The artifact's identity is deemed to be verified if any of thespecified matchers match against the identity.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"issuer": schema.StringAttribute{
											Description:         "Issuer specifies the regex pattern to match against to verifythe OIDC issuer in the Fulcio certificate. The pattern must be avalid Go regular expression.",
											MarkdownDescription: "Issuer specifies the regex pattern to match against to verifythe OIDC issuer in the Fulcio certificate. The pattern must be avalid Go regular expression.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"subject": schema.StringAttribute{
											Description:         "Subject specifies the regex pattern to match against to verifythe identity subject in the Fulcio certificate. The pattern mustbe a valid Go regular expression.",
											MarkdownDescription: "Subject specifies the regex pattern to match against to verifythe identity subject in the Fulcio certificate. The pattern mustbe a valid Go regular expression.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"provider": schema.StringAttribute{
								Description:         "Provider specifies the technology used to sign the OCI Artifact.",
								MarkdownDescription: "Provider specifies the technology used to sign the OCI Artifact.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("cosign"),
								},
							},

							"secret_ref": schema.SingleNestedAttribute{
								Description:         "SecretRef specifies the Kubernetes Secret containing thetrusted public keys.",
								MarkdownDescription: "SecretRef specifies the Kubernetes Secret containing thetrusted public keys.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the referent.",
										MarkdownDescription: "Name of the referent.",
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

func (r *SourceToolkitFluxcdIoOcirepositoryV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_source_toolkit_fluxcd_io_oci_repository_v1beta2_manifest")

	var model SourceToolkitFluxcdIoOcirepositoryV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("source.toolkit.fluxcd.io/v1beta2")
	model.Kind = pointer.String("OCIRepository")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
