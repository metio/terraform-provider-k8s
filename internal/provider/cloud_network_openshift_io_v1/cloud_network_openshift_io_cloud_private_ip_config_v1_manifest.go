/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package cloud_network_openshift_io_v1

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
	_ datasource.DataSource = &CloudNetworkOpenshiftIoCloudPrivateIpconfigV1Manifest{}
)

func NewCloudNetworkOpenshiftIoCloudPrivateIpconfigV1Manifest() datasource.DataSource {
	return &CloudNetworkOpenshiftIoCloudPrivateIpconfigV1Manifest{}
}

type CloudNetworkOpenshiftIoCloudPrivateIpconfigV1Manifest struct{}

type CloudNetworkOpenshiftIoCloudPrivateIpconfigV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Node *string `tfsdk:"node" json:"node,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CloudNetworkOpenshiftIoCloudPrivateIpconfigV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cloud_network_openshift_io_cloud_private_ip_config_v1_manifest"
}

func (r *CloudNetworkOpenshiftIoCloudPrivateIpconfigV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CloudPrivateIPConfig performs an assignment of a private IP address to the primary NIC associated with cloud VMs. This is done by specifying the IP and Kubernetes node which the IP should be assigned to. This CRD is intended to be used by the network plugin which manages the cluster network. The spec side represents the desired state requested by the network plugin, and the status side represents the current state that this CRD's controller has executed. No users will have permission to modify it, and if a cluster-admin decides to edit it for some reason, their changes will be overwritten the next time the network plugin reconciles the object. Note: the CR's name must specify the requested private IP address (can be IPv4 or IPv6).  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		MarkdownDescription: "CloudPrivateIPConfig performs an assignment of a private IP address to the primary NIC associated with cloud VMs. This is done by specifying the IP and Kubernetes node which the IP should be assigned to. This CRD is intended to be used by the network plugin which manages the cluster network. The spec side represents the desired state requested by the network plugin, and the status side represents the current state that this CRD's controller has executed. No users will have permission to modify it, and if a cluster-admin decides to edit it for some reason, their changes will be overwritten the next time the network plugin reconciles the object. Note: the CR's name must specify the requested private IP address (can be IPv4 or IPv6).  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
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
				Description:         "spec is the definition of the desired private IP request.",
				MarkdownDescription: "spec is the definition of the desired private IP request.",
				Attributes: map[string]schema.Attribute{
					"node": schema.StringAttribute{
						Description:         "node is the node name, as specified by the Kubernetes field: node.metadata.name",
						MarkdownDescription: "node is the node name, as specified by the Kubernetes field: node.metadata.name",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *CloudNetworkOpenshiftIoCloudPrivateIpconfigV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cloud_network_openshift_io_cloud_private_ip_config_v1_manifest")

	var model CloudNetworkOpenshiftIoCloudPrivateIpconfigV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("cloud.network.openshift.io/v1")
	model.Kind = pointer.String("CloudPrivateIPConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
