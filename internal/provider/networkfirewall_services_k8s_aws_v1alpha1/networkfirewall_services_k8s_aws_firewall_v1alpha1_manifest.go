/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package networkfirewall_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource = &NetworkfirewallServicesK8SAwsFirewallV1Alpha1Manifest{}
)

func NewNetworkfirewallServicesK8SAwsFirewallV1Alpha1Manifest() datasource.DataSource {
	return &NetworkfirewallServicesK8SAwsFirewallV1Alpha1Manifest{}
}

type NetworkfirewallServicesK8SAwsFirewallV1Alpha1Manifest struct{}

type NetworkfirewallServicesK8SAwsFirewallV1Alpha1ManifestData struct {
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
		DeleteProtection        *bool   `tfsdk:"delete_protection" json:"deleteProtection,omitempty"`
		Description             *string `tfsdk:"description" json:"description,omitempty"`
		EncryptionConfiguration *struct {
			KeyID *string `tfsdk:"key_id" json:"keyID,omitempty"`
			Type_ *string `tfsdk:"type_" json:"type_,omitempty"`
		} `tfsdk:"encryption_configuration" json:"encryptionConfiguration,omitempty"`
		FirewallName                   *string `tfsdk:"firewall_name" json:"firewallName,omitempty"`
		FirewallPolicyARN              *string `tfsdk:"firewall_policy_arn" json:"firewallPolicyARN,omitempty"`
		FirewallPolicyChangeProtection *bool   `tfsdk:"firewall_policy_change_protection" json:"firewallPolicyChangeProtection,omitempty"`
		SubnetChangeProtection         *bool   `tfsdk:"subnet_change_protection" json:"subnetChangeProtection,omitempty"`
		SubnetMappings                 *[]struct {
			IpAddressType *string `tfsdk:"ip_address_type" json:"ipAddressType,omitempty"`
			SubnetID      *string `tfsdk:"subnet_id" json:"subnetID,omitempty"`
		} `tfsdk:"subnet_mappings" json:"subnetMappings,omitempty"`
		Tags *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
		VpcID *string `tfsdk:"vpc_id" json:"vpcID,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *NetworkfirewallServicesK8SAwsFirewallV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_networkfirewall_services_k8s_aws_firewall_v1alpha1_manifest"
}

func (r *NetworkfirewallServicesK8SAwsFirewallV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Firewall is the Schema for the Firewalls API",
		MarkdownDescription: "Firewall is the Schema for the Firewalls API",
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
				Description:         "FirewallSpec defines the desired state of Firewall. The firewall defines the configuration settings for an Network Firewall firewall. These settings include the firewall policy, the subnets in your VPC to use for the firewall endpoints, and any tags that are attached to the firewall Amazon Web Services resource. The status of the firewall, for example whether it's ready to filter network traffic, is provided in the corresponding FirewallStatus. You can retrieve both objects by calling DescribeFirewall.",
				MarkdownDescription: "FirewallSpec defines the desired state of Firewall. The firewall defines the configuration settings for an Network Firewall firewall. These settings include the firewall policy, the subnets in your VPC to use for the firewall endpoints, and any tags that are attached to the firewall Amazon Web Services resource. The status of the firewall, for example whether it's ready to filter network traffic, is provided in the corresponding FirewallStatus. You can retrieve both objects by calling DescribeFirewall.",
				Attributes: map[string]schema.Attribute{
					"delete_protection": schema.BoolAttribute{
						Description:         "A flag indicating whether it is possible to delete the firewall. A setting of TRUE indicates that the firewall is protected against deletion. Use this setting to protect against accidentally deleting a firewall that is in use. When you create a firewall, the operation initializes this flag to TRUE.",
						MarkdownDescription: "A flag indicating whether it is possible to delete the firewall. A setting of TRUE indicates that the firewall is protected against deletion. Use this setting to protect against accidentally deleting a firewall that is in use. When you create a firewall, the operation initializes this flag to TRUE.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"description": schema.StringAttribute{
						Description:         "A description of the firewall.",
						MarkdownDescription: "A description of the firewall.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"encryption_configuration": schema.SingleNestedAttribute{
						Description:         "A complex type that contains settings for encryption of your firewall resources.",
						MarkdownDescription: "A complex type that contains settings for encryption of your firewall resources.",
						Attributes: map[string]schema.Attribute{
							"key_id": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type_": schema.StringAttribute{
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

					"firewall_name": schema.StringAttribute{
						Description:         "The descriptive name of the firewall. You can't change the name of a firewall after you create it.",
						MarkdownDescription: "The descriptive name of the firewall. You can't change the name of a firewall after you create it.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"firewall_policy_arn": schema.StringAttribute{
						Description:         "The Amazon Resource Name (ARN) of the FirewallPolicy that you want to use for the firewall.",
						MarkdownDescription: "The Amazon Resource Name (ARN) of the FirewallPolicy that you want to use for the firewall.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"firewall_policy_change_protection": schema.BoolAttribute{
						Description:         "A setting indicating whether the firewall is protected against a change to the firewall policy association. Use this setting to protect against accidentally modifying the firewall policy for a firewall that is in use. When you create a firewall, the operation initializes this setting to TRUE.",
						MarkdownDescription: "A setting indicating whether the firewall is protected against a change to the firewall policy association. Use this setting to protect against accidentally modifying the firewall policy for a firewall that is in use. When you create a firewall, the operation initializes this setting to TRUE.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"subnet_change_protection": schema.BoolAttribute{
						Description:         "A setting indicating whether the firewall is protected against changes to the subnet associations. Use this setting to protect against accidentally modifying the subnet associations for a firewall that is in use. When you create a firewall, the operation initializes this setting to TRUE.",
						MarkdownDescription: "A setting indicating whether the firewall is protected against changes to the subnet associations. Use this setting to protect against accidentally modifying the subnet associations for a firewall that is in use. When you create a firewall, the operation initializes this setting to TRUE.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"subnet_mappings": schema.ListNestedAttribute{
						Description:         "The public subnets to use for your Network Firewall firewalls. Each subnet must belong to a different Availability Zone in the VPC. Network Firewall creates a firewall endpoint in each subnet.",
						MarkdownDescription: "The public subnets to use for your Network Firewall firewalls. Each subnet must belong to a different Availability Zone in the VPC. Network Firewall creates a firewall endpoint in each subnet.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"ip_address_type": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"subnet_id": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "The key:value pairs to associate with the resource.",
						MarkdownDescription: "The key:value pairs to associate with the resource.",
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

					"vpc_id": schema.StringAttribute{
						Description:         "The unique identifier of the VPC where Network Firewall should create the firewall. You can't change this setting after you create the firewall.",
						MarkdownDescription: "The unique identifier of the VPC where Network Firewall should create the firewall. You can't change this setting after you create the firewall.",
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
	}
}

func (r *NetworkfirewallServicesK8SAwsFirewallV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_networkfirewall_services_k8s_aws_firewall_v1alpha1_manifest")

	var model NetworkfirewallServicesK8SAwsFirewallV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("networkfirewall.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("Firewall")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
