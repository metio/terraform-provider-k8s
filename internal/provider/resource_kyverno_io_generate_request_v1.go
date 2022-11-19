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

type KyvernoIoGenerateRequestV1Resource struct{}

var (
	_ resource.Resource = (*KyvernoIoGenerateRequestV1Resource)(nil)
)

type KyvernoIoGenerateRequestV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type KyvernoIoGenerateRequestV1GoModel struct {
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
		Context *struct {
			AdmissionRequestInfo *struct {
				AdmissionRequest *string `tfsdk:"admission_request" yaml:"admissionRequest,omitempty"`

				Operation *string `tfsdk:"operation" yaml:"operation,omitempty"`
			} `tfsdk:"admission_request_info" yaml:"admissionRequestInfo,omitempty"`

			UserInfo *struct {
				ClusterRoles *[]string `tfsdk:"cluster_roles" yaml:"clusterRoles,omitempty"`

				Roles *[]string `tfsdk:"roles" yaml:"roles,omitempty"`

				UserInfo *struct {
					Extra *map[string][]string `tfsdk:"extra" yaml:"extra,omitempty"`

					Groups *[]string `tfsdk:"groups" yaml:"groups,omitempty"`

					Uid *string `tfsdk:"uid" yaml:"uid,omitempty"`

					Username *string `tfsdk:"username" yaml:"username,omitempty"`
				} `tfsdk:"user_info" yaml:"userInfo,omitempty"`
			} `tfsdk:"user_info" yaml:"userInfo,omitempty"`
		} `tfsdk:"context" yaml:"context,omitempty"`

		Policy *string `tfsdk:"policy" yaml:"policy,omitempty"`

		Resource *struct {
			ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
		} `tfsdk:"resource" yaml:"resource,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewKyvernoIoGenerateRequestV1Resource() resource.Resource {
	return &KyvernoIoGenerateRequestV1Resource{}
}

func (r *KyvernoIoGenerateRequestV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_kyverno_io_generate_request_v1"
}

func (r *KyvernoIoGenerateRequestV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "GenerateRequest is a request to process generate rule.",
		MarkdownDescription: "GenerateRequest is a request to process generate rule.",
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
				Description:         "Spec is the information to identify the generate request.",
				MarkdownDescription: "Spec is the information to identify the generate request.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"context": {
						Description:         "Context ...",
						MarkdownDescription: "Context ...",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"admission_request_info": {
								Description:         "AdmissionRequestInfoObject stores the admission request and operation details",
								MarkdownDescription: "AdmissionRequestInfoObject stores the admission request and operation details",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"admission_request": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"operation": {
										Description:         "Operation is the type of resource operation being checked for admission control",
										MarkdownDescription: "Operation is the type of resource operation being checked for admission control",

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

							"user_info": {
								Description:         "RequestInfo contains permission info carried in an admission request.",
								MarkdownDescription: "RequestInfo contains permission info carried in an admission request.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"cluster_roles": {
										Description:         "ClusterRoles is a list of possible clusterRoles send the request.",
										MarkdownDescription: "ClusterRoles is a list of possible clusterRoles send the request.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"roles": {
										Description:         "Roles is a list of possible role send the request.",
										MarkdownDescription: "Roles is a list of possible role send the request.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"user_info": {
										Description:         "UserInfo is the userInfo carried in the admission request.",
										MarkdownDescription: "UserInfo is the userInfo carried in the admission request.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"extra": {
												Description:         "Any additional information provided by the authenticator.",
												MarkdownDescription: "Any additional information provided by the authenticator.",

												Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"groups": {
												Description:         "The names of groups this user is a part of.",
												MarkdownDescription: "The names of groups this user is a part of.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"uid": {
												Description:         "A unique value that identifies this user across time. If this user is deleted and another user by the same name is added, they will have different UIDs.",
												MarkdownDescription: "A unique value that identifies this user across time. If this user is deleted and another user by the same name is added, they will have different UIDs.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"username": {
												Description:         "The name that uniquely identifies this user among all active users.",
												MarkdownDescription: "The name that uniquely identifies this user among all active users.",

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

					"policy": {
						Description:         "Specifies the name of the policy.",
						MarkdownDescription: "Specifies the name of the policy.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"resource": {
						Description:         "ResourceSpec is the information to identify the generate request.",
						MarkdownDescription: "ResourceSpec is the information to identify the generate request.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"api_version": {
								Description:         "APIVersion specifies resource apiVersion.",
								MarkdownDescription: "APIVersion specifies resource apiVersion.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"kind": {
								Description:         "Kind specifies resource kind.",
								MarkdownDescription: "Kind specifies resource kind.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "Name specifies the resource name.",
								MarkdownDescription: "Name specifies the resource name.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"namespace": {
								Description:         "Namespace specifies resource namespace.",
								MarkdownDescription: "Namespace specifies resource namespace.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},
				}),

				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}, nil
}

func (r *KyvernoIoGenerateRequestV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_kyverno_io_generate_request_v1")

	var state KyvernoIoGenerateRequestV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel KyvernoIoGenerateRequestV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("kyverno.io/v1")
	goModel.Kind = utilities.Ptr("GenerateRequest")

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

func (r *KyvernoIoGenerateRequestV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kyverno_io_generate_request_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *KyvernoIoGenerateRequestV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_kyverno_io_generate_request_v1")

	var state KyvernoIoGenerateRequestV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel KyvernoIoGenerateRequestV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("kyverno.io/v1")
	goModel.Kind = utilities.Ptr("GenerateRequest")

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

func (r *KyvernoIoGenerateRequestV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_kyverno_io_generate_request_v1")
	// NO-OP: Terraform removes the state automatically for us
}
