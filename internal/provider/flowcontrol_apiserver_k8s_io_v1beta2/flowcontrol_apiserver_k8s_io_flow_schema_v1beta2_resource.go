/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package flowcontrol_apiserver_k8s_io_v1beta2

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	k8sTypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"time"
)

var (
	_ resource.Resource                = &FlowcontrolApiserverK8SIoFlowSchemaV1Beta2Resource{}
	_ resource.ResourceWithConfigure   = &FlowcontrolApiserverK8SIoFlowSchemaV1Beta2Resource{}
	_ resource.ResourceWithImportState = &FlowcontrolApiserverK8SIoFlowSchemaV1Beta2Resource{}
)

func NewFlowcontrolApiserverK8SIoFlowSchemaV1Beta2Resource() resource.Resource {
	return &FlowcontrolApiserverK8SIoFlowSchemaV1Beta2Resource{}
}

type FlowcontrolApiserverK8SIoFlowSchemaV1Beta2Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type FlowcontrolApiserverK8SIoFlowSchemaV1Beta2ResourceData struct {
	ID                  types.String `tfsdk:"id" json:"-"`
	ForceConflicts      types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager        types.String `tfsdk:"field_manager" json:"-"`
	DeletionPropagation types.String `tfsdk:"deletion_propagation" json:"-"`
	WaitForUpsert       types.List   `tfsdk:"wait_for_upsert" json:"-"`
	WaitForDelete       types.Object `tfsdk:"wait_for_delete" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

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

func (r *FlowcontrolApiserverK8SIoFlowSchemaV1Beta2Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_flowcontrol_apiserver_k8s_io_flow_schema_v1beta2"
}

func (r *FlowcontrolApiserverK8SIoFlowSchemaV1Beta2Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
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

			"force_conflicts": schema.BoolAttribute{
				Description:         "If 'true', server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "If `true`, server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"field_manager": schema.StringAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},

			"deletion_propagation": schema.StringAttribute{
				Description:         "Decides if a deletion will propagate to the dependents of the object, and how the garbage collector will handle the propagation.",
				MarkdownDescription: "Decides if a deletion will propagate to the dependents of the object, and how the garbage collector will handle the propagation.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("Orphan", "Background", "Foreground"),
				},
			},

			"wait_for_upsert": schema.ListNestedAttribute{
				Description:         "Wait for specific conditions after create/update of resources.",
				MarkdownDescription: "Wait for specific conditions after create/update of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"jsonpath": schema.StringAttribute{
							Description:         "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							MarkdownDescription: "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},
						"value": schema.StringAttribute{
							Description:         "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							MarkdownDescription: "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							Required:            false,
							Optional:            true,
							Computed:            true,
						},
						"timeout": schema.Int64Attribute{
							Description:         "The number of seconds to wait before giving up. Zero means check once and don't wait.",
							MarkdownDescription: "The number of seconds to wait before giving up. Zero means check once and don't wait.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(30),
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
						"poll_interval": schema.Int64Attribute{
							Description:         "The number of seconds to wait before checking again.",
							MarkdownDescription: "The number of seconds to wait before checking again.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(5),
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
					},
				},
			},

			"wait_for_delete": schema.SingleNestedAttribute{
				Description:         "Wait for deletion of resources.",
				MarkdownDescription: "Wait for deletion of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"timeout": schema.Int64Attribute{
						Description:         "The number of seconds to wait before giving up. Zero means check once and don't wait.",
						MarkdownDescription: "The number of seconds to wait before giving up. Zero means check once and don't wait.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             int64default.StaticInt64(30),
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
					"poll_interval": schema.Int64Attribute{
						Description:         "The number of seconds to wait before checking again.",
						MarkdownDescription: "The number of seconds to wait before checking again.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             int64default.StaticInt64(5),
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
				},
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            true,
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
						Computed:            true,
						Validators: []validator.Map{
							validators.AnnotationValidator(),
						},
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
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"matching_precedence": schema.Int64Attribute{
						Description:         "'matchingPrecedence' is used to choose among the FlowSchemas that match a given request. The chosen FlowSchema is among those with the numerically lowest (which we take to be logically highest) MatchingPrecedence.  Each MatchingPrecedence value must be ranged in [1,10000]. Note that if the precedence is not specified, it will be set to 1000 as default.",
						MarkdownDescription: "'matchingPrecedence' is used to choose among the FlowSchemas that match a given request. The chosen FlowSchema is among those with the numerically lowest (which we take to be logically highest) MatchingPrecedence.  Each MatchingPrecedence value must be ranged in [1,10000]. Note that if the precedence is not specified, it will be set to 1000 as default.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"priority_level_configuration": schema.SingleNestedAttribute{
						Description:         "PriorityLevelConfigurationReference contains information that points to the 'request-priority' being used.",
						MarkdownDescription: "PriorityLevelConfigurationReference contains information that points to the 'request-priority' being used.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "'name' is the name of the priority level configuration being referenced Required.",
								MarkdownDescription: "'name' is the name of the priority level configuration being referenced Required.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
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
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"verbs": schema.ListAttribute{
												Description:         "'verbs' is a list of matching verbs and may not be empty. '*' matches all verbs. If it is present, it must be the only entry. Required.",
												MarkdownDescription: "'verbs' is a list of matching verbs and may not be empty. '*' matches all verbs. If it is present, it must be the only entry. Required.",
												ElementType:         types.StringType,
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
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
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"cluster_scope": schema.BoolAttribute{
												Description:         "'clusterScope' indicates whether to match requests that do not specify a namespace (which happens either because the resource is not namespaced or the request targets all namespaces). If this field is omitted or false then the 'namespaces' field must contain a non-empty list.",
												MarkdownDescription: "'clusterScope' indicates whether to match requests that do not specify a namespace (which happens either because the resource is not namespaced or the request targets all namespaces). If this field is omitted or false then the 'namespaces' field must contain a non-empty list.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespaces": schema.ListAttribute{
												Description:         "'namespaces' is a list of target namespaces that restricts matches.  A request that specifies a target namespace matches only if either (a) this list contains that target namespace or (b) this list contains '*'.  Note that '*' matches any specified namespace but does not match a request that _does not specify_ a namespace (see the 'clusterScope' field for that). This list may be empty, but only if 'clusterScope' is true.",
												MarkdownDescription: "'namespaces' is a list of target namespaces that restricts matches.  A request that specifies a target namespace matches only if either (a) this list contains that target namespace or (b) this list contains '*'.  Note that '*' matches any specified namespace but does not match a request that _does not specify_ a namespace (see the 'clusterScope' field for that). This list may be empty, but only if 'clusterScope' is true.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"resources": schema.ListAttribute{
												Description:         "'resources' is a list of matching resources (i.e., lowercase and plural) with, if desired, subresource.  For example, [ 'services', 'nodes/status' ].  This list may not be empty. '*' matches all resources and, if present, must be the only entry. Required.",
												MarkdownDescription: "'resources' is a list of matching resources (i.e., lowercase and plural) with, if desired, subresource.  For example, [ 'services', 'nodes/status' ].  This list may not be empty. '*' matches all resources and, if present, must be the only entry. Required.",
												ElementType:         types.StringType,
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"verbs": schema.ListAttribute{
												Description:         "'verbs' is a list of matching verbs and may not be empty. '*' matches all verbs and, if present, must be the only entry. Required.",
												MarkdownDescription: "'verbs' is a list of matching verbs and may not be empty. '*' matches all verbs and, if present, must be the only entry. Required.",
												ElementType:         types.StringType,
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
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
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"kind": schema.StringAttribute{
												Description:         "'kind' indicates which one of the other fields is non-empty. Required",
												MarkdownDescription: "'kind' indicates which one of the other fields is non-empty. Required",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"service_account": schema.SingleNestedAttribute{
												Description:         "ServiceAccountSubject holds detailed information for service-account-kind subject.",
												MarkdownDescription: "ServiceAccountSubject holds detailed information for service-account-kind subject.",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "'name' is the name of matching ServiceAccount objects, or '*' to match regardless of name. Required.",
														MarkdownDescription: "'name' is the name of matching ServiceAccount objects, or '*' to match regardless of name. Required.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "'namespace' is the namespace of matching ServiceAccount objects. Required.",
														MarkdownDescription: "'namespace' is the namespace of matching ServiceAccount objects. Required.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"user": schema.SingleNestedAttribute{
												Description:         "UserSubject holds detailed information for user-kind subject.",
												MarkdownDescription: "UserSubject holds detailed information for user-kind subject.",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "'name' is the username that matches, or '*' to match all usernames. Required.",
														MarkdownDescription: "'name' is the username that matches, or '*' to match all usernames. Required.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
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

func (r *FlowcontrolApiserverK8SIoFlowSchemaV1Beta2Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedResourceDataError(request.ProviderData))
	}
}

func (r *FlowcontrolApiserverK8SIoFlowSchemaV1Beta2Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1beta2")

	var model FlowcontrolApiserverK8SIoFlowSchemaV1Beta2ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("flowcontrol.apiserver.k8s.io/v1beta2")
	model.Kind = pointer.String("FlowSchema")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "flowcontrol.apiserver.k8s.io", Version: "v1beta2", Resource: "flowschemas"}).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse FlowcontrolApiserverK8SIoFlowSchemaV1Beta2ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec
	if model.ForceConflicts.IsUnknown() {
		model.ForceConflicts = types.BoolNull()
	}
	if model.FieldManager.IsUnknown() {
		model.FieldManager = types.StringNull()
	}
	if model.DeletionPropagation.IsUnknown() {
		model.DeletionPropagation = types.StringNull()
	}
	if model.WaitForUpsert.IsUnknown() {
		model.WaitForUpsert = types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"jsonpath":      types.StringType,
				"value":         types.StringType,
				"timeout":       types.Int64Type,
				"poll_interval": types.Int64Type,
			},
		})
	}
	if model.WaitForDelete.IsUnknown() {
		model.WaitForDelete = types.ObjectNull(map[string]attr.Type{
			"timeout":       types.Int64Type,
			"poll_interval": types.Int64Type,
		})
	}

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *FlowcontrolApiserverK8SIoFlowSchemaV1Beta2Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1beta2")

	var data FlowcontrolApiserverK8SIoFlowSchemaV1Beta2ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "flowcontrol.apiserver.k8s.io", Version: "v1beta2", Resource: "flowschemas"}).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetResourceError(err, data.Metadata.Name))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse FlowcontrolApiserverK8SIoFlowSchemaV1Beta2ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec
	if data.ForceConflicts.IsUnknown() {
		data.ForceConflicts = types.BoolNull()
	}
	if data.FieldManager.IsUnknown() {
		data.FieldManager = types.StringNull()
	}
	if data.DeletionPropagation.IsUnknown() {
		data.DeletionPropagation = types.StringNull()
	}
	if data.WaitForUpsert.IsUnknown() {
		data.WaitForUpsert = types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"jsonpath":      types.StringType,
				"value":         types.StringType,
				"timeout":       types.Int64Type,
				"poll_interval": types.Int64Type,
			},
		})
	}
	if data.WaitForDelete.IsUnknown() {
		data.WaitForDelete = types.ObjectNull(map[string]attr.Type{
			"timeout":       types.Int64Type,
			"poll_interval": types.Int64Type,
		})
	}

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *FlowcontrolApiserverK8SIoFlowSchemaV1Beta2Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1beta2")

	var model FlowcontrolApiserverK8SIoFlowSchemaV1Beta2ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("flowcontrol.apiserver.k8s.io/v1beta2")
	model.Kind = pointer.String("FlowSchema")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "flowcontrol.apiserver.k8s.io", Version: "v1beta2", Resource: "flowschemas"}).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse FlowcontrolApiserverK8SIoFlowSchemaV1Beta2ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *FlowcontrolApiserverK8SIoFlowSchemaV1Beta2Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1beta2")

	var data FlowcontrolApiserverK8SIoFlowSchemaV1Beta2ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	deleteOptions := meta.DeleteOptions{}
	if !data.DeletionPropagation.IsNull() && !data.DeletionPropagation.IsUnknown() {
		deleteOptions.PropagationPolicy = utilities.MapDeletionPropagation(data.DeletionPropagation.ValueString())
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "flowcontrol.apiserver.k8s.io", Version: "v1beta2", Resource: "flowschemas"}).
		Delete(ctx, data.Metadata.Name, deleteOptions)
	if utilities.IsDeletionError(err) {
		response.Diagnostics.Append(utilities.DeleteError(err))
		return
	}

	if !data.WaitForDelete.IsNull() && !data.WaitForDelete.IsUnknown() {
		timeout := utilities.DetermineTimeout(data.WaitForDelete.Attributes())
		pollInterval := utilities.DeterminePollInterval(data.WaitForDelete.Attributes())

		startTime := time.Now()
		for {
			_, err := r.kubernetesClient.
				Resource(k8sSchema.GroupVersionResource{Group: "flowcontrol.apiserver.k8s.io", Version: "v1beta2", Resource: "flowschemas"}).
				Get(ctx, data.Metadata.Name, meta.GetOptions{})
			if utilities.IsNotFound(err) || timeout.Milliseconds() == 0 {
				break
			}
			if time.Now().After(startTime.Add(timeout)) {
				response.Diagnostics.Append(utilities.WaitTimeoutExceeded())
				return
			}
			time.Sleep(pollInterval)
		}
	}
}

func (r *FlowcontrolApiserverK8SIoFlowSchemaV1Beta2Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	if request.ID == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'name' Got: '%q'", request.ID),
		)
		return
	}
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	resource.ImportStatePassthroughID(ctx, path.Root("metadata").AtName("name"), request, response)
}
