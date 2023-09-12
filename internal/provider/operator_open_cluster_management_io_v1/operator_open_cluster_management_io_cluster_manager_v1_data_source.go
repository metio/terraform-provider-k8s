/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package operator_open_cluster_management_io_v1

import (
	"context"
	"encoding/json"
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
	_ datasource.DataSource              = &OperatorOpenClusterManagementIoClusterManagerV1DataSource{}
	_ datasource.DataSourceWithConfigure = &OperatorOpenClusterManagementIoClusterManagerV1DataSource{}
)

func NewOperatorOpenClusterManagementIoClusterManagerV1DataSource() datasource.DataSource {
	return &OperatorOpenClusterManagementIoClusterManagerV1DataSource{}
}

type OperatorOpenClusterManagementIoClusterManagerV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type OperatorOpenClusterManagementIoClusterManagerV1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

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

func (r *OperatorOpenClusterManagementIoClusterManagerV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operator_open_cluster_management_io_cluster_manager_v1"
}

func (r *OperatorOpenClusterManagementIoClusterManagerV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"mode": schema.StringAttribute{
											Description:         "Mode is either Enable, Disable, '' where '' is Disable by default. In Enable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=true'. In Disable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=false'.",
											MarkdownDescription: "Mode is either Enable, Disable, '' where '' is Disable by default. In Enable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=true'. In Disable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=false'.",
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"add_on_manager_image_pull_spec": schema.StringAttribute{
						Description:         "AddOnManagerImagePullSpec represents the desired image configuration of addon manager controller/webhook installed on hub.",
						MarkdownDescription: "AddOnManagerImagePullSpec represents the desired image configuration of addon manager controller/webhook installed on hub.",
						Required:            false,
						Optional:            false,
						Computed:            true,
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
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"port": schema.Int64Attribute{
												Description:         "Port represents the port of a webhook-server. The default value of Port is 443.",
												MarkdownDescription: "Port represents the port of a webhook-server. The default value of Port is 443.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"work_webhook_configuration": schema.SingleNestedAttribute{
										Description:         "WorkWebhookConfiguration represents the customized webhook-server configuration of work.",
										MarkdownDescription: "WorkWebhookConfiguration represents the customized webhook-server configuration of work.",
										Attributes: map[string]schema.Attribute{
											"address": schema.StringAttribute{
												Description:         "Address represents the address of a webhook-server. It could be in IP format or fqdn format. The Address must be reachable by apiserver of the hub cluster.",
												MarkdownDescription: "Address represents the address of a webhook-server. It could be in IP format or fqdn format. The Address must be reachable by apiserver of the hub cluster.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"port": schema.Int64Attribute{
												Description:         "Port represents the port of a webhook-server. The default value of Port is 443.",
												MarkdownDescription: "Port represents the port of a webhook-server. The default value of Port is 443.",
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

							"mode": schema.StringAttribute{
								Description:         "Mode can be Default or Hosted. In Default mode, the Hub is installed as a whole and all parts of Hub are deployed in the same cluster. In Hosted mode, only crd and configurations are installed on one cluster(defined as hub-cluster). Controllers run in another cluster (defined as management-cluster) and connect to the hub with the kubeconfig in secret of 'external-hub-kubeconfig'(a kubeconfig of hub-cluster with cluster-admin permission). Note: Do not modify the Mode field once it's applied.",
								MarkdownDescription: "Mode can be Default or Hosted. In Default mode, the Hub is installed as a whole and all parts of Hub are deployed in the same cluster. In Hosted mode, only crd and configurations are installed on one cluster(defined as hub-cluster). Controllers run in another cluster (defined as management-cluster) and connect to the hub with the kubeconfig in secret of 'external-hub-kubeconfig'(a kubeconfig of hub-cluster with cluster-admin permission). Note: Do not modify the Mode field once it's applied.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
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
								Optional:            false,
								Computed:            true,
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
											Optional:            false,
											Computed:            true,
										},

										"key": schema.StringAttribute{
											Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
											MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"operator": schema.StringAttribute{
											Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
											MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"toleration_seconds": schema.Int64Attribute{
											Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
											MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"value": schema.StringAttribute{
											Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
											MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"placement_image_pull_spec": schema.StringAttribute{
						Description:         "PlacementImagePullSpec represents the desired image configuration of placement controller/webhook installed on hub.",
						MarkdownDescription: "PlacementImagePullSpec represents the desired image configuration of placement controller/webhook installed on hub.",
						Required:            false,
						Optional:            false,
						Computed:            true,
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
								Optional:            false,
								Computed:            true,
							},

							"feature_gates": schema.ListNestedAttribute{
								Description:         "FeatureGates represents the list of feature gates for registration If it is set empty, default feature gates will be used. If it is set, featuregate/Foo is an example of one item in FeatureGates: 1. If featuregate/Foo does not exist, registration-operator will discard it 2. If featuregate/Foo exists and is false by default. It is now possible to set featuregate/Foo=[false|true] 3. If featuregate/Foo exists and is true by default. If a cluster-admin upgrading from 1 to 2 wants to continue having featuregate/Foo=false, he can set featuregate/Foo=false before upgrading. Let's say the cluster-admin wants featuregate/Foo=false.",
								MarkdownDescription: "FeatureGates represents the list of feature gates for registration If it is set empty, default feature gates will be used. If it is set, featuregate/Foo is an example of one item in FeatureGates: 1. If featuregate/Foo does not exist, registration-operator will discard it 2. If featuregate/Foo exists and is false by default. It is now possible to set featuregate/Foo=[false|true] 3. If featuregate/Foo exists and is true by default. If a cluster-admin upgrading from 1 to 2 wants to continue having featuregate/Foo=false, he can set featuregate/Foo=false before upgrading. Let's say the cluster-admin wants featuregate/Foo=false.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"feature": schema.StringAttribute{
											Description:         "Feature is the key of feature gate. e.g. featuregate/Foo.",
											MarkdownDescription: "Feature is the key of feature gate. e.g. featuregate/Foo.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"mode": schema.StringAttribute{
											Description:         "Mode is either Enable, Disable, '' where '' is Disable by default. In Enable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=true'. In Disable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=false'.",
											MarkdownDescription: "Mode is either Enable, Disable, '' where '' is Disable by default. In Enable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=true'. In Disable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=false'.",
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"registration_image_pull_spec": schema.StringAttribute{
						Description:         "RegistrationImagePullSpec represents the desired image of registration controller/webhook installed on hub.",
						MarkdownDescription: "RegistrationImagePullSpec represents the desired image of registration controller/webhook installed on hub.",
						Required:            false,
						Optional:            false,
						Computed:            true,
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
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"mode": schema.StringAttribute{
											Description:         "Mode is either Enable, Disable, '' where '' is Disable by default. In Enable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=true'. In Disable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=false'.",
											MarkdownDescription: "Mode is either Enable, Disable, '' where '' is Disable by default. In Enable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=true'. In Disable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=false'.",
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"work_image_pull_spec": schema.StringAttribute{
						Description:         "WorkImagePullSpec represents the desired image configuration of work controller/webhook installed on hub.",
						MarkdownDescription: "WorkImagePullSpec represents the desired image configuration of work controller/webhook installed on hub.",
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
	}
}

func (r *OperatorOpenClusterManagementIoClusterManagerV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *OperatorOpenClusterManagementIoClusterManagerV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_operator_open_cluster_management_io_cluster_manager_v1")

	var data OperatorOpenClusterManagementIoClusterManagerV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "operator.open-cluster-management.io", Version: "v1", Resource: "clustermanagers"}).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetResourceError(err, data.Metadata.Name))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse OperatorOpenClusterManagementIoClusterManagerV1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("operator.open-cluster-management.io/v1")
	data.Kind = pointer.String("ClusterManager")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
