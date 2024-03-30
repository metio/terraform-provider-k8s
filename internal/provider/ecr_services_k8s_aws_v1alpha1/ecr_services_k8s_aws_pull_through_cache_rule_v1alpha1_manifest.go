/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package ecr_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource = &EcrServicesK8SAwsPullThroughCacheRuleV1Alpha1Manifest{}
)

func NewEcrServicesK8SAwsPullThroughCacheRuleV1Alpha1Manifest() datasource.DataSource {
	return &EcrServicesK8SAwsPullThroughCacheRuleV1Alpha1Manifest{}
}

type EcrServicesK8SAwsPullThroughCacheRuleV1Alpha1Manifest struct{}

type EcrServicesK8SAwsPullThroughCacheRuleV1Alpha1ManifestData struct {
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
		EcrRepositoryPrefix *string `tfsdk:"ecr_repository_prefix" json:"ecrRepositoryPrefix,omitempty"`
		RegistryID          *string `tfsdk:"registry_id" json:"registryID,omitempty"`
		UpstreamRegistryURL *string `tfsdk:"upstream_registry_url" json:"upstreamRegistryURL,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *EcrServicesK8SAwsPullThroughCacheRuleV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_ecr_services_k8s_aws_pull_through_cache_rule_v1alpha1_manifest"
}

func (r *EcrServicesK8SAwsPullThroughCacheRuleV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "PullThroughCacheRule is the Schema for the PullThroughCacheRules API",
		MarkdownDescription: "PullThroughCacheRule is the Schema for the PullThroughCacheRules API",
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
				Description:         "PullThroughCacheRuleSpec defines the desired state of PullThroughCacheRule.The details of a pull through cache rule.",
				MarkdownDescription: "PullThroughCacheRuleSpec defines the desired state of PullThroughCacheRule.The details of a pull through cache rule.",
				Attributes: map[string]schema.Attribute{
					"ecr_repository_prefix": schema.StringAttribute{
						Description:         "The repository name prefix to use when caching images from the source registry.",
						MarkdownDescription: "The repository name prefix to use when caching images from the source registry.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"registry_id": schema.StringAttribute{
						Description:         "The Amazon Web Services account ID associated with the registry to createthe pull through cache rule for. If you do not specify a registry, the defaultregistry is assumed.",
						MarkdownDescription: "The Amazon Web Services account ID associated with the registry to createthe pull through cache rule for. If you do not specify a registry, the defaultregistry is assumed.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"upstream_registry_url": schema.StringAttribute{
						Description:         "The registry URL of the upstream public registry to use as the source forthe pull through cache rule.",
						MarkdownDescription: "The registry URL of the upstream public registry to use as the source forthe pull through cache rule.",
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

func (r *EcrServicesK8SAwsPullThroughCacheRuleV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_ecr_services_k8s_aws_pull_through_cache_rule_v1alpha1_manifest")

	var model EcrServicesK8SAwsPullThroughCacheRuleV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("ecr.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("PullThroughCacheRule")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
