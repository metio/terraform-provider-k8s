/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package edp_epam_com_v1alpha1

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
	_ datasource.DataSource = &EdpEpamComNexusBlobStoreV1Alpha1Manifest{}
)

func NewEdpEpamComNexusBlobStoreV1Alpha1Manifest() datasource.DataSource {
	return &EdpEpamComNexusBlobStoreV1Alpha1Manifest{}
}

type EdpEpamComNexusBlobStoreV1Alpha1Manifest struct{}

type EdpEpamComNexusBlobStoreV1Alpha1ManifestData struct {
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
		File *struct {
			Path *string `tfsdk:"path" json:"path,omitempty"`
		} `tfsdk:"file" json:"file,omitempty"`
		Name     *string `tfsdk:"name" json:"name,omitempty"`
		NexusRef *struct {
			Kind *string `tfsdk:"kind" json:"kind,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"nexus_ref" json:"nexusRef,omitempty"`
		S3 *struct {
			AdvancedBucketConnection *struct {
				Endpoint              *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
				ForcePathStyle        *bool   `tfsdk:"force_path_style" json:"forcePathStyle,omitempty"`
				MaxConnectionPoolSize *int64  `tfsdk:"max_connection_pool_size" json:"maxConnectionPoolSize,omitempty"`
				SignerType            *string `tfsdk:"signer_type" json:"signerType,omitempty"`
			} `tfsdk:"advanced_bucket_connection" json:"advancedBucketConnection,omitempty"`
			Bucket *struct {
				Expiration *int64  `tfsdk:"expiration" json:"expiration,omitempty"`
				Name       *string `tfsdk:"name" json:"name,omitempty"`
				Prefix     *string `tfsdk:"prefix" json:"prefix,omitempty"`
				Region     *string `tfsdk:"region" json:"region,omitempty"`
			} `tfsdk:"bucket" json:"bucket,omitempty"`
			BucketSecurity *struct {
				AccessKeyId *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
				} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
				Role            *string `tfsdk:"role" json:"role,omitempty"`
				SecretAccessKey *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
				} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
				SessionToken *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
				} `tfsdk:"session_token" json:"sessionToken,omitempty"`
			} `tfsdk:"bucket_security" json:"bucketSecurity,omitempty"`
			Encryption *struct {
				EncryptionKey  *string `tfsdk:"encryption_key" json:"encryptionKey,omitempty"`
				EncryptionType *string `tfsdk:"encryption_type" json:"encryptionType,omitempty"`
			} `tfsdk:"encryption" json:"encryption,omitempty"`
		} `tfsdk:"s3" json:"s3,omitempty"`
		SoftQuota *struct {
			Limit *int64  `tfsdk:"limit" json:"limit,omitempty"`
			Type  *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"soft_quota" json:"softQuota,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *EdpEpamComNexusBlobStoreV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_edp_epam_com_nexus_blob_store_v1alpha1_manifest"
}

func (r *EdpEpamComNexusBlobStoreV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "NexusBlobStore is the Schema for the nexusblobstores API.",
		MarkdownDescription: "NexusBlobStore is the Schema for the nexusblobstores API.",
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
				Description:         "NexusBlobStoreSpec defines the desired state of NexusBlobStore.",
				MarkdownDescription: "NexusBlobStoreSpec defines the desired state of NexusBlobStore.",
				Attributes: map[string]schema.Attribute{
					"file": schema.SingleNestedAttribute{
						Description:         "File type blobstore.",
						MarkdownDescription: "File type blobstore.",
						Attributes: map[string]schema.Attribute{
							"path": schema.StringAttribute{
								Description:         "The path to the blobstore contents. This can be an absolute path to anywhere on the system Nexus Repository Manager has access to it or can be a path relative to the sonatype-work directory.",
								MarkdownDescription: "The path to the blobstore contents. This can be an absolute path to anywhere on the system Nexus Repository Manager has access to it or can be a path relative to the sonatype-work directory.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"name": schema.StringAttribute{
						Description:         "Name of the BlobStore. Name should be unique across all BlobStores.",
						MarkdownDescription: "Name of the BlobStore. Name should be unique across all BlobStores.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"nexus_ref": schema.SingleNestedAttribute{
						Description:         "NexusRef is a reference to Nexus custom resource.",
						MarkdownDescription: "NexusRef is a reference to Nexus custom resource.",
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Description:         "Kind specifies the kind of the Nexus resource.",
								MarkdownDescription: "Kind specifies the kind of the Nexus resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name specifies the name of the Nexus resource.",
								MarkdownDescription: "Name specifies the name of the Nexus resource.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"s3": schema.SingleNestedAttribute{
						Description:         "S3 type blobstore.",
						MarkdownDescription: "S3 type blobstore.",
						Attributes: map[string]schema.Attribute{
							"advanced_bucket_connection": schema.SingleNestedAttribute{
								Description:         "A custom endpoint URL, signer type and whether path style access is enabled.",
								MarkdownDescription: "A custom endpoint URL, signer type and whether path style access is enabled.",
								Attributes: map[string]schema.Attribute{
									"endpoint": schema.StringAttribute{
										Description:         "A custom endpoint URL for third party object stores using the S3 API.",
										MarkdownDescription: "A custom endpoint URL for third party object stores using the S3 API.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"force_path_style": schema.BoolAttribute{
										Description:         "Setting this flag will result in path-style access being used for all requests.",
										MarkdownDescription: "Setting this flag will result in path-style access being used for all requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_connection_pool_size": schema.Int64Attribute{
										Description:         "Setting this value will override the default connection pool size of Nexus of the s3 client for this blobstore.",
										MarkdownDescription: "Setting this value will override the default connection pool size of Nexus of the s3 client for this blobstore.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"signer_type": schema.StringAttribute{
										Description:         "An API signature version which may be required for third party object stores using the S3 API.",
										MarkdownDescription: "An API signature version which may be required for third party object stores using the S3 API.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("DEFAULT", "S3SignerType", "AWSS3V4SignerType"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"bucket": schema.SingleNestedAttribute{
								Description:         "Details of the S3 bucket such as name and region.",
								MarkdownDescription: "Details of the S3 bucket such as name and region.",
								Attributes: map[string]schema.Attribute{
									"expiration": schema.Int64Attribute{
										Description:         "How many days until deleted blobs are finally removed from the S3 bucket (-1 to disable).",
										MarkdownDescription: "How many days until deleted blobs are finally removed from the S3 bucket (-1 to disable).",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "The name of the S3 bucket.",
										MarkdownDescription: "The name of the S3 bucket.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"prefix": schema.StringAttribute{
										Description:         "The S3 blob store (i.e. S3 object) key prefix.",
										MarkdownDescription: "The S3 blob store (i.e. S3 object) key prefix.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"region": schema.StringAttribute{
										Description:         "The AWS region to create a new S3 bucket in or an existing S3 bucket's region.",
										MarkdownDescription: "The AWS region to create a new S3 bucket in or an existing S3 bucket's region.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
							},

							"bucket_security": schema.SingleNestedAttribute{
								Description:         "Security details for granting access the S3 API.",
								MarkdownDescription: "Security details for granting access the S3 API.",
								Attributes: map[string]schema.Attribute{
									"access_key_id": schema.SingleNestedAttribute{
										Description:         "An IAM access key ID for granting access to the S3 bucket.",
										MarkdownDescription: "An IAM access key ID for granting access to the S3 bucket.",
										Attributes: map[string]schema.Attribute{
											"config_map_key_ref": schema.SingleNestedAttribute{
												Description:         "Selects a key of a ConfigMap.",
												MarkdownDescription: "Selects a key of a ConfigMap.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key to select.",
														MarkdownDescription: "The key to select.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_key_ref": schema.SingleNestedAttribute{
												Description:         "Selects a key of a secret.",
												MarkdownDescription: "Selects a key of a secret.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.",
														MarkdownDescription: "The key of the secret to select from.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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

									"role": schema.StringAttribute{
										Description:         "An IAM role to assume in order to access the S3 bucket.",
										MarkdownDescription: "An IAM role to assume in order to access the S3 bucket.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"secret_access_key": schema.SingleNestedAttribute{
										Description:         "The secret access key associated with the specified IAM access key ID.",
										MarkdownDescription: "The secret access key associated with the specified IAM access key ID.",
										Attributes: map[string]schema.Attribute{
											"config_map_key_ref": schema.SingleNestedAttribute{
												Description:         "Selects a key of a ConfigMap.",
												MarkdownDescription: "Selects a key of a ConfigMap.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key to select.",
														MarkdownDescription: "The key to select.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_key_ref": schema.SingleNestedAttribute{
												Description:         "Selects a key of a secret.",
												MarkdownDescription: "Selects a key of a secret.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.",
														MarkdownDescription: "The key of the secret to select from.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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

									"session_token": schema.SingleNestedAttribute{
										Description:         "An AWS STS session token associated with temporary security credentials which grant access to the S3 bucket.",
										MarkdownDescription: "An AWS STS session token associated with temporary security credentials which grant access to the S3 bucket.",
										Attributes: map[string]schema.Attribute{
											"config_map_key_ref": schema.SingleNestedAttribute{
												Description:         "Selects a key of a ConfigMap.",
												MarkdownDescription: "Selects a key of a ConfigMap.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key to select.",
														MarkdownDescription: "The key to select.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_key_ref": schema.SingleNestedAttribute{
												Description:         "Selects a key of a secret.",
												MarkdownDescription: "Selects a key of a secret.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.",
														MarkdownDescription: "The key of the secret to select from.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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
								Required: false,
								Optional: true,
								Computed: false,
							},

							"encryption": schema.SingleNestedAttribute{
								Description:         "The type of encryption to use if any.",
								MarkdownDescription: "The type of encryption to use if any.",
								Attributes: map[string]schema.Attribute{
									"encryption_key": schema.StringAttribute{
										Description:         "If using KMS encryption, you can supply a Key ID. If left blank, then the default will be used.",
										MarkdownDescription: "If using KMS encryption, you can supply a Key ID. If left blank, then the default will be used.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"encryption_type": schema.StringAttribute{
										Description:         "The type of S3 server side encryption to use.",
										MarkdownDescription: "The type of S3 server side encryption to use.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("none", "s3ManagedEncryption", "kmsManagedEncryption"),
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

					"soft_quota": schema.SingleNestedAttribute{
						Description:         "Settings to control the soft quota.",
						MarkdownDescription: "Settings to control the soft quota.",
						Attributes: map[string]schema.Attribute{
							"limit": schema.Int64Attribute{
								Description:         "The limit in MB.",
								MarkdownDescription: "The limit in MB.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
							},

							"type": schema.StringAttribute{
								Description:         "Type of the soft quota.",
								MarkdownDescription: "Type of the soft quota.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("spaceRemainingQuota", "spaceUsedQuota"),
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
	}
}

func (r *EdpEpamComNexusBlobStoreV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_edp_epam_com_nexus_blob_store_v1alpha1_manifest")

	var model EdpEpamComNexusBlobStoreV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("edp.epam.com/v1alpha1")
	model.Kind = pointer.String("NexusBlobStore")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
