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

type CouchbaseComCouchbaseReplicationV2Resource struct{}

var (
	_ resource.Resource = (*CouchbaseComCouchbaseReplicationV2Resource)(nil)
)

type CouchbaseComCouchbaseReplicationV2TerraformModel struct {
	Id              types.Int64  `tfsdk:"id"`
	YAML            types.String `tfsdk:"yaml"`
	ApiVersion      types.String `tfsdk:"api_version"`
	Kind            types.String `tfsdk:"kind"`
	Metadata        types.Object `tfsdk:"metadata"`
	Spec            types.Object `tfsdk:"spec"`
	ExplicitMapping types.Object `tfsdk:"explicit_mapping"`
}

type CouchbaseComCouchbaseReplicationV2GoModel struct {
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
		Bucket *string `tfsdk:"bucket" yaml:"bucket,omitempty"`

		CompressionType *string `tfsdk:"compression_type" yaml:"compressionType,omitempty"`

		FilterExpression *string `tfsdk:"filter_expression" yaml:"filterExpression,omitempty"`

		Paused *bool `tfsdk:"paused" yaml:"paused,omitempty"`

		RemoteBucket *string `tfsdk:"remote_bucket" yaml:"remoteBucket,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`

	ExplicitMapping *struct {
		AllowRules *[]struct {
			SourceKeyspace *struct {
				Collection *string `tfsdk:"collection" yaml:"collection,omitempty"`

				Scope *string `tfsdk:"scope" yaml:"scope,omitempty"`
			} `tfsdk:"source_keyspace" yaml:"sourceKeyspace,omitempty"`

			TargetKeyspace *struct {
				Collection *string `tfsdk:"collection" yaml:"collection,omitempty"`

				Scope *string `tfsdk:"scope" yaml:"scope,omitempty"`
			} `tfsdk:"target_keyspace" yaml:"targetKeyspace,omitempty"`
		} `tfsdk:"allow_rules" yaml:"allowRules,omitempty"`

		DenyRules *[]struct {
			SourceKeyspace *struct {
				Collection *string `tfsdk:"collection" yaml:"collection,omitempty"`

				Scope *string `tfsdk:"scope" yaml:"scope,omitempty"`
			} `tfsdk:"source_keyspace" yaml:"sourceKeyspace,omitempty"`
		} `tfsdk:"deny_rules" yaml:"denyRules,omitempty"`
	} `tfsdk:"explicit_mapping" yaml:"explicitMapping,omitempty"`
}

func NewCouchbaseComCouchbaseReplicationV2Resource() resource.Resource {
	return &CouchbaseComCouchbaseReplicationV2Resource{}
}

func (r *CouchbaseComCouchbaseReplicationV2Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_couchbase_com_couchbase_replication_v2"
}

func (r *CouchbaseComCouchbaseReplicationV2Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "The CouchbaseReplication resource represents a Couchbase-to-Couchbase, XDCR replication stream from a source bucket to a destination bucket.  This provides off-site backup, migration, and disaster recovery.",
		MarkdownDescription: "The CouchbaseReplication resource represents a Couchbase-to-Couchbase, XDCR replication stream from a source bucket to a destination bucket.  This provides off-site backup, migration, and disaster recovery.",
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
					},

					"compression_type": {
						Description:         "CompressionType is the type of compression to apply to the replication. When None, no compression will be applied to documents as they are transferred between clusters.  When Auto, Couchbase server will automatically compress documents as they are transferred to reduce bandwidth requirements. This field must be one of 'None' or 'Auto', defaulting to 'Auto'.",
						MarkdownDescription: "CompressionType is the type of compression to apply to the replication. When None, no compression will be applied to documents as they are transferred between clusters.  When Auto, Couchbase server will automatically compress documents as they are transferred to reduce bandwidth requirements. This field must be one of 'None' or 'Auto', defaulting to 'Auto'.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
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
					},
				}),

				Required: true,
				Optional: false,
				Computed: false,
			},

			"explicit_mapping": {
				Description:         "The explicit mappings to use for replication which are optional. For Scopes and Collection replication support we can specify a set of implicit and explicit mappings to use. If none is specified then it is assumed to be existing bucket level replication. https://docs.couchbase.com/server/current/learn/clusters-and-availability/xdcr-with-scopes-and-collections.html#explicit-mapping",
				MarkdownDescription: "The explicit mappings to use for replication which are optional. For Scopes and Collection replication support we can specify a set of implicit and explicit mappings to use. If none is specified then it is assumed to be existing bucket level replication. https://docs.couchbase.com/server/current/learn/clusters-and-availability/xdcr-with-scopes-and-collections.html#explicit-mapping",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"allow_rules": {
						Description:         "The list of explicit replications to carry out including any nested implicit replications: specifying a scope implicitly replicates all collections within it. There should be no duplicates, including more-specific duplicates, e.g. if you specify replication of a scope then you can only deny replication of collections within it.",
						MarkdownDescription: "The list of explicit replications to carry out including any nested implicit replications: specifying a scope implicitly replicates all collections within it. There should be no duplicates, including more-specific duplicates, e.g. if you specify replication of a scope then you can only deny replication of collections within it.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"source_keyspace": {
								Description:         "The source keyspace: where to replicate from. Source and target must match whether they have a collection or not, i.e. you cannot replicate from a scope to a collection.",
								MarkdownDescription: "The source keyspace: where to replicate from. Source and target must match whether they have a collection or not, i.e. you cannot replicate from a scope to a collection.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"collection": {
										Description:         "The optional collection within the scope. May be empty to just work at scope level.",
										MarkdownDescription: "The optional collection within the scope. May be empty to just work at scope level.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"scope": {
										Description:         "The scope to use.",
										MarkdownDescription: "The scope to use.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: true,
								Optional: false,
								Computed: false,
							},

							"target_keyspace": {
								Description:         "The target keyspace: where to replicate to. Source and target must match whether they have a collection or not, i.e. you cannot replicate from a scope to a collection.",
								MarkdownDescription: "The target keyspace: where to replicate to. Source and target must match whether they have a collection or not, i.e. you cannot replicate from a scope to a collection.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"collection": {
										Description:         "The optional collection within the scope. May be empty to just work at scope level.",
										MarkdownDescription: "The optional collection within the scope. May be empty to just work at scope level.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"scope": {
										Description:         "The scope to use.",
										MarkdownDescription: "The scope to use.",

										Type: types.StringType,

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

						Required: false,
						Optional: true,
						Computed: false,
					},

					"deny_rules": {
						Description:         "The list of explicit replications to prevent including any nested implicit denials: specifying a scope implicitly denies all collections within it. There should be no duplicates, including more-specific duplicates, e.g. if you specify denial of replication of a scope then you can only specify replication of collections within it.",
						MarkdownDescription: "The list of explicit replications to prevent including any nested implicit denials: specifying a scope implicitly denies all collections within it. There should be no duplicates, including more-specific duplicates, e.g. if you specify denial of replication of a scope then you can only specify replication of collections within it.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"source_keyspace": {
								Description:         "The source keyspace: where to block replication from.",
								MarkdownDescription: "The source keyspace: where to block replication from.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"collection": {
										Description:         "The optional collection within the scope. May be empty to just work at scope level.",
										MarkdownDescription: "The optional collection within the scope. May be empty to just work at scope level.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"scope": {
										Description:         "The scope to use.",
										MarkdownDescription: "The scope to use.",

										Type: types.StringType,

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

func (r *CouchbaseComCouchbaseReplicationV2Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_couchbase_com_couchbase_replication_v2")

	var state CouchbaseComCouchbaseReplicationV2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CouchbaseComCouchbaseReplicationV2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("couchbase.com/v2")
	goModel.Kind = utilities.Ptr("CouchbaseReplication")

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

func (r *CouchbaseComCouchbaseReplicationV2Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_couchbase_com_couchbase_replication_v2")
	// NO-OP: All data is already in Terraform state
}

func (r *CouchbaseComCouchbaseReplicationV2Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_couchbase_com_couchbase_replication_v2")

	var state CouchbaseComCouchbaseReplicationV2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CouchbaseComCouchbaseReplicationV2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("couchbase.com/v2")
	goModel.Kind = utilities.Ptr("CouchbaseReplication")

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

func (r *CouchbaseComCouchbaseReplicationV2Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_couchbase_com_couchbase_replication_v2")
	// NO-OP: Terraform removes the state automatically for us
}
