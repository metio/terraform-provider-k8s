/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package k8s_mariadb_com_v1alpha1

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
	_ datasource.DataSource = &K8SMariadbComDatabaseV1Alpha1Manifest{}
)

func NewK8SMariadbComDatabaseV1Alpha1Manifest() datasource.DataSource {
	return &K8SMariadbComDatabaseV1Alpha1Manifest{}
}

type K8SMariadbComDatabaseV1Alpha1Manifest struct{}

type K8SMariadbComDatabaseV1Alpha1ManifestData struct {
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
		CharacterSet *string `tfsdk:"character_set" json:"characterSet,omitempty"`
		Collate      *string `tfsdk:"collate" json:"collate,omitempty"`
		MariaDbRef   *struct {
			ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
			Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
			Name            *string `tfsdk:"name" json:"name,omitempty"`
			Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
			ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
			Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
			WaitForIt       *bool   `tfsdk:"wait_for_it" json:"waitForIt,omitempty"`
		} `tfsdk:"maria_db_ref" json:"mariaDbRef,omitempty"`
		Name            *string `tfsdk:"name" json:"name,omitempty"`
		RequeueInterval *string `tfsdk:"requeue_interval" json:"requeueInterval,omitempty"`
		RetryInterval   *string `tfsdk:"retry_interval" json:"retryInterval,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *K8SMariadbComDatabaseV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_k8s_mariadb_com_database_v1alpha1_manifest"
}

func (r *K8SMariadbComDatabaseV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Database is the Schema for the databases API. It is used to define a logical database as if you were running a 'CREATE DATABASE' statement.",
		MarkdownDescription: "Database is the Schema for the databases API. It is used to define a logical database as if you were running a 'CREATE DATABASE' statement.",
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
				Description:         "DatabaseSpec defines the desired state of Database",
				MarkdownDescription: "DatabaseSpec defines the desired state of Database",
				Attributes: map[string]schema.Attribute{
					"character_set": schema.StringAttribute{
						Description:         "CharacterSet to use in the Database.",
						MarkdownDescription: "CharacterSet to use in the Database.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"collate": schema.StringAttribute{
						Description:         "CharacterSet to use in the Database.",
						MarkdownDescription: "CharacterSet to use in the Database.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"maria_db_ref": schema.SingleNestedAttribute{
						Description:         "MariaDBRef is a reference to a MariaDB object.",
						MarkdownDescription: "MariaDBRef is a reference to a MariaDB object.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "API version of the referent.",
								MarkdownDescription: "API version of the referent.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"field_path": schema.StringAttribute{
								Description:         "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
								MarkdownDescription: "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								MarkdownDescription: "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resource_version": schema.StringAttribute{
								Description:         "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								MarkdownDescription: "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"uid": schema.StringAttribute{
								Description:         "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								MarkdownDescription: "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"wait_for_it": schema.BoolAttribute{
								Description:         "WaitForIt indicates whether the controller using this reference should wait for MariaDB to be ready.",
								MarkdownDescription: "WaitForIt indicates whether the controller using this reference should wait for MariaDB to be ready.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"name": schema.StringAttribute{
						Description:         "Name overrides the default Database name provided by metadata.name.",
						MarkdownDescription: "Name overrides the default Database name provided by metadata.name.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(80),
						},
					},

					"requeue_interval": schema.StringAttribute{
						Description:         "RequeueInterval is used to perform requeue reconciliations.",
						MarkdownDescription: "RequeueInterval is used to perform requeue reconciliations.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"retry_interval": schema.StringAttribute{
						Description:         "RetryInterval is the interval used to perform retries.",
						MarkdownDescription: "RetryInterval is the interval used to perform retries.",
						Required:            false,
						Optional:            true,
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

func (r *K8SMariadbComDatabaseV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_k8s_mariadb_com_database_v1alpha1_manifest")

	var model K8SMariadbComDatabaseV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("k8s.mariadb.com/v1alpha1")
	model.Kind = pointer.String("Database")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
