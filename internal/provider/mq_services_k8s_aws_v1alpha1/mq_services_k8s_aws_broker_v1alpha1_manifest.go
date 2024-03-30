/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package mq_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource = &MqServicesK8SAwsBrokerV1Alpha1Manifest{}
)

func NewMqServicesK8SAwsBrokerV1Alpha1Manifest() datasource.DataSource {
	return &MqServicesK8SAwsBrokerV1Alpha1Manifest{}
}

type MqServicesK8SAwsBrokerV1Alpha1Manifest struct{}

type MqServicesK8SAwsBrokerV1Alpha1ManifestData struct {
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
		AuthenticationStrategy  *string `tfsdk:"authentication_strategy" json:"authenticationStrategy,omitempty"`
		AutoMinorVersionUpgrade *bool   `tfsdk:"auto_minor_version_upgrade" json:"autoMinorVersionUpgrade,omitempty"`
		Configuration           *struct {
			Id       *string `tfsdk:"id" json:"id,omitempty"`
			Revision *int64  `tfsdk:"revision" json:"revision,omitempty"`
		} `tfsdk:"configuration" json:"configuration,omitempty"`
		CreatorRequestID  *string `tfsdk:"creator_request_id" json:"creatorRequestID,omitempty"`
		DeploymentMode    *string `tfsdk:"deployment_mode" json:"deploymentMode,omitempty"`
		EncryptionOptions *struct {
			KmsKeyID       *string `tfsdk:"kms_key_id" json:"kmsKeyID,omitempty"`
			UseAWSOwnedKey *bool   `tfsdk:"use_aws_owned_key" json:"useAWSOwnedKey,omitempty"`
		} `tfsdk:"encryption_options" json:"encryptionOptions,omitempty"`
		EngineType         *string `tfsdk:"engine_type" json:"engineType,omitempty"`
		EngineVersion      *string `tfsdk:"engine_version" json:"engineVersion,omitempty"`
		HostInstanceType   *string `tfsdk:"host_instance_type" json:"hostInstanceType,omitempty"`
		LdapServerMetadata *struct {
			Hosts                  *[]string `tfsdk:"hosts" json:"hosts,omitempty"`
			RoleBase               *string   `tfsdk:"role_base" json:"roleBase,omitempty"`
			RoleName               *string   `tfsdk:"role_name" json:"roleName,omitempty"`
			RoleSearchMatching     *string   `tfsdk:"role_search_matching" json:"roleSearchMatching,omitempty"`
			RoleSearchSubtree      *bool     `tfsdk:"role_search_subtree" json:"roleSearchSubtree,omitempty"`
			ServiceAccountPassword *string   `tfsdk:"service_account_password" json:"serviceAccountPassword,omitempty"`
			ServiceAccountUsername *string   `tfsdk:"service_account_username" json:"serviceAccountUsername,omitempty"`
			UserBase               *string   `tfsdk:"user_base" json:"userBase,omitempty"`
			UserRoleName           *string   `tfsdk:"user_role_name" json:"userRoleName,omitempty"`
			UserSearchMatching     *string   `tfsdk:"user_search_matching" json:"userSearchMatching,omitempty"`
			UserSearchSubtree      *bool     `tfsdk:"user_search_subtree" json:"userSearchSubtree,omitempty"`
		} `tfsdk:"ldap_server_metadata" json:"ldapServerMetadata,omitempty"`
		Logs *struct {
			Audit   *bool `tfsdk:"audit" json:"audit,omitempty"`
			General *bool `tfsdk:"general" json:"general,omitempty"`
		} `tfsdk:"logs" json:"logs,omitempty"`
		MaintenanceWindowStartTime *struct {
			DayOfWeek *string `tfsdk:"day_of_week" json:"dayOfWeek,omitempty"`
			TimeOfDay *string `tfsdk:"time_of_day" json:"timeOfDay,omitempty"`
			TimeZone  *string `tfsdk:"time_zone" json:"timeZone,omitempty"`
		} `tfsdk:"maintenance_window_start_time" json:"maintenanceWindowStartTime,omitempty"`
		Name               *string `tfsdk:"name" json:"name,omitempty"`
		PubliclyAccessible *bool   `tfsdk:"publicly_accessible" json:"publiclyAccessible,omitempty"`
		SecurityGroupRefs  *[]struct {
			From *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"security_group_refs" json:"securityGroupRefs,omitempty"`
		SecurityGroups *[]string `tfsdk:"security_groups" json:"securityGroups,omitempty"`
		StorageType    *string   `tfsdk:"storage_type" json:"storageType,omitempty"`
		SubnetIDs      *[]string `tfsdk:"subnet_i_ds" json:"subnetIDs,omitempty"`
		SubnetRefs     *[]struct {
			From *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"subnet_refs" json:"subnetRefs,omitempty"`
		Tags  *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
		Users *[]struct {
			ConsoleAccess *bool     `tfsdk:"console_access" json:"consoleAccess,omitempty"`
			Groups        *[]string `tfsdk:"groups" json:"groups,omitempty"`
			Password      *struct {
				Key       *string `tfsdk:"key" json:"key,omitempty"`
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"password" json:"password,omitempty"`
			Username *string `tfsdk:"username" json:"username,omitempty"`
		} `tfsdk:"users" json:"users,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *MqServicesK8SAwsBrokerV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_mq_services_k8s_aws_broker_v1alpha1_manifest"
}

func (r *MqServicesK8SAwsBrokerV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Broker is the Schema for the Brokers API",
		MarkdownDescription: "Broker is the Schema for the Brokers API",
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
				Description:         "BrokerSpec defines the desired state of Broker.",
				MarkdownDescription: "BrokerSpec defines the desired state of Broker.",
				Attributes: map[string]schema.Attribute{
					"authentication_strategy": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"auto_minor_version_upgrade": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"configuration": schema.SingleNestedAttribute{
						Description:         "A list of information about the configuration.Does not apply to RabbitMQ brokers.",
						MarkdownDescription: "A list of information about the configuration.Does not apply to RabbitMQ brokers.",
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"revision": schema.Int64Attribute{
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

					"creator_request_id": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"deployment_mode": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"encryption_options": schema.SingleNestedAttribute{
						Description:         "Does not apply to RabbitMQ brokers.Encryption options for the broker.",
						MarkdownDescription: "Does not apply to RabbitMQ brokers.Encryption options for the broker.",
						Attributes: map[string]schema.Attribute{
							"kms_key_id": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"use_aws_owned_key": schema.BoolAttribute{
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

					"engine_type": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"engine_version": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"host_instance_type": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"ldap_server_metadata": schema.SingleNestedAttribute{
						Description:         "Optional. The metadata of the LDAP server used to authenticate and authorizeconnections to the broker.Does not apply to RabbitMQ brokers.",
						MarkdownDescription: "Optional. The metadata of the LDAP server used to authenticate and authorizeconnections to the broker.Does not apply to RabbitMQ brokers.",
						Attributes: map[string]schema.Attribute{
							"hosts": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"role_base": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"role_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"role_search_matching": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"role_search_subtree": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_account_password": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_account_username": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"user_base": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"user_role_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"user_search_matching": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"user_search_subtree": schema.BoolAttribute{
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

					"logs": schema.SingleNestedAttribute{
						Description:         "The list of information about logs to be enabled for the specified broker.",
						MarkdownDescription: "The list of information about logs to be enabled for the specified broker.",
						Attributes: map[string]schema.Attribute{
							"audit": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"general": schema.BoolAttribute{
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

					"maintenance_window_start_time": schema.SingleNestedAttribute{
						Description:         "The scheduled time period relative to UTC during which Amazon MQ begins toapply pending updates or patches to the broker.",
						MarkdownDescription: "The scheduled time period relative to UTC during which Amazon MQ begins toapply pending updates or patches to the broker.",
						Attributes: map[string]schema.Attribute{
							"day_of_week": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"time_of_day": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"time_zone": schema.StringAttribute{
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
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"publicly_accessible": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"security_group_refs": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"from": schema.SingleNestedAttribute{
									Description:         "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
									MarkdownDescription: "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"security_groups": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"storage_type": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
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

					"subnet_refs": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"from": schema.SingleNestedAttribute{
									Description:         "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
									MarkdownDescription: "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"tags": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"users": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"console_access": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"groups": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"password": schema.SingleNestedAttribute{
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

								"username": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
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
	}
}

func (r *MqServicesK8SAwsBrokerV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_mq_services_k8s_aws_broker_v1alpha1_manifest")

	var model MqServicesK8SAwsBrokerV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("mq.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("Broker")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
