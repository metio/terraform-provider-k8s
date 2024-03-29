/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package tinkerbell_org_v1alpha1

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
	_ datasource.DataSource = &TinkerbellOrgStackV1Alpha1Manifest{}
)

func NewTinkerbellOrgStackV1Alpha1Manifest() datasource.DataSource {
	return &TinkerbellOrgStackV1Alpha1Manifest{}
}

type TinkerbellOrgStackV1Alpha1Manifest struct{}

type TinkerbellOrgStackV1Alpha1ManifestData struct {
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
		DnsResolverIP    *string   `tfsdk:"dns_resolver_ip" json:"dnsResolverIP,omitempty"`
		ImagePullSecrets *[]string `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
		Registry         *string   `tfsdk:"registry" json:"registry,omitempty"`
		Services         *struct {
			Hegel *struct {
				Image *struct {
					Repository *string `tfsdk:"repository" json:"repository,omitempty"`
					Tag        *string `tfsdk:"tag" json:"tag,omitempty"`
				} `tfsdk:"image" json:"image,omitempty"`
				TrustedProxies *[]string `tfsdk:"trusted_proxies" json:"trustedProxies,omitempty"`
			} `tfsdk:"hegel" json:"hegel,omitempty"`
			Rufio *struct {
				Image *struct {
					Repository *string `tfsdk:"repository" json:"repository,omitempty"`
					Tag        *string `tfsdk:"tag" json:"tag,omitempty"`
				} `tfsdk:"image" json:"image,omitempty"`
			} `tfsdk:"rufio" json:"rufio,omitempty"`
			Smee *struct {
				BackendConfigs *struct {
					BackendFileMode *struct {
						FilePath *string `tfsdk:"file_path" json:"filePath,omitempty"`
					} `tfsdk:"backend_file_mode" json:"backendFileMode,omitempty"`
					BackendKubeMode *struct {
						ConfigFilePath *string `tfsdk:"config_file_path" json:"configFilePath,omitempty"`
						KubeAPIURL     *string `tfsdk:"kube_api_url" json:"kubeAPIURL,omitempty"`
						KubeNamespace  *string `tfsdk:"kube_namespace" json:"kubeNamespace,omitempty"`
					} `tfsdk:"backend_kube_mode" json:"backendKubeMode,omitempty"`
				} `tfsdk:"backend_configs" json:"backendConfigs,omitempty"`
				DhcpConfigs *struct {
					IPForPacket           *string `tfsdk:"ip_for_packet" json:"IPForPacket,omitempty"`
					HttpIPXEBinaryAddress *string `tfsdk:"http_ipxe_binary_address" json:"httpIPXEBinaryAddress,omitempty"`
					HttpIPXEBinaryURI     *string `tfsdk:"http_ipxe_binary_uri" json:"httpIPXEBinaryURI,omitempty"`
					Ip                    *string `tfsdk:"ip" json:"ip,omitempty"`
					Port                  *int64  `tfsdk:"port" json:"port,omitempty"`
					SyslogIP              *string `tfsdk:"syslog_ip" json:"syslogIP,omitempty"`
					TftpAddress           *string `tfsdk:"tftp_address" json:"tftpAddress,omitempty"`
				} `tfsdk:"dhcp_configs" json:"dhcpConfigs,omitempty"`
				Image *struct {
					Repository *string `tfsdk:"repository" json:"repository,omitempty"`
					Tag        *string `tfsdk:"tag" json:"tag,omitempty"`
				} `tfsdk:"image" json:"image,omitempty"`
				IpxeConfigs *struct {
					EnableHTTPBinary  *bool     `tfsdk:"enable_http_binary" json:"enableHTTPBinary,omitempty"`
					EnableTLS         *bool     `tfsdk:"enable_tls" json:"enableTLS,omitempty"`
					ExtraKernelArgs   *string   `tfsdk:"extra_kernel_args" json:"extraKernelArgs,omitempty"`
					HookURL           *string   `tfsdk:"hook_url" json:"hookURL,omitempty"`
					Ip                *string   `tfsdk:"ip" json:"ip,omitempty"`
					Port              *int64    `tfsdk:"port" json:"port,omitempty"`
					TinkServerAddress *string   `tfsdk:"tink_server_address" json:"tinkServerAddress,omitempty"`
					TrustedProxies    *[]string `tfsdk:"trusted_proxies" json:"trustedProxies,omitempty"`
				} `tfsdk:"ipxe_configs" json:"ipxeConfigs,omitempty"`
				LogLevel      *string `tfsdk:"log_level" json:"logLevel,omitempty"`
				SyslogConfigs *struct {
					BindAddress *string `tfsdk:"bind_address" json:"bindAddress,omitempty"`
					Port        *int64  `tfsdk:"port" json:"port,omitempty"`
				} `tfsdk:"syslog_configs" json:"syslogConfigs,omitempty"`
				TftpConfigs *struct {
					Ip              *string `tfsdk:"ip" json:"ip,omitempty"`
					IpxeScriptPatch *string `tfsdk:"ipxe_script_patch" json:"ipxeScriptPatch,omitempty"`
					Port            *int64  `tfsdk:"port" json:"port,omitempty"`
					TftpTimeout     *int64  `tfsdk:"tftp_timeout" json:"tftpTimeout,omitempty"`
				} `tfsdk:"tftp_configs" json:"tftpConfigs,omitempty"`
			} `tfsdk:"smee" json:"smee,omitempty"`
			TinkController *struct {
				Image *struct {
					Repository *string `tfsdk:"repository" json:"repository,omitempty"`
					Tag        *string `tfsdk:"tag" json:"tag,omitempty"`
				} `tfsdk:"image" json:"image,omitempty"`
			} `tfsdk:"tink_controller" json:"tinkController,omitempty"`
			TinkServer *struct {
				EnableTLS *bool `tfsdk:"enable_tls" json:"enableTLS,omitempty"`
				Image     *struct {
					Repository *string `tfsdk:"repository" json:"repository,omitempty"`
					Tag        *string `tfsdk:"tag" json:"tag,omitempty"`
				} `tfsdk:"image" json:"image,omitempty"`
			} `tfsdk:"tink_server" json:"tinkServer,omitempty"`
		} `tfsdk:"services" json:"services,omitempty"`
		Version *string `tfsdk:"version" json:"version,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TinkerbellOrgStackV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_tinkerbell_org_stack_v1alpha1_manifest"
}

func (r *TinkerbellOrgStackV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Stack represents the tinkerbell stack that is being deployed in the kubernetes where the operator is deployed. Tinkerbell operator watches for different resources such as deployment, services, serviceAccounts, etc. One of those CRs is Stack which the operator will install the tink-stack based on its specs. Once the CR is deleted, the operator will delete all tinkerbell resources.",
		MarkdownDescription: "Stack represents the tinkerbell stack that is being deployed in the kubernetes where the operator is deployed. Tinkerbell operator watches for different resources such as deployment, services, serviceAccounts, etc. One of those CRs is Stack which the operator will install the tink-stack based on its specs. Once the CR is deleted, the operator will delete all tinkerbell resources.",
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
				Description:         "Spec describes the desired tinkerbell stack state.",
				MarkdownDescription: "Spec describes the desired tinkerbell stack state.",
				Attributes: map[string]schema.Attribute{
					"dns_resolver_ip": schema.StringAttribute{
						Description:         "DNSResolverIP is indicative of the resolver IP utilized for setting up the nginx server responsible for proxying to the Tinkerbell services and serving the Hook artifacts.",
						MarkdownDescription: "DNSResolverIP is indicative of the resolver IP utilized for setting up the nginx server responsible for proxying to the Tinkerbell services and serving the Hook artifacts.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image_pull_secrets": schema.ListAttribute{
						Description:         "ImagePullSecrets the secret name containing the docker auth config which should exist in the same namespace where the operator is deployed(typically tinkerbell)",
						MarkdownDescription: "ImagePullSecrets the secret name containing the docker auth config which should exist in the same namespace where the operator is deployed(typically tinkerbell)",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"registry": schema.StringAttribute{
						Description:         "Registry is the registry to use for all images. If this field is set, all tink service deployment images will be prefixed with this value. For example if the value here was set to docker.io, then smee image will be docker.io/tinkerbell/smee.",
						MarkdownDescription: "Registry is the registry to use for all images. If this field is set, all tink service deployment images will be prefixed with this value. For example if the value here was set to docker.io, then smee image will be docker.io/tinkerbell/smee.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"services": schema.SingleNestedAttribute{
						Description:         "Services contains all Tinkerbell Stack services.",
						MarkdownDescription: "Services contains all Tinkerbell Stack services.",
						Attributes: map[string]schema.Attribute{
							"hegel": schema.SingleNestedAttribute{
								Description:         "Hegel contains all the information and spec about smee.",
								MarkdownDescription: "Hegel contains all the information and spec about smee.",
								Attributes: map[string]schema.Attribute{
									"image": schema.SingleNestedAttribute{
										Description:         "Image specifies the details of a tinkerbell services images",
										MarkdownDescription: "Image specifies the details of a tinkerbell services images",
										Attributes: map[string]schema.Attribute{
											"repository": schema.StringAttribute{
												Description:         "Repository is used to set the image repository for tinkerbell services.",
												MarkdownDescription: "Repository is used to set the image repository for tinkerbell services.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tag": schema.StringAttribute{
												Description:         "Tag is used to set the image tag for tinkerbell services.",
												MarkdownDescription: "Tag is used to set the image tag for tinkerbell services.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"trusted_proxies": schema.ListAttribute{
										Description:         "TrustedProxies comma separated allowed CIDRs subnets to be used as trusted proxies",
										MarkdownDescription: "TrustedProxies comma separated allowed CIDRs subnets to be used as trusted proxies",
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

							"rufio": schema.SingleNestedAttribute{
								Description:         "Rufio contains all the information and spec about rufio.",
								MarkdownDescription: "Rufio contains all the information and spec about rufio.",
								Attributes: map[string]schema.Attribute{
									"image": schema.SingleNestedAttribute{
										Description:         "Image specifies the details of a tinkerbell services images",
										MarkdownDescription: "Image specifies the details of a tinkerbell services images",
										Attributes: map[string]schema.Attribute{
											"repository": schema.StringAttribute{
												Description:         "Repository is used to set the image repository for tinkerbell services.",
												MarkdownDescription: "Repository is used to set the image repository for tinkerbell services.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tag": schema.StringAttribute{
												Description:         "Tag is used to set the image tag for tinkerbell services.",
												MarkdownDescription: "Tag is used to set the image tag for tinkerbell services.",
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

							"smee": schema.SingleNestedAttribute{
								Description:         "Smee contains all the information and spec about smee.",
								MarkdownDescription: "Smee contains all the information and spec about smee.",
								Attributes: map[string]schema.Attribute{
									"backend_configs": schema.SingleNestedAttribute{
										Description:         "BackendConfigs contains the configurations for smee backend.",
										MarkdownDescription: "BackendConfigs contains the configurations for smee backend.",
										Attributes: map[string]schema.Attribute{
											"backend_file_mode": schema.SingleNestedAttribute{
												Description:         "BackendFileMode contains the file backend configurations for DHCP and the HTTP iPXE script.",
												MarkdownDescription: "BackendFileMode contains the file backend configurations for DHCP and the HTTP iPXE script.",
												Attributes: map[string]schema.Attribute{
													"file_path": schema.StringAttribute{
														Description:         "FilePath specifies the hardware yaml file path for the file backend.",
														MarkdownDescription: "FilePath specifies the hardware yaml file path for the file backend.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"backend_kube_mode": schema.SingleNestedAttribute{
												Description:         "BackendKubeMode contains the Kubernetes backend configurations for DHCP and the HTTP iPXE script.",
												MarkdownDescription: "BackendKubeMode contains the Kubernetes backend configurations for DHCP and the HTTP iPXE script.",
												Attributes: map[string]schema.Attribute{
													"config_file_path": schema.StringAttribute{
														Description:         "ConfigFilePath specifies the Kubernetes config file location.",
														MarkdownDescription: "ConfigFilePath specifies the Kubernetes config file location.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"kube_api_url": schema.StringAttribute{
														Description:         "KubeAPIURL specifies the Kubernetes API URL, used for in-cluster client construction.",
														MarkdownDescription: "KubeAPIURL specifies the Kubernetes API URL, used for in-cluster client construction.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"kube_namespace": schema.StringAttribute{
														Description:         "KubeNamespace specifies an optional Kubernetes namespace override to query hardware data from.",
														MarkdownDescription: "KubeNamespace specifies an optional Kubernetes namespace override to query hardware data from.",
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
										Required: true,
										Optional: false,
										Computed: false,
									},

									"dhcp_configs": schema.SingleNestedAttribute{
										Description:         "DHCPConfigs contains the DHCP server configurations.",
										MarkdownDescription: "DHCPConfigs contains the DHCP server configurations.",
										Attributes: map[string]schema.Attribute{
											"ip_for_packet": schema.StringAttribute{
												Description:         "IPForPacket IP address to use in DHCP packets",
												MarkdownDescription: "IPForPacket IP address to use in DHCP packets",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"http_ipxe_binary_address": schema.StringAttribute{
												Description:         "HTTPIPXEBinaryAddress specifies the http ipxe binary server address (IP:Port) to use in DHCP packets.",
												MarkdownDescription: "HTTPIPXEBinaryAddress specifies the http ipxe binary server address (IP:Port) to use in DHCP packets.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"http_ipxe_binary_uri": schema.StringAttribute{
												Description:         "HTTPIPXEBinaryURI specifies the http ipxe script server URL to use in DHCP packets.",
												MarkdownDescription: "HTTPIPXEBinaryURI specifies the http ipxe script server URL to use in DHCP packets.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ip": schema.StringAttribute{
												Description:         "IP is the local IP to listen on to serve TFTP binaries.",
												MarkdownDescription: "IP is the local IP to listen on to serve TFTP binaries.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"port": schema.Int64Attribute{
												Description:         "Port is the  local port to listen on to serve TFTP binaries.",
												MarkdownDescription: "Port is the  local port to listen on to serve TFTP binaries.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"syslog_ip": schema.StringAttribute{
												Description:         "SyslogIP specifies the syslog server IP address to use in DHCP packets.",
												MarkdownDescription: "SyslogIP specifies the syslog server IP address to use in DHCP packets.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tftp_address": schema.StringAttribute{
												Description:         "TFTPAddress specifies the tftp server address to use in DHCP packets.",
												MarkdownDescription: "TFTPAddress specifies the tftp server address to use in DHCP packets.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"image": schema.SingleNestedAttribute{
										Description:         "Image specifies the image repo and tag for Smee.",
										MarkdownDescription: "Image specifies the image repo and tag for Smee.",
										Attributes: map[string]schema.Attribute{
											"repository": schema.StringAttribute{
												Description:         "Repository is used to set the image repository for tinkerbell services.",
												MarkdownDescription: "Repository is used to set the image repository for tinkerbell services.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tag": schema.StringAttribute{
												Description:         "Tag is used to set the image tag for tinkerbell services.",
												MarkdownDescription: "Tag is used to set the image tag for tinkerbell services.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"ipxe_configs": schema.SingleNestedAttribute{
										Description:         "IPXEConfigs contains the iPXE configurations.",
										MarkdownDescription: "IPXEConfigs contains the iPXE configurations.",
										Attributes: map[string]schema.Attribute{
											"enable_http_binary": schema.BoolAttribute{
												Description:         "EnableHTTPBinary enable iPXE HTTP binary server.",
												MarkdownDescription: "EnableHTTPBinary enable iPXE HTTP binary server.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enable_tls": schema.BoolAttribute{
												Description:         "EnableTLS sets if the smee should run with TLS or not.",
												MarkdownDescription: "EnableTLS sets if the smee should run with TLS or not.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"extra_kernel_args": schema.StringAttribute{
												Description:         "ExtraKernelArgs specifies extra set of kernel args (k=v k=v) that are appended to the kernel cmdline iPXE script.",
												MarkdownDescription: "ExtraKernelArgs specifies extra set of kernel args (k=v k=v) that are appended to the kernel cmdline iPXE script.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"hook_url": schema.StringAttribute{
												Description:         "HookURL specifies the URL where OSIE(Hook) images are located.",
												MarkdownDescription: "HookURL specifies the URL where OSIE(Hook) images are located.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ip": schema.StringAttribute{
												Description:         "IP is the local IP to listen on to serve TFTP binaries.",
												MarkdownDescription: "IP is the local IP to listen on to serve TFTP binaries.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"port": schema.Int64Attribute{
												Description:         "Port is the  local port to listen on to serve TFTP binaries.",
												MarkdownDescription: "Port is the  local port to listen on to serve TFTP binaries.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"tink_server_address": schema.StringAttribute{
												Description:         "TinkServerAddress specifies the IP:Port of the tink server.",
												MarkdownDescription: "TinkServerAddress specifies the IP:Port of the tink server.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"trusted_proxies": schema.ListAttribute{
												Description:         "TrustedProxies comma separated allowed CIDRs subnets to be used as trusted proxies.",
												MarkdownDescription: "TrustedProxies comma separated allowed CIDRs subnets to be used as trusted proxies.",
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

									"log_level": schema.StringAttribute{
										Description:         "LogLevel sets the debug level for smee.",
										MarkdownDescription: "LogLevel sets the debug level for smee.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"syslog_configs": schema.SingleNestedAttribute{
										Description:         "SyslogConfigs contains the configurations of the syslog server.",
										MarkdownDescription: "SyslogConfigs contains the configurations of the syslog server.",
										Attributes: map[string]schema.Attribute{
											"bind_address": schema.StringAttribute{
												Description:         "IP is the local IP to listen on for syslog messages.",
												MarkdownDescription: "IP is the local IP to listen on for syslog messages.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"port": schema.Int64Attribute{
												Description:         "Port is the  local port to listen on for syslog messages.",
												MarkdownDescription: "Port is the  local port to listen on for syslog messages.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"tftp_configs": schema.SingleNestedAttribute{
										Description:         "TFTPConfigs contains the configurations of Tinkerbell TFTP server.",
										MarkdownDescription: "TFTPConfigs contains the configurations of Tinkerbell TFTP server.",
										Attributes: map[string]schema.Attribute{
											"ip": schema.StringAttribute{
												Description:         "IP is the local IP to listen on to serve TFTP binaries.",
												MarkdownDescription: "IP is the local IP to listen on to serve TFTP binaries.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"ipxe_script_patch": schema.StringAttribute{
												Description:         "IPXEScriptPatch specifies the iPXE script fragment to patch into served iPXE binaries served via TFTP or HTTP.",
												MarkdownDescription: "IPXEScriptPatch specifies the iPXE script fragment to patch into served iPXE binaries served via TFTP or HTTP.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"port": schema.Int64Attribute{
												Description:         "Port is the  local port to listen on to serve TFTP binaries.",
												MarkdownDescription: "Port is the  local port to listen on to serve TFTP binaries.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"tftp_timeout": schema.Int64Attribute{
												Description:         "TFTPTimeout specifies the iPXE tftp binary server requests timeout.",
												MarkdownDescription: "TFTPTimeout specifies the iPXE tftp binary server requests timeout.",
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

							"tink_controller": schema.SingleNestedAttribute{
								Description:         "TinkController contains all the information and spec about tink controller.",
								MarkdownDescription: "TinkController contains all the information and spec about tink controller.",
								Attributes: map[string]schema.Attribute{
									"image": schema.SingleNestedAttribute{
										Description:         "Image specifies the details of a tinkerbell services images",
										MarkdownDescription: "Image specifies the details of a tinkerbell services images",
										Attributes: map[string]schema.Attribute{
											"repository": schema.StringAttribute{
												Description:         "Repository is used to set the image repository for tinkerbell services.",
												MarkdownDescription: "Repository is used to set the image repository for tinkerbell services.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tag": schema.StringAttribute{
												Description:         "Tag is used to set the image tag for tinkerbell services.",
												MarkdownDescription: "Tag is used to set the image tag for tinkerbell services.",
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
								Required: true,
								Optional: false,
								Computed: false,
							},

							"tink_server": schema.SingleNestedAttribute{
								Description:         "TinkServer contains all the information and spec about tink server.",
								MarkdownDescription: "TinkServer contains all the information and spec about tink server.",
								Attributes: map[string]schema.Attribute{
									"enable_tls": schema.BoolAttribute{
										Description:         "EnableTLS sets if the tink server should run with TLS or not.",
										MarkdownDescription: "EnableTLS sets if the tink server should run with TLS or not.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"image": schema.SingleNestedAttribute{
										Description:         "Image specifies the details of a tinkerbell services images",
										MarkdownDescription: "Image specifies the details of a tinkerbell services images",
										Attributes: map[string]schema.Attribute{
											"repository": schema.StringAttribute{
												Description:         "Repository is used to set the image repository for tinkerbell services.",
												MarkdownDescription: "Repository is used to set the image repository for tinkerbell services.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tag": schema.StringAttribute{
												Description:         "Tag is used to set the image tag for tinkerbell services.",
												MarkdownDescription: "Tag is used to set the image tag for tinkerbell services.",
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
								Required: true,
								Optional: false,
								Computed: false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"version": schema.StringAttribute{
						Description:         "Version is the Tinkerbell CRD version.",
						MarkdownDescription: "Version is the Tinkerbell CRD version.",
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
	}
}

func (r *TinkerbellOrgStackV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_tinkerbell_org_stack_v1alpha1_manifest")

	var model TinkerbellOrgStackV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("tinkerbell.org/v1alpha1")
	model.Kind = pointer.String("Stack")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
