/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package flowcontrol_apiserver_k8s_io_v1beta3

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
	_ datasource.DataSource              = &FlowcontrolApiserverK8SIoFlowSchemaV1Beta3DataSource{}
	_ datasource.DataSourceWithConfigure = &FlowcontrolApiserverK8SIoFlowSchemaV1Beta3DataSource{}
)

func NewFlowcontrolApiserverK8SIoFlowSchemaV1Beta3DataSource() datasource.DataSource {
	return &FlowcontrolApiserverK8SIoFlowSchemaV1Beta3DataSource{}
}

type FlowcontrolApiserverK8SIoFlowSchemaV1Beta3DataSource struct {
	kubernetesClient dynamic.Interface
}

type FlowcontrolApiserverK8SIoFlowSchemaV1Beta3DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		DistinguisherMethod *struct {
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"distinguisher_method" json:"distinguisherMethod,omitempty"`
		MatchingPrecedence         *int64 `tfsdk:"matching_precedence" json:"matchingPrecedence,omitempty"`
		PriorityLevelConfiguration *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"priority_level_configuration" json:"priorityLevelConfiguration,omitempty"`
		Rules *[]struct {
			NonResourceRules *[]struct {
				NonResourceURLs *[]string `tfsdk:"non_resource_urls" json:"nonResourceURLs,omitempty"`
				Verbs           *[]string `tfsdk:"verbs" json:"verbs,omitempty"`
			} `tfsdk:"non_resource_rules" json:"nonResourceRules,omitempty"`
			ResourceRules *[]struct {
				ApiGroups    *[]string `tfsdk:"api_groups" json:"apiGroups,omitempty"`
				ClusterScope *bool     `tfsdk:"cluster_scope" json:"clusterScope,omitempty"`
				Namespaces   *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
				Resources    *[]string `tfsdk:"resources" json:"resources,omitempty"`
				Verbs        *[]string `tfsdk:"verbs" json:"verbs,omitempty"`
			} `tfsdk:"resource_rules" json:"resourceRules,omitempty"`
			Subjects *[]struct {
				Group *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"group" json:"group,omitempty"`
				Kind           *string `tfsdk:"kind" json:"kind,omitempty"`
				ServiceAccount *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				User *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"user" json:"user,omitempty"`
			} `tfsdk:"subjects" json:"subjects,omitempty"`
		} `tfsdk:"rules" json:"rules,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *FlowcontrolApiserverK8SIoFlowSchemaV1Beta3DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_flowcontrol_apiserver_k8s_io_flow_schema_v1beta3"
}

func (r *FlowcontrolApiserverK8SIoFlowSchemaV1Beta3DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "FlowSchema defines the schema of a group of flows. Note that a flow is made up of a set of inbound API requests with similar attributes and is identified by a pair of strings: the name of the FlowSchema and a 'flow distinguisher'.",
		MarkdownDescription: "FlowSchema defines the schema of a group of flows. Note that a flow is made up of a set of inbound API requests with similar attributes and is identified by a pair of strings: the name of the FlowSchema and a 'flow distinguisher'.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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
				Description:         "FlowSchemaSpec describes how the FlowSchema's specification looks like.",
				MarkdownDescription: "FlowSchemaSpec describes how the FlowSchema's specification looks like.",
				Attributes: map[string]schema.Attribute{
					"distinguisher_method": schema.SingleNestedAttribute{
						Description:         "FlowDistinguisherMethod specifies the method of a flow distinguisher.",
						MarkdownDescription: "FlowDistinguisherMethod specifies the method of a flow distinguisher.",
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Description:         "'type' is the type of flow distinguisher method The supported types are 'ByUser' and 'ByNamespace'. Required.",
								MarkdownDescription: "'type' is the type of flow distinguisher method The supported types are 'ByUser' and 'ByNamespace'. Required.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"matching_precedence": schema.Int64Attribute{
						Description:         "'matchingPrecedence' is used to choose among the FlowSchemas that match a given request. The chosen FlowSchema is among those with the numerically lowest (which we take to be logically highest) MatchingPrecedence.  Each MatchingPrecedence value must be ranged in [1,10000]. Note that if the precedence is not specified, it will be set to 1000 as default.",
						MarkdownDescription: "'matchingPrecedence' is used to choose among the FlowSchemas that match a given request. The chosen FlowSchema is among those with the numerically lowest (which we take to be logically highest) MatchingPrecedence.  Each MatchingPrecedence value must be ranged in [1,10000]. Note that if the precedence is not specified, it will be set to 1000 as default.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"priority_level_configuration": schema.SingleNestedAttribute{
						Description:         "PriorityLevelConfigurationReference contains information that points to the 'request-priority' being used.",
						MarkdownDescription: "PriorityLevelConfigurationReference contains information that points to the 'request-priority' being used.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "'name' is the name of the priority level configuration being referenced Required.",
								MarkdownDescription: "'name' is the name of the priority level configuration being referenced Required.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"rules": schema.ListNestedAttribute{
						Description:         "'rules' describes which requests will match this flow schema. This FlowSchema matches a request if and only if at least one member of rules matches the request. if it is an empty slice, there will be no requests matching the FlowSchema.",
						MarkdownDescription: "'rules' describes which requests will match this flow schema. This FlowSchema matches a request if and only if at least one member of rules matches the request. if it is an empty slice, there will be no requests matching the FlowSchema.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"non_resource_rules": schema.ListNestedAttribute{
									Description:         "'nonResourceRules' is a list of NonResourcePolicyRules that identify matching requests according to their verb and the target non-resource URL.",
									MarkdownDescription: "'nonResourceRules' is a list of NonResourcePolicyRules that identify matching requests according to their verb and the target non-resource URL.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"non_resource_urls": schema.ListAttribute{
												Description:         "'nonResourceURLs' is a set of url prefixes that a user should have access to and may not be empty. For example:  - '/healthz' is legal  - '/hea*' is illegal  - '/hea' is legal but matches nothing  - '/hea/*' also matches nothing  - '/healthz/*' matches all per-component health checks.'*' matches all non-resource urls. if it is present, it must be the only entry. Required.",
												MarkdownDescription: "'nonResourceURLs' is a set of url prefixes that a user should have access to and may not be empty. For example:  - '/healthz' is legal  - '/hea*' is illegal  - '/hea' is legal but matches nothing  - '/hea/*' also matches nothing  - '/healthz/*' matches all per-component health checks.'*' matches all non-resource urls. if it is present, it must be the only entry. Required.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"verbs": schema.ListAttribute{
												Description:         "'verbs' is a list of matching verbs and may not be empty. '*' matches all verbs. If it is present, it must be the only entry. Required.",
												MarkdownDescription: "'verbs' is a list of matching verbs and may not be empty. '*' matches all verbs. If it is present, it must be the only entry. Required.",
												ElementType:         types.StringType,
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

								"resource_rules": schema.ListNestedAttribute{
									Description:         "'resourceRules' is a slice of ResourcePolicyRules that identify matching requests according to their verb and the target resource. At least one of 'resourceRules' and 'nonResourceRules' has to be non-empty.",
									MarkdownDescription: "'resourceRules' is a slice of ResourcePolicyRules that identify matching requests according to their verb and the target resource. At least one of 'resourceRules' and 'nonResourceRules' has to be non-empty.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"api_groups": schema.ListAttribute{
												Description:         "'apiGroups' is a list of matching API groups and may not be empty. '*' matches all API groups and, if present, must be the only entry. Required.",
												MarkdownDescription: "'apiGroups' is a list of matching API groups and may not be empty. '*' matches all API groups and, if present, must be the only entry. Required.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"cluster_scope": schema.BoolAttribute{
												Description:         "'clusterScope' indicates whether to match requests that do not specify a namespace (which happens either because the resource is not namespaced or the request targets all namespaces). If this field is omitted or false then the 'namespaces' field must contain a non-empty list.",
												MarkdownDescription: "'clusterScope' indicates whether to match requests that do not specify a namespace (which happens either because the resource is not namespaced or the request targets all namespaces). If this field is omitted or false then the 'namespaces' field must contain a non-empty list.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"namespaces": schema.ListAttribute{
												Description:         "'namespaces' is a list of target namespaces that restricts matches.  A request that specifies a target namespace matches only if either (a) this list contains that target namespace or (b) this list contains '*'.  Note that '*' matches any specified namespace but does not match a request that _does not specify_ a namespace (see the 'clusterScope' field for that). This list may be empty, but only if 'clusterScope' is true.",
												MarkdownDescription: "'namespaces' is a list of target namespaces that restricts matches.  A request that specifies a target namespace matches only if either (a) this list contains that target namespace or (b) this list contains '*'.  Note that '*' matches any specified namespace but does not match a request that _does not specify_ a namespace (see the 'clusterScope' field for that). This list may be empty, but only if 'clusterScope' is true.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"resources": schema.ListAttribute{
												Description:         "'resources' is a list of matching resources (i.e., lowercase and plural) with, if desired, subresource.  For example, [ 'services', 'nodes/status' ].  This list may not be empty. '*' matches all resources and, if present, must be the only entry. Required.",
												MarkdownDescription: "'resources' is a list of matching resources (i.e., lowercase and plural) with, if desired, subresource.  For example, [ 'services', 'nodes/status' ].  This list may not be empty. '*' matches all resources and, if present, must be the only entry. Required.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"verbs": schema.ListAttribute{
												Description:         "'verbs' is a list of matching verbs and may not be empty. '*' matches all verbs and, if present, must be the only entry. Required.",
												MarkdownDescription: "'verbs' is a list of matching verbs and may not be empty. '*' matches all verbs and, if present, must be the only entry. Required.",
												ElementType:         types.StringType,
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

								"subjects": schema.ListNestedAttribute{
									Description:         "subjects is the list of normal user, serviceaccount, or group that this rule cares about. There must be at least one member in this slice. A slice that includes both the system:authenticated and system:unauthenticated user groups matches every request. Required.",
									MarkdownDescription: "subjects is the list of normal user, serviceaccount, or group that this rule cares about. There must be at least one member in this slice. A slice that includes both the system:authenticated and system:unauthenticated user groups matches every request. Required.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"group": schema.SingleNestedAttribute{
												Description:         "GroupSubject holds detailed information for group-kind subject.",
												MarkdownDescription: "GroupSubject holds detailed information for group-kind subject.",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "name is the user group that matches, or '*' to match all user groups. See https://github.com/kubernetes/apiserver/blob/master/pkg/authentication/user/user.go for some well-known group names. Required.",
														MarkdownDescription: "name is the user group that matches, or '*' to match all user groups. See https://github.com/kubernetes/apiserver/blob/master/pkg/authentication/user/user.go for some well-known group names. Required.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"kind": schema.StringAttribute{
												Description:         "'kind' indicates which one of the other fields is non-empty. Required",
												MarkdownDescription: "'kind' indicates which one of the other fields is non-empty. Required",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"service_account": schema.SingleNestedAttribute{
												Description:         "ServiceAccountSubject holds detailed information for service-account-kind subject.",
												MarkdownDescription: "ServiceAccountSubject holds detailed information for service-account-kind subject.",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "'name' is the name of matching ServiceAccount objects, or '*' to match regardless of name. Required.",
														MarkdownDescription: "'name' is the name of matching ServiceAccount objects, or '*' to match regardless of name. Required.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"namespace": schema.StringAttribute{
														Description:         "'namespace' is the namespace of matching ServiceAccount objects. Required.",
														MarkdownDescription: "'namespace' is the namespace of matching ServiceAccount objects. Required.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"user": schema.SingleNestedAttribute{
												Description:         "UserSubject holds detailed information for user-kind subject.",
												MarkdownDescription: "UserSubject holds detailed information for user-kind subject.",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "'name' is the username that matches, or '*' to match all usernames. Required.",
														MarkdownDescription: "'name' is the username that matches, or '*' to match all usernames. Required.",
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
									},
									Required: false,
									Optional: false,
									Computed: true,
								},
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

func (r *FlowcontrolApiserverK8SIoFlowSchemaV1Beta3DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *FlowcontrolApiserverK8SIoFlowSchemaV1Beta3DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1beta3")

	var data FlowcontrolApiserverK8SIoFlowSchemaV1Beta3DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "flowcontrol.apiserver.k8s.io", Version: "v1beta3", Resource: "FlowSchema"}).
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

	var readResponse FlowcontrolApiserverK8SIoFlowSchemaV1Beta3DataSourceData
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

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("flowcontrol.apiserver.k8s.io/v1beta3")
	data.Kind = pointer.String("FlowSchema")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
