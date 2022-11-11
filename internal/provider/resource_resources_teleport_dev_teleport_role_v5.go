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

type ResourcesTeleportDevTeleportRoleV5Resource struct{}

var (
	_ resource.Resource = (*ResourcesTeleportDevTeleportRoleV5Resource)(nil)
)

type ResourcesTeleportDevTeleportRoleV5TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type ResourcesTeleportDevTeleportRoleV5GoModel struct {
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
		Allow *struct {
			App_labels *map[string]string `tfsdk:"app_labels" yaml:"app_labels,omitempty"`

			Aws_role_arns *[]string `tfsdk:"aws_role_arns" yaml:"aws_role_arns,omitempty"`

			Cluster_labels *map[string]string `tfsdk:"cluster_labels" yaml:"cluster_labels,omitempty"`

			Db_labels *map[string]string `tfsdk:"db_labels" yaml:"db_labels,omitempty"`

			Db_names *[]string `tfsdk:"db_names" yaml:"db_names,omitempty"`

			Db_users *[]string `tfsdk:"db_users" yaml:"db_users,omitempty"`

			Host_groups *[]string `tfsdk:"host_groups" yaml:"host_groups,omitempty"`

			Host_sudoers *[]string `tfsdk:"host_sudoers" yaml:"host_sudoers,omitempty"`

			Impersonate *struct {
				Roles *[]string `tfsdk:"roles" yaml:"roles,omitempty"`

				Users *[]string `tfsdk:"users" yaml:"users,omitempty"`

				Where *string `tfsdk:"where" yaml:"where,omitempty"`
			} `tfsdk:"impersonate" yaml:"impersonate,omitempty"`

			Join_sessions *[]struct {
				Kinds *[]string `tfsdk:"kinds" yaml:"kinds,omitempty"`

				Modes *[]string `tfsdk:"modes" yaml:"modes,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Roles *[]string `tfsdk:"roles" yaml:"roles,omitempty"`
			} `tfsdk:"join_sessions" yaml:"join_sessions,omitempty"`

			Kubernetes_groups *[]string `tfsdk:"kubernetes_groups" yaml:"kubernetes_groups,omitempty"`

			Kubernetes_labels *map[string]string `tfsdk:"kubernetes_labels" yaml:"kubernetes_labels,omitempty"`

			Kubernetes_users *[]string `tfsdk:"kubernetes_users" yaml:"kubernetes_users,omitempty"`

			Logins *[]string `tfsdk:"logins" yaml:"logins,omitempty"`

			Node_labels *map[string]string `tfsdk:"node_labels" yaml:"node_labels,omitempty"`

			Request *struct {
				Annotations *map[string][]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

				Claims_to_roles *[]struct {
					Claim *string `tfsdk:"claim" yaml:"claim,omitempty"`

					Roles *[]string `tfsdk:"roles" yaml:"roles,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"claims_to_roles" yaml:"claims_to_roles,omitempty"`

				Roles *[]string `tfsdk:"roles" yaml:"roles,omitempty"`

				Search_as_roles *[]string `tfsdk:"search_as_roles" yaml:"search_as_roles,omitempty"`

				Suggested_reviewers *[]string `tfsdk:"suggested_reviewers" yaml:"suggested_reviewers,omitempty"`

				Thresholds *[]struct {
					Approve *int64 `tfsdk:"approve" yaml:"approve,omitempty"`

					Deny *int64 `tfsdk:"deny" yaml:"deny,omitempty"`

					Filter *string `tfsdk:"filter" yaml:"filter,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"thresholds" yaml:"thresholds,omitempty"`
			} `tfsdk:"request" yaml:"request,omitempty"`

			Require_session_join *[]struct {
				Count *int64 `tfsdk:"count" yaml:"count,omitempty"`

				Filter *string `tfsdk:"filter" yaml:"filter,omitempty"`

				Kinds *[]string `tfsdk:"kinds" yaml:"kinds,omitempty"`

				Modes *[]string `tfsdk:"modes" yaml:"modes,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				On_leave *string `tfsdk:"on_leave" yaml:"on_leave,omitempty"`
			} `tfsdk:"require_session_join" yaml:"require_session_join,omitempty"`

			Review_requests *struct {
				Claims_to_roles *[]struct {
					Claim *string `tfsdk:"claim" yaml:"claim,omitempty"`

					Roles *[]string `tfsdk:"roles" yaml:"roles,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"claims_to_roles" yaml:"claims_to_roles,omitempty"`

				Preview_as_roles *[]string `tfsdk:"preview_as_roles" yaml:"preview_as_roles,omitempty"`

				Roles *[]string `tfsdk:"roles" yaml:"roles,omitempty"`

				Where *string `tfsdk:"where" yaml:"where,omitempty"`
			} `tfsdk:"review_requests" yaml:"review_requests,omitempty"`

			Rules *[]struct {
				Actions *[]string `tfsdk:"actions" yaml:"actions,omitempty"`

				Resources *[]string `tfsdk:"resources" yaml:"resources,omitempty"`

				Verbs *[]string `tfsdk:"verbs" yaml:"verbs,omitempty"`

				Where *string `tfsdk:"where" yaml:"where,omitempty"`
			} `tfsdk:"rules" yaml:"rules,omitempty"`

			Windows_desktop_labels *map[string]string `tfsdk:"windows_desktop_labels" yaml:"windows_desktop_labels,omitempty"`

			Windows_desktop_logins *[]string `tfsdk:"windows_desktop_logins" yaml:"windows_desktop_logins,omitempty"`
		} `tfsdk:"allow" yaml:"allow,omitempty"`

		Deny *struct {
			App_labels *map[string]string `tfsdk:"app_labels" yaml:"app_labels,omitempty"`

			Aws_role_arns *[]string `tfsdk:"aws_role_arns" yaml:"aws_role_arns,omitempty"`

			Cluster_labels *map[string]string `tfsdk:"cluster_labels" yaml:"cluster_labels,omitempty"`

			Db_labels *map[string]string `tfsdk:"db_labels" yaml:"db_labels,omitempty"`

			Db_names *[]string `tfsdk:"db_names" yaml:"db_names,omitempty"`

			Db_users *[]string `tfsdk:"db_users" yaml:"db_users,omitempty"`

			Host_groups *[]string `tfsdk:"host_groups" yaml:"host_groups,omitempty"`

			Host_sudoers *[]string `tfsdk:"host_sudoers" yaml:"host_sudoers,omitempty"`

			Impersonate *struct {
				Roles *[]string `tfsdk:"roles" yaml:"roles,omitempty"`

				Users *[]string `tfsdk:"users" yaml:"users,omitempty"`

				Where *string `tfsdk:"where" yaml:"where,omitempty"`
			} `tfsdk:"impersonate" yaml:"impersonate,omitempty"`

			Join_sessions *[]struct {
				Kinds *[]string `tfsdk:"kinds" yaml:"kinds,omitempty"`

				Modes *[]string `tfsdk:"modes" yaml:"modes,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Roles *[]string `tfsdk:"roles" yaml:"roles,omitempty"`
			} `tfsdk:"join_sessions" yaml:"join_sessions,omitempty"`

			Kubernetes_groups *[]string `tfsdk:"kubernetes_groups" yaml:"kubernetes_groups,omitempty"`

			Kubernetes_labels *map[string]string `tfsdk:"kubernetes_labels" yaml:"kubernetes_labels,omitempty"`

			Kubernetes_users *[]string `tfsdk:"kubernetes_users" yaml:"kubernetes_users,omitempty"`

			Logins *[]string `tfsdk:"logins" yaml:"logins,omitempty"`

			Node_labels *map[string]string `tfsdk:"node_labels" yaml:"node_labels,omitempty"`

			Request *struct {
				Annotations *map[string][]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

				Claims_to_roles *[]struct {
					Claim *string `tfsdk:"claim" yaml:"claim,omitempty"`

					Roles *[]string `tfsdk:"roles" yaml:"roles,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"claims_to_roles" yaml:"claims_to_roles,omitempty"`

				Roles *[]string `tfsdk:"roles" yaml:"roles,omitempty"`

				Search_as_roles *[]string `tfsdk:"search_as_roles" yaml:"search_as_roles,omitempty"`

				Suggested_reviewers *[]string `tfsdk:"suggested_reviewers" yaml:"suggested_reviewers,omitempty"`

				Thresholds *[]struct {
					Approve *int64 `tfsdk:"approve" yaml:"approve,omitempty"`

					Deny *int64 `tfsdk:"deny" yaml:"deny,omitempty"`

					Filter *string `tfsdk:"filter" yaml:"filter,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"thresholds" yaml:"thresholds,omitempty"`
			} `tfsdk:"request" yaml:"request,omitempty"`

			Require_session_join *[]struct {
				Count *int64 `tfsdk:"count" yaml:"count,omitempty"`

				Filter *string `tfsdk:"filter" yaml:"filter,omitempty"`

				Kinds *[]string `tfsdk:"kinds" yaml:"kinds,omitempty"`

				Modes *[]string `tfsdk:"modes" yaml:"modes,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				On_leave *string `tfsdk:"on_leave" yaml:"on_leave,omitempty"`
			} `tfsdk:"require_session_join" yaml:"require_session_join,omitempty"`

			Review_requests *struct {
				Claims_to_roles *[]struct {
					Claim *string `tfsdk:"claim" yaml:"claim,omitempty"`

					Roles *[]string `tfsdk:"roles" yaml:"roles,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"claims_to_roles" yaml:"claims_to_roles,omitempty"`

				Preview_as_roles *[]string `tfsdk:"preview_as_roles" yaml:"preview_as_roles,omitempty"`

				Roles *[]string `tfsdk:"roles" yaml:"roles,omitempty"`

				Where *string `tfsdk:"where" yaml:"where,omitempty"`
			} `tfsdk:"review_requests" yaml:"review_requests,omitempty"`

			Rules *[]struct {
				Actions *[]string `tfsdk:"actions" yaml:"actions,omitempty"`

				Resources *[]string `tfsdk:"resources" yaml:"resources,omitempty"`

				Verbs *[]string `tfsdk:"verbs" yaml:"verbs,omitempty"`

				Where *string `tfsdk:"where" yaml:"where,omitempty"`
			} `tfsdk:"rules" yaml:"rules,omitempty"`

			Windows_desktop_labels *map[string]string `tfsdk:"windows_desktop_labels" yaml:"windows_desktop_labels,omitempty"`

			Windows_desktop_logins *[]string `tfsdk:"windows_desktop_logins" yaml:"windows_desktop_logins,omitempty"`
		} `tfsdk:"deny" yaml:"deny,omitempty"`

		Options *struct {
			Cert_extensions *[]struct {
				Mode *int64 `tfsdk:"mode" yaml:"mode,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Type *int64 `tfsdk:"type" yaml:"type,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`
			} `tfsdk:"cert_extensions" yaml:"cert_extensions,omitempty"`

			Cert_format *string `tfsdk:"cert_format" yaml:"cert_format,omitempty"`

			Client_idle_timeout *string `tfsdk:"client_idle_timeout" yaml:"client_idle_timeout,omitempty"`

			Create_host_user *bool `tfsdk:"create_host_user" yaml:"create_host_user,omitempty"`

			Desktop_clipboard *bool `tfsdk:"desktop_clipboard" yaml:"desktop_clipboard,omitempty"`

			Desktop_directory_sharing *bool `tfsdk:"desktop_directory_sharing" yaml:"desktop_directory_sharing,omitempty"`

			Disconnect_expired_cert *bool `tfsdk:"disconnect_expired_cert" yaml:"disconnect_expired_cert,omitempty"`

			Enhanced_recording *[]string `tfsdk:"enhanced_recording" yaml:"enhanced_recording,omitempty"`

			Forward_agent *bool `tfsdk:"forward_agent" yaml:"forward_agent,omitempty"`

			Lock *string `tfsdk:"lock" yaml:"lock,omitempty"`

			Max_connections *int64 `tfsdk:"max_connections" yaml:"max_connections,omitempty"`

			Max_kubernetes_connections *int64 `tfsdk:"max_kubernetes_connections" yaml:"max_kubernetes_connections,omitempty"`

			Max_session_ttl *string `tfsdk:"max_session_ttl" yaml:"max_session_ttl,omitempty"`

			Max_sessions *int64 `tfsdk:"max_sessions" yaml:"max_sessions,omitempty"`

			Permit_x11_forwarding *bool `tfsdk:"permit_x11_forwarding" yaml:"permit_x11_forwarding,omitempty"`

			Pin_source_ip *bool `tfsdk:"pin_source_ip" yaml:"pin_source_ip,omitempty"`

			Port_forwarding *bool `tfsdk:"port_forwarding" yaml:"port_forwarding,omitempty"`

			Record_session *struct {
				Default *string `tfsdk:"default" yaml:"default,omitempty"`

				Desktop *bool `tfsdk:"desktop" yaml:"desktop,omitempty"`

				Ssh *string `tfsdk:"ssh" yaml:"ssh,omitempty"`
			} `tfsdk:"record_session" yaml:"record_session,omitempty"`

			Request_access *string `tfsdk:"request_access" yaml:"request_access,omitempty"`

			Request_prompt *string `tfsdk:"request_prompt" yaml:"request_prompt,omitempty"`

			Require_session_mfa *int64 `tfsdk:"require_session_mfa" yaml:"require_session_mfa,omitempty"`

			Ssh_file_copy *bool `tfsdk:"ssh_file_copy" yaml:"ssh_file_copy,omitempty"`
		} `tfsdk:"options" yaml:"options,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewResourcesTeleportDevTeleportRoleV5Resource() resource.Resource {
	return &ResourcesTeleportDevTeleportRoleV5Resource{}
}

func (r *ResourcesTeleportDevTeleportRoleV5Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_resources_teleport_dev_teleport_role_v5"
}

func (r *ResourcesTeleportDevTeleportRoleV5Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Role is the Schema for the roles API",
		MarkdownDescription: "Role is the Schema for the roles API",
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
				Description:         "Role resource definition v5 from Teleport",
				MarkdownDescription: "Role resource definition v5 from Teleport",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"allow": {
						Description:         "Allow is the set of conditions evaluated to grant access.",
						MarkdownDescription: "Allow is the set of conditions evaluated to grant access.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"app_labels": {
								Description:         "AppLabels is a map of labels used as part of the RBAC system.",
								MarkdownDescription: "AppLabels is a map of labels used as part of the RBAC system.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"aws_role_arns": {
								Description:         "AWSRoleARNs is a list of AWS role ARNs this role is allowed to assume.",
								MarkdownDescription: "AWSRoleARNs is a list of AWS role ARNs this role is allowed to assume.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cluster_labels": {
								Description:         "ClusterLabels is a map of node labels (used to dynamically grant access to clusters).",
								MarkdownDescription: "ClusterLabels is a map of node labels (used to dynamically grant access to clusters).",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"db_labels": {
								Description:         "DatabaseLabels are used in RBAC system to allow/deny access to databases.",
								MarkdownDescription: "DatabaseLabels are used in RBAC system to allow/deny access to databases.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"db_names": {
								Description:         "DatabaseNames is a list of database names this role is allowed to connect to.",
								MarkdownDescription: "DatabaseNames is a list of database names this role is allowed to connect to.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"db_users": {
								Description:         "DatabaseUsers is a list of databaes users this role is allowed to connect as.",
								MarkdownDescription: "DatabaseUsers is a list of databaes users this role is allowed to connect as.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"host_groups": {
								Description:         "HostGroups is a list of groups for created users to be added to",
								MarkdownDescription: "HostGroups is a list of groups for created users to be added to",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"host_sudoers": {
								Description:         "HostSudoers is a list of entries to include in a users sudoer file",
								MarkdownDescription: "HostSudoers is a list of entries to include in a users sudoer file",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"impersonate": {
								Description:         "Impersonate specifies what users and roles this role is allowed to impersonate by issuing certificates or other possible means.",
								MarkdownDescription: "Impersonate specifies what users and roles this role is allowed to impersonate by issuing certificates or other possible means.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"roles": {
										Description:         "Roles is a list of resources this role is allowed to impersonate",
										MarkdownDescription: "Roles is a list of resources this role is allowed to impersonate",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"users": {
										Description:         "Users is a list of resources this role is allowed to impersonate, could be an empty list or a Wildcard pattern",
										MarkdownDescription: "Users is a list of resources this role is allowed to impersonate, could be an empty list or a Wildcard pattern",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"where": {
										Description:         "Where specifies optional advanced matcher",
										MarkdownDescription: "Where specifies optional advanced matcher",

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

							"join_sessions": {
								Description:         "JoinSessions specifies policies to allow users to join other sessions.",
								MarkdownDescription: "JoinSessions specifies policies to allow users to join other sessions.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"kinds": {
										Description:         "Kinds are the session kinds this policy applies to.",
										MarkdownDescription: "Kinds are the session kinds this policy applies to.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"modes": {
										Description:         "Modes is a list of permitted participant modes for this policy.",
										MarkdownDescription: "Modes is a list of permitted participant modes for this policy.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "Name is the name of the policy.",
										MarkdownDescription: "Name is the name of the policy.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"roles": {
										Description:         "Roles is a list of roles that you can join the session of.",
										MarkdownDescription: "Roles is a list of roles that you can join the session of.",

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

							"kubernetes_groups": {
								Description:         "KubeGroups is a list of kubernetes groups",
								MarkdownDescription: "KubeGroups is a list of kubernetes groups",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"kubernetes_labels": {
								Description:         "KubernetesLabels is a map of kubernetes cluster labels used for RBAC.",
								MarkdownDescription: "KubernetesLabels is a map of kubernetes cluster labels used for RBAC.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"kubernetes_users": {
								Description:         "KubeUsers is an optional kubernetes users to impersonate",
								MarkdownDescription: "KubeUsers is an optional kubernetes users to impersonate",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"logins": {
								Description:         "Logins is a list of *nix system logins.",
								MarkdownDescription: "Logins is a list of *nix system logins.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"node_labels": {
								Description:         "NodeLabels is a map of node labels (used to dynamically grant access to nodes).",
								MarkdownDescription: "NodeLabels is a map of node labels (used to dynamically grant access to nodes).",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"request": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"annotations": {
										Description:         "Annotations is a collection of annotations to be programmatically appended to pending access requests at the time of their creation. These annotations serve as a mechanism to propagate extra information to plugins.  Since these annotations support variable interpolation syntax, they also offer a mechanism for forwarding claims from an external identity provider, to a plugin via '{{external.trait_name}}' style substitutions.",
										MarkdownDescription: "Annotations is a collection of annotations to be programmatically appended to pending access requests at the time of their creation. These annotations serve as a mechanism to propagate extra information to plugins.  Since these annotations support variable interpolation syntax, they also offer a mechanism for forwarding claims from an external identity provider, to a plugin via '{{external.trait_name}}' style substitutions.",

										Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"claims_to_roles": {
										Description:         "ClaimsToRoles specifies a mapping from claims (traits) to teleport roles.",
										MarkdownDescription: "ClaimsToRoles specifies a mapping from claims (traits) to teleport roles.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"claim": {
												Description:         "Claim is a claim name.",
												MarkdownDescription: "Claim is a claim name.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"roles": {
												Description:         "Roles is a list of static teleport roles to match.",
												MarkdownDescription: "Roles is a list of static teleport roles to match.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value": {
												Description:         "Value is a claim value to match.",
												MarkdownDescription: "Value is a claim value to match.",

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

									"roles": {
										Description:         "Roles is the name of roles which will match the request rule.",
										MarkdownDescription: "Roles is the name of roles which will match the request rule.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"search_as_roles": {
										Description:         "SearchAsRoles is a list of extra roles which should apply to a user while they are searching for resources as part of a Resource Access Request, and defines the underlying roles which will be requested as part of any Resource Access Request.",
										MarkdownDescription: "SearchAsRoles is a list of extra roles which should apply to a user while they are searching for resources as part of a Resource Access Request, and defines the underlying roles which will be requested as part of any Resource Access Request.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"suggested_reviewers": {
										Description:         "SuggestedReviewers is a list of reviewer suggestions.  These can be teleport usernames, but that is not a requirement.",
										MarkdownDescription: "SuggestedReviewers is a list of reviewer suggestions.  These can be teleport usernames, but that is not a requirement.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"thresholds": {
										Description:         "Thresholds is a list of thresholds, one of which must be met in order for reviews to trigger a state-transition.  If no thresholds are provided, a default threshold of 1 for approval and denial is used.",
										MarkdownDescription: "Thresholds is a list of thresholds, one of which must be met in order for reviews to trigger a state-transition.  If no thresholds are provided, a default threshold of 1 for approval and denial is used.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"approve": {
												Description:         "Approve is the number of matching approvals needed for state-transition.",
												MarkdownDescription: "Approve is the number of matching approvals needed for state-transition.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"deny": {
												Description:         "Deny is the number of denials needed for state-transition.",
												MarkdownDescription: "Deny is the number of denials needed for state-transition.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"filter": {
												Description:         "Filter is an optional predicate used to determine which reviews count toward this threshold.",
												MarkdownDescription: "Filter is an optional predicate used to determine which reviews count toward this threshold.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "Name is the optional human-readable name of the threshold.",
												MarkdownDescription: "Name is the optional human-readable name of the threshold.",

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

							"require_session_join": {
								Description:         "RequireSessionJoin specifies policies for required users to start a session.",
								MarkdownDescription: "RequireSessionJoin specifies policies for required users to start a session.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"count": {
										Description:         "Count is the amount of people that need to be matched for this policy to be fulfilled.",
										MarkdownDescription: "Count is the amount of people that need to be matched for this policy to be fulfilled.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"filter": {
										Description:         "Filter is a predicate that determines what users count towards this policy.",
										MarkdownDescription: "Filter is a predicate that determines what users count towards this policy.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"kinds": {
										Description:         "Kinds are the session kinds this policy applies to.",
										MarkdownDescription: "Kinds are the session kinds this policy applies to.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"modes": {
										Description:         "Modes is the list of modes that may be used to fulfill this policy.",
										MarkdownDescription: "Modes is the list of modes that may be used to fulfill this policy.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "Name is the name of the policy.",
										MarkdownDescription: "Name is the name of the policy.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"on_leave": {
										Description:         "OnLeave is the behaviour that's used when the policy is no longer fulfilled for a live session.",
										MarkdownDescription: "OnLeave is the behaviour that's used when the policy is no longer fulfilled for a live session.",

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

							"review_requests": {
								Description:         "ReviewRequests defines conditions for submitting access reviews.",
								MarkdownDescription: "ReviewRequests defines conditions for submitting access reviews.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"claims_to_roles": {
										Description:         "ClaimsToRoles specifies a mapping from claims (traits) to teleport roles.",
										MarkdownDescription: "ClaimsToRoles specifies a mapping from claims (traits) to teleport roles.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"claim": {
												Description:         "Claim is a claim name.",
												MarkdownDescription: "Claim is a claim name.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"roles": {
												Description:         "Roles is a list of static teleport roles to match.",
												MarkdownDescription: "Roles is a list of static teleport roles to match.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value": {
												Description:         "Value is a claim value to match.",
												MarkdownDescription: "Value is a claim value to match.",

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

									"preview_as_roles": {
										Description:         "PreviewAsRoles is a list of extra roles which should apply to a reviewer while they are viewing a Resource Access Request for the purposes of viewing details such as the hostname and labels of requested resources.",
										MarkdownDescription: "PreviewAsRoles is a list of extra roles which should apply to a reviewer while they are viewing a Resource Access Request for the purposes of viewing details such as the hostname and labels of requested resources.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"roles": {
										Description:         "Roles is the name of roles which may be reviewed.",
										MarkdownDescription: "Roles is the name of roles which may be reviewed.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"where": {
										Description:         "Where is an optional predicate which further limits which requests are reviewable.",
										MarkdownDescription: "Where is an optional predicate which further limits which requests are reviewable.",

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
								Description:         "Rules is a list of rules and their access levels. Rules are a high level construct used for access control.",
								MarkdownDescription: "Rules is a list of rules and their access levels. Rules are a high level construct used for access control.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"actions": {
										Description:         "Actions specifies optional actions taken when this rule matches",
										MarkdownDescription: "Actions specifies optional actions taken when this rule matches",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"resources": {
										Description:         "Resources is a list of resources",
										MarkdownDescription: "Resources is a list of resources",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"verbs": {
										Description:         "Verbs is a list of verbs",
										MarkdownDescription: "Verbs is a list of verbs",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"where": {
										Description:         "Where specifies optional advanced matcher",
										MarkdownDescription: "Where specifies optional advanced matcher",

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

							"windows_desktop_labels": {
								Description:         "WindowsDesktopLabels are used in the RBAC system to allow/deny access to Windows desktops.",
								MarkdownDescription: "WindowsDesktopLabels are used in the RBAC system to allow/deny access to Windows desktops.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"windows_desktop_logins": {
								Description:         "WindowsDesktopLogins is a list of desktop login names allowed/denied for Windows desktops.",
								MarkdownDescription: "WindowsDesktopLogins is a list of desktop login names allowed/denied for Windows desktops.",

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

					"deny": {
						Description:         "Deny is the set of conditions evaluated to deny access. Deny takes priority over allow.",
						MarkdownDescription: "Deny is the set of conditions evaluated to deny access. Deny takes priority over allow.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"app_labels": {
								Description:         "AppLabels is a map of labels used as part of the RBAC system.",
								MarkdownDescription: "AppLabels is a map of labels used as part of the RBAC system.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"aws_role_arns": {
								Description:         "AWSRoleARNs is a list of AWS role ARNs this role is allowed to assume.",
								MarkdownDescription: "AWSRoleARNs is a list of AWS role ARNs this role is allowed to assume.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cluster_labels": {
								Description:         "ClusterLabels is a map of node labels (used to dynamically grant access to clusters).",
								MarkdownDescription: "ClusterLabels is a map of node labels (used to dynamically grant access to clusters).",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"db_labels": {
								Description:         "DatabaseLabels are used in RBAC system to allow/deny access to databases.",
								MarkdownDescription: "DatabaseLabels are used in RBAC system to allow/deny access to databases.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"db_names": {
								Description:         "DatabaseNames is a list of database names this role is allowed to connect to.",
								MarkdownDescription: "DatabaseNames is a list of database names this role is allowed to connect to.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"db_users": {
								Description:         "DatabaseUsers is a list of databaes users this role is allowed to connect as.",
								MarkdownDescription: "DatabaseUsers is a list of databaes users this role is allowed to connect as.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"host_groups": {
								Description:         "HostGroups is a list of groups for created users to be added to",
								MarkdownDescription: "HostGroups is a list of groups for created users to be added to",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"host_sudoers": {
								Description:         "HostSudoers is a list of entries to include in a users sudoer file",
								MarkdownDescription: "HostSudoers is a list of entries to include in a users sudoer file",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"impersonate": {
								Description:         "Impersonate specifies what users and roles this role is allowed to impersonate by issuing certificates or other possible means.",
								MarkdownDescription: "Impersonate specifies what users and roles this role is allowed to impersonate by issuing certificates or other possible means.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"roles": {
										Description:         "Roles is a list of resources this role is allowed to impersonate",
										MarkdownDescription: "Roles is a list of resources this role is allowed to impersonate",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"users": {
										Description:         "Users is a list of resources this role is allowed to impersonate, could be an empty list or a Wildcard pattern",
										MarkdownDescription: "Users is a list of resources this role is allowed to impersonate, could be an empty list or a Wildcard pattern",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"where": {
										Description:         "Where specifies optional advanced matcher",
										MarkdownDescription: "Where specifies optional advanced matcher",

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

							"join_sessions": {
								Description:         "JoinSessions specifies policies to allow users to join other sessions.",
								MarkdownDescription: "JoinSessions specifies policies to allow users to join other sessions.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"kinds": {
										Description:         "Kinds are the session kinds this policy applies to.",
										MarkdownDescription: "Kinds are the session kinds this policy applies to.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"modes": {
										Description:         "Modes is a list of permitted participant modes for this policy.",
										MarkdownDescription: "Modes is a list of permitted participant modes for this policy.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "Name is the name of the policy.",
										MarkdownDescription: "Name is the name of the policy.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"roles": {
										Description:         "Roles is a list of roles that you can join the session of.",
										MarkdownDescription: "Roles is a list of roles that you can join the session of.",

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

							"kubernetes_groups": {
								Description:         "KubeGroups is a list of kubernetes groups",
								MarkdownDescription: "KubeGroups is a list of kubernetes groups",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"kubernetes_labels": {
								Description:         "KubernetesLabels is a map of kubernetes cluster labels used for RBAC.",
								MarkdownDescription: "KubernetesLabels is a map of kubernetes cluster labels used for RBAC.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"kubernetes_users": {
								Description:         "KubeUsers is an optional kubernetes users to impersonate",
								MarkdownDescription: "KubeUsers is an optional kubernetes users to impersonate",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"logins": {
								Description:         "Logins is a list of *nix system logins.",
								MarkdownDescription: "Logins is a list of *nix system logins.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"node_labels": {
								Description:         "NodeLabels is a map of node labels (used to dynamically grant access to nodes).",
								MarkdownDescription: "NodeLabels is a map of node labels (used to dynamically grant access to nodes).",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"request": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"annotations": {
										Description:         "Annotations is a collection of annotations to be programmatically appended to pending access requests at the time of their creation. These annotations serve as a mechanism to propagate extra information to plugins.  Since these annotations support variable interpolation syntax, they also offer a mechanism for forwarding claims from an external identity provider, to a plugin via '{{external.trait_name}}' style substitutions.",
										MarkdownDescription: "Annotations is a collection of annotations to be programmatically appended to pending access requests at the time of their creation. These annotations serve as a mechanism to propagate extra information to plugins.  Since these annotations support variable interpolation syntax, they also offer a mechanism for forwarding claims from an external identity provider, to a plugin via '{{external.trait_name}}' style substitutions.",

										Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"claims_to_roles": {
										Description:         "ClaimsToRoles specifies a mapping from claims (traits) to teleport roles.",
										MarkdownDescription: "ClaimsToRoles specifies a mapping from claims (traits) to teleport roles.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"claim": {
												Description:         "Claim is a claim name.",
												MarkdownDescription: "Claim is a claim name.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"roles": {
												Description:         "Roles is a list of static teleport roles to match.",
												MarkdownDescription: "Roles is a list of static teleport roles to match.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value": {
												Description:         "Value is a claim value to match.",
												MarkdownDescription: "Value is a claim value to match.",

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

									"roles": {
										Description:         "Roles is the name of roles which will match the request rule.",
										MarkdownDescription: "Roles is the name of roles which will match the request rule.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"search_as_roles": {
										Description:         "SearchAsRoles is a list of extra roles which should apply to a user while they are searching for resources as part of a Resource Access Request, and defines the underlying roles which will be requested as part of any Resource Access Request.",
										MarkdownDescription: "SearchAsRoles is a list of extra roles which should apply to a user while they are searching for resources as part of a Resource Access Request, and defines the underlying roles which will be requested as part of any Resource Access Request.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"suggested_reviewers": {
										Description:         "SuggestedReviewers is a list of reviewer suggestions.  These can be teleport usernames, but that is not a requirement.",
										MarkdownDescription: "SuggestedReviewers is a list of reviewer suggestions.  These can be teleport usernames, but that is not a requirement.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"thresholds": {
										Description:         "Thresholds is a list of thresholds, one of which must be met in order for reviews to trigger a state-transition.  If no thresholds are provided, a default threshold of 1 for approval and denial is used.",
										MarkdownDescription: "Thresholds is a list of thresholds, one of which must be met in order for reviews to trigger a state-transition.  If no thresholds are provided, a default threshold of 1 for approval and denial is used.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"approve": {
												Description:         "Approve is the number of matching approvals needed for state-transition.",
												MarkdownDescription: "Approve is the number of matching approvals needed for state-transition.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"deny": {
												Description:         "Deny is the number of denials needed for state-transition.",
												MarkdownDescription: "Deny is the number of denials needed for state-transition.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"filter": {
												Description:         "Filter is an optional predicate used to determine which reviews count toward this threshold.",
												MarkdownDescription: "Filter is an optional predicate used to determine which reviews count toward this threshold.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "Name is the optional human-readable name of the threshold.",
												MarkdownDescription: "Name is the optional human-readable name of the threshold.",

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

							"require_session_join": {
								Description:         "RequireSessionJoin specifies policies for required users to start a session.",
								MarkdownDescription: "RequireSessionJoin specifies policies for required users to start a session.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"count": {
										Description:         "Count is the amount of people that need to be matched for this policy to be fulfilled.",
										MarkdownDescription: "Count is the amount of people that need to be matched for this policy to be fulfilled.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"filter": {
										Description:         "Filter is a predicate that determines what users count towards this policy.",
										MarkdownDescription: "Filter is a predicate that determines what users count towards this policy.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"kinds": {
										Description:         "Kinds are the session kinds this policy applies to.",
										MarkdownDescription: "Kinds are the session kinds this policy applies to.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"modes": {
										Description:         "Modes is the list of modes that may be used to fulfill this policy.",
										MarkdownDescription: "Modes is the list of modes that may be used to fulfill this policy.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "Name is the name of the policy.",
										MarkdownDescription: "Name is the name of the policy.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"on_leave": {
										Description:         "OnLeave is the behaviour that's used when the policy is no longer fulfilled for a live session.",
										MarkdownDescription: "OnLeave is the behaviour that's used when the policy is no longer fulfilled for a live session.",

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

							"review_requests": {
								Description:         "ReviewRequests defines conditions for submitting access reviews.",
								MarkdownDescription: "ReviewRequests defines conditions for submitting access reviews.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"claims_to_roles": {
										Description:         "ClaimsToRoles specifies a mapping from claims (traits) to teleport roles.",
										MarkdownDescription: "ClaimsToRoles specifies a mapping from claims (traits) to teleport roles.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"claim": {
												Description:         "Claim is a claim name.",
												MarkdownDescription: "Claim is a claim name.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"roles": {
												Description:         "Roles is a list of static teleport roles to match.",
												MarkdownDescription: "Roles is a list of static teleport roles to match.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value": {
												Description:         "Value is a claim value to match.",
												MarkdownDescription: "Value is a claim value to match.",

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

									"preview_as_roles": {
										Description:         "PreviewAsRoles is a list of extra roles which should apply to a reviewer while they are viewing a Resource Access Request for the purposes of viewing details such as the hostname and labels of requested resources.",
										MarkdownDescription: "PreviewAsRoles is a list of extra roles which should apply to a reviewer while they are viewing a Resource Access Request for the purposes of viewing details such as the hostname and labels of requested resources.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"roles": {
										Description:         "Roles is the name of roles which may be reviewed.",
										MarkdownDescription: "Roles is the name of roles which may be reviewed.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"where": {
										Description:         "Where is an optional predicate which further limits which requests are reviewable.",
										MarkdownDescription: "Where is an optional predicate which further limits which requests are reviewable.",

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
								Description:         "Rules is a list of rules and their access levels. Rules are a high level construct used for access control.",
								MarkdownDescription: "Rules is a list of rules and their access levels. Rules are a high level construct used for access control.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"actions": {
										Description:         "Actions specifies optional actions taken when this rule matches",
										MarkdownDescription: "Actions specifies optional actions taken when this rule matches",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"resources": {
										Description:         "Resources is a list of resources",
										MarkdownDescription: "Resources is a list of resources",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"verbs": {
										Description:         "Verbs is a list of verbs",
										MarkdownDescription: "Verbs is a list of verbs",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"where": {
										Description:         "Where specifies optional advanced matcher",
										MarkdownDescription: "Where specifies optional advanced matcher",

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

							"windows_desktop_labels": {
								Description:         "WindowsDesktopLabels are used in the RBAC system to allow/deny access to Windows desktops.",
								MarkdownDescription: "WindowsDesktopLabels are used in the RBAC system to allow/deny access to Windows desktops.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"windows_desktop_logins": {
								Description:         "WindowsDesktopLogins is a list of desktop login names allowed/denied for Windows desktops.",
								MarkdownDescription: "WindowsDesktopLogins is a list of desktop login names allowed/denied for Windows desktops.",

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

					"options": {
						Description:         "Options is for OpenSSH options like agent forwarding.",
						MarkdownDescription: "Options is for OpenSSH options like agent forwarding.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"cert_extensions": {
								Description:         "CertExtensions specifies the key/values",
								MarkdownDescription: "CertExtensions specifies the key/values",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"mode": {
										Description:         "Mode is the type of extension to be used -- currently critical-option is not supported",
										MarkdownDescription: "Mode is the type of extension to be used -- currently critical-option is not supported",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "Name specifies the key to be used in the cert extension.",
										MarkdownDescription: "Name specifies the key to be used in the cert extension.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": {
										Description:         "Type represents the certificate type being extended, only ssh is supported at this time.",
										MarkdownDescription: "Type represents the certificate type being extended, only ssh is supported at this time.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"value": {
										Description:         "Value specifies the valueg to be used in the cert extension.",
										MarkdownDescription: "Value specifies the valueg to be used in the cert extension.",

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

							"cert_format": {
								Description:         "CertificateFormat defines the format of the user certificate to allow compatibility with older versions of OpenSSH.",
								MarkdownDescription: "CertificateFormat defines the format of the user certificate to allow compatibility with older versions of OpenSSH.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"client_idle_timeout": {
								Description:         "ClientIdleTimeout sets disconnect clients on idle timeout behavior, if set to 0 means do not disconnect, otherwise is set to the idle duration.",
								MarkdownDescription: "ClientIdleTimeout sets disconnect clients on idle timeout behavior, if set to 0 means do not disconnect, otherwise is set to the idle duration.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"create_host_user": {
								Description:         "CreateHostUser allows users to be automatically created on a host",
								MarkdownDescription: "CreateHostUser allows users to be automatically created on a host",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"desktop_clipboard": {
								Description:         "DesktopClipboard indicates whether clipboard sharing is allowed between the user's workstation and the remote desktop. It defaults to true unless explicitly set to false.",
								MarkdownDescription: "DesktopClipboard indicates whether clipboard sharing is allowed between the user's workstation and the remote desktop. It defaults to true unless explicitly set to false.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"desktop_directory_sharing": {
								Description:         "DesktopDirectorySharing indicates whether directory sharing is allowed between the user's workstation and the remote desktop. It defaults to false unless explicitly set to true.",
								MarkdownDescription: "DesktopDirectorySharing indicates whether directory sharing is allowed between the user's workstation and the remote desktop. It defaults to false unless explicitly set to true.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"disconnect_expired_cert": {
								Description:         "DisconnectExpiredCert sets disconnect clients on expired certificates.",
								MarkdownDescription: "DisconnectExpiredCert sets disconnect clients on expired certificates.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enhanced_recording": {
								Description:         "BPF defines what events to record for the BPF-based session recorder.",
								MarkdownDescription: "BPF defines what events to record for the BPF-based session recorder.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"forward_agent": {
								Description:         "ForwardAgent is SSH agent forwarding.",
								MarkdownDescription: "ForwardAgent is SSH agent forwarding.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"lock": {
								Description:         "Lock specifies the locking mode (strict|best_effort) to be applied with the role.",
								MarkdownDescription: "Lock specifies the locking mode (strict|best_effort) to be applied with the role.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"max_connections": {
								Description:         "MaxConnections defines the maximum number of concurrent connections a user may hold.",
								MarkdownDescription: "MaxConnections defines the maximum number of concurrent connections a user may hold.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"max_kubernetes_connections": {
								Description:         "MaxKubernetesConnections defines the maximum number of concurrent Kubernetes sessions a user may hold.",
								MarkdownDescription: "MaxKubernetesConnections defines the maximum number of concurrent Kubernetes sessions a user may hold.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"max_session_ttl": {
								Description:         "MaxSessionTTL defines how long a SSH session can last for.",
								MarkdownDescription: "MaxSessionTTL defines how long a SSH session can last for.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"max_sessions": {
								Description:         "MaxSessions defines the maximum number of concurrent sessions per connection.",
								MarkdownDescription: "MaxSessions defines the maximum number of concurrent sessions per connection.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"permit_x11_forwarding": {
								Description:         "PermitX11Forwarding authorizes use of X11 forwarding.",
								MarkdownDescription: "PermitX11Forwarding authorizes use of X11 forwarding.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pin_source_ip": {
								Description:         "PinSourceIP forces the same client IP for certificate generation and usage",
								MarkdownDescription: "PinSourceIP forces the same client IP for certificate generation and usage",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"port_forwarding": {
								Description:         "PortForwarding defines if the certificate will have 'permit-port-forwarding' in the certificate. PortForwarding is 'yes' if not set, that's why this is a pointer",
								MarkdownDescription: "PortForwarding defines if the certificate will have 'permit-port-forwarding' in the certificate. PortForwarding is 'yes' if not set, that's why this is a pointer",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"record_session": {
								Description:         "RecordDesktopSession indicates whether desktop access sessions should be recorded. It defaults to true unless explicitly set to false.",
								MarkdownDescription: "RecordDesktopSession indicates whether desktop access sessions should be recorded. It defaults to true unless explicitly set to false.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"default": {
										Description:         "Default indicates the default value for the services.",
										MarkdownDescription: "Default indicates the default value for the services.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"desktop": {
										Description:         "Desktop indicates whether desktop sessions should be recorded. It defaults to true unless explicitly set to false.",
										MarkdownDescription: "Desktop indicates whether desktop sessions should be recorded. It defaults to true unless explicitly set to false.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"ssh": {
										Description:         "SSH indicates the session mode used on SSH sessions.",
										MarkdownDescription: "SSH indicates the session mode used on SSH sessions.",

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

							"request_access": {
								Description:         "RequestAccess defines the access request stategy (optional|note|always) where optional is the default.",
								MarkdownDescription: "RequestAccess defines the access request stategy (optional|note|always) where optional is the default.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"request_prompt": {
								Description:         "RequestPrompt is an optional message which tells users what they aught to",
								MarkdownDescription: "RequestPrompt is an optional message which tells users what they aught to",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"require_session_mfa": {
								Description:         "RequireMFAType is the type of MFA requirement enforced for this user.",
								MarkdownDescription: "RequireMFAType is the type of MFA requirement enforced for this user.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ssh_file_copy": {
								Description:         "SSHFileCopy indicates whether remote file operations via SCP or SFTP are allowed over an SSH session. It defaults to true unless explicitly set to false.",
								MarkdownDescription: "SSHFileCopy indicates whether remote file operations via SCP or SFTP are allowed over an SSH session. It defaults to true unless explicitly set to false.",

								Type: types.BoolType,

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

func (r *ResourcesTeleportDevTeleportRoleV5Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_resources_teleport_dev_teleport_role_v5")

	var state ResourcesTeleportDevTeleportRoleV5TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ResourcesTeleportDevTeleportRoleV5GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("resources.teleport.dev/v5")
	goModel.Kind = utilities.Ptr("TeleportRole")

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

func (r *ResourcesTeleportDevTeleportRoleV5Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_resources_teleport_dev_teleport_role_v5")
	// NO-OP: All data is already in Terraform state
}

func (r *ResourcesTeleportDevTeleportRoleV5Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_resources_teleport_dev_teleport_role_v5")

	var state ResourcesTeleportDevTeleportRoleV5TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ResourcesTeleportDevTeleportRoleV5GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("resources.teleport.dev/v5")
	goModel.Kind = utilities.Ptr("TeleportRole")

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

func (r *ResourcesTeleportDevTeleportRoleV5Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_resources_teleport_dev_teleport_role_v5")
	// NO-OP: Terraform removes the state automatically for us
}
