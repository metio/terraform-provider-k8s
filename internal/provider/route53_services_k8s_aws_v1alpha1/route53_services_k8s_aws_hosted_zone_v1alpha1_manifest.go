/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package route53_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource = &Route53ServicesK8SAwsHostedZoneV1Alpha1Manifest{}
)

func NewRoute53ServicesK8SAwsHostedZoneV1Alpha1Manifest() datasource.DataSource {
	return &Route53ServicesK8SAwsHostedZoneV1Alpha1Manifest{}
}

type Route53ServicesK8SAwsHostedZoneV1Alpha1Manifest struct{}

type Route53ServicesK8SAwsHostedZoneV1Alpha1ManifestData struct {
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
		DelegationSetID  *string `tfsdk:"delegation_set_id" json:"delegationSetID,omitempty"`
		HostedZoneConfig *struct {
			Comment     *string `tfsdk:"comment" json:"comment,omitempty"`
			PrivateZone *bool   `tfsdk:"private_zone" json:"privateZone,omitempty"`
		} `tfsdk:"hosted_zone_config" json:"hostedZoneConfig,omitempty"`
		Name *string `tfsdk:"name" json:"name,omitempty"`
		Tags *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
		Vpc *struct {
			VpcID     *string `tfsdk:"vpc_id" json:"vpcID,omitempty"`
			VpcRegion *string `tfsdk:"vpc_region" json:"vpcRegion,omitempty"`
		} `tfsdk:"vpc" json:"vpc,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *Route53ServicesK8SAwsHostedZoneV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_route53_services_k8s_aws_hosted_zone_v1alpha1_manifest"
}

func (r *Route53ServicesK8SAwsHostedZoneV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "HostedZone is the Schema for the HostedZones API",
		MarkdownDescription: "HostedZone is the Schema for the HostedZones API",
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
				Description:         "HostedZoneSpec defines the desired state of HostedZone.A complex type that contains general information about the hosted zone.",
				MarkdownDescription: "HostedZoneSpec defines the desired state of HostedZone.A complex type that contains general information about the hosted zone.",
				Attributes: map[string]schema.Attribute{
					"delegation_set_id": schema.StringAttribute{
						Description:         "If you want to associate a reusable delegation set with this hosted zone,the ID that Amazon Route 53 assigned to the reusable delegation set whenyou created it. For more information about reusable delegation sets, seeCreateReusableDelegationSet (https://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateReusableDelegationSet.html).If you are using a reusable delegation set to create a public hosted zonefor a subdomain, make sure that the parent hosted zone doesn't use one ormore of the same name servers. If you have overlapping nameservers, the operationwill cause a ConflictingDomainsExist error.",
						MarkdownDescription: "If you want to associate a reusable delegation set with this hosted zone,the ID that Amazon Route 53 assigned to the reusable delegation set whenyou created it. For more information about reusable delegation sets, seeCreateReusableDelegationSet (https://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateReusableDelegationSet.html).If you are using a reusable delegation set to create a public hosted zonefor a subdomain, make sure that the parent hosted zone doesn't use one ormore of the same name servers. If you have overlapping nameservers, the operationwill cause a ConflictingDomainsExist error.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"hosted_zone_config": schema.SingleNestedAttribute{
						Description:         "(Optional) A complex type that contains the following optional values:   * For public and private hosted zones, an optional comment   * For private hosted zones, an optional PrivateZone elementIf you don't specify a comment or the PrivateZone element, omit HostedZoneConfigand the other elements.",
						MarkdownDescription: "(Optional) A complex type that contains the following optional values:   * For public and private hosted zones, an optional comment   * For private hosted zones, an optional PrivateZone elementIf you don't specify a comment or the PrivateZone element, omit HostedZoneConfigand the other elements.",
						Attributes: map[string]schema.Attribute{
							"comment": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"private_zone": schema.BoolAttribute{
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

					"name": schema.StringAttribute{
						Description:         "The name of the domain. Specify a fully qualified domain name, for example,www.example.com. The trailing dot is optional; Amazon Route 53 assumes thatthe domain name is fully qualified. This means that Route 53 treats www.example.com(without a trailing dot) and www.example.com. (with a trailing dot) as identical.If you're creating a public hosted zone, this is the name you have registeredwith your DNS registrar. If your domain name is registered with a registrarother than Route 53, change the name servers for your domain to the set ofNameServers that CreateHostedZone returns in DelegationSet.",
						MarkdownDescription: "The name of the domain. Specify a fully qualified domain name, for example,www.example.com. The trailing dot is optional; Amazon Route 53 assumes thatthe domain name is fully qualified. This means that Route 53 treats www.example.com(without a trailing dot) and www.example.com. (with a trailing dot) as identical.If you're creating a public hosted zone, this is the name you have registeredwith your DNS registrar. If your domain name is registered with a registrarother than Route 53, change the name servers for your domain to the set ofNameServers that CreateHostedZone returns in DelegationSet.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "A complex type that contains a list of the tags that you want to add to thespecified health check or hosted zone and/or the tags that you want to editValue for.You can add a maximum of 10 tags to a health check or a hosted zone.",
						MarkdownDescription: "A complex type that contains a list of the tags that you want to add to thespecified health check or hosted zone and/or the tags that you want to editValue for.You can add a maximum of 10 tags to a health check or a hosted zone.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"value": schema.StringAttribute{
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

					"vpc": schema.SingleNestedAttribute{
						Description:         "(Private hosted zones only) A complex type that contains information aboutthe Amazon VPC that you're associating with this hosted zone.You can specify only one Amazon VPC when you create a private hosted zone.If you are associating a VPC with a hosted zone with this request, the paramatersVPCId and VPCRegion are also required.To associate additional Amazon VPCs with the hosted zone, use AssociateVPCWithHostedZone(https://docs.aws.amazon.com/Route53/latest/APIReference/API_AssociateVPCWithHostedZone.html)after you create a hosted zone.",
						MarkdownDescription: "(Private hosted zones only) A complex type that contains information aboutthe Amazon VPC that you're associating with this hosted zone.You can specify only one Amazon VPC when you create a private hosted zone.If you are associating a VPC with a hosted zone with this request, the paramatersVPCId and VPCRegion are also required.To associate additional Amazon VPCs with the hosted zone, use AssociateVPCWithHostedZone(https://docs.aws.amazon.com/Route53/latest/APIReference/API_AssociateVPCWithHostedZone.html)after you create a hosted zone.",
						Attributes: map[string]schema.Attribute{
							"vpc_id": schema.StringAttribute{
								Description:         "(Private hosted zones only) The ID of an Amazon VPC.",
								MarkdownDescription: "(Private hosted zones only) The ID of an Amazon VPC.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"vpc_region": schema.StringAttribute{
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *Route53ServicesK8SAwsHostedZoneV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_route53_services_k8s_aws_hosted_zone_v1alpha1_manifest")

	var model Route53ServicesK8SAwsHostedZoneV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("route53.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("HostedZone")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
