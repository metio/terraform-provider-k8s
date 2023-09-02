/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package couchbase_com_v2

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &CouchbaseComCouchbaseReplicationV2DataSource{}
	_ datasource.DataSourceWithConfigure = &CouchbaseComCouchbaseReplicationV2DataSource{}
)

func NewCouchbaseComCouchbaseReplicationV2DataSource() datasource.DataSource {
	return &CouchbaseComCouchbaseReplicationV2DataSource{}
}

type CouchbaseComCouchbaseReplicationV2DataSource struct {
	kubernetesClient dynamic.Interface
}

type CouchbaseComCouchbaseReplicationV2DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	ExplicitMapping *struct {
		AllowRules *[]struct {
			SourceKeyspace *struct {
				Collection *string `tfsdk:"collection" json:"collection,omitempty"`
				Scope      *string `tfsdk:"scope" json:"scope,omitempty"`
			} `tfsdk:"source_keyspace" json:"sourceKeyspace,omitempty"`
			TargetKeyspace *struct {
				Collection *string `tfsdk:"collection" json:"collection,omitempty"`
				Scope      *string `tfsdk:"scope" json:"scope,omitempty"`
			} `tfsdk:"target_keyspace" json:"targetKeyspace,omitempty"`
		} `tfsdk:"allow_rules" json:"allowRules,omitempty"`
		DenyRules *[]struct {
			SourceKeyspace *struct {
				Collection *string `tfsdk:"collection" json:"collection,omitempty"`
				Scope      *string `tfsdk:"scope" json:"scope,omitempty"`
			} `tfsdk:"source_keyspace" json:"sourceKeyspace,omitempty"`
		} `tfsdk:"deny_rules" json:"denyRules,omitempty"`
	} `tfsdk:"explicit_mapping" json:"explicitMapping,omitempty"`
	Spec *struct {
		Bucket           *string `tfsdk:"bucket" json:"bucket,omitempty"`
		CompressionType  *string `tfsdk:"compression_type" json:"compressionType,omitempty"`
		FilterExpression *string `tfsdk:"filter_expression" json:"filterExpression,omitempty"`
		Paused           *bool   `tfsdk:"paused" json:"paused,omitempty"`
		RemoteBucket     *string `tfsdk:"remote_bucket" json:"remoteBucket,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CouchbaseComCouchbaseReplicationV2DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_couchbase_com_couchbase_replication_v2"
}

func (r *CouchbaseComCouchbaseReplicationV2DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "The CouchbaseReplication resource represents a Couchbase-to-Couchbase, XDCR replication stream from a source bucket to a destination bucket.  This provides off-site backup, migration, and disaster recovery.",
		MarkdownDescription: "The CouchbaseReplication resource represents a Couchbase-to-Couchbase, XDCR replication stream from a source bucket to a destination bucket.  This provides off-site backup, migration, and disaster recovery.",
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

			"explicit_mapping": schema.SingleNestedAttribute{
				Description:         "The explicit mappings to use for replication which are optional. For Scopes and Collection replication support we can specify a set of implicit and explicit mappings to use. If none is specified then it is assumed to be existing bucket level replication. https://docs.couchbase.com/server/current/learn/clusters-and-availability/xdcr-with-scopes-and-collections.html#explicit-mapping",
				MarkdownDescription: "The explicit mappings to use for replication which are optional. For Scopes and Collection replication support we can specify a set of implicit and explicit mappings to use. If none is specified then it is assumed to be existing bucket level replication. https://docs.couchbase.com/server/current/learn/clusters-and-availability/xdcr-with-scopes-and-collections.html#explicit-mapping",
				Attributes: map[string]schema.Attribute{
					"allow_rules": schema.ListNestedAttribute{
						Description:         "The list of explicit replications to carry out including any nested implicit replications: specifying a scope implicitly replicates all collections within it. There should be no duplicates, including more-specific duplicates, e.g. if you specify replication of a scope then you can only deny replication of collections within it.",
						MarkdownDescription: "The list of explicit replications to carry out including any nested implicit replications: specifying a scope implicitly replicates all collections within it. There should be no duplicates, including more-specific duplicates, e.g. if you specify replication of a scope then you can only deny replication of collections within it.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"source_keyspace": schema.SingleNestedAttribute{
									Description:         "The source keyspace: where to replicate from. Source and target must match whether they have a collection or not, i.e. you cannot replicate from a scope to a collection.",
									MarkdownDescription: "The source keyspace: where to replicate from. Source and target must match whether they have a collection or not, i.e. you cannot replicate from a scope to a collection.",
									Attributes: map[string]schema.Attribute{
										"collection": schema.StringAttribute{
											Description:         "The optional collection within the scope. May be empty to just work at scope level.",
											MarkdownDescription: "The optional collection within the scope. May be empty to just work at scope level.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"scope": schema.StringAttribute{
											Description:         "The scope to use.",
											MarkdownDescription: "The scope to use.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"target_keyspace": schema.SingleNestedAttribute{
									Description:         "The target keyspace: where to replicate to. Source and target must match whether they have a collection or not, i.e. you cannot replicate from a scope to a collection.",
									MarkdownDescription: "The target keyspace: where to replicate to. Source and target must match whether they have a collection or not, i.e. you cannot replicate from a scope to a collection.",
									Attributes: map[string]schema.Attribute{
										"collection": schema.StringAttribute{
											Description:         "The optional collection within the scope. May be empty to just work at scope level.",
											MarkdownDescription: "The optional collection within the scope. May be empty to just work at scope level.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"scope": schema.StringAttribute{
											Description:         "The scope to use.",
											MarkdownDescription: "The scope to use.",
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"deny_rules": schema.ListNestedAttribute{
						Description:         "The list of explicit replications to prevent including any nested implicit denials: specifying a scope implicitly denies all collections within it. There should be no duplicates, including more-specific duplicates, e.g. if you specify denial of replication of a scope then you can only specify replication of collections within it.",
						MarkdownDescription: "The list of explicit replications to prevent including any nested implicit denials: specifying a scope implicitly denies all collections within it. There should be no duplicates, including more-specific duplicates, e.g. if you specify denial of replication of a scope then you can only specify replication of collections within it.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"source_keyspace": schema.SingleNestedAttribute{
									Description:         "The source keyspace: where to block replication from.",
									MarkdownDescription: "The source keyspace: where to block replication from.",
									Attributes: map[string]schema.Attribute{
										"collection": schema.StringAttribute{
											Description:         "The optional collection within the scope. May be empty to just work at scope level.",
											MarkdownDescription: "The optional collection within the scope. May be empty to just work at scope level.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"scope": schema.StringAttribute{
											Description:         "The scope to use.",
											MarkdownDescription: "The scope to use.",
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "CouchbaseReplicationSpec allows configuration of an XDCR replication.",
				MarkdownDescription: "CouchbaseReplicationSpec allows configuration of an XDCR replication.",
				Attributes: map[string]schema.Attribute{
					"bucket": schema.StringAttribute{
						Description:         "Bucket is the source bucket to replicate from.  This refers to the Couchbase bucket name, not the resource name of the bucket.  A bucket with this name must be defined on this cluster.  Legal bucket names have a maximum length of 100 characters and may be composed of any character from 'a-z', 'A-Z', '0-9' and '-_%.'.",
						MarkdownDescription: "Bucket is the source bucket to replicate from.  This refers to the Couchbase bucket name, not the resource name of the bucket.  A bucket with this name must be defined on this cluster.  Legal bucket names have a maximum length of 100 characters and may be composed of any character from 'a-z', 'A-Z', '0-9' and '-_%.'.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"compression_type": schema.StringAttribute{
						Description:         "CompressionType is the type of compression to apply to the replication. When None, no compression will be applied to documents as they are transferred between clusters.  When Auto, Couchbase server will automatically compress documents as they are transferred to reduce bandwidth requirements. This field must be one of 'None' or 'Auto', defaulting to 'Auto'.",
						MarkdownDescription: "CompressionType is the type of compression to apply to the replication. When None, no compression will be applied to documents as they are transferred between clusters.  When Auto, Couchbase server will automatically compress documents as they are transferred to reduce bandwidth requirements. This field must be one of 'None' or 'Auto', defaulting to 'Auto'.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"filter_expression": schema.StringAttribute{
						Description:         "FilterExpression allows certain documents to be filtered out of the replication.",
						MarkdownDescription: "FilterExpression allows certain documents to be filtered out of the replication.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"paused": schema.BoolAttribute{
						Description:         "Paused allows a replication to be stopped and restarted without having to restart the replication from the beginning.",
						MarkdownDescription: "Paused allows a replication to be stopped and restarted without having to restart the replication from the beginning.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"remote_bucket": schema.StringAttribute{
						Description:         "RemoteBucket is the remote bucket name to synchronize to.  This refers to the Couchbase bucket name, not the resource name of the bucket.  Legal bucket names have a maximum length of 100 characters and may be composed of any character from 'a-z', 'A-Z', '0-9' and '-_%.'.",
						MarkdownDescription: "RemoteBucket is the remote bucket name to synchronize to.  This refers to the Couchbase bucket name, not the resource name of the bucket.  Legal bucket names have a maximum length of 100 characters and may be composed of any character from 'a-z', 'A-Z', '0-9' and '-_%.'.",
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

func (r *CouchbaseComCouchbaseReplicationV2DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *CouchbaseComCouchbaseReplicationV2DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_couchbase_com_couchbase_replication_v2")

	var data CouchbaseComCouchbaseReplicationV2DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "couchbase.com", Version: "v2", Resource: "CouchbaseReplication"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
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

	var readResponse CouchbaseComCouchbaseReplicationV2DataSourceData
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
	data.ApiVersion = pointer.String("couchbase.com/v2")
	data.Kind = pointer.String("CouchbaseReplication")
	data.Metadata = readResponse.Metadata
	data.ExplicitMapping = readResponse.ExplicitMapping
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
