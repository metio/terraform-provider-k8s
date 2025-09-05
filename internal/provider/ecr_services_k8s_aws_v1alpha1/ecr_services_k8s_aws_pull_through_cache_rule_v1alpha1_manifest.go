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
		CredentialARN *string `tfsdk:"credential_arn" json:"credentialARN,omitempty"`
		CredentialRef *struct {
			From *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"credential_ref" json:"credentialRef,omitempty"`
		CustomRoleARN *string `tfsdk:"custom_role_arn" json:"customRoleARN,omitempty"`
		CustomRoleRef *struct {
			From *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"custom_role_ref" json:"customRoleRef,omitempty"`
		EcrRepositoryPrefix      *string `tfsdk:"ecr_repository_prefix" json:"ecrRepositoryPrefix,omitempty"`
		RegistryID               *string `tfsdk:"registry_id" json:"registryID,omitempty"`
		UpstreamRegistry         *string `tfsdk:"upstream_registry" json:"upstreamRegistry,omitempty"`
		UpstreamRegistryURL      *string `tfsdk:"upstream_registry_url" json:"upstreamRegistryURL,omitempty"`
		UpstreamRepositoryPrefix *string `tfsdk:"upstream_repository_prefix" json:"upstreamRepositoryPrefix,omitempty"`
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
				Description:         "PullThroughCacheRuleSpec defines the desired state of PullThroughCacheRule. The details of a pull through cache rule.",
				MarkdownDescription: "PullThroughCacheRuleSpec defines the desired state of PullThroughCacheRule. The details of a pull through cache rule.",
				Attributes: map[string]schema.Attribute{
					"credential_arn": schema.StringAttribute{
						Description:         "The Amazon Resource Name (ARN) of the Amazon Web Services Secrets Manager secret that identifies the credentials to authenticate to the upstream registry. Regex Pattern: '^arn:aws:secretsmanager:[a-zA-Z0-9-:]+:secret:ecr-pullthroughcache/[a-zA-Z0-9/_+=.@-]+$'",
						MarkdownDescription: "The Amazon Resource Name (ARN) of the Amazon Web Services Secrets Manager secret that identifies the credentials to authenticate to the upstream registry. Regex Pattern: '^arn:aws:secretsmanager:[a-zA-Z0-9-:]+:secret:ecr-pullthroughcache/[a-zA-Z0-9/_+=.@-]+$'",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"credential_ref": schema.SingleNestedAttribute{
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",
						Attributes: map[string]schema.Attribute{
							"from": schema.SingleNestedAttribute{
								Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"custom_role_arn": schema.StringAttribute{
						Description:         "Amazon Resource Name (ARN) of the IAM role to be assumed by Amazon ECR to authenticate to the ECR upstream registry. This role must be in the same account as the registry that you are configuring.",
						MarkdownDescription: "Amazon Resource Name (ARN) of the IAM role to be assumed by Amazon ECR to authenticate to the ECR upstream registry. This role must be in the same account as the registry that you are configuring.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"custom_role_ref": schema.SingleNestedAttribute{
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",
						Attributes: map[string]schema.Attribute{
							"from": schema.SingleNestedAttribute{
								Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"ecr_repository_prefix": schema.StringAttribute{
						Description:         "The repository name prefix to use when caching images from the source registry. There is always an assumed / applied to the end of the prefix. If you specify ecr-public as the prefix, Amazon ECR treats that as ecr-public/. Regex Pattern: '^((?:[a-z0-9]+(?:[._-][a-z0-9]+)*/)*[a-z0-9]+(?:[._-][a-z0-9]+)*/?|ROOT)$'",
						MarkdownDescription: "The repository name prefix to use when caching images from the source registry. There is always an assumed / applied to the end of the prefix. If you specify ecr-public as the prefix, Amazon ECR treats that as ecr-public/. Regex Pattern: '^((?:[a-z0-9]+(?:[._-][a-z0-9]+)*/)*[a-z0-9]+(?:[._-][a-z0-9]+)*/?|ROOT)$'",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"registry_id": schema.StringAttribute{
						Description:         "The Amazon Web Services account ID associated with the registry to create the pull through cache rule for. If you do not specify a registry, the default registry is assumed. Regex Pattern: '^[0-9]{12}$'",
						MarkdownDescription: "The Amazon Web Services account ID associated with the registry to create the pull through cache rule for. If you do not specify a registry, the default registry is assumed. Regex Pattern: '^[0-9]{12}$'",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"upstream_registry": schema.StringAttribute{
						Description:         "The name of the upstream registry.",
						MarkdownDescription: "The name of the upstream registry.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"upstream_registry_url": schema.StringAttribute{
						Description:         "The registry URL of the upstream public registry to use as the source for the pull through cache rule. The following is the syntax to use for each supported upstream registry. * Amazon ECR (ecr) – .dkr.ecr..amazonaws.com * Amazon ECR Public (ecr-public) – public.ecr.aws * Docker Hub (docker-hub) – registry-1.docker.io * GitHub Container Registry (github-container-registry) – ghcr.io * GitLab Container Registry (gitlab-container-registry) – registry.gitlab.com * Kubernetes (k8s) – registry.k8s.io * Microsoft Azure Container Registry (azure-container-registry) – .azurecr.io * Quay (quay) – quay.io",
						MarkdownDescription: "The registry URL of the upstream public registry to use as the source for the pull through cache rule. The following is the syntax to use for each supported upstream registry. * Amazon ECR (ecr) – .dkr.ecr..amazonaws.com * Amazon ECR Public (ecr-public) – public.ecr.aws * Docker Hub (docker-hub) – registry-1.docker.io * GitHub Container Registry (github-container-registry) – ghcr.io * GitLab Container Registry (gitlab-container-registry) – registry.gitlab.com * Kubernetes (k8s) – registry.k8s.io * Microsoft Azure Container Registry (azure-container-registry) – .azurecr.io * Quay (quay) – quay.io",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"upstream_repository_prefix": schema.StringAttribute{
						Description:         "The repository name prefix of the upstream registry to match with the upstream repository name. When this field isn't specified, Amazon ECR will use the ROOT. Regex Pattern: '^((?:[a-z0-9]+(?:[._-][a-z0-9]+)*/)*[a-z0-9]+(?:[._-][a-z0-9]+)*/?|ROOT)$'",
						MarkdownDescription: "The repository name prefix of the upstream registry to match with the upstream repository name. When this field isn't specified, Amazon ECR will use the ROOT. Regex Pattern: '^((?:[a-z0-9]+(?:[._-][a-z0-9]+)*/)*[a-z0-9]+(?:[._-][a-z0-9]+)*/?|ROOT)$'",
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
