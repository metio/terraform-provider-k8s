/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"regexp"

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

type GetambassadorIoTCPMappingV3Alpha1Resource struct{}

var (
	_ resource.Resource = (*GetambassadorIoTCPMappingV3Alpha1Resource)(nil)
)

type GetambassadorIoTCPMappingV3Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type GetambassadorIoTCPMappingV3Alpha1GoModel struct {
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
		Address *string `tfsdk:"address" yaml:"address,omitempty"`

		Ambassador_id *[]string `tfsdk:"ambassador_id" yaml:"ambassador_id,omitempty"`

		Circuit_breakers *[]struct {
			Max_connections *int64 `tfsdk:"max_connections" yaml:"max_connections,omitempty"`

			Max_pending_requests *int64 `tfsdk:"max_pending_requests" yaml:"max_pending_requests,omitempty"`

			Max_requests *int64 `tfsdk:"max_requests" yaml:"max_requests,omitempty"`

			Max_retries *int64 `tfsdk:"max_retries" yaml:"max_retries,omitempty"`

			Priority *string `tfsdk:"priority" yaml:"priority,omitempty"`
		} `tfsdk:"circuit_breakers" yaml:"circuit_breakers,omitempty"`

		Cluster_tag *string `tfsdk:"cluster_tag" yaml:"cluster_tag,omitempty"`

		Enable_ipv4 *bool `tfsdk:"enable_ipv4" yaml:"enable_ipv4,omitempty"`

		Enable_ipv6 *bool `tfsdk:"enable_ipv6" yaml:"enable_ipv6,omitempty"`

		Host *string `tfsdk:"host" yaml:"host,omitempty"`

		Idle_timeout_ms *string `tfsdk:"idle_timeout_ms" yaml:"idle_timeout_ms,omitempty"`

		Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

		Resolver *string `tfsdk:"resolver" yaml:"resolver,omitempty"`

		Service *string `tfsdk:"service" yaml:"service,omitempty"`

		Stats_name *string `tfsdk:"stats_name" yaml:"stats_name,omitempty"`

		Tls *string `tfsdk:"tls" yaml:"tls,omitempty"`

		V2ExplicitTLS *struct {
			ServiceScheme *string `tfsdk:"service_scheme" yaml:"serviceScheme,omitempty"`

			Tls *string `tfsdk:"tls" yaml:"tls,omitempty"`
		} `tfsdk:"v2_explicit_tls" yaml:"v2ExplicitTLS,omitempty"`

		Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewGetambassadorIoTCPMappingV3Alpha1Resource() resource.Resource {
	return &GetambassadorIoTCPMappingV3Alpha1Resource{}
}

func (r *GetambassadorIoTCPMappingV3Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_getambassador_io_tcp_mapping_v3alpha1"
}

func (r *GetambassadorIoTCPMappingV3Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "TCPMapping is the Schema for the tcpmappings API",
		MarkdownDescription: "TCPMapping is the Schema for the tcpmappings API",
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
				Description:         "TCPMappingSpec defines the desired state of TCPMapping",
				MarkdownDescription: "TCPMappingSpec defines the desired state of TCPMapping",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"address": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ambassador_id": {
						Description:         "AmbassadorID declares which Ambassador instances should pay attention to this resource. If no value is provided, the default is:  	ambassador_id: 	- 'default'  TODO(lukeshu): In v3alpha2, consider renaming all of the 'ambassador_id' (singular) fields to 'ambassador_ids' (plural).",
						MarkdownDescription: "AmbassadorID declares which Ambassador instances should pay attention to this resource. If no value is provided, the default is:  	ambassador_id: 	- 'default'  TODO(lukeshu): In v3alpha2, consider renaming all of the 'ambassador_id' (singular) fields to 'ambassador_ids' (plural).",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"circuit_breakers": {
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

					"cluster_tag": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_ipv4": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_ipv6": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"host": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"idle_timeout_ms": {
						Description:         "FIXME(lukeshu): Surely this should be an 'int'?",
						MarkdownDescription: "FIXME(lukeshu): Surely this should be an 'int'?",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"port": {
						Description:         "Port isn't a pointer because it's required.",
						MarkdownDescription: "Port isn't a pointer because it's required.",

						Type: types.Int64Type,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"resolver": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"service": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"stats_name": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"tls": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"v2_explicit_tls": {
						Description:         "V2ExplicitTLS controls some vanity/stylistic elements when converting from v3alpha1 to v2.  The values in an V2ExplicitTLS should not in any way affect the runtime operation of Emissary; except that it may affect internal names in the Envoy config, which may in turn affect stats names.  But it should not affect any end-user observable behavior.",
						MarkdownDescription: "V2ExplicitTLS controls some vanity/stylistic elements when converting from v3alpha1 to v2.  The values in an V2ExplicitTLS should not in any way affect the runtime operation of Emissary; except that it may affect internal names in the Envoy config, which may in turn affect stats names.  But it should not affect any end-user observable behavior.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"service_scheme": {
								Description:         "ServiceScheme specifies how to spell and capitalize the scheme-part of the service URL.  Acceptable values are 'http://' (case-insensitive), 'https://' (case-insensitive), or ''.  The value is used if it agrees with whether or not this resource enables TLS origination, or if something else in the resource overrides the scheme.",
								MarkdownDescription: "ServiceScheme specifies how to spell and capitalize the scheme-part of the service URL.  Acceptable values are 'http://' (case-insensitive), 'https://' (case-insensitive), or ''.  The value is used if it agrees with whether or not this resource enables TLS origination, or if something else in the resource overrides the scheme.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.RegexMatches(regexp.MustCompile(`^([hH][tT][tT][pP][sS]?://)?$`), ""),
								},
							},

							"tls": {
								Description:         "TLS controls whether and how to represent the 'tls' field when its value could be implied by the 'service' field.  In v2, there were a lot of different ways to spell an 'empty' value, and this field specifies which way to spell it (and will therefore only be used if the value will indeed be empty).   | Value        | Representation                        | Meaning of representation          |  |--------------+---------------------------------------+------------------------------------|  | ''           | omit the field                        | defer to service (no TLSContext)   |  | 'null'       | store an explicit 'null' in the field | defer to service (no TLSContext)   |  | 'string'     | store an empty string in the field    | defer to service (no TLSContext)   |  | 'bool:false' | store a Boolean 'false' in the field  | defer to service (no TLSContext)   |  | 'bool:true'  | store a Boolean 'true' in the field   | originate TLS (no TLSContext)      |  If the meaning of the representation contradicts anything else (if a TLSContext is to be used, or in the case of 'bool:true' if TLS is not to be originated), then this field is ignored.",
								MarkdownDescription: "TLS controls whether and how to represent the 'tls' field when its value could be implied by the 'service' field.  In v2, there were a lot of different ways to spell an 'empty' value, and this field specifies which way to spell it (and will therefore only be used if the value will indeed be empty).   | Value        | Representation                        | Meaning of representation          |  |--------------+---------------------------------------+------------------------------------|  | ''           | omit the field                        | defer to service (no TLSContext)   |  | 'null'       | store an explicit 'null' in the field | defer to service (no TLSContext)   |  | 'string'     | store an empty string in the field    | defer to service (no TLSContext)   |  | 'bool:false' | store a Boolean 'false' in the field  | defer to service (no TLSContext)   |  | 'bool:true'  | store a Boolean 'true' in the field   | originate TLS (no TLSContext)      |  If the meaning of the representation contradicts anything else (if a TLSContext is to be used, or in the case of 'bool:true' if TLS is not to be originated), then this field is ignored.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("", "null", "bool:true", "bool:false", "string"),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"weight": {
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
		},
	}, nil
}

func (r *GetambassadorIoTCPMappingV3Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_getambassador_io_tcp_mapping_v3alpha1")

	var state GetambassadorIoTCPMappingV3Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel GetambassadorIoTCPMappingV3Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("getambassador.io/v3alpha1")
	goModel.Kind = utilities.Ptr("TCPMapping")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *GetambassadorIoTCPMappingV3Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_getambassador_io_tcp_mapping_v3alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *GetambassadorIoTCPMappingV3Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_getambassador_io_tcp_mapping_v3alpha1")

	var state GetambassadorIoTCPMappingV3Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel GetambassadorIoTCPMappingV3Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("getambassador.io/v3alpha1")
	goModel.Kind = utilities.Ptr("TCPMapping")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *GetambassadorIoTCPMappingV3Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_getambassador_io_tcp_mapping_v3alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
