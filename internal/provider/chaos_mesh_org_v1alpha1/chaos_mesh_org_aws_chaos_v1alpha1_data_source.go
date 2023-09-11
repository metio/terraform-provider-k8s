/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package chaos_mesh_org_v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &ChaosMeshOrgAwschaosV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &ChaosMeshOrgAwschaosV1Alpha1DataSource{}
)

func NewChaosMeshOrgAwschaosV1Alpha1DataSource() datasource.DataSource {
	return &ChaosMeshOrgAwschaosV1Alpha1DataSource{}
}

type ChaosMeshOrgAwschaosV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type ChaosMeshOrgAwschaosV1Alpha1DataSourceData struct {
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

func (r *ChaosMeshOrgAwschaosV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_chaos_mesh_org_aws_chaos_v1alpha1"
}

func (r *ChaosMeshOrgAwschaosV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "AWSChaos is the Schema for the awschaos API",
		MarkdownDescription: "AWSChaos is the Schema for the awschaos API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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
				Description:         "AWSChaosSpec is the content of the specification for an AWSChaos",
				MarkdownDescription: "AWSChaosSpec is the content of the specification for an AWSChaos",
				Attributes: map[string]schema.Attribute{
					"action": schema.StringAttribute{
						Description:         "Action defines the specific aws chaos action. Supported action: ec2-stop / ec2-restart / detach-volume Default action: ec2-stop",
						MarkdownDescription: "Action defines the specific aws chaos action. Supported action: ec2-stop / ec2-restart / detach-volume Default action: ec2-stop",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"aws_region": schema.StringAttribute{
						Description:         "AWSRegion defines the region of aws.",
						MarkdownDescription: "AWSRegion defines the region of aws.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"device_name": schema.StringAttribute{
						Description:         "DeviceName indicates the name of the device. Needed in detach-volume.",
						MarkdownDescription: "DeviceName indicates the name of the device. Needed in detach-volume.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"duration": schema.StringAttribute{
						Description:         "Duration represents the duration of the chaos action.",
						MarkdownDescription: "Duration represents the duration of the chaos action.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"ec2_instance": schema.StringAttribute{
						Description:         "Ec2Instance indicates the ID of the ec2 instance.",
						MarkdownDescription: "Ec2Instance indicates the ID of the ec2 instance.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"endpoint": schema.StringAttribute{
						Description:         "Endpoint indicates the endpoint of the aws server. Just used it in test now.",
						MarkdownDescription: "Endpoint indicates the endpoint of the aws server. Just used it in test now.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"remote_cluster": schema.StringAttribute{
						Description:         "RemoteCluster represents the remote cluster where the chaos will be deployed",
						MarkdownDescription: "RemoteCluster represents the remote cluster where the chaos will be deployed",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"secret_name": schema.StringAttribute{
						Description:         "SecretName defines the name of kubernetes secret.",
						MarkdownDescription: "SecretName defines the name of kubernetes secret.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"volume_id": schema.StringAttribute{
						Description:         "EbsVolume indicates the ID of the EBS volume. Needed in detach-volume.",
						MarkdownDescription: "EbsVolume indicates the ID of the EBS volume. Needed in detach-volume.",
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
	}
}

func (r *ChaosMeshOrgAwschaosV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *ChaosMeshOrgAwschaosV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_chaos_mesh_org_aws_chaos_v1alpha1")

	var data ChaosMeshOrgAwschaosV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "chaos-mesh.org", Version: "v1alpha1", Resource: "awschaos"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse ChaosMeshOrgAwschaosV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("chaos-mesh.org/v1alpha1")
	data.Kind = pointer.String("AWSChaos")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
