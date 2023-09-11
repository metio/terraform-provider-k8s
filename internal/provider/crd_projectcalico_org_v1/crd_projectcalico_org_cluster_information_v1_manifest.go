/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package crd_projectcalico_org_v1

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
	_ datasource.DataSource = &CrdProjectcalicoOrgClusterInformationV1Manifest{}
)

func NewCrdProjectcalicoOrgClusterInformationV1Manifest() datasource.DataSource {
	return &CrdProjectcalicoOrgClusterInformationV1Manifest{}
}

type CrdProjectcalicoOrgClusterInformationV1Manifest struct{}

type CrdProjectcalicoOrgClusterInformationV1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		CalicoVersion  *string `tfsdk:"calico_version" json:"calicoVersion,omitempty"`
		ClusterGUID    *string `tfsdk:"cluster_guid" json:"clusterGUID,omitempty"`
		ClusterType    *string `tfsdk:"cluster_type" json:"clusterType,omitempty"`
		DatastoreReady *bool   `tfsdk:"datastore_ready" json:"datastoreReady,omitempty"`
		Variant        *string `tfsdk:"variant" json:"variant,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CrdProjectcalicoOrgClusterInformationV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_crd_projectcalico_org_cluster_information_v1_manifest"
}

func (r *CrdProjectcalicoOrgClusterInformationV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClusterInformation contains the cluster specific information.",
		MarkdownDescription: "ClusterInformation contains the cluster specific information.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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
				Description:         "ClusterInformationSpec contains the values of describing the cluster.",
				MarkdownDescription: "ClusterInformationSpec contains the values of describing the cluster.",
				Attributes: map[string]schema.Attribute{
					"calico_version": schema.StringAttribute{
						Description:         "CalicoVersion is the version of Calico that the cluster is running",
						MarkdownDescription: "CalicoVersion is the version of Calico that the cluster is running",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cluster_guid": schema.StringAttribute{
						Description:         "ClusterGUID is the GUID of the cluster",
						MarkdownDescription: "ClusterGUID is the GUID of the cluster",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cluster_type": schema.StringAttribute{
						Description:         "ClusterType describes the type of the cluster",
						MarkdownDescription: "ClusterType describes the type of the cluster",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"datastore_ready": schema.BoolAttribute{
						Description:         "DatastoreReady is used during significant datastore migrations to signal to components such as Felix that it should wait before accessing the datastore.",
						MarkdownDescription: "DatastoreReady is used during significant datastore migrations to signal to components such as Felix that it should wait before accessing the datastore.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"variant": schema.StringAttribute{
						Description:         "Variant declares which variant of Calico should be active.",
						MarkdownDescription: "Variant declares which variant of Calico should be active.",
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

func (r *CrdProjectcalicoOrgClusterInformationV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_crd_projectcalico_org_cluster_information_v1_manifest")

	var model CrdProjectcalicoOrgClusterInformationV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("crd.projectcalico.org/v1")
	model.Kind = pointer.String("ClusterInformation")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
