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
	_ datasource.DataSource              = &HiveOpenshiftIoClusterPoolV1DataSource{}
	_ datasource.DataSourceWithConfigure = &HiveOpenshiftIoClusterPoolV1DataSource{}
)

func NewHiveOpenshiftIoClusterPoolV1DataSource() datasource.DataSource {
	return &HiveOpenshiftIoClusterPoolV1DataSource{}
}

type HiveOpenshiftIoClusterPoolV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type HiveOpenshiftIoClusterPoolV1DataSourceData struct {
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
		Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
		BaseDomain    *string            `tfsdk:"base_domain" json:"baseDomain,omitempty"`
		ClaimLifetime *struct {
			Default *string `tfsdk:"default" json:"default,omitempty"`
			Maximum *string `tfsdk:"maximum" json:"maximum,omitempty"`
		} `tfsdk:"claim_lifetime" json:"claimLifetime,omitempty"`
		HibernateAfter    *string `tfsdk:"hibernate_after" json:"hibernateAfter,omitempty"`
		HibernationConfig *struct {
			ResumeTimeout *string `tfsdk:"resume_timeout" json:"resumeTimeout,omitempty"`
		} `tfsdk:"hibernation_config" json:"hibernationConfig,omitempty"`
		ImageSetRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"image_set_ref" json:"imageSetRef,omitempty"`
		InstallAttemptsLimit           *int64 `tfsdk:"install_attempts_limit" json:"installAttemptsLimit,omitempty"`
		InstallConfigSecretTemplateRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"install_config_secret_template_ref" json:"installConfigSecretTemplateRef,omitempty"`
		Inventory *[]struct {
			Kind *string `tfsdk:"kind" json:"kind,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"inventory" json:"inventory,omitempty"`
		Labels        *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		MaxConcurrent *int64             `tfsdk:"max_concurrent" json:"maxConcurrent,omitempty"`
		MaxSize       *int64             `tfsdk:"max_size" json:"maxSize,omitempty"`
		Platform      *struct {
			AgentBareMetal *struct {
				AgentSelector *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
				} `tfsdk:"agent_selector" json:"agentSelector,omitempty"`
			} `tfsdk:"agent_bare_metal" json:"agentBareMetal,omitempty"`
			Alibabacloud *struct {
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
				PrivateLink *struct {
					AdditionalAllowedPrincipals *[]string `tfsdk:"additional_allowed_principals" json:"additionalAllowedPrincipals,omitempty"`
					Enabled                     *bool     `tfsdk:"enabled" json:"enabled,omitempty"`
				} `tfsdk:"private_link" json:"privateLink,omitempty"`
				Region   *string            `tfsdk:"region" json:"region,omitempty"`
				UserTags *map[string]string `tfsdk:"user_tags" json:"userTags,omitempty"`
			} `tfsdk:"aws" json:"aws,omitempty"`
			Azure *struct {
				BaseDomainResourceGroupName *string `tfsdk:"base_domain_resource_group_name" json:"baseDomainResourceGroupName,omitempty"`
				CloudName                   *string `tfsdk:"cloud_name" json:"cloudName,omitempty"`
				CredentialsSecretRef        *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" json:"credentialsSecretRef,omitempty"`
				Region *string `tfsdk:"region" json:"region,omitempty"`
			} `tfsdk:"azure" json:"azure,omitempty"`
			Baremetal *struct {
				LibvirtSSHPrivateKeySecretRef *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"libvirt_ssh_private_key_secret_ref" json:"libvirtSSHPrivateKeySecretRef,omitempty"`
			} `tfsdk:"baremetal" json:"baremetal,omitempty"`
			Gcp *struct {
				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" json:"credentialsSecretRef,omitempty"`
				Region *string `tfsdk:"region" json:"region,omitempty"`
			} `tfsdk:"gcp" json:"gcp,omitempty"`
			Ibmcloud *struct {
				AccountID            *string `tfsdk:"account_id" json:"accountID,omitempty"`
				CisInstanceCRN       *string `tfsdk:"cis_instance_crn" json:"cisInstanceCRN,omitempty"`
				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" json:"credentialsSecretRef,omitempty"`
				Region *string `tfsdk:"region" json:"region,omitempty"`
			} `tfsdk:"ibmcloud" json:"ibmcloud,omitempty"`
			None      *map[string]string `tfsdk:"none" json:"none,omitempty"`
			Openstack *struct {
				CertificatesSecretRef *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"certificates_secret_ref" json:"certificatesSecretRef,omitempty"`
				Cloud                *string `tfsdk:"cloud" json:"cloud,omitempty"`
				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" json:"credentialsSecretRef,omitempty"`
				TrunkSupport *bool `tfsdk:"trunk_support" json:"trunkSupport,omitempty"`
			} `tfsdk:"openstack" json:"openstack,omitempty"`
			Ovirt *struct {
				CertificatesSecretRef *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"certificates_secret_ref" json:"certificatesSecretRef,omitempty"`
				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" json:"credentialsSecretRef,omitempty"`
				Ovirt_cluster_id   *string `tfsdk:"ovirt_cluster_id" json:"ovirt_cluster_id,omitempty"`
				Ovirt_network_name *string `tfsdk:"ovirt_network_name" json:"ovirt_network_name,omitempty"`
				Storage_domain_id  *string `tfsdk:"storage_domain_id" json:"storage_domain_id,omitempty"`
			} `tfsdk:"ovirt" json:"ovirt,omitempty"`
			Vsphere *struct {
				CertificatesSecretRef *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"certificates_secret_ref" json:"certificatesSecretRef,omitempty"`
				Cluster              *string `tfsdk:"cluster" json:"cluster,omitempty"`
				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" json:"credentialsSecretRef,omitempty"`
				Datacenter       *string `tfsdk:"datacenter" json:"datacenter,omitempty"`
				DefaultDatastore *string `tfsdk:"default_datastore" json:"defaultDatastore,omitempty"`
				Folder           *string `tfsdk:"folder" json:"folder,omitempty"`
				Network          *string `tfsdk:"network" json:"network,omitempty"`
				VCenter          *string `tfsdk:"v_center" json:"vCenter,omitempty"`
			} `tfsdk:"vsphere" json:"vsphere,omitempty"`
		} `tfsdk:"platform" json:"platform,omitempty"`
		PullSecretRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"pull_secret_ref" json:"pullSecretRef,omitempty"`
		RunningCount     *int64 `tfsdk:"running_count" json:"runningCount,omitempty"`
		Size             *int64 `tfsdk:"size" json:"size,omitempty"`
		SkipMachinePools *bool  `tfsdk:"skip_machine_pools" json:"skipMachinePools,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *HiveOpenshiftIoClusterPoolV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_hive_openshift_io_cluster_pool_v1"
}

func (r *HiveOpenshiftIoClusterPoolV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClusterPool represents a pool of clusters that should be kept ready to be given out to users. Clusters are removed from the pool once claimed and then automatically replaced with a new one.",
		MarkdownDescription: "ClusterPool represents a pool of clusters that should be kept ready to be given out to users. Clusters are removed from the pool once claimed and then automatically replaced with a new one.",
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
				Description:         "ClusterPoolSpec defines the desired state of the ClusterPool.",
				MarkdownDescription: "ClusterPoolSpec defines the desired state of the ClusterPool.",
				Attributes: map[string]schema.Attribute{
					"annotations": schema.MapAttribute{
						Description:         "Annotations to be applied to new ClusterDeployments created for the pool. ClusterDeployments that have already been claimed will not be affected when this value is modified.",
						MarkdownDescription: "Annotations to be applied to new ClusterDeployments created for the pool. ClusterDeployments that have already been claimed will not be affected when this value is modified.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"base_domain": schema.StringAttribute{
						Description:         "BaseDomain is the base domain to use for all clusters created in this pool.",
						MarkdownDescription: "BaseDomain is the base domain to use for all clusters created in this pool.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"claim_lifetime": schema.SingleNestedAttribute{
						Description:         "ClaimLifetime defines the lifetimes for claims for the cluster pool.",
						MarkdownDescription: "ClaimLifetime defines the lifetimes for claims for the cluster pool.",
						Attributes: map[string]schema.Attribute{
							"default": schema.StringAttribute{
								Description:         "Default is the default lifetime of the claim when no lifetime is set on the claim itself. This is a Duration value; see https://pkg.go.dev/time#ParseDuration for accepted formats. Note: due to discrepancies in validation vs parsing, we use a Pattern instead of 'Format=duration'. See https://bugzilla.redhat.com/show_bug.cgi?id=2050332 https://github.com/kubernetes/apimachinery/issues/131 https://github.com/kubernetes/apiextensions-apiserver/issues/56",
								MarkdownDescription: "Default is the default lifetime of the claim when no lifetime is set on the claim itself. This is a Duration value; see https://pkg.go.dev/time#ParseDuration for accepted formats. Note: due to discrepancies in validation vs parsing, we use a Pattern instead of 'Format=duration'. See https://bugzilla.redhat.com/show_bug.cgi?id=2050332 https://github.com/kubernetes/apimachinery/issues/131 https://github.com/kubernetes/apiextensions-apiserver/issues/56",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"maximum": schema.StringAttribute{
								Description:         "Maximum is the maximum lifetime of the claim after it is assigned a cluster. If the claim still exists when the lifetime has elapsed, the claim will be deleted by Hive. The lifetime of a claim is the mimimum of the lifetimes set by the cluster pool and the claim itself. This is a Duration value; see https://pkg.go.dev/time#ParseDuration for accepted formats. Note: due to discrepancies in validation vs parsing, we use a Pattern instead of 'Format=duration'. See https://bugzilla.redhat.com/show_bug.cgi?id=2050332 https://github.com/kubernetes/apimachinery/issues/131 https://github.com/kubernetes/apiextensions-apiserver/issues/56",
								MarkdownDescription: "Maximum is the maximum lifetime of the claim after it is assigned a cluster. If the claim still exists when the lifetime has elapsed, the claim will be deleted by Hive. The lifetime of a claim is the mimimum of the lifetimes set by the cluster pool and the claim itself. This is a Duration value; see https://pkg.go.dev/time#ParseDuration for accepted formats. Note: due to discrepancies in validation vs parsing, we use a Pattern instead of 'Format=duration'. See https://bugzilla.redhat.com/show_bug.cgi?id=2050332 https://github.com/kubernetes/apimachinery/issues/131 https://github.com/kubernetes/apiextensions-apiserver/issues/56",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"hibernate_after": schema.StringAttribute{
						Description:         "HibernateAfter will be applied to new ClusterDeployments created for the pool. HibernateAfter will transition clusters in the clusterpool to hibernating power state after it has been running for the given duration. The time that a cluster has been running is the time since the cluster was installed or the time since the cluster last came out of hibernation. This is a Duration value; see https://pkg.go.dev/time#ParseDuration for accepted formats. Note: due to discrepancies in validation vs parsing, we use a Pattern instead of 'Format=duration'. See https://bugzilla.redhat.com/show_bug.cgi?id=2050332 https://github.com/kubernetes/apimachinery/issues/131 https://github.com/kubernetes/apiextensions-apiserver/issues/56",
						MarkdownDescription: "HibernateAfter will be applied to new ClusterDeployments created for the pool. HibernateAfter will transition clusters in the clusterpool to hibernating power state after it has been running for the given duration. The time that a cluster has been running is the time since the cluster was installed or the time since the cluster last came out of hibernation. This is a Duration value; see https://pkg.go.dev/time#ParseDuration for accepted formats. Note: due to discrepancies in validation vs parsing, we use a Pattern instead of 'Format=duration'. See https://bugzilla.redhat.com/show_bug.cgi?id=2050332 https://github.com/kubernetes/apimachinery/issues/131 https://github.com/kubernetes/apiextensions-apiserver/issues/56",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"hibernation_config": schema.SingleNestedAttribute{
						Description:         "HibernationConfig configures the hibernation/resume behavior of ClusterDeployments owned by the ClusterPool.",
						MarkdownDescription: "HibernationConfig configures the hibernation/resume behavior of ClusterDeployments owned by the ClusterPool.",
						Attributes: map[string]schema.Attribute{
							"resume_timeout": schema.StringAttribute{
								Description:         "ResumeTimeout is the maximum amount of time we will wait for an unclaimed ClusterDeployment to resume from hibernation (e.g. at the behest of runningCount, or in preparation for being claimed). If this time is exceeded, the ClusterDeployment will be considered Broken and we will replace it. The default (unspecified or zero) means no timeout -- we will allow the ClusterDeployment to continue trying to resume 'forever'. This is a Duration value; see https://pkg.go.dev/time#ParseDuration for accepted formats. Note: due to discrepancies in validation vs parsing, we use a Pattern instead of 'Format=duration'. See https://bugzilla.redhat.com/show_bug.cgi?id=2050332 https://github.com/kubernetes/apimachinery/issues/131 https://github.com/kubernetes/apiextensions-apiserver/issues/56",
								MarkdownDescription: "ResumeTimeout is the maximum amount of time we will wait for an unclaimed ClusterDeployment to resume from hibernation (e.g. at the behest of runningCount, or in preparation for being claimed). If this time is exceeded, the ClusterDeployment will be considered Broken and we will replace it. The default (unspecified or zero) means no timeout -- we will allow the ClusterDeployment to continue trying to resume 'forever'. This is a Duration value; see https://pkg.go.dev/time#ParseDuration for accepted formats. Note: due to discrepancies in validation vs parsing, we use a Pattern instead of 'Format=duration'. See https://bugzilla.redhat.com/show_bug.cgi?id=2050332 https://github.com/kubernetes/apimachinery/issues/131 https://github.com/kubernetes/apiextensions-apiserver/issues/56",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"image_set_ref": schema.SingleNestedAttribute{
						Description:         "ImageSetRef is a reference to a ClusterImageSet. The release image specified in the ClusterImageSet will be used by clusters created for this cluster pool.",
						MarkdownDescription: "ImageSetRef is a reference to a ClusterImageSet. The release image specified in the ClusterImageSet will be used by clusters created for this cluster pool.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name is the name of the ClusterImageSet that this refers to",
								MarkdownDescription: "Name is the name of the ClusterImageSet that this refers to",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"install_attempts_limit": schema.Int64Attribute{
						Description:         "InstallAttemptsLimit is the maximum number of times Hive will attempt to install the cluster.",
						MarkdownDescription: "InstallAttemptsLimit is the maximum number of times Hive will attempt to install the cluster.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"install_config_secret_template_ref": schema.SingleNestedAttribute{
						Description:         "InstallConfigSecretTemplateRef is a secret with the key install-config.yaml consisting of the content of the install-config.yaml to be used as a template for all clusters in this pool. Cluster specific settings (name, basedomain) will be injected dynamically when the ClusterDeployment install-config Secret is generated.",
						MarkdownDescription: "InstallConfigSecretTemplateRef is a secret with the key install-config.yaml consisting of the content of the install-config.yaml to be used as a template for all clusters in this pool. Cluster specific settings (name, basedomain) will be injected dynamically when the ClusterDeployment install-config Secret is generated.",
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

					"inventory": schema.ListNestedAttribute{
						Description:         "Inventory maintains a list of entries consumed by the ClusterPool to customize the default ClusterDeployment.",
						MarkdownDescription: "Inventory maintains a list of entries consumed by the ClusterPool to customize the default ClusterDeployment.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"kind": schema.StringAttribute{
									Description:         "Kind denotes the kind of the referenced resource. The default is ClusterDeploymentCustomization, which is also currently the only supported value.",
									MarkdownDescription: "Kind denotes the kind of the referenced resource. The default is ClusterDeploymentCustomization, which is also currently the only supported value.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"name": schema.StringAttribute{
									Description:         "Name is the name of the referenced resource.",
									MarkdownDescription: "Name is the name of the referenced resource.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"labels": schema.MapAttribute{
						Description:         "Labels to be applied to new ClusterDeployments created for the pool. ClusterDeployments that have already been claimed will not be affected when this value is modified.",
						MarkdownDescription: "Labels to be applied to new ClusterDeployments created for the pool. ClusterDeployments that have already been claimed will not be affected when this value is modified.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"max_concurrent": schema.Int64Attribute{
						Description:         "MaxConcurrent is the maximum number of clusters that will be provisioned or deprovisioned at an time. This includes the claimed clusters being deprovisioned. By default there is no limit.",
						MarkdownDescription: "MaxConcurrent is the maximum number of clusters that will be provisioned or deprovisioned at an time. This includes the claimed clusters being deprovisioned. By default there is no limit.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"max_size": schema.Int64Attribute{
						Description:         "MaxSize is the maximum number of clusters that will be provisioned including clusters that have been claimed and ones waiting to be used. By default there is no limit.",
						MarkdownDescription: "MaxSize is the maximum number of clusters that will be provisioned including clusters that have been claimed and ones waiting to be used. By default there is no limit.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"platform": schema.SingleNestedAttribute{
						Description:         "Platform encompasses the desired platform for the cluster.",
						MarkdownDescription: "Platform encompasses the desired platform for the cluster.",
						Attributes: map[string]schema.Attribute{
							"agent_bare_metal": schema.SingleNestedAttribute{
								Description:         "AgentBareMetal is the configuration used when performing an Assisted Agent based installation to bare metal.",
								MarkdownDescription: "AgentBareMetal is the configuration used when performing an Assisted Agent based installation to bare metal.",
								Attributes: map[string]schema.Attribute{
									"agent_selector": schema.SingleNestedAttribute{
										Description:         "AgentSelector is a label selector used for associating relevant custom resources with this cluster. (Agent, BareMetalHost, etc)",
										MarkdownDescription: "AgentSelector is a label selector used for associating relevant custom resources with this cluster. (Agent, BareMetalHost, etc)",
										Attributes: map[string]schema.Attribute{
											"match_expressions": schema.ListNestedAttribute{
												Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "key is the label key that the selector applies to.",
															MarkdownDescription: "key is the label key that the selector applies to.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"operator": schema.StringAttribute{
															Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
															MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"values": schema.ListAttribute{
															Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
															MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"match_labels": schema.MapAttribute{
												Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												ElementType:         types.StringType,
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

							"alibabacloud": schema.SingleNestedAttribute{
								Description:         "AlibabaCloud is the configuration used when installing on Alibaba Cloud",
								MarkdownDescription: "AlibabaCloud is the configuration used when installing on Alibaba Cloud",
								Attributes: map[string]schema.Attribute{
									"credentials_secret_ref": schema.SingleNestedAttribute{
										Description:         "CredentialsSecretRef refers to a secret that contains Alibaba Cloud account access credentials.",
										MarkdownDescription: "CredentialsSecretRef refers to a secret that contains Alibaba Cloud account access credentials.",
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
										Description:         "Region specifies the Alibaba Cloud region where the cluster will be created.",
										MarkdownDescription: "Region specifies the Alibaba Cloud region where the cluster will be created.",
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
								Description:         "AWS is the configuration used when installing on AWS.",
								MarkdownDescription: "AWS is the configuration used when installing on AWS.",
								Attributes: map[string]schema.Attribute{
									"credentials_assume_role": schema.SingleNestedAttribute{
										Description:         "CredentialsAssumeRole refers to the IAM role that must be assumed to obtain AWS account access for the cluster operations.",
										MarkdownDescription: "CredentialsAssumeRole refers to the IAM role that must be assumed to obtain AWS account access for the cluster operations.",
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
										Description:         "CredentialsSecretRef refers to a secret that contains the AWS account access credentials.",
										MarkdownDescription: "CredentialsSecretRef refers to a secret that contains the AWS account access credentials.",
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

									"private_link": schema.SingleNestedAttribute{
										Description:         "PrivateLink allows uses to enable access to the cluster's API server using AWS PrivateLink. AWS PrivateLink includes a pair of VPC Endpoint Service and VPC Endpoint accross AWS accounts and allows clients to connect to services using AWS's internal networking instead of the Internet.",
										MarkdownDescription: "PrivateLink allows uses to enable access to the cluster's API server using AWS PrivateLink. AWS PrivateLink includes a pair of VPC Endpoint Service and VPC Endpoint accross AWS accounts and allows clients to connect to services using AWS's internal networking instead of the Internet.",
										Attributes: map[string]schema.Attribute{
											"additional_allowed_principals": schema.ListAttribute{
												Description:         "AdditionalAllowedPrincipals is a list of additional allowed principal ARNs to be configured for the Private Link cluster's VPC Endpoint Service. ARNs provided as AdditionalAllowedPrincipals will be configured for the cluster's VPC Endpoint Service in addition to the IAM entity used by Hive.",
												MarkdownDescription: "AdditionalAllowedPrincipals is a list of additional allowed principal ARNs to be configured for the Private Link cluster's VPC Endpoint Service. ARNs provided as AdditionalAllowedPrincipals will be configured for the cluster's VPC Endpoint Service in addition to the IAM entity used by Hive.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"enabled": schema.BoolAttribute{
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

									"region": schema.StringAttribute{
										Description:         "Region specifies the AWS region where the cluster will be created.",
										MarkdownDescription: "Region specifies the AWS region where the cluster will be created.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"user_tags": schema.MapAttribute{
										Description:         "UserTags specifies additional tags for AWS resources created for the cluster.",
										MarkdownDescription: "UserTags specifies additional tags for AWS resources created for the cluster.",
										ElementType:         types.StringType,
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
								Description:         "Azure is the configuration used when installing on Azure.",
								MarkdownDescription: "Azure is the configuration used when installing on Azure.",
								Attributes: map[string]schema.Attribute{
									"base_domain_resource_group_name": schema.StringAttribute{
										Description:         "BaseDomainResourceGroupName specifies the resource group where the azure DNS zone for the base domain is found",
										MarkdownDescription: "BaseDomainResourceGroupName specifies the resource group where the azure DNS zone for the base domain is found",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"cloud_name": schema.StringAttribute{
										Description:         "cloudName is the name of the Azure cloud environment which can be used to configure the Azure SDK with the appropriate Azure API endpoints. If empty, the value is equal to 'AzurePublicCloud'.",
										MarkdownDescription: "cloudName is the name of the Azure cloud environment which can be used to configure the Azure SDK with the appropriate Azure API endpoints. If empty, the value is equal to 'AzurePublicCloud'.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"credentials_secret_ref": schema.SingleNestedAttribute{
										Description:         "CredentialsSecretRef refers to a secret that contains the Azure account access credentials.",
										MarkdownDescription: "CredentialsSecretRef refers to a secret that contains the Azure account access credentials.",
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
										Description:         "Region specifies the Azure region where the cluster will be created.",
										MarkdownDescription: "Region specifies the Azure region where the cluster will be created.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"baremetal": schema.SingleNestedAttribute{
								Description:         "BareMetal is the configuration used when installing on bare metal.",
								MarkdownDescription: "BareMetal is the configuration used when installing on bare metal.",
								Attributes: map[string]schema.Attribute{
									"libvirt_ssh_private_key_secret_ref": schema.SingleNestedAttribute{
										Description:         "LibvirtSSHPrivateKeySecretRef is the reference to the secret that contains the private SSH key to use for access to the libvirt provisioning host. The SSH private key is expected to be in the secret data under the 'ssh-privatekey' key.",
										MarkdownDescription: "LibvirtSSHPrivateKeySecretRef is the reference to the secret that contains the private SSH key to use for access to the libvirt provisioning host. The SSH private key is expected to be in the secret data under the 'ssh-privatekey' key.",
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

							"gcp": schema.SingleNestedAttribute{
								Description:         "GCP is the configuration used when installing on Google Cloud Platform.",
								MarkdownDescription: "GCP is the configuration used when installing on Google Cloud Platform.",
								Attributes: map[string]schema.Attribute{
									"credentials_secret_ref": schema.SingleNestedAttribute{
										Description:         "CredentialsSecretRef refers to a secret that contains the GCP account access credentials.",
										MarkdownDescription: "CredentialsSecretRef refers to a secret that contains the GCP account access credentials.",
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
										Description:         "Region specifies the GCP region where the cluster will be created.",
										MarkdownDescription: "Region specifies the GCP region where the cluster will be created.",
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
								Description:         "IBMCloud is the configuration used when installing on IBM Cloud",
								MarkdownDescription: "IBMCloud is the configuration used when installing on IBM Cloud",
								Attributes: map[string]schema.Attribute{
									"account_id": schema.StringAttribute{
										Description:         "AccountID is the IBM Cloud Account ID. AccountID is DEPRECATED and is gathered via the IBM Cloud API for the provided credentials. This field will be ignored.",
										MarkdownDescription: "AccountID is the IBM Cloud Account ID. AccountID is DEPRECATED and is gathered via the IBM Cloud API for the provided credentials. This field will be ignored.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"cis_instance_crn": schema.StringAttribute{
										Description:         "CISInstanceCRN is the IBM Cloud Internet Services Instance CRN CISInstanceCRN is DEPRECATED and gathered via the IBM Cloud API for the provided credentials and cluster deployment base domain. This field will be ignored.",
										MarkdownDescription: "CISInstanceCRN is the IBM Cloud Internet Services Instance CRN CISInstanceCRN is DEPRECATED and gathered via the IBM Cloud API for the provided credentials and cluster deployment base domain. This field will be ignored.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"credentials_secret_ref": schema.SingleNestedAttribute{
										Description:         "CredentialsSecretRef refers to a secret that contains IBM Cloud account access credentials.",
										MarkdownDescription: "CredentialsSecretRef refers to a secret that contains IBM Cloud account access credentials.",
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
										Description:         "Region specifies the IBM Cloud region where the cluster will be created.",
										MarkdownDescription: "Region specifies the IBM Cloud region where the cluster will be created.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"none": schema.MapAttribute{
								Description:         "None indicates platform-agnostic install. https://docs.openshift.com/container-platform/4.7/installing/installing_platform_agnostic/installing-platform-agnostic.html",
								MarkdownDescription: "None indicates platform-agnostic install. https://docs.openshift.com/container-platform/4.7/installing/installing_platform_agnostic/installing-platform-agnostic.html",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"openstack": schema.SingleNestedAttribute{
								Description:         "OpenStack is the configuration used when installing on OpenStack",
								MarkdownDescription: "OpenStack is the configuration used when installing on OpenStack",
								Attributes: map[string]schema.Attribute{
									"certificates_secret_ref": schema.SingleNestedAttribute{
										Description:         "CertificatesSecretRef refers to a secret that contains CA certificates necessary for communicating with the OpenStack. There is additional configuration required for the OpenShift cluster to trust the certificates provided in this secret. The 'clouds.yaml' file included in the credentialsSecretRef Secret must also include a reference to the certificate bundle file for the OpenShift cluster being created to trust the OpenStack endpoints. The 'clouds.yaml' file must set the 'cacert' field to either '/etc/openstack-ca/<key name containing the trust bundle in credentialsSecretRef Secret>' or '/etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem'.  For example, '''clouds.yaml clouds: shiftstack: auth: ... cacert: '/etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem' '''",
										MarkdownDescription: "CertificatesSecretRef refers to a secret that contains CA certificates necessary for communicating with the OpenStack. There is additional configuration required for the OpenShift cluster to trust the certificates provided in this secret. The 'clouds.yaml' file included in the credentialsSecretRef Secret must also include a reference to the certificate bundle file for the OpenShift cluster being created to trust the OpenStack endpoints. The 'clouds.yaml' file must set the 'cacert' field to either '/etc/openstack-ca/<key name containing the trust bundle in credentialsSecretRef Secret>' or '/etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem'.  For example, '''clouds.yaml clouds: shiftstack: auth: ... cacert: '/etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem' '''",
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
										Description:         "Cloud will be used to indicate the OS_CLOUD value to use the right section from the clouds.yaml in the CredentialsSecretRef.",
										MarkdownDescription: "Cloud will be used to indicate the OS_CLOUD value to use the right section from the clouds.yaml in the CredentialsSecretRef.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"credentials_secret_ref": schema.SingleNestedAttribute{
										Description:         "CredentialsSecretRef refers to a secret that contains the OpenStack account access credentials.",
										MarkdownDescription: "CredentialsSecretRef refers to a secret that contains the OpenStack account access credentials.",
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

									"trunk_support": schema.BoolAttribute{
										Description:         "TrunkSupport indicates whether or not to use trunk ports in your OpenShift cluster.",
										MarkdownDescription: "TrunkSupport indicates whether or not to use trunk ports in your OpenShift cluster.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"ovirt": schema.SingleNestedAttribute{
								Description:         "Ovirt is the configuration used when installing on oVirt",
								MarkdownDescription: "Ovirt is the configuration used when installing on oVirt",
								Attributes: map[string]schema.Attribute{
									"certificates_secret_ref": schema.SingleNestedAttribute{
										Description:         "CertificatesSecretRef refers to a secret that contains the oVirt CA certificates necessary for communicating with oVirt.",
										MarkdownDescription: "CertificatesSecretRef refers to a secret that contains the oVirt CA certificates necessary for communicating with oVirt.",
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
										Description:         "CredentialsSecretRef refers to a secret that contains the oVirt account access credentials with fields: ovirt_url, ovirt_username, ovirt_password, ovirt_ca_bundle",
										MarkdownDescription: "CredentialsSecretRef refers to a secret that contains the oVirt account access credentials with fields: ovirt_url, ovirt_username, ovirt_password, ovirt_ca_bundle",
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

									"ovirt_cluster_id": schema.StringAttribute{
										Description:         "The target cluster under which all VMs will run",
										MarkdownDescription: "The target cluster under which all VMs will run",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"ovirt_network_name": schema.StringAttribute{
										Description:         "The target network of all the network interfaces of the nodes. Omitting defaults to ovirtmgmt network which is a default network for evert ovirt cluster.",
										MarkdownDescription: "The target network of all the network interfaces of the nodes. Omitting defaults to ovirtmgmt network which is a default network for evert ovirt cluster.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"storage_domain_id": schema.StringAttribute{
										Description:         "The target storage domain under which all VM disk would be created.",
										MarkdownDescription: "The target storage domain under which all VM disk would be created.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"vsphere": schema.SingleNestedAttribute{
								Description:         "VSphere is the configuration used when installing on vSphere",
								MarkdownDescription: "VSphere is the configuration used when installing on vSphere",
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

									"cluster": schema.StringAttribute{
										Description:         "Cluster is the name of the cluster virtual machines will be cloned into.",
										MarkdownDescription: "Cluster is the name of the cluster virtual machines will be cloned into.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"credentials_secret_ref": schema.SingleNestedAttribute{
										Description:         "CredentialsSecretRef refers to a secret that contains the vSphere account access credentials: GOVC_USERNAME, GOVC_PASSWORD fields.",
										MarkdownDescription: "CredentialsSecretRef refers to a secret that contains the vSphere account access credentials: GOVC_USERNAME, GOVC_PASSWORD fields.",
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

									"datacenter": schema.StringAttribute{
										Description:         "Datacenter is the name of the datacenter to use in the vCenter.",
										MarkdownDescription: "Datacenter is the name of the datacenter to use in the vCenter.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"default_datastore": schema.StringAttribute{
										Description:         "DefaultDatastore is the default datastore to use for provisioning volumes.",
										MarkdownDescription: "DefaultDatastore is the default datastore to use for provisioning volumes.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"folder": schema.StringAttribute{
										Description:         "Folder is the name of the folder that will be used and/or created for virtual machines.",
										MarkdownDescription: "Folder is the name of the folder that will be used and/or created for virtual machines.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"network": schema.StringAttribute{
										Description:         "Network specifies the name of the network to be used by the cluster.",
										MarkdownDescription: "Network specifies the name of the network to be used by the cluster.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"v_center": schema.StringAttribute{
										Description:         "VCenter is the domain name or IP address of the vCenter.",
										MarkdownDescription: "VCenter is the domain name or IP address of the vCenter.",
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

					"pull_secret_ref": schema.SingleNestedAttribute{
						Description:         "PullSecretRef is the reference to the secret to use when pulling images.",
						MarkdownDescription: "PullSecretRef is the reference to the secret to use when pulling images.",
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

					"running_count": schema.Int64Attribute{
						Description:         "RunningCount is the number of clusters we should keep running. The remainder will be kept hibernated until claimed. By default no clusters will be kept running (all will be hibernated).",
						MarkdownDescription: "RunningCount is the number of clusters we should keep running. The remainder will be kept hibernated until claimed. By default no clusters will be kept running (all will be hibernated).",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"size": schema.Int64Attribute{
						Description:         "Size is the default number of clusters that we should keep provisioned and waiting for use.",
						MarkdownDescription: "Size is the default number of clusters that we should keep provisioned and waiting for use.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"skip_machine_pools": schema.BoolAttribute{
						Description:         "SkipMachinePools allows creating clusterpools where the machinepools are not managed by hive after cluster creation",
						MarkdownDescription: "SkipMachinePools allows creating clusterpools where the machinepools are not managed by hive after cluster creation",
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

func (r *HiveOpenshiftIoClusterPoolV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *HiveOpenshiftIoClusterPoolV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_hive_openshift_io_cluster_pool_v1")

	var data HiveOpenshiftIoClusterPoolV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "hive.openshift.io", Version: "v1", Resource: "ClusterPool"}).
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

	var readResponse HiveOpenshiftIoClusterPoolV1DataSourceData
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
	data.Kind = pointer.String("ClusterPool")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
