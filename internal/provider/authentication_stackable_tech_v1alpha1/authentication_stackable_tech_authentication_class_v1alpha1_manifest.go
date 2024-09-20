/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package authentication_stackable_tech_v1alpha1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	_ datasource.DataSource = &AuthenticationStackableTechAuthenticationClassV1Alpha1Manifest{}
)

func NewAuthenticationStackableTechAuthenticationClassV1Alpha1Manifest() datasource.DataSource {
	return &AuthenticationStackableTechAuthenticationClassV1Alpha1Manifest{}
}

type AuthenticationStackableTechAuthenticationClassV1Alpha1Manifest struct{}

type AuthenticationStackableTechAuthenticationClassV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Protocol *struct {
			Ldap *struct {
				BindCredentials *struct {
					Scope *struct {
						Node     *bool     `tfsdk:"node" json:"node,omitempty"`
						Pod      *bool     `tfsdk:"pod" json:"pod,omitempty"`
						Services *[]string `tfsdk:"services" json:"services,omitempty"`
					} `tfsdk:"scope" json:"scope,omitempty"`
					SecretClass *string `tfsdk:"secret_class" json:"secretClass,omitempty"`
				} `tfsdk:"bind_credentials" json:"bindCredentials,omitempty"`
				EmailField     *string `tfsdk:"email_field" json:"emailField,omitempty"`
				FirstnameField *string `tfsdk:"firstname_field" json:"firstnameField,omitempty"`
				GroupField     *string `tfsdk:"group_field" json:"groupField,omitempty"`
				Hostname       *string `tfsdk:"hostname" json:"hostname,omitempty"`
				LastnameField  *string `tfsdk:"lastname_field" json:"lastnameField,omitempty"`
				Port           *int64  `tfsdk:"port" json:"port,omitempty"`
				SearchBase     *string `tfsdk:"search_base" json:"searchBase,omitempty"`
				SearchFilter   *string `tfsdk:"search_filter" json:"searchFilter,omitempty"`
				Tls            *struct {
					Insecure           *map[string]string `tfsdk:"insecure" json:"insecure,omitempty"`
					MutualVerification *struct {
						SecretClass *string `tfsdk:"secret_class" json:"secretClass,omitempty"`
					} `tfsdk:"mutual_verification" json:"mutualVerification,omitempty"`
					ServerVerification *struct {
						ServerCaCert *struct {
							Configmap   *string `tfsdk:"configmap" json:"configmap,omitempty"`
							Path        *string `tfsdk:"path" json:"path,omitempty"`
							Secret      *string `tfsdk:"secret" json:"secret,omitempty"`
							SecretClass *string `tfsdk:"secret_class" json:"secretClass,omitempty"`
						} `tfsdk:"server_ca_cert" json:"serverCaCert,omitempty"`
					} `tfsdk:"server_verification" json:"serverVerification,omitempty"`
					SystemProvided *map[string]string `tfsdk:"system_provided" json:"systemProvided,omitempty"`
				} `tfsdk:"tls" json:"tls,omitempty"`
				UidField *string `tfsdk:"uid_field" json:"uidField,omitempty"`
			} `tfsdk:"ldap" json:"ldap,omitempty"`
		} `tfsdk:"protocol" json:"protocol,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AuthenticationStackableTechAuthenticationClassV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_authentication_stackable_tech_authentication_class_v1alpha1_manifest"
}

func (r *AuthenticationStackableTechAuthenticationClassV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Auto-generated derived type for AuthenticationClassSpec via 'CustomResource'",
		MarkdownDescription: "Auto-generated derived type for AuthenticationClassSpec via 'CustomResource'",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"protocol": schema.SingleNestedAttribute{
						Description:         "Protocol used for authentication",
						MarkdownDescription: "Protocol used for authentication",
						Attributes: map[string]schema.Attribute{
							"ldap": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"bind_credentials": schema.SingleNestedAttribute{
										Description:         "In case you need a special account for searching the LDAP server you can specify it here",
										MarkdownDescription: "In case you need a special account for searching the LDAP server you can specify it here",
										Attributes: map[string]schema.Attribute{
											"scope": schema.SingleNestedAttribute{
												Description:         "[Scope](https://docs.stackable.tech/secret-operator/scope.html) of the [SecretClass](https://docs.stackable.tech/secret-operator/secretclass.html)",
												MarkdownDescription: "[Scope](https://docs.stackable.tech/secret-operator/scope.html) of the [SecretClass](https://docs.stackable.tech/secret-operator/secretclass.html)",
												Attributes: map[string]schema.Attribute{
													"node": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"pod": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"services": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
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

											"secret_class": schema.StringAttribute{
												Description:         "[SecretClass](https://docs.stackable.tech/secret-operator/secretclass.html) containing the LDAP bind credentials",
												MarkdownDescription: "[SecretClass](https://docs.stackable.tech/secret-operator/secretclass.html) containing the LDAP bind credentials",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"email_field": schema.StringAttribute{
										Description:         "The name of the email field",
										MarkdownDescription: "The name of the email field",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"firstname_field": schema.StringAttribute{
										Description:         "The name of the firstname field",
										MarkdownDescription: "The name of the firstname field",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"group_field": schema.StringAttribute{
										Description:         "The name of the group field",
										MarkdownDescription: "The name of the group field",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"hostname": schema.StringAttribute{
										Description:         "Hostname of the LDAP server",
										MarkdownDescription: "Hostname of the LDAP server",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"lastname_field": schema.StringAttribute{
										Description:         "The name of the lastname field",
										MarkdownDescription: "The name of the lastname field",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"port": schema.Int64Attribute{
										Description:         "Port of the LDAP server",
										MarkdownDescription: "Port of the LDAP server",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"search_base": schema.StringAttribute{
										Description:         "LDAP search base",
										MarkdownDescription: "LDAP search base",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"search_filter": schema.StringAttribute{
										Description:         "LDAP query to filter users",
										MarkdownDescription: "LDAP query to filter users",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tls": schema.SingleNestedAttribute{
										Description:         "Use a TLS connection. If not specified no TLS will be used",
										MarkdownDescription: "Use a TLS connection. If not specified no TLS will be used",
										Attributes: map[string]schema.Attribute{
											"insecure": schema.MapAttribute{
												Description:         "Use TLS but don't verify certificates. We have to use an empty struct instead of an empty Enum because of limitations of [kube-rs](https://github.com/kube-rs/kube-rs/)",
												MarkdownDescription: "Use TLS but don't verify certificates. We have to use an empty struct instead of an empty Enum because of limitations of [kube-rs](https://github.com/kube-rs/kube-rs/)",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"mutual_verification": schema.SingleNestedAttribute{
												Description:         "Use TLS and ca certificate to verify the server and the client",
												MarkdownDescription: "Use TLS and ca certificate to verify the server and the client",
												Attributes: map[string]schema.Attribute{
													"secret_class": schema.StringAttribute{
														Description:         "[SecretClass](https://docs.stackable.tech/secret-operator/secretclass.html) which will provide ca.crt, tls.crt and tls.key",
														MarkdownDescription: "[SecretClass](https://docs.stackable.tech/secret-operator/secretclass.html) which will provide ca.crt, tls.crt and tls.key",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"server_verification": schema.SingleNestedAttribute{
												Description:         "Use TLS and ca certificate to verify the server",
												MarkdownDescription: "Use TLS and ca certificate to verify the server",
												Attributes: map[string]schema.Attribute{
													"server_ca_cert": schema.SingleNestedAttribute{
														Description:         "Ca cert to verify the server",
														MarkdownDescription: "Ca cert to verify the server",
														Attributes: map[string]schema.Attribute{
															"configmap": schema.StringAttribute{
																Description:         "Name of the ConfigMap containing the ca cert. Key must be 'ca.crt'.",
																MarkdownDescription: "Name of the ConfigMap containing the ca cert. Key must be 'ca.crt'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"path": schema.StringAttribute{
																Description:         "Path to the ca cert",
																MarkdownDescription: "Path to the ca cert",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "Name of the Secret containing the ca cert. Key must be 'ca.crt'.",
																MarkdownDescription: "Name of the Secret containing the ca cert. Key must be 'ca.crt'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_class": schema.StringAttribute{
																Description:         "Name of the SecretClass which will provide the ca cert",
																MarkdownDescription: "Name of the SecretClass which will provide the ca cert",
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

											"system_provided": schema.MapAttribute{
												Description:         "Use TLS and the ca certificates provided by the system - in this case the Docker image - to verify the server. This can be useful when you e.g. use public AWS S3 or other public available services.",
												MarkdownDescription: "Use TLS and the ca certificates provided by the system - in this case the Docker image - to verify the server. This can be useful when you e.g. use public AWS S3 or other public available services.",
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

									"uid_field": schema.StringAttribute{
										Description:         "The name of the username field",
										MarkdownDescription: "The name of the username field",
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
						Required: true,
						Optional: false,
						Computed: false,
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *AuthenticationStackableTechAuthenticationClassV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_authentication_stackable_tech_authentication_class_v1alpha1_manifest")

	var model AuthenticationStackableTechAuthenticationClassV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("authentication.stackable.tech/v1alpha1")
	model.Kind = pointer.String("AuthenticationClass")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
