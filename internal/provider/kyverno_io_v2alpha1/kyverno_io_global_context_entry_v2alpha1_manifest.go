/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kyverno_io_v2alpha1

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
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &KyvernoIoGlobalContextEntryV2Alpha1Manifest{}
)

func NewKyvernoIoGlobalContextEntryV2Alpha1Manifest() datasource.DataSource {
	return &KyvernoIoGlobalContextEntryV2Alpha1Manifest{}
}

type KyvernoIoGlobalContextEntryV2Alpha1Manifest struct{}

type KyvernoIoGlobalContextEntryV2Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ApiCall *struct {
			Data *[]struct {
				Key   *string            `tfsdk:"key" json:"key,omitempty"`
				Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"data" json:"data,omitempty"`
			Method          *string `tfsdk:"method" json:"method,omitempty"`
			RefreshInterval *string `tfsdk:"refresh_interval" json:"refreshInterval,omitempty"`
			RetryLimit      *int64  `tfsdk:"retry_limit" json:"retryLimit,omitempty"`
			Service         *struct {
				CaBundle *string `tfsdk:"ca_bundle" json:"caBundle,omitempty"`
				Url      *string `tfsdk:"url" json:"url,omitempty"`
			} `tfsdk:"service" json:"service,omitempty"`
			UrlPath *string `tfsdk:"url_path" json:"urlPath,omitempty"`
		} `tfsdk:"api_call" json:"apiCall,omitempty"`
		KubernetesResource *struct {
			Group     *string `tfsdk:"group" json:"group,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Resource  *string `tfsdk:"resource" json:"resource,omitempty"`
			Version   *string `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"kubernetes_resource" json:"kubernetesResource,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KyvernoIoGlobalContextEntryV2Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kyverno_io_global_context_entry_v2alpha1_manifest"
}

func (r *KyvernoIoGlobalContextEntryV2Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "GlobalContextEntry declares resources to be cached.",
		MarkdownDescription: "GlobalContextEntry declares resources to be cached.",
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
				Description:         "Spec declares policy exception behaviors.",
				MarkdownDescription: "Spec declares policy exception behaviors.",
				Attributes: map[string]schema.Attribute{
					"api_call": schema.SingleNestedAttribute{
						Description:         "Stores results from an API call which will be cached.Mutually exclusive with KubernetesResource.This can be used to make calls to external (non-Kubernetes API server) services.It can also be used to make calls to the Kubernetes API server in such cases:1. A POST is needed to create a resource.2. Finer-grained control is needed. Example: To restrict the number of resources cached.",
						MarkdownDescription: "Stores results from an API call which will be cached.Mutually exclusive with KubernetesResource.This can be used to make calls to external (non-Kubernetes API server) services.It can also be used to make calls to the Kubernetes API server in such cases:1. A POST is needed to create a resource.2. Finer-grained control is needed. Example: To restrict the number of resources cached.",
						Attributes: map[string]schema.Attribute{
							"data": schema.ListNestedAttribute{
								Description:         "The data object specifies the POST data sent to the server.Only applicable when the method field is set to POST.",
								MarkdownDescription: "The data object specifies the POST data sent to the server.Only applicable when the method field is set to POST.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "Key is a unique identifier for the data value",
											MarkdownDescription: "Key is a unique identifier for the data value",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"value": schema.MapAttribute{
											Description:         "Value is the data value",
											MarkdownDescription: "Value is the data value",
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

							"method": schema.StringAttribute{
								Description:         "Method is the HTTP request type (GET or POST). Defaults to GET.",
								MarkdownDescription: "Method is the HTTP request type (GET or POST). Defaults to GET.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("GET", "POST"),
								},
							},

							"refresh_interval": schema.StringAttribute{
								Description:         "RefreshInterval defines the interval in duration at which to poll the APICall.The duration is a sequence of decimal numbers, each with optional fraction and a unit suffix,such as '300ms', '1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
								MarkdownDescription: "RefreshInterval defines the interval in duration at which to poll the APICall.The duration is a sequence of decimal numbers, each with optional fraction and a unit suffix,such as '300ms', '1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"retry_limit": schema.Int64Attribute{
								Description:         "RetryLimit defines the number of times the APICall should be retried in case of failure.",
								MarkdownDescription: "RetryLimit defines the number of times the APICall should be retried in case of failure.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
							},

							"service": schema.SingleNestedAttribute{
								Description:         "Service is an API call to a JSON web service.This is used for non-Kubernetes API server calls.It's mutually exclusive with the URLPath field.",
								MarkdownDescription: "Service is an API call to a JSON web service.This is used for non-Kubernetes API server calls.It's mutually exclusive with the URLPath field.",
								Attributes: map[string]schema.Attribute{
									"ca_bundle": schema.StringAttribute{
										Description:         "CABundle is a PEM encoded CA bundle which will be used to validatethe server certificate.",
										MarkdownDescription: "CABundle is a PEM encoded CA bundle which will be used to validatethe server certificate.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"url": schema.StringAttribute{
										Description:         "URL is the JSON web service URL. A typical form is'https://{service}.{namespace}:{port}/{path}'.",
										MarkdownDescription: "URL is the JSON web service URL. A typical form is'https://{service}.{namespace}:{port}/{path}'.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"url_path": schema.StringAttribute{
								Description:         "URLPath is the URL path to be used in the HTTP GET or POST request to theKubernetes API server (e.g. '/api/v1/namespaces' or  '/apis/apps/v1/deployments').The format required is the same format used by the 'kubectl get --raw' command.See https://kyverno.io/docs/writing-policies/external-data-sources/#variables-from-kubernetes-api-server-callsfor details.It's mutually exclusive with the Service field.",
								MarkdownDescription: "URLPath is the URL path to be used in the HTTP GET or POST request to theKubernetes API server (e.g. '/api/v1/namespaces' or  '/apis/apps/v1/deployments').The format required is the same format used by the 'kubectl get --raw' command.See https://kyverno.io/docs/writing-policies/external-data-sources/#variables-from-kubernetes-api-server-callsfor details.It's mutually exclusive with the Service field.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"kubernetes_resource": schema.SingleNestedAttribute{
						Description:         "Stores a list of Kubernetes resources which will be cached.Mutually exclusive with APICall.",
						MarkdownDescription: "Stores a list of Kubernetes resources which will be cached.Mutually exclusive with APICall.",
						Attributes: map[string]schema.Attribute{
							"group": schema.StringAttribute{
								Description:         "Group defines the group of the resource.",
								MarkdownDescription: "Group defines the group of the resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace defines the namespace of the resource. Leave empty for cluster scoped resources.If left empty for namespaced resources, all resources from all namespaces will be cached.",
								MarkdownDescription: "Namespace defines the namespace of the resource. Leave empty for cluster scoped resources.If left empty for namespaced resources, all resources from all namespaces will be cached.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resource": schema.StringAttribute{
								Description:         "Resource defines the type of the resource.Requires the pluralized form of the resource kind in lowercase. (Ex., 'deployments')",
								MarkdownDescription: "Resource defines the type of the resource.Requires the pluralized form of the resource kind in lowercase. (Ex., 'deployments')",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"version": schema.StringAttribute{
								Description:         "Version defines the version of the resource.",
								MarkdownDescription: "Version defines the version of the resource.",
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
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *KyvernoIoGlobalContextEntryV2Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kyverno_io_global_context_entry_v2alpha1_manifest")

	var model KyvernoIoGlobalContextEntryV2Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("kyverno.io/v2alpha1")
	model.Kind = pointer.String("GlobalContextEntry")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
