/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package operator_open_cluster_management_io_v1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	_ datasource.DataSource = &OperatorOpenClusterManagementIoClusterManagerV1Manifest{}
)

func NewOperatorOpenClusterManagementIoClusterManagerV1Manifest() datasource.DataSource {
	return &OperatorOpenClusterManagementIoClusterManagerV1Manifest{}
}

type OperatorOpenClusterManagementIoClusterManagerV1Manifest struct{}

type OperatorOpenClusterManagementIoClusterManagerV1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		AddOnManagerConfiguration *struct {
			FeatureGates *[]struct {
				Feature *string `tfsdk:"feature" json:"feature,omitempty"`
				Mode    *string `tfsdk:"mode" json:"mode,omitempty"`
			} `tfsdk:"feature_gates" json:"featureGates,omitempty"`
		} `tfsdk:"add_on_manager_configuration" json:"addOnManagerConfiguration,omitempty"`
		AddOnManagerImagePullSpec *string `tfsdk:"add_on_manager_image_pull_spec" json:"addOnManagerImagePullSpec,omitempty"`
		DeployOption              *struct {
			Hosted *struct {
				RegistrationWebhookConfiguration *struct {
					Address *string `tfsdk:"address" json:"address,omitempty"`
					Port    *int64  `tfsdk:"port" json:"port,omitempty"`
				} `tfsdk:"registration_webhook_configuration" json:"registrationWebhookConfiguration,omitempty"`
				WorkWebhookConfiguration *struct {
					Address *string `tfsdk:"address" json:"address,omitempty"`
					Port    *int64  `tfsdk:"port" json:"port,omitempty"`
				} `tfsdk:"work_webhook_configuration" json:"workWebhookConfiguration,omitempty"`
			} `tfsdk:"hosted" json:"hosted,omitempty"`
			Mode *string `tfsdk:"mode" json:"mode,omitempty"`
		} `tfsdk:"deploy_option" json:"deployOption,omitempty"`
		NodePlacement *struct {
			NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
			Tolerations  *[]struct {
				Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
				Key               *string `tfsdk:"key" json:"key,omitempty"`
				Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
				TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
				Value             *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"tolerations" json:"tolerations,omitempty"`
		} `tfsdk:"node_placement" json:"nodePlacement,omitempty"`
		PlacementImagePullSpec    *string `tfsdk:"placement_image_pull_spec" json:"placementImagePullSpec,omitempty"`
		RegistrationConfiguration *struct {
			AutoApproveUsers *[]string `tfsdk:"auto_approve_users" json:"autoApproveUsers,omitempty"`
			FeatureGates     *[]struct {
				Feature *string `tfsdk:"feature" json:"feature,omitempty"`
				Mode    *string `tfsdk:"mode" json:"mode,omitempty"`
			} `tfsdk:"feature_gates" json:"featureGates,omitempty"`
		} `tfsdk:"registration_configuration" json:"registrationConfiguration,omitempty"`
		RegistrationImagePullSpec *string `tfsdk:"registration_image_pull_spec" json:"registrationImagePullSpec,omitempty"`
		WorkConfiguration         *struct {
			FeatureGates *[]struct {
				Feature *string `tfsdk:"feature" json:"feature,omitempty"`
				Mode    *string `tfsdk:"mode" json:"mode,omitempty"`
			} `tfsdk:"feature_gates" json:"featureGates,omitempty"`
		} `tfsdk:"work_configuration" json:"workConfiguration,omitempty"`
		WorkImagePullSpec *string `tfsdk:"work_image_pull_spec" json:"workImagePullSpec,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OperatorOpenClusterManagementIoClusterManagerV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operator_open_cluster_management_io_cluster_manager_v1_manifest"
}

func (r *OperatorOpenClusterManagementIoClusterManagerV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClusterManager configures the controllers on the hub that govern registration and work distribution for attached Klusterlets. In Default mode, ClusterManager will only be deployed in open-cluster-management-hub namespace. In Hosted mode, ClusterManager will be deployed in the namespace with the same name as cluster manager.",
		MarkdownDescription: "ClusterManager configures the controllers on the hub that govern registration and work distribution for attached Klusterlets. In Default mode, ClusterManager will only be deployed in open-cluster-management-hub namespace. In Hosted mode, ClusterManager will be deployed in the namespace with the same name as cluster manager.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

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
				Description:         "Spec represents a desired deployment configuration of controllers that govern registration and work distribution for attached Klusterlets.",
				MarkdownDescription: "Spec represents a desired deployment configuration of controllers that govern registration and work distribution for attached Klusterlets.",
				Attributes: map[string]schema.Attribute{
					"add_on_manager_configuration": schema.SingleNestedAttribute{
						Description:         "AddOnManagerConfiguration contains the configuration of addon manager",
						MarkdownDescription: "AddOnManagerConfiguration contains the configuration of addon manager",
						Attributes: map[string]schema.Attribute{
							"feature_gates": schema.ListNestedAttribute{
								Description:         "FeatureGates represents the list of feature gates for addon manager If it is set empty, default feature gates will be used. If it is set, featuregate/Foo is an example of one item in FeatureGates: 1. If featuregate/Foo does not exist, registration-operator will discard it 2. If featuregate/Foo exists and is false by default. It is now possible to set featuregate/Foo=[false|true] 3. If featuregate/Foo exists and is true by default. If a cluster-admin upgrading from 1 to 2 wants to continue having featuregate/Foo=false, he can set featuregate/Foo=false before upgrading. Let's say the cluster-admin wants featuregate/Foo=false.",
								MarkdownDescription: "FeatureGates represents the list of feature gates for addon manager If it is set empty, default feature gates will be used. If it is set, featuregate/Foo is an example of one item in FeatureGates: 1. If featuregate/Foo does not exist, registration-operator will discard it 2. If featuregate/Foo exists and is false by default. It is now possible to set featuregate/Foo=[false|true] 3. If featuregate/Foo exists and is true by default. If a cluster-admin upgrading from 1 to 2 wants to continue having featuregate/Foo=false, he can set featuregate/Foo=false before upgrading. Let's say the cluster-admin wants featuregate/Foo=false.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"feature": schema.StringAttribute{
											Description:         "Feature is the key of feature gate. e.g. featuregate/Foo.",
											MarkdownDescription: "Feature is the key of feature gate. e.g. featuregate/Foo.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"mode": schema.StringAttribute{
											Description:         "Mode is either Enable, Disable, '' where '' is Disable by default. In Enable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=true'. In Disable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=false'.",
											MarkdownDescription: "Mode is either Enable, Disable, '' where '' is Disable by default. In Enable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=true'. In Disable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=false'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Enable", "Disable"),
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

					"add_on_manager_image_pull_spec": schema.StringAttribute{
						Description:         "AddOnManagerImagePullSpec represents the desired image configuration of addon manager controller/webhook installed on hub.",
						MarkdownDescription: "AddOnManagerImagePullSpec represents the desired image configuration of addon manager controller/webhook installed on hub.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"deploy_option": schema.SingleNestedAttribute{
						Description:         "DeployOption contains the options of deploying a cluster-manager Default mode is used if DeployOption is not set.",
						MarkdownDescription: "DeployOption contains the options of deploying a cluster-manager Default mode is used if DeployOption is not set.",
						Attributes: map[string]schema.Attribute{
							"hosted": schema.SingleNestedAttribute{
								Description:         "Hosted includes configurations we needs for clustermanager in the Hosted mode.",
								MarkdownDescription: "Hosted includes configurations we needs for clustermanager in the Hosted mode.",
								Attributes: map[string]schema.Attribute{
									"registration_webhook_configuration": schema.SingleNestedAttribute{
										Description:         "RegistrationWebhookConfiguration represents the customized webhook-server configuration of registration.",
										MarkdownDescription: "RegistrationWebhookConfiguration represents the customized webhook-server configuration of registration.",
										Attributes: map[string]schema.Attribute{
											"address": schema.StringAttribute{
												Description:         "Address represents the address of a webhook-server. It could be in IP format or fqdn format. The Address must be reachable by apiserver of the hub cluster.",
												MarkdownDescription: "Address represents the address of a webhook-server. It could be in IP format or fqdn format. The Address must be reachable by apiserver of the hub cluster.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])$`), ""),
												},
											},

											"port": schema.Int64Attribute{
												Description:         "Port represents the port of a webhook-server. The default value of Port is 443.",
												MarkdownDescription: "Port represents the port of a webhook-server. The default value of Port is 443.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtMost(65535),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"work_webhook_configuration": schema.SingleNestedAttribute{
										Description:         "WorkWebhookConfiguration represents the customized webhook-server configuration of work.",
										MarkdownDescription: "WorkWebhookConfiguration represents the customized webhook-server configuration of work.",
										Attributes: map[string]schema.Attribute{
											"address": schema.StringAttribute{
												Description:         "Address represents the address of a webhook-server. It could be in IP format or fqdn format. The Address must be reachable by apiserver of the hub cluster.",
												MarkdownDescription: "Address represents the address of a webhook-server. It could be in IP format or fqdn format. The Address must be reachable by apiserver of the hub cluster.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])$`), ""),
												},
											},

											"port": schema.Int64Attribute{
												Description:         "Port represents the port of a webhook-server. The default value of Port is 443.",
												MarkdownDescription: "Port represents the port of a webhook-server. The default value of Port is 443.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtMost(65535),
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

							"mode": schema.StringAttribute{
								Description:         "Mode can be Default or Hosted. In Default mode, the Hub is installed as a whole and all parts of Hub are deployed in the same cluster. In Hosted mode, only crd and configurations are installed on one cluster(defined as hub-cluster). Controllers run in another cluster (defined as management-cluster) and connect to the hub with the kubeconfig in secret of 'external-hub-kubeconfig'(a kubeconfig of hub-cluster with cluster-admin permission). Note: Do not modify the Mode field once it's applied.",
								MarkdownDescription: "Mode can be Default or Hosted. In Default mode, the Hub is installed as a whole and all parts of Hub are deployed in the same cluster. In Hosted mode, only crd and configurations are installed on one cluster(defined as hub-cluster). Controllers run in another cluster (defined as management-cluster) and connect to the hub with the kubeconfig in secret of 'external-hub-kubeconfig'(a kubeconfig of hub-cluster with cluster-admin permission). Note: Do not modify the Mode field once it's applied.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Default", "Hosted"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"node_placement": schema.SingleNestedAttribute{
						Description:         "NodePlacement enables explicit control over the scheduling of the deployed pods.",
						MarkdownDescription: "NodePlacement enables explicit control over the scheduling of the deployed pods.",
						Attributes: map[string]schema.Attribute{
							"node_selector": schema.MapAttribute{
								Description:         "NodeSelector defines which Nodes the Pods are scheduled on. The default is an empty list.",
								MarkdownDescription: "NodeSelector defines which Nodes the Pods are scheduled on. The default is an empty list.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tolerations": schema.ListNestedAttribute{
								Description:         "Tolerations is attached by pods to tolerate any taint that matches the triple <key,value,effect> using the matching operator <operator>. The default is an empty list.",
								MarkdownDescription: "Tolerations is attached by pods to tolerate any taint that matches the triple <key,value,effect> using the matching operator <operator>. The default is an empty list.",
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

					"placement_image_pull_spec": schema.StringAttribute{
						Description:         "PlacementImagePullSpec represents the desired image configuration of placement controller/webhook installed on hub.",
						MarkdownDescription: "PlacementImagePullSpec represents the desired image configuration of placement controller/webhook installed on hub.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"registration_configuration": schema.SingleNestedAttribute{
						Description:         "RegistrationConfiguration contains the configuration of registration",
						MarkdownDescription: "RegistrationConfiguration contains the configuration of registration",
						Attributes: map[string]schema.Attribute{
							"auto_approve_users": schema.ListAttribute{
								Description:         "AutoApproveUser represents a list of users that can auto approve CSR and accept client. If the credential of the bootstrap-hub-kubeconfig matches to the users, the cluster created by the bootstrap-hub-kubeconfig will be auto-registered into the hub cluster. This takes effect only when ManagedClusterAutoApproval feature gate is enabled.",
								MarkdownDescription: "AutoApproveUser represents a list of users that can auto approve CSR and accept client. If the credential of the bootstrap-hub-kubeconfig matches to the users, the cluster created by the bootstrap-hub-kubeconfig will be auto-registered into the hub cluster. This takes effect only when ManagedClusterAutoApproval feature gate is enabled.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"feature_gates": schema.ListNestedAttribute{
								Description:         "FeatureGates represents the list of feature gates for registration If it is set empty, default feature gates will be used. If it is set, featuregate/Foo is an example of one item in FeatureGates: 1. If featuregate/Foo does not exist, registration-operator will discard it 2. If featuregate/Foo exists and is false by default. It is now possible to set featuregate/Foo=[false|true] 3. If featuregate/Foo exists and is true by default. If a cluster-admin upgrading from 1 to 2 wants to continue having featuregate/Foo=false, he can set featuregate/Foo=false before upgrading. Let's say the cluster-admin wants featuregate/Foo=false.",
								MarkdownDescription: "FeatureGates represents the list of feature gates for registration If it is set empty, default feature gates will be used. If it is set, featuregate/Foo is an example of one item in FeatureGates: 1. If featuregate/Foo does not exist, registration-operator will discard it 2. If featuregate/Foo exists and is false by default. It is now possible to set featuregate/Foo=[false|true] 3. If featuregate/Foo exists and is true by default. If a cluster-admin upgrading from 1 to 2 wants to continue having featuregate/Foo=false, he can set featuregate/Foo=false before upgrading. Let's say the cluster-admin wants featuregate/Foo=false.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"feature": schema.StringAttribute{
											Description:         "Feature is the key of feature gate. e.g. featuregate/Foo.",
											MarkdownDescription: "Feature is the key of feature gate. e.g. featuregate/Foo.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"mode": schema.StringAttribute{
											Description:         "Mode is either Enable, Disable, '' where '' is Disable by default. In Enable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=true'. In Disable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=false'.",
											MarkdownDescription: "Mode is either Enable, Disable, '' where '' is Disable by default. In Enable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=true'. In Disable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=false'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Enable", "Disable"),
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

					"registration_image_pull_spec": schema.StringAttribute{
						Description:         "RegistrationImagePullSpec represents the desired image of registration controller/webhook installed on hub.",
						MarkdownDescription: "RegistrationImagePullSpec represents the desired image of registration controller/webhook installed on hub.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"work_configuration": schema.SingleNestedAttribute{
						Description:         "WorkConfiguration contains the configuration of work",
						MarkdownDescription: "WorkConfiguration contains the configuration of work",
						Attributes: map[string]schema.Attribute{
							"feature_gates": schema.ListNestedAttribute{
								Description:         "FeatureGates represents the list of feature gates for work If it is set empty, default feature gates will be used. If it is set, featuregate/Foo is an example of one item in FeatureGates: 1. If featuregate/Foo does not exist, registration-operator will discard it 2. If featuregate/Foo exists and is false by default. It is now possible to set featuregate/Foo=[false|true] 3. If featuregate/Foo exists and is true by default. If a cluster-admin upgrading from 1 to 2 wants to continue having featuregate/Foo=false, he can set featuregate/Foo=false before upgrading. Let's say the cluster-admin wants featuregate/Foo=false.",
								MarkdownDescription: "FeatureGates represents the list of feature gates for work If it is set empty, default feature gates will be used. If it is set, featuregate/Foo is an example of one item in FeatureGates: 1. If featuregate/Foo does not exist, registration-operator will discard it 2. If featuregate/Foo exists and is false by default. It is now possible to set featuregate/Foo=[false|true] 3. If featuregate/Foo exists and is true by default. If a cluster-admin upgrading from 1 to 2 wants to continue having featuregate/Foo=false, he can set featuregate/Foo=false before upgrading. Let's say the cluster-admin wants featuregate/Foo=false.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"feature": schema.StringAttribute{
											Description:         "Feature is the key of feature gate. e.g. featuregate/Foo.",
											MarkdownDescription: "Feature is the key of feature gate. e.g. featuregate/Foo.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"mode": schema.StringAttribute{
											Description:         "Mode is either Enable, Disable, '' where '' is Disable by default. In Enable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=true'. In Disable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=false'.",
											MarkdownDescription: "Mode is either Enable, Disable, '' where '' is Disable by default. In Enable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=true'. In Disable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=false'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Enable", "Disable"),
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

					"work_image_pull_spec": schema.StringAttribute{
						Description:         "WorkImagePullSpec represents the desired image configuration of work controller/webhook installed on hub.",
						MarkdownDescription: "WorkImagePullSpec represents the desired image configuration of work controller/webhook installed on hub.",
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

func (r *OperatorOpenClusterManagementIoClusterManagerV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_operator_open_cluster_management_io_cluster_manager_v1_manifest")

	var model OperatorOpenClusterManagementIoClusterManagerV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("operator.open-cluster-management.io/v1")
	model.Kind = pointer.String("ClusterManager")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
