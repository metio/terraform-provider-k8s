/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package elbv2_k8s_aws_v1alpha1

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
	_ datasource.DataSource = &Elbv2K8SAwsTargetGroupBindingV1Alpha1Manifest{}
)

func NewElbv2K8SAwsTargetGroupBindingV1Alpha1Manifest() datasource.DataSource {
	return &Elbv2K8SAwsTargetGroupBindingV1Alpha1Manifest{}
}

type Elbv2K8SAwsTargetGroupBindingV1Alpha1Manifest struct{}

type Elbv2K8SAwsTargetGroupBindingV1Alpha1ManifestData struct {
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
		Networking *struct {
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
		ServiceRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
			Port *string `tfsdk:"port" json:"port,omitempty"`
		} `tfsdk:"service_ref" json:"serviceRef,omitempty"`
		TargetGroupARN *string `tfsdk:"target_group_arn" json:"targetGroupARN,omitempty"`
		TargetType     *string `tfsdk:"target_type" json:"targetType,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *Elbv2K8SAwsTargetGroupBindingV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_elbv2_k8s_aws_target_group_binding_v1alpha1_manifest"
}

func (r *Elbv2K8SAwsTargetGroupBindingV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "TargetGroupBinding is the Schema for the TargetGroupBinding API",
		MarkdownDescription: "TargetGroupBinding is the Schema for the TargetGroupBinding API",
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
				Description:         "TargetGroupBindingSpec defines the desired state of TargetGroupBinding",
				MarkdownDescription: "TargetGroupBindingSpec defines the desired state of TargetGroupBinding",
				Attributes: map[string]schema.Attribute{
					"networking": schema.SingleNestedAttribute{
						Description:         "networking provides the networking setup for ELBV2 LoadBalancer to access targets in TargetGroup.",
						MarkdownDescription: "networking provides the networking setup for ELBV2 LoadBalancer to access targets in TargetGroup.",
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
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"security_group": schema.SingleNestedAttribute{
														Description:         "SecurityGroup defines a SecurityGroup peer. If specified, none of the other fields can be set.",
														MarkdownDescription: "SecurityGroup defines a SecurityGroup peer. If specified, none of the other fields can be set.",
														Attributes: map[string]schema.Attribute{
															"group_id": schema.StringAttribute{
																Description:         "GroupID is the EC2 SecurityGroupID.",
																MarkdownDescription: "GroupID is the EC2 SecurityGroupID.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
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
														Optional:            true,
														Computed:            false,
													},

													"protocol": schema.StringAttribute{
														Description:         "The protocol which traffic must match. If protocol is unspecified, it defaults to TCP.",
														MarkdownDescription: "The protocol which traffic must match. If protocol is unspecified, it defaults to TCP.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("TCP", "UDP"),
														},
													},
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
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

					"service_ref": schema.SingleNestedAttribute{
						Description:         "serviceRef is a reference to a Kubernetes Service and ServicePort.",
						MarkdownDescription: "serviceRef is a reference to a Kubernetes Service and ServicePort.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name is the name of the Service.",
								MarkdownDescription: "Name is the name of the Service.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"port": schema.StringAttribute{
								Description:         "Port is the port of the ServicePort.",
								MarkdownDescription: "Port is the port of the ServicePort.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"target_group_arn": schema.StringAttribute{
						Description:         "targetGroupARN is the Amazon Resource Name (ARN) for the TargetGroup.",
						MarkdownDescription: "targetGroupARN is the Amazon Resource Name (ARN) for the TargetGroup.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"target_type": schema.StringAttribute{
						Description:         "targetType is the TargetType of TargetGroup. If unspecified, it will be automatically inferred.",
						MarkdownDescription: "targetType is the TargetType of TargetGroup. If unspecified, it will be automatically inferred.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("instance", "ip"),
						},
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *Elbv2K8SAwsTargetGroupBindingV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_elbv2_k8s_aws_target_group_binding_v1alpha1_manifest")

	var model Elbv2K8SAwsTargetGroupBindingV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("elbv2.k8s.aws/v1alpha1")
	model.Kind = pointer.String("TargetGroupBinding")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
