/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type CamelApacheOrgCamelCatalogV1Resource struct{}

var (
	_ resource.Resource = (*CamelApacheOrgCamelCatalogV1Resource)(nil)
)

type CamelApacheOrgCamelCatalogV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type CamelApacheOrgCamelCatalogV1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Artifacts *struct {
			ArtifactId *string `tfsdk:"artifact_id" yaml:"artifactId,omitempty"`

			Dataformats *[]string `tfsdk:"dataformats" yaml:"dataformats,omitempty"`

			Dependencies *[]struct {
				ArtifactId *string `tfsdk:"artifact_id" yaml:"artifactId,omitempty"`

				Exclusions *[]struct {
					ArtifactId *string `tfsdk:"artifact_id" yaml:"artifactId,omitempty"`

					GroupId *string `tfsdk:"group_id" yaml:"groupId,omitempty"`
				} `tfsdk:"exclusions" yaml:"exclusions,omitempty"`

				GroupId *string `tfsdk:"group_id" yaml:"groupId,omitempty"`

				Version *string `tfsdk:"version" yaml:"version,omitempty"`
			} `tfsdk:"dependencies" yaml:"dependencies,omitempty"`

			Exclusions *[]struct {
				ArtifactId *string `tfsdk:"artifact_id" yaml:"artifactId,omitempty"`

				GroupId *string `tfsdk:"group_id" yaml:"groupId,omitempty"`
			} `tfsdk:"exclusions" yaml:"exclusions,omitempty"`

			GroupId *string `tfsdk:"group_id" yaml:"groupId,omitempty"`

			JavaTypes *[]string `tfsdk:"java_types" yaml:"javaTypes,omitempty"`

			Languages *[]string `tfsdk:"languages" yaml:"languages,omitempty"`

			Schemes *[]struct {
				Consumer *struct {
					Dependencies *[]struct {
						ArtifactId *string `tfsdk:"artifact_id" yaml:"artifactId,omitempty"`

						Exclusions *[]struct {
							ArtifactId *string `tfsdk:"artifact_id" yaml:"artifactId,omitempty"`

							GroupId *string `tfsdk:"group_id" yaml:"groupId,omitempty"`
						} `tfsdk:"exclusions" yaml:"exclusions,omitempty"`

						GroupId *string `tfsdk:"group_id" yaml:"groupId,omitempty"`

						Version *string `tfsdk:"version" yaml:"version,omitempty"`
					} `tfsdk:"dependencies" yaml:"dependencies,omitempty"`
				} `tfsdk:"consumer" yaml:"consumer,omitempty"`

				Http *bool `tfsdk:"http" yaml:"http,omitempty"`

				Id *string `tfsdk:"id" yaml:"id,omitempty"`

				Passive *bool `tfsdk:"passive" yaml:"passive,omitempty"`

				Producer *struct {
					Dependencies *[]struct {
						ArtifactId *string `tfsdk:"artifact_id" yaml:"artifactId,omitempty"`

						Exclusions *[]struct {
							ArtifactId *string `tfsdk:"artifact_id" yaml:"artifactId,omitempty"`

							GroupId *string `tfsdk:"group_id" yaml:"groupId,omitempty"`
						} `tfsdk:"exclusions" yaml:"exclusions,omitempty"`

						GroupId *string `tfsdk:"group_id" yaml:"groupId,omitempty"`

						Version *string `tfsdk:"version" yaml:"version,omitempty"`
					} `tfsdk:"dependencies" yaml:"dependencies,omitempty"`
				} `tfsdk:"producer" yaml:"producer,omitempty"`
			} `tfsdk:"schemes" yaml:"schemes,omitempty"`

			Version *string `tfsdk:"version" yaml:"version,omitempty"`
		} `tfsdk:"artifacts" yaml:"artifacts,omitempty"`

		Loaders *struct {
			ArtifactId *string `tfsdk:"artifact_id" yaml:"artifactId,omitempty"`

			Dependencies *[]struct {
				ArtifactId *string `tfsdk:"artifact_id" yaml:"artifactId,omitempty"`

				GroupId *string `tfsdk:"group_id" yaml:"groupId,omitempty"`

				Version *string `tfsdk:"version" yaml:"version,omitempty"`
			} `tfsdk:"dependencies" yaml:"dependencies,omitempty"`

			GroupId *string `tfsdk:"group_id" yaml:"groupId,omitempty"`

			Languages *[]string `tfsdk:"languages" yaml:"languages,omitempty"`

			Metadata *map[string]string `tfsdk:"metadata" yaml:"metadata,omitempty"`

			Version *string `tfsdk:"version" yaml:"version,omitempty"`
		} `tfsdk:"loaders" yaml:"loaders,omitempty"`

		Runtime *struct {
			ApplicationClass *string `tfsdk:"application_class" yaml:"applicationClass,omitempty"`

			Capabilities *struct {
				Dependencies *[]struct {
					ArtifactId *string `tfsdk:"artifact_id" yaml:"artifactId,omitempty"`

					GroupId *string `tfsdk:"group_id" yaml:"groupId,omitempty"`

					Version *string `tfsdk:"version" yaml:"version,omitempty"`
				} `tfsdk:"dependencies" yaml:"dependencies,omitempty"`

				Metadata *map[string]string `tfsdk:"metadata" yaml:"metadata,omitempty"`
			} `tfsdk:"capabilities" yaml:"capabilities,omitempty"`

			Dependencies *[]struct {
				ArtifactId *string `tfsdk:"artifact_id" yaml:"artifactId,omitempty"`

				GroupId *string `tfsdk:"group_id" yaml:"groupId,omitempty"`

				Version *string `tfsdk:"version" yaml:"version,omitempty"`
			} `tfsdk:"dependencies" yaml:"dependencies,omitempty"`

			Metadata *map[string]string `tfsdk:"metadata" yaml:"metadata,omitempty"`

			Provider *string `tfsdk:"provider" yaml:"provider,omitempty"`

			Version *string `tfsdk:"version" yaml:"version,omitempty"`
		} `tfsdk:"runtime" yaml:"runtime,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewCamelApacheOrgCamelCatalogV1Resource() resource.Resource {
	return &CamelApacheOrgCamelCatalogV1Resource{}
}

func (r *CamelApacheOrgCamelCatalogV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_camel_apache_org_camel_catalog_v1"
}

func (r *CamelApacheOrgCamelCatalogV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "CamelCatalog represents the languages, components, data formats and capabilities enabled on a given runtime provider. The catalog may be statically generated.",
		MarkdownDescription: "CamelCatalog represents the languages, components, data formats and capabilities enabled on a given runtime provider. The catalog may be statically generated.",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "the desired state of the catalog",
				MarkdownDescription: "the desired state of the catalog",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"artifacts": {
						Description:         "artifacts required by this catalog",
						MarkdownDescription: "artifacts required by this catalog",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"artifact_id": {
								Description:         "Maven Artifact",
								MarkdownDescription: "Maven Artifact",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"dataformats": {
								Description:         "accepted data formats",
								MarkdownDescription: "accepted data formats",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"dependencies": {
								Description:         "required dependencies",
								MarkdownDescription: "required dependencies",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"artifact_id": {
										Description:         "Maven Artifact",
										MarkdownDescription: "Maven Artifact",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"exclusions": {
										Description:         "provide a list of artifacts to exclude for this dependency",
										MarkdownDescription: "provide a list of artifacts to exclude for this dependency",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"artifact_id": {
												Description:         "Maven Artifact",
												MarkdownDescription: "Maven Artifact",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"group_id": {
												Description:         "Maven Group",
												MarkdownDescription: "Maven Group",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"group_id": {
										Description:         "Maven Group",
										MarkdownDescription: "Maven Group",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"version": {
										Description:         "Maven Version",
										MarkdownDescription: "Maven Version",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"exclusions": {
								Description:         "provide a list of artifacts to exclude for this dependency",
								MarkdownDescription: "provide a list of artifacts to exclude for this dependency",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"artifact_id": {
										Description:         "Maven Artifact",
										MarkdownDescription: "Maven Artifact",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"group_id": {
										Description:         "Maven Group",
										MarkdownDescription: "Maven Group",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"group_id": {
								Description:         "Maven Group",
								MarkdownDescription: "Maven Group",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"java_types": {
								Description:         "the Java types used by the artifact feature (ie, component, data format, ...)",
								MarkdownDescription: "the Java types used by the artifact feature (ie, component, data format, ...)",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"languages": {
								Description:         "accepted languages",
								MarkdownDescription: "accepted languages",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"schemes": {
								Description:         "accepted URI schemes",
								MarkdownDescription: "accepted URI schemes",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"consumer": {
										Description:         "required scope for consumer",
										MarkdownDescription: "required scope for consumer",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"dependencies": {
												Description:         "list of dependencies needed for this scope",
												MarkdownDescription: "list of dependencies needed for this scope",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"artifact_id": {
														Description:         "Maven Artifact",
														MarkdownDescription: "Maven Artifact",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"exclusions": {
														Description:         "provide a list of artifacts to exclude for this dependency",
														MarkdownDescription: "provide a list of artifacts to exclude for this dependency",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"artifact_id": {
																Description:         "Maven Artifact",
																MarkdownDescription: "Maven Artifact",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"group_id": {
																Description:         "Maven Group",
																MarkdownDescription: "Maven Group",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"group_id": {
														Description:         "Maven Group",
														MarkdownDescription: "Maven Group",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"version": {
														Description:         "Maven Version",
														MarkdownDescription: "Maven Version",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"http": {
										Description:         "is a HTTP based scheme",
										MarkdownDescription: "is a HTTP based scheme",

										Type: types.BoolType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"id": {
										Description:         "the ID (ie, timer in a timer:xyz URI)",
										MarkdownDescription: "the ID (ie, timer in a timer:xyz URI)",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"passive": {
										Description:         "is a passive scheme",
										MarkdownDescription: "is a passive scheme",

										Type: types.BoolType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"producer": {
										Description:         "required scope for producers",
										MarkdownDescription: "required scope for producers",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"dependencies": {
												Description:         "list of dependencies needed for this scope",
												MarkdownDescription: "list of dependencies needed for this scope",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"artifact_id": {
														Description:         "Maven Artifact",
														MarkdownDescription: "Maven Artifact",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"exclusions": {
														Description:         "provide a list of artifacts to exclude for this dependency",
														MarkdownDescription: "provide a list of artifacts to exclude for this dependency",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"artifact_id": {
																Description:         "Maven Artifact",
																MarkdownDescription: "Maven Artifact",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"group_id": {
																Description:         "Maven Group",
																MarkdownDescription: "Maven Group",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"group_id": {
														Description:         "Maven Group",
														MarkdownDescription: "Maven Group",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"version": {
														Description:         "Maven Version",
														MarkdownDescription: "Maven Version",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"version": {
								Description:         "Maven Version",
								MarkdownDescription: "Maven Version",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"loaders": {
						Description:         "loaders required by this catalog",
						MarkdownDescription: "loaders required by this catalog",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"artifact_id": {
								Description:         "Maven Artifact",
								MarkdownDescription: "Maven Artifact",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"dependencies": {
								Description:         "a list of additional dependencies required beside the base one",
								MarkdownDescription: "a list of additional dependencies required beside the base one",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"artifact_id": {
										Description:         "Maven Artifact",
										MarkdownDescription: "Maven Artifact",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"group_id": {
										Description:         "Maven Group",
										MarkdownDescription: "Maven Group",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"version": {
										Description:         "Maven Version",
										MarkdownDescription: "Maven Version",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"group_id": {
								Description:         "Maven Group",
								MarkdownDescription: "Maven Group",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"languages": {
								Description:         "a list of DSLs supported",
								MarkdownDescription: "a list of DSLs supported",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"metadata": {
								Description:         "Deprecated: never used a set of general metadata for various purposes",
								MarkdownDescription: "Deprecated: never used a set of general metadata for various purposes",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"version": {
								Description:         "Maven Version",
								MarkdownDescription: "Maven Version",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"runtime": {
						Description:         "the runtime targeted for the catalog",
						MarkdownDescription: "the runtime targeted for the catalog",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"application_class": {
								Description:         "application entry point (main) to be executed",
								MarkdownDescription: "application entry point (main) to be executed",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"capabilities": {
								Description:         "features offered by this runtime",
								MarkdownDescription: "features offered by this runtime",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"dependencies": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"artifact_id": {
												Description:         "Maven Artifact",
												MarkdownDescription: "Maven Artifact",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"group_id": {
												Description:         "Maven Group",
												MarkdownDescription: "Maven Group",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"version": {
												Description:         "Maven Version",
												MarkdownDescription: "Maven Version",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},

									"metadata": {
										Description:         "Deprecated: not in use",
										MarkdownDescription: "Deprecated: not in use",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"dependencies": {
								Description:         "list of dependencies needed to run the application",
								MarkdownDescription: "list of dependencies needed to run the application",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"artifact_id": {
										Description:         "Maven Artifact",
										MarkdownDescription: "Maven Artifact",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"group_id": {
										Description:         "Maven Group",
										MarkdownDescription: "Maven Group",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"version": {
										Description:         "Maven Version",
										MarkdownDescription: "Maven Version",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: true,
								Optional: false,
								Computed: false,
							},

							"metadata": {
								Description:         "set of metadata",
								MarkdownDescription: "set of metadata",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"provider": {
								Description:         "Camel main application provider, ie, Camel Quarkus",
								MarkdownDescription: "Camel main application provider, ie, Camel Quarkus",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"version": {
								Description:         "Camel K Runtime version",
								MarkdownDescription: "Camel K Runtime version",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *CamelApacheOrgCamelCatalogV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_camel_apache_org_camel_catalog_v1")

	var state CamelApacheOrgCamelCatalogV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CamelApacheOrgCamelCatalogV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("camel.apache.org/v1")
	goModel.Kind = utilities.Ptr("CamelCatalog")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CamelApacheOrgCamelCatalogV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_camel_apache_org_camel_catalog_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *CamelApacheOrgCamelCatalogV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_camel_apache_org_camel_catalog_v1")

	var state CamelApacheOrgCamelCatalogV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CamelApacheOrgCamelCatalogV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("camel.apache.org/v1")
	goModel.Kind = utilities.Ptr("CamelCatalog")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CamelApacheOrgCamelCatalogV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_camel_apache_org_camel_catalog_v1")
	// NO-OP: Terraform removes the state automatically for us
}
