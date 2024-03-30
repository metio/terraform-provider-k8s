/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package hive_openshift_io_v1

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
	_ datasource.DataSource = &HiveOpenshiftIoClusterProvisionV1Manifest{}
)

func NewHiveOpenshiftIoClusterProvisionV1Manifest() datasource.DataSource {
	return &HiveOpenshiftIoClusterProvisionV1Manifest{}
}

type HiveOpenshiftIoClusterProvisionV1Manifest struct{}

type HiveOpenshiftIoClusterProvisionV1ManifestData struct {
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
		AdminKubeconfigSecretRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"admin_kubeconfig_secret_ref" json:"adminKubeconfigSecretRef,omitempty"`
		AdminPasswordSecretRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"admin_password_secret_ref" json:"adminPasswordSecretRef,omitempty"`
		Attempt              *int64 `tfsdk:"attempt" json:"attempt,omitempty"`
		ClusterDeploymentRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"cluster_deployment_ref" json:"clusterDeploymentRef,omitempty"`
		ClusterID         *string            `tfsdk:"cluster_id" json:"clusterID,omitempty"`
		InfraID           *string            `tfsdk:"infra_id" json:"infraID,omitempty"`
		InstallLog        *string            `tfsdk:"install_log" json:"installLog,omitempty"`
		Metadata          *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
		MetadataJSON      *string            `tfsdk:"metadata_json" json:"metadataJSON,omitempty"`
		PrevClusterID     *string            `tfsdk:"prev_cluster_id" json:"prevClusterID,omitempty"`
		PrevInfraID       *string            `tfsdk:"prev_infra_id" json:"prevInfraID,omitempty"`
		PrevProvisionName *string            `tfsdk:"prev_provision_name" json:"prevProvisionName,omitempty"`
		Stage             *string            `tfsdk:"stage" json:"stage,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *HiveOpenshiftIoClusterProvisionV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_hive_openshift_io_cluster_provision_v1_manifest"
}

func (r *HiveOpenshiftIoClusterProvisionV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClusterProvision is the Schema for the clusterprovisions API",
		MarkdownDescription: "ClusterProvision is the Schema for the clusterprovisions API",
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
				Description:         "ClusterProvisionSpec defines the results of provisioning a cluster.",
				MarkdownDescription: "ClusterProvisionSpec defines the results of provisioning a cluster.",
				Attributes: map[string]schema.Attribute{
					"admin_kubeconfig_secret_ref": schema.SingleNestedAttribute{
						Description:         "AdminKubeconfigSecretRef references the secret containing the admin kubeconfig for this cluster.",
						MarkdownDescription: "AdminKubeconfigSecretRef references the secret containing the admin kubeconfig for this cluster.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"admin_password_secret_ref": schema.SingleNestedAttribute{
						Description:         "AdminPasswordSecretRef references the secret containing the admin username/password which can be used to login to this cluster.",
						MarkdownDescription: "AdminPasswordSecretRef references the secret containing the admin username/password which can be used to login to this cluster.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"attempt": schema.Int64Attribute{
						Description:         "Attempt is which attempt number of the cluster deployment that this ClusterProvision is",
						MarkdownDescription: "Attempt is which attempt number of the cluster deployment that this ClusterProvision is",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"cluster_deployment_ref": schema.SingleNestedAttribute{
						Description:         "ClusterDeploymentRef references the cluster deployment provisioned.",
						MarkdownDescription: "ClusterDeploymentRef references the cluster deployment provisioned.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"cluster_id": schema.StringAttribute{
						Description:         "ClusterID is a globally unique identifier for this cluster generated during installation. Used for reporting metrics among other places.",
						MarkdownDescription: "ClusterID is a globally unique identifier for this cluster generated during installation. Used for reporting metrics among other places.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"infra_id": schema.StringAttribute{
						Description:         "InfraID is an identifier for this cluster generated during installation and used for tagging/naming resources in cloud providers.",
						MarkdownDescription: "InfraID is an identifier for this cluster generated during installation and used for tagging/naming resources in cloud providers.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"install_log": schema.StringAttribute{
						Description:         "InstallLog is the log from the installer.",
						MarkdownDescription: "InstallLog is the log from the installer.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"metadata": schema.MapAttribute{
						Description:         "Metadata is the metadata.json generated by the installer, providing metadata information about the cluster created. NOTE: This is not used because it didn't work (it was always empty). We think because the thing it's storing (ClusterMetadata from installer) is not a runtime.Object, so can't be put in a RawExtension.",
						MarkdownDescription: "Metadata is the metadata.json generated by the installer, providing metadata information about the cluster created. NOTE: This is not used because it didn't work (it was always empty). We think because the thing it's storing (ClusterMetadata from installer) is not a runtime.Object, so can't be put in a RawExtension.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"metadata_json": schema.StringAttribute{
						Description:         "MetadataJSON is a JSON representation of the ClusterMetadata produced by the installer. We don't use a runtime.RawExtension because ClusterMetadata isn't a runtime.Object. We don't use ClusterMetadata itself because we don't want our API consumers to need to pull in the installer code and its dependencies.",
						MarkdownDescription: "MetadataJSON is a JSON representation of the ClusterMetadata produced by the installer. We don't use a runtime.RawExtension because ClusterMetadata isn't a runtime.Object. We don't use ClusterMetadata itself because we don't want our API consumers to need to pull in the installer code and its dependencies.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							validators.Base64Validator(),
						},
					},

					"prev_cluster_id": schema.StringAttribute{
						Description:         "PrevClusterID is the cluster ID of the previous failed provision attempt.",
						MarkdownDescription: "PrevClusterID is the cluster ID of the previous failed provision attempt.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"prev_infra_id": schema.StringAttribute{
						Description:         "PrevInfraID is the infra ID of the previous failed provision attempt.",
						MarkdownDescription: "PrevInfraID is the infra ID of the previous failed provision attempt.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"prev_provision_name": schema.StringAttribute{
						Description:         "PrevProvisionName is the name of the previous failed provision attempt.",
						MarkdownDescription: "PrevProvisionName is the name of the previous failed provision attempt.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"stage": schema.StringAttribute{
						Description:         "Stage is the stage of provisioning that the cluster deployment has reached.",
						MarkdownDescription: "Stage is the stage of provisioning that the cluster deployment has reached.",
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

func (r *HiveOpenshiftIoClusterProvisionV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_hive_openshift_io_cluster_provision_v1_manifest")

	var model HiveOpenshiftIoClusterProvisionV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("hive.openshift.io/v1")
	model.Kind = pointer.String("ClusterProvision")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
