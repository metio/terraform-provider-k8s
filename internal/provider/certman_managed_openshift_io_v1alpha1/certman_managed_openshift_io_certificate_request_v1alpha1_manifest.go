/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package certman_managed_openshift_io_v1alpha1

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
	_ datasource.DataSource = &CertmanManagedOpenshiftIoCertificateRequestV1Alpha1Manifest{}
)

func NewCertmanManagedOpenshiftIoCertificateRequestV1Alpha1Manifest() datasource.DataSource {
	return &CertmanManagedOpenshiftIoCertificateRequestV1Alpha1Manifest{}
}

type CertmanManagedOpenshiftIoCertificateRequestV1Alpha1Manifest struct{}

type CertmanManagedOpenshiftIoCertificateRequestV1Alpha1ManifestData struct {
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
		AcmeDNSDomain     *string `tfsdk:"acme_dns_domain" json:"acmeDNSDomain,omitempty"`
		ApiURL            *string `tfsdk:"api_url" json:"apiURL,omitempty"`
		CertificateSecret *struct {
			ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
			Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
			Name            *string `tfsdk:"name" json:"name,omitempty"`
			Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
			ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
			Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
		} `tfsdk:"certificate_secret" json:"certificateSecret,omitempty"`
		DnsNames *[]string `tfsdk:"dns_names" json:"dnsNames,omitempty"`
		Email    *string   `tfsdk:"email" json:"email,omitempty"`
		Platform *struct {
			Aws *struct {
				Credentials *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"credentials" json:"credentials,omitempty"`
				Region *string `tfsdk:"region" json:"region,omitempty"`
			} `tfsdk:"aws" json:"aws,omitempty"`
			Azure *struct {
				Credentials *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"credentials" json:"credentials,omitempty"`
				ResourceGroupName *string `tfsdk:"resource_group_name" json:"resourceGroupName,omitempty"`
			} `tfsdk:"azure" json:"azure,omitempty"`
			Gcp *struct {
				Credentials *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"credentials" json:"credentials,omitempty"`
			} `tfsdk:"gcp" json:"gcp,omitempty"`
			Mock *struct {
				AnswerDNSChallengeErrorString                 *string `tfsdk:"answer_dns_challenge_error_string" json:"answerDNSChallengeErrorString,omitempty"`
				AnswerDNSChallengeFQDN                        *string `tfsdk:"answer_dns_challenge_fqdn" json:"answerDNSChallengeFQDN,omitempty"`
				DeleteAcmeChallengeResourceRecordsErrorString *string `tfsdk:"delete_acme_challenge_resource_records_error_string" json:"deleteAcmeChallengeResourceRecordsErrorString,omitempty"`
				ValidateDNSWriteAccessBool                    *bool   `tfsdk:"validate_dns_write_access_bool" json:"validateDNSWriteAccessBool,omitempty"`
				ValidateDNSWriteAccessErrorString             *string `tfsdk:"validate_dns_write_access_error_string" json:"validateDNSWriteAccessErrorString,omitempty"`
			} `tfsdk:"mock" json:"mock,omitempty"`
		} `tfsdk:"platform" json:"platform,omitempty"`
		RenewBeforeDays *int64  `tfsdk:"renew_before_days" json:"renewBeforeDays,omitempty"`
		WebConsoleURL   *string `tfsdk:"web_console_url" json:"webConsoleURL,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CertmanManagedOpenshiftIoCertificateRequestV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_certman_managed_openshift_io_certificate_request_v1alpha1_manifest"
}

func (r *CertmanManagedOpenshiftIoCertificateRequestV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CertificateRequest is the Schema for the certificaterequests API",
		MarkdownDescription: "CertificateRequest is the Schema for the certificaterequests API",
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
				Description:         "CertificateRequestSpec defines the desired state of CertificateRequest",
				MarkdownDescription: "CertificateRequestSpec defines the desired state of CertificateRequest",
				Attributes: map[string]schema.Attribute{
					"acme_dns_domain": schema.StringAttribute{
						Description:         "ACMEDNSDomain is the DNS zone that will house the TXT records needed for the certificate to be created. In Route53 this would be the public Route53 hosted zone (the Domain Name not the ZoneID)",
						MarkdownDescription: "ACMEDNSDomain is the DNS zone that will house the TXT records needed for the certificate to be created. In Route53 this would be the public Route53 hosted zone (the Domain Name not the ZoneID)",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"api_url": schema.StringAttribute{
						Description:         "APIURL is the URL where the cluster's API can be accessed.",
						MarkdownDescription: "APIURL is the URL where the cluster's API can be accessed.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"certificate_secret": schema.SingleNestedAttribute{
						Description:         "CertificateSecret is the reference to the secret where certificates are stored.",
						MarkdownDescription: "CertificateSecret is the reference to the secret where certificates are stored.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "API version of the referent.",
								MarkdownDescription: "API version of the referent.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"field_path": schema.StringAttribute{
								Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
								MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resource_version": schema.StringAttribute{
								Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"uid": schema.StringAttribute{
								Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"dns_names": schema.ListAttribute{
						Description:         "DNSNames is a list of subject alt names to be used on the Certificate.",
						MarkdownDescription: "DNSNames is a list of subject alt names to be used on the Certificate.",
						ElementType:         types.StringType,
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"email": schema.StringAttribute{
						Description:         "Let's Encrypt will use this to contact you about expiring certificates, and issues related to your account.",
						MarkdownDescription: "Let's Encrypt will use this to contact you about expiring certificates, and issues related to your account.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"platform": schema.SingleNestedAttribute{
						Description:         "Platform contains specific cloud provider information such as credentials and secrets for the cluster infrastructure.",
						MarkdownDescription: "Platform contains specific cloud provider information such as credentials and secrets for the cluster infrastructure.",
						Attributes: map[string]schema.Attribute{
							"aws": schema.SingleNestedAttribute{
								Description:         "AWSPlatformSecrets contains secrets for clusters on the AWS platform.",
								MarkdownDescription: "AWSPlatformSecrets contains secrets for clusters on the AWS platform.",
								Attributes: map[string]schema.Attribute{
									"credentials": schema.SingleNestedAttribute{
										Description:         "Credentials refers to a secret that contains the AWS account access credentials.",
										MarkdownDescription: "Credentials refers to a secret that contains the AWS account access credentials.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"region": schema.StringAttribute{
										Description:         "Region specifies the AWS region where the cluster will be created.",
										MarkdownDescription: "Region specifies the AWS region where the cluster will be created.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"azure": schema.SingleNestedAttribute{
								Description:         "AzurePlatformSecrets contains secrets for clusters on the Azure platform.",
								MarkdownDescription: "AzurePlatformSecrets contains secrets for clusters on the Azure platform.",
								Attributes: map[string]schema.Attribute{
									"credentials": schema.SingleNestedAttribute{
										Description:         "Credentials refers to a secret that contains the AZURE account access credentials.",
										MarkdownDescription: "Credentials refers to a secret that contains the AZURE account access credentials.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"resource_group_name": schema.StringAttribute{
										Description:         "ResourceGroupName refers to the resource group that contains the dns zone.",
										MarkdownDescription: "ResourceGroupName refers to the resource group that contains the dns zone.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"gcp": schema.SingleNestedAttribute{
								Description:         "GCPPlatformSecrets contains secrets for clusters on the GCP platform.",
								MarkdownDescription: "GCPPlatformSecrets contains secrets for clusters on the GCP platform.",
								Attributes: map[string]schema.Attribute{
									"credentials": schema.SingleNestedAttribute{
										Description:         "Credentials refers to a secret that contains the GCP account access credentials.",
										MarkdownDescription: "Credentials refers to a secret that contains the GCP account access credentials.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

							"mock": schema.SingleNestedAttribute{
								Description:         "MockPlatformSecrets indicates a mock client should be generated, which doesn't interact with any platform",
								MarkdownDescription: "MockPlatformSecrets indicates a mock client should be generated, which doesn't interact with any platform",
								Attributes: map[string]schema.Attribute{
									"answer_dns_challenge_error_string": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"answer_dns_challenge_fqdn": schema.StringAttribute{
										Description:         "these options configure the return values for the mock client's functions",
										MarkdownDescription: "these options configure the return values for the mock client's functions",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"delete_acme_challenge_resource_records_error_string": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"validate_dns_write_access_bool": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"validate_dns_write_access_error_string": schema.StringAttribute{
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
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"renew_before_days": schema.Int64Attribute{
						Description:         "Number of days before expiration to reissue certificate. NOTE: Keeping 'renew' in JSON for backward-compatibility.",
						MarkdownDescription: "Number of days before expiration to reissue certificate. NOTE: Keeping 'renew' in JSON for backward-compatibility.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"web_console_url": schema.StringAttribute{
						Description:         "WebConsoleURL is the URL for the cluster's web console UI.",
						MarkdownDescription: "WebConsoleURL is the URL for the cluster's web console UI.",
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

func (r *CertmanManagedOpenshiftIoCertificateRequestV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_certman_managed_openshift_io_certificate_request_v1alpha1_manifest")

	var model CertmanManagedOpenshiftIoCertificateRequestV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("certman.managed.openshift.io/v1alpha1")
	model.Kind = pointer.String("CertificateRequest")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
