/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

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

type NetworkingIstioIoGatewayV1Alpha3Resource struct{}

var (
	_ resource.Resource = (*NetworkingIstioIoGatewayV1Alpha3Resource)(nil)
)

type NetworkingIstioIoGatewayV1Alpha3TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type NetworkingIstioIoGatewayV1Alpha3GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Selector *map[string]string `tfsdk:"selector" yaml:"selector,omitempty"`

		Servers *[]struct {
			Bind *string `tfsdk:"bind" yaml:"bind,omitempty"`

			DefaultEndpoint *string `tfsdk:"default_endpoint" yaml:"defaultEndpoint,omitempty"`

			Hosts *[]string `tfsdk:"hosts" yaml:"hosts,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Port *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Number *int64 `tfsdk:"number" yaml:"number,omitempty"`

				Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`

				TargetPort *int64 `tfsdk:"target_port" yaml:"targetPort,omitempty"`
			} `tfsdk:"port" yaml:"port,omitempty"`

			Tls *struct {
				CaCertificates *string `tfsdk:"ca_certificates" yaml:"caCertificates,omitempty"`

				CipherSuites *[]string `tfsdk:"cipher_suites" yaml:"cipherSuites,omitempty"`

				CredentialName *string `tfsdk:"credential_name" yaml:"credentialName,omitempty"`

				HttpsRedirect *bool `tfsdk:"https_redirect" yaml:"httpsRedirect,omitempty"`

				MaxProtocolVersion *string `tfsdk:"max_protocol_version" yaml:"maxProtocolVersion,omitempty"`

				MinProtocolVersion *string `tfsdk:"min_protocol_version" yaml:"minProtocolVersion,omitempty"`

				Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

				PrivateKey *string `tfsdk:"private_key" yaml:"privateKey,omitempty"`

				ServerCertificate *string `tfsdk:"server_certificate" yaml:"serverCertificate,omitempty"`

				SubjectAltNames *[]string `tfsdk:"subject_alt_names" yaml:"subjectAltNames,omitempty"`

				VerifyCertificateHash *[]string `tfsdk:"verify_certificate_hash" yaml:"verifyCertificateHash,omitempty"`

				VerifyCertificateSpki *[]string `tfsdk:"verify_certificate_spki" yaml:"verifyCertificateSpki,omitempty"`
			} `tfsdk:"tls" yaml:"tls,omitempty"`
		} `tfsdk:"servers" yaml:"servers,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewNetworkingIstioIoGatewayV1Alpha3Resource() resource.Resource {
	return &NetworkingIstioIoGatewayV1Alpha3Resource{}
}

func (r *NetworkingIstioIoGatewayV1Alpha3Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networking_istio_io_gateway_v1alpha3"
}

func (r *NetworkingIstioIoGatewayV1Alpha3Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "Configuration affecting edge load balancer. See more details at: https://istio.io/docs/reference/config/networking/gateway.html",
				MarkdownDescription: "Configuration affecting edge load balancer. See more details at: https://istio.io/docs/reference/config/networking/gateway.html",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"selector": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"servers": {
						Description:         "A list of server specifications.",
						MarkdownDescription: "A list of server specifications.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"bind": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"default_endpoint": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"hosts": {
								Description:         "One or more hosts exposed by this gateway.",
								MarkdownDescription: "One or more hosts exposed by this gateway.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "An optional name of the server, when set must be unique across all servers.",
								MarkdownDescription: "An optional name of the server, when set must be unique across all servers.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Label assigned to the port.",
										MarkdownDescription: "Label assigned to the port.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"number": {
										Description:         "A valid non-negative integer port number.",
										MarkdownDescription: "A valid non-negative integer port number.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"protocol": {
										Description:         "The protocol exposed on the port.",
										MarkdownDescription: "The protocol exposed on the port.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"target_port": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tls": {
								Description:         "Set of TLS related options that govern the server's behavior.",
								MarkdownDescription: "Set of TLS related options that govern the server's behavior.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"ca_certificates": {
										Description:         "REQUIRED if mode is 'MUTUAL'.",
										MarkdownDescription: "REQUIRED if mode is 'MUTUAL'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"cipher_suites": {
										Description:         "Optional: If specified, only support the specified cipher list.",
										MarkdownDescription: "Optional: If specified, only support the specified cipher list.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"credential_name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"https_redirect": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"max_protocol_version": {
										Description:         "Optional: Maximum TLS protocol version.",
										MarkdownDescription: "Optional: Maximum TLS protocol version.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"min_protocol_version": {
										Description:         "Optional: Minimum TLS protocol version.",
										MarkdownDescription: "Optional: Minimum TLS protocol version.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"mode": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"private_key": {
										Description:         "REQUIRED if mode is 'SIMPLE' or 'MUTUAL'.",
										MarkdownDescription: "REQUIRED if mode is 'SIMPLE' or 'MUTUAL'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"server_certificate": {
										Description:         "REQUIRED if mode is 'SIMPLE' or 'MUTUAL'.",
										MarkdownDescription: "REQUIRED if mode is 'SIMPLE' or 'MUTUAL'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"subject_alt_names": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"verify_certificate_hash": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"verify_certificate_spki": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

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
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *NetworkingIstioIoGatewayV1Alpha3Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_networking_istio_io_gateway_v1alpha3")

	var state NetworkingIstioIoGatewayV1Alpha3TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel NetworkingIstioIoGatewayV1Alpha3GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("networking.istio.io/v1alpha3")
	goModel.Kind = utilities.Ptr("Gateway")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *NetworkingIstioIoGatewayV1Alpha3Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_networking_istio_io_gateway_v1alpha3")
	// NO-OP: All data is already in Terraform state
}

func (r *NetworkingIstioIoGatewayV1Alpha3Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_networking_istio_io_gateway_v1alpha3")

	var state NetworkingIstioIoGatewayV1Alpha3TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel NetworkingIstioIoGatewayV1Alpha3GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("networking.istio.io/v1alpha3")
	goModel.Kind = utilities.Ptr("Gateway")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *NetworkingIstioIoGatewayV1Alpha3Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_networking_istio_io_gateway_v1alpha3")
	// NO-OP: Terraform removes the state automatically for us
}
