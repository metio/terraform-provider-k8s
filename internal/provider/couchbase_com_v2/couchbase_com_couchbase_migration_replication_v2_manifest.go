/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package couchbase_com_v2

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &CouchbaseComCouchbaseMigrationReplicationV2Manifest{}
)

func NewCouchbaseComCouchbaseMigrationReplicationV2Manifest() datasource.DataSource {
	return &CouchbaseComCouchbaseMigrationReplicationV2Manifest{}
}

type CouchbaseComCouchbaseMigrationReplicationV2Manifest struct{}

type CouchbaseComCouchbaseMigrationReplicationV2ManifestData struct {
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

	MigrationMapping *struct {
		Mappings *[]struct {
			Filter         *string `tfsdk:"filter" json:"filter,omitempty"`
			TargetKeyspace *struct {
				Collection *string `tfsdk:"collection" json:"collection,omitempty"`
				Scope      *string `tfsdk:"scope" json:"scope,omitempty"`
			} `tfsdk:"target_keyspace" json:"targetKeyspace,omitempty"`
		} `tfsdk:"mappings" json:"mappings,omitempty"`
	} `tfsdk:"migration_mapping" json:"migrationMapping,omitempty"`
	Spec *struct {
		Bucket           *string `tfsdk:"bucket" json:"bucket,omitempty"`
		CompressionType  *string `tfsdk:"compression_type" json:"compressionType,omitempty"`
		FilterExpression *string `tfsdk:"filter_expression" json:"filterExpression,omitempty"`
		Paused           *bool   `tfsdk:"paused" json:"paused,omitempty"`
		RemoteBucket     *string `tfsdk:"remote_bucket" json:"remoteBucket,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CouchbaseComCouchbaseMigrationReplicationV2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_couchbase_com_couchbase_migration_replication_v2_manifest"
}

func (r *CouchbaseComCouchbaseMigrationReplicationV2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "The CouchbaseScopeMigration resource represents the use of the special migration mapping within XDCR to take a filtered list from the default scope and collection of the source bucket, replicate it to named scopes and collections within the target bucket. The bucket-to-bucket replication cannot duplicate any used by the CouchbaseReplication resource, as these two types of replication are mutually exclusive between buckets. https://docs.couchbase.com/server/current/learn/clusters-and-availability/xdcr-with-scopes-and-collections.html#migration",
		MarkdownDescription: "The CouchbaseScopeMigration resource represents the use of the special migration mapping within XDCR to take a filtered list from the default scope and collection of the source bucket, replicate it to named scopes and collections within the target bucket. The bucket-to-bucket replication cannot duplicate any used by the CouchbaseReplication resource, as these two types of replication are mutually exclusive between buckets. https://docs.couchbase.com/server/current/learn/clusters-and-availability/xdcr-with-scopes-and-collections.html#migration",
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

			"migration_mapping": schema.SingleNestedAttribute{
				Description:         "The migration mappings to use, should never be empty as that is just an implicit bucket-to-bucket replication then.",
				MarkdownDescription: "The migration mappings to use, should never be empty as that is just an implicit bucket-to-bucket replication then.",
				Attributes: map[string]schema.Attribute{
					"mappings": schema.ListNestedAttribute{
						Description:         "The migration mappings to use, should never be empty as that is just an implicit bucket-to-bucket replication then.",
						MarkdownDescription: "The migration mappings to use, should never be empty as that is just an implicit bucket-to-bucket replication then.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"filter": schema.StringAttribute{
									Description:         "A filter to select from the source default scope and collection. Defaults to select everything in the default scope and collection.",
									MarkdownDescription: "A filter to select from the source default scope and collection. Defaults to select everything in the default scope and collection.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"target_keyspace": schema.SingleNestedAttribute{
									Description:         "The destination of our migration, must be a scope and collection.",
									MarkdownDescription: "The destination of our migration, must be a scope and collection.",
									Attributes: map[string]schema.Attribute{
										"collection": schema.StringAttribute{
											Description:         "The optional collection within the scope. May be empty to just work at scope level.",
											MarkdownDescription: "The optional collection within the scope. May be empty to just work at scope level.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
												stringvalidator.LengthAtMost(251),
												stringvalidator.RegexMatches(regexp.MustCompile(`^(_default|[a-zA-Z0-9\-][a-zA-Z0-9\-%_]{0,250})$`), ""),
											},
										},

										"scope": schema.StringAttribute{
											Description:         "The scope to use.",
											MarkdownDescription: "The scope to use.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
												stringvalidator.LengthAtMost(251),
												stringvalidator.RegexMatches(regexp.MustCompile(`^(_default|[a-zA-Z0-9\-][a-zA-Z0-9\-%_]{0,250})$`), ""),
											},
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},
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

			"spec": schema.SingleNestedAttribute{
				Description:         "CouchbaseReplicationSpec allows configuration of an XDCR replication.",
				MarkdownDescription: "CouchbaseReplicationSpec allows configuration of an XDCR replication.",
				Attributes: map[string]schema.Attribute{
					"bucket": schema.StringAttribute{
						Description:         "Bucket is the source bucket to replicate from.  This refers to the Couchbase bucket name, not the resource name of the bucket.  A bucket with this name must be defined on this cluster.  Legal bucket names have a maximum length of 100 characters and may be composed of any character from 'a-z', 'A-Z', '0-9' and '-_%.'.",
						MarkdownDescription: "Bucket is the source bucket to replicate from.  This refers to the Couchbase bucket name, not the resource name of the bucket.  A bucket with this name must be defined on this cluster.  Legal bucket names have a maximum length of 100 characters and may be composed of any character from 'a-z', 'A-Z', '0-9' and '-_%.'.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(100),
							stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9-_%\.]{1,100}$`), ""),
						},
					},

					"compression_type": schema.StringAttribute{
						Description:         "CompressionType is the type of compression to apply to the replication. When None, no compression will be applied to documents as they are transferred between clusters.  When Auto, Couchbase server will automatically compress documents as they are transferred to reduce bandwidth requirements. This field must be one of 'None' or 'Auto', defaulting to 'Auto'.",
						MarkdownDescription: "CompressionType is the type of compression to apply to the replication. When None, no compression will be applied to documents as they are transferred between clusters.  When Auto, Couchbase server will automatically compress documents as they are transferred to reduce bandwidth requirements. This field must be one of 'None' or 'Auto', defaulting to 'Auto'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("None", "Auto"),
						},
					},

					"filter_expression": schema.StringAttribute{
						Description:         "FilterExpression allows certain documents to be filtered out of the replication.",
						MarkdownDescription: "FilterExpression allows certain documents to be filtered out of the replication.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"paused": schema.BoolAttribute{
						Description:         "Paused allows a replication to be stopped and restarted without having to restart the replication from the beginning.",
						MarkdownDescription: "Paused allows a replication to be stopped and restarted without having to restart the replication from the beginning.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"remote_bucket": schema.StringAttribute{
						Description:         "RemoteBucket is the remote bucket name to synchronize to.  This refers to the Couchbase bucket name, not the resource name of the bucket.  Legal bucket names have a maximum length of 100 characters and may be composed of any character from 'a-z', 'A-Z', '0-9' and '-_%.'.",
						MarkdownDescription: "RemoteBucket is the remote bucket name to synchronize to.  This refers to the Couchbase bucket name, not the resource name of the bucket.  Legal bucket names have a maximum length of 100 characters and may be composed of any character from 'a-z', 'A-Z', '0-9' and '-_%.'.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(100),
							stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9-_%\.]{1,100}$`), ""),
						},
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *CouchbaseComCouchbaseMigrationReplicationV2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_couchbase_com_couchbase_migration_replication_v2_manifest")

	var model CouchbaseComCouchbaseMigrationReplicationV2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("couchbase.com/v2")
	model.Kind = pointer.String("CouchbaseMigrationReplication")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
