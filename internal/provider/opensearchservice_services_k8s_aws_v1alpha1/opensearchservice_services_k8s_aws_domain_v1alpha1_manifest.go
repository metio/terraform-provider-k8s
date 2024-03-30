/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package opensearchservice_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource = &OpensearchserviceServicesK8SAwsDomainV1Alpha1Manifest{}
)

func NewOpensearchserviceServicesK8SAwsDomainV1Alpha1Manifest() datasource.DataSource {
	return &OpensearchserviceServicesK8SAwsDomainV1Alpha1Manifest{}
}

type OpensearchserviceServicesK8SAwsDomainV1Alpha1Manifest struct{}

type OpensearchserviceServicesK8SAwsDomainV1Alpha1ManifestData struct {
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
		AccessPolicies          *string            `tfsdk:"access_policies" json:"accessPolicies,omitempty"`
		AdvancedOptions         *map[string]string `tfsdk:"advanced_options" json:"advancedOptions,omitempty"`
		AdvancedSecurityOptions *struct {
			AnonymousAuthEnabled        *bool `tfsdk:"anonymous_auth_enabled" json:"anonymousAuthEnabled,omitempty"`
			Enabled                     *bool `tfsdk:"enabled" json:"enabled,omitempty"`
			InternalUserDatabaseEnabled *bool `tfsdk:"internal_user_database_enabled" json:"internalUserDatabaseEnabled,omitempty"`
			MasterUserOptions           *struct {
				MasterUserARN      *string `tfsdk:"master_user_arn" json:"masterUserARN,omitempty"`
				MasterUserName     *string `tfsdk:"master_user_name" json:"masterUserName,omitempty"`
				MasterUserPassword *struct {
					Key       *string `tfsdk:"key" json:"key,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"master_user_password" json:"masterUserPassword,omitempty"`
			} `tfsdk:"master_user_options" json:"masterUserOptions,omitempty"`
			SAMLOptions *struct {
				Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
				Idp     *struct {
					EntityID        *string `tfsdk:"entity_id" json:"entityID,omitempty"`
					MetadataContent *string `tfsdk:"metadata_content" json:"metadataContent,omitempty"`
				} `tfsdk:"idp" json:"idp,omitempty"`
				MasterBackendRole     *string `tfsdk:"master_backend_role" json:"masterBackendRole,omitempty"`
				MasterUserName        *string `tfsdk:"master_user_name" json:"masterUserName,omitempty"`
				RolesKey              *string `tfsdk:"roles_key" json:"rolesKey,omitempty"`
				SessionTimeoutMinutes *int64  `tfsdk:"session_timeout_minutes" json:"sessionTimeoutMinutes,omitempty"`
				SubjectKey            *string `tfsdk:"subject_key" json:"subjectKey,omitempty"`
			} `tfsdk:"s_aml_options" json:"sAMLOptions,omitempty"`
		} `tfsdk:"advanced_security_options" json:"advancedSecurityOptions,omitempty"`
		AutoTuneOptions *struct {
			DesiredState         *string `tfsdk:"desired_state" json:"desiredState,omitempty"`
			MaintenanceSchedules *[]struct {
				CronExpressionForRecurrence *string `tfsdk:"cron_expression_for_recurrence" json:"cronExpressionForRecurrence,omitempty"`
				Duration                    *struct {
					Unit  *string `tfsdk:"unit" json:"unit,omitempty"`
					Value *int64  `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"duration" json:"duration,omitempty"`
				StartAt *string `tfsdk:"start_at" json:"startAt,omitempty"`
			} `tfsdk:"maintenance_schedules" json:"maintenanceSchedules,omitempty"`
		} `tfsdk:"auto_tune_options" json:"autoTuneOptions,omitempty"`
		ClusterConfig *struct {
			ColdStorageOptions *struct {
				Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
			} `tfsdk:"cold_storage_options" json:"coldStorageOptions,omitempty"`
			DedicatedMasterCount   *int64  `tfsdk:"dedicated_master_count" json:"dedicatedMasterCount,omitempty"`
			DedicatedMasterEnabled *bool   `tfsdk:"dedicated_master_enabled" json:"dedicatedMasterEnabled,omitempty"`
			DedicatedMasterType    *string `tfsdk:"dedicated_master_type" json:"dedicatedMasterType,omitempty"`
			InstanceCount          *int64  `tfsdk:"instance_count" json:"instanceCount,omitempty"`
			InstanceType           *string `tfsdk:"instance_type" json:"instanceType,omitempty"`
			WarmCount              *int64  `tfsdk:"warm_count" json:"warmCount,omitempty"`
			WarmEnabled            *bool   `tfsdk:"warm_enabled" json:"warmEnabled,omitempty"`
			WarmType               *string `tfsdk:"warm_type" json:"warmType,omitempty"`
			ZoneAwarenessConfig    *struct {
				AvailabilityZoneCount *int64 `tfsdk:"availability_zone_count" json:"availabilityZoneCount,omitempty"`
			} `tfsdk:"zone_awareness_config" json:"zoneAwarenessConfig,omitempty"`
			ZoneAwarenessEnabled *bool `tfsdk:"zone_awareness_enabled" json:"zoneAwarenessEnabled,omitempty"`
		} `tfsdk:"cluster_config" json:"clusterConfig,omitempty"`
		CognitoOptions *struct {
			Enabled        *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			IdentityPoolID *string `tfsdk:"identity_pool_id" json:"identityPoolID,omitempty"`
			RoleARN        *string `tfsdk:"role_arn" json:"roleARN,omitempty"`
			UserPoolID     *string `tfsdk:"user_pool_id" json:"userPoolID,omitempty"`
		} `tfsdk:"cognito_options" json:"cognitoOptions,omitempty"`
		DomainEndpointOptions *struct {
			CustomEndpoint               *string `tfsdk:"custom_endpoint" json:"customEndpoint,omitempty"`
			CustomEndpointCertificateARN *string `tfsdk:"custom_endpoint_certificate_arn" json:"customEndpointCertificateARN,omitempty"`
			CustomEndpointEnabled        *bool   `tfsdk:"custom_endpoint_enabled" json:"customEndpointEnabled,omitempty"`
			EnforceHTTPS                 *bool   `tfsdk:"enforce_https" json:"enforceHTTPS,omitempty"`
			TlsSecurityPolicy            *string `tfsdk:"tls_security_policy" json:"tlsSecurityPolicy,omitempty"`
		} `tfsdk:"domain_endpoint_options" json:"domainEndpointOptions,omitempty"`
		EbsOptions *struct {
			EbsEnabled *bool   `tfsdk:"ebs_enabled" json:"ebsEnabled,omitempty"`
			Iops       *int64  `tfsdk:"iops" json:"iops,omitempty"`
			Throughput *int64  `tfsdk:"throughput" json:"throughput,omitempty"`
			VolumeSize *int64  `tfsdk:"volume_size" json:"volumeSize,omitempty"`
			VolumeType *string `tfsdk:"volume_type" json:"volumeType,omitempty"`
		} `tfsdk:"ebs_options" json:"ebsOptions,omitempty"`
		EncryptionAtRestOptions *struct {
			Enabled  *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			KmsKeyID *string `tfsdk:"kms_key_id" json:"kmsKeyID,omitempty"`
		} `tfsdk:"encryption_at_rest_options" json:"encryptionAtRestOptions,omitempty"`
		EngineVersion        *string `tfsdk:"engine_version" json:"engineVersion,omitempty"`
		LogPublishingOptions *struct {
			CloudWatchLogsLogGroupARN *string `tfsdk:"cloud_watch_logs_log_group_arn" json:"cloudWatchLogsLogGroupARN,omitempty"`
			Enabled                   *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
		} `tfsdk:"log_publishing_options" json:"logPublishingOptions,omitempty"`
		Name                        *string `tfsdk:"name" json:"name,omitempty"`
		NodeToNodeEncryptionOptions *struct {
			Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
		} `tfsdk:"node_to_node_encryption_options" json:"nodeToNodeEncryptionOptions,omitempty"`
		Tags *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
		VpcOptions *struct {
			SecurityGroupIDs *[]string `tfsdk:"security_group_i_ds" json:"securityGroupIDs,omitempty"`
			SubnetIDs        *[]string `tfsdk:"subnet_i_ds" json:"subnetIDs,omitempty"`
		} `tfsdk:"vpc_options" json:"vpcOptions,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OpensearchserviceServicesK8SAwsDomainV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_opensearchservice_services_k8s_aws_domain_v1alpha1_manifest"
}

func (r *OpensearchserviceServicesK8SAwsDomainV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Domain is the Schema for the Domains API",
		MarkdownDescription: "Domain is the Schema for the Domains API",
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
				Description:         "DomainSpec defines the desired state of Domain.",
				MarkdownDescription: "DomainSpec defines the desired state of Domain.",
				Attributes: map[string]schema.Attribute{
					"access_policies": schema.StringAttribute{
						Description:         "IAM access policy as a JSON-formatted string.",
						MarkdownDescription: "IAM access policy as a JSON-formatted string.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"advanced_options": schema.MapAttribute{
						Description:         "Option to allow references to indices in an HTTP request body. Must be falsewhen configuring access to individual sub-resources. By default, the valueis true. See Advanced cluster parameters (http://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html#createdomain-configure-advanced-options)for more information.",
						MarkdownDescription: "Option to allow references to indices in an HTTP request body. Must be falsewhen configuring access to individual sub-resources. By default, the valueis true. See Advanced cluster parameters (http://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html#createdomain-configure-advanced-options)for more information.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"advanced_security_options": schema.SingleNestedAttribute{
						Description:         "Specifies advanced security options.",
						MarkdownDescription: "Specifies advanced security options.",
						Attributes: map[string]schema.Attribute{
							"anonymous_auth_enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"internal_user_database_enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"master_user_options": schema.SingleNestedAttribute{
								Description:         "Credentials for the master user: username and password, ARN, or both.",
								MarkdownDescription: "Credentials for the master user: username and password, ARN, or both.",
								Attributes: map[string]schema.Attribute{
									"master_user_arn": schema.StringAttribute{
										Description:         "The Amazon Resource Name (ARN) of the domain. See Identifiers for IAM Entities(http://docs.aws.amazon.com/IAM/latest/UserGuide/index.html) in Using AWSIdentity and Access Management for more information.",
										MarkdownDescription: "The Amazon Resource Name (ARN) of the domain. See Identifiers for IAM Entities(http://docs.aws.amazon.com/IAM/latest/UserGuide/index.html) in Using AWSIdentity and Access Management for more information.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"master_user_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"master_user_password": schema.SingleNestedAttribute{
										Description:         "SecretKeyReference combines a k8s corev1.SecretReference with aspecific key within the referred-to Secret",
										MarkdownDescription: "SecretKeyReference combines a k8s corev1.SecretReference with aspecific key within the referred-to Secret",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "Key is the key within the secret",
												MarkdownDescription: "Key is the key within the secret",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "name is unique within a namespace to reference a secret resource.",
												MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "namespace defines the space within which the secret name must be unique.",
												MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
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

							"s_aml_options": schema.SingleNestedAttribute{
								Description:         "The SAML application configuration for the domain.",
								MarkdownDescription: "The SAML application configuration for the domain.",
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"idp": schema.SingleNestedAttribute{
										Description:         "The SAML identity povider's information.",
										MarkdownDescription: "The SAML identity povider's information.",
										Attributes: map[string]schema.Attribute{
											"entity_id": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"metadata_content": schema.StringAttribute{
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

									"master_backend_role": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"master_user_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"roles_key": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"session_timeout_minutes": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"subject_key": schema.StringAttribute{
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"auto_tune_options": schema.SingleNestedAttribute{
						Description:         "Specifies Auto-Tune options.",
						MarkdownDescription: "Specifies Auto-Tune options.",
						Attributes: map[string]schema.Attribute{
							"desired_state": schema.StringAttribute{
								Description:         "The Auto-Tune desired state. Valid values are ENABLED and DISABLED.",
								MarkdownDescription: "The Auto-Tune desired state. Valid values are ENABLED and DISABLED.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"maintenance_schedules": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"cron_expression_for_recurrence": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"duration": schema.SingleNestedAttribute{
											Description:         "The maintenance schedule duration: duration value and duration unit. SeeAuto-Tune for Amazon OpenSearch Service (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/auto-tune.html)for more information.",
											MarkdownDescription: "The maintenance schedule duration: duration value and duration unit. SeeAuto-Tune for Amazon OpenSearch Service (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/auto-tune.html)for more information.",
											Attributes: map[string]schema.Attribute{
												"unit": schema.StringAttribute{
													Description:         "The unit of a maintenance schedule duration. Valid value is HOUR. See Auto-Tunefor Amazon OpenSearch Service (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/auto-tune.html)for more information.",
													MarkdownDescription: "The unit of a maintenance schedule duration. Valid value is HOUR. See Auto-Tunefor Amazon OpenSearch Service (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/auto-tune.html)for more information.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.Int64Attribute{
													Description:         "Integer to specify the value of a maintenance schedule duration. See Auto-Tunefor Amazon OpenSearch Service (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/auto-tune.html)for more information.",
													MarkdownDescription: "Integer to specify the value of a maintenance schedule duration. See Auto-Tunefor Amazon OpenSearch Service (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/auto-tune.html)for more information.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"start_at": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												validators.DateTime64Validator(),
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

					"cluster_config": schema.SingleNestedAttribute{
						Description:         "Configuration options for a domain. Specifies the instance type and numberof instances in the domain.",
						MarkdownDescription: "Configuration options for a domain. Specifies the instance type and numberof instances in the domain.",
						Attributes: map[string]schema.Attribute{
							"cold_storage_options": schema.SingleNestedAttribute{
								Description:         "Specifies the configuration for cold storage options such as enabled",
								MarkdownDescription: "Specifies the configuration for cold storage options such as enabled",
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
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

							"dedicated_master_count": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dedicated_master_enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dedicated_master_type": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"instance_count": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"instance_type": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"warm_count": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"warm_enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"warm_type": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"zone_awareness_config": schema.SingleNestedAttribute{
								Description:         "The zone awareness configuration for the domain cluster, such as the numberof availability zones.",
								MarkdownDescription: "The zone awareness configuration for the domain cluster, such as the numberof availability zones.",
								Attributes: map[string]schema.Attribute{
									"availability_zone_count": schema.Int64Attribute{
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

							"zone_awareness_enabled": schema.BoolAttribute{
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

					"cognito_options": schema.SingleNestedAttribute{
						Description:         "Options to specify the Cognito user and identity pools for OpenSearch Dashboardsauthentication. For more information, see Configuring Amazon Cognito authenticationfor OpenSearch Dashboards (http://docs.aws.amazon.com/opensearch-service/latest/developerguide/cognito-auth.html).",
						MarkdownDescription: "Options to specify the Cognito user and identity pools for OpenSearch Dashboardsauthentication. For more information, see Configuring Amazon Cognito authenticationfor OpenSearch Dashboards (http://docs.aws.amazon.com/opensearch-service/latest/developerguide/cognito-auth.html).",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"identity_pool_id": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"role_arn": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"user_pool_id": schema.StringAttribute{
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

					"domain_endpoint_options": schema.SingleNestedAttribute{
						Description:         "Options to specify configurations that will be applied to the domain endpoint.",
						MarkdownDescription: "Options to specify configurations that will be applied to the domain endpoint.",
						Attributes: map[string]schema.Attribute{
							"custom_endpoint": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"custom_endpoint_certificate_arn": schema.StringAttribute{
								Description:         "The Amazon Resource Name (ARN) of the domain. See Identifiers for IAM Entities(http://docs.aws.amazon.com/IAM/latest/UserGuide/index.html) in Using AWSIdentity and Access Management for more information.",
								MarkdownDescription: "The Amazon Resource Name (ARN) of the domain. See Identifiers for IAM Entities(http://docs.aws.amazon.com/IAM/latest/UserGuide/index.html) in Using AWSIdentity and Access Management for more information.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"custom_endpoint_enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enforce_https": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls_security_policy": schema.StringAttribute{
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

					"ebs_options": schema.SingleNestedAttribute{
						Description:         "Options to enable, disable, and specify the type and size of EBS storagevolumes.",
						MarkdownDescription: "Options to enable, disable, and specify the type and size of EBS storagevolumes.",
						Attributes: map[string]schema.Attribute{
							"ebs_enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"iops": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"throughput": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"volume_size": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"volume_type": schema.StringAttribute{
								Description:         "The type of EBS volume, standard, gp2, gp3 or io1. See Configuring EBS-basedStorage (http://docs.aws.amazon.com/opensearch-service/latest/developerguide/opensearch-createupdatedomains.html#opensearch-createdomain-configure-ebs)for more information.",
								MarkdownDescription: "The type of EBS volume, standard, gp2, gp3 or io1. See Configuring EBS-basedStorage (http://docs.aws.amazon.com/opensearch-service/latest/developerguide/opensearch-createupdatedomains.html#opensearch-createdomain-configure-ebs)for more information.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"encryption_at_rest_options": schema.SingleNestedAttribute{
						Description:         "Options for encryption of data at rest.",
						MarkdownDescription: "Options for encryption of data at rest.",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kms_key_id": schema.StringAttribute{
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

					"engine_version": schema.StringAttribute{
						Description:         "String of format Elasticsearch_X.Y or OpenSearch_X.Y to specify the engineversion for the Amazon OpenSearch Service domain. For example, 'OpenSearch_1.0'or 'Elasticsearch_7.9'. For more information, see Creating and managing AmazonOpenSearch Service domains (http://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html#createdomains).",
						MarkdownDescription: "String of format Elasticsearch_X.Y or OpenSearch_X.Y to specify the engineversion for the Amazon OpenSearch Service domain. For example, 'OpenSearch_1.0'or 'Elasticsearch_7.9'. For more information, see Creating and managing AmazonOpenSearch Service domains (http://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html#createdomains).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"log_publishing_options": schema.SingleNestedAttribute{
						Description:         "Map of LogType and LogPublishingOption, each containing options to publisha given type of OpenSearch log.",
						MarkdownDescription: "Map of LogType and LogPublishingOption, each containing options to publisha given type of OpenSearch log.",
						Attributes: map[string]schema.Attribute{
							"cloud_watch_logs_log_group_arn": schema.StringAttribute{
								Description:         "ARN of the Cloudwatch log group to publish logs to.",
								MarkdownDescription: "ARN of the Cloudwatch log group to publish logs to.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enabled": schema.BoolAttribute{
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

					"name": schema.StringAttribute{
						Description:         "The name of the Amazon OpenSearch Service domain you're creating. Domainnames are unique across the domains owned by an account within an AWS region.Domain names must start with a lowercase letter and can contain the followingcharacters: a-z (lowercase), 0-9, and - (hyphen).",
						MarkdownDescription: "The name of the Amazon OpenSearch Service domain you're creating. Domainnames are unique across the domains owned by an account within an AWS region.Domain names must start with a lowercase letter and can contain the followingcharacters: a-z (lowercase), 0-9, and - (hyphen).",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"node_to_node_encryption_options": schema.SingleNestedAttribute{
						Description:         "Node-to-node encryption options.",
						MarkdownDescription: "Node-to-node encryption options.",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
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

					"tags": schema.ListNestedAttribute{
						Description:         "A list of Tag added during domain creation.",
						MarkdownDescription: "A list of Tag added during domain creation.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Description:         "A string of length from 1 to 128 characters that specifies the key for atag. Tag keys must be unique for the domain to which they're attached.",
									MarkdownDescription: "A string of length from 1 to 128 characters that specifies the key for atag. Tag keys must be unique for the domain to which they're attached.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "A string of length from 0 to 256 characters that specifies the value fora tag. Tag values can be null and don't have to be unique in a tag set.",
									MarkdownDescription: "A string of length from 0 to 256 characters that specifies the value fora tag. Tag values can be null and don't have to be unique in a tag set.",
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

					"vpc_options": schema.SingleNestedAttribute{
						Description:         "Options to specify the subnets and security groups for a VPC endpoint. Formore information, see Launching your Amazon OpenSearch Service domains usinga VPC (http://docs.aws.amazon.com/opensearch-service/latest/developerguide/vpc.html).",
						MarkdownDescription: "Options to specify the subnets and security groups for a VPC endpoint. Formore information, see Launching your Amazon OpenSearch Service domains usinga VPC (http://docs.aws.amazon.com/opensearch-service/latest/developerguide/vpc.html).",
						Attributes: map[string]schema.Attribute{
							"security_group_i_ds": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"subnet_i_ds": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
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
	}
}

func (r *OpensearchserviceServicesK8SAwsDomainV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_opensearchservice_services_k8s_aws_domain_v1alpha1_manifest")

	var model OpensearchserviceServicesK8SAwsDomainV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("opensearchservice.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("Domain")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
