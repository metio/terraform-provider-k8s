/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package tf_tungsten_io_v1alpha1

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
	_ datasource.DataSource = &TfTungstenIoRabbitmqV1Alpha1Manifest{}
)

func NewTfTungstenIoRabbitmqV1Alpha1Manifest() datasource.DataSource {
	return &TfTungstenIoRabbitmqV1Alpha1Manifest{}
}

type TfTungstenIoRabbitmqV1Alpha1Manifest struct{}

type TfTungstenIoRabbitmqV1Alpha1ManifestData struct {
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
		CommonConfiguration *struct {
			AuthParameters *struct {
				AuthMode               *string `tfsdk:"auth_mode" json:"authMode,omitempty"`
				KeystoneAuthParameters *struct {
					Address           *string `tfsdk:"address" json:"address,omitempty"`
					AdminPassword     *string `tfsdk:"admin_password" json:"adminPassword,omitempty"`
					AdminPort         *int64  `tfsdk:"admin_port" json:"adminPort,omitempty"`
					AdminTenant       *string `tfsdk:"admin_tenant" json:"adminTenant,omitempty"`
					AdminUsername     *string `tfsdk:"admin_username" json:"adminUsername,omitempty"`
					AuthProtocol      *string `tfsdk:"auth_protocol" json:"authProtocol,omitempty"`
					Insecure          *bool   `tfsdk:"insecure" json:"insecure,omitempty"`
					Port              *int64  `tfsdk:"port" json:"port,omitempty"`
					ProjectDomainName *string `tfsdk:"project_domain_name" json:"projectDomainName,omitempty"`
					Region            *string `tfsdk:"region" json:"region,omitempty"`
					UserDomainName    *string `tfsdk:"user_domain_name" json:"userDomainName,omitempty"`
				} `tfsdk:"keystone_auth_parameters" json:"keystoneAuthParameters,omitempty"`
				KeystoneSecretName *string `tfsdk:"keystone_secret_name" json:"keystoneSecretName,omitempty"`
			} `tfsdk:"auth_parameters" json:"authParameters,omitempty"`
			Distribution     *string            `tfsdk:"distribution" json:"distribution,omitempty"`
			ImagePullSecrets *[]string          `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
			LogLevel         *string            `tfsdk:"log_level" json:"logLevel,omitempty"`
			NodeSelector     *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
			Tolerations      *[]struct {
				Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
				Key               *string `tfsdk:"key" json:"key,omitempty"`
				Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
				TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
				Value             *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"tolerations" json:"tolerations,omitempty"`
		} `tfsdk:"common_configuration" json:"commonConfiguration,omitempty"`
		ServiceConfiguration *struct {
			ClusterPartitionHandling *string `tfsdk:"cluster_partition_handling" json:"clusterPartitionHandling,omitempty"`
			Containers               *[]struct {
				Command *[]string `tfsdk:"command" json:"command,omitempty"`
				Image   *string   `tfsdk:"image" json:"image,omitempty"`
				Name    *string   `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"containers" json:"containers,omitempty"`
			ErlEpmdPort       *int64  `tfsdk:"erl_epmd_port" json:"erlEpmdPort,omitempty"`
			ErlangCookie      *string `tfsdk:"erlang_cookie" json:"erlangCookie,omitempty"`
			MirroredQueueMode *string `tfsdk:"mirrored_queue_mode" json:"mirroredQueueMode,omitempty"`
			Password          *string `tfsdk:"password" json:"password,omitempty"`
			Port              *int64  `tfsdk:"port" json:"port,omitempty"`
			Secret            *string `tfsdk:"secret" json:"secret,omitempty"`
			TcpListenOptions  *struct {
				Backlog       *int64 `tfsdk:"backlog" json:"backlog,omitempty"`
				ExitOnClose   *bool  `tfsdk:"exit_on_close" json:"exitOnClose,omitempty"`
				LingerOn      *bool  `tfsdk:"linger_on" json:"lingerOn,omitempty"`
				LingerTimeout *int64 `tfsdk:"linger_timeout" json:"lingerTimeout,omitempty"`
				Nodelay       *bool  `tfsdk:"nodelay" json:"nodelay,omitempty"`
			} `tfsdk:"tcp_listen_options" json:"tcpListenOptions,omitempty"`
			User  *string `tfsdk:"user" json:"user,omitempty"`
			Vhost *string `tfsdk:"vhost" json:"vhost,omitempty"`
		} `tfsdk:"service_configuration" json:"serviceConfiguration,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TfTungstenIoRabbitmqV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_tf_tungsten_io_rabbitmq_v1alpha1_manifest"
}

func (r *TfTungstenIoRabbitmqV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Rabbitmq is the Schema for the rabbitmqs API.",
		MarkdownDescription: "Rabbitmq is the Schema for the rabbitmqs API.",
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
				Description:         "RabbitmqSpec is the Spec for the cassandras API.",
				MarkdownDescription: "RabbitmqSpec is the Spec for the cassandras API.",
				Attributes: map[string]schema.Attribute{
					"common_configuration": schema.SingleNestedAttribute{
						Description:         "PodConfiguration is the common services struct.",
						MarkdownDescription: "PodConfiguration is the common services struct.",
						Attributes: map[string]schema.Attribute{
							"auth_parameters": schema.SingleNestedAttribute{
								Description:         "AuthParameters auth parameters",
								MarkdownDescription: "AuthParameters auth parameters",
								Attributes: map[string]schema.Attribute{
									"auth_mode": schema.StringAttribute{
										Description:         "AuthenticationMode auth mode",
										MarkdownDescription: "AuthenticationMode auth mode",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("noauth", "keystone"),
										},
									},

									"keystone_auth_parameters": schema.SingleNestedAttribute{
										Description:         "KeystoneAuthParameters keystone parameters",
										MarkdownDescription: "KeystoneAuthParameters keystone parameters",
										Attributes: map[string]schema.Attribute{
											"address": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"admin_password": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"admin_port": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"admin_tenant": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"admin_username": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"auth_protocol": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"insecure": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"port": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"project_domain_name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"region": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"user_domain_name": schema.StringAttribute{
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

									"keystone_secret_name": schema.StringAttribute{
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

							"distribution": schema.StringAttribute{
								Description:         "OS family",
								MarkdownDescription: "OS family",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image_pull_secrets": schema.ListAttribute{
								Description:         "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
								MarkdownDescription: "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"log_level": schema.StringAttribute{
								Description:         "Kubernetes Cluster Configuration",
								MarkdownDescription: "Kubernetes Cluster Configuration",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("info", "debug", "warning", "error", "critical", "none"),
								},
							},

							"node_selector": schema.MapAttribute{
								Description:         "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
								MarkdownDescription: "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tolerations": schema.ListNestedAttribute{
								Description:         "If specified, the pod's tolerations.",
								MarkdownDescription: "If specified, the pod's tolerations.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"effect": schema.StringAttribute{
											Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
											MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"key": schema.StringAttribute{
											Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
											MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
											MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"toleration_seconds": schema.Int64Attribute{
											Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
											MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
											MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"service_configuration": schema.SingleNestedAttribute{
						Description:         "RabbitmqConfiguration is the Spec for the cassandras API.",
						MarkdownDescription: "RabbitmqConfiguration is the Spec for the cassandras API.",
						Attributes: map[string]schema.Attribute{
							"cluster_partition_handling": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"containers": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"command": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"image": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
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

							"erl_epmd_port": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"erlang_cookie": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"mirrored_queue_mode": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("exactly", "all", "nodes"),
								},
							},

							"password": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secret": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tcp_listen_options": schema.SingleNestedAttribute{
								Description:         "TCPListenOptionsConfig is configuration for RabbitMQ TCP listen",
								MarkdownDescription: "TCPListenOptionsConfig is configuration for RabbitMQ TCP listen",
								Attributes: map[string]schema.Attribute{
									"backlog": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"exit_on_close": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"linger_on": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"linger_timeout": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"nodelay": schema.BoolAttribute{
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

							"user": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"vhost": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
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
	}
}

func (r *TfTungstenIoRabbitmqV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_tf_tungsten_io_rabbitmq_v1alpha1_manifest")

	var model TfTungstenIoRabbitmqV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("tf.tungsten.io/v1alpha1")
	model.Kind = pointer.String("Rabbitmq")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
