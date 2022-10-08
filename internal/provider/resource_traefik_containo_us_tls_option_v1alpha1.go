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

type TraefikContainoUsTLSOptionV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*TraefikContainoUsTLSOptionV1Alpha1Resource)(nil)
)

type TraefikContainoUsTLSOptionV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type TraefikContainoUsTLSOptionV1Alpha1GoModel struct {
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
		CurvePreferences *[]string `tfsdk:"curve_preferences" yaml:"curvePreferences,omitempty"`

		MaxVersion *string `tfsdk:"max_version" yaml:"maxVersion,omitempty"`

		MinVersion *string `tfsdk:"min_version" yaml:"minVersion,omitempty"`

		PreferServerCipherSuites *bool `tfsdk:"prefer_server_cipher_suites" yaml:"preferServerCipherSuites,omitempty"`

		SniStrict *bool `tfsdk:"sni_strict" yaml:"sniStrict,omitempty"`

		AlpnProtocols *[]string `tfsdk:"alpn_protocols" yaml:"alpnProtocols,omitempty"`

		CipherSuites *[]string `tfsdk:"cipher_suites" yaml:"cipherSuites,omitempty"`

		ClientAuth *struct {
			ClientAuthType *string `tfsdk:"client_auth_type" yaml:"clientAuthType,omitempty"`

			SecretNames *[]string `tfsdk:"secret_names" yaml:"secretNames,omitempty"`
		} `tfsdk:"client_auth" yaml:"clientAuth,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewTraefikContainoUsTLSOptionV1Alpha1Resource() resource.Resource {
	return &TraefikContainoUsTLSOptionV1Alpha1Resource{}
}

func (r *TraefikContainoUsTLSOptionV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_traefik_containo_us_tls_option_v1alpha1"
}

func (r *TraefikContainoUsTLSOptionV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "TLSOption is the CRD implementation of a Traefik TLS Option, allowing to configure some parameters of the TLS connection. More info: https://doc.traefik.io/traefik/v2.9/https/tls/#tls-options",
		MarkdownDescription: "TLSOption is the CRD implementation of a Traefik TLS Option, allowing to configure some parameters of the TLS connection. More info: https://doc.traefik.io/traefik/v2.9/https/tls/#tls-options",
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
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.RequiresReplace(),
						},
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
				Description:         "TLSOptionSpec defines the desired state of a TLSOption.",
				MarkdownDescription: "TLSOptionSpec defines the desired state of a TLSOption.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"curve_preferences": {
						Description:         "CurvePreferences defines the preferred elliptic curves in a specific order. More info: https://doc.traefik.io/traefik/v2.9/https/tls/#curve-preferences",
						MarkdownDescription: "CurvePreferences defines the preferred elliptic curves in a specific order. More info: https://doc.traefik.io/traefik/v2.9/https/tls/#curve-preferences",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"max_version": {
						Description:         "MaxVersion defines the maximum TLS version that Traefik will accept. Possible values: VersionTLS10, VersionTLS11, VersionTLS12, VersionTLS13. Default: None.",
						MarkdownDescription: "MaxVersion defines the maximum TLS version that Traefik will accept. Possible values: VersionTLS10, VersionTLS11, VersionTLS12, VersionTLS13. Default: None.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"min_version": {
						Description:         "MinVersion defines the minimum TLS version that Traefik will accept. Possible values: VersionTLS10, VersionTLS11, VersionTLS12, VersionTLS13. Default: VersionTLS10.",
						MarkdownDescription: "MinVersion defines the minimum TLS version that Traefik will accept. Possible values: VersionTLS10, VersionTLS11, VersionTLS12, VersionTLS13. Default: VersionTLS10.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"prefer_server_cipher_suites": {
						Description:         "PreferServerCipherSuites defines whether the server chooses a cipher suite among his own instead of among the client's. It is enabled automatically when minVersion or maxVersion is set. Deprecated: https://github.com/golang/go/issues/45430",
						MarkdownDescription: "PreferServerCipherSuites defines whether the server chooses a cipher suite among his own instead of among the client's. It is enabled automatically when minVersion or maxVersion is set. Deprecated: https://github.com/golang/go/issues/45430",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"sni_strict": {
						Description:         "SniStrict defines whether Traefik allows connections from clients connections that do not specify a server_name extension.",
						MarkdownDescription: "SniStrict defines whether Traefik allows connections from clients connections that do not specify a server_name extension.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"alpn_protocols": {
						Description:         "ALPNProtocols defines the list of supported application level protocols for the TLS handshake, in order of preference. More info: https://doc.traefik.io/traefik/v2.9/https/tls/#alpn-protocols",
						MarkdownDescription: "ALPNProtocols defines the list of supported application level protocols for the TLS handshake, in order of preference. More info: https://doc.traefik.io/traefik/v2.9/https/tls/#alpn-protocols",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"cipher_suites": {
						Description:         "CipherSuites defines the list of supported cipher suites for TLS versions up to TLS 1.2. More info: https://doc.traefik.io/traefik/v2.9/https/tls/#cipher-suites",
						MarkdownDescription: "CipherSuites defines the list of supported cipher suites for TLS versions up to TLS 1.2. More info: https://doc.traefik.io/traefik/v2.9/https/tls/#cipher-suites",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"client_auth": {
						Description:         "ClientAuth defines the server's policy for TLS Client Authentication.",
						MarkdownDescription: "ClientAuth defines the server's policy for TLS Client Authentication.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"client_auth_type": {
								Description:         "ClientAuthType defines the client authentication type to apply.",
								MarkdownDescription: "ClientAuthType defines the client authentication type to apply.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"secret_names": {
								Description:         "SecretNames defines the names of the referenced Kubernetes Secret storing certificate details.",
								MarkdownDescription: "SecretNames defines the names of the referenced Kubernetes Secret storing certificate details.",

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

				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}, nil
}

func (r *TraefikContainoUsTLSOptionV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_traefik_containo_us_tls_option_v1alpha1")

	var state TraefikContainoUsTLSOptionV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel TraefikContainoUsTLSOptionV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("traefik.containo.us/v1alpha1")
	goModel.Kind = utilities.Ptr("TLSOption")

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

func (r *TraefikContainoUsTLSOptionV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_traefik_containo_us_tls_option_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *TraefikContainoUsTLSOptionV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_traefik_containo_us_tls_option_v1alpha1")

	var state TraefikContainoUsTLSOptionV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel TraefikContainoUsTLSOptionV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("traefik.containo.us/v1alpha1")
	goModel.Kind = utilities.Ptr("TLSOption")

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

func (r *TraefikContainoUsTLSOptionV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_traefik_containo_us_tls_option_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
