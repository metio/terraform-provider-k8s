/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"regexp"

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

type HiveOpenshiftIoClusterDeploymentV1Resource struct{}

var (
	_ resource.Resource = (*HiveOpenshiftIoClusterDeploymentV1Resource)(nil)
)

type HiveOpenshiftIoClusterDeploymentV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type HiveOpenshiftIoClusterDeploymentV1GoModel struct {
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
		BaseDomain *string `tfsdk:"base_domain" yaml:"baseDomain,omitempty"`

		BoundServiceAccountSigningKeySecretRef *struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"bound_service_account_signing_key_secret_ref" yaml:"boundServiceAccountSigningKeySecretRef,omitempty"`

		CertificateBundles *[]struct {
			CertificateSecretRef *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"certificate_secret_ref" yaml:"certificateSecretRef,omitempty"`

			Generate *bool `tfsdk:"generate" yaml:"generate,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"certificate_bundles" yaml:"certificateBundles,omitempty"`

		ClusterInstallRef *struct {
			Group *string `tfsdk:"group" yaml:"group,omitempty"`

			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Version *string `tfsdk:"version" yaml:"version,omitempty"`
		} `tfsdk:"cluster_install_ref" yaml:"clusterInstallRef,omitempty"`

		ClusterMetadata *struct {
			AdminKubeconfigSecretRef *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"admin_kubeconfig_secret_ref" yaml:"adminKubeconfigSecretRef,omitempty"`

			AdminPasswordSecretRef *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"admin_password_secret_ref" yaml:"adminPasswordSecretRef,omitempty"`

			ClusterID *string `tfsdk:"cluster_id" yaml:"clusterID,omitempty"`

			InfraID *string `tfsdk:"infra_id" yaml:"infraID,omitempty"`
		} `tfsdk:"cluster_metadata" yaml:"clusterMetadata,omitempty"`

		ClusterName *string `tfsdk:"cluster_name" yaml:"clusterName,omitempty"`

		ClusterPoolRef *struct {
			ClaimName *string `tfsdk:"claim_name" yaml:"claimName,omitempty"`

			ClaimedTimestamp *string `tfsdk:"claimed_timestamp" yaml:"claimedTimestamp,omitempty"`

			ClusterDeploymentCustomization *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"cluster_deployment_customization" yaml:"clusterDeploymentCustomization,omitempty"`

			Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

			PoolName *string `tfsdk:"pool_name" yaml:"poolName,omitempty"`
		} `tfsdk:"cluster_pool_ref" yaml:"clusterPoolRef,omitempty"`

		ControlPlaneConfig *struct {
			ApiServerIPOverride *string `tfsdk:"api_server_ip_override" yaml:"apiServerIPOverride,omitempty"`

			ApiURLOverride *string `tfsdk:"api_url_override" yaml:"apiURLOverride,omitempty"`

			ServingCertificates *struct {
				Additional *[]struct {
					Domain *string `tfsdk:"domain" yaml:"domain,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"additional" yaml:"additional,omitempty"`

				Default *string `tfsdk:"default" yaml:"default,omitempty"`
			} `tfsdk:"serving_certificates" yaml:"servingCertificates,omitempty"`
		} `tfsdk:"control_plane_config" yaml:"controlPlaneConfig,omitempty"`

		HibernateAfter *string `tfsdk:"hibernate_after" yaml:"hibernateAfter,omitempty"`

		Ingress *[]struct {
			Domain *string `tfsdk:"domain" yaml:"domain,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			NamespaceSelector *struct {
				MatchExpressions *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

				MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
			} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

			RouteSelector *struct {
				MatchExpressions *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

				MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
			} `tfsdk:"route_selector" yaml:"routeSelector,omitempty"`

			ServingCertificate *string `tfsdk:"serving_certificate" yaml:"servingCertificate,omitempty"`
		} `tfsdk:"ingress" yaml:"ingress,omitempty"`

		InstallAttemptsLimit *int64 `tfsdk:"install_attempts_limit" yaml:"installAttemptsLimit,omitempty"`

		Installed *bool `tfsdk:"installed" yaml:"installed,omitempty"`

		ManageDNS *bool `tfsdk:"manage_dns" yaml:"manageDNS,omitempty"`

		Platform *struct {
			AgentBareMetal *struct {
				AgentSelector *struct {
					MatchExpressions *[]struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

						Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
					} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

					MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
				} `tfsdk:"agent_selector" yaml:"agentSelector,omitempty"`
			} `tfsdk:"agent_bare_metal" yaml:"agentBareMetal,omitempty"`

			Alibabacloud *struct {
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

				PrivateLink *struct {
					Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
				} `tfsdk:"private_link" yaml:"privateLink,omitempty"`

				Region *string `tfsdk:"region" yaml:"region,omitempty"`

				UserTags *map[string]string `tfsdk:"user_tags" yaml:"userTags,omitempty"`
			} `tfsdk:"aws" yaml:"aws,omitempty"`

			Azure *struct {
				BaseDomainResourceGroupName *string `tfsdk:"base_domain_resource_group_name" yaml:"baseDomainResourceGroupName,omitempty"`

				CloudName *string `tfsdk:"cloud_name" yaml:"cloudName,omitempty"`

				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" yaml:"credentialsSecretRef,omitempty"`

				Region *string `tfsdk:"region" yaml:"region,omitempty"`
			} `tfsdk:"azure" yaml:"azure,omitempty"`

			Baremetal *struct {
				LibvirtSSHPrivateKeySecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"libvirt_ssh_private_key_secret_ref" yaml:"libvirtSSHPrivateKeySecretRef,omitempty"`
			} `tfsdk:"baremetal" yaml:"baremetal,omitempty"`

			Gcp *struct {
				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" yaml:"credentialsSecretRef,omitempty"`

				Region *string `tfsdk:"region" yaml:"region,omitempty"`
			} `tfsdk:"gcp" yaml:"gcp,omitempty"`

			Ibmcloud *struct {
				AccountID *string `tfsdk:"account_id" yaml:"accountID,omitempty"`

				CisInstanceCRN *string `tfsdk:"cis_instance_crn" yaml:"cisInstanceCRN,omitempty"`

				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" yaml:"credentialsSecretRef,omitempty"`

				Region *string `tfsdk:"region" yaml:"region,omitempty"`
			} `tfsdk:"ibmcloud" yaml:"ibmcloud,omitempty"`

			None *map[string]string `tfsdk:"none" yaml:"none,omitempty"`

			Openstack *struct {
				CertificatesSecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"certificates_secret_ref" yaml:"certificatesSecretRef,omitempty"`

				Cloud *string `tfsdk:"cloud" yaml:"cloud,omitempty"`

				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" yaml:"credentialsSecretRef,omitempty"`

				TrunkSupport *bool `tfsdk:"trunk_support" yaml:"trunkSupport,omitempty"`
			} `tfsdk:"openstack" yaml:"openstack,omitempty"`

			Ovirt *struct {
				CertificatesSecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"certificates_secret_ref" yaml:"certificatesSecretRef,omitempty"`

				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" yaml:"credentialsSecretRef,omitempty"`

				Ovirt_cluster_id *string `tfsdk:"ovirt_cluster_id" yaml:"ovirt_cluster_id,omitempty"`

				Ovirt_network_name *string `tfsdk:"ovirt_network_name" yaml:"ovirt_network_name,omitempty"`

				Storage_domain_id *string `tfsdk:"storage_domain_id" yaml:"storage_domain_id,omitempty"`
			} `tfsdk:"ovirt" yaml:"ovirt,omitempty"`

			Vsphere *struct {
				CertificatesSecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"certificates_secret_ref" yaml:"certificatesSecretRef,omitempty"`

				Cluster *string `tfsdk:"cluster" yaml:"cluster,omitempty"`

				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" yaml:"credentialsSecretRef,omitempty"`

				Datacenter *string `tfsdk:"datacenter" yaml:"datacenter,omitempty"`

				DefaultDatastore *string `tfsdk:"default_datastore" yaml:"defaultDatastore,omitempty"`

				Folder *string `tfsdk:"folder" yaml:"folder,omitempty"`

				Network *string `tfsdk:"network" yaml:"network,omitempty"`

				VCenter *string `tfsdk:"v_center" yaml:"vCenter,omitempty"`
			} `tfsdk:"vsphere" yaml:"vsphere,omitempty"`
		} `tfsdk:"platform" yaml:"platform,omitempty"`

		PowerState *string `tfsdk:"power_state" yaml:"powerState,omitempty"`

		PreserveOnDelete *bool `tfsdk:"preserve_on_delete" yaml:"preserveOnDelete,omitempty"`

		Provisioning *struct {
			ImageSetRef *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"image_set_ref" yaml:"imageSetRef,omitempty"`

			InstallConfigSecretRef *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"install_config_secret_ref" yaml:"installConfigSecretRef,omitempty"`

			InstallerEnv *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`

				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`

					FieldRef *struct {
						ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

						FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`
					} `tfsdk:"field_ref" yaml:"fieldRef,omitempty"`

					ResourceFieldRef *struct {
						ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

						Divisor utilities.IntOrString `tfsdk:"divisor" yaml:"divisor,omitempty"`

						Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
					} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`

					SecretKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
			} `tfsdk:"installer_env" yaml:"installerEnv,omitempty"`

			InstallerImageOverride *string `tfsdk:"installer_image_override" yaml:"installerImageOverride,omitempty"`

			ManifestsConfigMapRef *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"manifests_config_map_ref" yaml:"manifestsConfigMapRef,omitempty"`

			ManifestsSecretRef *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"manifests_secret_ref" yaml:"manifestsSecretRef,omitempty"`

			ReleaseImage *string `tfsdk:"release_image" yaml:"releaseImage,omitempty"`

			SshKnownHosts *[]string `tfsdk:"ssh_known_hosts" yaml:"sshKnownHosts,omitempty"`

			SshPrivateKeySecretRef *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"ssh_private_key_secret_ref" yaml:"sshPrivateKeySecretRef,omitempty"`
		} `tfsdk:"provisioning" yaml:"provisioning,omitempty"`

		PullSecretRef *struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"pull_secret_ref" yaml:"pullSecretRef,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewHiveOpenshiftIoClusterDeploymentV1Resource() resource.Resource {
	return &HiveOpenshiftIoClusterDeploymentV1Resource{}
}

func (r *HiveOpenshiftIoClusterDeploymentV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hive_openshift_io_cluster_deployment_v1"
}

func (r *HiveOpenshiftIoClusterDeploymentV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "ClusterDeployment is the Schema for the clusterdeployments API",
		MarkdownDescription: "ClusterDeployment is the Schema for the clusterdeployments API",
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
				Description:         "ClusterDeploymentSpec defines the desired state of ClusterDeployment",
				MarkdownDescription: "ClusterDeploymentSpec defines the desired state of ClusterDeployment",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"base_domain": {
						Description:         "BaseDomain is the base domain to which the cluster should belong.",
						MarkdownDescription: "BaseDomain is the base domain to which the cluster should belong.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"bound_service_account_signing_key_secret_ref": {
						Description:         "BoundServiceAccountSignkingKeySecretRef refers to a Secret that contains a 'bound-service-account-signing-key.key' data key pointing to the private key that will be used to sign ServiceAccount objects. Primarily used to provision AWS clusters to use Amazon's Security Token Service.",
						MarkdownDescription: "BoundServiceAccountSignkingKeySecretRef refers to a Secret that contains a 'bound-service-account-signing-key.key' data key pointing to the private key that will be used to sign ServiceAccount objects. Primarily used to provision AWS clusters to use Amazon's Security Token Service.",

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

					"certificate_bundles": {
						Description:         "CertificateBundles is a list of certificate bundles associated with this cluster",
						MarkdownDescription: "CertificateBundles is a list of certificate bundles associated with this cluster",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"certificate_secret_ref": {
								Description:         "CertificateSecretRef is the reference to the secret that contains the certificate bundle. If the certificate bundle is to be generated, it will be generated with the name in this reference. Otherwise, it is expected that the secret should exist in the same namespace as the ClusterDeployment",
								MarkdownDescription: "CertificateSecretRef is the reference to the secret that contains the certificate bundle. If the certificate bundle is to be generated, it will be generated with the name in this reference. Otherwise, it is expected that the secret should exist in the same namespace as the ClusterDeployment",

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

							"generate": {
								Description:         "Generate indicates whether this bundle should have real certificates generated for it.",
								MarkdownDescription: "Generate indicates whether this bundle should have real certificates generated for it.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "Name is an identifier that must be unique within the bundle and must be referenced by an ingress or by the control plane serving certs",
								MarkdownDescription: "Name is an identifier that must be unique within the bundle and must be referenced by an ingress or by the control plane serving certs",

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

					"cluster_install_ref": {
						Description:         "ClusterInstallLocalReference provides reference to an object that implements the hivecontract ClusterInstall. The namespace of the object is same as the ClusterDeployment. This cannot be set when Provisioning is also set.",
						MarkdownDescription: "ClusterInstallLocalReference provides reference to an object that implements the hivecontract ClusterInstall. The namespace of the object is same as the ClusterDeployment. This cannot be set when Provisioning is also set.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"group": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"kind": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"version": {
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

					"cluster_metadata": {
						Description:         "ClusterMetadata contains metadata information about the installed cluster.",
						MarkdownDescription: "ClusterMetadata contains metadata information about the installed cluster.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"admin_kubeconfig_secret_ref": {
								Description:         "AdminKubeconfigSecretRef references the secret containing the admin kubeconfig for this cluster.",
								MarkdownDescription: "AdminKubeconfigSecretRef references the secret containing the admin kubeconfig for this cluster.",

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

							"admin_password_secret_ref": {
								Description:         "AdminPasswordSecretRef references the secret containing the admin username/password which can be used to login to this cluster.",
								MarkdownDescription: "AdminPasswordSecretRef references the secret containing the admin username/password which can be used to login to this cluster.",

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

							"cluster_id": {
								Description:         "ClusterID is a globally unique identifier for this cluster generated during installation. Used for reporting metrics among other places.",
								MarkdownDescription: "ClusterID is a globally unique identifier for this cluster generated during installation. Used for reporting metrics among other places.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"infra_id": {
								Description:         "InfraID is an identifier for this cluster generated during installation and used for tagging/naming resources in cloud providers.",
								MarkdownDescription: "InfraID is an identifier for this cluster generated during installation and used for tagging/naming resources in cloud providers.",

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

					"cluster_name": {
						Description:         "ClusterName is the friendly name of the cluster. It is used for subdomains, some resource tagging, and other instances where a friendly name for the cluster is useful.",
						MarkdownDescription: "ClusterName is the friendly name of the cluster. It is used for subdomains, some resource tagging, and other instances where a friendly name for the cluster is useful.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"cluster_pool_ref": {
						Description:         "ClusterPoolRef is a reference to the ClusterPool that this ClusterDeployment originated from.",
						MarkdownDescription: "ClusterPoolRef is a reference to the ClusterPool that this ClusterDeployment originated from.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"claim_name": {
								Description:         "ClaimName is the name of the ClusterClaim that claimed the cluster from the pool.",
								MarkdownDescription: "ClaimName is the name of the ClusterClaim that claimed the cluster from the pool.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"claimed_timestamp": {
								Description:         "ClaimedTimestamp is the time this cluster was assigned to a ClusterClaim. This is only used for ClusterDeployments belonging to ClusterPools.",
								MarkdownDescription: "ClaimedTimestamp is the time this cluster was assigned to a ClusterClaim. This is only used for ClusterDeployments belonging to ClusterPools.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									validators.DateTime64Validator(),
								},
							},

							"cluster_deployment_customization": {
								Description:         "CustomizationRef is the ClusterPool Inventory claimed customization for this ClusterDeployment. The Customization exists in the ClusterPool namespace.",
								MarkdownDescription: "CustomizationRef is the ClusterPool Inventory claimed customization for this ClusterDeployment. The Customization exists in the ClusterPool namespace.",

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

							"namespace": {
								Description:         "Namespace is the namespace where the ClusterPool resides.",
								MarkdownDescription: "Namespace is the namespace where the ClusterPool resides.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"pool_name": {
								Description:         "PoolName is the name of the ClusterPool for which the cluster was created.",
								MarkdownDescription: "PoolName is the name of the ClusterPool for which the cluster was created.",

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

					"control_plane_config": {
						Description:         "ControlPlaneConfig contains additional configuration for the target cluster's control plane",
						MarkdownDescription: "ControlPlaneConfig contains additional configuration for the target cluster's control plane",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"api_server_ip_override": {
								Description:         "APIServerIPOverride is the optional override of the API server IP address. Hive will use this IP address for creating TCP connections. Port from the original API server URL will be used. This field can be used when repointing the APIServer's DNS is not viable option.",
								MarkdownDescription: "APIServerIPOverride is the optional override of the API server IP address. Hive will use this IP address for creating TCP connections. Port from the original API server URL will be used. This field can be used when repointing the APIServer's DNS is not viable option.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"api_url_override": {
								Description:         "APIURLOverride is the optional URL override to which Hive will transition for communication with the API server of the remote cluster. When a remote cluster is created, Hive will initially communicate using the API URL established during installation. If an API URL Override is specified, Hive will periodically attempt to connect to the remote cluster using the override URL. Once Hive has determined that the override URL is active, Hive will use the override URL for further communications with the API server of the remote cluster.",
								MarkdownDescription: "APIURLOverride is the optional URL override to which Hive will transition for communication with the API server of the remote cluster. When a remote cluster is created, Hive will initially communicate using the API URL established during installation. If an API URL Override is specified, Hive will periodically attempt to connect to the remote cluster using the override URL. Once Hive has determined that the override URL is active, Hive will use the override URL for further communications with the API server of the remote cluster.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"serving_certificates": {
								Description:         "ServingCertificates specifies serving certificates for the control plane",
								MarkdownDescription: "ServingCertificates specifies serving certificates for the control plane",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"additional": {
										Description:         "Additional is a list of additional domains and certificates that are also associated with the control plane's api endpoint.",
										MarkdownDescription: "Additional is a list of additional domains and certificates that are also associated with the control plane's api endpoint.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"domain": {
												Description:         "Domain is the domain of the additional control plane certificate",
												MarkdownDescription: "Domain is the domain of the additional control plane certificate",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "Name references a CertificateBundle in the ClusterDeployment.Spec that should be used for this additional certificate.",
												MarkdownDescription: "Name references a CertificateBundle in the ClusterDeployment.Spec that should be used for this additional certificate.",

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

									"default": {
										Description:         "Default references the name of a CertificateBundle in the ClusterDeployment that should be used for the control plane's default endpoint.",
										MarkdownDescription: "Default references the name of a CertificateBundle in the ClusterDeployment that should be used for the control plane's default endpoint.",

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

					"hibernate_after": {
						Description:         "HibernateAfter will transition a cluster to hibernating power state after it has been running for the given duration. The time that a cluster has been running is the time since the cluster was installed or the time since the cluster last came out of hibernation. This is a Duration value; see https://pkg.go.dev/time#ParseDuration for accepted formats. Note: due to discrepancies in validation vs parsing, we use a Pattern instead of 'Format=duration'. See https://bugzilla.redhat.com/show_bug.cgi?id=2050332 https://github.com/kubernetes/apimachinery/issues/131 https://github.com/kubernetes/apiextensions-apiserver/issues/56",
						MarkdownDescription: "HibernateAfter will transition a cluster to hibernating power state after it has been running for the given duration. The time that a cluster has been running is the time since the cluster was installed or the time since the cluster last came out of hibernation. This is a Duration value; see https://pkg.go.dev/time#ParseDuration for accepted formats. Note: due to discrepancies in validation vs parsing, we use a Pattern instead of 'Format=duration'. See https://bugzilla.redhat.com/show_bug.cgi?id=2050332 https://github.com/kubernetes/apimachinery/issues/131 https://github.com/kubernetes/apiextensions-apiserver/issues/56",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ns|us|µs|ms|s|m|h))+$`), ""),
						},
					},

					"ingress": {
						Description:         "Ingress allows defining desired clusteringress/shards to be configured on the cluster.",
						MarkdownDescription: "Ingress allows defining desired clusteringress/shards to be configured on the cluster.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"domain": {
								Description:         "Domain (sometimes referred to as shard) is the full DNS suffix that the resulting IngressController object will service (eg abcd.mycluster.mydomain.com).",
								MarkdownDescription: "Domain (sometimes referred to as shard) is the full DNS suffix that the resulting IngressController object will service (eg abcd.mycluster.mydomain.com).",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"name": {
								Description:         "Name of the ClusterIngress object to create.",
								MarkdownDescription: "Name of the ClusterIngress object to create.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"namespace_selector": {
								Description:         "NamespaceSelector allows filtering the list of namespaces serviced by the ingress controller.",
								MarkdownDescription: "NamespaceSelector allows filtering the list of namespaces serviced by the ingress controller.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"match_expressions": {
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "key is the label key that the selector applies to.",
												MarkdownDescription: "key is the label key that the selector applies to.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"operator": {
												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"values": {
												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"match_labels": {
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"route_selector": {
								Description:         "RouteSelector allows filtering the set of Routes serviced by the ingress controller",
								MarkdownDescription: "RouteSelector allows filtering the set of Routes serviced by the ingress controller",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"match_expressions": {
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "key is the label key that the selector applies to.",
												MarkdownDescription: "key is the label key that the selector applies to.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"operator": {
												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"values": {
												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"match_labels": {
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"serving_certificate": {
								Description:         "ServingCertificate references a CertificateBundle in the ClusterDeployment.Spec that should be used for this Ingress",
								MarkdownDescription: "ServingCertificate references a CertificateBundle in the ClusterDeployment.Spec that should be used for this Ingress",

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

					"install_attempts_limit": {
						Description:         "InstallAttemptsLimit is the maximum number of times Hive will attempt to install the cluster.",
						MarkdownDescription: "InstallAttemptsLimit is the maximum number of times Hive will attempt to install the cluster.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"installed": {
						Description:         "Installed is true if the cluster has been installed",
						MarkdownDescription: "Installed is true if the cluster has been installed",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"manage_dns": {
						Description:         "ManageDNS specifies whether a DNSZone should be created and managed automatically for this ClusterDeployment",
						MarkdownDescription: "ManageDNS specifies whether a DNSZone should be created and managed automatically for this ClusterDeployment",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"platform": {
						Description:         "Platform is the configuration for the specific platform upon which to perform the installation.",
						MarkdownDescription: "Platform is the configuration for the specific platform upon which to perform the installation.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"agent_bare_metal": {
								Description:         "AgentBareMetal is the configuration used when performing an Assisted Agent based installation to bare metal.",
								MarkdownDescription: "AgentBareMetal is the configuration used when performing an Assisted Agent based installation to bare metal.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"agent_selector": {
										Description:         "AgentSelector is a label selector used for associating relevant custom resources with this cluster. (Agent, BareMetalHost, etc)",
										MarkdownDescription: "AgentSelector is a label selector used for associating relevant custom resources with this cluster. (Agent, BareMetalHost, etc)",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"match_expressions": {
												Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "key is the label key that the selector applies to.",
														MarkdownDescription: "key is the label key that the selector applies to.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"operator": {
														Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"values": {
														Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"match_labels": {
												Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

												Type: types.MapType{ElemType: types.StringType},

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

							"alibabacloud": {
								Description:         "AlibabaCloud is the configuration used when installing on Alibaba Cloud",
								MarkdownDescription: "AlibabaCloud is the configuration used when installing on Alibaba Cloud",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"credentials_secret_ref": {
										Description:         "CredentialsSecretRef refers to a secret that contains Alibaba Cloud account access credentials.",
										MarkdownDescription: "CredentialsSecretRef refers to a secret that contains Alibaba Cloud account access credentials.",

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
										Description:         "Region specifies the Alibaba Cloud region where the cluster will be created.",
										MarkdownDescription: "Region specifies the Alibaba Cloud region where the cluster will be created.",

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
								Description:         "AWS is the configuration used when installing on AWS.",
								MarkdownDescription: "AWS is the configuration used when installing on AWS.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"credentials_assume_role": {
										Description:         "CredentialsAssumeRole refers to the IAM role that must be assumed to obtain AWS account access for the cluster operations.",
										MarkdownDescription: "CredentialsAssumeRole refers to the IAM role that must be assumed to obtain AWS account access for the cluster operations.",

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
										Description:         "CredentialsSecretRef refers to a secret that contains the AWS account access credentials.",
										MarkdownDescription: "CredentialsSecretRef refers to a secret that contains the AWS account access credentials.",

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

									"private_link": {
										Description:         "PrivateLink allows uses to enable access to the cluster's API server using AWS PrivateLink. AWS PrivateLink includes a pair of VPC Endpoint Service and VPC Endpoint accross AWS accounts and allows clients to connect to services using AWS's internal networking instead of the Internet.",
										MarkdownDescription: "PrivateLink allows uses to enable access to the cluster's API server using AWS PrivateLink. AWS PrivateLink includes a pair of VPC Endpoint Service and VPC Endpoint accross AWS accounts and allows clients to connect to services using AWS's internal networking instead of the Internet.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"enabled": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"region": {
										Description:         "Region specifies the AWS region where the cluster will be created.",
										MarkdownDescription: "Region specifies the AWS region where the cluster will be created.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"user_tags": {
										Description:         "UserTags specifies additional tags for AWS resources created for the cluster.",
										MarkdownDescription: "UserTags specifies additional tags for AWS resources created for the cluster.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"azure": {
								Description:         "Azure is the configuration used when installing on Azure.",
								MarkdownDescription: "Azure is the configuration used when installing on Azure.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"base_domain_resource_group_name": {
										Description:         "BaseDomainResourceGroupName specifies the resource group where the azure DNS zone for the base domain is found",
										MarkdownDescription: "BaseDomainResourceGroupName specifies the resource group where the azure DNS zone for the base domain is found",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

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
										Description:         "CredentialsSecretRef refers to a secret that contains the Azure account access credentials.",
										MarkdownDescription: "CredentialsSecretRef refers to a secret that contains the Azure account access credentials.",

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
										Description:         "Region specifies the Azure region where the cluster will be created.",
										MarkdownDescription: "Region specifies the Azure region where the cluster will be created.",

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

							"baremetal": {
								Description:         "BareMetal is the configuration used when installing on bare metal.",
								MarkdownDescription: "BareMetal is the configuration used when installing on bare metal.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"libvirt_ssh_private_key_secret_ref": {
										Description:         "LibvirtSSHPrivateKeySecretRef is the reference to the secret that contains the private SSH key to use for access to the libvirt provisioning host. The SSH private key is expected to be in the secret data under the 'ssh-privatekey' key.",
										MarkdownDescription: "LibvirtSSHPrivateKeySecretRef is the reference to the secret that contains the private SSH key to use for access to the libvirt provisioning host. The SSH private key is expected to be in the secret data under the 'ssh-privatekey' key.",

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

							"gcp": {
								Description:         "GCP is the configuration used when installing on Google Cloud Platform.",
								MarkdownDescription: "GCP is the configuration used when installing on Google Cloud Platform.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"credentials_secret_ref": {
										Description:         "CredentialsSecretRef refers to a secret that contains the GCP account access credentials.",
										MarkdownDescription: "CredentialsSecretRef refers to a secret that contains the GCP account access credentials.",

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
										Description:         "Region specifies the GCP region where the cluster will be created.",
										MarkdownDescription: "Region specifies the GCP region where the cluster will be created.",

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
								Description:         "IBMCloud is the configuration used when installing on IBM Cloud",
								MarkdownDescription: "IBMCloud is the configuration used when installing on IBM Cloud",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"account_id": {
										Description:         "AccountID is the IBM Cloud Account ID. AccountID is DEPRECATED and is gathered via the IBM Cloud API for the provided credentials. This field will be ignored.",
										MarkdownDescription: "AccountID is the IBM Cloud Account ID. AccountID is DEPRECATED and is gathered via the IBM Cloud API for the provided credentials. This field will be ignored.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"cis_instance_crn": {
										Description:         "CISInstanceCRN is the IBM Cloud Internet Services Instance CRN CISInstanceCRN is DEPRECATED and gathered via the IBM Cloud API for the provided credentials and cluster deployment base domain. This field will be ignored.",
										MarkdownDescription: "CISInstanceCRN is the IBM Cloud Internet Services Instance CRN CISInstanceCRN is DEPRECATED and gathered via the IBM Cloud API for the provided credentials and cluster deployment base domain. This field will be ignored.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"credentials_secret_ref": {
										Description:         "CredentialsSecretRef refers to a secret that contains IBM Cloud account access credentials.",
										MarkdownDescription: "CredentialsSecretRef refers to a secret that contains IBM Cloud account access credentials.",

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
										Description:         "Region specifies the IBM Cloud region where the cluster will be created.",
										MarkdownDescription: "Region specifies the IBM Cloud region where the cluster will be created.",

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

							"none": {
								Description:         "None indicates platform-agnostic install. https://docs.openshift.com/container-platform/4.7/installing/installing_platform_agnostic/installing-platform-agnostic.html",
								MarkdownDescription: "None indicates platform-agnostic install. https://docs.openshift.com/container-platform/4.7/installing/installing_platform_agnostic/installing-platform-agnostic.html",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"openstack": {
								Description:         "OpenStack is the configuration used when installing on OpenStack",
								MarkdownDescription: "OpenStack is the configuration used when installing on OpenStack",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"certificates_secret_ref": {
										Description:         "CertificatesSecretRef refers to a secret that contains CA certificates necessary for communicating with the OpenStack. There is additional configuration required for the OpenShift cluster to trust the certificates provided in this secret. The 'clouds.yaml' file included in the credentialsSecretRef Secret must also include a reference to the certificate bundle file for the OpenShift cluster being created to trust the OpenStack endpoints. The 'clouds.yaml' file must set the 'cacert' field to either '/etc/openstack-ca/<key name containing the trust bundle in credentialsSecretRef Secret>' or '/etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem'.  For example, '''clouds.yaml clouds:   shiftstack:     auth: ...     cacert: '/etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem' '''",
										MarkdownDescription: "CertificatesSecretRef refers to a secret that contains CA certificates necessary for communicating with the OpenStack. There is additional configuration required for the OpenShift cluster to trust the certificates provided in this secret. The 'clouds.yaml' file included in the credentialsSecretRef Secret must also include a reference to the certificate bundle file for the OpenShift cluster being created to trust the OpenStack endpoints. The 'clouds.yaml' file must set the 'cacert' field to either '/etc/openstack-ca/<key name containing the trust bundle in credentialsSecretRef Secret>' or '/etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem'.  For example, '''clouds.yaml clouds:   shiftstack:     auth: ...     cacert: '/etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem' '''",

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
										Description:         "Cloud will be used to indicate the OS_CLOUD value to use the right section from the clouds.yaml in the CredentialsSecretRef.",
										MarkdownDescription: "Cloud will be used to indicate the OS_CLOUD value to use the right section from the clouds.yaml in the CredentialsSecretRef.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"credentials_secret_ref": {
										Description:         "CredentialsSecretRef refers to a secret that contains the OpenStack account access credentials.",
										MarkdownDescription: "CredentialsSecretRef refers to a secret that contains the OpenStack account access credentials.",

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

									"trunk_support": {
										Description:         "TrunkSupport indicates whether or not to use trunk ports in your OpenShift cluster.",
										MarkdownDescription: "TrunkSupport indicates whether or not to use trunk ports in your OpenShift cluster.",

										Type: types.BoolType,

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
								Description:         "Ovirt is the configuration used when installing on oVirt",
								MarkdownDescription: "Ovirt is the configuration used when installing on oVirt",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"certificates_secret_ref": {
										Description:         "CertificatesSecretRef refers to a secret that contains the oVirt CA certificates necessary for communicating with oVirt.",
										MarkdownDescription: "CertificatesSecretRef refers to a secret that contains the oVirt CA certificates necessary for communicating with oVirt.",

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
										Description:         "CredentialsSecretRef refers to a secret that contains the oVirt account access credentials with fields: ovirt_url, ovirt_username, ovirt_password, ovirt_ca_bundle",
										MarkdownDescription: "CredentialsSecretRef refers to a secret that contains the oVirt account access credentials with fields: ovirt_url, ovirt_username, ovirt_password, ovirt_ca_bundle",

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

									"ovirt_cluster_id": {
										Description:         "The target cluster under which all VMs will run",
										MarkdownDescription: "The target cluster under which all VMs will run",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"ovirt_network_name": {
										Description:         "The target network of all the network interfaces of the nodes. Omitting defaults to ovirtmgmt network which is a default network for evert ovirt cluster.",
										MarkdownDescription: "The target network of all the network interfaces of the nodes. Omitting defaults to ovirtmgmt network which is a default network for evert ovirt cluster.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"storage_domain_id": {
										Description:         "The target storage domain under which all VM disk would be created.",
										MarkdownDescription: "The target storage domain under which all VM disk would be created.",

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

							"vsphere": {
								Description:         "VSphere is the configuration used when installing on vSphere",
								MarkdownDescription: "VSphere is the configuration used when installing on vSphere",

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

									"cluster": {
										Description:         "Cluster is the name of the cluster virtual machines will be cloned into.",
										MarkdownDescription: "Cluster is the name of the cluster virtual machines will be cloned into.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"credentials_secret_ref": {
										Description:         "CredentialsSecretRef refers to a secret that contains the vSphere account access credentials: GOVC_USERNAME, GOVC_PASSWORD fields.",
										MarkdownDescription: "CredentialsSecretRef refers to a secret that contains the vSphere account access credentials: GOVC_USERNAME, GOVC_PASSWORD fields.",

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

									"datacenter": {
										Description:         "Datacenter is the name of the datacenter to use in the vCenter.",
										MarkdownDescription: "Datacenter is the name of the datacenter to use in the vCenter.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"default_datastore": {
										Description:         "DefaultDatastore is the default datastore to use for provisioning volumes.",
										MarkdownDescription: "DefaultDatastore is the default datastore to use for provisioning volumes.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"folder": {
										Description:         "Folder is the name of the folder that will be used and/or created for virtual machines.",
										MarkdownDescription: "Folder is the name of the folder that will be used and/or created for virtual machines.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"network": {
										Description:         "Network specifies the name of the network to be used by the cluster.",
										MarkdownDescription: "Network specifies the name of the network to be used by the cluster.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"v_center": {
										Description:         "VCenter is the domain name or IP address of the vCenter.",
										MarkdownDescription: "VCenter is the domain name or IP address of the vCenter.",

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

						Required: true,
						Optional: false,
						Computed: false,
					},

					"power_state": {
						Description:         "PowerState indicates whether a cluster should be running or hibernating. When omitted, PowerState defaults to the Running state.",
						MarkdownDescription: "PowerState indicates whether a cluster should be running or hibernating. When omitted, PowerState defaults to the Running state.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("", "Running", "Hibernating"),
						},
					},

					"preserve_on_delete": {
						Description:         "PreserveOnDelete allows the user to disconnect a cluster from Hive without deprovisioning it. This can also be used to abandon ongoing cluster deprovision.",
						MarkdownDescription: "PreserveOnDelete allows the user to disconnect a cluster from Hive without deprovisioning it. This can also be used to abandon ongoing cluster deprovision.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"provisioning": {
						Description:         "Provisioning contains settings used only for initial cluster provisioning. May be unset in the case of adopted clusters.",
						MarkdownDescription: "Provisioning contains settings used only for initial cluster provisioning. May be unset in the case of adopted clusters.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"image_set_ref": {
								Description:         "ImageSetRef is a reference to a ClusterImageSet. If a value is specified for ReleaseImage, that will take precedence over the one from the ClusterImageSet.",
								MarkdownDescription: "ImageSetRef is a reference to a ClusterImageSet. If a value is specified for ReleaseImage, that will take precedence over the one from the ClusterImageSet.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name is the name of the ClusterImageSet that this refers to",
										MarkdownDescription: "Name is the name of the ClusterImageSet that this refers to",

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

							"install_config_secret_ref": {
								Description:         "InstallConfigSecretRef is the reference to a secret that contains an openshift-install InstallConfig. This file will be passed through directly to the installer. Any version of InstallConfig can be used, provided it can be parsed by the openshift-install version for the release you are provisioning.",
								MarkdownDescription: "InstallConfigSecretRef is the reference to a secret that contains an openshift-install InstallConfig. This file will be passed through directly to the installer. Any version of InstallConfig can be used, provided it can be parsed by the openshift-install version for the release you are provisioning.",

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

							"installer_env": {
								Description:         "InstallerEnv are extra environment variables to pass through to the installer. This may be used to enable additional features of the installer.",
								MarkdownDescription: "InstallerEnv are extra environment variables to pass through to the installer. This may be used to enable additional features of the installer.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
										MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"value": {
										Description:         "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
										MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"value_from": {
										Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
										MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_map_key_ref": {
												Description:         "Selects a key of a ConfigMap.",
												MarkdownDescription: "Selects a key of a ConfigMap.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key to select.",
														MarkdownDescription: "The key to select.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"optional": {
														Description:         "Specify whether the ConfigMap or its key must be defined",
														MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"field_ref": {
												Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
												MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"api_version": {
														Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
														MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"field_path": {
														Description:         "Path of the field to select in the specified API version.",
														MarkdownDescription: "Path of the field to select in the specified API version.",

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

											"resource_field_ref": {
												Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
												MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"container_name": {
														Description:         "Container name: required for volumes, optional for env vars",
														MarkdownDescription: "Container name: required for volumes, optional for env vars",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"divisor": {
														Description:         "Specifies the output format of the exposed resources, defaults to '1'",
														MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",

														Type: utilities.IntOrStringType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"resource": {
														Description:         "Required: resource to select",
														MarkdownDescription: "Required: resource to select",

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

											"secret_key_ref": {
												Description:         "Selects a key of a secret in the pod's namespace",
												MarkdownDescription: "Selects a key of a secret in the pod's namespace",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"optional": {
														Description:         "Specify whether the Secret or its key must be defined",
														MarkdownDescription: "Specify whether the Secret or its key must be defined",

														Type: types.BoolType,

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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"installer_image_override": {
								Description:         "InstallerImageOverride allows specifying a URI for the installer image, normally gleaned from the metadata within the ReleaseImage.",
								MarkdownDescription: "InstallerImageOverride allows specifying a URI for the installer image, normally gleaned from the metadata within the ReleaseImage.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"manifests_config_map_ref": {
								Description:         "ManifestsConfigMapRef is a reference to user-provided manifests to add to or replace manifests that are generated by the installer. It serves the same purpose as, and is mutually exclusive with, ManifestsSecretRef.",
								MarkdownDescription: "ManifestsConfigMapRef is a reference to user-provided manifests to add to or replace manifests that are generated by the installer. It serves the same purpose as, and is mutually exclusive with, ManifestsSecretRef.",

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

							"manifests_secret_ref": {
								Description:         "ManifestsSecretRef is a reference to user-provided manifests to add to or replace manifests that are generated by the installer. It serves the same purpose as, and is mutually exclusive with, ManifestsConfigMapRef.",
								MarkdownDescription: "ManifestsSecretRef is a reference to user-provided manifests to add to or replace manifests that are generated by the installer. It serves the same purpose as, and is mutually exclusive with, ManifestsConfigMapRef.",

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

							"release_image": {
								Description:         "ReleaseImage is the image containing metadata for all components that run in the cluster, and is the primary and best way to specify what specific version of OpenShift you wish to install.",
								MarkdownDescription: "ReleaseImage is the image containing metadata for all components that run in the cluster, and is the primary and best way to specify what specific version of OpenShift you wish to install.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ssh_known_hosts": {
								Description:         "SSHKnownHosts are known hosts to be configured in the hive install manager pod to avoid ssh prompts. Use of ssh in the install pod is somewhat limited today (failure log gathering from cluster, some bare metal provisioning scenarios), so this setting is often not needed.",
								MarkdownDescription: "SSHKnownHosts are known hosts to be configured in the hive install manager pod to avoid ssh prompts. Use of ssh in the install pod is somewhat limited today (failure log gathering from cluster, some bare metal provisioning scenarios), so this setting is often not needed.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ssh_private_key_secret_ref": {
								Description:         "SSHPrivateKeySecretRef is the reference to the secret that contains the private SSH key to use for access to compute instances. This private key should correspond to the public key included in the InstallConfig. The private key is used by Hive to gather logs on the target cluster if there are install failures. The SSH private key is expected to be in the secret data under the 'ssh-privatekey' key.",
								MarkdownDescription: "SSHPrivateKeySecretRef is the reference to the secret that contains the private SSH key to use for access to compute instances. This private key should correspond to the public key included in the InstallConfig. The private key is used by Hive to gather logs on the target cluster if there are install failures. The SSH private key is expected to be in the secret data under the 'ssh-privatekey' key.",

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

					"pull_secret_ref": {
						Description:         "PullSecretRef is the reference to the secret to use when pulling images.",
						MarkdownDescription: "PullSecretRef is the reference to the secret to use when pulling images.",

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
		},
	}, nil
}

func (r *HiveOpenshiftIoClusterDeploymentV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_hive_openshift_io_cluster_deployment_v1")

	var state HiveOpenshiftIoClusterDeploymentV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel HiveOpenshiftIoClusterDeploymentV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("hive.openshift.io/v1")
	goModel.Kind = utilities.Ptr("ClusterDeployment")

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

func (r *HiveOpenshiftIoClusterDeploymentV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_hive_openshift_io_cluster_deployment_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *HiveOpenshiftIoClusterDeploymentV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_hive_openshift_io_cluster_deployment_v1")

	var state HiveOpenshiftIoClusterDeploymentV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel HiveOpenshiftIoClusterDeploymentV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("hive.openshift.io/v1")
	goModel.Kind = utilities.Ptr("ClusterDeployment")

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

func (r *HiveOpenshiftIoClusterDeploymentV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_hive_openshift_io_cluster_deployment_v1")
	// NO-OP: Terraform removes the state automatically for us
}
