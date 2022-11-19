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

type SecurityIstioIoAuthorizationPolicyV1Beta1Resource struct{}

var (
	_ resource.Resource = (*SecurityIstioIoAuthorizationPolicyV1Beta1Resource)(nil)
)

type SecurityIstioIoAuthorizationPolicyV1Beta1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type SecurityIstioIoAuthorizationPolicyV1Beta1GoModel struct {
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
		Action *string `tfsdk:"action" yaml:"action,omitempty"`

		Provider *struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"provider" yaml:"provider,omitempty"`

		Rules *[]struct {
			From *[]struct {
				Source *struct {
					IpBlocks *[]string `tfsdk:"ip_blocks" yaml:"ipBlocks,omitempty"`

					Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

					NotIpBlocks *[]string `tfsdk:"not_ip_blocks" yaml:"notIpBlocks,omitempty"`

					NotNamespaces *[]string `tfsdk:"not_namespaces" yaml:"notNamespaces,omitempty"`

					NotPrincipals *[]string `tfsdk:"not_principals" yaml:"notPrincipals,omitempty"`

					NotRemoteIpBlocks *[]string `tfsdk:"not_remote_ip_blocks" yaml:"notRemoteIpBlocks,omitempty"`

					NotRequestPrincipals *[]string `tfsdk:"not_request_principals" yaml:"notRequestPrincipals,omitempty"`

					Principals *[]string `tfsdk:"principals" yaml:"principals,omitempty"`

					RemoteIpBlocks *[]string `tfsdk:"remote_ip_blocks" yaml:"remoteIpBlocks,omitempty"`

					RequestPrincipals *[]string `tfsdk:"request_principals" yaml:"requestPrincipals,omitempty"`
				} `tfsdk:"source" yaml:"source,omitempty"`
			} `tfsdk:"from" yaml:"from,omitempty"`

			To *[]struct {
				Operation *struct {
					Hosts *[]string `tfsdk:"hosts" yaml:"hosts,omitempty"`

					Methods *[]string `tfsdk:"methods" yaml:"methods,omitempty"`

					NotHosts *[]string `tfsdk:"not_hosts" yaml:"notHosts,omitempty"`

					NotMethods *[]string `tfsdk:"not_methods" yaml:"notMethods,omitempty"`

					NotPaths *[]string `tfsdk:"not_paths" yaml:"notPaths,omitempty"`

					NotPorts *[]string `tfsdk:"not_ports" yaml:"notPorts,omitempty"`

					Paths *[]string `tfsdk:"paths" yaml:"paths,omitempty"`

					Ports *[]string `tfsdk:"ports" yaml:"ports,omitempty"`
				} `tfsdk:"operation" yaml:"operation,omitempty"`
			} `tfsdk:"to" yaml:"to,omitempty"`

			When *[]struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				NotValues *[]string `tfsdk:"not_values" yaml:"notValues,omitempty"`

				Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
			} `tfsdk:"when" yaml:"when,omitempty"`
		} `tfsdk:"rules" yaml:"rules,omitempty"`

		Selector *struct {
			MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
		} `tfsdk:"selector" yaml:"selector,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewSecurityIstioIoAuthorizationPolicyV1Beta1Resource() resource.Resource {
	return &SecurityIstioIoAuthorizationPolicyV1Beta1Resource{}
}

func (r *SecurityIstioIoAuthorizationPolicyV1Beta1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_istio_io_authorization_policy_v1beta1"
}

func (r *SecurityIstioIoAuthorizationPolicyV1Beta1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "Configuration for access control on workloads. See more details at: https://istio.io/docs/reference/config/security/authorization-policy.html",
				MarkdownDescription: "Configuration for access control on workloads. See more details at: https://istio.io/docs/reference/config/security/authorization-policy.html",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"action": {
						Description:         "Optional.",
						MarkdownDescription: "Optional.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("ALLOW", "DENY", "AUDIT", "CUSTOM"),
						},
					},

					"provider": {
						Description:         "Specifies detailed configuration of the CUSTOM action.",
						MarkdownDescription: "Specifies detailed configuration of the CUSTOM action.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Specifies the name of the extension provider.",
								MarkdownDescription: "Specifies the name of the extension provider.",

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

					"rules": {
						Description:         "Optional.",
						MarkdownDescription: "Optional.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"from": {
								Description:         "Optional.",
								MarkdownDescription: "Optional.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"source": {
										Description:         "Source specifies the source of a request.",
										MarkdownDescription: "Source specifies the source of a request.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"ip_blocks": {
												Description:         "Optional.",
												MarkdownDescription: "Optional.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespaces": {
												Description:         "Optional.",
												MarkdownDescription: "Optional.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"not_ip_blocks": {
												Description:         "Optional.",
												MarkdownDescription: "Optional.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"not_namespaces": {
												Description:         "Optional.",
												MarkdownDescription: "Optional.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"not_principals": {
												Description:         "Optional.",
												MarkdownDescription: "Optional.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"not_remote_ip_blocks": {
												Description:         "Optional.",
												MarkdownDescription: "Optional.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"not_request_principals": {
												Description:         "Optional.",
												MarkdownDescription: "Optional.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"principals": {
												Description:         "Optional.",
												MarkdownDescription: "Optional.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"remote_ip_blocks": {
												Description:         "Optional.",
												MarkdownDescription: "Optional.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"request_principals": {
												Description:         "Optional.",
												MarkdownDescription: "Optional.",

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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to": {
								Description:         "Optional.",
								MarkdownDescription: "Optional.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"operation": {
										Description:         "Operation specifies the operation of a request.",
										MarkdownDescription: "Operation specifies the operation of a request.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"hosts": {
												Description:         "Optional.",
												MarkdownDescription: "Optional.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"methods": {
												Description:         "Optional.",
												MarkdownDescription: "Optional.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"not_hosts": {
												Description:         "Optional.",
												MarkdownDescription: "Optional.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"not_methods": {
												Description:         "Optional.",
												MarkdownDescription: "Optional.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"not_paths": {
												Description:         "Optional.",
												MarkdownDescription: "Optional.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"not_ports": {
												Description:         "Optional.",
												MarkdownDescription: "Optional.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"paths": {
												Description:         "Optional.",
												MarkdownDescription: "Optional.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ports": {
												Description:         "Optional.",
												MarkdownDescription: "Optional.",

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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"when": {
								Description:         "Optional.",
								MarkdownDescription: "Optional.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"key": {
										Description:         "The name of an Istio attribute.",
										MarkdownDescription: "The name of an Istio attribute.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"not_values": {
										Description:         "Optional.",
										MarkdownDescription: "Optional.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"values": {
										Description:         "Optional.",
										MarkdownDescription: "Optional.",

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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"selector": {
						Description:         "Optional.",
						MarkdownDescription: "Optional.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"match_labels": {
								Description:         "",
								MarkdownDescription: "",

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

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *SecurityIstioIoAuthorizationPolicyV1Beta1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_security_istio_io_authorization_policy_v1beta1")

	var state SecurityIstioIoAuthorizationPolicyV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SecurityIstioIoAuthorizationPolicyV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("security.istio.io/v1beta1")
	goModel.Kind = utilities.Ptr("AuthorizationPolicy")

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

func (r *SecurityIstioIoAuthorizationPolicyV1Beta1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_security_istio_io_authorization_policy_v1beta1")
	// NO-OP: All data is already in Terraform state
}

func (r *SecurityIstioIoAuthorizationPolicyV1Beta1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_security_istio_io_authorization_policy_v1beta1")

	var state SecurityIstioIoAuthorizationPolicyV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SecurityIstioIoAuthorizationPolicyV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("security.istio.io/v1beta1")
	goModel.Kind = utilities.Ptr("AuthorizationPolicy")

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

func (r *SecurityIstioIoAuthorizationPolicyV1Beta1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_security_istio_io_authorization_policy_v1beta1")
	// NO-OP: Terraform removes the state automatically for us
}
