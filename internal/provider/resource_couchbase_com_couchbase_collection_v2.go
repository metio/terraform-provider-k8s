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

type CouchbaseComCouchbaseCollectionV2Resource struct{}

var (
	_ resource.Resource = (*CouchbaseComCouchbaseCollectionV2Resource)(nil)
)

type CouchbaseComCouchbaseCollectionV2TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type CouchbaseComCouchbaseCollectionV2GoModel struct {
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

	Spec *struct {
		MaxTTL *string `tfsdk:"max_ttl" yaml:"maxTTL,omitempty"`

		Name *string `tfsdk:"name" yaml:"name,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewCouchbaseComCouchbaseCollectionV2Resource() resource.Resource {
	return &CouchbaseComCouchbaseCollectionV2Resource{}
}

func (r *CouchbaseComCouchbaseCollectionV2Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_couchbase_com_couchbase_collection_v2"
}

func (r *CouchbaseComCouchbaseCollectionV2Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "CouchbaseCollection represent the finest grained size of data storage in Couchbase. Collections contain all documents and indexes in the system.  Collections also form the finest grain basis for role-based access control (RBAC) and cross-datacenter replication (XDCR).  In order to be considered by the Operator, every collection must be referenced by a 'CouchbaseScope' or 'CouchbaseScopeGroup' resource.",
		MarkdownDescription: "CouchbaseCollection represent the finest grained size of data storage in Couchbase. Collections contain all documents and indexes in the system.  Collections also form the finest grain basis for role-based access control (RBAC) and cross-datacenter replication (XDCR).  In order to be considered by the Operator, every collection must be referenced by a 'CouchbaseScope' or 'CouchbaseScopeGroup' resource.",
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

			"spec": {
				Description:         "Spec defines the desired state of the resource.",
				MarkdownDescription: "Spec defines the desired state of the resource.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"max_ttl": {
						Description:         "MaxTTL defines how long a document is permitted to exist for, without modification, until it is automatically deleted.  This field takes precedence over any TTL defined at the bucket level.  This is a default, and maximum time-to-live and may be set to a lower value by the client.  If the client specifies a higher value, then it is truncated to the maximum durability.  Documents are removed by Couchbase, after they have expired, when either accessed, the expiry pager is run, or the bucket is compacted.  When set to 0, then documents are not expired by default.  This field must be a duration in the range 0-2147483648s, defaulting to 0.  More info: https://golang.org/pkg/time/#ParseDuration",
						MarkdownDescription: "MaxTTL defines how long a document is permitted to exist for, without modification, until it is automatically deleted.  This field takes precedence over any TTL defined at the bucket level.  This is a default, and maximum time-to-live and may be set to a lower value by the client.  If the client specifies a higher value, then it is truncated to the maximum durability.  Documents are removed by Couchbase, after they have expired, when either accessed, the expiry pager is run, or the bucket is compacted.  When set to 0, then documents are not expired by default.  This field must be a duration in the range 0-2147483648s, defaulting to 0.  More info: https://golang.org/pkg/time/#ParseDuration",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"name": {
						Description:         "Name specifies the name of the collection.  By default, the metadata.name is used to define the collection name, however, due to the limited character set, this field can be used to override the default and provide the full functionality. Additionally the 'metadata.name' field is a DNS label, and thus limited to 63 characters, this field must be used if the name is longer than this limit. Collection names must be 1-251 characters in length, contain only [a-zA-Z0-9_-%] and not start with either _ or %.",
						MarkdownDescription: "Name specifies the name of the collection.  By default, the metadata.name is used to define the collection name, however, due to the limited character set, this field can be used to override the default and provide the full functionality. Additionally the 'metadata.name' field is a DNS label, and thus limited to 63 characters, this field must be used if the name is longer than this limit. Collection names must be 1-251 characters in length, contain only [a-zA-Z0-9_-%] and not start with either _ or %.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.LengthAtLeast(1),

							stringvalidator.LengthAtMost(251),

							stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9\-][a-zA-Z0-9\-%_]{0,250}$`), ""),
						},
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *CouchbaseComCouchbaseCollectionV2Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_couchbase_com_couchbase_collection_v2")

	var state CouchbaseComCouchbaseCollectionV2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CouchbaseComCouchbaseCollectionV2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("couchbase.com/v2")
	goModel.Kind = utilities.Ptr("CouchbaseCollection")

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

func (r *CouchbaseComCouchbaseCollectionV2Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_couchbase_com_couchbase_collection_v2")
	// NO-OP: All data is already in Terraform state
}

func (r *CouchbaseComCouchbaseCollectionV2Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_couchbase_com_couchbase_collection_v2")

	var state CouchbaseComCouchbaseCollectionV2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CouchbaseComCouchbaseCollectionV2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("couchbase.com/v2")
	goModel.Kind = utilities.Ptr("CouchbaseCollection")

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

func (r *CouchbaseComCouchbaseCollectionV2Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_couchbase_com_couchbase_collection_v2")
	// NO-OP: Terraform removes the state automatically for us
}
