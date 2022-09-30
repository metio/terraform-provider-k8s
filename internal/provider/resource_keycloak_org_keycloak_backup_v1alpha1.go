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

type KeycloakOrgKeycloakBackupV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*KeycloakOrgKeycloakBackupV1Alpha1Resource)(nil)
)

type KeycloakOrgKeycloakBackupV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type KeycloakOrgKeycloakBackupV1Alpha1GoModel struct {
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
		StorageClassName *string `tfsdk:"storage_class_name" yaml:"storageClassName,omitempty"`

		Aws *struct {
			Schedule *string `tfsdk:"schedule" yaml:"schedule,omitempty"`

			CredentialsSecretName *string `tfsdk:"credentials_secret_name" yaml:"credentialsSecretName,omitempty"`

			EncryptionKeySecretName *string `tfsdk:"encryption_key_secret_name" yaml:"encryptionKeySecretName,omitempty"`
		} `tfsdk:"aws" yaml:"aws,omitempty"`

		InstanceSelector *struct {
			MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`

			MatchExpressions *[]struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

				Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
			} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`
		} `tfsdk:"instance_selector" yaml:"instanceSelector,omitempty"`

		Restore *bool `tfsdk:"restore" yaml:"restore,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewKeycloakOrgKeycloakBackupV1Alpha1Resource() resource.Resource {
	return &KeycloakOrgKeycloakBackupV1Alpha1Resource{}
}

func (r *KeycloakOrgKeycloakBackupV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_keycloak_org_keycloak_backup_v1alpha1"
}

func (r *KeycloakOrgKeycloakBackupV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "KeycloakBackup is the Schema for the keycloakbackups API.",
		MarkdownDescription: "KeycloakBackup is the Schema for the keycloakbackups API.",
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
				Description:         "KeycloakBackupSpec defines the desired state of KeycloakBackup.",
				MarkdownDescription: "KeycloakBackupSpec defines the desired state of KeycloakBackup.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"storage_class_name": {
						Description:         "Name of the StorageClass for Postgresql Backup Persistent Volume Claim",
						MarkdownDescription: "Name of the StorageClass for Postgresql Backup Persistent Volume Claim",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"aws": {
						Description:         "If provided, an automatic database backup will be created on AWS S3 instead of a local Persistent Volume. If this property is not provided - a local Persistent Volume backup will be chosen.",
						MarkdownDescription: "If provided, an automatic database backup will be created on AWS S3 instead of a local Persistent Volume. If this property is not provided - a local Persistent Volume backup will be chosen.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"schedule": {
								Description:         "If specified, it will be used as a schedule for creating a CronJob.",
								MarkdownDescription: "If specified, it will be used as a schedule for creating a CronJob.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"credentials_secret_name": {
								Description:         "Provides a secret name used for connecting to AWS S3 Service. The secret needs to be in the following form:      apiVersion: v1     kind: Secret     metadata:       name: <Secret name>     type: Opaque     stringData:       AWS_S3_BUCKET_NAME: <S3 Bucket Name>       AWS_ACCESS_KEY_ID: <AWS Access Key ID>       AWS_SECRET_ACCESS_KEY: <AWS Secret Key>  For more information, please refer to the Operator documentation.",
								MarkdownDescription: "Provides a secret name used for connecting to AWS S3 Service. The secret needs to be in the following form:      apiVersion: v1     kind: Secret     metadata:       name: <Secret name>     type: Opaque     stringData:       AWS_S3_BUCKET_NAME: <S3 Bucket Name>       AWS_ACCESS_KEY_ID: <AWS Access Key ID>       AWS_SECRET_ACCESS_KEY: <AWS Secret Key>  For more information, please refer to the Operator documentation.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"encryption_key_secret_name": {
								Description:         "If provided, the database backup will be encrypted. Provides a secret name used for encrypting database data. The secret needs to be in the following form:      apiVersion: v1     kind: Secret     metadata:       name: <Secret name>     type: Opaque     stringData:       GPG_PUBLIC_KEY: <GPG Public Key>       GPG_TRUST_MODEL: <GPG Trust Model>       GPG_RECIPIENT: <GPG Recipient>  For more information, please refer to the Operator documentation.",
								MarkdownDescription: "If provided, the database backup will be encrypted. Provides a secret name used for encrypting database data. The secret needs to be in the following form:      apiVersion: v1     kind: Secret     metadata:       name: <Secret name>     type: Opaque     stringData:       GPG_PUBLIC_KEY: <GPG Public Key>       GPG_TRUST_MODEL: <GPG Trust Model>       GPG_RECIPIENT: <GPG Recipient>  For more information, please refer to the Operator documentation.",

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

					"instance_selector": {
						Description:         "Selector for looking up Keycloak Custom Resources.",
						MarkdownDescription: "Selector for looking up Keycloak Custom Resources.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"match_labels": {
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"restore": {
						Description:         "Controls automatic restore behavior. Currently not implemented.  In the future this will be used to trigger automatic restore for a given KeycloakBackup. Each backup will correspond to a single snapshot of the database (stored either in a Persistent Volume or AWS). If a user wants to restore it, all he/she needs to do is to change this flag to true. Potentially, it will be possible to restore a single backup multiple times.",
						MarkdownDescription: "Controls automatic restore behavior. Currently not implemented.  In the future this will be used to trigger automatic restore for a given KeycloakBackup. Each backup will correspond to a single snapshot of the database (stored either in a Persistent Volume or AWS). If a user wants to restore it, all he/she needs to do is to change this flag to true. Potentially, it will be possible to restore a single backup multiple times.",

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
		},
	}, nil
}

func (r *KeycloakOrgKeycloakBackupV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_keycloak_org_keycloak_backup_v1alpha1")

	var state KeycloakOrgKeycloakBackupV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel KeycloakOrgKeycloakBackupV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("keycloak.org/v1alpha1")
	goModel.Kind = utilities.Ptr("KeycloakBackup")

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

func (r *KeycloakOrgKeycloakBackupV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_keycloak_org_keycloak_backup_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *KeycloakOrgKeycloakBackupV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_keycloak_org_keycloak_backup_v1alpha1")

	var state KeycloakOrgKeycloakBackupV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel KeycloakOrgKeycloakBackupV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("keycloak.org/v1alpha1")
	goModel.Kind = utilities.Ptr("KeycloakBackup")

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

func (r *KeycloakOrgKeycloakBackupV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_keycloak_org_keycloak_backup_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
