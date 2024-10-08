/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package multicluster_crd_antrea_io_v1alpha2

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
	_ datasource.DataSource = &MulticlusterCrdAntreaIoClusterSetV1Alpha2Manifest{}
)

func NewMulticlusterCrdAntreaIoClusterSetV1Alpha2Manifest() datasource.DataSource {
	return &MulticlusterCrdAntreaIoClusterSetV1Alpha2Manifest{}
}

type MulticlusterCrdAntreaIoClusterSetV1Alpha2Manifest struct{}

type MulticlusterCrdAntreaIoClusterSetV1Alpha2ManifestData struct {
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
		ClusterID *string `tfsdk:"cluster_id" json:"clusterID,omitempty"`
		Leaders   *[]struct {
			ClusterID *string `tfsdk:"cluster_id" json:"clusterID,omitempty"`
			Secret    *string `tfsdk:"secret" json:"secret,omitempty"`
			Server    *string `tfsdk:"server" json:"server,omitempty"`
		} `tfsdk:"leaders" json:"leaders,omitempty"`
		Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *MulticlusterCrdAntreaIoClusterSetV1Alpha2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_multicluster_crd_antrea_io_cluster_set_v1alpha2_manifest"
}

func (r *MulticlusterCrdAntreaIoClusterSetV1Alpha2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClusterSet represents a ClusterSet.",
		MarkdownDescription: "ClusterSet represents a ClusterSet.",
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
				Description:         "ClusterSetSpec defines the desired state of ClusterSet.",
				MarkdownDescription: "ClusterSetSpec defines the desired state of ClusterSet.",
				Attributes: map[string]schema.Attribute{
					"cluster_id": schema.StringAttribute{
						Description:         "ClusterID identifies the local cluster.",
						MarkdownDescription: "ClusterID identifies the local cluster.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"leaders": schema.ListNestedAttribute{
						Description:         "Leaders include leader clusters known to the member clusters.",
						MarkdownDescription: "Leaders include leader clusters known to the member clusters.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"cluster_id": schema.StringAttribute{
									Description:         "Identify a leader cluster in the ClusterSet.",
									MarkdownDescription: "Identify a leader cluster in the ClusterSet.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"secret": schema.StringAttribute{
									Description:         "Name of the Secret resource in the member cluster, which stores the token to access the leader cluster's API server.",
									MarkdownDescription: "Name of the Secret resource in the member cluster, which stores the token to access the leader cluster's API server.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"server": schema.StringAttribute{
									Description:         "API server endpoint of the leader cluster. E.g. 'https://172.18.0.1:6443', 'https://example.com:6443'.",
									MarkdownDescription: "API server endpoint of the leader cluster. E.g. 'https://172.18.0.1:6443', 'https://example.com:6443'.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"namespace": schema.StringAttribute{
						Description:         "The leader cluster Namespace in which the ClusterSet is defined. Used in a member cluster.",
						MarkdownDescription: "The leader cluster Namespace in which the ClusterSet is defined. Used in a member cluster.",
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

func (r *MulticlusterCrdAntreaIoClusterSetV1Alpha2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_multicluster_crd_antrea_io_cluster_set_v1alpha2_manifest")

	var model MulticlusterCrdAntreaIoClusterSetV1Alpha2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("multicluster.crd.antrea.io/v1alpha2")
	model.Kind = pointer.String("ClusterSet")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
