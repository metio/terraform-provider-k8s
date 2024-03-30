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
	_ datasource.DataSource = &ClustersClusternetIoClusterRegistrationRequestV1Beta1Manifest{}
)

func NewClustersClusternetIoClusterRegistrationRequestV1Beta1Manifest() datasource.DataSource {
	return &ClustersClusternetIoClusterRegistrationRequestV1Beta1Manifest{}
}

type ClustersClusternetIoClusterRegistrationRequestV1Beta1Manifest struct{}

type ClustersClusternetIoClusterRegistrationRequestV1Beta1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ClusterId        *string            `tfsdk:"cluster_id" json:"clusterId,omitempty"`
		ClusterLabels    *map[string]string `tfsdk:"cluster_labels" json:"clusterLabels,omitempty"`
		ClusterName      *string            `tfsdk:"cluster_name" json:"clusterName,omitempty"`
		ClusterNamespace *string            `tfsdk:"cluster_namespace" json:"clusterNamespace,omitempty"`
		ClusterType      *string            `tfsdk:"cluster_type" json:"clusterType,omitempty"`
		SyncMode         *string            `tfsdk:"sync_mode" json:"syncMode,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ClustersClusternetIoClusterRegistrationRequestV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_clusters_clusternet_io_cluster_registration_request_v1beta1_manifest"
}

func (r *ClustersClusternetIoClusterRegistrationRequestV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClusterRegistrationRequest is the Schema for the clusterregistrationrequests API",
		MarkdownDescription: "ClusterRegistrationRequest is the Schema for the clusterregistrationrequests API",
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
				Description:         "ClusterRegistrationRequestSpec defines the desired state of ClusterRegistrationRequest",
				MarkdownDescription: "ClusterRegistrationRequestSpec defines the desired state of ClusterRegistrationRequest",
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

					"cluster_labels": schema.MapAttribute{
						Description:         "ClusterLabels is the labels of the child cluster.",
						MarkdownDescription: "ClusterLabels is the labels of the child cluster.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cluster_name": schema.StringAttribute{
						Description:         "ClusterName is the cluster name.a lower case alphanumeric characters or '-', and must start and end with an alphanumeric character",
						MarkdownDescription: "ClusterName is the cluster name.a lower case alphanumeric characters or '-', and must start and end with an alphanumeric character",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(30),
							stringvalidator.RegexMatches(regexp.MustCompile(`[a-z0-9]([-a-z0-9]*[a-z0-9])?([a-z0-9]([-a-z0-9]*[a-z0-9]))*`), ""),
						},
					},

					"cluster_namespace": schema.StringAttribute{
						Description:         "ClusterNamespace is the dedicated namespace of the cluster.",
						MarkdownDescription: "ClusterNamespace is the dedicated namespace of the cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(63),
							stringvalidator.RegexMatches(regexp.MustCompile(`[a-z0-9]([-a-z0-9]*[a-z0-9])?`), ""),
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
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Push", "Pull", "Dual"),
						},
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *ClustersClusternetIoClusterRegistrationRequestV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_clusters_clusternet_io_cluster_registration_request_v1beta1_manifest")

	var model ClustersClusternetIoClusterRegistrationRequestV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("clusters.clusternet.io/v1beta1")
	model.Kind = pointer.String("ClusterRegistrationRequest")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
