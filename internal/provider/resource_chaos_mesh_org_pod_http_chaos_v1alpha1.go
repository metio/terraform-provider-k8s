/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

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

type ChaosMeshOrgPodHttpChaosV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*ChaosMeshOrgPodHttpChaosV1Alpha1Resource)(nil)
)

type ChaosMeshOrgPodHttpChaosV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type ChaosMeshOrgPodHttpChaosV1Alpha1GoModel struct {
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
		Rules *[]struct {
			Actions *struct {
				Abort *bool `tfsdk:"abort" yaml:"abort,omitempty"`

				Delay *string `tfsdk:"delay" yaml:"delay,omitempty"`

				Patch *struct {
					Body *struct {
						Type *string `tfsdk:"type" yaml:"type,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"body" yaml:"body,omitempty"`

					Headers *[]string `tfsdk:"headers" yaml:"headers,omitempty"`

					Queries *[]string `tfsdk:"queries" yaml:"queries,omitempty"`
				} `tfsdk:"patch" yaml:"patch,omitempty"`

				Replace *struct {
					Body *string `tfsdk:"body" yaml:"body,omitempty"`

					Code *int64 `tfsdk:"code" yaml:"code,omitempty"`

					Headers *map[string]string `tfsdk:"headers" yaml:"headers,omitempty"`

					Method *string `tfsdk:"method" yaml:"method,omitempty"`

					Path *string `tfsdk:"path" yaml:"path,omitempty"`

					Queries *map[string]string `tfsdk:"queries" yaml:"queries,omitempty"`
				} `tfsdk:"replace" yaml:"replace,omitempty"`
			} `tfsdk:"actions" yaml:"actions,omitempty"`

			Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

			Selector *struct {
				Code *int64 `tfsdk:"code" yaml:"code,omitempty"`

				Method *string `tfsdk:"method" yaml:"method,omitempty"`

				Path *string `tfsdk:"path" yaml:"path,omitempty"`

				Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

				Request_headers *map[string]string `tfsdk:"request_headers" yaml:"request_headers,omitempty"`

				Response_headers *map[string]string `tfsdk:"response_headers" yaml:"response_headers,omitempty"`
			} `tfsdk:"selector" yaml:"selector,omitempty"`

			Source *string `tfsdk:"source" yaml:"source,omitempty"`

			Target *string `tfsdk:"target" yaml:"target,omitempty"`
		} `tfsdk:"rules" yaml:"rules,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewChaosMeshOrgPodHttpChaosV1Alpha1Resource() resource.Resource {
	return &ChaosMeshOrgPodHttpChaosV1Alpha1Resource{}
}

func (r *ChaosMeshOrgPodHttpChaosV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_chaos_mesh_org_pod_http_chaos_v1alpha1"
}

func (r *ChaosMeshOrgPodHttpChaosV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "PodHttpChaos is the Schema for the podhttpchaos API",
		MarkdownDescription: "PodHttpChaos is the Schema for the podhttpchaos API",
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
				Description:         "PodHttpChaosSpec defines the desired state of PodHttpChaos.",
				MarkdownDescription: "PodHttpChaosSpec defines the desired state of PodHttpChaos.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"rules": {
						Description:         "Rules are a list of injection rule for http request.",
						MarkdownDescription: "Rules are a list of injection rule for http request.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"actions": {
								Description:         "Actions contains rules to inject target.",
								MarkdownDescription: "Actions contains rules to inject target.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"abort": {
										Description:         "Abort is a rule to abort a http session.",
										MarkdownDescription: "Abort is a rule to abort a http session.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"delay": {
										Description:         "Delay represents the delay of the target request/response. A duration string is a possibly unsigned sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
										MarkdownDescription: "Delay represents the delay of the target request/response. A duration string is a possibly unsigned sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"patch": {
										Description:         "Patch is a rule to patch some contents in target.",
										MarkdownDescription: "Patch is a rule to patch some contents in target.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"body": {
												Description:         "Body is a rule to patch message body of target.",
												MarkdownDescription: "Body is a rule to patch message body of target.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"type": {
														Description:         "Type represents the patch type, only support 'JSON' as [merge patch json](https://tools.ietf.org/html/rfc7396) currently.",
														MarkdownDescription: "Type represents the patch type, only support 'JSON' as [merge patch json](https://tools.ietf.org/html/rfc7396) currently.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"value": {
														Description:         "Value is the patch contents.",
														MarkdownDescription: "Value is the patch contents.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"headers": {
												Description:         "Headers is a rule to append http headers of target. For example: '[['Set-Cookie', '<one cookie>'], ['Set-Cookie', '<another cookie>']]'.",
												MarkdownDescription: "Headers is a rule to append http headers of target. For example: '[['Set-Cookie', '<one cookie>'], ['Set-Cookie', '<another cookie>']]'.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"queries": {
												Description:         "Queries is a rule to append uri queries of target(Request only). For example: '[['foo', 'bar'], ['foo', 'unknown']]'.",
												MarkdownDescription: "Queries is a rule to append uri queries of target(Request only). For example: '[['foo', 'bar'], ['foo', 'unknown']]'.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"replace": {
										Description:         "Replace is a rule to replace some contents in target.",
										MarkdownDescription: "Replace is a rule to replace some contents in target.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"body": {
												Description:         "Body is a rule to replace http message body in target.",
												MarkdownDescription: "Body is a rule to replace http message body in target.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													validators.Base64Validator(),
												},
											},

											"code": {
												Description:         "Code is a rule to replace http status code in response.",
												MarkdownDescription: "Code is a rule to replace http status code in response.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"headers": {
												Description:         "Headers is a rule to replace http headers of target. The key-value pairs represent header name and header value pairs.",
												MarkdownDescription: "Headers is a rule to replace http headers of target. The key-value pairs represent header name and header value pairs.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"method": {
												Description:         "Method is a rule to replace http method in request.",
												MarkdownDescription: "Method is a rule to replace http method in request.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"path": {
												Description:         "Path is rule to to replace uri path in http request.",
												MarkdownDescription: "Path is rule to to replace uri path in http request.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"queries": {
												Description:         "Queries is a rule to replace uri queries in http request. For example, with value '{ 'foo': 'unknown' }', the '/?foo=bar' will be altered to '/?foo=unknown',",
												MarkdownDescription: "Queries is a rule to replace uri queries in http request. For example, with value '{ 'foo': 'unknown' }', the '/?foo=bar' will be altered to '/?foo=unknown',",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: true,
								Optional: false,
								Computed: false,
							},

							"port": {
								Description:         "Port represents the target port to be proxy of.",
								MarkdownDescription: "Port represents the target port to be proxy of.",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"selector": {
								Description:         "Selector contains the rules to select target.",
								MarkdownDescription: "Selector contains the rules to select target.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"code": {
										Description:         "Code is a rule to select target by http status code in response.",
										MarkdownDescription: "Code is a rule to select target by http status code in response.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"method": {
										Description:         "Method is a rule to select target by http method in request.",
										MarkdownDescription: "Method is a rule to select target by http method in request.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"path": {
										Description:         "Path is a rule to select target by uri path in http request.",
										MarkdownDescription: "Path is a rule to select target by uri path in http request.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"port": {
										Description:         "Port is a rule to select server listening on specific port.",
										MarkdownDescription: "Port is a rule to select server listening on specific port.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"request_headers": {
										Description:         "RequestHeaders is a rule to select target by http headers in request. The key-value pairs represent header name and header value pairs.",
										MarkdownDescription: "RequestHeaders is a rule to select target by http headers in request. The key-value pairs represent header name and header value pairs.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"response_headers": {
										Description:         "ResponseHeaders is a rule to select target by http headers in response. The key-value pairs represent header name and header value pairs.",
										MarkdownDescription: "ResponseHeaders is a rule to select target by http headers in response. The key-value pairs represent header name and header value pairs.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: true,
								Optional: false,
								Computed: false,
							},

							"source": {
								Description:         "Source represents the source of current rules",
								MarkdownDescription: "Source represents the source of current rules",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"target": {
								Description:         "Target is the object to be selected and injected, <Request|Response>.",
								MarkdownDescription: "Target is the object to be selected and injected, <Request|Response>.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

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

func (r *ChaosMeshOrgPodHttpChaosV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_chaos_mesh_org_pod_http_chaos_v1alpha1")

	var state ChaosMeshOrgPodHttpChaosV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ChaosMeshOrgPodHttpChaosV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("chaos-mesh.org/v1alpha1")
	goModel.Kind = utilities.Ptr("PodHttpChaos")

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

func (r *ChaosMeshOrgPodHttpChaosV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_chaos_mesh_org_pod_http_chaos_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *ChaosMeshOrgPodHttpChaosV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_chaos_mesh_org_pod_http_chaos_v1alpha1")

	var state ChaosMeshOrgPodHttpChaosV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ChaosMeshOrgPodHttpChaosV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("chaos-mesh.org/v1alpha1")
	goModel.Kind = utilities.Ptr("PodHttpChaos")

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

func (r *ChaosMeshOrgPodHttpChaosV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_chaos_mesh_org_pod_http_chaos_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
