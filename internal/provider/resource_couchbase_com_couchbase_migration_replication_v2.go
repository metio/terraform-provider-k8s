/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"regexp"

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

type CouchbaseComCouchbaseMigrationReplicationV2Resource struct{}

var (
	_ resource.Resource = (*CouchbaseComCouchbaseMigrationReplicationV2Resource)(nil)
)

type CouchbaseComCouchbaseMigrationReplicationV2TerraformModel struct {
	Id               types.Int64  `tfsdk:"id"`
	YAML             types.String `tfsdk:"yaml"`
	ApiVersion       types.String `tfsdk:"api_version"`
	Kind             types.String `tfsdk:"kind"`
	Metadata         types.Object `tfsdk:"metadata"`
	MigrationMapping types.Object `tfsdk:"migration_mapping"`
	Spec             types.Object `tfsdk:"spec"`
}

type CouchbaseComCouchbaseMigrationReplicationV2GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	MigrationMapping *struct {
		Mappings *[]struct {
			Filter *string `tfsdk:"filter" yaml:"filter,omitempty"`

			TargetKeyspace *struct {
				Collection *string `tfsdk:"collection" yaml:"collection,omitempty"`

				Scope *string `tfsdk:"scope" yaml:"scope,omitempty"`
			} `tfsdk:"target_keyspace" yaml:"targetKeyspace,omitempty"`
		} `tfsdk:"mappings" yaml:"mappings,omitempty"`
	} `tfsdk:"migration_mapping" yaml:"migrationMapping,omitempty"`

	Spec *struct {
		Bucket *string `tfsdk:"bucket" yaml:"bucket,omitempty"`

		CompressionType *string `tfsdk:"compression_type" yaml:"compressionType,omitempty"`

		FilterExpression *string `tfsdk:"filter_expression" yaml:"filterExpression,omitempty"`

		Paused *bool `tfsdk:"paused" yaml:"paused,omitempty"`

		RemoteBucket *string `tfsdk:"remote_bucket" yaml:"remoteBucket,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewCouchbaseComCouchbaseMigrationReplicationV2Resource() resource.Resource {
	return &CouchbaseComCouchbaseMigrationReplicationV2Resource{}
}

func (r *CouchbaseComCouchbaseMigrationReplicationV2Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_couchbase_com_couchbase_migration_replication_v2"
}

func (r *CouchbaseComCouchbaseMigrationReplicationV2Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "The CouchbaseScopeMigration resource represents the use of the special migration mapping within XDCR to take a filtered list from the default scope and collection of the source bucket, replicate it to named scopes and collections within the target bucket. The bucket-to-bucket replication cannot duplicate any used by the CouchbaseReplication resource, as these two types of replication are mutually exclusive between buckets. https://docs.couchbase.com/server/current/learn/clusters-and-availability/xdcr-with-scopes-and-collections.html#migration",
		MarkdownDescription: "The CouchbaseScopeMigration resource represents the use of the special migration mapping within XDCR to take a filtered list from the default scope and collection of the source bucket, replicate it to named scopes and collections within the target bucket. The bucket-to-bucket replication cannot duplicate any used by the CouchbaseReplication resource, as these two types of replication are mutually exclusive between buckets. https://docs.couchbase.com/server/current/learn/clusters-and-availability/xdcr-with-scopes-and-collections.html#migration",
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

			"migration_mapping": {
				Description:         "The migration mappings to use, should never be empty as that is just an implicit bucket-to-bucket replication then.",
				MarkdownDescription: "The migration mappings to use, should never be empty as that is just an implicit bucket-to-bucket replication then.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"mappings": {
						Description:         "The migration mappings to use, should never be empty as that is just an implicit bucket-to-bucket replication then.",
						MarkdownDescription: "The migration mappings to use, should never be empty as that is just an implicit bucket-to-bucket replication then.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"filter": {
								Description:         "A filter to select from the source default scope and collection. Defaults to select everything in the default scope and collection.",
								MarkdownDescription: "A filter to select from the source default scope and collection. Defaults to select everything in the default scope and collection.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"target_keyspace": {
								Description:         "The destination of our migration, must be a scope and collection.",
								MarkdownDescription: "The destination of our migration, must be a scope and collection.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"collection": {
										Description:         "The optional collection within the scope. May be empty to just work at scope level.",
										MarkdownDescription: "The optional collection within the scope. May be empty to just work at scope level.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.LengthAtLeast(1),

											stringvalidator.LengthAtMost(251),

											stringvalidator.RegexMatches(regexp.MustCompile(`^(_default|[a-zA-Z0-9\-][a-zA-Z0-9\-%_]{0,250})$`), ""),
										},
									},

									"scope": {
										Description:         "The scope to use.",
										MarkdownDescription: "The scope to use.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.LengthAtLeast(1),

											stringvalidator.LengthAtMost(251),

											stringvalidator.RegexMatches(regexp.MustCompile(`^(_default|[a-zA-Z0-9\-][a-zA-Z0-9\-%_]{0,250})$`), ""),
										},
									},
								}),

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},
				}),

				Required: true,
				Optional: false,
				Computed: false,
			},

			"spec": {
				Description:         "CouchbaseReplicationSpec allows configuration of an XDCR replication.",
				MarkdownDescription: "CouchbaseReplicationSpec allows configuration of an XDCR replication.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"bucket": {
						Description:         "Bucket is the source bucket to replicate from.  This refers to the Couchbase bucket name, not the resource name of the bucket.  A bucket with this name must be defined on this cluster.  Legal bucket names have a maximum length of 100 characters and may be composed of any character from 'a-z', 'A-Z', '0-9' and '-_%.'.",
						MarkdownDescription: "Bucket is the source bucket to replicate from.  This refers to the Couchbase bucket name, not the resource name of the bucket.  A bucket with this name must be defined on this cluster.  Legal bucket names have a maximum length of 100 characters and may be composed of any character from 'a-z', 'A-Z', '0-9' and '-_%.'.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.LengthAtMost(100),

							stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9-_%\.]{1,100}$`), ""),
						},
					},

					"compression_type": {
						Description:         "CompressionType is the type of compression to apply to the replication. When None, no compression will be applied to documents as they are transferred between clusters.  When Auto, Couchbase server will automatically compress documents as they are transferred to reduce bandwidth requirements. This field must be one of 'None' or 'Auto', defaulting to 'Auto'.",
						MarkdownDescription: "CompressionType is the type of compression to apply to the replication. When None, no compression will be applied to documents as they are transferred between clusters.  When Auto, Couchbase server will automatically compress documents as they are transferred to reduce bandwidth requirements. This field must be one of 'None' or 'Auto', defaulting to 'Auto'.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("None", "Auto"),
						},
					},

					"filter_expression": {
						Description:         "FilterExpression allows certain documents to be filtered out of the replication.",
						MarkdownDescription: "FilterExpression allows certain documents to be filtered out of the replication.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"paused": {
						Description:         "Paused allows a replication to be stopped and restarted without having to restart the replication from the beginning.",
						MarkdownDescription: "Paused allows a replication to be stopped and restarted without having to restart the replication from the beginning.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"remote_bucket": {
						Description:         "RemoteBucket is the remote bucket name to synchronize to.  This refers to the Couchbase bucket name, not the resource name of the bucket.  Legal bucket names have a maximum length of 100 characters and may be composed of any character from 'a-z', 'A-Z', '0-9' and '-_%.'.",
						MarkdownDescription: "RemoteBucket is the remote bucket name to synchronize to.  This refers to the Couchbase bucket name, not the resource name of the bucket.  Legal bucket names have a maximum length of 100 characters and may be composed of any character from 'a-z', 'A-Z', '0-9' and '-_%.'.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.LengthAtMost(100),

							stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9-_%\.]{1,100}$`), ""),
						},
					},
				}),

				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}, nil
}

func (r *CouchbaseComCouchbaseMigrationReplicationV2Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_couchbase_com_couchbase_migration_replication_v2")

	var state CouchbaseComCouchbaseMigrationReplicationV2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CouchbaseComCouchbaseMigrationReplicationV2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("couchbase.com/v2")
	goModel.Kind = utilities.Ptr("CouchbaseMigrationReplication")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CouchbaseComCouchbaseMigrationReplicationV2Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_couchbase_com_couchbase_migration_replication_v2")
	// NO-OP: All data is already in Terraform state
}

func (r *CouchbaseComCouchbaseMigrationReplicationV2Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_couchbase_com_couchbase_migration_replication_v2")

	var state CouchbaseComCouchbaseMigrationReplicationV2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CouchbaseComCouchbaseMigrationReplicationV2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("couchbase.com/v2")
	goModel.Kind = utilities.Ptr("CouchbaseMigrationReplication")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CouchbaseComCouchbaseMigrationReplicationV2Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_couchbase_com_couchbase_migration_replication_v2")
	// NO-OP: Terraform removes the state automatically for us
}
