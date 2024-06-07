/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package multicluster_x_k8s_io_v1alpha1

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
	_ datasource.DataSource = &MulticlusterXK8SIoServiceImportV1Alpha1Manifest{}
)

func NewMulticlusterXK8SIoServiceImportV1Alpha1Manifest() datasource.DataSource {
	return &MulticlusterXK8SIoServiceImportV1Alpha1Manifest{}
}

type MulticlusterXK8SIoServiceImportV1Alpha1Manifest struct{}

type MulticlusterXK8SIoServiceImportV1Alpha1ManifestData struct {
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
		Ips   *[]string `tfsdk:"ips" json:"ips,omitempty"`
		Ports *[]struct {
			AppProtocol *string `tfsdk:"app_protocol" json:"appProtocol,omitempty"`
			Name        *string `tfsdk:"name" json:"name,omitempty"`
			Port        *int64  `tfsdk:"port" json:"port,omitempty"`
			Protocol    *string `tfsdk:"protocol" json:"protocol,omitempty"`
		} `tfsdk:"ports" json:"ports,omitempty"`
		SessionAffinity       *string `tfsdk:"session_affinity" json:"sessionAffinity,omitempty"`
		SessionAffinityConfig *struct {
			ClientIP *struct {
				TimeoutSeconds *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
			} `tfsdk:"client_ip" json:"clientIP,omitempty"`
		} `tfsdk:"session_affinity_config" json:"sessionAffinityConfig,omitempty"`
		Type *string `tfsdk:"type" json:"type,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *MulticlusterXK8SIoServiceImportV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_multicluster_x_k8s_io_service_import_v1alpha1_manifest"
}

func (r *MulticlusterXK8SIoServiceImportV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ServiceImport describes a service imported from clusters in a ClusterSet.",
		MarkdownDescription: "ServiceImport describes a service imported from clusters in a ClusterSet.",
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
				Description:         "spec defines the behavior of a ServiceImport.",
				MarkdownDescription: "spec defines the behavior of a ServiceImport.",
				Attributes: map[string]schema.Attribute{
					"ips": schema.ListAttribute{
						Description:         "ip will be used as the VIP for this service when type is ClusterSetIP.",
						MarkdownDescription: "ip will be used as the VIP for this service when type is ClusterSetIP.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ports": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"app_protocol": schema.StringAttribute{
									Description:         "The application protocol for this port.This is used as a hint for implementations to offer richer behavior for protocols that they understand.This field follows standard Kubernetes label syntax.Valid values are either:* Un-prefixed protocol names - reserved for IANA standard service names (as perRFC-6335 and https://www.iana.org/assignments/service-names).* Kubernetes-defined prefixed names:  * 'kubernetes.io/h2c' - HTTP/2 over cleartext as described in https://www.rfc-editor.org/rfc/rfc7540* Other protocols should use implementation-defined prefixed names such asmycompany.com/my-custom-protocol.Field can be enabled with ServiceAppProtocol feature gate.",
									MarkdownDescription: "The application protocol for this port.This is used as a hint for implementations to offer richer behavior for protocols that they understand.This field follows standard Kubernetes label syntax.Valid values are either:* Un-prefixed protocol names - reserved for IANA standard service names (as perRFC-6335 and https://www.iana.org/assignments/service-names).* Kubernetes-defined prefixed names:  * 'kubernetes.io/h2c' - HTTP/2 over cleartext as described in https://www.rfc-editor.org/rfc/rfc7540* Other protocols should use implementation-defined prefixed names such asmycompany.com/my-custom-protocol.Field can be enabled with ServiceAppProtocol feature gate.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "The name of this port within the service. This must be a DNS_LABEL.All ports within a ServiceSpec must have unique names. When consideringthe endpoints for a Service, this must match the 'name' field in theEndpointPort.Optional if only one ServicePort is defined on this service.",
									MarkdownDescription: "The name of this port within the service. This must be a DNS_LABEL.All ports within a ServiceSpec must have unique names. When consideringthe endpoints for a Service, this must match the 'name' field in theEndpointPort.Optional if only one ServicePort is defined on this service.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"port": schema.Int64Attribute{
									Description:         "The port that will be exposed by this service.",
									MarkdownDescription: "The port that will be exposed by this service.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"protocol": schema.StringAttribute{
									Description:         "The IP protocol for this port. Supports 'TCP', 'UDP', and 'SCTP'.Default is TCP.",
									MarkdownDescription: "The IP protocol for this port. Supports 'TCP', 'UDP', and 'SCTP'.Default is TCP.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"session_affinity": schema.StringAttribute{
						Description:         "Supports 'ClientIP' and 'None'. Used to maintain session affinity.Enable client IP based session affinity.Must be ClientIP or None.Defaults to None.Ignored when type is HeadlessMore info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
						MarkdownDescription: "Supports 'ClientIP' and 'None'. Used to maintain session affinity.Enable client IP based session affinity.Must be ClientIP or None.Defaults to None.Ignored when type is HeadlessMore info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"session_affinity_config": schema.SingleNestedAttribute{
						Description:         "sessionAffinityConfig contains session affinity configuration.",
						MarkdownDescription: "sessionAffinityConfig contains session affinity configuration.",
						Attributes: map[string]schema.Attribute{
							"client_ip": schema.SingleNestedAttribute{
								Description:         "clientIP contains the configurations of Client IP based session affinity.",
								MarkdownDescription: "clientIP contains the configurations of Client IP based session affinity.",
								Attributes: map[string]schema.Attribute{
									"timeout_seconds": schema.Int64Attribute{
										Description:         "timeoutSeconds specifies the seconds of ClientIP type session sticky time.The value must be >0 && <=86400(for 1 day) if ServiceAffinity == 'ClientIP'.Default value is 10800(for 3 hours).",
										MarkdownDescription: "timeoutSeconds specifies the seconds of ClientIP type session sticky time.The value must be >0 && <=86400(for 1 day) if ServiceAffinity == 'ClientIP'.Default value is 10800(for 3 hours).",
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

					"type": schema.StringAttribute{
						Description:         "type defines the type of this service.Must be ClusterSetIP or Headless.",
						MarkdownDescription: "type defines the type of this service.Must be ClusterSetIP or Headless.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("ClusterSetIP", "Headless"),
						},
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *MulticlusterXK8SIoServiceImportV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_multicluster_x_k8s_io_service_import_v1alpha1_manifest")

	var model MulticlusterXK8SIoServiceImportV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("multicluster.x-k8s.io/v1alpha1")
	model.Kind = pointer.String("ServiceImport")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
