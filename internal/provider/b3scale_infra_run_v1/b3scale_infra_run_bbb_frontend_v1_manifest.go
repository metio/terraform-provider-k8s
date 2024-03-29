/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package b3scale_infra_run_v1

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
	_ datasource.DataSource = &B3ScaleInfraRunBbbfrontendV1Manifest{}
)

func NewB3ScaleInfraRunBbbfrontendV1Manifest() datasource.DataSource {
	return &B3ScaleInfraRunBbbfrontendV1Manifest{}
}

type B3ScaleInfraRunBbbfrontendV1Manifest struct{}

type B3ScaleInfraRunBbbfrontendV1ManifestData struct {
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
		Credentials *struct {
			Key       *string `tfsdk:"key" json:"key,omitempty"`
			SecretRef *struct {
				Key  *string `tfsdk:"key" json:"key,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
		} `tfsdk:"credentials" json:"credentials,omitempty"`
		Settings *struct {
			Create_default_params  *map[string]string `tfsdk:"create_default_params" json:"create_default_params,omitempty"`
			Create_override_params *map[string]string `tfsdk:"create_override_params" json:"create_override_params,omitempty"`
			Default_presentation   *struct {
				Force *bool   `tfsdk:"force" json:"force,omitempty"`
				Url   *string `tfsdk:"url" json:"url,omitempty"`
			} `tfsdk:"default_presentation" json:"default_presentation,omitempty"`
			Required_tags *[]string `tfsdk:"required_tags" json:"required_tags,omitempty"`
		} `tfsdk:"settings" json:"settings,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *B3ScaleInfraRunBbbfrontendV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_b3scale_infra_run_bbb_frontend_v1_manifest"
}

func (r *B3ScaleInfraRunBbbfrontendV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "Desired state of the BBBFrontend resource.",
				MarkdownDescription: "Desired state of the BBBFrontend resource.",
				Attributes: map[string]schema.Attribute{
					"credentials": schema.SingleNestedAttribute{
						Description:         "Predefined credentials for the B3scale instance",
						MarkdownDescription: "Predefined credentials for the B3scale instance",
						Attributes: map[string]schema.Attribute{
							"key": schema.StringAttribute{
								Description:         "Predefined key for B3scale instance, will be defined if not set",
								MarkdownDescription: "Predefined key for B3scale instance, will be defined if not set",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secret_ref": schema.SingleNestedAttribute{
								Description:         "SecretRef is a reference to a key in a Secret resource containing the key to connect to the BBB instance.",
								MarkdownDescription: "SecretRef is a reference to a key in a Secret resource containing the key to connect to the BBB instance.",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "The key of the entry in the Secret resource's 'data' field to be used.",
										MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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

					"settings": schema.SingleNestedAttribute{
						Description:         "Settings defines the B3Scale instance settings",
						MarkdownDescription: "Settings defines the B3Scale instance settings",
						Attributes: map[string]schema.Attribute{
							"create_default_params": schema.MapAttribute{
								Description:         "See https://github.com/b3scale/b3scale#configure-create-parameter-defaults-and-overrides",
								MarkdownDescription: "See https://github.com/b3scale/b3scale#configure-create-parameter-defaults-and-overrides",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"create_override_params": schema.MapAttribute{
								Description:         "See https://github.com/b3scale/b3scale#configure-create-parameter-defaults-and-overrides",
								MarkdownDescription: "See https://github.com/b3scale/b3scale#configure-create-parameter-defaults-and-overrides",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"default_presentation": schema.SingleNestedAttribute{
								Description:         "See https://github.com/b3scale/b3scale#middleware-configuration",
								MarkdownDescription: "See https://github.com/b3scale/b3scale#middleware-configuration",
								Attributes: map[string]schema.Attribute{
									"force": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"url": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"required_tags": schema.ListAttribute{
								Description:         "See https://github.com/b3scale/b3scale#middleware-configuration",
								MarkdownDescription: "See https://github.com/b3scale/b3scale#middleware-configuration",
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *B3ScaleInfraRunBbbfrontendV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_b3scale_infra_run_bbb_frontend_v1_manifest")

	var model B3ScaleInfraRunBbbfrontendV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("b3scale.infra.run/v1")
	model.Kind = pointer.String("BBBFrontend")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
