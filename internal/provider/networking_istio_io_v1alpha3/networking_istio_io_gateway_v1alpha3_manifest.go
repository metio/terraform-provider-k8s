/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package networking_istio_io_v1alpha3

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &NetworkingIstioIoGatewayV1Alpha3Manifest{}
)

func NewNetworkingIstioIoGatewayV1Alpha3Manifest() datasource.DataSource {
	return &NetworkingIstioIoGatewayV1Alpha3Manifest{}
}

type NetworkingIstioIoGatewayV1Alpha3Manifest struct{}

type NetworkingIstioIoGatewayV1Alpha3ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
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
		Selector *map[string]string `tfsdk:"selector" json:"selector,omitempty"`
		Servers  *[]struct {
			Bind            *string   `tfsdk:"bind" json:"bind,omitempty"`
			DefaultEndpoint *string   `tfsdk:"default_endpoint" json:"defaultEndpoint,omitempty"`
			Hosts           *[]string `tfsdk:"hosts" json:"hosts,omitempty"`
			Name            *string   `tfsdk:"name" json:"name,omitempty"`
			Port            *struct {
				Name       *string `tfsdk:"name" json:"name,omitempty"`
				Number     *int64  `tfsdk:"number" json:"number,omitempty"`
				Protocol   *string `tfsdk:"protocol" json:"protocol,omitempty"`
				TargetPort *int64  `tfsdk:"target_port" json:"targetPort,omitempty"`
			} `tfsdk:"port" json:"port,omitempty"`
			Tls *struct {
				CaCertificates        *string   `tfsdk:"ca_certificates" json:"caCertificates,omitempty"`
				CipherSuites          *[]string `tfsdk:"cipher_suites" json:"cipherSuites,omitempty"`
				CredentialName        *string   `tfsdk:"credential_name" json:"credentialName,omitempty"`
				HttpsRedirect         *bool     `tfsdk:"https_redirect" json:"httpsRedirect,omitempty"`
				MaxProtocolVersion    *string   `tfsdk:"max_protocol_version" json:"maxProtocolVersion,omitempty"`
				MinProtocolVersion    *string   `tfsdk:"min_protocol_version" json:"minProtocolVersion,omitempty"`
				Mode                  *string   `tfsdk:"mode" json:"mode,omitempty"`
				PrivateKey            *string   `tfsdk:"private_key" json:"privateKey,omitempty"`
				ServerCertificate     *string   `tfsdk:"server_certificate" json:"serverCertificate,omitempty"`
				SubjectAltNames       *[]string `tfsdk:"subject_alt_names" json:"subjectAltNames,omitempty"`
				VerifyCertificateHash *[]string `tfsdk:"verify_certificate_hash" json:"verifyCertificateHash,omitempty"`
				VerifyCertificateSpki *[]string `tfsdk:"verify_certificate_spki" json:"verifyCertificateSpki,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
		} `tfsdk:"servers" json:"servers,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *NetworkingIstioIoGatewayV1Alpha3Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_networking_istio_io_gateway_v1alpha3_manifest"
}

func (r *NetworkingIstioIoGatewayV1Alpha3Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

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
				Description:         "Configuration affecting edge load balancer. See more details at: https://istio.io/docs/reference/config/networking/gateway.html",
				MarkdownDescription: "Configuration affecting edge load balancer. See more details at: https://istio.io/docs/reference/config/networking/gateway.html",
				Attributes: map[string]schema.Attribute{
					"selector": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"servers": schema.ListNestedAttribute{
						Description:         "A list of server specifications.",
						MarkdownDescription: "A list of server specifications.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"bind": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"default_endpoint": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"hosts": schema.ListAttribute{
									Description:         "One or more hosts exposed by this gateway.",
									MarkdownDescription: "One or more hosts exposed by this gateway.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "An optional name of the server, when set must be unique across all servers.",
									MarkdownDescription: "An optional name of the server, when set must be unique across all servers.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"port": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Label assigned to the port.",
											MarkdownDescription: "Label assigned to the port.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"number": schema.Int64Attribute{
											Description:         "A valid non-negative integer port number.",
											MarkdownDescription: "A valid non-negative integer port number.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"protocol": schema.StringAttribute{
											Description:         "The protocol exposed on the port.",
											MarkdownDescription: "The protocol exposed on the port.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"target_port": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"tls": schema.SingleNestedAttribute{
									Description:         "Set of TLS related options that govern the server's behavior.",
									MarkdownDescription: "Set of TLS related options that govern the server's behavior.",
									Attributes: map[string]schema.Attribute{
										"ca_certificates": schema.StringAttribute{
											Description:         "REQUIRED if mode is 'MUTUAL' or 'OPTIONAL_MUTUAL'.",
											MarkdownDescription: "REQUIRED if mode is 'MUTUAL' or 'OPTIONAL_MUTUAL'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"cipher_suites": schema.ListAttribute{
											Description:         "Optional: If specified, only support the specified cipher list.",
											MarkdownDescription: "Optional: If specified, only support the specified cipher list.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"credential_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"https_redirect": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"max_protocol_version": schema.StringAttribute{
											Description:         "Optional: Maximum TLS protocol version.",
											MarkdownDescription: "Optional: Maximum TLS protocol version.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("TLS_AUTO", "TLSV1_0", "TLSV1_1", "TLSV1_2", "TLSV1_3"),
											},
										},

										"min_protocol_version": schema.StringAttribute{
											Description:         "Optional: Minimum TLS protocol version.",
											MarkdownDescription: "Optional: Minimum TLS protocol version.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("TLS_AUTO", "TLSV1_0", "TLSV1_1", "TLSV1_2", "TLSV1_3"),
											},
										},

										"mode": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("PASSTHROUGH", "SIMPLE", "MUTUAL", "AUTO_PASSTHROUGH", "ISTIO_MUTUAL", "OPTIONAL_MUTUAL"),
											},
										},

										"private_key": schema.StringAttribute{
											Description:         "REQUIRED if mode is 'SIMPLE' or 'MUTUAL'.",
											MarkdownDescription: "REQUIRED if mode is 'SIMPLE' or 'MUTUAL'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"server_certificate": schema.StringAttribute{
											Description:         "REQUIRED if mode is 'SIMPLE' or 'MUTUAL'.",
											MarkdownDescription: "REQUIRED if mode is 'SIMPLE' or 'MUTUAL'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"subject_alt_names": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"verify_certificate_hash": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"verify_certificate_spki": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
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

func (r *NetworkingIstioIoGatewayV1Alpha3Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_networking_istio_io_gateway_v1alpha3_manifest")

	var model NetworkingIstioIoGatewayV1Alpha3ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("networking.istio.io/v1alpha3")
	model.Kind = pointer.String("Gateway")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"YAML Error: "+err.Error(),
		)
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}