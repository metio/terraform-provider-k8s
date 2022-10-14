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

type HiveOpenshiftIoHiveConfigV1Resource struct{}

var (
	_ resource.Resource = (*HiveOpenshiftIoHiveConfigV1Resource)(nil)
)

type HiveOpenshiftIoHiveConfigV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type HiveOpenshiftIoHiveConfigV1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		AdditionalCertificateAuthoritiesSecretRef *[]struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"additional_certificate_authorities_secret_ref" yaml:"additionalCertificateAuthoritiesSecretRef,omitempty"`

		ArgoCDConfig *struct {
			Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

			Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
		} `tfsdk:"argo_cd_config" yaml:"argoCDConfig,omitempty"`

		AwsPrivateLink *struct {
			AssociatedVPCs *[]struct {
				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" yaml:"credentialsSecretRef,omitempty"`

				Region *string `tfsdk:"region" yaml:"region,omitempty"`

				VpcID *string `tfsdk:"vpc_id" yaml:"vpcID,omitempty"`
			} `tfsdk:"associated_vp_cs" yaml:"associatedVPCs,omitempty"`

			CredentialsSecretRef *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"credentials_secret_ref" yaml:"credentialsSecretRef,omitempty"`

			DnsRecordType *string `tfsdk:"dns_record_type" yaml:"dnsRecordType,omitempty"`

			EndpointVPCInventory *[]struct {
				Region *string `tfsdk:"region" yaml:"region,omitempty"`

				Subnets *[]struct {
					AvailabilityZone *string `tfsdk:"availability_zone" yaml:"availabilityZone,omitempty"`

					SubnetID *string `tfsdk:"subnet_id" yaml:"subnetID,omitempty"`
				} `tfsdk:"subnets" yaml:"subnets,omitempty"`

				VpcID *string `tfsdk:"vpc_id" yaml:"vpcID,omitempty"`
			} `tfsdk:"endpoint_vpc_inventory" yaml:"endpointVPCInventory,omitempty"`
		} `tfsdk:"aws_private_link" yaml:"awsPrivateLink,omitempty"`

		Backup *struct {
			MinBackupPeriodSeconds *int64 `tfsdk:"min_backup_period_seconds" yaml:"minBackupPeriodSeconds,omitempty"`

			Velero *struct {
				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
			} `tfsdk:"velero" yaml:"velero,omitempty"`
		} `tfsdk:"backup" yaml:"backup,omitempty"`

		ControllersConfig *struct {
			Controllers *[]struct {
				Config *struct {
					ClientBurst *int64 `tfsdk:"client_burst" yaml:"clientBurst,omitempty"`

					ClientQPS *int64 `tfsdk:"client_qps" yaml:"clientQPS,omitempty"`

					ConcurrentReconciles *int64 `tfsdk:"concurrent_reconciles" yaml:"concurrentReconciles,omitempty"`

					QueueBurst *int64 `tfsdk:"queue_burst" yaml:"queueBurst,omitempty"`

					QueueQPS *int64 `tfsdk:"queue_qps" yaml:"queueQPS,omitempty"`

					Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`
				} `tfsdk:"config" yaml:"config,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"controllers" yaml:"controllers,omitempty"`

			Default *struct {
				ClientBurst *int64 `tfsdk:"client_burst" yaml:"clientBurst,omitempty"`

				ClientQPS *int64 `tfsdk:"client_qps" yaml:"clientQPS,omitempty"`

				ConcurrentReconciles *int64 `tfsdk:"concurrent_reconciles" yaml:"concurrentReconciles,omitempty"`

				QueueBurst *int64 `tfsdk:"queue_burst" yaml:"queueBurst,omitempty"`

				QueueQPS *int64 `tfsdk:"queue_qps" yaml:"queueQPS,omitempty"`

				Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`
			} `tfsdk:"default" yaml:"default,omitempty"`
		} `tfsdk:"controllers_config" yaml:"controllersConfig,omitempty"`

		DeleteProtection *string `tfsdk:"delete_protection" yaml:"deleteProtection,omitempty"`

		DeprovisionsDisabled *bool `tfsdk:"deprovisions_disabled" yaml:"deprovisionsDisabled,omitempty"`

		DisabledControllers *[]string `tfsdk:"disabled_controllers" yaml:"disabledControllers,omitempty"`

		ExportMetrics *bool `tfsdk:"export_metrics" yaml:"exportMetrics,omitempty"`

		FailedProvisionConfig *struct {
			Aws *struct {
				Bucket *string `tfsdk:"bucket" yaml:"bucket,omitempty"`

				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" yaml:"credentialsSecretRef,omitempty"`

				Region *string `tfsdk:"region" yaml:"region,omitempty"`

				ServiceEndpoint *string `tfsdk:"service_endpoint" yaml:"serviceEndpoint,omitempty"`
			} `tfsdk:"aws" yaml:"aws,omitempty"`

			RetryReasons *[]string `tfsdk:"retry_reasons" yaml:"retryReasons,omitempty"`

			SkipGatherLogs *bool `tfsdk:"skip_gather_logs" yaml:"skipGatherLogs,omitempty"`
		} `tfsdk:"failed_provision_config" yaml:"failedProvisionConfig,omitempty"`

		FeatureGates *struct {
			Custom *struct {
				Enabled *[]string `tfsdk:"enabled" yaml:"enabled,omitempty"`
			} `tfsdk:"custom" yaml:"custom,omitempty"`

			FeatureSet *string `tfsdk:"feature_set" yaml:"featureSet,omitempty"`
		} `tfsdk:"feature_gates" yaml:"featureGates,omitempty"`

		GlobalPullSecretRef *struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"global_pull_secret_ref" yaml:"globalPullSecretRef,omitempty"`

		LogLevel *string `tfsdk:"log_level" yaml:"logLevel,omitempty"`

		MaintenanceMode *bool `tfsdk:"maintenance_mode" yaml:"maintenanceMode,omitempty"`

		ManagedDomains *[]struct {
			Aws *struct {
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

				ResourceGroupName *string `tfsdk:"resource_group_name" yaml:"resourceGroupName,omitempty"`
			} `tfsdk:"azure" yaml:"azure,omitempty"`

			Domains *[]string `tfsdk:"domains" yaml:"domains,omitempty"`

			Gcp *struct {
				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" yaml:"credentialsSecretRef,omitempty"`
			} `tfsdk:"gcp" yaml:"gcp,omitempty"`
		} `tfsdk:"managed_domains" yaml:"managedDomains,omitempty"`

		MetricsConfig *struct {
			MetricsWithDuration *[]struct {
				Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"metrics_with_duration" yaml:"metricsWithDuration,omitempty"`
		} `tfsdk:"metrics_config" yaml:"metricsConfig,omitempty"`

		ReleaseImageVerificationConfigMapRef *struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
		} `tfsdk:"release_image_verification_config_map_ref" yaml:"releaseImageVerificationConfigMapRef,omitempty"`

		ServiceProviderCredentialsConfig *struct {
			Aws *struct {
				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" yaml:"credentialsSecretRef,omitempty"`
			} `tfsdk:"aws" yaml:"aws,omitempty"`
		} `tfsdk:"service_provider_credentials_config" yaml:"serviceProviderCredentialsConfig,omitempty"`

		SyncSetReapplyInterval *string `tfsdk:"sync_set_reapply_interval" yaml:"syncSetReapplyInterval,omitempty"`

		TargetNamespace *string `tfsdk:"target_namespace" yaml:"targetNamespace,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewHiveOpenshiftIoHiveConfigV1Resource() resource.Resource {
	return &HiveOpenshiftIoHiveConfigV1Resource{}
}

func (r *HiveOpenshiftIoHiveConfigV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hive_openshift_io_hive_config_v1"
}

func (r *HiveOpenshiftIoHiveConfigV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "HiveConfig is the Schema for the hives API",
		MarkdownDescription: "HiveConfig is the Schema for the hives API",
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
				Description:         "HiveConfigSpec defines the desired state of Hive",
				MarkdownDescription: "HiveConfigSpec defines the desired state of Hive",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"additional_certificate_authorities_secret_ref": {
						Description:         "AdditionalCertificateAuthoritiesSecretRef is a list of references to secrets in the TargetNamespace that contain an additional Certificate Authority to use when communicating with target clusters. These certificate authorities will be used in addition to any self-signed CA generated by each cluster on installation. The cert data should be stored in the Secret key named 'ca.crt'.",
						MarkdownDescription: "AdditionalCertificateAuthoritiesSecretRef is a list of references to secrets in the TargetNamespace that contain an additional Certificate Authority to use when communicating with target clusters. These certificate authorities will be used in addition to any self-signed CA generated by each cluster on installation. The cert data should be stored in the Secret key named 'ca.crt'.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

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

					"argo_cd_config": {
						Description:         "ArgoCD specifies configuration for ArgoCD integration. If enabled, Hive will automatically add provisioned clusters to ArgoCD, and remove them when they are deprovisioned.",
						MarkdownDescription: "ArgoCD specifies configuration for ArgoCD integration. If enabled, Hive will automatically add provisioned clusters to ArgoCD, and remove them when they are deprovisioned.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"enabled": {
								Description:         "Enabled dictates if ArgoCD gitops integration is enabled. If not specified, the default is disabled.",
								MarkdownDescription: "Enabled dictates if ArgoCD gitops integration is enabled. If not specified, the default is disabled.",

								Type: types.BoolType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"namespace": {
								Description:         "Namespace specifies the namespace where ArgoCD is installed. Used for the location of cluster secrets. Defaults to 'argocd'",
								MarkdownDescription: "Namespace specifies the namespace where ArgoCD is installed. Used for the location of cluster secrets. Defaults to 'argocd'",

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

					"aws_private_link": {
						Description:         "AWSPrivateLink defines the configuration for the aws-private-link controller. It provides 3 major pieces of information required by the controller, 1. The Credentials that should be used to create AWS PrivateLink resources other than     what exist in the customer's account. 2. A list of VPCs that can be used by the controller to choose one to create AWS VPC Endpoints     for the AWS VPC Endpoint Services created for ClusterDeployments in their     corresponding regions. 3. A list of VPCs that should be able to resolve the DNS addresses setup for Private Link.",
						MarkdownDescription: "AWSPrivateLink defines the configuration for the aws-private-link controller. It provides 3 major pieces of information required by the controller, 1. The Credentials that should be used to create AWS PrivateLink resources other than     what exist in the customer's account. 2. A list of VPCs that can be used by the controller to choose one to create AWS VPC Endpoints     for the AWS VPC Endpoint Services created for ClusterDeployments in their     corresponding regions. 3. A list of VPCs that should be able to resolve the DNS addresses setup for Private Link.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"associated_vp_cs": {
								Description:         "AssociatedVPCs is the list of VPCs that should be able to resolve the DNS addresses setup for Private Link. This allows clients in VPC to resolve the AWS PrivateLink address using AWS's default DNS resolver for Private Route53 Hosted Zones.  This list should at minimum include the VPC where the current Hive controller is running.",
								MarkdownDescription: "AssociatedVPCs is the list of VPCs that should be able to resolve the DNS addresses setup for Private Link. This allows clients in VPC to resolve the AWS PrivateLink address using AWS's default DNS resolver for Private Route53 Hosted Zones.  This list should at minimum include the VPC where the current Hive controller is running.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"credentials_secret_ref": {
										Description:         "CredentialsSecretRef references a secret in the TargetNamespace that will be used to authenticate with AWS for associating the VPC with the Private HostedZone created for PrivateLink. When not provided, the common credentials for the controller should be used.",
										MarkdownDescription: "CredentialsSecretRef references a secret in the TargetNamespace that will be used to authenticate with AWS for associating the VPC with the Private HostedZone created for PrivateLink. When not provided, the common credentials for the controller should be used.",

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
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"vpc_id": {
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
								Description:         "CredentialsSecretRef references a secret in the TargetNamespace that will be used to authenticate with AWS for creating the resources for AWS PrivateLink.",
								MarkdownDescription: "CredentialsSecretRef references a secret in the TargetNamespace that will be used to authenticate with AWS for creating the resources for AWS PrivateLink.",

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

							"dns_record_type": {
								Description:         "DNSRecordType defines what type of DNS record should be created in Private Hosted Zone for the customer cluster's API endpoint (which is the VPC Endpoint's regional DNS name).",
								MarkdownDescription: "DNSRecordType defines what type of DNS record should be created in Private Hosted Zone for the customer cluster's API endpoint (which is the VPC Endpoint's regional DNS name).",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("Alias", "ARecord"),
								},
							},

							"endpoint_vpc_inventory": {
								Description:         "EndpointVPCInventory is a list of VPCs and the corresponding subnets in various AWS regions. The controller uses this list to choose a VPC for creating AWS VPC Endpoints. Since the VPC Endpoints must be in the same region as the ClusterDeployment, we must have VPCs in that region to be able to setup Private Link.",
								MarkdownDescription: "EndpointVPCInventory is a list of VPCs and the corresponding subnets in various AWS regions. The controller uses this list to choose a VPC for creating AWS VPC Endpoints. Since the VPC Endpoints must be in the same region as the ClusterDeployment, we must have VPCs in that region to be able to setup Private Link.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"region": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"subnets": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"availability_zone": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"subnet_id": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},

									"vpc_id": {
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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"backup": {
						Description:         "Backup specifies configuration for backup integration. If absent, backup integration will be disabled.",
						MarkdownDescription: "Backup specifies configuration for backup integration. If absent, backup integration will be disabled.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"min_backup_period_seconds": {
								Description:         "MinBackupPeriodSeconds specifies that a minimum of MinBackupPeriodSeconds will occur in between each backup. This is used to rate limit backups. This potentially batches together multiple changes into 1 backup. No backups will be lost as changes that happen during this interval are queued up and will result in a backup happening once the interval has been completed.",
								MarkdownDescription: "MinBackupPeriodSeconds specifies that a minimum of MinBackupPeriodSeconds will occur in between each backup. This is used to rate limit backups. This potentially batches together multiple changes into 1 backup. No backups will be lost as changes that happen during this interval are queued up and will result in a backup happening once the interval has been completed.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"velero": {
								Description:         "Velero specifies configuration for the Velero backup integration.",
								MarkdownDescription: "Velero specifies configuration for the Velero backup integration.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"enabled": {
										Description:         "Enabled dictates if Velero backup integration is enabled. If not specified, the default is disabled.",
										MarkdownDescription: "Enabled dictates if Velero backup integration is enabled. If not specified, the default is disabled.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"namespace": {
										Description:         "Namespace specifies in which namespace velero backup objects should be created. If not specified, the default is a namespace named 'velero'.",
										MarkdownDescription: "Namespace specifies in which namespace velero backup objects should be created. If not specified, the default is a namespace named 'velero'.",

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

					"controllers_config": {
						Description:         "ControllersConfig is used to configure different hive controllers",
						MarkdownDescription: "ControllersConfig is used to configure different hive controllers",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"controllers": {
								Description:         "Controllers contains a list of configurations for different controllers",
								MarkdownDescription: "Controllers contains a list of configurations for different controllers",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"config": {
										Description:         "ControllerConfig contains the configuration for the controller specified by Name field",
										MarkdownDescription: "ControllerConfig contains the configuration for the controller specified by Name field",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"client_burst": {
												Description:         "ClientBurst specifies client rate limiter burst for a controller",
												MarkdownDescription: "ClientBurst specifies client rate limiter burst for a controller",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"client_qps": {
												Description:         "ClientQPS specifies client rate limiter QPS for a controller",
												MarkdownDescription: "ClientQPS specifies client rate limiter QPS for a controller",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"concurrent_reconciles": {
												Description:         "ConcurrentReconciles specifies number of concurrent reconciles for a controller",
												MarkdownDescription: "ConcurrentReconciles specifies number of concurrent reconciles for a controller",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"queue_burst": {
												Description:         "QueueBurst specifies workqueue rate limiter burst for a controller",
												MarkdownDescription: "QueueBurst specifies workqueue rate limiter burst for a controller",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"queue_qps": {
												Description:         "QueueQPS specifies workqueue rate limiter QPS for a controller",
												MarkdownDescription: "QueueQPS specifies workqueue rate limiter QPS for a controller",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"replicas": {
												Description:         "Replicas specifies the number of replicas the specific controller pod should use. This is ONLY for controllers that have been split out into their own pods. This is ignored for all others.",
												MarkdownDescription: "Replicas specifies the number of replicas the specific controller pod should use. This is ONLY for controllers that have been split out into their own pods. This is ignored for all others.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},

									"name": {
										Description:         "Name specifies the name of the controller",
										MarkdownDescription: "Name specifies the name of the controller",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("clusterDeployment", "clusterrelocate", "clusterstate", "clusterversion", "controlPlaneCerts", "dnsendpoint", "dnszone", "remoteingress", "remotemachineset", "machinepool", "syncidentityprovider", "unreachable", "velerobackup", "clusterprovision", "clusterDeprovision", "clusterpool", "clusterpoolnamespace", "hibernation", "clusterclaim", "metrics", "clustersync"),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"default": {
								Description:         "Default specifies default configuration for all the controllers, can be used to override following coded defaults default for concurrent reconciles is 5 default for client qps is 5 default for client burst is 10 default for queue qps is 10 default for queue burst is 100",
								MarkdownDescription: "Default specifies default configuration for all the controllers, can be used to override following coded defaults default for concurrent reconciles is 5 default for client qps is 5 default for client burst is 10 default for queue qps is 10 default for queue burst is 100",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"client_burst": {
										Description:         "ClientBurst specifies client rate limiter burst for a controller",
										MarkdownDescription: "ClientBurst specifies client rate limiter burst for a controller",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"client_qps": {
										Description:         "ClientQPS specifies client rate limiter QPS for a controller",
										MarkdownDescription: "ClientQPS specifies client rate limiter QPS for a controller",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"concurrent_reconciles": {
										Description:         "ConcurrentReconciles specifies number of concurrent reconciles for a controller",
										MarkdownDescription: "ConcurrentReconciles specifies number of concurrent reconciles for a controller",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"queue_burst": {
										Description:         "QueueBurst specifies workqueue rate limiter burst for a controller",
										MarkdownDescription: "QueueBurst specifies workqueue rate limiter burst for a controller",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"queue_qps": {
										Description:         "QueueQPS specifies workqueue rate limiter QPS for a controller",
										MarkdownDescription: "QueueQPS specifies workqueue rate limiter QPS for a controller",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"replicas": {
										Description:         "Replicas specifies the number of replicas the specific controller pod should use. This is ONLY for controllers that have been split out into their own pods. This is ignored for all others.",
										MarkdownDescription: "Replicas specifies the number of replicas the specific controller pod should use. This is ONLY for controllers that have been split out into their own pods. This is ignored for all others.",

										Type: types.Int64Type,

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

					"delete_protection": {
						Description:         "DeleteProtection can be set to 'enabled' to turn on automatic delete protection for ClusterDeployments. When enabled, Hive will add the 'hive.openshift.io/protected-delete' annotation to new ClusterDeployments. Once a ClusterDeployment has been installed, a user must remove the annotation from a ClusterDeployment prior to deleting it.",
						MarkdownDescription: "DeleteProtection can be set to 'enabled' to turn on automatic delete protection for ClusterDeployments. When enabled, Hive will add the 'hive.openshift.io/protected-delete' annotation to new ClusterDeployments. Once a ClusterDeployment has been installed, a user must remove the annotation from a ClusterDeployment prior to deleting it.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("enabled"),
						},
					},

					"deprovisions_disabled": {
						Description:         "DeprovisionsDisabled can be set to true to block deprovision jobs from running.",
						MarkdownDescription: "DeprovisionsDisabled can be set to true to block deprovision jobs from running.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"disabled_controllers": {
						Description:         "DisabledControllers allows selectively disabling Hive controllers by name. The name of an individual controller matches the name of the controller as seen in the Hive logging output.",
						MarkdownDescription: "DisabledControllers allows selectively disabling Hive controllers by name. The name of an individual controller matches the name of the controller as seen in the Hive logging output.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"export_metrics": {
						Description:         "ExportMetrics specifies whether the operator should enable metrics for hive controllers to be extracted for prometheus. When set to true, the operator deploys ServiceMonitors so that the prometheus instances that extract metrics. The operator also sets up RBAC in the TargetNamespace so that openshift prometheus in the cluster can list/access objects required to pull metrics.",
						MarkdownDescription: "ExportMetrics specifies whether the operator should enable metrics for hive controllers to be extracted for prometheus. When set to true, the operator deploys ServiceMonitors so that the prometheus instances that extract metrics. The operator also sets up RBAC in the TargetNamespace so that openshift prometheus in the cluster can list/access objects required to pull metrics.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"failed_provision_config": {
						Description:         "FailedProvisionConfig is used to configure settings related to handling provision failures.",
						MarkdownDescription: "FailedProvisionConfig is used to configure settings related to handling provision failures.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"aws": {
								Description:         "FailedProvisionAWSConfig contains AWS-specific info to upload log files.",
								MarkdownDescription: "FailedProvisionAWSConfig contains AWS-specific info to upload log files.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"bucket": {
										Description:         "Bucket is the S3 bucket to store the logs in.",
										MarkdownDescription: "Bucket is the S3 bucket to store the logs in.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"credentials_secret_ref": {
										Description:         "CredentialsSecretRef references a secret in the TargetNamespace that will be used to authenticate with AWS S3. It will need permission to upload logs to S3. Secret should have keys named aws_access_key_id and aws_secret_access_key that contain the AWS credentials. Example Secret:   data:     aws_access_key_id: minio     aws_secret_access_key: minio123",
										MarkdownDescription: "CredentialsSecretRef references a secret in the TargetNamespace that will be used to authenticate with AWS S3. It will need permission to upload logs to S3. Secret should have keys named aws_access_key_id and aws_secret_access_key that contain the AWS credentials. Example Secret:   data:     aws_access_key_id: minio     aws_secret_access_key: minio123",

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
										Description:         "Region is the AWS region to use for S3 operations. This defaults to us-east-1. For AWS China, use cn-northwest-1.",
										MarkdownDescription: "Region is the AWS region to use for S3 operations. This defaults to us-east-1. For AWS China, use cn-northwest-1.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"service_endpoint": {
										Description:         "ServiceEndpoint is the url to connect to an S3 compatible provider.",
										MarkdownDescription: "ServiceEndpoint is the url to connect to an S3 compatible provider.",

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

							"retry_reasons": {
								Description:         "RetryReasons is a list of installFailingReason strings from the [additional-]install-log-regexes ConfigMaps. If specified, Hive will only retry a failed installation if it results in one of the listed reasons. If omitted (not the same thing as empty!), Hive will retry regardless of the failure reason. (The total number of install attempts is still constrained by ClusterDeployment.Spec.InstallAttemptsLimit.)",
								MarkdownDescription: "RetryReasons is a list of installFailingReason strings from the [additional-]install-log-regexes ConfigMaps. If specified, Hive will only retry a failed installation if it results in one of the listed reasons. If omitted (not the same thing as empty!), Hive will retry regardless of the failure reason. (The total number of install attempts is still constrained by ClusterDeployment.Spec.InstallAttemptsLimit.)",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"skip_gather_logs": {
								Description:         "DEPRECATED: This flag is no longer respected and will be removed in the future.",
								MarkdownDescription: "DEPRECATED: This flag is no longer respected and will be removed in the future.",

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

					"feature_gates": {
						Description:         "FeatureGateSelection allows selecting feature gates for the controller.",
						MarkdownDescription: "FeatureGateSelection allows selecting feature gates for the controller.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"custom": {
								Description:         "custom allows the enabling or disabling of any feature. Because of its nature, this setting cannot be validated.  If you have any typos or accidentally apply invalid combinations might cause unknown behavior. featureSet must equal 'Custom' must be set to use this field.",
								MarkdownDescription: "custom allows the enabling or disabling of any feature. Because of its nature, this setting cannot be validated.  If you have any typos or accidentally apply invalid combinations might cause unknown behavior. featureSet must equal 'Custom' must be set to use this field.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"enabled": {
										Description:         "enabled is a list of all feature gates that you want to force on",
										MarkdownDescription: "enabled is a list of all feature gates that you want to force on",

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

							"feature_set": {
								Description:         "featureSet changes the list of features in the cluster.  The default is empty.  Be very careful adjusting this setting.",
								MarkdownDescription: "featureSet changes the list of features in the cluster.  The default is empty.  Be very careful adjusting this setting.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("", "Custom"),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"global_pull_secret_ref": {
						Description:         "GlobalPullSecretRef is used to specify a pull secret that will be used globally by all of the cluster deployments. For each cluster deployment, the contents of GlobalPullSecret will be merged with the specific pull secret for a cluster deployment(if specified), with precedence given to the contents of the pull secret for the cluster deployment. The global pull secret is assumed to be in the TargetNamespace.",
						MarkdownDescription: "GlobalPullSecretRef is used to specify a pull secret that will be used globally by all of the cluster deployments. For each cluster deployment, the contents of GlobalPullSecret will be merged with the specific pull secret for a cluster deployment(if specified), with precedence given to the contents of the pull secret for the cluster deployment. The global pull secret is assumed to be in the TargetNamespace.",

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

					"log_level": {
						Description:         "LogLevel is the level of logging to use for the Hive controllers. Acceptable levels, from coarsest to finest, are panic, fatal, error, warn, info, debug, and trace. The default level is info.",
						MarkdownDescription: "LogLevel is the level of logging to use for the Hive controllers. Acceptable levels, from coarsest to finest, are panic, fatal, error, warn, info, debug, and trace. The default level is info.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"maintenance_mode": {
						Description:         "MaintenanceMode can be set to true to disable the hive controllers in situations where we need to ensure nothing is running that will add or act upon finalizers on Hive types. This should rarely be needed. Sets replicas to 0 for the hive-controllers deployment to accomplish this.",
						MarkdownDescription: "MaintenanceMode can be set to true to disable the hive controllers in situations where we need to ensure nothing is running that will add or act upon finalizers on Hive types. This should rarely be needed. Sets replicas to 0 for the hive-controllers deployment to accomplish this.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"managed_domains": {
						Description:         "ManagedDomains is the list of DNS domains that are managed by the Hive cluster When specifying 'manageDNS: true' in a ClusterDeployment, the ClusterDeployment's baseDomain should be a direct child of one of these domains, otherwise the ClusterDeployment creation will result in a validation error.",
						MarkdownDescription: "ManagedDomains is the list of DNS domains that are managed by the Hive cluster When specifying 'manageDNS: true' in a ClusterDeployment, the ClusterDeployment's baseDomain should be a direct child of one of these domains, otherwise the ClusterDeployment creation will result in a validation error.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"aws": {
								Description:         "AWS contains AWS-specific settings for external DNS",
								MarkdownDescription: "AWS contains AWS-specific settings for external DNS",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"credentials_secret_ref": {
										Description:         "CredentialsSecretRef references a secret in the TargetNamespace that will be used to authenticate with AWS Route53. It will need permission to manage entries for the domain listed in the parent ManageDNSConfig object. Secret should have AWS keys named 'aws_access_key_id' and 'aws_secret_access_key'.",
										MarkdownDescription: "CredentialsSecretRef references a secret in the TargetNamespace that will be used to authenticate with AWS Route53. It will need permission to manage entries for the domain listed in the parent ManageDNSConfig object. Secret should have AWS keys named 'aws_access_key_id' and 'aws_secret_access_key'.",

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
										Description:         "Region is the AWS region to use for route53 operations. This defaults to us-east-1. For AWS China, use cn-northwest-1.",
										MarkdownDescription: "Region is the AWS region to use for route53 operations. This defaults to us-east-1. For AWS China, use cn-northwest-1.",

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

							"azure": {
								Description:         "Azure contains Azure-specific settings for external DNS",
								MarkdownDescription: "Azure contains Azure-specific settings for external DNS",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"cloud_name": {
										Description:         "CloudName is the name of the Azure cloud environment which can be used to configure the Azure SDK with the appropriate Azure API endpoints. If empty, the value is equal to 'AzurePublicCloud'.",
										MarkdownDescription: "CloudName is the name of the Azure cloud environment which can be used to configure the Azure SDK with the appropriate Azure API endpoints. If empty, the value is equal to 'AzurePublicCloud'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("", "AzurePublicCloud", "AzureUSGovernmentCloud", "AzureChinaCloud", "AzureGermanCloud"),
										},
									},

									"credentials_secret_ref": {
										Description:         "CredentialsSecretRef references a secret in the TargetNamespace that will be used to authenticate with Azure DNS. It wil need permission to manage entries in each of the managed domains listed in the parent ManageDNSConfig object. Secret should have a key named 'osServicePrincipal.json'",
										MarkdownDescription: "CredentialsSecretRef references a secret in the TargetNamespace that will be used to authenticate with Azure DNS. It wil need permission to manage entries in each of the managed domains listed in the parent ManageDNSConfig object. Secret should have a key named 'osServicePrincipal.json'",

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

									"resource_group_name": {
										Description:         "ResourceGroupName specifies the Azure resource group containing the DNS zones for the domains being managed.",
										MarkdownDescription: "ResourceGroupName specifies the Azure resource group containing the DNS zones for the domains being managed.",

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

							"domains": {
								Description:         "Domains is the list of domains that hive will be managing entries for with the provided credentials.",
								MarkdownDescription: "Domains is the list of domains that hive will be managing entries for with the provided credentials.",

								Type: types.ListType{ElemType: types.StringType},

								Required: true,
								Optional: false,
								Computed: false,
							},

							"gcp": {
								Description:         "GCP contains GCP-specific settings for external DNS",
								MarkdownDescription: "GCP contains GCP-specific settings for external DNS",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"credentials_secret_ref": {
										Description:         "CredentialsSecretRef references a secret in the TargetNamespace that will be used to authenticate with GCP DNS. It will need permission to manage entries in each of the managed domains for this cluster. listed in the parent ManageDNSConfig object. Secret should have a key named 'osServiceAccount.json'. The credentials must specify the project to use.",
										MarkdownDescription: "CredentialsSecretRef references a secret in the TargetNamespace that will be used to authenticate with GCP DNS. It will need permission to manage entries in each of the managed domains for this cluster. listed in the parent ManageDNSConfig object. Secret should have a key named 'osServiceAccount.json'. The credentials must specify the project to use.",

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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"metrics_config": {
						Description:         "MetricsConfig encapsulates metrics specific configurations, like opting in for certain metrics.",
						MarkdownDescription: "MetricsConfig encapsulates metrics specific configurations, like opting in for certain metrics.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"metrics_with_duration": {
								Description:         "Optional metrics and their configurations",
								MarkdownDescription: "Optional metrics and their configurations",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"duration": {
										Description:         "Duration is the minimum time taken - the relevant metric will be logged only if the value reported by that metric is more than the time mentioned here. For example, if a user opts-in for current clusters stopping and mentions 1 hour here, only the clusters stopping for more than an hour will be reported. This is a Duration value; see https://pkg.go.dev/time#ParseDuration for accepted formats.",
										MarkdownDescription: "Duration is the minimum time taken - the relevant metric will be logged only if the value reported by that metric is more than the time mentioned here. For example, if a user opts-in for current clusters stopping and mentions 1 hour here, only the clusters stopping for more than an hour will be reported. This is a Duration value; see https://pkg.go.dev/time#ParseDuration for accepted formats.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ns|us|µs|ms|s|m|h))+$`), ""),
										},
									},

									"name": {
										Description:         "Name of the metric. It will correspond to an optional relevant metric in hive",
										MarkdownDescription: "Name of the metric. It will correspond to an optional relevant metric in hive",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("currentStopping", "currentResuming", "currentWaitingForCO", "currentClusterSyncFailing", "cumulativeHibernated", "cumulativeResumed"),
										},
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

					"release_image_verification_config_map_ref": {
						Description:         "ReleaseImageVerificationConfigMapRef is a reference to the ConfigMap that will be used to verify release images.  The config map structure is exactly the same as the config map used for verification of release images for OpenShift 4 during upgrades. Therefore you can usually set this to the config map shipped as part of OpenShift (openshift-config-managed/release-verification).  See https://github.com/openshift/cluster-update-keys for more details. The keys within the config map in the data field define how verification is performed:  verifier-public-key-*: One or more GPG public keys in ASCII form that must have signed the                        release image by digest.  store-*: A URL (scheme file://, http://, or https://) location that contains signatures. These          signatures are in the atomic container signature format. The URL will have the digest          of the image appended to it as '<STORE>/<ALGO>=<DIGEST>/signature-<NUMBER>' as described          in the container image signing format. The docker-image-manifest section of the          signature must match the release image digest. Signatures are searched starting at          NUMBER 1 and incrementing if the signature exists but is not valid. The signature is a          GPG signed and encrypted JSON message. The file store is provided for testing only at          the current time, although future versions of the CVO might allow host mounting of          signatures.  See https://github.com/containers/image/blob/ab49b0a48428c623a8f03b41b9083d48966b34a9/docs/signature-protocols.md for a description of the signature store  The returned verifier will require that any new release image will only be considered verified if each provided public key has signed the release image digest. The signature may be in any store and the lookup order is internally defined.  If not set, no verification will be performed.",
						MarkdownDescription: "ReleaseImageVerificationConfigMapRef is a reference to the ConfigMap that will be used to verify release images.  The config map structure is exactly the same as the config map used for verification of release images for OpenShift 4 during upgrades. Therefore you can usually set this to the config map shipped as part of OpenShift (openshift-config-managed/release-verification).  See https://github.com/openshift/cluster-update-keys for more details. The keys within the config map in the data field define how verification is performed:  verifier-public-key-*: One or more GPG public keys in ASCII form that must have signed the                        release image by digest.  store-*: A URL (scheme file://, http://, or https://) location that contains signatures. These          signatures are in the atomic container signature format. The URL will have the digest          of the image appended to it as '<STORE>/<ALGO>=<DIGEST>/signature-<NUMBER>' as described          in the container image signing format. The docker-image-manifest section of the          signature must match the release image digest. Signatures are searched starting at          NUMBER 1 and incrementing if the signature exists but is not valid. The signature is a          GPG signed and encrypted JSON message. The file store is provided for testing only at          the current time, although future versions of the CVO might allow host mounting of          signatures.  See https://github.com/containers/image/blob/ab49b0a48428c623a8f03b41b9083d48966b34a9/docs/signature-protocols.md for a description of the signature store  The returned verifier will require that any new release image will only be considered verified if each provided public key has signed the release image digest. The signature may be in any store and the lookup order is internally defined.  If not set, no verification will be performed.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name of the ConfigMap",
								MarkdownDescription: "Name of the ConfigMap",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"namespace": {
								Description:         "Namespace of the ConfigMap",
								MarkdownDescription: "Namespace of the ConfigMap",

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

					"service_provider_credentials_config": {
						Description:         "ServiceProviderCredentialsConfig is used to configure credentials related to being a service provider on various cloud platforms.",
						MarkdownDescription: "ServiceProviderCredentialsConfig is used to configure credentials related to being a service provider on various cloud platforms.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"aws": {
								Description:         "AWS is used to configure credentials related to being a service provider on AWS.",
								MarkdownDescription: "AWS is used to configure credentials related to being a service provider on AWS.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"credentials_secret_ref": {
										Description:         "CredentialsSecretRef references a secret in the TargetNamespace that will be used to authenticate with AWS to become the Service Provider. Being a Service Provider allows the controllers to assume the role in customer AWS accounts to manager clusters.",
										MarkdownDescription: "CredentialsSecretRef references a secret in the TargetNamespace that will be used to authenticate with AWS to become the Service Provider. Being a Service Provider allows the controllers to assume the role in customer AWS accounts to manager clusters.",

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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"sync_set_reapply_interval": {
						Description:         "SyncSetReapplyInterval is a string duration indicating how much time must pass before SyncSet resources will be reapplied. The default reapply interval is two hours.",
						MarkdownDescription: "SyncSetReapplyInterval is a string duration indicating how much time must pass before SyncSet resources will be reapplied. The default reapply interval is two hours.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"target_namespace": {
						Description:         "TargetNamespace is the namespace where the core Hive components should be run. Defaults to 'hive'. Will be created if it does not already exist. All resource references in HiveConfig can be assumed to be in the TargetNamespace. NOTE: Whereas it is possible to edit this value, causing hive to 'move' its core components to the new namespace, the old namespace is not deleted, as it will still contain resources created by kubernetes and/or other OpenShift controllers.",
						MarkdownDescription: "TargetNamespace is the namespace where the core Hive components should be run. Defaults to 'hive'. Will be created if it does not already exist. All resource references in HiveConfig can be assumed to be in the TargetNamespace. NOTE: Whereas it is possible to edit this value, causing hive to 'move' its core components to the new namespace, the old namespace is not deleted, as it will still contain resources created by kubernetes and/or other OpenShift controllers.",

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
		},
	}, nil
}

func (r *HiveOpenshiftIoHiveConfigV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_hive_openshift_io_hive_config_v1")

	var state HiveOpenshiftIoHiveConfigV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel HiveOpenshiftIoHiveConfigV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("hive.openshift.io/v1")
	goModel.Kind = utilities.Ptr("HiveConfig")

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

func (r *HiveOpenshiftIoHiveConfigV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_hive_openshift_io_hive_config_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *HiveOpenshiftIoHiveConfigV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_hive_openshift_io_hive_config_v1")

	var state HiveOpenshiftIoHiveConfigV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel HiveOpenshiftIoHiveConfigV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("hive.openshift.io/v1")
	goModel.Kind = utilities.Ptr("HiveConfig")

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

func (r *HiveOpenshiftIoHiveConfigV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_hive_openshift_io_hive_config_v1")
	// NO-OP: Terraform removes the state automatically for us
}
