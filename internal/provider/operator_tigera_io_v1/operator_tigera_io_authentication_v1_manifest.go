/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package operator_tigera_io_v1

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
	_ datasource.DataSource = &OperatorTigeraIoAuthenticationV1Manifest{}
)

func NewOperatorTigeraIoAuthenticationV1Manifest() datasource.DataSource {
	return &OperatorTigeraIoAuthenticationV1Manifest{}
}

type OperatorTigeraIoAuthenticationV1Manifest struct{}

type OperatorTigeraIoAuthenticationV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		DexDeployment *struct {
			Spec *struct {
				Template *struct {
					Spec *struct {
						Containers *[]struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Resources *struct {
								Claims *[]struct {
									Name *string `tfsdk:"name" json:"name,omitempty"`
								} `tfsdk:"claims" json:"claims,omitempty"`
								Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
								Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
							} `tfsdk:"resources" json:"resources,omitempty"`
						} `tfsdk:"containers" json:"containers,omitempty"`
						InitContainers *[]struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Resources *struct {
								Claims *[]struct {
									Name *string `tfsdk:"name" json:"name,omitempty"`
								} `tfsdk:"claims" json:"claims,omitempty"`
								Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
								Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
							} `tfsdk:"resources" json:"resources,omitempty"`
						} `tfsdk:"init_containers" json:"initContainers,omitempty"`
					} `tfsdk:"spec" json:"spec,omitempty"`
				} `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"dex_deployment" json:"dexDeployment,omitempty"`
		GroupsPrefix *string `tfsdk:"groups_prefix" json:"groupsPrefix,omitempty"`
		Ldap         *struct {
			GroupSearch *struct {
				BaseDN        *string `tfsdk:"base_dn" json:"baseDN,omitempty"`
				Filter        *string `tfsdk:"filter" json:"filter,omitempty"`
				NameAttribute *string `tfsdk:"name_attribute" json:"nameAttribute,omitempty"`
				UserMatchers  *[]struct {
					GroupAttribute *string `tfsdk:"group_attribute" json:"groupAttribute,omitempty"`
					UserAttribute  *string `tfsdk:"user_attribute" json:"userAttribute,omitempty"`
				} `tfsdk:"user_matchers" json:"userMatchers,omitempty"`
			} `tfsdk:"group_search" json:"groupSearch,omitempty"`
			Host       *string `tfsdk:"host" json:"host,omitempty"`
			StartTLS   *bool   `tfsdk:"start_tls" json:"startTLS,omitempty"`
			UserSearch *struct {
				BaseDN        *string `tfsdk:"base_dn" json:"baseDN,omitempty"`
				Filter        *string `tfsdk:"filter" json:"filter,omitempty"`
				NameAttribute *string `tfsdk:"name_attribute" json:"nameAttribute,omitempty"`
			} `tfsdk:"user_search" json:"userSearch,omitempty"`
		} `tfsdk:"ldap" json:"ldap,omitempty"`
		ManagerDomain *string `tfsdk:"manager_domain" json:"managerDomain,omitempty"`
		Oidc          *struct {
			EmailVerification *string   `tfsdk:"email_verification" json:"emailVerification,omitempty"`
			GroupsClaim       *string   `tfsdk:"groups_claim" json:"groupsClaim,omitempty"`
			GroupsPrefix      *string   `tfsdk:"groups_prefix" json:"groupsPrefix,omitempty"`
			IssuerURL         *string   `tfsdk:"issuer_url" json:"issuerURL,omitempty"`
			PromptTypes       *[]string `tfsdk:"prompt_types" json:"promptTypes,omitempty"`
			RequestedScopes   *[]string `tfsdk:"requested_scopes" json:"requestedScopes,omitempty"`
			Type              *string   `tfsdk:"type" json:"type,omitempty"`
			UsernameClaim     *string   `tfsdk:"username_claim" json:"usernameClaim,omitempty"`
			UsernamePrefix    *string   `tfsdk:"username_prefix" json:"usernamePrefix,omitempty"`
		} `tfsdk:"oidc" json:"oidc,omitempty"`
		Openshift *struct {
			IssuerURL *string `tfsdk:"issuer_url" json:"issuerURL,omitempty"`
		} `tfsdk:"openshift" json:"openshift,omitempty"`
		UsernamePrefix *string `tfsdk:"username_prefix" json:"usernamePrefix,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OperatorTigeraIoAuthenticationV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operator_tigera_io_authentication_v1_manifest"
}

func (r *OperatorTigeraIoAuthenticationV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Authentication is the Schema for the authentications API",
		MarkdownDescription: "Authentication is the Schema for the authentications API",
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
				Description:         "AuthenticationSpec defines the desired state of Authentication",
				MarkdownDescription: "AuthenticationSpec defines the desired state of Authentication",
				Attributes: map[string]schema.Attribute{
					"dex_deployment": schema.SingleNestedAttribute{
						Description:         "DexDeployment configures the Dex Deployment.",
						MarkdownDescription: "DexDeployment configures the Dex Deployment.",
						Attributes: map[string]schema.Attribute{
							"spec": schema.SingleNestedAttribute{
								Description:         "Spec is the specification of the Dex Deployment.",
								MarkdownDescription: "Spec is the specification of the Dex Deployment.",
								Attributes: map[string]schema.Attribute{
									"template": schema.SingleNestedAttribute{
										Description:         "Template describes the Dex Deployment pod that will be created.",
										MarkdownDescription: "Template describes the Dex Deployment pod that will be created.",
										Attributes: map[string]schema.Attribute{
											"spec": schema.SingleNestedAttribute{
												Description:         "Spec is the Dex Deployment's PodSpec.",
												MarkdownDescription: "Spec is the Dex Deployment's PodSpec.",
												Attributes: map[string]schema.Attribute{
													"containers": schema.ListNestedAttribute{
														Description:         "Containers is a list of Dex containers. If specified, this overrides the specified Dex Deployment containers. If omitted, the Dex Deployment will use its default values for its containers.",
														MarkdownDescription: "Containers is a list of Dex containers. If specified, this overrides the specified Dex Deployment containers. If omitted, the Dex Deployment will use its default values for its containers.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is an enum which identifies the Dex Deployment container by name. Supported values are: tigera-dex",
																	MarkdownDescription: "Name is an enum which identifies the Dex Deployment container by name. Supported values are: tigera-dex",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("tigera-dex"),
																	},
																},

																"resources": schema.SingleNestedAttribute{
																	Description:         "Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named Dex Deployment container's resources. If omitted, the Dex Deployment will use its default value for this container's resources.",
																	MarkdownDescription: "Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named Dex Deployment container's resources. If omitted, the Dex Deployment will use its default value for this container's resources.",
																	Attributes: map[string]schema.Attribute{
																		"claims": schema.ListNestedAttribute{
																			Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
																			MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																						MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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

																		"limits": schema.MapAttribute{
																			Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"requests": schema.MapAttribute{
																			Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"init_containers": schema.ListNestedAttribute{
														Description:         "InitContainers is a list of Dex init containers. If specified, this overrides the specified Dex Deployment init containers. If omitted, the Dex Deployment will use its default values for its init containers.",
														MarkdownDescription: "InitContainers is a list of Dex init containers. If specified, this overrides the specified Dex Deployment init containers. If omitted, the Dex Deployment will use its default values for its init containers.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is an enum which identifies the Dex Deployment init container by name. Supported values are: tigera-dex-tls-key-cert-provisioner",
																	MarkdownDescription: "Name is an enum which identifies the Dex Deployment init container by name. Supported values are: tigera-dex-tls-key-cert-provisioner",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("tigera-dex-tls-key-cert-provisioner"),
																	},
																},

																"resources": schema.SingleNestedAttribute{
																	Description:         "Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named Dex Deployment init container's resources. If omitted, the Dex Deployment will use its default value for this init container's resources.",
																	MarkdownDescription: "Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named Dex Deployment init container's resources. If omitted, the Dex Deployment will use its default value for this init container's resources.",
																	Attributes: map[string]schema.Attribute{
																		"claims": schema.ListNestedAttribute{
																			Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
																			MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																						MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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

																		"limits": schema.MapAttribute{
																			Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"requests": schema.MapAttribute{
																			Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"groups_prefix": schema.StringAttribute{
						Description:         "If specified, GroupsPrefix is prepended to each group obtained from the identity provider. Note that Kibana does not support a groups prefix, so this prefix is removed from Kubernetes Groups when translating log access ClusterRoleBindings into Elastic.",
						MarkdownDescription: "If specified, GroupsPrefix is prepended to each group obtained from the identity provider. Note that Kibana does not support a groups prefix, so this prefix is removed from Kubernetes Groups when translating log access ClusterRoleBindings into Elastic.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ldap": schema.SingleNestedAttribute{
						Description:         "LDAP contains the configuration needed to setup LDAP authentication.",
						MarkdownDescription: "LDAP contains the configuration needed to setup LDAP authentication.",
						Attributes: map[string]schema.Attribute{
							"group_search": schema.SingleNestedAttribute{
								Description:         "Group search configuration to find the groups that a user is in.",
								MarkdownDescription: "Group search configuration to find the groups that a user is in.",
								Attributes: map[string]schema.Attribute{
									"base_dn": schema.StringAttribute{
										Description:         "BaseDN to start the search from. For example 'cn=groups,dc=example,dc=com'",
										MarkdownDescription: "BaseDN to start the search from. For example 'cn=groups,dc=example,dc=com'",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"filter": schema.StringAttribute{
										Description:         "Optional filter to apply when searching the directory. For example '(objectClass=posixGroup)'",
										MarkdownDescription: "Optional filter to apply when searching the directory. For example '(objectClass=posixGroup)'",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name_attribute": schema.StringAttribute{
										Description:         "The attribute of the group that represents its name. This attribute can be used to apply RBAC to a user group.",
										MarkdownDescription: "The attribute of the group that represents its name. This attribute can be used to apply RBAC to a user group.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"user_matchers": schema.ListNestedAttribute{
										Description:         "Following list contains field pairs that are used to match a user to a group. It adds an additional requirement to the filter that an attribute in the group must match the user's attribute value.",
										MarkdownDescription: "Following list contains field pairs that are used to match a user to a group. It adds an additional requirement to the filter that an attribute in the group must match the user's attribute value.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"group_attribute": schema.StringAttribute{
													Description:         "The attribute of a group that links it to a user.",
													MarkdownDescription: "The attribute of a group that links it to a user.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"user_attribute": schema.StringAttribute{
													Description:         "The attribute of a user that links it to a group.",
													MarkdownDescription: "The attribute of a user that links it to a group.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
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

							"host": schema.StringAttribute{
								Description:         "The host and port of the LDAP server. Example: ad.example.com:636",
								MarkdownDescription: "The host and port of the LDAP server. Example: ad.example.com:636",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"start_tls": schema.BoolAttribute{
								Description:         "StartTLS whether to enable the startTLS feature for establishing TLS on an existing LDAP session. If true, the ldap:// protocol is used and then issues a StartTLS command, otherwise, connections will use the ldaps:// protocol.",
								MarkdownDescription: "StartTLS whether to enable the startTLS feature for establishing TLS on an existing LDAP session. If true, the ldap:// protocol is used and then issues a StartTLS command, otherwise, connections will use the ldaps:// protocol.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"user_search": schema.SingleNestedAttribute{
								Description:         "User entry search configuration to match the credentials with a user.",
								MarkdownDescription: "User entry search configuration to match the credentials with a user.",
								Attributes: map[string]schema.Attribute{
									"base_dn": schema.StringAttribute{
										Description:         "BaseDN to start the search from. For example 'cn=users,dc=example,dc=com'",
										MarkdownDescription: "BaseDN to start the search from. For example 'cn=users,dc=example,dc=com'",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"filter": schema.StringAttribute{
										Description:         "Optional filter to apply when searching the directory. For example '(objectClass=person)'",
										MarkdownDescription: "Optional filter to apply when searching the directory. For example '(objectClass=person)'",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name_attribute": schema.StringAttribute{
										Description:         "A mapping of the attribute that is used as the username. This attribute can be used to apply RBAC to a user. Default: uid",
										MarkdownDescription: "A mapping of the attribute that is used as the username. This attribute can be used to apply RBAC to a user. Default: uid",
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

					"manager_domain": schema.StringAttribute{
						Description:         "ManagerDomain is the domain name of the Manager",
						MarkdownDescription: "ManagerDomain is the domain name of the Manager",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"oidc": schema.SingleNestedAttribute{
						Description:         "OIDC contains the configuration needed to setup OIDC authentication.",
						MarkdownDescription: "OIDC contains the configuration needed to setup OIDC authentication.",
						Attributes: map[string]schema.Attribute{
							"email_verification": schema.StringAttribute{
								Description:         "Some providers do not include the claim 'email_verified' when there is no verification in the user enrollment process or if they are acting as a proxy for another identity provider. By default those tokens are deemed invalid. To skip this check, set the value to 'InsecureSkip'. Default: Verify",
								MarkdownDescription: "Some providers do not include the claim 'email_verified' when there is no verification in the user enrollment process or if they are acting as a proxy for another identity provider. By default those tokens are deemed invalid. To skip this check, set the value to 'InsecureSkip'. Default: Verify",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Verify", "InsecureSkip"),
								},
							},

							"groups_claim": schema.StringAttribute{
								Description:         "GroupsClaim specifies which claim to use from the OIDC provider as the group.",
								MarkdownDescription: "GroupsClaim specifies which claim to use from the OIDC provider as the group.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"groups_prefix": schema.StringAttribute{
								Description:         "Deprecated. Please use Authentication.Spec.GroupsPrefix instead.",
								MarkdownDescription: "Deprecated. Please use Authentication.Spec.GroupsPrefix instead.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"issuer_url": schema.StringAttribute{
								Description:         "IssuerURL is the URL to the OIDC provider.",
								MarkdownDescription: "IssuerURL is the URL to the OIDC provider.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"prompt_types": schema.ListAttribute{
								Description:         "PromptTypes is an optional list of string values that specifies whether the identity provider prompts the end user for re-authentication and consent. See the RFC for more information on prompt types: https://openid.net/specs/openid-connect-core-1_0.html. Default: 'Consent'",
								MarkdownDescription: "PromptTypes is an optional list of string values that specifies whether the identity provider prompts the end user for re-authentication and consent. See the RFC for more information on prompt types: https://openid.net/specs/openid-connect-core-1_0.html. Default: 'Consent'",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"requested_scopes": schema.ListAttribute{
								Description:         "RequestedScopes is a list of scopes to request from the OIDC provider. If not provided, the following scopes are requested: ['openid', 'email', 'profile', 'groups', 'offline_access'].",
								MarkdownDescription: "RequestedScopes is a list of scopes to request from the OIDC provider. If not provided, the following scopes are requested: ['openid', 'email', 'profile', 'groups', 'offline_access'].",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "Default: 'Dex'",
								MarkdownDescription: "Default: 'Dex'",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Dex", "Tigera"),
								},
							},

							"username_claim": schema.StringAttribute{
								Description:         "UsernameClaim specifies which claim to use from the OIDC provider as the username.",
								MarkdownDescription: "UsernameClaim specifies which claim to use from the OIDC provider as the username.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"username_prefix": schema.StringAttribute{
								Description:         "Deprecated. Please use Authentication.Spec.UsernamePrefix instead.",
								MarkdownDescription: "Deprecated. Please use Authentication.Spec.UsernamePrefix instead.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"openshift": schema.SingleNestedAttribute{
						Description:         "Openshift contains the configuration needed to setup Openshift OAuth authentication.",
						MarkdownDescription: "Openshift contains the configuration needed to setup Openshift OAuth authentication.",
						Attributes: map[string]schema.Attribute{
							"issuer_url": schema.StringAttribute{
								Description:         "IssuerURL is the URL to the Openshift OAuth provider. Ex.: https://api.my-ocp-domain.com:6443",
								MarkdownDescription: "IssuerURL is the URL to the Openshift OAuth provider. Ex.: https://api.my-ocp-domain.com:6443",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"username_prefix": schema.StringAttribute{
						Description:         "If specified, UsernamePrefix is prepended to each user obtained from the identity provider. Note that Kibana does not support a user prefix, so this prefix is removed from Kubernetes User when translating log access ClusterRoleBindings into Elastic.",
						MarkdownDescription: "If specified, UsernamePrefix is prepended to each user obtained from the identity provider. Note that Kibana does not support a user prefix, so this prefix is removed from Kubernetes User when translating log access ClusterRoleBindings into Elastic.",
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

func (r *OperatorTigeraIoAuthenticationV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_operator_tigera_io_authentication_v1_manifest")

	var model OperatorTigeraIoAuthenticationV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("operator.tigera.io/v1")
	model.Kind = pointer.String("Authentication")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
