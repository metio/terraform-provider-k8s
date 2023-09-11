/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package services_k8s_aws_v1alpha1

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
	_ datasource.DataSource              = &ServicesK8SAwsAdoptedResourceV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &ServicesK8SAwsAdoptedResourceV1Alpha1DataSource{}
)

func NewServicesK8SAwsAdoptedResourceV1Alpha1DataSource() datasource.DataSource {
	return &ServicesK8SAwsAdoptedResourceV1Alpha1DataSource{}
}

type ServicesK8SAwsAdoptedResourceV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type ServicesK8SAwsAdoptedResourceV1Alpha1DataSourceData struct {
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
		Aws *struct {
			AdditionalKeys *map[string]string `tfsdk:"additional_keys" json:"additionalKeys,omitempty"`
			Arn            *string            `tfsdk:"arn" json:"arn,omitempty"`
			NameOrID       *string            `tfsdk:"name_or_id" json:"nameOrID,omitempty"`
		} `tfsdk:"aws" json:"aws,omitempty"`
		Kubernetes *struct {
			Group    *string `tfsdk:"group" json:"group,omitempty"`
			Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
			Metadata *struct {
				Annotations     *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				GenerateName    *string            `tfsdk:"generate_name" json:"generateName,omitempty"`
				Labels          *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				Name            *string            `tfsdk:"name" json:"name,omitempty"`
				Namespace       *string            `tfsdk:"namespace" json:"namespace,omitempty"`
				OwnerReferences *[]struct {
					ApiVersion         *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					BlockOwnerDeletion *bool   `tfsdk:"block_owner_deletion" json:"blockOwnerDeletion,omitempty"`
					Controller         *bool   `tfsdk:"controller" json:"controller,omitempty"`
					Kind               *string `tfsdk:"kind" json:"kind,omitempty"`
					Name               *string `tfsdk:"name" json:"name,omitempty"`
					Uid                *string `tfsdk:"uid" json:"uid,omitempty"`
				} `tfsdk:"owner_references" json:"ownerReferences,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
		} `tfsdk:"kubernetes" json:"kubernetes,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ServicesK8SAwsAdoptedResourceV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_services_k8s_aws_adopted_resource_v1alpha1"
}

func (r *ServicesK8SAwsAdoptedResourceV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "AdoptedResource is the schema for the AdoptedResource API.",
		MarkdownDescription: "AdoptedResource is the schema for the AdoptedResource API.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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
				Description:         "AdoptedResourceSpec defines the desired state of the AdoptedResource.",
				MarkdownDescription: "AdoptedResourceSpec defines the desired state of the AdoptedResource.",
				Attributes: map[string]schema.Attribute{
					"aws": schema.SingleNestedAttribute{
						Description:         "AWSIdentifiers provide all unique ways to reference an AWS resource.",
						MarkdownDescription: "AWSIdentifiers provide all unique ways to reference an AWS resource.",
						Attributes: map[string]schema.Attribute{
							"additional_keys": schema.MapAttribute{
								Description:         "AdditionalKeys represents any additional arbitrary identifiers used when describing the target resource.",
								MarkdownDescription: "AdditionalKeys represents any additional arbitrary identifiers used when describing the target resource.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"arn": schema.StringAttribute{
								Description:         "ARN is the AWS Resource Name for the resource. It is a globally unique identifier.",
								MarkdownDescription: "ARN is the AWS Resource Name for the resource. It is a globally unique identifier.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name_or_id": schema.StringAttribute{
								Description:         "NameOrId is a user-supplied string identifier for the resource. It may or may not be globally unique, depending on the type of resource.",
								MarkdownDescription: "NameOrId is a user-supplied string identifier for the resource. It may or may not be globally unique, depending on the type of resource.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"kubernetes": schema.SingleNestedAttribute{
						Description:         "ResourceWithMetadata provides the values necessary to create a Kubernetes resource and override any of its metadata values.",
						MarkdownDescription: "ResourceWithMetadata provides the values necessary to create a Kubernetes resource and override any of its metadata values.",
						Attributes: map[string]schema.Attribute{
							"group": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"kind": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"metadata": schema.SingleNestedAttribute{
								Description:         "ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create. It is not possible to use 'metav1.ObjectMeta' inside spec, as the controller-gen automatically converts this to an arbitrary string-string map. https://github.com/kubernetes-sigs/controller-tools/issues/385  Active discussion about inclusion of this field in the spec is happening in this PR: https://github.com/kubernetes-sigs/controller-tools/pull/395  Until this is allowed, or if it never is, we will produce a subset of the object meta that contains only the fields which the user is allowed to modify in the metadata.",
								MarkdownDescription: "ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create. It is not possible to use 'metav1.ObjectMeta' inside spec, as the controller-gen automatically converts this to an arbitrary string-string map. https://github.com/kubernetes-sigs/controller-tools/issues/385  Active discussion about inclusion of this field in the spec is happening in this PR: https://github.com/kubernetes-sigs/controller-tools/pull/395  Until this is allowed, or if it never is, we will produce a subset of the object meta that contains only the fields which the user is allowed to modify in the metadata.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
										MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"generate_name": schema.StringAttribute{
										Description:         "GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server.  If this field is specified and the generated name exists, the server will NOT return a 409 - instead, it will either return 201 Created or 500 with Reason ServerTimeout indicating a unique name could not be found in the time allotted, and the client should retry (optionally after the time indicated in the Retry-After header).  Applied only if Name is not specified. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#idempotency",
										MarkdownDescription: "GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server.  If this field is specified and the generated name exists, the server will NOT return a 409 - instead, it will either return 201 Created or 500 with Reason ServerTimeout indicating a unique name could not be found in the time allotted, and the client should retry (optionally after the time indicated in the Retry-After header).  Applied only if Name is not specified. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#idempotency",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"labels": schema.MapAttribute{
										Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
										MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"name": schema.StringAttribute{
										Description:         "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
										MarkdownDescription: "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"namespace": schema.StringAttribute{
										Description:         "Namespace defines the space within each name must be unique. An empty namespace is equivalent to the 'default' namespace, but 'default' is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.  Must be a DNS_LABEL. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/namespaces",
										MarkdownDescription: "Namespace defines the space within each name must be unique. An empty namespace is equivalent to the 'default' namespace, but 'default' is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.  Must be a DNS_LABEL. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/namespaces",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"owner_references": schema.ListNestedAttribute{
										Description:         "List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller.",
										MarkdownDescription: "List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"api_version": schema.StringAttribute{
													Description:         "API version of the referent.",
													MarkdownDescription: "API version of the referent.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"block_owner_deletion": schema.BoolAttribute{
													Description:         "If true, AND if the owner has the 'foregroundDeletion' finalizer, then the owner cannot be deleted from the key-value store until this reference is removed. Defaults to false. To set this field, a user needs 'delete' permission of the owner, otherwise 422 (Unprocessable Entity) will be returned.",
													MarkdownDescription: "If true, AND if the owner has the 'foregroundDeletion' finalizer, then the owner cannot be deleted from the key-value store until this reference is removed. Defaults to false. To set this field, a user needs 'delete' permission of the owner, otherwise 422 (Unprocessable Entity) will be returned.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"controller": schema.BoolAttribute{
													Description:         "If true, this reference points to the managing controller.",
													MarkdownDescription: "If true, this reference points to the managing controller.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"kind": schema.StringAttribute{
													Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
													MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
													MarkdownDescription: "Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"uid": schema.StringAttribute{
													Description:         "UID of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#uids",
													MarkdownDescription: "UID of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#uids",
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
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *ServicesK8SAwsAdoptedResourceV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *ServicesK8SAwsAdoptedResourceV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_services_k8s_aws_adopted_resource_v1alpha1")

	var data ServicesK8SAwsAdoptedResourceV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "services.k8s.aws", Version: "v1alpha1", Resource: "adoptedresources"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse ServicesK8SAwsAdoptedResourceV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("services.k8s.aws/v1alpha1")
	data.Kind = pointer.String("AdoptedResource")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
