/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package config_openshift_io_v1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	_ datasource.DataSource = &ConfigOpenshiftIoNetworkV1Manifest{}
)

func NewConfigOpenshiftIoNetworkV1Manifest() datasource.DataSource {
	return &ConfigOpenshiftIoNetworkV1Manifest{}
}

type ConfigOpenshiftIoNetworkV1Manifest struct{}

type ConfigOpenshiftIoNetworkV1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ClusterNetwork *[]struct {
			Cidr       *string `tfsdk:"cidr" json:"cidr,omitempty"`
			HostPrefix *int64  `tfsdk:"host_prefix" json:"hostPrefix,omitempty"`
		} `tfsdk:"cluster_network" json:"clusterNetwork,omitempty"`
		ExternalIP *struct {
			AutoAssignCIDRs *[]string `tfsdk:"auto_assign_cid_rs" json:"autoAssignCIDRs,omitempty"`
			Policy          *struct {
				AllowedCIDRs  *[]string `tfsdk:"allowed_cid_rs" json:"allowedCIDRs,omitempty"`
				RejectedCIDRs *[]string `tfsdk:"rejected_cid_rs" json:"rejectedCIDRs,omitempty"`
			} `tfsdk:"policy" json:"policy,omitempty"`
		} `tfsdk:"external_ip" json:"externalIP,omitempty"`
		NetworkType          *string   `tfsdk:"network_type" json:"networkType,omitempty"`
		ServiceNetwork       *[]string `tfsdk:"service_network" json:"serviceNetwork,omitempty"`
		ServiceNodePortRange *string   `tfsdk:"service_node_port_range" json:"serviceNodePortRange,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ConfigOpenshiftIoNetworkV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_config_openshift_io_network_v1_manifest"
}

func (r *ConfigOpenshiftIoNetworkV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Network holds cluster-wide information about Network. The canonical name is 'cluster'. It is used to configure the desired network configuration, such as: IP address pools for services/pod IPs, network plugin, etc. Please view network.spec for an explanation on what applies when configuring this resource.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		MarkdownDescription: "Network holds cluster-wide information about Network. The canonical name is 'cluster'. It is used to configure the desired network configuration, such as: IP address pools for services/pod IPs, network plugin, etc. Please view network.spec for an explanation on what applies when configuring this resource.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

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
				Description:         "spec holds user settable values for configuration. As a general rule, this SHOULD NOT be read directly. Instead, you should consume the NetworkStatus, as it indicates the currently deployed configuration. Currently, most spec fields are immutable after installation. Please view the individual ones for further details on each.",
				MarkdownDescription: "spec holds user settable values for configuration. As a general rule, this SHOULD NOT be read directly. Instead, you should consume the NetworkStatus, as it indicates the currently deployed configuration. Currently, most spec fields are immutable after installation. Please view the individual ones for further details on each.",
				Attributes: map[string]schema.Attribute{
					"cluster_network": schema.ListNestedAttribute{
						Description:         "IP address pool to use for pod IPs. This field is immutable after installation.",
						MarkdownDescription: "IP address pool to use for pod IPs. This field is immutable after installation.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"cidr": schema.StringAttribute{
									Description:         "The complete block for pod IPs.",
									MarkdownDescription: "The complete block for pod IPs.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"host_prefix": schema.Int64Attribute{
									Description:         "The size (prefix) of block to allocate to each node. If this field is not used by the plugin, it can be left unset.",
									MarkdownDescription: "The size (prefix) of block to allocate to each node. If this field is not used by the plugin, it can be left unset.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(0),
									},
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"external_ip": schema.SingleNestedAttribute{
						Description:         "externalIP defines configuration for controllers that affect Service.ExternalIP. If nil, then ExternalIP is not allowed to be set.",
						MarkdownDescription: "externalIP defines configuration for controllers that affect Service.ExternalIP. If nil, then ExternalIP is not allowed to be set.",
						Attributes: map[string]schema.Attribute{
							"auto_assign_cid_rs": schema.ListAttribute{
								Description:         "autoAssignCIDRs is a list of CIDRs from which to automatically assign Service.ExternalIP. These are assigned when the service is of type LoadBalancer. In general, this is only useful for bare-metal clusters. In Openshift 3.x, this was misleadingly called 'IngressIPs'. Automatically assigned External IPs are not affected by any ExternalIPPolicy rules. Currently, only one entry may be provided.",
								MarkdownDescription: "autoAssignCIDRs is a list of CIDRs from which to automatically assign Service.ExternalIP. These are assigned when the service is of type LoadBalancer. In general, this is only useful for bare-metal clusters. In Openshift 3.x, this was misleadingly called 'IngressIPs'. Automatically assigned External IPs are not affected by any ExternalIPPolicy rules. Currently, only one entry may be provided.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"policy": schema.SingleNestedAttribute{
								Description:         "policy is a set of restrictions applied to the ExternalIP field. If nil or empty, then ExternalIP is not allowed to be set.",
								MarkdownDescription: "policy is a set of restrictions applied to the ExternalIP field. If nil or empty, then ExternalIP is not allowed to be set.",
								Attributes: map[string]schema.Attribute{
									"allowed_cid_rs": schema.ListAttribute{
										Description:         "allowedCIDRs is the list of allowed CIDRs.",
										MarkdownDescription: "allowedCIDRs is the list of allowed CIDRs.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"rejected_cid_rs": schema.ListAttribute{
										Description:         "rejectedCIDRs is the list of disallowed CIDRs. These take precedence over allowedCIDRs.",
										MarkdownDescription: "rejectedCIDRs is the list of disallowed CIDRs. These take precedence over allowedCIDRs.",
										ElementType:         types.StringType,
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

					"network_type": schema.StringAttribute{
						Description:         "NetworkType is the plugin that is to be deployed (e.g. OpenShiftSDN). This should match a value that the cluster-network-operator understands, or else no networking will be installed. Currently supported values are: - OpenShiftSDN This field is immutable after installation.",
						MarkdownDescription: "NetworkType is the plugin that is to be deployed (e.g. OpenShiftSDN). This should match a value that the cluster-network-operator understands, or else no networking will be installed. Currently supported values are: - OpenShiftSDN This field is immutable after installation.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service_network": schema.ListAttribute{
						Description:         "IP address pool for services. Currently, we only support a single entry here. This field is immutable after installation.",
						MarkdownDescription: "IP address pool for services. Currently, we only support a single entry here. This field is immutable after installation.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service_node_port_range": schema.StringAttribute{
						Description:         "The port range allowed for Services of type NodePort. If not specified, the default of 30000-32767 will be used. Such Services without a NodePort specified will have one automatically allocated from this range. This parameter can be updated after the cluster is installed.",
						MarkdownDescription: "The port range allowed for Services of type NodePort. If not specified, the default of 30000-32767 will be used. Such Services without a NodePort specified will have one automatically allocated from this range. This parameter can be updated after the cluster is installed.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]{1,4}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])-([0-9]{1,4}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])$`), ""),
						},
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *ConfigOpenshiftIoNetworkV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_config_openshift_io_network_v1_manifest")

	var model ConfigOpenshiftIoNetworkV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("config.openshift.io/v1")
	model.Kind = pointer.String("Network")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
