/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"

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

type KafkaStrimziIoKafkaUserV1Beta2Resource struct{}

var (
	_ resource.Resource = (*KafkaStrimziIoKafkaUserV1Beta2Resource)(nil)
)

type KafkaStrimziIoKafkaUserV1Beta2TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type KafkaStrimziIoKafkaUserV1Beta2GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Authentication *struct {
			Password *struct {
				ValueFrom *struct {
					SecretKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
			} `tfsdk:"password" yaml:"password,omitempty"`

			Type *string `tfsdk:"type" yaml:"type,omitempty"`
		} `tfsdk:"authentication" yaml:"authentication,omitempty"`

		Authorization *struct {
			Acls *[]struct {
				Host *string `tfsdk:"host" yaml:"host,omitempty"`

				Operation *string `tfsdk:"operation" yaml:"operation,omitempty"`

				Resource *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					PatternType *string `tfsdk:"pattern_type" yaml:"patternType,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"resource" yaml:"resource,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`
			} `tfsdk:"acls" yaml:"acls,omitempty"`

			Type *string `tfsdk:"type" yaml:"type,omitempty"`
		} `tfsdk:"authorization" yaml:"authorization,omitempty"`

		Quotas *struct {
			ConsumerByteRate *int64 `tfsdk:"consumer_byte_rate" yaml:"consumerByteRate,omitempty"`

			ControllerMutationRate *float64 `tfsdk:"controller_mutation_rate" yaml:"controllerMutationRate,omitempty"`

			ProducerByteRate *int64 `tfsdk:"producer_byte_rate" yaml:"producerByteRate,omitempty"`

			RequestPercentage *int64 `tfsdk:"request_percentage" yaml:"requestPercentage,omitempty"`
		} `tfsdk:"quotas" yaml:"quotas,omitempty"`

		Template *struct {
			Secret *struct {
				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

					Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`
				} `tfsdk:"metadata" yaml:"metadata,omitempty"`
			} `tfsdk:"secret" yaml:"secret,omitempty"`
		} `tfsdk:"template" yaml:"template,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewKafkaStrimziIoKafkaUserV1Beta2Resource() resource.Resource {
	return &KafkaStrimziIoKafkaUserV1Beta2Resource{}
}

func (r *KafkaStrimziIoKafkaUserV1Beta2Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_kafka_strimzi_io_kafka_user_v1beta2"
}

func (r *KafkaStrimziIoKafkaUserV1Beta2Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "The specification of the user.",
				MarkdownDescription: "The specification of the user.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"authentication": {
						Description:         "Authentication mechanism enabled for this Kafka user. The supported authentication mechanisms are 'scram-sha-512', 'tls', and 'tls-external'. * 'scram-sha-512' generates a secret with SASL SCRAM-SHA-512 credentials.* 'tls' generates a secret with user certificate for mutual TLS authentication.* 'tls-external' does not generate a user certificate.   But prepares the user for using mutual TLS authentication using a user certificate generated outside the User Operator.  ACLs and quotas set for this user are configured in the 'CN=<username>' format.Authentication is optional. If authentication is not configured, no credentials are generated. ACLs and quotas set for the user are configured in the '<username>' format suitable for SASL authentication.",
						MarkdownDescription: "Authentication mechanism enabled for this Kafka user. The supported authentication mechanisms are 'scram-sha-512', 'tls', and 'tls-external'. * 'scram-sha-512' generates a secret with SASL SCRAM-SHA-512 credentials.* 'tls' generates a secret with user certificate for mutual TLS authentication.* 'tls-external' does not generate a user certificate.   But prepares the user for using mutual TLS authentication using a user certificate generated outside the User Operator.  ACLs and quotas set for this user are configured in the 'CN=<username>' format.Authentication is optional. If authentication is not configured, no credentials are generated. ACLs and quotas set for the user are configured in the '<username>' format suitable for SASL authentication.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"password": {
								Description:         "Specify the password for the user. If not set, a new password is generated by the User Operator.",
								MarkdownDescription: "Specify the password for the user. If not set, a new password is generated by the User Operator.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"value_from": {
										Description:         "Secret from which the password should be read.",
										MarkdownDescription: "Secret from which the password should be read.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"secret_key_ref": {
												Description:         "Selects a key of a Secret in the resource's namespace.",
												MarkdownDescription: "Selects a key of a Secret in the resource's namespace.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"optional": {
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

							"type": {
								Description:         "Authentication type.",
								MarkdownDescription: "Authentication type.",

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

					"authorization": {
						Description:         "Authorization rules for this Kafka user.",
						MarkdownDescription: "Authorization rules for this Kafka user.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"acls": {
								Description:         "List of ACL rules which should be applied to this user.",
								MarkdownDescription: "List of ACL rules which should be applied to this user.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"host": {
										Description:         "The host from which the action described in the ACL rule is allowed or denied.",
										MarkdownDescription: "The host from which the action described in the ACL rule is allowed or denied.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"operation": {
										Description:         "Operation which will be allowed or denied. Supported operations are: Read, Write, Create, Delete, Alter, Describe, ClusterAction, AlterConfigs, DescribeConfigs, IdempotentWrite and All.",
										MarkdownDescription: "Operation which will be allowed or denied. Supported operations are: Read, Write, Create, Delete, Alter, Describe, ClusterAction, AlterConfigs, DescribeConfigs, IdempotentWrite and All.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"resource": {
										Description:         "Indicates the resource for which given ACL rule applies.",
										MarkdownDescription: "Indicates the resource for which given ACL rule applies.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "Name of resource for which given ACL rule applies. Can be combined with 'patternType' field to use prefix pattern.",
												MarkdownDescription: "Name of resource for which given ACL rule applies. Can be combined with 'patternType' field to use prefix pattern.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pattern_type": {
												Description:         "Describes the pattern used in the resource field. The supported types are 'literal' and 'prefix'. With 'literal' pattern type, the resource field will be used as a definition of a full name. With 'prefix' pattern type, the resource name will be used only as a prefix. Default value is 'literal'.",
												MarkdownDescription: "Describes the pattern used in the resource field. The supported types are 'literal' and 'prefix'. With 'literal' pattern type, the resource field will be used as a definition of a full name. With 'prefix' pattern type, the resource name will be used only as a prefix. Default value is 'literal'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"type": {
												Description:         "Resource type. The available resource types are 'topic', 'group', 'cluster', and 'transactionalId'.",
												MarkdownDescription: "Resource type. The available resource types are 'topic', 'group', 'cluster', and 'transactionalId'.",

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

									"type": {
										Description:         "The type of the rule. Currently the only supported type is 'allow'. ACL rules with type 'allow' are used to allow user to execute the specified operations. Default value is 'allow'.",
										MarkdownDescription: "The type of the rule. Currently the only supported type is 'allow'. ACL rules with type 'allow' are used to allow user to execute the specified operations. Default value is 'allow'.",

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

							"type": {
								Description:         "Authorization type. Currently the only supported type is 'simple'. 'simple' authorization type uses Kafka's 'kafka.security.authorizer.AclAuthorizer' class for authorization.",
								MarkdownDescription: "Authorization type. Currently the only supported type is 'simple'. 'simple' authorization type uses Kafka's 'kafka.security.authorizer.AclAuthorizer' class for authorization.",

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

					"quotas": {
						Description:         "Quotas on requests to control the broker resources used by clients. Network bandwidth and request rate quotas can be enforced.Kafka documentation for Kafka User quotas can be found at http://kafka.apache.org/documentation/#design_quotas.",
						MarkdownDescription: "Quotas on requests to control the broker resources used by clients. Network bandwidth and request rate quotas can be enforced.Kafka documentation for Kafka User quotas can be found at http://kafka.apache.org/documentation/#design_quotas.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"consumer_byte_rate": {
								Description:         "A quota on the maximum bytes per-second that each client group can fetch from a broker before the clients in the group are throttled. Defined on a per-broker basis.",
								MarkdownDescription: "A quota on the maximum bytes per-second that each client group can fetch from a broker before the clients in the group are throttled. Defined on a per-broker basis.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),
								},
							},

							"controller_mutation_rate": {
								Description:         "A quota on the rate at which mutations are accepted for the create topics request, the create partitions request and the delete topics request. The rate is accumulated by the number of partitions created or deleted.",
								MarkdownDescription: "A quota on the rate at which mutations are accepted for the create topics request, the create partitions request and the delete topics request. The rate is accumulated by the number of partitions created or deleted.",

								Type: types.NumberType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									float64validator.AtLeast(0),
								},
							},

							"producer_byte_rate": {
								Description:         "A quota on the maximum bytes per-second that each client group can publish to a broker before the clients in the group are throttled. Defined on a per-broker basis.",
								MarkdownDescription: "A quota on the maximum bytes per-second that each client group can publish to a broker before the clients in the group are throttled. Defined on a per-broker basis.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),
								},
							},

							"request_percentage": {
								Description:         "A quota on the maximum CPU utilization of each client group as a percentage of network and I/O threads.",
								MarkdownDescription: "A quota on the maximum CPU utilization of each client group as a percentage of network and I/O threads.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"template": {
						Description:         "Template to specify how Kafka User 'Secrets' are generated.",
						MarkdownDescription: "Template to specify how Kafka User 'Secrets' are generated.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"secret": {
								Description:         "Template for KafkaUser resources. The template allows users to specify how the 'Secret' with password or TLS certificates is generated.",
								MarkdownDescription: "Template for KafkaUser resources. The template allows users to specify how the 'Secret' with password or TLS certificates is generated.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"metadata": {
										Description:         "Metadata applied to the resource.",
										MarkdownDescription: "Metadata applied to the resource.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"annotations": {
												Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
												MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"labels": {
												Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
												MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

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
		},
	}, nil
}

func (r *KafkaStrimziIoKafkaUserV1Beta2Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_kafka_strimzi_io_kafka_user_v1beta2")

	var state KafkaStrimziIoKafkaUserV1Beta2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel KafkaStrimziIoKafkaUserV1Beta2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("kafka.strimzi.io/v1beta2")
	goModel.Kind = utilities.Ptr("KafkaUser")

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

func (r *KafkaStrimziIoKafkaUserV1Beta2Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kafka_strimzi_io_kafka_user_v1beta2")
	// NO-OP: All data is already in Terraform state
}

func (r *KafkaStrimziIoKafkaUserV1Beta2Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_kafka_strimzi_io_kafka_user_v1beta2")

	var state KafkaStrimziIoKafkaUserV1Beta2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel KafkaStrimziIoKafkaUserV1Beta2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("kafka.strimzi.io/v1beta2")
	goModel.Kind = utilities.Ptr("KafkaUser")

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

func (r *KafkaStrimziIoKafkaUserV1Beta2Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_kafka_strimzi_io_kafka_user_v1beta2")
	// NO-OP: Terraform removes the state automatically for us
}
