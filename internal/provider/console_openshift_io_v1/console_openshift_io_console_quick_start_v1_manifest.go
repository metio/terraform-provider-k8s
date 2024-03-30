/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package console_openshift_io_v1

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
	_ datasource.DataSource = &ConsoleOpenshiftIoConsoleQuickStartV1Manifest{}
)

func NewConsoleOpenshiftIoConsoleQuickStartV1Manifest() datasource.DataSource {
	return &ConsoleOpenshiftIoConsoleQuickStartV1Manifest{}
}

type ConsoleOpenshiftIoConsoleQuickStartV1Manifest struct{}

type ConsoleOpenshiftIoConsoleQuickStartV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		AccessReviewResources *[]struct {
			Group       *string `tfsdk:"group" json:"group,omitempty"`
			Name        *string `tfsdk:"name" json:"name,omitempty"`
			Namespace   *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Resource    *string `tfsdk:"resource" json:"resource,omitempty"`
			Subresource *string `tfsdk:"subresource" json:"subresource,omitempty"`
			Verb        *string `tfsdk:"verb" json:"verb,omitempty"`
			Version     *string `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"access_review_resources" json:"accessReviewResources,omitempty"`
		Conclusion      *string   `tfsdk:"conclusion" json:"conclusion,omitempty"`
		Description     *string   `tfsdk:"description" json:"description,omitempty"`
		DisplayName     *string   `tfsdk:"display_name" json:"displayName,omitempty"`
		DurationMinutes *int64    `tfsdk:"duration_minutes" json:"durationMinutes,omitempty"`
		Icon            *string   `tfsdk:"icon" json:"icon,omitempty"`
		Introduction    *string   `tfsdk:"introduction" json:"introduction,omitempty"`
		NextQuickStart  *[]string `tfsdk:"next_quick_start" json:"nextQuickStart,omitempty"`
		Prerequisites   *[]string `tfsdk:"prerequisites" json:"prerequisites,omitempty"`
		Tags            *[]string `tfsdk:"tags" json:"tags,omitempty"`
		Tasks           *[]struct {
			Description *string `tfsdk:"description" json:"description,omitempty"`
			Review      *struct {
				FailedTaskHelp *string `tfsdk:"failed_task_help" json:"failedTaskHelp,omitempty"`
				Instructions   *string `tfsdk:"instructions" json:"instructions,omitempty"`
			} `tfsdk:"review" json:"review,omitempty"`
			Summary *struct {
				Failed  *string `tfsdk:"failed" json:"failed,omitempty"`
				Success *string `tfsdk:"success" json:"success,omitempty"`
			} `tfsdk:"summary" json:"summary,omitempty"`
			Title *string `tfsdk:"title" json:"title,omitempty"`
		} `tfsdk:"tasks" json:"tasks,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ConsoleOpenshiftIoConsoleQuickStartV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_console_openshift_io_console_quick_start_v1_manifest"
}

func (r *ConsoleOpenshiftIoConsoleQuickStartV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ConsoleQuickStart is an extension for guiding user through various workflows in the OpenShift web console.  Compatibility level 2: Stable within a major release for a minimum of 9 months or 3 minor releases (whichever is longer).",
		MarkdownDescription: "ConsoleQuickStart is an extension for guiding user through various workflows in the OpenShift web console.  Compatibility level 2: Stable within a major release for a minimum of 9 months or 3 minor releases (whichever is longer).",
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
				Description:         "ConsoleQuickStartSpec is the desired quick start configuration.",
				MarkdownDescription: "ConsoleQuickStartSpec is the desired quick start configuration.",
				Attributes: map[string]schema.Attribute{
					"access_review_resources": schema.ListNestedAttribute{
						Description:         "accessReviewResources contains a list of resources that the user's access will be reviewed against in order for the user to complete the Quick Start. The Quick Start will be hidden if any of the access reviews fail.",
						MarkdownDescription: "accessReviewResources contains a list of resources that the user's access will be reviewed against in order for the user to complete the Quick Start. The Quick Start will be hidden if any of the access reviews fail.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"group": schema.StringAttribute{
									Description:         "Group is the API Group of the Resource.  '*' means all.",
									MarkdownDescription: "Group is the API Group of the Resource.  '*' means all.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name is the name of the resource being requested for a 'get' or deleted for a 'delete'. '' (empty) means all.",
									MarkdownDescription: "Name is the name of the resource being requested for a 'get' or deleted for a 'delete'. '' (empty) means all.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace is the namespace of the action being requested.  Currently, there is no distinction between no namespace and all namespaces '' (empty) is defaulted for LocalSubjectAccessReviews '' (empty) is empty for cluster-scoped resources '' (empty) means 'all' for namespace scoped resources from a SubjectAccessReview or SelfSubjectAccessReview",
									MarkdownDescription: "Namespace is the namespace of the action being requested.  Currently, there is no distinction between no namespace and all namespaces '' (empty) is defaulted for LocalSubjectAccessReviews '' (empty) is empty for cluster-scoped resources '' (empty) means 'all' for namespace scoped resources from a SubjectAccessReview or SelfSubjectAccessReview",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"resource": schema.StringAttribute{
									Description:         "Resource is one of the existing resource types.  '*' means all.",
									MarkdownDescription: "Resource is one of the existing resource types.  '*' means all.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"subresource": schema.StringAttribute{
									Description:         "Subresource is one of the existing resource types.  '' means none.",
									MarkdownDescription: "Subresource is one of the existing resource types.  '' means none.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"verb": schema.StringAttribute{
									Description:         "Verb is a kubernetes resource API verb, like: get, list, watch, create, update, delete, proxy.  '*' means all.",
									MarkdownDescription: "Verb is a kubernetes resource API verb, like: get, list, watch, create, update, delete, proxy.  '*' means all.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"version": schema.StringAttribute{
									Description:         "Version is the API Version of the Resource.  '*' means all.",
									MarkdownDescription: "Version is the API Version of the Resource.  '*' means all.",
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

					"conclusion": schema.StringAttribute{
						Description:         "conclusion sums up the Quick Start and suggests the possible next steps. (includes markdown)",
						MarkdownDescription: "conclusion sums up the Quick Start and suggests the possible next steps. (includes markdown)",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"description": schema.StringAttribute{
						Description:         "description is the description of the Quick Start. (includes markdown)",
						MarkdownDescription: "description is the description of the Quick Start. (includes markdown)",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
							stringvalidator.LengthAtMost(256),
						},
					},

					"display_name": schema.StringAttribute{
						Description:         "displayName is the display name of the Quick Start.",
						MarkdownDescription: "displayName is the display name of the Quick Start.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
						},
					},

					"duration_minutes": schema.Int64Attribute{
						Description:         "durationMinutes describes approximately how many minutes it will take to complete the Quick Start.",
						MarkdownDescription: "durationMinutes describes approximately how many minutes it will take to complete the Quick Start.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},

					"icon": schema.StringAttribute{
						Description:         "icon is a base64 encoded image that will be displayed beside the Quick Start display name. The icon should be an vector image for easy scaling. The size of the icon should be 40x40.",
						MarkdownDescription: "icon is a base64 encoded image that will be displayed beside the Quick Start display name. The icon should be an vector image for easy scaling. The size of the icon should be 40x40.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"introduction": schema.StringAttribute{
						Description:         "introduction describes the purpose of the Quick Start. (includes markdown)",
						MarkdownDescription: "introduction describes the purpose of the Quick Start. (includes markdown)",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
						},
					},

					"next_quick_start": schema.ListAttribute{
						Description:         "nextQuickStart is a list of the following Quick Starts, suggested for the user to try.",
						MarkdownDescription: "nextQuickStart is a list of the following Quick Starts, suggested for the user to try.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"prerequisites": schema.ListAttribute{
						Description:         "prerequisites contains all prerequisites that need to be met before taking a Quick Start. (includes markdown)",
						MarkdownDescription: "prerequisites contains all prerequisites that need to be met before taking a Quick Start. (includes markdown)",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tags": schema.ListAttribute{
						Description:         "tags is a list of strings that describe the Quick Start.",
						MarkdownDescription: "tags is a list of strings that describe the Quick Start.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tasks": schema.ListNestedAttribute{
						Description:         "tasks is the list of steps the user has to perform to complete the Quick Start.",
						MarkdownDescription: "tasks is the list of steps the user has to perform to complete the Quick Start.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"description": schema.StringAttribute{
									Description:         "description describes the steps needed to complete the task. (includes markdown)",
									MarkdownDescription: "description describes the steps needed to complete the task. (includes markdown)",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
									},
								},

								"review": schema.SingleNestedAttribute{
									Description:         "review contains instructions to validate the task is complete. The user will select 'Yes' or 'No'. using a radio button, which indicates whether the step was completed successfully.",
									MarkdownDescription: "review contains instructions to validate the task is complete. The user will select 'Yes' or 'No'. using a radio button, which indicates whether the step was completed successfully.",
									Attributes: map[string]schema.Attribute{
										"failed_task_help": schema.StringAttribute{
											Description:         "failedTaskHelp contains suggestions for a failed task review and is shown at the end of task. (includes markdown)",
											MarkdownDescription: "failedTaskHelp contains suggestions for a failed task review and is shown at the end of task. (includes markdown)",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
											},
										},

										"instructions": schema.StringAttribute{
											Description:         "instructions contains steps that user needs to take in order to validate his work after going through a task. (includes markdown)",
											MarkdownDescription: "instructions contains steps that user needs to take in order to validate his work after going through a task. (includes markdown)",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"summary": schema.SingleNestedAttribute{
									Description:         "summary contains information about the passed step.",
									MarkdownDescription: "summary contains information about the passed step.",
									Attributes: map[string]schema.Attribute{
										"failed": schema.StringAttribute{
											Description:         "failed briefly describes the unsuccessfully passed task. (includes markdown)",
											MarkdownDescription: "failed briefly describes the unsuccessfully passed task. (includes markdown)",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
												stringvalidator.LengthAtMost(128),
											},
										},

										"success": schema.StringAttribute{
											Description:         "success describes the succesfully passed task.",
											MarkdownDescription: "success describes the succesfully passed task.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"title": schema.StringAttribute{
									Description:         "title describes the task and is displayed as a step heading.",
									MarkdownDescription: "title describes the task and is displayed as a step heading.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
									},
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *ConsoleOpenshiftIoConsoleQuickStartV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_console_openshift_io_console_quick_start_v1_manifest")

	var model ConsoleOpenshiftIoConsoleQuickStartV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("console.openshift.io/v1")
	model.Kind = pointer.String("ConsoleQuickStart")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
