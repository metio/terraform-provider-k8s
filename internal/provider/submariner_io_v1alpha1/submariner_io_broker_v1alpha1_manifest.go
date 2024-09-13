/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package submariner_io_v1alpha1

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
	_ datasource.DataSource = &SubmarinerIoBrokerV1Alpha1Manifest{}
)

func NewSubmarinerIoBrokerV1Alpha1Manifest() datasource.DataSource {
	return &SubmarinerIoBrokerV1Alpha1Manifest{}
}

type SubmarinerIoBrokerV1Alpha1Manifest struct{}

type SubmarinerIoBrokerV1Alpha1ManifestData struct {
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
		ClustersetIPCIDRRange       *string   `tfsdk:"clusterset_ipcidr_range" json:"clustersetIPCIDRRange,omitempty"`
		ClustersetIPEnabled         *bool     `tfsdk:"clusterset_ip_enabled" json:"clustersetIPEnabled,omitempty"`
		Components                  *[]string `tfsdk:"components" json:"components,omitempty"`
		DefaultCustomDomains        *[]string `tfsdk:"default_custom_domains" json:"defaultCustomDomains,omitempty"`
		DefaultGlobalnetClusterSize *int64    `tfsdk:"default_globalnet_cluster_size" json:"defaultGlobalnetClusterSize,omitempty"`
		GlobalnetCIDRRange          *string   `tfsdk:"globalnet_cidr_range" json:"globalnetCIDRRange,omitempty"`
		GlobalnetEnabled            *bool     `tfsdk:"globalnet_enabled" json:"globalnetEnabled,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SubmarinerIoBrokerV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_submariner_io_broker_v1alpha1_manifest"
}

func (r *SubmarinerIoBrokerV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Broker is the Schema for the brokers API.",
		MarkdownDescription: "Broker is the Schema for the brokers API.",
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
				Description:         "BrokerSpec defines the desired state of Broker.",
				MarkdownDescription: "BrokerSpec defines the desired state of Broker.",
				Attributes: map[string]schema.Attribute{
					"clusterset_ipcidr_range": schema.StringAttribute{
						Description:         "ClustersetIP supernet range for allocating ClustersetIPCIDRs to each cluster.",
						MarkdownDescription: "ClustersetIP supernet range for allocating ClustersetIPCIDRs to each cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"clusterset_ip_enabled": schema.BoolAttribute{
						Description:         "Enable ClustersetIP default for connecting clusters.",
						MarkdownDescription: "Enable ClustersetIP default for connecting clusters.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"components": schema.ListAttribute{
						Description:         "List of the components to be installed - any of [service-discovery, connectivity].",
						MarkdownDescription: "List of the components to be installed - any of [service-discovery, connectivity].",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"default_custom_domains": schema.ListAttribute{
						Description:         "List of domains to use for multi-cluster service discovery.",
						MarkdownDescription: "List of domains to use for multi-cluster service discovery.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"default_globalnet_cluster_size": schema.Int64Attribute{
						Description:         "Default cluster size for GlobalCIDR allocated to each cluster (amount of global IPs).",
						MarkdownDescription: "Default cluster size for GlobalCIDR allocated to each cluster (amount of global IPs).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"globalnet_cidr_range": schema.StringAttribute{
						Description:         "GlobalCIDR supernet range for allocating GlobalCIDRs to each cluster.",
						MarkdownDescription: "GlobalCIDR supernet range for allocating GlobalCIDRs to each cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"globalnet_enabled": schema.BoolAttribute{
						Description:         "Enable support for Overlapping CIDRs in connecting clusters.",
						MarkdownDescription: "Enable support for Overlapping CIDRs in connecting clusters.",
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

func (r *SubmarinerIoBrokerV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_submariner_io_broker_v1alpha1_manifest")

	var model SubmarinerIoBrokerV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("submariner.io/v1alpha1")
	model.Kind = pointer.String("Broker")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
