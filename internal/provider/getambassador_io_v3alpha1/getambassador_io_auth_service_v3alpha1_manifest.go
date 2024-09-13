/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package getambassador_io_v3alpha1

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
	_ datasource.DataSource = &GetambassadorIoAuthServiceV3Alpha1Manifest{}
)

func NewGetambassadorIoAuthServiceV3Alpha1Manifest() datasource.DataSource {
	return &GetambassadorIoAuthServiceV3Alpha1Manifest{}
}

type GetambassadorIoAuthServiceV3Alpha1Manifest struct{}

type GetambassadorIoAuthServiceV3Alpha1ManifestData struct {
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
		Add_auth_headers              *map[string]string `tfsdk:"add_auth_headers" json:"add_auth_headers,omitempty"`
		Add_linkerd_headers           *bool              `tfsdk:"add_linkerd_headers" json:"add_linkerd_headers,omitempty"`
		Allow_request_body            *bool              `tfsdk:"allow_request_body" json:"allow_request_body,omitempty"`
		Allowed_authorization_headers *[]string          `tfsdk:"allowed_authorization_headers" json:"allowed_authorization_headers,omitempty"`
		Allowed_request_headers       *[]string          `tfsdk:"allowed_request_headers" json:"allowed_request_headers,omitempty"`
		Ambassador_id                 *[]string          `tfsdk:"ambassador_id" json:"ambassador_id,omitempty"`
		Auth_service                  *string            `tfsdk:"auth_service" json:"auth_service,omitempty"`
		Circuit_breakers              *[]struct {
			Max_connections      *int64  `tfsdk:"max_connections" json:"max_connections,omitempty"`
			Max_pending_requests *int64  `tfsdk:"max_pending_requests" json:"max_pending_requests,omitempty"`
			Max_requests         *int64  `tfsdk:"max_requests" json:"max_requests,omitempty"`
			Max_retries          *int64  `tfsdk:"max_retries" json:"max_retries,omitempty"`
			Priority             *string `tfsdk:"priority" json:"priority,omitempty"`
		} `tfsdk:"circuit_breakers" json:"circuit_breakers,omitempty"`
		Failure_mode_allow *bool `tfsdk:"failure_mode_allow" json:"failure_mode_allow,omitempty"`
		Include_body       *struct {
			Allow_partial *bool  `tfsdk:"allow_partial" json:"allow_partial,omitempty"`
			Max_bytes     *int64 `tfsdk:"max_bytes" json:"max_bytes,omitempty"`
		} `tfsdk:"include_body" json:"include_body,omitempty"`
		Path_prefix      *string `tfsdk:"path_prefix" json:"path_prefix,omitempty"`
		Proto            *string `tfsdk:"proto" json:"proto,omitempty"`
		Protocol_version *string `tfsdk:"protocol_version" json:"protocol_version,omitempty"`
		Stats_name       *string `tfsdk:"stats_name" json:"stats_name,omitempty"`
		Status_on_error  *struct {
			Code *int64 `tfsdk:"code" json:"code,omitempty"`
		} `tfsdk:"status_on_error" json:"status_on_error,omitempty"`
		Timeout_ms    *int64  `tfsdk:"timeout_ms" json:"timeout_ms,omitempty"`
		Tls           *string `tfsdk:"tls" json:"tls,omitempty"`
		V2ExplicitTLS *struct {
			ServiceScheme *string `tfsdk:"service_scheme" json:"serviceScheme,omitempty"`
			Tls           *string `tfsdk:"tls" json:"tls,omitempty"`
		} `tfsdk:"v2_explicit_tls" json:"v2ExplicitTLS,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GetambassadorIoAuthServiceV3Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_getambassador_io_auth_service_v3alpha1_manifest"
}

func (r *GetambassadorIoAuthServiceV3Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "AuthService is the Schema for the authservices API",
		MarkdownDescription: "AuthService is the Schema for the authservices API",
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
				Description:         "AuthServiceSpec defines the desired state of AuthService",
				MarkdownDescription: "AuthServiceSpec defines the desired state of AuthService",
				Attributes: map[string]schema.Attribute{
					"add_auth_headers": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"add_linkerd_headers": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"allow_request_body": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"allowed_authorization_headers": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"allowed_request_headers": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ambassador_id": schema.ListAttribute{
						Description:         "AmbassadorID declares which Ambassador instances should pay attention to this resource. If no value is provided, the default is: ambassador_id: - 'default'",
						MarkdownDescription: "AmbassadorID declares which Ambassador instances should pay attention to this resource. If no value is provided, the default is: ambassador_id: - 'default'",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"auth_service": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"circuit_breakers": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"max_connections": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"max_pending_requests": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"max_requests": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"max_retries": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"priority": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("default", "high"),
									},
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"failure_mode_allow": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"include_body": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"allow_partial": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"max_bytes": schema.Int64Attribute{
								Description:         "These aren't pointer types because they are required.",
								MarkdownDescription: "These aren't pointer types because they are required.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"path_prefix": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"proto": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("http", "grpc"),
						},
					},

					"protocol_version": schema.StringAttribute{
						Description:         "ProtocolVersion is the envoy api transport protocol version",
						MarkdownDescription: "ProtocolVersion is the envoy api transport protocol version",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("v2", "v3"),
						},
					},

					"stats_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"status_on_error": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"code": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"timeout_ms": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tls": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"v2_explicit_tls": schema.SingleNestedAttribute{
						Description:         "V2ExplicitTLS controls some vanity/stylistic elements when converting from v3alpha1 to v2. The values in an V2ExplicitTLS should not in any way affect the runtime operation of Emissary; except that it may affect internal names in the Envoy config, which may in turn affect stats names. But it should not affect any end-user observable behavior.",
						MarkdownDescription: "V2ExplicitTLS controls some vanity/stylistic elements when converting from v3alpha1 to v2. The values in an V2ExplicitTLS should not in any way affect the runtime operation of Emissary; except that it may affect internal names in the Envoy config, which may in turn affect stats names. But it should not affect any end-user observable behavior.",
						Attributes: map[string]schema.Attribute{
							"service_scheme": schema.StringAttribute{
								Description:         "ServiceScheme specifies how to spell and capitalize the scheme-part of the service URL. Acceptable values are 'http://' (case-insensitive), 'https://' (case-insensitive), or ''. The value is used if it agrees with whether or not this resource enables TLS origination, or if something else in the resource overrides the scheme.",
								MarkdownDescription: "ServiceScheme specifies how to spell and capitalize the scheme-part of the service URL. Acceptable values are 'http://' (case-insensitive), 'https://' (case-insensitive), or ''. The value is used if it agrees with whether or not this resource enables TLS origination, or if something else in the resource overrides the scheme.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^([hH][tT][tT][pP][sS]?://)?$`), ""),
								},
							},

							"tls": schema.StringAttribute{
								Description:         "TLS controls whether and how to represent the 'tls' field when its value could be implied by the 'service' field. In v2, there were a lot of different ways to spell an 'empty' value, and this field specifies which way to spell it (and will therefore only be used if the value will indeed be empty). | Value | Representation | Meaning of representation | |--------------+---------------------------------------+------------------------------------| | '' | omit the field | defer to service (no TLSContext) | | 'null' | store an explicit 'null' in the field | defer to service (no TLSContext) | | 'string' | store an empty string in the field | defer to service (no TLSContext) | | 'bool:false' | store a Boolean 'false' in the field | defer to service (no TLSContext) | | 'bool:true' | store a Boolean 'true' in the field | originate TLS (no TLSContext) | If the meaning of the representation contradicts anything else (if a TLSContext is to be used, or in the case of 'bool:true' if TLS is not to be originated), then this field is ignored.",
								MarkdownDescription: "TLS controls whether and how to represent the 'tls' field when its value could be implied by the 'service' field. In v2, there were a lot of different ways to spell an 'empty' value, and this field specifies which way to spell it (and will therefore only be used if the value will indeed be empty). | Value | Representation | Meaning of representation | |--------------+---------------------------------------+------------------------------------| | '' | omit the field | defer to service (no TLSContext) | | 'null' | store an explicit 'null' in the field | defer to service (no TLSContext) | | 'string' | store an empty string in the field | defer to service (no TLSContext) | | 'bool:false' | store a Boolean 'false' in the field | defer to service (no TLSContext) | | 'bool:true' | store a Boolean 'true' in the field | originate TLS (no TLSContext) | If the meaning of the representation contradicts anything else (if a TLSContext is to be used, or in the case of 'bool:true' if TLS is not to be originated), then this field is ignored.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("", "null", "bool:true", "bool:false", "string"),
								},
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
		},
	}
}

func (r *GetambassadorIoAuthServiceV3Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_getambassador_io_auth_service_v3alpha1_manifest")

	var model GetambassadorIoAuthServiceV3Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("getambassador.io/v3alpha1")
	model.Kind = pointer.String("AuthService")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
