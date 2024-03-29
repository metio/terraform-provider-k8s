/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package chisel_operator_io_v1

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
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &ChiselOperatorIoExitNodeProvisionerV1Manifest{}
)

func NewChiselOperatorIoExitNodeProvisionerV1Manifest() datasource.DataSource {
	return &ChiselOperatorIoExitNodeProvisionerV1Manifest{}
}

type ChiselOperatorIoExitNodeProvisionerV1Manifest struct{}

type ChiselOperatorIoExitNodeProvisionerV1ManifestData struct {
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
		AWS *struct {
			Auth           *string `tfsdk:"auth" json:"auth,omitempty"`
			Region         *string `tfsdk:"region" json:"region,omitempty"`
			Security_group *string `tfsdk:"security_group" json:"security_group,omitempty"`
			Size           *string `tfsdk:"size" json:"size,omitempty"`
		} `tfsdk:"aws" json:"AWS,omitempty"`
		DigitalOcean *struct {
			Auth             *string   `tfsdk:"auth" json:"auth,omitempty"`
			Region           *string   `tfsdk:"region" json:"region,omitempty"`
			Size             *string   `tfsdk:"size" json:"size,omitempty"`
			Ssh_fingerprints *[]string `tfsdk:"ssh_fingerprints" json:"ssh_fingerprints,omitempty"`
		} `tfsdk:"digital_ocean" json:"DigitalOcean,omitempty"`
		Linode *struct {
			Auth   *string `tfsdk:"auth" json:"auth,omitempty"`
			Region *string `tfsdk:"region" json:"region,omitempty"`
			Size   *string `tfsdk:"size" json:"size,omitempty"`
		} `tfsdk:"linode" json:"Linode,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ChiselOperatorIoExitNodeProvisionerV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_chisel_operator_io_exit_node_provisioner_v1_manifest"
}

func (r *ChiselOperatorIoExitNodeProvisionerV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Auto-generated derived type for ExitNodeProvisionerSpec via 'CustomResource'",
		MarkdownDescription: "Auto-generated derived type for ExitNodeProvisionerSpec via 'CustomResource'",
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
				Description:         "ExitNodeProvisioner is a custom resource that represents a Chisel exit node provisioner on a cloud provider.",
				MarkdownDescription: "ExitNodeProvisioner is a custom resource that represents a Chisel exit node provisioner on a cloud provider.",
				Attributes: map[string]schema.Attribute{
					"aws": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"auth": schema.StringAttribute{
								Description:         "Reference to a secret containing the AWS access key ID and secret access key, under the 'access_key_id' and 'secret_access_key' secret keys",
								MarkdownDescription: "Reference to a secret containing the AWS access key ID and secret access key, under the 'access_key_id' and 'secret_access_key' secret keys",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"region": schema.StringAttribute{
								Description:         "Region ID for the AWS region to provision the exit node in See https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html",
								MarkdownDescription: "Region ID for the AWS region to provision the exit node in See https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"security_group": schema.StringAttribute{
								Description:         "Security group name to use for the exit node, uses the default security group if not specified",
								MarkdownDescription: "Security group name to use for the exit node, uses the default security group if not specified",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"size": schema.StringAttribute{
								Description:         "Size for the EC2 instance See https://aws.amazon.com/ec2/instance-types/",
								MarkdownDescription: "Size for the EC2 instance See https://aws.amazon.com/ec2/instance-types/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"digital_ocean": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"auth": schema.StringAttribute{
								Description:         "Reference to a secret containing the DigitalOcean API token, under the 'DIGITALOCEAN_TOKEN' secret key",
								MarkdownDescription: "Reference to a secret containing the DigitalOcean API token, under the 'DIGITALOCEAN_TOKEN' secret key",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"region": schema.StringAttribute{
								Description:         "Region ID of the DigitalOcean datacenter to provision the exit node in If empty, DigitalOcean will randomly select a region for you, which might not be what you want See https://slugs.do-api.dev/",
								MarkdownDescription: "Region ID of the DigitalOcean datacenter to provision the exit node in If empty, DigitalOcean will randomly select a region for you, which might not be what you want See https://slugs.do-api.dev/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"size": schema.StringAttribute{
								Description:         "Size for the DigitalOcean droplet See https://slugs.do-api.dev/",
								MarkdownDescription: "Size for the DigitalOcean droplet See https://slugs.do-api.dev/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ssh_fingerprints": schema.ListAttribute{
								Description:         "SSH key fingerprints to add to the exit node",
								MarkdownDescription: "SSH key fingerprints to add to the exit node",
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

					"linode": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"auth": schema.StringAttribute{
								Description:         "Name of the secret containing the Linode API token, under the 'LINODE_TOKEN' secret key",
								MarkdownDescription: "Name of the secret containing the Linode API token, under the 'LINODE_TOKEN' secret key",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"region": schema.StringAttribute{
								Description:         "Region ID of the Linode datacenter to provision the exit node in See https://api.linode.com/v4/regions",
								MarkdownDescription: "Region ID of the Linode datacenter to provision the exit node in See https://api.linode.com/v4/regions",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"size": schema.StringAttribute{
								Description:         "Size for the Linode instance See https://api.linode.com/v4/linode/",
								MarkdownDescription: "Size for the Linode instance See https://api.linode.com/v4/linode/",
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
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *ChiselOperatorIoExitNodeProvisionerV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_chisel_operator_io_exit_node_provisioner_v1_manifest")

	var model ChiselOperatorIoExitNodeProvisionerV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("chisel-operator.io/v1")
	model.Kind = pointer.String("ExitNodeProvisioner")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
