/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package resources_teleport_dev_v6

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &ResourcesTeleportDevTeleportRoleV6Manifest{}
)

func NewResourcesTeleportDevTeleportRoleV6Manifest() datasource.DataSource {
	return &ResourcesTeleportDevTeleportRoleV6Manifest{}
}

type ResourcesTeleportDevTeleportRoleV6Manifest struct{}

type ResourcesTeleportDevTeleportRoleV6ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Allow *struct {
			App_labels                *map[string]string `tfsdk:"app_labels" json:"app_labels,omitempty"`
			App_labels_expression     *string            `tfsdk:"app_labels_expression" json:"app_labels_expression,omitempty"`
			Aws_role_arns             *[]string          `tfsdk:"aws_role_arns" json:"aws_role_arns,omitempty"`
			Azure_identities          *[]string          `tfsdk:"azure_identities" json:"azure_identities,omitempty"`
			Cluster_labels            *map[string]string `tfsdk:"cluster_labels" json:"cluster_labels,omitempty"`
			Cluster_labels_expression *string            `tfsdk:"cluster_labels_expression" json:"cluster_labels_expression,omitempty"`
			Db_labels                 *map[string]string `tfsdk:"db_labels" json:"db_labels,omitempty"`
			Db_labels_expression      *string            `tfsdk:"db_labels_expression" json:"db_labels_expression,omitempty"`
			Db_names                  *[]string          `tfsdk:"db_names" json:"db_names,omitempty"`
			Db_permissions            *[]struct {
				Match       *map[string]string `tfsdk:"match" json:"match,omitempty"`
				Permissions *[]string          `tfsdk:"permissions" json:"permissions,omitempty"`
			} `tfsdk:"db_permissions" json:"db_permissions,omitempty"`
			Db_roles                     *[]string          `tfsdk:"db_roles" json:"db_roles,omitempty"`
			Db_service_labels            *map[string]string `tfsdk:"db_service_labels" json:"db_service_labels,omitempty"`
			Db_service_labels_expression *string            `tfsdk:"db_service_labels_expression" json:"db_service_labels_expression,omitempty"`
			Db_users                     *[]string          `tfsdk:"db_users" json:"db_users,omitempty"`
			Desktop_groups               *[]string          `tfsdk:"desktop_groups" json:"desktop_groups,omitempty"`
			Gcp_service_accounts         *[]string          `tfsdk:"gcp_service_accounts" json:"gcp_service_accounts,omitempty"`
			Group_labels                 *map[string]string `tfsdk:"group_labels" json:"group_labels,omitempty"`
			Group_labels_expression      *string            `tfsdk:"group_labels_expression" json:"group_labels_expression,omitempty"`
			Host_groups                  *[]string          `tfsdk:"host_groups" json:"host_groups,omitempty"`
			Host_sudoers                 *[]string          `tfsdk:"host_sudoers" json:"host_sudoers,omitempty"`
			Impersonate                  *struct {
				Roles *[]string `tfsdk:"roles" json:"roles,omitempty"`
				Users *[]string `tfsdk:"users" json:"users,omitempty"`
				Where *string   `tfsdk:"where" json:"where,omitempty"`
			} `tfsdk:"impersonate" json:"impersonate,omitempty"`
			Join_sessions *[]struct {
				Kinds *[]string `tfsdk:"kinds" json:"kinds,omitempty"`
				Modes *[]string `tfsdk:"modes" json:"modes,omitempty"`
				Name  *string   `tfsdk:"name" json:"name,omitempty"`
				Roles *[]string `tfsdk:"roles" json:"roles,omitempty"`
			} `tfsdk:"join_sessions" json:"join_sessions,omitempty"`
			Kubernetes_groups            *[]string          `tfsdk:"kubernetes_groups" json:"kubernetes_groups,omitempty"`
			Kubernetes_labels            *map[string]string `tfsdk:"kubernetes_labels" json:"kubernetes_labels,omitempty"`
			Kubernetes_labels_expression *string            `tfsdk:"kubernetes_labels_expression" json:"kubernetes_labels_expression,omitempty"`
			Kubernetes_resources         *[]struct {
				Kind      *string   `tfsdk:"kind" json:"kind,omitempty"`
				Name      *string   `tfsdk:"name" json:"name,omitempty"`
				Namespace *string   `tfsdk:"namespace" json:"namespace,omitempty"`
				Verbs     *[]string `tfsdk:"verbs" json:"verbs,omitempty"`
			} `tfsdk:"kubernetes_resources" json:"kubernetes_resources,omitempty"`
			Kubernetes_users       *[]string          `tfsdk:"kubernetes_users" json:"kubernetes_users,omitempty"`
			Logins                 *[]string          `tfsdk:"logins" json:"logins,omitempty"`
			Node_labels            *map[string]string `tfsdk:"node_labels" json:"node_labels,omitempty"`
			Node_labels_expression *string            `tfsdk:"node_labels_expression" json:"node_labels_expression,omitempty"`
			Request                *struct {
				Annotations     *map[string][]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Claims_to_roles *[]struct {
					Claim *string   `tfsdk:"claim" json:"claim,omitempty"`
					Roles *[]string `tfsdk:"roles" json:"roles,omitempty"`
					Value *string   `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"claims_to_roles" json:"claims_to_roles,omitempty"`
				Max_duration        *string   `tfsdk:"max_duration" json:"max_duration,omitempty"`
				Roles               *[]string `tfsdk:"roles" json:"roles,omitempty"`
				Search_as_roles     *[]string `tfsdk:"search_as_roles" json:"search_as_roles,omitempty"`
				Suggested_reviewers *[]string `tfsdk:"suggested_reviewers" json:"suggested_reviewers,omitempty"`
				Thresholds          *[]struct {
					Approve *int64  `tfsdk:"approve" json:"approve,omitempty"`
					Deny    *int64  `tfsdk:"deny" json:"deny,omitempty"`
					Filter  *string `tfsdk:"filter" json:"filter,omitempty"`
					Name    *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"thresholds" json:"thresholds,omitempty"`
			} `tfsdk:"request" json:"request,omitempty"`
			Require_session_join *[]struct {
				Count    *int64    `tfsdk:"count" json:"count,omitempty"`
				Filter   *string   `tfsdk:"filter" json:"filter,omitempty"`
				Kinds    *[]string `tfsdk:"kinds" json:"kinds,omitempty"`
				Modes    *[]string `tfsdk:"modes" json:"modes,omitempty"`
				Name     *string   `tfsdk:"name" json:"name,omitempty"`
				On_leave *string   `tfsdk:"on_leave" json:"on_leave,omitempty"`
			} `tfsdk:"require_session_join" json:"require_session_join,omitempty"`
			Review_requests *struct {
				Claims_to_roles *[]struct {
					Claim *string   `tfsdk:"claim" json:"claim,omitempty"`
					Roles *[]string `tfsdk:"roles" json:"roles,omitempty"`
					Value *string   `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"claims_to_roles" json:"claims_to_roles,omitempty"`
				Preview_as_roles *[]string `tfsdk:"preview_as_roles" json:"preview_as_roles,omitempty"`
				Roles            *[]string `tfsdk:"roles" json:"roles,omitempty"`
				Where            *string   `tfsdk:"where" json:"where,omitempty"`
			} `tfsdk:"review_requests" json:"review_requests,omitempty"`
			Rules *[]struct {
				Actions   *[]string `tfsdk:"actions" json:"actions,omitempty"`
				Resources *[]string `tfsdk:"resources" json:"resources,omitempty"`
				Verbs     *[]string `tfsdk:"verbs" json:"verbs,omitempty"`
				Where     *string   `tfsdk:"where" json:"where,omitempty"`
			} `tfsdk:"rules" json:"rules,omitempty"`
			Spiffe *[]struct {
				Dns_sans *[]string `tfsdk:"dns_sans" json:"dns_sans,omitempty"`
				Ip_sans  *[]string `tfsdk:"ip_sans" json:"ip_sans,omitempty"`
				Path     *string   `tfsdk:"path" json:"path,omitempty"`
			} `tfsdk:"spiffe" json:"spiffe,omitempty"`
			Windows_desktop_labels            *map[string]string `tfsdk:"windows_desktop_labels" json:"windows_desktop_labels,omitempty"`
			Windows_desktop_labels_expression *string            `tfsdk:"windows_desktop_labels_expression" json:"windows_desktop_labels_expression,omitempty"`
			Windows_desktop_logins            *[]string          `tfsdk:"windows_desktop_logins" json:"windows_desktop_logins,omitempty"`
		} `tfsdk:"allow" json:"allow,omitempty"`
		Deny *struct {
			App_labels                *map[string]string `tfsdk:"app_labels" json:"app_labels,omitempty"`
			App_labels_expression     *string            `tfsdk:"app_labels_expression" json:"app_labels_expression,omitempty"`
			Aws_role_arns             *[]string          `tfsdk:"aws_role_arns" json:"aws_role_arns,omitempty"`
			Azure_identities          *[]string          `tfsdk:"azure_identities" json:"azure_identities,omitempty"`
			Cluster_labels            *map[string]string `tfsdk:"cluster_labels" json:"cluster_labels,omitempty"`
			Cluster_labels_expression *string            `tfsdk:"cluster_labels_expression" json:"cluster_labels_expression,omitempty"`
			Db_labels                 *map[string]string `tfsdk:"db_labels" json:"db_labels,omitempty"`
			Db_labels_expression      *string            `tfsdk:"db_labels_expression" json:"db_labels_expression,omitempty"`
			Db_names                  *[]string          `tfsdk:"db_names" json:"db_names,omitempty"`
			Db_permissions            *[]struct {
				Match       *map[string]string `tfsdk:"match" json:"match,omitempty"`
				Permissions *[]string          `tfsdk:"permissions" json:"permissions,omitempty"`
			} `tfsdk:"db_permissions" json:"db_permissions,omitempty"`
			Db_roles                     *[]string          `tfsdk:"db_roles" json:"db_roles,omitempty"`
			Db_service_labels            *map[string]string `tfsdk:"db_service_labels" json:"db_service_labels,omitempty"`
			Db_service_labels_expression *string            `tfsdk:"db_service_labels_expression" json:"db_service_labels_expression,omitempty"`
			Db_users                     *[]string          `tfsdk:"db_users" json:"db_users,omitempty"`
			Desktop_groups               *[]string          `tfsdk:"desktop_groups" json:"desktop_groups,omitempty"`
			Gcp_service_accounts         *[]string          `tfsdk:"gcp_service_accounts" json:"gcp_service_accounts,omitempty"`
			Group_labels                 *map[string]string `tfsdk:"group_labels" json:"group_labels,omitempty"`
			Group_labels_expression      *string            `tfsdk:"group_labels_expression" json:"group_labels_expression,omitempty"`
			Host_groups                  *[]string          `tfsdk:"host_groups" json:"host_groups,omitempty"`
			Host_sudoers                 *[]string          `tfsdk:"host_sudoers" json:"host_sudoers,omitempty"`
			Impersonate                  *struct {
				Roles *[]string `tfsdk:"roles" json:"roles,omitempty"`
				Users *[]string `tfsdk:"users" json:"users,omitempty"`
				Where *string   `tfsdk:"where" json:"where,omitempty"`
			} `tfsdk:"impersonate" json:"impersonate,omitempty"`
			Join_sessions *[]struct {
				Kinds *[]string `tfsdk:"kinds" json:"kinds,omitempty"`
				Modes *[]string `tfsdk:"modes" json:"modes,omitempty"`
				Name  *string   `tfsdk:"name" json:"name,omitempty"`
				Roles *[]string `tfsdk:"roles" json:"roles,omitempty"`
			} `tfsdk:"join_sessions" json:"join_sessions,omitempty"`
			Kubernetes_groups            *[]string          `tfsdk:"kubernetes_groups" json:"kubernetes_groups,omitempty"`
			Kubernetes_labels            *map[string]string `tfsdk:"kubernetes_labels" json:"kubernetes_labels,omitempty"`
			Kubernetes_labels_expression *string            `tfsdk:"kubernetes_labels_expression" json:"kubernetes_labels_expression,omitempty"`
			Kubernetes_resources         *[]struct {
				Kind      *string   `tfsdk:"kind" json:"kind,omitempty"`
				Name      *string   `tfsdk:"name" json:"name,omitempty"`
				Namespace *string   `tfsdk:"namespace" json:"namespace,omitempty"`
				Verbs     *[]string `tfsdk:"verbs" json:"verbs,omitempty"`
			} `tfsdk:"kubernetes_resources" json:"kubernetes_resources,omitempty"`
			Kubernetes_users       *[]string          `tfsdk:"kubernetes_users" json:"kubernetes_users,omitempty"`
			Logins                 *[]string          `tfsdk:"logins" json:"logins,omitempty"`
			Node_labels            *map[string]string `tfsdk:"node_labels" json:"node_labels,omitempty"`
			Node_labels_expression *string            `tfsdk:"node_labels_expression" json:"node_labels_expression,omitempty"`
			Request                *struct {
				Annotations     *map[string][]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Claims_to_roles *[]struct {
					Claim *string   `tfsdk:"claim" json:"claim,omitempty"`
					Roles *[]string `tfsdk:"roles" json:"roles,omitempty"`
					Value *string   `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"claims_to_roles" json:"claims_to_roles,omitempty"`
				Max_duration        *string   `tfsdk:"max_duration" json:"max_duration,omitempty"`
				Roles               *[]string `tfsdk:"roles" json:"roles,omitempty"`
				Search_as_roles     *[]string `tfsdk:"search_as_roles" json:"search_as_roles,omitempty"`
				Suggested_reviewers *[]string `tfsdk:"suggested_reviewers" json:"suggested_reviewers,omitempty"`
				Thresholds          *[]struct {
					Approve *int64  `tfsdk:"approve" json:"approve,omitempty"`
					Deny    *int64  `tfsdk:"deny" json:"deny,omitempty"`
					Filter  *string `tfsdk:"filter" json:"filter,omitempty"`
					Name    *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"thresholds" json:"thresholds,omitempty"`
			} `tfsdk:"request" json:"request,omitempty"`
			Require_session_join *[]struct {
				Count    *int64    `tfsdk:"count" json:"count,omitempty"`
				Filter   *string   `tfsdk:"filter" json:"filter,omitempty"`
				Kinds    *[]string `tfsdk:"kinds" json:"kinds,omitempty"`
				Modes    *[]string `tfsdk:"modes" json:"modes,omitempty"`
				Name     *string   `tfsdk:"name" json:"name,omitempty"`
				On_leave *string   `tfsdk:"on_leave" json:"on_leave,omitempty"`
			} `tfsdk:"require_session_join" json:"require_session_join,omitempty"`
			Review_requests *struct {
				Claims_to_roles *[]struct {
					Claim *string   `tfsdk:"claim" json:"claim,omitempty"`
					Roles *[]string `tfsdk:"roles" json:"roles,omitempty"`
					Value *string   `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"claims_to_roles" json:"claims_to_roles,omitempty"`
				Preview_as_roles *[]string `tfsdk:"preview_as_roles" json:"preview_as_roles,omitempty"`
				Roles            *[]string `tfsdk:"roles" json:"roles,omitempty"`
				Where            *string   `tfsdk:"where" json:"where,omitempty"`
			} `tfsdk:"review_requests" json:"review_requests,omitempty"`
			Rules *[]struct {
				Actions   *[]string `tfsdk:"actions" json:"actions,omitempty"`
				Resources *[]string `tfsdk:"resources" json:"resources,omitempty"`
				Verbs     *[]string `tfsdk:"verbs" json:"verbs,omitempty"`
				Where     *string   `tfsdk:"where" json:"where,omitempty"`
			} `tfsdk:"rules" json:"rules,omitempty"`
			Spiffe *[]struct {
				Dns_sans *[]string `tfsdk:"dns_sans" json:"dns_sans,omitempty"`
				Ip_sans  *[]string `tfsdk:"ip_sans" json:"ip_sans,omitempty"`
				Path     *string   `tfsdk:"path" json:"path,omitempty"`
			} `tfsdk:"spiffe" json:"spiffe,omitempty"`
			Windows_desktop_labels            *map[string]string `tfsdk:"windows_desktop_labels" json:"windows_desktop_labels,omitempty"`
			Windows_desktop_labels_expression *string            `tfsdk:"windows_desktop_labels_expression" json:"windows_desktop_labels_expression,omitempty"`
			Windows_desktop_logins            *[]string          `tfsdk:"windows_desktop_logins" json:"windows_desktop_logins,omitempty"`
		} `tfsdk:"deny" json:"deny,omitempty"`
		Options *struct {
			Cert_extensions *[]struct {
				Mode  *string `tfsdk:"mode" json:"mode,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"cert_extensions" json:"cert_extensions,omitempty"`
			Cert_format                    *string   `tfsdk:"cert_format" json:"cert_format,omitempty"`
			Client_idle_timeout            *string   `tfsdk:"client_idle_timeout" json:"client_idle_timeout,omitempty"`
			Create_db_user                 *bool     `tfsdk:"create_db_user" json:"create_db_user,omitempty"`
			Create_db_user_mode            *string   `tfsdk:"create_db_user_mode" json:"create_db_user_mode,omitempty"`
			Create_desktop_user            *bool     `tfsdk:"create_desktop_user" json:"create_desktop_user,omitempty"`
			Create_host_user               *bool     `tfsdk:"create_host_user" json:"create_host_user,omitempty"`
			Create_host_user_default_shell *string   `tfsdk:"create_host_user_default_shell" json:"create_host_user_default_shell,omitempty"`
			Create_host_user_mode          *string   `tfsdk:"create_host_user_mode" json:"create_host_user_mode,omitempty"`
			Desktop_clipboard              *bool     `tfsdk:"desktop_clipboard" json:"desktop_clipboard,omitempty"`
			Desktop_directory_sharing      *bool     `tfsdk:"desktop_directory_sharing" json:"desktop_directory_sharing,omitempty"`
			Device_trust_mode              *string   `tfsdk:"device_trust_mode" json:"device_trust_mode,omitempty"`
			Disconnect_expired_cert        *bool     `tfsdk:"disconnect_expired_cert" json:"disconnect_expired_cert,omitempty"`
			Enhanced_recording             *[]string `tfsdk:"enhanced_recording" json:"enhanced_recording,omitempty"`
			Forward_agent                  *bool     `tfsdk:"forward_agent" json:"forward_agent,omitempty"`
			Idp                            *struct {
				Saml *struct {
					Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
				} `tfsdk:"saml" json:"saml,omitempty"`
			} `tfsdk:"idp" json:"idp,omitempty"`
			Lock                       *string `tfsdk:"lock" json:"lock,omitempty"`
			Max_connections            *int64  `tfsdk:"max_connections" json:"max_connections,omitempty"`
			Max_kubernetes_connections *int64  `tfsdk:"max_kubernetes_connections" json:"max_kubernetes_connections,omitempty"`
			Max_session_ttl            *string `tfsdk:"max_session_ttl" json:"max_session_ttl,omitempty"`
			Max_sessions               *int64  `tfsdk:"max_sessions" json:"max_sessions,omitempty"`
			Mfa_verification_interval  *string `tfsdk:"mfa_verification_interval" json:"mfa_verification_interval,omitempty"`
			Permit_x11_forwarding      *bool   `tfsdk:"permit_x11_forwarding" json:"permit_x11_forwarding,omitempty"`
			Pin_source_ip              *bool   `tfsdk:"pin_source_ip" json:"pin_source_ip,omitempty"`
			Port_forwarding            *bool   `tfsdk:"port_forwarding" json:"port_forwarding,omitempty"`
			Record_session             *struct {
				Default *string `tfsdk:"default" json:"default,omitempty"`
				Desktop *bool   `tfsdk:"desktop" json:"desktop,omitempty"`
				Ssh     *string `tfsdk:"ssh" json:"ssh,omitempty"`
			} `tfsdk:"record_session" json:"record_session,omitempty"`
			Request_access      *string `tfsdk:"request_access" json:"request_access,omitempty"`
			Request_prompt      *string `tfsdk:"request_prompt" json:"request_prompt,omitempty"`
			Require_session_mfa *string `tfsdk:"require_session_mfa" json:"require_session_mfa,omitempty"`
			Ssh_file_copy       *bool   `tfsdk:"ssh_file_copy" json:"ssh_file_copy,omitempty"`
		} `tfsdk:"options" json:"options,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ResourcesTeleportDevTeleportRoleV6Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_resources_teleport_dev_teleport_role_v6_manifest"
}

func (r *ResourcesTeleportDevTeleportRoleV6Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Role is the Schema for the roles API",
		MarkdownDescription: "Role is the Schema for the roles API",
		Attributes: map[string]schema.Attribute{
			"yaml": schema.StringAttribute{
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
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
						Optional:            true,
						Computed:            false,
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
						Computed:            false,
						Validators: []validator.Map{
							validators.AnnotationValidator(),
						},
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "Role resource definition v6 from Teleport",
				MarkdownDescription: "Role resource definition v6 from Teleport",
				Attributes: map[string]schema.Attribute{
					"allow": schema.SingleNestedAttribute{
						Description:         "Allow is the set of conditions evaluated to grant access.",
						MarkdownDescription: "Allow is the set of conditions evaluated to grant access.",
						Attributes: map[string]schema.Attribute{
							"app_labels": schema.MapAttribute{
								Description:         "AppLabels is a map of labels used as part of the RBAC system.",
								MarkdownDescription: "AppLabels is a map of labels used as part of the RBAC system.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"app_labels_expression": schema.StringAttribute{
								Description:         "AppLabelsExpression is a predicate expression used to allow/deny access to Apps.",
								MarkdownDescription: "AppLabelsExpression is a predicate expression used to allow/deny access to Apps.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"aws_role_arns": schema.ListAttribute{
								Description:         "AWSRoleARNs is a list of AWS role ARNs this role is allowed to assume.",
								MarkdownDescription: "AWSRoleARNs is a list of AWS role ARNs this role is allowed to assume.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"azure_identities": schema.ListAttribute{
								Description:         "AzureIdentities is a list of Azure identities this role is allowed to assume.",
								MarkdownDescription: "AzureIdentities is a list of Azure identities this role is allowed to assume.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cluster_labels": schema.MapAttribute{
								Description:         "ClusterLabels is a map of node labels (used to dynamically grant access to clusters).",
								MarkdownDescription: "ClusterLabels is a map of node labels (used to dynamically grant access to clusters).",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cluster_labels_expression": schema.StringAttribute{
								Description:         "ClusterLabelsExpression is a predicate expression used to allow/deny access to remote Teleport clusters.",
								MarkdownDescription: "ClusterLabelsExpression is a predicate expression used to allow/deny access to remote Teleport clusters.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"db_labels": schema.MapAttribute{
								Description:         "DatabaseLabels are used in RBAC system to allow/deny access to databases.",
								MarkdownDescription: "DatabaseLabels are used in RBAC system to allow/deny access to databases.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"db_labels_expression": schema.StringAttribute{
								Description:         "DatabaseLabelsExpression is a predicate expression used to allow/deny access to Databases.",
								MarkdownDescription: "DatabaseLabelsExpression is a predicate expression used to allow/deny access to Databases.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"db_names": schema.ListAttribute{
								Description:         "DatabaseNames is a list of database names this role is allowed to connect to.",
								MarkdownDescription: "DatabaseNames is a list of database names this role is allowed to connect to.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"db_permissions": schema.ListNestedAttribute{
								Description:         "DatabasePermissions specifies a set of permissions that will be granted to the database user when using automatic database user provisioning.",
								MarkdownDescription: "DatabasePermissions specifies a set of permissions that will be granted to the database user when using automatic database user provisioning.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"match": schema.MapAttribute{
											Description:         "Match is a list of object labels that must be matched for the permission to be granted.",
											MarkdownDescription: "Match is a list of object labels that must be matched for the permission to be granted.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"permissions": schema.ListAttribute{
											Description:         "Permission is the list of string representations of the permission to be given, e.g. SELECT, INSERT, UPDATE, ...",
											MarkdownDescription: "Permission is the list of string representations of the permission to be given, e.g. SELECT, INSERT, UPDATE, ...",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"db_roles": schema.ListAttribute{
								Description:         "DatabaseRoles is a list of databases roles for automatic user creation.",
								MarkdownDescription: "DatabaseRoles is a list of databases roles for automatic user creation.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"db_service_labels": schema.MapAttribute{
								Description:         "DatabaseServiceLabels are used in RBAC system to allow/deny access to Database Services.",
								MarkdownDescription: "DatabaseServiceLabels are used in RBAC system to allow/deny access to Database Services.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"db_service_labels_expression": schema.StringAttribute{
								Description:         "DatabaseServiceLabelsExpression is a predicate expression used to allow/deny access to Database Services.",
								MarkdownDescription: "DatabaseServiceLabelsExpression is a predicate expression used to allow/deny access to Database Services.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"db_users": schema.ListAttribute{
								Description:         "DatabaseUsers is a list of databases users this role is allowed to connect as.",
								MarkdownDescription: "DatabaseUsers is a list of databases users this role is allowed to connect as.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"desktop_groups": schema.ListAttribute{
								Description:         "DesktopGroups is a list of groups for created desktop users to be added to",
								MarkdownDescription: "DesktopGroups is a list of groups for created desktop users to be added to",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"gcp_service_accounts": schema.ListAttribute{
								Description:         "GCPServiceAccounts is a list of GCP service accounts this role is allowed to assume.",
								MarkdownDescription: "GCPServiceAccounts is a list of GCP service accounts this role is allowed to assume.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"group_labels": schema.MapAttribute{
								Description:         "GroupLabels is a map of labels used as part of the RBAC system.",
								MarkdownDescription: "GroupLabels is a map of labels used as part of the RBAC system.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"group_labels_expression": schema.StringAttribute{
								Description:         "GroupLabelsExpression is a predicate expression used to allow/deny access to user groups.",
								MarkdownDescription: "GroupLabelsExpression is a predicate expression used to allow/deny access to user groups.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"host_groups": schema.ListAttribute{
								Description:         "HostGroups is a list of groups for created users to be added to",
								MarkdownDescription: "HostGroups is a list of groups for created users to be added to",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"host_sudoers": schema.ListAttribute{
								Description:         "HostSudoers is a list of entries to include in a users sudoer file",
								MarkdownDescription: "HostSudoers is a list of entries to include in a users sudoer file",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"impersonate": schema.SingleNestedAttribute{
								Description:         "Impersonate specifies what users and roles this role is allowed to impersonate by issuing certificates or other possible means.",
								MarkdownDescription: "Impersonate specifies what users and roles this role is allowed to impersonate by issuing certificates or other possible means.",
								Attributes: map[string]schema.Attribute{
									"roles": schema.ListAttribute{
										Description:         "Roles is a list of resources this role is allowed to impersonate",
										MarkdownDescription: "Roles is a list of resources this role is allowed to impersonate",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"users": schema.ListAttribute{
										Description:         "Users is a list of resources this role is allowed to impersonate, could be an empty list or a Wildcard pattern",
										MarkdownDescription: "Users is a list of resources this role is allowed to impersonate, could be an empty list or a Wildcard pattern",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"where": schema.StringAttribute{
										Description:         "Where specifies optional advanced matcher",
										MarkdownDescription: "Where specifies optional advanced matcher",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"join_sessions": schema.ListNestedAttribute{
								Description:         "JoinSessions specifies policies to allow users to join other sessions.",
								MarkdownDescription: "JoinSessions specifies policies to allow users to join other sessions.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"kinds": schema.ListAttribute{
											Description:         "Kinds are the session kinds this policy applies to.",
											MarkdownDescription: "Kinds are the session kinds this policy applies to.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"modes": schema.ListAttribute{
											Description:         "Modes is a list of permitted participant modes for this policy.",
											MarkdownDescription: "Modes is a list of permitted participant modes for this policy.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name is the name of the policy.",
											MarkdownDescription: "Name is the name of the policy.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"roles": schema.ListAttribute{
											Description:         "Roles is a list of roles that you can join the session of.",
											MarkdownDescription: "Roles is a list of roles that you can join the session of.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"kubernetes_groups": schema.ListAttribute{
								Description:         "KubeGroups is a list of kubernetes groups",
								MarkdownDescription: "KubeGroups is a list of kubernetes groups",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kubernetes_labels": schema.MapAttribute{
								Description:         "KubernetesLabels is a map of kubernetes cluster labels used for RBAC.",
								MarkdownDescription: "KubernetesLabels is a map of kubernetes cluster labels used for RBAC.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kubernetes_labels_expression": schema.StringAttribute{
								Description:         "KubernetesLabelsExpression is a predicate expression used to allow/deny access to kubernetes clusters.",
								MarkdownDescription: "KubernetesLabelsExpression is a predicate expression used to allow/deny access to kubernetes clusters.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kubernetes_resources": schema.ListNestedAttribute{
								Description:         "KubernetesResources is the Kubernetes Resources this Role grants access to.",
								MarkdownDescription: "KubernetesResources is the Kubernetes Resources this Role grants access to.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"kind": schema.StringAttribute{
											Description:         "Kind specifies the Kubernetes Resource type. At the moment only 'pod' is supported.",
											MarkdownDescription: "Kind specifies the Kubernetes Resource type. At the moment only 'pod' is supported.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name is the resource name. It supports wildcards.",
											MarkdownDescription: "Name is the resource name. It supports wildcards.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"namespace": schema.StringAttribute{
											Description:         "Namespace is the resource namespace. It supports wildcards.",
											MarkdownDescription: "Namespace is the resource namespace. It supports wildcards.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"verbs": schema.ListAttribute{
											Description:         "Verbs are the allowed Kubernetes verbs for the following resource.",
											MarkdownDescription: "Verbs are the allowed Kubernetes verbs for the following resource.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"kubernetes_users": schema.ListAttribute{
								Description:         "KubeUsers is an optional kubernetes users to impersonate",
								MarkdownDescription: "KubeUsers is an optional kubernetes users to impersonate",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"logins": schema.ListAttribute{
								Description:         "Logins is a list of *nix system logins.",
								MarkdownDescription: "Logins is a list of *nix system logins.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_labels": schema.MapAttribute{
								Description:         "NodeLabels is a map of node labels (used to dynamically grant access to nodes).",
								MarkdownDescription: "NodeLabels is a map of node labels (used to dynamically grant access to nodes).",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_labels_expression": schema.StringAttribute{
								Description:         "NodeLabelsExpression is a predicate expression used to allow/deny access to SSH nodes.",
								MarkdownDescription: "NodeLabelsExpression is a predicate expression used to allow/deny access to SSH nodes.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"request": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations is a collection of annotations to be programmatically appended to pending Access Requests at the time of their creation. These annotations serve as a mechanism to propagate extra information to plugins. Since these annotations support variable interpolation syntax, they also offer a mechanism for forwarding claims from an external identity provider, to a plugin via '{{external.trait_name}}' style substitutions.",
										MarkdownDescription: "Annotations is a collection of annotations to be programmatically appended to pending Access Requests at the time of their creation. These annotations serve as a mechanism to propagate extra information to plugins. Since these annotations support variable interpolation syntax, they also offer a mechanism for forwarding claims from an external identity provider, to a plugin via '{{external.trait_name}}' style substitutions.",
										ElementType:         types.ListType{ElemType: types.StringType},
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"claims_to_roles": schema.ListNestedAttribute{
										Description:         "ClaimsToRoles specifies a mapping from claims (traits) to teleport roles.",
										MarkdownDescription: "ClaimsToRoles specifies a mapping from claims (traits) to teleport roles.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"claim": schema.StringAttribute{
													Description:         "Claim is a claim name.",
													MarkdownDescription: "Claim is a claim name.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"roles": schema.ListAttribute{
													Description:         "Roles is a list of static teleport roles to match.",
													MarkdownDescription: "Roles is a list of static teleport roles to match.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "Value is a claim value to match.",
													MarkdownDescription: "Value is a claim value to match.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"max_duration": schema.StringAttribute{
										Description:         "MaxDuration is the amount of time the access will be granted for. If this is zero, the default duration is used.",
										MarkdownDescription: "MaxDuration is the amount of time the access will be granted for. If this is zero, the default duration is used.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"roles": schema.ListAttribute{
										Description:         "Roles is the name of roles which will match the request rule.",
										MarkdownDescription: "Roles is the name of roles which will match the request rule.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"search_as_roles": schema.ListAttribute{
										Description:         "SearchAsRoles is a list of extra roles which should apply to a user while they are searching for resources as part of a Resource Access Request, and defines the underlying roles which will be requested as part of any Resource Access Request.",
										MarkdownDescription: "SearchAsRoles is a list of extra roles which should apply to a user while they are searching for resources as part of a Resource Access Request, and defines the underlying roles which will be requested as part of any Resource Access Request.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"suggested_reviewers": schema.ListAttribute{
										Description:         "SuggestedReviewers is a list of reviewer suggestions. These can be teleport usernames, but that is not a requirement.",
										MarkdownDescription: "SuggestedReviewers is a list of reviewer suggestions. These can be teleport usernames, but that is not a requirement.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"thresholds": schema.ListNestedAttribute{
										Description:         "Thresholds is a list of thresholds, one of which must be met in order for reviews to trigger a state-transition. If no thresholds are provided, a default threshold of 1 for approval and denial is used.",
										MarkdownDescription: "Thresholds is a list of thresholds, one of which must be met in order for reviews to trigger a state-transition. If no thresholds are provided, a default threshold of 1 for approval and denial is used.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"approve": schema.Int64Attribute{
													Description:         "Approve is the number of matching approvals needed for state-transition.",
													MarkdownDescription: "Approve is the number of matching approvals needed for state-transition.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"deny": schema.Int64Attribute{
													Description:         "Deny is the number of denials needed for state-transition.",
													MarkdownDescription: "Deny is the number of denials needed for state-transition.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"filter": schema.StringAttribute{
													Description:         "Filter is an optional predicate used to determine which reviews count toward this threshold.",
													MarkdownDescription: "Filter is an optional predicate used to determine which reviews count toward this threshold.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name is the optional human-readable name of the threshold.",
													MarkdownDescription: "Name is the optional human-readable name of the threshold.",
													Required:            false,
													Optional:            true,
													Computed:            false,
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

							"require_session_join": schema.ListNestedAttribute{
								Description:         "RequireSessionJoin specifies policies for required users to start a session.",
								MarkdownDescription: "RequireSessionJoin specifies policies for required users to start a session.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"count": schema.Int64Attribute{
											Description:         "Count is the amount of people that need to be matched for this policy to be fulfilled.",
											MarkdownDescription: "Count is the amount of people that need to be matched for this policy to be fulfilled.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"filter": schema.StringAttribute{
											Description:         "Filter is a predicate that determines what users count towards this policy.",
											MarkdownDescription: "Filter is a predicate that determines what users count towards this policy.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"kinds": schema.ListAttribute{
											Description:         "Kinds are the session kinds this policy applies to.",
											MarkdownDescription: "Kinds are the session kinds this policy applies to.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"modes": schema.ListAttribute{
											Description:         "Modes is the list of modes that may be used to fulfill this policy.",
											MarkdownDescription: "Modes is the list of modes that may be used to fulfill this policy.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name is the name of the policy.",
											MarkdownDescription: "Name is the name of the policy.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"on_leave": schema.StringAttribute{
											Description:         "OnLeave is the behaviour that's used when the policy is no longer fulfilled for a live session.",
											MarkdownDescription: "OnLeave is the behaviour that's used when the policy is no longer fulfilled for a live session.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"review_requests": schema.SingleNestedAttribute{
								Description:         "ReviewRequests defines conditions for submitting access reviews.",
								MarkdownDescription: "ReviewRequests defines conditions for submitting access reviews.",
								Attributes: map[string]schema.Attribute{
									"claims_to_roles": schema.ListNestedAttribute{
										Description:         "ClaimsToRoles specifies a mapping from claims (traits) to teleport roles.",
										MarkdownDescription: "ClaimsToRoles specifies a mapping from claims (traits) to teleport roles.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"claim": schema.StringAttribute{
													Description:         "Claim is a claim name.",
													MarkdownDescription: "Claim is a claim name.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"roles": schema.ListAttribute{
													Description:         "Roles is a list of static teleport roles to match.",
													MarkdownDescription: "Roles is a list of static teleport roles to match.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "Value is a claim value to match.",
													MarkdownDescription: "Value is a claim value to match.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"preview_as_roles": schema.ListAttribute{
										Description:         "PreviewAsRoles is a list of extra roles which should apply to a reviewer while they are viewing a Resource Access Request for the purposes of viewing details such as the hostname and labels of requested resources.",
										MarkdownDescription: "PreviewAsRoles is a list of extra roles which should apply to a reviewer while they are viewing a Resource Access Request for the purposes of viewing details such as the hostname and labels of requested resources.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"roles": schema.ListAttribute{
										Description:         "Roles is the name of roles which may be reviewed.",
										MarkdownDescription: "Roles is the name of roles which may be reviewed.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"where": schema.StringAttribute{
										Description:         "Where is an optional predicate which further limits which requests are reviewable.",
										MarkdownDescription: "Where is an optional predicate which further limits which requests are reviewable.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"rules": schema.ListNestedAttribute{
								Description:         "Rules is a list of rules and their access levels. Rules are a high level construct used for access control.",
								MarkdownDescription: "Rules is a list of rules and their access levels. Rules are a high level construct used for access control.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"actions": schema.ListAttribute{
											Description:         "Actions specifies optional actions taken when this rule matches",
											MarkdownDescription: "Actions specifies optional actions taken when this rule matches",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"resources": schema.ListAttribute{
											Description:         "Resources is a list of resources",
											MarkdownDescription: "Resources is a list of resources",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"verbs": schema.ListAttribute{
											Description:         "Verbs is a list of verbs",
											MarkdownDescription: "Verbs is a list of verbs",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"where": schema.StringAttribute{
											Description:         "Where specifies optional advanced matcher",
											MarkdownDescription: "Where specifies optional advanced matcher",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"spiffe": schema.ListNestedAttribute{
								Description:         "SPIFFE is used to allow or deny access to a role holder to generating a SPIFFE SVID.",
								MarkdownDescription: "SPIFFE is used to allow or deny access to a role holder to generating a SPIFFE SVID.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"dns_sans": schema.ListAttribute{
											Description:         "DNSSANs specifies matchers for the SPIFFE ID DNS SANs. Each requested DNS SAN is compared against all matchers configured and if any match, the condition is considered to be met. The matcher by default allows '*' to be used to indicate zero or more of any character. Prepend '^' and append '$' to instead switch to matching using the Go regex syntax. Example: *.example.com would match foo.example.com",
											MarkdownDescription: "DNSSANs specifies matchers for the SPIFFE ID DNS SANs. Each requested DNS SAN is compared against all matchers configured and if any match, the condition is considered to be met. The matcher by default allows '*' to be used to indicate zero or more of any character. Prepend '^' and append '$' to instead switch to matching using the Go regex syntax. Example: *.example.com would match foo.example.com",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ip_sans": schema.ListAttribute{
											Description:         "IPSANs specifies matchers for the SPIFFE ID IP SANs. Each requested IP SAN is compared against all matchers configured and if any match, the condition is considered to be met. The matchers should be specified using CIDR notation, it supports IPv4 and IPv6. Examples: - 10.0.0.0/24 would match 10.0.0.0 to 10.255.255.255 - 10.0.0.42/32 would match only 10.0.0.42",
											MarkdownDescription: "IPSANs specifies matchers for the SPIFFE ID IP SANs. Each requested IP SAN is compared against all matchers configured and if any match, the condition is considered to be met. The matchers should be specified using CIDR notation, it supports IPv4 and IPv6. Examples: - 10.0.0.0/24 would match 10.0.0.0 to 10.255.255.255 - 10.0.0.42/32 would match only 10.0.0.42",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"path": schema.StringAttribute{
											Description:         "Path specifies a matcher for the SPIFFE ID path. It should not include the trust domain and should start with a leading slash. The matcher by default allows '*' to be used to indicate zero or more of any character. Prepend '^' and append '$' to instead switch to matching using the Go regex syntax. Example: - /svc/foo/*/bar would match /svc/foo/baz/bar - ^/svc/foo/.*/bar$ would match /svc/foo/baz/bar",
											MarkdownDescription: "Path specifies a matcher for the SPIFFE ID path. It should not include the trust domain and should start with a leading slash. The matcher by default allows '*' to be used to indicate zero or more of any character. Prepend '^' and append '$' to instead switch to matching using the Go regex syntax. Example: - /svc/foo/*/bar would match /svc/foo/baz/bar - ^/svc/foo/.*/bar$ would match /svc/foo/baz/bar",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"windows_desktop_labels": schema.MapAttribute{
								Description:         "WindowsDesktopLabels are used in the RBAC system to allow/deny access to Windows desktops.",
								MarkdownDescription: "WindowsDesktopLabels are used in the RBAC system to allow/deny access to Windows desktops.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"windows_desktop_labels_expression": schema.StringAttribute{
								Description:         "WindowsDesktopLabelsExpression is a predicate expression used to allow/deny access to Windows desktops.",
								MarkdownDescription: "WindowsDesktopLabelsExpression is a predicate expression used to allow/deny access to Windows desktops.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"windows_desktop_logins": schema.ListAttribute{
								Description:         "WindowsDesktopLogins is a list of desktop login names allowed/denied for Windows desktops.",
								MarkdownDescription: "WindowsDesktopLogins is a list of desktop login names allowed/denied for Windows desktops.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"deny": schema.SingleNestedAttribute{
						Description:         "Deny is the set of conditions evaluated to deny access. Deny takes priority over allow.",
						MarkdownDescription: "Deny is the set of conditions evaluated to deny access. Deny takes priority over allow.",
						Attributes: map[string]schema.Attribute{
							"app_labels": schema.MapAttribute{
								Description:         "AppLabels is a map of labels used as part of the RBAC system.",
								MarkdownDescription: "AppLabels is a map of labels used as part of the RBAC system.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"app_labels_expression": schema.StringAttribute{
								Description:         "AppLabelsExpression is a predicate expression used to allow/deny access to Apps.",
								MarkdownDescription: "AppLabelsExpression is a predicate expression used to allow/deny access to Apps.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"aws_role_arns": schema.ListAttribute{
								Description:         "AWSRoleARNs is a list of AWS role ARNs this role is allowed to assume.",
								MarkdownDescription: "AWSRoleARNs is a list of AWS role ARNs this role is allowed to assume.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"azure_identities": schema.ListAttribute{
								Description:         "AzureIdentities is a list of Azure identities this role is allowed to assume.",
								MarkdownDescription: "AzureIdentities is a list of Azure identities this role is allowed to assume.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cluster_labels": schema.MapAttribute{
								Description:         "ClusterLabels is a map of node labels (used to dynamically grant access to clusters).",
								MarkdownDescription: "ClusterLabels is a map of node labels (used to dynamically grant access to clusters).",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cluster_labels_expression": schema.StringAttribute{
								Description:         "ClusterLabelsExpression is a predicate expression used to allow/deny access to remote Teleport clusters.",
								MarkdownDescription: "ClusterLabelsExpression is a predicate expression used to allow/deny access to remote Teleport clusters.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"db_labels": schema.MapAttribute{
								Description:         "DatabaseLabels are used in RBAC system to allow/deny access to databases.",
								MarkdownDescription: "DatabaseLabels are used in RBAC system to allow/deny access to databases.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"db_labels_expression": schema.StringAttribute{
								Description:         "DatabaseLabelsExpression is a predicate expression used to allow/deny access to Databases.",
								MarkdownDescription: "DatabaseLabelsExpression is a predicate expression used to allow/deny access to Databases.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"db_names": schema.ListAttribute{
								Description:         "DatabaseNames is a list of database names this role is allowed to connect to.",
								MarkdownDescription: "DatabaseNames is a list of database names this role is allowed to connect to.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"db_permissions": schema.ListNestedAttribute{
								Description:         "DatabasePermissions specifies a set of permissions that will be granted to the database user when using automatic database user provisioning.",
								MarkdownDescription: "DatabasePermissions specifies a set of permissions that will be granted to the database user when using automatic database user provisioning.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"match": schema.MapAttribute{
											Description:         "Match is a list of object labels that must be matched for the permission to be granted.",
											MarkdownDescription: "Match is a list of object labels that must be matched for the permission to be granted.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"permissions": schema.ListAttribute{
											Description:         "Permission is the list of string representations of the permission to be given, e.g. SELECT, INSERT, UPDATE, ...",
											MarkdownDescription: "Permission is the list of string representations of the permission to be given, e.g. SELECT, INSERT, UPDATE, ...",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"db_roles": schema.ListAttribute{
								Description:         "DatabaseRoles is a list of databases roles for automatic user creation.",
								MarkdownDescription: "DatabaseRoles is a list of databases roles for automatic user creation.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"db_service_labels": schema.MapAttribute{
								Description:         "DatabaseServiceLabels are used in RBAC system to allow/deny access to Database Services.",
								MarkdownDescription: "DatabaseServiceLabels are used in RBAC system to allow/deny access to Database Services.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"db_service_labels_expression": schema.StringAttribute{
								Description:         "DatabaseServiceLabelsExpression is a predicate expression used to allow/deny access to Database Services.",
								MarkdownDescription: "DatabaseServiceLabelsExpression is a predicate expression used to allow/deny access to Database Services.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"db_users": schema.ListAttribute{
								Description:         "DatabaseUsers is a list of databases users this role is allowed to connect as.",
								MarkdownDescription: "DatabaseUsers is a list of databases users this role is allowed to connect as.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"desktop_groups": schema.ListAttribute{
								Description:         "DesktopGroups is a list of groups for created desktop users to be added to",
								MarkdownDescription: "DesktopGroups is a list of groups for created desktop users to be added to",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"gcp_service_accounts": schema.ListAttribute{
								Description:         "GCPServiceAccounts is a list of GCP service accounts this role is allowed to assume.",
								MarkdownDescription: "GCPServiceAccounts is a list of GCP service accounts this role is allowed to assume.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"group_labels": schema.MapAttribute{
								Description:         "GroupLabels is a map of labels used as part of the RBAC system.",
								MarkdownDescription: "GroupLabels is a map of labels used as part of the RBAC system.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"group_labels_expression": schema.StringAttribute{
								Description:         "GroupLabelsExpression is a predicate expression used to allow/deny access to user groups.",
								MarkdownDescription: "GroupLabelsExpression is a predicate expression used to allow/deny access to user groups.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"host_groups": schema.ListAttribute{
								Description:         "HostGroups is a list of groups for created users to be added to",
								MarkdownDescription: "HostGroups is a list of groups for created users to be added to",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"host_sudoers": schema.ListAttribute{
								Description:         "HostSudoers is a list of entries to include in a users sudoer file",
								MarkdownDescription: "HostSudoers is a list of entries to include in a users sudoer file",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"impersonate": schema.SingleNestedAttribute{
								Description:         "Impersonate specifies what users and roles this role is allowed to impersonate by issuing certificates or other possible means.",
								MarkdownDescription: "Impersonate specifies what users and roles this role is allowed to impersonate by issuing certificates or other possible means.",
								Attributes: map[string]schema.Attribute{
									"roles": schema.ListAttribute{
										Description:         "Roles is a list of resources this role is allowed to impersonate",
										MarkdownDescription: "Roles is a list of resources this role is allowed to impersonate",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"users": schema.ListAttribute{
										Description:         "Users is a list of resources this role is allowed to impersonate, could be an empty list or a Wildcard pattern",
										MarkdownDescription: "Users is a list of resources this role is allowed to impersonate, could be an empty list or a Wildcard pattern",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"where": schema.StringAttribute{
										Description:         "Where specifies optional advanced matcher",
										MarkdownDescription: "Where specifies optional advanced matcher",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"join_sessions": schema.ListNestedAttribute{
								Description:         "JoinSessions specifies policies to allow users to join other sessions.",
								MarkdownDescription: "JoinSessions specifies policies to allow users to join other sessions.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"kinds": schema.ListAttribute{
											Description:         "Kinds are the session kinds this policy applies to.",
											MarkdownDescription: "Kinds are the session kinds this policy applies to.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"modes": schema.ListAttribute{
											Description:         "Modes is a list of permitted participant modes for this policy.",
											MarkdownDescription: "Modes is a list of permitted participant modes for this policy.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name is the name of the policy.",
											MarkdownDescription: "Name is the name of the policy.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"roles": schema.ListAttribute{
											Description:         "Roles is a list of roles that you can join the session of.",
											MarkdownDescription: "Roles is a list of roles that you can join the session of.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"kubernetes_groups": schema.ListAttribute{
								Description:         "KubeGroups is a list of kubernetes groups",
								MarkdownDescription: "KubeGroups is a list of kubernetes groups",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kubernetes_labels": schema.MapAttribute{
								Description:         "KubernetesLabels is a map of kubernetes cluster labels used for RBAC.",
								MarkdownDescription: "KubernetesLabels is a map of kubernetes cluster labels used for RBAC.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kubernetes_labels_expression": schema.StringAttribute{
								Description:         "KubernetesLabelsExpression is a predicate expression used to allow/deny access to kubernetes clusters.",
								MarkdownDescription: "KubernetesLabelsExpression is a predicate expression used to allow/deny access to kubernetes clusters.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kubernetes_resources": schema.ListNestedAttribute{
								Description:         "KubernetesResources is the Kubernetes Resources this Role grants access to.",
								MarkdownDescription: "KubernetesResources is the Kubernetes Resources this Role grants access to.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"kind": schema.StringAttribute{
											Description:         "Kind specifies the Kubernetes Resource type. At the moment only 'pod' is supported.",
											MarkdownDescription: "Kind specifies the Kubernetes Resource type. At the moment only 'pod' is supported.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name is the resource name. It supports wildcards.",
											MarkdownDescription: "Name is the resource name. It supports wildcards.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"namespace": schema.StringAttribute{
											Description:         "Namespace is the resource namespace. It supports wildcards.",
											MarkdownDescription: "Namespace is the resource namespace. It supports wildcards.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"verbs": schema.ListAttribute{
											Description:         "Verbs are the allowed Kubernetes verbs for the following resource.",
											MarkdownDescription: "Verbs are the allowed Kubernetes verbs for the following resource.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"kubernetes_users": schema.ListAttribute{
								Description:         "KubeUsers is an optional kubernetes users to impersonate",
								MarkdownDescription: "KubeUsers is an optional kubernetes users to impersonate",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"logins": schema.ListAttribute{
								Description:         "Logins is a list of *nix system logins.",
								MarkdownDescription: "Logins is a list of *nix system logins.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_labels": schema.MapAttribute{
								Description:         "NodeLabels is a map of node labels (used to dynamically grant access to nodes).",
								MarkdownDescription: "NodeLabels is a map of node labels (used to dynamically grant access to nodes).",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_labels_expression": schema.StringAttribute{
								Description:         "NodeLabelsExpression is a predicate expression used to allow/deny access to SSH nodes.",
								MarkdownDescription: "NodeLabelsExpression is a predicate expression used to allow/deny access to SSH nodes.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"request": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations is a collection of annotations to be programmatically appended to pending Access Requests at the time of their creation. These annotations serve as a mechanism to propagate extra information to plugins. Since these annotations support variable interpolation syntax, they also offer a mechanism for forwarding claims from an external identity provider, to a plugin via '{{external.trait_name}}' style substitutions.",
										MarkdownDescription: "Annotations is a collection of annotations to be programmatically appended to pending Access Requests at the time of their creation. These annotations serve as a mechanism to propagate extra information to plugins. Since these annotations support variable interpolation syntax, they also offer a mechanism for forwarding claims from an external identity provider, to a plugin via '{{external.trait_name}}' style substitutions.",
										ElementType:         types.ListType{ElemType: types.StringType},
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"claims_to_roles": schema.ListNestedAttribute{
										Description:         "ClaimsToRoles specifies a mapping from claims (traits) to teleport roles.",
										MarkdownDescription: "ClaimsToRoles specifies a mapping from claims (traits) to teleport roles.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"claim": schema.StringAttribute{
													Description:         "Claim is a claim name.",
													MarkdownDescription: "Claim is a claim name.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"roles": schema.ListAttribute{
													Description:         "Roles is a list of static teleport roles to match.",
													MarkdownDescription: "Roles is a list of static teleport roles to match.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "Value is a claim value to match.",
													MarkdownDescription: "Value is a claim value to match.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"max_duration": schema.StringAttribute{
										Description:         "MaxDuration is the amount of time the access will be granted for. If this is zero, the default duration is used.",
										MarkdownDescription: "MaxDuration is the amount of time the access will be granted for. If this is zero, the default duration is used.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"roles": schema.ListAttribute{
										Description:         "Roles is the name of roles which will match the request rule.",
										MarkdownDescription: "Roles is the name of roles which will match the request rule.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"search_as_roles": schema.ListAttribute{
										Description:         "SearchAsRoles is a list of extra roles which should apply to a user while they are searching for resources as part of a Resource Access Request, and defines the underlying roles which will be requested as part of any Resource Access Request.",
										MarkdownDescription: "SearchAsRoles is a list of extra roles which should apply to a user while they are searching for resources as part of a Resource Access Request, and defines the underlying roles which will be requested as part of any Resource Access Request.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"suggested_reviewers": schema.ListAttribute{
										Description:         "SuggestedReviewers is a list of reviewer suggestions. These can be teleport usernames, but that is not a requirement.",
										MarkdownDescription: "SuggestedReviewers is a list of reviewer suggestions. These can be teleport usernames, but that is not a requirement.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"thresholds": schema.ListNestedAttribute{
										Description:         "Thresholds is a list of thresholds, one of which must be met in order for reviews to trigger a state-transition. If no thresholds are provided, a default threshold of 1 for approval and denial is used.",
										MarkdownDescription: "Thresholds is a list of thresholds, one of which must be met in order for reviews to trigger a state-transition. If no thresholds are provided, a default threshold of 1 for approval and denial is used.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"approve": schema.Int64Attribute{
													Description:         "Approve is the number of matching approvals needed for state-transition.",
													MarkdownDescription: "Approve is the number of matching approvals needed for state-transition.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"deny": schema.Int64Attribute{
													Description:         "Deny is the number of denials needed for state-transition.",
													MarkdownDescription: "Deny is the number of denials needed for state-transition.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"filter": schema.StringAttribute{
													Description:         "Filter is an optional predicate used to determine which reviews count toward this threshold.",
													MarkdownDescription: "Filter is an optional predicate used to determine which reviews count toward this threshold.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name is the optional human-readable name of the threshold.",
													MarkdownDescription: "Name is the optional human-readable name of the threshold.",
													Required:            false,
													Optional:            true,
													Computed:            false,
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

							"require_session_join": schema.ListNestedAttribute{
								Description:         "RequireSessionJoin specifies policies for required users to start a session.",
								MarkdownDescription: "RequireSessionJoin specifies policies for required users to start a session.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"count": schema.Int64Attribute{
											Description:         "Count is the amount of people that need to be matched for this policy to be fulfilled.",
											MarkdownDescription: "Count is the amount of people that need to be matched for this policy to be fulfilled.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"filter": schema.StringAttribute{
											Description:         "Filter is a predicate that determines what users count towards this policy.",
											MarkdownDescription: "Filter is a predicate that determines what users count towards this policy.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"kinds": schema.ListAttribute{
											Description:         "Kinds are the session kinds this policy applies to.",
											MarkdownDescription: "Kinds are the session kinds this policy applies to.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"modes": schema.ListAttribute{
											Description:         "Modes is the list of modes that may be used to fulfill this policy.",
											MarkdownDescription: "Modes is the list of modes that may be used to fulfill this policy.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name is the name of the policy.",
											MarkdownDescription: "Name is the name of the policy.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"on_leave": schema.StringAttribute{
											Description:         "OnLeave is the behaviour that's used when the policy is no longer fulfilled for a live session.",
											MarkdownDescription: "OnLeave is the behaviour that's used when the policy is no longer fulfilled for a live session.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"review_requests": schema.SingleNestedAttribute{
								Description:         "ReviewRequests defines conditions for submitting access reviews.",
								MarkdownDescription: "ReviewRequests defines conditions for submitting access reviews.",
								Attributes: map[string]schema.Attribute{
									"claims_to_roles": schema.ListNestedAttribute{
										Description:         "ClaimsToRoles specifies a mapping from claims (traits) to teleport roles.",
										MarkdownDescription: "ClaimsToRoles specifies a mapping from claims (traits) to teleport roles.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"claim": schema.StringAttribute{
													Description:         "Claim is a claim name.",
													MarkdownDescription: "Claim is a claim name.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"roles": schema.ListAttribute{
													Description:         "Roles is a list of static teleport roles to match.",
													MarkdownDescription: "Roles is a list of static teleport roles to match.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "Value is a claim value to match.",
													MarkdownDescription: "Value is a claim value to match.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"preview_as_roles": schema.ListAttribute{
										Description:         "PreviewAsRoles is a list of extra roles which should apply to a reviewer while they are viewing a Resource Access Request for the purposes of viewing details such as the hostname and labels of requested resources.",
										MarkdownDescription: "PreviewAsRoles is a list of extra roles which should apply to a reviewer while they are viewing a Resource Access Request for the purposes of viewing details such as the hostname and labels of requested resources.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"roles": schema.ListAttribute{
										Description:         "Roles is the name of roles which may be reviewed.",
										MarkdownDescription: "Roles is the name of roles which may be reviewed.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"where": schema.StringAttribute{
										Description:         "Where is an optional predicate which further limits which requests are reviewable.",
										MarkdownDescription: "Where is an optional predicate which further limits which requests are reviewable.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"rules": schema.ListNestedAttribute{
								Description:         "Rules is a list of rules and their access levels. Rules are a high level construct used for access control.",
								MarkdownDescription: "Rules is a list of rules and their access levels. Rules are a high level construct used for access control.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"actions": schema.ListAttribute{
											Description:         "Actions specifies optional actions taken when this rule matches",
											MarkdownDescription: "Actions specifies optional actions taken when this rule matches",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"resources": schema.ListAttribute{
											Description:         "Resources is a list of resources",
											MarkdownDescription: "Resources is a list of resources",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"verbs": schema.ListAttribute{
											Description:         "Verbs is a list of verbs",
											MarkdownDescription: "Verbs is a list of verbs",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"where": schema.StringAttribute{
											Description:         "Where specifies optional advanced matcher",
											MarkdownDescription: "Where specifies optional advanced matcher",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"spiffe": schema.ListNestedAttribute{
								Description:         "SPIFFE is used to allow or deny access to a role holder to generating a SPIFFE SVID.",
								MarkdownDescription: "SPIFFE is used to allow or deny access to a role holder to generating a SPIFFE SVID.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"dns_sans": schema.ListAttribute{
											Description:         "DNSSANs specifies matchers for the SPIFFE ID DNS SANs. Each requested DNS SAN is compared against all matchers configured and if any match, the condition is considered to be met. The matcher by default allows '*' to be used to indicate zero or more of any character. Prepend '^' and append '$' to instead switch to matching using the Go regex syntax. Example: *.example.com would match foo.example.com",
											MarkdownDescription: "DNSSANs specifies matchers for the SPIFFE ID DNS SANs. Each requested DNS SAN is compared against all matchers configured and if any match, the condition is considered to be met. The matcher by default allows '*' to be used to indicate zero or more of any character. Prepend '^' and append '$' to instead switch to matching using the Go regex syntax. Example: *.example.com would match foo.example.com",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ip_sans": schema.ListAttribute{
											Description:         "IPSANs specifies matchers for the SPIFFE ID IP SANs. Each requested IP SAN is compared against all matchers configured and if any match, the condition is considered to be met. The matchers should be specified using CIDR notation, it supports IPv4 and IPv6. Examples: - 10.0.0.0/24 would match 10.0.0.0 to 10.255.255.255 - 10.0.0.42/32 would match only 10.0.0.42",
											MarkdownDescription: "IPSANs specifies matchers for the SPIFFE ID IP SANs. Each requested IP SAN is compared against all matchers configured and if any match, the condition is considered to be met. The matchers should be specified using CIDR notation, it supports IPv4 and IPv6. Examples: - 10.0.0.0/24 would match 10.0.0.0 to 10.255.255.255 - 10.0.0.42/32 would match only 10.0.0.42",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"path": schema.StringAttribute{
											Description:         "Path specifies a matcher for the SPIFFE ID path. It should not include the trust domain and should start with a leading slash. The matcher by default allows '*' to be used to indicate zero or more of any character. Prepend '^' and append '$' to instead switch to matching using the Go regex syntax. Example: - /svc/foo/*/bar would match /svc/foo/baz/bar - ^/svc/foo/.*/bar$ would match /svc/foo/baz/bar",
											MarkdownDescription: "Path specifies a matcher for the SPIFFE ID path. It should not include the trust domain and should start with a leading slash. The matcher by default allows '*' to be used to indicate zero or more of any character. Prepend '^' and append '$' to instead switch to matching using the Go regex syntax. Example: - /svc/foo/*/bar would match /svc/foo/baz/bar - ^/svc/foo/.*/bar$ would match /svc/foo/baz/bar",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"windows_desktop_labels": schema.MapAttribute{
								Description:         "WindowsDesktopLabels are used in the RBAC system to allow/deny access to Windows desktops.",
								MarkdownDescription: "WindowsDesktopLabels are used in the RBAC system to allow/deny access to Windows desktops.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"windows_desktop_labels_expression": schema.StringAttribute{
								Description:         "WindowsDesktopLabelsExpression is a predicate expression used to allow/deny access to Windows desktops.",
								MarkdownDescription: "WindowsDesktopLabelsExpression is a predicate expression used to allow/deny access to Windows desktops.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"windows_desktop_logins": schema.ListAttribute{
								Description:         "WindowsDesktopLogins is a list of desktop login names allowed/denied for Windows desktops.",
								MarkdownDescription: "WindowsDesktopLogins is a list of desktop login names allowed/denied for Windows desktops.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"options": schema.SingleNestedAttribute{
						Description:         "Options is for OpenSSH options like agent forwarding.",
						MarkdownDescription: "Options is for OpenSSH options like agent forwarding.",
						Attributes: map[string]schema.Attribute{
							"cert_extensions": schema.ListNestedAttribute{
								Description:         "CertExtensions specifies the key/values",
								MarkdownDescription: "CertExtensions specifies the key/values",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"mode": schema.StringAttribute{
											Description:         "Mode is the type of extension to be used -- currently critical-option is not supported. 0 is 'extension'.",
											MarkdownDescription: "Mode is the type of extension to be used -- currently critical-option is not supported. 0 is 'extension'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name specifies the key to be used in the cert extension.",
											MarkdownDescription: "Name specifies the key to be used in the cert extension.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"type": schema.StringAttribute{
											Description:         "Type represents the certificate type being extended, only ssh is supported at this time. 0 is 'ssh'.",
											MarkdownDescription: "Type represents the certificate type being extended, only ssh is supported at this time. 0 is 'ssh'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Value specifies the value to be used in the cert extension.",
											MarkdownDescription: "Value specifies the value to be used in the cert extension.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"cert_format": schema.StringAttribute{
								Description:         "CertificateFormat defines the format of the user certificate to allow compatibility with older versions of OpenSSH.",
								MarkdownDescription: "CertificateFormat defines the format of the user certificate to allow compatibility with older versions of OpenSSH.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"client_idle_timeout": schema.StringAttribute{
								Description:         "ClientIdleTimeout sets disconnect clients on idle timeout behavior, if set to 0 means do not disconnect, otherwise is set to the idle duration.",
								MarkdownDescription: "ClientIdleTimeout sets disconnect clients on idle timeout behavior, if set to 0 means do not disconnect, otherwise is set to the idle duration.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"create_db_user": schema.BoolAttribute{
								Description:         "CreateDatabaseUser enabled automatic database user creation.",
								MarkdownDescription: "CreateDatabaseUser enabled automatic database user creation.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"create_db_user_mode": schema.StringAttribute{
								Description:         "CreateDatabaseUserMode allows users to be automatically created on a database when not set to off. 0 is 'unspecified', 1 is 'off', 2 is 'keep', 3 is 'best_effort_drop'.",
								MarkdownDescription: "CreateDatabaseUserMode allows users to be automatically created on a database when not set to off. 0 is 'unspecified', 1 is 'off', 2 is 'keep', 3 is 'best_effort_drop'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"create_desktop_user": schema.BoolAttribute{
								Description:         "CreateDesktopUser allows users to be automatically created on a Windows desktop",
								MarkdownDescription: "CreateDesktopUser allows users to be automatically created on a Windows desktop",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"create_host_user": schema.BoolAttribute{
								Description:         "CreateHostUser allows users to be automatically created on a host",
								MarkdownDescription: "CreateHostUser allows users to be automatically created on a host",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"create_host_user_default_shell": schema.StringAttribute{
								Description:         "CreateHostUserDefaultShell is used to configure the default shell for newly provisioned host users.",
								MarkdownDescription: "CreateHostUserDefaultShell is used to configure the default shell for newly provisioned host users.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"create_host_user_mode": schema.StringAttribute{
								Description:         "CreateHostUserMode allows users to be automatically created on a host when not set to off. 0 is 'unspecified'; 1 is 'off'; 2 is 'drop' (removed for v15 and above), 3 is 'keep'; 4 is 'insecure-drop'.",
								MarkdownDescription: "CreateHostUserMode allows users to be automatically created on a host when not set to off. 0 is 'unspecified'; 1 is 'off'; 2 is 'drop' (removed for v15 and above), 3 is 'keep'; 4 is 'insecure-drop'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"desktop_clipboard": schema.BoolAttribute{
								Description:         "DesktopClipboard indicates whether clipboard sharing is allowed between the user's workstation and the remote desktop. It defaults to true unless explicitly set to false.",
								MarkdownDescription: "DesktopClipboard indicates whether clipboard sharing is allowed between the user's workstation and the remote desktop. It defaults to true unless explicitly set to false.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"desktop_directory_sharing": schema.BoolAttribute{
								Description:         "DesktopDirectorySharing indicates whether directory sharing is allowed between the user's workstation and the remote desktop. It defaults to false unless explicitly set to true.",
								MarkdownDescription: "DesktopDirectorySharing indicates whether directory sharing is allowed between the user's workstation and the remote desktop. It defaults to false unless explicitly set to true.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"device_trust_mode": schema.StringAttribute{
								Description:         "DeviceTrustMode is the device authorization mode used for the resources associated with the role. See DeviceTrust.Mode.",
								MarkdownDescription: "DeviceTrustMode is the device authorization mode used for the resources associated with the role. See DeviceTrust.Mode.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"disconnect_expired_cert": schema.BoolAttribute{
								Description:         "DisconnectExpiredCert sets disconnect clients on expired certificates.",
								MarkdownDescription: "DisconnectExpiredCert sets disconnect clients on expired certificates.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enhanced_recording": schema.ListAttribute{
								Description:         "BPF defines what events to record for the BPF-based session recorder.",
								MarkdownDescription: "BPF defines what events to record for the BPF-based session recorder.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"forward_agent": schema.BoolAttribute{
								Description:         "ForwardAgent is SSH agent forwarding.",
								MarkdownDescription: "ForwardAgent is SSH agent forwarding.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"idp": schema.SingleNestedAttribute{
								Description:         "IDP is a set of options related to accessing IdPs within Teleport. Requires Teleport Enterprise.",
								MarkdownDescription: "IDP is a set of options related to accessing IdPs within Teleport. Requires Teleport Enterprise.",
								Attributes: map[string]schema.Attribute{
									"saml": schema.SingleNestedAttribute{
										Description:         "SAML are options related to the Teleport SAML IdP.",
										MarkdownDescription: "SAML are options related to the Teleport SAML IdP.",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Enabled is set to true if this option allows access to the Teleport SAML IdP.",
												MarkdownDescription: "Enabled is set to true if this option allows access to the Teleport SAML IdP.",
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

							"lock": schema.StringAttribute{
								Description:         "Lock specifies the locking mode (strict|best_effort) to be applied with the role.",
								MarkdownDescription: "Lock specifies the locking mode (strict|best_effort) to be applied with the role.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_connections": schema.Int64Attribute{
								Description:         "MaxConnections defines the maximum number of concurrent connections a user may hold.",
								MarkdownDescription: "MaxConnections defines the maximum number of concurrent connections a user may hold.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_kubernetes_connections": schema.Int64Attribute{
								Description:         "MaxKubernetesConnections defines the maximum number of concurrent Kubernetes sessions a user may hold.",
								MarkdownDescription: "MaxKubernetesConnections defines the maximum number of concurrent Kubernetes sessions a user may hold.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_session_ttl": schema.StringAttribute{
								Description:         "MaxSessionTTL defines how long a SSH session can last for.",
								MarkdownDescription: "MaxSessionTTL defines how long a SSH session can last for.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_sessions": schema.Int64Attribute{
								Description:         "MaxSessions defines the maximum number of concurrent sessions per connection.",
								MarkdownDescription: "MaxSessions defines the maximum number of concurrent sessions per connection.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"mfa_verification_interval": schema.StringAttribute{
								Description:         "MFAVerificationInterval optionally defines the maximum duration that can elapse between successive MFA verifications. This variable is used to ensure that users are periodically prompted to verify their identity, enhancing security by preventing prolonged sessions without re-authentication when using tsh proxy * derivatives. It's only effective if the session requires MFA. If not set, defaults to 'max_session_ttl'.",
								MarkdownDescription: "MFAVerificationInterval optionally defines the maximum duration that can elapse between successive MFA verifications. This variable is used to ensure that users are periodically prompted to verify their identity, enhancing security by preventing prolonged sessions without re-authentication when using tsh proxy * derivatives. It's only effective if the session requires MFA. If not set, defaults to 'max_session_ttl'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"permit_x11_forwarding": schema.BoolAttribute{
								Description:         "PermitX11Forwarding authorizes use of X11 forwarding.",
								MarkdownDescription: "PermitX11Forwarding authorizes use of X11 forwarding.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pin_source_ip": schema.BoolAttribute{
								Description:         "PinSourceIP forces the same client IP for certificate generation and usage",
								MarkdownDescription: "PinSourceIP forces the same client IP for certificate generation and usage",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port_forwarding": schema.BoolAttribute{
								Description:         "PortForwarding defines if the certificate will have 'permit-port-forwarding' in the certificate. PortForwarding is 'yes' if not set, that's why this is a pointer",
								MarkdownDescription: "PortForwarding defines if the certificate will have 'permit-port-forwarding' in the certificate. PortForwarding is 'yes' if not set, that's why this is a pointer",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"record_session": schema.SingleNestedAttribute{
								Description:         "RecordDesktopSession indicates whether desktop access sessions should be recorded. It defaults to true unless explicitly set to false.",
								MarkdownDescription: "RecordDesktopSession indicates whether desktop access sessions should be recorded. It defaults to true unless explicitly set to false.",
								Attributes: map[string]schema.Attribute{
									"default": schema.StringAttribute{
										Description:         "Default indicates the default value for the services.",
										MarkdownDescription: "Default indicates the default value for the services.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"desktop": schema.BoolAttribute{
										Description:         "Desktop indicates whether desktop sessions should be recorded. It defaults to true unless explicitly set to false.",
										MarkdownDescription: "Desktop indicates whether desktop sessions should be recorded. It defaults to true unless explicitly set to false.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ssh": schema.StringAttribute{
										Description:         "SSH indicates the session mode used on SSH sessions.",
										MarkdownDescription: "SSH indicates the session mode used on SSH sessions.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"request_access": schema.StringAttribute{
								Description:         "RequestAccess defines the request strategy (optional|note|always) where optional is the default.",
								MarkdownDescription: "RequestAccess defines the request strategy (optional|note|always) where optional is the default.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"request_prompt": schema.StringAttribute{
								Description:         "RequestPrompt is an optional message which tells users what they aught to request.",
								MarkdownDescription: "RequestPrompt is an optional message which tells users what they aught to request.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"require_session_mfa": schema.StringAttribute{
								Description:         "RequireMFAType is the type of MFA requirement enforced for this user. 0 is 'OFF', 1 is 'SESSION', 2 is 'SESSION_AND_HARDWARE_KEY', 3 is 'HARDWARE_KEY_TOUCH', 4 is 'HARDWARE_KEY_PIN', 5 is 'HARDWARE_KEY_TOUCH_AND_PIN'.",
								MarkdownDescription: "RequireMFAType is the type of MFA requirement enforced for this user. 0 is 'OFF', 1 is 'SESSION', 2 is 'SESSION_AND_HARDWARE_KEY', 3 is 'HARDWARE_KEY_TOUCH', 4 is 'HARDWARE_KEY_PIN', 5 is 'HARDWARE_KEY_TOUCH_AND_PIN'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ssh_file_copy": schema.BoolAttribute{
								Description:         "SSHFileCopy indicates whether remote file operations via SCP or SFTP are allowed over an SSH session. It defaults to true unless explicitly set to false.",
								MarkdownDescription: "SSHFileCopy indicates whether remote file operations via SCP or SFTP are allowed over an SSH session. It defaults to true unless explicitly set to false.",
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
	}
}

func (r *ResourcesTeleportDevTeleportRoleV6Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_resources_teleport_dev_teleport_role_v6_manifest")

	var model ResourcesTeleportDevTeleportRoleV6ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("resources.teleport.dev/v6")
	model.Kind = pointer.String("TeleportRole")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
