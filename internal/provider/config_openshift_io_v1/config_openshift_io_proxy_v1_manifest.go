/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package config_openshift_io_v1

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
	_ datasource.DataSource = &ConfigOpenshiftIoProxyV1Manifest{}
)

func NewConfigOpenshiftIoProxyV1Manifest() datasource.DataSource {
	return &ConfigOpenshiftIoProxyV1Manifest{}
}

type ConfigOpenshiftIoProxyV1Manifest struct{}

type ConfigOpenshiftIoProxyV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		HttpProxy          *string   `tfsdk:"http_proxy" json:"httpProxy,omitempty"`
		HttpsProxy         *string   `tfsdk:"https_proxy" json:"httpsProxy,omitempty"`
		NoProxy            *string   `tfsdk:"no_proxy" json:"noProxy,omitempty"`
		ReadinessEndpoints *[]string `tfsdk:"readiness_endpoints" json:"readinessEndpoints,omitempty"`
		TrustedCA          *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"trusted_ca" json:"trustedCA,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ConfigOpenshiftIoProxyV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_config_openshift_io_proxy_v1_manifest"
}

func (r *ConfigOpenshiftIoProxyV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Proxy holds cluster-wide information on how to configure default proxies for the cluster. The canonical name is 'cluster'  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		MarkdownDescription: "Proxy holds cluster-wide information on how to configure default proxies for the cluster. The canonical name is 'cluster'  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
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
				Description:         "Spec holds user-settable values for the proxy configuration",
				MarkdownDescription: "Spec holds user-settable values for the proxy configuration",
				Attributes: map[string]schema.Attribute{
					"http_proxy": schema.StringAttribute{
						Description:         "httpProxy is the URL of the proxy for HTTP requests.  Empty means unset and will not result in an env var.",
						MarkdownDescription: "httpProxy is the URL of the proxy for HTTP requests.  Empty means unset and will not result in an env var.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"https_proxy": schema.StringAttribute{
						Description:         "httpsProxy is the URL of the proxy for HTTPS requests.  Empty means unset and will not result in an env var.",
						MarkdownDescription: "httpsProxy is the URL of the proxy for HTTPS requests.  Empty means unset and will not result in an env var.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"no_proxy": schema.StringAttribute{
						Description:         "noProxy is a comma-separated list of hostnames and/or CIDRs and/or IPs for which the proxy should not be used. Empty means unset and will not result in an env var.",
						MarkdownDescription: "noProxy is a comma-separated list of hostnames and/or CIDRs and/or IPs for which the proxy should not be used. Empty means unset and will not result in an env var.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"readiness_endpoints": schema.ListAttribute{
						Description:         "readinessEndpoints is a list of endpoints used to verify readiness of the proxy.",
						MarkdownDescription: "readinessEndpoints is a list of endpoints used to verify readiness of the proxy.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"trusted_ca": schema.SingleNestedAttribute{
						Description:         "trustedCA is a reference to a ConfigMap containing a CA certificate bundle. The trustedCA field should only be consumed by a proxy validator. The validator is responsible for reading the certificate bundle from the required key 'ca-bundle.crt', merging it with the system default trust bundle, and writing the merged trust bundle to a ConfigMap named 'trusted-ca-bundle' in the 'openshift-config-managed' namespace. Clients that expect to make proxy connections must use the trusted-ca-bundle for all HTTPS requests to the proxy, and may use the trusted-ca-bundle for non-proxy HTTPS requests as well.  The namespace for the ConfigMap referenced by trustedCA is 'openshift-config'. Here is an example ConfigMap (in yaml):  apiVersion: v1 kind: ConfigMap metadata: name: user-ca-bundle namespace: openshift-config data: ca-bundle.crt: | -----BEGIN CERTIFICATE----- Custom CA certificate bundle. -----END CERTIFICATE-----",
						MarkdownDescription: "trustedCA is a reference to a ConfigMap containing a CA certificate bundle. The trustedCA field should only be consumed by a proxy validator. The validator is responsible for reading the certificate bundle from the required key 'ca-bundle.crt', merging it with the system default trust bundle, and writing the merged trust bundle to a ConfigMap named 'trusted-ca-bundle' in the 'openshift-config-managed' namespace. Clients that expect to make proxy connections must use the trusted-ca-bundle for all HTTPS requests to the proxy, and may use the trusted-ca-bundle for non-proxy HTTPS requests as well.  The namespace for the ConfigMap referenced by trustedCA is 'openshift-config'. Here is an example ConfigMap (in yaml):  apiVersion: v1 kind: ConfigMap metadata: name: user-ca-bundle namespace: openshift-config data: ca-bundle.crt: | -----BEGIN CERTIFICATE----- Custom CA certificate bundle. -----END CERTIFICATE-----",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "name is the metadata.name of the referenced config map",
								MarkdownDescription: "name is the metadata.name of the referenced config map",
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

func (r *ConfigOpenshiftIoProxyV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_config_openshift_io_proxy_v1_manifest")

	var model ConfigOpenshiftIoProxyV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("config.openshift.io/v1")
	model.Kind = pointer.String("Proxy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
