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
	"strings"
)

var (
	_ resource.Resource                = &HiveOpenshiftIoClusterDeprovisionV1Resource{}
	_ resource.ResourceWithConfigure   = &HiveOpenshiftIoClusterDeprovisionV1Resource{}
	_ resource.ResourceWithImportState = &HiveOpenshiftIoClusterDeprovisionV1Resource{}
)

func NewHiveOpenshiftIoClusterDeprovisionV1Resource() resource.Resource {
	return &HiveOpenshiftIoClusterDeprovisionV1Resource{}
}

type HiveOpenshiftIoClusterDeprovisionV1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type HiveOpenshiftIoClusterDeprovisionV1ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitFor        types.List   `tfsdk:"wait_for" json:"-"`

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

func (r *HiveOpenshiftIoClusterDeprovisionV1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_hive_openshift_io_cluster_deprovision_v1"
}

func (r *HiveOpenshiftIoClusterDeprovisionV1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
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
				Description:         "ClusterDeprovisionSpec defines the desired state of ClusterDeprovision",
				MarkdownDescription: "ClusterDeprovisionSpec defines the desired state of ClusterDeprovision",
				Attributes: map[string]schema.Attribute{
					"base_domain": schema.StringAttribute{
						Description:         "BaseDomain is the DNS base domain.",
						MarkdownDescription: "BaseDomain is the DNS base domain.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cluster_id": schema.StringAttribute{
						Description:         "ClusterID is a globally unique identifier for the cluster to deprovision. It will be used if specified.",
						MarkdownDescription: "ClusterID is a globally unique identifier for the cluster to deprovision. It will be used if specified.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cluster_name": schema.StringAttribute{
						Description:         "ClusterName is the friendly name of the cluster. It is used for subdomains, some resource tagging, and other instances where a friendly name for the cluster is useful.",
						MarkdownDescription: "ClusterName is the friendly name of the cluster. It is used for subdomains, some resource tagging, and other instances where a friendly name for the cluster is useful.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"infra_id": schema.StringAttribute{
						Description:         "InfraID is the identifier generated during installation for a cluster. It is used for tagging/naming resources in cloud providers.",
						MarkdownDescription: "InfraID is the identifier generated during installation for a cluster. It is used for tagging/naming resources in cloud providers.",
						Required:            true,
						Optional:            false,
						Computed:            false,
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
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"credentials_secret_ref": schema.SingleNestedAttribute{
										Description:         "CredentialsSecretRef is the Alibaba account credentials to use for deprovisioning the cluster",
										MarkdownDescription: "CredentialsSecretRef is the Alibaba account credentials to use for deprovisioning the cluster",
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
										Description:         "Region is the Alibaba region for this deprovision",
										MarkdownDescription: "Region is the Alibaba region for this deprovision",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
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
												Optional:            true,
												Computed:            false,
											},

											"role_arn": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"credentials_secret_ref": schema.SingleNestedAttribute{
										Description:         "CredentialsSecretRef is the AWS account credentials to use for deprovisioning the cluster",
										MarkdownDescription: "CredentialsSecretRef is the AWS account credentials to use for deprovisioning the cluster",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"hosted_zone_role": schema.StringAttribute{
										Description:         "HostedZoneRole is the role to assume when performing operations on a hosted zone owned by another account.",
										MarkdownDescription: "HostedZoneRole is the role to assume when performing operations on a hosted zone owned by another account.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"region": schema.StringAttribute{
										Description:         "Region is the AWS region for this deprovisioning",
										MarkdownDescription: "Region is the AWS region for this deprovisioning",
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
								Description:         "Azure contains Azure-specific deprovision settings",
								MarkdownDescription: "Azure contains Azure-specific deprovision settings",
								Attributes: map[string]schema.Attribute{
									"cloud_name": schema.StringAttribute{
										Description:         "cloudName is the name of the Azure cloud environment which can be used to configure the Azure SDK with the appropriate Azure API endpoints. If empty, the value is equal to 'AzurePublicCloud'.",
										MarkdownDescription: "cloudName is the name of the Azure cloud environment which can be used to configure the Azure SDK with the appropriate Azure API endpoints. If empty, the value is equal to 'AzurePublicCloud'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("", "AzurePublicCloud", "AzureUSGovernmentCloud", "AzureChinaCloud", "AzureGermanCloud"),
										},
									},

									"credentials_secret_ref": schema.SingleNestedAttribute{
										Description:         "CredentialsSecretRef is the Azure account credentials to use for deprovisioning the cluster",
										MarkdownDescription: "CredentialsSecretRef is the Azure account credentials to use for deprovisioning the cluster",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"resource_group_name": schema.StringAttribute{
										Description:         "ResourceGroupName is the name of the resource group where the cluster was installed. Required for new deprovisions (schema notwithstanding).",
										MarkdownDescription: "ResourceGroupName is the name of the resource group where the cluster was installed. Required for new deprovisions (schema notwithstanding).",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
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
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"region": schema.StringAttribute{
										Description:         "Region is the GCP region for this deprovision",
										MarkdownDescription: "Region is the GCP region for this deprovision",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"ibmcloud": schema.SingleNestedAttribute{
								Description:         "IBMCloud contains IBM Cloud specific deprovision settings",
								MarkdownDescription: "IBMCloud contains IBM Cloud specific deprovision settings",
								Attributes: map[string]schema.Attribute{
									"base_domain": schema.StringAttribute{
										Description:         "BaseDomain is the DNS base domain. TODO: Use the non-platform-specific BaseDomain field.",
										MarkdownDescription: "BaseDomain is the DNS base domain. TODO: Use the non-platform-specific BaseDomain field.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"credentials_secret_ref": schema.SingleNestedAttribute{
										Description:         "CredentialsSecretRef is the IBM Cloud credentials to use for deprovisioning the cluster",
										MarkdownDescription: "CredentialsSecretRef is the IBM Cloud credentials to use for deprovisioning the cluster",
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
										Description:         "Region specifies the IBM Cloud region",
										MarkdownDescription: "Region specifies the IBM Cloud region",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
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
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"cloud": schema.StringAttribute{
										Description:         "Cloud is the secion in the clouds.yaml secret below to use for auth/connectivity.",
										MarkdownDescription: "Cloud is the secion in the clouds.yaml secret below to use for auth/connectivity.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"credentials_secret_ref": schema.SingleNestedAttribute{
										Description:         "CredentialsSecretRef is the OpenStack account credentials to use for deprovisioning the cluster",
										MarkdownDescription: "CredentialsSecretRef is the OpenStack account credentials to use for deprovisioning the cluster",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
								Required: false,
								Optional: true,
								Computed: false,
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
												Optional:            true,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"cluster_id": schema.StringAttribute{
										Description:         "The oVirt cluster ID",
										MarkdownDescription: "The oVirt cluster ID",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"credentials_secret_ref": schema.SingleNestedAttribute{
										Description:         "CredentialsSecretRef is the oVirt account credentials to use for deprovisioning the cluster secret fields: ovirt_url, ovirt_username, ovirt_password, ovirt_ca_bundle",
										MarkdownDescription: "CredentialsSecretRef is the oVirt account credentials to use for deprovisioning the cluster secret fields: ovirt_url, ovirt_username, ovirt_password, ovirt_ca_bundle",
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
												Optional:            true,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"credentials_secret_ref": schema.SingleNestedAttribute{
										Description:         "CredentialsSecretRef is the vSphere account credentials to use for deprovisioning the cluster",
										MarkdownDescription: "CredentialsSecretRef is the vSphere account credentials to use for deprovisioning the cluster",
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

									"v_center": schema.StringAttribute{
										Description:         "VCenter is the vSphere vCenter hostname.",
										MarkdownDescription: "VCenter is the vSphere vCenter hostname.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *HiveOpenshiftIoClusterDeprovisionV1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
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

func (r *HiveOpenshiftIoClusterDeprovisionV1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_hive_openshift_io_cluster_deprovision_v1")

	var model HiveOpenshiftIoClusterDeprovisionV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("hive.openshift.io/v1")
	model.Kind = pointer.String("ClusterDeprovision")

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

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "hive.openshift.io", Version: "v1", Resource: "clusterdeprovisions"}).
		Namespace(model.Metadata.Namespace).
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

	var readResponse HiveOpenshiftIoClusterDeprovisionV1ResourceData
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

func (r *HiveOpenshiftIoClusterDeprovisionV1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_hive_openshift_io_cluster_deprovision_v1")

	var data HiveOpenshiftIoClusterDeprovisionV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "hive.openshift.io", Version: "v1", Resource: "clusterdeprovisions"}).
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

	var readResponse HiveOpenshiftIoClusterDeprovisionV1ResourceData
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

func (r *HiveOpenshiftIoClusterDeprovisionV1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_hive_openshift_io_cluster_deprovision_v1")

	var model HiveOpenshiftIoClusterDeprovisionV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("hive.openshift.io/v1")
	model.Kind = pointer.String("ClusterDeprovision")

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

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "hive.openshift.io", Version: "v1", Resource: "clusterdeprovisions"}).
		Namespace(model.Metadata.Namespace).
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

	var readResponse HiveOpenshiftIoClusterDeprovisionV1ResourceData
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

func (r *HiveOpenshiftIoClusterDeprovisionV1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_hive_openshift_io_cluster_deprovision_v1")

	var data HiveOpenshiftIoClusterDeprovisionV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "hive.openshift.io", Version: "v1", Resource: "clusterdeprovisions"}).
		Namespace(data.Metadata.Namespace).
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

func (r *HiveOpenshiftIoClusterDeprovisionV1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	idParts := strings.Split(request.ID, "/")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'namespace/name' Got: '%q'", request.ID),
		)
		return
	}

	namespace := idParts[0]
	name := idParts[1]
	tflog.Trace(ctx, "parsed import ID", map[string]interface{}{
		"namespace": namespace,
		"name":      name,
	})
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("namespace"), namespace)...)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("name"), name)...)
}
