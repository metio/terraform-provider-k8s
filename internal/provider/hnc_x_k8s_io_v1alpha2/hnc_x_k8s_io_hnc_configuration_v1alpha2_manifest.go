/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package hnc_x_k8s_io_v1alpha2

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
	_ datasource.DataSource = &HncXK8SIoHncconfigurationV1Alpha2Manifest{}
)

func NewHncXK8SIoHncconfigurationV1Alpha2Manifest() datasource.DataSource {
	return &HncXK8SIoHncconfigurationV1Alpha2Manifest{}
}

type HncXK8SIoHncconfigurationV1Alpha2Manifest struct{}

type HncXK8SIoHncconfigurationV1Alpha2ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Resources *[]struct {
			Group    *string `tfsdk:"group" json:"group,omitempty"`
			Mode     *string `tfsdk:"mode" json:"mode,omitempty"`
			Resource *string `tfsdk:"resource" json:"resource,omitempty"`
		} `tfsdk:"resources" json:"resources,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *HncXK8SIoHncconfigurationV1Alpha2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_hnc_x_k8s_io_hnc_configuration_v1alpha2_manifest"
}

func (r *HncXK8SIoHncconfigurationV1Alpha2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "HNCConfiguration is a cluster-wide configuration for HNC as a whole. See details in http://bit.ly/hnc-type-configuration",
		MarkdownDescription: "HNCConfiguration is a cluster-wide configuration for HNC as a whole. See details in http://bit.ly/hnc-type-configuration",
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
				Description:         "HNCConfigurationSpec defines the desired state of HNC configuration.",
				MarkdownDescription: "HNCConfigurationSpec defines the desired state of HNC configuration.",
				Attributes: map[string]schema.Attribute{
					"resources": schema.ListNestedAttribute{
						Description:         "Resources defines the cluster-wide settings for resource synchronization. Note that 'roles' and 'rolebindings' are pre-configured by HNC with 'Propagate' mode and are omitted in the spec. Any configuration of 'roles' or 'rolebindings' are not allowed. To learn more, see https://github.com/kubernetes-sigs/hierarchical-namespaces/blob/master/docs/user-guide/how-to.md#admin-types",
						MarkdownDescription: "Resources defines the cluster-wide settings for resource synchronization. Note that 'roles' and 'rolebindings' are pre-configured by HNC with 'Propagate' mode and are omitted in the spec. Any configuration of 'roles' or 'rolebindings' are not allowed. To learn more, see https://github.com/kubernetes-sigs/hierarchical-namespaces/blob/master/docs/user-guide/how-to.md#admin-types",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"group": schema.StringAttribute{
									Description:         "Group of the resource defined below. This is used to unambiguously identify the resource. It may be omitted for core resources (e.g. 'secrets').",
									MarkdownDescription: "Group of the resource defined below. This is used to unambiguously identify the resource. It may be omitted for core resources (e.g. 'secrets').",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"mode": schema.StringAttribute{
									Description:         "Synchronization mode of the kind. If the field is empty, it will be treated as 'Propagate'.",
									MarkdownDescription: "Synchronization mode of the kind. If the field is empty, it will be treated as 'Propagate'.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("Propagate", "Ignore", "Remove", "AllowPropagate"),
									},
								},

								"resource": schema.StringAttribute{
									Description:         "Resource to be configured.",
									MarkdownDescription: "Resource to be configured.",
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

func (r *HncXK8SIoHncconfigurationV1Alpha2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_hnc_x_k8s_io_hnc_configuration_v1alpha2_manifest")

	var model HncXK8SIoHncconfigurationV1Alpha2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("hnc.x-k8s.io/v1alpha2")
	model.Kind = pointer.String("HNCConfiguration")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
