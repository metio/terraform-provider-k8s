/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package apps_kubeblocks_io_v1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &AppsKubeblocksIoClusterDefinitionV1Manifest{}
)

func NewAppsKubeblocksIoClusterDefinitionV1Manifest() datasource.DataSource {
	return &AppsKubeblocksIoClusterDefinitionV1Manifest{}
}

type AppsKubeblocksIoClusterDefinitionV1Manifest struct{}

type AppsKubeblocksIoClusterDefinitionV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Topologies *[]struct {
			Components *[]struct {
				CompDef *string `tfsdk:"comp_def" json:"compDef,omitempty"`
				Name    *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"components" json:"components,omitempty"`
			Default *bool   `tfsdk:"default" json:"default,omitempty"`
			Name    *string `tfsdk:"name" json:"name,omitempty"`
			Orders  *struct {
				Provision *[]string `tfsdk:"provision" json:"provision,omitempty"`
				Terminate *[]string `tfsdk:"terminate" json:"terminate,omitempty"`
				Update    *[]string `tfsdk:"update" json:"update,omitempty"`
			} `tfsdk:"orders" json:"orders,omitempty"`
		} `tfsdk:"topologies" json:"topologies,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppsKubeblocksIoClusterDefinitionV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apps_kubeblocks_io_cluster_definition_v1_manifest"
}

func (r *AppsKubeblocksIoClusterDefinitionV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClusterDefinition defines the topology for databases or storage systems, offering a variety of topological configurations to meet diverse deployment needs and scenarios. It includes a list of Components, each linked to a ComponentDefinition, which enhances reusability and reduce redundancy. For example, widely used components such as etcd and Zookeeper can be defined once and reused across multiple ClusterDefinitions, simplifying the setup of new systems. Additionally, ClusterDefinition also specifies the sequence of startup, upgrade, and shutdown for Components, ensuring a controlled and predictable management of component lifecycles.",
		MarkdownDescription: "ClusterDefinition defines the topology for databases or storage systems, offering a variety of topological configurations to meet diverse deployment needs and scenarios. It includes a list of Components, each linked to a ComponentDefinition, which enhances reusability and reduce redundancy. For example, widely used components such as etcd and Zookeeper can be defined once and reused across multiple ClusterDefinitions, simplifying the setup of new systems. Additionally, ClusterDefinition also specifies the sequence of startup, upgrade, and shutdown for Components, ensuring a controlled and predictable management of component lifecycles.",
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
				Description:         "ClusterDefinitionSpec defines the desired state of ClusterDefinition.",
				MarkdownDescription: "ClusterDefinitionSpec defines the desired state of ClusterDefinition.",
				Attributes: map[string]schema.Attribute{
					"topologies": schema.ListNestedAttribute{
						Description:         "Topologies defines all possible topologies within the cluster.",
						MarkdownDescription: "Topologies defines all possible topologies within the cluster.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"components": schema.ListNestedAttribute{
									Description:         "Components specifies the components in the topology.",
									MarkdownDescription: "Components specifies the components in the topology.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"comp_def": schema.StringAttribute{
												Description:         "Specifies the exact name, name prefix, or regular expression pattern for matching the name of the ComponentDefinition custom resource (CR) that defines the Component's characteristics and behavior. The system selects the ComponentDefinition CR with the latest version that matches the pattern. This approach allows: 1. Precise selection by providing the exact name of a ComponentDefinition CR. 2. Flexible and automatic selection of the most up-to-date ComponentDefinition CR by specifying a name prefix or regular expression pattern. Once set, this field cannot be updated.",
												MarkdownDescription: "Specifies the exact name, name prefix, or regular expression pattern for matching the name of the ComponentDefinition custom resource (CR) that defines the Component's characteristics and behavior. The system selects the ComponentDefinition CR with the latest version that matches the pattern. This approach allows: 1. Precise selection by providing the exact name of a ComponentDefinition CR. 2. Flexible and automatic selection of the most up-to-date ComponentDefinition CR by specifying a name prefix or regular expression pattern. Once set, this field cannot be updated.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(64),
												},
											},

											"name": schema.StringAttribute{
												Description:         "Defines the unique identifier of the component within the cluster topology. It follows IANA Service naming rules and is used as part of the Service's DNS name. The name must start with a lowercase letter, can contain lowercase letters, numbers, and hyphens, and must end with a lowercase letter or number. Cannot be updated once set.",
												MarkdownDescription: "Defines the unique identifier of the component within the cluster topology. It follows IANA Service naming rules and is used as part of the Service's DNS name. The name must start with a lowercase letter, can contain lowercase letters, numbers, and hyphens, and must end with a lowercase letter or number. Cannot be updated once set.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(16),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([a-z0-9\-]*[a-z0-9])?$`), ""),
												},
											},
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"default": schema.BoolAttribute{
									Description:         "Default indicates whether this topology serves as the default configuration. When set to true, this topology is automatically used unless another is explicitly specified.",
									MarkdownDescription: "Default indicates whether this topology serves as the default configuration. When set to true, this topology is automatically used unless another is explicitly specified.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name is the unique identifier for the cluster topology. Cannot be updated.",
									MarkdownDescription: "Name is the unique identifier for the cluster topology. Cannot be updated.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(32),
									},
								},

								"orders": schema.SingleNestedAttribute{
									Description:         "Specifies the sequence in which components within a cluster topology are started, stopped, and upgraded. This ordering is crucial for maintaining the correct dependencies and operational flow across components.",
									MarkdownDescription: "Specifies the sequence in which components within a cluster topology are started, stopped, and upgraded. This ordering is crucial for maintaining the correct dependencies and operational flow across components.",
									Attributes: map[string]schema.Attribute{
										"provision": schema.ListAttribute{
											Description:         "Specifies the order for creating and initializing components. This is designed for components that depend on one another. Components without dependencies can be grouped together. Components that can be provisioned independently or have no dependencies can be listed together in the same stage, separated by commas.",
											MarkdownDescription: "Specifies the order for creating and initializing components. This is designed for components that depend on one another. Components without dependencies can be grouped together. Components that can be provisioned independently or have no dependencies can be listed together in the same stage, separated by commas.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"terminate": schema.ListAttribute{
											Description:         "Outlines the order for stopping and deleting components. This sequence is designed for components that require a graceful shutdown or have interdependencies. Components that can be terminated independently or have no dependencies can be listed together in the same stage, separated by commas.",
											MarkdownDescription: "Outlines the order for stopping and deleting components. This sequence is designed for components that require a graceful shutdown or have interdependencies. Components that can be terminated independently or have no dependencies can be listed together in the same stage, separated by commas.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"update": schema.ListAttribute{
											Description:         "Update determines the order for updating components' specifications, such as image upgrades or resource scaling. This sequence is designed for components that have dependencies or require specific update procedures. Components that can be updated independently or have no dependencies can be listed together in the same stage, separated by commas.",
											MarkdownDescription: "Update determines the order for updating components' specifications, such as image upgrades or resource scaling. This sequence is designed for components that have dependencies or require specific update procedures. Components that can be updated independently or have no dependencies can be listed together in the same stage, separated by commas.",
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

func (r *AppsKubeblocksIoClusterDefinitionV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apps_kubeblocks_io_cluster_definition_v1_manifest")

	var model AppsKubeblocksIoClusterDefinitionV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("apps.kubeblocks.io/v1")
	model.Kind = pointer.String("ClusterDefinition")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
