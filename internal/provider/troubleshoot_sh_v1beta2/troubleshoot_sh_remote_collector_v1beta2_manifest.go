/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package troubleshoot_sh_v1beta2

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
	_ datasource.DataSource = &TroubleshootShRemoteCollectorV1Beta2Manifest{}
)

func NewTroubleshootShRemoteCollectorV1Beta2Manifest() datasource.DataSource {
	return &TroubleshootShRemoteCollectorV1Beta2Manifest{}
}

type TroubleshootShRemoteCollectorV1Beta2Manifest struct{}

type TroubleshootShRemoteCollectorV1Beta2ManifestData struct {
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
		AfterCollection *[]struct {
			Callback *struct {
				Method    *string `tfsdk:"method" json:"method,omitempty"`
				RedactUri *string `tfsdk:"redact_uri" json:"redactUri,omitempty"`
				Uri       *string `tfsdk:"uri" json:"uri,omitempty"`
			} `tfsdk:"callback" json:"callback,omitempty"`
			UploadResultsTo *struct {
				Method    *string `tfsdk:"method" json:"method,omitempty"`
				RedactUri *string `tfsdk:"redact_uri" json:"redactUri,omitempty"`
				Uri       *string `tfsdk:"uri" json:"uri,omitempty"`
			} `tfsdk:"upload_results_to" json:"uploadResultsTo,omitempty"`
		} `tfsdk:"after_collection" json:"afterCollection,omitempty"`
		Collectors *[]struct {
			BlockDevices *struct {
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
			} `tfsdk:"block_devices" json:"blockDevices,omitempty"`
			Certificate *struct {
				CertificatePath *string `tfsdk:"certificate_path" json:"certificatePath,omitempty"`
				CollectorName   *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude         UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
				KeyPath         *string `tfsdk:"key_path" json:"keyPath,omitempty"`
			} `tfsdk:"certificate" json:"certificate,omitempty"`
			Cpu *struct {
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
			} `tfsdk:"cpu" json:"cpu,omitempty"`
			DiskUsage *struct {
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
				Path          *string `tfsdk:"path" json:"path,omitempty"`
			} `tfsdk:"disk_usage" json:"diskUsage,omitempty"`
			FilesystemPerformance *struct {
				BackgroundIOPSWarmupSeconds *int64  `tfsdk:"background_iops_warmup_seconds" json:"backgroundIOPSWarmupSeconds,omitempty"`
				BackgroundReadIOPS          *int64  `tfsdk:"background_read_iops" json:"backgroundReadIOPS,omitempty"`
				BackgroundReadIOPSJobs      *int64  `tfsdk:"background_read_iops_jobs" json:"backgroundReadIOPSJobs,omitempty"`
				BackgroundWriteIOPS         *int64  `tfsdk:"background_write_iops" json:"backgroundWriteIOPS,omitempty"`
				BackgroundWriteIOPSJobs     *int64  `tfsdk:"background_write_iops_jobs" json:"backgroundWriteIOPSJobs,omitempty"`
				CollectorName               *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Datasync                    *bool   `tfsdk:"datasync" json:"datasync,omitempty"`
				Directory                   *string `tfsdk:"directory" json:"directory,omitempty"`
				EnableBackgroundIOPS        *bool   `tfsdk:"enable_background_iops" json:"enableBackgroundIOPS,omitempty"`
				Exclude                     UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
				FileSize                    *string `tfsdk:"file_size" json:"fileSize,omitempty"`
				OperationSize               *int64  `tfsdk:"operation_size" json:"operationSize,omitempty"`
				RunTime                     *string `tfsdk:"run_time" json:"runTime,omitempty"`
				Sync                        *bool   `tfsdk:"sync" json:"sync,omitempty"`
				Timeout                     *string `tfsdk:"timeout" json:"timeout,omitempty"`
			} `tfsdk:"filesystem_performance" json:"filesystemPerformance,omitempty"`
			HostOS *struct {
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
			} `tfsdk:"host_os" json:"hostOS,omitempty"`
			HostServices *struct {
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
			} `tfsdk:"host_services" json:"hostServices,omitempty"`
			Http *struct {
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
				Get           *struct {
					Headers            *map[string]string `tfsdk:"headers" json:"headers,omitempty"`
					InsecureSkipVerify *bool              `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
					Proxy              *string            `tfsdk:"proxy" json:"proxy,omitempty"`
					Timeout            *string            `tfsdk:"timeout" json:"timeout,omitempty"`
					Tls                *struct {
						Cacert     *string `tfsdk:"cacert" json:"cacert,omitempty"`
						ClientCert *string `tfsdk:"client_cert" json:"clientCert,omitempty"`
						ClientKey  *string `tfsdk:"client_key" json:"clientKey,omitempty"`
						Secret     *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"secret" json:"secret,omitempty"`
						SkipVerify *bool `tfsdk:"skip_verify" json:"skipVerify,omitempty"`
					} `tfsdk:"tls" json:"tls,omitempty"`
					Url *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"get" json:"get,omitempty"`
				Post *struct {
					Body               *string            `tfsdk:"body" json:"body,omitempty"`
					Headers            *map[string]string `tfsdk:"headers" json:"headers,omitempty"`
					InsecureSkipVerify *bool              `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
					Proxy              *string            `tfsdk:"proxy" json:"proxy,omitempty"`
					Timeout            *string            `tfsdk:"timeout" json:"timeout,omitempty"`
					Tls                *struct {
						Cacert     *string `tfsdk:"cacert" json:"cacert,omitempty"`
						ClientCert *string `tfsdk:"client_cert" json:"clientCert,omitempty"`
						ClientKey  *string `tfsdk:"client_key" json:"clientKey,omitempty"`
						Secret     *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"secret" json:"secret,omitempty"`
						SkipVerify *bool `tfsdk:"skip_verify" json:"skipVerify,omitempty"`
					} `tfsdk:"tls" json:"tls,omitempty"`
					Url *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"post" json:"post,omitempty"`
				Put *struct {
					Body               *string            `tfsdk:"body" json:"body,omitempty"`
					Headers            *map[string]string `tfsdk:"headers" json:"headers,omitempty"`
					InsecureSkipVerify *bool              `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
					Proxy              *string            `tfsdk:"proxy" json:"proxy,omitempty"`
					Timeout            *string            `tfsdk:"timeout" json:"timeout,omitempty"`
					Tls                *struct {
						Cacert     *string `tfsdk:"cacert" json:"cacert,omitempty"`
						ClientCert *string `tfsdk:"client_cert" json:"clientCert,omitempty"`
						ClientKey  *string `tfsdk:"client_key" json:"clientKey,omitempty"`
						Secret     *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"secret" json:"secret,omitempty"`
						SkipVerify *bool `tfsdk:"skip_verify" json:"skipVerify,omitempty"`
					} `tfsdk:"tls" json:"tls,omitempty"`
					Url *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"put" json:"put,omitempty"`
			} `tfsdk:"http" json:"http,omitempty"`
			HttpLoadBalancer *struct {
				Address       *string `tfsdk:"address" json:"address,omitempty"`
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
				Path          *string `tfsdk:"path" json:"path,omitempty"`
				Port          *int64  `tfsdk:"port" json:"port,omitempty"`
				Timeout       *string `tfsdk:"timeout" json:"timeout,omitempty"`
			} `tfsdk:"http_load_balancer" json:"httpLoadBalancer,omitempty"`
			Ipv4Interfaces *struct {
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
			} `tfsdk:"ipv4_interfaces" json:"ipv4Interfaces,omitempty"`
			KernelModules *struct {
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
			} `tfsdk:"kernel_modules" json:"kernelModules,omitempty"`
			Memory *struct {
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
			} `tfsdk:"memory" json:"memory,omitempty"`
			SubnetAvailable *struct {
				CIDRRangeAlloc *string `tfsdk:"cidr_range_alloc" json:"CIDRRangeAlloc,omitempty"`
				CollectorName  *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				DesiredCIDR    *int64  `tfsdk:"desired_cidr" json:"desiredCIDR,omitempty"`
				Exclude        UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
			} `tfsdk:"subnet_available" json:"subnetAvailable,omitempty"`
			SystemPackages *struct {
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
			} `tfsdk:"system_packages" json:"systemPackages,omitempty"`
			TcpConnect *struct {
				Address       *string `tfsdk:"address" json:"address,omitempty"`
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
				Timeout       *string `tfsdk:"timeout" json:"timeout,omitempty"`
			} `tfsdk:"tcp_connect" json:"tcpConnect,omitempty"`
			TcpLoadBalancer *struct {
				Address       *string `tfsdk:"address" json:"address,omitempty"`
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
				Port          *int64  `tfsdk:"port" json:"port,omitempty"`
				Timeout       *string `tfsdk:"timeout" json:"timeout,omitempty"`
			} `tfsdk:"tcp_load_balancer" json:"tcpLoadBalancer,omitempty"`
			TcpPortStatus *struct {
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
				Interface     *string `tfsdk:"interface" json:"interface,omitempty"`
				Port          *int64  `tfsdk:"port" json:"port,omitempty"`
			} `tfsdk:"tcp_port_status" json:"tcpPortStatus,omitempty"`
			Time *struct {
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
			} `tfsdk:"time" json:"time,omitempty"`
			UdpPortStatus *struct {
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
				Interface     *string `tfsdk:"interface" json:"interface,omitempty"`
				Port          *int64  `tfsdk:"port" json:"port,omitempty"`
			} `tfsdk:"udp_port_status" json:"udpPortStatus,omitempty"`
		} `tfsdk:"collectors" json:"collectors,omitempty"`
		NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
		Uri          *string            `tfsdk:"uri" json:"uri,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TroubleshootShRemoteCollectorV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_troubleshoot_sh_remote_collector_v1beta2_manifest"
}

func (r *TroubleshootShRemoteCollectorV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "RemoteCollector is the Schema for the remote collectors API",
		MarkdownDescription: "RemoteCollector is the Schema for the remote collectors API",
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
				Description:         "RemoteCollectorSpec defines the desired state of the RemoteCollector",
				MarkdownDescription: "RemoteCollectorSpec defines the desired state of the RemoteCollector",
				Attributes: map[string]schema.Attribute{
					"after_collection": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"callback": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"method": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"redact_uri": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"uri": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"upload_results_to": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"method": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"redact_uri": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"uri": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"collectors": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"block_devices": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"certificate": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"certificate_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"key_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"cpu": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"disk_usage": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"filesystem_performance": schema.SingleNestedAttribute{
									Description:         "RemoteFilesystemPerformance benchmarks sequential write latency on a single file. The optional background IOPS feature attempts to mimic real-world conditions by running read and write workloads prior to and during benchmark execution.",
									MarkdownDescription: "RemoteFilesystemPerformance benchmarks sequential write latency on a single file. The optional background IOPS feature attempts to mimic real-world conditions by running read and write workloads prior to and during benchmark execution.",
									Attributes: map[string]schema.Attribute{
										"background_iops_warmup_seconds": schema.Int64Attribute{
											Description:         "How long to run the background IOPS read and write workloads prior to starting the benchmarks.",
											MarkdownDescription: "How long to run the background IOPS read and write workloads prior to starting the benchmarks.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"background_read_iops": schema.Int64Attribute{
											Description:         "The target read IOPS to run while benchmarking. This is a limit and there is no guarantee it will be reached. This is the total IOPS for all background read jobs.",
											MarkdownDescription: "The target read IOPS to run while benchmarking. This is a limit and there is no guarantee it will be reached. This is the total IOPS for all background read jobs.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"background_read_iops_jobs": schema.Int64Attribute{
											Description:         "Number of threads to use for background read IOPS. This should be set high enough to reach the target specified in BackgrounReadIOPS.",
											MarkdownDescription: "Number of threads to use for background read IOPS. This should be set high enough to reach the target specified in BackgrounReadIOPS.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"background_write_iops": schema.Int64Attribute{
											Description:         "The target write IOPS to run while benchmarking. This is a limit and there is no guarantee it will be reached. This is the total IOPS for all background write jobs.",
											MarkdownDescription: "The target write IOPS to run while benchmarking. This is a limit and there is no guarantee it will be reached. This is the total IOPS for all background write jobs.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"background_write_iops_jobs": schema.Int64Attribute{
											Description:         "Number of threads to use for background write IOPS. This should be set high enough to reach the target specified in BackgroundWriteIOPS. Example: If BackgroundWriteIOPS is 100 and write latency is 10ms then a single job would barely be able to reach 100 IOPS so this should be at least 2.",
											MarkdownDescription: "Number of threads to use for background write IOPS. This should be set high enough to reach the target specified in BackgroundWriteIOPS. Example: If BackgroundWriteIOPS is 100 and write latency is 10ms then a single job would barely be able to reach 100 IOPS so this should be at least 2.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"datasync": schema.BoolAttribute{
											Description:         "Whether to call datasync on the file after each write. Skipped if Sync is also true. Does not apply to background IOPS task.",
											MarkdownDescription: "Whether to call datasync on the file after each write. Skipped if Sync is also true. Does not apply to background IOPS task.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"directory": schema.StringAttribute{
											Description:         "The directory where the benchmark will create files.",
											MarkdownDescription: "The directory where the benchmark will create files.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"enable_background_iops": schema.BoolAttribute{
											Description:         "Enable the background IOPS feature.",
											MarkdownDescription: "Enable the background IOPS feature.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"file_size": schema.StringAttribute{
											Description:         "The size of the file used in the benchmark. The number of IO operations for the benchmark will be FileSize / OperationSizeBytes. Accepts valid Kubernetes resource units such as Mi.",
											MarkdownDescription: "The size of the file used in the benchmark. The number of IO operations for the benchmark will be FileSize / OperationSizeBytes. Accepts valid Kubernetes resource units such as Mi.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"operation_size": schema.Int64Attribute{
											Description:         "The size of each write operation performed while benchmarking. This does not apply to the background IOPS feature if enabled, since those must be fixed at 4096.",
											MarkdownDescription: "The size of each write operation performed while benchmarking. This does not apply to the background IOPS feature if enabled, since those must be fixed at 4096.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"run_time": schema.StringAttribute{
											Description:         "Limit runtime. The test will run until it completes the configured I/O workload or until it has run for this specified amount of time, whichever occurs first. When the unit is omitted, the value is interpreted in seconds. Defaults to 120 seconds. Set to '0' to disable.",
											MarkdownDescription: "Limit runtime. The test will run until it completes the configured I/O workload or until it has run for this specified amount of time, whichever occurs first. When the unit is omitted, the value is interpreted in seconds. Defaults to 120 seconds. Set to '0' to disable.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"sync": schema.BoolAttribute{
											Description:         "Whether to call sync on the file after each write. Does not apply to background IOPS task.",
											MarkdownDescription: "Whether to call sync on the file after each write. Does not apply to background IOPS task.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"timeout": schema.StringAttribute{
											Description:         "Total timeout, including background IOPS setup and warmup if enabled.",
											MarkdownDescription: "Total timeout, including background IOPS setup and warmup if enabled.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"host_os": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"host_services": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"http": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"get": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"headers": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"insecure_skip_verify": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"proxy": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"timeout": schema.StringAttribute{
													Description:         "Timeout is the time to wait for a server's response. Its a duration e.g 15s, 2h30m. Missing value or empty string or means no timeout.",
													MarkdownDescription: "Timeout is the time to wait for a server's response. Its a duration e.g 15s, 2h30m. Missing value or empty string or means no timeout.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tls": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"cacert": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"client_cert": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"client_key": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"secret": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"namespace": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"skip_verify": schema.BoolAttribute{
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

												"url": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"post": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"body": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"headers": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"insecure_skip_verify": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"proxy": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"timeout": schema.StringAttribute{
													Description:         "Timeout is the time to wait for a server's response. Its a duration e.g 15s, 2h30m. Missing value or empty string or means no timeout.",
													MarkdownDescription: "Timeout is the time to wait for a server's response. Its a duration e.g 15s, 2h30m. Missing value or empty string or means no timeout.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tls": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"cacert": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"client_cert": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"client_key": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"secret": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"namespace": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"skip_verify": schema.BoolAttribute{
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

												"url": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"put": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"body": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"headers": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"insecure_skip_verify": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"proxy": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"timeout": schema.StringAttribute{
													Description:         "Timeout is the time to wait for a server's response. Its a duration e.g 15s, 2h30m. Missing value or empty string or means no timeout.",
													MarkdownDescription: "Timeout is the time to wait for a server's response. Its a duration e.g 15s, 2h30m. Missing value or empty string or means no timeout.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tls": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"cacert": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"client_cert": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"client_key": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"secret": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"namespace": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"skip_verify": schema.BoolAttribute{
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

												"url": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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

								"http_load_balancer": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"address": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"port": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"timeout": schema.StringAttribute{
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

								"ipv4_interfaces": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"kernel_modules": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"memory": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"subnet_available": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"cidr_range_alloc": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"desired_cidr": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"system_packages": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"tcp_connect": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"address": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"timeout": schema.StringAttribute{
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

								"tcp_load_balancer": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"address": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"port": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"timeout": schema.StringAttribute{
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

								"tcp_port_status": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"interface": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"port": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"time": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"udp_port_status": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"interface": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"port": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"node_selector": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"uri": schema.StringAttribute{
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
		},
	}
}

func (r *TroubleshootShRemoteCollectorV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_troubleshoot_sh_remote_collector_v1beta2_manifest")

	var model TroubleshootShRemoteCollectorV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("troubleshoot.sh/v1beta2")
	model.Kind = pointer.String("RemoteCollector")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
