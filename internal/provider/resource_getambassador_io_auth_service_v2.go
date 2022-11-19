/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type GetambassadorIoAuthServiceV2Resource struct{}

var (
	_ resource.Resource = (*GetambassadorIoAuthServiceV2Resource)(nil)
)

type GetambassadorIoAuthServiceV2TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type GetambassadorIoAuthServiceV2GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Add_auth_headers *map[string]string `tfsdk:"add_auth_headers" yaml:"add_auth_headers,omitempty"`

		Add_linkerd_headers *bool `tfsdk:"add_linkerd_headers" yaml:"add_linkerd_headers,omitempty"`

		Allow_request_body *bool `tfsdk:"allow_request_body" yaml:"allow_request_body,omitempty"`

		Allowed_authorization_headers *[]string `tfsdk:"allowed_authorization_headers" yaml:"allowed_authorization_headers,omitempty"`

		Allowed_request_headers *[]string `tfsdk:"allowed_request_headers" yaml:"allowed_request_headers,omitempty"`

		Ambassador_id *[]string `tfsdk:"ambassador_id" yaml:"ambassador_id,omitempty"`

		Auth_service *string `tfsdk:"auth_service" yaml:"auth_service,omitempty"`

		Failure_mode_allow *bool `tfsdk:"failure_mode_allow" yaml:"failure_mode_allow,omitempty"`

		Include_body *struct {
			Allow_partial *bool `tfsdk:"allow_partial" yaml:"allow_partial,omitempty"`

			Max_bytes *int64 `tfsdk:"max_bytes" yaml:"max_bytes,omitempty"`
		} `tfsdk:"include_body" yaml:"include_body,omitempty"`

		Path_prefix *string `tfsdk:"path_prefix" yaml:"path_prefix,omitempty"`

		Proto *string `tfsdk:"proto" yaml:"proto,omitempty"`

		Protocol_version *string `tfsdk:"protocol_version" yaml:"protocol_version,omitempty"`

		Status_on_error *struct {
			Code *int64 `tfsdk:"code" yaml:"code,omitempty"`
		} `tfsdk:"status_on_error" yaml:"status_on_error,omitempty"`

		Timeout_ms *int64 `tfsdk:"timeout_ms" yaml:"timeout_ms,omitempty"`

		Tls *bool `tfsdk:"tls" yaml:"tls,omitempty"`

		V3CircuitBreakers *[]struct {
			Max_connections *int64 `tfsdk:"max_connections" yaml:"max_connections,omitempty"`

			Max_pending_requests *int64 `tfsdk:"max_pending_requests" yaml:"max_pending_requests,omitempty"`

			Max_requests *int64 `tfsdk:"max_requests" yaml:"max_requests,omitempty"`

			Max_retries *int64 `tfsdk:"max_retries" yaml:"max_retries,omitempty"`

			Priority *string `tfsdk:"priority" yaml:"priority,omitempty"`
		} `tfsdk:"v3_circuit_breakers" yaml:"v3CircuitBreakers,omitempty"`

		V3StatsName *string `tfsdk:"v3_stats_name" yaml:"v3StatsName,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewGetambassadorIoAuthServiceV2Resource() resource.Resource {
	return &GetambassadorIoAuthServiceV2Resource{}
}

func (r *GetambassadorIoAuthServiceV2Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_getambassador_io_auth_service_v2"
}

func (r *GetambassadorIoAuthServiceV2Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "AuthService is the Schema for the authservices API",
		MarkdownDescription: "AuthService is the Schema for the authservices API",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "AuthServiceSpec defines the desired state of AuthService",
				MarkdownDescription: "AuthServiceSpec defines the desired state of AuthService",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"add_auth_headers": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"add_linkerd_headers": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"allow_request_body": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"allowed_authorization_headers": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"allowed_request_headers": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ambassador_id": {
						Description:         "AmbassadorID declares which Ambassador instances should pay attention to this resource.  May either be a string or a list of strings.  If no value is provided, the default is:  	ambassador_id: 	- 'default'",
						MarkdownDescription: "AmbassadorID declares which Ambassador instances should pay attention to this resource.  May either be a string or a list of strings.  If no value is provided, the default is:  	ambassador_id: 	- 'default'",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"auth_service": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"failure_mode_allow": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"include_body": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"allow_partial": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"max_bytes": {
								Description:         "These aren't pointer types because they are required.",
								MarkdownDescription: "These aren't pointer types because they are required.",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"path_prefix": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"proto": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("http", "grpc"),
						},
					},

					"protocol_version": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("v2", "v3"),
						},
					},

					"status_on_error": {
						Description:         "Why isn't this just an int??",
						MarkdownDescription: "Why isn't this just an int??",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"code": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"timeout_ms": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"tls": {
						Description:         "BoolOrString is a type that can hold a Boolean or a string.",
						MarkdownDescription: "BoolOrString is a type that can hold a Boolean or a string.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"v3_circuit_breakers": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"max_connections": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"max_pending_requests": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"max_requests": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"max_retries": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"priority": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("default", "high"),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"v3_stats_name": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *GetambassadorIoAuthServiceV2Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_getambassador_io_auth_service_v2")

	var state GetambassadorIoAuthServiceV2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel GetambassadorIoAuthServiceV2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("getambassador.io/v2")
	goModel.Kind = utilities.Ptr("AuthService")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *GetambassadorIoAuthServiceV2Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_getambassador_io_auth_service_v2")
	// NO-OP: All data is already in Terraform state
}

func (r *GetambassadorIoAuthServiceV2Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_getambassador_io_auth_service_v2")

	var state GetambassadorIoAuthServiceV2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel GetambassadorIoAuthServiceV2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("getambassador.io/v2")
	goModel.Kind = utilities.Ptr("AuthService")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *GetambassadorIoAuthServiceV2Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_getambassador_io_auth_service_v2")
	// NO-OP: Terraform removes the state automatically for us
}
