/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kafka_strimzi_io_v1beta1

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
	_ datasource.DataSource              = &KafkaStrimziIoKafkaUserV1Beta1DataSource{}
	_ datasource.DataSourceWithConfigure = &KafkaStrimziIoKafkaUserV1Beta1DataSource{}
)

func NewKafkaStrimziIoKafkaUserV1Beta1DataSource() datasource.DataSource {
	return &KafkaStrimziIoKafkaUserV1Beta1DataSource{}
}

type KafkaStrimziIoKafkaUserV1Beta1DataSource struct {
	kubernetesClient dynamic.Interface
}

type KafkaStrimziIoKafkaUserV1Beta1DataSourceData struct {
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
		Authentication *struct {
			Password *struct {
				ValueFrom *struct {
					SecretKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"password" json:"password,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"authentication" json:"authentication,omitempty"`
		Authorization *struct {
			Acls *[]struct {
				Host       *string   `tfsdk:"host" json:"host,omitempty"`
				Operation  *string   `tfsdk:"operation" json:"operation,omitempty"`
				Operations *[]string `tfsdk:"operations" json:"operations,omitempty"`
				Resource   *struct {
					Name        *string `tfsdk:"name" json:"name,omitempty"`
					PatternType *string `tfsdk:"pattern_type" json:"patternType,omitempty"`
					Type        *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"resource" json:"resource,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"acls" json:"acls,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"authorization" json:"authorization,omitempty"`
		Quotas *struct {
			ConsumerByteRate       *int64     `tfsdk:"consumer_byte_rate" json:"consumerByteRate,omitempty"`
			ControllerMutationRate *big.Float `tfsdk:"controller_mutation_rate" json:"controllerMutationRate,omitempty"`
			ProducerByteRate       *int64     `tfsdk:"producer_byte_rate" json:"producerByteRate,omitempty"`
			RequestPercentage      *int64     `tfsdk:"request_percentage" json:"requestPercentage,omitempty"`
		} `tfsdk:"quotas" json:"quotas,omitempty"`
		Template *struct {
			Secret *struct {
				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
			} `tfsdk:"secret" json:"secret,omitempty"`
		} `tfsdk:"template" json:"template,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KafkaStrimziIoKafkaUserV1Beta1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kafka_strimzi_io_kafka_user_v1beta1"
}

func (r *KafkaStrimziIoKafkaUserV1Beta1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "The specification of the user.",
				MarkdownDescription: "The specification of the user.",
				Attributes: map[string]schema.Attribute{
					"authentication": schema.SingleNestedAttribute{
						Description:         "Authentication mechanism enabled for this Kafka user. The supported authentication mechanisms are 'scram-sha-512', 'tls', and 'tls-external'. * 'scram-sha-512' generates a secret with SASL SCRAM-SHA-512 credentials.* 'tls' generates a secret with user certificate for mutual TLS authentication.* 'tls-external' does not generate a user certificate.   But prepares the user for using mutual TLS authentication using a user certificate generated outside the User Operator.  ACLs and quotas set for this user are configured in the 'CN=<username>' format.Authentication is optional. If authentication is not configured, no credentials are generated. ACLs and quotas set for the user are configured in the '<username>' format suitable for SASL authentication.",
						MarkdownDescription: "Authentication mechanism enabled for this Kafka user. The supported authentication mechanisms are 'scram-sha-512', 'tls', and 'tls-external'. * 'scram-sha-512' generates a secret with SASL SCRAM-SHA-512 credentials.* 'tls' generates a secret with user certificate for mutual TLS authentication.* 'tls-external' does not generate a user certificate.   But prepares the user for using mutual TLS authentication using a user certificate generated outside the User Operator.  ACLs and quotas set for this user are configured in the 'CN=<username>' format.Authentication is optional. If authentication is not configured, no credentials are generated. ACLs and quotas set for the user are configured in the '<username>' format suitable for SASL authentication.",
						Attributes: map[string]schema.Attribute{
							"password": schema.SingleNestedAttribute{
								Description:         "Specify the password for the user. If not set, a new password is generated by the User Operator.",
								MarkdownDescription: "Specify the password for the user. If not set, a new password is generated by the User Operator.",
								Attributes: map[string]schema.Attribute{
									"value_from": schema.SingleNestedAttribute{
										Description:         "Secret from which the password should be read.",
										MarkdownDescription: "Secret from which the password should be read.",
										Attributes: map[string]schema.Attribute{
											"secret_key_ref": schema.SingleNestedAttribute{
												Description:         "Selects a key of a Secret in the resource's namespace.",
												MarkdownDescription: "Selects a key of a Secret in the resource's namespace.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"optional": schema.BoolAttribute{
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

							"type": schema.StringAttribute{
								Description:         "Authentication type.",
								MarkdownDescription: "Authentication type.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"authorization": schema.SingleNestedAttribute{
						Description:         "Authorization rules for this Kafka user.",
						MarkdownDescription: "Authorization rules for this Kafka user.",
						Attributes: map[string]schema.Attribute{
							"acls": schema.ListNestedAttribute{
								Description:         "List of ACL rules which should be applied to this user.",
								MarkdownDescription: "List of ACL rules which should be applied to this user.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"host": schema.StringAttribute{
											Description:         "The host from which the action described in the ACL rule is allowed or denied.",
											MarkdownDescription: "The host from which the action described in the ACL rule is allowed or denied.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"operation": schema.StringAttribute{
											Description:         "Operation which will be allowed or denied. Supported operations are: Read, Write, Create, Delete, Alter, Describe, ClusterAction, AlterConfigs, DescribeConfigs, IdempotentWrite and All.",
											MarkdownDescription: "Operation which will be allowed or denied. Supported operations are: Read, Write, Create, Delete, Alter, Describe, ClusterAction, AlterConfigs, DescribeConfigs, IdempotentWrite and All.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"operations": schema.ListAttribute{
											Description:         "List of operations which will be allowed or denied. Supported operations are: Read, Write, Create, Delete, Alter, Describe, ClusterAction, AlterConfigs, DescribeConfigs, IdempotentWrite and All.",
											MarkdownDescription: "List of operations which will be allowed or denied. Supported operations are: Read, Write, Create, Delete, Alter, Describe, ClusterAction, AlterConfigs, DescribeConfigs, IdempotentWrite and All.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"resource": schema.SingleNestedAttribute{
											Description:         "Indicates the resource for which given ACL rule applies.",
											MarkdownDescription: "Indicates the resource for which given ACL rule applies.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of resource for which given ACL rule applies. Can be combined with 'patternType' field to use prefix pattern.",
													MarkdownDescription: "Name of resource for which given ACL rule applies. Can be combined with 'patternType' field to use prefix pattern.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"pattern_type": schema.StringAttribute{
													Description:         "Describes the pattern used in the resource field. The supported types are 'literal' and 'prefix'. With 'literal' pattern type, the resource field will be used as a definition of a full name. With 'prefix' pattern type, the resource name will be used only as a prefix. Default value is 'literal'.",
													MarkdownDescription: "Describes the pattern used in the resource field. The supported types are 'literal' and 'prefix'. With 'literal' pattern type, the resource field will be used as a definition of a full name. With 'prefix' pattern type, the resource name will be used only as a prefix. Default value is 'literal'.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"type": schema.StringAttribute{
													Description:         "Resource type. The available resource types are 'topic', 'group', 'cluster', and 'transactionalId'.",
													MarkdownDescription: "Resource type. The available resource types are 'topic', 'group', 'cluster', and 'transactionalId'.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"type": schema.StringAttribute{
											Description:         "The type of the rule. Currently the only supported type is 'allow'. ACL rules with type 'allow' are used to allow user to execute the specified operations. Default value is 'allow'.",
											MarkdownDescription: "The type of the rule. Currently the only supported type is 'allow'. ACL rules with type 'allow' are used to allow user to execute the specified operations. Default value is 'allow'.",
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

							"type": schema.StringAttribute{
								Description:         "Authorization type. Currently the only supported type is 'simple'. 'simple' authorization type uses Kafka's 'kafka.security.authorizer.AclAuthorizer' class for authorization.",
								MarkdownDescription: "Authorization type. Currently the only supported type is 'simple'. 'simple' authorization type uses Kafka's 'kafka.security.authorizer.AclAuthorizer' class for authorization.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"quotas": schema.SingleNestedAttribute{
						Description:         "Quotas on requests to control the broker resources used by clients. Network bandwidth and request rate quotas can be enforced.Kafka documentation for Kafka User quotas can be found at http://kafka.apache.org/documentation/#design_quotas.",
						MarkdownDescription: "Quotas on requests to control the broker resources used by clients. Network bandwidth and request rate quotas can be enforced.Kafka documentation for Kafka User quotas can be found at http://kafka.apache.org/documentation/#design_quotas.",
						Attributes: map[string]schema.Attribute{
							"consumer_byte_rate": schema.Int64Attribute{
								Description:         "A quota on the maximum bytes per-second that each client group can fetch from a broker before the clients in the group are throttled. Defined on a per-broker basis.",
								MarkdownDescription: "A quota on the maximum bytes per-second that each client group can fetch from a broker before the clients in the group are throttled. Defined on a per-broker basis.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"controller_mutation_rate": types.NumberType{
								Description:         "A quota on the rate at which mutations are accepted for the create topics request, the create partitions request and the delete topics request. The rate is accumulated by the number of partitions created or deleted.",
								MarkdownDescription: "A quota on the rate at which mutations are accepted for the create topics request, the create partitions request and the delete topics request. The rate is accumulated by the number of partitions created or deleted.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"producer_byte_rate": schema.Int64Attribute{
								Description:         "A quota on the maximum bytes per-second that each client group can publish to a broker before the clients in the group are throttled. Defined on a per-broker basis.",
								MarkdownDescription: "A quota on the maximum bytes per-second that each client group can publish to a broker before the clients in the group are throttled. Defined on a per-broker basis.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"request_percentage": schema.Int64Attribute{
								Description:         "A quota on the maximum CPU utilization of each client group as a percentage of network and I/O threads.",
								MarkdownDescription: "A quota on the maximum CPU utilization of each client group as a percentage of network and I/O threads.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"template": schema.SingleNestedAttribute{
						Description:         "Template to specify how Kafka User 'Secrets' are generated.",
						MarkdownDescription: "Template to specify how Kafka User 'Secrets' are generated.",
						Attributes: map[string]schema.Attribute{
							"secret": schema.SingleNestedAttribute{
								Description:         "Template for KafkaUser resources. The template allows users to specify how the 'Secret' with password or TLS certificates is generated.",
								MarkdownDescription: "Template for KafkaUser resources. The template allows users to specify how the 'Secret' with password or TLS certificates is generated.",
								Attributes: map[string]schema.Attribute{
									"metadata": schema.SingleNestedAttribute{
										Description:         "Metadata applied to the resource.",
										MarkdownDescription: "Metadata applied to the resource.",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "Annotations added to the Kubernetes resource.",
												MarkdownDescription: "Annotations added to the Kubernetes resource.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"labels": schema.MapAttribute{
												Description:         "Labels added to the Kubernetes resource.",
												MarkdownDescription: "Labels added to the Kubernetes resource.",
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
		},
	}
}

func (r *KafkaStrimziIoKafkaUserV1Beta1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *KafkaStrimziIoKafkaUserV1Beta1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_kafka_strimzi_io_kafka_user_v1beta1")

	var data KafkaStrimziIoKafkaUserV1Beta1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "kafka.strimzi.io", Version: "v1beta1", Resource: "KafkaUser"}).
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

	var readResponse KafkaStrimziIoKafkaUserV1Beta1DataSourceData
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
	data.ApiVersion = pointer.String("kafka.strimzi.io/v1beta1")
	data.Kind = pointer.String("KafkaUser")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
