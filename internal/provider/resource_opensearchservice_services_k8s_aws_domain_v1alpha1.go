/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

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

type OpensearchserviceServicesK8SAwsDomainV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*OpensearchserviceServicesK8SAwsDomainV1Alpha1Resource)(nil)
)

type OpensearchserviceServicesK8SAwsDomainV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type OpensearchserviceServicesK8SAwsDomainV1Alpha1GoModel struct {
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
		AccessPolicies *string `tfsdk:"access_policies" yaml:"accessPolicies,omitempty"`

		AdvancedOptions *map[string]string `tfsdk:"advanced_options" yaml:"advancedOptions,omitempty"`

		AdvancedSecurityOptions *struct {
			AnonymousAuthEnabled *bool `tfsdk:"anonymous_auth_enabled" yaml:"anonymousAuthEnabled,omitempty"`

			Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

			InternalUserDatabaseEnabled *bool `tfsdk:"internal_user_database_enabled" yaml:"internalUserDatabaseEnabled,omitempty"`

			MasterUserOptions *struct {
				MasterUserARN *string `tfsdk:"master_user_arn" yaml:"masterUserARN,omitempty"`

				MasterUserName *string `tfsdk:"master_user_name" yaml:"masterUserName,omitempty"`

				MasterUserPassword *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
				} `tfsdk:"master_user_password" yaml:"masterUserPassword,omitempty"`
			} `tfsdk:"master_user_options" yaml:"masterUserOptions,omitempty"`

			SAMLOptions *struct {
				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Idp *struct {
					EntityID *string `tfsdk:"entity_id" yaml:"entityID,omitempty"`

					MetadataContent *string `tfsdk:"metadata_content" yaml:"metadataContent,omitempty"`
				} `tfsdk:"idp" yaml:"idp,omitempty"`

				MasterBackendRole *string `tfsdk:"master_backend_role" yaml:"masterBackendRole,omitempty"`

				MasterUserName *string `tfsdk:"master_user_name" yaml:"masterUserName,omitempty"`

				RolesKey *string `tfsdk:"roles_key" yaml:"rolesKey,omitempty"`

				SessionTimeoutMinutes *int64 `tfsdk:"session_timeout_minutes" yaml:"sessionTimeoutMinutes,omitempty"`

				SubjectKey *string `tfsdk:"subject_key" yaml:"subjectKey,omitempty"`
			} `tfsdk:"s_aml_options" yaml:"sAMLOptions,omitempty"`
		} `tfsdk:"advanced_security_options" yaml:"advancedSecurityOptions,omitempty"`

		AutoTuneOptions *struct {
			DesiredState *string `tfsdk:"desired_state" yaml:"desiredState,omitempty"`

			MaintenanceSchedules *[]struct {
				CronExpressionForRecurrence *string `tfsdk:"cron_expression_for_recurrence" yaml:"cronExpressionForRecurrence,omitempty"`

				Duration *struct {
					Unit *string `tfsdk:"unit" yaml:"unit,omitempty"`

					Value *int64 `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"duration" yaml:"duration,omitempty"`

				StartAt *string `tfsdk:"start_at" yaml:"startAt,omitempty"`
			} `tfsdk:"maintenance_schedules" yaml:"maintenanceSchedules,omitempty"`
		} `tfsdk:"auto_tune_options" yaml:"autoTuneOptions,omitempty"`

		ClusterConfig *struct {
			ColdStorageOptions *struct {
				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
			} `tfsdk:"cold_storage_options" yaml:"coldStorageOptions,omitempty"`

			DedicatedMasterCount *int64 `tfsdk:"dedicated_master_count" yaml:"dedicatedMasterCount,omitempty"`

			DedicatedMasterEnabled *bool `tfsdk:"dedicated_master_enabled" yaml:"dedicatedMasterEnabled,omitempty"`

			DedicatedMasterType *string `tfsdk:"dedicated_master_type" yaml:"dedicatedMasterType,omitempty"`

			InstanceCount *int64 `tfsdk:"instance_count" yaml:"instanceCount,omitempty"`

			InstanceType *string `tfsdk:"instance_type" yaml:"instanceType,omitempty"`

			WarmCount *int64 `tfsdk:"warm_count" yaml:"warmCount,omitempty"`

			WarmEnabled *bool `tfsdk:"warm_enabled" yaml:"warmEnabled,omitempty"`

			WarmType *string `tfsdk:"warm_type" yaml:"warmType,omitempty"`

			ZoneAwarenessConfig *struct {
				AvailabilityZoneCount *int64 `tfsdk:"availability_zone_count" yaml:"availabilityZoneCount,omitempty"`
			} `tfsdk:"zone_awareness_config" yaml:"zoneAwarenessConfig,omitempty"`

			ZoneAwarenessEnabled *bool `tfsdk:"zone_awareness_enabled" yaml:"zoneAwarenessEnabled,omitempty"`
		} `tfsdk:"cluster_config" yaml:"clusterConfig,omitempty"`

		CognitoOptions *struct {
			Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

			IdentityPoolID *string `tfsdk:"identity_pool_id" yaml:"identityPoolID,omitempty"`

			RoleARN *string `tfsdk:"role_arn" yaml:"roleARN,omitempty"`

			UserPoolID *string `tfsdk:"user_pool_id" yaml:"userPoolID,omitempty"`
		} `tfsdk:"cognito_options" yaml:"cognitoOptions,omitempty"`

		DomainEndpointOptions *struct {
			CustomEndpoint *string `tfsdk:"custom_endpoint" yaml:"customEndpoint,omitempty"`

			CustomEndpointCertificateARN *string `tfsdk:"custom_endpoint_certificate_arn" yaml:"customEndpointCertificateARN,omitempty"`

			CustomEndpointEnabled *bool `tfsdk:"custom_endpoint_enabled" yaml:"customEndpointEnabled,omitempty"`

			EnforceHTTPS *bool `tfsdk:"enforce_https" yaml:"enforceHTTPS,omitempty"`

			TlsSecurityPolicy *string `tfsdk:"tls_security_policy" yaml:"tlsSecurityPolicy,omitempty"`
		} `tfsdk:"domain_endpoint_options" yaml:"domainEndpointOptions,omitempty"`

		EbsOptions *struct {
			EbsEnabled *bool `tfsdk:"ebs_enabled" yaml:"ebsEnabled,omitempty"`

			Iops *int64 `tfsdk:"iops" yaml:"iops,omitempty"`

			Throughput *int64 `tfsdk:"throughput" yaml:"throughput,omitempty"`

			VolumeSize *int64 `tfsdk:"volume_size" yaml:"volumeSize,omitempty"`

			VolumeType *string `tfsdk:"volume_type" yaml:"volumeType,omitempty"`
		} `tfsdk:"ebs_options" yaml:"ebsOptions,omitempty"`

		EncryptionAtRestOptions *struct {
			Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

			KmsKeyID *string `tfsdk:"kms_key_id" yaml:"kmsKeyID,omitempty"`
		} `tfsdk:"encryption_at_rest_options" yaml:"encryptionAtRestOptions,omitempty"`

		EngineVersion *string `tfsdk:"engine_version" yaml:"engineVersion,omitempty"`

		LogPublishingOptions *struct {
			CloudWatchLogsLogGroupARN *string `tfsdk:"cloud_watch_logs_log_group_arn" yaml:"cloudWatchLogsLogGroupARN,omitempty"`

			Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
		} `tfsdk:"log_publishing_options" yaml:"logPublishingOptions,omitempty"`

		Name *string `tfsdk:"name" yaml:"name,omitempty"`

		NodeToNodeEncryptionOptions *struct {
			Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
		} `tfsdk:"node_to_node_encryption_options" yaml:"nodeToNodeEncryptionOptions,omitempty"`

		Tags *[]struct {
			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"tags" yaml:"tags,omitempty"`

		VpcOptions *struct {
			SecurityGroupIDs *[]string `tfsdk:"security_group_i_ds" yaml:"securityGroupIDs,omitempty"`

			SubnetIDs *[]string `tfsdk:"subnet_i_ds" yaml:"subnetIDs,omitempty"`
		} `tfsdk:"vpc_options" yaml:"vpcOptions,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewOpensearchserviceServicesK8SAwsDomainV1Alpha1Resource() resource.Resource {
	return &OpensearchserviceServicesK8SAwsDomainV1Alpha1Resource{}
}

func (r *OpensearchserviceServicesK8SAwsDomainV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_opensearchservice_services_k8s_aws_domain_v1alpha1"
}

func (r *OpensearchserviceServicesK8SAwsDomainV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Domain is the Schema for the Domains API",
		MarkdownDescription: "Domain is the Schema for the Domains API",
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
				Description:         "DomainSpec defines the desired state of Domain.",
				MarkdownDescription: "DomainSpec defines the desired state of Domain.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"access_policies": {
						Description:         "IAM access policy as a JSON-formatted string.",
						MarkdownDescription: "IAM access policy as a JSON-formatted string.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"advanced_options": {
						Description:         "Option to allow references to indices in an HTTP request body. Must be false when configuring access to individual sub-resources. By default, the value is true. See Advanced cluster parameters (http://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html#createdomain-configure-advanced-options) for more information.",
						MarkdownDescription: "Option to allow references to indices in an HTTP request body. Must be false when configuring access to individual sub-resources. By default, the value is true. See Advanced cluster parameters (http://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html#createdomain-configure-advanced-options) for more information.",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"advanced_security_options": {
						Description:         "Specifies advanced security options.",
						MarkdownDescription: "Specifies advanced security options.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"anonymous_auth_enabled": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enabled": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"internal_user_database_enabled": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"master_user_options": {
								Description:         "Credentials for the master user: username and password, ARN, or both.",
								MarkdownDescription: "Credentials for the master user: username and password, ARN, or both.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"master_user_arn": {
										Description:         "The Amazon Resource Name (ARN) of the domain. See Identifiers for IAM Entities (http://docs.aws.amazon.com/IAM/latest/UserGuide/index.html) in Using AWS Identity and Access Management for more information.",
										MarkdownDescription: "The Amazon Resource Name (ARN) of the domain. See Identifiers for IAM Entities (http://docs.aws.amazon.com/IAM/latest/UserGuide/index.html) in Using AWS Identity and Access Management for more information.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"master_user_name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"master_user_password": {
										Description:         "SecretKeyReference combines a k8s corev1.SecretReference with a specific key within the referred-to Secret",
										MarkdownDescription: "SecretKeyReference combines a k8s corev1.SecretReference with a specific key within the referred-to Secret",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "Key is the key within the secret",
												MarkdownDescription: "Key is the key within the secret",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "Name is unique within a namespace to reference a secret resource.",
												MarkdownDescription: "Name is unique within a namespace to reference a secret resource.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespace": {
												Description:         "Namespace defines the space within which the secret name must be unique.",
												MarkdownDescription: "Namespace defines the space within which the secret name must be unique.",

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

							"s_aml_options": {
								Description:         "The SAML application configuration for the domain.",
								MarkdownDescription: "The SAML application configuration for the domain.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"enabled": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"idp": {
										Description:         "The SAML identity povider's information.",
										MarkdownDescription: "The SAML identity povider's information.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"entity_id": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"metadata_content": {
												Description:         "",
												MarkdownDescription: "",

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

									"master_backend_role": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"master_user_name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"roles_key": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"session_timeout_minutes": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"subject_key": {
										Description:         "",
										MarkdownDescription: "",

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

					"auto_tune_options": {
						Description:         "Specifies Auto-Tune options.",
						MarkdownDescription: "Specifies Auto-Tune options.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"desired_state": {
								Description:         "The Auto-Tune desired state. Valid values are ENABLED and DISABLED.",
								MarkdownDescription: "The Auto-Tune desired state. Valid values are ENABLED and DISABLED.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"maintenance_schedules": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"cron_expression_for_recurrence": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"duration": {
										Description:         "The maintenance schedule duration: duration value and duration unit. See Auto-Tune for Amazon OpenSearch Service (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/auto-tune.html) for more information.",
										MarkdownDescription: "The maintenance schedule duration: duration value and duration unit. See Auto-Tune for Amazon OpenSearch Service (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/auto-tune.html) for more information.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"unit": {
												Description:         "The unit of a maintenance schedule duration. Valid value is HOUR. See Auto-Tune for Amazon OpenSearch Service (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/auto-tune.html) for more information.",
												MarkdownDescription: "The unit of a maintenance schedule duration. Valid value is HOUR. See Auto-Tune for Amazon OpenSearch Service (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/auto-tune.html) for more information.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value": {
												Description:         "Integer to specify the value of a maintenance schedule duration. See Auto-Tune for Amazon OpenSearch Service (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/auto-tune.html) for more information.",
												MarkdownDescription: "Integer to specify the value of a maintenance schedule duration. See Auto-Tune for Amazon OpenSearch Service (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/auto-tune.html) for more information.",

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

									"start_at": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											validators.DateTime64Validator(),
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

					"cluster_config": {
						Description:         "Configuration options for a domain. Specifies the instance type and number of instances in the domain.",
						MarkdownDescription: "Configuration options for a domain. Specifies the instance type and number of instances in the domain.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"cold_storage_options": {
								Description:         "Specifies the configuration for cold storage options such as enabled",
								MarkdownDescription: "Specifies the configuration for cold storage options such as enabled",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"enabled": {
										Description:         "",
										MarkdownDescription: "",

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

							"dedicated_master_count": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"dedicated_master_enabled": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"dedicated_master_type": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"instance_count": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"instance_type": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"warm_count": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"warm_enabled": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"warm_type": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"zone_awareness_config": {
								Description:         "The zone awareness configuration for the domain cluster, such as the number of availability zones.",
								MarkdownDescription: "The zone awareness configuration for the domain cluster, such as the number of availability zones.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"availability_zone_count": {
										Description:         "",
										MarkdownDescription: "",

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

							"zone_awareness_enabled": {
								Description:         "",
								MarkdownDescription: "",

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

					"cognito_options": {
						Description:         "Options to specify the Cognito user and identity pools for OpenSearch Dashboards authentication. For more information, see Configuring Amazon Cognito authentication for OpenSearch Dashboards (http://docs.aws.amazon.com/opensearch-service/latest/developerguide/cognito-auth.html).",
						MarkdownDescription: "Options to specify the Cognito user and identity pools for OpenSearch Dashboards authentication. For more information, see Configuring Amazon Cognito authentication for OpenSearch Dashboards (http://docs.aws.amazon.com/opensearch-service/latest/developerguide/cognito-auth.html).",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"enabled": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"identity_pool_id": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"role_arn": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"user_pool_id": {
								Description:         "",
								MarkdownDescription: "",

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

					"domain_endpoint_options": {
						Description:         "Options to specify configurations that will be applied to the domain endpoint.",
						MarkdownDescription: "Options to specify configurations that will be applied to the domain endpoint.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"custom_endpoint": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"custom_endpoint_certificate_arn": {
								Description:         "The Amazon Resource Name (ARN) of the domain. See Identifiers for IAM Entities (http://docs.aws.amazon.com/IAM/latest/UserGuide/index.html) in Using AWS Identity and Access Management for more information.",
								MarkdownDescription: "The Amazon Resource Name (ARN) of the domain. See Identifiers for IAM Entities (http://docs.aws.amazon.com/IAM/latest/UserGuide/index.html) in Using AWS Identity and Access Management for more information.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"custom_endpoint_enabled": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enforce_https": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tls_security_policy": {
								Description:         "",
								MarkdownDescription: "",

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

					"ebs_options": {
						Description:         "Options to enable, disable, and specify the type and size of EBS storage volumes.",
						MarkdownDescription: "Options to enable, disable, and specify the type and size of EBS storage volumes.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"ebs_enabled": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"iops": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"throughput": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"volume_size": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"volume_type": {
								Description:         "The type of EBS volume, standard, gp2, gp3 or io1. See Configuring EBS-based Storage (http://docs.aws.amazon.com/opensearch-service/latest/developerguide/opensearch-createupdatedomains.html#opensearch-createdomain-configure-ebs) for more information.",
								MarkdownDescription: "The type of EBS volume, standard, gp2, gp3 or io1. See Configuring EBS-based Storage (http://docs.aws.amazon.com/opensearch-service/latest/developerguide/opensearch-createupdatedomains.html#opensearch-createdomain-configure-ebs) for more information.",

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

					"encryption_at_rest_options": {
						Description:         "Options for encryption of data at rest.",
						MarkdownDescription: "Options for encryption of data at rest.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"enabled": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"kms_key_id": {
								Description:         "",
								MarkdownDescription: "",

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

					"engine_version": {
						Description:         "String of format Elasticsearch_X.Y or OpenSearch_X.Y to specify the engine version for the Amazon OpenSearch Service domain. For example, 'OpenSearch_1.0' or 'Elasticsearch_7.9'. For more information, see Creating and managing Amazon OpenSearch Service domains (http://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html#createdomains).",
						MarkdownDescription: "String of format Elasticsearch_X.Y or OpenSearch_X.Y to specify the engine version for the Amazon OpenSearch Service domain. For example, 'OpenSearch_1.0' or 'Elasticsearch_7.9'. For more information, see Creating and managing Amazon OpenSearch Service domains (http://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html#createdomains).",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"log_publishing_options": {
						Description:         "Map of LogType and LogPublishingOption, each containing options to publish a given type of OpenSearch log.",
						MarkdownDescription: "Map of LogType and LogPublishingOption, each containing options to publish a given type of OpenSearch log.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"cloud_watch_logs_log_group_arn": {
								Description:         "ARN of the Cloudwatch log group to publish logs to.",
								MarkdownDescription: "ARN of the Cloudwatch log group to publish logs to.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enabled": {
								Description:         "",
								MarkdownDescription: "",

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

					"name": {
						Description:         "The name of the Amazon OpenSearch Service domain you're creating. Domain names are unique across the domains owned by an account within an AWS region. Domain names must start with a lowercase letter and can contain the following characters: a-z (lowercase), 0-9, and - (hyphen).",
						MarkdownDescription: "The name of the Amazon OpenSearch Service domain you're creating. Domain names are unique across the domains owned by an account within an AWS region. Domain names must start with a lowercase letter and can contain the following characters: a-z (lowercase), 0-9, and - (hyphen).",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"node_to_node_encryption_options": {
						Description:         "Node-to-node encryption options.",
						MarkdownDescription: "Node-to-node encryption options.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"enabled": {
								Description:         "",
								MarkdownDescription: "",

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

					"tags": {
						Description:         "A list of Tag added during domain creation.",
						MarkdownDescription: "A list of Tag added during domain creation.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"key": {
								Description:         "A string of length from 1 to 128 characters that specifies the key for a tag. Tag keys must be unique for the domain to which they're attached.",
								MarkdownDescription: "A string of length from 1 to 128 characters that specifies the key for a tag. Tag keys must be unique for the domain to which they're attached.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"value": {
								Description:         "A string of length from 0 to 256 characters that specifies the value for a tag. Tag values can be null and don't have to be unique in a tag set.",
								MarkdownDescription: "A string of length from 0 to 256 characters that specifies the value for a tag. Tag values can be null and don't have to be unique in a tag set.",

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

					"vpc_options": {
						Description:         "Options to specify the subnets and security groups for a VPC endpoint. For more information, see Launching your Amazon OpenSearch Service domains using a VPC (http://docs.aws.amazon.com/opensearch-service/latest/developerguide/vpc.html).",
						MarkdownDescription: "Options to specify the subnets and security groups for a VPC endpoint. For more information, see Launching your Amazon OpenSearch Service domains using a VPC (http://docs.aws.amazon.com/opensearch-service/latest/developerguide/vpc.html).",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"security_group_i_ds": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"subnet_i_ds": {
								Description:         "",
								MarkdownDescription: "",

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
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *OpensearchserviceServicesK8SAwsDomainV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_opensearchservice_services_k8s_aws_domain_v1alpha1")

	var state OpensearchserviceServicesK8SAwsDomainV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel OpensearchserviceServicesK8SAwsDomainV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("opensearchservice.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("Domain")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *OpensearchserviceServicesK8SAwsDomainV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_opensearchservice_services_k8s_aws_domain_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *OpensearchserviceServicesK8SAwsDomainV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_opensearchservice_services_k8s_aws_domain_v1alpha1")

	var state OpensearchserviceServicesK8SAwsDomainV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel OpensearchserviceServicesK8SAwsDomainV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("opensearchservice.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("Domain")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *OpensearchserviceServicesK8SAwsDomainV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_opensearchservice_services_k8s_aws_domain_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
