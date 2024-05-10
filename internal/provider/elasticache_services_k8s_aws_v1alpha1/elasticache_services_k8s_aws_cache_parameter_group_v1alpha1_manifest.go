/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package elasticache_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource = &ElasticacheServicesK8SAwsCacheParameterGroupV1Alpha1Manifest{}
)

func NewElasticacheServicesK8SAwsCacheParameterGroupV1Alpha1Manifest() datasource.DataSource {
	return &ElasticacheServicesK8SAwsCacheParameterGroupV1Alpha1Manifest{}
}

type ElasticacheServicesK8SAwsCacheParameterGroupV1Alpha1Manifest struct{}

type ElasticacheServicesK8SAwsCacheParameterGroupV1Alpha1ManifestData struct {
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
		CacheParameterGroupFamily *string `tfsdk:"cache_parameter_group_family" json:"cacheParameterGroupFamily,omitempty"`
		CacheParameterGroupName   *string `tfsdk:"cache_parameter_group_name" json:"cacheParameterGroupName,omitempty"`
		Description               *string `tfsdk:"description" json:"description,omitempty"`
		ParameterNameValues       *[]struct {
			ParameterName  *string `tfsdk:"parameter_name" json:"parameterName,omitempty"`
			ParameterValue *string `tfsdk:"parameter_value" json:"parameterValue,omitempty"`
		} `tfsdk:"parameter_name_values" json:"parameterNameValues,omitempty"`
		Tags *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ElasticacheServicesK8SAwsCacheParameterGroupV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_elasticache_services_k8s_aws_cache_parameter_group_v1alpha1_manifest"
}

func (r *ElasticacheServicesK8SAwsCacheParameterGroupV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CacheParameterGroup is the Schema for the CacheParameterGroups API",
		MarkdownDescription: "CacheParameterGroup is the Schema for the CacheParameterGroups API",
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
				Description:         "CacheParameterGroupSpec defines the desired state of CacheParameterGroup.Represents the output of a CreateCacheParameterGroup operation.",
				MarkdownDescription: "CacheParameterGroupSpec defines the desired state of CacheParameterGroup.Represents the output of a CreateCacheParameterGroup operation.",
				Attributes: map[string]schema.Attribute{
					"cache_parameter_group_family": schema.StringAttribute{
						Description:         "The name of the cache parameter group family that the cache parameter groupcan be used with.Valid values are: memcached1.4 | memcached1.5 | memcached1.6 | redis2.6 |redis2.8 | redis3.2 | redis4.0 | redis5.0 | redis6.x",
						MarkdownDescription: "The name of the cache parameter group family that the cache parameter groupcan be used with.Valid values are: memcached1.4 | memcached1.5 | memcached1.6 | redis2.6 |redis2.8 | redis3.2 | redis4.0 | redis5.0 | redis6.x",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"cache_parameter_group_name": schema.StringAttribute{
						Description:         "A user-specified name for the cache parameter group.",
						MarkdownDescription: "A user-specified name for the cache parameter group.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"description": schema.StringAttribute{
						Description:         "A user-specified description for the cache parameter group.",
						MarkdownDescription: "A user-specified description for the cache parameter group.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"parameter_name_values": schema.ListNestedAttribute{
						Description:         "An array of parameter names and values for the parameter update. You mustsupply at least one parameter name and value; subsequent arguments are optional.A maximum of 20 parameters may be modified per request.",
						MarkdownDescription: "An array of parameter names and values for the parameter update. You mustsupply at least one parameter name and value; subsequent arguments are optional.A maximum of 20 parameters may be modified per request.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"parameter_name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"parameter_value": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "A list of tags to be added to this resource. A tag is a key-value pair. Atag key must be accompanied by a tag value, although null is accepted.",
						MarkdownDescription: "A list of tags to be added to this resource. A tag is a key-value pair. Atag key must be accompanied by a tag value, although null is accepted.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *ElasticacheServicesK8SAwsCacheParameterGroupV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_elasticache_services_k8s_aws_cache_parameter_group_v1alpha1_manifest")

	var model ElasticacheServicesK8SAwsCacheParameterGroupV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("elasticache.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("CacheParameterGroup")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
