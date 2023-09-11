/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package elbv2_k8s_aws_v1beta1

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
	_ datasource.DataSource              = &Elbv2K8SAwsTargetGroupBindingV1Beta1DataSource{}
	_ datasource.DataSourceWithConfigure = &Elbv2K8SAwsTargetGroupBindingV1Beta1DataSource{}
)

func NewElbv2K8SAwsTargetGroupBindingV1Beta1DataSource() datasource.DataSource {
	return &Elbv2K8SAwsTargetGroupBindingV1Beta1DataSource{}
}

type Elbv2K8SAwsTargetGroupBindingV1Beta1DataSource struct {
	kubernetesClient dynamic.Interface
}

type Elbv2K8SAwsTargetGroupBindingV1Beta1DataSourceData struct {
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
		IpAddressType *string `tfsdk:"ip_address_type" json:"ipAddressType,omitempty"`
		Networking    *struct {
			Ingress *[]struct {
				From *[]struct {
					IpBlock *struct {
						Cidr *string `tfsdk:"cidr" json:"cidr,omitempty"`
					} `tfsdk:"ip_block" json:"ipBlock,omitempty"`
					SecurityGroup *struct {
						GroupID *string `tfsdk:"group_id" json:"groupID,omitempty"`
					} `tfsdk:"security_group" json:"securityGroup,omitempty"`
				} `tfsdk:"from" json:"from,omitempty"`
				Ports *[]struct {
					Port     *string `tfsdk:"port" json:"port,omitempty"`
					Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
				} `tfsdk:"ports" json:"ports,omitempty"`
			} `tfsdk:"ingress" json:"ingress,omitempty"`
		} `tfsdk:"networking" json:"networking,omitempty"`
		NodeSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
		ServiceRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
			Port *string `tfsdk:"port" json:"port,omitempty"`
		} `tfsdk:"service_ref" json:"serviceRef,omitempty"`
		TargetGroupARN *string `tfsdk:"target_group_arn" json:"targetGroupARN,omitempty"`
		TargetType     *string `tfsdk:"target_type" json:"targetType,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *Elbv2K8SAwsTargetGroupBindingV1Beta1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_elbv2_k8s_aws_target_group_binding_v1beta1"
}

func (r *Elbv2K8SAwsTargetGroupBindingV1Beta1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "TargetGroupBinding is the Schema for the TargetGroupBinding API",
		MarkdownDescription: "TargetGroupBinding is the Schema for the TargetGroupBinding API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
				Description:         "TargetGroupBindingSpec defines the desired state of TargetGroupBinding",
				MarkdownDescription: "TargetGroupBindingSpec defines the desired state of TargetGroupBinding",
				Attributes: map[string]schema.Attribute{
					"ip_address_type": schema.StringAttribute{
						Description:         "ipAddressType specifies whether the target group is of type IPv4 or IPv6. If unspecified, it will be automatically inferred.",
						MarkdownDescription: "ipAddressType specifies whether the target group is of type IPv4 or IPv6. If unspecified, it will be automatically inferred.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"networking": schema.SingleNestedAttribute{
						Description:         "networking defines the networking rules to allow ELBV2 LoadBalancer to access targets in TargetGroup.",
						MarkdownDescription: "networking defines the networking rules to allow ELBV2 LoadBalancer to access targets in TargetGroup.",
						Attributes: map[string]schema.Attribute{
							"ingress": schema.ListNestedAttribute{
								Description:         "List of ingress rules to allow ELBV2 LoadBalancer to access targets in TargetGroup.",
								MarkdownDescription: "List of ingress rules to allow ELBV2 LoadBalancer to access targets in TargetGroup.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"from": schema.ListNestedAttribute{
											Description:         "List of peers which should be able to access the targets in TargetGroup. At least one NetworkingPeer should be specified.",
											MarkdownDescription: "List of peers which should be able to access the targets in TargetGroup. At least one NetworkingPeer should be specified.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"ip_block": schema.SingleNestedAttribute{
														Description:         "IPBlock defines an IPBlock peer. If specified, none of the other fields can be set.",
														MarkdownDescription: "IPBlock defines an IPBlock peer. If specified, none of the other fields can be set.",
														Attributes: map[string]schema.Attribute{
															"cidr": schema.StringAttribute{
																Description:         "CIDR is the network CIDR. Both IPV4 or IPV6 CIDR are accepted.",
																MarkdownDescription: "CIDR is the network CIDR. Both IPV4 or IPV6 CIDR are accepted.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"security_group": schema.SingleNestedAttribute{
														Description:         "SecurityGroup defines a SecurityGroup peer. If specified, none of the other fields can be set.",
														MarkdownDescription: "SecurityGroup defines a SecurityGroup peer. If specified, none of the other fields can be set.",
														Attributes: map[string]schema.Attribute{
															"group_id": schema.StringAttribute{
																Description:         "GroupID is the EC2 SecurityGroupID.",
																MarkdownDescription: "GroupID is the EC2 SecurityGroupID.",
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
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"ports": schema.ListNestedAttribute{
											Description:         "List of ports which should be made accessible on the targets in TargetGroup. If ports is empty or unspecified, it defaults to all ports with TCP.",
											MarkdownDescription: "List of ports which should be made accessible on the targets in TargetGroup. If ports is empty or unspecified, it defaults to all ports with TCP.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"port": schema.StringAttribute{
														Description:         "The port which traffic must match. When NodePort endpoints(instance TargetType) is used, this must be a numerical port. When Port endpoints(ip TargetType) is used, this can be either numerical or named port on pods. if port is unspecified, it defaults to all ports.",
														MarkdownDescription: "The port which traffic must match. When NodePort endpoints(instance TargetType) is used, this must be a numerical port. When Port endpoints(ip TargetType) is used, this can be either numerical or named port on pods. if port is unspecified, it defaults to all ports.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"protocol": schema.StringAttribute{
														Description:         "The protocol which traffic must match. If protocol is unspecified, it defaults to TCP.",
														MarkdownDescription: "The protocol which traffic must match. If protocol is unspecified, it defaults to TCP.",
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

					"node_selector": schema.SingleNestedAttribute{
						Description:         "node selector for instance type target groups to only register certain nodes",
						MarkdownDescription: "node selector for instance type target groups to only register certain nodes",
						Attributes: map[string]schema.Attribute{
							"match_expressions": schema.ListNestedAttribute{
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "key is the label key that the selector applies to.",
											MarkdownDescription: "key is the label key that the selector applies to.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"operator": schema.StringAttribute{
											Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											ElementType:         types.StringType,
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

							"match_labels": schema.MapAttribute{
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

					"service_ref": schema.SingleNestedAttribute{
						Description:         "serviceRef is a reference to a Kubernetes Service and ServicePort.",
						MarkdownDescription: "serviceRef is a reference to a Kubernetes Service and ServicePort.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name is the name of the Service.",
								MarkdownDescription: "Name is the name of the Service.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"port": schema.StringAttribute{
								Description:         "Port is the port of the ServicePort.",
								MarkdownDescription: "Port is the port of the ServicePort.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"target_group_arn": schema.StringAttribute{
						Description:         "targetGroupARN is the Amazon Resource Name (ARN) for the TargetGroup.",
						MarkdownDescription: "targetGroupARN is the Amazon Resource Name (ARN) for the TargetGroup.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"target_type": schema.StringAttribute{
						Description:         "targetType is the TargetType of TargetGroup. If unspecified, it will be automatically inferred.",
						MarkdownDescription: "targetType is the TargetType of TargetGroup. If unspecified, it will be automatically inferred.",
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

func (r *Elbv2K8SAwsTargetGroupBindingV1Beta1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *Elbv2K8SAwsTargetGroupBindingV1Beta1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_elbv2_k8s_aws_target_group_binding_v1beta1")

	var data Elbv2K8SAwsTargetGroupBindingV1Beta1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "elbv2.k8s.aws", Version: "v1beta1", Resource: "targetgroupbindings"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse Elbv2K8SAwsTargetGroupBindingV1Beta1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("elbv2.k8s.aws/v1beta1")
	data.Kind = pointer.String("TargetGroupBinding")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
