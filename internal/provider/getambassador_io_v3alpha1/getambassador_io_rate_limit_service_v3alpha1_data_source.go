/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package getambassador_io_v3alpha1

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
	_ datasource.DataSource              = &GetambassadorIoRateLimitServiceV3Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &GetambassadorIoRateLimitServiceV3Alpha1DataSource{}
)

func NewGetambassadorIoRateLimitServiceV3Alpha1DataSource() datasource.DataSource {
	return &GetambassadorIoRateLimitServiceV3Alpha1DataSource{}
}

type GetambassadorIoRateLimitServiceV3Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type GetambassadorIoRateLimitServiceV3Alpha1DataSourceData struct {
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
		Ambassador_id     *[]string `tfsdk:"ambassador_id" json:"ambassador_id,omitempty"`
		Domain            *string   `tfsdk:"domain" json:"domain,omitempty"`
		Failure_mode_deny *bool     `tfsdk:"failure_mode_deny" json:"failure_mode_deny,omitempty"`
		Protocol_version  *string   `tfsdk:"protocol_version" json:"protocol_version,omitempty"`
		Service           *string   `tfsdk:"service" json:"service,omitempty"`
		Stats_name        *string   `tfsdk:"stats_name" json:"stats_name,omitempty"`
		Timeout_ms        *int64    `tfsdk:"timeout_ms" json:"timeout_ms,omitempty"`
		Tls               *string   `tfsdk:"tls" json:"tls,omitempty"`
		V2ExplicitTLS     *struct {
			ServiceScheme *string `tfsdk:"service_scheme" json:"serviceScheme,omitempty"`
			Tls           *string `tfsdk:"tls" json:"tls,omitempty"`
		} `tfsdk:"v2_explicit_tls" json:"v2ExplicitTLS,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GetambassadorIoRateLimitServiceV3Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_getambassador_io_rate_limit_service_v3alpha1"
}

func (r *GetambassadorIoRateLimitServiceV3Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "RateLimitService is the Schema for the ratelimitservices API",
		MarkdownDescription: "RateLimitService is the Schema for the ratelimitservices API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
				Description:         "RateLimitServiceSpec defines the desired state of RateLimitService",
				MarkdownDescription: "RateLimitServiceSpec defines the desired state of RateLimitService",
				Attributes: map[string]schema.Attribute{
					"ambassador_id": schema.ListAttribute{
						Description:         "Common to all Ambassador objects.",
						MarkdownDescription: "Common to all Ambassador objects.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"domain": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"failure_mode_deny": schema.BoolAttribute{
						Description:         "FailureModeDeny when set to true, envoy will deny traffic if it is unable to communicate with the rate limit service.",
						MarkdownDescription: "FailureModeDeny when set to true, envoy will deny traffic if it is unable to communicate with the rate limit service.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"protocol_version": schema.StringAttribute{
						Description:         "ProtocolVersion is the envoy api transport protocol version",
						MarkdownDescription: "ProtocolVersion is the envoy api transport protocol version",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"service": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"stats_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"timeout_ms": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"tls": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"v2_explicit_tls": schema.SingleNestedAttribute{
						Description:         "V2ExplicitTLS controls some vanity/stylistic elements when converting from v3alpha1 to v2.  The values in an V2ExplicitTLS should not in any way affect the runtime operation of Emissary; except that it may affect internal names in the Envoy config, which may in turn affect stats names.  But it should not affect any end-user observable behavior.",
						MarkdownDescription: "V2ExplicitTLS controls some vanity/stylistic elements when converting from v3alpha1 to v2.  The values in an V2ExplicitTLS should not in any way affect the runtime operation of Emissary; except that it may affect internal names in the Envoy config, which may in turn affect stats names.  But it should not affect any end-user observable behavior.",
						Attributes: map[string]schema.Attribute{
							"service_scheme": schema.StringAttribute{
								Description:         "ServiceScheme specifies how to spell and capitalize the scheme-part of the service URL.  Acceptable values are 'http://' (case-insensitive), 'https://' (case-insensitive), or ''.  The value is used if it agrees with whether or not this resource enables TLS origination, or if something else in the resource overrides the scheme.",
								MarkdownDescription: "ServiceScheme specifies how to spell and capitalize the scheme-part of the service URL.  Acceptable values are 'http://' (case-insensitive), 'https://' (case-insensitive), or ''.  The value is used if it agrees with whether or not this resource enables TLS origination, or if something else in the resource overrides the scheme.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"tls": schema.StringAttribute{
								Description:         "TLS controls whether and how to represent the 'tls' field when its value could be implied by the 'service' field.  In v2, there were a lot of different ways to spell an 'empty' value, and this field specifies which way to spell it (and will therefore only be used if the value will indeed be empty).  | Value        | Representation                        | Meaning of representation          | |--------------+---------------------------------------+------------------------------------| | ''           | omit the field                        | defer to service (no TLSContext)   | | 'null'       | store an explicit 'null' in the field | defer to service (no TLSContext)   | | 'string'     | store an empty string in the field    | defer to service (no TLSContext)   | | 'bool:false' | store a Boolean 'false' in the field  | defer to service (no TLSContext)   | | 'bool:true'  | store a Boolean 'true' in the field   | originate TLS (no TLSContext)      |  If the meaning of the representation contradicts anything else (if a TLSContext is to be used, or in the case of 'bool:true' if TLS is not to be originated), then this field is ignored.",
								MarkdownDescription: "TLS controls whether and how to represent the 'tls' field when its value could be implied by the 'service' field.  In v2, there were a lot of different ways to spell an 'empty' value, and this field specifies which way to spell it (and will therefore only be used if the value will indeed be empty).  | Value        | Representation                        | Meaning of representation          | |--------------+---------------------------------------+------------------------------------| | ''           | omit the field                        | defer to service (no TLSContext)   | | 'null'       | store an explicit 'null' in the field | defer to service (no TLSContext)   | | 'string'     | store an empty string in the field    | defer to service (no TLSContext)   | | 'bool:false' | store a Boolean 'false' in the field  | defer to service (no TLSContext)   | | 'bool:true'  | store a Boolean 'true' in the field   | originate TLS (no TLSContext)      |  If the meaning of the representation contradicts anything else (if a TLSContext is to be used, or in the case of 'bool:true' if TLS is not to be originated), then this field is ignored.",
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
		},
	}
}

func (r *GetambassadorIoRateLimitServiceV3Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *GetambassadorIoRateLimitServiceV3Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_getambassador_io_rate_limit_service_v3alpha1")

	var data GetambassadorIoRateLimitServiceV3Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "getambassador.io", Version: "v3alpha1", Resource: "RateLimitService"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse GetambassadorIoRateLimitServiceV3Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("getambassador.io/v3alpha1")
	data.Kind = pointer.String("RateLimitService")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}