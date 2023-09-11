/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package multicluster_x_k8s_io_v1alpha1

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
	_ datasource.DataSource              = &MulticlusterXK8SIoServiceImportV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &MulticlusterXK8SIoServiceImportV1Alpha1DataSource{}
)

func NewMulticlusterXK8SIoServiceImportV1Alpha1DataSource() datasource.DataSource {
	return &MulticlusterXK8SIoServiceImportV1Alpha1DataSource{}
}

type MulticlusterXK8SIoServiceImportV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type MulticlusterXK8SIoServiceImportV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

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

func (r *MulticlusterXK8SIoServiceImportV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_multicluster_x_k8s_io_service_import_v1alpha1"
}

func (r *MulticlusterXK8SIoServiceImportV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ServiceImport describes a service imported from clusters in a ClusterSet.",
		MarkdownDescription: "ServiceImport describes a service imported from clusters in a ClusterSet.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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

			"spec": schema.SingleNestedAttribute{
				Description:         "spec defines the behavior of a ServiceImport.",
				MarkdownDescription: "spec defines the behavior of a ServiceImport.",
				Attributes: map[string]schema.Attribute{
					"ips": schema.ListAttribute{
						Description:         "ip will be used as the VIP for this service when type is ClusterSetIP.",
						MarkdownDescription: "ip will be used as the VIP for this service when type is ClusterSetIP.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"ports": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"app_protocol": schema.StringAttribute{
									Description:         "The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and http://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names such as mycompany.com/my-custom-protocol. Field can be enabled with ServiceAppProtocol feature gate.",
									MarkdownDescription: "The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and http://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names such as mycompany.com/my-custom-protocol. Field can be enabled with ServiceAppProtocol feature gate.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"name": schema.StringAttribute{
									Description:         "The name of this port within the service. This must be a DNS_LABEL. All ports within a ServiceSpec must have unique names. When considering the endpoints for a Service, this must match the 'name' field in the EndpointPort. Optional if only one ServicePort is defined on this service.",
									MarkdownDescription: "The name of this port within the service. This must be a DNS_LABEL. All ports within a ServiceSpec must have unique names. When considering the endpoints for a Service, this must match the 'name' field in the EndpointPort. Optional if only one ServicePort is defined on this service.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"port": schema.Int64Attribute{
									Description:         "The port that will be exposed by this service.",
									MarkdownDescription: "The port that will be exposed by this service.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"protocol": schema.StringAttribute{
									Description:         "The IP protocol for this port. Supports 'TCP', 'UDP', and 'SCTP'. Default is TCP.",
									MarkdownDescription: "The IP protocol for this port. Supports 'TCP', 'UDP', and 'SCTP'. Default is TCP.",
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

					"session_affinity": schema.StringAttribute{
						Description:         "Supports 'ClientIP' and 'None'. Used to maintain session affinity. Enable client IP based session affinity. Must be ClientIP or None. Defaults to None. Ignored when type is Headless More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
						MarkdownDescription: "Supports 'ClientIP' and 'None'. Used to maintain session affinity. Enable client IP based session affinity. Must be ClientIP or None. Defaults to None. Ignored when type is Headless More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
						Required:            false,
						Optional:            false,
						Computed:            true,
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
										Description:         "timeoutSeconds specifies the seconds of ClientIP type session sticky time. The value must be >0 && <=86400(for 1 day) if ServiceAffinity == 'ClientIP'. Default value is 10800(for 3 hours).",
										MarkdownDescription: "timeoutSeconds specifies the seconds of ClientIP type session sticky time. The value must be >0 && <=86400(for 1 day) if ServiceAffinity == 'ClientIP'. Default value is 10800(for 3 hours).",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"type": schema.StringAttribute{
						Description:         "type defines the type of this service. Must be ClusterSetIP or Headless.",
						MarkdownDescription: "type defines the type of this service. Must be ClusterSetIP or Headless.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *MulticlusterXK8SIoServiceImportV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *MulticlusterXK8SIoServiceImportV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_multicluster_x_k8s_io_service_import_v1alpha1")

	var data MulticlusterXK8SIoServiceImportV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "multicluster.x-k8s.io", Version: "v1alpha1", Resource: "serviceimports"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse MulticlusterXK8SIoServiceImportV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("multicluster.x-k8s.io/v1alpha1")
	data.Kind = pointer.String("ServiceImport")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
