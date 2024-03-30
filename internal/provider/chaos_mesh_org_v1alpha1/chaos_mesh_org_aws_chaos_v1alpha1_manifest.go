/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package chaos_mesh_org_v1alpha1

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
	_ datasource.DataSource = &ChaosMeshOrgAwschaosV1Alpha1Manifest{}
)

func NewChaosMeshOrgAwschaosV1Alpha1Manifest() datasource.DataSource {
	return &ChaosMeshOrgAwschaosV1Alpha1Manifest{}
}

type ChaosMeshOrgAwschaosV1Alpha1Manifest struct{}

type ChaosMeshOrgAwschaosV1Alpha1ManifestData struct {
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
		Action        *string `tfsdk:"action" json:"action,omitempty"`
		AwsRegion     *string `tfsdk:"aws_region" json:"awsRegion,omitempty"`
		DeviceName    *string `tfsdk:"device_name" json:"deviceName,omitempty"`
		Duration      *string `tfsdk:"duration" json:"duration,omitempty"`
		Ec2Instance   *string `tfsdk:"ec2_instance" json:"ec2Instance,omitempty"`
		Endpoint      *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
		RemoteCluster *string `tfsdk:"remote_cluster" json:"remoteCluster,omitempty"`
		SecretName    *string `tfsdk:"secret_name" json:"secretName,omitempty"`
		VolumeID      *string `tfsdk:"volume_id" json:"volumeID,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ChaosMeshOrgAwschaosV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_chaos_mesh_org_aws_chaos_v1alpha1_manifest"
}

func (r *ChaosMeshOrgAwschaosV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "AWSChaos is the Schema for the awschaos API",
		MarkdownDescription: "AWSChaos is the Schema for the awschaos API",
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
				Description:         "AWSChaosSpec is the content of the specification for an AWSChaos",
				MarkdownDescription: "AWSChaosSpec is the content of the specification for an AWSChaos",
				Attributes: map[string]schema.Attribute{
					"action": schema.StringAttribute{
						Description:         "Action defines the specific aws chaos action. Supported action: ec2-stop / ec2-restart / detach-volume Default action: ec2-stop",
						MarkdownDescription: "Action defines the specific aws chaos action. Supported action: ec2-stop / ec2-restart / detach-volume Default action: ec2-stop",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("ec2-stop", "ec2-restart", "detach-volume"),
						},
					},

					"aws_region": schema.StringAttribute{
						Description:         "AWSRegion defines the region of aws.",
						MarkdownDescription: "AWSRegion defines the region of aws.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"device_name": schema.StringAttribute{
						Description:         "DeviceName indicates the name of the device. Needed in detach-volume.",
						MarkdownDescription: "DeviceName indicates the name of the device. Needed in detach-volume.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"duration": schema.StringAttribute{
						Description:         "Duration represents the duration of the chaos action.",
						MarkdownDescription: "Duration represents the duration of the chaos action.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ec2_instance": schema.StringAttribute{
						Description:         "Ec2Instance indicates the ID of the ec2 instance.",
						MarkdownDescription: "Ec2Instance indicates the ID of the ec2 instance.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"endpoint": schema.StringAttribute{
						Description:         "Endpoint indicates the endpoint of the aws server. Just used it in test now.",
						MarkdownDescription: "Endpoint indicates the endpoint of the aws server. Just used it in test now.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"remote_cluster": schema.StringAttribute{
						Description:         "RemoteCluster represents the remote cluster where the chaos will be deployed",
						MarkdownDescription: "RemoteCluster represents the remote cluster where the chaos will be deployed",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"secret_name": schema.StringAttribute{
						Description:         "SecretName defines the name of kubernetes secret.",
						MarkdownDescription: "SecretName defines the name of kubernetes secret.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"volume_id": schema.StringAttribute{
						Description:         "EbsVolume indicates the ID of the EBS volume. Needed in detach-volume.",
						MarkdownDescription: "EbsVolume indicates the ID of the EBS volume. Needed in detach-volume.",
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

func (r *ChaosMeshOrgAwschaosV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_chaos_mesh_org_aws_chaos_v1alpha1_manifest")

	var model ChaosMeshOrgAwschaosV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("chaos-mesh.org/v1alpha1")
	model.Kind = pointer.String("AWSChaos")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
