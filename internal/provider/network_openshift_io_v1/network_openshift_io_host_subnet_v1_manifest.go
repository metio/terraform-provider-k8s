/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package network_openshift_io_v1

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
	_ datasource.DataSource = &NetworkOpenshiftIoHostSubnetV1Manifest{}
)

func NewNetworkOpenshiftIoHostSubnetV1Manifest() datasource.DataSource {
	return &NetworkOpenshiftIoHostSubnetV1Manifest{}
}

type NetworkOpenshiftIoHostSubnetV1Manifest struct{}

type NetworkOpenshiftIoHostSubnetV1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	EgressCIDRs *[]string `tfsdk:"egress_cid_rs" json:"egressCIDRs,omitempty"`
	EgressIPs   *[]string `tfsdk:"egress_i_ps" json:"egressIPs,omitempty"`
	Host        *string   `tfsdk:"host" json:"host,omitempty"`
	HostIP      *string   `tfsdk:"host_ip" json:"hostIP,omitempty"`
	Subnet      *string   `tfsdk:"subnet" json:"subnet,omitempty"`
}

func (r *NetworkOpenshiftIoHostSubnetV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_network_openshift_io_host_subnet_v1_manifest"
}

func (r *NetworkOpenshiftIoHostSubnetV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "HostSubnet describes the container subnet network on a node. The HostSubnet object must have the same name as the Node object it corresponds to.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		MarkdownDescription: "HostSubnet describes the container subnet network on a node. The HostSubnet object must have the same name as the Node object it corresponds to.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
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

			"egress_cid_rs": schema.ListAttribute{
				Description:         "EgressCIDRs is the list of CIDR ranges available for automatically assigning egress IPs to this node from. If this field is set then EgressIPs should be treated as read-only.",
				MarkdownDescription: "EgressCIDRs is the list of CIDR ranges available for automatically assigning egress IPs to this node from. If this field is set then EgressIPs should be treated as read-only.",
				ElementType:         types.StringType,
				Required:            false,
				Optional:            true,
				Computed:            false,
			},

			"egress_i_ps": schema.ListAttribute{
				Description:         "EgressIPs is the list of automatic egress IP addresses currently hosted by this node. If EgressCIDRs is empty, this can be set by hand; if EgressCIDRs is set then the master will overwrite the value here with its own allocation of egress IPs.",
				MarkdownDescription: "EgressIPs is the list of automatic egress IP addresses currently hosted by this node. If EgressCIDRs is empty, this can be set by hand; if EgressCIDRs is set then the master will overwrite the value here with its own allocation of egress IPs.",
				ElementType:         types.StringType,
				Required:            false,
				Optional:            true,
				Computed:            false,
			},

			"host": schema.StringAttribute{
				Description:         "Host is the name of the node. (This is the same as the object's name, but both fields must be set.)",
				MarkdownDescription: "Host is the name of the node. (This is the same as the object's name, but both fields must be set.)",
				Required:            true,
				Optional:            false,
				Computed:            false,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9.-]+$`), ""),
				},
			},

			"host_ip": schema.StringAttribute{
				Description:         "HostIP is the IP address to be used as a VTEP by other nodes in the overlay network",
				MarkdownDescription: "HostIP is the IP address to be used as a VTEP by other nodes in the overlay network",
				Required:            true,
				Optional:            false,
				Computed:            false,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]|[0-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[0-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$`), ""),
				},
			},

			"subnet": schema.StringAttribute{
				Description:         "Subnet is the CIDR range of the overlay network assigned to the node for its pods",
				MarkdownDescription: "Subnet is the CIDR range of the overlay network assigned to the node for its pods",
				Required:            true,
				Optional:            false,
				Computed:            false,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]|[0-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[0-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/([0-9]|[12][0-9]|3[0-2])$`), ""),
				},
			},
		},
	}
}

func (r *NetworkOpenshiftIoHostSubnetV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_network_openshift_io_host_subnet_v1_manifest")

	var model NetworkOpenshiftIoHostSubnetV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("network.openshift.io/v1")
	model.Kind = pointer.String("HostSubnet")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
