/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package cluster_ipfs_io_v1alpha1

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
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &ClusterIpfsIoIpfsClusterV1Alpha1Manifest{}
)

func NewClusterIpfsIoIpfsClusterV1Alpha1Manifest() datasource.DataSource {
	return &ClusterIpfsIoIpfsClusterV1Alpha1Manifest{}
}

type ClusterIpfsIoIpfsClusterV1Alpha1Manifest struct{}

type ClusterIpfsIoIpfsClusterV1Alpha1ManifestData struct {
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
		ClusterStorage *string `tfsdk:"cluster_storage" json:"clusterStorage,omitempty"`
		Follows        *[]struct {
			Name     *string `tfsdk:"name" json:"name,omitempty"`
			Template *string `tfsdk:"template" json:"template,omitempty"`
		} `tfsdk:"follows" json:"follows,omitempty"`
		IpfsResources *struct {
			Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
			Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"ipfs_resources" json:"ipfsResources,omitempty"`
		IpfsStorage *string `tfsdk:"ipfs_storage" json:"ipfsStorage,omitempty"`
		Networking  *struct {
			CircuitRelays *int64 `tfsdk:"circuit_relays" json:"circuitRelays,omitempty"`
			Public        *bool  `tfsdk:"public" json:"public,omitempty"`
		} `tfsdk:"networking" json:"networking,omitempty"`
		Replicas   *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
		Reprovider *struct {
			Interval *string `tfsdk:"interval" json:"interval,omitempty"`
			Strategy *string `tfsdk:"strategy" json:"strategy,omitempty"`
		} `tfsdk:"reprovider" json:"reprovider,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ClusterIpfsIoIpfsClusterV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cluster_ipfs_io_ipfs_cluster_v1alpha1_manifest"
}

func (r *ClusterIpfsIoIpfsClusterV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "IpfsCluster is the Schema for the ipfs API.",
		MarkdownDescription: "IpfsCluster is the Schema for the ipfs API.",
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
				Description:         "IpfsClusterSpec defines the desired state of the IpfsCluster.",
				MarkdownDescription: "IpfsClusterSpec defines the desired state of the IpfsCluster.",
				Attributes: map[string]schema.Attribute{
					"cluster_storage": schema.StringAttribute{
						Description:         "clusterStorage defines the amount of storage to be used by IPFS Cluster.",
						MarkdownDescription: "clusterStorage defines the amount of storage to be used by IPFS Cluster.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"follows": schema.ListNestedAttribute{
						Description:         "follows defines the list of other IPFS Clusters this one should follow.",
						MarkdownDescription: "follows defines the list of other IPFS Clusters this one should follow.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"template": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"ipfs_resources": schema.SingleNestedAttribute{
						Description:         "ipfsResources specifies the resource requirements for each IPFS container. If this value is omitted, then the operator will automatically determine these settings based on the storage sizes used.",
						MarkdownDescription: "ipfsResources specifies the resource requirements for each IPFS container. If this value is omitted, then the operator will automatically determine these settings based on the storage sizes used.",
						Attributes: map[string]schema.Attribute{
							"limits": schema.MapAttribute{
								Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"requests": schema.MapAttribute{
								Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

					"ipfs_storage": schema.StringAttribute{
						Description:         "ipfsStorage defines the total storage to be allocated by this resource.",
						MarkdownDescription: "ipfsStorage defines the total storage to be allocated by this resource.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"networking": schema.SingleNestedAttribute{
						Description:         "networking defines network configuration settings.",
						MarkdownDescription: "networking defines network configuration settings.",
						Attributes: map[string]schema.Attribute{
							"circuit_relays": schema.Int64Attribute{
								Description:         "circuitRelays defines how many CircuitRelays should be created.",
								MarkdownDescription: "circuitRelays defines how many CircuitRelays should be created.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"public": schema.BoolAttribute{
								Description:         "public is a switch which defines whether this IPFSCluster will use the global IPFS network or create its own.",
								MarkdownDescription: "public is a switch which defines whether this IPFSCluster will use the global IPFS network or create its own.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"replicas": schema.Int64Attribute{
						Description:         "replicas sets the number of replicas of IPFS Cluster nodes we should be running.",
						MarkdownDescription: "replicas sets the number of replicas of IPFS Cluster nodes we should be running.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"reprovider": schema.SingleNestedAttribute{
						Description:         "reprovider Describes the settings that each IPFS node should use when reproviding content.",
						MarkdownDescription: "reprovider Describes the settings that each IPFS node should use when reproviding content.",
						Attributes: map[string]schema.Attribute{
							"interval": schema.StringAttribute{
								Description:         "Interval sets the time between rounds of reproviding local content to the routing system. Defaults to '12h'.",
								MarkdownDescription: "Interval sets the time between rounds of reproviding local content to the routing system. Defaults to '12h'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"strategy": schema.StringAttribute{
								Description:         "Strategy specifies the reprovider strategy, defaults to 'all'.",
								MarkdownDescription: "Strategy specifies the reprovider strategy, defaults to 'all'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("all", "pinned", "roots"),
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

func (r *ClusterIpfsIoIpfsClusterV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cluster_ipfs_io_ipfs_cluster_v1alpha1_manifest")

	var model ClusterIpfsIoIpfsClusterV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("cluster.ipfs.io/v1alpha1")
	model.Kind = pointer.String("IpfsCluster")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
