/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kyverno_io_v1beta1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
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
	"strings"
)

var (
	_ resource.Resource                = &KyvernoIoUpdateRequestV1Beta1Resource{}
	_ resource.ResourceWithConfigure   = &KyvernoIoUpdateRequestV1Beta1Resource{}
	_ resource.ResourceWithImportState = &KyvernoIoUpdateRequestV1Beta1Resource{}
)

func NewKyvernoIoUpdateRequestV1Beta1Resource() resource.Resource {
	return &KyvernoIoUpdateRequestV1Beta1Resource{}
}

type KyvernoIoUpdateRequestV1Beta1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type KyvernoIoUpdateRequestV1Beta1ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitFor        types.List   `tfsdk:"wait_for" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Context *struct {
			AdmissionRequestInfo *struct {
				AdmissionRequest *struct {
					DryRun *bool `tfsdk:"dry_run" json:"dryRun,omitempty"`
					Kind   *struct {
						Group   *string `tfsdk:"group" json:"group,omitempty"`
						Kind    *string `tfsdk:"kind" json:"kind,omitempty"`
						Version *string `tfsdk:"version" json:"version,omitempty"`
					} `tfsdk:"kind" json:"kind,omitempty"`
					Name        *string            `tfsdk:"name" json:"name,omitempty"`
					Namespace   *string            `tfsdk:"namespace" json:"namespace,omitempty"`
					Object      *map[string]string `tfsdk:"object" json:"object,omitempty"`
					OldObject   *map[string]string `tfsdk:"old_object" json:"oldObject,omitempty"`
					Operation   *string            `tfsdk:"operation" json:"operation,omitempty"`
					Options     *map[string]string `tfsdk:"options" json:"options,omitempty"`
					RequestKind *struct {
						Group   *string `tfsdk:"group" json:"group,omitempty"`
						Kind    *string `tfsdk:"kind" json:"kind,omitempty"`
						Version *string `tfsdk:"version" json:"version,omitempty"`
					} `tfsdk:"request_kind" json:"requestKind,omitempty"`
					RequestResource *struct {
						Group    *string `tfsdk:"group" json:"group,omitempty"`
						Resource *string `tfsdk:"resource" json:"resource,omitempty"`
						Version  *string `tfsdk:"version" json:"version,omitempty"`
					} `tfsdk:"request_resource" json:"requestResource,omitempty"`
					RequestSubResource *string `tfsdk:"request_sub_resource" json:"requestSubResource,omitempty"`
					Resource           *struct {
						Group    *string `tfsdk:"group" json:"group,omitempty"`
						Resource *string `tfsdk:"resource" json:"resource,omitempty"`
						Version  *string `tfsdk:"version" json:"version,omitempty"`
					} `tfsdk:"resource" json:"resource,omitempty"`
					SubResource *string `tfsdk:"sub_resource" json:"subResource,omitempty"`
					Uid         *string `tfsdk:"uid" json:"uid,omitempty"`
					UserInfo    *struct {
						Extra    *map[string][]string `tfsdk:"extra" json:"extra,omitempty"`
						Groups   *[]string            `tfsdk:"groups" json:"groups,omitempty"`
						Uid      *string              `tfsdk:"uid" json:"uid,omitempty"`
						Username *string              `tfsdk:"username" json:"username,omitempty"`
					} `tfsdk:"user_info" json:"userInfo,omitempty"`
				} `tfsdk:"admission_request" json:"admissionRequest,omitempty"`
				Operation *string `tfsdk:"operation" json:"operation,omitempty"`
			} `tfsdk:"admission_request_info" json:"admissionRequestInfo,omitempty"`
			UserInfo *struct {
				ClusterRoles *[]string `tfsdk:"cluster_roles" json:"clusterRoles,omitempty"`
				Roles        *[]string `tfsdk:"roles" json:"roles,omitempty"`
				UserInfo     *struct {
					Extra    *map[string][]string `tfsdk:"extra" json:"extra,omitempty"`
					Groups   *[]string            `tfsdk:"groups" json:"groups,omitempty"`
					Uid      *string              `tfsdk:"uid" json:"uid,omitempty"`
					Username *string              `tfsdk:"username" json:"username,omitempty"`
				} `tfsdk:"user_info" json:"userInfo,omitempty"`
			} `tfsdk:"user_info" json:"userInfo,omitempty"`
		} `tfsdk:"context" json:"context,omitempty"`
		DeleteDownstream *bool   `tfsdk:"delete_downstream" json:"deleteDownstream,omitempty"`
		Policy           *string `tfsdk:"policy" json:"policy,omitempty"`
		RequestType      *string `tfsdk:"request_type" json:"requestType,omitempty"`
		Resource         *struct {
			ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
			Namespace  *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"resource" json:"resource,omitempty"`
		Rule        *string `tfsdk:"rule" json:"rule,omitempty"`
		Synchronize *bool   `tfsdk:"synchronize" json:"synchronize,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KyvernoIoUpdateRequestV1Beta1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kyverno_io_update_request_v1beta1"
}

func (r *KyvernoIoUpdateRequestV1Beta1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "UpdateRequest is a request to process mutate and generate rules in background.",
		MarkdownDescription: "UpdateRequest is a request to process mutate and generate rules in background.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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

			"field_manager": schema.BoolAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"wait_for": schema.ListNestedAttribute{
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
						"timeout": schema.StringAttribute{
							Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("30s"),
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
				Description:         "ResourceSpec is the information to identify the trigger resource.",
				MarkdownDescription: "ResourceSpec is the information to identify the trigger resource.",
				Attributes: map[string]schema.Attribute{
					"context": schema.SingleNestedAttribute{
						Description:         "Context ...",
						MarkdownDescription: "Context ...",
						Attributes: map[string]schema.Attribute{
							"admission_request_info": schema.SingleNestedAttribute{
								Description:         "AdmissionRequestInfoObject stores the admission request and operation details",
								MarkdownDescription: "AdmissionRequestInfoObject stores the admission request and operation details",
								Attributes: map[string]schema.Attribute{
									"admission_request": schema.SingleNestedAttribute{
										Description:         "AdmissionRequest describes the admission.Attributes for the admission request.",
										MarkdownDescription: "AdmissionRequest describes the admission.Attributes for the admission request.",
										Attributes: map[string]schema.Attribute{
											"dry_run": schema.BoolAttribute{
												Description:         "DryRun indicates that modifications will definitely not be persisted for this request. Defaults to false.",
												MarkdownDescription: "DryRun indicates that modifications will definitely not be persisted for this request. Defaults to false.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"kind": schema.SingleNestedAttribute{
												Description:         "Kind is the fully-qualified type of object being submitted (for example, v1.Pod or autoscaling.v1.Scale)",
												MarkdownDescription: "Kind is the fully-qualified type of object being submitted (for example, v1.Pod or autoscaling.v1.Scale)",
												Attributes: map[string]schema.Attribute{
													"group": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"kind": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"version": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": schema.StringAttribute{
												Description:         "Name is the name of the object as presented in the request.  On a CREATE operation, the client may omit name and rely on the server to generate the name.  If that is the case, this field will contain an empty string.",
												MarkdownDescription: "Name is the name of the object as presented in the request.  On a CREATE operation, the client may omit name and rely on the server to generate the name.  If that is the case, this field will contain an empty string.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace is the namespace associated with the request (if any).",
												MarkdownDescription: "Namespace is the namespace associated with the request (if any).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"object": schema.MapAttribute{
												Description:         "Object is the object from the incoming request.",
												MarkdownDescription: "Object is the object from the incoming request.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"old_object": schema.MapAttribute{
												Description:         "OldObject is the existing object. Only populated for DELETE and UPDATE requests.",
												MarkdownDescription: "OldObject is the existing object. Only populated for DELETE and UPDATE requests.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"operation": schema.StringAttribute{
												Description:         "Operation is the operation being performed. This may be different than the operation requested. e.g. a patch can result in either a CREATE or UPDATE Operation.",
												MarkdownDescription: "Operation is the operation being performed. This may be different than the operation requested. e.g. a patch can result in either a CREATE or UPDATE Operation.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"options": schema.MapAttribute{
												Description:         "Options is the operation option structure of the operation being performed. e.g. 'meta.k8s.io/v1.DeleteOptions' or 'meta.k8s.io/v1.CreateOptions'. This may be different than the options the caller provided. e.g. for a patch request the performed Operation might be a CREATE, in which case the Options will a 'meta.k8s.io/v1.CreateOptions' even though the caller provided 'meta.k8s.io/v1.PatchOptions'.",
												MarkdownDescription: "Options is the operation option structure of the operation being performed. e.g. 'meta.k8s.io/v1.DeleteOptions' or 'meta.k8s.io/v1.CreateOptions'. This may be different than the options the caller provided. e.g. for a patch request the performed Operation might be a CREATE, in which case the Options will a 'meta.k8s.io/v1.CreateOptions' even though the caller provided 'meta.k8s.io/v1.PatchOptions'.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"request_kind": schema.SingleNestedAttribute{
												Description:         "RequestKind is the fully-qualified type of the original API request (for example, v1.Pod or autoscaling.v1.Scale). If this is specified and differs from the value in 'kind', an equivalent match and conversion was performed.  For example, if deployments can be modified via apps/v1 and apps/v1beta1, and a webhook registered a rule of 'apiGroups:['apps'], apiVersions:['v1'], resources: ['deployments']' and 'matchPolicy: Equivalent', an API request to apps/v1beta1 deployments would be converted and sent to the webhook with 'kind: {group:'apps', version:'v1', kind:'Deployment'}' (matching the rule the webhook registered for), and 'requestKind: {group:'apps', version:'v1beta1', kind:'Deployment'}' (indicating the kind of the original API request).  See documentation for the 'matchPolicy' field in the webhook configuration type for more details.",
												MarkdownDescription: "RequestKind is the fully-qualified type of the original API request (for example, v1.Pod or autoscaling.v1.Scale). If this is specified and differs from the value in 'kind', an equivalent match and conversion was performed.  For example, if deployments can be modified via apps/v1 and apps/v1beta1, and a webhook registered a rule of 'apiGroups:['apps'], apiVersions:['v1'], resources: ['deployments']' and 'matchPolicy: Equivalent', an API request to apps/v1beta1 deployments would be converted and sent to the webhook with 'kind: {group:'apps', version:'v1', kind:'Deployment'}' (matching the rule the webhook registered for), and 'requestKind: {group:'apps', version:'v1beta1', kind:'Deployment'}' (indicating the kind of the original API request).  See documentation for the 'matchPolicy' field in the webhook configuration type for more details.",
												Attributes: map[string]schema.Attribute{
													"group": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"kind": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"version": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"request_resource": schema.SingleNestedAttribute{
												Description:         "RequestResource is the fully-qualified resource of the original API request (for example, v1.pods). If this is specified and differs from the value in 'resource', an equivalent match and conversion was performed.  For example, if deployments can be modified via apps/v1 and apps/v1beta1, and a webhook registered a rule of 'apiGroups:['apps'], apiVersions:['v1'], resources: ['deployments']' and 'matchPolicy: Equivalent', an API request to apps/v1beta1 deployments would be converted and sent to the webhook with 'resource: {group:'apps', version:'v1', resource:'deployments'}' (matching the resource the webhook registered for), and 'requestResource: {group:'apps', version:'v1beta1', resource:'deployments'}' (indicating the resource of the original API request).  See documentation for the 'matchPolicy' field in the webhook configuration type.",
												MarkdownDescription: "RequestResource is the fully-qualified resource of the original API request (for example, v1.pods). If this is specified and differs from the value in 'resource', an equivalent match and conversion was performed.  For example, if deployments can be modified via apps/v1 and apps/v1beta1, and a webhook registered a rule of 'apiGroups:['apps'], apiVersions:['v1'], resources: ['deployments']' and 'matchPolicy: Equivalent', an API request to apps/v1beta1 deployments would be converted and sent to the webhook with 'resource: {group:'apps', version:'v1', resource:'deployments'}' (matching the resource the webhook registered for), and 'requestResource: {group:'apps', version:'v1beta1', resource:'deployments'}' (indicating the resource of the original API request).  See documentation for the 'matchPolicy' field in the webhook configuration type.",
												Attributes: map[string]schema.Attribute{
													"group": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"resource": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"version": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"request_sub_resource": schema.StringAttribute{
												Description:         "RequestSubResource is the name of the subresource of the original API request, if any (for example, 'status' or 'scale') If this is specified and differs from the value in 'subResource', an equivalent match and conversion was performed. See documentation for the 'matchPolicy' field in the webhook configuration type.",
												MarkdownDescription: "RequestSubResource is the name of the subresource of the original API request, if any (for example, 'status' or 'scale') If this is specified and differs from the value in 'subResource', an equivalent match and conversion was performed. See documentation for the 'matchPolicy' field in the webhook configuration type.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"resource": schema.SingleNestedAttribute{
												Description:         "Resource is the fully-qualified resource being requested (for example, v1.pods)",
												MarkdownDescription: "Resource is the fully-qualified resource being requested (for example, v1.pods)",
												Attributes: map[string]schema.Attribute{
													"group": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"resource": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"version": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: true,
												Optional: false,
												Computed: false,
											},

											"sub_resource": schema.StringAttribute{
												Description:         "SubResource is the subresource being requested, if any (for example, 'status' or 'scale')",
												MarkdownDescription: "SubResource is the subresource being requested, if any (for example, 'status' or 'scale')",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"uid": schema.StringAttribute{
												Description:         "UID is an identifier for the individual request/response. It allows us to distinguish instances of requests which are otherwise identical (parallel requests, requests when earlier requests did not modify etc) The UID is meant to track the round trip (request/response) between the KAS and the WebHook, not the user request. It is suitable for correlating log entries between the webhook and apiserver, for either auditing or debugging.",
												MarkdownDescription: "UID is an identifier for the individual request/response. It allows us to distinguish instances of requests which are otherwise identical (parallel requests, requests when earlier requests did not modify etc) The UID is meant to track the round trip (request/response) between the KAS and the WebHook, not the user request. It is suitable for correlating log entries between the webhook and apiserver, for either auditing or debugging.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"user_info": schema.SingleNestedAttribute{
												Description:         "UserInfo is information about the requesting user",
												MarkdownDescription: "UserInfo is information about the requesting user",
												Attributes: map[string]schema.Attribute{
													"extra": schema.MapAttribute{
														Description:         "Any additional information provided by the authenticator.",
														MarkdownDescription: "Any additional information provided by the authenticator.",
														ElementType:         types.ListType{ElemType: types.StringType},
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"groups": schema.ListAttribute{
														Description:         "The names of groups this user is a part of.",
														MarkdownDescription: "The names of groups this user is a part of.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"uid": schema.StringAttribute{
														Description:         "A unique value that identifies this user across time. If this user is deleted and another user by the same name is added, they will have different UIDs.",
														MarkdownDescription: "A unique value that identifies this user across time. If this user is deleted and another user by the same name is added, they will have different UIDs.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"username": schema.StringAttribute{
														Description:         "The name that uniquely identifies this user among all active users.",
														MarkdownDescription: "The name that uniquely identifies this user among all active users.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: true,
												Optional: false,
												Computed: false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"operation": schema.StringAttribute{
										Description:         "Operation is the type of resource operation being checked for admission control",
										MarkdownDescription: "Operation is the type of resource operation being checked for admission control",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"user_info": schema.SingleNestedAttribute{
								Description:         "RequestInfo contains permission info carried in an admission request.",
								MarkdownDescription: "RequestInfo contains permission info carried in an admission request.",
								Attributes: map[string]schema.Attribute{
									"cluster_roles": schema.ListAttribute{
										Description:         "ClusterRoles is a list of possible clusterRoles send the request.",
										MarkdownDescription: "ClusterRoles is a list of possible clusterRoles send the request.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"roles": schema.ListAttribute{
										Description:         "Roles is a list of possible role send the request.",
										MarkdownDescription: "Roles is a list of possible role send the request.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"user_info": schema.SingleNestedAttribute{
										Description:         "UserInfo is the userInfo carried in the admission request.",
										MarkdownDescription: "UserInfo is the userInfo carried in the admission request.",
										Attributes: map[string]schema.Attribute{
											"extra": schema.MapAttribute{
												Description:         "Any additional information provided by the authenticator.",
												MarkdownDescription: "Any additional information provided by the authenticator.",
												ElementType:         types.ListType{ElemType: types.StringType},
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"groups": schema.ListAttribute{
												Description:         "The names of groups this user is a part of.",
												MarkdownDescription: "The names of groups this user is a part of.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"uid": schema.StringAttribute{
												Description:         "A unique value that identifies this user across time. If this user is deleted and another user by the same name is added, they will have different UIDs.",
												MarkdownDescription: "A unique value that identifies this user across time. If this user is deleted and another user by the same name is added, they will have different UIDs.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"username": schema.StringAttribute{
												Description:         "The name that uniquely identifies this user among all active users.",
												MarkdownDescription: "The name that uniquely identifies this user among all active users.",
												Required:            false,
												Optional:            true,
												Computed:            false,
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
						Required: true,
						Optional: false,
						Computed: false,
					},

					"delete_downstream": schema.BoolAttribute{
						Description:         "DeleteDownstream represents whether the downstream needs to be deleted.",
						MarkdownDescription: "DeleteDownstream represents whether the downstream needs to be deleted.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"policy": schema.StringAttribute{
						Description:         "Specifies the name of the policy.",
						MarkdownDescription: "Specifies the name of the policy.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"request_type": schema.StringAttribute{
						Description:         "Type represents request type for background processing",
						MarkdownDescription: "Type represents request type for background processing",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("mutate", "generate"),
						},
					},

					"resource": schema.SingleNestedAttribute{
						Description:         "ResourceSpec is the information to identify the trigger resource.",
						MarkdownDescription: "ResourceSpec is the information to identify the trigger resource.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "APIVersion specifies resource apiVersion.",
								MarkdownDescription: "APIVersion specifies resource apiVersion.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind specifies resource kind.",
								MarkdownDescription: "Kind specifies resource kind.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name specifies the resource name.",
								MarkdownDescription: "Name specifies the resource name.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace specifies resource namespace.",
								MarkdownDescription: "Namespace specifies resource namespace.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"rule": schema.StringAttribute{
						Description:         "Rule is the associate rule name of the current UR.",
						MarkdownDescription: "Rule is the associate rule name of the current UR.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"synchronize": schema.BoolAttribute{
						Description:         "Synchronize represents the sync behavior of the corresponding rule Optional. Defaults to 'false' if not specified.",
						MarkdownDescription: "Synchronize represents the sync behavior of the corresponding rule Optional. Defaults to 'false' if not specified.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *KyvernoIoUpdateRequestV1Beta1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *dynamic.DynamicClient, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *KyvernoIoUpdateRequestV1Beta1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_kyverno_io_update_request_v1beta1")

	var model KyvernoIoUpdateRequestV1Beta1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("kyverno.io/v1beta1")
	model.Kind = pointer.String("UpdateRequest")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
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

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "kyverno.io", Version: "v1beta1", Resource: "UpdateRequest"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while creating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse KyvernoIoUpdateRequestV1Beta1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *KyvernoIoUpdateRequestV1Beta1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kyverno_io_update_request_v1beta1")

	var data KyvernoIoUpdateRequestV1Beta1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "kyverno.io", Version: "v1beta1", Resource: "UpdateRequest"}).
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

	var readResponse KyvernoIoUpdateRequestV1Beta1ResourceData
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

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *KyvernoIoUpdateRequestV1Beta1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_kyverno_io_update_request_v1beta1")

	var model KyvernoIoUpdateRequestV1Beta1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("kyverno.io/v1beta1")
	model.Kind = pointer.String("UpdateRequest")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
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

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "kyverno.io", Version: "v1beta1", Resource: "UpdateRequest"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while updating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse KyvernoIoUpdateRequestV1Beta1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *KyvernoIoUpdateRequestV1Beta1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_kyverno_io_update_request_v1beta1")

	var data KyvernoIoUpdateRequestV1Beta1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "kyverno.io", Version: "v1beta1", Resource: "UpdateRequest"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to DELETE resource",
			"An unexpected error occurred while deleting the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"DELETE Error: "+err.Error(),
		)
		return
	}
}

func (r *KyvernoIoUpdateRequestV1Beta1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	idParts := strings.Split(request.ID, "/")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'namespace/name' Got: '%q'", request.ID),
		)
		return
	}

	namespace := idParts[0]
	name := idParts[1]
	tflog.Trace(ctx, "parsed import ID", map[string]interface{}{
		"namespace": namespace,
		"name":      name,
	})
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("namespace"), namespace)...)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("name"), name)...)
}
