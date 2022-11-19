/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"

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

type OperatorOpenClusterManagementIoClusterManagerV1Resource struct{}

var (
	_ resource.Resource = (*OperatorOpenClusterManagementIoClusterManagerV1Resource)(nil)
)

type OperatorOpenClusterManagementIoClusterManagerV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type OperatorOpenClusterManagementIoClusterManagerV1GoModel struct {
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
		DeployOption *struct {
			Hosted *struct {
				RegistrationWebhookConfiguration *struct {
					Address *string `tfsdk:"address" yaml:"address,omitempty"`

					Port *int64 `tfsdk:"port" yaml:"port,omitempty"`
				} `tfsdk:"registration_webhook_configuration" yaml:"registrationWebhookConfiguration,omitempty"`

				WorkWebhookConfiguration *struct {
					Address *string `tfsdk:"address" yaml:"address,omitempty"`

					Port *int64 `tfsdk:"port" yaml:"port,omitempty"`
				} `tfsdk:"work_webhook_configuration" yaml:"workWebhookConfiguration,omitempty"`
			} `tfsdk:"hosted" yaml:"hosted,omitempty"`

			Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`
		} `tfsdk:"deploy_option" yaml:"deployOption,omitempty"`

		NodePlacement *struct {
			NodeSelector *map[string]string `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

			Tolerations *[]struct {
				Effect *string `tfsdk:"effect" yaml:"effect,omitempty"`

				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

				TolerationSeconds *int64 `tfsdk:"toleration_seconds" yaml:"tolerationSeconds,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`
			} `tfsdk:"tolerations" yaml:"tolerations,omitempty"`
		} `tfsdk:"node_placement" yaml:"nodePlacement,omitempty"`

		PlacementImagePullSpec *string `tfsdk:"placement_image_pull_spec" yaml:"placementImagePullSpec,omitempty"`

		RegistrationConfiguration *struct {
			FeatureGates *[]struct {
				Feature *string `tfsdk:"feature" yaml:"feature,omitempty"`

				Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`
			} `tfsdk:"feature_gates" yaml:"featureGates,omitempty"`
		} `tfsdk:"registration_configuration" yaml:"registrationConfiguration,omitempty"`

		RegistrationImagePullSpec *string `tfsdk:"registration_image_pull_spec" yaml:"registrationImagePullSpec,omitempty"`

		WorkConfiguration *struct {
			FeatureGates *[]struct {
				Feature *string `tfsdk:"feature" yaml:"feature,omitempty"`

				Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`
			} `tfsdk:"feature_gates" yaml:"featureGates,omitempty"`
		} `tfsdk:"work_configuration" yaml:"workConfiguration,omitempty"`

		WorkImagePullSpec *string `tfsdk:"work_image_pull_spec" yaml:"workImagePullSpec,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewOperatorOpenClusterManagementIoClusterManagerV1Resource() resource.Resource {
	return &OperatorOpenClusterManagementIoClusterManagerV1Resource{}
}

func (r *OperatorOpenClusterManagementIoClusterManagerV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_operator_open_cluster_management_io_cluster_manager_v1"
}

func (r *OperatorOpenClusterManagementIoClusterManagerV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "ClusterManager configures the controllers on the hub that govern registration and work distribution for attached Klusterlets. In Default mode, ClusterManager will only be deployed in open-cluster-management-hub namespace. In Hosted mode, ClusterManager will be deployed in the namespace with the same name as cluster manager.",
		MarkdownDescription: "ClusterManager configures the controllers on the hub that govern registration and work distribution for attached Klusterlets. In Default mode, ClusterManager will only be deployed in open-cluster-management-hub namespace. In Hosted mode, ClusterManager will be deployed in the namespace with the same name as cluster manager.",
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
				Description:         "Spec represents a desired deployment configuration of controllers that govern registration and work distribution for attached Klusterlets.",
				MarkdownDescription: "Spec represents a desired deployment configuration of controllers that govern registration and work distribution for attached Klusterlets.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"deploy_option": {
						Description:         "DeployOption contains the options of deploying a cluster-manager Default mode is used if DeployOption is not set.",
						MarkdownDescription: "DeployOption contains the options of deploying a cluster-manager Default mode is used if DeployOption is not set.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"hosted": {
								Description:         "Hosted includes configurations we needs for clustermanager in the Hosted mode.",
								MarkdownDescription: "Hosted includes configurations we needs for clustermanager in the Hosted mode.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"registration_webhook_configuration": {
										Description:         "RegistrationWebhookConfiguration represents the customized webhook-server configuration of registration.",
										MarkdownDescription: "RegistrationWebhookConfiguration represents the customized webhook-server configuration of registration.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"address": {
												Description:         "Address represents the address of a webhook-server. It could be in IP format or fqdn format. The Address must be reachable by apiserver of the hub cluster.",
												MarkdownDescription: "Address represents the address of a webhook-server. It could be in IP format or fqdn format. The Address must be reachable by apiserver of the hub cluster.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.RegexMatches(regexp.MustCompile(`^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])$`), ""),
												},
											},

											"port": {
												Description:         "Port represents the port of a webhook-server. The default value of Port is 443.",
												MarkdownDescription: "Port represents the port of a webhook-server. The default value of Port is 443.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtMost(65535),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"work_webhook_configuration": {
										Description:         "WorkWebhookConfiguration represents the customized webhook-server configuration of work.",
										MarkdownDescription: "WorkWebhookConfiguration represents the customized webhook-server configuration of work.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"address": {
												Description:         "Address represents the address of a webhook-server. It could be in IP format or fqdn format. The Address must be reachable by apiserver of the hub cluster.",
												MarkdownDescription: "Address represents the address of a webhook-server. It could be in IP format or fqdn format. The Address must be reachable by apiserver of the hub cluster.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.RegexMatches(regexp.MustCompile(`^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])$`), ""),
												},
											},

											"port": {
												Description:         "Port represents the port of a webhook-server. The default value of Port is 443.",
												MarkdownDescription: "Port represents the port of a webhook-server. The default value of Port is 443.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtMost(65535),
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

							"mode": {
								Description:         "Mode can be Default or Hosted. In Default mode, the Hub is installed as a whole and all parts of Hub are deployed in the same cluster. In Hosted mode, only crd and configurations are installed on one cluster(defined as hub-cluster). Controllers run in another cluster (defined as management-cluster) and connect to the hub with the kubeconfig in secret of 'external-hub-kubeconfig'(a kubeconfig of hub-cluster with cluster-admin permission). Note: Do not modify the Mode field once it's applied.",
								MarkdownDescription: "Mode can be Default or Hosted. In Default mode, the Hub is installed as a whole and all parts of Hub are deployed in the same cluster. In Hosted mode, only crd and configurations are installed on one cluster(defined as hub-cluster). Controllers run in another cluster (defined as management-cluster) and connect to the hub with the kubeconfig in secret of 'external-hub-kubeconfig'(a kubeconfig of hub-cluster with cluster-admin permission). Note: Do not modify the Mode field once it's applied.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("Default", "Hosted"),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"node_placement": {
						Description:         "NodePlacement enables explicit control over the scheduling of the deployed pods.",
						MarkdownDescription: "NodePlacement enables explicit control over the scheduling of the deployed pods.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"node_selector": {
								Description:         "NodeSelector defines which Nodes the Pods are scheduled on. The default is an empty list.",
								MarkdownDescription: "NodeSelector defines which Nodes the Pods are scheduled on. The default is an empty list.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tolerations": {
								Description:         "Tolerations is attached by pods to tolerate any taint that matches the triple <key,value,effect> using the matching operator <operator>. The default is an empty list.",
								MarkdownDescription: "Tolerations is attached by pods to tolerate any taint that matches the triple <key,value,effect> using the matching operator <operator>. The default is an empty list.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"effect": {
										Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
										MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"key": {
										Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
										MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"operator": {
										Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
										MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"toleration_seconds": {
										Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
										MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"value": {
										Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
										MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",

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

					"placement_image_pull_spec": {
						Description:         "PlacementImagePullSpec represents the desired image configuration of placement controller/webhook installed on hub.",
						MarkdownDescription: "PlacementImagePullSpec represents the desired image configuration of placement controller/webhook installed on hub.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"registration_configuration": {
						Description:         "RegistrationConfiguration contains the configuration of registration",
						MarkdownDescription: "RegistrationConfiguration contains the configuration of registration",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"feature_gates": {
								Description:         "FeatureGates represents the list of feature gates for registration If it is set empty, default feature gates will be used. If it is set, featuregate/Foo is an example of one item in FeatureGates:   1. If featuregate/Foo does not exist, registration-operator will discard it   2. If featuregate/Foo exists and is false by default. It is now possible to set featuregate/Foo=[false|true]   3. If featuregate/Foo exists and is true by default. If a cluster-admin upgrading from 1 to 2 wants to continue having featuregate/Foo=false,  	he can set featuregate/Foo=false before upgrading. Let's say the cluster-admin wants featuregate/Foo=false.",
								MarkdownDescription: "FeatureGates represents the list of feature gates for registration If it is set empty, default feature gates will be used. If it is set, featuregate/Foo is an example of one item in FeatureGates:   1. If featuregate/Foo does not exist, registration-operator will discard it   2. If featuregate/Foo exists and is false by default. It is now possible to set featuregate/Foo=[false|true]   3. If featuregate/Foo exists and is true by default. If a cluster-admin upgrading from 1 to 2 wants to continue having featuregate/Foo=false,  	he can set featuregate/Foo=false before upgrading. Let's say the cluster-admin wants featuregate/Foo=false.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"feature": {
										Description:         "Feature is the key of feature gate. e.g. featuregate/Foo.",
										MarkdownDescription: "Feature is the key of feature gate. e.g. featuregate/Foo.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"mode": {
										Description:         "Mode is either Enable, Disable, '' where '' is Disable by default. In Enable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=true'. In Disable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=false'.",
										MarkdownDescription: "Mode is either Enable, Disable, '' where '' is Disable by default. In Enable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=true'. In Disable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=false'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("Enable", "Disable"),
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

					"registration_image_pull_spec": {
						Description:         "RegistrationImagePullSpec represents the desired image of registration controller/webhook installed on hub.",
						MarkdownDescription: "RegistrationImagePullSpec represents the desired image of registration controller/webhook installed on hub.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"work_configuration": {
						Description:         "WorkConfiguration contains the configuration of work",
						MarkdownDescription: "WorkConfiguration contains the configuration of work",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"feature_gates": {
								Description:         "FeatureGates represents the list of feature gates for work If it is set empty, default feature gates will be used. If it is set, featuregate/Foo is an example of one item in FeatureGates:   1. If featuregate/Foo does not exist, registration-operator will discard it   2. If featuregate/Foo exists and is false by default. It is now possible to set featuregate/Foo=[false|true]   3. If featuregate/Foo exists and is true by default. If a cluster-admin upgrading from 1 to 2 wants to continue having featuregate/Foo=false,  	he can set featuregate/Foo=false before upgrading. Let's say the cluster-admin wants featuregate/Foo=false.",
								MarkdownDescription: "FeatureGates represents the list of feature gates for work If it is set empty, default feature gates will be used. If it is set, featuregate/Foo is an example of one item in FeatureGates:   1. If featuregate/Foo does not exist, registration-operator will discard it   2. If featuregate/Foo exists and is false by default. It is now possible to set featuregate/Foo=[false|true]   3. If featuregate/Foo exists and is true by default. If a cluster-admin upgrading from 1 to 2 wants to continue having featuregate/Foo=false,  	he can set featuregate/Foo=false before upgrading. Let's say the cluster-admin wants featuregate/Foo=false.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"feature": {
										Description:         "Feature is the key of feature gate. e.g. featuregate/Foo.",
										MarkdownDescription: "Feature is the key of feature gate. e.g. featuregate/Foo.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"mode": {
										Description:         "Mode is either Enable, Disable, '' where '' is Disable by default. In Enable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=true'. In Disable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=false'.",
										MarkdownDescription: "Mode is either Enable, Disable, '' where '' is Disable by default. In Enable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=true'. In Disable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=false'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("Enable", "Disable"),
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

					"work_image_pull_spec": {
						Description:         "WorkImagePullSpec represents the desired image configuration of work controller/webhook installed on hub.",
						MarkdownDescription: "WorkImagePullSpec represents the desired image configuration of work controller/webhook installed on hub.",

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

func (r *OperatorOpenClusterManagementIoClusterManagerV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_operator_open_cluster_management_io_cluster_manager_v1")

	var state OperatorOpenClusterManagementIoClusterManagerV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel OperatorOpenClusterManagementIoClusterManagerV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("operator.open-cluster-management.io/v1")
	goModel.Kind = utilities.Ptr("ClusterManager")

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

func (r *OperatorOpenClusterManagementIoClusterManagerV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_operator_open_cluster_management_io_cluster_manager_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *OperatorOpenClusterManagementIoClusterManagerV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_operator_open_cluster_management_io_cluster_manager_v1")

	var state OperatorOpenClusterManagementIoClusterManagerV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel OperatorOpenClusterManagementIoClusterManagerV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("operator.open-cluster-management.io/v1")
	goModel.Kind = utilities.Ptr("ClusterManager")

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

func (r *OperatorOpenClusterManagementIoClusterManagerV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_operator_open_cluster_management_io_cluster_manager_v1")
	// NO-OP: Terraform removes the state automatically for us
}
