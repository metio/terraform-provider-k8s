/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package cert_manager_io_v1

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
	_ datasource.DataSource = &CertManagerIoCertificateRequestV1Manifest{}
)

func NewCertManagerIoCertificateRequestV1Manifest() datasource.DataSource {
	return &CertManagerIoCertificateRequestV1Manifest{}
}

type CertManagerIoCertificateRequestV1Manifest struct{}

type CertManagerIoCertificateRequestV1ManifestData struct {
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
		Description:         "A CertificateRequest is used to request a signed certificate from one of the configured issuers.  All fields within the CertificateRequest's 'spec' are immutable after creation. A CertificateRequest will either succeed or fail, as denoted by its 'status.state' field.  A CertificateRequest is a one-shot resource, meaning it represents a single point in time request for a certificate and cannot be re-used.",
		MarkdownDescription: "A CertificateRequest is used to request a signed certificate from one of the configured issuers.  All fields within the CertificateRequest's 'spec' are immutable after creation. A CertificateRequest will either succeed or fail, as denoted by its 'status.state' field.  A CertificateRequest is a one-shot resource, meaning it represents a single point in time request for a certificate and cannot be re-used.",
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
				Description:         "Desired state of the CertificateRequest resource.",
				MarkdownDescription: "Desired state of the CertificateRequest resource.",
				Attributes: map[string]schema.Attribute{
					"duration": schema.StringAttribute{
						Description:         "The requested 'duration' (i.e. lifetime) of the Certificate. This option may be ignored/overridden by some issuer types.",
						MarkdownDescription: "The requested 'duration' (i.e. lifetime) of the Certificate. This option may be ignored/overridden by some issuer types.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"extra": schema.MapAttribute{
						Description:         "Extra contains extra attributes of the user that created the CertificateRequest. Populated by the cert-manager webhook on creation and immutable.",
						MarkdownDescription: "Extra contains extra attributes of the user that created the CertificateRequest. Populated by the cert-manager webhook on creation and immutable.",
						ElementType:         types.ListType{ElemType: types.StringType},
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"groups": schema.ListAttribute{
						Description:         "Groups contains group membership of the user that created the CertificateRequest. Populated by the cert-manager webhook on creation and immutable.",
						MarkdownDescription: "Groups contains group membership of the user that created the CertificateRequest. Populated by the cert-manager webhook on creation and immutable.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"is_ca": schema.BoolAttribute{
						Description:         "IsCA will request to mark the certificate as valid for certificate signing when submitting to the issuer. This will automatically add the 'cert sign' usage to the list of 'usages'.",
						MarkdownDescription: "IsCA will request to mark the certificate as valid for certificate signing when submitting to the issuer. This will automatically add the 'cert sign' usage to the list of 'usages'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"issuer_ref": schema.SingleNestedAttribute{
						Description:         "IssuerRef is a reference to the issuer for this CertificateRequest.  If the 'kind' field is not set, or set to 'Issuer', an Issuer resource with the given name in the same namespace as the CertificateRequest will be used.  If the 'kind' field is set to 'ClusterIssuer', a ClusterIssuer with the provided name will be used. The 'name' field in this stanza is required at all times. The group field refers to the API group of the issuer which defaults to 'cert-manager.io' if empty.",
						MarkdownDescription: "IssuerRef is a reference to the issuer for this CertificateRequest.  If the 'kind' field is not set, or set to 'Issuer', an Issuer resource with the given name in the same namespace as the CertificateRequest will be used.  If the 'kind' field is set to 'ClusterIssuer', a ClusterIssuer with the provided name will be used. The 'name' field in this stanza is required at all times. The group field refers to the API group of the issuer which defaults to 'cert-manager.io' if empty.",
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
						Description:         "The PEM-encoded x509 certificate signing request to be submitted to the CA for signing.",
						MarkdownDescription: "The PEM-encoded x509 certificate signing request to be submitted to the CA for signing.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							validators.Base64Validator(),
						},
					},

					"uid": schema.StringAttribute{
						Description:         "UID contains the uid of the user that created the CertificateRequest. Populated by the cert-manager webhook on creation and immutable.",
						MarkdownDescription: "UID contains the uid of the user that created the CertificateRequest. Populated by the cert-manager webhook on creation and immutable.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"usages": schema.ListAttribute{
						Description:         "Usages is the set of x509 usages that are requested for the certificate. If usages are set they SHOULD be encoded inside the CSR spec Defaults to 'digital signature' and 'key encipherment' if not specified.",
						MarkdownDescription: "Usages is the set of x509 usages that are requested for the certificate. If usages are set they SHOULD be encoded inside the CSR spec Defaults to 'digital signature' and 'key encipherment' if not specified.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"username": schema.StringAttribute{
						Description:         "Username contains the name of the user that created the CertificateRequest. Populated by the cert-manager webhook on creation and immutable.",
						MarkdownDescription: "Username contains the name of the user that created the CertificateRequest. Populated by the cert-manager webhook on creation and immutable.",
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
	}
}

func (r *CertManagerIoCertificateRequestV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cert_manager_io_certificate_request_v1_manifest")

	var model CertManagerIoCertificateRequestV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("cert-manager.io/v1")
	model.Kind = pointer.String("CertificateRequest")

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
