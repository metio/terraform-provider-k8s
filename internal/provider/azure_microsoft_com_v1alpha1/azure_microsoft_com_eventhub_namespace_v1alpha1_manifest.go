/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package azure_microsoft_com_v1alpha1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &AzureMicrosoftComEventhubNamespaceV1Alpha1Manifest{}
)

func NewAzureMicrosoftComEventhubNamespaceV1Alpha1Manifest() datasource.DataSource {
	return &AzureMicrosoftComEventhubNamespaceV1Alpha1Manifest{}
}

type AzureMicrosoftComEventhubNamespaceV1Alpha1Manifest struct{}

type AzureMicrosoftComEventhubNamespaceV1Alpha1ManifestData struct {
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
		Location    *string `tfsdk:"location" json:"location,omitempty"`
		NetworkRule *struct {
			DefaultAction *string `tfsdk:"default_action" json:"defaultAction,omitempty"`
			IpRules       *[]struct {
				IpMask *string `tfsdk:"ip_mask" json:"ipMask,omitempty"`
			} `tfsdk:"ip_rules" json:"ipRules,omitempty"`
			VirtualNetworkRules *[]struct {
				IgnoreMissingServiceEndpoint *bool   `tfsdk:"ignore_missing_service_endpoint" json:"ignoreMissingServiceEndpoint,omitempty"`
				SubnetId                     *string `tfsdk:"subnet_id" json:"subnetId,omitempty"`
			} `tfsdk:"virtual_network_rules" json:"virtualNetworkRules,omitempty"`
		} `tfsdk:"network_rule" json:"networkRule,omitempty"`
		Properties *struct {
			IsAutoInflateEnabled   *bool  `tfsdk:"is_auto_inflate_enabled" json:"isAutoInflateEnabled,omitempty"`
			KafkaEnabled           *bool  `tfsdk:"kafka_enabled" json:"kafkaEnabled,omitempty"`
			MaximumThroughputUnits *int64 `tfsdk:"maximum_throughput_units" json:"maximumThroughputUnits,omitempty"`
		} `tfsdk:"properties" json:"properties,omitempty"`
		ResourceGroup *string `tfsdk:"resource_group" json:"resourceGroup,omitempty"`
		Sku           *struct {
			Capacity *int64  `tfsdk:"capacity" json:"capacity,omitempty"`
			Name     *string `tfsdk:"name" json:"name,omitempty"`
			Tier     *string `tfsdk:"tier" json:"tier,omitempty"`
		} `tfsdk:"sku" json:"sku,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AzureMicrosoftComEventhubNamespaceV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_azure_microsoft_com_eventhub_namespace_v1alpha1_manifest"
}

func (r *AzureMicrosoftComEventhubNamespaceV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "EventhubNamespace is the Schema for the eventhubnamespaces API",
		MarkdownDescription: "EventhubNamespace is the Schema for the eventhubnamespaces API",
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
				Description:         "EventhubNamespaceSpec defines the desired state of EventhubNamespace",
				MarkdownDescription: "EventhubNamespaceSpec defines the desired state of EventhubNamespace",
				Attributes: map[string]schema.Attribute{
					"location": schema.StringAttribute{
						Description:         "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run 'make' to regenerate code after modifying this file",
						MarkdownDescription: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run 'make' to regenerate code after modifying this file",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"network_rule": schema.SingleNestedAttribute{
						Description:         "EventhubNamespaceNetworkRule defines the namespace network rule",
						MarkdownDescription: "EventhubNamespaceNetworkRule defines the namespace network rule",
						Attributes: map[string]schema.Attribute{
							"default_action": schema.StringAttribute{
								Description:         "DefaultAction defined as a string",
								MarkdownDescription: "DefaultAction defined as a string",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ip_rules": schema.ListNestedAttribute{
								Description:         "IPRules - List of IpRules",
								MarkdownDescription: "IPRules - List of IpRules",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"ip_mask": schema.StringAttribute{
											Description:         "IPMask - IPv4 address 1.1.1.1 or CIDR notation 1.1.0.0/24",
											MarkdownDescription: "IPMask - IPv4 address 1.1.1.1 or CIDR notation 1.1.0.0/24",
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

							"virtual_network_rules": schema.ListNestedAttribute{
								Description:         "VirtualNetworkRules - List VirtualNetwork Rules",
								MarkdownDescription: "VirtualNetworkRules - List VirtualNetwork Rules",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"ignore_missing_service_endpoint": schema.BoolAttribute{
											Description:         "IgnoreMissingVnetServiceEndpoint - Value that indicates whether to ignore missing VNet Service Endpoint",
											MarkdownDescription: "IgnoreMissingVnetServiceEndpoint - Value that indicates whether to ignore missing VNet Service Endpoint",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"subnet_id": schema.StringAttribute{
											Description:         "Subnet - Full Resource ID of Virtual Network Subnet",
											MarkdownDescription: "Subnet - Full Resource ID of Virtual Network Subnet",
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

					"properties": schema.SingleNestedAttribute{
						Description:         "EventhubNamespaceProperties defines the namespace properties",
						MarkdownDescription: "EventhubNamespaceProperties defines the namespace properties",
						Attributes: map[string]schema.Attribute{
							"is_auto_inflate_enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kafka_enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"maximum_throughput_units": schema.Int64Attribute{
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

					"resource_group": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
							stringvalidator.RegexMatches(regexp.MustCompile(`^[-\w\._\(\)]+$`), ""),
						},
					},

					"sku": schema.SingleNestedAttribute{
						Description:         "EventhubNamespaceSku defines the sku",
						MarkdownDescription: "EventhubNamespaceSku defines the sku",
						Attributes: map[string]schema.Attribute{
							"capacity": schema.Int64Attribute{
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

							"tier": schema.StringAttribute{
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

func (r *AzureMicrosoftComEventhubNamespaceV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_azure_microsoft_com_eventhub_namespace_v1alpha1_manifest")

	var model AzureMicrosoftComEventhubNamespaceV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("azure.microsoft.com/v1alpha1")
	model.Kind = pointer.String("EventhubNamespace")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
