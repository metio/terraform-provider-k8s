/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package config_karmada_io_v1alpha1

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
	_ datasource.DataSource              = &ConfigKarmadaIoResourceInterpreterWebhookConfigurationV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &ConfigKarmadaIoResourceInterpreterWebhookConfigurationV1Alpha1DataSource{}
)

func NewConfigKarmadaIoResourceInterpreterWebhookConfigurationV1Alpha1DataSource() datasource.DataSource {
	return &ConfigKarmadaIoResourceInterpreterWebhookConfigurationV1Alpha1DataSource{}
}

type ConfigKarmadaIoResourceInterpreterWebhookConfigurationV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type ConfigKarmadaIoResourceInterpreterWebhookConfigurationV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Webhooks *[]struct {
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
		InterpreterContextVersions *[]string `tfsdk:"interpreter_context_versions" json:"interpreterContextVersions,omitempty"`
		Name                       *string   `tfsdk:"name" json:"name,omitempty"`
		Rules                      *[]struct {
			ApiGroups   *[]string `tfsdk:"api_groups" json:"apiGroups,omitempty"`
			ApiVersions *[]string `tfsdk:"api_versions" json:"apiVersions,omitempty"`
			Kinds       *[]string `tfsdk:"kinds" json:"kinds,omitempty"`
			Operations  *[]string `tfsdk:"operations" json:"operations,omitempty"`
		} `tfsdk:"rules" json:"rules,omitempty"`
		TimeoutSeconds *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
	} `tfsdk:"webhooks" json:"webhooks,omitempty"`
}

func (r *ConfigKarmadaIoResourceInterpreterWebhookConfigurationV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_config_karmada_io_resource_interpreter_webhook_configuration_v1alpha1"
}

func (r *ConfigKarmadaIoResourceInterpreterWebhookConfigurationV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ResourceInterpreterWebhookConfiguration describes the configuration of webhooks which take the responsibility to tell karmada the details of the resource object, especially for custom resources.",
		MarkdownDescription: "ResourceInterpreterWebhookConfiguration describes the configuration of webhooks which take the responsibility to tell karmada the details of the resource object, especially for custom resources.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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

			"webhooks": schema.ListNestedAttribute{
				Description:         "Webhooks is a list of webhooks and the affected resources and operations.",
				MarkdownDescription: "Webhooks is a list of webhooks and the affected resources and operations.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"client_config": schema.SingleNestedAttribute{
							Description:         "ClientConfig defines how to communicate with the hook.",
							MarkdownDescription: "ClientConfig defines how to communicate with the hook.",
							Attributes: map[string]schema.Attribute{
								"ca_bundle": schema.StringAttribute{
									Description:         "'caBundle' is a PEM encoded CA bundle which will be used to validate the webhook's server certificate. If unspecified, system trust roots on the apiserver are used.",
									MarkdownDescription: "'caBundle' is a PEM encoded CA bundle which will be used to validate the webhook's server certificate. If unspecified, system trust roots on the apiserver are used.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"service": schema.SingleNestedAttribute{
									Description:         "'service' is a reference to the service for this webhook. Either 'service' or 'url' must be specified.  If the webhook is running within the cluster, then you should use 'service'.",
									MarkdownDescription: "'service' is a reference to the service for this webhook. Either 'service' or 'url' must be specified.  If the webhook is running within the cluster, then you should use 'service'.",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "'name' is the name of the service. Required",
											MarkdownDescription: "'name' is the name of the service. Required",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"namespace": schema.StringAttribute{
											Description:         "'namespace' is the namespace of the service. Required",
											MarkdownDescription: "'namespace' is the namespace of the service. Required",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"path": schema.StringAttribute{
											Description:         "'path' is an optional URL path which will be sent in any request to this service.",
											MarkdownDescription: "'path' is an optional URL path which will be sent in any request to this service.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"port": schema.Int64Attribute{
											Description:         "If specified, the port on the service that hosting webhook. Default to 443 for backward compatibility. 'port' should be a valid port number (1-65535, inclusive).",
											MarkdownDescription: "If specified, the port on the service that hosting webhook. Default to 443 for backward compatibility. 'port' should be a valid port number (1-65535, inclusive).",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"url": schema.StringAttribute{
									Description:         "'url' gives the location of the webhook, in standard URL form ('scheme://host:port/path'). Exactly one of 'url' or 'service' must be specified.  The 'host' should not refer to a service running in the cluster; use the 'service' field instead. The host might be resolved via external DNS in some apiservers (e.g., 'kube-apiserver' cannot resolve in-cluster DNS as that would be a layering violation). 'host' may also be an IP address.  Please note that using 'localhost' or '127.0.0.1' as a 'host' is risky unless you take great care to run this webhook on all hosts which run an apiserver which might need to make calls to this webhook. Such installs are likely to be non-portable, i.e., not easy to turn up in a new cluster.  The scheme must be 'https'; the URL must begin with 'https://'.  A path is optional, and if present may be any string permissible in a URL. You may use the path to pass an arbitrary string to the webhook, for example, a cluster identifier.  Attempting to use a user or basic auth e.g. 'user:password@' is not allowed. Fragments ('#...') and query parameters ('?...') are not allowed, either.",
									MarkdownDescription: "'url' gives the location of the webhook, in standard URL form ('scheme://host:port/path'). Exactly one of 'url' or 'service' must be specified.  The 'host' should not refer to a service running in the cluster; use the 'service' field instead. The host might be resolved via external DNS in some apiservers (e.g., 'kube-apiserver' cannot resolve in-cluster DNS as that would be a layering violation). 'host' may also be an IP address.  Please note that using 'localhost' or '127.0.0.1' as a 'host' is risky unless you take great care to run this webhook on all hosts which run an apiserver which might need to make calls to this webhook. Such installs are likely to be non-portable, i.e., not easy to turn up in a new cluster.  The scheme must be 'https'; the URL must begin with 'https://'.  A path is optional, and if present may be any string permissible in a URL. You may use the path to pass an arbitrary string to the webhook, for example, a cluster identifier.  Attempting to use a user or basic auth e.g. 'user:password@' is not allowed. Fragments ('#...') and query parameters ('?...') are not allowed, either.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
							Required: false,
							Optional: false,
							Computed: true,
						},

						"interpreter_context_versions": schema.ListAttribute{
							Description:         "InterpreterContextVersions is an ordered list of preferred 'ResourceInterpreterContext' versions the Webhook expects. Karmada will try to use first version in the list which it supports. If none of the versions specified in this list supported by Karmada, validation will fail for this object. If a persisted webhook configuration specifies allowed versions and does not include any versions known to the Karmada, calls to the webhook will fail and be subject to the failure policy.",
							MarkdownDescription: "InterpreterContextVersions is an ordered list of preferred 'ResourceInterpreterContext' versions the Webhook expects. Karmada will try to use first version in the list which it supports. If none of the versions specified in this list supported by Karmada, validation will fail for this object. If a persisted webhook configuration specifies allowed versions and does not include any versions known to the Karmada, calls to the webhook will fail and be subject to the failure policy.",
							ElementType:         types.StringType,
							Required:            false,
							Optional:            false,
							Computed:            true,
						},

						"name": schema.StringAttribute{
							Description:         "Name is the full-qualified name of the webhook.",
							MarkdownDescription: "Name is the full-qualified name of the webhook.",
							Required:            false,
							Optional:            false,
							Computed:            true,
						},

						"rules": schema.ListNestedAttribute{
							Description:         "Rules describes what operations on what resources the webhook cares about. The webhook cares about an operation if it matches any Rule.",
							MarkdownDescription: "Rules describes what operations on what resources the webhook cares about. The webhook cares about an operation if it matches any Rule.",
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"api_groups": schema.ListAttribute{
										Description:         "APIGroups is the API groups the resources belong to. '*' is all groups. If '*' is present, the length of the slice must be one. For example: ['apps', 'batch', 'example.io'] means matches 3 groups. ['*'] means matches all group  Note: The group cloud be empty, e.g the 'core' group of kubernetes, in that case use [''].",
										MarkdownDescription: "APIGroups is the API groups the resources belong to. '*' is all groups. If '*' is present, the length of the slice must be one. For example: ['apps', 'batch', 'example.io'] means matches 3 groups. ['*'] means matches all group  Note: The group cloud be empty, e.g the 'core' group of kubernetes, in that case use [''].",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"api_versions": schema.ListAttribute{
										Description:         "APIVersions is the API versions the resources belong to. '*' is all versions. If '*' is present, the length of the slice must be one. For example: ['v1alpha1', 'v1beta1'] means matches 2 versions. ['*'] means matches all versions.",
										MarkdownDescription: "APIVersions is the API versions the resources belong to. '*' is all versions. If '*' is present, the length of the slice must be one. For example: ['v1alpha1', 'v1beta1'] means matches 2 versions. ['*'] means matches all versions.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"kinds": schema.ListAttribute{
										Description:         "Kinds is a list of resources this rule applies to. If '*' is present, the length of the slice must be one. For example: ['Deployment', 'Pod'] means matches Deployment and Pod. ['*'] means apply to all resources.",
										MarkdownDescription: "Kinds is a list of resources this rule applies to. If '*' is present, the length of the slice must be one. For example: ['Deployment', 'Pod'] means matches Deployment and Pod. ['*'] means apply to all resources.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"operations": schema.ListAttribute{
										Description:         "Operations is the operations the hook cares about. If '*' is present, the length of the slice must be one.",
										MarkdownDescription: "Operations is the operations the hook cares about. If '*' is present, the length of the slice must be one.",
										ElementType:         types.StringType,
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

						"timeout_seconds": schema.Int64Attribute{
							Description:         "TimeoutSeconds specifies the timeout for this webhook. After the timeout passes, the webhook call will be ignored or the API call will fail based on the failure policy. The timeout value must be between 1 and 30 seconds. Default to 10 seconds.",
							MarkdownDescription: "TimeoutSeconds specifies the timeout for this webhook. After the timeout passes, the webhook call will be ignored or the API call will fail based on the failure policy. The timeout value must be between 1 and 30 seconds. Default to 10 seconds.",
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
	}
}

func (r *ConfigKarmadaIoResourceInterpreterWebhookConfigurationV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *ConfigKarmadaIoResourceInterpreterWebhookConfigurationV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_config_karmada_io_resource_interpreter_webhook_configuration_v1alpha1")

	var data ConfigKarmadaIoResourceInterpreterWebhookConfigurationV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "config.karmada.io", Version: "v1alpha1", Resource: "resourceinterpreterwebhookconfigurations"}).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetResourceError(err, data.Metadata.Name))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse ConfigKarmadaIoResourceInterpreterWebhookConfigurationV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("config.karmada.io/v1alpha1")
	data.Kind = pointer.String("ResourceInterpreterWebhookConfiguration")
	data.Metadata = readResponse.Metadata
	data.Webhooks = readResponse.Webhooks

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
