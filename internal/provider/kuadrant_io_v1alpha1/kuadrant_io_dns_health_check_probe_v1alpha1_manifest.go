/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kuadrant_io_v1alpha1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &KuadrantIoDnshealthCheckProbeV1Alpha1Manifest{}
)

func NewKuadrantIoDnshealthCheckProbeV1Alpha1Manifest() datasource.DataSource {
	return &KuadrantIoDnshealthCheckProbeV1Alpha1Manifest{}
}

type KuadrantIoDnshealthCheckProbeV1Alpha1Manifest struct{}

type KuadrantIoDnshealthCheckProbeV1Alpha1ManifestData struct {
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
		AdditionalHeadersRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"additional_headers_ref" json:"additionalHeadersRef,omitempty"`
		Address                  *string `tfsdk:"address" json:"address,omitempty"`
		AllowInsecureCertificate *bool   `tfsdk:"allow_insecure_certificate" json:"allowInsecureCertificate,omitempty"`
		FailureThreshold         *int64  `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
		Hostname                 *string `tfsdk:"hostname" json:"hostname,omitempty"`
		Interval                 *string `tfsdk:"interval" json:"interval,omitempty"`
		Path                     *string `tfsdk:"path" json:"path,omitempty"`
		Port                     *int64  `tfsdk:"port" json:"port,omitempty"`
		Protocol                 *string `tfsdk:"protocol" json:"protocol,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KuadrantIoDnshealthCheckProbeV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kuadrant_io_dns_health_check_probe_v1alpha1_manifest"
}

func (r *KuadrantIoDnshealthCheckProbeV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DNSHealthCheckProbe is the Schema for the dnshealthcheckprobes API",
		MarkdownDescription: "DNSHealthCheckProbe is the Schema for the dnshealthcheckprobes API",
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
				Description:         "DNSHealthCheckProbeSpec defines the desired state of DNSHealthCheckProbe",
				MarkdownDescription: "DNSHealthCheckProbeSpec defines the desired state of DNSHealthCheckProbe",
				Attributes: map[string]schema.Attribute{
					"additional_headers_ref": schema.SingleNestedAttribute{
						Description:         "AdditionalHeadersRef refers to a secret that contains extra headers to send in the probe request, this is primarily useful if an authentication token is required by the endpoint.",
						MarkdownDescription: "AdditionalHeadersRef refers to a secret that contains extra headers to send in the probe request, this is primarily useful if an authentication token is required by the endpoint.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
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

					"address": schema.StringAttribute{
						Description:         "Address to connect to the host on (IP Address (A Record) or hostname (CNAME)).",
						MarkdownDescription: "Address to connect to the host on (IP Address (A Record) or hostname (CNAME)).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"allow_insecure_certificate": schema.BoolAttribute{
						Description:         "AllowInsecureCertificate will instruct the health check probe to not fail on a self-signed or otherwise invalid SSL certificate this is primarily used in development or testing environments and is set by the --insecure-health-checks flag",
						MarkdownDescription: "AllowInsecureCertificate will instruct the health check probe to not fail on a self-signed or otherwise invalid SSL certificate this is primarily used in development or testing environments and is set by the --insecure-health-checks flag",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"failure_threshold": schema.Int64Attribute{
						Description:         "FailureThreshold is a limit of consecutive failures that must occur for a host to be considered unhealthy",
						MarkdownDescription: "FailureThreshold is a limit of consecutive failures that must occur for a host to be considered unhealthy",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"hostname": schema.StringAttribute{
						Description:         "Hostname is the value sent in the host header, to route the request to the correct service Represents a root host of the parent DNS Record.",
						MarkdownDescription: "Hostname is the value sent in the host header, to route the request to the correct service Represents a root host of the parent DNS Record.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"interval": schema.StringAttribute{
						Description:         "Interval defines how frequently this probe should execute",
						MarkdownDescription: "Interval defines how frequently this probe should execute",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"path": schema.StringAttribute{
						Description:         "Path is the path to append to the host to reach the expected health check. Must start with '?' or '/', contain only valid URL characters and end with alphanumeric char or '/'. For example '/' or '/healthz' are common",
						MarkdownDescription: "Path is the path to append to the host to reach the expected health check. Must start with '?' or '/', contain only valid URL characters and end with alphanumeric char or '/'. For example '/' or '/healthz' are common",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\?|\/)[\w\-.~:\/?#\[\]@!$&'()*+,;=]+(?:[a-zA-Z0-9]|\/){1}$`), ""),
						},
					},

					"port": schema.Int64Attribute{
						Description:         "Port to connect to the host on. Must be either 80, 443 or 1024-49151",
						MarkdownDescription: "Port to connect to the host on. Must be either 80, 443 or 1024-49151",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"protocol": schema.StringAttribute{
						Description:         "Protocol to use when connecting to the host, valid values are 'HTTP' or 'HTTPS'",
						MarkdownDescription: "Protocol to use when connecting to the host, valid values are 'HTTP' or 'HTTPS'",
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

func (r *KuadrantIoDnshealthCheckProbeV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kuadrant_io_dns_health_check_probe_v1alpha1_manifest")

	var model KuadrantIoDnshealthCheckProbeV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("kuadrant.io/v1alpha1")
	model.Kind = pointer.String("DNSHealthCheckProbe")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
