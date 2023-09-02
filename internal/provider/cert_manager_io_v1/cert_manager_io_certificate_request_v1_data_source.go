/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package cert_manager_io_v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &CertManagerIoCertificateRequestV1DataSource{}
	_ datasource.DataSourceWithConfigure = &CertManagerIoCertificateRequestV1DataSource{}
)

func NewCertManagerIoCertificateRequestV1DataSource() datasource.DataSource {
	return &CertManagerIoCertificateRequestV1DataSource{}
}

type CertManagerIoCertificateRequestV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type CertManagerIoCertificateRequestV1DataSourceData struct {
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

func (r *CertManagerIoCertificateRequestV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cert_manager_io_certificate_request_v1"
}

func (r *CertManagerIoCertificateRequestV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "Desired state of the CertificateRequest resource.",
				MarkdownDescription: "Desired state of the CertificateRequest resource.",
				Attributes: map[string]schema.Attribute{
					"duration": schema.StringAttribute{
						Description:         "The requested 'duration' (i.e. lifetime) of the Certificate. This option may be ignored/overridden by some issuer types.",
						MarkdownDescription: "The requested 'duration' (i.e. lifetime) of the Certificate. This option may be ignored/overridden by some issuer types.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"extra": schema.MapAttribute{
						Description:         "Extra contains extra attributes of the user that created the CertificateRequest. Populated by the cert-manager webhook on creation and immutable.",
						MarkdownDescription: "Extra contains extra attributes of the user that created the CertificateRequest. Populated by the cert-manager webhook on creation and immutable.",
						ElementType:         types.ListType{ElemType: types.StringType},
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"groups": schema.ListAttribute{
						Description:         "Groups contains group membership of the user that created the CertificateRequest. Populated by the cert-manager webhook on creation and immutable.",
						MarkdownDescription: "Groups contains group membership of the user that created the CertificateRequest. Populated by the cert-manager webhook on creation and immutable.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"is_ca": schema.BoolAttribute{
						Description:         "IsCA will request to mark the certificate as valid for certificate signing when submitting to the issuer. This will automatically add the 'cert sign' usage to the list of 'usages'.",
						MarkdownDescription: "IsCA will request to mark the certificate as valid for certificate signing when submitting to the issuer. This will automatically add the 'cert sign' usage to the list of 'usages'.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"issuer_ref": schema.SingleNestedAttribute{
						Description:         "IssuerRef is a reference to the issuer for this CertificateRequest.  If the 'kind' field is not set, or set to 'Issuer', an Issuer resource with the given name in the same namespace as the CertificateRequest will be used.  If the 'kind' field is set to 'ClusterIssuer', a ClusterIssuer with the provided name will be used. The 'name' field in this stanza is required at all times. The group field refers to the API group of the issuer which defaults to 'cert-manager.io' if empty.",
						MarkdownDescription: "IssuerRef is a reference to the issuer for this CertificateRequest.  If the 'kind' field is not set, or set to 'Issuer', an Issuer resource with the given name in the same namespace as the CertificateRequest will be used.  If the 'kind' field is set to 'ClusterIssuer', a ClusterIssuer with the provided name will be used. The 'name' field in this stanza is required at all times. The group field refers to the API group of the issuer which defaults to 'cert-manager.io' if empty.",
						Attributes: map[string]schema.Attribute{
							"group": schema.StringAttribute{
								Description:         "Group of the resource being referred to.",
								MarkdownDescription: "Group of the resource being referred to.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the resource being referred to.",
								MarkdownDescription: "Kind of the resource being referred to.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the resource being referred to.",
								MarkdownDescription: "Name of the resource being referred to.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"request": schema.StringAttribute{
						Description:         "The PEM-encoded x509 certificate signing request to be submitted to the CA for signing.",
						MarkdownDescription: "The PEM-encoded x509 certificate signing request to be submitted to the CA for signing.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"uid": schema.StringAttribute{
						Description:         "UID contains the uid of the user that created the CertificateRequest. Populated by the cert-manager webhook on creation and immutable.",
						MarkdownDescription: "UID contains the uid of the user that created the CertificateRequest. Populated by the cert-manager webhook on creation and immutable.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"usages": schema.ListAttribute{
						Description:         "Usages is the set of x509 usages that are requested for the certificate. If usages are set they SHOULD be encoded inside the CSR spec Defaults to 'digital signature' and 'key encipherment' if not specified.",
						MarkdownDescription: "Usages is the set of x509 usages that are requested for the certificate. If usages are set they SHOULD be encoded inside the CSR spec Defaults to 'digital signature' and 'key encipherment' if not specified.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"username": schema.StringAttribute{
						Description:         "Username contains the name of the user that created the CertificateRequest. Populated by the cert-manager webhook on creation and immutable.",
						MarkdownDescription: "Username contains the name of the user that created the CertificateRequest. Populated by the cert-manager webhook on creation and immutable.",
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
	}
}

func (r *CertManagerIoCertificateRequestV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *CertManagerIoCertificateRequestV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_cert_manager_io_certificate_request_v1")

	var data CertManagerIoCertificateRequestV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "cert-manager.io", Version: "v1", Resource: "CertificateRequest"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
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

	var readResponse CertManagerIoCertificateRequestV1DataSourceData
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
	data.ApiVersion = pointer.String("cert-manager.io/v1")
	data.Kind = pointer.String("CertificateRequest")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
