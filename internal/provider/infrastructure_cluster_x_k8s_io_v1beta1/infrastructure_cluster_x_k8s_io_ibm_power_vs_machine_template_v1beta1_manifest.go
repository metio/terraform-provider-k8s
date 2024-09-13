/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package infrastructure_cluster_x_k8s_io_v1beta1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &InfrastructureClusterXK8SIoIbmpowerVsmachineTemplateV1Beta1Manifest{}
)

func NewInfrastructureClusterXK8SIoIbmpowerVsmachineTemplateV1Beta1Manifest() datasource.DataSource {
	return &InfrastructureClusterXK8SIoIbmpowerVsmachineTemplateV1Beta1Manifest{}
}

type InfrastructureClusterXK8SIoIbmpowerVsmachineTemplateV1Beta1Manifest struct{}

type InfrastructureClusterXK8SIoIbmpowerVsmachineTemplateV1Beta1ManifestData struct {
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
		Template *struct {
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
		} `tfsdk:"template" json:"template,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *InfrastructureClusterXK8SIoIbmpowerVsmachineTemplateV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_infrastructure_cluster_x_k8s_io_ibm_power_vs_machine_template_v1beta1_manifest"
}

func (r *InfrastructureClusterXK8SIoIbmpowerVsmachineTemplateV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "IBMPowerVSMachineTemplate is the Schema for the ibmpowervsmachinetemplates API.",
		MarkdownDescription: "IBMPowerVSMachineTemplate is the Schema for the ibmpowervsmachinetemplates API.",
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
				Description:         "IBMPowerVSMachineTemplateSpec defines the desired state of IBMPowerVSMachineTemplate.",
				MarkdownDescription: "IBMPowerVSMachineTemplateSpec defines the desired state of IBMPowerVSMachineTemplate.",
				Attributes: map[string]schema.Attribute{
					"template": schema.SingleNestedAttribute{
						Description:         "IBMPowerVSMachineTemplateResource holds the IBMPowerVSMachine spec.",
						MarkdownDescription: "IBMPowerVSMachineTemplateResource holds the IBMPowerVSMachine spec.",
						Attributes: map[string]schema.Attribute{
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
												Description:         "Regular expression to match resource, In case of multiple resources matches the provided regular expression the first matched resource will be selected",
												MarkdownDescription: "Regular expression to match resource, In case of multiple resources matches the provided regular expression the first matched resource will be selected",
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
										Description:         "ImageRef is an optional reference to a provider-specific resource that holds the details for provisioning the Image for a Cluster.",
										MarkdownDescription: "ImageRef is an optional reference to a provider-specific resource that holds the details for provisioning the Image for a Cluster.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
												MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"memory": schema.StringAttribute{
										Description:         "Memory is Amount of memory allocated (in GB)",
										MarkdownDescription: "Memory is Amount of memory allocated (in GB)",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"network": schema.SingleNestedAttribute{
										Description:         "Network is the reference to the Network to use for this instance.",
										MarkdownDescription: "Network is the reference to the Network to use for this instance.",
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
												Description:         "Regular expression to match resource, In case of multiple resources matches the provided regular expression the first matched resource will be selected",
												MarkdownDescription: "Regular expression to match resource, In case of multiple resources matches the provided regular expression the first matched resource will be selected",
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

									"proc_type": schema.StringAttribute{
										Description:         "ProcType is the processor type, e.g: dedicated, shared, capped",
										MarkdownDescription: "ProcType is the processor type, e.g: dedicated, shared, capped",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"processors": schema.StringAttribute{
										Description:         "Processors is Number of processors allocated.",
										MarkdownDescription: "Processors is Number of processors allocated.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(\.)?(\d)?(\d)?$`), ""),
										},
									},

									"provider_id": schema.StringAttribute{
										Description:         "ProviderID is the unique identifier as specified by the cloud provider.",
										MarkdownDescription: "ProviderID is the unique identifier as specified by the cloud provider.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"service_instance_id": schema.StringAttribute{
										Description:         "ServiceInstanceID is the id of the power cloud instance where the vsi instance will get deployed.",
										MarkdownDescription: "ServiceInstanceID is the id of the power cloud instance where the vsi instance will get deployed.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
										},
									},

									"ssh_key": schema.StringAttribute{
										Description:         "SSHKey is the name of the SSH key pair provided to the vsi for authenticating users.",
										MarkdownDescription: "SSHKey is the name of the SSH key pair provided to the vsi for authenticating users.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"sys_type": schema.StringAttribute{
										Description:         "SysType is the System type used to host the vsi.",
										MarkdownDescription: "SysType is the System type used to host the vsi.",
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
						Required: true,
						Optional: false,
						Computed: false,
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *InfrastructureClusterXK8SIoIbmpowerVsmachineTemplateV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_infrastructure_cluster_x_k8s_io_ibm_power_vs_machine_template_v1beta1_manifest")

	var model InfrastructureClusterXK8SIoIbmpowerVsmachineTemplateV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("infrastructure.cluster.x-k8s.io/v1beta1")
	model.Kind = pointer.String("IBMPowerVSMachineTemplate")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
