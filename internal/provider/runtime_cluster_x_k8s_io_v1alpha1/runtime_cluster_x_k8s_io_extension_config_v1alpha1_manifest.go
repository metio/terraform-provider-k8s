/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package runtime_cluster_x_k8s_io_v1alpha1

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
	_ datasource.DataSource = &RuntimeClusterXK8SIoExtensionConfigV1Alpha1Manifest{}
)

func NewRuntimeClusterXK8SIoExtensionConfigV1Alpha1Manifest() datasource.DataSource {
	return &RuntimeClusterXK8SIoExtensionConfigV1Alpha1Manifest{}
}

type RuntimeClusterXK8SIoExtensionConfigV1Alpha1Manifest struct{}

type RuntimeClusterXK8SIoExtensionConfigV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ClientConfig *struct {
			CaBundle *string `tfsdk:"ca_bundle" json:"caBundle,omitempty"`
			Service  *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				Path      *string `tfsdk:"path" json:"path,omitempty"`
				Port      *int64  `tfsdk:"port" json:"port,omitempty"`
			} `tfsdk:"service" json:"service,omitempty"`
			Url *string `tfsdk:"url" json:"url,omitempty"`
		} `tfsdk:"client_config" json:"clientConfig,omitempty"`
		NamespaceSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
		Settings *map[string]string `tfsdk:"settings" json:"settings,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *RuntimeClusterXK8SIoExtensionConfigV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_runtime_cluster_x_k8s_io_extension_config_v1alpha1_manifest"
}

func (r *RuntimeClusterXK8SIoExtensionConfigV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ExtensionConfig is the Schema for the ExtensionConfig API.",
		MarkdownDescription: "ExtensionConfig is the Schema for the ExtensionConfig API.",
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
				Description:         "ExtensionConfigSpec is the desired state of the ExtensionConfig",
				MarkdownDescription: "ExtensionConfigSpec is the desired state of the ExtensionConfig",
				Attributes: map[string]schema.Attribute{
					"client_config": schema.SingleNestedAttribute{
						Description:         "ClientConfig defines how to communicate with the Extension server.",
						MarkdownDescription: "ClientConfig defines how to communicate with the Extension server.",
						Attributes: map[string]schema.Attribute{
							"ca_bundle": schema.StringAttribute{
								Description:         "CABundle is a PEM encoded CA bundle which will be used to validate the Extension server's server certificate.",
								MarkdownDescription: "CABundle is a PEM encoded CA bundle which will be used to validate the Extension server's server certificate.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									validators.Base64Validator(),
								},
							},

							"service": schema.SingleNestedAttribute{
								Description:         "Service is a reference to the Kubernetes service for the Extension server. Note: Exactly one of 'url' or 'service' must be specified. If the Extension server is running within a cluster, then you should use 'service'.",
								MarkdownDescription: "Service is a reference to the Kubernetes service for the Extension server. Note: Exactly one of 'url' or 'service' must be specified. If the Extension server is running within a cluster, then you should use 'service'.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name is the name of the service.",
										MarkdownDescription: "Name is the name of the service.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "Namespace is the namespace of the service.",
										MarkdownDescription: "Namespace is the namespace of the service.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"path": schema.StringAttribute{
										Description:         "Path is an optional URL path and if present may be any string permissible in a URL. If a path is set it will be used as prefix to the hook-specific path.",
										MarkdownDescription: "Path is an optional URL path and if present may be any string permissible in a URL. If a path is set it will be used as prefix to the hook-specific path.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"port": schema.Int64Attribute{
										Description:         "Port is the port on the service that's hosting the Extension server. Defaults to 443. Port should be a valid port number (1-65535, inclusive).",
										MarkdownDescription: "Port is the port on the service that's hosting the Extension server. Defaults to 443. Port should be a valid port number (1-65535, inclusive).",
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
								Description:         "URL gives the location of the Extension server, in standard URL form ('scheme://host:port/path'). Note: Exactly one of 'url' or 'service' must be specified. The scheme must be 'https'. The 'host' should not refer to a service running in the cluster; use the 'service' field instead. A path is optional, and if present may be any string permissible in a URL. If a path is set it will be used as prefix to the hook-specific path. Attempting to use a user or basic auth e.g. 'user:password@' is not allowed. Fragments ('#...') and query parameters ('?...') are not allowed either.",
								MarkdownDescription: "URL gives the location of the Extension server, in standard URL form ('scheme://host:port/path'). Note: Exactly one of 'url' or 'service' must be specified. The scheme must be 'https'. The 'host' should not refer to a service running in the cluster; use the 'service' field instead. A path is optional, and if present may be any string permissible in a URL. If a path is set it will be used as prefix to the hook-specific path. Attempting to use a user or basic auth e.g. 'user:password@' is not allowed. Fragments ('#...') and query parameters ('?...') are not allowed either.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"namespace_selector": schema.SingleNestedAttribute{
						Description:         "NamespaceSelector decides whether to call the hook for an object based on whether the namespace for that object matches the selector. Defaults to the empty LabelSelector, which matches all objects.",
						MarkdownDescription: "NamespaceSelector decides whether to call the hook for an object based on whether the namespace for that object matches the selector. Defaults to the empty LabelSelector, which matches all objects.",
						Attributes: map[string]schema.Attribute{
							"match_expressions": schema.ListNestedAttribute{
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "key is the label key that the selector applies to.",
											MarkdownDescription: "key is the label key that the selector applies to.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											ElementType:         types.StringType,
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

							"match_labels": schema.MapAttribute{
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

					"settings": schema.MapAttribute{
						Description:         "Settings defines key value pairs to be passed to all calls to all supported RuntimeExtensions. Note: Settings can be overridden on the ClusterClass.",
						MarkdownDescription: "Settings defines key value pairs to be passed to all calls to all supported RuntimeExtensions. Note: Settings can be overridden on the ClusterClass.",
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
	}
}

func (r *RuntimeClusterXK8SIoExtensionConfigV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_runtime_cluster_x_k8s_io_extension_config_v1alpha1_manifest")

	var model RuntimeClusterXK8SIoExtensionConfigV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("runtime.cluster.x-k8s.io/v1alpha1")
	model.Kind = pointer.String("ExtensionConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
