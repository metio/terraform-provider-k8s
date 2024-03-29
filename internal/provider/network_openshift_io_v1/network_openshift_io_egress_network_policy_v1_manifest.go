/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package network_openshift_io_v1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &NetworkOpenshiftIoEgressNetworkPolicyV1Manifest{}
)

func NewNetworkOpenshiftIoEgressNetworkPolicyV1Manifest() datasource.DataSource {
	return &NetworkOpenshiftIoEgressNetworkPolicyV1Manifest{}
}

type NetworkOpenshiftIoEgressNetworkPolicyV1Manifest struct{}

type NetworkOpenshiftIoEgressNetworkPolicyV1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
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
		Egress *[]struct {
			To *struct {
				CidrSelector *string `tfsdk:"cidr_selector" json:"cidrSelector,omitempty"`
				DnsName      *string `tfsdk:"dns_name" json:"dnsName,omitempty"`
			} `tfsdk:"to" json:"to,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"egress" json:"egress,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *NetworkOpenshiftIoEgressNetworkPolicyV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_network_openshift_io_egress_network_policy_v1_manifest"
}

func (r *NetworkOpenshiftIoEgressNetworkPolicyV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "EgressNetworkPolicy describes the current egress network policy for a Namespace. When using the 'redhat/openshift-ovs-multitenant' network plugin, traffic from a pod to an IP address outside the cluster will be checked against each EgressNetworkPolicyRule in the pod's namespace's EgressNetworkPolicy, in order. If no rule matches (or no EgressNetworkPolicy is present) then the traffic will be allowed by default.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		MarkdownDescription: "EgressNetworkPolicy describes the current egress network policy for a Namespace. When using the 'redhat/openshift-ovs-multitenant' network plugin, traffic from a pod to an IP address outside the cluster will be checked against each EgressNetworkPolicyRule in the pod's namespace's EgressNetworkPolicy, in order. If no rule matches (or no EgressNetworkPolicy is present) then the traffic will be allowed by default.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
				Description:         "spec is the specification of the current egress network policy",
				MarkdownDescription: "spec is the specification of the current egress network policy",
				Attributes: map[string]schema.Attribute{
					"egress": schema.ListNestedAttribute{
						Description:         "egress contains the list of egress policy rules",
						MarkdownDescription: "egress contains the list of egress policy rules",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"to": schema.SingleNestedAttribute{
									Description:         "to is the target that traffic is allowed/denied to",
									MarkdownDescription: "to is the target that traffic is allowed/denied to",
									Attributes: map[string]schema.Attribute{
										"cidr_selector": schema.StringAttribute{
											Description:         "CIDRSelector is the CIDR range to allow/deny traffic to. If this is set, dnsName must be unset Ideally we would have liked to use the cidr openapi format for this property. But openshift-sdn only supports v4 while specifying the cidr format allows both v4 and v6 cidrs We are therefore using a regex pattern to validate instead.",
											MarkdownDescription: "CIDRSelector is the CIDR range to allow/deny traffic to. If this is set, dnsName must be unset Ideally we would have liked to use the cidr openapi format for this property. But openshift-sdn only supports v4 while specifying the cidr format allows both v4 and v6 cidrs We are therefore using a regex pattern to validate instead.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]|[0-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[0-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/([0-9]|[12][0-9]|3[0-2])$`), ""),
											},
										},

										"dns_name": schema.StringAttribute{
											Description:         "DNSName is the domain name to allow/deny traffic to. If this is set, cidrSelector must be unset",
											MarkdownDescription: "DNSName is the domain name to allow/deny traffic to. If this is set, cidrSelector must be unset",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^([A-Za-z0-9-]+\.)*[A-Za-z0-9-]+\.?$`), ""),
											},
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"type": schema.StringAttribute{
									Description:         "type marks this as an 'Allow' or 'Deny' rule",
									MarkdownDescription: "type marks this as an 'Allow' or 'Deny' rule",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.RegexMatches(regexp.MustCompile(`^Allow|Deny$`), ""),
									},
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *NetworkOpenshiftIoEgressNetworkPolicyV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_network_openshift_io_egress_network_policy_v1_manifest")

	var model NetworkOpenshiftIoEgressNetworkPolicyV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("network.openshift.io/v1")
	model.Kind = pointer.String("EgressNetworkPolicy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
