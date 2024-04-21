/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package ceph_rook_io_v1

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
	_ datasource.DataSource = &CephRookIoCephObjectStoreUserV1Manifest{}
)

func NewCephRookIoCephObjectStoreUserV1Manifest() datasource.DataSource {
	return &CephRookIoCephObjectStoreUserV1Manifest{}
}

type CephRookIoCephObjectStoreUserV1Manifest struct{}

type CephRookIoCephObjectStoreUserV1ManifestData struct {
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
		Capabilities *struct {
			Amz_cache     *string `tfsdk:"amz_cache" json:"amz-cache,omitempty"`
			Bilog         *string `tfsdk:"bilog" json:"bilog,omitempty"`
			Bucket        *string `tfsdk:"bucket" json:"bucket,omitempty"`
			Buckets       *string `tfsdk:"buckets" json:"buckets,omitempty"`
			Datalog       *string `tfsdk:"datalog" json:"datalog,omitempty"`
			Info          *string `tfsdk:"info" json:"info,omitempty"`
			Mdlog         *string `tfsdk:"mdlog" json:"mdlog,omitempty"`
			Metadata      *string `tfsdk:"metadata" json:"metadata,omitempty"`
			Oidc_provider *string `tfsdk:"oidc_provider" json:"oidc-provider,omitempty"`
			Ratelimit     *string `tfsdk:"ratelimit" json:"ratelimit,omitempty"`
			Roles         *string `tfsdk:"roles" json:"roles,omitempty"`
			Usage         *string `tfsdk:"usage" json:"usage,omitempty"`
			User          *string `tfsdk:"user" json:"user,omitempty"`
			User_policy   *string `tfsdk:"user_policy" json:"user-policy,omitempty"`
			Users         *string `tfsdk:"users" json:"users,omitempty"`
			Zone          *string `tfsdk:"zone" json:"zone,omitempty"`
		} `tfsdk:"capabilities" json:"capabilities,omitempty"`
		ClusterNamespace *string `tfsdk:"cluster_namespace" json:"clusterNamespace,omitempty"`
		DisplayName      *string `tfsdk:"display_name" json:"displayName,omitempty"`
		Quotas           *struct {
			MaxBuckets *int64  `tfsdk:"max_buckets" json:"maxBuckets,omitempty"`
			MaxObjects *int64  `tfsdk:"max_objects" json:"maxObjects,omitempty"`
			MaxSize    *string `tfsdk:"max_size" json:"maxSize,omitempty"`
		} `tfsdk:"quotas" json:"quotas,omitempty"`
		Store *string `tfsdk:"store" json:"store,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CephRookIoCephObjectStoreUserV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_ceph_rook_io_ceph_object_store_user_v1_manifest"
}

func (r *CephRookIoCephObjectStoreUserV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CephObjectStoreUser represents a Ceph Object Store Gateway User",
		MarkdownDescription: "CephObjectStoreUser represents a Ceph Object Store Gateway User",
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
				Description:         "ObjectStoreUserSpec represent the spec of an Objectstoreuser",
				MarkdownDescription: "ObjectStoreUserSpec represent the spec of an Objectstoreuser",
				Attributes: map[string]schema.Attribute{
					"capabilities": schema.SingleNestedAttribute{
						Description:         "Additional admin-level capabilities for the Ceph object store user",
						MarkdownDescription: "Additional admin-level capabilities for the Ceph object store user",
						Attributes: map[string]schema.Attribute{
							"amz_cache": schema.StringAttribute{
								Description:         "Add capabilities for user to send request to RGW Cache API header. Documented in https://docs.ceph.com/en/quincy/radosgw/rgw-cache/#cache-api",
								MarkdownDescription: "Add capabilities for user to send request to RGW Cache API header. Documented in https://docs.ceph.com/en/quincy/radosgw/rgw-cache/#cache-api",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("*", "read", "write", "read, write"),
								},
							},

							"bilog": schema.StringAttribute{
								Description:         "Add capabilities for user to change bucket index logging. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								MarkdownDescription: "Add capabilities for user to change bucket index logging. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("*", "read", "write", "read, write"),
								},
							},

							"bucket": schema.StringAttribute{
								Description:         "Admin capabilities to read/write Ceph object store buckets. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								MarkdownDescription: "Admin capabilities to read/write Ceph object store buckets. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("*", "read", "write", "read, write"),
								},
							},

							"buckets": schema.StringAttribute{
								Description:         "Admin capabilities to read/write Ceph object store buckets. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								MarkdownDescription: "Admin capabilities to read/write Ceph object store buckets. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("*", "read", "write", "read, write"),
								},
							},

							"datalog": schema.StringAttribute{
								Description:         "Add capabilities for user to change data logging. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								MarkdownDescription: "Add capabilities for user to change data logging. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("*", "read", "write", "read, write"),
								},
							},

							"info": schema.StringAttribute{
								Description:         "Admin capabilities to read/write information about the user. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								MarkdownDescription: "Admin capabilities to read/write information about the user. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("*", "read", "write", "read, write"),
								},
							},

							"mdlog": schema.StringAttribute{
								Description:         "Add capabilities for user to change metadata logging. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								MarkdownDescription: "Add capabilities for user to change metadata logging. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("*", "read", "write", "read, write"),
								},
							},

							"metadata": schema.StringAttribute{
								Description:         "Admin capabilities to read/write Ceph object store metadata. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								MarkdownDescription: "Admin capabilities to read/write Ceph object store metadata. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("*", "read", "write", "read, write"),
								},
							},

							"oidc_provider": schema.StringAttribute{
								Description:         "Add capabilities for user to change oidc provider. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								MarkdownDescription: "Add capabilities for user to change oidc provider. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("*", "read", "write", "read, write"),
								},
							},

							"ratelimit": schema.StringAttribute{
								Description:         "Add capabilities for user to set rate limiter for user and bucket. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								MarkdownDescription: "Add capabilities for user to set rate limiter for user and bucket. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("*", "read", "write", "read, write"),
								},
							},

							"roles": schema.StringAttribute{
								Description:         "Admin capabilities to read/write roles for user. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								MarkdownDescription: "Admin capabilities to read/write roles for user. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("*", "read", "write", "read, write"),
								},
							},

							"usage": schema.StringAttribute{
								Description:         "Admin capabilities to read/write Ceph object store usage. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								MarkdownDescription: "Admin capabilities to read/write Ceph object store usage. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("*", "read", "write", "read, write"),
								},
							},

							"user": schema.StringAttribute{
								Description:         "Admin capabilities to read/write Ceph object store users. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								MarkdownDescription: "Admin capabilities to read/write Ceph object store users. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("*", "read", "write", "read, write"),
								},
							},

							"user_policy": schema.StringAttribute{
								Description:         "Add capabilities for user to change user policies. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								MarkdownDescription: "Add capabilities for user to change user policies. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("*", "read", "write", "read, write"),
								},
							},

							"users": schema.StringAttribute{
								Description:         "Admin capabilities to read/write Ceph object store users. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								MarkdownDescription: "Admin capabilities to read/write Ceph object store users. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("*", "read", "write", "read, write"),
								},
							},

							"zone": schema.StringAttribute{
								Description:         "Admin capabilities to read/write Ceph object store zones. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								MarkdownDescription: "Admin capabilities to read/write Ceph object store zones. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("*", "read", "write", "read, write"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"cluster_namespace": schema.StringAttribute{
						Description:         "The namespace where the parent CephCluster and CephObjectStore are found",
						MarkdownDescription: "The namespace where the parent CephCluster and CephObjectStore are found",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"display_name": schema.StringAttribute{
						Description:         "The display name for the ceph users",
						MarkdownDescription: "The display name for the ceph users",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"quotas": schema.SingleNestedAttribute{
						Description:         "ObjectUserQuotaSpec can be used to set quotas for the object store user to limit their usage. See the [Ceph docs](https://docs.ceph.com/en/latest/radosgw/admin/?#quota-management) for more",
						MarkdownDescription: "ObjectUserQuotaSpec can be used to set quotas for the object store user to limit their usage. See the [Ceph docs](https://docs.ceph.com/en/latest/radosgw/admin/?#quota-management) for more",
						Attributes: map[string]schema.Attribute{
							"max_buckets": schema.Int64Attribute{
								Description:         "Maximum bucket limit for the ceph user",
								MarkdownDescription: "Maximum bucket limit for the ceph user",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_objects": schema.Int64Attribute{
								Description:         "Maximum number of objects across all the user's buckets",
								MarkdownDescription: "Maximum number of objects across all the user's buckets",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_size": schema.StringAttribute{
								Description:         "Maximum size limit of all objects across all the user's bucketsSee https://pkg.go.dev/k8s.io/apimachinery/pkg/api/resource#Quantity for more info.",
								MarkdownDescription: "Maximum size limit of all objects across all the user's bucketsSee https://pkg.go.dev/k8s.io/apimachinery/pkg/api/resource#Quantity for more info.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"store": schema.StringAttribute{
						Description:         "The store the user will be created in",
						MarkdownDescription: "The store the user will be created in",
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
	}
}

func (r *CephRookIoCephObjectStoreUserV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_ceph_rook_io_ceph_object_store_user_v1_manifest")

	var model CephRookIoCephObjectStoreUserV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("ceph.rook.io/v1")
	model.Kind = pointer.String("CephObjectStoreUser")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
