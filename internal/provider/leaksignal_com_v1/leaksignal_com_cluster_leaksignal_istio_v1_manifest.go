/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package leaksignal_com_v1

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
	_ datasource.DataSource = &LeaksignalComClusterLeaksignalIstioV1Manifest{}
)

func NewLeaksignalComClusterLeaksignalIstioV1Manifest() datasource.DataSource {
	return &LeaksignalComClusterLeaksignalIstioV1Manifest{}
}

type LeaksignalComClusterLeaksignalIstioV1Manifest struct{}

type LeaksignalComClusterLeaksignalIstioV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ApiKey                 *string `tfsdk:"api_key" json:"apiKey,omitempty"`
		CaBundle               *string `tfsdk:"ca_bundle" json:"caBundle,omitempty"`
		EnableStreaming        *bool   `tfsdk:"enable_streaming" json:"enableStreaming,omitempty"`
		FailOpen               *bool   `tfsdk:"fail_open" json:"failOpen,omitempty"`
		GrpcMode               *string `tfsdk:"grpc_mode" json:"grpcMode,omitempty"`
		Native                 *bool   `tfsdk:"native" json:"native,omitempty"`
		NativeProxyMemoryLimit *string `tfsdk:"native_proxy_memory_limit" json:"nativeProxyMemoryLimit,omitempty"`
		NativeRepo             *string `tfsdk:"native_repo" json:"nativeRepo,omitempty"`
		ProxyHash              *string `tfsdk:"proxy_hash" json:"proxyHash,omitempty"`
		ProxyPrefix            *string `tfsdk:"proxy_prefix" json:"proxyPrefix,omitempty"`
		ProxyPullLocation      *string `tfsdk:"proxy_pull_location" json:"proxyPullLocation,omitempty"`
		ProxyVersion           *string `tfsdk:"proxy_version" json:"proxyVersion,omitempty"`
		RefreshPodsOnStale     *bool   `tfsdk:"refresh_pods_on_stale" json:"refreshPodsOnStale,omitempty"`
		RefreshPodsOnUpdate    *bool   `tfsdk:"refresh_pods_on_update" json:"refreshPodsOnUpdate,omitempty"`
		Tls                    *bool   `tfsdk:"tls" json:"tls,omitempty"`
		UpstreamLocation       *string `tfsdk:"upstream_location" json:"upstreamLocation,omitempty"`
		UpstreamPort           *int64  `tfsdk:"upstream_port" json:"upstreamPort,omitempty"`
		WorkloadSelector       *struct {
			Labels *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		} `tfsdk:"workload_selector" json:"workloadSelector,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *LeaksignalComClusterLeaksignalIstioV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_leaksignal_com_cluster_leaksignal_istio_v1_manifest"
}

func (r *LeaksignalComClusterLeaksignalIstioV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Deploy LeakSignal Proxy in all istio-enabled namespaces, can be overriden by local LeaksignalIstios.",
		MarkdownDescription: "Deploy LeakSignal Proxy in all istio-enabled namespaces, can be overriden by local LeaksignalIstios.",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"api_key": schema.StringAttribute{
						Description:         "API Key from the LeakSignal Command dashboard. Alternatively, the deployment name from LeakAgent.",
						MarkdownDescription: "API Key from the LeakSignal Command dashboard. Alternatively, the deployment name from LeakAgent.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"ca_bundle": schema.StringAttribute{
						Description:         "Location of CA bundle in istio-proxy. Default is '/etc/ssl/certs/ca-certificates.crt' which is suitable for Istio. OpenShift Service Mesh requires '/etc/ssl/certs/ca-bundle.crt'.",
						MarkdownDescription: "Location of CA bundle in istio-proxy. Default is '/etc/ssl/certs/ca-certificates.crt' which is suitable for Istio. OpenShift Service Mesh requires '/etc/ssl/certs/ca-bundle.crt'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_streaming": schema.BoolAttribute{
						Description:         "If 'true' (default), then L4 streams are also scanned by LeakSignal Proxy.",
						MarkdownDescription: "If 'true' (default), then L4 streams are also scanned by LeakSignal Proxy.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"fail_open": schema.BoolAttribute{
						Description:         "If 'true' (default), if LeakSignal Proxy has a failure, then all traffic is routed around it.",
						MarkdownDescription: "If 'true' (default), if LeakSignal Proxy has a failure, then all traffic is routed around it.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"grpc_mode": schema.StringAttribute{
						Description:         "Whether to use Google GRPC or Envoy GRPC for WASM deployments.",
						MarkdownDescription: "Whether to use Google GRPC or Envoy GRPC for WASM deployments.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("default", "envoy"),
						},
					},

					"native": schema.BoolAttribute{
						Description:         "If 'true' (not default), istio-proxy containers are updated to a corresponding image with support for dynamic plugins, and the native LeakSignal Proxy module is installed.",
						MarkdownDescription: "If 'true' (not default), istio-proxy containers are updated to a corresponding image with support for dynamic plugins, and the native LeakSignal Proxy module is installed.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"native_proxy_memory_limit": schema.StringAttribute{
						Description:         "Alternative memory limit for Istio sidecars running native modules. Useful to mitigate a surge of memory usage when loading the proxy.",
						MarkdownDescription: "Alternative memory limit for Istio sidecars running native modules. Useful to mitigate a surge of memory usage when loading the proxy.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"native_repo": schema.StringAttribute{
						Description:         "Default is 'leaksignal/istio-proxy'. If no tag is specified, it is inferred from the existing proxy image on each given pod.",
						MarkdownDescription: "Default is 'leaksignal/istio-proxy'. If no tag is specified, it is inferred from the existing proxy image on each given pod.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"proxy_hash": schema.StringAttribute{
						Description:         "Hash of the downloaded bundle for LeakSignal Proxy. Will depend on your version and deployment mechanism (nginx, envoy, WASM).",
						MarkdownDescription: "Hash of the downloaded bundle for LeakSignal Proxy. Will depend on your version and deployment mechanism (nginx, envoy, WASM).",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"proxy_prefix": schema.StringAttribute{
						Description:         "Prefix of binary to pull. Defaults to 's3/leakproxy'. For LeakAgent deployments, use 'proxy'.",
						MarkdownDescription: "Prefix of binary to pull. Defaults to 's3/leakproxy'. For LeakAgent deployments, use 'proxy'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"proxy_pull_location": schema.StringAttribute{
						Description:         "Format 'https?://domain(:port)?/'. Defaults to 'https://leakproxy.s3.us-west-2.amazonaws.com/'.",
						MarkdownDescription: "Format 'https?://domain(:port)?/'. Defaults to 'https://leakproxy.s3.us-west-2.amazonaws.com/'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"proxy_version": schema.StringAttribute{
						Description:         "Version string for LeakSignal Proxy deployment.",
						MarkdownDescription: "Version string for LeakSignal Proxy deployment.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"refresh_pods_on_stale": schema.BoolAttribute{
						Description:         "Detects pods that should have leaksignal deployed, but dont, and restarts them.",
						MarkdownDescription: "Detects pods that should have leaksignal deployed, but dont, and restarts them.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"refresh_pods_on_update": schema.BoolAttribute{
						Description:         "For WASM mode, redeploys all pods with Istio sidecars affected by a LeakSignal Proxy upgrade. This provides more consistent behavior. Default is 'true'.",
						MarkdownDescription: "For WASM mode, redeploys all pods with Istio sidecars affected by a LeakSignal Proxy upgrade. This provides more consistent behavior. Default is 'true'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tls": schema.BoolAttribute{
						Description:         "If 'true' (default), TLS/HTTPS is used for telemetry upload and downloading LeakSignal Proxy. LeakAgent is usually 'false'.",
						MarkdownDescription: "If 'true' (default), TLS/HTTPS is used for telemetry upload and downloading LeakSignal Proxy. LeakAgent is usually 'false'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"upstream_location": schema.StringAttribute{
						Description:         "Hostname of upstream location to send metrics to. Default is 'ingestion.app.leaksignal.com'.",
						MarkdownDescription: "Hostname of upstream location to send metrics to. Default is 'ingestion.app.leaksignal.com'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"upstream_port": schema.Int64Attribute{
						Description:         "Port of upstream ingestion. Defaults to 80/443 depending on 'tls'. Recommended 8121 for LeakAgent.",
						MarkdownDescription: "Port of upstream ingestion. Defaults to 80/443 depending on 'tls'. Recommended 8121 for LeakAgent.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"workload_selector": schema.SingleNestedAttribute{
						Description:         "Pod selector for workloads.",
						MarkdownDescription: "Pod selector for workloads.",
						Attributes: map[string]schema.Attribute{
							"labels": schema.MapAttribute{
								Description:         "Labels to match any pod before deploying LeakSignal.",
								MarkdownDescription: "Labels to match any pod before deploying LeakSignal.",
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *LeaksignalComClusterLeaksignalIstioV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_leaksignal_com_cluster_leaksignal_istio_v1_manifest")

	var model LeaksignalComClusterLeaksignalIstioV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("leaksignal.com/v1")
	model.Kind = pointer.String("ClusterLeaksignalIstio")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
