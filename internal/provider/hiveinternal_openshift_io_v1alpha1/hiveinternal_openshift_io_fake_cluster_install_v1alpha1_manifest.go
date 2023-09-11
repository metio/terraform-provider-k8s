/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package hiveinternal_openshift_io_v1alpha1

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
	_ datasource.DataSource = &HiveinternalOpenshiftIoFakeClusterInstallV1Alpha1Manifest{}
)

func NewHiveinternalOpenshiftIoFakeClusterInstallV1Alpha1Manifest() datasource.DataSource {
	return &HiveinternalOpenshiftIoFakeClusterInstallV1Alpha1Manifest{}
}

type HiveinternalOpenshiftIoFakeClusterInstallV1Alpha1Manifest struct{}

type HiveinternalOpenshiftIoFakeClusterInstallV1Alpha1ManifestData struct {
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
		ClusterDeploymentRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"cluster_deployment_ref" json:"clusterDeploymentRef,omitempty"`
		ClusterMetadata *struct {
			AdminKubeconfigSecretRef *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"admin_kubeconfig_secret_ref" json:"adminKubeconfigSecretRef,omitempty"`
			AdminPasswordSecretRef *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"admin_password_secret_ref" json:"adminPasswordSecretRef,omitempty"`
			ClusterID *string `tfsdk:"cluster_id" json:"clusterID,omitempty"`
			InfraID   *string `tfsdk:"infra_id" json:"infraID,omitempty"`
			Platform  *struct {
				Aws *struct {
					HostedZoneRole *string `tfsdk:"hosted_zone_role" json:"hostedZoneRole,omitempty"`
				} `tfsdk:"aws" json:"aws,omitempty"`
				Azure *struct {
					ResourceGroupName *string `tfsdk:"resource_group_name" json:"resourceGroupName,omitempty"`
				} `tfsdk:"azure" json:"azure,omitempty"`
			} `tfsdk:"platform" json:"platform,omitempty"`
		} `tfsdk:"cluster_metadata" json:"clusterMetadata,omitempty"`
		ImageSetRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"image_set_ref" json:"imageSetRef,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *HiveinternalOpenshiftIoFakeClusterInstallV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_hiveinternal_openshift_io_fake_cluster_install_v1alpha1_manifest"
}

func (r *HiveinternalOpenshiftIoFakeClusterInstallV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "FakeClusterInstall represents a fake request to provision an agent based cluster.",
		MarkdownDescription: "FakeClusterInstall represents a fake request to provision an agent based cluster.",
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
				Description:         "FakeClusterInstallSpec defines the desired state of the FakeClusterInstall.",
				MarkdownDescription: "FakeClusterInstallSpec defines the desired state of the FakeClusterInstall.",
				Attributes: map[string]schema.Attribute{
					"cluster_deployment_ref": schema.SingleNestedAttribute{
						Description:         "ClusterDeploymentRef is a reference to the ClusterDeployment associated with this AgentClusterInstall.",
						MarkdownDescription: "ClusterDeploymentRef is a reference to the ClusterDeployment associated with this AgentClusterInstall.",
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

					"cluster_metadata": schema.SingleNestedAttribute{
						Description:         "ClusterMetadata contains metadata information about the installed cluster. It should be populated once the cluster install is completed. (it can be populated sooner if desired, but Hive will not copy back to ClusterDeployment until the Installed condition goes True.",
						MarkdownDescription: "ClusterMetadata contains metadata information about the installed cluster. It should be populated once the cluster install is completed. (it can be populated sooner if desired, but Hive will not copy back to ClusterDeployment until the Installed condition goes True.",
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
								Required: true,
								Optional: false,
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

							"cluster_id": schema.StringAttribute{
								Description:         "ClusterID is a globally unique identifier for this cluster generated during installation. Used for reporting metrics among other places.",
								MarkdownDescription: "ClusterID is a globally unique identifier for this cluster generated during installation. Used for reporting metrics among other places.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"infra_id": schema.StringAttribute{
								Description:         "InfraID is an identifier for this cluster generated during installation and used for tagging/naming resources in cloud providers.",
								MarkdownDescription: "InfraID is an identifier for this cluster generated during installation and used for tagging/naming resources in cloud providers.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"platform": schema.SingleNestedAttribute{
								Description:         "Platform holds platform-specific cluster metadata",
								MarkdownDescription: "Platform holds platform-specific cluster metadata",
								Attributes: map[string]schema.Attribute{
									"aws": schema.SingleNestedAttribute{
										Description:         "AWS holds AWS-specific cluster metadata",
										MarkdownDescription: "AWS holds AWS-specific cluster metadata",
										Attributes: map[string]schema.Attribute{
											"hosted_zone_role": schema.StringAttribute{
												Description:         "HostedZoneRole is the role to assume when performing operations on a hosted zone owned by another account.",
												MarkdownDescription: "HostedZoneRole is the role to assume when performing operations on a hosted zone owned by another account.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"azure": schema.SingleNestedAttribute{
										Description:         "Azure holds azure-specific cluster metadata",
										MarkdownDescription: "Azure holds azure-specific cluster metadata",
										Attributes: map[string]schema.Attribute{
											"resource_group_name": schema.StringAttribute{
												Description:         "ResourceGroupName is the name of the resource group in which the cluster resources were created.",
												MarkdownDescription: "ResourceGroupName is the name of the resource group in which the cluster resources were created.",
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
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"image_set_ref": schema.SingleNestedAttribute{
						Description:         "ImageSetRef is a reference to a ClusterImageSet. The release image specified in the ClusterImageSet will be used to install the cluster.",
						MarkdownDescription: "ImageSetRef is a reference to a ClusterImageSet. The release image specified in the ClusterImageSet will be used to install the cluster.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name is the name of the ClusterImageSet that this refers to",
								MarkdownDescription: "Name is the name of the ClusterImageSet that this refers to",
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
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *HiveinternalOpenshiftIoFakeClusterInstallV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_hiveinternal_openshift_io_fake_cluster_install_v1alpha1_manifest")

	var model HiveinternalOpenshiftIoFakeClusterInstallV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("hiveinternal.openshift.io/v1alpha1")
	model.Kind = pointer.String("FakeClusterInstall")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
