/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package app_terraform_io_v1alpha2

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &AppTerraformIoModuleV1Alpha2Manifest{}
)

func NewAppTerraformIoModuleV1Alpha2Manifest() datasource.DataSource {
	return &AppTerraformIoModuleV1Alpha2Manifest{}
}

type AppTerraformIoModuleV1Alpha2Manifest struct{}

type AppTerraformIoModuleV1Alpha2ManifestData struct {
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
		DestroyOnDeletion *bool `tfsdk:"destroy_on_deletion" json:"destroyOnDeletion,omitempty"`
		Module            *struct {
			Source  *string `tfsdk:"source" json:"source,omitempty"`
			Version *string `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"module" json:"module,omitempty"`
		Name         *string `tfsdk:"name" json:"name,omitempty"`
		Organization *string `tfsdk:"organization" json:"organization,omitempty"`
		Outputs      *[]struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Sensitive *bool   `tfsdk:"sensitive" json:"sensitive,omitempty"`
		} `tfsdk:"outputs" json:"outputs,omitempty"`
		RestartedAt *string `tfsdk:"restarted_at" json:"restartedAt,omitempty"`
		Token       *struct {
			SecretKeyRef *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
		} `tfsdk:"token" json:"token,omitempty"`
		Variables *[]struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"variables" json:"variables,omitempty"`
		Workspace *struct {
			Id   *string `tfsdk:"id" json:"id,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"workspace" json:"workspace,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppTerraformIoModuleV1Alpha2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_app_terraform_io_module_v1alpha2_manifest"
}

func (r *AppTerraformIoModuleV1Alpha2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Module is the Schema for the modules API Module implements the API-driven Run Workflow More information: - https://developer.hashicorp.com/terraform/cloud-docs/run/api",
		MarkdownDescription: "Module is the Schema for the modules API Module implements the API-driven Run Workflow More information: - https://developer.hashicorp.com/terraform/cloud-docs/run/api",
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
				Description:         "ModuleSpec defines the desired state of Module.",
				MarkdownDescription: "ModuleSpec defines the desired state of Module.",
				Attributes: map[string]schema.Attribute{
					"destroy_on_deletion": schema.BoolAttribute{
						Description:         "Specify whether or not to execute a Destroy run when the object is deleted from the Kubernetes. Default: 'false'.",
						MarkdownDescription: "Specify whether or not to execute a Destroy run when the object is deleted from the Kubernetes. Default: 'false'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"module": schema.SingleNestedAttribute{
						Description:         "Module source and version to execute.",
						MarkdownDescription: "Module source and version to execute.",
						Attributes: map[string]schema.Attribute{
							"source": schema.StringAttribute{
								Description:         "Non local Terraform module source. More information: - https://developer.hashicorp.com/terraform/language/modules/sources",
								MarkdownDescription: "Non local Terraform module source. More information: - https://developer.hashicorp.com/terraform/language/modules/sources",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},

							"version": schema.StringAttribute{
								Description:         "Terraform module version.",
								MarkdownDescription: "Terraform module version.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"name": schema.StringAttribute{
						Description:         "Name of the module that will be uploaded and executed. Default: 'this'.",
						MarkdownDescription: "Name of the module that will be uploaded and executed. Default: 'this'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
						},
					},

					"organization": schema.StringAttribute{
						Description:         "Organization name where the Workspace will be created. More information: - https://developer.hashicorp.com/terraform/cloud-docs/users-teams-organizations/organizations",
						MarkdownDescription: "Organization name where the Workspace will be created. More information: - https://developer.hashicorp.com/terraform/cloud-docs/users-teams-organizations/organizations",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
						},
					},

					"outputs": schema.ListNestedAttribute{
						Description:         "Module outputs to store in ConfigMap(non-sensitive) or Secret(sensitive).",
						MarkdownDescription: "Module outputs to store in ConfigMap(non-sensitive) or Secret(sensitive).",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Output name must match with the module output.",
									MarkdownDescription: "Output name must match with the module output.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
									},
								},

								"sensitive": schema.BoolAttribute{
									Description:         "Specify whether or not the output is sensitive. Default: 'false'.",
									MarkdownDescription: "Specify whether or not the output is sensitive. Default: 'false'.",
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

					"restarted_at": schema.StringAttribute{
						Description:         "Allows executing a new Run without changing any Workspace or Module attributes. Example: kubectl patch <KIND> <NAME> --type=merge --patch '{'spec': {'restartedAt': '''date -u -Iseconds'''}}'",
						MarkdownDescription: "Allows executing a new Run without changing any Workspace or Module attributes. Example: kubectl patch <KIND> <NAME> --type=merge --patch '{'spec': {'restartedAt': '''date -u -Iseconds'''}}'",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
						},
					},

					"token": schema.SingleNestedAttribute{
						Description:         "API Token to be used for API calls.",
						MarkdownDescription: "API Token to be used for API calls.",
						Attributes: map[string]schema.Attribute{
							"secret_key_ref": schema.SingleNestedAttribute{
								Description:         "Selects a key of a secret in the workspace's namespace",
								MarkdownDescription: "Selects a key of a secret in the workspace's namespace",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "The key of the secret to select from.  Must be a valid secret key.",
										MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"optional": schema.BoolAttribute{
										Description:         "Specify whether the Secret or its key must be defined",
										MarkdownDescription: "Specify whether the Secret or its key must be defined",
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
						Required: true,
						Optional: false,
						Computed: false,
					},

					"variables": schema.ListNestedAttribute{
						Description:         "Variables to pass to the module, they must exist in the Workspace.",
						MarkdownDescription: "Variables to pass to the module, they must exist in the Workspace.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Variable name must exist in the Workspace.",
									MarkdownDescription: "Variable name must exist in the Workspace.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
									},
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"workspace": schema.SingleNestedAttribute{
						Description:         "Workspace to execute the module.",
						MarkdownDescription: "Workspace to execute the module.",
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Description:         "Module Workspace ID. Must match pattern: '^ws-[a-zA-Z0-9]+$'",
								MarkdownDescription: "Module Workspace ID. Must match pattern: '^ws-[a-zA-Z0-9]+$'",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^ws-[a-zA-Z0-9]+$`), ""),
								},
							},

							"name": schema.StringAttribute{
								Description:         "Module Workspace Name.",
								MarkdownDescription: "Module Workspace Name.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
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

func (r *AppTerraformIoModuleV1Alpha2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_app_terraform_io_module_v1alpha2_manifest")

	var model AppTerraformIoModuleV1Alpha2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("app.terraform.io/v1alpha2")
	model.Kind = pointer.String("Module")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
