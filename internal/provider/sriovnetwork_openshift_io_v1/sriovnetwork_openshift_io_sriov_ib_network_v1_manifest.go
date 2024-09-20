/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package sriovnetwork_openshift_io_v1

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
	_ datasource.DataSource = &SriovnetworkOpenshiftIoSriovIbnetworkV1Manifest{}
)

func NewSriovnetworkOpenshiftIoSriovIbnetworkV1Manifest() datasource.DataSource {
	return &SriovnetworkOpenshiftIoSriovIbnetworkV1Manifest{}
}

type SriovnetworkOpenshiftIoSriovIbnetworkV1Manifest struct{}

type SriovnetworkOpenshiftIoSriovIbnetworkV1ManifestData struct {
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
		Capabilities     *string `tfsdk:"capabilities" json:"capabilities,omitempty"`
		Ipam             *string `tfsdk:"ipam" json:"ipam,omitempty"`
		LinkState        *string `tfsdk:"link_state" json:"linkState,omitempty"`
		MetaPlugins      *string `tfsdk:"meta_plugins" json:"metaPlugins,omitempty"`
		NetworkNamespace *string `tfsdk:"network_namespace" json:"networkNamespace,omitempty"`
		ResourceName     *string `tfsdk:"resource_name" json:"resourceName,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SriovnetworkOpenshiftIoSriovIbnetworkV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_sriovnetwork_openshift_io_sriov_ib_network_v1_manifest"
}

func (r *SriovnetworkOpenshiftIoSriovIbnetworkV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "SriovIBNetwork is the Schema for the sriovibnetworks API",
		MarkdownDescription: "SriovIBNetwork is the Schema for the sriovibnetworks API",
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
				Description:         "SriovIBNetworkSpec defines the desired state of SriovIBNetwork",
				MarkdownDescription: "SriovIBNetworkSpec defines the desired state of SriovIBNetwork",
				Attributes: map[string]schema.Attribute{
					"capabilities": schema.StringAttribute{
						Description:         "Capabilities to be configured for this network. Capabilities supported: (infinibandGUID), e.g. '{'infinibandGUID': true}'",
						MarkdownDescription: "Capabilities to be configured for this network. Capabilities supported: (infinibandGUID), e.g. '{'infinibandGUID': true}'",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ipam": schema.StringAttribute{
						Description:         "IPAM configuration to be used for this network.",
						MarkdownDescription: "IPAM configuration to be used for this network.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"link_state": schema.StringAttribute{
						Description:         "VF link state (enable|disable|auto)",
						MarkdownDescription: "VF link state (enable|disable|auto)",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("auto", "enable", "disable"),
						},
					},

					"meta_plugins": schema.StringAttribute{
						Description:         "MetaPluginsConfig configuration to be used in order to chain metaplugins to the sriov interface returned by the operator.",
						MarkdownDescription: "MetaPluginsConfig configuration to be used in order to chain metaplugins to the sriov interface returned by the operator.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"network_namespace": schema.StringAttribute{
						Description:         "Namespace of the NetworkAttachmentDefinition custom resource",
						MarkdownDescription: "Namespace of the NetworkAttachmentDefinition custom resource",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"resource_name": schema.StringAttribute{
						Description:         "SRIOV Network device plugin endpoint resource name",
						MarkdownDescription: "SRIOV Network device plugin endpoint resource name",
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

func (r *SriovnetworkOpenshiftIoSriovIbnetworkV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_sriovnetwork_openshift_io_sriov_ib_network_v1_manifest")

	var model SriovnetworkOpenshiftIoSriovIbnetworkV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("sriovnetwork.openshift.io/v1")
	model.Kind = pointer.String("SriovIBNetwork")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
