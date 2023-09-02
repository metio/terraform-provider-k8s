/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package camel_apache_org_v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &CamelApacheOrgCamelCatalogV1DataSource{}
	_ datasource.DataSourceWithConfigure = &CamelApacheOrgCamelCatalogV1DataSource{}
)

func NewCamelApacheOrgCamelCatalogV1DataSource() datasource.DataSource {
	return &CamelApacheOrgCamelCatalogV1DataSource{}
}

type CamelApacheOrgCamelCatalogV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type CamelApacheOrgCamelCatalogV1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Artifacts *struct {
			ArtifactId   *string   `tfsdk:"artifact_id" json:"artifactId,omitempty"`
			Dataformats  *[]string `tfsdk:"dataformats" json:"dataformats,omitempty"`
			Dependencies *[]struct {
				ArtifactId *string `tfsdk:"artifact_id" json:"artifactId,omitempty"`
				Exclusions *[]struct {
					ArtifactId *string `tfsdk:"artifact_id" json:"artifactId,omitempty"`
					GroupId    *string `tfsdk:"group_id" json:"groupId,omitempty"`
				} `tfsdk:"exclusions" json:"exclusions,omitempty"`
				GroupId *string `tfsdk:"group_id" json:"groupId,omitempty"`
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
						Exclusions *[]struct {
							ArtifactId *string `tfsdk:"artifact_id" json:"artifactId,omitempty"`
							GroupId    *string `tfsdk:"group_id" json:"groupId,omitempty"`
						} `tfsdk:"exclusions" json:"exclusions,omitempty"`
						GroupId *string `tfsdk:"group_id" json:"groupId,omitempty"`
						Version *string `tfsdk:"version" json:"version,omitempty"`
					} `tfsdk:"dependencies" json:"dependencies,omitempty"`
				} `tfsdk:"consumer" json:"consumer,omitempty"`
				Http     *bool   `tfsdk:"http" json:"http,omitempty"`
				Id       *string `tfsdk:"id" json:"id,omitempty"`
				Passive  *bool   `tfsdk:"passive" json:"passive,omitempty"`
				Producer *struct {
					Dependencies *[]struct {
						ArtifactId *string `tfsdk:"artifact_id" json:"artifactId,omitempty"`
						Exclusions *[]struct {
							ArtifactId *string `tfsdk:"artifact_id" json:"artifactId,omitempty"`
							GroupId    *string `tfsdk:"group_id" json:"groupId,omitempty"`
						} `tfsdk:"exclusions" json:"exclusions,omitempty"`
						GroupId *string `tfsdk:"group_id" json:"groupId,omitempty"`
						Version *string `tfsdk:"version" json:"version,omitempty"`
					} `tfsdk:"dependencies" json:"dependencies,omitempty"`
				} `tfsdk:"producer" json:"producer,omitempty"`
			} `tfsdk:"schemes" json:"schemes,omitempty"`
			Version *string `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"artifacts" json:"artifacts,omitempty"`
		Loaders *struct {
			ArtifactId   *string `tfsdk:"artifact_id" json:"artifactId,omitempty"`
			Dependencies *[]struct {
				ArtifactId *string `tfsdk:"artifact_id" json:"artifactId,omitempty"`
				GroupId    *string `tfsdk:"group_id" json:"groupId,omitempty"`
				Version    *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"dependencies" json:"dependencies,omitempty"`
			GroupId   *string            `tfsdk:"group_id" json:"groupId,omitempty"`
			Languages *[]string          `tfsdk:"languages" json:"languages,omitempty"`
			Metadata  *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
			Version   *string            `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"loaders" json:"loaders,omitempty"`
		Runtime *struct {
			ApplicationClass *string `tfsdk:"application_class" json:"applicationClass,omitempty"`
			Capabilities     *struct {
				Dependencies *[]struct {
					ArtifactId *string `tfsdk:"artifact_id" json:"artifactId,omitempty"`
					GroupId    *string `tfsdk:"group_id" json:"groupId,omitempty"`
					Version    *string `tfsdk:"version" json:"version,omitempty"`
				} `tfsdk:"dependencies" json:"dependencies,omitempty"`
			} `tfsdk:"capabilities" json:"capabilities,omitempty"`
			Dependencies *[]struct {
				ArtifactId *string `tfsdk:"artifact_id" json:"artifactId,omitempty"`
				GroupId    *string `tfsdk:"group_id" json:"groupId,omitempty"`
				Version    *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"dependencies" json:"dependencies,omitempty"`
			Metadata *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
			Provider *string            `tfsdk:"provider" json:"provider,omitempty"`
			Version  *string            `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"runtime" json:"runtime,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CamelApacheOrgCamelCatalogV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_camel_apache_org_camel_catalog_v1"
}

func (r *CamelApacheOrgCamelCatalogV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
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
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"dataformats": schema.ListAttribute{
								Description:         "accepted data formats",
								MarkdownDescription: "accepted data formats",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"dependencies": schema.ListNestedAttribute{
								Description:         "required dependencies",
								MarkdownDescription: "required dependencies",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"artifact_id": schema.StringAttribute{
											Description:         "Maven Artifact",
											MarkdownDescription: "Maven Artifact",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"exclusions": schema.ListNestedAttribute{
											Description:         "provide a list of artifacts to exclude for this dependency",
											MarkdownDescription: "provide a list of artifacts to exclude for this dependency",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"artifact_id": schema.StringAttribute{
														Description:         "Maven Artifact",
														MarkdownDescription: "Maven Artifact",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"group_id": schema.StringAttribute{
														Description:         "Maven Group",
														MarkdownDescription: "Maven Group",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"group_id": schema.StringAttribute{
											Description:         "Maven Group",
											MarkdownDescription: "Maven Group",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"version": schema.StringAttribute{
											Description:         "Maven Version",
											MarkdownDescription: "Maven Version",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"exclusions": schema.ListNestedAttribute{
								Description:         "provide a list of artifacts to exclude for this dependency",
								MarkdownDescription: "provide a list of artifacts to exclude for this dependency",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"artifact_id": schema.StringAttribute{
											Description:         "Maven Artifact",
											MarkdownDescription: "Maven Artifact",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"group_id": schema.StringAttribute{
											Description:         "Maven Group",
											MarkdownDescription: "Maven Group",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"group_id": schema.StringAttribute{
								Description:         "Maven Group",
								MarkdownDescription: "Maven Group",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"java_types": schema.ListAttribute{
								Description:         "the Java types used by the artifact feature (ie, component, data format, ...)",
								MarkdownDescription: "the Java types used by the artifact feature (ie, component, data format, ...)",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"languages": schema.ListAttribute{
								Description:         "accepted languages",
								MarkdownDescription: "accepted languages",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"exclusions": schema.ListNestedAttribute{
																Description:         "provide a list of artifacts to exclude for this dependency",
																MarkdownDescription: "provide a list of artifacts to exclude for this dependency",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"artifact_id": schema.StringAttribute{
																			Description:         "Maven Artifact",
																			MarkdownDescription: "Maven Artifact",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"group_id": schema.StringAttribute{
																			Description:         "Maven Group",
																			MarkdownDescription: "Maven Group",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"group_id": schema.StringAttribute{
																Description:         "Maven Group",
																MarkdownDescription: "Maven Group",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"version": schema.StringAttribute{
																Description:         "Maven Version",
																MarkdownDescription: "Maven Version",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"http": schema.BoolAttribute{
											Description:         "is a HTTP based scheme",
											MarkdownDescription: "is a HTTP based scheme",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"id": schema.StringAttribute{
											Description:         "the ID (ie, timer in a timer:xyz URI)",
											MarkdownDescription: "the ID (ie, timer in a timer:xyz URI)",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"passive": schema.BoolAttribute{
											Description:         "is a passive scheme",
											MarkdownDescription: "is a passive scheme",
											Required:            false,
											Optional:            false,
											Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"exclusions": schema.ListNestedAttribute{
																Description:         "provide a list of artifacts to exclude for this dependency",
																MarkdownDescription: "provide a list of artifacts to exclude for this dependency",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"artifact_id": schema.StringAttribute{
																			Description:         "Maven Artifact",
																			MarkdownDescription: "Maven Artifact",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"group_id": schema.StringAttribute{
																			Description:         "Maven Group",
																			MarkdownDescription: "Maven Group",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"group_id": schema.StringAttribute{
																Description:         "Maven Group",
																MarkdownDescription: "Maven Group",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"version": schema.StringAttribute{
																Description:         "Maven Version",
																MarkdownDescription: "Maven Version",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"version": schema.StringAttribute{
								Description:         "Maven Version",
								MarkdownDescription: "Maven Version",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"loaders": schema.SingleNestedAttribute{
						Description:         "loaders required by this catalog",
						MarkdownDescription: "loaders required by this catalog",
						Attributes: map[string]schema.Attribute{
							"artifact_id": schema.StringAttribute{
								Description:         "Maven Artifact",
								MarkdownDescription: "Maven Artifact",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"dependencies": schema.ListNestedAttribute{
								Description:         "a list of additional dependencies required beside the base one",
								MarkdownDescription: "a list of additional dependencies required beside the base one",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"artifact_id": schema.StringAttribute{
											Description:         "Maven Artifact",
											MarkdownDescription: "Maven Artifact",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"group_id": schema.StringAttribute{
											Description:         "Maven Group",
											MarkdownDescription: "Maven Group",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"version": schema.StringAttribute{
											Description:         "Maven Version",
											MarkdownDescription: "Maven Version",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"group_id": schema.StringAttribute{
								Description:         "Maven Group",
								MarkdownDescription: "Maven Group",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"languages": schema.ListAttribute{
								Description:         "a list of DSLs supported",
								MarkdownDescription: "a list of DSLs supported",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"metadata": schema.MapAttribute{
								Description:         "the metadata of the loader",
								MarkdownDescription: "the metadata of the loader",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"version": schema.StringAttribute{
								Description:         "Maven Version",
								MarkdownDescription: "Maven Version",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"runtime": schema.SingleNestedAttribute{
						Description:         "the runtime targeted for the catalog",
						MarkdownDescription: "the runtime targeted for the catalog",
						Attributes: map[string]schema.Attribute{
							"application_class": schema.StringAttribute{
								Description:         "application entry point (main) to be executed",
								MarkdownDescription: "application entry point (main) to be executed",
								Required:            false,
								Optional:            false,
								Computed:            true,
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
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"group_id": schema.StringAttribute{
													Description:         "Maven Group",
													MarkdownDescription: "Maven Group",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"version": schema.StringAttribute{
													Description:         "Maven Version",
													MarkdownDescription: "Maven Version",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"dependencies": schema.ListNestedAttribute{
								Description:         "list of dependencies needed to run the application",
								MarkdownDescription: "list of dependencies needed to run the application",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"artifact_id": schema.StringAttribute{
											Description:         "Maven Artifact",
											MarkdownDescription: "Maven Artifact",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"group_id": schema.StringAttribute{
											Description:         "Maven Group",
											MarkdownDescription: "Maven Group",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"version": schema.StringAttribute{
											Description:         "Maven Version",
											MarkdownDescription: "Maven Version",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"metadata": schema.MapAttribute{
								Description:         "set of metadata",
								MarkdownDescription: "set of metadata",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"provider": schema.StringAttribute{
								Description:         "Camel main application provider, ie, Camel Quarkus",
								MarkdownDescription: "Camel main application provider, ie, Camel Quarkus",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"version": schema.StringAttribute{
								Description:         "Camel K Runtime version",
								MarkdownDescription: "Camel K Runtime version",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *CamelApacheOrgCamelCatalogV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *CamelApacheOrgCamelCatalogV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_camel_apache_org_camel_catalog_v1")

	var data CamelApacheOrgCamelCatalogV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "camel.apache.org", Version: "v1", Resource: "CamelCatalog"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse CamelApacheOrgCamelCatalogV1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("camel.apache.org/v1")
	data.Kind = pointer.String("CamelCatalog")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
