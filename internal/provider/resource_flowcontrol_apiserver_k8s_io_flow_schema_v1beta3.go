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

type FlowcontrolApiserverK8SIoFlowSchemaV1Beta3Resource struct{}

var (
	_ resource.Resource = (*FlowcontrolApiserverK8SIoFlowSchemaV1Beta3Resource)(nil)
)

type FlowcontrolApiserverK8SIoFlowSchemaV1Beta3TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type FlowcontrolApiserverK8SIoFlowSchemaV1Beta3GoModel struct {
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
		DistinguisherMethod *struct {
			Type *string `tfsdk:"type" yaml:"type,omitempty"`
		} `tfsdk:"distinguisher_method" yaml:"distinguisherMethod,omitempty"`

		MatchingPrecedence *int64 `tfsdk:"matching_precedence" yaml:"matchingPrecedence,omitempty"`

		PriorityLevelConfiguration *struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"priority_level_configuration" yaml:"priorityLevelConfiguration,omitempty"`

		Rules *[]struct {
			NonResourceRules *[]struct {
				NonResourceURLs *[]string `tfsdk:"non_resource_urls" yaml:"nonResourceURLs,omitempty"`

				Verbs *[]string `tfsdk:"verbs" yaml:"verbs,omitempty"`
			} `tfsdk:"non_resource_rules" yaml:"nonResourceRules,omitempty"`

			ResourceRules *[]struct {
				ApiGroups *[]string `tfsdk:"api_groups" yaml:"apiGroups,omitempty"`

				ClusterScope *bool `tfsdk:"cluster_scope" yaml:"clusterScope,omitempty"`

				Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

				Resources *[]string `tfsdk:"resources" yaml:"resources,omitempty"`

				Verbs *[]string `tfsdk:"verbs" yaml:"verbs,omitempty"`
			} `tfsdk:"resource_rules" yaml:"resourceRules,omitempty"`

			Subjects *[]struct {
				Group *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"group" yaml:"group,omitempty"`

				Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

				ServiceAccount *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
				} `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

				User *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"user" yaml:"user,omitempty"`
			} `tfsdk:"subjects" yaml:"subjects,omitempty"`
		} `tfsdk:"rules" yaml:"rules,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewFlowcontrolApiserverK8SIoFlowSchemaV1Beta3Resource() resource.Resource {
	return &FlowcontrolApiserverK8SIoFlowSchemaV1Beta3Resource{}
}

func (r *FlowcontrolApiserverK8SIoFlowSchemaV1Beta3Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_flowcontrol_apiserver_k8s_io_flow_schema_v1beta3"
}

func (r *FlowcontrolApiserverK8SIoFlowSchemaV1Beta3Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "FlowSchema defines the schema of a group of flows. Note that a flow is made up of a set of inbound API requests with similar attributes and is identified by a pair of strings: the name of the FlowSchema and a 'flow distinguisher'.",
		MarkdownDescription: "FlowSchema defines the schema of a group of flows. Note that a flow is made up of a set of inbound API requests with similar attributes and is identified by a pair of strings: the name of the FlowSchema and a 'flow distinguisher'.",
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
				Description:         "FlowSchemaSpec describes how the FlowSchema's specification looks like.",
				MarkdownDescription: "FlowSchemaSpec describes how the FlowSchema's specification looks like.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"distinguisher_method": {
						Description:         "FlowDistinguisherMethod specifies the method of a flow distinguisher.",
						MarkdownDescription: "FlowDistinguisherMethod specifies the method of a flow distinguisher.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"type": {
								Description:         "'type' is the type of flow distinguisher method The supported types are 'ByUser' and 'ByNamespace'. Required.",
								MarkdownDescription: "'type' is the type of flow distinguisher method The supported types are 'ByUser' and 'ByNamespace'. Required.",

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

					"matching_precedence": {
						Description:         "'matchingPrecedence' is used to choose among the FlowSchemas that match a given request. The chosen FlowSchema is among those with the numerically lowest (which we take to be logically highest) MatchingPrecedence.  Each MatchingPrecedence value must be ranged in [1,10000]. Note that if the precedence is not specified, it will be set to 1000 as default.",
						MarkdownDescription: "'matchingPrecedence' is used to choose among the FlowSchemas that match a given request. The chosen FlowSchema is among those with the numerically lowest (which we take to be logically highest) MatchingPrecedence.  Each MatchingPrecedence value must be ranged in [1,10000]. Note that if the precedence is not specified, it will be set to 1000 as default.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"priority_level_configuration": {
						Description:         "PriorityLevelConfigurationReference contains information that points to the 'request-priority' being used.",
						MarkdownDescription: "PriorityLevelConfigurationReference contains information that points to the 'request-priority' being used.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "'name' is the name of the priority level configuration being referenced Required.",
								MarkdownDescription: "'name' is the name of the priority level configuration being referenced Required.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"rules": {
						Description:         "'rules' describes which requests will match this flow schema. This FlowSchema matches a request if and only if at least one member of rules matches the request. if it is an empty slice, there will be no requests matching the FlowSchema.",
						MarkdownDescription: "'rules' describes which requests will match this flow schema. This FlowSchema matches a request if and only if at least one member of rules matches the request. if it is an empty slice, there will be no requests matching the FlowSchema.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"non_resource_rules": {
								Description:         "'nonResourceRules' is a list of NonResourcePolicyRules that identify matching requests according to their verb and the target non-resource URL.",
								MarkdownDescription: "'nonResourceRules' is a list of NonResourcePolicyRules that identify matching requests according to their verb and the target non-resource URL.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"non_resource_urls": {
										Description:         "'nonResourceURLs' is a set of url prefixes that a user should have access to and may not be empty. For example:  - '/healthz' is legal  - '/hea*' is illegal  - '/hea' is legal but matches nothing  - '/hea/*' also matches nothing  - '/healthz/*' matches all per-component health checks.'*' matches all non-resource urls. if it is present, it must be the only entry. Required.",
										MarkdownDescription: "'nonResourceURLs' is a set of url prefixes that a user should have access to and may not be empty. For example:  - '/healthz' is legal  - '/hea*' is illegal  - '/hea' is legal but matches nothing  - '/hea/*' also matches nothing  - '/healthz/*' matches all per-component health checks.'*' matches all non-resource urls. if it is present, it must be the only entry. Required.",

										Type: types.ListType{ElemType: types.StringType},

										Required: true,
										Optional: false,
										Computed: false,
									},

									"verbs": {
										Description:         "'verbs' is a list of matching verbs and may not be empty. '*' matches all verbs. If it is present, it must be the only entry. Required.",
										MarkdownDescription: "'verbs' is a list of matching verbs and may not be empty. '*' matches all verbs. If it is present, it must be the only entry. Required.",

										Type: types.ListType{ElemType: types.StringType},

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resource_rules": {
								Description:         "'resourceRules' is a slice of ResourcePolicyRules that identify matching requests according to their verb and the target resource. At least one of 'resourceRules' and 'nonResourceRules' has to be non-empty.",
								MarkdownDescription: "'resourceRules' is a slice of ResourcePolicyRules that identify matching requests according to their verb and the target resource. At least one of 'resourceRules' and 'nonResourceRules' has to be non-empty.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"api_groups": {
										Description:         "'apiGroups' is a list of matching API groups and may not be empty. '*' matches all API groups and, if present, must be the only entry. Required.",
										MarkdownDescription: "'apiGroups' is a list of matching API groups and may not be empty. '*' matches all API groups and, if present, must be the only entry. Required.",

										Type: types.ListType{ElemType: types.StringType},

										Required: true,
										Optional: false,
										Computed: false,
									},

									"cluster_scope": {
										Description:         "'clusterScope' indicates whether to match requests that do not specify a namespace (which happens either because the resource is not namespaced or the request targets all namespaces). If this field is omitted or false then the 'namespaces' field must contain a non-empty list.",
										MarkdownDescription: "'clusterScope' indicates whether to match requests that do not specify a namespace (which happens either because the resource is not namespaced or the request targets all namespaces). If this field is omitted or false then the 'namespaces' field must contain a non-empty list.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"namespaces": {
										Description:         "'namespaces' is a list of target namespaces that restricts matches.  A request that specifies a target namespace matches only if either (a) this list contains that target namespace or (b) this list contains '*'.  Note that '*' matches any specified namespace but does not match a request that _does not specify_ a namespace (see the 'clusterScope' field for that). This list may be empty, but only if 'clusterScope' is true.",
										MarkdownDescription: "'namespaces' is a list of target namespaces that restricts matches.  A request that specifies a target namespace matches only if either (a) this list contains that target namespace or (b) this list contains '*'.  Note that '*' matches any specified namespace but does not match a request that _does not specify_ a namespace (see the 'clusterScope' field for that). This list may be empty, but only if 'clusterScope' is true.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"resources": {
										Description:         "'resources' is a list of matching resources (i.e., lowercase and plural) with, if desired, subresource.  For example, [ 'services', 'nodes/status' ].  This list may not be empty. '*' matches all resources and, if present, must be the only entry. Required.",
										MarkdownDescription: "'resources' is a list of matching resources (i.e., lowercase and plural) with, if desired, subresource.  For example, [ 'services', 'nodes/status' ].  This list may not be empty. '*' matches all resources and, if present, must be the only entry. Required.",

										Type: types.ListType{ElemType: types.StringType},

										Required: true,
										Optional: false,
										Computed: false,
									},

									"verbs": {
										Description:         "'verbs' is a list of matching verbs and may not be empty. '*' matches all verbs and, if present, must be the only entry. Required.",
										MarkdownDescription: "'verbs' is a list of matching verbs and may not be empty. '*' matches all verbs and, if present, must be the only entry. Required.",

										Type: types.ListType{ElemType: types.StringType},

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"subjects": {
								Description:         "subjects is the list of normal user, serviceaccount, or group that this rule cares about. There must be at least one member in this slice. A slice that includes both the system:authenticated and system:unauthenticated user groups matches every request. Required.",
								MarkdownDescription: "subjects is the list of normal user, serviceaccount, or group that this rule cares about. There must be at least one member in this slice. A slice that includes both the system:authenticated and system:unauthenticated user groups matches every request. Required.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"group": {
										Description:         "GroupSubject holds detailed information for group-kind subject.",
										MarkdownDescription: "GroupSubject holds detailed information for group-kind subject.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "name is the user group that matches, or '*' to match all user groups. See https://github.com/kubernetes/apiserver/blob/master/pkg/authentication/user/user.go for some well-known group names. Required.",
												MarkdownDescription: "name is the user group that matches, or '*' to match all user groups. See https://github.com/kubernetes/apiserver/blob/master/pkg/authentication/user/user.go for some well-known group names. Required.",

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

									"kind": {
										Description:         "'kind' indicates which one of the other fields is non-empty. Required",
										MarkdownDescription: "'kind' indicates which one of the other fields is non-empty. Required",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"service_account": {
										Description:         "ServiceAccountSubject holds detailed information for service-account-kind subject.",
										MarkdownDescription: "ServiceAccountSubject holds detailed information for service-account-kind subject.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "'name' is the name of matching ServiceAccount objects, or '*' to match regardless of name. Required.",
												MarkdownDescription: "'name' is the name of matching ServiceAccount objects, or '*' to match regardless of name. Required.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"namespace": {
												Description:         "'namespace' is the namespace of matching ServiceAccount objects. Required.",
												MarkdownDescription: "'namespace' is the namespace of matching ServiceAccount objects. Required.",

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

									"user": {
										Description:         "UserSubject holds detailed information for user-kind subject.",
										MarkdownDescription: "UserSubject holds detailed information for user-kind subject.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "'name' is the username that matches, or '*' to match all usernames. Required.",
												MarkdownDescription: "'name' is the username that matches, or '*' to match all usernames. Required.",

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

func (r *FlowcontrolApiserverK8SIoFlowSchemaV1Beta3Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1beta3")

	var state FlowcontrolApiserverK8SIoFlowSchemaV1Beta3TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel FlowcontrolApiserverK8SIoFlowSchemaV1Beta3GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("flowcontrol.apiserver.k8s.io/v1beta3")
	goModel.Kind = utilities.Ptr("FlowSchema")

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

func (r *FlowcontrolApiserverK8SIoFlowSchemaV1Beta3Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1beta3")
	// NO-OP: All data is already in Terraform state
}

func (r *FlowcontrolApiserverK8SIoFlowSchemaV1Beta3Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1beta3")

	var state FlowcontrolApiserverK8SIoFlowSchemaV1Beta3TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel FlowcontrolApiserverK8SIoFlowSchemaV1Beta3GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("flowcontrol.apiserver.k8s.io/v1beta3")
	goModel.Kind = utilities.Ptr("FlowSchema")

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

func (r *FlowcontrolApiserverK8SIoFlowSchemaV1Beta3Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1beta3")
	// NO-OP: Terraform removes the state automatically for us
}
