/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package getambassador_io_v1

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
	_ datasource.DataSource              = &GetambassadorIoAuthServiceV1DataSource{}
	_ datasource.DataSourceWithConfigure = &GetambassadorIoAuthServiceV1DataSource{}
)

func NewGetambassadorIoAuthServiceV1DataSource() datasource.DataSource {
	return &GetambassadorIoAuthServiceV1DataSource{}
}

type GetambassadorIoAuthServiceV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type GetambassadorIoAuthServiceV1DataSourceData struct {
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
		Add_auth_headers              *map[string]string `tfsdk:"add_auth_headers" json:"add_auth_headers,omitempty"`
		Add_linkerd_headers           *bool              `tfsdk:"add_linkerd_headers" json:"add_linkerd_headers,omitempty"`
		Allow_request_body            *bool              `tfsdk:"allow_request_body" json:"allow_request_body,omitempty"`
		Allowed_authorization_headers *[]string          `tfsdk:"allowed_authorization_headers" json:"allowed_authorization_headers,omitempty"`
		Allowed_request_headers       *[]string          `tfsdk:"allowed_request_headers" json:"allowed_request_headers,omitempty"`
		Ambassador_id                 *[]string          `tfsdk:"ambassador_id" json:"ambassador_id,omitempty"`
		Auth_service                  *string            `tfsdk:"auth_service" json:"auth_service,omitempty"`
		Failure_mode_allow            *bool              `tfsdk:"failure_mode_allow" json:"failure_mode_allow,omitempty"`
		Include_body                  *struct {
			Allow_partial *bool  `tfsdk:"allow_partial" json:"allow_partial,omitempty"`
			Max_bytes     *int64 `tfsdk:"max_bytes" json:"max_bytes,omitempty"`
		} `tfsdk:"include_body" json:"include_body,omitempty"`
		Path_prefix      *string `tfsdk:"path_prefix" json:"path_prefix,omitempty"`
		Proto            *string `tfsdk:"proto" json:"proto,omitempty"`
		Protocol_version *string `tfsdk:"protocol_version" json:"protocol_version,omitempty"`
		Status_on_error  *struct {
			Code *int64 `tfsdk:"code" json:"code,omitempty"`
		} `tfsdk:"status_on_error" json:"status_on_error,omitempty"`
		Timeout_ms        *int64 `tfsdk:"timeout_ms" json:"timeout_ms,omitempty"`
		Tls               *bool  `tfsdk:"tls" json:"tls,omitempty"`
		V3CircuitBreakers *[]struct {
			Max_connections      *int64  `tfsdk:"max_connections" json:"max_connections,omitempty"`
			Max_pending_requests *int64  `tfsdk:"max_pending_requests" json:"max_pending_requests,omitempty"`
			Max_requests         *int64  `tfsdk:"max_requests" json:"max_requests,omitempty"`
			Max_retries          *int64  `tfsdk:"max_retries" json:"max_retries,omitempty"`
			Priority             *string `tfsdk:"priority" json:"priority,omitempty"`
		} `tfsdk:"v3_circuit_breakers" json:"v3CircuitBreakers,omitempty"`
		V3StatsName *string `tfsdk:"v3_stats_name" json:"v3StatsName,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GetambassadorIoAuthServiceV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_getambassador_io_auth_service_v1"
}

func (r *GetambassadorIoAuthServiceV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "AuthService is the Schema for the authservices API",
		MarkdownDescription: "AuthService is the Schema for the authservices API",
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
				Description:         "AuthServiceSpec defines the desired state of AuthService",
				MarkdownDescription: "AuthServiceSpec defines the desired state of AuthService",
				Attributes: map[string]schema.Attribute{
					"add_auth_headers": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"add_linkerd_headers": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"allow_request_body": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"allowed_authorization_headers": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"allowed_request_headers": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"ambassador_id": schema.ListAttribute{
						Description:         "AmbassadorID declares which Ambassador instances should pay attention to this resource.  May either be a string or a list of strings.  If no value is provided, the default is:  ambassador_id: - 'default'",
						MarkdownDescription: "AmbassadorID declares which Ambassador instances should pay attention to this resource.  May either be a string or a list of strings.  If no value is provided, the default is:  ambassador_id: - 'default'",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"auth_service": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"failure_mode_allow": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"include_body": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"allow_partial": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"max_bytes": schema.Int64Attribute{
								Description:         "These aren't pointer types because they are required.",
								MarkdownDescription: "These aren't pointer types because they are required.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"path_prefix": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"proto": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"protocol_version": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"status_on_error": schema.SingleNestedAttribute{
						Description:         "Why isn't this just an int??",
						MarkdownDescription: "Why isn't this just an int??",
						Attributes: map[string]schema.Attribute{
							"code": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"timeout_ms": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"tls": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"v3_circuit_breakers": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"max_connections": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"max_pending_requests": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"max_requests": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"max_retries": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"priority": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"v3_stats_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
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

func (r *GetambassadorIoAuthServiceV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *GetambassadorIoAuthServiceV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_getambassador_io_auth_service_v1")

	var data GetambassadorIoAuthServiceV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "getambassador.io", Version: "v1", Resource: "AuthService"}).
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

	var readResponse GetambassadorIoAuthServiceV1DataSourceData
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
	data.ApiVersion = pointer.String("getambassador.io/v1")
	data.Kind = pointer.String("AuthService")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}