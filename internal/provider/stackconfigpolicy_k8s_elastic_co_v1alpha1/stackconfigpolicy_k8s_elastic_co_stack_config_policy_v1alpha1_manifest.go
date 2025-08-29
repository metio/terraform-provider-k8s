/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package stackconfigpolicy_k8s_elastic_co_v1alpha1

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
	_ datasource.DataSource = &StackconfigpolicyK8SElasticCoStackConfigPolicyV1Alpha1Manifest{}
)

func NewStackconfigpolicyK8SElasticCoStackConfigPolicyV1Alpha1Manifest() datasource.DataSource {
	return &StackconfigpolicyK8SElasticCoStackConfigPolicyV1Alpha1Manifest{}
}

type StackconfigpolicyK8SElasticCoStackConfigPolicyV1Alpha1Manifest struct{}

type StackconfigpolicyK8SElasticCoStackConfigPolicyV1Alpha1ManifestData struct {
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
		Elasticsearch *struct {
			ClusterSettings        *map[string]string `tfsdk:"cluster_settings" json:"clusterSettings,omitempty"`
			Config                 *map[string]string `tfsdk:"config" json:"config,omitempty"`
			IndexLifecyclePolicies *map[string]string `tfsdk:"index_lifecycle_policies" json:"indexLifecyclePolicies,omitempty"`
			IndexTemplates         *struct {
				ComponentTemplates       *map[string]string `tfsdk:"component_templates" json:"componentTemplates,omitempty"`
				ComposableIndexTemplates *map[string]string `tfsdk:"composable_index_templates" json:"composableIndexTemplates,omitempty"`
			} `tfsdk:"index_templates" json:"indexTemplates,omitempty"`
			IngestPipelines           *map[string]string `tfsdk:"ingest_pipelines" json:"ingestPipelines,omitempty"`
			SecretMounts              *map[string]string `tfsdk:"secret_mounts" json:"secretMounts,omitempty"`
			SecureSettings            *map[string]string `tfsdk:"secure_settings" json:"secureSettings,omitempty"`
			SecurityRoleMappings      *map[string]string `tfsdk:"security_role_mappings" json:"securityRoleMappings,omitempty"`
			SnapshotLifecyclePolicies *map[string]string `tfsdk:"snapshot_lifecycle_policies" json:"snapshotLifecyclePolicies,omitempty"`
			SnapshotRepositories      *map[string]string `tfsdk:"snapshot_repositories" json:"snapshotRepositories,omitempty"`
		} `tfsdk:"elasticsearch" json:"elasticsearch,omitempty"`
		Kibana *struct {
			Config         *map[string]string `tfsdk:"config" json:"config,omitempty"`
			SecureSettings *map[string]string `tfsdk:"secure_settings" json:"secureSettings,omitempty"`
		} `tfsdk:"kibana" json:"kibana,omitempty"`
		ResourceSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"resource_selector" json:"resourceSelector,omitempty"`
		SecureSettings *[]struct {
			Entries *[]struct {
				Key  *string `tfsdk:"key" json:"key,omitempty"`
				Path *string `tfsdk:"path" json:"path,omitempty"`
			} `tfsdk:"entries" json:"entries,omitempty"`
			SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
		} `tfsdk:"secure_settings" json:"secureSettings,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *StackconfigpolicyK8SElasticCoStackConfigPolicyV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_stackconfigpolicy_k8s_elastic_co_stack_config_policy_v1alpha1_manifest"
}

func (r *StackconfigpolicyK8SElasticCoStackConfigPolicyV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "StackConfigPolicy represents a StackConfigPolicy resource in a Kubernetes cluster.",
		MarkdownDescription: "StackConfigPolicy represents a StackConfigPolicy resource in a Kubernetes cluster.",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"elasticsearch": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"cluster_settings": schema.MapAttribute{
								Description:         "ClusterSettings holds the Elasticsearch cluster settings (/_cluster/settings)",
								MarkdownDescription: "ClusterSettings holds the Elasticsearch cluster settings (/_cluster/settings)",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"config": schema.MapAttribute{
								Description:         "Config holds the settings that go into elasticsearch.yml.",
								MarkdownDescription: "Config holds the settings that go into elasticsearch.yml.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"index_lifecycle_policies": schema.MapAttribute{
								Description:         "IndexLifecyclePolicies holds the Index Lifecycle policies settings (/_ilm/policy)",
								MarkdownDescription: "IndexLifecyclePolicies holds the Index Lifecycle policies settings (/_ilm/policy)",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"index_templates": schema.SingleNestedAttribute{
								Description:         "IndexTemplates holds the Index and Component Templates settings",
								MarkdownDescription: "IndexTemplates holds the Index and Component Templates settings",
								Attributes: map[string]schema.Attribute{
									"component_templates": schema.MapAttribute{
										Description:         "ComponentTemplates holds the Component Templates settings (/_component_template)",
										MarkdownDescription: "ComponentTemplates holds the Component Templates settings (/_component_template)",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"composable_index_templates": schema.MapAttribute{
										Description:         "ComposableIndexTemplates holds the Index Templates settings (/_index_template)",
										MarkdownDescription: "ComposableIndexTemplates holds the Index Templates settings (/_index_template)",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"ingest_pipelines": schema.MapAttribute{
								Description:         "IngestPipelines holds the Ingest Pipelines settings (/_ingest/pipeline)",
								MarkdownDescription: "IngestPipelines holds the Ingest Pipelines settings (/_ingest/pipeline)",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secret_mounts": schema.MapAttribute{
								Description:         "SecretMounts are additional Secrets that need to be mounted into the Elasticsearch pods.",
								MarkdownDescription: "SecretMounts are additional Secrets that need to be mounted into the Elasticsearch pods.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secure_settings": schema.MapAttribute{
								Description:         "SecureSettings are additional Secrets that contain data to be configured to Elasticsearch's keystore.",
								MarkdownDescription: "SecureSettings are additional Secrets that contain data to be configured to Elasticsearch's keystore.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"security_role_mappings": schema.MapAttribute{
								Description:         "SecurityRoleMappings holds the Role Mappings settings (/_security/role_mapping)",
								MarkdownDescription: "SecurityRoleMappings holds the Role Mappings settings (/_security/role_mapping)",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"snapshot_lifecycle_policies": schema.MapAttribute{
								Description:         "SnapshotLifecyclePolicies holds the Snapshot Lifecycle Policies settings (/_slm/policy)",
								MarkdownDescription: "SnapshotLifecyclePolicies holds the Snapshot Lifecycle Policies settings (/_slm/policy)",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"snapshot_repositories": schema.MapAttribute{
								Description:         "SnapshotRepositories holds the Snapshot Repositories settings (/_snapshot)",
								MarkdownDescription: "SnapshotRepositories holds the Snapshot Repositories settings (/_snapshot)",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"kibana": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"config": schema.MapAttribute{
								Description:         "Config holds the settings that go into kibana.yml.",
								MarkdownDescription: "Config holds the settings that go into kibana.yml.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secure_settings": schema.MapAttribute{
								Description:         "SecureSettings are additional Secrets that contain data to be configured to Kibana's keystore.",
								MarkdownDescription: "SecureSettings are additional Secrets that contain data to be configured to Kibana's keystore.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"resource_selector": schema.SingleNestedAttribute{
						Description:         "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
						MarkdownDescription: "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
						Attributes: map[string]schema.Attribute{
							"match_expressions": schema.ListNestedAttribute{
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "key is the label key that the selector applies to.",
											MarkdownDescription: "key is the label key that the selector applies to.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											ElementType:         types.StringType,
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

							"match_labels": schema.MapAttribute{
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"secure_settings": schema.ListNestedAttribute{
						Description:         "Deprecated: SecureSettings only applies to Elasticsearch and is deprecated. It must be set per application instead.",
						MarkdownDescription: "Deprecated: SecureSettings only applies to Elasticsearch and is deprecated. It must be set per application instead.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"entries": schema.ListNestedAttribute{
									Description:         "Entries define how to project each key-value pair in the secret to filesystem paths. If not defined, all keys will be projected to similarly named paths in the filesystem. If defined, only the specified keys will be projected to the corresponding paths.",
									MarkdownDescription: "Entries define how to project each key-value pair in the secret to filesystem paths. If not defined, all keys will be projected to similarly named paths in the filesystem. If defined, only the specified keys will be projected to the corresponding paths.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "Key is the key contained in the secret.",
												MarkdownDescription: "Key is the key contained in the secret.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"path": schema.StringAttribute{
												Description:         "Path is the relative file path to map the key to. Path must not be an absolute file path and must not contain any '..' components.",
												MarkdownDescription: "Path is the relative file path to map the key to. Path must not be an absolute file path and must not contain any '..' components.",
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

								"secret_name": schema.StringAttribute{
									Description:         "SecretName is the name of the secret.",
									MarkdownDescription: "SecretName is the name of the secret.",
									Required:            true,
									Optional:            false,
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

func (r *StackconfigpolicyK8SElasticCoStackConfigPolicyV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_stackconfigpolicy_k8s_elastic_co_stack_config_policy_v1alpha1_manifest")

	var model StackconfigpolicyK8SElasticCoStackConfigPolicyV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("stackconfigpolicy.k8s.elastic.co/v1alpha1")
	model.Kind = pointer.String("StackConfigPolicy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
