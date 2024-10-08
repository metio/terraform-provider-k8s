/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package keycloak_org_v1alpha1

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
	_ datasource.DataSource = &KeycloakOrgKeycloakBackupV1Alpha1Manifest{}
)

func NewKeycloakOrgKeycloakBackupV1Alpha1Manifest() datasource.DataSource {
	return &KeycloakOrgKeycloakBackupV1Alpha1Manifest{}
}

type KeycloakOrgKeycloakBackupV1Alpha1Manifest struct{}

type KeycloakOrgKeycloakBackupV1Alpha1ManifestData struct {
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
		Aws *struct {
			CredentialsSecretName   *string `tfsdk:"credentials_secret_name" json:"credentialsSecretName,omitempty"`
			EncryptionKeySecretName *string `tfsdk:"encryption_key_secret_name" json:"encryptionKeySecretName,omitempty"`
			Schedule                *string `tfsdk:"schedule" json:"schedule,omitempty"`
		} `tfsdk:"aws" json:"aws,omitempty"`
		InstanceSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"instance_selector" json:"instanceSelector,omitempty"`
		Restore          *bool   `tfsdk:"restore" json:"restore,omitempty"`
		StorageClassName *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KeycloakOrgKeycloakBackupV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_keycloak_org_keycloak_backup_v1alpha1_manifest"
}

func (r *KeycloakOrgKeycloakBackupV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "KeycloakBackup is the Schema for the keycloakbackups API.",
		MarkdownDescription: "KeycloakBackup is the Schema for the keycloakbackups API.",
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
				Description:         "KeycloakBackupSpec defines the desired state of KeycloakBackup.",
				MarkdownDescription: "KeycloakBackupSpec defines the desired state of KeycloakBackup.",
				Attributes: map[string]schema.Attribute{
					"aws": schema.SingleNestedAttribute{
						Description:         "If provided, an automatic database backup will be created on AWS S3 instead of a local Persistent Volume. If this property is not provided - a local Persistent Volume backup will be chosen.",
						MarkdownDescription: "If provided, an automatic database backup will be created on AWS S3 instead of a local Persistent Volume. If this property is not provided - a local Persistent Volume backup will be chosen.",
						Attributes: map[string]schema.Attribute{
							"credentials_secret_name": schema.StringAttribute{
								Description:         "Provides a secret name used for connecting to AWS S3 Service. The secret needs to be in the following form: apiVersion: v1 kind: Secret metadata: name: <Secret name> type: Opaque stringData: AWS_S3_BUCKET_NAME: <S3 Bucket Name> AWS_ACCESS_KEY_ID: <AWS Access Key ID> AWS_SECRET_ACCESS_KEY: <AWS Secret Key> For more information, please refer to the Operator documentation.",
								MarkdownDescription: "Provides a secret name used for connecting to AWS S3 Service. The secret needs to be in the following form: apiVersion: v1 kind: Secret metadata: name: <Secret name> type: Opaque stringData: AWS_S3_BUCKET_NAME: <S3 Bucket Name> AWS_ACCESS_KEY_ID: <AWS Access Key ID> AWS_SECRET_ACCESS_KEY: <AWS Secret Key> For more information, please refer to the Operator documentation.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"encryption_key_secret_name": schema.StringAttribute{
								Description:         "If provided, the database backup will be encrypted. Provides a secret name used for encrypting database data. The secret needs to be in the following form: apiVersion: v1 kind: Secret metadata: name: <Secret name> type: Opaque stringData: GPG_PUBLIC_KEY: <GPG Public Key> GPG_TRUST_MODEL: <GPG Trust Model> GPG_RECIPIENT: <GPG Recipient> For more information, please refer to the Operator documentation.",
								MarkdownDescription: "If provided, the database backup will be encrypted. Provides a secret name used for encrypting database data. The secret needs to be in the following form: apiVersion: v1 kind: Secret metadata: name: <Secret name> type: Opaque stringData: GPG_PUBLIC_KEY: <GPG Public Key> GPG_TRUST_MODEL: <GPG Trust Model> GPG_RECIPIENT: <GPG Recipient> For more information, please refer to the Operator documentation.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"schedule": schema.StringAttribute{
								Description:         "If specified, it will be used as a schedule for creating a CronJob.",
								MarkdownDescription: "If specified, it will be used as a schedule for creating a CronJob.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"instance_selector": schema.SingleNestedAttribute{
						Description:         "Selector for looking up Keycloak Custom Resources.",
						MarkdownDescription: "Selector for looking up Keycloak Custom Resources.",
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

					"restore": schema.BoolAttribute{
						Description:         "Controls automatic restore behavior. Currently not implemented. In the future this will be used to trigger automatic restore for a given KeycloakBackup. Each backup will correspond to a single snapshot of the database (stored either in a Persistent Volume or AWS). If a user wants to restore it, all he/she needs to do is to change this flag to true. Potentially, it will be possible to restore a single backup multiple times.",
						MarkdownDescription: "Controls automatic restore behavior. Currently not implemented. In the future this will be used to trigger automatic restore for a given KeycloakBackup. Each backup will correspond to a single snapshot of the database (stored either in a Persistent Volume or AWS). If a user wants to restore it, all he/she needs to do is to change this flag to true. Potentially, it will be possible to restore a single backup multiple times.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"storage_class_name": schema.StringAttribute{
						Description:         "Name of the StorageClass for Postgresql Backup Persistent Volume Claim",
						MarkdownDescription: "Name of the StorageClass for Postgresql Backup Persistent Volume Claim",
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

func (r *KeycloakOrgKeycloakBackupV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_keycloak_org_keycloak_backup_v1alpha1_manifest")

	var model KeycloakOrgKeycloakBackupV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("keycloak.org/v1alpha1")
	model.Kind = pointer.String("KeycloakBackup")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
