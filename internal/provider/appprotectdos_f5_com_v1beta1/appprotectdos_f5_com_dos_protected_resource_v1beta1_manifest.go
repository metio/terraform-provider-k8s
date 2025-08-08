/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package appprotectdos_f5_com_v1beta1

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
	_ datasource.DataSource = &AppprotectdosF5ComDosProtectedResourceV1Beta1Manifest{}
)

func NewAppprotectdosF5ComDosProtectedResourceV1Beta1Manifest() datasource.DataSource {
	return &AppprotectdosF5ComDosProtectedResourceV1Beta1Manifest{}
}

type AppprotectdosF5ComDosProtectedResourceV1Beta1Manifest struct{}

type AppprotectdosF5ComDosProtectedResourceV1Beta1ManifestData struct {
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
		ApDosMonitor *struct {
			Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
			Timeout  *int64  `tfsdk:"timeout" json:"timeout,omitempty"`
			Uri      *string `tfsdk:"uri" json:"uri,omitempty"`
		} `tfsdk:"ap_dos_monitor" json:"apDosMonitor,omitempty"`
		ApDosPolicy      *string `tfsdk:"ap_dos_policy" json:"apDosPolicy,omitempty"`
		DosAccessLogDest *string `tfsdk:"dos_access_log_dest" json:"dosAccessLogDest,omitempty"`
		DosSecurityLog   *struct {
			ApDosLogConf *string `tfsdk:"ap_dos_log_conf" json:"apDosLogConf,omitempty"`
			DosLogDest   *string `tfsdk:"dos_log_dest" json:"dosLogDest,omitempty"`
			Enable       *bool   `tfsdk:"enable" json:"enable,omitempty"`
		} `tfsdk:"dos_security_log" json:"dosSecurityLog,omitempty"`
		Enable *bool   `tfsdk:"enable" json:"enable,omitempty"`
		Name   *string `tfsdk:"name" json:"name,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppprotectdosF5ComDosProtectedResourceV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_appprotectdos_f5_com_dos_protected_resource_v1beta1_manifest"
}

func (r *AppprotectdosF5ComDosProtectedResourceV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DosProtectedResource defines a Dos protected resource.",
		MarkdownDescription: "DosProtectedResource defines a Dos protected resource.",
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
				Description:         "DosProtectedResourceSpec defines the properties and values a DosProtectedResource can have.",
				MarkdownDescription: "DosProtectedResourceSpec defines the properties and values a DosProtectedResource can have.",
				Attributes: map[string]schema.Attribute{
					"ap_dos_monitor": schema.SingleNestedAttribute{
						Description:         "ApDosMonitor is how NGINX App Protect DoS monitors the stress level of the protected object. The monitor requests are sent from localhost (127.0.0.1). Default value: URI - None, protocol - http1, timeout - NGINX App Protect DoS default.",
						MarkdownDescription: "ApDosMonitor is how NGINX App Protect DoS monitors the stress level of the protected object. The monitor requests are sent from localhost (127.0.0.1). Default value: URI - None, protocol - http1, timeout - NGINX App Protect DoS default.",
						Attributes: map[string]schema.Attribute{
							"protocol": schema.StringAttribute{
								Description:         "Protocol determines if the server listens on http1 / http2 / grpc / websocket. The default is http1.",
								MarkdownDescription: "Protocol determines if the server listens on http1 / http2 / grpc / websocket. The default is http1.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("http1", "http2", "grpc", "websocket"),
								},
							},

							"timeout": schema.Int64Attribute{
								Description:         "Timeout determines how long (in seconds) should NGINX App Protect DoS wait for a response. Default is 10 seconds for http1/http2 and 5 seconds for grpc.",
								MarkdownDescription: "Timeout determines how long (in seconds) should NGINX App Protect DoS wait for a response. Default is 10 seconds for http1/http2 and 5 seconds for grpc.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"uri": schema.StringAttribute{
								Description:         "URI is the destination to the desired protected object in the nginx.conf:",
								MarkdownDescription: "URI is the destination to the desired protected object in the nginx.conf:",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"ap_dos_policy": schema.StringAttribute{
						Description:         "ApDosPolicy is the namespace/name of a ApDosPolicy resource",
						MarkdownDescription: "ApDosPolicy is the namespace/name of a ApDosPolicy resource",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"dos_access_log_dest": schema.StringAttribute{
						Description:         "DosAccessLogDest is the network address for the access logs",
						MarkdownDescription: "DosAccessLogDest is the network address for the access logs",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"dos_security_log": schema.SingleNestedAttribute{
						Description:         "DosSecurityLog defines the security log of the DosProtectedResource.",
						MarkdownDescription: "DosSecurityLog defines the security log of the DosProtectedResource.",
						Attributes: map[string]schema.Attribute{
							"ap_dos_log_conf": schema.StringAttribute{
								Description:         "ApDosLogConf is the namespace/name of a APDosLogConf resource",
								MarkdownDescription: "ApDosLogConf is the namespace/name of a APDosLogConf resource",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dos_log_dest": schema.StringAttribute{
								Description:         "DosLogDest is the network address of a logging service, can be either IP or DNS name.",
								MarkdownDescription: "DosLogDest is the network address of a logging service, can be either IP or DNS name.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable": schema.BoolAttribute{
								Description:         "Enable enables the security logging feature if set to true",
								MarkdownDescription: "Enable enables the security logging feature if set to true",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable": schema.BoolAttribute{
						Description:         "Enable enables the DOS feature if set to true",
						MarkdownDescription: "Enable enables the DOS feature if set to true",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "Name is the name of protected object, max of 63 characters.",
						MarkdownDescription: "Name is the name of protected object, max of 63 characters.",
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

func (r *AppprotectdosF5ComDosProtectedResourceV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_appprotectdos_f5_com_dos_protected_resource_v1beta1_manifest")

	var model AppprotectdosF5ComDosProtectedResourceV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("appprotectdos.f5.com/v1beta1")
	model.Kind = pointer.String("DosProtectedResource")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
