/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package hive_openshift_io_v1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &HiveOpenshiftIoClusterDeploymentV1Manifest{}
)

func NewHiveOpenshiftIoClusterDeploymentV1Manifest() datasource.DataSource {
	return &HiveOpenshiftIoClusterDeploymentV1Manifest{}
}

type HiveOpenshiftIoClusterDeploymentV1Manifest struct{}

type HiveOpenshiftIoClusterDeploymentV1ManifestData struct {
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
		BaseDomain                             *string `tfsdk:"base_domain" json:"baseDomain,omitempty"`
		BoundServiceAccountSigningKeySecretRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"bound_service_account_signing_key_secret_ref" json:"boundServiceAccountSigningKeySecretRef,omitempty"`
		CertificateBundles *[]struct {
			CertificateSecretRef *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"certificate_secret_ref" json:"certificateSecretRef,omitempty"`
			Generate *bool   `tfsdk:"generate" json:"generate,omitempty"`
			Name     *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"certificate_bundles" json:"certificateBundles,omitempty"`
		ClusterInstallRef *struct {
			Group   *string `tfsdk:"group" json:"group,omitempty"`
			Kind    *string `tfsdk:"kind" json:"kind,omitempty"`
			Name    *string `tfsdk:"name" json:"name,omitempty"`
			Version *string `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"cluster_install_ref" json:"clusterInstallRef,omitempty"`
		ClusterMetadata *struct {
			AdminKubeconfigSecretRef *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"admin_kubeconfig_secret_ref" json:"adminKubeconfigSecretRef,omitempty"`
			AdminPasswordSecretRef *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"admin_password_secret_ref" json:"adminPasswordSecretRef,omitempty"`
			ClusterID *string `tfsdk:"cluster_id" json:"clusterID,omitempty"`
			InfraID   *string `tfsdk:"infra_id" json:"infraID,omitempty"`
			Platform  *struct {
				Aws *struct {
					HostedZoneRole *string `tfsdk:"hosted_zone_role" json:"hostedZoneRole,omitempty"`
				} `tfsdk:"aws" json:"aws,omitempty"`
				Azure *struct {
					ResourceGroupName *string `tfsdk:"resource_group_name" json:"resourceGroupName,omitempty"`
				} `tfsdk:"azure" json:"azure,omitempty"`
				Gcp *struct {
					NetworkProjectID *string `tfsdk:"network_project_id" json:"networkProjectID,omitempty"`
				} `tfsdk:"gcp" json:"gcp,omitempty"`
			} `tfsdk:"platform" json:"platform,omitempty"`
		} `tfsdk:"cluster_metadata" json:"clusterMetadata,omitempty"`
		ClusterName    *string `tfsdk:"cluster_name" json:"clusterName,omitempty"`
		ClusterPoolRef *struct {
			ClaimName                      *string `tfsdk:"claim_name" json:"claimName,omitempty"`
			ClaimedTimestamp               *string `tfsdk:"claimed_timestamp" json:"claimedTimestamp,omitempty"`
			ClusterDeploymentCustomization *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"cluster_deployment_customization" json:"clusterDeploymentCustomization,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			PoolName  *string `tfsdk:"pool_name" json:"poolName,omitempty"`
		} `tfsdk:"cluster_pool_ref" json:"clusterPoolRef,omitempty"`
		ControlPlaneConfig *struct {
			ApiServerIPOverride *string `tfsdk:"api_server_ip_override" json:"apiServerIPOverride,omitempty"`
			ApiURLOverride      *string `tfsdk:"api_url_override" json:"apiURLOverride,omitempty"`
			ServingCertificates *struct {
				Additional *[]struct {
					Domain *string `tfsdk:"domain" json:"domain,omitempty"`
					Name   *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"additional" json:"additional,omitempty"`
				Default *string `tfsdk:"default" json:"default,omitempty"`
			} `tfsdk:"serving_certificates" json:"servingCertificates,omitempty"`
		} `tfsdk:"control_plane_config" json:"controlPlaneConfig,omitempty"`
		HibernateAfter *string `tfsdk:"hibernate_after" json:"hibernateAfter,omitempty"`
		Ingress        *[]struct {
			Domain             *string `tfsdk:"domain" json:"domain,omitempty"`
			HttpErrorCodePages *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"http_error_code_pages" json:"httpErrorCodePages,omitempty"`
			Name              *string `tfsdk:"name" json:"name,omitempty"`
			NamespaceSelector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
			RouteSelector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"route_selector" json:"routeSelector,omitempty"`
			ServingCertificate *string `tfsdk:"serving_certificate" json:"servingCertificate,omitempty"`
		} `tfsdk:"ingress" json:"ingress,omitempty"`
		InstallAttemptsLimit *int64 `tfsdk:"install_attempts_limit" json:"installAttemptsLimit,omitempty"`
		Installed            *bool  `tfsdk:"installed" json:"installed,omitempty"`
		ManageDNS            *bool  `tfsdk:"manage_dns" json:"manageDNS,omitempty"`
		Platform             *struct {
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
				PrivateServiceConnect *struct {
					Enabled           *bool `tfsdk:"enabled" json:"enabled,omitempty"`
					ServiceAttachment *struct {
						Subnet *struct {
							Cidr     *string `tfsdk:"cidr" json:"cidr,omitempty"`
							Existing *struct {
								Name    *string `tfsdk:"name" json:"name,omitempty"`
								Project *string `tfsdk:"project" json:"project,omitempty"`
							} `tfsdk:"existing" json:"existing,omitempty"`
						} `tfsdk:"subnet" json:"subnet,omitempty"`
					} `tfsdk:"service_attachment" json:"serviceAttachment,omitempty"`
				} `tfsdk:"private_service_connect" json:"privateServiceConnect,omitempty"`
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
		PowerState       *string `tfsdk:"power_state" json:"powerState,omitempty"`
		PreserveOnDelete *bool   `tfsdk:"preserve_on_delete" json:"preserveOnDelete,omitempty"`
		Provisioning     *struct {
			ImageSetRef *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"image_set_ref" json:"imageSetRef,omitempty"`
			InstallConfigSecretRef *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"install_config_secret_ref" json:"installConfigSecretRef,omitempty"`
			InstallerEnv *[]struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					FieldRef *struct {
						ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
						FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
					} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
					ResourceFieldRef *struct {
						ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
						Divisor       *string `tfsdk:"divisor" json:"divisor,omitempty"`
						Resource      *string `tfsdk:"resource" json:"resource,omitempty"`
					} `tfsdk:"resource_field_ref" json:"resourceFieldRef,omitempty"`
					SecretKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"installer_env" json:"installerEnv,omitempty"`
			InstallerImageOverride *string `tfsdk:"installer_image_override" json:"installerImageOverride,omitempty"`
			ManifestsConfigMapRef  *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"manifests_config_map_ref" json:"manifestsConfigMapRef,omitempty"`
			ManifestsSecretRef *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"manifests_secret_ref" json:"manifestsSecretRef,omitempty"`
			ReleaseImage           *string   `tfsdk:"release_image" json:"releaseImage,omitempty"`
			SshKnownHosts          *[]string `tfsdk:"ssh_known_hosts" json:"sshKnownHosts,omitempty"`
			SshPrivateKeySecretRef *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"ssh_private_key_secret_ref" json:"sshPrivateKeySecretRef,omitempty"`
		} `tfsdk:"provisioning" json:"provisioning,omitempty"`
		PullSecretRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"pull_secret_ref" json:"pullSecretRef,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *HiveOpenshiftIoClusterDeploymentV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_hive_openshift_io_cluster_deployment_v1_manifest"
}

func (r *HiveOpenshiftIoClusterDeploymentV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClusterDeployment is the Schema for the clusterdeployments API",
		MarkdownDescription: "ClusterDeployment is the Schema for the clusterdeployments API",
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
				Description:         "ClusterDeploymentSpec defines the desired state of ClusterDeployment",
				MarkdownDescription: "ClusterDeploymentSpec defines the desired state of ClusterDeployment",
				Attributes: map[string]schema.Attribute{
					"base_domain": schema.StringAttribute{
						Description:         "BaseDomain is the base domain to which the cluster should belong.",
						MarkdownDescription: "BaseDomain is the base domain to which the cluster should belong.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"bound_service_account_signing_key_secret_ref": schema.SingleNestedAttribute{
						Description:         "BoundServiceAccountSigningKeySecretRef refers to a Secret that contains a 'bound-service-account-signing-key.key' data key pointing to the private key that will be used to sign ServiceAccount objects. Primarily used to provision AWS clusters to use Amazon's Security Token Service.",
						MarkdownDescription: "BoundServiceAccountSigningKeySecretRef refers to a Secret that contains a 'bound-service-account-signing-key.key' data key pointing to the private key that will be used to sign ServiceAccount objects. Primarily used to provision AWS clusters to use Amazon's Security Token Service.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"certificate_bundles": schema.ListNestedAttribute{
						Description:         "CertificateBundles is a list of certificate bundles associated with this cluster",
						MarkdownDescription: "CertificateBundles is a list of certificate bundles associated with this cluster",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"certificate_secret_ref": schema.SingleNestedAttribute{
									Description:         "CertificateSecretRef is the reference to the secret that contains the certificate bundle. If the certificate bundle is to be generated, it will be generated with the name in this reference. Otherwise, it is expected that the secret should exist in the same namespace as the ClusterDeployment",
									MarkdownDescription: "CertificateSecretRef is the reference to the secret that contains the certificate bundle. If the certificate bundle is to be generated, it will be generated with the name in this reference. Otherwise, it is expected that the secret should exist in the same namespace as the ClusterDeployment",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
											MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"generate": schema.BoolAttribute{
									Description:         "Generate indicates whether this bundle should have real certificates generated for it.",
									MarkdownDescription: "Generate indicates whether this bundle should have real certificates generated for it.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name is an identifier that must be unique within the bundle and must be referenced by an ingress or by the control plane serving certs",
									MarkdownDescription: "Name is an identifier that must be unique within the bundle and must be referenced by an ingress or by the control plane serving certs",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"cluster_install_ref": schema.SingleNestedAttribute{
						Description:         "ClusterInstallLocalReference provides reference to an object that implements the hivecontract ClusterInstall. The namespace of the object is same as the ClusterDeployment. This cannot be set when Provisioning is also set.",
						MarkdownDescription: "ClusterInstallLocalReference provides reference to an object that implements the hivecontract ClusterInstall. The namespace of the object is same as the ClusterDeployment. This cannot be set when Provisioning is also set.",
						Attributes: map[string]schema.Attribute{
							"group": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"version": schema.StringAttribute{
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

					"cluster_metadata": schema.SingleNestedAttribute{
						Description:         "ClusterMetadata contains metadata information about the installed cluster.",
						MarkdownDescription: "ClusterMetadata contains metadata information about the installed cluster.",
						Attributes: map[string]schema.Attribute{
							"admin_kubeconfig_secret_ref": schema.SingleNestedAttribute{
								Description:         "AdminKubeconfigSecretRef references the secret containing the admin kubeconfig for this cluster.",
								MarkdownDescription: "AdminKubeconfigSecretRef references the secret containing the admin kubeconfig for this cluster.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
							},

							"admin_password_secret_ref": schema.SingleNestedAttribute{
								Description:         "AdminPasswordSecretRef references the secret containing the admin username/password which can be used to login to this cluster.",
								MarkdownDescription: "AdminPasswordSecretRef references the secret containing the admin username/password which can be used to login to this cluster.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"cluster_id": schema.StringAttribute{
								Description:         "ClusterID is a globally unique identifier for this cluster generated during installation. Used for reporting metrics among other places.",
								MarkdownDescription: "ClusterID is a globally unique identifier for this cluster generated during installation. Used for reporting metrics among other places.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"infra_id": schema.StringAttribute{
								Description:         "InfraID is an identifier for this cluster generated during installation and used for tagging/naming resources in cloud providers.",
								MarkdownDescription: "InfraID is an identifier for this cluster generated during installation and used for tagging/naming resources in cloud providers.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"platform": schema.SingleNestedAttribute{
								Description:         "Platform holds platform-specific cluster metadata",
								MarkdownDescription: "Platform holds platform-specific cluster metadata",
								Attributes: map[string]schema.Attribute{
									"aws": schema.SingleNestedAttribute{
										Description:         "AWS holds AWS-specific cluster metadata",
										MarkdownDescription: "AWS holds AWS-specific cluster metadata",
										Attributes: map[string]schema.Attribute{
											"hosted_zone_role": schema.StringAttribute{
												Description:         "HostedZoneRole is the role to assume when performing operations on a hosted zone owned by another account.",
												MarkdownDescription: "HostedZoneRole is the role to assume when performing operations on a hosted zone owned by another account.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"azure": schema.SingleNestedAttribute{
										Description:         "Azure holds azure-specific cluster metadata",
										MarkdownDescription: "Azure holds azure-specific cluster metadata",
										Attributes: map[string]schema.Attribute{
											"resource_group_name": schema.StringAttribute{
												Description:         "ResourceGroupName is the name of the resource group in which the cluster resources were created.",
												MarkdownDescription: "ResourceGroupName is the name of the resource group in which the cluster resources were created.",
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
										Description:         "GCP holds GCP-specific cluster metadata",
										MarkdownDescription: "GCP holds GCP-specific cluster metadata",
										Attributes: map[string]schema.Attribute{
											"network_project_id": schema.StringAttribute{
												Description:         "NetworkProjectID is used for shared VPC setups",
												MarkdownDescription: "NetworkProjectID is used for shared VPC setups",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"cluster_name": schema.StringAttribute{
						Description:         "ClusterName is the friendly name of the cluster. It is used for subdomains, some resource tagging, and other instances where a friendly name for the cluster is useful.",
						MarkdownDescription: "ClusterName is the friendly name of the cluster. It is used for subdomains, some resource tagging, and other instances where a friendly name for the cluster is useful.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"cluster_pool_ref": schema.SingleNestedAttribute{
						Description:         "ClusterPoolRef is a reference to the ClusterPool that this ClusterDeployment originated from.",
						MarkdownDescription: "ClusterPoolRef is a reference to the ClusterPool that this ClusterDeployment originated from.",
						Attributes: map[string]schema.Attribute{
							"claim_name": schema.StringAttribute{
								Description:         "ClaimName is the name of the ClusterClaim that claimed the cluster from the pool.",
								MarkdownDescription: "ClaimName is the name of the ClusterClaim that claimed the cluster from the pool.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"claimed_timestamp": schema.StringAttribute{
								Description:         "ClaimedTimestamp is the time this cluster was assigned to a ClusterClaim. This is only used for ClusterDeployments belonging to ClusterPools.",
								MarkdownDescription: "ClaimedTimestamp is the time this cluster was assigned to a ClusterClaim. This is only used for ClusterDeployments belonging to ClusterPools.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									validators.DateTime64Validator(),
								},
							},

							"cluster_deployment_customization": schema.SingleNestedAttribute{
								Description:         "CustomizationRef is the ClusterPool Inventory claimed customization for this ClusterDeployment. The Customization exists in the ClusterPool namespace.",
								MarkdownDescription: "CustomizationRef is the ClusterPool Inventory claimed customization for this ClusterDeployment. The Customization exists in the ClusterPool namespace.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace is the namespace where the ClusterPool resides.",
								MarkdownDescription: "Namespace is the namespace where the ClusterPool resides.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"pool_name": schema.StringAttribute{
								Description:         "PoolName is the name of the ClusterPool for which the cluster was created.",
								MarkdownDescription: "PoolName is the name of the ClusterPool for which the cluster was created.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"control_plane_config": schema.SingleNestedAttribute{
						Description:         "ControlPlaneConfig contains additional configuration for the target cluster's control plane",
						MarkdownDescription: "ControlPlaneConfig contains additional configuration for the target cluster's control plane",
						Attributes: map[string]schema.Attribute{
							"api_server_ip_override": schema.StringAttribute{
								Description:         "APIServerIPOverride is the optional override of the API server IP address. Hive will use this IP address for creating TCP connections. Port from the original API server URL will be used. This field can be used when repointing the APIServer's DNS is not viable option.",
								MarkdownDescription: "APIServerIPOverride is the optional override of the API server IP address. Hive will use this IP address for creating TCP connections. Port from the original API server URL will be used. This field can be used when repointing the APIServer's DNS is not viable option.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"api_url_override": schema.StringAttribute{
								Description:         "APIURLOverride is the optional URL override to which Hive will transition for communication with the API server of the remote cluster. When a remote cluster is created, Hive will initially communicate using the API URL established during installation. If an API URL Override is specified, Hive will periodically attempt to connect to the remote cluster using the override URL. Once Hive has determined that the override URL is active, Hive will use the override URL for further communications with the API server of the remote cluster.",
								MarkdownDescription: "APIURLOverride is the optional URL override to which Hive will transition for communication with the API server of the remote cluster. When a remote cluster is created, Hive will initially communicate using the API URL established during installation. If an API URL Override is specified, Hive will periodically attempt to connect to the remote cluster using the override URL. Once Hive has determined that the override URL is active, Hive will use the override URL for further communications with the API server of the remote cluster.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"serving_certificates": schema.SingleNestedAttribute{
								Description:         "ServingCertificates specifies serving certificates for the control plane",
								MarkdownDescription: "ServingCertificates specifies serving certificates for the control plane",
								Attributes: map[string]schema.Attribute{
									"additional": schema.ListNestedAttribute{
										Description:         "Additional is a list of additional domains and certificates that are also associated with the control plane's api endpoint.",
										MarkdownDescription: "Additional is a list of additional domains and certificates that are also associated with the control plane's api endpoint.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"domain": schema.StringAttribute{
													Description:         "Domain is the domain of the additional control plane certificate",
													MarkdownDescription: "Domain is the domain of the additional control plane certificate",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name references a CertificateBundle in the ClusterDeployment.Spec that should be used for this additional certificate.",
													MarkdownDescription: "Name references a CertificateBundle in the ClusterDeployment.Spec that should be used for this additional certificate.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"default": schema.StringAttribute{
										Description:         "Default references the name of a CertificateBundle in the ClusterDeployment that should be used for the control plane's default endpoint.",
										MarkdownDescription: "Default references the name of a CertificateBundle in the ClusterDeployment that should be used for the control plane's default endpoint.",
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

					"hibernate_after": schema.StringAttribute{
						Description:         "HibernateAfter will transition a cluster to hibernating power state after it has been running for the given duration. The time that a cluster has been running is the time since the cluster was installed or the time since the cluster last came out of hibernation. This is a Duration value; see https://pkg.go.dev/time#ParseDuration for accepted formats. Note: due to discrepancies in validation vs parsing, we use a Pattern instead of 'Format=duration'. See https://bugzilla.redhat.com/show_bug.cgi?id=2050332 https://github.com/kubernetes/apimachinery/issues/131 https://github.com/kubernetes/apiextensions-apiserver/issues/56",
						MarkdownDescription: "HibernateAfter will transition a cluster to hibernating power state after it has been running for the given duration. The time that a cluster has been running is the time since the cluster was installed or the time since the cluster last came out of hibernation. This is a Duration value; see https://pkg.go.dev/time#ParseDuration for accepted formats. Note: due to discrepancies in validation vs parsing, we use a Pattern instead of 'Format=duration'. See https://bugzilla.redhat.com/show_bug.cgi?id=2050332 https://github.com/kubernetes/apimachinery/issues/131 https://github.com/kubernetes/apiextensions-apiserver/issues/56",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ns|us|µs|ms|s|m|h))+$`), ""),
						},
					},

					"ingress": schema.ListNestedAttribute{
						Description:         "Ingress allows defining desired clusteringress/shards to be configured on the cluster.",
						MarkdownDescription: "Ingress allows defining desired clusteringress/shards to be configured on the cluster.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"domain": schema.StringAttribute{
									Description:         "Domain (sometimes referred to as shard) is the full DNS suffix that the resulting IngressController object will service (eg abcd.mycluster.mydomain.com).",
									MarkdownDescription: "Domain (sometimes referred to as shard) is the full DNS suffix that the resulting IngressController object will service (eg abcd.mycluster.mydomain.com).",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"http_error_code_pages": schema.SingleNestedAttribute{
									Description:         "HttpErrorCodePages allows configuring custom HTTP error pages using the IngressController object",
									MarkdownDescription: "HttpErrorCodePages allows configuring custom HTTP error pages using the IngressController object",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "name is the metadata.name of the referenced config map",
											MarkdownDescription: "name is the metadata.name of the referenced config map",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"name": schema.StringAttribute{
									Description:         "Name of the ClusterIngress object to create.",
									MarkdownDescription: "Name of the ClusterIngress object to create.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace_selector": schema.SingleNestedAttribute{
									Description:         "NamespaceSelector allows filtering the list of namespaces serviced by the ingress controller.",
									MarkdownDescription: "NamespaceSelector allows filtering the list of namespaces serviced by the ingress controller.",
									Attributes: map[string]schema.Attribute{
										"match_expressions": schema.ListNestedAttribute{
											Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
											MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "key is the label key that the selector applies to.",
														MarkdownDescription: "key is the label key that the selector applies to.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"operator": schema.StringAttribute{
														Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"values": schema.ListAttribute{
														Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"match_labels": schema.MapAttribute{
											Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
											MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"route_selector": schema.SingleNestedAttribute{
									Description:         "RouteSelector allows filtering the set of Routes serviced by the ingress controller",
									MarkdownDescription: "RouteSelector allows filtering the set of Routes serviced by the ingress controller",
									Attributes: map[string]schema.Attribute{
										"match_expressions": schema.ListNestedAttribute{
											Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
											MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "key is the label key that the selector applies to.",
														MarkdownDescription: "key is the label key that the selector applies to.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"operator": schema.StringAttribute{
														Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"values": schema.ListAttribute{
														Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"match_labels": schema.MapAttribute{
											Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
											MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"serving_certificate": schema.StringAttribute{
									Description:         "ServingCertificate references a CertificateBundle in the ClusterDeployment.Spec that should be used for this Ingress",
									MarkdownDescription: "ServingCertificate references a CertificateBundle in the ClusterDeployment.Spec that should be used for this Ingress",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"install_attempts_limit": schema.Int64Attribute{
						Description:         "InstallAttemptsLimit is the maximum number of times Hive will attempt to install the cluster.",
						MarkdownDescription: "InstallAttemptsLimit is the maximum number of times Hive will attempt to install the cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"installed": schema.BoolAttribute{
						Description:         "Installed is true if the cluster has been installed",
						MarkdownDescription: "Installed is true if the cluster has been installed",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"manage_dns": schema.BoolAttribute{
						Description:         "ManageDNS specifies whether a DNSZone should be created and managed automatically for this ClusterDeployment",
						MarkdownDescription: "ManageDNS specifies whether a DNSZone should be created and managed automatically for this ClusterDeployment",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"platform": schema.SingleNestedAttribute{
						Description:         "Platform is the configuration for the specific platform upon which to perform the installation.",
						MarkdownDescription: "Platform is the configuration for the specific platform upon which to perform the installation.",
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
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"operator": schema.StringAttribute{
															Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
															MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"values": schema.ListAttribute{
															Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
															MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"match_labels": schema.MapAttribute{
												Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												ElementType:         types.StringType,
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
										Description:         "CredentialsSecretRef refers to a secret that contains the AWS account access credentials.",
										MarkdownDescription: "CredentialsSecretRef refers to a secret that contains the AWS account access credentials.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
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
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
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

									"region": schema.StringAttribute{
										Description:         "Region specifies the AWS region where the cluster will be created.",
										MarkdownDescription: "Region specifies the AWS region where the cluster will be created.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"user_tags": schema.MapAttribute{
										Description:         "UserTags specifies additional tags for AWS resources created for the cluster.",
										MarkdownDescription: "UserTags specifies additional tags for AWS resources created for the cluster.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"azure": schema.SingleNestedAttribute{
								Description:         "Azure is the configuration used when installing on Azure.",
								MarkdownDescription: "Azure is the configuration used when installing on Azure.",
								Attributes: map[string]schema.Attribute{
									"base_domain_resource_group_name": schema.StringAttribute{
										Description:         "BaseDomainResourceGroupName specifies the resource group where the azure DNS zone for the base domain is found",
										MarkdownDescription: "BaseDomainResourceGroupName specifies the resource group where the azure DNS zone for the base domain is found",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

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
										Description:         "CredentialsSecretRef refers to a secret that contains the Azure account access credentials.",
										MarkdownDescription: "CredentialsSecretRef refers to a secret that contains the Azure account access credentials.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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
										Description:         "Region specifies the Azure region where the cluster will be created.",
										MarkdownDescription: "Region specifies the Azure region where the cluster will be created.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
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
												Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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

							"gcp": schema.SingleNestedAttribute{
								Description:         "GCP is the configuration used when installing on Google Cloud Platform.",
								MarkdownDescription: "GCP is the configuration used when installing on Google Cloud Platform.",
								Attributes: map[string]schema.Attribute{
									"credentials_secret_ref": schema.SingleNestedAttribute{
										Description:         "CredentialsSecretRef refers to a secret that contains the GCP account access credentials.",
										MarkdownDescription: "CredentialsSecretRef refers to a secret that contains the GCP account access credentials.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"private_service_connect": schema.SingleNestedAttribute{
										Description:         "PrivateSericeConnect allows users to enable access to the cluster's API server using GCP Private Service Connect. It includes a forwarding rule paired with a Service Attachment across GCP accounts and allows clients to connect to services using GCP internal networking of using public load balancers.",
										MarkdownDescription: "PrivateSericeConnect allows users to enable access to the cluster's API server using GCP Private Service Connect. It includes a forwarding rule paired with a Service Attachment across GCP accounts and allows clients to connect to services using GCP internal networking of using public load balancers.",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Enabled specifies if Private Service Connect is to be enabled on the cluster.",
												MarkdownDescription: "Enabled specifies if Private Service Connect is to be enabled on the cluster.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"service_attachment": schema.SingleNestedAttribute{
												Description:         "ServiceAttachment configures the service attachment to be used by the cluster.",
												MarkdownDescription: "ServiceAttachment configures the service attachment to be used by the cluster.",
												Attributes: map[string]schema.Attribute{
													"subnet": schema.SingleNestedAttribute{
														Description:         "Subnet configures the subnetwork that contains the service attachment.",
														MarkdownDescription: "Subnet configures the subnetwork that contains the service attachment.",
														Attributes: map[string]schema.Attribute{
															"cidr": schema.StringAttribute{
																Description:         "Cidr specifies the cidr to use when creating a service attachment subnet.",
																MarkdownDescription: "Cidr specifies the cidr to use when creating a service attachment subnet.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"existing": schema.SingleNestedAttribute{
																Description:         "Existing specifies a pre-existing subnet to use instead of creating a new service attachment subnet. This is required when using BYO VPCs. It must be in the same region as the api-int load balancer, be configured with a purpose of 'Private Service Connect', and have sufficient routing and firewall rules to access the api-int load balancer.",
																MarkdownDescription: "Existing specifies a pre-existing subnet to use instead of creating a new service attachment subnet. This is required when using BYO VPCs. It must be in the same region as the api-int load balancer, be configured with a purpose of 'Private Service Connect', and have sufficient routing and firewall rules to access the api-int load balancer.",
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "Name specifies the name of the existing subnet.",
																		MarkdownDescription: "Name specifies the name of the existing subnet.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"project": schema.StringAttribute{
																		Description:         "Project specifies the project the subnet exists in. This is required for Shared VPC.",
																		MarkdownDescription: "Project specifies the project the subnet exists in. This is required for Shared VPC.",
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

									"region": schema.StringAttribute{
										Description:         "Region specifies the GCP region where the cluster will be created.",
										MarkdownDescription: "Region specifies the GCP region where the cluster will be created.",
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
								Description:         "IBMCloud is the configuration used when installing on IBM Cloud",
								MarkdownDescription: "IBMCloud is the configuration used when installing on IBM Cloud",
								Attributes: map[string]schema.Attribute{
									"account_id": schema.StringAttribute{
										Description:         "AccountID is the IBM Cloud Account ID. AccountID is DEPRECATED and is gathered via the IBM Cloud API for the provided credentials. This field will be ignored.",
										MarkdownDescription: "AccountID is the IBM Cloud Account ID. AccountID is DEPRECATED and is gathered via the IBM Cloud API for the provided credentials. This field will be ignored.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"cis_instance_crn": schema.StringAttribute{
										Description:         "CISInstanceCRN is the IBM Cloud Internet Services Instance CRN CISInstanceCRN is DEPRECATED and gathered via the IBM Cloud API for the provided credentials and cluster deployment base domain. This field will be ignored.",
										MarkdownDescription: "CISInstanceCRN is the IBM Cloud Internet Services Instance CRN CISInstanceCRN is DEPRECATED and gathered via the IBM Cloud API for the provided credentials and cluster deployment base domain. This field will be ignored.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"credentials_secret_ref": schema.SingleNestedAttribute{
										Description:         "CredentialsSecretRef refers to a secret that contains IBM Cloud account access credentials.",
										MarkdownDescription: "CredentialsSecretRef refers to a secret that contains IBM Cloud account access credentials.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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
										Description:         "Region specifies the IBM Cloud region where the cluster will be created.",
										MarkdownDescription: "Region specifies the IBM Cloud region where the cluster will be created.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"none": schema.MapAttribute{
								Description:         "None indicates platform-agnostic install. https://docs.openshift.com/container-platform/4.7/installing/installing_platform_agnostic/installing-platform-agnostic.html",
								MarkdownDescription: "None indicates platform-agnostic install. https://docs.openshift.com/container-platform/4.7/installing/installing_platform_agnostic/installing-platform-agnostic.html",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"openstack": schema.SingleNestedAttribute{
								Description:         "OpenStack is the configuration used when installing on OpenStack",
								MarkdownDescription: "OpenStack is the configuration used when installing on OpenStack",
								Attributes: map[string]schema.Attribute{
									"certificates_secret_ref": schema.SingleNestedAttribute{
										Description:         "CertificatesSecretRef refers to a secret that contains CA certificates necessary for communicating with the OpenStack. There is additional configuration required for the OpenShift cluster to trust the certificates provided in this secret. The 'clouds.yaml' file included in the credentialsSecretRef Secret must also include a reference to the certificate bundle file for the OpenShift cluster being created to trust the OpenStack endpoints. The 'clouds.yaml' file must set the 'cacert' field to either '/etc/openstack-ca/<key name containing the trust bundle in credentialsSecretRef Secret>' or '/etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem'. For example, '''clouds.yaml clouds: shiftstack: auth: ... cacert: '/etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem' '''",
										MarkdownDescription: "CertificatesSecretRef refers to a secret that contains CA certificates necessary for communicating with the OpenStack. There is additional configuration required for the OpenShift cluster to trust the certificates provided in this secret. The 'clouds.yaml' file included in the credentialsSecretRef Secret must also include a reference to the certificate bundle file for the OpenShift cluster being created to trust the OpenStack endpoints. The 'clouds.yaml' file must set the 'cacert' field to either '/etc/openstack-ca/<key name containing the trust bundle in credentialsSecretRef Secret>' or '/etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem'. For example, '''clouds.yaml clouds: shiftstack: auth: ... cacert: '/etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem' '''",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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
										Description:         "Cloud will be used to indicate the OS_CLOUD value to use the right section from the clouds.yaml in the CredentialsSecretRef.",
										MarkdownDescription: "Cloud will be used to indicate the OS_CLOUD value to use the right section from the clouds.yaml in the CredentialsSecretRef.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"credentials_secret_ref": schema.SingleNestedAttribute{
										Description:         "CredentialsSecretRef refers to a secret that contains the OpenStack account access credentials.",
										MarkdownDescription: "CredentialsSecretRef refers to a secret that contains the OpenStack account access credentials.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"trunk_support": schema.BoolAttribute{
										Description:         "TrunkSupport indicates whether or not to use trunk ports in your OpenShift cluster.",
										MarkdownDescription: "TrunkSupport indicates whether or not to use trunk ports in your OpenShift cluster.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
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
												Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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
										Description:         "CredentialsSecretRef refers to a secret that contains the oVirt account access credentials with fields: ovirt_url, ovirt_username, ovirt_password, ovirt_ca_bundle",
										MarkdownDescription: "CredentialsSecretRef refers to a secret that contains the oVirt account access credentials with fields: ovirt_url, ovirt_username, ovirt_password, ovirt_ca_bundle",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"ovirt_cluster_id": schema.StringAttribute{
										Description:         "The target cluster under which all VMs will run",
										MarkdownDescription: "The target cluster under which all VMs will run",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"ovirt_network_name": schema.StringAttribute{
										Description:         "The target network of all the network interfaces of the nodes. Omitting defaults to ovirtmgmt network which is a default network for evert ovirt cluster.",
										MarkdownDescription: "The target network of all the network interfaces of the nodes. Omitting defaults to ovirtmgmt network which is a default network for evert ovirt cluster.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage_domain_id": schema.StringAttribute{
										Description:         "The target storage domain under which all VM disk would be created.",
										MarkdownDescription: "The target storage domain under which all VM disk would be created.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
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
												Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"cluster": schema.StringAttribute{
										Description:         "Cluster is the name of the cluster virtual machines will be cloned into.",
										MarkdownDescription: "Cluster is the name of the cluster virtual machines will be cloned into.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"credentials_secret_ref": schema.SingleNestedAttribute{
										Description:         "CredentialsSecretRef refers to a secret that contains the vSphere account access credentials: GOVC_USERNAME, GOVC_PASSWORD fields.",
										MarkdownDescription: "CredentialsSecretRef refers to a secret that contains the vSphere account access credentials: GOVC_USERNAME, GOVC_PASSWORD fields.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"datacenter": schema.StringAttribute{
										Description:         "Datacenter is the name of the datacenter to use in the vCenter.",
										MarkdownDescription: "Datacenter is the name of the datacenter to use in the vCenter.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"default_datastore": schema.StringAttribute{
										Description:         "DefaultDatastore is the default datastore to use for provisioning volumes.",
										MarkdownDescription: "DefaultDatastore is the default datastore to use for provisioning volumes.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"folder": schema.StringAttribute{
										Description:         "Folder is the name of the folder that will be used and/or created for virtual machines.",
										MarkdownDescription: "Folder is the name of the folder that will be used and/or created for virtual machines.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"network": schema.StringAttribute{
										Description:         "Network specifies the name of the network to be used by the cluster.",
										MarkdownDescription: "Network specifies the name of the network to be used by the cluster.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"v_center": schema.StringAttribute{
										Description:         "VCenter is the domain name or IP address of the vCenter.",
										MarkdownDescription: "VCenter is the domain name or IP address of the vCenter.",
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
						Required: true,
						Optional: false,
						Computed: false,
					},

					"power_state": schema.StringAttribute{
						Description:         "PowerState indicates whether a cluster should be running or hibernating. When omitted, PowerState defaults to the Running state.",
						MarkdownDescription: "PowerState indicates whether a cluster should be running or hibernating. When omitted, PowerState defaults to the Running state.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("", "Running", "Hibernating"),
						},
					},

					"preserve_on_delete": schema.BoolAttribute{
						Description:         "PreserveOnDelete allows the user to disconnect a cluster from Hive without deprovisioning it. This can also be used to abandon ongoing cluster deprovision.",
						MarkdownDescription: "PreserveOnDelete allows the user to disconnect a cluster from Hive without deprovisioning it. This can also be used to abandon ongoing cluster deprovision.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"provisioning": schema.SingleNestedAttribute{
						Description:         "Provisioning contains settings used only for initial cluster provisioning. May be unset in the case of adopted clusters.",
						MarkdownDescription: "Provisioning contains settings used only for initial cluster provisioning. May be unset in the case of adopted clusters.",
						Attributes: map[string]schema.Attribute{
							"image_set_ref": schema.SingleNestedAttribute{
								Description:         "ImageSetRef is a reference to a ClusterImageSet. If a value is specified for ReleaseImage, that will take precedence over the one from the ClusterImageSet.",
								MarkdownDescription: "ImageSetRef is a reference to a ClusterImageSet. If a value is specified for ReleaseImage, that will take precedence over the one from the ClusterImageSet.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name is the name of the ClusterImageSet that this refers to",
										MarkdownDescription: "Name is the name of the ClusterImageSet that this refers to",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"install_config_secret_ref": schema.SingleNestedAttribute{
								Description:         "InstallConfigSecretRef is the reference to a secret that contains an openshift-install InstallConfig. This file will be passed through directly to the installer. Any version of InstallConfig can be used, provided it can be parsed by the openshift-install version for the release you are provisioning.",
								MarkdownDescription: "InstallConfigSecretRef is the reference to a secret that contains an openshift-install InstallConfig. This file will be passed through directly to the installer. Any version of InstallConfig can be used, provided it can be parsed by the openshift-install version for the release you are provisioning.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"installer_env": schema.ListNestedAttribute{
								Description:         "InstallerEnv are extra environment variables to pass through to the installer. This may be used to enable additional features of the installer.",
								MarkdownDescription: "InstallerEnv are extra environment variables to pass through to the installer. This may be used to enable additional features of the installer.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
											MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
											MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value_from": schema.SingleNestedAttribute{
											Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
											MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
													Description:         "Selects a key of a ConfigMap.",
													MarkdownDescription: "Selects a key of a ConfigMap.",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key to select.",
															MarkdownDescription: "The key to select.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
															MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"optional": schema.BoolAttribute{
															Description:         "Specify whether the ConfigMap or its key must be defined",
															MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"field_ref": schema.SingleNestedAttribute{
													Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
													MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
													Attributes: map[string]schema.Attribute{
														"api_version": schema.StringAttribute{
															Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
															MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"field_path": schema.StringAttribute{
															Description:         "Path of the field to select in the specified API version.",
															MarkdownDescription: "Path of the field to select in the specified API version.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"resource_field_ref": schema.SingleNestedAttribute{
													Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
													MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
													Attributes: map[string]schema.Attribute{
														"container_name": schema.StringAttribute{
															Description:         "Container name: required for volumes, optional for env vars",
															MarkdownDescription: "Container name: required for volumes, optional for env vars",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"divisor": schema.StringAttribute{
															Description:         "Specifies the output format of the exposed resources, defaults to '1'",
															MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"resource": schema.StringAttribute{
															Description:         "Required: resource to select",
															MarkdownDescription: "Required: resource to select",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"secret_key_ref": schema.SingleNestedAttribute{
													Description:         "Selects a key of a secret in the pod's namespace",
													MarkdownDescription: "Selects a key of a secret in the pod's namespace",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key of the secret to select from. Must be a valid secret key.",
															MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
															MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"optional": schema.BoolAttribute{
															Description:         "Specify whether the Secret or its key must be defined",
															MarkdownDescription: "Specify whether the Secret or its key must be defined",
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
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"installer_image_override": schema.StringAttribute{
								Description:         "InstallerImageOverride allows specifying a URI for the installer image, normally gleaned from the metadata within the ReleaseImage.",
								MarkdownDescription: "InstallerImageOverride allows specifying a URI for the installer image, normally gleaned from the metadata within the ReleaseImage.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"manifests_config_map_ref": schema.SingleNestedAttribute{
								Description:         "ManifestsConfigMapRef is a reference to user-provided manifests to add to or replace manifests that are generated by the installer. It serves the same purpose as, and is mutually exclusive with, ManifestsSecretRef.",
								MarkdownDescription: "ManifestsConfigMapRef is a reference to user-provided manifests to add to or replace manifests that are generated by the installer. It serves the same purpose as, and is mutually exclusive with, ManifestsSecretRef.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"manifests_secret_ref": schema.SingleNestedAttribute{
								Description:         "ManifestsSecretRef is a reference to user-provided manifests to add to or replace manifests that are generated by the installer. It serves the same purpose as, and is mutually exclusive with, ManifestsConfigMapRef.",
								MarkdownDescription: "ManifestsSecretRef is a reference to user-provided manifests to add to or replace manifests that are generated by the installer. It serves the same purpose as, and is mutually exclusive with, ManifestsConfigMapRef.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"release_image": schema.StringAttribute{
								Description:         "ReleaseImage is the image containing metadata for all components that run in the cluster, and is the primary and best way to specify what specific version of OpenShift you wish to install.",
								MarkdownDescription: "ReleaseImage is the image containing metadata for all components that run in the cluster, and is the primary and best way to specify what specific version of OpenShift you wish to install.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ssh_known_hosts": schema.ListAttribute{
								Description:         "SSHKnownHosts are known hosts to be configured in the hive install manager pod to avoid ssh prompts. Use of ssh in the install pod is somewhat limited today (failure log gathering from cluster, some bare metal provisioning scenarios), so this setting is often not needed.",
								MarkdownDescription: "SSHKnownHosts are known hosts to be configured in the hive install manager pod to avoid ssh prompts. Use of ssh in the install pod is somewhat limited today (failure log gathering from cluster, some bare metal provisioning scenarios), so this setting is often not needed.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ssh_private_key_secret_ref": schema.SingleNestedAttribute{
								Description:         "SSHPrivateKeySecretRef is the reference to the secret that contains the private SSH key to use for access to compute instances. This private key should correspond to the public key included in the InstallConfig. The private key is used by Hive to gather logs on the target cluster if there are install failures. The SSH private key is expected to be in the secret data under the 'ssh-privatekey' key.",
								MarkdownDescription: "SSHPrivateKeySecretRef is the reference to the secret that contains the private SSH key to use for access to compute instances. This private key should correspond to the public key included in the InstallConfig. The private key is used by Hive to gather logs on the target cluster if there are install failures. The SSH private key is expected to be in the secret data under the 'ssh-privatekey' key.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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

					"pull_secret_ref": schema.SingleNestedAttribute{
						Description:         "PullSecretRef is the reference to the secret to use when pulling images.",
						MarkdownDescription: "PullSecretRef is the reference to the secret to use when pulling images.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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
		},
	}
}

func (r *HiveOpenshiftIoClusterDeploymentV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_hive_openshift_io_cluster_deployment_v1_manifest")

	var model HiveOpenshiftIoClusterDeploymentV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("hive.openshift.io/v1")
	model.Kind = pointer.String("ClusterDeployment")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
