/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package operator_marin3r_3scale_net_v1alpha1

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
	_ datasource.DataSource = &OperatorMarin3R3ScaleNetDiscoveryServiceCertificateV1Alpha1Manifest{}
)

func NewOperatorMarin3R3ScaleNetDiscoveryServiceCertificateV1Alpha1Manifest() datasource.DataSource {
	return &OperatorMarin3R3ScaleNetDiscoveryServiceCertificateV1Alpha1Manifest{}
}

type OperatorMarin3R3ScaleNetDiscoveryServiceCertificateV1Alpha1Manifest struct{}

type OperatorMarin3R3ScaleNetDiscoveryServiceCertificateV1Alpha1ManifestData struct {
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
		CertificateRenewal *struct {
			Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
		} `tfsdk:"certificate_renewal" json:"certificateRenewal,omitempty"`
		CommonName *string   `tfsdk:"common_name" json:"commonName,omitempty"`
		Hosts      *[]string `tfsdk:"hosts" json:"hosts,omitempty"`
		IsCA       *bool     `tfsdk:"is_ca" json:"isCA,omitempty"`
		SecretRef  *struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
		Server *bool `tfsdk:"server" json:"server,omitempty"`
		Signer *struct {
			CaSigned *struct {
				CaSecretRef *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"ca_secret_ref" json:"caSecretRef,omitempty"`
			} `tfsdk:"ca_signed" json:"caSigned,omitempty"`
			SelfSigned *map[string]string `tfsdk:"self_signed" json:"selfSigned,omitempty"`
		} `tfsdk:"signer" json:"signer,omitempty"`
		ValidFor *int64 `tfsdk:"valid_for" json:"validFor,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OperatorMarin3R3ScaleNetDiscoveryServiceCertificateV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operator_marin3r_3scale_net_discovery_service_certificate_v1alpha1_manifest"
}

func (r *OperatorMarin3R3ScaleNetDiscoveryServiceCertificateV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DiscoveryServiceCertificate is an internal resource used to create certificates. This resource is used by the DiscoveryService controller to create the required certificates for the different components. Direct use of DiscoveryServiceCertificate objects is discouraged.",
		MarkdownDescription: "DiscoveryServiceCertificate is an internal resource used to create certificates. This resource is used by the DiscoveryService controller to create the required certificates for the different components. Direct use of DiscoveryServiceCertificate objects is discouraged.",
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
				Description:         "DiscoveryServiceCertificateSpec defines the desired state of DiscoveryServiceCertificate",
				MarkdownDescription: "DiscoveryServiceCertificateSpec defines the desired state of DiscoveryServiceCertificate",
				Attributes: map[string]schema.Attribute{
					"certificate_renewal": schema.SingleNestedAttribute{
						Description:         "CertificateRenewalConfig configures the certificate renewal process. If unset default behavior is to renew the certificate but not notify of renewals.",
						MarkdownDescription: "CertificateRenewalConfig configures the certificate renewal process. If unset default behavior is to renew the certificate but not notify of renewals.",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Enabled is a flag to enable or disable renewal of the certificate",
								MarkdownDescription: "Enabled is a flag to enable or disable renewal of the certificate",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"common_name": schema.StringAttribute{
						Description:         "CommonName is the CommonName of the certificate",
						MarkdownDescription: "CommonName is the CommonName of the certificate",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"hosts": schema.ListAttribute{
						Description:         "Hosts is the list of hosts the certificate is valid for. Only use when 'IsServerCertificate' is true. If unset, the CommonName field will be used to populate the valid hosts of the certificate.",
						MarkdownDescription: "Hosts is the list of hosts the certificate is valid for. Only use when 'IsServerCertificate' is true. If unset, the CommonName field will be used to populate the valid hosts of the certificate.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"is_ca": schema.BoolAttribute{
						Description:         "IsCA is a boolean specifying that the certificate is a CA",
						MarkdownDescription: "IsCA is a boolean specifying that the certificate is a CA",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"secret_ref": schema.SingleNestedAttribute{
						Description:         "SecretRef is a reference to the secret that will hold the certificate and the private key.",
						MarkdownDescription: "SecretRef is a reference to the secret that will hold the certificate and the private key.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "name is unique within a namespace to reference a secret resource.",
								MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "namespace defines the space within which the secret name must be unique.",
								MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"server": schema.BoolAttribute{
						Description:         "IsServerCertificate is a boolean specifying if the certificate should be issued with server auth usage enabled",
						MarkdownDescription: "IsServerCertificate is a boolean specifying if the certificate should be issued with server auth usage enabled",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"signer": schema.SingleNestedAttribute{
						Description:         "Signer specifies  the signer to use to create this certificate. Supported signers are CertManager and SelfSigned.",
						MarkdownDescription: "Signer specifies  the signer to use to create this certificate. Supported signers are CertManager and SelfSigned.",
						Attributes: map[string]schema.Attribute{
							"ca_signed": schema.SingleNestedAttribute{
								Description:         "CASigned holds specific configuration for the CASigned signer",
								MarkdownDescription: "CASigned holds specific configuration for the CASigned signer",
								Attributes: map[string]schema.Attribute{
									"ca_secret_ref": schema.SingleNestedAttribute{
										Description:         "A reference to a Secret containing the CA",
										MarkdownDescription: "A reference to a Secret containing the CA",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "name is unique within a namespace to reference a secret resource.",
												MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "namespace defines the space within which the secret name must be unique.",
												MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"self_signed": schema.MapAttribute{
								Description:         "SelfSigned holds specific configuration for the SelfSigned signer",
								MarkdownDescription: "SelfSigned holds specific configuration for the SelfSigned signer",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"valid_for": schema.Int64Attribute{
						Description:         "ValidFor specifies the validity of the certificate in seconds",
						MarkdownDescription: "ValidFor specifies the validity of the certificate in seconds",
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

func (r *OperatorMarin3R3ScaleNetDiscoveryServiceCertificateV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_operator_marin3r_3scale_net_discovery_service_certificate_v1alpha1_manifest")

	var model OperatorMarin3R3ScaleNetDiscoveryServiceCertificateV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("operator.marin3r.3scale.net/v1alpha1")
	model.Kind = pointer.String("DiscoveryServiceCertificate")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
