/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package hive_openshift_io_v1

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
	_ datasource.DataSource              = &HiveOpenshiftIoClusterDeprovisionV1DataSource{}
	_ datasource.DataSourceWithConfigure = &HiveOpenshiftIoClusterDeprovisionV1DataSource{}
)

func NewHiveOpenshiftIoClusterDeprovisionV1DataSource() datasource.DataSource {
	return &HiveOpenshiftIoClusterDeprovisionV1DataSource{}
}

type HiveOpenshiftIoClusterDeprovisionV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type HiveOpenshiftIoClusterDeprovisionV1DataSourceData struct {
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
		BaseDomain  *string `tfsdk:"base_domain" json:"baseDomain,omitempty"`
		ClusterID   *string `tfsdk:"cluster_id" json:"clusterID,omitempty"`
		ClusterName *string `tfsdk:"cluster_name" json:"clusterName,omitempty"`
		InfraID     *string `tfsdk:"infra_id" json:"infraID,omitempty"`
		Platform    *struct {
			Alibabacloud *struct {
				BaseDomain           *string `tfsdk:"base_domain" json:"baseDomain,omitempty"`
				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" json:"credentialsSecretRef,omitempty"`
				Region *string `tfsdk:"region" json:"region,omitempty"`
			} `tfsdk:"alibabacloud" json:"alibabacloud,omitempty"`
			Aws *struct {
				CredentialsAssumeRole *struct {
					ExternalID *string `tfsdk:"external_id" json:"externalID,omitempty"`
					RoleARN    *string `tfsdk:"role_arn" json:"roleARN,omitempty"`
				} `tfsdk:"credentials_assume_role" json:"credentialsAssumeRole,omitempty"`
				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" json:"credentialsSecretRef,omitempty"`
				HostedZoneRole *string `tfsdk:"hosted_zone_role" json:"hostedZoneRole,omitempty"`
				Region         *string `tfsdk:"region" json:"region,omitempty"`
			} `tfsdk:"aws" json:"aws,omitempty"`
			Azure *struct {
				CloudName            *string `tfsdk:"cloud_name" json:"cloudName,omitempty"`
				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" json:"credentialsSecretRef,omitempty"`
				ResourceGroupName *string `tfsdk:"resource_group_name" json:"resourceGroupName,omitempty"`
			} `tfsdk:"azure" json:"azure,omitempty"`
			Gcp *struct {
				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" json:"credentialsSecretRef,omitempty"`
				Region *string `tfsdk:"region" json:"region,omitempty"`
			} `tfsdk:"gcp" json:"gcp,omitempty"`
			Ibmcloud *struct {
				BaseDomain           *string `tfsdk:"base_domain" json:"baseDomain,omitempty"`
				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" json:"credentialsSecretRef,omitempty"`
				Region *string `tfsdk:"region" json:"region,omitempty"`
			} `tfsdk:"ibmcloud" json:"ibmcloud,omitempty"`
			Openstack *struct {
				CertificatesSecretRef *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"certificates_secret_ref" json:"certificatesSecretRef,omitempty"`
				Cloud                *string `tfsdk:"cloud" json:"cloud,omitempty"`
				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" json:"credentialsSecretRef,omitempty"`
			} `tfsdk:"openstack" json:"openstack,omitempty"`
			Ovirt *struct {
				CertificatesSecretRef *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"certificates_secret_ref" json:"certificatesSecretRef,omitempty"`
				ClusterID            *string `tfsdk:"cluster_id" json:"clusterID,omitempty"`
				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" json:"credentialsSecretRef,omitempty"`
			} `tfsdk:"ovirt" json:"ovirt,omitempty"`
			Vsphere *struct {
				CertificatesSecretRef *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"certificates_secret_ref" json:"certificatesSecretRef,omitempty"`
				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" json:"credentialsSecretRef,omitempty"`
				VCenter *string `tfsdk:"v_center" json:"vCenter,omitempty"`
			} `tfsdk:"vsphere" json:"vsphere,omitempty"`
		} `tfsdk:"platform" json:"platform,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *HiveOpenshiftIoClusterDeprovisionV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_hive_openshift_io_cluster_deprovision_v1"
}

func (r *HiveOpenshiftIoClusterDeprovisionV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClusterDeprovision is the Schema for the clusterdeprovisions API",
		MarkdownDescription: "ClusterDeprovision is the Schema for the clusterdeprovisions API",
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
				Description:         "ClusterDeprovisionSpec defines the desired state of ClusterDeprovision",
				MarkdownDescription: "ClusterDeprovisionSpec defines the desired state of ClusterDeprovision",
				Attributes: map[string]schema.Attribute{
					"base_domain": schema.StringAttribute{
						Description:         "BaseDomain is the DNS base domain.",
						MarkdownDescription: "BaseDomain is the DNS base domain.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"cluster_id": schema.StringAttribute{
						Description:         "ClusterID is a globally unique identifier for the cluster to deprovision. It will be used if specified.",
						MarkdownDescription: "ClusterID is a globally unique identifier for the cluster to deprovision. It will be used if specified.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"cluster_name": schema.StringAttribute{
						Description:         "ClusterName is the friendly name of the cluster. It is used for subdomains, some resource tagging, and other instances where a friendly name for the cluster is useful.",
						MarkdownDescription: "ClusterName is the friendly name of the cluster. It is used for subdomains, some resource tagging, and other instances where a friendly name for the cluster is useful.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"infra_id": schema.StringAttribute{
						Description:         "InfraID is the identifier generated during installation for a cluster. It is used for tagging/naming resources in cloud providers.",
						MarkdownDescription: "InfraID is the identifier generated during installation for a cluster. It is used for tagging/naming resources in cloud providers.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"platform": schema.SingleNestedAttribute{
						Description:         "Platform contains platform-specific configuration for a ClusterDeprovision",
						MarkdownDescription: "Platform contains platform-specific configuration for a ClusterDeprovision",
						Attributes: map[string]schema.Attribute{
							"alibabacloud": schema.SingleNestedAttribute{
								Description:         "AlibabaCloud contains Alibaba Cloud specific deprovision settings",
								MarkdownDescription: "AlibabaCloud contains Alibaba Cloud specific deprovision settings",
								Attributes: map[string]schema.Attribute{
									"base_domain": schema.StringAttribute{
										Description:         "BaseDomain is the DNS base domain. TODO: Use the non-platform-specific BaseDomain field.",
										MarkdownDescription: "BaseDomain is the DNS base domain. TODO: Use the non-platform-specific BaseDomain field.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"credentials_secret_ref": schema.SingleNestedAttribute{
										Description:         "CredentialsSecretRef is the Alibaba account credentials to use for deprovisioning the cluster",
										MarkdownDescription: "CredentialsSecretRef is the Alibaba account credentials to use for deprovisioning the cluster",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"region": schema.StringAttribute{
										Description:         "Region is the Alibaba region for this deprovision",
										MarkdownDescription: "Region is the Alibaba region for this deprovision",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"aws": schema.SingleNestedAttribute{
								Description:         "AWS contains AWS-specific deprovision settings",
								MarkdownDescription: "AWS contains AWS-specific deprovision settings",
								Attributes: map[string]schema.Attribute{
									"credentials_assume_role": schema.SingleNestedAttribute{
										Description:         "CredentialsAssumeRole refers to the IAM role that must be assumed to obtain AWS account access for deprovisioning the cluster.",
										MarkdownDescription: "CredentialsAssumeRole refers to the IAM role that must be assumed to obtain AWS account access for deprovisioning the cluster.",
										Attributes: map[string]schema.Attribute{
											"external_id": schema.StringAttribute{
												Description:         "ExternalID is random string generated by platform so that assume role is protected from confused deputy problem. more info: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html",
												MarkdownDescription: "ExternalID is random string generated by platform so that assume role is protected from confused deputy problem. more info: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"role_arn": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"credentials_secret_ref": schema.SingleNestedAttribute{
										Description:         "CredentialsSecretRef is the AWS account credentials to use for deprovisioning the cluster",
										MarkdownDescription: "CredentialsSecretRef is the AWS account credentials to use for deprovisioning the cluster",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"hosted_zone_role": schema.StringAttribute{
										Description:         "HostedZoneRole is the role to assume when performing operations on a hosted zone owned by another account.",
										MarkdownDescription: "HostedZoneRole is the role to assume when performing operations on a hosted zone owned by another account.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"region": schema.StringAttribute{
										Description:         "Region is the AWS region for this deprovisioning",
										MarkdownDescription: "Region is the AWS region for this deprovisioning",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"azure": schema.SingleNestedAttribute{
								Description:         "Azure contains Azure-specific deprovision settings",
								MarkdownDescription: "Azure contains Azure-specific deprovision settings",
								Attributes: map[string]schema.Attribute{
									"cloud_name": schema.StringAttribute{
										Description:         "cloudName is the name of the Azure cloud environment which can be used to configure the Azure SDK with the appropriate Azure API endpoints. If empty, the value is equal to 'AzurePublicCloud'.",
										MarkdownDescription: "cloudName is the name of the Azure cloud environment which can be used to configure the Azure SDK with the appropriate Azure API endpoints. If empty, the value is equal to 'AzurePublicCloud'.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"credentials_secret_ref": schema.SingleNestedAttribute{
										Description:         "CredentialsSecretRef is the Azure account credentials to use for deprovisioning the cluster",
										MarkdownDescription: "CredentialsSecretRef is the Azure account credentials to use for deprovisioning the cluster",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"resource_group_name": schema.StringAttribute{
										Description:         "ResourceGroupName is the name of the resource group where the cluster was installed. Required for new deprovisions (schema notwithstanding).",
										MarkdownDescription: "ResourceGroupName is the name of the resource group where the cluster was installed. Required for new deprovisions (schema notwithstanding).",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"gcp": schema.SingleNestedAttribute{
								Description:         "GCP contains GCP-specific deprovision settings",
								MarkdownDescription: "GCP contains GCP-specific deprovision settings",
								Attributes: map[string]schema.Attribute{
									"credentials_secret_ref": schema.SingleNestedAttribute{
										Description:         "CredentialsSecretRef is the GCP account credentials to use for deprovisioning the cluster",
										MarkdownDescription: "CredentialsSecretRef is the GCP account credentials to use for deprovisioning the cluster",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"region": schema.StringAttribute{
										Description:         "Region is the GCP region for this deprovision",
										MarkdownDescription: "Region is the GCP region for this deprovision",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"ibmcloud": schema.SingleNestedAttribute{
								Description:         "IBMCloud contains IBM Cloud specific deprovision settings",
								MarkdownDescription: "IBMCloud contains IBM Cloud specific deprovision settings",
								Attributes: map[string]schema.Attribute{
									"base_domain": schema.StringAttribute{
										Description:         "BaseDomain is the DNS base domain. TODO: Use the non-platform-specific BaseDomain field.",
										MarkdownDescription: "BaseDomain is the DNS base domain. TODO: Use the non-platform-specific BaseDomain field.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"credentials_secret_ref": schema.SingleNestedAttribute{
										Description:         "CredentialsSecretRef is the IBM Cloud credentials to use for deprovisioning the cluster",
										MarkdownDescription: "CredentialsSecretRef is the IBM Cloud credentials to use for deprovisioning the cluster",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"region": schema.StringAttribute{
										Description:         "Region specifies the IBM Cloud region",
										MarkdownDescription: "Region specifies the IBM Cloud region",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"openstack": schema.SingleNestedAttribute{
								Description:         "OpenStack contains OpenStack-specific deprovision settings",
								MarkdownDescription: "OpenStack contains OpenStack-specific deprovision settings",
								Attributes: map[string]schema.Attribute{
									"certificates_secret_ref": schema.SingleNestedAttribute{
										Description:         "CertificatesSecretRef refers to a secret that contains CA certificates necessary for communicating with the OpenStack.",
										MarkdownDescription: "CertificatesSecretRef refers to a secret that contains CA certificates necessary for communicating with the OpenStack.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"cloud": schema.StringAttribute{
										Description:         "Cloud is the secion in the clouds.yaml secret below to use for auth/connectivity.",
										MarkdownDescription: "Cloud is the secion in the clouds.yaml secret below to use for auth/connectivity.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"credentials_secret_ref": schema.SingleNestedAttribute{
										Description:         "CredentialsSecretRef is the OpenStack account credentials to use for deprovisioning the cluster",
										MarkdownDescription: "CredentialsSecretRef is the OpenStack account credentials to use for deprovisioning the cluster",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
								Required: false,
								Optional: false,
								Computed: true,
							},

							"ovirt": schema.SingleNestedAttribute{
								Description:         "Ovirt contains oVirt-specific deprovision settings",
								MarkdownDescription: "Ovirt contains oVirt-specific deprovision settings",
								Attributes: map[string]schema.Attribute{
									"certificates_secret_ref": schema.SingleNestedAttribute{
										Description:         "CertificatesSecretRef refers to a secret that contains the oVirt CA certificates necessary for communicating with the oVirt.",
										MarkdownDescription: "CertificatesSecretRef refers to a secret that contains the oVirt CA certificates necessary for communicating with the oVirt.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"cluster_id": schema.StringAttribute{
										Description:         "The oVirt cluster ID",
										MarkdownDescription: "The oVirt cluster ID",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"credentials_secret_ref": schema.SingleNestedAttribute{
										Description:         "CredentialsSecretRef is the oVirt account credentials to use for deprovisioning the cluster secret fields: ovirt_url, ovirt_username, ovirt_password, ovirt_ca_bundle",
										MarkdownDescription: "CredentialsSecretRef is the oVirt account credentials to use for deprovisioning the cluster secret fields: ovirt_url, ovirt_username, ovirt_password, ovirt_ca_bundle",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
								Required: false,
								Optional: false,
								Computed: true,
							},

							"vsphere": schema.SingleNestedAttribute{
								Description:         "VSphere contains VMWare vSphere-specific deprovision settings",
								MarkdownDescription: "VSphere contains VMWare vSphere-specific deprovision settings",
								Attributes: map[string]schema.Attribute{
									"certificates_secret_ref": schema.SingleNestedAttribute{
										Description:         "CertificatesSecretRef refers to a secret that contains the vSphere CA certificates necessary for communicating with the VCenter.",
										MarkdownDescription: "CertificatesSecretRef refers to a secret that contains the vSphere CA certificates necessary for communicating with the VCenter.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"credentials_secret_ref": schema.SingleNestedAttribute{
										Description:         "CredentialsSecretRef is the vSphere account credentials to use for deprovisioning the cluster",
										MarkdownDescription: "CredentialsSecretRef is the vSphere account credentials to use for deprovisioning the cluster",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"v_center": schema.StringAttribute{
										Description:         "VCenter is the vSphere vCenter hostname.",
										MarkdownDescription: "VCenter is the vSphere vCenter hostname.",
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
						Required: false,
						Optional: false,
						Computed: true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *HiveOpenshiftIoClusterDeprovisionV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *HiveOpenshiftIoClusterDeprovisionV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_hive_openshift_io_cluster_deprovision_v1")

	var data HiveOpenshiftIoClusterDeprovisionV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "hive.openshift.io", Version: "v1", Resource: "ClusterDeprovision"}).
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

	var readResponse HiveOpenshiftIoClusterDeprovisionV1DataSourceData
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
	data.ApiVersion = pointer.String("hive.openshift.io/v1")
	data.Kind = pointer.String("ClusterDeprovision")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
