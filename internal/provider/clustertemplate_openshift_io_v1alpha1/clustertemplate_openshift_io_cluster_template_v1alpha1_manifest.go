/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package clustertemplate_openshift_io_v1alpha1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	_ datasource.DataSource = &ClustertemplateOpenshiftIoClusterTemplateV1Alpha1Manifest{}
)

func NewClustertemplateOpenshiftIoClusterTemplateV1Alpha1Manifest() datasource.DataSource {
	return &ClustertemplateOpenshiftIoClusterTemplateV1Alpha1Manifest{}
}

type ClustertemplateOpenshiftIoClusterTemplateV1Alpha1Manifest struct{}

type ClustertemplateOpenshiftIoClusterTemplateV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ClusterDefinition       *string   `tfsdk:"cluster_definition" json:"clusterDefinition,omitempty"`
		ClusterSetup            *[]string `tfsdk:"cluster_setup" json:"clusterSetup,omitempty"`
		Cost                    *int64    `tfsdk:"cost" json:"cost,omitempty"`
		SkipClusterRegistration *bool     `tfsdk:"skip_cluster_registration" json:"skipClusterRegistration,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ClustertemplateOpenshiftIoClusterTemplateV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_clustertemplate_openshift_io_cluster_template_v1alpha1_manifest"
}

func (r *ClustertemplateOpenshiftIoClusterTemplateV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Template of a cluster - both installation and post-install setup are defined as ArgoCD application spec. Any application source is supported - typically a Helm chart",
		MarkdownDescription: "Template of a cluster - both installation and post-install setup are defined as ArgoCD application spec. Any application source is supported - typically a Helm chart",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"cluster_definition": schema.StringAttribute{
						Description:         "ArgoCD applicationset name which is used for installation of the cluster",
						MarkdownDescription: "ArgoCD applicationset name which is used for installation of the cluster",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"cluster_setup": schema.ListAttribute{
						Description:         "Array of ArgoCD applicationset names which are used for post installation setup of the cluster",
						MarkdownDescription: "Array of ArgoCD applicationset names which are used for post installation setup of the cluster",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cost": schema.Int64Attribute{
						Description:         "Cost of the cluster, used for quotas",
						MarkdownDescription: "Cost of the cluster, used for quotas",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},

					"skip_cluster_registration": schema.BoolAttribute{
						Description:         "Skip the registration of the cluster to the hub cluster",
						MarkdownDescription: "Skip the registration of the cluster to the hub cluster",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *ClustertemplateOpenshiftIoClusterTemplateV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_clustertemplate_openshift_io_cluster_template_v1alpha1_manifest")

	var model ClustertemplateOpenshiftIoClusterTemplateV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("clustertemplate.openshift.io/v1alpha1")
	model.Kind = pointer.String("ClusterTemplate")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
