/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package ceph_rook_io_v1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"net/http"
)

var (
	_ datasource.DataSource              = &CephRookIoCephObjectStoreUserV1DataSource{}
	_ datasource.DataSourceWithConfigure = &CephRookIoCephObjectStoreUserV1DataSource{}
)

func NewCephRookIoCephObjectStoreUserV1DataSource() datasource.DataSource {
	return &CephRookIoCephObjectStoreUserV1DataSource{}
}

type CephRookIoCephObjectStoreUserV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type CephRookIoCephObjectStoreUserV1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

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

func (r *CephRookIoCephObjectStoreUserV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_ceph_rook_io_ceph_object_store_user_v1"
}

func (r *CephRookIoCephObjectStoreUserV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CephObjectStoreUser represents a Ceph Object Store Gateway User",
		MarkdownDescription: "CephObjectStoreUser represents a Ceph Object Store Gateway User",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
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
								Optional:            false,
								Computed:            true,
							},

							"bilog": schema.StringAttribute{
								Description:         "Add capabilities for user to change bucket index logging. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								MarkdownDescription: "Add capabilities for user to change bucket index logging. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"bucket": schema.StringAttribute{
								Description:         "Admin capabilities to read/write Ceph object store buckets. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								MarkdownDescription: "Admin capabilities to read/write Ceph object store buckets. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"buckets": schema.StringAttribute{
								Description:         "Admin capabilities to read/write Ceph object store buckets. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								MarkdownDescription: "Admin capabilities to read/write Ceph object store buckets. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"datalog": schema.StringAttribute{
								Description:         "Add capabilities for user to change data logging. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								MarkdownDescription: "Add capabilities for user to change data logging. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"info": schema.StringAttribute{
								Description:         "Admin capabilities to read/write information about the user. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								MarkdownDescription: "Admin capabilities to read/write information about the user. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"mdlog": schema.StringAttribute{
								Description:         "Add capabilities for user to change metadata logging. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								MarkdownDescription: "Add capabilities for user to change metadata logging. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"metadata": schema.StringAttribute{
								Description:         "Admin capabilities to read/write Ceph object store metadata. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								MarkdownDescription: "Admin capabilities to read/write Ceph object store metadata. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"oidc_provider": schema.StringAttribute{
								Description:         "Add capabilities for user to change oidc provider. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								MarkdownDescription: "Add capabilities for user to change oidc provider. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"ratelimit": schema.StringAttribute{
								Description:         "Add capabilities for user to set rate limiter for user and bucket. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								MarkdownDescription: "Add capabilities for user to set rate limiter for user and bucket. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"roles": schema.StringAttribute{
								Description:         "Admin capabilities to read/write roles for user. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								MarkdownDescription: "Admin capabilities to read/write roles for user. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"usage": schema.StringAttribute{
								Description:         "Admin capabilities to read/write Ceph object store usage. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								MarkdownDescription: "Admin capabilities to read/write Ceph object store usage. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"user": schema.StringAttribute{
								Description:         "Admin capabilities to read/write Ceph object store users. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								MarkdownDescription: "Admin capabilities to read/write Ceph object store users. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"user_policy": schema.StringAttribute{
								Description:         "Add capabilities for user to change user policies. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								MarkdownDescription: "Add capabilities for user to change user policies. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"users": schema.StringAttribute{
								Description:         "Admin capabilities to read/write Ceph object store users. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								MarkdownDescription: "Admin capabilities to read/write Ceph object store users. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"zone": schema.StringAttribute{
								Description:         "Admin capabilities to read/write Ceph object store zones. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								MarkdownDescription: "Admin capabilities to read/write Ceph object store zones. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"cluster_namespace": schema.StringAttribute{
						Description:         "The namespace where the parent CephCluster and CephObjectStore are found",
						MarkdownDescription: "The namespace where the parent CephCluster and CephObjectStore are found",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"display_name": schema.StringAttribute{
						Description:         "The display name for the ceph users",
						MarkdownDescription: "The display name for the ceph users",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"quotas": schema.SingleNestedAttribute{
						Description:         "ObjectUserQuotaSpec can be used to set quotas for the object store user to limit their usage. See the [Ceph docs](https://docs.ceph.com/en/latest/radosgw/admin/?#quota-management) for more",
						MarkdownDescription: "ObjectUserQuotaSpec can be used to set quotas for the object store user to limit their usage. See the [Ceph docs](https://docs.ceph.com/en/latest/radosgw/admin/?#quota-management) for more",
						Attributes: map[string]schema.Attribute{
							"max_buckets": schema.Int64Attribute{
								Description:         "Maximum bucket limit for the ceph user",
								MarkdownDescription: "Maximum bucket limit for the ceph user",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"max_objects": schema.Int64Attribute{
								Description:         "Maximum number of objects across all the user's buckets",
								MarkdownDescription: "Maximum number of objects across all the user's buckets",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"max_size": schema.StringAttribute{
								Description:         "Maximum size limit of all objects across all the user's buckets See https://pkg.go.dev/k8s.io/apimachinery/pkg/api/resource#Quantity for more info.",
								MarkdownDescription: "Maximum size limit of all objects across all the user's buckets See https://pkg.go.dev/k8s.io/apimachinery/pkg/api/resource#Quantity for more info.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"store": schema.StringAttribute{
						Description:         "The store the user will be created in",
						MarkdownDescription: "The store the user will be created in",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *CephRookIoCephObjectStoreUserV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *CephRookIoCephObjectStoreUserV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_ceph_rook_io_ceph_object_store_user_v1")

	var data CephRookIoCephObjectStoreUserV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "ceph.rook.io", Version: "v1", Resource: "cephobjectstoreusers"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		var statusError *k8sErrors.StatusError
		if errors.As(err, &statusError) {
			if statusError.Status().Code == http.StatusNotFound {
				response.Diagnostics.AddError(
					"Unable to find resource",
					fmt.Sprintf("The requested resource cannot be found. "+
						"Make sure that it does exist in your cluster and you have set the correct name and namespace configured.\n\n"+
						"Namespace: %s\n"+
						"Name: %s", data.Metadata.Namespace, data.Metadata.Name),
				)
				return
			}
		} else {
			response.Diagnostics.AddError(
				"Unable to GET resource",
				fmt.Sprintf("An unexpected error occurred while reading the resource. "+
					"Please report this issue to the provider developers.\n\n"+
					"GET Error (%T): %s", err, err.Error()),
			)
		}
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse CephRookIoCephObjectStoreUserV1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("ceph.rook.io/v1")
	data.Kind = pointer.String("CephObjectStoreUser")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
