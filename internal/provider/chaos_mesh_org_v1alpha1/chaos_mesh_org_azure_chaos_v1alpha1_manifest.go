/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package chaos_mesh_org_v1alpha1

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
	_ datasource.DataSource = &ChaosMeshOrgAzureChaosV1Alpha1Manifest{}
)

func NewChaosMeshOrgAzureChaosV1Alpha1Manifest() datasource.DataSource {
	return &ChaosMeshOrgAzureChaosV1Alpha1Manifest{}
}

type ChaosMeshOrgAzureChaosV1Alpha1Manifest struct{}

type ChaosMeshOrgAzureChaosV1Alpha1ManifestData struct {
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
		Action            *string `tfsdk:"action" json:"action,omitempty"`
		DiskName          *string `tfsdk:"disk_name" json:"diskName,omitempty"`
		Duration          *string `tfsdk:"duration" json:"duration,omitempty"`
		Lun               *int64  `tfsdk:"lun" json:"lun,omitempty"`
		RemoteCluster     *string `tfsdk:"remote_cluster" json:"remoteCluster,omitempty"`
		ResourceGroupName *string `tfsdk:"resource_group_name" json:"resourceGroupName,omitempty"`
		SecretName        *string `tfsdk:"secret_name" json:"secretName,omitempty"`
		SubscriptionID    *string `tfsdk:"subscription_id" json:"subscriptionID,omitempty"`
		VmName            *string `tfsdk:"vm_name" json:"vmName,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ChaosMeshOrgAzureChaosV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_chaos_mesh_org_azure_chaos_v1alpha1_manifest"
}

func (r *ChaosMeshOrgAzureChaosV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "AzureChaos is the Schema for the azurechaos API",
		MarkdownDescription: "AzureChaos is the Schema for the azurechaos API",
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
				Description:         "AzureChaosSpec is the content of the specification for an AzureChaos",
				MarkdownDescription: "AzureChaosSpec is the content of the specification for an AzureChaos",
				Attributes: map[string]schema.Attribute{
					"action": schema.StringAttribute{
						Description:         "Action defines the specific azure chaos action. Supported action: vm-stop / vm-restart / disk-detach Default action: vm-stop",
						MarkdownDescription: "Action defines the specific azure chaos action. Supported action: vm-stop / vm-restart / disk-detach Default action: vm-stop",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("vm-stop", "vm-restart", "disk-detach"),
						},
					},

					"disk_name": schema.StringAttribute{
						Description:         "DiskName indicates the name of the disk. Needed in disk-detach.",
						MarkdownDescription: "DiskName indicates the name of the disk. Needed in disk-detach.",
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

					"lun": schema.Int64Attribute{
						Description:         "LUN indicates the Logical Unit Number of the data disk. Needed in disk-detach.",
						MarkdownDescription: "LUN indicates the Logical Unit Number of the data disk. Needed in disk-detach.",
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

					"resource_group_name": schema.StringAttribute{
						Description:         "ResourceGroupName defines the name of ResourceGroup",
						MarkdownDescription: "ResourceGroupName defines the name of ResourceGroup",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"secret_name": schema.StringAttribute{
						Description:         "SecretName defines the name of kubernetes secret. It is used for Azure credentials.",
						MarkdownDescription: "SecretName defines the name of kubernetes secret. It is used for Azure credentials.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"subscription_id": schema.StringAttribute{
						Description:         "SubscriptionID defines the id of Azure subscription.",
						MarkdownDescription: "SubscriptionID defines the id of Azure subscription.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"vm_name": schema.StringAttribute{
						Description:         "VMName defines the name of Virtual Machine",
						MarkdownDescription: "VMName defines the name of Virtual Machine",
						Required:            true,
						Optional:            false,
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

func (r *ChaosMeshOrgAzureChaosV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_chaos_mesh_org_azure_chaos_v1alpha1_manifest")

	var model ChaosMeshOrgAzureChaosV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("chaos-mesh.org/v1alpha1")
	model.Kind = pointer.String("AzureChaos")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
