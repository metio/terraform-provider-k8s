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
	_ datasource.DataSource              = &Elbv2K8SAwsIngressClassParamsV1Beta1DataSource{}
	_ datasource.DataSourceWithConfigure = &Elbv2K8SAwsIngressClassParamsV1Beta1DataSource{}
)

func NewElbv2K8SAwsIngressClassParamsV1Beta1DataSource() datasource.DataSource {
	return &Elbv2K8SAwsIngressClassParamsV1Beta1DataSource{}
}

type Elbv2K8SAwsIngressClassParamsV1Beta1DataSource struct {
	kubernetesClient dynamic.Interface
}

type Elbv2K8SAwsIngressClassParamsV1Beta1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Group *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"group" json:"group,omitempty"`
		InboundCIDRs           *[]string `tfsdk:"inbound_cid_rs" json:"inboundCIDRs,omitempty"`
		IpAddressType          *string   `tfsdk:"ip_address_type" json:"ipAddressType,omitempty"`
		LoadBalancerAttributes *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"load_balancer_attributes" json:"loadBalancerAttributes,omitempty"`
		NamespaceSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
		Scheme    *string `tfsdk:"scheme" json:"scheme,omitempty"`
		SslPolicy *string `tfsdk:"ssl_policy" json:"sslPolicy,omitempty"`
		Subnets   *struct {
			Ids  *[]string            `tfsdk:"ids" json:"ids,omitempty"`
			Tags *map[string][]string `tfsdk:"tags" json:"tags,omitempty"`
		} `tfsdk:"subnets" json:"subnets,omitempty"`
		Tags *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *Elbv2K8SAwsIngressClassParamsV1Beta1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_elbv2_k8s_aws_ingress_class_params_v1beta1"
}

func (r *Elbv2K8SAwsIngressClassParamsV1Beta1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "IngressClassParams is the Schema for the IngressClassParams API",
		MarkdownDescription: "IngressClassParams is the Schema for the IngressClassParams API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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
				Description:         "IngressClassParamsSpec defines the desired state of IngressClassParams",
				MarkdownDescription: "IngressClassParamsSpec defines the desired state of IngressClassParams",
				Attributes: map[string]schema.Attribute{
					"group": schema.SingleNestedAttribute{
						Description:         "Group defines the IngressGroup for all Ingresses that belong to IngressClass with this IngressClassParams.",
						MarkdownDescription: "Group defines the IngressGroup for all Ingresses that belong to IngressClass with this IngressClassParams.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name is the name of IngressGroup.",
								MarkdownDescription: "Name is the name of IngressGroup.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"inbound_cid_rs": schema.ListAttribute{
						Description:         "InboundCIDRs specifies the CIDRs that are allowed to access the Ingresses that belong to IngressClass with this IngressClassParams.",
						MarkdownDescription: "InboundCIDRs specifies the CIDRs that are allowed to access the Ingresses that belong to IngressClass with this IngressClassParams.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"ip_address_type": schema.StringAttribute{
						Description:         "IPAddressType defines the ip address type for all Ingresses that belong to IngressClass with this IngressClassParams.",
						MarkdownDescription: "IPAddressType defines the ip address type for all Ingresses that belong to IngressClass with this IngressClassParams.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"load_balancer_attributes": schema.ListNestedAttribute{
						Description:         "LoadBalancerAttributes define the custom attributes to LoadBalancers for all Ingress that that belong to IngressClass with this IngressClassParams.",
						MarkdownDescription: "LoadBalancerAttributes define the custom attributes to LoadBalancers for all Ingress that that belong to IngressClass with this IngressClassParams.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Description:         "The key of the attribute.",
									MarkdownDescription: "The key of the attribute.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"value": schema.StringAttribute{
									Description:         "The value of the attribute.",
									MarkdownDescription: "The value of the attribute.",
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

					"namespace_selector": schema.SingleNestedAttribute{
						Description:         "NamespaceSelector restrict the namespaces of Ingresses that are allowed to specify the IngressClass with this IngressClassParams. * if absent or present but empty, it selects all namespaces.",
						MarkdownDescription: "NamespaceSelector restrict the namespaces of Ingresses that are allowed to specify the IngressClass with this IngressClassParams. * if absent or present but empty, it selects all namespaces.",
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

					"scheme": schema.StringAttribute{
						Description:         "Scheme defines the scheme for all Ingresses that belong to IngressClass with this IngressClassParams.",
						MarkdownDescription: "Scheme defines the scheme for all Ingresses that belong to IngressClass with this IngressClassParams.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"ssl_policy": schema.StringAttribute{
						Description:         "SSLPolicy specifies the SSL Policy for all Ingresses that belong to IngressClass with this IngressClassParams.",
						MarkdownDescription: "SSLPolicy specifies the SSL Policy for all Ingresses that belong to IngressClass with this IngressClassParams.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"subnets": schema.SingleNestedAttribute{
						Description:         "Subnets defines the subnets for all Ingresses that belong to IngressClass with this IngressClassParams.",
						MarkdownDescription: "Subnets defines the subnets for all Ingresses that belong to IngressClass with this IngressClassParams.",
						Attributes: map[string]schema.Attribute{
							"ids": schema.ListAttribute{
								Description:         "IDs specify the resource IDs of subnets. Exactly one of this or 'tags' must be specified.",
								MarkdownDescription: "IDs specify the resource IDs of subnets. Exactly one of this or 'tags' must be specified.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"tags": schema.MapAttribute{
								Description:         "Tags specifies subnets in the load balancer's VPC where each tag specified in the map key contains one of the values in the corresponding value list. Exactly one of this or 'ids' must be specified.",
								MarkdownDescription: "Tags specifies subnets in the load balancer's VPC where each tag specified in the map key contains one of the values in the corresponding value list. Exactly one of this or 'ids' must be specified.",
								ElementType:         types.ListType{ElemType: types.StringType},
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "Tags defines list of Tags on AWS resources provisioned for Ingresses that belong to IngressClass with this IngressClassParams.",
						MarkdownDescription: "Tags defines list of Tags on AWS resources provisioned for Ingresses that belong to IngressClass with this IngressClassParams.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Description:         "The key of the tag.",
									MarkdownDescription: "The key of the tag.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"value": schema.StringAttribute{
									Description:         "The value of the tag.",
									MarkdownDescription: "The value of the tag.",
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
		},
	}
}

func (r *Elbv2K8SAwsIngressClassParamsV1Beta1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *Elbv2K8SAwsIngressClassParamsV1Beta1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_elbv2_k8s_aws_ingress_class_params_v1beta1")

	var data Elbv2K8SAwsIngressClassParamsV1Beta1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "elbv2.k8s.aws", Version: "v1beta1", Resource: "IngressClassParams"}).
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

	var readResponse Elbv2K8SAwsIngressClassParamsV1Beta1DataSourceData
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

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("elbv2.k8s.aws/v1beta1")
	data.Kind = pointer.String("IngressClassParams")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
