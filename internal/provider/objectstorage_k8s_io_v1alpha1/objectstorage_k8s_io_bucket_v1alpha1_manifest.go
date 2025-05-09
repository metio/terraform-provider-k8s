/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package objectstorage_k8s_io_v1alpha1

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
	_ datasource.DataSource = &ObjectstorageK8SIoBucketV1Alpha1Manifest{}
)

func NewObjectstorageK8SIoBucketV1Alpha1Manifest() datasource.DataSource {
	return &ObjectstorageK8SIoBucketV1Alpha1Manifest{}
}

type ObjectstorageK8SIoBucketV1Alpha1Manifest struct{}

type ObjectstorageK8SIoBucketV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		BucketClaim *struct {
			ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
			Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
			Name            *string `tfsdk:"name" json:"name,omitempty"`
			Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
			ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
			Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
		} `tfsdk:"bucket_claim" json:"bucketClaim,omitempty"`
		BucketClassName  *string            `tfsdk:"bucket_class_name" json:"bucketClassName,omitempty"`
		DeletionPolicy   *string            `tfsdk:"deletion_policy" json:"deletionPolicy,omitempty"`
		DriverName       *string            `tfsdk:"driver_name" json:"driverName,omitempty"`
		ExistingBucketID *string            `tfsdk:"existing_bucket_id" json:"existingBucketID,omitempty"`
		Parameters       *map[string]string `tfsdk:"parameters" json:"parameters,omitempty"`
		Protocols        *[]string          `tfsdk:"protocols" json:"protocols,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ObjectstorageK8SIoBucketV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_objectstorage_k8s_io_bucket_v1alpha1_manifest"
}

func (r *ObjectstorageK8SIoBucketV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"bucket_claim": schema.SingleNestedAttribute{
						Description:         "Name of the BucketClaim that resulted in the creation of this Bucket In case the Bucket object was created manually, then this should refer to the BucketClaim with which this Bucket should be bound",
						MarkdownDescription: "Name of the BucketClaim that resulted in the creation of this Bucket In case the Bucket object was created manually, then this should refer to the BucketClaim with which this Bucket should be bound",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "API version of the referent.",
								MarkdownDescription: "API version of the referent.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"field_path": schema.StringAttribute{
								Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.",
								MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resource_version": schema.StringAttribute{
								Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"uid": schema.StringAttribute{
								Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"bucket_class_name": schema.StringAttribute{
						Description:         "Name of the BucketClass specified in the BucketRequest",
						MarkdownDescription: "Name of the BucketClass specified in the BucketRequest",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"deletion_policy": schema.StringAttribute{
						Description:         "DeletionPolicy is used to specify how COSI should handle deletion of this bucket. There are 2 possible values: - Retain: Indicates that the bucket should not be deleted from the OSP (default) - Delete: Indicates that the bucket should be deleted from the OSP once all the workloads accessing this bucket are done",
						MarkdownDescription: "DeletionPolicy is used to specify how COSI should handle deletion of this bucket. There are 2 possible values: - Retain: Indicates that the bucket should not be deleted from the OSP (default) - Delete: Indicates that the bucket should be deleted from the OSP once all the workloads accessing this bucket are done",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"driver_name": schema.StringAttribute{
						Description:         "DriverName is the name of driver associated with this bucket",
						MarkdownDescription: "DriverName is the name of driver associated with this bucket",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"existing_bucket_id": schema.StringAttribute{
						Description:         "ExistingBucketID is the unique id of the bucket in the OSP. This field should be used to specify a bucket that has been created outside of COSI. This field will be empty when the Bucket is dynamically provisioned by COSI.",
						MarkdownDescription: "ExistingBucketID is the unique id of the bucket in the OSP. This field should be used to specify a bucket that has been created outside of COSI. This field will be empty when the Bucket is dynamically provisioned by COSI.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"parameters": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"protocols": schema.ListAttribute{
						Description:         "Protocols are the set of data APIs this bucket is expected to support. The possible values for protocol are: - S3: Indicates Amazon S3 protocol - Azure: Indicates Microsoft Azure BlobStore protocol - GCS: Indicates Google Cloud Storage protocol",
						MarkdownDescription: "Protocols are the set of data APIs this bucket is expected to support. The possible values for protocol are: - S3: Indicates Amazon S3 protocol - Azure: Indicates Microsoft Azure BlobStore protocol - GCS: Indicates Google Cloud Storage protocol",
						ElementType:         types.StringType,
						Required:            true,
						Optional:            false,
						Computed:            false,
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *ObjectstorageK8SIoBucketV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_objectstorage_k8s_io_bucket_v1alpha1_manifest")

	var model ObjectstorageK8SIoBucketV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("objectstorage.k8s.io/v1alpha1")
	model.Kind = pointer.String("Bucket")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
