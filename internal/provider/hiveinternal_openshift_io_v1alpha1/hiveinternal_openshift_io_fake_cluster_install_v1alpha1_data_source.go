/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package hiveinternal_openshift_io_v1alpha1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"net/http"
)

var (
	_ datasource.DataSource              = &HiveinternalOpenshiftIoFakeClusterInstallV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &HiveinternalOpenshiftIoFakeClusterInstallV1Alpha1DataSource{}
)

func NewHiveinternalOpenshiftIoFakeClusterInstallV1Alpha1DataSource() datasource.DataSource {
	return &HiveinternalOpenshiftIoFakeClusterInstallV1Alpha1DataSource{}
}

type HiveinternalOpenshiftIoFakeClusterInstallV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type HiveinternalOpenshiftIoFakeClusterInstallV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

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

func (r *HiveinternalOpenshiftIoFakeClusterInstallV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_hiveinternal_openshift_io_fake_cluster_install_v1alpha1"
}

func (r *HiveinternalOpenshiftIoFakeClusterInstallV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
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
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
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
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"admin_password_secret_ref": schema.SingleNestedAttribute{
								Description:         "AdminPasswordSecretRef references the secret containing the admin username/password which can be used to login to this cluster.",
								MarkdownDescription: "AdminPasswordSecretRef references the secret containing the admin username/password which can be used to login to this cluster.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"cluster_id": schema.StringAttribute{
								Description:         "ClusterID is a globally unique identifier for this cluster generated during installation. Used for reporting metrics among other places.",
								MarkdownDescription: "ClusterID is a globally unique identifier for this cluster generated during installation. Used for reporting metrics among other places.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"infra_id": schema.StringAttribute{
								Description:         "InfraID is an identifier for this cluster generated during installation and used for tagging/naming resources in cloud providers.",
								MarkdownDescription: "InfraID is an identifier for this cluster generated during installation and used for tagging/naming resources in cloud providers.",
								Required:            false,
								Optional:            false,
								Computed:            true,
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
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"azure": schema.SingleNestedAttribute{
										Description:         "Azure holds azure-specific cluster metadata",
										MarkdownDescription: "Azure holds azure-specific cluster metadata",
										Attributes: map[string]schema.Attribute{
											"resource_group_name": schema.StringAttribute{
												Description:         "ResourceGroupName is the name of the resource group in which the cluster resources were created.",
												MarkdownDescription: "ResourceGroupName is the name of the resource group in which the cluster resources were created.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"image_set_ref": schema.SingleNestedAttribute{
						Description:         "ImageSetRef is a reference to a ClusterImageSet. The release image specified in the ClusterImageSet will be used to install the cluster.",
						MarkdownDescription: "ImageSetRef is a reference to a ClusterImageSet. The release image specified in the ClusterImageSet will be used to install the cluster.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name is the name of the ClusterImageSet that this refers to",
								MarkdownDescription: "Name is the name of the ClusterImageSet that this refers to",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *HiveinternalOpenshiftIoFakeClusterInstallV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *HiveinternalOpenshiftIoFakeClusterInstallV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_hiveinternal_openshift_io_fake_cluster_install_v1alpha1")

	var data HiveinternalOpenshiftIoFakeClusterInstallV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "hiveinternal.openshift.io", Version: "v1alpha1", Resource: "fakeclusterinstalls"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		var statusError *k8sErrors.StatusError
		if errors.As(err, &statusError) {
			if statusError.Status().Code == http.StatusNotFound {
				response.Diagnostics.AddError(
					"Unable to find resource",
					fmt.Sprintf("The requested resource cannot be found. "+
						"Make sure that it does exist in your cluster and you have set the correct name and namespace configured.\n\n"+
						"Namespace: %s\n"+
						"Name: %s", data.Metadata.Namespace, data.Metadata.Name),
				)
				return
			}
		} else {
			response.Diagnostics.AddError(
				"Unable to GET resource",
				fmt.Sprintf("An unexpected error occurred while reading the resource. "+
					"Please report this issue to the provider developers.\n\n"+
					"GET Error (%T): %s", err, err.Error()),
			)
		}
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse HiveinternalOpenshiftIoFakeClusterInstallV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("hiveinternal.openshift.io/v1alpha1")
	data.Kind = pointer.String("FakeClusterInstall")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
