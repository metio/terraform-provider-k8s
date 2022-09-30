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

type KeycloakOrgKeycloakUserV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*KeycloakOrgKeycloakUserV1Alpha1Resource)(nil)
)

type KeycloakOrgKeycloakUserV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type KeycloakOrgKeycloakUserV1Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		RealmSelector *struct {
			MatchExpressions *[]struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

				Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
			} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

			MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
		} `tfsdk:"realm_selector" yaml:"realmSelector,omitempty"`

		User *struct {
			ClientRoles *map[string][]string `tfsdk:"client_roles" yaml:"clientRoles,omitempty"`

			Email *string `tfsdk:"email" yaml:"email,omitempty"`

			Groups *[]string `tfsdk:"groups" yaml:"groups,omitempty"`

			RealmRoles *[]string `tfsdk:"realm_roles" yaml:"realmRoles,omitempty"`

			Attributes *map[string][]string `tfsdk:"attributes" yaml:"attributes,omitempty"`

			Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

			FirstName *string `tfsdk:"first_name" yaml:"firstName,omitempty"`

			Credentials *[]struct {
				Temporary *bool `tfsdk:"temporary" yaml:"temporary,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`
			} `tfsdk:"credentials" yaml:"credentials,omitempty"`

			EmailVerified *bool `tfsdk:"email_verified" yaml:"emailVerified,omitempty"`

			FederatedIdentities *[]struct {
				IdentityProvider *string `tfsdk:"identity_provider" yaml:"identityProvider,omitempty"`

				UserId *string `tfsdk:"user_id" yaml:"userId,omitempty"`

				UserName *string `tfsdk:"user_name" yaml:"userName,omitempty"`
			} `tfsdk:"federated_identities" yaml:"federatedIdentities,omitempty"`

			LastName *string `tfsdk:"last_name" yaml:"lastName,omitempty"`

			Id *string `tfsdk:"id" yaml:"id,omitempty"`

			RequiredActions *[]string `tfsdk:"required_actions" yaml:"requiredActions,omitempty"`

			Username *string `tfsdk:"username" yaml:"username,omitempty"`
		} `tfsdk:"user" yaml:"user,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewKeycloakOrgKeycloakUserV1Alpha1Resource() resource.Resource {
	return &KeycloakOrgKeycloakUserV1Alpha1Resource{}
}

func (r *KeycloakOrgKeycloakUserV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_keycloak_org_keycloak_user_v1alpha1"
}

func (r *KeycloakOrgKeycloakUserV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "KeycloakUser is the Schema for the keycloakusers API.",
		MarkdownDescription: "KeycloakUser is the Schema for the keycloakusers API.",
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
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.RequiresReplace(),
						},
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
				Description:         "KeycloakUserSpec defines the desired state of KeycloakUser.",
				MarkdownDescription: "KeycloakUserSpec defines the desired state of KeycloakUser.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"realm_selector": {
						Description:         "Selector for looking up KeycloakRealm Custom Resources.",
						MarkdownDescription: "Selector for looking up KeycloakRealm Custom Resources.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"match_expressions": {
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"key": {
										Description:         "key is the label key that the selector applies to.",
										MarkdownDescription: "key is the label key that the selector applies to.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"operator": {
										Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
										MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"values": {
										Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
										MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

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

							"match_labels": {
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"user": {
						Description:         "Keycloak User REST object.",
						MarkdownDescription: "Keycloak User REST object.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"client_roles": {
								Description:         "A set of Client Roles.",
								MarkdownDescription: "A set of Client Roles.",

								Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"email": {
								Description:         "Email.",
								MarkdownDescription: "Email.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"groups": {
								Description:         "A set of Groups.",
								MarkdownDescription: "A set of Groups.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"realm_roles": {
								Description:         "A set of Realm Roles.",
								MarkdownDescription: "A set of Realm Roles.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"attributes": {
								Description:         "A set of Attributes.",
								MarkdownDescription: "A set of Attributes.",

								Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enabled": {
								Description:         "User enabled flag.",
								MarkdownDescription: "User enabled flag.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"first_name": {
								Description:         "First Name.",
								MarkdownDescription: "First Name.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"credentials": {
								Description:         "A set of Credentials.",
								MarkdownDescription: "A set of Credentials.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"temporary": {
										Description:         "True if this credential object is temporary.",
										MarkdownDescription: "True if this credential object is temporary.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": {
										Description:         "Credential Type.",
										MarkdownDescription: "Credential Type.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"value": {
										Description:         "Credential Value.",
										MarkdownDescription: "Credential Value.",

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

							"email_verified": {
								Description:         "True if email has already been verified.",
								MarkdownDescription: "True if email has already been verified.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"federated_identities": {
								Description:         "A set of Federated Identities.",
								MarkdownDescription: "A set of Federated Identities.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"identity_provider": {
										Description:         "Federated Identity Provider.",
										MarkdownDescription: "Federated Identity Provider.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"user_id": {
										Description:         "Federated Identity User ID.",
										MarkdownDescription: "Federated Identity User ID.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"user_name": {
										Description:         "Federated Identity User Name.",
										MarkdownDescription: "Federated Identity User Name.",

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

							"last_name": {
								Description:         "Last Name.",
								MarkdownDescription: "Last Name.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"id": {
								Description:         "User ID.",
								MarkdownDescription: "User ID.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"required_actions": {
								Description:         "A set of Required Actions.",
								MarkdownDescription: "A set of Required Actions.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"username": {
								Description:         "User Name.",
								MarkdownDescription: "User Name.",

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

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *KeycloakOrgKeycloakUserV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_keycloak_org_keycloak_user_v1alpha1")

	var state KeycloakOrgKeycloakUserV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel KeycloakOrgKeycloakUserV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("keycloak.org/v1alpha1")
	goModel.Kind = utilities.Ptr("KeycloakUser")

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

func (r *KeycloakOrgKeycloakUserV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_keycloak_org_keycloak_user_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *KeycloakOrgKeycloakUserV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_keycloak_org_keycloak_user_v1alpha1")

	var state KeycloakOrgKeycloakUserV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel KeycloakOrgKeycloakUserV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("keycloak.org/v1alpha1")
	goModel.Kind = utilities.Ptr("KeycloakUser")

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

func (r *KeycloakOrgKeycloakUserV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_keycloak_org_keycloak_user_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
