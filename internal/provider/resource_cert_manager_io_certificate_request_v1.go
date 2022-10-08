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

type CertManagerIoCertificateRequestV1Resource struct{}

var (
	_ resource.Resource = (*CertManagerIoCertificateRequestV1Resource)(nil)
)

type CertManagerIoCertificateRequestV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type CertManagerIoCertificateRequestV1GoModel struct {
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
		Usages *[]string `tfsdk:"usages" yaml:"usages,omitempty"`

		Username *string `tfsdk:"username" yaml:"username,omitempty"`

		Extra *map[string][]string `tfsdk:"extra" yaml:"extra,omitempty"`

		Groups *[]string `tfsdk:"groups" yaml:"groups,omitempty"`

		IsCA *bool `tfsdk:"is_ca" yaml:"isCA,omitempty"`

		Request *string `tfsdk:"request" yaml:"request,omitempty"`

		Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

		IssuerRef *struct {
			Group *string `tfsdk:"group" yaml:"group,omitempty"`

			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"issuer_ref" yaml:"issuerRef,omitempty"`

		Uid *string `tfsdk:"uid" yaml:"uid,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewCertManagerIoCertificateRequestV1Resource() resource.Resource {
	return &CertManagerIoCertificateRequestV1Resource{}
}

func (r *CertManagerIoCertificateRequestV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cert_manager_io_certificate_request_v1"
}

func (r *CertManagerIoCertificateRequestV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "A CertificateRequest is used to request a signed certificate from one of the configured issuers.  All fields within the CertificateRequest's 'spec' are immutable after creation. A CertificateRequest will either succeed or fail, as denoted by its 'status.state' field.  A CertificateRequest is a one-shot resource, meaning it represents a single point in time request for a certificate and cannot be re-used.",
		MarkdownDescription: "A CertificateRequest is used to request a signed certificate from one of the configured issuers.  All fields within the CertificateRequest's 'spec' are immutable after creation. A CertificateRequest will either succeed or fail, as denoted by its 'status.state' field.  A CertificateRequest is a one-shot resource, meaning it represents a single point in time request for a certificate and cannot be re-used.",
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
				Description:         "Desired state of the CertificateRequest resource.",
				MarkdownDescription: "Desired state of the CertificateRequest resource.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"usages": {
						Description:         "Usages is the set of x509 usages that are requested for the certificate. If usages are set they SHOULD be encoded inside the CSR spec Defaults to 'digital signature' and 'key encipherment' if not specified.",
						MarkdownDescription: "Usages is the set of x509 usages that are requested for the certificate. If usages are set they SHOULD be encoded inside the CSR spec Defaults to 'digital signature' and 'key encipherment' if not specified.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"username": {
						Description:         "Username contains the name of the user that created the CertificateRequest. Populated by the cert-manager webhook on creation and immutable.",
						MarkdownDescription: "Username contains the name of the user that created the CertificateRequest. Populated by the cert-manager webhook on creation and immutable.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"extra": {
						Description:         "Extra contains extra attributes of the user that created the CertificateRequest. Populated by the cert-manager webhook on creation and immutable.",
						MarkdownDescription: "Extra contains extra attributes of the user that created the CertificateRequest. Populated by the cert-manager webhook on creation and immutable.",

						Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"groups": {
						Description:         "Groups contains group membership of the user that created the CertificateRequest. Populated by the cert-manager webhook on creation and immutable.",
						MarkdownDescription: "Groups contains group membership of the user that created the CertificateRequest. Populated by the cert-manager webhook on creation and immutable.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"is_ca": {
						Description:         "IsCA will request to mark the certificate as valid for certificate signing when submitting to the issuer. This will automatically add the 'cert sign' usage to the list of 'usages'.",
						MarkdownDescription: "IsCA will request to mark the certificate as valid for certificate signing when submitting to the issuer. This will automatically add the 'cert sign' usage to the list of 'usages'.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"request": {
						Description:         "The PEM-encoded x509 certificate signing request to be submitted to the CA for signing.",
						MarkdownDescription: "The PEM-encoded x509 certificate signing request to be submitted to the CA for signing.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"duration": {
						Description:         "The requested 'duration' (i.e. lifetime) of the Certificate. This option may be ignored/overridden by some issuer types.",
						MarkdownDescription: "The requested 'duration' (i.e. lifetime) of the Certificate. This option may be ignored/overridden by some issuer types.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"issuer_ref": {
						Description:         "IssuerRef is a reference to the issuer for this CertificateRequest.  If the 'kind' field is not set, or set to 'Issuer', an Issuer resource with the given name in the same namespace as the CertificateRequest will be used.  If the 'kind' field is set to 'ClusterIssuer', a ClusterIssuer with the provided name will be used. The 'name' field in this stanza is required at all times. The group field refers to the API group of the issuer which defaults to 'cert-manager.io' if empty.",
						MarkdownDescription: "IssuerRef is a reference to the issuer for this CertificateRequest.  If the 'kind' field is not set, or set to 'Issuer', an Issuer resource with the given name in the same namespace as the CertificateRequest will be used.  If the 'kind' field is set to 'ClusterIssuer', a ClusterIssuer with the provided name will be used. The 'name' field in this stanza is required at all times. The group field refers to the API group of the issuer which defaults to 'cert-manager.io' if empty.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"group": {
								Description:         "Group of the resource being referred to.",
								MarkdownDescription: "Group of the resource being referred to.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"kind": {
								Description:         "Kind of the resource being referred to.",
								MarkdownDescription: "Kind of the resource being referred to.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "Name of the resource being referred to.",
								MarkdownDescription: "Name of the resource being referred to.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"uid": {
						Description:         "UID contains the uid of the user that created the CertificateRequest. Populated by the cert-manager webhook on creation and immutable.",
						MarkdownDescription: "UID contains the uid of the user that created the CertificateRequest. Populated by the cert-manager webhook on creation and immutable.",

						Type: types.StringType,

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

func (r *CertManagerIoCertificateRequestV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_cert_manager_io_certificate_request_v1")

	var state CertManagerIoCertificateRequestV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CertManagerIoCertificateRequestV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("cert-manager.io/v1")
	goModel.Kind = utilities.Ptr("CertificateRequest")

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

func (r *CertManagerIoCertificateRequestV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cert_manager_io_certificate_request_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *CertManagerIoCertificateRequestV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_cert_manager_io_certificate_request_v1")

	var state CertManagerIoCertificateRequestV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CertManagerIoCertificateRequestV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("cert-manager.io/v1")
	goModel.Kind = utilities.Ptr("CertificateRequest")

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

func (r *CertManagerIoCertificateRequestV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_cert_manager_io_certificate_request_v1")
	// NO-OP: Terraform removes the state automatically for us
}
