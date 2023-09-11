/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package extensions_istio_io_v1alpha1

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
	_ datasource.DataSource              = &ExtensionsIstioIoWasmPluginV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &ExtensionsIstioIoWasmPluginV1Alpha1DataSource{}
)

func NewExtensionsIstioIoWasmPluginV1Alpha1DataSource() datasource.DataSource {
	return &ExtensionsIstioIoWasmPluginV1Alpha1DataSource{}
}

type ExtensionsIstioIoWasmPluginV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type ExtensionsIstioIoWasmPluginV1Alpha1DataSourceData struct {
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
		FailStrategy    *string `tfsdk:"fail_strategy" json:"failStrategy,omitempty"`
		ImagePullPolicy *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
		ImagePullSecret *string `tfsdk:"image_pull_secret" json:"imagePullSecret,omitempty"`
		Match           *[]struct {
			Mode  *string `tfsdk:"mode" json:"mode,omitempty"`
			Ports *[]struct {
				Number *int64 `tfsdk:"number" json:"number,omitempty"`
			} `tfsdk:"ports" json:"ports,omitempty"`
		} `tfsdk:"match" json:"match,omitempty"`
		Phase        *string            `tfsdk:"phase" json:"phase,omitempty"`
		PluginConfig *map[string]string `tfsdk:"plugin_config" json:"pluginConfig,omitempty"`
		PluginName   *string            `tfsdk:"plugin_name" json:"pluginName,omitempty"`
		Priority     *int64             `tfsdk:"priority" json:"priority,omitempty"`
		Selector     *struct {
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"selector" json:"selector,omitempty"`
		Sha256    *string `tfsdk:"sha256" json:"sha256,omitempty"`
		TargetRef *struct {
			Group     *string `tfsdk:"group" json:"group,omitempty"`
			Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"target_ref" json:"targetRef,omitempty"`
		Type            *string `tfsdk:"type" json:"type,omitempty"`
		Url             *string `tfsdk:"url" json:"url,omitempty"`
		VerificationKey *string `tfsdk:"verification_key" json:"verificationKey,omitempty"`
		VmConfig        *struct {
			Env *[]struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *string `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"env" json:"env,omitempty"`
		} `tfsdk:"vm_config" json:"vmConfig,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ExtensionsIstioIoWasmPluginV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_extensions_istio_io_wasm_plugin_v1alpha1"
}

func (r *ExtensionsIstioIoWasmPluginV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "Extend the functionality provided by the Istio proxy through WebAssembly filters. See more details at: https://istio.io/docs/reference/config/proxy_extensions/wasm-plugin.html",
				MarkdownDescription: "Extend the functionality provided by the Istio proxy through WebAssembly filters. See more details at: https://istio.io/docs/reference/config/proxy_extensions/wasm-plugin.html",
				Attributes: map[string]schema.Attribute{
					"fail_strategy": schema.StringAttribute{
						Description:         "Specifies the failure behavior for the plugin due to fatal errors.",
						MarkdownDescription: "Specifies the failure behavior for the plugin due to fatal errors.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"image_pull_policy": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"image_pull_secret": schema.StringAttribute{
						Description:         "Credentials to use for OCI image pulling.",
						MarkdownDescription: "Credentials to use for OCI image pulling.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"match": schema.ListNestedAttribute{
						Description:         "Specifies the criteria to determine which traffic is passed to WasmPlugin.",
						MarkdownDescription: "Specifies the criteria to determine which traffic is passed to WasmPlugin.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"mode": schema.StringAttribute{
									Description:         "Criteria for selecting traffic by their direction.",
									MarkdownDescription: "Criteria for selecting traffic by their direction.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"ports": schema.ListNestedAttribute{
									Description:         "Criteria for selecting traffic by their destination port.",
									MarkdownDescription: "Criteria for selecting traffic by their destination port.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"number": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"phase": schema.StringAttribute{
						Description:         "Determines where in the filter chain this 'WasmPlugin' is to be injected.",
						MarkdownDescription: "Determines where in the filter chain this 'WasmPlugin' is to be injected.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"plugin_config": schema.MapAttribute{
						Description:         "The configuration that will be passed on to the plugin.",
						MarkdownDescription: "The configuration that will be passed on to the plugin.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"plugin_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"priority": schema.Int64Attribute{
						Description:         "Determines ordering of 'WasmPlugins' in the same 'phase'.",
						MarkdownDescription: "Determines ordering of 'WasmPlugins' in the same 'phase'.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"selector": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"match_labels": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"sha256": schema.StringAttribute{
						Description:         "SHA256 checksum that will be used to verify Wasm module or OCI container.",
						MarkdownDescription: "SHA256 checksum that will be used to verify Wasm module or OCI container.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"target_ref": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"group": schema.StringAttribute{
								Description:         "group is the group of the target resource.",
								MarkdownDescription: "group is the group of the target resource.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"kind": schema.StringAttribute{
								Description:         "kind is kind of the target resource.",
								MarkdownDescription: "kind is kind of the target resource.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "name is the name of the target resource.",
								MarkdownDescription: "name is the name of the target resource.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"namespace": schema.StringAttribute{
								Description:         "namespace is the namespace of the referent.",
								MarkdownDescription: "namespace is the namespace of the referent.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"type": schema.StringAttribute{
						Description:         "Specifies the type of Wasm Extension to be used.",
						MarkdownDescription: "Specifies the type of Wasm Extension to be used.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"url": schema.StringAttribute{
						Description:         "URL of a Wasm module or OCI container.",
						MarkdownDescription: "URL of a Wasm module or OCI container.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"verification_key": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"vm_config": schema.SingleNestedAttribute{
						Description:         "Configuration for a Wasm VM.",
						MarkdownDescription: "Configuration for a Wasm VM.",
						Attributes: map[string]schema.Attribute{
							"env": schema.ListNestedAttribute{
								Description:         "Specifies environment variables to be injected to this VM.",
								MarkdownDescription: "Specifies environment variables to be injected to this VM.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"value": schema.StringAttribute{
											Description:         "Value for the environment variable.",
											MarkdownDescription: "Value for the environment variable.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"value_from": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
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
		},
	}
}

func (r *ExtensionsIstioIoWasmPluginV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *ExtensionsIstioIoWasmPluginV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_extensions_istio_io_wasm_plugin_v1alpha1")

	var data ExtensionsIstioIoWasmPluginV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "extensions.istio.io", Version: "v1alpha1", Resource: "wasmplugins"}).
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

	var readResponse ExtensionsIstioIoWasmPluginV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("extensions.istio.io/v1alpha1")
	data.Kind = pointer.String("WasmPlugin")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
