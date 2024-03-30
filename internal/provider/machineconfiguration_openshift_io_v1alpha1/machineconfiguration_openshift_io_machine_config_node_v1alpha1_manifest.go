/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package machineconfiguration_openshift_io_v1alpha1

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
	_ datasource.DataSource = &MachineconfigurationOpenshiftIoMachineConfigNodeV1Alpha1Manifest{}
)

func NewMachineconfigurationOpenshiftIoMachineConfigNodeV1Alpha1Manifest() datasource.DataSource {
	return &MachineconfigurationOpenshiftIoMachineConfigNodeV1Alpha1Manifest{}
}

type MachineconfigurationOpenshiftIoMachineConfigNodeV1Alpha1Manifest struct{}

type MachineconfigurationOpenshiftIoMachineConfigNodeV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ConfigVersion *struct {
			Desired *string `tfsdk:"desired" json:"desired,omitempty"`
		} `tfsdk:"config_version" json:"configVersion,omitempty"`
		Node *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"node" json:"node,omitempty"`
		Pool *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"pool" json:"pool,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *MachineconfigurationOpenshiftIoMachineConfigNodeV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_machineconfiguration_openshift_io_machine_config_node_v1alpha1_manifest"
}

func (r *MachineconfigurationOpenshiftIoMachineConfigNodeV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "MachineConfigNode describes the health of the Machines on the system Compatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
		MarkdownDescription: "MachineConfigNode describes the health of the Machines on the system Compatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
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
				Description:         "spec describes the configuration of the machine config node.",
				MarkdownDescription: "spec describes the configuration of the machine config node.",
				Attributes: map[string]schema.Attribute{
					"config_version": schema.SingleNestedAttribute{
						Description:         "configVersion holds the desired config version for the node targeted by this machine config node resource. The desired version represents the machine config the node will attempt to update to. This gets set before the machine config operator validates the new machine config against the current machine config.",
						MarkdownDescription: "configVersion holds the desired config version for the node targeted by this machine config node resource. The desired version represents the machine config the node will attempt to update to. This gets set before the machine config operator validates the new machine config against the current machine config.",
						Attributes: map[string]schema.Attribute{
							"desired": schema.StringAttribute{
								Description:         "desired is the name of the machine config that the the node should be upgraded to. This value is set when the machine config pool generates a new version of its rendered configuration. When this value is changed, the machine config daemon starts the node upgrade process. This value gets set in the machine config node spec once the machine config has been targeted for upgrade and before it is validated. Must be a lowercase RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) It may consist of only alphanumeric characters, hyphens (-) and periods (.) and must be at most 253 characters in length.",
								MarkdownDescription: "desired is the name of the machine config that the the node should be upgraded to. This value is set when the machine config pool generates a new version of its rendered configuration. When this value is changed, the machine config daemon starts the node upgrade process. This value gets set in the machine config node spec once the machine config has been targeted for upgrade and before it is validated. Must be a lowercase RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) It may consist of only alphanumeric characters, hyphens (-) and periods (.) and must be at most 253 characters in length.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtMost(253),
									stringvalidator.RegexMatches(regexp.MustCompile(`^([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])(\.([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9]))*$`), ""),
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"node": schema.SingleNestedAttribute{
						Description:         "node contains a reference to the node for this machine config node.",
						MarkdownDescription: "node contains a reference to the node for this machine config node.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "name is the object name. Must be a lowercase RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) It may consist of only alphanumeric characters, hyphens (-) and periods (.) and must be at most 253 characters in length.",
								MarkdownDescription: "name is the object name. Must be a lowercase RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) It may consist of only alphanumeric characters, hyphens (-) and periods (.) and must be at most 253 characters in length.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtMost(253),
									stringvalidator.RegexMatches(regexp.MustCompile(`^([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])(\.([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9]))*$`), ""),
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"pool": schema.SingleNestedAttribute{
						Description:         "pool contains a reference to the machine config pool that this machine config node's referenced node belongs to.",
						MarkdownDescription: "pool contains a reference to the machine config pool that this machine config node's referenced node belongs to.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "name is the object name. Must be a lowercase RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) It may consist of only alphanumeric characters, hyphens (-) and periods (.) and must be at most 253 characters in length.",
								MarkdownDescription: "name is the object name. Must be a lowercase RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) It may consist of only alphanumeric characters, hyphens (-) and periods (.) and must be at most 253 characters in length.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtMost(253),
									stringvalidator.RegexMatches(regexp.MustCompile(`^([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])(\.([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9]))*$`), ""),
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

func (r *MachineconfigurationOpenshiftIoMachineConfigNodeV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_machineconfiguration_openshift_io_machine_config_node_v1alpha1_manifest")

	var model MachineconfigurationOpenshiftIoMachineConfigNodeV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("machineconfiguration.openshift.io/v1alpha1")
	model.Kind = pointer.String("MachineConfigNode")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
