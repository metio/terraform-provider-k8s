/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package operator_marin3r_3scale_net_v1alpha1

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
	_ datasource.DataSource = &OperatorMarin3R3ScaleNetDiscoveryServiceV1Alpha1Manifest{}
)

func NewOperatorMarin3R3ScaleNetDiscoveryServiceV1Alpha1Manifest() datasource.DataSource {
	return &OperatorMarin3R3ScaleNetDiscoveryServiceV1Alpha1Manifest{}
}

type OperatorMarin3R3ScaleNetDiscoveryServiceV1Alpha1Manifest struct{}

type OperatorMarin3R3ScaleNetDiscoveryServiceV1Alpha1ManifestData struct {
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
		Debug       *bool   `tfsdk:"debug" json:"debug,omitempty"`
		Image       *string `tfsdk:"image" json:"image,omitempty"`
		MetricsPort *int64  `tfsdk:"metrics_port" json:"metricsPort,omitempty"`
		PkiConfg    *struct {
			RootCertificateAuthority *struct {
				Duration   *string `tfsdk:"duration" json:"duration,omitempty"`
				SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
			} `tfsdk:"root_certificate_authority" json:"rootCertificateAuthority,omitempty"`
			ServerCertificate *struct {
				Duration   *string `tfsdk:"duration" json:"duration,omitempty"`
				SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
			} `tfsdk:"server_certificate" json:"serverCertificate,omitempty"`
		} `tfsdk:"pki_confg" json:"pkiConfg,omitempty"`
		PodPriorityClass *string `tfsdk:"pod_priority_class" json:"podPriorityClass,omitempty"`
		ProbePort        *int64  `tfsdk:"probe_port" json:"probePort,omitempty"`
		Resources        *struct {
			Claims *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"claims" json:"claims,omitempty"`
			Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
			Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"resources" json:"resources,omitempty"`
		ServiceConfig *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"service_config" json:"serviceConfig,omitempty"`
		XdsServerPort *int64 `tfsdk:"xds_server_port" json:"xdsServerPort,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OperatorMarin3R3ScaleNetDiscoveryServiceV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operator_marin3r_3scale_net_discovery_service_v1alpha1_manifest"
}

func (r *OperatorMarin3R3ScaleNetDiscoveryServiceV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DiscoveryService represents an envoy discovery service server. Only one instance per namespace is currently supported.",
		MarkdownDescription: "DiscoveryService represents an envoy discovery service server. Only one instance per namespace is currently supported.",
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
				Description:         "DiscoveryServiceSpec defines the desired state of DiscoveryService",
				MarkdownDescription: "DiscoveryServiceSpec defines the desired state of DiscoveryService",
				Attributes: map[string]schema.Attribute{
					"debug": schema.BoolAttribute{
						Description:         "Debug enables debugging log level for the discovery service controllers. It is safe to use since secret data is never shown in the logs.",
						MarkdownDescription: "Debug enables debugging log level for the discovery service controllers. It is safe to use since secret data is never shown in the logs.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image": schema.StringAttribute{
						Description:         "Image holds the image to use for the discovery service Deployment",
						MarkdownDescription: "Image holds the image to use for the discovery service Deployment",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"metrics_port": schema.Int64Attribute{
						Description:         "MetricsPort is the port where metrics are served. Defaults to 8383.",
						MarkdownDescription: "MetricsPort is the port where metrics are served. Defaults to 8383.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"pki_confg": schema.SingleNestedAttribute{
						Description:         "PKIConfig has configuration for the PKI that marin3r manages for the different certificates it requires",
						MarkdownDescription: "PKIConfig has configuration for the PKI that marin3r manages for the different certificates it requires",
						Attributes: map[string]schema.Attribute{
							"root_certificate_authority": schema.SingleNestedAttribute{
								Description:         "CertificateOptions specifies options to generate the server certificate used both for the xDS server and the mutating webhook server.",
								MarkdownDescription: "CertificateOptions specifies options to generate the server certificate used both for the xDS server and the mutating webhook server.",
								Attributes: map[string]schema.Attribute{
									"duration": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"secret_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
							},

							"server_certificate": schema.SingleNestedAttribute{
								Description:         "CertificateOptions specifies options to generate the server certificate used both for the xDS server and the mutating webhook server.",
								MarkdownDescription: "CertificateOptions specifies options to generate the server certificate used both for the xDS server and the mutating webhook server.",
								Attributes: map[string]schema.Attribute{
									"duration": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"secret_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"pod_priority_class": schema.StringAttribute{
						Description:         "PriorityClass to assign the discovery service Pod to",
						MarkdownDescription: "PriorityClass to assign the discovery service Pod to",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"probe_port": schema.Int64Attribute{
						Description:         "ProbePort is the port where healthz endpoint is served. Defaults to 8384.",
						MarkdownDescription: "ProbePort is the port where healthz endpoint is served. Defaults to 8384.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"resources": schema.SingleNestedAttribute{
						Description:         "Resources holds the Resource Requirements to use for the discovery service Deployment. When not set it defaults to no resource requests nor limits. CPU and Memory resources are supported.",
						MarkdownDescription: "Resources holds the Resource Requirements to use for the discovery service Deployment. When not set it defaults to no resource requests nor limits. CPU and Memory resources are supported.",
						Attributes: map[string]schema.Attribute{
							"claims": schema.ListNestedAttribute{
								Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
								MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
											MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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

							"limits": schema.MapAttribute{
								Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"requests": schema.MapAttribute{
								Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

					"service_config": schema.SingleNestedAttribute{
						Description:         "ServiceConfig configures the way the DiscoveryService endpoints are exposed",
						MarkdownDescription: "ServiceConfig configures the way the DiscoveryService endpoints are exposed",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "ServiceType is an enum with the available discovery service Service types",
								MarkdownDescription: "ServiceType is an enum with the available discovery service Service types",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"xds_server_port": schema.Int64Attribute{
						Description:         "XdsServerPort is the port where the xDS server listens. Defaults to 18000.",
						MarkdownDescription: "XdsServerPort is the port where the xDS server listens. Defaults to 18000.",
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

func (r *OperatorMarin3R3ScaleNetDiscoveryServiceV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_operator_marin3r_3scale_net_discovery_service_v1alpha1_manifest")

	var model OperatorMarin3R3ScaleNetDiscoveryServiceV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("operator.marin3r.3scale.net/v1alpha1")
	model.Kind = pointer.String("DiscoveryService")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
