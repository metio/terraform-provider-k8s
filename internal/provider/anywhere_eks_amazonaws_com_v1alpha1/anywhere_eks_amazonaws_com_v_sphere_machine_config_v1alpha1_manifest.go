/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package anywhere_eks_amazonaws_com_v1alpha1

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
	_ datasource.DataSource = &AnywhereEksAmazonawsComVsphereMachineConfigV1Alpha1Manifest{}
)

func NewAnywhereEksAmazonawsComVsphereMachineConfigV1Alpha1Manifest() datasource.DataSource {
	return &AnywhereEksAmazonawsComVsphereMachineConfigV1Alpha1Manifest{}
}

type AnywhereEksAmazonawsComVsphereMachineConfigV1Alpha1Manifest struct{}

type AnywhereEksAmazonawsComVsphereMachineConfigV1Alpha1ManifestData struct {
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
		CloneMode           *string `tfsdk:"clone_mode" json:"cloneMode,omitempty"`
		Datastore           *string `tfsdk:"datastore" json:"datastore,omitempty"`
		DiskGiB             *int64  `tfsdk:"disk_gi_b" json:"diskGiB,omitempty"`
		Folder              *string `tfsdk:"folder" json:"folder,omitempty"`
		HostOSConfiguration *struct {
			BottlerocketConfiguration *struct {
				Boot *struct {
					BootKernelParameters *map[string][]string `tfsdk:"boot_kernel_parameters" json:"bootKernelParameters,omitempty"`
				} `tfsdk:"boot" json:"boot,omitempty"`
				Kernel *struct {
					SysctlSettings *map[string]string `tfsdk:"sysctl_settings" json:"sysctlSettings,omitempty"`
				} `tfsdk:"kernel" json:"kernel,omitempty"`
				Kubernetes *struct {
					AllowedUnsafeSysctls *[]string `tfsdk:"allowed_unsafe_sysctls" json:"allowedUnsafeSysctls,omitempty"`
					ClusterDNSIPs        *[]string `tfsdk:"cluster_dnsi_ps" json:"clusterDNSIPs,omitempty"`
					MaxPods              *int64    `tfsdk:"max_pods" json:"maxPods,omitempty"`
				} `tfsdk:"kubernetes" json:"kubernetes,omitempty"`
			} `tfsdk:"bottlerocket_configuration" json:"bottlerocketConfiguration,omitempty"`
			CertBundles *[]struct {
				Data *string `tfsdk:"data" json:"data,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"cert_bundles" json:"certBundles,omitempty"`
			NtpConfiguration *struct {
				Servers *[]string `tfsdk:"servers" json:"servers,omitempty"`
			} `tfsdk:"ntp_configuration" json:"ntpConfiguration,omitempty"`
		} `tfsdk:"host_os_configuration" json:"hostOSConfiguration,omitempty"`
		MemoryMiB         *int64    `tfsdk:"memory_mi_b" json:"memoryMiB,omitempty"`
		NumCPUs           *int64    `tfsdk:"num_cp_us" json:"numCPUs,omitempty"`
		OsFamily          *string   `tfsdk:"os_family" json:"osFamily,omitempty"`
		ResourcePool      *string   `tfsdk:"resource_pool" json:"resourcePool,omitempty"`
		StoragePolicyName *string   `tfsdk:"storage_policy_name" json:"storagePolicyName,omitempty"`
		Tags              *[]string `tfsdk:"tags" json:"tags,omitempty"`
		Template          *string   `tfsdk:"template" json:"template,omitempty"`
		Users             *[]struct {
			Name              *string   `tfsdk:"name" json:"name,omitempty"`
			SshAuthorizedKeys *[]string `tfsdk:"ssh_authorized_keys" json:"sshAuthorizedKeys,omitempty"`
		} `tfsdk:"users" json:"users,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AnywhereEksAmazonawsComVsphereMachineConfigV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_anywhere_eks_amazonaws_com_v_sphere_machine_config_v1alpha1_manifest"
}

func (r *AnywhereEksAmazonawsComVsphereMachineConfigV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VSphereMachineConfig is the Schema for the vspheremachineconfigs API.",
		MarkdownDescription: "VSphereMachineConfig is the Schema for the vspheremachineconfigs API.",
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
				Description:         "VSphereMachineConfigSpec defines the desired state of VSphereMachineConfig.",
				MarkdownDescription: "VSphereMachineConfigSpec defines the desired state of VSphereMachineConfig.",
				Attributes: map[string]schema.Attribute{
					"clone_mode": schema.StringAttribute{
						Description:         "CloneMode describes the clone mode to be used when cloning vSphere VMs.",
						MarkdownDescription: "CloneMode describes the clone mode to be used when cloning vSphere VMs.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("fullClone", "linkedClone"),
						},
					},

					"datastore": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"disk_gi_b": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"folder": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"host_os_configuration": schema.SingleNestedAttribute{
						Description:         "HostOSConfiguration defines the configuration settings on the host OS.",
						MarkdownDescription: "HostOSConfiguration defines the configuration settings on the host OS.",
						Attributes: map[string]schema.Attribute{
							"bottlerocket_configuration": schema.SingleNestedAttribute{
								Description:         "BottlerocketConfiguration defines the Bottlerocket configuration on the host OS. These settings only take effect when the 'osFamily' is bottlerocket.",
								MarkdownDescription: "BottlerocketConfiguration defines the Bottlerocket configuration on the host OS. These settings only take effect when the 'osFamily' is bottlerocket.",
								Attributes: map[string]schema.Attribute{
									"boot": schema.SingleNestedAttribute{
										Description:         "Boot defines the boot settings for bottlerocket.",
										MarkdownDescription: "Boot defines the boot settings for bottlerocket.",
										Attributes: map[string]schema.Attribute{
											"boot_kernel_parameters": schema.MapAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.ListType{ElemType: types.StringType},
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"kernel": schema.SingleNestedAttribute{
										Description:         "Kernel defines the kernel settings for bottlerocket.",
										MarkdownDescription: "Kernel defines the kernel settings for bottlerocket.",
										Attributes: map[string]schema.Attribute{
											"sysctl_settings": schema.MapAttribute{
												Description:         "SysctlSettings defines the kernel sysctl settings to set for bottlerocket nodes.",
												MarkdownDescription: "SysctlSettings defines the kernel sysctl settings to set for bottlerocket nodes.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"kubernetes": schema.SingleNestedAttribute{
										Description:         "Kubernetes defines the Kubernetes settings on the host OS.",
										MarkdownDescription: "Kubernetes defines the Kubernetes settings on the host OS.",
										Attributes: map[string]schema.Attribute{
											"allowed_unsafe_sysctls": schema.ListAttribute{
												Description:         "AllowedUnsafeSysctls defines the list of unsafe sysctls that can be set on a node.",
												MarkdownDescription: "AllowedUnsafeSysctls defines the list of unsafe sysctls that can be set on a node.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"cluster_dnsi_ps": schema.ListAttribute{
												Description:         "ClusterDNSIPs defines IP addresses of the DNS servers.",
												MarkdownDescription: "ClusterDNSIPs defines IP addresses of the DNS servers.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_pods": schema.Int64Attribute{
												Description:         "MaxPods defines the maximum number of pods that can run on a node.",
												MarkdownDescription: "MaxPods defines the maximum number of pods that can run on a node.",
												Required:            false,
												Optional:            true,
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

							"cert_bundles": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"data": schema.StringAttribute{
											Description:         "Data defines the cert bundle data.",
											MarkdownDescription: "Data defines the cert bundle data.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name defines the cert bundle name.",
											MarkdownDescription: "Name defines the cert bundle name.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"ntp_configuration": schema.SingleNestedAttribute{
								Description:         "NTPConfiguration defines the NTP configuration on the host OS.",
								MarkdownDescription: "NTPConfiguration defines the NTP configuration on the host OS.",
								Attributes: map[string]schema.Attribute{
									"servers": schema.ListAttribute{
										Description:         "Servers defines a list of NTP servers to be configured on the host OS.",
										MarkdownDescription: "Servers defines a list of NTP servers to be configured on the host OS.",
										ElementType:         types.StringType,
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

					"memory_mi_b": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"num_cp_us": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"os_family": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"resource_pool": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"storage_policy_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tags": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"template": schema.StringAttribute{
						Description:         "Template field is the template to use for provisioning the VM. It must include the Kubernetes version(s). For example, a template used for Kubernetes 1.27 could be ubuntu-2204-1.27.",
						MarkdownDescription: "Template field is the template to use for provisioning the VM. It must include the Kubernetes version(s). For example, a template used for Kubernetes 1.27 could be ubuntu-2204-1.27.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"users": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"ssh_authorized_keys": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            true,
									Optional:            false,
									Computed:            false,
								},
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
	}
}

func (r *AnywhereEksAmazonawsComVsphereMachineConfigV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_anywhere_eks_amazonaws_com_v_sphere_machine_config_v1alpha1_manifest")

	var model AnywhereEksAmazonawsComVsphereMachineConfigV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("anywhere.eks.amazonaws.com/v1alpha1")
	model.Kind = pointer.String("VSphereMachineConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
