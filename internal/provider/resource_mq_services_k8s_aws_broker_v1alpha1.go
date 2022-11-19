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

type MqServicesK8SAwsBrokerV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*MqServicesK8SAwsBrokerV1Alpha1Resource)(nil)
)

type MqServicesK8SAwsBrokerV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type MqServicesK8SAwsBrokerV1Alpha1GoModel struct {
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
		AuthenticationStrategy *string `tfsdk:"authentication_strategy" yaml:"authenticationStrategy,omitempty"`

		AutoMinorVersionUpgrade *bool `tfsdk:"auto_minor_version_upgrade" yaml:"autoMinorVersionUpgrade,omitempty"`

		Configuration *struct {
			Id *string `tfsdk:"id" yaml:"id,omitempty"`

			Revision *int64 `tfsdk:"revision" yaml:"revision,omitempty"`
		} `tfsdk:"configuration" yaml:"configuration,omitempty"`

		CreatorRequestID *string `tfsdk:"creator_request_id" yaml:"creatorRequestID,omitempty"`

		DeploymentMode *string `tfsdk:"deployment_mode" yaml:"deploymentMode,omitempty"`

		EncryptionOptions *struct {
			KmsKeyID *string `tfsdk:"kms_key_id" yaml:"kmsKeyID,omitempty"`

			UseAWSOwnedKey *bool `tfsdk:"use_aws_owned_key" yaml:"useAWSOwnedKey,omitempty"`
		} `tfsdk:"encryption_options" yaml:"encryptionOptions,omitempty"`

		EngineType *string `tfsdk:"engine_type" yaml:"engineType,omitempty"`

		EngineVersion *string `tfsdk:"engine_version" yaml:"engineVersion,omitempty"`

		HostInstanceType *string `tfsdk:"host_instance_type" yaml:"hostInstanceType,omitempty"`

		LdapServerMetadata *struct {
			Hosts *[]string `tfsdk:"hosts" yaml:"hosts,omitempty"`

			RoleBase *string `tfsdk:"role_base" yaml:"roleBase,omitempty"`

			RoleName *string `tfsdk:"role_name" yaml:"roleName,omitempty"`

			RoleSearchMatching *string `tfsdk:"role_search_matching" yaml:"roleSearchMatching,omitempty"`

			RoleSearchSubtree *bool `tfsdk:"role_search_subtree" yaml:"roleSearchSubtree,omitempty"`

			ServiceAccountPassword *string `tfsdk:"service_account_password" yaml:"serviceAccountPassword,omitempty"`

			ServiceAccountUsername *string `tfsdk:"service_account_username" yaml:"serviceAccountUsername,omitempty"`

			UserBase *string `tfsdk:"user_base" yaml:"userBase,omitempty"`

			UserRoleName *string `tfsdk:"user_role_name" yaml:"userRoleName,omitempty"`

			UserSearchMatching *string `tfsdk:"user_search_matching" yaml:"userSearchMatching,omitempty"`

			UserSearchSubtree *bool `tfsdk:"user_search_subtree" yaml:"userSearchSubtree,omitempty"`
		} `tfsdk:"ldap_server_metadata" yaml:"ldapServerMetadata,omitempty"`

		Logs *struct {
			Audit *bool `tfsdk:"audit" yaml:"audit,omitempty"`

			General *bool `tfsdk:"general" yaml:"general,omitempty"`
		} `tfsdk:"logs" yaml:"logs,omitempty"`

		MaintenanceWindowStartTime *struct {
			DayOfWeek *string `tfsdk:"day_of_week" yaml:"dayOfWeek,omitempty"`

			TimeOfDay *string `tfsdk:"time_of_day" yaml:"timeOfDay,omitempty"`

			TimeZone *string `tfsdk:"time_zone" yaml:"timeZone,omitempty"`
		} `tfsdk:"maintenance_window_start_time" yaml:"maintenanceWindowStartTime,omitempty"`

		Name *string `tfsdk:"name" yaml:"name,omitempty"`

		PubliclyAccessible *bool `tfsdk:"publicly_accessible" yaml:"publiclyAccessible,omitempty"`

		SecurityGroups *[]string `tfsdk:"security_groups" yaml:"securityGroups,omitempty"`

		StorageType *string `tfsdk:"storage_type" yaml:"storageType,omitempty"`

		SubnetIDs *[]string `tfsdk:"subnet_i_ds" yaml:"subnetIDs,omitempty"`

		Tags *map[string]string `tfsdk:"tags" yaml:"tags,omitempty"`

		Users *[]struct {
			ConsoleAccess *bool `tfsdk:"console_access" yaml:"consoleAccess,omitempty"`

			Groups *[]string `tfsdk:"groups" yaml:"groups,omitempty"`

			Password *struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
			} `tfsdk:"password" yaml:"password,omitempty"`

			Username *string `tfsdk:"username" yaml:"username,omitempty"`
		} `tfsdk:"users" yaml:"users,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewMqServicesK8SAwsBrokerV1Alpha1Resource() resource.Resource {
	return &MqServicesK8SAwsBrokerV1Alpha1Resource{}
}

func (r *MqServicesK8SAwsBrokerV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mq_services_k8s_aws_broker_v1alpha1"
}

func (r *MqServicesK8SAwsBrokerV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Broker is the Schema for the Brokers API",
		MarkdownDescription: "Broker is the Schema for the Brokers API",
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
				Description:         "BrokerSpec defines the desired state of Broker.",
				MarkdownDescription: "BrokerSpec defines the desired state of Broker.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"authentication_strategy": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"auto_minor_version_upgrade": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"configuration": {
						Description:         "A list of information about the configuration.  Does not apply to RabbitMQ brokers.",
						MarkdownDescription: "A list of information about the configuration.  Does not apply to RabbitMQ brokers.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"id": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"revision": {
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

					"creator_request_id": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"deployment_mode": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"encryption_options": {
						Description:         "Does not apply to RabbitMQ brokers.  Encryption options for the broker.",
						MarkdownDescription: "Does not apply to RabbitMQ brokers.  Encryption options for the broker.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"kms_key_id": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"use_aws_owned_key": {
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

					"engine_type": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"engine_version": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"host_instance_type": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"ldap_server_metadata": {
						Description:         "Optional. The metadata of the LDAP server used to authenticate and authorize connections to the broker.  Does not apply to RabbitMQ brokers.",
						MarkdownDescription: "Optional. The metadata of the LDAP server used to authenticate and authorize connections to the broker.  Does not apply to RabbitMQ brokers.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"hosts": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"role_base": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"role_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"role_search_matching": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"role_search_subtree": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"service_account_password": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"service_account_username": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"user_base": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"user_role_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"user_search_matching": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"user_search_subtree": {
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

					"logs": {
						Description:         "The list of information about logs to be enabled for the specified broker.",
						MarkdownDescription: "The list of information about logs to be enabled for the specified broker.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"audit": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"general": {
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

					"maintenance_window_start_time": {
						Description:         "The scheduled time period relative to UTC during which Amazon MQ begins to apply pending updates or patches to the broker.",
						MarkdownDescription: "The scheduled time period relative to UTC during which Amazon MQ begins to apply pending updates or patches to the broker.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"day_of_week": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"time_of_day": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"time_zone": {
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

					"name": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"publicly_accessible": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"security_groups": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"storage_type": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

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

					"tags": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"users": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"console_access": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"groups": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"password": {
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

							"username": {
								Description:         "",
								MarkdownDescription: "",

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
		},
	}, nil
}

func (r *MqServicesK8SAwsBrokerV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_mq_services_k8s_aws_broker_v1alpha1")

	var state MqServicesK8SAwsBrokerV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel MqServicesK8SAwsBrokerV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("mq.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("Broker")

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

func (r *MqServicesK8SAwsBrokerV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_mq_services_k8s_aws_broker_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *MqServicesK8SAwsBrokerV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_mq_services_k8s_aws_broker_v1alpha1")

	var state MqServicesK8SAwsBrokerV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel MqServicesK8SAwsBrokerV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("mq.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("Broker")

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

func (r *MqServicesK8SAwsBrokerV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_mq_services_k8s_aws_broker_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
