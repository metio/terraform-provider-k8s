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

type CertificatesK8SIoCertificateSigningRequestV1Resource struct{}

var (
	_ resource.Resource = (*CertificatesK8SIoCertificateSigningRequestV1Resource)(nil)
)

type CertificatesK8SIoCertificateSigningRequestV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type CertificatesK8SIoCertificateSigningRequestV1GoModel struct {
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
		ExpirationSeconds *int64 `tfsdk:"expiration_seconds" yaml:"expirationSeconds,omitempty"`

		Extra *map[string][]string `tfsdk:"extra" yaml:"extra,omitempty"`

		Groups *[]string `tfsdk:"groups" yaml:"groups,omitempty"`

		Request *string `tfsdk:"request" yaml:"request,omitempty"`

		SignerName *string `tfsdk:"signer_name" yaml:"signerName,omitempty"`

		Uid *string `tfsdk:"uid" yaml:"uid,omitempty"`

		Usages *[]string `tfsdk:"usages" yaml:"usages,omitempty"`

		Username *string `tfsdk:"username" yaml:"username,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewCertificatesK8SIoCertificateSigningRequestV1Resource() resource.Resource {
	return &CertificatesK8SIoCertificateSigningRequestV1Resource{}
}

func (r *CertificatesK8SIoCertificateSigningRequestV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_certificates_k8s_io_certificate_signing_request_v1"
}

func (r *CertificatesK8SIoCertificateSigningRequestV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "CertificateSigningRequest objects provide a mechanism to obtain x509 certificates by submitting a certificate signing request, and having it asynchronously approved and issued.Kubelets use this API to obtain: 1. client certificates to authenticate to kube-apiserver (with the 'kubernetes.io/kube-apiserver-client-kubelet' signerName). 2. serving certificates for TLS endpoints kube-apiserver can connect to securely (with the 'kubernetes.io/kubelet-serving' signerName).This API can be used to request client certificates to authenticate to kube-apiserver (with the 'kubernetes.io/kube-apiserver-client' signerName), or to obtain certificates from custom non-Kubernetes signers.",
		MarkdownDescription: "CertificateSigningRequest objects provide a mechanism to obtain x509 certificates by submitting a certificate signing request, and having it asynchronously approved and issued.Kubelets use this API to obtain: 1. client certificates to authenticate to kube-apiserver (with the 'kubernetes.io/kube-apiserver-client-kubelet' signerName). 2. serving certificates for TLS endpoints kube-apiserver can connect to securely (with the 'kubernetes.io/kubelet-serving' signerName).This API can be used to request client certificates to authenticate to kube-apiserver (with the 'kubernetes.io/kube-apiserver-client' signerName), or to obtain certificates from custom non-Kubernetes signers.",
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
				Description:         "CertificateSigningRequestSpec contains the certificate request.",
				MarkdownDescription: "CertificateSigningRequestSpec contains the certificate request.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"expiration_seconds": {
						Description:         "expirationSeconds is the requested duration of validity of the issued certificate. The certificate signer may issue a certificate with a different validity duration so a client must check the delta between the notBefore and and notAfter fields in the issued certificate to determine the actual duration.The v1.22+ in-tree implementations of the well-known Kubernetes signers will honor this field as long as the requested duration is not greater than the maximum duration they will honor per the --cluster-signing-duration CLI flag to the Kubernetes controller manager.Certificate signers may not honor this field for various reasons:  1. Old signer that is unaware of the field (such as the in-tree     implementations prior to v1.22)  2. Signer whose configured maximum is shorter than the requested duration  3. Signer whose configured minimum is longer than the requested durationThe minimum valid value for expirationSeconds is 600, i.e. 10 minutes.",
						MarkdownDescription: "expirationSeconds is the requested duration of validity of the issued certificate. The certificate signer may issue a certificate with a different validity duration so a client must check the delta between the notBefore and and notAfter fields in the issued certificate to determine the actual duration.The v1.22+ in-tree implementations of the well-known Kubernetes signers will honor this field as long as the requested duration is not greater than the maximum duration they will honor per the --cluster-signing-duration CLI flag to the Kubernetes controller manager.Certificate signers may not honor this field for various reasons:  1. Old signer that is unaware of the field (such as the in-tree     implementations prior to v1.22)  2. Signer whose configured maximum is shorter than the requested duration  3. Signer whose configured minimum is longer than the requested durationThe minimum valid value for expirationSeconds is 600, i.e. 10 minutes.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"extra": {
						Description:         "extra contains extra attributes of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.",
						MarkdownDescription: "extra contains extra attributes of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.",

						Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"groups": {
						Description:         "groups contains group membership of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.",
						MarkdownDescription: "groups contains group membership of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"request": {
						Description:         "request contains an x509 certificate signing request encoded in a 'CERTIFICATE REQUEST' PEM block. When serialized as JSON or YAML, the data is additionally base64-encoded.",
						MarkdownDescription: "request contains an x509 certificate signing request encoded in a 'CERTIFICATE REQUEST' PEM block. When serialized as JSON or YAML, the data is additionally base64-encoded.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							validators.Base64Validator(),
						},
					},

					"signer_name": {
						Description:         "signerName indicates the requested signer, and is a qualified name.List/watch requests for CertificateSigningRequests can filter on this field using a 'spec.signerName=NAME' fieldSelector.Well-known Kubernetes signers are: 1. 'kubernetes.io/kube-apiserver-client': issues client certificates that can be used to authenticate to kube-apiserver.  Requests for this signer are never auto-approved by kube-controller-manager, can be issued by the 'csrsigning' controller in kube-controller-manager. 2. 'kubernetes.io/kube-apiserver-client-kubelet': issues client certificates that kubelets use to authenticate to kube-apiserver.  Requests for this signer can be auto-approved by the 'csrapproving' controller in kube-controller-manager, and can be issued by the 'csrsigning' controller in kube-controller-manager. 3. 'kubernetes.io/kubelet-serving' issues serving certificates that kubelets use to serve TLS endpoints, which kube-apiserver can connect to securely.  Requests for this signer are never auto-approved by kube-controller-manager, and can be issued by the 'csrsigning' controller in kube-controller-manager.More details are available at https://k8s.io/docs/reference/access-authn-authz/certificate-signing-requests/#kubernetes-signersCustom signerNames can also be specified. The signer defines: 1. Trust distribution: how trust (CA bundles) are distributed. 2. Permitted subjects: and behavior when a disallowed subject is requested. 3. Required, permitted, or forbidden x509 extensions in the request (including whether subjectAltNames are allowed, which types, restrictions on allowed values) and behavior when a disallowed extension is requested. 4. Required, permitted, or forbidden key usages / extended key usages. 5. Expiration/certificate lifetime: whether it is fixed by the signer, configurable by the admin. 6. Whether or not requests for CA certificates are allowed.",
						MarkdownDescription: "signerName indicates the requested signer, and is a qualified name.List/watch requests for CertificateSigningRequests can filter on this field using a 'spec.signerName=NAME' fieldSelector.Well-known Kubernetes signers are: 1. 'kubernetes.io/kube-apiserver-client': issues client certificates that can be used to authenticate to kube-apiserver.  Requests for this signer are never auto-approved by kube-controller-manager, can be issued by the 'csrsigning' controller in kube-controller-manager. 2. 'kubernetes.io/kube-apiserver-client-kubelet': issues client certificates that kubelets use to authenticate to kube-apiserver.  Requests for this signer can be auto-approved by the 'csrapproving' controller in kube-controller-manager, and can be issued by the 'csrsigning' controller in kube-controller-manager. 3. 'kubernetes.io/kubelet-serving' issues serving certificates that kubelets use to serve TLS endpoints, which kube-apiserver can connect to securely.  Requests for this signer are never auto-approved by kube-controller-manager, and can be issued by the 'csrsigning' controller in kube-controller-manager.More details are available at https://k8s.io/docs/reference/access-authn-authz/certificate-signing-requests/#kubernetes-signersCustom signerNames can also be specified. The signer defines: 1. Trust distribution: how trust (CA bundles) are distributed. 2. Permitted subjects: and behavior when a disallowed subject is requested. 3. Required, permitted, or forbidden x509 extensions in the request (including whether subjectAltNames are allowed, which types, restrictions on allowed values) and behavior when a disallowed extension is requested. 4. Required, permitted, or forbidden key usages / extended key usages. 5. Expiration/certificate lifetime: whether it is fixed by the signer, configurable by the admin. 6. Whether or not requests for CA certificates are allowed.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"uid": {
						Description:         "uid contains the uid of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.",
						MarkdownDescription: "uid contains the uid of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"usages": {
						Description:         "usages specifies a set of key usages requested in the issued certificate.Requests for TLS client certificates typically request: 'digital signature', 'key encipherment', 'client auth'.Requests for TLS serving certificates typically request: 'key encipherment', 'digital signature', 'server auth'.Valid values are: 'signing', 'digital signature', 'content commitment', 'key encipherment', 'key agreement', 'data encipherment', 'cert sign', 'crl sign', 'encipher only', 'decipher only', 'any', 'server auth', 'client auth', 'code signing', 'email protection', 's/mime', 'ipsec end system', 'ipsec tunnel', 'ipsec user', 'timestamping', 'ocsp signing', 'microsoft sgc', 'netscape sgc'",
						MarkdownDescription: "usages specifies a set of key usages requested in the issued certificate.Requests for TLS client certificates typically request: 'digital signature', 'key encipherment', 'client auth'.Requests for TLS serving certificates typically request: 'key encipherment', 'digital signature', 'server auth'.Valid values are: 'signing', 'digital signature', 'content commitment', 'key encipherment', 'key agreement', 'data encipherment', 'cert sign', 'crl sign', 'encipher only', 'decipher only', 'any', 'server auth', 'client auth', 'code signing', 'email protection', 's/mime', 'ipsec end system', 'ipsec tunnel', 'ipsec user', 'timestamping', 'ocsp signing', 'microsoft sgc', 'netscape sgc'",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"username": {
						Description:         "username contains the name of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.",
						MarkdownDescription: "username contains the name of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.",

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

func (r *CertificatesK8SIoCertificateSigningRequestV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_certificates_k8s_io_certificate_signing_request_v1")

	var state CertificatesK8SIoCertificateSigningRequestV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CertificatesK8SIoCertificateSigningRequestV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("certificates.k8s.io/v1")
	goModel.Kind = utilities.Ptr("CertificateSigningRequest")

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

func (r *CertificatesK8SIoCertificateSigningRequestV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_certificates_k8s_io_certificate_signing_request_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *CertificatesK8SIoCertificateSigningRequestV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_certificates_k8s_io_certificate_signing_request_v1")

	var state CertificatesK8SIoCertificateSigningRequestV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CertificatesK8SIoCertificateSigningRequestV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("certificates.k8s.io/v1")
	goModel.Kind = utilities.Ptr("CertificateSigningRequest")

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

func (r *CertificatesK8SIoCertificateSigningRequestV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_certificates_k8s_io_certificate_signing_request_v1")
	// NO-OP: Terraform removes the state automatically for us
}
