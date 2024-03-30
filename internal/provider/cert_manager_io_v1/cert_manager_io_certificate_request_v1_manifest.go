/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package cert_manager_io_v1

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
	_ datasource.DataSource = &CertManagerIoCertificateRequestV1Manifest{}
)

func NewCertManagerIoCertificateRequestV1Manifest() datasource.DataSource {
	return &CertManagerIoCertificateRequestV1Manifest{}
}

type CertManagerIoCertificateRequestV1Manifest struct{}

type CertManagerIoCertificateRequestV1ManifestData struct {
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
		Duration  *string              `tfsdk:"duration" json:"duration,omitempty"`
		Extra     *map[string][]string `tfsdk:"extra" json:"extra,omitempty"`
		Groups    *[]string            `tfsdk:"groups" json:"groups,omitempty"`
		IsCA      *bool                `tfsdk:"is_ca" json:"isCA,omitempty"`
		IssuerRef *struct {
			Group *string `tfsdk:"group" json:"group,omitempty"`
			Kind  *string `tfsdk:"kind" json:"kind,omitempty"`
			Name  *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"issuer_ref" json:"issuerRef,omitempty"`
		Request  *string   `tfsdk:"request" json:"request,omitempty"`
		Uid      *string   `tfsdk:"uid" json:"uid,omitempty"`
		Usages   *[]string `tfsdk:"usages" json:"usages,omitempty"`
		Username *string   `tfsdk:"username" json:"username,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CertManagerIoCertificateRequestV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cert_manager_io_certificate_request_v1_manifest"
}

func (r *CertManagerIoCertificateRequestV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "A CertificateRequest is used to request a signed certificate from one of theconfigured issuers.All fields within the CertificateRequest's 'spec' are immutable after creation.A CertificateRequest will either succeed or fail, as denoted by its 'Ready' statuscondition and its 'status.failureTime' field.A CertificateRequest is a one-shot resource, meaning it represents a singlepoint in time request for a certificate and cannot be re-used.",
		MarkdownDescription: "A CertificateRequest is used to request a signed certificate from one of theconfigured issuers.All fields within the CertificateRequest's 'spec' are immutable after creation.A CertificateRequest will either succeed or fail, as denoted by its 'Ready' statuscondition and its 'status.failureTime' field.A CertificateRequest is a one-shot resource, meaning it represents a singlepoint in time request for a certificate and cannot be re-used.",
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
				Description:         "Specification of the desired state of the CertificateRequest resource.https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
				MarkdownDescription: "Specification of the desired state of the CertificateRequest resource.https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
				Attributes: map[string]schema.Attribute{
					"duration": schema.StringAttribute{
						Description:         "Requested 'duration' (i.e. lifetime) of the Certificate. Note that theissuer may choose to ignore the requested duration, just like any otherrequested attribute.",
						MarkdownDescription: "Requested 'duration' (i.e. lifetime) of the Certificate. Note that theissuer may choose to ignore the requested duration, just like any otherrequested attribute.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"extra": schema.MapAttribute{
						Description:         "Extra contains extra attributes of the user that created the CertificateRequest.Populated by the cert-manager webhook on creation and immutable.",
						MarkdownDescription: "Extra contains extra attributes of the user that created the CertificateRequest.Populated by the cert-manager webhook on creation and immutable.",
						ElementType:         types.ListType{ElemType: types.StringType},
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"groups": schema.ListAttribute{
						Description:         "Groups contains group membership of the user that created the CertificateRequest.Populated by the cert-manager webhook on creation and immutable.",
						MarkdownDescription: "Groups contains group membership of the user that created the CertificateRequest.Populated by the cert-manager webhook on creation and immutable.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"is_ca": schema.BoolAttribute{
						Description:         "Requested basic constraints isCA value. Note that the issuer may chooseto ignore the requested isCA value, just like any other requested attribute.NOTE: If the CSR in the 'Request' field has a BasicConstraints extension,it must have the same isCA value as specified here.If true, this will automatically add the 'cert sign' usage to the listof requested 'usages'.",
						MarkdownDescription: "Requested basic constraints isCA value. Note that the issuer may chooseto ignore the requested isCA value, just like any other requested attribute.NOTE: If the CSR in the 'Request' field has a BasicConstraints extension,it must have the same isCA value as specified here.If true, this will automatically add the 'cert sign' usage to the listof requested 'usages'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"issuer_ref": schema.SingleNestedAttribute{
						Description:         "Reference to the issuer responsible for issuing the certificate.If the issuer is namespace-scoped, it must be in the same namespaceas the Certificate. If the issuer is cluster-scoped, it can be usedfrom any namespace.The 'name' field of the reference must always be specified.",
						MarkdownDescription: "Reference to the issuer responsible for issuing the certificate.If the issuer is namespace-scoped, it must be in the same namespaceas the Certificate. If the issuer is cluster-scoped, it can be usedfrom any namespace.The 'name' field of the reference must always be specified.",
						Attributes: map[string]schema.Attribute{
							"group": schema.StringAttribute{
								Description:         "Group of the resource being referred to.",
								MarkdownDescription: "Group of the resource being referred to.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the resource being referred to.",
								MarkdownDescription: "Kind of the resource being referred to.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the resource being referred to.",
								MarkdownDescription: "Name of the resource being referred to.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"request": schema.StringAttribute{
						Description:         "The PEM-encoded X.509 certificate signing request to be submitted to theissuer for signing.If the CSR has a BasicConstraints extension, its isCA attribute mustmatch the 'isCA' value of this CertificateRequest.If the CSR has a KeyUsage extension, its key usages must match thekey usages in the 'usages' field of this CertificateRequest.If the CSR has a ExtKeyUsage extension, its extended key usagesmust match the extended key usages in the 'usages' field of thisCertificateRequest.",
						MarkdownDescription: "The PEM-encoded X.509 certificate signing request to be submitted to theissuer for signing.If the CSR has a BasicConstraints extension, its isCA attribute mustmatch the 'isCA' value of this CertificateRequest.If the CSR has a KeyUsage extension, its key usages must match thekey usages in the 'usages' field of this CertificateRequest.If the CSR has a ExtKeyUsage extension, its extended key usagesmust match the extended key usages in the 'usages' field of thisCertificateRequest.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							validators.Base64Validator(),
						},
					},

					"uid": schema.StringAttribute{
						Description:         "UID contains the uid of the user that created the CertificateRequest.Populated by the cert-manager webhook on creation and immutable.",
						MarkdownDescription: "UID contains the uid of the user that created the CertificateRequest.Populated by the cert-manager webhook on creation and immutable.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"usages": schema.ListAttribute{
						Description:         "Requested key usages and extended key usages.NOTE: If the CSR in the 'Request' field has uses the KeyUsage orExtKeyUsage extension, these extensions must have the same valuesas specified here without any additional values.If unset, defaults to 'digital signature' and 'key encipherment'.",
						MarkdownDescription: "Requested key usages and extended key usages.NOTE: If the CSR in the 'Request' field has uses the KeyUsage orExtKeyUsage extension, these extensions must have the same valuesas specified here without any additional values.If unset, defaults to 'digital signature' and 'key encipherment'.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"username": schema.StringAttribute{
						Description:         "Username contains the name of the user that created the CertificateRequest.Populated by the cert-manager webhook on creation and immutable.",
						MarkdownDescription: "Username contains the name of the user that created the CertificateRequest.Populated by the cert-manager webhook on creation and immutable.",
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

func (r *CertManagerIoCertificateRequestV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cert_manager_io_certificate_request_v1_manifest")

	var model CertManagerIoCertificateRequestV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("cert-manager.io/v1")
	model.Kind = pointer.String("CertificateRequest")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
