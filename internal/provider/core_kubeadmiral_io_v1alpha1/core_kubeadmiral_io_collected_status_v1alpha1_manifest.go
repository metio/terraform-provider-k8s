/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package core_kubeadmiral_io_v1alpha1

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
	_ datasource.DataSource = &CoreKubeadmiralIoCollectedStatusV1Alpha1Manifest{}
)

func NewCoreKubeadmiralIoCollectedStatusV1Alpha1Manifest() datasource.DataSource {
	return &CoreKubeadmiralIoCollectedStatusV1Alpha1Manifest{}
}

type CoreKubeadmiralIoCollectedStatusV1Alpha1Manifest struct{}

type CoreKubeadmiralIoCollectedStatusV1Alpha1ManifestData struct {
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

	Clusters *[]struct {
		Cluster         *string            `tfsdk:"cluster" json:"cluster,omitempty"`
		CollectedFields *map[string]string `tfsdk:"collected_fields" json:"collectedFields,omitempty"`
		Error           *string            `tfsdk:"error" json:"error,omitempty"`
	} `tfsdk:"clusters" json:"clusters,omitempty"`
	LastUpdateTime *string `tfsdk:"last_update_time" json:"lastUpdateTime,omitempty"`
}

func (r *CoreKubeadmiralIoCollectedStatusV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_core_kubeadmiral_io_collected_status_v1alpha1_manifest"
}

func (r *CoreKubeadmiralIoCollectedStatusV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CollectedStatus stores the collected fields of Kubernetes objects from member clusters, that are propagated by a FederatedObject.",
		MarkdownDescription: "CollectedStatus stores the collected fields of Kubernetes objects from member clusters, that are propagated by a FederatedObject.",
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

			"clusters": schema.ListNestedAttribute{
				Description:         "Clusters is the list of member clusters and collected fields for its propagated Kubernetes object.",
				MarkdownDescription: "Clusters is the list of member clusters and collected fields for its propagated Kubernetes object.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"cluster": schema.StringAttribute{
							Description:         "Cluster is the name of the member cluster.",
							MarkdownDescription: "Cluster is the name of the member cluster.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},

						"collected_fields": schema.MapAttribute{
							Description:         "CollectedFields is the the set of fields collected for the Kubernetes object.",
							MarkdownDescription: "CollectedFields is the the set of fields collected for the Kubernetes object.",
							ElementType:         types.StringType,
							Required:            true,
							Optional:            false,
							Computed:            false,
						},

						"error": schema.StringAttribute{
							Description:         "Error records any errors encountered while collecting fields from the cluster.",
							MarkdownDescription: "Error records any errors encountered while collecting fields from the cluster.",
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

			"last_update_time": schema.StringAttribute{
				Description:         "LastUpdateTime is the last time that a collection was performed.",
				MarkdownDescription: "LastUpdateTime is the last time that a collection was performed.",
				Required:            true,
				Optional:            false,
				Computed:            false,
				Validators: []validator.String{
					validators.DateTime64Validator(),
				},
			},
		},
	}
}

func (r *CoreKubeadmiralIoCollectedStatusV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_core_kubeadmiral_io_collected_status_v1alpha1_manifest")

	var model CoreKubeadmiralIoCollectedStatusV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("core.kubeadmiral.io/v1alpha1")
	model.Kind = pointer.String("CollectedStatus")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
