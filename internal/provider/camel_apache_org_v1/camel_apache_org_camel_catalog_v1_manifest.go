/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package camel_apache_org_v1

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
	_ datasource.DataSource = &CamelApacheOrgCamelCatalogV1Manifest{}
)

func NewCamelApacheOrgCamelCatalogV1Manifest() datasource.DataSource {
	return &CamelApacheOrgCamelCatalogV1Manifest{}
}

type CamelApacheOrgCamelCatalogV1Manifest struct{}

type CamelApacheOrgCamelCatalogV1ManifestData struct {
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
		Artifacts *struct {
			ArtifactId   *string   `tfsdk:"artifact_id" json:"artifactId,omitempty"`
			Classifier   *string   `tfsdk:"classifier" json:"classifier,omitempty"`
			Dataformats  *[]string `tfsdk:"dataformats" json:"dataformats,omitempty"`
			Dependencies *[]struct {
				ArtifactId *string `tfsdk:"artifact_id" json:"artifactId,omitempty"`
				Classifier *string `tfsdk:"classifier" json:"classifier,omitempty"`
				Exclusions *[]struct {
					ArtifactId *string `tfsdk:"artifact_id" json:"artifactId,omitempty"`
					GroupId    *string `tfsdk:"group_id" json:"groupId,omitempty"`
				} `tfsdk:"exclusions" json:"exclusions,omitempty"`
				GroupId *string `tfsdk:"group_id" json:"groupId,omitempty"`
				Type    *string `tfsdk:"type" json:"type,omitempty"`
				Version *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"dependencies" json:"dependencies,omitempty"`
			Exclusions *[]struct {
				ArtifactId *string `tfsdk:"artifact_id" json:"artifactId,omitempty"`
				GroupId    *string `tfsdk:"group_id" json:"groupId,omitempty"`
			} `tfsdk:"exclusions" json:"exclusions,omitempty"`
			GroupId   *string   `tfsdk:"group_id" json:"groupId,omitempty"`
			JavaTypes *[]string `tfsdk:"java_types" json:"javaTypes,omitempty"`
			Languages *[]string `tfsdk:"languages" json:"languages,omitempty"`
			Schemes   *[]struct {
				Consumer *struct {
					Dependencies *[]struct {
						ArtifactId *string `tfsdk:"artifact_id" json:"artifactId,omitempty"`
						Classifier *string `tfsdk:"classifier" json:"classifier,omitempty"`
						Exclusions *[]struct {
							ArtifactId *string `tfsdk:"artifact_id" json:"artifactId,omitempty"`
							GroupId    *string `tfsdk:"group_id" json:"groupId,omitempty"`
						} `tfsdk:"exclusions" json:"exclusions,omitempty"`
						GroupId *string `tfsdk:"group_id" json:"groupId,omitempty"`
						Type    *string `tfsdk:"type" json:"type,omitempty"`
						Version *string `tfsdk:"version" json:"version,omitempty"`
					} `tfsdk:"dependencies" json:"dependencies,omitempty"`
				} `tfsdk:"consumer" json:"consumer,omitempty"`
				Http     *bool   `tfsdk:"http" json:"http,omitempty"`
				Id       *string `tfsdk:"id" json:"id,omitempty"`
				Passive  *bool   `tfsdk:"passive" json:"passive,omitempty"`
				Producer *struct {
					Dependencies *[]struct {
						ArtifactId *string `tfsdk:"artifact_id" json:"artifactId,omitempty"`
						Classifier *string `tfsdk:"classifier" json:"classifier,omitempty"`
						Exclusions *[]struct {
							ArtifactId *string `tfsdk:"artifact_id" json:"artifactId,omitempty"`
							GroupId    *string `tfsdk:"group_id" json:"groupId,omitempty"`
						} `tfsdk:"exclusions" json:"exclusions,omitempty"`
						GroupId *string `tfsdk:"group_id" json:"groupId,omitempty"`
						Type    *string `tfsdk:"type" json:"type,omitempty"`
						Version *string `tfsdk:"version" json:"version,omitempty"`
					} `tfsdk:"dependencies" json:"dependencies,omitempty"`
				} `tfsdk:"producer" json:"producer,omitempty"`
			} `tfsdk:"schemes" json:"schemes,omitempty"`
			Type    *string `tfsdk:"type" json:"type,omitempty"`
			Version *string `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"artifacts" json:"artifacts,omitempty"`
		Loaders *struct {
			ArtifactId   *string `tfsdk:"artifact_id" json:"artifactId,omitempty"`
			Classifier   *string `tfsdk:"classifier" json:"classifier,omitempty"`
			Dependencies *[]struct {
				ArtifactId *string `tfsdk:"artifact_id" json:"artifactId,omitempty"`
				Classifier *string `tfsdk:"classifier" json:"classifier,omitempty"`
				GroupId    *string `tfsdk:"group_id" json:"groupId,omitempty"`
				Type       *string `tfsdk:"type" json:"type,omitempty"`
				Version    *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"dependencies" json:"dependencies,omitempty"`
			GroupId   *string            `tfsdk:"group_id" json:"groupId,omitempty"`
			Languages *[]string          `tfsdk:"languages" json:"languages,omitempty"`
			Metadata  *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
			Type      *string            `tfsdk:"type" json:"type,omitempty"`
			Version   *string            `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"loaders" json:"loaders,omitempty"`
		Runtime *struct {
			ApplicationClass *string `tfsdk:"application_class" json:"applicationClass,omitempty"`
			Capabilities     *struct {
				Dependencies *[]struct {
					ArtifactId *string `tfsdk:"artifact_id" json:"artifactId,omitempty"`
					Classifier *string `tfsdk:"classifier" json:"classifier,omitempty"`
					GroupId    *string `tfsdk:"group_id" json:"groupId,omitempty"`
					Type       *string `tfsdk:"type" json:"type,omitempty"`
					Version    *string `tfsdk:"version" json:"version,omitempty"`
				} `tfsdk:"dependencies" json:"dependencies,omitempty"`
			} `tfsdk:"capabilities" json:"capabilities,omitempty"`
			Dependencies *[]struct {
				ArtifactId *string `tfsdk:"artifact_id" json:"artifactId,omitempty"`
				Classifier *string `tfsdk:"classifier" json:"classifier,omitempty"`
				GroupId    *string `tfsdk:"group_id" json:"groupId,omitempty"`
				Type       *string `tfsdk:"type" json:"type,omitempty"`
				Version    *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"dependencies" json:"dependencies,omitempty"`
			Metadata *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
			Provider *string            `tfsdk:"provider" json:"provider,omitempty"`
			Version  *string            `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"runtime" json:"runtime,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CamelApacheOrgCamelCatalogV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_camel_apache_org_camel_catalog_v1_manifest"
}

func (r *CamelApacheOrgCamelCatalogV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CamelCatalog represents the languages, components, data formats and capabilities enabled on a given runtime provider. The catalog may be statically generated.",
		MarkdownDescription: "CamelCatalog represents the languages, components, data formats and capabilities enabled on a given runtime provider. The catalog may be statically generated.",
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
				Description:         "the desired state of the catalog",
				MarkdownDescription: "the desired state of the catalog",
				Attributes: map[string]schema.Attribute{
					"artifacts": schema.SingleNestedAttribute{
						Description:         "artifacts required by this catalog",
						MarkdownDescription: "artifacts required by this catalog",
						Attributes: map[string]schema.Attribute{
							"artifact_id": schema.StringAttribute{
								Description:         "Maven Artifact",
								MarkdownDescription: "Maven Artifact",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"classifier": schema.StringAttribute{
								Description:         "Maven Classifier",
								MarkdownDescription: "Maven Classifier",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dataformats": schema.ListAttribute{
								Description:         "accepted data formats",
								MarkdownDescription: "accepted data formats",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dependencies": schema.ListNestedAttribute{
								Description:         "required dependencies",
								MarkdownDescription: "required dependencies",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"artifact_id": schema.StringAttribute{
											Description:         "Maven Artifact",
											MarkdownDescription: "Maven Artifact",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"classifier": schema.StringAttribute{
											Description:         "Maven Classifier",
											MarkdownDescription: "Maven Classifier",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclusions": schema.ListNestedAttribute{
											Description:         "provide a list of artifacts to exclude for this dependency",
											MarkdownDescription: "provide a list of artifacts to exclude for this dependency",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"artifact_id": schema.StringAttribute{
														Description:         "Maven Artifact",
														MarkdownDescription: "Maven Artifact",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"group_id": schema.StringAttribute{
														Description:         "Maven Group",
														MarkdownDescription: "Maven Group",
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

										"group_id": schema.StringAttribute{
											Description:         "Maven Group",
											MarkdownDescription: "Maven Group",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"type": schema.StringAttribute{
											Description:         "Maven Type",
											MarkdownDescription: "Maven Type",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"version": schema.StringAttribute{
											Description:         "Maven Version",
											MarkdownDescription: "Maven Version",
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

							"exclusions": schema.ListNestedAttribute{
								Description:         "provide a list of artifacts to exclude for this dependency",
								MarkdownDescription: "provide a list of artifacts to exclude for this dependency",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"artifact_id": schema.StringAttribute{
											Description:         "Maven Artifact",
											MarkdownDescription: "Maven Artifact",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"group_id": schema.StringAttribute{
											Description:         "Maven Group",
											MarkdownDescription: "Maven Group",
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

							"group_id": schema.StringAttribute{
								Description:         "Maven Group",
								MarkdownDescription: "Maven Group",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"java_types": schema.ListAttribute{
								Description:         "the Java types used by the artifact feature (ie, component, data format, ...)",
								MarkdownDescription: "the Java types used by the artifact feature (ie, component, data format, ...)",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"languages": schema.ListAttribute{
								Description:         "accepted languages",
								MarkdownDescription: "accepted languages",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"schemes": schema.ListNestedAttribute{
								Description:         "accepted URI schemes",
								MarkdownDescription: "accepted URI schemes",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"consumer": schema.SingleNestedAttribute{
											Description:         "required scope for consumer",
											MarkdownDescription: "required scope for consumer",
											Attributes: map[string]schema.Attribute{
												"dependencies": schema.ListNestedAttribute{
													Description:         "list of dependencies needed for this scope",
													MarkdownDescription: "list of dependencies needed for this scope",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"artifact_id": schema.StringAttribute{
																Description:         "Maven Artifact",
																MarkdownDescription: "Maven Artifact",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"classifier": schema.StringAttribute{
																Description:         "Maven Classifier",
																MarkdownDescription: "Maven Classifier",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"exclusions": schema.ListNestedAttribute{
																Description:         "provide a list of artifacts to exclude for this dependency",
																MarkdownDescription: "provide a list of artifacts to exclude for this dependency",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"artifact_id": schema.StringAttribute{
																			Description:         "Maven Artifact",
																			MarkdownDescription: "Maven Artifact",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"group_id": schema.StringAttribute{
																			Description:         "Maven Group",
																			MarkdownDescription: "Maven Group",
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

															"group_id": schema.StringAttribute{
																Description:         "Maven Group",
																MarkdownDescription: "Maven Group",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"type": schema.StringAttribute{
																Description:         "Maven Type",
																MarkdownDescription: "Maven Type",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"version": schema.StringAttribute{
																Description:         "Maven Version",
																MarkdownDescription: "Maven Version",
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"http": schema.BoolAttribute{
											Description:         "is a HTTP based scheme",
											MarkdownDescription: "is a HTTP based scheme",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"id": schema.StringAttribute{
											Description:         "the ID (ie, timer in a timer:xyz URI)",
											MarkdownDescription: "the ID (ie, timer in a timer:xyz URI)",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"passive": schema.BoolAttribute{
											Description:         "is a passive scheme",
											MarkdownDescription: "is a passive scheme",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"producer": schema.SingleNestedAttribute{
											Description:         "required scope for producers",
											MarkdownDescription: "required scope for producers",
											Attributes: map[string]schema.Attribute{
												"dependencies": schema.ListNestedAttribute{
													Description:         "list of dependencies needed for this scope",
													MarkdownDescription: "list of dependencies needed for this scope",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"artifact_id": schema.StringAttribute{
																Description:         "Maven Artifact",
																MarkdownDescription: "Maven Artifact",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"classifier": schema.StringAttribute{
																Description:         "Maven Classifier",
																MarkdownDescription: "Maven Classifier",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"exclusions": schema.ListNestedAttribute{
																Description:         "provide a list of artifacts to exclude for this dependency",
																MarkdownDescription: "provide a list of artifacts to exclude for this dependency",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"artifact_id": schema.StringAttribute{
																			Description:         "Maven Artifact",
																			MarkdownDescription: "Maven Artifact",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"group_id": schema.StringAttribute{
																			Description:         "Maven Group",
																			MarkdownDescription: "Maven Group",
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

															"group_id": schema.StringAttribute{
																Description:         "Maven Group",
																MarkdownDescription: "Maven Group",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"type": schema.StringAttribute{
																Description:         "Maven Type",
																MarkdownDescription: "Maven Type",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"version": schema.StringAttribute{
																Description:         "Maven Version",
																MarkdownDescription: "Maven Version",
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": schema.StringAttribute{
								Description:         "Maven Type",
								MarkdownDescription: "Maven Type",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"version": schema.StringAttribute{
								Description:         "Maven Version",
								MarkdownDescription: "Maven Version",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"loaders": schema.SingleNestedAttribute{
						Description:         "loaders required by this catalog",
						MarkdownDescription: "loaders required by this catalog",
						Attributes: map[string]schema.Attribute{
							"artifact_id": schema.StringAttribute{
								Description:         "Maven Artifact",
								MarkdownDescription: "Maven Artifact",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"classifier": schema.StringAttribute{
								Description:         "Maven Classifier",
								MarkdownDescription: "Maven Classifier",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dependencies": schema.ListNestedAttribute{
								Description:         "a list of additional dependencies required beside the base one",
								MarkdownDescription: "a list of additional dependencies required beside the base one",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"artifact_id": schema.StringAttribute{
											Description:         "Maven Artifact",
											MarkdownDescription: "Maven Artifact",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"classifier": schema.StringAttribute{
											Description:         "Maven Classifier",
											MarkdownDescription: "Maven Classifier",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"group_id": schema.StringAttribute{
											Description:         "Maven Group",
											MarkdownDescription: "Maven Group",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"type": schema.StringAttribute{
											Description:         "Maven Type",
											MarkdownDescription: "Maven Type",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"version": schema.StringAttribute{
											Description:         "Maven Version",
											MarkdownDescription: "Maven Version",
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

							"group_id": schema.StringAttribute{
								Description:         "Maven Group",
								MarkdownDescription: "Maven Group",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"languages": schema.ListAttribute{
								Description:         "a list of DSLs supported",
								MarkdownDescription: "a list of DSLs supported",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"metadata": schema.MapAttribute{
								Description:         "the metadata of the loader",
								MarkdownDescription: "the metadata of the loader",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "Maven Type",
								MarkdownDescription: "Maven Type",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"version": schema.StringAttribute{
								Description:         "Maven Version",
								MarkdownDescription: "Maven Version",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"runtime": schema.SingleNestedAttribute{
						Description:         "the runtime targeted for the catalog",
						MarkdownDescription: "the runtime targeted for the catalog",
						Attributes: map[string]schema.Attribute{
							"application_class": schema.StringAttribute{
								Description:         "application entry point (main) to be executed",
								MarkdownDescription: "application entry point (main) to be executed",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"capabilities": schema.SingleNestedAttribute{
								Description:         "features offered by this runtime",
								MarkdownDescription: "features offered by this runtime",
								Attributes: map[string]schema.Attribute{
									"dependencies": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"artifact_id": schema.StringAttribute{
													Description:         "Maven Artifact",
													MarkdownDescription: "Maven Artifact",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"classifier": schema.StringAttribute{
													Description:         "Maven Classifier",
													MarkdownDescription: "Maven Classifier",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"group_id": schema.StringAttribute{
													Description:         "Maven Group",
													MarkdownDescription: "Maven Group",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"type": schema.StringAttribute{
													Description:         "Maven Type",
													MarkdownDescription: "Maven Type",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"version": schema.StringAttribute{
													Description:         "Maven Version",
													MarkdownDescription: "Maven Version",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"dependencies": schema.ListNestedAttribute{
								Description:         "list of dependencies needed to run the application",
								MarkdownDescription: "list of dependencies needed to run the application",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"artifact_id": schema.StringAttribute{
											Description:         "Maven Artifact",
											MarkdownDescription: "Maven Artifact",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"classifier": schema.StringAttribute{
											Description:         "Maven Classifier",
											MarkdownDescription: "Maven Classifier",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"group_id": schema.StringAttribute{
											Description:         "Maven Group",
											MarkdownDescription: "Maven Group",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"type": schema.StringAttribute{
											Description:         "Maven Type",
											MarkdownDescription: "Maven Type",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"version": schema.StringAttribute{
											Description:         "Maven Version",
											MarkdownDescription: "Maven Version",
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

							"metadata": schema.MapAttribute{
								Description:         "set of metadata",
								MarkdownDescription: "set of metadata",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"provider": schema.StringAttribute{
								Description:         "Camel main application provider, ie, Camel Quarkus",
								MarkdownDescription: "Camel main application provider, ie, Camel Quarkus",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"version": schema.StringAttribute{
								Description:         "Camel K Runtime version",
								MarkdownDescription: "Camel K Runtime version",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
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

func (r *CamelApacheOrgCamelCatalogV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_camel_apache_org_camel_catalog_v1_manifest")

	var model CamelApacheOrgCamelCatalogV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("camel.apache.org/v1")
	model.Kind = pointer.String("CamelCatalog")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
