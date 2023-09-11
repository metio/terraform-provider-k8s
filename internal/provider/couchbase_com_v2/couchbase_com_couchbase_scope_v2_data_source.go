/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package couchbase_com_v2

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
	_ datasource.DataSource              = &CouchbaseComCouchbaseScopeV2DataSource{}
	_ datasource.DataSourceWithConfigure = &CouchbaseComCouchbaseScopeV2DataSource{}
)

func NewCouchbaseComCouchbaseScopeV2DataSource() datasource.DataSource {
	return &CouchbaseComCouchbaseScopeV2DataSource{}
}

type CouchbaseComCouchbaseScopeV2DataSource struct {
	kubernetesClient dynamic.Interface
}

type CouchbaseComCouchbaseScopeV2DataSourceData struct {
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
		Collections *struct {
			Managed                   *bool `tfsdk:"managed" json:"managed,omitempty"`
			PreserveDefaultCollection *bool `tfsdk:"preserve_default_collection" json:"preserveDefaultCollection,omitempty"`
			Resources                 *[]struct {
				Kind *string `tfsdk:"kind" json:"kind,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			Selector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"selector" json:"selector,omitempty"`
		} `tfsdk:"collections" json:"collections,omitempty"`
		DefaultScope *bool   `tfsdk:"default_scope" json:"defaultScope,omitempty"`
		Name         *string `tfsdk:"name" json:"name,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CouchbaseComCouchbaseScopeV2DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_couchbase_com_couchbase_scope_v2"
}

func (r *CouchbaseComCouchbaseScopeV2DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CouchbaseScope represents a logical unit of data storage that sits between buckets and collections e.g. a bucket may contain multiple scopes, and a scope may contain multiple collections.  At present, scopes are not nested, so provide only a single level of abstraction.  Scopes provide a coarser grained basis for role-based access control (RBAC) and cross-datacenter replication (XDCR) than collections, but finer that buckets. In order to be considered by the Operator, a scope must be referenced by either a 'CouchbaseBucket' or 'CouchbaseEphemeralBucket' resource.",
		MarkdownDescription: "CouchbaseScope represents a logical unit of data storage that sits between buckets and collections e.g. a bucket may contain multiple scopes, and a scope may contain multiple collections.  At present, scopes are not nested, so provide only a single level of abstraction.  Scopes provide a coarser grained basis for role-based access control (RBAC) and cross-datacenter replication (XDCR) than collections, but finer that buckets. In order to be considered by the Operator, a scope must be referenced by either a 'CouchbaseBucket' or 'CouchbaseEphemeralBucket' resource.",
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
				Description:         "Spec defines the desired state of the resource.",
				MarkdownDescription: "Spec defines the desired state of the resource.",
				Attributes: map[string]schema.Attribute{
					"collections": schema.SingleNestedAttribute{
						Description:         "Collections defines how to collate collections included in this scope or scope group. Any of the provided methods may be used to collate a set of collections to manage.  Collated collections must have unique names, otherwise it is considered ambiguous, and an error condition.",
						MarkdownDescription: "Collections defines how to collate collections included in this scope or scope group. Any of the provided methods may be used to collate a set of collections to manage.  Collated collections must have unique names, otherwise it is considered ambiguous, and an error condition.",
						Attributes: map[string]schema.Attribute{
							"managed": schema.BoolAttribute{
								Description:         "Managed indicates whether collections within this scope are managed. If not then you can dynamically create and delete collections with the Couchbase UI or SDKs.",
								MarkdownDescription: "Managed indicates whether collections within this scope are managed. If not then you can dynamically create and delete collections with the Couchbase UI or SDKs.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"preserve_default_collection": schema.BoolAttribute{
								Description:         "PreserveDefaultCollection indicates whether the Operator should manage the default collection within the default scope.  The default collection can be deleted, but can not be recreated by Couchbase Server.  By setting this field to 'true', the Operator will implicitly manage the default collection within the default scope.  The default collection cannot be modified and will have no document time-to-live (TTL).  When set to 'false', the operator will not manage the default collection, which will be deleted and cannot be used or recreated.",
								MarkdownDescription: "PreserveDefaultCollection indicates whether the Operator should manage the default collection within the default scope.  The default collection can be deleted, but can not be recreated by Couchbase Server.  By setting this field to 'true', the Operator will implicitly manage the default collection within the default scope.  The default collection cannot be modified and will have no document time-to-live (TTL).  When set to 'false', the operator will not manage the default collection, which will be deleted and cannot be used or recreated.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"resources": schema.ListNestedAttribute{
								Description:         "Resources is an explicit list of named resources that will be considered for inclusion in this scope or scopes.  If a resource reference doesn't match a resource, then no error conditions are raised due to undefined resource creation ordering and eventual consistency.",
								MarkdownDescription: "Resources is an explicit list of named resources that will be considered for inclusion in this scope or scopes.  If a resource reference doesn't match a resource, then no error conditions are raised due to undefined resource creation ordering and eventual consistency.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"kind": schema.StringAttribute{
											Description:         "Kind indicates the kind of resource that is being referenced.  A scope can only reference 'CouchbaseCollection' and 'CouchbaseCollectionGroup' resource kinds.  This field defaults to 'CouchbaseCollection' if not specified.",
											MarkdownDescription: "Kind indicates the kind of resource that is being referenced.  A scope can only reference 'CouchbaseCollection' and 'CouchbaseCollectionGroup' resource kinds.  This field defaults to 'CouchbaseCollection' if not specified.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"name": schema.StringAttribute{
											Description:         "Name is the name of the Kubernetes resource name that is being referenced. Legal collection names have a maximum length of 251 characters and may be composed of any character from 'a-z', 'A-Z', '0-9' and '_-%'.",
											MarkdownDescription: "Name is the name of the Kubernetes resource name that is being referenced. Legal collection names have a maximum length of 251 characters and may be composed of any character from 'a-z', 'A-Z', '0-9' and '_-%'.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"selector": schema.SingleNestedAttribute{
								Description:         "Selector allows resources to be implicitly considered for inclusion in this scope or scopes.  More info: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#labelselector-v1-meta",
								MarkdownDescription: "Selector allows resources to be implicitly considered for inclusion in this scope or scopes.  More info: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#labelselector-v1-meta",
								Attributes: map[string]schema.Attribute{
									"match_expressions": schema.ListNestedAttribute{
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "key is the label key that the selector applies to.",
													MarkdownDescription: "key is the label key that the selector applies to.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"operator": schema.StringAttribute{
													Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
													MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"values": schema.ListAttribute{
													Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
													MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"match_labels": schema.MapAttribute{
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										ElementType:         types.StringType,
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
						Required: false,
						Optional: false,
						Computed: true,
					},

					"default_scope": schema.BoolAttribute{
						Description:         "DefaultScope indicates whether this resource represents the default scope for a bucket.  When set to 'true', this allows the user to refer to and manage collections within the default scope.  When not defined, the Operator will implicitly manage the default scope as the default scope can not be deleted from Couchbase Server.  The Operator defined default scope will also have the 'persistDefaultCollection' flag set to 'true'.  Only one default scope is permitted to be contained in a bucket.",
						MarkdownDescription: "DefaultScope indicates whether this resource represents the default scope for a bucket.  When set to 'true', this allows the user to refer to and manage collections within the default scope.  When not defined, the Operator will implicitly manage the default scope as the default scope can not be deleted from Couchbase Server.  The Operator defined default scope will also have the 'persistDefaultCollection' flag set to 'true'.  Only one default scope is permitted to be contained in a bucket.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"name": schema.StringAttribute{
						Description:         "Name specifies the name of the scope.  By default, the metadata.name is used to define the scope name, however, due to the limited character set, this field can be used to override the default and provide the full functionality. Additionally the 'metadata.name' field is a DNS label, and thus limited to 63 characters, this field must be used if the name is longer than this limit. Scope names must be 1-251 characters in length, contain only [a-zA-Z0-9_-%] and not start with either _ or %.",
						MarkdownDescription: "Name specifies the name of the scope.  By default, the metadata.name is used to define the scope name, however, due to the limited character set, this field can be used to override the default and provide the full functionality. Additionally the 'metadata.name' field is a DNS label, and thus limited to 63 characters, this field must be used if the name is longer than this limit. Scope names must be 1-251 characters in length, contain only [a-zA-Z0-9_-%] and not start with either _ or %.",
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

func (r *CouchbaseComCouchbaseScopeV2DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *CouchbaseComCouchbaseScopeV2DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_couchbase_com_couchbase_scope_v2")

	var data CouchbaseComCouchbaseScopeV2DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "couchbase.com", Version: "v2", Resource: "couchbasescopes"}).
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

	var readResponse CouchbaseComCouchbaseScopeV2DataSourceData
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
	data.Kind = pointer.String("CouchbaseScope")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
