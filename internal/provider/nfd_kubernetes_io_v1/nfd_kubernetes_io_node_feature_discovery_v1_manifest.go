/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package nfd_kubernetes_io_v1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &NfdKubernetesIoNodeFeatureDiscoveryV1Manifest{}
)

func NewNfdKubernetesIoNodeFeatureDiscoveryV1Manifest() datasource.DataSource {
	return &NfdKubernetesIoNodeFeatureDiscoveryV1Manifest{}
}

type NfdKubernetesIoNodeFeatureDiscoveryV1Manifest struct{}

type NfdKubernetesIoNodeFeatureDiscoveryV1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
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
		ExtraLabelNs   *[]string `tfsdk:"extra_label_ns" json:"extraLabelNs,omitempty"`
		Instance       *string   `tfsdk:"instance" json:"instance,omitempty"`
		LabelWhiteList *string   `tfsdk:"label_white_list" json:"labelWhiteList,omitempty"`
		Operand        *struct {
			Image           *string `tfsdk:"image" json:"image,omitempty"`
			ImagePullPolicy *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
			ServicePort     *int64  `tfsdk:"service_port" json:"servicePort,omitempty"`
		} `tfsdk:"operand" json:"operand,omitempty"`
		PrunerOnDelete  *bool     `tfsdk:"pruner_on_delete" json:"prunerOnDelete,omitempty"`
		ResourceLabels  *[]string `tfsdk:"resource_labels" json:"resourceLabels,omitempty"`
		TopologyUpdater *bool     `tfsdk:"topology_updater" json:"topologyUpdater,omitempty"`
		WorkerConfig    *struct {
			ConfigData *string `tfsdk:"config_data" json:"configData,omitempty"`
		} `tfsdk:"worker_config" json:"workerConfig,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *NfdKubernetesIoNodeFeatureDiscoveryV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_nfd_kubernetes_io_node_feature_discovery_v1_manifest"
}

func (r *NfdKubernetesIoNodeFeatureDiscoveryV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "NodeFeatureDiscovery is the Schema for the nodefeaturediscoveries API",
		MarkdownDescription: "NodeFeatureDiscovery is the Schema for the nodefeaturediscoveries API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

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
				Description:         "NodeFeatureDiscoverySpec defines the desired state of NodeFeatureDiscovery",
				MarkdownDescription: "NodeFeatureDiscoverySpec defines the desired state of NodeFeatureDiscovery",
				Attributes: map[string]schema.Attribute{
					"extra_label_ns": schema.ListAttribute{
						Description:         "ExtraLabelNs defines the list of of allowed extra label namespaces By default, only allow labels in the default 'feature.node.kubernetes.io' label namespace",
						MarkdownDescription: "ExtraLabelNs defines the list of of allowed extra label namespaces By default, only allow labels in the default 'feature.node.kubernetes.io' label namespace",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"instance": schema.StringAttribute{
						Description:         "Instance name. Used to separate annotation namespaces for multiple parallel deployments.",
						MarkdownDescription: "Instance name. Used to separate annotation namespaces for multiple parallel deployments.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"label_white_list": schema.StringAttribute{
						Description:         "LabelWhiteList defines a regular expression for filtering feature labels based on their name. Each label must match against the given reqular expression in order to be published.",
						MarkdownDescription: "LabelWhiteList defines a regular expression for filtering feature labels based on their name. Each label must match against the given reqular expression in order to be published.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"operand": schema.SingleNestedAttribute{
						Description:         "OperandSpec describes configuration options for the operand",
						MarkdownDescription: "OperandSpec describes configuration options for the operand",
						Attributes: map[string]schema.Attribute{
							"image": schema.StringAttribute{
								Description:         "Image defines the image to pull for the NFD operand [defaults to registry.k8s.io/nfd/node-feature-discovery]",
								MarkdownDescription: "Image defines the image to pull for the NFD operand [defaults to registry.k8s.io/nfd/node-feature-discovery]",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`[a-zA-Z0-9\-]+`), ""),
								},
							},

							"image_pull_policy": schema.StringAttribute{
								Description:         "ImagePullPolicy defines Image pull policy for the NFD operand image [defaults to Always]",
								MarkdownDescription: "ImagePullPolicy defines Image pull policy for the NFD operand image [defaults to Always]",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_port": schema.Int64Attribute{
								Description:         "ServicePort specifies the TCP port that nfd-master listens for incoming requests.",
								MarkdownDescription: "ServicePort specifies the TCP port that nfd-master listens for incoming requests.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"pruner_on_delete": schema.BoolAttribute{
						Description:         "PruneOnDelete defines whether the NFD-master prune should be enabled or not. If enabled, the Operator will deploy an NFD-Master prune job that will remove all NFD labels (and other NFD-managed assets such as annotations, extended resources and taints) from the cluster nodes.",
						MarkdownDescription: "PruneOnDelete defines whether the NFD-master prune should be enabled or not. If enabled, the Operator will deploy an NFD-Master prune job that will remove all NFD labels (and other NFD-managed assets such as annotations, extended resources and taints) from the cluster nodes.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"resource_labels": schema.ListAttribute{
						Description:         "ResourceLabels defines the list of features to be advertised as extended resources instead of labels.",
						MarkdownDescription: "ResourceLabels defines the list of features to be advertised as extended resources instead of labels.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"topology_updater": schema.BoolAttribute{
						Description:         "Deploy the NFD-Topology-Updater NFD-Topology-Updater is a daemon responsible for examining allocated resources on a worker node to account for resources available to be allocated to new pod on a per-zone basis https://kubernetes-sigs.github.io/node-feature-discovery/master/get-started/introduction.html#nfd-topology-updater",
						MarkdownDescription: "Deploy the NFD-Topology-Updater NFD-Topology-Updater is a daemon responsible for examining allocated resources on a worker node to account for resources available to be allocated to new pod on a per-zone basis https://kubernetes-sigs.github.io/node-feature-discovery/master/get-started/introduction.html#nfd-topology-updater",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"worker_config": schema.SingleNestedAttribute{
						Description:         "WorkerConfig describes configuration options for the NFD worker.",
						MarkdownDescription: "WorkerConfig describes configuration options for the NFD worker.",
						Attributes: map[string]schema.Attribute{
							"config_data": schema.StringAttribute{
								Description:         "BinaryData holds the NFD configuration file",
								MarkdownDescription: "BinaryData holds the NFD configuration file",
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *NfdKubernetesIoNodeFeatureDiscoveryV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_nfd_kubernetes_io_node_feature_discovery_v1_manifest")

	var model NfdKubernetesIoNodeFeatureDiscoveryV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("nfd.kubernetes.io/v1")
	model.Kind = pointer.String("NodeFeatureDiscovery")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
