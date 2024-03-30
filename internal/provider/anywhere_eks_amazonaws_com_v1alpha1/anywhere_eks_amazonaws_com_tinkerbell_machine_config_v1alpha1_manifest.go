/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package anywhere_eks_amazonaws_com_v1alpha1

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
	_ datasource.DataSource = &AnywhereEksAmazonawsComTinkerbellMachineConfigV1Alpha1Manifest{}
)

func NewAnywhereEksAmazonawsComTinkerbellMachineConfigV1Alpha1Manifest() datasource.DataSource {
	return &AnywhereEksAmazonawsComTinkerbellMachineConfigV1Alpha1Manifest{}
}

type AnywhereEksAmazonawsComTinkerbellMachineConfigV1Alpha1Manifest struct{}

type AnywhereEksAmazonawsComTinkerbellMachineConfigV1Alpha1ManifestData struct {
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
		HardwareSelector    *map[string]string `tfsdk:"hardware_selector" json:"hardwareSelector,omitempty"`
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
		OsFamily    *string `tfsdk:"os_family" json:"osFamily,omitempty"`
		OsImageURL  *string `tfsdk:"os_image_url" json:"osImageURL,omitempty"`
		TemplateRef *struct {
			Kind *string `tfsdk:"kind" json:"kind,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"template_ref" json:"templateRef,omitempty"`
		Users *[]struct {
			Name              *string   `tfsdk:"name" json:"name,omitempty"`
			SshAuthorizedKeys *[]string `tfsdk:"ssh_authorized_keys" json:"sshAuthorizedKeys,omitempty"`
		} `tfsdk:"users" json:"users,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AnywhereEksAmazonawsComTinkerbellMachineConfigV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_anywhere_eks_amazonaws_com_tinkerbell_machine_config_v1alpha1_manifest"
}

func (r *AnywhereEksAmazonawsComTinkerbellMachineConfigV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "TinkerbellMachineConfig is the Schema for the tinkerbellmachineconfigs API.",
		MarkdownDescription: "TinkerbellMachineConfig is the Schema for the tinkerbellmachineconfigs API.",
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
				Description:         "TinkerbellMachineConfigSpec defines the desired state of TinkerbellMachineConfig.",
				MarkdownDescription: "TinkerbellMachineConfigSpec defines the desired state of TinkerbellMachineConfig.",
				Attributes: map[string]schema.Attribute{
					"hardware_selector": schema.MapAttribute{
						Description:         "HardwareSelector models a simple key-value selector used in Tinkerbell provisioning.",
						MarkdownDescription: "HardwareSelector models a simple key-value selector used in Tinkerbell provisioning.",
						ElementType:         types.StringType,
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

					"os_family": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"os_image_url": schema.StringAttribute{
						Description:         "OSImageURL can be used to override the default OS image path to pull from a local server. OSImageURL is a URL to the OS image used during provisioning. It must include the Kubernetes version(s). For example, a URL used for Kubernetes 1.27 could be http://localhost:8080/ubuntu-2204-1.27.tgz",
						MarkdownDescription: "OSImageURL can be used to override the default OS image path to pull from a local server. OSImageURL is a URL to the OS image used during provisioning. It must include the Kubernetes version(s). For example, a URL used for Kubernetes 1.27 could be http://localhost:8080/ubuntu-2204-1.27.tgz",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"template_ref": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
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

func (r *AnywhereEksAmazonawsComTinkerbellMachineConfigV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_anywhere_eks_amazonaws_com_tinkerbell_machine_config_v1alpha1_manifest")

	var model AnywhereEksAmazonawsComTinkerbellMachineConfigV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("anywhere.eks.amazonaws.com/v1alpha1")
	model.Kind = pointer.String("TinkerbellMachineConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
