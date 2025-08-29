/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package apiextensions_crossplane_io_v1

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
	_ datasource.DataSource = &ApiextensionsCrossplaneIoCompositionRevisionV1Manifest{}
)

func NewApiextensionsCrossplaneIoCompositionRevisionV1Manifest() datasource.DataSource {
	return &ApiextensionsCrossplaneIoCompositionRevisionV1Manifest{}
}

type ApiextensionsCrossplaneIoCompositionRevisionV1Manifest struct{}

type ApiextensionsCrossplaneIoCompositionRevisionV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		CompositeTypeRef *struct {
			ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
		} `tfsdk:"composite_type_ref" json:"compositeTypeRef,omitempty"`
		Mode     *string `tfsdk:"mode" json:"mode,omitempty"`
		Pipeline *[]struct {
			Credentials *[]struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				SecretRef *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
				Source *string `tfsdk:"source" json:"source,omitempty"`
			} `tfsdk:"credentials" json:"credentials,omitempty"`
			FunctionRef *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"function_ref" json:"functionRef,omitempty"`
			Input        *map[string]string `tfsdk:"input" json:"input,omitempty"`
			Requirements *struct {
				RequiredResources *[]struct {
					ApiVersion      *string            `tfsdk:"api_version" json:"apiVersion,omitempty"`
					Kind            *string            `tfsdk:"kind" json:"kind,omitempty"`
					MatchLabels     *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					Name            *string            `tfsdk:"name" json:"name,omitempty"`
					Namespace       *string            `tfsdk:"namespace" json:"namespace,omitempty"`
					RequirementName *string            `tfsdk:"requirement_name" json:"requirementName,omitempty"`
				} `tfsdk:"required_resources" json:"requiredResources,omitempty"`
			} `tfsdk:"requirements" json:"requirements,omitempty"`
			Step *string `tfsdk:"step" json:"step,omitempty"`
		} `tfsdk:"pipeline" json:"pipeline,omitempty"`
		Revision                          *int64  `tfsdk:"revision" json:"revision,omitempty"`
		WriteConnectionSecretsToNamespace *string `tfsdk:"write_connection_secrets_to_namespace" json:"writeConnectionSecretsToNamespace,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ApiextensionsCrossplaneIoCompositionRevisionV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apiextensions_crossplane_io_composition_revision_v1_manifest"
}

func (r *ApiextensionsCrossplaneIoCompositionRevisionV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "A CompositionRevision represents a revision of a Composition. Crossplane creates new revisions when there are changes to the Composition. Crossplane creates and manages CompositionRevisions. Don't directly edit CompositionRevisions.",
		MarkdownDescription: "A CompositionRevision represents a revision of a Composition. Crossplane creates new revisions when there are changes to the Composition. Crossplane creates and manages CompositionRevisions. Don't directly edit CompositionRevisions.",
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
				Description:         "CompositionRevisionSpec specifies the desired state of the composition revision.",
				MarkdownDescription: "CompositionRevisionSpec specifies the desired state of the composition revision.",
				Attributes: map[string]schema.Attribute{
					"composite_type_ref": schema.SingleNestedAttribute{
						Description:         "CompositeTypeRef specifies the type of composite resource that this composition is compatible with.",
						MarkdownDescription: "CompositeTypeRef specifies the type of composite resource that this composition is compatible with.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "APIVersion of the type.",
								MarkdownDescription: "APIVersion of the type.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the type.",
								MarkdownDescription: "Kind of the type.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"mode": schema.StringAttribute{
						Description:         "Mode controls what type or 'mode' of Composition will be used. 'Pipeline' indicates that a Composition specifies a pipeline of functions, each of which is responsible for producing composed resources that Crossplane should create or update.",
						MarkdownDescription: "Mode controls what type or 'mode' of Composition will be used. 'Pipeline' indicates that a Composition specifies a pipeline of functions, each of which is responsible for producing composed resources that Crossplane should create or update.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Pipeline"),
						},
					},

					"pipeline": schema.ListNestedAttribute{
						Description:         "Pipeline is a list of function steps that will be used when a composite resource referring to this composition is created. The Pipeline is only used by the 'Pipeline' mode of Composition. It is ignored by other modes.",
						MarkdownDescription: "Pipeline is a list of function steps that will be used when a composite resource referring to this composition is created. The Pipeline is only used by the 'Pipeline' mode of Composition. It is ignored by other modes.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"credentials": schema.ListNestedAttribute{
									Description:         "Credentials are optional credentials that the function needs.",
									MarkdownDescription: "Credentials are optional credentials that the function needs.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of this set of credentials.",
												MarkdownDescription: "Name of this set of credentials.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"secret_ref": schema.SingleNestedAttribute{
												Description:         "A SecretRef is a reference to a secret containing credentials that should be supplied to the function.",
												MarkdownDescription: "A SecretRef is a reference to a secret containing credentials that should be supplied to the function.",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Name of the secret.",
														MarkdownDescription: "Name of the secret.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the secret.",
														MarkdownDescription: "Namespace of the secret.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"source": schema.StringAttribute{
												Description:         "Source of the function credentials.",
												MarkdownDescription: "Source of the function credentials.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("None", "Secret"),
												},
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"function_ref": schema.SingleNestedAttribute{
									Description:         "FunctionRef is a reference to the function this step should execute.",
									MarkdownDescription: "FunctionRef is a reference to the function this step should execute.",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of the referenced Function.",
											MarkdownDescription: "Name of the referenced Function.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"input": schema.MapAttribute{
									Description:         "Input is an optional, arbitrary Kubernetes resource (i.e. a resource with an apiVersion and kind) that will be passed to the function as the 'input' of its RunFunctionRequest.",
									MarkdownDescription: "Input is an optional, arbitrary Kubernetes resource (i.e. a resource with an apiVersion and kind) that will be passed to the function as the 'input' of its RunFunctionRequest.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"requirements": schema.SingleNestedAttribute{
									Description:         "Requirements are resource requirements that will be satisfied before this pipeline step is called for the first time. This allows pre-populating required resources without requiring a function to request them first.",
									MarkdownDescription: "Requirements are resource requirements that will be satisfied before this pipeline step is called for the first time. This allows pre-populating required resources without requiring a function to request them first.",
									Attributes: map[string]schema.Attribute{
										"required_resources": schema.ListNestedAttribute{
											Description:         "RequiredResources is a list of resources that must be fetched before this function is called.",
											MarkdownDescription: "RequiredResources is a list of resources that must be fetched before this function is called.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"api_version": schema.StringAttribute{
														Description:         "APIVersion of the required resource.",
														MarkdownDescription: "APIVersion of the required resource.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"kind": schema.StringAttribute{
														Description:         "Kind of the required resource.",
														MarkdownDescription: "Kind of the required resource.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"match_labels": schema.MapAttribute{
														Description:         "MatchLabels specifies the set of labels to match for finding the required resource. When specified, Name is ignored.",
														MarkdownDescription: "MatchLabels specifies the set of labels to match for finding the required resource. When specified, Name is ignored.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the required resource.",
														MarkdownDescription: "Name of the required resource.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the required resource if it is namespaced.",
														MarkdownDescription: "Namespace of the required resource if it is namespaced.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"requirement_name": schema.StringAttribute{
														Description:         "RequirementName is the unique name to identify this required resource in the Required Resources map in the function request.",
														MarkdownDescription: "RequirementName is the unique name to identify this required resource in the Required Resources map in the function request.",
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

								"step": schema.StringAttribute{
									Description:         "Step name. Must be unique within its Pipeline.",
									MarkdownDescription: "Step name. Must be unique within its Pipeline.",
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

					"revision": schema.Int64Attribute{
						Description:         "Revision number. Newer revisions have larger numbers. This number can change. When a Composition transitions from state A -> B -> A there will be only two CompositionRevisions. Crossplane will edit the original CompositionRevision to change its revision number from 0 to 2.",
						MarkdownDescription: "Revision number. Newer revisions have larger numbers. This number can change. When a Composition transitions from state A -> B -> A there will be only two CompositionRevisions. Crossplane will edit the original CompositionRevision to change its revision number from 0 to 2.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"write_connection_secrets_to_namespace": schema.StringAttribute{
						Description:         "WriteConnectionSecretsToNamespace specifies the namespace in which the connection secrets of composite resource dynamically provisioned using this composition will be created.",
						MarkdownDescription: "WriteConnectionSecretsToNamespace specifies the namespace in which the connection secrets of composite resource dynamically provisioned using this composition will be created.",
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

func (r *ApiextensionsCrossplaneIoCompositionRevisionV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apiextensions_crossplane_io_composition_revision_v1_manifest")

	var model ApiextensionsCrossplaneIoCompositionRevisionV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("apiextensions.crossplane.io/v1")
	model.Kind = pointer.String("CompositionRevision")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
