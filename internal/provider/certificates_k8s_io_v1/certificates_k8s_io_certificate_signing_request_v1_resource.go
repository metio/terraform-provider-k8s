/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package certificates_k8s_io_v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	k8sTypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ resource.Resource                = &CertificatesK8SIoCertificateSigningRequestV1Resource{}
	_ resource.ResourceWithConfigure   = &CertificatesK8SIoCertificateSigningRequestV1Resource{}
	_ resource.ResourceWithImportState = &CertificatesK8SIoCertificateSigningRequestV1Resource{}
)

func NewCertificatesK8SIoCertificateSigningRequestV1Resource() resource.Resource {
	return &CertificatesK8SIoCertificateSigningRequestV1Resource{}
}

type CertificatesK8SIoCertificateSigningRequestV1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type CertificatesK8SIoCertificateSigningRequestV1ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitFor        types.List   `tfsdk:"wait_for" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ExpirationSeconds *int64               `tfsdk:"expiration_seconds" json:"expirationSeconds,omitempty"`
		Extra             *map[string][]string `tfsdk:"extra" json:"extra,omitempty"`
		Groups            *[]string            `tfsdk:"groups" json:"groups,omitempty"`
		Request           *string              `tfsdk:"request" json:"request,omitempty"`
		SignerName        *string              `tfsdk:"signer_name" json:"signerName,omitempty"`
		Uid               *string              `tfsdk:"uid" json:"uid,omitempty"`
		Usages            *[]string            `tfsdk:"usages" json:"usages,omitempty"`
		Username          *string              `tfsdk:"username" json:"username,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CertificatesK8SIoCertificateSigningRequestV1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_certificates_k8s_io_certificate_signing_request_v1"
}

func (r *CertificatesK8SIoCertificateSigningRequestV1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CertificateSigningRequest objects provide a mechanism to obtain x509 certificates by submitting a certificate signing request, and having it asynchronously approved and issued.Kubelets use this API to obtain: 1. client certificates to authenticate to kube-apiserver (with the 'kubernetes.io/kube-apiserver-client-kubelet' signerName). 2. serving certificates for TLS endpoints kube-apiserver can connect to securely (with the 'kubernetes.io/kubelet-serving' signerName).This API can be used to request client certificates to authenticate to kube-apiserver (with the 'kubernetes.io/kube-apiserver-client' signerName), or to obtain certificates from custom non-Kubernetes signers.",
		MarkdownDescription: "CertificateSigningRequest objects provide a mechanism to obtain x509 certificates by submitting a certificate signing request, and having it asynchronously approved and issued.Kubelets use this API to obtain: 1. client certificates to authenticate to kube-apiserver (with the 'kubernetes.io/kube-apiserver-client-kubelet' signerName). 2. serving certificates for TLS endpoints kube-apiserver can connect to securely (with the 'kubernetes.io/kubelet-serving' signerName).This API can be used to request client certificates to authenticate to kube-apiserver (with the 'kubernetes.io/kube-apiserver-client' signerName), or to obtain certificates from custom non-Kubernetes signers.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"force_conflicts": schema.BoolAttribute{
				Description:         "If 'true', server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "If `true`, server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"field_manager": schema.BoolAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"wait_for": schema.ListNestedAttribute{
				Description:         "Wait for specific conditions after create/update of resources.",
				MarkdownDescription: "Wait for specific conditions after create/update of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"jsonpath": schema.StringAttribute{
							Description:         "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							MarkdownDescription: "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},
						"value": schema.StringAttribute{
							Description:         "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							MarkdownDescription: "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							Required:            false,
							Optional:            true,
							Computed:            true,
						},
						"timeout": schema.StringAttribute{
							Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("30s"),
						},
					},
				},
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            true,
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
						Computed:            true,
						Validators: []validator.Map{
							validators.AnnotationValidator(),
						},
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "CertificateSigningRequestSpec contains the certificate request.",
				MarkdownDescription: "CertificateSigningRequestSpec contains the certificate request.",
				Attributes: map[string]schema.Attribute{
					"expiration_seconds": schema.Int64Attribute{
						Description:         "expirationSeconds is the requested duration of validity of the issued certificate. The certificate signer may issue a certificate with a different validity duration so a client must check the delta between the notBefore and and notAfter fields in the issued certificate to determine the actual duration.The v1.22+ in-tree implementations of the well-known Kubernetes signers will honor this field as long as the requested duration is not greater than the maximum duration they will honor per the --cluster-signing-duration CLI flag to the Kubernetes controller manager.Certificate signers may not honor this field for various reasons:  1. Old signer that is unaware of the field (such as the in-tree     implementations prior to v1.22)  2. Signer whose configured maximum is shorter than the requested duration  3. Signer whose configured minimum is longer than the requested durationThe minimum valid value for expirationSeconds is 600, i.e. 10 minutes.",
						MarkdownDescription: "expirationSeconds is the requested duration of validity of the issued certificate. The certificate signer may issue a certificate with a different validity duration so a client must check the delta between the notBefore and and notAfter fields in the issued certificate to determine the actual duration.The v1.22+ in-tree implementations of the well-known Kubernetes signers will honor this field as long as the requested duration is not greater than the maximum duration they will honor per the --cluster-signing-duration CLI flag to the Kubernetes controller manager.Certificate signers may not honor this field for various reasons:  1. Old signer that is unaware of the field (such as the in-tree     implementations prior to v1.22)  2. Signer whose configured maximum is shorter than the requested duration  3. Signer whose configured minimum is longer than the requested durationThe minimum valid value for expirationSeconds is 600, i.e. 10 minutes.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"extra": schema.MapAttribute{
						Description:         "extra contains extra attributes of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.",
						MarkdownDescription: "extra contains extra attributes of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.",
						ElementType:         types.ListType{ElemType: types.StringType},
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"groups": schema.ListAttribute{
						Description:         "groups contains group membership of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.",
						MarkdownDescription: "groups contains group membership of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"request": schema.StringAttribute{
						Description:         "request contains an x509 certificate signing request encoded in a 'CERTIFICATE REQUEST' PEM block. When serialized as JSON or YAML, the data is additionally base64-encoded.",
						MarkdownDescription: "request contains an x509 certificate signing request encoded in a 'CERTIFICATE REQUEST' PEM block. When serialized as JSON or YAML, the data is additionally base64-encoded.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							validators.Base64Validator(),
						},
					},

					"signer_name": schema.StringAttribute{
						Description:         "signerName indicates the requested signer, and is a qualified name.List/watch requests for CertificateSigningRequests can filter on this field using a 'spec.signerName=NAME' fieldSelector.Well-known Kubernetes signers are: 1. 'kubernetes.io/kube-apiserver-client': issues client certificates that can be used to authenticate to kube-apiserver.  Requests for this signer are never auto-approved by kube-controller-manager, can be issued by the 'csrsigning' controller in kube-controller-manager. 2. 'kubernetes.io/kube-apiserver-client-kubelet': issues client certificates that kubelets use to authenticate to kube-apiserver.  Requests for this signer can be auto-approved by the 'csrapproving' controller in kube-controller-manager, and can be issued by the 'csrsigning' controller in kube-controller-manager. 3. 'kubernetes.io/kubelet-serving' issues serving certificates that kubelets use to serve TLS endpoints, which kube-apiserver can connect to securely.  Requests for this signer are never auto-approved by kube-controller-manager, and can be issued by the 'csrsigning' controller in kube-controller-manager.More details are available at https://k8s.io/docs/reference/access-authn-authz/certificate-signing-requests/#kubernetes-signersCustom signerNames can also be specified. The signer defines: 1. Trust distribution: how trust (CA bundles) are distributed. 2. Permitted subjects: and behavior when a disallowed subject is requested. 3. Required, permitted, or forbidden x509 extensions in the request (including whether subjectAltNames are allowed, which types, restrictions on allowed values) and behavior when a disallowed extension is requested. 4. Required, permitted, or forbidden key usages / extended key usages. 5. Expiration/certificate lifetime: whether it is fixed by the signer, configurable by the admin. 6. Whether or not requests for CA certificates are allowed.",
						MarkdownDescription: "signerName indicates the requested signer, and is a qualified name.List/watch requests for CertificateSigningRequests can filter on this field using a 'spec.signerName=NAME' fieldSelector.Well-known Kubernetes signers are: 1. 'kubernetes.io/kube-apiserver-client': issues client certificates that can be used to authenticate to kube-apiserver.  Requests for this signer are never auto-approved by kube-controller-manager, can be issued by the 'csrsigning' controller in kube-controller-manager. 2. 'kubernetes.io/kube-apiserver-client-kubelet': issues client certificates that kubelets use to authenticate to kube-apiserver.  Requests for this signer can be auto-approved by the 'csrapproving' controller in kube-controller-manager, and can be issued by the 'csrsigning' controller in kube-controller-manager. 3. 'kubernetes.io/kubelet-serving' issues serving certificates that kubelets use to serve TLS endpoints, which kube-apiserver can connect to securely.  Requests for this signer are never auto-approved by kube-controller-manager, and can be issued by the 'csrsigning' controller in kube-controller-manager.More details are available at https://k8s.io/docs/reference/access-authn-authz/certificate-signing-requests/#kubernetes-signersCustom signerNames can also be specified. The signer defines: 1. Trust distribution: how trust (CA bundles) are distributed. 2. Permitted subjects: and behavior when a disallowed subject is requested. 3. Required, permitted, or forbidden x509 extensions in the request (including whether subjectAltNames are allowed, which types, restrictions on allowed values) and behavior when a disallowed extension is requested. 4. Required, permitted, or forbidden key usages / extended key usages. 5. Expiration/certificate lifetime: whether it is fixed by the signer, configurable by the admin. 6. Whether or not requests for CA certificates are allowed.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"uid": schema.StringAttribute{
						Description:         "uid contains the uid of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.",
						MarkdownDescription: "uid contains the uid of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"usages": schema.ListAttribute{
						Description:         "usages specifies a set of key usages requested in the issued certificate.Requests for TLS client certificates typically request: 'digital signature', 'key encipherment', 'client auth'.Requests for TLS serving certificates typically request: 'key encipherment', 'digital signature', 'server auth'.Valid values are: 'signing', 'digital signature', 'content commitment', 'key encipherment', 'key agreement', 'data encipherment', 'cert sign', 'crl sign', 'encipher only', 'decipher only', 'any', 'server auth', 'client auth', 'code signing', 'email protection', 's/mime', 'ipsec end system', 'ipsec tunnel', 'ipsec user', 'timestamping', 'ocsp signing', 'microsoft sgc', 'netscape sgc'",
						MarkdownDescription: "usages specifies a set of key usages requested in the issued certificate.Requests for TLS client certificates typically request: 'digital signature', 'key encipherment', 'client auth'.Requests for TLS serving certificates typically request: 'key encipherment', 'digital signature', 'server auth'.Valid values are: 'signing', 'digital signature', 'content commitment', 'key encipherment', 'key agreement', 'data encipherment', 'cert sign', 'crl sign', 'encipher only', 'decipher only', 'any', 'server auth', 'client auth', 'code signing', 'email protection', 's/mime', 'ipsec end system', 'ipsec tunnel', 'ipsec user', 'timestamping', 'ocsp signing', 'microsoft sgc', 'netscape sgc'",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"username": schema.StringAttribute{
						Description:         "username contains the name of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.",
						MarkdownDescription: "username contains the name of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.",
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

func (r *CertificatesK8SIoCertificateSigningRequestV1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *dynamic.DynamicClient, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *CertificatesK8SIoCertificateSigningRequestV1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_certificates_k8s_io_certificate_signing_request_v1")

	var model CertificatesK8SIoCertificateSigningRequestV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("certificates.k8s.io/v1")
	model.Kind = pointer.String("CertificateSigningRequest")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "certificates.k8s.io", Version: "v1", Resource: "CertificateSigningRequest"}).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while creating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse CertificatesK8SIoCertificateSigningRequestV1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *CertificatesK8SIoCertificateSigningRequestV1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_certificates_k8s_io_certificate_signing_request_v1")

	var data CertificatesK8SIoCertificateSigningRequestV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "certificates.k8s.io", Version: "v1", Resource: "CertificateSigningRequest"}).
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

	var readResponse CertificatesK8SIoCertificateSigningRequestV1ResourceData
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

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *CertificatesK8SIoCertificateSigningRequestV1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_certificates_k8s_io_certificate_signing_request_v1")

	var model CertificatesK8SIoCertificateSigningRequestV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("certificates.k8s.io/v1")
	model.Kind = pointer.String("CertificateSigningRequest")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "certificates.k8s.io", Version: "v1", Resource: "CertificateSigningRequest"}).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while updating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse CertificatesK8SIoCertificateSigningRequestV1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *CertificatesK8SIoCertificateSigningRequestV1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_certificates_k8s_io_certificate_signing_request_v1")

	var data CertificatesK8SIoCertificateSigningRequestV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "certificates.k8s.io", Version: "v1", Resource: "CertificateSigningRequest"}).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to DELETE resource",
			"An unexpected error occurred while deleting the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"DELETE Error: "+err.Error(),
		)
		return
	}
}

func (r *CertificatesK8SIoCertificateSigningRequestV1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	if request.ID == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'name' Got: '%q'", request.ID),
		)
		return
	}
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	resource.ImportStatePassthroughID(ctx, path.Root("metadata").AtName("name"), request, response)
}
