/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package infrastructure_cluster_x_k8s_io_v1beta1

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
	_ datasource.DataSource              = &InfrastructureClusterXK8SIoIbmpowerVsmachineV1Beta1DataSource{}
	_ datasource.DataSourceWithConfigure = &InfrastructureClusterXK8SIoIbmpowerVsmachineV1Beta1DataSource{}
)

func NewInfrastructureClusterXK8SIoIbmpowerVsmachineV1Beta1DataSource() datasource.DataSource {
	return &InfrastructureClusterXK8SIoIbmpowerVsmachineV1Beta1DataSource{}
}

type InfrastructureClusterXK8SIoIbmpowerVsmachineV1Beta1DataSource struct {
	kubernetesClient dynamic.Interface
}

type InfrastructureClusterXK8SIoIbmpowerVsmachineV1Beta1DataSourceData struct {
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
		Image *struct {
			Id    *string `tfsdk:"id" json:"id,omitempty"`
			Name  *string `tfsdk:"name" json:"name,omitempty"`
			Regex *string `tfsdk:"regex" json:"regex,omitempty"`
		} `tfsdk:"image" json:"image,omitempty"`
		ImageRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"image_ref" json:"imageRef,omitempty"`
		Memory  *string `tfsdk:"memory" json:"memory,omitempty"`
		Network *struct {
			Id    *string `tfsdk:"id" json:"id,omitempty"`
			Name  *string `tfsdk:"name" json:"name,omitempty"`
			Regex *string `tfsdk:"regex" json:"regex,omitempty"`
		} `tfsdk:"network" json:"network,omitempty"`
		ProcType          *string `tfsdk:"proc_type" json:"procType,omitempty"`
		Processors        *string `tfsdk:"processors" json:"processors,omitempty"`
		ProviderID        *string `tfsdk:"provider_id" json:"providerID,omitempty"`
		ServiceInstanceID *string `tfsdk:"service_instance_id" json:"serviceInstanceID,omitempty"`
		SshKey            *string `tfsdk:"ssh_key" json:"sshKey,omitempty"`
		SysType           *string `tfsdk:"sys_type" json:"sysType,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *InfrastructureClusterXK8SIoIbmpowerVsmachineV1Beta1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_infrastructure_cluster_x_k8s_io_ibm_power_vs_machine_v1beta1"
}

func (r *InfrastructureClusterXK8SIoIbmpowerVsmachineV1Beta1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "IBMPowerVSMachine is the Schema for the ibmpowervsmachines API.",
		MarkdownDescription: "IBMPowerVSMachine is the Schema for the ibmpowervsmachines API.",
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
				Description:         "IBMPowerVSMachineSpec defines the desired state of IBMPowerVSMachine.",
				MarkdownDescription: "IBMPowerVSMachineSpec defines the desired state of IBMPowerVSMachine.",
				Attributes: map[string]schema.Attribute{
					"image": schema.SingleNestedAttribute{
						Description:         "Image is the reference to the Image from which to create the machine instance.",
						MarkdownDescription: "Image is the reference to the Image from which to create the machine instance.",
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Description:         "ID of resource",
								MarkdownDescription: "ID of resource",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "Name of resource",
								MarkdownDescription: "Name of resource",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"regex": schema.StringAttribute{
								Description:         "Regular expression to match resource, In case of multiple resources matches the provided regular expression the first matched resource will be selected",
								MarkdownDescription: "Regular expression to match resource, In case of multiple resources matches the provided regular expression the first matched resource will be selected",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"image_ref": schema.SingleNestedAttribute{
						Description:         "ImageRef is an optional reference to a provider-specific resource that holds the details for provisioning the Image for a Cluster.",
						MarkdownDescription: "ImageRef is an optional reference to a provider-specific resource that holds the details for provisioning the Image for a Cluster.",
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

					"memory": schema.StringAttribute{
						Description:         "Memory is Amount of memory allocated (in GB)",
						MarkdownDescription: "Memory is Amount of memory allocated (in GB)",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"network": schema.SingleNestedAttribute{
						Description:         "Network is the reference to the Network to use for this instance.",
						MarkdownDescription: "Network is the reference to the Network to use for this instance.",
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Description:         "ID of resource",
								MarkdownDescription: "ID of resource",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "Name of resource",
								MarkdownDescription: "Name of resource",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"regex": schema.StringAttribute{
								Description:         "Regular expression to match resource, In case of multiple resources matches the provided regular expression the first matched resource will be selected",
								MarkdownDescription: "Regular expression to match resource, In case of multiple resources matches the provided regular expression the first matched resource will be selected",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"proc_type": schema.StringAttribute{
						Description:         "ProcType is the processor type, e.g: dedicated, shared, capped",
						MarkdownDescription: "ProcType is the processor type, e.g: dedicated, shared, capped",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"processors": schema.StringAttribute{
						Description:         "Processors is Number of processors allocated.",
						MarkdownDescription: "Processors is Number of processors allocated.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"provider_id": schema.StringAttribute{
						Description:         "ProviderID is the unique identifier as specified by the cloud provider.",
						MarkdownDescription: "ProviderID is the unique identifier as specified by the cloud provider.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"service_instance_id": schema.StringAttribute{
						Description:         "ServiceInstanceID is the id of the power cloud instance where the vsi instance will get deployed.",
						MarkdownDescription: "ServiceInstanceID is the id of the power cloud instance where the vsi instance will get deployed.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"ssh_key": schema.StringAttribute{
						Description:         "SSHKey is the name of the SSH key pair provided to the vsi for authenticating users.",
						MarkdownDescription: "SSHKey is the name of the SSH key pair provided to the vsi for authenticating users.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"sys_type": schema.StringAttribute{
						Description:         "SysType is the System type used to host the vsi.",
						MarkdownDescription: "SysType is the System type used to host the vsi.",
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

func (r *InfrastructureClusterXK8SIoIbmpowerVsmachineV1Beta1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *InfrastructureClusterXK8SIoIbmpowerVsmachineV1Beta1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_infrastructure_cluster_x_k8s_io_ibm_power_vs_machine_v1beta1")

	var data InfrastructureClusterXK8SIoIbmpowerVsmachineV1Beta1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "infrastructure.cluster.x-k8s.io", Version: "v1beta1", Resource: "ibmpowervsmachines"}).
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

	var readResponse InfrastructureClusterXK8SIoIbmpowerVsmachineV1Beta1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("infrastructure.cluster.x-k8s.io/v1beta1")
	data.Kind = pointer.String("IBMPowerVSMachine")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
