/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package console_openshift_io_v1alpha1

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
	_ datasource.DataSource = &ConsoleOpenshiftIoConsolePluginV1Alpha1Manifest{}
)

func NewConsoleOpenshiftIoConsolePluginV1Alpha1Manifest() datasource.DataSource {
	return &ConsoleOpenshiftIoConsolePluginV1Alpha1Manifest{}
}

type ConsoleOpenshiftIoConsolePluginV1Alpha1Manifest struct{}

type ConsoleOpenshiftIoConsolePluginV1Alpha1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		DisplayName *string `tfsdk:"display_name" json:"displayName,omitempty"`
		Proxy       *[]struct {
			Alias         *string `tfsdk:"alias" json:"alias,omitempty"`
			Authorize     *bool   `tfsdk:"authorize" json:"authorize,omitempty"`
			CaCertificate *string `tfsdk:"ca_certificate" json:"caCertificate,omitempty"`
			Service       *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				Port      *int64  `tfsdk:"port" json:"port,omitempty"`
			} `tfsdk:"service" json:"service,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"proxy" json:"proxy,omitempty"`
		Service *struct {
			BasePath  *string `tfsdk:"base_path" json:"basePath,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Port      *int64  `tfsdk:"port" json:"port,omitempty"`
		} `tfsdk:"service" json:"service,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ConsoleOpenshiftIoConsolePluginV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_console_openshift_io_console_plugin_v1alpha1_manifest"
}

func (r *ConsoleOpenshiftIoConsolePluginV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ConsolePlugin is an extension for customizing OpenShift web console by dynamically loading code from another service running on the cluster.  Compatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
		MarkdownDescription: "ConsolePlugin is an extension for customizing OpenShift web console by dynamically loading code from another service running on the cluster.  Compatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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
				Description:         "ConsolePluginSpec is the desired plugin configuration.",
				MarkdownDescription: "ConsolePluginSpec is the desired plugin configuration.",
				Attributes: map[string]schema.Attribute{
					"display_name": schema.StringAttribute{
						Description:         "displayName is the display name of the plugin.",
						MarkdownDescription: "displayName is the display name of the plugin.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
						},
					},

					"proxy": schema.ListNestedAttribute{
						Description:         "proxy is a list of proxies that describe various service type to which the plugin needs to connect to.",
						MarkdownDescription: "proxy is a list of proxies that describe various service type to which the plugin needs to connect to.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"alias": schema.StringAttribute{
									Description:         "alias is a proxy name that identifies the plugin's proxy. An alias name should be unique per plugin. The console backend exposes following proxy endpoint:  /api/proxy/plugin/<plugin-name>/<proxy-alias>/<request-path>?<optional-query-parameters>  Request example path:  /api/proxy/plugin/acm/search/pods?namespace=openshift-apiserver",
									MarkdownDescription: "alias is a proxy name that identifies the plugin's proxy. An alias name should be unique per plugin. The console backend exposes following proxy endpoint:  /api/proxy/plugin/<plugin-name>/<proxy-alias>/<request-path>?<optional-query-parameters>  Request example path:  /api/proxy/plugin/acm/search/pods?namespace=openshift-apiserver",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(128),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z0-9-_]+$`), ""),
									},
								},

								"authorize": schema.BoolAttribute{
									Description:         "authorize indicates if the proxied request should contain the logged-in user's OpenShift access token in the 'Authorization' request header. For example:  Authorization: Bearer sha256~kV46hPnEYhCWFnB85r5NrprAxggzgb6GOeLbgcKNsH0  By default the access token is not part of the proxied request.",
									MarkdownDescription: "authorize indicates if the proxied request should contain the logged-in user's OpenShift access token in the 'Authorization' request header. For example:  Authorization: Bearer sha256~kV46hPnEYhCWFnB85r5NrprAxggzgb6GOeLbgcKNsH0  By default the access token is not part of the proxied request.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ca_certificate": schema.StringAttribute{
									Description:         "caCertificate provides the cert authority certificate contents, in case the proxied Service is using custom service CA. By default, the service CA bundle provided by the service-ca operator is used.",
									MarkdownDescription: "caCertificate provides the cert authority certificate contents, in case the proxied Service is using custom service CA. By default, the service CA bundle provided by the service-ca operator is used.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.RegexMatches(regexp.MustCompile(`^-----BEGIN CERTIFICATE-----([\s\S]*)-----END CERTIFICATE-----\s?$`), ""),
									},
								},

								"service": schema.SingleNestedAttribute{
									Description:         "service is an in-cluster Service that the plugin will connect to. The Service must use HTTPS. The console backend exposes an endpoint in order to proxy communication between the plugin and the Service. Note: service field is required for now, since currently only 'Service' type is supported.",
									MarkdownDescription: "service is an in-cluster Service that the plugin will connect to. The Service must use HTTPS. The console backend exposes an endpoint in order to proxy communication between the plugin and the Service. Note: service field is required for now, since currently only 'Service' type is supported.",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "name of Service that the plugin needs to connect to.",
											MarkdownDescription: "name of Service that the plugin needs to connect to.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
												stringvalidator.LengthAtMost(128),
											},
										},

										"namespace": schema.StringAttribute{
											Description:         "namespace of Service that the plugin needs to connect to",
											MarkdownDescription: "namespace of Service that the plugin needs to connect to",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
												stringvalidator.LengthAtMost(128),
											},
										},

										"port": schema.Int64Attribute{
											Description:         "port on which the Service that the plugin needs to connect to is listening on.",
											MarkdownDescription: "port on which the Service that the plugin needs to connect to is listening on.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(1),
												int64validator.AtMost(65535),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"type": schema.StringAttribute{
									Description:         "type is the type of the console plugin's proxy. Currently only 'Service' is supported.",
									MarkdownDescription: "type is the type of the console plugin's proxy. Currently only 'Service' is supported.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.RegexMatches(regexp.MustCompile(`^(Service)$`), ""),
									},
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"service": schema.SingleNestedAttribute{
						Description:         "service is a Kubernetes Service that exposes the plugin using a deployment with an HTTP server. The Service must use HTTPS and Service serving certificate. The console backend will proxy the plugins assets from the Service using the service CA bundle.",
						MarkdownDescription: "service is a Kubernetes Service that exposes the plugin using a deployment with an HTTP server. The Service must use HTTPS and Service serving certificate. The console backend will proxy the plugins assets from the Service using the service CA bundle.",
						Attributes: map[string]schema.Attribute{
							"base_path": schema.StringAttribute{
								Description:         "basePath is the path to the plugin's assets. The primary asset it the manifest file called 'plugin-manifest.json', which is a JSON document that contains metadata about the plugin and the extensions.",
								MarkdownDescription: "basePath is the path to the plugin's assets. The primary asset it the manifest file called 'plugin-manifest.json', which is a JSON document that contains metadata about the plugin and the extensions.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.RegexMatches(regexp.MustCompile(`^/`), ""),
								},
							},

							"name": schema.StringAttribute{
								Description:         "name of Service that is serving the plugin assets.",
								MarkdownDescription: "name of Service that is serving the plugin assets.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(128),
								},
							},

							"namespace": schema.StringAttribute{
								Description:         "namespace of Service that is serving the plugin assets.",
								MarkdownDescription: "namespace of Service that is serving the plugin assets.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(128),
								},
							},

							"port": schema.Int64Attribute{
								Description:         "port on which the Service that is serving the plugin is listening to.",
								MarkdownDescription: "port on which the Service that is serving the plugin is listening to.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(65535),
								},
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
		},
	}
}

func (r *ConsoleOpenshiftIoConsolePluginV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_console_openshift_io_console_plugin_v1alpha1_manifest")

	var model ConsoleOpenshiftIoConsolePluginV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("console.openshift.io/v1alpha1")
	model.Kind = pointer.String("ConsolePlugin")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
