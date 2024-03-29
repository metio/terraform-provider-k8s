/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package upgrade_managed_openshift_io_v1alpha1

import (
	"context"
	"fmt"
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
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &UpgradeManagedOpenshiftIoUpgradeConfigV1Alpha1Manifest{}
)

func NewUpgradeManagedOpenshiftIoUpgradeConfigV1Alpha1Manifest() datasource.DataSource {
	return &UpgradeManagedOpenshiftIoUpgradeConfigV1Alpha1Manifest{}
}

type UpgradeManagedOpenshiftIoUpgradeConfigV1Alpha1Manifest struct{}

type UpgradeManagedOpenshiftIoUpgradeConfigV1Alpha1ManifestData struct {
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
		PDBForceDrainTimeout *int64 `tfsdk:"pdb_force_drain_timeout" json:"PDBForceDrainTimeout,omitempty"`
		CapacityReservation  *bool  `tfsdk:"capacity_reservation" json:"capacityReservation,omitempty"`
		Desired              *struct {
			Channel *string `tfsdk:"channel" json:"channel,omitempty"`
			Image   *string `tfsdk:"image" json:"image,omitempty"`
			Version *string `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"desired" json:"desired,omitempty"`
		Type      *string `tfsdk:"type" json:"type,omitempty"`
		UpgradeAt *string `tfsdk:"upgrade_at" json:"upgradeAt,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *UpgradeManagedOpenshiftIoUpgradeConfigV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_upgrade_managed_openshift_io_upgrade_config_v1alpha1_manifest"
}

func (r *UpgradeManagedOpenshiftIoUpgradeConfigV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "UpgradeConfig is the Schema for the upgradeconfigs API",
		MarkdownDescription: "UpgradeConfig is the Schema for the upgradeconfigs API",
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
				Description:         "UpgradeConfigSpec defines the desired state of UpgradeConfig and upgrade window and freeze window",
				MarkdownDescription: "UpgradeConfigSpec defines the desired state of UpgradeConfig and upgrade window and freeze window",
				Attributes: map[string]schema.Attribute{
					"pdb_force_drain_timeout": schema.Int64Attribute{
						Description:         "The maximum grace period granted to a node whose drain is blocked by a Pod Disruption Budget, before that drain is forced. Measured in minutes. The minimum accepted value is 0 and in this case it will trigger force drain after the expectedNodeDrainTime lapsed.",
						MarkdownDescription: "The maximum grace period granted to a node whose drain is blocked by a Pod Disruption Budget, before that drain is forced. Measured in minutes. The minimum accepted value is 0 and in this case it will trigger force drain after the expectedNodeDrainTime lapsed.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},

					"capacity_reservation": schema.BoolAttribute{
						Description:         "Specify if scaling up an extra node for capacity reservation before upgrade starts is needed",
						MarkdownDescription: "Specify if scaling up an extra node for capacity reservation before upgrade starts is needed",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"desired": schema.SingleNestedAttribute{
						Description:         "Specify the desired OpenShift release",
						MarkdownDescription: "Specify the desired OpenShift release",
						Attributes: map[string]schema.Attribute{
							"channel": schema.StringAttribute{
								Description:         "Channel used for upgrades",
								MarkdownDescription: "Channel used for upgrades",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image": schema.StringAttribute{
								Description:         "Image reference used for upgrades",
								MarkdownDescription: "Image reference used for upgrades",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"version": schema.StringAttribute{
								Description:         "Version of openshift release",
								MarkdownDescription: "Version of openshift release",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"type": schema.StringAttribute{
						Description:         "Type indicates the ClusterUpgrader implementation to use to perform an upgrade of the cluster",
						MarkdownDescription: "Type indicates the ClusterUpgrader implementation to use to perform an upgrade of the cluster",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("OSD", "ARO"),
						},
					},

					"upgrade_at": schema.StringAttribute{
						Description:         "Specify the upgrade start time",
						MarkdownDescription: "Specify the upgrade start time",
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

func (r *UpgradeManagedOpenshiftIoUpgradeConfigV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_upgrade_managed_openshift_io_upgrade_config_v1alpha1_manifest")

	var model UpgradeManagedOpenshiftIoUpgradeConfigV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("upgrade.managed.openshift.io/v1alpha1")
	model.Kind = pointer.String("UpgradeConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
