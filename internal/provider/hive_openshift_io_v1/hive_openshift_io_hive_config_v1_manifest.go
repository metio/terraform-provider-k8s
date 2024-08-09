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
	_ datasource.DataSource = &HiveOpenshiftIoHiveConfigV1Manifest{}
)

func NewHiveOpenshiftIoHiveConfigV1Manifest() datasource.DataSource {
	return &HiveOpenshiftIoHiveConfigV1Manifest{}
}

type HiveOpenshiftIoHiveConfigV1Manifest struct{}

type HiveOpenshiftIoHiveConfigV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		AdditionalCertificateAuthoritiesSecretRef *[]struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"additional_certificate_authorities_secret_ref" json:"additionalCertificateAuthoritiesSecretRef,omitempty"`
		ArgoCDConfig *struct {
			Enabled   *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"argo_cd_config" json:"argoCDConfig,omitempty"`
		AwsPrivateLink *struct {
			AssociatedVPCs *[]struct {
				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" json:"credentialsSecretRef,omitempty"`
				Region *string `tfsdk:"region" json:"region,omitempty"`
				VpcID  *string `tfsdk:"vpc_id" json:"vpcID,omitempty"`
			} `tfsdk:"associated_vp_cs" json:"associatedVPCs,omitempty"`
			CredentialsSecretRef *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"credentials_secret_ref" json:"credentialsSecretRef,omitempty"`
			DnsRecordType        *string `tfsdk:"dns_record_type" json:"dnsRecordType,omitempty"`
			EndpointVPCInventory *[]struct {
				Region  *string `tfsdk:"region" json:"region,omitempty"`
				Subnets *[]struct {
					AvailabilityZone *string `tfsdk:"availability_zone" json:"availabilityZone,omitempty"`
					SubnetID         *string `tfsdk:"subnet_id" json:"subnetID,omitempty"`
				} `tfsdk:"subnets" json:"subnets,omitempty"`
				VpcID *string `tfsdk:"vpc_id" json:"vpcID,omitempty"`
			} `tfsdk:"endpoint_vpc_inventory" json:"endpointVPCInventory,omitempty"`
		} `tfsdk:"aws_private_link" json:"awsPrivateLink,omitempty"`
		Backup *struct {
			MinBackupPeriodSeconds *int64 `tfsdk:"min_backup_period_seconds" json:"minBackupPeriodSeconds,omitempty"`
			Velero                 *struct {
				Enabled   *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"velero" json:"velero,omitempty"`
		} `tfsdk:"backup" json:"backup,omitempty"`
		ControllersConfig *struct {
			Controllers *[]struct {
				Config *struct {
					ClientBurst          *int64 `tfsdk:"client_burst" json:"clientBurst,omitempty"`
					ClientQPS            *int64 `tfsdk:"client_qps" json:"clientQPS,omitempty"`
					ConcurrentReconciles *int64 `tfsdk:"concurrent_reconciles" json:"concurrentReconciles,omitempty"`
					QueueBurst           *int64 `tfsdk:"queue_burst" json:"queueBurst,omitempty"`
					QueueQPS             *int64 `tfsdk:"queue_qps" json:"queueQPS,omitempty"`
					Replicas             *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
					Resources            *struct {
						Claims *[]struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"claims" json:"claims,omitempty"`
						Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
						Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
					} `tfsdk:"resources" json:"resources,omitempty"`
				} `tfsdk:"config" json:"config,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"controllers" json:"controllers,omitempty"`
			Default *struct {
				ClientBurst          *int64 `tfsdk:"client_burst" json:"clientBurst,omitempty"`
				ClientQPS            *int64 `tfsdk:"client_qps" json:"clientQPS,omitempty"`
				ConcurrentReconciles *int64 `tfsdk:"concurrent_reconciles" json:"concurrentReconciles,omitempty"`
				QueueBurst           *int64 `tfsdk:"queue_burst" json:"queueBurst,omitempty"`
				QueueQPS             *int64 `tfsdk:"queue_qps" json:"queueQPS,omitempty"`
				Replicas             *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
				Resources            *struct {
					Claims *[]struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"claims" json:"claims,omitempty"`
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
			} `tfsdk:"default" json:"default,omitempty"`
		} `tfsdk:"controllers_config" json:"controllersConfig,omitempty"`
		DeleteProtection *string `tfsdk:"delete_protection" json:"deleteProtection,omitempty"`
		DeploymentConfig *[]struct {
			DeploymentName *string `tfsdk:"deployment_name" json:"deploymentName,omitempty"`
			Resources      *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
		} `tfsdk:"deployment_config" json:"deploymentConfig,omitempty"`
		DeprovisionsDisabled  *bool     `tfsdk:"deprovisions_disabled" json:"deprovisionsDisabled,omitempty"`
		DisabledControllers   *[]string `tfsdk:"disabled_controllers" json:"disabledControllers,omitempty"`
		ExportMetrics         *bool     `tfsdk:"export_metrics" json:"exportMetrics,omitempty"`
		FailedProvisionConfig *struct {
			Aws *struct {
				Bucket               *string `tfsdk:"bucket" json:"bucket,omitempty"`
				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" json:"credentialsSecretRef,omitempty"`
				Region          *string `tfsdk:"region" json:"region,omitempty"`
				ServiceEndpoint *string `tfsdk:"service_endpoint" json:"serviceEndpoint,omitempty"`
			} `tfsdk:"aws" json:"aws,omitempty"`
			RetryReasons   *[]string `tfsdk:"retry_reasons" json:"retryReasons,omitempty"`
			SkipGatherLogs *bool     `tfsdk:"skip_gather_logs" json:"skipGatherLogs,omitempty"`
		} `tfsdk:"failed_provision_config" json:"failedProvisionConfig,omitempty"`
		FeatureGates *struct {
			Custom *struct {
				Enabled *[]string `tfsdk:"enabled" json:"enabled,omitempty"`
			} `tfsdk:"custom" json:"custom,omitempty"`
			FeatureSet *string `tfsdk:"feature_set" json:"featureSet,omitempty"`
		} `tfsdk:"feature_gates" json:"featureGates,omitempty"`
		GlobalPullSecretRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"global_pull_secret_ref" json:"globalPullSecretRef,omitempty"`
		LogLevel                *string `tfsdk:"log_level" json:"logLevel,omitempty"`
		MachinePoolPollInterval *string `tfsdk:"machine_pool_poll_interval" json:"machinePoolPollInterval,omitempty"`
		MaintenanceMode         *bool   `tfsdk:"maintenance_mode" json:"maintenanceMode,omitempty"`
		ManagedDomains          *[]struct {
			Aws *struct {
				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" json:"credentialsSecretRef,omitempty"`
				Region *string `tfsdk:"region" json:"region,omitempty"`
			} `tfsdk:"aws" json:"aws,omitempty"`
			Azure *struct {
				CloudName            *string `tfsdk:"cloud_name" json:"cloudName,omitempty"`
				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" json:"credentialsSecretRef,omitempty"`
				ResourceGroupName *string `tfsdk:"resource_group_name" json:"resourceGroupName,omitempty"`
			} `tfsdk:"azure" json:"azure,omitempty"`
			Domains *[]string `tfsdk:"domains" json:"domains,omitempty"`
			Gcp     *struct {
				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" json:"credentialsSecretRef,omitempty"`
			} `tfsdk:"gcp" json:"gcp,omitempty"`
		} `tfsdk:"managed_domains" json:"managedDomains,omitempty"`
		MetricsConfig *struct {
			AdditionalClusterDeploymentLabels *map[string]string `tfsdk:"additional_cluster_deployment_labels" json:"additionalClusterDeploymentLabels,omitempty"`
			MetricsWithDuration               *[]struct {
				Duration *string `tfsdk:"duration" json:"duration,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"metrics_with_duration" json:"metricsWithDuration,omitempty"`
		} `tfsdk:"metrics_config" json:"metricsConfig,omitempty"`
		PrivateLink *struct {
			Gcp *struct {
				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" json:"credentialsSecretRef,omitempty"`
				EndpointVPCInventory *[]struct {
					Network *string `tfsdk:"network" json:"network,omitempty"`
					Subnets *[]struct {
						Region *string `tfsdk:"region" json:"region,omitempty"`
						Subnet *string `tfsdk:"subnet" json:"subnet,omitempty"`
					} `tfsdk:"subnets" json:"subnets,omitempty"`
				} `tfsdk:"endpoint_vpc_inventory" json:"endpointVPCInventory,omitempty"`
			} `tfsdk:"gcp" json:"gcp,omitempty"`
		} `tfsdk:"private_link" json:"privateLink,omitempty"`
		ReleaseImageVerificationConfigMapRef *struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"release_image_verification_config_map_ref" json:"releaseImageVerificationConfigMapRef,omitempty"`
		ServiceProviderCredentialsConfig *struct {
			Aws *struct {
				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" json:"credentialsSecretRef,omitempty"`
			} `tfsdk:"aws" json:"aws,omitempty"`
		} `tfsdk:"service_provider_credentials_config" json:"serviceProviderCredentialsConfig,omitempty"`
		SyncSetReapplyInterval *string `tfsdk:"sync_set_reapply_interval" json:"syncSetReapplyInterval,omitempty"`
		TargetNamespace        *string `tfsdk:"target_namespace" json:"targetNamespace,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *HiveOpenshiftIoHiveConfigV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_hive_openshift_io_hive_config_v1_manifest"
}

func (r *HiveOpenshiftIoHiveConfigV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "HiveConfig is the Schema for the hives API",
		MarkdownDescription: "HiveConfig is the Schema for the hives API",
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
				Description:         "HiveConfigSpec defines the desired state of Hive",
				MarkdownDescription: "HiveConfigSpec defines the desired state of Hive",
				Attributes: map[string]schema.Attribute{
					"additional_certificate_authorities_secret_ref": schema.ListNestedAttribute{
						Description:         "AdditionalCertificateAuthoritiesSecretRef is a list of references to secrets in the TargetNamespace that contain an additional Certificate Authority to use when communicating with target clusters. These certificate authorities will be used in addition to any self-signed CA generated by each cluster on installation. The cert data should be stored in the Secret key named 'ca.crt'.",
						MarkdownDescription: "AdditionalCertificateAuthoritiesSecretRef is a list of references to secrets in the TargetNamespace that contain an additional Certificate Authority to use when communicating with target clusters. These certificate authorities will be used in addition to any self-signed CA generated by each cluster on installation. The cert data should be stored in the Secret key named 'ca.crt'.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
									MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

					"argo_cd_config": schema.SingleNestedAttribute{
						Description:         "ArgoCD specifies configuration for ArgoCD integration. If enabled, Hive will automatically add provisioned clusters to ArgoCD, and remove them when they are deprovisioned.",
						MarkdownDescription: "ArgoCD specifies configuration for ArgoCD integration. If enabled, Hive will automatically add provisioned clusters to ArgoCD, and remove them when they are deprovisioned.",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Enabled dictates if ArgoCD gitops integration is enabled. If not specified, the default is disabled.",
								MarkdownDescription: "Enabled dictates if ArgoCD gitops integration is enabled. If not specified, the default is disabled.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace specifies the namespace where ArgoCD is installed. Used for the location of cluster secrets. Defaults to 'argocd'",
								MarkdownDescription: "Namespace specifies the namespace where ArgoCD is installed. Used for the location of cluster secrets. Defaults to 'argocd'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"aws_private_link": schema.SingleNestedAttribute{
						Description:         "AWSPrivateLink defines the configuration for the aws-private-link controller. It provides 3 major pieces of information required by the controller, 1. The Credentials that should be used to create AWS PrivateLink resources other than what exist in the customer's account. 2. A list of VPCs that can be used by the controller to choose one to create AWS VPC Endpoints for the AWS VPC Endpoint Services created for ClusterDeployments in their corresponding regions. 3. A list of VPCs that should be able to resolve the DNS addresses setup for Private Link.",
						MarkdownDescription: "AWSPrivateLink defines the configuration for the aws-private-link controller. It provides 3 major pieces of information required by the controller, 1. The Credentials that should be used to create AWS PrivateLink resources other than what exist in the customer's account. 2. A list of VPCs that can be used by the controller to choose one to create AWS VPC Endpoints for the AWS VPC Endpoint Services created for ClusterDeployments in their corresponding regions. 3. A list of VPCs that should be able to resolve the DNS addresses setup for Private Link.",
						Attributes: map[string]schema.Attribute{
							"associated_vp_cs": schema.ListNestedAttribute{
								Description:         "AssociatedVPCs is the list of VPCs that should be able to resolve the DNS addresses setup for Private Link. This allows clients in VPC to resolve the AWS PrivateLink address using AWS's default DNS resolver for Private Route53 Hosted Zones.  This list should at minimum include the VPC where the current Hive controller is running.",
								MarkdownDescription: "AssociatedVPCs is the list of VPCs that should be able to resolve the DNS addresses setup for Private Link. This allows clients in VPC to resolve the AWS PrivateLink address using AWS's default DNS resolver for Private Route53 Hosted Zones.  This list should at minimum include the VPC where the current Hive controller is running.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"credentials_secret_ref": schema.SingleNestedAttribute{
											Description:         "CredentialsSecretRef references a secret in the TargetNamespace that will be used to authenticate with AWS for associating the VPC with the Private HostedZone created for PrivateLink. When not provided, the common credentials for the controller should be used.",
											MarkdownDescription: "CredentialsSecretRef references a secret in the TargetNamespace that will be used to authenticate with AWS for associating the VPC with the Private HostedZone created for PrivateLink. When not provided, the common credentials for the controller should be used.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
													MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"vpc_id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
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

							"credentials_secret_ref": schema.SingleNestedAttribute{
								Description:         "CredentialsSecretRef references a secret in the TargetNamespace that will be used to authenticate with AWS for creating the resources for AWS PrivateLink.",
								MarkdownDescription: "CredentialsSecretRef references a secret in the TargetNamespace that will be used to authenticate with AWS for creating the resources for AWS PrivateLink.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
										MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
							},

							"dns_record_type": schema.StringAttribute{
								Description:         "DNSRecordType defines what type of DNS record should be created in Private Hosted Zone for the customer cluster's API endpoint (which is the VPC Endpoint's regional DNS name).",
								MarkdownDescription: "DNSRecordType defines what type of DNS record should be created in Private Hosted Zone for the customer cluster's API endpoint (which is the VPC Endpoint's regional DNS name).",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Alias", "ARecord"),
								},
							},

							"endpoint_vpc_inventory": schema.ListNestedAttribute{
								Description:         "EndpointVPCInventory is a list of VPCs and the corresponding subnets in various AWS regions. The controller uses this list to choose a VPC for creating AWS VPC Endpoints. Since the VPC Endpoints must be in the same region as the ClusterDeployment, we must have VPCs in that region to be able to setup Private Link.",
								MarkdownDescription: "EndpointVPCInventory is a list of VPCs and the corresponding subnets in various AWS regions. The controller uses this list to choose a VPC for creating AWS VPC Endpoints. Since the VPC Endpoints must be in the same region as the ClusterDeployment, we must have VPCs in that region to be able to setup Private Link.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"region": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"subnets": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"availability_zone": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"subnet_id": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"vpc_id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"backup": schema.SingleNestedAttribute{
						Description:         "Backup specifies configuration for backup integration. If absent, backup integration will be disabled.",
						MarkdownDescription: "Backup specifies configuration for backup integration. If absent, backup integration will be disabled.",
						Attributes: map[string]schema.Attribute{
							"min_backup_period_seconds": schema.Int64Attribute{
								Description:         "MinBackupPeriodSeconds specifies that a minimum of MinBackupPeriodSeconds will occur in between each backup. This is used to rate limit backups. This potentially batches together multiple changes into 1 backup. No backups will be lost as changes that happen during this interval are queued up and will result in a backup happening once the interval has been completed.",
								MarkdownDescription: "MinBackupPeriodSeconds specifies that a minimum of MinBackupPeriodSeconds will occur in between each backup. This is used to rate limit backups. This potentially batches together multiple changes into 1 backup. No backups will be lost as changes that happen during this interval are queued up and will result in a backup happening once the interval has been completed.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"velero": schema.SingleNestedAttribute{
								Description:         "Velero specifies configuration for the Velero backup integration.",
								MarkdownDescription: "Velero specifies configuration for the Velero backup integration.",
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Description:         "Enabled dictates if Velero backup integration is enabled. If not specified, the default is disabled.",
										MarkdownDescription: "Enabled dictates if Velero backup integration is enabled. If not specified, the default is disabled.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "Namespace specifies in which namespace velero backup objects should be created. If not specified, the default is a namespace named 'velero'.",
										MarkdownDescription: "Namespace specifies in which namespace velero backup objects should be created. If not specified, the default is a namespace named 'velero'.",
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

					"controllers_config": schema.SingleNestedAttribute{
						Description:         "ControllersConfig is used to configure different hive controllers",
						MarkdownDescription: "ControllersConfig is used to configure different hive controllers",
						Attributes: map[string]schema.Attribute{
							"controllers": schema.ListNestedAttribute{
								Description:         "Controllers contains a list of configurations for different controllers",
								MarkdownDescription: "Controllers contains a list of configurations for different controllers",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"config": schema.SingleNestedAttribute{
											Description:         "ControllerConfig contains the configuration for the controller specified by Name field",
											MarkdownDescription: "ControllerConfig contains the configuration for the controller specified by Name field",
											Attributes: map[string]schema.Attribute{
												"client_burst": schema.Int64Attribute{
													Description:         "ClientBurst specifies client rate limiter burst for a controller",
													MarkdownDescription: "ClientBurst specifies client rate limiter burst for a controller",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"client_qps": schema.Int64Attribute{
													Description:         "ClientQPS specifies client rate limiter QPS for a controller",
													MarkdownDescription: "ClientQPS specifies client rate limiter QPS for a controller",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"concurrent_reconciles": schema.Int64Attribute{
													Description:         "ConcurrentReconciles specifies number of concurrent reconciles for a controller",
													MarkdownDescription: "ConcurrentReconciles specifies number of concurrent reconciles for a controller",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"queue_burst": schema.Int64Attribute{
													Description:         "QueueBurst specifies workqueue rate limiter burst for a controller",
													MarkdownDescription: "QueueBurst specifies workqueue rate limiter burst for a controller",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"queue_qps": schema.Int64Attribute{
													Description:         "QueueQPS specifies workqueue rate limiter QPS for a controller",
													MarkdownDescription: "QueueQPS specifies workqueue rate limiter QPS for a controller",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"replicas": schema.Int64Attribute{
													Description:         "Replicas specifies the number of replicas the specific controller pod should use. This is ONLY for controllers that have been split out into their own pods. This is ignored for all others.",
													MarkdownDescription: "Replicas specifies the number of replicas the specific controller pod should use. This is ONLY for controllers that have been split out into their own pods. This is ignored for all others.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"resources": schema.SingleNestedAttribute{
													Description:         "Resources describes the compute resource requirements of the controller container. This is ONLY for controllers that have been split out into their own pods. This is ignored for all others.",
													MarkdownDescription: "Resources describes the compute resource requirements of the controller container. This is ONLY for controllers that have been split out into their own pods. This is ignored for all others.",
													Attributes: map[string]schema.Attribute{
														"claims": schema.ListNestedAttribute{
															Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
															MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																		MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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

														"limits": schema.MapAttribute{
															Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
															MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"requests": schema.MapAttribute{
															Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
															MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"name": schema.StringAttribute{
											Description:         "Name specifies the name of the controller",
											MarkdownDescription: "Name specifies the name of the controller",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("clusterDeployment", "clusterrelocate", "clusterstate", "clusterversion", "controlPlaneCerts", "dnsendpoint", "dnszone", "remoteingress", "remotemachineset", "machinepool", "syncidentityprovider", "unreachable", "velerobackup", "clusterprovision", "clusterDeprovision", "clusterpool", "clusterpoolnamespace", "hibernation", "clusterclaim", "metrics", "clustersync"),
											},
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"default": schema.SingleNestedAttribute{
								Description:         "Default specifies default configuration for all the controllers, can be used to override following coded defaults default for concurrent reconciles is 5 default for client qps is 5 default for client burst is 10 default for queue qps is 10 default for queue burst is 100",
								MarkdownDescription: "Default specifies default configuration for all the controllers, can be used to override following coded defaults default for concurrent reconciles is 5 default for client qps is 5 default for client burst is 10 default for queue qps is 10 default for queue burst is 100",
								Attributes: map[string]schema.Attribute{
									"client_burst": schema.Int64Attribute{
										Description:         "ClientBurst specifies client rate limiter burst for a controller",
										MarkdownDescription: "ClientBurst specifies client rate limiter burst for a controller",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"client_qps": schema.Int64Attribute{
										Description:         "ClientQPS specifies client rate limiter QPS for a controller",
										MarkdownDescription: "ClientQPS specifies client rate limiter QPS for a controller",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"concurrent_reconciles": schema.Int64Attribute{
										Description:         "ConcurrentReconciles specifies number of concurrent reconciles for a controller",
										MarkdownDescription: "ConcurrentReconciles specifies number of concurrent reconciles for a controller",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"queue_burst": schema.Int64Attribute{
										Description:         "QueueBurst specifies workqueue rate limiter burst for a controller",
										MarkdownDescription: "QueueBurst specifies workqueue rate limiter burst for a controller",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"queue_qps": schema.Int64Attribute{
										Description:         "QueueQPS specifies workqueue rate limiter QPS for a controller",
										MarkdownDescription: "QueueQPS specifies workqueue rate limiter QPS for a controller",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"replicas": schema.Int64Attribute{
										Description:         "Replicas specifies the number of replicas the specific controller pod should use. This is ONLY for controllers that have been split out into their own pods. This is ignored for all others.",
										MarkdownDescription: "Replicas specifies the number of replicas the specific controller pod should use. This is ONLY for controllers that have been split out into their own pods. This is ignored for all others.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "Resources describes the compute resource requirements of the controller container. This is ONLY for controllers that have been split out into their own pods. This is ignored for all others.",
										MarkdownDescription: "Resources describes the compute resource requirements of the controller container. This is ONLY for controllers that have been split out into their own pods. This is ignored for all others.",
										Attributes: map[string]schema.Attribute{
											"claims": schema.ListNestedAttribute{
												Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
												MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
															MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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

											"limits": schema.MapAttribute{
												Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"requests": schema.MapAttribute{
												Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

					"delete_protection": schema.StringAttribute{
						Description:         "DeleteProtection can be set to 'enabled' to turn on automatic delete protection for ClusterDeployments. When enabled, Hive will add the 'hive.openshift.io/protected-delete' annotation to new ClusterDeployments. Once a ClusterDeployment has been installed, a user must remove the annotation from a ClusterDeployment prior to deleting it.",
						MarkdownDescription: "DeleteProtection can be set to 'enabled' to turn on automatic delete protection for ClusterDeployments. When enabled, Hive will add the 'hive.openshift.io/protected-delete' annotation to new ClusterDeployments. Once a ClusterDeployment has been installed, a user must remove the annotation from a ClusterDeployment prior to deleting it.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("enabled"),
						},
					},

					"deployment_config": schema.ListNestedAttribute{
						Description:         "DeploymentConfig is used to configure (pods/containers of) the Deployments generated by hive-operator.",
						MarkdownDescription: "DeploymentConfig is used to configure (pods/containers of) the Deployments generated by hive-operator.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"deployment_name": schema.StringAttribute{
									Description:         "DeploymentName is the name of one of the Deployments/StatefulSets managed by hive-operator. NOTE: At this time each deployment has only one container. In the future, we may provide a way to specify which container this DeploymentConfig will be applied to.",
									MarkdownDescription: "DeploymentName is the name of one of the Deployments/StatefulSets managed by hive-operator. NOTE: At this time each deployment has only one container. In the future, we may provide a way to specify which container this DeploymentConfig will be applied to.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("hive-controllers", "hive-clustersync", "hiveadmission"),
									},
								},

								"resources": schema.SingleNestedAttribute{
									Description:         "Resources allows customization of the resource (memory, CPU, etc.) limits and requests used by containers in the Deployment/StatefulSet named by DeploymentName.",
									MarkdownDescription: "Resources allows customization of the resource (memory, CPU, etc.) limits and requests used by containers in the Deployment/StatefulSet named by DeploymentName.",
									Attributes: map[string]schema.Attribute{
										"claims": schema.ListNestedAttribute{
											Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
											MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
														MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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

										"limits": schema.MapAttribute{
											Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
											MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"requests": schema.MapAttribute{
											Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
											MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"deprovisions_disabled": schema.BoolAttribute{
						Description:         "DeprovisionsDisabled can be set to true to block deprovision jobs from running.",
						MarkdownDescription: "DeprovisionsDisabled can be set to true to block deprovision jobs from running.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"disabled_controllers": schema.ListAttribute{
						Description:         "DisabledControllers allows selectively disabling Hive controllers by name. The name of an individual controller matches the name of the controller as seen in the Hive logging output.",
						MarkdownDescription: "DisabledControllers allows selectively disabling Hive controllers by name. The name of an individual controller matches the name of the controller as seen in the Hive logging output.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"export_metrics": schema.BoolAttribute{
						Description:         "ExportMetrics has been disabled and has no effect. If upgrading from a version where it was active, please be aware of the following in your HiveConfig.Spec.TargetNamespace (default 'hive' if unset): 1) ServiceMonitors named hive-controllers and hive-clustersync; 2) Role and RoleBinding named prometheus-k8s; 3) The 'openshift.io/cluster-monitoring' metadata.label on the Namespace itself. You may wish to delete these resources. Or you may wish to continue using them to enable monitoring in your environment; but be aware that hive will no longer reconcile them.",
						MarkdownDescription: "ExportMetrics has been disabled and has no effect. If upgrading from a version where it was active, please be aware of the following in your HiveConfig.Spec.TargetNamespace (default 'hive' if unset): 1) ServiceMonitors named hive-controllers and hive-clustersync; 2) Role and RoleBinding named prometheus-k8s; 3) The 'openshift.io/cluster-monitoring' metadata.label on the Namespace itself. You may wish to delete these resources. Or you may wish to continue using them to enable monitoring in your environment; but be aware that hive will no longer reconcile them.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"failed_provision_config": schema.SingleNestedAttribute{
						Description:         "FailedProvisionConfig is used to configure settings related to handling provision failures.",
						MarkdownDescription: "FailedProvisionConfig is used to configure settings related to handling provision failures.",
						Attributes: map[string]schema.Attribute{
							"aws": schema.SingleNestedAttribute{
								Description:         "FailedProvisionAWSConfig contains AWS-specific info to upload log files.",
								MarkdownDescription: "FailedProvisionAWSConfig contains AWS-specific info to upload log files.",
								Attributes: map[string]schema.Attribute{
									"bucket": schema.StringAttribute{
										Description:         "Bucket is the S3 bucket to store the logs in.",
										MarkdownDescription: "Bucket is the S3 bucket to store the logs in.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"credentials_secret_ref": schema.SingleNestedAttribute{
										Description:         "CredentialsSecretRef references a secret in the TargetNamespace that will be used to authenticate with AWS S3. It will need permission to upload logs to S3. Secret should have keys named aws_access_key_id and aws_secret_access_key that contain the AWS credentials. Example Secret: data: aws_access_key_id: minio aws_secret_access_key: minio123",
										MarkdownDescription: "CredentialsSecretRef references a secret in the TargetNamespace that will be used to authenticate with AWS S3. It will need permission to upload logs to S3. Secret should have keys named aws_access_key_id and aws_secret_access_key that contain the AWS credentials. Example Secret: data: aws_access_key_id: minio aws_secret_access_key: minio123",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
												MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
										Description:         "Region is the AWS region to use for S3 operations. This defaults to us-east-1. For AWS China, use cn-northwest-1.",
										MarkdownDescription: "Region is the AWS region to use for S3 operations. This defaults to us-east-1. For AWS China, use cn-northwest-1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"service_endpoint": schema.StringAttribute{
										Description:         "ServiceEndpoint is the url to connect to an S3 compatible provider.",
										MarkdownDescription: "ServiceEndpoint is the url to connect to an S3 compatible provider.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"retry_reasons": schema.ListAttribute{
								Description:         "RetryReasons is a list of installFailingReason strings from the [additional-]install-log-regexes ConfigMaps. If specified, Hive will only retry a failed installation if it results in one of the listed reasons. If omitted (not the same thing as empty!), Hive will retry regardless of the failure reason. (The total number of install attempts is still constrained by ClusterDeployment.Spec.InstallAttemptsLimit.)",
								MarkdownDescription: "RetryReasons is a list of installFailingReason strings from the [additional-]install-log-regexes ConfigMaps. If specified, Hive will only retry a failed installation if it results in one of the listed reasons. If omitted (not the same thing as empty!), Hive will retry regardless of the failure reason. (The total number of install attempts is still constrained by ClusterDeployment.Spec.InstallAttemptsLimit.)",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"skip_gather_logs": schema.BoolAttribute{
								Description:         "DEPRECATED: This flag is no longer respected and will be removed in the future.",
								MarkdownDescription: "DEPRECATED: This flag is no longer respected and will be removed in the future.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"feature_gates": schema.SingleNestedAttribute{
						Description:         "FeatureGateSelection allows selecting feature gates for the controller.",
						MarkdownDescription: "FeatureGateSelection allows selecting feature gates for the controller.",
						Attributes: map[string]schema.Attribute{
							"custom": schema.SingleNestedAttribute{
								Description:         "custom allows the enabling or disabling of any feature. Because of its nature, this setting cannot be validated.  If you have any typos or accidentally apply invalid combinations might cause unknown behavior. featureSet must equal 'Custom' must be set to use this field.",
								MarkdownDescription: "custom allows the enabling or disabling of any feature. Because of its nature, this setting cannot be validated.  If you have any typos or accidentally apply invalid combinations might cause unknown behavior. featureSet must equal 'Custom' must be set to use this field.",
								Attributes: map[string]schema.Attribute{
									"enabled": schema.ListAttribute{
										Description:         "enabled is a list of all feature gates that you want to force on",
										MarkdownDescription: "enabled is a list of all feature gates that you want to force on",
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

							"feature_set": schema.StringAttribute{
								Description:         "featureSet changes the list of features in the cluster.  The default is empty.  Be very careful adjusting this setting.",
								MarkdownDescription: "featureSet changes the list of features in the cluster.  The default is empty.  Be very careful adjusting this setting.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("", "Custom"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"global_pull_secret_ref": schema.SingleNestedAttribute{
						Description:         "GlobalPullSecretRef is used to specify a pull secret that will be used globally by all of the cluster deployments. For each cluster deployment, the contents of GlobalPullSecret will be merged with the specific pull secret for a cluster deployment(if specified), with precedence given to the contents of the pull secret for the cluster deployment. The global pull secret is assumed to be in the TargetNamespace.",
						MarkdownDescription: "GlobalPullSecretRef is used to specify a pull secret that will be used globally by all of the cluster deployments. For each cluster deployment, the contents of GlobalPullSecret will be merged with the specific pull secret for a cluster deployment(if specified), with precedence given to the contents of the pull secret for the cluster deployment. The global pull secret is assumed to be in the TargetNamespace.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
								MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"log_level": schema.StringAttribute{
						Description:         "LogLevel is the level of logging to use for the Hive controllers. Acceptable levels, from coarsest to finest, are panic, fatal, error, warn, info, debug, and trace. The default level is info.",
						MarkdownDescription: "LogLevel is the level of logging to use for the Hive controllers. Acceptable levels, from coarsest to finest, are panic, fatal, error, warn, info, debug, and trace. The default level is info.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"machine_pool_poll_interval": schema.StringAttribute{
						Description:         "MachinePoolPollInterval is a string duration indicating how much time must pass before checking whether remote resources related to MachinePools need to be reapplied. Set to zero to disable polling -- we'll only reconcile when hub objects change. The default interval is 30m.",
						MarkdownDescription: "MachinePoolPollInterval is a string duration indicating how much time must pass before checking whether remote resources related to MachinePools need to be reapplied. Set to zero to disable polling -- we'll only reconcile when hub objects change. The default interval is 30m.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"maintenance_mode": schema.BoolAttribute{
						Description:         "MaintenanceMode can be set to true to disable the hive controllers in situations where we need to ensure nothing is running that will add or act upon finalizers on Hive types. This should rarely be needed. Sets replicas to 0 for the hive-controllers deployment to accomplish this.",
						MarkdownDescription: "MaintenanceMode can be set to true to disable the hive controllers in situations where we need to ensure nothing is running that will add or act upon finalizers on Hive types. This should rarely be needed. Sets replicas to 0 for the hive-controllers deployment to accomplish this.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"managed_domains": schema.ListNestedAttribute{
						Description:         "ManagedDomains is the list of DNS domains that are managed by the Hive cluster When specifying 'manageDNS: true' in a ClusterDeployment, the ClusterDeployment's baseDomain should be a direct child of one of these domains, otherwise the ClusterDeployment creation will result in a validation error.",
						MarkdownDescription: "ManagedDomains is the list of DNS domains that are managed by the Hive cluster When specifying 'manageDNS: true' in a ClusterDeployment, the ClusterDeployment's baseDomain should be a direct child of one of these domains, otherwise the ClusterDeployment creation will result in a validation error.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"aws": schema.SingleNestedAttribute{
									Description:         "AWS contains AWS-specific settings for external DNS",
									MarkdownDescription: "AWS contains AWS-specific settings for external DNS",
									Attributes: map[string]schema.Attribute{
										"credentials_secret_ref": schema.SingleNestedAttribute{
											Description:         "CredentialsSecretRef references a secret in the TargetNamespace that will be used to authenticate with AWS Route53. It will need permission to manage entries for the domain listed in the parent ManageDNSConfig object. Secret should have AWS keys named 'aws_access_key_id' and 'aws_secret_access_key'.",
											MarkdownDescription: "CredentialsSecretRef references a secret in the TargetNamespace that will be used to authenticate with AWS Route53. It will need permission to manage entries for the domain listed in the parent ManageDNSConfig object. Secret should have AWS keys named 'aws_access_key_id' and 'aws_secret_access_key'.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
													MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
											Description:         "Region is the AWS region to use for route53 operations. This defaults to us-east-1. For AWS China, use cn-northwest-1.",
											MarkdownDescription: "Region is the AWS region to use for route53 operations. This defaults to us-east-1. For AWS China, use cn-northwest-1.",
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
									Description:         "Azure contains Azure-specific settings for external DNS",
									MarkdownDescription: "Azure contains Azure-specific settings for external DNS",
									Attributes: map[string]schema.Attribute{
										"cloud_name": schema.StringAttribute{
											Description:         "CloudName is the name of the Azure cloud environment which can be used to configure the Azure SDK with the appropriate Azure API endpoints. If empty, the value is equal to 'AzurePublicCloud'.",
											MarkdownDescription: "CloudName is the name of the Azure cloud environment which can be used to configure the Azure SDK with the appropriate Azure API endpoints. If empty, the value is equal to 'AzurePublicCloud'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("", "AzurePublicCloud", "AzureUSGovernmentCloud", "AzureChinaCloud", "AzureGermanCloud"),
											},
										},

										"credentials_secret_ref": schema.SingleNestedAttribute{
											Description:         "CredentialsSecretRef references a secret in the TargetNamespace that will be used to authenticate with Azure DNS. It wil need permission to manage entries in each of the managed domains listed in the parent ManageDNSConfig object. Secret should have a key named 'osServicePrincipal.json'",
											MarkdownDescription: "CredentialsSecretRef references a secret in the TargetNamespace that will be used to authenticate with Azure DNS. It wil need permission to manage entries in each of the managed domains listed in the parent ManageDNSConfig object. Secret should have a key named 'osServicePrincipal.json'",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
													MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
											Description:         "ResourceGroupName specifies the Azure resource group containing the DNS zones for the domains being managed.",
											MarkdownDescription: "ResourceGroupName specifies the Azure resource group containing the DNS zones for the domains being managed.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"domains": schema.ListAttribute{
									Description:         "Domains is the list of domains that hive will be managing entries for with the provided credentials.",
									MarkdownDescription: "Domains is the list of domains that hive will be managing entries for with the provided credentials.",
									ElementType:         types.StringType,
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"gcp": schema.SingleNestedAttribute{
									Description:         "GCP contains GCP-specific settings for external DNS",
									MarkdownDescription: "GCP contains GCP-specific settings for external DNS",
									Attributes: map[string]schema.Attribute{
										"credentials_secret_ref": schema.SingleNestedAttribute{
											Description:         "CredentialsSecretRef references a secret in the TargetNamespace that will be used to authenticate with GCP DNS. It will need permission to manage entries in each of the managed domains for this cluster. listed in the parent ManageDNSConfig object. Secret should have a key named 'osServiceAccount.json'. The credentials must specify the project to use.",
											MarkdownDescription: "CredentialsSecretRef references a secret in the TargetNamespace that will be used to authenticate with GCP DNS. It will need permission to manage entries in each of the managed domains for this cluster. listed in the parent ManageDNSConfig object. Secret should have a key named 'osServiceAccount.json'. The credentials must specify the project to use.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
													MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"metrics_config": schema.SingleNestedAttribute{
						Description:         "MetricsConfig encapsulates metrics specific configurations, like opting in for certain metrics.",
						MarkdownDescription: "MetricsConfig encapsulates metrics specific configurations, like opting in for certain metrics.",
						Attributes: map[string]schema.Attribute{
							"additional_cluster_deployment_labels": schema.MapAttribute{
								Description:         "AdditionalClusterDeploymentLabels allows configuration of additional labels to be applied to certain metrics. The keys can be any string value suitable for a metric label (see https://prometheus.io/docs/concepts/data_model/#metric-names-and-labels). The values can be any ClusterDeployment label key (from metadata.labels). When observing an affected metric, hive will label it with the specified metric key, and copy the value from the specified ClusterDeployment label. For example, including {'ocp_major_version': 'hive.openshift.io/version-major'} will cause affected metrics to include a label key ocp_major_version with the value from the hive.openshift.io/version-major ClusterDeployment label -- e.g. '4'. NOTE: Avoid ClusterDeployment labels whose values are unbounded, such as those representing cluster names or IDs, as these will cause your prometheus database to grow indefinitely. Affected metrics are those whose type implements the metricsWithDynamicLabels interface found in pkg/controller/metrics/metrics_with_dynamic_labels.go",
								MarkdownDescription: "AdditionalClusterDeploymentLabels allows configuration of additional labels to be applied to certain metrics. The keys can be any string value suitable for a metric label (see https://prometheus.io/docs/concepts/data_model/#metric-names-and-labels). The values can be any ClusterDeployment label key (from metadata.labels). When observing an affected metric, hive will label it with the specified metric key, and copy the value from the specified ClusterDeployment label. For example, including {'ocp_major_version': 'hive.openshift.io/version-major'} will cause affected metrics to include a label key ocp_major_version with the value from the hive.openshift.io/version-major ClusterDeployment label -- e.g. '4'. NOTE: Avoid ClusterDeployment labels whose values are unbounded, such as those representing cluster names or IDs, as these will cause your prometheus database to grow indefinitely. Affected metrics are those whose type implements the metricsWithDynamicLabels interface found in pkg/controller/metrics/metrics_with_dynamic_labels.go",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"metrics_with_duration": schema.ListNestedAttribute{
								Description:         "Optional metrics and their configurations",
								MarkdownDescription: "Optional metrics and their configurations",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"duration": schema.StringAttribute{
											Description:         "Duration is the minimum time taken - the relevant metric will be logged only if the value reported by that metric is more than the time mentioned here. For example, if a user opts-in for current clusters stopping and mentions 1 hour here, only the clusters stopping for more than an hour will be reported. This is a Duration value; see https://pkg.go.dev/time#ParseDuration for accepted formats.",
											MarkdownDescription: "Duration is the minimum time taken - the relevant metric will be logged only if the value reported by that metric is more than the time mentioned here. For example, if a user opts-in for current clusters stopping and mentions 1 hour here, only the clusters stopping for more than an hour will be reported. This is a Duration value; see https://pkg.go.dev/time#ParseDuration for accepted formats.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ns|us|µs|ms|s|m|h))+$`), ""),
											},
										},

										"name": schema.StringAttribute{
											Description:         "Name of the metric. It will correspond to an optional relevant metric in hive",
											MarkdownDescription: "Name of the metric. It will correspond to an optional relevant metric in hive",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("currentStopping", "currentResuming", "currentWaitingForCO", "currentClusterSyncFailing", "cumulativeHibernated", "cumulativeResumed"),
											},
										},
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

					"private_link": schema.SingleNestedAttribute{
						Description:         "PrivateLink is used to configure the privatelink controller.",
						MarkdownDescription: "PrivateLink is used to configure the privatelink controller.",
						Attributes: map[string]schema.Attribute{
							"gcp": schema.SingleNestedAttribute{
								Description:         "GCP is the configuration for GCP hub and link resources.",
								MarkdownDescription: "GCP is the configuration for GCP hub and link resources.",
								Attributes: map[string]schema.Attribute{
									"credentials_secret_ref": schema.SingleNestedAttribute{
										Description:         "CredentialsSecretRef references a secret in the TargetNamespace that will be used to authenticate with GCP for creating the resources for GCP Private Service Connect",
										MarkdownDescription: "CredentialsSecretRef references a secret in the TargetNamespace that will be used to authenticate with GCP for creating the resources for GCP Private Service Connect",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
												MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"endpoint_vpc_inventory": schema.ListNestedAttribute{
										Description:         "EndpointVPCInventory is a list of VPCs and the corresponding subnets in various GCP regions. The controller uses this list to choose a VPC for creating GCP Endpoints. Since the VPC Endpoints must be in the same region as the ClusterDeployment, we must have VPCs in that region to be able to setup Private Service Connect.",
										MarkdownDescription: "EndpointVPCInventory is a list of VPCs and the corresponding subnets in various GCP regions. The controller uses this list to choose a VPC for creating GCP Endpoints. Since the VPC Endpoints must be in the same region as the ClusterDeployment, we must have VPCs in that region to be able to setup Private Service Connect.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"network": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"subnets": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"subnet": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},
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

					"release_image_verification_config_map_ref": schema.SingleNestedAttribute{
						Description:         "ReleaseImageVerificationConfigMapRef is a reference to the ConfigMap that will be used to verify release images.  The config map structure is exactly the same as the config map used for verification of release images for OpenShift 4 during upgrades. Therefore you can usually set this to the config map shipped as part of OpenShift (openshift-config-managed/release-verification).  See https://github.com/openshift/cluster-update-keys for more details. The keys within the config map in the data field define how verification is performed:  verifier-public-key-*: One or more GPG public keys in ASCII form that must have signed the release image by digest.  store-*: A URL (scheme file://, http://, or https://) location that contains signatures. These signatures are in the atomic container signature format. The URL will have the digest of the image appended to it as '<STORE>/<ALGO>=<DIGEST>/signature-<NUMBER>' as described in the container image signing format. The docker-image-manifest section of the signature must match the release image digest. Signatures are searched starting at NUMBER 1 and incrementing if the signature exists but is not valid. The signature is a GPG signed and encrypted JSON message. The file store is provided for testing only at the current time, although future versions of the CVO might allow host mounting of signatures.  See https://github.com/containers/image/blob/ab49b0a48428c623a8f03b41b9083d48966b34a9/docs/signature-protocols.md for a description of the signature store  The returned verifier will require that any new release image will only be considered verified if each provided public key has signed the release image digest. The signature may be in any store and the lookup order is internally defined.  If not set, no verification will be performed.",
						MarkdownDescription: "ReleaseImageVerificationConfigMapRef is a reference to the ConfigMap that will be used to verify release images.  The config map structure is exactly the same as the config map used for verification of release images for OpenShift 4 during upgrades. Therefore you can usually set this to the config map shipped as part of OpenShift (openshift-config-managed/release-verification).  See https://github.com/openshift/cluster-update-keys for more details. The keys within the config map in the data field define how verification is performed:  verifier-public-key-*: One or more GPG public keys in ASCII form that must have signed the release image by digest.  store-*: A URL (scheme file://, http://, or https://) location that contains signatures. These signatures are in the atomic container signature format. The URL will have the digest of the image appended to it as '<STORE>/<ALGO>=<DIGEST>/signature-<NUMBER>' as described in the container image signing format. The docker-image-manifest section of the signature must match the release image digest. Signatures are searched starting at NUMBER 1 and incrementing if the signature exists but is not valid. The signature is a GPG signed and encrypted JSON message. The file store is provided for testing only at the current time, although future versions of the CVO might allow host mounting of signatures.  See https://github.com/containers/image/blob/ab49b0a48428c623a8f03b41b9083d48966b34a9/docs/signature-protocols.md for a description of the signature store  The returned verifier will require that any new release image will only be considered verified if each provided public key has signed the release image digest. The signature may be in any store and the lookup order is internally defined.  If not set, no verification will be performed.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the ConfigMap",
								MarkdownDescription: "Name of the ConfigMap",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the ConfigMap",
								MarkdownDescription: "Namespace of the ConfigMap",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"service_provider_credentials_config": schema.SingleNestedAttribute{
						Description:         "ServiceProviderCredentialsConfig is used to configure credentials related to being a service provider on various cloud platforms.",
						MarkdownDescription: "ServiceProviderCredentialsConfig is used to configure credentials related to being a service provider on various cloud platforms.",
						Attributes: map[string]schema.Attribute{
							"aws": schema.SingleNestedAttribute{
								Description:         "AWS is used to configure credentials related to being a service provider on AWS.",
								MarkdownDescription: "AWS is used to configure credentials related to being a service provider on AWS.",
								Attributes: map[string]schema.Attribute{
									"credentials_secret_ref": schema.SingleNestedAttribute{
										Description:         "CredentialsSecretRef references a secret in the TargetNamespace that will be used to authenticate with AWS to become the Service Provider. Being a Service Provider allows the controllers to assume the role in customer AWS accounts to manager clusters.",
										MarkdownDescription: "CredentialsSecretRef references a secret in the TargetNamespace that will be used to authenticate with AWS to become the Service Provider. Being a Service Provider allows the controllers to assume the role in customer AWS accounts to manager clusters.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
												MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

					"sync_set_reapply_interval": schema.StringAttribute{
						Description:         "SyncSetReapplyInterval is a string duration indicating how much time must pass before SyncSet resources will be reapplied. The default reapply interval is two hours.",
						MarkdownDescription: "SyncSetReapplyInterval is a string duration indicating how much time must pass before SyncSet resources will be reapplied. The default reapply interval is two hours.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"target_namespace": schema.StringAttribute{
						Description:         "TargetNamespace is the namespace where the core Hive components should be run. Defaults to 'hive'. Will be created if it does not already exist. All resource references in HiveConfig can be assumed to be in the TargetNamespace. NOTE: Whereas it is possible to edit this value, causing hive to 'move' its core components to the new namespace, the old namespace is not deleted, as it will still contain resources created by kubernetes and/or other OpenShift controllers.",
						MarkdownDescription: "TargetNamespace is the namespace where the core Hive components should be run. Defaults to 'hive'. Will be created if it does not already exist. All resource references in HiveConfig can be assumed to be in the TargetNamespace. NOTE: Whereas it is possible to edit this value, causing hive to 'move' its core components to the new namespace, the old namespace is not deleted, as it will still contain resources created by kubernetes and/or other OpenShift controllers.",
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

func (r *HiveOpenshiftIoHiveConfigV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_hive_openshift_io_hive_config_v1_manifest")

	var model HiveOpenshiftIoHiveConfigV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("hive.openshift.io/v1")
	model.Kind = pointer.String("HiveConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
