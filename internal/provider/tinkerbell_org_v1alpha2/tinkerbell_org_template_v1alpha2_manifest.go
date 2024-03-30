/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package tinkerbell_org_v1alpha2

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
	_ datasource.DataSource = &TinkerbellOrgTemplateV1Alpha2Manifest{}
)

func NewTinkerbellOrgTemplateV1Alpha2Manifest() datasource.DataSource {
	return &TinkerbellOrgTemplateV1Alpha2Manifest{}
}

type TinkerbellOrgTemplateV1Alpha2Manifest struct{}

type TinkerbellOrgTemplateV1Alpha2ManifestData struct {
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
		Actions *[]struct {
			Args       *[]string          `tfsdk:"args" json:"args,omitempty"`
			Cmd        *string            `tfsdk:"cmd" json:"cmd,omitempty"`
			Env        *map[string]string `tfsdk:"env" json:"env,omitempty"`
			Image      *string            `tfsdk:"image" json:"image,omitempty"`
			Name       *string            `tfsdk:"name" json:"name,omitempty"`
			Namespaces *struct {
				Network *string `tfsdk:"network" json:"network,omitempty"`
				Pid     *int64  `tfsdk:"pid" json:"pid,omitempty"`
			} `tfsdk:"namespaces" json:"namespaces,omitempty"`
			Volumes *[]string `tfsdk:"volumes" json:"volumes,omitempty"`
		} `tfsdk:"actions" json:"actions,omitempty"`
		Env     *map[string]string `tfsdk:"env" json:"env,omitempty"`
		Volumes *[]string          `tfsdk:"volumes" json:"volumes,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TinkerbellOrgTemplateV1Alpha2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_tinkerbell_org_template_v1alpha2_manifest"
}

func (r *TinkerbellOrgTemplateV1Alpha2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Template defines a set of actions to be run on a target machine. The template is renderedprior to execution where it is exposed to Hardware and user defined data. Most fields within theTemplateSpec may contain templates values excluding .TemplateSpec.Actions[].Name.See https://pkg.go.dev/text/template for more details.",
		MarkdownDescription: "Template defines a set of actions to be run on a target machine. The template is renderedprior to execution where it is exposed to Hardware and user defined data. Most fields within theTemplateSpec may contain templates values excluding .TemplateSpec.Actions[].Name.See https://pkg.go.dev/text/template for more details.",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"actions": schema.ListNestedAttribute{
						Description:         "Actions defines the set of actions to be run on a target machine. Actions are run sequentiallyin the order they are specified. At least 1 action must be specified. Names of actionsmust be unique within a Template.",
						MarkdownDescription: "Actions defines the set of actions to be run on a target machine. Actions are run sequentiallyin the order they are specified. At least 1 action must be specified. Names of actionsmust be unique within a Template.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"args": schema.ListAttribute{
									Description:         "Args are a set of arguments to be passed to the command executed by the container onlaunch.",
									MarkdownDescription: "Args are a set of arguments to be passed to the command executed by the container onlaunch.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"cmd": schema.StringAttribute{
									Description:         "Cmd defines the command to use when launching the image. It overrides the default commandof the action. It must be a unix path to an executable program.",
									MarkdownDescription: "Cmd defines the command to use when launching the image. It overrides the default commandof the action. It must be a unix path to an executable program.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.RegexMatches(regexp.MustCompile(`^(/[^/ ]*)+/?$`), ""),
									},
								},

								"env": schema.MapAttribute{
									Description:         "Env defines environment variables used when launching the container.",
									MarkdownDescription: "Env defines environment variables used when launching the container.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"image": schema.StringAttribute{
									Description:         "Image is an OCI image.",
									MarkdownDescription: "Image is an OCI image.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name is a name for the action.",
									MarkdownDescription: "Name is a name for the action.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespaces": schema.SingleNestedAttribute{
									Description:         "Namespace defines the Linux namespaces this container should execute in.",
									MarkdownDescription: "Namespace defines the Linux namespaces this container should execute in.",
									Attributes: map[string]schema.Attribute{
										"network": schema.StringAttribute{
											Description:         "Network defines the network namespace.",
											MarkdownDescription: "Network defines the network namespace.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"pid": schema.Int64Attribute{
											Description:         "PID defines the PID namespace",
											MarkdownDescription: "PID defines the PID namespace",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"volumes": schema.ListAttribute{
									Description:         "Volumes defines the volumes to mount into the container.",
									MarkdownDescription: "Volumes defines the volumes to mount into the container.",
									ElementType:         types.StringType,
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

					"env": schema.MapAttribute{
						Description:         "Env defines environment variables to be available in all actions. If an action specifiesthe same environment variable it will take precedence.",
						MarkdownDescription: "Env defines environment variables to be available in all actions. If an action specifiesthe same environment variable it will take precedence.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"volumes": schema.ListAttribute{
						Description:         "Volumes to be mounted on all actions. If an action specifies the same volume it will takeprecedence.",
						MarkdownDescription: "Volumes to be mounted on all actions. If an action specifies the same volume it will takeprecedence.",
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
	}
}

func (r *TinkerbellOrgTemplateV1Alpha2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_tinkerbell_org_template_v1alpha2_manifest")

	var model TinkerbellOrgTemplateV1Alpha2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("tinkerbell.org/v1alpha2")
	model.Kind = pointer.String("Template")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
