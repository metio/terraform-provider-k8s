/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package extensions_istio_io_v1alpha1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	_ datasource.DataSource = &ExtensionsIstioIoWasmPluginV1Alpha1Manifest{}
)

func NewExtensionsIstioIoWasmPluginV1Alpha1Manifest() datasource.DataSource {
	return &ExtensionsIstioIoWasmPluginV1Alpha1Manifest{}
}

type ExtensionsIstioIoWasmPluginV1Alpha1Manifest struct{}

type ExtensionsIstioIoWasmPluginV1Alpha1ManifestData struct {
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
		TargetRefs *[]struct {
			Group     *string `tfsdk:"group" json:"group,omitempty"`
			Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"target_refs" json:"targetRefs,omitempty"`
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

func (r *ExtensionsIstioIoWasmPluginV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_extensions_istio_io_wasm_plugin_v1alpha1_manifest"
}

func (r *ExtensionsIstioIoWasmPluginV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "Extend the functionality provided by the Istio proxy through WebAssembly filters. See more details at: https://istio.io/docs/reference/config/proxy_extensions/wasm-plugin.html",
				MarkdownDescription: "Extend the functionality provided by the Istio proxy through WebAssembly filters. See more details at: https://istio.io/docs/reference/config/proxy_extensions/wasm-plugin.html",
				Attributes: map[string]schema.Attribute{
					"fail_strategy": schema.StringAttribute{
						Description:         "Specifies the failure behavior for the plugin due to fatal errors.Valid Options: FAIL_CLOSE, FAIL_OPEN",
						MarkdownDescription: "Specifies the failure behavior for the plugin due to fatal errors.Valid Options: FAIL_CLOSE, FAIL_OPEN",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("FAIL_CLOSE", "FAIL_OPEN"),
						},
					},

					"image_pull_policy": schema.StringAttribute{
						Description:         "The pull behaviour to be applied when fetching Wasm module by either OCI image or 'http/https'.Valid Options: IfNotPresent, Always",
						MarkdownDescription: "The pull behaviour to be applied when fetching Wasm module by either OCI image or 'http/https'.Valid Options: IfNotPresent, Always",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("UNSPECIFIED_POLICY", "IfNotPresent", "Always"),
						},
					},

					"image_pull_secret": schema.StringAttribute{
						Description:         "Credentials to use for OCI image pulling.",
						MarkdownDescription: "Credentials to use for OCI image pulling.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
							stringvalidator.LengthAtMost(253),
						},
					},

					"match": schema.ListNestedAttribute{
						Description:         "Specifies the criteria to determine which traffic is passed to WasmPlugin.",
						MarkdownDescription: "Specifies the criteria to determine which traffic is passed to WasmPlugin.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"mode": schema.StringAttribute{
									Description:         "Criteria for selecting traffic by their direction.Valid Options: CLIENT, SERVER, CLIENT_AND_SERVER",
									MarkdownDescription: "Criteria for selecting traffic by their direction.Valid Options: CLIENT, SERVER, CLIENT_AND_SERVER",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("UNDEFINED", "CLIENT", "SERVER", "CLIENT_AND_SERVER"),
									},
								},

								"ports": schema.ListNestedAttribute{
									Description:         "Criteria for selecting traffic by their destination port.",
									MarkdownDescription: "Criteria for selecting traffic by their destination port.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"number": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
													int64validator.AtMost(65535),
												},
											},
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

					"phase": schema.StringAttribute{
						Description:         "Determines where in the filter chain this 'WasmPlugin' is to be injected.Valid Options: AUTHN, AUTHZ, STATS",
						MarkdownDescription: "Determines where in the filter chain this 'WasmPlugin' is to be injected.Valid Options: AUTHN, AUTHZ, STATS",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("UNSPECIFIED_PHASE", "AUTHN", "AUTHZ", "STATS"),
						},
					},

					"plugin_config": schema.MapAttribute{
						Description:         "The configuration that will be passed on to the plugin.",
						MarkdownDescription: "The configuration that will be passed on to the plugin.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"plugin_name": schema.StringAttribute{
						Description:         "The plugin name to be used in the Envoy configuration (used to be called 'rootID').",
						MarkdownDescription: "The plugin name to be used in the Envoy configuration (used to be called 'rootID').",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
							stringvalidator.LengthAtMost(256),
						},
					},

					"priority": schema.Int64Attribute{
						Description:         "Determines ordering of 'WasmPlugins' in the same 'phase'.",
						MarkdownDescription: "Determines ordering of 'WasmPlugins' in the same 'phase'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"selector": schema.SingleNestedAttribute{
						Description:         "Criteria used to select the specific set of pods/VMs on which this plugin configuration should be applied.",
						MarkdownDescription: "Criteria used to select the specific set of pods/VMs on which this plugin configuration should be applied.",
						Attributes: map[string]schema.Attribute{
							"match_labels": schema.MapAttribute{
								Description:         "One or more labels that indicate a specific set of pods/VMs on which a policy should be applied.",
								MarkdownDescription: "One or more labels that indicate a specific set of pods/VMs on which a policy should be applied.",
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

					"sha256": schema.StringAttribute{
						Description:         "SHA256 checksum that will be used to verify Wasm module or OCI container.",
						MarkdownDescription: "SHA256 checksum that will be used to verify Wasm module or OCI container.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`(^$|^[a-f0-9]{64}$)`), ""),
						},
					},

					"target_ref": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"group": schema.StringAttribute{
								Description:         "group is the group of the target resource.",
								MarkdownDescription: "group is the group of the target resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtMost(253),
									stringvalidator.RegexMatches(regexp.MustCompile(`^$|^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
								},
							},

							"kind": schema.StringAttribute{
								Description:         "kind is kind of the target resource.",
								MarkdownDescription: "kind is kind of the target resource.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(63),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z]([-a-zA-Z0-9]*[a-zA-Z0-9])?$`), ""),
								},
							},

							"name": schema.StringAttribute{
								Description:         "name is the name of the target resource.",
								MarkdownDescription: "name is the name of the target resource.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(253),
								},
							},

							"namespace": schema.StringAttribute{
								Description:         "namespace is the namespace of the referent.",
								MarkdownDescription: "namespace is the namespace of the referent.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"target_refs": schema.ListNestedAttribute{
						Description:         "Optional.",
						MarkdownDescription: "Optional.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"group": schema.StringAttribute{
									Description:         "group is the group of the target resource.",
									MarkdownDescription: "group is the group of the target resource.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(253),
										stringvalidator.RegexMatches(regexp.MustCompile(`^$|^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
									},
								},

								"kind": schema.StringAttribute{
									Description:         "kind is kind of the target resource.",
									MarkdownDescription: "kind is kind of the target resource.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(63),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z]([-a-zA-Z0-9]*[a-zA-Z0-9])?$`), ""),
									},
								},

								"name": schema.StringAttribute{
									Description:         "name is the name of the target resource.",
									MarkdownDescription: "name is the name of the target resource.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(253),
									},
								},

								"namespace": schema.StringAttribute{
									Description:         "namespace is the namespace of the referent.",
									MarkdownDescription: "namespace is the namespace of the referent.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"type": schema.StringAttribute{
						Description:         "Specifies the type of Wasm Extension to be used.Valid Options: HTTP, NETWORK",
						MarkdownDescription: "Specifies the type of Wasm Extension to be used.Valid Options: HTTP, NETWORK",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("UNSPECIFIED_PLUGIN_TYPE", "HTTP", "NETWORK"),
						},
					},

					"url": schema.StringAttribute{
						Description:         "URL of a Wasm module or OCI container.",
						MarkdownDescription: "URL of a Wasm module or OCI container.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
						},
					},

					"verification_key": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
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
											Description:         "Name of the environment variable.",
											MarkdownDescription: "Name of the environment variable.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
												stringvalidator.LengthAtMost(256),
											},
										},

										"value": schema.StringAttribute{
											Description:         "Value for the environment variable.",
											MarkdownDescription: "Value for the environment variable.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(2048),
											},
										},

										"value_from": schema.StringAttribute{
											Description:         "Source for the environment variable's value.Valid Options: INLINE, HOST",
											MarkdownDescription: "Source for the environment variable's value.Valid Options: INLINE, HOST",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("INLINE", "HOST"),
											},
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
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *ExtensionsIstioIoWasmPluginV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_extensions_istio_io_wasm_plugin_v1alpha1_manifest")

	var model ExtensionsIstioIoWasmPluginV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("extensions.istio.io/v1alpha1")
	model.Kind = pointer.String("WasmPlugin")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
