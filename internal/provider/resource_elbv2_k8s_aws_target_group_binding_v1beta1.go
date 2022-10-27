/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

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

type Elbv2K8SAwsTargetGroupBindingV1Beta1Resource struct{}

var (
	_ resource.Resource = (*Elbv2K8SAwsTargetGroupBindingV1Beta1Resource)(nil)
)

type Elbv2K8SAwsTargetGroupBindingV1Beta1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type Elbv2K8SAwsTargetGroupBindingV1Beta1GoModel struct {
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
		IpAddressType *string `tfsdk:"ip_address_type" yaml:"ipAddressType,omitempty"`

		Networking *struct {
			Ingress *[]struct {
				From *[]struct {
					IpBlock *struct {
						Cidr *string `tfsdk:"cidr" yaml:"cidr,omitempty"`
					} `tfsdk:"ip_block" yaml:"ipBlock,omitempty"`

					SecurityGroup *struct {
						GroupID *string `tfsdk:"group_id" yaml:"groupID,omitempty"`
					} `tfsdk:"security_group" yaml:"securityGroup,omitempty"`
				} `tfsdk:"from" yaml:"from,omitempty"`

				Ports *[]struct {
					Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

					Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`
				} `tfsdk:"ports" yaml:"ports,omitempty"`
			} `tfsdk:"ingress" yaml:"ingress,omitempty"`
		} `tfsdk:"networking" yaml:"networking,omitempty"`

		NodeSelector *struct {
			MatchExpressions *[]struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

				Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
			} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

			MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
		} `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

		ServiceRef *struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`
		} `tfsdk:"service_ref" yaml:"serviceRef,omitempty"`

		TargetGroupARN *string `tfsdk:"target_group_arn" yaml:"targetGroupARN,omitempty"`

		TargetType *string `tfsdk:"target_type" yaml:"targetType,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewElbv2K8SAwsTargetGroupBindingV1Beta1Resource() resource.Resource {
	return &Elbv2K8SAwsTargetGroupBindingV1Beta1Resource{}
}

func (r *Elbv2K8SAwsTargetGroupBindingV1Beta1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_elbv2_k8s_aws_target_group_binding_v1beta1"
}

func (r *Elbv2K8SAwsTargetGroupBindingV1Beta1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "TargetGroupBinding is the Schema for the TargetGroupBinding API",
		MarkdownDescription: "TargetGroupBinding is the Schema for the TargetGroupBinding API",
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
				Description:         "TargetGroupBindingSpec defines the desired state of TargetGroupBinding",
				MarkdownDescription: "TargetGroupBindingSpec defines the desired state of TargetGroupBinding",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"ip_address_type": {
						Description:         "ipAddressType specifies whether the target group is of type IPv4 or IPv6. If unspecified, it will be automatically inferred.",
						MarkdownDescription: "ipAddressType specifies whether the target group is of type IPv4 or IPv6. If unspecified, it will be automatically inferred.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("ipv4", "ipv6"),
						},
					},

					"networking": {
						Description:         "networking defines the networking rules to allow ELBV2 LoadBalancer to access targets in TargetGroup.",
						MarkdownDescription: "networking defines the networking rules to allow ELBV2 LoadBalancer to access targets in TargetGroup.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"ingress": {
								Description:         "List of ingress rules to allow ELBV2 LoadBalancer to access targets in TargetGroup.",
								MarkdownDescription: "List of ingress rules to allow ELBV2 LoadBalancer to access targets in TargetGroup.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"from": {
										Description:         "List of peers which should be able to access the targets in TargetGroup. At least one NetworkingPeer should be specified.",
										MarkdownDescription: "List of peers which should be able to access the targets in TargetGroup. At least one NetworkingPeer should be specified.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"ip_block": {
												Description:         "IPBlock defines an IPBlock peer. If specified, none of the other fields can be set.",
												MarkdownDescription: "IPBlock defines an IPBlock peer. If specified, none of the other fields can be set.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"cidr": {
														Description:         "CIDR is the network CIDR. Both IPV4 or IPV6 CIDR are accepted.",
														MarkdownDescription: "CIDR is the network CIDR. Both IPV4 or IPV6 CIDR are accepted.",

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

											"security_group": {
												Description:         "SecurityGroup defines a SecurityGroup peer. If specified, none of the other fields can be set.",
												MarkdownDescription: "SecurityGroup defines a SecurityGroup peer. If specified, none of the other fields can be set.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"group_id": {
														Description:         "GroupID is the EC2 SecurityGroupID.",
														MarkdownDescription: "GroupID is the EC2 SecurityGroupID.",

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
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},

									"ports": {
										Description:         "List of ports which should be made accessible on the targets in TargetGroup. If ports is empty or unspecified, it defaults to all ports with TCP.",
										MarkdownDescription: "List of ports which should be made accessible on the targets in TargetGroup. If ports is empty or unspecified, it defaults to all ports with TCP.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"port": {
												Description:         "The port which traffic must match. When NodePort endpoints(instance TargetType) is used, this must be a numerical port. When Port endpoints(ip TargetType) is used, this can be either numerical or named port on pods. if port is unspecified, it defaults to all ports.",
												MarkdownDescription: "The port which traffic must match. When NodePort endpoints(instance TargetType) is used, this must be a numerical port. When Port endpoints(ip TargetType) is used, this can be either numerical or named port on pods. if port is unspecified, it defaults to all ports.",

												Type: utilities.IntOrStringType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"protocol": {
												Description:         "The protocol which traffic must match. If protocol is unspecified, it defaults to TCP.",
												MarkdownDescription: "The protocol which traffic must match. If protocol is unspecified, it defaults to TCP.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("TCP", "UDP"),
												},
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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"node_selector": {
						Description:         "node selector for instance type target groups to only register certain nodes",
						MarkdownDescription: "node selector for instance type target groups to only register certain nodes",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"match_expressions": {
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"key": {
										Description:         "key is the label key that the selector applies to.",
										MarkdownDescription: "key is the label key that the selector applies to.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"operator": {
										Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
										MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"values": {
										Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
										MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

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

							"match_labels": {
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

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

					"service_ref": {
						Description:         "serviceRef is a reference to a Kubernetes Service and ServicePort.",
						MarkdownDescription: "serviceRef is a reference to a Kubernetes Service and ServicePort.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name is the name of the Service.",
								MarkdownDescription: "Name is the name of the Service.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"port": {
								Description:         "Port is the port of the ServicePort.",
								MarkdownDescription: "Port is the port of the ServicePort.",

								Type: utilities.IntOrStringType{},

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"target_group_arn": {
						Description:         "targetGroupARN is the Amazon Resource Name (ARN) for the TargetGroup.",
						MarkdownDescription: "targetGroupARN is the Amazon Resource Name (ARN) for the TargetGroup.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.LengthAtLeast(1),
						},
					},

					"target_type": {
						Description:         "targetType is the TargetType of TargetGroup. If unspecified, it will be automatically inferred.",
						MarkdownDescription: "targetType is the TargetType of TargetGroup. If unspecified, it will be automatically inferred.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("instance", "ip"),
						},
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *Elbv2K8SAwsTargetGroupBindingV1Beta1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_elbv2_k8s_aws_target_group_binding_v1beta1")

	var state Elbv2K8SAwsTargetGroupBindingV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel Elbv2K8SAwsTargetGroupBindingV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("elbv2.k8s.aws/v1beta1")
	goModel.Kind = utilities.Ptr("TargetGroupBinding")

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

func (r *Elbv2K8SAwsTargetGroupBindingV1Beta1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_elbv2_k8s_aws_target_group_binding_v1beta1")
	// NO-OP: All data is already in Terraform state
}

func (r *Elbv2K8SAwsTargetGroupBindingV1Beta1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_elbv2_k8s_aws_target_group_binding_v1beta1")

	var state Elbv2K8SAwsTargetGroupBindingV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel Elbv2K8SAwsTargetGroupBindingV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("elbv2.k8s.aws/v1beta1")
	goModel.Kind = utilities.Ptr("TargetGroupBinding")

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

func (r *Elbv2K8SAwsTargetGroupBindingV1Beta1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_elbv2_k8s_aws_target_group_binding_v1beta1")
	// NO-OP: Terraform removes the state automatically for us
}
