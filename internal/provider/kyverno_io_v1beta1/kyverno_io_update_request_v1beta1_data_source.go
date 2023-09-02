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
	_ datasource.DataSource              = &KyvernoIoUpdateRequestV1Beta1DataSource{}
	_ datasource.DataSourceWithConfigure = &KyvernoIoUpdateRequestV1Beta1DataSource{}
)

func NewKyvernoIoUpdateRequestV1Beta1DataSource() datasource.DataSource {
	return &KyvernoIoUpdateRequestV1Beta1DataSource{}
}

type KyvernoIoUpdateRequestV1Beta1DataSource struct {
	kubernetesClient dynamic.Interface
}

type KyvernoIoUpdateRequestV1Beta1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

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

func (r *KyvernoIoUpdateRequestV1Beta1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kyverno_io_update_request_v1beta1"
}

func (r *KyvernoIoUpdateRequestV1Beta1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
												Optional:            false,
												Computed:            true,
											},

											"kind": schema.SingleNestedAttribute{
												Description:         "Kind is the fully-qualified type of object being submitted (for example, v1.Pod or autoscaling.v1.Scale)",
												MarkdownDescription: "Kind is the fully-qualified type of object being submitted (for example, v1.Pod or autoscaling.v1.Scale)",
												Attributes: map[string]schema.Attribute{
													"group": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"kind": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"version": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"name": schema.StringAttribute{
												Description:         "Name is the name of the object as presented in the request.  On a CREATE operation, the client may omit name and rely on the server to generate the name.  If that is the case, this field will contain an empty string.",
												MarkdownDescription: "Name is the name of the object as presented in the request.  On a CREATE operation, the client may omit name and rely on the server to generate the name.  If that is the case, this field will contain an empty string.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace is the namespace associated with the request (if any).",
												MarkdownDescription: "Namespace is the namespace associated with the request (if any).",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"object": schema.MapAttribute{
												Description:         "Object is the object from the incoming request.",
												MarkdownDescription: "Object is the object from the incoming request.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"old_object": schema.MapAttribute{
												Description:         "OldObject is the existing object. Only populated for DELETE and UPDATE requests.",
												MarkdownDescription: "OldObject is the existing object. Only populated for DELETE and UPDATE requests.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"operation": schema.StringAttribute{
												Description:         "Operation is the operation being performed. This may be different than the operation requested. e.g. a patch can result in either a CREATE or UPDATE Operation.",
												MarkdownDescription: "Operation is the operation being performed. This may be different than the operation requested. e.g. a patch can result in either a CREATE or UPDATE Operation.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"options": schema.MapAttribute{
												Description:         "Options is the operation option structure of the operation being performed. e.g. 'meta.k8s.io/v1.DeleteOptions' or 'meta.k8s.io/v1.CreateOptions'. This may be different than the options the caller provided. e.g. for a patch request the performed Operation might be a CREATE, in which case the Options will a 'meta.k8s.io/v1.CreateOptions' even though the caller provided 'meta.k8s.io/v1.PatchOptions'.",
												MarkdownDescription: "Options is the operation option structure of the operation being performed. e.g. 'meta.k8s.io/v1.DeleteOptions' or 'meta.k8s.io/v1.CreateOptions'. This may be different than the options the caller provided. e.g. for a patch request the performed Operation might be a CREATE, in which case the Options will a 'meta.k8s.io/v1.CreateOptions' even though the caller provided 'meta.k8s.io/v1.PatchOptions'.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"request_kind": schema.SingleNestedAttribute{
												Description:         "RequestKind is the fully-qualified type of the original API request (for example, v1.Pod or autoscaling.v1.Scale). If this is specified and differs from the value in 'kind', an equivalent match and conversion was performed.  For example, if deployments can be modified via apps/v1 and apps/v1beta1, and a webhook registered a rule of 'apiGroups:['apps'], apiVersions:['v1'], resources: ['deployments']' and 'matchPolicy: Equivalent', an API request to apps/v1beta1 deployments would be converted and sent to the webhook with 'kind: {group:'apps', version:'v1', kind:'Deployment'}' (matching the rule the webhook registered for), and 'requestKind: {group:'apps', version:'v1beta1', kind:'Deployment'}' (indicating the kind of the original API request).  See documentation for the 'matchPolicy' field in the webhook configuration type for more details.",
												MarkdownDescription: "RequestKind is the fully-qualified type of the original API request (for example, v1.Pod or autoscaling.v1.Scale). If this is specified and differs from the value in 'kind', an equivalent match and conversion was performed.  For example, if deployments can be modified via apps/v1 and apps/v1beta1, and a webhook registered a rule of 'apiGroups:['apps'], apiVersions:['v1'], resources: ['deployments']' and 'matchPolicy: Equivalent', an API request to apps/v1beta1 deployments would be converted and sent to the webhook with 'kind: {group:'apps', version:'v1', kind:'Deployment'}' (matching the rule the webhook registered for), and 'requestKind: {group:'apps', version:'v1beta1', kind:'Deployment'}' (indicating the kind of the original API request).  See documentation for the 'matchPolicy' field in the webhook configuration type for more details.",
												Attributes: map[string]schema.Attribute{
													"group": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"kind": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"version": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"request_resource": schema.SingleNestedAttribute{
												Description:         "RequestResource is the fully-qualified resource of the original API request (for example, v1.pods). If this is specified and differs from the value in 'resource', an equivalent match and conversion was performed.  For example, if deployments can be modified via apps/v1 and apps/v1beta1, and a webhook registered a rule of 'apiGroups:['apps'], apiVersions:['v1'], resources: ['deployments']' and 'matchPolicy: Equivalent', an API request to apps/v1beta1 deployments would be converted and sent to the webhook with 'resource: {group:'apps', version:'v1', resource:'deployments'}' (matching the resource the webhook registered for), and 'requestResource: {group:'apps', version:'v1beta1', resource:'deployments'}' (indicating the resource of the original API request).  See documentation for the 'matchPolicy' field in the webhook configuration type.",
												MarkdownDescription: "RequestResource is the fully-qualified resource of the original API request (for example, v1.pods). If this is specified and differs from the value in 'resource', an equivalent match and conversion was performed.  For example, if deployments can be modified via apps/v1 and apps/v1beta1, and a webhook registered a rule of 'apiGroups:['apps'], apiVersions:['v1'], resources: ['deployments']' and 'matchPolicy: Equivalent', an API request to apps/v1beta1 deployments would be converted and sent to the webhook with 'resource: {group:'apps', version:'v1', resource:'deployments'}' (matching the resource the webhook registered for), and 'requestResource: {group:'apps', version:'v1beta1', resource:'deployments'}' (indicating the resource of the original API request).  See documentation for the 'matchPolicy' field in the webhook configuration type.",
												Attributes: map[string]schema.Attribute{
													"group": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"resource": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"version": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"request_sub_resource": schema.StringAttribute{
												Description:         "RequestSubResource is the name of the subresource of the original API request, if any (for example, 'status' or 'scale') If this is specified and differs from the value in 'subResource', an equivalent match and conversion was performed. See documentation for the 'matchPolicy' field in the webhook configuration type.",
												MarkdownDescription: "RequestSubResource is the name of the subresource of the original API request, if any (for example, 'status' or 'scale') If this is specified and differs from the value in 'subResource', an equivalent match and conversion was performed. See documentation for the 'matchPolicy' field in the webhook configuration type.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"resource": schema.SingleNestedAttribute{
												Description:         "Resource is the fully-qualified resource being requested (for example, v1.pods)",
												MarkdownDescription: "Resource is the fully-qualified resource being requested (for example, v1.pods)",
												Attributes: map[string]schema.Attribute{
													"group": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"resource": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"version": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"sub_resource": schema.StringAttribute{
												Description:         "SubResource is the subresource being requested, if any (for example, 'status' or 'scale')",
												MarkdownDescription: "SubResource is the subresource being requested, if any (for example, 'status' or 'scale')",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"uid": schema.StringAttribute{
												Description:         "UID is an identifier for the individual request/response. It allows us to distinguish instances of requests which are otherwise identical (parallel requests, requests when earlier requests did not modify etc) The UID is meant to track the round trip (request/response) between the KAS and the WebHook, not the user request. It is suitable for correlating log entries between the webhook and apiserver, for either auditing or debugging.",
												MarkdownDescription: "UID is an identifier for the individual request/response. It allows us to distinguish instances of requests which are otherwise identical (parallel requests, requests when earlier requests did not modify etc) The UID is meant to track the round trip (request/response) between the KAS and the WebHook, not the user request. It is suitable for correlating log entries between the webhook and apiserver, for either auditing or debugging.",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
														Optional:            false,
														Computed:            true,
													},

													"groups": schema.ListAttribute{
														Description:         "The names of groups this user is a part of.",
														MarkdownDescription: "The names of groups this user is a part of.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"uid": schema.StringAttribute{
														Description:         "A unique value that identifies this user across time. If this user is deleted and another user by the same name is added, they will have different UIDs.",
														MarkdownDescription: "A unique value that identifies this user across time. If this user is deleted and another user by the same name is added, they will have different UIDs.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"username": schema.StringAttribute{
														Description:         "The name that uniquely identifies this user among all active users.",
														MarkdownDescription: "The name that uniquely identifies this user among all active users.",
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
										Required: false,
										Optional: false,
										Computed: true,
									},

									"operation": schema.StringAttribute{
										Description:         "Operation is the type of resource operation being checked for admission control",
										MarkdownDescription: "Operation is the type of resource operation being checked for admission control",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
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
										Optional:            false,
										Computed:            true,
									},

									"roles": schema.ListAttribute{
										Description:         "Roles is a list of possible role send the request.",
										MarkdownDescription: "Roles is a list of possible role send the request.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
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
												Optional:            false,
												Computed:            true,
											},

											"groups": schema.ListAttribute{
												Description:         "The names of groups this user is a part of.",
												MarkdownDescription: "The names of groups this user is a part of.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"uid": schema.StringAttribute{
												Description:         "A unique value that identifies this user across time. If this user is deleted and another user by the same name is added, they will have different UIDs.",
												MarkdownDescription: "A unique value that identifies this user across time. If this user is deleted and another user by the same name is added, they will have different UIDs.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"username": schema.StringAttribute{
												Description:         "The name that uniquely identifies this user among all active users.",
												MarkdownDescription: "The name that uniquely identifies this user among all active users.",
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
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"delete_downstream": schema.BoolAttribute{
						Description:         "DeleteDownstream represents whether the downstream needs to be deleted.",
						MarkdownDescription: "DeleteDownstream represents whether the downstream needs to be deleted.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"policy": schema.StringAttribute{
						Description:         "Specifies the name of the policy.",
						MarkdownDescription: "Specifies the name of the policy.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"request_type": schema.StringAttribute{
						Description:         "Type represents request type for background processing",
						MarkdownDescription: "Type represents request type for background processing",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"resource": schema.SingleNestedAttribute{
						Description:         "ResourceSpec is the information to identify the trigger resource.",
						MarkdownDescription: "ResourceSpec is the information to identify the trigger resource.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "APIVersion specifies resource apiVersion.",
								MarkdownDescription: "APIVersion specifies resource apiVersion.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind specifies resource kind.",
								MarkdownDescription: "Kind specifies resource kind.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "Name specifies the resource name.",
								MarkdownDescription: "Name specifies the resource name.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace specifies resource namespace.",
								MarkdownDescription: "Namespace specifies resource namespace.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"rule": schema.StringAttribute{
						Description:         "Rule is the associate rule name of the current UR.",
						MarkdownDescription: "Rule is the associate rule name of the current UR.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"synchronize": schema.BoolAttribute{
						Description:         "Synchronize represents the sync behavior of the corresponding rule Optional. Defaults to 'false' if not specified.",
						MarkdownDescription: "Synchronize represents the sync behavior of the corresponding rule Optional. Defaults to 'false' if not specified.",
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
	}
}

func (r *KyvernoIoUpdateRequestV1Beta1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *KyvernoIoUpdateRequestV1Beta1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_kyverno_io_update_request_v1beta1")

	var data KyvernoIoUpdateRequestV1Beta1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
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

	var readResponse KyvernoIoUpdateRequestV1Beta1DataSourceData
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

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("kyverno.io/v1beta1")
	data.Kind = pointer.String("UpdateRequest")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
