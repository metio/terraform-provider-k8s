/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package infrastructure_cluster_x_k8s_io_v1beta2

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
	_ datasource.DataSource = &InfrastructureClusterXK8SIoIbmpowerVsmachineV1Beta2Manifest{}
)

func NewInfrastructureClusterXK8SIoIbmpowerVsmachineV1Beta2Manifest() datasource.DataSource {
	return &InfrastructureClusterXK8SIoIbmpowerVsmachineV1Beta2Manifest{}
}

type InfrastructureClusterXK8SIoIbmpowerVsmachineV1Beta2Manifest struct{}

type InfrastructureClusterXK8SIoIbmpowerVsmachineV1Beta2ManifestData struct {
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
		Image *struct {
			Id    *string `tfsdk:"id" json:"id,omitempty"`
			Name  *string `tfsdk:"name" json:"name,omitempty"`
			Regex *string `tfsdk:"regex" json:"regex,omitempty"`
		} `tfsdk:"image" json:"image,omitempty"`
		ImageRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"image_ref" json:"imageRef,omitempty"`
		MemoryGiB *int64 `tfsdk:"memory_gi_b" json:"memoryGiB,omitempty"`
		Network   *struct {
			Id    *string `tfsdk:"id" json:"id,omitempty"`
			Name  *string `tfsdk:"name" json:"name,omitempty"`
			Regex *string `tfsdk:"regex" json:"regex,omitempty"`
		} `tfsdk:"network" json:"network,omitempty"`
		ProcessorType   *string `tfsdk:"processor_type" json:"processorType,omitempty"`
		Processors      *string `tfsdk:"processors" json:"processors,omitempty"`
		ProviderID      *string `tfsdk:"provider_id" json:"providerID,omitempty"`
		ServiceInstance *struct {
			Id    *string `tfsdk:"id" json:"id,omitempty"`
			Name  *string `tfsdk:"name" json:"name,omitempty"`
			Regex *string `tfsdk:"regex" json:"regex,omitempty"`
		} `tfsdk:"service_instance" json:"serviceInstance,omitempty"`
		ServiceInstanceID *string `tfsdk:"service_instance_id" json:"serviceInstanceID,omitempty"`
		SshKey            *string `tfsdk:"ssh_key" json:"sshKey,omitempty"`
		SystemType        *string `tfsdk:"system_type" json:"systemType,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *InfrastructureClusterXK8SIoIbmpowerVsmachineV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_infrastructure_cluster_x_k8s_io_ibm_power_vs_machine_v1beta2_manifest"
}

func (r *InfrastructureClusterXK8SIoIbmpowerVsmachineV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "IBMPowerVSMachine is the Schema for the ibmpowervsmachines API.",
		MarkdownDescription: "IBMPowerVSMachine is the Schema for the ibmpowervsmachines API.",
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
				Description:         "IBMPowerVSMachineSpec defines the desired state of IBMPowerVSMachine.",
				MarkdownDescription: "IBMPowerVSMachineSpec defines the desired state of IBMPowerVSMachine.",
				Attributes: map[string]schema.Attribute{
					"image": schema.SingleNestedAttribute{
						Description:         "Image the reference to the image which is used to create the instance.supported image identifier in IBMPowerVSResourceReference are Name and ID and that can be obtained from IBM Cloud UI or IBM Cloud cli.",
						MarkdownDescription: "Image the reference to the image which is used to create the instance.supported image identifier in IBMPowerVSResourceReference are Name and ID and that can be obtained from IBM Cloud UI or IBM Cloud cli.",
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Description:         "ID of resource",
								MarkdownDescription: "ID of resource",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},

							"name": schema.StringAttribute{
								Description:         "Name of resource",
								MarkdownDescription: "Name of resource",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},

							"regex": schema.StringAttribute{
								Description:         "Regular expression to match resource,In case of multiple resources matches the provided regular expression the first matched resource will be selected",
								MarkdownDescription: "Regular expression to match resource,In case of multiple resources matches the provided regular expression the first matched resource will be selected",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"image_ref": schema.SingleNestedAttribute{
						Description:         "ImageRef is an optional reference to a provider-specific resource that holdsthe details for provisioning the Image for a Cluster.",
						MarkdownDescription: "ImageRef is an optional reference to a provider-specific resource that holdsthe details for provisioning the Image for a Cluster.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
								MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"memory_gi_b": schema.Int64Attribute{
						Description:         "memoryGiB is the size of a virtual machine's memory, in GiB.maximum value for the MemoryGiB depends on the selected SystemType.when SystemType is set to e880 maximum MemoryGiB value is 7463 GiB.when SystemType is set to e980 maximum MemoryGiB value is 15307 GiB.when SystemType is set to s922 maximum MemoryGiB value is 942 GiB.The minimum memory is 2 GiB.When omitted, this means the user has no opinion and the platform is left to choose a reasonabledefault, which is subject to change over time. The current default is 2.",
						MarkdownDescription: "memoryGiB is the size of a virtual machine's memory, in GiB.maximum value for the MemoryGiB depends on the selected SystemType.when SystemType is set to e880 maximum MemoryGiB value is 7463 GiB.when SystemType is set to e980 maximum MemoryGiB value is 15307 GiB.when SystemType is set to s922 maximum MemoryGiB value is 942 GiB.The minimum memory is 2 GiB.When omitted, this means the user has no opinion and the platform is left to choose a reasonabledefault, which is subject to change over time. The current default is 2.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"network": schema.SingleNestedAttribute{
						Description:         "Network is the reference to the Network to use for this instance.supported network identifier in IBMPowerVSResourceReference are Name, ID and RegEx and that can be obtained from IBM Cloud UI or IBM Cloud cli.",
						MarkdownDescription: "Network is the reference to the Network to use for this instance.supported network identifier in IBMPowerVSResourceReference are Name, ID and RegEx and that can be obtained from IBM Cloud UI or IBM Cloud cli.",
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Description:         "ID of resource",
								MarkdownDescription: "ID of resource",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},

							"name": schema.StringAttribute{
								Description:         "Name of resource",
								MarkdownDescription: "Name of resource",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},

							"regex": schema.StringAttribute{
								Description:         "Regular expression to match resource,In case of multiple resources matches the provided regular expression the first matched resource will be selected",
								MarkdownDescription: "Regular expression to match resource,In case of multiple resources matches the provided regular expression the first matched resource will be selected",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"processor_type": schema.StringAttribute{
						Description:         "processorType is the VM instance processor type.It must be set to one of the following values: Dedicated, Capped or Shared.Dedicated: resources are allocated for a specific client, The hypervisor makes a 1:1 binding of a partition’s processor to a physical processor core.Shared: Shared among other clients.Capped: Shared, but resources do not expand beyond those that are requested, the amount of CPU time is Capped to the value specified for the entitlement.if the processorType is selected as Dedicated, then processors value cannot be fractional.When omitted, this means that the user has no opinion and the platform is left to choose areasonable default, which is subject to change over time. The current default is Shared.",
						MarkdownDescription: "processorType is the VM instance processor type.It must be set to one of the following values: Dedicated, Capped or Shared.Dedicated: resources are allocated for a specific client, The hypervisor makes a 1:1 binding of a partition’s processor to a physical processor core.Shared: Shared among other clients.Capped: Shared, but resources do not expand beyond those that are requested, the amount of CPU time is Capped to the value specified for the entitlement.if the processorType is selected as Dedicated, then processors value cannot be fractional.When omitted, this means that the user has no opinion and the platform is left to choose areasonable default, which is subject to change over time. The current default is Shared.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Dedicated", "Shared", "Capped", ""),
						},
					},

					"processors": schema.StringAttribute{
						Description:         "processors is the number of virtual processors in a virtual machine.when the processorType is selected as Dedicated the processors value cannot be fractional.maximum value for the Processors depends on the selected SystemType.when SystemType is set to e880 or e980 maximum Processors value is 143.when SystemType is set to s922 maximum Processors value is 15.minimum value for Processors depends on the selected ProcessorType.when ProcessorType is set as Shared or Capped, The minimum processors is 0.25.when ProcessorType is set as Dedicated, The minimum processors is 1.When omitted, this means that the user has no opinion and the platform is left to choose areasonable default, which is subject to change over time. The default is set based on the selected ProcessorType.when ProcessorType selected as Dedicated, the default is set to 1.when ProcessorType selected as Shared or Capped, the default is set to 0.25.",
						MarkdownDescription: "processors is the number of virtual processors in a virtual machine.when the processorType is selected as Dedicated the processors value cannot be fractional.maximum value for the Processors depends on the selected SystemType.when SystemType is set to e880 or e980 maximum Processors value is 143.when SystemType is set to s922 maximum Processors value is 15.minimum value for Processors depends on the selected ProcessorType.when ProcessorType is set as Shared or Capped, The minimum processors is 0.25.when ProcessorType is set as Dedicated, The minimum processors is 1.When omitted, this means that the user has no opinion and the platform is left to choose areasonable default, which is subject to change over time. The default is set based on the selected ProcessorType.when ProcessorType selected as Dedicated, the default is set to 1.when ProcessorType selected as Shared or Capped, the default is set to 0.25.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"provider_id": schema.StringAttribute{
						Description:         "ProviderID is the unique identifier as specified by the cloud provider.",
						MarkdownDescription: "ProviderID is the unique identifier as specified by the cloud provider.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service_instance": schema.SingleNestedAttribute{
						Description:         "serviceInstance is the reference to the Power VS workspace on which the server instance(VM) will be created.Power VS workspace is a container for all Power VS instances at a specific geographic region.serviceInstance can be created via IBM Cloud catalog or CLI.supported serviceInstance identifier in PowerVSResource are Name and ID and that can be obtained from IBM Cloud UI or IBM Cloud cli.More detail about Power VS service instance.https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-creating-power-virtual-serverwhen omitted system will dynamically create the service instance",
						MarkdownDescription: "serviceInstance is the reference to the Power VS workspace on which the server instance(VM) will be created.Power VS workspace is a container for all Power VS instances at a specific geographic region.serviceInstance can be created via IBM Cloud catalog or CLI.supported serviceInstance identifier in PowerVSResource are Name and ID and that can be obtained from IBM Cloud UI or IBM Cloud cli.More detail about Power VS service instance.https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-creating-power-virtual-serverwhen omitted system will dynamically create the service instance",
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Description:         "ID of resource",
								MarkdownDescription: "ID of resource",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},

							"name": schema.StringAttribute{
								Description:         "Name of resource",
								MarkdownDescription: "Name of resource",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},

							"regex": schema.StringAttribute{
								Description:         "Regular expression to match resource,In case of multiple resources matches the provided regular expression the first matched resource will be selected",
								MarkdownDescription: "Regular expression to match resource,In case of multiple resources matches the provided regular expression the first matched resource will be selected",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"service_instance_id": schema.StringAttribute{
						Description:         "ServiceInstanceID is the id of the power cloud instance where the vsi instance will get deployed.Deprecated: use ServiceInstance instead",
						MarkdownDescription: "ServiceInstanceID is the id of the power cloud instance where the vsi instance will get deployed.Deprecated: use ServiceInstance instead",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"ssh_key": schema.StringAttribute{
						Description:         "SSHKey is the name of the SSH key pair provided to the vsi for authenticating users.",
						MarkdownDescription: "SSHKey is the name of the SSH key pair provided to the vsi for authenticating users.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"system_type": schema.StringAttribute{
						Description:         "systemType is the System type used to host the instance.systemType determines the number of cores and memory that is available.Few of the supported SystemTypes are s922,e880,e980.e880 systemType available only in Dallas Datacenters.e980 systemType available in Datacenters except Dallas and Washington.When omitted, this means that the user has no opinion and the platform is left to choose areasonable default, which is subject to change over time. The current default is s922 which is generally available.",
						MarkdownDescription: "systemType is the System type used to host the instance.systemType determines the number of cores and memory that is available.Few of the supported SystemTypes are s922,e880,e980.e880 systemType available only in Dallas Datacenters.e980 systemType available in Datacenters except Dallas and Washington.When omitted, this means that the user has no opinion and the platform is left to choose areasonable default, which is subject to change over time. The current default is s922 which is generally available.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("s922", "e880", "e980", "s1022", ""),
						},
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *InfrastructureClusterXK8SIoIbmpowerVsmachineV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_infrastructure_cluster_x_k8s_io_ibm_power_vs_machine_v1beta2_manifest")

	var model InfrastructureClusterXK8SIoIbmpowerVsmachineV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("infrastructure.cluster.x-k8s.io/v1beta2")
	model.Kind = pointer.String("IBMPowerVSMachine")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
