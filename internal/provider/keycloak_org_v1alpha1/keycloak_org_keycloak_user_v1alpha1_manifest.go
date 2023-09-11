/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package keycloak_org_v1alpha1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &KeycloakOrgKeycloakUserV1Alpha1Manifest{}
)

func NewKeycloakOrgKeycloakUserV1Alpha1Manifest() datasource.DataSource {
	return &KeycloakOrgKeycloakUserV1Alpha1Manifest{}
}

type KeycloakOrgKeycloakUserV1Alpha1Manifest struct{}

type KeycloakOrgKeycloakUserV1Alpha1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
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
		RealmSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"realm_selector" json:"realmSelector,omitempty"`
		User *struct {
			Attributes  *map[string][]string `tfsdk:"attributes" json:"attributes,omitempty"`
			ClientRoles *map[string][]string `tfsdk:"client_roles" json:"clientRoles,omitempty"`
			Credentials *[]struct {
				Temporary *bool   `tfsdk:"temporary" json:"temporary,omitempty"`
				Type      *string `tfsdk:"type" json:"type,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"credentials" json:"credentials,omitempty"`
			Email               *string `tfsdk:"email" json:"email,omitempty"`
			EmailVerified       *bool   `tfsdk:"email_verified" json:"emailVerified,omitempty"`
			Enabled             *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			FederatedIdentities *[]struct {
				IdentityProvider *string `tfsdk:"identity_provider" json:"identityProvider,omitempty"`
				UserId           *string `tfsdk:"user_id" json:"userId,omitempty"`
				UserName         *string `tfsdk:"user_name" json:"userName,omitempty"`
			} `tfsdk:"federated_identities" json:"federatedIdentities,omitempty"`
			FirstName       *string   `tfsdk:"first_name" json:"firstName,omitempty"`
			Groups          *[]string `tfsdk:"groups" json:"groups,omitempty"`
			Id              *string   `tfsdk:"id" json:"id,omitempty"`
			LastName        *string   `tfsdk:"last_name" json:"lastName,omitempty"`
			RealmRoles      *[]string `tfsdk:"realm_roles" json:"realmRoles,omitempty"`
			RequiredActions *[]string `tfsdk:"required_actions" json:"requiredActions,omitempty"`
			Username        *string   `tfsdk:"username" json:"username,omitempty"`
		} `tfsdk:"user" json:"user,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KeycloakOrgKeycloakUserV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_keycloak_org_keycloak_user_v1alpha1_manifest"
}

func (r *KeycloakOrgKeycloakUserV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "KeycloakUser is the Schema for the keycloakusers API.",
		MarkdownDescription: "KeycloakUser is the Schema for the keycloakusers API.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

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
				Description:         "KeycloakUserSpec defines the desired state of KeycloakUser.",
				MarkdownDescription: "KeycloakUserSpec defines the desired state of KeycloakUser.",
				Attributes: map[string]schema.Attribute{
					"realm_selector": schema.SingleNestedAttribute{
						Description:         "Selector for looking up KeycloakRealm Custom Resources.",
						MarkdownDescription: "Selector for looking up KeycloakRealm Custom Resources.",
						Attributes: map[string]schema.Attribute{
							"match_expressions": schema.ListNestedAttribute{
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "key is the label key that the selector applies to.",
											MarkdownDescription: "key is the label key that the selector applies to.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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

							"match_labels": schema.MapAttribute{
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

					"user": schema.SingleNestedAttribute{
						Description:         "Keycloak User REST object.",
						MarkdownDescription: "Keycloak User REST object.",
						Attributes: map[string]schema.Attribute{
							"attributes": schema.MapAttribute{
								Description:         "A set of Attributes.",
								MarkdownDescription: "A set of Attributes.",
								ElementType:         types.ListType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"client_roles": schema.MapAttribute{
								Description:         "A set of Client Roles.",
								MarkdownDescription: "A set of Client Roles.",
								ElementType:         types.ListType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"credentials": schema.ListNestedAttribute{
								Description:         "A set of Credentials.",
								MarkdownDescription: "A set of Credentials.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"temporary": schema.BoolAttribute{
											Description:         "True if this credential object is temporary.",
											MarkdownDescription: "True if this credential object is temporary.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"type": schema.StringAttribute{
											Description:         "Credential Type.",
											MarkdownDescription: "Credential Type.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Credential Value.",
											MarkdownDescription: "Credential Value.",
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

							"email": schema.StringAttribute{
								Description:         "Email.",
								MarkdownDescription: "Email.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"email_verified": schema.BoolAttribute{
								Description:         "True if email has already been verified.",
								MarkdownDescription: "True if email has already been verified.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enabled": schema.BoolAttribute{
								Description:         "User enabled flag.",
								MarkdownDescription: "User enabled flag.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"federated_identities": schema.ListNestedAttribute{
								Description:         "A set of Federated Identities.",
								MarkdownDescription: "A set of Federated Identities.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"identity_provider": schema.StringAttribute{
											Description:         "Federated Identity Provider.",
											MarkdownDescription: "Federated Identity Provider.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"user_id": schema.StringAttribute{
											Description:         "Federated Identity User ID.",
											MarkdownDescription: "Federated Identity User ID.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"user_name": schema.StringAttribute{
											Description:         "Federated Identity User Name.",
											MarkdownDescription: "Federated Identity User Name.",
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

							"first_name": schema.StringAttribute{
								Description:         "First Name.",
								MarkdownDescription: "First Name.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"groups": schema.ListAttribute{
								Description:         "A set of Groups.",
								MarkdownDescription: "A set of Groups.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"id": schema.StringAttribute{
								Description:         "User ID.",
								MarkdownDescription: "User ID.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"last_name": schema.StringAttribute{
								Description:         "Last Name.",
								MarkdownDescription: "Last Name.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"realm_roles": schema.ListAttribute{
								Description:         "A set of Realm Roles.",
								MarkdownDescription: "A set of Realm Roles.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"required_actions": schema.ListAttribute{
								Description:         "A set of Required Actions.",
								MarkdownDescription: "A set of Required Actions.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"username": schema.StringAttribute{
								Description:         "User Name.",
								MarkdownDescription: "User Name.",
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
		},
	}
}

func (r *KeycloakOrgKeycloakUserV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_keycloak_org_keycloak_user_v1alpha1_manifest")

	var model KeycloakOrgKeycloakUserV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("keycloak.org/v1alpha1")
	model.Kind = pointer.String("KeycloakUser")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
