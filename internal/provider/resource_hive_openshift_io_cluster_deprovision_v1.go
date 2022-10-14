/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

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

type HiveOpenshiftIoClusterDeprovisionV1Resource struct{}

var (
	_ resource.Resource = (*HiveOpenshiftIoClusterDeprovisionV1Resource)(nil)
)

type HiveOpenshiftIoClusterDeprovisionV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type HiveOpenshiftIoClusterDeprovisionV1GoModel struct {
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
		ClusterID *string `tfsdk:"cluster_id" yaml:"clusterID,omitempty"`

		ClusterName *string `tfsdk:"cluster_name" yaml:"clusterName,omitempty"`

		InfraID *string `tfsdk:"infra_id" yaml:"infraID,omitempty"`

		Platform *struct {
			Alibabacloud *struct {
				BaseDomain *string `tfsdk:"base_domain" yaml:"baseDomain,omitempty"`

				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" yaml:"credentialsSecretRef,omitempty"`

				Region *string `tfsdk:"region" yaml:"region,omitempty"`
			} `tfsdk:"alibabacloud" yaml:"alibabacloud,omitempty"`

			Aws *struct {
				CredentialsAssumeRole *struct {
					ExternalID *string `tfsdk:"external_id" yaml:"externalID,omitempty"`

					RoleARN *string `tfsdk:"role_arn" yaml:"roleARN,omitempty"`
				} `tfsdk:"credentials_assume_role" yaml:"credentialsAssumeRole,omitempty"`

				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" yaml:"credentialsSecretRef,omitempty"`

				Region *string `tfsdk:"region" yaml:"region,omitempty"`
			} `tfsdk:"aws" yaml:"aws,omitempty"`

			Azure *struct {
				CloudName *string `tfsdk:"cloud_name" yaml:"cloudName,omitempty"`

				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" yaml:"credentialsSecretRef,omitempty"`
			} `tfsdk:"azure" yaml:"azure,omitempty"`

			Gcp *struct {
				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" yaml:"credentialsSecretRef,omitempty"`

				Region *string `tfsdk:"region" yaml:"region,omitempty"`
			} `tfsdk:"gcp" yaml:"gcp,omitempty"`

			Ibmcloud *struct {
				BaseDomain *string `tfsdk:"base_domain" yaml:"baseDomain,omitempty"`

				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" yaml:"credentialsSecretRef,omitempty"`

				Region *string `tfsdk:"region" yaml:"region,omitempty"`
			} `tfsdk:"ibmcloud" yaml:"ibmcloud,omitempty"`

			Openstack *struct {
				CertificatesSecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"certificates_secret_ref" yaml:"certificatesSecretRef,omitempty"`

				Cloud *string `tfsdk:"cloud" yaml:"cloud,omitempty"`

				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" yaml:"credentialsSecretRef,omitempty"`
			} `tfsdk:"openstack" yaml:"openstack,omitempty"`

			Ovirt *struct {
				CertificatesSecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"certificates_secret_ref" yaml:"certificatesSecretRef,omitempty"`

				ClusterID *string `tfsdk:"cluster_id" yaml:"clusterID,omitempty"`

				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" yaml:"credentialsSecretRef,omitempty"`
			} `tfsdk:"ovirt" yaml:"ovirt,omitempty"`

			Vsphere *struct {
				CertificatesSecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"certificates_secret_ref" yaml:"certificatesSecretRef,omitempty"`

				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" yaml:"credentialsSecretRef,omitempty"`

				VCenter *string `tfsdk:"v_center" yaml:"vCenter,omitempty"`
			} `tfsdk:"vsphere" yaml:"vsphere,omitempty"`
		} `tfsdk:"platform" yaml:"platform,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewHiveOpenshiftIoClusterDeprovisionV1Resource() resource.Resource {
	return &HiveOpenshiftIoClusterDeprovisionV1Resource{}
}

func (r *HiveOpenshiftIoClusterDeprovisionV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hive_openshift_io_cluster_deprovision_v1"
}

func (r *HiveOpenshiftIoClusterDeprovisionV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "ClusterDeprovision is the Schema for the clusterdeprovisions API",
		MarkdownDescription: "ClusterDeprovision is the Schema for the clusterdeprovisions API",
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
				Description:         "ClusterDeprovisionSpec defines the desired state of ClusterDeprovision",
				MarkdownDescription: "ClusterDeprovisionSpec defines the desired state of ClusterDeprovision",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"cluster_id": {
						Description:         "ClusterID is a globally unique identifier for the cluster to deprovision. It will be used if specified.",
						MarkdownDescription: "ClusterID is a globally unique identifier for the cluster to deprovision. It will be used if specified.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"cluster_name": {
						Description:         "ClusterName is the friendly name of the cluster. It is used for subdomains, some resource tagging, and other instances where a friendly name for the cluster is useful.",
						MarkdownDescription: "ClusterName is the friendly name of the cluster. It is used for subdomains, some resource tagging, and other instances where a friendly name for the cluster is useful.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"infra_id": {
						Description:         "InfraID is the identifier generated during installation for a cluster. It is used for tagging/naming resources in cloud providers.",
						MarkdownDescription: "InfraID is the identifier generated during installation for a cluster. It is used for tagging/naming resources in cloud providers.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"platform": {
						Description:         "Platform contains platform-specific configuration for a ClusterDeprovision",
						MarkdownDescription: "Platform contains platform-specific configuration for a ClusterDeprovision",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"alibabacloud": {
								Description:         "AlibabaCloud contains Alibaba Cloud specific deprovision settings",
								MarkdownDescription: "AlibabaCloud contains Alibaba Cloud specific deprovision settings",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"base_domain": {
										Description:         "BaseDomain is the DNS base domain",
										MarkdownDescription: "BaseDomain is the DNS base domain",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"credentials_secret_ref": {
										Description:         "CredentialsSecretRef is the Alibaba account credentials to use for deprovisioning the cluster",
										MarkdownDescription: "CredentialsSecretRef is the Alibaba account credentials to use for deprovisioning the cluster",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

									"region": {
										Description:         "Region is the Alibaba region for this deprovision",
										MarkdownDescription: "Region is the Alibaba region for this deprovision",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"aws": {
								Description:         "AWS contains AWS-specific deprovision settings",
								MarkdownDescription: "AWS contains AWS-specific deprovision settings",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"credentials_assume_role": {
										Description:         "CredentialsAssumeRole refers to the IAM role that must be assumed to obtain AWS account access for deprovisioning the cluster.",
										MarkdownDescription: "CredentialsAssumeRole refers to the IAM role that must be assumed to obtain AWS account access for deprovisioning the cluster.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"external_id": {
												Description:         "ExternalID is random string generated by platform so that assume role is protected from confused deputy problem. more info: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html",
												MarkdownDescription: "ExternalID is random string generated by platform so that assume role is protected from confused deputy problem. more info: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"role_arn": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"credentials_secret_ref": {
										Description:         "CredentialsSecretRef is the AWS account credentials to use for deprovisioning the cluster",
										MarkdownDescription: "CredentialsSecretRef is the AWS account credentials to use for deprovisioning the cluster",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"region": {
										Description:         "Region is the AWS region for this deprovisioning",
										MarkdownDescription: "Region is the AWS region for this deprovisioning",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"azure": {
								Description:         "Azure contains Azure-specific deprovision settings",
								MarkdownDescription: "Azure contains Azure-specific deprovision settings",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"cloud_name": {
										Description:         "cloudName is the name of the Azure cloud environment which can be used to configure the Azure SDK with the appropriate Azure API endpoints. If empty, the value is equal to 'AzurePublicCloud'.",
										MarkdownDescription: "cloudName is the name of the Azure cloud environment which can be used to configure the Azure SDK with the appropriate Azure API endpoints. If empty, the value is equal to 'AzurePublicCloud'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("", "AzurePublicCloud", "AzureUSGovernmentCloud", "AzureChinaCloud", "AzureGermanCloud"),
										},
									},

									"credentials_secret_ref": {
										Description:         "CredentialsSecretRef is the Azure account credentials to use for deprovisioning the cluster",
										MarkdownDescription: "CredentialsSecretRef is the Azure account credentials to use for deprovisioning the cluster",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

												Type: types.StringType,

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

							"gcp": {
								Description:         "GCP contains GCP-specific deprovision settings",
								MarkdownDescription: "GCP contains GCP-specific deprovision settings",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"credentials_secret_ref": {
										Description:         "CredentialsSecretRef is the GCP account credentials to use for deprovisioning the cluster",
										MarkdownDescription: "CredentialsSecretRef is the GCP account credentials to use for deprovisioning the cluster",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"region": {
										Description:         "Region is the GCP region for this deprovision",
										MarkdownDescription: "Region is the GCP region for this deprovision",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ibmcloud": {
								Description:         "IBMCloud contains IBM Cloud specific deprovision settings",
								MarkdownDescription: "IBMCloud contains IBM Cloud specific deprovision settings",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"base_domain": {
										Description:         "BaseDomain is the DNS base domain",
										MarkdownDescription: "BaseDomain is the DNS base domain",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"credentials_secret_ref": {
										Description:         "CredentialsSecretRef is the IBM Cloud credentials to use for deprovisioning the cluster",
										MarkdownDescription: "CredentialsSecretRef is the IBM Cloud credentials to use for deprovisioning the cluster",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

									"region": {
										Description:         "Region specifies the IBM Cloud region",
										MarkdownDescription: "Region specifies the IBM Cloud region",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"openstack": {
								Description:         "OpenStack contains OpenStack-specific deprovision settings",
								MarkdownDescription: "OpenStack contains OpenStack-specific deprovision settings",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"certificates_secret_ref": {
										Description:         "CertificatesSecretRef refers to a secret that contains CA certificates necessary for communicating with the OpenStack.",
										MarkdownDescription: "CertificatesSecretRef refers to a secret that contains CA certificates necessary for communicating with the OpenStack.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"cloud": {
										Description:         "Cloud is the secion in the clouds.yaml secret below to use for auth/connectivity.",
										MarkdownDescription: "Cloud is the secion in the clouds.yaml secret below to use for auth/connectivity.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"credentials_secret_ref": {
										Description:         "CredentialsSecretRef is the OpenStack account credentials to use for deprovisioning the cluster",
										MarkdownDescription: "CredentialsSecretRef is the OpenStack account credentials to use for deprovisioning the cluster",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

												Type: types.StringType,

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

							"ovirt": {
								Description:         "Ovirt contains oVirt-specific deprovision settings",
								MarkdownDescription: "Ovirt contains oVirt-specific deprovision settings",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"certificates_secret_ref": {
										Description:         "CertificatesSecretRef refers to a secret that contains the oVirt CA certificates necessary for communicating with the oVirt.",
										MarkdownDescription: "CertificatesSecretRef refers to a secret that contains the oVirt CA certificates necessary for communicating with the oVirt.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

									"cluster_id": {
										Description:         "The oVirt cluster ID",
										MarkdownDescription: "The oVirt cluster ID",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"credentials_secret_ref": {
										Description:         "CredentialsSecretRef is the oVirt account credentials to use for deprovisioning the cluster secret fields: ovirt_url, ovirt_username, ovirt_password, ovirt_ca_bundle",
										MarkdownDescription: "CredentialsSecretRef is the oVirt account credentials to use for deprovisioning the cluster secret fields: ovirt_url, ovirt_username, ovirt_password, ovirt_ca_bundle",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"vsphere": {
								Description:         "VSphere contains VMWare vSphere-specific deprovision settings",
								MarkdownDescription: "VSphere contains VMWare vSphere-specific deprovision settings",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"certificates_secret_ref": {
										Description:         "CertificatesSecretRef refers to a secret that contains the vSphere CA certificates necessary for communicating with the VCenter.",
										MarkdownDescription: "CertificatesSecretRef refers to a secret that contains the vSphere CA certificates necessary for communicating with the VCenter.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

									"credentials_secret_ref": {
										Description:         "CredentialsSecretRef is the vSphere account credentials to use for deprovisioning the cluster",
										MarkdownDescription: "CredentialsSecretRef is the vSphere account credentials to use for deprovisioning the cluster",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

									"v_center": {
										Description:         "VCenter is the vSphere vCenter hostname.",
										MarkdownDescription: "VCenter is the vSphere vCenter hostname.",

										Type: types.StringType,

										Required: true,
										Optional: false,
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

func (r *HiveOpenshiftIoClusterDeprovisionV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_hive_openshift_io_cluster_deprovision_v1")

	var state HiveOpenshiftIoClusterDeprovisionV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel HiveOpenshiftIoClusterDeprovisionV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("hive.openshift.io/v1")
	goModel.Kind = utilities.Ptr("ClusterDeprovision")

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

func (r *HiveOpenshiftIoClusterDeprovisionV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_hive_openshift_io_cluster_deprovision_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *HiveOpenshiftIoClusterDeprovisionV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_hive_openshift_io_cluster_deprovision_v1")

	var state HiveOpenshiftIoClusterDeprovisionV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel HiveOpenshiftIoClusterDeprovisionV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("hive.openshift.io/v1")
	goModel.Kind = utilities.Ptr("ClusterDeprovision")

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

func (r *HiveOpenshiftIoClusterDeprovisionV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_hive_openshift_io_cluster_deprovision_v1")
	// NO-OP: Terraform removes the state automatically for us
}
