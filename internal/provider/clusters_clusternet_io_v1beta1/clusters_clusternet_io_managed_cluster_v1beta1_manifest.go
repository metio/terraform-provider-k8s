/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package clusters_clusternet_io_v1beta1

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
	_ datasource.DataSource = &ClustersClusternetIoManagedClusterV1Beta1Manifest{}
)

func NewClustersClusternetIoManagedClusterV1Beta1Manifest() datasource.DataSource {
	return &ClustersClusternetIoManagedClusterV1Beta1Manifest{}
}

type ClustersClusternetIoManagedClusterV1Beta1Manifest struct{}

type ClustersClusternetIoManagedClusterV1Beta1ManifestData struct {
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
		ClusterId           *string `tfsdk:"cluster_id" json:"clusterId,omitempty"`
		ClusterInitBaseName *string `tfsdk:"cluster_init_base_name" json:"clusterInitBaseName,omitempty"`
		ClusterType         *string `tfsdk:"cluster_type" json:"clusterType,omitempty"`
		SyncMode            *string `tfsdk:"sync_mode" json:"syncMode,omitempty"`
		Taints              *[]struct {
			Effect    *string `tfsdk:"effect" json:"effect,omitempty"`
			Key       *string `tfsdk:"key" json:"key,omitempty"`
			TimeAdded *string `tfsdk:"time_added" json:"timeAdded,omitempty"`
			Value     *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"taints" json:"taints,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ClustersClusternetIoManagedClusterV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_clusters_clusternet_io_managed_cluster_v1beta1_manifest"
}

func (r *ClustersClusternetIoManagedClusterV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ManagedCluster is the Schema for the managedclusters API",
		MarkdownDescription: "ManagedCluster is the Schema for the managedclusters API",
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
				Description:         "ManagedClusterSpec defines the desired state of ManagedCluster",
				MarkdownDescription: "ManagedClusterSpec defines the desired state of ManagedCluster",
				Attributes: map[string]schema.Attribute{
					"cluster_id": schema.StringAttribute{
						Description:         "ClusterID, a Random (Version 4) UUID, is a unique value in time and space value representing for child cluster.It is typically generated by the clusternet agent on the successful creation of a 'clusternet-agent' Leasein the child cluster.Also it is not allowed to change on PUT operations.",
						MarkdownDescription: "ClusterID, a Random (Version 4) UUID, is a unique value in time and space value representing for child cluster.It is typically generated by the clusternet agent on the successful creation of a 'clusternet-agent' Leasein the child cluster.Also it is not allowed to change on PUT operations.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
						},
					},

					"cluster_init_base_name": schema.StringAttribute{
						Description:         "ClusterInitBaseName denotes the name of a Base used for initialization.Also a taint 'clusters.clusternet.io/initialization:NoSchedule' will be added during the operation and removedafter successful initialization.If this cluster has got an annotation 'clusters.clusternet.io/skip-cluster-init', this field will be empty.Normally this field is fully managed by clusternet-controller-manager and immutable.",
						MarkdownDescription: "ClusterInitBaseName denotes the name of a Base used for initialization.Also a taint 'clusters.clusternet.io/initialization:NoSchedule' will be added during the operation and removedafter successful initialization.If this cluster has got an annotation 'clusters.clusternet.io/skip-cluster-init', this field will be empty.Normally this field is fully managed by clusternet-controller-manager and immutable.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(512),
						},
					},

					"cluster_type": schema.StringAttribute{
						Description:         "ClusterType denotes the type of the child cluster.",
						MarkdownDescription: "ClusterType denotes the type of the child cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"sync_mode": schema.StringAttribute{
						Description:         "SyncMode decides how to sync resources from parent cluster to child cluster.",
						MarkdownDescription: "SyncMode decides how to sync resources from parent cluster to child cluster.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Push", "Pull", "Dual"),
						},
					},

					"taints": schema.ListNestedAttribute{
						Description:         "Taints has the 'effect' on any resource that does not tolerate the Taint.",
						MarkdownDescription: "Taints has the 'effect' on any resource that does not tolerate the Taint.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"effect": schema.StringAttribute{
									Description:         "Required. The effect of the taint on podsthat do not tolerate the taint.Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
									MarkdownDescription: "Required. The effect of the taint on podsthat do not tolerate the taint.Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"key": schema.StringAttribute{
									Description:         "Required. The taint key to be applied to a node.",
									MarkdownDescription: "Required. The taint key to be applied to a node.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"time_added": schema.StringAttribute{
									Description:         "TimeAdded represents the time at which the taint was added.It is only written for NoExecute taints.",
									MarkdownDescription: "TimeAdded represents the time at which the taint was added.It is only written for NoExecute taints.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										validators.DateTime64Validator(),
									},
								},

								"value": schema.StringAttribute{
									Description:         "The taint value corresponding to the taint key.",
									MarkdownDescription: "The taint value corresponding to the taint key.",
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

func (r *ClustersClusternetIoManagedClusterV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_clusters_clusternet_io_managed_cluster_v1beta1_manifest")

	var model ClustersClusternetIoManagedClusterV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("clusters.clusternet.io/v1beta1")
	model.Kind = pointer.String("ManagedCluster")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
